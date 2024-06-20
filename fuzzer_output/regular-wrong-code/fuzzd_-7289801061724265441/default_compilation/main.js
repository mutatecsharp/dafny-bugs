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
    static fm0(p0, p1, p2, p3, globalState) {
      return false;
    };
    static fm1(p0, p1, p2, globalState) {
      return (_dafny.ZERO).minus(new BigNumber(-300));
    };
    static fm2(p0, globalState) {
      return _module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), function (_0_i0) {
  return new BigNumber(-446);
}));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-787)), function (_1_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length),new BigNumber((_dafny.Seq.UnicodeFromString("fivufydg")).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(390)), function (_2_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-437)), new BigNumber(967))).length),new BigNumber((_dafny.Seq.UnicodeFromString("pevdpvb")).length));
      })), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), function (_3_i2) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),_dafny.ONE);
      }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-719),new BigNumber((_dafny.MultiSet.fromElements(true, true, false, false, false)).cardinality())))));
    };
    static fm5(globalState) {
      return _dafny.Seq.of(!(!(true)));
    };
    static fm6(globalState) {
      let _source0 = _module.D2.create_DC8(_dafny.Seq.of(false, false, false));
      if (_source0.is_DC9) {
        let _4___mcc_h0 = (_source0).cf14;
        let _5___mcc_h1 = (_source0).cf15;
        let _6_cf15 = _5___mcc_h1;
        let _7_cf14 = _4___mcc_h0;
        return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("gm"), _dafny.Seq.UnicodeFromString("byjjmnp"), _dafny.Seq.UnicodeFromString("l"), _dafny.Seq.UnicodeFromString("dkpsd"));
      } else {
        let _8___mcc_h2 = (_source0).cf13;
        let _9_cf13 = _8___mcc_h2;
        return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ihjrffkga"), _dafny.Seq.UnicodeFromString("xiypy"), _dafny.Seq.UnicodeFromString("kxg"));
      }
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ruoqh"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_10_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("irmgafus")));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('n'.codePointAt(0)))).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)))).length)))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), function (_11_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), function (_12_i1) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length);
      })).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(768), new BigNumber(353))));
    };
    static m0(globalState) {
      let r0 = _dafny.Map.Empty;
      let _13_v0;
      _13_v0 = new BigNumber(234);
      let _14_v1;
      _14_v1 = _dafny.Seq.UnicodeFromString("kmmfs");
      let _15_v2;
      _15_v2 = _dafny.MultiSet.fromElements(new BigNumber((_14_v1).length));
      let _16_v3;
      _16_v3 = false;
      let _17_v4;
      _17_v4 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_13_v0);
      let _18_v5;
      _18_v5 = _dafny.MultiSet.fromElements(_16_v3);
      let _19_v6;
      _19_v6 = _dafny.Seq.of(_16_v3, _16_v3);
      let _20_v7;
      _20_v7 = _dafny.Set.fromElements(_13_v0, (_dafny.ZERO).minus(_13_v0), new BigNumber((_19_v6).length), new BigNumber(417));
      let _21_v8;
      let _nw0 = Array((new BigNumber(16)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _13_v0;
      _nw0[(_dafny.ONE).toNumber()] = _13_v0;
      _nw0[(new BigNumber(2)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(3)).toNumber()] = new BigNumber(43);
      _nw0[(new BigNumber(4)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(5)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(6)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(7)).toNumber()] = (((_15_v2).contains(_13_v0)) ? ((_15_v2).get(_13_v0)) : ((((_17_v4).contains(_module.__default.fm0(_13_v0, _16_v3, _16_v3, _14_v1, globalState))) ? ((_17_v4).get(_module.__default.fm0(_13_v0, _16_v3, _16_v3, _14_v1, globalState))) : (_13_v0))));
      _nw0[(new BigNumber(8)).toNumber()] = new BigNumber(554);
      _nw0[(new BigNumber(9)).toNumber()] = (new BigNumber((_18_v5).cardinality())).plus(new BigNumber((_14_v1).length));
      _nw0[(new BigNumber(10)).toNumber()] = (((_18_v5).contains(_16_v3)) ? ((_18_v5).get(_16_v3)) : (_13_v0));
      _nw0[(new BigNumber(11)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(12)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(13)).toNumber()] = _13_v0;
      _nw0[(new BigNumber(14)).toNumber()] = (new BigNumber(380)).multipliedBy(_module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_19_v6).length)), _16_v3, _20_v7, globalState));
      _nw0[(new BigNumber(15)).toNumber()] = _13_v0;
      _21_v8 = _nw0;
      let _22_v9;
      _22_v9 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_21_v8);
      let _index0 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length));
      (_21_v8)[_index0] = new BigNumber((_22_v9).length);
      let _index1 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length));
      (_21_v8)[_index1] = _13_v0;
      let _23_v10;
      let _nw1 = Array((new BigNumber(19)).toNumber()).fill(false);
      _23_v10 = _nw1;
      let _index2 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length));
      (_23_v10)[_index2] = _16_v3;
      let _24_v11;
      _24_v11 = _module.D3.create_DC10(_20_v7);
      let _25_v12;
      _25_v12 = _dafny.Seq.of(_module.__default.fm1(_13_v0, _16_v3, (_24_v11).dtor_cf16, globalState), _13_v0);
      let _26_v13;
      _26_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(985),_17_v4);
      let _27_v14;
      _27_v14 = new _dafny.CodePoint('p'.codePointAt(0));
      let _index3 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length));
      let _rhs0 = _19_v6;
      let _rhs1 = _13_v0;
      let _rhs2 = _16_v3;
      let _rhs3 = ((_dafny.MultiSet.FromArray(_25_v12)).Union(_module.__default.fm8((_25_v12)[_module.__default.safeIndex(new BigNumber(((((_26_v13).contains((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))])) ? ((_26_v13).get((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))])) : (_dafny.Map.Empty.slice().updateUnsafe(_16_v3,_13_v0)))).length), new BigNumber((_25_v12).length))], _16_v3, _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_27_v14), new BigNumber((_19_v6).length), globalState))).Intersect(_dafny.MultiSet.fromElements(_13_v0));
      let _rhs4 = _16_v3;
      let _lhs0 = globalState;
      let _lhs1 = _23_v10;
      let _lhs2 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length));
      _19_v6 = _rhs0;
      _lhs0.f20 = _rhs1;
      _16_v3 = _rhs2;
      _15_v2 = _rhs3;
      _lhs1[_lhs2] = _rhs4;
      let _hi0 = (_13_v0).plus((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))]);
      for (let _28_i0 = new BigNumber((_dafny.Seq.of(_13_v0, _13_v0, (_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))])).length); _28_i0.isLessThan(_hi0); _28_i0 = _28_i0.plus(_dafny.ONE)) {
        let _index4 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length));
        (_23_v10)[_index4] = ((new BigNumber((_dafny.MultiSet.fromElements(_13_v0)).cardinality())).isLessThan(_28_i0)) && (((_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))]) === (_16_v3));
        let _29_v15;
        _29_v15 = _dafny.Map.Empty.slice().updateUnsafe((_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))],_23_v10);
        _29_v15 = (_29_v15).update(((_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))]) || (_16_v3), _23_v10);
        let _30_v16;
        _30_v16 = _dafny.Map.Empty.slice().updateUnsafe(_21_v8,(_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))]);
        let _31_v17;
        _31_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_16_v3);
        let _32_v18;
        _32_v18 = _dafny.Map.Empty.slice().updateUnsafe(_31_v17,_16_v3);
        let _33_v19;
        _33_v19 = _dafny.Seq.of(_32_v18, _32_v18);
        let _34_v20;
        let _nw2 = Array((new BigNumber(13)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = (_30_v16).Merge(_30_v16);
        _nw2[(_dafny.ONE).toNumber()] = _30_v16;
        _nw2[(new BigNumber(2)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(3)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(4)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(5)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_21_v8,new BigNumber(96));
        _nw2[(new BigNumber(7)).toNumber()] = ((_30_v16).update(_21_v8, new BigNumber((_15_v2).cardinality()))).update((((_22_v9).contains(_16_v3)) ? ((_22_v9).get(_16_v3)) : (_21_v8)), (_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))]);
        _nw2[(new BigNumber(8)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_21_v8,(_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))]);
        _nw2[(new BigNumber(10)).toNumber()] = _30_v16;
        _nw2[(new BigNumber(11)).toNumber()] = ((_30_v16).update(_21_v8, new BigNumber(((_33_v19)[_module.__default.safeIndex((_25_v12)[_module.__default.safeIndex(new BigNumber((_25_v12).length), new BigNumber((_25_v12).length))], new BigNumber((_33_v19).length))]).length))).Merge(_30_v16);
        _nw2[(new BigNumber(12)).toNumber()] = _30_v16;
        _34_v20 = _nw2;
        let _35_v21;
        let _nw3 = new _module.C0();
        _nw3.__ctor((_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))]);
        _35_v21 = _nw3;
        let _36_v22;
        _36_v22 = _dafny.MultiSet.fromElements(_35_v21);
        let _index5 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_34_v20).length));
        (_34_v20)[_index5] = (_dafny.Map.Empty.slice().updateUnsafe(_21_v8,(((_36_v22).contains(_35_v21)) ? ((_36_v22).get(_35_v21)) : (_13_v0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_21_v8,new BigNumber((_31_v17).length)));
        let _index6 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_34_v20).length));
        let _rhs5 = (_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))];
        let _rhs6 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_14_v1, _14_v1), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("g"), _module.__default.safeIndex(new BigNumber((_25_v12).length), new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)), new _dafny.CodePoint('e'.codePointAt(0))));
        let _rhs7 = _module.__default.fm1(((_dafny.ZERO).minus((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))])).plus(_28_i0), !_dafny.areEqual(_19_v6, _dafny.Seq.of(_16_v3)), _20_v7, globalState);
        let _rhs8 = _30_v16;
        let _lhs3 = globalState;
        let _lhs4 = _34_v20;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_34_v20).length));
        _16_v3 = _rhs5;
        _16_v3 = _rhs6;
        _lhs3.f20 = _rhs7;
        _lhs4[_lhs5] = _rhs8;
        let _37_v23;
        _37_v23 = _module.D0.create_DC0(_25_v12);
        let _38_v24;
        let _nw4 = Array((new BigNumber(3)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _module.D0.create_DC0(_25_v12);
        _nw4[(_dafny.ONE).toNumber()] = _37_v23;
        _nw4[(new BigNumber(2)).toNumber()] = _module.__default.fm2((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))], globalState);
        _38_v24 = _nw4;
        let _39_v25;
        _39_v25 = _dafny.Map.Empty.slice().updateUnsafe(_35_v21,_38_v24);
        _39_v25 = (_39_v25).update(_35_v21, _38_v24);
      }
      let _index7 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length));
      (_23_v10)[_index7] = (_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))];
      let _40_v26;
      _40_v26 = _dafny.Map.Empty.slice().updateUnsafe((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))],(_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))]);
      let _41_v27;
      _41_v27 = _dafny.Map.Empty.slice().updateUnsafe(((!(_16_v3)) ? (_13_v0) : (_13_v0)),((_module.__default.fm0(new BigNumber((_40_v26).length), _16_v3, _16_v3, _dafny.Seq.UnicodeFromString("whejvaw"), globalState)) ? (_16_v3) : (true)));
      _16_v3 = !((((_41_v27).contains(_13_v0)) ? ((_41_v27).get(_13_v0)) : ((_23_v10)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_23_v10).length))])));
      let _42_v28;
      _42_v28 = _dafny.Map.Empty.slice().updateUnsafe(_23_v10,_module.__default.fm0((_21_v8)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_21_v8).length))], _16_v3, false, _14_v1, globalState));
      _16_v3 = (((_42_v28).contains(_23_v10)) ? ((_42_v28).get(_23_v10)) : ((_16_v3) === (_16_v3)));
      let _43_v29;
      let _nw5 = new _module.C0();
      _nw5.__ctor(_16_v3);
      _43_v29 = _nw5;
      let _44_v30;
      _44_v30 = _dafny.MultiSet.fromElements(_43_v29);
      let _45_v31;
      _45_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_44_v30).cardinality()),_21_v8);
      r0 = _45_v31;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _46_v0;
      _46_v0 = _dafny.Seq.UnicodeFromString("ljlofb");
      let _47_v1;
      let _nw6 = Array((new BigNumber(5)).toNumber()).fill([]);
      _47_v1 = _nw6;
      let _48_v2;
      _48_v2 = new BigNumber(135);
      let _49_v3;
      _49_v3 = _dafny.Seq.of(_48_v2);
      let _50_v4;
      let _nw7 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _50_v4 = _nw7;
      let _51_v5;
      _51_v5 = false;
      let _52_v6;
      _52_v6 = _dafny.Map.Empty.slice().updateUnsafe(_51_v5,new BigNumber(584));
      let _53_v7;
      _53_v7 = new _dafny.CodePoint('t'.codePointAt(0));
      let _54_v8;
      _54_v8 = _dafny.Seq.of(_51_v5);
      let _55_v9;
      _55_v9 = _dafny.Map.Empty.slice().updateUnsafe((_46_v0)[_module.__default.safeIndex(_48_v2, new BigNumber((_46_v0).length))],new BigNumber((_dafny.Seq.UnicodeFromString("apnhjorfd")).length));
      let _56_v10;
      _56_v10 = _dafny.Seq.of(_55_v9, _dafny.Map.Empty.slice().updateUnsafe(_53_v7,_48_v2));
      let _57_v11;
      _57_v11 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,(_56_v10)[_module.__default.safeIndex(_48_v2, new BigNumber((_56_v10).length))]);
      let _58_v12;
      _58_v12 = _dafny.Seq.of(_46_v0);
      let _59_globalState;
      let _nw8 = new _module.GlobalState();
      _nw8.__ctor(_46_v0, _46_v0, false, _47_v1, _49_v3, true, new BigNumber(604), false, _50_v4, (_52_v6).Merge(_52_v6), new BigNumber(503), _dafny.Seq.update(_46_v0, _module.__default.safeIndex(_48_v2, new BigNumber((_46_v0).length)), _53_v7), _dafny.Seq.Concat(_54_v8, _dafny.Seq.update(_54_v8, _module.__default.safeIndex(new BigNumber(9), new BigNumber((_54_v8).length)), _51_v5)), new BigNumber(996), _49_v3, _50_v4, false, new BigNumber(938), (_57_v11).Merge(_57_v11), _58_v12, new BigNumber(465), new BigNumber(48));
      _59_globalState = _nw8;
      let _60_v13;
      let _init0 = function (_61_i0) {
        return false;
      };
      let _nw9 = Array((new BigNumber(8)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw9.length); _i0_0++) {
        _nw9[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _60_v13 = _nw9;
      let _62_v14;
      _62_v14 = _dafny.Seq.of(_60_v13);
      _60_v13 = (_62_v14)[_module.__default.safeIndex(_48_v2, new BigNumber((_62_v14).length))];
      let _hi1 = _48_v2;
      for (let _63_i1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_48_v2, _48_v2, _48_v2, _48_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_46_v0).length),_48_v2)).length)), _module.__default.safeIndex(_48_v2, new BigNumber((_dafny.Seq.of(_48_v2, _48_v2, _48_v2, _48_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_46_v0).length),_48_v2)).length))).length)), _48_v2)).length), _48_v2); _63_i1.isLessThan(_hi1); _63_i1 = _63_i1.plus(_dafny.ONE)) {
        let _64_v15;
        _64_v15 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,new BigNumber(174));
        let _rhs9 = _64_v15;
        let _rhs10 = (_48_v2).multipliedBy(_63_i1);
        let _lhs6 = _59_globalState;
        _64_v15 = _rhs9;
        _lhs6.f20 = _rhs10;
        let _65_v16;
        _65_v16 = _dafny.Seq.of(_50_v4);
        let _66_v17;
        _66_v17 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,(_65_v16)[_module.__default.safeIndex(_48_v2, new BigNumber((_65_v16).length))]);
        let _67_v18;
        _67_v18 = _dafny.Seq.of(_50_v4, _50_v4, (((_66_v17).contains(_48_v2)) ? ((_66_v17).get(_48_v2)) : (_50_v4)));
        _67_v18 = _dafny.Seq.update(_65_v16, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("vpxkaki")).length), new BigNumber((_65_v16).length)), _50_v4);
        _54_v8 = _54_v8;
        _51_v5 = _51_v5;
      }
      _52_v6 = (_52_v6).update(_51_v5, _48_v2);
      let _68_v19;
      let _out0;
      _out0 = _module.__default.m0(_59_globalState);
      _68_v19 = _out0;
      if ((false) || (false)) {
        let _69_v20;
        _69_v20 = _dafny.Seq.of(_dafny.Seq.Concat(_49_v3, _49_v3), _49_v3, _49_v3);
        _69_v20 = _69_v20;
        let _index8 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length));
        (_60_v13)[_index8] = _51_v5;
        let _index9 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length));
        (_60_v13)[_index9] = _51_v5;
        let _70_v21;
        _70_v21 = _dafny.MultiSet.fromElements((_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))]);
        let _index10 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length));
        (_60_v13)[_index10] = !(!(!(_70_v21).contains(_51_v5)));
        _46_v0 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dyxqgdxc"), _46_v0), _module.__default.safeIndex(new BigNumber((function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of _dafny.IntegerRange(new BigNumber(938), new BigNumber(402))) {
            let _71_v22 = _compr_0;
            if (((new BigNumber(938)).isLessThanOrEqualTo(_71_v22)) && ((_71_v22).isLessThan(new BigNumber(402)))) {
              _coll0.push([(_71_v22).multipliedBy(_48_v2),(_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))]]);
            }
          }
          return _coll0;
        }()).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dyxqgdxc"), _46_v0)).length)), _53_v7), _46_v0);
        let _72_v23;
        let _nw10 = Array((new BigNumber(17)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_48_v2, (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))], _51_v5, _46_v0, _59_globalState);
        _nw10[(_dafny.ONE).toNumber()] = _51_v5;
        _nw10[(new BigNumber(2)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(3)).toNumber()] = !(_51_v5);
        _nw10[(new BigNumber(4)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(5)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(6)).toNumber()] = _module.__default.fm0(new BigNumber(-251), false, _51_v5, _46_v0, _59_globalState);
        _nw10[(new BigNumber(7)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(8)).toNumber()] = (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))];
        _nw10[(new BigNumber(9)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(10)).toNumber()] = (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))];
        _nw10[(new BigNumber(11)).toNumber()] = _module.__default.fm0(new BigNumber((_46_v0).length), (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))], _51_v5, _46_v0, _59_globalState);
        _nw10[(new BigNumber(12)).toNumber()] = _51_v5;
        _nw10[(new BigNumber(13)).toNumber()] = (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))];
        _nw10[(new BigNumber(14)).toNumber()] = (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))];
        _nw10[(new BigNumber(15)).toNumber()] = (_60_v13)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_60_v13).length))];
        _nw10[(new BigNumber(16)).toNumber()] = _51_v5;
        _72_v23 = _nw10;
        let _73_v24;
        _73_v24 = _dafny.Map.Empty.slice().updateUnsafe(_72_v23,_48_v2);
        _73_v24 = ((_73_v24).Merge(_73_v24)).Merge(_73_v24);
      } else {
        let _init1 = ((_74_v2) => function (_75_i2) {
          return _module.__default.safeDivisionInt(_75_i2, _74_v2);
        })(_48_v2);
        let _nw11 = Array((new BigNumber(24)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw11.length); _i0_1++) {
          _nw11[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _50_v4 = _nw11;
        let _76_v25;
        let _nw12 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
        _76_v25 = _nw12;
        let _77_v26;
        _77_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_49_v3, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_46_v0).length)), new BigNumber((_49_v3).length)), _48_v2),_76_v25);
        let _78_v27;
        _78_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_54_v8),_49_v3);
        let _79_v28;
        _79_v28 = _dafny.Seq.of(_54_v8);
        _77_v26 = (_77_v26).update((((_78_v27).contains(_79_v28)) ? ((_78_v27).get(_79_v28)) : (_49_v3)), _76_v25);
        let _80_v29;
        _80_v29 = _dafny.Set.fromElements(_60_v13);
        _48_v2 = new BigNumber((_80_v29).length);
        let _index11 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length));
        (_60_v13)[_index11] = _51_v5;
        let _index12 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length));
        (_60_v13)[_index12] = _51_v5;
        let _81_v30;
        let _nw13 = Array((new BigNumber(16)).toNumber()).fill([]);
        _81_v30 = _nw13;
        let _82_v31;
        let _nw14 = Array((new BigNumber(7)).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = _50_v4;
        _nw14[(_dafny.ONE).toNumber()] = _50_v4;
        _nw14[(new BigNumber(2)).toNumber()] = _50_v4;
        _nw14[(new BigNumber(3)).toNumber()] = _50_v4;
        _nw14[(new BigNumber(4)).toNumber()] = _50_v4;
        _nw14[(new BigNumber(5)).toNumber()] = _50_v4;
        _nw14[(new BigNumber(6)).toNumber()] = _50_v4;
        _82_v31 = _nw14;
        let _83_v32;
        _83_v32 = _dafny.Seq.of(_82_v31);
        let _index13 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_81_v30).length));
        (_81_v30)[_index13] = (_83_v32)[_module.__default.safeIndex(_48_v2, new BigNumber((_83_v32).length))];
        let _84_v33;
        _84_v33 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,new BigNumber((_54_v8).length));
        let _85_v34;
        _85_v34 = _dafny.MultiSet.fromElements((_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))]);
        let _86_v35;
        _86_v35 = _dafny.Map.Empty.slice().updateUnsafe((((_85_v34).contains(_module.__default.fm0(_48_v2, (_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))], (_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))], _46_v0, _59_globalState))) ? ((_85_v34).get(_module.__default.fm0(_48_v2, (_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))], (_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))], _46_v0, _59_globalState))) : (_48_v2)),(_60_v13)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length))]);
        let _87_v36;
        _87_v36 = _dafny.Set.fromElements(new BigNumber((_52_v6).length));
        let _index14 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length));
        let _index15 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_81_v30).length));
        let _rhs11 = true;
        let _rhs12 = _82_v31;
        let _rhs13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_48_v2, new BigNumber((_86_v35).length)),_48_v2);
        let _rhs14 = _module.__default.safeDivisionInt(new BigNumber(633), (_48_v2).minus(_module.__default.fm1(_48_v2, false, _87_v36, _59_globalState)));
        let _lhs7 = _60_v13;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_60_v13).length));
        let _lhs9 = _81_v30;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_81_v30).length));
        let _lhs11 = _59_globalState;
        _lhs7[_lhs8] = _rhs11;
        _lhs9[_lhs10] = _rhs12;
        _84_v33 = _rhs13;
        _lhs11.f20 = _rhs14;
      }
      let _hi2 = _48_v2;
      for (let _88_i3 = (new BigNumber(941)).plus(_48_v2); _88_i3.isLessThan(_hi2); _88_i3 = _88_i3.plus(_dafny.ONE)) {
        let _89_v37;
        _89_v37 = _module.D0.create_DC0(_49_v3);
        let _pat_let_tv0 = _49_v3;
        _49_v3 = (function (_pat_let0_0) {
          return function (_90_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_91_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_91_dt__update_hcf0_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_89_v37)).dtor_cf0;
        let _92_v38;
        _92_v38 = _module.D0.create_DC3(_50_v4, _88_i3);
        let _source1 = _92_v38;
        if (_source1.is_DC1) {
          let _93___mcc_h0 = (_source1).cf1;
          let _94___mcc_h1 = (_source1).cf2;
          let _95___mcc_h2 = (_source1).cf3;
          let _96_cf3 = _95___mcc_h2;
          let _97_cf2 = _94___mcc_h1;
          let _98_cf1 = _93___mcc_h0;
          _50_v4 = _50_v4;
          let _index16 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_60_v13).length));
          (_60_v13)[_index16] = !(_97_cf2) || (_97_cf2);
          let _99_v39;
          _99_v39 = _dafny.MultiSet.fromElements(!(_51_v5), _dafny.Seq.contains(_49_v3, (_dafny.ZERO).minus(_48_v2)));
          let _index17 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_60_v13).length));
          let _rhs15 = !(_51_v5) || (_97_cf2);
          let _rhs16 = (_99_v39).Union(_99_v39);
          let _lhs12 = _60_v13;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_60_v13).length));
          _lhs12[_lhs13] = _rhs15;
          _99_v39 = _rhs16;
          let _index18 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_60_v13).length));
          (_60_v13)[_index18] = (_51_v5) === ((_88_i3).isLessThanOrEqualTo(_88_i3));
          _89_v37 = ((_97_cf2) ? (_89_v37) : (_module.__default.fm2(new BigNumber((_54_v8).length), _59_globalState)));
        } else if (_source1.is_DC2) {
          let _100___mcc_h3 = (_source1).cf4;
          let _101_cf4 = _100___mcc_h3;
          let _102_v40;
          _102_v40 = _dafny.Seq.of(_50_v4, _50_v4);
          let _rhs17 = _dafny.Seq.of(_50_v4, _50_v4, _50_v4, _50_v4);
          let _rhs18 = _60_v13;
          let _rhs19 = false;
          let _rhs20 = (_48_v2).plus((_88_i3).minus(_48_v2));
          let _lhs14 = _59_globalState;
          _102_v40 = _rhs17;
          _60_v13 = _rhs18;
          _51_v5 = _rhs19;
          _lhs14.f20 = _rhs20;
          _51_v5 = !(_51_v5) || (((_51_v5) ? (!(_51_v5)) : (_51_v5)));
          (_59_globalState).f20 = _48_v2;
          let _103_v41;
          _103_v41 = _dafny.Set.fromElements(_89_v37, _89_v37);
          _103_v41 = (_103_v41).Union(_103_v41);
        } else if (_source1.is_DC3) {
          let _104___mcc_h4 = (_source1).cf5;
          let _105___mcc_h5 = (_source1).cf6;
          let _106_cf6 = _105___mcc_h5;
          let _107_cf5 = _104___mcc_h4;
          _51_v5 = (!(true)) || (!(!(_51_v5)) || (_51_v5));
          (_59_globalState).f20 = (_88_i3).multipliedBy((_dafny.ZERO).minus((_88_i3).minus(new BigNumber(538))));
          let _rhs21 = (new BigNumber((_49_v3).length)).plus(_106_cf6);
          let _rhs22 = _51_v5;
          _106_cf6 = _rhs21;
          _51_v5 = _rhs22;
          let _108_v42;
          _108_v42 = _dafny.Set.fromElements(_48_v2);
          (_59_globalState).f20 = (_module.__default.fm1(_88_i3, _51_v5, _108_v42, _59_globalState)).minus(_88_i3);
        } else if (_source1.is_DC0) {
          let _109___mcc_h6 = (_source1).cf0;
          let _110_cf0 = _109___mcc_h6;
          let _111_v43;
          _111_v43 = _dafny.Set.fromElements(_48_v2, _48_v2, _88_i3);
          _111_v43 = _111_v43;
          let _112_v44;
          _112_v44 = _dafny.Set.fromElements(_51_v5, (_54_v8)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_54_v8).length))], !((_54_v8)[_module.__default.safeIndex(_48_v2, new BigNumber((_54_v8).length))]), _51_v5);
          _51_v5 = (_112_v44).IsSubsetOf(_112_v44);
          let _113_v45;
          let _nw15 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _113_v45 = _nw15;
          let _114_v46;
          _114_v46 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,new BigNumber(303));
          let _115_v47;
          _115_v47 = _dafny.Seq.of(_114_v46, _114_v46);
          let _index19 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_113_v45).length));
          (_113_v45)[_index19] = _dafny.Seq.Concat(_115_v47, _115_v47);
          let _index20 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_113_v45).length));
          (_113_v45)[_index20] = _module.__default.fm3(_59_globalState);
          let _index21 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_60_v13).length));
          (_60_v13)[_index21] = !(((_51_v5) ? (_51_v5) : (_51_v5)));
          let _116_v48;
          _116_v48 = _module.D0.create_DC2(_51_v5);
          let _index22 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_60_v13).length));
          (_60_v13)[_index22] = (_116_v48).dtor_cf4;
        } else {
          let _117___mcc_h7 = (_source1).cf7;
          let _118_cf7 = _117___mcc_h7;
          let _119_v49;
          _119_v49 = _dafny.Map.Empty.slice().updateUnsafe(_88_i3,_88_i3);
          let _120_v50;
          let _nw16 = new _module.C0();
          _nw16.__ctor((_119_v49).equals(_dafny.Map.Empty.slice().updateUnsafe(_88_i3,_48_v2)));
          _120_v50 = _nw16;
          let _121_v51;
          _121_v51 = _dafny.MultiSet.fromElements((_120_v50).f22);
          _51_v5 = (_121_v51).IsProperSubsetOf((_dafny.MultiSet.FromArray(_module.__default.fm5(_59_globalState))).update(_51_v5, _module.__default.abs(_48_v2)));
          _68_v19 = _68_v19;
          (_59_globalState).f20 = _88_i3;
        }
        (_59_globalState).f20 = ((_51_v5) ? (_module.__default.safeDivisionInt(new BigNumber((_52_v6).length), _48_v2)) : (_48_v2));
        let _122_v52;
        _122_v52 = _dafny.Set.fromElements(_52_v6);
        let _123_v53;
        _123_v53 = _module.D1.create_DC5(_46_v0);
        let _124_v54;
        _124_v54 = _dafny.Map.Empty.slice().updateUnsafe((_122_v52).IsProperSubsetOf(_122_v52),_module.D0.create_DC3(_50_v4, new BigNumber(((_123_v53).dtor_cf8).length)));
        let _125_v55;
        _125_v55 = _dafny.MultiSet.fromElements((_51_v5) === (_51_v5));
        let _126_v56;
        let _init2 = ((_127_v0) => function (_128_i4) {
          return _127_v0;
        })(_46_v0);
        let _nw17 = Array((new BigNumber(26)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw17.length); _i0_2++) {
          _nw17[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _126_v56 = _nw17;
        let _129_v57;
        _129_v57 = _module.D0.create_DC1(_50_v4, _51_v5, _126_v56);
        let _rhs23 = _dafny.Map.Empty.slice().updateUnsafe(_51_v5,((false) ? (_92_v38) : (_92_v38)));
        let _rhs24 = (_dafny.MultiSet.FromArray(_54_v8)).Union(_dafny.MultiSet.fromElements((_129_v57).dtor_cf2, _51_v5));
        _124_v54 = _rhs23;
        _125_v55 = _rhs24;
      }
      _46_v0 = _46_v0;
      let _130_v58;
      let _out1;
      _out1 = _module.__default.m0(_59_globalState);
      _130_v58 = _out1;
      (_59_globalState).f20 = _48_v2;
      let _131_v59;
      _131_v59 = _module.D0.create_DC2(false);
      let _132_i5;
      _132_i5 = _dafny.ZERO;
      L0: {
        while (!((_131_v59).dtor_cf4) || (true)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_132_i5)) {
              break L0;
            }
            _132_i5 = (_132_i5).plus(_dafny.ONE);
            let _133_v60;
            let _nw18 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
            _133_v60 = _nw18;
            let _index23 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_133_v60).length));
            (_133_v60)[_index23] = _module.__default.fm6(_59_globalState);
            let _134_v61;
            _134_v61 = _dafny.Map.Empty.slice().updateUnsafe(_51_v5,_46_v0);
            let _135_v62;
            _135_v62 = _dafny.MultiSet.fromElements((((_134_v61).contains(_51_v5)) ? ((_134_v61).get(_51_v5)) : (_dafny.Seq.UnicodeFromString("lrw"))));
            let _index24 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_133_v60).length));
            (_133_v60)[_index24] = ((_135_v62).Union(_135_v62)).update(_46_v0, _module.__default.abs(_48_v2));
            (_59_globalState).f11 = _46_v0;
            _51_v5 = _51_v5;
            (_59_globalState).f20 = (_dafny.ZERO).minus((_49_v3)[_module.__default.safeIndex(_48_v2, new BigNumber((_49_v3).length))]);
          }
        }
      }
      let _136_v63;
      let _out2;
      _out2 = _module.__default.m0(_59_globalState);
      _136_v63 = _out2;
      if (!(!(!(_48_v2).isEqualTo(new BigNumber(27))))) {
        let _137_v64;
        let _out3;
        _out3 = _module.__default.m0(_59_globalState);
        _137_v64 = _out3;
        if (_51_v5) {
          _50_v4 = _50_v4;
          _51_v5 = _51_v5;
          let _138_v65;
          _138_v65 = _module.D0.create_DC0(_49_v3);
          let _139_v66;
          _139_v66 = _dafny.Seq.of(_module.D0.create_DC4(_138_v65), _138_v65, _module.D0.create_DC4(_138_v65), _138_v65, _138_v65);
          let _140_v67;
          _140_v67 = _dafny.Map.Empty.slice().updateUnsafe(_53_v7,(_139_v66)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_139_v66).length))]);
          let _141_v68;
          _141_v68 = _module.D0.create_DC4((((_140_v67).contains(_53_v7)) ? ((_140_v67).get(_53_v7)) : (_138_v65)));
          let _142_v69;
          let _nw19 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _142_v69 = _nw19;
          let _143_v70;
          _143_v70 = _module.D0.create_DC4(((_51_v5) ? (_138_v65) : (_module.D0.create_DC1(_50_v4, _51_v5, _142_v69))));
          _143_v70 = _module.D0.create_DC4(_138_v65);
          let _144_v71;
          _144_v71 = _dafny.Set.fromElements(_51_v5);
          let _145_v72;
          _145_v72 = _dafny.Seq.of(_144_v71, _144_v71);
          let _146_v73;
          let _nw20 = Array((new BigNumber(9)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_48_v2);
          _nw20[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_48_v2);
          _nw20[(new BigNumber(2)).toNumber()] = _49_v3;
          _nw20[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_49_v3, _dafny.Seq.of(new BigNumber((_144_v71).length)));
          _nw20[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_49_v3, _dafny.Seq.update(_dafny.Seq.of(_48_v2), _module.__default.safeIndex(_48_v2, new BigNumber((_dafny.Seq.of(_48_v2)).length)), new BigNumber((_dafny.Seq.of(_54_v8, _54_v8, _54_v8, _54_v8)).length)));
          _nw20[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_48_v2, _48_v2);
          _nw20[(new BigNumber(6)).toNumber()] = _49_v3;
          _nw20[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), ((_147_v2) => function (_148_i6) {
            return _147_v2;
          })(_48_v2)), _module.__default.safeIndex((_dafny.ZERO).minus(_48_v2), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), ((_149_v2) => function (_150_i6) {
            return _149_v2;
          })(_48_v2))).length)), _48_v2);
          _nw20[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_48_v2), _module.__default.safeIndex((_dafny.ZERO).minus(_48_v2), new BigNumber((_dafny.Seq.of(_48_v2)).length)), new BigNumber((_145_v72).length));
          _146_v73 = _nw20;
          let _index25 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_146_v73).length));
          (_146_v73)[_index25] = _49_v3;
          let _index26 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_146_v73).length));
          (_146_v73)[_index26] = _dafny.Seq.Concat(_dafny.Seq.update(_49_v3, _module.__default.safeIndex(_48_v2, new BigNumber((_49_v3).length)), new BigNumber((_dafny.Seq.of(_48_v2)).length)), _49_v3);
          let _151_v74;
          _151_v74 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_48_v2);
          let _152_v75;
          _152_v75 = _dafny.Set.fromElements(_48_v2, _48_v2);
          (_59_globalState).f20 = (_dafny.ZERO).minus(_module.__default.fm1((((_151_v74).contains(_48_v2)) ? ((_151_v74).get(_48_v2)) : (_module.__default.fm1(_48_v2, true, _152_v75, _59_globalState))), !(_dafny.Set.fromElements(_51_v5)).equals(_144_v71), _152_v75, _59_globalState));
        } else {
          let _153_v76;
          let _nw21 = new _module.C0();
          _nw21.__ctor(_51_v5);
          _153_v76 = _nw21;
          let _154_v77;
          _154_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_49_v3).length),_153_v76);
          (_59_globalState).f20 = (new BigNumber((((_51_v5) ? (_154_v77) : (_154_v77))).length)).minus(new BigNumber(61));
          let _nw22 = new _module.C0();
          _nw22.__ctor((_153_v76).f22);
          _153_v76 = _nw22;
          let _155_v78;
          let _out4;
          _out4 = _module.__default.m0(_59_globalState);
          _155_v78 = _out4;
          let _156_v79;
          let _out5;
          _out5 = _module.__default.m0(_59_globalState);
          _156_v79 = _out5;
          let _157_v80;
          let _nw23 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _157_v80 = _nw23;
          let _158_v81;
          _158_v81 = _dafny.Set.fromElements(_48_v2);
          let _index27 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_157_v80).length));
          (_157_v80)[_index27] = (_dafny.Set.fromElements(_48_v2, new BigNumber((_49_v3).length))).Difference(_158_v81);
          let _index28 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_157_v80).length));
          (_157_v80)[_index28] = (_158_v81).Difference(_158_v81);
        }
        _51_v5 = _51_v5;
        if (_51_v5) {
          let _159_v82;
          let _out6;
          _out6 = _module.__default.m0(_59_globalState);
          _159_v82 = _out6;
          let _160_v83;
          _160_v83 = _dafny.MultiSet.fromElements(_50_v4);
          let _161_v84;
          _161_v84 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(_48_v2), (((_160_v83).contains(_50_v4)) ? ((_160_v83).get(_50_v4)) : ((_dafny.ZERO).minus(_48_v2)))),_51_v5);
          _161_v84 = (_161_v84).update(_48_v2, _51_v5);
          _51_v5 = _51_v5;
          let _162_v85;
          let _nw24 = new _module.C0();
          _nw24.__ctor((_54_v8)[_module.__default.safeIndex(_48_v2, new BigNumber((_54_v8).length))]);
          _162_v85 = _nw24;
          let _163_v86;
          _163_v86 = _dafny.Map.Empty.slice().updateUnsafe(_162_v85,_51_v5);
          let _164_v87;
          _164_v87 = _dafny.Set.fromElements((_163_v86).update(_162_v85, _51_v5), _163_v86, _163_v86);
          let _165_v88;
          _165_v88 = _dafny.Set.fromElements(_48_v2);
          let _166_v89;
          _166_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_module.__default.fm1(new BigNumber((_164_v87).length), _51_v5, _165_v88, _59_globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(328)), ((_167_v7) => function (_168_i7) {
            return _167_v7;
          })(_53_v7))).length)),(_48_v2).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gkst")).length)));
          _166_v89 = _166_v89;
          _51_v5 = _dafny.Seq.IsPrefixOf(_46_v0, _dafny.Seq.of(_53_v7));
        } else {
          (_59_globalState).f20 = (_dafny.ZERO).minus(_48_v2);
          _51_v5 = _51_v5;
          let _169_v90;
          let _out7;
          _out7 = _module.__default.m0(_59_globalState);
          _169_v90 = _out7;
          let _170_v91;
          _170_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_52_v6).length), _48_v2),_53_v7);
          let _index29 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_60_v13).length));
          (_60_v13)[_index29] = _51_v5;
          let _171_v92;
          let _nw25 = new _module.C0();
          _nw25.__ctor(_51_v5);
          _171_v92 = _nw25;
          let _index30 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_60_v13).length));
          (_60_v13)[_index30] = !((_dafny.ZERO).minus(_48_v2)).isEqualTo(_48_v2);
          let _172_v94;
          _172_v94 = _dafny.MultiSet.fromElements(_48_v2);
          let _173_v95;
          _173_v95 = _dafny.Map.Empty.slice().updateUnsafe(_50_v4,_171_v92);
          let _index31 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_60_v13).length));
          let _index32 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_60_v13).length));
          let _rhs25 = function () {
            let _coll1 = new _dafny.Map();
            for (const _compr_1 of (_172_v94).Elements) {
              let _174_v93 = _compr_1;
              if ((_172_v94).contains(_174_v93)) {
                _coll1.push([_module.__default.safeModuloInt(_174_v93, _48_v2),_53_v7]);
              }
            }
            return _coll1;
          }();
          let _rhs26 = _dafny.Seq.Concat(_module.__default.fm7((_171_v92).f22, _48_v2, _59_globalState), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("oqkka"), _module.__default.safeIndex(new BigNumber((_52_v6).length), new BigNumber((_dafny.Seq.UnicodeFromString("oqkka")).length)), (_46_v0)[_module.__default.safeIndex(_48_v2, new BigNumber((_46_v0).length))]));
          let _rhs27 = _51_v5;
          let _rhs28 = _171_v92;
          let _rhs29 = (_dafny.Map.Empty.slice().updateUnsafe(_50_v4,_171_v92)).equals(_173_v95);
          let _lhs15 = _59_globalState;
          let _lhs16 = _60_v13;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_60_v13).length));
          let _lhs18 = _60_v13;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_60_v13).length));
          _170_v91 = _rhs25;
          _lhs15.f0 = _rhs26;
          _lhs16[_lhs17] = _rhs27;
          _171_v92 = _rhs28;
          _lhs18[_lhs19] = _rhs29;
          let _index33 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_50_v4).length));
          (_50_v4)[_index33] = _48_v2;
          let _175_v96;
          _175_v96 = _dafny.Set.fromElements(_48_v2);
          let _index34 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_50_v4).length));
          (_50_v4)[_index34] = _module.__default.fm1(new BigNumber(-810), _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-288)), ((_176_v2) => function (_177_i8) {
            return _176_v2;
          })(_48_v2)), _49_v3), _175_v96, _59_globalState);
        }
        _48_v2 = _48_v2;
      } else {
        let _rhs30 = _53_v7;
        let _rhs31 = _51_v5;
        let _rhs32 = _module.__default.safeModuloInt(_48_v2, _48_v2);
        let _lhs20 = _59_globalState;
        _53_v7 = _rhs30;
        _51_v5 = _rhs31;
        _lhs20.f20 = _rhs32;
        _51_v5 = _51_v5;
        _48_v2 = _48_v2;
        _48_v2 = _48_v2;
        if ((((_51_v5) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(_48_v2))) : (_48_v2))).isLessThan(new BigNumber((_46_v0).length))) {
          _53_v7 = _53_v7;
          let _178_v97;
          _178_v97 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_48_v2, _48_v2),_module.__default.safeDivisionInt(new BigNumber(-178), _48_v2));
          _178_v97 = _178_v97;
          let _179_v98;
          _179_v98 = _module.D0.create_DC3(_50_v4, (_dafny.ZERO).minus((_48_v2).multipliedBy(_48_v2)));
          let _180_v99;
          let _nw26 = Array((new BigNumber(21)).toNumber()).fill([]);
          _180_v99 = _nw26;
          let _181_v100;
          _181_v100 = _dafny.Set.fromElements(_48_v2, _48_v2);
          let _index35 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_50_v4).length));
          (_50_v4)[_index35] = (_module.__default.fm1(_48_v2, _51_v5, _181_v100, _59_globalState)).minus(_48_v2);
          let _182_v101;
          _182_v101 = _dafny.Map.Empty.slice().updateUnsafe(_49_v3,_48_v2);
          let _index36 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_50_v4).length));
          let _rhs33 = _179_v98;
          let _rhs34 = ((!(!(_51_v5)) || (_51_v5)) ? (_180_v99) : (_180_v99));
          let _rhs35 = _module.__default.fm1((new BigNumber((_182_v101).length)).plus(_48_v2), !_dafny.areEqual(_46_v0, _46_v0), _181_v100, _59_globalState);
          let _rhs36 = (_48_v2).multipliedBy(_48_v2);
          let _rhs37 = _module.__default.safeDivisionInt(_48_v2, _48_v2);
          let _lhs21 = _50_v4;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_50_v4).length));
          let _lhs23 = _59_globalState;
          _179_v98 = _rhs33;
          _180_v99 = _rhs34;
          _lhs21[_lhs22] = _rhs35;
          _lhs23.f20 = _rhs36;
          _48_v2 = _rhs37;
          let _183_v102;
          let _nw27 = new _module.C0();
          _nw27.__ctor(true);
          _183_v102 = _nw27;
          let _184_v103;
          _184_v103 = _dafny.Map.Empty.slice().updateUnsafe(true,_54_v8);
          let _185_v104;
          _185_v104 = _module.D2.create_DC8(_dafny.Seq.of(_51_v5, _51_v5));
          let _186_v105;
          let _nw28 = Array((new BigNumber(27)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _54_v8;
          _nw28[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_51_v5);
          _nw28[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_54_v8, _54_v8);
          _nw28[(new BigNumber(3)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(4)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(5)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm5(_59_globalState), _54_v8);
          _nw28[(new BigNumber(7)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(8)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_51_v5);
          _nw28[(new BigNumber(10)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(true);
          _nw28[(new BigNumber(12)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_54_v8, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("fllbype")).length), new BigNumber((_54_v8).length)), _51_v5);
          _nw28[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_51_v5, (_183_v102).f22, !(!(!(_51_v5))), _51_v5, _51_v5);
          _nw28[(new BigNumber(15)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_51_v5);
          _nw28[(new BigNumber(17)).toNumber()] = (((_184_v103).contains((_183_v102).f22)) ? ((_184_v103).get((_183_v102).f22)) : (_dafny.Seq.of(_51_v5)));
          _nw28[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(!(_51_v5));
          _nw28[(new BigNumber(19)).toNumber()] = (_185_v104).dtor_cf13;
          _nw28[(new BigNumber(20)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(21)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(22)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_54_v8, _54_v8);
          _nw28[(new BigNumber(24)).toNumber()] = _54_v8;
          _nw28[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_54_v8, _dafny.Seq.of(_51_v5, false, !(_51_v5), _51_v5));
          _nw28[(new BigNumber(26)).toNumber()] = _54_v8;
          _186_v105 = _nw28;
          let _index37 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_186_v105).length));
          (_186_v105)[_index37] = _dafny.Seq.Concat(_module.__default.fm5(_59_globalState), _54_v8);
          let _index38 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_186_v105).length));
          (_186_v105)[_index38] = _dafny.Seq.Concat(_module.__default.fm5(_59_globalState), _dafny.Seq.Concat(_54_v8, _54_v8));
        } else {
          let _index39 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_50_v4).length));
          (_50_v4)[_index39] = new BigNumber((_54_v8).length);
          let _187_v106;
          let _init3 = ((_188_v0) => function (_189_i9) {
            return _188_v0;
          })(_46_v0);
          let _nw29 = Array((new BigNumber(17)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw29.length); _i0_3++) {
            _nw29[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _187_v106 = _nw29;
          let _index40 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_187_v106).length));
          (_187_v106)[_index40] = _module.__default.fm7(_51_v5, _48_v2, _59_globalState);
          let _190_v107;
          _190_v107 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_51_v5,_48_v2));
          let _index41 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_50_v4).length));
          let _index42 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_187_v106).length));
          let _rhs38 = _46_v0;
          let _rhs39 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_190_v107).length), _48_v2), new BigNumber((_46_v0).length));
          let _rhs40 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), ((_191_v7) => function (_192_i10) {
            return _191_v7;
          })(_53_v7));
          let _rhs41 = _50_v4;
          let _rhs42 = _48_v2;
          let _lhs24 = _50_v4;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_50_v4).length));
          let _lhs26 = _187_v106;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_187_v106).length));
          let _lhs28 = _59_globalState;
          _46_v0 = _rhs38;
          _lhs24[_lhs25] = _rhs39;
          _lhs26[_lhs27] = _rhs40;
          _50_v4 = _rhs41;
          _lhs28.f20 = _rhs42;
          let _193_v108;
          let _out8;
          _out8 = _module.__default.m0(_59_globalState);
          _193_v108 = _out8;
          let _index43 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_60_v13).length));
          (_60_v13)[_index43] = _51_v5;
          let _index44 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_60_v13).length));
          (_60_v13)[_index44] = _51_v5;
          let _194_v109;
          _194_v109 = _dafny.MultiSet.fromElements(_48_v2);
          _194_v109 = _194_v109;
          let _195_v110;
          _195_v110 = _dafny.Set.fromElements(_131_v59);
          let _pat_let_tv1 = _60_v13;
          let _pat_let_tv2 = _60_v13;
          let _196_v111;
          _196_v111 = _dafny.Map.Empty.slice().updateUnsafe(!(_195_v110).contains(function (_pat_let2_0) {
            return function (_197_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_198_dt__update_hcf4_h0) {
                  return _module.D0.create_DC2(_198_dt__update_hcf4_h0);
                }(_pat_let3_0);
              }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_pat_let_tv1).length))]);
            }(_pat_let2_0);
          }(_131_v59)),!(!(!((_60_v13)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_60_v13).length))]))) || (!(false)));
          _196_v111 = (_196_v111).update(false, _module.__default.fm0((_dafny.ZERO).minus(_48_v2), !((_60_v13)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_60_v13).length))]), _module.__default.fm0(_48_v2, _51_v5, _51_v5, _dafny.Seq.UnicodeFromString("t"), _59_globalState), (_187_v106)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_187_v106).length))], _59_globalState));
        }
      }
      let _199_v112;
      let _out9;
      _out9 = _module.__default.m0(_59_globalState);
      _199_v112 = _out9;
      let _200_v113;
      _200_v113 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_49_v3);
      let _201_v114;
      _201_v114 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_48_v2);
      let _rhs43 = (_200_v113).update(_48_v2, _dafny.Seq.of((((_201_v114).contains(_48_v2)) ? ((_201_v114).get(_48_v2)) : (_48_v2))));
      let _rhs44 = _48_v2;
      let _rhs45 = !_dafny.Seq.contains(_58_v12, _dafny.Seq.Concat(_46_v0, _46_v0));
      let _lhs29 = _59_globalState;
      _200_v113 = _rhs43;
      _lhs29.f20 = _rhs44;
      _51_v5 = _rhs45;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_50_v4).length))) {
        let _202_i11 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_202_i11)) && ((_202_i11).isLessThan(new BigNumber((_50_v4).length))))) {
          (_50_v4)[(_202_i11)] = (_202_i11).multipliedBy(_48_v2);
        }
      }
      let _203_i12;
      _203_i12 = _dafny.ZERO;
      L1: {
        while (_51_v5) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_203_i12)) {
              break L1;
            }
            _203_i12 = (_203_i12).plus(_dafny.ONE);
            let _204_v115;
            let _nw30 = new _module.C0();
            _nw30.__ctor(_51_v5);
            _204_v115 = _nw30;
            (_59_globalState).f20 = new BigNumber(612);
            (_59_globalState).f20 = _48_v2;
            if ((_204_v115).f22) {
              let _205_v116;
              _205_v116 = _dafny.Seq.of(_204_v115, _204_v115);
              let _206_v117;
              _206_v117 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_204_v115), _205_v116)).length), new BigNumber(73), _48_v2);
              _206_v117 = _206_v117;
              (_59_globalState).f20 = (_48_v2).minus(_module.__default.safeDivisionInt(_48_v2, _48_v2));
              let _207_v118;
              _207_v118 = _dafny.MultiSet.fromElements(true);
              let _208_v119;
              _208_v119 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_207_v118);
              _208_v119 = (((new BigNumber(50)).isLessThanOrEqualTo(_48_v2)) ? (_208_v119) : (_208_v119));
              let _209_v120;
              let _nw31 = new _module.C0();
              _nw31.__ctor(true);
              _209_v120 = _nw31;
              let _210_v121;
              let _out10;
              _out10 = _module.__default.m0(_59_globalState);
              _210_v121 = _out10;
            } else {
              let _211_v122;
              let _init4 = ((_212_v0) => function (_213_i13) {
                return _212_v0;
              })(_46_v0);
              let _nw32 = Array((new BigNumber(12)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw32.length); _i0_4++) {
                _nw32[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _211_v122 = _nw32;
              let _214_v123;
              _214_v123 = _module.D0.create_DC1(_50_v4, !(false), _211_v122);
              let _rhs46 = (((_204_v115).f22) && ((_204_v115).f22)) || ((_214_v123).dtor_cf2);
              let _rhs47 = _48_v2;
              let _rhs48 = _60_v13;
              let _rhs49 = _48_v2;
              let _lhs30 = _59_globalState;
              let _lhs31 = _59_globalState;
              _51_v5 = _rhs46;
              _lhs30.f20 = _rhs47;
              _60_v13 = _rhs48;
              _lhs31.f20 = _rhs49;
              (_59_globalState).f20 = _48_v2;
              let _215_v125;
              _215_v125 = _dafny.Seq.of((function () {
                let _coll2 = new _dafny.Map();
                for (const _compr_2 of _dafny.IntegerRange(new BigNumber(266), new BigNumber(17))) {
                  let _216_v124 = _compr_2;
                  if (((new BigNumber(266)).isLessThanOrEqualTo(_216_v124)) && ((_216_v124).isLessThan(new BigNumber(17)))) {
                    _coll2.push([(_216_v124).minus(new BigNumber((_46_v0).length)),_48_v2]);
                  }
                }
                return _coll2;
              }()).Merge(_201_v114), (_201_v114).Merge(_dafny.Map.Empty.slice().updateUnsafe(_48_v2,new BigNumber((_dafny.Set.fromElements(_48_v2)).length))), _201_v114);
              let _217_v126;
              _217_v126 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC2((_204_v115).f22),_215_v125);
              _215_v125 = (((_217_v126).contains(_module.D0.create_DC2((_204_v115).f22))) ? ((_217_v126).get(_module.D0.create_DC2((_204_v115).f22))) : (_215_v125));
              let _218_v127;
              _218_v127 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_60_v13);
              _218_v127 = _218_v127;
              let _219_v128;
              _219_v128 = _module.D0.create_DC0(_49_v3);
              let _220_v129;
              _220_v129 = _module.D0.create_DC4(_219_v128);
              let _pat_let_tv3 = _204_v115;
              let _pat_let_tv4 = _219_v128;
              let _221_v130;
              let _nw33 = Array((new BigNumber(17)).toNumber());
              _nw33[(_dafny.ZERO).toNumber()] = _220_v129;
              _nw33[(_dafny.ONE).toNumber()] = _220_v129;
              _nw33[(new BigNumber(2)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(3)).toNumber()] = _module.D0.create_DC4(_219_v128);
              _nw33[(new BigNumber(4)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(5)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(6)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(7)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(8)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(9)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(10)).toNumber()] = function (_pat_let4_0) {
                return function (_222_dt__update__tmp_h2) {
                  return function (_pat_let5_0) {
                    return function (_223_dt__update_hcf7_h0) {
                      return _module.D0.create_DC4(_223_dt__update_hcf7_h0);
                    }(_pat_let5_0);
                  }(_module.D0.create_DC2((_pat_let_tv3).f22));
                }(_pat_let4_0);
              }(_220_v129);
              _nw33[(new BigNumber(11)).toNumber()] = function (_pat_let6_0) {
                return function (_224_dt__update__tmp_h3) {
                  return function (_pat_let7_0) {
                    return function (_225_dt__update_hcf7_h1) {
                      return _module.D0.create_DC4(_225_dt__update_hcf7_h1);
                    }(_pat_let7_0);
                  }(_pat_let_tv4);
                }(_pat_let6_0);
              }(_220_v129);
              _nw33[(new BigNumber(12)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(13)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(14)).toNumber()] = _220_v129;
              _nw33[(new BigNumber(15)).toNumber()] = _module.D0.create_DC4(_219_v128);
              _nw33[(new BigNumber(16)).toNumber()] = _220_v129;
              _221_v130 = _nw33;
              let _index45 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_221_v130).length));
              (_221_v130)[_index45] = _220_v129;
              let _index46 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_221_v130).length));
              (_221_v130)[_index46] = ((_51_v5) ? ((((_204_v115).f22) ? (_220_v129) : (_220_v129))) : (_220_v129));
            }
          }
        }
      }
      process.stdout.write((_46_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_48_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_49_v3, _dafny.Seq.of(new BigNumber(135)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_51_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_52_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(135)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_53_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_54_v8, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_55_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_56_v10, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(9)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(135))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(135),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(135))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_58_v12, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ljlofb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_59_globalState.f0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_59_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_59_globalState).f4, _dafny.Seq.of(new BigNumber(135)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_59_globalState).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(584)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_59_globalState.f11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_59_globalState).f12, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_59_globalState).f14, _dafny.Seq.of(new BigNumber(135)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_59_globalState).f18).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(135),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(135))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_59_globalState.f19, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ljlofb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_59_globalState.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_62_v14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_68_v19).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_130_v58).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v59).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_132_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_136_v63).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_199_v112).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_200_v113).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.of(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_201_v114).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_203_i12));
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
    static create_DC2(cf4) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf5, cf6) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D0(4);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get is_DC4() { return this.$tag === 4; }
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
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 4) {
        return "D0.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([], false, []);
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
    static create_DC6(cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D1(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + this.cf8.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC6(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC9(cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D2(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC9(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC11(cf17, cf18, cf19) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(null, null, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.Seq.UnicodeFromString("");
      this.f11 = _dafny.Seq.UnicodeFromString("");
      this.f19 = _dafny.Seq.of();
      this.f20 = _dafny.ZERO;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f2 = false;
      this._f3 = [];
      this._f4 = _dafny.Seq.of();
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = [];
      this._f9 = _dafny.Map.Empty;
      this._f10 = _dafny.ZERO;
      this._f12 = _dafny.Seq.of();
      this._f13 = _dafny.ZERO;
      this._f14 = _dafny.Seq.of();
      this._f15 = [];
      this._f16 = false;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.Map.Empty;
      this._f21 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this)._f21 = f21;
      return;
    }
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f22 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22) {
      let _this = this;
      (_this)._f22 = f22;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f22;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
