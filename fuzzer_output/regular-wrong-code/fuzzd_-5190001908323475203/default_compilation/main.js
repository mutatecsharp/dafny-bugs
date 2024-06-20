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
      return _module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sppkdbw")).length),_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-658),_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), (_module.D23.create_DC66(true, false, new BigNumber((_dafny.Seq.UnicodeFromString("tv")).length), new _dafny.CodePoint('n'.codePointAt(0)))).dtor_cf115, new _dafny.CodePoint('y'.codePointAt(0)))))).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("av"), _dafny.Seq.UnicodeFromString("iysixoa"))).length));
    };
    static fm1(p0, p1, p2, globalState) {
      return ((true) && (false)) && (!_dafny.Seq.contains(_dafny.Seq.of(true, true, true, !(false)), false));
    };
    static fm2(p0, p1, globalState) {
      return _dafny.Set.fromElements((new BigNumber(-694)).minus(new BigNumber(827)), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,false))).length));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("dkry")), (_module.D29.create_DC80(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("hhavv"), _dafny.Seq.UnicodeFromString("anqfsyd")))).dtor_cf137);
    };
    static fm4(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_0_i0) {
        return new BigNumber(940);
      }),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(559), new BigNumber(10), new BigNumber(822), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(21), new BigNumber(743))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(21)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(743)))) {
            _coll0.push([(_1_v0).plus(new BigNumber(186)),new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)]);
          }
        }
        return _coll0;
      }()).length))).length)))).cardinality()))).cardinality()), new BigNumber(614)),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-27)), function (_2_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_module.D12.create_DC32(false)).dtor_cf55)).length);
      }),!(true)));
    };
    static fm7(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(264), (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("emmbf")).length))).cardinality())).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))).length))), (new BigNumber((_dafny.Seq.UnicodeFromString("mmkep")).length)).minus(new BigNumber(-486)), new BigNumber(520), (new BigNumber(712)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-864)))).length))));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true)));
    };
    static fm13(p0, p1, p2, globalState) {
      if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("jnrfyvd"), _dafny.Seq.UnicodeFromString("wtttsx"))) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC5(!(false), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-930))).cardinality())), _dafny.Seq.of(new BigNumber(260), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(515), new BigNumber(-320))) {
    let _3_v0 = _compr_1;
    if (((new BigNumber(515)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(-320)))) {
      _coll1.push([(_3_v0).multipliedBy(new BigNumber(17)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), function (_4_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length)]);
    }
  }
  return _coll1;
}()).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length)))), new BigNumber(86));
    };
    static fm18(globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm19(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(147),_module.D21.create_DC58(new BigNumber(462)))).length)))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),new _dafny.CodePoint('a'.codePointAt(0)))).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),new BigNumber(106))).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-636), new BigNumber(620))) {
          let _5_v0 = _compr_2;
          if (((new BigNumber(-636)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(620)))) {
            _coll2.push([(_5_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())),new BigNumber((function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of _dafny.IntegerRange(new BigNumber(759), new BigNumber(-371))) {
                let _6_v1 = _compr_3;
                if (((new BigNumber(759)).isLessThanOrEqualTo(_6_v1)) && ((_6_v1).isLessThan(new BigNumber(-371)))) {
                  _coll3.push([(_6_v1).multipliedBy(new BigNumber(-511)),false]);
                }
              }
              return _coll3;
            }()).length)]);
          }
        }
        return _coll2;
      }()));
    };
    static fm20(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _dafny.Seq.UnicodeFromString("dftv")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-393)), function (_7_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("dxmd")));
    };
    static fm21(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-961)), new BigNumber(-525), new BigNumber(796)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-308)), function (_8_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), function (_9_i1) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("fhati")).length);
        })).length));
      }), _dafny.Seq.of(new BigNumber(521), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ltildgwl")).length))).length))));
    };
    static fm22(globalState) {
      return function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (((false) ? (_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true), false))) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(false, (_module.D14.create_DC40(_module.D6.create_DC18(_module.D6.create_DC15(new _dafny.CodePoint('l'.codePointAt(0)))), false, true)).dtor_cf66))))).Elements) {
          let _10_v0 = _compr_4;
          if (_dafny.Seq.contains(((false) ? (_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true), false))) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(false, (_module.D14.create_DC40(_module.D6.create_DC18(_module.D6.create_DC15(new _dafny.CodePoint('l'.codePointAt(0)))), false, true)).dtor_cf66)))), _10_v0)) {
            _coll4.add(_10_v0);
          }
        }
        return _coll4;
      }();
    };
    static fm23(globalState) {
      return _module.D1.create_DC5(!(!(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(673)), function (_11_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length)).isEqualTo(new BigNumber((_dafny.Seq.of(false)).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(567)), _dafny.Seq.of(new BigNumber(731), new BigNumber(368))), new BigNumber((_dafny.Seq.of(true, true, true)).length));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jcestcjpt")).length),true)).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length)), new BigNumber(860)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(572), new BigNumber(-2), new BigNumber(633)));
    };
    static fm26(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('q'.codePointAt(0));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(true, false)).Difference(_dafny.Set.fromElements(!(true)))).Difference(((true) ? (_dafny.Set.fromElements(!(!(!(false))))) : (_dafny.Set.fromElements(false))));
    };
    static fm28(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("elpv");
    };
    static fm29(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(594)).multipliedBy(new BigNumber(437)),(new BigNumber(-159)).plus(new BigNumber(716)));
    };
    static fm30(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).Intersect(_dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))));
    };
    static fm33(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(-809), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D23.create_DC66(true, true, new BigNumber(442), new _dafny.CodePoint('f'.codePointAt(0)))).dtor_cf112,new BigNumber(721))).length)).multipliedBy(new BigNumber(335)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)).minus(new BigNumber(881)), new BigNumber(889));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(true)));
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("phtihs"), _dafny.Seq.UnicodeFromString("aq")), _dafny.Seq.UnicodeFromString("vw"));
    };
    static fm36(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(453),!(false))).length))).length),!(false))).length), new BigNumber(137)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(367), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(995))).cardinality()))).length))));
    };
    static fm37(p0, p1, globalState) {
      if ((false) === (false)) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else if (!(true)) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }
    };
    static fm38(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(!(true))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_12_i0) {
        return _dafny.Set.fromElements(true, !(true));
      }));
    };
    static fm39(p0, p1, globalState) {
      return _module.D1.create_DC5(!(true), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(907)), function (_13_i0) {
  return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length));
})).length), (_module.D6.create_DC17(true, new BigNumber(-786), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))).dtor_cf32, new BigNumber(598)), new BigNumber(259));
    };
    static fm40(p0, p1, p2, globalState) {
      return _module.D9.create_DC25();
    };
    static fm41(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(!(false)), _dafny.Seq.of(!(true))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-967)), function (_14_i0) {
        return _dafny.Seq.of(false, false);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_15_i1) {
        return _dafny.Seq.of(false, true, true, false);
      })));
    };
    static fm42(p0, p1, globalState) {
      if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(true)), _dafny.MultiSet.fromElements(true, false, !(true)))) {
        if (!(true)) {
          return _module.D6.create_DC17(false, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-748), new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()))),true)).length))).length), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)));
        } else {
          return _module.D6.create_DC17(false, new BigNumber(409), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)));
        }
      } else {
        return _module.D6.create_DC17(!(true), new BigNumber(-333), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)));
      }
    };
    static fm43(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(800),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(993),!(false)))).length));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      if (((true) ? (true) : (false))) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)));
      } else {
        return _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)));
      }
    };
    static fm45(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(522)).plus(new BigNumber((_dafny.Seq.of(!(true), false)).length)),true);
    };
    static fm46(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false));
    };
    static fm47(p0, p1, globalState) {
      if ((new BigNumber(-838)).isLessThanOrEqualTo(new BigNumber(146))) {
        return _module.D6.create_DC18(_module.D6.create_DC17(!(false), new BigNumber(16), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))));
      } else {
        return _module.D6.create_DC18(_module.D6.create_DC16(true));
      }
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return _module.D12.create_DC29((_dafny.MultiSet.fromElements(new BigNumber(239))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(521))).length))));
    };
    static fm49(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(false, true)).Union(_dafny.MultiSet.fromElements(true));
    };
    static fm50(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of((_module.D1.create_DC5(!(false), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gr")).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(807))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("re")).length))).length), new BigNumber(-637), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(726))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-441))).length)), new BigNumber(323))).dtor_cf9), _dafny.Seq.of(false, true, false, true, false)));
    };
    static fm51(p0, p1, globalState) {
      return _dafny.Seq.of(_module.D0.create_DC0(new BigNumber(466)));
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return _module.D12.create_DC32(((false) ? (false) : (false)));
    };
    static fm53(p0, p1, globalState) {
      return _module.D2.create_DC8(false, _dafny.Seq.UnicodeFromString("tb"));
    };
    static fm54(globalState) {
      return _dafny.Seq.of((function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(14), new BigNumber(-272))) {
          let _16_v0 = _compr_5;
          if (((new BigNumber(14)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(-272)))) {
            _coll5.add((_16_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("tfbs")).length)));
          }
        }
        return _coll5;
      }()).Intersect(_dafny.Set.fromElements(new BigNumber(545))));
    };
    static fm55(p0, globalState) {
      return _module.D8.create_DC22((new BigNumber(-47)).minus(new BigNumber(897)), _module.D1.create_DC3(_dafny.Set.fromElements(true, false, !(!(false)), !(false))), _module.D0.create_DC0(new BigNumber(856)));
    };
    static fm56(globalState) {
      return _module.D0.create_DC0(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(400), new BigNumber((_dafny.Seq.UnicodeFromString("sa")).length))).Difference(_dafny.MultiSet.fromElements(new BigNumber(626)))).cardinality()));
    };
    static fm57(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber(-776))))).Merge((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_dafny.Seq.of(new BigNumber(918)))).Keys.Elements) {
          let _17_v0 = _compr_6;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_dafny.Seq.of(new BigNumber(918)))).contains(_17_v0)) {
            _coll6.push([_17_v0,_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(713))]);
          }
        }
        return _coll6;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(804)))));
    };
    static fm58(p0, p1, globalState) {
      return _module.D0.create_DC1(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(432))).length)),new BigNumber(324))).Merge(function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-502), new BigNumber(829))) {
    let _18_v0 = _compr_7;
    if (((new BigNumber(-502)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(829)))) {
      _coll7.push([(_18_v0).multipliedBy(new BigNumber(358)),new BigNumber(-191)]);
    }
  }
  return _coll7;
}())).length));
    };
    static fm59(globalState) {
      return _dafny.Seq.of(_module.D15.create_DC44(_module.D15.create_DC44(_module.D15.create_DC44(_module.D15.create_DC43(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(260)), function (_19_i0) {
  return new BigNumber(-529);
})).length))))), _module.D15.create_DC44(_module.D15.create_DC44(_module.D15.create_DC42(_dafny.MultiSet.fromElements(true, false, !(false))))), _module.D15.create_DC44(_module.D15.create_DC43(new BigNumber((_dafny.Seq.of(new BigNumber(686), (_module.D2.create_DC9(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(115),false)).length), new BigNumber(-426), _dafny.Seq.Create(_module.__default.abs(new BigNumber(103)), function (_20_i1) {
  return new _dafny.CodePoint('h'.codePointAt(0));
}), false, _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),false)).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_21_i2) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length)))).dtor_cf17)).length))), _module.D15.create_DC44(_module.D15.create_DC42(_dafny.MultiSet.fromElements(true))), _module.D15.create_DC44(_module.D15.create_DC44(_module.D15.create_DC43(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length)))).cardinality())))));
    };
    static fm60(globalState) {
      if (!((_dafny.Set.fromElements(new BigNumber(-364), new BigNumber(513))).IsProperSubsetOf(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-76)))))) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, !(true)),new BigNumber((_dafny.Set.fromElements(true)).length));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, !(false), true, true),new BigNumber(435));
      }
    };
    static fm61(p0, p1, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(896)), function (_22_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_23_i1) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length);
        });
      }), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-203)), function (_24_i2) {
        return new BigNumber(-910);
      }), _dafny.Seq.of(new BigNumber(533), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-670), new BigNumber(717))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-282),new BigNumber((function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-841)), function (_25_i3) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        })).length))).length),new BigNumber((function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(319), new BigNumber(45))) {
            let _26_v0 = _compr_9;
            if (((new BigNumber(319)).isLessThanOrEqualTo(_26_v0)) && ((_26_v0).isLessThan(new BigNumber(45)))) {
              _coll9.push([_module.__default.safeDivisionInt(_26_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_27_i4) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length)),new BigNumber((_dafny.Set.fromElements(true)).length)]);
            }
          }
          return _coll9;
        }()).length))).length),!(!(!(false))))).Keys.Elements) {
          let _28_v1 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-841)), function (_25_i3) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length))).length),new BigNumber((function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(319), new BigNumber(45))) {
              let _26_v0 = _compr_10;
              if (((new BigNumber(319)).isLessThanOrEqualTo(_26_v0)) && ((_26_v0).isLessThan(new BigNumber(45)))) {
                _coll10.push([_module.__default.safeDivisionInt(_26_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_27_i4) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length)),new BigNumber((_dafny.Set.fromElements(true)).length)]);
              }
            }
            return _coll10;
          }()).length))).length),!(!(!(false))))).contains(_28_v1)) {
            _coll8.add(_module.__default.safeDivisionInt(_28_v1, new BigNumber((_dafny.Seq.of(true)).length)));
          }
        }
        return _coll8;
      }()).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)))))).Intersect((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(722), new BigNumber((_dafny.Seq.of(new BigNumber(663))).length)))).Difference((_module.D30.create_DC82(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(817), new BigNumber(778)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(642)), function (_29_i5) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("neutv")).length);
})).length)), _dafny.Seq.of(new BigNumber(574), new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), function (_30_i6) {
  return new BigNumber((_dafny.Set.fromElements(new BigNumber(5), new BigNumber((_dafny.Seq.of(true)).length))).length);
})).length), new BigNumber(197)))))).dtor_cf141));
    };
    static fm62(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ldysnsn"), _dafny.Seq.UnicodeFromString("tddexouk"), _dafny.Seq.UnicodeFromString("cm"), _dafny.Seq.UnicodeFromString("yicgsp"))).Union((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("iuli"), _dafny.Seq.UnicodeFromString("wobeyci"))).Intersect(function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_31_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }))).Elements) {
          let _32_v0 = _compr_11;
          if ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_31_i0) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }))).contains(_32_v0)) {
            _coll11.add(_32_v0);
          }
        }
        return _coll11;
      }()));
    };
    static fm63(p0, p1, p2, globalState) {
      return _module.D17.create_DC48(!(!(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), function (_33_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length), new BigNumber((_dafny.Set.fromElements(!(true), !(false), false)).length))).cardinality()))).length)).isEqualTo(new BigNumber(-705))), (true) && (!(true)), _dafny.areEqual(_dafny.Seq.UnicodeFromString("wivlux"), _dafny.Seq.UnicodeFromString("nw")));
    };
    static fm64(globalState) {
      if (false) {
        return _module.D2.create_DC7(_dafny.Set.fromElements(new BigNumber(-133)));
      } else {
        return _module.D2.create_DC7(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("jalxahvpr"))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), function (_34_i0) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length), new BigNumber(570)));
      }
    };
    static m0(p0, p1, globalState) {
      let _35_v0;
      _35_v0 = true;
      let _36_i0;
      _36_i0 = _dafny.ZERO;
      L0: {
        while (_35_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_36_i0)) {
              break L0;
            }
            _36_i0 = (_36_i0).plus(_dafny.ONE);
            let _37_v1;
            _37_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kvrx")).length),_35_v0);
            if ((_35_v0) || ((((((_37_v1).contains(p0)) ? ((_37_v1).get(p0)) : (_35_v0))) ? (_35_v0) : (_35_v0)))) {
              let _38_v2;
              _38_v2 = new BigNumber(519);
              _38_v2 = _module.__default.safeModuloInt((_38_v2).minus(_38_v2), p0);
              _38_v2 = _module.__default.safeModuloInt(_module.__default.fm0(globalState), new BigNumber((_37_v1).length));
              let _39_v3;
              let _nw0 = new _module.C3();
              _nw0.__ctor(_38_v2, _35_v0);
              _39_v3 = _nw0;
              let _40_v4;
              let _nw1 = new _module.C5();
              _nw1.__ctor(false);
              _40_v4 = _nw1;
              let _41_v5;
              _41_v5 = _dafny.Map.Empty.slice().updateUnsafe((_40_v4).f10,_40_v4);
              let _42_v6;
              _42_v6 = _dafny.Map.Empty.slice().updateUnsafe(_38_v2,_40_v4);
              let _43_v7;
              _43_v7 = new _dafny.CodePoint('h'.codePointAt(0));
              let _44_v8;
              let _nw2 = Array((new BigNumber(19)).toNumber());
              _nw2[(_dafny.ZERO).toNumber()] = _40_v4;
              _nw2[(_dafny.ONE).toNumber()] = _40_v4;
              _nw2[(new BigNumber(2)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(3)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(4)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(5)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(6)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(7)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(8)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(9)).toNumber()] = ((_35_v0) ? (_40_v4) : (_40_v4));
              _nw2[(new BigNumber(10)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(11)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(12)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(13)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(14)).toNumber()] = (((_41_v5).contains((_40_v4).f10)) ? ((_41_v5).get((_40_v4).f10)) : (_40_v4));
              _nw2[(new BigNumber(15)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(16)).toNumber()] = (((_42_v6).contains((_39_v3).fm9(true, p0, _43_v7, globalState))) ? ((_42_v6).get((_39_v3).fm9(true, p0, _43_v7, globalState))) : (_40_v4));
              _nw2[(new BigNumber(17)).toNumber()] = _40_v4;
              _nw2[(new BigNumber(18)).toNumber()] = _40_v4;
              _44_v8 = _nw2;
              let _index0 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_44_v8).length));
              (_44_v8)[_index0] = _40_v4;
              let _45_v9;
              let _nw3 = Array((new BigNumber(2)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = (_40_v4).f10;
              _nw3[(_dafny.ONE).toNumber()] = (_40_v4).f10;
              _45_v9 = _nw3;
              let _46_v10;
              _46_v10 = _dafny.MultiSet.fromElements((_40_v4).f10);
              let _47_v11;
              _47_v11 = _dafny.Map.Empty.slice().updateUnsafe((((_46_v10).contains(_35_v0)) ? ((_46_v10).get(_35_v0)) : ((_39_v3).f14)),_module.__default.safeModuloInt(new BigNumber(274), p0));
              let _index1 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_44_v8).length));
              let _rhs0 = (_dafny.ZERO).minus(p0);
              let _rhs1 = _40_v4;
              let _rhs2 = p1;
              let _rhs3 = _45_v9;
              let _rhs4 = _module.__default.fm29(_module.__default.safeDivisionInt((_39_v3).f14, p1), !(_46_v10).equals(_dafny.MultiSet.fromElements((_40_v4).f10)), globalState);
              let _lhs0 = _44_v8;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_44_v8).length));
              _38_v2 = _rhs0;
              _lhs0[_lhs1] = _rhs1;
              _38_v2 = _rhs2;
              _45_v9 = _rhs3;
              _47_v11 = _rhs4;
              _35_v0 = (_40_v4).f10;
            } else {
              let _48_v12;
              let _nw4 = new _module.C3();
              _nw4.__ctor(new BigNumber((_37_v1).length), _35_v0);
              _48_v12 = _nw4;
              _48_v12 = _48_v12;
              let _49_v13;
              let _nw5 = Array((new BigNumber(19)).toNumber()).fill(false);
              _49_v13 = _nw5;
              let _index2 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_49_v13).length));
              (_49_v13)[_index2] = !(_35_v0) || (false);
              let _index3 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_49_v13).length));
              (_49_v13)[_index3] = !(_48_v12.f7);
              _49_v13 = _49_v13;
              let _50_v14;
              _50_v14 = new BigNumber(929);
              let _51_v15;
              _51_v15 = _dafny.Set.fromElements(_48_v12, _48_v12);
              _50_v14 = (p0).plus(_module.__default.safeModuloInt(new BigNumber(694), new BigNumber((_51_v15).length)));
              let _52_v16;
              let _nw6 = Array((new BigNumber(4)).toNumber()).fill([]);
              _52_v16 = _nw6;
              let _index4 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_52_v16).length));
              (_52_v16)[_index4] = _49_v13;
              let _index5 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_52_v16).length));
              (_52_v16)[_index5] = _49_v13;
            }
            _35_v0 = !(p0).isEqualTo(_module.__default.safeModuloInt(p1, p1));
            if (_35_v0) {
              let _53_v17;
              _53_v17 = _dafny.MultiSet.fromElements(_35_v0, _35_v0);
              let _54_v18;
              _54_v18 = _dafny.Seq.of(_53_v17);
              let _55_v19;
              _55_v19 = _dafny.MultiSet.fromElements(_53_v17, _53_v17, (_54_v18)[_module.__default.safeIndex(p1, new BigNumber((_54_v18).length))], _53_v17);
              let _56_v20;
              let _nw7 = new _module.C3();
              _nw7.__ctor((((_55_v19).contains(_dafny.MultiSet.fromElements(_35_v0))) ? ((_55_v19).get(_dafny.MultiSet.fromElements(_35_v0))) : (p0)), _35_v0);
              _56_v20 = _nw7;
              let _57_v21;
              _57_v21 = new BigNumber(195);
              _57_v21 = new BigNumber((_dafny.Seq.UnicodeFromString("uwk")).length);
              let _58_v22;
              _58_v22 = new _dafny.CodePoint('c'.codePointAt(0));
              _58_v22 = _58_v22;
              _35_v0 = _35_v0;
              let _59_v23;
              _59_v23 = _module.D1.create_DC3(_module.__default.fm27(_35_v0, _58_v22, _35_v0, _35_v0, globalState));
              let _60_v24;
              _60_v24 = _module.D1.create_DC6(_59_v23);
              _60_v24 = _60_v24;
            } else {
              let _61_v25;
              let _nw8 = Array((new BigNumber(19)).toNumber()).fill(false);
              _61_v25 = _nw8;
              let _62_v26;
              let _nw9 = Array((new BigNumber(21)).toNumber());
              _62_v26 = _nw9;
              let _63_v27;
              _63_v27 = _dafny.Seq.of(_62_v26, _62_v26);
              let _64_v28;
              _64_v28 = _dafny.Seq.of((_63_v27)[_module.__default.safeIndex(p1, new BigNumber((_63_v27).length))], _62_v26);
              let _index6 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length));
              (_61_v25)[_index6] = (new BigNumber((_64_v28).length)).isLessThanOrEqualTo(p0);
              let _index7 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length));
              (_61_v25)[_index7] = true;
              let _65_v29;
              _65_v29 = new BigNumber(816);
              _65_v29 = p0;
              _65_v29 = p1;
              let _66_v30;
              _66_v30 = _dafny.Seq.UnicodeFromString("knmujqefl");
              let _67_v31;
              _67_v31 = _dafny.Set.fromElements(_61_v25);
              let _68_v32;
              _68_v32 = new _dafny.CodePoint('a'.codePointAt(0));
              let _69_v33;
              let _nw10 = new _module.C2();
              _nw10.__ctor(_67_v31, _68_v32, _35_v0);
              _69_v33 = _nw10;
              let _70_v34;
              _70_v34 = _dafny.Seq.of(_69_v33, _69_v33, _69_v33);
              let _71_v35;
              _71_v35 = _dafny.MultiSet.fromElements(_61_v25, _61_v25, _61_v25, _61_v25);
              let _72_v36;
              let _nw11 = new _module.C6();
              _nw11.__ctor((_61_v25)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length))], _35_v0);
              _72_v36 = _nw11;
              let _73_v37;
              _73_v37 = _dafny.Set.fromElements(_72_v36);
              let _74_v38;
              _74_v38 = _module.D8.create_DC21(p1, _67_v31, new BigNumber((_73_v37).length));
              let _pat_let_tv0 = _67_v31;
              let _75_v39;
              _75_v39 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(function (_pat_let0_0) {
                return function (_76_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_77_dt__update_hcf39_h0) {
                      return _module.D8.create_DC21((_76_dt__update__tmp_h0).dtor_cf38, _77_dt__update_hcf39_h0, (_76_dt__update__tmp_h0).dtor_cf40);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_74_v38))).length));
              let _78_v40;
              _78_v40 = _dafny.Seq.of(_35_v0, _35_v0);
              let _79_v41;
              _79_v41 = _dafny.Set.fromElements((_61_v25)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length))]);
              let _80_v42;
              _80_v42 = _dafny.Map.Empty.slice().updateUnsafe(true,_65_v29);
              let _81_v43;
              _81_v43 = _dafny.Set.fromElements(_65_v29, p1, p0, p0);
              let _82_v44;
              _82_v44 = _dafny.MultiSet.fromElements((_69_v33).f16);
              let _83_v45;
              let _nw12 = Array((new BigNumber(28)).toNumber());
              _nw12[(_dafny.ZERO).toNumber()] = _65_v29;
              _nw12[(_dafny.ONE).toNumber()] = _65_v29;
              _nw12[(new BigNumber(2)).toNumber()] = p1;
              _nw12[(new BigNumber(3)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(4)).toNumber()] = new BigNumber((_75_v39).cardinality());
              _nw12[(new BigNumber(5)).toNumber()] = p1;
              _nw12[(new BigNumber(6)).toNumber()] = new BigNumber((_80_v42).length);
              _nw12[(new BigNumber(7)).toNumber()] = p1;
              _nw12[(new BigNumber(8)).toNumber()] = p1;
              _nw12[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p1);
              _nw12[(new BigNumber(10)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(11)).toNumber()] = new BigNumber((_66_v30).length);
              _nw12[(new BigNumber(12)).toNumber()] = p1;
              _nw12[(new BigNumber(13)).toNumber()] = p1;
              _nw12[(new BigNumber(14)).toNumber()] = new BigNumber((_81_v43).length);
              _nw12[(new BigNumber(15)).toNumber()] = new BigNumber((_66_v30).length);
              _nw12[(new BigNumber(16)).toNumber()] = new BigNumber((_82_v44).cardinality());
              _nw12[(new BigNumber(17)).toNumber()] = p0;
              _nw12[(new BigNumber(18)).toNumber()] = new BigNumber((_66_v30).length);
              _nw12[(new BigNumber(19)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(20)).toNumber()] = p1;
              _nw12[(new BigNumber(21)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(22)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(23)).toNumber()] = p1;
              _nw12[(new BigNumber(24)).toNumber()] = p0;
              _nw12[(new BigNumber(25)).toNumber()] = p1;
              _nw12[(new BigNumber(26)).toNumber()] = _65_v29;
              _nw12[(new BigNumber(27)).toNumber()] = _65_v29;
              _83_v45 = _nw12;
              let _84_v46;
              _84_v46 = _dafny.Map.Empty.slice().updateUnsafe(_83_v45,_81_v43);
              let _85_v47;
              _85_v47 = _module.D13.create_DC34(_84_v46);
              let _86_v48;
              _86_v48 = _dafny.Seq.of(_85_v47);
              let _87_v49;
              _87_v49 = _module.D17.create_DC49(new BigNumber((_79_v41).length), _86_v48, (_dafny.ZERO).minus(new BigNumber((_78_v40).length)), (_61_v25)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length))]);
              let _88_v50;
              _88_v50 = _module.D2.create_DC7(_81_v43);
              let _89_v51;
              let _nw13 = Array((new BigNumber(6)).toNumber());
              _nw13[(_dafny.ZERO).toNumber()] = _88_v50;
              _nw13[(_dafny.ONE).toNumber()] = _88_v50;
              _nw13[(new BigNumber(2)).toNumber()] = _88_v50;
              _nw13[(new BigNumber(3)).toNumber()] = _88_v50;
              _nw13[(new BigNumber(4)).toNumber()] = _88_v50;
              _nw13[(new BigNumber(5)).toNumber()] = _module.D2.create_DC7(_81_v43);
              _89_v51 = _nw13;
              let _90_v52;
              let _nw14 = new _module.C4();
              _nw14.__ctor(_89_v51, (_61_v25)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length))]);
              _90_v52 = _nw14;
              let _91_v53;
              _91_v53 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_72_v36.f7,_90_v52), _dafny.Map.Empty.slice().updateUnsafe(_72_v36.f7,_90_v52));
              let _92_v55;
              let _nw15 = Array((new BigNumber(21)).toNumber());
              _nw15[(_dafny.ZERO).toNumber()] = _65_v29;
              _nw15[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((p0).minus(new BigNumber(761)));
              _nw15[(new BigNumber(2)).toNumber()] = (p1).plus(new BigNumber((_66_v30).length));
              _nw15[(new BigNumber(3)).toNumber()] = new BigNumber((_70_v34).length);
              _nw15[(new BigNumber(4)).toNumber()] = (new BigNumber((_71_v35).cardinality())).multipliedBy(_65_v29);
              _nw15[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p1);
              _nw15[(new BigNumber(6)).toNumber()] = new BigNumber((_75_v39).cardinality());
              _nw15[(new BigNumber(7)).toNumber()] = new BigNumber((_75_v39).cardinality());
              _nw15[(new BigNumber(8)).toNumber()] = new BigNumber((_78_v40).length);
              _nw15[(new BigNumber(9)).toNumber()] = _module.__default.fm0(globalState);
              _nw15[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_65_v29);
              _nw15[(new BigNumber(11)).toNumber()] = new BigNumber((_66_v30).length);
              _nw15[(new BigNumber(12)).toNumber()] = (_87_v49).dtor_cf82;
              _nw15[(new BigNumber(13)).toNumber()] = (new BigNumber((_dafny.MultiSet.FromArray(_78_v40)).cardinality())).multipliedBy(new BigNumber((_91_v53).cardinality()));
              _nw15[(new BigNumber(14)).toNumber()] = _65_v29;
              _nw15[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_72_v36).fm9((_61_v25)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length))], p0, (_69_v33).f16, globalState)));
              _nw15[(new BigNumber(16)).toNumber()] = p0;
              _nw15[(new BigNumber(17)).toNumber()] = _65_v29;
              _nw15[(new BigNumber(18)).toNumber()] = p1;
              _nw15[(new BigNumber(19)).toNumber()] = _65_v29;
              _nw15[(new BigNumber(20)).toNumber()] = new BigNumber(((_81_v43).Intersect(function () {
                let _coll12 = new _dafny.Set();
                for (const _compr_12 of _dafny.IntegerRange(new BigNumber(744), new BigNumber(-968))) {
                  let _93_v54 = _compr_12;
                  if (((new BigNumber(744)).isLessThanOrEqualTo(_93_v54)) && ((_93_v54).isLessThan(new BigNumber(-968)))) {
                    _coll12.add(_module.__default.safeModuloInt(_93_v54, new BigNumber(450)));
                  }
                }
                return _coll12;
              }())).length);
              _92_v55 = _nw15;
              let _index8 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_92_v55).length));
              (_92_v55)[_index8] = p0;
              let _index9 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_92_v55).length));
              let _index10 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length));
              let _rhs5 = (p1).plus(_module.__default.safeModuloInt(p1, new BigNumber((_79_v41).length)));
              let _rhs6 = new BigNumber((_79_v41).length);
              let _rhs7 = _35_v0;
              let _lhs2 = _92_v55;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_92_v55).length));
              let _lhs4 = _61_v25;
              let _lhs5 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_61_v25).length));
              _65_v29 = _rhs5;
              _lhs2[_lhs3] = _rhs6;
              _lhs4[_lhs5] = _rhs7;
              let _94_v56;
              _94_v56 = _dafny.Map.Empty.slice().updateUnsafe(true,_72_v36.f7);
              _94_v56 = (_94_v56).update((_72_v36).fm8(p0, globalState), (true) || (false));
            }
            let _95_v57;
            _95_v57 = _module.D27.create_DC75(_37_v1);
            let _96_v58;
            let _nw16 = new _module.C3();
            _nw16.__ctor(p0, !((_95_v57).dtor_cf130).equals(_37_v1));
            _96_v58 = _nw16;
            let _97_v59;
            _97_v59 = _dafny.Seq.of(p1);
            let _98_v60;
            let _nw17 = new _module.C7();
            _nw17.__ctor(_dafny.Map.Empty.slice().updateUnsafe(true,_97_v59), _35_v0);
            _98_v60 = _nw17;
            let _99_v61;
            let _init0 = ((_100_p1) => function (_101_i1) {
              return (_101_i1).multipliedBy(_100_p1);
            })(p1);
            let _nw18 = Array((new BigNumber(29)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw18.length); _i0_0++) {
              _nw18[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _99_v61 = _nw18;
            let _102_v62;
            let _nw19 = Array((new BigNumber(26)).toNumber());
            _nw19[(_dafny.ZERO).toNumber()] = _99_v61;
            _nw19[(_dafny.ONE).toNumber()] = _99_v61;
            _nw19[(new BigNumber(2)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(3)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(4)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(5)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(6)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(7)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(8)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(9)).toNumber()] = ((_35_v0) ? (_99_v61) : (_99_v61));
            _nw19[(new BigNumber(10)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(11)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(12)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(13)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(14)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(15)).toNumber()] = ((_96_v58.f7) ? (_99_v61) : (_99_v61));
            _nw19[(new BigNumber(16)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(17)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(18)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(19)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(20)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(21)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(22)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(23)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(24)).toNumber()] = _99_v61;
            _nw19[(new BigNumber(25)).toNumber()] = _99_v61;
            _102_v62 = _nw19;
            let _index11 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_102_v62).length));
            (_102_v62)[_index11] = ((_35_v0) ? (_99_v61) : (_99_v61));
            let _103_v63;
            _103_v63 = _module.D18.create_DC50(_96_v58);
            let _index12 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_102_v62).length));
            let _rhs8 = _96_v58;
            let _rhs9 = _98_v60;
            let _rhs10 = _35_v0;
            let _rhs11 = _99_v61;
            let _rhs12 = _103_v63;
            let _lhs6 = _102_v62;
            let _lhs7 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_102_v62).length));
            _96_v58 = _rhs8;
            _98_v60 = _rhs9;
            _35_v0 = _rhs10;
            _lhs6[_lhs7] = _rhs11;
            _103_v63 = _rhs12;
          }
        }
      }
      let _104_v64;
      _104_v64 = _dafny.Seq.UnicodeFromString("ed");
      let _105_v65;
      _105_v65 = _dafny.Seq.of(_104_v64);
      let _106_v66;
      _106_v66 = _dafny.Map.Empty.slice().updateUnsafe(_35_v0,p0);
      let _107_v67;
      _107_v67 = _dafny.MultiSet.fromElements(p1);
      let _108_v68;
      _108_v68 = _dafny.MultiSet.fromElements(p1, p1, (_dafny.ZERO).minus(new BigNumber((_107_v67).cardinality())), p0, p0);
      let _109_v69;
      _109_v69 = _dafny.MultiSet.fromElements(false);
      let _110_v70;
      _110_v70 = _dafny.Seq.of(_109_v69, _109_v69);
      let _111_v71;
      let _nw20 = Array((new BigNumber(27)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = _104_v64;
      _nw20[(_dafny.ONE).toNumber()] = _104_v64;
      _nw20[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("pxnpy");
      _nw20[(new BigNumber(3)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(4)).toNumber()] = _module.__default.fm28(_35_v0, p0, globalState);
      _nw20[(new BigNumber(5)).toNumber()] = (_105_v65)[_module.__default.safeIndex(p1, new BigNumber((_105_v65).length))];
      _nw20[(new BigNumber(6)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), function (_112_i2) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      });
      _nw20[(new BigNumber(8)).toNumber()] = _module.__default.fm35(new BigNumber(730), p0, globalState);
      _nw20[(new BigNumber(9)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(10)).toNumber()] = _module.__default.fm28(_35_v0, (_dafny.ZERO).minus(new BigNumber((_106_v66).length)), globalState);
      _nw20[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_104_v64, _104_v64);
      _nw20[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_104_v64, _104_v64);
      _nw20[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_104_v64, _104_v64);
      _nw20[(new BigNumber(14)).toNumber()] = _module.__default.fm28(_35_v0, new BigNumber((_module.__default.fm34(new BigNumber((_107_v67).cardinality()), _35_v0, _35_v0, new BigNumber(((_110_v70)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_110_v70).length))]).cardinality()), globalState)).length), globalState);
      _nw20[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("hccoqdanf");
      _nw20[(new BigNumber(16)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(17)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(18)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_104_v64, _104_v64);
      _nw20[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("jlpfou");
      _nw20[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wniyk"), _104_v64);
      _nw20[(new BigNumber(22)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("vclldfv");
      _nw20[(new BigNumber(24)).toNumber()] = _104_v64;
      _nw20[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("vimrxtn");
      _nw20[(new BigNumber(26)).toNumber()] = _104_v64;
      _111_v71 = _nw20;
      let _113_v72;
      _113_v72 = _module.D28.create_DC77(_111_v71);
      _111_v71 = (_113_v72).dtor_cf134;
      let _114_v73;
      _114_v73 = new BigNumber(999);
      let _115_v74;
      _115_v74 = _dafny.Map.Empty.slice().updateUnsafe(_114_v73,_114_v73);
      _114_v73 = (((_115_v74).contains(p1)) ? ((_115_v74).get(p1)) : ((p1).minus(p1)));
      let _116_v75;
      _116_v75 = _dafny.Seq.of(_35_v0);
      let _117_v76;
      _117_v76 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(p0),_116_v75);
      _114_v73 = new BigNumber(((((_117_v76).contains(_module.D0.create_DC1(p0))) ? ((_117_v76).get(_module.D0.create_DC1(p0))) : (_dafny.Seq.Concat(_116_v75, _116_v75)))).length);
      let _118_v77;
      _118_v77 = _dafny.Map.Empty.slice().updateUnsafe(_104_v64,_35_v0);
      _35_v0 = (((_118_v77).contains(_dafny.Seq.Concat(_104_v64, _104_v64))) ? ((_118_v77).get(_dafny.Seq.Concat(_104_v64, _104_v64))) : (_35_v0));
      let _119_v78;
      _119_v78 = _dafny.Set.fromElements(new BigNumber(535));
      let _120_v79;
      _120_v79 = _module.D2.create_DC7(_119_v78);
      let _pat_let_tv1 = _119_v78;
      let _pat_let_tv2 = _35_v0;
      let _pat_let_tv3 = globalState;
      let _121_v81;
      let _nw21 = Array((new BigNumber(25)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = function (_pat_let2_0) {
        return function (_122_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_123_dt__update_hcf13_h0) {
              return _module.D2.create_DC7(_123_dt__update_hcf13_h0);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let2_0);
      }(_120_v79);
      _nw21[(_dafny.ONE).toNumber()] = _120_v79;
      _nw21[(new BigNumber(2)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(3)).toNumber()] = _module.D2.create_DC7(_119_v78);
      _nw21[(new BigNumber(4)).toNumber()] = function (_pat_let4_0) {
        return function (_124_dt__update__tmp_h2) {
          return function (_pat_let5_0) {
            return function (_125_dt__update_hcf13_h1) {
              return _module.D2.create_DC7(_125_dt__update_hcf13_h1);
            }(_pat_let5_0);
          }(_module.__default.fm2(new BigNumber(644), _pat_let_tv2, _pat_let_tv3));
        }(_pat_let4_0);
      }(_module.__default.fm64(globalState));
      _nw21[(new BigNumber(5)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(6)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(7)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(8)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(9)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(10)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(11)).toNumber()] = _module.D2.create_DC7(_dafny.Set.fromElements(_114_v73));
      _nw21[(new BigNumber(12)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(13)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(14)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(15)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(16)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(17)).toNumber()] = _module.D2.create_DC7(_dafny.Set.fromElements(_114_v73, new BigNumber((function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of _dafny.IntegerRange(new BigNumber(75), new BigNumber(633))) {
    let _126_v80 = _compr_13;
    if (((new BigNumber(75)).isLessThanOrEqualTo(_126_v80)) && ((_126_v80).isLessThan(new BigNumber(633)))) {
      _coll13.push([_module.__default.safeModuloInt(_126_v80, p0),new BigNumber(779)]);
    }
  }
  return _coll13;
}()).length)));
      _nw21[(new BigNumber(18)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(19)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(20)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(21)).toNumber()] = _module.D2.create_DC7(_dafny.Set.fromElements(_114_v73));
      _nw21[(new BigNumber(22)).toNumber()] = _120_v79;
      _nw21[(new BigNumber(23)).toNumber()] = _module.D2.create_DC7(_119_v78);
      _nw21[(new BigNumber(24)).toNumber()] = _120_v79;
      _121_v81 = _nw21;
      let _127_v82;
      let _nw22 = new _module.C4();
      _nw22.__ctor(_121_v81, (p1).isLessThanOrEqualTo(_114_v73));
      _127_v82 = _nw22;
      return;
    }
    static Main(__noArgsParameter) {
      let _128_v0;
      _128_v0 = new BigNumber(-231);
      let _129_v1;
      _129_v1 = _dafny.Seq.of(_dafny.Set.fromElements(_128_v0));
      let _130_v2;
      _130_v2 = true;
      let _131_v3;
      _131_v3 = _dafny.Seq.of(_130_v2);
      let _132_v4;
      _132_v4 = _dafny.MultiSet.fromElements(true, (_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))]);
      let _133_v5;
      let _init1 = ((_134_v3) => function (_135_i0) {
        return _134_v3;
      })(_131_v3);
      let _nw23 = Array((new BigNumber(29)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw23.length); _i0_1++) {
        _nw23[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _133_v5 = _nw23;
      let _136_globalState;
      let _nw24 = new _module.GlobalState();
      _nw24.__ctor(false, false, _129_v1, _132_v4, new BigNumber(290), _133_v5);
      _136_globalState = _nw24;
      if (_130_v2) {
        let _137_v6;
        _137_v6 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_module.__default.fm0(_136_globalState));
        _137_v6 = (_137_v6).update(_130_v2, ((_130_v2) ? (_128_v0) : (_128_v0)));
        let _138_v7;
        let _nw25 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _138_v7 = _nw25;
        _138_v7 = _138_v7;
        _137_v6 = (_137_v6).update((_130_v2) || (_130_v2), _128_v0);
        let _139_v8;
        _139_v8 = _dafny.Seq.of(_128_v0, _128_v0);
        _128_v0 = (new BigNumber((_139_v8).length)).multipliedBy(_128_v0);
        if (false) {
          _128_v0 = (_128_v0).minus(_module.__default.safeModuloInt(_128_v0, _128_v0));
          let _140_v9;
          let _nw26 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _140_v9 = _nw26;
          let _141_v10;
          _141_v10 = _dafny.Set.fromElements(_140_v9, _140_v9);
          let _142_v11;
          _142_v11 = _module.D0.create_DC2(_130_v2, _128_v0, _141_v10, new BigNumber(558));
          let _pat_let_tv4 = _128_v0;
          let _pat_let_tv5 = _141_v10;
          _128_v0 = (function (_pat_let6_0) {
            return function (_143_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_144_dt__update_hcf5_h0) {
                  return function (_pat_let8_0) {
                    return function (_145_dt__update_hcf4_h0) {
                      return _module.D0.create_DC2((_143_dt__update__tmp_h0).dtor_cf2, (_143_dt__update__tmp_h0).dtor_cf3, _145_dt__update_hcf4_h0, _144_dt__update_hcf5_h0);
                    }(_pat_let8_0);
                  }(_pat_let_tv5);
                }(_pat_let7_0);
              }(_pat_let_tv4);
            }(_pat_let6_0);
          }(_142_v11)).dtor_cf3;
          _130_v2 = _130_v2;
          _128_v0 = _module.__default.safeDivisionInt(_128_v0, _128_v0);
          let _146_v12;
          _146_v12 = _dafny.Seq.UnicodeFromString("h");
          let _147_v13;
          _147_v13 = _dafny.Set.fromElements((_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))]);
          let _148_v14;
          _148_v14 = new _dafny.CodePoint('w'.codePointAt(0));
          let _149_v15;
          _149_v15 = _dafny.Set.fromElements(_148_v14);
          let _150_v16;
          _150_v16 = _dafny.Set.fromElements(_128_v0, _128_v0);
          let _151_v17;
          _151_v17 = _dafny.Map.Empty.slice().updateUnsafe(_150_v16,_130_v2);
          let _152_v18;
          let _nw27 = Array((new BigNumber(20)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _130_v2;
          _nw27[(_dafny.ONE).toNumber()] = _130_v2;
          _nw27[(new BigNumber(2)).toNumber()] = (_128_v0).isLessThanOrEqualTo(_128_v0);
          _nw27[(new BigNumber(3)).toNumber()] = !_dafny.areEqual(_146_v12, _146_v12);
          _nw27[(new BigNumber(4)).toNumber()] = !((_130_v2) && (_130_v2));
          _nw27[(new BigNumber(5)).toNumber()] = _130_v2;
          _nw27[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(_130_v2)).equals(_147_v13);
          _nw27[(new BigNumber(7)).toNumber()] = true;
          _nw27[(new BigNumber(8)).toNumber()] = (_147_v13).IsProperSubsetOf(_147_v13);
          _nw27[(new BigNumber(9)).toNumber()] = _130_v2;
          _nw27[(new BigNumber(10)).toNumber()] = !(_module.__default.fm1(_dafny.Seq.of(_module.D0.create_DC0(_128_v0)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_153_i1) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length), new BigNumber((_137_v6).length), _136_globalState));
          _nw27[(new BigNumber(11)).toNumber()] = (_132_v4).contains(_130_v2);
          _nw27[(new BigNumber(12)).toNumber()] = _130_v2;
          _nw27[(new BigNumber(13)).toNumber()] = (_128_v0).isLessThan(_128_v0);
          _nw27[(new BigNumber(14)).toNumber()] = _130_v2;
          _nw27[(new BigNumber(15)).toNumber()] = (_dafny.Set.fromElements(_130_v2, false)).IsProperSubsetOf(_147_v13);
          _nw27[(new BigNumber(16)).toNumber()] = (_149_v15).IsProperSubsetOf(_149_v15);
          _nw27[(new BigNumber(17)).toNumber()] = (_130_v2) && (_130_v2);
          _nw27[(new BigNumber(18)).toNumber()] = !((_151_v17).contains(_module.__default.fm2(_128_v0, _130_v2, _136_globalState)));
          _nw27[(new BigNumber(19)).toNumber()] = _130_v2;
          _152_v18 = _nw27;
          _152_v18 = _152_v18;
        } else {
          let _154_v19;
          let _nw28 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _154_v19 = _nw28;
          let _155_v20;
          _155_v20 = new _dafny.CodePoint('i'.codePointAt(0));
          let _index13 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_154_v19).length));
          (_154_v19)[_index13] = _155_v20;
          let _index14 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_154_v19).length));
          (_154_v19)[_index14] = _155_v20;
          let _156_v21;
          let _nw29 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
          _156_v21 = _nw29;
          let _157_v22;
          _157_v22 = _dafny.MultiSet.fromElements(new BigNumber(340));
          let _index15 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_156_v21).length));
          (_156_v21)[_index15] = _157_v22;
          let _index16 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_156_v21).length));
          let _rhs13 = _130_v2;
          let _rhs14 = (_dafny.ZERO).minus(_128_v0);
          let _rhs15 = _157_v22;
          let _lhs8 = _156_v21;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_156_v21).length));
          _130_v2 = _rhs13;
          _128_v0 = _rhs14;
          _lhs8[_lhs9] = _rhs15;
          let _158_v23;
          _158_v23 = _dafny.Seq.UnicodeFromString("bdayhkro");
          _158_v23 = _dafny.Seq.Concat(_158_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), ((_159_v19) => function (_160_i2) {
            return (_159_v19)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_159_v19).length))];
          })(_154_v19)));
          _module.__default.m0(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gvuwlpyo"), _158_v23)).length), _128_v0, _136_globalState);
          _module.__default.m0((_dafny.ZERO).minus(_module.__default.fm0(_136_globalState)), _128_v0, _136_globalState);
        }
      } else {
        let _161_v24;
        _161_v24 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("wqkta")).length));
        _161_v24 = _161_v24;
        let _162_v25;
        _162_v25 = _dafny.Seq.of(_161_v24);
        let _163_v26;
        _163_v26 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_128_v0);
        let _164_v27;
        let _nw30 = Array((new BigNumber(28)).toNumber());
        _nw30[(_dafny.ZERO).toNumber()] = _130_v2;
        _nw30[(_dafny.ONE).toNumber()] = !(_130_v2);
        _nw30[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_162_v25, _128_v0, _128_v0, _136_globalState);
        _nw30[(new BigNumber(3)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(4)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(5)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(6)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(7)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(8)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(9)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(10)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(11)).toNumber()] = !(_130_v2);
        _nw30[(new BigNumber(12)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(13)).toNumber()] = (_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))];
        _nw30[(new BigNumber(14)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(15)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(16)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(17)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(18)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(19)).toNumber()] = _module.__default.fm1(_162_v25, new BigNumber((_163_v26).length), _128_v0, _136_globalState);
        _nw30[(new BigNumber(20)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(21)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(22)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(23)).toNumber()] = false;
        _nw30[(new BigNumber(24)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(25)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(26)).toNumber()] = _130_v2;
        _nw30[(new BigNumber(27)).toNumber()] = _130_v2;
        _164_v27 = _nw30;
        let _165_v28;
        _165_v28 = _dafny.Seq.UnicodeFromString("mre");
        let _166_v29;
        _166_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_164_v27,_165_v28),_128_v0);
        let _167_v30;
        _167_v30 = _dafny.Map.Empty.slice().updateUnsafe(_164_v27,_165_v28);
        _166_v29 = (_166_v29).update(_167_v30, _128_v0);
        let _168_v31;
        let _nw31 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _168_v31 = _nw31;
        let _169_v32;
        _169_v32 = _dafny.MultiSet.fromElements(_168_v31);
        let _170_v33;
        _170_v33 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,(((_169_v32).contains(_168_v31)) ? ((_169_v32).get(_168_v31)) : (_128_v0)));
        _module.__default.m0(_128_v0, new BigNumber((_170_v33).length), _136_globalState);
        let _171_v34;
        _171_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_130_v2, new BigNumber((_132_v4).cardinality()), _128_v0, new BigNumber((_165_v28).length), _136_globalState),_130_v2);
        let _172_v35;
        _172_v35 = new _dafny.CodePoint('t'.codePointAt(0));
        _171_v34 = (_171_v34).update(_dafny.Seq.of(_dafny.Seq.update(_165_v28, _module.__default.safeIndex(_128_v0, new BigNumber((_165_v28).length)), _172_v35), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-651)), ((_173_v35) => function (_174_i3) {
          return _173_v35;
        })(_172_v35))), _130_v2);
        let _nw32 = Array((new BigNumber(9)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _130_v2;
        _nw32[(_dafny.ONE).toNumber()] = (_132_v4).IsProperSubsetOf(_132_v4);
        _nw32[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), ((_175_v0) => function (_176_i4) {
          return _175_v0;
        })(_128_v0)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), ((_177_v35) => function (_178_i5) {
          return _177_v35;
        })(_172_v35))).length)));
        _nw32[(new BigNumber(3)).toNumber()] = (_128_v0).isLessThanOrEqualTo(new BigNumber(-697));
        _nw32[(new BigNumber(4)).toNumber()] = _130_v2;
        _nw32[(new BigNumber(5)).toNumber()] = !(_130_v2);
        _nw32[(new BigNumber(6)).toNumber()] = _130_v2;
        _nw32[(new BigNumber(7)).toNumber()] = (_128_v0).isLessThan(new BigNumber((_165_v28).length));
        _nw32[(new BigNumber(8)).toNumber()] = _130_v2;
        _164_v27 = _nw32;
      }
      let _179_v36;
      _179_v36 = _dafny.MultiSet.fromElements(_128_v0, _128_v0, new BigNumber(826));
      let _hi0 = new BigNumber((_179_v36).cardinality());
      for (let _180_i6 = _128_v0; _180_i6.isLessThan(_hi0); _180_i6 = _180_i6.plus(_dafny.ONE)) {
        _179_v36 = _179_v36;
        if (_130_v2) {
          let _181_v37;
          let _init2 = ((_182_v36, _183_v3, _184_v0) => function (_185_i7) {
            return (_dafny.MultiSet.fromElements(new BigNumber((_183_v3).length), _184_v0)).IsProperSubsetOf(_182_v36);
          })(_179_v36, _131_v3, _128_v0);
          let _nw33 = Array((new BigNumber(7)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw33.length); _i0_2++) {
            _nw33[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _181_v37 = _nw33;
          let _index17 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length));
          (_181_v37)[_index17] = _130_v2;
          let _index18 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length));
          (_181_v37)[_index18] = _130_v2;
          let _index19 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length));
          (_181_v37)[_index19] = !((_181_v37)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length))]);
          let _186_v38;
          _186_v38 = _dafny.Seq.UnicodeFromString("lragd");
          _186_v38 = _dafny.Seq.Concat(_186_v38, _186_v38);
          _130_v2 = (_181_v37)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length))];
          _module.__default.m0(_180_i6, (new BigNumber((function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_module.__default.fm4(_180_i6, _136_globalState)).Keys.Elements) {
              let _187_v39 = _compr_14;
              if ((_module.__default.fm4(_180_i6, _136_globalState)).contains(_187_v39)) {
                _coll14.push([_187_v39,(_181_v37)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_181_v37).length))]]);
              }
            }
            return _coll14;
          }()).length)).multipliedBy(new BigNumber(650)), _136_globalState);
        } else {
          let _188_v40;
          _188_v40 = _module.D0.create_DC0(_180_i6);
          let _189_v41;
          _189_v41 = _dafny.Seq.of(_188_v40, function (_pat_let9_0) {
            return function (_190_dt__update__tmp_h1) {
              return function (_pat_let10_0) {
                return function (_191_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_191_dt__update_hcf0_h0);
                }(_pat_let10_0);
              }(_180_i6);
            }(_pat_let9_0);
          }(_module.D0.create_DC0(_180_i6)), _188_v40, _188_v40);
          let _192_v42;
          _192_v42 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), ((_193_v0) => function (_194_i8) {
            return _193_v0;
          })(_128_v0))).length), _180_i6);
          _130_v2 = _module.__default.fm1(_dafny.Seq.Concat(_189_v41, _189_v41), _180_i6, new BigNumber((_dafny.Seq.Concat(_192_v42, _192_v42)).length), _136_globalState);
          let _195_v43;
          _195_v43 = _dafny.Seq.UnicodeFromString("mb");
          _195_v43 = _195_v43;
          _module.__default.m0(_128_v0, (_192_v42)[_module.__default.safeIndex(_128_v0, new BigNumber((_192_v42).length))], _136_globalState);
          _128_v0 = _128_v0;
          let _196_v44;
          _196_v44 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_130_v2);
          let _197_v45;
          _197_v45 = _dafny.MultiSet.fromElements(_196_v44, _196_v44, _196_v44, _196_v44);
          _130_v2 = (_197_v45).IsProperSubsetOf(_dafny.MultiSet.fromElements(_196_v44, _196_v44));
        }
        let _198_v46;
        _198_v46 = _dafny.Seq.UnicodeFromString("pxbij");
        _198_v46 = _dafny.Seq.Concat(_198_v46, _198_v46);
        if (true) {
          let _199_v47;
          _199_v47 = _module.D0.create_DC0(_128_v0);
          let _200_v48;
          let _nw34 = Array((new BigNumber(6)).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = _module.__default.fm1(_dafny.Seq.of(_199_v47, _199_v47), _128_v0, _128_v0, _136_globalState);
          _nw34[(_dafny.ONE).toNumber()] = !(new BigNumber(-499)).isEqualTo(_128_v0);
          _nw34[(new BigNumber(2)).toNumber()] = _130_v2;
          _nw34[(new BigNumber(3)).toNumber()] = !(_130_v2) || (_130_v2);
          _nw34[(new BigNumber(4)).toNumber()] = !(_130_v2);
          _nw34[(new BigNumber(5)).toNumber()] = ((!(_130_v2)) ? (false) : (!(_130_v2)));
          _200_v48 = _nw34;
          let _index20 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_200_v48).length));
          (_200_v48)[_index20] = _130_v2;
          let _index21 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_200_v48).length));
          (_200_v48)[_index21] = (_128_v0).isLessThanOrEqualTo(_180_i6);
          _128_v0 = _128_v0;
          let _201_v49;
          let _init3 = ((_202_v48) => function (_203_i9) {
            return (_203_i9).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_202_v48)[_module.__default.safeIndex(new BigNumber(152), new BigNumber((_202_v48).length))],true)).length));
          })(_200_v48);
          let _nw35 = Array((new BigNumber(12)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw35.length); _i0_3++) {
            _nw35[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _201_v49 = _nw35;
          let _index22 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_201_v49).length));
          (_201_v49)[_index22] = new BigNumber(543);
          let _index23 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_201_v49).length));
          (_201_v49)[_index23] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_180_i6), (_180_i6).plus(_128_v0));
          let _204_v50;
          _204_v50 = _dafny.Seq.of(_201_v49);
          _module.__default.m0(new BigNumber((_204_v50).length), (new BigNumber((_198_v46).length)).plus(_180_i6), _136_globalState);
          let _205_v51;
          _205_v51 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_180_i6);
          _199_v47 = _module.D0.create_DC0(new BigNumber((_205_v51).length));
        } else {
          let _206_v52;
          _206_v52 = _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_130_v2),_130_v2)).length)));
          let _207_v53;
          _207_v53 = _module.D0.create_DC1(_128_v0);
          let _208_v54;
          _208_v54 = _dafny.Set.fromElements(_180_i6, _128_v0);
          let _209_v55;
          let _nw36 = Array((new BigNumber(13)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = (_206_v52).dtor_cf0;
          _nw36[(_dafny.ONE).toNumber()] = new BigNumber((_198_v46).length);
          _nw36[(new BigNumber(2)).toNumber()] = _180_i6;
          _nw36[(new BigNumber(3)).toNumber()] = _180_i6;
          _nw36[(new BigNumber(4)).toNumber()] = (_207_v53).dtor_cf1;
          _nw36[(new BigNumber(5)).toNumber()] = new BigNumber(49);
          _nw36[(new BigNumber(6)).toNumber()] = new BigNumber((_208_v54).length);
          _nw36[(new BigNumber(7)).toNumber()] = _128_v0;
          _nw36[(new BigNumber(8)).toNumber()] = _128_v0;
          _nw36[(new BigNumber(9)).toNumber()] = _128_v0;
          _nw36[(new BigNumber(10)).toNumber()] = _180_i6;
          _nw36[(new BigNumber(11)).toNumber()] = new BigNumber(895);
          _nw36[(new BigNumber(12)).toNumber()] = _180_i6;
          _209_v55 = _nw36;
          let _210_v56;
          _210_v56 = _dafny.Map.Empty.slice().updateUnsafe(_180_i6,_209_v55);
          _210_v56 = (_210_v56).update(_128_v0, _209_v55);
          let _211_v57;
          let _nw37 = Array((new BigNumber(5)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _130_v2;
          _nw37[(_dafny.ONE).toNumber()] = _130_v2;
          _nw37[(new BigNumber(2)).toNumber()] = _130_v2;
          _nw37[(new BigNumber(3)).toNumber()] = _130_v2;
          _nw37[(new BigNumber(4)).toNumber()] = _130_v2;
          _211_v57 = _nw37;
          let _212_v58;
          _212_v58 = _dafny.Map.Empty.slice().updateUnsafe(_211_v57,_130_v2);
          _128_v0 = new BigNumber((_212_v58).length);
          let _index24 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_211_v57).length));
          (_211_v57)[_index24] = (_180_i6).isEqualTo(_128_v0);
          let _index25 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_211_v57).length));
          (_211_v57)[_index25] = (_130_v2) === (_130_v2);
          let _213_v59;
          let _nw38 = Array((new BigNumber(4)).toNumber()).fill(_module.D0.Default());
          _213_v59 = _nw38;
          let _214_v60;
          _214_v60 = _dafny.Seq.of(_213_v59);
          _214_v60 = _dafny.Seq.update(_214_v60, _module.__default.safeIndex(_128_v0, new BigNumber((_214_v60).length)), _213_v59);
          let _index26 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_211_v57).length));
          (_211_v57)[_index26] = (_211_v57)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_211_v57).length))];
        }
      }
      _130_v2 = _130_v2;
      let _215_v61;
      _215_v61 = _module.D0.create_DC0(_128_v0);
      let _216_v62;
      _216_v62 = _dafny.Map.Empty.slice().updateUnsafe(_215_v61,((_130_v2) ? (_130_v2) : (_130_v2)));
      let _217_v63;
      _217_v63 = _dafny.MultiSet.fromElements(_179_v36);
      if ((((_216_v62).contains(_215_v61)) ? ((_216_v62).get(_215_v61)) : ((_217_v63).IsProperSubsetOf(_217_v63)))) {
        let _218_v64;
        let _nw39 = Array((new BigNumber(2)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = !(_130_v2);
        _nw39[(_dafny.ONE).toNumber()] = (_132_v4).IsProperSubsetOf(_132_v4);
        _218_v64 = _nw39;
        let _index27 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_218_v64).length));
        (_218_v64)[_index27] = _130_v2;
        let _index28 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_218_v64).length));
        (_218_v64)[_index28] = true;
        let _219_v65;
        _219_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_132_v4).cardinality()),_128_v0);
        _219_v65 = (_219_v65).update(_128_v0, (_128_v0).multipliedBy(_module.__default.fm0(_136_globalState)));
        let _220_v66;
        _220_v66 = _dafny.Map.Empty.slice().updateUnsafe(!(_130_v2) || (_module.__default.fm1(_dafny.Seq.of(_215_v61), _128_v0, _128_v0, _136_globalState)),_132_v4);
        let _221_v67;
        let _init4 = ((_222_v0) => function (_223_i10) {
          return (_223_i10).plus(_222_v0);
        })(_128_v0);
        let _nw40 = Array((new BigNumber(26)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw40.length); _i0_4++) {
          _nw40[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _221_v67 = _nw40;
        let _224_v68;
        _224_v68 = _dafny.Set.fromElements(_221_v67, _221_v67);
        let _225_v69;
        _225_v69 = _module.D0.create_DC2(true, _128_v0, _224_v68, _128_v0);
        let _226_v70;
        _226_v70 = _dafny.Seq.of(_224_v68);
        let _227_v71;
        _227_v71 = _module.D0.create_DC2((_225_v69).dtor_cf2, _128_v0, (_226_v70)[_module.__default.safeIndex(_128_v0, new BigNumber((_226_v70).length))], _128_v0);
        _220_v66 = (_220_v66).update(!((_227_v71).dtor_cf2), _132_v4);
        let _228_v72;
        _228_v72 = _dafny.Set.fromElements(true);
        let _229_v73;
        _229_v73 = _module.D1.create_DC3(_228_v72);
        _228_v72 = (_229_v73).dtor_cf6;
        let _230_v74;
        _230_v74 = _dafny.Set.fromElements(_218_v64);
        let _231_v75;
        _231_v75 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_130_v2);
        let _232_v76;
        _232_v76 = new _dafny.CodePoint('g'.codePointAt(0));
        let _233_v77;
        let _nw41 = new _module.C2();
        _nw41.__ctor((_module.D8.create_DC21(_128_v0, _230_v74, new BigNumber((_231_v75).length))).dtor_cf39, _232_v76, (_218_v64)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_218_v64).length))]);
        _233_v77 = _nw41;
      } else {
        let _234_v78;
        let _init5 = ((_235_v2) => function (_236_i11) {
          return _235_v2;
        })(_130_v2);
        let _nw42 = Array((new BigNumber(13)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw42.length); _i0_5++) {
          _nw42[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _234_v78 = _nw42;
        let _index29 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_234_v78).length));
        (_234_v78)[_index29] = _130_v2;
        let _237_v79;
        _237_v79 = _dafny.Seq.UnicodeFromString("nkcb");
        let _238_v80;
        _238_v80 = _dafny.Seq.of(_237_v79, _dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), function (_239_i12) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        }));
        let _index30 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_234_v78).length));
        (_234_v78)[_index30] = _dafny.Seq.IsPrefixOf((_238_v80)[_module.__default.safeIndex(_128_v0, new BigNumber((_238_v80).length))], _237_v79);
        _module.__default.m0(_128_v0, _128_v0, _136_globalState);
        let _240_v81;
        _240_v81 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_130_v2);
        let _241_v82;
        _241_v82 = _dafny.Seq.of(_module.D0.create_DC0(_128_v0));
        _240_v81 = (_240_v81).update(_130_v2, !(_module.__default.fm1(_241_v82, _128_v0, (_dafny.ZERO).minus(_128_v0), _136_globalState)) || ((_234_v78)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_234_v78).length))]));
        let _242_v83;
        _242_v83 = _dafny.Map.Empty.slice().updateUnsafe((_234_v78)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_234_v78).length))],_128_v0);
        let _243_v84;
        let _init6 = ((_244_v0) => function (_245_i13) {
          return _module.__default.safeDivisionInt(_245_i13, _244_v0);
        })(_128_v0);
        let _nw43 = Array((new BigNumber(19)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw43.length); _i0_6++) {
          _nw43[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _243_v84 = _nw43;
        let _246_v85;
        _246_v85 = _module.D13.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_243_v84,_module.__default.fm2(new BigNumber(540), _130_v2, _136_globalState)));
        let _247_v86;
        _247_v86 = _dafny.Seq.of(_246_v85);
        let _248_v87;
        _248_v87 = _module.D17.create_DC49(new BigNumber((_242_v83).length), _247_v86, new BigNumber((_242_v83).length), _130_v2);
        _248_v87 = _248_v87;
        let _249_v88;
        _249_v88 = _module.D6.create_DC16((_234_v78)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_234_v78).length))]);
        let _250_v89;
        _250_v89 = _dafny.Map.Empty.slice().updateUnsafe((_215_v61).dtor_cf0,(_249_v88).dtor_cf30);
        _250_v89 = (_250_v89).update(_128_v0, (new BigNumber(588)).isLessThanOrEqualTo(_128_v0));
      }
      let _251_v90;
      _251_v90 = _dafny.Seq.UnicodeFromString("kpq");
      _251_v90 = _dafny.Seq.UnicodeFromString("mg");
      if (_130_v2) {
        _128_v0 = _128_v0;
        let _252_v91;
        let _nw44 = new _module.C1();
        _nw44.__ctor(_130_v2);
        _252_v91 = _nw44;
        let _253_v92;
        _253_v92 = _dafny.Seq.of(_252_v91);
        let _254_v93;
        let _nw45 = Array((new BigNumber(27)).toNumber());
        _nw45[(_dafny.ZERO).toNumber()] = _252_v91;
        _nw45[(_dafny.ONE).toNumber()] = _252_v91;
        _nw45[(new BigNumber(2)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(3)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(4)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(5)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(6)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(7)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(8)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(9)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(10)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(11)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(12)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(13)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(14)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(15)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(16)).toNumber()] = (_253_v92)[_module.__default.safeIndex(_128_v0, new BigNumber((_253_v92).length))];
        _nw45[(new BigNumber(17)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(18)).toNumber()] = (_253_v92)[_module.__default.safeIndex(_128_v0, new BigNumber((_253_v92).length))];
        _nw45[(new BigNumber(19)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(20)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(21)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(22)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(23)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(24)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(25)).toNumber()] = _252_v91;
        _nw45[(new BigNumber(26)).toNumber()] = _252_v91;
        _254_v93 = _nw45;
        let _255_v94;
        _255_v94 = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,_130_v2);
        let _256_v95;
        _256_v95 = _dafny.Seq.of(_254_v93);
        let _257_v96;
        _257_v96 = _dafny.Map.Empty.slice().updateUnsafe(_215_v61,_128_v0);
        let _258_v97;
        _258_v97 = _module.D2.create_DC9(_128_v0, _128_v0, _251_v90, _130_v2, _257_v96);
        let _259_v98;
        _259_v98 = _dafny.Map.Empty.slice().updateUnsafe(_258_v97,_130_v2);
        let _260_v99;
        let _nw46 = Array((new BigNumber(25)).toNumber());
        _nw46[(_dafny.ZERO).toNumber()] = _255_v94;
        _nw46[(_dafny.ONE).toNumber()] = _255_v94;
        _nw46[(new BigNumber(2)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(3)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_254_v93,_130_v2)).update((_256_v95)[_module.__default.safeIndex(_128_v0, new BigNumber((_256_v95).length))], _130_v2);
        _nw46[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,(((_259_v98).contains(_module.D2.create_DC9(new BigNumber((_179_v36).cardinality()), new BigNumber(94), _251_v90, _130_v2, _257_v96))) ? ((_259_v98).get(_module.D2.create_DC9(new BigNumber((_179_v36).cardinality()), new BigNumber(94), _251_v90, _130_v2, _257_v96))) : (_130_v2)));
        _nw46[(new BigNumber(6)).toNumber()] = ((_130_v2) ? (_255_v94) : (_255_v94));
        _nw46[(new BigNumber(7)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,_130_v2);
        _nw46[(new BigNumber(9)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(10)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(11)).toNumber()] = (_255_v94).Merge(_255_v94);
        _nw46[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,(_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))]);
        _nw46[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,_130_v2);
        _nw46[(new BigNumber(14)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_254_v93,false)).update(_254_v93, _130_v2);
        _nw46[(new BigNumber(16)).toNumber()] = (_255_v94).update((_module.D24.create_DC67(_254_v93)).dtor_cf116, _130_v2);
        _nw46[(new BigNumber(17)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(18)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_254_v93,_130_v2);
        _nw46[(new BigNumber(20)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(21)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(22)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(23)).toNumber()] = _255_v94;
        _nw46[(new BigNumber(24)).toNumber()] = _255_v94;
        _260_v99 = _nw46;
        let _index31 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_260_v99).length));
        (_260_v99)[_index31] = _255_v94;
        let _261_v100;
        _261_v100 = _module.D24.create_DC67(_254_v93);
        let _index32 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_260_v99).length));
        (_260_v99)[_index32] = (_255_v94).update((_261_v100).dtor_cf116, _130_v2);
        let _262_v101;
        let _init7 = ((_263_v90) => function (_264_i14) {
          return (_264_i14).plus(new BigNumber((_263_v90).length));
        })(_251_v90);
        let _nw47 = Array((new BigNumber(20)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw47.length); _i0_7++) {
          _nw47[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _262_v101 = _nw47;
        _262_v101 = _262_v101;
        let _265_v102;
        _265_v102 = _dafny.Map.Empty.slice().updateUnsafe((_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))],_254_v93);
        let _266_v103;
        _266_v103 = _module.D22.create_DC62(!(_130_v2), _128_v0);
        let _267_v104;
        let _nw48 = Array((new BigNumber(26)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _254_v93;
        _nw48[(_dafny.ONE).toNumber()] = _254_v93;
        _nw48[(new BigNumber(2)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(3)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(4)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(5)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(6)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(7)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(8)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(9)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(10)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(11)).toNumber()] = (_256_v95)[_module.__default.safeIndex(_128_v0, new BigNumber((_256_v95).length))];
        _nw48[(new BigNumber(12)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(13)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(14)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(15)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(16)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(17)).toNumber()] = ((_130_v2) ? (_254_v93) : (_254_v93));
        _nw48[(new BigNumber(18)).toNumber()] = (_261_v100).dtor_cf116;
        _nw48[(new BigNumber(19)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(20)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(21)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(22)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(23)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(24)).toNumber()] = _254_v93;
        _nw48[(new BigNumber(25)).toNumber()] = (((_265_v102).contains((_266_v103).dtor_cf106)) ? ((_265_v102).get((_266_v103).dtor_cf106)) : (_254_v93));
        _267_v104 = _nw48;
        let _index33 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_267_v104).length));
        (_267_v104)[_index33] = ((_130_v2) ? (_254_v93) : (_254_v93));
        let _268_v105;
        let _init8 = ((_269_v0) => function (_270_i15) {
          return _module.D2.create_DC7(_dafny.Set.fromElements(_269_v0, _269_v0, _269_v0, _269_v0));
        })(_128_v0);
        let _nw49 = Array((new BigNumber(7)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw49.length); _i0_8++) {
          _nw49[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _268_v105 = _nw49;
        let _271_v106;
        _271_v106 = _module.D22.create_DC61(_268_v105);
        let _272_v107;
        _272_v107 = _module.D17.create_DC48(_130_v2, !(_130_v2), _130_v2);
        let _pat_let_tv6 = _130_v2;
        let _273_v108;
        let _nw50 = Array((new BigNumber(9)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _272_v107;
        _nw50[(_dafny.ONE).toNumber()] = _272_v107;
        _nw50[(new BigNumber(2)).toNumber()] = _272_v107;
        _nw50[(new BigNumber(3)).toNumber()] = function (_pat_let11_0) {
          return function (_274_dt__update__tmp_h2) {
            return function (_pat_let12_0) {
              return function (_275_dt__update_hcf77_h0) {
                return _module.D17.create_DC48(_275_dt__update_hcf77_h0, (_274_dt__update__tmp_h2).dtor_cf78, (_274_dt__update__tmp_h2).dtor_cf79);
              }(_pat_let12_0);
            }(_pat_let_tv6);
          }(_pat_let11_0);
        }(_272_v107);
        _nw50[(new BigNumber(4)).toNumber()] = _272_v107;
        _nw50[(new BigNumber(5)).toNumber()] = _272_v107;
        _nw50[(new BigNumber(6)).toNumber()] = ((_130_v2) ? (_272_v107) : (_272_v107));
        _nw50[(new BigNumber(7)).toNumber()] = ((!(_130_v2)) ? (_272_v107) : (_272_v107));
        _nw50[(new BigNumber(8)).toNumber()] = _module.D17.create_DC48(_130_v2, _130_v2, _130_v2);
        _273_v108 = _nw50;
        let _index34 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_273_v108).length));
        (_273_v108)[_index34] = _272_v107;
        let _276_v109;
        _276_v109 = new _dafny.CodePoint('m'.codePointAt(0));
        let _277_v110;
        _277_v110 = _dafny.Set.fromElements(_130_v2);
        let _278_v111;
        _278_v111 = _dafny.Map.Empty.slice().updateUnsafe(_277_v110,_128_v0);
        let _279_v112;
        let _nw51 = Array((new BigNumber(23)).toNumber()).fill(false);
        _279_v112 = _nw51;
        let _280_v113;
        _280_v113 = _module.D3.create_DC10(_279_v112);
        let _pat_let_tv7 = _279_v112;
        let _281_v114;
        _281_v114 = _dafny.Set.fromElements((function (_pat_let13_0) {
          return function (_282_dt__update__tmp_h3) {
            return function (_pat_let14_0) {
              return function (_283_dt__update_hcf21_h0) {
                return _module.D3.create_DC10(_283_dt__update_hcf21_h0);
              }(_pat_let14_0);
            }(_pat_let_tv7);
          }(_pat_let13_0);
        }(_280_v113)).dtor_cf21);
        let _284_v115;
        _284_v115 = _dafny.Seq.of((((_278_v111).contains(_277_v110)) ? ((_278_v111).get(_277_v110)) : (new BigNumber(-10))), new BigNumber((_281_v114).length));
        let _285_v116;
        _285_v116 = _dafny.Map.Empty.slice().updateUnsafe(false,_128_v0);
        let _286_v117;
        _286_v117 = _module.D24.create_DC68(new BigNumber(79), _130_v2, _285_v116);
        let _pat_let_tv8 = _130_v2;
        let _index35 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_267_v104).length));
        let _index36 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_273_v108).length));
        let _rhs16 = _254_v93;
        let _rhs17 = _271_v106;
        let _rhs18 = _module.__default.fm63(_dafny.Seq.update(_284_v115, _module.__default.safeIndex(_128_v0, new BigNumber((_284_v115).length)), _128_v0), (function (_pat_let15_0) {
          return function (_287_dt__update__tmp_h4) {
            return function (_pat_let16_0) {
              return function (_288_dt__update_hcf118_h0) {
                return _module.D24.create_DC68((_287_dt__update__tmp_h4).dtor_cf117, _288_dt__update_hcf118_h0, (_287_dt__update__tmp_h4).dtor_cf119);
              }(_pat_let16_0);
            }(_pat_let_tv8);
          }(_pat_let15_0);
        }(_286_v117)).dtor_cf118, _130_v2, _136_globalState);
        let _rhs19 = _276_v109;
        let _lhs10 = _267_v104;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_267_v104).length));
        let _lhs12 = _273_v108;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_273_v108).length));
        _lhs10[_lhs11] = _rhs16;
        _271_v106 = _rhs17;
        _lhs12[_lhs13] = _rhs18;
        _276_v109 = _rhs19;
        _128_v0 = _module.__default.safeDivisionInt(_128_v0, (_284_v115)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_284_v115).length))]);
      } else {
        _module.__default.m0(_128_v0, _128_v0, _136_globalState);
        let _289_v118;
        let _nw52 = new _module.C6();
        _nw52.__ctor(_130_v2, _130_v2);
        _289_v118 = _nw52;
        _131_v3 = _131_v3;
        let _290_v119;
        let _nw53 = new _module.C6();
        _nw53.__ctor(_130_v2, _130_v2);
        _290_v119 = _nw53;
        let _291_v120;
        _291_v120 = _dafny.Seq.of(_128_v0);
        let _292_v121;
        let _nw54 = Array((new BigNumber(2)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = (_291_v120)[_module.__default.safeIndex((_dafny.ZERO).minus(_128_v0), new BigNumber((_291_v120).length))];
        _nw54[(_dafny.ONE).toNumber()] = new BigNumber((_291_v120).length);
        _292_v121 = _nw54;
        let _index37 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_292_v121).length));
        (_292_v121)[_index37] = _128_v0;
        let _index38 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_292_v121).length));
        (_292_v121)[_index38] = (((_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))]) ? (new BigNumber((_251_v90).length)) : (_128_v0));
      }
      let _293_v122;
      _293_v122 = _dafny.Map.Empty.slice().updateUnsafe(_130_v2,_128_v0);
      let _294_v123;
      let _nw55 = new _module.C0();
      _nw55.__ctor(_130_v2, _130_v2);
      _294_v123 = _nw55;
      let _295_v124;
      let _nw56 = Array((new BigNumber(5)).toNumber());
      _nw56[(_dafny.ZERO).toNumber()] = _294_v123;
      _nw56[(_dafny.ONE).toNumber()] = _294_v123;
      _nw56[(new BigNumber(2)).toNumber()] = _294_v123;
      _nw56[(new BigNumber(3)).toNumber()] = _294_v123;
      _nw56[(new BigNumber(4)).toNumber()] = _294_v123;
      _295_v124 = _nw56;
      let _296_v125;
      _296_v125 = _dafny.Map.Empty.slice().updateUnsafe(_295_v124,(_294_v123).f12);
      let _297_v126;
      _297_v126 = _dafny.Seq.of(new BigNumber(-882));
      let _298_v127;
      _298_v127 = _dafny.Seq.of(new BigNumber((_297_v126).length));
      let _299_v128;
      _299_v128 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber((_131_v3).length)),new BigNumber((_298_v127).length));
      let _300_v129;
      _300_v129 = _module.D2.create_DC9(_128_v0, new BigNumber((_296_v125).length), _251_v90, _294_v123.f13, _299_v128);
      _293_v122 = (_293_v122).update((_300_v129).dtor_cf19, _128_v0);
      let _301_v130;
      let _nw57 = Array((new BigNumber(3)).toNumber());
      _nw57[(_dafny.ZERO).toNumber()] = _294_v123.f13;
      _nw57[(_dafny.ONE).toNumber()] = _294_v123.f13;
      _nw57[(new BigNumber(2)).toNumber()] = _294_v123.f13;
      _301_v130 = _nw57;
      let _302_v131;
      let _init9 = function (_303_i16) {
        return _module.__default.safeModuloInt(_303_i16, new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).length));
      };
      let _nw58 = Array((new BigNumber(12)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw58.length); _i0_9++) {
        _nw58[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _302_v131 = _nw58;
      let _304_v132;
      _304_v132 = _module.D21.create_DC59(_301_v130, _251_v90, _302_v131, _130_v2);
      let _305_v133;
      _305_v133 = _dafny.Set.fromElements((_304_v132).dtor_cf103);
      let _306_v134;
      _306_v134 = _dafny.Seq.of(_305_v133);
      if (_dafny.Seq.IsProperPrefixOf(_306_v134, _306_v134)) {
        let _index39 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length));
        (_301_v130)[_index39] = _294_v123.f13;
        let _index40 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length));
        (_301_v130)[_index40] = ((_294_v123).f12) === (!((_294_v123).f12));
        _128_v0 = _128_v0;
        if ((((_130_v2) ? (_294_v123.f13) : ((_301_v130)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length))]))) || (_294_v123.f13)) {
          let _307_v135;
          let _init10 = ((_308_v0) => function (_309_i17) {
            return _module.D2.create_DC7((_module.D2.create_DC7(_dafny.Set.fromElements(_308_v0))).dtor_cf13);
          })(_128_v0);
          let _nw59 = Array((new BigNumber(29)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw59.length); _i0_10++) {
            _nw59[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _307_v135 = _nw59;
          let _310_v136;
          let _nw60 = new _module.C4();
          _nw60.__ctor(_307_v135, false);
          _310_v136 = _nw60;
          let _311_v137;
          _311_v137 = _dafny.Map.Empty.slice().updateUnsafe((_294_v123).f12,_310_v136);
          _311_v137 = (_311_v137).update((_301_v130)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length))], _310_v136);
          _130_v2 = (_131_v3)[_module.__default.safeIndex(_128_v0, new BigNumber((_131_v3).length))];
          let _index41 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length));
          (_302_v131)[_index41] = _128_v0;
          let _index42 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length));
          (_302_v131)[_index42] = _module.__default.safeDivisionInt(_128_v0, _128_v0);
          let _312_v138;
          _312_v138 = new _dafny.CodePoint('j'.codePointAt(0));
          (_294_v123).f13 = (_module.D6.create_DC17(_294_v123.f13, new BigNumber(642), _312_v138, _312_v138)).dtor_cf31;
          let _index43 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length));
          let _index44 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length));
          let _rhs20 = (_294_v123).f12;
          let _rhs21 = !(!(!((((_dafny.ZERO).minus((_302_v131)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length))])).plus((_310_v136).fm15(_136_globalState))).isLessThanOrEqualTo(_128_v0))));
          let _rhs22 = !(((_dafny.ZERO).minus((_302_v131)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length))])).minus(new BigNumber(485))).isEqualTo(_128_v0);
          let _rhs23 = _128_v0;
          let _lhs14 = _301_v130;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length));
          let _lhs16 = _294_v123;
          let _lhs17 = _302_v131;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_302_v131).length));
          _lhs14[_lhs15] = _rhs20;
          _130_v2 = _rhs21;
          _lhs16.f13 = _rhs22;
          _lhs17[_lhs18] = _rhs23;
        } else {
          let _313_v139;
          _313_v139 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_128_v0);
          (_294_v123).f13 = (_179_v36).equals(_dafny.MultiSet.fromElements(_128_v0, _128_v0, new BigNumber(334), (((_313_v139).contains((_dafny.ZERO).minus(_128_v0))) ? ((_313_v139).get((_dafny.ZERO).minus(_128_v0))) : (_128_v0))));
          let _314_v140;
          let _nw61 = Array((new BigNumber(24)).toNumber()).fill([]);
          _314_v140 = _nw61;
          let _315_v141;
          let _init11 = ((_316_v123) => function (_317_i18) {
            return (_316_v123).f12;
          })(_294_v123);
          let _nw62 = Array((new BigNumber(14)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw62.length); _i0_11++) {
            _nw62[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _315_v141 = _nw62;
          let _index45 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_314_v140).length));
          (_314_v140)[_index45] = _315_v141;
          let _318_v142;
          _318_v142 = _module.D15.create_DC43(_128_v0);
          let _319_v143;
          let _nw63 = Array((new BigNumber(13)).toNumber()).fill(_module.D2.Default());
          _319_v143 = _nw63;
          let _320_v144;
          let _nw64 = new _module.C4();
          _nw64.__ctor(_319_v143, !((_301_v130)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length))]));
          _320_v144 = _nw64;
          let _321_v145;
          _321_v145 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_320_v144,_320_v144.f7));
          let _index46 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_314_v140).length));
          let _rhs24 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sqlveg"), _dafny.Seq.Concat(_251_v90, _251_v90))).length);
          let _rhs25 = _module.__default.safeModuloInt((_318_v142).dtor_cf70, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_321_v145, _321_v145)).length)));
          let _rhs26 = _315_v141;
          let _rhs27 = (_294_v123).f12;
          let _lhs19 = _314_v140;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_314_v140).length));
          let _lhs21 = _294_v123;
          _128_v0 = _rhs24;
          _128_v0 = _rhs25;
          _lhs19[_lhs20] = _rhs26;
          _lhs21.f13 = _rhs27;
          _128_v0 = (_dafny.ZERO).minus(_128_v0);
          _130_v2 = (_130_v2) || ((_128_v0).isLessThan(_128_v0));
          let _322_v146;
          _322_v146 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs28 = _322_v146;
          let _rhs29 = _128_v0;
          let _rhs30 = _302_v131;
          _322_v146 = _rhs28;
          _128_v0 = _rhs29;
          _302_v131 = _rhs30;
        }
        let _323_v147;
        _323_v147 = _dafny.Seq.of(_215_v61, _215_v61);
        let _324_v148;
        let _nw65 = Array((new BigNumber(5)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = _130_v2;
        _nw65[(_dafny.ONE).toNumber()] = _294_v123.f13;
        _nw65[(new BigNumber(2)).toNumber()] = false;
        _nw65[(new BigNumber(3)).toNumber()] = (_294_v123).f12;
        _nw65[(new BigNumber(4)).toNumber()] = _130_v2;
        _324_v148 = _nw65;
        let _325_v149;
        _325_v149 = _dafny.Set.fromElements(_324_v148);
        let _326_v150;
        _326_v150 = new _dafny.CodePoint('e'.codePointAt(0));
        let _327_v151;
        let _nw66 = new _module.C2();
        _nw66.__ctor(_325_v149, _326_v150, false);
        _327_v151 = _nw66;
        let _328_v152;
        _328_v152 = _dafny.Seq.of(_327_v151);
        let _index47 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_301_v130).length));
        (_301_v130)[_index47] = _module.__default.fm1(_323_v147, _module.__default.safeModuloInt(new BigNumber((_328_v152).length), _128_v0), (new BigNumber(861)).minus(_128_v0), _136_globalState);
        let _329_v153;
        let _nw67 = Array((new BigNumber(21)).toNumber()).fill(_module.D2.Default());
        _329_v153 = _nw67;
        let _330_v154;
        let _nw68 = new _module.C4();
        _nw68.__ctor(_329_v153, _294_v123.f13);
        _330_v154 = _nw68;
      } else {
        let _331_v155;
        let _nw69 = new _module.C6();
        _nw69.__ctor(_294_v123.f13, (_294_v123.f13) || ((_294_v123).f12));
        _331_v155 = _nw69;
        let _nw70 = new _module.C6();
        _nw70.__ctor(_130_v2, (_305_v133).IsSubsetOf(_dafny.Set.fromElements(_294_v123.f13)));
        _331_v155 = _nw70;
        let _index48 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length));
        (_301_v130)[_index48] = true;
        let _332_v156;
        _332_v156 = _dafny.Set.fromElements(_301_v130, _301_v130);
        let _index49 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length));
        let _rhs31 = !(!(_128_v0).isEqualTo(((_294_v123.f13) ? (_128_v0) : (_128_v0))));
        let _rhs32 = (_dafny.Set.fromElements(_301_v130)).IsProperSubsetOf(_332_v156);
        let _lhs22 = _294_v123;
        let _lhs23 = _301_v130;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length));
        _lhs22.f13 = _rhs31;
        _lhs23[_lhs24] = _rhs32;
        _131_v3 = _dafny.Seq.update(_131_v3, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ivahwxl")).length), new BigNumber((_131_v3).length)), (_301_v130)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length))]);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_251_v90, _251_v90), _dafny.Seq.Concat(_251_v90, _251_v90))) {
          let _rhs33 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_128_v0)).plus(_128_v0));
          let _rhs34 = _294_v123.f13;
          let _rhs35 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_333_i19) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), _dafny.Seq.Concat(_251_v90, _251_v90));
          _128_v0 = _rhs33;
          _130_v2 = _rhs34;
          _251_v90 = _rhs35;
          let _334_v157;
          _334_v157 = _module.D12.create_DC30(_128_v0, (_331_v155).f9, _297_v126, _128_v0);
          let _335_v158;
          _335_v158 = _dafny.Seq.of(_334_v157, _334_v157, _334_v157, _334_v157, _334_v157);
          let _336_v159;
          _336_v159 = _dafny.Map.Empty.slice().updateUnsafe(_335_v158,_294_v123.f13);
          let _337_v160;
          _337_v160 = _dafny.Seq.of(_293_v122, _293_v122);
          _336_v159 = (_336_v159).update(_dafny.Seq.of(_334_v157), _dafny.Seq.IsProperPrefixOf(_337_v160, _337_v160));
          let _338_v161;
          let _init12 = ((_339_v3, _340_v0, _341_v130, _342_v123) => function (_343_i20) {
            return (((_339_v3)[_module.__default.safeIndex(_340_v0, new BigNumber((_339_v3).length))]) ? ((_341_v130)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_341_v130).length))]) : ((_342_v123).f12));
          })(_131_v3, _128_v0, _301_v130, _294_v123);
          let _nw71 = Array((new BigNumber(10)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw71.length); _i0_12++) {
            _nw71[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _338_v161 = _nw71;
          let _344_v162;
          _344_v162 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_128_v0);
          let _rhs36 = new BigNumber(((_344_v162).Merge(_344_v162)).length);
          let _rhs37 = _338_v161;
          _128_v0 = _rhs36;
          _338_v161 = _rhs37;
          (_294_v123).f13 = true;
          let _345_v163;
          _345_v163 = _module.D13.create_DC35();
          let _346_v164;
          _346_v164 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(539),((_130_v2) ? (_345_v163) : (_345_v163)));
          let _347_v165;
          _347_v165 = _dafny.Map.Empty.slice().updateUnsafe(_294_v123.f13,_297_v126);
          let _348_v166;
          let _nw72 = new _module.C7();
          _nw72.__ctor(_347_v165, (_294_v123).f12);
          _348_v166 = _nw72;
          let _349_v167;
          let _nw73 = Array((new BigNumber(2)).toNumber());
          _nw73[(_dafny.ZERO).toNumber()] = _348_v166;
          _nw73[(_dafny.ONE).toNumber()] = _348_v166;
          _349_v167 = _nw73;
          let _350_v168;
          _350_v168 = _dafny.Seq.of(_349_v167, _349_v167, _349_v167);
          let _351_v169;
          let _nw74 = new _module.C1();
          _nw74.__ctor(_348_v166.f7);
          _351_v169 = _nw74;
          let _352_v170;
          _352_v170 = _dafny.MultiSet.fromElements(_351_v169);
          let _353_v171;
          _353_v171 = _dafny.Map.Empty.slice().updateUnsafe(_294_v123.f13,_351_v169);
          let _354_v172;
          let _nw75 = Array((new BigNumber(28)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _349_v167;
          _nw75[(_dafny.ONE).toNumber()] = _349_v167;
          _nw75[(new BigNumber(2)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(3)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(4)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(5)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(6)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(7)).toNumber()] = (_350_v168)[_module.__default.safeIndex((((_352_v170).contains((((_353_v171).contains((_331_v155).f9)) ? ((_353_v171).get((_331_v155).f9)) : (_351_v169)))) ? ((_352_v170).get((((_353_v171).contains((_331_v155).f9)) ? ((_353_v171).get((_331_v155).f9)) : (_351_v169)))) : (new BigNumber(-16))), new BigNumber((_350_v168).length))];
          _nw75[(new BigNumber(8)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(9)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(10)).toNumber()] = (_350_v168)[_module.__default.safeIndex((((_344_v162).contains(_128_v0)) ? ((_344_v162).get(_128_v0)) : (_128_v0)), new BigNumber((_350_v168).length))];
          _nw75[(new BigNumber(11)).toNumber()] = (((_294_v123).f12) ? (_349_v167) : (_349_v167));
          _nw75[(new BigNumber(12)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(13)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(14)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(15)).toNumber()] = (((_331_v155).f9) ? (_349_v167) : (_349_v167));
          _nw75[(new BigNumber(16)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(17)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(18)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(19)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(20)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(21)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(22)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(23)).toNumber()] = (((_331_v155).f9) ? (_349_v167) : (_349_v167));
          _nw75[(new BigNumber(24)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(25)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(26)).toNumber()] = _349_v167;
          _nw75[(new BigNumber(27)).toNumber()] = _349_v167;
          _354_v172 = _nw75;
          let _index50 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_354_v172).length));
          (_354_v172)[_index50] = _349_v167;
          let _355_v174;
          _355_v174 = _dafny.Set.fromElements(new BigNumber(48), _128_v0);
          let _356_v177;
          _356_v177 = new _dafny.CodePoint('x'.codePointAt(0));
          let _357_v178;
          _357_v178 = _dafny.Map.Empty.slice().updateUnsafe(_348_v166,new BigNumber(((_module.D25.create_DC70(function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of (function () {
    let _coll16 = new _dafny.Map();
    for (const _compr_16 of (_dafny.Seq.update(_251_v90, _module.__default.safeIndex(new BigNumber((_131_v3).length), new BigNumber((_251_v90).length)), _356_v177)).Elements) {
      let _358_v176 = _compr_16;
      if (_dafny.Seq.contains(_dafny.Seq.update(_251_v90, _module.__default.safeIndex(new BigNumber((_131_v3).length), new BigNumber((_251_v90).length)), _356_v177), _358_v176)) {
        _coll16.push([_358_v176,(_294_v123).f12]);
      }
    }
    return _coll16;
  }()).Keys.Elements) {
    let _359_v175 = _compr_15;
    if ((function () {
      let _coll17 = new _dafny.Map();
      for (const _compr_17 of (_dafny.Seq.update(_251_v90, _module.__default.safeIndex(new BigNumber((_131_v3).length), new BigNumber((_251_v90).length)), _356_v177)).Elements) {
        let _358_v176 = _compr_17;
        if (_dafny.Seq.contains(_dafny.Seq.update(_251_v90, _module.__default.safeIndex(new BigNumber((_131_v3).length), new BigNumber((_251_v90).length)), _356_v177), _358_v176)) {
          _coll17.push([_358_v176,(_294_v123).f12]);
        }
      }
      return _coll17;
    }()).contains(_359_v175)) {
      _coll15.push([_359_v175,_128_v0]);
    }
  }
  return _coll15;
}())).dtor_cf125).length));
          let _index51 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_354_v172).length));
          let _rhs38 = (_128_v0).minus(_128_v0);
          let _rhs39 = ((_346_v164).Merge(function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_355_v174).Elements) {
              let _360_v173 = _compr_18;
              if ((_355_v174).contains(_360_v173)) {
                _coll18.push([(_360_v173).minus(_128_v0),_module.D13.create_DC35()]);
              }
            }
            return _coll18;
          }())).update(new BigNumber((_297_v126).length), _345_v163);
          let _rhs40 = _349_v167;
          let _rhs41 = !((_357_v178).Merge(_357_v178)).contains(_348_v166);
          let _lhs25 = _354_v172;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_354_v172).length));
          let _lhs27 = _294_v123;
          _128_v0 = _rhs38;
          _346_v164 = _rhs39;
          _lhs25[_lhs26] = _rhs40;
          _lhs27.f13 = _rhs41;
        } else {
          _305_v133 = ((_305_v133).Difference(_305_v133)).Difference(_305_v133);
          let _361_v179;
          let _init13 = ((_362_v123, _363_v155) => function (_364_i21) {
            return _dafny.Set.fromElements(_362_v123.f13, (_363_v155).f9);
          })(_294_v123, _331_v155);
          let _nw76 = Array((new BigNumber(15)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw76.length); _i0_13++) {
            _nw76[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _361_v179 = _nw76;
          let _365_v180;
          let _nw77 = Array((new BigNumber(8)).toNumber());
          _365_v180 = _nw77;
          let _366_v181;
          let _nw78 = new _module.C5();
          _nw78.__ctor(_130_v2);
          _366_v181 = _nw78;
          let _367_v182;
          _367_v182 = _module.D26.create_DC73(_366_v181);
          let _368_v183;
          _368_v183 = new _dafny.CodePoint('r'.codePointAt(0));
          let _369_v184;
          let _nw79 = new _module.C2();
          _nw79.__ctor(_332_v156, _368_v183, (_301_v130)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length))]);
          _369_v184 = _nw79;
          let _370_v185;
          _370_v185 = _dafny.Map.Empty.slice().updateUnsafe((_367_v182).dtor_cf129,_369_v184);
          let _index52 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_365_v180).length));
          (_365_v180)[_index52] = (((_370_v185).contains(_366_v181)) ? ((_370_v185).get(_366_v181)) : (_369_v184));
          let _index53 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length));
          let _index54 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_365_v180).length));
          let _rhs42 = _361_v179;
          let _rhs43 = false;
          let _rhs44 = _128_v0;
          let _rhs45 = _369_v184;
          let _lhs28 = _301_v130;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_301_v130).length));
          let _lhs30 = _365_v180;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_365_v180).length));
          _361_v179 = _rhs42;
          _lhs28[_lhs29] = _rhs43;
          _128_v0 = _rhs44;
          _lhs30[_lhs31] = _rhs45;
          let _index55 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          (_302_v131)[_index55] = new BigNumber((_251_v90).length);
          let _371_v186;
          _371_v186 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_128_v0);
          let _index56 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          (_302_v131)[_index56] = new BigNumber(((_371_v186).Merge(_371_v186)).length);
          let _372_v187;
          _372_v187 = _module.D15.create_DC43(new BigNumber((_dafny.Seq.update(_251_v90, _module.__default.safeIndex((_331_v155).fm12((_294_v123).f12, _128_v0, new BigNumber((_297_v126).length), (_294_v123).f12, _136_globalState), new BigNumber((_251_v90).length)), (_369_v184).f16)).length));
          let _index57 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _index58 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _index59 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _rhs46 = _128_v0;
          let _rhs47 = ((_302_v131)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length))]).minus((_302_v131)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length))]);
          let _rhs48 = (_372_v187).dtor_cf70;
          let _rhs49 = !((_331_v155).f9);
          let _rhs50 = _module.__default.fm0(_136_globalState);
          let _lhs32 = _302_v131;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _lhs34 = _302_v131;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _lhs36 = _302_v131;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          let _lhs38 = _294_v123;
          let _lhs39 = _302_v131;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_302_v131).length));
          _lhs32[_lhs33] = _rhs46;
          _lhs34[_lhs35] = _rhs47;
          _lhs36[_lhs37] = _rhs48;
          _lhs38.f13 = _rhs49;
          _lhs39[_lhs40] = _rhs50;
          let _373_v188;
          _373_v188 = _module.D13.create_DC35();
          _373_v188 = _373_v188;
        }
        _302_v131 = _302_v131;
      }
      let _374_v189;
      let _nw80 = new _module.C6();
      _nw80.__ctor(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("tpio"), _251_v90), (_294_v123).f12);
      _374_v189 = _nw80;
      _128_v0 = _module.__default.safeDivisionInt((_128_v0).minus(_128_v0), new BigNumber((_251_v90).length));
      _128_v0 = _128_v0;
      _130_v2 = (_131_v3)[_module.__default.safeIndex((_297_v126)[_module.__default.safeIndex(_128_v0, new BigNumber((_297_v126).length))], new BigNumber((_131_v3).length))];
      let _375_v190;
      _375_v190 = _dafny.Map.Empty.slice().updateUnsafe(_128_v0,_251_v90);
      (_294_v123).f13 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((((_375_v190).contains(_128_v0)) ? ((_375_v190).get(_128_v0)) : (_dafny.Seq.UnicodeFromString("tjbvvw"))), _dafny.Seq.UnicodeFromString("gwgsuo")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), function (_376_i22) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }));
      let _index61 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_302_v131).length));
      (_302_v131)[_index61] = _128_v0;
      let _index62 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_302_v131).length));
      (_302_v131)[_index62] = _128_v0;
      let _377_v191;
      let _nw81 = Array((new BigNumber(13)).toNumber());
      _nw81[(_dafny.ZERO).toNumber()] = _301_v130;
      _nw81[(_dafny.ONE).toNumber()] = _301_v130;
      _nw81[(new BigNumber(2)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(3)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(4)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(5)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(6)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(7)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(8)).toNumber()] = ((_130_v2) ? (_301_v130) : (_301_v130));
      _nw81[(new BigNumber(9)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(10)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(11)).toNumber()] = _301_v130;
      _nw81[(new BigNumber(12)).toNumber()] = _301_v130;
      _377_v191 = _nw81;
      _377_v191 = _377_v191;
      let _rhs51 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), ((_378_v123) => function (_379_i23) {
        return (((_378_v123).f12) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (new _dafny.CodePoint('x'.codePointAt(0))));
      })(_294_v123));
      let _rhs52 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_298_v127, _298_v127));
      let _rhs53 = _302_v131;
      _251_v90 = _rhs51;
      _179_v36 = _rhs52;
      _302_v131 = _rhs53;
      process.stdout.write(_dafny.toString(_128_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_129_v1, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-231))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_130_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_131_v3, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v4).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[_dafny.ZERO], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[_dafny.ONE], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(2)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(3)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(4)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(5)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(6)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(7)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(8)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(9)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(10)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(11)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(12)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(13)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(14)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(15)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(16)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(17)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(18)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(19)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(20)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(21)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(22)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(23)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(24)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(25)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(26)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(27)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_133_v5)[new BigNumber(28)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_136_globalState).f2, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-231))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_globalState).f3).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[_dafny.ZERO], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[_dafny.ONE], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(2)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(3)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(4)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(5)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(6)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(7)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(8)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(9)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(10)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(11)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(12)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(13)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(14)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(15)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(16)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(17)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(18)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(19)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(20)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(21)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(22)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(23)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(24)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(25)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(26)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(27)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_136_globalState).f5)[new BigNumber(28)], _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v36).equals(_dafny.MultiSet.fromElements(_dafny.ONE, _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_215_v61).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v62).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber(462)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_217_v63).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(462), new BigNumber(462), new BigNumber(826))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_251_v90, _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_293_v122).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(462)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_294_v123).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_294_v123.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v124)[_dafny.ZERO]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v124)[_dafny.ZERO].f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v124)[_dafny.ONE]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v124)[_dafny.ONE].f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v124)[new BigNumber(2)]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v124)[new BigNumber(2)].f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v124)[new BigNumber(3)]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v124)[new BigNumber(3)].f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v124)[new BigNumber(4)]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v124)[new BigNumber(4)].f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_296_v125).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_297_v126, _dafny.Seq.of(new BigNumber(-882)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_298_v127, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_299_v128).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.ONE),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v129).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v129).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_300_v129).dtor_cf18).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v129).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_300_v129).dtor_cf20).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_dafny.ONE),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v130)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v130)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v130)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v131)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf100)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf100)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf100)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_304_v132).dtor_cf101).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v132).dtor_cf102)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_304_v132).dtor_cf103));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_305_v133).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_306_v134, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_374_v189).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_375_v190).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Seq.UnicodeFromString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffmgmg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(12)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(12)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_377_v191)[new BigNumber(12)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
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
    static create_DC4(cf7, cf8) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC5(cf9, cf10, cf11) {
      let $dt = new D1(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D1(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false);
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
    static create_DC8(cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC9(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D2(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf14) + ", " + this.cf15.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + this.cf18.toVerbatimString(true) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC11(cf22, cf23, cf24, cf25) {
      let $dt = new D3(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC10(cf21) {
      let $dt = new D3(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf26) {
      let $dt = new D3(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC13(cf27) {
      let $dt = new D4(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27;
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf28) {
      let $dt = new D5(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28);
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
    static create_DC16(cf30) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC17(cf31, cf32, cf33, cf34) {
      let $dt = new D6(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC15(cf29) {
      let $dt = new D6(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC18(cf35) {
      let $dt = new D6(3);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false);
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
    static create_DC19(cf36) {
      let $dt = new D7(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36;
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf38, cf39, cf40) {
      let $dt = new D8(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC22(cf41, cf42, cf43) {
      let $dt = new D8(1);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC20(cf37) {
      let $dt = new D8(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC24() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC25() {
      let $dt = new D9(1);
      return $dt;
    }
    static create_DC26(cf45) {
      let $dt = new D9(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC23(cf44) {
      let $dt = new D9(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24";
      } else if (this.$tag === 1) {
        return "D9.DC25";
      } else if (this.$tag === 2) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf44) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf44 === other.cf44;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24();
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
    static create_DC27(cf46) {
      let $dt = new D10(0);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf47) {
      let $dt = new D11(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC30(cf49, cf50, cf51, cf52) {
      let $dt = new D12(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC31(cf53, cf54) {
      let $dt = new D12(1);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC32(cf55) {
      let $dt = new D12(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC29(cf48) {
      let $dt = new D12(3);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC33(cf56) {
      let $dt = new D12(4);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get is_DC29() { return this.$tag === 3; }
    get is_DC33() { return this.$tag === 4; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 4) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53) && this.cf54 === other.cf54;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf55 === other.cf55;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC30(_dafny.ZERO, false, _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC35() {
      let $dt = new D13(0);
      return $dt;
    }
    static create_DC36(cf58, cf59, cf60) {
      let $dt = new D13(1);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC34(cf57) {
      let $dt = new D13(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC37(cf61) {
      let $dt = new D13(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get is_DC37() { return this.$tag === 3; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35";
      } else if (this.$tag === 1) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf61) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC35();
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
    static create_DC39(cf63, cf64) {
      let $dt = new D14(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC40(cf65, cf66, cf67) {
      let $dt = new D14(1);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC38(cf62) {
      let $dt = new D14(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC41(cf68) {
      let $dt = new D14(3);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC39" + "(" + this.cf63.toVerbatimString(true) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf68, other.cf68);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC39(_dafny.Seq.UnicodeFromString(""), []);
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
    static create_DC43(cf70) {
      let $dt = new D15(0);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC42(cf69) {
      let $dt = new D15(1);
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC44(cf71) {
      let $dt = new D15(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC43(_dafny.ZERO);
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
    static create_DC46(cf73, cf74, cf75) {
      let $dt = new D16(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC45(cf72) {
      let $dt = new D16(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC46(_dafny.ZERO, [], new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC48(cf77, cf78, cf79) {
      let $dt = new D17(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC49(cf80, cf81, cf82, cf83) {
      let $dt = new D17(1);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC47(cf76) {
      let $dt = new D17(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf77 === other.cf77 && this.cf78 === other.cf78 && this.cf79 === other.cf79;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC48(false, false, false);
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
    static create_DC51(cf85, cf86, cf87) {
      let $dt = new D18(0);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC52() {
      let $dt = new D18(1);
      return $dt;
    }
    static create_DC50(cf84) {
      let $dt = new D18(2);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC51" + "(" + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC52";
      } else if (this.$tag === 2) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85) && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC51(_dafny.Seq.of(), _dafny.ZERO, false);
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
    static create_DC53(cf88) {
      let $dt = new D19(0);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC53" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf88, other.cf88);
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf90, cf91, cf92) {
      let $dt = new D20(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC56(cf93, cf94, cf95, cf96, cf97) {
      let $dt = new D20(1);
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC54(cf89) {
      let $dt = new D20(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC55" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC56" + "(" + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC54" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf93, other.cf93) && this.cf94 === other.cf94 && this.cf95 === other.cf95 && _dafny.areEqual(this.cf96, other.cf96) && this.cf97 === other.cf97;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf89 === other.cf89;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC55(false, _dafny.ZERO, []);
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
    static create_DC58(cf99) {
      let $dt = new D21(0);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC59(cf100, cf101, cf102, cf103) {
      let $dt = new D21(1);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC57(cf98) {
      let $dt = new D21(2);
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC60(cf104) {
      let $dt = new D21(3);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get is_DC60() { return this.$tag === 3; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC58" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC59" + "(" + _dafny.toString(this.cf100) + ", " + this.cf101.toVerbatimString(true) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC60" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf99, other.cf99);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf100 === other.cf100 && _dafny.areEqual(this.cf101, other.cf101) && this.cf102 === other.cf102 && this.cf103 === other.cf103;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf98 === other.cf98;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC58(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62(cf106, cf107) {
      let $dt = new D22(0);
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC61(cf105) {
      let $dt = new D22(1);
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC63(cf108) {
      let $dt = new D22(2);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get is_DC63() { return this.$tag === 2; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC62" + "(" + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC61" + "(" + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC63" + "(" + _dafny.toString(this.cf108) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf106 === other.cf106 && _dafny.areEqual(this.cf107, other.cf107);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf105 === other.cf105;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC62(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC65(cf110, cf111) {
      let $dt = new D23(0);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC66(cf112, cf113, cf114, cf115) {
      let $dt = new D23(1);
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC64(cf109) {
      let $dt = new D23(2);
      $dt.cf109 = cf109;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get is_DC66() { return this.$tag === 1; }
    get is_DC64() { return this.$tag === 2; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf109() { return this.cf109; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC65" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC66" + "(" + _dafny.toString(this.cf112) + ", " + _dafny.toString(this.cf113) + ", " + _dafny.toString(this.cf114) + ", " + _dafny.toString(this.cf115) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC64" + "(" + _dafny.toString(this.cf109) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf110, other.cf110) && _dafny.areEqual(this.cf111, other.cf111);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf112 === other.cf112 && this.cf113 === other.cf113 && _dafny.areEqual(this.cf114, other.cf114) && _dafny.areEqual(this.cf115, other.cf115);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf109, other.cf109);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC65(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC68(cf117, cf118, cf119) {
      let $dt = new D24(0);
      $dt.cf117 = cf117;
      $dt.cf118 = cf118;
      $dt.cf119 = cf119;
      return $dt;
    }
    static create_DC69(cf120, cf121, cf122, cf123, cf124) {
      let $dt = new D24(1);
      $dt.cf120 = cf120;
      $dt.cf121 = cf121;
      $dt.cf122 = cf122;
      $dt.cf123 = cf123;
      $dt.cf124 = cf124;
      return $dt;
    }
    static create_DC67(cf116) {
      let $dt = new D24(2);
      $dt.cf116 = cf116;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC69() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf120() { return this.cf120; }
    get dtor_cf121() { return this.cf121; }
    get dtor_cf122() { return this.cf122; }
    get dtor_cf123() { return this.cf123; }
    get dtor_cf124() { return this.cf124; }
    get dtor_cf116() { return this.cf116; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC68" + "(" + _dafny.toString(this.cf117) + ", " + _dafny.toString(this.cf118) + ", " + _dafny.toString(this.cf119) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC69" + "(" + _dafny.toString(this.cf120) + ", " + _dafny.toString(this.cf121) + ", " + _dafny.toString(this.cf122) + ", " + _dafny.toString(this.cf123) + ", " + _dafny.toString(this.cf124) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC67" + "(" + _dafny.toString(this.cf116) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf117, other.cf117) && this.cf118 === other.cf118 && _dafny.areEqual(this.cf119, other.cf119);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf120, other.cf120) && _dafny.areEqual(this.cf121, other.cf121) && _dafny.areEqual(this.cf122, other.cf122) && this.cf123 === other.cf123 && _dafny.areEqual(this.cf124, other.cf124);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf116 === other.cf116;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC68(_dafny.ZERO, false, _dafny.Map.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC71(cf126, cf127) {
      let $dt = new D25(0);
      $dt.cf126 = cf126;
      $dt.cf127 = cf127;
      return $dt;
    }
    static create_DC70(cf125) {
      let $dt = new D25(1);
      $dt.cf125 = cf125;
      return $dt;
    }
    static create_DC72(cf128) {
      let $dt = new D25(2);
      $dt.cf128 = cf128;
      return $dt;
    }
    get is_DC71() { return this.$tag === 0; }
    get is_DC70() { return this.$tag === 1; }
    get is_DC72() { return this.$tag === 2; }
    get dtor_cf126() { return this.cf126; }
    get dtor_cf127() { return this.cf127; }
    get dtor_cf125() { return this.cf125; }
    get dtor_cf128() { return this.cf128; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC71" + "(" + _dafny.toString(this.cf126) + ", " + _dafny.toString(this.cf127) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC70" + "(" + _dafny.toString(this.cf125) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC72" + "(" + _dafny.toString(this.cf128) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf126 === other.cf126 && _dafny.areEqual(this.cf127, other.cf127);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf125, other.cf125);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf128, other.cf128);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC71(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC74() {
      let $dt = new D26(0);
      return $dt;
    }
    static create_DC73(cf129) {
      let $dt = new D26(1);
      $dt.cf129 = cf129;
      return $dt;
    }
    get is_DC74() { return this.$tag === 0; }
    get is_DC73() { return this.$tag === 1; }
    get dtor_cf129() { return this.cf129; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC74";
      } else if (this.$tag === 1) {
        return "D26.DC73" + "(" + _dafny.toString(this.cf129) + ")";
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
        return other.$tag === 1 && this.cf129 === other.cf129;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC74();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC76(cf131, cf132, cf133) {
      let $dt = new D27(0);
      $dt.cf131 = cf131;
      $dt.cf132 = cf132;
      $dt.cf133 = cf133;
      return $dt;
    }
    static create_DC75(cf130) {
      let $dt = new D27(1);
      $dt.cf130 = cf130;
      return $dt;
    }
    get is_DC76() { return this.$tag === 0; }
    get is_DC75() { return this.$tag === 1; }
    get dtor_cf131() { return this.cf131; }
    get dtor_cf132() { return this.cf132; }
    get dtor_cf133() { return this.cf133; }
    get dtor_cf130() { return this.cf130; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC76" + "(" + _dafny.toString(this.cf131) + ", " + _dafny.toString(this.cf132) + ", " + _dafny.toString(this.cf133) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC75" + "(" + _dafny.toString(this.cf130) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf131, other.cf131) && _dafny.areEqual(this.cf132, other.cf132) && _dafny.areEqual(this.cf133, other.cf133);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf130, other.cf130);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC76(_dafny.Set.Empty, new _dafny.CodePoint('D'.codePointAt(0)), new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC78(cf135) {
      let $dt = new D28(0);
      $dt.cf135 = cf135;
      return $dt;
    }
    static create_DC77(cf134) {
      let $dt = new D28(1);
      $dt.cf134 = cf134;
      return $dt;
    }
    static create_DC79(cf136) {
      let $dt = new D28(2);
      $dt.cf136 = cf136;
      return $dt;
    }
    get is_DC78() { return this.$tag === 0; }
    get is_DC77() { return this.$tag === 1; }
    get is_DC79() { return this.$tag === 2; }
    get dtor_cf135() { return this.cf135; }
    get dtor_cf134() { return this.cf134; }
    get dtor_cf136() { return this.cf136; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC78" + "(" + _dafny.toString(this.cf135) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC77" + "(" + _dafny.toString(this.cf134) + ")";
      } else if (this.$tag === 2) {
        return "D28.DC79" + "(" + _dafny.toString(this.cf136) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf135, other.cf135);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf134 === other.cf134;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf136, other.cf136);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC78(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC81(cf138, cf139, cf140) {
      let $dt = new D29(0);
      $dt.cf138 = cf138;
      $dt.cf139 = cf139;
      $dt.cf140 = cf140;
      return $dt;
    }
    static create_DC80(cf137) {
      let $dt = new D29(1);
      $dt.cf137 = cf137;
      return $dt;
    }
    get is_DC81() { return this.$tag === 0; }
    get is_DC80() { return this.$tag === 1; }
    get dtor_cf138() { return this.cf138; }
    get dtor_cf139() { return this.cf139; }
    get dtor_cf140() { return this.cf140; }
    get dtor_cf137() { return this.cf137; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC81" + "(" + _dafny.toString(this.cf138) + ", " + _dafny.toString(this.cf139) + ", " + _dafny.toString(this.cf140) + ")";
      } else if (this.$tag === 1) {
        return "D29.DC80" + "(" + _dafny.toString(this.cf137) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf138, other.cf138) && this.cf139 === other.cf139 && this.cf140 === other.cf140;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf137, other.cf137);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D29.create_DC81(new _dafny.CodePoint('D'.codePointAt(0)), false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
        }
      };
    }
  }

  $module.D30 = class D30 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC83(cf142, cf143, cf144, cf145) {
      let $dt = new D30(0);
      $dt.cf142 = cf142;
      $dt.cf143 = cf143;
      $dt.cf144 = cf144;
      $dt.cf145 = cf145;
      return $dt;
    }
    static create_DC84(cf146, cf147) {
      let $dt = new D30(1);
      $dt.cf146 = cf146;
      $dt.cf147 = cf147;
      return $dt;
    }
    static create_DC82(cf141) {
      let $dt = new D30(2);
      $dt.cf141 = cf141;
      return $dt;
    }
    get is_DC83() { return this.$tag === 0; }
    get is_DC84() { return this.$tag === 1; }
    get is_DC82() { return this.$tag === 2; }
    get dtor_cf142() { return this.cf142; }
    get dtor_cf143() { return this.cf143; }
    get dtor_cf144() { return this.cf144; }
    get dtor_cf145() { return this.cf145; }
    get dtor_cf146() { return this.cf146; }
    get dtor_cf147() { return this.cf147; }
    get dtor_cf141() { return this.cf141; }
    toString() {
      if (this.$tag === 0) {
        return "D30.DC83" + "(" + _dafny.toString(this.cf142) + ", " + _dafny.toString(this.cf143) + ", " + _dafny.toString(this.cf144) + ", " + _dafny.toString(this.cf145) + ")";
      } else if (this.$tag === 1) {
        return "D30.DC84" + "(" + _dafny.toString(this.cf146) + ", " + _dafny.toString(this.cf147) + ")";
      } else if (this.$tag === 2) {
        return "D30.DC82" + "(" + _dafny.toString(this.cf141) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf142, other.cf142) && _dafny.areEqual(this.cf143, other.cf143) && this.cf144 === other.cf144 && _dafny.areEqual(this.cf145, other.cf145);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf146 === other.cf146 && _dafny.areEqual(this.cf147, other.cf147);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf141, other.cf141);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D30.create_DC83(_dafny.Map.Empty, _dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D30.Default();
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
      this._f0 = false;
      this._f1 = false;
      this._f2 = _dafny.Seq.of();
      this._f3 = _dafny.MultiSet.Empty;
      this._f4 = _dafny.ZERO;
      this._f5 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f13 = false;
      this._f12 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      return;
    }
    fm16(p0, globalState) {
      let _this = this;
      return _this.f13;
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      let _source0 = _module.D0.create_DC1(new BigNumber(-379));
      if (_source0.is_DC1) {
        let _380___mcc_h0 = (_source0).cf1;
        let _381_cf1 = _380___mcc_h0;
        return _dafny.Seq.of((_this).f12);
      } else if (_source0.is_DC2) {
        let _382___mcc_h1 = (_source0).cf2;
        let _383___mcc_h2 = (_source0).cf3;
        let _384___mcc_h3 = (_source0).cf4;
        let _385___mcc_h4 = (_source0).cf5;
        let _386_cf5 = _385___mcc_h4;
        let _387_cf4 = _384___mcc_h3;
        let _388_cf3 = _383___mcc_h2;
        let _389_cf2 = _382___mcc_h1;
        return _dafny.Seq.Concat(_dafny.Seq.of(_389_cf2, false), _dafny.Seq.of(_389_cf2));
      } else {
        let _390___mcc_h5 = (_source0).cf0;
        let _391_cf0 = _390___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(_this.f13));
      }
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f7 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f7) {
      let _this = this;
      (_this)._f7 = f7;
      return;
    }
    fm31(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Set.fromElements(_this.f7, true, false);
    };
    fm32(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D2.create_DC9(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7)).length), (_module.D1.create_DC5(_this.f7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-515)), function (_392_i0) {
  return new BigNumber(-556);
}), new BigNumber((_dafny.Seq.UnicodeFromString("ywsvi")).length))).dtor_cf11, _dafny.Seq.UnicodeFromString("nemcnlivt"), _this.f7, _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_module.D1.create_DC5(_this.f7, _dafny.Seq.of(new BigNumber(781), new BigNumber(663)), new BigNumber(-754)))).length))).length))),(_module.D3.create_DC11(new BigNumber(87), new BigNumber((_dafny.MultiSet.fromElements(_this.f7)).cardinality()), new BigNumber(546), _this.f7)).dtor_cf23))).dtor_cf19;
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      let _393_v0;
      _393_v0 = _module.D1.create_DC4(_this.f7, _this.f7);
      let _394_v1;
      _394_v1 = _dafny.Map.Empty.slice().updateUnsafe(_393_v0,p1);
      let _395_v2;
      _395_v2 = _dafny.Set.fromElements(new BigNumber(-86), new BigNumber(-766));
      _394_v1 = (_394_v1).update(function (_pat_let17_0) {
        return function (_396_dt__update__tmp_h0) {
          return function (_pat_let18_0) {
            return function (_397_dt__update_hcf8_h0) {
              return _module.D1.create_DC4((_396_dt__update__tmp_h0).dtor_cf7, _397_dt__update_hcf8_h0);
            }(_pat_let18_0);
          }(false);
        }(_pat_let17_0);
      }(_module.D1.create_DC4(!(p2), p2)), new BigNumber(((_dafny.Set.fromElements(p1, p1, p1)).Union(_395_v2)).length));
      let _398_v3;
      let _nw82 = new _module.C0();
      _nw82.__ctor(!((p2) === (p2)), ((_this.f7) ? (_this.f7) : (p0)));
      _398_v3 = _nw82;
      let _399_v4;
      let _nw83 = Array((new BigNumber(24)).toNumber()).fill(false);
      _399_v4 = _nw83;
      let _index63 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_399_v4).length));
      (_399_v4)[_index63] = !(false) || (!(_398_v3.f13));
      let _400_v5;
      _400_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f7);
      let _index64 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_399_v4).length));
      (_399_v4)[_index64] = (((_400_v5).contains(p1)) ? ((_400_v5).get(p1)) : ((((_398_v3).f12) ? (p0) : ((_398_v3).f12))));
      let _401_i0;
      _401_i0 = _dafny.ZERO;
      L1: {
        while (!(true)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_401_i0)) {
              break L1;
            }
            _401_i0 = (_401_i0).plus(_dafny.ONE);
            let _402_v6;
            _402_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
            let _403_v7;
            _403_v7 = _dafny.Seq.of(p1, p1);
            let _404_v8;
            _404_v8 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_402_v6).length)), _403_v7);
            let _405_v9;
            _405_v9 = _dafny.Seq.of(_404_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-313)), ((_406_v7, _407_p1) => function (_408_i1) {
              return _dafny.Seq.update(_406_v7, _module.__default.safeIndex(new BigNumber(-573), new BigNumber((_406_v7).length)), _407_p1);
            })(_403_v7, p1)), _404_v8);
            let _409_v10;
            _409_v10 = _dafny.Seq.of((_399_v4)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_399_v4).length))]);
            let _410_v11;
            _410_v11 = _dafny.Map.Empty.slice().updateUnsafe(_409_v10,_409_v10);
            let _411_v12;
            _411_v12 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_dafny.Seq.update((_405_v9)[_module.__default.safeIndex(p1, new BigNumber((_405_v9).length))], _module.__default.safeIndex(new BigNumber((_410_v11).length), new BigNumber(((_405_v9)[_module.__default.safeIndex(p1, new BigNumber((_405_v9).length))]).length)), _403_v7));
            _411_v12 = (_411_v12).update(p1, _404_v8);
            let _412_v13;
            _412_v13 = new BigNumber(-438);
            let _413_v14;
            _413_v14 = _dafny.MultiSet.fromElements(p1);
            let _414_v15;
            _414_v15 = new _dafny.CodePoint('n'.codePointAt(0));
            let _415_v16;
            _415_v16 = _dafny.Seq.UnicodeFromString("lrennfvri");
            let _416_v17;
            _416_v17 = _dafny.Map.Empty.slice().updateUnsafe(_415_v16,p1);
            _412_v13 = (((_413_v14).IsProperSubsetOf(_module.__default.fm33(p1, _414_v15, globalState))) ? (p1) : (_module.__default.safeDivisionInt(new BigNumber((_416_v17).length), _412_v13)));
            _412_v13 = (_dafny.ZERO).minus(p1);
            _412_v13 = _module.__default.safeModuloInt(new BigNumber(697), (p1).minus(_412_v13));
          }
        }
      }
      let _417_v18;
      let _nw84 = new _module.C0();
      _nw84.__ctor(p2, _this.f7);
      _417_v18 = _nw84;
      let _418_v19;
      _418_v19 = _dafny.Seq.UnicodeFromString("ppm");
      _418_v19 = _418_v19;
      let _419_v20;
      _419_v20 = _dafny.Map.Empty.slice().updateUnsafe(_417_v18.f13,p2);
      r0 = (_419_v20).Merge(_419_v20);
      let _420_v21;
      _420_v21 = _dafny.Seq.of(_this.f7);
      let _421_v22;
      _421_v22 = _dafny.Set.fromElements(p2, _417_v18.f13, (_398_v3).fm16(false, globalState), (_420_v21)[_module.__default.safeIndex(p1, new BigNumber((_420_v21).length))], (_398_v3).f12);
      let _422_v23;
      _422_v23 = _dafny.Map.Empty.slice().updateUnsafe((_399_v4)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_399_v4).length))],(_421_v22).Intersect(_421_v22));
      r1 = (((_422_v23).contains(p2)) ? ((_422_v23).get(p2)) : (_421_v22));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _source1 = _module.D9.create_DC24();
      if (_source1.is_DC24) {
        let _423_v0;
        _423_v0 = new BigNumber(456);
        let _424_v1;
        _424_v1 = _dafny.Set.fromElements(_423_v0, _423_v0);
        let _425_v2;
        _425_v2 = _dafny.Seq.of(_this.f7);
        let _426_v3;
        _426_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_425_v2).length));
        let _427_v4;
        let _nw85 = Array((new BigNumber(2)).toNumber()).fill(false);
        _427_v4 = _nw85;
        let _428_v5;
        _428_v5 = _dafny.Set.fromElements(_427_v4, _427_v4);
        let _429_v6;
        _429_v6 = _dafny.MultiSet.fromElements(_423_v0, _423_v0);
        let _430_v7;
        _430_v7 = _module.D8.create_DC21(new BigNumber((_module.__default.fm34(_module.__default.fm0(globalState), _this.f7, (_this).fm32(_this.f7, _424_v1, new BigNumber((_426_v3).length), _423_v0, globalState), _423_v0, globalState)).length), _428_v5, new BigNumber((_429_v6).cardinality()));
        _423_v0 = (_423_v0).plus((_430_v7).dtor_cf40);
        _423_v0 = _423_v0;
        (_this).f7 = _this.f7;
        (_this).f7 = _this.f7;
      } else if (_source1.is_DC25) {
        let _431_v8;
        _431_v8 = new BigNumber(319);
        let _432_v9;
        _432_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_431_v8, _431_v8),_this.f7);
        let _433_v10;
        _433_v10 = _dafny.Set.fromElements(_this.f7, true);
        let _434_v11;
        _434_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _435_v12;
        _435_v12 = _dafny.Map.Empty.slice().updateUnsafe(_431_v8,_434_v11);
        let _436_v13;
        _436_v13 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm35(new BigNumber((_433_v10).length), _431_v8, globalState)).length), new BigNumber((_435_v12).length));
        _432_v9 = (_dafny.Map.Empty.slice().updateUnsafe(_431_v8,_this.f7)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),(_this).fm32(false, _436_v13, _431_v8, _431_v8, globalState)));
        let _437_v14;
        _437_v14 = _dafny.Seq.of(_this.f7);
        let _438_v15;
        _438_v15 = _dafny.Seq.of(_431_v8, new BigNumber((_437_v14).length));
        let _439_v16;
        _439_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_438_v15);
        let _440_v17;
        _440_v17 = _dafny.Seq.UnicodeFromString("qkse");
        let _441_v18;
        _441_v18 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, _this.f7, _this.f7, _this.f7);
        let _442_v19;
        let _nw86 = new _module.C0();
        _nw86.__ctor(_this.f7, _this.f7);
        _442_v19 = _nw86;
        let _443_v20;
        _443_v20 = _dafny.Seq.of(_442_v19, _442_v19);
        let _444_v21;
        _444_v21 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_437_v14, _module.__default.safeIndex((((_441_v18).contains(_this.f7)) ? ((_441_v18).get(_this.f7)) : (new BigNumber((_443_v20).length))), new BigNumber((_437_v14).length)), _this.f7), _437_v14), _437_v14);
        let _rhs54 = new BigNumber((_444_v21).length);
        let _rhs55 = _module.__default.fm36(!(_this.f7), (_this.f7) && (_this.f7), globalState);
        let _rhs56 = (((_dafny.Set.fromElements(new BigNumber(972))).IsDisjointFrom(_436_v13)) ? (((false) ? (_440_v17) : (_440_v17))) : (_440_v17));
        let _rhs57 = _dafny.Seq.Concat(_dafny.Seq.update(_440_v17, _module.__default.safeIndex(_431_v8, new BigNumber((_440_v17).length)), _434_v11), _dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), ((_445_v11) => function (_446_i0) {
          return _445_v11;
        })(_434_v11)));
        let _rhs58 = _module.__default.safeModuloInt(new BigNumber((_440_v17).length), _431_v8);
        _431_v8 = _rhs54;
        _439_v16 = _rhs55;
        _440_v17 = _rhs56;
        _440_v17 = _rhs57;
        _431_v8 = _rhs58;
        let _447_v22;
        _447_v22 = _dafny.Map.Empty.slice().updateUnsafe(_431_v8,_441_v18);
        _432_v9 = (_432_v9).update(new BigNumber(((_447_v22)).length), (_431_v8).isLessThan((_dafny.ZERO).minus(_431_v8)));
        if ((_442_v19).f12) {
          (_this).f7 = true;
          _440_v17 = _440_v17;
          let _448_v23;
          _448_v23 = _dafny.Map.Empty.slice().updateUnsafe((_442_v19).fm16(false, globalState),(_dafny.ZERO).minus(_431_v8));
          (_442_v19).f13 = !(((_442_v19.f13) ? ((_442_v19).f12) : ((_442_v19).f12))) || (!((_448_v23).update(!(!((_437_v14)[_module.__default.safeIndex(_431_v8, new BigNumber((_437_v14).length))])), _431_v8)).contains(_442_v19.f13));
          let _449_v24;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_442_v19).f12, (_442_v19).f12);
          _449_v24 = _nw87;
          (_442_v19).f13 = (_442_v19).f12;
        } else {
          let _450_v25;
          let _init14 = ((_451_v15) => function (_452_i1) {
            return _451_v15;
          })(_438_v15);
          let _nw88 = Array((new BigNumber(13)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw88.length); _i0_14++) {
            _nw88[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _450_v25 = _nw88;
          let _index65 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_450_v25).length));
          (_450_v25)[_index65] = _dafny.Seq.of(_431_v8);
          let _453_v26;
          _453_v26 = _dafny.Map.Empty.slice().updateUnsafe(_431_v8,_436_v13);
          let _454_v28;
          _454_v28 = _dafny.Set.fromElements(_436_v13, (((_453_v26).contains(_431_v8)) ? ((_453_v26).get(_431_v8)) : (function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of (_dafny.Set.fromElements(_module.__default.fm0(globalState), _431_v8)).Elements) {
              let _455_v27 = _compr_19;
              if ((_dafny.Set.fromElements(_module.__default.fm0(globalState), _431_v8)).contains(_455_v27)) {
                _coll19.add((_455_v27).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-79),true)).length)));
              }
            }
            return _coll19;
          }())), _436_v13, (_436_v13).Difference(_436_v13));
          let _456_v29;
          let _init15 = function (_457_i2) {
            return false;
          };
          let _nw89 = Array((new BigNumber(23)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw89.length); _i0_15++) {
            _nw89[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _456_v29 = _nw89;
          let _458_v30;
          _458_v30 = _dafny.Map.Empty.slice().updateUnsafe(_436_v13,_442_v19.f13);
          let _459_v32;
          _459_v32 = _dafny.Seq.of(_456_v29, _456_v29, _456_v29, _456_v29);
          let _index66 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_450_v25).length));
          let _rhs59 = _438_v15;
          let _rhs60 = _434_v11;
          let _rhs61 = ((_454_v28).Union(function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of (_458_v30).Keys.Elements) {
              let _460_v31 = _compr_20;
              if ((_458_v30).contains(_460_v31)) {
                _coll20.add(_460_v31);
              }
            }
            return _coll20;
          }())).Intersect(_dafny.Set.fromElements(_436_v13));
          let _rhs62 = _434_v11;
          let _rhs63 = (_459_v32)[_module.__default.safeIndex(_431_v8, new BigNumber((_459_v32).length))];
          let _lhs41 = _450_v25;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_450_v25).length));
          _lhs41[_lhs42] = _rhs59;
          _434_v11 = _rhs60;
          _454_v28 = _rhs61;
          _434_v11 = _rhs62;
          _456_v29 = _rhs63;
          let _461_v33;
          _461_v33 = _dafny.Map.Empty.slice().updateUnsafe(_434_v11,_442_v19.f13);
          let _462_v34;
          _462_v34 = _module.D0.create_DC0(_431_v8);
          let _rhs64 = (_442_v19.f13) === (_this.f7);
          let _rhs65 = _442_v19.f13;
          let _rhs66 = _module.__default.fm1(_dafny.Seq.of(_462_v34), _431_v8, _module.__default.fm0(globalState), globalState);
          let _rhs67 = _461_v33;
          let _rhs68 = _module.__default.fm0(globalState);
          let _lhs43 = _this;
          let _lhs44 = _442_v19;
          let _lhs45 = _442_v19;
          _lhs43.f7 = _rhs64;
          _lhs44.f13 = _rhs65;
          _lhs45.f13 = _rhs66;
          _461_v33 = _rhs67;
          _431_v8 = _rhs68;
          let _index67 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_456_v29).length));
          (_456_v29)[_index67] = _this.f7;
          let _index68 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_456_v29).length));
          (_456_v29)[_index68] = _this.f7;
          let _index69 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_456_v29).length));
          (_456_v29)[_index69] = _dafny.Seq.IsPrefixOf(_440_v17, _440_v17);
          let _463_v35;
          let _nw90 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _463_v35 = _nw90;
          let _index70 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_463_v35).length));
          (_463_v35)[_index70] = (_module.__default.fm0(globalState)).multipliedBy(_431_v8);
          let _index71 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_463_v35).length));
          (_463_v35)[_index71] = _431_v8;
        }
      } else if (_source1.is_DC26) {
        let _464___mcc_h0 = (_source1).cf45;
        let _465_cf45 = _464___mcc_h0;
        _465_cf45 = _465_cf45;
        let _466_v36;
        _466_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,!(_this.f7));
        let _467_v37;
        _467_v37 = new BigNumber(941);
        _466_v36 = (_466_v36).update(!(_467_v37).isEqualTo(_467_v37), _this.f7);
        let _468_v38;
        _468_v38 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(617)).plus(_467_v37),_467_v37);
        _468_v38 = _dafny.Map.Empty.slice().updateUnsafe(_467_v37,new BigNumber(566));
        let _469_v39;
        let _nw91 = Array((new BigNumber(6)).toNumber()).fill(_module.D1.Default());
        _469_v39 = _nw91;
        _469_v39 = _469_v39;
      } else {
        let _470___mcc_h1 = (_source1).cf44;
        let _471_cf44 = _470___mcc_h1;
        let _472_v40;
        _472_v40 = new BigNumber(458);
        let _473_v41;
        _473_v41 = _dafny.Map.Empty.slice().updateUnsafe(_472_v40,_472_v40);
        let _474_v42;
        _474_v42 = _dafny.Seq.UnicodeFromString("befgtwleg");
        let _475_v43;
        let _nw92 = new _module.C0();
        _nw92.__ctor(((((_473_v41).contains(_472_v40)) ? ((_473_v41).get(_472_v40)) : (new BigNumber((_474_v42).length)))).isEqualTo(_472_v40), _this.f7);
        _475_v43 = _nw92;
        let _index72 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_471_cf44).length));
        (_471_cf44)[_index72] = _472_v40;
        let _index73 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_471_cf44).length));
        (_471_cf44)[_index73] = ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_472_v40, _472_v40)).cardinality()))).minus(_472_v40);
        let _476_v44;
        let _nw93 = Array((new BigNumber(21)).toNumber()).fill(_module.D0.Default());
        _476_v44 = _nw93;
        let _477_v45;
        _477_v45 = _module.D0.create_DC0((_471_cf44)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_471_cf44).length))]);
        let _index74 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_476_v44).length));
        (_476_v44)[_index74] = _477_v45;
        let _pat_let_tv9 = _471_cf44;
        let _pat_let_tv10 = _471_cf44;
        let _index75 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_476_v44).length));
        let _rhs69 = function (_pat_let19_0) {
          return function (_478_dt__update__tmp_h0) {
            return function (_pat_let20_0) {
              return function (_479_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_479_dt__update_hcf0_h0);
              }(_pat_let20_0);
            }((_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_pat_let_tv9).length))]);
          }(_pat_let19_0);
        }(_477_v45);
        let _rhs70 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_472_v40), (new BigNumber((_dafny.Seq.UnicodeFromString("edflhdb")).length)).plus(new BigNumber((_474_v42).length)));
        let _lhs46 = _476_v44;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_476_v44).length));
        _lhs46[_lhs47] = _rhs69;
        _472_v40 = _rhs70;
        let _480_v46;
        _480_v46 = new _dafny.CodePoint('w'.codePointAt(0));
        let _481_v47;
        _481_v47 = _dafny.Set.fromElements(_480_v46, _480_v46);
        let _482_v48;
        _482_v48 = _481_v47;
        let _483_v49;
        _483_v49 = _dafny.Map.Empty.slice().updateUnsafe((_482_v48),_475_v43.f13);
        let _484_v50;
        _484_v50 = _dafny.Map.Empty.slice().updateUnsafe((_475_v43.f13) && (true),_483_v49);
        _484_v50 = (_484_v50).update(_this.f7, _483_v49);
      }
      let _485_v51;
      let _nw94 = new _module.C0();
      _nw94.__ctor((false) === (false), _this.f7);
      _485_v51 = _nw94;
      let _486_v52;
      let _nw95 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _486_v52 = _nw95;
      let _487_v53;
      _487_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("gca")).length),_this.f7);
      let _488_v54;
      _488_v54 = new BigNumber(259);
      let _489_v55;
      _489_v55 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37(true, _module.D9.create_DC25(), globalState),(((_487_v53).contains(_488_v54)) ? ((_487_v53).get(_488_v54)) : (_this.f7)));
      let _index76 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
      (_486_v52)[_index76] = (new BigNumber((_489_v55).length)).multipliedBy(_488_v54);
      let _index77 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
      (_486_v52)[_index77] = (_dafny.ZERO).minus((_488_v54).minus(_488_v54));
      let _490_v56;
      let _init16 = ((_491_v51) => function (_492_i3) {
        return _dafny.Seq.of(_dafny.Set.fromElements(_491_v51.f13), _dafny.Set.fromElements(_491_v51.f13, _this.f7, (_491_v51).f12, true), _dafny.Set.fromElements((_491_v51).f12));
      })(_485_v51);
      let _nw96 = Array((new BigNumber(6)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw96.length); _i0_16++) {
        _nw96[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _490_v56 = _nw96;
      let _index78 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_490_v56).length));
      (_490_v56)[_index78] = _dafny.Seq.of(_dafny.Set.fromElements(_485_v51.f13, _this.f7, (_485_v51).f12), _dafny.Set.fromElements(false, (_485_v51).f12));
      let _index79 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_490_v56).length));
      (_490_v56)[_index79] = _module.__default.fm38((_485_v51).f12, globalState);
      let _493_v57;
      _493_v57 = _dafny.Seq.of(_488_v54);
      if (((_493_v57)[_module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_493_v57).length))]).isLessThanOrEqualTo(_488_v54)) {
        let _494_v58;
        _494_v58 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,new BigNumber(238));
        let _495_v59;
        _495_v59 = _dafny.MultiSet.fromElements(_494_v58);
        let _496_v60;
        _496_v60 = _dafny.MultiSet.fromElements((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
        let _497_v61;
        _497_v61 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_496_v60).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(_485_v51.f13,(_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]));
        let _498_v62;
        _498_v62 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_497_v61));
        (_this).f7 = (_495_v59).IsDisjointFrom(((_498_v62)[_module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_498_v62).length))]).Difference(_495_v59));
        let _index80 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
        (_486_v52)[_index80] = _488_v54;
        let _499_v63;
        _499_v63 = _dafny.Seq.of(!(_this.f7));
        (_this).f7 = _dafny.Seq.IsPrefixOf(_499_v63, ((_this.f7) ? (_499_v63) : (_499_v63)));
        let _500_v64;
        _500_v64 = _dafny.Seq.UnicodeFromString("mwepdcn");
        let _501_v65;
        _501_v65 = new _dafny.CodePoint('g'.codePointAt(0));
        let _rhs71 = (_488_v54).isLessThanOrEqualTo(_488_v54);
        let _rhs72 = !(((_485_v51.f13) ? (_485_v51.f13) : (_485_v51.f13)));
        let _rhs73 = ((_485_v51.f13) ? (_485_v51.f13) : (_485_v51.f13));
        let _rhs74 = _dafny.Seq.Concat(_dafny.Seq.update(_500_v64, _module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_500_v64).length)), _501_v65), _500_v64);
        let _rhs75 = ((_this.f7) ? (_501_v65) : (_501_v65));
        let _lhs48 = _this;
        let _lhs49 = _485_v51;
        let _lhs50 = _485_v51;
        _lhs48.f7 = _rhs71;
        _lhs49.f13 = _rhs72;
        _lhs50.f13 = _rhs73;
        _500_v64 = _rhs74;
        _501_v65 = _rhs75;
        _488_v54 = _module.__default.fm0(globalState);
      } else {
        (_485_v51).f13 = (_485_v51).f12;
        let _502_v66;
        let _nw97 = new _module.C0();
        _nw97.__ctor(_485_v51.f13, _485_v51.f13);
        _502_v66 = _nw97;
        let _503_v67;
        _503_v67 = _module.D3.create_DC11(new BigNumber(-973), new BigNumber(984), _488_v54, _502_v66.f13);
        let _504_v68;
        _504_v68 = _module.D3.create_DC12(_503_v67);
        let _pat_let_tv11 = _503_v67;
        _504_v68 = function (_pat_let21_0) {
          return function (_505_dt__update__tmp_h1) {
            return function (_pat_let22_0) {
              return function (_506_dt__update_hcf26_h0) {
                return _module.D3.create_DC12(_506_dt__update_hcf26_h0);
              }(_pat_let22_0);
            }(_pat_let_tv11);
          }(_pat_let21_0);
        }(_504_v68);
        let _507_v69;
        _507_v69 = new _dafny.CodePoint('s'.codePointAt(0));
        let _508_v70;
        _508_v70 = _dafny.Seq.UnicodeFromString("tnehap");
        if (_dafny.Seq.contains(_508_v70, _507_v69)) {
          let _509_v71;
          _509_v71 = _dafny.Set.fromElements(new BigNumber(-964));
          let _510_v72;
          _510_v72 = _dafny.Set.fromElements(_509_v71);
          let _511_v73;
          _511_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_510_v72).length),_488_v54);
          let _512_v74;
          _512_v74 = _dafny.Map.Empty.slice().updateUnsafe((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))],_511_v73);
          _512_v74 = (_512_v74).Merge(_dafny.Map.Empty.slice().updateUnsafe(_488_v54,_511_v73));
          let _513_v75;
          let _nw98 = Array((new BigNumber(21)).toNumber()).fill(false);
          _513_v75 = _nw98;
          let _514_v76;
          _514_v76 = _dafny.Seq.of(_513_v75, _513_v75);
          let _515_v77;
          let _nw99 = Array((new BigNumber(8)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = _502_v66.f13;
          _nw99[(_dafny.ONE).toNumber()] = _485_v51.f13;
          _nw99[(new BigNumber(2)).toNumber()] = (_485_v51).fm16((_485_v51).f12, globalState);
          _nw99[(new BigNumber(3)).toNumber()] = _485_v51.f13;
          _nw99[(new BigNumber(4)).toNumber()] = (_485_v51).fm16((_485_v51).f12, globalState);
          _nw99[(new BigNumber(5)).toNumber()] = (_485_v51).f12;
          _nw99[(new BigNumber(6)).toNumber()] = _502_v66.f13;
          _nw99[(new BigNumber(7)).toNumber()] = !_dafny.Seq.contains(_514_v76, _513_v75);
          _515_v77 = _nw99;
          _513_v75 = _515_v77;
          let _516_v78;
          _516_v78 = _dafny.MultiSet.fromElements(_508_v70);
          let _517_v79;
          _517_v79 = _module.D0.create_DC0((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _518_v80;
          _518_v80 = _dafny.Map.Empty.slice().updateUnsafe(_517_v79,_488_v54);
          let _519_v81;
          _519_v81 = _dafny.Map.Empty.slice().updateUnsafe(true,_507_v69);
          let _520_v82;
          _520_v82 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC9(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-963)), ((_521_v54) => function (_522_i4) {
  return _521_v54;
})(_488_v54))).length), new BigNumber((_516_v78).cardinality()), _508_v70, (_485_v51).f12, _518_v80)).dtor_cf19,_507_v69), _519_v81);
          let _523_v83;
          let _nw100 = new _module.C0();
          _nw100.__ctor((_520_v82).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(_519_v81))), !((_485_v51).f12));
          _523_v83 = _nw100;
          let _524_v84;
          _524_v84 = _dafny.Set.fromElements(_486_v52, _486_v52);
          let _525_v85;
          _525_v85 = _module.D0.create_DC2(((_502_v66).f12) || ((_this).fm32(!((_485_v51).f12), _509_v71, _488_v54, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], globalState)), _module.__default.safeModuloInt((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]), _524_v84, new BigNumber((_509_v71).length));
          _525_v85 = _525_v85;
          let _526_v86;
          _526_v86 = _dafny.Set.fromElements((_485_v51).f12, false);
          let _527_v87;
          _527_v87 = _dafny.Seq.of(!(_502_v66.f13), _502_v66.f13, _485_v51.f13);
          _488_v54 = ((_this.f7) ? (new BigNumber((_526_v86).length)) : ((_488_v54).plus(new BigNumber((_dafny.Set.fromElements(_488_v54, new BigNumber(340), new BigNumber((_527_v87).length))).length))));
        } else {
          (_this).f7 = _this.f7;
          let _528_v88;
          _528_v88 = _module.D1.create_DC4(_this.f7, _485_v51.f13);
          let _529_v89;
          _529_v89 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_485_v51.f13);
          let _530_v90;
          let _nw101 = new _module.C0();
          _nw101.__ctor((_485_v51).f12, !(_dafny.Map.Empty.slice().updateUnsafe((_528_v88).dtor_cf7,_485_v51.f13)).equals(_529_v89));
          _530_v90 = _nw101;
          let _index81 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _rhs76 = true;
          let _rhs77 = _488_v54;
          let _lhs51 = _this;
          let _lhs52 = _486_v52;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          _lhs51.f7 = _rhs76;
          _lhs52[_lhs53] = _rhs77;
          let _531_v91;
          _531_v91 = _dafny.Set.fromElements(_485_v51.f13, _485_v51.f13);
          let _532_v92;
          let _nw102 = new _module.C0();
          _nw102.__ctor(!(_531_v91).contains((_485_v51).f12), _this.f7);
          _532_v92 = _nw102;
          let _533_v93;
          _533_v93 = _dafny.Seq.of(_508_v70);
          let _534_v94;
          _534_v94 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("cr"), (_533_v93)[_module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_533_v93).length))]);
          let _index82 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _rhs78 = (_dafny.ZERO).minus(((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).multipliedBy((_488_v54).multipliedBy((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])));
          let _rhs79 = (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))];
          let _rhs80 = ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).plus(_488_v54);
          let _rhs81 = (_534_v94).IsDisjointFrom(_dafny.Set.fromElements(_508_v70, _508_v70, _508_v70, _508_v70, _dafny.Seq.UnicodeFromString("ct")));
          let _lhs54 = _486_v52;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _lhs56 = _502_v66;
          _488_v54 = _rhs78;
          _lhs54[_lhs55] = _rhs79;
          _488_v54 = _rhs80;
          _lhs56.f13 = _rhs81;
        }
        if (!(((_485_v51).f12) === ((_502_v66).fm16(_485_v51.f13, globalState)))) {
          let _535_v95;
          let _init17 = ((_536_v69) => function (_537_i5) {
            return _536_v69;
          })(_507_v69);
          let _nw103 = Array((new BigNumber(29)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw103.length); _i0_17++) {
            _nw103[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _535_v95 = _nw103;
          let _index83 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_535_v95).length));
          (_535_v95)[_index83] = _507_v69;
          let _index84 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_535_v95).length));
          (_535_v95)[_index84] = _507_v69;
          let _538_v96;
          _538_v96 = _dafny.Set.fromElements(_485_v51.f13, _485_v51.f13);
          let _539_v98;
          _539_v98 = _dafny.MultiSet.fromElements(new BigNumber((_538_v96).length), _488_v54, new BigNumber((function () {
            let _coll21 = new _dafny.Set();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(168), new BigNumber(983))) {
              let _540_v97 = _compr_21;
              if (((new BigNumber(168)).isLessThanOrEqualTo(_540_v97)) && ((_540_v97).isLessThan(new BigNumber(983)))) {
                _coll21.add(_module.__default.safeDivisionInt(_540_v97, _488_v54));
              }
            }
            return _coll21;
          }()).length), _488_v54);
          let _541_v99;
          _541_v99 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,_539_v98);
          let _542_v100;
          _542_v100 = _module.D0.create_DC0(new BigNumber(498));
          let _543_v101;
          _543_v101 = _dafny.Seq.of(_542_v100);
          let _544_v102;
          _544_v102 = _dafny.Set.fromElements(_486_v52);
          let _545_v103;
          _545_v103 = _module.D0.create_DC2(_485_v51.f13, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _544_v102, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _546_v104;
          let _nw104 = Array((new BigNumber(11)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _485_v51.f13;
          _nw104[(_dafny.ONE).toNumber()] = _this.f7;
          _nw104[(new BigNumber(2)).toNumber()] = !(_485_v51.f13);
          _nw104[(new BigNumber(3)).toNumber()] = !((_539_v98).IsDisjointFrom((((_541_v99).contains(_502_v66.f13)) ? ((_541_v99).get(_502_v66.f13)) : (_539_v98))));
          _nw104[(new BigNumber(4)).toNumber()] = !(_485_v51.f13);
          _nw104[(new BigNumber(5)).toNumber()] = (_485_v51).f12;
          _nw104[(new BigNumber(6)).toNumber()] = (_502_v66).f12;
          _nw104[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_543_v101, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber(-545), globalState);
          _nw104[(new BigNumber(8)).toNumber()] = (((_545_v103).dtor_cf2) ? (_485_v51.f13) : (!(_502_v66.f13)));
          _nw104[(new BigNumber(9)).toNumber()] = _this.f7;
          _nw104[(new BigNumber(10)).toNumber()] = (_485_v51).f12;
          _546_v104 = _nw104;
          let _547_v105;
          _547_v105 = _dafny.Seq.of(true);
          let _index85 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_546_v104).length));
          (_546_v104)[_index85] = (_547_v105)[_module.__default.safeIndex(_488_v54, new BigNumber((_547_v105).length))];
          let _548_v106;
          let _init18 = ((_549_v57, _550_v52, _551_v54) => function (_552_i6) {
            return _dafny.Seq.update(_549_v57, _module.__default.safeIndex((_550_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_550_v52).length))], new BigNumber((_549_v57).length)), _551_v54);
          })(_493_v57, _486_v52, _488_v54);
          let _nw105 = Array((new BigNumber(19)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
            _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _548_v106 = _nw105;
          let _553_v107;
          _553_v107 = _dafny.Map.Empty.slice().updateUnsafe((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))],_508_v70);
          let _index86 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_548_v106).length));
          (_548_v106)[_index86] = _dafny.Seq.update(_dafny.Seq.Concat(_493_v57, _493_v57), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((((_553_v107).contains(_488_v54)) ? ((_553_v107).get(_488_v54)) : (_dafny.Seq.update(_508_v70, _module.__default.safeIndex(_488_v54, new BigNumber((_508_v70).length)), (_535_v95)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_535_v95).length))])))).length)), new BigNumber((_dafny.Seq.Concat(_493_v57, _493_v57)).length)), _module.__default.fm0(globalState));
          let _index87 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_546_v104).length));
          let _index89 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_548_v106).length));
          let _rhs82 = new BigNumber((_508_v70).length);
          let _rhs83 = ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).isLessThanOrEqualTo(_488_v54);
          let _rhs84 = _dafny.Seq.Concat(_493_v57, _493_v57);
          let _lhs57 = _486_v52;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _lhs59 = _546_v104;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_546_v104).length));
          let _lhs61 = _548_v106;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_548_v106).length));
          _lhs57[_lhs58] = _rhs82;
          _lhs59[_lhs60] = _rhs83;
          _lhs61[_lhs62] = _rhs84;
          let _554_v108;
          _554_v108 = _module.D3.create_DC10(_546_v104);
          _546_v104 = (_554_v108).dtor_cf21;
          let _555_v109;
          let _init19 = ((_556_v105) => function (_557_i7) {
            return _556_v105;
          })(_547_v105);
          let _nw106 = Array((new BigNumber(2)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw106.length); _i0_19++) {
            _nw106[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _555_v109 = _nw106;
          let _index90 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_555_v109).length));
          (_555_v109)[_index90] = (_485_v51).fm17(new BigNumber((_508_v70).length), (_502_v66).f12, globalState);
          let _index91 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_555_v109).length));
          (_555_v109)[_index91] = (_502_v66).fm17((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _485_v51.f13, globalState);
          _488_v54 = _488_v54;
        } else {
          let _558_v110;
          _558_v110 = _dafny.Seq.of(false);
          let _559_v111;
          _559_v111 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dwb"),_558_v110);
          _488_v54 = new BigNumber((((_559_v111).update(_508_v70, _558_v110)).Merge(_559_v111)).length);
          let _560_v112;
          let _nw107 = Array((new BigNumber(16)).toNumber()).fill(false);
          _560_v112 = _nw107;
          let _561_v113;
          _561_v113 = _dafny.Set.fromElements(_560_v112, _560_v112, _560_v112, _560_v112, _560_v112);
          let _562_v114;
          _562_v114 = _dafny.Map.Empty.slice().updateUnsafe(_507_v69,_561_v113);
          _562_v114 = (_562_v114).update(_507_v69, (_561_v113).Union(_561_v113));
          _488_v54 = ((new BigNumber((_493_v57).length)).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])))).multipliedBy(_488_v54);
          let _index92 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          (_486_v52)[_index92] = ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).minus(_488_v54);
          _module.__default.m0((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _488_v54, globalState);
        }
      }
      let _563_v115;
      _563_v115 = _module.D0.create_DC0((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
      let _564_v116;
      _564_v116 = _dafny.Seq.UnicodeFromString("to");
      let _565_v117;
      _565_v117 = _dafny.Seq.of(_563_v115, _module.D0.create_DC0(new BigNumber((_564_v116).length)), _563_v115);
      let _566_v118;
      _566_v118 = new _dafny.CodePoint('e'.codePointAt(0));
      let _567_v119;
      _567_v119 = _dafny.Set.fromElements(_566_v118, _566_v118, _566_v118, _566_v118, _566_v118);
      let _568_v120;
      _568_v120 = _567_v119;
      let _569_v121;
      _569_v121 = _dafny.Map.Empty.slice().updateUnsafe(_568_v120,(_485_v51).f12);
      if (_module.__default.fm1(_dafny.Seq.Concat(_565_v117, _565_v117), new BigNumber((_569_v121).length), (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], globalState)) {
        let _570_v122;
        _570_v122 = _dafny.Seq.of(_564_v116, _564_v116);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_570_v122, _570_v122), _dafny.Seq.Create(_module.__default.abs(new BigNumber(275)), ((_571_v116) => function (_572_i8) {
          return _571_v116;
        })(_564_v116)))) {
          let _573_v123;
          _573_v123 = _dafny.Map.Empty.slice().updateUnsafe(_563_v115,_488_v54);
          let _574_v124;
          _574_v124 = _module.D2.create_DC9((_488_v54).multipliedBy(_488_v54), _488_v54, _dafny.Seq.UnicodeFromString("v"), (_485_v51).f12, _573_v123);
          let _index93 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _rhs85 = _module.__default.safeDivisionInt(_488_v54, (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])));
          let _rhs86 = _485_v51.f13;
          let _rhs87 = _574_v124;
          let _lhs63 = _486_v52;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _lhs65 = _485_v51;
          _lhs63[_lhs64] = _rhs85;
          _lhs65.f13 = _rhs86;
          _574_v124 = _rhs87;
          let _575_v125;
          _575_v125 = _dafny.Map.Empty.slice().updateUnsafe(((_485_v51).f12) === ((_485_v51).f12),(_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          _488_v54 = (((_575_v125).contains(_485_v51.f13)) ? ((_575_v125).get(_485_v51.f13)) : (((_485_v51.f13) ? (_module.__default.fm0(globalState)) : ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]))));
          (_this).f7 = _this.f7;
          let _576_v126;
          _576_v126 = _module.D2.create_DC8((_485_v51).f12, _564_v116);
          let _rhs88 = _module.__default.safeDivisionInt(_488_v54, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("giukdrl"), _564_v116)).length));
          let _rhs89 = _576_v126;
          let _rhs90 = _485_v51.f13;
          let _rhs91 = _dafny.Seq.Concat(_dafny.Seq.Concat(_564_v116, _564_v116), _dafny.Seq.Concat(_dafny.Seq.update(_564_v116, _module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_564_v116).length)), _566_v118), _564_v116));
          let _lhs66 = _485_v51;
          _488_v54 = _rhs88;
          _576_v126 = _rhs89;
          _lhs66.f13 = _rhs90;
          _564_v116 = _rhs91;
          let _577_v127;
          _577_v127 = _module.D3.create_DC11((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _488_v54, _488_v54, _485_v51.f13);
          let _578_v128;
          _578_v128 = _dafny.Seq.of((_577_v127).dtor_cf25);
          let _579_v129;
          _579_v129 = _dafny.MultiSet.fromElements(_488_v54);
          let _580_v130;
          _580_v130 = _dafny.Seq.of(_579_v129);
          let _581_v131;
          let _nw108 = Array((new BigNumber(24)).toNumber());
          _nw108[(_dafny.ZERO).toNumber()] = true;
          _nw108[(_dafny.ONE).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(2)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(3)).toNumber()] = _485_v51.f13;
          _nw108[(new BigNumber(4)).toNumber()] = (_578_v128)[_module.__default.safeIndex(_488_v54, new BigNumber((_578_v128).length))];
          _nw108[(new BigNumber(5)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(6)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(7)).toNumber()] = _485_v51.f13;
          _nw108[(new BigNumber(8)).toNumber()] = ((_580_v130)[_module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_580_v130).length))]).IsProperSubsetOf((_dafny.MultiSet.FromArray(_493_v57)).update((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _module.__default.abs((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])));
          _nw108[(new BigNumber(9)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(10)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(11)).toNumber()] = _this.f7;
          _nw108[(new BigNumber(12)).toNumber()] = _dafny.areEqual((_module.__default.fm39((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_485_v51).f12, globalState)).dtor_cf10, _dafny.Seq.of(_module.__default.fm0(globalState)));
          _nw108[(new BigNumber(13)).toNumber()] = (_488_v54).isLessThanOrEqualTo(_488_v54);
          _nw108[(new BigNumber(14)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(15)).toNumber()] = _this.f7;
          _nw108[(new BigNumber(16)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(17)).toNumber()] = ((_this.f7) ? ((_485_v51).f12) : ((_485_v51).f12));
          _nw108[(new BigNumber(18)).toNumber()] = _485_v51.f13;
          _nw108[(new BigNumber(19)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(20)).toNumber()] = (new BigNumber((_564_v116).length)).isLessThanOrEqualTo(_488_v54);
          _nw108[(new BigNumber(21)).toNumber()] = (_485_v51).f12;
          _nw108[(new BigNumber(22)).toNumber()] = false;
          _nw108[(new BigNumber(23)).toNumber()] = (_579_v129).IsProperSubsetOf(_dafny.MultiSet.fromElements(_488_v54));
          _581_v131 = _nw108;
          let _index94 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_581_v131).length));
          (_581_v131)[_index94] = (_485_v51).f12;
          let _index95 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_581_v131).length));
          (_581_v131)[_index95] = (_488_v54).isLessThan((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
        } else {
          let _index96 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          (_486_v52)[_index96] = (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))];
          let _582_v132;
          _582_v132 = _dafny.Seq.of((_485_v51).f12);
          let _583_v133;
          _583_v133 = _dafny.MultiSet.fromElements(_582_v132);
          let _584_v134;
          _584_v134 = _module.D2.create_DC8(false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_585_v118) => function (_586_i9) {
  return _585_v118;
})(_566_v118)));
          let _587_v135;
          _587_v135 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm35(new BigNumber((_583_v133).cardinality()), new BigNumber(-39), globalState), (_584_v134).dtor_cf15), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_module.__default.fm35(new BigNumber((_583_v133).cardinality()), new BigNumber(-39), globalState), (_584_v134).dtor_cf15)).length)), new _dafny.CodePoint('h'.codePointAt(0))), _module.__default.safeIndex(_488_v54, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm35(new BigNumber((_583_v133).cardinality()), new BigNumber(-39), globalState), (_584_v134).dtor_cf15), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_module.__default.fm35(new BigNumber((_583_v133).cardinality()), new BigNumber(-39), globalState), (_584_v134).dtor_cf15)).length)), new _dafny.CodePoint('h'.codePointAt(0)))).length)), _566_v118),_493_v57);
          _587_v135 = ((_485_v51.f13) ? ((_587_v135).Merge(_587_v135)) : (_587_v135));
          let _588_v136;
          let _nw109 = new _module.C0();
          _nw109.__ctor(_this.f7, _dafny.Seq.IsProperPrefixOf(_493_v57, _493_v57));
          _588_v136 = _nw109;
          _488_v54 = (_dafny.ZERO).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _index97 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          (_486_v52)[_index97] = (((_588_v136).f12) ? ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]) : ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]));
        }
        if ((_485_v51).f12) {
          let _589_v137;
          _589_v137 = _dafny.Map.Empty.slice().updateUnsafe(true,(_485_v51).f12);
          _589_v137 = (_589_v137).update(_this.f7, _module.__default.fm1(_565_v117, _488_v54, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], globalState));
          let _590_v138;
          _590_v138 = _module.D6.create_DC17((_485_v51).f12, (_dafny.ZERO).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]), _566_v118, _566_v118);
          let _591_v139;
          _591_v139 = _module.D6.create_DC18(_590_v138);
          let _592_v140;
          _592_v140 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,_module.__default.fm40((((_487_v53).contains((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])) ? ((_487_v53).get((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])) : (_this.f7)), _591_v139, _488_v54, globalState));
          let _593_v141;
          _593_v141 = _module.D9.create_DC25();
          _592_v140 = (_592_v140).update((((_487_v53).contains(new BigNumber((_dafny.Set.fromElements(_564_v116)).length))) ? ((_487_v53).get(new BigNumber((_dafny.Set.fromElements(_564_v116)).length))) : (_485_v51.f13)), _593_v141);
          (_this).f7 = false;
          let _594_v142;
          let _nw110 = Array((new BigNumber(16)).toNumber()).fill(false);
          _594_v142 = _nw110;
          let _595_v143;
          _595_v143 = _module.D3.create_DC10(_594_v142);
          let _596_v144;
          _596_v144 = _module.D3.create_DC12(_595_v143);
          let _pat_let_tv12 = _595_v143;
          _596_v144 = function (_pat_let23_0) {
            return function (_597_dt__update__tmp_h2) {
              return function (_pat_let24_0) {
                return function (_598_dt__update_hcf26_h1) {
                  return _module.D3.create_DC12(_598_dt__update_hcf26_h1);
                }(_pat_let24_0);
              }(_pat_let_tv12);
            }(_pat_let23_0);
          }(_596_v144);
          let _599_v145;
          let _nw111 = Array((_dafny.ONE).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = _564_v116;
          _599_v145 = _nw111;
          let _index98 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_599_v145).length));
          (_599_v145)[_index98] = _564_v116;
          let _index99 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_594_v142).length));
          (_594_v142)[_index99] = (_module.__default.fm2((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_485_v51).f12, globalState)).IsSubsetOf(_dafny.Set.fromElements((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]));
          let _index100 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_599_v145).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_594_v142).length));
          let _rhs92 = _564_v116;
          let _rhs93 = _485_v51.f13;
          let _lhs67 = _599_v145;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_599_v145).length));
          let _lhs69 = _594_v142;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_594_v142).length));
          _lhs67[_lhs68] = _rhs92;
          _lhs69[_lhs70] = _rhs93;
        } else {
          let _600_v146;
          let _nw112 = Array((new BigNumber(10)).toNumber()).fill([]);
          _600_v146 = _nw112;
          let _601_v147;
          let _init20 = ((_602_v51) => function (_603_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe(_602_v51.f13,(_602_v51).f12);
          })(_485_v51);
          let _nw113 = Array((new BigNumber(19)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw113.length); _i0_20++) {
            _nw113[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _601_v147 = _nw113;
          let _604_v148;
          _604_v148 = _601_v147;
          let _605_v149;
          let _nw114 = Array((_dafny.ONE).toNumber());
          _nw114[(_dafny.ZERO).toNumber()] = _604_v148;
          _605_v149 = _nw114;
          let _index102 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_600_v146).length));
          (_600_v146)[_index102] = _605_v149;
          let _index103 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_600_v146).length));
          (_600_v146)[_index103] = _605_v149;
          let _606_v150;
          let _init21 = ((_607_v51) => function (_608_i11) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(_607_v51.f13, _607_v51.f13, _607_v51.f13, (_607_v51).f12, true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(92)), ((_609_v51) => function (_610_i12) {
              return _dafny.Seq.of((_609_v51).f12);
            })(_607_v51)));
          })(_485_v51);
          let _nw115 = Array((new BigNumber(25)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw115.length); _i0_21++) {
            _nw115[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _606_v150 = _nw115;
          let _index104 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_606_v150).length));
          (_606_v150)[_index104] = _module.__default.fm41(_485_v51.f13, globalState);
          let _611_v151;
          _611_v151 = _dafny.Seq.of((_485_v51).f12, true, _485_v51.f13);
          let _612_v152;
          _612_v152 = _dafny.Seq.of(_611_v151, _611_v151, _dafny.Seq.update(_611_v151, _module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_611_v151).length)), true), _dafny.Seq.Concat(_611_v151, _dafny.Seq.update(_dafny.Seq.update(_611_v151, _module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_611_v151).length)), (_485_v51).f12), _module.__default.safeIndex(_488_v54, new BigNumber((_dafny.Seq.update(_611_v151, _module.__default.safeIndex((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], new BigNumber((_611_v151).length)), (_485_v51).f12)).length)), _485_v51.f13)), _dafny.Seq.Concat(_611_v151, _611_v151));
          let _index105 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_606_v150).length));
          (_606_v150)[_index105] = _612_v152;
          let _613_v153;
          let _init22 = function (_614_i13) {
            return true;
          };
          let _nw116 = Array((new BigNumber(11)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw116.length); _i0_22++) {
            _nw116[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _613_v153 = _nw116;
          let _index106 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_613_v153).length));
          (_613_v153)[_index106] = (_485_v51).f12;
          let _index107 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_613_v153).length));
          (_613_v153)[_index107] = _this.f7;
          let _index108 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          (_486_v52)[_index108] = (_493_v57)[_module.__default.safeIndex(_module.__default.safeModuloInt((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _488_v54), new BigNumber((_493_v57).length))];
          _488_v54 = (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))];
        }
        let _615_v154;
        _615_v154 = _dafny.Map.Empty.slice().updateUnsafe(_486_v52,_488_v54);
        let _616_v155;
        _616_v155 = _dafny.Seq.of(_this.f7);
        _615_v154 = ((_this.f7) ? (_615_v154) : (_dafny.Map.Empty.slice().updateUnsafe(_486_v52,new BigNumber((_dafny.Seq.update(_616_v155, _module.__default.safeIndex(_488_v54, new BigNumber((_616_v155).length)), _485_v51.f13)).length))));
        if (_this.f7) {
          _564_v116 = _564_v116;
          let _617_v156;
          _617_v156 = _module.D1.create_DC4(_this.f7, _this.f7);
          let _pat_let_tv13 = _485_v51;
          let _618_v157;
          _618_v157 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let25_0) {
            return function (_619_dt__update__tmp_h3) {
              return function (_pat_let26_0) {
                return function (_620_dt__update_hcf7_h0) {
                  return _module.D1.create_DC4(_620_dt__update_hcf7_h0, (_619_dt__update__tmp_h3).dtor_cf8);
                }(_pat_let26_0);
              }(_pat_let_tv13.f13);
            }(_pat_let25_0);
          }(_617_v156),_this.f7);
          let _621_v159;
          _621_v159 = _dafny.Set.fromElements(_617_v156);
          let _rhs94 = function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_621_v159).Elements) {
              let _622_v158 = _compr_22;
              if ((_621_v159).contains(_622_v158)) {
                _coll22.push([_622_v158,_this.f7]);
              }
            }
            return _coll22;
          }();
          let _rhs95 = _this.f7;
          let _lhs71 = _485_v51;
          _618_v157 = _rhs94;
          _lhs71.f13 = _rhs95;
          _566_v118 = _566_v118;
          let _623_v160;
          _623_v160 = _dafny.Set.fromElements((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _624_v161;
          _624_v161 = _dafny.MultiSet.fromElements(_486_v52, _486_v52, _486_v52, _486_v52, _486_v52);
          let _rhs96 = _488_v54;
          let _rhs97 = (_623_v160).Difference(_623_v160);
          let _rhs98 = _624_v161;
          _488_v54 = _rhs96;
          _623_v160 = _rhs97;
          _624_v161 = _rhs98;
          let _625_v162;
          _625_v162 = _module.D1.create_DC5(_485_v51.f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), ((_626_v54) => function (_627_i14) {
  return _626_v54;
})(_488_v54)), new BigNumber(394));
          let _index109 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _rhs99 = ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).isLessThanOrEqualTo((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _rhs100 = _485_v51.f13;
          let _rhs101 = _module.__default.fm0(globalState);
          let _rhs102 = (_625_v162).dtor_cf9;
          let _lhs72 = _this;
          let _lhs73 = _485_v51;
          let _lhs74 = _486_v52;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _lhs76 = _485_v51;
          _lhs72.f7 = _rhs99;
          _lhs73.f13 = _rhs100;
          _lhs74[_lhs75] = _rhs101;
          _lhs76.f13 = _rhs102;
        } else {
          (_485_v51).f13 = _485_v51.f13;
          let _628_v163;
          _628_v163 = _dafny.Map.Empty.slice().updateUnsafe(_485_v51,_486_v52);
          let _629_v164;
          _629_v164 = _dafny.MultiSet.fromElements(_486_v52, _486_v52, (((_628_v163).contains(_485_v51)) ? ((_628_v163).get(_485_v51)) : (_486_v52)));
          _629_v164 = _629_v164;
          let _630_v165;
          _630_v165 = _dafny.Map.Empty.slice().updateUnsafe((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))],(_485_v51).fm17(_488_v54, _485_v51.f13, globalState));
          _630_v165 = (_630_v165).update(_488_v54, _616_v155);
          let _631_v166;
          _631_v166 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_485_v51.f13,(_485_v51).f12)).length));
          _631_v166 = (_631_v166).update((_485_v51).f12, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _632_v167;
          let _nw117 = Array((new BigNumber(18)).toNumber());
          _nw117[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw117[(_dafny.ONE).toNumber()] = (_485_v51).f12;
          _nw117[(new BigNumber(2)).toNumber()] = (_this).fm32(_485_v51.f13, _dafny.Set.fromElements((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]), (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], globalState);
          _nw117[(new BigNumber(3)).toNumber()] = (_485_v51).f12;
          _nw117[(new BigNumber(4)).toNumber()] = (_485_v51).f12;
          _nw117[(new BigNumber(5)).toNumber()] = (_485_v51).fm16(true, globalState);
          _nw117[(new BigNumber(6)).toNumber()] = _this.f7;
          _nw117[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_565_v117, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _488_v54, globalState);
          _nw117[(new BigNumber(8)).toNumber()] = _this.f7;
          _nw117[(new BigNumber(9)).toNumber()] = !((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).isEqualTo((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          _nw117[(new BigNumber(10)).toNumber()] = (_485_v51).f12;
          _nw117[(new BigNumber(11)).toNumber()] = (_485_v51).f12;
          _nw117[(new BigNumber(12)).toNumber()] = !(!(_485_v51.f13)) || (_this.f7);
          _nw117[(new BigNumber(13)).toNumber()] = (_485_v51.f13) === ((_485_v51).f12);
          _nw117[(new BigNumber(14)).toNumber()] = _485_v51.f13;
          _nw117[(new BigNumber(15)).toNumber()] = _485_v51.f13;
          _nw117[(new BigNumber(16)).toNumber()] = !(((_this.f7) ? ((_485_v51).f12) : ((_485_v51).f12)));
          _nw117[(new BigNumber(17)).toNumber()] = !(_this.f7);
          _632_v167 = _nw117;
          let _rhs103 = _dafny.Seq.IsPrefixOf(_564_v116, _dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), ((_633_v118) => function (_634_i15) {
            return _633_v118;
          })(_566_v118)));
          let _rhs104 = _632_v167;
          let _lhs77 = _this;
          _lhs77.f7 = _rhs103;
          _632_v167 = _rhs104;
        }
        let _635_v168;
        _635_v168 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(((_485_v51).f12) ? (_485_v51.f13) : (_485_v51.f13)));
        (_this).f7 = (((_635_v168).contains(_this.f7)) ? ((_635_v168).get(_this.f7)) : (_485_v51.f13));
      } else {
        if (false) {
          _488_v54 = _488_v54;
          let _636_v169;
          let _nw118 = new _module.C0();
          _nw118.__ctor(_this.f7, _485_v51.f13);
          _636_v169 = _nw118;
          (_485_v51).f13 = _485_v51.f13;
          let _637_v170;
          _637_v170 = _dafny.Map.Empty.slice().updateUnsafe(_485_v51.f13,new _dafny.CodePoint('e'.codePointAt(0)));
          let _638_v171;
          _638_v171 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? ((_485_v51).f12) : (true)),new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_566_v118)).Merge(_637_v170)).length));
          _638_v171 = (_638_v171).update((_485_v51).f12, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _639_v172;
          _639_v172 = _dafny.MultiSet.fromElements(_636_v169.f13, false);
          _module.__default.m0((((_639_v172).contains((_485_v51).f12)) ? ((_639_v172).get((_485_v51).f12)) : ((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])), _488_v54, globalState);
        } else {
          _488_v54 = _module.__default.safeDivisionInt(_488_v54, _488_v54);
          let _index110 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _rhs105 = ((_488_v54).plus(_488_v54)).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _rhs106 = _564_v116;
          let _lhs78 = _486_v52;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          _lhs78[_lhs79] = _rhs105;
          _564_v116 = _rhs106;
          let _640_v173;
          let _nw119 = Array((_dafny.ONE).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = (_485_v51).f12;
          _640_v173 = _nw119;
          let _641_v174;
          _641_v174 = _dafny.Set.fromElements(_640_v173, _640_v173);
          let _642_v175;
          _642_v175 = _module.D8.create_DC21(new BigNumber((_564_v116).length), _641_v174, (_dafny.ZERO).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]));
          _module.__default.m0(((_this.f7) ? (_488_v54) : ((_642_v175).dtor_cf38)), (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], globalState);
          _488_v54 = (_488_v54).multipliedBy(_488_v54);
          let _643_v176;
          _643_v176 = _dafny.MultiSet.fromElements(_488_v54);
          (_485_v51).f13 = !(((_643_v176).Difference(_dafny.MultiSet.fromElements((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))], _488_v54))).IsProperSubsetOf(_643_v176));
        }
        let _644_v177;
        _644_v177 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,(_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(163), _488_v54)));
        _644_v177 = (_644_v177).update(_485_v51.f13, (((_644_v177).contains((_485_v51).fm16((_485_v51).f12, globalState))) ? ((_644_v177).get((_485_v51).fm16((_485_v51).f12, globalState))) : (_488_v54)));
        let _645_v178;
        _645_v178 = _dafny.Set.fromElements(_486_v52);
        let _646_v179;
        _646_v179 = _module.D0.create_DC2(_485_v51.f13, new BigNumber(856), _645_v178, _488_v54);
        let _647_v180;
        _647_v180 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).fm16(_485_v51.f13, globalState),_646_v179);
        _module.__default.m0((_dafny.ZERO).minus(new BigNumber((_647_v180).length)), _488_v54, globalState);
        let _rhs107 = (_485_v51).f12;
        let _rhs108 = _dafny.Seq.Concat(_493_v57, _493_v57);
        let _lhs80 = _485_v51;
        _lhs80.f13 = _rhs107;
        _493_v57 = _rhs108;
        if ((_485_v51).f12) {
          let _648_v181;
          _648_v181 = _dafny.Set.fromElements(_485_v51.f13);
          let _649_v182;
          _649_v182 = _dafny.Set.fromElements(new BigNumber((_648_v181).length));
          let _650_v183;
          let _nw120 = Array((new BigNumber(11)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _485_v51.f13;
          _nw120[(_dafny.ONE).toNumber()] = _485_v51.f13;
          _nw120[(new BigNumber(2)).toNumber()] = _485_v51.f13;
          _nw120[(new BigNumber(3)).toNumber()] = (_485_v51).f12;
          _nw120[(new BigNumber(4)).toNumber()] = true;
          _nw120[(new BigNumber(5)).toNumber()] = _485_v51.f13;
          _nw120[(new BigNumber(6)).toNumber()] = (_this).fm32(_485_v51.f13, _649_v182, new BigNumber(893), (_dafny.ZERO).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]), globalState);
          _nw120[(new BigNumber(7)).toNumber()] = _485_v51.f13;
          _nw120[(new BigNumber(8)).toNumber()] = _485_v51.f13;
          _nw120[(new BigNumber(9)).toNumber()] = false;
          _nw120[(new BigNumber(10)).toNumber()] = (_485_v51).f12;
          _650_v183 = _nw120;
          let _651_v184;
          _651_v184 = _dafny.Set.fromElements(_650_v183);
          (_485_v51).f13 = (_651_v184).contains(_650_v183);
          let _652_v185;
          _652_v185 = _dafny.Map.Empty.slice().updateUnsafe(_650_v183,(_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          _652_v185 = (_652_v185).update(_650_v183, (_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _653_v186;
          _653_v186 = _module.D1.create_DC4(_485_v51.f13, (_485_v51).f12);
          let _index111 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_650_v183).length));
          (_650_v183)[_index111] = (_488_v54).isLessThanOrEqualTo(new BigNumber((_648_v181).length));
          let _index112 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _index113 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_650_v183).length));
          let _rhs109 = new BigNumber(301);
          let _rhs110 = _653_v186;
          let _rhs111 = ((_485_v51).f12) && (_this.f7);
          let _rhs112 = ((_this.f7) ? (true) : ((_485_v51).f12));
          let _lhs81 = _486_v52;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length));
          let _lhs83 = _650_v183;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_650_v183).length));
          let _lhs85 = _485_v51;
          _lhs81[_lhs82] = _rhs109;
          _653_v186 = _rhs110;
          _lhs83[_lhs84] = _rhs111;
          _lhs85.f13 = _rhs112;
          let _654_v187;
          _654_v187 = _dafny.Map.Empty.slice().updateUnsafe((_485_v51).f12,_493_v57);
          let _655_v188;
          _655_v188 = _dafny.Seq.of(true);
          _654_v187 = (_654_v187).update(!((_655_v188)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(254)), ((_656_v52) => function (_657_i16) {
            return (_656_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_656_v52).length))];
          })(_486_v52))).length), new BigNumber((_655_v188).length))]) || (_485_v51.f13), _493_v57);
          (_this).f7 = (((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]).minus((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))])).isEqualTo((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
        } else {
          let _658_v189;
          _658_v189 = _dafny.Set.fromElements(_488_v54);
          let _659_v190;
          _659_v190 = _dafny.Set.fromElements(_485_v51.f13);
          (_this).f7 = !(!(_659_v190).contains(!(!((_658_v189).IsSubsetOf(_658_v189)))));
          let _660_v191;
          _660_v191 = _module.D0.create_DC1((_486_v52)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_486_v52).length))]);
          let _661_v192;
          let _init23 = ((_662_v51) => function (_663_i17) {
            return (_662_v51).f12;
          })(_485_v51);
          let _nw121 = Array((new BigNumber(22)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw121.length); _i0_23++) {
            _nw121[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _661_v192 = _nw121;
          let _664_v193;
          let _nw122 = Array((new BigNumber(3)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _661_v192;
          _nw122[(_dafny.ONE).toNumber()] = _661_v192;
          _nw122[(new BigNumber(2)).toNumber()] = _661_v192;
          _664_v193 = _nw122;
          let _rhs113 = _660_v191;
          let _rhs114 = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), _488_v54);
          let _rhs115 = _485_v51.f13;
          let _rhs116 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(147)), ((_665_v118) => function (_666_i18) {
            return _665_v118;
          })(_566_v118));
          let _rhs117 = _664_v193;
          let _lhs86 = _485_v51;
          _660_v191 = _rhs113;
          _488_v54 = _rhs114;
          _lhs86.f13 = _rhs115;
          _564_v116 = _rhs116;
          _664_v193 = _rhs117;
          (_485_v51).f13 = _485_v51.f13;
          let _667_v194;
          _667_v194 = _dafny.Seq.of(_487_v53, _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_488_v54),(_485_v51).f12), (_487_v53).Merge(_487_v53), _487_v53, _487_v53);
          _667_v194 = _667_v194;
          let _668_v195;
          let _nw123 = new _module.C0();
          _nw123.__ctor(_this.f7, _this.f7);
          _668_v195 = _nw123;
        }
      }
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f7 = false;
      this._f15 = _dafny.Set.Empty;
      this._f16 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f15, f16, f7) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f7 = f7;
      return;
    }
    fm24(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      if (!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), function (_669_i1) {
        return _dafny.Seq.UnicodeFromString("nohtpuim");
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(85)), function (_670_i0) {
        return (_this).f16;
      }))) {
        let _671_v0;
        _671_v0 = _dafny.Seq.of(p1);
        let _672_v1;
        _672_v1 = _dafny.Seq.of(p0);
        let _673_v2;
        _673_v2 = _module.D1.create_DC5((p2) === (p0), _dafny.Seq.Concat(_dafny.Seq.update(_671_v0, _module.__default.safeIndex(p1, new BigNumber((_671_v0).length)), p1), _671_v0), new BigNumber((_672_v1).length));
        let _source2 = _673_v2;
        if (_source2.is_DC4) {
          let _674___mcc_h0 = (_source2).cf7;
          let _675___mcc_h1 = (_source2).cf8;
          let _676_cf8 = _675___mcc_h1;
          let _677_cf7 = _674___mcc_h0;
          let _678_v3;
          _678_v3 = new BigNumber(528);
          _678_v3 = _module.__default.fm0(globalState);
          _678_v3 = _678_v3;
          let _rhs118 = _678_v3;
          let _rhs119 = _module.__default.safeDivisionInt(p1, _678_v3);
          let _rhs120 = (p1).isLessThan(_678_v3);
          _678_v3 = _rhs118;
          _678_v3 = _rhs119;
          _677_cf7 = _rhs120;
          _678_v3 = (p1).minus(_678_v3);
        } else if (_source2.is_DC5) {
          let _679___mcc_h2 = (_source2).cf9;
          let _680___mcc_h3 = (_source2).cf10;
          let _681___mcc_h4 = (_source2).cf11;
          let _682_cf11 = _681___mcc_h4;
          let _683_cf10 = _680___mcc_h3;
          let _684_cf9 = _679___mcc_h2;
          let _685_v4;
          let _init24 = ((_686_p1) => function (_687_i2) {
            return _module.__default.safeDivisionInt(_687_i2, _686_p1);
          })(p1);
          let _nw124 = Array((new BigNumber(11)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw124.length); _i0_24++) {
            _nw124[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _685_v4 = _nw124;
          _685_v4 = _685_v4;
          let _688_v5;
          let _init25 = function (_689_i3) {
            return _this.f7;
          };
          let _nw125 = Array((new BigNumber(2)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw125.length); _i0_25++) {
            _nw125[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _688_v5 = _nw125;
          let _690_v6;
          _690_v6 = _dafny.MultiSet.fromElements(_688_v5, _688_v5, _688_v5, _688_v5);
          _690_v6 = _690_v6;
          let _691_v7;
          _691_v7 = _dafny.Set.fromElements(p0);
          let _692_v8;
          _692_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_691_v7).length));
          let _693_v9;
          _693_v9 = _dafny.MultiSet.fromElements(new BigNumber((_692_v8).length));
          let _694_v10;
          _694_v10 = _dafny.Seq.of(_693_v9, _693_v9);
          let _695_v11;
          _695_v11 = _module.D1.create_DC3(_691_v7);
          let _696_v12;
          _696_v12 = _module.D0.create_DC0(_682_cf11);
          let _697_v13;
          _697_v13 = _dafny.Seq.of(_696_v12, _696_v12);
          let _698_v14;
          let _nw126 = Array((new BigNumber(13)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw126[(_dafny.ONE).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw126[(new BigNumber(3)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(4)).toNumber()] = _module.__default.fm26(_684_cf9, p1, _682_cf11, globalState);
          _nw126[(new BigNumber(5)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(6)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(7)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(9)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(10)).toNumber()] = (_this).f16;
          _nw126[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw126[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _698_v14 = _nw126;
          (_this).m13((_dafny.MultiSet.FromArray(_694_v10)).IsDisjointFrom(_dafny.MultiSet.fromElements(_693_v9, _module.__default.fm25(_682_cf11, (_this).f16, new BigNumber((_dafny.Seq.of(p2, p0)).length), _695_v11, globalState))), p0, _module.__default.fm1(_697_v13, _682_cf11, _682_cf11, globalState), _698_v14, globalState);
          let _699_v15;
          _699_v15 = _dafny.Seq.UnicodeFromString("wawknevf");
          _699_v15 = _699_v15;
        } else if (_source2.is_DC3) {
          let _700___mcc_h5 = (_source2).cf6;
          let _701_cf6 = _700___mcc_h5;
          let _702_v16;
          _702_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.safeModuloInt(p1, p1));
          _702_v16 = (_702_v16).update(p1, _module.__default.fm0(globalState));
          let _703_v17;
          let _nw127 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _703_v17 = _nw127;
          let _index114 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_703_v17).length));
          (_703_v17)[_index114] = (_dafny.ZERO).minus(p1);
          let _index115 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_703_v17).length));
          (_703_v17)[_index115] = _module.__default.fm0(globalState);
          (_this).f7 = p0;
          let _index116 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_703_v17).length));
          (_703_v17)[_index116] = (p1).plus((_dafny.ZERO).minus((_703_v17)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_703_v17).length))]));
        } else {
          let _704___mcc_h6 = (_source2).cf12;
          let _705_cf12 = _704___mcc_h6;
          let _706_v18;
          _706_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p2));
          r0 = ((_this).fm24(false, p0, p1, globalState)).Merge(_706_v18);
          let _707_v19;
          _707_v19 = new BigNumber(410);
          let _708_v20;
          let _nw128 = Array((new BigNumber(28)).toNumber()).fill(false);
          _708_v20 = _nw128;
          let _709_v21;
          _709_v21 = _dafny.Seq.of(_module.D0.create_DC0(new BigNumber(-726)));
          let _index117 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_708_v20).length));
          (_708_v20)[_index117] = !(_module.__default.fm1(_709_v21, _707_v19, new BigNumber(891), globalState));
          let _710_v22;
          let _nw129 = new _module.C0();
          _nw129.__ctor(_this.f7, !(p0));
          _710_v22 = _nw129;
          let _index118 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_708_v20).length));
          let _rhs121 = (_707_v19).multipliedBy(_707_v19);
          let _rhs122 = p2;
          let _rhs123 = _710_v22;
          let _rhs124 = p1;
          let _lhs87 = _708_v20;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_708_v20).length));
          _707_v19 = _rhs121;
          _lhs87[_lhs88] = _rhs122;
          _710_v22 = _rhs123;
          _707_v19 = _rhs124;
          let _711_v23;
          _711_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,_671_v0);
          _711_v23 = (_711_v23).update(p1, _671_v0);
          let _712_v24;
          let _init26 = function (_713_i4) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_714_i5) {
              return (_this).f16;
            }), _dafny.Seq.UnicodeFromString("iwiju"));
          };
          let _nw130 = Array((new BigNumber(26)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw130.length); _i0_26++) {
            _nw130[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _712_v24 = _nw130;
          let _715_v25;
          _715_v25 = _dafny.Seq.UnicodeFromString("rnjpk");
          let _index119 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_712_v24).length));
          (_712_v24)[_index119] = _715_v25;
          let _716_v26;
          let _nw131 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _716_v26 = _nw131;
          let _717_v27;
          _717_v27 = _dafny.Map.Empty.slice().updateUnsafe((_710_v22).f12,_707_v19);
          let _index120 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_716_v26).length));
          (_716_v26)[_index120] = _717_v27;
          let _718_v28;
          _718_v28 = new _dafny.CodePoint('d'.codePointAt(0));
          let _719_v29;
          _719_v29 = _dafny.MultiSet.fromElements(_707_v19);
          let _720_v30;
          _720_v30 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(80), p1), _719_v29, _719_v29, _719_v29);
          let _721_v31;
          _721_v31 = _dafny.Set.fromElements(_this.f7, (((_706_v18).contains(p2)) ? ((_706_v18).get(p2)) : ((_708_v20)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_708_v20).length))])), _dafny.areEqual(_715_v25, _715_v25), (_720_v30).IsDisjointFrom(_720_v30), (_708_v20)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_708_v20).length))]);
          let _index121 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_712_v24).length));
          let _index122 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_716_v26).length));
          let _rhs125 = _715_v25;
          let _rhs126 = (_717_v27).Merge(_717_v27);
          let _rhs127 = _721_v31;
          let _rhs128 = _module.__default.fm26(!(_this.f7), _module.__default.safeDivisionInt(p1, (_671_v0)[_module.__default.safeIndex(p1, new BigNumber((_671_v0).length))]), new BigNumber(236), globalState);
          let _lhs89 = _712_v24;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_712_v24).length));
          let _lhs91 = _716_v26;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_716_v26).length));
          _lhs89[_lhs90] = _rhs125;
          _lhs91[_lhs92] = _rhs126;
          r1 = _rhs127;
          _718_v28 = _rhs128;
        }
        let _722_v32;
        _722_v32 = _dafny.Set.fromElements(p2);
        let _723_v33;
        let _init27 = ((_724_p1) => function (_725_i6) {
          return _module.__default.safeDivisionInt(_725_i6, _724_p1);
        })(p1);
        let _nw132 = Array((new BigNumber(12)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw132.length); _i0_27++) {
          _nw132[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _723_v33 = _nw132;
        let _726_v34;
        _726_v34 = _dafny.Map.Empty.slice().updateUnsafe(_722_v32,_723_v33);
        let _727_v35;
        let _init28 = ((_728_p0, _729_p2) => function (_730_i7) {
          return (_module.D1.create_DC4(_728_p0, _729_p2)).dtor_cf8;
        })(p0, p2);
        let _nw133 = Array((new BigNumber(20)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw133.length); _i0_28++) {
          _nw133[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _727_v35 = _nw133;
        let _731_v36;
        _731_v36 = _dafny.Map.Empty.slice().updateUnsafe(_726_v34,_727_v35);
        _731_v36 = (_731_v36).update((_726_v34).Merge(_726_v34), _727_v35);
        let _732_v37;
        _732_v37 = new BigNumber(584);
        let _733_v38;
        _733_v38 = _dafny.Map.Empty.slice().updateUnsafe(!((p0) || (true)),_module.__default.safeDivisionInt(new BigNumber((_671_v0).length), new BigNumber(428)));
        _732_v37 = (((_733_v38).contains(p0)) ? ((_733_v38).get(p0)) : (p1));
        let _index123 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_727_v35).length));
        (_727_v35)[_index123] = p2;
        let _index124 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_727_v35).length));
        (_727_v35)[_index124] = p2;
        _module.__default.m0(_module.__default.safeModuloInt(p1, _732_v37), _732_v37, globalState);
      } else {
        let _734_v39;
        _734_v39 = new BigNumber(-291);
        let _735_v40;
        _735_v40 = _dafny.Seq.UnicodeFromString("uk");
        _734_v39 = (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_734_v39)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),new BigNumber((_735_v40).length)))).length));
        let _736_v41;
        _736_v41 = _module.D0.create_DC0(p1);
        let _737_v42;
        _737_v42 = _dafny.Seq.of(function (_pat_let27_0) {
          return function (_738_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_739_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_739_dt__update_hcf0_h0);
              }(_pat_let28_0);
            }(new BigNumber(778));
          }(_pat_let27_0);
        }(_module.D0.create_DC0(p1)));
        let _740_v43;
        _740_v43 = _module.D6.create_DC17(_module.__default.fm1(_737_v42, new BigNumber(472), p1, globalState), _734_v39, (_this).f16, new _dafny.CodePoint('u'.codePointAt(0)));
        let _pat_let_tv14 = _734_v39;
        let _741_v44;
        let _nw134 = Array((new BigNumber(19)).toNumber());
        _nw134[(_dafny.ZERO).toNumber()] = _this.f7;
        _nw134[(_dafny.ONE).toNumber()] = p2;
        _nw134[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_dafny.Seq.of(_736_v41), (_740_v43).dtor_cf32, p1, globalState);
        _nw134[(new BigNumber(3)).toNumber()] = !(p2);
        _nw134[(new BigNumber(4)).toNumber()] = p0;
        _nw134[(new BigNumber(5)).toNumber()] = _module.__default.fm1(_dafny.Seq.of(_736_v41, function (_pat_let29_0) {
          return function (_742_dt__update__tmp_h1) {
            return function (_pat_let30_0) {
              return function (_743_dt__update_hcf0_h1) {
                return _module.D0.create_DC0(_743_dt__update_hcf0_h1);
              }(_pat_let30_0);
            }(_pat_let_tv14);
          }(_pat_let29_0);
        }(_736_v41), _736_v41, _736_v41), _734_v39, _734_v39, globalState);
        _nw134[(new BigNumber(6)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(7)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(8)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(9)).toNumber()] = p2;
        _nw134[(new BigNumber(10)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(11)).toNumber()] = p0;
        _nw134[(new BigNumber(12)).toNumber()] = p2;
        _nw134[(new BigNumber(13)).toNumber()] = p0;
        _nw134[(new BigNumber(14)).toNumber()] = p0;
        _nw134[(new BigNumber(15)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(16)).toNumber()] = _this.f7;
        _nw134[(new BigNumber(17)).toNumber()] = p0;
        _nw134[(new BigNumber(18)).toNumber()] = true;
        _741_v44 = _nw134;
        let _744_v45;
        _744_v45 = _module.D8.create_DC21(_734_v39, (_this).f15, _734_v39);
        if (((_744_v45).dtor_cf39).IsProperSubsetOf(_dafny.Set.fromElements(_741_v44))) {
          let _745_v46;
          _745_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),new BigNumber((_dafny.Seq.UnicodeFromString("wl")).length));
          let _746_v47;
          _746_v47 = _dafny.MultiSet.fromElements(p1, p1);
          _734_v39 = (new BigNumber((_dafny.Seq.UnicodeFromString("xngciqxxs")).length)).minus((((_745_v46).contains(_734_v39)) ? ((_745_v46).get(_734_v39)) : ((_dafny.ZERO).minus((((_746_v47).contains(p1)) ? ((_746_v47).get(p1)) : (p1))))));
          let _747_v48;
          _747_v48 = _module.D0.create_DC1((p1).minus(_734_v39));
          _747_v48 = _747_v48;
          let _748_v49;
          let _nw135 = Array((new BigNumber(14)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = new BigNumber(-717);
          _nw135[(_dafny.ONE).toNumber()] = _734_v39;
          _nw135[(new BigNumber(2)).toNumber()] = _734_v39;
          _nw135[(new BigNumber(3)).toNumber()] = _734_v39;
          _nw135[(new BigNumber(4)).toNumber()] = _module.__default.fm0(globalState);
          _nw135[(new BigNumber(5)).toNumber()] = p1;
          _nw135[(new BigNumber(6)).toNumber()] = _734_v39;
          _nw135[(new BigNumber(7)).toNumber()] = p1;
          _nw135[(new BigNumber(8)).toNumber()] = p1;
          _nw135[(new BigNumber(9)).toNumber()] = _734_v39;
          _nw135[(new BigNumber(10)).toNumber()] = _734_v39;
          _nw135[(new BigNumber(11)).toNumber()] = p1;
          _nw135[(new BigNumber(12)).toNumber()] = p1;
          _nw135[(new BigNumber(13)).toNumber()] = p1;
          _748_v49 = _nw135;
          let _749_v50;
          _749_v50 = _dafny.Map.Empty.slice().updateUnsafe(_748_v49,_this.f7);
          _749_v50 = (_749_v50).update(_748_v49, p0);
          let _750_v51;
          _750_v51 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f7) ? (p0) : (_this.f7)),p0);
          r0 = _750_v51;
          (_this).f7 = ((_module.__default.fm0(globalState)).multipliedBy(p1)).isEqualTo((_dafny.ZERO).minus(p1));
        } else {
          let _751_v52;
          let _nw136 = Array((new BigNumber(20)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = (_this).f16;
          _nw136[(_dafny.ONE).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(2)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(3)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(4)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(5)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(6)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(7)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw136[(new BigNumber(10)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _nw136[(new BigNumber(12)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(13)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(14)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(15)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(16)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(17)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(18)).toNumber()] = (_this).f16;
          _nw136[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _751_v52 = _nw136;
          (_this).m13(p2, p0, !(p0) || (_this.f7), _751_v52, globalState);
          _734_v39 = (_734_v39).plus(_734_v39);
          let _752_v53;
          _752_v53 = _dafny.MultiSet.fromElements(!(p0));
          let _753_v54;
          _753_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p1);
          let _754_v55;
          _754_v55 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_this.f7);
          let _755_v56;
          _755_v56 = _dafny.Seq.of((((_754_v55).contains(new _dafny.CodePoint('y'.codePointAt(0)))) ? ((_754_v55).get(new _dafny.CodePoint('y'.codePointAt(0)))) : (p0)));
          let _756_v57;
          _756_v57 = _dafny.Seq.of(_734_v39);
          let _757_v58;
          _757_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _758_v59;
          let _nw137 = Array((new BigNumber(27)).toNumber());
          _nw137[(_dafny.ZERO).toNumber()] = _734_v39;
          _nw137[(_dafny.ONE).toNumber()] = (((_752_v53).contains(p0)) ? ((_752_v53).get(p0)) : (p1));
          _nw137[(new BigNumber(2)).toNumber()] = p1;
          _nw137[(new BigNumber(3)).toNumber()] = p1;
          _nw137[(new BigNumber(4)).toNumber()] = (((_753_v54).contains(p0)) ? ((_753_v54).get(p0)) : (_734_v39));
          _nw137[(new BigNumber(5)).toNumber()] = new BigNumber((_755_v56).length);
          _nw137[(new BigNumber(6)).toNumber()] = p1;
          _nw137[(new BigNumber(7)).toNumber()] = p1;
          _nw137[(new BigNumber(8)).toNumber()] = _734_v39;
          _nw137[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_759_i8) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(104), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_760_i8) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length)), (_this).f16)).length);
          _nw137[(new BigNumber(10)).toNumber()] = p1;
          _nw137[(new BigNumber(11)).toNumber()] = (_756_v57)[_module.__default.safeIndex(p1, new BigNumber((_756_v57).length))];
          _nw137[(new BigNumber(12)).toNumber()] = p1;
          _nw137[(new BigNumber(13)).toNumber()] = _734_v39;
          _nw137[(new BigNumber(14)).toNumber()] = _734_v39;
          _nw137[(new BigNumber(15)).toNumber()] = p1;
          _nw137[(new BigNumber(16)).toNumber()] = _734_v39;
          _nw137[(new BigNumber(17)).toNumber()] = _module.__default.fm0(globalState);
          _nw137[(new BigNumber(18)).toNumber()] = new BigNumber(-396);
          _nw137[(new BigNumber(19)).toNumber()] = (_756_v57)[_module.__default.safeIndex(_734_v39, new BigNumber((_756_v57).length))];
          _nw137[(new BigNumber(20)).toNumber()] = new BigNumber(-866);
          _nw137[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), function (_761_i9) {
            return (_this).f16;
          })).length);
          _nw137[(new BigNumber(22)).toNumber()] = _734_v39;
          _nw137[(new BigNumber(23)).toNumber()] = p1;
          _nw137[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw137[(new BigNumber(25)).toNumber()] = new BigNumber((_757_v58).length);
          _nw137[(new BigNumber(26)).toNumber()] = p1;
          _758_v59 = _nw137;
          let _762_v60;
          _762_v60 = _dafny.Seq.of(_758_v59);
          let _763_v61;
          _763_v61 = _dafny.Seq.of(_762_v60, _762_v60, _762_v60, _dafny.Seq.Concat(_762_v60, _762_v60), _762_v60);
          _734_v39 = new BigNumber((_dafny.Seq.update((_763_v61)[_module.__default.safeIndex((_dafny.ZERO).minus(_734_v39), new BigNumber((_763_v61).length))], _module.__default.safeIndex(new BigNumber((_module.__default.fm27(_this.f7, new _dafny.CodePoint('c'.codePointAt(0)), _this.f7, (_740_v43).dtor_cf31, globalState)).length), new BigNumber(((_763_v61)[_module.__default.safeIndex((_dafny.ZERO).minus(_734_v39), new BigNumber((_763_v61).length))]).length)), _758_v59)).length);
          (_this).f7 = (p1).isLessThan(p1);
          let _index125 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_758_v59).length));
          (_758_v59)[_index125] = new BigNumber((_735_v40).length);
          let _index126 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_758_v59).length));
          (_758_v59)[_index126] = p1;
        }
        (_this).f7 = p2;
        (_this).f7 = _dafny.areEqual(_dafny.Seq.Concat(_735_v40, _735_v40), _dafny.Seq.update(_735_v40, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_735_v40).length)), (_this).f16));
        let _764_v62;
        let _init29 = ((_765_v40) => function (_766_i10) {
          return _765_v40;
        })(_735_v40);
        let _nw138 = Array((new BigNumber(3)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw138.length); _i0_29++) {
          _nw138[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _764_v62 = _nw138;
        let _index127 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_764_v62).length));
        (_764_v62)[_index127] = _735_v40;
        let _767_v63;
        _767_v63 = _dafny.Seq.of(new BigNumber(938));
        let _768_v64;
        _768_v64 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_769_i11) {
          return (_this).f16;
        }), _module.__default.safeIndex(new BigNumber(72), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_770_i11) {
          return (_this).f16;
        })).length)), (_this).f16)).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_767_v63).length))), _734_v39);
        let _771_v65;
        _771_v65 = _dafny.Map.Empty.slice().updateUnsafe(_734_v39,_dafny.MultiSet.fromElements(_module.__default.fm0(globalState), p1));
        let _index128 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_764_v62).length));
        let _rhs129 = !((_768_v64).IsDisjointFrom((((_771_v65).contains(p1)) ? ((_771_v65).get(p1)) : (_768_v64))));
        let _rhs130 = p1;
        let _rhs131 = _this.f7;
        let _rhs132 = _735_v40;
        let _lhs93 = _this;
        let _lhs94 = _this;
        let _lhs95 = _764_v62;
        let _lhs96 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_764_v62).length));
        _lhs93.f7 = _rhs129;
        _734_v39 = _rhs130;
        _lhs94.f7 = _rhs131;
        _lhs95[_lhs96] = _rhs132;
      }
      if ((_module.D3.create_DC11(p1, p1, p1, _this.f7)).dtor_cf25) {
        let _772_v66;
        _772_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p1);
        let _773_v67;
        _773_v67 = _dafny.Seq.UnicodeFromString("n");
        let _774_v68;
        let _nw139 = new _module.C0();
        _nw139.__ctor(_this.f7, (new BigNumber((_772_v66).length)).isEqualTo(new BigNumber((_773_v67).length)));
        _774_v68 = _nw139;
        (_this).f7 = (new BigNumber((_dafny.Seq.Concat(_773_v67, _773_v67)).length)).isLessThanOrEqualTo(p1);
        let _775_v69;
        _775_v69 = _module.D0.create_DC0(p1);
        let _776_v70;
        _776_v70 = _dafny.Seq.of(_775_v69);
        let _777_v71;
        _777_v71 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm1(_776_v70, p1, p1, globalState));
        (_this).f7 = (((_777_v71).contains(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1))) ? ((_777_v71).get(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1))) : (p2));
        let _778_v72;
        _778_v72 = new BigNumber(408);
        _778_v72 = _778_v72;
        let _779_v73;
        _779_v73 = _dafny.Map.Empty.slice().updateUnsafe(_774_v68.f13,(_this).f16);
        let _780_v74;
        _780_v74 = _dafny.Seq.of(new BigNumber(318), _778_v72);
        let _781_v75;
        let _nw140 = Array((new BigNumber(25)).toNumber());
        _nw140[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw140[(_dafny.ONE).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(2)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(3)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(4)).toNumber()] = (((_779_v73).contains(!(p2))) ? ((_779_v73).get(!(p2))) : ((_this).f16));
        _nw140[(new BigNumber(5)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(6)).toNumber()] = (_773_v67)[_module.__default.safeIndex(new BigNumber((_780_v74).length), new BigNumber((_773_v67).length))];
        _nw140[(new BigNumber(7)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(8)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(9)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(10)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(11)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(12)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(13)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(14)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(15)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
        _nw140[(new BigNumber(17)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(18)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw140[(new BigNumber(20)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(21)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(22)).toNumber()] = (_this).f16;
        _nw140[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw140[(new BigNumber(24)).toNumber()] = (_this).f16;
        _781_v75 = _nw140;
        (_this).m13(p2, _this.f7, (p0) === ((_774_v68).f12), _781_v75, globalState);
      } else {
        if (p2) {
          let _782_v76;
          let _nw141 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
          _782_v76 = _nw141;
          let _783_v77;
          _783_v77 = _dafny.Set.fromElements(p0);
          let _index129 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_782_v76).length));
          (_782_v76)[_index129] = _783_v77;
          let _index130 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_782_v76).length));
          (_782_v76)[_index130] = _module.__default.fm27(p0, (_this).f16, p2, p2, globalState);
          let _784_v78;
          let _nw142 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _784_v78 = _nw142;
          let _index131 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_784_v78).length));
          (_784_v78)[_index131] = p1;
          let _785_v79;
          _785_v79 = _dafny.Seq.of(false);
          let _index132 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_784_v78).length));
          (_784_v78)[_index132] = (_module.__default.safeModuloInt(new BigNumber((_785_v79).length), p1)).plus((p1).multipliedBy(p1));
          let _786_v80;
          _786_v80 = _dafny.Map.Empty.slice().updateUnsafe(_784_v78,_784_v78);
          _786_v80 = (_786_v80).update(_784_v78, _784_v78);
          (_this).f7 = p2;
          let _787_v81;
          _787_v81 = new _dafny.CodePoint('o'.codePointAt(0));
          let _788_v82;
          _788_v82 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
          let _789_v83;
          _789_v83 = _dafny.Map.Empty.slice().updateUnsafe(p1,_788_v82);
          _787_v81 = ((p2) ? (_module.__default.fm26(p0, new BigNumber((_789_v83).length), (_784_v78)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_784_v78).length))], globalState)) : ((_this).f16));
        } else {
          let _790_v84;
          _790_v84 = _dafny.Seq.UnicodeFromString("ymragmkba");
          _790_v84 = _790_v84;
          let _791_v85;
          _791_v85 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          _791_v85 = (_791_v85).update(!(p0) || (p2), p2);
          let _792_v86;
          _792_v86 = new BigNumber(718);
          let _793_v87;
          let _init30 = ((_794_p0) => function (_795_i12) {
            return _794_p0;
          })(p0);
          let _nw143 = Array((new BigNumber(22)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw143.length); _i0_30++) {
            _nw143[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _793_v87 = _nw143;
          let _796_v88;
          _796_v88 = _dafny.MultiSet.fromElements(_793_v87);
          let _797_v89;
          _797_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber(((_dafny.MultiSet.fromElements(_793_v87, _793_v87, _793_v87, _793_v87, _793_v87)).Difference(_796_v88)).cardinality()));
          let _798_v90;
          _798_v90 = _dafny.Seq.of(false, p0);
          let _799_v91;
          _799_v91 = _dafny.MultiSet.fromElements(p0, p0);
          let _rhs133 = new BigNumber((_797_v89).length);
          let _rhs134 = !(!_dafny.areEqual(_790_v84, _790_v84)) || ((_799_v91).IsSubsetOf(_dafny.MultiSet.FromArray(_798_v90)));
          let _lhs97 = _this;
          _792_v86 = _rhs133;
          _lhs97.f7 = _rhs134;
          let _800_v92;
          let _init31 = function (_801_i13) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          };
          let _nw144 = Array((new BigNumber(3)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw144.length); _i0_31++) {
            _nw144[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _800_v92 = _nw144;
          (_this).m13(p0, (p0) === (_this.f7), !(p2) || (p2), _800_v92, globalState);
          let _802_v93;
          _802_v93 = new _dafny.CodePoint('b'.codePointAt(0));
          _802_v93 = (_this).f16;
        }
        if (p0) {
          (_this).f7 = p0;
          let _803_v94;
          _803_v94 = _dafny.Seq.UnicodeFromString("cksvuxoeh");
          let _804_v95;
          _804_v95 = _dafny.Map.Empty.slice().updateUnsafe(_803_v94,_this.f7);
          let _805_v96;
          _805_v96 = _dafny.Seq.of(_this.f7);
          let _806_v97;
          _806_v97 = _dafny.MultiSet.fromElements(false, false, (_805_v96)[_module.__default.safeIndex(p1, new BigNumber((_805_v96).length))], _this.f7, _this.f7);
          let _807_v98;
          let _nw145 = Array((_dafny.ONE).toNumber()).fill(false);
          _807_v98 = _nw145;
          let _808_v99;
          _808_v99 = _dafny.Map.Empty.slice().updateUnsafe(p0,_807_v98);
          let _809_v100;
          _809_v100 = _dafny.Map.Empty.slice().updateUnsafe(_808_v99,_805_v96);
          let _810_v101;
          _810_v101 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D1.create_DC4(p2, true)).dtor_cf7);
          let _811_v102;
          let _nw146 = Array((new BigNumber(22)).toNumber());
          _nw146[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p2);
          _nw146[(_dafny.ONE).toNumber()] = _dafny.Seq.of((((_804_v95).contains(_module.__default.fm28(p2, new BigNumber((_dafny.Seq.of(p0)).length), globalState))) ? ((_804_v95).get(_module.__default.fm28(p2, new BigNumber((_dafny.Seq.of(p0)).length), globalState))) : (p2)), p0, p0, _this.f7, _this.f7);
          _nw146[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_805_v96, _805_v96);
          _nw146[(new BigNumber(3)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(4)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(5)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(6)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(7)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(8)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(9)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_805_v96, _module.__default.safeIndex(new BigNumber((_806_v97).cardinality()), new BigNumber((_805_v96).length)), _this.f7);
          _nw146[(new BigNumber(11)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(12)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_805_v96, _805_v96);
          _nw146[(new BigNumber(14)).toNumber()] = (((_809_v100).contains(_808_v99)) ? ((_809_v100).get(_808_v99)) : (_805_v96));
          _nw146[(new BigNumber(15)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((((_810_v101).contains(p0)) ? ((_810_v101).get(p0)) : (p0)), p0, p0), _805_v96);
          _nw146[(new BigNumber(17)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(p2, _this.f7);
          _nw146[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(p2, true);
          _nw146[(new BigNumber(20)).toNumber()] = _805_v96;
          _nw146[(new BigNumber(21)).toNumber()] = _805_v96;
          _811_v102 = _nw146;
          let _index133 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_811_v102).length));
          (_811_v102)[_index133] = _dafny.Seq.of(!(p0));
          let _index134 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_811_v102).length));
          (_811_v102)[_index134] = _805_v96;
          let _812_v103;
          _812_v103 = new BigNumber(-377);
          _812_v103 = _812_v103;
          let _813_v104;
          let _nw147 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _813_v104 = _nw147;
          let _814_v105;
          _814_v105 = _module.D9.create_DC23(_813_v104);
          _813_v104 = (_814_v105).dtor_cf44;
          let _815_v107;
          _815_v107 = _dafny.Set.fromElements(p1, _812_v103);
          (_this).f7 = !(((_this.f7) ? (p2) : ((function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(140), new BigNumber(998))) {
              let _816_v106 = _compr_23;
              if (((new BigNumber(140)).isLessThanOrEqualTo(_816_v106)) && ((_816_v106).isLessThan(new BigNumber(998)))) {
                _coll23.add((_816_v106).multipliedBy(p1));
              }
            }
            return _coll23;
          }()).IsDisjointFrom(_815_v107))));
        } else {
          (_this).f7 = _this.f7;
          let _817_v108;
          let _init32 = ((_818_p1) => function (_819_i14) {
            return _module.__default.safeModuloInt(_819_i14, _818_p1);
          })(p1);
          let _nw148 = Array((new BigNumber(7)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw148.length); _i0_32++) {
            _nw148[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _817_v108 = _nw148;
          let _820_v109;
          _820_v109 = _dafny.Map.Empty.slice().updateUnsafe(_817_v108,_dafny.Seq.UnicodeFromString("guk"));
          let _821_v110;
          _821_v110 = _dafny.Seq.UnicodeFromString("xoh");
          _820_v109 = (_820_v109).update(_817_v108, _821_v110);
          let _822_v111;
          let _nw149 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
          _822_v111 = _nw149;
          let _823_v112;
          let _nw150 = Array((new BigNumber(12)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _817_v108;
          _nw150[(_dafny.ONE).toNumber()] = _817_v108;
          _nw150[(new BigNumber(2)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(3)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(4)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(5)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(6)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(7)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(8)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(9)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(10)).toNumber()] = _817_v108;
          _nw150[(new BigNumber(11)).toNumber()] = _817_v108;
          _823_v112 = _nw150;
          let _index135 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length));
          (_823_v112)[_index135] = _817_v108;
          let _824_v113;
          _824_v113 = new BigNumber(-572);
          let _825_v114;
          _825_v114 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_824_v113));
          let _index136 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length));
          let _rhs135 = p2;
          let _rhs136 = _817_v108;
          let _rhs137 = _822_v111;
          let _rhs138 = _817_v108;
          let _rhs139 = (((_825_v114).contains((_824_v113).minus(p1))) ? ((_825_v114).get((_824_v113).minus(p1))) : (p1));
          let _lhs98 = _this;
          let _lhs99 = _823_v112;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length));
          _lhs98.f7 = _rhs135;
          _817_v108 = _rhs136;
          _822_v111 = _rhs137;
          _lhs99[_lhs100] = _rhs138;
          _824_v113 = _rhs139;
          let _826_v115;
          let _nw151 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _826_v115 = _nw151;
          _826_v115 = _826_v115;
          let _827_v116;
          _827_v116 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber(855));
          let _arr0 = (_823_v112)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length))];
          let _index137 = _module.__default.safeIndex(new BigNumber(692), new BigNumber(((_823_v112)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length))]).length));
          _arr0[_index137] = new BigNumber((_827_v116).length);
          let _arr1 = (_823_v112)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length))];
          let _index138 = _module.__default.safeIndex(new BigNumber(692), new BigNumber(((_823_v112)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_823_v112).length))]).length));
          _arr1[_index138] = new BigNumber((_dafny.Seq.UnicodeFromString("b")).length);
        }
        let _828_v117;
        _828_v117 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f7);
        let _829_v118;
        _829_v118 = _module.D0.create_DC0(p1);
        let _830_v119;
        _830_v119 = _dafny.Seq.UnicodeFromString("yher");
        r0 = ((_828_v117).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.Seq.of(_829_v118, _829_v118), p1, new BigNumber((_830_v119).length), globalState),p0))).Merge(_828_v117);
        let _831_v120;
        _831_v120 = new BigNumber(120);
        let _832_v121;
        _832_v121 = _dafny.Seq.of(p1);
        _831_v120 = ((_dafny.ZERO).minus(new BigNumber((_832_v121).length))).multipliedBy(p1);
        let _833_v122;
        let _nw152 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _833_v122 = _nw152;
        let _834_v123;
        _834_v123 = _dafny.MultiSet.fromElements(_831_v120);
        let _index139 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_833_v122).length));
        (_833_v122)[_index139] = (_dafny.ZERO).minus((((_834_v123).contains((_832_v121)[_module.__default.safeIndex(p1, new BigNumber((_832_v121).length))])) ? ((_834_v123).get((_832_v121)[_module.__default.safeIndex(p1, new BigNumber((_832_v121).length))])) : (new BigNumber((_830_v119).length))));
        let _index140 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_833_v122).length));
        (_833_v122)[_index140] = _module.__default.safeModuloInt(p1, _module.__default.safeModuloInt(p1, new BigNumber((_830_v119).length)));
      }
      let _hi1 = (p1).plus(p1);
      for (let _835_i15 = p1; _835_i15.isLessThan(_hi1); _835_i15 = _835_i15.plus(_dafny.ONE)) {
        (_this).f7 = !(p2);
        let _836_v124;
        let _nw153 = Array((new BigNumber(17)).toNumber()).fill(false);
        _836_v124 = _nw153;
        let _index141 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_836_v124).length));
        (_836_v124)[_index141] = _this.f7;
        let _837_v125;
        _837_v125 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("rwpthm"));
        let _index142 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_836_v124).length));
        (_836_v124)[_index142] = (_837_v125).IsDisjointFrom(_837_v125);
        let _838_v126;
        _838_v126 = new BigNumber(-90);
        let _839_v127;
        _839_v127 = _dafny.Seq.of((_836_v124)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_836_v124).length))], (_836_v124)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_836_v124).length))], p0);
        _838_v126 = (new BigNumber((_dafny.Seq.of((_839_v127)[_module.__default.safeIndex(_838_v126, new BigNumber((_839_v127).length))], p0, p0)).length)).plus(new BigNumber((_839_v127).length));
        let _840_v128;
        _840_v128 = _module.D6.create_DC16(p0);
        _840_v128 = _840_v128;
      }
      let _841_i16;
      _841_i16 = _dafny.ZERO;
      L2: {
        while ((_module.__default.fm0(globalState)).isEqualTo(p1)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_841_i16)) {
              break L2;
            }
            _841_i16 = (_841_i16).plus(_dafny.ONE);
            let _842_v129;
            _842_v129 = new BigNumber(-698);
            _842_v129 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_842_v129, _842_v129)).multipliedBy(_module.__default.safeModuloInt(_842_v129, new BigNumber(198))));
            let _843_v130;
            _843_v130 = new _dafny.CodePoint('p'.codePointAt(0));
            let _844_v131;
            _844_v131 = _dafny.Seq.UnicodeFromString("pmp");
            _843_v130 = ((p2) ? ((_844_v131)[_module.__default.safeIndex(p1, new BigNumber((_844_v131).length))]) : ((_this).f16));
            let _845_v132;
            let _nw154 = Array((new BigNumber(4)).toNumber()).fill(false);
            _845_v132 = _nw154;
            let _846_v133;
            _846_v133 = _dafny.Map.Empty.slice().updateUnsafe(_845_v132,_844_v131);
            let _847_v134;
            _847_v134 = _module.D3.create_DC10(_845_v132);
            _846_v133 = (_846_v133).update((_847_v134).dtor_cf21, _844_v131);
            let _848_v135;
            _848_v135 = _dafny.Seq.of(_module.D0.create_DC0(_842_v129));
            let _849_v136;
            _849_v136 = _module.D0.create_DC0((_dafny.ZERO).minus(p1));
            let _850_v137;
            _850_v137 = _dafny.Map.Empty.slice().updateUnsafe(_845_v132,p1);
            let _851_v138;
            _851_v138 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.Seq.update(_dafny.Seq.update(_848_v135, _module.__default.safeIndex(_842_v129, new BigNumber((_848_v135).length)), _849_v136), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.update(_848_v135, _module.__default.safeIndex(_842_v129, new BigNumber((_848_v135).length)), _849_v136)).length)), _849_v136), (_dafny.ZERO).minus(_842_v129), new BigNumber((_850_v137).length), globalState),_dafny.Seq.Create(_module.__default.abs(new BigNumber(485)), ((_852_v131) => function (_853_i17) {
              return _852_v131;
            })(_844_v131)));
            let _854_v139;
            _854_v139 = _dafny.Seq.of(_this.f7);
            let _855_v140;
            _855_v140 = _dafny.Map.Empty.slice().updateUnsafe(true,(_module.D3.create_DC11(new BigNumber((_854_v139).length), p1, p1, _this.f7)).dtor_cf25);
            let _856_v141;
            _856_v141 = _dafny.Seq.of(_844_v131);
            _851_v138 = (_851_v138).update((_855_v140).equals(_855_v140), _dafny.Seq.Concat(_dafny.Seq.update(_856_v141, _module.__default.safeIndex(_842_v129, new BigNumber((_856_v141).length)), _844_v131), _856_v141));
          }
        }
      }
      let _857_v142;
      _857_v142 = _module.D0.create_DC0(p1);
      let _858_v143;
      let _nw155 = Array((new BigNumber(2)).toNumber());
      _nw155[(_dafny.ZERO).toNumber()] = p1;
      _nw155[(_dafny.ONE).toNumber()] = (_857_v142).dtor_cf0;
      _858_v143 = _nw155;
      let _index143 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_858_v143).length));
      (_858_v143)[_index143] = p1;
      let _index144 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_858_v143).length));
      (_858_v143)[_index144] = p1;
      let _859_v144;
      _859_v144 = _module.D6.create_DC16(true);
      let _860_v145;
      _860_v145 = _dafny.Map.Empty.slice().updateUnsafe((_859_v144).dtor_cf30,(_858_v143)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_858_v143).length))]);
      _860_v145 = (_860_v145).update(_this.f7, p1);
      let _861_v146;
      _861_v146 = _module.D0.create_DC1(new BigNumber(-88));
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = p2;
      let _pat_let_tv17 = p0;
      let _pat_let_tv18 = p2;
      let _pat_let_tv19 = p2;
      let _pat_let_tv20 = p2;
      let _pat_let_tv21 = p2;
      let _pat_let_tv22 = p2;
      let _pat_let_tv23 = p0;
      r0 = function (_source3) {
        if (_source3.is_DC1) {
          let _862___mcc_h7 = (_source3).cf1;
          let _863_cf1 = _862___mcc_h7;
          return (((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv15,_this.f7)).update(true, _pat_let_tv16)).update(_pat_let_tv17, _this.f7)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv18,_pat_let_tv19));
        } else if (_source3.is_DC2) {
          let _864___mcc_h8 = (_source3).cf2;
          let _865___mcc_h9 = (_source3).cf3;
          let _866___mcc_h10 = (_source3).cf4;
          let _867___mcc_h11 = (_source3).cf5;
          let _868_cf5 = _867___mcc_h11;
          let _869_cf4 = _866___mcc_h10;
          let _870_cf3 = _865___mcc_h9;
          let _871_cf2 = _864___mcc_h8;
          return _dafny.Map.Empty.slice().updateUnsafe(_871_cf2,_pat_let_tv20);
        } else {
          let _872___mcc_h12 = (_source3).cf0;
          let _873_cf0 = _872___mcc_h12;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv21,_this.f7)).update((_module.D6.create_DC17(_pat_let_tv22, new BigNumber(742), (_this).f16, (_this).f16)).dtor_cf31, _pat_let_tv23);
        }
      }(_861_v146);
      let _874_v147;
      _874_v147 = _dafny.Set.fromElements(_this.f7);
      r1 = (_874_v147).Intersect((_874_v147).Intersect(_874_v147));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _875_v0;
      _875_v0 = new _dafny.CodePoint('d'.codePointAt(0));
      _875_v0 = _875_v0;
      let _876_v1;
      _876_v1 = _dafny.Seq.of(false);
      let _877_v2;
      _877_v2 = _dafny.Seq.of((_this).f15);
      let _878_v3;
      _878_v3 = new BigNumber(504);
      let _879_v4;
      _879_v4 = _module.D8.create_DC21(new BigNumber((_876_v1).length), (_877_v2)[_module.__default.safeIndex(_878_v3, new BigNumber((_877_v2).length))], _878_v3);
      let _source4 = _879_v4;
      if (_source4.is_DC21) {
        let _880___mcc_h0 = (_source4).cf38;
        let _881___mcc_h1 = (_source4).cf39;
        let _882___mcc_h2 = (_source4).cf40;
        let _883_cf40 = _882___mcc_h2;
        let _884_cf39 = _881___mcc_h1;
        let _885_cf38 = _880___mcc_h0;
        let _886_v5;
        let _nw156 = new _module.C0();
        _nw156.__ctor(!(true), _this.f7);
        _886_v5 = _nw156;
        let _887_v6;
        let _nw157 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
        _887_v6 = _nw157;
        let _888_v7;
        _888_v7 = _887_v6;
        let _889_v8;
        _889_v8 = _dafny.Seq.UnicodeFromString("jwqfqdd");
        let _890_v9;
        _890_v9 = _dafny.Map.Empty.slice().updateUnsafe(_888_v7,_889_v8);
        let _891_v10;
        _891_v10 = _dafny.Set.fromElements(_886_v5.f13, _this.f7, (_886_v5).f12, _this.f7, _886_v5.f13);
        let _892_v11;
        _892_v11 = _module.D1.create_DC3(_891_v10);
        let _893_v12;
        _893_v12 = _module.D0.create_DC0(_878_v3);
        let _894_v13;
        _894_v13 = _module.D8.create_DC22(_883_cf40, _892_v11, _893_v12);
        let _895_v14;
        _895_v14 = _dafny.Map.Empty.slice().updateUnsafe(_894_v13,_885_cf38);
        let _896_v15;
        _896_v15 = _module.D8.create_DC22((new BigNumber(((((_890_v9).contains(_887_v6)) ? ((_890_v9).get(_887_v6)) : (_889_v8))).length)).plus((((_895_v14).contains(_894_v13)) ? ((_895_v14).get(_894_v13)) : (_885_cf38))), _module.D1.create_DC3(_891_v10), _893_v12);
        let _source5 = _896_v15;
        if (_source5.is_DC21) {
          let _897___mcc_h7 = (_source5).cf38;
          let _898___mcc_h8 = (_source5).cf39;
          let _899___mcc_h9 = (_source5).cf40;
          let _900_cf40 = _899___mcc_h9;
          let _901_cf39 = _898___mcc_h8;
          let _902_cf38 = _897___mcc_h7;
          let _903_v16;
          let _nw158 = new _module.C0();
          _nw158.__ctor(true, false);
          _903_v16 = _nw158;
          _875_v0 = (_this).f16;
          (_886_v5).f13 = ((_this.f7) ? ((_886_v5).f12) : (true));
          let _904_v17;
          _904_v17 = _dafny.Map.Empty.slice().updateUnsafe(_900_cf40,new BigNumber((_889_v8).length));
          let _905_v18;
          _905_v18 = _dafny.Seq.of(new BigNumber((_891_v10).length), new BigNumber((_904_v17).length));
          let _906_v19;
          _906_v19 = _module.D1.create_DC5((_886_v5).f12, _905_v18, _900_cf40);
          let _907_v20;
          let _nw159 = new _module.C0();
          _nw159.__ctor((_902_cf38).isLessThanOrEqualTo(_878_v3), (_906_v19).dtor_cf9);
          _907_v20 = _nw159;
        } else if (_source5.is_DC22) {
          let _908___mcc_h10 = (_source5).cf41;
          let _909___mcc_h11 = (_source5).cf42;
          let _910___mcc_h12 = (_source5).cf43;
          let _911_cf43 = _910___mcc_h12;
          let _912_cf42 = _909___mcc_h11;
          let _913_cf41 = _908___mcc_h10;
          let _914_v22;
          let _init33 = ((_915_v8, _916_cf41, _917_cf40) => function (_918_i0) {
            return _dafny.Seq.of(new BigNumber(454), new BigNumber((function () {
              let _coll24 = new _dafny.Map();
              for (const _compr_24 of _dafny.IntegerRange(new BigNumber(261), new BigNumber(923))) {
                let _919_v21 = _compr_24;
                if (((new BigNumber(261)).isLessThanOrEqualTo(_919_v21)) && ((_919_v21).isLessThan(new BigNumber(923)))) {
                  _coll24.push([_module.__default.safeModuloInt(_919_v21, (_dafny.ZERO).minus(_916_cf41)),_917_cf40]);
                }
              }
              return _coll24;
            }()).length), new BigNumber((_915_v8).length));
          })(_889_v8, _913_cf41, _883_cf40);
          let _nw160 = Array((new BigNumber(17)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw160.length); _i0_33++) {
            _nw160[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _914_v22 = _nw160;
          _914_v22 = ((_886_v5.f13) ? (_914_v22) : (_914_v22));
          _913_cf41 = _913_cf41;
          _876_v1 = _876_v1;
          _878_v3 = (_913_cf41).multipliedBy(_913_cf41);
        } else {
          let _920___mcc_h13 = (_source5).cf37;
          let _921_cf37 = _920___mcc_h13;
          let _922_v23;
          _922_v23 = _dafny.Map.Empty.slice().updateUnsafe(_878_v3,_886_v5.f13);
          let _923_v24;
          _923_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(_883_cf40, _this.f7, globalState),_922_v23);
          let _rhs140 = _923_v24;
          let _rhs141 = new BigNumber(342);
          _923_v24 = _rhs140;
          _883_cf40 = _rhs141;
          let _924_v25;
          let _nw161 = Array((new BigNumber(26)).toNumber()).fill(_module.D2.Default());
          _924_v25 = _nw161;
          let _925_v26;
          _925_v26 = _dafny.Set.fromElements(new BigNumber(984));
          let _926_v27;
          _926_v27 = _module.D2.create_DC7(_925_v26);
          let _index145 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_924_v25).length));
          (_924_v25)[_index145] = _926_v27;
          let _index146 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_924_v25).length));
          (_924_v25)[_index146] = _926_v27;
          let _927_v28;
          let _nw162 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _927_v28 = _nw162;
          let _928_v29;
          _928_v29 = _dafny.Seq.of((((_876_v1)[_module.__default.safeIndex(_883_cf40, new BigNumber((_876_v1).length))]) ? (_927_v28) : (_927_v28)), _927_v28);
          _927_v28 = (_928_v29)[_module.__default.safeIndex((_883_cf40).multipliedBy(_885_cf38), new BigNumber((_928_v29).length))];
          let _index147 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_927_v28).length));
          (_927_v28)[_index147] = _875_v0;
          let _index148 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_927_v28).length));
          (_927_v28)[_index148] = (_this).f16;
        }
        (_this).f7 = !(_this.f7) || ((_this.f7) === ((_886_v5).f12));
        let _929_v30;
        _929_v30 = _dafny.MultiSet.fromElements(_878_v3, new BigNumber((_dafny.Seq.update(_876_v1, _module.__default.safeIndex(_885_cf38, new BigNumber((_876_v1).length)), _this.f7)).length));
        let _930_v31;
        let _nw163 = Array((new BigNumber(3)).toNumber());
        _nw163[(_dafny.ZERO).toNumber()] = (((_929_v30).contains(_883_cf40)) ? ((_929_v30).get(_883_cf40)) : (_883_cf40));
        _nw163[(_dafny.ONE).toNumber()] = (_878_v3).multipliedBy(_885_cf38);
        _nw163[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_889_v8, _889_v8)).length);
        _930_v31 = _nw163;
        let _index149 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_930_v31).length));
        (_930_v31)[_index149] = (_878_v3).plus(_878_v3);
        let _index150 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_930_v31).length));
        (_930_v31)[_index150] = (_dafny.ZERO).minus(_883_cf40);
      } else if (_source4.is_DC22) {
        let _931___mcc_h3 = (_source4).cf41;
        let _932___mcc_h4 = (_source4).cf42;
        let _933___mcc_h5 = (_source4).cf43;
        let _934_cf43 = _933___mcc_h5;
        let _935_cf42 = _932___mcc_h4;
        let _936_cf41 = _931___mcc_h3;
        let _937_v32;
        _937_v32 = _dafny.Set.fromElements((_this).f16);
        (_this).f7 = (_module.__default.fm30(_this.f7, _878_v3, _this.f7, globalState)).IsDisjointFrom(_937_v32);
        let _938_v33;
        _938_v33 = _dafny.Map.Empty.slice().updateUnsafe(_934_cf43,new BigNumber(437));
        let _939_v34;
        _939_v34 = _module.D2.create_DC9(new BigNumber(532), _936_cf41, _dafny.Seq.UnicodeFromString("yuln"), _this.f7, _938_v33);
        (_this).f7 = (_939_v34).dtor_cf19;
        if ((_876_v1)[_module.__default.safeIndex(_878_v3, new BigNumber((_876_v1).length))]) {
          let _940_v35;
          _940_v35 = _dafny.Seq.UnicodeFromString("reopw");
          let _rhs142 = _module.__default.safeModuloInt(_936_cf41, new BigNumber((_dafny.Seq.of(_this.f7)).length));
          let _rhs143 = _940_v35;
          _878_v3 = _rhs142;
          _940_v35 = _rhs143;
          let _941_v36;
          let _nw164 = new _module.C1();
          _nw164.__ctor(_this.f7);
          _941_v36 = _nw164;
          let _942_v37;
          _942_v37 = _dafny.Map.Empty.slice().updateUnsafe(_941_v36,_878_v3);
          let _943_v38;
          _943_v38 = _dafny.MultiSet.fromElements(new BigNumber((_942_v37).length), _878_v3);
          (_this).f7 = (_this.f7) === ((_878_v3).isLessThanOrEqualTo((((_943_v38).contains(_936_cf41)) ? ((_943_v38).get(_936_cf41)) : (_936_cf41))));
          _936_cf41 = _module.__default.safeModuloInt(_878_v3, _878_v3);
          let _944_v39;
          _944_v39 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_878_v3).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), ((_945_v0) => function (_946_i1) {
            return _945_v0;
          })(_875_v0))).length))),_878_v3);
          _944_v39 = (_944_v39).update(_878_v3, _936_cf41);
          let _947_v40;
          _947_v40 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),_936_cf41);
          _878_v3 = (_878_v3).plus((((_947_v40).contains(_941_v36.f7)) ? ((_947_v40).get(_941_v36.f7)) : (_936_cf41)));
        } else {
          let _948_v41;
          let _nw165 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _948_v41 = _nw165;
          let _949_v42;
          _949_v42 = _dafny.Seq.UnicodeFromString("yt");
          let _index151 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_948_v41).length));
          (_948_v41)[_index151] = _dafny.Seq.Concat(_949_v42, _949_v42);
          let _index152 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_948_v41).length));
          (_948_v41)[_index152] = _949_v42;
          let _950_v43;
          _950_v43 = _dafny.Set.fromElements(_this.f7, _this.f7);
          let _951_v44;
          _951_v44 = _dafny.Set.fromElements(new BigNumber(263), new BigNumber((_950_v43).length));
          let _952_v45;
          _952_v45 = _dafny.Seq.of(new BigNumber((_951_v44).length));
          let _953_v46;
          _953_v46 = _module.D9.create_DC26(_952_v45);
          _953_v46 = _953_v46;
          let _954_v47;
          let _nw166 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _954_v47 = _nw166;
          let _index153 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_954_v47).length));
          (_954_v47)[_index153] = _936_cf41;
          let _index154 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_954_v47).length));
          (_954_v47)[_index154] = _936_cf41;
          _878_v3 = _936_cf41;
          let _955_v48;
          _955_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber(((_948_v41)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_948_v41).length))]).length));
          _955_v48 = (_955_v48).update(_this.f7, _936_cf41);
        }
        let _956_v49;
        _956_v49 = _dafny.Seq.UnicodeFromString("t");
        let _957_v50;
        _957_v50 = _dafny.Map.Empty.slice().updateUnsafe(_936_cf41,_956_v49);
        _956_v49 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ix"), (((_957_v50).contains(_878_v3)) ? ((_957_v50).get(_878_v3)) : (_956_v49)));
      } else {
        let _958___mcc_h6 = (_source4).cf37;
        let _959_cf37 = _958___mcc_h6;
        let _960_v51;
        _960_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
        _960_v51 = (_960_v51).update(!(_878_v3).isEqualTo(new BigNumber(-684)), _this.f7);
        let _961_v52;
        let _nw167 = new _module.C0();
        _nw167.__ctor((_878_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_962_i2) {
          return (_this).f16;
        })).length)), (_878_v3).isLessThan(_878_v3));
        _961_v52 = _nw167;
        let _963_v53;
        _963_v53 = _dafny.Map.Empty.slice().updateUnsafe(_878_v3,(_961_v52).f12);
        let _964_v54;
        let _nw168 = Array((new BigNumber(3)).toNumber());
        _nw168[(_dafny.ZERO).toNumber()] = !(_963_v53).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), ((_965_v0) => function (_966_i3) {
          return _965_v0;
        })(_875_v0))).length));
        _nw168[(_dafny.ONE).toNumber()] = _961_v52.f13;
        _nw168[(new BigNumber(2)).toNumber()] = _961_v52.f13;
        _964_v54 = _nw168;
        let _index155 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_964_v54).length));
        (_964_v54)[_index155] = (_961_v52).f12;
        let _index156 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_964_v54).length));
        (_964_v54)[_index156] = !((_961_v52).f12);
        let _967_v55;
        _967_v55 = _dafny.Set.fromElements((_963_v53).Merge(_963_v53), _963_v53);
        let _968_v56;
        _968_v56 = _dafny.Seq.UnicodeFromString("rck");
        let _969_v58;
        _969_v58 = _dafny.Map.Empty.slice().updateUnsafe(_878_v3,_module.__default.fm28(true, (_dafny.ZERO).minus(new BigNumber(-784)), globalState));
        let _rhs144 = (_967_v55).Difference(_dafny.Set.fromElements(function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(221))) {
            let _970_v57 = _compr_25;
            if (((new BigNumber(667)).isLessThanOrEqualTo(_970_v57)) && ((_970_v57).isLessThan(new BigNumber(221)))) {
              _coll25.push([_module.__default.safeDivisionInt(_970_v57, _878_v3),_961_v52.f13]);
            }
          }
          return _coll25;
        }(), _963_v53));
        let _rhs145 = (((_969_v58).contains(((_this.f7) ? (_878_v3) : (_878_v3)))) ? ((_969_v58).get(((_this.f7) ? (_878_v3) : (_878_v3)))) : (_968_v56));
        let _rhs146 = _module.__default.fm0(globalState);
        let _rhs147 = _875_v0;
        _967_v55 = _rhs144;
        _968_v56 = _rhs145;
        _878_v3 = _rhs146;
        _875_v0 = _rhs147;
      }
      let _971_v59;
      let _nw169 = new _module.C1();
      _nw169.__ctor(_this.f7);
      _971_v59 = _nw169;
      let _972_v60;
      let _nw170 = Array((new BigNumber(3)).toNumber()).fill(false);
      _972_v60 = _nw170;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_972_v60).length))) {
        let _973_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_973_i4)) && ((_973_i4).isLessThan(new BigNumber((_972_v60).length))))) {
          (_972_v60)[(_973_i4)] = ((_dafny.MultiSet.fromElements(_878_v3)).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_878_v3)).length)))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ltjmuktis")).length)));
        }
      }
      let _974_i5;
      _974_i5 = _dafny.ZERO;
      L3: {
        while (!((_dafny.ZERO).minus(new BigNumber((_876_v1).length))).isEqualTo(new BigNumber(373))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_974_i5)) {
              break L3;
            }
            _974_i5 = (_974_i5).plus(_dafny.ONE);
            let _975_v61;
            let _nw171 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _975_v61 = _nw171;
            let _index157 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_975_v61).length));
            (_975_v61)[_index157] = _875_v0;
            let _index158 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_975_v61).length));
            (_975_v61)[_index158] = _875_v0;
            let _index159 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_975_v61).length));
            (_975_v61)[_index159] = ((_this.f7) ? (_875_v0) : (_module.__default.fm37(_this.f7, _module.D9.create_DC25(), globalState)));
            let _976_v62;
            _976_v62 = _dafny.Map.Empty.slice().updateUnsafe(_878_v3,_this.f7);
            let _977_v63;
            _977_v63 = _dafny.Map.Empty.slice().updateUnsafe(_878_v3,_878_v3);
            _976_v62 = (_976_v62).update((((_977_v63).contains(_878_v3)) ? ((_977_v63).get(_878_v3)) : (_878_v3)), _this.f7);
            let _index160 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_975_v61).length));
            (_975_v61)[_index160] = _875_v0;
          }
        }
      }
      let _978_v64;
      let _init34 = ((_979_v3, _980_v0) => function (_981_i6) {
        return _module.D6.create_DC17(_this.f7, _979_v3, (_this).f16, _980_v0);
      })(_878_v3, _875_v0);
      let _nw172 = Array((new BigNumber(16)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw172.length); _i0_34++) {
        _nw172[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _978_v64 = _nw172;
      let _982_v65;
      _982_v65 = _module.D6.create_DC17(_this.f7, _878_v3, (_this).f16, new _dafny.CodePoint('b'.codePointAt(0)));
      let _index161 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_978_v64).length));
      (_978_v64)[_index161] = _982_v65;
      let _index162 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_978_v64).length));
      (_978_v64)[_index162] = _module.__default.fm42(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,new _dafny.CodePoint('j'.codePointAt(0)))).length), (_878_v3).isLessThanOrEqualTo(_878_v3), globalState);
      return;
    }
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _983_v0;
      _983_v0 = _dafny.Seq.of(p0);
      let _984_v1;
      _984_v1 = new BigNumber(608);
      let _985_v2;
      _985_v2 = _dafny.Seq.of(_dafny.Seq.update(_983_v0, _module.__default.safeIndex(_984_v1, new BigNumber((_983_v0).length)), _this.f7), _983_v0);
      (_this).f7 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), function (_986_i0) {
        return _dafny.Seq.of(_this.f7);
      }), _985_v2);
      let _987_v3;
      _987_v3 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_984_v1);
      let _988_v4;
      _988_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,_984_v1);
      let _hi2 = new BigNumber((_988_v4).length);
      for (let _989_i1 = new BigNumber((_987_v3).length); _989_i1.isLessThan(_hi2); _989_i1 = _989_i1.plus(_dafny.ONE)) {
        let _990_v5;
        _990_v5 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_989_i1,_module.__default.fm0(globalState)), (_987_v3).Merge(_module.__default.fm29(_module.__default.fm0(globalState), p0, globalState)));
        let _991_v6;
        _991_v6 = _dafny.Seq.of(_989_i1, new BigNumber(-871));
        let _rhs148 = (_991_v6)[_module.__default.safeIndex(_984_v1, new BigNumber((_991_v6).length))];
        let _rhs149 = p1;
        let _rhs150 = _984_v1;
        let _rhs151 = _990_v5;
        let _lhs101 = _this;
        _984_v1 = _rhs148;
        _lhs101.f7 = _rhs149;
        _984_v1 = _rhs150;
        _990_v5 = _rhs151;
        _984_v1 = _989_i1;
        (_this).f7 = p0;
        let _992_v7;
        _992_v7 = _module.D6.create_DC16(p0);
        (_this).f7 = (_992_v7).dtor_cf30;
      }
      let _993_v8;
      _993_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
      let _994_v9;
      _994_v9 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,p1);
      let _995_i2;
      _995_i2 = _dafny.ZERO;
      L4: {
        while ((((_993_v8).contains((((_994_v9).contains(_984_v1)) ? ((_994_v9).get(_984_v1)) : (p1)))) ? ((_993_v8).get((((_994_v9).contains(_984_v1)) ? ((_994_v9).get(_984_v1)) : (p1)))) : (!(false)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_995_i2)) {
              break L4;
            }
            _995_i2 = (_995_i2).plus(_dafny.ONE);
            let _996_v10;
            let _nw173 = Array((new BigNumber(21)).toNumber()).fill([]);
            _996_v10 = _nw173;
            let _997_v11;
            _997_v11 = _dafny.Set.fromElements(_984_v1, new BigNumber(551));
            let _998_v12;
            _998_v12 = _dafny.Seq.of((_this).f16);
            let _999_v13;
            _999_v13 = _dafny.Set.fromElements(p1, _this.f7);
            let _1000_v14;
            _1000_v14 = _module.D1.create_DC3(_999_v13);
            let _1001_v15;
            _1001_v15 = _module.D0.create_DC0(new BigNumber((_988_v4).length));
            let _1002_v16;
            _1002_v16 = _module.D8.create_DC22(_984_v1, _1000_v14, _1001_v15);
            let _1003_v17;
            let _nw174 = Array((new BigNumber(21)).toNumber());
            _nw174[(_dafny.ZERO).toNumber()] = _984_v1;
            _nw174[(_dafny.ONE).toNumber()] = _984_v1;
            _nw174[(new BigNumber(2)).toNumber()] = _module.__default.fm0(globalState);
            _nw174[(new BigNumber(3)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(4)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(5)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(6)).toNumber()] = new BigNumber((_997_v11).length);
            _nw174[(new BigNumber(7)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(8)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(730),new BigNumber((_983_v0).length))).length);
            _nw174[(new BigNumber(10)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(11)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(12)).toNumber()] = new BigNumber((_983_v0).length);
            _nw174[(new BigNumber(13)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(14)).toNumber()] = new BigNumber((_998_v12).length);
            _nw174[(new BigNumber(15)).toNumber()] = new BigNumber(800);
            _nw174[(new BigNumber(16)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(17)).toNumber()] = _984_v1;
            _nw174[(new BigNumber(18)).toNumber()] = _module.__default.fm0(globalState);
            _nw174[(new BigNumber(19)).toNumber()] = (_1002_v16).dtor_cf41;
            _nw174[(new BigNumber(20)).toNumber()] = _984_v1;
            _1003_v17 = _nw174;
            let _1004_v18;
            _1004_v18 = _dafny.Seq.of(_1003_v17, _1003_v17);
            let _index163 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length));
            (_996_v10)[_index163] = (_1004_v18)[_module.__default.safeIndex(_984_v1, new BigNumber((_1004_v18).length))];
            let _index164 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length));
            (_996_v10)[_index164] = _1003_v17;
            if (p2) {
              let _1005_v19;
              _1005_v19 = new _dafny.CodePoint('b'.codePointAt(0));
              _1005_v19 = (_this).f16;
              _1005_v19 = (_this).f16;
              let _1006_v20;
              let _init35 = ((_1007_v12) => function (_1008_i3) {
                return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(803)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), ((_1009_v12) => function (_1010_i4) {
                  return new BigNumber((_1009_v12).length);
                })(_1007_v12)));
              })(_998_v12);
              let _nw175 = Array((new BigNumber(3)).toNumber());
              for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw175.length); _i0_35++) {
                _nw175[_i0_35] = _init35(new BigNumber(_i0_35));
              }
              _1006_v20 = _nw175;
              let _init36 = ((_1011_v4) => function (_1012_i5) {
                return _dafny.Seq.of(new BigNumber((_1011_v4).length));
              })(_988_v4);
              let _nw176 = Array((new BigNumber(14)).toNumber());
              for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw176.length); _i0_36++) {
                _nw176[_i0_36] = _init36(new BigNumber(_i0_36));
              }
              _1006_v20 = _nw176;
              _984_v1 = _984_v1;
              _998_v12 = _998_v12;
            } else {
              let _1013_v21;
              _1013_v21 = _dafny.Seq.of(_1001_v15);
              let _1014_v22;
              _1014_v22 = _dafny.Set.fromElements(_993_v8, _993_v8, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.Seq.update(_1013_v21, _module.__default.safeIndex(new BigNumber((_module.__default.fm43((_this).f16, _this.f7, p2, globalState)).length), new BigNumber((_1013_v21).length)), _module.D0.create_DC0(_984_v1)), _984_v1, _984_v1, globalState),!(!(true))));
              let _rhs152 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_998_v12, _dafny.Seq.Concat(_998_v12, _998_v12)), _module.__default.safeIndex(_984_v1, new BigNumber((_dafny.Seq.Concat(_998_v12, _dafny.Seq.Concat(_998_v12, _998_v12))).length)), (_this).f16)).length);
              let _rhs153 = _1014_v22;
              _984_v1 = _rhs152;
              _1014_v22 = _rhs153;
              let _index165 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length));
              (_996_v10)[_index165] = (_996_v10)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length))];
              let _1015_v23;
              let _nw177 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
              _1015_v23 = _nw177;
              let _1016_v24;
              _1016_v24 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_1017_i6) {
                return (_this).f16;
              }));
              let _index166 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1015_v23).length));
              (_1015_v23)[_index166] = _1016_v24;
              let _index167 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1015_v23).length));
              (_1015_v23)[_index167] = _1016_v24;
              let _rhs154 = _984_v1;
              let _rhs155 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qr"), _998_v12);
              _984_v1 = _rhs154;
              _998_v12 = _rhs155;
              let _1018_v25;
              let _nw178 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1018_v25 = _nw178;
              let _index168 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1018_v25).length));
              (_1018_v25)[_index168] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_1019_i7) {
                return (_this).f16;
              });
              let _index169 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1018_v25).length));
              (_1018_v25)[_index169] = _dafny.Seq.Concat(_998_v12, _998_v12);
            }
            let _arr2 = (_996_v10)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length))];
            let _index170 = _module.__default.safeIndex(new BigNumber(934), new BigNumber(((_996_v10)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length))]).length));
            _arr2[_index170] = _module.__default.fm0(globalState);
            let _arr3 = (_996_v10)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length))];
            let _index171 = _module.__default.safeIndex(new BigNumber(934), new BigNumber(((_996_v10)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_996_v10).length))]).length));
            _arr3[_index171] = _984_v1;
            let _1020_v26;
            _1020_v26 = _module.D6.create_DC16(true);
            (_this).f7 = (_1020_v26).dtor_cf30;
          }
        }
      }
      let _1021_v27;
      _1021_v27 = new _dafny.CodePoint('a'.codePointAt(0));
      _1021_v27 = _1021_v27;
      let _1022_v28;
      _1022_v28 = _module.D0.create_DC0(_984_v1);
      let _1023_v29;
      _1023_v29 = _dafny.Seq.of(_1022_v28, _1022_v28);
      if (((!((_984_v1).isEqualTo(new BigNumber(889)))) ? (_module.__default.fm1(_1023_v29, _984_v1, new BigNumber((_993_v8).length), globalState)) : (!(_this.f7) || (true)))) {
        _984_v1 = _984_v1;
        let _1024_v30;
        let _nw179 = new _module.C0();
        _nw179.__ctor(p0, !(((p0) ? (p1) : (p2))));
        _1024_v30 = _nw179;
        let _1025_v31;
        let _nw180 = new _module.C0();
        _nw180.__ctor(_this.f7, _1024_v30.f13);
        _1025_v31 = _nw180;
        let _1026_v32;
        _1026_v32 = _dafny.Seq.UnicodeFromString("hpqblmqyy");
        let _1027_v33;
        _1027_v33 = _dafny.Seq.of(new BigNumber((_1026_v32).length), _984_v1);
        let _1028_v34;
        _1028_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1027_v33,!((_1024_v30).f12));
        _1028_v34 = (_1028_v34).update(_1027_v33, _this.f7);
        _984_v1 = _module.__default.fm0(globalState);
      } else {
        _984_v1 = (_dafny.ZERO).minus(_984_v1);
        let _1029_v35;
        _1029_v35 = _dafny.MultiSet.fromElements(_984_v1);
        let _1030_v36;
        _1030_v36 = _dafny.MultiSet.fromElements(_984_v1, (_dafny.ZERO).minus(new BigNumber((_1029_v35).cardinality())));
        let _1031_v37;
        _1031_v37 = _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_1030_v36);
        let _1032_v38;
        _1032_v38 = _dafny.Seq.of(_987_v3, _987_v3, _987_v3, _987_v3, _dafny.Map.Empty.slice().updateUnsafe(_984_v1,_984_v1));
        let _1033_v39;
        _1033_v39 = _dafny.Seq.UnicodeFromString("fvugstq");
        let _source6 = _module.__default.fm44(_module.__default.safeModuloInt(_984_v1, _984_v1), (_dafny.Map.Empty.slice().updateUnsafe(_984_v1,_1030_v36)).Merge(_1031_v37), _1021_v27, (_1032_v38)[_module.__default.safeIndex(new BigNumber((_module.__default.fm45(new BigNumber((_1033_v39).length), new BigNumber((_1033_v39).length), globalState)).length), new BigNumber((_1032_v38).length))], globalState);
        let _1034___mcc_h0 = _source6;
        let _1035_cf47 = _1034___mcc_h0;
        _1033_v39 = _1033_v39;
        let _1036_v40;
        let _nw181 = new _module.C1();
        _nw181.__ctor((((_994_v9).contains(new BigNumber(868))) ? ((_994_v9).get(new BigNumber(868))) : (p1)));
        _1036_v40 = _nw181;
        let _1037_v41;
        let _nw182 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1037_v41 = _nw182;
        let _1038_v42;
        _1038_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1037_v41,new _dafny.CodePoint('e'.codePointAt(0)));
        let _1039_v43;
        _1039_v43 = _dafny.Set.fromElements(_this.f7);
        let _1040_v44;
        _1040_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1039_v43,_987_v3);
        let _rhs156 = _1038_v42;
        let _rhs157 = (new BigNumber(((_1040_v44).Merge(_1040_v44)).length)).minus((new BigNumber(816)).minus(_984_v1));
        let _rhs158 = !(p2);
        let _lhs102 = _this;
        _1038_v42 = _rhs156;
        _984_v1 = _rhs157;
        _lhs102.f7 = _rhs158;
        let _1041_v45;
        let _nw183 = new _module.C1();
        _nw183.__ctor(p2);
        _1041_v45 = _nw183;
        _1041_v45 = _1041_v45;
        let _1042_v46;
        _1042_v46 = _module.D1.create_DC4(p1, p2);
        let _1043_v47;
        _1043_v47 = _module.D1.create_DC6(_1042_v46);
        _1043_v47 = _1043_v47;
        _994_v9 = (_994_v9).update(_984_v1, p1);
        if ((_module.__default.fm1(_1023_v29, new BigNumber(5), (_dafny.ZERO).minus(new BigNumber(((_this).f15).length)), globalState)) === (_this.f7)) {
          _984_v1 = _984_v1;
          let _1044_v48;
          let _nw184 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1044_v48 = _nw184;
          let _1045_v49;
          _1045_v49 = _dafny.Set.fromElements(_1033_v39);
          let _1046_v50;
          _1046_v50 = _dafny.Seq.of(_1030_v36);
          let _1047_v51;
          _1047_v51 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1045_v49).length)), new BigNumber(((_1046_v50)[_module.__default.safeIndex(_984_v1, new BigNumber((_1046_v50).length))]).cardinality()));
          let _1048_v52;
          _1048_v52 = _module.D2.create_DC8(_this.f7, _1033_v39);
          let _1049_v53;
          _1049_v53 = _dafny.MultiSet.fromElements(_1048_v52);
          let _1050_v54;
          let _nw185 = Array((new BigNumber(9)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = _984_v1;
          _nw185[(_dafny.ONE).toNumber()] = _984_v1;
          _nw185[(new BigNumber(2)).toNumber()] = _984_v1;
          _nw185[(new BigNumber(3)).toNumber()] = new BigNumber((_988_v4).length);
          _nw185[(new BigNumber(4)).toNumber()] = (_1047_v51)[_module.__default.safeIndex(_984_v1, new BigNumber((_1047_v51).length))];
          _nw185[(new BigNumber(5)).toNumber()] = _984_v1;
          _nw185[(new BigNumber(6)).toNumber()] = _984_v1;
          _nw185[(new BigNumber(7)).toNumber()] = _984_v1;
          _nw185[(new BigNumber(8)).toNumber()] = new BigNumber((_1049_v53).cardinality());
          _1050_v54 = _nw185;
          let _1051_v55;
          _1051_v55 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1050_v54);
          let _nw186 = Array((new BigNumber(12)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = _1050_v54;
          _nw186[(_dafny.ONE).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(2)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(3)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(4)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(5)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(6)).toNumber()] = (((_1051_v55).contains(p1)) ? ((_1051_v55).get(p1)) : (_1050_v54));
          _nw186[(new BigNumber(7)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(8)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(9)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(10)).toNumber()] = _1050_v54;
          _nw186[(new BigNumber(11)).toNumber()] = _1050_v54;
          _1044_v48 = _nw186;
          _1033_v39 = _1033_v39;
          let _1052_v56;
          _1052_v56 = _module.D6.create_DC17(_this.f7, (_1047_v51)[_module.__default.safeIndex(_984_v1, new BigNumber((_1047_v51).length))], new _dafny.CodePoint('b'.codePointAt(0)), (_this).f16);
          let _rhs159 = (_this).f16;
          let _rhs160 = ((p1) ? ((_1052_v56).dtor_cf31) : (p1));
          let _rhs161 = _dafny.MultiSet.fromElements(_984_v1);
          let _rhs162 = (_dafny.ZERO).minus(new BigNumber(((_dafny.Set.fromElements(_984_v1, _984_v1)).Union(_module.__default.fm2(_984_v1, _this.f7, globalState))).length));
          let _rhs163 = (_this).f16;
          let _lhs103 = _this;
          _1021_v27 = _rhs159;
          _lhs103.f7 = _rhs160;
          _1030_v36 = _rhs161;
          _984_v1 = _rhs162;
          _1021_v27 = _rhs163;
          let _1053_v57;
          let _init37 = ((_1054_v0, _1055_v9) => function (_1056_i8) {
            return (_1054_v0)[_module.__default.safeIndex(new BigNumber((_1055_v9).length), new BigNumber((_1054_v0).length))];
          })(_983_v0, _994_v9);
          let _nw187 = Array((new BigNumber(6)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw187.length); _i0_37++) {
            _nw187[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1053_v57 = _nw187;
          _1053_v57 = _1053_v57;
        } else {
          let _1057_v58;
          _1057_v58 = _dafny.Seq.of(new BigNumber(-595));
          let _1058_v59;
          _1058_v59 = _dafny.Set.fromElements(_984_v1, new BigNumber((_1057_v58).length), _984_v1, (_dafny.ZERO).minus((_1057_v58)[_module.__default.safeIndex((_dafny.ZERO).minus(_984_v1), new BigNumber((_1057_v58).length))]), new BigNumber((_1030_v36).cardinality()));
          let _1059_v60;
          let _nw188 = new _module.C0();
          _nw188.__ctor(p1, (_1058_v59).IsSubsetOf(_1058_v59));
          _1059_v60 = _nw188;
          let _1060_v62;
          _1060_v62 = _dafny.Set.fromElements(function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_1030_v36).Elements) {
              let _1061_v61 = _compr_26;
              if ((_1030_v36).contains(_1061_v61)) {
                _coll26.push([_module.__default.safeModuloInt(_1061_v61, _984_v1),_984_v1]);
              }
            }
            return _coll26;
          }());
          _993_v8 = (_993_v8).update((_1060_v62).IsSubsetOf(_1060_v62), _this.f7);
          (_this).f7 = (_1059_v60).f12;
          _module.__default.m0(new BigNumber((_988_v4).length), (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(_984_v1), _984_v1)), globalState);
          _984_v1 = (_dafny.ZERO).minus((((_988_v4).contains((_983_v0)[_module.__default.safeIndex(_984_v1, new BigNumber((_983_v0).length))])) ? ((_988_v4).get((_983_v0)[_module.__default.safeIndex(_984_v1, new BigNumber((_983_v0).length))])) : (_984_v1)));
        }
      }
      let _1062_v63;
      let _nw189 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1062_v63 = _nw189;
      let _index172 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1062_v63).length));
      (_1062_v63)[_index172] = !(p1);
      let _index173 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1062_v63).length));
      (_1062_v63)[_index173] = p0;
      return;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f7 = false;
      this._f14 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f14, f7) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f7 = f7;
      return;
    }
    fm8(p0, globalState) {
      let _this = this;
      return _this.f7;
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f14;
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1063_v0;
      let _nw190 = new _module.C0();
      _nw190.__ctor(_this.f7, !(_this.f7) || (!(_this.f7)));
      _1063_v0 = _nw190;
      if (!(_this.f7)) {
        (_this).f7 = (_1063_v0).f12;
        (_1063_v0).f13 = (_1063_v0).f12;
        let _1064_v1;
        _1064_v1 = _dafny.MultiSet.fromElements(_module.D3.create_DC11(new BigNumber(273), p0, new BigNumber(187), false), _module.D3.create_DC11(new BigNumber(159), (_this).f14, (_this).f14, _this.f7));
        let _1065_v2;
        _1065_v2 = _module.D3.create_DC11((_this).f14, p0, (_this).f14, _1063_v0.f13);
        _1064_v1 = _dafny.MultiSet.fromElements(_1065_v2, _1065_v2);
        let _1066_v3;
        _1066_v3 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f14)).length)));
        let _1067_v4;
        let _nw191 = new _module.C0();
        _nw191.__ctor(_this.f7, (_1066_v3).IsProperSubsetOf(_1066_v3));
        _1067_v4 = _nw191;
        let _1068_v5;
        _1068_v5 = _dafny.Set.fromElements(_1063_v0.f13);
        let _1069_v6;
        _1069_v6 = _dafny.Map.Empty.slice().updateUnsafe((_1063_v0).f12,new BigNumber((_1068_v5).length));
        let _1070_v7;
        _1070_v7 = _1069_v6;
        let _1071_v8;
        _1071_v8 = _dafny.Set.fromElements(_1069_v6, _dafny.Map.Empty.slice().updateUnsafe(_1063_v0.f13,(_this).f14));
        let _1072_v9;
        _1072_v9 = _dafny.MultiSet.fromElements(true, _this.f7);
        let _1073_v10;
        _1073_v10 = _dafny.Seq.of(false, (_1071_v8).IsSubsetOf(_dafny.Set.fromElements(((_1070_v7)).update((_1063_v0).f12, new BigNumber(949)), _1069_v6)), (_1067_v4).f12, (_1072_v9).contains((_1063_v0).f12));
        _1073_v10 = _1073_v10;
      } else {
        if (_this.f7) {
          let _1074_v11;
          _1074_v11 = _dafny.Seq.of(_1063_v0.f13);
          let _1075_v12;
          _1075_v12 = new _dafny.CodePoint('t'.codePointAt(0));
          r1 = (_this).fm9(_1063_v0.f13, ((_this).f14).minus(new BigNumber((_dafny.Seq.update(_1074_v11, _module.__default.safeIndex(p0, new BigNumber((_1074_v11).length)), _1063_v0.f13)).length)), _1075_v12, globalState);
          (_1063_v0).f13 = !(!(_1063_v0.f13));
          let _1076_v13;
          let _nw192 = Array((new BigNumber(10)).toNumber()).fill(false);
          _1076_v13 = _nw192;
          let _1077_v14;
          _1077_v14 = _module.D3.create_DC10(_1076_v13);
          let _1078_v15;
          _1078_v15 = _module.D3.create_DC12(_1077_v14);
          let _rhs164 = p0;
          let _rhs165 = _1078_v15;
          r1 = _rhs164;
          _1078_v15 = _rhs165;
          let _1079_v16;
          _1079_v16 = _dafny.MultiSet.fromElements(_1063_v0, _1063_v0);
          let _1080_v17;
          _1080_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1079_v16).cardinality()),(_1063_v0).f12);
          let _1081_v18;
          let _nw193 = new _module.C0();
          _nw193.__ctor((((_1080_v17).contains(p0)) ? ((_1080_v17).get(p0)) : (_1063_v0.f13)), !(!((_1063_v0).f12)));
          _1081_v18 = _nw193;
          let _1082_v19;
          let _init38 = ((_1083_p0) => function (_1084_i0) {
            return _module.__default.safeDivisionInt(_1084_i0, _1083_p0);
          })(p0);
          let _nw194 = Array((new BigNumber(20)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw194.length); _i0_38++) {
            _nw194[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1082_v19 = _nw194;
          let _index174 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1082_v19).length));
          (_1082_v19)[_index174] = ((_this).f14).multipliedBy(p0);
          let _index175 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1082_v19).length));
          let _rhs166 = p0;
          let _rhs167 = ((_dafny.ZERO).minus(p0)).plus(new BigNumber(-334));
          let _rhs168 = _1081_v18;
          let _lhs104 = _1082_v19;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1082_v19).length));
          _lhs104[_lhs105] = _rhs166;
          r3 = _rhs167;
          _1081_v18 = _rhs168;
        } else {
          let _1085_v20;
          let _nw195 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1085_v20 = _nw195;
          let _index176 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1085_v20).length));
          (_1085_v20)[_index176] = p0;
          let _1086_v21;
          _1086_v21 = new _dafny.CodePoint('e'.codePointAt(0));
          let _index177 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1085_v20).length));
          let _rhs169 = (_this).fm9(_1063_v0.f13, (_this).f14, (_module.D6.create_DC17((_1063_v0).f12, p0, _1086_v21, _1086_v21)).dtor_cf33, globalState);
          let _rhs170 = _this.f7;
          let _rhs171 = p0;
          let _rhs172 = new BigNumber(811);
          let _lhs106 = _1085_v20;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1085_v20).length));
          let _lhs108 = _1063_v0;
          _lhs106[_lhs107] = _rhs169;
          _lhs108.f13 = _rhs170;
          r3 = _rhs171;
          r1 = _rhs172;
          let _1087_v22;
          _1087_v22 = _dafny.Seq.of((_this).f14, p0, p0);
          let _1088_v23;
          _1088_v23 = _dafny.MultiSet.fromElements(_1087_v22, _module.__default.fm21(_1063_v0.f13, globalState));
          _1088_v23 = _dafny.MultiSet.fromElements(_1087_v22);
          let _1089_v24;
          _1089_v24 = _dafny.Seq.of((_1063_v0).fm16(_1063_v0.f13, globalState));
          let _1090_v25;
          let _nw196 = new _module.C0();
          _nw196.__ctor((_1089_v24)[_module.__default.safeIndex((_1085_v20)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_1085_v20).length))], new BigNumber((_1089_v24).length))], _1063_v0.f13);
          _1090_v25 = _nw196;
          let _1091_v26;
          let _nw197 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _1091_v26 = _nw197;
          _1091_v26 = _1091_v26;
          let _1092_v27;
          _1092_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(769),_this.f7);
          let _1093_v28;
          _1093_v28 = _dafny.MultiSet.fromElements(false);
          let _rhs173 = _1086_v21;
          let _rhs174 = _module.__default.safeModuloInt(new BigNumber((_1093_v28).cardinality()), p0);
          let _rhs175 = _dafny.Map.Empty.slice().updateUnsafe((_1087_v22)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_1085_v20)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_1085_v20).length))])).length), new BigNumber((_1087_v22).length))],_1063_v0.f13);
          _1086_v21 = _rhs173;
          r3 = _rhs174;
          _1092_v27 = _rhs175;
        }
        let _1094_v29;
        let _init39 = ((_1095_v0) => function (_1096_i1) {
          return (_1095_v0).f12;
        })(_1063_v0);
        let _nw198 = Array((new BigNumber(19)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw198.length); _i0_39++) {
          _nw198[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1094_v29 = _nw198;
        let _1097_v30;
        _1097_v30 = _dafny.Seq.of(_1094_v29);
        _1097_v30 = _1097_v30;
        let _1098_v31;
        _1098_v31 = new _dafny.CodePoint('x'.codePointAt(0));
        _1098_v31 = _1098_v31;
        if ((_1063_v0).f12) {
          let _1099_v32;
          let _nw199 = Array((new BigNumber(12)).toNumber());
          _nw199[(_dafny.ZERO).toNumber()] = ((_this).f14).plus(new BigNumber(-954));
          _nw199[(_dafny.ONE).toNumber()] = new BigNumber((((_this.f7) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), ((_1100_v31) => function (_1101_i2) {
            return _1100_v31;
          })(_1098_v31))) : (p1))).length);
          _nw199[(new BigNumber(2)).toNumber()] = p0;
          _nw199[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(919));
          _nw199[(new BigNumber(4)).toNumber()] = new BigNumber(566);
          _nw199[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw199[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw199[(new BigNumber(7)).toNumber()] = p0;
          _nw199[(new BigNumber(8)).toNumber()] = (_this).f14;
          _nw199[(new BigNumber(9)).toNumber()] = (_this).f14;
          _nw199[(new BigNumber(10)).toNumber()] = p0;
          _nw199[(new BigNumber(11)).toNumber()] = (_this).f14;
          _1099_v32 = _nw199;
          let _index178 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1099_v32).length));
          (_1099_v32)[_index178] = new BigNumber((p1).length);
          let _1102_v33;
          _1102_v33 = _dafny.Seq.of(_this.f7, !((_1063_v0).f12));
          let _index179 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1099_v32).length));
          (_1099_v32)[_index179] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_1063_v0).f12), _1102_v33)).length);
          let _index180 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1094_v29).length));
          (_1094_v29)[_index180] = (_1063_v0).f12;
          let _index181 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1099_v32).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1099_v32).length));
          let _index183 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1094_v29).length));
          let _rhs176 = p0;
          let _rhs177 = (_dafny.ZERO).minus(p0);
          let _rhs178 = _1063_v0.f13;
          let _rhs179 = _1098_v31;
          let _lhs109 = _1099_v32;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1099_v32).length));
          let _lhs111 = _1099_v32;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1099_v32).length));
          let _lhs113 = _1094_v29;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1094_v29).length));
          _lhs109[_lhs110] = _rhs176;
          _lhs111[_lhs112] = _rhs177;
          _lhs113[_lhs114] = _rhs178;
          _1098_v31 = _rhs179;
          let _1103_v34;
          _1103_v34 = _dafny.Seq.of((_1099_v32)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_1099_v32).length))]);
          let _index184 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1099_v32).length));
          (_1099_v32)[_index184] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1063_v0.f13,new BigNumber(33))).length)).multipliedBy((_1103_v34)[_module.__default.safeIndex(p0, new BigNumber((_1103_v34).length))]);
          let _1104_v35;
          _1104_v35 = _dafny.Seq.UnicodeFromString("wloks");
          let _1105_v36;
          let _nw200 = Array((new BigNumber(17)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1105_v36 = _nw200;
          let _rhs180 = p1;
          let _rhs181 = _1105_v36;
          _1104_v35 = _rhs180;
          _1105_v36 = _rhs181;
          r3 = (new BigNumber((_1104_v35).length)).multipliedBy((_this).f14);
          r1 = new BigNumber(626);
        } else {
          r3 = p0;
          (_1063_v0).f13 = !(_1063_v0.f13);
          r3 = (_dafny.ZERO).minus(((_this).f14).plus((_this).f14));
          let _1106_v37;
          _1106_v37 = _dafny.Seq.UnicodeFromString("xyjc");
          _1106_v37 = _1106_v37;
          let _1107_v38;
          let _nw201 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1107_v38 = _nw201;
          let _index185 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1107_v38).length));
          (_1107_v38)[_index185] = p0;
          let _index186 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1107_v38).length));
          (_1107_v38)[_index186] = p0;
        }
        let _1108_v39;
        let _1109_v40;
        let _out0;
        let _out1;
        let _outcollector0 = (_this).m11(globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _1108_v39 = _out0;
        _1109_v40 = _out1;
      }
      let _1110_v41;
      _1110_v41 = _dafny.Seq.of(_this.f7, _1063_v0.f13);
      let _1111_v42;
      _1111_v42 = _dafny.Set.fromElements((_1063_v0).f12);
      let _1112_v43;
      _1112_v43 = _dafny.Set.fromElements(new BigNumber((_1111_v42).length));
      let _hi3 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1110_v41, _module.__default.safeIndex(new BigNumber((_1112_v43).length), new BigNumber((_1110_v41).length)), _1063_v0.f13), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_1063_v0).f12)).length), new BigNumber((_dafny.Seq.update(_1110_v41, _module.__default.safeIndex(new BigNumber((_1112_v43).length), new BigNumber((_1110_v41).length)), _1063_v0.f13)).length)), (_1063_v0).f12)).length);
      for (let _1113_i3 = p0; _1113_i3.isLessThan(_hi3); _1113_i3 = _1113_i3.plus(_dafny.ONE)) {
        (_1063_v0).f13 = _1063_v0.f13;
        r1 = (p0).minus((_this).f14);
        let _1114_v44;
        let _nw202 = new _module.C0();
        _nw202.__ctor(_this.f7, (_1113_i3).isLessThanOrEqualTo((_this).f14));
        _1114_v44 = _nw202;
        let _1115_v45;
        _1115_v45 = _module.D3.create_DC11(_1113_i3, (_this).f14, p0, (_1063_v0).f12);
        let _nw203 = new _module.C0();
        _nw203.__ctor(_this.f7, (_1115_v45).dtor_cf25);
        _1114_v44 = _nw203;
        let _1116_v46;
        let _init40 = ((_1117_v42, _1118_v44) => function (_1119_i4) {
          return (_dafny.Set.fromElements((_1118_v44).f12, _1118_v44.f13)).IsProperSubsetOf(_1117_v42);
        })(_1111_v42, _1114_v44);
        let _nw204 = Array((new BigNumber(25)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw204.length); _i0_40++) {
          _nw204[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1116_v46 = _nw204;
        let _index187 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1116_v46).length));
        (_1116_v46)[_index187] = _1063_v0.f13;
        let _index188 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1116_v46).length));
        let _rhs182 = _1113_i3;
        let _rhs183 = !(_1063_v0.f13) || ((_1114_v44).f12);
        let _lhs115 = _1116_v46;
        let _lhs116 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1116_v46).length));
        r1 = _rhs182;
        _lhs115[_lhs116] = _rhs183;
      }
      r3 = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
      let _1120_v47;
      let _nw205 = Array((_dafny.ONE).toNumber()).fill(false);
      _1120_v47 = _nw205;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1120_v47).length))) {
        let _1121_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1121_i5)) && ((_1121_i5).isLessThan(new BigNumber((_1120_v47).length))))) {
          (_1120_v47)[(_1121_i5)] = _1063_v0.f13;
        }
      }
      let _1122_v48;
      _1122_v48 = _dafny.MultiSet.fromElements((_this).f14);
      if (!(_1122_v48).contains((_dafny.ZERO).minus((((_1122_v48).contains(new BigNumber((p1).length))) ? ((_1122_v48).get(new BigNumber((p1).length))) : ((_this).f14))))) {
        let _1123_v49;
        _1123_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(new BigNumber(869), globalState),_1063_v0.f13);
        let _1124_v50;
        let _nw206 = Array((new BigNumber(20)).toNumber());
        _nw206[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1063_v0.f13,(_1063_v0).f12);
        _nw206[(_dafny.ONE).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(2)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(3)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(4)).toNumber()] = (_1123_v49).update((_1063_v0).f12, _1063_v0.f13);
        _nw206[(new BigNumber(5)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(6)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(7)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),_1063_v0.f13);
        _nw206[(new BigNumber(9)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(10)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(11)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(12)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(13)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(14)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(15)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(16)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(17)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(18)).toNumber()] = _1123_v49;
        _nw206[(new BigNumber(19)).toNumber()] = _1123_v49;
        _1124_v50 = _nw206;
        let _1125_v51;
        _1125_v51 = _1124_v50;
        let _1126_v52;
        let _nw207 = Array((new BigNumber(4)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = _1125_v51;
        _nw207[(_dafny.ONE).toNumber()] = _1124_v50;
        _nw207[(new BigNumber(2)).toNumber()] = _1124_v50;
        _nw207[(new BigNumber(3)).toNumber()] = _1125_v51;
        _1126_v52 = _nw207;
        _1126_v52 = _1126_v52;
        r1 = _module.__default.fm0(globalState);
        let _1127_v53;
        _1127_v53 = new _dafny.CodePoint('b'.codePointAt(0));
        r2 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_1128_v53) => function (_1129_i6) {
          return _1128_v53;
        })(_1127_v53)), _1127_v53);
        if (!(_1123_v49).contains((_1063_v0).f12)) {
          r1 = (_this).f14;
          let _1130_v54;
          _1130_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1120_v47);
          let _1131_v55;
          _1131_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1130_v54).length),(_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((p1).length), p0, new BigNumber((_1112_v43).length))).update(new BigNumber((_dafny.Seq.of((_this).f14)).length), _module.__default.abs(new BigNumber((_1112_v43).length)))).cardinality())));
          let _1132_v56;
          _1132_v56 = _dafny.Seq.of(_1131_v55);
          _1132_v56 = _1132_v56;
          _1131_v55 = (_1131_v55).update((_this).f14, new BigNumber(710));
          let _1133_v57;
          _1133_v57 = _dafny.Seq.of(new BigNumber(39), _module.__default.fm0(globalState), new BigNumber((p1).length), (_this).f14);
          let _1134_v58;
          let _nw208 = new _module.C0();
          _nw208.__ctor((_1063_v0).f12, _dafny.Seq.contains(_1133_v57, p0));
          _1134_v58 = _nw208;
          r1 = (_this).f14;
        } else {
          r2 = (_1063_v0).f12;
          _1127_v53 = _1127_v53;
          let _1135_v59;
          _1135_v59 = _dafny.Seq.of(p0);
          let _1136_v60;
          _1136_v60 = _module.D1.create_DC5(_this.f7, _dafny.Seq.update(_1135_v59, _module.__default.safeIndex((_this).f14, new BigNumber((_1135_v59).length)), p0), new BigNumber((_1135_v59).length));
          let _1137_v61;
          _1137_v61 = _dafny.MultiSet.fromElements((_1063_v0).f12);
          let _1138_v62;
          _1138_v62 = _dafny.Seq.of((_1137_v61).update((_1063_v0).f12, _module.__default.abs(p0)));
          let _1139_v64;
          let _nw209 = new _module.C0();
          _nw209.__ctor((_1063_v0).f12, (_module.__default.fm22(globalState)).IsSubsetOf(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (_1138_v62).Elements) {
              let _1140_v63 = _compr_27;
              if (_dafny.Seq.contains(_1138_v62, _1140_v63)) {
                _coll27.add(_1140_v63);
              }
            }
            return _coll27;
          }()));
          _1139_v64 = _nw209;
          let _rhs184 = (_1139_v64).fm16((_1063_v0).f12, globalState);
          let _rhs185 = _module.__default.fm23(globalState);
          let _rhs186 = (_1112_v43).Union(_1112_v43);
          let _rhs187 = _1063_v0;
          let _lhs117 = _1063_v0;
          _lhs117.f13 = _rhs184;
          _1136_v60 = _rhs185;
          _1112_v43 = _rhs186;
          _1139_v64 = _rhs187;
          (_this).f7 = _dafny.areEqual(_1135_v59, _1135_v59);
          (_1063_v0).f13 = (_1110_v41)[_module.__default.safeIndex((_this).f14, new BigNumber((_1110_v41).length))];
        }
        let _1141_v65;
        _1141_v65 = _dafny.Seq.of(p1);
        let _1142_v66;
        _1142_v66 = _dafny.Seq.of(p0, (_this).f14);
        let _1143_v67;
        _1143_v67 = _module.D3.create_DC11(new BigNumber(((((_1063_v0).f12) ? (_1141_v65) : (_1141_v65))).length), (new BigNumber((_1142_v66).length)).multipliedBy(p0), p0, (_1063_v0).f12);
        let _source7 = _1143_v67;
        if (_source7.is_DC11) {
          let _1144___mcc_h0 = (_source7).cf22;
          let _1145___mcc_h1 = (_source7).cf23;
          let _1146___mcc_h2 = (_source7).cf24;
          let _1147___mcc_h3 = (_source7).cf25;
          let _1148_cf25 = _1147___mcc_h3;
          let _1149_cf24 = _1146___mcc_h2;
          let _1150_cf23 = _1145___mcc_h1;
          let _1151_cf22 = _1144___mcc_h0;
          (_this).f7 = _1063_v0.f13;
          let _index189 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1120_v47).length));
          (_1120_v47)[_index189] = (_1063_v0).f12;
          let _index190 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1120_v47).length));
          (_1120_v47)[_index190] = (_1063_v0).f12;
          _1150_cf23 = (_dafny.ZERO).minus((_1142_v66)[_module.__default.safeIndex(p0, new BigNumber((_1142_v66).length))]);
          _1141_v65 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(471)), ((_1152_p1) => function (_1153_i7) {
            return _1152_p1;
          })(p1));
        } else if (_source7.is_DC10) {
          let _1154___mcc_h4 = (_source7).cf21;
          let _1155_cf21 = _1154___mcc_h4;
          (_this).f7 = (_this).fm8((_this).f14, globalState);
          (_1063_v0).f13 = (_1063_v0).f12;
          let _1156_v68;
          let _init41 = ((_1157_v53) => function (_1158_i8) {
            return _1157_v53;
          })(_1127_v53);
          let _nw210 = Array((new BigNumber(13)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw210.length); _i0_41++) {
            _nw210[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1156_v68 = _nw210;
          let _1159_v69;
          _1159_v69 = _dafny.MultiSet.fromElements(_1156_v68);
          _1159_v69 = _1159_v69;
          let _1160_v70;
          _1160_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1122_v48,(_1063_v0).f12);
          _1160_v70 = (_1160_v70).update((_1122_v48).Difference(_1122_v48), (((_1063_v0).f12) ? ((_1063_v0).f12) : (_1063_v0.f13)));
        } else {
          let _1161___mcc_h5 = (_source7).cf26;
          let _1162_cf26 = _1161___mcc_h5;
          let _1163_v71;
          let _nw211 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _1163_v71 = _nw211;
          let _1164_v72;
          _1164_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v0.f13,_dafny.MultiSet.fromElements(_1063_v0.f13));
          let _index191 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1163_v71).length));
          (_1163_v71)[_index191] = _1164_v72;
          let _index192 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1163_v71).length));
          let _rhs188 = _1063_v0.f13;
          let _rhs189 = (_this).fm8(p0, globalState);
          let _rhs190 = new BigNumber((_dafny.Seq.of((_1122_v48).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f14)))).length);
          let _rhs191 = (_1164_v72).Merge(_1164_v72);
          let _lhs118 = _1063_v0;
          let _lhs119 = _this;
          let _lhs120 = _1163_v71;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1163_v71).length));
          _lhs118.f13 = _rhs188;
          _lhs119.f7 = _rhs189;
          r3 = _rhs190;
          _lhs120[_lhs121] = _rhs191;
          let _index193 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_1120_v47).length));
          (_1120_v47)[_index193] = _1063_v0.f13;
          let _1165_v73;
          _1165_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1127_v53,_this.f7);
          let _index194 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_1120_v47).length));
          (_1120_v47)[_index194] = (((_this).fm8((_this).f14, globalState)) ? ((((_1165_v73).contains(_1127_v53)) ? ((_1165_v73).get(_1127_v53)) : ((_1063_v0).f12))) : ((_1063_v0).f12));
          (_1063_v0).f13 = _this.f7;
          let _1166_v74;
          let _init42 = function (_1167_i9) {
            return _dafny.MultiSet.fromElements((_this).f14);
          };
          let _nw212 = Array((new BigNumber(13)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw212.length); _i0_42++) {
            _nw212[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1166_v74 = _nw212;
          _1166_v74 = (_1166_v74);
        }
      } else {
        r3 = (_this).f14;
        let _1168_v75;
        let _nw213 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _1168_v75 = _nw213;
        let _index195 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1168_v75).length));
        (_1168_v75)[_index195] = new BigNumber(((_1063_v0).fm17((_this).f14, _this.f7, globalState)).length);
        let _index196 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1168_v75).length));
        (_1168_v75)[_index196] = p0;
        let _1169_v76;
        _1169_v76 = _dafny.Seq.of(p0, p0, (_this).f14);
        let _1170_v77;
        _1170_v77 = _dafny.Seq.of((new BigNumber((_1169_v76).length)).multipliedBy((_this).f14));
        let _index197 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1168_v75).length));
        (_1168_v75)[_index197] = (_1169_v76)[_module.__default.safeIndex(p0, new BigNumber((_1169_v76).length))];
        let _1171_v78;
        let _nw214 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1171_v78 = _nw214;
        let _rhs192 = _1171_v78;
        let _rhs193 = (_this).f14;
        let _rhs194 = p0;
        _1171_v78 = _rhs192;
        r1 = _rhs193;
        r3 = _rhs194;
        let _1172_v79;
        _1172_v79 = _dafny.Set.fromElements(_1120_v47);
        let _1173_v80;
        _1173_v80 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1174_v81;
        let _nw215 = new _module.C2();
        _nw215.__ctor(_1172_v79, _1173_v80, (_1063_v0).f12);
        _1174_v81 = _nw215;
        let _1175_v82;
        _1175_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1174_v81,(_this).f14);
        let _1176_v83;
        _1176_v83 = _module.D3.create_DC11(p0, (_this).f14, (((_1175_v82).contains(_1174_v81)) ? ((_1175_v82).get(_1174_v81)) : ((_this).f14)), _1174_v81.f7);
        _1176_v83 = _1176_v83;
      }
      let _1177_v84;
      _1177_v84 = _module.D3.create_DC10(_1120_v47);
      let _1178_v85;
      _1178_v85 = _dafny.Seq.of(_1120_v47);
      let _1179_v86;
      _1179_v86 = _dafny.Seq.of(_1120_v47, (_1178_v85)[_module.__default.safeIndex(new BigNumber(-433), new BigNumber((_1178_v85).length))]);
      let _1180_v87;
      let _nw216 = Array((new BigNumber(27)).toNumber());
      _nw216[(_dafny.ZERO).toNumber()] = _1120_v47;
      _nw216[(_dafny.ONE).toNumber()] = (_1177_v84).dtor_cf21;
      _nw216[(new BigNumber(2)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(3)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(4)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(5)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(6)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(7)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(8)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(9)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(10)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(11)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(12)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(13)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(14)).toNumber()] = (((_1063_v0).f12) ? (_1120_v47) : (_1120_v47));
      _nw216[(new BigNumber(15)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(16)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(17)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(18)).toNumber()] = (_1179_v86)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_1179_v86).length))];
      _nw216[(new BigNumber(19)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(20)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(21)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(22)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(23)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(24)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(25)).toNumber()] = _1120_v47;
      _nw216[(new BigNumber(26)).toNumber()] = _1120_v47;
      _1180_v87 = _nw216;
      r0 = _1180_v87;
      r1 = (_this).f14;
      let _1181_v88;
      _1181_v88 = _module.D0.create_DC0(p0);
      let _pat_let_tv24 = _1122_v48;
      let _pat_let_tv25 = p0;
      let _pat_let_tv26 = _1112_v43;
      let _pat_let_tv27 = _1063_v0;
      let _pat_let_tv28 = _1063_v0;
      r2 = function (_source8) {
        if (_source8.is_DC1) {
          let _1182___mcc_h6 = (_source8).cf1;
          let _1183_cf1 = _1182___mcc_h6;
          return ((_module.D12.create_DC29(_pat_let_tv24)).dtor_cf48).IsSubsetOf(_dafny.MultiSet.fromElements(_pat_let_tv25, new BigNumber((_pat_let_tv26).length)));
        } else if (_source8.is_DC2) {
          let _1184___mcc_h7 = (_source8).cf2;
          let _1185___mcc_h8 = (_source8).cf3;
          let _1186___mcc_h9 = (_source8).cf4;
          let _1187___mcc_h10 = (_source8).cf5;
          let _1188_cf5 = _1187___mcc_h10;
          let _1189_cf4 = _1186___mcc_h9;
          let _1190_cf3 = _1185___mcc_h8;
          let _1191_cf2 = _1184___mcc_h7;
          return (_pat_let_tv27).f12;
        } else {
          let _1192___mcc_h11 = (_source8).cf0;
          let _1193_cf0 = _1192___mcc_h11;
          return !((_pat_let_tv28).f12);
        }
      }(_1181_v88);
      r3 = (new BigNumber((_1110_v41).length)).multipliedBy(((_this).f14).multipliedBy(p0));
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      let _1194_v0;
      let _init43 = function (_1195_i1) {
        return !(_this.f7);
      };
      let _nw217 = Array((new BigNumber(28)).toNumber());
      for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw217.length); _i0_43++) {
        _nw217[_i0_43] = _init43(new BigNumber(_i0_43));
      }
      _1194_v0 = _nw217;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1194_v0).length))) {
        let _1196_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1196_i0)) && ((_1196_i0).isLessThan(new BigNumber((_1194_v0).length))))) {
          (_1194_v0)[(_1196_i0)] = _this.f7;
        }
      }
      let _1197_v1;
      let _1198_v2;
      let _1199_v3;
      let _1200_v4;
      let _out2;
      let _out3;
      let _out4;
      let _out5;
      let _outcollector1 = (_this).m12(globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _out4 = _outcollector1[2];
      _out5 = _outcollector1[3];
      _1197_v1 = _out2;
      _1198_v2 = _out3;
      _1199_v3 = _out4;
      _1200_v4 = _out5;
      let _1201_i2;
      _1201_i2 = _dafny.ZERO;
      L5: {
        while ((p1).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), function (_1208_i3) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1201_i2)) {
              break L5;
            }
            _1201_i2 = (_1201_i2).plus(_dafny.ONE);
            let _1202_v5;
            _1202_v5 = new BigNumber(-13);
            let _1203_v6;
            _1203_v6 = _dafny.Set.fromElements(p2);
            let _rhs195 = !((_1203_v6).Intersect(_1203_v6)).equals((_1203_v6).Intersect(_1203_v6));
            let _rhs196 = (_1202_v5).plus(new BigNumber(679));
            let _lhs122 = _this;
            _lhs122.f7 = _rhs195;
            _1202_v5 = _rhs196;
            r1 = _1203_v6;
            let _1204_v7;
            _1204_v7 = _dafny.Seq.UnicodeFromString("pqvfuxln");
            let _1205_v8;
            _1205_v8 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1204_v7).length),p2)).length));
            _1197_v1 = ((_1205_v8).Intersect(_1205_v8)).IsDisjointFrom((_1205_v8).Union(_1205_v8));
            let _1206_v9;
            _1206_v9 = new _dafny.CodePoint('x'.codePointAt(0));
            let _1207_v10;
            _1207_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v5,_1206_v9);
            _1207_v10 = (_1207_v10).update(_1202_v5, _1206_v9);
          }
        }
      }
      let _index198 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1194_v0).length));
      (_1194_v0)[_index198] = p0;
      let _1209_v11;
      _1209_v11 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),false);
      let _index199 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1194_v0).length));
      (_1194_v0)[_index199] = !(_1209_v11).contains(_1197_v1);
      let _1210_v12;
      _1210_v12 = _module.D3.create_DC10(_1194_v0);
      let _pat_let_tv29 = _1194_v0;
      _1210_v12 = function (_pat_let31_0) {
        return function (_1211_dt__update__tmp_h0) {
          return function (_pat_let32_0) {
            return function (_1212_dt__update_hcf21_h0) {
              return _module.D3.create_DC10(_1212_dt__update_hcf21_h0);
            }(_pat_let32_0);
          }(_pat_let_tv29);
        }(_pat_let31_0);
      }(_1210_v12);
      let _1213_v13;
      _1213_v13 = _dafny.MultiSet.fromElements((_this).f14, p1, p1);
      let _1214_v14;
      _1214_v14 = new BigNumber(484);
      let _rhs197 = _1213_v13;
      let _rhs198 = (p1).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f14))), _1214_v14));
      let _rhs199 = (_this).f14;
      _1213_v13 = _rhs197;
      _1214_v14 = _rhs198;
      _1214_v14 = _rhs199;
      let _1215_v15;
      _1215_v15 = _dafny.Set.fromElements(_1200_v4, _1198_v2);
      r0 = (_1209_v11).update(p2, (new BigNumber((_1215_v15).length)).isLessThan(_1214_v14));
      let _1216_v16;
      _1216_v16 = _dafny.Seq.of(_this.f7);
      let _1217_v17;
      _1217_v17 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1218_v18;
      _1218_v18 = _dafny.Set.fromElements(_this.f7, true);
      r1 = (_module.__default.fm27((_1216_v16)[_module.__default.safeIndex((_this).f14, new BigNumber((_1216_v16).length))], _1217_v17, _1199_v3, p0, globalState)).Difference((_1218_v18).Difference(_1218_v18));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _1219_v1;
      _1219_v1 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f14),(_this).f14);
      let _1220_v2;
      _1220_v2 = _dafny.Seq.of(new BigNumber((function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_1219_v1).Keys.Elements) {
          let _1221_v0 = _compr_28;
          if ((_1219_v1).contains(_1221_v0)) {
            _coll28.push([(_1221_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("exnsnbf")).length)),(_this).f14]);
          }
        }
        return _coll28;
      }()).length));
      if (!(_dafny.Seq.IsProperPrefixOf(_1220_v2, _1220_v2))) {
        let _1222_v3;
        _1222_v3 = new BigNumber(980);
        let _1223_v4;
        _1223_v4 = _dafny.Seq.UnicodeFromString("kd");
        _1222_v3 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1223_v4, _1223_v4), _dafny.Seq.UnicodeFromString("aa"))).length);
        let _1224_v5;
        _1224_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1223_v4);
        _1222_v3 = new BigNumber(((((_this.f7) ? (_1224_v5) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1223_v4)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_1225_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })))).length);
        _1222_v3 = (_1220_v2)[_module.__default.safeIndex(_1222_v3, new BigNumber((_1220_v2).length))];
        let _1226_v6;
        _1226_v6 = _module.D1.create_DC6(_module.D1.create_DC4(_this.f7, _this.f7));
        let _1227_v7;
        _1227_v7 = _dafny.Seq.of(true);
        let _1228_v8;
        _1228_v8 = _dafny.Seq.of(_dafny.Seq.update(_1227_v7, _module.__default.safeIndex(_1222_v3, new BigNumber((_1227_v7).length)), _this.f7), _dafny.Seq.Concat(_1227_v7, _1227_v7));
        let _rhs200 = _1226_v6;
        let _rhs201 = ((_this.f7) ? (_this.f7) : ((_1222_v3).isLessThanOrEqualTo(_1222_v3)));
        let _rhs202 = (_this).f14;
        let _rhs203 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-714)), ((_1229_v7) => function (_1230_i1) {
          return _dafny.Seq.Concat(_1229_v7, _1229_v7);
        })(_1227_v7));
        let _lhs123 = _this;
        _1226_v6 = _rhs200;
        _lhs123.f7 = _rhs201;
        _1222_v3 = _rhs202;
        _1228_v8 = _rhs203;
        let _1231_v9;
        _1231_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,true);
        let _1232_v10;
        _1232_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("f"), _1223_v4),(((_1231_v9).contains((_this).f14)) ? ((_1231_v9).get((_this).f14)) : (_this.f7)));
        let _1233_v11;
        _1233_v11 = _module.D2.create_DC8(_this.f7, _1223_v4);
        (_this).f7 = (((_1232_v10).contains(!(_this.f7) || (_this.f7))) ? ((_1232_v10).get(!(_this.f7) || (_this.f7))) : ((_1233_v11).dtor_cf14));
      } else {
        let _1234_v12;
        _1234_v12 = _dafny.Seq.of(_this.f7, _this.f7, true, _this.f7);
        if (((_1234_v12)[_module.__default.safeIndex((_this).f14, new BigNumber((_1234_v12).length))]) || (_this.f7)) {
          let _1235_v13;
          _1235_v13 = _module.D6.create_DC16(!((((_1234_v12)[_module.__default.safeIndex(new BigNumber((_1220_v2).length), new BigNumber((_1234_v12).length))]) ? (!(_this.f7)) : (_this.f7))));
          _1235_v13 = function (_pat_let33_0) {
            return function (_1236_dt__update__tmp_h0) {
              return function (_pat_let34_0) {
                return function (_1237_dt__update_hcf30_h0) {
                  return _module.D6.create_DC16(_1237_dt__update_hcf30_h0);
                }(_pat_let34_0);
              }(_this.f7);
            }(_pat_let33_0);
          }(_1235_v13);
          let _1238_v14;
          _1238_v14 = new BigNumber(305);
          let _1239_v15;
          let _nw218 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1239_v15 = _nw218;
          let _1240_v16;
          _1240_v16 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qasccg")).length), new BigNumber((_dafny.MultiSet.fromElements(_1239_v15, _1239_v15, _1239_v15)).cardinality()), _1238_v14);
          _1238_v14 = (_1238_v14).multipliedBy((_1238_v14).minus(new BigNumber((_1240_v16).length)));
          let _1241_v17;
          _1241_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new _dafny.CodePoint('d'.codePointAt(0)));
          let _1242_v18;
          _1242_v18 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1243_v19;
          _1243_v19 = _module.D0.create_DC0((_this).f14);
          let _1244_v20;
          _1244_v20 = _dafny.Seq.of(_1243_v19);
          let _1245_v21;
          _1245_v21 = _dafny.MultiSet.fromElements(_this.f7, _module.__default.fm1(_1244_v20, (_this).f14, _1238_v14, globalState));
          let _1246_v22;
          _1246_v22 = _dafny.Map.Empty.slice().updateUnsafe((((_1241_v17).contains(_this.f7)) ? ((_1241_v17).get(_this.f7)) : (_1242_v18)),new BigNumber((_1245_v21).cardinality()));
          let _1247_v23;
          let _nw219 = new _module.C0();
          _nw219.__ctor(!(!(_1219_v1).contains(new BigNumber((_1246_v22).length))), _this.f7);
          _1247_v23 = _nw219;
          let _1248_v24;
          _1248_v24 = _dafny.Seq.UnicodeFromString("twgvgcqtp");
          let _1249_v25;
          let _nw220 = Array((new BigNumber(2)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = _1238_v14;
          _nw220[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1248_v24).length));
          _1249_v25 = _nw220;
          let _1250_v26;
          _1250_v26 = _module.D9.create_DC23(_1249_v25);
          let _1251_v27;
          _1251_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v26,_1249_v25);
          let _1252_v28;
          _1252_v28 = _dafny.MultiSet.fromElements(_1251_v27, (_1251_v27).update(_1250_v26, _1249_v25), _1251_v27);
          _module.__default.m0((((_1252_v28).contains(_1251_v27)) ? ((_1252_v28).get(_1251_v27)) : ((_this).f14)), ((_this).f14).multipliedBy(_1238_v14), globalState);
          let _1253_v29;
          _1253_v29 = _module.D3.create_DC10(_1239_v15);
          let _1254_v30;
          _1254_v30 = _module.D3.create_DC12(_1253_v29);
          let _pat_let_tv30 = globalState;
          let _pat_let_tv31 = _1238_v14;
          let _pat_let_tv32 = _1247_v23;
          let _1255_v31;
          _1255_v31 = _dafny.MultiSet.fromElements(function (_pat_let35_0) {
            return function (_1256_dt__update__tmp_h1) {
              return function (_pat_let36_0) {
                return function (_1257_dt__update_hcf26_h0) {
                  return _module.D3.create_DC12(_1257_dt__update_hcf26_h0);
                }(_pat_let36_0);
              }(_module.D3.create_DC11(new BigNumber(562), new BigNumber((_module.__default.fm46((_this).f14, _pat_let_tv30)).length), _pat_let_tv31, (_pat_let_tv32).f12));
            }(_pat_let35_0);
          }(_1254_v30));
          _1255_v31 = _1255_v31;
        } else {
          let _1258_v32;
          let _init44 = function (_1259_i2) {
            return (_1259_i2).plus((_dafny.ZERO).minus((_this).f14));
          };
          let _nw221 = Array((new BigNumber(9)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw221.length); _i0_44++) {
            _nw221[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1258_v32 = _nw221;
          let _1260_v33;
          _1260_v33 = _dafny.MultiSet.fromElements(_1258_v32, _1258_v32, _1258_v32);
          let _index200 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length));
          (_1258_v32)[_index200] = new BigNumber((_1260_v33).cardinality());
          let _index201 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length));
          (_1258_v32)[_index201] = new BigNumber((_1220_v2).length);
          let _1261_v34;
          _1261_v34 = new _dafny.CodePoint('b'.codePointAt(0));
          let _1262_v35;
          _1262_v35 = _dafny.Set.fromElements(_1261_v34);
          let _1263_v36;
          _1263_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1261_v34);
          let _1264_v38;
          _1264_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1263_v36,function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of (_dafny.Set.fromElements(_1261_v34)).Elements) {
              let _1265_v37 = _compr_29;
              if ((_dafny.Set.fromElements(_1261_v34)).contains(_1265_v37)) {
                _coll29.add(_1265_v37);
              }
            }
            return _coll29;
          }());
          let _1266_v39;
          _1266_v39 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1261_v34,(_this).f14));
          let _1267_v41;
          let _nw222 = Array((new BigNumber(12)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), _1261_v34, _1261_v34);
          _nw222[(_dafny.ONE).toNumber()] = _1262_v35;
          _nw222[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_1261_v34, _1261_v34, _1261_v34);
          _nw222[(new BigNumber(3)).toNumber()] = (_1262_v35).Difference(_1262_v35);
          _nw222[(new BigNumber(4)).toNumber()] = _1262_v35;
          _nw222[(new BigNumber(5)).toNumber()] = (((_1264_v38).contains(_1263_v36)) ? ((_1264_v38).get(_1263_v36)) : (_1262_v35));
          _nw222[(new BigNumber(6)).toNumber()] = _1262_v35;
          _nw222[(new BigNumber(7)).toNumber()] = (_1262_v35).Union(_1262_v35);
          _nw222[(new BigNumber(8)).toNumber()] = (_1262_v35).Intersect(_1262_v35);
          _nw222[(new BigNumber(9)).toNumber()] = function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of ((_1266_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]), new BigNumber((_1266_v39).length))]).Keys.Elements) {
              let _1268_v40 = _compr_30;
              if (((_1266_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]), new BigNumber((_1266_v39).length))]).contains(_1268_v40)) {
                _coll30.add(_1268_v40);
              }
            }
            return _coll30;
          }();
          _nw222[(new BigNumber(10)).toNumber()] = (_1262_v35).Union(_1262_v35);
          _nw222[(new BigNumber(11)).toNumber()] = _1262_v35;
          _1267_v41 = _nw222;
          let _index202 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_1267_v41).length));
          (_1267_v41)[_index202] = _1262_v35;
          let _index203 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_1267_v41).length));
          (_1267_v41)[_index203] = ((_1262_v35).Difference(_1262_v35)).Intersect(_1262_v35);
          let _1269_v42;
          _1269_v42 = _dafny.Set.fromElements((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))], (_this).f14);
          let _1270_v43;
          _1270_v43 = _dafny.MultiSet.fromElements(((_this).f14).multipliedBy((_this).f14), (_dafny.ZERO).minus((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]), _module.__default.safeDivisionInt((_this).f14, new BigNumber((_1269_v42).length)), (_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]);
          _1270_v43 = (_module.__default.fm33(new BigNumber(257), _1261_v34, globalState)).Difference(_1270_v43);
          let _1271_v44;
          _1271_v44 = _module.D0.create_DC0((_this).f14);
          let _1272_v45;
          _1272_v45 = _dafny.Seq.of(_1271_v44);
          let _1273_v46;
          _1273_v46 = _dafny.Seq.UnicodeFromString("htrgmixse");
          let _1274_v47;
          _1274_v47 = _dafny.Seq.of(_module.__default.fm21(_module.__default.fm1(_1272_v45, new BigNumber((_1273_v46).length), (_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))], globalState), globalState));
          _1220_v2 = (_1274_v47)[_module.__default.safeIndex(((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]).plus((_1258_v32)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1258_v32).length))]), new BigNumber((_1274_v47).length))];
          let _1275_v48;
          let _nw223 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _1275_v48 = _nw223;
          let _index204 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_1275_v48).length));
          (_1275_v48)[_index204] = _1220_v2;
          let _index205 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_1275_v48).length));
          let _rhs204 = _1220_v2;
          let _rhs205 = _1273_v46;
          let _lhs124 = _1275_v48;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_1275_v48).length));
          _lhs124[_lhs125] = _rhs204;
          _1273_v46 = _rhs205;
        }
        (_this).f7 = true;
        let _1276_v49;
        _1276_v49 = _dafny.Set.fromElements((_this).f14);
        (_this).f7 = ((_1276_v49).Union(_1276_v49)).IsProperSubsetOf(_1276_v49);
        let _1277_v50;
        _1277_v50 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, _this.f7, true);
        if ((_module.__default.safeDivisionInt(new BigNumber((_1277_v50).cardinality()), (_this).f14)).isLessThanOrEqualTo((_this).f14)) {
          (_this).f7 = (_module.D1.create_DC4(!(true), _this.f7)).dtor_cf7;
          (_this).f7 = _this.f7;
          (_this).f7 = _this.f7;
          let _1278_v51;
          _1278_v51 = _dafny.Seq.UnicodeFromString("tu");
          let _1279_v52;
          _1279_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1278_v51, _1278_v51),_this.f7);
          _1279_v52 = (_1279_v52).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_1280_i3) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }), _this.f7);
          (_this).f7 = _this.f7;
        } else {
          let _1281_v53;
          _1281_v53 = _dafny.Seq.UnicodeFromString("wsnf");
          let _1282_v54;
          let _nw224 = Array((new BigNumber(21)).toNumber());
          _nw224[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw224[(_dafny.ONE).toNumber()] = _this.f7;
          _nw224[(new BigNumber(2)).toNumber()] = (_this).fm8(_module.__default.fm0(globalState), globalState);
          _nw224[(new BigNumber(3)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(4)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(5)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(6)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(7)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(8)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(9)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(10)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(11)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(12)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(13)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(14)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(15)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(16)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(17)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(18)).toNumber()] = _this.f7;
          _nw224[(new BigNumber(19)).toNumber()] = (_this).fm8(new BigNumber((_1281_v53).length), globalState);
          _nw224[(new BigNumber(20)).toNumber()] = _this.f7;
          _1282_v54 = _nw224;
          let _1283_v55;
          _1283_v55 = _dafny.Set.fromElements(_1282_v54);
          let _1284_v56;
          _1284_v56 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1285_v57;
          let _nw225 = new _module.C2();
          _nw225.__ctor(_1283_v55, _1284_v56, _this.f7);
          _1285_v57 = _nw225;
          _1285_v57 = _1285_v57;
          let _1286_v58;
          _1286_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1277_v50).IsSubsetOf(_1277_v50),_this.f7);
          _1286_v58 = (_1286_v58).update(!(_this.f7) || (_this.f7), _dafny.areEqual(_1281_v53, _dafny.Seq.update(_1281_v53, _module.__default.safeIndex(new BigNumber(732), new BigNumber((_1281_v53).length)), (_1285_v57).f16)));
          let _1287_v59;
          _1287_v59 = new BigNumber(-744);
          _1287_v59 = (_this).f14;
          _module.__default.m0(_1287_v59, _1287_v59, globalState);
          let _1288_v60;
          let _nw226 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1288_v60 = _nw226;
          let _1289_v61;
          _1289_v61 = _dafny.Set.fromElements(_this.f7, _this.f7);
          let _index206 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1282_v54).length));
          (_1282_v54)[_index206] = (_1289_v61).IsProperSubsetOf(_1289_v61);
          let _1290_v62;
          _1290_v62 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("acfbl"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-775)), ((_1291_v57) => function (_1292_i4) {
            return (_1291_v57).f16;
          })(_1285_v57)), _1281_v53);
          let _index207 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1282_v54).length));
          let _rhs206 = _dafny.Seq.Concat((_1290_v62)[_module.__default.safeIndex(new BigNumber((_1276_v49).length), new BigNumber((_1290_v62).length))], _1281_v53);
          let _rhs207 = _1288_v60;
          let _rhs208 = _this.f7;
          let _rhs209 = _dafny.Seq.UnicodeFromString("acvrlugbv");
          let _lhs126 = _1282_v54;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1282_v54).length));
          _1281_v53 = _rhs206;
          _1288_v60 = _rhs207;
          _lhs126[_lhs127] = _rhs208;
          _1281_v53 = _rhs209;
        }
        let _1293_v63;
        _1293_v63 = _dafny.Set.fromElements(_this.f7, _this.f7, _this.f7, _this.f7, _this.f7);
        let _1294_v64;
        _1294_v64 = _module.D1.create_DC3(_1293_v63);
        let _source9 = _1294_v64;
        if (_source9.is_DC4) {
          let _1295___mcc_h0 = (_source9).cf7;
          let _1296___mcc_h1 = (_source9).cf8;
          let _1297_cf8 = _1296___mcc_h1;
          let _1298_cf7 = _1295___mcc_h0;
          let _1299_v65;
          _1299_v65 = _module.D0.create_DC0((_this).f14);
          let _1300_v66;
          _1300_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1293_v63,false);
          let _1301_v67;
          _1301_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.Seq.of(_1299_v65, _1299_v65), (_this).f14, new BigNumber(449), globalState),_1300_v66);
          _1301_v67 = (_1301_v67).update(_this.f7, _1300_v66);
          let _1302_v68;
          let _nw227 = new _module.C1();
          _nw227.__ctor(_1297_cf8);
          _1302_v68 = _nw227;
          let _1303_v69;
          _1303_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1297_cf8,_1302_v68);
          let _1304_v70;
          _1304_v70 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1305_v71;
          _1305_v71 = _dafny.MultiSet.fromElements(((_this).f14).plus(new BigNumber(((_1303_v69).update(_1298_cf7, _1302_v68)).length)), (_this).f14, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1304_v70,_1302_v68.f7)).length)));
          let _1306_v72;
          _1306_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1298_cf7,(_this).f14);
          _1305_v71 = (((!(_1298_cf7)) ? ((_dafny.MultiSet.fromElements((_this).f14)).update((_this).f14, _module.__default.abs((_this).f14))) : (_1305_v71))).Intersect(_dafny.MultiSet.fromElements((_this).f14, (((_1306_v72).contains(_1297_cf8)) ? ((_1306_v72).get(_1297_cf8)) : ((_this).f14))));
          let _1307_v73;
          _1307_v73 = _dafny.Seq.of(_1304_v70, _1304_v70);
          let _1308_v74;
          _1308_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1220_v2,(_1307_v73)[_module.__default.safeIndex((_this).f14, new BigNumber((_1307_v73).length))]);
          _1308_v74 = _1308_v74;
          (_1302_v68).f7 = !(_1297_cf8) || (_1302_v68.f7);
        } else if (_source9.is_DC5) {
          let _1309___mcc_h2 = (_source9).cf9;
          let _1310___mcc_h3 = (_source9).cf10;
          let _1311___mcc_h4 = (_source9).cf11;
          let _1312_cf11 = _1311___mcc_h4;
          let _1313_cf10 = _1310___mcc_h3;
          let _1314_cf9 = _1309___mcc_h2;
          _1219_v1 = (_1219_v1).update((_this).f14, (_1312_cf11).plus(_1312_cf11));
          let _1315_v75;
          _1315_v75 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1316_v76;
          _1316_v76 = _module.D6.create_DC17(_1314_cf9, new BigNumber(579), new _dafny.CodePoint('f'.codePointAt(0)), _1315_v75);
          let _1317_v77;
          _1317_v77 = _module.D6.create_DC18(_1316_v76);
          let _1318_v78;
          let _init45 = function (_1319_i5) {
            return false;
          };
          let _nw228 = Array((new BigNumber(19)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw228.length); _i0_45++) {
            _nw228[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1318_v78 = _nw228;
          let _1320_v79;
          _1320_v79 = _module.D12.create_DC30(new BigNumber(335), _1314_cf9, _1220_v2, (_this).f14);
          let _pat_let_tv33 = _1220_v2;
          let _rhs210 = _module.__default.fm47((function (_pat_let37_0) {
            return function (_1321_dt__update__tmp_h2) {
              return function (_pat_let38_0) {
                return function (_1322_dt__update_hcf49_h0) {
                  return _module.D12.create_DC30(_1322_dt__update_hcf49_h0, (_1321_dt__update__tmp_h2).dtor_cf50, (_1321_dt__update__tmp_h2).dtor_cf51, (_1321_dt__update__tmp_h2).dtor_cf52);
                }(_pat_let38_0);
              }(new BigNumber((_pat_let_tv33).length));
            }(_pat_let37_0);
          }(_1320_v79)).dtor_cf51, _module.__default.safeModuloInt(_1312_cf11, _1312_cf11), globalState);
          let _rhs211 = _1318_v78;
          let _rhs212 = _1318_v78;
          _1317_v77 = _rhs210;
          _1318_v78 = _rhs211;
          _1318_v78 = _rhs212;
          _1234_v12 = _1234_v12;
          let _1323_v80;
          let _init46 = function (_1324_i6) {
            return (_1324_i6).plus(new BigNumber((_dafny.Seq.UnicodeFromString("gdgr")).length));
          };
          let _nw229 = Array((new BigNumber(27)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw229.length); _i0_46++) {
            _nw229[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1323_v80 = _nw229;
          _1323_v80 = _1323_v80;
        } else if (_source9.is_DC3) {
          let _1325___mcc_h5 = (_source9).cf6;
          let _1326_cf6 = _1325___mcc_h5;
          let _1327_v81;
          let _nw230 = new _module.C1();
          _nw230.__ctor(_this.f7);
          _1327_v81 = _nw230;
          (_this).f7 = _this.f7;
          let _1328_v82;
          _1328_v82 = _dafny.Seq.UnicodeFromString("s");
          let _1329_v83;
          _1329_v83 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1330_v84;
          _1330_v84 = _dafny.MultiSet.fromElements((_this).f14);
          let _1331_v85;
          _1331_v85 = _module.D12.create_DC29(_1330_v84);
          let _1332_v86;
          let _nw231 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1332_v86 = _nw231;
          let _1333_v87;
          _1333_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1332_v86);
          let _1334_v88;
          _1334_v88 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1332_v86);
          let _1335_v89;
          _1335_v89 = _dafny.MultiSet.fromElements((((_1333_v87).contains((_dafny.ZERO).minus((_this).f14))) ? ((_1333_v87).get((_dafny.ZERO).minus((_this).f14))) : ((((_1334_v88).contains(_this.f7)) ? ((_1334_v88).get(_this.f7)) : (_1332_v86)))));
          let _1336_v90;
          let _nw232 = Array((new BigNumber(2)).toNumber()).fill(false);
          _1336_v90 = _nw232;
          let _1337_v91;
          _1337_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1329_v83,_1336_v90);
          let _1338_v92;
          _1338_v92 = _module.D0.create_DC0((_this).f14);
          let _1339_v93;
          _1339_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1332_v86,(_module.D2.create_DC9(new BigNumber(-590), (_this).f14, _1328_v82, _this.f7, _dafny.Map.Empty.slice().updateUnsafe(_1338_v92,(_this).f14))).dtor_cf16);
          let _1340_v94;
          _1340_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1338_v92,(_dafny.ZERO).minus(new BigNumber((_1219_v1).length)));
          let _1341_v95;
          _1341_v95 = _module.D2.create_DC9((_this).f14, new BigNumber((_1234_v12).length), _1328_v82, _this.f7, _1340_v94);
          let _1342_v96;
          _1342_v96 = _dafny.Seq.of(function (_pat_let39_0) {
            return function (_1343_dt__update__tmp_h3) {
              return function (_pat_let40_0) {
                return function (_1344_dt__update_hcf17_h0) {
                  return function (_pat_let41_0) {
                    return function (_1345_dt__update_hcf16_h0) {
                      return _module.D2.create_DC9(_1345_dt__update_hcf16_h0, _1344_dt__update_hcf17_h0, (_1343_dt__update__tmp_h3).dtor_cf18, (_1343_dt__update__tmp_h3).dtor_cf19, (_1343_dt__update__tmp_h3).dtor_cf20);
                    }(_pat_let41_0);
                  }((_this).f14);
                }(_pat_let40_0);
              }(new BigNumber(-22));
            }(_pat_let39_0);
          }(_1341_v95));
          let _1346_v97;
          let _nw233 = Array((new BigNumber(27)).toNumber());
          _nw233[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1328_v82, _1328_v82)).length);
          _nw233[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1220_v2).length));
          _nw233[(new BigNumber(2)).toNumber()] = (_this).fm9(_this.f7, new BigNumber((_dafny.MultiSet.fromElements(_this.f7)).cardinality()), _1329_v83, globalState);
          _nw233[(new BigNumber(3)).toNumber()] = ((_this).f14).minus(new BigNumber((_1330_v84).cardinality()));
          _nw233[(new BigNumber(4)).toNumber()] = _module.__default.fm0(globalState);
          _nw233[(new BigNumber(5)).toNumber()] = ((_this).f14).minus(new BigNumber(((_1331_v85).dtor_cf48).cardinality()));
          _nw233[(new BigNumber(6)).toNumber()] = (_this).fm9(_this.f7, (_this).f14, _1329_v83, globalState);
          _nw233[(new BigNumber(7)).toNumber()] = new BigNumber((_1335_v89).cardinality());
          _nw233[(new BigNumber(8)).toNumber()] = new BigNumber(638);
          _nw233[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.update(_1328_v82, _module.__default.safeIndex(new BigNumber((_1337_v91).length), new BigNumber((_1328_v82).length)), new _dafny.CodePoint('l'.codePointAt(0)))).length);
          _nw233[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_this.f7)).length);
          _nw233[(new BigNumber(11)).toNumber()] = new BigNumber(483);
          _nw233[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_this).f14);
          _nw233[(new BigNumber(13)).toNumber()] = ((_this).f14).minus(new BigNumber(709));
          _nw233[(new BigNumber(14)).toNumber()] = _module.__default.fm0(globalState);
          _nw233[(new BigNumber(15)).toNumber()] = (((_1339_v93).contains(_1332_v86)) ? ((_1339_v93).get(_1332_v86)) : (new BigNumber(787)));
          _nw233[(new BigNumber(16)).toNumber()] = (_this).f14;
          _nw233[(new BigNumber(17)).toNumber()] = ((_this.f7) ? (new BigNumber(978)) : ((_this).f14));
          _nw233[(new BigNumber(18)).toNumber()] = new BigNumber(-411);
          _nw233[(new BigNumber(19)).toNumber()] = (_this).f14;
          _nw233[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1328_v82, _dafny.Seq.update(_1328_v82, _module.__default.safeIndex(new BigNumber((_1234_v12).length), new BigNumber((_1328_v82).length)), _1329_v83))).length);
          _nw233[(new BigNumber(21)).toNumber()] = (_this).f14;
          _nw233[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1342_v96).length), (_this).f14);
          _nw233[(new BigNumber(23)).toNumber()] = ((_this.f7) ? ((_this).f14) : (new BigNumber((_1277_v50).cardinality())));
          _nw233[(new BigNumber(24)).toNumber()] = (_this).f14;
          _nw233[(new BigNumber(25)).toNumber()] = new BigNumber(-18);
          _nw233[(new BigNumber(26)).toNumber()] = (_this).f14;
          _1346_v97 = _nw233;
          let _index208 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1332_v86).length));
          (_1332_v86)[_index208] = (_this).f14;
          let _index209 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1332_v86).length));
          (_1332_v86)[_index209] = (_dafny.ZERO).minus((_this).f14);
          let _1347_v98;
          _1347_v98 = _module.D0.create_DC1((_this).f14);
          let _pat_let_tv34 = _1329_v83;
          _1347_v98 = function (_pat_let42_0) {
            return function (_1348_dt__update__tmp_h4) {
              return function (_pat_let43_0) {
                return function (_1351_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_1351_dt__update_hcf1_h0);
                }(_pat_let43_0);
              }((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(576)), ((_1349_v83) => function (_1350_i7) {
                return _1349_v83;
              })(_pat_let_tv34))).length)).multipliedBy(new BigNumber(725)));
            }(_pat_let42_0);
          }(_1347_v98);
        } else {
          let _1352___mcc_h6 = (_source9).cf12;
          let _1353_cf12 = _1352___mcc_h6;
          let _1354_v99;
          _1354_v99 = new BigNumber(365);
          _1354_v99 = (_this).f14;
          _1354_v99 = ((_dafny.ZERO).minus(_1354_v99)).minus((_this).f14);
          let _1355_v100;
          let _init47 = function (_1356_i8) {
            return !(true);
          };
          let _nw234 = Array((new BigNumber(25)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw234.length); _i0_47++) {
            _nw234[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1355_v100 = _nw234;
          let _1357_v101;
          _1357_v101 = _dafny.Set.fromElements(_1355_v100);
          let _1358_v102;
          _1358_v102 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1359_v103;
          _1359_v103 = _dafny.Seq.of(_module.D0.create_DC0((_this).f14));
          let _1360_v104;
          let _nw235 = new _module.C2();
          _nw235.__ctor(_1357_v101, _1358_v102, (_module.__default.fm1(_1359_v103, _1354_v99, _1354_v99, globalState)) === (_this.f7));
          _1360_v104 = _nw235;
          let _1361_v105;
          let _nw236 = Array((new BigNumber(2)).toNumber());
          _nw236[(_dafny.ZERO).toNumber()] = _1354_v99;
          _nw236[(_dafny.ONE).toNumber()] = _1354_v99;
          _1361_v105 = _nw236;
          let _1362_v106;
          _1362_v106 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_1219_v1).length));
          let _1363_v107;
          _1363_v107 = _dafny.Map.Empty.slice().updateUnsafe(_1354_v99,_this.f7);
          let _1364_v108;
          _1364_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1362_v106,(_module.__default.fm48((_dafny.ZERO).minus(_1354_v99), _1354_v99, (((_1363_v107).contains(_1354_v99)) ? ((_1363_v107).get(_1354_v99)) : (_this.f7)), _this.f7, globalState)).dtor_cf48);
          let _1365_v109;
          _1365_v109 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm2(new BigNumber(516), true, globalState)).length));
          let _index210 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1361_v105).length));
          (_1361_v105)[_index210] = new BigNumber(((((_1364_v108).contains(_1362_v106)) ? ((_1364_v108).get(_1362_v106)) : (_1365_v109))).cardinality());
          let _index211 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1361_v105).length));
          (_1361_v105)[_index211] = new BigNumber((_dafny.Seq.of(_1361_v105, ((false) ? (_1361_v105) : (_1361_v105)))).length);
        }
      }
      let _1366_v110;
      let _nw237 = new _module.C0();
      _nw237.__ctor(_this.f7, _this.f7);
      _1366_v110 = _nw237;
      let _1367_v111;
      _1367_v111 = _dafny.Seq.of((_this).fm8((_this).f14, globalState));
      let _1368_v112;
      let _nw238 = Array((new BigNumber(20)).toNumber());
      _nw238[(_dafny.ZERO).toNumber()] = (_1366_v110).f12;
      _nw238[(_dafny.ONE).toNumber()] = _this.f7;
      _nw238[(new BigNumber(2)).toNumber()] = _1366_v110.f13;
      _nw238[(new BigNumber(3)).toNumber()] = (_1367_v111)[_module.__default.safeIndex((((_1219_v1).contains((_this).f14)) ? ((_1219_v1).get((_this).f14)) : ((_this).f14)), new BigNumber((_1367_v111).length))];
      _nw238[(new BigNumber(4)).toNumber()] = _this.f7;
      _nw238[(new BigNumber(5)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(6)).toNumber()] = _1366_v110.f13;
      _nw238[(new BigNumber(7)).toNumber()] = !((_1366_v110).fm16(_this.f7, globalState));
      _nw238[(new BigNumber(8)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(9)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(10)).toNumber()] = _1366_v110.f13;
      _nw238[(new BigNumber(11)).toNumber()] = _this.f7;
      _nw238[(new BigNumber(12)).toNumber()] = _this.f7;
      _nw238[(new BigNumber(13)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(14)).toNumber()] = _this.f7;
      _nw238[(new BigNumber(15)).toNumber()] = _this.f7;
      _nw238[(new BigNumber(16)).toNumber()] = !((_1366_v110).f12);
      _nw238[(new BigNumber(17)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(18)).toNumber()] = (_1366_v110).f12;
      _nw238[(new BigNumber(19)).toNumber()] = _this.f7;
      _1368_v112 = _nw238;
      let _1369_v113;
      _1369_v113 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1370_v114;
      _1370_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1368_v112,_1369_v113);
      let _1371_v115;
      _1371_v115 = _dafny.Seq.UnicodeFromString("acjxxafo");
      _1370_v114 = (_1370_v114).update(_1368_v112, (((_module.D1.create_DC4(_1366_v110.f13, (_1367_v111)[_module.__default.safeIndex(new BigNumber((_1371_v115).length), new BigNumber((_1367_v111).length))])).dtor_cf8) ? (_1369_v113) : (_1369_v113)));
      _1369_v113 = _1369_v113;
      let _1372_v116;
      _1372_v116 = _dafny.Set.fromElements(_this.f7);
      let _1373_v117;
      _1373_v117 = _dafny.Set.fromElements(_1372_v116, _1372_v116, _1372_v116, _1372_v116, (_1372_v116).Difference(_1372_v116));
      _1373_v117 = _1373_v117;
      let _1374_v118;
      _1374_v118 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("cejtyhlrx"), _dafny.Seq.UnicodeFromString("vervc"));
      let _1375_v119;
      _1375_v119 = _dafny.MultiSet.fromElements(_this.f7);
      let _1376_v120;
      _1376_v120 = _dafny.Seq.of(_1375_v119, _1375_v119);
      let _1377_v121;
      _1377_v121 = _dafny.MultiSet.fromElements((_this).f14);
      let _1378_v122;
      _1378_v122 = _module.D2.create_DC8(false, _1371_v115);
      let _1379_v123;
      let _nw239 = Array((new BigNumber(28)).toNumber());
      _nw239[(_dafny.ZERO).toNumber()] = (_this).f14;
      _nw239[(_dafny.ONE).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(2)).toNumber()] = (_1220_v2)[_module.__default.safeIndex((_this).f14, new BigNumber((_1220_v2).length))];
      _nw239[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1374_v118).cardinality()),(_this).fm9(_1366_v110.f13, (_this).f14, _1369_v113, globalState)))).length));
      _nw239[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f14);
      _nw239[(new BigNumber(5)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(6)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(7)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(8)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(9)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(10)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(11)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(12)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(13)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(14)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(15)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(16)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(17)).toNumber()] = new BigNumber(84);
      _nw239[(new BigNumber(18)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(19)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(20)).toNumber()] = new BigNumber((_1376_v120).length);
      _nw239[(new BigNumber(21)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(22)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(23)).toNumber()] = (((_1377_v121).contains((_dafny.ZERO).minus((_this).f14))) ? ((_1377_v121).get((_dafny.ZERO).minus((_this).f14))) : ((_this).f14));
      _nw239[(new BigNumber(24)).toNumber()] = _module.__default.fm0(globalState);
      _nw239[(new BigNumber(25)).toNumber()] = new BigNumber((_1372_v116).length);
      _nw239[(new BigNumber(26)).toNumber()] = (_this).f14;
      _nw239[(new BigNumber(27)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1378_v122).dtor_cf15).length));
      _1379_v123 = _nw239;
      let _1380_v124;
      _1380_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v123,_dafny.Set.fromElements((_this).f14, _module.__default.fm0(globalState)));
      let _1381_v125;
      _1381_v125 = _module.D13.create_DC34(_1380_v124);
      (_1366_v110).f13 = !((_1381_v125).dtor_cf57).contains(_1379_v123);
      return;
    }
    m11(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1382_v0;
      _1382_v0 = _module.D1.create_DC4(_this.f7, _this.f7);
      let _1383_i0;
      _1383_i0 = _dafny.ZERO;
      L6: {
        while (function (_source10) {
          if (_source10.is_DC4) {
            let _1387___mcc_h0 = (_source10).cf7;
            let _1388___mcc_h1 = (_source10).cf8;
            let _1389_cf8 = _1388___mcc_h1;
            let _1390_cf7 = _1387___mcc_h0;
            return _1390_cf7;
          } else if (_source10.is_DC5) {
            let _1391___mcc_h2 = (_source10).cf9;
            let _1392___mcc_h3 = (_source10).cf10;
            let _1393___mcc_h4 = (_source10).cf11;
            let _1394_cf11 = _1393___mcc_h4;
            let _1395_cf10 = _1392___mcc_h3;
            let _1396_cf9 = _1391___mcc_h2;
            return _this.f7;
          } else if (_source10.is_DC3) {
            let _1397___mcc_h5 = (_source10).cf6;
            let _1398_cf6 = _1397___mcc_h5;
            return (_this.f7) && (_this.f7);
          } else {
            let _1399___mcc_h6 = (_source10).cf12;
            let _1400_cf12 = _1399___mcc_h6;
            return !(((_this).f14).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(_this.f7, _this.f7, _this.f7)).length)));
          }
        }(_1382_v0)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1383_i0)) {
              break L6;
            }
            _1383_i0 = (_1383_i0).plus(_dafny.ONE);
            let _1384_v1;
            _1384_v1 = new BigNumber(123);
            _1384_v1 = _module.__default.safeModuloInt((_this).f14, _1384_v1);
            let _1385_v2;
            _1385_v2 = _dafny.Seq.UnicodeFromString("kxhikfk");
            _module.__default.m0((_this).f14, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pdhqaeccs"), _1385_v2)).length), globalState);
            let _1386_v3;
            _1386_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v2,_1384_v1);
            _1384_v1 = (((_1386_v3).contains(_1385_v2)) ? ((_1386_v3).get(_1385_v2)) : ((_dafny.ZERO).minus(new BigNumber((_1385_v2).length))));
            _1385_v2 = _dafny.Seq.UnicodeFromString("n");
          }
        }
      }
      let _1401_v4;
      let _init48 = function (_1402_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f14);
      };
      let _nw240 = Array((_dafny.ONE).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw240.length); _i0_48++) {
        _nw240[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1401_v4 = _nw240;
      let _1403_v5;
      _1403_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1401_v4);
      _1403_v5 = (_1403_v5).update(_this.f7, _1401_v4);
      let _1404_v6;
      let _nw241 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1404_v6 = _nw241;
      let _1405_v7;
      _1405_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1404_v6,(_this).f14);
      let _1406_v8;
      _1406_v8 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, true, _this.f7);
      let _1407_v9;
      _1407_v9 = _dafny.Seq.UnicodeFromString("qifiugigg");
      let _1408_v10;
      _1408_v10 = _dafny.Seq.of(_this.f7);
      let _1409_v11;
      _1409_v11 = _module.D12.create_DC30((_this).f14, !(_this.f7), _dafny.Seq.of((_this).f14), new BigNumber((_1408_v10).length));
      let _1410_v12;
      let _nw242 = Array((new BigNumber(23)).toNumber());
      _nw242[(_dafny.ZERO).toNumber()] = (_this).f14;
      _nw242[(_dafny.ONE).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(2)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(3)).toNumber()] = ((_this.f7) ? ((_this).f14) : ((_this).f14));
      _nw242[(new BigNumber(4)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(5)).toNumber()] = (new BigNumber((_1405_v7).length)).plus(new BigNumber((_module.__default.fm21(_this.f7, globalState)).length));
      _nw242[(new BigNumber(6)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(7)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(8)).toNumber()] = _module.__default.fm0(globalState);
      _nw242[(new BigNumber(9)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(10)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(11)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((((_1406_v8).contains(_this.f7)) ? ((_1406_v8).get(_this.f7)) : ((_this).f14)), (_this).f14);
      _nw242[(new BigNumber(13)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(14)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(15)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(16)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(17)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(18)).toNumber()] = new BigNumber((_1407_v9).length);
      _nw242[(new BigNumber(19)).toNumber()] = (_1409_v11).dtor_cf52;
      _nw242[(new BigNumber(20)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(21)).toNumber()] = (_this).f14;
      _nw242[(new BigNumber(22)).toNumber()] = (_this).f14;
      _1410_v12 = _nw242;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1404_v6).length))) {
        let _1411_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1411_i2)) && ((_1411_i2).isLessThan(new BigNumber((_1404_v6).length))))) {
          (_1404_v6)[(_1411_i2)] = _module.__default.safeModuloInt(_1411_i2, (_this).f14);
        }
      }
      let _1412_v13;
      _1412_v13 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f14));
      let _1413_v14;
      _1413_v14 = _dafny.MultiSet.fromElements(new BigNumber((_1412_v13).cardinality()));
      let _1414_i3;
      _1414_i3 = _dafny.ZERO;
      L7: {
        while ((_1412_v13).IsSubsetOf(_1412_v13)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1414_i3)) {
              break L7;
            }
            _1414_i3 = (_1414_i3).plus(_dafny.ONE);
            let _1415_v15;
            _1415_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f7)).cardinality()),_this.f7);
            let _1416_v16;
            _1416_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(((_1415_v15).contains((_dafny.ZERO).minus((_this).f14))) ? ((_1415_v15).get((_dafny.ZERO).minus((_this).f14))) : (_this.f7)));
            let _1417_v17;
            _1417_v17 = _dafny.Seq.of(_module.D0.create_DC0(new BigNumber(308)));
            _1416_v16 = (_1416_v16).update(_this.f7, _module.__default.fm1(_1417_v17, (_this).f14, new BigNumber((_1415_v15).length), globalState));
            let _1418_v18;
            let _init49 = ((_1419_v10, _1420_v8) => function (_1421_i4) {
              return (_1419_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1420_v8).cardinality()))).length), new BigNumber((_1419_v10).length))];
            })(_1408_v10, _1406_v8);
            let _nw243 = Array((new BigNumber(2)).toNumber());
            for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw243.length); _i0_49++) {
              _nw243[_i0_49] = _init49(new BigNumber(_i0_49));
            }
            _1418_v18 = _nw243;
            let _index212 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1418_v18).length));
            (_1418_v18)[_index212] = (new BigNumber((_1408_v10).length)).isLessThan((_this).f14);
            let _index213 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1418_v18).length));
            (_1418_v18)[_index213] = _this.f7;
            _1410_v12 = _1410_v12;
            let _index214 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1418_v18).length));
            (_1418_v18)[_index214] = ((_1406_v8).Union(_1406_v8)).IsProperSubsetOf((_1406_v8).Intersect(_dafny.MultiSet.fromElements((_1418_v18)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_1418_v18).length))], true, true)));
          }
        }
      }
      let _1422_v19;
      _1422_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1408_v10);
      let _1423_v20;
      _1423_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(false, _1422_v19, _this.f7, globalState),_this.f7);
      let _1424_v21;
      _1424_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
      let _1425_v22;
      _1425_v22 = _dafny.Seq.of(_1424_v21);
      let _1426_v23;
      _1426_v23 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1425_v22)[_module.__default.safeIndex((_this).f14, new BigNumber((_1425_v22).length))]);
      let _1427_v24;
      _1427_v24 = _dafny.Seq.of(new BigNumber((_1423_v20).length), new BigNumber((_1407_v9).length), new BigNumber(((((_1426_v23).contains(_this.f7)) ? ((_1426_v23).get(_this.f7)) : ((_1424_v21).update(_this.f7, _this.f7)))).length), (_this).f14, ((_this).f14).plus(new BigNumber(573)));
      let _1428_v25;
      let _nw244 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1428_v25 = _nw244;
      let _index215 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1428_v25).length));
      (_1428_v25)[_index215] = !(_this.f7);
      let _1429_v26;
      _1429_v26 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1430_v27;
      _1430_v27 = _dafny.Set.fromElements((_this).f14);
      let _index216 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1428_v25).length));
      let _rhs213 = (!(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f14)).contains(_this.f7)) === (_this.f7);
      let _rhs214 = _1427_v24;
      let _rhs215 = (((_this.f7) ? (_1430_v27) : (_1430_v27))).IsSubsetOf(_1430_v27);
      let _rhs216 = _1429_v26;
      let _rhs217 = _this.f7;
      let _lhs128 = _1428_v25;
      let _lhs129 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1428_v25).length));
      let _lhs130 = _this;
      r1 = _rhs213;
      _1427_v24 = _rhs214;
      _lhs128[_lhs129] = _rhs215;
      _1429_v26 = _rhs216;
      _lhs130.f7 = _rhs217;
      let _1431_v28;
      let _nw245 = Array((new BigNumber(14)).toNumber()).fill([]);
      _1431_v28 = _nw245;
      let _rhs218 = _1431_v28;
      let _rhs219 = ((!((((_1428_v25)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1428_v25).length))]) ? (!(_this.f7)) : (_this.f7)))) ? ((_1408_v10)[_module.__default.safeIndex(new BigNumber((_1407_v9).length), new BigNumber((_1408_v10).length))]) : (_this.f7));
      _1431_v28 = _rhs218;
      r0 = _rhs219;
      let _1432_v29;
      _1432_v29 = _module.D6.create_DC15(_1429_v26);
      let _pat_let_tv35 = _1429_v26;
      let _1433_v30;
      _1433_v30 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let44_0) {
        return function (_1434_dt__update__tmp_h0) {
          return function (_pat_let45_0) {
            return function (_1435_dt__update_hcf29_h0) {
              return _module.D6.create_DC15(_1435_dt__update_hcf29_h0);
            }(_pat_let45_0);
          }(_pat_let_tv35);
        }(_pat_let44_0);
      }(_1432_v29),(_1428_v25)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1428_v25).length))]);
      r0 = !(_1433_v30).contains(_1432_v29);
      let _1436_v31;
      _1436_v31 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-313)), ((_1437_v26) => function (_1438_i5) {
        return _1437_v26;
      })(_1429_v26)));
      r1 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.update(_1436_v31, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), ((_1439_v26) => function (_1440_i6) {
        return _1439_v26;
      })(_1429_v26))).length), new BigNumber((_1436_v31).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-208)), ((_1441_v26) => function (_1442_i7) {
        return _1441_v26;
      })(_1429_v26))), _1436_v31), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), ((_1443_v9) => function (_1444_i8) {
        return _1443_v9;
      })(_1407_v9)), _1436_v31));
      return [r0, r1];
    }
    m12(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = false;
      let r3 = _dafny.MultiSet.Empty;
      let _1445_v0;
      _1445_v0 = new _dafny.CodePoint('p'.codePointAt(0));
      let _1446_v1;
      _1446_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), function (_1447_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _module.__default.safeIndex((_this).f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), function (_1448_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length)), _1445_v0));
      let _1449_v2;
      _1449_v2 = _dafny.Seq.UnicodeFromString("vau");
      _1446_v1 = (_1446_v1).update(true, _dafny.Seq.Concat(_module.__default.fm35((_this).f14, (_dafny.ZERO).minus((_this).f14), globalState), _1449_v2));
      let _1450_v3;
      _1450_v3 = _dafny.Seq.of(_this.f7);
      let _1451_v4;
      _1451_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_1450_v3).length));
      let _1452_v5;
      _1452_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1451_v4).Merge(_1451_v4)).length),_this.f7);
      _1452_v5 = (_1452_v5).update((_dafny.ZERO).minus((_this).f14), _this.f7);
      _module.__default.m0((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f14, (_this).f14)), (_this).f14, globalState);
      r0 = _this.f7;
      let _hi4 = (_this).f14;
      for (let _1453_i1 = _module.__default.fm0(globalState); _1453_i1.isLessThan(_hi4); _1453_i1 = _1453_i1.plus(_dafny.ONE)) {
        let _1454_v6;
        _1454_v6 = _module.D12.create_DC32(_this.f7);
        let _1455_v7;
        _1455_v7 = _dafny.Set.fromElements(_1449_v2);
        let _1456_v8;
        _1456_v8 = _module.D0.create_DC0((_this).f14);
        let _1457_v9;
        _1457_v9 = _dafny.Seq.of(_1456_v8);
        let _1458_v10;
        _1458_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1449_v2,_this.f7);
        let _1459_v11;
        let _init50 = ((_1460_i1) => function (_1461_i2) {
          return (_1461_i2).plus(_1460_i1);
        })(_1453_i1);
        let _nw246 = Array((new BigNumber(2)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw246.length); _i0_50++) {
          _nw246[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1459_v11 = _nw246;
        let _1462_v12;
        _1462_v12 = _dafny.MultiSet.fromElements(_1459_v11);
        let _1463_v13;
        _1463_v13 = _dafny.Set.fromElements(_1445_v0);
        let _1464_v14;
        _1464_v14 = _dafny.Seq.of((_this).f14, (_this).f14, (_this).f14);
        let _1465_v15;
        _1465_v15 = _dafny.Seq.of(new BigNumber((_1464_v14).length));
        let _1466_v16;
        let _nw247 = Array((new BigNumber(23)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = _this.f7;
        _nw247[(_dafny.ONE).toNumber()] = _this.f7;
        _nw247[(new BigNumber(2)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(3)).toNumber()] = (_this.f7) || (_this.f7);
        _nw247[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_1457_v9, _1453_i1, new BigNumber(503), globalState);
        _nw247[(new BigNumber(5)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(6)).toNumber()] = !((_1450_v3)[_module.__default.safeIndex((_this).f14, new BigNumber((_1450_v3).length))]);
        _nw247[(new BigNumber(7)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(8)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(9)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(10)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(11)).toNumber()] = !(_this.f7) || (_this.f7);
        _nw247[(new BigNumber(12)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(13)).toNumber()] = (((_1458_v10).contains(_1449_v2)) ? ((_1458_v10).get(_1449_v2)) : (_this.f7));
        _nw247[(new BigNumber(14)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(15)).toNumber()] = true;
        _nw247[(new BigNumber(16)).toNumber()] = !(_this.f7);
        _nw247[(new BigNumber(17)).toNumber()] = ((_1462_v12).update(_1459_v11, _module.__default.abs(_1453_i1))).IsProperSubsetOf(_1462_v12);
        _nw247[(new BigNumber(18)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(19)).toNumber()] = _module.__default.fm1(_1457_v9, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm50(_1463_v13, _1455_v7, (_dafny.ZERO).minus(_module.__default.fm0(globalState)), globalState), _module.__default.safeIndex((_this).f14, new BigNumber((_module.__default.fm50(_1463_v13, _1455_v7, (_dafny.ZERO).minus(_module.__default.fm0(globalState)), globalState)).length)), _this.f7), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f14), new BigNumber((_dafny.Seq.update(_module.__default.fm50(_1463_v13, _1455_v7, (_dafny.ZERO).minus(_module.__default.fm0(globalState)), globalState), _module.__default.safeIndex((_this).f14, new BigNumber((_module.__default.fm50(_1463_v13, _1455_v7, (_dafny.ZERO).minus(_module.__default.fm0(globalState)), globalState)).length)), _this.f7)).length)), _this.f7)).length), new BigNumber(295), globalState);
        _nw247[(new BigNumber(20)).toNumber()] = _this.f7;
        _nw247[(new BigNumber(21)).toNumber()] = false;
        _nw247[(new BigNumber(22)).toNumber()] = (_1453_i1).isEqualTo(new BigNumber((_1465_v15).length));
        _1466_v16 = _nw247;
        let _rhs220 = _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_1450_v3, _1450_v3));
        let _rhs221 = _1454_v6;
        let _rhs222 = _1455_v7;
        let _rhs223 = _1466_v16;
        _1450_v3 = _rhs220;
        _1454_v6 = _rhs221;
        _1455_v7 = _rhs222;
        _1466_v16 = _rhs223;
        let _index217 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1466_v16).length));
        (_1466_v16)[_index217] = _this.f7;
        let _1467_v17;
        _1467_v17 = _dafny.Seq.of(_1449_v2, _dafny.Seq.of(_1445_v0), _1449_v2);
        let _index218 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1466_v16).length));
        (_1466_v16)[_index218] = _module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_1468_v8) => function (_1469_i3) {
          return _1468_v8;
        })(_1456_v8)), _module.__default.safeDivisionInt(_1453_i1, new BigNumber((_1467_v17).length)), _1453_i1, globalState);
        let _1470_v18;
        _1470_v18 = _dafny.MultiSet.fromElements((_1467_v17)[_module.__default.safeIndex((_this).f14, new BigNumber((_1467_v17).length))], _1449_v2);
        let _1471_v19;
        _1471_v19 = _dafny.Seq.of(_1470_v18);
        r2 = !((_dafny.MultiSet.FromArray(_1467_v17)).equals((_1471_v19)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f14), new BigNumber((_1471_v19).length))])) || (_this.f7);
        _1449_v2 = _1449_v2;
      }
      let _1472_v20;
      let _nw248 = Array((new BigNumber(8)).toNumber());
      _nw248[(_dafny.ZERO).toNumber()] = new BigNumber((_1450_v3).length);
      _nw248[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.of(true, !(_this.f7))).length);
      _nw248[(new BigNumber(2)).toNumber()] = (_this).f14;
      _nw248[(new BigNumber(3)).toNumber()] = (_this).f14;
      _nw248[(new BigNumber(4)).toNumber()] = (_this).f14;
      _nw248[(new BigNumber(5)).toNumber()] = (_this).f14;
      _nw248[(new BigNumber(6)).toNumber()] = (_this).f14;
      _nw248[(new BigNumber(7)).toNumber()] = (_this).f14;
      _1472_v20 = _nw248;
      let _1473_v21;
      _1473_v21 = _module.D9.create_DC23(_1472_v20);
      let _1474_v22;
      _1474_v22 = _dafny.Seq.of((_1473_v21).dtor_cf44);
      let _1475_v23;
      _1475_v23 = _dafny.Set.fromElements(((_this).f14).minus((_this).f14), new BigNumber((_1474_v22).length), (_this).f14);
      let _1476_v24;
      _1476_v24 = _module.D2.create_DC7(_1475_v23);
      let _pat_let_tv36 = _1475_v23;
      _1475_v23 = (function (_pat_let46_0) {
        return function (_1477_dt__update__tmp_h0) {
          return function (_pat_let47_0) {
            return function (_1478_dt__update_hcf13_h0) {
              return _module.D2.create_DC7(_1478_dt__update_hcf13_h0);
            }(_pat_let47_0);
          }(_pat_let_tv36);
        }(_pat_let46_0);
      }(_1476_v24)).dtor_cf13;
      r0 = _this.f7;
      let _1479_v25;
      _1479_v25 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
      r1 = (_dafny.MultiSet.fromElements(!(_this.f7))).Intersect(_1479_v25);
      r2 = !(true);
      r3 = _1479_v25;
      return [r0, r1, r2, r3];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f7 = false;
      this._f11 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f11, f7) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f7 = f7;
      return;
    }
    fm15(globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), function (_1480_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length)).plus((new BigNumber(-417)).multipliedBy(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gmpsycksp")).length)), new BigNumber((_dafny.Seq.UnicodeFromString("x")).length))).length)));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      let _1481_v0;
      _1481_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1482_v1;
      _1482_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1481_v0).length),_this.f7);
      let _1483_v2;
      let _nw249 = new _module.C0();
      _nw249.__ctor((((_1482_v1).contains((_dafny.ZERO).minus(p1))) ? ((_1482_v1).get((_dafny.ZERO).minus(p1))) : (p0)), p2);
      _1483_v2 = _nw249;
      let _1484_i0;
      _1484_i0 = _dafny.ZERO;
      L8: {
        while ((_1483_v2).f12) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1484_i0)) {
              break L8;
            }
            _1484_i0 = (_1484_i0).plus(_dafny.ONE);
            let _1485_v3;
            _1485_v3 = _dafny.Seq.of(p0, _1483_v2.f13, _this.f7);
            let _1486_v4;
            let _nw250 = new _module.C0();
            _nw250.__ctor(!(new BigNumber((_1485_v3).length)).isEqualTo(_module.__default.fm0(globalState)), p0);
            _1486_v4 = _nw250;
            let _1487_v5;
            _1487_v5 = new BigNumber(503);
            _1487_v5 = _1487_v5;
            let _1488_v6;
            _1488_v6 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_1483_v2.f13)).length));
            let _1489_v7;
            _1489_v7 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.Map.Empty.slice().updateUnsafe(false,_1488_v6)).contains(_1483_v2.f13),_1485_v3);
            _1489_v7 = (_1489_v7).update(!(!((_1483_v2).f12)), _1485_v3);
            _module.__default.m0(p1, new BigNumber(10), globalState);
          }
        }
      }
      let _hi5 = p1;
      for (let _1490_i1 = p1; _1490_i1.isLessThan(_hi5); _1490_i1 = _1490_i1.plus(_dafny.ONE)) {
        let _1491_v9;
        _1491_v9 = _dafny.MultiSet.fromElements(_1483_v2.f13);
        let _1492_v10;
        _1492_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1490_i1);
        let _1493_v11;
        _1493_v11 = _dafny.Set.fromElements(new BigNumber((_1492_v10).length), new BigNumber((_1492_v10).length));
        let _1494_v12;
        _1494_v12 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1493_v11).length),new BigNumber(542))).length));
        let _1495_v13;
        _1495_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1494_v12,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), function (_1496_i2) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length)));
        let _1497_v14;
        _1497_v14 = _dafny.Seq.of(p1, _1490_i1, new BigNumber((_1495_v13).length), _module.__default.fm0(globalState));
        let _1498_v15;
        _1498_v15 = _module.D1.create_DC5(true, _1494_v12, p1);
        let _1499_v16;
        _1499_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1491_v9,_1498_v15);
        let _1500_v17;
        let _nw251 = Array((new BigNumber(12)).toNumber());
        _nw251[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw251[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of (_1499_v16).Keys.Elements) {
            let _1501_v8 = _compr_31;
            if ((_1499_v16).contains(_1501_v8)) {
              _coll31.push([_1501_v8,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1490_i1)).length), new BigNumber(-527), p1)).length)]);
            }
          }
          return _coll31;
        }()).length));
        _nw251[(new BigNumber(2)).toNumber()] = p1;
        _nw251[(new BigNumber(3)).toNumber()] = p1;
        _nw251[(new BigNumber(4)).toNumber()] = _1490_i1;
        _nw251[(new BigNumber(5)).toNumber()] = p1;
        _nw251[(new BigNumber(6)).toNumber()] = p1;
        _nw251[(new BigNumber(7)).toNumber()] = _1490_i1;
        _nw251[(new BigNumber(8)).toNumber()] = (p1).minus(new BigNumber(669));
        _nw251[(new BigNumber(9)).toNumber()] = _1490_i1;
        _nw251[(new BigNumber(10)).toNumber()] = _1490_i1;
        _nw251[(new BigNumber(11)).toNumber()] = p1;
        _1500_v17 = _nw251;
        let _index219 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        (_1500_v17)[_index219] = _1490_i1;
        let _1502_v18;
        _1502_v18 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1490_i1), p1, new BigNumber(146));
        let _1503_v19;
        _1503_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1494_v12);
        let _index220 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        (_1500_v17)[_index220] = (new BigNumber(324)).multipliedBy(((((_1502_v18).contains(new BigNumber(-348))) ? ((_1502_v18).get(new BigNumber(-348))) : (new BigNumber(((((_1503_v19).contains((_1483_v2).f12)) ? ((_1503_v19).get((_1483_v2).f12)) : (_1497_v14))).length)))).multipliedBy(_1490_i1));
        let _1504_v20;
        _1504_v20 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1505_v21;
        _1505_v21 = _dafny.Seq.UnicodeFromString("anfua");
        let _1506_v22;
        let _nw252 = new _module.C0();
        _nw252.__ctor((_1493_v11).equals(_1493_v11), _dafny.Seq.contains(_1505_v21, _1504_v20));
        _1506_v22 = _nw252;
        let _1507_v23;
        _1507_v23 = _module.D0.create_DC1((_dafny.ZERO).minus(_1490_i1));
        let _1508_v24;
        _1508_v24 = _dafny.Set.fromElements(!(false), _1483_v2.f13, p0, _1506_v22.f13, _this.f7);
        let _1509_v25;
        let _nw253 = Array((new BigNumber(25)).toNumber());
        _nw253[(_dafny.ZERO).toNumber()] = _1507_v23;
        _nw253[(_dafny.ONE).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(2)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(3)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(4)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(5)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(6)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(7)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(8)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(9)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(10)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(11)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(12)).toNumber()] = _module.D0.create_DC1(p1);
        _nw253[(new BigNumber(13)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(14)).toNumber()] = _module.D0.create_DC1((_dafny.ZERO).minus(p1));
        _nw253[(new BigNumber(15)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(16)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(17)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(18)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(19)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(20)).toNumber()] = _module.D0.create_DC1(new BigNumber((_1508_v24).length));
        _nw253[(new BigNumber(21)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(22)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(23)).toNumber()] = _1507_v23;
        _nw253[(new BigNumber(24)).toNumber()] = _1507_v23;
        _1509_v25 = _nw253;
        let _index221 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1509_v25).length));
        (_1509_v25)[_index221] = _1507_v23;
        let _1510_v26;
        let _nw254 = Array((new BigNumber(9)).toNumber()).fill(false);
        _1510_v26 = _nw254;
        let _1511_v27;
        _1511_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1506_v22.f13,_1510_v26);
        let _index222 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _index223 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _index224 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1509_v25).length));
        let _index225 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _rhs224 = ((_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))]).plus(new BigNumber(((_1511_v27).Merge(_1511_v27)).length));
        let _rhs225 = (_this).fm15(globalState);
        let _rhs226 = _1507_v23;
        let _rhs227 = (_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))];
        let _rhs228 = (_1483_v2).f12;
        let _lhs131 = _1500_v17;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _lhs133 = _1500_v17;
        let _lhs134 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _lhs135 = _1509_v25;
        let _lhs136 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_1509_v25).length));
        let _lhs137 = _1500_v17;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
        let _lhs139 = _1483_v2;
        _lhs131[_lhs132] = _rhs224;
        _lhs133[_lhs134] = _rhs225;
        _lhs135[_lhs136] = _rhs226;
        _lhs137[_lhs138] = _rhs227;
        _lhs139.f13 = _rhs228;
        if (!(p1).isEqualTo(new BigNumber(423))) {
          let _1512_v28;
          let _nw255 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _1512_v28 = _nw255;
          let _index226 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1512_v28).length));
          (_1512_v28)[_index226] = _1503_v19;
          let _1513_v29;
          _1513_v29 = _dafny.Seq.of(_1503_v19, _1503_v19, _dafny.Map.Empty.slice().updateUnsafe((_1483_v2).f12,_1494_v12));
          let _index227 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1512_v28).length));
          (_1512_v28)[_index227] = (_1513_v29)[_module.__default.safeIndex((_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))], new BigNumber((_1513_v29).length))];
          let _1514_v30;
          _1514_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1493_v11);
          _1514_v30 = (_1514_v30).update(p1, _module.__default.fm2((_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))], !(_1506_v22.f13), globalState));
          let _1515_v31;
          _1515_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1481_v0).Merge(_1481_v0),(_1506_v22).f12);
          _1515_v31 = (_1515_v31).update(_1481_v0, p0);
          (_1506_v22).f13 = !(_1506_v22.f13) || (p0);
          let _index228 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length));
          (_1500_v17)[_index228] = ((_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))]).multipliedBy(_1490_i1);
        } else {
          _module.__default.m0((((_1502_v18).contains(new BigNumber(753))) ? ((_1502_v18).get(new BigNumber(753))) : (p1)), (((_1492_v10).contains(_1483_v2.f13)) ? ((_1492_v10).get(_1483_v2.f13)) : (_1490_i1)), globalState);
          let _1516_v32;
          _1516_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1507_v23,(_1490_i1).plus(p1));
          _1516_v32 = _1516_v32;
          let _1517_v33;
          let _nw256 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _1517_v33 = _nw256;
          _1517_v33 = _1517_v33;
          let _index229 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_1510_v26).length));
          (_1510_v26)[_index229] = (_1483_v2).f12;
          let _index230 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_1510_v26).length));
          (_1510_v26)[_index230] = (((_1502_v18).contains((_1500_v17)[_module.__default.safeIndex(new BigNumber(560), new BigNumber((_1500_v17).length))])) ? ((_1506_v22).f12) : ((((_1506_v22).f12) ? ((_1483_v2).f12) : (_1483_v2.f13))));
          let _1518_v34;
          _1518_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1491_v9,new BigNumber(822));
          _1518_v34 = _1518_v34;
        }
      }
      let _1519_v36;
      _1519_v36 = _dafny.Seq.of((_1483_v2).f12, p2, false, _1483_v2.f13, (_1483_v2).f12);
      let _1520_v37;
      _1520_v37 = _dafny.MultiSet.fromElements(_1519_v36);
      let _1521_v38;
      _1521_v38 = new _dafny.CodePoint('x'.codePointAt(0));
      let _1522_v39;
      _1522_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll32 = new _dafny.Map();
        for (const _compr_32 of (_1520_v37).Elements) {
          let _1523_v35 = _compr_32;
          if ((_1520_v37).contains(_1523_v35)) {
            _coll32.push([_1523_v35,p2]);
          }
        }
        return _coll32;
      }()).length),_1521_v38);
      _1522_v39 = (_1522_v39).update((_dafny.ZERO).minus(p1), _1521_v38);
      if ((_1483_v2.f13) || (p2)) {
        let _1524_v40;
        _1524_v40 = new BigNumber(128);
        _1524_v40 = p1;
        let _1525_v41;
        _1525_v41 = _dafny.Set.fromElements(p2, _1483_v2.f13, _1483_v2.f13);
        let _1526_v42;
        _1526_v42 = _module.D1.create_DC3(_1525_v41);
        let _1527_v43;
        _1527_v43 = _module.D1.create_DC6(_1526_v42);
        let _source11 = _1527_v43;
        if (_source11.is_DC4) {
          let _1528___mcc_h0 = (_source11).cf7;
          let _1529___mcc_h1 = (_source11).cf8;
          let _1530_cf8 = _1529___mcc_h1;
          let _1531_cf7 = _1528___mcc_h0;
          let _1532_v44;
          _1532_v44 = _dafny.Seq.of(p1);
          let _1533_v45;
          _1533_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_1532_v44, _1532_v44),_dafny.Map.Empty.slice().updateUnsafe((_1483_v2).f12,p1));
          let _1534_v46;
          _1534_v46 = _dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC5(p2, _dafny.Seq.of(new BigNumber((_1525_v41).length)), _1524_v40)).dtor_cf9,_1524_v40);
          _1533_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1534_v46);
          _1524_v40 = new BigNumber(-24);
          let _1535_v47;
          _1535_v47 = _dafny.MultiSet.fromElements(_1521_v38, _module.__default.fm18(globalState), _1521_v38);
          let _1536_v48;
          _1536_v48 = _dafny.Map.Empty.slice().updateUnsafe(!(_1525_v41).equals(_1525_v41),((_1535_v47).update(_1521_v38, _module.__default.abs(_1524_v40))).Union(_1535_v47));
          _1535_v47 = (((_1536_v48).contains(_1483_v2.f13)) ? ((_1536_v48).get(_1483_v2.f13)) : (_dafny.MultiSet.fromElements(_1521_v38)));
          let _1537_v49;
          let _nw257 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1537_v49 = _nw257;
          let _1538_v50;
          _1538_v50 = _dafny.Seq.of(_1537_v49);
          let _1539_v51;
          _1539_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v40,_1538_v50);
          _1539_v51 = (_1539_v51).update(_1524_v40, _dafny.Seq.Concat(_1538_v50, _dafny.Seq.of(_1537_v49, _1537_v49, _1537_v49, _1537_v49)));
        } else if (_source11.is_DC5) {
          let _1540___mcc_h2 = (_source11).cf9;
          let _1541___mcc_h3 = (_source11).cf10;
          let _1542___mcc_h4 = (_source11).cf11;
          let _1543_cf11 = _1542___mcc_h4;
          let _1544_cf10 = _1541___mcc_h3;
          let _1545_cf9 = _1540___mcc_h2;
          let _1546_v52;
          _1546_v52 = _dafny.Seq.UnicodeFromString("vu");
          _1546_v52 = _dafny.Seq.update(_1546_v52, _module.__default.safeIndex(_1524_v40, new BigNumber((_1546_v52).length)), _1521_v38);
          _1545_cf9 = true;
          let _1547_v53;
          _1547_v53 = _dafny.MultiSet.fromElements(_1483_v2);
          let _1548_v54;
          _1548_v54 = _dafny.Seq.of(_1483_v2);
          let _1549_v55;
          _1549_v55 = _dafny.Seq.of((_1548_v54)[_module.__default.safeIndex(_1524_v40, new BigNumber((_1548_v54).length))]);
          let _1550_v56;
          _1550_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1548_v54);
          let _1551_v57;
          let _nw258 = Array((new BigNumber(17)).toNumber());
          _nw258[(_dafny.ZERO).toNumber()] = _1547_v53;
          _nw258[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_1483_v2, _1483_v2)).Difference(_1547_v53);
          _nw258[(new BigNumber(2)).toNumber()] = (_1547_v53).Difference(_1547_v53);
          _nw258[(new BigNumber(3)).toNumber()] = (_1547_v53).Intersect(_1547_v53);
          _nw258[(new BigNumber(4)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_1549_v55);
          _nw258[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(((_1483_v2.f13) ? ((((_1550_v56).contains((_1483_v2).f12)) ? ((_1550_v56).get((_1483_v2).f12)) : (_1549_v55))) : (_dafny.Seq.of(_1483_v2, _1483_v2))));
          _nw258[(new BigNumber(7)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(8)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(9)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(10)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_1549_v55);
          _nw258[(new BigNumber(12)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_1483_v2, _1483_v2);
          _nw258[(new BigNumber(14)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(15)).toNumber()] = _1547_v53;
          _nw258[(new BigNumber(16)).toNumber()] = _1547_v53;
          _1551_v57 = _nw258;
          let _index231 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_1551_v57).length));
          (_1551_v57)[_index231] = _1547_v53;
          let _index232 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_1551_v57).length));
          (_1551_v57)[_index232] = _1547_v53;
          _1524_v40 = _1543_cf11;
        } else if (_source11.is_DC3) {
          let _1552___mcc_h5 = (_source11).cf6;
          let _1553_cf6 = _1552___mcc_h5;
          let _1554_v58;
          let _nw259 = new _module.C0();
          _nw259.__ctor((p0) && ((_1483_v2).f12), (_1483_v2).f12);
          _1554_v58 = _nw259;
          let _1555_v59;
          let _init51 = ((_1556_v2) => function (_1557_i3) {
            return _1556_v2.f13;
          })(_1483_v2);
          let _nw260 = Array((new BigNumber(17)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw260.length); _i0_51++) {
            _nw260[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1555_v59 = _nw260;
          let _1558_v60;
          _1558_v60 = _dafny.Seq.of(_1555_v59, _1555_v59, (((_1483_v2).f12) ? (_1555_v59) : (_1555_v59)), _1555_v59);
          _1555_v59 = (_1558_v60)[_module.__default.safeIndex((_1524_v40).multipliedBy(_1524_v40), new BigNumber((_1558_v60).length))];
          let _1559_v61;
          let _nw261 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1559_v61 = _nw261;
          let _rhs229 = false;
          let _rhs230 = _1559_v61;
          let _rhs231 = _1524_v40;
          let _lhs140 = _1483_v2;
          _lhs140.f13 = _rhs229;
          _1559_v61 = _rhs230;
          _1524_v40 = _rhs231;
          (_this).f7 = _1554_v58.f13;
        } else {
          let _1560___mcc_h6 = (_source11).cf12;
          let _1561_cf12 = _1560___mcc_h6;
          let _1562_v62;
          _1562_v62 = _dafny.Set.fromElements(p1);
          let _1563_v63;
          _1563_v63 = _dafny.Map.Empty.slice().updateUnsafe((_1483_v2).f12,new BigNumber((_1562_v62).length));
          let _1564_v64;
          _1564_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v40,p1);
          let _1565_v65;
          _1565_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1564_v64,(_1483_v2).f12);
          let _1566_v66;
          _1566_v66 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)));
          let _1567_v67;
          _1567_v67 = _dafny.Seq.UnicodeFromString("rfqf");
          let _1568_v68;
          _1568_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1567_v67,(_1483_v2).f12);
          let _1569_v69;
          let _nw262 = Array((new BigNumber(28)).toNumber());
          _nw262[(_dafny.ZERO).toNumber()] = new BigNumber((_1566_v66).cardinality());
          _nw262[(_dafny.ONE).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(2)).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(3)).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(4)).toNumber()] = p1;
          _nw262[(new BigNumber(5)).toNumber()] = p1;
          _nw262[(new BigNumber(6)).toNumber()] = new BigNumber(986);
          _nw262[(new BigNumber(7)).toNumber()] = new BigNumber((_1568_v68).length);
          _nw262[(new BigNumber(8)).toNumber()] = p1;
          _nw262[(new BigNumber(9)).toNumber()] = p1;
          _nw262[(new BigNumber(10)).toNumber()] = p1;
          _nw262[(new BigNumber(11)).toNumber()] = p1;
          _nw262[(new BigNumber(12)).toNumber()] = p1;
          _nw262[(new BigNumber(13)).toNumber()] = new BigNumber((_1567_v67).length);
          _nw262[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw262[(new BigNumber(15)).toNumber()] = p1;
          _nw262[(new BigNumber(16)).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(17)).toNumber()] = p1;
          _nw262[(new BigNumber(18)).toNumber()] = p1;
          _nw262[(new BigNumber(19)).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(20)).toNumber()] = p1;
          _nw262[(new BigNumber(21)).toNumber()] = new BigNumber((_1564_v64).length);
          _nw262[(new BigNumber(22)).toNumber()] = new BigNumber((_1563_v63).length);
          _nw262[(new BigNumber(23)).toNumber()] = p1;
          _nw262[(new BigNumber(24)).toNumber()] = _1524_v40;
          _nw262[(new BigNumber(25)).toNumber()] = p1;
          _nw262[(new BigNumber(26)).toNumber()] = p1;
          _nw262[(new BigNumber(27)).toNumber()] = _1524_v40;
          _1569_v69 = _nw262;
          let _1570_v70;
          _1570_v70 = _dafny.Set.fromElements(_1569_v69);
          let _1571_v71;
          _1571_v71 = _dafny.MultiSet.fromElements((_module.D0.create_DC2(p2, _1524_v40, _1570_v70, p1)).dtor_cf2);
          let _1572_v72;
          _1572_v72 = _dafny.Seq.of(_1524_v40);
          let _1573_v73;
          _1573_v73 = _dafny.Map.Empty.slice().updateUnsafe((((_1565_v65).contains(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1571_v71).cardinality())))) ? ((_1565_v65).get(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1571_v71).cardinality())))) : ((_1483_v2).f12)),_1572_v72);
          let _1574_v74;
          let _nw263 = Array((new BigNumber(10)).toNumber());
          _nw263[(_dafny.ZERO).toNumber()] = (((_1563_v63).contains((_1483_v2).f12)) ? ((_1563_v63).get((_1483_v2).f12)) : (p1));
          _nw263[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_this).fm15(globalState), _1524_v40);
          _nw263[(new BigNumber(2)).toNumber()] = _1524_v40;
          _nw263[(new BigNumber(3)).toNumber()] = _1524_v40;
          _nw263[(new BigNumber(4)).toNumber()] = new BigNumber((_1573_v73).length);
          _nw263[(new BigNumber(5)).toNumber()] = _1524_v40;
          _nw263[(new BigNumber(6)).toNumber()] = p1;
          _nw263[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(701), p1);
          _nw263[(new BigNumber(8)).toNumber()] = new BigNumber(395);
          _nw263[(new BigNumber(9)).toNumber()] = ((_module.D0.create_DC1((_dafny.ZERO).minus(p1))).dtor_cf1).minus(_1524_v40);
          _1574_v74 = _nw263;
          let _index233 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length));
          (_1569_v69)[_index233] = (((_1564_v64).contains(_1524_v40)) ? ((_1564_v64).get(_1524_v40)) : ((_1572_v72)[_module.__default.safeIndex(_1524_v40, new BigNumber((_1572_v72).length))]));
          let _1575_v75;
          let _nw264 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1575_v75 = _nw264;
          let _index234 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length));
          let _rhs232 = _module.__default.safeDivisionInt(p1, _1524_v40);
          let _rhs233 = _1524_v40;
          let _rhs234 = _1575_v75;
          let _rhs235 = _1519_v36;
          let _rhs236 = (_1483_v2).f12;
          let _lhs141 = _1569_v69;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length));
          let _lhs143 = _this;
          _1524_v40 = _rhs232;
          _lhs141[_lhs142] = _rhs233;
          _1575_v75 = _rhs234;
          _1519_v36 = _rhs235;
          _lhs143.f7 = _rhs236;
          let _1576_v76;
          let _nw265 = Array((new BigNumber(4)).toNumber()).fill([]);
          _1576_v76 = _nw265;
          let _1577_v77;
          let _nw266 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1577_v77 = _nw266;
          let _index235 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1576_v76).length));
          (_1576_v76)[_index235] = _1577_v77;
          let _1578_v78;
          let _nw267 = Array((new BigNumber(25)).toNumber()).fill([]);
          _1578_v78 = _nw267;
          let _index236 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length));
          (_1578_v78)[_index236] = _1575_v75;
          let _index237 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1576_v76).length));
          let _index238 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length));
          let _rhs237 = _1577_v77;
          let _rhs238 = false;
          let _rhs239 = p2;
          let _rhs240 = ((true) ? (p0) : ((_1483_v2.f13) || ((_1483_v2).f12)));
          let _rhs241 = _1575_v75;
          let _lhs144 = _1576_v76;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1576_v76).length));
          let _lhs146 = _1483_v2;
          let _lhs147 = _1483_v2;
          let _lhs148 = _1483_v2;
          let _lhs149 = _1578_v78;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length));
          _lhs144[_lhs145] = _rhs237;
          _lhs146.f13 = _rhs238;
          _lhs147.f13 = _rhs239;
          _lhs148.f13 = _rhs240;
          _lhs149[_lhs150] = _rhs241;
          let _1579_v79;
          _1579_v79 = _dafny.Seq.of((_1578_v78)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length))], (_1578_v78)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length))], (_1578_v78)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length))], _1575_v75, (_1578_v78)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length))]);
          let _1580_v80;
          _1580_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1483_v2).f12,(_1579_v79)[_module.__default.safeIndex((_1569_v69)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length))], new BigNumber((_1579_v79).length))]);
          let _rhs242 = _1572_v72;
          let _rhs243 = (((_1580_v80).contains(false)) ? ((_1580_v80).get(false)) : ((_1578_v78)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_1578_v78).length))]));
          _1572_v72 = _rhs242;
          _1575_v75 = _rhs243;
          let _1581_v81;
          _1581_v81 = _module.D0.create_DC1(_module.__default.safeModuloInt(_module.__default.fm0(globalState), _1524_v40));
          let _index239 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length));
          let _rhs244 = p2;
          let _rhs245 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_1567_v67, _1567_v67), _module.__default.safeIndex((_1581_v81).dtor_cf1, new BigNumber((_dafny.Seq.Concat(_1567_v67, _1567_v67)).length)), _1521_v38), _1567_v67);
          let _rhs246 = (_1572_v72)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(301)), function (_1582_i4) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length), new BigNumber((_1572_v72).length))];
          let _rhs247 = _1581_v81;
          let _rhs248 = (_1483_v2).f12;
          let _lhs151 = _this;
          let _lhs152 = _1569_v69;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1569_v69).length));
          let _lhs154 = _1483_v2;
          _lhs151.f7 = _rhs244;
          _1567_v67 = _rhs245;
          _lhs152[_lhs153] = _rhs246;
          _1581_v81 = _rhs247;
          _lhs154.f13 = _rhs248;
        }
        let _1583_v82;
        _1583_v82 = _module.D0.create_DC0(p1);
        _1524_v40 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p1, _1524_v40), (_1583_v82).dtor_cf0);
        let _1584_v83;
        _1584_v83 = _dafny.Seq.UnicodeFromString("btmsw");
        let _1585_v84;
        _1585_v84 = _module.D2.create_DC8((p1).isLessThanOrEqualTo(_1524_v40), _1584_v83);
        let _1586_v85;
        _1586_v85 = _dafny.Seq.of(_1583_v82, _1583_v82);
        let _pat_let_tv37 = _1586_v85;
        let _pat_let_tv38 = p1;
        let _pat_let_tv39 = _1524_v40;
        let _pat_let_tv40 = globalState;
        _1585_v84 = function (_pat_let48_0) {
          return function (_1587_dt__update__tmp_h0) {
            return function (_pat_let49_0) {
              return function (_1588_dt__update_hcf14_h0) {
                return _module.D2.create_DC8(_1588_dt__update_hcf14_h0, (_1587_dt__update__tmp_h0).dtor_cf15);
              }(_pat_let49_0);
            }(_module.__default.fm1(_pat_let_tv37, _pat_let_tv38, _pat_let_tv39, _pat_let_tv40));
          }(_pat_let48_0);
        }(_1585_v84);
        let _1589_v86;
        let _nw268 = Array((new BigNumber(24)).toNumber());
        _nw268[(_dafny.ZERO).toNumber()] = p2;
        _nw268[(_dafny.ONE).toNumber()] = p2;
        _nw268[(new BigNumber(2)).toNumber()] = (_1483_v2).fm16(true, globalState);
        _nw268[(new BigNumber(3)).toNumber()] = p0;
        _nw268[(new BigNumber(4)).toNumber()] = (_1483_v2).f12;
        _nw268[(new BigNumber(5)).toNumber()] = (_1483_v2).f12;
        _nw268[(new BigNumber(6)).toNumber()] = _1483_v2.f13;
        _nw268[(new BigNumber(7)).toNumber()] = (false) === (true);
        _nw268[(new BigNumber(8)).toNumber()] = true;
        _nw268[(new BigNumber(9)).toNumber()] = p0;
        _nw268[(new BigNumber(10)).toNumber()] = (((_1519_v36)[_module.__default.safeIndex(p1, new BigNumber((_1519_v36).length))]) ? ((_1483_v2).f12) : ((_1483_v2).f12));
        _nw268[(new BigNumber(11)).toNumber()] = true;
        _nw268[(new BigNumber(12)).toNumber()] = ((_1483_v2.f13) ? ((_1483_v2).f12) : ((_1483_v2).f12));
        _nw268[(new BigNumber(13)).toNumber()] = _1483_v2.f13;
        _nw268[(new BigNumber(14)).toNumber()] = (_1483_v2).f12;
        _nw268[(new BigNumber(15)).toNumber()] = (_1483_v2).f12;
        _nw268[(new BigNumber(16)).toNumber()] = false;
        _nw268[(new BigNumber(17)).toNumber()] = _this.f7;
        _nw268[(new BigNumber(18)).toNumber()] = ((_this.f7) ? (false) : (p0));
        _nw268[(new BigNumber(19)).toNumber()] = (_1483_v2).f12;
        _nw268[(new BigNumber(20)).toNumber()] = p0;
        _nw268[(new BigNumber(21)).toNumber()] = _1483_v2.f13;
        _nw268[(new BigNumber(22)).toNumber()] = (_1524_v40).isLessThanOrEqualTo(p1);
        _nw268[(new BigNumber(23)).toNumber()] = _1483_v2.f13;
        _1589_v86 = _nw268;
        let _index240 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_1589_v86).length));
        (_1589_v86)[_index240] = _1483_v2.f13;
        let _index241 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_1589_v86).length));
        (_1589_v86)[_index241] = (p0) || (_this.f7);
      } else {
        (_1483_v2).f13 = false;
        (_this).f7 = _1483_v2.f13;
        let _1590_v87;
        _1590_v87 = _module.D1.create_DC5((new BigNumber(35)).isLessThan(p1), _dafny.Seq.of(p1), p1);
        _1590_v87 = _1590_v87;
        let _1591_v88;
        let _nw269 = Array((new BigNumber(21)).toNumber());
        _nw269[(_dafny.ZERO).toNumber()] = (_1483_v2).f12;
        _nw269[(_dafny.ONE).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(2)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(3)).toNumber()] = p0;
        _nw269[(new BigNumber(4)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(5)).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(6)).toNumber()] = true;
        _nw269[(new BigNumber(7)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(8)).toNumber()] = p0;
        _nw269[(new BigNumber(9)).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(10)).toNumber()] = _this.f7;
        _nw269[(new BigNumber(11)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(12)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(13)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(14)).toNumber()] = _1483_v2.f13;
        _nw269[(new BigNumber(15)).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(16)).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(17)).toNumber()] = _this.f7;
        _nw269[(new BigNumber(18)).toNumber()] = true;
        _nw269[(new BigNumber(19)).toNumber()] = (_1483_v2).f12;
        _nw269[(new BigNumber(20)).toNumber()] = false;
        _1591_v88 = _nw269;
        let _1592_v89;
        _1592_v89 = _dafny.Seq.of(_1591_v88);
        if ((_1483_v2).fm16(_dafny.areEqual(_dafny.Seq.of(_1591_v88), _1592_v89), globalState)) {
          let _1593_v90;
          let _init52 = ((_1594_p1) => function (_1595_i5) {
            return (_1595_i5).plus(_1594_p1);
          })(p1);
          let _nw270 = Array((new BigNumber(29)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw270.length); _i0_52++) {
            _nw270[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _1593_v90 = _nw270;
          let _1596_v91;
          _1596_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1521_v38,_1593_v90);
          _1596_v91 = ((true) ? ((_1596_v91).Merge(_1596_v91)) : (((_1596_v91).update(_1521_v38, _1593_v90)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1521_v38,_1593_v90))));
          let _1597_v92;
          let _nw271 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _1597_v92 = _nw271;
          let _1598_v93;
          _1598_v93 = _dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC10(_1591_v88)).dtor_cf21,_1597_v92);
          let _1599_v94;
          _1599_v94 = _1597_v92;
          let _1600_v95;
          _1600_v95 = _dafny.Seq.of(_1597_v92);
          let _1601_v96;
          _1601_v96 = _dafny.MultiSet.fromElements(false);
          let _1602_v97;
          let _nw272 = Array((new BigNumber(26)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = _1597_v92;
          _nw272[(_dafny.ONE).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(2)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(3)).toNumber()] = (((_1598_v93).contains(_1591_v88)) ? ((_1598_v93).get(_1591_v88)) : (_1597_v92));
          _nw272[(new BigNumber(4)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(5)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(6)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(7)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(8)).toNumber()] = (_1599_v94);
          _nw272[(new BigNumber(9)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(10)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(11)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(12)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(13)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(14)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(15)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(16)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(17)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(18)).toNumber()] = (_1600_v95)[_module.__default.safeIndex((((_1601_v96).contains(true)) ? ((_1601_v96).get(true)) : (p1)), new BigNumber((_1600_v95).length))];
          _nw272[(new BigNumber(19)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(20)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(21)).toNumber()] = (_1599_v94);
          _nw272[(new BigNumber(22)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(23)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(24)).toNumber()] = _1597_v92;
          _nw272[(new BigNumber(25)).toNumber()] = _1597_v92;
          _1602_v97 = _nw272;
          let _index242 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1602_v97).length));
          (_1602_v97)[_index242] = _1597_v92;
          let _index243 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1591_v88).length));
          (_1591_v88)[_index243] = _1483_v2.f13;
          let _1603_v98;
          _1603_v98 = _dafny.MultiSet.fromElements(_1593_v90);
          let _index244 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1602_v97).length));
          let _index245 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1591_v88).length));
          let _rhs249 = _1521_v38;
          let _rhs250 = (((_1483_v2).f12) ? (_1597_v92) : (_1597_v92));
          let _rhs251 = (_1483_v2).f12;
          let _rhs252 = _1603_v98;
          let _lhs155 = _1602_v97;
          let _lhs156 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1602_v97).length));
          let _lhs157 = _1591_v88;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1591_v88).length));
          _1521_v38 = _rhs249;
          _lhs155[_lhs156] = _rhs250;
          _lhs157[_lhs158] = _rhs251;
          _1603_v98 = _rhs252;
          let _1604_v99;
          let _nw273 = Array((new BigNumber(17)).toNumber()).fill([]);
          _1604_v99 = _nw273;
          let _1605_v100;
          let _init53 = function (_1606_i6) {
            return _dafny.Seq.UnicodeFromString("m");
          };
          let _nw274 = Array((new BigNumber(22)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw274.length); _i0_53++) {
            _nw274[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _1605_v100 = _nw274;
          let _index246 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1604_v99).length));
          (_1604_v99)[_index246] = _1605_v100;
          let _1607_v101;
          _1607_v101 = _dafny.Seq.of(p1);
          let _1608_v102;
          _1608_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1607_v101,_1605_v100);
          let _index247 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1604_v99).length));
          (_1604_v99)[_index247] = (((_1608_v102).contains(((!(_1483_v2.f13)) ? (_1607_v101) : (_1607_v101)))) ? ((_1608_v102).get(((!(_1483_v2.f13)) ? (_1607_v101) : (_1607_v101)))) : (_1605_v100));
          let _index248 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1591_v88).length));
          (_1591_v88)[_index248] = _this.f7;
          let _1609_v103;
          _1609_v103 = new BigNumber(951);
          _1609_v103 = p1;
        } else {
          (_1483_v2).f13 = ((_this.f7) ? (p2) : ((_1483_v2).f12));
          let _1610_v104;
          _1610_v104 = _dafny.Seq.of(p1, p1);
          let _1611_v105;
          let _nw275 = new _module.C0();
          _nw275.__ctor(!(_1483_v2.f13), (_1519_v36)[_module.__default.safeIndex((_1610_v104)[_module.__default.safeIndex(p1, new BigNumber((_1610_v104).length))], new BigNumber((_1519_v36).length))]);
          _1611_v105 = _nw275;
          let _1612_v106;
          _1612_v106 = _dafny.Seq.UnicodeFromString("ukiih");
          _1612_v106 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tgmrpobis"), _dafny.Seq.Concat(_1612_v106, _1612_v106));
          let _1613_v107;
          _1613_v107 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1610_v104)[_module.__default.safeIndex(p1, new BigNumber((_1610_v104).length))]);
          let _1614_v108;
          _1614_v108 = new BigNumber(410);
          let _1615_v109;
          _1615_v109 = _module.D0.create_DC1(_1614_v108);
          let _pat_let_tv41 = p2;
          let _pat_let_tv42 = p1;
          let _1616_v110;
          _1616_v110 = _dafny.MultiSet.fromElements(_1590_v87, _1590_v87, function (_pat_let50_0) {
            return function (_1617_dt__update__tmp_h1) {
              return function (_pat_let51_0) {
                return function (_1618_dt__update_hcf9_h0) {
                  return function (_pat_let52_0) {
                    return function (_1619_dt__update_hcf11_h0) {
                      return _module.D1.create_DC5(_1618_dt__update_hcf9_h0, (_1617_dt__update__tmp_h1).dtor_cf10, _1619_dt__update_hcf11_h0);
                    }(_pat_let52_0);
                  }((_dafny.ZERO).minus(_pat_let_tv42));
                }(_pat_let51_0);
              }(_pat_let_tv41);
            }(_pat_let50_0);
          }(_1590_v87));
          let _1620_v111;
          _1620_v111 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(globalState),p1);
          let _1621_v112;
          _1621_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1481_v0,_1620_v111);
          let _rhs253 = !(_1621_v112).equals(_1621_v112);
          let _rhs254 = _module.__default.fm19(p1, globalState);
          let _rhs255 = _1614_v108;
          let _rhs256 = _1615_v109;
          let _rhs257 = _1616_v110;
          let _lhs159 = _1483_v2;
          _lhs159.f13 = _rhs253;
          _1613_v107 = _rhs254;
          _1614_v108 = _rhs255;
          _1615_v109 = _rhs256;
          _1616_v110 = _rhs257;
          (_1483_v2).f13 = _this.f7;
        }
        let _1622_v113;
        _1622_v113 = _dafny.Seq.UnicodeFromString("l");
        (_1483_v2).f13 = ((_1483_v2.f13) ? (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("w"), _1622_v113)) : (((false) ? ((_1483_v2).f12) : (_1483_v2.f13))));
      }
      (_this).f7 = (_1483_v2).f12;
      r0 = _1481_v0;
      let _1623_v114;
      _1623_v114 = _dafny.Set.fromElements(!(_this.f7));
      r1 = _1623_v114;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _1624_v0;
      _1624_v0 = new BigNumber(361);
      let _1625_v1;
      _1625_v1 = _dafny.Seq.UnicodeFromString("xl");
      let _1626_v2;
      _1626_v2 = _module.D2.create_DC8(_this.f7, _1625_v1);
      _1624_v0 = new BigNumber((_dafny.Seq.Concat((_1626_v2).dtor_cf15, _1625_v1)).length);
      let _hi6 = _1624_v0;
      for (let _1627_i0 = _1624_v0; _1627_i0.isLessThan(_hi6); _1627_i0 = _1627_i0.plus(_dafny.ONE)) {
        let _1628_v3;
        _1628_v3 = _module.D0.create_DC0(_1627_i0);
        let _1629_v4;
        _1629_v4 = _dafny.Seq.of(_1628_v3);
        let _1630_v5;
        let _nw276 = new _module.C0();
        _nw276.__ctor(_module.__default.fm1(_1629_v4, _1627_i0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_1625_v1, _module.__default.safeIndex(_1624_v0, new BigNumber((_1625_v1).length)), new _dafny.CodePoint('e'.codePointAt(0)))).length)), globalState), _this.f7);
        _1630_v5 = _nw276;
        let _1631_v6;
        _1631_v6 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1632_v7;
        _1632_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1630_v5.f13,_1631_v6);
        let _1633_v8;
        _1633_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1630_v5.f13,_dafny.Seq.update(_dafny.Seq.update((((_1630_v5).f12) ? (_module.__default.fm20((_1630_v5).fm16(_1630_v5.f13, globalState), (_1630_v5).f12, globalState)) : (_dafny.Seq.UnicodeFromString("gcosxqa"))), _module.__default.safeIndex(_1624_v0, new BigNumber(((((_1630_v5).f12) ? (_module.__default.fm20((_1630_v5).fm16(_1630_v5.f13, globalState), (_1630_v5).f12, globalState)) : (_dafny.Seq.UnicodeFromString("gcosxqa")))).length)), _1631_v6), _module.__default.safeIndex(_1624_v0, new BigNumber((_dafny.Seq.update((((_1630_v5).f12) ? (_module.__default.fm20((_1630_v5).fm16(_1630_v5.f13, globalState), (_1630_v5).f12, globalState)) : (_dafny.Seq.UnicodeFromString("gcosxqa"))), _module.__default.safeIndex(_1624_v0, new BigNumber(((((_1630_v5).f12) ? (_module.__default.fm20((_1630_v5).fm16(_1630_v5.f13, globalState), (_1630_v5).f12, globalState)) : (_dafny.Seq.UnicodeFromString("gcosxqa")))).length)), _1631_v6)).length)), (((_1632_v7).contains(_this.f7)) ? ((_1632_v7).get(_this.f7)) : (_1631_v6))));
        _1625_v1 = (((_1633_v8).contains(_this.f7)) ? ((_1633_v8).get(_this.f7)) : (_1625_v1));
        if ((_1630_v5).fm16((_1630_v5).f12, globalState)) {
          let _1634_v9;
          let _nw277 = Array((new BigNumber(5)).toNumber()).fill(_module.D3.Default());
          _1634_v9 = _nw277;
          let _1635_v10;
          _1635_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1631_v6, _1631_v6)).length),_1634_v9);
          _module.__default.m0(new BigNumber((_1635_v10).length), new BigNumber((_1625_v1).length), globalState);
          _1624_v0 = ((_this).fm15(globalState)).plus(_1624_v0);
          (_this).f7 = (_1630_v5).f12;
          _1624_v0 = _module.__default.safeModuloInt(_1627_i0, _1627_i0);
          let _1636_v11;
          _1636_v11 = _dafny.Set.fromElements(_1627_i0);
          _1636_v11 = _module.__default.fm2(_module.__default.fm0(globalState), (((_1630_v5).f12) ? (_1630_v5.f13) : (_1630_v5.f13)), globalState);
        } else {
          let _1637_v12;
          _1637_v12 = _dafny.Seq.of(_1624_v0);
          let _1638_v13;
          _1638_v13 = _dafny.Set.fromElements(_1630_v5.f13);
          let _1639_v14;
          _1639_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1637_v12,new BigNumber((_1638_v13).length));
          _1624_v0 = (new BigNumber(((_1639_v14).Merge(_1639_v14)).length)).multipliedBy(_1627_i0);
          _1625_v1 = _1625_v1;
          let _1640_v15;
          _1640_v15 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-993)), ((_1641_v0) => function (_1642_i1) {
            return _1641_v0;
          })(_1624_v0)));
          let _1643_v16;
          _1643_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1640_v15).Difference(_1640_v15),_1630_v5.f13);
          _1643_v16 = (_1643_v16).update(_1640_v15, _1630_v5.f13);
          let _1644_v17;
          let _nw278 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1644_v17 = _nw278;
          _1644_v17 = _1644_v17;
          _1624_v0 = (_dafny.ZERO).minus(((_1624_v0).multipliedBy(_1627_i0)).multipliedBy(_1624_v0));
        }
        let _1645_v18;
        _1645_v18 = _module.D2.create_DC9(_1624_v0, _1624_v0, _1625_v1, true, _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1624_v0),_1624_v0));
        let _1646_v19;
        _1646_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1624_v0),new BigNumber(459));
        let _1647_v20;
        _1647_v20 = _module.D2.create_DC9((_1645_v18).dtor_cf17, _1627_i0, _1625_v1, false, _1646_v19);
        let _1648_v21;
        _1648_v21 = _dafny.Seq.of((_1645_v18).dtor_cf17, new BigNumber(718));
        _1624_v0 = (new BigNumber((_1648_v21).length)).plus(_1627_i0);
      }
      let _1649_v22;
      let _nw279 = Array((new BigNumber(8)).toNumber()).fill([]);
      _1649_v22 = _nw279;
      let _1650_v23;
      let _nw280 = Array((new BigNumber(3)).toNumber());
      _nw280[(_dafny.ZERO).toNumber()] = _1624_v0;
      _nw280[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1624_v0);
      _nw280[(new BigNumber(2)).toNumber()] = _1624_v0;
      _1650_v23 = _nw280;
      let _1651_v24;
      _1651_v24 = _dafny.Seq.of(_1650_v23);
      let _index249 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_1649_v22).length));
      (_1649_v22)[_index249] = (_1651_v24)[_module.__default.safeIndex(_1624_v0, new BigNumber((_1651_v24).length))];
      let _index250 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_1649_v22).length));
      (_1649_v22)[_index250] = _1650_v23;
      let _1652_v25;
      let _nw281 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
      _1652_v25 = _nw281;
      let _1653_v26;
      _1653_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f7)),(_module.D1.create_DC4(true, _this.f7)).dtor_cf7);
      let _index251 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1652_v25).length));
      (_1652_v25)[_index251] = _1653_v26;
      let _index252 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1652_v25).length));
      (_1652_v25)[_index252] = (_1653_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f7),_module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), ((_1654_v0) => function (_1655_i2) {
        return _module.D0.create_DC0(_1654_v0);
      })(_1624_v0)), _1624_v0, _1624_v0, globalState)));
      let _1656_v27;
      _1656_v27 = _dafny.Set.fromElements((_1649_v22)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_1649_v22).length))]);
      let _1657_v28;
      _1657_v28 = _module.D0.create_DC2(_this.f7, _1624_v0, _1656_v27, _1624_v0);
      let _hi7 = _1624_v0;
      for (let _1658_i3 = (_1657_v28).dtor_cf3; _1658_i3.isLessThan(_hi7); _1658_i3 = _1658_i3.plus(_dafny.ONE)) {
        let _1659_v29;
        _1659_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
        let _1660_v30;
        _1660_v30 = _dafny.Seq.of(_1659_v29, _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7), _1659_v29, _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7));
        let _1661_v31;
        _1661_v31 = _dafny.Seq.of(_1659_v29, (_1660_v30)[_module.__default.safeIndex(_1624_v0, new BigNumber((_1660_v30).length))], _1659_v29);
        let _1662_v32;
        _1662_v32 = _dafny.MultiSet.fromElements(_1658_i3, (_dafny.ZERO).minus(_1624_v0), new BigNumber(((_1661_v31)[_module.__default.safeIndex(_1658_i3, new BigNumber((_1661_v31).length))]).length));
        (_this).f7 = !(!(_1662_v32).contains(_1658_i3));
        _1624_v0 = _1624_v0;
        (_this).f7 = _this.f7;
        (_this).m10(_1624_v0, globalState);
      }
      let _1663_v33;
      let _nw282 = new _module.C0();
      _nw282.__ctor(!(_this.f7), _this.f7);
      _1663_v33 = _nw282;
      let _1664_v34;
      _1664_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1663_v33,_1663_v33.f13);
      let _1665_v35;
      _1665_v35 = _dafny.MultiSet.fromElements(_module.__default.fm0(globalState));
      if ((((_1664_v34).contains(_1663_v33)) ? ((_1664_v34).get(_1663_v33)) : ((_1665_v35).IsSubsetOf(_1665_v35)))) {
        (_1663_v33).f13 = false;
        let _1666_v36;
        let _init54 = function (_1667_i4) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        };
        let _nw283 = Array((new BigNumber(22)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw283.length); _i0_54++) {
          _nw283[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1666_v36 = _nw283;
        let _1668_v37;
        _1668_v37 = new _dafny.CodePoint('h'.codePointAt(0));
        let _index253 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1666_v36).length));
        (_1666_v36)[_index253] = ((_this.f7) ? (_1668_v37) : (_1668_v37));
        let _index254 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1666_v36).length));
        (_1666_v36)[_index254] = _1668_v37;
        (_1663_v33).f13 = !(_1624_v0).isEqualTo(_1624_v0);
        let _1669_v38;
        _1669_v38 = _dafny.Set.fromElements(_1624_v0);
        (_1663_v33).f13 = (_1669_v38).IsSubsetOf(_1669_v38);
        let _1670_v39;
        let _nw284 = Array((new BigNumber(7)).toNumber());
        _1670_v39 = _nw284;
        let _1671_v40;
        let _nw285 = new _module.C3();
        _nw285.__ctor(_1624_v0, !(_1663_v33.f13));
        _1671_v40 = _nw285;
        let _index255 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1670_v39).length));
        (_1670_v39)[_index255] = _1671_v40;
        let _index256 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1670_v39).length));
        (_1670_v39)[_index256] = _1671_v40;
      } else {
        _1624_v0 = _1624_v0;
        let _1672_v41;
        _1672_v41 = _module.D0.create_DC0(new BigNumber(508));
        let _1673_v42;
        _1673_v42 = _dafny.Set.fromElements(_1624_v0);
        let _1674_v43;
        _1674_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1624_v0,_1624_v0);
        let _rhs258 = _dafny.Seq.UnicodeFromString("efqt");
        let _rhs259 = _1625_v1;
        let _rhs260 = ((_1624_v0).minus(_1624_v0)).multipliedBy(_1624_v0);
        let _rhs261 = (_1672_v41).dtor_cf0;
        let _rhs262 = (function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of (_1674_v43).Keys.Elements) {
            let _1675_v44 = _compr_33;
            if ((_1674_v43).contains(_1675_v44)) {
              _coll33.add((_1675_v44).plus((_dafny.ZERO).minus(new BigNumber(-336))));
            }
          }
          return _coll33;
        }()).IsSubsetOf((_1673_v42).Union(_1673_v42));
        let _lhs160 = _this;
        _1625_v1 = _rhs258;
        _1625_v1 = _rhs259;
        _1624_v0 = _rhs260;
        _1624_v0 = _rhs261;
        _lhs160.f7 = _rhs262;
        let _1676_v45;
        _1676_v45 = _dafny.MultiSet.fromElements(false, _this.f7, _1663_v33.f13);
        let _1677_v46;
        _1677_v46 = _dafny.Seq.of(_1676_v45);
        let _1678_v47;
        _1678_v47 = _dafny.Set.fromElements(_1663_v33.f13, false, _this.f7, !(_1663_v33.f13) || (_1663_v33.f13));
        let _rhs263 = _1677_v46;
        let _rhs264 = _module.__default.fm27((_1663_v33).f12, new _dafny.CodePoint('t'.codePointAt(0)), !(!(_1624_v0).isEqualTo(_1624_v0)), _this.f7, globalState);
        let _rhs265 = (_1624_v0).minus(_module.__default.safeDivisionInt(_1624_v0, _1624_v0));
        let _rhs266 = new BigNumber((_1625_v1).length);
        let _rhs267 = ((_this).fm15(globalState)).minus(_1624_v0);
        _1677_v46 = _rhs263;
        _1678_v47 = _rhs264;
        _1624_v0 = _rhs265;
        _1624_v0 = _rhs266;
        _1624_v0 = _rhs267;
        let _1679_v48;
        _1679_v48 = new _dafny.CodePoint('c'.codePointAt(0));
        _1679_v48 = _1679_v48;
        if (!(((_1678_v47).Union(_1678_v47)).IsDisjointFrom(_1678_v47))) {
          let _1680_v49;
          let _nw286 = new _module.C3();
          _nw286.__ctor(_1624_v0, _this.f7);
          _1680_v49 = _nw286;
          let _1681_v50;
          _1681_v50 = _dafny.Map.Empty.slice().updateUnsafe(((_1663_v33.f13) ? (new BigNumber((_1678_v47).length)) : (_1624_v0)),!((new BigNumber(-68)).isLessThan((_1680_v49).f14)));
          (_1663_v33).f13 = (((_1681_v50).contains(((_1663_v33.f13) ? (_1624_v0) : (_1624_v0)))) ? ((_1681_v50).get(((_1663_v33.f13) ? (_1624_v0) : (_1624_v0)))) : (_this.f7));
          let _1682_v51;
          _1682_v51 = _dafny.Seq.of((_1663_v33).f12);
          let _1683_v52;
          _1683_v52 = _dafny.Map.Empty.slice().updateUnsafe((_1682_v51)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-478)), ((_1684_v48) => function (_1685_i5) {
            return _1684_v48;
          })(_1679_v48))).length), new BigNumber((_1682_v51).length))],_1676_v45);
          _1683_v52 = (_1683_v52).update((_1663_v33).f12, _dafny.MultiSet.FromArray(_1682_v51));
          let _1686_v53;
          _1686_v53 = _dafny.Seq.of(_1625_v1);
          _1686_v53 = _1686_v53;
          let _1687_v54;
          _1687_v54 = _dafny.Map.Empty.slice().updateUnsafe((_1663_v33).f12,_1674_v43);
          let _1688_v55;
          let _nw287 = new _module.C1();
          _nw287.__ctor(!((((_1687_v54).contains(_1663_v33.f13)) ? ((_1687_v54).get(_1663_v33.f13)) : (_dafny.Map.Empty.slice().updateUnsafe(_1624_v0,new BigNumber((_1625_v1).length))))).contains(new BigNumber((_dafny.Seq.UnicodeFromString("fbw")).length)));
          _1688_v55 = _nw287;
        } else {
          _1625_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), _1679_v48), _1625_v1), _dafny.Seq.Concat(_1625_v1, _1625_v1));
          let _index257 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_1650_v23).length));
          (_1650_v23)[_index257] = new BigNumber((_dafny.Seq.Concat(_1625_v1, _1625_v1)).length);
          let _index258 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_1650_v23).length));
          (_1650_v23)[_index258] = (_this).fm15(globalState);
          let _1689_v56;
          _1689_v56 = _dafny.Seq.of(!(!(false)));
          _1624_v0 = new BigNumber((((((!(!(_1663_v33.f13))) ? (_1663_v33.f13) : (_1663_v33.f13))) ? (_1689_v56) : (_1689_v56))).length);
          let _1690_v57;
          _1690_v57 = _module.D13.create_DC35();
          let _1691_v58;
          _1691_v58 = _module.D13.create_DC37(_1690_v57);
          let _1692_v59;
          _1692_v59 = _module.D13.create_DC37(_1691_v58);
          let _1693_v60;
          _1693_v60 = _module.D13.create_DC37(_1691_v58);
          let _1694_v61;
          _1694_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1693_v60,_1676_v45);
          (_this).f7 = (_1676_v45).IsSubsetOf((((_1694_v61).contains(_1693_v60)) ? ((_1694_v61).get(_1693_v60)) : (_1676_v45)));
          let _1695_v62;
          _1695_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1649_v22)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_1649_v22).length))],_module.D1.create_DC4(true, _1663_v33.f13));
          let _1696_v63;
          _1696_v63 = _module.D1.create_DC4(_1663_v33.f13, _1663_v33.f13);
          (_1663_v33).f13 = (_1695_v62).equals(_dafny.Map.Empty.slice().updateUnsafe((_1649_v22)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_1649_v22).length))],_1696_v63));
        }
      }
      return;
    }
    m9(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let _1697_v0;
      _1697_v0 = _dafny.Seq.of(_this.f7);
      let _1698_v1;
      let _nw288 = new _module.C1();
      _nw288.__ctor(_dafny.Seq.IsPrefixOf(_1697_v0, _1697_v0));
      _1698_v1 = _nw288;
      let _1699_v2;
      _1699_v2 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1700_v3;
      _1700_v3 = _dafny.Set.fromElements(_1699_v2);
      let _1701_v4;
      _1701_v4 = _1700_v3;
      if (function (_source12) {
        let _1702___mcc_h0 = _source12;
        let _1703_cf47 = _1702___mcc_h0;
        return _this.f7;
      }(_1701_v4)) {
        let _1704_v5;
        _1704_v5 = new BigNumber(877);
        let _1705_v6;
        _1705_v6 = _dafny.Seq.of(_1704_v5, _1704_v5);
        let _1706_v7;
        _1706_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,!_dafny.areEqual(_1705_v6, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), ((_1707_v5) => function (_1708_i0) {
          return _1707_v5;
        })(_1704_v5)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), ((_1709_v2) => function (_1710_i1) {
          return _1709_v2;
        })(_1699_v2))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), ((_1711_v5) => function (_1712_i0) {
          return _1711_v5;
        })(_1704_v5))).length)), _1704_v5)));
        _1706_v7 = _1706_v7;
        let _1713_v8;
        let _nw289 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
        _1713_v8 = _nw289;
        let _index259 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1713_v8).length));
        (_1713_v8)[_index259] = _1705_v6;
        let _1714_v9;
        _1714_v9 = _dafny.Set.fromElements(_this.f7);
        let _index260 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1713_v8).length));
        let _rhs268 = !((_1714_v9).Union(_1714_v9)).contains(_this.f7);
        let _rhs269 = _1705_v6;
        let _rhs270 = _1705_v6;
        let _rhs271 = true;
        let _rhs272 = _this.f7;
        let _lhs161 = _this;
        let _lhs162 = _1713_v8;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1713_v8).length));
        let _lhs164 = _this;
        let _lhs165 = _this;
        _lhs161.f7 = _rhs268;
        _1705_v6 = _rhs269;
        _lhs162[_lhs163] = _rhs270;
        _lhs164.f7 = _rhs271;
        _lhs165.f7 = _rhs272;
        (_this).f7 = _this.f7;
        let _1715_v10;
        _1715_v10 = _module.D9.create_DC26((_1713_v8)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_1713_v8).length))]);
        let _1716_v11;
        _1716_v11 = _module.D12.create_DC31(_1715_v10, _this.f7);
        if (!((_1716_v11).dtor_cf54) || (_this.f7)) {
          r2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.Seq.UnicodeFromString("suupjuvr"));
          let _1717_v12;
          _1717_v12 = _module.D0.create_DC0(_1704_v5);
          let _1718_v13;
          let _nw290 = new _module.C3();
          _nw290.__ctor(((_module.__default.fm1(_dafny.Seq.of(_1717_v12, _1717_v12, _1717_v12, _1717_v12, _1717_v12), _1704_v5, _module.__default.fm0(globalState), globalState)) ? (_1704_v5) : (_1704_v5)), _this.f7);
          _1718_v13 = _nw290;
          (_this).f7 = ((_this.f7) ? (_this.f7) : ((_1697_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), ((_1719_v13) => function (_1720_i2) {
            return (_1719_v13).f14;
          })(_1718_v13))).length), new BigNumber((_1697_v0).length))]));
          let _1721_v14;
          _1721_v14 = _dafny.MultiSet.fromElements(_1704_v5);
          _1700_v3 = _module.__default.fm30((_1721_v14).IsDisjointFrom(_1721_v14), _1704_v5, (new BigNumber(718)).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_1704_v5))), globalState);
          let _1722_v15;
          _1722_v15 = _dafny.Seq.of(_1717_v12);
          (_this).f7 = ((_module.__default.fm1(_module.__default.fm51(_module.__default.fm18(globalState), _this.f7, globalState), _1704_v5, _1704_v5, globalState)) ? (_this.f7) : (_module.__default.fm1(_1722_v15, _1704_v5, new BigNumber((_1705_v6).length), globalState)));
        } else {
          _1704_v5 = _1704_v5;
          let _1723_v16;
          _1723_v16 = _dafny.Seq.UnicodeFromString("mj");
          let _1724_v17;
          let _nw291 = Array((new BigNumber(7)).toNumber());
          _nw291[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1723_v16, _1723_v16);
          _nw291[(_dafny.ONE).toNumber()] = _1723_v16;
          _nw291[(new BigNumber(2)).toNumber()] = _1723_v16;
          _nw291[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1723_v16, _1723_v16);
          _nw291[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(460)), ((_1725_v2) => function (_1726_i3) {
            return _1725_v2;
          })(_1699_v2));
          _nw291[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), ((_1727_v2) => function (_1728_i4) {
            return _1727_v2;
          })(_1699_v2));
          _nw291[(new BigNumber(6)).toNumber()] = _1723_v16;
          _1724_v17 = _nw291;
          let _index261 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1724_v17).length));
          (_1724_v17)[_index261] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), ((_1729_v2) => function (_1730_i5) {
            return _1729_v2;
          })(_1699_v2));
          let _1731_v18;
          _1731_v18 = _dafny.Seq.of(_1723_v16, _1723_v16, _1723_v16);
          let _index262 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1724_v17).length));
          (_1724_v17)[_index262] = _dafny.Seq.Concat(((_this.f7) ? (_1723_v16) : (_1723_v16)), _dafny.Seq.Concat((_1731_v18)[_module.__default.safeIndex(_1704_v5, new BigNumber((_1731_v18).length))], _1723_v16));
          (_this).f7 = _this.f7;
          let _1732_v19;
          _1732_v19 = _dafny.Set.fromElements(_1714_v9);
          let _1733_v20;
          _1733_v20 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1732_v19).length)));
          let _1734_v21;
          _1734_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1733_v20,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-716)), ((_1735_v2) => function (_1736_i6) {
            return _1735_v2;
          })(_1699_v2)));
          _1734_v21 = _1734_v21;
          let _1737_v22;
          _1737_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1704_v5,_this.f7);
          (_this).f7 = (((_1737_v22).contains(_1704_v5)) ? ((_1737_v22).get(_1704_v5)) : (_this.f7));
        }
        let _1738_v23;
        _1738_v23 = _dafny.MultiSet.fromElements(false, _this.f7, !(false), _this.f7);
        let _1739_v24;
        _1739_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(476),(_1738_v23).update(_this.f7, _module.__default.abs(_1704_v5)));
        (_this).f7 = (_1738_v23).IsProperSubsetOf((((_1739_v24).contains(_1704_v5)) ? ((_1739_v24).get(_1704_v5)) : (_1738_v23)));
      } else {
        let _1740_v25;
        let _nw292 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1740_v25 = _nw292;
        let _1741_v26;
        _1741_v26 = _dafny.Set.fromElements(_1740_v25);
        let _1742_v27;
        _1742_v27 = new BigNumber(72);
        let _1743_v28;
        _1743_v28 = _dafny.Set.fromElements(_this.f7, _this.f7, _this.f7, _this.f7, _this.f7);
        let _1744_v29;
        let _nw293 = new _module.C2();
        _nw293.__ctor(_1741_v26, _1699_v2, !(_1742_v27).isEqualTo(new BigNumber((_1743_v28).length)));
        _1744_v29 = _nw293;
        (_this).f7 = (_1697_v0)[_module.__default.safeIndex(_1742_v27, new BigNumber((_1697_v0).length))];
        (_this).f7 = _this.f7;
        (_this).f7 = _this.f7;
        (_this).f7 = _this.f7;
      }
      let _1745_v30;
      _1745_v30 = new BigNumber(928);
      if ((_1697_v0)[_module.__default.safeIndex(_1745_v30, new BigNumber((_1697_v0).length))]) {
        if (_module.__default.fm1(_module.__default.fm51(_module.__default.fm18(globalState), _this.f7, globalState), _module.__default.safeDivisionInt(_1745_v30, _1745_v30), (_1745_v30).multipliedBy(_1745_v30), globalState)) {
          let _1746_v31;
          let _nw294 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
          _1746_v31 = _nw294;
          let _1747_v32;
          let _init55 = ((_1748_v30, _1749_v0) => function (_1750_i7) {
            return _module.D3.create_DC11(new BigNumber(263), _1748_v30, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1749_v0).length),_this.f7)).length), _this.f7);
          })(_1745_v30, _1697_v0);
          let _nw295 = Array((new BigNumber(19)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw295.length); _i0_55++) {
            _nw295[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _1747_v32 = _nw295;
          let _1751_v33;
          _1751_v33 = _dafny.Set.fromElements(_1747_v32);
          let _index263 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1746_v31).length));
          (_1746_v31)[_index263] = _1751_v33;
          let _index264 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1746_v31).length));
          (_1746_v31)[_index264] = ((_1751_v33).Difference(_1751_v33)).Intersect(_1751_v33);
          _1745_v30 = _1745_v30;
          let _1752_v34;
          _1752_v34 = _dafny.Seq.UnicodeFromString("dcfdxmh");
          let _1753_v35;
          _1753_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1699_v2,_dafny.Seq.Concat(_1752_v34, _1752_v34));
          let _1754_v36;
          let _nw296 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1754_v36 = _nw296;
          let _index265 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1754_v36).length));
          (_1754_v36)[_index265] = _1745_v30;
          let _1755_v37;
          _1755_v37 = _dafny.Seq.of(_1745_v30);
          let _1756_v38;
          _1756_v38 = _dafny.Set.fromElements(_1745_v30, _1745_v30, (_dafny.ZERO).minus(new BigNumber((_1755_v37).length)), _1745_v30, _1745_v30);
          let _1757_v39;
          _1757_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1745_v30);
          let _pat_let_tv43 = _1753_v35;
          let _index266 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1754_v36).length));
          let _rhs273 = _module.__default.safeDivisionInt(_1745_v30, _1745_v30);
          let _rhs274 = (function (_pat_let53_0) {
            return function (_1758_dt__update__tmp_h0) {
              return function (_pat_let54_0) {
                return function (_1759_dt__update_hcf62_h0) {
                  return _module.D14.create_DC38(_1759_dt__update_hcf62_h0);
                }(_pat_let54_0);
              }(_pat_let_tv43);
            }(_pat_let53_0);
          }(_module.D14.create_DC38(_1753_v35))).dtor_cf62;
          let _rhs275 = _module.__default.safeModuloInt((_this).fm15(globalState), _1745_v30);
          let _rhs276 = (_dafny.ZERO).minus(_1745_v30);
          let _rhs277 = (_1698_v1).fm32((_1745_v30).isLessThan(_1745_v30), _1756_v38, new BigNumber(728), new BigNumber((_1757_v39).length), globalState);
          let _lhs166 = _1754_v36;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1754_v36).length));
          let _lhs168 = _this;
          r1 = _rhs273;
          _1753_v35 = _rhs274;
          _lhs166[_lhs167] = _rhs275;
          r1 = _rhs276;
          _lhs168.f7 = _rhs277;
          let _1760_v40;
          let _nw297 = Array((new BigNumber(8)).toNumber());
          _nw297[(_dafny.ZERO).toNumber()] = true;
          _nw297[(_dafny.ONE).toNumber()] = false;
          _nw297[(new BigNumber(2)).toNumber()] = _this.f7;
          _nw297[(new BigNumber(3)).toNumber()] = _this.f7;
          _nw297[(new BigNumber(4)).toNumber()] = _this.f7;
          _nw297[(new BigNumber(5)).toNumber()] = _this.f7;
          _nw297[(new BigNumber(6)).toNumber()] = _this.f7;
          _nw297[(new BigNumber(7)).toNumber()] = _this.f7;
          _1760_v40 = _nw297;
          let _1761_v41;
          _1761_v41 = _dafny.Set.fromElements(_1760_v40, _1760_v40);
          let _1762_v42;
          let _nw298 = new _module.C2();
          _nw298.__ctor(_1761_v41, _1699_v2, _this.f7);
          _1762_v42 = _nw298;
          let _1763_v43;
          _1763_v43 = _dafny.MultiSet.fromElements(_this.f7);
          let _1764_v44;
          _1764_v44 = _module.D15.create_DC42(_dafny.MultiSet.fromElements(_this.f7));
          let _1765_v45;
          _1765_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1697_v0);
          let _1766_v46;
          _1766_v46 = _module.D0.create_DC0((_1754_v36)[_module.__default.safeIndex(new BigNumber(120), new BigNumber((_1754_v36).length))]);
          let _1767_v47;
          _1767_v47 = _dafny.Seq.of(_1766_v46);
          let _1768_v48;
          _1768_v48 = _dafny.Seq.of(_1763_v43, (_1764_v44).dtor_cf69, (_1763_v43).update(true, _module.__default.abs(_1745_v30)), (_module.__default.fm49(!(_this.f7), _1765_v45, _this.f7, globalState)).Difference(_dafny.MultiSet.fromElements(_module.__default.fm1(_1767_v47, _1745_v30, _1745_v30, globalState))), (_1763_v43).Difference(_dafny.MultiSet.fromElements(false, _this.f7)));
          r1 = new BigNumber((_1768_v48).length);
        } else {
          let _1769_v49;
          _1769_v49 = _dafny.Set.fromElements(_this.f7);
          let _1770_v50;
          _1770_v50 = _module.D1.create_DC3(_1769_v49);
          let _1771_v51;
          _1771_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1745_v30)).cardinality()),_this.f7);
          let _1772_v52;
          _1772_v52 = _dafny.Seq.UnicodeFromString("syvodnk");
          let _1773_v53;
          _1773_v53 = _module.D3.create_DC11(_1745_v30, _1745_v30, new BigNumber((_1772_v52).length), _this.f7);
          let _1774_v54;
          _1774_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1745_v30,_1772_v52);
          let _1775_v55;
          _1775_v55 = _dafny.Set.fromElements(_1745_v30, (_dafny.ZERO).minus(_1745_v30));
          let _1776_v56;
          let _nw299 = Array((new BigNumber(26)).toNumber());
          _nw299[(_dafny.ZERO).toNumber()] = _1745_v30;
          _nw299[(_dafny.ONE).toNumber()] = new BigNumber(694);
          _nw299[(new BigNumber(2)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(3)).toNumber()] = _module.__default.fm0(globalState);
          _nw299[(new BigNumber(4)).toNumber()] = new BigNumber(((_dafny.Set.fromElements(_this.f7, _this.f7)).Difference((_1770_v50).dtor_cf6)).length);
          _nw299[(new BigNumber(5)).toNumber()] = ((true) ? (_1745_v30) : (_1745_v30));
          _nw299[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(138), new BigNumber((_1769_v49).length));
          _nw299[(new BigNumber(7)).toNumber()] = (_1745_v30).plus(_1745_v30);
          _nw299[(new BigNumber(8)).toNumber()] = ((_this.f7) ? (_1745_v30) : (new BigNumber((_1697_v0).length)));
          _nw299[(new BigNumber(9)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(10)).toNumber()] = new BigNumber((((_this.f7) ? (_1771_v51) : (_1771_v51))).length);
          _nw299[(new BigNumber(11)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(12)).toNumber()] = (new BigNumber(919)).plus(_1745_v30);
          _nw299[(new BigNumber(13)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(14)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(15)).toNumber()] = (_1773_v53).dtor_cf23;
          _nw299[(new BigNumber(16)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(17)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(18)).toNumber()] = (_this).fm15(globalState);
          _nw299[(new BigNumber(19)).toNumber()] = new BigNumber(772);
          _nw299[(new BigNumber(20)).toNumber()] = ((_this.f7) ? (_module.__default.fm0(globalState)) : (_1745_v30));
          _nw299[(new BigNumber(21)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(22)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(23)).toNumber()] = _1745_v30;
          _nw299[(new BigNumber(24)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1772_v52, (((_1774_v54).contains((_dafny.ZERO).minus(new BigNumber((_1775_v55).length)))) ? ((_1774_v54).get((_dafny.ZERO).minus(new BigNumber((_1775_v55).length)))) : (_1772_v52)))).length);
          _nw299[(new BigNumber(25)).toNumber()] = _1745_v30;
          _1776_v56 = _nw299;
          let _index267 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length));
          (_1776_v56)[_index267] = (_dafny.ZERO).minus(_1745_v30);
          let _index268 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length));
          (_1776_v56)[_index268] = new BigNumber((_dafny.Seq.Concat(_1697_v0, _1697_v0)).length);
          let _1777_v57;
          let _init56 = ((_1778_v0) => function (_1779_i8) {
            return (_dafny.MultiSet.fromElements(_1778_v0)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1778_v0));
          })(_1697_v0);
          let _nw300 = Array((new BigNumber(10)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw300.length); _i0_56++) {
            _nw300[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _1777_v57 = _nw300;
          _1777_v57 = _1777_v57;
          let _1780_v58;
          let _nw301 = Array((new BigNumber(19)).toNumber()).fill(_module.D13.Default());
          _1780_v58 = _nw301;
          let _1781_v60;
          _1781_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1776_v56,function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of _dafny.IntegerRange(new BigNumber(262), new BigNumber(582))) {
              let _1782_v59 = _compr_34;
              if (((new BigNumber(262)).isLessThanOrEqualTo(_1782_v59)) && ((_1782_v59).isLessThan(new BigNumber(582)))) {
                _coll34.add(_module.__default.safeDivisionInt(_1782_v59, (_1776_v56)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length))]));
              }
            }
            return _coll34;
          }());
          let _1783_v61;
          _1783_v61 = _module.D13.create_DC34(_1781_v60);
          let _1784_v62;
          _1784_v62 = _module.D13.create_DC37(_1783_v61);
          let _index269 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1780_v58).length));
          (_1780_v58)[_index269] = _1784_v62;
          let _index270 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1780_v58).length));
          (_1780_v58)[_index270] = _module.D13.create_DC37(_1783_v61);
          let _1785_v63;
          _1785_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(_this.f7, globalState),_dafny.Seq.IsPrefixOf(_1772_v52, _1772_v52));
          let _1786_v64;
          _1786_v64 = _dafny.Seq.of(new BigNumber((_1772_v52).length), (_1776_v56)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length))], (_1776_v56)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length))], (_1776_v56)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length))], new BigNumber(28));
          _1785_v63 = (_1785_v63).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1786_v64,_this.f7));
          let _index271 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1776_v56).length));
          (_1776_v56)[_index271] = (_dafny.ZERO).minus(new BigNumber((_1772_v52).length));
        }
        r1 = _module.__default.safeDivisionInt((_1745_v30).plus(_1745_v30), _1745_v30);
        let _1787_v65;
        _1787_v65 = _module.D6.create_DC15(_1699_v2);
        _1787_v65 = _1787_v65;
        _1697_v0 = _dafny.Seq.Concat(_1697_v0, _dafny.Seq.Concat(_dafny.Seq.update(_1697_v0, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1697_v0)).cardinality()), new BigNumber((_1697_v0).length)), _this.f7), _dafny.Seq.of(_this.f7)));
        let _1788_v66;
        let _nw302 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1788_v66 = _nw302;
        let _1789_v67;
        _1789_v67 = _module.D0.create_DC2(!(_this.f7), new BigNumber((_1697_v0).length), _dafny.Set.fromElements(_1788_v66), _1745_v30);
        let _1790_v68;
        _1790_v68 = _module.D0.create_DC2(_this.f7, (_1789_v67).dtor_cf5, _dafny.Set.fromElements(_1788_v66, _1788_v66), _module.__default.fm0(globalState));
        r1 = (_1789_v67).dtor_cf5;
      } else {
        let _1791_v69;
        _1791_v69 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1697_v0)[_module.__default.safeIndex(_1745_v30, new BigNumber((_1697_v0).length))]);
        let _1792_v70;
        _1792_v70 = _dafny.MultiSet.fromElements(new BigNumber((_1791_v69).length));
        let _1793_v71;
        _1793_v71 = _dafny.MultiSet.fromElements(_1699_v2);
        let _1794_v72;
        _1794_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_this.f7)).update(_this.f7, _module.__default.abs(_1745_v30)),_this.f7);
        let _1795_v74;
        _1795_v74 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(((_1792_v70).contains((((_1793_v71).contains(_1699_v2)) ? ((_1793_v71).get(_1699_v2)) : (new BigNumber((function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of (_1794_v72).Keys.Elements) {
            let _1797_v73 = _compr_36;
            if ((_1794_v72).contains(_1797_v73)) {
              _coll36.add(_1797_v73);
            }
          }
          return _coll36;
        }()).length))))) ? ((_1792_v70).get((((_1793_v71).contains(_1699_v2)) ? ((_1793_v71).get(_1699_v2)) : (new BigNumber((function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of (_1794_v72).Keys.Elements) {
            let _1796_v73 = _compr_35;
            if ((_1794_v72).contains(_1796_v73)) {
              _coll35.add(_1796_v73);
            }
          }
          return _coll35;
        }()).length))))) : (_1745_v30)));
        _1795_v74 = (_1795_v74).update(_this.f7, _1745_v30);
        _1697_v0 = _dafny.Seq.Concat(_1697_v0, _1697_v0);
        let _1798_v75;
        _1798_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1697_v0,_1745_v30);
        _1798_v75 = (_1798_v75).Merge((_1798_v75).Merge(_1798_v75));
        let _1799_v76;
        _1799_v76 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, _this.f7, _this.f7, _this.f7);
        let _1800_v77;
        _1800_v77 = _dafny.Seq.UnicodeFromString("snip");
        _1745_v30 = _module.__default.safeDivisionInt((((_1799_v76).contains(_this.f7)) ? ((_1799_v76).get(_this.f7)) : (_1745_v30)), new BigNumber((_1800_v77).length));
        _1745_v30 = ((_this).fm15(globalState)).minus((new BigNumber((_1800_v77).length)).minus(_1745_v30));
      }
      let _1801_v78;
      _1801_v78 = _dafny.Seq.UnicodeFromString("sbunisxop");
      let _1802_v79;
      let _nw303 = new _module.C3();
      _nw303.__ctor(_1745_v30, !(_this.f7));
      _1802_v79 = _nw303;
      let _1803_v80;
      _1803_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1802_v79);
      let _1804_v81;
      _1804_v81 = _dafny.MultiSet.fromElements(_this.f7);
      let _1805_v82;
      _1805_v82 = _module.D0.create_DC0(new BigNumber((_1804_v81).cardinality()));
      let _1806_v83;
      _1806_v83 = _dafny.Seq.of(_1805_v82, _1805_v82, _1805_v82);
      let _1807_v84;
      _1807_v84 = _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('n'.codePointAt(0)));
      if (_module.__default.fm1(_module.__default.fm51(_1699_v2, _this.f7, globalState), new BigNumber((_dafny.Seq.Concat(_1801_v78, _module.__default.fm35(new BigNumber((_dafny.Seq.of(_1802_v79, (((_1803_v80).contains(_module.__default.fm1(_1806_v83, _1745_v30, _1745_v30, globalState))) ? ((_1803_v80).get(_module.__default.fm1(_1806_v83, _1745_v30, _1745_v30, globalState))) : (_1802_v79)), _1802_v79, _1802_v79)).length), new BigNumber((_1807_v84).length), globalState))).length), (new BigNumber((_1801_v78).length)).minus((_dafny.ZERO).minus(_1745_v30)), globalState)) {
        (_this).f7 = !(!(_this.f7)) || (_this.f7);
        _1745_v30 = _module.__default.safeModuloInt(_1745_v30, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-723)), ((_1808_v2) => function (_1809_i9) {
          return _1808_v2;
        })(_1699_v2))).length));
        if ((_module.__default.fm30(!(_this.f7), (_1802_v79).f14, false, globalState)).equals(_1700_v3)) {
          let _1810_v85;
          let _nw304 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1810_v85 = _nw304;
          let _1811_v86;
          _1811_v86 = _module.D9.create_DC23(_1810_v85);
          let _rhs278 = ((_dafny.ZERO).minus((_1802_v79).f14)).minus((_1802_v79).f14);
          let _rhs279 = _1811_v86;
          r1 = _rhs278;
          _1811_v86 = _rhs279;
          _1745_v30 = _module.__default.safeDivisionInt((_1802_v79).f14, ((_1802_v79).f14).minus((_1802_v79).fm9(_this.f7, _1745_v30, _1699_v2, globalState)));
          let _1812_v87;
          _1812_v87 = _module.D6.create_DC17(!(_this.f7), (_1802_v79).f14, _module.__default.fm26(_this.f7, _1745_v30, (_1802_v79).f14, globalState), _1699_v2);
          let _1813_v88;
          _1813_v88 = _module.D6.create_DC18(_1812_v87);
          let _1814_v89;
          _1814_v89 = _module.D14.create_DC40(_1813_v88, _this.f7, _this.f7);
          _1814_v89 = _1814_v89;
          let _1815_v90;
          let _init57 = ((_1816_v79) => function (_1817_i10) {
            return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,(_1816_v79).f14), _dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),new BigNumber(676)), _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7), _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7))).cardinality())));
          })(_1802_v79);
          let _nw305 = Array((new BigNumber(12)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw305.length); _i0_57++) {
            _nw305[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _1815_v90 = _nw305;
          let _1818_v91;
          _1818_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_dafny.ZERO).minus((_1802_v79).f14));
          let _1819_v92;
          _1819_v92 = _dafny.Seq.of(_1818_v91, _1818_v91, _1818_v91);
          let _index272 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1815_v90).length));
          (_1815_v90)[_index272] = _1819_v92;
          let _1820_v93;
          _1820_v93 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1818_v91).length),_1697_v0);
          let _1821_v94;
          _1821_v94 = _dafny.Seq.of(_dafny.Seq.of(_this.f7, _this.f7), (((_1820_v93).contains((_1802_v79).f14)) ? ((_1820_v93).get((_1802_v79).f14)) : (_1697_v0)), _1697_v0);
          let _1822_v95;
          _1822_v95 = _dafny.Map.Empty.slice().updateUnsafe((_1697_v0)[_module.__default.safeIndex((_dafny.ZERO).minus((((_1804_v81).contains(_this.f7)) ? ((_1804_v81).get(_this.f7)) : (new BigNumber((_1804_v81).cardinality())))), new BigNumber((_1697_v0).length))],_1821_v94);
          let _1823_v96;
          _1823_v96 = _dafny.Seq.of(_1821_v94);
          let _index273 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1815_v90).length));
          let _rhs280 = _1819_v92;
          let _rhs281 = (((_1822_v95).contains((_1802_v79).fm8((_1802_v79).f14, globalState))) ? ((_1822_v95).get((_1802_v79).fm8((_1802_v79).f14, globalState))) : (((_this.f7) ? ((_1823_v96)[_module.__default.safeIndex((_1802_v79).f14, new BigNumber((_1823_v96).length))]) : (_1821_v94))));
          let _lhs169 = _1815_v90;
          let _lhs170 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1815_v90).length));
          _lhs169[_lhs170] = _rhs280;
          _1821_v94 = _rhs281;
          let _1824_v97;
          _1824_v97 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
          _1824_v97 = _1824_v97;
        } else {
          _1801_v78 = _dafny.Seq.UnicodeFromString("kefhqejr");
          let _1825_v98;
          _1825_v98 = _module.D12.create_DC32(_this.f7);
          let _1826_v99;
          _1826_v99 = _dafny.Map.Empty.slice().updateUnsafe((_1802_v79).f14,_this.f7);
          let _1827_v100;
          _1827_v100 = _dafny.MultiSet.fromElements((_1802_v79).f14, new BigNumber((_1826_v99).length));
          let _1828_v101;
          _1828_v101 = _dafny.Seq.of(_1827_v100);
          let _1829_v102;
          _1829_v102 = _dafny.Seq.of((_1802_v79).f14);
          let _1830_v103;
          _1830_v103 = _dafny.Map.Empty.slice().updateUnsafe((_1802_v79).f14,(_1802_v79).f14);
          let _1831_v104;
          _1831_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1745_v30,_1830_v103);
          let _1832_v105;
          let _nw306 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1832_v105 = _nw306;
          let _1833_v106;
          _1833_v106 = _dafny.MultiSet.fromElements(_1832_v105);
          _1825_v98 = _module.__default.fm52(_dafny.Seq.Concat(_1828_v101, _dafny.Seq.of(_dafny.MultiSet.FromArray(_1829_v102), _1827_v100, _1827_v100, (_1827_v100).update(new BigNumber(((((_1831_v104).contains(new BigNumber((_1833_v106).cardinality()))) ? ((_1831_v104).get(new BigNumber((_1833_v106).cardinality()))) : (_1830_v103))).length), _module.__default.abs((_1802_v79).f14)), _dafny.MultiSet.FromArray(_1829_v102))), _1745_v30, _module.__default.safeModuloInt((_1802_v79).f14, _1745_v30), _this.f7, globalState);
          let _1834_v107;
          _1834_v107 = _dafny.Set.fromElements(_1832_v105, _1832_v105);
          let _1835_v108;
          _1835_v108 = _module.D9.create_DC25();
          let _1836_v109;
          let _nw307 = new _module.C2();
          _nw307.__ctor(_1834_v107, _module.__default.fm37(_this.f7, _1835_v108, globalState), (new BigNumber(826)).isLessThanOrEqualTo(new BigNumber(255)));
          _1836_v109 = _nw307;
          r1 = _module.__default.safeModuloInt((_1802_v79).f14, (_1745_v30).multipliedBy((_1802_v79).f14));
          let _1837_v110;
          let _nw308 = new _module.C0();
          _nw308.__ctor(_this.f7, false);
          _1837_v110 = _nw308;
        }
        let _1838_v111;
        _1838_v111 = _dafny.Seq.of((_1802_v79).f14, (_1802_v79).f14);
        let _1839_v112;
        _1839_v112 = _module.D9.create_DC26(_1838_v111);
        let _1840_v113;
        let _nw309 = Array((new BigNumber(10)).toNumber());
        _nw309[(_dafny.ZERO).toNumber()] = _1839_v112;
        _nw309[(_dafny.ONE).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(2)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(3)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(4)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(5)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(6)).toNumber()] = ((_this.f7) ? (_module.D9.create_DC26(_1838_v111)) : (_1839_v112));
        _nw309[(new BigNumber(7)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(8)).toNumber()] = _1839_v112;
        _nw309[(new BigNumber(9)).toNumber()] = _1839_v112;
        _1840_v113 = _nw309;
        let _index274 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1840_v113).length));
        (_1840_v113)[_index274] = _1839_v112;
        let _index275 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1840_v113).length));
        (_1840_v113)[_index275] = _module.D9.create_DC26(_1838_v111);
        (_this).f7 = _this.f7;
      } else {
        let _1841_v114;
        _1841_v114 = _dafny.Set.fromElements(_1745_v30);
        _1745_v30 = new BigNumber(((_1841_v114).Difference(_1841_v114)).length);
        if ((_dafny.MultiSet.FromArray(_1697_v0)).IsDisjointFrom(_1804_v81)) {
          _1745_v30 = (_1802_v79).f14;
          (_this).f7 = (new BigNumber(674)).isLessThan((_1802_v79).f14);
          let _1842_v115;
          let _nw310 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1842_v115 = _nw310;
          let _index276 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_1842_v115).length));
          (_1842_v115)[_index276] = !(_this.f7) || (true);
          let _index277 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_1842_v115).length));
          (_1842_v115)[_index277] = true;
          _1842_v115 = _1842_v115;
          let _1843_v116;
          _1843_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1699_v2,_dafny.Seq.UnicodeFromString("h"));
          let _1844_v117;
          _1844_v117 = _module.D14.create_DC38(_1843_v116);
          let _1845_v118;
          _1845_v118 = _dafny.Map.Empty.slice().updateUnsafe((_1802_v79).f14,_1844_v117);
          let _1846_v119;
          _1846_v119 = _module.D14.create_DC41((((_1845_v118).contains(_1745_v30)) ? ((_1845_v118).get(_1745_v30)) : (_1844_v117)));
          _1846_v119 = _1846_v119;
        } else {
          (_this).f7 = _this.f7;
          (_this).f7 = _this.f7;
          let _1847_v120;
          let _nw311 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1847_v120 = _nw311;
          let _1848_v121;
          let _nw312 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1848_v121 = _nw312;
          let _index278 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1847_v120).length));
          (_1847_v120)[_index278] = ((_this.f7) ? (_1848_v121) : (_1848_v121));
          let _index279 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1847_v120).length));
          let _nw313 = Array((new BigNumber(9)).toNumber());
          _nw313[(_dafny.ZERO).toNumber()] = _1699_v2;
          _nw313[(_dafny.ONE).toNumber()] = _1699_v2;
          _nw313[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw313[(new BigNumber(3)).toNumber()] = _1699_v2;
          _nw313[(new BigNumber(4)).toNumber()] = _1699_v2;
          _nw313[(new BigNumber(5)).toNumber()] = _1699_v2;
          _nw313[(new BigNumber(6)).toNumber()] = (((_1807_v84).contains(!(_this.f7))) ? ((_1807_v84).get(!(_this.f7))) : (_1699_v2));
          _nw313[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw313[(new BigNumber(8)).toNumber()] = _1699_v2;
          (_1847_v120)[_index279] = _nw313;
          _1801_v78 = _1801_v78;
          let _1849_v122;
          _1849_v122 = _dafny.MultiSet.fromElements(new BigNumber(209));
          let _1850_v123;
          _1850_v123 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),_this.f7);
          let _1851_v124;
          let _nw314 = Array((new BigNumber(21)).toNumber());
          _nw314[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw314[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements((_1802_v79).f14, (_1802_v79).f14)).IsProperSubsetOf(_1849_v122);
          _nw314[(new BigNumber(2)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(3)).toNumber()] = true;
          _nw314[(new BigNumber(4)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(5)).toNumber()] = (_this.f7) === (false);
          _nw314[(new BigNumber(6)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(7)).toNumber()] = !((((_1850_v123).contains(_this.f7)) ? ((_1850_v123).get(_this.f7)) : (_this.f7)));
          _nw314[(new BigNumber(8)).toNumber()] = (!(_this.f7)) && (_this.f7);
          _nw314[(new BigNumber(9)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(10)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(11)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(12)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(13)).toNumber()] = ((_this.f7) ? (!(_this.f7)) : (_this.f7));
          _nw314[(new BigNumber(14)).toNumber()] = ((_1802_v79).f14).isLessThan((_1802_v79).f14);
          _nw314[(new BigNumber(15)).toNumber()] = ((_dafny.ZERO).minus(_1745_v30)).isLessThan((_1802_v79).f14);
          _nw314[(new BigNumber(16)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(17)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(18)).toNumber()] = _this.f7;
          _nw314[(new BigNumber(19)).toNumber()] = (_1697_v0)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState))), new BigNumber((_1697_v0).length))];
          _nw314[(new BigNumber(20)).toNumber()] = _this.f7;
          _1851_v124 = _nw314;
          let _index280 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1851_v124).length));
          (_1851_v124)[_index280] = _this.f7;
          let _index281 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1851_v124).length));
          let _rhs282 = (_1801_v78)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_1841_v114).length), new BigNumber((_1849_v122).cardinality())), new BigNumber((_1801_v78).length))];
          let _rhs283 = _this.f7;
          let _lhs171 = _1851_v124;
          let _lhs172 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1851_v124).length));
          r0 = _rhs282;
          _lhs171[_lhs172] = _rhs283;
        }
        let _1852_v125;
        let _nw315 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1852_v125 = _nw315;
        if (_dafny.Seq.contains(_dafny.Seq.of(_1852_v125, _1852_v125, _1852_v125), _1852_v125)) {
          let _index282 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1852_v125).length));
          (_1852_v125)[_index282] = (_1802_v79).f14;
          let _1853_v126;
          let _nw316 = Array((new BigNumber(12)).toNumber()).fill([]);
          _1853_v126 = _nw316;
          let _1854_v127;
          let _init58 = function (_1855_i11) {
            return _this.f7;
          };
          let _nw317 = Array((new BigNumber(13)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw317.length); _i0_58++) {
            _nw317[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1854_v127 = _nw317;
          let _index283 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1853_v126).length));
          (_1853_v126)[_index283] = _1854_v127;
          let _1856_v128;
          _1856_v128 = _dafny.MultiSet.fromElements(_1745_v30);
          let _index284 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1852_v125).length));
          let _index285 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1853_v126).length));
          let _rhs284 = (((_1856_v128).contains((_1802_v79).fm9(_this.f7, (_dafny.ZERO).minus(_1745_v30), _1699_v2, globalState))) ? ((_1856_v128).get((_1802_v79).fm9(_this.f7, (_dafny.ZERO).minus(_1745_v30), _1699_v2, globalState))) : (_module.__default.fm0(globalState)));
          let _rhs285 = _this.f7;
          let _rhs286 = (_1802_v79).f14;
          let _rhs287 = _1854_v127;
          let _lhs173 = _this;
          let _lhs174 = _1852_v125;
          let _lhs175 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1852_v125).length));
          let _lhs176 = _1853_v126;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1853_v126).length));
          _1745_v30 = _rhs284;
          _lhs173.f7 = _rhs285;
          _lhs174[_lhs175] = _rhs286;
          _lhs176[_lhs177] = _rhs287;
          _1745_v30 = (_1802_v79).f14;
          let _1857_v129;
          let _init59 = ((_1858_v79) => function (_1859_i12) {
            return (_1859_i12).minus((_1858_v79).f14);
          })(_1802_v79);
          let _nw318 = Array((new BigNumber(7)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw318.length); _i0_59++) {
            _nw318[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _1857_v129 = _nw318;
          _1857_v129 = _1857_v129;
          let _1860_v130;
          _1860_v130 = _module.D6.create_DC17(_this.f7, (_this).fm15(globalState), _1699_v2, _1699_v2);
          let _1861_v131;
          _1861_v131 = _module.D6.create_DC18(_1860_v130);
          let _1862_v132;
          _1862_v132 = _module.D6.create_DC18(_1860_v130);
          let _1863_v133;
          _1863_v133 = _module.D14.create_DC40(_1862_v132, false, _this.f7);
          let _1864_v134;
          let _init60 = ((_1865_v2) => function (_1866_i13) {
            return _1865_v2;
          })(_1699_v2);
          let _nw319 = Array((new BigNumber(5)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw319.length); _i0_60++) {
            _nw319[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _1864_v134 = _nw319;
          let _1867_v135;
          _1867_v135 = _dafny.MultiSet.fromElements(_1864_v134, _1864_v134);
          r1 = (((_1863_v133).dtor_cf67) ? ((_1802_v79).f14) : ((((_1867_v135).contains(_1864_v134)) ? ((_1867_v135).get(_1864_v134)) : ((_1802_v79).f14))));
          let _1868_v136;
          let _nw320 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
          _1868_v136 = _nw320;
          _1868_v136 = _1868_v136;
        } else {
          let _index286 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1852_v125).length));
          (_1852_v125)[_index286] = (_1802_v79).f14;
          let _index287 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1852_v125).length));
          (_1852_v125)[_index287] = _module.__default.safeDivisionInt(((_1802_v79).f14).multipliedBy(_1745_v30), ((_1802_v79).f14).minus(new BigNumber(239)));
          let _1869_v137;
          let _init61 = function (_1870_i14) {
            return _this.f7;
          };
          let _nw321 = Array((new BigNumber(10)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw321.length); _i0_61++) {
            _nw321[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _1869_v137 = _nw321;
          let _index288 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1869_v137).length));
          (_1869_v137)[_index288] = _this.f7;
          let _index289 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1869_v137).length));
          let _rhs288 = (_this.f7) && (_this.f7);
          let _rhs289 = !(false) || (_this.f7);
          let _rhs290 = new BigNumber(690);
          let _lhs178 = _1869_v137;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1869_v137).length));
          let _lhs180 = _this;
          _lhs178[_lhs179] = _rhs288;
          _lhs180.f7 = _rhs289;
          _1745_v30 = _rhs290;
          let _index290 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1869_v137).length));
          (_1869_v137)[_index290] = true;
          let _1871_v138;
          let _nw322 = new _module.C1();
          _nw322.__ctor(_this.f7);
          _1871_v138 = _nw322;
          let _rhs291 = _module.__default.safeModuloInt((_1802_v79).f14, (_1852_v125)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_1852_v125).length))]);
          let _rhs292 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_1802_v79).f14), (_1852_v125)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_1852_v125).length))]);
          let _rhs293 = (_1802_v79).f14;
          r1 = _rhs291;
          _1745_v30 = _rhs292;
          _1745_v30 = _rhs293;
        }
        let _1872_v139;
        let _nw323 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1872_v139 = _nw323;
        let _1873_v140;
        _1873_v140 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1802_v79).f14);
        let _1874_v141;
        _1874_v141 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1873_v140).length),_this.f7);
        let _1875_v142;
        _1875_v142 = _dafny.Map.Empty.slice().updateUnsafe(_1874_v141,((_this.f7) ? (_1872_v139) : (_1872_v139)));
        _1872_v139 = (((_1875_v142).contains(_1874_v141)) ? ((_1875_v142).get(_1874_v141)) : (_1872_v139));
        (_this).f7 = ((_this.f7) ? ((_1802_v79).fm8(_1745_v30, globalState)) : (_this.f7));
      }
      let _1876_v143;
      let _nw324 = Array((new BigNumber(3)).toNumber()).fill(false);
      _1876_v143 = _nw324;
      let _index291 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      (_1876_v143)[_index291] = !(_this.f7);
      let _1877_v144;
      _1877_v144 = _module.D14.create_DC38(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),_dafny.Seq.UnicodeFromString("balalo")));
      let _1878_v145;
      _1878_v145 = _module.D15.create_DC43((_1802_v79).f14);
      let _pat_let_tv44 = _1801_v78;
      let _pat_let_tv45 = _1699_v2;
      let _pat_let_tv46 = _1745_v30;
      let _index292 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _rhs294 = !(_this.f7);
      let _rhs295 = function (_source13) {
        if (_source13.is_DC39) {
          let _1879___mcc_h1 = (_source13).cf63;
          let _1880___mcc_h2 = (_source13).cf64;
          let _1881_cf64 = _1880___mcc_h2;
          let _1882_cf63 = _1879___mcc_h1;
          if (_this.f7) {
            return _this.f7;
          } else {
            return _this.f7;
          }
        } else if (_source13.is_DC40) {
          let _1883___mcc_h3 = (_source13).cf65;
          let _1884___mcc_h4 = (_source13).cf66;
          let _1885___mcc_h5 = (_source13).cf67;
          let _1886_cf67 = _1885___mcc_h5;
          let _1887_cf66 = _1884___mcc_h4;
          let _1888_cf65 = _1883___mcc_h3;
          return _1886_cf67;
        } else if (_source13.is_DC38) {
          let _1889___mcc_h6 = (_source13).cf62;
          let _1890_cf62 = _1889___mcc_h6;
          return _this.f7;
        } else {
          let _1891___mcc_h7 = (_source13).cf68;
          let _1892_cf68 = _1891___mcc_h7;
          return _this.f7;
        }
      }(_1877_v144);
      let _rhs296 = new BigNumber(688);
      let _rhs297 = function (_source14) {
        if (_source14.is_DC43) {
          let _1893___mcc_h8 = (_source14).cf70;
          let _1894_cf70 = _1893___mcc_h8;
          return false;
        } else if (_source14.is_DC42) {
          let _1895___mcc_h9 = (_source14).cf69;
          let _1896_cf69 = _1895___mcc_h9;
          if (_this.f7) {
            return _this.f7;
          } else {
            return _this.f7;
          }
        } else {
          let _1897___mcc_h10 = (_source14).cf71;
          let _1898_cf71 = _1897___mcc_h10;
          return !(!_dafny.Seq.contains(_pat_let_tv44, _pat_let_tv45));
        }
      }(function (_pat_let55_0) {
        return function (_1899_dt__update__tmp_h1) {
          return function (_pat_let56_0) {
            return function (_1900_dt__update_hcf70_h0) {
              return _module.D15.create_DC43(_1900_dt__update_hcf70_h0);
            }(_pat_let56_0);
          }(_pat_let_tv46);
        }(_pat_let55_0);
      }(_1878_v145));
      let _lhs181 = _this;
      let _lhs182 = _this;
      let _lhs183 = _1876_v143;
      let _lhs184 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      _lhs181.f7 = _rhs294;
      _lhs182.f7 = _rhs295;
      r1 = _rhs296;
      _lhs183[_lhs184] = _rhs297;
      let _index293 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _index294 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _index295 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _rhs298 = ((_1745_v30).multipliedBy(new BigNumber(948))).minus(_1745_v30);
      let _rhs299 = (_1876_v143)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length))];
      let _rhs300 = (_1698_v1).fm32(!((_1876_v143)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length))]), function () {
        let _coll37 = new _dafny.Set();
        for (const _compr_37 of _dafny.IntegerRange(new BigNumber(258), new BigNumber(576))) {
          let _1901_v146 = _compr_37;
          if (((new BigNumber(258)).isLessThanOrEqualTo(_1901_v146)) && ((_1901_v146).isLessThan(new BigNumber(576)))) {
            _coll37.add(_module.__default.safeDivisionInt(_1901_v146, (_1802_v79).f14));
          }
        }
        return _coll37;
      }(), new BigNumber(831), _1745_v30, globalState);
      let _rhs301 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_1878_v145).dtor_cf70, _1745_v30));
      let _rhs302 = _this.f7;
      let _lhs185 = _1876_v143;
      let _lhs186 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _lhs187 = _1876_v143;
      let _lhs188 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      let _lhs189 = _1876_v143;
      let _lhs190 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1876_v143).length));
      _1745_v30 = _rhs298;
      _lhs185[_lhs186] = _rhs299;
      _lhs187[_lhs188] = _rhs300;
      r1 = _rhs301;
      _lhs189[_lhs190] = _rhs302;
      r0 = _module.__default.fm18(globalState);
      r1 = _1745_v30;
      let _1902_v147;
      _1902_v147 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1801_v78);
      r2 = (_1902_v147).update(false, _1801_v78);
      return [r0, r1, r2];
    }
    m10(p0, globalState) {
      let _this = this;
      let _1903_v0;
      _1903_v0 = new BigNumber(921);
      let _1904_v1;
      _1904_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f7);
      _1903_v0 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(p0, _1903_v0)).minus(new BigNumber((_1904_v1).length)));
      (_this).f7 = true;
      let _1905_v2;
      _1905_v2 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1906_v3;
      _1906_v3 = _dafny.Seq.UnicodeFromString("oislpn");
      let _1907_i0;
      _1907_i0 = _dafny.ZERO;
      L9: {
        while (_dafny.Seq.contains(_1906_v3, ((_this.f7) ? (_1905_v2) : (_module.__default.fm18(globalState))))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1907_i0)) {
              break L9;
            }
            _1907_i0 = (_1907_i0).plus(_dafny.ONE);
            let _1908_v4;
            _1908_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1906_v3,_this.f7);
            _1908_v4 = (_1908_v4).update(_dafny.Seq.Concat(_1906_v3, _1906_v3), _this.f7);
            (_this).f7 = _this.f7;
            let _1909_v5;
            let _nw325 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
            _1909_v5 = _nw325;
            let _1910_v6;
            _1910_v6 = _module.D0.create_DC0(new BigNumber(512));
            let _1911_v7;
            _1911_v7 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, _module.__default.fm1(_dafny.Seq.of(_1910_v6, _1910_v6, _1910_v6), _1903_v0, _1903_v0, globalState));
            let _index296 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1909_v5).length));
            (_1909_v5)[_index296] = new BigNumber((_1911_v7).cardinality());
            let _1912_v8;
            let _init62 = function (_1913_i1) {
              return _this.f7;
            };
            let _nw326 = Array((new BigNumber(17)).toNumber());
            for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw326.length); _i0_62++) {
              _nw326[_i0_62] = _init62(new BigNumber(_i0_62));
            }
            _1912_v8 = _nw326;
            let _1914_v9;
            _1914_v9 = _dafny.MultiSet.fromElements(_1912_v8, _1912_v8);
            let _index297 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1909_v5).length));
            (_1909_v5)[_index297] = new BigNumber(((_1914_v9).Difference((_1914_v9).update(_1912_v8, _module.__default.abs(p0)))).cardinality());
            let _1915_v10;
            _1915_v10 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_1916_v3) => function (_1917_i2) {
              return new BigNumber((_1916_v3).length);
            })(_1906_v3))).length));
            let _index298 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1909_v5).length));
            (_1909_v5)[_index298] = (((_1915_v10).contains(p0)) ? ((_1915_v10).get(p0)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(474)), function (_1918_i3) {
              return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("nnw")).length));
            })).length)));
          }
        }
      }
      let _1919_v11;
      _1919_v11 = _dafny.Set.fromElements(_1905_v2, _1905_v2);
      let _1920_v12;
      _1920_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1906_v3);
      let _1921_v13;
      _1921_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,true);
      let _1922_v15;
      _1922_v15 = _dafny.Set.fromElements(_1905_v2);
      let _source15 = (function () {
        let _coll38 = new _dafny.Set();
        for (const _compr_38 of ((((_1920_v12).contains((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) ? ((_1920_v12).get((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) : (_1906_v3))).Elements) {
          let _1923_v14 = _compr_38;
          if (_dafny.Seq.contains((((_1920_v12).contains((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) ? ((_1920_v12).get((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) : (_1906_v3)), _1923_v14)) {
            _coll38.add(_1923_v14);
          }
        }
        return _coll38;
      }()).Difference(_1922_v15);
      let _1924___mcc_h0 = _source15;
      let _1925_cf47 = _1924___mcc_h0;
      _1906_v3 = _dafny.Seq.UnicodeFromString("f");
      let _1926_v16;
      _1926_v16 = _module.D1.create_DC3(_dafny.Set.fromElements(_this.f7));
      let _1927_v17;
      _1927_v17 = _module.D0.create_DC0(new BigNumber(-529));
      let _1928_v18;
      _1928_v18 = _module.D8.create_DC22(p0, _1926_v16, _1927_v17);
      let _1929_v19;
      _1929_v19 = _dafny.Seq.of(_1928_v18, _1928_v18, _1928_v18);
      _1929_v19 = _1929_v19;
      (_this).f7 = _this.f7;
      _1903_v0 = p0;
      if ((_this.f7) && (_this.f7)) {
        let _1930_v21;
        let _nw327 = new _module.C0();
        _nw327.__ctor((_dafny.Set.fromElements(_1905_v2, _1905_v2)).IsProperSubsetOf(function () {
          let _coll39 = new _dafny.Set();
          for (const _compr_39 of (_1906_v3).Elements) {
            let _1931_v20 = _compr_39;
            if (_dafny.Seq.contains(_1906_v3, _1931_v20)) {
              _coll39.add(_1931_v20);
            }
          }
          return _coll39;
        }()), _this.f7);
        _1930_v21 = _nw327;
        let _1932_v22;
        _1932_v22 = _module.D6.create_DC16(false);
        if ((_1932_v22).dtor_cf30) {
          let _1933_v23;
          _1933_v23 = _dafny.Seq.of(p0);
          let _1934_v24;
          _1934_v24 = _dafny.Set.fromElements(_1933_v23);
          let _1935_v25;
          _1935_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,_1934_v24);
          let _1936_v26;
          _1936_v26 = _module.D0.create_DC0(p0);
          let _1937_v27;
          _1937_v27 = _dafny.Seq.of(_1936_v26, _1936_v26);
          let _1938_v28;
          _1938_v28 = _dafny.MultiSet.fromElements(_1930_v21.f13, _1930_v21.f13);
          let _1939_v29;
          _1939_v29 = _dafny.Set.fromElements(!((_1930_v21).f12));
          let _1940_v30;
          _1940_v30 = _module.D1.create_DC3(_1939_v29);
          let _1941_v31;
          _1941_v31 = _module.D8.create_DC22(_1903_v0, _1940_v30, _1936_v26);
          let _rhs303 = ((((((_1921_v13).contains((_1930_v21).f12)) ? ((_1921_v13).get((_1930_v21).f12)) : (true))) ? (_1934_v24) : ((((_1935_v25).contains(_1930_v21.f13)) ? ((_1935_v25).get(_1930_v21.f13)) : (_1934_v24))))).IsProperSubsetOf(((_1930_v21.f13) ? (_dafny.Set.fromElements(_1933_v23)) : (_1934_v24)));
          let _rhs304 = _module.__default.fm1(_1937_v27, new BigNumber((_1938_v28).cardinality()), (_1941_v31).dtor_cf41, globalState);
          let _rhs305 = new BigNumber((_dafny.Seq.Concat(_1933_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_1942_p0) => function (_1943_i4) {
            return (_dafny.ZERO).minus((_dafny.ZERO).minus(_1942_p0));
          })(p0)))).length);
          let _lhs191 = _1930_v21;
          let _lhs192 = _1930_v21;
          _lhs191.f13 = _rhs303;
          _lhs192.f13 = _rhs304;
          _1903_v0 = _rhs305;
          let _1944_v32;
          let _nw328 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1944_v32 = _nw328;
          let _index299 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length));
          (_1944_v32)[_index299] = _module.__default.safeModuloInt(p0, _1903_v0);
          let _index300 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length));
          (_1944_v32)[_index300] = _1903_v0;
          let _1945_v33;
          _1945_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1905_v2,_1944_v32);
          let _1946_v34;
          let _nw329 = Array((new BigNumber(27)).toNumber()).fill([]);
          _1946_v34 = _nw329;
          let _1947_v35;
          let _init63 = ((_1948_v23) => function (_1949_i5) {
            return _1948_v23;
          })(_1933_v23);
          let _nw330 = Array((new BigNumber(28)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw330.length); _i0_63++) {
            _nw330[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _1947_v35 = _nw330;
          let _index301 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_1946_v34).length));
          (_1946_v34)[_index301] = _1947_v35;
          let _index302 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_1946_v34).length));
          let _rhs306 = _module.__default.safeModuloInt((_1944_v32)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length))], (((_1930_v21).f12) ? (p0) : (new BigNumber((_1921_v13).length))));
          let _rhs307 = ((_1945_v33).Merge(_1945_v33)).Merge(_1945_v33);
          let _rhs308 = !(_1930_v21.f13);
          let _rhs309 = _module.__default.fm18(globalState);
          let _rhs310 = _1947_v35;
          let _lhs193 = _1930_v21;
          let _lhs194 = _1946_v34;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_1946_v34).length));
          _1903_v0 = _rhs306;
          _1945_v33 = _rhs307;
          _lhs193.f13 = _rhs308;
          _1905_v2 = _rhs309;
          _lhs194[_lhs195] = _rhs310;
          let _1950_v36;
          _1950_v36 = _dafny.Seq.of(!(!(_this.f7)), _this.f7);
          let _1951_v37;
          _1951_v37 = _dafny.Seq.of((_1950_v36)[_module.__default.safeIndex(p0, new BigNumber((_1950_v36).length))]);
          let _1952_v38;
          _1952_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1930_v21.f13,p0);
          let _1953_v39;
          _1953_v39 = _dafny.Seq.of(_1952_v38, _1952_v38);
          let _1954_v40;
          _1954_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1906_v3).length),_1905_v2);
          let _1955_v41;
          _1955_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1953_v39,_dafny.Seq.update(_dafny.Seq.Concat(_1951_v37, _dafny.Seq.of(_1930_v21.f13, (_1930_v21).f12, (_1930_v21).f12, _this.f7, (_1930_v21).f12)), _module.__default.safeIndex(new BigNumber((_1954_v40).length), new BigNumber((_dafny.Seq.Concat(_1951_v37, _dafny.Seq.of(_1930_v21.f13, (_1930_v21).f12, (_1930_v21).f12, _this.f7, (_1930_v21).f12))).length)), (_1930_v21).f12));
          _1951_v37 = (((_1955_v41).contains(_1953_v39)) ? ((_1955_v41).get(_1953_v39)) : (_1951_v37));
          let _index303 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length));
          let _rhs311 = ((_1944_v32)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length))]).minus(_1903_v0);
          let _rhs312 = _1930_v21.f13;
          let _lhs196 = _1944_v32;
          let _lhs197 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1944_v32).length));
          let _lhs198 = _1930_v21;
          _lhs196[_lhs197] = _rhs311;
          _lhs198.f13 = _rhs312;
        } else {
          let _1956_v42;
          let _nw331 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1956_v42 = _nw331;
          let _1957_v43;
          let _nw332 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1957_v43 = _nw332;
          let _index304 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length));
          (_1956_v42)[_index304] = _1957_v43;
          let _1958_v44;
          _1958_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1903_v0,_1903_v0);
          let _1959_v45;
          _1959_v45 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          let _1960_v46;
          _1960_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(_1930_v21.f13),(_1930_v21).f12);
          let _index305 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length));
          let _nw333 = Array((new BigNumber(21)).toNumber());
          _nw333[(_dafny.ZERO).toNumber()] = p0;
          _nw333[(_dafny.ONE).toNumber()] = ((_1930_v21.f13) ? (_1903_v0) : (p0));
          _nw333[(new BigNumber(2)).toNumber()] = new BigNumber(-965);
          _nw333[(new BigNumber(3)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(4)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(5)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(6)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(7)).toNumber()] = _module.__default.fm0(globalState);
          _nw333[(new BigNumber(8)).toNumber()] = p0;
          _nw333[(new BigNumber(9)).toNumber()] = p0;
          _nw333[(new BigNumber(10)).toNumber()] = p0;
          _nw333[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(857), new BigNumber((_1958_v44).length)));
          _nw333[(new BigNumber(12)).toNumber()] = p0;
          _nw333[(new BigNumber(13)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(14)).toNumber()] = (p0).minus(new BigNumber((_1958_v44).length));
          _nw333[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((p0).multipliedBy(new BigNumber(((_module.__default.fm53(new BigNumber((_1959_v45).length), _1930_v21.f13, globalState)).dtor_cf15).length)));
          _nw333[(new BigNumber(16)).toNumber()] = p0;
          _nw333[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1903_v0, p0));
          _nw333[(new BigNumber(18)).toNumber()] = (_1903_v0).multipliedBy(new BigNumber(101));
          _nw333[(new BigNumber(19)).toNumber()] = _1903_v0;
          _nw333[(new BigNumber(20)).toNumber()] = (new BigNumber(507)).multipliedBy(new BigNumber(((_1960_v46).update(_1932_v22, (_1930_v21).f12)).length));
          (_1956_v42)[_index305] = _nw333;
          let _1961_v47;
          let _nw334 = Array((new BigNumber(16)).toNumber());
          _nw334[(_dafny.ZERO).toNumber()] = !(_1930_v21.f13);
          _nw334[(_dafny.ONE).toNumber()] = _this.f7;
          _nw334[(new BigNumber(2)).toNumber()] = (_1930_v21).f12;
          _nw334[(new BigNumber(3)).toNumber()] = _this.f7;
          _nw334[(new BigNumber(4)).toNumber()] = _1930_v21.f13;
          _nw334[(new BigNumber(5)).toNumber()] = _1930_v21.f13;
          _nw334[(new BigNumber(6)).toNumber()] = (_1930_v21).f12;
          _nw334[(new BigNumber(7)).toNumber()] = _this.f7;
          _nw334[(new BigNumber(8)).toNumber()] = (_1930_v21).f12;
          _nw334[(new BigNumber(9)).toNumber()] = (_1930_v21).f12;
          _nw334[(new BigNumber(10)).toNumber()] = _this.f7;
          _nw334[(new BigNumber(11)).toNumber()] = _1930_v21.f13;
          _nw334[(new BigNumber(12)).toNumber()] = _1930_v21.f13;
          _nw334[(new BigNumber(13)).toNumber()] = _this.f7;
          _nw334[(new BigNumber(14)).toNumber()] = (_1930_v21).f12;
          _nw334[(new BigNumber(15)).toNumber()] = (_1930_v21).f12;
          _1961_v47 = _nw334;
          let _1962_v48;
          _1962_v48 = _dafny.Set.fromElements(_1961_v47);
          let _1963_v49;
          _1963_v49 = _dafny.Seq.of(_1930_v21.f13, (_1930_v21).f12);
          let _1964_v50;
          _1964_v50 = _module.D13.create_DC36(_1905_v2, _1963_v49, true);
          let _1965_v51;
          let _nw335 = new _module.C2();
          _nw335.__ctor(_1962_v48, (_1964_v50).dtor_cf58, _this.f7);
          _1965_v51 = _nw335;
          let _arr4 = (_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))];
          let _index306 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]).length));
          _arr4[_index306] = _1903_v0;
          let _arr5 = (_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))];
          let _index307 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]).length));
          _arr5[_index307] = _1903_v0;
          let _arr6 = (_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))];
          let _index308 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]).length));
          _arr6[_index308] = ((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))])[_module.__default.safeIndex(new BigNumber(695), new BigNumber(((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]).length))];
          let _1966_v52;
          _1966_v52 = _module.D9.create_DC24();
          let _arr7 = (_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))];
          let _index309 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]).length));
          _arr7[_index309] = new BigNumber(((_module.D16.create_DC45(_dafny.Map.Empty.slice().updateUnsafe(_1966_v52,(_1956_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1956_v42).length))]))).dtor_cf72).length);
        }
        let _1967_v53;
        let _init64 = ((_1968_p0, _1969_v0) => function (_1970_i6) {
          return (_module.D12.create_DC30((_dafny.ZERO).minus(_1968_p0), false, _dafny.Seq.of(_1969_v0), _1968_p0)).dtor_cf51;
        })(p0, _1903_v0);
        let _nw336 = Array((new BigNumber(5)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw336.length); _i0_64++) {
          _nw336[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _1967_v53 = _nw336;
        let _1971_v54;
        _1971_v54 = _dafny.Seq.of(_1903_v0);
        let _index310 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1967_v53).length));
        (_1967_v53)[_index310] = _dafny.Seq.Concat(_dafny.Seq.of(p0), _1971_v54);
        let _1972_v55;
        _1972_v55 = _dafny.MultiSet.fromElements(_this.f7);
        let _index311 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1967_v53).length));
        (_1967_v53)[_index311] = _dafny.Seq.Concat(_dafny.Seq.update(_1971_v54, _module.__default.safeIndex(new BigNumber((_1906_v3).length), new BigNumber((_1971_v54).length)), new BigNumber((_1972_v55).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(983)), ((_1973_p0) => function (_1974_i7) {
          return (_dafny.ZERO).minus(_1973_p0);
        })(p0)));
        _1972_v55 = _1972_v55;
        let _1975_v56;
        let _nw337 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1975_v56 = _nw337;
        let _1976_v57;
        _1976_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-271),_1975_v56);
        let _1977_v58;
        let _nw338 = Array((new BigNumber(3)).toNumber());
        _nw338[(_dafny.ZERO).toNumber()] = _1930_v21.f13;
        _nw338[(_dafny.ONE).toNumber()] = (_1903_v0).isLessThanOrEqualTo(new BigNumber((_1976_v57).length));
        _nw338[(new BigNumber(2)).toNumber()] = (p0).isLessThan(new BigNumber(17));
        _1977_v58 = _nw338;
        let _index312 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1975_v56).length));
        (_1975_v56)[_index312] = _1930_v21.f13;
        let _1978_v59;
        _1978_v59 = _dafny.Set.fromElements((_1930_v21).f12, false);
        let _1979_v60;
        _1979_v60 = _module.D1.create_DC3(_1978_v59);
        let _1980_v61;
        _1980_v61 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.update(_1906_v3, _module.__default.safeIndex(p0, new BigNumber((_1906_v3).length)), _1905_v2)).length));
        let _1981_v62;
        _1981_v62 = _module.D8.create_DC22(new BigNumber(446), _1979_v60, _1980_v61);
        let _1982_v63;
        _1982_v63 = _dafny.Seq.of(true, _module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(578)), ((_1983_v61) => function (_1984_i8) {
          return _1983_v61;
        })(_1980_v61)), _1903_v0, p0, globalState), _1930_v21.f13);
        let _index313 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1975_v56).length));
        let _rhs313 = (_1930_v21).f12;
        let _rhs314 = (_1982_v63)[_module.__default.safeIndex(new BigNumber((_1906_v3).length), new BigNumber((_1982_v63).length))];
        let _rhs315 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1906_v3, _1906_v3), _1906_v3);
        let _rhs316 = _1905_v2;
        let _rhs317 = _1981_v62;
        let _lhs199 = _1975_v56;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1975_v56).length));
        let _lhs201 = _1930_v21;
        let _lhs202 = _1930_v21;
        _lhs199[_lhs200] = _rhs313;
        _lhs201.f13 = _rhs314;
        _lhs202.f13 = _rhs315;
        _1905_v2 = _rhs316;
        _1981_v62 = _rhs317;
      } else {
        if (((_this.f7) ? (true) : (_this.f7))) {
          let _1985_v64;
          let _init65 = function (_1986_i9) {
            return _dafny.MultiSet.fromElements(_this.f7, !(_this.f7));
          };
          let _nw339 = Array((new BigNumber(15)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw339.length); _i0_65++) {
            _nw339[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _1985_v64 = _nw339;
          _1985_v64 = _1985_v64;
          let _1987_v65;
          _1987_v65 = _module.D0.create_DC0(p0);
          let _1988_v66;
          _1988_v66 = _dafny.Seq.of(_module.D0.create_DC0(p0), _1987_v65, _1987_v65);
          let _1989_v67;
          _1989_v67 = _dafny.Seq.of(_this.f7);
          let _1990_v68;
          _1990_v68 = _dafny.Seq.of(new BigNumber((_1989_v67).length));
          let _1991_v69;
          let _nw340 = new _module.C1();
          _nw340.__ctor(_module.__default.fm1(_1988_v66, new BigNumber((_1906_v3).length), (_dafny.ZERO).minus((_1990_v68)[_module.__default.safeIndex(p0, new BigNumber((_1990_v68).length))]), globalState));
          _1991_v69 = _nw340;
          let _1992_v70;
          let _nw341 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1992_v70 = _nw341;
          let _index314 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1992_v70).length));
          (_1992_v70)[_index314] = _this.f7;
          let _index315 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1992_v70).length));
          (_1992_v70)[_index315] = _this.f7;
          let _1993_v71;
          _1993_v71 = _dafny.Map.Empty.slice().updateUnsafe((_1992_v70)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_1992_v70).length))],(_1903_v0).minus(_1903_v0));
          _1993_v71 = _module.__default.fm43(_1905_v2, _this.f7, _this.f7, globalState);
          let _1994_v72;
          _1994_v72 = _dafny.MultiSet.fromElements(_1903_v0);
          let _1995_v73;
          let _nw342 = new _module.C3();
          _nw342.__ctor((_1903_v0).minus(new BigNumber(151)), (_1994_v72).IsSubsetOf(_1994_v72));
          _1995_v73 = _nw342;
        } else {
          let _1996_v74;
          _1996_v74 = _module.D6.create_DC17(_this.f7, _1903_v0, _1905_v2, _1905_v2);
          let _1997_v75;
          _1997_v75 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),(_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState))));
          let _1998_v77;
          _1998_v77 = _dafny.Set.fromElements(new BigNumber((_1997_v75).length));
          let _1999_v78;
          let _nw343 = Array((new BigNumber(16)).toNumber());
          _nw343[(_dafny.ZERO).toNumber()] = p0;
          _nw343[(_dafny.ONE).toNumber()] = new BigNumber(-252);
          _nw343[(new BigNumber(2)).toNumber()] = _1903_v0;
          _nw343[(new BigNumber(3)).toNumber()] = (new BigNumber(250)).plus((_1996_v74).dtor_cf32);
          _nw343[(new BigNumber(4)).toNumber()] = _1903_v0;
          _nw343[(new BigNumber(5)).toNumber()] = (_1903_v0).plus(_1903_v0);
          _nw343[(new BigNumber(6)).toNumber()] = p0;
          _nw343[(new BigNumber(7)).toNumber()] = _1903_v0;
          _nw343[(new BigNumber(8)).toNumber()] = new BigNumber((_1904_v1).length);
          _nw343[(new BigNumber(9)).toNumber()] = _1903_v0;
          _nw343[(new BigNumber(10)).toNumber()] = (((_1997_v75).contains(p0)) ? ((_1997_v75).get(p0)) : (new BigNumber((function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_1998_v77).Elements) {
              let _2000_v76 = _compr_40;
              if ((_1998_v77).contains(_2000_v76)) {
                _coll40.push([_module.__default.safeModuloInt(_2000_v76, new BigNumber((_1906_v3).length)),_this.f7]);
              }
            }
            return _coll40;
          }()).length)));
          _nw343[(new BigNumber(11)).toNumber()] = p0;
          _nw343[(new BigNumber(12)).toNumber()] = _1903_v0;
          _nw343[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_module.__default.fm0(globalState)).plus(new BigNumber((_1906_v3).length)));
          _nw343[(new BigNumber(14)).toNumber()] = (_1903_v0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), ((_2001_p0) => function (_2002_i10) {
            return _2001_p0;
          })(p0))).length));
          _nw343[(new BigNumber(15)).toNumber()] = p0;
          _1999_v78 = _nw343;
          let _index316 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length));
          (_1999_v78)[_index316] = p0;
          let _index317 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length));
          (_1999_v78)[_index317] = (new BigNumber(276)).plus(p0);
          let _2003_v79;
          _2003_v79 = _dafny.Seq.of(_this.f7, false);
          let _2004_v80;
          _2004_v80 = _dafny.Seq.of(_this.f7, _this.f7, (_2003_v79)[_module.__default.safeIndex(_1903_v0, new BigNumber((_2003_v79).length))], _this.f7, _this.f7);
          let _2005_v81;
          _2005_v81 = _dafny.Seq.of(new BigNumber((_2004_v80).length));
          let _2006_v82;
          _2006_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1999_v78,_1903_v0);
          let _index318 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length));
          (_1999_v78)[_index318] = (_2005_v81)[_module.__default.safeIndex(_module.__default.safeModuloInt((((_2006_v82).contains(_1999_v78)) ? ((_2006_v82).get(_1999_v78)) : ((_1999_v78)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length))])), (_2005_v81)[_module.__default.safeIndex(_1903_v0, new BigNumber((_2005_v81).length))]), new BigNumber((_2005_v81).length))];
          let _2007_v83;
          _2007_v83 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("p"), _dafny.Seq.update(_dafny.Seq.Concat((((_1920_v12).contains((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) ? ((_1920_v12).get((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) : (_1906_v3)), _1906_v3), _module.__default.safeIndex((_1999_v78)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length))], new BigNumber((_dafny.Seq.Concat((((_1920_v12).contains((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) ? ((_1920_v12).get((((_1921_v13).contains(_this.f7)) ? ((_1921_v13).get(_this.f7)) : (_this.f7)))) : (_1906_v3)), _1906_v3)).length)), _1905_v2), _1906_v3, _dafny.Seq.UnicodeFromString("volltvv"), _1906_v3);
          let _2008_v84;
          _2008_v84 = _module.D9.create_DC26(_2005_v81);
          let _2009_v85;
          _2009_v85 = _module.D12.create_DC31(_2008_v84, _this.f7);
          let _2010_v86;
          _2010_v86 = _dafny.Map.Empty.slice().updateUnsafe(true,_2003_v79);
          _2007_v83 = _module.__default.fm3(_this.f7, _1903_v0, new BigNumber((_module.__default.fm49((_2009_v85).dtor_cf54, _2010_v86, _this.f7, globalState)).cardinality()), ((_this.f7) ? ((_1999_v78)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_1999_v78).length))]) : (_1903_v0)), globalState);
          let _2011_v87;
          _2011_v87 = _dafny.MultiSet.fromElements((p0).multipliedBy(p0));
          _2011_v87 = _2011_v87;
          (_this).f7 = _this.f7;
        }
        (_this).f7 = true;
        let _2012_v88;
        _2012_v88 = _dafny.Seq.of(new BigNumber(382), (_1903_v0).plus(_1903_v0), p0);
        _1903_v0 = (_2012_v88)[_module.__default.safeIndex(p0, new BigNumber((_2012_v88).length))];
        _1906_v3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), ((_2013_v2) => function (_2014_i11) {
          return _2013_v2;
        })(_1905_v2));
        let _2015_v89;
        _2015_v89 = _dafny.Set.fromElements(_this.f7, _this.f7, _this.f7, _this.f7, _this.f7);
        let _2016_v90;
        _2016_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2015_v89).Union(_2015_v89)).length),_1906_v3);
        _2016_v90 = (_2016_v90).update(_module.__default.safeDivisionInt(p0, p0), _1906_v3);
      }
      let _2017_v91;
      let _nw344 = new _module.C3();
      _nw344.__ctor(p0, _this.f7);
      _2017_v91 = _nw344;
      let _2018_v92;
      _2018_v92 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2017_v91);
      let _hi8 = ((_this.f7) ? (new BigNumber(-406)) : (new BigNumber((_2018_v92).length)));
      for (let _2019_i12 = (_this).fm15(globalState); _2019_i12.isLessThan(_hi8); _2019_i12 = _2019_i12.plus(_dafny.ONE)) {
        let _2020_v93;
        let _init66 = function (_2021_i13) {
          return _this.f7;
        };
        let _nw345 = Array((new BigNumber(13)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw345.length); _i0_66++) {
          _nw345[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2020_v93 = _nw345;
        let _index319 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2020_v93).length));
        (_2020_v93)[_index319] = _this.f7;
        let _2022_v94;
        let _nw346 = Array((_dafny.ONE).toNumber());
        _nw346[(_dafny.ZERO).toNumber()] = _1903_v0;
        _2022_v94 = _nw346;
        let _2023_v95;
        _2023_v95 = _dafny.Set.fromElements(_2022_v94);
        let _2024_v96;
        _2024_v96 = _module.D0.create_DC2(_this.f7, p0, _2023_v95, p0);
        let _2025_v97;
        _2025_v97 = _dafny.MultiSet.fromElements(_this.f7);
        let _index320 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2020_v93).length));
        (_2020_v93)[_index320] = ((_2024_v96).dtor_cf3).isEqualTo(new BigNumber(((_2025_v97).Union(_2025_v97)).cardinality()));
        let _index321 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2020_v93).length));
        let _rhs318 = false;
        let _rhs319 = (_2020_v93)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_2020_v93).length))];
        let _lhs203 = _2020_v93;
        let _lhs204 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2020_v93).length));
        let _lhs205 = _this;
        _lhs203[_lhs204] = _rhs318;
        _lhs205.f7 = _rhs319;
        _1903_v0 = _1903_v0;
        let _2026_v98;
        _2026_v98 = _dafny.Seq.of(_1904_v1);
        (_this).f7 = (new BigNumber(777)).isLessThanOrEqualTo(new BigNumber(((_2026_v98)[_module.__default.safeIndex((_2017_v91).f14, new BigNumber((_2026_v98).length))]).length));
      }
      return;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10) {
      let _this = this;
      (_this)._f10 = f10;
      return;
    }
    m7(globalState) {
      let _this = this;
      let _2027_v0;
      _2027_v0 = new _dafny.CodePoint('n'.codePointAt(0));
      let _2028_v1;
      _2028_v1 = _dafny.Seq.UnicodeFromString("xb");
      let _2029_v2;
      _2029_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2027_v0,new BigNumber((_2028_v1).length));
      let _2030_v3;
      _2030_v3 = new BigNumber(68);
      let _hi9 = _2030_v3;
      for (let _2031_i0 = (((_2029_v2).contains(_2027_v0)) ? ((_2029_v2).get(_2027_v0)) : (_2030_v3)); _2031_i0.isLessThan(_hi9); _2031_i0 = _2031_i0.plus(_dafny.ONE)) {
        if ((_this).f10) {
          let _2032_v4;
          _2032_v4 = true;
          _2032_v4 = (_this).f10;
          let _2033_v5;
          _2033_v5 = _dafny.MultiSet.fromElements(_2031_i0);
          let _rhs320 = (_this).f10;
          let _rhs321 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), ((_2034_v0) => function (_2035_i1) {
            return _2034_v0;
          })(_2027_v0));
          let _rhs322 = _2033_v5;
          _2032_v4 = _rhs320;
          _2028_v1 = _rhs321;
          _2033_v5 = _rhs322;
          let _2036_v6;
          _2036_v6 = _dafny.MultiSet.fromElements((_2031_i0).isLessThanOrEqualTo(_module.__default.fm0(globalState)), !(((_this).f10) && ((_this).f10)), !(_2032_v4), (_this).f10);
          let _2037_v8;
          _2037_v8 = _dafny.Set.fromElements(_2030_v3);
          let _2038_v9;
          _2038_v9 = _dafny.Seq.of(_2032_v4, _2032_v4, _2032_v4);
          let _2039_v10;
          _2039_v10 = _module.D0.create_DC1(new BigNumber((_2028_v1).length));
          let _2040_v11;
          _2040_v11 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let57_0) {
            return function (_2041_dt__update__tmp_h0) {
              return function (_pat_let58_0) {
                return function (_2042_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_2042_dt__update_hcf1_h0);
                }(_pat_let58_0);
              }(new BigNumber(-760));
            }(_pat_let57_0);
          }(_2039_v10)).dtor_cf1,_2030_v3);
          let _rhs323 = (function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_2037_v8).Elements) {
              let _2043_v7 = _compr_41;
              if ((_2037_v8).contains(_2043_v7)) {
                _coll41.push([_module.__default.safeModuloInt(_2043_v7, _2031_i0),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2028_v1).length),_2032_v4)]);
              }
            }
            return _coll41;
          }()).contains(new BigNumber(-407));
          let _rhs324 = _2036_v6;
          let _rhs325 = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(_2032_v4, (_this).f10, (_2038_v9)[_module.__default.safeIndex(new BigNumber((_2040_v11).length), new BigNumber((_2038_v9).length))]), _2038_v9, _2038_v9)).length));
          _2032_v4 = _rhs323;
          _2036_v6 = _rhs324;
          _2030_v3 = _rhs325;
          _2032_v4 = _2032_v4;
          let _2044_v12;
          _2044_v12 = _module.D1.create_DC4(_2032_v4, !(_2032_v4));
          let _2045_v13;
          _2045_v13 = _dafny.Seq.of(function (_pat_let59_0) {
            return function (_2046_dt__update__tmp_h1) {
              return function (_pat_let60_0) {
                return function (_2047_dt__update_hcf8_h0) {
                  return _module.D1.create_DC4((_2046_dt__update__tmp_h1).dtor_cf7, _2047_dt__update_hcf8_h0);
                }(_pat_let60_0);
              }(false);
            }(_pat_let59_0);
          }(_2044_v12));
          _2045_v13 = _dafny.Seq.Concat(_2045_v13, _dafny.Seq.Concat(_2045_v13, _2045_v13));
        } else {
          let _2048_v14;
          let _nw347 = Array((new BigNumber(2)).toNumber()).fill(false);
          _2048_v14 = _nw347;
          let _2049_v15;
          _2049_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2048_v14,_2031_i0);
          _2049_v15 = _2049_v15;
          let _2050_v16;
          _2050_v16 = true;
          _2050_v16 = (_this).f10;
          let _2051_v17;
          _2051_v17 = _module.D0.create_DC0(_2030_v3);
          let _2052_v18;
          _2052_v18 = _dafny.Seq.of(_2051_v17);
          _2050_v16 = _module.__default.fm1(_2052_v18, _2031_i0, ((_dafny.ZERO).minus(_2030_v3)).plus(_2030_v3), globalState);
          let _index322 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_2048_v14).length));
          (_2048_v14)[_index322] = !(_2050_v16) || (!(_2050_v16));
          let _index323 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_2048_v14).length));
          let _rhs326 = _2050_v16;
          let _rhs327 = true;
          let _lhs206 = _2048_v14;
          let _lhs207 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_2048_v14).length));
          _lhs206[_lhs207] = _rhs326;
          _2050_v16 = _rhs327;
          let _2053_v19;
          _2053_v19 = _dafny.Set.fromElements(_2031_i0, _module.__default.fm0(globalState), _2030_v3);
          let _2054_v20;
          _2054_v20 = _module.D2.create_DC7(_2053_v19);
          _2053_v19 = (function (_pat_let61_0) {
            return function (_2055_dt__update__tmp_h2) {
              return function (_pat_let62_0) {
                return function (_2057_dt__update_hcf13_h0) {
                  return _module.D2.create_DC7(_2057_dt__update_hcf13_h0);
                }(_pat_let62_0);
              }(function () {
                let _coll42 = new _dafny.Set();
                for (const _compr_42 of _dafny.IntegerRange(new BigNumber(133), new BigNumber(-938))) {
                  let _2056_v21 = _compr_42;
                  if (((new BigNumber(133)).isLessThanOrEqualTo(_2056_v21)) && ((_2056_v21).isLessThan(new BigNumber(-938)))) {
                    _coll42.add(_module.__default.safeDivisionInt(_2056_v21, new BigNumber(462)));
                  }
                }
                return _coll42;
              }());
            }(_pat_let61_0);
          }(_2054_v20)).dtor_cf13;
        }
        let _2058_v22;
        _2058_v22 = _module.D1.create_DC4((_this).f10, (_this).f10);
        let _2059_v23;
        let _nw348 = Array((new BigNumber(7)).toNumber());
        _nw348[(_dafny.ZERO).toNumber()] = (_this).f10;
        _nw348[(_dafny.ONE).toNumber()] = false;
        _nw348[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw348[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw348[(new BigNumber(4)).toNumber()] = (_this).f10;
        _nw348[(new BigNumber(5)).toNumber()] = (_this).f10;
        _nw348[(new BigNumber(6)).toNumber()] = ((_2058_v22).dtor_cf8) && ((_this).f10);
        _2059_v23 = _nw348;
        let _index324 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_2059_v23).length));
        (_2059_v23)[_index324] = false;
        let _2060_v24;
        _2060_v24 = false;
        let _2061_v25;
        _2061_v25 = _module.D0.create_DC0(_2030_v3);
        let _2062_v26;
        _2062_v26 = _dafny.Seq.of(_2061_v25, _2061_v25);
        let _2063_v27;
        _2063_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2028_v1,_module.__default.fm1(_2062_v26, _2031_i0, _2030_v3, globalState));
        let _2064_v28;
        _2064_v28 = _dafny.Seq.of(_2028_v1, _2028_v1);
        let _2065_v29;
        _2065_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2031_i0,_2060_v24);
        let _2066_v30;
        _2066_v30 = _dafny.Seq.of(_2031_i0);
        let _index325 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_2059_v23).length));
        let _rhs328 = (_module.D2.create_DC8((_this).f10, _2028_v1)).dtor_cf14;
        let _rhs329 = (((_2063_v27).contains(_dafny.Seq.UnicodeFromString("akjinxm"))) ? ((_2063_v27).get(_dafny.Seq.UnicodeFromString("akjinxm"))) : ((!(_2060_v24)) || ((_this).f10)));
        let _rhs330 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2028_v1, _2028_v1), ((_2060_v24) ? (_2028_v1) : ((_2064_v28)[_module.__default.safeIndex(new BigNumber((_2065_v29).length), new BigNumber((_2064_v28).length))])));
        let _rhs331 = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
        let _rhs332 = !((_this).f10) || (_module.__default.fm1(_2062_v26, new BigNumber((_2066_v30).length), _2030_v3, globalState));
        let _lhs208 = _2059_v23;
        let _lhs209 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_2059_v23).length));
        _lhs208[_lhs209] = _rhs328;
        _2060_v24 = _rhs329;
        _2028_v1 = _rhs330;
        _2030_v3 = _rhs331;
        _2060_v24 = _rhs332;
        let _2067_v31;
        _2067_v31 = _module.D1.create_DC4((_this).f10, _2060_v24);
        let _2068_v32;
        _2068_v32 = _module.D1.create_DC6(_2067_v31);
        let _2069_v33;
        _2069_v33 = _dafny.MultiSet.fromElements(_2030_v3);
        let _2070_v34;
        _2070_v34 = _dafny.Set.fromElements((_2059_v23)[_module.__default.safeIndex(new BigNumber(584), new BigNumber((_2059_v23).length))], (_this).f10);
        let _2071_v35;
        _2071_v35 = _dafny.Seq.of(_module.D1.create_DC6(_2067_v31), _module.__default.fm14(_2031_i0, new BigNumber((_2069_v33).cardinality()), new BigNumber((_2070_v34).length), _module.__default.fm1(_dafny.Seq.of(_2061_v25, _2061_v25, _module.D0.create_DC0(_2031_i0), _2061_v25, _2061_v25), new BigNumber((_2028_v1).length), _2030_v3, globalState), globalState), _2067_v31, _2067_v31, _2067_v31);
        let _pat_let_tv47 = _2067_v31;
        let _2072_v36;
        let _nw349 = Array((new BigNumber(19)).toNumber());
        _nw349[(_dafny.ZERO).toNumber()] = _2068_v32;
        _nw349[(_dafny.ONE).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(2)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(3)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(4)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(5)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(6)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(7)).toNumber()] = _module.D1.create_DC6(_2067_v31);
        _nw349[(new BigNumber(8)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(9)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(10)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(11)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(12)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(13)).toNumber()] = _module.D1.create_DC6((_2071_v35)[_module.__default.safeIndex(_2031_i0, new BigNumber((_2071_v35).length))]);
        _nw349[(new BigNumber(14)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(15)).toNumber()] = _module.D1.create_DC6(_2067_v31);
        _nw349[(new BigNumber(16)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(17)).toNumber()] = _2068_v32;
        _nw349[(new BigNumber(18)).toNumber()] = function (_pat_let63_0) {
          return function (_2073_dt__update__tmp_h3) {
            return function (_pat_let64_0) {
              return function (_2074_dt__update_hcf12_h0) {
                return _module.D1.create_DC6(_2074_dt__update_hcf12_h0);
              }(_pat_let64_0);
            }(_pat_let_tv47);
          }(_pat_let63_0);
        }(_module.D1.create_DC6(_2067_v31));
        _2072_v36 = _nw349;
        let _2075_v37;
        _2075_v37 = _dafny.Set.fromElements(_2072_v36);
        _2060_v24 = (_2075_v37).IsProperSubsetOf(_2075_v37);
        let _2076_v38;
        _2076_v38 = _dafny.Seq.of((((_2059_v23)[_module.__default.safeIndex(new BigNumber(584), new BigNumber((_2059_v23).length))]) ? (_2066_v30) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), ((_2077_i0) => function (_2078_i2) {
          return _2077_i0;
        })(_2031_i0)))), _dafny.Seq.Concat(_2066_v30, _2066_v30), _dafny.Seq.of(_2030_v3));
        _2066_v30 = (_2076_v38)[_module.__default.safeIndex((_2030_v3).multipliedBy(_2030_v3), new BigNumber((_2076_v38).length))];
      }
      let _2079_v39;
      _2079_v39 = false;
      let _2080_v41;
      _2080_v41 = _dafny.Seq.of(new BigNumber((function () {
        let _coll43 = new _dafny.Map();
        for (const _compr_43 of (_2028_v1).Elements) {
          let _2081_v40 = _compr_43;
          if (_dafny.Seq.contains(_2028_v1, _2081_v40)) {
            _coll43.push([_2081_v40,_2030_v3]);
          }
        }
        return _coll43;
      }()).length));
      let _2082_v42;
      _2082_v42 = _dafny.MultiSet.fromElements(_2028_v1, _dafny.Seq.UnicodeFromString("fogjs"), _2028_v1);
      _2079_v39 = ((new BigNumber((_2082_v42).cardinality())).plus(_2030_v3)).isLessThan(new BigNumber((_2080_v41).length));
      let _2083_i3;
      _2083_i3 = _dafny.ZERO;
      L10: {
        while (_2079_v39) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2083_i3)) {
              break L10;
            }
            _2083_i3 = (_2083_i3).plus(_dafny.ONE);
            let _2084_v43;
            _2084_v43 = _module.D2.create_DC8(_2079_v39, _2028_v1);
            _2030_v3 = new BigNumber(((_2084_v43).dtor_cf15).length);
            let _2085_v44;
            let _init67 = ((_2086_v3) => function (_2087_i4) {
              return _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2086_v3);
            })(_2030_v3);
            let _nw350 = Array((new BigNumber(15)).toNumber());
            for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw350.length); _i0_67++) {
              _nw350[_i0_67] = _init67(new BigNumber(_i0_67));
            }
            _2085_v44 = _nw350;
            let _2088_v45;
            _2088_v45 = _dafny.MultiSet.fromElements(_2030_v3);
            let _2089_v46;
            _2089_v46 = _dafny.Map.Empty.slice().updateUnsafe(_2030_v3,_2030_v3);
            let _2090_v47;
            _2090_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(((_2088_v45).contains(_2030_v3)) ? ((_2088_v45).get(_2030_v3)) : ((((_2089_v46).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2079_v39,_2030_v3)).length))) ? ((_2089_v46).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2079_v39,_2030_v3)).length))) : (new BigNumber((_2028_v1).length))))));
            let _index326 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2085_v44).length));
            (_2085_v44)[_index326] = _2090_v47;
            let _2091_v48;
            _2091_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2030_v3,_2079_v39);
            let _index327 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2085_v44).length));
            (_2085_v44)[_index327] = (((((_2091_v48).contains(new BigNumber(543))) ? ((_2091_v48).get(new BigNumber(543))) : (_2079_v39))) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2030_v3)) : (_dafny.Map.Empty.slice().updateUnsafe(!(_2079_v39),_2030_v3)));
            let _2092_v49;
            let _nw351 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
            _2092_v49 = _nw351;
            let _2093_v50;
            _2093_v50 = _dafny.Set.fromElements(_2092_v49);
            let _2094_v51;
            _2094_v51 = _module.D0.create_DC2(_2079_v39, new BigNumber(693), _2093_v50, _2030_v3);
            let _2095_v52;
            _2095_v52 = _dafny.MultiSet.fromElements(_2094_v51);
            let _2096_v53;
            _2096_v53 = _dafny.Set.fromElements(_2095_v52);
            _2096_v53 = _dafny.Set.fromElements((_2095_v52).update(_2094_v51, _module.__default.abs(new BigNumber(285))));
            let _2097_v54;
            _2097_v54 = _dafny.Set.fromElements(new BigNumber(-256));
            _2097_v54 = (function () {
              let _coll44 = new _dafny.Set();
              for (const _compr_44 of _dafny.IntegerRange(new BigNumber(467), new BigNumber(541))) {
                let _2098_v55 = _compr_44;
                if (((new BigNumber(467)).isLessThanOrEqualTo(_2098_v55)) && ((_2098_v55).isLessThan(new BigNumber(541)))) {
                  _coll44.add((_2098_v55).minus(_2030_v3));
                }
              }
              return _coll44;
            }()).Difference(_2097_v54);
          }
        }
      }
      let _2099_v56;
      let _init68 = ((_2100_v1) => function (_2101_i5) {
        return (_2101_i5).minus(new BigNumber((_2100_v1).length));
      })(_2028_v1);
      let _nw352 = Array((new BigNumber(22)).toNumber());
      for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw352.length); _i0_68++) {
        _nw352[_i0_68] = _init68(new BigNumber(_i0_68));
      }
      _2099_v56 = _nw352;
      let _2102_v57;
      _2102_v57 = _dafny.Seq.of((_this).f10, _2079_v39);
      let _index328 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
      (_2099_v56)[_index328] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2102_v57).length),_module.__default.fm0(globalState))).length);
      let _index329 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
      (_2099_v56)[_index329] = _2030_v3;
      let _2103_i6;
      _2103_i6 = _dafny.ZERO;
      L11: {
        while ((_this).f10) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2103_i6)) {
              break L11;
            }
            _2103_i6 = (_2103_i6).plus(_dafny.ONE);
            _2030_v3 = _2030_v3;
            _2079_v39 = _2079_v39;
            let _2104_v58;
            let _nw353 = Array((new BigNumber(18)).toNumber()).fill([]);
            _2104_v58 = _nw353;
            let _2105_v59;
            _2105_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2099_v56);
            let _index330 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_2104_v58).length));
            (_2104_v58)[_index330] = (((_2105_v59).contains((_this).f10)) ? ((_2105_v59).get((_this).f10)) : (_2099_v56));
            let _index331 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_2104_v58).length));
            (_2104_v58)[_index331] = _2099_v56;
            let _2106_v60;
            _2106_v60 = _dafny.MultiSet.fromElements((_this).f10, _2079_v39, _2079_v39, !(_2079_v39), (_this).f10);
            let _2107_v61;
            _2107_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.MultiSet.fromElements(_2079_v39));
            let _index332 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
            let _index333 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
            let _rhs333 = (_2099_v56)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length))];
            let _rhs334 = _module.__default.fm0(globalState);
            let _rhs335 = (((_2107_v61).contains(_2079_v39)) ? ((_2107_v61).get(_2079_v39)) : (_2106_v60));
            let _rhs336 = (new BigNumber((_dafny.Seq.Concat(_2102_v57, _2102_v57)).length)).multipliedBy(_2030_v3);
            let _lhs210 = _2099_v56;
            let _lhs211 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
            let _lhs212 = _2099_v56;
            let _lhs213 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length));
            _lhs210[_lhs211] = _rhs333;
            _2030_v3 = _rhs334;
            _2106_v60 = _rhs335;
            _lhs212[_lhs213] = _rhs336;
          }
        }
      }
      let _2108_v62;
      _2108_v62 = _dafny.Set.fromElements((_2099_v56)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length))], _2030_v3, (_2099_v56)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length))], _2030_v3);
      _2030_v3 = (new BigNumber(((_2108_v62).Difference(_2108_v62)).length)).plus(_module.__default.safeDivisionInt((_2099_v56)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_2099_v56).length))], new BigNumber(148)));
      return;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let _2109_v0;
      _2109_v0 = new BigNumber(984);
      let _hi10 = new BigNumber((function () {
        let _coll45 = new _dafny.Map();
        for (const _compr_45 of _dafny.IntegerRange(new BigNumber(219), new BigNumber(267))) {
          let _2110_v1 = _compr_45;
          if (((new BigNumber(219)).isLessThanOrEqualTo(_2110_v1)) && ((_2110_v1).isLessThan(new BigNumber(267)))) {
            _coll45.push([_module.__default.safeDivisionInt(_2110_v1, new BigNumber((_dafny.Seq.UnicodeFromString("kja")).length)),p1]);
          }
        }
        return _coll45;
      }()).length);
      for (let _2111_i0 = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), _2109_v0); _2111_i0.isLessThan(_hi10); _2111_i0 = _2111_i0.plus(_dafny.ONE)) {
        let _2112_v2;
        _2112_v2 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber(641), _2109_v0), _2109_v0);
        _2112_v2 = _2112_v2;
        let _2113_v3;
        _2113_v3 = true;
        _2113_v3 = p1;
        let _2114_v4;
        _2114_v4 = _dafny.Seq.of(p1);
        let _2115_v5;
        _2115_v5 = _dafny.Seq.UnicodeFromString("ocvmfl");
        let _2116_v6;
        _2116_v6 = _dafny.MultiSet.fromElements((_2114_v4)[_module.__default.safeIndex(_2111_i0, new BigNumber((_2114_v4).length))], (_2111_i0).isEqualTo(new BigNumber((_2115_v5).length)), (_this).f10, (_this).f10, p1);
        let _2117_v7;
        _2117_v7 = _module.D0.create_DC1(_2109_v0);
        let _2118_v8;
        _2118_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2117_v7,_2113_v3);
        let _2119_v9;
        _2119_v9 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2120_v10;
        let _nw354 = Array((new BigNumber(16)).toNumber());
        _nw354[(_dafny.ZERO).toNumber()] = _2113_v3;
        _nw354[(_dafny.ONE).toNumber()] = (_2111_i0).isLessThanOrEqualTo(new BigNumber(858));
        _nw354[(new BigNumber(2)).toNumber()] = p1;
        _nw354[(new BigNumber(3)).toNumber()] = p1;
        _nw354[(new BigNumber(4)).toNumber()] = (_2109_v0).isEqualTo(_2111_i0);
        _nw354[(new BigNumber(5)).toNumber()] = p1;
        _nw354[(new BigNumber(6)).toNumber()] = !((((_2118_v8).contains(_2117_v7)) ? ((_2118_v8).get(_2117_v7)) : (_2113_v3)));
        _nw354[(new BigNumber(7)).toNumber()] = !((_this).f10);
        _nw354[(new BigNumber(8)).toNumber()] = _2113_v3;
        _nw354[(new BigNumber(9)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), function (_2121_i1) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _module.__default.safeIndex(_2109_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), function (_2122_i1) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length)), _2119_v9), _2115_v5);
        _nw354[(new BigNumber(10)).toNumber()] = false;
        _nw354[(new BigNumber(11)).toNumber()] = (_this).f10;
        _nw354[(new BigNumber(12)).toNumber()] = !((_2113_v3) === (!(_2113_v3)));
        _nw354[(new BigNumber(13)).toNumber()] = (_2111_i0).isLessThanOrEqualTo(_2111_i0);
        _nw354[(new BigNumber(14)).toNumber()] = (_this).f10;
        _nw354[(new BigNumber(15)).toNumber()] = p1;
        _2120_v10 = _nw354;
        let _rhs337 = _2116_v6;
        let _rhs338 = _2120_v10;
        _2116_v6 = _rhs337;
        _2120_v10 = _rhs338;
        let _2123_v11;
        let _nw355 = Array((new BigNumber(26)).toNumber()).fill(_module.D2.Default());
        _2123_v11 = _nw355;
        let _2124_v12;
        let _nw356 = new _module.C4();
        _nw356.__ctor(_2123_v11, false);
        _2124_v12 = _nw356;
      }
      let _2125_v13;
      _2125_v13 = _dafny.Seq.of(!(p1), (_this).f10, p1);
      _2125_v13 = _2125_v13;
      let _2126_i2;
      _2126_i2 = _dafny.ZERO;
      L12: {
        while ((_this).f10) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2126_i2)) {
              break L12;
            }
            _2126_i2 = (_2126_i2).plus(_dafny.ONE);
            let _2127_v14;
            _2127_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_dafny.ZERO).minus(_2109_v0));
            let _2128_v15;
            _2128_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_2127_v14).contains(p1)) ? ((_2127_v14).get(p1)) : (new BigNumber((_2125_v13).length))),(_this).f10);
            if ((((_2128_v15).contains(_2109_v0)) ? ((_2128_v15).get(_2109_v0)) : (!(p1)))) {
              let _2129_v16;
              _2129_v16 = false;
              _2129_v16 = (_module.__default.safeModuloInt(_2109_v0, _2109_v0)).isLessThan(_2109_v0);
              let _2130_v17;
              let _init69 = function (_2131_i3) {
                return !(false);
              };
              let _nw357 = Array((new BigNumber(14)).toNumber());
              for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw357.length); _i0_69++) {
                _nw357[_i0_69] = _init69(new BigNumber(_i0_69));
              }
              _2130_v17 = _nw357;
              let _2132_v18;
              let _nw358 = Array((_dafny.ONE).toNumber());
              _nw358[(_dafny.ZERO).toNumber()] = _2130_v17;
              _2132_v18 = _nw358;
              let _index334 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_2132_v18).length));
              (_2132_v18)[_index334] = _2130_v17;
              let _index335 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_2132_v18).length));
              (_2132_v18)[_index335] = _2130_v17;
              let _2133_v19;
              _2133_v19 = _dafny.Set.fromElements((_this).f10, _2129_v16);
              let _2134_v20;
              _2134_v20 = _dafny.Seq.of(new BigNumber((_2133_v19).length));
              let _2135_v21;
              _2135_v21 = _module.D1.create_DC5(_2129_v16, _2134_v20, _2109_v0);
              let _2136_v22;
              _2136_v22 = _dafny.MultiSet.fromElements(_2129_v16, true, _2129_v16, false, false);
              let _2137_v23;
              let _nw359 = Array((new BigNumber(26)).toNumber());
              _nw359[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_2109_v0, _2109_v0);
              _nw359[(_dafny.ONE).toNumber()] = (_2135_v21).dtor_cf10;
              _nw359[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-441)), ((_2138_v0) => function (_2139_i4) {
                return _2138_v0;
              })(_2109_v0));
              _nw359[(new BigNumber(3)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(4)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(5)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(6)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_2109_v0);
              _nw359[(new BigNumber(8)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_2134_v20, _module.__default.safeIndex(_2109_v0, new BigNumber((_2134_v20).length)), _2109_v0);
              _nw359[(new BigNumber(10)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_2134_v20, _module.__default.safeIndex(new BigNumber((_2125_v13).length), new BigNumber((_2134_v20).length)), _2109_v0);
              _nw359[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_2109_v0);
              _nw359[(new BigNumber(13)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(803)), ((_2140_v20) => function (_2141_i5) {
                return new BigNumber((_2140_v20).length);
              })(_2134_v20));
              _nw359[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_2134_v20, _module.__default.safeIndex(_2109_v0, new BigNumber((_2134_v20).length)), _2109_v0);
              _nw359[(new BigNumber(16)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_2109_v0);
              _nw359[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(856)), ((_2142_v0) => function (_2143_i6) {
                return _2142_v0;
              })(_2109_v0));
              _nw359[(new BigNumber(19)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(_2109_v0), new BigNumber((_2136_v22).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(969)), ((_2144_v0) => function (_2145_i7) {
                return _2144_v0;
              })(_2109_v0))).length));
              _nw359[(new BigNumber(21)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(22)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(23)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(24)).toNumber()] = _2134_v20;
              _nw359[(new BigNumber(25)).toNumber()] = _2134_v20;
              _2137_v23 = _nw359;
              let _2146_v24;
              _2146_v24 = _dafny.Map.Empty.slice().updateUnsafe((_2109_v0).minus(_2109_v0),_2137_v23);
              let _2147_v25;
              _2147_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2109_v0,_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber(149)));
              _2146_v24 = (_2146_v24).update((new BigNumber((((((_2147_v25).contains(_2109_v0)) ? ((_2147_v25).get(_2109_v0)) : (_2127_v14))).update((_this).f10, _2109_v0)).length)).minus(_2109_v0), _2137_v23);
              let _2148_v26;
              _2148_v26 = _dafny.MultiSet.fromElements(new BigNumber(847));
              let _2149_v27;
              let _nw360 = new _module.C1();
              _nw360.__ctor((_2148_v26).IsDisjointFrom(_2148_v26));
              _2149_v27 = _nw360;
              let _2150_v28;
              let _nw361 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
              _2150_v28 = _nw361;
              let _2151_v29;
              _2151_v29 = _dafny.Set.fromElements(_2150_v28);
              _2129_v16 = !(!(_2129_v16)) || ((_2151_v29).IsProperSubsetOf(_2151_v29));
            } else {
              _2109_v0 = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
              let _2152_v30;
              _2152_v30 = _module.D0.create_DC0(_2109_v0);
              let _2153_v31;
              let _nw362 = Array((new BigNumber(5)).toNumber());
              _nw362[(_dafny.ZERO).toNumber()] = _module.__default.fm1(_dafny.Seq.of(_2152_v30, _2152_v30), _2109_v0, _2109_v0, globalState);
              _nw362[(_dafny.ONE).toNumber()] = p1;
              _nw362[(new BigNumber(2)).toNumber()] = (_this).f10;
              _nw362[(new BigNumber(3)).toNumber()] = true;
              _nw362[(new BigNumber(4)).toNumber()] = p1;
              _2153_v31 = _nw362;
              let _2154_v32;
              _2154_v32 = _module.D3.create_DC10(_2153_v31);
              let _2155_v33;
              _2155_v33 = _dafny.Seq.UnicodeFromString("hv");
              let _2156_v34;
              _2156_v34 = new _dafny.CodePoint('u'.codePointAt(0));
              let _rhs339 = _module.__default.fm0(globalState);
              let _rhs340 = _2154_v32;
              let _rhs341 = (((_this).f10) ? (_dafny.Seq.UnicodeFromString("jph")) : (_dafny.Seq.UnicodeFromString("di")));
              let _rhs342 = _2156_v34;
              _2109_v0 = _rhs339;
              _2154_v32 = _rhs340;
              _2155_v33 = _rhs341;
              _2156_v34 = _rhs342;
              let _2157_v35;
              _2157_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2152_v30,new BigNumber(650));
              _2155_v33 = _dafny.Seq.Concat(_2155_v33, (_module.D2.create_DC9(_2109_v0, new BigNumber((_2127_v14).length), _2155_v33, (_this).f10, _2157_v35)).dtor_cf18);
              let _2158_v36;
              _2158_v36 = false;
              _2158_v36 = _2158_v36;
              let _2159_v37;
              let _nw363 = Array((new BigNumber(15)).toNumber()).fill([]);
              _2159_v37 = _nw363;
              let _2160_v38;
              _2160_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2109_v0),_2153_v31);
              let _2161_v39;
              let _nw364 = Array((new BigNumber(8)).toNumber());
              _nw364[(_dafny.ZERO).toNumber()] = _2153_v31;
              _nw364[(_dafny.ONE).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(2)).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(3)).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(4)).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(5)).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(6)).toNumber()] = _2153_v31;
              _nw364[(new BigNumber(7)).toNumber()] = (((_2160_v38).contains(new BigNumber(908))) ? ((_2160_v38).get(new BigNumber(908))) : (_2153_v31));
              _2161_v39 = _nw364;
              let _2162_v40;
              _2162_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2109_v0,_2161_v39);
              let _index336 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2159_v37).length));
              (_2159_v37)[_index336] = (((_2162_v40).contains(_2109_v0)) ? ((_2162_v40).get(_2109_v0)) : (_2161_v39));
              let _index337 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2159_v37).length));
              (_2159_v37)[_index337] = _2161_v39;
            }
            _2109_v0 = (_2109_v0).multipliedBy(_2109_v0);
            let _2163_v41;
            let _init70 = function (_2164_i8) {
              return ((_this).f10) || (true);
            };
            let _nw365 = Array((new BigNumber(3)).toNumber());
            for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw365.length); _i0_70++) {
              _nw365[_i0_70] = _init70(new BigNumber(_i0_70));
            }
            _2163_v41 = _nw365;
            let _index338 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length));
            (_2163_v41)[_index338] = !_dafny.Seq.contains(_2125_v13, (_this).f10);
            let _index339 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length));
            (_2163_v41)[_index339] = !(p1);
            let _2165_v42;
            _2165_v42 = _module.D6.create_DC16(false);
            let _2166_v43;
            _2166_v43 = _dafny.MultiSet.fromElements((_2165_v42).dtor_cf30, (_2125_v13)[_module.__default.safeIndex(_2109_v0, new BigNumber((_2125_v13).length))]);
            let _2167_v44;
            _2167_v44 = _module.D3.create_DC11(new BigNumber((_2166_v43).cardinality()), _2109_v0, _2109_v0, (_this).f10);
            if ((_2167_v44).dtor_cf25) {
              let _2168_v46;
              let _init71 = ((_2169_v15) => function (_2170_i9) {
                return _module.D2.create_DC7(function () {
  let _coll46 = new _dafny.Set();
  for (const _compr_46 of (_dafny.MultiSet.fromElements(new BigNumber((_2169_v15).length))).Elements) {
    let _2171_v45 = _compr_46;
    if ((_dafny.MultiSet.fromElements(new BigNumber((_2169_v15).length))).contains(_2171_v45)) {
      _coll46.add(_module.__default.safeModuloInt(_2171_v45, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(322))).cardinality())));
    }
  }
  return _coll46;
}());
              })(_2128_v15);
              let _nw366 = Array((new BigNumber(29)).toNumber());
              for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw366.length); _i0_71++) {
                _nw366[_i0_71] = _init71(new BigNumber(_i0_71));
              }
              _2168_v46 = _nw366;
              let _2172_v47;
              let _nw367 = new _module.C4();
              _nw367.__ctor((((_2163_v41)[_module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length))]) ? (_2168_v46) : (_2168_v46)), (_this).f10);
              _2172_v47 = _nw367;
              _2109_v0 = (_2109_v0).plus(_2109_v0);
              let _2173_v48;
              _2173_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2109_v0,_dafny.Seq.UnicodeFromString("krvabqmix"));
              let _2174_v49;
              _2174_v49 = _dafny.Seq.UnicodeFromString("rr");
              _2173_v48 = (_2173_v48).update(_2109_v0, _2174_v49);
              _2109_v0 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("djnwyp")).length), _2109_v0)).minus(_2109_v0);
              _2109_v0 = _2109_v0;
            } else {
              let _index340 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length));
              (_2163_v41)[_index340] = p1;
              _2109_v0 = (_2109_v0).multipliedBy(_2109_v0);
              let _2175_v50;
              _2175_v50 = _module.D0.create_DC0(_2109_v0);
              let _2176_v51;
              _2176_v51 = _dafny.Seq.of(_2175_v50);
              let _index341 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length));
              (_2163_v41)[_index341] = _module.__default.fm1(_2176_v51, _2109_v0, (_2109_v0).multipliedBy(_2109_v0), globalState);
              let _index342 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_2163_v41).length));
              (_2163_v41)[_index342] = p1;
              _2109_v0 = _2109_v0;
            }
          }
        }
      }
      let _2177_v52;
      _2177_v52 = _dafny.Set.fromElements(p1, p1);
      let _2178_v53;
      let _nw368 = new _module.C3();
      _nw368.__ctor(new BigNumber((_2177_v52).length), p1);
      _2178_v53 = _nw368;
      let _2179_v54;
      _2179_v54 = _dafny.MultiSet.fromElements(_2178_v53);
      let _2180_v55;
      _2180_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(659),_2179_v54);
      let _hi11 = (_2178_v53).f14;
      for (let _2181_i10 = ((p1) ? (new BigNumber((_2180_v55).length)) : (new BigNumber((_2125_v13).length))); _2181_i10.isLessThan(_hi11); _2181_i10 = _2181_i10.plus(_dafny.ONE)) {
        let _2182_v56;
        _2182_v56 = _dafny.MultiSet.fromElements(_2181_i10);
        let _2183_v57;
        let _nw369 = new _module.C1();
        _nw369.__ctor(p1);
        _2183_v57 = _nw369;
        _2109_v0 = _module.__default.safeModuloInt((_2178_v53).f14, _module.__default.safeModuloInt(_2181_i10, (((_2182_v56).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2183_v57,_dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), function (_2185_i11) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }))).length))) ? ((_2182_v56).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2183_v57,_dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), function (_2184_i11) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }))).length))) : (_2109_v0))));
        let _rhs343 = !(_2183_v57.f7);
        let _rhs344 = p1;
        let _lhs214 = _2183_v57;
        let _lhs215 = _2183_v57;
        _lhs214.f7 = _rhs343;
        _lhs215.f7 = _rhs344;
        let _2186_v58;
        _2186_v58 = _dafny.Seq.UnicodeFromString("gmfjacoit");
        let _2187_v59;
        let _nw370 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _2187_v59 = _nw370;
        let _2188_v60;
        _2188_v60 = _module.D14.create_DC39(_2186_v58, _2187_v59);
        let _2189_v61;
        let _nw371 = new _module.C0();
        _nw371.__ctor(_dafny.Seq.IsProperPrefixOf((_2188_v60).dtor_cf63, _dafny.Seq.UnicodeFromString("egn")), _2183_v57.f7);
        _2189_v61 = _nw371;
        _2186_v58 = _2186_v58;
      }
      let _2190_v62;
      _2190_v62 = true;
      let _2191_v63;
      _2191_v63 = _dafny.Seq.of(_2109_v0, (_2178_v53).f14);
      let _2192_v64;
      _2192_v64 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_2191_v63, _2191_v63),false);
      let _2193_v65;
      _2193_v65 = _dafny.MultiSet.fromElements(p1);
      let _rhs345 = !(!((((_2192_v64).contains(p1)) ? ((_2192_v64).get(p1)) : ((_2193_v65).IsSubsetOf(_2193_v65)))));
      let _rhs346 = _2125_v13;
      _2190_v62 = _rhs345;
      _2125_v13 = _rhs346;
      let _2194_v66;
      _2194_v66 = _module.D0.create_DC0(new BigNumber((_2125_v13).length));
      let _2195_v67;
      _2195_v67 = _dafny.MultiSet.fromElements(_2109_v0, _2109_v0);
      let _2196_v68;
      _2196_v68 = _dafny.Seq.UnicodeFromString("bxc");
      let _2197_v69;
      _2197_v69 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2198_v70;
      _2198_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2197_v69,new BigNumber((_dafny.Seq.UnicodeFromString("ab")).length));
      let _2199_v71;
      _2199_v71 = _dafny.Map.Empty.slice().updateUnsafe(true,_2198_v70);
      let _2200_v73;
      _2200_v73 = _dafny.MultiSet.fromElements(_2197_v69, _2197_v69);
      let _2201_i12;
      _2201_i12 = _dafny.ZERO;
      L13: {
        while (_module.__default.fm1(_dafny.Seq.update(_dafny.Seq.of(_2194_v66, _2194_v66), _module.__default.safeIndex(new BigNumber((_2195_v67).cardinality()), new BigNumber((_dafny.Seq.of(_2194_v66, _2194_v66)).length)), _module.D0.create_DC0(new BigNumber((_2196_v68).length))), (_dafny.ZERO).minus(_2109_v0), new BigNumber(((((_2199_v71).contains(true)) ? ((_2199_v71).get(true)) : (function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of (_2200_v73).Elements) {
            let _2243_v72 = _compr_47;
            if ((_2200_v73).contains(_2243_v72)) {
              _coll47.push([_2243_v72,(_dafny.ZERO).minus(new BigNumber(-746))]);
            }
          }
          return _coll47;
        }()))).length), globalState)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2201_i12)) {
              break L13;
            }
            _2201_i12 = (_2201_i12).plus(_dafny.ONE);
            _2125_v13 = _dafny.Seq.Concat(_2125_v13, _2125_v13);
            let _2202_v74;
            _2202_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2196_v68,false);
            let _2203_v75;
            _2203_v75 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? ((_2178_v53).f14) : (new BigNumber((_2202_v74).length))),((p1) ? (p1) : ((_this).f10)));
            _2203_v75 = (_2203_v75).update((_2178_v53).f14, !(((_this).f10) === ((_this).f10)));
            _2109_v0 = _2109_v0;
            let _2204_v76;
            let _init72 = ((_2205_v62) => function (_2206_i13) {
              return _2205_v62;
            })(_2190_v62);
            let _nw372 = Array((new BigNumber(27)).toNumber());
            for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw372.length); _i0_72++) {
              _nw372[_i0_72] = _init72(new BigNumber(_i0_72));
            }
            _2204_v76 = _nw372;
            let _2207_v77;
            _2207_v77 = _dafny.Seq.of(_2204_v76);
            let _2208_v78;
            _2208_v78 = _module.D6.create_DC15(_module.__default.fm26(_2190_v62, new BigNumber((_2207_v77).length), (_dafny.ZERO).minus((_2178_v53).f14), globalState));
            let _2209_v79;
            _2209_v79 = _module.D6.create_DC18(_2208_v78);
            let _source16 = _module.D14.create_DC40(_2209_v79, !(_2190_v62), (_this).f10);
            if (_source16.is_DC39) {
              let _2210___mcc_h0 = (_source16).cf63;
              let _2211___mcc_h1 = (_source16).cf64;
              let _2212_cf64 = _2211___mcc_h1;
              let _2213_cf63 = _2210___mcc_h0;
              let _2214_v80;
              _2214_v80 = _dafny.Map.Empty.slice().updateUnsafe((_2178_v53).f14,_2192_v64);
              let _2215_v81;
              let _nw373 = Array((new BigNumber(15)).toNumber());
              _nw373[(_dafny.ZERO).toNumber()] = _2192_v64;
              _nw373[(_dafny.ONE).toNumber()] = _module.__default.fm34(_2109_v0, p1, _2190_v62, (_2178_v53).f14, globalState);
              _nw373[(new BigNumber(2)).toNumber()] = _2192_v64;
              _nw373[(new BigNumber(3)).toNumber()] = _2192_v64;
              _nw373[(new BigNumber(4)).toNumber()] = _module.__default.fm34(new BigNumber(309), (_this).f10, (_this).f10, new BigNumber((_2193_v65).cardinality()), globalState);
              _nw373[(new BigNumber(5)).toNumber()] = _module.__default.fm34(new BigNumber(376), (_this).f10, (_this).f10, _2109_v0, globalState);
              _nw373[(new BigNumber(6)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(p1,_2190_v62)).update(p1, (_this).f10)).update((_2178_v53).fm8(_2109_v0, globalState), p1);
              _nw373[(new BigNumber(7)).toNumber()] = _2192_v64;
              _nw373[(new BigNumber(8)).toNumber()] = (_2192_v64).Merge((((_2214_v80).contains(new BigNumber((_2196_v68).length))) ? ((_2214_v80).get(new BigNumber((_2196_v68).length))) : (_2192_v64)));
              _nw373[(new BigNumber(9)).toNumber()] = (_2192_v64).update(p1, _2190_v62);
              _nw373[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2190_v62,true);
              _nw373[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2190_v62,_2190_v62);
              _nw373[(new BigNumber(12)).toNumber()] = _2192_v64;
              _nw373[(new BigNumber(13)).toNumber()] = _2192_v64;
              _nw373[(new BigNumber(14)).toNumber()] = _2192_v64;
              _2215_v81 = _nw373;
              let _index343 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2215_v81).length));
              (_2215_v81)[_index343] = _2192_v64;
              let _index344 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2215_v81).length));
              (_2215_v81)[_index344] = _2192_v64;
              _2197_v69 = _2197_v69;
              _2190_v62 = false;
              _2192_v64 = (_2192_v64).Merge(_2192_v64);
            } else if (_source16.is_DC40) {
              let _2216___mcc_h2 = (_source16).cf65;
              let _2217___mcc_h3 = (_source16).cf66;
              let _2218___mcc_h4 = (_source16).cf67;
              let _2219_cf67 = _2218___mcc_h4;
              let _2220_cf66 = _2217___mcc_h3;
              let _2221_cf65 = _2216___mcc_h2;
              let _2222_v82;
              _2222_v82 = _dafny.MultiSet.fromElements(_2196_v68, _2196_v68, _dafny.Seq.Concat(_2196_v68, _2196_v68));
              _2109_v0 = (((_2222_v82).contains(_2196_v68)) ? ((_2222_v82).get(_2196_v68)) : ((_2178_v53).f14));
              _2196_v68 = _2196_v68;
              _2109_v0 = _2109_v0;
              let _2223_v83;
              _2223_v83 = _dafny.Map.Empty.slice().updateUnsafe(false,_2125_v13);
              _2193_v65 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2125_v13, (((_2223_v83).contains(_2219_cf67)) ? ((_2223_v83).get(_2219_cf67)) : (_2125_v13))))).Difference(_2193_v65);
            } else if (_source16.is_DC38) {
              let _2224___mcc_h5 = (_source16).cf62;
              let _2225_cf62 = _2224___mcc_h5;
              let _2226_v84;
              let _nw374 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
              _2226_v84 = _nw374;
              let _2227_v85;
              let _nw375 = Array((new BigNumber(16)).toNumber()).fill(_module.D6.Default());
              _2227_v85 = _nw375;
              let _2228_v86;
              _2228_v86 = _dafny.Seq.of(_2227_v85, _2227_v85);
              let _2229_v87;
              _2229_v87 = _dafny.Seq.of(_2227_v85, _2227_v85, _2227_v85, (_2228_v86)[_module.__default.safeIndex((_2178_v53).f14, new BigNumber((_2228_v86).length))], _2227_v85);
              let _rhs347 = _2109_v0;
              let _rhs348 = _2226_v84;
              let _rhs349 = _dafny.Seq.Concat(_2229_v87, _2229_v87);
              let _rhs350 = _module.__default.fm0(globalState);
              _2109_v0 = _rhs347;
              _2226_v84 = _rhs348;
              _2228_v86 = _rhs349;
              _2109_v0 = _rhs350;
              _2109_v0 = (((false) ? (new BigNumber(-261)) : (_2109_v0))).minus(new BigNumber(725));
              let _2230_v88;
              let _init73 = ((_2231_v53) => function (_2232_i14) {
                return (_2232_i14).plus((_2231_v53).f14);
              })(_2178_v53);
              let _nw376 = Array((new BigNumber(3)).toNumber());
              for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw376.length); _i0_73++) {
                _nw376[_i0_73] = _init73(new BigNumber(_i0_73));
              }
              _2230_v88 = _nw376;
              let _index345 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2230_v88).length));
              (_2230_v88)[_index345] = (_2178_v53).f14;
              let _index346 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2230_v88).length));
              (_2230_v88)[_index346] = (_2178_v53).f14;
              let _2233_v89;
              _2233_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_2109_v0, (_2230_v88)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2230_v88).length))]),_module.__default.fm0(globalState));
              _2233_v89 = (_2233_v89).update(_module.__default.fm0(globalState), _2109_v0);
            } else {
              let _2234___mcc_h6 = (_source16).cf68;
              let _2235_cf68 = _2234___mcc_h6;
              let _index347 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2204_v76).length));
              (_2204_v76)[_index347] = (_this).f10;
              let _2236_v90;
              _2236_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2125_v13);
              let _index348 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2204_v76).length));
              (_2204_v76)[_index348] = ((_module.__default.fm49(p1, _2236_v90, false, globalState)).update(p1, _module.__default.abs((((_2195_v67).contains((_2178_v53).f14)) ? ((_2195_v67).get((_2178_v53).f14)) : ((_dafny.ZERO).minus((_2178_v53).f14)))))).equals(_2193_v65);
              let _index349 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2204_v76).length));
              (_2204_v76)[_index349] = _2190_v62;
              _2190_v62 = (_2204_v76)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2204_v76).length))];
              let _2237_v91;
              let _nw377 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _2237_v91 = _nw377;
              let _2238_v92;
              _2238_v92 = _module.D16.create_DC46((_2178_v53).f14, _2237_v91, _2197_v69);
              let _2239_v93;
              _2239_v93 = _module.D9.create_DC23(_2237_v91);
              let _2240_v94;
              _2240_v94 = _dafny.Seq.of(_2237_v91, _2237_v91, _2237_v91, _2237_v91, _2237_v91);
              let _2241_v95;
              _2241_v95 = _module.D14.create_DC39(_module.__default.fm20(_2190_v62, _2190_v62, globalState), _2237_v91);
              let _2242_v96;
              let _nw378 = Array((new BigNumber(28)).toNumber());
              _nw378[(_dafny.ZERO).toNumber()] = _2237_v91;
              _nw378[(_dafny.ONE).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(2)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(3)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(4)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(5)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(6)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(7)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(8)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(9)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(10)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(11)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(12)).toNumber()] = (_2238_v92).dtor_cf74;
              _nw378[(new BigNumber(13)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(14)).toNumber()] = (_2239_v93).dtor_cf44;
              _nw378[(new BigNumber(15)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(16)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(17)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(18)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(19)).toNumber()] = (_2240_v94)[_module.__default.safeIndex((_2178_v53).f14, new BigNumber((_2240_v94).length))];
              _nw378[(new BigNumber(20)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(21)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(22)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(23)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(24)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(25)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(26)).toNumber()] = _2237_v91;
              _nw378[(new BigNumber(27)).toNumber()] = (_2241_v95).dtor_cf64;
              _2242_v96 = _nw378;
              _2242_v96 = _2242_v96;
            }
          }
        }
      }
      let _pat_let_tv48 = _2109_v0;
      r0 = function (_pat_let65_0) {
        return function (_2244_dt__update__tmp_h0) {
          return function (_pat_let66_0) {
            return function (_2245_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_2245_dt__update_hcf0_h0);
            }(_pat_let66_0);
          }(_pat_let_tv48);
        }(_pat_let65_0);
      }(_2194_v66);
      return r0;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f7 = false;
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f9, f7) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f7 = f7;
      return;
    }
    fm8(p0, globalState) {
      let _this = this;
      return _this.f7;
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber(101), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(390)), function (_2246_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-704)), function (_2247_i1) {
          return new BigNumber((_dafny.MultiSet.fromElements((_this).f9, (_this).f9)).cardinality());
        });
      })).length))).multipliedBy(new BigNumber(((function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of _dafny.IntegerRange(new BigNumber(888), new BigNumber(769))) {
          let _2248_v0 = _compr_48;
          if (((new BigNumber(888)).isLessThanOrEqualTo(_2248_v0)) && ((_2248_v0).isLessThan(new BigNumber(769)))) {
            _coll48.push([(_2248_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("oxffjj")).length)),(_this).f9]);
          }
        }
        return _coll48;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-992),new BigNumber(215))).length))).length)))).dtor_cf1,_this.f7))).length));
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(254);
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _2249_v0;
      let _init74 = function (_2250_i0) {
        return _this.f7;
      };
      let _nw379 = Array((new BigNumber(27)).toNumber());
      for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw379.length); _i0_74++) {
        _nw379[_i0_74] = _init74(new BigNumber(_i0_74));
      }
      _2249_v0 = _nw379;
      let _index350 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
      (_2249_v0)[_index350] = (_this).fm8(p0, globalState);
      let _index351 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
      (_2249_v0)[_index351] = true;
      let _2251_i1;
      _2251_i1 = _dafny.ZERO;
      L14: {
        while (!_dafny.areEqual(p1, _dafny.Seq.Concat(p1, p1))) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2251_i1)) {
              break L14;
            }
            _2251_i1 = (_2251_i1).plus(_dafny.ONE);
            r1 = (_dafny.ZERO).minus(new BigNumber(-752));
            let _2252_v1;
            _2252_v1 = _dafny.Seq.of(_dafny.Seq.of(p0, p0));
            let _2253_v2;
            _2253_v2 = _dafny.Seq.of((_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))]);
            let _2254_v3;
            _2254_v3 = _dafny.MultiSet.fromElements((_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))], (_this).f9);
            let _2255_v4;
            _2255_v4 = _dafny.Seq.of(p0);
            let _2256_v5;
            _2256_v5 = _dafny.Seq.of(new BigNumber((_2255_v4).length));
            let _2257_v6;
            _2257_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f9);
            let _2258_v7;
            let _nw380 = Array((new BigNumber(18)).toNumber());
            _nw380[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2252_v1, _dafny.Seq.update(_2252_v1, _module.__default.safeIndex(p0, new BigNumber((_2252_v1).length)), _dafny.Seq.of(new BigNumber((_2253_v2).length), p0)))).length);
            _nw380[(_dafny.ONE).toNumber()] = new BigNumber((_2254_v3).cardinality());
            _nw380[(new BigNumber(2)).toNumber()] = (_2256_v5)[_module.__default.safeIndex(p0, new BigNumber((_2256_v5).length))];
            _nw380[(new BigNumber(3)).toNumber()] = p0;
            _nw380[(new BigNumber(4)).toNumber()] = p0;
            _nw380[(new BigNumber(5)).toNumber()] = (_this).fm12(_this.f7, new BigNumber((_2257_v6).length), p0, (_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))], globalState);
            _nw380[(new BigNumber(6)).toNumber()] = p0;
            _nw380[(new BigNumber(7)).toNumber()] = (new BigNumber(764)).plus(new BigNumber(-606));
            _nw380[(new BigNumber(8)).toNumber()] = _module.__default.fm0(globalState);
            _nw380[(new BigNumber(9)).toNumber()] = p0;
            _nw380[(new BigNumber(10)).toNumber()] = (_this).fm9((_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))], p0, new _dafny.CodePoint('g'.codePointAt(0)), globalState);
            _nw380[(new BigNumber(11)).toNumber()] = (p0).plus(new BigNumber((p1).length));
            _nw380[(new BigNumber(12)).toNumber()] = new BigNumber((p1).length);
            _nw380[(new BigNumber(13)).toNumber()] = p0;
            _nw380[(new BigNumber(14)).toNumber()] = p0;
            _nw380[(new BigNumber(15)).toNumber()] = new BigNumber(611);
            _nw380[(new BigNumber(16)).toNumber()] = new BigNumber((_2257_v6).length);
            _nw380[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2255_v4).length), p0);
            _2258_v7 = _nw380;
            let _index352 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_2258_v7).length));
            (_2258_v7)[_index352] = new BigNumber(560);
            let _index353 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_2258_v7).length));
            (_2258_v7)[_index353] = p0;
            let _2259_v8;
            _2259_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
            let _2260_v9;
            _2260_v9 = new _dafny.CodePoint('w'.codePointAt(0));
            let _index354 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_2258_v7).length));
            (_2258_v7)[_index354] = new BigNumber((_dafny.Seq.Concat((((_2259_v8).contains(new BigNumber((_dafny.Seq.UnicodeFromString("ydewlhgde")).length))) ? ((_2259_v8).get(new BigNumber((_dafny.Seq.UnicodeFromString("ydewlhgde")).length))) : (p1)), _dafny.Seq.update(_dafny.Seq.update(p1, _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((p1).length)), _2260_v9), _module.__default.safeIndex((_2258_v7)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_2258_v7).length))], new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((p1).length)), _2260_v9)).length)), _2260_v9))).length);
            (_this).f7 = !(_this.f7);
          }
        }
      }
      let _index355 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
      (_2249_v0)[_index355] = true;
      let _2261_v10;
      _2261_v10 = new _dafny.CodePoint('m'.codePointAt(0));
      let _2262_v11;
      _2262_v11 = _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), _2261_v10, _2261_v10);
      _2262_v11 = _dafny.Seq.Concat(_dafny.Seq.of(_2261_v10), _dafny.Seq.of(_module.__default.fm13((_this).f9, p0, (_this).f9, globalState)));
      let _2263_v12;
      _2263_v12 = _dafny.MultiSet.fromElements(_this.f7, true);
      let _index356 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
      (_2249_v0)[_index356] = (((((_this).f9) ? (_this.f7) : ((_this).f9))) ? ((_2263_v12).IsSubsetOf(_2263_v12)) : ((_this).f9));
      let _2264_v13;
      _2264_v13 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((p1).length), p0, p0)).length), p0);
      let _2265_i2;
      _2265_i2 = _dafny.ZERO;
      L15: {
        while ((new BigNumber((_2264_v13).length)).isEqualTo(p0)) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2265_i2)) {
              break L15;
            }
            _2265_i2 = (_2265_i2).plus(_dafny.ONE);
            let _2266_v14;
            let _nw381 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2266_v14 = _nw381;
            let _index357 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2266_v14).length));
            (_2266_v14)[_index357] = _dafny.Seq.UnicodeFromString("wlg");
            let _2267_v15;
            _2267_v15 = _dafny.Seq.of(_2262_v11, _dafny.Seq.Concat(p1, _2262_v11));
            let _2268_v16;
            _2268_v16 = _dafny.Seq.of(p0);
            let _index358 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2266_v14).length));
            (_2266_v14)[_index358] = (_2267_v15)[_module.__default.safeIndex(((_this.f7) ? (new BigNumber((_2268_v16).length)) : (p0)), new BigNumber((_2267_v15).length))];
            _2261_v10 = new _dafny.CodePoint('m'.codePointAt(0));
            let _2269_v17;
            let _init75 = ((_2270_p0) => function (_2271_i3) {
              return (_2271_i3).multipliedBy(new BigNumber((_dafny.Seq.of(_2270_p0)).length));
            })(p0);
            let _nw382 = Array((new BigNumber(4)).toNumber());
            for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw382.length); _i0_75++) {
              _nw382[_i0_75] = _init75(new BigNumber(_i0_75));
            }
            _2269_v17 = _nw382;
            let _2272_v18;
            _2272_v18 = _dafny.Set.fromElements(_2269_v17);
            let _2273_v19;
            _2273_v19 = _dafny.Seq.of(_2272_v18, _2272_v18);
            let _2274_v21;
            _2274_v21 = _module.D0.create_DC2(_this.f7, (_dafny.ZERO).minus(p0), (_2273_v19)[_module.__default.safeIndex(new BigNumber((function () {
  let _coll49 = new _dafny.Set();
  for (const _compr_49 of _dafny.IntegerRange(new BigNumber(287), new BigNumber(279))) {
    let _2275_v20 = _compr_49;
    if (((new BigNumber(287)).isLessThanOrEqualTo(_2275_v20)) && ((_2275_v20).isLessThan(new BigNumber(279)))) {
      _coll49.add((_2275_v20).plus(p0));
    }
  }
  return _coll49;
}()).length), new BigNumber((_2273_v19).length))], (_this).fm9(_this.f7, p0, _2261_v10, globalState));
            let _index359 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
            (_2249_v0)[_index359] = (_2274_v21).dtor_cf2;
            let _index360 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length));
            (_2249_v0)[_index360] = !((_this).f9);
          }
        }
      }
      let _2276_v22;
      let _nw383 = Array((new BigNumber(19)).toNumber()).fill([]);
      _2276_v22 = _nw383;
      r0 = _2276_v22;
      let _2277_v23;
      let _init76 = ((_2278_p0) => function (_2279_i4) {
        return (_2279_i4).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber((_dafny.Seq.of(_2278_p0)).length))).length));
      })(p0);
      let _nw384 = Array((new BigNumber(6)).toNumber());
      for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw384.length); _i0_76++) {
        _nw384[_i0_76] = _init76(new BigNumber(_i0_76));
      }
      _2277_v23 = _nw384;
      let _2280_v24;
      _2280_v24 = _dafny.Seq.of(_2249_v0);
      let _2281_v25;
      _2281_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2277_v23,new BigNumber((_2280_v24).length));
      r1 = (((_this).fm12(_this.f7, p0, p0, (_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))], globalState)).plus(p0)).plus((new BigNumber((_2262_v11).length)).multipliedBy(new BigNumber((_2281_v25).length)));
      r2 = (_2249_v0)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2249_v0).length))];
      r3 = (new BigNumber((p1).length)).multipliedBy(p0);
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      (_this).f7 = p2;
      let _2282_v0;
      _2282_v0 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ajaosb"));
      let _2283_v1;
      _2283_v1 = _dafny.Seq.UnicodeFromString("onrgrefdt");
      let _hi12 = (p1).plus(p1);
      for (let _2284_i0 = new BigNumber((_dafny.Seq.Concat((_2282_v0)[_module.__default.safeIndex(p1, new BigNumber((_2282_v0).length))], _2283_v1)).length); _2284_i0.isLessThan(_hi12); _2284_i0 = _2284_i0.plus(_dafny.ONE)) {
        let _2285_v2;
        let _init77 = function (_2286_i1) {
          return (_this).f9;
        };
        let _nw385 = Array((new BigNumber(25)).toNumber());
        for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw385.length); _i0_77++) {
          _nw385[_i0_77] = _init77(new BigNumber(_i0_77));
        }
        _2285_v2 = _nw385;
        let _index361 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2285_v2).length));
        (_2285_v2)[_index361] = _this.f7;
        let _2287_v3;
        _2287_v3 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index362 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2285_v2).length));
        let _rhs351 = p0;
        let _rhs352 = _2287_v3;
        let _lhs216 = _2285_v2;
        let _lhs217 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_2285_v2).length));
        _lhs216[_lhs217] = _rhs351;
        _2287_v3 = _rhs352;
        let _2288_v4;
        _2288_v4 = _dafny.Map.Empty.slice().updateUnsafe((_2285_v2)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_2285_v2).length))],(_2285_v2)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_2285_v2).length))]);
        let _2289_v5;
        let _nw386 = new _module.C0();
        _nw386.__ctor(p0, (((_this).f9) ? ((((_2288_v4).contains(p0)) ? ((_2288_v4).get(p0)) : ((_this).f9))) : (p2)));
        _2289_v5 = _nw386;
        let _2290_v6;
        _2290_v6 = new BigNumber(88);
        _2290_v6 = _2284_i0;
        _2290_v6 = ((_2290_v6).multipliedBy(_2290_v6)).minus(_module.__default.safeModuloInt(_2284_i0, _2290_v6));
      }
      let _2291_v7;
      let _nw387 = Array((new BigNumber(2)).toNumber()).fill(_module.D14.Default());
      _2291_v7 = _nw387;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2291_v7).length))) {
        let _2292_i2 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2292_i2)) && ((_2292_i2).isLessThan(new BigNumber((_2291_v7).length))))) {
          (_2291_v7)[(_2292_i2)] = ((p2) ? (_module.D14.create_DC40(_module.D6.create_DC18(_module.D6.create_DC15(new _dafny.CodePoint('o'.codePointAt(0)))), p0, p2)) : (_module.D14.create_DC40(_module.D6.create_DC18(_module.D6.create_DC15(new _dafny.CodePoint('i'.codePointAt(0)))), _this.f7, true)));
        }
      }
      if ((_this).fm8(p1, globalState)) {
        let _2293_v8;
        _2293_v8 = new BigNumber(754);
        let _2294_v9;
        _2294_v9 = _dafny.Seq.of(_2293_v8);
        _2293_v8 = (_2294_v9)[_module.__default.safeIndex(p1, new BigNumber((_2294_v9).length))];
        let _2295_v10;
        _2295_v10 = _dafny.Seq.of(p0);
        _2293_v8 = new BigNumber((_dafny.Seq.update(_2295_v10, _module.__default.safeIndex((p1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,false)).length)), new BigNumber((_2295_v10).length)), (false) === (p2))).length);
        let _2296_v11;
        _2296_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2297_v12;
        _2297_v12 = _module.D6.create_DC15(_2296_v11);
        _2297_v12 = _module.D6.create_DC15((_2283_v1)[_module.__default.safeIndex(new BigNumber((_2283_v1).length), new BigNumber((_2283_v1).length))]);
        let _2298_v13;
        let _nw388 = Array((new BigNumber(4)).toNumber()).fill(false);
        _2298_v13 = _nw388;
        let _2299_v14;
        _2299_v14 = _dafny.Set.fromElements(_2298_v13, _2298_v13);
        let _2300_v15;
        let _nw389 = new _module.C2();
        _nw389.__ctor(_2299_v14, _2296_v11, p2);
        _2300_v15 = _nw389;
        _2293_v8 = new BigNumber((_dafny.Seq.UnicodeFromString("rkdxg")).length);
      } else {
        let _2301_v16;
        _2301_v16 = new BigNumber(619);
        _2301_v16 = p1;
        let _2302_v17;
        _2302_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2283_v1);
        let _2303_v18;
        _2303_v18 = _dafny.Seq.of(false, p0);
        let _2304_v19;
        _2304_v19 = _module.D0.create_DC0(p1);
        let _2305_v20;
        _2305_v20 = _dafny.Seq.of(_2304_v19, _2304_v19, _2304_v19, _2304_v19, _2304_v19);
        let _2306_v21;
        _2306_v21 = _dafny.Seq.of(p1, _2301_v16);
        let _rhs353 = ((_2301_v16).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_2303_v18)[_module.__default.safeIndex(p1, new BigNumber((_2303_v18).length))])).length))) === ((_this).f9);
        let _rhs354 = ((_this).fm12(p0, _2301_v16, new BigNumber(478), p2, globalState)).minus(p1);
        let _rhs355 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_module.__default.fm1(_dafny.Seq.update(_2305_v20, _module.__default.safeIndex(new BigNumber((_2306_v21).length), new BigNumber((_2305_v20).length)), _module.D0.create_DC0(new BigNumber((_2283_v1).length))), p1, _2301_v16, globalState)), _2303_v18);
        let _rhs356 = ((_2302_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_2283_v1))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2301_v16,_2283_v1));
        let _lhs218 = _this;
        let _lhs219 = _this;
        _lhs218.f7 = _rhs353;
        _2301_v16 = _rhs354;
        _lhs219.f7 = _rhs355;
        _2302_v17 = _rhs356;
        _2301_v16 = new BigNumber(635);
        _2301_v16 = _2301_v16;
        let _2307_v22;
        let _init78 = ((_2308_p1) => function (_2309_i3) {
          return (_2309_i3).plus(_2308_p1);
        })(p1);
        let _nw390 = Array((new BigNumber(14)).toNumber());
        for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw390.length); _i0_78++) {
          _nw390[_i0_78] = _init78(new BigNumber(_i0_78));
        }
        _2307_v22 = _nw390;
        let _index363 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_2307_v22).length));
        (_2307_v22)[_index363] = new BigNumber((_dafny.MultiSet.FromArray(_2306_v21)).cardinality());
        let _index364 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_2307_v22).length));
        (_2307_v22)[_index364] = new BigNumber((((_2302_v17).Merge(_2302_v17)).Merge(_2302_v17)).length);
      }
      if (p0) {
        let _2310_v23;
        _2310_v23 = _module.D13.create_DC35();
        let _2311_v24;
        _2311_v24 = _dafny.Seq.of(_2310_v23);
        let _2312_v25;
        _2312_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-987)), function (_2313_i4) {
          return _module.D13.create_DC35();
        }), _2311_v24),p1);
        let _2314_v26;
        _2314_v26 = new _dafny.CodePoint('u'.codePointAt(0));
        let _2315_v27;
        _2315_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2314_v26);
        _2312_v25 = (_2312_v25).update(_2311_v24, new BigNumber((_2315_v27).length));
        let _2316_v28;
        let _nw391 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
        _2316_v28 = _nw391;
        let _2317_v29;
        _2317_v29 = _dafny.Seq.of(p1, p1);
        let _2318_v30;
        _2318_v30 = _dafny.Set.fromElements(_2317_v29);
        let _index365 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_2316_v28).length));
        (_2316_v28)[_index365] = (_2318_v30).Difference(_2318_v30);
        let _2319_v31;
        _2319_v31 = _dafny.Seq.of(p2, _module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), ((_2320_p1) => function (_2321_i5) {
          return _module.D0.create_DC0(_2320_p1);
        })(p1)), p1, p1, globalState), (_this).f9);
        let _2322_v32;
        _2322_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2317_v29,_2319_v31);
        let _index366 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_2316_v28).length));
        (_2316_v28)[_index366] = (_dafny.Set.fromElements(_2317_v29, _dafny.Seq.of(new BigNumber((_2319_v31).length), p1), _2317_v29)).Intersect(function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of (_2322_v32).Keys.Elements) {
            let _2323_v33 = _compr_50;
            if ((_2322_v32).contains(_2323_v33)) {
              _coll50.add(_2323_v33);
            }
          }
          return _coll50;
        }());
        let _2324_v34;
        let _nw392 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2324_v34 = _nw392;
        let _index367 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length));
        (_2324_v34)[_index367] = p1;
        let _index368 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length));
        (_2324_v34)[_index368] = _module.__default.safeModuloInt(p1, p1);
        let _index369 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length));
        (_2324_v34)[_index369] = p1;
        let _index370 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length));
        (_2324_v34)[_index370] = (((_this).f9) ? (_module.__default.safeDivisionInt((_this).fm9((_this).f9, (_dafny.ZERO).minus((_2324_v34)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length))]), _2314_v26, globalState), (_2324_v34)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length))])) : ((_dafny.ZERO).minus((_2324_v34)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_2324_v34).length))])));
      } else {
        (_this).f7 = p2;
        let _2325_v35;
        _2325_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(769),_dafny.MultiSet.fromElements(_this.f7));
        let _2326_v36;
        _2326_v36 = _dafny.MultiSet.fromElements(!(_this.f7), p0);
        if (!(((_dafny.MultiSet.fromElements((_this).f9, true)).IsProperSubsetOf((((_2325_v35).contains(new BigNumber(-296))) ? ((_2325_v35).get(new BigNumber(-296))) : (_2326_v36)))) || ((_this).f9))) {
          let _2327_v37;
          _2327_v37 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),p1);
          let _2328_v38;
          _2328_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2327_v37);
          let _2329_v39;
          let _nw393 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2329_v39 = _nw393;
          let _2330_v40;
          _2330_v40 = _dafny.Seq.of(_this.f7);
          let _2331_v41;
          _2331_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2329_v39,_2330_v40);
          let _2332_v42;
          _2332_v42 = _dafny.MultiSet.fromElements(new BigNumber(((_2331_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2329_v39,_2330_v40))).length), p1, new BigNumber((((p2) ? (_dafny.MultiSet.fromElements(p2, (_this).f9)) : (_dafny.MultiSet.fromElements(_this.f7, p0, p0, p0, _this.f7)))).cardinality()), p1);
          let _2333_v43;
          _2333_v43 = new _dafny.CodePoint('g'.codePointAt(0));
          let _2334_v44;
          _2334_v44 = _dafny.Set.fromElements(p1);
          let _2335_v45;
          _2335_v45 = _dafny.Set.fromElements(p1, new BigNumber((_2334_v44).length), p1, new BigNumber((_2283_v1).length), p1);
          let _2336_v46;
          _2336_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2283_v1);
          let _2337_v47;
          _2337_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2336_v46).contains(p1)) ? ((_2336_v46).get(p1)) : (_2283_v1))).length),(_this).f9);
          let _2338_v48;
          _2338_v48 = _dafny.Seq.of(new BigNumber((_2283_v1).length));
          let _2339_v49;
          let _nw394 = Array((new BigNumber(15)).toNumber());
          _nw394[(_dafny.ZERO).toNumber()] = p1;
          _nw394[(_dafny.ONE).toNumber()] = p1;
          _nw394[(new BigNumber(2)).toNumber()] = p1;
          _nw394[(new BigNumber(3)).toNumber()] = p1;
          _nw394[(new BigNumber(4)).toNumber()] = p1;
          _nw394[(new BigNumber(5)).toNumber()] = p1;
          _nw394[(new BigNumber(6)).toNumber()] = p1;
          _nw394[(new BigNumber(7)).toNumber()] = p1;
          _nw394[(new BigNumber(8)).toNumber()] = p1;
          _nw394[(new BigNumber(9)).toNumber()] = p1;
          _nw394[(new BigNumber(10)).toNumber()] = p1;
          _nw394[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm0(globalState))).cardinality());
          _nw394[(new BigNumber(12)).toNumber()] = p1;
          _nw394[(new BigNumber(13)).toNumber()] = (_2338_v48)[_module.__default.safeIndex(new BigNumber(548), new BigNumber((_2338_v48).length))];
          _nw394[(new BigNumber(14)).toNumber()] = p1;
          _2339_v49 = _nw394;
          let _2340_v50;
          _2340_v50 = _dafny.Map.Empty.slice().updateUnsafe(false,_2339_v49);
          let _2341_v51;
          let _nw395 = Array((new BigNumber(27)).toNumber());
          _nw395[(_dafny.ZERO).toNumber()] = new BigNumber((_2283_v1).length);
          _nw395[(_dafny.ONE).toNumber()] = p1;
          _nw395[(new BigNumber(2)).toNumber()] = p1;
          _nw395[(new BigNumber(3)).toNumber()] = p1;
          _nw395[(new BigNumber(4)).toNumber()] = p1;
          _nw395[(new BigNumber(5)).toNumber()] = p1;
          _nw395[(new BigNumber(6)).toNumber()] = p1;
          _nw395[(new BigNumber(7)).toNumber()] = p1;
          _nw395[(new BigNumber(8)).toNumber()] = p1;
          _nw395[(new BigNumber(9)).toNumber()] = p1;
          _nw395[(new BigNumber(10)).toNumber()] = p1;
          _nw395[(new BigNumber(11)).toNumber()] = p1;
          _nw395[(new BigNumber(12)).toNumber()] = p1;
          _nw395[(new BigNumber(13)).toNumber()] = p1;
          _nw395[(new BigNumber(14)).toNumber()] = p1;
          _nw395[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.update(_2283_v1, _module.__default.safeIndex(new BigNumber((_2330_v40).length), new BigNumber((_2283_v1).length)), _2333_v43)).length);
          _nw395[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw395[(new BigNumber(17)).toNumber()] = p1;
          _nw395[(new BigNumber(18)).toNumber()] = p1;
          _nw395[(new BigNumber(19)).toNumber()] = p1;
          _nw395[(new BigNumber(20)).toNumber()] = p1;
          _nw395[(new BigNumber(21)).toNumber()] = new BigNumber((_2334_v44).length);
          _nw395[(new BigNumber(22)).toNumber()] = new BigNumber(153);
          _nw395[(new BigNumber(23)).toNumber()] = new BigNumber(((_2332_v42).update(new BigNumber((_dafny.Seq.UnicodeFromString("vkp")).length), _module.__default.abs(new BigNumber((_2337_v47).length)))).cardinality());
          _nw395[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2332_v42).cardinality()));
          _nw395[(new BigNumber(25)).toNumber()] = new BigNumber(441);
          _nw395[(new BigNumber(26)).toNumber()] = new BigNumber((_2340_v50).length);
          _2341_v51 = _nw395;
          let _2342_v52;
          _2342_v52 = _dafny.MultiSet.fromElements(_2341_v51, _2339_v49, _2341_v51);
          let _2343_v53;
          _2343_v53 = new BigNumber(882);
          let _2344_v54;
          _2344_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2341_v51);
          let _2345_v55;
          _2345_v55 = _module.D17.create_DC47(_2342_v52);
          let _rhs357 = (_2328_v38).update(new BigNumber((_2344_v54).length), _2327_v37);
          let _rhs358 = _2332_v42;
          let _rhs359 = (((_2345_v55).dtor_cf76).Difference(_2342_v52)).update(_2341_v51, _module.__default.abs(_2343_v53));
          let _rhs360 = p1;
          _2328_v38 = _rhs357;
          _2332_v42 = _rhs358;
          _2342_v52 = _rhs359;
          _2343_v53 = _rhs360;
          _module.__default.m0(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_2330_v40).length)), _2343_v53), _2343_v53, globalState);
          _2343_v53 = (_2343_v53).multipliedBy(new BigNumber(821));
          let _2346_v56;
          let _nw396 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
          _2346_v56 = _nw396;
          let _2347_v57;
          _2347_v57 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_2346_v56);
          _2347_v57 = (_2347_v57).update(p1, _2346_v56);
          let _2348_v58;
          let _nw397 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2348_v58 = _nw397;
          let _index371 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2348_v58).length));
          (_2348_v58)[_index371] = _module.__default.fm28(p2, _2343_v53, globalState);
          let _index372 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2348_v58).length));
          (_2348_v58)[_index372] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), ((_2349_v43) => function (_2350_i6) {
            return _2349_v43;
          })(_2333_v43));
        } else {
          let _2351_v59;
          _2351_v59 = _dafny.Map.Empty.slice().updateUnsafe((p1).isEqualTo(new BigNumber(323)),new BigNumber(-596));
          _2351_v59 = (_2351_v59).update(_this.f7, new BigNumber(147));
          let _2352_v60;
          let _nw398 = new _module.C3();
          _nw398.__ctor(p1, p2);
          _2352_v60 = _nw398;
          let _2353_v61;
          _2353_v61 = new BigNumber(569);
          let _2354_v62;
          let _nw399 = new _module.C3();
          _nw399.__ctor(p1, p0);
          _2354_v62 = _nw399;
          let _2355_v63;
          _2355_v63 = _module.D18.create_DC50(_2354_v62);
          let _rhs361 = _2352_v60;
          let _rhs362 = p1;
          let _rhs363 = (_2353_v61).isLessThanOrEqualTo(_2353_v61);
          let _rhs364 = (_2355_v63).dtor_cf84;
          let _rhs365 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hulf"), _2283_v1);
          let _lhs220 = _this;
          _2352_v60 = _rhs361;
          _2353_v61 = _rhs362;
          _lhs220.f7 = _rhs363;
          _2354_v62 = _rhs364;
          _2283_v1 = _rhs365;
          let _2356_v64;
          let _nw400 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2356_v64 = _nw400;
          let _2357_v65;
          _2357_v65 = _dafny.Seq.of((_2352_v60).f14, p1);
          let _2358_v66;
          _2358_v66 = _module.D12.create_DC30(_2353_v61, !(_2354_v62.f7), _2357_v65, (_2352_v60).f14);
          let _2359_v67;
          _2359_v67 = _dafny.Seq.of(_2354_v62.f7, p0, (_2358_v66).dtor_cf50, p2, _2354_v62.f7);
          let _2360_v68;
          _2360_v68 = _dafny.Set.fromElements((_2359_v67)[_module.__default.safeIndex(p1, new BigNumber((_2359_v67).length))]);
          let _2361_v69;
          _2361_v69 = _dafny.MultiSet.fromElements(p1, (_2352_v60).f14, (_dafny.ZERO).minus(new BigNumber((_2360_v68).length)), new BigNumber(100), _2353_v61);
          let _2362_v70;
          let _nw401 = Array((new BigNumber(8)).toNumber()).fill(false);
          _2362_v70 = _nw401;
          let _2363_v71;
          _2363_v71 = _dafny.Set.fromElements(_2362_v70);
          let _index373 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_2356_v64).length));
          (_2356_v64)[_index373] = ((((_2361_v69).contains(new BigNumber((_2363_v71).length))) ? ((_2361_v69).get(new BigNumber((_2363_v71).length))) : (p1))).plus(p1);
          let _index374 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_2356_v64).length));
          (_2356_v64)[_index374] = (_dafny.ZERO).minus((_2352_v60).f14);
          (_2354_v62).f7 = !(_2354_v62.f7);
          let _2364_v72;
          _2364_v72 = _dafny.Set.fromElements((_2352_v60).f14, p1, (_2352_v60).f14);
          let _2365_v73;
          _2365_v73 = _dafny.Seq.of(_module.__default.fm2(_2353_v61, _2354_v62.f7, globalState), (_2364_v72).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_2354_v62, _2354_v62)).cardinality()))));
          let _2366_v74;
          let _init79 = ((_2367_v67) => function (_2368_i7) {
            return _2367_v67;
          })(_2359_v67);
          let _nw402 = Array((_dafny.ONE).toNumber());
          for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw402.length); _i0_79++) {
            _nw402[_i0_79] = _init79(new BigNumber(_i0_79));
          }
          _2366_v74 = _nw402;
          let _index375 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_2366_v74).length));
          (_2366_v74)[_index375] = _2359_v67;
          let _2369_v75;
          _2369_v75 = _dafny.MultiSet.fromElements(_2362_v70, _2362_v70, _2362_v70, _2362_v70);
          let _2370_v76;
          _2370_v76 = _2325_v35;
          let _2371_v77;
          _2371_v77 = _dafny.Seq.of(_dafny.Seq.of(p0, (_2352_v60).fm8(_2353_v61, globalState), _this.f7));
          let _index376 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_2366_v74).length));
          let _rhs366 = _module.__default.fm54(globalState);
          let _rhs367 = (_2371_v77)[_module.__default.safeIndex((_2352_v60).f14, new BigNumber((_2371_v77).length))];
          let _rhs368 = _2369_v75;
          let _rhs369 = _2370_v76;
          let _rhs370 = _2356_v64;
          let _lhs221 = _2366_v74;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_2366_v74).length));
          _2365_v73 = _rhs366;
          _lhs221[_lhs222] = _rhs367;
          _2369_v75 = _rhs368;
          _2370_v76 = _rhs369;
          _2356_v64 = _rhs370;
        }
        let _2372_v78;
        let _nw403 = new _module.C3();
        _nw403.__ctor(new BigNumber(962), !(p0));
        _2372_v78 = _nw403;
        let _2373_v79;
        _2373_v79 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2374_v80;
        _2374_v80 = _dafny.Set.fromElements(p2);
        let _2375_v81;
        let _nw404 = new _module.C3();
        _nw404.__ctor(p1, (_2374_v80).IsProperSubsetOf(_module.__default.fm27(_this.f7, _2373_v79, p2, p0, globalState)));
        _2375_v81 = _nw404;
        let _2376_v82;
        let _nw405 = Array((new BigNumber(18)).toNumber()).fill(false);
        _2376_v82 = _nw405;
        let _2377_v83;
        _2377_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2376_v82);
        _2376_v82 = (((_2377_v83).contains(_this.f7)) ? ((_2377_v83).get(_this.f7)) : (_2376_v82));
      }
      let _2378_v84;
      _2378_v84 = _dafny.Set.fromElements(false);
      r1 = (_2378_v84).Difference((_2378_v84).Union(_2378_v84));
      let _2379_v85;
      _2379_v85 = _dafny.Map.Empty.slice().updateUnsafe(false,p2);
      r0 = (_2379_v85).Merge(_2379_v85);
      let _2380_v86;
      _2380_v86 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2381_v87;
      _2381_v87 = _module.D1.create_DC3(_2378_v84);
      r1 = ((_dafny.Seq.contains(_2283_v1, _2380_v86)) ? ((_2381_v87).dtor_cf6) : ((_2378_v84).Intersect(_2378_v84)));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      if (_this.f7) {
        let _2382_v0;
        _2382_v0 = new BigNumber(-390);
        let _2383_v1;
        let _nw406 = Array((new BigNumber(4)).toNumber());
        _nw406[(_dafny.ZERO).toNumber()] = _this.f7;
        _nw406[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw406[(new BigNumber(2)).toNumber()] = _this.f7;
        _nw406[(new BigNumber(3)).toNumber()] = false;
        _2383_v1 = _nw406;
        let _2384_v2;
        _2384_v2 = _dafny.Set.fromElements(_2383_v1, _2383_v1, _2383_v1, _2383_v1, _2383_v1);
        let _2385_v3;
        _2385_v3 = new _dafny.CodePoint('t'.codePointAt(0));
        let _2386_v4;
        let _nw407 = new _module.C2();
        _nw407.__ctor(_2384_v2, _2385_v3, (_this).f9);
        _2386_v4 = _nw407;
        let _2387_v5;
        _2387_v5 = _dafny.Seq.of(_2386_v4);
        let _2388_v6;
        _2388_v6 = _dafny.Seq.of(new BigNumber((_2387_v5).length), new BigNumber(127));
        let _2389_v7;
        let _init80 = ((_2390_v0) => function (_2391_i0) {
          return _module.__default.safeDivisionInt(_2391_i0, _2390_v0);
        })(_2382_v0);
        let _nw408 = Array((new BigNumber(5)).toNumber());
        for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw408.length); _i0_80++) {
          _nw408[_i0_80] = _init80(new BigNumber(_i0_80));
        }
        _2389_v7 = _nw408;
        let _2392_v8;
        _2392_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2382_v0,_2389_v7);
        let _2393_v9;
        _2393_v9 = _dafny.Map.Empty.slice().updateUnsafe((_2382_v0).isLessThan(_2382_v0),_module.__default.safeDivisionInt(new BigNumber((_2388_v6).length), new BigNumber((_2392_v8).length)));
        _2393_v9 = _2393_v9;
        let _2394_v10;
        let _nw409 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2394_v10 = _nw409;
        let _index377 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_2394_v10).length));
        (_2394_v10)[_index377] = ((_this.f7) ? (new _dafny.CodePoint('r'.codePointAt(0))) : (new _dafny.CodePoint('o'.codePointAt(0))));
        let _2395_v11;
        _2395_v11 = _dafny.Seq.of(_2386_v4.f7);
        let _2396_v12;
        _2396_v12 = _dafny.Set.fromElements((_this).f9, _dafny.Seq.IsProperPrefixOf(_2395_v11, _2395_v11));
        let _index378 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_2394_v10).length));
        let _rhs371 = _module.__default.fm13(_2386_v4.f7, _2382_v0, (_this).f9, globalState);
        let _rhs372 = _2396_v12;
        let _lhs223 = _2394_v10;
        let _lhs224 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_2394_v10).length));
        _lhs223[_lhs224] = _rhs371;
        _2396_v12 = _rhs372;
        let _2397_v13;
        _2397_v13 = _dafny.MultiSet.fromElements(_2386_v4.f7);
        let _2398_v14;
        _2398_v14 = _module.D9.create_DC26(_dafny.Seq.of(new BigNumber((_2397_v13).cardinality())));
        let _pat_let_tv49 = _2388_v6;
        let _rhs373 = (_this).f9;
        let _rhs374 = (_this).f9;
        let _rhs375 = function (_pat_let67_0) {
          return function (_2399_dt__update__tmp_h0) {
            return function (_pat_let68_0) {
              return function (_2400_dt__update_hcf45_h0) {
                return _module.D9.create_DC26(_2400_dt__update_hcf45_h0);
              }(_pat_let68_0);
            }(_pat_let_tv49);
          }(_pat_let67_0);
        }(_2398_v14);
        let _rhs376 = _dafny.Seq.Concat(_2395_v11, _2395_v11);
        let _lhs225 = _this;
        let _lhs226 = _2386_v4;
        _lhs225.f7 = _rhs373;
        _lhs226.f7 = _rhs374;
        _2398_v14 = _rhs375;
        _2395_v11 = _rhs376;
        let _2401_v15;
        let _nw410 = new _module.C5();
        _nw410.__ctor(_2386_v4.f7);
        _2401_v15 = _nw410;
        let _2402_v16;
        _2402_v16 = _dafny.Set.fromElements(_2382_v0, _2382_v0);
        let _2403_v17;
        _2403_v17 = _module.D2.create_DC7(_2402_v16);
        let _2404_v18;
        _2404_v18 = _dafny.Seq.of(_2402_v16);
        let _pat_let_tv50 = _2402_v16;
        let _pat_let_tv51 = _2402_v16;
        let _2405_v19;
        let _nw411 = Array((new BigNumber(19)).toNumber());
        _nw411[(_dafny.ZERO).toNumber()] = _2403_v17;
        _nw411[(_dafny.ONE).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(2)).toNumber()] = _module.D2.create_DC7((_2404_v18)[_module.__default.safeIndex(_2382_v0, new BigNumber((_2404_v18).length))]);
        _nw411[(new BigNumber(3)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(4)).toNumber()] = function (_pat_let69_0) {
          return function (_2406_dt__update__tmp_h1) {
            return function (_pat_let70_0) {
              return function (_2407_dt__update_hcf13_h0) {
                return _module.D2.create_DC7(_2407_dt__update_hcf13_h0);
              }(_pat_let70_0);
            }(_pat_let_tv50);
          }(_pat_let69_0);
        }(_module.D2.create_DC7(_2402_v16));
        _nw411[(new BigNumber(5)).toNumber()] = function (_pat_let71_0) {
          return function (_2408_dt__update__tmp_h2) {
            return function (_pat_let72_0) {
              return function (_2409_dt__update_hcf13_h1) {
                return _module.D2.create_DC7(_2409_dt__update_hcf13_h1);
              }(_pat_let72_0);
            }(_pat_let_tv51);
          }(_pat_let71_0);
        }(_module.D2.create_DC7(_2402_v16));
        _nw411[(new BigNumber(6)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(7)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(8)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(9)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(10)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(11)).toNumber()] = _module.D2.create_DC7(_2402_v16);
        _nw411[(new BigNumber(12)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(13)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(14)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(15)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(16)).toNumber()] = _module.D2.create_DC7(_dafny.Set.fromElements(_2382_v0, _2382_v0, _2382_v0, _2382_v0, _2382_v0));
        _nw411[(new BigNumber(17)).toNumber()] = _2403_v17;
        _nw411[(new BigNumber(18)).toNumber()] = _module.D2.create_DC7(_2402_v16);
        _2405_v19 = _nw411;
        let _2410_v20;
        let _nw412 = new _module.C4();
        _nw412.__ctor(_2405_v19, (_this).f9);
        _2410_v20 = _nw412;
        let _2411_v21;
        _2411_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2410_v20);
        let _2412_v22;
        _2412_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2385_v3,_2410_v20);
        let _2413_v23;
        _2413_v23 = _dafny.Seq.of((((_2412_v22).contains(_2385_v3)) ? ((_2412_v22).get(_2385_v3)) : (_2410_v20)), _2410_v20);
        let _2414_v24;
        let _nw413 = Array((new BigNumber(22)).toNumber());
        _nw413[(_dafny.ZERO).toNumber()] = _2410_v20;
        _nw413[(_dafny.ONE).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(2)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(3)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(4)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(5)).toNumber()] = (((_2411_v21).contains((_this).f9)) ? ((_2411_v21).get((_this).f9)) : (_2410_v20));
        _nw413[(new BigNumber(6)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(7)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(8)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(9)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(10)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(11)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(12)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(13)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(14)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(15)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(16)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(17)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(18)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(19)).toNumber()] = _2410_v20;
        _nw413[(new BigNumber(20)).toNumber()] = (((_this).f9) ? (_2410_v20) : (_2410_v20));
        _nw413[(new BigNumber(21)).toNumber()] = (_2413_v23)[_module.__default.safeIndex(_2382_v0, new BigNumber((_2413_v23).length))];
        _2414_v24 = _nw413;
        let _index379 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_2414_v24).length));
        (_2414_v24)[_index379] = (_2413_v23)[_module.__default.safeIndex(_2382_v0, new BigNumber((_2413_v23).length))];
        let _index380 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_2414_v24).length));
        (_2414_v24)[_index380] = _2410_v20;
      } else {
        let _2415_v25;
        let _init81 = function (_2416_i1) {
          return _module.D2.create_DC7(_dafny.Set.fromElements(new BigNumber(-717)));
        };
        let _nw414 = Array((new BigNumber(6)).toNumber());
        for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw414.length); _i0_81++) {
          _nw414[_i0_81] = _init81(new BigNumber(_i0_81));
        }
        _2415_v25 = _nw414;
        let _2417_v26;
        let _nw415 = new _module.C4();
        _nw415.__ctor(_2415_v25, (!(true)) || (false));
        _2417_v26 = _nw415;
        let _2418_v27;
        let _nw416 = Array((new BigNumber(16)).toNumber()).fill(false);
        _2418_v27 = _nw416;
        let _2419_v28;
        _2419_v28 = _dafny.Set.fromElements(_2418_v27);
        let _2420_v29;
        _2420_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2418_v27);
        let _2421_v30;
        let _nw417 = Array((new BigNumber(6)).toNumber());
        _nw417[(_dafny.ZERO).toNumber()] = _2419_v28;
        _nw417[(_dafny.ONE).toNumber()] = (_2419_v28).Difference(_2419_v28);
        _nw417[(new BigNumber(2)).toNumber()] = _2419_v28;
        _nw417[(new BigNumber(3)).toNumber()] = _2419_v28;
        _nw417[(new BigNumber(4)).toNumber()] = (_2419_v28).Intersect(_2419_v28);
        _nw417[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_2418_v27, _2418_v27, _2418_v27, (((_2420_v29).contains(!((_this).f9))) ? ((_2420_v29).get(!((_this).f9))) : (_2418_v27)));
        _2421_v30 = _nw417;
        let _index381 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_2421_v30).length));
        (_2421_v30)[_index381] = _2419_v28;
        let _index382 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_2421_v30).length));
        (_2421_v30)[_index382] = _2419_v28;
        let _2422_v31;
        _2422_v31 = new _dafny.CodePoint('d'.codePointAt(0));
        let _2423_v32;
        _2423_v32 = new BigNumber(43);
        let _2424_v33;
        _2424_v33 = _dafny.Seq.of(_2423_v32);
        let _2425_v34;
        _2425_v34 = _dafny.MultiSet.fromElements(_this.f7, !(_this.f7), (_this).f9);
        let _2426_v35;
        _2426_v35 = _dafny.Seq.of(_2422_v31);
        let _2427_v36;
        _2427_v36 = _module.D0.create_DC0(_2423_v32);
        let _2428_v37;
        _2428_v37 = _dafny.Seq.of(_2427_v36, _2427_v36);
        let _2429_v38;
        _2429_v38 = _dafny.MultiSet.fromElements(new BigNumber(623));
        let _2430_v39;
        let _nw418 = Array((new BigNumber(15)).toNumber());
        _nw418[(_dafny.ZERO).toNumber()] = _2423_v32;
        _nw418[(_dafny.ONE).toNumber()] = (_2423_v32).multipliedBy(new BigNumber((_2424_v33).length));
        _nw418[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((((_2425_v34).contains(_this.f7)) ? ((_2425_v34).get(_this.f7)) : (new BigNumber(985))));
        _nw418[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("caerf")).length);
        _nw418[(new BigNumber(4)).toNumber()] = (new BigNumber(841)).minus(new BigNumber((_2426_v35).length));
        _nw418[(new BigNumber(5)).toNumber()] = _2423_v32;
        _nw418[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_2423_v32, new BigNumber((_2425_v34).cardinality()));
        _nw418[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).fm12((_this).fm8(_2423_v32, globalState), _2423_v32, _2423_v32, _module.__default.fm1(_2428_v37, _2423_v32, _2423_v32, globalState), globalState))).length);
        _nw418[(new BigNumber(8)).toNumber()] = _2423_v32;
        _nw418[(new BigNumber(9)).toNumber()] = _2423_v32;
        _nw418[(new BigNumber(10)).toNumber()] = new BigNumber((((_2429_v38).update(new BigNumber((_2426_v35).length), _module.__default.abs(new BigNumber(387)))).Union((_dafny.MultiSet.fromElements(new BigNumber((_2429_v38).cardinality()))).update(_2423_v32, _module.__default.abs((_this).fm12((_this).f9, _2423_v32, new BigNumber((_2429_v38).cardinality()), !((_this).f9), globalState))))).cardinality());
        _nw418[(new BigNumber(11)).toNumber()] = (((_2425_v34).contains((_this).f9)) ? ((_2425_v34).get((_this).f9)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-608)), function (_2431_i2) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        })).length)));
        _nw418[(new BigNumber(12)).toNumber()] = _2423_v32;
        _nw418[(new BigNumber(13)).toNumber()] = ((_dafny.ZERO).minus(_2423_v32)).minus(_2423_v32);
        _nw418[(new BigNumber(14)).toNumber()] = new BigNumber(971);
        _2430_v39 = _nw418;
        let _index383 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
        (_2430_v39)[_index383] = _module.__default.safeDivisionInt(_2423_v32, _2423_v32);
        let _index384 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
        let _rhs377 = _module.__default.fm13(((_this.f7) ? (!(_this.f7)) : (false)), _2423_v32, _this.f7, globalState);
        let _rhs378 = _this.f7;
        let _rhs379 = ((new BigNumber((_dafny.Seq.of(_this.f7)).length)).multipliedBy(_module.__default.fm0(globalState))).plus(new BigNumber(-956));
        let _rhs380 = (!((_this).f9)) && ((_this).f9);
        let _rhs381 = (_this).f9;
        let _lhs227 = _this;
        let _lhs228 = _2430_v39;
        let _lhs229 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
        let _lhs230 = _this;
        let _lhs231 = _this;
        _2422_v31 = _rhs377;
        _lhs227.f7 = _rhs378;
        _lhs228[_lhs229] = _rhs379;
        _lhs230.f7 = _rhs380;
        _lhs231.f7 = _rhs381;
        if (_this.f7) {
          let _2432_v40;
          _2432_v40 = _dafny.Seq.of(_this.f7);
          let _2433_v41;
          _2433_v41 = _dafny.Set.fromElements(_this.f7, (_2432_v40)[_module.__default.safeIndex(new BigNumber(903), new BigNumber((_2432_v40).length))], (_this).f9);
          let _2434_v42;
          _2434_v42 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2433_v41);
          let _index385 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
          (_2430_v39)[_index385] = (_2423_v32).minus((_dafny.ZERO).minus(new BigNumber((_2434_v42).length)));
          let _2435_v43;
          _2435_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_2430_v39)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length))]);
          let _2436_v44;
          let _nw419 = new _module.C3();
          _nw419.__ctor(new BigNumber(((_2435_v43).Merge(_module.__default.fm43(new _dafny.CodePoint('k'.codePointAt(0)), (_this).f9, _this.f7, globalState))).length), _this.f7);
          _2436_v44 = _nw419;
          let _index386 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
          (_2430_v39)[_index386] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_2423_v32))));
          (_this).f7 = (_this).fm8((_2436_v44).f14, globalState);
          (_this).f7 = (_2423_v32).isLessThan((_2423_v32).plus(new BigNumber(918)));
        } else {
          _2426_v35 = _2426_v35;
          let _2437_v45;
          let _nw420 = new _module.C5();
          _nw420.__ctor((_this).f9);
          _2437_v45 = _nw420;
          (_this).f7 = ((!(!((_2437_v45).f10)) || (_this.f7)) ? (false) : ((_2437_v45).f10));
          _2423_v32 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hpuhsuu"), (((_2437_v45).f10) ? (_dafny.Seq.UnicodeFromString("rkajlx")) : (_dafny.Seq.UnicodeFromString("aj"))))).length);
          let _index387 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length));
          (_2430_v39)[_index387] = (_2430_v39)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length))];
        }
        _module.__default.m0((_2430_v39)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length))], (_2430_v39)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_2430_v39).length))], globalState);
      }
      let _2438_v46;
      _2438_v46 = new BigNumber(91);
      let _2439_i3;
      _2439_i3 = _dafny.ZERO;
      L16: {
        while (!(_2438_v46).isEqualTo(_2438_v46)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2439_i3)) {
              break L16;
            }
            _2439_i3 = (_2439_i3).plus(_dafny.ONE);
            let _2440_v47;
            _2440_v47 = new _dafny.CodePoint('s'.codePointAt(0));
            let _2441_v48;
            _2441_v48 = _dafny.Seq.of(_2440_v47, _2440_v47);
            let _2442_v49;
            _2442_v49 = _dafny.Seq.of(_2438_v46, new BigNumber((_dafny.Seq.of(_2438_v46)).length), _2438_v46, _2438_v46, new BigNumber(-967));
            let _2443_v50;
            _2443_v50 = _dafny.Seq.of((_this).f9);
            let _2444_v51;
            let _nw421 = Array((new BigNumber(10)).toNumber());
            _nw421[(_dafny.ZERO).toNumber()] = _this.f7;
            _nw421[(_dafny.ONE).toNumber()] = !(_this.f7);
            _nw421[(new BigNumber(2)).toNumber()] = !(_this.f7) || ((_this).f9);
            _nw421[(new BigNumber(3)).toNumber()] = (_this).f9;
            _nw421[(new BigNumber(4)).toNumber()] = !(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), ((_2445_v47) => function (_2446_i4) {
              return _2445_v47;
            })(_2440_v47)), (_2441_v48)[_module.__default.safeIndex(new BigNumber((_2442_v49).length), new BigNumber((_2441_v48).length))]));
            _nw421[(new BigNumber(5)).toNumber()] = _this.f7;
            _nw421[(new BigNumber(6)).toNumber()] = _dafny.Seq.contains(_2441_v48, _2440_v47);
            _nw421[(new BigNumber(7)).toNumber()] = _this.f7;
            _nw421[(new BigNumber(8)).toNumber()] = !((_this).f9);
            _nw421[(new BigNumber(9)).toNumber()] = (new BigNumber((_2443_v50).length)).isEqualTo((_dafny.ZERO).minus(_2438_v46));
            _2444_v51 = _nw421;
            let _index388 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_2444_v51).length));
            (_2444_v51)[_index388] = _this.f7;
            let _index389 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_2444_v51).length));
            (_2444_v51)[_index389] = _this.f7;
            _2438_v46 = new BigNumber(625);
            _2438_v46 = _module.__default.safeModuloInt((((_this).f9) ? (new BigNumber(314)) : (_2438_v46)), new BigNumber(-620));
            _2440_v47 = _2440_v47;
          }
        }
      }
      let _2447_v52;
      let _init82 = function (_2448_i5) {
        return _this.f7;
      };
      let _nw422 = Array((new BigNumber(28)).toNumber());
      for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw422.length); _i0_82++) {
        _nw422[_i0_82] = _init82(new BigNumber(_i0_82));
      }
      _2447_v52 = _nw422;
      let _2449_v53;
      _2449_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2447_v52,_2438_v46);
      _2449_v53 = (_2449_v53).update(_2447_v52, _2438_v46);
      let _2450_v54;
      _2450_v54 = _dafny.Seq.UnicodeFromString("fnenprj");
      let _2451_v55;
      _2451_v55 = _dafny.Seq.of(_2438_v46);
      let _2452_v56;
      _2452_v56 = _module.D9.create_DC26(_2451_v55);
      let _pat_let_tv52 = _2450_v54;
      let _pat_let_tv53 = _2450_v54;
      let _pat_let_tv54 = _2450_v54;
      let _pat_let_tv55 = _2450_v54;
      _2450_v54 = function (_source17) {
        if (_source17.is_DC24) {
          return _pat_let_tv52;
        } else if (_source17.is_DC25) {
          return _pat_let_tv53;
        } else if (_source17.is_DC26) {
          let _2453___mcc_h0 = (_source17).cf45;
          let _2454_cf45 = _2453___mcc_h0;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), function (_2455_i6) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          });
        } else {
          let _2456___mcc_h1 = (_source17).cf44;
          let _2457_cf44 = _2456___mcc_h1;
          return _dafny.Seq.Concat(_pat_let_tv54, _pat_let_tv55);
        }
      }(_2452_v56);
      let _2458_v57;
      let _nw423 = new _module.C3();
      _nw423.__ctor(((_this.f7) ? (_2438_v46) : (_2438_v46)), !((_this).f9));
      _2458_v57 = _nw423;
      let _hi13 = new BigNumber((_dafny.Seq.of(_2438_v46)).length);
      for (let _2459_i7 = (_2438_v46).multipliedBy(_2438_v46); _2459_i7.isLessThan(_hi13); _2459_i7 = _2459_i7.plus(_dafny.ONE)) {
        if (((_2458_v57).f14).isLessThanOrEqualTo((_2458_v57).f14)) {
          _2438_v46 = _2459_i7;
          _2438_v46 = _2459_i7;
          let _2460_v58;
          _2460_v58 = _module.D0.create_DC0(_2438_v46);
          let _2461_v59;
          _2461_v59 = _dafny.Seq.of(_2460_v58, _2460_v58);
          let _2462_v60;
          _2462_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2450_v54,_2459_i7);
          (_this).f7 = ((_module.__default.fm1(_2461_v59, new BigNumber((_2462_v60).length), _2438_v46, globalState)) ? ((_this).f9) : ((_this).f9));
          let _2463_v61;
          let _nw424 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2463_v61 = _nw424;
          let _index390 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_2463_v61).length));
          (_2463_v61)[_index390] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f9), _dafny.Seq.of((_this).f9))).length);
          let _index391 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_2463_v61).length));
          (_2463_v61)[_index391] = new BigNumber((_2450_v54).length);
          _2438_v46 = ((_2458_v57).f14).multipliedBy((_2458_v57).f14);
        } else {
          let _2464_v62;
          _2464_v62 = _dafny.Seq.of(_this.f7, _this.f7);
          let _2465_v63;
          _2465_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(729),new BigNumber((_2464_v62).length));
          let _2466_v64;
          _2466_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2465_v63,(_2459_i7).minus((_2458_v57).f14));
          _2466_v64 = (_dafny.Map.Empty.slice().updateUnsafe(_2465_v63,new BigNumber((_dafny.Seq.UnicodeFromString("mdo")).length))).update(_dafny.Map.Empty.slice().updateUnsafe(_2438_v46,(_2458_v57).f14), (_2451_v55)[_module.__default.safeIndex((_2458_v57).f14, new BigNumber((_2451_v55).length))]);
          (_this).f7 = (_this).fm8((_2458_v57).f14, globalState);
          (_this).f7 = _this.f7;
          let _2467_v65;
          let _init83 = ((_2468_v57) => function (_2469_i8) {
            return (_2469_i8).multipliedBy((_2468_v57).f14);
          })(_2458_v57);
          let _nw425 = Array((new BigNumber(9)).toNumber());
          for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw425.length); _i0_83++) {
            _nw425[_i0_83] = _init83(new BigNumber(_i0_83));
          }
          _2467_v65 = _nw425;
          let _index392 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_2467_v65).length));
          (_2467_v65)[_index392] = _2459_i7;
          let _index393 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_2467_v65).length));
          (_2467_v65)[_index393] = new BigNumber(857);
          let _2470_v66;
          _2470_v66 = new _dafny.CodePoint('p'.codePointAt(0));
          let _2471_v67;
          _2471_v67 = _dafny.Set.fromElements(_2470_v66);
          let _2472_v68;
          _2472_v68 = _dafny.Set.fromElements(_2450_v54);
          let _2473_v69;
          let _nw426 = Array((new BigNumber(20)).toNumber());
          _nw426[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f7, (_this).f9, (_this).f9, (_this).f9), _2464_v62);
          _nw426[(_dafny.ONE).toNumber()] = _module.__default.fm50(_2471_v67, _2472_v68, new BigNumber((_dafny.MultiSet.fromElements(_module.D14.create_DC39(_dafny.Seq.UnicodeFromString("j"), _2467_v65))).cardinality()), globalState);
          _nw426[(new BigNumber(2)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(3)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(4)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(5)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(6)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(true), _2464_v62);
          _nw426[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_2464_v62, _2464_v62);
          _nw426[(new BigNumber(9)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(10)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(11)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_2464_v62, _dafny.Seq.of(_this.f7, (_this).f9));
          _nw426[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_2464_v62, _dafny.Seq.of((_this).f9));
          _nw426[(new BigNumber(14)).toNumber()] = _module.__default.fm50(_2471_v67, _2472_v68, new BigNumber((_dafny.Set.fromElements(_this.f7)).length), globalState);
          _nw426[(new BigNumber(15)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(16)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(17)).toNumber()] = _2464_v62;
          _nw426[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_this.f7);
          _nw426[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_2464_v62, _2464_v62);
          _2473_v69 = _nw426;
          let _index394 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_2473_v69).length));
          (_2473_v69)[_index394] = _2464_v62;
          let _index395 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_2473_v69).length));
          (_2473_v69)[_index395] = _dafny.Seq.of(_this.f7, (_this).f9, (_this).f9);
        }
        let _2474_v70;
        let _init84 = ((_2475_v57) => function (_2476_i9) {
          return _dafny.Set.fromElements((_2475_v57).f14);
        })(_2458_v57);
        let _nw427 = Array((new BigNumber(25)).toNumber());
        for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw427.length); _i0_84++) {
          _nw427[_i0_84] = _init84(new BigNumber(_i0_84));
        }
        _2474_v70 = _nw427;
        let _index396 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2474_v70).length));
        (_2474_v70)[_index396] = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_2458_v57).f14, (_2458_v57).f14)).cardinality()), new BigNumber(484));
        let _2477_v71;
        _2477_v71 = _dafny.Set.fromElements(_2459_i7);
        let _index397 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2474_v70).length));
        (_2474_v70)[_index397] = _2477_v71;
        let _2478_v72;
        _2478_v72 = _dafny.Seq.of(_this.f7);
        let _2479_v73;
        _2479_v73 = _dafny.Map.Empty.slice().updateUnsafe((_2458_v57).f14,_2478_v72);
        let _2480_v74;
        _2480_v74 = _dafny.Seq.of(_dafny.Seq.of(_this.f7), _dafny.Seq.Concat((((_2479_v73).contains((_2458_v57).f14)) ? ((_2479_v73).get((_2458_v57).f14)) : (_2478_v72)), _2478_v72), _2478_v72);
        _2480_v74 = _2480_v74;
        let _2481_v75;
        _2481_v75 = new _dafny.CodePoint('b'.codePointAt(0));
        (_this).f7 = !_dafny.Seq.contains(_dafny.Seq.Concat(_2450_v54, _2450_v54), _2481_v75);
      }
      return;
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f7 = false;
      this._f8 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f8, f7) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f7 = f7;
      return;
    }
    fm8(p0, globalState) {
      let _this = this;
      return _this.f7;
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-831), new BigNumber((function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(628), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f7)).length)))).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, false, _this.f7)).cardinality()))).length))).Keys.Elements) {
          let _2482_v0 = _compr_51;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(628), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f7)).length)))).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, false, _this.f7)).cardinality()))).length))).contains(_2482_v0)) {
            _coll51.push([(_2482_v0).minus(new BigNumber((function () {
              let _coll52 = new _dafny.Map();
              for (const _compr_52 of _dafny.IntegerRange(new BigNumber(573), new BigNumber(822))) {
                let _2483_v1 = _compr_52;
                if (((new BigNumber(573)).isLessThanOrEqualTo(_2483_v1)) && ((_2483_v1).isLessThan(new BigNumber(822)))) {
                  _coll52.push([(_2483_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ugnal")).length)),_this.f7]);
                }
              }
              return _coll52;
            }()).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f7, _this.f7)).length), new BigNumber(734))).length)]);
          }
        }
        return _coll51;
      }()).length)))).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hyvc"), _dafny.Seq.UnicodeFromString("f"))).length));
    };
    fm10(p0, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(485), new BigNumber(-242))).cardinality())).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,!(_this.f7))).length)));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Set.Empty;
      let _hi14 = (_module.__default.fm0(globalState)).minus((_this).fm10(false, globalState));
      for (let _2484_i0 = p1; _2484_i0.isLessThan(_hi14); _2484_i0 = _2484_i0.plus(_dafny.ONE)) {
        let _2485_v0;
        _2485_v0 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), ((_2486_p1) => function (_2487_i1) {
          return _2486_p1;
        })(p1)),_this.f7);
        let _2488_v1;
        _2488_v1 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,p0)).length));
        let _2489_v2;
        _2489_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2484_i0,_module.__default.fm11((((_2485_v0).contains(_2488_v1)) ? ((_2485_v0).get(_2488_v1)) : (false)), _this.f7, _2484_i0, _this.f7, globalState));
        _2489_v2 = (_2489_v2).update(p1, _dafny.Map.Empty.slice().updateUnsafe(p2,!(true)));
        if (_this.f7) {
          let _2490_v3;
          _2490_v3 = new BigNumber(565);
          let _2491_v4;
          _2491_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2490_v3);
          let _2492_v5;
          _2492_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,_2490_v3),p2);
          _2490_v3 = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_2491_v4,p0)).Merge(_2492_v5)).update(_2491_v4, false)).length);
          _2490_v3 = _2490_v3;
          _2490_v3 = _module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus((p1).plus(p1)));
          (_this).f7 = _this.f7;
          _2490_v3 = p1;
        } else {
          (_this).f7 = !(_this.f7);
          let _2493_v6;
          _2493_v6 = _dafny.MultiSet.fromElements(_2484_i0);
          let _2494_v7;
          _2494_v7 = _dafny.MultiSet.fromElements(false, p0, true, _this.f7);
          let _2495_v8;
          _2495_v8 = _dafny.Seq.UnicodeFromString("nmrn");
          let _2496_v9;
          let _nw428 = Array((new BigNumber(10)).toNumber());
          _nw428[(_dafny.ZERO).toNumber()] = p1;
          _nw428[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p2, _this.f7, _this.f7, _this.f7, p2)).length));
          _nw428[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(_2484_i0));
          _nw428[(new BigNumber(3)).toNumber()] = new BigNumber((((_this.f7) ? (_2493_v6) : (_2493_v6))).cardinality());
          _nw428[(new BigNumber(4)).toNumber()] = new BigNumber(888);
          _nw428[(new BigNumber(5)).toNumber()] = ((p2) ? (_2484_i0) : (_2484_i0));
          _nw428[(new BigNumber(6)).toNumber()] = new BigNumber((_2494_v7).cardinality());
          _nw428[(new BigNumber(7)).toNumber()] = (new BigNumber(29)).plus(_2484_i0);
          _nw428[(new BigNumber(8)).toNumber()] = _2484_i0;
          _nw428[(new BigNumber(9)).toNumber()] = new BigNumber((_2495_v8).length);
          _2496_v9 = _nw428;
          let _2497_v10;
          let _nw429 = Array((new BigNumber(27)).toNumber()).fill(false);
          _2497_v10 = _nw429;
          let _index398 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2497_v10).length));
          (_2497_v10)[_index398] = _this.f7;
          let _index399 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2497_v10).length));
          let _rhs382 = _2496_v9;
          let _rhs383 = p0;
          let _lhs232 = _2497_v10;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2497_v10).length));
          _2496_v9 = _rhs382;
          _lhs232[_lhs233] = _rhs383;
          let _2498_v11;
          _2498_v11 = new _dafny.CodePoint('g'.codePointAt(0));
          _2498_v11 = _2498_v11;
          (_this).m5(globalState);
          let _2499_v12;
          _2499_v12 = new BigNumber(455);
          _2499_v12 = _2499_v12;
        }
        (_this).f7 = p2;
        let _2500_v13;
        let _nw430 = Array((new BigNumber(19)).toNumber()).fill(_module.D1.Default());
        _2500_v13 = _nw430;
        let _2501_v14;
        _2501_v14 = _module.D1.create_DC6(_module.D1.create_DC5(p0, _dafny.Seq.of(_module.__default.fm0(globalState)), _2484_i0));
        let _2502_v15;
        _2502_v15 = _module.D1.create_DC6(_2501_v14);
        let _index400 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2500_v13).length));
        (_2500_v13)[_index400] = _2502_v15;
        let _index401 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2500_v13).length));
        (_2500_v13)[_index401] = _2502_v15;
      }
      let _hi15 = p1;
      for (let _2503_i2 = p1; _2503_i2.isLessThan(_hi15); _2503_i2 = _2503_i2.plus(_dafny.ONE)) {
        let _2504_v16;
        _2504_v16 = _dafny.Seq.UnicodeFromString("fxlh");
        _2504_v16 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nsxh"), _2504_v16);
        let _2505_v17;
        let _nw431 = new _module.C5();
        _nw431.__ctor((p1).isLessThanOrEqualTo(_2503_i2));
        _2505_v17 = _nw431;
        if ((_2505_v17).f10) {
          let _2506_v18;
          _2506_v18 = new _dafny.CodePoint('y'.codePointAt(0));
          let _2507_v19;
          _2507_v19 = _dafny.Seq.of(_this.f7, p2);
          let _2508_v20;
          _2508_v20 = _module.D13.create_DC36(_2506_v18, _2507_v19, _this.f7);
          let _2509_v21;
          let _nw432 = new _module.C6();
          _nw432.__ctor((_2508_v20).dtor_cf60, ((p0) ? ((_2505_v17).f10) : (p2)));
          _2509_v21 = _nw432;
          _2509_v21 = _2509_v21;
          let _2510_v22;
          _2510_v22 = _dafny.Set.fromElements(_module.__default.fm21(p0, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), ((_2511_p1) => function (_2512_i3) {
            return _2511_p1;
          })(p1)));
          _2510_v22 = _2510_v22;
          _2504_v16 = _2504_v16;
          let _2513_v23;
          let _nw433 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _2513_v23 = _nw433;
          let _index402 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2513_v23).length));
          (_2513_v23)[_index402] = new BigNumber(450);
          let _2514_v24;
          let _nw434 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _2514_v24 = _nw434;
          let _2515_v25;
          _2515_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2506_v18,_2514_v24);
          let _2516_v26;
          _2516_v26 = _2514_v24;
          let _2517_v27;
          _2517_v27 = _dafny.Seq.of(_2515_v25, _2515_v25, _2515_v25, _dafny.Map.Empty.slice().updateUnsafe(_2506_v18,_2516_v26));
          let _index403 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2513_v23).length));
          (_2513_v23)[_index403] = new BigNumber(((((_2517_v27)[_module.__default.safeIndex(_2503_i2, new BigNumber((_2517_v27).length))]).Merge((_2515_v25).update(_2506_v18, _2514_v24))).Merge(_2515_v25)).length);
          let _index404 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2513_v23).length));
          (_2513_v23)[_index404] = p1;
        } else {
          let _2518_v28;
          let _init85 = function (_2519_i4) {
            return _module.__default.safeDivisionInt(_2519_i4, new BigNumber(93));
          };
          let _nw435 = Array((new BigNumber(17)).toNumber());
          for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw435.length); _i0_85++) {
            _nw435[_i0_85] = _init85(new BigNumber(_i0_85));
          }
          _2518_v28 = _nw435;
          let _index405 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_2518_v28).length));
          (_2518_v28)[_index405] = _2503_i2;
          let _index406 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_2518_v28).length));
          (_2518_v28)[_index406] = _module.__default.safeDivisionInt(((_dafny.ZERO).minus(p1)).multipliedBy(_2503_i2), (new BigNumber((_dafny.Seq.UnicodeFromString("oxawko")).length)).minus(_2503_i2));
          let _2520_v29;
          let _nw436 = Array((new BigNumber(9)).toNumber());
          _nw436[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw436[(_dafny.ONE).toNumber()] = p2;
          _nw436[(new BigNumber(2)).toNumber()] = (_2505_v17).f10;
          _nw436[(new BigNumber(3)).toNumber()] = p0;
          _nw436[(new BigNumber(4)).toNumber()] = (_2505_v17).f10;
          _nw436[(new BigNumber(5)).toNumber()] = (_2505_v17).f10;
          _nw436[(new BigNumber(6)).toNumber()] = ((_2505_v17).f10) === (p0);
          _nw436[(new BigNumber(7)).toNumber()] = !_dafny.areEqual(_2504_v16, _2504_v16);
          _nw436[(new BigNumber(8)).toNumber()] = p2;
          _2520_v29 = _nw436;
          _2520_v29 = ((p0) ? (_2520_v29) : (_2520_v29));
          (_this).f7 = _this.f7;
          (_2505_v17).m7(globalState);
          let _2521_v30;
          _2521_v30 = new _dafny.CodePoint('a'.codePointAt(0));
          _2521_v30 = _2521_v30;
        }
        let _2522_v31;
        _2522_v31 = new BigNumber(753);
        _2522_v31 = p1;
      }
      let _2523_v32;
      _2523_v32 = _dafny.Seq.UnicodeFromString("nbgnqbndj");
      _2523_v32 = ((p0) ? (_dafny.Seq.Concat(_2523_v32, _2523_v32)) : (_dafny.Seq.Concat(_2523_v32, _dafny.Seq.UnicodeFromString("ngt"))));
      let _2524_v33;
      _2524_v33 = _module.D0.create_DC0(p1);
      let _2525_v34;
      _2525_v34 = _dafny.Seq.of(_2524_v33);
      let _2526_v35;
      _2526_v35 = new _dafny.CodePoint('e'.codePointAt(0));
      let _2527_v36;
      _2527_v36 = _dafny.Seq.of(_2523_v32, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_2528_v32, _2529_p1) => function (_2530_i5) {
        return (_2528_v32)[_module.__default.safeIndex(_2529_p1, new BigNumber((_2528_v32).length))];
      })(_2523_v32, p1)), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_2525_v34, p1, p1, globalState),_this.f7)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_2531_v32, _2532_p1) => function (_2533_i5) {
        return (_2531_v32)[_module.__default.safeIndex(_2532_p1, new BigNumber((_2531_v32).length))];
      })(_2523_v32, p1))).length)), _2526_v35), _dafny.Seq.UnicodeFromString("ojoyengg"), _dafny.Seq.Concat(_2523_v32, _2523_v32), _dafny.Seq.Concat(_2523_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_2534_v35) => function (_2535_i6) {
        return _2534_v35;
      })(_2526_v35))));
      _2523_v32 = (_2527_v36)[_module.__default.safeIndex(p1, new BigNumber((_2527_v36).length))];
      let _2536_i7;
      _2536_i7 = _dafny.ZERO;
      L17: {
        while ((p1).isLessThanOrEqualTo(p1)) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2536_i7)) {
              break L17;
            }
            _2536_i7 = (_2536_i7).plus(_dafny.ONE);
            let _2537_v37;
            let _init86 = function (_2538_i8) {
              return false;
            };
            let _nw437 = Array((new BigNumber(9)).toNumber());
            for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw437.length); _i0_86++) {
              _nw437[_i0_86] = _init86(new BigNumber(_i0_86));
            }
            _2537_v37 = _nw437;
            let _2539_v38;
            let _nw438 = new _module.C6();
            _nw438.__ctor(_this.f7, p0);
            _2539_v38 = _nw438;
            let _2540_v39;
            _2540_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2539_v38,p2);
            let _index407 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
            (_2537_v37)[_index407] = !(_2540_v39).contains(_2539_v38);
            let _2541_v40;
            let _nw439 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2541_v40 = _nw439;
            let _index408 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length));
            (_2541_v40)[_index408] = _2523_v32;
            let _index409 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
            let _index410 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length));
            let _rhs384 = _dafny.Seq.IsProperPrefixOf(_2523_v32, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("duoejddgs"), _2523_v32));
            let _rhs385 = _dafny.Seq.Concat(_2523_v32, _dafny.Seq.Concat(_2523_v32, _2523_v32));
            let _lhs234 = _2537_v37;
            let _lhs235 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
            let _lhs236 = _2541_v40;
            let _lhs237 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length));
            _lhs234[_lhs235] = _rhs384;
            _lhs236[_lhs237] = _rhs385;
            (_2539_v38).f7 = p2;
            if (_2539_v38.f7) {
              let _2542_v41;
              _2542_v41 = new BigNumber(479);
              _2542_v41 = (_dafny.ZERO).minus((p1).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(45), _2542_v41)));
              let _2543_v42;
              let _nw440 = Array((new BigNumber(6)).toNumber());
              _nw440[(_dafny.ZERO).toNumber()] = _2542_v41;
              _nw440[(_dafny.ONE).toNumber()] = new BigNumber(149);
              _nw440[(new BigNumber(2)).toNumber()] = p1;
              _nw440[(new BigNumber(3)).toNumber()] = _2542_v41;
              _nw440[(new BigNumber(4)).toNumber()] = p1;
              _nw440[(new BigNumber(5)).toNumber()] = new BigNumber((_2523_v32).length);
              _2543_v42 = _nw440;
              let _2544_v43;
              _2544_v43 = _module.D14.create_DC39(_2523_v32, _2543_v42);
              let _2545_v44;
              _2545_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2526_v35,(_2541_v40)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length))]);
              let _2546_v45;
              _2546_v45 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_2526_v35), (((_2545_v44).contains(_2526_v35)) ? ((_2545_v44).get(_2526_v35)) : (_dafny.Seq.update(_dafny.Seq.of(_2526_v35), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_2523_v32, _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2523_v32).length)), new _dafny.CodePoint('v'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.of(_2526_v35)).length)), _2526_v35)))));
              let _index411 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
              let _rhs386 = _dafny.Seq.IsProperPrefixOf((_2544_v43).dtor_cf63, ((_2539_v38.f7) ? (_2523_v32) : (_2523_v32)));
              let _rhs387 = (_2546_v45)[_module.__default.safeIndex(_2542_v41, new BigNumber((_2546_v45).length))];
              let _lhs238 = _2539_v38;
              let _lhs239 = _2537_v37;
              let _lhs240 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
              _lhs238.f7 = _rhs386;
              _lhs239[_lhs240] = _rhs387;
              _2542_v41 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(p1)).minus(_2542_v41));
              let _2547_v46;
              _2547_v46 = _dafny.Set.fromElements(_2539_v38.f7, true);
              let _2548_v47;
              _2548_v47 = _dafny.Seq.of(_2542_v41, new BigNumber(138));
              let _2549_v48;
              _2549_v48 = _module.D1.create_DC5(_2539_v38.f7, _2548_v47, new BigNumber(616));
              r1 = (_2547_v46).Difference(((false) ? (_module.__default.fm27(_2539_v38.f7, _2526_v35, (_2537_v37)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length))], !((_2549_v48).dtor_cf9), globalState)) : (_2547_v46)));
              _2542_v41 = _module.__default.safeDivisionInt(p1, p1);
            } else {
              let _2550_v49;
              _2550_v49 = new BigNumber(329);
              _2550_v49 = p1;
              let _2551_v50;
              let _nw441 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
              _2551_v50 = _nw441;
              let _index412 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2551_v50).length));
              (_2551_v50)[_index412] = new BigNumber(35);
              let _2552_v51;
              _2552_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2550_v49);
              let _index413 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2551_v50).length));
              let _rhs388 = !(p0);
              let _rhs389 = (((_2552_v51).contains(_2539_v38.f7)) ? ((_2552_v51).get(_2539_v38.f7)) : (_2550_v49));
              let _rhs390 = p1;
              let _rhs391 = (_2539_v38).fm9(!((_2537_v37)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length))]), p1, _2526_v35, globalState);
              let _rhs392 = _2550_v49;
              let _lhs241 = _this;
              let _lhs242 = _2551_v50;
              let _lhs243 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2551_v50).length));
              _lhs241.f7 = _rhs388;
              _2550_v49 = _rhs389;
              _lhs242[_lhs243] = _rhs390;
              _2550_v49 = _rhs391;
              _2550_v49 = _rhs392;
              (_2539_v38).f7 = (_2537_v37)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length))];
              let _pat_let_tv56 = _2550_v49;
              let _pat_let_tv57 = _2550_v49;
              let _pat_let_tv58 = p1;
              let _index414 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
              (_2537_v37)[_index414] = (function (_pat_let73_0) {
                return function (_2553_dt__update__tmp_h1) {
                  return function (_pat_let74_0) {
                    return function (_2554_dt__update_hcf22_h0) {
                      return function (_pat_let75_0) {
                        return function (_2555_dt__update_hcf23_h0) {
                          return function (_pat_let76_0) {
                            return function (_2556_dt__update_hcf24_h0) {
                              return _module.D3.create_DC11(_2554_dt__update_hcf22_h0, _2555_dt__update_hcf23_h0, _2556_dt__update_hcf24_h0, (_2553_dt__update__tmp_h1).dtor_cf25);
                            }(_pat_let76_0);
                          }(_pat_let_tv58);
                        }(_pat_let75_0);
                      }(_pat_let_tv57);
                    }(_pat_let74_0);
                  }(_pat_let_tv56);
                }(_pat_let73_0);
              }(_module.D3.create_DC11(_module.__default.fm0(globalState), (_2551_v50)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2551_v50).length))], p1, _2539_v38.f7))).dtor_cf25;
              let _index415 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2551_v50).length));
              (_2551_v50)[_index415] = p1;
            }
            if (p0) {
              let _2557_v52;
              _2557_v52 = new BigNumber(150);
              _2557_v52 = _2557_v52;
              let _2558_v53;
              _2558_v53 = _dafny.Seq.of(_this.f7);
              let _2559_v54;
              _2559_v54 = _dafny.Set.fromElements(new BigNumber((_2558_v53).length));
              let _2560_v55;
              _2560_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2524_v33,new BigNumber((_2559_v54).length));
              let _2561_v56;
              _2561_v56 = _module.D2.create_DC9(_2557_v52, _2557_v52, (_2541_v40)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length))], p0, _2560_v55);
              let _2562_v57;
              _2562_v57 = _dafny.Set.fromElements(new BigNumber(((_2561_v56).dtor_cf18).length));
              let _2563_v58;
              _2563_v58 = _dafny.MultiSet.fromElements(p1, (_2557_v52).minus(new BigNumber((_2559_v54).length)));
              _2563_v58 = _2563_v58;
              let _2564_v59;
              _2564_v59 = _module.D1.create_DC3(_dafny.Set.fromElements(_2539_v38.f7, p0, _2539_v38.f7));
              let _2565_v60;
              _2565_v60 = _dafny.Seq.of(_2557_v52, p1, new BigNumber(872), new BigNumber(((_2541_v40)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_2541_v40).length))]).length), new BigNumber((_2558_v53).length));
              let _2566_v61;
              _2566_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2564_v59,new BigNumber((_2565_v60).length));
              _2566_v61 = _2566_v61;
              let _2567_v62;
              _2567_v62 = _module.D2.create_DC7(_2559_v54);
              let _2568_v63;
              _2568_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2526_v35,_2567_v62);
              let _2569_v64;
              _2569_v64 = _dafny.Set.fromElements(_2537_v37, _2537_v37);
              let _2570_v65;
              let _nw442 = new _module.C2();
              _nw442.__ctor(_2569_v64, _2526_v35, _this.f7);
              _2570_v65 = _nw442;
              let _2571_v66;
              _2571_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2557_v52,_2570_v65);
              let _2572_v67;
              _2572_v67 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2571_v66);
              let _2573_v68;
              _2573_v68 = _module.D3.create_DC10(_2537_v37);
              let _2574_v69;
              let _nw443 = Array((new BigNumber(4)).toNumber());
              _nw443[(_dafny.ZERO).toNumber()] = (_2573_v68).dtor_cf21;
              _nw443[(_dafny.ONE).toNumber()] = _2537_v37;
              _nw443[(new BigNumber(2)).toNumber()] = _2537_v37;
              _nw443[(new BigNumber(3)).toNumber()] = _2537_v37;
              _2574_v69 = _nw443;
              let _index416 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2574_v69).length));
              (_2574_v69)[_index416] = _2537_v37;
              let _index417 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2574_v69).length));
              let _rhs393 = ((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), ((_2575_p1) => function (_2576_i9) {
                return _2575_p1;
              })(p1)), _2565_v60)) ? (!(_2562_v57).equals(_2562_v57)) : (_module.__default.fm1(_dafny.Seq.update(_2525_v34, _module.__default.safeIndex(_2557_v52, new BigNumber((_2525_v34).length)), _2524_v33), (_dafny.ZERO).minus(_2557_v52), new BigNumber(805), globalState)));
              let _rhs394 = (_2568_v63).Merge(_2568_v63);
              let _rhs395 = _2572_v67;
              let _rhs396 = _2537_v37;
              let _lhs244 = _2539_v38;
              let _lhs245 = _2574_v69;
              let _lhs246 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2574_v69).length));
              _lhs244.f7 = _rhs393;
              _2568_v63 = _rhs394;
              _2572_v67 = _rhs395;
              _lhs245[_lhs246] = _rhs396;
              let _index418 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2537_v37).length));
              (_2537_v37)[_index418] = ((false) ? (_this.f7) : (_2539_v38.f7));
            } else {
              let _2577_v70;
              _2577_v70 = new BigNumber(126);
              let _2578_v71;
              _2578_v71 = _dafny.Map.Empty.slice().updateUnsafe(true,_2537_v37);
              _2577_v70 = new BigNumber((_2578_v71).length);
              let _2579_v72;
              let _nw444 = new _module.C5();
              _nw444.__ctor(p2);
              _2579_v72 = _nw444;
              let _2580_v73;
              _2580_v73 = _dafny.Seq.of(_2537_v37);
              let _2581_v74;
              _2581_v74 = _dafny.Set.fromElements(_2537_v37, _2537_v37, (_2580_v73)[_module.__default.safeIndex(_2577_v70, new BigNumber((_2580_v73).length))]);
              let _2582_v75;
              let _nw445 = new _module.C2();
              _nw445.__ctor(_2581_v74, _2526_v35, (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(534)), ((_2583_v35, _2584_p0, _2585_v38) => function (_2586_i10) {
                return (_module.D13.create_DC36(_2583_v35, _dafny.Seq.of(_2584_p0, _2585_v38.f7), true)).dtor_cf58;
              })(_2526_v35, p0, _2539_v38))).length)).isEqualTo(_2577_v70));
              _2582_v75 = _nw445;
              _2537_v37 = _2537_v37;
              (_2582_v75).m3(globalState);
            }
          }
        }
      }
      let _hi16 = p1;
      for (let _2587_i11 = p1; _2587_i11.isLessThan(_hi16); _2587_i11 = _2587_i11.plus(_dafny.ONE)) {
        let _2588_v76;
        _2588_v76 = _dafny.MultiSet.fromElements(new BigNumber(621));
        let _2589_v77;
        _2589_v77 = _module.D12.create_DC29((_2588_v76).Union(_2588_v76));
        let _source18 = _2589_v77;
        if (_source18.is_DC30) {
          let _2590___mcc_h0 = (_source18).cf49;
          let _2591___mcc_h1 = (_source18).cf50;
          let _2592___mcc_h2 = (_source18).cf51;
          let _2593___mcc_h3 = (_source18).cf52;
          let _2594_cf52 = _2593___mcc_h3;
          let _2595_cf51 = _2592___mcc_h2;
          let _2596_cf50 = _2591___mcc_h1;
          let _2597_cf49 = _2590___mcc_h0;
          _2523_v32 = _2523_v32;
          let _2598_v78;
          let _nw446 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2598_v78 = _nw446;
          let _2599_v79;
          _2599_v79 = _dafny.Seq.of(_2598_v78, _2598_v78);
          _2598_v78 = (_2599_v79)[_module.__default.safeIndex(_2587_i11, new BigNumber((_2599_v79).length))];
          _2597_cf49 = _module.__default.fm0(globalState);
          _2523_v32 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2523_v32, _dafny.Seq.update(_2523_v32, _module.__default.safeIndex(_2594_cf52, new BigNumber((_2523_v32).length)), _2526_v35)), _2523_v32);
        } else if (_source18.is_DC31) {
          let _2600___mcc_h4 = (_source18).cf53;
          let _2601___mcc_h5 = (_source18).cf54;
          let _2602_cf54 = _2601___mcc_h5;
          let _2603_cf53 = _2600___mcc_h4;
          let _2604_v80;
          _2604_v80 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
          let _2605_v81;
          let _nw447 = new _module.C3();
          _nw447.__ctor((new BigNumber(((_2604_v80).update(p0, _module.__default.abs(_2587_i11))).cardinality())).minus(p1), p2);
          _2605_v81 = _nw447;
          let _2606_v82;
          _2606_v82 = _dafny.Seq.of(_2587_i11);
          let _2607_v83;
          _2607_v83 = _dafny.Seq.of((_this).fm8((_2606_v82)[_module.__default.safeIndex((_2605_v81).f14, new BigNumber((_2606_v82).length))], globalState), p2, _this.f7);
          (_this).f7 = (_2607_v83)[_module.__default.safeIndex((new BigNumber((_2588_v76).cardinality())).plus((_2605_v81).f14), new BigNumber((_2607_v83).length))];
          let _2608_v84;
          _2608_v84 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_2523_v32).length));
          let _2609_v85;
          _2609_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(728),(_2605_v81).f14);
          let _2610_v86;
          let _nw448 = Array((new BigNumber(15)).toNumber());
          _nw448[(_dafny.ZERO).toNumber()] = _2587_i11;
          _nw448[(_dafny.ONE).toNumber()] = p1;
          _nw448[(new BigNumber(2)).toNumber()] = (_2605_v81).f14;
          _nw448[(new BigNumber(3)).toNumber()] = _2587_i11;
          _nw448[(new BigNumber(4)).toNumber()] = _2587_i11;
          _nw448[(new BigNumber(5)).toNumber()] = new BigNumber((_2608_v84).length);
          _nw448[(new BigNumber(6)).toNumber()] = _2587_i11;
          _nw448[(new BigNumber(7)).toNumber()] = (_2605_v81).fm9(p0, p1, new _dafny.CodePoint('b'.codePointAt(0)), globalState);
          _nw448[(new BigNumber(8)).toNumber()] = _2587_i11;
          _nw448[(new BigNumber(9)).toNumber()] = _2587_i11;
          _nw448[(new BigNumber(10)).toNumber()] = p1;
          _nw448[(new BigNumber(11)).toNumber()] = (_2605_v81).f14;
          _nw448[(new BigNumber(12)).toNumber()] = new BigNumber(((_2609_v85).update(p1, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length))).length);
          _nw448[(new BigNumber(13)).toNumber()] = new BigNumber(794);
          _nw448[(new BigNumber(14)).toNumber()] = _2587_i11;
          _2610_v86 = _nw448;
          let _2611_v87;
          _2611_v87 = _dafny.MultiSet.fromElements(_2610_v86, _2610_v86, _2610_v86, _2610_v86, _2610_v86);
          _2611_v87 = _2611_v87;
          _2602_cf54 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), ((_2612_v35) => function (_2613_i12) {
            return _2612_v35;
          })(_2526_v35)), _module.__default.fm20(p2, _2602_cf54, globalState)), _2523_v32));
        } else if (_source18.is_DC32) {
          let _2614___mcc_h6 = (_source18).cf55;
          let _2615_cf55 = _2614___mcc_h6;
          let _2616_v88;
          _2616_v88 = _dafny.Seq.of(_this.f7, (_this).fm8(_2587_i11, globalState), p2);
          let _2617_v89;
          _2617_v89 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_2616_v88, _module.__default.safeIndex(p1, new BigNumber((_2616_v88).length)), _this.f7)).length), _2587_i11, _module.__default.fm0(globalState), _2587_i11, p1);
          let _2618_v90;
          _2618_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2617_v89);
          _2618_v90 = _2618_v90;
          let _rhs397 = _2615_cf55;
          let _rhs398 = (((p1).isLessThanOrEqualTo((_dafny.ZERO).minus(p1))) ? (p0) : (!(_2615_cf55)));
          _2615_cf55 = _rhs397;
          _2615_cf55 = _rhs398;
          _2526_v35 = _2526_v35;
          let _2619_v91;
          _2619_v91 = new BigNumber(-542);
          let _2620_v92;
          _2620_v92 = _dafny.Seq.of((_dafny.ZERO).minus(p1));
          _2619_v91 = _module.__default.safeDivisionInt(p1, (_2620_v92)[_module.__default.safeIndex(p1, new BigNumber((_2620_v92).length))]);
        } else if (_source18.is_DC29) {
          let _2621___mcc_h7 = (_source18).cf48;
          let _2622_cf48 = _2621___mcc_h7;
          (_this).f7 = false;
          (_this).f7 = _this.f7;
          let _2623_v93;
          _2623_v93 = _dafny.Set.fromElements(_2587_i11);
          r1 = _module.__default.fm27((_dafny.Set.fromElements(p1, _2587_i11, p1)).IsProperSubsetOf(_2623_v93), _2526_v35, _this.f7, !(p2) || (p2), globalState);
          let _2624_v94;
          let _init87 = ((_2625_i11) => function (_2626_i13) {
            return (_2626_i13).minus(_2625_i11);
          })(_2587_i11);
          let _nw449 = Array((_dafny.ONE).toNumber());
          for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw449.length); _i0_87++) {
            _nw449[_i0_87] = _init87(new BigNumber(_i0_87));
          }
          _2624_v94 = _nw449;
          let _index419 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2624_v94).length));
          (_2624_v94)[_index419] = _2587_i11;
          let _index420 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2624_v94).length));
          (_2624_v94)[_index420] = _2587_i11;
        } else {
          let _2627___mcc_h8 = (_source18).cf56;
          let _2628_cf56 = _2627___mcc_h8;
          let _2629_v95;
          _2629_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2587_i11,p2);
          let _2630_v96;
          _2630_v96 = _dafny.Seq.of(new BigNumber((_module.__default.fm27(p2, _2526_v35, p0, (((_2629_v95).contains(_2587_i11)) ? ((_2629_v95).get(_2587_i11)) : (_this.f7)), globalState)).length));
          let _rhs399 = (_2587_i11).isEqualTo((_2587_i11).minus(new BigNumber((_2630_v96).length)));
          let _rhs400 = (p0) && (p2);
          let _lhs247 = _this;
          let _lhs248 = _this;
          _lhs247.f7 = _rhs399;
          _lhs248.f7 = _rhs400;
          let _2631_v97;
          _2631_v97 = new BigNumber(971);
          let _2632_v98;
          _2632_v98 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).fm8(new BigNumber((_dafny.Seq.update(_2523_v32, _module.__default.safeIndex(new BigNumber(379), new BigNumber((_2523_v32).length)), _2526_v35)).length), globalState));
          _2631_v97 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2631_v97).plus(new BigNumber((((p0) ? (_2632_v98) : (_2632_v98))).length))));
          let _2633_v99;
          _2633_v99 = _module.D6.create_DC15(_2526_v35);
          let _2634_v100;
          _2634_v100 = _module.D9.create_DC25();
          let _2635_v101;
          _2635_v101 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_2523_v32);
          let _rhs401 = _2523_v32;
          let _rhs402 = _module.__default.fm37(false, _2634_v100, globalState);
          let _rhs403 = p1;
          let _rhs404 = _2633_v99;
          let _rhs405 = (((_2635_v101).contains((_2631_v97).isEqualTo(_2631_v97))) ? ((_2635_v101).get((_2631_v97).isEqualTo(_2631_v97))) : (_dafny.Seq.Concat(_2523_v32, _dafny.Seq.UnicodeFromString("vrwpnkxg"))));
          _2523_v32 = _rhs401;
          _2526_v35 = _rhs402;
          _2631_v97 = _rhs403;
          _2633_v99 = _rhs404;
          _2523_v32 = _rhs405;
          let _2636_v102;
          _2636_v102 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-34)), ((_2637_v35) => function (_2638_i14) {
            return _2637_v35;
          })(_2526_v35)), _2523_v32);
          let _rhs406 = _2636_v102;
          let _rhs407 = new BigNumber((_2523_v32).length);
          _2636_v102 = _rhs406;
          _2631_v97 = _rhs407;
        }
        if (true) {
          let _2639_v103;
          let _nw450 = Array((new BigNumber(27)).toNumber()).fill(false);
          _2639_v103 = _nw450;
          let _2640_v104;
          _2640_v104 = _dafny.Set.fromElements(_2639_v103, _2639_v103);
          let _2641_v105;
          let _nw451 = new _module.C2();
          _nw451.__ctor(_2640_v104, _2526_v35, !(p2));
          _2641_v105 = _nw451;
          _2641_v105 = _2641_v105;
          let _2642_v106;
          let _init88 = ((_2643_p2) => function (_2644_i15) {
            return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(493),_2643_p2));
          })(p2);
          let _nw452 = Array((_dafny.ONE).toNumber());
          for (let _i0_88 = 0; _i0_88 < new BigNumber(_nw452.length); _i0_88++) {
            _nw452[_i0_88] = _init88(new BigNumber(_i0_88));
          }
          _2642_v106 = _nw452;
          let _2645_v107;
          _2645_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),p2);
          let _2646_v108;
          _2646_v108 = _dafny.Seq.of(_2645_v107, _2645_v107, _2645_v107, _2645_v107, _2645_v107);
          let _index421 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_2642_v106).length));
          (_2642_v106)[_index421] = _dafny.Seq.Concat(_dafny.Seq.update(_2646_v108, _module.__default.safeIndex(_2587_i11, new BigNumber((_2646_v108).length)), _2645_v107), _dafny.Seq.update(_2646_v108, _module.__default.safeIndex(_2587_i11, new BigNumber((_2646_v108).length)), _2645_v107));
          let _index422 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_2642_v106).length));
          (_2642_v106)[_index422] = _2646_v108;
          let _2647_v109;
          let _nw453 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _2647_v109 = _nw453;
          _2647_v109 = _2647_v109;
          let _2648_v110;
          let _nw454 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
          _2648_v110 = _nw454;
          let _index423 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2648_v110).length));
          (_2648_v110)[_index423] = _2527_v36;
          let _2649_v111;
          let _nw455 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2649_v111 = _nw455;
          let _2650_v112;
          let _init89 = ((_2651_v35) => function (_2652_i16) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), ((_2653_v35) => function (_2654_i17) {
              return _2653_v35;
            })(_2651_v35));
          })(_2526_v35);
          let _nw456 = Array((new BigNumber(6)).toNumber());
          for (let _i0_89 = 0; _i0_89 < new BigNumber(_nw456.length); _i0_89++) {
            _nw456[_i0_89] = _init89(new BigNumber(_i0_89));
          }
          _2650_v112 = _nw456;
          let _index424 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2648_v110).length));
          let _rhs408 = _2527_v36;
          let _rhs409 = _2639_v103;
          let _rhs410 = _2649_v111;
          let _rhs411 = (_2641_v105).f16;
          let _rhs412 = _2650_v112;
          let _lhs249 = _2648_v110;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2648_v110).length));
          _lhs249[_lhs250] = _rhs408;
          _2639_v103 = _rhs409;
          _2649_v111 = _rhs410;
          _2526_v35 = _rhs411;
          _2650_v112 = _rhs412;
          let _2655_v113;
          _2655_v113 = _dafny.Seq.of(_this.f7);
          let _2656_v114;
          let _nw457 = new _module.C6();
          _nw457.__ctor(!(p2), !(_dafny.areEqual(_2655_v113, _2655_v113)));
          _2656_v114 = _nw457;
        } else {
          let _2657_v115;
          _2657_v115 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f7);
          let _2658_v116;
          _2658_v116 = _2657_v115;
          r0 = (_2658_v116);
          let _2659_v117;
          _2659_v117 = _dafny.Seq.of(p0, p2);
          (_this).f7 = (_2659_v117)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_module.__default.fm0(globalState), _2587_i11), new BigNumber((_2659_v117).length))];
          let _2660_v118;
          let _nw458 = new _module.C6();
          _nw458.__ctor((_2587_i11).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(p0, p0, p0, _this.f7, p2)).cardinality())), p0);
          _2660_v118 = _nw458;
          _2660_v118 = _2660_v118;
          (_2660_v118).f7 = _this.f7;
          (_this).f7 = (_this).fm8(p1, globalState);
        }
        let _2661_v119;
        _2661_v119 = new BigNumber(322);
        _2661_v119 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(_2587_i11, _2587_i11))).multipliedBy(_2587_i11);
        (_this).f7 = _this.f7;
      }
      let _2662_v120;
      _2662_v120 = _module.D6.create_DC17(p0, p1, _2526_v35, _2526_v35);
      let _2663_v121;
      _2663_v121 = _module.D6.create_DC18(_2662_v120);
      let _pat_let_tv59 = p2;
      let _pat_let_tv60 = p2;
      let _pat_let_tv61 = p0;
      let _pat_let_tv62 = _2523_v32;
      let _pat_let_tv63 = p2;
      let _pat_let_tv64 = p2;
      r0 = function (_source19) {
        if (_source19.is_DC16) {
          let _2664___mcc_h9 = (_source19).cf30;
          let _2665_cf30 = _2664___mcc_h9;
          return (_dafny.Map.Empty.slice().updateUnsafe(!(_2665_cf30),_pat_let_tv59)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_this.f7));
        } else if (_source19.is_DC17) {
          let _2666___mcc_h10 = (_source19).cf31;
          let _2667___mcc_h11 = (_source19).cf32;
          let _2668___mcc_h12 = (_source19).cf33;
          let _2669___mcc_h13 = (_source19).cf34;
          let _2670_cf34 = _2669___mcc_h13;
          let _2671_cf33 = _2668___mcc_h12;
          let _2672_cf32 = _2667___mcc_h11;
          let _2673_cf31 = _2666___mcc_h10;
          return (_dafny.Map.Empty.slice().updateUnsafe(_this.f7,false)).update(_pat_let_tv60, !(_2673_cf31));
        } else if (_source19.is_DC15) {
          let _2674___mcc_h14 = (_source19).cf29;
          let _2675_cf29 = _2674___mcc_h14;
          return (_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC8(_pat_let_tv61, _pat_let_tv62)).dtor_cf14,false)).update(!(!(true)), _this.f7);
        } else {
          let _2676___mcc_h15 = (_source19).cf35;
          let _2677_cf35 = _2676___mcc_h15;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv63,_pat_let_tv64);
        }
      }(_2663_v121);
      let _2678_v122;
      _2678_v122 = _dafny.Set.fromElements(p0);
      r1 = (_2678_v122).Difference(_2678_v122);
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _2679_v0;
      _2679_v0 = new BigNumber(-924);
      let _2680_v1;
      _2680_v1 = _dafny.Seq.of(_this.f7, true);
      let _2681_v2;
      _2681_v2 = new _dafny.CodePoint('q'.codePointAt(0));
      let _2682_v3;
      _2682_v3 = _dafny.Set.fromElements(_2681_v2, _2681_v2);
      let _2683_v5;
      _2683_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of (_2682_v3).Elements) {
          let _2684_v4 = _compr_53;
          if ((_2682_v3).contains(_2684_v4)) {
            _coll53.add(_2684_v4);
          }
        }
        return _coll53;
      }()).length),_this.f7);
      let _2685_v6;
      _2685_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2683_v5).length),_2679_v0);
      let _hi17 = ((_dafny.ZERO).minus(new BigNumber((_2680_v1).length))).multipliedBy(new BigNumber((_2685_v6).length));
      for (let _2686_i0 = _2679_v0; _2686_i0.isLessThan(_hi17); _2686_i0 = _2686_i0.plus(_dafny.ONE)) {
        _2679_v0 = _2679_v0;
        let _2687_v7;
        _2687_v7 = _dafny.MultiSet.fromElements(_this.f7);
        let _2688_v8;
        _2688_v8 = _dafny.Seq.of(_2687_v7);
        let _2689_v9;
        _2689_v9 = _dafny.MultiSet.fromElements((_2688_v8)[_module.__default.safeIndex(new BigNumber((_2680_v1).length), new BigNumber((_2688_v8).length))]);
        _2689_v9 = (_2689_v9).Union(_2689_v9);
        (_this).f7 = _this.f7;
        let _2690_v10;
        _2690_v10 = _dafny.Seq.UnicodeFromString("ygaj");
        if (!_dafny.Seq.contains(_2690_v10, _2681_v2)) {
          let _2691_v11;
          _2691_v11 = _dafny.MultiSet.fromElements(_2686_i0);
          _2691_v11 = _2691_v11;
          let _2692_v12;
          let _nw459 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _2692_v12 = _nw459;
          let _index425 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_2692_v12).length));
          (_2692_v12)[_index425] = (_dafny.Map.Empty.slice().updateUnsafe(_2686_i0,new BigNumber(-272))).Merge(_2685_v6);
          let _2693_v13;
          _2693_v13 = _dafny.Seq.of(_2685_v6);
          let _index426 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_2692_v12).length));
          (_2692_v12)[_index426] = (_2693_v13)[_module.__default.safeIndex(_2686_i0, new BigNumber((_2693_v13).length))];
          (_this).f7 = !(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_2694_v2) => function (_2695_i1) {
            return _2694_v2;
          })(_2681_v2)), _module.__default.safeIndex(_2686_i0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_2696_v2) => function (_2697_i1) {
            return _2696_v2;
          })(_2681_v2))).length)), _2681_v2)).length), _2679_v0)).isEqualTo(_2686_i0);
          let _2698_v14;
          _2698_v14 = _dafny.Set.fromElements(_2686_i0, (_dafny.ZERO).minus(_2686_i0), _2679_v0);
          _2679_v0 = (_dafny.ZERO).minus(new BigNumber((_2698_v14).length));
          let _2699_v15;
          _2699_v15 = _dafny.Seq.of(_2679_v0);
          _2679_v0 = (_2699_v15)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_2679_v0)), new BigNumber((_2699_v15).length))];
        } else {
          (_this).f7 = _this.f7;
          (_this).f7 = _this.f7;
          let _2700_v16;
          _2700_v16 = _2682_v3;
          let _2701_v17;
          _2701_v17 = _dafny.Seq.of(_2700_v16);
          let _2702_v18;
          _2702_v18 = _dafny.Map.Empty.slice().updateUnsafe(((((_this).f8).contains(_this.f7)) ? (((_this).f8).get(_this.f7)) : (_dafny.Seq.of(_2686_i0))),_dafny.Seq.update(_2701_v17, _module.__default.safeIndex(new BigNumber((_2690_v10).length), new BigNumber((_2701_v17).length)), _2700_v16));
          _2702_v18 = _2702_v18;
          let _2703_v19;
          _2703_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2681_v2);
          (_this).f7 = (_module.__default.safeDivisionInt(new BigNumber((_2703_v19).length), _2679_v0)).isLessThan((((_2685_v6).contains(new BigNumber((_2690_v10).length))) ? ((_2685_v6).get(new BigNumber((_2690_v10).length))) : (_2679_v0)));
          let _2704_v20;
          _2704_v20 = _dafny.Set.fromElements(_this.f7);
          let _2705_v21;
          _2705_v21 = _module.D1.create_DC3(_2704_v20);
          let _2706_v22;
          _2706_v22 = _module.D0.create_DC0(_2679_v0);
          let _2707_v23;
          _2707_v23 = _module.D8.create_DC22(new BigNumber(939), _2705_v21, _2706_v22);
          let _rhs413 = _2679_v0;
          let _rhs414 = _module.__default.fm55(_2681_v2, globalState);
          let _rhs415 = _2690_v10;
          let _rhs416 = _2683_v5;
          _2679_v0 = _rhs413;
          _2707_v23 = _rhs414;
          _2690_v10 = _rhs415;
          _2683_v5 = _rhs416;
        }
      }
      let _2708_v24;
      _2708_v24 = _dafny.Set.fromElements(_2679_v0, new BigNumber((_dafny.Seq.of(_this.f7)).length), (_this).fm9(_this.f7, new BigNumber(958), _2681_v2, globalState));
      if ((_2708_v24).IsProperSubsetOf(_2708_v24)) {
        let _2709_v25;
        let _nw460 = Array((new BigNumber(26)).toNumber()).fill(false);
        _2709_v25 = _nw460;
        let _index427 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length));
        (_2709_v25)[_index427] = _this.f7;
        let _index428 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length));
        (_2709_v25)[_index428] = _this.f7;
        let _2710_v26;
        _2710_v26 = _dafny.Seq.of(new BigNumber(-672));
        let _2711_v27;
        _2711_v27 = _dafny.Seq.of(_2710_v26, _2710_v26, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), ((_2712_v0) => function (_2713_i3) {
          return _2712_v0;
        })(_2679_v0)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("owdl")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), ((_2714_v0) => function (_2715_i3) {
          return _2714_v0;
        })(_2679_v0))).length)), new BigNumber((_module.__default.fm20(_this.f7, _this.f7, globalState)).length)));
        if (!_dafny.Seq.contains(_2711_v27, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), ((_2716_v0) => function (_2717_i2) {
          return _2716_v0;
        })(_2679_v0)), _dafny.Seq.of(_2679_v0)))) {
          let _2718_v28;
          let _nw461 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _2718_v28 = _nw461;
          let _2719_v29;
          _2719_v29 = _dafny.MultiSet.fromElements(true);
          let _2720_v30;
          _2720_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2679_v0,_2719_v29);
          let _index429 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2718_v28).length));
          (_2718_v28)[_index429] = _2720_v30;
          let _index430 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2718_v28).length));
          (_2718_v28)[_index430] = _2720_v30;
          _2679_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_2679_v0));
          _2679_v0 = _2679_v0;
          _2679_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2680_v1).length)));
          let _2721_v31;
          _2721_v31 = _dafny.Seq.UnicodeFromString("cljigeqyq");
          _2721_v31 = _2721_v31;
        } else {
          let _2722_v32;
          let _nw462 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2722_v32 = _nw462;
          let _2723_v33;
          _2723_v33 = _module.D1.create_DC3(_dafny.Set.fromElements(_this.f7, _this.f7));
          let _2724_v34;
          _2724_v34 = _dafny.Map.Empty.slice().updateUnsafe((_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))],(_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))]);
          let _2725_v35;
          let _nw463 = new _module.C3();
          _nw463.__ctor((_2710_v26)[_module.__default.safeIndex(new BigNumber((_2724_v34).length), new BigNumber((_2710_v26).length))], (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))]);
          _2725_v35 = _nw463;
          let _2726_v36;
          _2726_v36 = _dafny.Map.Empty.slice().updateUnsafe((_2723_v33).dtor_cf6,_2725_v35);
          let _2727_v37;
          let _nw464 = Array((new BigNumber(10)).toNumber());
          _nw464[(_dafny.ZERO).toNumber()] = _2679_v0;
          _nw464[(_dafny.ONE).toNumber()] = new BigNumber(339);
          _nw464[(new BigNumber(2)).toNumber()] = new BigNumber(-621);
          _nw464[(new BigNumber(3)).toNumber()] = _2679_v0;
          _nw464[(new BigNumber(4)).toNumber()] = new BigNumber((_2680_v1).length);
          _nw464[(new BigNumber(5)).toNumber()] = _2679_v0;
          _nw464[(new BigNumber(6)).toNumber()] = new BigNumber(585);
          _nw464[(new BigNumber(7)).toNumber()] = new BigNumber((_2726_v36).length);
          _nw464[(new BigNumber(8)).toNumber()] = (_2725_v35).f14;
          _nw464[(new BigNumber(9)).toNumber()] = _2679_v0;
          _2727_v37 = _nw464;
          let _2728_v38;
          _2728_v38 = _dafny.Seq.of(_2727_v37, _2727_v37);
          let _2729_v39;
          _2729_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2709_v25,_2727_v37);
          let _2730_v40;
          let _nw465 = Array((new BigNumber(15)).toNumber());
          _nw465[(_dafny.ZERO).toNumber()] = _2727_v37;
          _nw465[(_dafny.ONE).toNumber()] = (_2728_v38)[_module.__default.safeIndex(_2679_v0, new BigNumber((_2728_v38).length))];
          _nw465[(new BigNumber(2)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(3)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(4)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(5)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(6)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(7)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(8)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(9)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(10)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(11)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(12)).toNumber()] = (((_2729_v39).contains(_2709_v25)) ? ((_2729_v39).get(_2709_v25)) : (_2727_v37));
          _nw465[(new BigNumber(13)).toNumber()] = _2727_v37;
          _nw465[(new BigNumber(14)).toNumber()] = _2727_v37;
          _2730_v40 = _nw465;
          let _index431 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_2722_v32).length));
          (_2722_v32)[_index431] = _2730_v40;
          let _index432 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_2722_v32).length));
          (_2722_v32)[_index432] = ((_this.f7) ? (_2730_v40) : (_2730_v40));
          let _2731_v41;
          _2731_v41 = _dafny.Set.fromElements(_2727_v37);
          let _index433 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length));
          let _rhs417 = _this.f7;
          let _rhs418 = (new BigNumber(-325)).plus(_2679_v0);
          let _rhs419 = ((_2731_v41).Difference(_2731_v41)).Difference(_dafny.Set.fromElements(_2727_v37));
          let _rhs420 = (_dafny.ZERO).minus((_2679_v0).minus((_2725_v35).f14));
          let _lhs251 = _2709_v25;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length));
          _lhs251[_lhs252] = _rhs417;
          _2679_v0 = _rhs418;
          _2731_v41 = _rhs419;
          _2679_v0 = _rhs420;
          let _2732_v42;
          _2732_v42 = _dafny.Seq.UnicodeFromString("qpswxaqat");
          let _rhs421 = new BigNumber((_2732_v42).length);
          let _rhs422 = (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))];
          let _rhs423 = (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))];
          let _lhs253 = _this;
          let _lhs254 = _this;
          _2679_v0 = _rhs421;
          _lhs253.f7 = _rhs422;
          _lhs254.f7 = _rhs423;
          _2679_v0 = (_module.__default.safeDivisionInt((_2725_v35).f14, new BigNumber((_dafny.MultiSet.fromElements((_2725_v35).f14, _2679_v0)).cardinality()))).multipliedBy(((_2725_v35).f14).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f7)).length)));
          let _2733_v43;
          _2733_v43 = _module.D0.create_DC0((_2725_v35).f14);
          let _2734_v44;
          let _nw466 = new _module.C6();
          _nw466.__ctor(_module.__default.fm1(_dafny.Seq.of(_2733_v43, _module.__default.fm56(globalState)), (_2725_v35).f14, (_dafny.ZERO).minus(_2679_v0), globalState), !_dafny.Seq.contains(_dafny.Seq.of(_2709_v25, _2709_v25), _2709_v25));
          _2734_v44 = _nw466;
        }
        (_this).f7 = (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))];
        let _2735_v45;
        _2735_v45 = _dafny.Set.fromElements(_2709_v25, _2709_v25, _2709_v25);
        let _2736_v46;
        _2736_v46 = _module.D8.create_DC20(_2735_v45);
        _2736_v46 = _2736_v46;
        if (_this.f7) {
          let _2737_v47;
          let _nw467 = new _module.C5();
          _nw467.__ctor(true);
          _2737_v47 = _nw467;
          _2679_v0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2709_v25,_2679_v0)).length);
          let _2738_v48;
          _2738_v48 = _module.D0.create_DC1(_2679_v0);
          _2738_v48 = _2738_v48;
          (_this).f7 = (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))];
          let _2739_v49;
          let _nw468 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _2739_v49 = _nw468;
          let _index434 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_2739_v49).length));
          (_2739_v49)[_index434] = (_dafny.ZERO).minus(_2679_v0);
          let _2740_v50;
          _2740_v50 = _dafny.Seq.UnicodeFromString("yiih");
          let _index435 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_2739_v49).length));
          (_2739_v49)[_index435] = new BigNumber((_2740_v50).length);
        } else {
          let _nw469 = Array((new BigNumber(9)).toNumber()).fill(false);
          _2709_v25 = _nw469;
          let _2741_v51;
          _2741_v51 = _dafny.Seq.UnicodeFromString("aykqqpd");
          _2741_v51 = _2741_v51;
          (_this).f7 = (_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))];
          let _2742_v52;
          _2742_v52 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
          let _2743_v53;
          let _nw470 = Array((new BigNumber(17)).toNumber());
          _nw470[(_dafny.ZERO).toNumber()] = _2679_v0;
          _nw470[(_dafny.ONE).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(2)).toNumber()] = new BigNumber(-997);
          _nw470[(new BigNumber(3)).toNumber()] = new BigNumber((_2741_v51).length);
          _nw470[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2742_v52).cardinality()), _2679_v0);
          _nw470[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2741_v51).length));
          _nw470[(new BigNumber(6)).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(7)).toNumber()] = (((_2680_v1)[_module.__default.safeIndex(_2679_v0, new BigNumber((_2680_v1).length))]) ? (_2679_v0) : (_2679_v0));
          _nw470[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2710_v26).length), _2679_v0);
          _nw470[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_2741_v51).length)), new BigNumber(802));
          _nw470[(new BigNumber(10)).toNumber()] = new BigNumber(-196);
          _nw470[(new BigNumber(11)).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(12)).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(13)).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(14)).toNumber()] = _2679_v0;
          _nw470[(new BigNumber(15)).toNumber()] = new BigNumber((_2741_v51).length);
          _nw470[(new BigNumber(16)).toNumber()] = _2679_v0;
          _2743_v53 = _nw470;
          let _index436 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_2743_v53).length));
          (_2743_v53)[_index436] = _2679_v0;
          let _index437 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_2743_v53).length));
          (_2743_v53)[_index437] = _2679_v0;
          let _2744_v54;
          _2744_v54 = _module.D1.create_DC4(_this.f7, _this.f7);
          let _2745_v55;
          _2745_v55 = _module.D1.create_DC6(_2744_v54);
          let _2746_v56;
          _2746_v56 = _module.D18.create_DC51(_2680_v1, _2679_v0, true);
          let _2747_v57;
          _2747_v57 = _dafny.Seq.of(_dafny.Seq.of((_2746_v56).dtor_cf87), _2680_v1, _2680_v1);
          let _rhs424 = _dafny.Seq.update(_2680_v1, _module.__default.safeIndex(_2679_v0, new BigNumber((_2680_v1).length)), !((_2709_v25)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_2709_v25).length))]));
          let _rhs425 = _2743_v53;
          let _rhs426 = (new BigNumber((_2747_v57).length)).isLessThanOrEqualTo(_2679_v0);
          let _rhs427 = _2745_v55;
          let _lhs255 = _this;
          _2680_v1 = _rhs424;
          _2743_v53 = _rhs425;
          _lhs255.f7 = _rhs426;
          _2745_v55 = _rhs427;
        }
      } else {
        let _2748_v58;
        _2748_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f7);
        let _2749_v59;
        let _nw471 = Array((new BigNumber(7)).toNumber()).fill(false);
        _2749_v59 = _nw471;
        let _2750_v60;
        let _nw472 = new _module.C2();
        _nw472.__ctor(_dafny.Set.fromElements(_2749_v59), new _dafny.CodePoint('t'.codePointAt(0)), _this.f7);
        _2750_v60 = _nw472;
        let _2751_v61;
        _2751_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2750_v60);
        let _2752_v62;
        _2752_v62 = _dafny.Seq.UnicodeFromString("xiamatg");
        _2748_v58 = (_2748_v58).update((_2679_v0).isLessThan(new BigNumber((_2751_v61).length)), !_dafny.areEqual(_2752_v62, _2752_v62));
        let _2753_v63;
        _2753_v63 = _dafny.Seq.of(_2752_v62, _2752_v62);
        let _2754_v64;
        let _nw473 = new _module.C0();
        _nw473.__ctor((_2679_v0).isLessThan(_2679_v0), _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), ((_2755_v2) => function (_2756_i4) {
          return _2755_v2;
        })(_2681_v2)), (_2753_v63)[_module.__default.safeIndex(_2679_v0, new BigNumber((_2753_v63).length))]));
        _2754_v64 = _nw473;
        (_this).f7 = true;
        (_2754_v64).f13 = _2754_v64.f13;
        _2679_v0 = (_2679_v0).multipliedBy(((_this).fm10(_this.f7, globalState)).plus(_2679_v0));
      }
      _2679_v0 = _2679_v0;
      _2679_v0 = _2679_v0;
      _2679_v0 = new BigNumber(348);
      let _2757_v65;
      let _init90 = ((_2758_v0) => function (_2759_i5) {
        return _module.D9.create_DC26(_dafny.Seq.of(new BigNumber(920), _2758_v0));
      })(_2679_v0);
      let _nw474 = Array((new BigNumber(4)).toNumber());
      for (let _i0_90 = 0; _i0_90 < new BigNumber(_nw474.length); _i0_90++) {
        _nw474[_i0_90] = _init90(new BigNumber(_i0_90));
      }
      _2757_v65 = _nw474;
      let _2760_v66;
      _2760_v66 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(999)), ((_2761_v0) => function (_2762_i6) {
        return _2761_v0;
      })(_2679_v0))).length));
      let _index438 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_2757_v65).length));
      (_2757_v65)[_index438] = _module.D9.create_DC26(_2760_v66);
      let _2763_v67;
      _2763_v67 = _module.D9.create_DC26(_2760_v66);
      let _2764_v68;
      _2764_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2679_v0,(_module.D9.create_DC26(_2760_v66)).dtor_cf45);
      let _2765_v69;
      _2765_v69 = _dafny.Seq.UnicodeFromString("msetw");
      let _pat_let_tv65 = _2764_v68;
      let _pat_let_tv66 = _2765_v69;
      let _pat_let_tv67 = _2760_v66;
      let _pat_let_tv68 = _2764_v68;
      let _pat_let_tv69 = _2765_v69;
      let _pat_let_tv70 = _2679_v0;
      let _index439 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_2757_v65).length));
      (_2757_v65)[_index439] = function (_pat_let77_0) {
        return function (_2768_dt__update__tmp_h1) {
          return function (_pat_let80_0) {
            return function (_2771_dt__update_hcf45_h1) {
              return _module.D9.create_DC26(_2771_dt__update_hcf45_h1);
            }(_pat_let80_0);
          }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), ((_2769_v0) => function (_2770_i7) {
            return _2769_v0;
          })(_pat_let_tv70)));
        }(_pat_let77_0);
      }(function (_pat_let78_0) {
        return function (_2766_dt__update__tmp_h0) {
          return function (_pat_let79_0) {
            return function (_2767_dt__update_hcf45_h0) {
              return _module.D9.create_DC26(_2767_dt__update_hcf45_h0);
            }(_pat_let79_0);
          }((((_pat_let_tv68).contains(new BigNumber((_pat_let_tv69).length))) ? ((_pat_let_tv65).get(new BigNumber((_pat_let_tv66).length))) : (_pat_let_tv67)));
        }(_pat_let78_0);
      }(_2763_v67));
      return;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      (_this).m5(globalState);
      let _2772_i0;
      _2772_i0 = _dafny.ZERO;
      L18: {
        while (!(new BigNumber(641)).isEqualTo((_dafny.ZERO).minus((_this).fm10(!(_this.f7), globalState)))) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2772_i0)) {
              break L18;
            }
            _2772_i0 = (_2772_i0).plus(_dafny.ONE);
            let _2773_v0;
            let _nw475 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
            _2773_v0 = _nw475;
            let _2774_v1;
            _2774_v1 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_2775_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})).length));
            let _2776_v2;
            _2776_v2 = _dafny.Seq.of(_2774_v1);
            let _2777_v3;
            _2777_v3 = _dafny.MultiSet.fromElements(_this.f7, _module.__default.fm1(_2776_v2, p0, new BigNumber((p1).length), globalState));
            let _2778_v4;
            _2778_v4 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f7, false), _2777_v3);
            let _index440 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2773_v0).length));
            (_2773_v0)[_index440] = (_2778_v4)[_module.__default.safeIndex(p0, new BigNumber((_2778_v4).length))];
            let _2779_v5;
            _2779_v5 = _module.D2.create_DC8(_this.f7, p1);
            let _2780_v6;
            _2780_v6 = _dafny.Seq.of((_2779_v5).dtor_cf14);
            let _index441 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2773_v0).length));
            (_2773_v0)[_index441] = (_2777_v3).update((_2780_v6)[_module.__default.safeIndex(p0, new BigNumber((_2780_v6).length))], _module.__default.abs((p0).multipliedBy(p0)));
            let _2781_v7;
            _2781_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
            (_this).f7 = (((_2781_v7).contains(_this.f7)) ? ((_2781_v7).get(_this.f7)) : (_this.f7));
            let _2782_v8;
            let _nw476 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2782_v8 = _nw476;
            let _2783_v9;
            _2783_v9 = new _dafny.CodePoint('a'.codePointAt(0));
            let _index442 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_2782_v8).length));
            (_2782_v8)[_index442] = _2783_v9;
            let _2784_v10;
            let _init91 = ((_2785_p0) => function (_2786_i2) {
              return (_2785_p0).isLessThan(new BigNumber(41));
            })(p0);
            let _nw477 = Array((new BigNumber(8)).toNumber());
            for (let _i0_91 = 0; _i0_91 < new BigNumber(_nw477.length); _i0_91++) {
              _nw477[_i0_91] = _init91(new BigNumber(_i0_91));
            }
            _2784_v10 = _nw477;
            let _2787_v11;
            _2787_v11 = _dafny.Set.fromElements(true);
            let _2788_v12;
            _2788_v12 = _dafny.Seq.of(_2787_v11);
            let _index443 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2784_v10).length));
            (_2784_v10)[_index443] = (new BigNumber((p1).length)).isLessThan(new BigNumber((_2788_v12).length));
            let _2789_v13;
            let _nw478 = Array((new BigNumber(4)).toNumber());
            _nw478[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((function (_pat_let81_0) {
              return function (_2790_dt__update__tmp_h0) {
                return function (_pat_let82_0) {
                  return function (_2791_dt__update_hcf14_h0) {
                    return _module.D2.create_DC8(_2791_dt__update_hcf14_h0, (_2790_dt__update__tmp_h0).dtor_cf15);
                  }(_pat_let82_0);
                }(_this.f7);
              }(_pat_let81_0);
            }(_module.D2.create_DC8(_this.f7, p1))).dtor_cf15, (_2779_v5).dtor_cf15);
            _nw478[(_dafny.ONE).toNumber()] = p1;
            _nw478[(new BigNumber(2)).toNumber()] = p1;
            _nw478[(new BigNumber(3)).toNumber()] = p1;
            _2789_v13 = _nw478;
            let _index444 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2789_v13).length));
            (_2789_v13)[_index444] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), function (_2792_i3) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            });
            let _index445 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_2782_v8).length));
            let _index446 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2784_v10).length));
            let _index447 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2789_v13).length));
            let _rhs428 = (_dafny.ZERO).minus(p0);
            let _rhs429 = (p1)[_module.__default.safeIndex(p0, new BigNumber((p1).length))];
            let _rhs430 = _this.f7;
            let _rhs431 = _dafny.Seq.UnicodeFromString("fgnlwup");
            let _lhs256 = _2782_v8;
            let _lhs257 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_2782_v8).length));
            let _lhs258 = _2784_v10;
            let _lhs259 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2784_v10).length));
            let _lhs260 = _2789_v13;
            let _lhs261 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_2789_v13).length));
            r1 = _rhs428;
            _lhs256[_lhs257] = _rhs429;
            _lhs258[_lhs259] = _rhs430;
            _lhs260[_lhs261] = _rhs431;
            let _2793_v14;
            let _nw479 = Array((new BigNumber(6)).toNumber());
            _nw479[(_dafny.ZERO).toNumber()] = p0;
            _nw479[(_dafny.ONE).toNumber()] = p0;
            _nw479[(new BigNumber(2)).toNumber()] = new BigNumber(-651);
            _nw479[(new BigNumber(3)).toNumber()] = p0;
            _nw479[(new BigNumber(4)).toNumber()] = p0;
            _nw479[(new BigNumber(5)).toNumber()] = new BigNumber(801);
            _2793_v14 = _nw479;
            let _2794_v15;
            _2794_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2793_v14,_dafny.Set.fromElements(p0));
            let _2795_v16;
            _2795_v16 = _module.D13.create_DC34(_2794_v15);
            let _2796_v17;
            _2796_v17 = _dafny.Seq.of(_2795_v16, _2795_v16, _2795_v16, _2795_v16, _2795_v16);
            let _2797_v18;
            _2797_v18 = _module.D17.create_DC49(p0, _2796_v17, p0, (_2784_v10)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2784_v10).length))]);
            r1 = (_2797_v18).dtor_cf82;
          }
        }
      }
      let _2798_v19;
      _2798_v19 = _dafny.Set.fromElements(_this.f7, _this.f7);
      let _2799_v20;
      _2799_v20 = new _dafny.CodePoint('n'.codePointAt(0));
      let _rhs432 = (_this).fm9(_this.f7, new BigNumber((_dafny.Set.fromElements(new BigNumber((_2798_v19).length), new BigNumber((_dafny.Seq.UnicodeFromString("qnyywoju")).length), new BigNumber(149), p0, new BigNumber(-648))).length), _2799_v20, globalState);
      let _rhs433 = new BigNumber(310);
      r3 = _rhs432;
      r3 = _rhs433;
      let _2800_v21;
      _2800_v21 = _dafny.Set.fromElements(p0, (_this).fm9(_this.f7, p0, _2799_v20, globalState));
      let _2801_v22;
      _2801_v22 = _dafny.Seq.of(p0);
      let _2802_v23;
      _2802_v23 = _module.D1.create_DC5(_this.f7, _2801_v22, p0);
      let _2803_v24;
      _2803_v24 = _module.D12.create_DC31(_module.D9.create_DC26(_2801_v22), (_2802_v23).dtor_cf9);
      let _2804_v25;
      _2804_v25 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-892)), ((_2805_v20) => function (_2806_i4) {
        return _2805_v20;
      })(_2799_v20)));
      let _2807_v26;
      _2807_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_2808_v20) => function (_2809_i5) {
        return _2808_v20;
      })(_2799_v20)));
      let _2810_v27;
      _2810_v27 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(p1)).length)), p1), _2804_v25, _2804_v25, _dafny.Seq.of((((_2807_v26).contains(_this.f7)) ? ((_2807_v26).get(_this.f7)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_2811_v20) => function (_2812_i6) {
        return _2811_v20;
      })(_2799_v20)))), p1));
      let _2813_v28;
      _2813_v28 = _module.D0.create_DC0(p0);
      let _2814_v29;
      _2814_v29 = _dafny.Seq.of(_2813_v28, _2813_v28);
      let _2815_v30;
      let _init92 = ((_2816_p0) => function (_2817_i8) {
        return (_2817_i8).multipliedBy(_2816_p0);
      })(p0);
      let _nw480 = Array((new BigNumber(19)).toNumber());
      for (let _i0_92 = 0; _i0_92 < new BigNumber(_nw480.length); _i0_92++) {
        _nw480[_i0_92] = _init92(new BigNumber(_i0_92));
      }
      _2815_v30 = _nw480;
      let _2818_v31;
      _2818_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2815_v30,_2800_v21);
      let _2819_v32;
      _2819_v32 = _module.D13.create_DC34(_2818_v31);
      let _2820_v33;
      _2820_v33 = _dafny.Seq.of(_2819_v32);
      let _2821_v34;
      _2821_v34 = _dafny.Seq.of(_this.f7);
      let _2822_v35;
      _2822_v35 = _module.D17.create_DC49(p0, _2820_v33, new BigNumber((_2821_v34).length), _this.f7);
      let _2823_v36;
      let _nw481 = Array((new BigNumber(21)).toNumber());
      _nw481[(_dafny.ZERO).toNumber()] = _this.f7;
      _nw481[(_dafny.ONE).toNumber()] = _this.f7;
      _nw481[(new BigNumber(2)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(3)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(4)).toNumber()] = (_2803_v24).dtor_cf54;
      _nw481[(new BigNumber(5)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.update((_2810_v27)[_module.__default.safeIndex(p0, new BigNumber((_2810_v27).length))], _module.__default.safeIndex(p0, new BigNumber(((_2810_v27)[_module.__default.safeIndex(p0, new BigNumber((_2810_v27).length))]).length)), p1), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update((_2810_v27)[_module.__default.safeIndex(p0, new BigNumber((_2810_v27).length))], _module.__default.safeIndex(p0, new BigNumber(((_2810_v27)[_module.__default.safeIndex(p0, new BigNumber((_2810_v27).length))]).length)), p1)).length)), _dafny.Seq.UnicodeFromString("eopjgc")), p1);
      _nw481[(new BigNumber(6)).toNumber()] = (_module.__default.fm1(_2814_v29, p0, p0, globalState)) || (_this.f7);
      _nw481[(new BigNumber(7)).toNumber()] = !(true);
      _nw481[(new BigNumber(8)).toNumber()] = !(_this.f7) || (_this.f7);
      _nw481[(new BigNumber(9)).toNumber()] = ((_this.f7) ? (_this.f7) : (_this.f7));
      _nw481[(new BigNumber(10)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(11)).toNumber()] = true;
      _nw481[(new BigNumber(12)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(13)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(14)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-331)), ((_2824_p1) => function (_2825_i7) {
        return _2824_p1;
      })(p1)), _2804_v25);
      _nw481[(new BigNumber(15)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(16)).toNumber()] = (_2822_v35).dtor_cf83;
      _nw481[(new BigNumber(17)).toNumber()] = _this.f7;
      _nw481[(new BigNumber(18)).toNumber()] = _dafny.Seq.contains(p1, _2799_v20);
      _nw481[(new BigNumber(19)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), ((_2826_v20) => function (_2827_i9) {
        return _2826_v20;
      })(_2799_v20)), _2799_v20);
      _nw481[(new BigNumber(20)).toNumber()] = true;
      _2823_v36 = _nw481;
      let _index448 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_2823_v36).length));
      (_2823_v36)[_index448] = (_this).fm8(p0, globalState);
      let _2828_v37;
      _2828_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_this.f7));
      let _index449 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_2823_v36).length));
      let _rhs434 = (_2800_v21).Intersect((_2800_v21).Union(_2800_v21));
      let _rhs435 = (p0).isLessThan(new BigNumber((_2801_v22).length));
      let _rhs436 = (((_2828_v37).contains(p0)) ? ((_2828_v37).get(p0)) : (_this.f7));
      let _rhs437 = _2798_v19;
      let _rhs438 = _this.f7;
      let _lhs262 = _this;
      let _lhs263 = _2823_v36;
      let _lhs264 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_2823_v36).length));
      let _lhs265 = _this;
      _2800_v21 = _rhs434;
      _lhs262.f7 = _rhs435;
      _lhs263[_lhs264] = _rhs436;
      _2798_v19 = _rhs437;
      _lhs265.f7 = _rhs438;
      r2 = (_2823_v36)[_module.__default.safeIndex(new BigNumber(473), new BigNumber((_2823_v36).length))];
      let _2829_v38;
      _2829_v38 = _dafny.Set.fromElements(_2823_v36, _2823_v36, _2823_v36, _2823_v36);
      let _2830_v39;
      let _nw482 = new _module.C2();
      _nw482.__ctor((_2829_v38).Intersect(_2829_v38), _2799_v20, !(_this.f7));
      _2830_v39 = _nw482;
      let _2831_v40;
      let _nw483 = Array((new BigNumber(8)).toNumber()).fill([]);
      _2831_v40 = _nw483;
      r0 = _2831_v40;
      r1 = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
      r2 = ((_this.f7) ? (false) : (_this.f7));
      let _2832_v41;
      _2832_v41 = _module.D9.create_DC25();
      let _pat_let_tv71 = p0;
      let _pat_let_tv72 = p0;
      let _pat_let_tv73 = p0;
      let _pat_let_tv74 = p0;
      let _pat_let_tv75 = p0;
      r3 = function (_source20) {
        if (_source20.is_DC24) {
          return _module.__default.safeModuloInt(_pat_let_tv71, _pat_let_tv72);
        } else if (_source20.is_DC25) {
          return _pat_let_tv73;
        } else if (_source20.is_DC26) {
          let _2833___mcc_h0 = (_source20).cf45;
          let _2834_cf45 = _2833___mcc_h0;
          return (_dafny.ZERO).minus((_dafny.ZERO).minus(_pat_let_tv74));
        } else {
          let _2835___mcc_h1 = (_source20).cf44;
          let _2836_cf44 = _2835___mcc_h1;
          return _pat_let_tv75;
        }
      }(_2832_v41);
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let _2837_v0;
      let _nw484 = new _module.C0();
      _nw484.__ctor(_this.f7, _this.f7);
      _2837_v0 = _nw484;
      let _2838_v1;
      _2838_v1 = new BigNumber(593);
      _2838_v1 = (_2838_v1).plus(_module.__default.safeModuloInt(_2838_v1, new BigNumber(-343)));
      let _2839_v2;
      _2839_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2837_v0.f13,true);
      let _2840_i0;
      _2840_i0 = _dafny.ZERO;
      L19: {
        while (((new BigNumber(542)).multipliedBy(new BigNumber((_2839_v2).length))).isLessThan(_2838_v1)) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2840_i0)) {
              break L19;
            }
            _2840_i0 = (_2840_i0).plus(_dafny.ONE);
            (_2837_v0).f13 = _this.f7;
            let _2841_v3;
            _2841_v3 = (_2839_v2).Merge(_2839_v2);
            let _2842_v4;
            let _nw485 = new _module.C3();
            _nw485.__ctor(_module.__default.fm0(globalState), _this.f7);
            _2842_v4 = _nw485;
            let _2843_v5;
            let _nw486 = new _module.C6();
            _nw486.__ctor(!((_2837_v0).f12), _2837_v0.f13);
            _2843_v5 = _nw486;
            let _rhs439 = _2841_v3;
            let _rhs440 = !((_2837_v0).f12);
            let _rhs441 = _2838_v1;
            let _rhs442 = _2842_v4;
            let _rhs443 = _2843_v5;
            let _lhs266 = _2837_v0;
            _2841_v3 = _rhs439;
            _lhs266.f13 = _rhs440;
            _2838_v1 = _rhs441;
            _2842_v4 = _rhs442;
            _2843_v5 = _rhs443;
            let _2844_v6;
            _2844_v6 = new _dafny.CodePoint('i'.codePointAt(0));
            let _2845_v7;
            _2845_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2837_v0).f12,_2838_v1);
            let _2846_v8;
            _2846_v8 = _dafny.Seq.of(_2845_v7);
            let _2847_v9;
            _2847_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2844_v6,(_2846_v8)[_module.__default.safeIndex(_2838_v1, new BigNumber((_2846_v8).length))]);
            let _2848_v10;
            _2848_v10 = _dafny.Seq.of(_2838_v1, _2838_v1);
            _2847_v9 = _module.__default.fm57(_dafny.Seq.UnicodeFromString("oqkrpxyi"), _2838_v1, (_2848_v10)[_module.__default.safeIndex(_2838_v1, new BigNumber((_2848_v10).length))], globalState);
            let _2849_v11;
            _2849_v11 = _dafny.Seq.UnicodeFromString("ivk");
            let _2850_v12;
            _2850_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2838_v1, _2838_v1),_2849_v11);
            let _2851_v13;
            _2851_v13 = _dafny.Set.fromElements(_2838_v1, _2838_v1, _2838_v1, _2838_v1, _2838_v1);
            _2850_v12 = (_2850_v12).update((_2851_v13).Intersect(_2851_v13), _2849_v11);
          }
        }
      }
      (_this).f7 = _2837_v0.f13;
      (_2837_v0).f13 = (_2837_v0).f12;
      let _hi18 = new BigNumber(-902);
      for (let _2852_i1 = _2838_v1; _2852_i1.isLessThan(_hi18); _2852_i1 = _2852_i1.plus(_dafny.ONE)) {
        let _2853_v14;
        let _nw487 = new _module.C1();
        _nw487.__ctor(_2837_v0.f13);
        _2853_v14 = _nw487;
        let _2854_v15;
        _2854_v15 = _dafny.Seq.UnicodeFromString("svdmeqnc");
        let _2855_v16;
        let _nw488 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2855_v16 = _nw488;
        let _2856_v17;
        _2856_v17 = _dafny.Seq.of(_2855_v16);
        let _2857_v18;
        _2857_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2854_v15).length),(_2856_v17)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_2856_v17).length))]);
        let _2858_v19;
        let _2859_v20;
        let _out6;
        let _out7;
        let _outcollector2 = (_this).m6((new BigNumber(62)).plus(_2852_i1), _2857_v18, _2838_v1, (_2838_v1).isLessThan(_2852_i1), globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _2858_v19 = _out6;
        _2859_v20 = _out7;
        let _2860_v21;
        _2860_v21 = _dafny.Seq.of(_2838_v1);
        let _2861_v23;
        let _nw489 = Array((new BigNumber(11)).toNumber());
        _nw489[(_dafny.ZERO).toNumber()] = _2837_v0.f13;
        _nw489[(_dafny.ONE).toNumber()] = _2859_v20;
        _nw489[(new BigNumber(2)).toNumber()] = false;
        _nw489[(new BigNumber(3)).toNumber()] = _2837_v0.f13;
        _nw489[(new BigNumber(4)).toNumber()] = (_2837_v0).f12;
        _nw489[(new BigNumber(5)).toNumber()] = (_2837_v0).f12;
        _nw489[(new BigNumber(6)).toNumber()] = (function () {
          let _coll54 = new _dafny.Set();
          for (const _compr_54 of (_2860_v21).Elements) {
            let _2862_v22 = _compr_54;
            if (_dafny.Seq.contains(_2860_v21, _2862_v22)) {
              _coll54.add(_module.__default.safeDivisionInt(_2862_v22, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(110), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber((_dafny.Seq.of(true, true)).length)))).length)));
            }
          }
          return _coll54;
        }()).IsDisjointFrom(_2858_v19);
        _nw489[(new BigNumber(7)).toNumber()] = _this.f7;
        _nw489[(new BigNumber(8)).toNumber()] = false;
        _nw489[(new BigNumber(9)).toNumber()] = false;
        _nw489[(new BigNumber(10)).toNumber()] = _2837_v0.f13;
        _2861_v23 = _nw489;
        let _index450 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length));
        (_2861_v23)[_index450] = true;
        let _index451 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length));
        let _rhs444 = _2838_v1;
        let _rhs445 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2854_v15, _2854_v15), _2854_v15)).length);
        let _rhs446 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_2838_v1), _2838_v1)).isLessThan(_2838_v1);
        let _lhs267 = _2861_v23;
        let _lhs268 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length));
        _2838_v1 = _rhs444;
        _2838_v1 = _rhs445;
        _lhs267[_lhs268] = _rhs446;
        if (_this.f7) {
          _2838_v1 = _2838_v1;
          _2838_v1 = new BigNumber((_2854_v15).length);
          let _2863_v24;
          _2863_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2855_v16,_2858_v19);
          let _2864_v25;
          _2864_v25 = _module.D13.create_DC34(_2863_v24);
          let _2865_v26;
          _2865_v26 = _dafny.Seq.of(_2864_v25, _2864_v25, _2864_v25, _2864_v25);
          let _2866_v27;
          _2866_v27 = _dafny.Seq.of((_2861_v23)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length))]);
          (_2837_v0).f13 = (_module.D17.create_DC49(_module.__default.fm0(globalState), _2865_v26, _2838_v1, (_2866_v27)[_module.__default.safeIndex(_2852_i1, new BigNumber((_2866_v27).length))])).dtor_cf83;
          let _2867_v28;
          _2867_v28 = _dafny.Map.Empty.slice().updateUnsafe((_2861_v23)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length))],_2852_i1);
          _2867_v28 = (_2867_v28).update((_2837_v0).f12, new BigNumber(((_2853_v14).fm31(_2838_v1, !((_2837_v0).f12), _2854_v15, (_2837_v0).f12, globalState)).length));
          let _2868_v29;
          _2868_v29 = _dafny.Seq.of(_2861_v23);
          _2861_v23 = (_2868_v29)[_module.__default.safeIndex(_2838_v1, new BigNumber((_2868_v29).length))];
        } else {
          let _2869_v30;
          let _nw490 = new _module.C6();
          _nw490.__ctor(_2837_v0.f13, !(false));
          _2869_v30 = _nw490;
          let _index452 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_2855_v16).length));
          (_2855_v16)[_index452] = _2852_i1;
          let _index453 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_2855_v16).length));
          (_2855_v16)[_index453] = _2838_v1;
          let _2870_v31;
          _2870_v31 = _dafny.Map.Empty.slice().updateUnsafe((_2837_v0).f12,_2861_v23);
          let _2871_v32;
          _2871_v32 = _module.D3.create_DC10(_2861_v23);
          let _pat_let_tv76 = _2861_v23;
          let _pat_let_tv77 = _2861_v23;
          let _2872_v33;
          let _nw491 = Array((new BigNumber(23)).toNumber());
          _nw491[(_dafny.ZERO).toNumber()] = _module.D3.create_DC10((((_2870_v31).contains((_2837_v0).f12)) ? ((_2870_v31).get((_2837_v0).f12)) : (_2861_v23)));
          _nw491[(_dafny.ONE).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(2)).toNumber()] = _module.D3.create_DC10(_2861_v23);
          _nw491[(new BigNumber(3)).toNumber()] = function (_pat_let83_0) {
            return function (_2873_dt__update__tmp_h0) {
              return function (_pat_let84_0) {
                return function (_2874_dt__update_hcf21_h0) {
                  return _module.D3.create_DC10(_2874_dt__update_hcf21_h0);
                }(_pat_let84_0);
              }(_pat_let_tv76);
            }(_pat_let83_0);
          }(_2871_v32);
          _nw491[(new BigNumber(4)).toNumber()] = function (_pat_let85_0) {
            return function (_2875_dt__update__tmp_h1) {
              return function (_pat_let86_0) {
                return function (_2876_dt__update_hcf21_h1) {
                  return _module.D3.create_DC10(_2876_dt__update_hcf21_h1);
                }(_pat_let86_0);
              }(_pat_let_tv77);
            }(_pat_let85_0);
          }(_2871_v32);
          _nw491[(new BigNumber(5)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(6)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(7)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(8)).toNumber()] = _module.D3.create_DC10(_2861_v23);
          _nw491[(new BigNumber(9)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(10)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(11)).toNumber()] = _module.D3.create_DC10(_2861_v23);
          _nw491[(new BigNumber(12)).toNumber()] = _module.D3.create_DC10(_2861_v23);
          _nw491[(new BigNumber(13)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(14)).toNumber()] = ((true) ? (_2871_v32) : (_2871_v32));
          _nw491[(new BigNumber(15)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(16)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(17)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(18)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(19)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(20)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(21)).toNumber()] = _2871_v32;
          _nw491[(new BigNumber(22)).toNumber()] = _module.D3.create_DC10(_2861_v23);
          _2872_v33 = _nw491;
          let _index454 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2872_v33).length));
          (_2872_v33)[_index454] = _2871_v32;
          let _index455 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2872_v33).length));
          (_2872_v33)[_index455] = _2871_v32;
          (_2837_v0).f13 = (_2861_v23)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2861_v23).length))];
          let _2877_v34;
          _2877_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2837_v0.f13,_2839_v2);
          let _2878_v35;
          _2878_v35 = _dafny.Map.Empty.slice().updateUnsafe((_2855_v16)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_2855_v16).length))],_2877_v34);
          let _2879_v36;
          _2879_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2852_i1,(_2869_v30).f9);
          _2878_v35 = (_2878_v35).update(_module.__default.safeDivisionInt(_2838_v1, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2854_v15).length),_2879_v36)).length))), _2877_v34);
        }
      }
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let _2880_v0;
      _2880_v0 = _dafny.Seq.UnicodeFromString("ck");
      _2880_v0 = _dafny.Seq.UnicodeFromString("hosrfwfoe");
      let _rhs447 = _dafny.Seq.UnicodeFromString("njabs");
      let _rhs448 = _dafny.Seq.Concat(_2880_v0, _dafny.Seq.UnicodeFromString("uesfgn"));
      _2880_v0 = _rhs447;
      _2880_v0 = _rhs448;
      let _2881_v1;
      _2881_v1 = _module.D2.create_DC8(_this.f7, _dafny.Seq.UnicodeFromString("myneoghrv"));
      let _2882_v2;
      _2882_v2 = _dafny.Seq.of((_2881_v1).dtor_cf14);
      r1 = !((_2882_v2)[_module.__default.safeIndex(new BigNumber((_2882_v2).length), new BigNumber((_2882_v2).length))]);
      let _2883_v4;
      let _init93 = ((_2884_p0) => function (_2885_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll55 = new _dafny.Set();
          for (const _compr_55 of (_dafny.Map.Empty.slice().updateUnsafe(_2884_p0,_2884_p0)).Keys.Elements) {
            let _2886_v3 = _compr_55;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_2884_p0,_2884_p0)).contains(_2886_v3)) {
              _coll55.add((_2886_v3).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bcoa")).length)));
            }
          }
          return _coll55;
        }(),true);
      })(p0);
      let _nw492 = Array((new BigNumber(15)).toNumber());
      for (let _i0_93 = 0; _i0_93 < new BigNumber(_nw492.length); _i0_93++) {
        _nw492[_i0_93] = _init93(new BigNumber(_i0_93));
      }
      _2883_v4 = _nw492;
      let _2887_v5;
      _2887_v5 = _dafny.Set.fromElements(p2);
      let _2888_v6;
      _2888_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2887_v5,p3);
      let _index456 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_2883_v4).length));
      (_2883_v4)[_index456] = ((_this.f7) ? (_2888_v6) : (_2888_v6));
      let _2889_v7;
      _2889_v7 = _module.D0.create_DC0(_module.__default.fm0(globalState));
      let _2890_v8;
      _2890_v8 = _dafny.Seq.of(_2889_v7);
      let _2891_v9;
      _2891_v9 = _dafny.Map.Empty.slice().updateUnsafe(((_module.__default.fm1(_2890_v8, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f7),_this.f7)).length), p0, globalState)) ? (p3) : (p3)),(_2888_v6).Merge(_2888_v6));
      let _2892_v10;
      _2892_v10 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
      let _index457 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_2883_v4).length));
      (_2883_v4)[_index457] = (((_2891_v9).contains((_this).fm8((_dafny.ZERO).minus(p0), globalState))) ? ((_2891_v9).get((_this).fm8((_dafny.ZERO).minus(p0), globalState))) : ((_2888_v6).update(_module.__default.fm2(new BigNumber((_2892_v10).cardinality()), _this.f7, globalState), _this.f7)));
      let _2893_i1;
      _2893_i1 = _dafny.ZERO;
      L20: {
        while (true) {
          C20: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2893_i1)) {
              break L20;
            }
            _2893_i1 = (_2893_i1).plus(_dafny.ONE);
            (_this).f7 = _this.f7;
            let _2894_v11;
            _2894_v11 = new BigNumber(-180);
            _2894_v11 = p2;
            let _2895_v12;
            _2895_v12 = new _dafny.CodePoint('u'.codePointAt(0));
            _2895_v12 = _2895_v12;
            _module.__default.m0(new BigNumber(215), _2894_v11, globalState);
          }
        }
      }
      let _2896_v13;
      let _nw493 = new _module.C6();
      _nw493.__ctor(p3, p3);
      _2896_v13 = _nw493;
      r0 = (_2887_v5).Difference(_2887_v5);
      r1 = (_2896_v13).f9;
      return [r0, r1];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return ((_this).f6).minus(new BigNumber(387));
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0))), _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_2897_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _2898_v0;
      _2898_v0 = true;
      let _2899_v1;
      _2899_v1 = _module.D1.create_DC4(_2898_v0, _2898_v0);
      let _source21 = ((_2898_v0) ? (_2899_v1) : (_module.D1.create_DC4(_2898_v0, _2898_v0)));
      if (_source21.is_DC4) {
        let _2900___mcc_h0 = (_source21).cf7;
        let _2901___mcc_h1 = (_source21).cf8;
        let _2902_cf8 = _2901___mcc_h1;
        let _2903_cf7 = _2900___mcc_h0;
        if (_2903_cf7) {
          let _2904_v2;
          _2904_v2 = _dafny.Seq.UnicodeFromString("jqrmug");
          _2904_v2 = _2904_v2;
          r3 = (_dafny.ZERO).minus((_this).f6);
          let _2905_v3;
          _2905_v3 = _module.D0.create_DC0((_this).f6);
          let _2906_v4;
          _2906_v4 = _dafny.MultiSet.fromElements(_2905_v3);
          _2906_v4 = (_2906_v4).Intersect(_dafny.MultiSet.fromElements(_2905_v3, _2905_v3));
          let _2907_v5;
          _2907_v5 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f6), (_this).f6);
          let _2908_v6;
          _2908_v6 = _dafny.Seq.of((_this).f6);
          let _2909_v7;
          _2909_v7 = _dafny.Seq.of(_2898_v0, _2903_cf7);
          let _2910_v8;
          _2910_v8 = _dafny.Set.fromElements(_2903_cf7);
          let _2911_v9;
          let _nw494 = Array((new BigNumber(29)).toNumber());
          _nw494[(_dafny.ZERO).toNumber()] = _2907_v5;
          _nw494[(_dafny.ONE).toNumber()] = ((true) ? (_dafny.MultiSet.FromArray(_2908_v6)) : (_2907_v5));
          _nw494[(new BigNumber(2)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(3)).toNumber()] = (_2907_v5).Union(_2907_v5);
          _nw494[(new BigNumber(4)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_2908_v6);
          _nw494[(new BigNumber(6)).toNumber()] = (_2907_v5).Intersect(_dafny.MultiSet.fromElements((_this).f6));
          _nw494[(new BigNumber(7)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(8)).toNumber()] = (_2907_v5).Intersect(_2907_v5);
          _nw494[(new BigNumber(9)).toNumber()] = (_2907_v5).Difference(_2907_v5);
          _nw494[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_2903_cf7)).cardinality()), (_this).f6, (_this).f6);
          _nw494[(new BigNumber(11)).toNumber()] = (_2907_v5).Intersect(_2907_v5);
          _nw494[(new BigNumber(12)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(((_2902_cf8) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(215)), function (_2912_i0) {
            return (_this).f6;
          })) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-510)), function (_2913_i1) {
            return new BigNumber(714);
          }))));
          _nw494[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f6)).Union(_2907_v5);
          _nw494[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements((_this).f6, (_this).f6);
          _nw494[(new BigNumber(16)).toNumber()] = _module.__default.fm7(_2902_cf8, _2898_v0, _2902_cf8, globalState);
          _nw494[(new BigNumber(17)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(18)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_2909_v7).length), (_this).f6)).update((_this).f6, _module.__default.abs((_dafny.ZERO).minus(new BigNumber(-677))));
          _nw494[(new BigNumber(19)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(20)).toNumber()] = (_2907_v5).Union(_dafny.MultiSet.fromElements((_this).f6));
          _nw494[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.FromArray(_2908_v6);
          _nw494[(new BigNumber(22)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(23)).toNumber()] = _module.__default.fm7(_2898_v0, _2898_v0, _2903_cf7, globalState);
          _nw494[(new BigNumber(24)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(25)).toNumber()] = ((_2907_v5).update(new BigNumber((_2910_v8).length), _module.__default.abs((_this).f6))).Intersect(_2907_v5);
          _nw494[(new BigNumber(26)).toNumber()] = _module.__default.fm7(_2898_v0, _2898_v0, _2898_v0, globalState);
          _nw494[(new BigNumber(27)).toNumber()] = _2907_v5;
          _nw494[(new BigNumber(28)).toNumber()] = _dafny.MultiSet.fromElements((_this).f6, (_this).f6);
          _2911_v9 = _nw494;
          _2911_v9 = _2911_v9;
          let _2914_v10;
          _2914_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2902_cf8,new BigNumber(-326));
          _2914_v10 = (_2914_v10).update(_2898_v0, (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f6, (_this).f6)));
        } else {
          let _2915_v11;
          let _nw495 = new _module.C0();
          _nw495.__ctor(_2903_cf7, _2903_cf7);
          _2915_v11 = _nw495;
          let _2916_v12;
          _2916_v12 = new _dafny.CodePoint('w'.codePointAt(0));
          _2916_v12 = _2916_v12;
          let _2917_v13;
          let _nw496 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2917_v13 = _nw496;
          let _2918_v14;
          let _nw497 = Array((new BigNumber(9)).toNumber()).fill(false);
          _2918_v14 = _nw497;
          let _index458 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_2917_v13).length));
          (_2917_v13)[_index458] = _2918_v14;
          let _index459 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_2917_v13).length));
          (_2917_v13)[_index459] = _2918_v14;
          let _2919_v15;
          _2919_v15 = _dafny.Seq.UnicodeFromString("etpudh");
          let _2920_v16;
          _2920_v16 = _dafny.MultiSet.fromElements(new BigNumber((_2919_v15).length));
          let _2921_v17;
          _2921_v17 = _dafny.MultiSet.fromElements((_2915_v11).f12);
          let _2922_v18;
          _2922_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2915_v11).f12,(_this).f6);
          let _2923_v19;
          _2923_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2903_cf7);
          let _2924_v20;
          let _nw498 = new _module.C6();
          _nw498.__ctor((_2915_v11).f12, _2898_v0);
          _2924_v20 = _nw498;
          let _2925_v21;
          _2925_v21 = _module.D20.create_DC54(_2924_v20);
          let _2926_v22;
          _2926_v22 = _dafny.Seq.of((_this).f6, (_this).f6);
          let _2927_v23;
          let _nw499 = Array((new BigNumber(29)).toNumber());
          _nw499[(_dafny.ZERO).toNumber()] = new BigNumber(357);
          _nw499[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_this).f6, (_this).f6);
          _nw499[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(3)).toNumber()] = _module.__default.fm0(globalState);
          _nw499[(new BigNumber(4)).toNumber()] = new BigNumber((_2920_v16).cardinality());
          _nw499[(new BigNumber(5)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(6)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(7)).toNumber()] = ((_this).f6).minus((_this).f6);
          _nw499[(new BigNumber(8)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(9)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(10)).toNumber()] = (new BigNumber(-569)).minus((_this).f6);
          _nw499[(new BigNumber(11)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(12)).toNumber()] = new BigNumber(515);
          _nw499[(new BigNumber(13)).toNumber()] = ((!((_2915_v11).f12)) ? ((_this).f6) : ((_this).f6));
          _nw499[(new BigNumber(14)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(15)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-843)), ((_2928_v12) => function (_2929_i2) {
            return _2928_v12;
          })(_2916_v12))).length);
          _nw499[(new BigNumber(17)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(18)).toNumber()] = new BigNumber((_2921_v17).cardinality());
          _nw499[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt((_this).f6, new BigNumber((_2922_v18).length));
          _nw499[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
          _nw499[(new BigNumber(21)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(22)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_2915_v11).fm17((_this).f6, _2898_v0, globalState)).length));
          _nw499[(new BigNumber(24)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(25)).toNumber()] = (new BigNumber((_2923_v19).length)).multipliedBy(new BigNumber(-587));
          _nw499[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2925_v21).dtor_cf89,_dafny.MultiSet.FromArray(_dafny.Seq.update(_2926_v22, _module.__default.safeIndex((_this).f6, new BigNumber((_2926_v22).length)), (_this).f6)))).length);
          _nw499[(new BigNumber(27)).toNumber()] = (_this).f6;
          _nw499[(new BigNumber(28)).toNumber()] = (_this).f6;
          _2927_v23 = _nw499;
          let _2930_v24;
          _2930_v24 = _dafny.Seq.of(_2918_v14);
          let _index460 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_2927_v23).length));
          (_2927_v23)[_index460] = (new BigNumber((_2930_v24).length)).multipliedBy((_this).f6);
          let _index461 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_2927_v23).length));
          (_2927_v23)[_index461] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f6).plus(new BigNumber(-563)))));
          let _2931_v25;
          _2931_v25 = _module.D0.create_DC1(((_this).f6).minus((_2927_v23)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_2927_v23).length))]));
          _2931_v25 = ((_2898_v0) ? (_2931_v25) : (_module.__default.fm58(_2903_cf7, _2915_v11.f13, globalState)));
        }
        let _2932_v26;
        let _nw500 = Array((new BigNumber(2)).toNumber());
        _nw500[(_dafny.ZERO).toNumber()] = false;
        _nw500[(_dafny.ONE).toNumber()] = _2898_v0;
        _2932_v26 = _nw500;
        let _2933_v27;
        _2933_v27 = _dafny.Set.fromElements(_2932_v26);
        let _2934_v28;
        _2934_v28 = new _dafny.CodePoint('q'.codePointAt(0));
        let _2935_v29;
        let _nw501 = new _module.C2();
        _nw501.__ctor(_2933_v27, _2934_v28, _2902_cf8);
        _2935_v29 = _nw501;
        let _2936_v30;
        _2936_v30 = _dafny.Seq.of(_2935_v29, _2935_v29);
        let _2937_v31;
        _2937_v31 = _module.D0.create_DC0((_this).f6);
        let _2938_v32;
        _2938_v32 = _dafny.Seq.of(_2937_v31);
        let _2939_v33;
        _2939_v33 = _dafny.Seq.of(_2902_cf8);
        let _2940_v34;
        _2940_v34 = _dafny.Seq.of(_module.__default.fm1(_2938_v32, new BigNumber((_2939_v33).length), new BigNumber(-168), globalState));
        let _2941_v35;
        _2941_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_2936_v30, _dafny.Seq.update(_2936_v30, _module.__default.safeIndex((_this).f6, new BigNumber((_2936_v30).length)), _2935_v29)), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_2936_v30, _dafny.Seq.update(_2936_v30, _module.__default.safeIndex((_this).f6, new BigNumber((_2936_v30).length)), _2935_v29))).length)), _2935_v29),(_2940_v34)[_module.__default.safeIndex((_this).f6, new BigNumber((_2940_v34).length))]);
        let _2942_v36;
        _2942_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2902_cf8,true);
        let _2943_v37;
        let _nw502 = new _module.C6();
        _nw502.__ctor(false, _module.__default.fm1(_2938_v32, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2902_cf8,_2942_v36)).length), (_this).f6, globalState));
        _2943_v37 = _nw502;
        let _2944_v38;
        _2944_v38 = _dafny.MultiSet.fromElements(_2943_v37, _2943_v37);
        _2941_v35 = (_2941_v35).update(_dafny.Seq.update(_2936_v30, _module.__default.safeIndex((_this).f6, new BigNumber((_2936_v30).length)), _2935_v29), (_2944_v38).IsDisjointFrom(_2944_v38));
        let _2945_v39;
        let _init94 = function (_2946_i3) {
          return _module.D2.create_DC7(_dafny.Set.fromElements((_this).f6, (_this).f6));
        };
        let _nw503 = Array((new BigNumber(21)).toNumber());
        for (let _i0_94 = 0; _i0_94 < new BigNumber(_nw503.length); _i0_94++) {
          _nw503[_i0_94] = _init94(new BigNumber(_i0_94));
        }
        _2945_v39 = _nw503;
        let _2947_v40;
        let _nw504 = new _module.C4();
        _nw504.__ctor(_2945_v39, _2902_cf8);
        _2947_v40 = _nw504;
        let _2948_v41;
        _2948_v41 = _dafny.Set.fromElements(_2934_v28, _2934_v28);
        if (!(!(_2948_v41).contains(_2934_v28))) {
          let _2949_v42;
          _2949_v42 = _dafny.MultiSet.fromElements((_this).f6);
          let _rhs449 = _2898_v0;
          let _rhs450 = _2935_v29.f7;
          let _rhs451 = !(!(_2949_v42).contains(new BigNumber((_dafny.Seq.of((_this).f6, new BigNumber(912))).length))) || (!(_dafny.Seq.IsProperPrefixOf(_2939_v33, _2939_v33)));
          r2 = _rhs449;
          r0 = _rhs450;
          _2903_cf7 = _rhs451;
          let _2950_v43;
          let _nw505 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _2950_v43 = _nw505;
          let _index462 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length));
          (_2950_v43)[_index462] = (_this).f6;
          let _index463 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length));
          (_2950_v43)[_index463] = new BigNumber(727);
          let _2951_v44;
          _2951_v44 = _dafny.Set.fromElements(_2950_v43, _2950_v43);
          let _2952_v45;
          _2952_v45 = _module.D0.create_DC2(_2935_v29.f7, (_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))], _2951_v44, (_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]);
          let _2953_v46;
          _2953_v46 = _dafny.Set.fromElements(_2935_v29.f7, (_2952_v45).dtor_cf2, (_2943_v37).f9, (_2943_v37).f9, _2935_v29.f7);
          let _2954_v47;
          _2954_v47 = _dafny.Set.fromElements(_2953_v46, _2953_v46);
          let _2955_v48;
          let _nw506 = new _module.C2();
          _nw506.__ctor(_2933_v27, _2934_v28, (_2943_v37).f9);
          _2955_v48 = _nw506;
          let _2956_v49;
          _2956_v49 = _dafny.Map.Empty.slice().updateUnsafe((_2954_v47).IsDisjointFrom(_2954_v47),_2955_v48);
          let _2957_v50;
          _2957_v50 = _dafny.Seq.UnicodeFromString("ubha");
          let _2958_v51;
          let _nw507 = new _module.C5();
          _nw507.__ctor(_2935_v29.f7);
          _2958_v51 = _nw507;
          let _2959_v52;
          _2959_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2937_v31,(_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]);
          let _2960_v53;
          _2960_v53 = _module.D2.create_DC9(_module.__default.safeModuloInt((_this).f6, (_dafny.ZERO).minus((_this).f6)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.__default.fm1(_dafny.Seq.of(_module.D0.create_DC0((_this).f6)), (_this).f6, (_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))], globalState))).length))), _2957_v50, ((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2958_v51,_2953_v46)).length)), _2959_v52);
          let _2961_v54;
          _2961_v54 = _dafny.Map.Empty.slice().updateUnsafe((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))],(_this).f6);
          let _rhs452 = _2956_v49;
          let _rhs453 = _2960_v53;
          let _rhs454 = (((_2961_v54).contains((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))])) ? ((_2961_v54).get((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))])) : ((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]));
          _2956_v49 = _rhs452;
          _2960_v53 = _rhs453;
          r1 = _rhs454;
          let _2962_v55;
          _2962_v55 = _dafny.Map.Empty.slice().updateUnsafe((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))],_2903_cf7);
          let _2963_v56;
          _2963_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2962_v55,_2898_v0);
          let _2964_v57;
          _2964_v57 = _dafny.Seq.of((_this).f6);
          let _2965_v58;
          _2965_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2898_v0,_2964_v57);
          let _2966_v59;
          let _nw508 = new _module.C7();
          _nw508.__ctor(_2965_v58, false);
          _2966_v59 = _nw508;
          let _2967_v60;
          _2967_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2966_v59);
          let _2968_v61;
          _2968_v61 = _module.D12.create_DC30((_this).f6, (_2943_v37).f9, _module.__default.fm21((((_2963_v56).contains(_2962_v55)) ? ((_2963_v56).get(_2962_v55)) : ((_2943_v37).fm8((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))], globalState))), globalState), new BigNumber((_2967_v60).length));
          let _2969_v62;
          _2969_v62 = _dafny.Seq.of(_2968_v61, _2968_v61);
          let _2970_v63;
          _2970_v63 = _dafny.Seq.of(_2969_v62);
          _2970_v63 = _2970_v63;
          r2 = ((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]).isLessThanOrEqualTo((_2950_v43)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_2950_v43).length))]);
        } else {
          let _2971_v64;
          let _nw509 = Array((new BigNumber(22)).toNumber());
          _2971_v64 = _nw509;
          let _2972_v65;
          let _nw510 = new _module.C2();
          _nw510.__ctor(_2933_v27, _2934_v28, (_2940_v34)[_module.__default.safeIndex((_this).f6, new BigNumber((_2940_v34).length))]);
          _2972_v65 = _nw510;
          let _index464 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_2971_v64).length));
          (_2971_v64)[_index464] = _2972_v65;
          let _index465 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_2971_v64).length));
          let _nw511 = new _module.C2();
          _nw511.__ctor(_2933_v27, _2934_v28, _2902_cf8);
          (_2971_v64)[_index465] = _nw511;
          let _2973_v66;
          let _init95 = function (_2974_i4) {
            return _module.__default.safeModuloInt(_2974_i4, (_this).f6);
          };
          let _nw512 = Array((new BigNumber(29)).toNumber());
          for (let _i0_95 = 0; _i0_95 < new BigNumber(_nw512.length); _i0_95++) {
            _nw512[_i0_95] = _init95(new BigNumber(_i0_95));
          }
          _2973_v66 = _nw512;
          let _2975_v67;
          _2975_v67 = _dafny.Seq.UnicodeFromString("xsi");
          let _index466 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2973_v66).length));
          (_2973_v66)[_index466] = (_dafny.ZERO).minus(new BigNumber((_2975_v67).length));
          let _2976_v68;
          _2976_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2975_v67,(_2943_v37).f9);
          let _2977_v69;
          _2977_v69 = _dafny.MultiSet.fromElements((new BigNumber(453)).minus((_this).f6), (_dafny.ZERO).minus((_this).f6), new BigNumber((_2976_v68).length));
          let _index467 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2973_v66).length));
          let _rhs455 = (_this).f6;
          let _rhs456 = new BigNumber(923);
          let _rhs457 = _2977_v69;
          let _lhs269 = _2973_v66;
          let _lhs270 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2973_v66).length));
          r1 = _rhs455;
          _lhs269[_lhs270] = _rhs456;
          _2977_v69 = _rhs457;
          let _2978_v70;
          _2978_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
          let _2979_v71;
          _2979_v71 = _module.D6.create_DC15(_2934_v28);
          let _2980_v72;
          _2980_v72 = _module.D6.create_DC18(_2979_v71);
          let _2981_v73;
          _2981_v73 = _dafny.Seq.of(_2980_v72);
          let _2982_v74;
          _2982_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2981_v73,_2903_cf7);
          _2978_v70 = (_2978_v70).update((_this).f6, new BigNumber((_2982_v74).length));
          let _2983_v75;
          _2983_v75 = _dafny.Map.Empty.slice().updateUnsafe((_2943_v37).f9,(_2973_v66)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_2973_v66).length))]);
          _2983_v75 = _2983_v75;
          let _2984_v76;
          let _nw513 = new _module.C4();
          _nw513.__ctor(_2945_v39, (_2943_v37).f9);
          _2984_v76 = _nw513;
        }
      } else if (_source21.is_DC5) {
        let _2985___mcc_h2 = (_source21).cf9;
        let _2986___mcc_h3 = (_source21).cf10;
        let _2987___mcc_h4 = (_source21).cf11;
        let _2988_cf11 = _2987___mcc_h4;
        let _2989_cf10 = _2986___mcc_h3;
        let _2990_cf9 = _2985___mcc_h2;
        let _2991_v77;
        let _nw514 = new _module.C6();
        _nw514.__ctor(_2898_v0, _2898_v0);
        _2991_v77 = _nw514;
        let _2992_v78;
        _2992_v78 = _dafny.Set.fromElements(_2991_v77, _2991_v77);
        let _2993_v79;
        _2993_v79 = _dafny.Set.fromElements(new BigNumber((_2992_v78).length));
        let _2994_v80;
        _2994_v80 = _dafny.Seq.UnicodeFromString("sraop");
        let _2995_v81;
        _2995_v81 = _module.D2.create_DC7(_2993_v79);
        let _2996_v82;
        let _nw515 = Array((new BigNumber(19)).toNumber());
        _nw515[(_dafny.ZERO).toNumber()] = _2993_v79;
        _nw515[(_dafny.ONE).toNumber()] = (_2993_v79).Union(_2993_v79);
        _nw515[(new BigNumber(2)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(3)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(4)).toNumber()] = (_2993_v79).Union(_dafny.Set.fromElements(new BigNumber((_2994_v80).length)));
        _nw515[(new BigNumber(5)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(6)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(7)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(8)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(9)).toNumber()] = (_2995_v81).dtor_cf13;
        _nw515[(new BigNumber(10)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(11)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements((_this).f6);
        _nw515[(new BigNumber(13)).toNumber()] = (_2993_v79).Difference(_2993_v79);
        _nw515[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements((_this).f6, new BigNumber(-181));
        _nw515[(new BigNumber(15)).toNumber()] = _module.__default.fm2((_this).f6, _2990_cf9, globalState);
        _nw515[(new BigNumber(16)).toNumber()] = _2993_v79;
        _nw515[(new BigNumber(17)).toNumber()] = (_module.__default.fm2(new BigNumber((_2994_v80).length), false, globalState)).Union(_2993_v79);
        _nw515[(new BigNumber(18)).toNumber()] = _2993_v79;
        _2996_v82 = _nw515;
        _2996_v82 = _2996_v82;
        let _2997_v83;
        _2997_v83 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2998_v84;
        _2998_v84 = _dafny.Set.fromElements(_2997_v83);
        let _2999_v85;
        _2999_v85 = _module.D0.create_DC0(_2988_cf11);
        let _3000_v86;
        _3000_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).fm6(_2999_v85, globalState)).length),_2994_v80);
        let _3001_v87;
        _3001_v87 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2988_cf11));
        let _3002_v88;
        let _nw516 = Array((new BigNumber(26)).toNumber());
        _nw516[(_dafny.ZERO).toNumber()] = new BigNumber(-691);
        _nw516[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw516[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_2988_cf11, (_this).f6);
        _nw516[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).f6, new BigNumber(-107));
        _nw516[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2994_v80, _module.__default.safeIndex((_this).f6, new BigNumber((_2994_v80).length)), _2997_v83), _2994_v80)).length);
        _nw516[(new BigNumber(5)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(6)).toNumber()] = (new BigNumber(76)).minus((_this).f6);
        _nw516[(new BigNumber(7)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(8)).toNumber()] = (_this).f6;
        _nw516[(new BigNumber(9)).toNumber()] = new BigNumber((_2993_v79).length);
        _nw516[(new BigNumber(10)).toNumber()] = new BigNumber(((_2998_v84).Union(_2998_v84)).length);
        _nw516[(new BigNumber(11)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(12)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements(_2994_v80, _dafny.Seq.of(_2997_v83, _2997_v83))).update(_2994_v80, _module.__default.abs((_this).f6))).cardinality()));
        _nw516[(new BigNumber(14)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(15)).toNumber()] = _2988_cf11;
        _nw516[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2994_v80, (((_3000_v86).contains((_this).f6)) ? ((_3000_v86).get((_this).f6)) : (_2994_v80)))).length);
        _nw516[(new BigNumber(17)).toNumber()] = (new BigNumber((_module.__default.fm21(_2990_cf9, globalState)).length)).plus((_this).f6);
        _nw516[(new BigNumber(18)).toNumber()] = new BigNumber((_2994_v80).length);
        _nw516[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw516[(new BigNumber(20)).toNumber()] = (_this).f6;
        _nw516[(new BigNumber(21)).toNumber()] = (_this).f6;
        _nw516[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2994_v80).length), (_dafny.ZERO).minus(_2988_cf11));
        _nw516[(new BigNumber(23)).toNumber()] = new BigNumber(((_3001_v87).Intersect(_3001_v87)).cardinality());
        _nw516[(new BigNumber(24)).toNumber()] = new BigNumber(957);
        _nw516[(new BigNumber(25)).toNumber()] = _2988_cf11;
        _3002_v88 = _nw516;
        let _index468 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_3002_v88).length));
        (_3002_v88)[_index468] = (_dafny.ZERO).minus((_this).f6);
        let _index469 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_3002_v88).length));
        (_3002_v88)[_index469] = (_this).f6;
        let _3003_v89;
        _3003_v89 = _dafny.Seq.of(!(!(_2898_v0)), _2990_cf9);
        _3003_v89 = _3003_v89;
        let _3004_v90;
        let _init96 = ((_3005_v81) => function (_3006_i5) {
          return _3005_v81;
        })(_2995_v81);
        let _nw517 = Array((new BigNumber(8)).toNumber());
        for (let _i0_96 = 0; _i0_96 < new BigNumber(_nw517.length); _i0_96++) {
          _nw517[_i0_96] = _init96(new BigNumber(_i0_96));
        }
        _3004_v90 = _nw517;
        let _3007_v91;
        let _nw518 = new _module.C4();
        _nw518.__ctor(_3004_v90, _2898_v0);
        _3007_v91 = _nw518;
        let _3008_v92;
        _3008_v92 = _dafny.Set.fromElements(_3007_v91);
        let _3009_v93;
        _3009_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2898_v0,_3008_v92);
        let _3010_v94;
        _3010_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        let _3011_v95;
        _3011_v95 = _dafny.MultiSet.fromElements(_3010_v94, _3010_v94);
        let _3012_v96;
        _3012_v96 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-743),(_3002_v88)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_3002_v88).length))]));
        let _3013_v97;
        _3013_v97 = _dafny.Seq.of(_3001_v87, (_3001_v87).update(new BigNumber((_2994_v80).length), _module.__default.abs(new BigNumber((_dafny.Seq.of(!(_2990_cf9))).length))));
        _3009_v93 = (_3009_v93).update(!(((_dafny.MultiSet.FromArray(_3012_v96)).update(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3002_v88,_2991_v77.f7)).length),(_3002_v88)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_3002_v88).length))]), _module.__default.abs(new BigNumber(((_3013_v97)[_module.__default.safeIndex(new BigNumber(271), new BigNumber((_3013_v97).length))]).cardinality())))).IsProperSubsetOf(_3011_v95)), _dafny.Set.fromElements(_3007_v91, _3007_v91));
      } else if (_source21.is_DC3) {
        let _3014___mcc_h5 = (_source21).cf6;
        let _3015_cf6 = _3014___mcc_h5;
        let _3016_v98;
        let _nw519 = Array((new BigNumber(29)).toNumber());
        _3016_v98 = _nw519;
        let _3017_v99;
        let _nw520 = new _module.C6();
        _nw520.__ctor(_2898_v0, _2898_v0);
        _3017_v99 = _nw520;
        let _index470 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_3016_v98).length));
        (_3016_v98)[_index470] = _3017_v99;
        let _index471 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_3016_v98).length));
        (_3016_v98)[_index471] = _3017_v99;
        let _3018_v100;
        let _nw521 = new _module.C1();
        _nw521.__ctor(_2898_v0);
        _3018_v100 = _nw521;
        let _3019_v101;
        _3019_v101 = _dafny.MultiSet.fromElements(_3018_v100);
        let _3020_v102;
        let _nw522 = new _module.C6();
        _nw522.__ctor(!((((_3017_v99).f9) ? (_2898_v0) : (_2898_v0))), (_3019_v101).IsProperSubsetOf(_dafny.MultiSet.fromElements(_3018_v100, _3018_v100)));
        _3020_v102 = _nw522;
        let _3021_v103;
        let _nw523 = new _module.C5();
        _nw523.__ctor((_3017_v99).f9);
        _3021_v103 = _nw523;
        r2 = (_3020_v102).fm8((_this).f6, globalState);
      } else {
        let _3022___mcc_h6 = (_source21).cf12;
        let _3023_cf12 = _3022___mcc_h6;
        let _3024_v104;
        _3024_v104 = _dafny.Map.Empty.slice().updateUnsafe(_2898_v0,_module.__default.fm21(!(false), globalState));
        let _3025_v105;
        let _nw524 = new _module.C7();
        _nw524.__ctor(_3024_v104, _2898_v0);
        _3025_v105 = _nw524;
        let _3026_v106;
        _3026_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3025_v105);
        let _3027_v107;
        _3027_v107 = _dafny.Map.Empty.slice().updateUnsafe(_3025_v105.f7,(_this).f6);
        let _3028_v108;
        _3028_v108 = _dafny.Seq.of((_3026_v106).Merge((_3026_v106).update((((_3027_v107).contains(_3025_v105.f7)) ? ((_3027_v107).get(_3025_v105.f7)) : ((_this).f6)), _3025_v105)), _3026_v106, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_3025_v105.f7)).length),_3025_v105)).Merge(_3026_v106), _3026_v106);
        _3028_v108 = _3028_v108;
        let _3029_v109;
        let _nw525 = Array((new BigNumber(25)).toNumber()).fill(_module.D13.Default());
        _3029_v109 = _nw525;
        let _3030_v110;
        let _nw526 = Array((new BigNumber(16)).toNumber()).fill([]);
        _3030_v110 = _nw526;
        let _3031_v111;
        _3031_v111 = _module.D0.create_DC0((_this).f6);
        let _3032_v112;
        _3032_v112 = _dafny.Seq.of(_3031_v111);
        let _3033_v113;
        _3033_v113 = _module.D12.create_DC33(_module.D12.create_DC30(new BigNumber(171), _module.__default.fm1(_3032_v112, (_this).f6, (_this).f6, globalState), _dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), function (_3034_i6) {
  return (_this).f6;
}), (_this).f6));
        let _3035_v114;
        _3035_v114 = _module.D12.create_DC32(false);
        let _pat_let_tv78 = _3035_v114;
        let _3036_v115;
        let _nw527 = Array((new BigNumber(15)).toNumber());
        _nw527[(_dafny.ZERO).toNumber()] = _3033_v113;
        _nw527[(_dafny.ONE).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(2)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(3)).toNumber()] = _module.D12.create_DC33(_3035_v114);
        _nw527[(new BigNumber(4)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(5)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(6)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(7)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(8)).toNumber()] = function (_pat_let87_0) {
          return function (_3037_dt__update__tmp_h0) {
            return function (_pat_let88_0) {
              return function (_3038_dt__update_hcf56_h0) {
                return _module.D12.create_DC33(_3038_dt__update_hcf56_h0);
              }(_pat_let88_0);
            }(_pat_let_tv78);
          }(_pat_let87_0);
        }(_3033_v113);
        _nw527[(new BigNumber(9)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(10)).toNumber()] = _module.D12.create_DC33(_3035_v114);
        _nw527[(new BigNumber(11)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(12)).toNumber()] = _module.D12.create_DC33(_3035_v114);
        _nw527[(new BigNumber(13)).toNumber()] = _3033_v113;
        _nw527[(new BigNumber(14)).toNumber()] = _3033_v113;
        _3036_v115 = _nw527;
        let _index472 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_3030_v110).length));
        (_3030_v110)[_index472] = _3036_v115;
        let _3039_v116;
        _3039_v116 = _module.D21.create_DC57(_3036_v115);
        let _index473 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_3030_v110).length));
        let _rhs458 = false;
        let _rhs459 = _3029_v109;
        let _rhs460 = (_3039_v116).dtor_cf98;
        let _lhs271 = _3030_v110;
        let _lhs272 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_3030_v110).length));
        _2898_v0 = _rhs458;
        _3029_v109 = _rhs459;
        _lhs271[_lhs272] = _rhs460;
        r3 = (_this).f6;
        let _3040_v117;
        _3040_v117 = _dafny.Seq.UnicodeFromString("ao");
        let _3041_v118;
        _3041_v118 = _module.D2.create_DC8(((_this).f6).isLessThan((_this).f6), _3040_v117);
        let _source22 = _3041_v118;
        if (_source22.is_DC8) {
          let _3042___mcc_h7 = (_source22).cf14;
          let _3043___mcc_h8 = (_source22).cf15;
          let _3044_cf15 = _3043___mcc_h8;
          let _3045_cf14 = _3042___mcc_h7;
          let _3046_v119;
          _3046_v119 = _dafny.Set.fromElements((_this).f6, (_this).f6);
          let _3047_v120;
          _3047_v120 = _module.D2.create_DC7(_3046_v119);
          let _3048_v121;
          let _nw528 = Array((new BigNumber(18)).toNumber());
          _nw528[(_dafny.ZERO).toNumber()] = _3047_v120;
          _nw528[(_dafny.ONE).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(2)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(3)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(4)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(5)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(6)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(7)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(8)).toNumber()] = _module.D2.create_DC7(_3046_v119);
          _nw528[(new BigNumber(9)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(10)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(11)).toNumber()] = _module.D2.create_DC7(_3046_v119);
          _nw528[(new BigNumber(12)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(13)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(14)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(15)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(16)).toNumber()] = _3047_v120;
          _nw528[(new BigNumber(17)).toNumber()] = _3047_v120;
          _3048_v121 = _nw528;
          let _3049_v122;
          _3049_v122 = _module.D22.create_DC61(_3048_v121);
          let _pat_let_tv79 = _3048_v121;
          let _3050_v123;
          _3050_v123 = _dafny.Seq.of((function (_pat_let89_0) {
            return function (_3051_dt__update__tmp_h1) {
              return function (_pat_let90_0) {
                return function (_3052_dt__update_hcf105_h0) {
                  return _module.D22.create_DC61(_3052_dt__update_hcf105_h0);
                }(_pat_let90_0);
              }(_pat_let_tv79);
            }(_pat_let89_0);
          }(_3049_v122)).dtor_cf105);
          let _3053_v124;
          let _nw529 = new _module.C4();
          _nw529.__ctor((_3050_v123)[_module.__default.safeIndex((_this).f6, new BigNumber((_3050_v123).length))], !(false) || (_2898_v0));
          _3053_v124 = _nw529;
          (_3025_v105).f7 = _2898_v0;
          _3045_cf14 = (!(_2898_v0)) || (_3045_cf14);
          let _3054_v125;
          _3054_v125 = _module.D15.create_DC43(new BigNumber((_dafny.Seq.UnicodeFromString("xnka")).length));
          let _3055_v126;
          _3055_v126 = _module.D15.create_DC44(_3054_v125);
          let _3056_v127;
          _3056_v127 = _dafny.Seq.of(_3055_v126, _3055_v126, _3055_v126, _3055_v126, _3055_v126);
          _3056_v127 = _module.__default.fm59(globalState);
        } else if (_source22.is_DC9) {
          let _3057___mcc_h9 = (_source22).cf16;
          let _3058___mcc_h10 = (_source22).cf17;
          let _3059___mcc_h11 = (_source22).cf18;
          let _3060___mcc_h12 = (_source22).cf19;
          let _3061___mcc_h13 = (_source22).cf20;
          let _3062_cf20 = _3061___mcc_h13;
          let _3063_cf19 = _3060___mcc_h12;
          let _3064_cf18 = _3059___mcc_h11;
          let _3065_cf17 = _3058___mcc_h10;
          let _3066_cf16 = _3057___mcc_h9;
          _3065_cf17 = (_3066_cf16).plus((_this).f6);
          let _3067_v128;
          _3067_v128 = _dafny.Map.Empty.slice().updateUnsafe(_3063_cf19,_3039_v116);
          let _3068_v129;
          _3068_v129 = _dafny.Seq.of(_3067_v128);
          _3068_v129 = ((_2898_v0) ? (_dafny.Seq.of((_3067_v128).update(_2898_v0, _3039_v116), _3067_v128, _3067_v128)) : (_3068_v129));
          let _3069_v130;
          _3069_v130 = new _dafny.CodePoint('l'.codePointAt(0));
          _3069_v130 = _3069_v130;
          let _3070_v131;
          _3070_v131 = _dafny.Map.Empty.slice().updateUnsafe(_3066_cf16,_3025_v105.f7);
          let _3071_v132;
          _3071_v132 = _module.D23.create_DC64(_dafny.Map.Empty.slice().updateUnsafe(_3066_cf16,_3069_v130));
          let _3072_v133;
          _3072_v133 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3069_v130);
          let _3073_v135;
          _3073_v135 = _dafny.Seq.of(_3065_cf17);
          let _rhs461 = ((_3063_cf19) ? (_3025_v105.f7) : (!((_3071_v132).dtor_cf109).equals(_3072_v133)));
          let _rhs462 = _3065_cf17;
          let _rhs463 = (function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of (_3073_v135).Elements) {
              let _3074_v134 = _compr_56;
              if (_dafny.Seq.contains(_3073_v135, _3074_v134)) {
                _coll56.push([_module.__default.safeDivisionInt(_3074_v134, _3066_cf16),!(_3025_v105.f7)]);
              }
            }
            return _coll56;
          }()).Merge((_3070_v131).Merge(_module.__default.fm45(_3065_cf17, (_this).f6, globalState)));
          r0 = _rhs461;
          _3066_cf16 = _rhs462;
          _3070_v131 = _rhs463;
        } else {
          let _3075___mcc_h14 = (_source22).cf13;
          let _3076_cf13 = _3075___mcc_h14;
          let _3077_v136;
          _3077_v136 = _dafny.Set.fromElements(_2898_v0);
          r3 = new BigNumber(((_module.__default.fm60(globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_3077_v136,new BigNumber(315)))).length);
          let _3078_v137;
          let _nw530 = Array((new BigNumber(28)).toNumber()).fill(false);
          _3078_v137 = _nw530;
          let _3079_v138;
          let _nw531 = Array((new BigNumber(9)).toNumber());
          _nw531[(_dafny.ZERO).toNumber()] = _3078_v137;
          _nw531[(_dafny.ONE).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(2)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(3)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(4)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(5)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(6)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(7)).toNumber()] = _3078_v137;
          _nw531[(new BigNumber(8)).toNumber()] = _3078_v137;
          _3079_v138 = _nw531;
          let _index474 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_3079_v138).length));
          (_3079_v138)[_index474] = _3078_v137;
          let _3080_v139;
          _3080_v139 = _dafny.Seq.of(_3025_v105.f7, !(!(_2898_v0)), true);
          let _index475 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_3078_v137).length));
          (_3078_v137)[_index475] = !_dafny.Seq.contains(_3080_v139, _2898_v0);
          let _index476 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_3078_v137).length));
          (_3078_v137)[_index476] = _3025_v105.f7;
          let _3081_v140;
          let _init97 = ((_3082_v117) => function (_3083_i7) {
            return _module.__default.safeDivisionInt(_3083_i7, new BigNumber((_3082_v117).length));
          })(_3040_v117);
          let _nw532 = Array((new BigNumber(29)).toNumber());
          for (let _i0_97 = 0; _i0_97 < new BigNumber(_nw532.length); _i0_97++) {
            _nw532[_i0_97] = _init97(new BigNumber(_i0_97));
          }
          _3081_v140 = _nw532;
          let _index477 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_3081_v140).length));
          (_3081_v140)[_index477] = new BigNumber(13);
          let _3084_v141;
          _3084_v141 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3025_v105.f7);
          let _index478 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_3079_v138).length));
          let _index479 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_3078_v137).length));
          let _index480 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_3078_v137).length));
          let _index481 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_3081_v140).length));
          let _rhs464 = _3078_v137;
          let _rhs465 = (_this).f6;
          let _rhs466 = _2898_v0;
          let _rhs467 = (((_3084_v141).contains((_this).f6)) ? ((_3084_v141).get((_this).f6)) : (_2898_v0));
          let _rhs468 = (_dafny.ZERO).minus((_this).f6);
          let _lhs273 = _3079_v138;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_3079_v138).length));
          let _lhs275 = _3078_v137;
          let _lhs276 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_3078_v137).length));
          let _lhs277 = _3078_v137;
          let _lhs278 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_3078_v137).length));
          let _lhs279 = _3081_v140;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_3081_v140).length));
          _lhs273[_lhs274] = _rhs464;
          r1 = _rhs465;
          _lhs275[_lhs276] = _rhs466;
          _lhs277[_lhs278] = _rhs467;
          _lhs279[_lhs280] = _rhs468;
          let _3085_v142;
          let _nw533 = new _module.C1();
          _nw533.__ctor(!((_dafny.Map.Empty.slice().updateUnsafe(_3025_v105.f7,_3025_v105.f7))).equals(_dafny.Map.Empty.slice().updateUnsafe(_3025_v105.f7,_3025_v105.f7)));
          _3085_v142 = _nw533;
          _3085_v142 = _3085_v142;
        }
      }
      let _3086_v143;
      _3086_v143 = _dafny.Seq.of((_this).f6, new BigNumber(-167));
      let _3087_v144;
      _3087_v144 = _dafny.Seq.UnicodeFromString("cosrqbc");
      let _3088_v145;
      _3088_v145 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _3089_v146;
      _3089_v146 = _dafny.MultiSet.fromElements(_2898_v0, _2898_v0, _2898_v0);
      let _3090_v147;
      let _nw534 = Array((new BigNumber(23)).toNumber()).fill(false);
      _3090_v147 = _nw534;
      let _3091_v148;
      _3091_v148 = _dafny.Set.fromElements(_3090_v147, _3090_v147, _3090_v147, _3090_v147);
      let _3092_v149;
      _3092_v149 = new _dafny.CodePoint('l'.codePointAt(0));
      let _3093_v150;
      let _nw535 = new _module.C2();
      _nw535.__ctor(_3091_v148, _3092_v149, _2898_v0);
      _3093_v150 = _nw535;
      let _3094_v151;
      _3094_v151 = _dafny.MultiSet.fromElements(_3093_v150, _3093_v150);
      let _3095_v152;
      _3095_v152 = _module.D0.create_DC0(_module.__default.fm0(globalState));
      let _3096_v153;
      let _nw536 = Array((new BigNumber(24)).toNumber());
      _nw536[(_dafny.ZERO).toNumber()] = new BigNumber((_3086_v143).length);
      _nw536[(_dafny.ONE).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(2)).toNumber()] = (_this).fm5(globalState);
      _nw536[(new BigNumber(3)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(4)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(5)).toNumber()] = (new BigNumber((_3087_v144).length)).multipliedBy((((_3088_v145).contains((_this).f6)) ? ((_3088_v145).get((_this).f6)) : ((_this).f6)));
      _nw536[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((_this).f6)).plus((_this).f6);
      _nw536[(new BigNumber(7)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(8)).toNumber()] = new BigNumber((_3089_v146).cardinality());
      _nw536[(new BigNumber(9)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f6), (_this).f6);
      _nw536[(new BigNumber(11)).toNumber()] = ((_this).f6).multipliedBy((_this).f6);
      _nw536[(new BigNumber(12)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(13)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(14)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(15)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(16)).toNumber()] = (((_3094_v151).contains(_3093_v150)) ? ((_3094_v151).get(_3093_v150)) : ((_this).f6));
      _nw536[(new BigNumber(17)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(18)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(19)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(20)).toNumber()] = new BigNumber(176);
      _nw536[(new BigNumber(21)).toNumber()] = (_this).f6;
      _nw536[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_this).fm6(function (_pat_let91_0) {
        return function (_3097_dt__update__tmp_h2) {
          return function (_pat_let92_0) {
            return function (_3098_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_3098_dt__update_hcf0_h0);
            }(_pat_let92_0);
          }((_this).f6);
        }(_pat_let91_0);
      }(_3095_v152), globalState), _3087_v144)).length);
      _nw536[(new BigNumber(23)).toNumber()] = new BigNumber(-750);
      _3096_v153 = _nw536;
      let _3099_v154;
      _3099_v154 = _dafny.Map.Empty.slice().updateUnsafe(_3096_v153,(_dafny.ZERO).minus((_this).f6));
      let _index482 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
      (_3096_v153)[_index482] = _module.__default.safeDivisionInt((((_3099_v154).contains(_3096_v153)) ? ((_3099_v154).get(_3096_v153)) : (new BigNumber(666))), (_this).f6);
      let _index483 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
      (_3096_v153)[_index483] = _module.__default.safeDivisionInt((_this).f6, (_this).f6);
      r3 = new BigNumber(435);
      let _3100_v155;
      _3100_v155 = _module.D2.create_DC8(_2898_v0, _3087_v144);
      let _hi19 = new BigNumber(-443);
      for (let _3101_i8 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3100_v155,_3087_v144)).length); _3101_i8.isLessThan(_hi19); _3101_i8 = _3101_i8.plus(_dafny.ONE)) {
        let _3102_v156;
        _3102_v156 = _dafny.Seq.of(_2898_v0, _2898_v0);
        let _3103_v157;
        _3103_v157 = _module.D13.create_DC36(_3092_v149, _3102_v156, _3093_v150.f7);
        let _3104_v158;
        _3104_v158 = _module.D13.create_DC37(_3103_v157);
        let _source23 = _3104_v158;
        if (_source23.is_DC35) {
          _module.__default.m0((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], _module.__default.safeDivisionInt((_this).f6, (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]), globalState);
          let _3105_v159;
          _3105_v159 = _dafny.Seq.of(_3095_v152, _3095_v152);
          let _pat_let_tv80 = _3096_v153;
          let _pat_let_tv81 = _3096_v153;
          r0 = ((_module.__default.fm1(_dafny.Seq.update(_3105_v159, _module.__default.safeIndex(_3101_i8, new BigNumber((_3105_v159).length)), function (_pat_let93_0) {
            return function (_3106_dt__update__tmp_h3) {
              return function (_pat_let94_0) {
                return function (_3107_dt__update_hcf0_h1) {
                  return _module.D0.create_DC0(_3107_dt__update_hcf0_h1);
                }(_pat_let94_0);
              }((_pat_let_tv81)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_pat_let_tv80).length))]);
            }(_pat_let93_0);
          }(_3095_v152)), new BigNumber((_3102_v156).length), (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], globalState)) ? (_2898_v0) : (false));
          let _3108_v160;
          _3108_v160 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_3086_v143, _3086_v143),_3093_v150);
          _3108_v160 = (_3108_v160).update(_3086_v143, _3093_v150);
          let _3109_v161;
          _3109_v161 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(205),_3092_v149);
          let _3110_v162;
          _3110_v162 = _dafny.Set.fromElements((((_3109_v161).contains(_3101_i8)) ? ((_3109_v161).get(_3101_i8)) : (new _dafny.CodePoint('a'.codePointAt(0)))), _3092_v149);
          let _3111_v164;
          _3111_v164 = _dafny.Seq.of(function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-994)), ((_3112_v144) => function (_3113_i9) {
              return _3112_v144;
            })(_3087_v144))).Elements) {
              let _3114_v163 = _compr_57;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-994)), ((_3115_v144) => function (_3113_i9) {
                return _3115_v144;
              })(_3087_v144)), _3114_v163)) {
                _coll57.add(_3114_v163);
              }
            }
            return _coll57;
          }());
          let _3116_v165;
          _3116_v165 = _dafny.Map.Empty.slice().updateUnsafe(_3093_v150.f7,_3093_v150.f7);
          _3102_v156 = _dafny.Seq.Concat(_3102_v156, _module.__default.fm50(_3110_v162, (_3111_v164)[_module.__default.safeIndex(_3101_i8, new BigNumber((_3111_v164).length))], new BigNumber((_3116_v165).length), globalState));
        } else if (_source23.is_DC36) {
          let _3117___mcc_h15 = (_source23).cf58;
          let _3118___mcc_h16 = (_source23).cf59;
          let _3119___mcc_h17 = (_source23).cf60;
          let _3120_cf60 = _3119___mcc_h17;
          let _3121_cf59 = _3118___mcc_h16;
          let _3122_cf58 = _3117___mcc_h15;
          let _index484 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index484] = _2898_v0;
          let _index485 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index485] = _2898_v0;
          r3 = _module.__default.safeDivisionInt((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _index486 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          (_3096_v153)[_index486] = ((_this).f6).plus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          r3 = _3101_i8;
        } else if (_source23.is_DC34) {
          let _3123___mcc_h18 = (_source23).cf57;
          let _3124_cf57 = _3123___mcc_h18;
          let _3125_v166;
          let _nw537 = new _module.C6();
          _nw537.__ctor(_3093_v150.f7, _3093_v150.f7);
          _3125_v166 = _nw537;
          let _3126_v167;
          let _nw538 = Array((new BigNumber(12)).toNumber());
          _nw538[(_dafny.ZERO).toNumber()] = _3125_v166;
          _nw538[(_dafny.ONE).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(2)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(3)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(4)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(5)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(6)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(7)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(8)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(9)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(10)).toNumber()] = _3125_v166;
          _nw538[(new BigNumber(11)).toNumber()] = _3125_v166;
          _3126_v167 = _nw538;
          let _index487 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_3126_v167).length));
          (_3126_v167)[_index487] = _3125_v166;
          let _3127_v168;
          let _nw539 = Array((new BigNumber(18)).toNumber());
          _nw539[(_dafny.ZERO).toNumber()] = _3087_v144;
          _nw539[(_dafny.ONE).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(2)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(3)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(4)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("oviiryq");
          _nw539[(new BigNumber(6)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(7)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(8)).toNumber()] = ((!(false)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), ((_3128_v149) => function (_3129_i10) {
            return _3128_v149;
          })(_3092_v149))) : (_3087_v144));
          _nw539[(new BigNumber(9)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("bwnc");
          _nw539[(new BigNumber(11)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(12)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(13)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_3087_v144, _module.__default.safeIndex(_3101_i8, new BigNumber((_3087_v144).length)), _3092_v149);
          _nw539[(new BigNumber(15)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(16)).toNumber()] = _3087_v144;
          _nw539[(new BigNumber(17)).toNumber()] = _3087_v144;
          _3127_v168 = _nw539;
          let _index488 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_3127_v168).length));
          (_3127_v168)[_index488] = _3087_v144;
          let _3130_v169;
          _3130_v169 = _dafny.Seq.of(_3125_v166, _3125_v166);
          let _3131_v170;
          _3131_v170 = _dafny.Set.fromElements(_3092_v149, _3092_v149, _3092_v149, _3092_v149, _3092_v149);
          let _3132_v171;
          _3132_v171 = _3131_v170;
          let _3133_v172;
          _3133_v172 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3132_v171);
          let _3134_v173;
          _3134_v173 = _dafny.Map.Empty.slice().updateUnsafe(_3133_v172,_module.__default.fm26(!(_3093_v150.f7), (_this).f6, new BigNumber((_3102_v156).length), globalState));
          let _3135_v174;
          _3135_v174 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(globalState),_3133_v172);
          let _3136_v175;
          _3136_v175 = _module.D21.create_DC58((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _3137_v176;
          let _nw540 = Array((new BigNumber(22)).toNumber()).fill(_module.D2.Default());
          _3137_v176 = _nw540;
          let _3138_v177;
          let _nw541 = new _module.C4();
          _nw541.__ctor(_3137_v176, _2898_v0);
          _3138_v177 = _nw541;
          let _3139_v178;
          _3139_v178 = _dafny.Map.Empty.slice().updateUnsafe(_3125_v166.f7,_3138_v177);
          let _index489 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_3126_v167).length));
          let _index490 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_3127_v168).length));
          let _index491 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          let _rhs469 = (_3130_v169)[_module.__default.safeIndex((_this).f6, new BigNumber((_3130_v169).length))];
          let _rhs470 = (((_3134_v173).contains((((_3135_v174).contains((_3136_v175).dtor_cf99)) ? ((_3135_v174).get((_3136_v175).dtor_cf99)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3139_v178)).length),_3132_v171))))) ? ((_3134_v173).get((((_3135_v174).contains((_3136_v175).dtor_cf99)) ? ((_3135_v174).get((_3136_v175).dtor_cf99)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_3139_v178)).length),_3132_v171))))) : (_module.__default.fm13(_3093_v150.f7, new BigNumber((_3087_v144).length), _2898_v0, globalState)));
          let _rhs471 = _3087_v144;
          let _rhs472 = ((_dafny.ZERO).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))])).isEqualTo((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _rhs473 = ((_this).f6).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _lhs281 = _3126_v167;
          let _lhs282 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_3126_v167).length));
          let _lhs283 = _3127_v168;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_3127_v168).length));
          let _lhs285 = _3096_v153;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          _lhs281[_lhs282] = _rhs469;
          _3092_v149 = _rhs470;
          _lhs283[_lhs284] = _rhs471;
          r2 = _rhs472;
          _lhs285[_lhs286] = _rhs473;
          _3093_v150 = _3093_v150;
          _2899_v1 = _2899_v1;
          r3 = (_this).f6;
        } else {
          let _3140___mcc_h19 = (_source23).cf61;
          let _3141_cf61 = _3140___mcc_h19;
          let _3142_v179;
          let _nw542 = Array((new BigNumber(5)).toNumber());
          _3142_v179 = _nw542;
          let _3143_v180;
          let _nw543 = Array((new BigNumber(10)).toNumber()).fill([]);
          _3143_v180 = _nw543;
          let _index492 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_3143_v180).length));
          (_3143_v180)[_index492] = _3096_v153;
          let _index493 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index493] = !(new BigNumber(386)).isEqualTo((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _index494 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_3143_v180).length));
          let _index495 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_3090_v147).length));
          let _rhs474 = _3142_v179;
          let _rhs475 = (_3087_v144)[_module.__default.safeIndex(new BigNumber((_3087_v144).length), new BigNumber((_3087_v144).length))];
          let _rhs476 = _2898_v0;
          let _rhs477 = _3096_v153;
          let _rhs478 = _3093_v150.f7;
          let _lhs287 = _3143_v180;
          let _lhs288 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_3143_v180).length));
          let _lhs289 = _3090_v147;
          let _lhs290 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_3090_v147).length));
          _3142_v179 = _rhs474;
          _3092_v149 = _rhs475;
          r2 = _rhs476;
          _lhs287[_lhs288] = _rhs477;
          _lhs289[_lhs290] = _rhs478;
          let _3144_v181;
          _3144_v181 = _dafny.Map.Empty.slice().updateUnsafe(false,_3101_i8);
          let _pat_let_tv82 = _3144_v181;
          let _pat_let_tv83 = _3093_v150;
          let _pat_let_tv84 = _3087_v144;
          let _pat_let_tv85 = _3087_v144;
          (_3093_v150).f7 = _module.__default.fm1(_dafny.Seq.of(_3095_v152, _3095_v152, function (_pat_let95_0) {
            return function (_3145_dt__update__tmp_h4) {
              return function (_pat_let96_0) {
                return function (_3146_dt__update_hcf0_h2) {
                  return _module.D0.create_DC0(_3146_dt__update_hcf0_h2);
                }(_pat_let96_0);
              }(new BigNumber(((_pat_let_tv82).update(!(_pat_let_tv83.f7), new BigNumber((_pat_let_tv84).length))).length));
            }(_pat_let95_0);
          }(_3095_v152), function (_pat_let97_0) {
            return function (_3149_dt__update__tmp_h6) {
              return function (_pat_let100_0) {
                return function (_3150_dt__update_hcf0_h4) {
                  return _module.D0.create_DC0(_3150_dt__update_hcf0_h4);
                }(_pat_let100_0);
              }(_3101_i8);
            }(_pat_let97_0);
          }(function (_pat_let98_0) {
            return function (_3147_dt__update__tmp_h5) {
              return function (_pat_let99_0) {
                return function (_3148_dt__update_hcf0_h3) {
                  return _module.D0.create_DC0(_3148_dt__update_hcf0_h3);
                }(_pat_let99_0);
              }(new BigNumber((_pat_let_tv85).length));
            }(_pat_let98_0);
          }(_3095_v152))), (_dafny.ZERO).minus(new BigNumber((_3088_v145).length)), (_this).f6, globalState);
          let _3151_v182;
          let _init98 = function (_3152_i11) {
            return _module.D2.create_DC7(_dafny.Set.fromElements((_this).f6));
          };
          let _nw544 = Array((new BigNumber(8)).toNumber());
          for (let _i0_98 = 0; _i0_98 < new BigNumber(_nw544.length); _i0_98++) {
            _nw544[_i0_98] = _init98(new BigNumber(_i0_98));
          }
          _3151_v182 = _nw544;
          let _3153_v183;
          _3153_v183 = _module.D22.create_DC61(_3151_v182);
          let _3154_v184;
          let _nw545 = new _module.C5();
          _nw545.__ctor((_3090_v147)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_3090_v147).length))]);
          _3154_v184 = _nw545;
          let _3155_v185;
          let _nw546 = new _module.C4();
          _nw546.__ctor(_3151_v182, _3093_v150.f7);
          _3155_v185 = _nw546;
          let _rhs479 = (_3088_v145).update(_3101_i8, _module.__default.safeModuloInt((_dafny.ZERO).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]), (_this).f6));
          let _rhs480 = false;
          let _rhs481 = _module.D22.create_DC61((_3155_v185).f11);
          let _rhs482 = _3154_v184;
          let _rhs483 = _3155_v185;
          let _lhs291 = _3093_v150;
          _3088_v145 = _rhs479;
          _lhs291.f7 = _rhs480;
          _3153_v183 = _rhs481;
          _3154_v184 = _rhs482;
          _3155_v185 = _rhs483;
          let _index496 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          (_3096_v153)[_index496] = (_dafny.ZERO).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
        }
        let _3156_v186;
        _3156_v186 = _dafny.Map.Empty.slice().updateUnsafe(_3093_v150.f7,_3086_v143);
        let _3157_v187;
        let _nw547 = new _module.C7();
        _nw547.__ctor(_3156_v186, !(_3093_v150.f7));
        _3157_v187 = _nw547;
        _3157_v187 = _3157_v187;
        let _index497 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_3090_v147).length));
        (_3090_v147)[_index497] = _3093_v150.f7;
        let _index498 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_3090_v147).length));
        (_3090_v147)[_index498] = ((_this).f6).isLessThan((_this).f6);
        if ((_3157_v187).fm8(((_3093_v150.f7) ? ((_this).f6) : (new BigNumber(391))), globalState)) {
          let _3158_v188;
          _3158_v188 = _dafny.Map.Empty.slice().updateUnsafe(_3101_i8,new _dafny.CodePoint('u'.codePointAt(0)));
          let _3159_v189;
          _3159_v189 = _module.D23.create_DC64((_3158_v188).Merge(_3158_v188));
          let _pat_let_tv86 = _3158_v188;
          let _rhs484 = _dafny.Seq.of(new BigNumber(-903));
          let _rhs485 = function (_pat_let101_0) {
            return function (_3160_dt__update__tmp_h7) {
              return function (_pat_let102_0) {
                return function (_3161_dt__update_hcf109_h0) {
                  return _module.D23.create_DC64(_3161_dt__update_hcf109_h0);
                }(_pat_let102_0);
              }(_pat_let_tv86);
            }(_pat_let101_0);
          }(_3159_v189);
          _3086_v143 = _rhs484;
          _3159_v189 = _rhs485;
          _3086_v143 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_3086_v143, _dafny.Seq.of((_this).f6, (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], (_this).f6, new BigNumber((_3086_v143).length), (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))])), _3086_v143), _module.__default.safeIndex((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_3086_v143, _dafny.Seq.of((_this).f6, (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], (_this).f6, new BigNumber((_3086_v143).length), (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))])), _3086_v143)).length)), (_this).f6);
          let _3162_v190;
          _3162_v190 = _dafny.Map.Empty.slice().updateUnsafe(_3093_v150.f7,_2898_v0);
          let _3163_v191;
          _3163_v191 = _3162_v190;
          let _3164_v192;
          _3164_v192 = _dafny.MultiSet.fromElements(_module.__default.fm21(_3093_v150.f7, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_3165_i12) {
            return new BigNumber(689);
          }));
          let _rhs486 = (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))];
          let _rhs487 = _module.__default.safeDivisionInt(new BigNumber((_3087_v144).length), (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _rhs488 = _3163_v191;
          let _rhs489 = new BigNumber(((_module.__default.fm61((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], false, globalState)).Intersect(_3164_v192)).cardinality());
          let _rhs490 = _3093_v150.f7;
          let _lhs292 = _3093_v150;
          r3 = _rhs486;
          r3 = _rhs487;
          _3163_v191 = _rhs488;
          r1 = _rhs489;
          _lhs292.f7 = _rhs490;
          _module.__default.m0((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], (_dafny.ZERO).minus(_3101_i8), globalState);
          let _3166_v194;
          let _init99 = ((_3167_v188, _3168_v144, _3169_v149) => function (_3170_i13) {
            return (function () {
              let _coll58 = new _dafny.Map();
              for (const _compr_58 of (_dafny.Seq.of(_3168_v144)).Elements) {
                let _3171_v193 = _compr_58;
                if (_dafny.Seq.contains(_dafny.Seq.of(_3168_v144), _3171_v193)) {
                  _coll58.push([_3171_v193,_3169_v149]);
                }
              }
              return _coll58;
            }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-365)), ((_3172_v188, _3173_v149) => function (_3174_i14) {
              return (((_3172_v188).contains((_this).f6)) ? ((_3172_v188).get((_this).f6)) : (_3173_v149));
            })(_3167_v188, _3169_v149)),_3169_v149));
          })(_3158_v188, _3087_v144, _3092_v149);
          let _nw548 = Array((new BigNumber(26)).toNumber());
          for (let _i0_99 = 0; _i0_99 < new BigNumber(_nw548.length); _i0_99++) {
            _nw548[_i0_99] = _init99(new BigNumber(_i0_99));
          }
          _3166_v194 = _nw548;
          let _3175_v195;
          _3175_v195 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("plrvjxso"),_3092_v149);
          let _index499 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_3166_v194).length));
          (_3166_v194)[_index499] = _3175_v195;
          let _3176_v197;
          _3176_v197 = _dafny.MultiSet.fromElements(_3087_v144, _3087_v144);
          let _index500 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_3166_v194).length));
          (_3166_v194)[_index500] = ((!(!(_3093_v150.f7))) ? (_3175_v195) : (function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_3176_v197).Elements) {
              let _3177_v196 = _compr_59;
              if ((_3176_v197).contains(_3177_v196)) {
                _coll59.push([_3177_v196,_3092_v149]);
              }
            }
            return _coll59;
          }()));
        } else {
          r3 = (_3086_v143)[_module.__default.safeIndex(_3101_i8, new BigNumber((_3086_v143).length))];
          let _3178_v198;
          _3178_v198 = _dafny.Seq.of(_module.__default.fm50(_module.__default.fm30((_3090_v147)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_3090_v147).length))], (_this).f6, (_3090_v147)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_3090_v147).length))], globalState), _module.__default.fm62(_3093_v150.f7, (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], globalState), _3101_i8, globalState));
          _3178_v198 = _module.__default.fm41(_3093_v150.f7, globalState);
          r3 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_3101_i8));
          _3088_v145 = (_3088_v145).update((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f6, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_3089_v146).cardinality()))))));
          let _3179_v199;
          let _nw549 = new _module.C5();
          _nw549.__ctor(_2898_v0);
          _3179_v199 = _nw549;
        }
      }
      let _3180_v200;
      _3180_v200 = _dafny.Set.fromElements(_2898_v0, _2898_v0);
      let _3181_v201;
      _3181_v201 = _module.D1.create_DC3(_3180_v200);
      let _pat_let_tv87 = _3093_v150;
      if (function (_source24) {
        if (_source24.is_DC4) {
          let _3182___mcc_h20 = (_source24).cf7;
          let _3183___mcc_h21 = (_source24).cf8;
          let _3184_cf8 = _3183___mcc_h21;
          let _3185_cf7 = _3182___mcc_h20;
          return _3185_cf7;
        } else if (_source24.is_DC5) {
          let _3186___mcc_h22 = (_source24).cf9;
          let _3187___mcc_h23 = (_source24).cf10;
          let _3188___mcc_h24 = (_source24).cf11;
          let _3189_cf11 = _3188___mcc_h24;
          let _3190_cf10 = _3187___mcc_h23;
          let _3191_cf9 = _3186___mcc_h22;
          return _pat_let_tv87.f7;
        } else if (_source24.is_DC3) {
          let _3192___mcc_h25 = (_source24).cf6;
          let _3193_cf6 = _3192___mcc_h25;
          return false;
        } else {
          let _3194___mcc_h26 = (_source24).cf12;
          let _3195_cf12 = _3194___mcc_h26;
          return false;
        }
      }(((true) ? (_3181_v201) : (_3181_v201)))) {
        let _index501 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
        (_3096_v153)[_index501] = (_this).f6;
        r2 = _3093_v150.f7;
        let _3196_v202;
        _3196_v202 = _module.D6.create_DC16(_3093_v150.f7);
        if ((_3196_v202).dtor_cf30) {
          (_3093_v150).f7 = _3093_v150.f7;
          let _3197_v203;
          let _nw550 = new _module.C2();
          _nw550.__ctor(_dafny.Set.fromElements(_3090_v147, _3090_v147, _3090_v147, _3090_v147), _3092_v149, (new BigNumber((_3086_v143).length)).isLessThan((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]));
          _3197_v203 = _nw550;
          _3197_v203 = _3197_v203;
          let _3198_v204;
          _3198_v204 = _dafny.MultiSet.fromElements((_3087_v144)[_module.__default.safeIndex((_dafny.ZERO).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]), new BigNumber((_3087_v144).length))]);
          let _3199_v205;
          _3199_v205 = _dafny.MultiSet.fromElements(_3198_v204);
          let _3200_v206;
          _3200_v206 = _dafny.Seq.of(_3199_v205);
          _3199_v205 = ((_3200_v206)[_module.__default.safeIndex((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))], new BigNumber((_3200_v206).length))]).Intersect(_3199_v205);
          r1 = ((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]).multipliedBy((new BigNumber((_dafny.Seq.UnicodeFromString("tin")).length)).minus((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]));
          let _3201_v207;
          _3201_v207 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((((_3089_v146).contains(!(_3093_v150.f7))) ? ((_3089_v146).get(!(_3093_v150.f7))) : (_module.__default.fm0(globalState))), (_this).f6),(_3197_v203).f16);
          _3201_v207 = (_3201_v207).update((_this).f6, _3092_v149);
        } else {
          let _3202_v208;
          let _init100 = ((_3203_v153) => function (_3204_i15) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_3205_v153) => function (_3206_i16) {
              return (_3205_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3205_v153).length))];
            })(_3203_v153));
          })(_3096_v153);
          let _nw551 = Array((new BigNumber(6)).toNumber());
          for (let _i0_100 = 0; _i0_100 < new BigNumber(_nw551.length); _i0_100++) {
            _nw551[_i0_100] = _init100(new BigNumber(_i0_100));
          }
          _3202_v208 = _nw551;
          let _index502 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_3202_v208).length));
          (_3202_v208)[_index502] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), ((_3207_v150, _3208_v149) => function (_3209_i17) {
            return (_module.D6.create_DC17(_3207_v150.f7, (_this).f6, _3208_v149, _3208_v149)).dtor_cf32;
          })(_3093_v150, _3092_v149));
          let _3210_v209;
          _3210_v209 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber(756), (_dafny.ZERO).minus(new BigNumber((_3086_v143).length))), (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]);
          let _3211_v210;
          _3211_v210 = _dafny.Seq.of(_3086_v143);
          let _index503 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_3202_v208).length));
          let _rhs491 = _dafny.Seq.Concat((_3211_v210)[_module.__default.safeIndex(new BigNumber((_3087_v144).length), new BigNumber((_3211_v210).length))], _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-261), (_this).f6), _3086_v143));
          let _rhs492 = ((_3210_v209).Intersect(_3210_v209)).Difference(_3210_v209);
          let _lhs293 = _3202_v208;
          let _lhs294 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_3202_v208).length));
          _lhs293[_lhs294] = _rhs491;
          _3210_v209 = _rhs492;
          let _3212_v211;
          let _init101 = ((_3213_v209) => function (_3214_i18) {
            return _module.D2.create_DC7(_3213_v209);
          })(_3210_v209);
          let _nw552 = Array((new BigNumber(5)).toNumber());
          for (let _i0_101 = 0; _i0_101 < new BigNumber(_nw552.length); _i0_101++) {
            _nw552[_i0_101] = _init101(new BigNumber(_i0_101));
          }
          _3212_v211 = _nw552;
          let _3215_v212;
          let _nw553 = new _module.C4();
          _nw553.__ctor(_3212_v211, _2898_v0);
          _3215_v212 = _nw553;
          _3215_v212 = _3215_v212;
          let _index504 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index504] = ((_2898_v0) ? (_3093_v150.f7) : (_3093_v150.f7));
          let _index505 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index505] = _3093_v150.f7;
          let _3216_v213;
          _3216_v213 = _dafny.Seq.of(_3095_v152, _3095_v152);
          r0 = _module.__default.fm1(_3216_v213, ((_3093_v150.f7) ? (new BigNumber(778)) : ((_this).f6)), (new BigNumber((_3087_v144).length)).plus((_dafny.ZERO).minus((_this).f6)), globalState);
          let _3217_v214;
          let _nw554 = Array((new BigNumber(19)).toNumber());
          _3217_v214 = _nw554;
          let _3218_v215;
          _3218_v215 = _dafny.Map.Empty.slice().updateUnsafe(_3180_v200,_3217_v214);
          _3217_v214 = (((_3218_v215).contains(_3180_v200)) ? ((_3218_v215).get(_3180_v200)) : (_3217_v214));
        }
        if ((_module.__default.safeModuloInt((_this).f6, (_this).f6)).isLessThanOrEqualTo((_this).f6)) {
          let _index506 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_3090_v147).length));
          (_3090_v147)[_index506] = _2898_v0;
          let _3219_v216;
          let _nw555 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _3219_v216 = _nw555;
          let _index507 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_3219_v216).length));
          (_3219_v216)[_index507] = _3092_v149;
          let _3220_v217;
          _3220_v217 = _dafny.MultiSet.fromElements(_3092_v149, _3092_v149);
          let _3221_v218;
          _3221_v218 = _dafny.Seq.of(_3089_v146);
          let _index508 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_3090_v147).length));
          let _index509 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_3219_v216).length));
          let _index510 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          let _rhs493 = (_this).f6;
          let _rhs494 = (_3220_v217).IsDisjointFrom(_dafny.MultiSet.fromElements(_3092_v149, new _dafny.CodePoint('n'.codePointAt(0)), _3092_v149));
          let _rhs495 = ((_3093_v150.f7) ? (_3092_v149) : (_3092_v149));
          let _rhs496 = _module.__default.fm0(globalState);
          let _rhs497 = !(_dafny.Seq.IsProperPrefixOf(_3221_v218, _3221_v218));
          let _lhs295 = _3090_v147;
          let _lhs296 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_3090_v147).length));
          let _lhs297 = _3219_v216;
          let _lhs298 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_3219_v216).length));
          let _lhs299 = _3096_v153;
          let _lhs300 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          r3 = _rhs493;
          _lhs295[_lhs296] = _rhs494;
          _lhs297[_lhs298] = _rhs495;
          _lhs299[_lhs300] = _rhs496;
          r0 = _rhs497;
          r2 = true;
          let _3222_v219;
          let _3223_v220;
          let _out8;
          let _out9;
          let _outcollector3 = (_3093_v150).m2((_3090_v147)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_3090_v147).length))], (_this).f6, !_dafny.Seq.contains(_3087_v144, (_3219_v216)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_3219_v216).length))]), globalState);
          _out8 = _outcollector3[0];
          _out9 = _outcollector3[1];
          _3222_v219 = _out8;
          _3223_v220 = _out9;
          let _index511 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          (_3096_v153)[_index511] = (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))];
          let _3224_v221;
          let _nw556 = new _module.C2();
          _nw556.__ctor(_3091_v148, _3092_v149, _3093_v150.f7);
          _3224_v221 = _nw556;
        } else {
          _3090_v147 = _3090_v147;
          let _3225_v222;
          let _nw557 = Array((new BigNumber(6)).toNumber());
          _nw557[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_3087_v144, _3087_v144);
          _nw557[(_dafny.ONE).toNumber()] = _3087_v144;
          _nw557[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_3087_v144, _3087_v144);
          _nw557[(new BigNumber(3)).toNumber()] = _3087_v144;
          _nw557[(new BigNumber(4)).toNumber()] = _3087_v144;
          _nw557[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("wc");
          _3225_v222 = _nw557;
          let _index512 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_3225_v222).length));
          (_3225_v222)[_index512] = _dafny.Seq.Concat(_3087_v144, _dafny.Seq.UnicodeFromString("pbrvad"));
          let _index513 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_3225_v222).length));
          let _rhs498 = _dafny.Seq.Concat((_this).fm6(_3095_v152, globalState), _3087_v144);
          let _rhs499 = _3093_v150;
          let _rhs500 = _module.__default.safeModuloInt(new BigNumber((_3087_v144).length), new BigNumber(((_3088_v145).update((_this).f6, (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))])).length));
          let _rhs501 = !(!((new BigNumber((_3086_v143).length)).isLessThanOrEqualTo((_this).f6)));
          let _rhs502 = _2898_v0;
          let _lhs301 = _3225_v222;
          let _lhs302 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_3225_v222).length));
          let _lhs303 = _3093_v150;
          _lhs301[_lhs302] = _rhs498;
          _3093_v150 = _rhs499;
          r3 = _rhs500;
          r2 = _rhs501;
          _lhs303.f7 = _rhs502;
          (_3093_v150).f7 = _2898_v0;
          let _index514 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          let _rhs503 = (_this).f6;
          let _rhs504 = _3092_v149;
          let _lhs304 = _3096_v153;
          let _lhs305 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          _lhs304[_lhs305] = _rhs503;
          _3092_v149 = _rhs504;
          let _index515 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
          (_3096_v153)[_index515] = (_dafny.ZERO).minus((_this).f6);
        }
        let _3226_v223;
        _3226_v223 = _dafny.Map.Empty.slice().updateUnsafe(false,_3093_v150.f7);
        let _index516 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_3090_v147).length));
        (_3090_v147)[_index516] = (((_3226_v223).contains(_3093_v150.f7)) ? ((_3226_v223).get(_3093_v150.f7)) : (_2898_v0));
        let _index517 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_3090_v147).length));
        (_3090_v147)[_index517] = !(_3093_v150.f7) || (_2898_v0);
      } else {
        r3 = (_this).f6;
        let _index518 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
        (_3096_v153)[_index518] = (_module.__default.safeDivisionInt((_this).f6, (_this).fm5(globalState))).plus(_module.__default.fm0(globalState));
        let _3227_v224;
        _3227_v224 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), function (_3228_i19) {
          return (_this).f6;
        }));
        let _index519 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
        (_3096_v153)[_index519] = _module.__default.safeModuloInt((_this).f6, new BigNumber((_3227_v224).cardinality()));
        let _3229_v225;
        _3229_v225 = _module.D21.create_DC59(_3090_v147, _dafny.Seq.update(_3087_v144, _module.__default.safeIndex(new BigNumber((_3087_v144).length), new BigNumber((_3087_v144).length)), _3092_v149), _3096_v153, true);
        let _3230_v226;
        _3230_v226 = _dafny.Seq.of(_3229_v225, _3229_v225);
        let _index520 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
        let _rhs505 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f6, new BigNumber((_3087_v144).length)));
        let _rhs506 = _dafny.Seq.IsProperPrefixOf(_3230_v226, _dafny.Seq.Concat(_dafny.Seq.of(_3229_v225), _3230_v226));
        let _lhs306 = _3096_v153;
        let _lhs307 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length));
        let _lhs308 = _3093_v150;
        _lhs306[_lhs307] = _rhs505;
        _lhs308.f7 = _rhs506;
        let _3231_v227;
        _3231_v227 = _module.D9.create_DC24();
        _3231_v227 = _3231_v227;
      }
      let _3232_v228;
      _3232_v228 = _dafny.Map.Empty.slice().updateUnsafe(_2898_v0,_3093_v150.f7);
      _3232_v228 = (_3232_v228).update(_2898_v0, ((_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))]).isLessThan((_this).f6));
      let _3233_v229;
      let _nw558 = new _module.C3();
      _nw558.__ctor((_dafny.ZERO).minus((_this).f6), _3093_v150.f7);
      _3233_v229 = _nw558;
      r0 = (new BigNumber((_dafny.MultiSet.fromElements(_3233_v229, _3233_v229, _3233_v229)).cardinality())).isEqualTo(((_this).f6).plus(new BigNumber(207)));
      r1 = (_this).f6;
      r2 = _3233_v229.f7;
      r3 = (_3096_v153)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_3096_v153).length))];
      return [r0, r1, r2, r3];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
