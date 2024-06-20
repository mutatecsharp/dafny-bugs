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
    static fm0(p0, p1, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(984))).length))).length), new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(602), new BigNumber(705))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(602)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(705)))) {
            _coll0.push([_module.__default.safeDivisionInt(_0_v0, new BigNumber(803)),_dafny.Seq.of(new BigNumber(-786))]);
          }
        }
        return _coll0;
      }()).length))).minus((new BigNumber(376)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("fsmqogqg")).length)));
    };
    static fm1(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(32))).cardinality()))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(756)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(43), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(874))).length))).length)), _dafny.Seq.of(new BigNumber(268), new BigNumber(281)))));
    };
    static fm2(p0, p1, p2, globalState) {
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("hidhmds")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(427)), function (_1_i1) {
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0))))).cardinality());
      })).length))).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_2_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length));
    };
    static fm3(globalState) {
      return new _dafny.CodePoint('v'.codePointAt(0));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat((_module.D8.create_DC26(_module.D1.create_DC3(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))), _dafny.Seq.UnicodeFromString("cu"))).dtor_cf42, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ccsyvjc"), _dafny.Seq.UnicodeFromString("rdcttheo")));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(129),new BigNumber(577))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(711),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((_module.D8.create_DC23(_dafny.Set.fromElements(true))).dtor_cf36).length))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(202), new BigNumber(-405))) {
          let _3_v0 = _compr_1;
          if (((new BigNumber(202)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(-405)))) {
            _coll1.push([_module.__default.safeModuloInt(_3_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(684),true)).length)),!(true)]);
          }
        }
        return _coll1;
      }()).length)),new BigNumber(((_module.D19.create_DC44(_dafny.Map.Empty.slice().updateUnsafe(!(true),false))).dtor_cf67).length))));
    };
    static fm7(globalState) {
      return _dafny.Seq.of(new BigNumber(-417));
    };
    static fm13(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm14(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Union((_dafny.MultiSet.fromElements(true, false)).Union(_dafny.MultiSet.fromElements(true, true, true, !(!(!(false))), true)));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(155))).cardinality()))).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)),(new BigNumber((_dafny.Seq.of(true)).length)).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)))).cardinality())));
    };
    static fm16(globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vnbdljh")).length)))).length)).isLessThan(new BigNumber(286))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-275),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),!(false)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-59),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))).cardinality()))),true));
      }
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.of(true), _dafny.Seq.of(!(true), true, true)),(new BigNumber(-570)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("sx")).length)));
    };
    static fm18(globalState) {
      return (_module.D6.create_DC16(function () {
  let _coll2 = new _dafny.Set();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(418), new BigNumber(-290))) {
    let _4_v0 = _compr_2;
    if (((new BigNumber(418)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(-290)))) {
      _coll2.add((_4_v0).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(891), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(661)), function (_5_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length))).cardinality())));
    }
  }
  return _coll2;
}())).dtor_cf27;
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(!(true), true), _dafny.Seq.of(false, true)), _dafny.Seq.of(true));
    };
    static fm20(globalState) {
      return (((false) ? (_module.D20.create_DC48(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(!(false)), false, !(false)),new BigNumber(812)))) : (_module.D20.create_DC48(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),new BigNumber(930)))))).dtor_cf75;
    };
    static fm21(globalState) {
      return _module.D0.create_DC1((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false)).length), new BigNumber((_dafny.Seq.of(true, !(false))).length))), (new BigNumber((_dafny.Set.fromElements(new BigNumber(790))).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), function (_6_i0) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})).length)), (_module.D10.create_DC29(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length), false)).dtor_cf46);
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_7_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), function (_8_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })));
    };
    static fm23(globalState) {
      return _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_9_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("ytfqd")));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-428)), function (_10_i0) {
        return new BigNumber(814);
      });
    };
    static fm25(globalState) {
      return _module.D4.create_DC13((new BigNumber(994)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
  let _coll3 = new _dafny.Set();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(38), new BigNumber(-926))) {
    let _11_v0 = _compr_3;
    if (((new BigNumber(38)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(-926)))) {
      _coll3.add((_11_v0).plus(new BigNumber(960)));
    }
  }
  return _coll3;
}()).length)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("folgy"), _dafny.Seq.UnicodeFromString("q")));
    };
    static fm26(globalState) {
      return _module.D4.create_DC14((_module.D4.create_DC14(_module.D4.create_DC14(_module.D4.create_DC12(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))))).dtor_cf25);
    };
    static fm27(globalState) {
      return _module.D1.create_DC6();
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(!(false), true, true, !(true))).Difference(_dafny.Set.fromElements(!((_module.D1.create_DC4(true, false, new _dafny.CodePoint('e'.codePointAt(0)), new BigNumber(554), new BigNumber(784))).dtor_cf8), true, false, false))).Difference(_dafny.Set.fromElements(true, false, true, true, !(true)));
    };
    static fm29(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!(false), false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-424),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(643)), function (_12_i0) {
        return new BigNumber(398);
      }))).cardinality()))).length), new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(58), new BigNumber(520))).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xb")).length))),new BigNumber(311))).length))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(289))).length)))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(19), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(32), new BigNumber(765)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), function (_13_i1) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-402)), function (_14_i2) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length);
      })).length), new BigNumber((_dafny.Set.fromElements(true, false)).length), new BigNumber(-421))));
    };
    static fm30(p0, p1, globalState) {
      return (function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(232), new BigNumber(455))) {
            let _15_v1 = _compr_5;
            if (((new BigNumber(232)).isLessThanOrEqualTo(_15_v1)) && ((_15_v1).isLessThan(new BigNumber(455)))) {
              _coll5.add(_module.__default.safeModuloInt(_15_v1, new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)));
            }
          }
          return _coll5;
        }()).Elements) {
          let _16_v0 = _compr_4;
          if ((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(232), new BigNumber(455))) {
              let _17_v1 = _compr_6;
              if (((new BigNumber(232)).isLessThanOrEqualTo(_17_v1)) && ((_17_v1).isLessThan(new BigNumber(455)))) {
                _coll6.add(_module.__default.safeModuloInt(_17_v1, new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)));
              }
            }
            return _coll6;
          }()).contains(_16_v0)) {
            _coll4.push([_module.__default.safeDivisionInt(_16_v0, new BigNumber(845)),new _dafny.CodePoint('s'.codePointAt(0))]);
          }
        }
        return _coll4;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(393)), function (_18_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)),new _dafny.CodePoint('e'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(695),new _dafny.CodePoint('k'.codePointAt(0)))));
    };
    static fm31(p0, p1, globalState) {
      return _module.D6.create_DC19(_module.D6.create_DC19(_module.D6.create_DC18(false)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(false, true, new _dafny.CodePoint('r'.codePointAt(0)), new BigNumber(-724), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-318)), function (_19_i0) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length)),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(true, !(false), new _dafny.CodePoint('l'.codePointAt(0)), new BigNumber(-368), new BigNumber(-261)),false));
    };
    static fm33(p0, p1, globalState) {
      return _module.D3.create_DC10();
    };
    static fm34(p0, globalState) {
      return function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aqy"))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("bjuydaqli")))).Elements) {
          let _20_v0 = _compr_7;
          if (((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aqy"))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("bjuydaqli")))).contains(_20_v0)) {
            _coll7.push([_20_v0,_module.D1.create_DC5(!(false), !(true), !(true))]);
          }
        }
        return _coll7;
      }();
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), function (_21_i0) {
        return _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(51))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("rvhgqeka")).length));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), function (_22_i1) {
        return _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)), new BigNumber(-183), (_dafny.ZERO).minus(new BigNumber(-324)), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(!(true), true))).length))).length), new BigNumber(-687));
      }));
    };
    static fm36(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-904)), function (_23_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC33(new _dafny.CodePoint('r'.codePointAt(0)), !(false))).dtor_cf55,false);
      }));
    };
    static fm37(globalState) {
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ycqwr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), function (_24_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }))) {
        return _module.D8.create_DC26(_module.D1.create_DC3((_module.D8.create_DC26(_module.D1.create_DC3(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))), _dafny.Seq.UnicodeFromString("ndi"))).dtor_cf42), _dafny.Seq.UnicodeFromString("w"));
      } else {
        return _module.D8.create_DC26(_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), function (_25_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})), _dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), function (_26_i2) {
  return new _dafny.CodePoint('w'.codePointAt(0));
}));
      }
    };
    static fm38(globalState) {
      return _dafny.MultiSet.fromElements(_module.D3.create_DC11(new BigNumber(442), false, true));
    };
    static fm39(p0, p1, globalState) {
      return (function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(80), new BigNumber(406))) {
          let _27_v0 = _compr_8;
          if (((new BigNumber(80)).isLessThanOrEqualTo(_27_v0)) && ((_27_v0).isLessThan(new BigNumber(406)))) {
            _coll8.push([(_27_v0).minus(new BigNumber(-269)),_dafny.Set.fromElements(new BigNumber(522), new BigNumber((_dafny.Seq.UnicodeFromString("gc")).length))]);
          }
        }
        return _coll8;
      }()).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(977),_dafny.Set.fromElements(new BigNumber(826)))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),_dafny.Set.fromElements(new BigNumber(380))))));
    };
    static fm40(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(828)), function (_28_i0) {
        return _dafny.Seq.UnicodeFromString("q");
      }), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fil"))), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("canhopsk")));
    };
    static fm41(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("nw")).length))).Difference((_module.D13.create_DC35(_dafny.MultiSet.fromElements(new BigNumber(98)))).dtor_cf57), _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("bcleahc")).length))).cardinality()))), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-165)),new BigNumber((_dafny.Set.fromElements(true, true, true)).length))).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(709),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jjwyqqi")).length))).length))).Keys.Elements) {
          let _29_v0 = _compr_9;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(709),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("jjwyqqi")).length))).length))).contains(_29_v0)) {
            _coll9.push([_module.__default.safeModuloInt(_29_v0, new BigNumber(620)),false]);
          }
        }
        return _coll9;
      }()).length), new BigNumber(-251), new BigNumber((_dafny.Seq.of(!(true))).length), new BigNumber((_dafny.Seq.UnicodeFromString("uqgceiacv")).length))));
    };
    static fm42(p0, globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(-853)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements(true)).length)));
    };
    static fm43(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(!(false) || (false), false, new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), function (_30_i0) {
  return new BigNumber((_dafny.Seq.of(true, !(!(true)))).length);
}))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(439))))).length), _module.__default.safeModuloInt(new BigNumber(124), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ahdj")).length))).cardinality())));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let _31_v0;
      _31_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(4));
      let _32_v1;
      _32_v1 = _dafny.Seq.UnicodeFromString("nfp");
      _31_v0 = (_31_v0).update(new BigNumber((_32_v1).length), p1);
      let _33_v2;
      _33_v2 = _dafny.MultiSet.fromElements(p1);
      let _34_v3;
      _34_v3 = _dafny.Seq.of((_33_v2).update(new BigNumber((_31_v0).length), _module.__default.abs(_module.__default.fm0(p0, p0, globalState))));
      let _35_v4;
      _35_v4 = false;
      let _36_v5;
      _36_v5 = _dafny.Set.fromElements(_35_v4, _35_v4, _35_v4, true);
      let _37_v6;
      _37_v6 = new _dafny.CodePoint('w'.codePointAt(0));
      let _38_v7;
      _38_v7 = _dafny.Map.Empty.slice().updateUnsafe(_37_v6,p1);
      let _39_v8;
      _39_v8 = _dafny.Seq.of(_38_v7);
      let _40_v9;
      let _nw0 = new _module.C1();
      _nw0.__ctor(_34_v3, p1, new BigNumber((_36_v5).length), _39_v8);
      _40_v9 = _nw0;
      let _41_v10;
      _41_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,_40_v9);
      let _42_v11;
      let _nw1 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _42_v11 = _nw1;
      let _43_v12;
      let _nw2 = new _module.C4();
      _nw2.__ctor(_42_v11, (_40_v9).f24, (_40_v9).f25);
      _43_v12 = _nw2;
      let _44_v13;
      _44_v13 = _dafny.MultiSet.fromElements(_43_v12, _43_v12, _43_v12, _43_v12, _43_v12);
      let _45_v14;
      _45_v14 = _40_v9;
      _41_v10 = (_41_v10).update(_module.__default.fm0((((_44_v13).contains(_43_v12)) ? ((_44_v13).get(_43_v12)) : (new BigNumber(235))), new BigNumber(986), globalState), (_40_v9));
      let _46_v15;
      _46_v15 = _dafny.Map.Empty.slice().updateUnsafe((_43_v12).f28,(_40_v9).f24);
      let _index0 = _module.__default.safeIndex(new BigNumber(862), new BigNumber(((_43_v12).f28).length));
      ((_43_v12).f28)[_index0] = new BigNumber((_46_v15).length);
      let _index1 = _module.__default.safeIndex(new BigNumber(862), new BigNumber(((_43_v12).f28).length));
      ((_43_v12).f28)[_index1] = p1;
      let _47_v16;
      _47_v16 = _dafny.Map.Empty.slice().updateUnsafe(_35_v4,(_43_v12).f28);
      _47_v16 = (_47_v16).Merge((_47_v16).Merge(_47_v16));
      let _hi0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-669)), ((_48_v6) => function (_49_i1) {
        return _48_v6;
      })(_37_v6))).length);
      for (let _50_i0 = (_dafny.ZERO).minus(((_43_v12).f28)[_module.__default.safeIndex(new BigNumber(862), new BigNumber(((_43_v12).f28).length))]); _50_i0.isLessThan(_hi0); _50_i0 = _50_i0.plus(_dafny.ONE)) {
        let _51_v17;
        let _init0 = ((_52_v1) => function (_53_i2) {
          return !_dafny.Seq.contains(_52_v1, new _dafny.CodePoint('w'.codePointAt(0)));
        })(_32_v1);
        let _nw3 = Array((new BigNumber(15)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
          _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _51_v17 = _nw3;
        let _54_v18;
        _54_v18 = _dafny.Map.Empty.slice().updateUnsafe((_40_v9).f24,((true) ? (_51_v17) : (_51_v17)));
        _51_v17 = (((_54_v18).contains((_40_v9).f24)) ? ((_54_v18).get((_40_v9).f24)) : (_51_v17));
        let _55_v19;
        _55_v19 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(694)), ((_56_v6) => function (_57_i3) {
          return _56_v6;
        })(_37_v6)));
        let _58_v20;
        _58_v20 = _dafny.Seq.of(_50_i0, new BigNumber((_dafny.Seq.UnicodeFromString("ipsloox")).length));
        let _59_v21;
        _59_v21 = _dafny.Set.fromElements(new BigNumber((_55_v19).cardinality()), new BigNumber((_58_v20).length));
        let _60_v22;
        _60_v22 = _module.D10.create_DC29((_40_v9).f24, _35_v4);
        (globalState).f18 = ((_35_v4) ? ((((_43_v12).f28)[_module.__default.safeIndex(new BigNumber(862), new BigNumber(((_43_v12).f28).length))]).minus(new BigNumber((_59_v21).length))) : (_module.__default.safeModuloInt((_60_v22).dtor_cf45, ((_43_v12).f28)[_module.__default.safeIndex(new BigNumber(862), new BigNumber(((_43_v12).f28).length))])));
        let _61_v23;
        let _nw4 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
        _61_v23 = _nw4;
        let _62_v24;
        _62_v24 = _dafny.Seq.of(_35_v4);
        let _index2 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_61_v23).length));
        (_61_v23)[_index2] = _62_v24;
        let _index3 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_61_v23).length));
        let _rhs0 = _dafny.Seq.Concat(_62_v24, _dafny.Seq.update(_dafny.Seq.of(_35_v4, _35_v4, _35_v4, _35_v4), _module.__default.safeIndex(_50_i0, new BigNumber((_dafny.Seq.of(_35_v4, _35_v4, _35_v4, _35_v4)).length)), _35_v4));
        let _rhs1 = (_dafny.ZERO).minus(((_40_v9).f24).plus(p0));
        let _rhs2 = _36_v5;
        let _lhs0 = _61_v23;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_61_v23).length));
        let _lhs2 = globalState;
        _lhs0[_lhs1] = _rhs0;
        _lhs2.f1 = _rhs1;
        _36_v5 = _rhs2;
        let _63_v25;
        _63_v25 = _dafny.MultiSet.fromElements(_35_v4);
        let _64_v26;
        _64_v26 = _dafny.Map.Empty.slice().updateUnsafe((_63_v25).Difference(_63_v25),false);
        _64_v26 = (_64_v26).update(_63_v25, (_35_v4) === (true));
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_42_v11).length))) {
        let _65_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_65_i4)) && ((_65_i4).isLessThan(new BigNumber((_42_v11).length))))) {
          (_42_v11)[(_65_i4)] = (_65_i4).multipliedBy(p1);
        }
      }
      let _66_v27;
      _66_v27 = _dafny.Map.Empty.slice().updateUnsafe(_35_v4,p0);
      let _67_v28;
      _67_v28 = _module.D4.create_DC13(new BigNumber(-145), _66_v27, _32_v1);
      r0 = (_dafny.ZERO).minus(new BigNumber(((_67_v28).dtor_cf24).length));
      return r0;
    }
    static Main(__noArgsParameter) {
      let _68_v0;
      _68_v0 = new BigNumber(590);
      let _69_v1;
      _69_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,_68_v0);
      let _70_v2;
      _70_v2 = true;
      let _71_v3;
      let _init1 = function (_72_i0) {
        return (_72_i0).minus(new BigNumber(-719));
      };
      let _nw5 = Array((new BigNumber(24)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
        _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _71_v3 = _nw5;
      let _73_v4;
      _73_v4 = _module.D0.create_DC0(_71_v3);
      let _74_v5;
      _74_v5 = _dafny.Map.Empty.slice().updateUnsafe(_71_v3,(_73_v4).dtor_cf0);
      let _75_v6;
      _75_v6 = _dafny.Seq.of(_70_v2);
      let _76_v7;
      _76_v7 = new _dafny.CodePoint('c'.codePointAt(0));
      let _77_v8;
      _77_v8 = _module.D1.create_DC3(_dafny.Seq.of(_76_v7, new _dafny.CodePoint('h'.codePointAt(0))));
      let _78_v9;
      _78_v9 = _dafny.MultiSet.fromElements(_76_v7);
      let _79_v10;
      _79_v10 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray((_77_v8).dtor_cf7), _78_v9);
      let _80_v11;
      let _nw6 = Array((new BigNumber(2)).toNumber()).fill(false);
      _80_v11 = _nw6;
      let _81_v12;
      _81_v12 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-991)), ((_82_v0) => function (_83_i1) {
        return _82_v0;
      })(_68_v0))).length)),_80_v11);
      let _84_v13;
      _84_v13 = _dafny.Set.fromElements(_70_v2, _70_v2, _70_v2);
      let _85_v14;
      _85_v14 = _dafny.Seq.UnicodeFromString("nbntyekh");
      let _86_v15;
      _86_v15 = _module.D1.create_DC4(_70_v2, _70_v2, (_85_v14)[_module.__default.safeIndex(_68_v0, new BigNumber((_85_v14).length))], _68_v0, _68_v0);
      let _87_v16;
      _87_v16 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(382),_68_v0));
      let _88_globalState;
      let _nw7 = new _module.GlobalState();
      _nw7.__ctor(true, new BigNumber(554), true, (_69_v1).update(_70_v2, _68_v0), false, new BigNumber(-61), new _dafny.CodePoint('c'.codePointAt(0)), _74_v5, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_70_v2, _70_v2, !(false)), _75_v6), _module.__default.safeIndex((_dafny.ZERO).minus(_68_v0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_70_v2, _70_v2, !(false)), _75_v6)).length)), !(_70_v2)), new BigNumber(889), new BigNumber(204), _79_v10, new BigNumber(127), _81_v12, new BigNumber(996), new BigNumber(982), _84_v13, _85_v14, new BigNumber(-607), _dafny.MultiSet.fromElements(_80_v11, _80_v11), _75_v6, _dafny.Seq.of(_70_v2, (_86_v15).dtor_cf8), (((_87_v16).contains(_68_v0)) ? ((_87_v16).get(_68_v0)) : (function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(199), new BigNumber(972))) {
          let _89_v17 = _compr_10;
          if (((new BigNumber(199)).isLessThanOrEqualTo(_89_v17)) && ((_89_v17).isLessThan(new BigNumber(972)))) {
            _coll10.push([(_89_v17).minus(_68_v0),_68_v0]);
          }
        }
        return _coll10;
      }())), new BigNumber(790));
      _88_globalState = _nw7;
      let _index4 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
      (_80_v11)[_index4] = (_68_v0).isLessThanOrEqualTo(_68_v0);
      let _index5 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
      (_80_v11)[_index5] = _70_v2;
      _76_v7 = new _dafny.CodePoint('u'.codePointAt(0));
      let _90_v18;
      _90_v18 = _dafny.MultiSet.fromElements(!((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]), (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
      let _91_v19;
      _91_v19 = _dafny.Map.Empty.slice().updateUnsafe(_70_v2,false);
      let _92_v20;
      _92_v20 = _dafny.Seq.of(_68_v0, new BigNumber((_91_v19).length));
      if ((_module.__default.fm0(new BigNumber((_92_v20).length), _68_v0, _88_globalState)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_68_v0, (((_90_v18).contains(!((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]))) ? ((_90_v18).get(!((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]))) : (_68_v0))))) {
        let _93_v21;
        _93_v21 = _dafny.MultiSet.fromElements(_68_v0);
        let _94_v22;
        _94_v22 = _dafny.Map.Empty.slice().updateUnsafe(_76_v7,_93_v21);
        let _95_v23;
        _95_v23 = _dafny.Seq.of(_93_v21);
        (_88_globalState).f14 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_93_v21, _module.__default.fm1(_68_v0, _68_v0, _88_globalState), (((_94_v22).contains(_76_v7)) ? ((_94_v22).get(_76_v7)) : (_93_v21))), _95_v23)).length);
        _73_v4 = _module.D0.create_DC0(_71_v3);
        let _96_v24;
        _96_v24 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_85_v14).length), _68_v0, _88_globalState));
        _96_v24 = _96_v24;
        let _97_v25;
        _97_v25 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_84_v13);
        let _98_v26;
        _98_v26 = _dafny.Map.Empty.slice().updateUnsafe((_84_v13).IsProperSubsetOf((((_97_v25).contains(true)) ? ((_97_v25).get(true)) : (_dafny.Set.fromElements(_70_v2, _module.__default.fm2(new BigNumber(659), new BigNumber((_75_v6).length), _76_v7, _88_globalState))))),_71_v3);
        let _99_v27;
        let _nw8 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _99_v27 = _nw8;
        let _index6 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_99_v27).length));
        (_99_v27)[_index6] = (_85_v14)[_module.__default.safeIndex(_68_v0, new BigNumber((_85_v14).length))];
        let _100_v28;
        _100_v28 = _dafny.MultiSet.fromElements(_99_v27);
        let _index7 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_99_v27).length));
        let _index8 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
        let _rhs3 = _98_v26;
        let _rhs4 = _module.__default.fm3(_88_globalState);
        let _rhs5 = _68_v0;
        let _rhs6 = ((_100_v28).Union(_100_v28)).IsSubsetOf(_100_v28);
        let _rhs7 = _68_v0;
        let _lhs3 = _99_v27;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_99_v27).length));
        let _lhs5 = _88_globalState;
        let _lhs6 = _80_v11;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
        let _lhs8 = _88_globalState;
        _98_v26 = _rhs3;
        _lhs3[_lhs4] = _rhs4;
        _lhs5.f9 = _rhs5;
        _lhs6[_lhs7] = _rhs6;
        _lhs8.f1 = _rhs7;
        (_88_globalState).f9 = _module.__default.safeModuloInt((_68_v0).plus(new BigNumber(765)), new BigNumber(((_84_v13).Difference(_84_v13)).length));
      } else {
        let _101_v29;
        let _nw9 = Array((new BigNumber(4)).toNumber()).fill([]);
        _101_v29 = _nw9;
        let _102_v30;
        _102_v30 = _dafny.Map.Empty.slice().updateUnsafe(false,_71_v3);
        let _103_v31;
        _103_v31 = _module.D0.create_DC2((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], (_dafny.ZERO).minus(_68_v0), _76_v7);
        let _index9 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_101_v29).length));
        (_101_v29)[_index9] = (((_102_v30).contains(!((_103_v31).dtor_cf4))) ? ((_102_v30).get(!((_103_v31).dtor_cf4))) : (_71_v3));
        let _index10 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_101_v29).length));
        (_101_v29)[_index10] = _71_v3;
        let _104_v32;
        _104_v32 = _module.D0.create_DC1(new BigNumber(215), _68_v0, (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
        let _105_v33;
        _105_v33 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_85_v14);
        let _106_v34;
        let _nw10 = Array((new BigNumber(21)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_85_v14, _85_v14);
        _nw10[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_85_v14, _module.__default.safeIndex(_module.__default.fm0(_68_v0, _68_v0, _88_globalState), new BigNumber((_85_v14).length)), _76_v7);
        _nw10[(new BigNumber(2)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("orkpvq"), _module.__default.safeIndex(_68_v0, new BigNumber((_dafny.Seq.UnicodeFromString("orkpvq")).length)), _module.__default.fm3(_88_globalState));
        _nw10[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_85_v14, _module.__default.fm4((_104_v32).dtor_cf3, (_dafny.ZERO).minus(new BigNumber((_92_v20).length)), _76_v7, _module.__default.fm2(_68_v0, _68_v0, _76_v7, _88_globalState), _88_globalState));
        _nw10[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_85_v14, _85_v14);
        _nw10[(new BigNumber(6)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(7)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(8)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-260)), ((_107_v7) => function (_108_i2) {
          return _107_v7;
        })(_76_v7));
        _nw10[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat((((_105_v33).contains(new BigNumber((_84_v13).length))) ? ((_105_v33).get(new BigNumber((_84_v13).length))) : (_85_v14)), _dafny.Seq.update(_85_v14, _module.__default.safeIndex(new BigNumber((_92_v20).length), new BigNumber((_85_v14).length)), _76_v7));
        _nw10[(new BigNumber(11)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("jmsoaylng");
        _nw10[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_85_v14, _85_v14);
        _nw10[(new BigNumber(14)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(15)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(16)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("arypuvscv");
        _nw10[(new BigNumber(18)).toNumber()] = _85_v14;
        _nw10[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_85_v14, _module.__default.safeIndex(_68_v0, new BigNumber((_85_v14).length)), _76_v7);
        _nw10[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("bmic");
        _106_v34 = _nw10;
        let _109_v35;
        _109_v35 = _dafny.Seq.of(_106_v34, _106_v34);
        _106_v34 = (_109_v35)[_module.__default.safeIndex((_86_v15).dtor_cf11, new BigNumber((_109_v35).length))];
        (_88_globalState).f2 = (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))];
        let _110_v36;
        _110_v36 = _dafny.Set.fromElements(_85_v14, _85_v14, _module.__default.fm4((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_68_v0,_70_v2)).length), _76_v7, _70_v2, _88_globalState));
        let _111_v37;
        _111_v37 = _dafny.Map.Empty.slice().updateUnsafe((_110_v36).Union(_110_v36),((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) || ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]));
        _111_v37 = (_111_v37).update(_110_v36, !((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]));
        let _112_v38;
        _112_v38 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_76_v7);
        _70_v2 = _module.__default.fm2((_68_v0).plus(_68_v0), _module.__default.safeDivisionInt(_68_v0, _68_v0), (((_112_v38).contains((_104_v32).dtor_cf1)) ? ((_112_v38).get((_104_v32).dtor_cf1)) : (_76_v7)), _88_globalState);
      }
      let _113_v39;
      _113_v39 = _module.D1.create_DC5((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], (_68_v0).isLessThanOrEqualTo(_68_v0), (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
      _113_v39 = _113_v39;
      let _hi1 = (_92_v20)[_module.__default.safeIndex(_68_v0, new BigNumber((_92_v20).length))];
      for (let _114_i3 = _68_v0; _114_i3.isLessThan(_hi1); _114_i3 = _114_i3.plus(_dafny.ONE)) {
        _85_v14 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ug"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(922)), function (_115_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })), _module.__default.safeIndex((_dafny.ZERO).minus(_68_v0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ug"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(922)), function (_116_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }))).length)), _76_v7), _85_v14);
        let _117_v40;
        let _out0;
        _out0 = _module.__default.m0(_114_i3, _114_i3, _88_globalState);
        _117_v40 = _out0;
        let _118_v41;
        _118_v41 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_117_v40);
        let _119_v43;
        _119_v43 = _dafny.Seq.of(_118_v41, _dafny.Map.Empty.slice().updateUnsafe(_68_v0,new BigNumber(-223)), function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_92_v20).Elements) {
            let _120_v42 = _compr_11;
            if (_dafny.Seq.contains(_92_v20, _120_v42)) {
              _coll11.push([_module.__default.safeDivisionInt(_120_v42, new BigNumber(135)),new BigNumber(597)]);
            }
          }
          return _coll11;
        }(), _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_114_i3), _118_v41);
        let _121_v44;
        _121_v44 = _dafny.Map.Empty.slice().updateUnsafe(_70_v2,_119_v43);
        _121_v44 = (_121_v44).update(_70_v2, _119_v43);
        let _index11 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_71_v3).length));
        (_71_v3)[_index11] = _68_v0;
        let _index12 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_71_v3).length));
        (_71_v3)[_index12] = _117_v40;
      }
      (_88_globalState).f18 = new BigNumber((_85_v14).length);
      let _122_v45;
      let _nw11 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _122_v45 = _nw11;
      let _index13 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length));
      (_122_v45)[_index13] = _85_v14;
      let _123_v46;
      _123_v46 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_85_v14, _85_v14, _dafny.Seq.UnicodeFromString("ncd"), _85_v14, _85_v14),_dafny.Seq.update(_dafny.Seq.UnicodeFromString("lbmfhj"), _module.__default.safeIndex(new BigNumber(-133), new BigNumber((_dafny.Seq.UnicodeFromString("lbmfhj")).length)), _76_v7));
      let _124_v47;
      _124_v47 = _dafny.Seq.of(_85_v14);
      let _125_v48;
      _125_v48 = _dafny.Seq.of(_85_v14, (_124_v47)[_module.__default.safeIndex(_68_v0, new BigNumber((_124_v47).length))], _85_v14, _85_v14, _85_v14);
      let _index14 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length));
      (_122_v45)[_index14] = (((_123_v46).contains(_124_v47)) ? ((_123_v46).get(_124_v47)) : (_dafny.Seq.UnicodeFromString("anrxvxycj")));
      let _rhs8 = ((_module.__default.fm2(new BigNumber((_91_v19).length), _68_v0, _76_v7, _88_globalState)) ? ((_dafny.ZERO).minus(_68_v0)) : (((!(_70_v2)) ? (_68_v0) : ((_dafny.ZERO).minus(_68_v0)))));
      let _rhs9 = !(((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) === (_70_v2));
      let _lhs9 = _88_globalState;
      _lhs9.f14 = _rhs8;
      _70_v2 = _rhs9;
      let _126_v49;
      _126_v49 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kyghqhrgx")).length));
      let _127_v50;
      _127_v50 = _module.D0.create_DC1(_68_v0, _68_v0, !((_75_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(760),_126_v49)).length), new BigNumber((_75_v6).length))]));
      let _source0 = _127_v50;
      if (_source0.is_DC1) {
        let _128___mcc_h0 = (_source0).cf1;
        let _129___mcc_h1 = (_source0).cf2;
        let _130___mcc_h2 = (_source0).cf3;
        let _131_cf3 = _130___mcc_h2;
        let _132_cf2 = _129___mcc_h1;
        let _133_cf1 = _128___mcc_h0;
        let _134_v51;
        let _nw12 = new _module.C0();
        _nw12.__ctor(!(_131_cf3) || (_70_v2));
        _134_v51 = _nw12;
        let _135_v52;
        let _out1;
        _out1 = _module.__default.m0(_133_cf1, _module.__default.safeDivisionInt(_68_v0, _68_v0), _88_globalState);
        _135_v52 = _out1;
        _131_cf3 = (_134_v51).f34;
        _131_cf3 = (_134_v51).f34;
      } else if (_source0.is_DC2) {
        let _136___mcc_h3 = (_source0).cf4;
        let _137___mcc_h4 = (_source0).cf5;
        let _138___mcc_h5 = (_source0).cf6;
        let _139_cf6 = _138___mcc_h5;
        let _140_cf5 = _137___mcc_h4;
        let _141_cf4 = _136___mcc_h3;
        _70_v2 = !(((_module.__default.fm2(_68_v0, new BigNumber((_75_v6).length), new _dafny.CodePoint('t'.codePointAt(0)), _88_globalState)) ? (_141_cf4) : (_70_v2))) || (!((_84_v13).IsSubsetOf(_84_v13)));
        let _142_v53;
        _142_v53 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_92_v20);
        _142_v53 = _142_v53;
        (_88_globalState).f9 = (_140_cf5).minus(new BigNumber(969));
        let _index15 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_71_v3).length));
        (_71_v3)[_index15] = _module.__default.safeDivisionInt(_68_v0, _140_cf5);
        let _index16 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_71_v3).length));
        (_71_v3)[_index16] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(776), _140_cf5));
      } else {
        let _143___mcc_h6 = (_source0).cf0;
        let _144_cf0 = _143___mcc_h6;
        if ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) {
          _68_v0 = _68_v0;
          let _145_v54;
          let _nw13 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _145_v54 = _nw13;
          let _index17 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_145_v54).length));
          (_145_v54)[_index17] = _92_v20;
          let _index18 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_145_v54).length));
          (_145_v54)[_index18] = _92_v20;
          let _init2 = ((_146_v0) => function (_147_i5) {
            return (_147_i5).minus(_146_v0);
          })(_68_v0);
          let _nw14 = Array((new BigNumber(29)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw14.length); _i0_2++) {
            _nw14[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _71_v3 = _nw14;
          let _148_v55;
          _148_v55 = _dafny.Map.Empty.slice().updateUnsafe(_76_v7,_68_v0);
          let _149_v56;
          _149_v56 = _dafny.Seq.of(_148_v55);
          let _150_v57;
          let _nw15 = new _module.C4();
          _nw15.__ctor(_71_v3, _68_v0, _149_v56);
          _150_v57 = _nw15;
          let _151_v58;
          _151_v58 = _dafny.Seq.of(_150_v57, _150_v57);
          let _152_v59;
          let _nw16 = new _module.C4();
          _nw16.__ctor(_144_cf0, new BigNumber((_151_v58).length), _149_v56);
          _152_v59 = _nw16;
          let _153_v60;
          _153_v60 = _dafny.MultiSet.fromElements(new BigNumber(-453), _68_v0);
          let _154_v61;
          _154_v61 = _dafny.Seq.of(_153_v60, _dafny.MultiSet.fromElements(_68_v0));
          let _155_v62;
          let _nw17 = new _module.C1();
          _nw17.__ctor(_154_v61, _68_v0, (((_69_v1).contains(true)) ? ((_69_v1).get(true)) : (new BigNumber((_153_v60).cardinality()))), _149_v56);
          _155_v62 = _nw17;
          let _156_v63;
          _156_v63 = _155_v62;
          let _157_v64;
          _157_v64 = _dafny.Seq.of(_155_v62);
          let _158_v65;
          let _nw18 = Array((new BigNumber(26)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _155_v62;
          _nw18[(_dafny.ONE).toNumber()] = _155_v62;
          _nw18[(new BigNumber(2)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(3)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(4)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(5)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(6)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(7)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(8)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(9)).toNumber()] = (_156_v63);
          _nw18[(new BigNumber(10)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(11)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(12)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(13)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(14)).toNumber()] = (_157_v64)[_module.__default.safeIndex(new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length), new BigNumber((_157_v64).length))];
          _nw18[(new BigNumber(15)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(16)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(17)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(18)).toNumber()] = (_157_v64)[_module.__default.safeIndex(_68_v0, new BigNumber((_157_v64).length))];
          _nw18[(new BigNumber(19)).toNumber()] = (_157_v64)[_module.__default.safeIndex(_155_v62.f33, new BigNumber((_157_v64).length))];
          _nw18[(new BigNumber(20)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(21)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(22)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(23)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(24)).toNumber()] = _155_v62;
          _nw18[(new BigNumber(25)).toNumber()] = _155_v62;
          _158_v65 = _nw18;
          let _index19 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_158_v65).length));
          (_158_v65)[_index19] = _155_v62;
          let _index20 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_158_v65).length));
          let _rhs10 = _76_v7;
          let _rhs11 = (_dafny.ZERO).minus(_68_v0);
          let _rhs12 = _155_v62;
          let _rhs13 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.update((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))], _module.__default.safeIndex(_155_v62.f33, new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length)), _76_v7), _85_v14), _85_v14);
          let _lhs10 = _88_globalState;
          let _lhs11 = _158_v65;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_158_v65).length));
          _76_v7 = _rhs10;
          _lhs10.f18 = _rhs11;
          _lhs11[_lhs12] = _rhs12;
          _70_v2 = _rhs13;
        } else {
          let _159_v67;
          _159_v67 = _dafny.Set.fromElements(_76_v7, _76_v7);
          let _160_v68;
          _160_v68 = _module.D4.create_DC12(_92_v20);
          let _161_v69;
          _161_v69 = _dafny.Map.Empty.slice().updateUnsafe((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))],_160_v68);
          let _162_v70;
          _162_v70 = _dafny.Map.Empty.slice().updateUnsafe(_76_v7,new BigNumber((_161_v69).length));
          let _163_v71;
          _163_v71 = _dafny.Seq.of(function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_159_v67).Elements) {
              let _164_v66 = _compr_12;
              if ((_159_v67).contains(_164_v66)) {
                _coll12.push([_164_v66,_68_v0]);
              }
            }
            return _coll12;
          }(), _162_v70);
          let _165_v72;
          let _nw19 = new _module.C2();
          _nw19.__ctor((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length), _163_v71);
          _165_v72 = _nw19;
          let _166_v73;
          _166_v73 = _dafny.Seq.of(_165_v72);
          let _167_v74;
          _167_v74 = _dafny.MultiSet.fromElements((_166_v73)[_module.__default.safeIndex(_68_v0, new BigNumber((_166_v73).length))], _165_v72);
          let _168_v75;
          _168_v75 = _dafny.Map.Empty.slice().updateUnsafe(true,_165_v72);
          let _169_v76;
          let _nw20 = new _module.C4();
          _nw20.__ctor(_71_v3, (((_167_v74).contains((((_168_v75).contains((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])) ? ((_168_v75).get((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])) : (_165_v72)))) ? ((_167_v74).get((((_168_v75).contains((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])) ? ((_168_v75).get((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])) : (_165_v72)))) : (_68_v0)), _163_v71);
          _169_v76 = _nw20;
          _169_v76 = _169_v76;
          _69_v1 = _69_v1;
          let _170_v77;
          let _nw21 = new _module.C2();
          _nw21.__ctor((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_92_v20)[_module.__default.safeIndex((_169_v76).f24, new BigNumber((_92_v20).length))], new BigNumber(687))), _module.__default.fm42((_169_v76).f24, _88_globalState));
          _170_v77 = _nw21;
          (_88_globalState).f17 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yqdcd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-23)), ((_171_v7) => function (_172_i6) {
            return _171_v7;
          })(_76_v7)));
          let _173_v78;
          let _nw22 = new _module.C0();
          _nw22.__ctor((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
          _173_v78 = _nw22;
          _173_v78 = _173_v78;
        }
        (_88_globalState).f2 = _dafny.Seq.contains(_dafny.Seq.Concat(_92_v20, _92_v20), _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.fm0(_68_v0, _68_v0, _88_globalState)), _68_v0));
        let _174_v79;
        let _nw23 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
        _174_v79 = _nw23;
        let _175_v80;
        _175_v80 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_76_v7,_68_v0)).update(_76_v7, _68_v0));
        let _176_v81;
        let _nw24 = new _module.C2();
        _nw24.__ctor(false, _68_v0, (_175_v80));
        _176_v81 = _nw24;
        let _177_v82;
        _177_v82 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,new BigNumber(83));
        let _178_v83;
        let _nw25 = Array((new BigNumber(4)).toNumber());
        _nw25[(_dafny.ZERO).toNumber()] = _177_v82;
        _nw25[(_dafny.ONE).toNumber()] = _177_v82;
        _nw25[(new BigNumber(2)).toNumber()] = _177_v82;
        _nw25[(new BigNumber(3)).toNumber()] = _177_v82;
        _178_v83 = _nw25;
        let _179_v84;
        _179_v84 = _dafny.Map.Empty.slice().updateUnsafe(_176_v81,_178_v83);
        let _180_v85;
        let _nw26 = new _module.C3();
        _nw26.__ctor(_174_v79, (((_179_v84).contains(_176_v81)) ? ((_179_v84).get(_176_v81)) : (_178_v83)), new BigNumber(869), _dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_181_v7, _182_v0) => function (_183_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe(_181_v7,_182_v0);
        })(_76_v7, _68_v0)));
        _180_v85 = _nw26;
        (_88_globalState).f18 = _module.__default.safeDivisionInt(new BigNumber((_85_v14).length), new BigNumber((_dafny.Seq.of(_180_v85)).length));
        let _184_v87;
        _184_v87 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_76_v7,_68_v0));
        let _185_v88;
        _185_v88 = _dafny.Map.Empty.slice().updateUnsafe(_76_v7,_68_v0);
        let _186_v89;
        let _nw27 = new _module.C4();
        _nw27.__ctor(_144_cf0, _68_v0, _dafny.Seq.update(_184_v87, _module.__default.safeIndex(_68_v0, new BigNumber((_184_v87).length)), _185_v88));
        _186_v89 = _nw27;
        let _187_v90;
        _187_v90 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(!((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]),_76_v7)).update(_176_v81.f31, _76_v7),_186_v89);
        _81_v12 = (_81_v12).update((new BigNumber((function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(62), new BigNumber(277))) {
            let _188_v86 = _compr_13;
            if (((new BigNumber(62)).isLessThanOrEqualTo(_188_v86)) && ((_188_v86).isLessThan(new BigNumber(277)))) {
              _coll13.add(_module.__default.safeModuloInt(_188_v86, _68_v0));
            }
          }
          return _coll13;
        }()).length)).multipliedBy(new BigNumber((_187_v90).length)), _80_v11);
      }
      let _hi2 = (_dafny.ZERO).minus(_68_v0);
      for (let _189_i8 = _module.__default.fm0(_68_v0, _68_v0, _88_globalState); _189_i8.isLessThan(_hi2); _189_i8 = _189_i8.plus(_dafny.ONE)) {
        let _index21 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
        (_80_v11)[_index21] = (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))];
        (_88_globalState).f18 = _189_i8;
        let _190_v91;
        _190_v91 = _dafny.MultiSet.fromElements(_189_i8);
        let _index22 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
        (_80_v11)[_index22] = ((((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) ? ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) : (_70_v2))) || (!(_dafny.MultiSet.fromElements(_189_i8)).equals(_190_v91));
        let _index23 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_71_v3).length));
        (_71_v3)[_index23] = _189_i8;
        let _191_v92;
        let _nw28 = Array((new BigNumber(27)).toNumber());
        _nw28[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
        _nw28[(_dafny.ONE).toNumber()] = _76_v7;
        _nw28[(new BigNumber(2)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(3)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(4)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(5)).toNumber()] = (((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]) ? (_76_v7) : (_76_v7));
        _nw28[(new BigNumber(6)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(7)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw28[(new BigNumber(9)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(10)).toNumber()] = ((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))])[_module.__default.safeIndex(_189_i8, new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length))];
        _nw28[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw28[(new BigNumber(12)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(13)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(14)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(15)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(16)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(17)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(18)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(19)).toNumber()] = _module.__default.fm3(_88_globalState);
        _nw28[(new BigNumber(20)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(21)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(22)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(23)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(24)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(25)).toNumber()] = _76_v7;
        _nw28[(new BigNumber(26)).toNumber()] = _76_v7;
        _191_v92 = _nw28;
        let _index24 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_191_v92).length));
        (_191_v92)[_index24] = _76_v7;
        let _192_v93;
        _192_v93 = _dafny.Map.Empty.slice().updateUnsafe(_70_v2,_76_v7);
        let _193_v94;
        _193_v94 = _dafny.Seq.of(_80_v11);
        let _194_v95;
        _194_v95 = _dafny.Map.Empty.slice().updateUnsafe(_80_v11,_80_v11);
        let _195_v96;
        let _nw29 = new _module.C0();
        _nw29.__ctor(_70_v2);
        _195_v96 = _nw29;
        let _196_v97;
        _196_v97 = _dafny.Seq.of(_195_v96);
        let _197_v98;
        _197_v98 = _dafny.MultiSet.fromElements((_193_v94)[_module.__default.safeIndex(_68_v0, new BigNumber((_193_v94).length))], _80_v11, (((_194_v95).contains(_80_v11)) ? ((_194_v95).get(_80_v11)) : ((_193_v94)[_module.__default.safeIndex(new BigNumber((_196_v97).length), new BigNumber((_193_v94).length))])));
        let _index25 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_71_v3).length));
        let _index26 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_191_v92).length));
        let _rhs14 = _dafny.Seq.update(_dafny.Seq.Concat(_75_v6, _75_v6), _module.__default.safeIndex(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_192_v93).length)), new BigNumber((_module.__default.fm23(_88_globalState)).length)), new BigNumber((_dafny.Seq.Concat(_75_v6, _75_v6)).length)), (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
        let _rhs15 = _module.__default.fm0(_68_v0, _module.__default.safeDivisionInt(new BigNumber((_192_v93).length), _68_v0), _88_globalState);
        let _rhs16 = new BigNumber(525);
        let _rhs17 = _76_v7;
        let _rhs18 = (((_197_v98).contains(_80_v11)) ? ((_197_v98).get(_80_v11)) : (_189_i8));
        let _lhs13 = _88_globalState;
        let _lhs14 = _71_v3;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_71_v3).length));
        let _lhs16 = _191_v92;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_191_v92).length));
        let _lhs18 = _88_globalState;
        _lhs13.f20 = _rhs14;
        _lhs14[_lhs15] = _rhs15;
        _68_v0 = _rhs16;
        _lhs16[_lhs17] = _rhs17;
        _lhs18.f9 = _rhs18;
      }
      let _198_v99;
      let _init3 = ((_199_v0) => function (_200_i9) {
        return _dafny.Set.fromElements(_199_v0, _199_v0);
      })(_68_v0);
      let _nw30 = Array((new BigNumber(6)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw30.length); _i0_3++) {
        _nw30[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _198_v99 = _nw30;
      let _201_v100;
      let _init4 = ((_202_v6, _203_v0) => function (_204_i10) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_202_v6).length),_203_v0);
      })(_75_v6, _68_v0);
      let _nw31 = Array((new BigNumber(25)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw31.length); _i0_4++) {
        _nw31[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _201_v100 = _nw31;
      let _205_v102;
      _205_v102 = _dafny.Map.Empty.slice().updateUnsafe(_76_v7,_68_v0);
      let _206_v103;
      _206_v103 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_76_v7,new BigNumber(-1)), _205_v102);
      let _207_v104;
      let _nw32 = new _module.C3();
      _nw32.__ctor(_198_v99, _201_v100, _module.__default.safeModuloInt(_68_v0, _68_v0), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_208_v2, _209_v6) => function (_210_i11) {
        return function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_208_v2)).Keys.Elements) {
            let _211_v101 = _compr_14;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_208_v2)).contains(_211_v101)) {
              _coll14.push([_211_v101,new BigNumber((_209_v6).length)]);
            }
          }
          return _coll14;
        }();
      })(_70_v2, _75_v6)), _206_v103));
      _207_v104 = _nw32;
      let _212_v105;
      _212_v105 = _dafny.Map.Empty.slice().updateUnsafe(_68_v0,_70_v2);
      let _213_v106;
      _213_v106 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? ((((_212_v105).contains(new BigNumber((_75_v6).length))) ? ((_212_v105).get(new BigNumber((_75_v6).length))) : (true))) : ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])),_module.__default.fm43(_68_v0, !(_module.__default.fm2(_68_v0, _68_v0, _76_v7, _88_globalState)), _70_v2, _88_globalState));
      _213_v106 = (_213_v106).Merge(_dafny.Map.Empty.slice().updateUnsafe((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))],_86_v15));
      let _214_v107;
      _214_v107 = _module.D3.create_DC9(_75_v6);
      let _source1 = _214_v107;
      if (_source1.is_DC10) {
        (_88_globalState).f2 = !(_70_v2);
        let _215_v108;
        let _216_v109;
        let _217_v110;
        let _out2;
        let _out3;
        let _out4;
        let _outcollector0 = (_207_v104).m5(_70_v2, _88_globalState);
        _out2 = _outcollector0[0];
        _out3 = _outcollector0[1];
        _out4 = _outcollector0[2];
        _215_v108 = _out2;
        _216_v109 = _out3;
        _217_v110 = _out4;
        let _218_v111;
        let _nw33 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _218_v111 = _nw33;
        _218_v111 = _218_v111;
        (_88_globalState).f1 = _module.__default.fm0(_68_v0, (_68_v0).multipliedBy(_217_v110), _88_globalState);
      } else if (_source1.is_DC11) {
        let _219___mcc_h7 = (_source1).cf18;
        let _220___mcc_h8 = (_source1).cf19;
        let _221___mcc_h9 = (_source1).cf20;
        let _222_cf20 = _221___mcc_h9;
        let _223_cf19 = _220___mcc_h8;
        let _224_cf18 = _219___mcc_h7;
        let _225_v112;
        _225_v112 = _dafny.Set.fromElements(_75_v6, _75_v6, _75_v6, _75_v6);
        let _226_v113;
        _226_v113 = _dafny.Map.Empty.slice().updateUnsafe(_223_cf19,_225_v112);
        let _227_v114;
        _227_v114 = (((_226_v113).contains(_70_v2)) ? ((_226_v113).get(_70_v2)) : (_dafny.Set.fromElements(_75_v6, _75_v6)));
        let _source2 = _227_v114;
        let _228___mcc_h11 = _source2;
        let _229_cf65 = _228___mcc_h11;
        (_88_globalState).f17 = (_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))];
        (_207_v104).m6(_88_globalState);
        let _230_v115;
        let _nw34 = new _module.C4();
        _nw34.__ctor(_71_v3, _68_v0, _206_v103);
        _230_v115 = _nw34;
        let _231_v116;
        _231_v116 = _dafny.Map.Empty.slice().updateUnsafe(_92_v20,_223_cf19);
        let _232_v117;
        _232_v117 = _dafny.Seq.of(_69_v1);
        let _233_v118;
        let _234_v119;
        let _235_v120;
        let _236_v121;
        let _out5;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector1 = (_230_v115).m4(((!(!(_223_cf19))) ? ((((_231_v116).contains(_92_v20)) ? ((_231_v116).get(_92_v20)) : (!(true)))) : ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))])), !(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_232_v117, _module.__default.safeIndex(_68_v0, new BigNumber((_232_v117).length)), _69_v1), _232_v117)), _88_globalState);
        _out5 = _outcollector1[0];
        _out6 = _outcollector1[1];
        _out7 = _outcollector1[2];
        _out8 = _outcollector1[3];
        _233_v118 = _out5;
        _234_v119 = _out6;
        _235_v120 = _out7;
        _236_v121 = _out8;
        (_207_v104).m6(_88_globalState);
        if (!(_223_cf19)) {
          let _rhs19 = _dafny.Seq.Concat(_85_v14, _85_v14);
          let _rhs20 = _68_v0;
          let _lhs19 = _88_globalState;
          let _lhs20 = _88_globalState;
          _lhs19.f17 = _rhs19;
          _lhs20.f14 = _rhs20;
          (_88_globalState).f14 = (_68_v0).multipliedBy(((_92_v20)[_module.__default.safeIndex(_224_cf18, new BigNumber((_92_v20).length))]).plus(new BigNumber((_dafny.Set.fromElements(_80_v11)).length)));
          let _237_v122;
          let _out9;
          _out9 = _module.__default.m0(new BigNumber(928), (_224_cf18).multipliedBy(new BigNumber((_92_v20).length)), _88_globalState);
          _237_v122 = _out9;
          (_88_globalState).f21 = _module.__default.fm23(_88_globalState);
          let _index27 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
          (_80_v11)[_index27] = _223_cf19;
        } else {
          let _238_v123;
          let _nw35 = Array((new BigNumber(12)).toNumber()).fill(_module.D7.Default());
          _238_v123 = _nw35;
          let _239_v124;
          _239_v124 = _dafny.Map.Empty.slice().updateUnsafe(_70_v2,_238_v123);
          _239_v124 = (_239_v124).update(_70_v2, _238_v123);
          let _240_v125;
          let _nw36 = Array((new BigNumber(6)).toNumber());
          _240_v125 = _nw36;
          _240_v125 = _240_v125;
          (_88_globalState).f14 = (new BigNumber(369)).minus(_224_cf18);
          (_88_globalState).f2 = _222_cf20;
          (_88_globalState).f18 = _68_v0;
        }
        (_88_globalState).f9 = _68_v0;
      } else {
        let _241___mcc_h10 = (_source1).cf17;
        let _242_cf17 = _241___mcc_h10;
        let _243_v126;
        let _nw37 = new _module.C4();
        _nw37.__ctor(_71_v3, new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length), _206_v103);
        _243_v126 = _nw37;
        let _244_v127;
        _244_v127 = _module.D3.create_DC10();
        let _245_v128;
        _245_v128 = _dafny.Seq.of((_dafny.MultiSet.fromElements(_module.__default.fm0(_68_v0, _68_v0, _88_globalState), new BigNumber(((_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))]).length))).update(new BigNumber((_212_v105).length), _module.__default.abs(_68_v0)));
        let _246_v129;
        let _nw38 = new _module.C1();
        _nw38.__ctor(_dafny.Seq.update(_245_v128, _module.__default.safeIndex((_207_v104).fm8(_88_globalState), new BigNumber((_245_v128).length)), _dafny.MultiSet.fromElements(new BigNumber(-838), _68_v0, _68_v0)), _68_v0, new BigNumber((_69_v1).length), _206_v103);
        _246_v129 = _nw38;
        let _247_v130;
        _247_v130 = _dafny.Map.Empty.slice().updateUnsafe((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))],_243_v126);
        let _248_v131;
        _248_v131 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_246_v129.f33, new BigNumber((_242_cf17).length), (_dafny.ZERO).minus(_68_v0)), _92_v20, _dafny.Seq.of(new BigNumber(368)));
        let _rhs21 = (_module.D3.create_DC11(_246_v129.f33, _70_v2, _70_v2)).dtor_cf20;
        let _rhs22 = (((_247_v130).contains(_70_v2)) ? ((_247_v130).get(_70_v2)) : (_243_v126));
        let _rhs23 = _module.__default.fm33(!(!(new BigNumber((_248_v131).cardinality())).isEqualTo(_68_v0)), _246_v129.f33, _88_globalState);
        let _rhs24 = _68_v0;
        let _rhs25 = _246_v129;
        let _lhs21 = _88_globalState;
        let _lhs22 = _88_globalState;
        _lhs21.f2 = _rhs21;
        _243_v126 = _rhs22;
        _244_v127 = _rhs23;
        _lhs22.f18 = _rhs24;
        _246_v129 = _rhs25;
        let _249_v132;
        let _nw39 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
        _249_v132 = _nw39;
        _249_v132 = _249_v132;
        if (_70_v2) {
          (_88_globalState).f9 = new BigNumber(-974);
          (_88_globalState).f18 = (_68_v0).multipliedBy(new BigNumber(-481));
          let _index28 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
          (_80_v11)[_index28] = (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))];
          let _250_v133;
          let _nw40 = new _module.C0();
          _nw40.__ctor(!(true));
          _250_v133 = _nw40;
          let _251_v134;
          let _nw41 = Array((_dafny.ONE).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _250_v133;
          _251_v134 = _nw41;
          let _index29 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_251_v134).length));
          (_251_v134)[_index29] = _250_v133;
          let _index30 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_251_v134).length));
          (_251_v134)[_index30] = _250_v133;
          let _252_v135;
          let _nw42 = new _module.C4();
          _nw42.__ctor(_71_v3, _module.__default.fm0(_246_v129.f33, _246_v129.f33, _88_globalState), _dafny.Seq.Concat(_206_v103, _dafny.Seq.of(_205_v102)));
          _252_v135 = _nw42;
          let _rhs26 = _252_v135;
          let _rhs27 = (_122_v45)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_122_v45).length))];
          let _rhs28 = (_dafny.ZERO).minus((new BigNumber(495)).minus(((_246_v129).fm11(_70_v2, (_250_v133).f34, (_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], _68_v0, _88_globalState)).plus((_252_v135).f24)));
          let _lhs23 = _88_globalState;
          let _lhs24 = _88_globalState;
          _252_v135 = _rhs26;
          _lhs23.f17 = _rhs27;
          _lhs24.f9 = _rhs28;
        } else {
          let _253_v136;
          let _254_v137;
          let _255_v138;
          let _out10;
          let _out11;
          let _out12;
          let _outcollector2 = (_207_v104).m5((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))], _88_globalState);
          _out10 = _outcollector2[0];
          _out11 = _outcollector2[1];
          _out12 = _outcollector2[2];
          _253_v136 = _out10;
          _254_v137 = _out11;
          _255_v138 = _out12;
          (_88_globalState).f18 = _68_v0;
          (_88_globalState).f18 = _246_v129.f33;
          (_88_globalState).f14 = (_dafny.ZERO).minus(_246_v129.f33);
          (_88_globalState).f18 = _module.__default.fm0(_68_v0, (_dafny.ZERO).minus((_246_v129.f33).plus(_246_v129.f33)), _88_globalState);
        }
        let _256_v139;
        _256_v139 = _dafny.MultiSet.fromElements(_246_v129.f33, new BigNumber((_85_v14).length));
        let _index31 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length));
        (_80_v11)[_index31] = !((_256_v139).IsProperSubsetOf(_256_v139)) || ((_80_v11)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_80_v11).length))]);
      }
      (_88_globalState).f18 = _module.__default.safeModuloInt((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements(_68_v0, new BigNumber((_92_v20).length), _68_v0)).cardinality())).plus((_dafny.ZERO).minus(_68_v0))), _68_v0);
      (_88_globalState).f18 = _68_v0;
      _212_v105 = (_212_v105).update(_68_v0, !((_126_v49).IsSubsetOf(_126_v49)));
      process.stdout.write(_dafny.toString(_68_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v1).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(590)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_70_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_73_v4).dtor_cf0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_74_v5).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_75_v6, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_77_v8).dtor_cf7, _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v9).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_79_v10).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_81_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_84_v13).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_85_v14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v15).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v15).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v15).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v15).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v15).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(590),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(382),new BigNumber(590))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_88_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_88_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_88_globalState).f3).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(590)).updateUnsafe(true,new BigNumber(590)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_88_globalState).f7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_88_globalState).f8, _dafny.Seq.of(false, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_88_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_88_globalState).f11).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_88_globalState).f13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_88_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_88_globalState).f16).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_88_globalState.f17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_88_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_88_globalState.f19).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_88_globalState.f20, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_88_globalState.f21, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState.f22).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(382),new BigNumber(590)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v18).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_92_v20, _dafny.Seq.of(new BigNumber(590), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v39).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v39).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v39).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_122_v45)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_v46).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("ncd"), _dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh")),_dafny.Seq.UnicodeFromString("ubmfhj")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_124_v47, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("nbntyekh")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_125_v48, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh"), _dafny.Seq.UnicodeFromString("nbntyekh")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v49).equals(_dafny.Set.fromElements(new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v50).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v50).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v50).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v99)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v100)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v102).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_206_v103, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber(-1)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),new BigNumber(525))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_207_v104).f29)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_207_v104.f30)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v105).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(525),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v106).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D1.create_DC4(true, true, new _dafny.CodePoint('k'.codePointAt(0)), new BigNumber(590), new BigNumber(590))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_214_v107).dtor_cf17, _dafny.Seq.of(true))));
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
    static create_DC2(cf4, cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
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
    get dtor_cf6() { return this.cf6; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
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
    static create_DC4(cf8, cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf13, cf14, cf15) {
      let $dt = new D1(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D1(2);
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(3);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + this.cf7.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13 && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC8() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf16) + ")";
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
        return other.$tag === 1 && this.cf16 === other.cf16;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8();
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
    static create_DC10() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC11(cf18, cf19, cf20) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D3(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10();
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
    static create_DC13(cf22, cf23, cf24) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D4(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D4(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + this.cf24.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.ZERO, _dafny.Map.Empty, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC15(cf26) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
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
    static create_DC18(cf29) {
      let $dt = new D6(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC16(cf27) {
      let $dt = new D6(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC19(cf30) {
      let $dt = new D6(3);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf29 === other.cf29;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false);
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
    static create_DC21(cf32, cf33, cf34) {
      let $dt = new D7(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC20(cf31) {
      let $dt = new D7(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC22(cf35) {
      let $dt = new D7(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21(_module.D1.Default(), _module.D6.Default(), []);
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
    static create_DC24(cf37) {
      let $dt = new D8(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC25(cf38, cf39, cf40) {
      let $dt = new D8(1);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf41, cf42) {
      let $dt = new D8(2);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC23(cf36) {
      let $dt = new D8(3);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf41) + ", " + this.cf42.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(_dafny.ZERO);
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
    static create_DC27(cf43) {
      let $dt = new D9(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf43) + ")";
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
      return null;
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
    static create_DC29(cf45, cf46) {
      let $dt = new D10(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D10(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf44 === other.cf44;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(_dafny.ZERO, false);
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
    static create_DC31(cf48, cf49, cf50, cf51, cf52) {
      let $dt = new D11(0);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC30(cf47) {
      let $dt = new D11(1);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(false, _dafny.ZERO, [], false, _dafny.ZERO);
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
    static create_DC33(cf54, cf55) {
      let $dt = new D12(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC32(cf53) {
      let $dt = new D12(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC34(cf56) {
      let $dt = new D12(2);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC36(cf58, cf59) {
      let $dt = new D13(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC35(cf57) {
      let $dt = new D13(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC37(cf60) {
      let $dt = new D13(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC36(false, false);
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
    static create_DC38(cf61) {
      let $dt = new D14(0);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf61) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC39(cf62) {
      let $dt = new D15(0);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62);
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf64) {
      let $dt = new D16(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC40(cf63) {
      let $dt = new D16(1);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf63 === other.cf63;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC41(_dafny.ZERO);
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
    static create_DC42(cf65) {
      let $dt = new D17(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC43(cf66) {
      let $dt = new D18(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf66 === other.cf66;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC45(cf68, cf69, cf70, cf71) {
      let $dt = new D19(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC46(cf72, cf73) {
      let $dt = new D19(1);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC47(cf74) {
      let $dt = new D19(2);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC44(cf67) {
      let $dt = new D19(3);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get is_DC44() { return this.$tag === 3; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC44" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72) && this.cf73 === other.cf73;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC45(_dafny.ZERO, false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC49(cf76, cf77) {
      let $dt = new D20(0);
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC48(cf75) {
      let $dt = new D20(1);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf75, other.cf75);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC49(_dafny.MultiSet.Empty, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this.f2 = false;
      this.f9 = _dafny.ZERO;
      this.f14 = _dafny.ZERO;
      this.f17 = _dafny.Seq.UnicodeFromString("");
      this.f18 = _dafny.ZERO;
      this.f19 = _dafny.MultiSet.Empty;
      this.f20 = _dafny.Seq.of();
      this.f21 = _dafny.Seq.of();
      this.f22 = _dafny.Map.Empty;
      this._f0 = false;
      this._f3 = _dafny.Map.Empty;
      this._f4 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f7 = _dafny.Map.Empty;
      this._f8 = _dafny.Seq.of();
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.Set.Empty;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.Map.Empty;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.Set.Empty;
      this._f23 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this).f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this).f21 = f21;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f34 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f34) {
      let _this = this;
      (_this)._f34 = f34;
      return;
    }
    fm12(p0, globalState) {
      let _this = this;
      return _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber(-780))).plus(new BigNumber(-448)), new BigNumber((_dafny.Seq.UnicodeFromString("a")).length));
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
      this.f33 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f32, f33, f24, f25) {
      let _this = this;
      (_this)._f32 = f32;
      (_this).f33 = f33;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm10(globalState) {
      let _this = this;
      return false;
    };
    fm11(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("qgrpl")).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))))).length));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      if (((_this.f33).minus(_this.f33)).isLessThanOrEqualTo(_this.f33)) {
        let _257_v0;
        _257_v0 = true;
        (globalState).f2 = _257_v0;
        if (_257_v0) {
          (globalState).f1 = (_this).f24;
          let _258_v1;
          _258_v1 = new _dafny.CodePoint('i'.codePointAt(0));
          _258_v1 = _258_v1;
          let _259_v2;
          let _init5 = ((_260_v0) => function (_261_i0) {
            return _260_v0;
          })(_257_v0);
          let _nw43 = Array((new BigNumber(9)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
            _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _259_v2 = _nw43;
          let _262_v3;
          _262_v3 = _dafny.Seq.of(_257_v0);
          let _263_v4;
          _263_v4 = _dafny.Set.fromElements(_257_v0, (_262_v3)[_module.__default.safeIndex(new BigNumber(730), new BigNumber((_262_v3).length))], _257_v0);
          _259_v2 = (((_dafny.Set.fromElements(_257_v0)).IsSubsetOf(_263_v4)) ? (_259_v2) : (_259_v2));
          let _264_v5;
          _264_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,_257_v0);
          let _265_v6;
          _265_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_262_v3),(_this).fm10(globalState));
          let _266_v7;
          _266_v7 = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,_257_v0);
          let _267_v8;
          _267_v8 = _module.D3.create_DC9(_262_v3);
          let _268_v9;
          _268_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_266_v7).length),(_267_v8).dtor_cf17);
          (globalState).f2 = (((_264_v5).contains(_this.f33)) ? ((_264_v5).get(_this.f33)) : ((((_265_v6).contains(_dafny.Seq.of((((_268_v9).contains((_dafny.ZERO).minus((_this).f24))) ? ((_268_v9).get((_dafny.ZERO).minus((_this).f24))) : (_262_v3))))) ? ((_265_v6).get(_dafny.Seq.of((((_268_v9).contains((_dafny.ZERO).minus((_this).f24))) ? ((_268_v9).get((_dafny.ZERO).minus((_this).f24))) : (_262_v3))))) : (_257_v0))));
          let _269_v10;
          _269_v10 = _dafny.Seq.UnicodeFromString("clyji");
          let _270_v11;
          _270_v11 = _dafny.Set.fromElements(_269_v10);
          let _271_v12;
          _271_v12 = _dafny.Seq.of(_module.__default.safeDivisionInt((_this).f24, new BigNumber((_270_v11).length)), (_this).f24, new BigNumber((_dafny.MultiSet.fromElements(_257_v0)).cardinality()));
          let _rhs29 = _259_v2;
          let _rhs30 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), ((_272_v1) => function (_273_i1) {
            return _272_v1;
          })(_258_v1)))).length);
          let _rhs31 = (_this).f24;
          let _rhs32 = _271_v12;
          let _lhs25 = globalState;
          let _lhs26 = globalState;
          _259_v2 = _rhs29;
          _lhs25.f9 = _rhs30;
          _lhs26.f9 = _rhs31;
          _271_v12 = _rhs32;
        } else {
          let _274_v13;
          _274_v13 = _dafny.MultiSet.fromElements(_this.f33);
          let _275_v14;
          _275_v14 = _dafny.MultiSet.fromElements(_257_v0, _257_v0, _257_v0);
          let _276_v15;
          let _nw44 = new _module.C0();
          _nw44.__ctor(!(new BigNumber(((_274_v13).update(_this.f33, _module.__default.abs((_this).f24))).cardinality())).isEqualTo((((_275_v14).contains(_257_v0)) ? ((_275_v14).get(_257_v0)) : (new BigNumber((_dafny.Set.fromElements(true, !(_257_v0), _257_v0)).length)))));
          _276_v15 = _nw44;
          let _277_v16;
          let _init6 = function (_278_i2) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), function (_279_i3) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_280_i4) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            }));
          };
          let _nw45 = Array((new BigNumber(17)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
            _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _277_v16 = _nw45;
          let _281_v17;
          _281_v17 = _dafny.Seq.UnicodeFromString("wtih");
          let _282_v18;
          _282_v18 = new _dafny.CodePoint('q'.codePointAt(0));
          let _index32 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_277_v16).length));
          (_277_v16)[_index32] = _dafny.Seq.update(_281_v17, _module.__default.safeIndex((_this).fm11((_276_v15).f34, (_276_v15).f34, true, _this.f33, globalState), new BigNumber((_281_v17).length)), _282_v18);
          let _index33 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_277_v16).length));
          (_277_v16)[_index33] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vcjpd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), ((_283_v18) => function (_284_i5) {
            return _283_v18;
          })(_282_v18)));
          _276_v15 = _276_v15;
          let _285_v19;
          _285_v19 = _dafny.Map.Empty.slice().updateUnsafe((_276_v15).f34,_282_v18);
          let _286_v20;
          _286_v20 = _dafny.Map.Empty.slice().updateUnsafe(!((_276_v15).f34),_257_v0);
          _285_v19 = (_285_v19).update((((_286_v20).contains((_276_v15).f34)) ? ((_286_v20).get((_276_v15).f34)) : (_257_v0)), _282_v18);
          let _287_v21;
          _287_v21 = _module.D0.create_DC1((_this).f24, new BigNumber(949), ((_module.__default.fm2(_this.f33, (_dafny.ZERO).minus((_this).f24), _282_v18, globalState)) ? ((_276_v15).f34) : ((_276_v15).f34)));
          let _288_v22;
          _288_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(_257_v0),_this.f33);
          let _289_v23;
          _289_v23 = _dafny.Seq.of((_276_v15).f34, (_276_v15).f34, _257_v0);
          let _290_v24;
          _290_v24 = _dafny.Seq.of(((_257_v0) ? (_288_v22) : (_288_v22)), (_288_v22).update(true, new BigNumber((_289_v23).length)), _288_v22, _288_v22, _288_v22);
          let _rhs33 = (_276_v15).f34;
          let _rhs34 = _287_v21;
          let _rhs35 = (_this).f24;
          let _rhs36 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_291_v22) => function (_292_i6) {
            return _291_v22;
          })(_288_v22));
          let _lhs27 = globalState;
          let _lhs28 = globalState;
          _lhs27.f2 = _rhs33;
          _287_v21 = _rhs34;
          _lhs28.f1 = _rhs35;
          _290_v24 = _rhs36;
        }
        let _293_v25;
        let _nw46 = Array((new BigNumber(12)).toNumber()).fill([]);
        _293_v25 = _nw46;
        let _294_v26;
        _294_v26 = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,_257_v0);
        let _295_v27;
        _295_v27 = _dafny.Seq.of(_294_v26);
        let _296_v28;
        let _nw47 = Array((new BigNumber(25)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _module.__default.fm13(globalState);
        _nw47[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,_257_v0);
        _nw47[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,_257_v0);
        _nw47[(new BigNumber(3)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(4)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,_257_v0);
        _nw47[(new BigNumber(6)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(7)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(8)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,!(_257_v0));
        _nw47[(new BigNumber(10)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_257_v0);
        _nw47[(new BigNumber(12)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(13)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(14)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(15)).toNumber()] = (_295_v27)[_module.__default.safeIndex((_this).f24, new BigNumber((_295_v27).length))];
        _nw47[(new BigNumber(16)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(17)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(18)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(19)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(20)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(21)).toNumber()] = (_294_v26).update(_257_v0, _257_v0);
        _nw47[(new BigNumber(22)).toNumber()] = _294_v26;
        _nw47[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_257_v0,!(_257_v0));
        _nw47[(new BigNumber(24)).toNumber()] = _294_v26;
        _296_v28 = _nw47;
        let _index34 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_293_v25).length));
        (_293_v25)[_index34] = _296_v28;
        let _index35 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_293_v25).length));
        (_293_v25)[_index35] = _296_v28;
        let _297_v29;
        _297_v29 = _dafny.Set.fromElements(new BigNumber(342));
        _297_v29 = _297_v29;
        let _298_v30;
        let _nw48 = Array((new BigNumber(14)).toNumber()).fill(false);
        _298_v30 = _nw48;
        let _299_v31;
        _299_v31 = _dafny.Seq.of(_298_v30, _298_v30, _298_v30);
        let _300_v32;
        let _301_v33;
        let _302_v34;
        let _303_v35;
        let _out13;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector3 = (_this).m8((_299_v31)[_module.__default.safeIndex(_this.f33, new BigNumber((_299_v31).length))], globalState);
        _out13 = _outcollector3[0];
        _out14 = _outcollector3[1];
        _out15 = _outcollector3[2];
        _out16 = _outcollector3[3];
        _300_v32 = _out13;
        _301_v33 = _out14;
        _302_v34 = _out15;
        _303_v35 = _out16;
      } else {
        let _304_v36;
        _304_v36 = true;
        if (_304_v36) {
          (globalState).f2 = _304_v36;
          let _305_v37;
          _305_v37 = _dafny.Map.Empty.slice().updateUnsafe(_304_v36,_this.f33);
          let _306_v38;
          _306_v38 = _module.D0.create_DC1(_this.f33, new BigNumber((_305_v37).length), _304_v36);
          let _307_v39;
          _307_v39 = _dafny.Map.Empty.slice().updateUnsafe((_306_v38).dtor_cf3,_this.f33);
          let _308_v40;
          _308_v40 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_304_v36);
          let _309_v41;
          _309_v41 = new _dafny.CodePoint('c'.codePointAt(0));
          let _310_v42;
          _310_v42 = _dafny.MultiSet.fromElements(_304_v36);
          let _311_v43;
          _311_v43 = _module.D1.create_DC4((_this.f33).isLessThanOrEqualTo(new BigNumber((_307_v39).length)), (((_308_v40).contains(new _dafny.CodePoint('d'.codePointAt(0)))) ? ((_308_v40).get(new _dafny.CodePoint('d'.codePointAt(0)))) : (_304_v36)), _309_v41, (((_310_v42).contains(true)) ? ((_310_v42).get(true)) : (_this.f33)), (_this).fm11(_304_v36, !(_304_v36), _304_v36, _this.f33, globalState));
          _311_v43 = _311_v43;
          let _312_v44;
          let _nw49 = new _module.C0();
          _nw49.__ctor(_304_v36);
          _312_v44 = _nw49;
          _310_v42 = (_310_v42).Intersect(_module.__default.fm14(_module.__default.fm15(new BigNumber(224), _304_v36, (_this).f24, (_312_v44).f34, globalState), _this.f33, globalState));
          let _313_v45;
          _313_v45 = _dafny.Seq.of(_this.f33, _this.f33);
          let _314_v46;
          _314_v46 = _module.D4.create_DC12(_313_v45);
          let _315_v47;
          _315_v47 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(((_310_v42).update((_312_v44).f34, _module.__default.abs(_this.f33))).cardinality())).isLessThan(_this.f33),(_314_v46).dtor_cf21);
          _315_v47 = (_315_v47).update((_312_v44).f34, _dafny.Seq.update(_dafny.Seq.of(_this.f33, (_this).f24), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_this.f33)).length), new BigNumber((_dafny.Seq.of(_this.f33, (_this).f24)).length)), (_this).f24));
        } else {
          let _316_v48;
          _316_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,((_304_v36) ? ((_dafny.ZERO).minus((_this).f24)) : ((_this).f24)));
          _316_v48 = (_316_v48).update((_this).f24, (_this).f24);
          let _317_v49;
          let _nw50 = Array((new BigNumber(24)).toNumber()).fill(_module.D0.Default());
          _317_v49 = _nw50;
          _317_v49 = _317_v49;
          let _318_v50;
          _318_v50 = _module.D0.create_DC1((_dafny.ZERO).minus((_this).f24), new BigNumber(300), _304_v36);
          (globalState).f2 = !(((_this).f24).isEqualTo((_this.f33).plus((_318_v50).dtor_cf2)));
          let _319_v51;
          _319_v51 = _dafny.Seq.UnicodeFromString("i");
          (globalState).f17 = _319_v51;
          let _320_v52;
          _320_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), function (_321_i7) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }));
          _320_v52 = (_320_v52).update((_this).f24, _319_v51);
        }
        let _322_v53;
        let _init7 = function (_323_i8) {
          return (_323_i8).multipliedBy(_this.f33);
        };
        let _nw51 = Array((new BigNumber(20)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw51.length); _i0_7++) {
          _nw51[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _322_v53 = _nw51;
        let _index36 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_322_v53).length));
        (_322_v53)[_index36] = (_this).f24;
        let _index37 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_322_v53).length));
        (_322_v53)[_index37] = ((_this).f24).multipliedBy((_this).f24);
        (globalState).f14 = _this.f33;
        let _324_v54;
        let _init8 = function (_325_i9) {
          return _dafny.Seq.UnicodeFromString("u");
        };
        let _nw52 = Array((new BigNumber(6)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw52.length); _i0_8++) {
          _nw52[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _324_v54 = _nw52;
        let _326_v55;
        _326_v55 = _dafny.Seq.UnicodeFromString("xwklb");
        let _index38 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_324_v54).length));
        (_324_v54)[_index38] = _326_v55;
        let _index39 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_324_v54).length));
        (_324_v54)[_index39] = _326_v55;
        let _327_v56;
        _327_v56 = _dafny.Map.Empty.slice().updateUnsafe(_304_v36,_304_v36);
        let _328_v57;
        _328_v57 = _dafny.Map.Empty.slice().updateUnsafe(_304_v36,new BigNumber((_327_v56).length));
        let _rhs37 = (_module.D4.create_DC13(_this.f33, _328_v57, _326_v55)).dtor_cf22;
        let _rhs38 = (_this).f24;
        let _rhs39 = _304_v36;
        let _rhs40 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_329_i10) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        });
        let _lhs29 = globalState;
        let _lhs30 = globalState;
        let _lhs31 = globalState;
        _lhs29.f18 = _rhs37;
        _lhs30.f9 = _rhs38;
        _304_v36 = _rhs39;
        _lhs31.f17 = _rhs40;
      }
      (globalState).f9 = (_this).f24;
      let _330_v58;
      _330_v58 = true;
      r0 = _330_v58;
      _330_v58 = (_this).fm10(globalState);
      let _331_v59;
      _331_v59 = _dafny.Seq.UnicodeFromString("tvcjan");
      let _332_v60;
      _332_v60 = new _dafny.CodePoint('t'.codePointAt(0));
      let _333_v61;
      _333_v61 = _dafny.Seq.of(_331_v59, _331_v59, _module.__default.fm4(_330_v58, _this.f33, _332_v60, !(_330_v58), globalState));
      let _334_i11;
      _334_i11 = _dafny.ZERO;
      L0: {
        while (!_dafny.areEqual(_333_v61, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-168)), ((_376_v59) => function (_377_i12) {
          return _376_v59;
        })(_331_v59)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_334_i11)) {
              break L0;
            }
            _334_i11 = (_334_i11).plus(_dafny.ONE);
            let _335_v62;
            _335_v62 = _dafny.Set.fromElements(_this.f33);
            let _336_v63;
            _336_v63 = _dafny.Map.Empty.slice().updateUnsafe(_330_v58,new BigNumber((_335_v62).length));
            let _337_v64;
            _337_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,_332_v60);
            let _338_v65;
            _338_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_337_v64).length),!(_330_v58));
            let _339_v66;
            let _nw53 = Array((new BigNumber(19)).toNumber());
            _nw53[(_dafny.ZERO).toNumber()] = _this.f33;
            _nw53[(_dafny.ONE).toNumber()] = _this.f33;
            _nw53[(new BigNumber(2)).toNumber()] = new BigNumber((_336_v63).length);
            _nw53[(new BigNumber(3)).toNumber()] = (_this).f24;
            _nw53[(new BigNumber(4)).toNumber()] = new BigNumber(676);
            _nw53[(new BigNumber(5)).toNumber()] = new BigNumber((_331_v59).length);
            _nw53[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm16(globalState)).length);
            _nw53[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f24);
            _nw53[(new BigNumber(8)).toNumber()] = (_this).f24;
            _nw53[(new BigNumber(9)).toNumber()] = new BigNumber((_338_v65).length);
            _nw53[(new BigNumber(10)).toNumber()] = (_this).f24;
            _nw53[(new BigNumber(11)).toNumber()] = ((_this).f24).minus(new BigNumber((_dafny.Seq.of(_331_v59, _331_v59)).length));
            _nw53[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_this.f33, new BigNumber(-6));
            _nw53[(new BigNumber(13)).toNumber()] = (_this).f24;
            _nw53[(new BigNumber(14)).toNumber()] = (_this).f24;
            _nw53[(new BigNumber(15)).toNumber()] = new BigNumber((_338_v65).length);
            _nw53[(new BigNumber(16)).toNumber()] = _this.f33;
            _nw53[(new BigNumber(17)).toNumber()] = _module.__default.fm0(new BigNumber(-523), _this.f33, globalState);
            _nw53[(new BigNumber(18)).toNumber()] = (_this).f24;
            _339_v66 = _nw53;
            let _index40 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length));
            (_339_v66)[_index40] = _this.f33;
            let _index41 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length));
            let _rhs41 = (_this).f24;
            let _rhs42 = ((_330_v58) ? (_339_v66) : (_339_v66));
            let _lhs32 = _339_v66;
            let _lhs33 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length));
            _lhs32[_lhs33] = _rhs41;
            _339_v66 = _rhs42;
            let _340_v67;
            _340_v67 = _module.D0.create_DC2(_330_v58, new BigNumber(-118), _332_v60);
            let _341_v68;
            _341_v68 = _dafny.Map.Empty.slice().updateUnsafe(_340_v67,_330_v58);
            let _342_v69;
            _342_v69 = _dafny.Map.Empty.slice().updateUnsafe(_330_v58,((_330_v58) ? (_330_v58) : ((((_341_v68).contains(_340_v67)) ? ((_341_v68).get(_340_v67)) : (_330_v58)))));
            _342_v69 = (_342_v69).update(!((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]).isEqualTo((_this).f24), _330_v58);
            let _343_v70;
            _343_v70 = _module.D2.create_DC8();
            let _source3 = _343_v70;
            if (_source3.is_DC8) {
              let _344_v71;
              _344_v71 = _dafny.Seq.of(new BigNumber((_331_v59).length));
              let _345_v72;
              _345_v72 = _dafny.Seq.of(_330_v58, _330_v58);
              let _346_v73;
              let _nw54 = Array((new BigNumber(27)).toNumber());
              _nw54[(_dafny.ZERO).toNumber()] = _module.__default.fm17(new BigNumber(790), (_this).f24, new BigNumber(580), false, globalState);
              _nw54[(_dafny.ONE).toNumber()] = _336_v63;
              _nw54[(new BigNumber(2)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(3)).toNumber()] = (_336_v63).update(_330_v58, (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]);
              _nw54[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], _332_v60, globalState),(_this).f24);
              _nw54[(new BigNumber(5)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_330_v58,new BigNumber(641))).update(_330_v58, _this.f33);
              _nw54[(new BigNumber(6)).toNumber()] = _module.__default.fm17(new BigNumber((_module.__default.fm4(_330_v58, (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], _332_v60, _330_v58, globalState)).length), _this.f33, new BigNumber((_344_v71).length), _330_v58, globalState);
              _nw54[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_330_v58,new BigNumber(-683));
              _nw54[(new BigNumber(8)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_330_v58,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_330_v58,(_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]),new BigNumber((_331_v59).length))).length));
              _nw54[(new BigNumber(10)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(11)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(12)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(13)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(14)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_330_v58,_this.f33);
              _nw54[(new BigNumber(16)).toNumber()] = _module.__default.fm17(new BigNumber((_module.__default.fm18(globalState)).length), (_this).f24, new BigNumber((_345_v72).length), _330_v58, globalState);
              _nw54[(new BigNumber(17)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(18)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(19)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(20)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(21)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(22)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(23)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(24)).toNumber()] = (_336_v63).update(_330_v58, (_this).f24);
              _nw54[(new BigNumber(25)).toNumber()] = _336_v63;
              _nw54[(new BigNumber(26)).toNumber()] = _336_v63;
              _346_v73 = _nw54;
              let _347_v74;
              _347_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_348_v60) => function (_349_i13) {
                return _348_v60;
              })(_332_v60))).length),_346_v73);
              let _350_v75;
              _350_v75 = _dafny.Seq.of(_346_v73);
              _347_v74 = (_347_v74).update((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], (_350_v75)[_module.__default.safeIndex(_this.f33, new BigNumber((_350_v75).length))]);
              let _351_v76;
              _351_v76 = _module.D3.create_DC11(_this.f33, _330_v58, _330_v58);
              let _352_v77;
              _352_v77 = _dafny.MultiSet.fromElements(_this.f33, (_this).f24, _this.f33, (_this).fm11(_330_v58, !(_330_v58), _330_v58, _this.f33, globalState), (_351_v76).dtor_cf18);
              let _353_v78;
              _353_v78 = _dafny.Set.fromElements(_339_v66);
              let _rhs43 = _330_v58;
              let _rhs44 = _352_v77;
              let _rhs45 = (_353_v78).Intersect(_353_v78);
              let _lhs34 = globalState;
              _lhs34.f2 = _rhs43;
              _352_v77 = _rhs44;
              _353_v78 = _rhs45;
              let _354_v79;
              _354_v79 = _dafny.MultiSet.fromElements(_330_v58, _330_v58);
              let _355_v80;
              _355_v80 = _dafny.Map.Empty.slice().updateUnsafe(_354_v79,((_this).f24).multipliedBy((_this).f24));
              _355_v80 = _355_v80;
              let _356_v81;
              let _init9 = ((_357_v58) => function (_358_i14) {
                return _357_v58;
              })(_330_v58);
              let _nw55 = Array((new BigNumber(24)).toNumber());
              for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw55.length); _i0_9++) {
                _nw55[_i0_9] = _init9(new BigNumber(_i0_9));
              }
              _356_v81 = _nw55;
              let _359_v82;
              _359_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_330_v58, (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], globalState),_356_v81);
              let _360_v83;
              _360_v83 = _dafny.Seq.of(_dafny.Seq.of(_330_v58));
              _356_v81 = (((_359_v82).contains(_dafny.Seq.Concat(_360_v83, _dafny.Seq.update(_360_v83, _module.__default.safeIndex(new BigNumber((_331_v59).length), new BigNumber((_360_v83).length)), _dafny.Seq.of(true))))) ? ((_359_v82).get(_dafny.Seq.Concat(_360_v83, _dafny.Seq.update(_360_v83, _module.__default.safeIndex(new BigNumber((_331_v59).length), new BigNumber((_360_v83).length)), _dafny.Seq.of(true))))) : (_356_v81));
            } else {
              let _361___mcc_h0 = (_source3).cf16;
              let _362_cf16 = _361___mcc_h0;
              (globalState).f2 = !(_module.__default.safeModuloInt(new BigNumber((_338_v65).length), (_dafny.ZERO).minus((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]))).isEqualTo((_this).f24);
              let _363_v84;
              _363_v84 = _dafny.MultiSet.fromElements(_330_v58);
              let _364_v85;
              _364_v85 = _dafny.Seq.of(new BigNumber((_363_v84).cardinality()));
              let _365_v86;
              _365_v86 = _dafny.Seq.of(_330_v58, _330_v58, _330_v58);
              let _366_v87;
              _366_v87 = _dafny.Map.Empty.slice().updateUnsafe(_365_v86,_330_v58);
              let _367_v88;
              _367_v88 = _dafny.Map.Empty.slice().updateUnsafe(_332_v60,_330_v58);
              let _368_v89;
              _368_v89 = _dafny.MultiSet.fromElements(_330_v58);
              let _369_v90;
              _369_v90 = _dafny.Seq.of(_363_v84, (_368_v89));
              let _370_v91;
              _370_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,(_this).f24);
              let _371_v92;
              _371_v92 = _module.D1.create_DC4(_330_v58, _330_v58, _332_v60, new BigNumber((_370_v91).length), (_this).f24);
              let _372_v94;
              let _nw56 = Array((new BigNumber(24)).toNumber());
              _nw56[(_dafny.ZERO).toNumber()] = !((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]).isEqualTo(_this.f33);
              _nw56[(_dafny.ONE).toNumber()] = true;
              _nw56[(new BigNumber(2)).toNumber()] = _module.__default.fm2((_364_v85)[_module.__default.safeIndex(new BigNumber((_331_v59).length), new BigNumber((_364_v85).length))], _this.f33, _332_v60, globalState);
              _nw56[(new BigNumber(3)).toNumber()] = _dafny.Seq.contains(_364_v85, (_this).f24);
              _nw56[(new BigNumber(4)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(5)).toNumber()] = _module.__default.fm2(_this.f33, _this.f33, new _dafny.CodePoint('c'.codePointAt(0)), globalState);
              _nw56[(new BigNumber(6)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(7)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(8)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(9)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(10)).toNumber()] = !(true);
              _nw56[(new BigNumber(11)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(12)).toNumber()] = (((_366_v87).contains(_365_v86)) ? ((_366_v87).get(_365_v86)) : (_330_v58));
              _nw56[(new BigNumber(13)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(14)).toNumber()] = !(_module.__default.fm2((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], new BigNumber((_363_v84).cardinality()), _332_v60, globalState));
              _nw56[(new BigNumber(15)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(16)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(17)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(18)).toNumber()] = (new BigNumber((_367_v88).length)).isLessThanOrEqualTo(new BigNumber(((_369_v90)[_module.__default.safeIndex((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))], new BigNumber((_369_v90).length))]).cardinality()));
              _nw56[(new BigNumber(19)).toNumber()] = _module.__default.fm2(new BigNumber((_370_v91).length), (_371_v92).dtor_cf12, _332_v60, globalState);
              _nw56[(new BigNumber(20)).toNumber()] = false;
              _nw56[(new BigNumber(21)).toNumber()] = ((_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]).isLessThanOrEqualTo(new BigNumber(((function () {
                let _coll15 = new _dafny.Map();
                for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-804), new BigNumber(751))) {
                  let _373_v93 = _compr_15;
                  if (((new BigNumber(-804)).isLessThanOrEqualTo(_373_v93)) && ((_373_v93).isLessThan(new BigNumber(751)))) {
                    _coll15.push([_module.__default.safeDivisionInt(_373_v93, (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))]),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_330_v58,_332_v60)).length)]);
                  }
                }
                return _coll15;
              }()).update((_this).f24, (_339_v66)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_339_v66).length))])).length));
              _nw56[(new BigNumber(22)).toNumber()] = _330_v58;
              _nw56[(new BigNumber(23)).toNumber()] = _330_v58;
              _372_v94 = _nw56;
              _372_v94 = _372_v94;
              let _374_v95;
              _374_v95 = _dafny.Set.fromElements(_365_v86, _365_v86);
              _374_v95 = ((function () {
                let _coll16 = new _dafny.Set();
                for (const _compr_16 of (_module.__default.fm20(globalState)).Keys.Elements) {
                  let _375_v96 = _compr_16;
                  if ((_module.__default.fm20(globalState)).contains(_375_v96)) {
                    _coll16.add(_375_v96);
                  }
                }
                return _coll16;
              }()).Difference(_374_v95)).Difference(_374_v95);
              (globalState).f17 = _331_v59;
            }
            (globalState).f1 = (new BigNumber((_335_v62).length)).multipliedBy(_this.f33);
          }
        }
      }
      (globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_331_v59, _331_v59), _dafny.Seq.UnicodeFromString("ndbanqk"))).length);
      r0 = !(!(_330_v58));
      return r0;
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      (globalState).f18 = _this.f33;
      let _378_v0;
      _378_v0 = true;
      let _379_v1;
      _379_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_378_v0);
      let _380_v2;
      _380_v2 = _dafny.Seq.of(new BigNumber((_379_v1).length));
      let _381_v3;
      _381_v3 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_380_v2, _380_v2),(_this).f24);
      (globalState).f1 = (((_381_v3).contains(_380_v2)) ? ((_381_v3).get(_380_v2)) : ((_this).f24));
      if (_378_v0) {
        r0 = _378_v0;
        let _382_v4;
        _382_v4 = new _dafny.CodePoint('k'.codePointAt(0));
        let _383_v5;
        _383_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(546),p0);
        let _384_v6;
        _384_v6 = _module.D1.create_DC4(_378_v0, _378_v0, _382_v4, (_this).f24, new BigNumber((_383_v5).length));
        let _pat_let_tv0 = _382_v4;
        let _source4 = function (_pat_let0_0) {
          return function (_385_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_386_dt__update_hcf10_h0) {
                return _module.D1.create_DC4((_385_dt__update__tmp_h0).dtor_cf8, (_385_dt__update__tmp_h0).dtor_cf9, _386_dt__update_hcf10_h0, (_385_dt__update__tmp_h0).dtor_cf11, (_385_dt__update__tmp_h0).dtor_cf12);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_384_v6);
        if (_source4.is_DC4) {
          let _387___mcc_h0 = (_source4).cf8;
          let _388___mcc_h1 = (_source4).cf9;
          let _389___mcc_h2 = (_source4).cf10;
          let _390___mcc_h3 = (_source4).cf11;
          let _391___mcc_h4 = (_source4).cf12;
          let _392_cf12 = _391___mcc_h4;
          let _393_cf11 = _390___mcc_h3;
          let _394_cf10 = _389___mcc_h2;
          let _395_cf9 = _388___mcc_h1;
          let _396_cf8 = _387___mcc_h0;
          let _397_v7;
          let _nw57 = Array((new BigNumber(4)).toNumber()).fill([]);
          _397_v7 = _nw57;
          _397_v7 = _397_v7;
          (globalState).f2 = _378_v0;
          let _398_v8;
          _398_v8 = _module.D2.create_DC8();
          let _399_v9;
          _399_v9 = _dafny.Set.fromElements(_378_v0);
          let _400_v10;
          _400_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-689)), ((_401_v4) => function (_402_i0) {
            return _401_v4;
          })(_382_v4)),true);
          let _403_v11;
          _403_v11 = _dafny.Seq.UnicodeFromString("qojo");
          let _404_v12;
          _404_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_378_v0),_396_cf8);
          let _405_v13;
          _405_v13 = _dafny.Seq.of(_404_v12);
          let _rhs46 = _398_v8;
          let _rhs47 = (((_399_v9).IsSubsetOf(_399_v9)) ? (_396_cf8) : (((_396_cf8) ? (_378_v0) : (_395_cf9))));
          let _rhs48 = (_dafny.ZERO).minus((new BigNumber(((_400_v10).update(_dafny.Seq.update(_403_v11, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_380_v2),_this.f33)).length), new BigNumber((_403_v11).length)), _382_v4), _395_cf9)).length)).multipliedBy(((_this).f24).plus(new BigNumber((_405_v13).length))));
          let _lhs35 = globalState;
          _398_v8 = _rhs46;
          _396_cf8 = _rhs47;
          _lhs35.f1 = _rhs48;
          let _406_v14;
          _406_v14 = _dafny.Map.Empty.slice().updateUnsafe(_394_cf10,!(_395_cf9) || (_395_cf9));
          _406_v14 = (_406_v14).update(_382_v4, _378_v0);
        } else if (_source4.is_DC5) {
          let _407___mcc_h5 = (_source4).cf13;
          let _408___mcc_h6 = (_source4).cf14;
          let _409___mcc_h7 = (_source4).cf15;
          let _410_cf15 = _409___mcc_h7;
          let _411_cf14 = _408___mcc_h6;
          let _412_cf13 = _407___mcc_h5;
          let _413_v15;
          let _nw58 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
          _413_v15 = _nw58;
          let _414_v16;
          _414_v16 = _dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(_this.f33, new BigNumber((p0).length))]);
          let _index42 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_413_v15).length));
          (_413_v15)[_index42] = ((_378_v0) ? (_414_v16) : (_414_v16));
          let _415_v17;
          _415_v17 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_414_v16);
          let _index43 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_413_v15).length));
          (_413_v15)[_index43] = (((_415_v17).contains(false)) ? ((_415_v17).get(false)) : (_dafny.MultiSet.FromArray(p0)));
          let _416_v18;
          _416_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),(_this).f24);
          _416_v18 = (_416_v18).update(_this.f33, _this.f33);
          let _417_v19;
          let _init10 = ((_418_v0) => function (_419_i1) {
            return _418_v0;
          })(_378_v0);
          let _nw59 = Array((new BigNumber(11)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw59.length); _i0_10++) {
            _nw59[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _417_v19 = _nw59;
          let _420_v20;
          _420_v20 = _dafny.MultiSet.fromElements(_417_v19);
          let _421_v21;
          _421_v21 = _dafny.Set.fromElements(_378_v0);
          let _422_v22;
          _422_v22 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,new BigNumber((_421_v21).length));
          let _423_v23;
          _423_v23 = _dafny.Seq.UnicodeFromString("lxpkr");
          let _424_v24;
          let _out17;
          _out17 = _module.__default.m0(new BigNumber((((_411_cf14) ? (_dafny.MultiSet.fromElements(_417_v19, _417_v19)) : (_420_v20))).cardinality()), (_module.D4.create_DC13((_dafny.ZERO).minus((_this).f24), _422_v22, _423_v23)).dtor_cf22, globalState);
          _424_v24 = _out17;
          let _425_v25;
          _425_v25 = _module.D3.create_DC11((_this).f24, _410_cf15, _412_cf13);
          (_this).f33 = ((_425_v25).dtor_cf18).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f33)));
        } else if (_source4.is_DC6) {
          let _426_v26;
          let _init11 = ((_427_v0) => function (_428_i2) {
            return _427_v0;
          })(_378_v0);
          let _nw60 = Array((new BigNumber(21)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw60.length); _i0_11++) {
            _nw60[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _426_v26 = _nw60;
          let _429_v27;
          _429_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,_426_v26);
          let _430_v28;
          _430_v28 = _dafny.Seq.UnicodeFromString("lp");
          let _431_v29;
          _431_v29 = _dafny.Map.Empty.slice().updateUnsafe((_429_v27).Merge(_429_v27),_430_v28);
          let _432_v30;
          _432_v30 = _dafny.Seq.of(_430_v28);
          _431_v29 = (_431_v29).update(_429_v27, _dafny.Seq.Concat((_432_v30)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_380_v2).length)), new BigNumber((_432_v30).length))], _dafny.Seq.UnicodeFromString("ucibb")));
          let _433_v31;
          _433_v31 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_378_v0);
          _433_v31 = (_433_v31).update((_378_v0) && (false), _378_v0);
          let _434_v32;
          let _nw61 = new _module.C0();
          _nw61.__ctor(true);
          _434_v32 = _nw61;
          let _435_v33;
          let _nw62 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
          _435_v33 = _nw62;
          _435_v33 = _435_v33;
        } else {
          let _436___mcc_h8 = (_source4).cf7;
          let _437_cf7 = _436___mcc_h8;
          let _438_v34;
          _438_v34 = _dafny.MultiSet.fromElements(_378_v0, _378_v0);
          r0 = (_438_v34).equals((_438_v34).Difference(_438_v34));
          let _439_v35;
          _439_v35 = _dafny.Map.Empty.slice().updateUnsafe((_378_v0) === (_378_v0),((_this).f24).plus(new BigNumber((_437_cf7).length)));
          let _440_v36;
          let _nw63 = Array((new BigNumber(29)).toNumber()).fill(false);
          _440_v36 = _nw63;
          let _rhs49 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("pfkfe"), _module.__default.safeIndex(_this.f33, new BigNumber((_dafny.Seq.UnicodeFromString("pfkfe")).length)), _382_v4), _437_cf7)).length));
          let _rhs50 = _dafny.MultiSet.fromElements(_440_v36);
          let _rhs51 = (_this).f24;
          let _lhs36 = globalState;
          let _lhs37 = globalState;
          _439_v35 = _rhs49;
          _lhs36.f19 = _rhs50;
          _lhs37.f1 = _rhs51;
          (globalState).f9 = _this.f33;
          (globalState).f2 = false;
        }
        let _441_v37;
        let _nw64 = Array((new BigNumber(3)).toNumber()).fill(false);
        _441_v37 = _nw64;
        let _index44 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length));
        (_441_v37)[_index44] = _378_v0;
        let _index45 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length));
        (_441_v37)[_index45] = _module.__default.fm2((_this).f24, _this.f33, new _dafny.CodePoint('w'.codePointAt(0)), globalState);
        (globalState).f18 = (_this).f24;
        let _442_v38;
        _442_v38 = _module.D1.create_DC5((_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))], _378_v0, !(_378_v0));
        if ((_442_v38).dtor_cf13) {
          let _443_v39;
          _443_v39 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_this.f33);
          let _444_v40;
          _444_v40 = _module.D4.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("pao")).length), _443_v39, _dafny.Seq.UnicodeFromString("yvdcwh"));
          let _index46 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length));
          (_441_v37)[_index46] = (_module.__default.fm0(_this.f33, new BigNumber((_dafny.Seq.update((_444_v40).dtor_cf24, _module.__default.safeIndex((_this).f24, new BigNumber(((_444_v40).dtor_cf24).length)), _382_v4)).length), globalState)).isLessThanOrEqualTo((((_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))]) ? (new BigNumber((_380_v2).length)) : (_this.f33)));
          _379_v1 = (_379_v1).update(_this.f33, (!(_378_v0)) && (_378_v0));
          r0 = !(_378_v0);
          let _445_v41;
          let _nw65 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _445_v41 = _nw65;
          let _446_v42;
          _446_v42 = _dafny.Map.Empty.slice().updateUnsafe(false,_378_v0);
          let _index47 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_445_v41).length));
          (_445_v41)[_index47] = _446_v42;
          let _447_v43;
          _447_v43 = _dafny.Seq.UnicodeFromString("bxjwsbdy");
          let _index48 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_445_v41).length));
          let _rhs52 = _module.__default.safeDivisionInt((_this).f24, _this.f33);
          let _rhs53 = !_dafny.areEqual(_dafny.Seq.Concat(_447_v43, _dafny.Seq.UnicodeFromString("b")), _dafny.Seq.Concat(_447_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-744)), ((_448_v4) => function (_449_i3) {
            return _448_v4;
          })(_382_v4))));
          let _rhs54 = ((_446_v42).update((_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))], (_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))])).update((_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))], _378_v0);
          let _rhs55 = (function (_pat_let2_0) {
            return function (_450_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_451_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_451_dt__update_hcf1_h0, (_450_dt__update__tmp_h1).dtor_cf2, (_450_dt__update__tmp_h1).dtor_cf3);
                }(_pat_let3_0);
              }(_this.f33);
            }(_pat_let2_0);
          }(_module.__default.fm21(globalState))).dtor_cf2;
          let _lhs38 = globalState;
          let _lhs39 = _445_v41;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_445_v41).length));
          let _lhs41 = globalState;
          _lhs38.f18 = _rhs52;
          _378_v0 = _rhs53;
          _lhs39[_lhs40] = _rhs54;
          _lhs41.f1 = _rhs55;
          let _452_v44;
          let _453_v45;
          let _454_v46;
          let _455_v47;
          let _out18;
          let _out19;
          let _out20;
          let _out21;
          let _outcollector4 = (_this).m8(_441_v37, globalState);
          _out18 = _outcollector4[0];
          _out19 = _outcollector4[1];
          _out20 = _outcollector4[2];
          _out21 = _outcollector4[3];
          _452_v44 = _out18;
          _453_v45 = _out19;
          _454_v46 = _out20;
          _455_v47 = _out21;
        } else {
          let _456_v48;
          let _nw66 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _456_v48 = _nw66;
          let _index49 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_456_v48).length));
          (_456_v48)[_index49] = (_this).f24;
          let _index50 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_456_v48).length));
          let _rhs56 = _this.f33;
          let _rhs57 = _378_v0;
          let _lhs42 = _456_v48;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_456_v48).length));
          _lhs42[_lhs43] = _rhs56;
          r2 = _rhs57;
          r0 = _378_v0;
          let _457_v49;
          _457_v49 = _dafny.Map.Empty.slice().updateUnsafe(_384_v6,(_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))]);
          let _458_v50;
          let _nw67 = Array((new BigNumber(7)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _457_v49;
          _nw67[(_dafny.ONE).toNumber()] = (_457_v49).Merge(_457_v49);
          _nw67[(new BigNumber(2)).toNumber()] = _457_v49;
          _nw67[(new BigNumber(3)).toNumber()] = (_457_v49).update(_384_v6, (_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))]);
          _nw67[(new BigNumber(4)).toNumber()] = (_457_v49).Merge(_457_v49);
          _nw67[(new BigNumber(5)).toNumber()] = _457_v49;
          _nw67[(new BigNumber(6)).toNumber()] = _457_v49;
          _458_v50 = _nw67;
          _458_v50 = _458_v50;
          let _459_v51;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_378_v0);
          _459_v51 = _nw68;
          let _460_v52;
          _460_v52 = _dafny.Seq.of(_382_v4);
          let _461_v53;
          _461_v53 = _module.D1.create_DC3(_460_v52);
          let _462_v54;
          _462_v54 = _dafny.Set.fromElements(_461_v53, _461_v53, _module.D1.create_DC3(_460_v52), _461_v53, _module.D1.create_DC3(_460_v52));
          _459_v51 = ((!((_462_v54).IsSubsetOf(_462_v54))) ? (_459_v51) : (_459_v51));
          (globalState).f18 = ((((_441_v37)[_module.__default.safeIndex(new BigNumber(817), new BigNumber((_441_v37).length))]) ? (_this.f33) : ((_this).f24))).plus(new BigNumber(214));
        }
      } else {
        let _463_v55;
        _463_v55 = new _dafny.CodePoint('r'.codePointAt(0));
        let _464_v56;
        _464_v56 = _dafny.Seq.of(_463_v55);
        let _465_v57;
        _465_v57 = _dafny.Set.fromElements(_464_v56, _464_v56);
        _465_v57 = _465_v57;
        let _rhs58 = _378_v0;
        let _rhs59 = p0;
        let _rhs60 = (_dafny.ZERO).minus(_this.f33);
        let _rhs61 = (_module.__default.fm22(_this.f33, _378_v0, _378_v0, _module.__default.fm0(_this.f33, (_this).f24, globalState), globalState)).Intersect((_465_v57).Union(_dafny.Set.fromElements(_464_v56, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-411)), ((_466_v55) => function (_467_i4) {
          return _466_v55;
        })(_463_v55)))));
        let _rhs62 = _463_v55;
        let _lhs44 = globalState;
        let _lhs45 = globalState;
        r2 = _rhs58;
        _lhs44.f21 = _rhs59;
        _lhs45.f18 = _rhs60;
        _465_v57 = _rhs61;
        _463_v55 = _rhs62;
        if (_378_v0) {
          let _468_v58;
          _468_v58 = _dafny.MultiSet.fromElements(_this.f33, _this.f33);
          let _469_v59;
          let _nw69 = Array((new BigNumber(13)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = p0;
          _nw69[(_dafny.ONE).toNumber()] = p0;
          _nw69[(new BigNumber(2)).toNumber()] = p0;
          _nw69[(new BigNumber(3)).toNumber()] = p0;
          _nw69[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(p0, p0);
          _nw69[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_378_v0, false, _378_v0);
          _nw69[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.of(_378_v0));
          _nw69[(new BigNumber(7)).toNumber()] = p0;
          _nw69[(new BigNumber(8)).toNumber()] = p0;
          _nw69[(new BigNumber(9)).toNumber()] = _module.__default.fm23(globalState);
          _nw69[(new BigNumber(10)).toNumber()] = _module.__default.fm23(globalState);
          _nw69[(new BigNumber(11)).toNumber()] = p0;
          _nw69[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex((((_468_v58).contains(new BigNumber((p0).length))) ? ((_468_v58).get(new BigNumber((p0).length))) : (_this.f33)), new BigNumber((p0).length)), _378_v0);
          _469_v59 = _nw69;
          let _index51 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_469_v59).length));
          (_469_v59)[_index51] = p0;
          let _index52 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_469_v59).length));
          (_469_v59)[_index52] = p0;
          let _470_v60;
          _470_v60 = _dafny.Seq.of(_464_v56);
          (globalState).f2 = _module.__default.fm2(_module.__default.fm0(new BigNumber(((_470_v60)[_module.__default.safeIndex(_this.f33, new BigNumber((_470_v60).length))]).length), (_module.D3.create_DC11(_this.f33, ((_469_v59)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_469_v59).length))])[_module.__default.safeIndex((_this).f24, new BigNumber(((_469_v59)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_469_v59).length))]).length))], _378_v0)).dtor_cf18, globalState), _this.f33, _module.__default.fm3(globalState), globalState);
          let _471_v61;
          _471_v61 = _module.D4.create_DC13(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-751)), ((_472_v55) => function (_473_i5) {
  return _472_v55;
})(_463_v55))).length), _dafny.Map.Empty.slice().updateUnsafe(true,_this.f33), _464_v56);
          let _474_v62;
          _474_v62 = _dafny.Map.Empty.slice().updateUnsafe((_471_v61).dtor_cf22,new BigNumber(-176));
          _474_v62 = (_474_v62).update(_this.f33, ((_378_v0) ? (new BigNumber((_464_v56).length)) : ((_this).f24)));
          let _475_v63;
          let _nw70 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _475_v63 = _nw70;
          let _476_v64;
          _476_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,_dafny.Map.Empty.slice().updateUnsafe(_this.f33,_378_v0));
          let _index53 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_475_v63).length));
          (_475_v63)[_index53] = _module.__default.safeDivisionInt(_this.f33, new BigNumber((_476_v64).length));
          let _index54 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_475_v63).length));
          (_475_v63)[_index54] = _this.f33;
          let _477_v65;
          let _nw71 = new _module.C0();
          _nw71.__ctor(_378_v0);
          _477_v65 = _nw71;
        } else {
          let _478_v67;
          _478_v67 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_this.f33);
          let _479_v68;
          _479_v68 = _dafny.MultiSet.fromElements(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_380_v2).Elements) {
              let _480_v66 = _compr_17;
              if (_dafny.Seq.contains(_380_v2, _480_v66)) {
                _coll17.push([_module.__default.safeDivisionInt(_480_v66, _this.f33),_378_v0]);
              }
            }
            return _coll17;
          }(), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_378_v0), _379_v1, (_379_v1).update((((_478_v67).contains(true)) ? ((_478_v67).get(true)) : (_this.f33)), _378_v0), _379_v1);
          let _481_v69;
          _481_v69 = _module.D4.create_DC12(_380_v2);
          let _482_v70;
          _482_v70 = _module.D3.create_DC11(_this.f33, _378_v0, _378_v0);
          let _483_v71;
          _483_v71 = _dafny.MultiSet.fromElements(_378_v0, (_482_v70).dtor_cf19, _378_v0);
          let _484_v72;
          let _nw72 = Array((new BigNumber(17)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _this.f33;
          _nw72[(_dafny.ONE).toNumber()] = (_this).f24;
          _nw72[(new BigNumber(2)).toNumber()] = ((_dafny.ZERO).minus(_this.f33)).multipliedBy(_this.f33);
          _nw72[(new BigNumber(3)).toNumber()] = _this.f33;
          _nw72[(new BigNumber(4)).toNumber()] = ((_module.__default.fm2(_this.f33, (((_479_v68).contains(_379_v1)) ? ((_479_v68).get(_379_v1)) : (_this.f33)), _463_v55, globalState)) ? (new BigNumber(((_481_v69).dtor_cf21).length)) : (_this.f33));
          _nw72[(new BigNumber(5)).toNumber()] = ((_this).f24).multipliedBy(_this.f33);
          _nw72[(new BigNumber(6)).toNumber()] = new BigNumber((_483_v71).cardinality());
          _nw72[(new BigNumber(7)).toNumber()] = _this.f33;
          _nw72[(new BigNumber(8)).toNumber()] = new BigNumber(100);
          _nw72[(new BigNumber(9)).toNumber()] = _this.f33;
          _nw72[(new BigNumber(10)).toNumber()] = _this.f33;
          _nw72[(new BigNumber(11)).toNumber()] = (_this).f24;
          _nw72[(new BigNumber(12)).toNumber()] = new BigNumber((_464_v56).length);
          _nw72[(new BigNumber(13)).toNumber()] = _this.f33;
          _nw72[(new BigNumber(14)).toNumber()] = new BigNumber(856);
          _nw72[(new BigNumber(15)).toNumber()] = (new BigNumber((_464_v56).length)).plus((_this).f24);
          _nw72[(new BigNumber(16)).toNumber()] = new BigNumber(946);
          _484_v72 = _nw72;
          _484_v72 = _484_v72;
          let _index55 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_484_v72).length));
          (_484_v72)[_index55] = _this.f33;
          let _index56 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_484_v72).length));
          (_484_v72)[_index56] = (_this.f33).plus(new BigNumber(802));
          let _485_v73;
          let _nw73 = Array((new BigNumber(22)).toNumber()).fill(false);
          _485_v73 = _nw73;
          let _486_v74;
          _486_v74 = _dafny.Set.fromElements(_485_v73);
          let _487_v75;
          _487_v75 = _module.D1.create_DC5(_378_v0, (_486_v74).IsSubsetOf(_486_v74), _378_v0);
          _487_v75 = _487_v75;
          let _488_v76;
          let _nw74 = Array((new BigNumber(9)).toNumber()).fill(_module.D4.Default());
          _488_v76 = _nw74;
          let _index57 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_488_v76).length));
          (_488_v76)[_index57] = _481_v69;
          let _489_v77;
          _489_v77 = _module.D1.create_DC4(_378_v0, _378_v0, _463_v55, _this.f33, (_this).f24);
          let _pat_let_tv1 = _380_v2;
          let _index58 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_488_v76).length));
          let _rhs63 = function (_pat_let4_0) {
            return function (_490_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_491_dt__update_hcf21_h0) {
                  return _module.D4.create_DC12(_491_dt__update_hcf21_h0);
                }(_pat_let5_0);
              }(_pat_let_tv1);
            }(_pat_let4_0);
          }(_481_v69);
          let _rhs64 = (((_489_v77).dtor_cf9) ? ((_478_v67).Merge(_478_v67)) : (_478_v67));
          let _lhs46 = _488_v76;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_488_v76).length));
          _lhs46[_lhs47] = _rhs63;
          _478_v67 = _rhs64;
          let _492_v78;
          _492_v78 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_module.__default.fm24((_380_v2)[_module.__default.safeIndex((_this).f24, new BigNumber((_380_v2).length))], !(_378_v0), (_484_v72)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_484_v72).length))], (_dafny.ZERO).minus((_this).f24), globalState));
          _492_v78 = (_492_v78).update(_378_v0, _380_v2);
        }
        let _493_v79;
        let _nw75 = Array((new BigNumber(7)).toNumber());
        _nw75[(_dafny.ZERO).toNumber()] = _463_v55;
        _nw75[(_dafny.ONE).toNumber()] = _463_v55;
        _nw75[(new BigNumber(2)).toNumber()] = _463_v55;
        _nw75[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw75[(new BigNumber(4)).toNumber()] = (_464_v56)[_module.__default.safeIndex((_this).f24, new BigNumber((_464_v56).length))];
        _nw75[(new BigNumber(5)).toNumber()] = _463_v55;
        _nw75[(new BigNumber(6)).toNumber()] = _463_v55;
        _493_v79 = _nw75;
        let _494_v80;
        _494_v80 = _dafny.Seq.of(_493_v79);
        let _495_v81;
        _495_v81 = _module.D2.create_DC7((_494_v80)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_494_v80).length))]);
        let _source5 = _495_v81;
        if (_source5.is_DC8) {
          (globalState).f18 = (_this.f33).plus(((_378_v0) ? ((_this).f24) : ((_this).f24)));
          let _496_v82;
          let _nw76 = Array((new BigNumber(27)).toNumber()).fill([]);
          _496_v82 = _nw76;
          _496_v82 = _496_v82;
          let _497_v83;
          let _nw77 = new _module.C0();
          _nw77.__ctor(((_378_v0) ? (_378_v0) : (_378_v0)));
          _497_v83 = _nw77;
          _378_v0 = ((_this).f24).isLessThan((_this).f24);
        } else {
          let _498___mcc_h9 = (_source5).cf16;
          let _499_cf16 = _498___mcc_h9;
          let _500_v84;
          _500_v84 = _module.D1.create_DC5(_378_v0, _378_v0, true);
          let _pat_let_tv2 = _378_v0;
          let _501_v85;
          _501_v85 = _dafny.Map.Empty.slice().updateUnsafe((_380_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f33), new BigNumber((_380_v2).length))],function (_pat_let6_0) {
            return function (_502_dt__update__tmp_h3) {
              return function (_pat_let7_0) {
                return function (_503_dt__update_hcf13_h0) {
                  return _module.D1.create_DC5(_503_dt__update_hcf13_h0, (_502_dt__update__tmp_h3).dtor_cf14, (_502_dt__update__tmp_h3).dtor_cf15);
                }(_pat_let7_0);
              }(_pat_let_tv2);
            }(_pat_let6_0);
          }(_500_v84));
          let _504_v86;
          _504_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f33,_501_v85);
          _504_v86 = function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_module.__default.fm18(globalState)).Elements) {
              let _505_v87 = _compr_18;
              if ((_module.__default.fm18(globalState)).contains(_505_v87)) {
                _coll18.push([(_505_v87).multipliedBy(_this.f33),(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f33)),_500_v84)).Merge(_501_v85)]);
              }
            }
            return _coll18;
          }();
          let _506_v88;
          let _nw78 = new _module.C0();
          _nw78.__ctor(_378_v0);
          _506_v88 = _nw78;
          (globalState).f17 = _464_v56;
          let _507_v89;
          _507_v89 = _dafny.Map.Empty.slice().updateUnsafe(_464_v56,!(_378_v0));
          _507_v89 = (_507_v89).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(901)), ((_508_v55) => function (_509_i6) {
            return _508_v55;
          })(_463_v55)), ((_506_v88).f34) === ((_506_v88).f34));
        }
        (globalState).f14 = (_dafny.ZERO).minus((_this).f24);
      }
      let _510_v90;
      _510_v90 = _dafny.Seq.UnicodeFromString("frfeqaonm");
      let _511_v91;
      _511_v91 = _module.D4.create_DC13(_this.f33, _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_this.f33), _510_v90);
      let _source6 = _511_v91;
      if (_source6.is_DC13) {
        let _512___mcc_h10 = (_source6).cf22;
        let _513___mcc_h11 = (_source6).cf23;
        let _514___mcc_h12 = (_source6).cf24;
        let _515_cf24 = _514___mcc_h12;
        let _516_cf23 = _513___mcc_h11;
        let _517_cf22 = _512___mcc_h10;
        (globalState).f2 = _378_v0;
        r1 = (_dafny.ZERO).minus(_this.f33);
        if (_378_v0) {
          let _518_v92;
          _518_v92 = new _dafny.CodePoint('x'.codePointAt(0));
          let _rhs65 = _module.__default.fm2(_this.f33, _this.f33, _518_v92, globalState);
          let _rhs66 = (_this).fm11(_378_v0, _378_v0, !((_this).f24).isEqualTo(_this.f33), (_module.__default.fm25(globalState)).dtor_cf22, globalState);
          let _rhs67 = _378_v0;
          let _lhs48 = _this;
          let _lhs49 = globalState;
          _378_v0 = _rhs65;
          _lhs48.f33 = _rhs66;
          _lhs49.f2 = _rhs67;
          _379_v1 = (_379_v1).update(_517_cf22, true);
          (globalState).f2 = false;
          r2 = (_378_v0) === (!(_378_v0));
          let _519_v93;
          let _nw79 = Array((_dafny.ONE).toNumber()).fill(false);
          _519_v93 = _nw79;
          let _520_v94;
          _520_v94 = _dafny.Set.fromElements(_519_v93, _519_v93, _519_v93, _519_v93, _519_v93);
          _520_v94 = (_dafny.Set.fromElements(_519_v93, _519_v93, _519_v93)).Intersect(_520_v94);
        } else {
          let _521_v95;
          _521_v95 = new _dafny.CodePoint('l'.codePointAt(0));
          let _522_v96;
          _522_v96 = _dafny.MultiSet.fromElements((_379_v1).Merge(_379_v1));
          let _rhs68 = _521_v95;
          let _rhs69 = _522_v96;
          _521_v95 = _rhs68;
          _522_v96 = _rhs69;
          (globalState).f2 = ((_dafny.ZERO).minus((_this).f24)).isLessThan(_this.f33);
          let _523_v97;
          _523_v97 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(_515_cf24),false);
          let _524_v98;
          _524_v98 = _module.D1.create_DC3(_510_v90);
          _523_v97 = (_dafny.Map.Empty.slice().updateUnsafe(_524_v98,_378_v0)).Merge(_523_v97);
          _378_v0 = (_this).fm10(globalState);
          (_this).f33 = _this.f33;
        }
        let _525_v99;
        _525_v99 = new _dafny.CodePoint('n'.codePointAt(0));
        _525_v99 = _525_v99;
      } else if (_source6.is_DC12) {
        let _526___mcc_h13 = (_source6).cf21;
        let _527_cf21 = _526___mcc_h13;
        (globalState).f2 = _378_v0;
        (globalState).f2 = _dafny.areEqual(_dafny.Seq.Concat(_module.__default.fm23(globalState), p0), p0);
        (globalState).f18 = ((_this).f24).minus(_this.f33);
        let _528_v100;
        let _nw80 = Array((new BigNumber(7)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = false;
        _nw80[(_dafny.ONE).toNumber()] = _378_v0;
        _nw80[(new BigNumber(2)).toNumber()] = (_378_v0) === (_378_v0);
        _nw80[(new BigNumber(3)).toNumber()] = _378_v0;
        _nw80[(new BigNumber(4)).toNumber()] = _378_v0;
        _nw80[(new BigNumber(5)).toNumber()] = !(_dafny.areEqual(_510_v90, _dafny.Seq.UnicodeFromString("lruyipny")));
        _nw80[(new BigNumber(6)).toNumber()] = true;
        _528_v100 = _nw80;
        let _index59 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_528_v100).length));
        (_528_v100)[_index59] = _378_v0;
        let _index60 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_528_v100).length));
        (_528_v100)[_index60] = _378_v0;
      } else {
        let _529___mcc_h14 = (_source6).cf25;
        let _530_cf25 = _529___mcc_h14;
        let _531_v101;
        let _nw81 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _531_v101 = _nw81;
        let _532_v102;
        _532_v102 = _dafny.Map.Empty.slice().updateUnsafe(_378_v0,_378_v0);
        let _index61 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_531_v101).length));
        (_531_v101)[_index61] = _532_v102;
        let _533_v103;
        _533_v103 = _dafny.MultiSet.fromElements((_this).f24, _this.f33);
        let _index62 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_531_v101).length));
        let _rhs70 = (_dafny.ZERO).minus((_this).f24);
        let _rhs71 = _532_v102;
        let _rhs72 = ((_533_v103).Intersect(_533_v103)).Difference(_533_v103);
        let _lhs50 = globalState;
        let _lhs51 = _531_v101;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_531_v101).length));
        _lhs50.f9 = _rhs70;
        _lhs51[_lhs52] = _rhs71;
        _533_v103 = _rhs72;
        let _534_v104;
        _534_v104 = _dafny.MultiSet.fromElements(_module.__default.fm26(globalState));
        let _535_v105;
        let _init12 = ((_536_v0) => function (_537_i7) {
          return _536_v0;
        })(_378_v0);
        let _nw82 = Array((new BigNumber(22)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw82.length); _i0_12++) {
          _nw82[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _535_v105 = _nw82;
        let _538_v106;
        _538_v106 = _dafny.Map.Empty.slice().updateUnsafe(_535_v105,(_this).f24);
        let _539_v107;
        _539_v107 = _module.D4.create_DC13((((_538_v106).contains(_535_v105)) ? ((_538_v106).get(_535_v105)) : (new BigNumber((_dafny.MultiSet.fromElements(_this.f33)).cardinality()))), _module.__default.fm17((_this).f24, new BigNumber((_380_v2).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_378_v0,_378_v0)).length), true, globalState), _510_v90);
        let _540_v108;
        _540_v108 = _module.D4.create_DC14(_539_v107);
        let _541_v109;
        _541_v109 = _module.D4.create_DC14(_539_v107);
        let _542_v110;
        _542_v110 = _module.D4.create_DC14(_540_v108);
        let _543_v111;
        _543_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(97),new BigNumber(-620));
        let _544_v112;
        _544_v112 = new _dafny.CodePoint('y'.codePointAt(0));
        let _545_v113;
        let _nw83 = Array((new BigNumber(12)).toNumber());
        _nw83[(_dafny.ZERO).toNumber()] = ((_this).f24).plus((_380_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_534_v104).update(_542_v110, _module.__default.abs(_this.f33))).cardinality())), new BigNumber((_380_v2).length))]);
        _nw83[(_dafny.ONE).toNumber()] = _this.f33;
        _nw83[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw83[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this.f33).plus(new BigNumber((_dafny.Seq.UnicodeFromString("bgboi")).length)));
        _nw83[(new BigNumber(4)).toNumber()] = (new BigNumber((_533_v103).cardinality())).minus(new BigNumber((_543_v111).length));
        _nw83[(new BigNumber(5)).toNumber()] = (_this).f24;
        _nw83[(new BigNumber(6)).toNumber()] = new BigNumber((_379_v1).length);
        _nw83[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw83[(new BigNumber(8)).toNumber()] = _this.f33;
        _nw83[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw83[(new BigNumber(10)).toNumber()] = (_this).f24;
        _nw83[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update(_510_v90, _module.__default.safeIndex(_this.f33, new BigNumber((_510_v90).length)), _544_v112)).length);
        _545_v113 = _nw83;
        _545_v113 = _545_v113;
        let _546_v114;
        _546_v114 = _module.D1.create_DC6();
        _546_v114 = _module.__default.fm27(globalState);
        let _547_v115;
        _547_v115 = _dafny.Map.Empty.slice().updateUnsafe(_544_v112,(_this).f24);
        let _548_v116;
        _548_v116 = _dafny.Map.Empty.slice().updateUnsafe(_547_v115,(_this).f24);
        let _549_v117;
        _549_v117 = _dafny.Set.fromElements(_378_v0, _378_v0);
        let _550_v118;
        _550_v118 = _dafny.Seq.of((_548_v116).update(_547_v115, (_this).f24), _548_v116, _dafny.Map.Empty.slice().updateUnsafe(_547_v115,new BigNumber((_549_v117).length)));
        _548_v116 = ((_550_v118)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_550_v118).length))]).Merge(_548_v116);
      }
      (globalState).f14 = (_this).f24;
      (globalState).f14 = _this.f33;
      r0 = _378_v0;
      r1 = (_this).f24;
      let _551_v119;
      _551_v119 = _dafny.Set.fromElements(_this.f33);
      r2 = (function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_380_v2).Elements) {
          let _552_v120 = _compr_19;
          if (_dafny.Seq.contains(_380_v2, _552_v120)) {
            _coll19.add(_module.__default.safeDivisionInt(_552_v120, new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(true, true))).length)));
          }
        }
        return _coll19;
      }()).IsSubsetOf(_551_v119);
      return [r0, r1, r2];
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _553_v0;
      _553_v0 = false;
      r1 = ((_553_v0) ? (_553_v0) : (((_553_v0) ? (_553_v0) : (_553_v0))));
      let _554_v1;
      _554_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(319),(_dafny.ZERO).minus((_this).f24));
      let _555_v2;
      _555_v2 = _dafny.Seq.UnicodeFromString("nedubbppa");
      let _556_v3;
      let _nw84 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _556_v3 = _nw84;
      let _557_v4;
      _557_v4 = _dafny.Set.fromElements(_556_v3, _556_v3, _556_v3);
      let _558_v5;
      _558_v5 = _dafny.Set.fromElements(new BigNumber((_555_v2).length));
      let _559_v6;
      _559_v6 = _module.D6.create_DC16(_558_v5);
      let _560_v7;
      _560_v7 = _dafny.Seq.of(new BigNumber(((_559_v6).dtor_cf27).length));
      let _561_v8;
      _561_v8 = _dafny.Seq.of(_553_v0, _553_v0);
      let _562_v9;
      let _nw85 = Array((new BigNumber(26)).toNumber());
      _nw85[(_dafny.ZERO).toNumber()] = (_this).f24;
      _nw85[(_dafny.ONE).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(2)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(3)).toNumber()] = (((_554_v1).contains(_this.f33)) ? ((_554_v1).get(_this.f33)) : (_this.f33));
      _nw85[(new BigNumber(4)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(5)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(6)).toNumber()] = new BigNumber((_555_v2).length);
      _nw85[(new BigNumber(7)).toNumber()] = new BigNumber((_555_v2).length);
      _nw85[(new BigNumber(8)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(9)).toNumber()] = new BigNumber((_557_v4).length);
      _nw85[(new BigNumber(10)).toNumber()] = _module.__default.fm0((_this).f24, new BigNumber(580), globalState);
      _nw85[(new BigNumber(11)).toNumber()] = _this.f33;
      _nw85[(new BigNumber(12)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(13)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(14)).toNumber()] = _this.f33;
      _nw85[(new BigNumber(15)).toNumber()] = new BigNumber((_560_v7).length);
      _nw85[(new BigNumber(16)).toNumber()] = _this.f33;
      _nw85[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_561_v8).length));
      _nw85[(new BigNumber(18)).toNumber()] = new BigNumber(-528);
      _nw85[(new BigNumber(19)).toNumber()] = _this.f33;
      _nw85[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_560_v7).length));
      _nw85[(new BigNumber(21)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(22)).toNumber()] = (_this).f24;
      _nw85[(new BigNumber(23)).toNumber()] = _this.f33;
      _nw85[(new BigNumber(24)).toNumber()] = new BigNumber(108);
      _nw85[(new BigNumber(25)).toNumber()] = _this.f33;
      _562_v9 = _nw85;
      let _563_v10;
      _563_v10 = _dafny.MultiSet.fromElements(_562_v9);
      if (!((_563_v10).equals(_563_v10))) {
        let _564_v11;
        _564_v11 = _dafny.Map.Empty.slice().updateUnsafe(!(_553_v0),new BigNumber((_dafny.Seq.Concat(_561_v8, _561_v8)).length));
        _564_v11 = (_564_v11).update(_553_v0, (_this).f24);
        let _565_v12;
        let _nw86 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
        _565_v12 = _nw86;
        let _566_v13;
        _566_v13 = _module.D4.create_DC14(_module.D4.create_DC13(_this.f33, _564_v11, _555_v2));
        let _567_v14;
        _567_v14 = _dafny.Map.Empty.slice().updateUnsafe(_566_v13,_555_v2);
        let _index63 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_565_v12).length));
        (_565_v12)[_index63] = _567_v14;
        let _568_v15;
        _568_v15 = _module.D7.create_DC20(_567_v14);
        let _index64 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_565_v12).length));
        (_565_v12)[_index64] = (_568_v15).dtor_cf31;
        (globalState).f18 = _this.f33;
        (globalState).f2 = _553_v0;
        let _569_v16;
        _569_v16 = new _dafny.CodePoint('u'.codePointAt(0));
        let _570_v17;
        _570_v17 = _dafny.Map.Empty.slice().updateUnsafe(_553_v0,_553_v0);
        let _571_v18;
        _571_v18 = _dafny.MultiSet.fromElements(new BigNumber((_570_v17).length), (_this).f24);
        let _572_v19;
        _572_v19 = _dafny.MultiSet.fromElements((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_569_v16,!(_553_v0))).length)).isEqualTo(new BigNumber((_571_v18).cardinality())));
        _572_v19 = _module.__default.fm14(_554_v1, _this.f33, globalState);
      } else {
        (globalState).f2 = !_dafny.areEqual(_561_v8, _561_v8);
        if (!((_this).f24).isEqualTo(_module.__default.fm0((_this).f24, (_this).f24, globalState))) {
          let _573_v20;
          _573_v20 = _dafny.Set.fromElements(_553_v0, !(_553_v0), _553_v0, _553_v0, _553_v0);
          let _574_v21;
          _574_v21 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_555_v2).length), new BigNumber((_573_v20).length), _this.f33));
          let _575_v22;
          let _nw87 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _575_v22 = _nw87;
          let _index65 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_575_v22).length));
          (_575_v22)[_index65] = _module.__default.fm24(_this.f33, _553_v0, new BigNumber((_dafny.Seq.UnicodeFromString("dn")).length), _this.f33, globalState);
          let _index66 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((p0).length));
          (p0)[_index66] = (new BigNumber((_558_v5).length)).isEqualTo((_this).f24);
          let _index67 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_575_v22).length));
          let _index68 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((p0).length));
          let _rhs73 = _module.__default.fm24((_dafny.ZERO).minus(_this.f33), _553_v0, new BigNumber((_dafny.Seq.Concat(_555_v2, _555_v2)).length), new BigNumber(151), globalState);
          let _rhs74 = _574_v21;
          let _rhs75 = _this.f33;
          let _rhs76 = _dafny.Seq.of((_this).f24, new BigNumber(829), _module.__default.safeModuloInt(_this.f33, new BigNumber((_555_v2).length)));
          let _rhs77 = _553_v0;
          let _lhs53 = _this;
          let _lhs54 = _575_v22;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_575_v22).length));
          let _lhs56 = p0;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((p0).length));
          _560_v7 = _rhs73;
          _574_v21 = _rhs74;
          _lhs53.f33 = _rhs75;
          _lhs54[_lhs55] = _rhs76;
          _lhs56[_lhs57] = _rhs77;
          let _576_v23;
          let _nw88 = new _module.C0();
          _nw88.__ctor((_this.f33).isLessThanOrEqualTo(_this.f33));
          _576_v23 = _nw88;
          r1 = (p0)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((p0).length))];
          let _577_v24;
          _577_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_575_v22)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_575_v22).length))], _module.__default.safeIndex(_this.f33, new BigNumber(((_575_v22)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_575_v22).length))]).length)), _this.f33),_module.D6.create_DC17(_553_v0));
          let _578_v25;
          _578_v25 = _module.D6.create_DC17(!((_576_v23).f34));
          _577_v24 = (_577_v24).update(_560_v7, (((p0)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((p0).length))]) ? (_578_v25) : (_578_v25)));
          let _579_v26;
          let _nw89 = new _module.C0();
          _nw89.__ctor((_576_v23).f34);
          _579_v26 = _nw89;
        } else {
          let _580_v27;
          _580_v27 = _dafny.Set.fromElements(_553_v0);
          let _rhs78 = _553_v0;
          let _rhs79 = (_dafny.Set.fromElements(_553_v0)).Difference(_580_v27);
          let _rhs80 = false;
          let _lhs58 = globalState;
          _553_v0 = _rhs78;
          _580_v27 = _rhs79;
          _lhs58.f2 = _rhs80;
          r1 = _553_v0;
          let _581_v28;
          _581_v28 = _module.D1.create_DC5(_553_v0, _553_v0, _553_v0);
          let _582_v29;
          _582_v29 = _module.D6.create_DC17(_553_v0);
          let _pat_let_tv3 = _553_v0;
          let _583_v30;
          _583_v30 = _module.D7.create_DC21(function (_pat_let8_0) {
  return function (_584_dt__update__tmp_h0) {
    return function (_pat_let9_0) {
      return function (_585_dt__update_hcf15_h0) {
        return _module.D1.create_DC5((_584_dt__update__tmp_h0).dtor_cf13, (_584_dt__update__tmp_h0).dtor_cf14, _585_dt__update_hcf15_h0);
      }(_pat_let9_0);
    }(_pat_let_tv3);
  }(_pat_let8_0);
}(_581_v28), _module.D6.create_DC19(_582_v29), p0);
          let _586_v31;
          _586_v31 = _dafny.Map.Empty.slice().updateUnsafe(_553_v0,_583_v30);
          let _587_v32;
          _587_v32 = _dafny.MultiSet.fromElements(_553_v0);
          let _588_v33;
          _588_v33 = _dafny.Map.Empty.slice().updateUnsafe(_553_v0,true);
          let _rhs81 = ((new BigNumber((_561_v8).length)).minus(_this.f33)).minus((((_587_v32).contains(_553_v0)) ? ((_587_v32).get(_553_v0)) : (_this.f33)));
          let _rhs82 = new BigNumber((_561_v8).length);
          let _rhs83 = (_dafny.Map.Empty.slice().updateUnsafe(_553_v0,_583_v30)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_553_v0),_583_v30));
          let _rhs84 = ((_dafny.ZERO).minus((new BigNumber((_588_v33).length)).minus((_this).f24))).isLessThanOrEqualTo((_this).f24);
          let _lhs59 = globalState;
          let _lhs60 = globalState;
          let _lhs61 = globalState;
          _lhs59.f9 = _rhs81;
          _lhs60.f18 = _rhs82;
          _586_v31 = _rhs83;
          _lhs61.f2 = _rhs84;
          let _589_v34;
          _589_v34 = _dafny.Map.Empty.slice().updateUnsafe(_561_v8,(_this).f24);
          _589_v34 = (_589_v34).update(_dafny.Seq.Concat(_561_v8, _561_v8), new BigNumber(997));
          (globalState).f1 = (_dafny.ZERO).minus(new BigNumber(-374));
        }
        _562_v9 = _562_v9;
        let _index69 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((p0).length));
        (p0)[_index69] = (_553_v0) || (_553_v0);
        let _index70 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((p0).length));
        (p0)[_index70] = !(!(_553_v0));
        let _590_v35;
        _590_v35 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(841)), ((_591_v2) => function (_592_i0) {
          return (_591_v2)[_module.__default.safeIndex(_this.f33, new BigNumber((_591_v2).length))];
        })(_555_v2)));
        let _593_v36;
        let _nw90 = new _module.C0();
        _nw90.__ctor(!((_590_v35).IsDisjointFrom(_590_v35)));
        _593_v36 = _nw90;
      }
      (globalState).f2 = (_557_v4).IsDisjointFrom(_557_v4);
      let _594_v37;
      _594_v37 = _module.D1.create_DC5(_553_v0, true, _553_v0);
      let _595_v38;
      _595_v38 = _module.D6.create_DC16(_dafny.Set.fromElements(_this.f33));
      let _596_v39;
      _596_v39 = _module.D7.create_DC21(_594_v37, _module.D6.create_DC19(_595_v38), p0);
      let _597_v40;
      _597_v40 = _module.D7.create_DC22(_596_v39);
      let _598_v41;
      _598_v41 = _dafny.Map.Empty.slice().updateUnsafe(_597_v40,p0);
      _598_v41 = (_598_v41).update(_module.D7.create_DC22(_596_v39), p0);
      let _599_v42;
      _599_v42 = _dafny.MultiSet.fromElements((_this).f24, (_dafny.ZERO).minus(new BigNumber((_558_v5).length)));
      let _600_v43;
      _600_v43 = _dafny.Set.fromElements((new BigNumber((_599_v42).cardinality())).isLessThan(new BigNumber(-939)), _553_v0, ((!(_553_v0)) ? (_553_v0) : (_553_v0)), _553_v0, _553_v0);
      let _601_v44;
      _601_v44 = _dafny.MultiSet.fromElements((_module.__default.fm28(_553_v0, new BigNumber((_555_v2).length), _553_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_553_v0,_this.f33)).length), globalState)).IsProperSubsetOf(_600_v43), _553_v0, _553_v0);
      let _602_v45;
      _602_v45 = _module.D8.create_DC23(_dafny.Set.fromElements(_553_v0, _553_v0, _553_v0, !(_553_v0)));
      let _rhs85 = (_602_v45).dtor_cf36;
      let _rhs86 = ((new BigNumber(830)).plus(_this.f33)).plus(new BigNumber(689));
      let _rhs87 = (_601_v44).update(_553_v0, _module.__default.abs((_this).f24));
      let _lhs62 = globalState;
      _600_v43 = _rhs85;
      _lhs62.f18 = _rhs86;
      _601_v44 = _rhs87;
      let _hi3 = ((_this).f24).multipliedBy(new BigNumber((_555_v2).length));
      for (let _603_i1 = ((_553_v0) ? ((_this).f24) : (_this.f33)); _603_i1.isLessThan(_hi3); _603_i1 = _603_i1.plus(_dafny.ONE)) {
        _555_v2 = _555_v2;
        r1 = ((_this).f24).isLessThan((_dafny.ZERO).minus((_module.__default.fm0((_this).f24, new BigNumber(179), globalState)).minus(_this.f33)));
        if (_553_v0) {
          let _604_v46;
          _604_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
          let _rhs88 = _module.__default.safeModuloInt(new BigNumber((_604_v46).length), (_603_i1).plus(_this.f33));
          let _rhs89 = !((_559_v6).dtor_cf27).equals(function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(-841))) {
              let _605_v47 = _compr_20;
              if (((new BigNumber(280)).isLessThanOrEqualTo(_605_v47)) && ((_605_v47).isLessThan(new BigNumber(-841)))) {
                _coll20.add(_module.__default.safeModuloInt(_605_v47, new BigNumber((_601_v44).cardinality())));
              }
            }
            return _coll20;
          }());
          let _lhs63 = globalState;
          _lhs63.f18 = _rhs88;
          r1 = _rhs89;
          let _606_v48;
          let _init13 = ((_607_v2, _608_i1) => function (_609_i2) {
            return _dafny.Seq.update(_607_v2, _module.__default.safeIndex(_608_i1, new BigNumber((_607_v2).length)), new _dafny.CodePoint('v'.codePointAt(0)));
          })(_555_v2, _603_i1);
          let _nw91 = Array((new BigNumber(16)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw91.length); _i0_13++) {
            _nw91[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _606_v48 = _nw91;
          _606_v48 = _606_v48;
          let _610_v49;
          let _out22;
          _out22 = _module.__default.m0(_this.f33, ((_this).f24).minus(new BigNumber(-447)), globalState);
          _610_v49 = _out22;
          let _611_v50;
          _611_v50 = _module.D6.create_DC17(_553_v0);
          let _612_v51;
          _612_v51 = new _dafny.CodePoint('d'.codePointAt(0));
          let _613_v52;
          _613_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_555_v2).length),_module.__default.fm4(_553_v0, _610_v49, _612_v51, _553_v0, globalState));
          let _614_v53;
          _614_v53 = _dafny.Map.Empty.slice().updateUnsafe(_611_v50,_613_v52);
          let _615_v54;
          _615_v54 = _dafny.Map.Empty.slice().updateUnsafe(_610_v49,_613_v52);
          _614_v53 = (_614_v53).update(_611_v50, ((((_615_v54).contains(_610_v49)) ? ((_615_v54).get(_610_v49)) : (_dafny.Map.Empty.slice().updateUnsafe(_610_v49,_555_v2)))).Merge(_613_v52));
          let _rhs90 = _553_v0;
          let _rhs91 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(212)), ((_616_v51) => function (_617_i3) {
            return _616_v51;
          })(_612_v51)), _555_v2);
          let _lhs64 = globalState;
          let _lhs65 = globalState;
          _lhs64.f2 = _rhs90;
          _lhs65.f17 = _rhs91;
        } else {
          (globalState).f21 = _561_v8;
          let _618_v55;
          _618_v55 = _dafny.Map.Empty.slice().updateUnsafe(_553_v0,new BigNumber((_dafny.Seq.update(_560_v7, _module.__default.safeIndex(_603_i1, new BigNumber((_560_v7).length)), _603_i1)).length));
          let _619_v56;
          _619_v56 = _dafny.Map.Empty.slice().updateUnsafe(_553_v0,new BigNumber((_618_v55).length));
          r0 = ((((_619_v56).contains(_553_v0)) ? ((_619_v56).get(_553_v0)) : (new BigNumber(679)))).minus(_module.__default.safeDivisionInt(_603_i1, _this.f33));
          let _620_v57;
          let _nw92 = new _module.C0();
          _nw92.__ctor(_553_v0);
          _620_v57 = _nw92;
          let _621_v58;
          _621_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_622_i4) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length)),(_620_v57).f34);
          _621_v58 = _621_v58;
          let _index71 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((p0).length));
          (p0)[_index71] = false;
          let _623_v59;
          _623_v59 = new _dafny.CodePoint('s'.codePointAt(0));
          let _624_v60;
          _624_v60 = _dafny.Map.Empty.slice().updateUnsafe(_562_v9,_623_v59);
          let _index72 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((p0).length));
          let _rhs92 = !((((true) ? (new BigNumber(701)) : (_603_i1))).isLessThan((_dafny.ZERO).minus(_603_i1)));
          let _rhs93 = (_560_v7)[_module.__default.safeIndex((new BigNumber(102)).plus(new BigNumber(97)), new BigNumber((_560_v7).length))];
          let _rhs94 = (_dafny.Map.Empty.slice().updateUnsafe(_556_v3,_623_v59)).Merge((_624_v60).Merge(_dafny.Map.Empty.slice().updateUnsafe(_556_v3,_623_v59)));
          let _lhs66 = p0;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((p0).length));
          let _lhs68 = globalState;
          _lhs66[_lhs67] = _rhs92;
          _lhs68.f9 = _rhs93;
          _624_v60 = _rhs94;
        }
        r1 = (new BigNumber(134)).isEqualTo(_603_i1);
      }
      r0 = (_this.f33).multipliedBy((_module.__default.fm0((_this).f24, _this.f33, globalState)).multipliedBy(new BigNumber(610)));
      r1 = _553_v0;
      let _625_v61;
      _625_v61 = new _dafny.CodePoint('f'.codePointAt(0));
      let _626_v62;
      _626_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_601_v44).cardinality()),false);
      r2 = (_560_v7)[_module.__default.safeIndex(_module.__default.fm0((_dafny.ZERO).minus((_module.D1.create_DC4(_553_v0, _553_v0, _625_v61, new BigNumber((_560_v7).length), _module.__default.fm0(_this.f33, _module.__default.fm0(new BigNumber((_626_v62).length), (_this).f24, globalState), globalState))).dtor_cf12), _this.f33, globalState), new BigNumber((_560_v7).length))];
      r3 = (_555_v2)[_module.__default.safeIndex((new BigNumber(-930)).minus(new BigNumber(447)), new BigNumber((_555_v2).length))];
      return [r0, r1, r2, r3];
    }
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
      this.f31 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f31, f24, f25) {
      let _this = this;
      (_this).f31 = f31;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(_module.D1.create_DC6())).IsSubsetOf(_dafny.MultiSet.fromElements(_module.D1.create_DC6(), _module.D1.create_DC6()));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let _627_v0;
      _627_v0 = _module.D0.create_DC1((_this).f24, (_this).f24, _this.f31);
      if (function (_source7) {
        if (_source7.is_DC1) {
          let _628___mcc_h4 = (_source7).cf1;
          let _629___mcc_h5 = (_source7).cf2;
          let _630___mcc_h6 = (_source7).cf3;
          let _631_cf3 = _630___mcc_h6;
          let _632_cf2 = _629___mcc_h5;
          let _633_cf1 = _628___mcc_h4;
          return _631_cf3;
        } else if (_source7.is_DC2) {
          let _634___mcc_h7 = (_source7).cf4;
          let _635___mcc_h8 = (_source7).cf5;
          let _636___mcc_h9 = (_source7).cf6;
          let _637_cf6 = _636___mcc_h9;
          let _638_cf5 = _635___mcc_h8;
          let _639_cf4 = _634___mcc_h7;
          return !(_this.f31);
        } else {
          let _640___mcc_h10 = (_source7).cf0;
          let _641_cf0 = _640___mcc_h10;
          return _this.f31;
        }
      }(_627_v0)) {
        let _642_v1;
        let _init14 = function (_643_i0) {
          return !(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-591)), function (_644_i1) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length)).isEqualTo((_this).f24);
        };
        let _nw93 = Array((new BigNumber(14)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw93.length); _i0_14++) {
          _nw93[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _642_v1 = _nw93;
        let _rhs95 = _642_v1;
        let _rhs96 = (_this.f31) === ((_this.f31) && (_this.f31));
        let _lhs69 = globalState;
        _642_v1 = _rhs95;
        _lhs69.f2 = _rhs96;
        if (_this.f31) {
          let _645_v2;
          _645_v2 = _dafny.Seq.UnicodeFromString("dvbaaxfn");
          let _index73 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length));
          (_642_v1)[_index73] = ((_this).f24).isLessThanOrEqualTo(new BigNumber((_645_v2).length));
          let _index74 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length));
          (_642_v1)[_index74] = _this.f31;
          let _646_v3;
          _646_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(69)), function (_647_i2) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length));
          (globalState).f9 = ((_this).f24).minus(_module.__default.safeDivisionInt(new BigNumber((_646_v3).length), (_this).f24));
          (globalState).f18 = (_this).f24;
          let _648_v4;
          _648_v4 = _dafny.MultiSet.fromElements((_642_v1)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length))], (_642_v1)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length))], _this.f31);
          let _649_v5;
          _649_v5 = new _dafny.CodePoint('o'.codePointAt(0));
          let _650_v6;
          let _nw94 = Array((_dafny.ONE).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = _649_v5;
          _650_v6 = _nw94;
          let _index75 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_650_v6).length));
          (_650_v6)[_index75] = ((!(_this.f31)) ? (_module.__default.fm3(globalState)) : (new _dafny.CodePoint('t'.codePointAt(0))));
          let _651_v7;
          _651_v7 = _module.D0.create_DC2(_this.f31, (_this).f24, _649_v5);
          let _index76 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_650_v6).length));
          let _rhs97 = !((_this).fm9((_642_v1)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length))], _module.__default.fm0((_this).f24, (_this).f24, globalState), (_this).f24, globalState));
          let _rhs98 = _648_v4;
          let _rhs99 = (_651_v7).dtor_cf6;
          let _lhs70 = _642_v1;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length));
          let _lhs72 = _650_v6;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_650_v6).length));
          _lhs70[_lhs71] = _rhs97;
          _648_v4 = _rhs98;
          _lhs72[_lhs73] = _rhs99;
          let _index78 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length));
          (_642_v1)[_index78] = ((_642_v1)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_642_v1).length))]) || (_this.f31);
        } else {
          r0 = _this.f31;
          let _index79 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_642_v1).length));
          (_642_v1)[_index79] = (new BigNumber(155)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-796)), function (_652_i3) {
            return (_this).f24;
          })).length));
          let _index80 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_642_v1).length));
          (_642_v1)[_index80] = _this.f31;
          (globalState).f2 = (_this).fm9(_this.f31, (_this).f24, (_this).f24, globalState);
          let _653_v8;
          _653_v8 = _dafny.Seq.UnicodeFromString("xuqmdjtcc");
          let _654_v9;
          _654_v9 = new _dafny.CodePoint('e'.codePointAt(0));
          let _655_v10;
          _655_v10 = _module.D0.create_DC2((_642_v1)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_642_v1).length))], new BigNumber((_dafny.Seq.UnicodeFromString("qoctswai")).length), _654_v9);
          let _656_v11;
          _656_v11 = _dafny.Map.Empty.slice().updateUnsafe((_655_v10).dtor_cf6,(_this).f24);
          let _657_v12;
          let _nw95 = new _module.C1();
          _nw95.__ctor(_module.__default.fm29(new BigNumber(743), (_this).f24, globalState), (_this).f24, (_this).f24, _dafny.Seq.update((_this).f25, _module.__default.safeIndex(new BigNumber((_653_v8).length), new BigNumber(((_this).f25).length)), _656_v11));
          _657_v12 = _nw95;
          let _nw96 = new _module.C1();
          _nw96.__ctor((_657_v12).f32, (_this).f24, (_this).f24, (_this).f25);
          _657_v12 = _nw96;
        }
        let _658_v13;
        _658_v13 = _dafny.MultiSet.fromElements(_642_v1, _642_v1, _642_v1);
        let _659_v14;
        let _init15 = function (_660_i4) {
          return _module.__default.safeModuloInt(_660_i4, (_this).f24);
        };
        let _nw97 = Array((new BigNumber(19)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw97.length); _i0_15++) {
          _nw97[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _659_v14 = _nw97;
        let _661_v15;
        _661_v15 = _dafny.Map.Empty.slice().updateUnsafe((_658_v13).Difference(_658_v13),_659_v14);
        _661_v15 = (_661_v15).update(_658_v13, _659_v14);
        let _662_v16;
        _662_v16 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_663_i5) {
  return new _dafny.CodePoint('j'.codePointAt(0));
}));
        let _664_v17;
        _664_v17 = _module.D8.create_DC26(_662_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(623)), function (_665_i6) {
  return new _dafny.CodePoint('a'.codePointAt(0));
}));
        let _666_v18;
        _666_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_664_v17)).length));
        let _667_v19;
        _667_v19 = _dafny.Seq.of(_666_v18, _module.__default.fm17((_this).f24, (_this).f24, (_this).f24, _this.f31, globalState), _666_v18, _666_v18, _666_v18);
        _667_v19 = _667_v19;
        let _668_v20;
        _668_v20 = _module.D1.create_DC5(true, _this.f31, _this.f31);
        let _669_v21;
        _669_v21 = _module.D6.create_DC19(_module.D6.create_DC18(false));
        let _670_v22;
        _670_v22 = _module.D6.create_DC19(_669_v21);
        let _671_v23;
        _671_v23 = _module.D6.create_DC19((_module.D7.create_DC21(_668_v20, _670_v22, _642_v1)).dtor_cf33);
        let _source8 = _670_v22;
        if (_source8.is_DC17) {
          let _672___mcc_h0 = (_source8).cf28;
          let _673_cf28 = _672___mcc_h0;
          let _674_v24;
          _674_v24 = _dafny.Seq.UnicodeFromString("ibwh");
          (globalState).f17 = _674_v24;
          let _675_v25;
          _675_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_673_cf28);
          let _676_v26;
          _676_v26 = _module.D6.create_DC18((((_675_v25).contains(_673_cf28)) ? ((_675_v25).get(_673_cf28)) : (_673_cf28)));
          _673_cf28 = (_676_v26).dtor_cf29;
          let _677_v27;
          _677_v27 = _dafny.Seq.of((_this).f24);
          let _index81 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_659_v14).length));
          (_659_v14)[_index81] = _module.__default.safeDivisionInt(new BigNumber((_677_v27).length), (_this).f24);
          let _index82 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_659_v14).length));
          (_659_v14)[_index82] = _module.__default.safeDivisionInt((_this).f24, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_674_v24).length))));
          let _678_v28;
          _678_v28 = _dafny.Seq.of(_673_cf28, _this.f31, _673_cf28);
          let _679_v29;
          _679_v29 = _dafny.MultiSet.fromElements((_659_v14)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_659_v14).length))]);
          let _680_v30;
          _680_v30 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f24), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_678_v28).length))), _dafny.MultiSet.fromElements(new BigNumber((_675_v25).length)), (_679_v29).update(new BigNumber((_677_v27).length), _module.__default.abs((_659_v14)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_659_v14).length))])));
          let _681_v31;
          let _nw98 = new _module.C1();
          _nw98.__ctor(_680_v30, new BigNumber(-808), (_this).f24, (_this).f25);
          _681_v31 = _nw98;
        } else if (_source8.is_DC18) {
          let _682___mcc_h1 = (_source8).cf29;
          let _683_cf29 = _682___mcc_h1;
          let _684_v32;
          _684_v32 = new _dafny.CodePoint('d'.codePointAt(0));
          let _685_v33;
          _685_v33 = _dafny.Seq.of(_this.f31, _683_cf29);
          let _686_v34;
          _686_v34 = _dafny.Map.Empty.slice().updateUnsafe(_684_v32,_dafny.Seq.Concat(_685_v33, _685_v33));
          let _687_v35;
          _687_v35 = _dafny.MultiSet.fromElements(_683_cf29);
          let _688_v36;
          _688_v36 = _687_v35;
          let _rhs100 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_685_v33);
          let _rhs101 = _687_v35;
          let _rhs102 = false;
          let _lhs74 = globalState;
          _686_v34 = _rhs100;
          _688_v36 = _rhs101;
          _lhs74.f2 = _rhs102;
          _659_v14 = _659_v14;
          _659_v14 = _659_v14;
          _684_v32 = _684_v32;
        } else if (_source8.is_DC16) {
          let _689___mcc_h2 = (_source8).cf27;
          let _690_cf27 = _689___mcc_h2;
          let _691_v37;
          _691_v37 = _dafny.Set.fromElements(_this.f31);
          let _692_v38;
          _692_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f24, (_this).f24),_691_v37);
          _692_v38 = (_692_v38).update(new BigNumber(37), _691_v37);
          let _693_v39;
          _693_v39 = _dafny.Seq.UnicodeFromString("h");
          let _694_v40;
          let _nw99 = new _module.C0();
          _nw99.__ctor(((_dafny.ZERO).minus((_this).f24)).isLessThan(new BigNumber((_693_v39).length)));
          _694_v40 = _nw99;
          let _695_v41;
          _695_v41 = _dafny.Seq.of((_694_v40).f34);
          let _696_v42;
          _696_v42 = new _dafny.CodePoint('g'.codePointAt(0));
          let _697_v43;
          _697_v43 = _module.D2.create_DC8();
          let _698_v44;
          _698_v44 = _dafny.MultiSet.fromElements(_697_v43);
          let _699_v45;
          _699_v45 = _dafny.MultiSet.fromElements((_this).f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_694_v40).f34,_696_v42)).length), (_dafny.ZERO).minus(new BigNumber((_698_v44).cardinality())));
          (globalState).f21 = _dafny.Seq.update(((((_694_v40).f34) && (!((_694_v40).f34))) ? (_dafny.Seq.Concat(_module.__default.fm23(globalState), _695_v41)) : ((((_694_v40).f34) ? (_695_v41) : (_695_v41)))), _module.__default.safeIndex(((_this).f24).multipliedBy((((_699_v45).contains((_this).f24)) ? ((_699_v45).get((_this).f24)) : ((_this).f24))), new BigNumber((((((_694_v40).f34) && (!((_694_v40).f34))) ? (_dafny.Seq.Concat(_module.__default.fm23(globalState), _695_v41)) : ((((_694_v40).f34) ? (_695_v41) : (_695_v41))))).length)), _dafny.Seq.IsPrefixOf(_693_v39, _693_v39));
          _693_v39 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(663)), ((_700_v42) => function (_701_i7) {
            return _700_v42;
          })(_696_v42));
        } else {
          let _702___mcc_h3 = (_source8).cf30;
          let _703_cf30 = _702___mcc_h3;
          (globalState).f2 = _this.f31;
          (_this).f31 = _this.f31;
          let _704_v46;
          _704_v46 = _dafny.Seq.UnicodeFromString("ea");
          let _rhs103 = _704_v46;
          let _rhs104 = ((_this.f31) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), function (_705_i8) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })) : (_dafny.Seq.UnicodeFromString("hdlwah")));
          let _rhs105 = _704_v46;
          let _rhs106 = (new BigNumber(-216)).isEqualTo((_this).f24);
          let _lhs75 = globalState;
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          _lhs75.f17 = _rhs103;
          _lhs76.f17 = _rhs104;
          _lhs77.f17 = _rhs105;
          r0 = _rhs106;
          let _706_v47;
          _706_v47 = _module.D8.create_DC24((_this).f24);
          let _707_v48;
          _707_v48 = new _dafny.CodePoint('f'.codePointAt(0));
          let _708_v49;
          _708_v49 = _dafny.Map.Empty.slice().updateUnsafe(_707_v48,_707_v48);
          let _709_v51;
          _709_v51 = _dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll21 = new _dafny.Set();
            for (const _compr_21 of (_708_v49).Keys.Elements) {
              let _710_v50 = _compr_21;
              if ((_708_v49).contains(_710_v50)) {
                _coll21.add(_710_v50);
              }
            }
            return _coll21;
          }()).length), (_this).f24);
          let _711_v52;
          _711_v52 = _dafny.Set.fromElements(_this.f31, true, _this.f31);
          _706_v47 = _module.D8.create_DC24((((_709_v51).contains((_this).f24)) ? ((_709_v51).get((_this).f24)) : (new BigNumber((_711_v52).length))));
        }
      } else {
        let _712_v53;
        _712_v53 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f24).minus((_this).f24),(_this).f24);
        (globalState).f22 = _712_v53;
        let _713_v54;
        _713_v54 = _dafny.Seq.UnicodeFromString("o");
        let _714_v55;
        _714_v55 = _dafny.Seq.of((_this).f24, new BigNumber((_713_v54).length), new BigNumber(920));
        if (_dafny.Seq.IsProperPrefixOf(_714_v55, _dafny.Seq.of((_this).f24, (_this).f24, (_dafny.ZERO).minus((_this).f24), (_this).f24, (_this).f24))) {
          let _715_v56;
          _715_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,new BigNumber(170));
          let _716_v57;
          _716_v57 = _module.D4.create_DC13(((_this).f24).minus(new BigNumber(511)), (_715_v56).Merge(_715_v56), _713_v54);
          _716_v57 = _716_v57;
          (globalState).f18 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f24));
          _712_v53 = ((true) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)) : (_712_v53));
          let _717_v58;
          _717_v58 = _dafny.MultiSet.fromElements(true, _this.f31);
          let _718_v59;
          _718_v59 = new _dafny.CodePoint('t'.codePointAt(0));
          let _719_v60;
          _719_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_713_v54);
          let _720_v61;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_this.f31);
          _720_v61 = _nw100;
          let _721_v62;
          _721_v62 = _dafny.Seq.of(_720_v61, _720_v61);
          let _722_v63;
          _722_v63 = _dafny.Set.fromElements((_this).f24, new BigNumber((_721_v62).length));
          let _723_v64;
          let _nw101 = Array((new BigNumber(15)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _this.f31;
          _nw101[(_dafny.ONE).toNumber()] = !(false);
          _nw101[(new BigNumber(2)).toNumber()] = (_module.D6.create_DC17(_this.f31)).dtor_cf28;
          _nw101[(new BigNumber(3)).toNumber()] = _this.f31;
          _nw101[(new BigNumber(4)).toNumber()] = (_this.f31) === (_this.f31);
          _nw101[(new BigNumber(5)).toNumber()] = (_717_v58).IsDisjointFrom(_717_v58);
          _nw101[(new BigNumber(6)).toNumber()] = _this.f31;
          _nw101[(new BigNumber(7)).toNumber()] = _module.__default.fm2(new BigNumber(211), new BigNumber((_714_v55).length), _718_v59, globalState);
          _nw101[(new BigNumber(8)).toNumber()] = ((_this).f24).isLessThanOrEqualTo((_this).f24);
          _nw101[(new BigNumber(9)).toNumber()] = !((_719_v60).update((_this).f24, _713_v54)).contains((_this).f24);
          _nw101[(new BigNumber(10)).toNumber()] = !(_module.__default.fm30(_718_v59, !(_this.f31), globalState)).contains(new BigNumber((_722_v63).length));
          _nw101[(new BigNumber(11)).toNumber()] = !(new BigNumber(755)).isEqualTo((_this).f24);
          _nw101[(new BigNumber(12)).toNumber()] = true;
          _nw101[(new BigNumber(13)).toNumber()] = _this.f31;
          _nw101[(new BigNumber(14)).toNumber()] = (_720_v61).f34;
          _723_v64 = _nw101;
          _723_v64 = _723_v64;
          let _724_v65;
          _724_v65 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), function (_725_i9) {
            return (_this).f24;
          })).length), (_this).f24, _module.__default.fm0((_this).f24, (_this).f24, globalState));
          let _726_v66;
          _726_v66 = _dafny.Set.fromElements(_724_v65);
          _726_v66 = _726_v66;
        } else {
          let _rhs107 = false;
          let _rhs108 = _this.f31;
          let _rhs109 = new BigNumber(613);
          let _lhs78 = globalState;
          let _lhs79 = globalState;
          let _lhs80 = globalState;
          _lhs78.f2 = _rhs107;
          _lhs79.f2 = _rhs108;
          _lhs80.f1 = _rhs109;
          let _727_v67;
          let _init16 = ((_728_v55) => function (_729_i10) {
            return _dafny.areEqual(_728_v55, _dafny.Seq.of((_this).f24));
          })(_714_v55);
          let _nw102 = Array((new BigNumber(15)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw102.length); _i0_16++) {
            _nw102[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _727_v67 = _nw102;
          _727_v67 = _727_v67;
          let _rhs110 = _713_v54;
          let _rhs111 = (((_this).f24).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ljamq")).length))).minus(_module.__default.safeDivisionInt((_this).f24, (_this).f24));
          let _lhs81 = globalState;
          let _lhs82 = globalState;
          _lhs81.f17 = _rhs110;
          _lhs82.f18 = _rhs111;
          let _730_v68;
          let _nw103 = new _module.C0();
          _nw103.__ctor(_this.f31);
          _730_v68 = _nw103;
          (_this).f31 = !(true) || ((_730_v68).f34);
        }
        if (_this.f31) {
          let _731_v69;
          _731_v69 = _dafny.Seq.of(_713_v54);
          (_this).f31 = !_dafny.areEqual(_713_v54, _dafny.Seq.Concat((_731_v69)[_module.__default.safeIndex((_this).f24, new BigNumber((_731_v69).length))], _713_v54));
          let _732_v70;
          _732_v70 = _module.D4.create_DC12(_714_v55);
          _714_v55 = (_732_v70).dtor_cf21;
          let _733_v71;
          _733_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(new BigNumber(340)).minus((_this).f24));
          _733_v71 = (_733_v71).update(false, (_this).f24);
          (globalState).f14 = (_714_v55)[_module.__default.safeIndex((_this).f24, new BigNumber((_714_v55).length))];
          (globalState).f2 = !(_dafny.areEqual(_714_v55, _dafny.Seq.of((_this).f24, (_this).f24, (_this).f24)));
        } else {
          let _734_v72;
          _734_v72 = _dafny.Set.fromElements((_this).f24);
          let _735_v73;
          _735_v73 = _dafny.Set.fromElements(_734_v72);
          (globalState).f18 = new BigNumber((_735_v73).length);
          let _736_v74;
          _736_v74 = _dafny.MultiSet.fromElements((_this).f24);
          let _737_v75;
          let _nw104 = new _module.C1();
          _nw104.__ctor(_dafny.Seq.update(_dafny.Seq.of(_736_v74), _module.__default.safeIndex(new BigNumber((_714_v55).length), new BigNumber((_dafny.Seq.of(_736_v74)).length)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_738_i11) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })).length))), (_this).f24, (_this).f24, (_this).f25);
          _737_v75 = _nw104;
          let _739_v76;
          _739_v76 = _dafny.Map.Empty.slice().updateUnsafe(false,_737_v75);
          let _740_v77;
          _740_v77 = _737_v75;
          let _741_v78;
          let _nw105 = Array((new BigNumber(19)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = _737_v75;
          _nw105[(_dafny.ONE).toNumber()] = _737_v75;
          _nw105[(new BigNumber(2)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(3)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(4)).toNumber()] = (((_739_v76).contains(_this.f31)) ? ((_739_v76).get(_this.f31)) : (_737_v75));
          _nw105[(new BigNumber(5)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(6)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(7)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(8)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(9)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(10)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(11)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(12)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(13)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(14)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(15)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(16)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(17)).toNumber()] = _737_v75;
          _nw105[(new BigNumber(18)).toNumber()] = (_740_v77);
          _741_v78 = _nw105;
          let _742_v79;
          _742_v79 = _dafny.Seq.of(_737_v75);
          let _743_v80;
          _743_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f31);
          let _index83 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_741_v78).length));
          (_741_v78)[_index83] = ((_this.f31) ? ((_742_v79)[_module.__default.safeIndex(new BigNumber((_743_v80).length), new BigNumber((_742_v79).length))]) : (_737_v75));
          let _744_v81;
          _744_v81 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-79))),_737_v75);
          let _index84 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_741_v78).length));
          (_741_v78)[_index84] = ((_this.f31) ? ((((_744_v81).contains((_this).f24)) ? ((_744_v81).get((_this).f24)) : (_737_v75))) : (((!(_this.f31)) ? (_737_v75) : (_737_v75))));
          (globalState).f2 = !(!(!(_this.f31)));
          (globalState).f9 = (_737_v75).f24;
          let _745_v82;
          _745_v82 = _dafny.Set.fromElements(_this.f31);
          let _746_v83;
          _746_v83 = _dafny.Map.Empty.slice().updateUnsafe((_737_v75).f24,_this.f31);
          let _747_v84;
          _747_v84 = new _dafny.CodePoint('x'.codePointAt(0));
          let _748_v85;
          _748_v85 = _dafny.Set.fromElements(_747_v84, _747_v84);
          let _749_v86;
          let _nw106 = Array((new BigNumber(12)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = _745_v82;
          _nw106[(_dafny.ONE).toNumber()] = _745_v82;
          _nw106[(new BigNumber(2)).toNumber()] = _745_v82;
          _nw106[(new BigNumber(3)).toNumber()] = (_module.__default.fm28(_this.f31, (_this).f24, _this.f31, (_dafny.ZERO).minus((_737_v75).f24), globalState)).Union(_module.__default.fm28((((_746_v83).contains((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f24)))) ? ((_746_v83).get((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f24)))) : (_this.f31)), new BigNumber((_748_v85).length), _this.f31, (_this).f24, globalState));
          _nw106[(new BigNumber(4)).toNumber()] = (_745_v82).Difference(_745_v82);
          _nw106[(new BigNumber(5)).toNumber()] = _745_v82;
          _nw106[(new BigNumber(6)).toNumber()] = _745_v82;
          _nw106[(new BigNumber(7)).toNumber()] = _745_v82;
          _nw106[(new BigNumber(8)).toNumber()] = _module.__default.fm28(_this.f31, (_this).f24, _this.f31, new BigNumber(503), globalState);
          _nw106[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_this.f31, true, _this.f31);
          _nw106[(new BigNumber(10)).toNumber()] = (_745_v82).Difference(_745_v82);
          _nw106[(new BigNumber(11)).toNumber()] = _745_v82;
          _749_v86 = _nw106;
          let _index85 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_749_v86).length));
          (_749_v86)[_index85] = _745_v82;
          let _index86 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_749_v86).length));
          (_749_v86)[_index86] = _745_v82;
        }
        (globalState).f2 = _this.f31;
        (globalState).f17 = _dafny.Seq.UnicodeFromString("tlgqytw");
      }
      let _750_v87;
      let _init17 = function (_751_i12) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      };
      let _nw107 = Array((new BigNumber(26)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw107.length); _i0_17++) {
        _nw107[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _750_v87 = _nw107;
      let _752_v88;
      _752_v88 = new _dafny.CodePoint('y'.codePointAt(0));
      let _index87 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length));
      (_750_v87)[_index87] = _752_v88;
      let _753_v89;
      _753_v89 = _dafny.Seq.UnicodeFromString("pe");
      let _index88 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length));
      let _rhs112 = (_this).f24;
      let _rhs113 = (_this).f24;
      let _rhs114 = _752_v88;
      let _rhs115 = _753_v89;
      let _rhs116 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).f24, (_this).f24), _module.__default.safeModuloInt((_this).f24, new BigNumber((_753_v89).length)));
      let _lhs83 = globalState;
      let _lhs84 = globalState;
      let _lhs85 = _750_v87;
      let _lhs86 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length));
      let _lhs87 = globalState;
      let _lhs88 = globalState;
      _lhs83.f9 = _rhs112;
      _lhs84.f14 = _rhs113;
      _lhs85[_lhs86] = _rhs114;
      _lhs87.f17 = _rhs115;
      _lhs88.f9 = _rhs116;
      let _hi4 = (_this).f24;
      for (let _754_i13 = (_this).f24; _754_i13.isLessThan(_hi4); _754_i13 = _754_i13.plus(_dafny.ONE)) {
        let _755_v90;
        _755_v90 = _dafny.Seq.of(_this.f31, _this.f31, _this.f31, _this.f31);
        let _756_v91;
        _756_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).fm9(_module.__default.fm2(_754_i13, new BigNumber(968), (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], globalState), new BigNumber((_755_v90).length), _754_i13, globalState));
        _756_v91 = (_756_v91).update(_this.f31, _this.f31);
        if ((_module.__default.safeDivisionInt((_this).f24, _754_i13)).isEqualTo(new BigNumber((_module.__default.fm18(globalState)).length))) {
          let _757_v92;
          let _nw108 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _757_v92 = _nw108;
          let _index89 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_757_v92).length));
          (_757_v92)[_index89] = (_dafny.ZERO).minus(_754_i13);
          let _index90 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_757_v92).length));
          (_757_v92)[_index90] = (_this).f24;
          let _index91 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_757_v92).length));
          let _index92 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_757_v92).length));
          let _rhs117 = (_this).f24;
          let _rhs118 = _module.__default.safeModuloInt((_this).f24, new BigNumber(225));
          let _lhs89 = _757_v92;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_757_v92).length));
          let _lhs91 = _757_v92;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_757_v92).length));
          _lhs89[_lhs90] = _rhs117;
          _lhs91[_lhs92] = _rhs118;
          (globalState).f2 = (_this).fm9(_this.f31, _754_i13, (new BigNumber((_module.__default.fm4(_this.f31, _754_i13, new _dafny.CodePoint('g'.codePointAt(0)), _this.f31, globalState)).length)).multipliedBy((_this).f24), globalState);
          let _758_v93;
          let _nw109 = Array((new BigNumber(26)).toNumber()).fill(false);
          _758_v93 = _nw109;
          let _759_v94;
          let _nw110 = Array((new BigNumber(13)).toNumber()).fill(_module.D6.Default());
          _759_v94 = _nw110;
          let _760_v95;
          _760_v95 = _dafny.Set.fromElements((_757_v92)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_757_v92).length))], _754_i13, new BigNumber((_753_v89).length));
          let _index93 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_759_v94).length));
          (_759_v94)[_index93] = _module.D6.create_DC16(_760_v95);
          let _761_v96;
          _761_v96 = _module.D1.create_DC5(_this.f31, _this.f31, _this.f31);
          let _762_v97;
          _762_v97 = _module.D7.create_DC21(_761_v96, _module.__default.fm31(_this.f31, _this.f31, globalState), _758_v93);
          let _index94 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_759_v94).length));
          let _rhs119 = ((_this.f31) ? (_758_v93) : ((_762_v97).dtor_cf34));
          let _rhs120 = _module.D6.create_DC16(_760_v95);
          let _lhs93 = _759_v94;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_759_v94).length));
          _758_v93 = _rhs119;
          _lhs93[_lhs94] = _rhs120;
          let _763_v98;
          let _nw111 = Array((new BigNumber(2)).toNumber()).fill([]);
          _763_v98 = _nw111;
          let _index95 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_763_v98).length));
          (_763_v98)[_index95] = ((_this.f31) ? (_757_v92) : (_757_v92));
          let _764_v99;
          _764_v99 = _module.D0.create_DC0(_757_v92);
          let _index96 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_763_v98).length));
          (_763_v98)[_index96] = (_764_v99).dtor_cf0;
          (globalState).f9 = (_754_i13).multipliedBy(((_this).f24).minus((_757_v92)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_757_v92).length))]));
        } else {
          let _765_v100;
          _765_v100 = _dafny.MultiSet.fromElements(_754_i13, _754_i13);
          let _766_v101;
          _766_v101 = _dafny.Seq.of(_765_v100);
          let _767_v102;
          _767_v102 = _dafny.MultiSet.fromElements(_this.f31, _this.f31);
          let _768_v103;
          _768_v103 = _dafny.Set.fromElements(!(_this.f31));
          let _769_v104;
          _769_v104 = _dafny.Map.Empty.slice().updateUnsafe(_752_v88,(_this).f24);
          let _770_v105;
          let _nw112 = new _module.C1();
          _nw112.__ctor(_dafny.Seq.Concat(_766_v101, _766_v101), _754_i13, _module.__default.safeModuloInt(_754_i13, (((_767_v102).contains(_this.f31)) ? ((_767_v102).get(_this.f31)) : (new BigNumber((_768_v103).length)))), ((true) ? (_dafny.Seq.of(_769_v104, _769_v104, _769_v104)) : (_dafny.Seq.of(_769_v104))));
          _770_v105 = _nw112;
          (globalState).f2 = _this.f31;
          _767_v102 = _767_v102;
          let _771_v106;
          _771_v106 = _dafny.Set.fromElements(_754_i13);
          (globalState).f1 = _module.__default.safeDivisionInt(_770_v105.f33, _module.__default.safeModuloInt((_dafny.ZERO).minus(_754_i13), new BigNumber((_771_v106).length)));
          (globalState).f18 = (_this).f24;
        }
        if (!((_this).f24).isEqualTo(_754_i13)) {
          let _772_v107;
          let _nw113 = new _module.C0();
          _nw113.__ctor(_this.f31);
          _772_v107 = _nw113;
          let _773_v108;
          _773_v108 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_754_i13);
          (globalState).f17 = ((_module.__default.fm2(new BigNumber((_773_v108).length), new BigNumber(392), _module.__default.fm3(globalState), globalState)) ? (_753_v89) : (_753_v89));
          (globalState).f18 = _754_i13;
          let _774_v109;
          let _init18 = ((_775_v90, _776_v107) => function (_777_i14) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_775_v90), _dafny.Seq.of(_775_v90, _dafny.Seq.of((_776_v107).f34, (_776_v107).f34, false)));
          })(_755_v90, _772_v107);
          let _nw114 = Array((new BigNumber(5)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw114.length); _i0_18++) {
            _nw114[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _774_v109 = _nw114;
          let _778_v110;
          _778_v110 = _dafny.Seq.of(_755_v90);
          let _index97 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_774_v109).length));
          (_774_v109)[_index97] = _dafny.Seq.Concat(_778_v110, _778_v110);
          let _779_v111;
          _779_v111 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_755_v90);
          let _index98 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_774_v109).length));
          (_774_v109)[_index98] = _dafny.Seq.update((((_772_v107).f34) ? (_dafny.Seq.Concat(_778_v110, _778_v110)) : (((!((_772_v107).f34)) ? (_778_v110) : (_dafny.Seq.of(_dafny.Seq.of((_772_v107).f34), (((_779_v111).contains(new BigNumber(370))) ? ((_779_v111).get(new BigNumber(370))) : (_755_v90)), _755_v90, _755_v90))))), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("q")).length)), new BigNumber(((((_772_v107).f34) ? (_dafny.Seq.Concat(_778_v110, _778_v110)) : (((!((_772_v107).f34)) ? (_778_v110) : (_dafny.Seq.of(_dafny.Seq.of((_772_v107).f34), (((_779_v111).contains(new BigNumber(370))) ? ((_779_v111).get(new BigNumber(370))) : (_755_v90)), _755_v90, _755_v90)))))).length)), _755_v90);
          _755_v90 = _755_v90;
        } else {
          let _780_v112;
          _780_v112 = _dafny.MultiSet.fromElements((_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], _752_v88);
          let _781_v113;
          _781_v113 = _dafny.Seq.of((_this).f24, (_this).f24);
          (globalState).f14 = _module.__default.fm0(_module.__default.fm0(new BigNumber(233), _module.__default.fm0((_this).f24, (((_780_v112).contains(_752_v88)) ? ((_780_v112).get(_752_v88)) : (new BigNumber((_781_v113).length))), globalState), globalState), (_this).f24, globalState);
          let _782_v114;
          _782_v114 = _dafny.Map.Empty.slice().updateUnsafe(_753_v89,_dafny.Seq.of(_this.f31, _this.f31));
          let _783_v115;
          _783_v115 = _dafny.Set.fromElements((((_782_v114).contains(_753_v89)) ? ((_782_v114).get(_753_v89)) : (_755_v90)), _755_v90, _module.__default.fm23(globalState));
          let _784_v116;
          _784_v116 = _dafny.Seq.of(_755_v90, _755_v90);
          _783_v115 = function () {
            let _coll22 = new _dafny.Set();
            for (const _compr_22 of (_784_v116).Elements) {
              let _785_v117 = _compr_22;
              if (_dafny.Seq.contains(_784_v116, _785_v117)) {
                _coll22.add(_785_v117);
              }
            }
            return _coll22;
          }();
          let _786_v118;
          _786_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_783_v115).length));
          let _787_v119;
          let _out23;
          _out23 = _module.__default.m0((((_786_v118).contains((_this).f24)) ? ((_786_v118).get((_this).f24)) : (new BigNumber(-202))), (_dafny.ZERO).minus(new BigNumber((_756_v91).length)), globalState);
          _787_v119 = _out23;
          (globalState).f1 = ((_this.f31) ? (_754_i13) : (_module.__default.safeDivisionInt(_787_v119, new BigNumber(958))));
          (globalState).f14 = _787_v119;
        }
        let _788_v120;
        _788_v120 = _dafny.MultiSet.fromElements(_this.f31);
        let _789_v121;
        _789_v121 = _module.D6.create_DC18(false);
        let _790_v122;
        _790_v122 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,new BigNumber(513));
        let _791_v123;
        let _nw115 = Array((new BigNumber(11)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = (_789_v121).dtor_cf29;
        _nw115[(_dafny.ONE).toNumber()] = _this.f31;
        _nw115[(new BigNumber(2)).toNumber()] = (_this).fm9(_this.f31, (_this).f24, (_this).f24, globalState);
        _nw115[(new BigNumber(3)).toNumber()] = (_this).fm9(_this.f31, (_this).f24, new BigNumber((_790_v122).length), globalState);
        _nw115[(new BigNumber(4)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(5)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(6)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(7)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(8)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(9)).toNumber()] = _this.f31;
        _nw115[(new BigNumber(10)).toNumber()] = _this.f31;
        _791_v123 = _nw115;
        let _792_v124;
        _792_v124 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_791_v123);
        let _793_v125;
        _793_v125 = _dafny.Map.Empty.slice().updateUnsafe((_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))],((_788_v120).update(!(_this.f31), _module.__default.abs(new BigNumber((_792_v124).length)))).Union(_788_v120));
        _793_v125 = (_793_v125).update(_752_v88, _788_v120);
      }
      let _794_v126;
      _794_v126 = _module.D6.create_DC16(_module.__default.fm18(globalState));
      let _source9 = _module.D6.create_DC19(_794_v126);
      if (_source9.is_DC17) {
        let _795___mcc_h11 = (_source9).cf28;
        let _796_cf28 = _795___mcc_h11;
        let _797_v127;
        let _init19 = function (_798_i15) {
          return (_798_i15).plus((_this).f24);
        };
        let _nw116 = Array((new BigNumber(22)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw116.length); _i0_19++) {
          _nw116[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _797_v127 = _nw116;
        let _index99 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
        (_797_v127)[_index99] = (_this).f24;
        let _799_v128;
        _799_v128 = _dafny.Set.fromElements((_this).f24);
        let _800_v129;
        _800_v129 = _dafny.MultiSet.fromElements(_796_cf28, _this.f31);
        let _801_v131;
        _801_v131 = _dafny.MultiSet.fromElements(_799_v128, _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(51))).length), (_this).f24, new BigNumber((_800_v129).cardinality())), function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(103), new BigNumber(967))) {
            let _802_v130 = _compr_23;
            if (((new BigNumber(103)).isLessThanOrEqualTo(_802_v130)) && ((_802_v130).isLessThan(new BigNumber(967)))) {
              _coll23.add((_802_v130).plus(new BigNumber((_753_v89).length)));
            }
          }
          return _coll23;
        }());
        let _index100 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
        (_797_v127)[_index100] = ((_this.f31) ? ((((_801_v131).contains(_799_v128)) ? ((_801_v131).get(_799_v128)) : ((_this).f24))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_this.f31)).length)));
        if ((_this).fm9(true, new BigNumber(726), (_this).f24, globalState)) {
          let _index101 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length));
          (_750_v87)[_index101] = (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))];
          let _803_v132;
          _803_v132 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), ((_804_v87) => function (_805_i16) {
            return (_804_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_804_v87).length))];
          })(_750_v87)), ((_this.f31) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), ((_806_v88) => function (_807_i17) {
            return _806_v88;
          })(_752_v88))) : (_dafny.Seq.update(_753_v89, _module.__default.safeIndex((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], new BigNumber((_753_v89).length)), new _dafny.CodePoint('o'.codePointAt(0))))), _dafny.Seq.UnicodeFromString("hqys"));
          let _index102 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
          let _rhs121 = _753_v89;
          let _rhs122 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], (_this).f24));
          let _rhs123 = _803_v132;
          let _lhs95 = _797_v127;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
          _753_v89 = _rhs121;
          _lhs95[_lhs96] = _rhs122;
          _803_v132 = _rhs123;
          let _808_v133;
          _808_v133 = _dafny.Seq.of(_796_cf28);
          let _rhs124 = (_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))];
          let _rhs125 = _dafny.Seq.Concat(_808_v133, _808_v133);
          let _lhs97 = globalState;
          let _lhs98 = globalState;
          _lhs97.f1 = _rhs124;
          _lhs98.f20 = _rhs125;
          (_this).f31 = ((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))]).isLessThanOrEqualTo((_this).f24);
          let _809_v134;
          _809_v134 = _dafny.Seq.of(_797_v127, _797_v127);
          let _810_v135;
          _810_v135 = _module.D1.create_DC4(_796_cf28, _this.f31, _752_v88, (_dafny.ZERO).minus((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))]), (_this).f24);
          let _811_v136;
          let _nw117 = Array((new BigNumber(17)).toNumber());
          _nw117[(_dafny.ZERO).toNumber()] = _797_v127;
          _nw117[(_dafny.ONE).toNumber()] = _797_v127;
          _nw117[(new BigNumber(2)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(3)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(4)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(5)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(6)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(7)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(8)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(9)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(10)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(11)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(12)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(13)).toNumber()] = _797_v127;
          _nw117[(new BigNumber(14)).toNumber()] = (_809_v134)[_module.__default.safeIndex((_810_v135).dtor_cf12, new BigNumber((_809_v134).length))];
          _nw117[(new BigNumber(15)).toNumber()] = (_809_v134)[_module.__default.safeIndex((_this).f24, new BigNumber((_809_v134).length))];
          _nw117[(new BigNumber(16)).toNumber()] = _797_v127;
          _811_v136 = _nw117;
          let _index103 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_811_v136).length));
          (_811_v136)[_index103] = _797_v127;
          let _index104 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_811_v136).length));
          (_811_v136)[_index104] = _797_v127;
        } else {
          let _812_v137;
          _812_v137 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
          let _813_v138;
          _813_v138 = _dafny.Map.Empty.slice().updateUnsafe(((_796_cf28) ? (_796_cf28) : (_this.f31)),(_812_v137).Difference(_812_v137));
          (globalState).f9 = new BigNumber((_813_v138).length);
          let _814_v139;
          _814_v139 = _module.D6.create_DC16(_dafny.Set.fromElements(new BigNumber((_753_v89).length)));
          let _815_v140;
          _815_v140 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_814_v139);
          let _816_v141;
          _816_v141 = _dafny.Map.Empty.slice().updateUnsafe(_815_v140,_796_cf28);
          _816_v141 = ((_816_v141).Merge(_816_v141)).Merge(_816_v141);
          let _817_v142;
          _817_v142 = _dafny.Map.Empty.slice().updateUnsafe((_812_v137).update((_this).f24, _module.__default.abs((_this).f24)),_752_v88);
          let _818_v144;
          _818_v144 = _dafny.MultiSet.fromElements((_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))]);
          _817_v142 = (_817_v142).update(_dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of (_818_v144).Elements) {
              let _819_v143 = _compr_24;
              if ((_818_v144).contains(_819_v143)) {
                _coll24.push([_819_v143,_dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))])]);
              }
            }
            return _coll24;
          }()).length)), (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))]);
          let _820_v145;
          let _nw118 = new _module.C0();
          _nw118.__ctor(_this.f31);
          _820_v145 = _nw118;
          let _821_v146;
          let _nw119 = new _module.C1();
          _nw119.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), ((_822_v137) => function (_823_i18) {
            return _822_v137;
          })(_812_v137)), new BigNumber((_dafny.MultiSet.fromElements((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))])).cardinality()), (_this).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), ((_824_v88) => function (_825_i19) {
            return _dafny.Map.Empty.slice().updateUnsafe(_824_v88,new BigNumber(-786));
          })(_752_v88)));
          _821_v146 = _nw119;
        }
        (globalState).f14 = _module.__default.safeModuloInt(new BigNumber(-45), ((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))]).plus(new BigNumber(519)));
        let _826_v147;
        _826_v147 = _dafny.Seq.of(false, _796_cf28, _796_cf28, !(_796_cf28), _796_cf28);
        if (_dafny.Seq.contains(_826_v147, _dafny.Seq.contains(_826_v147, _796_cf28))) {
          let _index105 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length));
          (_750_v87)[_index105] = _752_v88;
          let _827_v148;
          _827_v148 = _dafny.Map.Empty.slice().updateUnsafe(_752_v88,_this.f31);
          _827_v148 = (_827_v148).update((_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], !_dafny.areEqual(_826_v147, _826_v147));
          r0 = !((_this.f31) === (((_this.f31) ? (true) : (_796_cf28))));
          (globalState).f2 = !_dafny.areEqual(_753_v89, _dafny.Seq.Concat(_753_v89, _dafny.Seq.UnicodeFromString("bxrlmgp")));
          (globalState).f1 = new BigNumber((_753_v89).length);
        } else {
          let _828_v149;
          _828_v149 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))]);
          let _829_v150;
          _829_v150 = _dafny.Set.fromElements((((_828_v149).contains(_796_cf28)) ? ((_828_v149).get(_796_cf28)) : (_752_v88)));
          let _830_v151;
          _830_v151 = _module.D3.create_DC11((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], _796_cf28, _module.__default.fm2(new BigNumber((_829_v150).length), (_this).f24, _module.__default.fm3(globalState), globalState));
          (globalState).f2 = (_830_v151).dtor_cf19;
          (globalState).f18 = ((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))]).minus((_this).f24);
          _796_cf28 = !((_module.__default.safeModuloInt((_this).f24, (_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))])).isLessThanOrEqualTo((_module.__default.fm0((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], (_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], globalState)).minus((_this).f24)));
          let _831_v152;
          _831_v152 = _dafny.MultiSet.fromElements((_this).f24, (_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], (_this).f24);
          let _832_v153;
          let _nw120 = new _module.C1();
          _nw120.__ctor(_dafny.Seq.of(_831_v152, (_831_v152).update(new BigNumber((_831_v152).cardinality()), _module.__default.abs(new BigNumber((_826_v147).length))), (_831_v152).update((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], _module.__default.abs((_this).f24)), _831_v152, _831_v152), new BigNumber((_753_v89).length), (_this).f24, (_this).f25);
          _832_v153 = _nw120;
          let _index106 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
          let _rhs126 = _832_v153;
          let _rhs127 = (new BigNumber((_831_v152).cardinality())).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_797_v127,(_832_v153).f24)).length));
          let _rhs128 = _this.f31;
          let _rhs129 = _753_v89;
          let _rhs130 = _module.__default.safeModuloInt((_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))], (_dafny.ZERO).minus((_832_v153).f24));
          let _lhs99 = globalState;
          let _lhs100 = globalState;
          let _lhs101 = _797_v127;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
          _832_v153 = _rhs126;
          _lhs99.f2 = _rhs127;
          r0 = _rhs128;
          _lhs100.f17 = _rhs129;
          _lhs101[_lhs102] = _rhs130;
          let _index107 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length));
          (_797_v127)[_index107] = (_797_v127)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_797_v127).length))];
        }
      } else if (_source9.is_DC18) {
        let _833___mcc_h12 = (_source9).cf29;
        let _834_cf29 = _833___mcc_h12;
        (globalState).f1 = (_this).f24;
        let _835_v154;
        _835_v154 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f24);
        _835_v154 = (_835_v154).update(_834_cf29, (new BigNumber(188)).minus((_this).f24));
        let _836_v155;
        let _nw121 = Array((new BigNumber(7)).toNumber()).fill(false);
        _836_v155 = _nw121;
        let _index108 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_836_v155).length));
        (_836_v155)[_index108] = _this.f31;
        let _837_v156;
        _837_v156 = _dafny.Set.fromElements(_836_v155);
        let _838_v157;
        _838_v157 = _dafny.Seq.of((_this).f24, (_this).f24);
        let _index109 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_836_v155).length));
        let _rhs131 = (_this).f24;
        let _rhs132 = (_838_v157)[_module.__default.safeIndex((_this).f24, new BigNumber((_838_v157).length))];
        let _rhs133 = _834_cf29;
        let _rhs134 = (_837_v156).Union(_837_v156);
        let _lhs103 = globalState;
        let _lhs104 = globalState;
        let _lhs105 = _836_v155;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_836_v155).length));
        _lhs103.f18 = _rhs131;
        _lhs104.f14 = _rhs132;
        _lhs105[_lhs106] = _rhs133;
        _837_v156 = _rhs134;
        (globalState).f14 = (_this).f24;
      } else if (_source9.is_DC16) {
        let _839___mcc_h13 = (_source9).cf27;
        let _840_cf27 = _839___mcc_h13;
        let _841_v158;
        _841_v158 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f24);
        let _842_v159;
        let _out24;
        _out24 = _module.__default.m0(_module.__default.fm0((_this).f24, (_this).f24, globalState), (((_841_v158).contains(_this.f31)) ? ((_841_v158).get(_this.f31)) : (new BigNumber(464))), globalState);
        _842_v159 = _out24;
        let _843_v160;
        _843_v160 = _module.D1.create_DC4(_this.f31, _this.f31, new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(983), (_this).f24);
        let _844_v161;
        _844_v161 = _dafny.Map.Empty.slice().updateUnsafe(_843_v160,_this.f31);
        let _845_v163;
        let _nw122 = Array((new BigNumber(12)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_843_v160,_this.f31);
        _nw122[(_dafny.ONE).toNumber()] = _844_v161;
        _nw122[(new BigNumber(2)).toNumber()] = (_844_v161).Merge(_844_v161);
        _nw122[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_843_v160,_this.f31)).Merge(_844_v161);
        _nw122[(new BigNumber(4)).toNumber()] = (_844_v161).Merge(_844_v161);
        _nw122[(new BigNumber(5)).toNumber()] = _844_v161;
        _nw122[(new BigNumber(6)).toNumber()] = _844_v161;
        _nw122[(new BigNumber(7)).toNumber()] = _844_v161;
        _nw122[(new BigNumber(8)).toNumber()] = function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_844_v161).Keys.Elements) {
            let _846_v162 = _compr_25;
            if ((_844_v161).contains(_846_v162)) {
              _coll25.push([_846_v162,_this.f31]);
            }
          }
          return _coll25;
        }();
        _nw122[(new BigNumber(9)).toNumber()] = (_844_v161).Merge(_844_v161);
        _nw122[(new BigNumber(10)).toNumber()] = (_844_v161).Merge(_dafny.Map.Empty.slice().updateUnsafe(_843_v160,_this.f31));
        _nw122[(new BigNumber(11)).toNumber()] = (_844_v161).Merge(_dafny.Map.Empty.slice().updateUnsafe(_843_v160,_module.__default.fm2(_842_v159, _842_v159, (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], globalState)));
        _845_v163 = _nw122;
        let _index110 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_845_v163).length));
        (_845_v163)[_index110] = (_dafny.Map.Empty.slice().updateUnsafe(_843_v160,_this.f31)).Merge(_844_v161);
        let _847_v164;
        _847_v164 = _dafny.MultiSet.fromElements(_this.f31, _this.f31);
        let _848_v165;
        _848_v165 = _dafny.Seq.of(_this.f31);
        let _849_v166;
        _849_v166 = _module.D4.create_DC13(_842_v159, _841_v158, _dafny.Seq.UnicodeFromString("libuqyvxh"));
        let _index111 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_845_v163).length));
        (_845_v163)[_index111] = _module.__default.fm32((_this).f24, _dafny.Seq.of(_847_v164, _847_v164, _847_v164, _847_v164), _848_v165, _dafny.Seq.update(_753_v89, _module.__default.safeIndex((_849_v166).dtor_cf22, new BigNumber((_753_v89).length)), _752_v88), globalState);
        (globalState).f18 = _module.__default.fm0((_this).f24, _842_v159, globalState);
        let _850_v167;
        let _nw123 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _850_v167 = _nw123;
        let _851_v168;
        let _nw124 = Array((new BigNumber(5)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_753_v89, _753_v89);
        _nw124[(_dafny.ONE).toNumber()] = ((_this.f31) ? (_753_v89) : (_module.__default.fm4(_this.f31, _842_v159, (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], _this.f31, globalState)));
        _nw124[(new BigNumber(2)).toNumber()] = _753_v89;
        _nw124[(new BigNumber(3)).toNumber()] = _753_v89;
        _nw124[(new BigNumber(4)).toNumber()] = _753_v89;
        _851_v168 = _nw124;
        let _852_v169;
        _852_v169 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_this.f31);
        let _853_v170;
        _853_v170 = _dafny.Seq.of(_842_v159);
        let _rhs135 = (_dafny.ZERO).minus((((_this).f24).multipliedBy((((_847_v164).contains(_module.__default.fm2(_module.__default.fm0((_this).f24, _842_v159, globalState), (_this).f24, _752_v88, globalState))) ? ((_847_v164).get(_module.__default.fm2(_module.__default.fm0((_this).f24, _842_v159, globalState), (_this).f24, _752_v88, globalState))) : (_842_v159)))).plus((_842_v159).multipliedBy(_842_v159)));
        let _rhs136 = ((false) ? (_850_v167) : (_850_v167));
        let _rhs137 = (_module.__default.fm16(globalState)).equals((_852_v169).Merge(_852_v169));
        let _rhs138 = ((_853_v170)[_module.__default.safeIndex(_842_v159, new BigNumber((_853_v170).length))]).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_842_v159,_this.f31)).length));
        let _rhs139 = _851_v168;
        let _lhs107 = globalState;
        let _lhs108 = globalState;
        _lhs107.f18 = _rhs135;
        _850_v167 = _rhs136;
        r0 = _rhs137;
        _lhs108.f14 = _rhs138;
        _851_v168 = _rhs139;
      } else {
        let _854___mcc_h14 = (_source9).cf30;
        let _855_cf30 = _854___mcc_h14;
        let _856_v171;
        _856_v171 = _module.D1.create_DC6();
        _856_v171 = _856_v171;
        let _857_v172;
        _857_v172 = _module.D6.create_DC17(_this.f31);
        (globalState).f2 = ((_857_v172).dtor_cf28) && (_this.f31);
        (globalState).f2 = _this.f31;
        let _858_v173;
        _858_v173 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_753_v89, _dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), ((_859_v88) => function (_860_i20) {
          return _859_v88;
        })(_752_v88)))).length), ((_this).f24).minus((_this).f24));
        _858_v173 = _dafny.Seq.of((_this).f24, (_this).f24);
      }
      let _861_v174;
      let _nw125 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
      _861_v174 = _nw125;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_861_v174).length))) {
        let _862_i21 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_862_i21)) && ((_862_i21).isLessThan(new BigNumber((_861_v174).length))))) {
          (_861_v174)[(_862_i21)] = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_module.__default.safeDivisionInt((_this).f24, (_this).f24));
        }
      }
      let _863_i22;
      _863_i22 = _dafny.ZERO;
      L1: {
        while (!(_this.f31)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_863_i22)) {
              break L1;
            }
            _863_i22 = (_863_i22).plus(_dafny.ONE);
            if (_this.f31) {
              let _864_v175;
              let _nw126 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _864_v175 = _nw126;
              let _865_v176;
              _865_v176 = _dafny.Map.Empty.slice().updateUnsafe(_864_v175,_this.f31);
              let _index112 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length));
              (_864_v175)[_index112] = new BigNumber(((_865_v176).Merge((_dafny.Map.Empty.slice().updateUnsafe(_864_v175,_this.f31)).update(_864_v175, _this.f31))).length);
              let _866_v177;
              _866_v177 = _dafny.MultiSet.fromElements((_this).f24);
              let _867_v178;
              _867_v178 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_dafny.Seq.UnicodeFromString("rfeykmiww")).length), (_this).f24, globalState),_this.f31);
              let _index113 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length));
              (_864_v175)[_index113] = ((_this.f31) ? ((((_866_v177).contains((_this).f24)) ? ((_866_v177).get((_this).f24)) : ((_this).f24))) : (new BigNumber((_867_v178).length)));
              (globalState).f9 = (_864_v175)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length))];
              (globalState).f2 = ((new BigNumber(519)).plus((_dafny.ZERO).minus((_864_v175)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length))]))).isLessThan((_864_v175)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length))]);
              let _868_v179;
              _868_v179 = _dafny.Map.Empty.slice().updateUnsafe(_864_v175,(_this).f24);
              let _869_v180;
              _869_v180 = _dafny.Seq.of(_864_v175);
              let _870_v181;
              _870_v181 = _dafny.Set.fromElements(_this.f31);
              _868_v179 = (_868_v179).update((_869_v180)[_module.__default.safeIndex(_module.__default.fm0(_module.__default.fm0(new BigNumber((_870_v181).length), (_864_v175)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length))], globalState), (_this).f24, globalState), new BigNumber((_869_v180).length))], (_this).f24);
              let _871_v182;
              _871_v182 = _dafny.Set.fromElements((_this).f24);
              let _872_v183;
              _872_v183 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,new BigNumber((_871_v182).length));
              (globalState).f2 = (_this).fm9(_this.f31, _module.__default.fm0(new BigNumber((_872_v183).length), (_864_v175)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_864_v175).length))], globalState), (_this).f24, globalState);
            } else {
              let _873_v184;
              _873_v184 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9(_this.f31, (_this).f24, new BigNumber(593), globalState),_this.f31);
              _873_v184 = (_873_v184).update(false, false);
              (globalState).f2 = _this.f31;
              let _874_v185;
              _874_v185 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
              let _875_v186;
              _875_v186 = _dafny.Seq.of(((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_873_v184).length))).update((_this).f24, new BigNumber(-306))).Merge(_874_v185), _874_v185);
              _875_v186 = _dafny.Seq.Concat(_875_v186, _875_v186);
              (globalState).f2 = !(_this.f31) || ((_this.f31) === (_this.f31));
              let _876_v187;
              let _out25;
              _out25 = _module.__default.m0((_this).f24, (_dafny.ZERO).minus((_this).f24), globalState);
              _876_v187 = _out25;
            }
            let _877_v188;
            _877_v188 = _dafny.MultiSet.fromElements((_this).f24);
            let _878_v189;
            let _nw127 = new _module.C0();
            _nw127.__ctor((_dafny.MultiSet.fromElements((_this).f24, new BigNumber((_753_v89).length))).IsSubsetOf(_877_v188));
            _878_v189 = _nw127;
            if (((_this).f24).isLessThan((_this).f24)) {
              let _879_v190;
              _879_v190 = _dafny.Set.fromElements(new BigNumber((_753_v89).length), (_dafny.ZERO).minus((_this).f24));
              let _880_v191;
              _880_v191 = _module.D6.create_DC16(_879_v190);
              _880_v191 = _880_v191;
              let _881_v192;
              _881_v192 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_module.__default.fm18(globalState));
              let _882_v194;
              _882_v194 = _dafny.Seq.of(_879_v190, function () {
                let _coll26 = new _dafny.Set();
                for (const _compr_26 of (_877_v188).Elements) {
                  let _883_v193 = _compr_26;
                  if ((_877_v188).contains(_883_v193)) {
                    _coll26.add((_883_v193).plus(new BigNumber(918)));
                  }
                }
                return _coll26;
              }(), _879_v190);
              let _884_v195;
              _884_v195 = _dafny.Seq.of((_this).f24);
              _879_v190 = (((_881_v192).contains((_this).f24)) ? ((_881_v192).get((_this).f24)) : (((_module.__default.fm2((_this).f24, (_this).f24, (_750_v87)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_750_v87).length))], globalState)) ? ((_882_v194)[_module.__default.safeIndex(new BigNumber((_884_v195).length), new BigNumber((_882_v194).length))]) : (_879_v190))));
              let _885_v196;
              _885_v196 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23(globalState),_module.__default.fm2((_this).f24, (_dafny.ZERO).minus((_this).f24), new _dafny.CodePoint('n'.codePointAt(0)), globalState));
              let _886_v197;
              _886_v197 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_878_v189).f34);
              let _887_v198;
              _887_v198 = _dafny.Seq.of((_878_v189).f34);
              _885_v196 = (_885_v196).update((((((_886_v197).contains((_this).fm9((_878_v189).f34, (_this).f24, (_this).f24, globalState))) ? ((_886_v197).get((_this).fm9((_878_v189).f34, (_this).f24, (_this).f24, globalState))) : ((_878_v189).f34))) ? (_887_v198) : (_887_v198)), (_878_v189).f34);
              let _888_v199;
              _888_v199 = _module.D1.create_DC3(_753_v89);
              _888_v199 = _module.D1.create_DC3(_753_v89);
              (_this).f31 = _this.f31;
            } else {
              let _889_v200;
              _889_v200 = _dafny.Set.fromElements(new BigNumber(104));
              let _890_v201;
              _890_v201 = _dafny.Map.Empty.slice().updateUnsafe((_878_v189).f34,_889_v200);
              let _891_v202;
              _891_v202 = _dafny.Seq.of(!((_878_v189).f34));
              let _892_v203;
              _892_v203 = _dafny.Map.Empty.slice().updateUnsafe(_890_v201,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _891_v202))).cardinality()));
              _892_v203 = (_892_v203).update(_890_v201, (_this).f24);
              let _893_v204;
              let _out26;
              _out26 = _module.__default.m0((_dafny.ZERO).minus((_this).f24), (_this).f24, globalState);
              _893_v204 = _out26;
              let _894_v205;
              let _init20 = function (_895_i23) {
                return !(_this.f31);
              };
              let _nw128 = Array((new BigNumber(26)).toNumber());
              for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw128.length); _i0_20++) {
                _nw128[_i0_20] = _init20(new BigNumber(_i0_20));
              }
              _894_v205 = _nw128;
              let _rhs140 = (new BigNumber((_dafny.Seq.UnicodeFromString("tlpy")).length)).multipliedBy((_dafny.ZERO).minus((_this).f24));
              let _rhs141 = _894_v205;
              _893_v204 = _rhs140;
              _894_v205 = _rhs141;
              let _896_v206;
              _896_v206 = _dafny.Seq.of(_893_v204);
              _896_v206 = _dafny.Seq.Concat(_896_v206, _dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), ((_897_v204) => function (_898_i24) {
                return _897_v204;
              })(_893_v204)));
              (globalState).f2 = (_878_v189).f34;
            }
            let _899_v207;
            let _init21 = ((_900_v189) => function (_901_i25) {
              return (_dafny.MultiSet.fromElements((_900_v189).f34)).Union(_dafny.MultiSet.fromElements(_this.f31));
            })(_878_v189);
            let _nw129 = Array((new BigNumber(23)).toNumber());
            for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw129.length); _i0_21++) {
              _nw129[_i0_21] = _init21(new BigNumber(_i0_21));
            }
            _899_v207 = _nw129;
            _899_v207 = _899_v207;
          }
        }
      }
      r0 = _this.f31;
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
      this.f30 = [];
      this._f29 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f29, f30, f24, f25) {
      let _this = this;
      (_this)._f29 = f29;
      (_this).f30 = f30;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm8(globalState) {
      let _this = this;
      return (_module.D0.create_DC2(false, (_this).f24, new _dafny.CodePoint('h'.codePointAt(0)))).dtor_cf5;
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let _902_v0;
      _902_v0 = _dafny.Seq.UnicodeFromString("olddh");
      let _903_v1;
      _903_v1 = new _dafny.CodePoint('w'.codePointAt(0));
      let _904_v2;
      _904_v2 = _dafny.Seq.of(_dafny.Seq.update(_902_v0, _module.__default.safeIndex((_this).f24, new BigNumber((_902_v0).length)), _903_v1));
      let _hi5 = new BigNumber((_904_v2).length);
      for (let _905_i0 = (_this).f24; _905_i0.isLessThan(_hi5); _905_i0 = _905_i0.plus(_dafny.ONE)) {
        let _906_v3;
        _906_v3 = true;
        (globalState).f17 = _dafny.Seq.Concat(_module.__default.fm4(false, (_this).f24, _903_v1, _906_v3, globalState), _dafny.Seq.UnicodeFromString("pglhejcp"));
        let _907_v4;
        let _init22 = ((_908_v3) => function (_909_i1) {
          return ((_908_v3) ? (_908_v3) : (_908_v3));
        })(_906_v3);
        let _nw130 = Array((new BigNumber(21)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw130.length); _i0_22++) {
          _nw130[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _907_v4 = _nw130;
        let _nw131 = Array((new BigNumber(2)).toNumber()).fill(false);
        _907_v4 = _nw131;
        let _910_v5;
        _910_v5 = _dafny.Seq.of(_906_v3);
        let _911_v6;
        _911_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_906_v3));
        let _912_v7;
        _912_v7 = _dafny.Seq.of(_911_v6, _dafny.MultiSet.fromElements(_910_v5, _910_v5), _911_v6);
        if (((_912_v7)[_module.__default.safeIndex((_this).f24, new BigNumber((_912_v7).length))]).contains(_dafny.Seq.Concat(_910_v5, _910_v5))) {
          let _913_v8;
          let _nw132 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _913_v8 = _nw132;
          let _914_v9;
          _914_v9 = _module.D2.create_DC7(_913_v8);
          let _915_v10;
          let _nw133 = Array((new BigNumber(6)).toNumber());
          _nw133[(_dafny.ZERO).toNumber()] = _913_v8;
          _nw133[(_dafny.ONE).toNumber()] = _913_v8;
          _nw133[(new BigNumber(2)).toNumber()] = _913_v8;
          _nw133[(new BigNumber(3)).toNumber()] = _913_v8;
          _nw133[(new BigNumber(4)).toNumber()] = _913_v8;
          _nw133[(new BigNumber(5)).toNumber()] = (_914_v9).dtor_cf16;
          _915_v10 = _nw133;
          _915_v10 = _915_v10;
          let _916_v11;
          let _nw134 = new _module.C0();
          _nw134.__ctor(!_dafny.areEqual(_910_v5, _dafny.Seq.of(true, false, _906_v3, _906_v3, _906_v3)));
          _916_v11 = _nw134;
          (globalState).f18 = _905_i0;
          let _index114 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_907_v4).length));
          (_907_v4)[_index114] = true;
          let _index115 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_907_v4).length));
          (_907_v4)[_index115] = _dafny.Seq.IsProperPrefixOf(_902_v0, _dafny.Seq.Concat(_dafny.Seq.update(_902_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("apcb")).length), new BigNumber((_902_v0).length)), _903_v1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(426)), function (_917_i2) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })));
          let _918_v12;
          _918_v12 = _dafny.MultiSet.fromElements(!((_916_v11).f34), (_907_v4)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_907_v4).length))], (_916_v11).f34);
          let _919_v13;
          _919_v13 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(711)), ((_920_v1) => function (_921_i3) {
            return _920_v1;
          })(_903_v1)));
          let _rhs142 = (((_918_v12).IsDisjointFrom(_dafny.MultiSet.fromElements((_916_v11).f34))) ? ((_this).f24) : (new BigNumber((_902_v0).length)));
          let _rhs143 = ((_919_v13).IsSubsetOf(_module.__default.fm22(_905_i0, (_916_v11).f34, (_916_v11).f34, new BigNumber((_902_v0).length), globalState))) && ((_916_v11).f34);
          let _rhs144 = (_this).f24;
          let _lhs109 = globalState;
          let _lhs110 = globalState;
          let _lhs111 = globalState;
          _lhs109.f14 = _rhs142;
          _lhs110.f2 = _rhs143;
          _lhs111.f9 = _rhs144;
        } else {
          let _922_v14;
          _922_v14 = _dafny.Map.Empty.slice().updateUnsafe(_902_v0,_906_v3);
          _922_v14 = (_922_v14).update(_902_v0, _906_v3);
          let _923_v15;
          _923_v15 = _dafny.Set.fromElements(_906_v3);
          let _924_v16;
          _924_v16 = _module.D0.create_DC2((new BigNumber((_923_v15).length)).isLessThan(_905_i0), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-589)), ((_925_i0) => function (_926_i4) {
  return _925_i0;
})(_905_i0))).length), new _dafny.CodePoint('g'.codePointAt(0)));
          _924_v16 = _924_v16;
          let _927_v18;
          let _init23 = function (_928_i5) {
            return function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of _dafny.IntegerRange(new BigNumber(291), new BigNumber(-29))) {
                let _929_v17 = _compr_27;
                if (((new BigNumber(291)).isLessThanOrEqualTo(_929_v17)) && ((_929_v17).isLessThan(new BigNumber(-29)))) {
                  _coll27.add(_module.__default.safeModuloInt(_929_v17, (_this).f24));
                }
              }
              return _coll27;
            }();
          };
          let _nw135 = Array((new BigNumber(9)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw135.length); _i0_23++) {
            _nw135[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _927_v18 = _nw135;
          let _rhs145 = (_this).f29;
          let _rhs146 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_902_v0, _902_v0), _902_v0)).length);
          let _lhs112 = globalState;
          _927_v18 = _rhs145;
          _lhs112.f1 = _rhs146;
          let _930_v19;
          _930_v19 = _module.D8.create_DC23(_923_v15);
          _930_v19 = _module.D8.create_DC23(_923_v15);
          let _931_v20;
          _931_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f24, (_this).f24, globalState),(_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_910_v5).length)).plus((_this).f24))));
          _931_v20 = (_931_v20).update(_905_i0, _module.__default.safeModuloInt(_905_i0, (_dafny.ZERO).minus(new BigNumber((_902_v0).length))));
        }
        (globalState).f18 = _module.__default.fm0(_905_i0, ((_this).f24).plus(_905_i0), globalState);
      }
      let _hi6 = (_this).f24;
      for (let _932_i6 = new BigNumber((_902_v0).length); _932_i6.isLessThan(_hi6); _932_i6 = _932_i6.plus(_dafny.ONE)) {
        (globalState).f2 = (new BigNumber((_902_v0).length)).isEqualTo((_932_i6).plus(new BigNumber((_dafny.Seq.UnicodeFromString("nhg")).length)));
        let _933_v21;
        _933_v21 = true;
        let _934_v22;
        _934_v22 = _dafny.Map.Empty.slice().updateUnsafe(_933_v21,_933_v21);
        _934_v22 = (_934_v22).update(!((_this).f24).isEqualTo(_932_i6), _933_v21);
        _902_v0 = _dafny.Seq.Concat(_902_v0, ((_933_v21) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), ((_935_v1) => function (_936_i7) {
          return _935_v1;
        })(_903_v1))) : (_902_v0)));
        (globalState).f2 = _933_v21;
      }
      let _937_v23;
      let _out27;
      _out27 = _module.__default.m0((_this).f24, (_this).f24, globalState);
      _937_v23 = _out27;
      let _938_v24;
      _938_v24 = true;
      let _939_v25;
      _939_v25 = _dafny.MultiSet.fromElements(_938_v24);
      let _940_v26;
      _940_v26 = _dafny.Seq.of(new BigNumber((_939_v25).cardinality()), (_this).f24, _937_v23);
      let _941_v27;
      _941_v27 = _dafny.Map.Empty.slice().updateUnsafe(_938_v24,_940_v26);
      let _942_v28;
      _942_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_941_v27).contains(_938_v24)) ? ((_941_v27).get(_938_v24)) : (_940_v26))).length),_938_v24);
      let _943_v29;
      _943_v29 = _dafny.MultiSet.fromElements(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_942_v28,_937_v23)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_942_v28,_937_v23))).length));
      _943_v29 = (_dafny.MultiSet.fromElements(_937_v23)).Intersect(_943_v29);
      let _source10 = _module.__default.fm33(_938_v24, (_dafny.ZERO).minus((_this).f24), globalState);
      if (_source10.is_DC10) {
        _938_v24 = _938_v24;
        let _944_v30;
        let _nw136 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _944_v30 = _nw136;
        let _945_v31;
        _945_v31 = _dafny.Set.fromElements(_944_v30, _944_v30);
        let _946_v32;
        _946_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(_938_v24),_938_v24);
        let _947_v33;
        _947_v33 = _dafny.Set.fromElements(!(_module.__default.fm2(new BigNumber(872), (_this).f24, new _dafny.CodePoint('x'.codePointAt(0)), globalState)));
        let _948_v34;
        _948_v34 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_940_v26));
        let _949_v35;
        _949_v35 = _dafny.Map.Empty.slice().updateUnsafe(_903_v1,(_dafny.ZERO).minus((_this).f24));
        let _950_v36;
        let _nw137 = new _module.C1();
        _nw137.__ctor(_948_v34, (_this).f24, (_this).f24, _dafny.Seq.of(_949_v35));
        _950_v36 = _nw137;
        let _951_v37;
        _951_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(new BigNumber((_946_v32).length), _937_v23, _903_v1, globalState),(_dafny.ZERO).minus(_937_v23));
        let _952_v38;
        _952_v38 = _dafny.Map.Empty.slice().updateUnsafe(_950_v36,_951_v37);
        let _953_v39;
        let _nw138 = Array((new BigNumber(20)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = new BigNumber((_945_v31).length);
        _nw138[(_dafny.ONE).toNumber()] = ((_this).f24).multipliedBy(_937_v23);
        _nw138[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f24), (_this).f24);
        _nw138[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(((_this).f24).multipliedBy(_937_v23));
        _nw138[(new BigNumber(4)).toNumber()] = (_this).f24;
        _nw138[(new BigNumber(5)).toNumber()] = _937_v23;
        _nw138[(new BigNumber(6)).toNumber()] = new BigNumber(((_946_v32).Merge(_946_v32)).length);
        _nw138[(new BigNumber(7)).toNumber()] = new BigNumber((_947_v33).length);
        _nw138[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("hwuajdbrj")).length), (_this).f24);
        _nw138[(new BigNumber(9)).toNumber()] = new BigNumber((_952_v38).length);
        _nw138[(new BigNumber(10)).toNumber()] = (_module.D1.create_DC4(_938_v24, _938_v24, _903_v1, (_this).f24, (_this).fm8(globalState))).dtor_cf12;
        _nw138[(new BigNumber(11)).toNumber()] = _950_v36.f33;
        _nw138[(new BigNumber(12)).toNumber()] = (_this).f24;
        _nw138[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw138[(new BigNumber(14)).toNumber()] = ((_this).f24).minus(_950_v36.f33);
        _nw138[(new BigNumber(15)).toNumber()] = new BigNumber((_947_v33).length);
        _nw138[(new BigNumber(16)).toNumber()] = (_this).f24;
        _nw138[(new BigNumber(17)).toNumber()] = (_940_v26)[_module.__default.safeIndex(_937_v23, new BigNumber((_940_v26).length))];
        _nw138[(new BigNumber(18)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("hhs")).length)).plus((_dafny.ZERO).minus(_937_v23));
        _nw138[(new BigNumber(19)).toNumber()] = (_this).f24;
        _953_v39 = _nw138;
        let _954_v40;
        _954_v40 = _dafny.Map.Empty.slice().updateUnsafe(_951_v37,(_dafny.ZERO).minus(_950_v36.f33));
        let _955_v41;
        _955_v41 = _dafny.Map.Empty.slice().updateUnsafe(_938_v24,(_951_v37).update(_938_v24, _950_v36.f33));
        let _nw139 = Array((new BigNumber(5)).toNumber());
        _nw139[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw139[(_dafny.ONE).toNumber()] = (((_954_v40).contains((((_955_v41).contains(_938_v24)) ? ((_955_v41).get(_938_v24)) : (_951_v37)))) ? ((_954_v40).get((((_955_v41).contains(_938_v24)) ? ((_955_v41).get(_938_v24)) : (_951_v37)))) : ((_this).f24));
        _nw139[(new BigNumber(2)).toNumber()] = _950_v36.f33;
        _nw139[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw139[(new BigNumber(4)).toNumber()] = _937_v23;
        _953_v39 = _nw139;
        let _956_v42;
        let _nw140 = Array((new BigNumber(25)).toNumber()).fill(false);
        _956_v42 = _nw140;
        let _957_v43;
        _957_v43 = _dafny.Map.Empty.slice().updateUnsafe(_902_v0,_956_v42);
        _956_v42 = (((_957_v43).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-401)), ((_960_v1) => function (_961_i8) {
          return _960_v1;
        })(_903_v1)))) ? ((_957_v43).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-401)), ((_958_v1) => function (_959_i8) {
          return _958_v1;
        })(_903_v1)))) : (_956_v42));
        (globalState).f17 = _dafny.Seq.Concat(_902_v0, _902_v0);
      } else if (_source10.is_DC11) {
        let _962___mcc_h0 = (_source10).cf18;
        let _963___mcc_h1 = (_source10).cf19;
        let _964___mcc_h2 = (_source10).cf20;
        let _965_cf20 = _964___mcc_h2;
        let _966_cf19 = _963___mcc_h1;
        let _967_cf18 = _962___mcc_h0;
        let _968_v44;
        _968_v44 = _dafny.Map.Empty.slice().updateUnsafe(_938_v24,!(!(true)) || (true));
        (globalState).f2 = (((_968_v44).contains(_module.__default.fm2(_937_v23, _937_v23, _903_v1, globalState))) ? ((_968_v44).get(_module.__default.fm2(_937_v23, _937_v23, _903_v1, globalState))) : (_966_cf19));
        let _969_v45;
        let _nw141 = Array((new BigNumber(20)).toNumber());
        _nw141[(_dafny.ZERO).toNumber()] = _966_cf19;
        _nw141[(_dafny.ONE).toNumber()] = _938_v24;
        _nw141[(new BigNumber(2)).toNumber()] = !(_965_cf20);
        _nw141[(new BigNumber(3)).toNumber()] = _966_cf19;
        _nw141[(new BigNumber(4)).toNumber()] = _966_cf19;
        _nw141[(new BigNumber(5)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(6)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(7)).toNumber()] = _966_cf19;
        _nw141[(new BigNumber(8)).toNumber()] = false;
        _nw141[(new BigNumber(9)).toNumber()] = true;
        _nw141[(new BigNumber(10)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(11)).toNumber()] = _965_cf20;
        _nw141[(new BigNumber(12)).toNumber()] = true;
        _nw141[(new BigNumber(13)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(14)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(15)).toNumber()] = !(false);
        _nw141[(new BigNumber(16)).toNumber()] = _938_v24;
        _nw141[(new BigNumber(17)).toNumber()] = _965_cf20;
        _nw141[(new BigNumber(18)).toNumber()] = _965_cf20;
        _nw141[(new BigNumber(19)).toNumber()] = _966_cf19;
        _969_v45 = _nw141;
        let _970_v46;
        _970_v46 = _dafny.Seq.of(_969_v45);
        _970_v46 = _970_v46;
        let _index116 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_969_v45).length));
        (_969_v45)[_index116] = _966_cf19;
        let _971_v47;
        _971_v47 = _dafny.Seq.of(_966_cf19, _938_v24);
        let _972_v48;
        _972_v48 = _dafny.MultiSet.fromElements(_971_v47);
        let _index117 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_969_v45).length));
        (_969_v45)[_index117] = (((_966_cf19) ? (_972_v48) : (_972_v48))).IsSubsetOf(_972_v48);
        let _index118 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_969_v45).length));
        (_969_v45)[_index118] = (_969_v45)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_969_v45).length))];
      } else {
        let _973___mcc_h3 = (_source10).cf17;
        let _974_cf17 = _973___mcc_h3;
        let _975_v49;
        _975_v49 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_937_v23), _dafny.MultiSet.fromElements(new BigNumber((_902_v0).length)), _943_v29, _943_v29, _943_v29);
        let _976_v51;
        let _nw142 = new _module.C1();
        _nw142.__ctor(_975_v49, new BigNumber((_902_v0).length), (_this).f24, _dafny.Seq.of(function () {
          let _coll28 = new _dafny.Map();
          for (const _compr_28 of (_902_v0).Elements) {
            let _977_v50 = _compr_28;
            if (_dafny.Seq.contains(_902_v0, _977_v50)) {
              _coll28.push([_977_v50,_937_v23]);
            }
          }
          return _coll28;
        }()));
        _976_v51 = _nw142;
        let _978_v52;
        _978_v52 = _dafny.Set.fromElements((_940_v26)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_940_v26).length))]);
        let _979_v54;
        _979_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_978_v52).Union(function () {
          let _coll29 = new _dafny.Set();
          for (const _compr_29 of _dafny.IntegerRange(new BigNumber(350), new BigNumber(677))) {
            let _980_v53 = _compr_29;
            if (((new BigNumber(350)).isLessThanOrEqualTo(_980_v53)) && ((_980_v53).isLessThan(new BigNumber(677)))) {
              _coll29.add((_980_v53).multipliedBy(_937_v23));
            }
          }
          return _coll29;
        }())).length),_dafny.Seq.UnicodeFromString("gsixn"));
        let _rhs147 = (((_979_v54).contains(_976_v51.f33)) ? ((_979_v54).get(_976_v51.f33)) : (_902_v0));
        let _rhs148 = ((_module.__default.fm2(new BigNumber((_902_v0).length), new BigNumber(189), _903_v1, globalState)) ? (_937_v23) : ((_this).f24));
        let _rhs149 = false;
        let _lhs113 = globalState;
        let _lhs114 = globalState;
        _lhs113.f17 = _rhs147;
        _lhs114.f14 = _rhs148;
        _938_v24 = _rhs149;
        let _981_v55;
        let _nw143 = Array((new BigNumber(28)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = _938_v24;
        _nw143[(_dafny.ONE).toNumber()] = _938_v24;
        _nw143[(new BigNumber(2)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(3)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(4)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(5)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(6)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(7)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(8)).toNumber()] = ((_this).f24).isEqualTo((_this).f24);
        _nw143[(new BigNumber(9)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements(_938_v24, _938_v24)).IsDisjointFrom((_939_v25).update(false, _module.__default.abs(_976_v51.f33)));
        _nw143[(new BigNumber(11)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(12)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(13)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(14)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(15)).toNumber()] = !(_976_v51.f33).isEqualTo(_976_v51.f33);
        _nw143[(new BigNumber(16)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(17)).toNumber()] = !(_938_v24);
        _nw143[(new BigNumber(18)).toNumber()] = (_974_cf17)[_module.__default.safeIndex(_937_v23, new BigNumber((_974_cf17).length))];
        _nw143[(new BigNumber(19)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(20)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(21)).toNumber()] = ((_938_v24) ? (_938_v24) : (_938_v24));
        _nw143[(new BigNumber(22)).toNumber()] = !(_938_v24);
        _nw143[(new BigNumber(23)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(24)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(25)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(26)).toNumber()] = _938_v24;
        _nw143[(new BigNumber(27)).toNumber()] = _938_v24;
        _981_v55 = _nw143;
        _981_v55 = _981_v55;
        let _982_v56;
        let _nw144 = new _module.C0();
        _nw144.__ctor(_dafny.Seq.IsProperPrefixOf(_902_v0, _902_v0));
        _982_v56 = _nw144;
      }
      let _983_v57;
      _983_v57 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_902_v0).length)),_937_v23);
      let _984_v58;
      _984_v58 = _dafny.Map.Empty.slice().updateUnsafe(_937_v23,_983_v57);
      let _985_v59;
      _985_v59 = _dafny.Map.Empty.slice().updateUnsafe(_938_v24,_937_v23);
      let _986_v60;
      _986_v60 = _module.D4.create_DC13(new BigNumber((_984_v58).length), _985_v59, _902_v0);
      let _987_v61;
      _987_v61 = _dafny.Map.Empty.slice().updateUnsafe(_986_v60,_dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), ((_988_v1) => function (_989_i9) {
        return _988_v1;
      })(_903_v1)));
      (globalState).f1 = new BigNumber((((((_938_v24) ? (_938_v24) : (_938_v24))) ? ((((_987_v61).contains(_986_v60)) ? ((_987_v61).get(_986_v60)) : (_902_v0))) : (_902_v0))).length);
      r0 = true;
      return r0;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _990_v0;
      _990_v0 = _dafny.Seq.of(p0);
      let _991_v1;
      _991_v1 = _module.D3.create_DC9(_dafny.Seq.update(_990_v0, _module.__default.safeIndex(new BigNumber(852), new BigNumber((_990_v0).length)), p0));
      let _992_i0;
      _992_i0 = _dafny.ZERO;
      L2: {
        let _pat_let_tv6 = p0;
        while (function (_source12) {
          if (_source12.is_DC10) {
            return _pat_let_tv6;
          } else if (_source12.is_DC11) {
            let _1069___mcc_h7 = (_source12).cf18;
            let _1070___mcc_h8 = (_source12).cf19;
            let _1071___mcc_h9 = (_source12).cf20;
            let _1072_cf20 = _1071___mcc_h9;
            let _1073_cf19 = _1070___mcc_h8;
            let _1074_cf18 = _1069___mcc_h7;
            return _1073_cf19;
          } else {
            let _1075___mcc_h10 = (_source12).cf17;
            let _1076_cf17 = _1075___mcc_h10;
            return true;
          }
        }(_991_v1)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_992_i0)) {
              break L2;
            }
            _992_i0 = (_992_i0).plus(_dafny.ONE);
            let _993_v2;
            _993_v2 = new _dafny.CodePoint('o'.codePointAt(0));
            let _994_v3;
            _994_v3 = _dafny.Seq.of(_993_v2, _993_v2, _993_v2);
            let _995_v4;
            _995_v4 = _module.D1.create_DC3(_dafny.Seq.update(_994_v3, _module.__default.safeIndex((_this).f24, new BigNumber((_994_v3).length)), _993_v2));
            let _996_v5;
            _996_v5 = _module.D8.create_DC26(_995_v4, _994_v3);
            let _source11 = _996_v5;
            if (_source11.is_DC24) {
              let _997___mcc_h0 = (_source11).cf37;
              let _998_cf37 = _997___mcc_h0;
              let _999_v6;
              _999_v6 = _module.D0.create_DC1(new BigNumber(980), (_this).f24, p0);
              let _1000_v7;
              _1000_v7 = _dafny.Seq.of((_this).f24);
              let _1001_v8;
              _1001_v8 = _module.D1.create_DC5(p0, p0, p0);
              let _1002_v9;
              _1002_v9 = _module.D6.create_DC18(p0);
              let _1003_v10;
              let _init24 = ((_1004_p0) => function (_1005_i1) {
                return _1004_p0;
              })(p0);
              let _nw145 = Array((new BigNumber(17)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw145.length); _i0_24++) {
                _nw145[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _1003_v10 = _nw145;
              let _1006_v11;
              _1006_v11 = _module.D7.create_DC21(_1001_v8, _module.D6.create_DC19(_1002_v9), _1003_v10);
              let _1007_v12;
              _1007_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1000_v7,(_1006_v11).dtor_cf34);
              let _pat_let_tv4 = _990_v0;
              let _pat_let_tv5 = _990_v0;
              let _1008_v13;
              let _nw146 = Array((new BigNumber(27)).toNumber());
              _nw146[(_dafny.ZERO).toNumber()] = _998_cf37;
              _nw146[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((function (_pat_let10_0) {
                return function (_1009_dt__update__tmp_h0) {
                  return function (_pat_let11_0) {
                    return function (_1010_dt__update_hcf3_h0) {
                      return _module.D0.create_DC1((_1009_dt__update__tmp_h0).dtor_cf1, (_1009_dt__update__tmp_h0).dtor_cf2, _1010_dt__update_hcf3_h0);
                    }(_pat_let11_0);
                  }((_pat_let_tv4)[_module.__default.safeIndex((_this).f24, new BigNumber((_pat_let_tv5).length))]);
                }(_pat_let10_0);
              }(_999_v6)).dtor_cf1);
              _nw146[(new BigNumber(2)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(3)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(4)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f24);
              _nw146[(new BigNumber(6)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(7)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(8)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(9)).toNumber()] = new BigNumber((_1007_v12).length);
              _nw146[(new BigNumber(10)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(11)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(12)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(13)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(14)).toNumber()] = new BigNumber((_994_v3).length);
              _nw146[(new BigNumber(15)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(16)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(17)).toNumber()] = new BigNumber(597);
              _nw146[(new BigNumber(18)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(19)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(20)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(21)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(22)).toNumber()] = new BigNumber(-506);
              _nw146[(new BigNumber(23)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(24)).toNumber()] = (_this).f24;
              _nw146[(new BigNumber(25)).toNumber()] = _998_cf37;
              _nw146[(new BigNumber(26)).toNumber()] = (_this).f24;
              _1008_v13 = _nw146;
              let _1011_v14;
              _1011_v14 = _dafny.Seq.of(_1008_v13, ((p0) ? (_1008_v13) : (_1008_v13)), _1008_v13);
              let _index119 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1008_v13).length));
              (_1008_v13)[_index119] = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(_998_cf37)))).length);
              let _index120 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1008_v13).length));
              (_1008_v13)[_index120] = _998_cf37;
              let _1012_v15;
              _1012_v15 = _dafny.Map.Empty.slice().updateUnsafe(_993_v2,(_this).f24);
              let _1013_v16;
              let _nw147 = new _module.C2();
              _nw147.__ctor(p0, (_this).f24, _dafny.Seq.of(_1012_v15));
              _1013_v16 = _nw147;
              let _1014_v17;
              _1014_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1013_v16,_1008_v13);
              let _1015_v18;
              let _nw148 = new _module.C0();
              _nw148.__ctor(p0);
              _1015_v18 = _nw148;
              let _1016_v19;
              _1016_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1015_v18,(_1015_v18).f34);
              let _1017_v20;
              _1017_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1013_v16).f24,true);
              let _1018_v21;
              _1018_v21 = _dafny.Set.fromElements(_1000_v7, _1000_v7, _1000_v7);
              let _1019_v22;
              _1019_v22 = _dafny.Set.fromElements(!((_1015_v18).f34));
              let _index121 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1008_v13).length));
              let _index122 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1008_v13).length));
              let _rhs150 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1011_v14, _dafny.Seq.of(_1008_v13)), _1011_v14);
              let _rhs151 = ((_this).f24).minus(new BigNumber((_1014_v17).length));
              let _rhs152 = _dafny.Seq.update(_dafny.Seq.update(_module.__default.fm4(p0, new BigNumber(((_1016_v19).Merge(_1016_v19)).length), _993_v2, p0, globalState), _module.__default.safeIndex((_1000_v7)[_module.__default.safeIndex((_this).f24, new BigNumber((_1000_v7).length))], new BigNumber((_module.__default.fm4(p0, new BigNumber(((_1016_v19).Merge(_1016_v19)).length), _993_v2, p0, globalState)).length)), _993_v2), _module.__default.safeIndex(new BigNumber((_994_v3).length), new BigNumber((_dafny.Seq.update(_module.__default.fm4(p0, new BigNumber(((_1016_v19).Merge(_1016_v19)).length), _993_v2, p0, globalState), _module.__default.safeIndex((_1000_v7)[_module.__default.safeIndex((_this).f24, new BigNumber((_1000_v7).length))], new BigNumber((_module.__default.fm4(p0, new BigNumber(((_1016_v19).Merge(_1016_v19)).length), _993_v2, p0, globalState)).length)), _993_v2)).length)), _993_v2);
              let _rhs153 = !((_1015_v18).f34) || (p0);
              let _rhs154 = _module.__default.safeModuloInt((_1000_v7)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1017_v20).length), (_this).f24, (_this).f24, new BigNumber((_1018_v21).length))).length), new BigNumber((_1000_v7).length))], new BigNumber((_1019_v22).length));
              let _lhs115 = _1008_v13;
              let _lhs116 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1008_v13).length));
              let _lhs117 = globalState;
              let _lhs118 = _1008_v13;
              let _lhs119 = _module.__default.safeIndex(new BigNumber(583), new BigNumber((_1008_v13).length));
              _1011_v14 = _rhs150;
              _lhs115[_lhs116] = _rhs151;
              _lhs117.f17 = _rhs152;
              r1 = _rhs153;
              _lhs118[_lhs119] = _rhs154;
              let _index123 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1008_v13).length));
              (_1008_v13)[_index123] = (_1013_v16).f24;
              (globalState).f1 = ((_module.D0.create_DC2(!((_1015_v18).f34), (_this).f24, _993_v2)).dtor_cf5).multipliedBy(_998_cf37);
              (globalState).f2 = !((_1015_v18).f34);
            } else if (_source11.is_DC25) {
              let _1020___mcc_h1 = (_source11).cf38;
              let _1021___mcc_h2 = (_source11).cf39;
              let _1022___mcc_h3 = (_source11).cf40;
              let _1023_cf40 = _1022___mcc_h3;
              let _1024_cf39 = _1021___mcc_h2;
              let _1025_cf38 = _1020___mcc_h1;
              let _1026_v23;
              _1026_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1025_cf38,_1023_cf40);
              let _1027_v24;
              _1027_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1024_cf39,new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1023_cf40, _1023_cf40, !(!(p0))), _module.__default.safeIndex(_1025_cf38, new BigNumber((_dafny.Seq.of(_1023_cf40, _1023_cf40, !(!(p0)))).length)), (((_1026_v23).contains((_this).f24)) ? ((_1026_v23).get((_this).f24)) : (_1023_cf40)))).length));
              let _1028_v25;
              _1028_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1025_cf38,new BigNumber(679));
              _1027_v24 = (_1027_v24).update((_1024_cf39) && (p0), _module.__default.safeDivisionInt((_this).f24, new BigNumber((_1028_v25).length)));
              (globalState).f1 = _module.__default.safeDivisionInt(_1025_cf38, _module.__default.safeDivisionInt(_1025_cf38, (_this).f24));
              let _rhs155 = p0;
              let _rhs156 = _1023_cf40;
              let _rhs157 = !_dafny.Seq.contains(_994_v3, _993_v2);
              let _rhs158 = ((true) ? (new _dafny.CodePoint('b'.codePointAt(0))) : (_993_v2));
              let _rhs159 = (_this).f24;
              let _lhs120 = globalState;
              let _lhs121 = globalState;
              let _lhs122 = globalState;
              _lhs120.f2 = _rhs155;
              _lhs121.f2 = _rhs156;
              r0 = _rhs157;
              _993_v2 = _rhs158;
              _lhs122.f1 = _rhs159;
              let _1029_v26;
              let _nw149 = Array((new BigNumber(19)).toNumber());
              _nw149[(_dafny.ZERO).toNumber()] = true;
              _nw149[(_dafny.ONE).toNumber()] = true;
              _nw149[(new BigNumber(2)).toNumber()] = p0;
              _nw149[(new BigNumber(3)).toNumber()] = _1024_cf39;
              _nw149[(new BigNumber(4)).toNumber()] = !(p0);
              _nw149[(new BigNumber(5)).toNumber()] = p0;
              _nw149[(new BigNumber(6)).toNumber()] = p0;
              _nw149[(new BigNumber(7)).toNumber()] = _1024_cf39;
              _nw149[(new BigNumber(8)).toNumber()] = false;
              _nw149[(new BigNumber(9)).toNumber()] = _1023_cf40;
              _nw149[(new BigNumber(10)).toNumber()] = !(p0);
              _nw149[(new BigNumber(11)).toNumber()] = _1023_cf40;
              _nw149[(new BigNumber(12)).toNumber()] = p0;
              _nw149[(new BigNumber(13)).toNumber()] = _1023_cf40;
              _nw149[(new BigNumber(14)).toNumber()] = _1023_cf40;
              _nw149[(new BigNumber(15)).toNumber()] = _1024_cf39;
              _nw149[(new BigNumber(16)).toNumber()] = !(!(true));
              _nw149[(new BigNumber(17)).toNumber()] = p0;
              _nw149[(new BigNumber(18)).toNumber()] = true;
              _1029_v26 = _nw149;
              let _1030_v27;
              _1030_v27 = _dafny.Seq.of(_1029_v26, _1029_v26);
              let _1031_v28;
              _1031_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1030_v27, _1030_v27),new BigNumber((_1027_v24).length));
              _1031_v28 = (_1031_v28).update(_1030_v27, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1027_v24).length))));
            } else if (_source11.is_DC26) {
              let _1032___mcc_h4 = (_source11).cf41;
              let _1033___mcc_h5 = (_source11).cf42;
              let _1034_cf42 = _1033___mcc_h5;
              let _1035_cf41 = _1032___mcc_h4;
              let _1036_v29;
              _1036_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_993_v2);
              _1036_v29 = (_1036_v29).update((_this).f24, _993_v2);
              _1034_cf42 = _994_v3;
              let _1037_v30;
              _1037_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
              let _1038_v31;
              _1038_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
              let _1039_v32;
              _1039_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _1040_v33;
              _1040_v33 = _dafny.Seq.of(new BigNumber((_994_v3).length), (_this).f24, (_this).f24, new BigNumber((_dafny.Set.fromElements(p0, p0)).length), new BigNumber((_1039_v32).length));
              let _1041_v34;
              _1041_v34 = _module.D8.create_DC25((_this).f24, p0, p0);
              let _1042_v35;
              _1042_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm14(_1037_v30, (((_1038_v31).contains(p0)) ? ((_1038_v31).get(p0)) : (new BigNumber((_1040_v33).length))), globalState)).cardinality()),(_1041_v34).dtor_cf39);
              let _1043_v36;
              let _nw150 = Array((new BigNumber(29)).toNumber()).fill(false);
              _1043_v36 = _nw150;
              let _1044_v37;
              _1044_v37 = _module.D1.create_DC5(p0, p0, !(p0));
              let _index124 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1043_v36).length));
              (_1043_v36)[_index124] = (_1044_v37).dtor_cf14;
              let _1045_v38;
              _1045_v38 = _dafny.Seq.of(_1043_v36, _1043_v36, _1043_v36);
              let _index125 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1043_v36).length));
              let _rhs160 = _1042_v35;
              let _rhs161 = !(true) || (p0);
              let _rhs162 = _1045_v38;
              let _lhs123 = _1043_v36;
              let _lhs124 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1043_v36).length));
              _1042_v35 = _rhs160;
              _lhs123[_lhs124] = _rhs161;
              _1045_v38 = _rhs162;
              let _1046_v39;
              _1046_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.D4.create_DC13(new BigNumber((_1040_v33).length), _1038_v31, _dafny.Seq.UnicodeFromString("de")),new BigNumber((_1040_v33).length));
              _1046_v39 = (_1046_v39).update(_module.__default.fm25(globalState), (_this).fm8(globalState));
            } else {
              let _1047___mcc_h6 = (_source11).cf36;
              let _1048_cf36 = _1047___mcc_h6;
              let _1049_v40;
              let _init25 = function (_1050_i2) {
                return _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
              };
              let _nw151 = Array((new BigNumber(25)).toNumber());
              for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw151.length); _i0_25++) {
                _nw151[_i0_25] = _init25(new BigNumber(_i0_25));
              }
              _1049_v40 = _nw151;
              let _1051_v41;
              let _nw152 = Array((new BigNumber(9)).toNumber());
              _nw152[(_dafny.ZERO).toNumber()] = _1049_v40;
              _nw152[(_dafny.ONE).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(2)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(3)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(4)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(5)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(6)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(7)).toNumber()] = _1049_v40;
              _nw152[(new BigNumber(8)).toNumber()] = _1049_v40;
              _1051_v41 = _nw152;
              let _index126 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1051_v41).length));
              (_1051_v41)[_index126] = ((!(p0)) ? (_1049_v40) : (_1049_v40));
              let _1052_v42;
              let _nw153 = Array((new BigNumber(27)).toNumber()).fill(false);
              _1052_v42 = _nw153;
              let _index127 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length));
              (_1052_v42)[_index127] = p0;
              let _1053_v43;
              _1053_v43 = _module.D1.create_DC5(p0, !(p0), !(p0));
              let _1054_v44;
              _1054_v44 = _dafny.Map.Empty.slice().updateUnsafe(_994_v3,_1053_v43);
              let _1055_v45;
              let _nw154 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1055_v45 = _nw154;
              let _index128 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1055_v45).length));
              (_1055_v45)[_index128] = _994_v3;
              let _1056_v46;
              _1056_v46 = _dafny.Seq.of(_1049_v40);
              let _index129 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1051_v41).length));
              let _index130 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length));
              let _index131 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1055_v45).length));
              let _rhs163 = (_1056_v46)[_module.__default.safeIndex((_this).f24, new BigNumber((_1056_v46).length))];
              let _rhs164 = _dafny.Seq.Concat(_994_v3, _994_v3);
              let _rhs165 = (new BigNumber(((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs((_this).f24))).cardinality())).isLessThan((_this).f24);
              let _rhs166 = (_module.__default.fm34(p0, globalState)).Merge(_1054_v44);
              let _rhs167 = _dafny.Seq.Concat(_994_v3, _994_v3);
              let _lhs125 = _1051_v41;
              let _lhs126 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1051_v41).length));
              let _lhs127 = globalState;
              let _lhs128 = _1052_v42;
              let _lhs129 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length));
              let _lhs130 = _1055_v45;
              let _lhs131 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1055_v45).length));
              _lhs125[_lhs126] = _rhs163;
              _lhs127.f17 = _rhs164;
              _lhs128[_lhs129] = _rhs165;
              _1054_v44 = _rhs166;
              _lhs130[_lhs131] = _rhs167;
              let _1057_v47;
              _1057_v47 = _dafny.Set.fromElements((_this).f24);
              let _1058_v48;
              _1058_v48 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1057_v47).length))).length));
              let _1059_v49;
              _1059_v49 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_1057_v47) : (_1057_v47)),new BigNumber(821));
              let _rhs168 = _1059_v49;
              let _rhs169 = (_this).f24;
              let _lhs132 = globalState;
              _1059_v49 = _rhs168;
              _lhs132.f9 = _rhs169;
              let _index132 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length));
              (_1052_v42)[_index132] = (_1052_v42)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length))];
              let _1060_v50;
              _1060_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v42)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length))],(_1052_v42)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length))]);
              let _1061_v51;
              _1061_v51 = _dafny.MultiSet.fromElements(_1060_v50);
              let _1062_v52;
              _1062_v52 = _dafny.MultiSet.fromElements(_module.__default.fm0(_module.__default.fm0(new BigNumber((_1061_v51).cardinality()), (_this).f24, globalState), new BigNumber((_module.__default.fm17(new BigNumber(157), new BigNumber(-613), (_this).f24, (_1052_v42)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length))], globalState)).length), globalState));
              let _1063_v53;
              _1063_v53 = _dafny.Map.Empty.slice().updateUnsafe(_993_v2,(_this).f24);
              let _1064_v54;
              _1064_v54 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v42)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1052_v42).length))],(_this).f25);
              let _1065_v55;
              let _nw155 = new _module.C1();
              _nw155.__ctor(_dafny.Seq.of(_1062_v52), (_this).f24, _module.__default.safeDivisionInt((_this).f24, (_this).f24), ((p0) ? (_dafny.Seq.of(_1063_v53)) : ((((_1064_v54).contains(p0)) ? ((_1064_v54).get(p0)) : ((_this).f25)))));
              _1065_v55 = _nw155;
            }
            let _1066_v56;
            _1066_v56 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f24)),p0);
            let _1067_v57;
            let _nw156 = new _module.C1();
            _nw156.__ctor(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_1066_v56).length))), (_this).f24, new BigNumber(939), _dafny.Seq.Concat((_this).f25, (_this).f25));
            _1067_v57 = _nw156;
            (globalState).f21 = _990_v0;
            let _1068_v58;
            _1068_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f24, (_dafny.ZERO).minus((_this).f24), _993_v2, globalState),!(_1067_v57.f33).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_990_v0,p0)).length)));
            _1068_v58 = (_1068_v58).update(p0, p0);
          }
        }
      }
      let _1077_v59;
      let _init26 = ((_1078_p0) => function (_1079_i3) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_1078_p0,_1078_p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1078_p0,_1078_p0));
      })(p0);
      let _nw157 = Array((new BigNumber(8)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw157.length); _i0_26++) {
        _nw157[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1077_v59 = _nw157;
      let _index133 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1077_v59).length));
      (_1077_v59)[_index133] = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1080_v60;
      _1080_v60 = _dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0)));
      let _1081_v61;
      let _nw158 = Array((new BigNumber(6)).toNumber()).fill(false);
      _1081_v61 = _nw158;
      let _index134 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
      (_1081_v61)[_index134] = !(p0) || (p0);
      let _1082_v62;
      _1082_v62 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1083_v63;
      _1083_v63 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p0)),!(!(p0)));
      let _1084_v64;
      _1084_v64 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
      let _index135 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1077_v59).length));
      let _index136 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
      let _rhs170 = _1083_v63;
      let _rhs171 = _1080_v60;
      let _rhs172 = ((_this).f24).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f24));
      let _rhs173 = new BigNumber(((_dafny.MultiSet.fromElements((_this).f24)).Intersect((_1084_v64).Intersect(_1084_v64))).cardinality());
      let _rhs174 = _1082_v62;
      let _lhs133 = _1077_v59;
      let _lhs134 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1077_v59).length));
      let _lhs135 = _1081_v61;
      let _lhs136 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
      let _lhs137 = globalState;
      _lhs133[_lhs134] = _rhs170;
      _1080_v60 = _rhs171;
      _lhs135[_lhs136] = _rhs172;
      _lhs137.f9 = _rhs173;
      _1082_v62 = _rhs174;
      let _1085_v65;
      _1085_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_module.__default.fm15((_this).f24, p0, (_this).f24, (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))], globalState));
      _1085_v65 = _1085_v65;
      if (!(((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]) === ((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]))) {
        let _1086_v66;
        _1086_v66 = _module.D3.create_DC11(((_this).f24).minus(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length)), p0, true);
        let _1087_v67;
        _1087_v67 = _dafny.Seq.UnicodeFromString("fidp");
        let _rhs175 = _1081_v61;
        let _rhs176 = _1086_v66;
        let _rhs177 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), ((_1088_v62) => function (_1089_i4) {
          return _1088_v62;
        })(_1082_v62)), _1087_v67);
        let _lhs138 = globalState;
        _1081_v61 = _rhs175;
        _1086_v66 = _rhs176;
        _lhs138.f17 = _rhs177;
        let _index137 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
        (_1081_v61)[_index137] = !(!((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]) || (!(p0))) || (p0);
        (globalState).f2 = (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))];
        let _1090_v68;
        _1090_v68 = _dafny.Seq.of((_this).f24);
        let _1091_v69;
        _1091_v69 = _module.D4.create_DC14(_module.D4.create_DC12(_1090_v68));
        let _1092_v70;
        _1092_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1091_v69,_1087_v67);
        let _1093_v71;
        _1093_v71 = _module.D7.create_DC20(_1092_v70);
        let _1094_v72;
        _1094_v72 = _module.D7.create_DC22(_1093_v71);
        _1094_v72 = _1094_v72;
        let _1095_v73;
        _1095_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
        let _1096_v74;
        _1096_v74 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),new BigNumber((_1087_v67).length));
        let _1097_v75;
        _1097_v75 = _module.D4.create_DC13((_this).f24, _1096_v74, _1087_v67);
        let _1098_v76;
        _1098_v76 = _dafny.Set.fromElements(new BigNumber(862));
        let _1099_v77;
        _1099_v77 = _dafny.MultiSet.fromElements(p0, p0, (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
        let _index138 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
        let _rhs178 = ((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]) && (!((new BigNumber((_1095_v73).length)).isLessThanOrEqualTo((_this).f24)));
        let _rhs179 = (_1097_v75).dtor_cf22;
        let _rhs180 = (_dafny.ZERO).minus(_module.__default.fm0(new BigNumber((_1098_v76).length), (((_1099_v77).contains((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))])) ? ((_1099_v77).get((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))])) : ((_this).f24)), globalState));
        let _rhs181 = (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))];
        let _rhs182 = ((_this).f24).minus(((_this).f24).plus((_this).f24));
        let _lhs139 = _1081_v61;
        let _lhs140 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
        let _lhs141 = globalState;
        let _lhs142 = globalState;
        let _lhs143 = globalState;
        _lhs139[_lhs140] = _rhs178;
        _lhs141.f14 = _rhs179;
        _lhs142.f9 = _rhs180;
        r0 = _rhs181;
        _lhs143.f1 = _rhs182;
      } else {
        (globalState).f2 = (((p0) ? (p0) : ((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]))) && (((p0) ? (p0) : (_module.__default.fm2(new BigNumber((_990_v0).length), (_this).f24, new _dafny.CodePoint('w'.codePointAt(0)), globalState))));
        let _1100_v78;
        _1100_v78 = _dafny.Seq.UnicodeFromString("omqcl");
        if (((!((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))])) ? (p0) : (_dafny.Seq.IsProperPrefixOf(_1100_v78, _1100_v78)))) {
          let _rhs183 = (new BigNumber((_dafny.Seq.Concat(_1100_v78, _dafny.Seq.UnicodeFromString("sqal"))).length)).plus((_this).f24);
          let _rhs184 = p0;
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          _lhs144.f1 = _rhs183;
          _lhs145.f2 = _rhs184;
          let _1101_v79;
          _1101_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(34),!(true));
          let _1102_v80;
          _1102_v80 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber((_1101_v79).length), (_this).f24));
          _1102_v80 = _1102_v80;
          let _1103_v81;
          let _nw159 = Array((new BigNumber(21)).toNumber()).fill([]);
          _1103_v81 = _nw159;
          let _1104_v82;
          _1104_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v63,(_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
          let _rhs185 = _1103_v81;
          let _rhs186 = (_dafny.ZERO).minus(((!(!((_this).f24).isEqualTo((_dafny.ZERO).minus((_this).f24)))) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), ((_1105_v61) => function (_1106_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe((_1105_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1105_v61).length))],(_this).f24);
          })(_1081_v61))).length)) : (new BigNumber(((_1104_v82).Merge(_1104_v82)).length))));
          let _lhs146 = globalState;
          _1103_v81 = _rhs185;
          _lhs146.f1 = _rhs186;
          let _1107_v83;
          _1107_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1081_v61);
          _1107_v83 = (_1107_v83).update(new BigNumber((_dafny.Seq.UnicodeFromString("nnv")).length), _1081_v61);
          let _index139 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
          (_1081_v61)[_index139] = (_990_v0)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f24, (_this).f24), new BigNumber((_990_v0).length))];
        } else {
          let _1108_v84;
          _1108_v84 = _dafny.Map.Empty.slice().updateUnsafe((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))],(_this).f24);
          _1108_v84 = _1108_v84;
          (globalState).f18 = (_this).f24;
          let _1109_v85;
          let _nw160 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1109_v85 = _nw160;
          let _index140 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1109_v85).length));
          (_1109_v85)[_index140] = (((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]) ? ((_dafny.ZERO).minus((_this).f24)) : ((_this).f24));
          let _index141 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_1109_v85).length));
          (_1109_v85)[_index141] = (_this).f24;
          (globalState).f2 = p0;
          (globalState).f14 = (_1109_v85)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_1109_v85).length))];
        }
        (globalState).f9 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f24, (_this).f24)), (_this).f24);
        if (p0) {
          let _1110_v86;
          _1110_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1082_v62,p0);
          let _index142 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length));
          (_1081_v61)[_index142] = !((((_1110_v86).contains(_1082_v62)) ? ((_1110_v86).get(_1082_v62)) : (p0)));
          (globalState).f9 = (_this).f24;
          let _1111_v87;
          let _init27 = ((_1112_v0) => function (_1113_i6) {
            return _1112_v0;
          })(_990_v0);
          let _nw161 = Array((new BigNumber(16)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw161.length); _i0_27++) {
            _nw161[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1111_v87 = _nw161;
          let _index143 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_1111_v87).length));
          (_1111_v87)[_index143] = _990_v0;
          let _index144 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_1111_v87).length));
          (_1111_v87)[_index144] = _dafny.Seq.update(_990_v0, _module.__default.safeIndex((_this).f24, new BigNumber((_990_v0).length)), false);
          let _1114_v88;
          _1114_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))], (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]),_1081_v61);
          _1114_v88 = _1114_v88;
          (globalState).f9 = (_this).f24;
        } else {
          let _1115_v89;
          _1115_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1082_v62);
          _1082_v62 = (((_1115_v89).contains(new BigNumber((_dafny.Seq.Concat(_990_v0, _990_v0)).length))) ? ((_1115_v89).get(new BigNumber((_dafny.Seq.Concat(_990_v0, _990_v0)).length))) : (_1082_v62));
          r1 = p0;
          let _1116_v90;
          _1116_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f24, (_this).fm8(globalState)),(((_1083_v63).contains(p0)) ? ((_1083_v63).get(p0)) : (p0)));
          _1116_v90 = (_1116_v90).update((_this).f24, (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
          let _1117_v91;
          let _nw162 = new _module.C2();
          _nw162.__ctor(p0, _module.__default.safeModuloInt((_this).f24, (_this).f24), (_this).f25);
          _1117_v91 = _nw162;
          let _1118_v92;
          _1118_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1082_v62,!(true));
          r0 = (((_1118_v92).contains(_1082_v62)) ? ((_1118_v92).get(_1082_v62)) : ((new BigNumber(123)).isLessThan((_this).f24)));
        }
        let _1119_v93;
        _1119_v93 = _dafny.Set.fromElements((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))], (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
        let _1120_v94;
        _1120_v94 = _module.D0.create_DC1(new BigNumber((_1119_v93).length), (_this).f24, (_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
        let _1121_v95;
        let _nw163 = new _module.C0();
        _nw163.__ctor((_1120_v94).dtor_cf3);
        _1121_v95 = _nw163;
      }
      _1081_v61 = _1081_v61;
      let _1122_i7;
      _1122_i7 = _dafny.ZERO;
      L3: {
        while (!((_this).f24).isEqualTo((_this).f24)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1122_i7)) {
              break L3;
            }
            _1122_i7 = (_1122_i7).plus(_dafny.ONE);
            (globalState).f1 = (((_this).f24).plus((_this).f24)).plus((_this).f24);
            let _1123_v96;
            _1123_v96 = _dafny.Set.fromElements(true);
            _1123_v96 = _1123_v96;
            let _1124_v97;
            _1124_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
            _1124_v97 = _1124_v97;
            let _1125_v98;
            _1125_v98 = _dafny.MultiSet.fromElements((_1081_v61)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_1081_v61).length))]);
            r1 = (_dafny.MultiSet.FromArray(_990_v0)).IsProperSubsetOf(_1125_v98);
          }
        }
      }
      let _1126_v99;
      _1126_v99 = _dafny.Seq.UnicodeFromString("lq");
      let _1127_v100;
      _1127_v100 = _dafny.Set.fromElements(new BigNumber((_1126_v99).length));
      r0 = (((_module.D6.create_DC16(_1127_v100)).dtor_cf27).Union(_1127_v100)).IsDisjointFrom((_dafny.Set.fromElements(new BigNumber(17))).Difference(_dafny.Set.fromElements((_this).f24)));
      r1 = !((p0) === (p0));
      r2 = (_this).f24;
      return [r0, r1, r2];
    }
    m6(globalState) {
      let _this = this;
      let _1128_v0;
      _1128_v0 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f24).plus((_this).f24),true);
      let _1129_v1;
      _1129_v1 = false;
      _1128_v0 = (_1128_v0).update((_this).f24, _1129_v1);
      (globalState).f2 = !((((_this).f24).minus((_this).f24)).isEqualTo((_this).f24));
      let _1130_v2;
      _1130_v2 = _module.D1.create_DC5(false, _1129_v1, _1129_v1);
      let _1131_v3;
      _1131_v3 = _module.D6.create_DC17(_1129_v1);
      let _1132_v4;
      _1132_v4 = _module.D6.create_DC19(_1131_v3);
      let _1133_v5;
      _1133_v5 = _dafny.Seq.of(_1129_v1, _1129_v1);
      let _1134_v6;
      let _nw164 = Array((new BigNumber(12)).toNumber());
      _nw164[(_dafny.ZERO).toNumber()] = _1129_v1;
      _nw164[(_dafny.ONE).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(2)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(3)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(4)).toNumber()] = !(_1129_v1);
      _nw164[(new BigNumber(5)).toNumber()] = (_1133_v5)[_module.__default.safeIndex((_this).f24, new BigNumber((_1133_v5).length))];
      _nw164[(new BigNumber(6)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(7)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(8)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(9)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(10)).toNumber()] = _1129_v1;
      _nw164[(new BigNumber(11)).toNumber()] = _1129_v1;
      _1134_v6 = _nw164;
      let _1135_v7;
      _1135_v7 = _module.D7.create_DC21(_1130_v2, _1132_v4, _1134_v6);
      let _1136_v8;
      _1136_v8 = new _dafny.CodePoint('k'.codePointAt(0));
      let _pat_let_tv7 = _1136_v8;
      let _pat_let_tv8 = globalState;
      let _source13 = function (_pat_let12_0) {
        return function (_1137_dt__update__tmp_h0) {
          return function (_pat_let13_0) {
            return function (_1138_dt__update_hcf33_h0) {
              return _module.D7.create_DC21((_1137_dt__update__tmp_h0).dtor_cf32, _1138_dt__update_hcf33_h0, (_1137_dt__update__tmp_h0).dtor_cf34);
            }(_pat_let13_0);
          }(_module.D6.create_DC19(_module.D6.create_DC18(_module.__default.fm2((_this).f24, (_this).f24, _pat_let_tv7, _pat_let_tv8))));
        }(_pat_let12_0);
      }(_1135_v7);
      if (_source13.is_DC21) {
        let _1139___mcc_h0 = (_source13).cf32;
        let _1140___mcc_h1 = (_source13).cf33;
        let _1141___mcc_h2 = (_source13).cf34;
        let _1142_cf34 = _1141___mcc_h2;
        let _1143_cf33 = _1140___mcc_h1;
        let _1144_cf32 = _1139___mcc_h0;
        (globalState).f2 = (_1129_v1) && (_1129_v1);
        (globalState).f2 = _1129_v1;
        let _1145_v9;
        _1145_v9 = _dafny.Seq.of(_1136_v8, _1136_v8);
        let _1146_v10;
        _1146_v10 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1145_v9).length)), (_this).f24, (_this).f24);
        let _1147_v11;
        _1147_v11 = _module.D6.create_DC18(_1129_v1);
        let _1148_v12;
        _1148_v12 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), ((_1149_v8) => function (_1150_i0) {
          return _1149_v8;
        })(_1136_v8)));
        let _1151_v13;
        _1151_v13 = _dafny.MultiSet.fromElements(_1129_v1);
        let _rhs187 = new BigNumber((_dafny.Seq.update(_1148_v12, _module.__default.safeIndex((_this).f24, new BigNumber((_1148_v12).length)), _1145_v9)).length);
        let _rhs188 = _1146_v10;
        let _rhs189 = _1133_v5;
        let _rhs190 = function (_pat_let14_0) {
          return function (_1152_dt__update__tmp_h1) {
            return function (_pat_let15_0) {
              return function (_1153_dt__update_hcf29_h0) {
                return _module.D6.create_DC18(_1153_dt__update_hcf29_h0);
              }(_pat_let15_0);
            }(true);
          }(_pat_let14_0);
        }(_1147_v11);
        let _rhs191 = _module.__default.fm0(((_this).f24).minus(new BigNumber(((_1151_v13).update(_1129_v1, _module.__default.abs((_this).f24))).cardinality())), new BigNumber(600), globalState);
        let _lhs147 = globalState;
        let _lhs148 = globalState;
        let _lhs149 = globalState;
        _lhs147.f18 = _rhs187;
        _1146_v10 = _rhs188;
        _lhs148.f20 = _rhs189;
        _1147_v11 = _rhs190;
        _lhs149.f18 = _rhs191;
        let _1154_v14;
        let _init28 = function (_1155_i1) {
          return (_1155_i1).plus((_this).f24);
        };
        let _nw165 = Array((new BigNumber(17)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw165.length); _i0_28++) {
          _nw165[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1154_v14 = _nw165;
        let _1156_v15;
        _1156_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v14,_1129_v1);
        let _1157_v16;
        _1157_v16 = _dafny.Seq.of(_1146_v10);
        let _1158_v17;
        _1158_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1145_v9,_1157_v16);
        let _1159_v18;
        _1159_v18 = _dafny.Seq.of(new BigNumber(57));
        let _rhs192 = (((_this).f24).minus((_this).f24)).plus((_this).fm8(globalState));
        let _rhs193 = _module.__default.fm3(globalState);
        let _rhs194 = _1156_v15;
        let _rhs195 = (((_1158_v17).contains(_1145_v9)) ? ((_1158_v17).get(_1145_v9)) : (_module.__default.fm35(!(_1129_v1), _1129_v1, globalState)));
        let _rhs196 = (_dafny.MultiSet.FromArray(_1159_v18)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_1128_v0).length), (_this).f24, (_this).f24, new BigNumber(709)));
        let _lhs150 = globalState;
        let _lhs151 = globalState;
        _lhs150.f1 = _rhs192;
        _1136_v8 = _rhs193;
        _1156_v15 = _rhs194;
        _1157_v16 = _rhs195;
        _lhs151.f2 = _rhs196;
      } else if (_source13.is_DC20) {
        let _1160___mcc_h3 = (_source13).cf31;
        let _1161_cf31 = _1160___mcc_h3;
        let _1162_v19;
        let _nw166 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1162_v19 = _nw166;
        let _1163_v20;
        _1163_v20 = _dafny.MultiSet.fromElements(_1129_v1, _1129_v1, _1129_v1);
        let _index145 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1162_v19).length));
        (_1162_v19)[_index145] = ((_this).f24).minus(new BigNumber((_1163_v20).cardinality()));
        let _1164_v21;
        _1164_v21 = _dafny.Seq.UnicodeFromString("gsj");
        let _index146 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1162_v19).length));
        let _rhs197 = _1164_v21;
        let _rhs198 = (_this).f24;
        let _lhs152 = globalState;
        let _lhs153 = _1162_v19;
        let _lhs154 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1162_v19).length));
        _lhs152.f17 = _rhs197;
        _lhs153[_lhs154] = _rhs198;
        _1129_v1 = _1129_v1;
        (globalState).f2 = _1129_v1;
        (globalState).f9 = ((_1162_v19)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_1162_v19).length))]).plus((_1162_v19)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_1162_v19).length))]);
      } else {
        let _1165___mcc_h4 = (_source13).cf35;
        let _1166_cf35 = _1165___mcc_h4;
        let _1167_v22;
        _1167_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1136_v8,(_this).f24);
        let _1168_v23;
        let _nw167 = new _module.C2();
        _nw167.__ctor(_1129_v1, new BigNumber(-718), _dafny.Seq.update(_dafny.Seq.of(_1167_v22), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.of(_1167_v22)).length)), _1167_v22));
        _1168_v23 = _nw167;
        let _1169_v24;
        _1169_v24 = _dafny.Seq.of((_this).f24);
        let _1170_v25;
        _1170_v25 = _dafny.Seq.UnicodeFromString("qe");
        let _1171_v27;
        _1171_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.of((_this).f24)).length));
        let _1172_v28;
        _1172_v28 = _dafny.MultiSet.fromElements((_this).f24);
        let _1173_v29;
        let _nw168 = Array((new BigNumber(13)).toNumber());
        _nw168[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw168[(_dafny.ONE).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(2)).toNumber()] = (_1169_v24)[_module.__default.safeIndex(new BigNumber((_1170_v25).length), new BigNumber((_1169_v24).length))];
        _nw168[(new BigNumber(3)).toNumber()] = new BigNumber((function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_1171_v27).Keys.Elements) {
            let _1174_v26 = _compr_30;
            if ((_1171_v27).contains(_1174_v26)) {
              _coll30.push([(_1174_v26).multipliedBy(new BigNumber((_1133_v5).length)),new BigNumber(-366)]);
            }
          }
          return _coll30;
        }()).length);
        _nw168[(new BigNumber(4)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(5)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(8)).toNumber()] = new BigNumber(495);
        _nw168[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(10)).toNumber()] = new BigNumber(488);
        _nw168[(new BigNumber(11)).toNumber()] = (_this).f24;
        _nw168[(new BigNumber(12)).toNumber()] = new BigNumber((_1172_v28).cardinality());
        _1173_v29 = _nw168;
        let _1175_v30;
        _1175_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("i")).length),_1173_v29);
        let _1176_v31;
        _1176_v31 = _module.D0.create_DC0(_1173_v29);
        let _1177_v32;
        _1177_v32 = _dafny.Set.fromElements(_1173_v29, (((_1175_v30).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1176_v31,_1173_v29)).length))) ? ((_1175_v30).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1176_v31,_1173_v29)).length))) : (_1173_v29)), _1173_v29, _1173_v29, _1173_v29);
        _1177_v32 = ((_1177_v32).Difference(_1177_v32)).Union(_1177_v32);
        (globalState).f21 = _1133_v5;
        let _1178_v33;
        _1178_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1168_v23.f31,((_dafny.ZERO).minus((_this).f24)).isLessThan((_dafny.ZERO).minus((_this).f24)));
        _1178_v33 = (_1178_v33).update(!(_1168_v23.f31), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("pfodpobt"), _1170_v25));
      }
      let _1179_v34;
      _1179_v34 = _dafny.Seq.UnicodeFromString("u");
      let _1180_v35;
      _1180_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1129_v1,_1179_v34);
      let _1181_v36;
      _1181_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_1129_v1) ? ((_this).f24) : ((_this).f24)),(_1180_v35).Merge(_1180_v35));
      _1181_v36 = (_1181_v36).update((_this).f24, _dafny.Map.Empty.slice().updateUnsafe(_1129_v1,_1179_v34));
      (globalState).f17 = _1179_v34;
      let _1182_v38;
      _1182_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1129_v1,new BigNumber((_dafny.Seq.of(_1129_v1, _1129_v1, false, !(_1129_v1))).length));
      let _1183_v39;
      _1183_v39 = _module.D4.create_DC14(_module.D4.create_DC13(new BigNumber((_1179_v34).length), _1182_v38, _1179_v34));
      let _1184_v40;
      _1184_v40 = _dafny.Set.fromElements(_1183_v39);
      let _1185_v41;
      _1185_v41 = _module.D7.create_DC20(function () {
  let _coll31 = new _dafny.Map();
  for (const _compr_31 of (_1184_v40).Elements) {
    let _1186_v37 = _compr_31;
    if ((_1184_v40).contains(_1186_v37)) {
      _coll31.push([_1186_v37,_1179_v34]);
    }
  }
  return _coll31;
}());
      let _1187_i2;
      _1187_i2 = _dafny.ZERO;
      L4: {
        let _pat_let_tv9 = _1136_v8;
        let _pat_let_tv10 = _1136_v8;
        let _pat_let_tv11 = _1136_v8;
        let _pat_let_tv12 = _1136_v8;
        let _pat_let_tv13 = _1136_v8;
        let _pat_let_tv14 = _1129_v1;
        while (function (_source14) {
          if (_source14.is_DC21) {
            let _1221___mcc_h5 = (_source14).cf32;
            let _1222___mcc_h6 = (_source14).cf33;
            let _1223___mcc_h7 = (_source14).cf34;
            let _1224_cf34 = _1223___mcc_h7;
            let _1225_cf33 = _1222___mcc_h6;
            let _1226_cf32 = _1221___mcc_h5;
            return (function () {
              let _coll32 = new _dafny.Set();
              for (const _compr_32 of (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv9,_pat_let_tv10)).Keys.Elements) {
                let _1227_v42 = _compr_32;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv11,_pat_let_tv12)).contains(_1227_v42)) {
                  _coll32.add(_1227_v42);
                }
              }
              return _coll32;
            }()).IsProperSubsetOf(_dafny.Set.fromElements(_pat_let_tv13));
          } else if (_source14.is_DC20) {
            let _1228___mcc_h8 = (_source14).cf31;
            let _1229_cf31 = _1228___mcc_h8;
            return !(!(!(true)));
          } else {
            let _1230___mcc_h9 = (_source14).cf35;
            let _1231_cf35 = _1230___mcc_h9;
            return (_pat_let_tv14) === (true);
          }
        }(_1185_v41)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1187_i2)) {
              break L4;
            }
            _1187_i2 = (_1187_i2).plus(_dafny.ONE);
            let _1188_v43;
            _1188_v43 = _dafny.MultiSet.fromElements(new BigNumber((_1179_v34).length));
            let _1189_v44;
            let _1190_v45;
            let _1191_v46;
            let _out28;
            let _out29;
            let _out30;
            let _outcollector5 = (_this).m5(_module.__default.fm2(new BigNumber((_1188_v43).cardinality()), (_this).f24, _1136_v8, globalState), globalState);
            _out28 = _outcollector5[0];
            _out29 = _outcollector5[1];
            _out30 = _outcollector5[2];
            _1189_v44 = _out28;
            _1190_v45 = _out29;
            _1191_v46 = _out30;
            let _1192_v47;
            _1192_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1190_v45,_1129_v1);
            let _1193_v48;
            let _nw169 = new _module.C2();
            _nw169.__ctor(_1189_v44, new BigNumber((_1192_v47).length), (_this).f25);
            _1193_v48 = _nw169;
            let _1194_v49;
            _1194_v49 = _dafny.Seq.of(_1193_v48);
            let _1195_v50;
            _1195_v50 = _dafny.Seq.of((_this).f24, (_dafny.ZERO).minus(new BigNumber((_1194_v49).length)));
            _1128_v0 = (_1128_v0).update((new BigNumber((_1195_v50).length)).minus(_1191_v46), _1193_v48.f31);
            if (_1190_v45) {
              (globalState).f18 = (_1191_v46).plus((_this).f24);
              let _1196_v51;
              _1196_v51 = _dafny.Seq.of(_1192_v47);
              let _1197_v52;
              _1197_v52 = _dafny.Seq.of(_1196_v51, _1196_v51);
              let _1198_v53;
              _1198_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1129_v1,_1196_v51);
              let _1199_v54;
              let _nw170 = Array((new BigNumber(29)).toNumber());
              _nw170[(_dafny.ZERO).toNumber()] = ((_1190_v45) ? (_1196_v51) : (_1196_v51));
              _nw170[(_dafny.ONE).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(2)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(3)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(4)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(5)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(6)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm36(_1133_v5, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_1200_v47) => function (_1201_i3) {
                return _1200_v47;
              })(_1192_v47)));
              _nw170[(new BigNumber(8)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(9)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(10)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(11)).toNumber()] = (_1197_v52)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f24), new BigNumber((_1197_v52).length))];
              _nw170[(new BigNumber(12)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(13)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _dafny.Seq.of(_1192_v47));
              _nw170[(new BigNumber(15)).toNumber()] = (((_1198_v53).contains(_1193_v48.f31)) ? ((_1198_v53).get(_1193_v48.f31)) : (_dafny.Seq.of(_1192_v47)));
              _nw170[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), ((_1202_v47, _1203_v1, _1204_v48) => function (_1205_i4) {
                return (_1202_v47).update(_1203_v1, _1204_v48.f31);
              })(_1192_v47, _1129_v1, _1193_v48));
              _nw170[(new BigNumber(17)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _1196_v51);
              _nw170[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _module.__default.fm36(_dafny.Seq.of(_1129_v1, _1129_v1), globalState));
              _nw170[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _dafny.Seq.update(_dafny.Seq.update(_1196_v51, _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1196_v51).length)), _1192_v47), _module.__default.safeIndex(_1191_v46, new BigNumber((_dafny.Seq.update(_1196_v51, _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1196_v51).length)), _1192_v47)).length)), _1192_v47));
              _nw170[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), ((_1206_v47) => function (_1207_i5) {
                return _1206_v47;
              })(_1192_v47));
              _nw170[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _1196_v51);
              _nw170[(new BigNumber(23)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(24)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(25)).toNumber()] = _1196_v51;
              _nw170[(new BigNumber(26)).toNumber()] = _dafny.Seq.of(_1192_v47, _1192_v47, _1192_v47);
              _nw170[(new BigNumber(27)).toNumber()] = _dafny.Seq.of(_1192_v47);
              _nw170[(new BigNumber(28)).toNumber()] = _dafny.Seq.Concat(_1196_v51, _dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), ((_1208_v47, _1209_v48) => function (_1210_i6) {
                return (_1208_v47).update(_1209_v48.f31, _1209_v48.f31);
              })(_1192_v47, _1193_v48)));
              _1199_v54 = _nw170;
              let _index147 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1199_v54).length));
              (_1199_v54)[_index147] = _1196_v51;
              let _index148 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1199_v54).length));
              (_1199_v54)[_index148] = _module.__default.fm36(_module.__default.fm23(globalState), globalState);
              (globalState).f14 = (_this).f24;
              let _1211_v55;
              _1211_v55 = _module.D1.create_DC3(_1179_v34);
              let _1212_v56;
              _1212_v56 = _module.D8.create_DC26(_1211_v55, _1179_v34);
              let _1213_v57;
              _1213_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(globalState),_1212_v56);
              _1213_v57 = (_1213_v57).update(_1191_v46, _module.__default.fm37(globalState));
              let _1214_v58;
              let _nw171 = new _module.C2();
              _nw171.__ctor((_1191_v46).isLessThan(_1191_v46), (_this).fm8(globalState), (_this).f25);
              _1214_v58 = _nw171;
            } else {
              let _1215_v59;
              _1215_v59 = _dafny.Seq.of(_1188_v43);
              let _1216_v60;
              let _nw172 = new _module.C1();
              _nw172.__ctor(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(694)), ((_1217_v43) => function (_1218_i7) {
                return _1217_v43;
              })(_1188_v43)), _1215_v59), (_this).f24, ((_this).f24).multipliedBy((_this).f24), (_this).f25);
              _1216_v60 = _nw172;
              let _1219_v61;
              let _nw173 = new _module.C1();
              _nw173.__ctor(_1215_v59, _1216_v60.f33, _1216_v60.f33, (_this).f25);
              _1219_v61 = _nw173;
              _1129_v1 = (((_1128_v0).contains(_1191_v46)) ? ((_1128_v0).get(_1191_v46)) : (_1193_v48.f31));
              (_1216_v60).f33 = _module.__default.safeModuloInt(_1219_v61.f33, (_this).f24);
              let _1220_v62;
              _1220_v62 = _dafny.MultiSet.fromElements(_1193_v48.f31, _1129_v1, (_1216_v60.f33).isLessThan((_this).f24), _1193_v48.f31, _1193_v48.f31);
              _1220_v62 = _1220_v62;
            }
            (globalState).f18 = (_1195_v50)[_module.__default.safeIndex((_this).f24, new BigNumber((_1195_v50).length))];
          }
        }
      }
      return;
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
      this._f28 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f28, f24, f25) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_this).f24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),(_this).f24))).length), (_this).f24, true);
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let _1232_v1;
      _1232_v1 = _dafny.MultiSet.fromElements((_this).f24);
      let _1233_v2;
      _1233_v2 = false;
      let _1234_v3;
      _1234_v3 = _dafny.Seq.of(_1233_v2, _1233_v2);
      let _1235_v4;
      _1235_v4 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f24);
      let _1236_v5;
      _1236_v5 = _dafny.Seq.of(_1235_v4);
      let _1237_v6;
      _1237_v6 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1238_v7;
      _1238_v7 = _module.D0.create_DC2(_1233_v2, new BigNumber(((_1236_v5)[_module.__default.safeIndex((_this).f24, new BigNumber((_1236_v5).length))]).length), _1237_v6);
      let _1239_v8;
      _1239_v8 = _dafny.Seq.UnicodeFromString("dnox");
      let _1240_v9;
      _1240_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1239_v8).length));
      let _1241_v10;
      _1241_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1240_v9);
      let _1242_v12;
      let _nw174 = Array((new BigNumber(27)).toNumber()).fill(false);
      _1242_v12 = _nw174;
      let _1243_v13;
      _1243_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v2,_1242_v12);
      let _1244_v14;
      let _nw175 = Array((new BigNumber(13)).toNumber());
      _nw175[(_dafny.ZERO).toNumber()] = function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_1232_v1).Elements) {
          let _1245_v0 = _compr_33;
          if ((_1232_v1).contains(_1245_v0)) {
            _coll33.push([_module.__default.safeModuloInt(_1245_v0, (_this).f24),new BigNumber(98)]);
          }
        }
        return _coll33;
      }();
      _nw175[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1233_v2,_1233_v2)).length),(_this).f24);
      _nw175[(new BigNumber(2)).toNumber()] = (((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_dafny.ZERO).minus((_this).f24))).update((_this).f24, (_this).f24)).update((_this).f24, new BigNumber(698))).update(new BigNumber((_1234_v3).length), (_1238_v7).dtor_cf5);
      _nw175[(new BigNumber(3)).toNumber()] = _module.__default.fm6((_this).f24, (_this).f24, _module.__default.fm0((_this).f24, (_this).f24, globalState), (_this).f24, globalState);
      _nw175[(new BigNumber(4)).toNumber()] = _1240_v9;
      _nw175[(new BigNumber(5)).toNumber()] = _1240_v9;
      _nw175[(new BigNumber(6)).toNumber()] = _1240_v9;
      _nw175[(new BigNumber(7)).toNumber()] = (((_1241_v10).contains((_this).f24)) ? ((_1241_v10).get((_this).f24)) : (_1240_v9));
      _nw175[(new BigNumber(8)).toNumber()] = _1240_v9;
      _nw175[(new BigNumber(9)).toNumber()] = _1240_v9;
      _nw175[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)).Merge(function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of (_dafny.Seq.update(_module.__default.fm7(globalState), _module.__default.safeIndex((_this).f24, new BigNumber((_module.__default.fm7(globalState)).length)), (_this).f24)).Elements) {
          let _1246_v11 = _compr_34;
          if (_dafny.Seq.contains(_dafny.Seq.update(_module.__default.fm7(globalState), _module.__default.safeIndex((_this).f24, new BigNumber((_module.__default.fm7(globalState)).length)), (_this).f24), _1246_v11)) {
            _coll34.push([(_1246_v11).minus((_this).f24),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1239_v8).length),_1233_v2))).cardinality())]);
          }
        }
        return _coll34;
      }());
      _nw175[(new BigNumber(11)).toNumber()] = (_1240_v9).Merge(_1240_v9);
      _nw175[(new BigNumber(12)).toNumber()] = (_1240_v9).update(new BigNumber((_1243_v13).length), (_dafny.ZERO).minus((_this).f24));
      _1244_v14 = _nw175;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1244_v14).length))) {
        let _1247_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1247_i0)) && ((_1247_i0).isLessThan(new BigNumber((_1244_v14).length))))) {
          (_1244_v14)[(_1247_i0)] = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
        }
      }
      let _1248_v15;
      _1248_v15 = _dafny.Set.fromElements(new BigNumber(575));
      let _hi7 = ((_this).f24).plus(new BigNumber((_1248_v15).length));
      for (let _1249_i1 = (_this).f24; _1249_i1.isLessThan(_hi7); _1249_i1 = _1249_i1.plus(_dafny.ONE)) {
        let _1250_v16;
        let _init29 = ((_1251_v8) => function (_1252_i2) {
          return _1251_v8;
        })(_1239_v8);
        let _nw176 = Array((new BigNumber(9)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw176.length); _i0_29++) {
          _nw176[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1250_v16 = _nw176;
        _1250_v16 = _1250_v16;
        (globalState).f14 = (new BigNumber((_dafny.Seq.update(_1239_v8, _module.__default.safeIndex((_this).f24, new BigNumber((_1239_v8).length)), _1237_v6)).length)).minus(_module.__default.safeModuloInt((_this).f24, (_this).f24));
        let _1253_v17;
        _1253_v17 = _dafny.Seq.of(new BigNumber((_1239_v8).length), _1249_i1, _1249_i1);
        let _1254_v18;
        _1254_v18 = _module.D1.create_DC4(((_this).f24).isLessThanOrEqualTo(_1249_i1), _1233_v2, _1237_v6, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm2(_1249_i1, new BigNumber((_1239_v8).length), _1237_v6, globalState)), _1234_v3)).length), (_1253_v17)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1239_v8, _module.__default.safeIndex(new BigNumber((_1253_v17).length), new BigNumber((_1239_v8).length)), _1237_v6)).length), new BigNumber((_1253_v17).length))]);
        _1254_v18 = _1254_v18;
        let _index149 = _module.__default.safeIndex(new BigNumber(864), new BigNumber(((_this).f28).length));
        ((_this).f28)[_index149] = (new BigNumber(350)).minus(_1249_i1);
        let _index150 = _module.__default.safeIndex(new BigNumber(864), new BigNumber(((_this).f28).length));
        let _rhs199 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1239_v8, _1239_v8), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f24), new BigNumber((_dafny.Seq.Concat(_1239_v8, _1239_v8)).length)), _1237_v6)).length));
        let _rhs200 = _1233_v2;
        let _lhs155 = (_this).f28;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(864), new BigNumber(((_this).f28).length));
        let _lhs157 = globalState;
        _lhs155[_lhs156] = _rhs199;
        _lhs157.f2 = _rhs200;
      }
      let _1255_v19;
      _1255_v19 = _dafny.MultiSet.fromElements(_1237_v6, _1237_v6);
      let _1256_v21;
      _1256_v21 = _dafny.Set.fromElements(_1237_v6, _1237_v6, _1237_v6);
      let _hi8 = (_this).f24;
      for (let _1257_i3 = new BigNumber(((function () {
        let _coll35 = new _dafny.Set();
        for (const _compr_35 of (_1255_v19).Elements) {
          let _1280_v20 = _compr_35;
          if ((_1255_v19).contains(_1280_v20)) {
            _coll35.add(_1280_v20);
          }
        }
        return _coll35;
      }()).Intersect(_1256_v21)).length); _1257_i3.isLessThan(_hi8); _1257_i3 = _1257_i3.plus(_dafny.ONE)) {
        if (_1233_v2) {
          let _index151 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1242_v12).length));
          (_1242_v12)[_index151] = (function (_pat_let16_0) {
            return function (_1258_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_1259_dt__update_hcf5_h0) {
                  return _module.D0.create_DC2((_1258_dt__update__tmp_h0).dtor_cf4, _1259_dt__update_hcf5_h0, (_1258_dt__update__tmp_h0).dtor_cf6);
                }(_pat_let17_0);
              }(_1257_i3);
            }(_pat_let16_0);
          }(_1238_v7)).dtor_cf4;
          let _index152 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1242_v12).length));
          (_1242_v12)[_index152] = _module.__default.fm2((_dafny.ZERO).minus((_1257_i3).plus(new BigNumber((_1240_v9).length))), _1257_i3, _1237_v6, globalState);
          let _1260_v22;
          _1260_v22 = _dafny.Seq.of(new BigNumber((_1239_v8).length));
          let _1261_v23;
          _1261_v23 = _dafny.Seq.of(_1232_v1, _dafny.MultiSet.fromElements(_1257_i3));
          let _1262_v24;
          _1262_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,_1257_i3);
          let _1263_v25;
          let _nw177 = new _module.C1();
          _nw177.__ctor(_1261_v23, (_this).f24, (_this).f24, _dafny.Seq.of(_1262_v24, _1262_v24, _1262_v24));
          _1263_v25 = _nw177;
          let _1264_v26;
          _1264_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1263_v25,_1233_v2);
          (globalState).f1 = _module.__default.safeDivisionInt((new BigNumber((_1260_v22).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1260_v22).length))), new BigNumber((_1264_v26).length));
          (globalState).f18 = _1257_i3;
          let _1265_v27;
          _1265_v27 = _dafny.MultiSet.fromElements(_1240_v9);
          (globalState).f2 = (_1265_v27).IsProperSubsetOf(_1265_v27);
          let _rhs201 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1257_i3));
          let _rhs202 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1239_v8, _dafny.Seq.UnicodeFromString("d")), _1239_v8);
          let _rhs203 = _1233_v2;
          let _lhs158 = globalState;
          _1248_v15 = _rhs201;
          _lhs158.f17 = _rhs202;
          r0 = _rhs203;
        } else {
          let _1266_v28;
          _1266_v28 = _dafny.Seq.of(_1257_i3);
          let _1267_v29;
          _1267_v29 = _module.D4.create_DC12(_1266_v28);
          let _1268_v30;
          _1268_v30 = _module.D4.create_DC14(_1267_v29);
          let _1269_v31;
          _1269_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v30,_1239_v8);
          let _1270_v32;
          _1270_v32 = _module.D7.create_DC20(_1269_v31);
          let _1271_v33;
          _1271_v33 = _dafny.MultiSet.fromElements(_1270_v32, _module.D7.create_DC20(_1269_v31), _1270_v32);
          _1271_v33 = _1271_v33;
          (globalState).f14 = ((_this).f24).multipliedBy(_1257_i3);
          let _1272_v34;
          _1272_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v2,_1233_v2);
          _1272_v34 = (_1272_v34).update(!(_1233_v2), _1233_v2);
          (globalState).f2 = _1233_v2;
          let _1273_v35;
          _1273_v35 = _dafny.MultiSet.fromElements(_1233_v2);
          let _1274_v36;
          _1274_v36 = _module.D3.create_DC10();
          let _1275_v37;
          _1275_v37 = _dafny.MultiSet.fromElements(_1274_v36);
          _1273_v35 = _dafny.MultiSet.fromElements((_1275_v37).IsProperSubsetOf(_1275_v37));
        }
        let _index153 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length));
        (_1242_v12)[_index153] = false;
        let _1276_v38;
        _1276_v38 = _dafny.Seq.of(_1239_v8);
        let _index154 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length));
        let _rhs204 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xho"), _1239_v8), _dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), ((_1277_v6) => function (_1278_i4) {
          return _1277_v6;
        })(_1237_v6)));
        let _rhs205 = _1255_v19;
        let _rhs206 = _dafny.Seq.of(_1239_v8, _1239_v8);
        let _rhs207 = _1233_v2;
        let _lhs159 = _1242_v12;
        let _lhs160 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length));
        let _lhs161 = globalState;
        _lhs159[_lhs160] = _rhs204;
        _1255_v19 = _rhs205;
        _1276_v38 = _rhs206;
        _lhs161.f2 = _rhs207;
        _1235_v4 = (_1235_v4).update(!((_1242_v12)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length))]), _1257_i3);
        let _1279_v39;
        _1279_v39 = _dafny.Seq.of((_this).f28, (_this).f28, (_this).f28);
        let _index155 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length));
        let _rhs208 = _1239_v8;
        let _rhs209 = _dafny.Seq.update(_1279_v39, _module.__default.safeIndex((_this).f24, new BigNumber((_1279_v39).length)), (_this).f28);
        let _rhs210 = _1248_v15;
        let _rhs211 = _1233_v2;
        let _rhs212 = !((_1242_v12)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length))]) || (_1233_v2);
        let _lhs162 = globalState;
        let _lhs163 = _1242_v12;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1242_v12).length));
        let _lhs165 = globalState;
        _lhs162.f17 = _rhs208;
        _1279_v39 = _rhs209;
        _1248_v15 = _rhs210;
        _lhs163[_lhs164] = _rhs211;
        _lhs165.f2 = _rhs212;
      }
      let _1281_i5;
      _1281_i5 = _dafny.ZERO;
      L5: {
        while (((_this).f24).isEqualTo((_this).f24)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1281_i5)) {
              break L5;
            }
            _1281_i5 = (_1281_i5).plus(_dafny.ONE);
            let _pat_let_tv15 = _1233_v2;
            let _1282_v40;
            _1282_v40 = _dafny.MultiSet.fromElements(function (_pat_let18_0) {
              return function (_1283_dt__update__tmp_h1) {
                return function (_pat_let19_0) {
                  return function (_1284_dt__update_hcf20_h0) {
                    return _module.D3.create_DC11((_1283_dt__update__tmp_h1).dtor_cf18, (_1283_dt__update__tmp_h1).dtor_cf19, _1284_dt__update_hcf20_h0);
                  }(_pat_let19_0);
                }(_pat_let_tv15);
              }(_pat_let18_0);
            }(_module.D3.create_DC11(_module.__default.fm0((_this).f24, new BigNumber(-151), globalState), false, _1233_v2)), _module.D3.create_DC11((_this).f24, _1233_v2, _1233_v2));
            if (!((_1282_v40).IsSubsetOf(_module.__default.fm38(globalState)))) {
              let _index156 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              ((_this).f28)[_index156] = new BigNumber((_1248_v15).length);
              let _index157 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              let _rhs213 = (((_this).f24).plus((_this).f24)).plus(_module.__default.fm0((_this).f24, (_this).f24, globalState));
              let _rhs214 = (_this).f24;
              let _lhs166 = (_this).f28;
              let _lhs167 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              let _lhs168 = globalState;
              _lhs166[_lhs167] = _rhs213;
              _lhs168.f14 = _rhs214;
              (globalState).f17 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_1285_v6) => function (_1286_i6) {
                return _1285_v6;
              })(_1237_v6)), _1239_v8);
              r0 = _1233_v2;
              let _index158 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              let _index159 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              let _rhs215 = _dafny.Seq.UnicodeFromString("f");
              let _rhs216 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(337), (_dafny.ZERO).minus((_this).f24)));
              let _rhs217 = new BigNumber((_1255_v19).cardinality());
              let _lhs169 = globalState;
              let _lhs170 = (_this).f28;
              let _lhs171 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              let _lhs172 = (_this).f28;
              let _lhs173 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length));
              _lhs169.f17 = _rhs215;
              _lhs170[_lhs171] = _rhs216;
              _lhs172[_lhs173] = _rhs217;
              let _1287_v41;
              let _nw178 = Array((new BigNumber(28)).toNumber());
              _1287_v41 = _nw178;
              let _1288_v42;
              let _nw179 = new _module.C2();
              _nw179.__ctor(_1233_v2, ((_this).f28)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length))], (_this).f25);
              _1288_v42 = _nw179;
              let _1289_v43;
              _1289_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1234_v3)[_module.__default.safeIndex(((_this).f28)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f28).length))], new BigNumber((_1234_v3).length))],_1288_v42);
              let _index160 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1287_v41).length));
              (_1287_v41)[_index160] = (((_1289_v43).contains(_1233_v2)) ? ((_1289_v43).get(_1233_v2)) : (_1288_v42));
              let _1290_v44;
              _1290_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1240_v9,_1288_v42);
              let _index161 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1287_v41).length));
              (_1287_v41)[_index161] = (((_1290_v44).contains((_1240_v9).Merge(_1240_v9))) ? ((_1290_v44).get((_1240_v9).Merge(_1240_v9))) : (_1288_v42));
            } else {
              let _1291_v45;
              let _nw180 = Array((new BigNumber(28)).toNumber()).fill([]);
              _1291_v45 = _nw180;
              let _1292_v46;
              _1292_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1233_v2);
              let _1293_v47;
              _1293_v47 = _module.D1.create_DC3(_module.__default.fm4(false, (_this).f24, _1237_v6, _1233_v2, globalState));
              let _1294_v48;
              _1294_v48 = _module.D8.create_DC26(_1293_v47, _dafny.Seq.UnicodeFromString("b"));
              let _1295_v49;
              let _nw181 = Array((new BigNumber(19)).toNumber());
              _nw181[(_dafny.ZERO).toNumber()] = _1239_v8;
              _nw181[(_dafny.ONE).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(2)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("an");
              _nw181[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_1296_v2, _1297_v6, _1298_v8) => function (_1299_i7) {
                return (_module.D1.create_DC4(_1296_v2, _1296_v2, _1297_v6, (_this).f24, new BigNumber((_dafny.Seq.of(new BigNumber((_1298_v8).length))).length))).dtor_cf10;
              })(_1233_v2, _1237_v6, _1239_v8));
              _nw181[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), ((_1300_v6) => function (_1301_i8) {
                return _1300_v6;
              })(_1237_v6));
              _nw181[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-797)), ((_1302_v6) => function (_1303_i9) {
                return _1302_v6;
              })(_1237_v6)), _module.__default.safeIndex(new BigNumber((_1292_v46).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-797)), ((_1304_v6) => function (_1305_i9) {
                return _1304_v6;
              })(_1237_v6))).length)), _1237_v6);
              _nw181[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("kncwpm");
              _nw181[(new BigNumber(8)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("mpjutg");
              _nw181[(new BigNumber(10)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("wx");
              _nw181[(new BigNumber(12)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(13)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), ((_1306_v6) => function (_1307_i10) {
                return _1306_v6;
              })(_1237_v6));
              _nw181[(new BigNumber(15)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(16)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(17)).toNumber()] = _1239_v8;
              _nw181[(new BigNumber(18)).toNumber()] = (_1294_v48).dtor_cf42;
              _1295_v49 = _nw181;
              let _1308_v50;
              _1308_v50 = _module.D10.create_DC28(_1295_v49);
              let _index162 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1291_v45).length));
              (_1291_v45)[_index162] = (_1308_v50).dtor_cf44;
              let _1309_v51;
              _1309_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v2,_1295_v49);
              let _index163 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1291_v45).length));
              (_1291_v45)[_index163] = (((_1309_v51).contains(((_this).f24).isLessThan((_dafny.ZERO).minus((_this).f24)))) ? ((_1309_v51).get(((_this).f24).isLessThan((_dafny.ZERO).minus((_this).f24)))) : (_1295_v49));
              _1235_v4 = _1235_v4;
              let _1310_v52;
              let _nw182 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
              _1310_v52 = _nw182;
              let _1311_v53;
              _1311_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_this).f24);
              let _1312_v54;
              let _nw183 = new _module.C3();
              _nw183.__ctor(_1310_v52, ((_1233_v2) ? (_1244_v14) : (_1244_v14)), _module.__default.safeDivisionInt((_module.D8.create_DC25((_this).f24, _1233_v2, _1233_v2)).dtor_cf38, (_this).f24), _dafny.Seq.update(_dafny.Seq.Concat((_this).f25, (_this).f25), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Concat((_this).f25, (_this).f25)).length)), _1311_v53));
              _1312_v54 = _nw183;
              let _1313_v56;
              _1313_v56 = _module.D1.create_DC4(_1233_v2, _1233_v2, _1237_v6, (_this).f24, new BigNumber((function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of _dafny.IntegerRange(new BigNumber(784), new BigNumber(805))) {
    let _1314_v55 = _compr_36;
    if (((new BigNumber(784)).isLessThanOrEqualTo(_1314_v55)) && ((_1314_v55).isLessThan(new BigNumber(805)))) {
      _coll36.push([(_1314_v55).plus((_this).f24),_1237_v6]);
    }
  }
  return _coll36;
}()).length));
              r0 = ((_this).f24).isLessThanOrEqualTo((_1313_v56).dtor_cf12);
              (globalState).f2 = !(_module.__default.fm2(((_this).f24).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), ((_1315_v2) => function (_1316_i11) {
                return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1315_v2)).length));
              })(_1233_v2)), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), ((_1317_v2) => function (_1318_i11) {
                return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1317_v2)).length));
              })(_1233_v2))).length)), (_this).f24)).length)), new BigNumber(-970), _1237_v6, globalState));
            }
            (globalState).f2 = false;
            let _1319_v57;
            let _nw184 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
            _1319_v57 = _nw184;
            let _1320_v58;
            _1320_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_this).f24);
            let _1321_v59;
            let _nw185 = new _module.C3();
            _nw185.__ctor(_1319_v57, _1244_v14, _module.__default.safeModuloInt((_this).f24, (_this).f24), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_this).f24), _1320_v58));
            _1321_v59 = _nw185;
            let _1322_v60;
            _1322_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v12,false);
            let _1323_v61;
            let _nw186 = new _module.C2();
            _nw186.__ctor((((_1322_v60).contains(_1242_v12)) ? ((_1322_v60).get(_1242_v12)) : (!(!(_1233_v2)))), (_this).f24, _dafny.Seq.update((_this).f25, _module.__default.safeIndex((_this).f24, new BigNumber(((_this).f25).length)), _1320_v58));
            _1323_v61 = _nw186;
          }
        }
      }
      let _1324_i12;
      _1324_i12 = _dafny.ZERO;
      L6: {
        while (_1233_v2) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1324_i12)) {
              break L6;
            }
            _1324_i12 = (_1324_i12).plus(_dafny.ONE);
            (globalState).f18 = (_this).f24;
            (globalState).f14 = (_this).f24;
            let _1325_v62;
            let _nw187 = new _module.C2();
            _nw187.__ctor(!(_1233_v2), (_this).f24, (_this).f25);
            _1325_v62 = _nw187;
            if (_1233_v2) {
              let _1326_v65;
              let _nw188 = Array((new BigNumber(22)).toNumber());
              _nw188[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements((_this).f24);
              _nw188[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements((_this).f24, (_this).f24);
              _nw188[(new BigNumber(2)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(3)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements((_this).f24, (_this).f24);
              _nw188[(new BigNumber(5)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(6)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(7)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(8)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(9)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(10)).toNumber()] = function () {
                let _coll37 = new _dafny.Set();
                for (const _compr_37 of _dafny.IntegerRange(new BigNumber(678), new BigNumber(758))) {
                  let _1327_v63 = _compr_37;
                  if (((new BigNumber(678)).isLessThanOrEqualTo(_1327_v63)) && ((_1327_v63).isLessThan(new BigNumber(758)))) {
                    _coll37.add((_1327_v63).multipliedBy((_this).f24));
                  }
                }
                return _coll37;
              }();
              _nw188[(new BigNumber(11)).toNumber()] = function () {
                let _coll38 = new _dafny.Set();
                for (const _compr_38 of _dafny.IntegerRange(new BigNumber(494), new BigNumber(-646))) {
                  let _1328_v64 = _compr_38;
                  if (((new BigNumber(494)).isLessThanOrEqualTo(_1328_v64)) && ((_1328_v64).isLessThan(new BigNumber(-646)))) {
                    _coll38.add((_1328_v64).plus((_this).f24));
                  }
                }
                return _coll38;
              }();
              _nw188[(new BigNumber(12)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(13)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(14)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(15)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(16)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(17)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(18)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(19)).toNumber()] = _1248_v15;
              _nw188[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements((_this).f24);
              _nw188[(new BigNumber(21)).toNumber()] = _1248_v15;
              _1326_v65 = _nw188;
              let _1329_v66;
              let _nw189 = new _module.C3();
              _nw189.__ctor(_1326_v65, _1244_v14, new BigNumber((_1239_v8).length), (_this).f25);
              _1329_v66 = _nw189;
              let _1330_v67;
              _1330_v67 = _dafny.Seq.of((_this).f24);
              let _1331_v68;
              _1331_v68 = _dafny.Seq.of((_1330_v67)[_module.__default.safeIndex((_this).f24, new BigNumber((_1330_v67).length))]);
              (globalState).f2 = ((_1233_v2) ? (!(_dafny.Seq.IsProperPrefixOf(_1331_v68, _dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_1332_i13) {
                return new BigNumber(717);
              })))) : (_1325_v62.f31));
              (globalState).f17 = _1239_v8;
              let _1333_v69;
              _1333_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v12,_1233_v2);
              _1333_v69 = (_1333_v69).update(_1242_v12, _1325_v62.f31);
              let _1334_v70;
              _1334_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1330_v67,_1325_v62.f31);
              let _index164 = _module.__default.safeIndex(new BigNumber(638), new BigNumber(((_this).f28).length));
              ((_this).f28)[_index164] = _module.__default.safeModuloInt(new BigNumber(100), (_dafny.ZERO).minus(new BigNumber((_1334_v70).length)));
              let _index165 = _module.__default.safeIndex(new BigNumber(638), new BigNumber(((_this).f28).length));
              ((_this).f28)[_index165] = (_this).f24;
            } else {
              let _1335_v71;
              let _init30 = ((_1336_v15) => function (_1337_i14) {
                return _1336_v15;
              })(_1248_v15);
              let _nw190 = Array((new BigNumber(23)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw190.length); _i0_30++) {
                _nw190[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1335_v71 = _nw190;
              let _1338_v73;
              _1338_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,_1325_v62.f31);
              let _1339_v74;
              _1339_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_this).f24);
              let _1340_v75;
              let _nw191 = new _module.C3();
              _nw191.__ctor(_1335_v71, _1244_v14, (_this).f24, _dafny.Seq.of(function () {
                let _coll39 = new _dafny.Map();
                for (const _compr_39 of (_1338_v73).Keys.Elements) {
                  let _1341_v72 = _compr_39;
                  if ((_1338_v73).contains(_1341_v72)) {
                    _coll39.push([_1341_v72,new BigNumber((_dafny.Set.fromElements((_module.D3.create_DC11((_this).f24, _1325_v62.f31, true)).dtor_cf20)).length)]);
                  }
                }
                return _coll39;
              }(), _1339_v74, _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_dafny.ZERO).minus((_this).f24)), _1339_v74));
              _1340_v75 = _nw191;
              _1340_v75 = _1340_v75;
              let _1342_v76;
              _1342_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1325_v62.f31,_1234_v3);
              let _1343_v77;
              let _nw192 = new _module.C2();
              _nw192.__ctor(!(_1342_v76).contains(_1233_v2), (_this).f24, _dafny.Seq.of(_1339_v74));
              _1343_v77 = _nw192;
              let _index166 = _module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length));
              ((_this).f28)[_index166] = (_this).f24;
              let _index167 = _module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length));
              let _rhs218 = _1325_v62.f31;
              let _rhs219 = false;
              let _rhs220 = _1325_v62.f31;
              let _rhs221 = (_this).f24;
              let _rhs222 = (_this).f24;
              let _lhs174 = globalState;
              let _lhs175 = globalState;
              let _lhs176 = globalState;
              let _lhs177 = (_this).f28;
              let _lhs178 = _module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length));
              let _lhs179 = globalState;
              _lhs174.f2 = _rhs218;
              _lhs175.f2 = _rhs219;
              _lhs176.f2 = _rhs220;
              _lhs177[_lhs178] = _rhs221;
              _lhs179.f9 = _rhs222;
              let _index168 = _module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length));
              ((_this).f28)[_index168] = ((_this).f28)[_module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length))];
              (globalState).f9 = ((_this).f28)[_module.__default.safeIndex(new BigNumber(42), new BigNumber(((_this).f28).length))];
            }
          }
        }
      }
      let _1344_v78;
      _1344_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v2,_1238_v7);
      let _hi9 = new BigNumber((_1344_v78).length);
      for (let _1345_i15 = new BigNumber((_module.__default.fm39((_this).f24, _1233_v2, globalState)).length); _1345_i15.isLessThan(_hi9); _1345_i15 = _1345_i15.plus(_dafny.ONE)) {
        (globalState).f9 = (_this).f24;
        let _1346_v79;
        let _nw193 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
        _1346_v79 = _nw193;
        let _1347_v80;
        _1347_v80 = _dafny.Set.fromElements(true);
        let _index169 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1346_v79).length));
        (_1346_v79)[_index169] = (_module.__default.fm28(_1233_v2, (_dafny.ZERO).minus(_1345_i15), true, (_dafny.ZERO).minus(_1345_i15), globalState)).Intersect(_1347_v80);
        let _index170 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1346_v79).length));
        (_1346_v79)[_index170] = (_1347_v80).Union((_1347_v80).Difference(_dafny.Set.fromElements(_1233_v2)));
        if (!(_1345_i15).isEqualTo((_this).f24)) {
          let _1348_v81;
          let _init31 = ((_1349_i15, _1350_v8) => function (_1351_i16) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1349_i15, new BigNumber(297), _1349_i15, _1349_i15, _1349_i15)).length),_1350_v8);
          })(_1345_i15, _1239_v8);
          let _nw194 = Array((new BigNumber(16)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw194.length); _i0_31++) {
            _nw194[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1348_v81 = _nw194;
          let _1352_v82;
          _1352_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1239_v8);
          let _index171 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1348_v81).length));
          (_1348_v81)[_index171] = _1352_v82;
          let _index172 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1348_v81).length));
          (_1348_v81)[_index172] = (_1352_v82).update(new BigNumber(-230), _1239_v8);
          (globalState).f9 = _1345_i15;
          let _1353_v84;
          _1353_v84 = _dafny.MultiSet.fromElements((function () {
            let _coll40 = new _dafny.Set();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(719), new BigNumber(-613))) {
              let _1354_v83 = _compr_40;
              if (((new BigNumber(719)).isLessThanOrEqualTo(_1354_v83)) && ((_1354_v83).isLessThan(new BigNumber(-613)))) {
                _coll40.add(_module.__default.safeModuloInt(_1354_v83, (_this).f24));
              }
            }
            return _coll40;
          }()).equals(_1248_v15), _1233_v2, (_1233_v2) === (_1233_v2), !(_1233_v2) || (_1233_v2));
          let _1355_v85;
          let _nw195 = new _module.C2();
          _nw195.__ctor(!(_1233_v2), _module.__default.safeDivisionInt(_1345_i15, _1345_i15), _dafny.Seq.update((_this).f25, _module.__default.safeIndex(_1345_i15, new BigNumber(((_this).f25).length)), _dafny.Map.Empty.slice().updateUnsafe(_1237_v6,(_this).f24)));
          _1355_v85 = _nw195;
          let _rhs223 = _1345_i15;
          let _rhs224 = _dafny.MultiSet.fromElements(!(_1233_v2));
          let _rhs225 = ((((_1233_v2) ? (_1233_v2) : (_1233_v2))) ? (_1233_v2) : (!(_1355_v85.f31) || (_1355_v85.f31)));
          let _rhs226 = (_module.__default.fm0(_1345_i15, (_this).f24, globalState)).plus(_1345_i15);
          let _rhs227 = _1355_v85;
          let _lhs180 = globalState;
          let _lhs181 = globalState;
          _lhs180.f9 = _rhs223;
          _1353_v84 = _rhs224;
          _1233_v2 = _rhs225;
          _lhs181.f14 = _rhs226;
          _1355_v85 = _rhs227;
          let _index173 = _module.__default.safeIndex(new BigNumber(768), new BigNumber(((_this).f28).length));
          ((_this).f28)[_index173] = _1345_i15;
          let _index174 = _module.__default.safeIndex(new BigNumber(768), new BigNumber(((_this).f28).length));
          ((_this).f28)[_index174] = _1345_i15;
          let _1356_v86;
          _1356_v86 = _module.D11.create_DC30(_1240_v9);
          let _1357_v87;
          _1357_v87 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1239_v8, _1239_v8),_1233_v2);
          let _index175 = _module.__default.safeIndex(new BigNumber(768), new BigNumber(((_this).f28).length));
          let _rhs228 = ((_this).f28)[_module.__default.safeIndex(new BigNumber(768), new BigNumber(((_this).f28).length))];
          let _rhs229 = _1237_v6;
          let _rhs230 = (_dafny.MultiSet.fromElements(_1240_v9)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1240_v9, (_1356_v86).dtor_cf47, _1240_v9));
          let _rhs231 = (((_1357_v87).contains(_1239_v8)) ? ((_1357_v87).get(_1239_v8)) : (!(_1233_v2)));
          let _lhs182 = (_this).f28;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(768), new BigNumber(((_this).f28).length));
          let _lhs184 = _1355_v85;
          let _lhs185 = globalState;
          _lhs182[_lhs183] = _rhs228;
          _1237_v6 = _rhs229;
          _lhs184.f31 = _rhs230;
          _lhs185.f2 = _rhs231;
        } else {
          let _1358_v88;
          _1358_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1345_i15,_1242_v12);
          let _rhs232 = _1233_v2;
          let _rhs233 = _dafny.Seq.Concat(_1239_v8, _1239_v8);
          let _rhs234 = _1358_v88;
          let _lhs186 = globalState;
          let _lhs187 = globalState;
          _lhs186.f2 = _rhs232;
          _lhs187.f17 = _rhs233;
          _1358_v88 = _rhs234;
          let _index176 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1242_v12).length));
          (_1242_v12)[_index176] = (_1234_v3)[_module.__default.safeIndex(_1345_i15, new BigNumber((_1234_v3).length))];
          let _index177 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1242_v12).length));
          (_1242_v12)[_index177] = (_1345_i15).isEqualTo(new BigNumber(((_1235_v4).update(_1233_v2, (_this).f24)).length));
          let _1359_v89;
          let _init32 = ((_1360_v12, _1361_v2) => function (_1362_i17) {
            return !((_1360_v12)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1360_v12).length))]) || (_1361_v2);
          })(_1242_v12, _1233_v2);
          let _nw196 = Array((new BigNumber(28)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw196.length); _i0_32++) {
            _nw196[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1359_v89 = _nw196;
          _1359_v89 = _1359_v89;
          let _1363_v90;
          _1363_v90 = _module.D8.create_DC25(_1345_i15, _1233_v2, (_1242_v12)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1242_v12).length))]);
          let _1364_v91;
          _1364_v91 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_1239_v8, _module.__default.safeIndex((_this).f24, new BigNumber((_1239_v8).length)), new _dafny.CodePoint('a'.codePointAt(0))), _module.__default.fm4((_1242_v12)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1242_v12).length))], _1345_i15, _1237_v6, (_1363_v90).dtor_cf40, globalState), _dafny.Seq.UnicodeFromString("vhh"), _dafny.Seq.Concat(_1239_v8, _1239_v8), _1239_v8);
          let _1365_v92;
          _1365_v92 = _dafny.Seq.of(_1364_v91, (((_1242_v12)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1242_v12).length))]) ? (_dafny.MultiSet.fromElements(_1239_v8, _1239_v8)) : (_1364_v91)), _dafny.MultiSet.fromElements(_1239_v8, _1239_v8, _dafny.Seq.UnicodeFromString("lqkivho")), ((_1364_v91).update(_1239_v8, _module.__default.abs(_1345_i15))).Intersect(_1364_v91));
          _1364_v91 = (_1365_v92)[_module.__default.safeIndex((_this).f24, new BigNumber((_1365_v92).length))];
          _1239_v8 = _dafny.Seq.Concat(_1239_v8, _1239_v8);
        }
        let _index178 = _module.__default.safeIndex(new BigNumber(384), new BigNumber(((_this).f28).length));
        ((_this).f28)[_index178] = (_this).f24;
        let _index179 = _module.__default.safeIndex(new BigNumber(384), new BigNumber(((_this).f28).length));
        ((_this).f28)[_index179] = (_this).f24;
      }
      let _1366_v93;
      _1366_v93 = _dafny.Seq.of(_1232_v1, _1232_v1);
      let _1367_v94;
      let _nw197 = new _module.C1();
      _nw197.__ctor(_1366_v93, (_this).f24, (_this).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(526)), ((_1368_v6) => function (_1369_i18) {
        return _dafny.Map.Empty.slice().updateUnsafe(_1368_v6,(_this).f24);
      })(_1237_v6)));
      _1367_v94 = _nw197;
      let _1370_v95;
      _1370_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1367_v94);
      let _1371_v96;
      _1371_v96 = _dafny.Seq.of(_module.__default.fm0((_this).f24, new BigNumber((_1370_v95).length), globalState));
      r0 = (new BigNumber(140)).isEqualTo((_1371_v96)[_module.__default.safeIndex(new BigNumber((_1248_v15).length), new BigNumber((_1371_v96).length))]);
      return r0;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let r3 = false;
      let _1372_v0;
      _1372_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f24);
      let _1373_v1;
      _1373_v1 = _dafny.Seq.of(_1372_v0);
      let _1374_v2;
      _1374_v2 = _dafny.MultiSet.fromElements(p1, p1);
      let _1375_v3;
      _1375_v3 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1376_v4;
      _1376_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1375_v3)).length),new BigNumber((_dafny.Seq.of(p1)).length));
      let _1377_v5;
      _1377_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.D3.create_DC11((_this).f24, p1, p1));
      let _1378_v6;
      _1378_v6 = _module.D12.create_DC32(_1377_v5);
      let _1379_v7;
      _1379_v7 = _dafny.Set.fromElements(p0);
      let _1380_v8;
      let _nw198 = Array((new BigNumber(16)).toNumber());
      _nw198[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1373_v1, _1373_v1)).length);
      _nw198[(_dafny.ONE).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(2)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(3)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(4)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(5)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(6)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(7)).toNumber()] = (((_1374_v2).contains(p0)) ? ((_1374_v2).get(p0)) : (new BigNumber((_1376_v4).length)));
      _nw198[(new BigNumber(8)).toNumber()] = ((p1) ? ((_dafny.ZERO).minus((_this).f24)) : ((_this).f24));
      _nw198[(new BigNumber(9)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(10)).toNumber()] = ((_this).f24).minus(new BigNumber((_dafny.Seq.of(p1)).length));
      _nw198[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((_1378_v6).dtor_cf53).length), (_this).f24);
      _nw198[(new BigNumber(12)).toNumber()] = new BigNumber(((_1379_v7).Intersect(_1379_v7)).length);
      _nw198[(new BigNumber(13)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(14)).toNumber()] = (_this).f24;
      _nw198[(new BigNumber(15)).toNumber()] = (_this).f24;
      _1380_v8 = _nw198;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1380_v8).length))) {
        let _1381_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1381_i0)) && ((_1381_i0).isLessThan(new BigNumber((_1380_v8).length))))) {
          (_1380_v8)[(_1381_i0)] = (_1381_i0).multipliedBy(_module.__default.safeModuloInt((_this).f24, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("spivrnx")).length))));
        }
      }
      let _1382_v9;
      let _nw199 = new _module.C0();
      _nw199.__ctor(true);
      _1382_v9 = _nw199;
      let _hi10 = (_dafny.ZERO).minus((_this).f24);
      for (let _1383_i1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p1)).length)); _1383_i1.isLessThan(_hi10); _1383_i1 = _1383_i1.plus(_dafny.ONE)) {
        if ((_1382_v9).f34) {
          _1375_v3 = _1375_v3;
          let _1384_v10;
          _1384_v10 = _dafny.Set.fromElements((_this).f28);
          let _1385_v11;
          _1385_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1386_v12;
          _1386_v12 = _dafny.Seq.of((_1382_v9).f34);
          let _1387_v13;
          let _nw200 = Array((new BigNumber(27)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = p1;
          _nw200[(_dafny.ONE).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(2)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(3)).toNumber()] = p1;
          _nw200[(new BigNumber(4)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(5)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(6)).toNumber()] = _module.__default.fm2(new BigNumber(111), _1383_i1, _1375_v3, globalState);
          _nw200[(new BigNumber(7)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(8)).toNumber()] = (((_1385_v11).contains(p1)) ? ((_1385_v11).get(p1)) : (p0));
          _nw200[(new BigNumber(9)).toNumber()] = p1;
          _nw200[(new BigNumber(10)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(11)).toNumber()] = p0;
          _nw200[(new BigNumber(12)).toNumber()] = p0;
          _nw200[(new BigNumber(13)).toNumber()] = false;
          _nw200[(new BigNumber(14)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(15)).toNumber()] = true;
          _nw200[(new BigNumber(16)).toNumber()] = false;
          _nw200[(new BigNumber(17)).toNumber()] = p1;
          _nw200[(new BigNumber(18)).toNumber()] = p0;
          _nw200[(new BigNumber(19)).toNumber()] = !(true);
          _nw200[(new BigNumber(20)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(21)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(22)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(23)).toNumber()] = (_1386_v12)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1382_v9).f34,(_1382_v9).f34)).length), new BigNumber((_1386_v12).length))];
          _nw200[(new BigNumber(24)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(25)).toNumber()] = (_1382_v9).f34;
          _nw200[(new BigNumber(26)).toNumber()] = (_module.D6.create_DC17(p0)).dtor_cf28;
          _1387_v13 = _nw200;
          let _1388_v14;
          _1388_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1384_v10).length),_1387_v13);
          _1388_v14 = (_1388_v14).update(_module.__default.fm0(new BigNumber(951), new BigNumber(-204), globalState), _1387_v13);
          (globalState).f2 = true;
          let _1389_v15;
          let _nw201 = new _module.C0();
          _nw201.__ctor((_1382_v9).f34);
          _1389_v15 = _nw201;
          (globalState).f1 = (_this).f24;
        } else {
          (globalState).f18 = (_this).f24;
          let _index180 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f28).length));
          ((_this).f28)[_index180] = _module.__default.fm0((_dafny.ZERO).minus(_module.__default.fm0((_this).f24, new BigNumber(572), globalState)), (_this).f24, globalState);
          let _1390_v16;
          let _nw202 = Array((new BigNumber(27)).toNumber()).fill(_module.D4.Default());
          _1390_v16 = _nw202;
          let _1391_v17;
          _1391_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1390_v16,(_1382_v9).f34);
          let _index181 = _module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f28).length));
          ((_this).f28)[_index181] = (((p0) ? (_1383_i1) : (new BigNumber(-994)))).minus(new BigNumber((_1391_v17).length));
          let _1392_v18;
          let _init33 = ((_1393_v9, _1394_p0) => function (_1395_i2) {
            return ((_1393_v9).f34) || (_1394_p0);
          })(_1382_v9, p0);
          let _nw203 = Array((new BigNumber(17)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw203.length); _i0_33++) {
            _nw203[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1392_v18 = _nw203;
          let _index182 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_1392_v18).length));
          (_1392_v18)[_index182] = (_1382_v9).f34;
          let _index183 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_1392_v18).length));
          (_1392_v18)[_index183] = (_1382_v9).f34;
          _1392_v18 = _1392_v18;
          let _1396_v19;
          let _nw204 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
          _1396_v19 = _nw204;
          let _1397_v20;
          let _nw205 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _1397_v20 = _nw205;
          let _1398_v21;
          _1398_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1375_v3,_1383_i1);
          let _1399_v22;
          let _nw206 = new _module.C3();
          _nw206.__ctor(_1396_v19, _1397_v20, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-388)), ((_this).f28)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f28).length))]), _dafny.Seq.update(_dafny.Seq.Concat((_this).f25, _dafny.Seq.update((_this).f25, _module.__default.safeIndex(_1383_i1, new BigNumber(((_this).f25).length)), _1398_v21)), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Concat((_this).f25, _dafny.Seq.update((_this).f25, _module.__default.safeIndex(_1383_i1, new BigNumber(((_this).f25).length)), _1398_v21))).length)), (_1398_v21).update(_1375_v3, ((_this).f28)[_module.__default.safeIndex(new BigNumber(694), new BigNumber(((_this).f28).length))])));
          _1399_v22 = _nw206;
          _1399_v22 = _1399_v22;
        }
        let _1400_v23;
        _1400_v23 = _dafny.Seq.UnicodeFromString("epg");
        let _1401_v24;
        _1401_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f24, new BigNumber((_1400_v23).length)),true);
        let _1402_v25;
        _1402_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1382_v9).f34,_1379_v7);
        _1401_v24 = (_1401_v24).update((((_1382_v9).f34) ? (_1383_i1) : (new BigNumber((_1402_v25).length))), !((_1382_v9).f34) || (!(p1)));
        r1 = new BigNumber(602);
        let _1403_v26;
        _1403_v26 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f24));
        _1401_v24 = (_1401_v24).update(new BigNumber((_1403_v26).cardinality()), false);
      }
      let _1404_v27;
      _1404_v27 = _dafny.Seq.UnicodeFromString("wlyo");
      let _index184 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length));
      (_1380_v8)[_index184] = _module.__default.safeModuloInt((_this).f24, (_dafny.ZERO).minus(new BigNumber((_1404_v27).length)));
      let _1405_v28;
      _1405_v28 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1404_v27).length)));
      let _index185 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length));
      (_1380_v8)[_index185] = (_module.__default.fm0((_this).f24, (_this).f24, globalState)).multipliedBy((_1405_v28)[_module.__default.safeIndex((_this).f24, new BigNumber((_1405_v28).length))]);
      let _1406_v29;
      _1406_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1375_v3,false);
      let _index186 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length));
      (_1380_v8)[_index186] = _module.__default.fm0((new BigNumber((_1406_v29).length)).multipliedBy((_this).f24), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), ((_1407_v3) => function (_1408_i3) {
        return _1407_v3;
      })(_1375_v3))).length), globalState);
      let _1409_v30;
      let _nw207 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
      _1409_v30 = _nw207;
      let _1410_v31;
      _1410_v31 = _dafny.Seq.of(p0);
      let _index187 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1409_v30).length));
      (_1409_v30)[_index187] = (((_1382_v9).f34) ? (_1410_v31) : (_1410_v31));
      let _index188 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1409_v30).length));
      (_1409_v30)[_index188] = _dafny.Seq.Concat(_1410_v31, _1410_v31);
      r0 = (_1380_v8)[_module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length))];
      r1 = (_this).f24;
      let _1411_v32;
      _1411_v32 = _dafny.Seq.of(_1405_v28, _1405_v28);
      let _1412_v33;
      _1412_v33 = _module.D4.create_DC12(_dafny.Seq.of((_1380_v8)[_module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length))], (_1380_v8)[_module.__default.safeIndex(new BigNumber(153), new BigNumber((_1380_v8).length))], (_this).f24, (_this).f24));
      let _1413_v34;
      _1413_v34 = _dafny.Seq.of(_1405_v28, (_1411_v32)[_module.__default.safeIndex((_this).f24, new BigNumber((_1411_v32).length))], (_1412_v33).dtor_cf21, _dafny.Seq.Concat(_1405_v28, _1405_v28));
      r2 = (_1411_v32)[_module.__default.safeIndex((_this).f24, new BigNumber((_1411_v32).length))];
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f24 = _dafny.ZERO;
      this._f25 = _dafny.Seq.of();
      this.f27 = _dafny.ZERO;
      this._f26 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f26, f27, f24, f25) {
      let _this = this;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let _hi11 = _this.f27;
      for (let _1414_i0 = _module.__default.safeModuloInt(_this.f27, _this.f27); _1414_i0.isLessThan(_hi11); _1414_i0 = _1414_i0.plus(_dafny.ONE)) {
        (globalState).f14 = ((_this).f24).minus(_this.f27);
        if (_module.__default.fm2((_this.f27).plus(new BigNumber(-676)), (_this).f24, _module.__default.fm3(globalState), globalState)) {
          let _1415_v0;
          let _init34 = function (_1416_i1) {
            return !(!(true));
          };
          let _nw208 = Array((new BigNumber(15)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw208.length); _i0_34++) {
            _nw208[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1415_v0 = _nw208;
          let _1417_v1;
          _1417_v1 = false;
          let _index189 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1415_v0).length));
          (_1415_v0)[_index189] = _1417_v1;
          let _index190 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1415_v0).length));
          (_1415_v0)[_index190] = !(_1417_v1) || (_1417_v1);
          let _1418_v2;
          let _nw209 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _1418_v2 = _nw209;
          let _1419_v3;
          let _nw210 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1419_v3 = _nw210;
          let _1420_v4;
          let _nw211 = new _module.C4();
          _nw211.__ctor(_1419_v3, new BigNumber((_dafny.Seq.of(_this.f27)).length), (_this).f25);
          _1420_v4 = _nw211;
          let _index191 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1418_v2).length));
          (_1418_v2)[_index191] = _dafny.Seq.of(_1420_v4);
          let _1421_v5;
          _1421_v5 = _dafny.MultiSet.fromElements(_1417_v1);
          let _1422_v6;
          _1422_v6 = _dafny.Seq.of(_1420_v4, _1420_v4);
          let _index192 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1418_v2).length));
          (_1418_v2)[_index192] = (((_1421_v5).IsDisjointFrom(_1421_v5)) ? (_dafny.Seq.Concat(_1422_v6, _1422_v6)) : (_1422_v6));
          let _1423_v7;
          let _nw212 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1423_v7 = _nw212;
          let _1424_v8;
          _1424_v8 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index193 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1423_v7).length));
          (_1423_v7)[_index193] = _1424_v8;
          let _index194 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1423_v7).length));
          (_1423_v7)[_index194] = _1424_v8;
          let _1425_v9;
          _1425_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1420_v4).f24,(_1420_v4).f24);
          let _1426_v10;
          _1426_v10 = _dafny.Seq.UnicodeFromString("iccnie");
          let _1427_v11;
          _1427_v11 = _dafny.Seq.of(new BigNumber((_1426_v10).length));
          let _1428_v12;
          _1428_v12 = _dafny.MultiSet.fromElements((((_1425_v9).contains((_1420_v4).f24)) ? ((_1425_v9).get((_1420_v4).f24)) : ((_1420_v4).f24)), (_1427_v11)[_module.__default.safeIndex((_this).f24, new BigNumber((_1427_v11).length))]);
          let _1429_v13;
          let _out31;
          _out31 = _module.__default.m0((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1414_i0, new BigNumber((_1428_v12).cardinality()))), _module.__default.safeDivisionInt(_1414_i0, _module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_1426_v10).length)), _1414_i0, globalState)), globalState);
          _1429_v13 = _out31;
          let _index195 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1415_v0).length));
          (_1415_v0)[_index195] = _1417_v1;
        } else {
          (globalState).f9 = _this.f27;
          let _1430_v14;
          _1430_v14 = true;
          let _1431_v15;
          let _nw213 = new _module.C0();
          _nw213.__ctor(_1430_v14);
          _1431_v15 = _nw213;
          let _1432_v16;
          _1432_v16 = _module.D6.create_DC18(!((_1431_v15).f34));
          let _1433_v17;
          _1433_v17 = _module.D6.create_DC19(_1432_v16);
          let _1434_v18;
          _1434_v18 = _dafny.Seq.UnicodeFromString("rflync");
          let _1435_v19;
          _1435_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1434_v18,_module.D6.create_DC18((_1431_v15).f34));
          let _pat_let_tv16 = _1432_v16;
          let _pat_let_tv17 = _1435_v19;
          let _pat_let_tv18 = _1434_v18;
          let _pat_let_tv19 = _1432_v16;
          let _pat_let_tv20 = _1435_v19;
          let _pat_let_tv21 = _1434_v18;
          let _pat_let_tv22 = _1432_v16;
          let _1436_v20;
          let _nw214 = Array((new BigNumber(23)).toNumber());
          _nw214[(_dafny.ZERO).toNumber()] = _1433_v17;
          _nw214[(_dafny.ONE).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(2)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(3)).toNumber()] = _module.D6.create_DC19(_1432_v16);
          _nw214[(new BigNumber(4)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(5)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(6)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(7)).toNumber()] = _module.D6.create_DC19(_1432_v16);
          _nw214[(new BigNumber(8)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(9)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(10)).toNumber()] = _module.D6.create_DC19(_1432_v16);
          _nw214[(new BigNumber(11)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(12)).toNumber()] = function (_pat_let20_0) {
            return function (_1437_dt__update__tmp_h0) {
              return function (_pat_let21_0) {
                return function (_1438_dt__update_hcf30_h0) {
                  return _module.D6.create_DC19(_1438_dt__update_hcf30_h0);
                }(_pat_let21_0);
              }(_pat_let_tv16);
            }(_pat_let20_0);
          }(_1433_v17);
          _nw214[(new BigNumber(13)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(14)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(15)).toNumber()] = ((false) ? (_1433_v17) : (_1433_v17));
          _nw214[(new BigNumber(16)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(17)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(18)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(19)).toNumber()] = function (_pat_let22_0) {
            return function (_1439_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_1440_dt__update_hcf30_h1) {
                  return _module.D6.create_DC19(_1440_dt__update_hcf30_h1);
                }(_pat_let23_0);
              }((((_pat_let_tv20).contains(_pat_let_tv21)) ? ((_pat_let_tv17).get(_pat_let_tv18)) : (_pat_let_tv19)));
            }(_pat_let22_0);
          }(_module.D6.create_DC19(_1432_v16));
          _nw214[(new BigNumber(20)).toNumber()] = function (_pat_let24_0) {
            return function (_1441_dt__update__tmp_h2) {
              return function (_pat_let25_0) {
                return function (_1442_dt__update_hcf30_h2) {
                  return _module.D6.create_DC19(_1442_dt__update_hcf30_h2);
                }(_pat_let25_0);
              }(_pat_let_tv22);
            }(_pat_let24_0);
          }(_1433_v17);
          _nw214[(new BigNumber(21)).toNumber()] = _1433_v17;
          _nw214[(new BigNumber(22)).toNumber()] = _1433_v17;
          _1436_v20 = _nw214;
          _1436_v20 = _1436_v20;
          let _1443_v21;
          _1443_v21 = _dafny.Seq.of((_1431_v15).f34);
          (globalState).f21 = (((_1431_v15).f34) ? (_1443_v21) : (_dafny.Seq.Concat(_1443_v21, _1443_v21)));
          let _1444_v22;
          let _nw215 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1444_v22 = _nw215;
          let _1445_v23;
          _1445_v23 = _dafny.MultiSet.fromElements(_1444_v22, _1444_v22, _1444_v22, _1444_v22);
          _1445_v23 = (_1445_v23).update(_1444_v22, _module.__default.abs(new BigNumber(((((_1431_v15).f34) ? (_dafny.Set.fromElements(_1444_v22)) : (_dafny.Set.fromElements(_1444_v22, _1444_v22)))).length)));
        }
        (globalState).f1 = _module.__default.safeModuloInt(new BigNumber(891), _this.f27);
        let _1446_v24;
        _1446_v24 = false;
        if (_1446_v24) {
          _1446_v24 = _1446_v24;
          let _1447_v25;
          let _out32;
          _out32 = (_this).m2(globalState);
          _1447_v25 = _out32;
          let _1448_v26;
          let _init35 = ((_1449_i0) => function (_1450_i2) {
            return (_1450_i2).plus(_1449_i0);
          })(_1414_i0);
          let _nw216 = Array((new BigNumber(18)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw216.length); _i0_35++) {
            _nw216[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1448_v26 = _nw216;
          let _index196 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_1448_v26).length));
          (_1448_v26)[_index196] = _this.f27;
          let _index197 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_1448_v26).length));
          (_1448_v26)[_index197] = _1414_i0;
          let _1451_v27;
          _1451_v27 = new _dafny.CodePoint('o'.codePointAt(0));
          let _1452_v28;
          _1452_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1451_v27,_1414_i0);
          let _1453_v29;
          let _nw217 = new _module.C4();
          _nw217.__ctor(_1448_v26, (_1448_v26)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_1448_v26).length))], _dafny.Seq.update(_dafny.Seq.Concat((_this).f25, (_this).f25), _module.__default.safeIndex(new BigNumber(-423), new BigNumber((_dafny.Seq.Concat((_this).f25, (_this).f25)).length)), _1452_v28));
          _1453_v29 = _nw217;
          (globalState).f9 = (_this).f24;
        } else {
          let _1454_v30;
          _1454_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1446_v24,_1414_i0);
          let _1455_v31;
          let _nw218 = Array((new BigNumber(2)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = _this.f27;
          _nw218[(_dafny.ONE).toNumber()] = (((_1454_v30).contains(_1446_v24)) ? ((_1454_v30).get(_1446_v24)) : ((_dafny.ZERO).minus(_this.f27)));
          _1455_v31 = _nw218;
          let _1456_v32;
          let _nw219 = new _module.C4();
          _nw219.__ctor(_1455_v31, (_this).f24, (_this).f25);
          _1456_v32 = _nw219;
          _1456_v32 = _1456_v32;
          let _rhs235 = _1446_v24;
          let _rhs236 = _module.__default.safeModuloInt(_this.f27, _1414_i0);
          let _rhs237 = ((_1446_v24) ? (false) : (false));
          let _lhs188 = globalState;
          let _lhs189 = globalState;
          let _lhs190 = globalState;
          _lhs188.f2 = _rhs235;
          _lhs189.f18 = _rhs236;
          _lhs190.f2 = _rhs237;
          let _1457_v33;
          _1457_v33 = _dafny.Seq.of(_1446_v24);
          let _1458_v34;
          _1458_v34 = _dafny.MultiSet.fromElements(new BigNumber((_1457_v33).length), (_this).f24, _1414_i0, ((_1456_v32).f24).multipliedBy((_this).f24), (new BigNumber(564)).plus(_this.f27));
          _1458_v34 = _1458_v34;
          (globalState).f17 = _dafny.Seq.UnicodeFromString("tbquywj");
          let _1459_v35;
          let _nw220 = new _module.C4();
          _nw220.__ctor(_1455_v31, (_1456_v32).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_1460_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),(_this).f24);
          }));
          _1459_v35 = _nw220;
          let _1461_v36;
          _1461_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(964),_1459_v35);
          _1461_v36 = _1461_v36;
        }
      }
      let _1462_v37;
      _1462_v37 = false;
      let _1463_v38;
      _1463_v38 = _dafny.Seq.of((_this).f24);
      let _pat_let_tv23 = _1462_v37;
      let _pat_let_tv24 = _1463_v38;
      let _1464_v39;
      _1464_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1462_v37,function (_pat_let26_0) {
        return function (_1465_dt__update__tmp_h3) {
          return function (_pat_let27_0) {
            return function (_1466_dt__update_hcf20_h0) {
              return function (_pat_let28_0) {
                return function (_1467_dt__update_hcf18_h0) {
                  return _module.D3.create_DC11(_1467_dt__update_hcf18_h0, (_1465_dt__update__tmp_h3).dtor_cf19, _1466_dt__update_hcf20_h0);
                }(_pat_let28_0);
              }(new BigNumber((_pat_let_tv24).length));
            }(_pat_let27_0);
          }(_pat_let_tv23);
        }(_pat_let26_0);
      }(_module.D3.create_DC11(_this.f27, _1462_v37, _1462_v37)));
      let _1468_v40;
      _1468_v40 = _module.D12.create_DC32(_1464_v39);
      let _source15 = _1468_v40;
      if (_source15.is_DC33) {
        let _1469___mcc_h0 = (_source15).cf54;
        let _1470___mcc_h1 = (_source15).cf55;
        let _1471_cf55 = _1470___mcc_h1;
        let _1472_cf54 = _1469___mcc_h0;
        let _1473_v41;
        _1473_v41 = _dafny.Set.fromElements(_this.f27);
        _1471_cf55 = !(!(((_dafny.Set.fromElements(_this.f27, (_this).f24)).Difference(_1473_v41)).IsSubsetOf(_1473_v41)));
        let _1474_v42;
        _1474_v42 = _dafny.Seq.UnicodeFromString("tdwmjxld");
        r0 = _module.__default.fm2((new BigNumber((_1474_v42).length)).multipliedBy((_this).f24), (_this.f27).multipliedBy(new BigNumber(891)), _1472_cf54, globalState);
        let _1475_v43;
        let _nw221 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1475_v43 = _nw221;
        let _index198 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1475_v43).length));
        (_1475_v43)[_index198] = _this.f27;
        let _1476_v44;
        _1476_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1471_cf55,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), function (_1477_i4) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length));
        let _index199 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1475_v43).length));
        (_1475_v43)[_index199] = (new BigNumber((_1476_v44).length)).plus(_module.__default.safeModuloInt((_this).f24, _this.f27));
        (_this).f27 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_1478_i5) {
          return (_this).f24;
        })).length));
      } else if (_source15.is_DC32) {
        let _1479___mcc_h2 = (_source15).cf53;
        let _1480_cf53 = _1479___mcc_h2;
        let _1481_v45;
        _1481_v45 = _dafny.Seq.of(true, _1462_v37);
        let _1482_v46;
        _1482_v46 = _dafny.Set.fromElements(_1481_v45);
        r0 = (_1482_v46).IsProperSubsetOf(_1482_v46);
        let _1483_v47;
        _1483_v47 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1463_v38));
        let _1484_v48;
        let _nw222 = new _module.C1();
        _nw222.__ctor(_1483_v47, _this.f27, _this.f27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-502)), function (_1485_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),(_this).f24);
        }));
        _1484_v48 = _nw222;
        let _1486_v49;
        _1486_v49 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1487_v51;
        _1487_v51 = _dafny.MultiSet.fromElements((_this).f24);
        let _1488_v52;
        _1488_v52 = _module.D13.create_DC35(_1487_v51);
        let _1489_v53;
        _1489_v53 = _dafny.Seq.of(_module.D3.create_DC11(new BigNumber(950), _1462_v37, _module.__default.fm2(new BigNumber(543), _this.f27, _1486_v49, globalState)));
        let _1490_v54;
        _1490_v54 = _dafny.MultiSet.fromElements(_1462_v37, _1462_v37);
        let _1491_v55;
        let _nw223 = Array((new BigNumber(16)).toNumber());
        _nw223[(_dafny.ZERO).toNumber()] = !(true);
        _nw223[(_dafny.ONE).toNumber()] = _1462_v37;
        _nw223[(new BigNumber(2)).toNumber()] = _1462_v37;
        _nw223[(new BigNumber(3)).toNumber()] = _module.__default.fm2(_1484_v48.f33, _1484_v48.f33, _1486_v49, globalState);
        _nw223[(new BigNumber(4)).toNumber()] = _1462_v37;
        _nw223[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_1484_v48.f33, _1484_v48.f33)).IsDisjointFrom(function () {
          let _coll41 = new _dafny.Set();
          for (const _compr_41 of _dafny.IntegerRange(new BigNumber(794), new BigNumber(-319))) {
            let _1492_v50 = _compr_41;
            if (((new BigNumber(794)).isLessThanOrEqualTo(_1492_v50)) && ((_1492_v50).isLessThan(new BigNumber(-319)))) {
              _coll41.add(_module.__default.safeDivisionInt(_1492_v50, (_this).f24));
            }
          }
          return _coll41;
        }());
        _nw223[(new BigNumber(6)).toNumber()] = true;
        _nw223[(new BigNumber(7)).toNumber()] = true;
        _nw223[(new BigNumber(8)).toNumber()] = ((_1488_v52).dtor_cf57).IsDisjointFrom(_1487_v51);
        _nw223[(new BigNumber(9)).toNumber()] = ((_1487_v51).update(new BigNumber((_1489_v53).length), _module.__default.abs(_1484_v48.f33))).IsSubsetOf(_1487_v51);
        _nw223[(new BigNumber(10)).toNumber()] = !(_1484_v48.f33).isEqualTo(_this.f27);
        _nw223[(new BigNumber(11)).toNumber()] = (_1490_v54).IsDisjointFrom(_1490_v54);
        _nw223[(new BigNumber(12)).toNumber()] = !(_module.__default.fm2(new BigNumber((_1490_v54).cardinality()), _this.f27, _1486_v49, globalState));
        _nw223[(new BigNumber(13)).toNumber()] = _1462_v37;
        _nw223[(new BigNumber(14)).toNumber()] = (_1481_v45)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f27), new BigNumber((_1481_v45).length))];
        _nw223[(new BigNumber(15)).toNumber()] = _1462_v37;
        _1491_v55 = _nw223;
        let _index200 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1491_v55).length));
        (_1491_v55)[_index200] = (_this.f27).isLessThanOrEqualTo((_this).f24);
        let _index201 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1491_v55).length));
        (_1491_v55)[_index201] = _1462_v37;
        let _1493_v56;
        _1493_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1486_v49,_1462_v37);
        let _index202 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1491_v55).length));
        let _rhs238 = _module.__default.safeDivisionInt(_this.f27, _module.__default.safeDivisionInt(_1484_v48.f33, new BigNumber(541)));
        let _rhs239 = (true) && ((((_1491_v55)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_1491_v55).length))]) ? (_1462_v37) : (_1462_v37)));
        let _rhs240 = !(_1493_v56).equals(_1493_v56);
        let _rhs241 = (_dafny.ZERO).minus(_module.__default.fm0(_module.__default.safeDivisionInt((_this).f24, new BigNumber((_1463_v38).length)), _this.f27, globalState));
        let _rhs242 = _1462_v37;
        let _lhs191 = globalState;
        let _lhs192 = globalState;
        let _lhs193 = _1484_v48;
        let _lhs194 = _1491_v55;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1491_v55).length));
        _lhs191.f14 = _rhs238;
        r0 = _rhs239;
        _lhs192.f2 = _rhs240;
        _lhs193.f33 = _rhs241;
        _lhs194[_lhs195] = _rhs242;
      } else {
        let _1494___mcc_h3 = (_source15).cf56;
        let _1495_cf56 = _1494___mcc_h3;
        let _1496_v57;
        let _nw224 = Array((new BigNumber(3)).toNumber()).fill(false);
        _1496_v57 = _nw224;
        let _index203 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1496_v57).length));
        (_1496_v57)[_index203] = false;
        let _1497_v58;
        _1497_v58 = _dafny.Seq.UnicodeFromString("kltkdmnl");
        let _1498_v59;
        _1498_v59 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1499_v60;
        _1499_v60 = _module.D8.create_DC25(_this.f27, _1462_v37, _1462_v37);
        let _1500_v61;
        _1500_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1498_v59,(_1499_v60).dtor_cf40);
        let _1501_v62;
        _1501_v62 = _dafny.Seq.of(_1462_v37);
        let _index204 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1496_v57).length));
        let _rhs243 = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), function (_1502_i7) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _1497_v58), _1497_v58);
        let _rhs244 = (((_1500_v61).contains(_1498_v59)) ? ((_1500_v61).get(_1498_v59)) : (_1462_v37));
        let _rhs245 = _1501_v62;
        let _lhs196 = _1496_v57;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1496_v57).length));
        let _lhs198 = globalState;
        let _lhs199 = globalState;
        _lhs196[_lhs197] = _rhs243;
        _lhs198.f2 = _rhs244;
        _lhs199.f20 = _rhs245;
        let _index205 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1496_v57).length));
        (_1496_v57)[_index205] = (_1496_v57)[_module.__default.safeIndex(new BigNumber(383), new BigNumber((_1496_v57).length))];
        let _1503_v63;
        _1503_v63 = _dafny.MultiSet.fromElements(_1462_v37);
        let _1504_v64;
        _1504_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v57,new BigNumber((_dafny.Seq.UnicodeFromString("ekhhngqe")).length));
        let _1505_v65;
        _1505_v65 = _dafny.Seq.of(_1496_v57, _1496_v57, _1496_v57, _1496_v57, _1496_v57);
        let _1506_v66;
        let _nw225 = Array((new BigNumber(24)).toNumber());
        _nw225[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-453), _this.f27);
        _nw225[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_this.f27);
        _nw225[(new BigNumber(2)).toNumber()] = _module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_1503_v63).cardinality())), new BigNumber((_1497_v58).length), globalState);
        _nw225[(new BigNumber(3)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(4)).toNumber()] = new BigNumber(685);
        _nw225[(new BigNumber(5)).toNumber()] = new BigNumber((_1504_v64).length);
        _nw225[(new BigNumber(6)).toNumber()] = (new BigNumber((_1497_v58).length)).plus(_this.f27);
        _nw225[(new BigNumber(7)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(8)).toNumber()] = (new BigNumber((_1463_v38).length)).multipliedBy((_this).f24);
        _nw225[(new BigNumber(9)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(10)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(_this.f27, new BigNumber(-1));
        _nw225[(new BigNumber(12)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1497_v58, _1497_v58)).length);
        _nw225[(new BigNumber(14)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(15)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.update(_1505_v65, _module.__default.safeIndex(new BigNumber((_1501_v62).length), new BigNumber((_1505_v65).length)), _1496_v57)).length);
        _nw225[(new BigNumber(17)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(18)).toNumber()] = (_this).f24;
        _nw225[(new BigNumber(19)).toNumber()] = _this.f27;
        _nw225[(new BigNumber(20)).toNumber()] = new BigNumber((_1497_v58).length);
        _nw225[(new BigNumber(21)).toNumber()] = new BigNumber(400);
        _nw225[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_1463_v38).length)).minus((_this).f24));
        _nw225[(new BigNumber(23)).toNumber()] = ((_1462_v37) ? (new BigNumber(662)) : (_module.__default.fm0(new BigNumber(-584), (_this).f24, globalState)));
        _1506_v66 = _nw225;
        let _1507_v67;
        _1507_v67 = _dafny.Seq.of(_1506_v66, _1506_v66, ((_1462_v37) ? (_1506_v66) : (_1506_v66)));
        _1506_v66 = (_1507_v67)[_module.__default.safeIndex((_this).f24, new BigNumber((_1507_v67).length))];
        (globalState).f2 = _1462_v37;
      }
      let _1508_v68;
      _1508_v68 = _dafny.MultiSet.fromElements(!(_1462_v37));
      let _1509_v69;
      _1509_v69 = _dafny.Seq.of(_1462_v37, _1462_v37, _1462_v37);
      let _1510_v70;
      _1510_v70 = _dafny.MultiSet.fromElements((_this).f24, _this.f27, (_this).f24, _this.f27, _this.f27);
      let _1511_v71;
      let _nw226 = new _module.C1();
      _nw226.__ctor(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_1509_v69).length)), _1510_v70), (_this).f24, _this.f27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(382)), function (_1512_i8) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),_this.f27);
      }));
      _1511_v71 = _nw226;
      let _1513_v72;
      _1513_v72 = _dafny.Set.fromElements(_1511_v71, _1511_v71);
      let _1514_v73;
      let _nw227 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _1514_v73 = _nw227;
      let _1515_v75;
      _1515_v75 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1516_v76;
      _1516_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1515_v75,_dafny.Seq.UnicodeFromString("uoejhqk"));
      let _1517_v77;
      let _nw228 = new _module.C4();
      _nw228.__ctor(_1514_v73, _this.f27, _dafny.Seq.of(function () {
        let _coll42 = new _dafny.Map();
        for (const _compr_42 of (_1516_v76).Keys.Elements) {
          let _1518_v74 = _compr_42;
          if ((_1516_v76).contains(_1518_v74)) {
            _coll42.push([_1518_v74,_1511_v71.f33]);
          }
        }
        return _coll42;
      }()));
      _1517_v77 = _nw228;
      let _1519_v78;
      _1519_v78 = _dafny.Set.fromElements(_1517_v77);
      let _1520_v79;
      _1520_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1462_v37);
      let _1521_v80;
      _1521_v80 = _dafny.Seq.UnicodeFromString("uuvexp");
      let _1522_v81;
      let _nw229 = Array((new BigNumber(23)).toNumber());
      _nw229[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f24);
      _nw229[(_dafny.ONE).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(2)).toNumber()] = _this.f27;
      _nw229[(new BigNumber(3)).toNumber()] = _this.f27;
      _nw229[(new BigNumber(4)).toNumber()] = _this.f27;
      _nw229[(new BigNumber(5)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(6)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bobuyv")).length));
      _nw229[(new BigNumber(8)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).f24);
      _nw229[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("k")).length);
      _nw229[(new BigNumber(11)).toNumber()] = new BigNumber(-676);
      _nw229[(new BigNumber(12)).toNumber()] = _module.__default.fm0((_this).f24, _this.f27, globalState);
      _nw229[(new BigNumber(13)).toNumber()] = _module.__default.fm0(new BigNumber((_1513_v72).length), _this.f27, globalState);
      _nw229[(new BigNumber(14)).toNumber()] = _this.f27;
      _nw229[(new BigNumber(15)).toNumber()] = new BigNumber((_1519_v78).length);
      _nw229[(new BigNumber(16)).toNumber()] = new BigNumber((_1520_v79).length);
      _nw229[(new BigNumber(17)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(18)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(19)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(20)).toNumber()] = _1511_v71.f33;
      _nw229[(new BigNumber(21)).toNumber()] = (_this).f24;
      _nw229[(new BigNumber(22)).toNumber()] = new BigNumber((_1521_v80).length);
      _1522_v81 = _nw229;
      let _1523_v82;
      _1523_v82 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_1508_v68).cardinality()), new BigNumber(461)),_1522_v81);
      let _1524_v83;
      let _nw230 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1524_v83 = _nw230;
      let _index206 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_1524_v83).length));
      (_1524_v83)[_index206] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_1525_i9) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      });
      let _index207 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_1524_v83).length));
      let _rhs246 = false;
      let _rhs247 = ((((_this).f26).contains(_1462_v37)) ? (((_this).f26).get(_1462_v37)) : (_dafny.Seq.IsProperPrefixOf(_1509_v69, _1509_v69)));
      let _rhs248 = (_1523_v82).update(new BigNumber((_1509_v69).length), _1514_v73);
      let _rhs249 = _1462_v37;
      let _rhs250 = _1521_v80;
      let _lhs200 = globalState;
      let _lhs201 = _1524_v83;
      let _lhs202 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_1524_v83).length));
      _lhs200.f2 = _rhs246;
      _1462_v37 = _rhs247;
      _1523_v82 = _rhs248;
      r0 = _rhs249;
      _lhs201[_lhs202] = _rhs250;
      _1514_v73 = _1522_v81;
      let _1526_v84;
      _1526_v84 = _dafny.Map.Empty.slice().updateUnsafe(false,_1511_v71.f33);
      let _1527_v85;
      _1527_v85 = _module.D4.create_DC13(_this.f27, _1526_v84, _1521_v80);
      let _1528_v86;
      _1528_v86 = _module.D4.create_DC14(_1527_v85);
      let _1529_v87;
      _1529_v87 = _module.D4.create_DC14(_1528_v86);
      let _1530_v88;
      _1530_v88 = _module.D4.create_DC14(_1528_v86);
      let _1531_v89;
      _1531_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v88,_1521_v80);
      let _1532_v90;
      _1532_v90 = _module.D7.create_DC20(_1531_v89);
      let _1533_v91;
      _1533_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1532_v90,(_1524_v83)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_1524_v83).length))]);
      _1533_v91 = (_1533_v91).update(_1532_v90, _dafny.Seq.UnicodeFromString("aohh"));
      let _1534_v92;
      let _nw231 = Array((new BigNumber(5)).toNumber());
      _1534_v92 = _nw231;
      let _1535_v93;
      let _nw232 = new _module.C4();
      _nw232.__ctor(_1522_v81, (_1517_v77).f24, (_this).f25);
      _1535_v93 = _nw232;
      let _1536_v94;
      _1536_v94 = _dafny.Seq.of(_1535_v93);
      let _1537_v95;
      _1537_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1462_v37,_1509_v69);
      let _index208 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1534_v92).length));
      (_1534_v92)[_index208] = (_1536_v94)[_module.__default.safeIndex(new BigNumber(((((_1537_v95).contains(_1462_v37)) ? ((_1537_v95).get(_1462_v37)) : (_1509_v69))).length), new BigNumber((_1536_v94).length))];
      let _index209 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1534_v92).length));
      (_1534_v92)[_index209] = (_1536_v94)[_module.__default.safeIndex(((_this).f24).multipliedBy(new BigNumber(492)), new BigNumber((_1536_v94).length))];
      r0 = ((_1511_v71).fm10(globalState)) || (!(_1462_v37));
      return r0;
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let _1538_v0;
      _1538_v0 = true;
      if (!(_1538_v0)) {
        (globalState).f9 = (_this).f24;
        let _1539_v1;
        _1539_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v0,new BigNumber(((_this).f26).length));
        let _1540_v2;
        _1540_v2 = _module.D4.create_DC13((_this).f24, _1539_v1, _dafny.Seq.UnicodeFromString("qiphwxsf"));
        if (((_1540_v2).dtor_cf22).isEqualTo(_this.f27)) {
          let _1541_v3;
          let _nw233 = Array((new BigNumber(28)).toNumber()).fill(false);
          _1541_v3 = _nw233;
          let _index210 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1541_v3).length));
          (_1541_v3)[_index210] = _1538_v0;
          let _index211 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1541_v3).length));
          (_1541_v3)[_index211] = _1538_v0;
          let _1542_v4;
          _1542_v4 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1543_v5;
          let _nw234 = new _module.C0();
          _nw234.__ctor(_1538_v0);
          _1543_v5 = _nw234;
          let _1544_v6;
          _1544_v6 = _dafny.Set.fromElements(_this.f27);
          let _1545_v7;
          _1545_v7 = _dafny.MultiSet.fromElements(new BigNumber((_1544_v6).length));
          let _1546_v8;
          _1546_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v4,new BigNumber((_1545_v7).cardinality()));
          let _1547_v9;
          let _nw235 = new _module.C2();
          _nw235.__ctor((_1543_v5).f34, (_this).f24, _dafny.Seq.of(_1546_v8));
          _1547_v9 = _nw235;
          let _1548_v10;
          let _nw236 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1548_v10 = _nw236;
          let _1549_v11;
          _1549_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1547_v9,_1548_v10);
          let _rhs251 = (_1549_v11).contains(_1547_v9);
          let _rhs252 = _module.__default.fm3(globalState);
          let _rhs253 = _1543_v5;
          let _lhs203 = globalState;
          _lhs203.f2 = _rhs251;
          _1542_v4 = _rhs252;
          _1543_v5 = _rhs253;
          (globalState).f14 = (_1547_v9).f24;
          let _1550_v12;
          _1550_v12 = _dafny.MultiSet.fromElements((_1541_v3)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1541_v3).length))]);
          let _1551_v13;
          _1551_v13 = _dafny.MultiSet.fromElements((_1550_v12).Difference(_1550_v12));
          _1551_v13 = (_1551_v13).update(_1550_v12, _module.__default.abs((_this).f24));
          let _1552_v14;
          let _nw237 = Array((new BigNumber(9)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = _1544_v6;
          _nw237[(_dafny.ONE).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(2)).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements((_1547_v9).f24);
          _nw237[(new BigNumber(4)).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(5)).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(6)).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(7)).toNumber()] = _1544_v6;
          _nw237[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements((_this).f24, (_this).f24);
          _1552_v14 = _nw237;
          let _1553_v15;
          _1553_v15 = _dafny.Seq.of(_1552_v14);
          let _1554_v16;
          _1554_v16 = _dafny.Seq.of((_1543_v5).f34, _1538_v0);
          let _1555_v17;
          let _nw238 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _1555_v17 = _nw238;
          let _1556_v18;
          let _nw239 = new _module.C3();
          _nw239.__ctor((_1553_v15)[_module.__default.safeIndex(new BigNumber((_1554_v16).length), new BigNumber((_1553_v15).length))], _1555_v17, (_this).f24, (_1547_v9).f25);
          _1556_v18 = _nw239;
        } else {
          let _1557_v19;
          let _nw240 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
          _1557_v19 = _nw240;
          let _1558_v21;
          let _init36 = function (_1559_i0) {
            return (_module.D11.create_DC30(function () {
  let _coll43 = new _dafny.Map();
  for (const _compr_43 of _dafny.IntegerRange(new BigNumber(429), new BigNumber(778))) {
    let _1560_v20 = _compr_43;
    if (((new BigNumber(429)).isLessThanOrEqualTo(_1560_v20)) && ((_1560_v20).isLessThan(new BigNumber(778)))) {
      _coll43.push([(_1560_v20).plus((_this).f24),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bgqm")).length))]);
    }
  }
  return _coll43;
}())).dtor_cf47;
          };
          let _nw241 = Array((new BigNumber(13)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw241.length); _i0_36++) {
            _nw241[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1558_v21 = _nw241;
          let _1561_v22;
          let _nw242 = new _module.C3();
          _nw242.__ctor(_1557_v19, _1558_v21, (_this).f24, (_this).f25);
          _1561_v22 = _nw242;
          let _1562_v23;
          _1562_v23 = _1561_v22;
          _1562_v23 = _1562_v23;
          let _1563_v24;
          _1563_v24 = _dafny.MultiSet.fromElements(_1561_v22, _1561_v22);
          let _1564_v25;
          _1564_v25 = _dafny.Set.fromElements((_1563_v24).IsDisjointFrom(_1563_v24));
          let _1565_v26;
          _1565_v26 = _module.D0.create_DC1((_this).f24, new BigNumber(279), _1538_v0);
          let _1566_v27;
          _1566_v27 = new _dafny.CodePoint('n'.codePointAt(0));
          let _pat_let_tv25 = _1538_v0;
          let _rhs254 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_1561_v22).f24, (_1561_v22).f24));
          let _rhs255 = _dafny.Set.fromElements((function (_pat_let29_0) {
            return function (_1567_dt__update__tmp_h0) {
              return function (_pat_let30_0) {
                return function (_1568_dt__update_hcf1_h0) {
                  return function (_pat_let31_0) {
                    return function (_1569_dt__update_hcf3_h0) {
                      return _module.D0.create_DC1(_1568_dt__update_hcf1_h0, (_1567_dt__update__tmp_h0).dtor_cf2, _1569_dt__update_hcf3_h0);
                    }(_pat_let31_0);
                  }(_pat_let_tv25);
                }(_pat_let30_0);
              }(_this.f27);
            }(_pat_let29_0);
          }(_1565_v26)).dtor_cf3, _module.__default.fm2(_this.f27, _module.__default.fm0(new BigNumber(-397), (_dafny.ZERO).minus((_1561_v22).f24), globalState), _1566_v27, globalState));
          let _rhs256 = (_this).f24;
          let _lhs204 = globalState;
          let _lhs205 = globalState;
          _lhs204.f14 = _rhs254;
          _1564_v25 = _rhs255;
          _lhs205.f18 = _rhs256;
          (globalState).f2 = ((_1538_v0) ? (_1538_v0) : (!(_1538_v0)));
          let _1570_v28;
          let _nw243 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1570_v28 = _nw243;
          _1570_v28 = _1570_v28;
          let _1571_v29;
          _1571_v29 = _dafny.MultiSet.fromElements(_1538_v0, _1538_v0);
          let _1572_v30;
          _1572_v30 = _dafny.Seq.of(_this.f27);
          let _1573_v31;
          _1573_v31 = _dafny.Set.fromElements((_this).f24, new BigNumber((_1572_v30).length), (_this).f24, _this.f27, _this.f27);
          let _1574_v32;
          _1574_v32 = _dafny.Set.fromElements(_1566_v27, _1566_v27);
          let _1575_v33;
          let _nw244 = Array((new BigNumber(29)).toNumber());
          _nw244[(_dafny.ZERO).toNumber()] = !(_1538_v0) || (_1538_v0);
          _nw244[(_dafny.ONE).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(2)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(3)).toNumber()] = ((_1561_v22).f24).isLessThanOrEqualTo((_this).f24);
          _nw244[(new BigNumber(4)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(5)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(6)).toNumber()] = false;
          _nw244[(new BigNumber(7)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(8)).toNumber()] = false;
          _nw244[(new BigNumber(9)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(10)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(11)).toNumber()] = (_this.f27).isLessThanOrEqualTo((_1561_v22).f24);
          _nw244[(new BigNumber(12)).toNumber()] = (_1565_v26).dtor_cf3;
          _nw244[(new BigNumber(13)).toNumber()] = true;
          _nw244[(new BigNumber(14)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(15)).toNumber()] = (_this.f27).isEqualTo(new BigNumber((_1571_v29).cardinality()));
          _nw244[(new BigNumber(16)).toNumber()] = ((_1561_v22).f24).isLessThan((_1561_v22).f24);
          _nw244[(new BigNumber(17)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(18)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(19)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(20)).toNumber()] = (_1573_v31).IsSubsetOf(_1573_v31);
          _nw244[(new BigNumber(21)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(22)).toNumber()] = _1538_v0;
          _nw244[(new BigNumber(23)).toNumber()] = !(((_this).f24).isLessThan(new BigNumber((_1574_v32).length)));
          _nw244[(new BigNumber(24)).toNumber()] = !(_1538_v0);
          _nw244[(new BigNumber(25)).toNumber()] = (_1573_v31).IsProperSubsetOf(_dafny.Set.fromElements(_this.f27, _this.f27));
          _nw244[(new BigNumber(26)).toNumber()] = ((_1538_v0) ? (_1538_v0) : (_1538_v0));
          _nw244[(new BigNumber(27)).toNumber()] = ((_this).f24).isEqualTo(_this.f27);
          _nw244[(new BigNumber(28)).toNumber()] = _1538_v0;
          _1575_v33 = _nw244;
          let _index212 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1575_v33).length));
          (_1575_v33)[_index212] = _1538_v0;
          let _1576_v34;
          _1576_v34 = _dafny.Seq.of(_1566_v27);
          let _1577_v35;
          _1577_v35 = _dafny.MultiSet.fromElements((_this).f24, new BigNumber((_1576_v34).length));
          let _1578_v36;
          let _nw245 = new _module.C0();
          _nw245.__ctor(_1538_v0);
          _1578_v36 = _nw245;
          let _1579_v37;
          _1579_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1578_v36,_dafny.MultiSet.FromArray(_1572_v30));
          let _index213 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1575_v33).length));
          (_1575_v33)[_index213] = ((_1577_v35).Intersect((((_1579_v37).contains(_1578_v36)) ? ((_1579_v37).get(_1578_v36)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f24)))))).IsSubsetOf(_1577_v35);
        }
        (globalState).f18 = _module.__default.safeDivisionInt((_this).f24, new BigNumber(798));
        let _1580_v38;
        _1580_v38 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1581_v39;
        _1581_v39 = _dafny.Seq.UnicodeFromString("dhvqnlf");
        _1538_v0 = ((_1538_v0) ? (_dafny.Seq.contains(_1581_v39, _1580_v38)) : (!(((_1538_v0) ? (_1538_v0) : (_1538_v0)))));
        (globalState).f17 = _1581_v39;
      } else {
        (globalState).f1 = (_dafny.ZERO).minus((_this).f24);
        let _1582_v40;
        _1582_v40 = new _dafny.CodePoint('t'.codePointAt(0));
        if (_module.__default.fm2((_this).f24, (_this).f24, _1582_v40, globalState)) {
          (globalState).f14 = (_this).f24;
          let _1583_v41;
          _1583_v41 = _module.D12.create_DC33(_1582_v40, _1538_v0);
          let _1584_v42;
          _1584_v42 = _module.D12.create_DC34(((!(_1538_v0)) ? (_1583_v41) : (_1583_v41)));
          let _1585_v43;
          let _nw246 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1585_v43 = _nw246;
          let _1586_v44;
          _1586_v44 = _dafny.Seq.of(_1585_v43);
          let _1587_v45;
          let _nw247 = new _module.C4();
          _nw247.__ctor((_1586_v44)[_module.__default.safeIndex((_this).f24, new BigNumber((_1586_v44).length))], new BigNumber(191), _dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), function (_1588_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_this.f27);
          }));
          _1587_v45 = _nw247;
          let _1589_v46;
          _1589_v46 = _dafny.Seq.of(_1538_v0, _1538_v0, _1538_v0);
          let _1590_v47;
          _1590_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(972)), ((_1591_v45, _1592_v0) => function (_1593_i2) {
            return (_module.D8.create_DC25((_1591_v45).f24, _1592_v0, _1592_v0)).dtor_cf38;
          })(_1587_v45, _1538_v0)),_1587_v45);
          let _1594_v48;
          _1594_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_1587_v45).f24, (_this).f24)).length),(_1587_v45).f24);
          let _1595_v49;
          _1595_v49 = _dafny.Seq.of(new BigNumber((_1594_v48).length), _module.__default.fm0((_1587_v45).f24, _this.f27, globalState));
          let _rhs257 = _1584_v42;
          let _rhs258 = _dafny.Seq.IsPrefixOf(_1589_v46, _1589_v46);
          let _rhs259 = (((_1590_v47).contains(_1595_v49)) ? ((_1590_v47).get(_1595_v49)) : (_1587_v45));
          _1584_v42 = _rhs257;
          _1538_v0 = _rhs258;
          _1587_v45 = _rhs259;
          let _1596_v50;
          let _nw248 = new _module.C4();
          _nw248.__ctor(_1585_v43, (_1587_v45).f24, (_1587_v45).f25);
          _1596_v50 = _nw248;
          let _1597_v51;
          _1597_v51 = _dafny.Map.Empty.slice().updateUnsafe((_1587_v45).f24,_1596_v50);
          let _1598_v52;
          _1598_v52 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_1597_v51).length));
          let _1599_v53;
          _1599_v53 = _dafny.MultiSet.fromElements(_1538_v0);
          let _1600_v54;
          _1600_v54 = _dafny.Seq.of((_1598_v52).update(_1538_v0, new BigNumber((_1599_v53).cardinality())), _1598_v52);
          _1600_v54 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1595_v49).length))).update(_1538_v0, _this.f27), _1598_v52);
          r0 = _1538_v0;
          (globalState).f2 = (_1538_v0) === (_1538_v0);
        } else {
          let _1601_v55;
          _1601_v55 = _dafny.Seq.UnicodeFromString("drbtc");
          let _1602_v56;
          _1602_v56 = _dafny.Set.fromElements(new BigNumber((_1601_v55).length));
          let _1603_v57;
          _1603_v57 = _dafny.Seq.of(_1538_v0, _1538_v0);
          let _1604_v58;
          _1604_v58 = _dafny.MultiSet.fromElements((_this).f24, new BigNumber((_1601_v55).length));
          let _1605_v60;
          _1605_v60 = _module.D6.create_DC16(_1602_v56);
          let _1606_v61;
          let _nw249 = Array((new BigNumber(28)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1602_v56;
          _nw249[(_dafny.ONE).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(2)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_this.f27, _this.f27, new BigNumber((_1603_v57).length));
          _nw249[(new BigNumber(4)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(5)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(6)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(7)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(8)).toNumber()] = function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_1604_v58).Elements) {
              let _1607_v59 = _compr_44;
              if ((_1604_v58).contains(_1607_v59)) {
                _coll44.add((_1607_v59).plus(new BigNumber((_dafny.Seq.UnicodeFromString("omdiun")).length)));
              }
            }
            return _coll44;
          }();
          _nw249[(new BigNumber(9)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(10)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(11)).toNumber()] = (_1605_v60).dtor_cf27;
          _nw249[(new BigNumber(12)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(13)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(14)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(15)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(16)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(17)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(18)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(19)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(20)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(new BigNumber(510), new BigNumber((_1602_v56).length));
          _nw249[(new BigNumber(22)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(23)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(24)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(25)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(26)).toNumber()] = _1602_v56;
          _nw249[(new BigNumber(27)).toNumber()] = _1602_v56;
          _1606_v61 = _nw249;
          let _1608_v62;
          let _nw250 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
          _1608_v62 = _nw250;
          let _1609_v63;
          _1609_v63 = (_this).f25;
          let _1610_v64;
          let _nw251 = new _module.C3();
          _nw251.__ctor(_1606_v61, _1608_v62, (_dafny.ZERO).minus((_this).f24), (_1609_v63));
          _1610_v64 = _nw251;
          let _1611_v65;
          _1611_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v0,_1610_v64);
          _1611_v65 = (_1611_v65).update(false, _1610_v64);
          let _1612_v66;
          _1612_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v0,_this.f27);
          let _1613_v67;
          _1613_v67 = _dafny.Map.Empty.slice().updateUnsafe(!(_1538_v0),_1612_v66);
          _1613_v67 = ((_1613_v67).Merge(_1613_v67)).Merge(_1613_v67);
          _1612_v66 = _module.__default.fm17(_module.__default.fm0(new BigNumber((_1601_v55).length), new BigNumber(-163), globalState), _this.f27, _this.f27, _1538_v0, globalState);
          let _1614_v68;
          let _nw252 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1614_v68 = _nw252;
          let _index214 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_1614_v68).length));
          (_1614_v68)[_index214] = _1582_v40;
          let _index215 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_1614_v68).length));
          (_1614_v68)[_index215] = _1582_v40;
          let _1615_v69;
          _1615_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_1603_v57).length)),(_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)), new BigNumber((_1601_v55).length))));
          let _1616_v70;
          _1616_v70 = _dafny.Seq.of(_this.f27);
          let _1617_v71;
          _1617_v71 = _dafny.MultiSet.fromElements(_1538_v0);
          _1615_v69 = (_1615_v69).update(_dafny.Seq.Concat(_1616_v70, _1616_v70), (_this.f27).minus(new BigNumber((_1617_v71).cardinality())));
        }
        let _1618_v72;
        _1618_v72 = _dafny.Set.fromElements(_this.f27);
        let _1619_v73;
        _1619_v73 = _dafny.Seq.of((_this).f24, new BigNumber(-217));
        let _1620_v75;
        let _nw253 = Array((new BigNumber(14)).toNumber());
        _nw253[(_dafny.ZERO).toNumber()] = (_1618_v72).Difference(_1618_v72);
        _nw253[(_dafny.ONE).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_this.f27, (_this).f24, (_this).f24);
        _nw253[(new BigNumber(3)).toNumber()] = (_dafny.Set.fromElements(_this.f27, _this.f27)).Union(_1618_v72);
        _nw253[(new BigNumber(4)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(5)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(6)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(7)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(8)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(9)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_this.f27, _this.f27, (_this).f24, new BigNumber((_1619_v73).length));
        _nw253[(new BigNumber(11)).toNumber()] = function () {
          let _coll45 = new _dafny.Set();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(234), new BigNumber(-416))) {
            let _1621_v74 = _compr_45;
            if (((new BigNumber(234)).isLessThanOrEqualTo(_1621_v74)) && ((_1621_v74).isLessThan(new BigNumber(-416)))) {
              _coll45.add(_module.__default.safeModuloInt(_1621_v74, (_this).f24));
            }
          }
          return _coll45;
        }();
        _nw253[(new BigNumber(12)).toNumber()] = _1618_v72;
        _nw253[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements((_this).f24, (_this).f24);
        _1620_v75 = _nw253;
        _1620_v75 = _1620_v75;
        let _1622_v76;
        let _nw254 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _1622_v76 = _nw254;
        let _1623_v77;
        let _nw255 = Array((new BigNumber(20)).toNumber()).fill(false);
        _1623_v77 = _nw255;
        let _1624_v78;
        _1624_v78 = _dafny.Map.Empty.slice().updateUnsafe(false,_1623_v77);
        let _index216 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1622_v76).length));
        (_1622_v76)[_index216] = (_1624_v78).Merge(_1624_v78);
        let _1625_v79;
        _1625_v79 = _1624_v78;
        let _index217 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1622_v76).length));
        (_1622_v76)[_index217] = (_1625_v79);
        (globalState).f18 = new BigNumber(772);
      }
      (globalState).f1 = new BigNumber((_module.__default.fm40(globalState)).length);
      (_this).m3((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f24, (_this).f24)), globalState);
      let _1626_v80;
      _1626_v80 = _dafny.MultiSet.fromElements(_this.f27);
      if ((_1626_v80).IsProperSubsetOf(_1626_v80)) {
        let _1627_v81;
        _1627_v81 = _dafny.Seq.of(_1538_v0);
        if (((_1538_v0) ? (!(_1538_v0) || (_1538_v0)) : ((_1627_v81)[_module.__default.safeIndex(new BigNumber(-585), new BigNumber((_1627_v81).length))]))) {
          (globalState).f18 = _this.f27;
          r0 = !(((!(_this.f27).isEqualTo((_this).f24)) ? (_1538_v0) : (!(_1538_v0))));
          let _1628_v82;
          _1628_v82 = _dafny.Seq.UnicodeFromString("lofb");
          _1538_v0 = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("th"), _1628_v82);
          let _1629_v83;
          _1629_v83 = _dafny.Set.fromElements(_1538_v0);
          _1629_v83 = (_1629_v83).Union((_1629_v83).Intersect(_1629_v83));
          (globalState).f17 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1628_v82, _1628_v82), _dafny.Seq.UnicodeFromString("fyxkjjg"));
        } else {
          (globalState).f2 = ((_1538_v0) ? (((((_this).f26).contains(_1538_v0)) ? (((_this).f26).get(_1538_v0)) : (_1538_v0))) : (_1538_v0));
          let _1630_v84;
          _1630_v84 = _module.D0.create_DC1(_this.f27, _this.f27, _1538_v0);
          let _1631_v85;
          _1631_v85 = _dafny.Seq.UnicodeFromString("qljejp");
          let _1632_v86;
          _1632_v86 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1631_v85).length)));
          _1630_v84 = _module.D0.create_DC1((_this).f24, ((_this).f24).multipliedBy(new BigNumber(632)), (_1632_v86).IsSubsetOf(_1632_v86));
          let _1633_v87;
          let _init37 = ((_1634_v0) => function (_1635_i3) {
            return _1634_v0;
          })(_1538_v0);
          let _nw256 = Array((new BigNumber(18)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw256.length); _i0_37++) {
            _nw256[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1633_v87 = _nw256;
          let _1636_v88;
          _1636_v88 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,_1538_v0);
          let _1637_v89;
          _1637_v89 = _dafny.Map.Empty.slice().updateUnsafe((((_1636_v88).contains(new BigNumber(918))) ? ((_1636_v88).get(new BigNumber(918))) : (_1538_v0)),(_this).f24);
          let _index218 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length));
          (_1633_v87)[_index218] = !(!(_1637_v89).equals(_dafny.Map.Empty.slice().updateUnsafe(_1538_v0,_this.f27)));
          let _1638_v90;
          _1638_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1631_v85,new BigNumber(327));
          let _index219 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length));
          let _rhs260 = _1538_v0;
          let _rhs261 = (_dafny.Map.Empty.slice().updateUnsafe(_1631_v85,_module.__default.fm0(_this.f27, _this.f27, globalState))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1631_v85,(_this).f24)).Merge(function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), ((_1639_v85) => function (_1640_i4) {
              return _1639_v85;
            })(_1631_v85))).Elements) {
              let _1641_v91 = _compr_46;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), ((_1642_v85) => function (_1640_i4) {
                return _1642_v85;
              })(_1631_v85)), _1641_v91)) {
                _coll46.push([_1641_v91,_this.f27]);
              }
            }
            return _coll46;
          }()));
          let _lhs206 = _1633_v87;
          let _lhs207 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length));
          _lhs206[_lhs207] = _rhs260;
          _1638_v90 = _rhs261;
          (globalState).f9 = (_this).f24;
          let _1643_v92;
          let _nw257 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _1643_v92 = _nw257;
          let _1644_v93;
          _1644_v93 = _dafny.MultiSet.fromElements((_1633_v87)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length))]);
          let _1645_v94;
          _1645_v94 = _dafny.Seq.of(new BigNumber((_1637_v89).length));
          let _1646_v95;
          _1646_v95 = _module.D3.create_DC11(_this.f27, (_1627_v81)[_module.__default.safeIndex(new BigNumber((_1645_v94).length), new BigNumber((_1627_v81).length))], true);
          let _index220 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length));
          let _rhs262 = (_1626_v80).Union(_1626_v80);
          let _rhs263 = _this.f27;
          let _rhs264 = (((_1633_v87)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length))]) ? ((new BigNumber((_1644_v93).cardinality())).isLessThan(_this.f27)) : ((_1646_v95).dtor_cf20));
          let _rhs265 = ((!((_1633_v87)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length))])) ? (_1643_v92) : (_1643_v92));
          let _lhs208 = globalState;
          let _lhs209 = _1633_v87;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1633_v87).length));
          _1626_v80 = _rhs262;
          _lhs208.f9 = _rhs263;
          _lhs209[_lhs210] = _rhs264;
          _1643_v92 = _rhs265;
        }
        let _1647_v96;
        let _nw258 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1647_v96 = _nw258;
        let _1648_v97;
        _1648_v97 = _module.D3.create_DC11(_this.f27, _1538_v0, _1538_v0);
        let _1649_v98;
        let _nw259 = Array((new BigNumber(27)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = _1647_v96;
        _nw259[(_dafny.ONE).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(2)).toNumber()] = ((_1538_v0) ? (_1647_v96) : (_1647_v96));
        _nw259[(new BigNumber(3)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(4)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(5)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(6)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(7)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(8)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(9)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(10)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(11)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(12)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(13)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(14)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(15)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(16)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(17)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(18)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(19)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(20)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(21)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(22)).toNumber()] = (((_1648_v97).dtor_cf19) ? (_1647_v96) : (_1647_v96));
        _nw259[(new BigNumber(23)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(24)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(25)).toNumber()] = _1647_v96;
        _nw259[(new BigNumber(26)).toNumber()] = _1647_v96;
        _1649_v98 = _nw259;
        let _index221 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length));
        (_1649_v98)[_index221] = _1647_v96;
        let _index222 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length));
        (_1649_v98)[_index222] = _1647_v96;
        let _1650_v99;
        _1650_v99 = _dafny.MultiSet.fromElements(_1538_v0, _1538_v0, _1538_v0, _1538_v0, _1538_v0);
        let _1651_v100;
        _1651_v100 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1652_v101;
        _1652_v101 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f24, (_this).f24, _1651_v100, globalState),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_this.f27)).length));
        if ((_dafny.MultiSet.fromElements(!(_1538_v0), false)).IsSubsetOf((_1650_v99).update(_1538_v0, _module.__default.abs(new BigNumber((_1652_v101).length))))) {
          r0 = (_1626_v80).IsDisjointFrom(_1626_v80);
          let _1653_v102;
          _1653_v102 = _dafny.Set.fromElements((_this).f24);
          let _1654_v103;
          _1654_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1653_v102,(_this).f24);
          _1654_v103 = (_1654_v103).update(_module.__default.fm18(globalState), _module.__default.safeModuloInt(new BigNumber((_1627_v81).length), _this.f27));
          let _arr0 = (_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))];
          let _index223 = _module.__default.safeIndex(new BigNumber(777), new BigNumber(((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))]).length));
          _arr0[_index223] = !(_1538_v0);
          let _arr1 = (_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))];
          let _index224 = _module.__default.safeIndex(new BigNumber(777), new BigNumber(((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))]).length));
          _arr1[_index224] = (_module.__default.fm0((_this).f24, (_this).f24, globalState)).isLessThanOrEqualTo(_this.f27);
          (globalState).f2 = _1538_v0;
          let _1655_v104;
          _1655_v104 = _dafny.Seq.of(_1651_v100, _1651_v100, _1651_v100);
          let _1656_v105;
          _1656_v105 = _module.D1.create_DC3(_1655_v104);
          let _1657_v106;
          _1657_v106 = _dafny.Set.fromElements(_1626_v80);
          let _1658_v107;
          _1658_v107 = _dafny.Seq.of(_dafny.Set.fromElements(_1626_v80, _1626_v80, _1626_v80));
          let _1659_v109;
          _1659_v109 = _dafny.Seq.of(_1626_v80, _1626_v80);
          let _1660_v110;
          let _nw260 = new _module.C1();
          _nw260.__ctor(_1659_v109, (_this).f24, (_this).f24, (_this).f25);
          _1660_v110 = _nw260;
          let _1661_v111;
          _1661_v111 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v110,((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))])[_module.__default.safeIndex(new BigNumber(777), new BigNumber(((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))]).length))]);
          let _1662_v112;
          _1662_v112 = _dafny.Seq.of((((_1626_v80).contains((_this).f24)) ? ((_1626_v80).get((_this).f24)) : (new BigNumber((_1661_v111).length))));
          let _1663_v113;
          _1663_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v110.f33,_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_1662_v112), _1626_v80));
          let _1664_v114;
          _1664_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1626_v80,_this.f27);
          let _1665_v116;
          let _nw261 = Array((new BigNumber(27)).toNumber());
          _nw261[(_dafny.ZERO).toNumber()] = _1657_v106;
          _nw261[(_dafny.ONE).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(2)).toNumber()] = (_1657_v106).Difference(_1657_v106);
          _nw261[(new BigNumber(3)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(4)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(5)).toNumber()] = (_1657_v106).Difference(_1657_v106);
          _nw261[(new BigNumber(6)).toNumber()] = ((_1658_v107)[_module.__default.safeIndex(new BigNumber((_1655_v104).length), new BigNumber((_1658_v107).length))]).Difference(_1657_v106);
          _nw261[(new BigNumber(7)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(8)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(9)).toNumber()] = ((_1538_v0) ? (_1657_v106) : (_1657_v106));
          _nw261[(new BigNumber(10)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(11)).toNumber()] = (_1657_v106).Intersect(_module.__default.fm41((_this).f24, _module.__default.fm4(_1538_v0, (_this).f24, _1651_v100, !(_1538_v0), globalState), !(!(false)), globalState));
          _nw261[(new BigNumber(12)).toNumber()] = function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of (_1657_v106).Elements) {
              let _1666_v108 = _compr_47;
              if ((_1657_v106).contains(_1666_v108)) {
                _coll47.add(_1666_v108);
              }
            }
            return _coll47;
          }();
          _nw261[(new BigNumber(13)).toNumber()] = ((_1658_v107)[_module.__default.safeIndex(new BigNumber((_1662_v112).length), new BigNumber((_1658_v107).length))]).Difference((((_1663_v113).contains((_dafny.ZERO).minus(_1660_v110.f33))) ? ((_1663_v113).get((_dafny.ZERO).minus(_1660_v110.f33))) : (_1657_v106)));
          _nw261[(new BigNumber(14)).toNumber()] = ((!(_1538_v0)) ? (_dafny.Set.fromElements(_1626_v80)) : (function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_1664_v114).Keys.Elements) {
              let _1667_v115 = _compr_48;
              if ((_1664_v114).contains(_1667_v115)) {
                _coll48.add(_1667_v115);
              }
            }
            return _coll48;
          }()));
          _nw261[(new BigNumber(15)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(_1626_v80, _dafny.MultiSet.fromElements((_this).f24, _1660_v110.f33));
          _nw261[(new BigNumber(17)).toNumber()] = (_1657_v106).Difference(_1657_v106);
          _nw261[(new BigNumber(18)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(19)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(20)).toNumber()] = (_1658_v107)[_module.__default.safeIndex(_this.f27, new BigNumber((_1658_v107).length))];
          _nw261[(new BigNumber(21)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(22)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(23)).toNumber()] = _dafny.Set.fromElements(_1626_v80);
          _nw261[(new BigNumber(24)).toNumber()] = _module.__default.fm41(_this.f27, _1655_v104, true, globalState);
          _nw261[(new BigNumber(25)).toNumber()] = _1657_v106;
          _nw261[(new BigNumber(26)).toNumber()] = (_1657_v106).Difference(_1657_v106);
          _1665_v116 = _nw261;
          let _1668_v117;
          _1668_v117 = _module.D8.create_DC25(_this.f27, ((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))])[_module.__default.safeIndex(new BigNumber(777), new BigNumber(((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))]).length))], _1538_v0);
          let _index225 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1665_v116).length));
          (_1665_v116)[_index225] = _module.__default.fm41((_this).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), ((_1669_v100) => function (_1670_i5) {
            return _1669_v100;
          })(_1651_v100)), (_1668_v117).dtor_cf39, globalState);
          let _1671_v118;
          let _nw262 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1671_v118 = _nw262;
          let _1672_v119;
          let _nw263 = Array((new BigNumber(4)).toNumber());
          _nw263[(_dafny.ZERO).toNumber()] = _1671_v118;
          _nw263[(_dafny.ONE).toNumber()] = _1671_v118;
          _nw263[(new BigNumber(2)).toNumber()] = _1671_v118;
          _nw263[(new BigNumber(3)).toNumber()] = _1671_v118;
          _1672_v119 = _nw263;
          let _index226 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1672_v119).length));
          (_1672_v119)[_index226] = _1671_v118;
          let _pat_let_tv26 = _1655_v104;
          let _index227 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1665_v116).length));
          let _index228 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1672_v119).length));
          let _rhs266 = ((_1538_v0) ? (_1656_v105) : (((((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))])[_module.__default.safeIndex(new BigNumber(777), new BigNumber(((_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))]).length))]) ? (_1656_v105) : (function (_pat_let32_0) {
            return function (_1673_dt__update__tmp_h1) {
              return function (_pat_let33_0) {
                return function (_1674_dt__update_hcf7_h0) {
                  return _module.D1.create_DC3(_1674_dt__update_hcf7_h0);
                }(_pat_let33_0);
              }(_pat_let_tv26);
            }(_pat_let32_0);
          }(_1656_v105)))));
          let _rhs267 = _1657_v106;
          let _rhs268 = _1671_v118;
          let _lhs211 = _1665_v116;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1665_v116).length));
          let _lhs213 = _1672_v119;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1672_v119).length));
          _1656_v105 = _rhs266;
          _lhs211[_lhs212] = _rhs267;
          _lhs213[_lhs214] = _rhs268;
        } else {
          (globalState).f14 = ((_1538_v0) ? ((_this).f24) : (_this.f27));
          let _index229 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length));
          (_1649_v98)[_index229] = (_1649_v98)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1649_v98).length))];
          let _1675_v120;
          let _nw264 = new _module.C2();
          _nw264.__ctor(((_1538_v0) ? (_1538_v0) : (_1538_v0)), (_this).f24, (_this).f25);
          _1675_v120 = _nw264;
          (globalState).f14 = (_dafny.ZERO).minus(_this.f27);
          let _1676_v121;
          _1676_v121 = _dafny.Seq.UnicodeFromString("tlgkqke");
          let _1677_v122;
          let _nw265 = new _module.C2();
          _nw265.__ctor(_1538_v0, (_this.f27).minus(new BigNumber((_1676_v121).length)), (_this).f25);
          _1677_v122 = _nw265;
        }
        let _1678_v123;
        let _init38 = function (_1679_i6) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sihbc"), _dafny.Seq.UnicodeFromString("egjfu"));
        };
        let _nw266 = Array((new BigNumber(13)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw266.length); _i0_38++) {
          _nw266[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1678_v123 = _nw266;
        let _1680_v124;
        _1680_v124 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), function (_1681_i7) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }));
        let _index230 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1678_v123).length));
        (_1678_v123)[_index230] = (_1680_v124)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f27), new BigNumber((_1680_v124).length))];
        let _1682_v125;
        _1682_v125 = _dafny.Seq.UnicodeFromString("wugxlqdqp");
        let _index231 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1678_v123).length));
        (_1678_v123)[_index231] = _1682_v125;
        let _1683_v126;
        _1683_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this.f27).multipliedBy(new BigNumber(350)));
        let _1684_v127;
        _1684_v127 = _dafny.Seq.of(new BigNumber((_1626_v80).cardinality()), _this.f27);
        _1683_v126 = (_1683_v126).update((_this).f24, (_1684_v127)[_module.__default.safeIndex(_this.f27, new BigNumber((_1684_v127).length))]);
      } else {
        let _1685_v128;
        _1685_v128 = _dafny.Seq.of(_this.f27, (_this).f24);
        let _1686_v129;
        let _nw267 = new _module.C1();
        _nw267.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-395)), ((_1687_v0) => function (_1688_i8) {
          return _dafny.MultiSet.fromElements(_this.f27, new BigNumber((_dafny.MultiSet.fromElements(_1687_v0, (_module.D10.create_DC29((_this).f24, _1687_v0)).dtor_cf46, _1687_v0)).cardinality()));
        })(_1538_v0)), new BigNumber((_dafny.Seq.update(_1685_v128, _module.__default.safeIndex(new BigNumber((_1626_v80).cardinality()), new BigNumber((_1685_v128).length)), _this.f27)).length), _this.f27, (_this).f25);
        _1686_v129 = _nw267;
        let _1689_v130;
        _1689_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v129,_1538_v0);
        r0 = (((((_1689_v130).contains(_1686_v129)) ? ((_1689_v130).get(_1686_v129)) : (_1538_v0))) ? (true) : (_1538_v0));
        let _1690_v131;
        _1690_v131 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v0,new BigNumber(80));
        let _1691_v132;
        _1691_v132 = _dafny.Seq.of(_1538_v0, true);
        let _1692_v133;
        let _nw268 = Array((new BigNumber(27)).toNumber());
        _nw268[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw268[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_1686_v129.f33, new BigNumber((_1690_v131).length));
        _nw268[(new BigNumber(2)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw268[(new BigNumber(4)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(5)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(6)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(7)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw268[(new BigNumber(9)).toNumber()] = (((_1690_v131).contains(_1538_v0)) ? ((_1690_v131).get(_1538_v0)) : ((_this).f24));
        _nw268[(new BigNumber(10)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(11)).toNumber()] = (_this).f24;
        _nw268[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1691_v132).length), _1686_v129.f33);
        _nw268[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt((_this).f24, _this.f27);
        _nw268[(new BigNumber(14)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(15)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_this.f27, _this.f27);
        _nw268[(new BigNumber(17)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(18)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(19)).toNumber()] = new BigNumber(676);
        _nw268[(new BigNumber(20)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(21)).toNumber()] = _1686_v129.f33;
        _nw268[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(356)).plus((_this).f24));
        _nw268[(new BigNumber(23)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(24)).toNumber()] = (_this).f24;
        _nw268[(new BigNumber(25)).toNumber()] = _this.f27;
        _nw268[(new BigNumber(26)).toNumber()] = _this.f27;
        _1692_v133 = _nw268;
        let _index232 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1692_v133).length));
        (_1692_v133)[_index232] = _1686_v129.f33;
        let _index233 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1692_v133).length));
        (_1692_v133)[_index233] = _module.__default.safeDivisionInt((_this).f24, new BigNumber(574));
        let _1693_v134;
        _1693_v134 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1694_v135;
        _1694_v135 = _dafny.Map.Empty.slice().updateUnsafe((((_1626_v80).contains(_1686_v129.f33)) ? ((_1626_v80).get(_1686_v129.f33)) : (_1686_v129.f33)),_1693_v134);
        let _1695_v137;
        _1695_v137 = _dafny.Seq.of(_1693_v134, new _dafny.CodePoint('f'.codePointAt(0)), _1693_v134, _1693_v134, new _dafny.CodePoint('q'.codePointAt(0)));
        let _1696_v138;
        _1696_v138 = _dafny.Seq.of(function () {
          let _coll49 = new _dafny.Map();
          for (const _compr_49 of (_1695_v137).Elements) {
            let _1697_v136 = _compr_49;
            if (_dafny.Seq.contains(_1695_v137, _1697_v136)) {
              _coll49.push([_1697_v136,true]);
            }
          }
          return _coll49;
        }());
        let _1698_v139;
        let _nw269 = new _module.C1();
        _nw269.__ctor((_1686_v129).f32, new BigNumber(207), new BigNumber((_1696_v138).length), (_this).f25);
        _1698_v139 = _nw269;
        let _1699_v140;
        _1699_v140 = _dafny.Seq.of(_1698_v139);
        let _1700_v141;
        let _nw270 = new _module.C2();
        _nw270.__ctor((new BigNumber((_dafny.Seq.of((_this).f24, new BigNumber((_1694_v135).length), new BigNumber(797), _this.f27, new BigNumber((_1699_v140).length))).length)).isLessThanOrEqualTo(_1686_v129.f33), (_this).f24, (_this).f25);
        _1700_v141 = _nw270;
        let _1701_v142;
        _1701_v142 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(420),_1686_v129.f33);
        let _1702_v143;
        _1702_v143 = _module.D8.create_DC24((((_1701_v142).contains(_this.f27)) ? ((_1701_v142).get(_this.f27)) : (_1686_v129.f33)));
        let _index234 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1692_v133).length));
        (_1692_v133)[_index234] = (_1702_v143).dtor_cf37;
        let _1703_v144;
        _1703_v144 = _dafny.Set.fromElements(_1700_v141.f31);
        _1538_v0 = (_dafny.Set.fromElements(_1538_v0)).IsProperSubsetOf(_1703_v144);
      }
      _1538_v0 = _1538_v0;
      let _1704_v145;
      let _nw271 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1704_v145 = _nw271;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1704_v145).length))) {
        let _1705_i9 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1705_i9)) && ((_1705_i9).isLessThan(new BigNumber((_1704_v145).length))))) {
          (_1704_v145)[(_1705_i9)] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_1538_v0), _dafny.Seq.of(_1538_v0)));
        }
      }
      r0 = !(_1538_v0);
      return r0;
    }
    m3(p0, globalState) {
      let _this = this;
      let _1706_v0;
      _1706_v0 = _dafny.Seq.UnicodeFromString("jhyjv");
      let _1707_v1;
      _1707_v1 = _dafny.MultiSet.fromElements(new BigNumber((_1706_v0).length), _module.__default.fm0(p0, _this.f27, globalState));
      let _1708_v2;
      _1708_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1707_v1,false);
      if ((((_1708_v2).contains(_1707_v1)) ? ((_1708_v2).get(_1707_v1)) : ((_module.__default.fm1(_this.f27, new BigNumber(889), globalState)).IsSubsetOf(_1707_v1)))) {
        if ((p0).isLessThan(new BigNumber(866))) {
          let _1709_v3;
          _1709_v3 = true;
          let _1710_v4;
          let _nw272 = new _module.C2();
          _nw272.__ctor(_1709_v3, p0, (_this).f25);
          _1710_v4 = _nw272;
          (globalState).f17 = _1706_v0;
          let _1711_v5;
          let _nw273 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _1711_v5 = _nw273;
          let _1712_v6;
          _1712_v6 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), ((_1713_p0) => function (_1714_i0) {
            return _1713_p0;
          })(p0))).length), new BigNumber(274), p0, _this.f27, p0);
          let _index235 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1711_v5).length));
          (_1711_v5)[_index235] = _dafny.Seq.update(_1712_v6, _module.__default.safeIndex(p0, new BigNumber((_1712_v6).length)), _this.f27);
          let _index236 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1711_v5).length));
          (_1711_v5)[_index236] = _1712_v6;
          let _1715_v7;
          let _nw274 = Array((new BigNumber(13)).toNumber()).fill(false);
          _1715_v7 = _nw274;
          let _index237 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1715_v7).length));
          (_1715_v7)[_index237] = (_this.f27).isLessThanOrEqualTo(_this.f27);
          let _index238 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1715_v7).length));
          (_1715_v7)[_index238] = _1709_v3;
          _1709_v3 = _1710_v4.f31;
        } else {
          let _1716_v8;
          let _nw275 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1716_v8 = _nw275;
          let _index239 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1716_v8).length));
          (_1716_v8)[_index239] = _1706_v0;
          let _index240 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1716_v8).length));
          (_1716_v8)[_index240] = _1706_v0;
          let _1717_v9;
          _1717_v9 = true;
          (globalState).f2 = !(_1717_v9);
          let _1718_v10;
          let _nw276 = new _module.C0();
          _nw276.__ctor((true) === (_1717_v9));
          _1718_v10 = _nw276;
          let _1719_v11;
          let _nw277 = Array((new BigNumber(15)).toNumber());
          _nw277[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f24);
          _nw277[(_dafny.ONE).toNumber()] = p0;
          _nw277[(new BigNumber(2)).toNumber()] = p0;
          _nw277[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("jk")).length);
          _nw277[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f24);
          _nw277[(new BigNumber(5)).toNumber()] = p0;
          _nw277[(new BigNumber(6)).toNumber()] = _this.f27;
          _nw277[(new BigNumber(7)).toNumber()] = (_this).f24;
          _nw277[(new BigNumber(8)).toNumber()] = _this.f27;
          _nw277[(new BigNumber(9)).toNumber()] = _this.f27;
          _nw277[(new BigNumber(10)).toNumber()] = (_this).f24;
          _nw277[(new BigNumber(11)).toNumber()] = p0;
          _nw277[(new BigNumber(12)).toNumber()] = _this.f27;
          _nw277[(new BigNumber(13)).toNumber()] = p0;
          _nw277[(new BigNumber(14)).toNumber()] = _this.f27;
          _1719_v11 = _nw277;
          let _1720_v12;
          _1720_v12 = _dafny.Seq.of(_1719_v11);
          let _1721_v13;
          _1721_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1719_v11);
          let _1722_v14;
          let _nw278 = Array((new BigNumber(28)).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = (_1720_v12)[_module.__default.safeIndex(p0, new BigNumber((_1720_v12).length))];
          _nw278[(_dafny.ONE).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(2)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(3)).toNumber()] = ((_1717_v9) ? (_1719_v11) : (_1719_v11));
          _nw278[(new BigNumber(4)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(5)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(6)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(7)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(8)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(9)).toNumber()] = (_1720_v12)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1720_v12).length))];
          _nw278[(new BigNumber(10)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(11)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(12)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(13)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(14)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(15)).toNumber()] = ((_1717_v9) ? (_1719_v11) : (_1719_v11));
          _nw278[(new BigNumber(16)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(17)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(18)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(19)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(20)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(21)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(22)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(23)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(24)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(25)).toNumber()] = _1719_v11;
          _nw278[(new BigNumber(26)).toNumber()] = (((_1721_v13).contains((_this).f24)) ? ((_1721_v13).get((_this).f24)) : (_1719_v11));
          _nw278[(new BigNumber(27)).toNumber()] = _1719_v11;
          _1722_v14 = _nw278;
          let _index241 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1722_v14).length));
          (_1722_v14)[_index241] = _1719_v11;
          let _index242 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1722_v14).length));
          (_1722_v14)[_index242] = _1719_v11;
          let _1723_v15;
          let _nw279 = Array((new BigNumber(27)).toNumber());
          _1723_v15 = _nw279;
          let _1724_v16;
          let _nw280 = Array((new BigNumber(16)).toNumber()).fill(false);
          _1724_v16 = _nw280;
          let _1725_v17;
          _1725_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_1724_v16);
          let _1726_v18;
          _1726_v18 = _module.D16.create_DC40(_1723_v15);
          let _rhs269 = ((new BigNumber(691)).multipliedBy(new BigNumber((_1706_v0).length))).multipliedBy(new BigNumber((_1725_v17).length));
          let _rhs270 = (_1726_v18).dtor_cf63;
          let _lhs215 = _this;
          _lhs215.f27 = _rhs269;
          _1723_v15 = _rhs270;
        }
        let _1727_v19;
        let _init39 = function (_1728_i1) {
          return true;
        };
        let _nw281 = Array((new BigNumber(25)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw281.length); _i0_39++) {
          _nw281[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1727_v19 = _nw281;
        let _1729_v20;
        _1729_v20 = true;
        let _index243 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_1727_v19).length));
        (_1727_v19)[_index243] = _1729_v20;
        let _1730_v21;
        _1730_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1729_v20,_1729_v20);
        let _1731_v22;
        _1731_v22 = _dafny.Seq.of(_1706_v0, _1706_v0, _1706_v0);
        let _index244 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_1727_v19).length));
        let _rhs271 = _1729_v20;
        let _rhs272 = (_1730_v21).Merge((_this).f26);
        let _rhs273 = new BigNumber(((_1731_v22)[_module.__default.safeIndex((_this).f24, new BigNumber((_1731_v22).length))]).length);
        let _lhs216 = _1727_v19;
        let _lhs217 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_1727_v19).length));
        let _lhs218 = globalState;
        _lhs216[_lhs217] = _rhs271;
        _1730_v21 = _rhs272;
        _lhs218.f14 = _rhs273;
        let _1732_v23;
        let _nw282 = new _module.C2();
        _nw282.__ctor((new BigNumber(-986)).isLessThanOrEqualTo(_this.f27), ((_dafny.ZERO).minus(_this.f27)).minus(p0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), ((_1733_p0) => function (_1734_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(84)), ((_1735_p0) => function (_1736_i3) {
            return _1735_p0;
          })(_1733_p0))).length));
        })(p0)));
        _1732_v23 = _nw282;
        _1732_v23 = ((!_dafny.areEqual(_dafny.Seq.UnicodeFromString("lylfv"), _1706_v0)) ? (_1732_v23) : (_1732_v23));
        let _index245 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_1727_v19).length));
        (_1727_v19)[_index245] = _1732_v23.f31;
        let _1737_v24;
        _1737_v24 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1738_v25;
        _1738_v25 = _dafny.Set.fromElements(_1727_v19, _1727_v19);
        let _1739_v26;
        _1739_v26 = _dafny.Seq.of(_1738_v25, _1738_v25);
        let _1740_v27;
        _1740_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-379)), function (_1741_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("fg")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-379)), function (_1742_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("fg"))).length)), _1737_v24), _module.__default.safeIndex(_this.f27, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-379)), function (_1743_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("fg")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-379)), function (_1744_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("fg"))).length)), _1737_v24)).length)), _1737_v24)).length),(_1738_v25).Union((_1739_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f27), new BigNumber((_1739_v26).length))]));
        let _1745_v28;
        _1745_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1737_v24,(_this).f24);
        let _1746_v29;
        let _nw283 = Array((new BigNumber(23)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = new BigNumber(427);
        _nw283[(_dafny.ONE).toNumber()] = p0;
        _nw283[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw283[(new BigNumber(3)).toNumber()] = _module.__default.fm0((_this).f24, _this.f27, globalState);
        _nw283[(new BigNumber(4)).toNumber()] = p0;
        _nw283[(new BigNumber(5)).toNumber()] = p0;
        _nw283[(new BigNumber(6)).toNumber()] = p0;
        _nw283[(new BigNumber(7)).toNumber()] = p0;
        _nw283[(new BigNumber(8)).toNumber()] = _this.f27;
        _nw283[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw283[(new BigNumber(10)).toNumber()] = _this.f27;
        _nw283[(new BigNumber(11)).toNumber()] = new BigNumber((_1745_v28).length);
        _nw283[(new BigNumber(12)).toNumber()] = p0;
        _nw283[(new BigNumber(13)).toNumber()] = new BigNumber(((_this).f26).length);
        _nw283[(new BigNumber(14)).toNumber()] = (_this).f24;
        _nw283[(new BigNumber(15)).toNumber()] = p0;
        _nw283[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_1706_v0).length))).length);
        _nw283[(new BigNumber(17)).toNumber()] = new BigNumber(-124);
        _nw283[(new BigNumber(18)).toNumber()] = p0;
        _nw283[(new BigNumber(19)).toNumber()] = _this.f27;
        _nw283[(new BigNumber(20)).toNumber()] = _this.f27;
        _nw283[(new BigNumber(21)).toNumber()] = _this.f27;
        _nw283[(new BigNumber(22)).toNumber()] = _this.f27;
        _1746_v29 = _nw283;
        let _1747_v30;
        _1747_v30 = _dafny.Seq.of(_1746_v29, _1746_v29);
        let _rhs274 = new BigNumber(((((_1740_v27).contains(_this.f27)) ? ((_1740_v27).get(_this.f27)) : (_1738_v25))).length);
        let _rhs275 = _module.__default.fm2((p0).plus(p0), p0, _1737_v24, globalState);
        let _rhs276 = (((_dafny.ZERO).minus(_this.f27)).multipliedBy(new BigNumber((_1747_v30).length))).isLessThan((_this.f27).minus((_this).f24));
        let _rhs277 = _dafny.Seq.Concat(_1706_v0, _1706_v0);
        let _rhs278 = true;
        let _lhs219 = globalState;
        let _lhs220 = globalState;
        let _lhs221 = globalState;
        let _lhs222 = globalState;
        _lhs219.f14 = _rhs274;
        _1729_v20 = _rhs275;
        _lhs220.f2 = _rhs276;
        _lhs221.f17 = _rhs277;
        _lhs222.f2 = _rhs278;
      } else {
        let _1748_v31;
        _1748_v31 = new _dafny.CodePoint('t'.codePointAt(0));
        _1748_v31 = _1748_v31;
        (globalState).f14 = p0;
        let _1749_v32;
        _1749_v32 = true;
        let _1750_v33;
        let _nw284 = Array((new BigNumber(2)).toNumber());
        _nw284[(_dafny.ZERO).toNumber()] = _1749_v32;
        _nw284[(_dafny.ONE).toNumber()] = _1749_v32;
        _1750_v33 = _nw284;
        let _1751_v34;
        _1751_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,!(_1749_v32));
        let _1752_v35;
        _1752_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1751_v34);
        let _1753_v36;
        let _nw285 = Array((new BigNumber(22)).toNumber());
        _nw285[(_dafny.ZERO).toNumber()] = (_1751_v34).Merge(_1751_v34);
        _nw285[(_dafny.ONE).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(2)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32);
        _nw285[(new BigNumber(4)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_module.__default.fm2((_this).f24, _this.f27, _1748_v31, globalState));
        _nw285[(new BigNumber(6)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(7)).toNumber()] = (_1751_v34).Merge((_1751_v34).update(_1750_v33, !(_1749_v32)));
        _nw285[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32);
        _nw285[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,false);
        _nw285[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32);
        _nw285[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32);
        _nw285[(new BigNumber(12)).toNumber()] = ((_1749_v32) ? (_1751_v34) : (_1751_v34));
        _nw285[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32));
        _nw285[(new BigNumber(14)).toNumber()] = (((_1752_v35).contains((_dafny.ZERO).minus(p0))) ? ((_1752_v35).get((_dafny.ZERO).minus(p0))) : (_1751_v34));
        _nw285[(new BigNumber(15)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,!(_1749_v32));
        _nw285[(new BigNumber(17)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(18)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(19)).toNumber()] = _1751_v34;
        _nw285[(new BigNumber(20)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_1750_v33,_1749_v32)).update(_1750_v33, _1749_v32)).Merge(_1751_v34);
        _nw285[(new BigNumber(21)).toNumber()] = _1751_v34;
        _1753_v36 = _nw285;
        let _index246 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1753_v36).length));
        (_1753_v36)[_index246] = ((_1751_v34).update(_1750_v33, _1749_v32)).update(_1750_v33, _1749_v32);
        let _1754_v37;
        _1754_v37 = _dafny.Seq.of(!(_1749_v32), _1749_v32);
        let _1755_v38;
        _1755_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1754_v37,_this.f27);
        let _index247 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1753_v36).length));
        (_1753_v36)[_index247] = _dafny.Map.Empty.slice().updateUnsafe(_1750_v33,((((_1755_v38).contains(_module.__default.fm23(globalState))) ? ((_1755_v38).get(_module.__default.fm23(globalState))) : ((_this).f24))).isLessThan(_this.f27));
        _1749_v32 = (_1754_v37)[_module.__default.safeIndex((_this).f24, new BigNumber((_1754_v37).length))];
        (globalState).f9 = (_this.f27).multipliedBy((_this).f24);
      }
      let _1756_v39;
      _1756_v39 = true;
      let _1757_i5;
      _1757_i5 = _dafny.ZERO;
      L7: {
        while (_1756_v39) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1757_i5)) {
              break L7;
            }
            _1757_i5 = (_1757_i5).plus(_dafny.ONE);
            let _1758_v40;
            _1758_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
            _1758_v40 = (_1758_v40).update((_this).f24, (new BigNumber((_dafny.Set.fromElements(_1756_v39, _1756_v39, _1756_v39)).length)).minus(p0));
            let _1759_v41;
            let _nw286 = new _module.C0();
            _nw286.__ctor(!(_1756_v39) || (_1756_v39));
            _1759_v41 = _nw286;
            let _1760_v42;
            _1760_v42 = new _dafny.CodePoint('o'.codePointAt(0));
            let _1761_v43;
            _1761_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("j"));
            let _1762_v44;
            _1762_v44 = _dafny.Seq.of((_1759_v41).f34, true);
            let _1763_v45;
            _1763_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1762_v44)[_module.__default.safeIndex((_this).f24, new BigNumber((_1762_v44).length))],_module.__default.fm3(globalState));
            let _1764_v46;
            let _nw287 = Array((new BigNumber(29)).toNumber());
            _nw287[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("c");
            _nw287[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jgelqg"), _1706_v0);
            _nw287[(new BigNumber(2)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(3)).toNumber()] = _module.__default.fm4(false, p0, _1760_v42, _module.__default.fm2((_this).f24, (_this).f24, _1760_v42, globalState), globalState);
            _nw287[(new BigNumber(4)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(5)).toNumber()] = (((_1761_v43).contains((_1759_v41).f34)) ? ((_1761_v43).get((_1759_v41).f34)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), ((_1765_v42) => function (_1766_i6) {
              return _1765_v42;
            })(_1760_v42))));
            _nw287[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _nw287[(new BigNumber(7)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_1706_v0, _module.__default.safeIndex(new BigNumber((_1763_v45).length), new BigNumber((_1706_v0).length)), _1760_v42), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_1706_v0, _module.__default.safeIndex(new BigNumber((_1763_v45).length), new BigNumber((_1706_v0).length)), _1760_v42)).length)), new _dafny.CodePoint('f'.codePointAt(0)));
            _nw287[(new BigNumber(9)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(10)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _nw287[(new BigNumber(12)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _nw287[(new BigNumber(14)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), ((_1767_v42) => function (_1768_i7) {
              return _1767_v42;
            })(_1760_v42));
            _nw287[(new BigNumber(16)).toNumber()] = _module.__default.fm4(_1756_v39, _this.f27, _1760_v42, (_1759_v41).f34, globalState);
            _nw287[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), ((_1769_v42) => function (_1770_i8) {
              return _1769_v42;
            })(_1760_v42));
            _nw287[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _nw287[(new BigNumber(19)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _nw287[(new BigNumber(21)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(22)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("kdm");
            _nw287[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("nlv");
            _nw287[(new BigNumber(25)).toNumber()] = _dafny.Seq.update(_1706_v0, _module.__default.safeIndex((_this).f24, new BigNumber((_1706_v0).length)), _1760_v42);
            _nw287[(new BigNumber(26)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(27)).toNumber()] = _1706_v0;
            _nw287[(new BigNumber(28)).toNumber()] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            _1764_v46 = _nw287;
            let _index248 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_1764_v46).length));
            (_1764_v46)[_index248] = _dafny.Seq.Concat(_1706_v0, _1706_v0);
            let _index249 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_1764_v46).length));
            (_1764_v46)[_index249] = _1706_v0;
            let _1771_v47;
            let _nw288 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
            _1771_v47 = _nw288;
            let _1772_v49;
            _1772_v49 = _dafny.Seq.of(_1707_v1);
            let _1773_v50;
            _1773_v50 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),(_this).f24);
            let _1774_v51;
            let _nw289 = new _module.C1();
            _nw289.__ctor(_1772_v49, new BigNumber((_1773_v50).length), new BigNumber((_1706_v0).length), (_this).f25);
            _1774_v51 = _nw289;
            let _1775_v52;
            _1775_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-30),_1774_v51);
            let _1776_v53;
            _1776_v53 = _dafny.Seq.of(_this.f27, p0);
            let _1777_v54;
            _1777_v54 = _dafny.Seq.of(new BigNumber((_1775_v52).length), (_1776_v53)[_module.__default.safeIndex(new BigNumber(-383), new BigNumber((_1776_v53).length))]);
            let _1778_v55;
            _1778_v55 = _module.D12.create_DC33(_1760_v42, (_1759_v41).f34);
            let _1779_v56;
            _1779_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1778_v55).dtor_cf54,(_1759_v41).f34);
            let _1780_v57;
            _1780_v57 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_1777_v54, _module.__default.safeIndex(_this.f27, new BigNumber((_1777_v54).length)), new BigNumber(132))).length), (_this).f24, new BigNumber((_1779_v56).length), new BigNumber((_1707_v1).cardinality()), new BigNumber(320));
            let _index250 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1771_v47).length));
            (_1771_v47)[_index250] = function () {
              let _coll50 = new _dafny.Map();
              for (const _compr_50 of (_1780_v57).Elements) {
                let _1781_v48 = _compr_50;
                if (_dafny.Seq.contains(_1780_v57, _1781_v48)) {
                  _coll50.push([(_1781_v48).plus(new BigNumber(9)),_this.f27]);
                }
              }
              return _coll50;
            }();
            let _1782_v58;
            _1782_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,_1776_v53);
            let _1783_v59;
            _1783_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1759_v41).f34,(((_1782_v58).contains(_1756_v39)) ? ((_1782_v58).get(_1756_v39)) : (_1777_v54)));
            let _1784_v60;
            _1784_v60 = _dafny.MultiSet.fromElements(_1759_v41);
            let _1785_v61;
            _1785_v61 = _module.D16.create_DC41(new BigNumber((_1784_v60).cardinality()));
            let _index251 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1771_v47).length));
            (_1771_v47)[_index251] = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,((_module.__default.fm2(new BigNumber((_dafny.Seq.UnicodeFromString("tvsoofsw")).length), p0, _1760_v42, globalState)) ? (new BigNumber((_1783_v59).length)) : ((_1785_v61).dtor_cf64)));
          }
        }
      }
      (globalState).f1 = p0;
      let _1786_v62;
      let _nw290 = new _module.C0();
      _nw290.__ctor(_1756_v39);
      _1786_v62 = _nw290;
      let _1787_v63;
      let _nw291 = Array((new BigNumber(25)).toNumber());
      _nw291[(_dafny.ZERO).toNumber()] = _1786_v62;
      _nw291[(_dafny.ONE).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(2)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(3)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(4)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(5)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(6)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(7)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(8)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(9)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(10)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(11)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(12)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(13)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(14)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(15)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(16)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(17)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(18)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(19)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(20)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(21)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(22)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(23)).toNumber()] = _1786_v62;
      _nw291[(new BigNumber(24)).toNumber()] = _1786_v62;
      _1787_v63 = _nw291;
      _1787_v63 = _1787_v63;
      let _1788_v64;
      let _init40 = ((_1789_p0) => function (_1790_i9) {
        return (_1790_i9).plus(_1789_p0);
      })(p0);
      let _nw292 = Array((new BigNumber(4)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw292.length); _i0_40++) {
        _nw292[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1788_v64 = _nw292;
      let _index252 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
      (_1788_v64)[_index252] = _this.f27;
      let _index253 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
      let _rhs279 = (new BigNumber((_dafny.Seq.Concat(_1706_v0, _1706_v0)).length)).isLessThanOrEqualTo((_this).f24);
      let _rhs280 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_1706_v0, _dafny.Seq.UnicodeFromString("tytwp"))).length), _module.__default.fm0((_this).f24, p0, globalState)));
      let _lhs223 = _1788_v64;
      let _lhs224 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
      _1756_v39 = _rhs279;
      _lhs223[_lhs224] = _rhs280;
      let _1791_v65;
      _1791_v65 = new _dafny.CodePoint('q'.codePointAt(0));
      let _1792_i10;
      _1792_i10 = _dafny.ZERO;
      L8: {
        while (_module.__default.fm2(p0, p0, _1791_v65, globalState)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1792_i10)) {
              break L8;
            }
            _1792_i10 = (_1792_i10).plus(_dafny.ONE);
            _1707_v1 = _1707_v1;
            if (((!(!(_1756_v39)) || (((((_this).f26).contains(_1756_v39)) ? (((_this).f26).get(_1756_v39)) : ((_1786_v62).f34)))) ? (false) : (!((_1786_v62).f34)))) {
              let _1793_v66;
              _1793_v66 = _dafny.Seq.of((_1786_v62).f34);
              let _1794_v67;
              _1794_v67 = _dafny.Seq.of(_1793_v66);
              let _1795_v68;
              _1795_v68 = _dafny.Set.fromElements((_1794_v67)[_module.__default.safeIndex(new BigNumber((_1706_v0).length), new BigNumber((_1794_v67).length))], _dafny.Seq.Concat(_1793_v66, _1793_v66), _1793_v66);
              let _1796_v69;
              _1796_v69 = _module.D13.create_DC36((_1786_v62).f34, _1756_v39);
              let _1797_v70;
              _1797_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1756_v39),(_1786_v62).f34);
              let _1798_v72;
              _1798_v72 = _1795_v68;
              _1795_v68 = (((_1796_v69).dtor_cf58) ? (_1795_v68) : ((function () {
                let _coll51 = new _dafny.Set();
                for (const _compr_51 of (_1797_v70).Keys.Elements) {
                  let _1799_v71 = _compr_51;
                  if ((_1797_v70).contains(_1799_v71)) {
                    _coll51.add(_1799_v71);
                  }
                }
                return _coll51;
              }()).Intersect((_1798_v72))));
              let _1800_v73;
              _1800_v73 = _dafny.Seq.of(_1707_v1, _1707_v1);
              let _1801_v74;
              _1801_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1756_v39,_1800_v73);
              let _1802_v75;
              let _nw293 = new _module.C1();
              _nw293.__ctor((((_1801_v74).contains((_1786_v62).f34)) ? ((_1801_v74).get((_1786_v62).f34)) : (_1800_v73)), _this.f27, (new BigNumber(325)).minus((_this).f24), _dafny.Seq.Concat((_this).f25, (_this).f25));
              _1802_v75 = _nw293;
              let _1803_v76;
              _1803_v76 = _dafny.Seq.of(_1802_v75);
              (globalState).f18 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_1803_v76, _dafny.Seq.of(_1802_v75, _1802_v75))).length)).multipliedBy(p0));
              let _1804_v77;
              let _nw294 = Array((new BigNumber(9)).toNumber()).fill(false);
              _1804_v77 = _nw294;
              let _1805_v78;
              let _nw295 = Array((new BigNumber(12)).toNumber());
              _nw295[(_dafny.ZERO).toNumber()] = _1804_v77;
              _nw295[(_dafny.ONE).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(2)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(3)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(4)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(5)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(6)).toNumber()] = ((false) ? (_1804_v77) : (_1804_v77));
              _nw295[(new BigNumber(7)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(8)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(9)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(10)).toNumber()] = _1804_v77;
              _nw295[(new BigNumber(11)).toNumber()] = _1804_v77;
              _1805_v78 = _nw295;
              let _index254 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_1805_v78).length));
              (_1805_v78)[_index254] = _1804_v77;
              let _index255 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_1805_v78).length));
              (_1805_v78)[_index255] = _1804_v77;
              (globalState).f1 = _1802_v75.f33;
            } else {
              let _1806_v79;
              _1806_v79 = _dafny.Set.fromElements((_this).f24);
              let _1807_v80;
              _1807_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))],_1806_v79);
              let _1808_v81;
              _1808_v81 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),p0);
              let _1809_v82;
              _1809_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1791_v65,(_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]), _dafny.Map.Empty.slice().updateUnsafe((_1706_v0)[_module.__default.safeIndex(new BigNumber(-891), new BigNumber((_1706_v0).length))],(_this).f24), _1808_v81, _1808_v81, _1808_v81));
              let _1810_v83;
              let _nw296 = new _module.C2();
              _nw296.__ctor(((((_1807_v80).contains(_module.__default.fm0(p0, _this.f27, globalState))) ? ((_1807_v80).get(_module.__default.fm0(p0, _this.f27, globalState))) : (_1806_v79))).IsDisjointFrom(_1806_v79), (_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))], (((_1809_v82).contains((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])) ? ((_1809_v82).get((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])) : ((_this).f25)));
              _1810_v83 = _nw296;
              _1810_v83 = _1810_v83;
              let _1811_v84;
              _1811_v84 = _module.D13.create_DC35(_1707_v1);
              let _1812_v85;
              _1812_v85 = _dafny.Seq.of(_this.f27, (_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))], (_dafny.ZERO).minus((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]));
              (_1810_v83).f31 = ((_1811_v84).dtor_cf57).IsDisjointFrom(_dafny.MultiSet.FromArray(_1812_v85));
              let _1813_v86;
              _1813_v86 = _dafny.Set.fromElements(_1756_v39, _1810_v83.f31);
              let _1814_v87;
              _1814_v87 = _dafny.Seq.of(_1813_v86);
              _1812_v85 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_1815_v64) => function (_1816_i11) {
                return (_1815_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1815_v64).length))];
              })(_1788_v64)), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_1817_v64) => function (_1818_i11) {
                return (_1817_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1817_v64).length))];
              })(_1788_v64))).length)), _this.f27), _module.__default.safeIndex(_this.f27, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_1819_v64) => function (_1820_i11) {
                return (_1819_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1819_v64).length))];
              })(_1788_v64)), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_1821_v64) => function (_1822_i11) {
                return (_1821_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1821_v64).length))];
              })(_1788_v64))).length)), _this.f27)).length)), new BigNumber((_1814_v87).length));
              let _index256 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
              (_1788_v64)[_index256] = (p0).minus((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]);
              let _1823_v88;
              _1823_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xuayu"),(_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])).length),_1810_v83.f31);
              let _1824_v89;
              let _nw297 = Array((new BigNumber(25)).toNumber()).fill(_module.D4.Default());
              _1824_v89 = _nw297;
              let _1825_v90;
              _1825_v90 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,(((((_1823_v88).contains(new BigNumber(302))) ? ((_1823_v88).get(new BigNumber(302))) : (_1810_v83.f31))) ? (_1824_v89) : (_1824_v89)));
              let _1826_v91;
              _1826_v91 = _dafny.Map.Empty.slice().updateUnsafe((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))],_1791_v65);
              let _rhs281 = (_1825_v90).Merge(_1825_v90);
              let _rhs282 = (_1786_v62).f34;
              let _rhs283 = (_1826_v91).Merge(_1826_v91);
              let _lhs225 = globalState;
              _1825_v90 = _rhs281;
              _lhs225.f2 = _rhs282;
              _1826_v91 = _rhs283;
            }
            if ((((_1756_v39) ? ((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]) : ((_this).f24))).isEqualTo((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])) {
              _1756_v39 = (_1786_v62).f34;
              let _1827_v92;
              _1827_v92 = _dafny.Seq.of(_1788_v64);
              let _1828_v93;
              let _nw298 = Array((new BigNumber(7)).toNumber());
              _nw298[(_dafny.ZERO).toNumber()] = _1788_v64;
              _nw298[(_dafny.ONE).toNumber()] = _1788_v64;
              _nw298[(new BigNumber(2)).toNumber()] = _1788_v64;
              _nw298[(new BigNumber(3)).toNumber()] = _1788_v64;
              _nw298[(new BigNumber(4)).toNumber()] = _1788_v64;
              _nw298[(new BigNumber(5)).toNumber()] = (_1827_v92)[_module.__default.safeIndex(_this.f27, new BigNumber((_1827_v92).length))];
              _nw298[(new BigNumber(6)).toNumber()] = _1788_v64;
              _1828_v93 = _nw298;
              let _index257 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1828_v93).length));
              (_1828_v93)[_index257] = _1788_v64;
              let _1829_v94;
              _1829_v94 = _dafny.Seq.of(new BigNumber(-469));
              let _1830_v95;
              _1830_v95 = _dafny.Seq.of(true, true);
              let _1831_v96;
              _1831_v96 = _dafny.Seq.of(_1830_v95);
              let _index258 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
              let _index259 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1828_v93).length));
              let _rhs284 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(-763), new BigNumber((_dafny.Seq.update(_1829_v94, _module.__default.safeIndex(new BigNumber((_module.__default.fm6(new BigNumber((_1829_v94).length), (_this).f24, (_this).f24, _this.f27, globalState)).length), new BigNumber((_1829_v94).length)), p0)).length)), (new BigNumber(423)).multipliedBy(_this.f27));
              let _rhs285 = (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(((_1831_v96)[_module.__default.safeIndex((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))], new BigNumber((_1831_v96).length))]).length))).length))).IsDisjointFrom(_dafny.Set.fromElements((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]));
              let _rhs286 = _1788_v64;
              let _rhs287 = _1706_v0;
              let _rhs288 = _1706_v0;
              let _lhs226 = _1788_v64;
              let _lhs227 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length));
              let _lhs228 = globalState;
              let _lhs229 = _1828_v93;
              let _lhs230 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1828_v93).length));
              let _lhs231 = globalState;
              _lhs226[_lhs227] = _rhs284;
              _lhs228.f2 = _rhs285;
              _lhs229[_lhs230] = _rhs286;
              _lhs231.f17 = _rhs287;
              _1706_v0 = _rhs288;
              let _1832_v97;
              let _nw299 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
              _1832_v97 = _nw299;
              let _1833_v98;
              _1833_v98 = _dafny.Seq.of(_1832_v97, _1832_v97);
              _1832_v97 = (_1833_v98)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f24), new BigNumber((_1833_v98).length))];
              let _1834_v99;
              let _out33;
              _out33 = _module.__default.m0(p0, (_this).f24, globalState);
              _1834_v99 = _out33;
              let _1835_v100;
              _1835_v100 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.fm23(globalState));
              let _1836_v101;
              _1836_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm1(new BigNumber(((_1707_v1).update(_this.f27, _module.__default.abs((_dafny.ZERO).minus(_this.f27)))).cardinality()), _module.__default.fm0(new BigNumber(((((_1835_v100).contains((_1786_v62).f34)) ? ((_1835_v100).get((_1786_v62).f34)) : (_dafny.Seq.of((_1786_v62).f34)))).length), _this.f27, globalState), globalState)).cardinality()),_1834_v99);
              let _1837_v103;
              _1837_v103 = _dafny.Seq.of(function () {
                let _coll52 = new _dafny.Map();
                for (const _compr_52 of _dafny.IntegerRange(new BigNumber(593), new BigNumber(594))) {
                  let _1838_v102 = _compr_52;
                  if (((new BigNumber(593)).isLessThanOrEqualTo(_1838_v102)) && ((_1838_v102).isLessThan(new BigNumber(594)))) {
                    _coll52.push([(_1838_v102).plus(_1834_v99),(_dafny.ZERO).minus((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])]);
                  }
                }
                return _coll52;
              }(), _1836_v101);
              let _1839_v104;
              _1839_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1756_v39,p0);
              let _1840_v105;
              _1840_v105 = _module.D3.create_DC9(_1830_v95);
              let _1841_v106;
              _1841_v106 = _module.D1.create_DC4((_1786_v62).f34, _1756_v39, _1791_v65, (_this).f24, new BigNumber(((_1840_v105).dtor_cf17).length));
              let _1842_v108;
              let _nw300 = Array((new BigNumber(22)).toNumber());
              _nw300[(_dafny.ZERO).toNumber()] = _1836_v101;
              _nw300[(_dafny.ONE).toNumber()] = (_1836_v101).Merge(_1836_v101);
              _nw300[(new BigNumber(2)).toNumber()] = (_1837_v103)[_module.__default.safeIndex((_this).f24, new BigNumber((_1837_v103).length))];
              _nw300[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f27),(_this).f24);
              _nw300[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),_1834_v99)).Merge(_1836_v101);
              _nw300[(new BigNumber(5)).toNumber()] = (_1836_v101).Merge(_1836_v101);
              _nw300[(new BigNumber(6)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,(_dafny.ZERO).minus(p0));
              _nw300[(new BigNumber(8)).toNumber()] = (_1836_v101).update(new BigNumber(187), (_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]);
              _nw300[(new BigNumber(9)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(10)).toNumber()] = ((true) ? (_1836_v101) : (_dafny.Map.Empty.slice().updateUnsafe((((_1839_v104).contains(_1756_v39)) ? ((_1839_v104).get(_1756_v39)) : ((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))])),p0)));
              _nw300[(new BigNumber(11)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(12)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1830_v95).length),_1834_v99);
              _nw300[(new BigNumber(14)).toNumber()] = ((_1756_v39) ? (_1836_v101) : (_module.__default.fm15((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))], (_1841_v106).dtor_cf8, new BigNumber(-368), (_1786_v62).f34, globalState)));
              _nw300[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f27,new BigNumber((_1829_v94).length))).Merge(function () {
                let _coll53 = new _dafny.Map();
                for (const _compr_53 of _dafny.IntegerRange(new BigNumber(794), new BigNumber(-694))) {
                  let _1843_v107 = _compr_53;
                  if (((new BigNumber(794)).isLessThanOrEqualTo(_1843_v107)) && ((_1843_v107).isLessThan(new BigNumber(-694)))) {
                    _coll53.push([(_1843_v107).minus(p0),_this.f27]);
                  }
                }
                return _coll53;
              }());
              _nw300[(new BigNumber(16)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(17)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(18)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(19)).toNumber()] = _1836_v101;
              _nw300[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1834_v99,(_this).f24);
              _nw300[(new BigNumber(21)).toNumber()] = _1836_v101;
              _1842_v108 = _nw300;
              let _index260 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1842_v108).length));
              (_1842_v108)[_index260] = _1836_v101;
              let _index261 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1842_v108).length));
              (_1842_v108)[_index261] = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,(_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]);
            } else {
              _1791_v65 = new _dafny.CodePoint('v'.codePointAt(0));
              let _1844_v110;
              _1844_v110 = _dafny.MultiSet.fromElements(_1791_v65);
              let _1845_v112;
              let _nw301 = new _module.C2();
              _nw301.__ctor(((_dafny.ZERO).minus((_this).f24)).isEqualTo(new BigNumber(318)), (_dafny.ZERO).minus(((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]).minus((_this).f24)), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_1846_v65, _1847_v64) => function (_1848_i12) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1846_v65,(_1847_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1847_v64).length))]);
              })(_1791_v65, _1788_v64)), _module.__default.safeIndex(_module.__default.fm0((_dafny.ZERO).minus(p0), p0, globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), ((_1849_v65, _1850_v64) => function (_1851_i12) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1849_v65,(_1850_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1850_v64).length))]);
              })(_1791_v65, _1788_v64))).length)), function () {
                let _coll54 = new _dafny.Map();
                for (const _compr_54 of (_1844_v110).Elements) {
                  let _1852_v109 = _compr_54;
                  if ((_1844_v110).contains(_1852_v109)) {
                    _coll54.push([_1852_v109,(_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]]);
                  }
                }
                return _coll54;
              }()), _dafny.Seq.Create(_module.__default.abs(new BigNumber(586)), ((_1853_v110) => function (_1854_i13) {
                return function () {
                  let _coll55 = new _dafny.Map();
                  for (const _compr_55 of (_1853_v110).Elements) {
                    let _1855_v111 = _compr_55;
                    if ((_1853_v110).contains(_1855_v111)) {
                      _coll55.push([_1855_v111,(_this).f24]);
                    }
                  }
                  return _coll55;
                }();
              })(_1844_v110))));
              _1845_v112 = _nw301;
              (globalState).f2 = (_1786_v62).f34;
              (_this).f27 = (p0).multipliedBy(_module.__default.safeModuloInt((_this).f24, new BigNumber((_1707_v1).cardinality())));
              (globalState).f2 = _1845_v112.f31;
            }
            let _1856_v114;
            _1856_v114 = _dafny.Seq.of((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))]);
            let _1857_v115;
            _1857_v115 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1856_v114).length),p0);
            let _1858_v116;
            _1858_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1791_v65,_1857_v115);
            let _1859_v117;
            _1859_v117 = _dafny.Map.Empty.slice().updateUnsafe((_1788_v64)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1788_v64).length))],(_1786_v62).f34);
            let _1860_v118;
            _1860_v118 = _dafny.Set.fromElements(_this.f27, (_dafny.ZERO).minus(new BigNumber(((function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of ((((_1858_v116).contains(_1791_v65)) ? ((_1858_v116).get(_1791_v65)) : (_1857_v115))).Keys.Elements) {
                let _1861_v113 = _compr_56;
                if (((((_1858_v116).contains(_1791_v65)) ? ((_1858_v116).get(_1791_v65)) : (_1857_v115))).contains(_1861_v113)) {
                  _coll56.push([(_1861_v113).plus(_this.f27),(_1786_v62).f34]);
                }
              }
              return _coll56;
            }()).Merge(_1859_v117)).length)));
            (globalState).f18 = new BigNumber((_1860_v118).length);
          }
        }
      }
      return;
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
