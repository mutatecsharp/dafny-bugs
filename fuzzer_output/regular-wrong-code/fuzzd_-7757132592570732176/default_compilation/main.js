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
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), function (_0_i0) {
        return new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-258)), new BigNumber((_dafny.Seq.of(false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ohtdbkn")).length)), new BigNumber(898), new BigNumber((_dafny.Set.fromElements(new BigNumber(474), new BigNumber((_dafny.Seq.of(new BigNumber(-238))).length))).length))).length);
      }), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), function (_1_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ajayh")).length), new BigNumber((_dafny.Seq.UnicodeFromString("li")).length), new BigNumber(980))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-351)), function (_2_i2) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), function (_3_i3) {
        return new BigNumber(-156);
      })));
    };
    static fm1(p0, p1, p2, globalState) {
      return (!(!(function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_4_i0) {
          return _dafny.Seq.of((_module.D9.create_DC26(new _dafny.CodePoint('o'.codePointAt(0)))).dtor_cf48);
        })).Elements) {
          let _5_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_4_i0) {
            return _dafny.Seq.of((_module.D9.create_DC26(new _dafny.CodePoint('o'.codePointAt(0)))).dtor_cf48);
          }), _5_v0)) {
            _coll0.add(_5_v0);
          }
        }
        return _coll0;
      }()).equals(_dafny.Set.fromElements(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0))))))) === (((false) ? (true) : (false)));
    };
    static fm2(globalState) {
      return (_dafny.ZERO).minus(((new BigNumber(72)).minus(new BigNumber(792))).plus(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(209), new BigNumber(-442))) {
          let _6_v0 = _compr_1;
          if (((new BigNumber(209)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(-442)))) {
            _coll1.push([(_6_v0).plus(new BigNumber(-890)),new BigNumber((_dafny.Seq.UnicodeFromString("o")).length)]);
          }
        }
        return _coll1;
      }()).length)));
    };
    static fm3(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(85)), function (_7_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ypcfvnmt")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), function (_8_i1) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length);
        })).length));
      })).length),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jfvw")).length),false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(!(false), true, false, !(false), false)).cardinality()),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(973),true))));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(((_module.D7.create_DC21(_dafny.MultiSet.fromElements(false))).dtor_cf39).IsProperSubsetOf(_dafny.MultiSet.fromElements(true)));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(153));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1(new BigNumber(433), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), function (_9_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("sogpqka")), _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), function (_10_i1) {
  return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-122))).cardinality());
}))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(270))).length))), !(false));
    };
    static fm17(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()), new BigNumber(97), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_11_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length), new BigNumber(929), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false, true)).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(590))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("enf")).length))).length))).length))), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('q'.codePointAt(0)))).length))).length))).length), new BigNumber(119)));
    };
    static fm20(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),false));
    };
    static fm22(p0, globalState) {
      if (false) {
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-982)), function (_12_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
      } else {
        return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ctdyvhdge"), _dafny.Seq.UnicodeFromString("xrw"))).Intersect(function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("wa"), _dafny.Seq.UnicodeFromString("ucjpv"))).Elements) {
            let _13_v0 = _compr_2;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("wa"), _dafny.Seq.UnicodeFromString("ucjpv")), _13_v0)) {
              _coll2.add(_13_v0);
            }
          }
          return _coll2;
        }());
      }
    };
    static fm23(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),new BigNumber(438)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),new BigNumber((_dafny.Seq.UnicodeFromString("w")).length)))).Intersect(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(540),new BigNumber((_dafny.Seq.of(new BigNumber(-219), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xrdi")).length),new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber(735), new BigNumber(847), new BigNumber(648))).length)))).Elements) {
          let _14_v0 = _compr_3;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(540),new BigNumber((_dafny.Seq.of(new BigNumber(-219), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xrdi")).length),new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber(735), new BigNumber(847), new BigNumber(648))).length)))).contains(_14_v0)) {
            _coll3.add(_14_v0);
          }
        }
        return _coll3;
      }())).Intersect(_dafny.Set.fromElements(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.of(new BigNumber(847), new BigNumber(96))).Elements) {
          let _15_v1 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(847), new BigNumber(96)), _15_v1)) {
            _coll4.push([(_15_v1).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),new BigNumber(298))).length)),new BigNumber(403)]);
          }
        }
        return _coll4;
      }(), (_module.D4.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),new BigNumber(68)))).dtor_cf20, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(720),new BigNumber((_dafny.Seq.UnicodeFromString("mabtyye")).length))));
    };
    static fm25(p0, p1, globalState) {
      let _source0 = _module.D17.create_DC46(_module.D17.create_DC44(_dafny.Set.fromElements(_module.D15.create_DC41(true))));
      if (_source0.is_DC45) {
        let _16___mcc_h0 = (_source0).cf69;
        let _17___mcc_h1 = (_source0).cf70;
        let _18_cf70 = _17___mcc_h1;
        let _19_cf69 = _16___mcc_h0;
        return new _dafny.CodePoint('p'.codePointAt(0));
      } else if (_source0.is_DC44) {
        let _20___mcc_h2 = (_source0).cf68;
        let _21_cf68 = _20___mcc_h2;
        return new _dafny.CodePoint('y'.codePointAt(0));
      } else {
        let _22___mcc_h3 = (_source0).cf71;
        let _23_cf71 = _22___mcc_h3;
        return new _dafny.CodePoint('h'.codePointAt(0));
      }
    };
    static fm30(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(!(true))).Union(_dafny.MultiSet.fromElements(true))).Union(((_module.D7.create_DC21(_dafny.MultiSet.fromElements(false))).dtor_cf39).Union(_dafny.MultiSet.fromElements(!(true))));
    };
    static fm31(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(539),_dafny.Map.Empty.slice().updateUnsafe((_module.D19.create_DC51(new BigNumber((_dafny.Seq.of(new BigNumber(-19), new BigNumber(-441))).length), new BigNumber((_dafny.Seq.UnicodeFromString("u")).length))).dtor_cf77,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, !(false))).length))).cardinality())));
    };
    static fm36(p0, p1, p2, globalState) {
      if (!(_dafny.Set.fromElements(false, false, false)).equals(_dafny.Set.fromElements(!(false), true, false, true, (_module.D0.create_DC3(new BigNumber(233), true)).dtor_cf8))) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), function (_24_i0) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        });
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-849)), function (_25_i1) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("hen"));
      }
    };
    static fm37(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D20.create_DC53();
      if (_source1.is_DC53) {
        return _module.D5.create_DC15(new BigNumber(-438), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ntyca")).length)), new BigNumber(976));
      } else if (_source1.is_DC52) {
        let _26___mcc_h0 = (_source1).cf79;
        let _27_cf79 = _26___mcc_h0;
        return _module.D5.create_DC15(new BigNumber(-615), new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).cardinality()));
      } else {
        let _28___mcc_h1 = (_source1).cf80;
        let _29_cf80 = _28___mcc_h1;
        return _module.D5.create_DC15(new BigNumber(-728), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("xseseylhc"))).length), new BigNumber((_dafny.Seq.UnicodeFromString("bgx")).length));
      }
    };
    static fm38(p0, p1, globalState) {
      let _source2 = _module.D12.create_DC35(new BigNumber(-704));
      if (_source2.is_DC33) {
        let _30___mcc_h0 = (_source2).cf58;
        let _31_cf58 = _30___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(886),_31_cf58)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(476),new BigNumber(-762)));
      } else if (_source2.is_DC34) {
        let _32___mcc_h1 = (_source2).cf59;
        let _33_cf59 = _32___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(535),new BigNumber(167));
      } else if (_source2.is_DC35) {
        let _34___mcc_h2 = (_source2).cf60;
        let _35_cf60 = _34___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(391),_35_cf60);
      } else {
        let _36___mcc_h3 = (_source2).cf57;
        let _37_cf57 = _36___mcc_h3;
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(844),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), function (_38_i0) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length));
        } else {
          return function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-112),new BigNumber(862))).Keys.Elements) {
              let _39_v0 = _compr_5;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-112),new BigNumber(862))).contains(_39_v0)) {
                _coll5.push([_module.__default.safeModuloInt(_39_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(50),new BigNumber((_dafny.Seq.UnicodeFromString("rlsvjvt")).length))).length)),new BigNumber((_dafny.Set.fromElements(false, true)).length)]);
              }
            }
            return _coll5;
          }();
        }
      }
    };
    static fm39(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true));
    };
    static fm40(p0, globalState) {
      return _module.D3.create_DC8((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(124))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(347))));
    };
    static fm41(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), function (_40_i0) {
        return _module.D6.create_DC18(true);
      });
    };
    static fm42(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(false),(new BigNumber((_dafny.Seq.of(false)).length)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-576)), function (_41_i0) {
        return _dafny.Seq.of(false);
      })).length),!(true))).length)));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return _module.D6.create_DC16((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(370), new BigNumber(695))) {
    let _42_v0 = _compr_6;
    if (((new BigNumber(370)).isLessThanOrEqualTo(_42_v0)) && ((_42_v0).isLessThan(new BigNumber(695)))) {
      _coll6.push([(_42_v0).plus(new BigNumber(956)),_module.D5.create_DC15(new BigNumber(-552), new BigNumber((_dafny.Seq.UnicodeFromString("polbyirqn")).length), new BigNumber(-321))]);
    }
  }
  return _coll6;
}()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(826),_module.D5.create_DC15(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), function (_43_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), function (_44_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), function (_45_i2) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})).length)))));
    };
    static fm44(globalState) {
      return ((function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (_dafny.Seq.of(_module.D6.create_DC18(true), _module.D6.create_DC18(true), _module.D6.create_DC18(false), _module.D6.create_DC18(true), _module.D6.create_DC18(false))).Elements) {
          let _46_v0 = _compr_7;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D6.create_DC18(true), _module.D6.create_DC18(true), _module.D6.create_DC18(false), _module.D6.create_DC18(true), _module.D6.create_DC18(false)), _46_v0)) {
            _coll7.add(_46_v0);
          }
        }
        return _coll7;
      }()).Union(_dafny.Set.fromElements(_module.D6.create_DC18(true), _module.D6.create_DC18(false)))).Intersect(_dafny.Set.fromElements(_module.D6.create_DC18(false)));
    };
    static fm45(p0, globalState) {
      return (_module.D9.create_DC26(new _dafny.CodePoint('k'.codePointAt(0)))).dtor_cf48;
    };
    static fm46(p0, p1, globalState) {
      return _module.D4.create_DC12((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), function (_47_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}))).IsSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))), new BigNumber(336), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(true)))),new BigNumber((_dafny.Seq.UnicodeFromString("vul")).length))).length));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),_dafny.Seq.UnicodeFromString("te"))).length),true)).Merge(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(633), new BigNumber(325))) {
          let _48_v0 = _compr_8;
          if (((new BigNumber(633)).isLessThanOrEqualTo(_48_v0)) && ((_48_v0).isLessThan(new BigNumber(325)))) {
            _coll8.push([(_48_v0).minus(new BigNumber(389)),false]);
          }
        }
        return _coll8;
      }());
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D0.create_DC3(new BigNumber(-618), !(true)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D0.create_DC3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)).length))).length), false))))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), function (_49_i0) {
        return _module.D0.create_DC3(new BigNumber(842), !(true));
      })));
    };
    static fm49(globalState) {
      return function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(856))).length)),_dafny.Seq.UnicodeFromString("ss"))).Keys.Elements) {
          let _50_v0 = _compr_9;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(856))).length)),_dafny.Seq.UnicodeFromString("ss"))).contains(_50_v0)) {
            _coll9.add(_50_v0);
          }
        }
        return _coll9;
      }();
    };
    static fm50(globalState) {
      return (((!(false)) ? (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rxvistgq")).length))) : (_dafny.Set.fromElements(new BigNumber(737), new BigNumber(-486))))).Difference((_dafny.Set.fromElements(new BigNumber(-904), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,false))).cardinality()), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), function (_51_i0) {
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality());
      }))).cardinality()))).Intersect(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))).Elements) {
            let _52_v1 = _compr_11;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))), _52_v1)) {
              _coll11.add(_52_v1);
            }
          }
          return _coll11;
        }()).Elements) {
          let _53_v0 = _compr_10;
          if ((function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))).Elements) {
              let _54_v1 = _compr_12;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))), _54_v1)) {
                _coll12.add(_54_v1);
              }
            }
            return _coll12;
          }()).contains(_53_v0)) {
            _coll10.push([_53_v0,new BigNumber((_dafny.MultiSet.fromElements(_module.D18.create_DC47(function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
    let _55_v2 = _compr_13;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_55_v2)) {
      _coll13.push([_55_v2,false]);
    }
  }
  return _coll13;
}()), _module.D18.create_DC47(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),true)))).cardinality())]);
          }
        }
        return _coll10;
      }()).length)), new BigNumber(-460), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length))).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_56_i1) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)))));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(765)), function (_57_i0) {
        return _module.D17.create_DC44(_dafny.Set.fromElements(_module.D15.create_DC41(true)));
      })).length)),new BigNumber((_dafny.Seq.UnicodeFromString("hct")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(298)), function (_58_i1) {
        return new BigNumber((_dafny.Set.fromElements(true)).length);
      })),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-268)), function (_59_i2) {
        return new BigNumber(-586);
      })).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-826), new BigNumber((_dafny.Seq.UnicodeFromString("lvbhnjfvo")).length), new BigNumber(40), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).cardinality()))).length),true)).length), new BigNumber(598)),new BigNumber(424)));
    };
    static fm52(globalState) {
      return (_dafny.MultiSet.fromElements(_module.D3.create_DC9(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-271))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(-945), new BigNumber(163))).cardinality()), new BigNumber((_dafny.Seq.of(true)).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-984)), function (_60_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("qdnyphax")).length);
})).length), new BigNumber(315))).length))).length))).cardinality()), function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of _dafny.IntegerRange(new BigNumber(643), new BigNumber(200))) {
    let _61_v0 = _compr_14;
    if (((new BigNumber(643)).isLessThanOrEqualTo(_61_v0)) && ((_61_v0).isLessThan(new BigNumber(200)))) {
      _coll14.push([(_61_v0).plus(new BigNumber(675)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), function (_62_i1) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length))]);
    }
  }
  return _coll14;
}(), new BigNumber((_dafny.Seq.UnicodeFromString("eaevsmgns")).length), (_module.D17.create_DC45(new BigNumber(648), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-451)), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(347), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))).cardinality()))).length), new BigNumber(217)))).cardinality()))))).dtor_cf69), _module.D3.create_DC9(new BigNumber(582), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(517),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),false)).length))), new BigNumber(-58), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length)))).Intersect((_dafny.MultiSet.fromElements(_module.D3.create_DC9(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, true, !(false), false)).length))).cardinality()), _dafny.Map.Empty.slice().updateUnsafe((_module.D19.create_DC51(new BigNumber(160), new BigNumber((_dafny.Seq.of(new BigNumber(185))).length))).dtor_cf78,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-564),(_module.D12.create_DC33(new BigNumber((_dafny.Seq.of(true)).length))).dtor_cf58)), new BigNumber(434), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new _dafny.CodePoint('l'.codePointAt(0)))).length)), _module.D3.create_DC9(new BigNumber(542), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(246),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(5),false)).length),new BigNumber(759))).length),new BigNumber(769))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("pf")).length)))).Difference(_dafny.MultiSet.fromElements(_module.D3.create_DC9(new BigNumber(-617), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-228))).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(337),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hjucomr")).length)))), new BigNumber((_dafny.Seq.UnicodeFromString("hckgvkew")).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), _module.D3.create_DC9(new BigNumber(-585), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)),function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-411),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-818),new BigNumber(-633))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(935),true)).length)))).length))).length))).Keys.Elements) {
    let _63_v1 = _compr_15;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-411),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-818),new BigNumber(-633))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(935),true)).length)))).length))).length))).contains(_63_v1)) {
      _coll15.push([(_63_v1).minus(new BigNumber(744)),new BigNumber(567)]);
    }
  }
  return _coll15;
}()), new BigNumber(-490), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(391),true)).length)), _module.D3.create_DC9(new BigNumber((_dafny.Seq.UnicodeFromString("iduro")).length), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-987),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("a")).length),new BigNumber((_dafny.Seq.of(false)).length))), new BigNumber(348), new BigNumber((function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of (_dafny.Seq.of(new BigNumber(213), new BigNumber((_dafny.Seq.UnicodeFromString("sokhgatk")).length))).Elements) {
    let _64_v2 = _compr_16;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(213), new BigNumber((_dafny.Seq.UnicodeFromString("sokhgatk")).length)), _64_v2)) {
      _coll16.push([_module.__default.safeDivisionInt(_64_v2, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-745))).cardinality())),true]);
    }
  }
  return _coll16;
}()).length)))));
    };
    static fm53(p0, p1, globalState) {
      return _module.D3.create_DC9(new BigNumber(576), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("mvqhguysd")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_65_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length))), new BigNumber(377), (_dafny.ZERO).minus(((false) ? (new BigNumber(-266)) : (new BigNumber(-394)))));
    };
    static fm54(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(150)))).contains(new BigNumber((_dafny.Set.fromElements(false, true, false)).length)),((false) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(609),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), function (_66_i0) {
        return new BigNumber(-711);
      })).length), new BigNumber((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Seq.of(new BigNumber(-746))).Elements) {
          let _67_v0 = _compr_17;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-746)), _67_v0)) {
            _coll17.push([_module.__default.safeModuloInt(_67_v0, new BigNumber((_dafny.Seq.UnicodeFromString("xflc")).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rbjpcvxsh"),true)).length)]);
          }
        }
        return _coll17;
      }()).length)))).cardinality()))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_module.D6.create_DC19(!(true), true, false, true, _dafny.Seq.of(!(true))))).length),new BigNumber(666)))));
    };
    static fm55(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),_module.D7.create_DC21(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_module.D7.create_DC21(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(false)))))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),_module.D7.create_DC21(_dafny.MultiSet.fromElements(false))));
    };
    static fm56(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_68_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-61),true);
      }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(788),false), function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(974), new BigNumber(-465))) {
          let _69_v0 = _compr_18;
          if (((new BigNumber(974)).isLessThanOrEqualTo(_69_v0)) && ((_69_v0).isLessThan(new BigNumber(-465)))) {
            _coll18.push([_module.__default.safeModuloInt(_69_v0, new BigNumber((_dafny.Set.fromElements(false, !(!(true)))).length)),false]);
          }
        }
        return _coll18;
      }()));
    };
    static fm57(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of _dafny.IntegerRange(new BigNumber(134), new BigNumber(571))) {
          let _70_v0 = _compr_19;
          if (((new BigNumber(134)).isLessThanOrEqualTo(_70_v0)) && ((_70_v0).isLessThan(new BigNumber(571)))) {
            _coll19.push([(_70_v0).minus(new BigNumber((function () {
              let _coll20 = new _dafny.Map();
              for (const _compr_20 of _dafny.IntegerRange(new BigNumber(574), new BigNumber(-763))) {
                let _71_v1 = _compr_20;
                if (((new BigNumber(574)).isLessThanOrEqualTo(_71_v1)) && ((_71_v1).isLessThan(new BigNumber(-763)))) {
                  _coll20.push([(_71_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).length)),true]);
                }
              }
              return _coll20;
            }()).length)),false]);
          }
        }
        return _coll19;
      }(),_dafny.Seq.of(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(48),false),_dafny.Seq.of(false)))).Merge((function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(652), new BigNumber(162))) {
            let _72_v3 = _compr_22;
            if (((new BigNumber(652)).isLessThanOrEqualTo(_72_v3)) && ((_72_v3).isLessThan(new BigNumber(162)))) {
              _coll22.push([_module.__default.safeModuloInt(_72_v3, new BigNumber(955)),false]);
            }
          }
          return _coll22;
        }(),true)).Keys.Elements) {
          let _73_v2 = _compr_21;
          if ((_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(652), new BigNumber(162))) {
              let _72_v3 = _compr_23;
              if (((new BigNumber(652)).isLessThanOrEqualTo(_72_v3)) && ((_72_v3).isLessThan(new BigNumber(162)))) {
                _coll23.push([_module.__default.safeModuloInt(_72_v3, new BigNumber(955)),false]);
              }
            }
            return _coll23;
          }(),true)).contains(_73_v2)) {
            _coll21.push([_73_v2,_dafny.Seq.of(false)]);
          }
        }
        return _coll21;
      }()).Merge(function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true))).length)), new BigNumber(729))).length))).Elements) {
            let _74_v5 = _compr_25;
            if ((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true))).length)), new BigNumber(729))).length))).contains(_74_v5)) {
              _coll25.push([(_74_v5).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("a")).length))).length))).length)),true]);
            }
          }
          return _coll25;
        }(),new BigNumber((_dafny.Set.fromElements(new BigNumber(712))).length))).Keys.Elements) {
          let _75_v4 = _compr_24;
          if ((_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true))).length)), new BigNumber(729))).length))).Elements) {
              let _74_v5 = _compr_26;
              if ((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true))).length)), new BigNumber(729))).length))).contains(_74_v5)) {
                _coll26.push([(_74_v5).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("a")).length))).length))).length)),true]);
              }
            }
            return _coll26;
          }(),new BigNumber((_dafny.Set.fromElements(new BigNumber(712))).length))).contains(_75_v4)) {
            _coll24.push([_75_v4,_dafny.Seq.of(!(false))]);
          }
        }
        return _coll24;
      }()));
    };
    static fm58(p0, globalState) {
      return _module.D15.create_DC41((_dafny.MultiSet.fromElements(new BigNumber(317), new BigNumber((_dafny.Seq.UnicodeFromString("cjorlfrb")).length), new BigNumber(636))).IsSubsetOf(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jk"),(_dafny.Seq.of(_module.D13.create_DC37(), _module.D13.create_DC37())))).length)), new BigNumber(880), new BigNumber((function () {
  let _coll27 = new _dafny.Map();
  for (const _compr_27 of _dafny.IntegerRange(new BigNumber(949), new BigNumber(919))) {
    let _76_v0 = _compr_27;
    if (((new BigNumber(949)).isLessThanOrEqualTo(_76_v0)) && ((_76_v0).isLessThan(new BigNumber(919)))) {
      _coll27.push([_module.__default.safeDivisionInt(_76_v0, new BigNumber((_dafny.Set.fromElements(false, true)).length)),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())]);
    }
  }
  return _coll27;
}()).length))));
    };
    static fm59(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_dafny.Set.fromElements(new BigNumber(170), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xuifnsqld")).length),new BigNumber(197))).length))).Elements) {
          let _77_v0 = _compr_28;
          if ((_dafny.Set.fromElements(new BigNumber(170), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xuifnsqld")).length),new BigNumber(197))).length))).contains(_77_v0)) {
            _coll28.push([(_77_v0).multipliedBy(new BigNumber(437)),new BigNumber(242)]);
          }
        }
        return _coll28;
      }(),new BigNumber(548))).Merge((_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(746), new BigNumber(-949))) {
          let _78_v1 = _compr_29;
          if (((new BigNumber(746)).isLessThanOrEqualTo(_78_v1)) && ((_78_v1).isLessThan(new BigNumber(-949)))) {
            _coll29.push([_module.__default.safeDivisionInt(_78_v1, new BigNumber(915)),new BigNumber(607)]);
          }
        }
        return _coll29;
      }(),new BigNumber(449))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(391)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(701)))))).cardinality()),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hcxgjmau"),true)).length))))).cardinality())),new BigNumber(737))));
    };
    static fm60(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),!(false)));
    };
    static fm61(globalState) {
      return _module.D15.create_DC42(_module.D15.create_DC41(true));
    };
    static fm62(p0, p1, p2, p3, globalState) {
      if (true) {
        return _dafny.MultiSet.fromElements(_dafny.Seq.of((_module.D6.create_DC18(false)).dtor_cf32, false), _dafny.Seq.of(false, true, false, true));
      } else {
        return (_dafny.MultiSet.fromElements(_dafny.Seq.of(!(true)))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, true)));
      }
    };
    static fm63(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(642))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of _dafny.IntegerRange(new BigNumber(360), new BigNumber(62))) {
          let _79_v0 = _compr_30;
          if (((new BigNumber(360)).isLessThanOrEqualTo(_79_v0)) && ((_79_v0).isLessThan(new BigNumber(62)))) {
            _coll30.push([(_79_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ihi")).length)),new BigNumber((_dafny.Seq.of(new BigNumber(90))).length)]);
          }
        }
        return _coll30;
      }())).length), new BigNumber(10), new BigNumber(817)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-504)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(494), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(746))).length), new BigNumber((_dafny.Seq.of(false, true)).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-829)), function (_80_i0) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), function (_81_i1) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length));
      })));
    };
    static fm64(globalState) {
      let _source3 = _module.D9.create_DC26(new _dafny.CodePoint('v'.codePointAt(0)));
      if (_source3.is_DC27) {
        let _82___mcc_h0 = (_source3).cf49;
        let _83___mcc_h1 = (_source3).cf50;
        let _84___mcc_h2 = (_source3).cf51;
        let _85___mcc_h3 = (_source3).cf52;
        let _86_cf52 = _85___mcc_h3;
        let _87_cf51 = _84___mcc_h2;
        let _88_cf50 = _83___mcc_h1;
        let _89_cf49 = _82___mcc_h0;
        return _module.D18.create_DC48();
      } else if (_source3.is_DC26) {
        let _90___mcc_h4 = (_source3).cf48;
        let _91_cf48 = _90___mcc_h4;
        return _module.D18.create_DC48();
      } else {
        let _92___mcc_h5 = (_source3).cf53;
        let _93_cf53 = _92___mcc_h5;
        return _module.D18.create_DC48();
      }
    };
    static fm65(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).Union(function () {
        let _coll31 = new _dafny.Set();
        for (const _compr_31 of (_dafny.MultiSet.fromElements(new BigNumber(451))).Elements) {
          let _94_v0 = _compr_31;
          if ((_dafny.MultiSet.fromElements(new BigNumber(451))).contains(_94_v0)) {
            _coll31.add((_94_v0).plus(new BigNumber(408)));
          }
        }
        return _coll31;
      }());
    };
    static fm66(p0, globalState) {
      return _module.D20.create_DC53();
    };
    static fm67(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)));
    };
    static fm68(p0, p1, globalState) {
      return _dafny.Set.fromElements(_module.D20.create_DC53());
    };
    static fm69(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(((_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
  let _coll32 = new _dafny.Set();
  for (const _compr_32 of _dafny.IntegerRange(new BigNumber(395), new BigNumber(-382))) {
    let _95_v0 = _compr_32;
    if (((new BigNumber(395)).isLessThanOrEqualTo(_95_v0)) && ((_95_v0).isLessThan(new BigNumber(-382)))) {
      _coll32.add((_95_v0).minus(new BigNumber(909)));
    }
  }
  return _coll32;
}()).length)))).dtor_cf14).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.UnicodeFromString("rdbdqhl")).length)));
    };
    static fm70(globalState) {
      return function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).Elements) {
          let _96_v0 = _compr_33;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).contains(_96_v0)) {
            _coll33.push([_96_v0,new BigNumber(((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false))).length)]);
          }
        }
        return _coll33;
      }();
    };
    static fm71(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(714)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), function (_97_i0) {
        return new BigNumber(754);
      }))).Difference(function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length)))).Elements) {
          let _98_v0 = _compr_34;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length))), _98_v0)) {
            _coll34.add(_98_v0);
          }
        }
        return _coll34;
      }());
    };
    static m0(globalState) {
      let r0 = [];
      let r1 = false;
      let _99_v0;
      _99_v0 = new BigNumber(826);
      let _100_v1;
      _100_v1 = _dafny.Seq.of(_99_v0);
      let _101_v2;
      _101_v2 = _dafny.Set.fromElements((_100_v1)[_module.__default.safeIndex(_99_v0, new BigNumber((_100_v1).length))], _99_v0);
      let _102_v3;
      let _init0 = ((_103_v0) => function (_104_i0) {
        return (_104_i0).multipliedBy(_103_v0);
      })(_99_v0);
      let _nw0 = Array((new BigNumber(15)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _102_v3 = _nw0;
      let _105_v4;
      _105_v4 = _module.D0.create_DC2(new BigNumber((_101_v2).length), _102_v3);
      let _106_v5;
      _106_v5 = _dafny.Set.fromElements(_105_v4, _105_v4);
      let _source4 = (((_106_v5).IsSubsetOf(_106_v5)) ? (_105_v4) : (_module.D0.create_DC2(_99_v0, _102_v3)));
      if (_source4.is_DC1) {
        let _107___mcc_h0 = (_source4).cf1;
        let _108___mcc_h1 = (_source4).cf2;
        let _109___mcc_h2 = (_source4).cf3;
        let _110___mcc_h3 = (_source4).cf4;
        let _111_cf4 = _110___mcc_h3;
        let _112_cf3 = _109___mcc_h2;
        let _113_cf2 = _108___mcc_h1;
        let _114_cf1 = _107___mcc_h0;
        if (!(!((_111_cf4) && (_111_cf4)))) {
          let _115_v6;
          _115_v6 = new _dafny.CodePoint('l'.codePointAt(0));
          let _116_v7;
          _116_v7 = _dafny.MultiSet.fromElements(_115_v6, _115_v6);
          let _117_v8;
          _117_v8 = _dafny.Set.fromElements(_116_v7, (_116_v7).Union(_116_v7), _dafny.MultiSet.fromElements(_115_v6));
          _117_v8 = _117_v8;
          (globalState).f8 = _111_cf4;
          let _118_v9;
          let _init1 = ((_119_cf4) => function (_120_i1) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_119_cf4, _119_cf4, _119_cf4, _119_cf4), _dafny.Seq.of(_119_cf4));
          })(_111_cf4);
          let _nw1 = Array((new BigNumber(11)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw1.length); _i0_1++) {
            _nw1[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _118_v9 = _nw1;
          let _121_v10;
          _121_v10 = _dafny.Seq.of(_111_cf4);
          let _rhs0 = (_121_v10)[_module.__default.safeIndex(_99_v0, new BigNumber((_121_v10).length))];
          let _rhs1 = _111_cf4;
          let _rhs2 = !((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(829), _99_v0))).isEqualTo(_99_v0);
          let _rhs3 = (_99_v0).multipliedBy(_module.__default.safeModuloInt(_112_cf3, (_100_v1)[_module.__default.safeIndex(_114_cf1, new BigNumber((_100_v1).length))]));
          let _rhs4 = ((_111_cf4) ? (_118_v9) : (_118_v9));
          let _lhs0 = globalState;
          _lhs0.f8 = _rhs0;
          r1 = _rhs1;
          r1 = _rhs2;
          _112_cf3 = _rhs3;
          _118_v9 = _rhs4;
          let _122_v11;
          let _nw2 = new _module.C16();
          _nw2.__ctor(_111_cf4, _112_cf3, _111_cf4, _module.__default.fm2(globalState), _114_cf1);
          _122_v11 = _nw2;
          (globalState).f8 = _111_cf4;
        } else {
          (globalState).f8 = _111_cf4;
          let _123_v12;
          let _nw3 = Array((new BigNumber(7)).toNumber()).fill(false);
          _123_v12 = _nw3;
          let _index0 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          (_123_v12)[_index0] = _111_cf4;
          let _124_v13;
          _124_v13 = _dafny.Map.Empty.slice().updateUnsafe(_99_v0,_112_cf3);
          let _125_v14;
          _125_v14 = _module.D4.create_DC12(_111_cf4, new BigNumber((_124_v13).length), _99_v0);
          let _126_v15;
          _126_v15 = _dafny.Seq.of((_125_v14).dtor_cf21, _111_cf4);
          let _127_v16;
          _127_v16 = _dafny.MultiSet.fromElements(_111_cf4, _module.__default.fm1(new BigNumber(838), _111_cf4, (_dafny.ZERO).minus(_114_cf1), globalState));
          let _index1 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          (_123_v12)[_index1] = (_126_v15)[_module.__default.safeIndex((((_127_v16).contains(_111_cf4)) ? ((_127_v16).get(_111_cf4)) : (_114_cf1)), new BigNumber((_126_v15).length))];
          _112_cf3 = new BigNumber(-68);
          (globalState).f8 = _module.__default.fm1((_112_cf3).minus(_114_cf1), _111_cf4, (_99_v0).multipliedBy(new BigNumber((_101_v2).length)), globalState);
          let _index2 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          let _index3 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          let _rhs5 = !(_111_cf4);
          let _rhs6 = _99_v0;
          let _rhs7 = (_123_v12)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length))];
          let _lhs1 = _123_v12;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          let _lhs3 = _123_v12;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_123_v12).length));
          _lhs1[_lhs2] = _rhs5;
          _112_cf3 = _rhs6;
          _lhs3[_lhs4] = _rhs7;
        }
        _99_v0 = _112_cf3;
        let _128_v17;
        let _init2 = ((_129_cf4) => function (_130_i2) {
          return ((!(true)) ? (_129_cf4) : (_129_cf4));
        })(_111_cf4);
        let _nw4 = Array((new BigNumber(10)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw4.length); _i0_2++) {
          _nw4[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _128_v17 = _nw4;
        let _131_v18;
        _131_v18 = _dafny.MultiSet.fromElements(false);
        let _index4 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_128_v17).length));
        (_128_v17)[_index4] = (new BigNumber((_131_v18).cardinality())).isLessThanOrEqualTo(_99_v0);
        let _132_v19;
        _132_v19 = _dafny.Seq.of(_111_cf4);
        let _index5 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_128_v17).length));
        (_128_v17)[_index5] = _dafny.areEqual(_dafny.Seq.Concat(_132_v19, _dafny.Seq.update(_132_v19, _module.__default.safeIndex(new BigNumber((_module.__default.fm50(globalState)).length), new BigNumber((_132_v19).length)), _111_cf4)), _dafny.Seq.of(_module.__default.fm1(_112_cf3, _111_cf4, _99_v0, globalState)));
        let _133_v20;
        _133_v20 = _dafny.Seq.of(_132_v19);
        let _134_v21;
        let _nw5 = new _module.C7();
        _nw5.__ctor(new BigNumber((_module.__default.fm0(globalState)).length), _114_cf1);
        _134_v21 = _nw5;
        let _135_v22;
        _135_v22 = _dafny.Map.Empty.slice().updateUnsafe(_134_v21,_132_v19);
        r1 = !_dafny.Seq.contains(_dafny.Seq.of((((_135_v22).contains(_134_v21)) ? ((_135_v22).get(_134_v21)) : (_dafny.Seq.of((_128_v17)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_128_v17).length))])))), _dafny.Seq.update((_133_v20)[_module.__default.safeIndex(_112_cf3, new BigNumber((_133_v20).length))], _module.__default.safeIndex(_99_v0, new BigNumber(((_133_v20)[_module.__default.safeIndex(_112_cf3, new BigNumber((_133_v20).length))]).length)), _111_cf4));
      } else if (_source4.is_DC2) {
        let _136___mcc_h4 = (_source4).cf5;
        let _137___mcc_h5 = (_source4).cf6;
        let _138_cf6 = _137___mcc_h5;
        let _139_cf5 = _136___mcc_h4;
        let _140_v23;
        _140_v23 = true;
        let _141_v24;
        _141_v24 = new _dafny.CodePoint('r'.codePointAt(0));
        let _142_v25;
        _142_v25 = _dafny.MultiSet.fromElements(_141_v24);
        let _143_v26;
        _143_v26 = _dafny.Map.Empty.slice().updateUnsafe(_138_cf6,new BigNumber((_142_v25).cardinality()));
        let _144_v27;
        let _nw6 = new _module.C16();
        _nw6.__ctor((false) && (false), _139_cf5, _140_v23, _module.__default.fm2(globalState), new BigNumber(((_143_v26).Merge(_143_v26)).length));
        _144_v27 = _nw6;
        _144_v27 = _144_v27;
        let _145_v28;
        _145_v28 = _module.D20.create_DC54(_module.__default.fm66(_141_v24, globalState));
        _145_v28 = _145_v28;
        let _rhs8 = _99_v0;
        let _rhs9 = _139_cf5;
        let _rhs10 = _dafny.areEqual(_dafny.Seq.Concat(_module.__default.fm0(globalState), _100_v1), _100_v1);
        _99_v0 = _rhs8;
        _99_v0 = _rhs9;
        r1 = _rhs10;
        let _146_v29;
        _146_v29 = _dafny.Map.Empty.slice().updateUnsafe(_144_v27.f11,((_dafny.ZERO).minus(_139_cf5)).plus(_99_v0));
        _99_v0 = new BigNumber((_146_v29).length);
      } else if (_source4.is_DC3) {
        let _147___mcc_h6 = (_source4).cf7;
        let _148___mcc_h7 = (_source4).cf8;
        let _149_cf8 = _148___mcc_h7;
        let _150_cf7 = _147___mcc_h6;
        _149_cf8 = _149_cf8;
        let _151_v30;
        let _nw7 = Array((new BigNumber(2)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = true;
        _nw7[(_dafny.ONE).toNumber()] = _149_cf8;
        _151_v30 = _nw7;
        let _152_v31;
        _152_v31 = _dafny.MultiSet.fromElements(_151_v30);
        let _153_v32;
        _153_v32 = _dafny.MultiSet.fromElements(_99_v0, _99_v0);
        let _154_v33;
        let _nw8 = new _module.C12();
        _nw8.__ctor((_152_v31).Intersect(_152_v31), _99_v0, (_153_v32).IsSubsetOf(_153_v32), _150_cf7, _99_v0);
        _154_v33 = _nw8;
        (globalState).f8 = _149_cf8;
        let _155_v34;
        _155_v34 = _dafny.Map.Empty.slice().updateUnsafe((_154_v33).f19,_module.D5.create_DC15((_154_v33).f19, (_154_v33).f19, new BigNumber(703)));
        let _156_v35;
        _156_v35 = _module.D6.create_DC16(_155_v34);
        _156_v35 = _156_v35;
      } else {
        let _157___mcc_h8 = (_source4).cf0;
        let _158_cf0 = _157___mcc_h8;
        let _159_v36;
        _159_v36 = true;
        r1 = _159_v36;
        let _160_v37;
        _160_v37 = new _dafny.CodePoint('g'.codePointAt(0));
        _160_v37 = _160_v37;
        let _161_v38;
        _161_v38 = _dafny.MultiSet.fromElements((_105_v4).dtor_cf6, _102_v3, _102_v3, _102_v3);
        let _162_v39;
        _162_v39 = _dafny.Seq.of(_161_v38, (_dafny.MultiSet.fromElements(_102_v3, _102_v3)).update(_102_v3, _module.__default.abs(new BigNumber(-378))), _161_v38);
        let _163_v40;
        _163_v40 = _dafny.Map.Empty.slice().updateUnsafe((_162_v39)[_module.__default.safeIndex(_99_v0, new BigNumber((_162_v39).length))],_99_v0);
        let _164_v41;
        _164_v41 = _dafny.Map.Empty.slice().updateUnsafe(_160_v37,new BigNumber(911));
        let _165_v42;
        _165_v42 = _dafny.Seq.of(_159_v36);
        _163_v40 = (_163_v40).update(((_module.__default.fm1(new BigNumber(172), _159_v36, new BigNumber((_164_v41).length), globalState)) ? (_161_v38) : (_161_v38)), new BigNumber((((_159_v36) ? (_165_v42) : (_dafny.Seq.update(_165_v42, _module.__default.safeIndex(_99_v0, new BigNumber((_165_v42).length)), _159_v36)))).length));
        let _166_v43;
        let _nw9 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _166_v43 = _nw9;
        let _index6 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_166_v43).length));
        (_166_v43)[_index6] = _dafny.Seq.UnicodeFromString("rhw");
        let _167_v44;
        _167_v44 = _dafny.Seq.UnicodeFromString("pcb");
        let _index7 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_166_v43).length));
        (_166_v43)[_index7] = _167_v44;
      }
      let _168_v45;
      _168_v45 = true;
      let _169_v46;
      _169_v46 = new _dafny.CodePoint('i'.codePointAt(0));
      let _170_v47;
      _170_v47 = _dafny.Set.fromElements(_169_v46);
      (globalState).f8 = (_module.__default.fm67(_168_v45, _168_v45, _99_v0, globalState)).equals(_170_v47);
      if (_module.__default.fm1(new BigNumber((_170_v47).length), false, _99_v0, globalState)) {
        if (true) {
          let _171_v48;
          let _nw10 = new _module.C15();
          _nw10.__ctor(_168_v45, _99_v0, _168_v45);
          _171_v48 = _nw10;
          let _172_v49;
          _172_v49 = _dafny.Seq.of(_171_v48);
          let _173_v50;
          _173_v50 = _dafny.Map.Empty.slice().updateUnsafe(_172_v49,_99_v0);
          let _174_v51;
          _174_v51 = _dafny.Map.Empty.slice().updateUnsafe((_173_v50).Merge(_dafny.Map.Empty.slice().updateUnsafe(_172_v49,_99_v0)),_171_v48.f16);
          _174_v51 = (_174_v51).update(_173_v50, _171_v48.f16);
          let _index8 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length));
          (_102_v3)[_index8] = _171_v48.f17;
          let _index9 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length));
          (_102_v3)[_index9] = _99_v0;
          let _175_v52;
          let _nw11 = Array((new BigNumber(11)).toNumber()).fill(false);
          _175_v52 = _nw11;
          let _176_v53;
          _176_v53 = _dafny.MultiSet.fromElements(_175_v52);
          let _177_v54;
          _177_v54 = _dafny.Seq.of(_171_v48.f16, _171_v48.f16);
          let _178_v55;
          _178_v55 = _dafny.MultiSet.fromElements(_168_v45, (_177_v54)[_module.__default.safeIndex(_171_v48.f17, new BigNumber((_177_v54).length))], _171_v48.f16, true, _168_v45);
          let _179_v56;
          let _nw12 = new _module.C12();
          _nw12.__ctor(_176_v53, (new BigNumber((_178_v55).cardinality())).multipliedBy(_171_v48.f17), _168_v45, _171_v48.f17, (_102_v3)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length))]);
          _179_v56 = _nw12;
          let _180_v57;
          _180_v57 = _dafny.Map.Empty.slice().updateUnsafe((_102_v3)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length))],_175_v52);
          let _181_v58;
          let _nw13 = Array((new BigNumber(17)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _175_v52;
          _nw13[(_dafny.ONE).toNumber()] = _175_v52;
          _nw13[(new BigNumber(2)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(3)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(4)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(5)).toNumber()] = ((_179_v56.f11) ? (_175_v52) : (_175_v52));
          _nw13[(new BigNumber(6)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(7)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(8)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(9)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(10)).toNumber()] = (((_180_v57).contains((_dafny.ZERO).minus(_99_v0))) ? ((_180_v57).get((_dafny.ZERO).minus(_99_v0))) : (_175_v52));
          _nw13[(new BigNumber(11)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(12)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(13)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(14)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(15)).toNumber()] = _175_v52;
          _nw13[(new BigNumber(16)).toNumber()] = _175_v52;
          _181_v58 = _nw13;
          let _rhs11 = _179_v56.f11;
          let _rhs12 = _179_v56;
          let _rhs13 = _171_v48.f16;
          let _rhs14 = (_102_v3)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length))];
          let _rhs15 = _181_v58;
          let _lhs5 = globalState;
          let _lhs6 = _171_v48;
          _lhs5.f8 = _rhs11;
          _179_v56 = _rhs12;
          r1 = _rhs13;
          _lhs6.f17 = _rhs14;
          _181_v58 = _rhs15;
          let _index10 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_102_v3).length));
          (_102_v3)[_index10] = new BigNumber((_100_v1).length);
          let _182_v59;
          _182_v59 = _dafny.Map.Empty.slice().updateUnsafe(_171_v48.f16,_179_v56.f11);
          let _183_v60;
          let _nw14 = new _module.C14();
          _nw14.__ctor((_171_v48.f16) === ((((_182_v59).contains(_179_v56.f11)) ? ((_182_v59).get(_179_v56.f11)) : (!(_179_v56.f11)))));
          _183_v60 = _nw14;
        } else {
          let _184_v61;
          let _nw15 = new _module.C4();
          _nw15.__ctor((_99_v0).multipliedBy(_99_v0), _168_v45);
          _184_v61 = _nw15;
          let _185_v62;
          let _nw16 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
          _185_v62 = _nw16;
          let _186_v63;
          _186_v63 = _dafny.MultiSet.fromElements(_168_v45, true, !(!(_168_v45)));
          let _index11 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_185_v62).length));
          (_185_v62)[_index11] = _186_v63;
          let _187_v64;
          _187_v64 = _dafny.Seq.UnicodeFromString("ncihubldw");
          let _188_v65;
          let _nw17 = new _module.C8();
          _nw17.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), ((_189_v0) => function (_190_i3) {
            return _189_v0;
          })(_99_v0)));
          _188_v65 = _nw17;
          let _191_v66;
          _191_v66 = _dafny.Map.Empty.slice().updateUnsafe(_188_v65,_module.__default.fm36(_99_v0, _99_v0, _169_v46, globalState));
          let _192_v67;
          _192_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cgjuj"),_dafny.Seq.UnicodeFromString("s"));
          let _193_v68;
          _193_v68 = _dafny.Map.Empty.slice().updateUnsafe(_99_v0,_169_v46);
          let _194_v69;
          let _nw18 = Array((new BigNumber(26)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _187_v64;
          _nw18[(_dafny.ONE).toNumber()] = _187_v64;
          _nw18[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("xkd");
          _nw18[(new BigNumber(3)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("hrjw");
          _nw18[(new BigNumber(5)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(6)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(7)).toNumber()] = _module.__default.fm36((_184_v61).f25, (_184_v61).f25, _169_v46, globalState);
          _nw18[(new BigNumber(8)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(9)).toNumber()] = (((_191_v66).contains(_188_v65)) ? ((_191_v66).get(_188_v65)) : ((((_192_v67).contains(_187_v64)) ? ((_192_v67).get(_187_v64)) : (_187_v64))));
          _nw18[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_187_v64, _module.__default.safeIndex((_184_v61).f25, new BigNumber((_187_v64).length)), (((_193_v68).contains(_module.__default.fm2(globalState))) ? ((_193_v68).get(_module.__default.fm2(globalState))) : (_169_v46)));
          _nw18[(new BigNumber(11)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(12)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(13)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(14)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(15)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(16)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(17)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
          _nw18[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), ((_195_v46) => function (_196_i4) {
            return _195_v46;
          })(_169_v46)), _module.__default.safeIndex((_184_v61).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), ((_197_v46) => function (_198_i4) {
            return _197_v46;
          })(_169_v46))).length)), _169_v46);
          _nw18[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("bljdb");
          _nw18[(new BigNumber(21)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(22)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(23)).toNumber()] = _187_v64;
          _nw18[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("xdgj");
          _nw18[(new BigNumber(25)).toNumber()] = _187_v64;
          _194_v69 = _nw18;
          let _199_v70;
          _199_v70 = _dafny.Map.Empty.slice().updateUnsafe(_194_v69,_184_v61);
          let _200_v71;
          _200_v71 = _dafny.Seq.of(_186_v63, _186_v63);
          let _201_v72;
          let _nw19 = Array((new BigNumber(12)).toNumber()).fill(false);
          _201_v72 = _nw19;
          let _202_v73;
          _202_v73 = _dafny.Seq.of(_201_v72);
          let _index12 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_185_v62).length));
          let _rhs16 = (((_199_v70).contains(_194_v69)) ? ((_199_v70).get(_194_v69)) : (_184_v61));
          let _rhs17 = (_200_v71)[_module.__default.safeIndex(new BigNumber((_202_v73).length), new BigNumber((_200_v71).length))];
          let _lhs7 = _185_v62;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_185_v62).length));
          _184_v61 = _rhs16;
          _lhs7[_lhs8] = _rhs17;
          let _203_v74;
          _203_v74 = _dafny.MultiSet.fromElements(_201_v72, _201_v72);
          let _204_v75;
          let _nw20 = new _module.C12();
          _nw20.__ctor(_203_v74, new BigNumber(211), true, new BigNumber(121), _99_v0);
          _204_v75 = _nw20;
          let _205_v76;
          _205_v76 = _dafny.Seq.of(_204_v75);
          _99_v0 = new BigNumber((_205_v76).length);
          _99_v0 = new BigNumber((_170_v47).length);
          _99_v0 = (_184_v61).f25;
          let _206_v77;
          _206_v77 = _module.D21.create_DC56(true, !(_168_v45), _99_v0);
          _99_v0 = (_206_v77).dtor_cf84;
        }
        if (_168_v45) {
          let _207_v78;
          let _init3 = ((_208_v45) => function (_209_i5) {
            return !((_208_v45) || (_208_v45));
          })(_168_v45);
          let _nw21 = Array((new BigNumber(16)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw21.length); _i0_3++) {
            _nw21[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _207_v78 = _nw21;
          let _index13 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_207_v78).length));
          (_207_v78)[_index13] = _168_v45;
          let _210_v79;
          _210_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(925),_168_v45);
          let _211_v80;
          _211_v80 = _dafny.Map.Empty.slice().updateUnsafe(_99_v0,!(_210_v79).contains(_99_v0));
          let _index14 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_207_v78).length));
          (_207_v78)[_index14] = (((_211_v80).contains(_module.__default.safeModuloInt(new BigNumber(102), _99_v0))) ? ((_211_v80).get(_module.__default.safeModuloInt(new BigNumber(102), _99_v0))) : (_168_v45));
          (globalState).f8 = _168_v45;
          _99_v0 = _99_v0;
          (globalState).f8 = (_207_v78)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_207_v78).length))];
          let _212_v81;
          let _nw22 = new _module.C10();
          _nw22.__ctor((_207_v78)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_207_v78).length))]);
          _212_v81 = _nw22;
        } else {
          let _index15 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length));
          (_102_v3)[_index15] = _99_v0;
          let _index16 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length));
          (_102_v3)[_index16] = _99_v0;
          let _213_v82;
          _213_v82 = _dafny.Map.Empty.slice().updateUnsafe(_100_v1,_168_v45);
          let _214_v83;
          _214_v83 = _dafny.Map.Empty.slice().updateUnsafe((_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))],new BigNumber(196));
          let _215_v85;
          _215_v85 = _dafny.Seq.of(function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_module.__default.fm68(_169_v46, _168_v45, globalState)).Elements) {
              let _216_v84 = _compr_35;
              if ((_module.__default.fm68(_169_v46, _168_v45, globalState)).contains(_216_v84)) {
                _coll35.push([_216_v84,_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_168_v45,_168_v45), _dafny.Map.Empty.slice().updateUnsafe(_168_v45,_168_v45))]);
              }
            }
            return _coll35;
          }());
          let _217_v87;
          _217_v87 = _module.D20.create_DC53();
          let _218_v88;
          _218_v88 = _dafny.Map.Empty.slice().updateUnsafe(_168_v45,_168_v45);
          let _219_v89;
          _219_v89 = _dafny.MultiSet.fromElements(_218_v88);
          let _220_v90;
          _220_v90 = _dafny.Seq.UnicodeFromString("tf");
          let _221_v91;
          _221_v91 = ((_219_v89).update(_218_v88, _module.__default.abs(new BigNumber((_220_v90).length)))).update(_218_v88, _module.__default.abs(_99_v0));
          let _222_v92;
          _222_v92 = _dafny.Map.Empty.slice().updateUnsafe(_217_v87,_221_v91);
          let _223_v93;
          _223_v93 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_217_v87,_219_v89), _222_v92);
          let _224_v94;
          let _nw23 = Array((new BigNumber(16)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = (new BigNumber((_213_v82).length)).isLessThan(_99_v0);
          _nw23[(_dafny.ONE).toNumber()] = !(!(!((_214_v83).contains(_module.__default.fm2(globalState)))));
          _nw23[(new BigNumber(2)).toNumber()] = false;
          _nw23[(new BigNumber(3)).toNumber()] = _168_v45;
          _nw23[(new BigNumber(4)).toNumber()] = !(_168_v45);
          _nw23[(new BigNumber(5)).toNumber()] = _168_v45;
          _nw23[(new BigNumber(6)).toNumber()] = _168_v45;
          _nw23[(new BigNumber(7)).toNumber()] = (new BigNumber(980)).isLessThanOrEqualTo(_99_v0);
          _nw23[(new BigNumber(8)).toNumber()] = false;
          _nw23[(new BigNumber(9)).toNumber()] = (_99_v0).isLessThan(new BigNumber((_100_v1).length));
          _nw23[(new BigNumber(10)).toNumber()] = !(_module.__default.fm2(globalState)).isEqualTo((_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))]);
          _nw23[(new BigNumber(11)).toNumber()] = (function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of (_215_v85).Elements) {
              let _225_v86 = _compr_36;
              if (_dafny.Seq.contains(_215_v85, _225_v86)) {
                _coll36.add(_225_v86);
              }
            }
            return _coll36;
          }()).IsDisjointFrom(_223_v93);
          _nw23[(new BigNumber(12)).toNumber()] = _168_v45;
          _nw23[(new BigNumber(13)).toNumber()] = _168_v45;
          _nw23[(new BigNumber(14)).toNumber()] = !_dafny.Seq.contains(_100_v1, _99_v0);
          _nw23[(new BigNumber(15)).toNumber()] = _168_v45;
          _224_v94 = _nw23;
          let _index17 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length));
          (_224_v94)[_index17] = !(_168_v45);
          let _226_v95;
          _226_v95 = _dafny.MultiSet.fromElements(_168_v45);
          let _index18 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length));
          (_224_v94)[_index18] = ((((_226_v95).contains(_168_v45)) ? ((_226_v95).get(_168_v45)) : (_99_v0))).isLessThan((_99_v0).plus((_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))]));
          let _227_v96;
          _227_v96 = _module.D15.create_DC41(_168_v45);
          let _228_v97;
          _228_v97 = _dafny.Map.Empty.slice().updateUnsafe(_169_v46,_227_v96);
          _228_v97 = (_228_v97).update(_169_v46, _227_v96);
          let _229_v98;
          _229_v98 = _module.D19.create_DC51(((_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))]).minus(new BigNumber((_220_v90).length)), (((_226_v95).contains(_168_v45)) ? ((_226_v95).get(_168_v45)) : ((_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))])));
          _229_v98 = ((!(_226_v95).equals(_226_v95)) ? (_229_v98) : ((((_224_v94)[_module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length))]) ? (_229_v98) : (_229_v98))));
          let _230_v99;
          let _nw24 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _230_v99 = _nw24;
          let _index19 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_230_v99).length));
          (_230_v99)[_index19] = _169_v46;
          let _index20 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length));
          let _index21 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_230_v99).length));
          let _index22 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length));
          let _rhs18 = (_102_v3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length))];
          let _rhs19 = new _dafny.CodePoint('w'.codePointAt(0));
          let _rhs20 = (false) && ((_224_v94)[_module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length))]);
          let _lhs9 = _102_v3;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_102_v3).length));
          let _lhs11 = _230_v99;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_230_v99).length));
          let _lhs13 = _224_v94;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_224_v94).length));
          _lhs9[_lhs10] = _rhs18;
          _lhs11[_lhs12] = _rhs19;
          _lhs13[_lhs14] = _rhs20;
        }
        let _231_v100;
        let _nw25 = new _module.C2();
        _nw25.__ctor((_99_v0).multipliedBy(_99_v0), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("mvjqaslaa"), _module.__default.safeIndex((_dafny.ZERO).minus(_99_v0), new BigNumber((_dafny.Seq.UnicodeFromString("mvjqaslaa")).length)), _169_v46), _99_v0, _99_v0, _module.__default.fm1(_99_v0, false, new BigNumber(-216), globalState));
        _231_v100 = _nw25;
        _231_v100 = _231_v100;
        let _index23 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_102_v3).length));
        (_102_v3)[_index23] = new BigNumber(-865);
        let _index24 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_102_v3).length));
        (_102_v3)[_index24] = _99_v0;
        let _232_v101;
        _232_v101 = _dafny.Map.Empty.slice().updateUnsafe((_231_v100).f26,_99_v0);
        (globalState).f8 = (_232_v101).equals((_232_v101).Merge(_232_v101));
      } else {
        let _233_v102;
        _233_v102 = _module.D12.create_DC35(_99_v0);
        let _source5 = _233_v102;
        if (_source5.is_DC33) {
          let _234___mcc_h9 = (_source5).cf58;
          let _235_cf58 = _234___mcc_h9;
          r1 = _168_v45;
          let _236_v103;
          let _nw26 = new _module.C15();
          _nw26.__ctor(_168_v45, new BigNumber(747), false);
          _236_v103 = _nw26;
          let _237_v104;
          _237_v104 = _dafny.Map.Empty.slice().updateUnsafe(_168_v45,_236_v103);
          let _238_v105;
          _238_v105 = _dafny.Set.fromElements((_module.D0.create_DC3(new BigNumber((_237_v104).length), _236_v103.f11)).dtor_cf8, _168_v45, !(_236_v103.f11));
          (globalState).f8 = !(_238_v105).contains(!(true) || (true));
          _99_v0 = _235_cf58;
          _168_v45 = _module.__default.fm1(_99_v0, !(_236_v103.f11), _235_cf58, globalState);
        } else if (_source5.is_DC34) {
          let _239___mcc_h10 = (_source5).cf59;
          let _240_cf59 = _239___mcc_h10;
          let _241_v106;
          let _init4 = ((_242_v45) => function (_243_i6) {
            return _242_v45;
          })(_168_v45);
          let _nw27 = Array((new BigNumber(4)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw27.length); _i0_4++) {
            _nw27[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _241_v106 = _nw27;
          let _244_v107;
          _244_v107 = _dafny.Seq.of(_241_v106);
          _244_v107 = _dafny.Seq.Concat(_dafny.Seq.Concat(_244_v107, _244_v107), _dafny.Seq.of(_241_v106, _241_v106, _241_v106, _241_v106, _241_v106));
          let _245_v108;
          let _nw28 = new _module.C17();
          _nw28.__ctor(_168_v45);
          _245_v108 = _nw28;
          let _246_v109;
          let _nw29 = Array((new BigNumber(2)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _245_v108;
          _nw29[(_dafny.ONE).toNumber()] = _245_v108;
          _246_v109 = _nw29;
          let _247_v110;
          _247_v110 = _245_v108;
          let _index25 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_246_v109).length));
          (_246_v109)[_index25] = (_247_v110);
          let _index26 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_246_v109).length));
          (_246_v109)[_index26] = _245_v108;
          (globalState).f8 = _168_v45;
          _168_v45 = _168_v45;
        } else if (_source5.is_DC35) {
          let _248___mcc_h11 = (_source5).cf60;
          let _249_cf60 = _248___mcc_h11;
          let _250_v111;
          _250_v111 = _dafny.Seq.UnicodeFromString("gmrtyw");
          _249_cf60 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_250_v111, _250_v111), _dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), ((_251_v46) => function (_252_i7) {
            return _251_v46;
          })(_169_v46)))).length);
          _99_v0 = _99_v0;
          let _253_v112;
          let _init5 = ((_254_v45) => function (_255_i8) {
            return _254_v45;
          })(_168_v45);
          let _nw30 = Array((new BigNumber(2)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw30.length); _i0_5++) {
            _nw30[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _253_v112 = _nw30;
          let _256_v113;
          _256_v113 = _dafny.MultiSet.fromElements(_253_v112, _253_v112);
          let _257_v114;
          _257_v114 = _dafny.Map.Empty.slice().updateUnsafe((_256_v113).IsSubsetOf(_256_v113),_249_cf60);
          let _258_v115;
          _258_v115 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(_249_cf60,_168_v45));
          let _259_v116;
          let _nw31 = new _module.C10();
          _nw31.__ctor(true);
          _259_v116 = _nw31;
          let _260_v118;
          _260_v118 = _dafny.Map.Empty.slice().updateUnsafe(_259_v116,_module.__default.fm70(globalState));
          _257_v114 = _module.__default.fm69((_258_v115).Merge(_dafny.Map.Empty.slice().updateUnsafe(_168_v45,_dafny.Map.Empty.slice().updateUnsafe(_99_v0,_168_v45))), _module.__default.safeModuloInt(_99_v0, _99_v0), _249_cf60, !(_dafny.Map.Empty.slice().updateUnsafe(_259_v116,function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_170_v47).Elements) {
              let _261_v117 = _compr_37;
              if ((_170_v47).contains(_261_v117)) {
                _coll37.push([_261_v117,new BigNumber(343)]);
              }
            }
            return _coll37;
          }())).equals(_260_v118), globalState);
          (globalState).f8 = _168_v45;
        } else {
          let _262___mcc_h12 = (_source5).cf57;
          let _263_cf57 = _262___mcc_h12;
          let _264_v119;
          _264_v119 = _dafny.Seq.of(_168_v45, _168_v45, _168_v45, !(_168_v45), _168_v45);
          _264_v119 = _dafny.Seq.Concat(_dafny.Seq.Concat(_264_v119, _264_v119), _264_v119);
          let _265_v120;
          let _nw32 = Array((new BigNumber(27)).toNumber()).fill(false);
          _265_v120 = _nw32;
          let _266_v121;
          let _nw33 = new _module.C4();
          _nw33.__ctor(_99_v0, _168_v45);
          _266_v121 = _nw33;
          let _267_v122;
          _267_v122 = _dafny.Map.Empty.slice().updateUnsafe(((_168_v45) ? (_265_v120) : (_265_v120)),_266_v121);
          _267_v122 = (_267_v122).update(_265_v120, _266_v121);
          (globalState).f8 = false;
          let _index27 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_102_v3).length));
          (_102_v3)[_index27] = (_266_v121).f25;
          let _index28 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_102_v3).length));
          (_102_v3)[_index28] = _99_v0;
        }
        let _268_v123;
        _268_v123 = _dafny.Seq.of(_168_v45);
        r1 = !((_268_v123)[_module.__default.safeIndex((_dafny.ZERO).minus(_99_v0), new BigNumber((_268_v123).length))]);
        _268_v123 = _268_v123;
        let _269_v124;
        let _init6 = ((_270_v45) => function (_271_i9) {
          return _270_v45;
        })(_168_v45);
        let _nw34 = Array((new BigNumber(27)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw34.length); _i0_6++) {
          _nw34[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _269_v124 = _nw34;
        let _272_v125;
        _272_v125 = _dafny.Map.Empty.slice().updateUnsafe(_269_v124,_168_v45);
        let _273_v126;
        let _nw35 = new _module.C4();
        _nw35.__ctor(_99_v0, _168_v45);
        _273_v126 = _nw35;
        let _274_v127;
        _274_v127 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_273_v126,!(_168_v45)),_272_v125);
        let _275_v128;
        _275_v128 = _dafny.Map.Empty.slice().updateUnsafe(_273_v126,_273_v126.f11);
        let _276_v129;
        _276_v129 = _dafny.Seq.UnicodeFromString("lof");
        let _277_v130;
        _277_v130 = _dafny.Seq.of(_276_v129, _dafny.Seq.UnicodeFromString("eeksf"));
        let _278_v131;
        _278_v131 = _dafny.Map.Empty.slice().updateUnsafe(!((_273_v126).fm6(_273_v126.f11, new BigNumber((_276_v129).length), new BigNumber((_277_v130).length), globalState)),_dafny.Map.Empty.slice().updateUnsafe(_269_v124,_168_v45));
        let _279_v133;
        _279_v133 = _dafny.MultiSet.fromElements(new BigNumber(719));
        let _280_v134;
        let _nw36 = Array((new BigNumber(20)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = _272_v125;
        _nw36[(_dafny.ONE).toNumber()] = _272_v125;
        _nw36[(new BigNumber(2)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(3)).toNumber()] = (_272_v125).Merge(_272_v125);
        _nw36[(new BigNumber(4)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(5)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(6)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(7)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(8)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(9)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(10)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(11)).toNumber()] = (((_274_v127).contains((_275_v128).update(_273_v126, (_273_v126).fm6(_273_v126.f11, _99_v0, _99_v0, globalState)))) ? ((_274_v127).get((_275_v128).update(_273_v126, (_273_v126).fm6(_273_v126.f11, _99_v0, _99_v0, globalState)))) : (_272_v125));
        _nw36[(new BigNumber(12)).toNumber()] = (_272_v125).Merge(_272_v125);
        _nw36[(new BigNumber(13)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(14)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(15)).toNumber()] = _272_v125;
        _nw36[(new BigNumber(16)).toNumber()] = (_272_v125).Merge(_272_v125);
        _nw36[(new BigNumber(17)).toNumber()] = (((_278_v131).contains(false)) ? ((_278_v131).get(false)) : (_dafny.Map.Empty.slice().updateUnsafe(_269_v124,true)));
        _nw36[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_269_v124,(_273_v126).fm6(!(_168_v45), new BigNumber(((function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of (_279_v133).Elements) {
            let _281_v132 = _compr_38;
            if ((_279_v133).contains(_281_v132)) {
              _coll38.push([_module.__default.safeModuloInt(_281_v132, new BigNumber(167)),_273_v126.f11]);
            }
          }
          return _coll38;
        }()).update(_99_v0, false)).length), _99_v0, globalState));
        _nw36[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_269_v124,_273_v126.f11);
        _280_v134 = _nw36;
        let _index29 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_280_v134).length));
        (_280_v134)[_index29] = _272_v125;
        let _index30 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_280_v134).length));
        (_280_v134)[_index30] = (_272_v125).update(_269_v124, true);
        let _282_v135;
        _282_v135 = _dafny.MultiSet.fromElements(_269_v124, _269_v124, _269_v124, _269_v124, _269_v124);
        let _283_v136;
        let _nw37 = new _module.C12();
        _nw37.__ctor(_282_v135, (_99_v0).multipliedBy(_99_v0), _273_v126.f11, ((true) ? (_99_v0) : (_99_v0)), (_99_v0).minus(new BigNumber((_268_v123).length)));
        _283_v136 = _nw37;
      }
      let _284_v137;
      _284_v137 = _dafny.MultiSet.fromElements(_168_v45, ((_168_v45) ? (_168_v45) : (_168_v45)), _module.__default.fm1(_99_v0, _168_v45, new BigNumber(338), globalState));
      let _285_v138;
      _285_v138 = _dafny.Seq.UnicodeFromString("xyydy");
      let _286_v139;
      _286_v139 = _dafny.Seq.of(_168_v45, _168_v45);
      let _287_v140;
      _287_v140 = _dafny.MultiSet.fromElements(new BigNumber((_101_v2).length), new BigNumber(784), _99_v0, new BigNumber((_285_v138).length), new BigNumber((_286_v139).length));
      _284_v137 = _module.__default.fm30((_dafny.MultiSet.fromElements(new BigNumber(363), _99_v0)).IsDisjointFrom(_287_v140), _168_v45, _99_v0, globalState);
      let _288_v141;
      _288_v141 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(537),_168_v45);
      _288_v141 = (_288_v141).update(_99_v0, _168_v45);
      let _289_v142;
      let _nw38 = new _module.C2();
      _nw38.__ctor(_99_v0, _285_v138, _99_v0, _99_v0, _168_v45);
      _289_v142 = _nw38;
      let _290_v143;
      _290_v143 = _dafny.Map.Empty.slice().updateUnsafe(_289_v142,_168_v45);
      if ((_168_v45) === ((((_290_v143).contains(_289_v142)) ? ((_290_v143).get(_289_v142)) : (_168_v45)))) {
        let _291_v144;
        let _init7 = ((_292_v45) => function (_293_i10) {
          return _292_v45;
        })(_168_v45);
        let _nw39 = Array((new BigNumber(29)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw39.length); _i0_7++) {
          _nw39[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _291_v144 = _nw39;
        let _294_v145;
        _294_v145 = _dafny.MultiSet.fromElements(_291_v144);
        let _295_v146;
        _295_v146 = _module.D21.create_DC55(_294_v145);
        let _296_v147;
        _296_v147 = _module.D15.create_DC41(_module.__default.fm1((_289_v142).f12, _168_v45, (_289_v142).f12, globalState));
        let _297_v148;
        _297_v148 = _module.D15.create_DC42(_296_v147);
        let _298_v149;
        _298_v149 = _291_v144;
        let _299_v150;
        _299_v150 = _dafny.Map.Empty.slice().updateUnsafe((_289_v142).f12,(_289_v142).f13);
        let _rhs21 = _295_v146;
        let _rhs22 = (_298_v149);
        let _rhs23 = _297_v148;
        let _rhs24 = ((_289_v142).f13).multipliedBy((((_299_v150).contains((_289_v142).f13)) ? ((_299_v150).get((_289_v142).f13)) : ((_289_v142).f13)));
        _295_v146 = _rhs21;
        _291_v144 = _rhs22;
        _297_v148 = _rhs23;
        _99_v0 = _rhs24;
        _287_v140 = ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_99_v0, (_289_v142).f13))).update((_289_v142).f12, _module.__default.abs((_289_v142).f12))).Union(_module.__default.fm12((_dafny.ZERO).minus((_289_v142).f12), _99_v0, new BigNumber((_dafny.Set.fromElements(_168_v45, _168_v45)).length), _168_v45, globalState));
        _168_v45 = _168_v45;
        let _300_v151;
        let _nw40 = new _module.C3();
        _nw40.__ctor(((_168_v45) ? ((_289_v142).f12) : (new BigNumber(534))), new BigNumber((_100_v1).length));
        _300_v151 = _nw40;
        _99_v0 = _module.__default.safeDivisionInt(new BigNumber(-313), _module.__default.safeModuloInt((_289_v142).f13, (_289_v142).f13));
      } else {
        let _301_v152;
        let _nw41 = new _module.C3();
        _nw41.__ctor((_289_v142).f13, (_289_v142).f12);
        _301_v152 = _nw41;
        _301_v152 = _301_v152;
        let _302_v153;
        let _nw42 = Array((new BigNumber(20)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = (_289_v142).f20;
        _nw42[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((_289_v142).f20, _285_v138);
        _nw42[(new BigNumber(2)).toNumber()] = _285_v138;
        _nw42[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_285_v138, _285_v138);
        _nw42[(new BigNumber(4)).toNumber()] = (_289_v142).f20;
        _nw42[(new BigNumber(5)).toNumber()] = _285_v138;
        _nw42[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_285_v138, _module.__default.safeIndex((_289_v142).f12, new BigNumber((_285_v138).length)), _169_v46);
        _nw42[(new BigNumber(7)).toNumber()] = ((_168_v45) ? ((_289_v142).f20) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), ((_303_v46) => function (_304_i11) {
          return _303_v46;
        })(_169_v46))));
        _nw42[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat((_289_v142).f20, (_289_v142).f20);
        _nw42[(new BigNumber(9)).toNumber()] = (_289_v142).f20;
        _nw42[(new BigNumber(10)).toNumber()] = _285_v138;
        _nw42[(new BigNumber(11)).toNumber()] = (_289_v142).f20;
        _nw42[(new BigNumber(12)).toNumber()] = _285_v138;
        _nw42[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("sjuxrcwi");
        _nw42[(new BigNumber(14)).toNumber()] = _dafny.Seq.update((_289_v142).f20, _module.__default.safeIndex((_289_v142).f13, new BigNumber(((_289_v142).f20).length)), _169_v46);
        _nw42[(new BigNumber(15)).toNumber()] = (_289_v142).f20;
        _nw42[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_285_v138, (_289_v142).f20);
        _nw42[(new BigNumber(17)).toNumber()] = _285_v138;
        _nw42[(new BigNumber(18)).toNumber()] = (_289_v142).f20;
        _nw42[(new BigNumber(19)).toNumber()] = (_289_v142).f20;
        _302_v153 = _nw42;
        let _index31 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_302_v153).length));
        (_302_v153)[_index31] = _dafny.Seq.Concat((_289_v142).f20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), ((_305_v46) => function (_306_i12) {
          return (_module.D9.create_DC26(_305_v46)).dtor_cf48;
        })(_169_v46)));
        let _index32 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_302_v153).length));
        (_302_v153)[_index32] = (_289_v142).f20;
        if (((_module.D6.create_DC17(_168_v45, _99_v0)).dtor_cf30) && ((_168_v45) === (_168_v45))) {
          let _307_v154;
          _307_v154 = _dafny.Map.Empty.slice().updateUnsafe(_99_v0,(_289_v142).f12);
          let _308_v155;
          _308_v155 = _dafny.Set.fromElements(_100_v1, _dafny.Seq.of((((_307_v154).contains((_289_v142).f12)) ? ((_307_v154).get((_289_v142).f12)) : (new BigNumber(-953)))));
          _168_v45 = (_module.__default.fm71(_168_v45, _168_v45, (_289_v142).f13, _168_v45, globalState)).IsSubsetOf((_308_v155).Difference(_dafny.Set.fromElements(_dafny.Seq.of((_289_v142).f12, (_289_v142).f12, (_289_v142).f12, new BigNumber(268), (_289_v142).f12))));
          (globalState).f8 = _168_v45;
          let _309_v156;
          _309_v156 = _dafny.Map.Empty.slice().updateUnsafe(((_289_v142).f12).isLessThanOrEqualTo((_289_v142).f13),_289_v142);
          _309_v156 = (_309_v156).update(_168_v45, _289_v142);
          let _310_v157;
          _310_v157 = _dafny.Map.Empty.slice().updateUnsafe((_289_v142).f12,_dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), ((_311_v142) => function (_312_i13) {
            return (_311_v142).f13;
          })(_289_v142)));
          _310_v157 = (_310_v157).update(new BigNumber((_100_v1).length), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(507)), ((_313_v1) => function (_314_i14) {
            return new BigNumber((_313_v1).length);
          })(_100_v1)), _100_v1));
          (globalState).f8 = (_301_v152).fm7(globalState);
        } else {
          _99_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_100_v1)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_100_v1).length))]));
          _168_v45 = _168_v45;
          let _315_v158;
          _315_v158 = _dafny.Map.Empty.slice().updateUnsafe((_289_v142).f12,new _dafny.CodePoint('x'.codePointAt(0)));
          let _316_v159;
          _316_v159 = _module.D6.create_DC18(_168_v45);
          _315_v158 = (_315_v158).update((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_102_v3,(_316_v159).dtor_cf32)).length)).multipliedBy((_dafny.ZERO).minus((_289_v142).f12)), _169_v46);
          _168_v45 = _168_v45;
          let _317_v160;
          let _nw43 = new _module.C0();
          _nw43.__ctor(new BigNumber((_285_v138).length), (_168_v45) && (!(_168_v45)), (_289_v142).f13, (_289_v142).f13);
          _317_v160 = _nw43;
          _317_v160 = _317_v160;
        }
        let _318_v161;
        _318_v161 = _dafny.Map.Empty.slice().updateUnsafe(_168_v45,_168_v45);
        let _319_v162;
        _319_v162 = _dafny.MultiSet.fromElements(_318_v161, _318_v161);
        let _320_v163;
        _320_v163 = _319_v162;
        let _321_v165;
        let _nw44 = new _module.C1();
        _nw44.__ctor(new BigNumber((function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of (_170_v47).Elements) {
            let _322_v164 = _compr_39;
            if ((_170_v47).contains(_322_v164)) {
              _coll39.push([_322_v164,(_289_v142).f12]);
            }
          }
          return _coll39;
        }()).length));
        _321_v165 = _nw44;
        let _323_v166;
        let _init8 = ((_324_v1) => function (_325_i15) {
          return _324_v1;
        })(_100_v1);
        let _nw45 = Array((new BigNumber(29)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw45.length); _i0_8++) {
          _nw45[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _323_v166 = _nw45;
        let _index33 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_323_v166).length));
        (_323_v166)[_index33] = _100_v1;
        let _326_v167;
        let _nw46 = new _module.C7();
        _nw46.__ctor((_dafny.ZERO).minus(_99_v0), _99_v0);
        _326_v167 = _nw46;
        let _index34 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_323_v166).length));
        let _rhs25 = _320_v163;
        let _rhs26 = ((!(_168_v45) || (_168_v45)) ? (_321_v165) : (_321_v165));
        let _rhs27 = (_dafny.Seq.of((_289_v142).f12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_289_v142,_326_v167)).length),_168_v45)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_289_v142).f12,_99_v0)).length), (_289_v142).f12, (_289_v142).f13));
        let _lhs15 = _323_v166;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_323_v166).length));
        _320_v163 = _rhs25;
        _321_v165 = _rhs26;
        _lhs15[_lhs16] = _rhs27;
        _102_v3 = _102_v3;
      }
      let _327_v168;
      let _init9 = ((_328_v137) => function (_329_i16) {
        return _328_v137;
      })(_284_v137);
      let _nw47 = Array((new BigNumber(6)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw47.length); _i0_9++) {
        _nw47[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _327_v168 = _nw47;
      r0 = _327_v168;
      r1 = (!(_168_v45) || (true)) === (_168_v45);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _330_v0;
      _330_v0 = _dafny.Seq.UnicodeFromString("wh");
      let _331_v1;
      _331_v1 = new BigNumber(-230);
      let _332_v2;
      _332_v2 = _dafny.Seq.of(_331_v1);
      let _333_v3;
      _333_v3 = _dafny.Seq.of(_332_v2);
      let _334_v4;
      _334_v4 = true;
      let _335_v5;
      _335_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(19),_334_v4);
      let _336_v6;
      _336_v6 = new _dafny.CodePoint('l'.codePointAt(0));
      let _337_v7;
      _337_v7 = _dafny.Map.Empty.slice().updateUnsafe(_336_v6,_334_v4);
      let _338_v9;
      _338_v9 = _dafny.Set.fromElements(_336_v6);
      let _339_v10;
      _339_v10 = _dafny.Seq.of(function () {
        let _coll40 = new _dafny.Map();
        for (const _compr_40 of (_338_v9).Elements) {
          let _340_v8 = _compr_40;
          if ((_338_v9).contains(_340_v8)) {
            _coll40.push([_340_v8,true]);
          }
        }
        return _coll40;
      }());
      let _341_v11;
      _341_v11 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_336_v6,_334_v4), _337_v7, _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_334_v4), (_339_v10)[_module.__default.safeIndex(_331_v1, new BigNumber((_339_v10).length))]);
      let _342_globalState;
      let _nw48 = new _module.GlobalState();
      _nw48.__ctor(new BigNumber(886), new BigNumber(542), false, new BigNumber(596), _dafny.Seq.Concat(_330_v0, _330_v0), _dafny.Seq.update(_dafny.Seq.update((_333_v3)[_module.__default.safeIndex(_331_v1, new BigNumber((_333_v3).length))], _module.__default.safeIndex(_331_v1, new BigNumber(((_333_v3)[_module.__default.safeIndex(_331_v1, new BigNumber((_333_v3).length))]).length)), _331_v1), _module.__default.safeIndex(new BigNumber((_335_v5).length), new BigNumber((_dafny.Seq.update((_333_v3)[_module.__default.safeIndex(_331_v1, new BigNumber((_333_v3).length))], _module.__default.safeIndex(_331_v1, new BigNumber(((_333_v3)[_module.__default.safeIndex(_331_v1, new BigNumber((_333_v3).length))]).length)), _331_v1)).length)), _331_v1), new BigNumber(285), new BigNumber(542), true, _341_v11, new BigNumber(803));
      _342_globalState = _nw48;
      let _343_i0;
      _343_i0 = _dafny.ZERO;
      L0: {
        while (_334_v4) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_343_i0)) {
              break L0;
            }
            _343_i0 = (_343_i0).plus(_dafny.ONE);
            _331_v1 = _331_v1;
            let _344_v12;
            let _345_v13;
            let _out0;
            let _out1;
            let _outcollector0 = _module.__default.m0(_342_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _344_v12 = _out0;
            _345_v13 = _out1;
            let _346_v14;
            _346_v14 = _dafny.Seq.of(_334_v4, _345_v13, !(_345_v13), _334_v4, _345_v13);
            let _347_v15;
            let _nw49 = Array((new BigNumber(14)).toNumber());
            _nw49[(_dafny.ZERO).toNumber()] = _332_v2;
            _nw49[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_331_v1);
            _nw49[(new BigNumber(2)).toNumber()] = _332_v2;
            _nw49[(new BigNumber(3)).toNumber()] = _332_v2;
            _nw49[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), ((_348_v1) => function (_349_i1) {
              return _348_v1;
            })(_331_v1));
            _nw49[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_332_v2, _module.__default.safeIndex(_331_v1, new BigNumber((_332_v2).length)), _331_v1);
            _nw49[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-352)), ((_350_v1) => function (_351_i2) {
              return _350_v1;
            })(_331_v1));
            _nw49[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_332_v2, _332_v2);
            _nw49[(new BigNumber(8)).toNumber()] = _332_v2;
            _nw49[(new BigNumber(9)).toNumber()] = _332_v2;
            _nw49[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_332_v2, _module.__default.safeIndex(_331_v1, new BigNumber((_332_v2).length)), new BigNumber((_346_v14).length));
            _nw49[(new BigNumber(11)).toNumber()] = _332_v2;
            _nw49[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_332_v2, _module.__default.fm0(_342_globalState));
            _nw49[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_331_v1, new BigNumber((_330_v0).length), _331_v1, _331_v1, _331_v1);
            _347_v15 = _nw49;
            let _352_v16;
            _352_v16 = _module.D0.create_DC0(_347_v15);
            _347_v15 = (((_334_v4) ? (_352_v16) : (_352_v16))).dtor_cf0;
            _331_v1 = _331_v1;
          }
        }
      }
      _330_v0 = _330_v0;
      let _353_v17;
      _353_v17 = _dafny.Set.fromElements(_334_v4, _334_v4, _334_v4);
      let _354_v18;
      _354_v18 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_330_v0, _module.__default.safeIndex(new BigNumber((_353_v17).length), new BigNumber((_330_v0).length)), _336_v6));
      let _355_v19;
      _355_v19 = _dafny.Seq.of(_330_v0, _330_v0);
      let _356_v20;
      _356_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber(37), _334_v4, _331_v1, _342_globalState),_334_v4);
      let _357_v21;
      _357_v21 = _dafny.MultiSet.fromElements(_331_v1);
      let _358_v22;
      _358_v22 = _dafny.Seq.of(_334_v4);
      let _359_v23;
      let _nw50 = Array((new BigNumber(27)).toNumber());
      _nw50[(_dafny.ZERO).toNumber()] = _334_v4;
      _nw50[(_dafny.ONE).toNumber()] = _334_v4;
      _nw50[(new BigNumber(2)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(3)).toNumber()] = (_354_v18).IsDisjointFrom(_dafny.MultiSet.FromArray(_355_v19));
      _nw50[(new BigNumber(4)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(5)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(6)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(7)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(8)).toNumber()] = (((_356_v20).contains(_334_v4)) ? ((_356_v20).get(_334_v4)) : (_334_v4));
      _nw50[(new BigNumber(9)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(10)).toNumber()] = (_module.__default.fm1(_331_v1, _module.__default.fm1(_331_v1, _334_v4, _331_v1, _342_globalState), _331_v1, _342_globalState)) || (_334_v4);
      _nw50[(new BigNumber(11)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(12)).toNumber()] = (_334_v4) || (_334_v4);
      _nw50[(new BigNumber(13)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(14)).toNumber()] = true;
      _nw50[(new BigNumber(15)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(16)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(17)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(18)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(19)).toNumber()] = (_module.D0.create_DC3(new BigNumber((_357_v21).cardinality()), true)).dtor_cf8;
      _nw50[(new BigNumber(20)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(21)).toNumber()] = (_358_v22)[_module.__default.safeIndex(new BigNumber(-828), new BigNumber((_358_v22).length))];
      _nw50[(new BigNumber(22)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_330_v0, _330_v0);
      _nw50[(new BigNumber(23)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(24)).toNumber()] = (_331_v1).isLessThanOrEqualTo(_331_v1);
      _nw50[(new BigNumber(25)).toNumber()] = _334_v4;
      _nw50[(new BigNumber(26)).toNumber()] = (!(_334_v4)) || (_module.__default.fm1(_331_v1, _334_v4, _331_v1, _342_globalState));
      _359_v23 = _nw50;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_359_v23).length))) {
        let _360_i3 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_360_i3)) && ((_360_i3).isLessThan(new BigNumber((_359_v23).length))))) {
          (_359_v23)[(_360_i3)] = _334_v4;
        }
      }
      let _index35 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length));
      (_359_v23)[_index35] = _334_v4;
      let _index36 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length));
      (_359_v23)[_index36] = (_331_v1).isLessThanOrEqualTo((_module.__default.fm2(_342_globalState)).multipliedBy(_331_v1));
      let _361_v24;
      _361_v24 = _dafny.Map.Empty.slice().updateUnsafe(_332_v2,(_359_v23)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length))]);
      let _362_v25;
      _362_v25 = _dafny.Seq.of(_361_v24);
      (_342_globalState).f8 = !(!(!(_334_v4) || ((_dafny.Map.Empty.slice().updateUnsafe(_332_v2,(_359_v23)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length))])).equals((_362_v25)[_module.__default.safeIndex(_331_v1, new BigNumber((_362_v25).length))]))));
      _336_v6 = _336_v6;
      _331_v1 = (((_359_v23)[_module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length))]) ? (_331_v1) : (_331_v1));
      let _hi0 = _module.__default.safeDivisionInt(_331_v1, _331_v1);
      for (let _363_i4 = _331_v1; _363_i4.isLessThan(_hi0); _363_i4 = _363_i4.plus(_dafny.ONE)) {
        _331_v1 = new BigNumber(122);
        _358_v22 = _358_v22;
        let _364_v26;
        _364_v26 = _dafny.Set.fromElements(_331_v1, _module.__default.safeDivisionInt(_363_i4, _363_i4), _363_i4, new BigNumber((_330_v0).length));
        let _365_v27;
        _365_v27 = _dafny.Map.Empty.slice().updateUnsafe(_363_i4,_364_v26);
        _364_v26 = ((((_365_v27).contains(_331_v1)) ? ((_365_v27).get(_331_v1)) : (_dafny.Set.fromElements(_363_i4)))).Intersect(_364_v26);
        _334_v4 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_342_globalState),_331_v1)).length), _363_i4)).isLessThan(new BigNumber((_module.__default.fm3(_342_globalState)).length));
      }
      _331_v1 = _331_v1;
      _358_v22 = _358_v22;
      let _366_i5;
      _366_i5 = _dafny.ZERO;
      L1: {
        while (_334_v4) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_366_i5)) {
              break L1;
            }
            _366_i5 = (_366_i5).plus(_dafny.ONE);
            let _367_v28;
            _367_v28 = _module.D0.create_DC3(_331_v1, _334_v4);
            _353_v17 = _module.__default.fm4(new BigNumber(-348), _367_v28, _module.__default.safeModuloInt(_331_v1, _331_v1), new _dafny.CodePoint('o'.codePointAt(0)), _342_globalState);
            _330_v0 = _330_v0;
            let _index37 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_359_v23).length));
            (_359_v23)[_index37] = _334_v4;
            let _368_v29;
            let _nw51 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
            _368_v29 = _nw51;
            let _369_v30;
            _369_v30 = _dafny.Map.Empty.slice().updateUnsafe(_368_v29,(_331_v1).isLessThanOrEqualTo(new BigNumber(58)));
            _369_v30 = (_369_v30).update(_368_v29, !(_334_v4));
          }
        }
      }
      let _370_v31;
      let _nw52 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _370_v31 = _nw52;
      _370_v31 = _370_v31;
      (_342_globalState).f8 = (false) === (_334_v4);
      let _371_v32;
      let _nw53 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _371_v32 = _nw53;
      let _372_v33;
      let _init10 = ((_373_v2) => function (_374_i6) {
        return _373_v2;
      })(_332_v2);
      let _nw54 = Array((new BigNumber(8)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw54.length); _i0_10++) {
        _nw54[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _372_v33 = _nw54;
      let _375_v34;
      _375_v34 = _module.D0.create_DC0(_372_v33);
      let _376_v35;
      _376_v35 = _dafny.Seq.of(_375_v34);
      let _pat_let_tv0 = _372_v33;
      let _pat_let_tv1 = _372_v33;
      let _index38 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_371_v32).length));
      (_371_v32)[_index38] = _dafny.Seq.Concat(_376_v35, _dafny.Seq.of(function (_pat_let0_0) {
        return function (_377_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_378_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_378_dt__update_hcf0_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_375_v34), _375_v34, function (_pat_let2_0) {
        return function (_379_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_380_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_380_dt__update_hcf0_h1);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let2_0);
      }(_375_v34)));
      let _index39 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_371_v32).length));
      (_371_v32)[_index39] = _dafny.Seq.of(_module.D0.create_DC0(_372_v33), _375_v34);
      _331_v1 = (_dafny.ZERO).minus(((new BigNumber((_353_v17).length)).multipliedBy(_331_v1)).multipliedBy(_331_v1));
      let _381_v36;
      let _init11 = ((_382_v22) => function (_383_i7) {
        return _dafny.MultiSet.fromElements(_382_v22, _382_v22);
      })(_358_v22);
      let _nw55 = Array((new BigNumber(13)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw55.length); _i0_11++) {
        _nw55[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _381_v36 = _nw55;
      let _384_v37;
      _384_v37 = _dafny.MultiSet.fromElements(_358_v22);
      let _index40 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_381_v36).length));
      (_381_v36)[_index40] = _384_v37;
      let _index41 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_381_v36).length));
      let _rhs28 = _334_v4;
      let _rhs29 = _384_v37;
      let _rhs30 = _331_v1;
      let _lhs17 = _342_globalState;
      let _lhs18 = _381_v36;
      let _lhs19 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_381_v36).length));
      _lhs17.f8 = _rhs28;
      _lhs18[_lhs19] = _rhs29;
      _331_v1 = _rhs30;
      process.stdout.write((_330_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_331_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_332_v2, _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_333_v3, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-230))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_334_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(19),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_336_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_338_v9).equals(_dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_339_v10, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_341_v11).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),true), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_342_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_342_globalState.f5, _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_342_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_342_globalState).f9).equals(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),true), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_342_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_343_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_353_v17).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v18).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wl")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_355_v19, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wh"), _dafny.Seq.UnicodeFromString("wh")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_356_v20).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_357_v21).equals(_dafny.MultiSet.fromElements(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_358_v22, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v23)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_361_v24).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-230)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_362_v25, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-230)),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_366_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_371_v32)[new BigNumber(13)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[_dafny.ONE], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_372_v33)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[_dafny.ONE], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_375_v34).dtor_cf0)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(-230)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_376_v35).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_381_v36)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_384_v37).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true)))));
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
    static create_DC3(cf7, cf8) {
      let $dt = new D0(2);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + this.cf2.toVerbatimString(true) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
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
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false);
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
    static create_DC5(cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf9) {
      let $dt = new D1(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D1(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(_dafny.MultiSet.Empty, false);
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
    static create_DC7(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf15, cf16, cf17, cf18) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC10(cf19) {
      let $dt = new D3(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO, _dafny.Map.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC12(cf21, cf22, cf23) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC11(cf20) {
      let $dt = new D4(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC15(cf26, cf27, cf28) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D5(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf25 === other.cf25;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC17(cf30, cf31) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC18(cf32) {
      let $dt = new D6(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf33, cf34, cf35, cf36, cf37) {
      let $dt = new D6(2);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D6(3);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC20(cf38) {
      let $dt = new D6(4);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get is_DC20() { return this.$tag === 4; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && this.cf35 === other.cf35 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, _dafny.ZERO);
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
    static create_DC22(cf40, cf41) {
      let $dt = new D7(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf42, cf43, cf44, cf45) {
      let $dt = new D7(1);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC21(cf39) {
      let $dt = new D7(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC24(cf46) {
      let $dt = new D7(3);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC23" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22(_dafny.ZERO, []);
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
    static create_DC25(cf47) {
      let $dt = new D8(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf47) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC27(cf49, cf50, cf51, cf52) {
      let $dt = new D9(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC26(cf48) {
      let $dt = new D9(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC28(cf53) {
      let $dt = new D9(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC27(false, _module.D7.Default(), _dafny.ZERO, false);
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
    static create_DC29(cf54) {
      let $dt = new D10(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf56) {
      let $dt = new D11(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC30(cf55) {
      let $dt = new D11(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(false);
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
    static create_DC33(cf58) {
      let $dt = new D12(0);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC34(cf59) {
      let $dt = new D12(1);
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC35(cf60) {
      let $dt = new D12(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC32(cf57) {
      let $dt = new D12(3);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf59 === other.cf59;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(_dafny.ZERO);
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
    static create_DC37() {
      let $dt = new D13(0);
      return $dt;
    }
    static create_DC36(cf61) {
      let $dt = new D13(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC38(cf62) {
      let $dt = new D13(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC37";
      } else if (this.$tag === 1) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf62) + ")";
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
        return other.$tag === 1 && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC37();
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
    static create_DC39(cf63) {
      let $dt = new D14(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf63 === other.cf63;
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf65) {
      let $dt = new D15(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC40(cf64) {
      let $dt = new D15(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC42(cf66) {
      let $dt = new D15(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC41" + "(" + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf65 === other.cf65;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC41(false);
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
    static create_DC43(cf67) {
      let $dt = new D16(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
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
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC45(cf69, cf70) {
      let $dt = new D17(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC44(cf68) {
      let $dt = new D17(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC46(cf71) {
      let $dt = new D17(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC45(_dafny.ZERO, _dafny.Set.Empty);
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
    static create_DC48() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC49(cf73, cf74, cf75) {
      let $dt = new D18(1);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC47(cf72) {
      let $dt = new D18(2);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC48";
      } else if (this.$tag === 1) {
        return "D18.DC49" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf72) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf73, other.cf73) && _dafny.areEqual(this.cf74, other.cf74) && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC48();
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
    static create_DC51(cf77, cf78) {
      let $dt = new D19(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC50(cf76) {
      let $dt = new D19(1);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC50" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC51(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC53() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC52(cf79) {
      let $dt = new D20(1);
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC54(cf80) {
      let $dt = new D20(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC53";
      } else if (this.$tag === 1) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC54" + "(" + _dafny.toString(this.cf80) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC53();
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
    static create_DC56(cf82, cf83, cf84) {
      let $dt = new D21(0);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC55(cf81) {
      let $dt = new D21(1);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC57(cf85) {
      let $dt = new D21(2);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf82 === other.cf82 && this.cf83 === other.cf83 && _dafny.areEqual(this.cf84, other.cf84);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC56(false, false, _dafny.ZERO);
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
    static create_DC58(cf86) {
      let $dt = new D22(0);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC58" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf86 === other.cf86;
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
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC59(cf87) {
      let $dt = new D23(0);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC59" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf87, other.cf87);
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
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC61(cf89, cf90) {
      let $dt = new D24(0);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC60(cf88) {
      let $dt = new D24(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC61" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC60" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf89, other.cf89) && _dafny.areEqual(this.cf90, other.cf90);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf88, other.cf88);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC61(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC62(cf91) {
      let $dt = new D25(0);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC62" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf91 === other.cf91;
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
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC63(cf92) {
      let $dt = new D26(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC63() { return this.$tag === 0; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC63" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92);
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
          return D26.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f5 = _dafny.Seq.of();
      this.f8 = false;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.Seq.UnicodeFromString("");
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
      this._f9 = _dafny.Set.Empty;
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
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
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this.f29 = false;
      this._f28 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f28, f29, f12, f13) {
      let _this = this;
      (_this)._f28 = f28;
      (_this).f29 = f29;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return _this.f29;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((_module.D0.create_DC1(new BigNumber(68), _dafny.Seq.UnicodeFromString("aqdyj"), new BigNumber((_dafny.Set.fromElements(_this.f29)).length), _this.f29)).dtor_cf2, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fgplbwu"), _dafny.Seq.UnicodeFromString("boygyd")));
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f27) {
      let _this = this;
      (_this).f27 = f27;
      return;
    }
    fm34(globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(_this.f27, _this.f27)).minus(_this.f27);
    };
    fm35(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, true)), _dafny.MultiSet.fromElements(true)), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(false)), _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(false, true, false))));
    };
    m24(globalState) {
      let _this = this;
      let _385_v0;
      _385_v0 = _dafny.Seq.UnicodeFromString("mbgh");
      let _386_v1;
      _386_v1 = new _dafny.CodePoint('y'.codePointAt(0));
      let _387_v2;
      let _nw56 = Array((new BigNumber(28)).toNumber());
      _nw56[(_dafny.ZERO).toNumber()] = _385_v0;
      _nw56[(_dafny.ONE).toNumber()] = _385_v0;
      _nw56[(new BigNumber(2)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(3)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(4)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_385_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(931)), function (_388_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }));
      _nw56[(new BigNumber(6)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(7)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), function (_389_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("deoknfa"));
      _nw56[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_385_v0, _385_v0);
      _nw56[(new BigNumber(10)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(11)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(12)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(13)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(14)).toNumber()] = _module.__default.fm36(_this.f27, new BigNumber((_385_v0).length), _386_v1, globalState);
      _nw56[(new BigNumber(15)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(16)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(17)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(18)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("qscv");
      _nw56[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_385_v0, _385_v0);
      _nw56[(new BigNumber(21)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(22)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_385_v0, _385_v0);
      _nw56[(new BigNumber(24)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("beqc");
      _nw56[(new BigNumber(26)).toNumber()] = _385_v0;
      _nw56[(new BigNumber(27)).toNumber()] = _385_v0;
      _387_v2 = _nw56;
      let _index42 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
      (_387_v2)[_index42] = _385_v0;
      let _390_v3;
      _390_v3 = _dafny.Map.Empty.slice().updateUnsafe(_386_v1,_385_v0);
      let _391_v4;
      _391_v4 = true;
      let _index43 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
      (_387_v2)[_index43] = _dafny.Seq.Concat((((_390_v3).contains(_386_v1)) ? ((_390_v3).get(_386_v1)) : (_module.__default.fm36(_this.f27, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_391_v4)).length)), _386_v1, globalState))), _385_v0);
      let _392_v5;
      _392_v5 = _module.D4.create_DC13(_this.f27);
      _392_v5 = _392_v5;
      let _393_v6;
      let _nw57 = new _module.C0();
      _nw57.__ctor(_this.f27, _391_v4, _module.__default.safeModuloInt(_this.f27, _this.f27), _this.f27);
      _393_v6 = _nw57;
      let _394_v7;
      _394_v7 = _dafny.Map.Empty.slice().updateUnsafe(!(_391_v4),_393_v6.f29);
      if ((((_394_v7).contains(_393_v6.f29)) ? ((_394_v7).get(_393_v6.f29)) : (((_393_v6).f28).isLessThanOrEqualTo((_393_v6).f28)))) {
        let _395_v8;
        _395_v8 = _dafny.Seq.of(_393_v6.f29, _393_v6.f29);
        let _396_v9;
        _396_v9 = _dafny.Seq.of(new BigNumber((_395_v8).length), new BigNumber(((_394_v7).update(_391_v4, _393_v6.f29)).length));
        (_this).f27 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_396_v9, _396_v9)).length), (_393_v6).f28);
        let _index44 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
        (_387_v2)[_index44] = _385_v0;
        let _397_v10;
        let _init12 = ((_398_v6) => function (_399_i2) {
          return (_399_i2).plus((_398_v6).f28);
        })(_393_v6);
        let _nw58 = Array((new BigNumber(11)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw58.length); _i0_12++) {
          _nw58[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _397_v10 = _nw58;
        let _index45 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length));
        (_397_v10)[_index45] = _this.f27;
        let _index46 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length));
        (_397_v10)[_index46] = (_396_v9)[_module.__default.safeIndex(new BigNumber((_395_v8).length), new BigNumber((_396_v9).length))];
        let _400_v11;
        _400_v11 = _dafny.Map.Empty.slice().updateUnsafe(_395_v8,_393_v6.f29);
        _400_v11 = (_400_v11).update(_395_v8, _391_v4);
        if (_393_v6.f29) {
          let _index47 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
          (_387_v2)[_index47] = (_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))];
          (_393_v6).f29 = _393_v6.f29;
          let _index48 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
          (_387_v2)[_index48] = _dafny.Seq.Concat(_385_v0, (_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))]);
          let _401_v12;
          let _init13 = ((_402_v6) => function (_403_i3) {
            return _402_v6.f29;
          })(_393_v6);
          let _nw59 = Array((new BigNumber(6)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw59.length); _i0_13++) {
            _nw59[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _401_v12 = _nw59;
          let _404_v13;
          _404_v13 = _dafny.MultiSet.fromElements((_391_v4) && (_393_v6.f29));
          let _405_v14;
          _405_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(416),_this.f27);
          let _406_v15;
          _406_v15 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_module.D3.create_DC9((_393_v6).f28, _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_397_v10)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length))]),_405_v14), (_393_v6).f28, (_397_v10)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length))])).dtor_cf17));
          let _407_v16;
          _407_v16 = _dafny.Map.Empty.slice().updateUnsafe(_406_v15,(_393_v6).f28);
          let _rhs31 = _401_v12;
          let _rhs32 = _401_v12;
          let _rhs33 = (_dafny.MultiSet.fromElements(_391_v4)).Union((_404_v13).Difference(_dafny.MultiSet.FromArray(_395_v8)));
          let _rhs34 = (_385_v0)[_module.__default.safeIndex(((_393_v6).f28).plus(new BigNumber((function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of (_407_v16).Keys.Elements) {
              let _408_v17 = _compr_41;
              if ((_407_v16).contains(_408_v17)) {
                _coll41.add(_408_v17);
              }
            }
            return _coll41;
          }()).length)), new BigNumber((_385_v0).length))];
          _401_v12 = _rhs31;
          _401_v12 = _rhs32;
          _404_v13 = _rhs33;
          _386_v1 = _rhs34;
          (_this).f27 = new BigNumber(-182);
        } else {
          let _409_v18;
          _409_v18 = _dafny.Set.fromElements(_393_v6.f29);
          let _410_v19;
          _410_v19 = _dafny.Seq.of(_409_v18);
          let _411_v20;
          _411_v20 = _dafny.Map.Empty.slice().updateUnsafe(_409_v18,_393_v6.f29);
          let _412_v21;
          _412_v21 = _dafny.MultiSet.fromElements((_397_v10)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length))], new BigNumber((_411_v20).length));
          let _413_v22;
          let _nw60 = Array((new BigNumber(11)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = false;
          _nw60[(_dafny.ONE).toNumber()] = _393_v6.f29;
          _nw60[(new BigNumber(2)).toNumber()] = ((_410_v19)[_module.__default.safeIndex((_393_v6).f28, new BigNumber((_410_v19).length))]).IsSubsetOf(_409_v18);
          _nw60[(new BigNumber(3)).toNumber()] = (_409_v18).equals(_409_v18);
          _nw60[(new BigNumber(4)).toNumber()] = (_409_v18).IsSubsetOf(_409_v18);
          _nw60[(new BigNumber(5)).toNumber()] = !(_module.__default.fm1((_393_v6).f28, _393_v6.f29, (_397_v10)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length))], globalState));
          _nw60[(new BigNumber(6)).toNumber()] = _391_v4;
          _nw60[(new BigNumber(7)).toNumber()] = (_412_v21).IsProperSubsetOf(_412_v21);
          _nw60[(new BigNumber(8)).toNumber()] = _393_v6.f29;
          _nw60[(new BigNumber(9)).toNumber()] = _393_v6.f29;
          _nw60[(new BigNumber(10)).toNumber()] = !(!(_391_v4));
          _413_v22 = _nw60;
          let _index49 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_413_v22).length));
          (_413_v22)[_index49] = !((_393_v6).f28).isEqualTo((_393_v6).f28);
          let _index50 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_413_v22).length));
          let _rhs35 = !(((_397_v10)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_397_v10).length))]).isLessThanOrEqualTo((_393_v6).f28));
          let _rhs36 = _this.f27;
          let _lhs20 = _413_v22;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_413_v22).length));
          let _lhs22 = _this;
          _lhs20[_lhs21] = _rhs35;
          _lhs22.f27 = _rhs36;
          let _414_v23;
          let _nw61 = Array((new BigNumber(25)).toNumber()).fill([]);
          _414_v23 = _nw61;
          _414_v23 = _414_v23;
          _397_v10 = _397_v10;
          _387_v2 = _387_v2;
          let _415_v24;
          _415_v24 = _dafny.Map.Empty.slice().updateUnsafe(_393_v6.f29,(_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))]);
          let _rhs37 = !((_415_v24).Merge(_415_v24)).contains(_391_v4);
          let _rhs38 = _393_v6.f29;
          let _lhs23 = globalState;
          _lhs23.f8 = _rhs37;
          _391_v4 = _rhs38;
        }
      } else {
        let _416_v25;
        let _nw62 = Array((new BigNumber(6)).toNumber()).fill(false);
        _416_v25 = _nw62;
        let _index51 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_416_v25).length));
        (_416_v25)[_index51] = (_393_v6).fm7(globalState);
        let _417_v26;
        _417_v26 = _module.D0.create_DC3(new BigNumber((_dafny.Seq.of(_393_v6.f29, true)).length), false);
        let _418_v27;
        _418_v27 = _dafny.Map.Empty.slice().updateUnsafe(_417_v26,_393_v6.f29);
        let _419_v28;
        _419_v28 = _dafny.Seq.of((_393_v6).f28);
        let _420_v29;
        _420_v29 = _dafny.Seq.of((((_418_v27).contains(_module.D0.create_DC3(new BigNumber((_dafny.Seq.update((_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))], _module.__default.safeIndex(_this.f27, new BigNumber(((_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))]).length)), _386_v1)).length), _module.__default.fm1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_419_v28).length))).length), _393_v6.f29, (_393_v6).f28, globalState)))) ? ((_418_v27).get(_module.D0.create_DC3(new BigNumber((_dafny.Seq.update((_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))], _module.__default.safeIndex(_this.f27, new BigNumber(((_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))]).length)), _386_v1)).length), _module.__default.fm1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_419_v28).length))).length), _393_v6.f29, (_393_v6).f28, globalState)))) : (_393_v6.f29)), _391_v4, _393_v6.f29);
        let _index52 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_416_v25).length));
        (_416_v25)[_index52] = (_420_v29)[_module.__default.safeIndex((_this.f27).plus((_393_v6).f28), new BigNumber((_420_v29).length))];
        if (_dafny.Seq.IsPrefixOf((_387_v2)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length))], _385_v0)) {
          let _421_v30;
          _421_v30 = _dafny.Map.Empty.slice().updateUnsafe(_393_v6.f29,(_393_v6).f28);
          let _422_v31;
          let _nw63 = new _module.C0();
          _nw63.__ctor((_393_v6).f28, _391_v4, (_393_v6).f28, new BigNumber(((_421_v30).Merge(_421_v30)).length));
          _422_v31 = _nw63;
          _416_v25 = _416_v25;
          let _423_v32;
          let _init14 = ((_424_v31) => function (_425_i4) {
            return (_425_i4).multipliedBy((_424_v31).f28);
          })(_422_v31);
          let _nw64 = Array((new BigNumber(14)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw64.length); _i0_14++) {
            _nw64[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _423_v32 = _nw64;
          let _426_v33;
          _426_v33 = _dafny.Map.Empty.slice().updateUnsafe(_393_v6.f29,_423_v32);
          let _427_v34;
          _427_v34 = _dafny.Map.Empty.slice().updateUnsafe(_426_v33,_416_v25);
          _427_v34 = (_427_v34).update(_426_v33, _416_v25);
          let _428_v35;
          let _nw65 = Array((new BigNumber(4)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_420_v29, _dafny.Seq.of((_416_v25)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_416_v25).length))]));
          _nw65[(_dafny.ONE).toNumber()] = _dafny.Seq.of(false);
          _nw65[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_420_v29, _module.__default.safeIndex(_this.f27, new BigNumber((_420_v29).length)), true);
          _nw65[(new BigNumber(3)).toNumber()] = _420_v29;
          _428_v35 = _nw65;
          let _index53 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_428_v35).length));
          (_428_v35)[_index53] = _dafny.Seq.Concat(_420_v29, _420_v29);
          let _429_v36;
          _429_v36 = _dafny.Set.fromElements(_391_v4);
          let _index54 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_428_v35).length));
          let _index55 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_416_v25).length));
          let _rhs39 = ((_393_v6).f28).minus((_393_v6).f28);
          let _rhs40 = _dafny.Seq.Concat(_420_v29, _dafny.Seq.Concat(_420_v29, _420_v29));
          let _rhs41 = (_393_v6.f29) || (_module.__default.fm1(_this.f27, false, new BigNumber((_429_v36).length), globalState));
          let _lhs24 = _this;
          let _lhs25 = _428_v35;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_428_v35).length));
          let _lhs27 = _416_v25;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_416_v25).length));
          _lhs24.f27 = _rhs39;
          _lhs25[_lhs26] = _rhs40;
          _lhs27[_lhs28] = _rhs41;
          let _430_v37;
          _430_v37 = _dafny.MultiSet.fromElements((_422_v31).f28);
          _430_v37 = (_430_v37).Difference(((_module.D1.create_DC5(_dafny.MultiSet.fromElements((_393_v6).f28), _391_v4)).dtor_cf10).Difference(_dafny.MultiSet.fromElements((_393_v6).f28)));
        } else {
          let _431_v38;
          let _init15 = ((_432_v6) => function (_433_i5) {
            return _module.__default.safeModuloInt(_433_i5, (_432_v6).f28);
          })(_393_v6);
          let _nw66 = Array((new BigNumber(24)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw66.length); _i0_15++) {
            _nw66[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _431_v38 = _nw66;
          let _434_v39;
          _434_v39 = _module.D0.create_DC2(new BigNumber((_419_v28).length), _431_v38);
          (_this).f27 = (_434_v39).dtor_cf5;
          let _435_v40;
          _435_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,(_393_v6).f28);
          let _pat_let_tv2 = _435_v40;
          let _436_v41;
          _436_v41 = _dafny.Seq.of(function (_pat_let4_0) {
            return function (_437_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_438_dt__update_hcf20_h0) {
                  return _module.D4.create_DC11(_438_dt__update_hcf20_h0);
                }(_pat_let5_0);
              }(_pat_let_tv2);
            }(_pat_let4_0);
          }(_module.D4.create_DC11(_435_v40)));
          let _439_v42;
          _439_v42 = _dafny.MultiSet.fromElements(_393_v6.f29);
          let _440_v43;
          _440_v43 = _module.D4.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_439_v42).cardinality()),(_393_v6).f28));
          _436_v41 = _dafny.Seq.Concat(_dafny.Seq.Concat(_436_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), ((_441_v40) => function (_442_i6) {
            return _module.D4.create_DC11(_441_v40);
          })(_435_v40))), _dafny.Seq.Concat(_436_v41, _dafny.Seq.update(_436_v41, _module.__default.safeIndex(_this.f27, new BigNumber((_436_v41).length)), _440_v43)));
          let _443_v44;
          let _nw67 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _443_v44 = _nw67;
          let _444_v45;
          _444_v45 = _dafny.Map.Empty.slice().updateUnsafe(_420_v29,(_393_v6).f28);
          let _index56 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_443_v44).length));
          (_443_v44)[_index56] = _444_v45;
          let _index57 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_443_v44).length));
          (_443_v44)[_index57] = (_444_v45).Merge(_444_v45);
          let _445_v46;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_this.f27, _393_v6.f29, new BigNumber((_419_v28).length), new BigNumber(-556));
          _445_v46 = _nw68;
          (_this).f27 = (new BigNumber(-286)).minus((_this.f27).plus((_393_v6).f28));
        }
        let _446_v47;
        _446_v47 = _module.D0.create_DC1((_393_v6).f28, _dafny.Seq.UnicodeFromString("qbumnuk"), new BigNumber(285), _391_v4);
        let _index58 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_387_v2).length));
        (_387_v2)[_index58] = _dafny.Seq.update(_dafny.Seq.Concat(_385_v0, _dafny.Seq.Concat((_446_v47).dtor_cf2, _385_v0)), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_447_v6) => function (_448_i7) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_447_v6).f28,_447_v6.f29)).length);
        })(_393_v6))).length))).length), new BigNumber((_dafny.Seq.Concat(_385_v0, _dafny.Seq.Concat((_446_v47).dtor_cf2, _385_v0))).length)), _386_v1);
        let _449_v48;
        _449_v48 = _dafny.Set.fromElements(_this.f27, new BigNumber(921), _this.f27, (_this.f27).multipliedBy((_393_v6).f28));
        _449_v48 = _449_v48;
        let _450_v49;
        let _init16 = function (_451_i8) {
          return (_451_i8).plus(_this.f27);
        };
        let _nw69 = Array((new BigNumber(28)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw69.length); _i0_16++) {
          _nw69[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _450_v49 = _nw69;
        let _452_v50;
        _452_v50 = _dafny.Seq.of(_450_v49, _450_v49, _450_v49);
        let _453_v51;
        _453_v51 = _module.D1.create_DC4(_452_v50);
        let _454_v52;
        _454_v52 = _dafny.Map.Empty.slice().updateUnsafe((_393_v6).f28,_453_v51);
        _454_v52 = (_454_v52).update((_393_v6).f28, _453_v51);
      }
      let _455_v53;
      _455_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,_393_v6.f29);
      let _456_v54;
      _456_v54 = _dafny.Set.fromElements(_393_v6.f29);
      let _457_v55;
      _457_v55 = _module.D0.create_DC1(new BigNumber((_455_v53).length), _385_v0, new BigNumber((_456_v54).length), _393_v6.f29);
      let _458_v56;
      _458_v56 = _dafny.MultiSet.fromElements(_this.f27);
      let _459_v57;
      _459_v57 = _dafny.Seq.of(_458_v56, _458_v56);
      let _pat_let_tv3 = _385_v0;
      let _pat_let_tv4 = _387_v2;
      let _pat_let_tv5 = _387_v2;
      let _pat_let_tv6 = _387_v2;
      let _pat_let_tv7 = _387_v2;
      let _pat_let_tv8 = _385_v0;
      let _pat_let_tv9 = _387_v2;
      let _pat_let_tv10 = _387_v2;
      let _pat_let_tv11 = _387_v2;
      let _pat_let_tv12 = _387_v2;
      _385_v0 = _dafny.Seq.update(function (_source6) {
        if (_source6.is_DC1) {
          let _460___mcc_h0 = (_source6).cf1;
          let _461___mcc_h1 = (_source6).cf2;
          let _462___mcc_h2 = (_source6).cf3;
          let _463___mcc_h3 = (_source6).cf4;
          let _464_cf4 = _463___mcc_h3;
          let _465_cf3 = _462___mcc_h2;
          let _466_cf2 = _461___mcc_h1;
          let _467_cf1 = _460___mcc_h0;
          return _pat_let_tv3;
        } else if (_source6.is_DC2) {
          let _468___mcc_h4 = (_source6).cf5;
          let _469___mcc_h5 = (_source6).cf6;
          let _470_cf6 = _469___mcc_h5;
          let _471_cf5 = _468___mcc_h4;
          return (_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_pat_let_tv4).length))];
        } else if (_source6.is_DC3) {
          let _472___mcc_h6 = (_source6).cf7;
          let _473___mcc_h7 = (_source6).cf8;
          let _474_cf8 = _473___mcc_h7;
          let _475_cf7 = _472___mcc_h6;
          return _dafny.Seq.UnicodeFromString("a");
        } else {
          let _476___mcc_h8 = (_source6).cf0;
          let _477_cf0 = _476___mcc_h8;
          return (_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_pat_let_tv6).length))];
        }
      }(_457_v55), _module.__default.safeIndex((new BigNumber(((_458_v56).update(_this.f27, _module.__default.abs(new BigNumber((_dafny.MultiSet.FromArray(_459_v57)).cardinality())))).cardinality())).multipliedBy((_393_v6).f28), new BigNumber((function (_source7) {
        if (_source7.is_DC1) {
          let _478___mcc_h9 = (_source7).cf1;
          let _479___mcc_h10 = (_source7).cf2;
          let _480___mcc_h11 = (_source7).cf3;
          let _481___mcc_h12 = (_source7).cf4;
          let _482_cf4 = _481___mcc_h12;
          let _483_cf3 = _480___mcc_h11;
          let _484_cf2 = _479___mcc_h10;
          let _485_cf1 = _478___mcc_h9;
          return _pat_let_tv8;
        } else if (_source7.is_DC2) {
          let _486___mcc_h13 = (_source7).cf5;
          let _487___mcc_h14 = (_source7).cf6;
          let _488_cf6 = _487___mcc_h14;
          let _489_cf5 = _486___mcc_h13;
          return (_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_pat_let_tv9).length))];
        } else if (_source7.is_DC3) {
          let _490___mcc_h15 = (_source7).cf7;
          let _491___mcc_h16 = (_source7).cf8;
          let _492_cf8 = _491___mcc_h16;
          let _493_cf7 = _490___mcc_h15;
          return _dafny.Seq.UnicodeFromString("a");
        } else {
          let _494___mcc_h17 = (_source7).cf0;
          let _495_cf0 = _494___mcc_h17;
          return (_pat_let_tv12)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_pat_let_tv11).length))];
        }
      }(_457_v55)).length)), _386_v1);
      let _496_v58;
      let _nw70 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _496_v58 = _nw70;
      let _index59 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_496_v58).length));
      (_496_v58)[_index59] = (_393_v6).f28;
      let _497_v59;
      let _init17 = ((_498_v6) => function (_499_i9) {
        return (_498_v6.f29) || (_498_v6.f29);
      })(_393_v6);
      let _nw71 = Array((new BigNumber(17)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw71.length); _i0_17++) {
        _nw71[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _497_v59 = _nw71;
      let _500_v60;
      _500_v60 = _dafny.Seq.of((_455_v53).equals(_455_v53), _dafny.Seq.IsProperPrefixOf(_385_v0, _385_v0), _391_v4, !(_393_v6.f29) || (_391_v4), !(_this.f27).isEqualTo(new BigNumber(971)));
      let _index60 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_496_v58).length));
      let _rhs42 = new BigNumber((_500_v60).length);
      let _rhs43 = _497_v59;
      let _rhs44 = (_dafny.ZERO).minus((_393_v6).f28);
      let _lhs29 = _496_v58;
      let _lhs30 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_496_v58).length));
      let _lhs31 = _this;
      _lhs29[_lhs30] = _rhs42;
      _497_v59 = _rhs43;
      _lhs31.f27 = _rhs44;
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T0, _module.T1];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f26, f20, f12, f13, f11) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f20 = f20;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f11 = f11;
      return;
    }
    fm18(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.MultiSet.fromElements(_this.f11, _this.f11)).Intersect(_dafny.MultiSet.fromElements(_this.f11, _this.f11))).cardinality())).isLessThan((_this).f12);
    };
    fm7(globalState) {
      let _this = this;
      return _this.f11;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((_this).f20, _dafny.Seq.Concat((_this).f20, (_this).f20));
    };
    fm5(p0, globalState) {
      let _this = this;
      let _source8 = ((false) ? (_module.D1.create_DC5(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f12, (_this).f26, (_this).f13)), _this.f11)) : (_module.D1.create_DC5(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f13)).length), new BigNumber(-86)), false)));
      if (_source8.is_DC5) {
        let _501___mcc_h0 = (_source8).cf10;
        let _502___mcc_h1 = (_source8).cf11;
        let _503_cf11 = _502___mcc_h1;
        let _504_cf10 = _501___mcc_h0;
        return _module.D0.create_DC1(new BigNumber((function () {
  let _coll42 = new _dafny.Map();
  for (const _compr_42 of _dafny.IntegerRange(new BigNumber(129), new BigNumber(-59))) {
    let _505_v0 = _compr_42;
    if (((new BigNumber(129)).isLessThanOrEqualTo(_505_v0)) && ((_505_v0).isLessThan(new BigNumber(-59)))) {
      _coll42.push([(_505_v0).multipliedBy(new BigNumber(750)),_this.f11]);
    }
  }
  return _coll42;
}()).length), _dafny.Seq.UnicodeFromString("nthyvsrum"), new BigNumber(788), _503_cf11);
      } else if (_source8.is_DC4) {
        let _506___mcc_h2 = (_source8).cf9;
        let _507_cf9 = _506___mcc_h2;
        return _module.D0.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("au")).length), (_this).f20, (_this).f26, _this.f11);
      } else {
        let _508___mcc_h3 = (_source8).cf12;
        let _509_cf12 = _508___mcc_h3;
        return _module.D0.create_DC1((_this).f26, (_this).f20, new BigNumber(575), !(_this.f11));
      }
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return !(_this.f11);
    };
    fm32(p0, globalState) {
      let _this = this;
      let _source9 = _module.D0.create_DC3(new BigNumber((_dafny.Seq.update((_this).f20, _module.__default.safeIndex((_this).f12, new BigNumber(((_this).f20).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length), _this.f11);
      if (_source9.is_DC1) {
        let _510___mcc_h0 = (_source9).cf1;
        let _511___mcc_h1 = (_source9).cf2;
        let _512___mcc_h2 = (_source9).cf3;
        let _513___mcc_h3 = (_source9).cf4;
        let _514_cf4 = _513___mcc_h3;
        let _515_cf3 = _512___mcc_h2;
        let _516_cf2 = _511___mcc_h1;
        let _517_cf1 = _510___mcc_h0;
        if (_514_cf4) {
          return _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of _dafny.IntegerRange(new BigNumber(65), new BigNumber(338))) {
              let _518_v0 = _compr_43;
              if (((new BigNumber(65)).isLessThanOrEqualTo(_518_v0)) && ((_518_v0).isLessThan(new BigNumber(338)))) {
                _coll43.push([_module.__default.safeModuloInt(_518_v0, (_this).f26),(_this).f26]);
              }
            }
            return _coll43;
          }(),_this.f11);
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_515_cf3,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f12)).length)),_this.f11);
        }
      } else if (_source9.is_DC2) {
        let _519___mcc_h4 = (_source9).cf5;
        let _520___mcc_h5 = (_source9).cf6;
        let _521_cf6 = _520___mcc_h5;
        let _522_cf5 = _519___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true, _this.f11, true, _this.f11, _this.f11)).cardinality()),new BigNumber(-522)),_this.f11);
      } else if (_source9.is_DC3) {
        let _523___mcc_h6 = (_source9).cf7;
        let _524___mcc_h7 = (_source9).cf8;
        let _525_cf8 = _524___mcc_h7;
        let _526_cf7 = _523___mcc_h6;
        return function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of (function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_526_cf7,(_this).f13))).Elements) {
              let _527_v2 = _compr_45;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_526_cf7,(_this).f13)), _527_v2)) {
                _coll45.push([_527_v2,false]);
              }
            }
            return _coll45;
          }()).Keys.Elements) {
            let _528_v1 = _compr_44;
            if ((function () {
              let _coll46 = new _dafny.Map();
              for (const _compr_46 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_526_cf7,(_this).f13))).Elements) {
                let _527_v2 = _compr_46;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_526_cf7,(_this).f13)), _527_v2)) {
                  _coll46.push([_527_v2,false]);
                }
              }
              return _coll46;
            }()).contains(_528_v1)) {
              _coll44.push([_528_v1,_this.f11]);
            }
          }
          return _coll44;
        }();
      } else {
        let _529___mcc_h8 = (_source9).cf0;
        let _530_cf0 = _529___mcc_h8;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f12),_this.f11)).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(501))) {
            let _531_v3 = _compr_47;
            if (((new BigNumber(809)).isLessThanOrEqualTo(_531_v3)) && ((_531_v3).isLessThan(new BigNumber(501)))) {
              _coll47.push([_module.__default.safeDivisionInt(_531_v3, (_this).f26),(_this).f13]);
            }
          }
          return _coll47;
        }(),_this.f11));
      }
    };
    fm33(p0, globalState) {
      let _this = this;
      let _source10 = _module.D3.create_DC10(_module.D3.create_DC10(_module.D3.create_DC9((_this).f13, _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_module.D4.create_DC11(_dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f26))).dtor_cf20), (_this).f26, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11)).length))));
      if (_source10.is_DC9) {
        let _532___mcc_h0 = (_source10).cf15;
        let _533___mcc_h1 = (_source10).cf16;
        let _534___mcc_h2 = (_source10).cf17;
        let _535___mcc_h3 = (_source10).cf18;
        let _536_cf18 = _535___mcc_h3;
        let _537_cf17 = _534___mcc_h2;
        let _538_cf16 = _533___mcc_h1;
        let _539_cf15 = _532___mcc_h0;
        return _this.f11;
      } else if (_source10.is_DC8) {
        let _540___mcc_h4 = (_source10).cf14;
        let _541_cf14 = _540___mcc_h4;
        return _this.f11;
      } else {
        let _542___mcc_h5 = (_source10).cf19;
        let _543_cf19 = _542___mcc_h5;
        return !(true);
      }
    };
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      if (_this.f11) {
        let _544_v0;
        let _nw72 = new _module.C1();
        _nw72.__ctor(p1);
        _544_v0 = _nw72;
        (globalState).f8 = _this.f11;
        let _545_v1;
        let _nw73 = Array((new BigNumber(16)).toNumber()).fill(false);
        _545_v1 = _nw73;
        let _546_v2;
        _546_v2 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f11) ? (_this.f11) : (_this.f11)),_545_v1);
        _546_v2 = (_546_v2).update(_this.f11, _545_v1);
        let _547_v3;
        _547_v3 = _dafny.Set.fromElements(_this.f11);
        let _548_v4;
        _548_v4 = _dafny.Set.fromElements((_this).f12);
        let _549_v5;
        _549_v5 = _dafny.Seq.of(new BigNumber((_548_v4).length), p1);
        let _550_v6;
        _550_v6 = _dafny.MultiSet.fromElements(p1);
        _547_v3 = (((_this).fm6(_this.f11, (_549_v5)[_module.__default.safeIndex(new BigNumber(324), new BigNumber((_549_v5).length))], new BigNumber((_550_v6).cardinality()), globalState)) ? ((_547_v3).Difference(_dafny.Set.fromElements(_this.f11))) : (_dafny.Set.fromElements(_this.f11)));
        (_this).f11 = (_this).fm33(((_this).f12).minus(p0), globalState);
      } else {
        let _551_v7;
        _551_v7 = new _dafny.CodePoint('s'.codePointAt(0));
        let _552_v8;
        let _init18 = function (_553_i0) {
          return _module.__default.safeDivisionInt(_553_i0, (_this).f13);
        };
        let _nw74 = Array((new BigNumber(13)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw74.length); _i0_18++) {
          _nw74[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _552_v8 = _nw74;
        let _rhs45 = ((p2).isLessThanOrEqualTo((_this).f12)) === (_this.f11);
        let _rhs46 = ((!(_this.f11)) ? (new _dafny.CodePoint('i'.codePointAt(0))) : (_551_v7));
        let _rhs47 = p3;
        let _lhs32 = _this;
        _lhs32.f11 = _rhs45;
        _551_v7 = _rhs46;
        _552_v8 = _rhs47;
        let _554_v9;
        _554_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(_this.f11,p1));
        let _rhs48 = _module.__default.safeModuloInt(p2, (_dafny.ZERO).minus(p0));
        let _rhs49 = (_554_v9).Merge(_554_v9);
        r1 = _rhs48;
        _554_v9 = _rhs49;
        (_this).f11 = true;
        let _555_v10;
        _555_v10 = _dafny.Map.Empty.slice().updateUnsafe(_551_v7,_this.f11);
        let _556_v11;
        _556_v11 = _dafny.Set.fromElements((_this).f12);
        (_this).f11 = _dafny.Seq.IsPrefixOf((_this).fm8((_this).f13, _555_v10, _556_v11, globalState), _dafny.Seq.UnicodeFromString("wwb"));
        let _557_v12;
        _557_v12 = _dafny.Seq.UnicodeFromString("x");
        _557_v12 = _dafny.Seq.Concat((_this).f20, (_this).f20);
      }
      let _558_v13;
      let _nw75 = Array((new BigNumber(5)).toNumber());
      _558_v13 = _nw75;
      let _559_v14;
      _559_v14 = _dafny.MultiSet.fromElements(_558_v13, _558_v13);
      _559_v14 = _559_v14;
      let _560_v15;
      _560_v15 = _module.D4.create_DC12(_this.f11, p1, p0);
      let _561_v16;
      _561_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
      let _pat_let_tv13 = p1;
      let _562_v17;
      _562_v17 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let6_0) {
        return function (_563_dt__update__tmp_h0) {
          return function (_pat_let7_0) {
            return function (_566_dt__update_hcf22_h0) {
              return function (_pat_let8_0) {
                return function (_567_dt__update_hcf21_h0) {
                  return _module.D4.create_DC12(_567_dt__update_hcf21_h0, _566_dt__update_hcf22_h0, (_563_dt__update__tmp_h0).dtor_cf23);
                }(_pat_let8_0);
              }(_this.f11);
            }(_pat_let7_0);
          }(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-544)), ((_564_p1) => function (_565_i1) {
            return _564_p1;
          })(_pat_let_tv13))).length));
        }(_pat_let6_0);
      }(_560_v15)).dtor_cf21,new BigNumber((((_this.f11) ? (_561_v16) : (_561_v16))).length));
      let _568_v18;
      _568_v18 = _dafny.Seq.of(_562_v17, _562_v17);
      let _569_v19;
      _569_v19 = _module.D0.create_DC1(p1, (_this).f20, (_this).f26, _this.f11);
      _562_v17 = ((_this.f11) ? ((_568_v18)[_module.__default.safeIndex((_this).f26, new BigNumber((_568_v18).length))]) : ((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber(((_569_v19).dtor_cf2).length))).Merge(_562_v17)));
      let _570_v20;
      let _nw76 = Array((new BigNumber(29)).toNumber()).fill([]);
      _570_v20 = _nw76;
      let _571_v21;
      _571_v21 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),_570_v20);
      let _572_v22;
      _572_v22 = _dafny.Map.Empty.slice().updateUnsafe((((_571_v21).contains(_this.f11)) ? ((_571_v21).get(_this.f11)) : (_570_v20)),_this.f11);
      if (!((_this.f11) === ((((_572_v22).contains(_570_v20)) ? ((_572_v22).get(_570_v20)) : (_this.f11))))) {
        let _index61 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length));
        (p3)[_index61] = (_dafny.ZERO).minus((_this).f26);
        let _index62 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length));
        (p3)[_index62] = p0;
        if ((_this.f11) && (_this.f11)) {
          r1 = _module.__default.safeModuloInt((_this).f13, (_this).f26);
          r1 = (_this).f12;
          (globalState).f8 = !(_this.f11);
          let _573_v23;
          let _init19 = function (_574_i2) {
            return _this.f11;
          };
          let _nw77 = Array((new BigNumber(22)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw77.length); _i0_19++) {
            _nw77[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _573_v23 = _nw77;
          let _index63 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_573_v23).length));
          (_573_v23)[_index63] = _this.f11;
          let _575_v24;
          let _init20 = ((_576_v15) => function (_577_i3) {
            return _576_v15;
          })(_560_v15);
          let _nw78 = Array((new BigNumber(4)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw78.length); _i0_20++) {
            _nw78[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _575_v24 = _nw78;
          let _index64 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_575_v24).length));
          (_575_v24)[_index64] = _560_v15;
          let _578_v25;
          _578_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(330),_this.f11);
          let _index65 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_573_v23).length));
          let _index66 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_575_v24).length));
          let _rhs50 = _this.f11;
          let _rhs51 = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((p3)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length))],_this.f11)).Merge(_578_v25)).length)).multipliedBy(p2);
          let _rhs52 = (((_this).f12).isLessThan(new BigNumber(767))) === (!(_this.f11));
          let _rhs53 = ((true) ? (_560_v15) : (_module.D4.create_DC12(_this.f11, new BigNumber(((_this).f20).length), (_dafny.ZERO).minus(p2))));
          let _rhs54 = ((_this.f11) ? (_573_v23) : (_573_v23));
          let _lhs33 = _573_v23;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_573_v23).length));
          let _lhs35 = p3;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length));
          let _lhs37 = globalState;
          let _lhs38 = _575_v24;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_575_v24).length));
          _lhs33[_lhs34] = _rhs50;
          _lhs35[_lhs36] = _rhs51;
          _lhs37.f8 = _rhs52;
          _lhs38[_lhs39] = _rhs53;
          _573_v23 = _rhs54;
          let _579_v26;
          _579_v26 = _dafny.Seq.UnicodeFromString("b");
          _579_v26 = _579_v26;
        } else {
          r1 = new BigNumber(758);
          let _index68 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length));
          (p3)[_index68] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f13, ((_this).f13).multipliedBy(p2)));
          (_this).f11 = _this.f11;
          let _580_v27;
          let _nw79 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _580_v27 = _nw79;
          let _nw80 = Array((new BigNumber(10)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = (p3)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length))];
          _nw80[(_dafny.ONE).toNumber()] = (p2).multipliedBy(p2);
          _nw80[(new BigNumber(2)).toNumber()] = (p3)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length))];
          _nw80[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f12));
          _nw80[(new BigNumber(4)).toNumber()] = new BigNumber(505);
          _nw80[(new BigNumber(5)).toNumber()] = p1;
          _nw80[(new BigNumber(6)).toNumber()] = p2;
          _nw80[(new BigNumber(7)).toNumber()] = new BigNumber(344);
          _nw80[(new BigNumber(8)).toNumber()] = (_this).f26;
          _nw80[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_this.f11)).length);
          _580_v27 = _nw80;
          let _581_v28;
          let _nw81 = Array((new BigNumber(10)).toNumber()).fill(false);
          _581_v28 = _nw81;
          r0 = _581_v28;
        }
        if (_this.f11) {
          _561_v16 = (_561_v16).update(!_dafny.areEqual((_this).f20, (_this).f20), _this.f11);
          r1 = ((new BigNumber(-753)).minus(_module.__default.fm2(globalState))).minus(p0);
          let _582_v29;
          _582_v29 = _dafny.MultiSet.fromElements(p2);
          let _583_v30;
          _583_v30 = _dafny.MultiSet.fromElements((_582_v29).IsDisjointFrom(_582_v29), _this.f11, _this.f11, _this.f11, !((_this).fm18(_this.f11, _this.f11, (_this).f26, p1, globalState)));
          _583_v30 = _583_v30;
          let _584_v31;
          _584_v31 = _dafny.Seq.of(new BigNumber(((_561_v16).Merge(_561_v16)).length));
          (globalState).f5 = _584_v31;
          let _585_v32;
          let _nw82 = new _module.C0();
          _nw82.__ctor(_module.__default.fm2(globalState), _this.f11, (p3)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length))], p2);
          _585_v32 = _nw82;
        } else {
          r1 = (_this).f12;
          let _586_v33;
          let _init21 = function (_587_i4) {
            return (_this).f20;
          };
          let _nw83 = Array((new BigNumber(14)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw83.length); _i0_21++) {
            _nw83[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _586_v33 = _nw83;
          let _588_v34;
          _588_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_586_v33);
          _588_v34 = (_588_v34).update((_this).f20, (_module.D5.create_DC14(_586_v33)).dtor_cf25);
          let _589_v35;
          _589_v35 = new _dafny.CodePoint('s'.codePointAt(0));
          let _590_v36;
          _590_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_591_i5) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_592_i5) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          })).length)), _589_v35),(_this).f20);
          _590_v36 = _590_v36;
          let _593_v37;
          _593_v37 = _module.D3.create_DC8(_562_v17);
          let _pat_let_tv14 = _562_v17;
          let _594_v38;
          _594_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,function (_pat_let9_0) {
            return function (_595_dt__update__tmp_h1) {
              return function (_pat_let10_0) {
                return function (_596_dt__update_hcf14_h0) {
                  return _module.D3.create_DC8(_596_dt__update_hcf14_h0);
                }(_pat_let10_0);
              }(_pat_let_tv14);
            }(_pat_let9_0);
          }(_593_v37));
          let _597_v39;
          _597_v39 = _dafny.MultiSet.fromElements(_594_v38, _594_v38);
          let _598_v40;
          _598_v40 = _dafny.Map.Empty.slice().updateUnsafe(_597_v39,_589_v35);
          _598_v40 = (_598_v40).update(_597_v39, _589_v35);
          let _599_v41;
          _599_v41 = _dafny.Seq.of((_this).f13);
          let _600_v42;
          _600_v42 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_599_v41, _module.__default.safeIndex((p3)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((p3).length))], new BigNumber((_599_v41).length)), p0), _599_v41);
          let _601_v43;
          _601_v43 = _dafny.MultiSet.fromElements(new BigNumber((_600_v42).cardinality()), p2, new BigNumber(-503), new BigNumber(235));
          _561_v16 = (_561_v16).update(true, (_601_v43).IsProperSubsetOf(_601_v43));
        }
        let _602_v44;
        _602_v44 = _dafny.Seq.of(_this.f11, _this.f11, !(_this.f11));
        let _603_v45;
        let _nw84 = new _module.C1();
        _nw84.__ctor(new BigNumber((_602_v44).length));
        _603_v45 = _nw84;
        let _604_v46;
        let _nw85 = Array((new BigNumber(16)).toNumber()).fill(_module.D0.Default());
        _604_v46 = _nw85;
        let _605_v47;
        let _init22 = function (_606_i6) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-821)), function (_607_i7) {
            return new BigNumber(652);
          });
        };
        let _nw86 = Array((new BigNumber(4)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw86.length); _i0_22++) {
          _nw86[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _605_v47 = _nw86;
        let _608_v48;
        _608_v48 = _module.D0.create_DC0(_605_v47);
        let _index69 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_604_v46).length));
        (_604_v46)[_index69] = _608_v48;
        let _index70 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_604_v46).length));
        (_604_v46)[_index70] = _module.D0.create_DC0((_608_v48).dtor_cf0);
      } else {
        r1 = (_module.__default.fm2(globalState)).minus(p2);
        (_this).f11 = !((new BigNumber((_dafny.Seq.Concat((_this).f20, (_this).f20)).length)).isLessThan((_this).f12));
        (globalState).f8 = (_this).fm7(globalState);
        let _609_v49;
        _609_v49 = _dafny.MultiSet.fromElements((_this).f26, new BigNumber(583), p2, (_this).f12);
        let _610_v50;
        _610_v50 = _dafny.Map.Empty.slice().updateUnsafe((_609_v49).Intersect(_609_v49),_this.f11);
        let _611_v51;
        _611_v51 = _dafny.Seq.of(_this.f11, !(_this.f11));
        let _612_v52;
        _612_v52 = _dafny.Map.Empty.slice().updateUnsafe(_611_v51,new BigNumber(-494));
        let _613_v53;
        _613_v53 = _dafny.Seq.of(new BigNumber(849), (_this).f26, (((_612_v52).contains(_dafny.Seq.of(_this.f11, _this.f11))) ? ((_612_v52).get(_dafny.Seq.of(_this.f11, _this.f11))) : ((_this).f12)), (_this).f12);
        let _614_v54;
        _614_v54 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(350)), function (_615_i8) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }));
        _610_v50 = (_610_v50).update(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_613_v53, _613_v53)), _dafny.Seq.IsProperPrefixOf(_614_v54, _614_v54));
        if (_this.f11) {
          let _616_v55;
          let _nw87 = Array((_dafny.ONE).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = (_613_v53)[_module.__default.safeIndex((_this).f12, new BigNumber((_613_v53).length))];
          _616_v55 = _nw87;
          let _617_v56;
          _617_v56 = _dafny.Seq.of(p3, p3, p3, p3);
          _616_v55 = (_617_v56)[_module.__default.safeIndex((p1).plus(new BigNumber(175)), new BigNumber((_617_v56).length))];
          _616_v55 = _616_v55;
          (_this).f11 = _this.f11;
          let _618_v57;
          _618_v57 = _dafny.MultiSet.fromElements(true);
          _562_v17 = (_562_v17).update((_618_v57).IsSubsetOf(_618_v57), p2);
          let _619_v58;
          _619_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p2);
          let _620_v59;
          let _nw88 = new _module.C1();
          _nw88.__ctor(((_this).f12).multipliedBy(new BigNumber((_619_v58).length)));
          _620_v59 = _nw88;
        } else {
          let _621_v60;
          _621_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f12, (_dafny.ZERO).minus((_this).f13)),_module.__default.fm2(globalState));
          let _622_v61;
          _622_v61 = new _dafny.CodePoint('f'.codePointAt(0));
          _621_v60 = (_621_v60).update(((_this).f13).multipliedBy(p0), new BigNumber((_dafny.Seq.update((_this).f20, _module.__default.safeIndex(new BigNumber(2), new BigNumber(((_this).f20).length)), _622_v61)).length));
          (globalState).f5 = _613_v53;
          let _index71 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          (p3)[_index71] = (_this).f26;
          let _index72 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          (p3)[_index72] = (_dafny.ZERO).minus((_this).f13);
          let _623_v62;
          _623_v62 = _dafny.Set.fromElements(_this.f11);
          let _624_v63;
          let _nw89 = Array((new BigNumber(17)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _this.f11;
          _nw89[(_dafny.ONE).toNumber()] = _this.f11;
          _nw89[(new BigNumber(2)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(3)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(4)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(5)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(6)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(7)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(8)).toNumber()] = (_this).fm6(_this.f11, new BigNumber(((_561_v16).update(_this.f11, _this.f11)).length), new BigNumber(964), globalState);
          _nw89[(new BigNumber(9)).toNumber()] = !(_this.f11);
          _nw89[(new BigNumber(10)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(11)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(12)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(13)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(14)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(15)).toNumber()] = _this.f11;
          _nw89[(new BigNumber(16)).toNumber()] = _this.f11;
          _624_v63 = _nw89;
          let _625_v64;
          _625_v64 = _dafny.Seq.of((_this).f12, new BigNumber((_623_v62).length));
          let _626_v65;
          _626_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_625_v64);
          let _627_v66;
          _627_v66 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_624_v63, _624_v63),new BigNumber(((_626_v65).Merge(_626_v65)).length));
          let _628_v67;
          _628_v67 = _dafny.Seq.of(_624_v63);
          let _629_v68;
          _629_v68 = _dafny.Set.fromElements(new BigNumber(-78));
          let _index73 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          let _index74 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          let _rhs55 = !((p0).isLessThan((_dafny.ZERO).minus((p3)[_module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length))])));
          let _rhs56 = (p3)[_module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length))];
          let _rhs57 = (((_627_v66).contains(_628_v67)) ? ((_627_v66).get(_628_v67)) : (new BigNumber((_611_v51).length)));
          let _rhs58 = _623_v62;
          let _rhs59 = !(((_623_v62).IsProperSubsetOf(_623_v62)) && ((_629_v68).IsSubsetOf(_629_v68)));
          let _lhs40 = globalState;
          let _lhs41 = p3;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          let _lhs43 = p3;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((p3).length));
          let _lhs45 = globalState;
          _lhs40.f8 = _rhs55;
          _lhs41[_lhs42] = _rhs56;
          _lhs43[_lhs44] = _rhs57;
          _623_v62 = _rhs58;
          _lhs45.f8 = _rhs59;
          let _630_v69;
          _630_v69 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f11) ? (p1) : (p2)),_611_v51);
          _630_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_611_v51);
        }
      }
      let _631_v70;
      _631_v70 = _module.D3.create_DC8(_562_v17);
      let _632_v71;
      _632_v71 = _module.D3.create_DC10(_631_v70);
      let _633_v72;
      _633_v72 = _module.D3.create_DC10(_module.D3.create_DC10(_631_v70));
      let _634_v73;
      _634_v73 = _module.D3.create_DC10(_632_v71);
      let _635_v74;
      _635_v74 = _module.D3.create_DC10(_633_v72);
      let _636_v75;
      _636_v75 = _module.D3.create_DC10(_633_v72);
      let _pat_let_tv15 = p2;
      r1 = function (_source11) {
        if (_source11.is_DC9) {
          let _637___mcc_h0 = (_source11).cf15;
          let _638___mcc_h1 = (_source11).cf16;
          let _639___mcc_h2 = (_source11).cf17;
          let _640___mcc_h3 = (_source11).cf18;
          let _641_cf18 = _640___mcc_h3;
          let _642_cf17 = _639___mcc_h2;
          let _643_cf16 = _638___mcc_h1;
          let _644_cf15 = _637___mcc_h0;
          return _module.__default.safeModuloInt(_641_cf18, new BigNumber((_dafny.MultiSet.fromElements(_this.f11, _this.f11)).cardinality()));
        } else if (_source11.is_DC8) {
          let _645___mcc_h4 = (_source11).cf14;
          let _646_cf14 = _645___mcc_h4;
          return (_this).f13;
        } else {
          let _647___mcc_h5 = (_source11).cf19;
          let _648_cf19 = _647___mcc_h5;
          return _pat_let_tv15;
        }
      }(_636_v75);
      let _649_v76;
      let _650_v77;
      let _out2;
      let _out3;
      let _outcollector1 = _module.__default.m0(globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _649_v76 = _out2;
      _650_v77 = _out3;
      let _651_v78;
      let _init23 = ((_652_v77) => function (_653_i9) {
        return !(_652_v77) || (_this.f11);
      })(_650_v77);
      let _nw90 = Array((new BigNumber(26)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw90.length); _i0_23++) {
        _nw90[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _651_v78 = _nw90;
      r0 = _651_v78;
      r1 = p1;
      r2 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus((_this).f12), (_this).f12);
      return [r0, r1, r2];
    }
    m23(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D3.Default();
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _hi1 = (_this).f13;
      for (let _654_i0 = (_this).f12; _654_i0.isLessThan(_hi1); _654_i0 = _654_i0.plus(_dafny.ONE)) {
        r0 = (_this).f20;
        r2 = _module.__default.safeDivisionInt((_this).f13, p2);
        let _655_v0;
        let _init24 = ((_656_p0) => function (_657_i1) {
          return _656_p0;
        })(p0);
        let _nw91 = Array((new BigNumber(26)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw91.length); _i0_24++) {
          _nw91[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _655_v0 = _nw91;
        let _nw92 = Array((new BigNumber(4)).toNumber()).fill(false);
        _655_v0 = _nw92;
        (_this).f11 = _this.f11;
      }
      let _658_v1;
      _658_v1 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),!(_this.f11));
      let _659_i2;
      _659_i2 = _dafny.ZERO;
      L2: {
        while (((_658_v1).Merge(_658_v1)).equals((_658_v1).Merge(_658_v1))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_659_i2)) {
              break L2;
            }
            _659_i2 = (_659_i2).plus(_dafny.ONE);
            let _660_v2;
            _660_v2 = new _dafny.CodePoint('s'.codePointAt(0));
            let _661_v3;
            _661_v3 = _module.D0.create_DC1((_this).f13, (_this).f20, p1, p0);
            let _source12 = _module.__default.fm37((_this).fm6(p0, new BigNumber((_module.__default.fm36((_this).f12, p2, _660_v2, globalState)).length), p2, globalState), (_661_v3).dtor_cf1, (_this).fm33(p2, globalState), p2, globalState);
            if (_source12.is_DC15) {
              let _662___mcc_h0 = (_source12).cf26;
              let _663___mcc_h1 = (_source12).cf27;
              let _664___mcc_h2 = (_source12).cf28;
              let _665_cf28 = _664___mcc_h2;
              let _666_cf27 = _663___mcc_h1;
              let _667_cf26 = _662___mcc_h0;
              _666_cf27 = (_this).f13;
              let _668_v4;
              _668_v4 = _dafny.Seq.of(_665_cf28);
              let _669_v5;
              let _nw93 = Array((new BigNumber(17)).toNumber());
              _nw93[(_dafny.ZERO).toNumber()] = _666_cf27;
              _nw93[(_dafny.ONE).toNumber()] = new BigNumber(61);
              _nw93[(new BigNumber(2)).toNumber()] = _667_cf26;
              _nw93[(new BigNumber(3)).toNumber()] = _666_cf27;
              _nw93[(new BigNumber(4)).toNumber()] = (_this).f26;
              _nw93[(new BigNumber(5)).toNumber()] = (_668_v4)[_module.__default.safeIndex(new BigNumber((_658_v1).length), new BigNumber((_668_v4).length))];
              _nw93[(new BigNumber(6)).toNumber()] = new BigNumber(((_this).f20).length);
              _nw93[(new BigNumber(7)).toNumber()] = new BigNumber(975);
              _nw93[(new BigNumber(8)).toNumber()] = (_this).f12;
              _nw93[(new BigNumber(9)).toNumber()] = (_this).f26;
              _nw93[(new BigNumber(10)).toNumber()] = (_this).f26;
              _nw93[(new BigNumber(11)).toNumber()] = p1;
              _nw93[(new BigNumber(12)).toNumber()] = _666_cf27;
              _nw93[(new BigNumber(13)).toNumber()] = _module.__default.fm2(globalState);
              _nw93[(new BigNumber(14)).toNumber()] = _667_cf26;
              _nw93[(new BigNumber(15)).toNumber()] = (_this).f13;
              _nw93[(new BigNumber(16)).toNumber()] = _665_cf28;
              _669_v5 = _nw93;
              let _670_v6;
              _670_v6 = _dafny.MultiSet.fromElements(_669_v5, _669_v5, _669_v5, _669_v5, _669_v5);
              let _671_v7;
              _671_v7 = _dafny.Seq.of(p0, p0);
              let _672_v8;
              _672_v8 = _dafny.MultiSet.fromElements(new BigNumber((_671_v7).length), (_this).f26);
              let _673_v9;
              _673_v9 = _dafny.MultiSet.fromElements(_660_v2, _660_v2, _660_v2, _660_v2, _660_v2);
              let _674_v10;
              _674_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
              let _675_v11;
              _675_v11 = _dafny.Map.Empty.slice().updateUnsafe(_661_v3,(_this).f13);
              let _676_v12;
              let _nw94 = Array((new BigNumber(19)).toNumber());
              _nw94[(_dafny.ZERO).toNumber()] = (_this).f13;
              _nw94[(_dafny.ONE).toNumber()] = _665_cf28;
              _nw94[(new BigNumber(2)).toNumber()] = new BigNumber((((_670_v6).update(_669_v5, _module.__default.abs((_this).f26))).Union(_670_v6)).cardinality());
              _nw94[(new BigNumber(3)).toNumber()] = ((_this).f26).multipliedBy(new BigNumber(((_this).f20).length));
              _nw94[(new BigNumber(4)).toNumber()] = _666_cf27;
              _nw94[(new BigNumber(5)).toNumber()] = p2;
              _nw94[(new BigNumber(6)).toNumber()] = p2;
              _nw94[(new BigNumber(7)).toNumber()] = (_this).f26;
              _nw94[(new BigNumber(8)).toNumber()] = _667_cf26;
              _nw94[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(p1, _667_cf26);
              _nw94[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), _module.__default.fm2(globalState));
              _nw94[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(((_this.f11) ? ((_this).f13) : ((_dafny.ZERO).minus(_666_cf27))));
              _nw94[(new BigNumber(12)).toNumber()] = new BigNumber(((_672_v8).Difference(_672_v8)).cardinality());
              _nw94[(new BigNumber(13)).toNumber()] = p1;
              _nw94[(new BigNumber(14)).toNumber()] = (((_673_v9).contains(_660_v2)) ? ((_673_v9).get(_660_v2)) : ((_this).f13));
              _nw94[(new BigNumber(15)).toNumber()] = (((_674_v10).contains(_this.f11)) ? ((_674_v10).get(_this.f11)) : ((_dafny.ZERO).minus(p1)));
              _nw94[(new BigNumber(16)).toNumber()] = (new BigNumber((_671_v7).length)).multipliedBy((_this).f13);
              _nw94[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(p1, (((_675_v11).contains(_661_v3)) ? ((_675_v11).get(_661_v3)) : (_666_cf27)));
              _nw94[(new BigNumber(18)).toNumber()] = (_this).f26;
              _676_v12 = _nw94;
              let _nw95 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _676_v12 = _nw95;
              let _677_v13;
              _677_v13 = _dafny.Set.fromElements((_this).f13, (_this).f13);
              let _678_v14;
              _678_v14 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f13).multipliedBy(_666_cf27),(_677_v13).Difference(_677_v13));
              _678_v14 = _678_v14;
              _676_v12 = _669_v5;
            } else {
              let _679___mcc_h3 = (_source12).cf25;
              let _680_cf25 = _679___mcc_h3;
              let _681_v15;
              let _nw96 = Array((new BigNumber(4)).toNumber());
              _nw96[(_dafny.ZERO).toNumber()] = (_this).f12;
              _nw96[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), function (_682_i3) {
                return _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f26));
              })).length);
              _nw96[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f13);
              _nw96[(new BigNumber(3)).toNumber()] = (_this).f13;
              _681_v15 = _nw96;
              let _683_v16;
              _683_v16 = _dafny.Map.Empty.slice().updateUnsafe(_681_v15,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(p0)).length))).cardinality()));
              let _684_v17;
              _684_v17 = _dafny.MultiSet.fromElements((((_683_v16).contains(_681_v15)) ? ((_683_v16).get(_681_v15)) : ((_this).f12)));
              let _685_v18;
              _685_v18 = _dafny.MultiSet.fromElements((((_684_v17).contains(p2)) ? ((_684_v17).get(p2)) : (p2)));
              let _686_v19;
              let _nw97 = new _module.C0();
              _nw97.__ctor((_this).f12, p0, (_this).f13, ((_this).f12).minus((((_684_v17).contains((_dafny.ZERO).minus(new BigNumber(-963)))) ? ((_684_v17).get((_dafny.ZERO).minus(new BigNumber(-963)))) : (p1))));
              _686_v19 = _nw97;
              _686_v19 = ((_this.f11) ? (_686_v19) : (((!(_686_v19.f29)) ? (_686_v19) : (_686_v19))));
              _660_v2 = _660_v2;
              let _687_v20;
              _687_v20 = _dafny.MultiSet.fromElements(_this.f11);
              _687_v20 = (_687_v20).Union(_687_v20);
              _660_v2 = _660_v2;
            }
            let _688_v21;
            let _init25 = ((_689_v2) => function (_690_i4) {
              return _dafny.Map.Empty.slice().updateUnsafe(_689_v2,_this.f11);
            })(_660_v2);
            let _nw98 = Array((new BigNumber(19)).toNumber());
            for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw98.length); _i0_25++) {
              _nw98[_i0_25] = _init25(new BigNumber(_i0_25));
            }
            _688_v21 = _nw98;
            let _691_v22;
            _691_v22 = _dafny.Map.Empty.slice().updateUnsafe(_660_v2,_this.f11);
            let _index75 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_688_v21).length));
            (_688_v21)[_index75] = (_691_v22).Merge((_dafny.Map.Empty.slice().updateUnsafe(_660_v2,_this.f11)).update(_660_v2, p0));
            let _692_v23;
            _692_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f12);
            let _693_v24;
            let _nw99 = Array((new BigNumber(20)).toNumber());
            _nw99[(_dafny.ZERO).toNumber()] = (_this).f13;
            _nw99[(_dafny.ONE).toNumber()] = p2;
            _nw99[(new BigNumber(2)).toNumber()] = new BigNumber((_692_v23).length);
            _nw99[(new BigNumber(3)).toNumber()] = p2;
            _nw99[(new BigNumber(4)).toNumber()] = (_this).f13;
            _nw99[(new BigNumber(5)).toNumber()] = (_this).f12;
            _nw99[(new BigNumber(6)).toNumber()] = (_this).f12;
            _nw99[(new BigNumber(7)).toNumber()] = (_this).f12;
            _nw99[(new BigNumber(8)).toNumber()] = p2;
            _nw99[(new BigNumber(9)).toNumber()] = p1;
            _nw99[(new BigNumber(10)).toNumber()] = (_this).f13;
            _nw99[(new BigNumber(11)).toNumber()] = p1;
            _nw99[(new BigNumber(12)).toNumber()] = (_this).f12;
            _nw99[(new BigNumber(13)).toNumber()] = (_this).f12;
            _nw99[(new BigNumber(14)).toNumber()] = p2;
            _nw99[(new BigNumber(15)).toNumber()] = (_this).f13;
            _nw99[(new BigNumber(16)).toNumber()] = p2;
            _nw99[(new BigNumber(17)).toNumber()] = (_this).f26;
            _nw99[(new BigNumber(18)).toNumber()] = (_this).f13;
            _nw99[(new BigNumber(19)).toNumber()] = p2;
            _693_v24 = _nw99;
            let _694_v25;
            _694_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,_693_v24);
            let _695_v26;
            _695_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_694_v25).contains(p2)) ? ((_694_v25).get(p2)) : (_693_v24)),((p0) ? ((_this).f12) : ((_this).f12)));
            let _696_v27;
            _696_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(928),(_this).f12);
            let _index76 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_688_v21).length));
            let _rhs60 = _module.__default.safeModuloInt((_this).f12, new BigNumber((((_this.f11) ? (_module.__default.fm38(_660_v2, _660_v2, globalState)) : (_696_v27))).length));
            let _rhs61 = _691_v22;
            let _rhs62 = (_695_v26).Merge(_695_v26);
            let _lhs46 = _688_v21;
            let _lhs47 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_688_v21).length));
            r2 = _rhs60;
            _lhs46[_lhs47] = _rhs61;
            _695_v26 = _rhs62;
            r2 = new BigNumber(-497);
            let _697_v28;
            let _nw100 = new _module.C0();
            _nw100.__ctor(p2, _this.f11, (_this).f26, new BigNumber(712));
            _697_v28 = _nw100;
          }
        }
      }
      let _698_v29;
      _698_v29 = new _dafny.CodePoint('f'.codePointAt(0));
      let _699_v30;
      _699_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,_698_v29);
      let _700_v32;
      _700_v32 = _dafny.MultiSet.fromElements(function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of _dafny.IntegerRange(new BigNumber(225), new BigNumber(317))) {
          let _701_v31 = _compr_48;
          if (((new BigNumber(225)).isLessThanOrEqualTo(_701_v31)) && ((_701_v31).isLessThan(new BigNumber(317)))) {
            _coll48.push([(_701_v31).plus(new BigNumber((_dafny.Seq.of(p0)).length)),_698_v29]);
          }
        }
        return _coll48;
      }());
      if (!(!(false) || ((_700_v32).contains(_699_v30)))) {
        let _702_v33;
        _702_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,p1);
        let _703_v34;
        let _nw101 = new _module.C0();
        _nw101.__ctor((p2).plus((((_702_v33).contains(_dafny.Seq.UnicodeFromString("hfxotwq"))) ? ((_702_v33).get(_dafny.Seq.UnicodeFromString("hfxotwq"))) : ((_dafny.ZERO).minus((_this).f13)))), _this.f11, new BigNumber(((_this).f20).length), (new BigNumber(((_this).f20).length)).minus((_this).f26));
        _703_v34 = _nw101;
        let _704_v35;
        _704_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f13),(_this).f12);
        let _705_v36;
        _705_v36 = _dafny.Seq.of(p0, ((_this).f13).isLessThan(p2), (new BigNumber(58)).isLessThanOrEqualTo((((_704_v35).contains(p1)) ? ((_704_v35).get(p1)) : (p1))), _703_v34.f29);
        if ((_705_v36)[_module.__default.safeIndex(p2, new BigNumber((_705_v36).length))]) {
          r2 = new BigNumber(91);
          let _706_v37;
          _706_v37 = _dafny.Set.fromElements(p0);
          let _707_v38;
          _707_v38 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),_706_v37);
          let _708_v39;
          let _nw102 = new _module.C0();
          _nw102.__ctor((_this).f26, _703_v34.f29, new BigNumber((_707_v38).length), new BigNumber(839));
          _708_v39 = _nw102;
          let _709_v40;
          let _710_v41;
          let _out4;
          let _out5;
          let _outcollector2 = _module.__default.m0(globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _709_v40 = _out4;
          _710_v41 = _out5;
          let _711_v42;
          _711_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f11);
          let _712_v43;
          _712_v43 = _dafny.Seq.of(_711_v42, _711_v42);
          let _713_v44;
          let _nw103 = Array((new BigNumber(16)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = new BigNumber(983);
          _nw103[(_dafny.ONE).toNumber()] = p1;
          _nw103[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw103[(new BigNumber(3)).toNumber()] = p1;
          _nw103[(new BigNumber(4)).toNumber()] = (_this).f12;
          _nw103[(new BigNumber(5)).toNumber()] = new BigNumber((_705_v36).length);
          _nw103[(new BigNumber(6)).toNumber()] = (_703_v34).f28;
          _nw103[(new BigNumber(7)).toNumber()] = (_703_v34).f28;
          _nw103[(new BigNumber(8)).toNumber()] = new BigNumber(717);
          _nw103[(new BigNumber(9)).toNumber()] = new BigNumber((_712_v43).length);
          _nw103[(new BigNumber(10)).toNumber()] = (_708_v39).f28;
          _nw103[(new BigNumber(11)).toNumber()] = new BigNumber(242);
          _nw103[(new BigNumber(12)).toNumber()] = ((_this).f26).plus((_708_v39).f28);
          _nw103[(new BigNumber(13)).toNumber()] = (_703_v34).f28;
          _nw103[(new BigNumber(14)).toNumber()] = (_703_v34).f28;
          _nw103[(new BigNumber(15)).toNumber()] = new BigNumber(248);
          _713_v44 = _nw103;
          _713_v44 = _713_v44;
          r2 = (_this).f13;
        } else {
          let _714_v45;
          let _nw104 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _714_v45 = _nw104;
          let _index77 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_714_v45).length));
          (_714_v45)[_index77] = (_this).f26;
          let _index78 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_714_v45).length));
          (_714_v45)[_index78] = (_this).f12;
          (globalState).f8 = _this.f11;
          let _index79 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_714_v45).length));
          (_714_v45)[_index79] = (_703_v34).f28;
          (_703_v34).f29 = !(p0);
          r2 = ((_this).f12).plus((_703_v34).f28);
        }
        let _715_v46;
        let _init26 = ((_716_v34) => function (_717_i5) {
          return _716_v34.f29;
        })(_703_v34);
        let _nw105 = Array((new BigNumber(11)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw105.length); _i0_26++) {
          _nw105[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _715_v46 = _nw105;
        let _718_v47;
        _718_v47 = _dafny.Map.Empty.slice().updateUnsafe(((_703_v34.f29) ? (_698_v29) : (_698_v29)),new BigNumber(798));
        let _719_v48;
        _719_v48 = _dafny.Seq.of(p1);
        let _rhs63 = _715_v46;
        let _rhs64 = _dafny.Seq.Concat((_this).f20, (_this).f20);
        let _rhs65 = (((_718_v47).contains(_698_v29)) ? ((_718_v47).get(_698_v29)) : (new BigNumber((_719_v48).length)));
        _715_v46 = _rhs63;
        r0 = _rhs64;
        r2 = _rhs65;
        if (_703_v34.f29) {
          let _720_v49;
          let _nw106 = new _module.C0();
          _nw106.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_703_v34).f28, (_this).f13)), _703_v34.f29, (_this).f26, (_719_v48)[_module.__default.safeIndex((_this).f26, new BigNumber((_719_v48).length))]);
          _720_v49 = _nw106;
          (_720_v49).f29 = _this.f11;
          r2 = ((_720_v49).f28).multipliedBy((_this).f12);
          let _721_v50;
          _721_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f26);
          let _722_v51;
          _722_v51 = _dafny.Set.fromElements(new BigNumber((_721_v50).length), new BigNumber(-61));
          _722_v51 = _722_v51;
          r2 = ((_720_v49.f29) ? (new BigNumber(-438)) : ((_this).f26));
        } else {
          let _rhs66 = new BigNumber((_dafny.Seq.of((_703_v34).f28, p2)).length);
          let _rhs67 = _dafny.Seq.Concat(_705_v36, _dafny.Seq.Concat(_705_v36, _module.__default.fm39(_703_v34.f29, false, globalState)));
          r2 = _rhs66;
          _705_v36 = _rhs67;
          let _723_v53;
          let _init27 = ((_724_v34) => function (_725_i6) {
            return (_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe((_724_v34).f28,_module.D5.create_DC15((_724_v34).f28, (_724_v34).f28, new BigNumber((function () {
  let _coll49 = new _dafny.Map();
  for (const _compr_49 of _dafny.IntegerRange(new BigNumber(916), new BigNumber(751))) {
    let _726_v52 = _compr_49;
    if (((new BigNumber(916)).isLessThanOrEqualTo(_726_v52)) && ((_726_v52).isLessThan(new BigNumber(751)))) {
      _coll49.push([_module.__default.safeModuloInt(_726_v52, new BigNumber(781)),(_this).f20]);
    }
  }
  return _coll49;
}()).length))))).dtor_cf29;
          })(_703_v34);
          let _nw107 = Array((new BigNumber(27)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw107.length); _i0_27++) {
            _nw107[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _723_v53 = _nw107;
          let _727_v54;
          _727_v54 = _dafny.MultiSet.fromElements(p1);
          let _728_v55;
          _728_v55 = _module.D1.create_DC5((_727_v54).Union(_727_v54), ((_dafny.ZERO).minus(p2)).isLessThanOrEqualTo(new BigNumber((_module.__default.fm36((_703_v34).f28, (_703_v34).f28, _698_v29, globalState)).length)));
          let _729_v56;
          _729_v56 = _module.D4.create_DC13((_this).f13);
          let _rhs68 = _698_v29;
          let _rhs69 = _module.__default.fm36(((_703_v34).f28).multipliedBy((_729_v56).dtor_cf24), new BigNumber(466), _698_v29, globalState);
          let _rhs70 = _723_v53;
          let _rhs71 = _module.D1.create_DC5(_727_v54, p0);
          let _rhs72 = new BigNumber(((_this).f20).length);
          _698_v29 = _rhs68;
          r0 = _rhs69;
          _723_v53 = _rhs70;
          _728_v55 = _rhs71;
          r2 = _rhs72;
          let _730_v57;
          _730_v57 = _dafny.Map.Empty.slice().updateUnsafe(!(_703_v34.f29),p2);
          _730_v57 = (_730_v57).update(!(_730_v57).contains(_this.f11), p2);
          (_this).f11 = p0;
          _704_v35 = (_704_v35).update(((_this).f13).plus((_this).f26), (_703_v34).f28);
        }
        let _731_v58;
        _731_v58 = _dafny.Set.fromElements(false);
        let _rhs73 = (((_705_v36)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_705_v36).length))]) ? (_this.f11) : (p0));
        let _rhs74 = (_this).f26;
        let _rhs75 = _dafny.Set.fromElements((_705_v36)[_module.__default.safeIndex(p1, new BigNumber((_705_v36).length))]);
        let _rhs76 = _698_v29;
        let _lhs48 = globalState;
        _lhs48.f8 = _rhs73;
        r2 = _rhs74;
        _731_v58 = _rhs75;
        _698_v29 = _rhs76;
      } else {
        (_this).f11 = p0;
        if (!((_this).fm7(globalState))) {
          let _732_v59;
          let _nw108 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _732_v59 = _nw108;
          let _index80 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length));
          (_732_v59)[_index80] = (_this).f26;
          let _index81 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length));
          (_732_v59)[_index81] = ((_this).f12).minus((_this).f12);
          let _733_v60;
          let _734_v61;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m0(globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _733_v60 = _out6;
          _734_v61 = _out7;
          let _735_v62;
          _735_v62 = _dafny.MultiSet.fromElements((_this).f13);
          let _736_v63;
          _736_v63 = _dafny.Map.Empty.slice().updateUnsafe(_735_v62,(_this).f13);
          let _737_v64;
          _737_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm33(p2, globalState),p2);
          let _index82 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length));
          let _rhs77 = (_this).f26;
          let _rhs78 = ((_dafny.ZERO).minus((((_737_v64).contains(p0)) ? ((_737_v64).get(p0)) : (p2)))).isLessThanOrEqualTo((((_736_v63).contains((_module.D1.create_DC5(_735_v62, _734_v61)).dtor_cf10)) ? ((_736_v63).get((_module.D1.create_DC5(_735_v62, _734_v61)).dtor_cf10)) : ((_this).f12)));
          let _rhs79 = _734_v61;
          let _rhs80 = p0;
          let _rhs81 = false;
          let _lhs49 = _732_v59;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length));
          let _lhs51 = _this;
          _lhs49[_lhs50] = _rhs77;
          _lhs51.f11 = _rhs78;
          r3 = _rhs79;
          _734_v61 = _rhs80;
          r3 = _rhs81;
          let _index83 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length));
          (_732_v59)[_index83] = (_this).f13;
          let _738_v65;
          _738_v65 = _module.D4.create_DC13((_dafny.ZERO).minus((_732_v59)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_732_v59).length))]));
          _738_v65 = _738_v65;
        } else {
          _698_v29 = _698_v29;
          let _739_v66;
          _739_v66 = _dafny.MultiSet.fromElements(_this.f11);
          let _740_v67;
          _740_v67 = _dafny.Set.fromElements(new BigNumber((_739_v66).cardinality()), new BigNumber((_658_v1).length), (_this).f26, (_dafny.ZERO).minus((_this).f26), _module.__default.fm2(globalState));
          let _741_v68;
          _741_v68 = _dafny.Seq.of(p1, (_this).f13);
          _740_v67 = ((function () {
            let _coll50 = new _dafny.Set();
            for (const _compr_50 of (_741_v68).Elements) {
              let _742_v69 = _compr_50;
              if (_dafny.Seq.contains(_741_v68, _742_v69)) {
                _coll50.add((_742_v69).plus(new BigNumber(956)));
              }
            }
            return _coll50;
          }()).Union(function () {
            let _coll51 = new _dafny.Set();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(896), new BigNumber(705))) {
              let _743_v70 = _compr_51;
              if (((new BigNumber(896)).isLessThanOrEqualTo(_743_v70)) && ((_743_v70).isLessThan(new BigNumber(705)))) {
                _coll51.add((_743_v70).plus((_this).f26));
              }
            }
            return _coll51;
          }())).Intersect(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of _dafny.IntegerRange(new BigNumber(352), new BigNumber(698))) {
              let _744_v71 = _compr_52;
              if (((new BigNumber(352)).isLessThanOrEqualTo(_744_v71)) && ((_744_v71).isLessThan(new BigNumber(698)))) {
                _coll52.add((_744_v71).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D4.create_DC12(_this.f11, new BigNumber(((_this).f20).length), new BigNumber((_dafny.Seq.of(_this.f11)).length)),_739_v66)).length)));
              }
            }
            return _coll52;
          }());
          let _745_v72;
          let _init28 = function (_746_i7) {
            return !(_this.f11);
          };
          let _nw109 = Array((new BigNumber(12)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw109.length); _i0_28++) {
            _nw109[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _745_v72 = _nw109;
          let _747_v73;
          _747_v73 = _dafny.Seq.of(_745_v72, _745_v72);
          _745_v72 = (_747_v73)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_747_v73).length))];
          r2 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_748_v29) => function (_749_i8) {
            return _748_v29;
          })(_698_v29)), _dafny.Seq.Concat((_this).f20, (_this).f20)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_750_v29) => function (_751_i8) {
            return _750_v29;
          })(_698_v29)), _dafny.Seq.Concat((_this).f20, (_this).f20))).length)), ((_this).f20)[_module.__default.safeIndex((_this).f12, new BigNumber(((_this).f20).length))])).length);
          let _752_v74;
          _752_v74 = _dafny.Seq.of(_this.f11, _this.f11, p0);
          let _753_v75;
          let _init29 = ((_754_v68) => function (_755_i9) {
            return (_755_i9).minus((_754_v68)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_754_v68).length))]);
          })(_741_v68);
          let _nw110 = Array((new BigNumber(23)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw110.length); _i0_29++) {
            _nw110[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _753_v75 = _nw110;
          let _756_v76;
          let _nw111 = Array((new BigNumber(22)).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = _753_v75;
          _nw111[(_dafny.ONE).toNumber()] = _753_v75;
          _nw111[(new BigNumber(2)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(3)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(4)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(5)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(6)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(7)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(8)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(9)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(10)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(11)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(12)).toNumber()] = ((_this.f11) ? (_753_v75) : (_753_v75));
          _nw111[(new BigNumber(13)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(14)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(15)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(16)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(17)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(18)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(19)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(20)).toNumber()] = _753_v75;
          _nw111[(new BigNumber(21)).toNumber()] = _753_v75;
          _756_v76 = _nw111;
          let _index84 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_756_v76).length));
          (_756_v76)[_index84] = _753_v75;
          let _757_v77;
          _757_v77 = _module.D6.create_DC18(p0);
          let _index85 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_756_v76).length));
          let _rhs82 = _this.f11;
          let _rhs83 = true;
          let _rhs84 = _this.f11;
          let _rhs85 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p0, p0, p0, (_757_v77).dtor_cf32), _752_v74), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, p0, p0, (_757_v77).dtor_cf32), _752_v74)).length)), _this.f11);
          let _rhs86 = _753_v75;
          let _lhs52 = globalState;
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          let _lhs55 = _756_v76;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_756_v76).length));
          _lhs52.f8 = _rhs82;
          _lhs53.f8 = _rhs83;
          _lhs54.f8 = _rhs84;
          _752_v74 = _rhs85;
          _lhs55[_lhs56] = _rhs86;
        }
        let _758_v78;
        _758_v78 = _dafny.Set.fromElements(_this.f11);
        let _759_v79;
        _759_v79 = _dafny.Seq.of((_this).f12, new BigNumber((_758_v78).length));
        let _760_v80;
        _760_v80 = _dafny.Set.fromElements(_759_v79, _759_v79);
        let _761_v81;
        _761_v81 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_759_v79);
        (globalState).f8 = ((_dafny.Set.fromElements((((_761_v81).contains(p0)) ? ((_761_v81).get(p0)) : (_759_v79)), _759_v79, _dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_762_i10) {
          return (_this).f13;
        }))).Intersect(_760_v80)).IsProperSubsetOf(_760_v80);
        r2 = _module.__default.fm2(globalState);
        let _763_v82;
        _763_v82 = _dafny.Seq.of(p0);
        _763_v82 = _dafny.Seq.Concat(_dafny.Seq.Concat(_763_v82, _763_v82), _763_v82);
      }
      r3 = p0;
      let _764_v83;
      _764_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(619));
      let _765_v84;
      _765_v84 = _module.D3.create_DC8(_764_v83);
      _765_v84 = _module.__default.fm40(_dafny.Seq.contains(_dafny.Seq.update((_this).f20, _module.__default.safeIndex((_this).f12, new BigNumber(((_this).f20).length)), new _dafny.CodePoint('y'.codePointAt(0))), _698_v29), globalState);
      let _766_v85;
      _766_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,((_this).f26).isEqualTo(new BigNumber(905)));
      let _767_v87;
      _767_v87 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-311)), ((_768_v29) => function (_769_i11) {
        return _768_v29;
      })(_698_v29)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("a")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-311)), ((_770_v29) => function (_771_i11) {
        return _770_v29;
      })(_698_v29))).length)), _698_v29),(_this).f13);
      _766_v85 = function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of ((_767_v87).Merge(_767_v87)).Keys.Elements) {
          let _772_v86 = _compr_53;
          if (((_767_v87).Merge(_767_v87)).contains(_772_v86)) {
            _coll53.push([_772_v86,!(p0) || (_this.f11)]);
          }
        }
        return _coll53;
      }();
      r0 = (_this).f20;
      let _773_v88;
      _773_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f12);
      let _774_v89;
      _774_v89 = _dafny.Map.Empty.slice().updateUnsafe(p2,_773_v88);
      let _775_v90;
      let _nw112 = new _module.C0();
      _nw112.__ctor(new BigNumber(((_this).f20).length), false, (_this).f13, (_this).f13);
      _775_v90 = _nw112;
      let _776_v91;
      _776_v91 = _module.D3.create_DC10(_module.D3.create_DC9(p1, _774_v89, _module.__default.fm2(globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm7(globalState),_775_v90)).length)));
      let _777_v92;
      _777_v92 = _module.D3.create_DC8(_764_v83);
      let _pat_let_tv16 = _777_v92;
      r1 = function (_pat_let11_0) {
        return function (_778_dt__update__tmp_h0) {
          return function (_pat_let12_0) {
            return function (_779_dt__update_hcf19_h0) {
              return _module.D3.create_DC10(_779_dt__update_hcf19_h0);
            }(_pat_let12_0);
          }(_pat_let_tv16);
        }(_pat_let11_0);
      }(_776_v91);
      let _780_v93;
      _780_v93 = _dafny.Map.Empty.slice().updateUnsafe(_775_v90.f29,(_this).f20);
      let _781_v94;
      _781_v94 = _dafny.Seq.of(p1, (_775_v90).f28, new BigNumber(385), new BigNumber(((((_780_v93).contains(p0)) ? ((_780_v93).get(p0)) : ((_this).f20))).length));
      let _782_v95;
      _782_v95 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(new BigNumber((_781_v94).length))).multipliedBy((_this).f13),_775_v90.f29);
      r2 = new BigNumber((_782_v95).length);
      r3 = (p1).isLessThanOrEqualTo(new BigNumber(240));
      return [r0, r1, r2, r3];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return (_module.D0.create_DC3((_this).f13, !(true))).dtor_cf8;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qei"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-489)), function (_783_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lotyl"), _dafny.Seq.UnicodeFromString("st")));
    };
    m22(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = [];
      let r3 = _dafny.Set.Empty;
      let _784_v0;
      _784_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber(455));
      let _785_v1;
      _785_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_784_v0);
      let _786_v2;
      _786_v2 = _module.D3.create_DC9((_this).f13, _785_v1, (_this).f12, (_this).f12);
      let _787_v3;
      _787_v3 = _module.D3.create_DC10(_786_v2);
      let _788_v4;
      _788_v4 = _module.D3.create_DC10(_786_v2);
      let _source13 = _788_v4;
      if (_source13.is_DC9) {
        let _789___mcc_h0 = (_source13).cf15;
        let _790___mcc_h1 = (_source13).cf16;
        let _791___mcc_h2 = (_source13).cf17;
        let _792___mcc_h3 = (_source13).cf18;
        let _793_cf18 = _792___mcc_h3;
        let _794_cf17 = _791___mcc_h2;
        let _795_cf16 = _790___mcc_h1;
        let _796_cf15 = _789___mcc_h0;
        let _797_v5;
        _797_v5 = new _dafny.CodePoint('j'.codePointAt(0));
        let _798_v6;
        _798_v6 = _dafny.Set.fromElements(_797_v5, _797_v5, _797_v5);
        if (!((_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), _797_v5, _797_v5, _797_v5, _797_v5)).IsDisjointFrom(_798_v6))) {
          _794_cf17 = (_this).f13;
          let _799_v7;
          _799_v7 = false;
          let _800_v8;
          _800_v8 = _dafny.Map.Empty.slice().updateUnsafe(!(_799_v7),_799_v7);
          let _801_v9;
          let _nw113 = Array((new BigNumber(3)).toNumber());
          _nw113[(_dafny.ZERO).toNumber()] = _799_v7;
          _nw113[(_dafny.ONE).toNumber()] = _799_v7;
          _nw113[(new BigNumber(2)).toNumber()] = _799_v7;
          _801_v9 = _nw113;
          let _802_v10;
          _802_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_800_v8).length),_801_v9);
          let _803_v11;
          let _nw114 = new _module.C0();
          _nw114.__ctor(new BigNumber(((_802_v10).Merge(_802_v10)).length), _799_v7, new BigNumber(463), _794_cf17);
          _803_v11 = _nw114;
          let _804_v12;
          let _init30 = ((_805_v11) => function (_806_i0) {
            return (_806_i0).minus((_805_v11).f28);
          })(_803_v11);
          let _nw115 = Array((new BigNumber(11)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw115.length); _i0_30++) {
            _nw115[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _804_v12 = _nw115;
          r1 = _804_v12;
          let _807_v13;
          _807_v13 = _dafny.Seq.of(!(_803_v11.f29));
          let _808_v14;
          _808_v14 = _dafny.MultiSet.fromElements((_this).f13);
          let _809_v15;
          let _nw116 = Array((new BigNumber(21)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_799_v7, _799_v7);
          _nw116[(_dafny.ONE).toNumber()] = _807_v13;
          _nw116[(new BigNumber(2)).toNumber()] = ((_803_v11.f29) ? (_807_v13) : (_807_v13));
          _nw116[(new BigNumber(3)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(true, false, _799_v7, _803_v11.f29);
          _nw116[(new BigNumber(5)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(6)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_799_v7, _799_v7);
          _nw116[(new BigNumber(8)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(9)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(10)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(11)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(12)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_807_v13, _807_v13);
          _nw116[(new BigNumber(14)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_799_v7, _803_v11.f29);
          _nw116[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_807_v13, _module.__default.safeIndex((_803_v11).f28, new BigNumber((_807_v13).length)), _803_v11.f29);
          _nw116[(new BigNumber(17)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(18)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(19)).toNumber()] = _807_v13;
          _nw116[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_807_v13, _807_v13), _module.__default.safeIndex(new BigNumber(((p0).update(_808_v14, _module.__default.abs(_794_cf17))).cardinality()), new BigNumber((_dafny.Seq.Concat(_807_v13, _807_v13)).length)), _799_v7);
          _809_v15 = _nw116;
          let _810_v16;
          _810_v16 = _dafny.Seq.of(_793_cf18, new BigNumber(579), (_803_v11).f28, (_this).f12, _794_cf17);
          let _811_v17;
          _811_v17 = _dafny.MultiSet.fromElements(false);
          let _rhs87 = _809_v15;
          let _rhs88 = (_module.__default.fm2(globalState)).minus(_796_cf15);
          let _rhs89 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_810_v16, _810_v16), _dafny.Seq.of(_793_cf18));
          let _rhs90 = ((_dafny.MultiSet.fromElements(_799_v7)).IsProperSubsetOf(_811_v17)) === (_dafny.areEqual(_810_v16, _810_v16));
          let _lhs57 = globalState;
          let _lhs58 = globalState;
          _809_v15 = _rhs87;
          _794_cf17 = _rhs88;
          _lhs57.f8 = _rhs89;
          _lhs58.f8 = _rhs90;
          let _812_v19;
          _812_v19 = _dafny.Map.Empty.slice().updateUnsafe((_796_cf15).isLessThanOrEqualTo(_793_cf18),function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of _dafny.IntegerRange(new BigNumber(601), new BigNumber(209))) {
              let _813_v18 = _compr_54;
              if (((new BigNumber(601)).isLessThanOrEqualTo(_813_v18)) && ((_813_v18).isLessThan(new BigNumber(209)))) {
                _coll54.push([(_813_v18).minus(new BigNumber((_810_v16).length)),_799_v7]);
              }
            }
            return _coll54;
          }());
          let _814_v20;
          _814_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_803_v11).f28)),_803_v11.f29);
          _812_v19 = (_812_v19).update(_799_v7, _814_v20);
        } else {
          let _815_v23;
          let _init31 = ((_816_cf18, _817_cf15) => function (_818_i1) {
            return (function () {
              let _coll55 = new _dafny.Map();
              for (const _compr_55 of (_dafny.Seq.of(_816_cf18)).Elements) {
                let _819_v21 = _compr_55;
                if (_dafny.Seq.contains(_dafny.Seq.of(_816_cf18), _819_v21)) {
                  _coll55.push([(_819_v21).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of((_this).f13))).length)),_dafny.Seq.of(_module.D6.create_DC18(false), _module.D6.create_DC18(!(true)))]);
                }
              }
              return _coll55;
            }()).Merge(function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of (_dafny.Set.fromElements((_this).f13, _817_cf15, _817_cf15, _817_cf15, new BigNumber((_dafny.Seq.of(true, false)).length))).Elements) {
                let _820_v22 = _compr_56;
                if ((_dafny.Set.fromElements((_this).f13, _817_cf15, _817_cf15, _817_cf15, new BigNumber((_dafny.Seq.of(true, false)).length))).contains(_820_v22)) {
                  _coll56.push([(_820_v22).plus(_816_cf18),_dafny.Seq.of(_module.D6.create_DC18(false))]);
                }
              }
              return _coll56;
            }());
          })(_793_cf18, _796_cf15);
          let _nw117 = Array((new BigNumber(20)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw117.length); _i0_31++) {
            _nw117[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _815_v23 = _nw117;
          let _821_v24;
          _821_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(455),_module.__default.fm41(globalState));
          let _index86 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_815_v23).length));
          (_815_v23)[_index86] = _821_v24;
          let _index87 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_815_v23).length));
          (_815_v23)[_index87] = (_821_v24).Merge(_821_v24);
          let _822_v25;
          _822_v25 = false;
          let _823_v26;
          let _nw118 = Array((new BigNumber(6)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _module.__default.fm1((_this).f12, _822_v25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), ((_824_v5) => function (_825_i2) {
            return _824_v5;
          })(_797_v5))).length), globalState);
          _nw118[(_dafny.ONE).toNumber()] = _822_v25;
          _nw118[(new BigNumber(2)).toNumber()] = _822_v25;
          _nw118[(new BigNumber(3)).toNumber()] = _822_v25;
          _nw118[(new BigNumber(4)).toNumber()] = _822_v25;
          _nw118[(new BigNumber(5)).toNumber()] = _822_v25;
          _823_v26 = _nw118;
          let _index88 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_823_v26).length));
          (_823_v26)[_index88] = _822_v25;
          let _826_v27;
          _826_v27 = _dafny.MultiSet.fromElements(_796_cf15, new BigNumber(-893));
          let _827_v28;
          let _nw119 = Array((new BigNumber(8)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of((((_826_v27).contains(_794_cf17)) ? ((_826_v27).get(_794_cf17)) : (_794_cf17)))).length);
          _nw119[(_dafny.ONE).toNumber()] = _796_cf15;
          _nw119[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw119[(new BigNumber(3)).toNumber()] = _796_cf15;
          _nw119[(new BigNumber(4)).toNumber()] = new BigNumber(338);
          _nw119[(new BigNumber(5)).toNumber()] = (_this).f13;
          _nw119[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("a"), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.UnicodeFromString("a")).length)), new _dafny.CodePoint('w'.codePointAt(0)))).length);
          _nw119[(new BigNumber(7)).toNumber()] = new BigNumber(625);
          _827_v28 = _nw119;
          let _828_v29;
          let _nw120 = Array((new BigNumber(7)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _827_v28;
          _nw120[(_dafny.ONE).toNumber()] = _827_v28;
          _nw120[(new BigNumber(2)).toNumber()] = _827_v28;
          _nw120[(new BigNumber(3)).toNumber()] = _827_v28;
          _nw120[(new BigNumber(4)).toNumber()] = ((_822_v25) ? (_827_v28) : (_827_v28));
          _nw120[(new BigNumber(5)).toNumber()] = _827_v28;
          _nw120[(new BigNumber(6)).toNumber()] = _827_v28;
          _828_v29 = _nw120;
          let _index89 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_828_v29).length));
          (_828_v29)[_index89] = _827_v28;
          let _index90 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_823_v26).length));
          let _index91 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_828_v29).length));
          let _rhs91 = _822_v25;
          let _rhs92 = !(_822_v25);
          let _rhs93 = _827_v28;
          let _lhs59 = globalState;
          let _lhs60 = _823_v26;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_823_v26).length));
          let _lhs62 = _828_v29;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_828_v29).length));
          _lhs59.f8 = _rhs91;
          _lhs60[_lhs61] = _rhs92;
          _lhs62[_lhs63] = _rhs93;
          let _829_v30;
          _829_v30 = _dafny.Map.Empty.slice().updateUnsafe((_823_v26)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_823_v26).length))],_793_cf18);
          let _830_v31;
          _830_v31 = _dafny.MultiSet.fromElements(_829_v30);
          let _831_v32;
          let _nw121 = new _module.C0();
          _nw121.__ctor(_793_cf18, (_823_v26)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_823_v26).length))], _796_cf15, new BigNumber((_830_v31).cardinality()));
          _831_v32 = _nw121;
          let _832_v33;
          let _nw122 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _832_v33 = _nw122;
          let _833_v34;
          let _nw123 = Array((new BigNumber(4)).toNumber());
          _nw123[(_dafny.ZERO).toNumber()] = _832_v33;
          _nw123[(_dafny.ONE).toNumber()] = _832_v33;
          _nw123[(new BigNumber(2)).toNumber()] = _832_v33;
          _nw123[(new BigNumber(3)).toNumber()] = _832_v33;
          _833_v34 = _nw123;
          let _index92 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_833_v34).length));
          (_833_v34)[_index92] = _832_v33;
          let _index93 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_833_v34).length));
          (_833_v34)[_index93] = _832_v33;
          let _834_v35;
          _834_v35 = _dafny.Seq.of((_module.D0.create_DC2(_796_cf15, _827_v28)).dtor_cf5);
          _796_cf15 = (new BigNumber((_dafny.Seq.Concat(_834_v35, _834_v35)).length)).plus(new BigNumber(92));
        }
        let _835_v36;
        _835_v36 = false;
        let _836_v37;
        _836_v37 = _dafny.Map.Empty.slice().updateUnsafe(_835_v36,(_this).f13);
        let _837_v38;
        _837_v38 = _module.D3.create_DC8(_836_v37);
        let _pat_let_tv17 = _835_v36;
        let _pat_let_tv18 = _793_cf18;
        let _pat_let_tv19 = _836_v37;
        let _838_v39;
        _838_v39 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_793_cf18),function (_pat_let13_0) {
          return function (_841_dt__update__tmp_h1) {
            return function (_pat_let16_0) {
              return function (_842_dt__update_hcf14_h1) {
                return _module.D3.create_DC8(_842_dt__update_hcf14_h1);
              }(_pat_let16_0);
            }(_pat_let_tv19);
          }(_pat_let13_0);
        }(function (_pat_let14_0) {
          return function (_839_dt__update__tmp_h0) {
            return function (_pat_let15_0) {
              return function (_840_dt__update_hcf14_h0) {
                return _module.D3.create_DC8(_840_dt__update_hcf14_h0);
              }(_pat_let15_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv17,_pat_let_tv18));
          }(_pat_let14_0);
        }(_837_v38)));
        _838_v39 = (_838_v39).update(_module.__default.fm2(globalState), _837_v38);
        let _843_v40;
        let _init32 = ((_844_v36) => function (_845_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(_844_v36,_844_v36);
        })(_835_v36);
        let _nw124 = Array((new BigNumber(7)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw124.length); _i0_32++) {
          _nw124[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _843_v40 = _nw124;
        let _846_v41;
        _846_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,_835_v36);
        let _index94 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_843_v40).length));
        (_843_v40)[_index94] = ((_835_v36) ? (_846_v41) : (_846_v41));
        let _index95 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_843_v40).length));
        (_843_v40)[_index95] = (_846_v41).Merge(_module.__default.fm42(_796_cf15, (_this).f12, globalState));
        let _847_v42;
        let _nw125 = Array((new BigNumber(2)).toNumber()).fill(false);
        _847_v42 = _nw125;
        let _index96 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_847_v42).length));
        (_847_v42)[_index96] = ((_dafny.ZERO).minus(_796_cf15)).isEqualTo(_796_cf15);
        let _848_v43;
        let _nw126 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _848_v43 = _nw126;
        let _849_v44;
        _849_v44 = _dafny.Seq.of(new BigNumber((_784_v0).length));
        let _index97 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_847_v42).length));
        let _rhs94 = (_this).f12;
        let _rhs95 = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_848_v43)).length), (_this).f13, _793_cf18, new BigNumber(476), (_this).f13), _module.__default.fm0(globalState)), _dafny.Seq.Concat(_849_v44, _849_v44));
        let _lhs64 = _847_v42;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_847_v42).length));
        _793_cf18 = _rhs94;
        _lhs64[_lhs65] = _rhs95;
      } else if (_source13.is_DC8) {
        let _850___mcc_h4 = (_source13).cf14;
        let _851_cf14 = _850___mcc_h4;
        let _852_v45;
        let _nw127 = Array((new BigNumber(22)).toNumber()).fill([]);
        _852_v45 = _nw127;
        let _853_v46;
        let _nw128 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _853_v46 = _nw128;
        let _index98 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_852_v45).length));
        (_852_v45)[_index98] = _853_v46;
        let _index99 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_852_v45).length));
        (_852_v45)[_index99] = _853_v46;
        let _854_v47;
        let _nw129 = new _module.C0();
        _nw129.__ctor(new BigNumber(-25), !(false), (_this).f12, (new BigNumber((_dafny.Seq.of(true)).length)).minus((_this).f12));
        _854_v47 = _nw129;
        let _855_v48;
        _855_v48 = new _dafny.CodePoint('a'.codePointAt(0));
        let _856_v49;
        _856_v49 = _dafny.Set.fromElements(_855_v48);
        let _857_v50;
        _857_v50 = _dafny.Map.Empty.slice().updateUnsafe((_854_v47).f28,_856_v49);
        let _858_v51;
        _858_v51 = _dafny.Seq.UnicodeFromString("a");
        let _859_v52;
        _859_v52 = _module.D6.create_DC20(_module.__default.fm43((_this).f13, _854_v47.f29, new BigNumber(((((_857_v50).contains(new BigNumber(297))) ? ((_857_v50).get(new BigNumber(297))) : (_856_v49))).length), _858_v51, globalState));
        let _860_v53;
        _860_v53 = _module.D6.create_DC18(false);
        let _861_v54;
        _861_v54 = _dafny.Set.fromElements(_860_v53);
        let _862_v55;
        _862_v55 = _module.D6.create_DC18(_854_v47.f29);
        let _863_v56;
        _863_v56 = _module.D6.create_DC20(_862_v55);
        let _rhs96 = _854_v47.f29;
        let _rhs97 = _dafny.areEqual(_858_v51, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vqio"), _858_v51), _module.__default.safeIndex((_854_v47).f28, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vqio"), _858_v51)).length)), _855_v48));
        let _rhs98 = (_861_v54).IsSubsetOf(_module.__default.fm44(globalState));
        let _rhs99 = _module.D6.create_DC20(_863_v56);
        let _lhs66 = globalState;
        r0 = _rhs96;
        r0 = _rhs97;
        _lhs66.f8 = _rhs98;
        _859_v52 = _rhs99;
        if (_854_v47.f29) {
          let _864_v57;
          _864_v57 = new BigNumber(494);
          let _865_v58;
          _865_v58 = _dafny.Seq.of(_853_v46);
          let _866_v59;
          _866_v59 = _module.D1.create_DC6(_module.D1.create_DC4(_865_v58));
          let _867_v60;
          _867_v60 = _dafny.Seq.of(_866_v59);
          let _868_v61;
          _868_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_867_v60);
          let _869_v62;
          _869_v62 = _dafny.Seq.of(_864_v57);
          let _870_v63;
          _870_v63 = _dafny.Seq.of(new BigNumber((_869_v62).length));
          let _871_v64;
          _871_v64 = _dafny.Seq.of(new BigNumber((_869_v62).length), (_854_v47).f28);
          let _872_v65;
          _872_v65 = _module.D4.create_DC12(!(_module.__default.fm1((_this).f12, _module.__default.fm1(new BigNumber((_870_v63).length), (_854_v47).fm7(globalState), _module.__default.fm2(globalState), globalState), new BigNumber((_871_v64).length), globalState)), new BigNumber((_851_cf14).length), (_854_v47).f28);
          let _rhs100 = ((_854_v47).f28).multipliedBy(new BigNumber(571));
          let _rhs101 = ((_this).f13).minus((_dafny.ZERO).minus((_854_v47).f28));
          let _rhs102 = (_872_v65).dtor_cf22;
          let _rhs103 = _868_v61;
          let _rhs104 = _858_v51;
          _864_v57 = _rhs100;
          _864_v57 = _rhs101;
          _864_v57 = _rhs102;
          _868_v61 = _rhs103;
          _858_v51 = _rhs104;
          let _873_v66;
          _873_v66 = _module.D5.create_DC15((_854_v47).f28, (_854_v47).f28, new BigNumber(156));
          let _pat_let_tv20 = _854_v47;
          let _pat_let_tv21 = _851_cf14;
          let _pat_let_tv22 = _854_v47;
          let _pat_let_tv23 = _784_v0;
          let _874_v67;
          let _nw130 = Array((new BigNumber(5)).toNumber());
          _nw130[(_dafny.ZERO).toNumber()] = function (_pat_let17_0) {
            return function (_877_dt__update__tmp_h3) {
              return function (_pat_let20_0) {
                return function (_878_dt__update_hcf26_h1) {
                  return function (_pat_let21_0) {
                    return function (_879_dt__update_hcf28_h0) {
                      return _module.D5.create_DC15(_878_dt__update_hcf26_h1, (_877_dt__update__tmp_h3).dtor_cf27, _879_dt__update_hcf28_h0);
                    }(_pat_let21_0);
                  }((_pat_let_tv22).f28);
                }(_pat_let20_0);
              }((_dafny.ZERO).minus(new BigNumber((_pat_let_tv21).length)));
            }(_pat_let17_0);
          }(function (_pat_let18_0) {
            return function (_875_dt__update__tmp_h2) {
              return function (_pat_let19_0) {
                return function (_876_dt__update_hcf26_h0) {
                  return _module.D5.create_DC15(_876_dt__update_hcf26_h0, (_875_dt__update__tmp_h2).dtor_cf27, (_875_dt__update__tmp_h2).dtor_cf28);
                }(_pat_let19_0);
              }((_pat_let_tv20).f28);
            }(_pat_let18_0);
          }(_873_v66));
          _nw130[(_dafny.ONE).toNumber()] = _873_v66;
          _nw130[(new BigNumber(2)).toNumber()] = _module.D5.create_DC15((_854_v47).f28, (_this).f12, (_854_v47).f28);
          _nw130[(new BigNumber(3)).toNumber()] = _module.D5.create_DC15((_this).f13, (_854_v47).f28, new BigNumber(119));
          _nw130[(new BigNumber(4)).toNumber()] = function (_pat_let22_0) {
            return function (_880_dt__update__tmp_h4) {
              return function (_pat_let23_0) {
                return function (_881_dt__update_hcf28_h1) {
                  return _module.D5.create_DC15((_880_dt__update__tmp_h4).dtor_cf26, (_880_dt__update__tmp_h4).dtor_cf27, _881_dt__update_hcf28_h1);
                }(_pat_let23_0);
              }((_dafny.ZERO).minus(new BigNumber((_pat_let_tv23).length)));
            }(_pat_let22_0);
          }(_873_v66);
          _874_v67 = _nw130;
          let _rhs105 = _874_v67;
          let _rhs106 = _854_v47.f29;
          let _rhs107 = (_870_v63)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_870_v63).length))];
          _874_v67 = _rhs105;
          r0 = _rhs106;
          _864_v57 = _rhs107;
          let _882_v68;
          _882_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-872),_853_v46);
          let _883_v69;
          _883_v69 = _module.D3.create_DC9((_this).f13, _785_v1, new BigNumber((_871_v64).length), (_854_v47).f28);
          r1 = (((_882_v68).contains((_869_v62)[_module.__default.safeIndex((_883_v69).dtor_cf17, new BigNumber((_869_v62).length))])) ? ((_882_v68).get((_869_v62)[_module.__default.safeIndex((_883_v69).dtor_cf17, new BigNumber((_869_v62).length))])) : ((_865_v58)[_module.__default.safeIndex((_854_v47).f28, new BigNumber((_865_v58).length))]));
          (globalState).f8 = _854_v47.f29;
          let _884_v70;
          _884_v70 = _module.D1.create_DC4(_dafny.Seq.Concat(_865_v58, _865_v58));
          _884_v70 = _884_v70;
        } else {
          _855_v48 = _module.__default.fm45((_this).f12, globalState);
          let _885_v71;
          _885_v71 = _dafny.Set.fromElements(!(_854_v47.f29), true);
          let _886_v72;
          _886_v72 = _dafny.Set.fromElements((_854_v47).f28, new BigNumber((_851_cf14).length), (_854_v47).f28, (_this).f13, _module.__default.fm2(globalState));
          _858_v51 = _dafny.Seq.Concat((_854_v47).fm8(new BigNumber((_885_v71).length), _dafny.Map.Empty.slice().updateUnsafe(_855_v48,_854_v47.f29), _886_v72, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), function (_887_i4) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }));
          let _888_v73;
          _888_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_854_v47).f28),_854_v47.f29);
          let _889_v74;
          _889_v74 = _dafny.Seq.of(true);
          let _890_v75;
          _890_v75 = _dafny.MultiSet.fromElements((((_888_v73).contains((_854_v47).f28)) ? ((_888_v73).get((_854_v47).f28)) : (_854_v47.f29)), (_854_v47).fm7(globalState), _854_v47.f29, (_889_v74)[_module.__default.safeIndex((_this).f12, new BigNumber((_889_v74).length))], _854_v47.f29);
          let _891_v76;
          _891_v76 = _module.D7.create_DC21((_890_v75).update(_854_v47.f29, _module.__default.abs((_this).f13)));
          _890_v75 = (_891_v76).dtor_cf39;
          let _892_v77;
          _892_v77 = new BigNumber(239);
          let _893_v78;
          let _nw131 = Array((new BigNumber(6)).toNumber()).fill(false);
          _893_v78 = _nw131;
          let _894_v79;
          _894_v79 = _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_854_v47.f29,_854_v47.f29)).update(_854_v47.f29, _854_v47.f29));
          let _895_v80;
          _895_v80 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ceucd"), _858_v51);
          let _896_v81;
          _896_v81 = _dafny.Seq.of(_893_v78, _893_v78);
          let _897_v82;
          _897_v82 = _dafny.Map.Empty.slice().updateUnsafe(_854_v47.f29,true);
          let _898_v83;
          _898_v83 = _dafny.MultiSet.fromElements(_897_v82);
          let _rhs108 = _dafny.Seq.of(!((_module.D7.create_DC23(_852_v45, _889_v74, _854_v47.f29, new BigNumber((_895_v80).length))).dtor_cf44), _854_v47.f29, !(_854_v47.f29) || (_854_v47.f29));
          let _rhs109 = ((_854_v47.f29) ? (false) : ((_this).fm7(globalState)));
          let _rhs110 = _module.__default.fm2(globalState);
          let _rhs111 = (_896_v81)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f12), new BigNumber((_896_v81).length))];
          let _rhs112 = ((_898_v83)).Intersect(_dafny.MultiSet.fromElements(_897_v82));
          let _lhs67 = _854_v47;
          _889_v74 = _rhs108;
          _lhs67.f29 = _rhs109;
          _892_v77 = _rhs110;
          _893_v78 = _rhs111;
          _894_v79 = _rhs112;
          let _899_v84;
          let _nw132 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
          _899_v84 = _nw132;
          let _900_v85;
          _900_v85 = _dafny.Map.Empty.slice().updateUnsafe(_899_v84,_889_v74);
          _889_v74 = (((_900_v85).contains(_899_v84)) ? ((_900_v85).get(_899_v84)) : (_dafny.Seq.of(_854_v47.f29, _854_v47.f29, _854_v47.f29)));
        }
      } else {
        let _901___mcc_h5 = (_source13).cf19;
        let _902_cf19 = _901___mcc_h5;
        let _903_v86;
        _903_v86 = new BigNumber(-242);
        let _904_v87;
        let _nw133 = Array((new BigNumber(28)).toNumber()).fill(false);
        _904_v87 = _nw133;
        let _index100 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_904_v87).length));
        (_904_v87)[_index100] = true;
        let _905_v88;
        _905_v88 = _dafny.Seq.UnicodeFromString("ghpgbut");
        let _index101 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_904_v87).length));
        let _rhs113 = (_this).f12;
        let _rhs114 = (_module.__default.safeModuloInt(new BigNumber(-326), (_this).f13)).multipliedBy((new BigNumber(-870)).plus(_module.__default.fm2(globalState)));
        let _rhs115 = (new BigNumber((_905_v88).length)).isLessThanOrEqualTo(_903_v86);
        let _lhs68 = _904_v87;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_904_v87).length));
        _903_v86 = _rhs113;
        _903_v86 = _rhs114;
        _lhs68[_lhs69] = _rhs115;
        _904_v87 = _904_v87;
        let _906_v89;
        let _nw134 = new _module.C2();
        _nw134.__ctor((((_904_v87)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_904_v87).length))]) ? (new BigNumber(-247)) : (new BigNumber(-417))), _dafny.Seq.Concat(_905_v88, _dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_907_v88) => function (_908_i5) {
          return (_907_v88)[_module.__default.safeIndex((_this).f12, new BigNumber((_907_v88).length))];
        })(_905_v88))), (_this).f13, _903_v86, (new BigNumber(463)).isLessThan(_903_v86));
        _906_v89 = _nw134;
        _906_v89 = _906_v89;
        let _909_v90;
        _909_v90 = _dafny.Seq.of(_906_v89.f11, false, (_904_v87)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_904_v87).length))]);
        let _910_v91;
        _910_v91 = _dafny.MultiSet.fromElements(_909_v90, _909_v90, _dafny.Seq.of(_906_v89.f11));
        let _911_v92;
        _911_v92 = _dafny.Map.Empty.slice().updateUnsafe(_910_v91,(_this).f12);
        _911_v92 = (_911_v92).update(_910_v91, (_this).f12);
      }
      let _912_v93;
      _912_v93 = false;
      let _913_v94;
      _913_v94 = _dafny.MultiSet.fromElements(_912_v93, _912_v93);
      let _914_v95;
      _914_v95 = _dafny.Seq.UnicodeFromString("nfsiu");
      let _915_v96;
      _915_v96 = _dafny.Seq.of(_912_v93, _912_v93);
      let _916_v97;
      _916_v97 = _dafny.Seq.of(new BigNumber((_915_v96).length));
      let _917_v98;
      _917_v98 = _dafny.Set.fromElements(new BigNumber((_916_v97).length), (_this).f13);
      let _918_v99;
      _918_v99 = _module.D6.create_DC19(true, !((_913_v94).IsDisjointFrom(_913_v94)), (_917_v98).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_919_i6) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}))).update((_this).f13, _914_v95)).length), (_this).f12, (_this).f13)), _912_v93, _dafny.Seq.Concat(_915_v96, _915_v96));
      let _source14 = _918_v99;
      if (_source14.is_DC17) {
        let _920___mcc_h6 = (_source14).cf30;
        let _921___mcc_h7 = (_source14).cf31;
        let _922_cf31 = _921___mcc_h7;
        let _923_cf30 = _920___mcc_h6;
        _922_cf31 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_this).f13, new BigNumber(403))).multipliedBy((_this).f12));
        let _rhs116 = ((_this).f12).isLessThan((_916_v97)[_module.__default.safeIndex((_this).f13, new BigNumber((_916_v97).length))]);
        let _rhs117 = new BigNumber((_916_v97).length);
        let _rhs118 = new BigNumber((_917_v98).length);
        let _lhs70 = globalState;
        _lhs70.f8 = _rhs116;
        _922_cf31 = _rhs117;
        _922_cf31 = _rhs118;
        let _924_v100;
        let _init33 = ((_925_v96) => function (_926_i7) {
          return (_926_i7).minus(new BigNumber((_925_v96).length));
        })(_915_v96);
        let _nw135 = Array((new BigNumber(25)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw135.length); _i0_33++) {
          _nw135[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _924_v100 = _nw135;
        let _927_v101;
        _927_v101 = _dafny.Seq.of(_924_v100, _924_v100);
        let _928_v102;
        _928_v102 = _dafny.Map.Empty.slice().updateUnsafe(_923_cf30,_924_v100);
        let _929_v103;
        let _nw136 = Array((new BigNumber(29)).toNumber());
        _nw136[(_dafny.ZERO).toNumber()] = _924_v100;
        _nw136[(_dafny.ONE).toNumber()] = _924_v100;
        _nw136[(new BigNumber(2)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(3)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(4)).toNumber()] = (_927_v101)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_927_v101).length))];
        _nw136[(new BigNumber(5)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(6)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(7)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(8)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(9)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(10)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(11)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(12)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(13)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(14)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(15)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(16)).toNumber()] = (_927_v101)[_module.__default.safeIndex(new BigNumber((_917_v98).length), new BigNumber((_927_v101).length))];
        _nw136[(new BigNumber(17)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(18)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(19)).toNumber()] = (((_928_v102).contains(!(!(_912_v93)))) ? ((_928_v102).get(!(!(_912_v93)))) : (_924_v100));
        _nw136[(new BigNumber(20)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(21)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(22)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(23)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(24)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(25)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(26)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(27)).toNumber()] = _924_v100;
        _nw136[(new BigNumber(28)).toNumber()] = _924_v100;
        _929_v103 = _nw136;
        let _930_v104;
        _930_v104 = _module.D7.create_DC22((_this).f12, _929_v103);
        let _931_v105;
        _931_v105 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,((_912_v93) ? (_module.D7.create_DC22(_922_cf31, _929_v103)) : (_930_v104)));
        _931_v105 = (_931_v105).update((_916_v97)[_module.__default.safeIndex((((_913_v94).contains(_912_v93)) ? ((_913_v94).get(_912_v93)) : ((_this).f12)), new BigNumber((_916_v97).length))], _930_v104);
        let _932_v106;
        _932_v106 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-733));
        let _933_v107;
        _933_v107 = new _dafny.CodePoint('q'.codePointAt(0));
        let _934_v108;
        _934_v108 = _dafny.Map.Empty.slice().updateUnsafe(_933_v107,_912_v93);
        let _935_v109;
        _935_v109 = _module.D7.create_DC23(_929_v103, _915_v96, _912_v93, (_this).f13);
        let _936_v110;
        let _nw137 = new _module.C2();
        _nw137.__ctor((new BigNumber((_932_v106).length)).multipliedBy(_922_cf31), (_this).fm8((_this).f13, _934_v108, _dafny.Set.fromElements((_this).f12), globalState), (_935_v109).dtor_cf45, new BigNumber(4), (_915_v96)[_module.__default.safeIndex((_this).f13, new BigNumber((_915_v96).length))]);
        _936_v110 = _nw137;
      } else if (_source14.is_DC18) {
        let _937___mcc_h8 = (_source14).cf32;
        let _938_cf32 = _937___mcc_h8;
        let _939_v111;
        _939_v111 = _dafny.Set.fromElements(_938_cf32, _938_cf32);
        _939_v111 = (_939_v111).Intersect(_939_v111);
        _938_cf32 = _dafny.areEqual(_dafny.Seq.of(_912_v93), _915_v96);
        let _940_v112;
        _940_v112 = new BigNumber(158);
        _940_v112 = _940_v112;
        let _941_v113;
        _941_v113 = _dafny.Map.Empty.slice().updateUnsafe(_914_v95,_915_v96);
        let _942_v115;
        _942_v115 = _dafny.Map.Empty.slice().updateUnsafe(_938_cf32,(_this).f12);
        let _943_v116;
        _943_v116 = new _dafny.CodePoint('h'.codePointAt(0));
        let _944_v117;
        _944_v117 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_914_v95, _module.__default.safeIndex(new BigNumber((_942_v115).length), new BigNumber((_914_v95).length)), _943_v116));
        _941_v113 = function () {
          let _coll57 = new _dafny.Map();
          for (const _compr_57 of (_944_v117).Elements) {
            let _945_v114 = _compr_57;
            if ((_944_v117).contains(_945_v114)) {
              _coll57.push([_945_v114,_915_v96]);
            }
          }
          return _coll57;
        }();
      } else if (_source14.is_DC19) {
        let _946___mcc_h9 = (_source14).cf33;
        let _947___mcc_h10 = (_source14).cf34;
        let _948___mcc_h11 = (_source14).cf35;
        let _949___mcc_h12 = (_source14).cf36;
        let _950___mcc_h13 = (_source14).cf37;
        let _951_cf37 = _950___mcc_h13;
        let _952_cf36 = _949___mcc_h12;
        let _953_cf35 = _948___mcc_h11;
        let _954_cf34 = _947___mcc_h10;
        let _955_cf33 = _946___mcc_h9;
        let _956_v118;
        let _nw138 = new _module.C0();
        _nw138.__ctor((new BigNumber(636)).minus(new BigNumber(402)), _912_v93, (_this).f13, _module.__default.safeDivisionInt(new BigNumber(788), new BigNumber(350)));
        _956_v118 = _nw138;
        let _957_v119;
        _957_v119 = new BigNumber(716);
        _957_v119 = (_this).f12;
        let _958_v120;
        let _nw139 = Array((new BigNumber(16)).toNumber()).fill([]);
        _958_v120 = _nw139;
        let _959_v121;
        _959_v121 = _module.D7.create_DC22((_956_v118).f28, ((_912_v93) ? (_958_v120) : (_958_v120)));
        let _rhs119 = _952_cf36;
        let _rhs120 = _916_v97;
        let _rhs121 = _959_v121;
        r0 = _rhs119;
        _916_v97 = _rhs120;
        _959_v121 = _rhs121;
        let _960_v122;
        let _nw140 = Array((new BigNumber(20)).toNumber()).fill(false);
        _960_v122 = _nw140;
        let _961_v123;
        _961_v123 = _dafny.Seq.of(_960_v122, _960_v122);
        _960_v122 = (_961_v123)[_module.__default.safeIndex(new BigNumber((_914_v95).length), new BigNumber((_961_v123).length))];
      } else if (_source14.is_DC16) {
        let _962___mcc_h14 = (_source14).cf29;
        let _963_cf29 = _962___mcc_h14;
        let _964_v124;
        _964_v124 = _dafny.Map.Empty.slice().updateUnsafe(!(_912_v93),_912_v93);
        let _965_v125;
        _965_v125 = _dafny.MultiSet.fromElements(_964_v124, _964_v124);
        let _966_v126;
        _966_v126 = _965_v125;
        _966_v126 = _966_v126;
        let _967_v127;
        let _nw141 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _967_v127 = _nw141;
        let _index102 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_967_v127).length));
        (_967_v127)[_index102] = (_this).f13;
        let _968_v128;
        _968_v128 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_912_v93, _912_v93, _912_v93)).length), (_this).f13);
        let _969_v129;
        _969_v129 = new _dafny.CodePoint('d'.codePointAt(0));
        let _970_v130;
        _970_v130 = _dafny.Set.fromElements(_969_v129, _969_v129);
        let _971_v131;
        _971_v131 = _dafny.Map.Empty.slice().updateUnsafe(_970_v130,new BigNumber((_915_v96).length));
        let _972_v132;
        _972_v132 = _module.D4.create_DC11(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f12));
        let _index103 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_967_v127).length));
        (_967_v127)[_index103] = (((_968_v128).contains((new BigNumber(584)).minus(new BigNumber(((_971_v131).update(_970_v130, new BigNumber((_dafny.Seq.of(_972_v132, _module.D4.create_DC11((_784_v0).update((((_913_v94).contains(_912_v93)) ? ((_913_v94).get(_912_v93)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ajqx")).length))), (_this).f13)))).length))).length)))) ? ((_968_v128).get((new BigNumber(584)).minus(new BigNumber(((_971_v131).update(_970_v130, new BigNumber((_dafny.Seq.of(_972_v132, _module.D4.create_DC11((_784_v0).update((((_913_v94).contains(_912_v93)) ? ((_913_v94).get(_912_v93)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ajqx")).length))), (_this).f13)))).length))).length)))) : (_module.__default.safeModuloInt((_this).f13, (_this).f12)));
        (globalState).f8 = _912_v93;
        let _index104 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_967_v127).length));
        let _rhs122 = _916_v97;
        let _rhs123 = (_967_v127)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_967_v127).length))];
        let _lhs71 = globalState;
        let _lhs72 = _967_v127;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_967_v127).length));
        _lhs71.f5 = _rhs122;
        _lhs72[_lhs73] = _rhs123;
      } else {
        let _973___mcc_h15 = (_source14).cf38;
        let _974_cf38 = _973___mcc_h15;
        let _975_v133;
        let _nw142 = new _module.C0();
        _nw142.__ctor((_this).f13, _module.__default.fm1((_this).f13, _912_v93, (_this).f13, globalState), new BigNumber(960), ((_this).f13).minus((_this).f13));
        _975_v133 = _nw142;
        let _976_v134;
        _976_v134 = _dafny.Map.Empty.slice().updateUnsafe(_912_v93,((_975_v133).f28).plus((_this).f13));
        let _977_v135;
        _977_v135 = _dafny.Map.Empty.slice().updateUnsafe((_975_v133).f28,_976_v134);
        _976_v134 = ((((_977_v135).contains((_975_v133).f28)) ? ((_977_v135).get((_975_v133).f28)) : (_976_v134))).update(false, ((_this).f12).plus(_module.__default.fm2(globalState)));
        (_975_v133).f29 = (_912_v93) === ((_912_v93) || (_912_v93));
        let _978_v136;
        _978_v136 = new BigNumber(758);
        let _979_v137;
        _979_v137 = _dafny.Set.fromElements(_975_v133.f29, _975_v133.f29, _912_v93);
        let _980_v138;
        _980_v138 = _dafny.Seq.of((_979_v137).Difference(_dafny.Set.fromElements(false, _912_v93)), _979_v137);
        _978_v136 = new BigNumber((_980_v138).length);
      }
      r0 = _912_v93;
      let _981_v139;
      _981_v139 = new BigNumber(974);
      _981_v139 = (_this).f12;
      let _982_v140;
      _982_v140 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.FromArray(_916_v97)).update((_this).f12, _module.__default.abs((_this).f13)),new BigNumber((_916_v97).length));
      _981_v139 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_916_v97, _module.__default.safeIndex(new BigNumber(6), new BigNumber((_916_v97).length)), (_this).f13), _dafny.Seq.of(_981_v139, (_dafny.ZERO).minus(new BigNumber((_982_v140).length)), (_this).f13))).length);
      let _983_v141;
      _983_v141 = _module.D6.create_DC20(_module.D6.create_DC17(_912_v93, (_this).f12));
      let _984_v142;
      _984_v142 = new _dafny.CodePoint('r'.codePointAt(0));
      let _985_v143;
      _985_v143 = _module.D0.create_DC1((_this).f12, _914_v95, (_this).f12, true);
      let _986_v144;
      _986_v144 = _dafny.Seq.of(_dafny.Seq.of(_984_v142), (_985_v143).dtor_cf2);
      let _987_v145;
      let _init34 = function (_988_i8) {
        return (_988_i8).minus((_this).f13);
      };
      let _nw143 = Array((new BigNumber(24)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw143.length); _i0_34++) {
        _nw143[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _987_v145 = _nw143;
      let _989_v146;
      let _nw144 = Array((new BigNumber(5)).toNumber());
      _nw144[(_dafny.ZERO).toNumber()] = _987_v145;
      _nw144[(_dafny.ONE).toNumber()] = _987_v145;
      _nw144[(new BigNumber(2)).toNumber()] = _987_v145;
      _nw144[(new BigNumber(3)).toNumber()] = _987_v145;
      _nw144[(new BigNumber(4)).toNumber()] = _987_v145;
      _989_v146 = _nw144;
      let _990_v147;
      _990_v147 = _module.D7.create_DC22((_this).f13, _989_v146);
      let _index105 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_987_v145).length));
      (_987_v145)[_index105] = (_916_v97)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_912_v93,_990_v147)).length), new BigNumber((_916_v97).length))];
      let _index106 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_987_v145).length));
      let _rhs124 = !(_912_v93);
      let _rhs125 = ((_this).f13).isLessThan((_this).f12);
      let _rhs126 = _module.__default.fm43((_dafny.ZERO).minus((_this).f13), (_912_v93) && (_912_v93), (_this).f12, _dafny.Seq.Concat(_914_v95, _914_v95), globalState);
      let _rhs127 = _dafny.Seq.of(_914_v95);
      let _rhs128 = _981_v139;
      let _lhs74 = _987_v145;
      let _lhs75 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_987_v145).length));
      _912_v93 = _rhs124;
      _912_v93 = _rhs125;
      _983_v141 = _rhs126;
      _986_v144 = _rhs127;
      _lhs74[_lhs75] = _rhs128;
      r0 = _912_v93;
      r1 = _987_v145;
      let _991_v148;
      let _nw145 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
      _991_v148 = _nw145;
      r2 = _991_v148;
      let _992_v149;
      _992_v149 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(898)), ((_993_v94) => function (_994_i9) {
        return new BigNumber((_dafny.Seq.of(_module.D7.create_DC21(_993_v94), _module.D7.create_DC21(_993_v94), _module.D7.create_DC21(_993_v94))).length);
      })(_913_v94));
      r3 = _dafny.Set.fromElements(_992_v149);
      return [r0, r1, r2, r3];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f11 = false;
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    __ctor(f25, f11) {
      let _this = this;
      (_this)._f25 = f25;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1((_this).f25, _dafny.Seq.UnicodeFromString("qsdsks"), (_this).f25, _this.f11);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f25).isLessThan((_this).f25);
    };
    fm28(p0, globalState) {
      let _this = this;
      return !(_this.f11);
    };
    fm29(p0, p1, globalState) {
      let _this = this;
      if (_this.f11) {
        return (!(_this.f11)) && (_this.f11);
      } else {
        return _this.f11;
      }
    };
    m21(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _995_v0;
      _995_v0 = _dafny.Seq.of(false);
      let _996_v1;
      _996_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_995_v0, _dafny.Seq.of(true, _this.f11)),p2);
      _996_v1 = (_dafny.Map.Empty.slice().updateUnsafe(_995_v0,new BigNumber((_995_v0).length))).Merge((_996_v1).Merge(_996_v1));
      let _997_v3;
      _997_v3 = _dafny.Map.Empty.slice().updateUnsafe(p2,function () {
        let _coll58 = new _dafny.Map();
        for (const _compr_58 of _dafny.IntegerRange(new BigNumber(665), new BigNumber(886))) {
          let _998_v2 = _compr_58;
          if (((new BigNumber(665)).isLessThanOrEqualTo(_998_v2)) && ((_998_v2).isLessThan(new BigNumber(886)))) {
            _coll58.push([_module.__default.safeModuloInt(_998_v2, p3),p3]);
          }
        }
        return _coll58;
      }());
      let _999_v4;
      _999_v4 = _module.D3.create_DC9((_this).f25, _997_v3, (_this).f25, (_this).f25);
      let _1000_v5;
      _1000_v5 = _module.D3.create_DC10(_999_v4);
      let _1001_i0;
      _1001_i0 = _dafny.ZERO;
      L3: {
        let _pat_let_tv24 = p2;
        let _pat_let_tv25 = p0;
        let _pat_let_tv26 = p0;
        while (function (_source15) {
          if (_source15.is_DC9) {
            let _1007___mcc_h0 = (_source15).cf15;
            let _1008___mcc_h1 = (_source15).cf16;
            let _1009___mcc_h2 = (_source15).cf17;
            let _1010___mcc_h3 = (_source15).cf18;
            let _1011_cf18 = _1010___mcc_h3;
            let _1012_cf17 = _1009___mcc_h2;
            let _1013_cf16 = _1008___mcc_h1;
            let _1014_cf15 = _1007___mcc_h0;
            return (new BigNumber((function () {
              let _coll59 = new _dafny.Set();
              for (const _compr_59 of _dafny.IntegerRange(new BigNumber(649), new BigNumber(528))) {
                let _1015_v6 = _compr_59;
                if (((new BigNumber(649)).isLessThanOrEqualTo(_1015_v6)) && ((_1015_v6).isLessThan(new BigNumber(528)))) {
                  _coll59.add((_1015_v6).plus(_pat_let_tv24));
                }
              }
              return _coll59;
            }()).length)).isEqualTo((_this).f25);
          } else if (_source15.is_DC8) {
            let _1016___mcc_h4 = (_source15).cf14;
            let _1017_cf14 = _1016___mcc_h4;
            return _this.f11;
          } else {
            let _1018___mcc_h5 = (_source15).cf19;
            let _1019_cf19 = _1018___mcc_h5;
            return !_dafny.areEqual(_pat_let_tv25, _pat_let_tv26);
          }
        }(_1000_v5)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1001_i0)) {
              break L3;
            }
            _1001_i0 = (_1001_i0).plus(_dafny.ONE);
            let _1002_v7;
            _1002_v7 = _module.D0.create_DC3(p2, _this.f11);
            r1 = ((_1002_v7).dtor_cf7).multipliedBy((_this).f25);
            let _1003_v8;
            _1003_v8 = _dafny.Seq.of(_module.__default.fm2(globalState));
            r1 = (new BigNumber(333)).minus(new BigNumber((_1003_v8).length));
            let _1004_v9;
            let _nw146 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
            _1004_v9 = _nw146;
            let _1005_v10;
            _1005_v10 = _dafny.Seq.UnicodeFromString("ccejxpoh");
            let _1006_v11;
            _1006_v11 = _dafny.Seq.of(_1004_v9);
            let _rhs129 = (_dafny.ZERO).minus(_module.__default.fm2(globalState));
            let _rhs130 = (_1006_v11)[_module.__default.safeIndex((new BigNumber((_1005_v10).length)).multipliedBy(p3), new BigNumber((_1006_v11).length))];
            let _rhs131 = (p2).minus(p3);
            let _rhs132 = _module.__default.safeModuloInt(p2, _module.__default.fm2(globalState));
            let _rhs133 = _dafny.Seq.Concat(p0, p0);
            r0 = _rhs129;
            _1004_v9 = _rhs130;
            r1 = _rhs131;
            r1 = _rhs132;
            _1005_v10 = _rhs133;
            r1 = new BigNumber((_1005_v10).length);
          }
        }
      }
      let _1020_v12;
      _1020_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p3, p3, p2),p3);
      let _1021_v13;
      _1021_v13 = _dafny.MultiSet.fromElements(_this.f11, (_this).fm6(_this.f11, (_this).f25, new BigNumber((_1020_v12).length), globalState));
      if (!((_module.__default.fm30(_this.f11, !(_this.f11), new BigNumber((_1021_v13).cardinality()), globalState)).IsProperSubsetOf(_module.__default.fm30(_this.f11, false, (_this).f25, globalState)))) {
        let _1022_v14;
        _1022_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),new BigNumber(970));
        let _1023_v15;
        _1023_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
        let _1024_v16;
        let _nw147 = Array((new BigNumber(14)).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = new BigNumber((_1022_v14).length);
        _nw147[(_dafny.ONE).toNumber()] = p3;
        _nw147[(new BigNumber(2)).toNumber()] = (_this).f25;
        _nw147[(new BigNumber(3)).toNumber()] = p2;
        _nw147[(new BigNumber(4)).toNumber()] = p3;
        _nw147[(new BigNumber(5)).toNumber()] = p3;
        _nw147[(new BigNumber(6)).toNumber()] = p2;
        _nw147[(new BigNumber(7)).toNumber()] = new BigNumber(-549);
        _nw147[(new BigNumber(8)).toNumber()] = new BigNumber((_1022_v14).length);
        _nw147[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1023_v15).length));
        _nw147[(new BigNumber(10)).toNumber()] = p2;
        _nw147[(new BigNumber(11)).toNumber()] = p2;
        _nw147[(new BigNumber(12)).toNumber()] = new BigNumber(-768);
        _nw147[(new BigNumber(13)).toNumber()] = p2;
        _1024_v16 = _nw147;
        let _1025_v17;
        let _nw148 = Array((new BigNumber(7)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = _1024_v16;
        _nw148[(_dafny.ONE).toNumber()] = _1024_v16;
        _nw148[(new BigNumber(2)).toNumber()] = _1024_v16;
        _nw148[(new BigNumber(3)).toNumber()] = _1024_v16;
        _nw148[(new BigNumber(4)).toNumber()] = _1024_v16;
        _nw148[(new BigNumber(5)).toNumber()] = _1024_v16;
        _nw148[(new BigNumber(6)).toNumber()] = _1024_v16;
        _1025_v17 = _nw148;
        let _index107 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length));
        (_1025_v17)[_index107] = _1024_v16;
        let _index108 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length));
        (_1025_v17)[_index108] = _1024_v16;
        let _1026_v18;
        _1026_v18 = _dafny.Seq.UnicodeFromString("o");
        let _1027_v19;
        _1027_v19 = _dafny.Seq.of(p0);
        let _1028_v20;
        _1028_v20 = _dafny.Seq.of((_1027_v19)[_module.__default.safeIndex((_module.D0.create_DC2(p3, _1024_v16)).dtor_cf5, new BigNumber((_1027_v19).length))], _1026_v18, _dafny.Seq.update(_1026_v18, _module.__default.safeIndex(p3, new BigNumber((_1026_v18).length)), (p0)[_module.__default.safeIndex(_module.__default.fm2(globalState), new BigNumber((p0).length))]));
        _1026_v18 = (_1028_v20)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f25, p3), new BigNumber((_1028_v20).length))];
        let _source16 = _module.D0.create_DC3(p2, _this.f11);
        if (_source16.is_DC1) {
          let _1029___mcc_h6 = (_source16).cf1;
          let _1030___mcc_h7 = (_source16).cf2;
          let _1031___mcc_h8 = (_source16).cf3;
          let _1032___mcc_h9 = (_source16).cf4;
          let _1033_cf4 = _1032___mcc_h9;
          let _1034_cf3 = _1031___mcc_h8;
          let _1035_cf2 = _1030___mcc_h7;
          let _1036_cf1 = _1029___mcc_h6;
          let _1037_v21;
          _1037_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber(167));
          let _1038_v22;
          _1038_v22 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gmgibfqjo")).length));
          let _rhs134 = ((((_1037_v21).contains(_1033_cf4)) ? ((_1037_v21).get(_1033_cf4)) : (new BigNumber((_995_v0).length)))).plus(new BigNumber((p0).length));
          let _rhs135 = ((_1021_v13).update(_this.f11, _module.__default.abs((_1038_v22)[_module.__default.safeIndex(new BigNumber((_1035_cf2).length), new BigNumber((_1038_v22).length))]))).IsDisjointFrom(_1021_v13);
          let _rhs136 = ((_1033_cf4) ? (_this.f11) : (_this.f11));
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          r0 = _rhs134;
          _lhs76.f8 = _rhs135;
          _lhs77.f8 = _rhs136;
          let _1039_v23;
          _1039_v23 = _dafny.Set.fromElements(_1034_cf3);
          let _1040_v24;
          _1040_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1033_cf4,_module.__default.fm0(globalState));
          _1033_cf4 = !(new BigNumber((_1039_v23).length)).isEqualTo((_dafny.ZERO).minus(new BigNumber(((((_1040_v24).contains((((_1023_v15).contains(_1033_cf4)) ? ((_1023_v15).get(_1033_cf4)) : (_this.f11)))) ? ((_1040_v24).get((((_1023_v15).contains(_1033_cf4)) ? ((_1023_v15).get(_1033_cf4)) : (_this.f11)))) : (_1038_v22))).length)));
          let _1041_v25;
          let _init35 = ((_1042_v0) => function (_1043_i1) {
            return _1042_v0;
          })(_995_v0);
          let _nw149 = Array((new BigNumber(25)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw149.length); _i0_35++) {
            _nw149[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1041_v25 = _nw149;
          _1041_v25 = _1041_v25;
          _1033_cf4 = _this.f11;
        } else if (_source16.is_DC2) {
          let _1044___mcc_h10 = (_source16).cf5;
          let _1045___mcc_h11 = (_source16).cf6;
          let _1046_cf6 = _1045___mcc_h11;
          let _1047_cf5 = _1044___mcc_h10;
          let _1048_v26;
          _1048_v26 = _module.D3.create_DC9(p3, _997_v3, (_this).f25, (_this).f25);
          _1048_v26 = _module.D3.create_DC9(p2, _module.__default.fm31(globalState), p3, (_dafny.ZERO).minus(_1047_cf5));
          let _1049_v27;
          _1049_v27 = _dafny.Seq.of(_1024_v16);
          let _1050_v28;
          _1050_v28 = _module.D1.create_DC4(_1049_v27);
          let _1051_v29;
          _1051_v29 = _module.D1.create_DC6(_1050_v28);
          _1051_v29 = _1051_v29;
          _1047_cf5 = (((_1021_v13).contains((((_1023_v15).contains(!(_this.f11))) ? ((_1023_v15).get(!(_this.f11))) : (_this.f11)))) ? ((_1021_v13).get((((_1023_v15).contains(!(_this.f11))) ? ((_1023_v15).get(!(_this.f11))) : (_this.f11)))) : (_module.__default.safeModuloInt(_1047_cf5, p2)));
          let _1052_v30;
          _1052_v30 = _dafny.Seq.of(p2);
          let _1053_v31;
          let _nw150 = new _module.C2();
          _nw150.__ctor(((_this.f11) ? (p2) : ((_this).f25)), p0, (_dafny.ZERO).minus((_1052_v30)[_module.__default.safeIndex(p2, new BigNumber((_1052_v30).length))]), new BigNumber(803), !(true));
          _1053_v31 = _nw150;
        } else if (_source16.is_DC3) {
          let _1054___mcc_h12 = (_source16).cf7;
          let _1055___mcc_h13 = (_source16).cf8;
          let _1056_cf8 = _1055___mcc_h13;
          let _1057_cf7 = _1054___mcc_h12;
          _1022_v14 = (_1022_v14).update(p2, p3);
          let _1058_v32;
          let _nw151 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
          _1058_v32 = _nw151;
          _1058_v32 = _1058_v32;
          let _1059_v33;
          let _1060_v34;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _1059_v33 = _out8;
          _1060_v34 = _out9;
          let _arr0 = (_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))];
          let _index109 = _module.__default.safeIndex(new BigNumber(322), new BigNumber(((_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))]).length));
          _arr0[_index109] = (_dafny.ZERO).minus((_this).f25);
          let _arr1 = (_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))];
          let _index110 = _module.__default.safeIndex(new BigNumber(322), new BigNumber(((_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))]).length));
          _arr1[_index110] = (p2).multipliedBy(_1057_cf7);
        } else {
          let _1061___mcc_h14 = (_source16).cf0;
          let _1062_cf0 = _1061___mcc_h14;
          _1026_v18 = p0;
          let _1063_v35;
          _1063_v35 = _module.D0.create_DC1(p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_1064_i2) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}), p3, _this.f11);
          let _1065_v36;
          _1065_v36 = _dafny.Seq.of(p3, p2, p3, new BigNumber((p0).length));
          let _1066_v37;
          _1066_v37 = _module.D0.create_DC2(p2, _1024_v16);
          let _1067_v38;
          let _nw152 = new _module.C2();
          _nw152.__ctor((_this).f25, _dafny.Seq.Concat((_1063_v35).dtor_cf2, _1026_v18), _module.__default.fm2(globalState), (_this).f25, (new BigNumber((_1065_v36).length)).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_1066_v37)).length))));
          _1067_v38 = _nw152;
          let _nw153 = new _module.C2();
          _nw153.__ctor(p2, _1026_v18, ((_this.f11) ? ((_this).f25) : (new BigNumber(143))), (_this).f25, _this.f11);
          _1067_v38 = _nw153;
          let _1068_v39;
          _1068_v39 = _dafny.Set.fromElements(_1067_v38.f11);
          let _1069_v40;
          _1069_v40 = _dafny.MultiSet.fromElements(_1068_v39, _1068_v39, _1068_v39, _1068_v39);
          _1023_v15 = (_1023_v15).update((_1069_v40).IsProperSubsetOf(_1069_v40), false);
          let _1070_v41;
          let _1071_v42;
          let _out10;
          let _out11;
          let _outcollector5 = _module.__default.m0(globalState);
          _out10 = _outcollector5[0];
          _out11 = _outcollector5[1];
          _1070_v41 = _out10;
          _1071_v42 = _out11;
        }
        let _1072_v43;
        _1072_v43 = _dafny.Set.fromElements(p3);
        r0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_1072_v43).length)), p3), new BigNumber((_dafny.Seq.Concat(p0, _1026_v18)).length));
        if (!(_this.f11)) {
          (globalState).f8 = !(!((_this).fm6(!((p2).isEqualTo(p3)), new BigNumber((_dafny.Set.fromElements(_this.f11, false, _this.f11, _this.f11, _this.f11)).length), (p3).plus(new BigNumber(764)), globalState)));
          r0 = (_this).f25;
          (globalState).f8 = (_this.f11) && ((_this).fm29(_this.f11, new BigNumber((_dafny.Seq.of(new BigNumber((p0).length))).length), globalState));
          let _index111 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p1).length));
          (p1)[_index111] = _this.f11;
          let _index112 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p1).length));
          (p1)[_index112] = ((_module.__default.fm1(p2, _this.f11, new BigNumber(803), globalState)) ? (!(!(!(_this.f11)))) : ((_this.f11) === (_this.f11)));
          let _1073_v44;
          _1073_v44 = _dafny.Map.Empty.slice().updateUnsafe(p2,(p1)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p1).length))]);
          let _1074_v45;
          _1074_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v44,!(_this.f11));
          let _1075_v46;
          let _nw154 = Array((new BigNumber(5)).toNumber()).fill(false);
          _1075_v46 = _nw154;
          let _1076_v47;
          _1076_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1074_v45).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(p1)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p1).length))]),!((_this).fm6(!(!(_this.f11)), p2, p3, globalState)))),_1075_v46);
          _1076_v47 = (_1076_v47).update(_1074_v45, _1075_v46);
        } else {
          let _1077_v48;
          _1077_v48 = _dafny.Seq.of(p2);
          r1 = (_1077_v48)[_module.__default.safeIndex(((_this).f25).multipliedBy((_this).f25), new BigNumber((_1077_v48).length))];
          let _1078_v49;
          _1078_v49 = _dafny.Seq.of(_1024_v16, (_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))], _1024_v16);
          let _1079_v50;
          _1079_v50 = _module.D1.create_DC4(_dafny.Seq.of((_1025_v17)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length))], _1024_v16));
          let _pat_let_tv27 = _1078_v49;
          let _pat_let_tv28 = _1078_v49;
          let _1080_v51;
          _1080_v51 = _dafny.Seq.of(_module.D1.create_DC4(_1078_v49), function (_pat_let24_0) {
            return function (_1083_dt__update__tmp_h1) {
              return function (_pat_let27_0) {
                return function (_1084_dt__update_hcf9_h1) {
                  return _module.D1.create_DC4(_1084_dt__update_hcf9_h1);
                }(_pat_let27_0);
              }(_pat_let_tv28);
            }(_pat_let24_0);
          }(function (_pat_let25_0) {
            return function (_1081_dt__update__tmp_h0) {
              return function (_pat_let26_0) {
                return function (_1082_dt__update_hcf9_h0) {
                  return _module.D1.create_DC4(_1082_dt__update_hcf9_h0);
                }(_pat_let26_0);
              }(_pat_let_tv27);
            }(_pat_let25_0);
          }(_1079_v50)));
          _1080_v51 = _1080_v51;
          (_this).f11 = (_module.__default.fm46(_this.f11, (_995_v0)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_995_v0).length))], globalState)).dtor_cf21;
          let _1085_v52;
          let _nw155 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _1085_v52 = _nw155;
          let _1086_v53;
          let _nw156 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _1086_v53 = _nw156;
          let _1087_v54;
          _1087_v54 = _module.D0.create_DC0(_1086_v53);
          let _1088_v55;
          _1088_v55 = _dafny.Seq.of(_1087_v54, _1087_v54);
          let _index113 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1085_v52).length));
          (_1085_v52)[_index113] = _dafny.Map.Empty.slice().updateUnsafe(_1088_v55,_dafny.Seq.UnicodeFromString("va"));
          let _1089_v56;
          _1089_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1088_v55,_1026_v18);
          let _index114 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1085_v52).length));
          (_1085_v52)[_index114] = (_1089_v56).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1088_v55,p0));
          let _index115 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1025_v17).length));
          (_1025_v17)[_index115] = _1024_v16;
        }
      } else {
        r1 = (_this).f25;
        let _1090_v57;
        let _nw157 = new _module.C2();
        _nw157.__ctor((_dafny.ZERO).minus(new BigNumber((_module.__default.fm39(_this.f11, _this.f11, globalState)).length)), _dafny.Seq.UnicodeFromString("njlurjslh"), new BigNumber(688), p2, !(_this.f11));
        _1090_v57 = _nw157;
        let _1091_v58;
        _1091_v58 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f11);
        let _1092_v59;
        _1092_v59 = new _dafny.CodePoint('d'.codePointAt(0));
        let _rhs137 = (_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f11)).Merge((_1091_v58).Merge(_1091_v58));
        let _rhs138 = _1092_v59;
        _1091_v58 = _rhs137;
        _1092_v59 = _rhs138;
        let _1093_v60;
        _1093_v60 = _dafny.Set.fromElements(_995_v0);
        let _1094_v61;
        let _nw158 = new _module.C2();
        _nw158.__ctor((_1090_v57).f26, p0, ((_1090_v57).f26).multipliedBy(p2), new BigNumber(((_dafny.Set.fromElements(_995_v0)).Intersect(_1093_v60)).length), _this.f11);
        _1094_v61 = _nw158;
        _1091_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-739),_this.f11);
      }
      r1 = p3;
      let _hi2 = (_this).f25;
      for (let _1095_i3 = (_module.__default.fm2(globalState)).minus(p3); _1095_i3.isLessThan(_hi2); _1095_i3 = _1095_i3.plus(_dafny.ONE)) {
        if (true) {
          let _1096_v62;
          let _nw159 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1096_v62 = _nw159;
          _1096_v62 = _1096_v62;
          (globalState).f8 = _dafny.areEqual(((_this.f11) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), function (_1097_i4) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          })) : (p0)), ((_this.f11) ? (p0) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_1098_i5) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }))));
          let _1099_v63;
          _1099_v63 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
          let _1100_v64;
          _1100_v64 = _dafny.Set.fromElements(new BigNumber((p0).length));
          (globalState).f8 = ((((_1099_v63).contains(p1)) ? ((_1099_v63).get(p1)) : (new BigNumber((_1100_v64).length)))).isLessThan(new BigNumber(-92));
          let _1101_v65;
          _1101_v65 = new _dafny.CodePoint('a'.codePointAt(0));
          _1101_v65 = ((_this.f11) ? (_1101_v65) : (_1101_v65));
          let _1102_v66;
          _1102_v66 = _dafny.Seq.UnicodeFromString("pq");
          _1102_v66 = _1102_v66;
        } else {
          let _1103_v67;
          let _init36 = ((_1104_i3) => function (_1105_i6) {
            return (_1105_i6).plus(_1104_i3);
          })(_1095_i3);
          let _nw160 = Array((new BigNumber(9)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw160.length); _i0_36++) {
            _nw160[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1103_v67 = _nw160;
          let _index116 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length));
          (_1103_v67)[_index116] = p2;
          let _index117 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length));
          (_1103_v67)[_index117] = p3;
          let _index118 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length));
          (p1)[_index118] = _this.f11;
          let _index119 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length));
          (p1)[_index119] = (((_1103_v67)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length))]).multipliedBy(p2)).isEqualTo(((_dafny.ZERO).minus(_1095_i3)).minus(_1095_i3));
          let _index120 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length));
          (p1)[_index120] = ((p1)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length))]) && (_this.f11);
          let _1106_v68;
          let _nw161 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1106_v68 = _nw161;
          let _1107_v69;
          _1107_v69 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index121 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1106_v68).length));
          (_1106_v68)[_index121] = _1107_v69;
          let _1108_v70;
          let _nw162 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _1108_v70 = _nw162;
          let _1109_v71;
          _1109_v71 = _dafny.Seq.of(_1095_i3);
          let _1110_v72;
          _1110_v72 = _dafny.MultiSet.fromElements(new BigNumber(531), p3);
          let _1111_v73;
          let _nw163 = Array((new BigNumber(6)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = _1109_v71;
          _nw163[(_dafny.ONE).toNumber()] = _1109_v71;
          _nw163[(new BigNumber(2)).toNumber()] = _1109_v71;
          _nw163[(new BigNumber(3)).toNumber()] = _1109_v71;
          _nw163[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(p2, _1095_i3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(91)), function (_1112_i7) {
            return new BigNumber(-349);
          })).length), (_1109_v71)[_module.__default.safeIndex(new BigNumber((_1110_v72).cardinality()), new BigNumber((_1109_v71).length))]);
          _nw163[(new BigNumber(5)).toNumber()] = _1109_v71;
          _1111_v73 = _nw163;
          let _1113_v74;
          _1113_v74 = _module.D0.create_DC0(_1111_v73);
          let _1114_v75;
          _1114_v75 = _dafny.Set.fromElements(_1113_v74);
          let _index122 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length));
          let _index123 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1106_v68).length));
          let _index124 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length));
          let _rhs139 = (_1114_v75).IsSubsetOf(_1114_v75);
          let _rhs140 = _1107_v69;
          let _rhs141 = _1108_v70;
          let _rhs142 = (_1103_v67)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length))];
          let _rhs143 = !_dafny.areEqual(p0, p0);
          let _lhs78 = p1;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((p1).length));
          let _lhs80 = _1106_v68;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1106_v68).length));
          let _lhs82 = _1103_v67;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_1103_v67).length));
          let _lhs84 = _this;
          _lhs78[_lhs79] = _rhs139;
          _lhs80[_lhs81] = _rhs140;
          _1108_v70 = _rhs141;
          _lhs82[_lhs83] = _rhs142;
          _lhs84.f11 = _rhs143;
          let _1115_v76;
          _1115_v76 = _module.D0.create_DC2(_1095_i3, _1103_v67);
          _1115_v76 = _1115_v76;
        }
        let _1116_v77;
        _1116_v77 = _dafny.Set.fromElements(_this.f11);
        let _1117_v78;
        let _nw164 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1117_v78 = _nw164;
        let _index125 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1117_v78).length));
        (_1117_v78)[_index125] = (_dafny.ZERO).minus(p2);
        let _1118_v79;
        _1118_v79 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(p0, p0)).length));
        let _index126 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1117_v78).length));
        let _rhs144 = ((_1116_v77).Union(_1116_v77)).Intersect(_1116_v77);
        let _rhs145 = (_1118_v79)[_module.__default.safeIndex(p3, new BigNumber((_1118_v79).length))];
        let _rhs146 = _this.f11;
        let _rhs147 = (_this).f25;
        let _lhs85 = _1117_v78;
        let _lhs86 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_1117_v78).length));
        let _lhs87 = globalState;
        _1116_v77 = _rhs144;
        _lhs85[_lhs86] = _rhs145;
        _lhs87.f8 = _rhs146;
        r1 = _rhs147;
        let _index127 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
        (p1)[_index127] = _this.f11;
        let _index128 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
        (p1)[_index128] = _this.f11;
        let _index129 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
        (p1)[_index129] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
      }
      let _1119_v80;
      _1119_v80 = _dafny.MultiSet.fromElements(p3, (_this).f25, (_this).f25);
      let _1120_v81;
      _1120_v81 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1121_v82;
      _1121_v82 = _dafny.Set.fromElements(_this.f11);
      let _1122_v83;
      _1122_v83 = _dafny.Seq.of((_this).f25);
      let _1123_v84;
      _1123_v84 = _module.D5.create_DC15(new BigNumber((_1122_v83).length), p2, p3);
      let _1124_v85;
      _1124_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1121_v82).length),_1123_v84);
      let _1125_v86;
      _1125_v86 = _module.D6.create_DC20(_module.D6.create_DC16(_1124_v85));
      let _1126_v87;
      _1126_v87 = _module.D6.create_DC20(_1125_v86);
      let _1127_v88;
      _1127_v88 = _module.D6.create_DC20(_1125_v86);
      let _1128_v89;
      _1128_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),p3);
      let _pat_let_tv29 = _1119_v80;
      let _pat_let_tv30 = _1119_v80;
      let _pat_let_tv31 = _1119_v80;
      let _pat_let_tv32 = _1119_v80;
      let _pat_let_tv33 = _1122_v83;
      let _pat_let_tv34 = _1122_v83;
      let _pat_let_tv35 = _995_v0;
      let _rhs148 = p3;
      let _rhs149 = function (_source17) {
        if (_source17.is_DC17) {
          let _1129___mcc_h15 = (_source17).cf30;
          let _1130___mcc_h16 = (_source17).cf31;
          let _1131_cf31 = _1130___mcc_h16;
          let _1132_cf30 = _1129___mcc_h15;
          return _pat_let_tv29;
        } else if (_source17.is_DC18) {
          let _1133___mcc_h17 = (_source17).cf32;
          let _1134_cf32 = _1133___mcc_h17;
          return (_pat_let_tv30).Intersect(_dafny.MultiSet.fromElements(new BigNumber(699)));
        } else if (_source17.is_DC19) {
          let _1135___mcc_h18 = (_source17).cf33;
          let _1136___mcc_h19 = (_source17).cf34;
          let _1137___mcc_h20 = (_source17).cf35;
          let _1138___mcc_h21 = (_source17).cf36;
          let _1139___mcc_h22 = (_source17).cf37;
          let _1140_cf37 = _1139___mcc_h22;
          let _1141_cf36 = _1138___mcc_h21;
          let _1142_cf35 = _1137___mcc_h20;
          let _1143_cf34 = _1136___mcc_h19;
          let _1144_cf33 = _1135___mcc_h18;
          return _pat_let_tv31;
        } else if (_source17.is_DC16) {
          let _1145___mcc_h23 = (_source17).cf29;
          let _1146_cf29 = _1145___mcc_h23;
          return _pat_let_tv32;
        } else {
          let _1147___mcc_h24 = (_source17).cf38;
          let _1148_cf38 = _1147___mcc_h24;
          return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_pat_let_tv33, _pat_let_tv34));
        }
      }(((_this.f11) ? (_1127_v88) : (function (_pat_let28_0) {
        return function (_1149_dt__update__tmp_h2) {
          return function (_pat_let29_0) {
            return function (_1150_dt__update_hcf38_h0) {
              return _module.D6.create_DC20(_1150_dt__update_hcf38_h0);
            }(_pat_let29_0);
          }(_module.D6.create_DC19(_this.f11, _this.f11, _this.f11, _this.f11, _pat_let_tv35));
        }(_pat_let28_0);
      }(_1127_v88))));
      let _rhs150 = new BigNumber((_1128_v89).length);
      let _rhs151 = _1120_v81;
      r1 = _rhs148;
      _1119_v80 = _rhs149;
      r0 = _rhs150;
      _1120_v81 = _rhs151;
      r0 = p3;
      r1 = p3;
      let _1151_v90;
      _1151_v90 = _dafny.Seq.of(p1, p1);
      r2 = (_1151_v90)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_1119_v80).cardinality()), new BigNumber((_995_v0).length)), new BigNumber((_1151_v90).length))];
      return [r0, r1, r2];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f24 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    fm26(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (function () {
        let _coll60 = new _dafny.Set();
        for (const _compr_60 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), function (_1152_i0) {
          return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), true))).cardinality()))).length);
        })).Elements) {
          let _1153_v0 = _compr_60;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), function (_1152_i0) {
            return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), true))).cardinality()))).length);
          }), _1153_v0)) {
            _coll60.add(_module.__default.safeDivisionInt(_1153_v0, new BigNumber(11)));
          }
        }
        return _coll60;
      }()).Union((function () {
        let _coll61 = new _dafny.Set();
        for (const _compr_61 of _dafny.IntegerRange(new BigNumber(-629), new BigNumber(423))) {
          let _1154_v1 = _compr_61;
          if (((new BigNumber(-629)).isLessThanOrEqualTo(_1154_v1)) && ((_1154_v1).isLessThan(new BigNumber(423)))) {
            _coll61.add((_1154_v1).multipliedBy((_this).f24));
          }
        }
        return _coll61;
      }()).Intersect(_dafny.Set.fromElements((_this).f24)));
    };
    fm27(p0, p1, p2, globalState) {
      let _this = this;
      return (((_this).f24).minus((_module.D0.create_DC1((_this).f24, _dafny.Seq.UnicodeFromString("hgoabqmce"), new BigNumber((_dafny.Seq.UnicodeFromString("ir")).length), true)).dtor_cf1)).isLessThan(_module.__default.safeModuloInt((_this).f24, (_this).f24));
    };
    m20(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _1155_v0;
      _1155_v0 = true;
      let _1156_v1;
      _1156_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1155_v0,(_this).f24);
      let _1157_v2;
      _1157_v2 = _module.D3.create_DC8(_1156_v1);
      let _source18 = _1157_v2;
      if (_source18.is_DC9) {
        let _1158___mcc_h0 = (_source18).cf15;
        let _1159___mcc_h1 = (_source18).cf16;
        let _1160___mcc_h2 = (_source18).cf17;
        let _1161___mcc_h3 = (_source18).cf18;
        let _1162_cf18 = _1161___mcc_h3;
        let _1163_cf17 = _1160___mcc_h2;
        let _1164_cf16 = _1159___mcc_h1;
        let _1165_cf15 = _1158___mcc_h0;
        (globalState).f8 = (_module.__default.fm2(globalState)).isLessThanOrEqualTo(_module.__default.fm2(globalState));
        let _1166_v3;
        _1166_v3 = _dafny.Seq.of(!(_1155_v0), _1155_v0);
        let _1167_v4;
        _1167_v4 = _dafny.Seq.of(_1166_v3);
        let _1168_v5;
        let _nw165 = new _module.C0();
        _nw165.__ctor(_1165_cf15, _1155_v0, new BigNumber((_dafny.MultiSet.FromArray(_1167_v4)).cardinality()), _1162_cf18);
        _1168_v5 = _nw165;
        let _1169_v6;
        let _nw166 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1169_v6 = _nw166;
        let _1170_v8;
        _1170_v8 = _dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll62 = new _dafny.Map();
          for (const _compr_62 of _dafny.IntegerRange(new BigNumber(326), new BigNumber(901))) {
            let _1171_v7 = _compr_62;
            if (((new BigNumber(326)).isLessThanOrEqualTo(_1171_v7)) && ((_1171_v7).isLessThan(new BigNumber(901)))) {
              _coll62.push([_module.__default.safeModuloInt(_1171_v7, (_1168_v5).f28),_1168_v5.f29]);
            }
          }
          return _coll62;
        }()).length));
        let _index130 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1169_v6).length));
        (_1169_v6)[_index130] = _1170_v8;
        let _1172_v9;
        let _nw167 = Array((new BigNumber(12)).toNumber()).fill([]);
        _1172_v9 = _nw167;
        let _1173_v10;
        _1173_v10 = _module.D7.create_DC23(_1172_v9, _1166_v3, false, new BigNumber(-32));
        let _1174_v11;
        _1174_v11 = _dafny.Seq.of((_this).f24);
        let _index131 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1169_v6).length));
        let _rhs152 = (_1170_v8).Difference((_dafny.MultiSet.FromArray(_1174_v11)).update(new BigNumber((_1166_v3).length), _module.__default.abs(_1165_cf15)));
        let _rhs153 = _1162_cf18;
        let _rhs154 = _1173_v10;
        let _rhs155 = (_1168_v5.f29) && (_1155_v0);
        let _lhs88 = _1169_v6;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1169_v6).length));
        let _lhs90 = globalState;
        _lhs88[_lhs89] = _rhs152;
        r0 = _rhs153;
        _1173_v10 = _rhs154;
        _lhs90.f8 = _rhs155;
        _1173_v10 = _1173_v10;
      } else if (_source18.is_DC8) {
        let _1175___mcc_h4 = (_source18).cf14;
        let _1176_cf14 = _1175___mcc_h4;
        if (_1155_v0) {
          let _1177_v12;
          let _nw168 = Array((new BigNumber(7)).toNumber());
          _nw168[(_dafny.ZERO).toNumber()] = !(!((_this).f24).isEqualTo((_this).f24));
          _nw168[(_dafny.ONE).toNumber()] = _1155_v0;
          _nw168[(new BigNumber(2)).toNumber()] = _1155_v0;
          _nw168[(new BigNumber(3)).toNumber()] = _1155_v0;
          _nw168[(new BigNumber(4)).toNumber()] = _1155_v0;
          _nw168[(new BigNumber(5)).toNumber()] = !(((_this).f24).isLessThanOrEqualTo((_this).f24));
          _nw168[(new BigNumber(6)).toNumber()] = _1155_v0;
          _1177_v12 = _nw168;
          let _1178_v13;
          let _nw169 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1178_v13 = _nw169;
          let _index132 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1178_v13).length));
          (_1178_v13)[_index132] = (_this).f24;
          let _1179_v14;
          _1179_v14 = _dafny.Seq.UnicodeFromString("ihtnsyxyb");
          let _index133 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1178_v13).length));
          let _rhs156 = _1177_v12;
          let _rhs157 = _1155_v0;
          let _rhs158 = (_this).f24;
          let _rhs159 = ((_this).f24).minus(new BigNumber((_1179_v14).length));
          let _lhs91 = globalState;
          let _lhs92 = _1178_v13;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1178_v13).length));
          _1177_v12 = _rhs156;
          _lhs91.f8 = _rhs157;
          _lhs92[_lhs93] = _rhs158;
          r0 = _rhs159;
          let _1180_v15;
          _1180_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,_1155_v0);
          _1180_v15 = (_1180_v15).update(!(_1155_v0) || (_1155_v0), _1155_v0);
          let _1181_v16;
          _1181_v16 = _dafny.Seq.of(!(_1155_v0));
          let _1182_v17;
          _1182_v17 = _module.D6.create_DC19(_1155_v0, _1155_v0, _1155_v0, _1155_v0, _1181_v16);
          let _1183_v18;
          _1183_v18 = _module.D6.create_DC20(_1182_v17);
          let _rhs160 = _module.__default.safeDivisionInt((_this).f24, (_1178_v13)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_1178_v13).length))]);
          let _rhs161 = _1183_v18;
          let _rhs162 = new BigNumber((_1180_v15).length);
          r0 = _rhs160;
          _1183_v18 = _rhs161;
          r0 = _rhs162;
          let _1184_v19;
          _1184_v19 = _dafny.MultiSet.fromElements(_1155_v0);
          let _1185_v20;
          _1185_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1184_v19,_1155_v0);
          let _1186_v21;
          _1186_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1185_v20,_1155_v0);
          let _1187_v23;
          _1187_v23 = _dafny.Seq.of(_1184_v19, _1184_v19);
          _1186_v21 = (_1186_v21).update(function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of (_1187_v23).Elements) {
              let _1188_v22 = _compr_63;
              if (_dafny.Seq.contains(_1187_v23, _1188_v22)) {
                _coll63.push([_1188_v22,_1155_v0]);
              }
            }
            return _coll63;
          }(), _1155_v0);
          let _1189_v24;
          let _nw170 = Array((new BigNumber(6)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _1178_v13;
          _nw170[(_dafny.ONE).toNumber()] = _1178_v13;
          _nw170[(new BigNumber(2)).toNumber()] = _1178_v13;
          _nw170[(new BigNumber(3)).toNumber()] = _1178_v13;
          _nw170[(new BigNumber(4)).toNumber()] = _1178_v13;
          _nw170[(new BigNumber(5)).toNumber()] = _1178_v13;
          _1189_v24 = _nw170;
          let _1190_v25;
          _1190_v25 = _module.D7.create_DC24(_module.D7.create_DC22((_1178_v13)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_1178_v13).length))], _1189_v24));
          let _1191_v26;
          _1191_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1189_v24);
          let _1192_v27;
          _1192_v27 = _dafny.Seq.of((_this).f24);
          let _pat_let_tv36 = _1191_v26;
          let _pat_let_tv37 = _1192_v27;
          let _pat_let_tv38 = _1178_v13;
          let _pat_let_tv39 = _1178_v13;
          let _pat_let_tv40 = _1192_v27;
          let _pat_let_tv41 = _1155_v0;
          let _pat_let_tv42 = _1189_v24;
          let _pat_let_tv43 = _1191_v26;
          let _pat_let_tv44 = _1192_v27;
          let _pat_let_tv45 = _1178_v13;
          let _pat_let_tv46 = _1178_v13;
          let _pat_let_tv47 = _1192_v27;
          let _pat_let_tv48 = _1155_v0;
          let _pat_let_tv49 = _1181_v16;
          let _pat_let_tv50 = _1155_v0;
          let _pat_let_tv51 = _1179_v14;
          _1190_v25 = function (_pat_let30_0) {
            return function (_1193_dt__update__tmp_h0) {
              return function (_pat_let31_0) {
                return function (_1194_dt__update_hcf46_h0) {
                  return _module.D7.create_DC24(_1194_dt__update_hcf46_h0);
                }(_pat_let31_0);
              }(_module.D7.create_DC23((((_pat_let_tv43).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv44)[_module.__default.safeIndex((_pat_let_tv46)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_pat_let_tv45).length))], new BigNumber((_pat_let_tv47).length))],!(_pat_let_tv48))).length))) ? ((_pat_let_tv36).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv37)[_module.__default.safeIndex((_pat_let_tv39)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_pat_let_tv38).length))], new BigNumber((_pat_let_tv40).length))],!(_pat_let_tv41))).length))) : (_pat_let_tv42)), _pat_let_tv49, _pat_let_tv50, new BigNumber((_pat_let_tv51).length)));
            }(_pat_let30_0);
          }(_module.D7.create_DC24(_module.D7.create_DC21(_1184_v19)));
        } else {
          let _1195_v28;
          _1195_v28 = _dafny.Set.fromElements(_1155_v0, true, _1155_v0);
          _1195_v28 = _1195_v28;
          let _1196_v29;
          let _nw171 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _1196_v29 = _nw171;
          let _1197_v30;
          _1197_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1196_v29,(_1155_v0) && (_1155_v0));
          _1197_v30 = (_1197_v30).update(_1196_v29, _1155_v0);
          let _1198_v31;
          _1198_v31 = _dafny.Seq.of(_1155_v0);
          let _1199_v32;
          _1199_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(new BigNumber((_dafny.MultiSet.FromArray(_1198_v31)).cardinality()), _1155_v0),_dafny.Set.fromElements((_this).f24));
          let _1200_v33;
          _1200_v33 = _module.D0.create_DC3((_this).f24, _1155_v0);
          let _1201_v34;
          _1201_v34 = _dafny.Set.fromElements((_this).f24);
          let _1202_v35;
          _1202_v35 = _dafny.Set.fromElements(new BigNumber((_1201_v34).length));
          _1199_v32 = (_1199_v32).Merge((_1199_v32).update(_1200_v33, _1202_v35));
          let _1203_v36;
          _1203_v36 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
          let _1204_v37;
          _1204_v37 = _dafny.Seq.of(new BigNumber(258));
          let _1205_v38;
          _1205_v38 = _dafny.Seq.of((_1204_v37)[_module.__default.safeIndex((((_1203_v36).contains(new BigNumber(473))) ? ((_1203_v36).get(new BigNumber(473))) : ((_this).f24)), new BigNumber((_1204_v37).length))]);
          let _1206_v39;
          _1206_v39 = _module.D1.create_DC5(((_1203_v36).update((_this).f24, _module.__default.abs((_this).f24))).Difference(_dafny.MultiSet.FromArray(_1204_v37)), _1155_v0);
          let _pat_let_tv52 = _1155_v0;
          let _pat_let_tv53 = _1155_v0;
          _1206_v39 = function (_pat_let32_0) {
            return function (_1207_dt__update__tmp_h1) {
              return function (_pat_let33_0) {
                return function (_1208_dt__update_hcf11_h0) {
                  return _module.D1.create_DC5((_1207_dt__update__tmp_h1).dtor_cf10, _1208_dt__update_hcf11_h0);
                }(_pat_let33_0);
              }((_pat_let_tv52) || (_pat_let_tv53));
            }(_pat_let32_0);
          }(_1206_v39);
          (globalState).f8 = (_1203_v36).IsSubsetOf(_1203_v36);
        }
        let _1209_v40;
        _1209_v40 = _dafny.Seq.UnicodeFromString("dvfyoivjb");
        (globalState).f8 = !(_dafny.Seq.IsProperPrefixOf(_1209_v40, _dafny.Seq.UnicodeFromString("llh"))) || (_1155_v0);
        (globalState).f8 = ((_this).f24).isEqualTo(new BigNumber(-218));
        let _1210_v41;
        let _nw172 = new _module.C2();
        _nw172.__ctor((_this).f24, _1209_v40, (_this).f24, (_this).f24, _1155_v0);
        _1210_v41 = _nw172;
      } else {
        let _1211___mcc_h5 = (_source18).cf19;
        let _1212_cf19 = _1211___mcc_h5;
        r0 = (_this).f24;
        let _1213_v42;
        let _init37 = function (_1214_i0) {
          return false;
        };
        let _nw173 = Array((new BigNumber(24)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw173.length); _i0_37++) {
          _nw173[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1213_v42 = _nw173;
        let _index134 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1213_v42).length));
        (_1213_v42)[_index134] = _1155_v0;
        let _index135 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1213_v42).length));
        (_1213_v42)[_index135] = _1155_v0;
        let _1215_v43;
        let _init38 = function (_1216_i1) {
          return _module.D5.create_DC15((_this).f24, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_1217_i2) {
  return new _dafny.CodePoint('e'.codePointAt(0));
}), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_1218_i2) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length)), new _dafny.CodePoint('b'.codePointAt(0)))).length), (_this).f24);
        };
        let _nw174 = Array((_dafny.ONE).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw174.length); _i0_38++) {
          _nw174[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1215_v43 = _nw174;
        let _1219_v44;
        _1219_v44 = _dafny.MultiSet.fromElements(_1215_v43);
        let _1220_v45;
        _1220_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1219_v44).Union(_1219_v44));
        let _1221_v46;
        _1221_v46 = _dafny.Set.fromElements((_this).f24);
        let _1222_v47;
        _1222_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1221_v46).length));
        _1220_v45 = (_1220_v45).update((((_1222_v47).contains(new BigNumber(-360))) ? ((_1222_v47).get(new BigNumber(-360))) : ((_this).f24)), _1219_v44);
        let _1223_v48;
        _1223_v48 = _module.D5.create_DC15(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber((_dafny.Seq.of((_1213_v42)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1213_v42).length))], _1155_v0, _1155_v0, _1155_v0, (_1213_v42)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1213_v42).length))])).length), (_1213_v42)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1213_v42).length))], (_this).f24, globalState),new BigNumber((_dafny.Seq.of((_this).f24)).length))).length), (_this).f24, (_this).f24);
        let _1224_v49;
        _1224_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1223_v48);
        let _1225_v50;
        _1225_v50 = _module.D6.create_DC16(_1224_v49);
        let _1226_v51;
        _1226_v51 = _module.D6.create_DC20(_1225_v50);
        let _1227_v52;
        _1227_v52 = _dafny.Seq.of(_module.D6.create_DC20(_1225_v50), _1226_v51, _1226_v51);
        r0 = new BigNumber((_1227_v52).length);
      }
      let _1228_v53;
      _1228_v53 = _dafny.Seq.UnicodeFromString("atdgxf");
      let _1229_v54;
      _1229_v54 = _module.D0.create_DC1((_this).f24, _1228_v53, (_this).f24, !(_1155_v0));
      if ((_1229_v54).dtor_cf4) {
        let _1230_v55;
        _1230_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1155_v0,_1155_v0);
        r0 = new BigNumber(((((_1230_v55).update(_1155_v0, _1155_v0)).Merge(_1230_v55)).Merge(_1230_v55)).length);
        let _1231_v56;
        _1231_v56 = new _dafny.CodePoint('r'.codePointAt(0));
        r1 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1228_v53, _module.__default.safeIndex((_this).f24, new BigNumber((_1228_v53).length)), _1231_v56), _module.__default.safeIndex((_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.fm2(globalState))).minus((_this).f24)), new BigNumber((_dafny.Seq.update(_1228_v53, _module.__default.safeIndex((_this).f24, new BigNumber((_1228_v53).length)), _1231_v56)).length)), _1231_v56)).length);
        let _1232_v57;
        _1232_v57 = _module.D9.create_DC26(_1231_v56);
        (globalState).f8 = !(_dafny.Seq.contains(_1228_v53, (_1232_v57).dtor_cf48));
        let _1233_v58;
        _1233_v58 = _dafny.MultiSet.fromElements(_1231_v56, _1231_v56);
        _1233_v58 = (_1233_v58).Intersect(_1233_v58);
        _1155_v0 = true;
      } else {
        let _1234_v59;
        _1234_v59 = _dafny.Seq.of((_this).f24);
        let _1235_v60;
        _1235_v60 = _module.D3.create_DC8(_1156_v1);
        let _1236_v61;
        _1236_v61 = _dafny.Seq.of(_module.D3.create_DC8(_1156_v1), _1235_v60);
        let _1237_v62;
        _1237_v62 = _module.D3.create_DC10((_1236_v61)[_module.__default.safeIndex(new BigNumber((_1234_v59).length), new BigNumber((_1236_v61).length))]);
        let _1238_v63;
        _1238_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v62,_1156_v1);
        let _1239_v64;
        let _nw175 = Array((new BigNumber(20)).toNumber());
        _nw175[(_dafny.ZERO).toNumber()] = _1156_v1;
        _nw175[(_dafny.ONE).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(2)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(3)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(4)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(5)).toNumber()] = (_1156_v1).Merge(_1156_v1);
        _nw175[(new BigNumber(6)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(7)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(8)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(9)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(10)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1155_v0,(_this).f24);
        _nw175[(new BigNumber(12)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1155_v0,new BigNumber((_1234_v59).length));
        _nw175[(new BigNumber(14)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(15)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(16)).toNumber()] = (((_1238_v63).contains(_1237_v62)) ? ((_1238_v63).get(_1237_v62)) : (_1156_v1));
        _nw175[(new BigNumber(17)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(18)).toNumber()] = _1156_v1;
        _nw175[(new BigNumber(19)).toNumber()] = (_1156_v1).Merge(_1156_v1);
        _1239_v64 = _nw175;
        let _index136 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1239_v64).length));
        (_1239_v64)[_index136] = _1156_v1;
        let _index137 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1239_v64).length));
        (_1239_v64)[_index137] = _1156_v1;
        _1155_v0 = _1155_v0;
        let _1240_v65;
        let _nw176 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1240_v65 = _nw176;
        let _index138 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1240_v65).length));
        (_1240_v65)[_index138] = _1155_v0;
        let _index139 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1240_v65).length));
        (_1240_v65)[_index139] = _1155_v0;
        (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f24, (_this).f24, _module.__default.fm2(globalState), (_this).f24, (_this).f24), _1234_v59);
        let _index140 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1240_v65).length));
        (_1240_v65)[_index140] = (_1155_v0) && (((_1155_v0) ? (_1155_v0) : ((_1240_v65)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1240_v65).length))])));
      }
      let _1241_v66;
      _1241_v66 = _dafny.Seq.of(_1155_v0);
      let _1242_v67;
      _1242_v67 = _dafny.Set.fromElements(_1155_v0, _1155_v0);
      let _1243_v68;
      _1243_v68 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f24, new BigNumber((_1241_v66).length), new BigNumber((_1242_v67).length)),_dafny.Seq.Concat(_1228_v53, _dafny.Seq.UnicodeFromString("omygfqai")));
      let _1244_v69;
      _1244_v69 = _dafny.Set.fromElements(new BigNumber(-491), (_this).f24, (_this).f24);
      let _1245_v70;
      _1245_v70 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, (_this).f24, (_this).f24);
      let _1246_v71;
      _1246_v71 = new _dafny.CodePoint('j'.codePointAt(0));
      _1243_v68 = (_1243_v68).update(_1244_v69, _dafny.Seq.update(_1228_v53, _module.__default.safeIndex((((_1245_v70).contains(new BigNumber(483))) ? ((_1245_v70).get(new BigNumber(483))) : ((_this).f24)), new BigNumber((_1228_v53).length)), _1246_v71));
      let _hi3 = ((_this).f24).minus(new BigNumber(179));
      for (let _1247_i3 = (_this).f24; _1247_i3.isLessThan(_hi3); _1247_i3 = _1247_i3.plus(_dafny.ONE)) {
        let _1248_v72;
        let _nw177 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1248_v72 = _nw177;
        _1248_v72 = _1248_v72;
        r1 = _module.__default.safeModuloInt((_this).f24, (_this).f24);
        let _1249_v73;
        let _init39 = function (_1250_i4) {
          return false;
        };
        let _nw178 = Array((new BigNumber(6)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw178.length); _i0_39++) {
          _nw178[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1249_v73 = _nw178;
        let _1251_v74;
        let _nw179 = Array((new BigNumber(13)).toNumber());
        _nw179[(_dafny.ZERO).toNumber()] = _1249_v73;
        _nw179[(_dafny.ONE).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(2)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(3)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(4)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(5)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(6)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(7)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(8)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(9)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(10)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(11)).toNumber()] = _1249_v73;
        _nw179[(new BigNumber(12)).toNumber()] = _1249_v73;
        _1251_v74 = _nw179;
        let _1252_v75;
        _1252_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1244_v69,_1251_v74);
        _1252_v75 = (_1252_v75).update(_dafny.Set.fromElements((_this).f24, (_this).f24, _1247_i3), _1251_v74);
        let _1253_v76;
        _1253_v76 = _dafny.Set.fromElements(_module.__default.fm47(_1155_v0, _1155_v0, _1156_v1, _1247_i3, globalState));
        let _1254_v77;
        let _nw180 = Array((new BigNumber(17)).toNumber()).fill([]);
        _1254_v77 = _nw180;
        let _pat_let_tv54 = _1254_v77;
        let _1255_v78;
        _1255_v78 = _module.D9.create_DC27(_1155_v0, function (_pat_let34_0) {
  return function (_1256_dt__update__tmp_h2) {
    return function (_pat_let35_0) {
      return function (_1257_dt__update_hcf41_h0) {
        return _module.D7.create_DC22((_1256_dt__update__tmp_h2).dtor_cf40, _1257_dt__update_hcf41_h0);
      }(_pat_let35_0);
    }(_pat_let_tv54);
  }(_pat_let34_0);
}(_module.D7.create_DC22(new BigNumber((_dafny.Seq.of((_this).f24, new BigNumber((_1253_v76).length), (_this).f24)).length), _1254_v77)), (_this).f24, _1155_v0);
        (globalState).f8 = (_this).fm27((_1255_v78).dtor_cf51, _1228_v53, (new BigNumber(-919)).isEqualTo(new BigNumber((_1241_v66).length)), globalState);
      }
      let _1258_v79;
      _1258_v79 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_1155_v0);
      let _1259_v80;
      _1259_v80 = _module.D6.create_DC18(_1155_v0);
      let _1260_v81;
      _1260_v81 = _module.D6.create_DC19(_1155_v0, (((_1258_v79).contains(_1155_v0)) ? ((_1258_v79).get(_1155_v0)) : ((_1259_v80).dtor_cf32)), !(_module.__default.fm1((_this).f24, _1155_v0, (_this).f24, globalState)), _1155_v0, _dafny.Seq.of(_1155_v0, _1155_v0, _1155_v0, _1155_v0, _1155_v0));
      let _rhs163 = (((_this).fm27((_this).f24, _dafny.Seq.UnicodeFromString("ldu"), _1155_v0, globalState)) ? (_1155_v0) : (_1155_v0));
      let _rhs164 = _module.D6.create_DC19(_1155_v0, !(true), _1155_v0, !(!((_1155_v0) === (_1155_v0))), _1241_v66);
      let _rhs165 = _1228_v53;
      let _rhs166 = ((_this).f24).plus((_this).f24);
      let _rhs167 = (new BigNumber(355)).isLessThanOrEqualTo((_this).f24);
      _1155_v0 = _rhs163;
      _1260_v81 = _rhs164;
      _1228_v53 = _rhs165;
      r1 = _rhs166;
      _1155_v0 = _rhs167;
      let _1261_v83;
      _1261_v83 = _module.D3.create_DC9((_this).f24, function () {
  let _coll64 = new _dafny.Map();
  for (const _compr_64 of _dafny.IntegerRange(new BigNumber(-620), new BigNumber(554))) {
    let _1262_v82 = _compr_64;
    if (((new BigNumber(-620)).isLessThanOrEqualTo(_1262_v82)) && ((_1262_v82).isLessThan(new BigNumber(554)))) {
      _coll64.push([(_1262_v82).multipliedBy((_this).f24),_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)]);
    }
  }
  return _coll64;
}(), new BigNumber((_1241_v66).length), _module.__default.fm2(globalState));
      let _1263_v84;
      _1263_v84 = _module.D3.create_DC10(_1261_v83);
      let _1264_v85;
      _1264_v85 = _module.D3.create_DC10(_1261_v83);
      let _pat_let_tv55 = _1263_v84;
      let _source19 = function (_pat_let36_0) {
        return function (_1265_dt__update__tmp_h3) {
          return function (_pat_let37_0) {
            return function (_1266_dt__update_hcf19_h0) {
              return _module.D3.create_DC10(_1266_dt__update_hcf19_h0);
            }(_pat_let37_0);
          }(_pat_let_tv55);
        }(_pat_let36_0);
      }(_1264_v85);
      if (_source19.is_DC9) {
        let _1267___mcc_h6 = (_source19).cf15;
        let _1268___mcc_h7 = (_source19).cf16;
        let _1269___mcc_h8 = (_source19).cf17;
        let _1270___mcc_h9 = (_source19).cf18;
        let _1271_cf18 = _1270___mcc_h9;
        let _1272_cf17 = _1269___mcc_h8;
        let _1273_cf16 = _1268___mcc_h7;
        let _1274_cf15 = _1267___mcc_h6;
        let _1275_v86;
        _1275_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1274_cf15,_1155_v0);
        let _1276_v87;
        _1276_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1274_cf15,(_this).f24);
        let _1277_v88;
        let _nw181 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _1277_v88 = _nw181;
        let _1278_v89;
        let _nw182 = Array((new BigNumber(25)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = _1277_v88;
        _nw182[(_dafny.ONE).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(2)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(3)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(4)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(5)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(6)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(7)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(8)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(9)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(10)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(11)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(12)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(13)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(14)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(15)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(16)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(17)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(18)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(19)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(20)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(21)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(22)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(23)).toNumber()] = _1277_v88;
        _nw182[(new BigNumber(24)).toNumber()] = _1277_v88;
        _1278_v89 = _nw182;
        let _1279_v90;
        _1279_v90 = _module.D7.create_DC22(_1272_cf17, _1278_v89);
        let _1280_v91;
        _1280_v91 = _module.D9.create_DC27(_1155_v0, _1279_v90, (_this).f24, true);
        let _1281_v92;
        let _nw183 = Array((new BigNumber(25)).toNumber());
        _nw183[(_dafny.ZERO).toNumber()] = _1155_v0;
        _nw183[(_dafny.ONE).toNumber()] = false;
        _nw183[(new BigNumber(2)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(3)).toNumber()] = false;
        _nw183[(new BigNumber(4)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(5)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(6)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(7)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(8)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(9)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(10)).toNumber()] = !(_1155_v0) || (_1155_v0);
        _nw183[(new BigNumber(11)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(12)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(13)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(14)).toNumber()] = (_1272_cf17).isEqualTo(_1271_cf18);
        _nw183[(new BigNumber(15)).toNumber()] = (_1229_v54).dtor_cf4;
        _nw183[(new BigNumber(16)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(17)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(18)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(19)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(20)).toNumber()] = (((_1275_v86).contains(new BigNumber((_1276_v87).length))) ? ((_1275_v86).get(new BigNumber((_1276_v87).length))) : (true));
        _nw183[(new BigNumber(21)).toNumber()] = (_this).fm27(_1274_cf15, (_module.D0.create_DC1((((_1245_v70).contains((_this).f24)) ? ((_1245_v70).get((_this).f24)) : ((_this).f24)), _1228_v53, _1272_cf17, !(_1155_v0))).dtor_cf2, (_1280_v91).dtor_cf52, globalState);
        _nw183[(new BigNumber(22)).toNumber()] = ((_1260_v81).dtor_cf36) === (_1155_v0);
        _nw183[(new BigNumber(23)).toNumber()] = _1155_v0;
        _nw183[(new BigNumber(24)).toNumber()] = _1155_v0;
        _1281_v92 = _nw183;
        let _index141 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1281_v92).length));
        (_1281_v92)[_index141] = true;
        let _index142 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1281_v92).length));
        (_1281_v92)[_index142] = !_dafny.Seq.contains(_1241_v66, _1155_v0);
        let _1282_v93;
        let _nw184 = new _module.C2();
        _nw184.__ctor(_1272_cf17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(369)), ((_1283_v71) => function (_1284_i5) {
          return _1283_v71;
        })(_1246_v71)), (_this).f24, new BigNumber(943), _1155_v0);
        _1282_v93 = _nw184;
        let _1285_v94;
        _1285_v94 = _dafny.Seq.of(_module.__default.fm2(globalState));
        _1274_cf15 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_1271_cf18, (_1285_v94)[_module.__default.safeIndex((_this).f24, new BigNumber((_1285_v94).length))])).multipliedBy(_module.__default.fm2(globalState)));
        r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1272_cf17), (_this).f24);
      } else if (_source19.is_DC8) {
        let _1286___mcc_h10 = (_source19).cf14;
        let _1287_cf14 = _1286___mcc_h10;
        let _1288_v95;
        let _nw185 = Array((new BigNumber(26)).toNumber());
        _nw185[(_dafny.ZERO).toNumber()] = _1155_v0;
        _nw185[(_dafny.ONE).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(2)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(3)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(4)).toNumber()] = true;
        _nw185[(new BigNumber(5)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(6)).toNumber()] = true;
        _nw185[(new BigNumber(7)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(8)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(9)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(10)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(11)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(12)).toNumber()] = true;
        _nw185[(new BigNumber(13)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(14)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(15)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(16)).toNumber()] = !(_1155_v0);
        _nw185[(new BigNumber(17)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(18)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(19)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(20)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(21)).toNumber()] = true;
        _nw185[(new BigNumber(22)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(23)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(24)).toNumber()] = _1155_v0;
        _nw185[(new BigNumber(25)).toNumber()] = _1155_v0;
        _1288_v95 = _nw185;
        (globalState).f8 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_1288_v95, _1288_v95), p0);
        if ((!(_1155_v0)) && (_1155_v0)) {
          let _1289_v96;
          let _1290_v97;
          let _out12;
          let _out13;
          let _outcollector6 = _module.__default.m0(globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _1289_v96 = _out12;
          _1290_v97 = _out13;
          let _index143 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_1288_v95).length));
          (_1288_v95)[_index143] = _1290_v97;
          let _index144 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_1288_v95).length));
          (_1288_v95)[_index144] = _module.__default.fm1(new BigNumber(315), !(_1290_v97) || (_1155_v0), ((_this).f24).multipliedBy(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1290_v97,(_this).f24)).update(_1290_v97, (_this).f24)).length)), globalState);
          let _1291_v98;
          let _nw186 = new _module.C1();
          _nw186.__ctor(new BigNumber(58));
          _1291_v98 = _nw186;
          r1 = (_this).f24;
          let _1292_v99;
          let _nw187 = new _module.C1();
          _nw187.__ctor(new BigNumber((_1228_v53).length));
          _1292_v99 = _nw187;
        } else {
          r0 = _module.__default.fm2(globalState);
          r0 = (_dafny.ZERO).minus((_this).f24);
          let _1293_v100;
          _1293_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1258_v79,_module.__default.safeModuloInt((_this).f24, (_this).f24));
          let _1294_v101;
          _1294_v101 = _dafny.Seq.of((_this).f24, (_this).f24);
          let _1295_v102;
          _1295_v102 = _dafny.MultiSet.fromElements(_1294_v101, _dafny.Seq.Create(_module.__default.abs(new BigNumber(315)), function (_1296_i6) {
            return (_this).f24;
          }));
          let _1297_v103;
          _1297_v103 = _module.D0.create_DC3((((_1295_v102).contains(_dafny.Seq.of(new BigNumber(262)))) ? ((_1295_v102).get(_dafny.Seq.of(new BigNumber(262)))) : ((_this).f24)), _1155_v0);
          let _1298_v104;
          _1298_v104 = _dafny.MultiSet.fromElements(_1297_v103, ((_1155_v0) ? (_1297_v103) : (_1297_v103)), _1297_v103);
          let _rhs168 = (_1293_v100).Merge(_1293_v100);
          let _rhs169 = (_1298_v104).Union(_module.__default.fm48(new BigNumber((_dafny.Seq.of(_1155_v0, _1155_v0)).length), _1246_v71, new BigNumber((_1241_v66).length), _1155_v0, globalState));
          _1293_v100 = _rhs168;
          _1298_v104 = _rhs169;
          (globalState).f8 = _1155_v0;
          r0 = (_this).f24;
        }
        _1155_v0 = ((_1155_v0) ? (_1155_v0) : (_1155_v0));
        let _1299_v105;
        _1299_v105 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
        r1 = ((_dafny.Seq.contains(_1228_v53, _1246_v71)) ? (((_1155_v0) ? ((((_1299_v105).contains((_this).f24)) ? ((_1299_v105).get((_this).f24)) : ((_this).f24))) : (new BigNumber(485)))) : (((_this).f24).plus((_dafny.ZERO).minus((_this).f24))));
      } else {
        let _1300___mcc_h11 = (_source19).cf19;
        let _1301_cf19 = _1300___mcc_h11;
        _1246_v71 = _1246_v71;
        let _1302_v106;
        _1302_v106 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-671)),new BigNumber(201));
        _1156_v1 = (_dafny.Map.Empty.slice().updateUnsafe(_1155_v0,new BigNumber((_1302_v106).length))).Merge((_1156_v1).update(_1155_v0, (_this).f24));
        let _1303_v107;
        _1303_v107 = _module.D9.create_DC26((_1228_v53)[_module.__default.safeIndex(new BigNumber((_1244_v69).length), new BigNumber((_1228_v53).length))]);
        let _source20 = _1303_v107;
        if (_source20.is_DC27) {
          let _1304___mcc_h12 = (_source20).cf49;
          let _1305___mcc_h13 = (_source20).cf50;
          let _1306___mcc_h14 = (_source20).cf51;
          let _1307___mcc_h15 = (_source20).cf52;
          let _1308_cf52 = _1307___mcc_h15;
          let _1309_cf51 = _1306___mcc_h14;
          let _1310_cf50 = _1305___mcc_h13;
          let _1311_cf49 = _1304___mcc_h12;
          r0 = (_dafny.ZERO).minus(((_1309_cf51).multipliedBy(_1309_cf51)).multipliedBy(_module.__default.fm2(globalState)));
          _1156_v1 = (_1156_v1).update(true, (_this).f24);
          let _1312_v108;
          let _nw188 = new _module.C4();
          _nw188.__ctor(_module.__default.safeModuloInt(_module.__default.fm2(globalState), new BigNumber((_dafny.Seq.of(true)).length)), _1155_v0);
          _1312_v108 = _nw188;
          let _1313_v109;
          let _nw189 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1313_v109 = _nw189;
          let _1314_v110;
          let _nw190 = Array((new BigNumber(25)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = _1313_v109;
          _nw190[(_dafny.ONE).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(2)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(3)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(4)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(5)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(6)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(7)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(8)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(9)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(10)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(11)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(12)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(13)).toNumber()] = (_module.D0.create_DC2(_1309_cf51, _1313_v109)).dtor_cf6;
          _nw190[(new BigNumber(14)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(15)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(16)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(17)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(18)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(19)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(20)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(21)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(22)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(23)).toNumber()] = _1313_v109;
          _nw190[(new BigNumber(24)).toNumber()] = _1313_v109;
          _1314_v110 = _nw190;
          let _1315_v111;
          _1315_v111 = _module.D7.create_DC23(_1314_v110, _1241_v66, !(_1308_cf52), (_this).f24);
          let _index145 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1313_v109).length));
          (_1313_v109)[_index145] = (_1315_v111).dtor_cf45;
          let _index146 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1313_v109).length));
          (_1313_v109)[_index146] = (_1312_v108).f25;
        } else if (_source20.is_DC26) {
          let _1316___mcc_h16 = (_source20).cf48;
          let _1317_cf48 = _1316___mcc_h16;
          let _1318_v112;
          let _init40 = function (_1319_i7) {
            return (_1319_i7).multipliedBy((_this).f24);
          };
          let _nw191 = Array((new BigNumber(25)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw191.length); _i0_40++) {
            _nw191[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1318_v112 = _nw191;
          let _1320_v113;
          _1320_v113 = _dafny.Seq.of(_1318_v112, _1318_v112);
          let _1321_v114;
          let _nw192 = new _module.C3();
          _nw192.__ctor(new BigNumber((_1320_v113).length), (_this).f24);
          _1321_v114 = _nw192;
          let _1322_v115;
          _1322_v115 = _module.D0.create_DC3((_this).f24, _1155_v0);
          let _1323_v116;
          let _init41 = ((_1324_v0) => function (_1325_i8) {
            return !(_1324_v0);
          })(_1155_v0);
          let _nw193 = Array((new BigNumber(27)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw193.length); _i0_41++) {
            _nw193[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1323_v116 = _nw193;
          let _1326_v117;
          _1326_v117 = _dafny.Seq.of(_1323_v116);
          let _rhs170 = _1322_v115;
          let _rhs171 = _1326_v117;
          _1322_v115 = _rhs170;
          _1326_v117 = _rhs171;
          let _1327_v119;
          _1327_v119 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)).length));
          let _1328_v120;
          _1328_v120 = _dafny.Seq.of(new BigNumber((function () {
            let _coll65 = new _dafny.Map();
            for (const _compr_65 of (_1327_v119).Elements) {
              let _1329_v118 = _compr_65;
              if (_dafny.Seq.contains(_1327_v119, _1329_v118)) {
                _coll65.push([(_1329_v118).plus((_this).f24),_1155_v0]);
              }
            }
            return _coll65;
          }()).length), (_dafny.ZERO).minus((_this).f24), (_this).f24, (_this).f24, (_this).f24);
          (globalState).f5 = _dafny.Seq.Concat(_1327_v119, _1328_v120);
          let _1330_v121;
          let _nw194 = Array((new BigNumber(13)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = _1323_v116;
          _nw194[(_dafny.ONE).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(2)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(3)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(4)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(5)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(6)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(7)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(8)).toNumber()] = ((!(_1155_v0)) ? (_1323_v116) : (_1323_v116));
          _nw194[(new BigNumber(9)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(10)).toNumber()] = (_1326_v117)[_module.__default.safeIndex(new BigNumber(375), new BigNumber((_1326_v117).length))];
          _nw194[(new BigNumber(11)).toNumber()] = _1323_v116;
          _nw194[(new BigNumber(12)).toNumber()] = _1323_v116;
          _1330_v121 = _nw194;
          _1330_v121 = ((_1155_v0) ? (_1330_v121) : (_1330_v121));
        } else {
          let _1331___mcc_h17 = (_source20).cf53;
          let _1332_cf53 = _1331___mcc_h17;
          let _1333_v122;
          let _nw195 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1333_v122 = _nw195;
          let _index147 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1333_v122).length));
          (_1333_v122)[_index147] = _1228_v53;
          let _index148 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1333_v122).length));
          (_1333_v122)[_index148] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-765)), ((_1334_v71) => function (_1335_i9) {
            return _1334_v71;
          })(_1246_v71));
          _1246_v71 = ((_1333_v122)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1333_v122).length))])[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rdjws"), _module.__default.safeIndex(new BigNumber(-193), new BigNumber((_dafny.Seq.UnicodeFromString("rdjws")).length)), _1246_v71)).length), (_this).f24), new BigNumber(((_1333_v122)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1333_v122).length))]).length))];
          let _1336_v123;
          let _nw196 = Array((new BigNumber(4)).toNumber()).fill(false);
          _1336_v123 = _nw196;
          let _1337_v124;
          _1337_v124 = _dafny.Seq.of(_1336_v123);
          _1337_v124 = _dafny.Seq.Concat(_1337_v124, _1337_v124);
          let _1338_v125;
          let _1339_v126;
          let _out14;
          let _out15;
          let _outcollector7 = _module.__default.m0(globalState);
          _out14 = _outcollector7[0];
          _out15 = _outcollector7[1];
          _1338_v125 = _out14;
          _1339_v126 = _out15;
        }
        let _1340_v127;
        _1340_v127 = _module.D5.create_DC15((_this).f24, (_this).f24, (_this).f24);
        let _1341_v128;
        _1341_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1340_v127);
        let _1342_v129;
        _1342_v129 = _module.D6.create_DC16(_1341_v128);
        let _1343_v130;
        _1343_v130 = _module.D6.create_DC20(_module.D6.create_DC20(_1342_v129));
        let _source21 = _1343_v130;
        if (_source21.is_DC17) {
          let _1344___mcc_h18 = (_source21).cf30;
          let _1345___mcc_h19 = (_source21).cf31;
          let _1346_cf31 = _1345___mcc_h19;
          let _1347_cf30 = _1344___mcc_h18;
          let _1348_v131;
          let _nw197 = Array((new BigNumber(8)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = _1347_cf30;
          _nw197[(_dafny.ONE).toNumber()] = (_1244_v69).IsDisjointFrom(_1244_v69);
          _nw197[(new BigNumber(2)).toNumber()] = _1155_v0;
          _nw197[(new BigNumber(3)).toNumber()] = _1155_v0;
          _nw197[(new BigNumber(4)).toNumber()] = _1155_v0;
          _nw197[(new BigNumber(5)).toNumber()] = _1155_v0;
          _nw197[(new BigNumber(6)).toNumber()] = _1155_v0;
          _nw197[(new BigNumber(7)).toNumber()] = (_1347_cf30) === (_1347_cf30);
          _1348_v131 = _nw197;
          _1348_v131 = _1348_v131;
          let _1349_v132;
          _1349_v132 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1346_cf31,_1155_v0),_module.__default.fm49(globalState));
          let _1350_v133;
          _1350_v133 = _dafny.Set.fromElements(_1244_v69);
          _1349_v132 = (_1349_v132).update(_module.__default.fm47(!((_this).fm27(_1346_cf31, _module.__default.fm36(_1346_cf31, _1346_cf31, _1246_v71, globalState), false, globalState)), _1347_cf30, _1156_v1, new BigNumber((_1228_v53).length), globalState), _1350_v133);
          let _1351_v134;
          let _init42 = function (_1352_i10) {
            return _module.__default.safeDivisionInt(_1352_i10, new BigNumber(744));
          };
          let _nw198 = Array((new BigNumber(18)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw198.length); _i0_42++) {
            _nw198[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1351_v134 = _nw198;
          let _index149 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1351_v134).length));
          (_1351_v134)[_index149] = ((_dafny.ZERO).minus(_1346_cf31)).plus(_1346_cf31);
          let _index150 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1351_v134).length));
          (_1351_v134)[_index150] = (_this).f24;
          r2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1228_v53, _1228_v53), _dafny.Seq.Concat(_1228_v53, _1228_v53));
        } else if (_source21.is_DC18) {
          let _1353___mcc_h20 = (_source21).cf32;
          let _1354_cf32 = _1353___mcc_h20;
          let _1355_v135;
          let _1356_v136;
          let _out16;
          let _out17;
          let _outcollector8 = _module.__default.m0(globalState);
          _out16 = _outcollector8[0];
          _out17 = _outcollector8[1];
          _1355_v135 = _out16;
          _1356_v136 = _out17;
          (globalState).f8 = _1354_cf32;
          let _1357_v137;
          let _nw199 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1357_v137 = _nw199;
          let _index151 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1357_v137).length));
          (_1357_v137)[_index151] = ((_this).f24).multipliedBy((_this).f24);
          let _index152 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1357_v137).length));
          (_1357_v137)[_index152] = (_this).f24;
          (globalState).f8 = ((true) ? (_1155_v0) : (_1354_cf32));
        } else if (_source21.is_DC19) {
          let _1358___mcc_h21 = (_source21).cf33;
          let _1359___mcc_h22 = (_source21).cf34;
          let _1360___mcc_h23 = (_source21).cf35;
          let _1361___mcc_h24 = (_source21).cf36;
          let _1362___mcc_h25 = (_source21).cf37;
          let _1363_cf37 = _1362___mcc_h25;
          let _1364_cf36 = _1361___mcc_h24;
          let _1365_cf35 = _1360___mcc_h23;
          let _1366_cf34 = _1359___mcc_h22;
          let _1367_cf33 = _1358___mcc_h21;
          let _1368_v138;
          _1368_v138 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1228_v53);
          _1228_v53 = (((_1368_v138).contains((_this).f24)) ? ((_1368_v138).get((_this).f24)) : (_1228_v53));
          let _1369_v139;
          _1369_v139 = _dafny.Seq.of((_this).f24);
          let _1370_v140;
          let _nw200 = Array((new BigNumber(4)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1369_v139;
          _nw200[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_module.__default.fm2(globalState), (_this).f24, (_this).f24, new BigNumber((_1245_v70).cardinality()), (_this).f24);
          _nw200[(new BigNumber(2)).toNumber()] = _1369_v139;
          _nw200[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), ((_1371_v139) => function (_1372_i11) {
            return (_1371_v139)[_module.__default.safeIndex((_this).f24, new BigNumber((_1371_v139).length))];
          })(_1369_v139));
          _1370_v140 = _nw200;
          let _1373_v141;
          _1373_v141 = _module.D0.create_DC0(_1370_v140);
          let _pat_let_tv56 = _1370_v140;
          _1373_v141 = function (_pat_let38_0) {
            return function (_1374_dt__update__tmp_h4) {
              return function (_pat_let39_0) {
                return function (_1375_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_1375_dt__update_hcf0_h0);
                }(_pat_let39_0);
              }(_pat_let_tv56);
            }(_pat_let38_0);
          }(_1373_v141);
          r0 = (_this).f24;
          let _1376_v142;
          _1376_v142 = _module.D4.create_DC13(new BigNumber((_dafny.MultiSet.FromArray(_1363_cf37)).cardinality()));
          let _1377_v143;
          _1377_v143 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_1376_v142);
          let _1378_v144;
          _1378_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1365_cf35,_1377_v143);
          _1378_v144 = _1378_v144;
        } else if (_source21.is_DC16) {
          let _1379___mcc_h26 = (_source21).cf29;
          let _1380_cf29 = _1379___mcc_h26;
          let _1381_v145;
          _1381_v145 = _dafny.Seq.of(new BigNumber((_1241_v66).length), (_this).f24, (_this).f24, _module.__default.fm2(globalState), _module.__default.fm2(globalState));
          (globalState).f8 = _dafny.Seq.contains(_1381_v145, ((_this).f24).minus((_this).f24));
          let _1382_v146;
          let _init43 = function (_1383_i12) {
            return ((_this).f24).isLessThan((_this).f24);
          };
          let _nw201 = Array((new BigNumber(6)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw201.length); _i0_43++) {
            _nw201[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1382_v146 = _nw201;
          let _1384_v147;
          _1384_v147 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1228_v53).length),_1155_v0);
          let _1385_v148;
          _1385_v148 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1381_v145);
          let _index153 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1382_v146).length));
          (_1382_v146)[_index153] = (((_1384_v147).contains((_dafny.ZERO).minus((_this).f24))) ? ((_1384_v147).get((_dafny.ZERO).minus((_this).f24))) : (_module.__default.fm1(new BigNumber(671), _1155_v0, new BigNumber((_1385_v148).length), globalState)));
          let _index154 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1382_v146).length));
          (_1382_v146)[_index154] = false;
          let _1386_v149;
          let _nw202 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1386_v149 = _nw202;
          _1386_v149 = _1386_v149;
          let _1387_v150;
          let _nw203 = Array((new BigNumber(3)).toNumber());
          _nw203[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1155_v0), _1241_v66);
          _nw203[(_dafny.ONE).toNumber()] = _1241_v66;
          _nw203[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1241_v66, _1241_v66);
          _1387_v150 = _nw203;
          let _index155 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1387_v150).length));
          (_1387_v150)[_index155] = _dafny.Seq.Concat(_1241_v66, _module.__default.fm39((_1382_v146)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1382_v146).length))], (_1382_v146)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1382_v146).length))], globalState));
          let _1388_v151;
          _1388_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1246_v71,_1241_v66);
          let _index156 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1387_v150).length));
          (_1387_v150)[_index156] = _dafny.Seq.Concat(_1241_v66, (((_1388_v151).contains(_1246_v71)) ? ((_1388_v151).get(_1246_v71)) : (_module.__default.fm39((_1382_v146)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1382_v146).length))], _1155_v0, globalState))));
        } else {
          let _1389___mcc_h27 = (_source21).cf38;
          let _1390_cf38 = _1389___mcc_h27;
          let _1391_v152;
          let _nw204 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1391_v152 = _nw204;
          let _1392_v153;
          _1392_v153 = _dafny.Map.Empty.slice().updateUnsafe(_1155_v0,_1246_v71);
          let _index157 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1391_v152).length));
          (_1391_v152)[_index157] = !(_1392_v153).contains(_1155_v0);
          let _index158 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1391_v152).length));
          (_1391_v152)[_index158] = _1155_v0;
          (globalState).f8 = !(_1155_v0);
          let _1393_v154;
          _1393_v154 = _module.D1.create_DC5(_dafny.MultiSet.fromElements((_this).f24, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("okfmdc"), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.UnicodeFromString("okfmdc")).length)), _1246_v71)).length)), _1155_v0);
          _1245_v70 = (_1245_v70).Difference((_1393_v154).dtor_cf10);
          r0 = (new BigNumber((_1228_v53).length)).plus((_this).f24);
        }
      }
      r0 = (_this).f24;
      r1 = (_this).f24;
      let _1394_v155;
      _1394_v155 = _dafny.MultiSet.fromElements(_1258_v79, _1258_v79, (_dafny.Map.Empty.slice().updateUnsafe(_1155_v0,(_1241_v66)[_module.__default.safeIndex((_this).f24, new BigNumber((_1241_v66).length))])).update(_1155_v0, _1155_v0));
      let _1395_v156;
      _1395_v156 = _1394_v155;
      let _pat_let_tv57 = _1228_v53;
      let _pat_let_tv58 = _1228_v53;
      r2 = _dafny.Seq.update(function (_source22) {
        let _1396___mcc_h28 = _source22;
        let _1397_cf47 = _1396___mcc_h28;
        return _dafny.Seq.Concat(_pat_let_tv57, _dafny.Seq.UnicodeFromString("sywamaofb"));
      }(_1395_v156), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f24), new BigNumber((function (_source23) {
        let _1398___mcc_h29 = _source23;
        let _1399_cf47 = _1398___mcc_h29;
        return _dafny.Seq.Concat(_pat_let_tv58, _dafny.Seq.UnicodeFromString("sywamaofb"));
      }(_1395_v156)).length)), ((_1155_v0) ? ((_1228_v53)[_module.__default.safeIndex((_this).f24, new BigNumber((_1228_v53).length))]) : (_1246_v71)));
      return [r0, r1, r2];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm24(p0, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("gfqdouu"))).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xjnse"), _dafny.Seq.UnicodeFromString("cfgxspivk")))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber(-283));
      } else {
        return (function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(339))).cardinality()), new BigNumber(153), new BigNumber(441), (_dafny.ZERO).minus(new BigNumber(-446)))).length))).Keys.Elements) {
            let _1400_v0 = _compr_66;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(339))).cardinality()), new BigNumber(153), new BigNumber(441), (_dafny.ZERO).minus(new BigNumber(-446)))).length))).contains(_1400_v0)) {
              _coll66.push([_1400_v0,new BigNumber(431)]);
            }
          }
          return _coll66;
        }()).Merge(function () {
          let _coll67 = new _dafny.Map();
          for (const _compr_67 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true)).Keys.Elements) {
            let _1401_v1 = _compr_67;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true)).contains(_1401_v1)) {
              _coll67.push([_1401_v1,new BigNumber(-180)]);
            }
          }
          return _coll67;
        }());
      }
    };
    m18(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.MultiSet.Empty;
      let r3 = _module.D0.Default();
      let _1402_v0;
      _1402_v0 = false;
      let _1403_v1;
      _1403_v1 = new BigNumber(602);
      let _1404_v2;
      _1404_v2 = _dafny.Seq.UnicodeFromString("hybdjuc");
      let _1405_v3;
      _1405_v3 = _module.D0.create_DC1(_1403_v1, _1404_v2, _1403_v1, _module.__default.fm1(_1403_v1, _1402_v0, _1403_v1, globalState));
      let _1406_v4;
      _1406_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1402_v0,_1405_v3);
      let _1407_v5;
      _1407_v5 = _dafny.Seq.of(_1402_v0);
      _1406_v4 = (_1406_v4).update((_1407_v5)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_1407_v5).length))], _1405_v3);
      _1402_v0 = _1402_v0;
      let _1408_i0;
      _1408_i0 = _dafny.ZERO;
      L4: {
        while (_1402_v0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1408_i0)) {
              break L4;
            }
            _1408_i0 = (_1408_i0).plus(_dafny.ONE);
            r1 = _1403_v1;
            let _1409_v6;
            let _nw205 = Array((new BigNumber(2)).toNumber()).fill([]);
            _1409_v6 = _nw205;
            let _nw206 = Array((new BigNumber(4)).toNumber()).fill([]);
            _1409_v6 = _nw206;
            r0 = !(new BigNumber(399)).isEqualTo(_module.__default.safeModuloInt(_1403_v1, _1403_v1));
            r0 = _1402_v0;
          }
        }
      }
      if (_1402_v0) {
        _1403_v1 = new BigNumber(183);
        let _1410_v7;
        let _nw207 = Array((new BigNumber(12)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = !(_1403_v1).isEqualTo(_1403_v1);
        _nw207[(_dafny.ONE).toNumber()] = _module.__default.fm1(_1403_v1, _1402_v0, (_dafny.ZERO).minus(_1403_v1), globalState);
        _nw207[(new BigNumber(2)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(3)).toNumber()] = (_1407_v5)[_module.__default.safeIndex(_1403_v1, new BigNumber((_1407_v5).length))];
        _nw207[(new BigNumber(4)).toNumber()] = ((_1402_v0) ? (_1402_v0) : (!(_1402_v0)));
        _nw207[(new BigNumber(5)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(6)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(7)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(8)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(9)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(10)).toNumber()] = _1402_v0;
        _nw207[(new BigNumber(11)).toNumber()] = _1402_v0;
        _1410_v7 = _nw207;
        let _index159 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1410_v7).length));
        (_1410_v7)[_index159] = _1402_v0;
        let _index160 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1410_v7).length));
        (_1410_v7)[_index160] = _module.__default.fm1(_1403_v1, _1402_v0, _1403_v1, globalState);
        (globalState).f8 = _module.__default.fm1(new BigNumber(842), (_1410_v7)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_1410_v7).length))], _1403_v1, globalState);
        let _1411_v8;
        _1411_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1403_v1).multipliedBy(_1403_v1),(_1410_v7)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_1410_v7).length))]);
        _1411_v8 = (_1411_v8).update(_1403_v1, (_1410_v7)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_1410_v7).length))]);
        let _1412_v9;
        _1412_v9 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_module.__default.fm2(globalState));
        let _1413_v10;
        _1413_v10 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1414_v11;
        _1414_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_1404_v2, _module.__default.safeIndex(new BigNumber((_1412_v9).length), new BigNumber((_1404_v2).length)), _1413_v10), _1404_v2),_1403_v1);
        _1414_v11 = (_1414_v11).update(_1404_v2, _1403_v1);
      } else {
        (globalState).f8 = _1402_v0;
        let _1415_v12;
        _1415_v12 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1416_v13;
        _1416_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1403_v1,_1415_v12);
        let _1417_v14;
        _1417_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1403_v1,_1402_v0);
        let _1418_v15;
        _1418_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1417_v14).length),_1402_v0);
        let _1419_v16;
        _1419_v16 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1416_v13).length),_1402_v0), _1417_v14);
        let _1420_v17;
        _1420_v17 = _dafny.Set.fromElements(_1402_v0, _1402_v0);
        let _1421_v18;
        _1421_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(new BigNumber((_1419_v16).cardinality()), new BigNumber((_1420_v17).length), globalState),false);
        _1402_v0 = !(_1421_v18).equals((_1421_v18).Merge(_1421_v18));
        let _1422_v19;
        let _init44 = ((_1423_v1) => function (_1424_i1) {
          return (_1424_i1).minus(_1423_v1);
        })(_1403_v1);
        let _nw208 = Array((new BigNumber(18)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw208.length); _i0_44++) {
          _nw208[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1422_v19 = _nw208;
        let _index161 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length));
        (_1422_v19)[_index161] = (new BigNumber((_1404_v2).length)).minus(_1403_v1);
        let _index162 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length));
        (_1422_v19)[_index162] = _1403_v1;
        let _1425_v20;
        let _init45 = ((_1426_v2) => function (_1427_i2) {
          return _1426_v2;
        })(_1404_v2);
        let _nw209 = Array((new BigNumber(12)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw209.length); _i0_45++) {
          _nw209[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1425_v20 = _nw209;
        let _1428_v21;
        _1428_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1422_v19)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length))],_1425_v20);
        _1428_v21 = (_1428_v21).update((_dafny.ZERO).minus((_1422_v19)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length))]), _1425_v20);
        let _1429_v22;
        _1429_v22 = _dafny.Seq.of(new BigNumber(952));
        let _1430_v23;
        _1430_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1402_v0,_1429_v22);
        (globalState).f5 = _dafny.Seq.Concat((((_1430_v23).contains(_1402_v0)) ? ((_1430_v23).get(_1402_v0)) : (_dafny.Seq.of(_1403_v1, (_1422_v19)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length))], (_1422_v19)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length))], (_1422_v19)[_module.__default.safeIndex(new BigNumber(972), new BigNumber((_1422_v19).length))], _1403_v1))), _1429_v22);
      }
      if (_1402_v0) {
        let _1431_v24;
        _1431_v24 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1432_v25;
        _1432_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v24,_1403_v1);
        let _1433_v26;
        _1433_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1403_v1,_1402_v0);
        let _1434_v27;
        _1434_v27 = _dafny.Seq.of(_1433_v26, (_1433_v26).update(_1403_v1, !(_1402_v0)), _1433_v26);
        let _1435_v28;
        let _nw210 = new _module.C0();
        _nw210.__ctor(_1403_v1, _1402_v0, (_dafny.ZERO).minus(new BigNumber((_1432_v25).length)), new BigNumber((_dafny.MultiSet.FromArray(_1434_v27)).cardinality()));
        _1435_v28 = _nw210;
        let _1436_v29;
        _1436_v29 = _dafny.MultiSet.fromElements(_1402_v0, _1402_v0);
        let _1437_v30;
        _1437_v30 = _dafny.Seq.of(_1436_v29, _1436_v29);
        let _1438_v31;
        _1438_v31 = _dafny.Map.Empty.slice().updateUnsafe(((_1435_v28.f29) ? (_1402_v0) : (!(true))),(_1437_v30)[_module.__default.safeIndex((_1435_v28).f28, new BigNumber((_1437_v30).length))]);
        _1438_v31 = ((_1402_v0) ? (_1438_v31) : (_1438_v31));
        (globalState).f8 = _1435_v28.f29;
        let _1439_v32;
        _1439_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1435_v28).f28,_module.__default.safeModuloInt(new BigNumber(183), new BigNumber(-824)));
        _1439_v32 = (_1439_v32).update((_1435_v28).f28, new BigNumber(682));
        r1 = _module.__default.safeDivisionInt(_1403_v1, _1403_v1);
      } else {
        let _1440_v33;
        _1440_v33 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1441_v34;
        _1441_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1440_v33,_1403_v1);
        let _1442_v35;
        _1442_v35 = _dafny.Set.fromElements(new BigNumber((_1404_v2).length), (((_1441_v34).contains(new _dafny.CodePoint('t'.codePointAt(0)))) ? ((_1441_v34).get(new _dafny.CodePoint('t'.codePointAt(0)))) : (new BigNumber(257))));
        let _1443_v36;
        let _nw211 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1443_v36 = _nw211;
        let _1444_v37;
        _1444_v37 = _dafny.MultiSet.fromElements(_1402_v0, _1402_v0, _1402_v0);
        let _1445_v38;
        let _out18;
        _out18 = (_this).m19(_1442_v35, _1443_v36, _1403_v1, _1444_v37, globalState);
        _1445_v38 = _out18;
        _1403_v1 = (_1403_v1).multipliedBy(((_1402_v0) ? (_1403_v1) : (new BigNumber(286))));
        let _index163 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1445_v38).length));
        (_1445_v38)[_index163] = _1402_v0;
        let _1446_v39;
        _1446_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v38,_1402_v0);
        let _index164 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1445_v38).length));
        let _rhs172 = (_1402_v0) === ((_1403_v1).isLessThanOrEqualTo(_1403_v1));
        let _rhs173 = _1402_v0;
        let _rhs174 = _1402_v0;
        let _rhs175 = _dafny.Seq.of(!((((_1446_v39).contains(_1445_v38)) ? ((_1446_v39).get(_1445_v38)) : (_1402_v0))), _1402_v0);
        let _rhs176 = ((_1403_v1).minus(_1403_v1)).isLessThan(_1403_v1);
        let _lhs94 = globalState;
        let _lhs95 = _1445_v38;
        let _lhs96 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1445_v38).length));
        _lhs94.f8 = _rhs172;
        _1402_v0 = _rhs173;
        r0 = _rhs174;
        _1407_v5 = _rhs175;
        _lhs95[_lhs96] = _rhs176;
        _1440_v33 = _1440_v33;
        let _1447_v40;
        let _init46 = ((_1448_v2) => function (_1449_i3) {
          return _1448_v2;
        })(_1404_v2);
        let _nw212 = Array((new BigNumber(12)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw212.length); _i0_46++) {
          _nw212[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1447_v40 = _nw212;
        let _1450_v41;
        _1450_v41 = _module.D5.create_DC14(_1447_v40);
        let _1451_v42;
        _1451_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1450_v41,(_1445_v38)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_1445_v38).length))]);
        let _1452_v43;
        _1452_v43 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((_1451_v42).update(_1450_v41, (_1445_v38)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_1445_v38).length))])).length)), _1403_v1, _1403_v1);
        let _1453_v44;
        _1453_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1440_v33,((_1452_v43)[_module.__default.safeIndex(_1403_v1, new BigNumber((_1452_v43).length))]).isLessThanOrEqualTo(_1403_v1));
        let _1454_v45;
        _1454_v45 = _dafny.MultiSet.fromElements(_1403_v1, _1403_v1, _1403_v1);
        _1453_v44 = (_1453_v44).update(_1440_v33, ((((_1454_v45).contains(_1403_v1)) ? ((_1454_v45).get(_1403_v1)) : (_1403_v1))).isLessThan(_1403_v1));
      }
      r0 = (_1403_v1).isLessThan(_1403_v1);
      r0 = _1402_v0;
      let _1455_v46;
      _1455_v46 = _dafny.Map.Empty.slice().updateUnsafe(true,_1403_v1);
      let _1456_v47;
      _1456_v47 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm50(globalState),(new BigNumber((_dafny.Seq.of(new BigNumber((_1455_v46).length))).length)).plus(_1403_v1));
      let _1457_v48;
      _1457_v48 = _dafny.Set.fromElements(_1403_v1);
      r1 = (((_1456_v47).contains(_1457_v48)) ? ((_1456_v47).get(_1457_v48)) : (_1403_v1));
      let _1458_v49;
      let _nw213 = Array((new BigNumber(4)).toNumber()).fill([]);
      _1458_v49 = _nw213;
      let _1459_v50;
      _1459_v50 = _module.D7.create_DC22(_1403_v1, _1458_v49);
      let _1460_v51;
      _1460_v51 = _module.D9.create_DC27(_1402_v0, _1459_v50, _1403_v1, (_1407_v5)[_module.__default.safeIndex(_1403_v1, new BigNumber((_1407_v5).length))]);
      let _1461_v52;
      let _nw214 = Array((new BigNumber(4)).toNumber());
      _nw214[(_dafny.ZERO).toNumber()] = _1402_v0;
      _nw214[(_dafny.ONE).toNumber()] = (_1460_v51).dtor_cf49;
      _nw214[(new BigNumber(2)).toNumber()] = _1402_v0;
      _nw214[(new BigNumber(3)).toNumber()] = !(_1402_v0);
      _1461_v52 = _nw214;
      let _1462_v53;
      _1462_v53 = _dafny.MultiSet.fromElements(_1461_v52);
      r2 = _1462_v53;
      let _1463_v54;
      let _init47 = ((_1464_v1) => function (_1465_i4) {
        return _dafny.Seq.of(_1464_v1);
      })(_1403_v1);
      let _nw215 = Array((new BigNumber(17)).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw215.length); _i0_47++) {
        _nw215[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _1463_v54 = _nw215;
      r3 = _module.D0.create_DC0(_1463_v54);
      return [r0, r1, r2, r3];
    }
    m19(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let _1466_v0;
      let _nw216 = Array((new BigNumber(24)).toNumber()).fill([]);
      _1466_v0 = _nw216;
      _1466_v0 = _1466_v0;
      let _1467_v1;
      _1467_v1 = true;
      let _1468_v2;
      _1468_v2 = _dafny.MultiSet.fromElements(p2);
      let _1469_v3;
      _1469_v3 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1470_v4;
      let _nw217 = Array((new BigNumber(17)).toNumber());
      _nw217[(_dafny.ZERO).toNumber()] = _1467_v1;
      _nw217[(_dafny.ONE).toNumber()] = _module.__default.fm1(p2, _1467_v1, new BigNumber(-777), globalState);
      _nw217[(new BigNumber(2)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(3)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(4)).toNumber()] = (new BigNumber(94)).isLessThan(new BigNumber(181));
      _nw217[(new BigNumber(5)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(6)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(7)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(8)).toNumber()] = !(_1467_v1) || (_1467_v1);
      _nw217[(new BigNumber(9)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(10)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(11)).toNumber()] = _module.__default.fm1(new BigNumber((_module.__default.fm51(new BigNumber((_1468_v2).cardinality()), p2, p2, _1469_v3, globalState)).length), _1467_v1, p2, globalState);
      _nw217[(new BigNumber(12)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(13)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(14)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(15)).toNumber()] = _1467_v1;
      _nw217[(new BigNumber(16)).toNumber()] = (p2).isLessThanOrEqualTo(p2);
      _1470_v4 = _nw217;
      let _index165 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
      (_1470_v4)[_index165] = !(_1467_v1);
      let _index166 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
      (_1470_v4)[_index166] = _1467_v1;
      let _index167 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
      (_1470_v4)[_index167] = true;
      let _1471_v5;
      _1471_v5 = _dafny.Seq.UnicodeFromString("r");
      if (!_dafny.areEqual(_1471_v5, (((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]) ? (_1471_v5) : (_1471_v5)))) {
        let _1472_v6;
        _1472_v6 = _dafny.Set.fromElements(p2, new BigNumber((_dafny.Seq.Concat(_1471_v5, _1471_v5)).length));
        let _1473_v7;
        _1473_v7 = p0;
        _1472_v6 = (_1473_v7);
        let _1474_v8;
        _1474_v8 = _module.D4.create_DC13(new BigNumber(473));
        let _1475_v9;
        _1475_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1474_v8,_dafny.MultiSet.fromElements(_1467_v1));
        let _1476_v10;
        _1476_v10 = _dafny.Map.Empty.slice().updateUnsafe(((_1467_v1) ? (_dafny.Map.Empty.slice().updateUnsafe(_module.D4.create_DC13(new BigNumber(461)),p3)) : (_1475_v9)),_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_1471_v5, _module.__default.safeIndex((((p3).contains((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))])) ? ((p3).get((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))])) : (p2)), new BigNumber((_1471_v5).length)), _1469_v3)).length), p2));
        _1476_v10 = (_1476_v10).update(_1475_v9, (p2).plus(p2));
        let _1477_v11;
        let _nw218 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1477_v11 = _nw218;
        let _1478_v12;
        let _init48 = ((_1479_v1, _1480_v4) => function (_1481_i0) {
          return _dafny.Set.fromElements(_1479_v1, !((_1480_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1480_v4).length))]));
        })(_1467_v1, _1470_v4);
        let _nw219 = Array((new BigNumber(6)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw219.length); _i0_48++) {
          _nw219[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1478_v12 = _nw219;
        let _index168 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_1477_v11).length));
        (_1477_v11)[_index168] = _1478_v12;
        let _index169 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_1477_v11).length));
        (_1477_v11)[_index169] = _1478_v12;
        let _index170 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((p1).length));
        (p1)[_index170] = _module.__default.fm2(globalState);
        let _1482_v13;
        _1482_v13 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _1483_v14;
        _1483_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1467_v1,true);
        let _1484_v15;
        _1484_v15 = _dafny.Seq.of((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], !((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]), (_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], _1467_v1);
        let _1485_v16;
        _1485_v16 = _dafny.Set.fromElements(_1483_v14, (((_1484_v15)[_module.__default.safeIndex(new BigNumber((_1471_v5).length), new BigNumber((_1484_v15).length))]) ? (_1483_v14) : (_1483_v14)), _dafny.Map.Empty.slice().updateUnsafe((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))],_1467_v1));
        let _1486_v17;
        _1486_v17 = _dafny.Seq.of(p2, new BigNumber(889));
        let _1487_v18;
        _1487_v18 = _dafny.Set.fromElements(_1467_v1, _1467_v1);
        let _index171 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((p1).length));
        let _rhs177 = (_1484_v15)[_module.__default.safeIndex(_module.__default.safeModuloInt(p2, new BigNumber((_1486_v17).length)), new BigNumber((_1484_v15).length))];
        let _rhs178 = new BigNumber((_1487_v18).length);
        let _rhs179 = _1482_v13;
        let _rhs180 = _1485_v16;
        let _rhs181 = (_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))];
        let _lhs97 = globalState;
        let _lhs98 = p1;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((p1).length));
        let _lhs100 = globalState;
        _lhs97.f8 = _rhs177;
        _lhs98[_lhs99] = _rhs178;
        _1482_v13 = _rhs179;
        _1485_v16 = _rhs180;
        _lhs100.f8 = _rhs181;
        let _index172 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((p1).length));
        (p1)[_index172] = (_1486_v17)[_module.__default.safeIndex(p2, new BigNumber((_1486_v17).length))];
      } else {
        let _1488_v19;
        _1488_v19 = new BigNumber(699);
        _1488_v19 = new BigNumber(970);
        let _1489_v20;
        _1489_v20 = _dafny.Seq.of(_1467_v1);
        let _1490_v21;
        _1490_v21 = _module.D6.create_DC19(_1467_v1, _1467_v1, (_1489_v20)[_module.__default.safeIndex(p2, new BigNumber((_1489_v20).length))], _1467_v1, _1489_v20);
        _1467_v1 = !((_1490_v21).dtor_cf33);
        let _index173 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
        (_1470_v4)[_index173] = (_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))];
        if (_1467_v1) {
          let _1491_v22;
          _1491_v22 = _dafny.Set.fromElements(_1467_v1, !((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]));
          let _1492_v23;
          _1492_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], _1467_v1)).Difference(_1491_v22),_dafny.Seq.UnicodeFromString("qaswmji"));
          _1492_v23 = (_1492_v23).update((_1491_v22).Difference(_1491_v22), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), ((_1493_v3) => function (_1494_i1) {
            return _1493_v3;
          })(_1469_v3)), _1471_v5));
          let _index174 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((p1).length));
          (p1)[_index174] = _1488_v19;
          let _index175 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((p1).length));
          (p1)[_index175] = p2;
          let _1495_v24;
          let _nw220 = Array((new BigNumber(19)).toNumber()).fill(_module.D6.Default());
          _1495_v24 = _nw220;
          let _1496_v25;
          _1496_v25 = _dafny.MultiSet.fromElements(_1495_v24);
          let _1497_v26;
          _1497_v26 = _module.D11.create_DC30(_1496_v25);
          let _rhs182 = _1469_v3;
          let _rhs183 = (_1496_v25).Union(((_1496_v25).update(_1495_v24, _module.__default.abs(p2))).Intersect((_1497_v26).dtor_cf55));
          _1469_v3 = _rhs182;
          _1496_v25 = _rhs183;
          let _index176 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
          (_1470_v4)[_index176] = (_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))];
          let _1498_v27;
          _1498_v27 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((p1).length))],_1469_v3);
          let _1499_v28;
          _1499_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))],new BigNumber((_dafny.Set.fromElements(new BigNumber((_1498_v27).length), (p1)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((p1).length))])).length));
          let _index177 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length));
          (_1470_v4)[_index177] = (new BigNumber((_1499_v28).length)).isEqualTo(_1488_v19);
        } else {
          let _1500_v29;
          _1500_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1471_v5,(_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]);
          let _1501_v30;
          _1501_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_1500_v29).contains(_1471_v5)) ? ((_1500_v29).get(_1471_v5)) : ((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))])),_1488_v19);
          _1501_v30 = _1501_v30;
          let _1502_v31;
          _1502_v31 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ssylvffrk"));
          let _1503_v32;
          _1503_v32 = _module.D0.create_DC1(p2, _1471_v5, p2, !(_1467_v1));
          _1502_v31 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-907)), ((_1504_v3) => function (_1505_i2) {
            return _1504_v3;
          })(_1469_v3)), _1471_v5, (_1503_v32).dtor_cf2, _1471_v5, _1471_v5);
          _1488_v19 = _1488_v19;
          _1488_v19 = p2;
          let _index178 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((p1).length));
          (p1)[_index178] = _1488_v19;
          let _1506_v33;
          _1506_v33 = _dafny.Seq.of(_1489_v20);
          let _1507_v34;
          _1507_v34 = _dafny.Seq.of((((p3).contains(_1467_v1)) ? ((p3).get(_1467_v1)) : (_1488_v19)));
          let _1508_v35;
          let _nw221 = new _module.C2();
          _nw221.__ctor(p2, _1471_v5, p2, (_1507_v34)[_module.__default.safeIndex(p2, new BigNumber((_1507_v34).length))], _1467_v1);
          _1508_v35 = _nw221;
          let _1509_v36;
          _1509_v36 = _dafny.MultiSet.fromElements(_1508_v35);
          let _index179 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((p1).length));
          let _rhs184 = p2;
          let _rhs185 = (_dafny.ZERO).minus((p2).multipliedBy((new BigNumber((_1471_v5).length)).plus(new BigNumber((_1506_v33).length))));
          let _rhs186 = _module.__default.safeDivisionInt(_1488_v19, _module.__default.safeDivisionInt(p2, (((_1509_v36).contains(_1508_v35)) ? ((_1509_v36).get(_1508_v35)) : (_1488_v19))));
          let _rhs187 = ((p2).multipliedBy(p2)).minus(_1488_v19);
          let _lhs101 = p1;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((p1).length));
          _lhs101[_lhs102] = _rhs184;
          _1488_v19 = _rhs185;
          _1488_v19 = _rhs186;
          _1488_v19 = _rhs187;
        }
        let _index180 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((p1).length));
        (p1)[_index180] = p2;
        let _1510_v37;
        _1510_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))],(_dafny.MultiSet.fromElements(_1467_v1, _1467_v1)).update(true, _module.__default.abs(p2)));
        let _1511_v38;
        _1511_v38 = _module.D12.create_DC32(_1510_v37);
        let _index181 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((p1).length));
        let _rhs188 = (_1488_v19).plus(_1488_v19);
        let _rhs189 = (_1511_v38).dtor_cf57;
        let _lhs103 = p1;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((p1).length));
        _lhs103[_lhs104] = _rhs188;
        _1510_v37 = _rhs189;
      }
      let _1512_v39;
      _1512_v39 = _module.D1.create_DC4(_dafny.Seq.of(p1, p1));
      let _1513_v40;
      _1513_v40 = _module.D1.create_DC6(_1512_v39);
      let _source24 = _1513_v40;
      if (_source24.is_DC5) {
        let _1514___mcc_h0 = (_source24).cf10;
        let _1515___mcc_h1 = (_source24).cf11;
        let _1516_cf11 = _1515___mcc_h1;
        let _1517_cf10 = _1514___mcc_h0;
        (globalState).f8 = _1516_cf11;
        let _1518_v41;
        _1518_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(546),p2);
        let _1519_v43;
        _1519_v43 = _module.D3.create_DC9(p2, function () {
  let _coll68 = new _dafny.Map();
  for (const _compr_68 of (_1468_v2).Elements) {
    let _1520_v42 = _compr_68;
    if ((_1468_v2).contains(_1520_v42)) {
      _coll68.push([(_1520_v42).multipliedBy(p2),_1518_v41]);
    }
  }
  return _coll68;
}(), new BigNumber((_1471_v5).length), p2);
        let _1521_v44;
        _1521_v44 = _dafny.MultiSet.fromElements(_module.D3.create_DC9(p2, _dafny.Map.Empty.slice().updateUnsafe(p2,_1518_v41), p2, p2), _1519_v43, _1519_v43);
        (globalState).f8 = (_module.__default.fm52(globalState)).IsSubsetOf(_1521_v44);
        let _1522_v45;
        _1522_v45 = _dafny.Set.fromElements(_1467_v1);
        (globalState).f8 = !((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]) || ((_1522_v45).IsDisjointFrom(_1522_v45));
        let _1523_v46;
        _1523_v46 = new BigNumber(-948);
        _1523_v46 = ((_1467_v1) ? (p2) : (p2));
      } else if (_source24.is_DC4) {
        let _1524___mcc_h2 = (_source24).cf9;
        let _1525_cf9 = _1524___mcc_h2;
        let _1526_v47;
        let _nw222 = new _module.C0();
        _nw222.__ctor(_module.__default.safeModuloInt(p2, p2), _1467_v1, p2, new BigNumber(483));
        _1526_v47 = _nw222;
        let _1527_v48;
        _1527_v48 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
        let _1528_v49;
        let _nw223 = new _module.C2();
        _nw223.__ctor(_module.__default.safeDivisionInt(p2, p2), _module.__default.fm36((_1526_v47).f28, (_1526_v47).f28, _1469_v3, globalState), (_1526_v47).f28, (_dafny.ZERO).minus(p2), (((_1527_v48).contains((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))])) ? ((_1527_v48).get((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))])) : (_1467_v1)));
        _1528_v49 = _nw223;
        _1467_v1 = (!(_dafny.MultiSet.fromElements((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], _1526_v47.f29)).equals(p3)) && (_1526_v47.f29);
        let _1529_v50;
        _1529_v50 = _dafny.Map.Empty.slice().updateUnsafe(((_1528_v49).f26).plus(new BigNumber(732)),_module.__default.fm1(p2, (_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], new BigNumber(469), globalState));
        _1529_v50 = (_1529_v50).update(_module.__default.safeDivisionInt((_1526_v47).f28, (_1528_v49).f26), ((_1526_v47).f28).isLessThanOrEqualTo(p2));
      } else {
        let _1530___mcc_h3 = (_source24).cf12;
        let _1531_cf12 = _1530___mcc_h3;
        let _1532_v51;
        let _nw224 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1532_v51 = _nw224;
        let _index182 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1532_v51).length));
        (_1532_v51)[_index182] = _1469_v3;
        let _index183 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1532_v51).length));
        (_1532_v51)[_index183] = new _dafny.CodePoint('v'.codePointAt(0));
        let _1533_v52;
        _1533_v52 = new BigNumber(605);
        _1533_v52 = p2;
        let _1534_v53;
        _1534_v53 = _dafny.Map.Empty.slice().updateUnsafe(false,!(_1467_v1));
        let _1535_v54;
        _1535_v54 = _dafny.Seq.of((p2).plus(p2), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("ivvcj")).length), new BigNumber((_1534_v53).length)), _module.__default.safeModuloInt((_dafny.ZERO).minus(p2), p2));
        let _1536_v55;
        _1536_v55 = _dafny.Set.fromElements((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))], _1467_v1, _1467_v1);
        _1533_v52 = (_1535_v54)[_module.__default.safeIndex(new BigNumber(((_1536_v55).Intersect(_dafny.Set.fromElements(_1467_v1))).length), new BigNumber((_1535_v54).length))];
        _1536_v55 = _1536_v55;
      }
      _1471_v5 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1471_v5, _1471_v5), _1471_v5), _module.__default.safeIndex((p2).minus(p2), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1471_v5, _1471_v5), _1471_v5)).length)), _1469_v3);
      let _1537_v56;
      _1537_v56 = _dafny.Seq.of((_1470_v4)[_module.__default.safeIndex(new BigNumber(601), new BigNumber((_1470_v4).length))]);
      let _1538_v57;
      _1538_v57 = _module.D0.create_DC1(_module.__default.fm2(globalState), _1471_v5, p2, (_1537_v56)[_module.__default.safeIndex(p2, new BigNumber((_1537_v56).length))]);
      r0 = ((!((_1538_v57).dtor_cf4)) ? (_1470_v4) : (_1470_v4));
      return r0;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return !(false);
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ce"), _dafny.Seq.UnicodeFromString("xawihm"));
    };
    m16(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = [];
      let _hi4 = (_this).f13;
      for (let _1539_i0 = (_this).f13; _1539_i0.isLessThan(_hi4); _1539_i0 = _1539_i0.plus(_dafny.ONE)) {
        let _1540_v0;
        _1540_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,((_this).f12).isLessThanOrEqualTo(_module.__default.fm2(globalState)));
        _1540_v0 = _1540_v0;
        let _1541_v1;
        _1541_v1 = new BigNumber(577);
        _1541_v1 = (_dafny.ZERO).minus(_1541_v1);
        let _1542_v2;
        _1542_v2 = _dafny.Seq.of(p0, p0, true, true);
        let _index184 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((p2).length));
        (p2)[_index184] = !((!(p0)) === ((_1542_v2)[_module.__default.safeIndex(_1539_i0, new BigNumber((_1542_v2).length))]));
        let _index185 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((p2).length));
        (p2)[_index185] = !(p0);
        let _1543_v3;
        let _init49 = ((_1544_v1, _1545_p2) => function (_1546_i1) {
          return ((_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,_1544_v1))).dtor_cf14).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1545_p2)[_module.__default.safeIndex(new BigNumber(272), new BigNumber((_1545_p2).length))],(_this).f12));
        })(_1541_v1, p2);
        let _nw225 = Array((new BigNumber(13)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw225.length); _i0_49++) {
          _nw225[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1543_v3 = _nw225;
        let _1547_v4;
        _1547_v4 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(272), new BigNumber((p2).length))],(_dafny.ZERO).minus(p3));
        let _index186 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1543_v3).length));
        (_1543_v3)[_index186] = _1547_v4;
        let _index187 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1543_v3).length));
        (_1543_v3)[_index187] = (_1547_v4).Merge(_1547_v4);
      }
      let _1548_i2;
      _1548_i2 = _dafny.ZERO;
      L5: {
        while (p0) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1548_i2)) {
              break L5;
            }
            _1548_i2 = (_1548_i2).plus(_dafny.ONE);
            let _1549_v5;
            _1549_v5 = new BigNumber(-181);
            _1549_v5 = _module.__default.safeDivisionInt(p3, _module.__default.fm2(globalState));
            let _1550_v6;
            _1550_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-437));
            let _1551_v7;
            _1551_v7 = _module.D3.create_DC8((_1550_v6).update(p0, new BigNumber(-342)));
            let _1552_v8;
            _1552_v8 = _module.D3.create_DC10(_1551_v7);
            let _1553_v9;
            _1553_v9 = _module.D3.create_DC10(_1551_v7);
            let _1554_v10;
            _1554_v10 = _module.D3.create_DC10(_1552_v8);
            let _source25 = _1554_v10;
            if (_source25.is_DC9) {
              let _1555___mcc_h0 = (_source25).cf15;
              let _1556___mcc_h1 = (_source25).cf16;
              let _1557___mcc_h2 = (_source25).cf17;
              let _1558___mcc_h3 = (_source25).cf18;
              let _1559_cf18 = _1558___mcc_h3;
              let _1560_cf17 = _1557___mcc_h2;
              let _1561_cf16 = _1556___mcc_h1;
              let _1562_cf15 = _1555___mcc_h0;
              let _1563_v11;
              _1563_v11 = _module.D0.create_DC1((_dafny.ZERO).minus(p3), p1, (_this).f13, p0);
              _1560_cf17 = (_1549_v5).minus((_dafny.ZERO).minus((_1563_v11).dtor_cf3));
              r1 = ((true) ? (p2) : (p2));
              let _1564_v12;
              let _1565_v13;
              let _1566_v14;
              let _out19;
              let _out20;
              let _out21;
              let _outcollector9 = (_this).m17((((_1550_v6).contains(p0)) ? ((_1550_v6).get(p0)) : ((_this).f13)), p0, globalState);
              _out19 = _outcollector9[0];
              _out20 = _outcollector9[1];
              _out21 = _outcollector9[2];
              _1564_v12 = _out19;
              _1565_v13 = _out20;
              _1566_v14 = _out21;
              let _1567_v15;
              _1567_v15 = _module.D0.create_DC3(new BigNumber(801), p0);
              _1564_v12 = (_1567_v15).dtor_cf8;
            } else if (_source25.is_DC8) {
              let _1568___mcc_h4 = (_source25).cf14;
              let _1569_cf14 = _1568___mcc_h4;
              let _1570_v16;
              _1570_v16 = _dafny.Seq.UnicodeFromString("otl");
              let _1571_v17;
              _1571_v17 = _dafny.Set.fromElements((_this).f13);
              let _1572_v18;
              _1572_v18 = _dafny.Seq.of(_1549_v5, new BigNumber(-121), (_dafny.ZERO).minus(new BigNumber(-505)), (_this).f13, p3);
              let _1573_v19;
              _1573_v19 = new _dafny.CodePoint('h'.codePointAt(0));
              let _1574_v20;
              _1574_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1573_v19,p0);
              let _1575_v22;
              _1575_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v17,(_this).fm8(new BigNumber(((_1572_v18)).length), _1574_v20, function () {
                let _coll69 = new _dafny.Set();
                for (const _compr_69 of (_1572_v18).Elements) {
                  let _1576_v21 = _compr_69;
                  if (_dafny.Seq.contains(_1572_v18, _1576_v21)) {
                    _coll69.add((_1576_v21).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)));
                  }
                }
                return _coll69;
              }(), globalState));
              _1570_v16 = (((_1575_v22).contains(_1571_v17)) ? ((_1575_v22).get(_1571_v17)) : (_dafny.Seq.Concat(p1, p1)));
              _1549_v5 = (_dafny.ZERO).minus((_this).f13);
              _1549_v5 = (_this).f12;
              let _index188 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((p2).length));
              (p2)[_index188] = ((_this).f13).isEqualTo((_this).f13);
              let _1577_v23;
              _1577_v23 = _dafny.MultiSet.fromElements(p0);
              let _1578_v24;
              _1578_v24 = _dafny.Seq.of(p0);
              let _index189 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((p2).length));
              (p2)[_index189] = !((_1577_v23).Union(_dafny.MultiSet.FromArray(_1578_v24))).contains(p0);
            } else {
              let _1579___mcc_h5 = (_source25).cf19;
              let _1580_cf19 = _1579___mcc_h5;
              let _1581_v25;
              _1581_v25 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ob")).length));
              let _rhs190 = (_dafny.ZERO).minus((_this).f13);
              let _rhs191 = (_1549_v5).multipliedBy((_this).f12);
              let _rhs192 = true;
              let _rhs193 = !(p0);
              let _rhs194 = _1581_v25;
              let _lhs105 = globalState;
              let _lhs106 = globalState;
              let _lhs107 = globalState;
              _1549_v5 = _rhs190;
              _1549_v5 = _rhs191;
              _lhs105.f8 = _rhs192;
              _lhs106.f8 = _rhs193;
              _lhs107.f5 = _rhs194;
              _1549_v5 = ((_1581_v25)[_module.__default.safeIndex(p3, new BigNumber((_1581_v25).length))]).multipliedBy(p3);
              let _1582_v26;
              _1582_v26 = _dafny.Set.fromElements(true, p0);
              let _1583_v27;
              _1583_v27 = _dafny.Seq.of(p0, p0, (_1582_v26).IsSubsetOf(_1582_v26));
              let _rhs195 = p0;
              let _rhs196 = p0;
              let _rhs197 = _module.__default.safeModuloInt((_this).f13, (_this).f12);
              let _rhs198 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
              let _rhs199 = new BigNumber((_1583_v27).length);
              let _lhs108 = globalState;
              let _lhs109 = globalState;
              _lhs108.f8 = _rhs195;
              _lhs109.f8 = _rhs196;
              _1549_v5 = _rhs197;
              _1549_v5 = _rhs198;
              _1549_v5 = _rhs199;
              (globalState).f8 = p0;
            }
            let _1584_v29;
            _1584_v29 = _dafny.Map.Empty.slice().updateUnsafe(p3,function () {
              let _coll70 = new _dafny.Map();
              for (const _compr_70 of _dafny.IntegerRange(new BigNumber(615), new BigNumber(-716))) {
                let _1585_v28 = _compr_70;
                if (((new BigNumber(615)).isLessThanOrEqualTo(_1585_v28)) && ((_1585_v28).isLessThan(new BigNumber(-716)))) {
                  _coll70.push([(_1585_v28).minus(p3),(_this).f12]);
                }
              }
              return _coll70;
            }());
            let _1586_v30;
            _1586_v30 = _dafny.Set.fromElements((_this).f13, _module.__default.fm2(globalState), p3, (_this).f12);
            let _1587_v31;
            _1587_v31 = _module.D3.create_DC9((_this).f12, _1584_v29, _1549_v5, new BigNumber((_1586_v30).length));
            let _1588_v32;
            _1588_v32 = _dafny.Seq.of(_1549_v5, p3, (_1587_v31).dtor_cf17, _module.__default.fm2(globalState));
            (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1588_v32, _dafny.Seq.of((_this).f13, (_this).f12)), _1588_v32);
            (globalState).f8 = p0;
          }
        }
      }
      let _1589_v33;
      _1589_v33 = _module.D1.create_DC5(_dafny.MultiSet.fromElements((_this).f13), p0);
      let _1590_v34;
      _1590_v34 = _module.D1.create_DC6(_1589_v33);
      let _1591_v35;
      _1591_v35 = _module.D1.create_DC6(_1590_v34);
      _1591_v35 = _1591_v35;
      let _1592_v36;
      _1592_v36 = _dafny.Seq.UnicodeFromString("paujkw");
      _1592_v36 = _dafny.Seq.UnicodeFromString("hy");
      let _1593_v37;
      let _1594_v38;
      let _out22;
      let _out23;
      let _outcollector10 = _module.__default.m0(globalState);
      _out22 = _outcollector10[0];
      _out23 = _outcollector10[1];
      _1593_v37 = _out22;
      _1594_v38 = _out23;
      let _1595_v40;
      _1595_v40 = _dafny.Seq.of((_this).f13);
      let _hi5 = new BigNumber((_module.__default.fm22(function () {
        let _coll71 = new _dafny.Map();
        for (const _compr_71 of (_1595_v40).Elements) {
          let _1596_v39 = _compr_71;
          if (_dafny.Seq.contains(_1595_v40, _1596_v39)) {
            _coll71.push([_module.__default.safeDivisionInt(_1596_v39, new BigNumber(285)),p3]);
          }
        }
        return _coll71;
      }(), globalState)).length);
      for (let _1597_i3 = (_this).f13; _1597_i3.isLessThan(_hi5); _1597_i3 = _1597_i3.plus(_dafny.ONE)) {
        (globalState).f8 = !(p0);
        let _1598_v41;
        _1598_v41 = new BigNumber(905);
        _1598_v41 = _1598_v41;
        let _1599_v42;
        _1599_v42 = _dafny.MultiSet.fromElements(_1598_v41);
        let _rhs200 = _1594_v38;
        let _rhs201 = (_dafny.ZERO).minus((_this).f12);
        let _rhs202 = (((_dafny.MultiSet.fromElements((_this).f12)).IsDisjointFrom(_1599_v42)) ? (!(p0)) : (!(p0)));
        let _rhs203 = _dafny.Seq.Concat(p1, p1);
        let _lhs110 = globalState;
        _lhs110.f8 = _rhs200;
        _1598_v41 = _rhs201;
        _1594_v38 = _rhs202;
        _1592_v36 = _rhs203;
        _1598_v41 = (((_this).f12).plus(new BigNumber((p1).length))).plus(_1597_i3);
      }
      let _1600_v43;
      _1600_v43 = _dafny.Seq.of(true);
      let _1601_v44;
      _1601_v44 = _dafny.MultiSet.fromElements(p3, p3);
      let _1602_v45;
      _1602_v45 = _dafny.MultiSet.fromElements((_module.D0.create_DC1(new BigNumber((_1601_v44).cardinality()), _1592_v36, new BigNumber(485), !(false))).dtor_cf1, p3);
      let _1603_v46;
      let _nw226 = Array((new BigNumber(14)).toNumber());
      _nw226[(_dafny.ZERO).toNumber()] = p3;
      _nw226[(_dafny.ONE).toNumber()] = p3;
      _nw226[(new BigNumber(2)).toNumber()] = new BigNumber((_1600_v43).length);
      _nw226[(new BigNumber(3)).toNumber()] = ((_1594_v38) ? ((_this).f12) : (p3));
      _nw226[(new BigNumber(4)).toNumber()] = (_this).f13;
      _nw226[(new BigNumber(5)).toNumber()] = new BigNumber(((_module.D1.create_DC5((_1601_v44).update(p3, _module.__default.abs((_this).f12)), _1594_v38)).dtor_cf10).cardinality());
      _nw226[(new BigNumber(6)).toNumber()] = p3;
      _nw226[(new BigNumber(7)).toNumber()] = new BigNumber(-5);
      _nw226[(new BigNumber(8)).toNumber()] = (_this).f13;
      _nw226[(new BigNumber(9)).toNumber()] = (p3).plus(new BigNumber(90));
      _nw226[(new BigNumber(10)).toNumber()] = (_this).f13;
      _nw226[(new BigNumber(11)).toNumber()] = p3;
      _nw226[(new BigNumber(12)).toNumber()] = ((_this).f12).plus((_this).f13);
      _nw226[(new BigNumber(13)).toNumber()] = p3;
      _1603_v46 = _nw226;
      r0 = _1603_v46;
      r1 = p2;
      return [r0, r1];
    }
    m17(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _1604_v0;
      _1604_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0);
      let _1605_v1;
      _1605_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1604_v0);
      let _1606_v3;
      _1606_v3 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1607_v4;
      _1607_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1606_v3,p1);
      let _1608_v5;
      _1608_v5 = _dafny.Set.fromElements(p0);
      let _1609_v6;
      _1609_v6 = _dafny.Seq.UnicodeFromString("agru");
      let _1610_v7;
      _1610_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1604_v0, (((_1605_v1).contains(p0)) ? ((_1605_v1).get(p0)) : (function () {
        let _coll72 = new _dafny.Map();
        for (const _compr_72 of _dafny.IntegerRange(new BigNumber(243), new BigNumber(483))) {
          let _1611_v2 = _compr_72;
          if (((new BigNumber(243)).isLessThanOrEqualTo(_1611_v2)) && ((_1611_v2).isLessThan(new BigNumber(483)))) {
            _coll72.push([(_1611_v2).minus((_this).f12),(_this).f13]);
          }
        }
        return _coll72;
      }()))),_dafny.Seq.IsProperPrefixOf((_this).fm8((_this).f13, _1607_v4, _1608_v5, globalState), _1609_v6));
      let _1612_v8;
      _1612_v8 = _dafny.Set.fromElements((_1604_v0).update((_this).f12, (_dafny.ZERO).minus((_this).f12)));
      _1610_v7 = (_1610_v7).update((_1612_v8).Union(_module.__default.fm23(p1, new BigNumber(-496), globalState)), true);
      let _1613_v9;
      _1613_v9 = _dafny.Seq.of((_this).f13);
      let _1614_v10;
      _1614_v10 = _dafny.Seq.Concat(_1613_v9, _dafny.Seq.of(new BigNumber(897), new BigNumber(70), new BigNumber(924)));
      let _source26 = _1614_v10;
      let _1615___mcc_h0 = _source26;
      let _1616_cf13 = _1615___mcc_h0;
      _1604_v0 = (_1604_v0).update((_this).f13, (_this).f12);
      if (p1) {
        r0 = p1;
        r1 = new BigNumber((_1616_cf13).length);
        r1 = (_this).f13;
        r1 = new BigNumber((function () {
          let _coll73 = new _dafny.Set();
          for (const _compr_73 of _dafny.IntegerRange(new BigNumber(524), new BigNumber(131))) {
            let _1617_v11 = _compr_73;
            if (((new BigNumber(524)).isLessThanOrEqualTo(_1617_v11)) && ((_1617_v11).isLessThan(new BigNumber(131)))) {
              _coll73.add((_1617_v11).multipliedBy((_this).f13));
            }
          }
          return _coll73;
        }()).length);
        let _1618_v12;
        let _nw227 = new _module.C2();
        _nw227.__ctor((_this).f12, _module.__default.fm36((_this).f12, new BigNumber(804), new _dafny.CodePoint('d'.codePointAt(0)), globalState), new BigNumber((_1609_v6).length), (_this).f12, (_1608_v5).IsSubsetOf(_1608_v5));
        _1618_v12 = _nw227;
      } else {
        let _1619_v13;
        _1619_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.update(_1613_v9, _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1613_v9).length)), (_this).f13), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.update(_1613_v9, _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1613_v9).length)), (_this).f13)).length)), new BigNumber((_1609_v6).length)),p1);
        _1619_v13 = _1619_v13;
        r1 = _module.__default.safeDivisionInt(p0, p0);
        (globalState).f8 = !(((p1) ? (new BigNumber((_dafny.Set.fromElements(_1606_v3, _1606_v3)).length)) : ((_this).f12))).isEqualTo((_this).f12);
        let _1620_v14;
        let _nw228 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1620_v14 = _nw228;
        let _index190 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1620_v14).length));
        (_1620_v14)[_index190] = new BigNumber((_module.__default.fm0(globalState)).length);
        let _index191 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1620_v14).length));
        (_1620_v14)[_index191] = (_this).f13;
        let _1621_v15;
        let _nw229 = Array((new BigNumber(11)).toNumber()).fill(_module.D3.Default());
        _1621_v15 = _nw229;
        let _1622_v16;
        _1622_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f13);
        let _1623_v17;
        _1623_v17 = _module.D3.create_DC8(_1622_v16);
        let _index192 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1621_v15).length));
        (_1621_v15)[_index192] = _1623_v17;
        let _index193 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1621_v15).length));
        (_1621_v15)[_index193] = _1623_v17;
      }
      (globalState).f8 = p1;
      let _1624_v18;
      _1624_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm36(p0, new BigNumber(854), _1606_v3, globalState));
      let _1625_v19;
      let _nw230 = new _module.C2();
      _nw230.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f12), (_this).f12), (((_1624_v18).contains(p1)) ? ((_1624_v18).get(p1)) : (_dafny.Seq.UnicodeFromString("oblpvb"))), p0, (_this).f13, p1);
      _1625_v19 = _nw230;
      let _1626_v20;
      _1626_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p1);
      _1626_v20 = (function () {
        let _coll74 = new _dafny.Map();
        for (const _compr_74 of (_1626_v20).Keys.Elements) {
          let _1627_v21 = _compr_74;
          if ((_1626_v20).contains(_1627_v21)) {
            _coll74.push([(_1627_v21).multipliedBy((_this).f12),p1]);
          }
        }
        return _coll74;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),p1));
      r1 = p0;
      let _1628_v22;
      let _nw231 = Array((_dafny.ONE).toNumber()).fill(_module.D7.Default());
      _1628_v22 = _nw231;
      let _rhs204 = (_this).f13;
      let _rhs205 = (_this).f13;
      let _rhs206 = ((!(p1)) ? (!(!(p1)) || (p1)) : (p1));
      let _rhs207 = _1628_v22;
      let _rhs208 = _dafny.Seq.contains(_1609_v6, _1606_v3);
      let _lhs111 = globalState;
      r1 = _rhs204;
      r1 = _rhs205;
      _lhs111.f8 = _rhs206;
      _1628_v22 = _rhs207;
      r0 = _rhs208;
      let _1629_v23;
      let _nw232 = Array((new BigNumber(3)).toNumber());
      _nw232[(_dafny.ZERO).toNumber()] = !(p1) || (p1);
      _nw232[(_dafny.ONE).toNumber()] = p1;
      _nw232[(new BigNumber(2)).toNumber()] = p1;
      _1629_v23 = _nw232;
      _1629_v23 = _1629_v23;
      let _1630_v24;
      _1630_v24 = _dafny.Seq.of(p1);
      r0 = (_1630_v24)[_module.__default.safeIndex((_this).f12, new BigNumber((_1630_v24).length))];
      r1 = p0;
      r2 = _1609_v6;
      return [r0, r1, r2];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f23 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f23) {
      let _this = this;
      (_this)._f23 = f23;
      return;
    }
    fm21(p0, p1, globalState) {
      let _this = this;
      return _module.D0.create_DC3((new BigNumber((_dafny.Set.fromElements(false, true)).length)).multipliedBy(new BigNumber(18)), !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(297),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(492))).cardinality())),new BigNumber(144))).length)).isEqualTo(new BigNumber(-162)));
    };
    m14(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _module.D0.Default();
      let _1631_v0;
      _1631_v0 = true;
      (globalState).f8 = _1631_v0;
      let _1632_v1;
      _1632_v1 = new BigNumber(556);
      _1632_v1 = _1632_v1;
      let _1633_i0;
      _1633_i0 = _dafny.ZERO;
      L6: {
        while (_module.__default.fm1((_dafny.ZERO).minus(_1632_v1), _1631_v0, (_dafny.ZERO).minus(_1632_v1), globalState)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1633_i0)) {
              break L6;
            }
            _1633_i0 = (_1633_i0).plus(_dafny.ONE);
            let _1634_v2;
            _1634_v2 = _dafny.Seq.of(_dafny.Seq.of(_1631_v0, !(_1631_v0)));
            let _1635_v3;
            _1635_v3 = _dafny.Seq.UnicodeFromString("ypssnf");
            let _1636_v4;
            let _nw233 = new _module.C0();
            _nw233.__ctor(((_dafny.ZERO).minus(_1632_v1)).plus(new BigNumber((_module.__default.fm36(_1632_v1, _1632_v1, _module.__default.fm45(new BigNumber((_1634_v2).length), globalState), globalState)).length)), _1631_v0, _module.__default.safeDivisionInt(new BigNumber((_1635_v3).length), _1632_v1), _1632_v1);
            _1636_v4 = _nw233;
            let _1637_v5;
            _1637_v5 = _dafny.Set.fromElements(_1636_v4.f29, _1631_v0);
            let _1638_v6;
            _1638_v6 = _dafny.MultiSet.fromElements(new BigNumber((_1637_v5).length));
            let _1639_v7;
            _1639_v7 = _module.D1.create_DC5(_1638_v6, _1631_v0);
            let _1640_v8;
            let _nw234 = Array((new BigNumber(3)).toNumber());
            _nw234[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_1632_v1);
            _nw234[(_dafny.ONE).toNumber()] = (_1639_v7).dtor_cf10;
            _nw234[(new BigNumber(2)).toNumber()] = _1638_v6;
            _1640_v8 = _nw234;
            let _1641_v9;
            _1641_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1640_v8,!(_1636_v4.f29));
            _1641_v9 = (_1641_v9).update(_1640_v8, _1631_v0);
            let _1642_v10;
            _1642_v10 = _dafny.MultiSet.fromElements(_1631_v0, _1636_v4.f29, false);
            let _rhs209 = _1631_v0;
            let _rhs210 = _1635_v3;
            let _rhs211 = ((_1636_v4).f28).minus(((_dafny.ZERO).minus(_1632_v1)).minus((_1636_v4).f28));
            let _rhs212 = _1631_v0;
            let _rhs213 = _1642_v10;
            let _lhs112 = globalState;
            _1631_v0 = _rhs209;
            _1635_v3 = _rhs210;
            _1632_v1 = _rhs211;
            _lhs112.f8 = _rhs212;
            _1642_v10 = _rhs213;
            _1632_v1 = (_1636_v4).f28;
          }
        }
      }
      let _1643_v11;
      _1643_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1631_v0,_1631_v0);
      let _1644_v12;
      _1644_v12 = _dafny.MultiSet.fromElements(_1643_v11, _1643_v11);
      let _1645_v13;
      _1645_v13 = _1644_v12;
      let _1646_v14;
      _1646_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1645_v13,new _dafny.CodePoint('r'.codePointAt(0)));
      let _1647_v15;
      _1647_v15 = _dafny.Set.fromElements(_1646_v14);
      _1647_v15 = (_1647_v15).Intersect(((_1631_v0) ? (_dafny.Set.fromElements(_1646_v14)) : (_1647_v15)));
      let _1648_i1;
      _1648_i1 = _dafny.ZERO;
      L7: {
        while (_1631_v0) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1648_i1)) {
              break L7;
            }
            _1648_i1 = (_1648_i1).plus(_dafny.ONE);
            let _1649_v16;
            let _nw235 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
            _1649_v16 = _nw235;
            let _1650_v17;
            let _init50 = ((_1651_v0) => function (_1652_i2) {
              return _1651_v0;
            })(_1631_v0);
            let _nw236 = Array((new BigNumber(6)).toNumber());
            for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw236.length); _i0_50++) {
              _nw236[_i0_50] = _init50(new BigNumber(_i0_50));
            }
            _1650_v17 = _nw236;
            let _1653_v18;
            _1653_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1650_v17,_1631_v0);
            let _index194 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1649_v16).length));
            (_1649_v16)[_index194] = _1653_v18;
            let _index195 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1649_v16).length));
            (_1649_v16)[_index195] = _dafny.Map.Empty.slice().updateUnsafe(_1650_v17,_1631_v0);
            if (((_1632_v1).isEqualTo(_1632_v1)) && (_1631_v0)) {
              _1632_v1 = ((_this).f23)[_module.__default.safeIndex(new BigNumber(267), new BigNumber(((_this).f23).length))];
              let _1654_v19;
              _1654_v19 = new _dafny.CodePoint('v'.codePointAt(0));
              let _1655_v20;
              _1655_v20 = _dafny.Seq.UnicodeFromString("jia");
              let _1656_v21;
              let _init51 = ((_1657_v1) => function (_1658_i3) {
                return (_1658_i3).plus(_1657_v1);
              })(_1632_v1);
              let _nw237 = Array((new BigNumber(23)).toNumber());
              for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw237.length); _i0_51++) {
                _nw237[_i0_51] = _init51(new BigNumber(_i0_51));
              }
              _1656_v21 = _nw237;
              let _rhs214 = !_dafny.Seq.contains(_1655_v20, ((true) ? (_1654_v19) : (_1654_v19)));
              let _rhs215 = _1631_v0;
              let _rhs216 = _1656_v21;
              let _rhs217 = _1632_v1;
              let _lhs113 = globalState;
              r0 = _rhs214;
              _lhs113.f8 = _rhs215;
              r1 = _rhs216;
              _1632_v1 = _rhs217;
              _1632_v1 = _1632_v1;
              let _1659_v22;
              _1659_v22 = _dafny.Seq.of(_1631_v0, _1631_v0);
              let _1660_v23;
              let _nw238 = new _module.C2();
              _nw238.__ctor(_1632_v1, _1655_v20, new BigNumber((_1659_v22).length), ((true) ? (_1632_v1) : (_1632_v1)), !(_1631_v0));
              _1660_v23 = _nw238;
              let _rhs218 = _1660_v23;
              let _rhs219 = _1631_v0;
              let _lhs114 = globalState;
              _1660_v23 = _rhs218;
              _lhs114.f8 = _rhs219;
              let _1661_v24;
              let _nw239 = new _module.C6();
              _nw239.__ctor();
              _1661_v24 = _nw239;
            } else {
              let _rhs220 = ((_1631_v0) ? (_1631_v0) : (_1631_v0));
              let _rhs221 = _1632_v1;
              let _lhs115 = globalState;
              _lhs115.f8 = _rhs220;
              _1632_v1 = _rhs221;
              let _1662_v25;
              _1662_v25 = _dafny.Seq.UnicodeFromString("fxdmehy");
              let _1663_v26;
              _1663_v26 = new _dafny.CodePoint('t'.codePointAt(0));
              let _1664_v27;
              _1664_v27 = _dafny.Map.Empty.slice().updateUnsafe(true,_1662_v25);
              let _1665_v28;
              _1665_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1663_v26,_1632_v1);
              let _1666_v29;
              _1666_v29 = _dafny.Seq.of(false);
              let _1667_v30;
              let _nw240 = Array((new BigNumber(21)).toNumber());
              _nw240[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1662_v25, _1662_v25);
              _nw240[(_dafny.ONE).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1662_v25, _module.__default.safeIndex(_1632_v1, new BigNumber((_1662_v25).length)), _1663_v26), _dafny.Seq.UnicodeFromString("yyrd"));
              _nw240[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(731)), ((_1668_v26) => function (_1669_i4) {
                return _1668_v26;
              })(_1663_v26)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-275)), function (_1670_i5) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              }));
              _nw240[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1662_v25, _1662_v25);
              _nw240[(new BigNumber(5)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(6)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(7)).toNumber()] = (((_1664_v27).contains(_1631_v0)) ? ((_1664_v27).get(_1631_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), ((_1671_v26) => function (_1672_i6) {
                return _1671_v26;
              })(_1663_v26))));
              _nw240[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1662_v25, _1662_v25);
              _nw240[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1662_v25, _1662_v25);
              _nw240[(new BigNumber(10)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(11)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(12)).toNumber()] = _module.__default.fm36(new BigNumber((_1665_v28).length), _1632_v1, _module.__default.fm45(_1632_v1, globalState), globalState);
              _nw240[(new BigNumber(13)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(14)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1662_v25, _dafny.Seq.UnicodeFromString("mrlgffk"));
              _nw240[(new BigNumber(16)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(17)).toNumber()] = _module.__default.fm36(_1632_v1, new BigNumber((_1666_v29).length), _1663_v26, globalState);
              _nw240[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-427)), ((_1673_v26) => function (_1674_i7) {
                return _1673_v26;
              })(_1663_v26)), _1662_v25);
              _nw240[(new BigNumber(19)).toNumber()] = _1662_v25;
              _nw240[(new BigNumber(20)).toNumber()] = _1662_v25;
              _1667_v30 = _nw240;
              _1667_v30 = _1667_v30;
              let _1675_v31;
              _1675_v31 = _module.D6.create_DC19(_1631_v0, _1631_v0, _1631_v0, _1631_v0, _1666_v29);
              let _index196 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1650_v17).length));
              (_1650_v17)[_index196] = _dafny.areEqual(_1666_v29, (_1675_v31).dtor_cf37);
              let _1676_v32;
              let _nw241 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _1676_v32 = _nw241;
              let _index197 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length));
              (_1676_v32)[_index197] = _1632_v1;
              let _1677_v33;
              _1677_v33 = _dafny.Set.fromElements(_1632_v1, _1632_v1, _1632_v1);
              let _1678_v34;
              _1678_v34 = _dafny.Set.fromElements(_1677_v33, _1677_v33);
              let _1679_v35;
              _1679_v35 = _dafny.Seq.of(_1678_v34);
              let _1680_v36;
              _1680_v36 = _dafny.MultiSet.fromElements(new BigNumber((_1666_v29).length));
              let _1681_v37;
              _1681_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1631_v0,_1632_v1);
              let _1682_v38;
              _1682_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1632_v1,_module.__default.fm1(_1632_v1, _1631_v0, new BigNumber((_1681_v37).length), globalState));
              let _index198 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1650_v17).length));
              let _index199 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length));
              let _rhs222 = new _dafny.CodePoint('l'.codePointAt(0));
              let _rhs223 = ((((_1649_v16)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_1649_v16).length))]).contains(_1650_v17)) ? (((_1649_v16)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_1649_v16).length))]).get(_1650_v17)) : (_1631_v0));
              let _rhs224 = _module.__default.fm1(_1632_v1, _1631_v0, (_dafny.ZERO).minus(new BigNumber(((_1679_v35)[_module.__default.safeIndex(_1632_v1, new BigNumber((_1679_v35).length))]).length)), globalState);
              let _rhs225 = ((true) ? (new _dafny.CodePoint('i'.codePointAt(0))) : (((_module.__default.fm1(new BigNumber((_1680_v36).cardinality()), (((_1682_v38).contains(_1632_v1)) ? ((_1682_v38).get(_1632_v1)) : (_1631_v0)), _1632_v1, globalState)) ? (_1663_v26) : (_1663_v26))));
              let _rhs226 = _1632_v1;
              let _lhs116 = _1650_v17;
              let _lhs117 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1650_v17).length));
              let _lhs118 = globalState;
              let _lhs119 = _1676_v32;
              let _lhs120 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length));
              _1663_v26 = _rhs222;
              _lhs116[_lhs117] = _rhs223;
              _lhs118.f8 = _rhs224;
              _1663_v26 = _rhs225;
              _lhs119[_lhs120] = _rhs226;
              let _1683_v39;
              _1683_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1676_v32)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length))],_1663_v26);
              let _1684_v40;
              let _nw242 = new _module.C7();
              _nw242.__ctor((_1676_v32)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length))], ((_this).f23)[_module.__default.safeIndex(new BigNumber((_1683_v39).length), new BigNumber(((_this).f23).length))]);
              _1684_v40 = _nw242;
              let _index200 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length));
              (_1676_v32)[_index200] = (_1676_v32)[_module.__default.safeIndex(new BigNumber(858), new BigNumber((_1676_v32).length))];
            }
            if (_1631_v0) {
              let _1685_v41;
              _1685_v41 = _dafny.MultiSet.fromElements(_1631_v0, !(_1631_v0), _1631_v0);
              let _1686_v42;
              _1686_v42 = _module.D7.create_DC21(_1685_v41);
              let _1687_v43;
              _1687_v43 = _module.D7.create_DC24(_1686_v42);
              let _1688_v44;
              _1688_v44 = _dafny.Map.Empty.slice().updateUnsafe(!(_1631_v0),_1686_v42);
              let _pat_let_tv59 = _1686_v42;
              _1687_v43 = function (_pat_let40_0) {
                return function (_1689_dt__update__tmp_h0) {
                  return function (_pat_let41_0) {
                    return function (_1690_dt__update_hcf46_h0) {
                      return _module.D7.create_DC24(_1690_dt__update_hcf46_h0);
                    }(_pat_let41_0);
                  }(_pat_let_tv59);
                }(_pat_let40_0);
              }(_module.D7.create_DC24((((_1688_v44).contains(_1631_v0)) ? ((_1688_v44).get(_1631_v0)) : (_1686_v42))));
              let _1691_v45;
              _1691_v45 = _dafny.Seq.UnicodeFromString("jgvpeq");
              let _1692_v46;
              let _nw243 = new _module.C2();
              _nw243.__ctor(_1632_v1, _1691_v45, _1632_v1, _1632_v1, _1631_v0);
              _1692_v46 = _nw243;
              let _1693_v47;
              _1693_v47 = _dafny.Seq.of(_1692_v46, _1692_v46);
              let _1694_v48;
              let _nw244 = Array((new BigNumber(6)).toNumber());
              _nw244[(_dafny.ZERO).toNumber()] = _1632_v1;
              _nw244[(_dafny.ONE).toNumber()] = _1632_v1;
              _nw244[(new BigNumber(2)).toNumber()] = new BigNumber((_1693_v47).length);
              _nw244[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_1632_v1);
              _nw244[(new BigNumber(4)).toNumber()] = _1632_v1;
              _nw244[(new BigNumber(5)).toNumber()] = _1632_v1;
              _1694_v48 = _nw244;
              let _1695_v49;
              _1695_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1694_v48,new BigNumber(-758));
              _1695_v49 = (_1695_v49).update(_1694_v48, new BigNumber((_1691_v45).length));
              let _1696_v50;
              _1696_v50 = new _dafny.CodePoint('x'.codePointAt(0));
              let _1697_v51;
              _1697_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v45,_1696_v50);
              _1697_v51 = (_1697_v51).update(_dafny.Seq.Concat(_1691_v45, _1691_v45), _1696_v50);
              _1631_v0 = _1631_v0;
              let _1698_v52;
              _1698_v52 = _dafny.MultiSet.fromElements(_1632_v1, _1632_v1, _1632_v1, _1632_v1);
              _1698_v52 = _1698_v52;
            } else {
              (globalState).f8 = _1631_v0;
              let _1699_v53;
              let _nw245 = new _module.C5();
              _nw245.__ctor(((_module.__default.fm1(_1632_v1, true, _module.__default.fm2(globalState), globalState)) ? ((_dafny.ZERO).minus(_module.__default.fm2(globalState))) : (new BigNumber(((_this).f23).length))));
              _1699_v53 = _nw245;
              let _1700_v54;
              _1700_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1631_v0,_dafny.Seq.of(_1631_v0));
              let _1701_v55;
              _1701_v55 = new _dafny.CodePoint('q'.codePointAt(0));
              let _1702_v56;
              _1702_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1701_v55,_1632_v1);
              let _1703_v57;
              _1703_v57 = _dafny.MultiSet.fromElements(_1702_v56);
              let _1704_v58;
              _1704_v58 = _dafny.Seq.of(false);
              _1632_v1 = (_dafny.ZERO).minus(new BigNumber(((((_1700_v54).contains(!((_dafny.MultiSet.fromElements(_1702_v56, _1702_v56)).IsSubsetOf(_1703_v57)))) ? ((_1700_v54).get(!((_dafny.MultiSet.fromElements(_1702_v56, _1702_v56)).IsSubsetOf(_1703_v57)))) : (_dafny.Seq.Concat(_1704_v58, _1704_v58)))).length));
              _1632_v1 = (_1699_v53).f24;
              _1632_v1 = (_1699_v53).f24;
            }
            _1632_v1 = _module.__default.fm2(globalState);
          }
        }
      }
      let _1705_v59;
      _1705_v59 = _dafny.MultiSet.fromElements(_1631_v0, _1631_v0, false, _1631_v0);
      let _1706_v60;
      _1706_v60 = _module.D4.create_DC13(new BigNumber(758));
      let _pat_let_tv60 = _1705_v59;
      let _pat_let_tv61 = _1705_v59;
      let _pat_let_tv62 = _1705_v59;
      _1705_v59 = function (_source27) {
        if (_source27.is_DC12) {
          let _1707___mcc_h0 = (_source27).cf21;
          let _1708___mcc_h1 = (_source27).cf22;
          let _1709___mcc_h2 = (_source27).cf23;
          let _1710_cf23 = _1709___mcc_h2;
          let _1711_cf22 = _1708___mcc_h1;
          let _1712_cf21 = _1707___mcc_h0;
          return _pat_let_tv60;
        } else if (_source27.is_DC13) {
          let _1713___mcc_h3 = (_source27).cf24;
          let _1714_cf24 = _1713___mcc_h3;
          return _pat_let_tv61;
        } else {
          let _1715___mcc_h4 = (_source27).cf20;
          let _1716_cf20 = _1715___mcc_h4;
          return _pat_let_tv62;
        }
      }(_1706_v60);
      r0 = _1631_v0;
      let _1717_v61;
      let _nw246 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1717_v61 = _nw246;
      r1 = _1717_v61;
      let _1718_v62;
      _1718_v62 = _module.D0.create_DC3(_1632_v1, _1631_v0);
      let _1719_v63;
      _1719_v63 = _dafny.Seq.of(_1631_v0);
      let _pat_let_tv63 = _1705_v59;
      let _pat_let_tv64 = _1719_v63;
      let _pat_let_tv65 = _1632_v1;
      let _pat_let_tv66 = _1719_v63;
      let _pat_let_tv67 = _1631_v0;
      let _pat_let_tv68 = _1705_v59;
      r2 = function (_pat_let42_0) {
        return function (_1720_dt__update__tmp_h1) {
          return function (_pat_let43_0) {
            return function (_1721_dt__update_hcf7_h0) {
              return _module.D0.create_DC3(_1721_dt__update_hcf7_h0, (_1720_dt__update__tmp_h1).dtor_cf8);
            }(_pat_let43_0);
          }((((_pat_let_tv68).contains(true)) ? ((_pat_let_tv63).get(true)) : (new BigNumber((_dafny.Seq.update(_pat_let_tv64, _module.__default.safeIndex(_pat_let_tv65, new BigNumber((_pat_let_tv66).length)), _pat_let_tv67)).length))));
        }(_pat_let42_0);
      }(_1718_v62);
      return [r0, r1, r2];
    }
    m15(globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Map.Empty;
      let _1722_v0;
      let _init52 = function (_1723_i0) {
        return _dafny.Seq.UnicodeFromString("dslysgcu");
      };
      let _nw247 = Array((new BigNumber(11)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw247.length); _i0_52++) {
        _nw247[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _1722_v0 = _nw247;
      let _1724_v1;
      _1724_v1 = _dafny.Seq.UnicodeFromString("barxrs");
      let _1725_v2;
      _1725_v2 = _dafny.Seq.of(_1724_v1, _1724_v1);
      let _1726_v3;
      _1726_v3 = new BigNumber(168);
      let _index201 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1722_v0).length));
      (_1722_v0)[_index201] = _dafny.Seq.Concat((_1725_v2)[_module.__default.safeIndex(_1726_v3, new BigNumber((_1725_v2).length))], _1724_v1);
      let _1727_v4;
      _1727_v4 = true;
      let _1728_v5;
      _1728_v5 = _dafny.Seq.of(_1727_v4, false, _1727_v4, (_1726_v3).isEqualTo(_1726_v3), _1727_v4);
      let _1729_v6;
      let _nw248 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _1729_v6 = _nw248;
      let _index202 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length));
      (_1729_v6)[_index202] = _1726_v3;
      let _1730_v7;
      _1730_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1727_v4,new BigNumber(455));
      let _index203 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1722_v0).length));
      let _index204 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length));
      let _rhs227 = _1724_v1;
      let _rhs228 = _module.__default.fm39(!(_1727_v4), _1727_v4, globalState);
      let _rhs229 = (_module.__default.safeModuloInt(_1726_v3, new BigNumber(390))).isLessThanOrEqualTo(new BigNumber(136));
      let _rhs230 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_1731_i1) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length);
      let _rhs231 = _module.__default.safeModuloInt(new BigNumber(765), (_1726_v3).multipliedBy(((_this).f23)[_module.__default.safeIndex(new BigNumber(((_1730_v7).update(_1727_v4, _1726_v3)).length), new BigNumber(((_this).f23).length))]));
      let _lhs121 = _1722_v0;
      let _lhs122 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1722_v0).length));
      let _lhs123 = globalState;
      let _lhs124 = _1729_v6;
      let _lhs125 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length));
      _lhs121[_lhs122] = _rhs227;
      _1728_v5 = _rhs228;
      _lhs123.f8 = _rhs229;
      _lhs124[_lhs125] = _rhs230;
      _1726_v3 = _rhs231;
      let _1732_v8;
      _1732_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))],_module.D3.create_DC8(_1730_v7));
      let _1733_v9;
      _1733_v9 = _dafny.Set.fromElements(_1727_v4, _1727_v4, _1727_v4);
      let _1734_v10;
      _1734_v10 = _dafny.Set.fromElements(new BigNumber((_1733_v9).length), new BigNumber(((_this).f23).length));
      let _1735_v11;
      _1735_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))],new BigNumber(548));
      let _1736_v12;
      _1736_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1726_v3,_1735_v11);
      let _1737_v13;
      _1737_v13 = _module.D3.create_DC10((((_1732_v8).contains(new BigNumber((_1734_v10).length))) ? ((_1732_v8).get(new BigNumber((_1734_v10).length))) : (_module.D3.create_DC9(_1726_v3, _1736_v12, (_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))], ((_this).f23)[_module.__default.safeIndex(_1726_v3, new BigNumber(((_this).f23).length))]))));
      _1737_v13 = _1737_v13;
      let _1738_v14;
      let _nw249 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1738_v14 = _nw249;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1738_v14).length))) {
        let _1739_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1739_i2)) && ((_1739_i2).isLessThan(new BigNumber((_1738_v14).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_1738_v14, (_1739_i2).toNumber(), ((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))]).isLessThanOrEqualTo(_1726_v3)));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _index205 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length));
      (_1729_v6)[_index205] = _1726_v3;
      r2 = (_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))];
      let _1740_v15;
      _1740_v15 = _dafny.MultiSet.fromElements(_1726_v3);
      let _1741_v16;
      _1741_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))],_1740_v15);
      r1 = (_dafny.ZERO).minus((((_1740_v15).contains((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))])) ? ((_1740_v15).get((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))])) : (_module.__default.safeModuloInt(new BigNumber((_1741_v16).length), (((_1735_v11).contains((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))])) ? ((_1735_v11).get((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))])) : (_1726_v3))))));
      let _1742_v17;
      _1742_v17 = _dafny.MultiSet.fromElements(true, _1727_v4);
      let _1743_v18;
      _1743_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1727_v4,_1742_v17);
      r0 = ((((_1743_v18).contains(!(_1727_v4))) ? ((_1743_v18).get(!(_1727_v4))) : (_dafny.MultiSet.FromArray(_1728_v5)))).Intersect((_dafny.MultiSet.fromElements(_1727_v4, _1727_v4, _1727_v4, _1727_v4, _1727_v4)).Union(_1742_v17));
      r1 = ((_dafny.ZERO).minus(_1726_v3)).minus(_module.__default.safeModuloInt(new BigNumber(-548), (_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))]));
      r2 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat((_1722_v0)[_module.__default.safeIndex(new BigNumber(668), new BigNumber((_1722_v0).length))], (_1722_v0)[_module.__default.safeIndex(new BigNumber(668), new BigNumber((_1722_v0).length))])).length), _1726_v3);
      let _1744_v19;
      _1744_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1727_v4,_1727_v4);
      let _1745_v20;
      _1745_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_1744_v19).contains(_1727_v4)) ? ((_1744_v19).get(_1727_v4)) : (_1727_v4)),_1727_v4);
      r3 = (((_1727_v4) ? (_1745_v20) : (_module.__default.fm42((_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))], (_1729_v6)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1729_v6).length))], globalState)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_1727_v4));
      return [r0, r1, r2, r3];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this.f21 = _dafny.Set.Empty;
      this.f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0, _module.T2];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f21, f22, f12, f13, f11, f20) {
      let _this = this;
      (_this).f21 = f21;
      (_this).f22 = f22;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f11 = f11;
      (_this)._f20 = f20;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(_this.f11, _this.f11, !(_this.f11), _this.f11))).IsProperSubsetOf((_dafny.Set.fromElements(_dafny.Set.fromElements(_this.f11, _this.f11), _dafny.Set.fromElements(true))).Difference(function () {
        let _coll75 = new _dafny.Set();
        for (const _compr_75 of (_dafny.Seq.of(_dafny.Set.fromElements(_this.f11, _this.f11, false, _this.f11, !(_this.f11)), _dafny.Set.fromElements(_this.f11, _this.f11))).Elements) {
          let _1746_v0 = _compr_75;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(_this.f11, _this.f11, false, _this.f11, !(_this.f11)), _dafny.Set.fromElements(_this.f11, _this.f11)), _1746_v0)) {
            _coll75.add(_1746_v0);
          }
        }
        return _coll75;
      }()));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      if (_this.f11) {
        return (_this).f20;
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), function (_1747_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        });
      }
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1((_this).f12, (_this).f20, (_this).f12, (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_this.f11, _this.f11, _this.f11)).length))).IsSubsetOf(_dafny.Set.fromElements((_this).f13)));
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return !(_this.f11);
    };
    fm18(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f11;
    };
    fm19(p0, p1, globalState) {
      let _this = this;
      return new _dafny.CodePoint('f'.codePointAt(0));
    };
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let _1748_v0;
      _1748_v0 = _dafny.Seq.UnicodeFromString("almbhp");
      _1748_v0 = _1748_v0;
      let _1749_v1;
      _1749_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,!(_this.f11));
      let _1750_v2;
      _1750_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,_this.f11);
      let _1751_v3;
      _1751_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_1749_v1).contains(_this.f11)) ? ((_1749_v1).get(_this.f11)) : ((((_1750_v2).contains(p0)) ? ((_1750_v2).get(p0)) : (_this.f11)))));
      let _1752_v4;
      _1752_v4 = _dafny.MultiSet.fromElements(_this.f11, _this.f11, _this.f11);
      let _1753_v5;
      _1753_v5 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f11)).length));
      let _1754_v6;
      _1754_v6 = _dafny.Set.fromElements(_1753_v5, _1753_v5);
      _1750_v2 = (_1750_v2).update(((((_1752_v4).contains(_this.f11)) ? ((_1752_v4).get(_this.f11)) : (new BigNumber((_1754_v6).length)))).plus((_dafny.ZERO).minus(p1)), !((_this).fm7(globalState)));
      if (!((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f12, _this.f22))).isEqualTo(_this.f22)) {
        let _1755_v7;
        _1755_v7 = _module.D0.create_DC3((p0).multipliedBy(p2), _this.f11);
        let _1756_v8;
        _1756_v8 = _dafny.Seq.of(false, _this.f11);
        let _1757_v9;
        let _nw250 = Array((new BigNumber(24)).toNumber());
        _nw250[(_dafny.ZERO).toNumber()] = _this.f11;
        _nw250[(_dafny.ONE).toNumber()] = _this.f11;
        _nw250[(new BigNumber(2)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(3)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(4)).toNumber()] = (!(_this.f11)) && (_this.f11);
        _nw250[(new BigNumber(5)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),_this.f11)).equals(_1749_v1);
        _nw250[(new BigNumber(6)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(7)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(8)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(9)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(10)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(11)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(12)).toNumber()] = true;
        _nw250[(new BigNumber(13)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(14)).toNumber()] = (true) || (!(_this.f11));
        _nw250[(new BigNumber(15)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(16)).toNumber()] = (_1756_v8)[_module.__default.safeIndex((_this).f13, new BigNumber((_1756_v8).length))];
        _nw250[(new BigNumber(17)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(18)).toNumber()] = (_this.f11) || (_this.f11);
        _nw250[(new BigNumber(19)).toNumber()] = (_this.f11) || (_this.f11);
        _nw250[(new BigNumber(20)).toNumber()] = false;
        _nw250[(new BigNumber(21)).toNumber()] = _this.f11;
        _nw250[(new BigNumber(22)).toNumber()] = (((_1750_v2).contains((_this).f13)) ? ((_1750_v2).get((_this).f13)) : (_this.f11));
        _nw250[(new BigNumber(23)).toNumber()] = !(!(_this.f11) || (_this.f11));
        _1757_v9 = _nw250;
        let _rhs232 = _1757_v9;
        let _rhs233 = _1755_v7;
        r0 = _rhs232;
        _1755_v7 = _rhs233;
        let _1758_v10;
        _1758_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber((_1748_v0).length));
        let _1759_v11;
        _1759_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1758_v10).length),p1);
        let _1760_v12;
        _1760_v12 = _dafny.MultiSet.fromElements((((_1759_v11).contains(p1)) ? ((_1759_v11).get(p1)) : (new BigNumber(((_this).f20).length))));
        let _1761_v13;
        _1761_v13 = _module.D1.create_DC5(_1760_v12, _this.f11);
        let _1762_v14;
        _1762_v14 = _module.D1.create_DC6(_1761_v13);
        let _1763_v15;
        _1763_v15 = _dafny.Seq.of(_1762_v14, _1762_v14);
        let _1764_v16;
        _1764_v16 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1765_v17;
        _1765_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1763_v15,_1764_v16);
        let _1766_v18;
        _1766_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(886),_1765_v17);
        let _rhs234 = ((p0).minus(new BigNumber((_module.__default.fm20(_this.f22, new BigNumber(((_this).f20).length), new BigNumber((_1749_v1).length), globalState)).length))).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f13))));
        let _rhs235 = (((_1760_v12).contains(_module.__default.safeDivisionInt(p2, p0))) ? ((_1760_v12).get(_module.__default.safeDivisionInt(p2, p0))) : ((_this).f12));
        let _rhs236 = (((_this.f11) ? (_this.f11) : (_this.f11))) === (_this.f11);
        let _rhs237 = (p2).multipliedBy(p2);
        let _rhs238 = new BigNumber((((((_1766_v18).contains((_this).f12)) ? ((_1766_v18).get((_this).f12)) : (_1765_v17))).Merge((_1765_v17).Merge(_1765_v17))).length);
        let _lhs126 = _this;
        let _lhs127 = _this;
        r1 = _rhs234;
        r1 = _rhs235;
        _lhs126.f11 = _rhs236;
        r1 = _rhs237;
        _lhs127.f22 = _rhs238;
        let _1767_v19;
        _1767_v19 = _dafny.Set.fromElements((_this).f13);
        let _1768_v20;
        _1768_v20 = _dafny.Seq.of((_1767_v19).Union(_1767_v19), _1767_v19, _dafny.Set.fromElements((_this).f13));
        let _index206 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length));
        (_1757_v9)[_index206] = !(!(_this.f11));
        let _1769_v21;
        _1769_v21 = _dafny.MultiSet.fromElements(p3);
        let _1770_v22;
        _1770_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1769_v21).cardinality()),_1767_v19);
        let _1771_v23;
        _1771_v23 = _dafny.Seq.of(_1757_v9, _1757_v9);
        let _1772_v24;
        _1772_v24 = _dafny.Set.fromElements((_1771_v23)[_module.__default.safeIndex(p0, new BigNumber((_1771_v23).length))]);
        let _1773_v25;
        _1773_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1764_v16,_this.f11);
        let _1774_v26;
        _1774_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1753_v5,(_this).fm8(new BigNumber((_1772_v24).length), _1773_v25, _1767_v19, globalState));
        let _index207 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length));
        let _rhs239 = _dafny.Seq.update(_dafny.Seq.Concat(_1768_v20, _dafny.Seq.Concat(_dafny.Seq.of(_1767_v19, _1767_v19), _1768_v20)), _module.__default.safeIndex(_module.__default.fm2(globalState), new BigNumber((_dafny.Seq.Concat(_1768_v20, _dafny.Seq.Concat(_dafny.Seq.of(_1767_v19, _1767_v19), _1768_v20))).length)), ((((_1770_v22).contains((_this).f13)) ? ((_1770_v22).get((_this).f13)) : (_dafny.Set.fromElements((((_1752_v4).contains(_this.f11)) ? ((_1752_v4).get(_this.f11)) : (new BigNumber(577))), p1)))).Union(_1767_v19));
        let _rhs240 = _this.f11;
        let _rhs241 = new BigNumber((_1774_v26).length);
        let _rhs242 = (_1756_v8)[_module.__default.safeIndex(new BigNumber((_1748_v0).length), new BigNumber((_1756_v8).length))];
        let _rhs243 = (_dafny.ZERO).minus(p2);
        let _lhs128 = _1757_v9;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length));
        let _lhs130 = _this;
        let _lhs131 = _this;
        let _lhs132 = _this;
        _1768_v20 = _rhs239;
        _lhs128[_lhs129] = _rhs240;
        _lhs130.f22 = _rhs241;
        _lhs131.f11 = _rhs242;
        _lhs132.f22 = _rhs243;
        if (_this.f11) {
          let _index208 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((p3).length));
          (p3)[_index208] = _this.f22;
          let _index209 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((p3).length));
          (p3)[_index209] = (new BigNumber((_dafny.MultiSet.fromElements(_this.f11, (_1757_v9)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length))])).cardinality())).plus((_this).f13);
          let _1775_v27;
          _1775_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_1748_v0);
          let _1776_v28;
          _1776_v28 = _module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1753_v5).length))).length), (_this).f20, (_this).f13, (_1757_v9)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length))]);
          _1748_v0 = (((_1775_v27).contains(false)) ? ((_1775_v27).get(false)) : ((_1776_v28).dtor_cf2));
          r1 = p0;
          let _1777_v29;
          _1777_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1750_v2,_1769_v21);
          _1777_v29 = (_1777_v29).update(_1751_v3, _1769_v21);
          let _1778_v30;
          let _init53 = function (_1779_i0) {
            return (_1779_i0).plus(_this.f22);
          };
          let _nw251 = Array((new BigNumber(16)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw251.length); _i0_53++) {
            _nw251[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _1778_v30 = _nw251;
          let _1780_v31;
          _1780_v31 = _dafny.Seq.of(_1778_v30);
          let _1781_v32;
          let _nw252 = Array((new BigNumber(18)).toNumber());
          _nw252[(_dafny.ZERO).toNumber()] = _1780_v31;
          _nw252[(_dafny.ONE).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1780_v31, _1780_v31);
          _nw252[(new BigNumber(3)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1780_v31, _1780_v31);
          _nw252[(new BigNumber(5)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(6)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1778_v30, _1778_v30), _1780_v31);
          _nw252[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1780_v31, _1780_v31);
          _nw252[(new BigNumber(9)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(10)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(11)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(12)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(13)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1780_v31, _1780_v31);
          _nw252[(new BigNumber(15)).toNumber()] = _1780_v31;
          _nw252[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_1780_v31, _module.__default.safeIndex(p1, new BigNumber((_1780_v31).length)), _1778_v30);
          _nw252[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_1778_v30), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_1778_v30)).length)), _1778_v30);
          _1781_v32 = _nw252;
          let _index210 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1781_v32).length));
          (_1781_v32)[_index210] = _dafny.Seq.Concat(_1780_v31, _dafny.Seq.of(_1778_v30, _1778_v30));
          let _index211 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1781_v32).length));
          (_1781_v32)[_index211] = _dafny.Seq.of(_1778_v30, _1778_v30, _1778_v30, _1778_v30, _1778_v30);
        } else {
          (globalState).f8 = (_1756_v8)[_module.__default.safeIndex((_this).f13, new BigNumber((_1756_v8).length))];
          let _1782_v33;
          _1782_v33 = _1753_v5;
          let _1783_v34;
          _1783_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1782_v33,_this.f22);
          _1783_v34 = (_1783_v34).update(_1782_v33, p2);
          let _1784_v35;
          let _nw253 = new _module.C2();
          _nw253.__ctor(((_this.f11) ? (new BigNumber((_dafny.Seq.of(_this.f22)).length)) : (p2)), _dafny.Seq.UnicodeFromString("tqwuuwg"), _module.__default.safeModuloInt(_this.f22, _module.__default.fm2(globalState)), new BigNumber((_1760_v12).cardinality()), !((p2).isLessThanOrEqualTo(new BigNumber(((_this).f20).length))));
          _1784_v35 = _nw253;
          let _index212 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length));
          (_1757_v9)[_index212] = _this.f11;
          let _1785_v36;
          _1785_v36 = _module.D6.create_DC17(false, p1);
          let _1786_v37;
          _1786_v37 = _module.D0.create_DC1((_1785_v36).dtor_cf31, (_this).f20, (_1784_v35).f26, (_1757_v9)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1757_v9).length))]);
          let _1787_v38;
          let _nw254 = Array((new BigNumber(4)).toNumber());
          _nw254[(_dafny.ZERO).toNumber()] = _1757_v9;
          _nw254[(_dafny.ONE).toNumber()] = _1757_v9;
          _nw254[(new BigNumber(2)).toNumber()] = _1757_v9;
          _nw254[(new BigNumber(3)).toNumber()] = _1757_v9;
          _1787_v38 = _nw254;
          let _1788_v39;
          let _1789_v40;
          let _out24;
          let _out25;
          let _outcollector11 = (_this).m12((_this).f20, _1786_v37, _1787_v38, globalState);
          _out24 = _outcollector11[0];
          _out25 = _outcollector11[1];
          _1788_v39 = _out24;
          _1789_v40 = _out25;
        }
        r1 = (_this).f12;
      } else {
        if (_this.f11) {
          let _1790_v41;
          _1790_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,new BigNumber(-956));
          _1790_v41 = ((_1790_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_this.f22))).Merge(_1790_v41);
          let _1791_v42;
          _1791_v42 = _dafny.Seq.of(false);
          (globalState).f8 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_1791_v42, _dafny.Seq.update(_dafny.Seq.of(_this.f11, _this.f11), _module.__default.safeIndex(_this.f22, new BigNumber((_dafny.Seq.of(_this.f11, _this.f11)).length)), _this.f11)), _module.__default.safeIndex(new BigNumber((_1791_v42).length), new BigNumber((_dafny.Seq.Concat(_1791_v42, _dafny.Seq.update(_dafny.Seq.of(_this.f11, _this.f11), _module.__default.safeIndex(_this.f22, new BigNumber((_dafny.Seq.of(_this.f11, _this.f11)).length)), _this.f11))).length)), _this.f11), _1791_v42));
          let _1792_v43;
          _1792_v43 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.of(false, _this.f11));
          let _1793_v44;
          _1793_v44 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1792_v43);
          _1793_v44 = (_1793_v44).update(p3, _1792_v43);
          let _1794_v45;
          let _init54 = function (_1795_i1) {
            return _module.D6.create_DC18(!(_this.f11));
          };
          let _nw255 = Array((new BigNumber(25)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw255.length); _i0_54++) {
            _nw255[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _1794_v45 = _nw255;
          let _1796_v46;
          _1796_v46 = _module.D6.create_DC18(!(_this.f11));
          let _index213 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1794_v45).length));
          (_1794_v45)[_index213] = _1796_v46;
          let _index214 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1794_v45).length));
          (_1794_v45)[_index214] = _1796_v46;
          let _1797_v47;
          let _nw256 = Array((new BigNumber(12)).toNumber()).fill([]);
          _1797_v47 = _nw256;
          _1797_v47 = _1797_v47;
        } else {
          (_this).f11 = _this.f11;
          let _1798_v48;
          let _nw257 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1798_v48 = _nw257;
          let _index215 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1798_v48).length));
          (_1798_v48)[_index215] = _this.f11;
          let _index216 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1798_v48).length));
          (_1798_v48)[_index216] = _this.f11;
          let _index217 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((p3).length));
          (p3)[_index217] = (_this).f12;
          let _index218 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((p3).length));
          (p3)[_index218] = (new BigNumber(694)).plus(new BigNumber(-748));
          (globalState).f8 = (_1798_v48)[_module.__default.safeIndex(new BigNumber(932), new BigNumber((_1798_v48).length))];
          let _1799_v49;
          let _nw258 = new _module.C2();
          _nw258.__ctor(p1, _1748_v0, _module.__default.fm2(globalState), (p2).plus(p0), _this.f11);
          _1799_v49 = _nw258;
        }
        if (_this.f11) {
          let _1800_v50;
          let _nw259 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1800_v50 = _nw259;
          let _1801_v51;
          _1801_v51 = _dafny.Seq.of(false);
          r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_module.D7.create_DC23(_1800_v50, _1801_v51, _this.f11, _this.f22)).dtor_cf45), ((_dafny.ZERO).minus(p0)).minus(p0));
          r1 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p0, p0), (_this).f13);
          let _1802_v52;
          let _init55 = function (_1803_i2) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          };
          let _nw260 = Array((new BigNumber(22)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw260.length); _i0_55++) {
            _nw260[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _1802_v52 = _nw260;
          _1802_v52 = _1802_v52;
          (globalState).f8 = false;
          let _index219 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((p3).length));
          (p3)[_index219] = p1;
          let _index220 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((p3).length));
          (p3)[_index220] = _this.f22;
        } else {
          r1 = (_dafny.ZERO).minus(((_this.f11) ? (p2) : (p0)));
          _1752_v4 = _1752_v4;
          _1748_v0 = (_this).f20;
          let _index221 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((p3).length));
          (p3)[_index221] = p2;
          let _index222 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((p3).length));
          let _rhs244 = _this.f11;
          let _rhs245 = _this.f22;
          let _rhs246 = (_dafny.ZERO).minus(p1);
          let _lhs133 = _this;
          let _lhs134 = _this;
          let _lhs135 = p3;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((p3).length));
          _lhs133.f11 = _rhs244;
          _lhs134.f22 = _rhs245;
          _lhs135[_lhs136] = _rhs246;
          let _1804_v53;
          let _nw261 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1804_v53 = _nw261;
          let _index223 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_1804_v53).length));
          (_1804_v53)[_index223] = true;
          let _1805_v54;
          _1805_v54 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1806_v55;
          _1806_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p1);
          let _1807_v56;
          _1807_v56 = _module.D0.create_DC1((_this).f13, (_this).f20, new BigNumber((_1806_v55).length), _this.f11);
          let _index224 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_1804_v53).length));
          let _rhs247 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1748_v0, _dafny.Seq.update(_1748_v0, _module.__default.safeIndex((_dafny.ZERO).minus(_this.f22), new BigNumber((_1748_v0).length)), _1805_v54)), _dafny.Seq.Concat(_1748_v0, (_1807_v56).dtor_cf2));
          let _rhs248 = ((false) ? (_1748_v0) : ((_this).f20));
          let _rhs249 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), ((_1808_v54) => function (_1809_i3) {
            return _1808_v54;
          })(_1805_v54)), new _dafny.CodePoint('i'.codePointAt(0)));
          let _lhs137 = _1804_v53;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_1804_v53).length));
          _1748_v0 = _rhs247;
          _1748_v0 = _rhs248;
          _lhs137[_lhs138] = _rhs249;
        }
        let _1810_v57;
        _1810_v57 = _dafny.Set.fromElements(_this.f11);
        let _1811_v58;
        _1811_v58 = _dafny.Seq.of(_1748_v0);
        let _rhs250 = (_this.f22).isEqualTo(_module.__default.fm2(globalState));
        let _rhs251 = (_this).f13;
        let _rhs252 = ((_this).f12).isEqualTo(p2);
        let _rhs253 = _dafny.Seq.Concat((_1811_v58)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f22), new BigNumber((_1811_v58).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_1812_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
        let _rhs254 = _1810_v57;
        let _lhs139 = globalState;
        let _lhs140 = _this;
        _lhs139.f8 = _rhs250;
        r1 = _rhs251;
        _lhs140.f11 = _rhs252;
        _1748_v0 = _rhs253;
        _1810_v57 = _rhs254;
        let _1813_v59;
        _1813_v59 = new _dafny.CodePoint('p'.codePointAt(0));
        let _1814_v60;
        let _nw262 = Array((new BigNumber(2)).toNumber()).fill([]);
        _1814_v60 = _nw262;
        let _1815_v61;
        _1815_v61 = _module.D7.create_DC22(p0, _1814_v60);
        let _1816_v62;
        _1816_v62 = _module.D9.create_DC27(!(_this.f11), _1815_v61, p0, _this.f11);
        let _1817_v63;
        _1817_v63 = _dafny.Seq.of(_this.f11, _this.f11);
        let _1818_v64;
        _1818_v64 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1813_v59),_this.f11),((_1816_v62).dtor_cf49) || ((_1817_v63)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length), new BigNumber((_1817_v63).length))]));
        let _1819_v65;
        _1819_v65 = _dafny.Set.fromElements(_1813_v59, _1813_v59);
        let _1820_v66;
        _1820_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1819_v65,_this.f11);
        let _1821_v68;
        _1821_v68 = _dafny.MultiSet.fromElements(_1813_v59);
        let _1822_v70;
        _1822_v70 = _dafny.Seq.of(function () {
          let _coll76 = new _dafny.Set();
          for (const _compr_76 of (_1821_v68).Elements) {
            let _1823_v69 = _compr_76;
            if ((_1821_v68).contains(_1823_v69)) {
              _coll76.add(_1823_v69);
            }
          }
          return _coll76;
        }());
        _1818_v64 = (_1818_v64).update((_1820_v66).Merge(function () {
          let _coll77 = new _dafny.Map();
          for (const _compr_77 of (_1822_v70).Elements) {
            let _1824_v67 = _compr_77;
            if (_dafny.Seq.contains(_1822_v70, _1824_v67)) {
              _coll77.push([_1824_v67,_this.f11]);
            }
          }
          return _coll77;
        }()), _this.f11);
        let _1825_v71;
        _1825_v71 = _module.D0.create_DC1(p1, (_this).f20, (_this).f12, (_module.__default.fm46(_this.f11, _this.f11, globalState)).dtor_cf21);
        let _1826_v72;
        let _nw263 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1826_v72 = _nw263;
        let _1827_v73;
        let _nw264 = Array((new BigNumber(3)).toNumber());
        _nw264[(_dafny.ZERO).toNumber()] = _1826_v72;
        _nw264[(_dafny.ONE).toNumber()] = _1826_v72;
        _nw264[(new BigNumber(2)).toNumber()] = _1826_v72;
        _1827_v73 = _nw264;
        let _1828_v74;
        _1828_v74 = _dafny.Seq.of(_1827_v73, _1827_v73, _1827_v73, _1827_v73, _1827_v73);
        let _1829_v75;
        let _1830_v76;
        let _out26;
        let _out27;
        let _outcollector12 = (_this).m12(_1748_v0, _1825_v71, (_1828_v74)[_module.__default.safeIndex(_this.f22, new BigNumber((_1828_v74).length))], globalState);
        _out26 = _outcollector12[0];
        _out27 = _outcollector12[1];
        _1829_v75 = _out26;
        _1830_v76 = _out27;
      }
      let _hi6 = new BigNumber((_1748_v0).length);
      for (let _1831_i5 = p1; _1831_i5.isLessThan(_hi6); _1831_i5 = _1831_i5.plus(_dafny.ONE)) {
        let _1832_v77;
        _1832_v77 = _dafny.Seq.of(_this.f11, true, _this.f11);
        let _1833_v78;
        _1833_v78 = _dafny.Seq.of((_1832_v77)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1753_v5)).cardinality()), new BigNumber((_1832_v77).length))], _this.f11);
        let _1834_v79;
        let _nw265 = Array((new BigNumber(11)).toNumber());
        _nw265[(_dafny.ZERO).toNumber()] = false;
        _nw265[(_dafny.ONE).toNumber()] = (true) || ((_this).fm7(globalState));
        _nw265[(new BigNumber(2)).toNumber()] = _this.f11;
        _nw265[(new BigNumber(3)).toNumber()] = _this.f11;
        _nw265[(new BigNumber(4)).toNumber()] = _this.f11;
        _nw265[(new BigNumber(5)).toNumber()] = _this.f11;
        _nw265[(new BigNumber(6)).toNumber()] = ((_this.f11) ? (_this.f11) : (_this.f11));
        _nw265[(new BigNumber(7)).toNumber()] = (((_1832_v77)[_module.__default.safeIndex(p0, new BigNumber((_1832_v77).length))]) ? (_this.f11) : (_this.f11));
        _nw265[(new BigNumber(8)).toNumber()] = !(_this.f11);
        _nw265[(new BigNumber(9)).toNumber()] = _this.f11;
        _nw265[(new BigNumber(10)).toNumber()] = (_this.f11) || (!(_this.f11));
        _1834_v79 = _nw265;
        let _index225 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length));
        (_1834_v79)[_index225] = _this.f11;
        let _1835_v80;
        _1835_v80 = _dafny.Set.fromElements(!(_this.f11), _this.f11);
        let _1836_v81;
        _1836_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1835_v80,p1);
        let _index226 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length));
        (_1834_v79)[_index226] = ((((_1836_v81).contains(_1835_v80)) ? ((_1836_v81).get(_1835_v80)) : (_this.f22))).isLessThan((_this.f22).plus(new BigNumber((_1749_v1).length)));
        let _index227 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length));
        (_1834_v79)[_index227] = (_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))];
        let _index228 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length));
        (_1834_v79)[_index228] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), function (_1837_i6) {
          return new BigNumber(-205);
        }), _1753_v5);
        let _1838_v82;
        _1838_v82 = _module.D0.create_DC2(new BigNumber((_dafny.Seq.of(_this.f11)).length), p3);
        let _1839_v83;
        _1839_v83 = _dafny.Seq.of(p3, p3);
        let _1840_v84;
        let _nw266 = Array((new BigNumber(29)).toNumber());
        _nw266[(_dafny.ZERO).toNumber()] = p3;
        _nw266[(_dafny.ONE).toNumber()] = p3;
        _nw266[(new BigNumber(2)).toNumber()] = p3;
        _nw266[(new BigNumber(3)).toNumber()] = p3;
        _nw266[(new BigNumber(4)).toNumber()] = p3;
        _nw266[(new BigNumber(5)).toNumber()] = p3;
        _nw266[(new BigNumber(6)).toNumber()] = p3;
        _nw266[(new BigNumber(7)).toNumber()] = p3;
        _nw266[(new BigNumber(8)).toNumber()] = p3;
        _nw266[(new BigNumber(9)).toNumber()] = p3;
        _nw266[(new BigNumber(10)).toNumber()] = p3;
        _nw266[(new BigNumber(11)).toNumber()] = p3;
        _nw266[(new BigNumber(12)).toNumber()] = p3;
        _nw266[(new BigNumber(13)).toNumber()] = p3;
        _nw266[(new BigNumber(14)).toNumber()] = p3;
        _nw266[(new BigNumber(15)).toNumber()] = p3;
        _nw266[(new BigNumber(16)).toNumber()] = p3;
        _nw266[(new BigNumber(17)).toNumber()] = p3;
        _nw266[(new BigNumber(18)).toNumber()] = p3;
        _nw266[(new BigNumber(19)).toNumber()] = (_1838_v82).dtor_cf6;
        _nw266[(new BigNumber(20)).toNumber()] = p3;
        _nw266[(new BigNumber(21)).toNumber()] = p3;
        _nw266[(new BigNumber(22)).toNumber()] = (_1839_v83)[_module.__default.safeIndex(_this.f22, new BigNumber((_1839_v83).length))];
        _nw266[(new BigNumber(23)).toNumber()] = p3;
        _nw266[(new BigNumber(24)).toNumber()] = p3;
        _nw266[(new BigNumber(25)).toNumber()] = p3;
        _nw266[(new BigNumber(26)).toNumber()] = p3;
        _nw266[(new BigNumber(27)).toNumber()] = p3;
        _nw266[(new BigNumber(28)).toNumber()] = p3;
        _1840_v84 = _nw266;
        let _1841_v85;
        _1841_v85 = _module.D7.create_DC22((_this).f12, _1840_v84);
        let _pat_let_tv69 = _1840_v84;
        let _source28 = function (_pat_let44_0) {
          return function (_1842_dt__update__tmp_h0) {
            return function (_pat_let45_0) {
              return function (_1843_dt__update_hcf41_h0) {
                return _module.D7.create_DC22((_1842_dt__update__tmp_h0).dtor_cf40, _1843_dt__update_hcf41_h0);
              }(_pat_let45_0);
            }(_pat_let_tv69);
          }(_pat_let44_0);
        }(_1841_v85);
        if (_source28.is_DC22) {
          let _1844___mcc_h0 = (_source28).cf40;
          let _1845___mcc_h1 = (_source28).cf41;
          let _1846_cf41 = _1845___mcc_h1;
          let _1847_cf40 = _1844___mcc_h0;
          let _1848_v86;
          let _nw267 = new _module.C8();
          _nw267.__ctor(_dafny.Seq.Concat(_1753_v5, _1753_v5));
          _1848_v86 = _nw267;
          _1847_cf40 = ((_1848_v86).f23)[_module.__default.safeIndex((((_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))]) ? (p2) : ((_this).f12)), new BigNumber(((_1848_v86).f23).length))];
          let _1849_v87;
          _1849_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1748_v0,_1834_v79);
          let _1850_v88;
          _1850_v88 = _dafny.Seq.of(_1834_v79);
          _1849_v87 = (_1849_v87).update(_dafny.Seq.Concat((_this).f20, (_this).f20), (_1850_v88)[_module.__default.safeIndex((_this).f13, new BigNumber((_1850_v88).length))]);
          let _1851_v89;
          _1851_v89 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1831_i5,!((_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))])));
          let _1852_v91;
          _1852_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll78 = new _dafny.Set();
            for (const _compr_78 of (_1851_v89).Elements) {
              let _1853_v90 = _compr_78;
              if ((_1851_v89).contains(_1853_v90)) {
                _coll78.add(_1853_v90);
              }
            }
            return _coll78;
          }()).length),_this.f22);
          let _1854_v92;
          _1854_v92 = _module.D0.create_DC3((_this).f12, (_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))]);
          let _1855_v93;
          _1855_v93 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1856_v94;
          _1856_v94 = _dafny.MultiSet.fromElements((_this).f13, _1831_i5);
          let _1857_v95;
          let _nw268 = Array((new BigNumber(18)).toNumber());
          _nw268[(_dafny.ZERO).toNumber()] = (_this).f12;
          _nw268[(_dafny.ONE).toNumber()] = new BigNumber(((_this).f20).length);
          _nw268[(new BigNumber(2)).toNumber()] = new BigNumber(((_1852_v91).Merge(_1852_v91)).length);
          _nw268[(new BigNumber(3)).toNumber()] = new BigNumber(((_1835_v80).Intersect(_module.__default.fm4(_1847_cf40, _1854_v92, new BigNumber((_1749_v1).length), _1855_v93, globalState))).length);
          _nw268[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1748_v0, _1748_v0)).length);
          _nw268[(new BigNumber(5)).toNumber()] = p0;
          _nw268[(new BigNumber(6)).toNumber()] = new BigNumber((_1856_v94).cardinality());
          _nw268[(new BigNumber(7)).toNumber()] = _this.f22;
          _nw268[(new BigNumber(8)).toNumber()] = _1847_cf40;
          _nw268[(new BigNumber(9)).toNumber()] = (_this).f12;
          _nw268[(new BigNumber(10)).toNumber()] = p2;
          _nw268[(new BigNumber(11)).toNumber()] = _1831_i5;
          _nw268[(new BigNumber(12)).toNumber()] = (_this.f22).plus(p1);
          _nw268[(new BigNumber(13)).toNumber()] = ((((_1856_v94).contains(p0)) ? ((_1856_v94).get(p0)) : (_this.f22))).multipliedBy(new BigNumber(781));
          _nw268[(new BigNumber(14)).toNumber()] = _1831_i5;
          _nw268[(new BigNumber(15)).toNumber()] = p0;
          _nw268[(new BigNumber(16)).toNumber()] = _this.f22;
          _nw268[(new BigNumber(17)).toNumber()] = p1;
          _1857_v95 = _nw268;
          _1857_v95 = p3;
        } else if (_source28.is_DC23) {
          let _1858___mcc_h2 = (_source28).cf42;
          let _1859___mcc_h3 = (_source28).cf43;
          let _1860___mcc_h4 = (_source28).cf44;
          let _1861___mcc_h5 = (_source28).cf45;
          let _1862_cf45 = _1861___mcc_h5;
          let _1863_cf44 = _1860___mcc_h4;
          let _1864_cf43 = _1859___mcc_h3;
          let _1865_cf42 = _1858___mcc_h2;
          let _index229 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length));
          (_1834_v79)[_index229] = _1863_cf44;
          let _1866_v96;
          let _init56 = function (_1867_i7) {
            return _module.__default.safeModuloInt(_1867_i7, new BigNumber((_dafny.MultiSet.fromElements((_this).f20)).cardinality()));
          };
          let _nw269 = Array((new BigNumber(2)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw269.length); _i0_56++) {
            _nw269[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _1866_v96 = _nw269;
          let _rhs255 = _1866_v96;
          let _rhs256 = _1748_v0;
          _1866_v96 = _rhs255;
          _1748_v0 = _rhs256;
          let _1868_v97;
          _1868_v97 = _module.D1.create_DC4(_1839_v83);
          let _1869_v98;
          _1869_v98 = _module.D1.create_DC6(_1868_v97);
          let _1870_v99;
          let _nw270 = Array((new BigNumber(3)).toNumber());
          _nw270[(_dafny.ZERO).toNumber()] = _1869_v98;
          _nw270[(_dafny.ONE).toNumber()] = _1869_v98;
          _nw270[(new BigNumber(2)).toNumber()] = _module.D1.create_DC6(_1868_v97);
          _1870_v99 = _nw270;
          let _index230 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_1870_v99).length));
          (_1870_v99)[_index230] = _1869_v98;
          let _index231 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_1870_v99).length));
          (_1870_v99)[_index231] = _1869_v98;
          r1 = ((_this).f13).minus(p1);
        } else if (_source28.is_DC21) {
          let _1871___mcc_h6 = (_source28).cf39;
          let _1872_cf39 = _1871___mcc_h6;
          (_this).f22 = (p2).minus(p1);
          let _index232 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p3).length));
          (p3)[_index232] = new BigNumber((_dafny.Seq.of(_this.f22, (_this).f13)).length);
          let _index233 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p3).length));
          (p3)[_index233] = new BigNumber(((_1750_v2).Merge(_1751_v3)).length);
          let _1873_v100;
          _1873_v100 = _module.D7.create_DC24(_module.D7.create_DC23(_1840_v84, _1833_v78, (_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))], p2));
          let _1874_v101;
          let _nw271 = Array((new BigNumber(8)).toNumber());
          _nw271[(_dafny.ZERO).toNumber()] = _1873_v100;
          _nw271[(_dafny.ONE).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(2)).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(3)).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(4)).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(5)).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(6)).toNumber()] = _1873_v100;
          _nw271[(new BigNumber(7)).toNumber()] = _1873_v100;
          _1874_v101 = _nw271;
          let _index234 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1874_v101).length));
          (_1874_v101)[_index234] = _1873_v100;
          let _index235 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1874_v101).length));
          let _rhs257 = (_this).f20;
          let _rhs258 = _1832_v77;
          let _rhs259 = _1873_v100;
          let _lhs141 = _1874_v101;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1874_v101).length));
          _1748_v0 = _rhs257;
          _1833_v78 = _rhs258;
          _lhs141[_lhs142] = _rhs259;
          let _1875_v102;
          let _nw272 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1875_v102 = _nw272;
          let _1876_v103;
          _1876_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1875_v102,(_dafny.ZERO).minus(_module.__default.fm2(globalState)));
          let _index236 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p3).length));
          (p3)[_index236] = new BigNumber((_1876_v103).length);
        } else {
          let _1877___mcc_h7 = (_source28).cf46;
          let _1878_cf46 = _1877___mcc_h7;
          r1 = _module.__default.safeDivisionInt(p0, (new BigNumber((_1832_v77).length)).minus((_this).f13));
          let _1879_v104;
          _1879_v104 = _dafny.Seq.of(_1834_v79, _1834_v79, _1834_v79);
          let _1880_v105;
          _1880_v105 = _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_1831_i5));
          let _1881_v106;
          _1881_v106 = _dafny.Map.Empty.slice().updateUnsafe((_1834_v79)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1834_v79).length))],p1);
          r0 = (_1879_v104)[_module.__default.safeIndex(new BigNumber((((_1880_v105).dtor_cf14).Merge(_1881_v106)).length), new BigNumber((_1879_v104).length))];
          let _1882_v107;
          _1882_v107 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1883_v108;
          _1883_v108 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f12).minus(p1),_1882_v107);
          _1883_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1831_i5,_1882_v107);
          let _1884_v109;
          let _nw273 = new _module.C1();
          _nw273.__ctor(p0);
          _1884_v109 = _nw273;
        }
      }
      let _1885_v110;
      let _nw274 = Array((new BigNumber(26)).toNumber()).fill(false);
      _1885_v110 = _nw274;
      let _index237 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
      (_1885_v110)[_index237] = (p0).isLessThanOrEqualTo(new BigNumber(226));
      let _1886_v111;
      _1886_v111 = _dafny.MultiSet.fromElements(_1753_v5);
      let _index238 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
      (_1885_v110)[_index238] = !((_1886_v111).IsSubsetOf(_1886_v111));
      if ((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) {
        (_this).f11 = (_this).fm18(_this.f11, (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], p1, p0, globalState);
        let _1887_v112;
        let _nw275 = new _module.C8();
        _nw275.__ctor(_dafny.Seq.Concat(_1753_v5, _1753_v5));
        _1887_v112 = _nw275;
        let _index239 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
        (_1885_v110)[_index239] = !(_dafny.Seq.IsProperPrefixOf((_this).f20, (_this).f20));
        r1 = (_this).f12;
        let _1888_v113;
        _1888_v113 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,p2);
        let _source29 = _module.__default.fm53((_1888_v113).Merge(_1888_v113), _module.__default.safeModuloInt(p0, (_this).f12), globalState);
        if (_source29.is_DC9) {
          let _1889___mcc_h8 = (_source29).cf15;
          let _1890___mcc_h9 = (_source29).cf16;
          let _1891___mcc_h10 = (_source29).cf17;
          let _1892___mcc_h11 = (_source29).cf18;
          let _1893_cf18 = _1892___mcc_h11;
          let _1894_cf17 = _1891___mcc_h10;
          let _1895_cf16 = _1890___mcc_h9;
          let _1896_cf15 = _1889___mcc_h8;
          let _1897_v114;
          _1897_v114 = _dafny.MultiSet.fromElements((_this).f12, p1);
          let _index240 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
          (_1885_v110)[_index240] = !(!(!((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1897_v114).cardinality()), p1))).isEqualTo((_dafny.ZERO).minus(_1894_cf17))));
          let _1898_v115;
          let _nw276 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1898_v115 = _nw276;
          let _1899_v116;
          _1899_v116 = _module.D7.create_DC22(p0, _1898_v115);
          let _1900_v117;
          _1900_v117 = _module.D9.create_DC27(true, _1899_v116, new BigNumber(-105), (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]);
          let _1901_v118;
          _1901_v118 = _module.D9.create_DC28(_module.D9.create_DC28(_1900_v117));
          _1901_v118 = _module.D9.create_DC28(_1900_v117);
          let _1902_v119;
          _1902_v119 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1903_v120;
          let _nw277 = Array((new BigNumber(14)).toNumber());
          _nw277[(_dafny.ZERO).toNumber()] = _1902_v119;
          _nw277[(_dafny.ONE).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(2)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(3)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _nw277[(new BigNumber(5)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(6)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(7)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(8)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(9)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(10)).toNumber()] = _1902_v119;
          _nw277[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw277[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw277[(new BigNumber(13)).toNumber()] = _1902_v119;
          _1903_v120 = _nw277;
          let _1904_v121;
          _1904_v121 = _module.D13.create_DC36(_1903_v120);
          let _1905_v122;
          let _nw278 = Array((new BigNumber(6)).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = (_1904_v121).dtor_cf61;
          _nw278[(_dafny.ONE).toNumber()] = _1903_v120;
          _nw278[(new BigNumber(2)).toNumber()] = _1903_v120;
          _nw278[(new BigNumber(3)).toNumber()] = (_1904_v121).dtor_cf61;
          _nw278[(new BigNumber(4)).toNumber()] = _1903_v120;
          _nw278[(new BigNumber(5)).toNumber()] = _1903_v120;
          _1905_v122 = _nw278;
          let _index241 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1905_v122).length));
          (_1905_v122)[_index241] = _1903_v120;
          let _index242 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1905_v122).length));
          (_1905_v122)[_index242] = _1903_v120;
          (globalState).f5 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), ((_1906_cf17) => function (_1907_i8) {
            return _1906_cf17;
          })(_1894_cf17));
        } else if (_source29.is_DC8) {
          let _1908___mcc_h12 = (_source29).cf14;
          let _1909_cf14 = _1908___mcc_h12;
          (globalState).f8 = ((_this).fm6((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], new BigNumber(732), _this.f22, globalState)) === (_this.f11);
          let _1910_v123;
          let _init57 = ((_1911_v2) => function (_1912_i9) {
            return _1911_v2;
          })(_1750_v2);
          let _nw279 = Array((new BigNumber(9)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw279.length); _i0_57++) {
            _nw279[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _1910_v123 = _nw279;
          let _index243 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1910_v123).length));
          (_1910_v123)[_index243] = _1751_v3;
          let _index244 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1910_v123).length));
          (_1910_v123)[_index244] = _1750_v2;
          let _1913_v124;
          let _init58 = function (_1914_i10) {
            return (_1914_i10).minus((_this).f13);
          };
          let _nw280 = Array((new BigNumber(11)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw280.length); _i0_58++) {
            _nw280[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1913_v124 = _nw280;
          _1913_v124 = (((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) ? (_1913_v124) : (p3));
          let _1915_v125;
          let _nw281 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
          _1915_v125 = _nw281;
          let _index245 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1915_v125).length));
          (_1915_v125)[_index245] = _1752_v4;
          let _1916_v126;
          _1916_v126 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1753_v5);
          let _1917_v127;
          _1917_v127 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1916_v126).length),_dafny.MultiSet.fromElements(_this.f11));
          let _1918_v128;
          _1918_v128 = _dafny.MultiSet.fromElements(new BigNumber(-915));
          let _index246 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1915_v125).length));
          (_1915_v125)[_index246] = (((((_1917_v127).contains(new BigNumber((_1918_v128).cardinality()))) ? ((_1917_v127).get(new BigNumber((_1918_v128).cardinality()))) : (_1752_v4))).Union(_1752_v4)).update((_dafny.MultiSet.fromElements(_this.f11)).IsProperSubsetOf(_module.__default.fm30(!((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]), (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], new BigNumber(545), globalState)), _module.__default.abs((p1).multipliedBy(_module.__default.fm2(globalState))));
        } else {
          let _1919___mcc_h13 = (_source29).cf19;
          let _1920_cf19 = _1919___mcc_h13;
          let _1921_v129;
          let _nw282 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1921_v129 = _nw282;
          let _index247 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1921_v129).length));
          (_1921_v129)[_index247] = (_this).f20;
          let _1922_v130;
          _1922_v130 = _dafny.Seq.of((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]);
          let _1923_v131;
          let _nw283 = Array((new BigNumber(12)).toNumber()).fill([]);
          _1923_v131 = _nw283;
          let _1924_v132;
          _1924_v132 = _module.D7.create_DC22(new BigNumber((_1922_v130).length), _1923_v131);
          let _1925_v133;
          _1925_v133 = _module.D9.create_DC27(_this.f11, _1924_v132, new BigNumber(-258), _this.f11);
          let _1926_v134;
          _1926_v134 = _module.D9.create_DC28(_1925_v133);
          let _index248 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1921_v129).length));
          let _rhs260 = (_this).f20;
          let _rhs261 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1926_v134,_this.f22)).length);
          let _lhs143 = _1921_v129;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1921_v129).length));
          _lhs143[_lhs144] = _rhs260;
          r1 = _rhs261;
          let _index249 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
          (_1885_v110)[_index249] = !(((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) && (((_this).f13).isLessThanOrEqualTo(p2)));
          let _1927_v135;
          let _nw284 = Array((new BigNumber(23)).toNumber()).fill(_module.D5.Default());
          _1927_v135 = _nw284;
          let _1928_v136;
          _1928_v136 = _module.D5.create_DC14(_1921_v129);
          let _index250 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1927_v135).length));
          (_1927_v135)[_index250] = _1928_v136;
          let _index251 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1927_v135).length));
          (_1927_v135)[_index251] = _1928_v136;
          let _1929_v137;
          _1929_v137 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1930_v138;
          _1930_v138 = _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), _1929_v137, _1929_v137, _1929_v137);
          let _1931_v139;
          _1931_v139 = _dafny.Map.Empty.slice().updateUnsafe(_1930_v138,p1);
          let _1932_v140;
          _1932_v140 = _dafny.Seq.of(_1748_v0, _1748_v0);
          let _1933_v141;
          _1933_v141 = _module.D0.create_DC1((_dafny.ZERO).minus(_this.f22), (_1932_v140)[_module.__default.safeIndex((_this).f12, new BigNumber((_1932_v140).length))], (_this).f12, (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]);
          _1931_v139 = (_1931_v139).update(_1930_v138, ((_1933_v141).dtor_cf3).multipliedBy(_this.f22));
        }
      } else {
        let _1934_v142;
        let _nw285 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1934_v142 = _nw285;
        let _index252 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length));
        (_1934_v142)[_index252] = _1885_v110;
        let _1935_v143;
        _1935_v143 = _1885_v110;
        let _index253 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length));
        (_1934_v142)[_index253] = (_1935_v143);
        let _1936_v144;
        _1936_v144 = _dafny.Set.fromElements((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], true);
        let _index254 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length));
        (_1885_v110)[_index254] = (new BigNumber(((_1936_v144).Union(_dafny.Set.fromElements((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], _this.f11))).length)).isLessThanOrEqualTo(_module.__default.fm2(globalState));
        (_this).f22 = ((((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) ? (p2) : ((_this).f12))).minus((p2).multipliedBy(p2));
        let _1937_v145;
        _1937_v145 = _dafny.Seq.of((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]);
        let _index255 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length));
        let _rhs262 = _dafny.Seq.Concat(_1937_v145, _dafny.Seq.of((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]));
        let _rhs263 = (((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) ? (new BigNumber(((_1936_v144).Union(_1936_v144)).length)) : (_this.f22));
        let _rhs264 = !(((_1752_v4).Intersect(_1752_v4)).IsDisjointFrom((_module.__default.fm30((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], _this.f22, globalState)).update(_this.f11, _module.__default.abs(new BigNumber(877)))));
        let _rhs265 = !(!(true)) || ((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]);
        let _rhs266 = (((_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))]) ? ((_1934_v142)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length))]) : ((_1934_v142)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length))]));
        let _lhs145 = _this;
        let _lhs146 = globalState;
        let _lhs147 = _1934_v142;
        let _lhs148 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1934_v142).length));
        _1937_v145 = _rhs262;
        r1 = _rhs263;
        _lhs145.f11 = _rhs264;
        _lhs146.f8 = _rhs265;
        _lhs147[_lhs148] = _rhs266;
        (_this).m13(((_this).f12).plus(p1), (_1885_v110)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_1885_v110).length))], p3, globalState);
      }
      r0 = _1885_v110;
      r1 = (p1).plus(((_this).f12).plus(p2));
      let _1938_v147;
      _1938_v147 = _dafny.Set.fromElements((_this).f13);
      r2 = ((function () {
        let _coll79 = new _dafny.Set();
        for (const _compr_79 of _dafny.IntegerRange(new BigNumber(552), new BigNumber(597))) {
          let _1939_v146 = _compr_79;
          if (((new BigNumber(552)).isLessThanOrEqualTo(_1939_v146)) && ((_1939_v146).isLessThan(new BigNumber(597)))) {
            _coll79.add(_module.__default.safeDivisionInt(_1939_v146, new BigNumber((_1748_v0).length)));
          }
        }
        return _coll79;
      }()).Intersect(_1938_v147)).Union(function () {
        let _coll80 = new _dafny.Set();
        for (const _compr_80 of (_1753_v5).Elements) {
          let _1940_v148 = _compr_80;
          if (_dafny.Seq.contains(_1753_v5, _1940_v148)) {
            _coll80.add(_module.__default.safeModuloInt(_1940_v148, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length)));
          }
        }
        return _coll80;
      }());
      return [r0, r1, r2];
    }
    m12(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _rhs267 = !(false);
      let _rhs268 = (_this.f22).multipliedBy(_module.__default.safeModuloInt(_this.f22, _this.f22));
      let _lhs149 = _this;
      let _lhs150 = _this;
      _lhs149.f11 = _rhs267;
      _lhs150.f22 = _rhs268;
      let _1941_v0;
      let _nw286 = Array((new BigNumber(7)).toNumber()).fill(false);
      _1941_v0 = _nw286;
      let _index256 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length));
      (p2)[_index256] = _1941_v0;
      let _1942_v1;
      _1942_v1 = _dafny.Seq.of((false) && (_this.f11));
      let _index257 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length));
      let _rhs269 = !(_this.f11);
      let _rhs270 = (_1942_v1)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f12, _this.f22), new BigNumber((_1942_v1).length))];
      let _rhs271 = _1941_v0;
      let _lhs151 = p2;
      let _lhs152 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length));
      r0 = _rhs269;
      r1 = _rhs270;
      _lhs151[_lhs152] = _rhs271;
      let _1943_v2;
      let _nw287 = Array((new BigNumber(15)).toNumber()).fill(_module.D3.Default());
      _1943_v2 = _nw287;
      let _1944_v3;
      _1944_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f13);
      let _1945_v4;
      _1945_v4 = _module.D3.create_DC8(_1944_v3);
      let _index258 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1943_v2).length));
      (_1943_v2)[_index258] = _1945_v4;
      let _index259 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1943_v2).length));
      (_1943_v2)[_index259] = _1945_v4;
      let _1946_v5;
      _1946_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_dafny.Set.fromElements((_this).f12, _this.f22, _this.f22));
      let _1947_v6;
      _1947_v6 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f22));
      let _1948_v7;
      _1948_v7 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1942_v1).length)), (_this).f12);
      _1946_v5 = (_1946_v5).update(new BigNumber(((_1947_v6)[_module.__default.safeIndex((_this).f12, new BigNumber((_1947_v6).length))]).cardinality()), (_1948_v7).Union(_dafny.Set.fromElements(new BigNumber(835), (_this).f12)));
      let _hi7 = (_dafny.ZERO).minus(((_dafny.ZERO).minus((_this).f13)).minus((_this).f13));
      for (let _1949_i0 = (_this).f13; _1949_i0.isLessThan(_hi7); _1949_i0 = _1949_i0.plus(_dafny.ONE)) {
        let _1950_v8;
        _1950_v8 = new _dafny.CodePoint('b'.codePointAt(0));
        _1950_v8 = _1950_v8;
        if (false) {
          let _1951_v9;
          let _nw288 = new _module.C7();
          _nw288.__ctor(_this.f22, new BigNumber(868));
          _1951_v9 = _nw288;
          (_this).f22 = _module.__default.fm2(globalState);
          r1 = !(((_this).f13).isLessThan((_this).f13)) || (_this.f11);
          let _1952_v10;
          _1952_v10 = _dafny.MultiSet.fromElements(_this.f11, _this.f11);
          let _1953_v11;
          _1953_v11 = _module.D12.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(false,_1952_v10));
          _1953_v11 = _1953_v11;
          let _arr2 = (p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))];
          let _index260 = _module.__default.safeIndex(new BigNumber(428), new BigNumber(((p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))]).length));
          _arr2[_index260] = (_1948_v7).IsSubsetOf(_dafny.Set.fromElements((_this).f13));
          let _arr3 = (p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))];
          let _index261 = _module.__default.safeIndex(new BigNumber(428), new BigNumber(((p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))]).length));
          _arr3[_index261] = _this.f11;
        } else {
          (_this).f22 = new BigNumber(368);
          let _1954_v12;
          _1954_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1942_v1).length),_this.f22);
          let _1955_v13;
          _1955_v13 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1954_v12).update((_this).f13, _1949_i0));
          let _rhs272 = _1950_v8;
          let _rhs273 = _module.__default.fm54((((_1942_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1948_v7).length)), new BigNumber((_1942_v1).length))]) ? (!(false)) : (_this.f11)), globalState);
          _1950_v8 = _rhs272;
          _1955_v13 = _rhs273;
          let _1956_v14;
          let _nw289 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1956_v14 = _nw289;
          let _index262 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1956_v14).length));
          (_1956_v14)[_index262] = _module.__default.safeDivisionInt(new BigNumber(((_module.D0.create_DC1(new BigNumber(382), p0, (_this).f13, true)).dtor_cf2).length), _this.f22);
          let _1957_v15;
          let _nw290 = Array((new BigNumber(26)).toNumber()).fill([]);
          _1957_v15 = _nw290;
          let _1958_v16;
          let _init59 = function (_1959_i1) {
            return _dafny.MultiSet.fromElements((_this).f12);
          };
          let _nw291 = Array((new BigNumber(26)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw291.length); _i0_59++) {
            _nw291[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _1958_v16 = _nw291;
          let _index263 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1957_v15).length));
          (_1957_v15)[_index263] = _1958_v16;
          let _index264 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1956_v14).length));
          (_1956_v14)[_index264] = (_this).f12;
          let _index265 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1956_v14).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1957_v15).length));
          let _index267 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1956_v14).length));
          let _rhs274 = (_this).f13;
          let _rhs275 = (_module.__default.fm50(globalState)).equals(_1948_v7);
          let _rhs276 = _1958_v16;
          let _rhs277 = _1950_v8;
          let _rhs278 = _module.__default.fm2(globalState);
          let _lhs153 = _1956_v14;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1956_v14).length));
          let _lhs155 = _this;
          let _lhs156 = _1957_v15;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1957_v15).length));
          let _lhs158 = _1956_v14;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_1956_v14).length));
          _lhs153[_lhs154] = _rhs274;
          _lhs155.f11 = _rhs275;
          _lhs156[_lhs157] = _rhs276;
          _1950_v8 = _rhs277;
          _lhs158[_lhs159] = _rhs278;
          _1950_v8 = _module.__default.fm45(_1949_i0, globalState);
          let _1960_v17;
          let _nw292 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _1960_v17 = _nw292;
          let _1961_v18;
          _1961_v18 = _dafny.Seq.of((_this).f13, new BigNumber(653));
          let _index268 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1960_v17).length));
          (_1960_v17)[_index268] = _1961_v18;
          let _index269 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1960_v17).length));
          (_1960_v17)[_index269] = _1961_v18;
        }
        if (_this.f11) {
          let _1962_v19;
          _1962_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1950_v8,new _dafny.CodePoint('g'.codePointAt(0)));
          _1962_v19 = (_1962_v19).update(_1950_v8, new _dafny.CodePoint('e'.codePointAt(0)));
          let _1963_v20;
          let _nw293 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _1963_v20 = _nw293;
          let _1964_v21;
          _1964_v21 = _dafny.Seq.of(_this.f22, _this.f22);
          let _1965_v22;
          let _nw294 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1965_v22 = _nw294;
          let _1966_v23;
          _1966_v23 = _dafny.MultiSet.fromElements(_1965_v22);
          let _1967_v24;
          _1967_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_dafny.ZERO).minus(new BigNumber((_1966_v23).cardinality())));
          let _1968_v25;
          _1968_v25 = _dafny.Set.fromElements((_this).f20);
          let _1969_v26;
          let _nw295 = Array((new BigNumber(22)).toNumber());
          _nw295[(_dafny.ZERO).toNumber()] = _module.__default.fm2(globalState);
          _nw295[(_dafny.ONE).toNumber()] = new BigNumber((_1944_v3).length);
          _nw295[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(3)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(4)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(5)).toNumber()] = new BigNumber((_1964_v21).length);
          _nw295[(new BigNumber(6)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(7)).toNumber()] = (_this).f13;
          _nw295[(new BigNumber(8)).toNumber()] = (_this).f13;
          _nw295[(new BigNumber(9)).toNumber()] = _this.f22;
          _nw295[(new BigNumber(10)).toNumber()] = new BigNumber(992);
          _nw295[(new BigNumber(11)).toNumber()] = _1949_i0;
          _nw295[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-454)), function (_1970_i2) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length);
          _nw295[(new BigNumber(13)).toNumber()] = _1949_i0;
          _nw295[(new BigNumber(14)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(15)).toNumber()] = (_this).f13;
          _nw295[(new BigNumber(16)).toNumber()] = _this.f22;
          _nw295[(new BigNumber(17)).toNumber()] = new BigNumber((_1967_v24).length);
          _nw295[(new BigNumber(18)).toNumber()] = new BigNumber((_1968_v25).length);
          _nw295[(new BigNumber(19)).toNumber()] = new BigNumber(((_this).f20).length);
          _nw295[(new BigNumber(20)).toNumber()] = (_this).f12;
          _nw295[(new BigNumber(21)).toNumber()] = new BigNumber((_1942_v1).length);
          _1969_v26 = _nw295;
          let _1971_v27;
          _1971_v27 = _dafny.Seq.of(_1969_v26);
          let _1972_v28;
          _1972_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_1971_v27);
          let _index270 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_1963_v20).length));
          (_1963_v20)[_index270] = (((_1972_v28).contains((_this).f13)) ? ((_1972_v28).get((_this).f13)) : (_dafny.Seq.update(_1971_v27, _module.__default.safeIndex(new BigNumber(-133), new BigNumber((_1971_v27).length)), _1965_v22)));
          let _index271 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_1963_v20).length));
          (_1963_v20)[_index271] = _dafny.Seq.Concat(_1971_v27, _1971_v27);
          let _1973_v29;
          _1973_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1944_v3,!(_this.f11));
          let _1974_v31;
          _1974_v31 = _module.D5.create_DC15(_1949_i0, _1949_i0, (_this).f13);
          let _1975_v32;
          _1975_v32 = _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1974_v31));
          let _1976_v33;
          _1976_v33 = _module.D6.create_DC20(_1975_v32);
          let _1977_v34;
          _1977_v34 = _module.D6.create_DC20(_1975_v32);
          let _1978_v35;
          _1978_v35 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f12), _1944_v3);
          let _1979_v36;
          _1979_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1977_v34,_1978_v35);
          _1973_v29 = function () {
            let _coll81 = new _dafny.Map();
            for (const _compr_81 of ((((_1979_v36).contains(_1977_v34)) ? ((_1979_v36).get(_1977_v34)) : (_dafny.Seq.of(_1944_v3)))).Elements) {
              let _1980_v30 = _compr_81;
              if (_dafny.Seq.contains((((_1979_v36).contains(_1977_v34)) ? ((_1979_v36).get(_1977_v34)) : (_dafny.Seq.of(_1944_v3))), _1980_v30)) {
                _coll81.push([_1980_v30,_dafny.Seq.contains(_1964_v21, _1949_i0)]);
              }
            }
            return _coll81;
          }();
          _1948_v7 = (_1948_v7).Intersect(function () {
            let _coll82 = new _dafny.Set();
            for (const _compr_82 of (_1967_v24).Keys.Elements) {
              let _1981_v37 = _compr_82;
              if ((_1967_v24).contains(_1981_v37)) {
                _coll82.add((_1981_v37).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(201)), function (_1982_i3) {
                  return new _dafny.CodePoint('u'.codePointAt(0));
                })).length))));
              }
            }
            return _coll82;
          }());
          let _1983_v38;
          let _nw296 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1983_v38 = _nw296;
          _1983_v38 = _1983_v38;
        } else {
          let _1984_v39;
          _1984_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_this.f11);
          (_this).f22 = new BigNumber((_1984_v39).length);
          let _index272 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_1941_v0).length));
          (_1941_v0)[_index272] = _this.f11;
          let _index273 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_1941_v0).length));
          (_1941_v0)[_index273] = ((_this.f11) ? (true) : (((_this.f11) ? (_this.f11) : (_this.f11))));
          let _1985_v40;
          let _nw297 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1985_v40 = _nw297;
          let _1986_v41;
          _1986_v41 = _dafny.Seq.of(_1985_v40);
          let _1987_v42;
          let _nw298 = Array((new BigNumber(9)).toNumber());
          _nw298[(_dafny.ZERO).toNumber()] = _1985_v40;
          _nw298[(_dafny.ONE).toNumber()] = (_1986_v41)[_module.__default.safeIndex(_this.f22, new BigNumber((_1986_v41).length))];
          _nw298[(new BigNumber(2)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(3)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(4)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(5)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(6)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(7)).toNumber()] = _1985_v40;
          _nw298[(new BigNumber(8)).toNumber()] = _1985_v40;
          _1987_v42 = _nw298;
          let _1988_v43;
          _1988_v43 = _module.D9.create_DC27(_this.f11, _module.D7.create_DC22((_dafny.ZERO).minus((_module.D5.create_DC15(new BigNumber((_1948_v7).length), _1949_i0, (_this).f12)).dtor_cf27), _1987_v42), (_this.f22).multipliedBy((_this).f13), false);
          _1988_v43 = _1988_v43;
          (globalState).f8 = (_1941_v0)[_module.__default.safeIndex(new BigNumber(64), new BigNumber((_1941_v0).length))];
          let _1989_v44;
          let _1990_v45;
          let _out28;
          let _out29;
          let _outcollector13 = _module.__default.m0(globalState);
          _out28 = _outcollector13[0];
          _out29 = _outcollector13[1];
          _1989_v44 = _out28;
          _1990_v45 = _out29;
        }
        if ((_module.__default.fm2(globalState)).isLessThan((new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1991_i0) => function (_1992_i4) {
          return _1991_i0;
        })(_1949_i0)))).cardinality())).multipliedBy((_this).f13))) {
          let _1993_v46;
          let _nw299 = Array((new BigNumber(21)).toNumber()).fill([]);
          _1993_v46 = _nw299;
          _1993_v46 = p2;
          (_this).f22 = (_this).f13;
          let _arr4 = (p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))];
          let _index274 = _module.__default.safeIndex(new BigNumber(863), new BigNumber(((p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))]).length));
          _arr4[_index274] = _this.f11;
          let _arr5 = (p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))];
          let _index275 = _module.__default.safeIndex(new BigNumber(863), new BigNumber(((p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))]).length));
          _arr5[_index275] = !(true) || (true);
          (_this).f22 = _this.f22;
          let _1994_v47;
          _1994_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1941_v0,(new BigNumber((p0).length)).multipliedBy(_this.f22));
          _1994_v47 = ((_1994_v47).Merge(_1994_v47)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1941_v0,new BigNumber(940)));
        } else {
          (globalState).f8 = _this.f11;
          (_this).f22 = (_this).f12;
          let _1995_v48;
          let _nw300 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1995_v48 = _nw300;
          let _index276 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1995_v48).length));
          (_1995_v48)[_index276] = p0;
          let _1996_v49;
          let _nw301 = new _module.C1();
          _nw301.__ctor((_this).f12);
          _1996_v49 = _nw301;
          let _1997_v50;
          _1997_v50 = _dafny.Seq.of(_1996_v49, _1996_v49, _1996_v49);
          let _index277 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length));
          let _index278 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1995_v48).length));
          let _rhs279 = (p2)[_module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length))];
          let _rhs280 = (_this).f20;
          let _rhs281 = _dafny.Seq.of(_1996_v49, _1996_v49, _1996_v49);
          let _rhs282 = new BigNumber(-688);
          let _lhs160 = p2;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((p2).length));
          let _lhs162 = _1995_v48;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1995_v48).length));
          let _lhs164 = _1996_v49;
          _lhs160[_lhs161] = _rhs279;
          _lhs162[_lhs163] = _rhs280;
          _1997_v50 = _rhs281;
          _lhs164.f27 = _rhs282;
          (globalState).f8 = (_this.f11) === (_this.f11);
          let _rhs283 = (_this.f11) || (_this.f11);
          let _rhs284 = (_dafny.ZERO).minus((_this).f12);
          let _lhs165 = _1996_v49;
          r1 = _rhs283;
          _lhs165.f27 = _rhs284;
        }
      }
      let _hi8 = new BigNumber(157);
      for (let _1998_i5 = (_this).f12; _1998_i5.isLessThan(_hi8); _1998_i5 = _1998_i5.plus(_dafny.ONE)) {
        let _1999_v51;
        let _2000_v52;
        let _out30;
        let _out31;
        let _outcollector14 = _module.__default.m0(globalState);
        _out30 = _outcollector14[0];
        _out31 = _outcollector14[1];
        _1999_v51 = _out30;
        _2000_v52 = _out31;
        (_this).f22 = new BigNumber(-302);
        let _2001_v53;
        _2001_v53 = _dafny.Seq.of((_this).f12, _1998_i5);
        let _2002_v54;
        _2002_v54 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2003_v55;
        let _nw302 = Array((new BigNumber(13)).toNumber());
        _nw302[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("omgjypkmn");
        _nw302[(_dafny.ONE).toNumber()] = (_this).f20;
        _nw302[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("kpqqegm");
        _nw302[(new BigNumber(3)).toNumber()] = _module.__default.fm36((_this).f12, new BigNumber((_2001_v53).length), (_module.D9.create_DC26(_2002_v54)).dtor_cf48, globalState);
        _nw302[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), ((_2004_v54) => function (_2005_i6) {
          return _2004_v54;
        })(_2002_v54));
        _nw302[(new BigNumber(5)).toNumber()] = (_this).f20;
        _nw302[(new BigNumber(6)).toNumber()] = p0;
        _nw302[(new BigNumber(7)).toNumber()] = (_this).f20;
        _nw302[(new BigNumber(8)).toNumber()] = (_this).f20;
        _nw302[(new BigNumber(9)).toNumber()] = (_this).f20;
        _nw302[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), ((_2006_v54) => function (_2007_i7) {
          return _2006_v54;
        })(_2002_v54));
        _nw302[(new BigNumber(11)).toNumber()] = p0;
        _nw302[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("rnwrb");
        _2003_v55 = _nw302;
        let _2008_v56;
        _2008_v56 = _module.D5.create_DC14((_module.D5.create_DC14(_2003_v55)).dtor_cf25);
        let _2009_v57;
        _2009_v57 = _dafny.MultiSet.fromElements(_1998_i5);
        let _2010_v58;
        _2010_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_2000_v52),_2009_v57);
        let _2011_v60;
        _2011_v60 = _dafny.MultiSet.fromElements(false, !((_this).fm6(_this.f11, _1998_i5, _1998_i5, globalState)));
        let _2012_v61;
        _2012_v61 = _dafny.Set.fromElements(_2011_v60);
        let _2013_v62;
        _2013_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2008_v56,(_2010_v58).equals(function () {
          let _coll83 = new _dafny.Map();
          for (const _compr_83 of (_2012_v61).Elements) {
            let _2014_v59 = _compr_83;
            if ((_2012_v61).contains(_2014_v59)) {
              _coll83.push([_2014_v59,_2009_v57]);
            }
          }
          return _coll83;
        }()));
        _2013_v62 = (_2013_v62).Merge(_2013_v62);
        _2000_v52 = ((_this.f11) ? (_2000_v52) : (_this.f11));
      }
      r0 = _this.f11;
      let _2015_v63;
      _2015_v63 = _dafny.Seq.of(_this.f22, (_this).f12, new BigNumber(174), (_dafny.ZERO).minus((_this).f13));
      let _2016_v64;
      _2016_v64 = _module.D4.create_DC12(_this.f11, _module.__default.fm2(globalState), new BigNumber((_2015_v63).length));
      r1 = (_2016_v64).dtor_cf21;
      return [r0, r1];
    }
    m13(p0, p1, p2, globalState) {
      let _this = this;
      let _2017_v0;
      _2017_v0 = _dafny.Seq.of(p1);
      let _index279 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
      (p2)[_index279] = new BigNumber((_dafny.Seq.Concat(_2017_v0, _dafny.Seq.of(_this.f11, _this.f11))).length);
      let _index280 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
      (p2)[_index280] = _module.__default.fm2(globalState);
      let _2018_v1;
      _2018_v1 = _dafny.Set.fromElements(new BigNumber(313));
      let _2019_v2;
      _2019_v2 = _dafny.Seq.of(_2018_v1, _2018_v1, _2018_v1, _2018_v1, _2018_v1);
      let _2020_v3;
      _2020_v3 = (_2019_v2)[_module.__default.safeIndex(new BigNumber(-741), new BigNumber((_2019_v2).length))];
      let _source30 = _2020_v3;
      let _2021___mcc_h0 = _source30;
      let _2022_cf54 = _2021___mcc_h0;
      let _2023_v4;
      _2023_v4 = _dafny.MultiSet.fromElements(p1);
      let _2024_v5;
      _2024_v5 = _module.D7.create_DC21(_2023_v4);
      let _2025_v6;
      _2025_v6 = _module.D7.create_DC24(_2024_v5);
      let _pat_let_tv70 = _2024_v5;
      _2025_v6 = function (_pat_let46_0) {
        return function (_2026_dt__update__tmp_h0) {
          return function (_pat_let47_0) {
            return function (_2027_dt__update_hcf46_h0) {
              return _module.D7.create_DC24(_2027_dt__update_hcf46_h0);
            }(_pat_let47_0);
          }(_pat_let_tv70);
        }(_pat_let46_0);
      }(_2025_v6);
      if (false) {
        (globalState).f8 = (_this.f11) === (_this.f11);
        let _2028_v7;
        let _nw303 = new _module.C0();
        _nw303.__ctor((_this).f13, p1, (_this).f13, _this.f22);
        _2028_v7 = _nw303;
        let _2029_v8;
        _2029_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2028_v7);
        let _2030_v9;
        _2030_v9 = _dafny.Set.fromElements(_2029_v8);
        let _2031_v10;
        _2031_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,_2030_v9);
        let _2032_v11;
        _2032_v11 = _dafny.MultiSet.fromElements(_this.f22);
        let _2033_v12;
        _2033_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2032_v11);
        let _2034_v13;
        let _nw304 = Array((new BigNumber(13)).toNumber());
        _nw304[(_dafny.ZERO).toNumber()] = true;
        _nw304[(_dafny.ONE).toNumber()] = _this.f11;
        _nw304[(new BigNumber(2)).toNumber()] = ((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]).isEqualTo(_this.f22);
        _nw304[(new BigNumber(3)).toNumber()] = !((((_2031_v10).contains(_2028_v7.f29)) ? ((_2031_v10).get(_2028_v7.f29)) : (_2030_v9))).contains(_2029_v8);
        _nw304[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.fromElements((_2028_v7).f28)).IsSubsetOf((((_2033_v12).contains(_this.f11)) ? ((_2033_v12).get(_this.f11)) : (_2032_v11)));
        _nw304[(new BigNumber(5)).toNumber()] = _this.f11;
        _nw304[(new BigNumber(6)).toNumber()] = p1;
        _nw304[(new BigNumber(7)).toNumber()] = !(p1) || (!(p1));
        _nw304[(new BigNumber(8)).toNumber()] = (_this.f22).isLessThanOrEqualTo((_this).f12);
        _nw304[(new BigNumber(9)).toNumber()] = _dafny.Seq.contains((_this).f20, new _dafny.CodePoint('d'.codePointAt(0)));
        _nw304[(new BigNumber(10)).toNumber()] = !((_this).f12).isEqualTo((_this).f12);
        _nw304[(new BigNumber(11)).toNumber()] = ((_this).f13).isLessThanOrEqualTo((_this).f12);
        _nw304[(new BigNumber(12)).toNumber()] = false;
        _2034_v13 = _nw304;
        let _index281 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_2034_v13).length));
        (_2034_v13)[_index281] = false;
        let _index282 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_2034_v13).length));
        (_2034_v13)[_index282] = !(p1);
        let _2035_v14;
        _2035_v14 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains((_this).f20, new _dafny.CodePoint('w'.codePointAt(0))),new BigNumber(35));
        let _2036_v15;
        _2036_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2028_v7.f29,p1);
        _2035_v14 = (_2035_v14).update((((_2036_v15).contains(true)) ? ((_2036_v15).get(true)) : (_2028_v7.f29)), ((_2028_v7.f29) ? ((_this).f13) : (_this.f22)));
        let _2037_v16;
        _2037_v16 = _dafny.Seq.UnicodeFromString("pybpnsivf");
        _2037_v16 = (_this).f20;
        let _index283 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        (p2)[_index283] = _module.__default.fm2(globalState);
      } else {
        let _index284 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        (p2)[_index284] = _module.__default.fm2(globalState);
        let _2038_v17;
        let _nw305 = new _module.C8();
        _nw305.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(305)), function (_2039_i0) {
          return (_this).f12;
        }));
        _2038_v17 = _nw305;
        (globalState).f8 = !(p1);
        let _2040_v18;
        _2040_v18 = _dafny.Seq.UnicodeFromString("jlq");
        let _2041_v19;
        _2041_v19 = _dafny.MultiSet.fromElements(new BigNumber(457));
        let _2042_v20;
        _2042_v20 = _dafny.MultiSet.fromElements((((_2041_v19).contains(_this.f22)) ? ((_2041_v19).get(_this.f22)) : (new BigNumber(737))));
        let _2043_v21;
        _2043_v21 = _module.D11.create_DC31(_this.f11);
        let _2044_v22;
        _2044_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2043_v21,(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]);
        let _index285 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        let _index286 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        let _rhs285 = (p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))];
        let _rhs286 = (p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))];
        let _rhs287 = _dafny.Seq.UnicodeFromString("uk");
        let _rhs288 = _this.f11;
        let _rhs289 = _module.__default.safeDivisionInt((((_2042_v20).contains(new BigNumber((_2044_v22).length))) ? ((_2042_v20).get(new BigNumber((_2044_v22).length))) : (new BigNumber((_2040_v18).length))), ((_this).f13).minus(p0));
        let _lhs166 = p2;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        let _lhs168 = _this;
        let _lhs169 = _this;
        let _lhs170 = p2;
        let _lhs171 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        _lhs166[_lhs167] = _rhs285;
        _lhs168.f22 = _rhs286;
        _2040_v18 = _rhs287;
        _lhs169.f11 = _rhs288;
        _lhs170[_lhs171] = _rhs289;
        (_this).f22 = _module.__default.safeModuloInt((_this).f13, _module.__default.safeModuloInt(p0, p0));
      }
      let _2045_v23;
      _2045_v23 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]);
      (globalState).f8 = (_2045_v23).contains(p1);
      let _2046_v24;
      _2046_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2023_v4,p1);
      let _2047_v25;
      _2047_v25 = _dafny.MultiSet.fromElements((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], new BigNumber((_2046_v24).length), (_this).f12, ((p1) ? ((_this).f13) : (new BigNumber(((_this).f20).length))));
      (_this).f22 = (((_2047_v25).contains(_module.__default.fm2(globalState))) ? ((_2047_v25).get(_module.__default.fm2(globalState))) : (_module.__default.safeDivisionInt(new BigNumber(38), p0)));
      let _hi9 = (_this).f13;
      for (let _2048_i1 = (_this).f13; _2048_i1.isLessThan(_hi9); _2048_i1 = _2048_i1.plus(_dafny.ONE)) {
        let _2049_v26;
        _2049_v26 = _dafny.MultiSet.fromElements((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], (p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]);
        let _2050_v27;
        _2050_v27 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2051_v28;
        _2051_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2050_v27,_this.f11);
        let _2052_v29;
        _2052_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2048_i1);
        let _2053_v30;
        _2053_v30 = _dafny.Seq.of((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], (_this).f12);
        let _2054_v31;
        _2054_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_2055_p2) => function (_2056_i2) {
          return (_2055_p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((_2055_p2).length))];
        })(p2)));
        let _2057_v32;
        _2057_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f12);
        let _2058_v33;
        _2058_v33 = _dafny.MultiSet.fromElements(_2057_v32, _2057_v32);
        let _2059_v34;
        _2059_v34 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))],p1);
        let _2060_v35;
        _2060_v35 = _dafny.Set.fromElements(p1);
        let _2061_v36;
        let _nw306 = Array((new BigNumber(23)).toNumber());
        _nw306[(_dafny.ZERO).toNumber()] = _this.f11;
        _nw306[(_dafny.ONE).toNumber()] = false;
        _nw306[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], (_this).f12)).IsProperSubsetOf(_2049_v26);
        _nw306[(new BigNumber(3)).toNumber()] = _this.f11;
        _nw306[(new BigNumber(4)).toNumber()] = _this.f11;
        _nw306[(new BigNumber(5)).toNumber()] = (((_2051_v28).contains(_2050_v27)) ? ((_2051_v28).get(_2050_v27)) : (_this.f11));
        _nw306[(new BigNumber(6)).toNumber()] = ((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]).isLessThan(new BigNumber((_2052_v29).length));
        _nw306[(new BigNumber(7)).toNumber()] = !_dafny.areEqual(_2053_v30, _2053_v30);
        _nw306[(new BigNumber(8)).toNumber()] = true;
        _nw306[(new BigNumber(9)).toNumber()] = !_dafny.Seq.contains((((_2054_v31).contains(_this.f11)) ? ((_2054_v31).get(_this.f11)) : (_2053_v30)), new BigNumber((_2017_v0).length));
        _nw306[(new BigNumber(10)).toNumber()] = p1;
        _nw306[(new BigNumber(11)).toNumber()] = _this.f11;
        _nw306[(new BigNumber(12)).toNumber()] = !(p1) || (_this.f11);
        _nw306[(new BigNumber(13)).toNumber()] = _this.f11;
        _nw306[(new BigNumber(14)).toNumber()] = p1;
        _nw306[(new BigNumber(15)).toNumber()] = true;
        _nw306[(new BigNumber(16)).toNumber()] = _module.__default.fm1(p0, _this.f11, (_this).f13, globalState);
        _nw306[(new BigNumber(17)).toNumber()] = !(p1);
        _nw306[(new BigNumber(18)).toNumber()] = (_2058_v33).IsSubsetOf(_2058_v33);
        _nw306[(new BigNumber(19)).toNumber()] = (((_2059_v34).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))])).length))) ? ((_2059_v34).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))])).length))) : (_this.f11));
        _nw306[(new BigNumber(20)).toNumber()] = !(_2060_v35).equals(_2060_v35);
        _nw306[(new BigNumber(21)).toNumber()] = !(!(_2057_v32).equals((_2057_v32).update(new BigNumber(320), new BigNumber(((_this).f20).length))));
        _nw306[(new BigNumber(22)).toNumber()] = p1;
        _2061_v36 = _nw306;
        let _2062_v37;
        let _nw307 = Array((new BigNumber(23)).toNumber()).fill([]);
        _2062_v37 = _nw307;
        let _2063_v38;
        _2063_v38 = _module.D7.create_DC23(_2062_v37, _2017_v0, p1, (_this).f12);
        let _index287 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_2061_v36).length));
        (_2061_v36)[_index287] = !((_2063_v38).dtor_cf44) || (true);
        let _index288 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_2061_v36).length));
        (_2061_v36)[_index288] = p1;
        let _index289 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        (p2)[_index289] = _this.f22;
        _2052_v29 = (_2052_v29).update((_2061_v36)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_2061_v36).length))], ((_dafny.ZERO).minus((_this).f13)).minus((_this).f13));
        let _index290 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        let _rhs290 = (new BigNumber(((_this).f20).length)).plus((new BigNumber(308)).minus(new BigNumber(((_this).fm8((_dafny.ZERO).minus(new BigNumber(((_this).f20).length)), _dafny.Map.Empty.slice().updateUnsafe(_2050_v27,p1), _2018_v1, globalState)).length)));
        let _rhs291 = ((new BigNumber(((_this).f20).length)).multipliedBy((_this).f13)).minus(_module.__default.fm2(globalState));
        let _rhs292 = _this.f11;
        let _lhs172 = p2;
        let _lhs173 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        let _lhs174 = _this;
        let _lhs175 = globalState;
        _lhs172[_lhs173] = _rhs290;
        _lhs174.f22 = _rhs291;
        _lhs175.f8 = _rhs292;
      }
      let _2064_v39;
      let _nw308 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2064_v39 = _nw308;
      let _2065_v40;
      _2065_v40 = _module.D5.create_DC14(_2064_v39);
      let _2066_v41;
      _2066_v41 = _dafny.MultiSet.fromElements(_2065_v40);
      let _2067_v42;
      _2067_v42 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(_2065_v40, _2065_v40), _2066_v41);
      let _2068_v43;
      _2068_v43 = _dafny.Seq.of((_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]));
      let _2069_v44;
      _2069_v44 = _dafny.MultiSet.fromElements((_this).f20);
      let _2070_v45;
      _2070_v45 = _dafny.Set.fromElements(_this.f11);
      let _2071_v46;
      let _nw309 = new _module.C4();
      _nw309.__ctor(new BigNumber(531), p1);
      _2071_v46 = _nw309;
      let _2072_v47;
      _2072_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v46,(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]);
      let _2073_v48;
      let _nw310 = Array((new BigNumber(20)).toNumber());
      _nw310[(_dafny.ZERO).toNumber()] = p0;
      _nw310[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2067_v42).length));
      _nw310[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_this).f13, new BigNumber((_dafny.Seq.UnicodeFromString("ggj")).length));
      _nw310[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm2(globalState));
      _nw310[(new BigNumber(4)).toNumber()] = (new BigNumber((_2068_v43).length)).minus((((_2069_v44).contains((_this).f20)) ? ((_2069_v44).get((_this).f20)) : ((_this).f13)));
      _nw310[(new BigNumber(5)).toNumber()] = _this.f22;
      _nw310[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw310[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_this).f13, new BigNumber((_2070_v45).length));
      _nw310[(new BigNumber(8)).toNumber()] = (_this).f13;
      _nw310[(new BigNumber(9)).toNumber()] = (_this).f13;
      _nw310[(new BigNumber(10)).toNumber()] = (_this).f13;
      _nw310[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))]);
      _nw310[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((_this).f13, p0);
      _nw310[(new BigNumber(13)).toNumber()] = (_this).f13;
      _nw310[(new BigNumber(14)).toNumber()] = (_this).f12;
      _nw310[(new BigNumber(15)).toNumber()] = new BigNumber(((_this).f20).length);
      _nw310[(new BigNumber(16)).toNumber()] = _this.f22;
      _nw310[(new BigNumber(17)).toNumber()] = p0;
      _nw310[(new BigNumber(18)).toNumber()] = (((_2072_v47).contains(_2071_v46)) ? ((_2072_v47).get(_2071_v46)) : (new BigNumber((_dafny.Seq.UnicodeFromString("uidytrhx")).length)));
      _nw310[(new BigNumber(19)).toNumber()] = p0;
      _2073_v48 = _nw310;
      _2073_v48 = _2073_v48;
      if (_this.f11) {
        let _index291 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length));
        (p2)[_index291] = _this.f22;
        let _2074_v49;
        _2074_v49 = new _dafny.CodePoint('i'.codePointAt(0));
        _2074_v49 = _2074_v49;
        (globalState).f8 = !(new BigNumber((_module.__default.fm36(_this.f22, (_this).f13, _2074_v49, globalState)).length)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality())));
        let _2075_v50;
        _2075_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f20).length),_2071_v46.f11);
        _2075_v50 = (_2075_v50).update(_module.__default.safeDivisionInt(_module.__default.fm2(globalState), (_this).f13), ((_2071_v46.f11) ? (_2071_v46.f11) : (!(_this.f11))));
        let _2076_v51;
        _2076_v51 = _module.D5.create_DC15((_this).f12, (_this).f12, p0);
        let _2077_v52;
        _2077_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2076_v51);
        let _2078_v53;
        _2078_v53 = _module.D6.create_DC16(_2077_v52);
        let _2079_v54;
        _2079_v54 = _dafny.Set.fromElements(_module.D6.create_DC16(_2077_v52), _2078_v53);
        let _2080_v55;
        _2080_v55 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_2071_v46.f11)),new BigNumber((_2079_v54).length));
        let _2081_v56;
        _2081_v56 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.update(_2068_v43, _module.__default.safeIndex(new BigNumber((_2017_v0).length), new BigNumber((_2068_v43).length)), (_this).f12)).length), _dafny.Seq.UnicodeFromString("c"), (_this).f12, _2071_v46.f11);
        let _2082_v57;
        let _nw311 = Array((new BigNumber(21)).toNumber());
        _nw311[(_dafny.ZERO).toNumber()] = _2071_v46.f11;
        _nw311[(_dafny.ONE).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(2)).toNumber()] = p1;
        _nw311[(new BigNumber(3)).toNumber()] = ((_this).f13).isEqualTo((_this).f13);
        _nw311[(new BigNumber(4)).toNumber()] = _this.f11;
        _nw311[(new BigNumber(5)).toNumber()] = true;
        _nw311[(new BigNumber(6)).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(7)).toNumber()] = p1;
        _nw311[(new BigNumber(8)).toNumber()] = (_this).fm7(globalState);
        _nw311[(new BigNumber(9)).toNumber()] = !(_2080_v55).equals(_dafny.Map.Empty.slice().updateUnsafe(_2071_v46.f11,_this.f22));
        _nw311[(new BigNumber(10)).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(11)).toNumber()] = false;
        _nw311[(new BigNumber(12)).toNumber()] = !(_module.__default.fm1(_this.f22, _this.f11, new BigNumber(906), globalState));
        _nw311[(new BigNumber(13)).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(14)).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(15)).toNumber()] = _this.f11;
        _nw311[(new BigNumber(16)).toNumber()] = _2071_v46.f11;
        _nw311[(new BigNumber(17)).toNumber()] = !((function (_pat_let48_0) {
          return function (_2083_dt__update__tmp_h1) {
            return function (_pat_let49_0) {
              return function (_2084_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_2083_dt__update__tmp_h1).dtor_cf1, _2084_dt__update_hcf2_h0, (_2083_dt__update__tmp_h1).dtor_cf3, (_2083_dt__update__tmp_h1).dtor_cf4);
              }(_pat_let49_0);
            }((_this).f20);
          }(_pat_let48_0);
        }(_2081_v56)).dtor_cf4);
        _nw311[(new BigNumber(18)).toNumber()] = (_2070_v45).equals(_2070_v45);
        _nw311[(new BigNumber(19)).toNumber()] = !(!((((_2075_v50).contains(p0)) ? ((_2075_v50).get(p0)) : (_2071_v46.f11))));
        _nw311[(new BigNumber(20)).toNumber()] = p1;
        _2082_v57 = _nw311;
        let _index292 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2082_v57).length));
        (_2082_v57)[_index292] = !(_this.f11);
        let _index293 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2082_v57).length));
        (_2082_v57)[_index293] = false;
      } else {
        let _2085_v58;
        _2085_v58 = _dafny.Seq.of((_this).f20);
        _2085_v58 = _2085_v58;
        let _2086_v59;
        _2086_v59 = _module.D7.create_DC21(_dafny.MultiSet.fromElements(_this.f11, true));
        let _2087_v60;
        _2087_v60 = _dafny.MultiSet.fromElements(_2071_v46.f11, _2071_v46.f11, p1, _this.f11);
        let _pat_let_tv71 = _2087_v60;
        let _2088_v61;
        _2088_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2071_v46.f11),function (_pat_let50_0) {
          return function (_2089_dt__update__tmp_h2) {
            return function (_pat_let51_0) {
              return function (_2090_dt__update_hcf39_h0) {
                return _module.D7.create_DC21(_2090_dt__update_hcf39_h0);
              }(_pat_let51_0);
            }(_pat_let_tv71);
          }(_pat_let50_0);
        }(_2086_v59));
        let _2091_v62;
        _2091_v62 = new _dafny.CodePoint('v'.codePointAt(0));
        let _2092_v63;
        _2092_v63 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), _2091_v62, _2091_v62),(_this).f13);
        let _2093_v64;
        _2093_v64 = _dafny.MultiSet.fromElements(_2091_v62);
        let _2094_v65;
        let _nw312 = Array((new BigNumber(2)).toNumber());
        _nw312[(_dafny.ZERO).toNumber()] = _2092_v63;
        _nw312[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2093_v64,(p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))])).Merge(_2092_v63);
        _2094_v65 = _nw312;
        let _index294 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_2094_v65).length));
        (_2094_v65)[_index294] = _2092_v63;
        let _2095_v66;
        _2095_v66 = _module.D15.create_DC40(_module.__default.fm55(new BigNumber(831), (_this).f13, p0, globalState));
        let _2096_v67;
        _2096_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f20);
        let _2097_v68;
        _2097_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber(398));
        let _2098_v69;
        _2098_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this.f22).plus(new BigNumber((_2096_v67).length)),(((_2097_v68).contains((_this).f12)) ? ((_2097_v68).get((_this).f12)) : ((_this).f12)));
        let _2099_v70;
        _2099_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v46.f11,_2071_v46.f11);
        let _index295 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_2094_v65).length));
        let _rhs293 = (((_2071_v46.f11) ? (_module.D15.create_DC40(_dafny.Map.Empty.slice().updateUnsafe(_2017_v0,_2086_v59))) : (_2095_v66))).dtor_cf64;
        let _rhs294 = _2092_v63;
        let _rhs295 = _2091_v62;
        let _rhs296 = (((_2098_v69).contains(new BigNumber((_2099_v70).length))) ? ((_2098_v69).get(new BigNumber((_2099_v70).length))) : (new BigNumber(((_this).f20).length)));
        let _rhs297 = !(((_2071_v46.f11) ? (_2099_v70) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2071_v46.f11)))).contains(_2071_v46.f11);
        let _lhs176 = _2094_v65;
        let _lhs177 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_2094_v65).length));
        let _lhs178 = _this;
        let _lhs179 = _2071_v46;
        _2088_v61 = _rhs293;
        _lhs176[_lhs177] = _rhs294;
        _2091_v62 = _rhs295;
        _lhs178.f22 = _rhs296;
        _lhs179.f11 = _rhs297;
        let _2100_v71;
        _2100_v71 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(p0, (_this).f13), (_this).f13);
        _2100_v71 = ((_dafny.MultiSet.fromElements((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))])).Difference(_2100_v71)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(525), (_this).f13, (p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], (_dafny.ZERO).minus((_this).f12))));
        let _2101_v72;
        let _nw313 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _2101_v72 = _nw313;
        let _2102_v73;
        _2102_v73 = _dafny.Seq.of(_dafny.Seq.of(_this.f22));
        let _index296 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2101_v72).length));
        (_2101_v72)[_index296] = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_2102_v73)[_module.__default.safeIndex((p2)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((p2).length))], new BigNumber((_2102_v73).length))]);
        let _2103_v74;
        _2103_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_dafny.Seq.Concat(_2068_v43, _dafny.Seq.of(p0)));
        let _index297 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2101_v72).length));
        let _rhs298 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,!(_2071_v46.f11) || (_2071_v46.f11))).length);
        let _rhs299 = _2103_v74;
        let _rhs300 = p1;
        let _rhs301 = (p0).multipliedBy(new BigNumber(202));
        let _rhs302 = _2091_v62;
        let _lhs180 = _this;
        let _lhs181 = _2101_v72;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2101_v72).length));
        let _lhs183 = _2071_v46;
        let _lhs184 = _this;
        _lhs180.f22 = _rhs298;
        _lhs181[_lhs182] = _rhs299;
        _lhs183.f11 = _rhs300;
        _lhs184.f22 = _rhs301;
        _2091_v62 = _rhs302;
        let _2104_v75;
        _2104_v75 = _dafny.Seq.UnicodeFromString("sdeejae");
        _2104_v75 = (_this).f20;
      }
      let _2105_v76;
      let _init60 = ((_2106_v0) => function (_2107_i3) {
        return _dafny.Seq.of((_2106_v0)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)))).cardinality()), new BigNumber((_2106_v0).length))], !((_module.D15.create_DC41(_this.f11)).dtor_cf65));
      })(_2017_v0);
      let _nw314 = Array((new BigNumber(25)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw314.length); _i0_60++) {
        _nw314[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _2105_v76 = _nw314;
      _2105_v76 = _2105_v76;
      return;
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(361)), function (_2108_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length), new BigNumber(-38))).length), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ecqjml"), _dafny.Seq.UnicodeFromString("fojk")), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("omya"), _dafny.Seq.UnicodeFromString("qfg"))).length), true);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f11;
    };
    fm16(p0, p1, globalState) {
      let _this = this;
      return !(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11)).length)))).contains(!(!(_this.f11)) || ((_module.D0.create_DC3(new BigNumber((_dafny.MultiSet.FromArray((_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), function (_2109_i0) {
  return new BigNumber((_dafny.Seq.of(_this.f11)).length);
})))).cardinality()), _this.f11)).dtor_cf8));
    };
    m9(p0, p1, globalState) {
      let _this = this;
      let _2110_v0;
      let _init61 = function (_2111_i0) {
        return _this.f11;
      };
      let _nw315 = Array((new BigNumber(25)).toNumber());
      for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw315.length); _i0_61++) {
        _nw315[_i0_61] = _init61(new BigNumber(_i0_61));
      }
      _2110_v0 = _nw315;
      let _2112_v1;
      _2112_v1 = _dafny.Set.fromElements(p0, p0);
      let _index298 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length));
      (_2110_v0)[_index298] = (_2112_v1).IsDisjointFrom(_2112_v1);
      let _index299 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length));
      (_2110_v0)[_index299] = _this.f11;
      let _2113_i1;
      _2113_i1 = _dafny.ZERO;
      L8: {
        while ((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2113_i1)) {
              break L8;
            }
            _2113_i1 = (_2113_i1).plus(_dafny.ONE);
            let _index300 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length));
            (_2110_v0)[_index300] = (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))];
            let _2114_v2;
            _2114_v2 = _dafny.MultiSet.fromElements(new BigNumber(232));
            (globalState).f8 = (_this).fm16(p0, (new BigNumber((_2114_v2).cardinality())).plus(p0), globalState);
            let _2115_v3;
            _2115_v3 = _dafny.Seq.of(p0, p0);
            let _2116_v4;
            _2116_v4 = _dafny.Seq.update(_2115_v3, _module.__default.safeIndex(p0, new BigNumber((_2115_v3).length)), p0);
            let _2117_v5;
            _2117_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-638)), function (_2118_i2) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            }),(_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]);
            let _2119_v6;
            _2119_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),new BigNumber((_2117_v5).length));
            let _2120_v7;
            _2120_v7 = _dafny.Seq.UnicodeFromString("ptqylp");
            _2119_v6 = (_2119_v6).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(142)), ((_2121_p0) => function (_2122_i3) {
              return _2121_p0;
            })(p0)), ((_dafny.ZERO).minus(new BigNumber((_2120_v7).length))).multipliedBy(p0));
            let _2123_v8;
            _2123_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2112_v1,new BigNumber((_2114_v2).cardinality()));
            _2123_v8 = (_2123_v8).update(_2112_v1, p0);
          }
        }
      }
      let _2124_v9;
      _2124_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]);
      let _2125_v10;
      _2125_v10 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState), new BigNumber(-777));
      let _2126_v11;
      _2126_v11 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((((_2125_v10).contains(p0)) ? ((_2125_v10).get(p0)) : (p0)))),(_this).fm6((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], p0, p0, globalState)), _2124_v9, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2112_v1).length),_this.f11)).update(p0, false));
      _2124_v9 = (_2126_v11)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, new BigNumber(-618)), new BigNumber((_2126_v11).length))];
      let _2127_v12;
      _2127_v12 = _dafny.Seq.of(_this.f11);
      let _2128_v13;
      _2128_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm6((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], new BigNumber((_2127_v12).length), p0, globalState),new BigNumber((_2112_v1).length));
      if ((_2127_v12)[_module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]),p0)).Merge(_2128_v13)).length), new BigNumber((_2127_v12).length))]) {
        let _2129_v14;
        _2129_v14 = new BigNumber(724);
        let _2130_v15;
        _2130_v15 = new _dafny.CodePoint('n'.codePointAt(0));
        let _2131_v16;
        _2131_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2130_v15,_this.f11);
        let _rhs303 = p0;
        let _rhs304 = !(((!(_this.f11) || ((((_2124_v9).contains(_2129_v14)) ? ((_2124_v9).get(_2129_v14)) : (true)))) ? ((p0).isLessThan(_2129_v14)) : (!(_this.f11) || ((((_2131_v16).contains(_2130_v15)) ? ((_2131_v16).get(_2130_v15)) : (false))))));
        let _rhs305 = _2129_v14;
        let _rhs306 = new BigNumber(182);
        let _rhs307 = (p0).isLessThanOrEqualTo(_2129_v14);
        let _lhs185 = globalState;
        let _lhs186 = _this;
        _2129_v14 = _rhs303;
        _lhs185.f8 = _rhs304;
        _2129_v14 = _rhs305;
        _2129_v14 = _rhs306;
        _lhs186.f11 = _rhs307;
        _2129_v14 = _2129_v14;
        if (_this.f11) {
          _2129_v14 = _2129_v14;
          let _2132_v17;
          _2132_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(p0, p0));
          let _2133_v18;
          _2133_v18 = _dafny.Seq.of(p0);
          let _2134_v19;
          _2134_v19 = (((_2132_v17).contains(p0)) ? ((_2132_v17).get(p0)) : (_2133_v18));
          let _2135_v20;
          let _nw316 = Array((new BigNumber(27)).toNumber());
          _nw316[(_dafny.ZERO).toNumber()] = _2134_v19;
          _nw316[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), ((_2136_p0) => function (_2137_i4) {
            return (_dafny.ZERO).minus(_2136_p0);
          })(p0));
          _nw316[(new BigNumber(2)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(3)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(4)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(5)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(6)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(7)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(8)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), ((_2138_v19, _2139_v0) => function (_2140_i5) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2138_v19,_module.D0.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length), (_2139_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2139_v0).length))]))).length);
          })(_2134_v19, _2110_v0));
          _nw316[(new BigNumber(10)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(11)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(12)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(13)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(14)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(15)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(16)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(p0, _2129_v14, _2129_v14);
          _nw316[(new BigNumber(18)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(19)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(20)).toNumber()] = _2133_v18;
          _nw316[(new BigNumber(21)).toNumber()] = _2133_v18;
          _nw316[(new BigNumber(22)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(23)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(24)).toNumber()] = _2134_v19;
          _nw316[(new BigNumber(25)).toNumber()] = _2133_v18;
          _nw316[(new BigNumber(26)).toNumber()] = _2134_v19;
          _2135_v20 = _nw316;
          let _2141_v21;
          let _2142_v22;
          let _2143_v23;
          let _out32;
          let _out33;
          let _out34;
          let _outcollector15 = (_this).m10(_this.f11, _2135_v20, p0, globalState);
          _out32 = _outcollector15[0];
          _out33 = _outcollector15[1];
          _out34 = _outcollector15[2];
          _2141_v21 = _out32;
          _2142_v22 = _out33;
          _2143_v23 = _out34;
          _2129_v14 = (_dafny.ZERO).minus(_module.__default.fm2(globalState));
          (_this).f11 = !_dafny.Seq.contains(_2127_v12, false);
          (globalState).f8 = !(((_this.f11) || (_this.f11)) || (_this.f11));
        } else {
          let _2144_v24;
          _2144_v24 = _dafny.MultiSet.fromElements(_this.f11, _this.f11, !(!((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))])));
          let _2145_v25;
          _2145_v25 = _dafny.Seq.of(_2144_v24);
          let _rhs308 = (_dafny.ZERO).minus(_2129_v14);
          let _rhs309 = _dafny.Seq.Concat(_2145_v25, _2145_v25);
          _2129_v14 = _rhs308;
          _2145_v25 = _rhs309;
          let _2146_v26;
          _2146_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))],_2130_v15);
          let _2147_v27;
          _2147_v27 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_2146_v26).length)));
          let _2148_v28;
          let _nw317 = new _module.C8();
          _nw317.__ctor(_2147_v27);
          _2148_v28 = _nw317;
          let _2149_v29;
          _2149_v29 = _dafny.Seq.UnicodeFromString("nvdlfmlk");
          let _2150_v30;
          _2150_v30 = _dafny.Seq.of(_2149_v29, _dafny.Seq.Concat(_2149_v29, _dafny.Seq.UnicodeFromString("sslfvi")));
          _2150_v30 = _dafny.Seq.of(_2149_v29);
          let _index301 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((p1).length));
          (p1)[_index301] = _module.__default.fm2(globalState);
          let _index302 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((p1).length));
          (p1)[_index302] = _module.__default.safeDivisionInt((new BigNumber((_2149_v29).length)).multipliedBy(p0), (_dafny.ZERO).minus(_2129_v14));
          (globalState).f5 = (_2148_v28).f23;
        }
        _2129_v14 = p0;
        let _2151_v31;
        let _nw318 = new _module.C6();
        _nw318.__ctor();
        _2151_v31 = _nw318;
      } else {
        let _2152_v32;
        _2152_v32 = _dafny.Seq.of(p1);
        let _2153_v33;
        _2153_v33 = _dafny.Map.Empty.slice().updateUnsafe((_2152_v32)[_module.__default.safeIndex(p0, new BigNumber((_2152_v32).length))],_2110_v0);
        _2153_v33 = (_2153_v33).update((_2152_v32)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_2152_v32).length))], _2110_v0);
        let _2154_v34;
        _2154_v34 = new BigNumber(725);
        let _2155_v35;
        _2155_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2154_v34,p0);
        let _2156_v36;
        _2156_v36 = _module.D11.create_DC31((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]);
        let _2157_v37;
        _2157_v37 = _dafny.Seq.of(_2156_v36);
        _2154_v34 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(((((_2155_v35).contains(new BigNumber((_2157_v37).length))) ? ((_2155_v35).get(new BigNumber((_2157_v37).length))) : ((_dafny.ZERO).minus(p0)))).multipliedBy(_2154_v34), new BigNumber((_2127_v12).length)));
        let _2158_v38;
        _2158_v38 = new _dafny.CodePoint('r'.codePointAt(0));
        _2158_v38 = _2158_v38;
        _2154_v34 = (_dafny.ZERO).minus(_2154_v34);
        let _2159_v39;
        _2159_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
        (globalState).f8 = ((new BigNumber((_2159_v39).length)).minus(new BigNumber(551))).isLessThanOrEqualTo((new BigNumber(-17)).multipliedBy(_2154_v34));
      }
      if ((((((_2124_v9).contains(new BigNumber((_2128_v13).length))) ? ((_2124_v9).get(new BigNumber((_2128_v13).length))) : ((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]))) ? (_dafny.Seq.IsProperPrefixOf(_2127_v12, _2127_v12)) : (_this.f11))) {
        let _2160_v40;
        let _nw319 = new _module.C6();
        _nw319.__ctor();
        _2160_v40 = _nw319;
        let _2161_v41;
        _2161_v41 = _dafny.Seq.UnicodeFromString("hgnjjve");
        let _2162_v42;
        _2162_v42 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm2(globalState)).multipliedBy(new BigNumber((_2161_v41).length)),_module.__default.safeDivisionInt(p0, p0));
        _2162_v42 = (_2162_v42).update((p0).plus((_dafny.ZERO).minus(p0)), (p0).minus(p0));
        let _2163_v43;
        _2163_v43 = _dafny.Set.fromElements(_2162_v42, _2162_v42, _2162_v42);
        _2163_v43 = (_dafny.Set.fromElements(_2162_v42, _2162_v42, _2162_v42)).Union(_2163_v43);
        let _2164_v44;
        let _nw320 = new _module.C4();
        _nw320.__ctor(_module.__default.safeDivisionInt(new BigNumber(457), p0), _module.__default.fm1((_dafny.ZERO).minus(p0), _this.f11, p0, globalState));
        _2164_v44 = _nw320;
        let _2165_v45;
        _2165_v45 = _dafny.Seq.of((_2164_v44).f25);
        let _2166_v46;
        _2166_v46 = _2165_v45;
        let _2167_v47;
        _2167_v47 = _dafny.Map.Empty.slice().updateUnsafe((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))],_2166_v46);
        _2167_v47 = (_2167_v47).update(_this.f11, _2166_v46);
      } else {
        let _index303 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length));
        (_2110_v0)[_index303] = !(_this.f11);
        let _2168_v48;
        _2168_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v9,_2127_v12);
        let _2169_v50;
        _2169_v50 = _dafny.Map.Empty.slice().updateUnsafe(_2110_v0,(_2168_v48).update(function () {
          let _coll84 = new _dafny.Map();
          for (const _compr_84 of _dafny.IntegerRange(new BigNumber(-830), new BigNumber(911))) {
            let _2170_v49 = _compr_84;
            if (((new BigNumber(-830)).isLessThanOrEqualTo(_2170_v49)) && ((_2170_v49).isLessThan(new BigNumber(911)))) {
              _coll84.push([_module.__default.safeDivisionInt(_2170_v49, p0),(_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]]);
            }
          }
          return _coll84;
        }(), _dafny.Seq.of((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))])));
        let _2171_v53;
        _2171_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v9,p0);
        let _2172_v55;
        _2172_v55 = _dafny.Seq.of(_2168_v48, _2168_v48, _2168_v48, function () {
          let _coll85 = new _dafny.Map();
          for (const _compr_85 of (_module.__default.fm56(p0, (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], _this.f11, (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], globalState)).Elements) {
            let _2173_v54 = _compr_85;
            if (_dafny.Seq.contains(_module.__default.fm56(p0, (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], _this.f11, (_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))], globalState), _2173_v54)) {
              _coll85.push([_2173_v54,_2127_v12]);
            }
          }
          return _coll85;
        }());
        let _2174_v56;
        let _nw321 = Array((new BigNumber(25)).toNumber());
        _nw321[(_dafny.ZERO).toNumber()] = _2168_v48;
        _nw321[(_dafny.ONE).toNumber()] = (_2168_v48).Merge((((_2169_v50).contains(_2110_v0)) ? ((_2169_v50).get(_2110_v0)) : (_dafny.Map.Empty.slice().updateUnsafe(_2124_v9,_dafny.Seq.of(_this.f11)))));
        _nw321[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll86 = new _dafny.Map();
          for (const _compr_86 of (_2125_v10).Elements) {
            let _2175_v51 = _compr_86;
            if ((_2125_v10).contains(_2175_v51)) {
              _coll86.push([(_2175_v51).minus(p0),_this.f11]);
            }
          }
          return _coll86;
        }(),_dafny.Seq.of((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]));
        _nw321[(new BigNumber(3)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(4)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(5)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2124_v9,_2127_v12);
        _nw321[(new BigNumber(7)).toNumber()] = (function () {
          let _coll87 = new _dafny.Map();
          for (const _compr_87 of (_2171_v53).Keys.Elements) {
            let _2176_v52 = _compr_87;
            if ((_2171_v53).contains(_2176_v52)) {
              _coll87.push([_2176_v52,_2127_v12]);
            }
          }
          return _coll87;
        }()).Merge(_2168_v48);
        _nw321[(new BigNumber(8)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(9)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(10)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2124_v9,_2127_v12);
        _nw321[(new BigNumber(12)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_2124_v9,_2127_v12))).Merge(_2168_v48);
        _nw321[(new BigNumber(13)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(14)).toNumber()] = (_2172_v55)[_module.__default.safeIndex(p0, new BigNumber((_2172_v55).length))];
        _nw321[(new BigNumber(15)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(16)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),(_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]),_2127_v12);
        _nw321[(new BigNumber(18)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(19)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(20)).toNumber()] = (((_2110_v0)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length))]) ? (_2168_v48) : (_2168_v48));
        _nw321[(new BigNumber(21)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(22)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(23)).toNumber()] = _2168_v48;
        _nw321[(new BigNumber(24)).toNumber()] = _2168_v48;
        _2174_v56 = _nw321;
        let _index304 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2174_v56).length));
        (_2174_v56)[_index304] = _module.__default.fm57(new BigNumber(-671), _this.f11, globalState);
        let _2177_v57;
        _2177_v57 = new BigNumber(-177);
        let _2178_v58;
        _2178_v58 = _dafny.Seq.UnicodeFromString("ehm");
        let _index305 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2174_v56).length));
        let _rhs310 = _2168_v48;
        let _rhs311 = _dafny.Seq.IsProperPrefixOf(_2178_v58, _2178_v58);
        let _rhs312 = (p0).plus(_2177_v57);
        let _rhs313 = _this.f11;
        let _lhs187 = _2174_v56;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2174_v56).length));
        let _lhs189 = _this;
        let _lhs190 = globalState;
        _lhs187[_lhs188] = _rhs310;
        _lhs189.f11 = _rhs311;
        _2177_v57 = _rhs312;
        _lhs190.f8 = _rhs313;
        (globalState).f8 = true;
        _2177_v57 = p0;
        let _index306 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((p1).length));
        (p1)[_index306] = p0;
        let _index307 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((p1).length));
        (p1)[_index307] = (_dafny.ZERO).minus(_2177_v57);
      }
      let _index308 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2110_v0).length));
      (_2110_v0)[_index308] = _this.f11;
      return;
    }
    m10(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let _2179_i0;
      _2179_i0 = _dafny.ZERO;
      L9: {
        while (_this.f11) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2179_i0)) {
              break L9;
            }
            _2179_i0 = (_2179_i0).plus(_dafny.ONE);
            let _2180_v0;
            let _nw322 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2180_v0 = _nw322;
            let _2181_v1;
            _2181_v1 = _dafny.Seq.of(_this.f11);
            let _rhs314 = (!(p2).isEqualTo(p2)) || (p0);
            let _rhs315 = _2180_v0;
            let _rhs316 = _dafny.Seq.Concat(_2181_v1, _2181_v1);
            let _rhs317 = _this.f11;
            let _rhs318 = _this.f11;
            let _lhs191 = _this;
            let _lhs192 = globalState;
            let _lhs193 = globalState;
            _lhs191.f11 = _rhs314;
            _2180_v0 = _rhs315;
            _2181_v1 = _rhs316;
            _lhs192.f8 = _rhs317;
            _lhs193.f8 = _rhs318;
            let _2182_v2;
            _2182_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            _2182_v2 = (_2182_v2).update(true, p0);
            let _2183_v3;
            _2183_v3 = _dafny.Seq.UnicodeFromString("t");
            (globalState).f8 = (_this).fm6((_module.__default.fm1(new BigNumber((_2183_v3).length), p0, p2, globalState)) === (p0), new BigNumber(678), p2, globalState);
            let _2184_v4;
            _2184_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2181_v1).length),p2);
            let _2185_v5;
            _2185_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2184_v4).length),_2184_v4);
            let _2186_v6;
            _2186_v6 = _module.D3.create_DC9(p2, _2185_v5, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), ((_2187_p2) => function (_2188_i1) {
  return _2187_p2;
})(p2))).length)), (_dafny.ZERO).minus(new BigNumber((_2183_v3).length)));
            let _2189_v7;
            _2189_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2186_v6,_2183_v3);
            let _pat_let_tv72 = p2;
            let _pat_let_tv73 = p2;
            _2189_v7 = (_2189_v7).update(function (_pat_let52_0) {
              return function (_2190_dt__update__tmp_h0) {
                return function (_pat_let53_0) {
                  return function (_2191_dt__update_hcf17_h0) {
                    return function (_pat_let54_0) {
                      return function (_2192_dt__update_hcf18_h0) {
                        return _module.D3.create_DC9((_2190_dt__update__tmp_h0).dtor_cf15, (_2190_dt__update__tmp_h0).dtor_cf16, _2191_dt__update_hcf17_h0, _2192_dt__update_hcf18_h0);
                      }(_pat_let54_0);
                    }(_pat_let_tv73);
                  }(_pat_let53_0);
                }(_pat_let_tv72);
              }(_pat_let52_0);
            }(_2186_v6), _module.__default.fm36(p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11)).length), new _dafny.CodePoint('r'.codePointAt(0)), globalState));
          }
        }
      }
      let _2193_v8;
      let _nw323 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _2193_v8 = _nw323;
      let _2194_v9;
      let _nw324 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2194_v9 = _nw324;
      let _2195_v10;
      _2195_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2193_v8,((p0) ? (_2194_v9) : (_2194_v9)));
      _2195_v10 = (_2195_v10).update(_2193_v8, _2194_v9);
      let _2196_v11;
      let _2197_v12;
      let _out35;
      let _out36;
      let _outcollector16 = _module.__default.m0(globalState);
      _out35 = _outcollector16[0];
      _out36 = _outcollector16[1];
      _2196_v11 = _out35;
      _2197_v12 = _out36;
      let _2198_v13;
      _2198_v13 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2199_v14;
      let _nw325 = Array((new BigNumber(7)).toNumber()).fill(false);
      _2199_v14 = _nw325;
      let _2200_v15;
      _2200_v15 = _2199_v14;
      let _2201_v16;
      _2201_v16 = _dafny.Seq.of(_2199_v14, (_2200_v15));
      let _rhs319 = _2198_v13;
      let _rhs320 = _dafny.areEqual(_2201_v16, _2201_v16);
      let _rhs321 = (!(_this.f11)) || (_this.f11);
      let _lhs194 = globalState;
      let _lhs195 = globalState;
      _2198_v13 = _rhs319;
      _lhs194.f8 = _rhs320;
      _lhs195.f8 = _rhs321;
      let _2202_v17;
      _2202_v17 = _module.D6.create_DC17(p0, p2);
      let _2203_v18;
      _2203_v18 = _module.D6.create_DC20(_2202_v17);
      let _2204_v19;
      _2204_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2203_v18,new BigNumber(515));
      r1 = (((_2204_v19).contains(_2203_v18)) ? ((_2204_v19).get(_2203_v18)) : (p2));
      let _2205_v20;
      _2205_v20 = _dafny.Seq.UnicodeFromString("cwra");
      let _2206_v21;
      let _nw326 = new _module.C2();
      _nw326.__ctor(p2, _2205_v20, p2, _module.__default.safeDivisionInt(p2, p2), p0);
      _2206_v21 = _nw326;
      _2206_v21 = _2206_v21;
      let _2207_v22;
      _2207_v22 = _dafny.Seq.of(_2197_v12, !(_2206_v21.f11), p0);
      let _2208_v23;
      _2208_v23 = _dafny.MultiSet.fromElements(_module.__default.fm39(_2197_v12, p0, globalState), _2207_v22, _2207_v22, _2207_v22, _2207_v22);
      let _2209_v24;
      _2209_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_2208_v23).contains(_2207_v22)) ? ((_2208_v23).get(_2207_v22)) : (p2)),_2205_v20);
      r0 = _2209_v24;
      r1 = p2;
      let _2210_v25;
      _2210_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2199_v14,_module.__default.fm1(p2, p0, (_dafny.ZERO).minus(p2), globalState));
      r2 = ((_dafny.Map.Empty.slice().updateUnsafe(_2199_v14,_this.f11)).Merge(_2210_v25)).Merge((_2210_v25).Merge(_2210_v25));
      return [r0, r1, r2];
    }
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return !((!(false)) === ((!(!((_module.D0.create_DC3((_this).f13, !(false))).dtor_cf8))) === (false)));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("pqlbnbxw");
    };
    fm14(globalState) {
      let _this = this;
      let _source31 = _module.D1.create_DC5(_dafny.MultiSet.fromElements((_this).f13, (_this).f12), !(true));
      if (_source31.is_DC5) {
        let _2211___mcc_h0 = (_source31).cf10;
        let _2212___mcc_h1 = (_source31).cf11;
        let _2213_cf11 = _2212___mcc_h1;
        let _2214_cf10 = _2211___mcc_h0;
        return (_this).f12;
      } else if (_source31.is_DC4) {
        let _2215___mcc_h2 = (_source31).cf9;
        let _2216_cf9 = _2215___mcc_h2;
        return (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f13, (_this).f13))));
      } else {
        let _2217___mcc_h3 = (_source31).cf12;
        let _2218_cf12 = _2217___mcc_h3;
        return (_this).f13;
      }
    };
    m8(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _2219_v0;
      _2219_v0 = false;
      let _2220_v1;
      _2220_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2219_v0);
      let _2221_v2;
      _2221_v2 = _dafny.Seq.of(_2219_v0);
      let _2222_v3;
      _2222_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2219_v0,_2219_v0);
      let _2223_v4;
      let _nw327 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _2223_v4 = _nw327;
      let _2224_v5;
      _2224_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2219_v0,(_this).f12);
      let _2225_v6;
      _2225_v6 = _dafny.Map.Empty.slice().updateUnsafe((((_2220_v1).contains(new BigNumber(806))) ? ((_2220_v1).get(new BigNumber(806))) : (_2219_v0)),_2224_v5);
      let _2226_v7;
      _2226_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2223_v4,new BigNumber(((((_2225_v6).contains(_2219_v0)) ? ((_2225_v6).get(_2219_v0)) : (_dafny.Map.Empty.slice().updateUnsafe(_2219_v0,(_this).f13)))).length));
      let _2227_v8;
      _2227_v8 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2228_v9;
      _2228_v9 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), _2227_v8);
      let _2229_v10;
      let _nw328 = Array((new BigNumber(27)).toNumber());
      _nw328[(_dafny.ZERO).toNumber()] = (_this).f12;
      _nw328[(_dafny.ONE).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(2)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(3)).toNumber()] = (_this).f13;
      _nw328[(new BigNumber(4)).toNumber()] = new BigNumber((_2220_v1).length);
      _nw328[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), function (_2230_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length);
      _nw328[(new BigNumber(6)).toNumber()] = (_this).f13;
      _nw328[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f13);
      _nw328[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_this).f12)).cardinality());
      _nw328[(new BigNumber(9)).toNumber()] = (_this).f13;
      _nw328[(new BigNumber(10)).toNumber()] = ((_2219_v0) ? ((_this).f13) : ((_module.__default.fm15(_2219_v0, (_this).f13, _2221_v2, _2219_v0, globalState)).dtor_cf1));
      _nw328[(new BigNumber(11)).toNumber()] = new BigNumber(((_2222_v3).Merge(_2222_v3)).length);
      _nw328[(new BigNumber(12)).toNumber()] = ((_this).f13).multipliedBy(new BigNumber((_2226_v7).length));
      _nw328[(new BigNumber(13)).toNumber()] = ((true) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f13))) : (new BigNumber((_dafny.Seq.UnicodeFromString("fj")).length)));
      _nw328[(new BigNumber(14)).toNumber()] = new BigNumber((((_2219_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_2219_v0,(_this).f12)) : (_dafny.Map.Empty.slice().updateUnsafe(_2219_v0,(_this).f12)))).length);
      _nw328[(new BigNumber(15)).toNumber()] = ((_this).f13).multipliedBy((_this).f13);
      _nw328[(new BigNumber(16)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(17)).toNumber()] = new BigNumber(182);
      _nw328[(new BigNumber(18)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(19)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(20)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(21)).toNumber()] = ((_this).f13).multipliedBy((_this).f13);
      _nw328[(new BigNumber(22)).toNumber()] = (_this).f12;
      _nw328[(new BigNumber(23)).toNumber()] = new BigNumber(-69);
      _nw328[(new BigNumber(24)).toNumber()] = new BigNumber((_2228_v9).cardinality());
      _nw328[(new BigNumber(25)).toNumber()] = new BigNumber(230);
      _nw328[(new BigNumber(26)).toNumber()] = ((_this).f12).plus((_this).f12);
      _2229_v10 = _nw328;
      let _index309 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      (_2229_v10)[_index309] = _module.__default.safeModuloInt((_this).f12, (_this).f12);
      let _index310 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      (_2229_v10)[_index310] = (_this).f12;
      let _2231_v11;
      _2231_v11 = _dafny.Seq.UnicodeFromString("scvtajpa");
      let _index311 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      (_2229_v10)[_index311] = new BigNumber((_2231_v11).length);
      let _2232_v12;
      let _init62 = ((_2233_v8) => function (_2234_i1) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), ((_2235_v8) => function (_2236_i2) {
          return _2235_v8;
        })(_2233_v8));
      })(_2227_v8);
      let _nw329 = Array((new BigNumber(22)).toNumber());
      for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw329.length); _i0_62++) {
        _nw329[_i0_62] = _init62(new BigNumber(_i0_62));
      }
      _2232_v12 = _nw329;
      let _index312 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length));
      (_2232_v12)[_index312] = _2231_v11;
      let _index313 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2232_v12).length));
      (_2232_v12)[_index313] = _2231_v11;
      let _2237_v13;
      _2237_v13 = _dafny.MultiSet.fromElements(_2219_v0);
      let _2238_v14;
      _2238_v14 = _dafny.MultiSet.fromElements((_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))]);
      let _pat_let_tv74 = _2229_v10;
      let _pat_let_tv75 = _2229_v10;
      let _pat_let_tv76 = _2219_v0;
      let _pat_let_tv77 = _2219_v0;
      let _pat_let_tv78 = _2219_v0;
      let _index314 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      let _index315 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      let _index316 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length));
      let _index317 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2232_v12).length));
      let _rhs322 = (_this).f12;
      let _rhs323 = function (_source32) {
        if (_source32.is_DC1) {
          let _2239___mcc_h0 = (_source32).cf1;
          let _2240___mcc_h1 = (_source32).cf2;
          let _2241___mcc_h2 = (_source32).cf3;
          let _2242___mcc_h3 = (_source32).cf4;
          let _2243_cf4 = _2242___mcc_h3;
          let _2244_cf3 = _2241___mcc_h2;
          let _2245_cf2 = _2240___mcc_h1;
          let _2246_cf1 = _2239___mcc_h0;
          return (_2244_cf3).minus(new BigNumber(-932));
        } else if (_source32.is_DC2) {
          let _2247___mcc_h4 = (_source32).cf5;
          let _2248___mcc_h5 = (_source32).cf6;
          let _2249_cf6 = _2248___mcc_h5;
          let _2250_cf5 = _2247___mcc_h4;
          return (_pat_let_tv75)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_pat_let_tv74).length))];
        } else if (_source32.is_DC3) {
          let _2251___mcc_h6 = (_source32).cf7;
          let _2252___mcc_h7 = (_source32).cf8;
          let _2253_cf8 = _2252___mcc_h7;
          let _2254_cf7 = _2251___mcc_h6;
          return (_this).f12;
        } else {
          let _2255___mcc_h8 = (_source32).cf0;
          let _2256_cf0 = _2255___mcc_h8;
          return new BigNumber(-906);
        }
      }(_module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber(((_2237_v13).update(_2219_v0, _module.__default.abs((_this).f12))).cardinality())), _2231_v11, (_this).f13, _2219_v0));
      let _rhs324 = _2231_v11;
      let _rhs325 = function (_source33) {
        if (_source33.is_DC5) {
          let _2257___mcc_h9 = (_source33).cf10;
          let _2258___mcc_h10 = (_source33).cf11;
          let _2259_cf11 = _2258___mcc_h10;
          let _2260_cf10 = _2257___mcc_h9;
          return _pat_let_tv76;
        } else if (_source33.is_DC4) {
          let _2261___mcc_h11 = (_source33).cf9;
          let _2262_cf9 = _2261___mcc_h11;
          return _pat_let_tv77;
        } else {
          let _2263___mcc_h12 = (_source33).cf12;
          let _2264_cf12 = _2263___mcc_h12;
          return _pat_let_tv78;
        }
      }(_module.D1.create_DC5(_2238_v14, _2219_v0));
      let _rhs326 = _2231_v11;
      let _lhs196 = _2229_v10;
      let _lhs197 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      let _lhs198 = _2229_v10;
      let _lhs199 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
      let _lhs200 = _2232_v12;
      let _lhs201 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length));
      let _lhs202 = _2232_v12;
      let _lhs203 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_2232_v12).length));
      _lhs196[_lhs197] = _rhs322;
      _lhs198[_lhs199] = _rhs323;
      _lhs200[_lhs201] = _rhs324;
      _2219_v0 = _rhs325;
      _lhs202[_lhs203] = _rhs326;
      let _2265_i3;
      _2265_i3 = _dafny.ZERO;
      L10: {
        while ((_this).fm7(globalState)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2265_i3)) {
              break L10;
            }
            _2265_i3 = (_2265_i3).plus(_dafny.ONE);
            let _2266_v15;
            _2266_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2227_v8,_module.__default.fm1((_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], _2219_v0, (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], globalState));
            let _2267_v16;
            _2267_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))],(_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))]);
            let _2268_v17;
            _2268_v17 = _dafny.Set.fromElements(new BigNumber((_2267_v16).length), (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))]);
            let _2269_v18;
            _2269_v18 = _dafny.Set.fromElements(new BigNumber(((_this).fm8((_this).f12, _2266_v15, _2268_v17, globalState)).length));
            let _2270_v19;
            _2270_v19 = _dafny.Seq.of(_2238_v14);
            let _2271_v20;
            let _nw330 = Array((new BigNumber(10)).toNumber());
            _nw330[(_dafny.ZERO).toNumber()] = _2238_v14;
            _nw330[(_dafny.ONE).toNumber()] = _2238_v14;
            _nw330[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_2269_v18).length), (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], (_this).f12, new BigNumber((_2221_v2).length));
            _nw330[(new BigNumber(3)).toNumber()] = _2238_v14;
            _nw330[(new BigNumber(4)).toNumber()] = (_2270_v19)[_module.__default.safeIndex((_this).f12, new BigNumber((_2270_v19).length))];
            _nw330[(new BigNumber(5)).toNumber()] = _2238_v14;
            _nw330[(new BigNumber(6)).toNumber()] = _2238_v14;
            _nw330[(new BigNumber(7)).toNumber()] = (_2238_v14).update((_this).f12, _module.__default.abs((_this).f13));
            _nw330[(new BigNumber(8)).toNumber()] = _2238_v14;
            _nw330[(new BigNumber(9)).toNumber()] = (_2270_v19)[_module.__default.safeIndex((_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], new BigNumber((_2270_v19).length))];
            _2271_v20 = _nw330;
            let _2272_v21;
            _2272_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2219_v0,_2271_v20);
            _2272_v21 = (_2272_v21).update(_2219_v0, _2271_v20);
            _2231_v11 = _dafny.Seq.Concat(_dafny.Seq.Concat((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _dafny.Seq.update(_2231_v11, _module.__default.safeIndex(new BigNumber(589), new BigNumber((_2231_v11).length)), _2227_v8)), _2231_v11);
            let _2273_v22;
            let _nw331 = new _module.C3();
            _nw331.__ctor((_this).f13, (_this).f13);
            _2273_v22 = _nw331;
            _2228_v9 = (_2228_v9).update(_2227_v8, _module.__default.abs(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f13)).Merge(_2267_v16)).length)));
          }
        }
      }
      let _hi10 = (_this).f12;
      for (let _2274_i4 = ((_2219_v0) ? ((_this).f12) : ((_this).f13)); _2274_i4.isLessThan(_hi10); _2274_i4 = _2274_i4.plus(_dafny.ONE)) {
        let _2275_v23;
        _2275_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(271),_2221_v2);
        let _2276_v24;
        let _nw332 = Array((new BigNumber(14)).toNumber());
        _nw332[(_dafny.ZERO).toNumber()] = _2221_v2;
        _nw332[(_dafny.ONE).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_2219_v0);
        _nw332[(new BigNumber(3)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_this).fm7(globalState), _2219_v0);
        _nw332[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2221_v2, _2221_v2);
        _nw332[(new BigNumber(6)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(7)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(8)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_2221_v2, _module.__default.safeIndex(_2274_i4, new BigNumber((_2221_v2).length)), _module.__default.fm1(new BigNumber(144), _2219_v0, new BigNumber((_dafny.Seq.UnicodeFromString("nwqyqwpb")).length), globalState)), _2221_v2);
        _nw332[(new BigNumber(10)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update((((_2275_v23).contains(_2274_i4)) ? ((_2275_v23).get(_2274_i4)) : (_dafny.Seq.of(_2219_v0, false))), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f12), new BigNumber(((((_2275_v23).contains(_2274_i4)) ? ((_2275_v23).get(_2274_i4)) : (_dafny.Seq.of(_2219_v0, false)))).length)), _2219_v0), _2221_v2), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update((((_2275_v23).contains(_2274_i4)) ? ((_2275_v23).get(_2274_i4)) : (_dafny.Seq.of(_2219_v0, false))), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f12), new BigNumber(((((_2275_v23).contains(_2274_i4)) ? ((_2275_v23).get(_2274_i4)) : (_dafny.Seq.of(_2219_v0, false)))).length)), _2219_v0), _2221_v2)).length)), _2219_v0);
        _nw332[(new BigNumber(12)).toNumber()] = _2221_v2;
        _nw332[(new BigNumber(13)).toNumber()] = _2221_v2;
        _2276_v24 = _nw332;
        let _2277_v25;
        _2277_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _module.__default.safeIndex((_this).f12, new BigNumber(((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))]).length)), _2227_v8)).length),new BigNumber((_dafny.Seq.of(_2274_i4, (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))], (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))])).length));
        let _2278_v26;
        _2278_v26 = _dafny.Seq.of(_2223_v4, _2229_v10, _2223_v4);
        let _index318 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        let _index319 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        let _rhs327 = _2276_v24;
        let _rhs328 = ((_2219_v0) ? (new BigNumber((_2277_v25).length)) : ((_this).f12));
        let _rhs329 = new BigNumber((_dafny.Seq.Concat(_2278_v26, _dafny.Seq.Concat(_dafny.Seq.of(_2229_v10), _2278_v26))).length);
        let _lhs204 = _2229_v10;
        let _lhs205 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        let _lhs206 = _2229_v10;
        let _lhs207 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        _2276_v24 = _rhs327;
        _lhs204[_lhs205] = _rhs328;
        _lhs206[_lhs207] = _rhs329;
        let _2279_v27;
        let _nw333 = new _module.C2();
        _nw333.__ctor((_this).f13, (_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], (_2274_i4).plus(new BigNumber((_dafny.Seq.of(_2219_v0, _2219_v0)).length)), ((_this).f12).minus((_this).f13), _2219_v0);
        _2279_v27 = _nw333;
        _2219_v0 = !(function () {
          let _coll88 = new _dafny.Map();
          for (const _compr_88 of (function () {
            let _coll89 = new _dafny.Map();
            for (const _compr_89 of (_dafny.Seq.of((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _dafny.Seq.UnicodeFromString("arocdqm"), (_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))])).Elements) {
              let _2280_v29 = _compr_89;
              if (_dafny.Seq.contains(_dafny.Seq.of((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _dafny.Seq.UnicodeFromString("arocdqm"), (_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))]), _2280_v29)) {
                _coll89.push([_2280_v29,new BigNumber((_2277_v25).length)]);
              }
            }
            return _coll89;
          }()).Keys.Elements) {
            let _2281_v28 = _compr_88;
            if ((function () {
              let _coll90 = new _dafny.Map();
              for (const _compr_90 of (_dafny.Seq.of((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _dafny.Seq.UnicodeFromString("arocdqm"), (_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))])).Elements) {
                let _2280_v29 = _compr_90;
                if (_dafny.Seq.contains(_dafny.Seq.of((_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))], _dafny.Seq.UnicodeFromString("arocdqm"), (_2232_v12)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2232_v12).length))]), _2280_v29)) {
                  _coll90.push([_2280_v29,new BigNumber((_2277_v25).length)]);
                }
              }
              return _coll90;
            }()).contains(_2281_v28)) {
              _coll88.push([_2281_v28,_dafny.Set.fromElements((_2279_v27).f26)]);
            }
          }
          return _coll88;
        }()).contains(_2231_v11);
        let _2282_v30;
        _2282_v30 = _module.D9.create_DC26(_2227_v8);
        let _2283_v31;
        let _nw334 = Array((new BigNumber(7)).toNumber());
        _nw334[(_dafny.ZERO).toNumber()] = _2227_v8;
        _nw334[(_dafny.ONE).toNumber()] = _2227_v8;
        _nw334[(new BigNumber(2)).toNumber()] = _2227_v8;
        _nw334[(new BigNumber(3)).toNumber()] = _2227_v8;
        _nw334[(new BigNumber(4)).toNumber()] = (_2282_v30).dtor_cf48;
        _nw334[(new BigNumber(5)).toNumber()] = _2227_v8;
        _nw334[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _2283_v31 = _nw334;
        let _2284_v32;
        _2284_v32 = _dafny.MultiSet.fromElements(_2223_v4, _2229_v10);
        let _index320 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        let _rhs330 = (_2229_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length))];
        let _rhs331 = (((_2284_v32).IsSubsetOf(_2284_v32)) ? (_2283_v31) : (_2283_v31));
        let _lhs208 = _2229_v10;
        let _lhs209 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2229_v10).length));
        _lhs208[_lhs209] = _rhs330;
        _2283_v31 = _rhs331;
      }
      r1 = _2219_v0;
      let _2285_v33;
      _2285_v33 = _module.D6.create_DC19(_2219_v0, _2219_v0, _2219_v0, true, _2221_v2);
      let _pat_let_tv79 = _2219_v0;
      let _pat_let_tv80 = _2224_v5;
      let _pat_let_tv81 = _2229_v10;
      let _pat_let_tv82 = _2229_v10;
      let _pat_let_tv83 = _2229_v10;
      let _pat_let_tv84 = _2229_v10;
      let _pat_let_tv85 = _2229_v10;
      let _pat_let_tv86 = _2229_v10;
      let _pat_let_tv87 = _2231_v11;
      let _pat_let_tv88 = _2224_v5;
      let _pat_let_tv89 = _2229_v10;
      r0 = function (_source34) {
        if (_source34.is_DC17) {
          let _2286___mcc_h13 = (_source34).cf30;
          let _2287___mcc_h14 = (_source34).cf31;
          let _2288_cf31 = _2287___mcc_h14;
          let _2289_cf30 = _2286___mcc_h13;
          if (_pat_let_tv79) {
            return _dafny.Seq.of((_this).f12);
          } else {
            return _dafny.Seq.of((_this).f12, new BigNumber((_pat_let_tv80).length), (_pat_let_tv82)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_pat_let_tv81).length))], (_dafny.ZERO).minus(_2288_cf31), _2288_cf31);
          }
        } else if (_source34.is_DC18) {
          let _2290___mcc_h15 = (_source34).cf32;
          let _2291_cf32 = _2290___mcc_h15;
          if (!((_module.D4.create_DC12(_2291_cf32, (_this).f13, (_pat_let_tv84)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_pat_let_tv83).length))])).dtor_cf21)) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-196)), function (_2292_i5) {
              return (_this).f13;
            });
          } else {
            return _dafny.Seq.of(new BigNumber(-548), (_this).f12);
          }
        } else if (_source34.is_DC19) {
          let _2293___mcc_h16 = (_source34).cf33;
          let _2294___mcc_h17 = (_source34).cf34;
          let _2295___mcc_h18 = (_source34).cf35;
          let _2296___mcc_h19 = (_source34).cf36;
          let _2297___mcc_h20 = (_source34).cf37;
          let _2298_cf37 = _2297___mcc_h20;
          let _2299_cf36 = _2296___mcc_h19;
          let _2300_cf35 = _2295___mcc_h18;
          let _2301_cf34 = _2294___mcc_h17;
          let _2302_cf33 = _2293___mcc_h16;
          return _dafny.Seq.of((_pat_let_tv86)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_pat_let_tv85).length))], new BigNumber((_pat_let_tv87).length));
        } else if (_source34.is_DC16) {
          let _2303___mcc_h21 = (_source34).cf29;
          let _2304_cf29 = _2303___mcc_h21;
          return _dafny.Seq.of(new BigNumber((_pat_let_tv88).length), (_this).f12);
        } else {
          let _2305___mcc_h22 = (_source34).cf38;
          let _2306_cf38 = _2305___mcc_h22;
          return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(826)), ((_2307_v10) => function (_2308_i6) {
            return (_2307_v10)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_2307_v10).length))];
          })(_pat_let_tv89)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f13)).length), (_this).f13));
        }
      }(_2285_v33);
      r1 = _2219_v0;
      return [r0, r1];
    }
  };

  $module.C12 = class C12 {
    constructor () {
      this._tname = "_module.C12";
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this.f18 = _dafny.MultiSet.Empty;
      this._f19 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f18, f19, f11, f12, f13) {
      let _this = this;
      (_this).f18 = f18;
      (_this)._f19 = f19;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source35 = _module.D0.create_DC3(new BigNumber((_dafny.Seq.of((_this).f19, (_this).f19, (_this).f13, new BigNumber((_dafny.Seq.UnicodeFromString("gw")).length), (_this).f13)).length), _this.f11);
      if (_source35.is_DC1) {
        let _2309___mcc_h0 = (_source35).cf1;
        let _2310___mcc_h1 = (_source35).cf2;
        let _2311___mcc_h2 = (_source35).cf3;
        let _2312___mcc_h3 = (_source35).cf4;
        let _2313_cf4 = _2312___mcc_h3;
        let _2314_cf3 = _2311___mcc_h2;
        let _2315_cf2 = _2310___mcc_h1;
        let _2316_cf1 = _2309___mcc_h0;
        return _module.D0.create_DC1(new BigNumber((function () {
  let _coll91 = new _dafny.Map();
  for (const _compr_91 of _dafny.IntegerRange(new BigNumber(772), new BigNumber(95))) {
    let _2317_v0 = _compr_91;
    if (((new BigNumber(772)).isLessThanOrEqualTo(_2317_v0)) && ((_2317_v0).isLessThan(new BigNumber(95)))) {
      _coll91.push([(_2317_v0).plus(_2314_cf3),_dafny.Seq.of(_this.f11)]);
    }
  }
  return _coll91;
}()).length), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("xyqxy"), _module.__default.safeIndex((_this).f19, new BigNumber((_dafny.Seq.UnicodeFromString("xyqxy")).length)), new _dafny.CodePoint('x'.codePointAt(0))), (_this).f12, false);
      } else if (_source35.is_DC2) {
        let _2318___mcc_h4 = (_source35).cf5;
        let _2319___mcc_h5 = (_source35).cf6;
        let _2320_cf6 = _2319___mcc_h5;
        let _2321_cf5 = _2318___mcc_h4;
        return _module.D0.create_DC1(new BigNumber(494), _dafny.Seq.UnicodeFromString("ccyu"), new BigNumber(785), _this.f11);
      } else if (_source35.is_DC3) {
        let _2322___mcc_h6 = (_source35).cf7;
        let _2323___mcc_h7 = (_source35).cf8;
        let _2324_cf8 = _2323___mcc_h7;
        let _2325_cf7 = _2322___mcc_h6;
        return _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(_this.f11, _this.f11)).length), _dafny.Seq.UnicodeFromString("apji"), (_this).f12, _this.f11);
      } else {
        let _2326___mcc_h8 = (_source35).cf0;
        let _2327_cf0 = _2326___mcc_h8;
        if (_this.f11) {
          return _module.D0.create_DC1((_this).f12, _dafny.Seq.UnicodeFromString("niwbwv"), (_this).f13, true);
        } else {
          return _module.D0.create_DC1((_this).f13, _dafny.Seq.UnicodeFromString("kidqydpry"), (_this).f13, _this.f11);
        }
      }
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f11;
    };
    fm7(globalState) {
      let _this = this;
      return _this.f11;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("f");
    };
    fm13(p0, globalState) {
      let _this = this;
      return _this.f11;
    };
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2328_v0;
      _2328_v0 = _dafny.Seq.of(p1);
      let _2329_v1;
      let _nw335 = new _module.C4();
      _nw335.__ctor(new BigNumber((_dafny.Seq.Concat(_2328_v0, _2328_v0)).length), true);
      _2329_v1 = _nw335;
      let _index321 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
      (p0)[_index321] = (_2329_v1).f25;
      let _index322 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
      (p0)[_index322] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.fm2(globalState), (_this).f13)));
      if (_this.f11) {
        (globalState).f8 = !(_module.__default.fm1(p1, true, ((_this).f12).plus((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]), globalState));
        let _2330_v2;
        _2330_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
        let _2331_v3;
        _2331_v3 = _dafny.Set.fromElements(_this.f11, false, _this.f11);
        let _2332_v4;
        _2332_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(17),_2331_v3);
        (globalState).f8 = (((!((_this).fm6((((_2330_v2).contains(true)) ? ((_2330_v2).get(true)) : (_module.__default.fm1((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))], _this.f11, (p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))], globalState))), (_dafny.ZERO).minus((_2329_v1).f25), new BigNumber((_2332_v4).length), globalState))) || (true)) ? (((_this.f11) ? (_this.f11) : (_this.f11))) : (!(((_dafny.ZERO).minus((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))])).isLessThan((_this).f12))));
        let _2333_v5;
        let _init63 = function (_2334_i0) {
          return (_2334_i0).multipliedBy(new BigNumber(139));
        };
        let _nw336 = Array((new BigNumber(10)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw336.length); _i0_63++) {
          _nw336[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2333_v5 = _nw336;
        _2333_v5 = ((_this.f11) ? (_2333_v5) : (p0));
        let _2335_v6;
        let _nw337 = new _module.C5();
        _nw337.__ctor((_this).f13);
        _2335_v6 = _nw337;
        if (((!(!(_this.f11)) || (_this.f11)) ? (_this.f11) : (true))) {
          let _2336_v7;
          _2336_v7 = _dafny.MultiSet.fromElements(_this.f11, _this.f11, _this.f11, _this.f11, _this.f11);
          let _2337_v8;
          _2337_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2336_v7);
          let _2338_v9;
          _2338_v9 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_2337_v8).length)).minus((_2329_v1).f25),((_this.f11) ? (_this.f11) : (_this.f11)));
          _2338_v9 = (_2338_v9).update((_dafny.ZERO).minus((_dafny.ZERO).minus(((_2329_v1).f25).minus((_2329_v1).f25))), _this.f11);
          let _2339_v10;
          let _nw338 = new _module.C1();
          _nw338.__ctor((_dafny.ZERO).minus(_module.__default.safeModuloInt((_2329_v1).f25, (_this).f12)));
          _2339_v10 = _nw338;
          (globalState).f8 = !(_this.f11);
          (globalState).f8 = !(_this.f11);
          let _2340_v11;
          _2340_v11 = _dafny.Map.Empty.slice().updateUnsafe((_2329_v1).f25,_2333_v5);
          _2340_v11 = (_2340_v11).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2339_v10).fm34(globalState),_2333_v5));
        } else {
          r0 = ((_this).f19).plus(((_module.__default.fm1((_2329_v1).f25, _this.f11, new BigNumber(-204), globalState)) ? ((_2328_v0)[_module.__default.safeIndex(new BigNumber(481), new BigNumber((_2328_v0).length))]) : ((_this).f19)));
          let _index323 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
          (p0)[_index323] = (new BigNumber(13)).plus((_this).f19);
          let _2341_v12;
          _2341_v12 = _dafny.Seq.UnicodeFromString("nvld");
          let _2342_v13;
          _2342_v13 = new _dafny.CodePoint('j'.codePointAt(0));
          let _2343_v14;
          let _nw339 = new _module.C2();
          _nw339.__ctor(_module.__default.fm2(globalState), _dafny.Seq.update(_dafny.Seq.Concat(_2341_v12, _2341_v12), _module.__default.safeIndex((_2329_v1).f25, new BigNumber((_dafny.Seq.Concat(_2341_v12, _2341_v12)).length)), _2342_v13), new BigNumber(422), ((_dafny.ZERO).minus((_this).f12)).minus(new BigNumber((_2341_v12).length)), !(_this.f11));
          _2343_v14 = _nw339;
          _2343_v14 = _2343_v14;
          let _2344_v15;
          _2344_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_dafny.ZERO).minus((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]));
          let _2345_v16;
          _2345_v16 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))],_module.__default.safeModuloInt((((_2344_v15).contains((_2335_v6).fm27((_2329_v1).f25, _dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), ((_2348_v13) => function (_2349_i1) {
            return _2348_v13;
          })(_2342_v13)), _this.f11, globalState))) ? ((_2344_v15).get((_2335_v6).fm27((_2329_v1).f25, _dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), ((_2346_v13) => function (_2347_i1) {
            return _2346_v13;
          })(_2342_v13)), _this.f11, globalState))) : ((_2329_v1).f25)), new BigNumber((_2344_v15).length)));
          let _2350_v17;
          _2350_v17 = _dafny.MultiSet.fromElements(_this.f11, _this.f11, !(_this.f11));
          let _2351_v18;
          _2351_v18 = _dafny.Seq.of(_this.f11, _this.f11);
          _2345_v16 = (_2345_v16).update((new BigNumber((_2350_v17).cardinality())).minus((_2329_v1).f25), new BigNumber((_2351_v18).length));
          (globalState).f8 = !(_this.f11) || (_this.f11);
        }
      } else {
        (globalState).f8 = _this.f11;
        let _2352_v19;
        _2352_v19 = _dafny.MultiSet.fromElements((_this).f12);
        let _2353_v20;
        _2353_v20 = new _dafny.CodePoint('e'.codePointAt(0));
        let _2354_v21;
        _2354_v21 = _dafny.Seq.of(_2353_v20, _2353_v20, _2353_v20, _2353_v20, _2353_v20);
        let _2355_v22;
        _2355_v22 = _module.D1.create_DC5(((_2352_v19).update((_2329_v1).f25, _module.__default.abs(new BigNumber((_2354_v21).length)))).Intersect(_2352_v19), _this.f11);
        _2355_v22 = _2355_v22;
        let _2356_v23;
        let _nw340 = Array((new BigNumber(3)).toNumber());
        _nw340[(_dafny.ZERO).toNumber()] = _this.f11;
        _nw340[(_dafny.ONE).toNumber()] = false;
        _nw340[(new BigNumber(2)).toNumber()] = (_2352_v19).IsDisjointFrom(_2352_v19);
        _2356_v23 = _nw340;
        let _2357_v24;
        _2357_v24 = _module.D9.create_DC26(_2353_v20);
        let _2358_v25;
        _2358_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,true);
        let _2359_v26;
        _2359_v26 = _module.D0.create_DC2(new BigNumber((_2358_v25).length), p0);
        let _2360_v27;
        _2360_v27 = _dafny.Seq.of(_2354_v21, _dafny.Seq.UnicodeFromString("yg"), _dafny.Seq.UnicodeFromString("oakqppwdg"), _module.__default.fm36(new BigNumber((_2354_v21).length), (_this).f19, _2353_v20, globalState), _dafny.Seq.UnicodeFromString("qwsviefmw"));
        let _2361_v28;
        _2361_v28 = _dafny.MultiSet.fromElements(_this.f11, _this.f11);
        let _2362_v29;
        _2362_v29 = _dafny.Seq.of(_this.f11, _this.f11, _this.f11);
        let _2363_v30;
        _2363_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_this.f11);
        let _2364_v32;
        _2364_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2353_v20,(_2362_v29)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll92 = new _dafny.Set();
          for (const _compr_92 of (_2363_v30).Keys.Elements) {
            let _2365_v31 = _compr_92;
            if ((_2363_v30).contains(_2365_v31)) {
              _coll92.add((_2365_v31).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length)));
            }
          }
          return _coll92;
        }()).length), new BigNumber((_2362_v29).length))]);
        let _2366_v33;
        _2366_v33 = _dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]);
        let _2367_v34;
        let _nw341 = Array((new BigNumber(28)).toNumber());
        _nw341[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_2354_v21, _2354_v21), _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm2(globalState)), new BigNumber((_dafny.Seq.Concat(_2354_v21, _2354_v21)).length)), (_2357_v24).dtor_cf48);
        _nw341[(_dafny.ONE).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(2)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(3)).toNumber()] = ((_this.f11) ? (_dafny.Seq.update(_2354_v21, _module.__default.safeIndex((_2359_v26).dtor_cf5, new BigNumber((_2354_v21).length)), _2353_v20)) : (_dafny.Seq.UnicodeFromString("llmohotq")));
        _nw341[(new BigNumber(4)).toNumber()] = (_2360_v27)[_module.__default.safeIndex((_dafny.ZERO).minus((_2329_v1).f25), new BigNumber((_2360_v27).length))];
        _nw341[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), function (_2368_i2) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }), _module.__default.safeIndex(_module.__default.fm2(globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), function (_2369_i2) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })).length)), new _dafny.CodePoint('o'.codePointAt(0)));
        _nw341[(new BigNumber(6)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_2354_v21, _module.__default.safeIndex((((_2361_v28).contains(_this.f11)) ? ((_2361_v28).get(_this.f11)) : ((_this).f12)), new BigNumber((_2354_v21).length)), _2353_v20), _dafny.Seq.UnicodeFromString("vx")), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2354_v21, _module.__default.safeIndex((((_2361_v28).contains(_this.f11)) ? ((_2361_v28).get(_this.f11)) : ((_this).f12)), new BigNumber((_2354_v21).length)), _2353_v20), _dafny.Seq.UnicodeFromString("vx"))).length)), new _dafny.CodePoint('q'.codePointAt(0)));
        _nw341[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), ((_2370_v20) => function (_2371_i3) {
          return _2370_v20;
        })(_2353_v20));
        _nw341[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-35)), ((_2372_v20) => function (_2373_i4) {
          return _2372_v20;
        })(_2353_v20));
        _nw341[(new BigNumber(10)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(11)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(12)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(13)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(14)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(15)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(16)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(17)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(18)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(19)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(20)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(21)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_2354_v21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), ((_2374_v20) => function (_2375_i5) {
          return _2374_v20;
        })(_2353_v20)));
        _nw341[(new BigNumber(23)).toNumber()] = _2354_v21;
        _nw341[(new BigNumber(24)).toNumber()] = _dafny.Seq.update((_this).fm8((_2329_v1).f25, _2364_v32, _2366_v33, globalState), _module.__default.safeIndex((_dafny.ZERO).minus((_2329_v1).f25), new BigNumber(((_this).fm8((_2329_v1).f25, _2364_v32, _2366_v33, globalState)).length)), _2353_v20);
        _nw341[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("tc");
        _nw341[(new BigNumber(26)).toNumber()] = ((_this.f11) ? (_2354_v21) : (_2354_v21));
        _nw341[(new BigNumber(27)).toNumber()] = _2354_v21;
        _2367_v34 = _nw341;
        let _index324 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2367_v34).length));
        (_2367_v34)[_index324] = _2354_v21;
        let _2376_v35;
        _2376_v35 = _dafny.Map.Empty.slice().updateUnsafe((_2329_v1).f25,(_2329_v1).f25);
        let _2377_v36;
        _2377_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(375),_2376_v35);
        let _2378_v37;
        _2378_v37 = _module.D3.create_DC9((_this).f19, _2377_v36, p1, (_this).f13);
        let _index325 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2367_v34).length));
        let _rhs332 = _2356_v23;
        let _rhs333 = _2354_v21;
        let _rhs334 = _2353_v20;
        let _rhs335 = _dafny.areEqual(_dafny.Seq.Concat(_2354_v21, _2354_v21), ((true) ? (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nxs"), _module.__default.safeIndex((_2329_v1).f25, new BigNumber((_dafny.Seq.UnicodeFromString("nxs")).length)), new _dafny.CodePoint('t'.codePointAt(0)))) : (_2354_v21)));
        let _rhs336 = ((new BigNumber((_dafny.Set.fromElements(_2378_v37)).length)).plus((_2329_v1).f25)).minus(p1);
        let _lhs210 = _2367_v34;
        let _lhs211 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2367_v34).length));
        let _lhs212 = _this;
        _2356_v23 = _rhs332;
        _lhs210[_lhs211] = _rhs333;
        _2353_v20 = _rhs334;
        _lhs212.f11 = _rhs335;
        r0 = _rhs336;
        let _index326 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        (p0)[_index326] = (_dafny.ZERO).minus((_this).f19);
        let _index327 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        (p0)[_index327] = (_module.__default.safeModuloInt((_2328_v0)[_module.__default.safeIndex((_this).f13, new BigNumber((_2328_v0).length))], (_this).f13)).multipliedBy((_this).f19);
      }
      let _2379_v38;
      let _nw342 = Array((new BigNumber(12)).toNumber()).fill(false);
      _2379_v38 = _nw342;
      let _index328 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
      (_2379_v38)[_index328] = _this.f11;
      let _2380_v39;
      let _init64 = function (_2381_i6) {
        return _module.__default.safeDivisionInt(_2381_i6, (_this).f19);
      };
      let _nw343 = Array((new BigNumber(25)).toNumber());
      for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw343.length); _i0_64++) {
        _nw343[_i0_64] = _init64(new BigNumber(_i0_64));
      }
      _2380_v39 = _nw343;
      let _index329 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
      let _rhs337 = _this.f11;
      let _rhs338 = _2380_v39;
      let _lhs213 = _2379_v38;
      let _lhs214 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
      _lhs213[_lhs214] = _rhs337;
      _2380_v39 = _rhs338;
      let _2382_v40;
      _2382_v40 = _module.D12.create_DC35(_module.__default.safeDivisionInt(p1, _module.__default.fm2(globalState)));
      let _source36 = _2382_v40;
      if (_source36.is_DC33) {
        let _2383___mcc_h0 = (_source36).cf58;
        let _2384_cf58 = _2383___mcc_h0;
        _2380_v39 = _2380_v39;
        let _2385_v41;
        let _nw344 = new _module.C10();
        _nw344.__ctor(_this.f11);
        _2385_v41 = _nw344;
        _2385_v41 = (((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]) ? (_2385_v41) : (_2385_v41));
        let _2386_v42;
        let _nw345 = new _module.C0();
        _nw345.__ctor((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))], _this.f11, (_2329_v1).f25, p1);
        _2386_v42 = _nw345;
        let _2387_v43;
        _2387_v43 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2386_v42);
        let _2388_v44;
        _2388_v44 = _dafny.Seq.of(_2386_v42.f29);
        let _2389_v45;
        _2389_v45 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2387_v43,(_2388_v44)[_module.__default.safeIndex((_2329_v1).f25, new BigNumber((_2388_v44).length))])).length));
        let _2390_v46;
        _2390_v46 = _dafny.Seq.UnicodeFromString("mdwgotbd");
        let _2391_v47;
        _2391_v47 = _module.D0.create_DC1(p1, _2390_v46, (_2329_v1).f25, _2386_v42.f29);
        let _2392_v48;
        _2392_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_2391_v47).dtor_cf2, _2390_v46)).length),(_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
        let _index330 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
        (_2379_v38)[_index330] = !(_2386_v42.f29) || ((new BigNumber((_2392_v48).length)).isLessThanOrEqualTo(new BigNumber((_2389_v45).cardinality())));
        _2390_v46 = _dafny.Seq.UnicodeFromString("xjd");
      } else if (_source36.is_DC34) {
        let _2393___mcc_h1 = (_source36).cf59;
        let _2394_cf59 = _2393___mcc_h1;
        let _2395_v49;
        _2395_v49 = _dafny.MultiSet.fromElements(((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]).isLessThan(new BigNumber((_2328_v0).length)), false);
        let _index331 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _index332 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _rhs339 = (((_2395_v49).contains(_this.f11)) ? ((_2395_v49).get(_this.f11)) : (p1));
        let _rhs340 = p1;
        let _lhs215 = p0;
        let _lhs216 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _lhs217 = p0;
        let _lhs218 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        _lhs215[_lhs216] = _rhs339;
        _lhs217[_lhs218] = _rhs340;
        let _2396_v50;
        _2396_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2379_v38);
        let _2397_v51;
        _2397_v51 = _2379_v38;
        let _2398_v52;
        let _nw346 = Array((new BigNumber(18)).toNumber());
        _nw346[(_dafny.ZERO).toNumber()] = _2379_v38;
        _nw346[(_dafny.ONE).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(2)).toNumber()] = (((_2396_v50).contains(new BigNumber(264))) ? ((_2396_v50).get(new BigNumber(264))) : (_2379_v38));
        _nw346[(new BigNumber(3)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(4)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(5)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(6)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(7)).toNumber()] = ((_this.f11) ? (_2379_v38) : (_2379_v38));
        _nw346[(new BigNumber(8)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(9)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(10)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(11)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(12)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(13)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(14)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(15)).toNumber()] = (_2397_v51);
        _nw346[(new BigNumber(16)).toNumber()] = _2379_v38;
        _nw346[(new BigNumber(17)).toNumber()] = (_2379_v38);
        _2398_v52 = _nw346;
        _2398_v52 = _2398_v52;
        let _index333 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        (p0)[_index333] = _module.__default.safeDivisionInt(((_2329_v1).f25).minus((_this).f12), (_2329_v1).f25);
        let _2399_v53;
        _2399_v53 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2400_v54;
        _2400_v54 = _dafny.Seq.of(_2399_v53);
        let _2401_v55;
        _2401_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2395_v49,new BigNumber(415));
        let _2402_v57;
        _2402_v57 = _dafny.Seq.of(p0, _2394_cf59);
        let _2403_v58;
        _2403_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2402_v57).length),p1);
        let _index334 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _rhs341 = (_this).f13;
        let _rhs342 = _2400_v54;
        let _rhs343 = function () {
          let _coll93 = new _dafny.Map();
          for (const _compr_93 of (_dafny.Seq.of(_2395_v49, _dafny.MultiSet.fromElements((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]))).Elements) {
            let _2404_v56 = _compr_93;
            if (_dafny.Seq.contains(_dafny.Seq.of(_2395_v49, _dafny.MultiSet.fromElements((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))])), _2404_v56)) {
              _coll93.push([_2404_v56,(_2329_v1).f25]);
            }
          }
          return _coll93;
        }();
        let _rhs344 = _2328_v0;
        let _rhs345 = (((_this).f13).plus((_2329_v1).f25)).multipliedBy(((((_2403_v58).contains((_2329_v1).f25)) ? ((_2403_v58).get((_2329_v1).f25)) : (new BigNumber((_2403_v58).length)))).minus((_2329_v1).f25));
        let _lhs219 = p0;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        _lhs219[_lhs220] = _rhs341;
        _2400_v54 = _rhs342;
        _2401_v55 = _rhs343;
        _2328_v0 = _rhs344;
        r0 = _rhs345;
      } else if (_source36.is_DC35) {
        let _2405___mcc_h2 = (_source36).cf60;
        let _2406_cf60 = _2405___mcc_h2;
        _2380_v39 = _2380_v39;
        let _2407_v59;
        let _2408_v60;
        let _out37;
        let _out38;
        let _outcollector17 = _module.__default.m0(globalState);
        _out37 = _outcollector17[0];
        _out38 = _outcollector17[1];
        _2407_v59 = _out37;
        _2408_v60 = _out38;
        let _index335 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        (p0)[_index335] = _module.__default.safeModuloInt((_2329_v1).f25, (_this).f13);
        let _2409_v61;
        _2409_v61 = _dafny.Seq.of(p0, _2380_v39);
        let _2410_v62;
        let _nw347 = Array((new BigNumber(8)).toNumber());
        _nw347[(_dafny.ZERO).toNumber()] = (_2409_v61)[_module.__default.safeIndex(new BigNumber((_module.__default.fm50(globalState)).length), new BigNumber((_2409_v61).length))];
        _nw347[(_dafny.ONE).toNumber()] = _2380_v39;
        _nw347[(new BigNumber(2)).toNumber()] = p0;
        _nw347[(new BigNumber(3)).toNumber()] = p0;
        _nw347[(new BigNumber(4)).toNumber()] = _2380_v39;
        _nw347[(new BigNumber(5)).toNumber()] = p0;
        _nw347[(new BigNumber(6)).toNumber()] = _2380_v39;
        _nw347[(new BigNumber(7)).toNumber()] = p0;
        _2410_v62 = _nw347;
        let _index336 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2410_v62).length));
        (_2410_v62)[_index336] = p0;
        let _2411_v63;
        _2411_v63 = _dafny.Seq.UnicodeFromString("hate");
        let _2412_v64;
        _2412_v64 = _dafny.Set.fromElements((_2329_v1).f25, (_this).f19, p1, p1);
        let _2413_v66;
        _2413_v66 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2414_v67;
        _2414_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2413_v66,_this.f11);
        let _2415_v68;
        _2415_v68 = _dafny.Map.Empty.slice().updateUnsafe((_2329_v1).f25,(_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
        let _index337 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2410_v62).length));
        let _index338 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _index339 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _rhs346 = p0;
        let _rhs347 = (_2329_v1).f25;
        let _rhs348 = (_this).fm8(_module.__default.safeDivisionInt((_this).f12, new BigNumber((function () {
          let _coll94 = new _dafny.Set();
          for (const _compr_94 of (_2412_v64).Elements) {
            let _2416_v65 = _compr_94;
            if ((_2412_v64).contains(_2416_v65)) {
              _coll94.add(_module.__default.safeModuloInt(_2416_v65, new BigNumber((_dafny.Seq.of(false, true)).length)));
            }
          }
          return _coll94;
        }()).length)), (_2414_v67).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2413_v66,_this.f11)), _dafny.Set.fromElements(p1, (p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))], (_this).f19, new BigNumber((_2415_v68).length), (_dafny.ZERO).minus(new BigNumber((_2415_v68).length))), globalState);
        let _rhs349 = _module.__default.fm2(globalState);
        let _lhs221 = _2410_v62;
        let _lhs222 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2410_v62).length));
        let _lhs223 = p0;
        let _lhs224 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        let _lhs225 = p0;
        let _lhs226 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
        _lhs221[_lhs222] = _rhs346;
        _lhs223[_lhs224] = _rhs347;
        _2411_v63 = _rhs348;
        _lhs225[_lhs226] = _rhs349;
      } else {
        let _2417___mcc_h3 = (_source36).cf57;
        let _2418_cf57 = _2417___mcc_h3;
        if (false) {
          _2379_v38 = _2379_v38;
          let _2419_v69;
          _2419_v69 = _dafny.Seq.UnicodeFromString("hsy");
          let _2420_v70;
          _2420_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-635),_2419_v69);
          let _2421_v71;
          _2421_v71 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2422_v72;
          _2422_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v71,_this.f11);
          let _2423_v73;
          _2423_v73 = _dafny.Set.fromElements(p1, (_this).f19);
          _2420_v70 = (_2420_v70).update((_2329_v1).f25, (((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]) ? ((_this).fm8((_2329_v1).f25, _2422_v72, _2423_v73, globalState)) : (_2419_v69)));
          let _2424_v74;
          let _nw348 = new _module.C6();
          _nw348.__ctor();
          _2424_v74 = _nw348;
          let _2425_v75;
          let _nw349 = Array((new BigNumber(4)).toNumber()).fill([]);
          _2425_v75 = _nw349;
          let _2426_v76;
          _2426_v76 = _module.D4.create_DC12(_this.f11, new BigNumber(-251), (p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]);
          let _2427_v77;
          _2427_v77 = _dafny.Seq.of(_this.f11);
          let _2428_v78;
          _2428_v78 = _module.D7.create_DC23(_2425_v75, _2427_v77, (_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))], (p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]);
          _2425_v75 = (((_2426_v76).dtor_cf21) ? ((_2428_v78).dtor_cf42) : (_2425_v75));
          let _2429_v79;
          _2429_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2419_v69,(_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
          (_this).f11 = (_2429_v79).contains(_2419_v69);
        } else {
          let _index340 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
          (_2379_v38)[_index340] = (_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))];
          let _2430_v80;
          _2430_v80 = new _dafny.CodePoint('a'.codePointAt(0));
          let _2431_v81;
          _2431_v81 = _dafny.Map.Empty.slice().updateUnsafe((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))],false);
          let _2432_v82;
          _2432_v82 = _dafny.Set.fromElements(_2431_v81);
          let _2433_v83;
          _2433_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), ((_2434_v80) => function (_2435_i7) {
            return _2434_v80;
          })(_2430_v80))).length),_module.__default.fm2(globalState));
          let _2436_v84;
          _2436_v84 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_2432_v82).length)), (_2329_v1).f25), (_2329_v1).f25, _module.__default.fm2(globalState), new BigNumber((((!(_this.f11)) ? (_2433_v83) : (_2433_v83))).length), ((_this).f12).plus((_2329_v1).f25));
          let _2437_v85;
          _2437_v85 = _dafny.Map.Empty.slice().updateUnsafe((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))],_2430_v80);
          let _rhs350 = (((_2437_v85).contains((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))])) ? ((_2437_v85).get((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))])) : (_2430_v80));
          let _rhs351 = (_dafny.Set.fromElements((_this).f12, (_2329_v1).f25, p1, (p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))])).Union(_module.__default.fm50(globalState));
          _2430_v80 = _rhs350;
          _2436_v84 = _rhs351;
          let _2438_v86;
          _2438_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2431_v81).length),_this.f11);
          _2438_v86 = (_2438_v86).update(_module.__default.safeModuloInt(new BigNumber(-393), (_2329_v1).f25), _this.f11);
          let _2439_v87;
          _2439_v87 = _dafny.Set.fromElements((((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]) ? (p0) : (p0)), _2380_v39, p0, p0);
          _2439_v87 = _2439_v87;
          r1 = !(!((_this).fm13(_module.__default.fm1((_2329_v1).f25, (_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))], (_this).f13, globalState), globalState)));
        }
        if (!(_this.f11)) {
          let _2440_v88;
          let _nw350 = new _module.C6();
          _nw350.__ctor();
          _2440_v88 = _nw350;
          let _2441_v89;
          _2441_v89 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
          let _2442_v90;
          _2442_v90 = new _dafny.CodePoint('d'.codePointAt(0));
          let _2443_v91;
          _2443_v91 = _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), _module.__default.fm45(new BigNumber((_2441_v89).length), globalState), _module.__default.fm45((_2329_v1).f25, globalState), _2442_v90, _2442_v90);
          _2443_v91 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_2444_v90) => function (_2445_i8) {
            return _2444_v90;
          })(_2442_v90));
          let _2446_v92;
          _2446_v92 = _dafny.MultiSet.fromElements(_2442_v90);
          let _2447_v93;
          let _nw351 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2447_v93 = _nw351;
          let _2448_v94;
          _2448_v94 = _module.D7.create_DC22((((_2446_v92).contains(new _dafny.CodePoint('u'.codePointAt(0)))) ? ((_2446_v92).get(new _dafny.CodePoint('u'.codePointAt(0)))) : (new BigNumber(235))), _2447_v93);
          let _2449_v95;
          _2449_v95 = _module.D7.create_DC24(_2448_v94);
          _2449_v95 = _2449_v95;
          let _2450_v96;
          let _2451_v97;
          let _out39;
          let _out40;
          let _outcollector18 = _module.__default.m0(globalState);
          _out39 = _outcollector18[0];
          _out40 = _outcollector18[1];
          _2450_v96 = _out39;
          _2451_v97 = _out40;
          let _2452_v98;
          _2452_v98 = _dafny.Set.fromElements(_module.__default.fm58(new BigNumber(754), globalState), _module.D15.create_DC41(_2451_v97));
          let _2453_v99;
          _2453_v99 = _module.D17.create_DC44(_2452_v98);
          let _index341 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length));
          (_2379_v38)[_index341] = ((_2453_v99).dtor_cf68).IsSubsetOf(_2452_v98);
        } else {
          let _2454_v100;
          _2454_v100 = new _dafny.CodePoint('t'.codePointAt(0));
          let _2455_v101;
          _2455_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2454_v100),(p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]);
          let _2456_v103;
          _2456_v103 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2454_v100),_this.f11);
          r1 = (_2455_v101).equals((function () {
            let _coll95 = new _dafny.Map();
            for (const _compr_95 of (_2456_v103).Keys.Elements) {
              let _2457_v102 = _compr_95;
              if ((_2456_v103).contains(_2457_v102)) {
                _coll95.push([_2457_v102,(_2329_v1).f25]);
              }
            }
            return _coll95;
          }()).Merge(_2455_v101));
          let _2458_v104;
          _2458_v104 = _dafny.Map.Empty.slice().updateUnsafe((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))],p1);
          let _2459_v105;
          _2459_v105 = _dafny.Map.Empty.slice().updateUnsafe(_2458_v104,(_2329_v1).f25);
          let _2460_v106;
          _2460_v106 = _dafny.Seq.UnicodeFromString("xuhbaxora");
          _2459_v105 = (_2459_v105).update((_2458_v104).Merge(_2458_v104), ((_2329_v1).f25).plus(new BigNumber((_2460_v106).length)));
          let _2461_v107;
          let _nw352 = Array((new BigNumber(27)).toNumber());
          _nw352[(_dafny.ZERO).toNumber()] = p0;
          _nw352[(_dafny.ONE).toNumber()] = p0;
          _nw352[(new BigNumber(2)).toNumber()] = p0;
          _nw352[(new BigNumber(3)).toNumber()] = p0;
          _nw352[(new BigNumber(4)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(5)).toNumber()] = p0;
          _nw352[(new BigNumber(6)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(7)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(8)).toNumber()] = p0;
          _nw352[(new BigNumber(9)).toNumber()] = p0;
          _nw352[(new BigNumber(10)).toNumber()] = p0;
          _nw352[(new BigNumber(11)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(12)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(13)).toNumber()] = p0;
          _nw352[(new BigNumber(14)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(15)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(16)).toNumber()] = p0;
          _nw352[(new BigNumber(17)).toNumber()] = p0;
          _nw352[(new BigNumber(18)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(19)).toNumber()] = p0;
          _nw352[(new BigNumber(20)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(21)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(22)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(23)).toNumber()] = _2380_v39;
          _nw352[(new BigNumber(24)).toNumber()] = p0;
          _nw352[(new BigNumber(25)).toNumber()] = p0;
          _nw352[(new BigNumber(26)).toNumber()] = _2380_v39;
          _2461_v107 = _nw352;
          let _2462_v108;
          _2462_v108 = _dafny.Seq.of(!(true));
          (globalState).f8 = (_module.D7.create_DC23(_2461_v107, _2462_v108, true, (_2329_v1).f25)).dtor_cf44;
          let _2463_v109;
          let _nw353 = Array((new BigNumber(4)).toNumber());
          _nw353[(_dafny.ZERO).toNumber()] = _2379_v38;
          _nw353[(_dafny.ONE).toNumber()] = _2379_v38;
          _nw353[(new BigNumber(2)).toNumber()] = _2379_v38;
          _nw353[(new BigNumber(3)).toNumber()] = _2379_v38;
          _2463_v109 = _nw353;
          let _rhs352 = (_2329_v1).f25;
          let _rhs353 = _2463_v109;
          r0 = _rhs352;
          _2463_v109 = _rhs353;
          let _index342 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_2463_v109).length));
          (_2463_v109)[_index342] = _2379_v38;
          let _2464_v110;
          _2464_v110 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), ((_2465_v1) => function (_2466_i9) {
            return (_2465_v1).f25;
          })(_2329_v1)));
          let _index343 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
          let _index344 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_2463_v109).length));
          let _rhs354 = _this.f11;
          let _rhs355 = _this.f11;
          let _rhs356 = (_2464_v110).IsProperSubsetOf(_2464_v110);
          let _rhs357 = (((_2462_v108)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))], new BigNumber((_2462_v108).length))]) ? ((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]) : ((_2329_v1).f25));
          let _rhs358 = (((((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]) ? (_this.f11) : ((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]))) ? (_2379_v38) : (_2379_v38));
          let _lhs227 = _this;
          let _lhs228 = _this;
          let _lhs229 = globalState;
          let _lhs230 = p0;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length));
          let _lhs232 = _2463_v109;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_2463_v109).length));
          _lhs227.f11 = _rhs354;
          _lhs228.f11 = _rhs355;
          _lhs229.f8 = _rhs356;
          _lhs230[_lhs231] = _rhs357;
          _lhs232[_lhs233] = _rhs358;
        }
        let _2467_v111;
        _2467_v111 = _dafny.Map.Empty.slice().updateUnsafe((_2329_v1).f25,new BigNumber(719));
        let _2468_v112;
        _2468_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v111,(_2328_v0)[_module.__default.safeIndex((_2329_v1).f25, new BigNumber((_2328_v0).length))]);
        let _2469_v113;
        _2469_v113 = _dafny.Seq.of(_2467_v111, _2467_v111);
        _2468_v112 = (_dafny.Map.Empty.slice().updateUnsafe((_2469_v113)[_module.__default.safeIndex((_this).f13, new BigNumber((_2469_v113).length))],(p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))])).Merge((_2468_v112).Merge(_module.__default.fm59((_this).f13, (_2329_v1).f25, _this.f11, (_2329_v1).f25, globalState)));
        let _2470_v114;
        _2470_v114 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_this.f11),true);
        let _2471_v115;
        _2471_v115 = _dafny.Set.fromElements((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
        _2470_v114 = (_2470_v114).update((_2471_v115).Intersect(_dafny.Set.fromElements((_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))])), (_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))]);
      }
      (globalState).f8 = (_2379_v38)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_2379_v38).length))];
      r0 = (_dafny.ZERO).minus(((_this).f19).plus((p0)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((p0).length))]));
      let _2472_v116;
      _2472_v116 = _dafny.Seq.UnicodeFromString("iowpq");
      r1 = !_dafny.areEqual(_2472_v116, _dafny.Seq.UnicodeFromString("fvs"));
      return [r0, r1];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C13 = class C13 {
    constructor () {
      this._tname = "_module.C13";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Seq.of();
      let _2473_v0;
      _2473_v0 = new BigNumber(388);
      let _hi11 = _2473_v0;
      for (let _2474_i0 = (_dafny.ZERO).minus(_2473_v0); _2474_i0.isLessThan(_hi11); _2474_i0 = _2474_i0.plus(_dafny.ONE)) {
        let _2475_v1;
        _2475_v1 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2476_v2;
        _2476_v2 = _dafny.Seq.UnicodeFromString("d");
        r0 = _dafny.Seq.contains(_2476_v2, _2475_v1);
        r0 = !((new BigNumber(-350)).minus(_2474_i0)).isEqualTo(_module.__default.fm2(globalState));
        let _2477_v3;
        _2477_v3 = false;
        let _2478_v4;
        _2478_v4 = _dafny.MultiSet.fromElements(_2474_i0);
        let _2479_v5;
        _2479_v5 = _module.D0.create_DC1(_2474_i0, _2476_v2, _2474_i0, _2477_v3);
        (globalState).f8 = _module.__default.fm1(_2474_i0, !((_2477_v3) && (_2477_v3)), (((_2478_v4).contains((_2479_v5).dtor_cf3)) ? ((_2478_v4).get((_2479_v5).dtor_cf3)) : ((_dafny.ZERO).minus(_2474_i0))), globalState);
        _2476_v2 = _dafny.Seq.UnicodeFromString("dpaumjqq");
      }
      let _2480_v6;
      _2480_v6 = false;
      let _2481_v7;
      _2481_v7 = _module.D0.create_DC3(_2473_v0, _2480_v6);
      let _source37 = _2481_v7;
      if (_source37.is_DC1) {
        let _2482___mcc_h0 = (_source37).cf1;
        let _2483___mcc_h1 = (_source37).cf2;
        let _2484___mcc_h2 = (_source37).cf3;
        let _2485___mcc_h3 = (_source37).cf4;
        let _2486_cf4 = _2485___mcc_h3;
        let _2487_cf3 = _2484___mcc_h2;
        let _2488_cf2 = _2483___mcc_h1;
        let _2489_cf1 = _2482___mcc_h0;
        let _2490_v8;
        _2490_v8 = _dafny.Seq.of(_2473_v0);
        _2489_cf1 = (_2489_cf1).multipliedBy(new BigNumber((_2490_v8).length));
        let _2491_v9;
        _2491_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2489_cf1,_2480_v6);
        _2480_v6 = (((_2491_v9).contains(_module.__default.safeDivisionInt(_2489_cf1, _2487_cf3))) ? ((_2491_v9).get(_module.__default.safeDivisionInt(_2489_cf1, _2487_cf3))) : (!(_2480_v6)));
        if (_2486_cf4) {
          let _2492_v10;
          _2492_v10 = new _dafny.CodePoint('r'.codePointAt(0));
          let _rhs359 = _2480_v6;
          let _rhs360 = new _dafny.CodePoint('y'.codePointAt(0));
          let _rhs361 = _2489_cf1;
          let _lhs234 = globalState;
          _lhs234.f8 = _rhs359;
          _2492_v10 = _rhs360;
          _2473_v0 = _rhs361;
          let _2493_v11;
          _2493_v11 = _dafny.MultiSet.fromElements(_2480_v6, _2480_v6);
          let _2494_v12;
          _2494_v12 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((((_2493_v11).contains(_2486_cf4)) ? ((_2493_v11).get(_2486_cf4)) : (_module.__default.fm2(globalState))))), _2473_v0, new BigNumber(-530), new BigNumber((_2491_v9).length), _2487_cf3);
          let _2495_v13;
          _2495_v13 = _dafny.Set.fromElements(_2486_cf4, !(_2480_v6));
          let _2496_v14;
          _2496_v14 = _dafny.MultiSet.fromElements(_2495_v13, _2495_v13);
          let _2497_v15;
          let _init65 = ((_2498_v0) => function (_2499_i1) {
            return _module.__default.safeModuloInt(_2499_i1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jf")).length),_2498_v0)).length));
          })(_2473_v0);
          let _nw354 = Array((new BigNumber(4)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw354.length); _i0_65++) {
            _nw354[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2497_v15 = _nw354;
          let _2500_v16;
          _2500_v16 = _module.D0.create_DC2(new BigNumber((_2496_v14).cardinality()), _2497_v15);
          let _2501_v17;
          _2501_v17 = _dafny.Seq.of(_2480_v6, _2480_v6, _2480_v6);
          let _rhs362 = ((_2494_v12).Difference(_dafny.Set.fromElements(_2489_cf1))).Intersect(_2494_v12);
          let _rhs363 = (_2500_v16).dtor_cf5;
          let _rhs364 = _dafny.Seq.IsPrefixOf(_2501_v17, _2501_v17);
          let _lhs235 = globalState;
          _2494_v12 = _rhs362;
          r2 = _rhs363;
          _lhs235.f8 = _rhs364;
          let _2502_v18;
          let _nw355 = Array((new BigNumber(8)).toNumber()).fill(false);
          _2502_v18 = _nw355;
          let _index345 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_2502_v18).length));
          (_2502_v18)[_index345] = _2486_cf4;
          let _index346 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_2502_v18).length));
          (_2502_v18)[_index346] = _2480_v6;
          _2480_v6 = (_2502_v18)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2502_v18).length))];
          let _2503_v19;
          _2503_v19 = _module.D0.create_DC1((_dafny.ZERO).minus(_2473_v0), _dafny.Seq.UnicodeFromString("igwwlgl"), _2489_cf1, _2486_cf4);
          r0 = _dafny.Seq.IsProperPrefixOf((_2503_v19).dtor_cf2, _2488_cf2);
        } else {
          let _2504_v20;
          let _nw356 = new _module.C3();
          _nw356.__ctor(_2487_cf3, new BigNumber(-62));
          _2504_v20 = _nw356;
          let _2505_v21;
          let _init66 = ((_2506_v0) => function (_2507_i2) {
            return (_2507_i2).multipliedBy(_2506_v0);
          })(_2473_v0);
          let _nw357 = Array((new BigNumber(29)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw357.length); _i0_66++) {
            _nw357[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2505_v21 = _nw357;
          let _index347 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2505_v21).length));
          (_2505_v21)[_index347] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("qpsgkquyo")).length), _2487_cf3);
          let _index348 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2505_v21).length));
          (_2505_v21)[_index348] = _2487_cf3;
          let _2508_v22;
          let _nw358 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Set.Empty);
          _2508_v22 = _nw358;
          _2508_v22 = _2508_v22;
          _2486_cf4 = _2480_v6;
          let _2509_v23;
          _2509_v23 = _dafny.Set.fromElements(new BigNumber((_2488_cf2).length));
          let _2510_v24;
          _2510_v24 = _module.D17.create_DC45((_dafny.ZERO).minus(new BigNumber((_2490_v8).length)), _2509_v23);
          r2 = (_2510_v24).dtor_cf69;
        }
        let _2511_v25;
        let _nw359 = Array((new BigNumber(28)).toNumber()).fill(_module.D7.Default());
        _2511_v25 = _nw359;
        let _2512_v26;
        _2512_v26 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2513_v27;
        _2513_v27 = _dafny.MultiSet.fromElements(_2486_cf4, _2486_cf4, true);
        let _2514_v28;
        _2514_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2512_v26),_2513_v27);
        let _2515_v29;
        _2515_v29 = _module.D9.create_DC26(_2512_v26);
        let _2516_v30;
        _2516_v30 = _dafny.Set.fromElements((_2515_v29).dtor_cf48, _2512_v26);
        let _2517_v31;
        _2517_v31 = _module.D7.create_DC24(_module.D7.create_DC21((((_2514_v28).contains(_2516_v30)) ? ((_2514_v28).get(_2516_v30)) : (_2513_v27))));
        let _index349 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2511_v25).length));
        (_2511_v25)[_index349] = _module.D7.create_DC24(_module.D7.create_DC24((_2517_v31).dtor_cf46));
        let _index350 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2511_v25).length));
        (_2511_v25)[_index350] = _2517_v31;
      } else if (_source37.is_DC2) {
        let _2518___mcc_h4 = (_source37).cf5;
        let _2519___mcc_h5 = (_source37).cf6;
        let _2520_cf6 = _2519___mcc_h5;
        let _2521_cf5 = _2518___mcc_h4;
        let _2522_v32;
        _2522_v32 = _module.D6.create_DC18(_2480_v6);
        _2522_v32 = _2522_v32;
        let _2523_v33;
        _2523_v33 = _dafny.Set.fromElements(_2480_v6);
        let _2524_v34;
        _2524_v34 = new _dafny.CodePoint('c'.codePointAt(0));
        _2523_v33 = (_module.__default.fm4(new BigNumber(992), _2481_v7, _2521_cf5, _2524_v34, globalState)).Union(_dafny.Set.fromElements(_2480_v6, _2480_v6, !(_2480_v6), _2480_v6, _module.__default.fm1(_2473_v0, _2480_v6, _2473_v0, globalState)));
        let _2525_v35;
        _2525_v35 = _dafny.Seq.of((_dafny.ZERO).minus(_2521_cf5));
        let _2526_v36;
        _2526_v36 = _2525_v35;
        let _2527_v37;
        _2527_v37 = _dafny.Set.fromElements(_2520_cf6);
        let _2528_v38;
        _2528_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2526_v36,new BigNumber((_2527_v37).length));
        let _2529_v40;
        _2529_v40 = _dafny.Seq.UnicodeFromString("wqem");
        let _2530_v41;
        _2530_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2529_v40,_2521_cf5);
        let _2531_v42;
        _2531_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_2473_v0, _2480_v6, _2473_v0, globalState),_2530_v41);
        let _2532_v43;
        _2532_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2480_v6,_2480_v6);
        let _2533_v44;
        let _nw360 = new _module.C9();
        _nw360.__ctor(function () {
          let _coll96 = new _dafny.Set();
          for (const _compr_96 of (_2528_v38).Keys.Elements) {
            let _2534_v39 = _compr_96;
            if ((_2528_v38).contains(_2534_v39)) {
              _coll96.add(_2534_v39);
            }
          }
          return _coll96;
        }(), (new BigNumber(((((_2531_v42).contains(!(_2480_v6))) ? ((_2531_v42).get(!(_2480_v6))) : (_2530_v41))).length)).multipliedBy(new BigNumber((_2532_v43).length)), _2473_v0, _2521_cf5, !_dafny.areEqual(_dafny.Seq.UnicodeFromString("v"), _2529_v40), _dafny.Seq.UnicodeFromString("rxog"));
        _2533_v44 = _nw360;
        let _index351 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2520_cf6).length));
        (_2520_cf6)[_index351] = new BigNumber(804);
        let _index352 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2520_cf6).length));
        let _rhs365 = _2533_v44;
        let _rhs366 = ((_2480_v6) ? (_2520_cf6) : (((true) ? (_2520_cf6) : (_2520_cf6))));
        let _rhs367 = _2480_v6;
        let _rhs368 = ((_module.__default.fm2(globalState)).multipliedBy(_2473_v0)).plus(_2473_v0);
        let _rhs369 = _2480_v6;
        let _lhs236 = globalState;
        let _lhs237 = _2520_cf6;
        let _lhs238 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2520_cf6).length));
        _2533_v44 = _rhs365;
        _2520_cf6 = _rhs366;
        _lhs236.f8 = _rhs367;
        _lhs237[_lhs238] = _rhs368;
        _2480_v6 = _rhs369;
        _2480_v6 = _2480_v6;
      } else if (_source37.is_DC3) {
        let _2535___mcc_h6 = (_source37).cf7;
        let _2536___mcc_h7 = (_source37).cf8;
        let _2537_cf8 = _2536___mcc_h7;
        let _2538_cf7 = _2535___mcc_h6;
        let _2539_v45;
        _2539_v45 = _dafny.Seq.of(_2480_v6, _2537_cf8, _2537_cf8);
        r3 = _dafny.Seq.Concat(_2539_v45, _2539_v45);
        let _2540_v46;
        _2540_v46 = _dafny.Seq.of((_dafny.ZERO).minus(_2473_v0));
        let _2541_v48;
        _2541_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2540_v46,function () {
          let _coll97 = new _dafny.Map();
          for (const _compr_97 of _dafny.IntegerRange(new BigNumber(407), new BigNumber(236))) {
            let _2542_v47 = _compr_97;
            if (((new BigNumber(407)).isLessThanOrEqualTo(_2542_v47)) && ((_2542_v47).isLessThan(new BigNumber(236)))) {
              _coll97.push([(_2542_v47).minus(_2473_v0),_2537_cf8]);
            }
          }
          return _coll97;
        }());
        let _2543_v49;
        _2543_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(453),_2480_v6);
        let _2544_v50;
        let _nw361 = Array((new BigNumber(3)).toNumber());
        _nw361[(_dafny.ZERO).toNumber()] = _2541_v48;
        _nw361[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(617)),(_2543_v49).update(_2473_v0, _2480_v6));
        _nw361[(new BigNumber(2)).toNumber()] = _2541_v48;
        _2544_v50 = _nw361;
        let _index353 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2544_v50).length));
        (_2544_v50)[_index353] = (_2541_v48).update(_2540_v46, (_2543_v49).update(new BigNumber(-790), true));
        let _2545_v51;
        _2545_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2537_cf8,_2473_v0);
        let _2546_v52;
        _2546_v52 = _dafny.Set.fromElements(_2543_v49);
        let _2547_v53;
        _2547_v53 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2548_v54;
        _2548_v54 = _dafny.Map.Empty.slice().updateUnsafe(_2547_v53,_2538_cf7);
        let _2549_v55;
        _2549_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2537_cf8,true);
        let _index354 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2544_v50).length));
        let _rhs370 = (((_2480_v6) ? (_2541_v48) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(57)), ((_2550_cf7) => function (_2551_i3) {
          return _2550_cf7;
        })(_2538_cf7)),_module.__default.fm47(true, _2537_cf8, _2545_v51, _2473_v0, globalState))))).update(_2540_v46, _module.__default.fm47(true, _2537_cf8, _dafny.Map.Empty.slice().updateUnsafe(_2537_cf8,new BigNumber((_2546_v52).length)), _2538_cf7, globalState));
        let _rhs371 = _dafny.Seq.contains(_dafny.Seq.Concat(_2539_v45, _2539_v45), _2480_v6);
        let _rhs372 = _2480_v6;
        let _rhs373 = (_module.__default.safeModuloInt(_2473_v0, _2538_cf7)).multipliedBy((((_2548_v54).contains(_2547_v53)) ? ((_2548_v54).get(_2547_v53)) : (_2538_cf7)));
        let _rhs374 = (_2549_v55).Merge((_2549_v55).Merge(_2549_v55));
        let _lhs239 = _2544_v50;
        let _lhs240 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2544_v50).length));
        let _lhs241 = globalState;
        _lhs239[_lhs240] = _rhs370;
        _lhs241.f8 = _rhs371;
        _2480_v6 = _rhs372;
        _2538_cf7 = _rhs373;
        r1 = _rhs374;
        if (_2480_v6) {
          let _2552_v56;
          let _nw362 = new _module.C1();
          _nw362.__ctor(_2473_v0);
          _2552_v56 = _nw362;
          let _2553_v57;
          let _nw363 = Array((new BigNumber(19)).toNumber()).fill(false);
          _2553_v57 = _nw363;
          let _2554_v58;
          _2554_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2553_v57,_2480_v6);
          let _2555_v59;
          _2555_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2540_v46,(_2554_v58).equals(_2554_v58));
          _2555_v59 = (_2555_v59).update(_dafny.Seq.update(_dafny.Seq.of(_2552_v56.f27, _2473_v0), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_2552_v56.f27)).length), new BigNumber((_dafny.Seq.of(_2552_v56.f27, _2473_v0)).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(901)), ((_2556_v53) => function (_2557_i4) {
            return _2556_v53;
          })(_2547_v53))).length)), _dafny.areEqual(_2539_v45, _2539_v45));
          (globalState).f8 = (_2539_v45)[_module.__default.safeIndex(_2538_cf7, new BigNumber((_2539_v45).length))];
          _2480_v6 = _2537_cf8;
          _2547_v53 = new _dafny.CodePoint('n'.codePointAt(0));
        } else {
          r2 = _2538_cf7;
          let _2558_v60;
          _2558_v60 = _dafny.Set.fromElements(false, _2537_cf8);
          r2 = (_2473_v0).plus(new BigNumber((_2558_v60).length));
          let _2559_v61;
          _2559_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_2547_v53);
          let _2560_v62;
          _2560_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2559_v61,_module.__default.fm2(globalState));
          _2560_v62 = (_2560_v62).update((_2559_v61).Merge(function () {
            let _coll98 = new _dafny.Map();
            for (const _compr_98 of (_2543_v49).Keys.Elements) {
              let _2561_v63 = _compr_98;
              if ((_2543_v49).contains(_2561_v63)) {
                _coll98.push([(_2561_v63).minus((_2540_v46)[_module.__default.safeIndex(_2473_v0, new BigNumber((_2540_v46).length))]),_2547_v53]);
              }
            }
            return _coll98;
          }()), (_dafny.ZERO).minus(_2473_v0));
          let _2562_v64;
          _2562_v64 = _dafny.Seq.UnicodeFromString("nivaav");
          let _2563_v65;
          let _nw364 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2563_v65 = _nw364;
          let _2564_v66;
          _2564_v66 = _module.D7.create_DC23(_2563_v65, _2539_v45, _2537_cf8, _2473_v0);
          let _2565_v67;
          _2565_v67 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_2537_cf8) ? (_2473_v0) : (new BigNumber((_2562_v64).length)))),_2564_v66);
          _2565_v67 = (_2565_v67).update(_2538_cf7, _2564_v66);
          let _2566_v68;
          let _nw365 = Array((new BigNumber(22)).toNumber()).fill(false);
          _2566_v68 = _nw365;
          _2566_v68 = _2566_v68;
        }
        let _2567_v69;
        let _init67 = ((_2568_v45, _2569_v6) => function (_2570_i5) {
          return _module.D15.create_DC40(_dafny.Map.Empty.slice().updateUnsafe(_2568_v45,_module.D7.create_DC21(_dafny.MultiSet.fromElements(_2569_v6))));
        })(_2539_v45, _2480_v6);
        let _nw366 = Array((new BigNumber(7)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw366.length); _i0_67++) {
          _nw366[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2567_v69 = _nw366;
        let _2571_v70;
        _2571_v70 = _module.D7.create_DC21(_module.__default.fm30(_2480_v6, _2480_v6, _2538_cf7, globalState));
        let _2572_v71;
        _2572_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2539_v45,_2571_v70);
        let _2573_v72;
        _2573_v72 = _module.D15.create_DC40(_2572_v71);
        let _index355 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2567_v69).length));
        (_2567_v69)[_index355] = _2573_v72;
        let _pat_let_tv90 = _2538_cf7;
        let _pat_let_tv91 = _2473_v0;
        let _pat_let_tv92 = globalState;
        let _index356 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2567_v69).length));
        (_2567_v69)[_index356] = function (_pat_let55_0) {
          return function (_2574_dt__update__tmp_h0) {
            return function (_pat_let56_0) {
              return function (_2575_dt__update_hcf64_h0) {
                return _module.D15.create_DC40(_2575_dt__update_hcf64_h0);
              }(_pat_let56_0);
            }(_module.__default.fm55(new BigNumber(488), _pat_let_tv90, _pat_let_tv91, _pat_let_tv92));
          }(_pat_let55_0);
        }(_2573_v72);
      } else {
        let _2576___mcc_h8 = (_source37).cf0;
        let _2577_cf0 = _2576___mcc_h8;
        let _2578_v73;
        _2578_v73 = _dafny.MultiSet.fromElements((_2473_v0).plus(_2473_v0), ((!(!(_2480_v6))) ? (_2473_v0) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_2473_v0)).length))), (_dafny.ZERO).minus(_2473_v0));
        let _2579_v74;
        _2579_v74 = _dafny.MultiSet.fromElements(_2480_v6, true);
        _2578_v73 = (((_2579_v74).equals((_2579_v74).update(_2480_v6, _module.__default.abs(_2473_v0)))) ? ((_2578_v73).Difference(_2578_v73)) : ((_2578_v73).Intersect((_2578_v73).update(_2473_v0, _module.__default.abs(new BigNumber(84))))));
        let _2580_v75;
        let _2581_v76;
        let _out41;
        let _out42;
        let _outcollector19 = _module.__default.m0(globalState);
        _out41 = _outcollector19[0];
        _out42 = _outcollector19[1];
        _2580_v75 = _out41;
        _2581_v76 = _out42;
        (globalState).f8 = !((_module.__default.fm2(globalState)).minus(_2473_v0)).isEqualTo(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ja")).length))).multipliedBy(_2473_v0));
        let _2582_v77;
        _2582_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2581_v76,_2480_v6);
        let _2583_v78;
        let _nw367 = Array((new BigNumber(26)).toNumber());
        _nw367[(_dafny.ZERO).toNumber()] = _2473_v0;
        _nw367[(_dafny.ONE).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(2)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(3)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(4)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(5)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(6)).toNumber()] = new BigNumber((_2582_v77).length);
        _nw367[(new BigNumber(7)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(8)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(9)).toNumber()] = new BigNumber((_2582_v77).length);
        _nw367[(new BigNumber(10)).toNumber()] = new BigNumber(718);
        _nw367[(new BigNumber(11)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(12)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(13)).toNumber()] = new BigNumber(329);
        _nw367[(new BigNumber(14)).toNumber()] = new BigNumber(895);
        _nw367[(new BigNumber(15)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(16)).toNumber()] = new BigNumber(825);
        _nw367[(new BigNumber(17)).toNumber()] = new BigNumber(-284);
        _nw367[(new BigNumber(18)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(19)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(20)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(21)).toNumber()] = _module.__default.fm2(globalState);
        _nw367[(new BigNumber(22)).toNumber()] = _2473_v0;
        _nw367[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), function (_2584_i6) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length);
        _nw367[(new BigNumber(24)).toNumber()] = new BigNumber(468);
        _nw367[(new BigNumber(25)).toNumber()] = new BigNumber(115);
        _2583_v78 = _nw367;
        let _2585_v79;
        let _nw368 = Array((_dafny.ONE).toNumber()).fill([]);
        _2585_v79 = _nw368;
        let _2586_v80;
        _2586_v80 = _dafny.Seq.UnicodeFromString("gka");
        let _2587_v81;
        _2587_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2586_v80,_2480_v6);
        let _2588_v82;
        _2588_v82 = _dafny.Seq.of(_2581_v76, (((_2587_v81).contains(_2586_v80)) ? ((_2587_v81).get(_2586_v80)) : (_2581_v76)));
        let _2589_v83;
        _2589_v83 = _module.D7.create_DC23(_2585_v79, _2588_v82, _2581_v76, _2473_v0);
        let _2590_v84;
        _2590_v84 = _dafny.Map.Empty.slice().updateUnsafe(_2583_v78,_2589_v83);
        let _2591_v85;
        _2591_v85 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_2588_v82);
        let _2592_v86;
        _2592_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2591_v85,_2581_v76);
        let _2593_v87;
        _2593_v87 = _dafny.Seq.of(new BigNumber((_2582_v77).length));
        let _2594_v88;
        _2594_v88 = _dafny.MultiSet.fromElements(_2593_v87);
        let _rhs375 = (new BigNumber((_2590_v84).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), function (_2595_i7) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })).length));
        let _rhs376 = (_2473_v0).isLessThan((((_2579_v74).contains(_2480_v6)) ? ((_2579_v74).get(_2480_v6)) : ((_dafny.ZERO).minus(_2473_v0))));
        let _rhs377 = _2480_v6;
        let _rhs378 = !((((_2592_v86).contains(_2591_v85)) ? ((_2592_v86).get(_2591_v85)) : (!((_2594_v88).IsDisjointFrom(_dafny.MultiSet.fromElements(_2593_v87))))));
        let _rhs379 = _2581_v76;
        let _lhs242 = globalState;
        r2 = _rhs375;
        r0 = _rhs376;
        _lhs242.f8 = _rhs377;
        _2581_v76 = _rhs378;
        _2581_v76 = _rhs379;
      }
      let _2596_i8;
      _2596_i8 = _dafny.ZERO;
      L11: {
        while (((((_2480_v6) ? (_2480_v6) : (true))) ? (_2480_v6) : (_2480_v6))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2596_i8)) {
              break L11;
            }
            _2596_i8 = (_2596_i8).plus(_dafny.ONE);
            let _2597_v89;
            _2597_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2480_v6,_2480_v6);
            let _2598_v90;
            _2598_v90 = _dafny.Set.fromElements(_2480_v6);
            let _2599_v91;
            let _nw369 = new _module.C11();
            _nw369.__ctor(((_2480_v6) ? ((_dafny.ZERO).minus(_2473_v0)) : (new BigNumber((_2597_v89).length))), new BigNumber((_2598_v90).length));
            _2599_v91 = _nw369;
            let _2600_v92;
            _2600_v92 = _dafny.Seq.UnicodeFromString("jf");
            let _2601_v93;
            _2601_v93 = new _dafny.CodePoint('u'.codePointAt(0));
            let _2602_v94;
            let _nw370 = new _module.C2();
            _nw370.__ctor(_2473_v0, _2600_v92, _2473_v0, new BigNumber((_module.__default.fm4(_2473_v0, _2481_v7, _2473_v0, _2601_v93, globalState)).length), _2480_v6);
            _2602_v94 = _nw370;
            let _2603_v95;
            _2603_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2602_v94,!(_2480_v6) || (_2480_v6));
            _2603_v95 = (_2603_v95).update(_2602_v94, _2480_v6);
            let _2604_v96;
            _2604_v96 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_2480_v6);
            let _2605_v97;
            _2605_v97 = _dafny.Map.Empty.slice().updateUnsafe((((_2604_v96).contains(_2473_v0)) ? ((_2604_v96).get(_2473_v0)) : (_2480_v6)),new BigNumber(831));
            let _2606_v98;
            _2606_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2480_v6,(new BigNumber((_2605_v97).length)).multipliedBy(_2473_v0));
            _2606_v98 = _2606_v98;
            _2473_v0 = (_2473_v0).multipliedBy((_2602_v94).f26);
          }
        }
      }
      let _2607_v99;
      let _init68 = ((_2608_v0) => function (_2609_i9) {
        return (_2609_i9).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2608_v0,_2608_v0)).length));
      })(_2473_v0);
      let _nw371 = Array((new BigNumber(20)).toNumber());
      for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw371.length); _i0_68++) {
        _nw371[_i0_68] = _init68(new BigNumber(_i0_68));
      }
      _2607_v99 = _nw371;
      let _2610_v100;
      let _nw372 = Array((new BigNumber(26)).toNumber());
      _nw372[(_dafny.ZERO).toNumber()] = _2607_v99;
      _nw372[(_dafny.ONE).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(2)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(3)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(4)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(5)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(6)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(7)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(8)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(9)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(10)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(11)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(12)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(13)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(14)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(15)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(16)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(17)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(18)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(19)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(20)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(21)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(22)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(23)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(24)).toNumber()] = _2607_v99;
      _nw372[(new BigNumber(25)).toNumber()] = _2607_v99;
      _2610_v100 = _nw372;
      let _2611_v101;
      _2611_v101 = _dafny.Seq.of(_2480_v6, _2480_v6, _2480_v6);
      let _2612_v102;
      _2612_v102 = _module.D7.create_DC23(_2610_v100, _dafny.Seq.Concat(_2611_v101, _2611_v101), _2480_v6, _2473_v0);
      let _source38 = _2612_v102;
      if (_source38.is_DC22) {
        let _2613___mcc_h9 = (_source38).cf40;
        let _2614___mcc_h10 = (_source38).cf41;
        let _2615_cf41 = _2614___mcc_h10;
        let _2616_cf40 = _2613___mcc_h9;
        let _2617_v103;
        _2617_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2607_v99,((_2480_v6) ? (_2473_v0) : (_2473_v0)));
        let _2618_v104;
        _2618_v104 = _module.D15.create_DC41(_2480_v6);
        let _2619_v105;
        _2619_v105 = _module.D15.create_DC42(_2618_v104);
        let _2620_v106;
        _2620_v106 = _module.D15.create_DC42(_2619_v105);
        let _2621_v107;
        _2621_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2620_v106,_2473_v0);
        _2617_v103 = (_2617_v103).update(_2607_v99, (_dafny.ZERO).minus(new BigNumber((_2621_v107).length)));
        if (!(((_2480_v6) ? (_2480_v6) : (_2480_v6))) || (!((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_module.__default.fm1(_2473_v0, _2480_v6, _2473_v0, globalState))).length)).isEqualTo(_2616_cf40)))) {
          let _2622_v108;
          let _nw373 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
          _2622_v108 = _nw373;
          let _index357 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2622_v108).length));
          (_2622_v108)[_index357] = _dafny.Seq.Concat(_2611_v101, _dafny.Seq.update(_dafny.Seq.of(_2480_v6), _module.__default.safeIndex(_2473_v0, new BigNumber((_dafny.Seq.of(_2480_v6)).length)), _2480_v6));
          let _2623_v109;
          _2623_v109 = _dafny.Seq.UnicodeFromString("oevt");
          let _2624_v110;
          _2624_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2480_v6,_2480_v6);
          let _2625_v111;
          _2625_v111 = _dafny.Seq.of(_2624_v110, _2624_v110);
          let _2626_v112;
          _2626_v112 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_2623_v109, _2623_v109)).length), new BigNumber(798), new BigNumber(((_2625_v111)[_module.__default.safeIndex(_2473_v0, new BigNumber((_2625_v111).length))]).length));
          let _index358 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2622_v108).length));
          let _rhs380 = ((_2473_v0).isLessThan(_2616_cf40)) || (_2480_v6);
          let _rhs381 = _2626_v112;
          let _rhs382 = !(_2480_v6) || (((_2480_v6) ? (_2480_v6) : (_2480_v6)));
          let _rhs383 = _2611_v101;
          let _lhs243 = globalState;
          let _lhs244 = globalState;
          let _lhs245 = globalState;
          let _lhs246 = _2622_v108;
          let _lhs247 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2622_v108).length));
          _lhs243.f8 = _rhs380;
          _lhs244.f5 = _rhs381;
          _lhs245.f8 = _rhs382;
          _lhs246[_lhs247] = _rhs383;
          _2480_v6 = _2480_v6;
          r2 = _2616_cf40;
          let _2627_v113;
          let _nw374 = Array((new BigNumber(2)).toNumber()).fill(false);
          _2627_v113 = _nw374;
          _2627_v113 = _2627_v113;
          let _index359 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index359] = _2473_v0;
          let _2628_v114;
          _2628_v114 = _dafny.MultiSet.fromElements(_2480_v6, _2480_v6, _2480_v6);
          let _2629_v115;
          _2629_v115 = _dafny.Seq.of(_2628_v114, (_2628_v114).update(true, _module.__default.abs(_2616_cf40)), _dafny.MultiSet.FromArray(_2611_v101));
          let _index360 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index360] = new BigNumber((_2629_v115).length);
        } else {
          let _index361 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index361] = _2473_v0;
          let _index362 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index362] = _2616_cf40;
          let _2630_v116;
          _2630_v116 = _dafny.Seq.of(new BigNumber(924));
          let _2631_v117;
          let _nw375 = Array((new BigNumber(12)).toNumber());
          _nw375[(_dafny.ZERO).toNumber()] = _2480_v6;
          _nw375[(_dafny.ONE).toNumber()] = ((_2480_v6) ? (_2480_v6) : (_2480_v6));
          _nw375[(new BigNumber(2)).toNumber()] = !((_2480_v6) === (false));
          _nw375[(new BigNumber(3)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(4)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(5)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(6)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(7)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(8)).toNumber()] = _module.__default.fm1(_2473_v0, _2480_v6, _2473_v0, globalState);
          _nw375[(new BigNumber(9)).toNumber()] = _2480_v6;
          _nw375[(new BigNumber(10)).toNumber()] = _module.__default.fm1(_2616_cf40, _2480_v6, new BigNumber((_2630_v116).length), globalState);
          _nw375[(new BigNumber(11)).toNumber()] = (_2616_cf40).isLessThanOrEqualTo(_2473_v0);
          _2631_v117 = _nw375;
          let _index363 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2631_v117).length));
          (_2631_v117)[_index363] = _2480_v6;
          let _2632_v118;
          _2632_v118 = _module.D9.create_DC26(new _dafny.CodePoint('p'.codePointAt(0)));
          let _2633_v119;
          _2633_v119 = _dafny.Map.Empty.slice().updateUnsafe((_2632_v118).dtor_cf48,false);
          let _2634_v120;
          _2634_v120 = new _dafny.CodePoint('l'.codePointAt(0));
          let _2635_v121;
          let _nw376 = Array((new BigNumber(13)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = (_2633_v119).Merge(_2633_v119);
          _nw376[(_dafny.ONE).toNumber()] = (_2633_v119).update(_2634_v120, !(_2480_v6));
          _nw376[(new BigNumber(2)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(3)).toNumber()] = (_module.D18.create_DC47(_2633_v119)).dtor_cf72;
          _nw376[(new BigNumber(4)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(5)).toNumber()] = _module.__default.fm60(_2480_v6, globalState);
          _nw376[(new BigNumber(6)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(7)).toNumber()] = _module.__default.fm60(_2480_v6, globalState);
          _nw376[(new BigNumber(8)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(9)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(10)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(11)).toNumber()] = _2633_v119;
          _nw376[(new BigNumber(12)).toNumber()] = _2633_v119;
          _2635_v121 = _nw376;
          let _index364 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_2607_v99).length));
          let _index365 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2607_v99).length));
          let _index366 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2631_v117).length));
          let _rhs384 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_2616_cf40), _2616_cf40)).minus((_dafny.ZERO).minus(_2616_cf40));
          let _rhs385 = _2473_v0;
          let _rhs386 = _2480_v6;
          let _rhs387 = _2635_v121;
          let _lhs248 = _2607_v99;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_2607_v99).length));
          let _lhs250 = _2607_v99;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2607_v99).length));
          let _lhs252 = _2631_v117;
          let _lhs253 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_2631_v117).length));
          _lhs248[_lhs249] = _rhs384;
          _lhs250[_lhs251] = _rhs385;
          _lhs252[_lhs253] = _rhs386;
          _2635_v121 = _rhs387;
          let _2636_v122;
          _2636_v122 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_2473_v0, _2616_cf40),_2473_v0);
          _2636_v122 = (_2636_v122).update(new BigNumber(118), _module.__default.safeDivisionInt(new BigNumber(-368), _2473_v0));
          (globalState).f8 = ((_2631_v117)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2631_v117).length))]) && (false);
          let _2637_v123;
          _2637_v123 = _dafny.MultiSet.fromElements(_2616_cf40, new BigNumber((_2636_v122).length));
          let _2638_v124;
          _2638_v124 = _dafny.MultiSet.fromElements(_2637_v123);
          let _2639_v125;
          _2639_v125 = _dafny.Seq.UnicodeFromString("qm");
          let _2640_v126;
          let _nw377 = new _module.C2();
          _nw377.__ctor(new BigNumber((_2638_v124).cardinality()), _2639_v125, _2616_cf40, (_2607_v99)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_2607_v99).length))], (_2631_v117)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_2631_v117).length))]);
          _2640_v126 = _nw377;
          let _2641_v127;
          _2641_v127 = _dafny.MultiSet.fromElements(((_2640_v126).f26).isLessThan((_2607_v99)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_2607_v99).length))]));
          _2641_v127 = _2641_v127;
        }
        let _index367 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_2610_v100).length));
        (_2610_v100)[_index367] = _2607_v99;
        let _index368 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_2610_v100).length));
        (_2610_v100)[_index368] = ((!(_2480_v6)) ? (_2607_v99) : (_2607_v99));
        let _2642_v128;
        _2642_v128 = _dafny.Seq.UnicodeFromString("fqbfsoug");
        _2642_v128 = (_module.D0.create_DC1(_module.__default.fm2(globalState), _2642_v128, _2473_v0, _2480_v6)).dtor_cf2;
      } else if (_source38.is_DC23) {
        let _2643___mcc_h11 = (_source38).cf42;
        let _2644___mcc_h12 = (_source38).cf43;
        let _2645___mcc_h13 = (_source38).cf44;
        let _2646___mcc_h14 = (_source38).cf45;
        let _2647_cf45 = _2646___mcc_h14;
        let _2648_cf44 = _2645___mcc_h13;
        let _2649_cf43 = _2644___mcc_h12;
        let _2650_cf42 = _2643___mcc_h11;
        let _2651_v129;
        _2651_v129 = _dafny.Seq.UnicodeFromString("adgdyssy");
        let _2652_v130;
        _2652_v130 = _dafny.Seq.of(new BigNumber(-629));
        let _2653_v131;
        _2653_v131 = _dafny.MultiSet.fromElements(_2473_v0);
        let _2654_v132;
        _2654_v132 = _dafny.Seq.of(_2647_cf45, (_2652_v130)[_module.__default.safeIndex(_2473_v0, new BigNumber((_2652_v130).length))], _2647_cf45, new BigNumber((_2611_v101).length), (((_2653_v131).contains(_2473_v0)) ? ((_2653_v131).get(_2473_v0)) : (_2473_v0)));
        let _2655_v133;
        let _nw378 = new _module.C2();
        _nw378.__ctor((_2473_v0).plus(_2647_cf45), _2651_v129, new BigNumber((_2652_v130).length), _module.__default.safeModuloInt((_2652_v130)[_module.__default.safeIndex(_2647_cf45, new BigNumber((_2652_v130).length))], _2647_cf45), _2648_cf44);
        _2655_v133 = _nw378;
        let _2656_v134;
        _2656_v134 = _dafny.Map.Empty.slice().updateUnsafe((_2655_v133).f26,_2651_v129);
        _2656_v134 = (_2656_v134).update((_2655_v133).f26, _2651_v129);
        if (_2480_v6) {
          _2648_cf44 = _2480_v6;
          let _2657_v135;
          _2657_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(749),_2648_cf44);
          let _2658_v136;
          _2658_v136 = _module.D0.create_DC1(_2647_cf45, _2651_v129, (_2655_v133).f26, _2648_cf44);
          let _2659_v137;
          _2659_v137 = new _dafny.CodePoint('n'.codePointAt(0));
          let _2660_v138;
          _2660_v138 = _dafny.Set.fromElements(_2473_v0);
          let _2661_v139;
          let _nw379 = Array((new BigNumber(16)).toNumber()).fill(_module.D0.Default());
          _2661_v139 = _nw379;
          let _2662_v140;
          _2662_v140 = _dafny.Set.fromElements(_2661_v139);
          let _2663_v141;
          _2663_v141 = _dafny.Map.Empty.slice().updateUnsafe(_2648_cf44,new BigNumber((_2662_v140).length));
          let _pat_let_tv93 = _2651_v129;
          let _pat_let_tv94 = _2651_v129;
          let _pat_let_tv95 = globalState;
          let _pat_let_tv96 = globalState;
          let _pat_let_tv97 = _2659_v137;
          let _pat_let_tv98 = globalState;
          let _pat_let_tv99 = _2473_v0;
          let _pat_let_tv100 = _2648_cf44;
          let _2664_v142;
          let _nw380 = Array((new BigNumber(27)).toNumber());
          _nw380[(_dafny.ZERO).toNumber()] = _module.D0.create_DC1((_2652_v130)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_2652_v130).length))], _dafny.Seq.UnicodeFromString("oabhtcmf"), new BigNumber((_2657_v135).length), _2648_cf44);
          _nw380[(_dafny.ONE).toNumber()] = function (_pat_let57_0) {
            return function (_2665_dt__update__tmp_h1) {
              return function (_pat_let58_0) {
                return function (_2666_dt__update_hcf3_h0) {
                  return function (_pat_let59_0) {
                    return function (_2667_dt__update_hcf2_h0) {
                      return _module.D0.create_DC1((_2665_dt__update__tmp_h1).dtor_cf1, _2667_dt__update_hcf2_h0, _2666_dt__update_hcf3_h0, (_2665_dt__update__tmp_h1).dtor_cf4);
                    }(_pat_let59_0);
                  }(_pat_let_tv94);
                }(_pat_let58_0);
              }((_dafny.ZERO).minus(new BigNumber((_pat_let_tv93).length)));
            }(_pat_let57_0);
          }(_2658_v136);
          _nw380[(new BigNumber(2)).toNumber()] = (_2655_v133).fm5(_2647_cf45, globalState);
          _nw380[(new BigNumber(3)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(4)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(5)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(6)).toNumber()] = function (_pat_let60_0) {
            return function (_2668_dt__update__tmp_h2) {
              return function (_pat_let61_0) {
                return function (_2669_dt__update_hcf2_h1) {
                  return function (_pat_let62_0) {
                    return function (_2670_dt__update_hcf1_h0) {
                      return _module.D0.create_DC1(_2670_dt__update_hcf1_h0, _2669_dt__update_hcf2_h1, (_2668_dt__update__tmp_h2).dtor_cf3, (_2668_dt__update__tmp_h2).dtor_cf4);
                    }(_pat_let62_0);
                  }(_pat_let_tv99);
                }(_pat_let61_0);
              }(_module.__default.fm36(_module.__default.fm2(_pat_let_tv95), _module.__default.fm2(_pat_let_tv96), _pat_let_tv97, _pat_let_tv98));
            }(_pat_let60_0);
          }(_2658_v136);
          _nw380[(new BigNumber(7)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(8)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(9)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(10)).toNumber()] = ((_2480_v6) ? (function (_pat_let63_0) {
            return function (_2671_dt__update__tmp_h3) {
              return function (_pat_let64_0) {
                return function (_2672_dt__update_hcf4_h0) {
                  return _module.D0.create_DC1((_2671_dt__update__tmp_h3).dtor_cf1, (_2671_dt__update__tmp_h3).dtor_cf2, (_2671_dt__update__tmp_h3).dtor_cf3, _2672_dt__update_hcf4_h0);
                }(_pat_let64_0);
              }(!(_pat_let_tv100));
            }(_pat_let63_0);
          }(_2658_v136)) : (_module.D0.create_DC1((_2655_v133).f26, _2651_v129, new BigNumber((_2660_v138).length), _2480_v6)));
          _nw380[(new BigNumber(11)).toNumber()] = _module.D0.create_DC1(new BigNumber((_module.__default.fm0(globalState)).length), _2651_v129, new BigNumber(774), _2648_cf44);
          _nw380[(new BigNumber(12)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(13)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(14)).toNumber()] = _module.D0.create_DC1(new BigNumber((_2657_v135).length), _dafny.Seq.UnicodeFromString("nfcbcyxby"), new BigNumber(133), _2480_v6);
          _nw380[(new BigNumber(15)).toNumber()] = _module.D0.create_DC1(_2647_cf45, _2651_v129, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2648_cf44,_2480_v6)).length), _2480_v6);
          _nw380[(new BigNumber(16)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(17)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(18)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(19)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(20)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(21)).toNumber()] = _module.__default.fm15(_2648_cf44, (_2655_v133).f26, _dafny.Seq.of(_2480_v6), _2480_v6, globalState);
          _nw380[(new BigNumber(22)).toNumber()] = _2658_v136;
          _nw380[(new BigNumber(23)).toNumber()] = (_2655_v133).fm5(_2473_v0, globalState);
          _nw380[(new BigNumber(24)).toNumber()] = _module.D0.create_DC1(new BigNumber((_2651_v129).length), _2651_v129, (_dafny.ZERO).minus((_2655_v133).f26), _2480_v6);
          _nw380[(new BigNumber(25)).toNumber()] = _module.D0.create_DC1(new BigNumber((_2663_v141).length), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("du"), _module.__default.safeIndex((_2655_v133).f26, new BigNumber((_dafny.Seq.UnicodeFromString("du")).length)), _2659_v137), _2647_cf45, _2480_v6);
          _nw380[(new BigNumber(26)).toNumber()] = _module.__default.fm15(_2480_v6, (_2655_v133).f26, _2649_cf43, _2480_v6, globalState);
          _2664_v142 = _nw380;
          let _index369 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_2664_v142).length));
          (_2664_v142)[_index369] = _module.D0.create_DC1(new BigNumber((_2651_v129).length), _dafny.Seq.UnicodeFromString("e"), _2647_cf45, _2480_v6);
          let _index370 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_2664_v142).length));
          (_2664_v142)[_index370] = _2658_v136;
          let _2673_v143;
          _2673_v143 = _dafny.MultiSet.fromElements(_2651_v129);
          (globalState).f8 = !((_2673_v143).IsDisjointFrom((_2673_v143).Union(_2673_v143)));
          let _2674_v144;
          _2674_v144 = _dafny.Set.fromElements((_dafny.MultiSet.fromElements(_2473_v0, _2473_v0, _2473_v0)).IsDisjointFrom(_2653_v131));
          let _rhs388 = (_2611_v101)[_module.__default.safeIndex(((_dafny.ZERO).minus(_2473_v0)).multipliedBy((_2655_v133).f26), new BigNumber((_2611_v101).length))];
          let _rhs389 = _2653_v131;
          let _rhs390 = _dafny.Map.Empty.slice().updateUnsafe((_2674_v144).IsDisjointFrom(_2674_v144),_2648_cf44);
          let _rhs391 = (_2674_v144).Intersect((_2674_v144).Union(_2674_v144));
          r0 = _rhs388;
          _2653_v131 = _rhs389;
          r1 = _rhs390;
          _2674_v144 = _rhs391;
          _2607_v99 = _2607_v99;
        } else {
          let _2675_v145;
          let _nw381 = Array((new BigNumber(7)).toNumber()).fill([]);
          _2675_v145 = _nw381;
          let _2676_v146;
          _2676_v146 = _dafny.Map.Empty.slice().updateUnsafe(!(!(false) || (_2480_v6)),_2675_v145);
          _2676_v146 = (_2676_v146).update(_2648_cf44, _2675_v145);
          _2647_cf45 = _2473_v0;
          (globalState).f8 = (_2655_v133).fm6(_2480_v6, (_2655_v133).f26, (_2655_v133).f26, globalState);
          let _index371 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index371] = new BigNumber(328);
          let _index372 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_2607_v99).length));
          (_2607_v99)[_index372] = (_module.__default.safeDivisionInt((_2655_v133).f26, new BigNumber(-509))).plus((_2655_v133).f26);
          _2647_cf45 = new BigNumber((_2654_v132).length);
        }
        let _2677_v147;
        let _nw382 = Array((new BigNumber(18)).toNumber()).fill(false);
        _2677_v147 = _nw382;
        let _2678_v148;
        _2678_v148 = _dafny.MultiSet.fromElements(_2677_v147, _2677_v147, _2677_v147, _2677_v147);
        let _2679_v149;
        let _nw383 = new _module.C12();
        _nw383.__ctor((_2678_v148).Intersect(_2678_v148), _2473_v0, _2648_cf44, _2647_cf45, _module.__default.safeDivisionInt(_2647_cf45, (_2655_v133).f26));
        _2679_v149 = _nw383;
      } else if (_source38.is_DC21) {
        let _2680___mcc_h15 = (_source38).cf39;
        let _2681_cf39 = _2680___mcc_h15;
        let _2682_v150;
        let _nw384 = new _module.C10();
        _nw384.__ctor(false);
        _2682_v150 = _nw384;
        let _2683_v151;
        _2683_v151 = _dafny.Set.fromElements(_2682_v150, _2682_v150);
        let _2684_v152;
        _2684_v152 = _dafny.Map.Empty.slice().updateUnsafe((_2683_v151).Union(_2683_v151),_2473_v0);
        _2684_v152 = _2684_v152;
        let _2685_v153;
        _2685_v153 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v0,_2480_v6);
        let _2686_v154;
        _2686_v154 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2473_v0,!(_2480_v6))).length), (_dafny.ZERO).minus(_2473_v0));
        let _2687_v155;
        let _nw385 = new _module.C0();
        _nw385.__ctor(_2473_v0, !((_2473_v0).isEqualTo(new BigNumber((_2685_v153).length))), (_dafny.ZERO).minus(_2473_v0), (_2686_v154)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_2686_v154).length))]);
        _2687_v155 = _nw385;
        r0 = _2687_v155.f29;
        let _2688_v156;
        let _nw386 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2688_v156 = _nw386;
        let _init69 = ((_2689_v155) => function (_2690_i10) {
          return _2689_v155.f29;
        })(_2687_v155);
        let _nw387 = Array((new BigNumber(11)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw387.length); _i0_69++) {
          _nw387[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2688_v156 = _nw387;
      } else {
        let _2691___mcc_h16 = (_source38).cf46;
        let _2692_cf46 = _2691___mcc_h16;
        let _2693_v157;
        let _init70 = ((_2694_v6) => function (_2695_i11) {
          return _2694_v6;
        })(_2480_v6);
        let _nw388 = Array((new BigNumber(7)).toNumber());
        for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw388.length); _i0_70++) {
          _nw388[_i0_70] = _init70(new BigNumber(_i0_70));
        }
        _2693_v157 = _nw388;
        let _2696_v158;
        let _init71 = function (_2697_i12) {
          return _module.D6.create_DC18(true);
        };
        let _nw389 = Array((new BigNumber(14)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw389.length); _i0_71++) {
          _nw389[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _2696_v158 = _nw389;
        let _2698_v159;
        _2698_v159 = _dafny.MultiSet.fromElements(_2696_v158, _2696_v158);
        let _2699_v160;
        _2699_v160 = _module.D11.create_DC30(_2698_v159);
        let _2700_v161;
        _2700_v161 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_2699_v160, _2699_v160, _2699_v160, _2699_v160, _2699_v160));
        let _2701_v162;
        _2701_v162 = _dafny.Seq.of(_2699_v160, _2699_v160);
        let _2702_v163;
        _2702_v163 = _dafny.Map.Empty.slice().updateUnsafe(_2693_v157,(_2700_v161).update(_2701_v162, _module.__default.abs(new BigNumber((_2611_v101).length))));
        r0 = ((((((_2702_v163).contains(_2693_v157)) ? ((_2702_v163).get(_2693_v157)) : (_2700_v161))).IsSubsetOf(_2700_v161)) ? (_module.__default.fm1(_2473_v0, _2480_v6, _2473_v0, globalState)) : (_2480_v6));
        let _2703_v164;
        _2703_v164 = _module.D7.create_DC23(_2610_v100, _2611_v101, true, _2473_v0);
        let _2704_v165;
        _2704_v165 = _module.D7.create_DC24(_2703_v164);
        let _2705_v166;
        _2705_v166 = _module.D7.create_DC24(_2703_v164);
        let _2706_v167;
        _2706_v167 = _dafny.Seq.of(_2705_v166);
        let _2707_v168;
        _2707_v168 = _dafny.Map.Empty.slice().updateUnsafe(_2706_v167,_2473_v0);
        let _2708_v169;
        _2708_v169 = _dafny.Map.Empty.slice().updateUnsafe(_2707_v168,_2693_v157);
        _2693_v157 = (((_2708_v169).contains(_2707_v168)) ? ((_2708_v169).get(_2707_v168)) : (_2693_v157));
        let _2709_v170;
        _2709_v170 = _dafny.Seq.UnicodeFromString("cefp");
        (globalState).f8 = _dafny.Seq.IsProperPrefixOf(_2709_v170, _dafny.Seq.UnicodeFromString("n"));
        let _2710_v171;
        _2710_v171 = _dafny.Seq.of(_2473_v0, _2473_v0);
        let _2711_v172;
        _2711_v172 = _dafny.Map.Empty.slice().updateUnsafe(_2710_v171,_2473_v0);
        let _index373 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_2607_v99).length));
        (_2607_v99)[_index373] = new BigNumber(((_2711_v172).update(_2710_v171, _2473_v0)).length);
        let _2712_v174;
        _2712_v174 = function () {
          let _coll99 = new _dafny.Set();
          for (const _compr_99 of _dafny.IntegerRange(new BigNumber(968), new BigNumber(401))) {
            let _2713_v173 = _compr_99;
            if (((new BigNumber(968)).isLessThanOrEqualTo(_2713_v173)) && ((_2713_v173).isLessThan(new BigNumber(401)))) {
              _coll99.add(_module.__default.safeDivisionInt(_2713_v173, _2473_v0));
            }
          }
          return _coll99;
        }();
        let _index374 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2693_v157).length));
        (_2693_v157)[_index374] = _2480_v6;
        let _2714_v175;
        _2714_v175 = _dafny.Set.fromElements(new BigNumber(-621), _2473_v0);
        let _index375 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_2607_v99).length));
        let _index376 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2693_v157).length));
        let _rhs392 = (_module.__default.safeDivisionInt(_2473_v0, _2473_v0)).multipliedBy(_2473_v0);
        let _rhs393 = _2714_v175;
        let _rhs394 = _2480_v6;
        let _rhs395 = false;
        let _rhs396 = (_2473_v0).minus(_2473_v0);
        let _lhs254 = _2607_v99;
        let _lhs255 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_2607_v99).length));
        let _lhs256 = _2693_v157;
        let _lhs257 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2693_v157).length));
        _lhs254[_lhs255] = _rhs392;
        _2712_v174 = _rhs393;
        _2480_v6 = _rhs394;
        _lhs256[_lhs257] = _rhs395;
        r2 = _rhs396;
      }
      let _2715_v176;
      _2715_v176 = _dafny.Seq.UnicodeFromString("qbbw");
      let _2716_v177;
      _2716_v177 = _dafny.Map.Empty.slice().updateUnsafe(_2473_v0,new BigNumber(855));
      let _2717_v178;
      _2717_v178 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber((_2716_v177).length));
      let _2718_v179;
      _2718_v179 = _dafny.Map.Empty.slice().updateUnsafe(_2717_v178,!(_2480_v6));
      let _2719_v180;
      let _nw390 = new _module.C2();
      _nw390.__ctor(_module.__default.safeDivisionInt(new BigNumber(297), _2473_v0), _2715_v176, (_dafny.ZERO).minus(new BigNumber((_2718_v179).length)), _2473_v0, _2480_v6);
      _2719_v180 = _nw390;
      let _2720_v181;
      let _init72 = ((_2721_v176) => function (_2722_i13) {
        return _2721_v176;
      })(_2715_v176);
      let _nw391 = Array((new BigNumber(25)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw391.length); _i0_72++) {
        _nw391[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2720_v181 = _nw391;
      let _2723_v182;
      _2723_v182 = _dafny.Map.Empty.slice().updateUnsafe(_2720_v181,_2480_v6);
      _2723_v182 = (_2723_v182).update(_2720_v181, _2480_v6);
      let _2724_v183;
      _2724_v183 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("q"),!(_2480_v6));
      let _2725_v185;
      _2725_v185 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2473_v0,(_2719_v180).f26)).length),_2480_v6);
      let _2726_v186;
      _2726_v186 = new _dafny.CodePoint('f'.codePointAt(0));
      let _2727_v187;
      _2727_v187 = _dafny.Seq.of(_module.__default.fm36((_2719_v180).f26, (_2719_v180).f26, _2726_v186, globalState), _2715_v176, _2715_v176);
      let _2728_v188;
      _2728_v188 = _dafny.Set.fromElements(_2473_v0, _2473_v0);
      r0 = !((_dafny.Set.fromElements((_2719_v180).f26)).Union(_2728_v188)).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(((_module.__default.fm1(new BigNumber((function () {
        let _coll100 = new _dafny.Set();
        for (const _compr_100 of (_2724_v183).Keys.Elements) {
          let _2729_v184 = _compr_100;
          if ((_2724_v183).contains(_2729_v184)) {
            _coll100.add(_2729_v184);
          }
        }
        return _coll100;
      }()).length), _2480_v6, (_dafny.ZERO).minus(new BigNumber((_2725_v185).length)), globalState)) ? ((_2727_v187)[_module.__default.safeIndex((_2719_v180).f26, new BigNumber((_2727_v187).length))]) : (_2715_v176)), _module.__default.safeIndex((_2719_v180).f26, new BigNumber((((_module.__default.fm1(new BigNumber((function () {
        let _coll101 = new _dafny.Set();
        for (const _compr_101 of (_2724_v183).Keys.Elements) {
          let _2730_v184 = _compr_101;
          if ((_2724_v183).contains(_2730_v184)) {
            _coll101.add(_2730_v184);
          }
        }
        return _coll101;
      }()).length), _2480_v6, (_dafny.ZERO).minus(new BigNumber((_2725_v185).length)), globalState)) ? ((_2727_v187)[_module.__default.safeIndex((_2719_v180).f26, new BigNumber((_2727_v187).length))]) : (_2715_v176))).length)), _2726_v186)).length)));
      let _2731_v189;
      _2731_v189 = _dafny.Map.Empty.slice().updateUnsafe(_2480_v6,_2480_v6);
      r1 = (_2731_v189).Merge(_2731_v189);
      r2 = (_2719_v180).f26;
      r3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2611_v101, _dafny.Seq.update(_2611_v101, _module.__default.safeIndex(_2473_v0, new BigNumber((_2611_v101).length)), false)), _2611_v101);
      return [r0, r1, r2, r3];
    }
  };

  $module.C14 = class C14 {
    constructor () {
      this._tname = "_module.C14";
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1((new BigNumber((_dafny.Set.fromElements(_this.f11)).length)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber(158))).length)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wxliobhom"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), function (_2732_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})), (new BigNumber((_dafny.Seq.of(_this.f11, _this.f11, (_module.D0.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("gbesth")).length), _this.f11)).dtor_cf8, _this.f11, false)).length)).plus((_dafny.ZERO).minus(new BigNumber(-157))), false);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ehqwd"), _dafny.Seq.UnicodeFromString("viqa"))).length)).isLessThan((new BigNumber((_dafny.Seq.of(false, true)).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(867)), function (_2733_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length)));
    };
    fm11(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ebhb")).length)))).Union(_dafny.MultiSet.fromElements(new BigNumber(-492)))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-929))).length)));
    };
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2734_v0;
      _2734_v0 = new BigNumber(-41);
      let _2735_v1;
      _2735_v1 = _dafny.Set.fromElements(p0);
      let _2736_v2;
      _2736_v2 = _dafny.Seq.UnicodeFromString("lhd");
      let _2737_v3;
      _2737_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2734_v0,_2734_v0);
      let _2738_v4;
      let _nw392 = Array((new BigNumber(28)).toNumber());
      _nw392[(_dafny.ZERO).toNumber()] = _2734_v0;
      _nw392[(_dafny.ONE).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(2)).toNumber()] = new BigNumber((_2735_v1).length);
      _nw392[(new BigNumber(3)).toNumber()] = new BigNumber(844);
      _nw392[(new BigNumber(4)).toNumber()] = new BigNumber(555);
      _nw392[(new BigNumber(5)).toNumber()] = new BigNumber(((_module.D1.create_DC4(p2)).dtor_cf9).length);
      _nw392[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_2734_v0);
      _nw392[(new BigNumber(7)).toNumber()] = new BigNumber((_2736_v2).length);
      _nw392[(new BigNumber(8)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(9)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(10)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(11)).toNumber()] = new BigNumber(734);
      _nw392[(new BigNumber(12)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(13)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(14)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(15)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_2734_v0);
      _nw392[(new BigNumber(17)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2734_v0,_this.f11)).length);
      _nw392[(new BigNumber(19)).toNumber()] = (((_2737_v3).contains(_2734_v0)) ? ((_2737_v3).get(_2734_v0)) : (_2734_v0));
      _nw392[(new BigNumber(20)).toNumber()] = new BigNumber(223);
      _nw392[(new BigNumber(21)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_2734_v0);
      _nw392[(new BigNumber(23)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(24)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(25)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(26)).toNumber()] = _2734_v0;
      _nw392[(new BigNumber(27)).toNumber()] = _2734_v0;
      _2738_v4 = _nw392;
      let _hi12 = (_module.D0.create_DC2((_dafny.ZERO).minus(_2734_v0), _2738_v4)).dtor_cf5;
      for (let _2739_i0 = _2734_v0; _2739_i0.isLessThan(_hi12); _2739_i0 = _2739_i0.plus(_dafny.ONE)) {
        r1 = _this.f11;
        let _2740_v5;
        _2740_v5 = new _dafny.CodePoint('v'.codePointAt(0));
        let _2741_v6;
        _2741_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2737_v3,_2740_v5);
        let _2742_v7;
        _2742_v7 = _dafny.MultiSet.fromElements(_this.f11, p0);
        let _2743_v8;
        _2743_v8 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2741_v6).length),new BigNumber((_2742_v7).cardinality()))).length), _2739_i0);
        r1 = (_module.__default.fm12(_2734_v0, _2734_v0, (((_2743_v8).contains(new BigNumber((_2742_v7).cardinality()))) ? ((_2743_v8).get(new BigNumber((_2742_v7).cardinality()))) : (_2734_v0)), p0, globalState)).IsSubsetOf((_dafny.MultiSet.FromArray(_module.__default.fm0(globalState))).Intersect(_2743_v8));
        let _2744_v9;
        let _nw393 = new _module.C1();
        _nw393.__ctor(new BigNumber((_2736_v2).length));
        _2744_v9 = _nw393;
        let _pat_let_tv101 = _2738_v4;
        let _source39 = function (_pat_let65_0) {
          return function (_2749_dt__update__tmp_h2) {
            return function (_pat_let70_0) {
              return function (_2750_dt__update_hcf58_h2) {
                return _module.D12.create_DC33(_2750_dt__update_hcf58_h2);
              }(_pat_let70_0);
            }(new BigNumber(226));
          }(_pat_let65_0);
        }(function (_pat_let66_0) {
          return function (_2747_dt__update__tmp_h1) {
            return function (_pat_let69_0) {
              return function (_2748_dt__update_hcf58_h1) {
                return _module.D12.create_DC33(_2748_dt__update_hcf58_h1);
              }(_pat_let69_0);
            }(new BigNumber((_dafny.Seq.of(_pat_let_tv101)).length));
          }(_pat_let66_0);
        }(function (_pat_let67_0) {
          return function (_2745_dt__update__tmp_h0) {
            return function (_pat_let68_0) {
              return function (_2746_dt__update_hcf58_h0) {
                return _module.D12.create_DC33(_2746_dt__update_hcf58_h0);
              }(_pat_let68_0);
            }(_2739_i0);
          }(_pat_let67_0);
        }(_module.D12.create_DC33(_2734_v0))));
        if (_source39.is_DC33) {
          let _2751___mcc_h0 = (_source39).cf58;
          let _2752_cf58 = _2751___mcc_h0;
          r1 = p0;
          let _index377 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_2738_v4).length));
          (_2738_v4)[_index377] = _2739_i0;
          let _index378 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_2738_v4).length));
          (_2738_v4)[_index378] = (_dafny.ZERO).minus(_module.__default.fm2(globalState));
          (globalState).f8 = p0;
          (globalState).f8 = false;
        } else if (_source39.is_DC34) {
          let _2753___mcc_h1 = (_source39).cf59;
          let _2754_cf59 = _2753___mcc_h1;
          let _2755_v10;
          let _nw394 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2755_v10 = _nw394;
          let _2756_v11;
          _2756_v11 = _module.D7.create_DC22(new BigNumber(-520), _2755_v10);
          let _2757_v12;
          _2757_v12 = _dafny.Seq.of(_module.D7.create_DC22(new BigNumber(548), _2755_v10), _2756_v11, _2756_v11, _2756_v11, _2756_v11);
          _2757_v12 = _2757_v12;
          let _2758_v13;
          _2758_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2743_v8,_this.f11);
          let _2759_v14;
          _2759_v14 = _dafny.Seq.of(_2744_v9.f27);
          let _2760_v15;
          _2760_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.FromArray(_2759_v14));
          _2758_v13 = (_2758_v13).update((((_2760_v15).contains((_this).fm6(p0, _2744_v9.f27, new BigNumber(618), globalState))) ? ((_2760_v15).get((_this).fm6(p0, _2744_v9.f27, new BigNumber(618), globalState))) : (_2743_v8)), p0);
          (_2744_v9).m24(globalState);
          (_2744_v9).m24(globalState);
        } else if (_source39.is_DC35) {
          let _2761___mcc_h2 = (_source39).cf60;
          let _2762_cf60 = _2761___mcc_h2;
          let _2763_v16;
          _2763_v16 = _dafny.Seq.of(_this.f11, !(false), p0, _this.f11, _this.f11);
          let _2764_v17;
          _2764_v17 = _module.D6.create_DC19(p0, _this.f11, p0, p0, _2763_v16);
          let _2765_v18;
          _2765_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2740_v5,(_dafny.ZERO).minus(_2739_i0));
          let _2766_v19;
          _2766_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_2742_v7).contains((_2764_v17).dtor_cf34)) ? ((_2742_v7).get((_2764_v17).dtor_cf34)) : (new BigNumber((_2765_v18).length))),p0);
          _2763_v16 = _dafny.Seq.of((((_2766_v19).contains(new BigNumber((p1).length))) ? ((_2766_v19).get(new BigNumber((p1).length))) : (p0)), p0, _this.f11, p0);
          _2738_v4 = _2738_v4;
          (_2744_v9).f27 = _2744_v9.f27;
          let _2767_v20;
          let _init73 = ((_2768_cf60) => function (_2769_i1) {
            return (_2768_cf60).isLessThanOrEqualTo(new BigNumber(790));
          })(_2762_cf60);
          let _nw395 = Array((new BigNumber(26)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw395.length); _i0_73++) {
            _nw395[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2767_v20 = _nw395;
          let _index379 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_2767_v20).length));
          (_2767_v20)[_index379] = p0;
          let _index380 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_2767_v20).length));
          (_2767_v20)[_index380] = p0;
        } else {
          let _2770___mcc_h3 = (_source39).cf57;
          let _2771_cf57 = _2770___mcc_h3;
          let _2772_v21;
          let _nw396 = new _module.C11();
          _nw396.__ctor(((_this.f11) ? (new BigNumber((_2736_v2).length)) : (new BigNumber((_2743_v8).cardinality()))), _2744_v9.f27);
          _2772_v21 = _nw396;
          let _2773_v22;
          _2773_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_2743_v8, _this.f11),_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), ((_2774_v5) => function (_2775_i2) {
            return _2774_v5;
          })(_2740_v5)));
          let _2776_v23;
          let _init74 = ((_2777_v5) => function (_2778_i3) {
            return (_module.D9.create_DC26(_2777_v5)).dtor_cf48;
          })(_2740_v5);
          let _nw397 = Array((new BigNumber(19)).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw397.length); _i0_74++) {
            _nw397[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2776_v23 = _nw397;
          let _index381 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2776_v23).length));
          (_2776_v23)[_index381] = new _dafny.CodePoint('g'.codePointAt(0));
          let _2779_v24;
          let _init75 = ((_2780_p0) => function (_2781_i4) {
            return _2780_p0;
          })(p0);
          let _nw398 = Array((_dafny.ONE).toNumber());
          for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw398.length); _i0_75++) {
            _nw398[_i0_75] = _init75(new BigNumber(_i0_75));
          }
          _2779_v24 = _nw398;
          let _2782_v25;
          _2782_v25 = _module.D1.create_DC5(_dafny.MultiSet.fromElements(_2739_i0, _2744_v9.f27), !(p0));
          let _2783_v26;
          _2783_v26 = _dafny.Set.fromElements(_2735_v1);
          let _index382 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2776_v23).length));
          let _rhs397 = _this.f11;
          let _rhs398 = _dafny.Map.Empty.slice().updateUnsafe(_2782_v25,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("whowf"), _2736_v2));
          let _rhs399 = ((_2783_v26).Union(_2783_v26)).IsSubsetOf(_2783_v26);
          let _rhs400 = ((_this.f11) ? (_2740_v5) : (_2740_v5));
          let _rhs401 = _2779_v24;
          let _lhs258 = globalState;
          let _lhs259 = _this;
          let _lhs260 = _2776_v23;
          let _lhs261 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2776_v23).length));
          _lhs258.f8 = _rhs397;
          _2773_v22 = _rhs398;
          _lhs259.f11 = _rhs399;
          _lhs260[_lhs261] = _rhs400;
          _2779_v24 = _rhs401;
          _2736_v2 = _2736_v2;
          let _2784_v27;
          _2784_v27 = _dafny.Map.Empty.slice().updateUnsafe((_2776_v23)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2776_v23).length))],_this.f11);
          _2736_v2 = _dafny.Seq.Concat((_2772_v21).fm8(_2744_v9.f27, _2784_v27, p1, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_2785_v5) => function (_2786_i5) {
            return _2785_v5;
          })(_2740_v5)));
        }
      }
      let _2787_v28;
      let _init76 = function (_2788_i6) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      };
      let _nw399 = Array((new BigNumber(18)).toNumber());
      for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw399.length); _i0_76++) {
        _nw399[_i0_76] = _init76(new BigNumber(_i0_76));
      }
      _2787_v28 = _nw399;
      let _2789_v29;
      _2789_v29 = new _dafny.CodePoint('s'.codePointAt(0));
      let _2790_v30;
      _2790_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2789_v29);
      let _index383 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length));
      (_2787_v28)[_index383] = (((_2790_v30).contains(_this.f11)) ? ((_2790_v30).get(_this.f11)) : (_2789_v29));
      let _index384 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length));
      (_2787_v28)[_index384] = _2789_v29;
      let _2791_v32;
      _2791_v32 = _dafny.Seq.of(_2736_v2, _2736_v2);
      let _2792_v33;
      _2792_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wueuik"),false);
      let _2793_i7;
      _2793_i7 = _dafny.ZERO;
      L12: {
        while (!(!(function () {
          let _coll103 = new _dafny.Map();
          for (const _compr_103 of (_2791_v32).Elements) {
            let _2807_v31 = _compr_103;
            if (_dafny.Seq.contains(_2791_v32, _2807_v31)) {
              _coll103.push([_2807_v31,p0]);
            }
          }
          return _coll103;
        }()).equals(_2792_v33)) || (_this.f11)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2793_i7)) {
              break L12;
            }
            _2793_i7 = (_2793_i7).plus(_dafny.ONE);
            (globalState).f8 = _this.f11;
            let _2794_v34;
            let _nw400 = new _module.C10();
            _nw400.__ctor(p0);
            _2794_v34 = _nw400;
            let _2795_v35;
            _2795_v35 = _dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))], (_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]);
            let _2796_v36;
            _2796_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2795_v35,p0);
            let _2797_v37;
            _2797_v37 = _module.D19.create_DC50(_2796_v36);
            let _2798_v39;
            let _nw401 = Array((new BigNumber(6)).toNumber());
            _nw401[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]),p0)).update(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))), false);
            _nw401[(_dafny.ONE).toNumber()] = (_2797_v37).dtor_cf76;
            _nw401[(new BigNumber(2)).toNumber()] = (_2797_v37).dtor_cf76;
            _nw401[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2795_v35,_module.__default.fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-464)), ((_2799_v28) => function (_2800_i8) {
              return (_2799_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2799_v28).length))];
            })(_2787_v28))).length), p0, new BigNumber((_dafny.MultiSet.fromElements(!(_this.f11), p0)).cardinality()), globalState));
            _nw401[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_2736_v2),p0)).Merge(_2796_v36);
            _nw401[(new BigNumber(5)).toNumber()] = (_2796_v36).Merge(function () {
              let _coll102 = new _dafny.Map();
              for (const _compr_102 of (_dafny.Seq.update(_dafny.Seq.of(_dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]), (_2795_v35).update(_2789_v29, _module.__default.abs(_2734_v0))), _module.__default.safeIndex(_2734_v0, new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]), (_2795_v35).update(_2789_v29, _module.__default.abs(_2734_v0)))).length)), _2795_v35)).Elements) {
                let _2801_v38 = _compr_102;
                if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.of(_dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]), (_2795_v35).update(_2789_v29, _module.__default.abs(_2734_v0))), _module.__default.safeIndex(_2734_v0, new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements((_2787_v28)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2787_v28).length))]), (_2795_v35).update(_2789_v29, _module.__default.abs(_2734_v0)))).length)), _2795_v35), _2801_v38)) {
                  _coll102.push([_2801_v38,false]);
                }
              }
              return _coll102;
            }());
            _2798_v39 = _nw401;
            let _index385 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_2798_v39).length));
            (_2798_v39)[_index385] = (_2796_v36).update(_2795_v35, p0);
            let _2802_v40;
            _2802_v40 = _module.D15.create_DC41(_this.f11);
            let _index386 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_2798_v39).length));
            let _rhs402 = _this.f11;
            let _rhs403 = _2738_v4;
            let _rhs404 = _2796_v36;
            let _rhs405 = (_2802_v40).dtor_cf65;
            let _rhs406 = _2736_v2;
            let _lhs262 = _2798_v39;
            let _lhs263 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_2798_v39).length));
            let _lhs264 = globalState;
            r1 = _rhs402;
            _2738_v4 = _rhs403;
            _lhs262[_lhs263] = _rhs404;
            _lhs264.f8 = _rhs405;
            _2736_v2 = _rhs406;
            let _2803_v41;
            let _nw402 = Array((new BigNumber(19)).toNumber()).fill(false);
            _2803_v41 = _nw402;
            let _index387 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_2803_v41).length));
            (_2803_v41)[_index387] = p0;
            let _2804_v42;
            _2804_v42 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2734_v0);
            let _2805_v43;
            _2805_v43 = _dafny.Seq.of(_2804_v42);
            let _2806_v44;
            _2806_v44 = _dafny.MultiSet.fromElements(false);
            let _index388 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_2803_v41).length));
            let _rhs407 = !_dafny.areEqual(_dafny.Seq.Concat(_2805_v43, _2805_v43), _dafny.Seq.Concat(_2805_v43, _2805_v43));
            let _rhs408 = (((_2806_v44).IsSubsetOf(_dafny.MultiSet.fromElements(p0))) ? (_2734_v0) : (_2734_v0));
            let _rhs409 = _this.f11;
            let _rhs410 = _2736_v2;
            let _rhs411 = ((_2734_v0).multipliedBy(new BigNumber((_dafny.Seq.of(_2734_v0)).length))).multipliedBy(_2734_v0);
            let _lhs265 = _2803_v41;
            let _lhs266 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_2803_v41).length));
            let _lhs267 = _this;
            _lhs265[_lhs266] = _rhs407;
            r0 = _rhs408;
            _lhs267.f11 = _rhs409;
            _2736_v2 = _rhs410;
            _2734_v0 = _rhs411;
          }
        }
      }
      let _2808_v45;
      _2808_v45 = _module.D4.create_DC12(_this.f11, new BigNumber(-585), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-289)), function (_2809_i9) {
  return new _dafny.CodePoint('o'.codePointAt(0));
})).length));
      let _2810_v46;
      _2810_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2734_v0);
      let _2811_v48;
      _2811_v48 = _dafny.Set.fromElements(_dafny.Seq.of(_module.__default.fm46(p0, _this.f11, globalState), _2808_v45, _module.D4.create_DC12(_this.f11, (((_2810_v46).contains(p0)) ? ((_2810_v46).get(p0)) : (new BigNumber((function () {
  let _coll104 = new _dafny.Set();
  for (const _compr_104 of _dafny.IntegerRange(new BigNumber(-622), new BigNumber(857))) {
    let _2812_v47 = _compr_104;
    if (((new BigNumber(-622)).isLessThanOrEqualTo(_2812_v47)) && ((_2812_v47).isLessThan(new BigNumber(857)))) {
      _coll104.add((_2812_v47).minus(_2734_v0));
    }
  }
  return _coll104;
}()).length))), _2734_v0), _2808_v45, _2808_v45));
      _2811_v48 = (_2811_v48).Difference(_2811_v48);
      let _2813_v49;
      _2813_v49 = _dafny.MultiSet.fromElements(_2738_v4, _2738_v4);
      (_this).f11 = !(!((_2813_v49).update(_2738_v4, _module.__default.abs(_2734_v0))).contains(_2738_v4));
      let _2814_v50;
      _2814_v50 = _module.D15.create_DC41(_this.f11);
      let _2815_v51;
      _2815_v51 = _module.D15.create_DC42(_2814_v50);
      let _2816_v52;
      let _nw403 = Array((new BigNumber(17)).toNumber());
      _nw403[(_dafny.ZERO).toNumber()] = _2815_v51;
      _nw403[(_dafny.ONE).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(2)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(3)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(4)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(5)).toNumber()] = _module.D15.create_DC42(_2814_v50);
      _nw403[(new BigNumber(6)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(7)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(8)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(9)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(10)).toNumber()] = _module.D15.create_DC42(_2814_v50);
      _nw403[(new BigNumber(11)).toNumber()] = _module.__default.fm61(globalState);
      _nw403[(new BigNumber(12)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(13)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(14)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(15)).toNumber()] = _2815_v51;
      _nw403[(new BigNumber(16)).toNumber()] = _2815_v51;
      _2816_v52 = _nw403;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2816_v52).length))) {
        let _2817_i10 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2817_i10)) && ((_2817_i10).isLessThan(new BigNumber((_2816_v52).length))))) {
          (_2816_v52)[(_2817_i10)] = _2815_v51;
        }
      }
      r0 = _2734_v0;
      r1 = p0;
      return [r0, r1];
    }
  };

  $module.C15 = class C15 {
    constructor () {
      this._tname = "_module.C15";
      this._f11 = false;
      this.f16 = false;
      this.f17 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    __ctor(f16, f17, f11) {
      let _this = this;
      (_this).f16 = f16;
      (_this).f17 = f17;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source40 = _module.D0.create_DC3(_this.f17, _this.f11);
      if (_source40.is_DC1) {
        let _2818___mcc_h0 = (_source40).cf1;
        let _2819___mcc_h1 = (_source40).cf2;
        let _2820___mcc_h2 = (_source40).cf3;
        let _2821___mcc_h3 = (_source40).cf4;
        let _2822_cf4 = _2821___mcc_h3;
        let _2823_cf3 = _2820___mcc_h2;
        let _2824_cf2 = _2819___mcc_h1;
        let _2825_cf1 = _2818___mcc_h0;
        return _module.D0.create_DC1(new BigNumber(622), _2824_cf2, new BigNumber((_dafny.Seq.of(_2825_cf1, new BigNumber((_2824_cf2).length))).length), _this.f11);
      } else if (_source40.is_DC2) {
        let _2826___mcc_h4 = (_source40).cf5;
        let _2827___mcc_h5 = (_source40).cf6;
        let _2828_cf6 = _2827___mcc_h5;
        let _2829_cf5 = _2826___mcc_h4;
        return _module.D0.create_DC1(_2829_cf5, _dafny.Seq.UnicodeFromString("bxbwjen"), _2829_cf5, !(_this.f11));
      } else if (_source40.is_DC3) {
        let _2830___mcc_h6 = (_source40).cf7;
        let _2831___mcc_h7 = (_source40).cf8;
        let _2832_cf8 = _2831___mcc_h7;
        let _2833_cf7 = _2830___mcc_h6;
        return _module.D0.create_DC1(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2832_cf8, _this.f16))).cardinality()), _dafny.Seq.UnicodeFromString("vtyog"), new BigNumber((_dafny.Seq.of(_2832_cf8)).length), !(_this.f16));
      } else {
        let _2834___mcc_h8 = (_source40).cf0;
        let _2835_cf0 = _2834___mcc_h8;
        return _module.D0.create_DC1(_this.f17, _dafny.Seq.UnicodeFromString("xuxniq"), new BigNumber(715), _this.f16);
      }
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return !(_this.f16);
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0))))));
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let _2836_v0;
      let _init77 = function (_2837_i0) {
        return _this.f11;
      };
      let _nw404 = Array((new BigNumber(16)).toNumber());
      for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw404.length); _i0_77++) {
        _nw404[_i0_77] = _init77(new BigNumber(_i0_77));
      }
      _2836_v0 = _nw404;
      let _2838_v1;
      _2838_v1 = _dafny.MultiSet.fromElements(_2836_v0);
      let _2839_v2;
      _2839_v2 = _dafny.Seq.UnicodeFromString("jtcekccsk");
      let _2840_v3;
      let _nw405 = new _module.C12();
      _nw405.__ctor(_2838_v1, new BigNumber((_dafny.Set.fromElements(_this.f17, p0)).length), _this.f16, new BigNumber((_dafny.Set.fromElements(new BigNumber((_2839_v2).length))).length), new BigNumber(372));
      _2840_v3 = _nw405;
      let _2841_v4;
      _2841_v4 = _module.D6.create_DC17(_2840_v3.f11, (_dafny.ZERO).minus(_this.f17));
      let _pat_let_tv102 = _2840_v3;
      let _pat_let_tv103 = _2839_v2;
      let _pat_let_tv104 = _2839_v2;
      let _pat_let_tv105 = _2839_v2;
      let _pat_let_tv106 = p1;
      let _pat_let_tv107 = p0;
      let _pat_let_tv108 = p2;
      let _pat_let_tv109 = p1;
      let _pat_let_tv110 = p2;
      let _pat_let_tv111 = _2840_v3;
      let _rhs412 = _this.f16;
      let _rhs413 = !(function (_source41) {
        if (_source41.is_DC17) {
          let _2842___mcc_h0 = (_source41).cf30;
          let _2843___mcc_h1 = (_source41).cf31;
          let _2844_cf31 = _2843___mcc_h1;
          let _2845_cf30 = _2842___mcc_h0;
          if (_this.f16) {
            return _pat_let_tv102.f11;
          } else {
            return _2845_cf30;
          }
        } else if (_source41.is_DC18) {
          let _2846___mcc_h2 = (_source41).cf32;
          let _2847_cf32 = _2846___mcc_h2;
          return (_dafny.MultiSet.fromElements(_pat_let_tv103, _pat_let_tv104)).contains(_pat_let_tv105);
        } else if (_source41.is_DC19) {
          let _2848___mcc_h3 = (_source41).cf33;
          let _2849___mcc_h4 = (_source41).cf34;
          let _2850___mcc_h5 = (_source41).cf35;
          let _2851___mcc_h6 = (_source41).cf36;
          let _2852___mcc_h7 = (_source41).cf37;
          let _2853_cf37 = _2852___mcc_h7;
          let _2854_cf36 = _2851___mcc_h6;
          let _2855_cf35 = _2850___mcc_h5;
          let _2856_cf34 = _2849___mcc_h4;
          let _2857_cf33 = _2848___mcc_h3;
          return (_dafny.MultiSet.fromElements(_pat_let_tv106)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv107,_this.f17)).length)));
        } else if (_source41.is_DC16) {
          let _2858___mcc_h8 = (_source41).cf29;
          let _2859_cf29 = _2858___mcc_h8;
          return !(_this.f16);
        } else {
          let _2860___mcc_h9 = (_source41).cf38;
          let _2861_cf38 = _2860___mcc_h9;
          return !(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_pat_let_tv108)[_module.__default.safeIndex(_pat_let_tv109, new BigNumber((_pat_let_tv110).length))])).contains(_pat_let_tv111.f11);
        }
      }(_2841_v4));
      let _rhs414 = _module.__default.fm2(globalState);
      let _rhs415 = _2840_v3;
      let _lhs268 = _this;
      let _lhs269 = _this;
      _lhs268.f11 = _rhs412;
      _lhs269.f11 = _rhs413;
      r1 = _rhs414;
      _2840_v3 = _rhs415;
      let _2862_i1;
      _2862_i1 = _dafny.ZERO;
      L13: {
        while (_2840_v3.f11) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2862_i1)) {
              break L13;
            }
            _2862_i1 = (_2862_i1).plus(_dafny.ONE);
            let _index389 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_2836_v0).length));
            (_2836_v0)[_index389] = _this.f16;
            let _index390 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_2836_v0).length));
            (_2836_v0)[_index390] = (_module.__default.fm2(globalState)).isEqualTo(p0);
            let _2863_v5;
            _2863_v5 = _dafny.Seq.of(p1);
            let _2864_v6;
            let _nw406 = new _module.C3();
            _nw406.__ctor((_2863_v5)[_module.__default.safeIndex(p1, new BigNumber((_2863_v5).length))], p1);
            _2864_v6 = _nw406;
            let _2865_v7;
            _2865_v7 = _dafny.Set.fromElements(_this.f16);
            let _2866_v8;
            _2866_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.update(p2, _module.__default.safeIndex(_this.f17, new BigNumber((p2).length)), _2840_v3.f11));
            let _2867_v9;
            _2867_v9 = _dafny.MultiSet.fromElements(p2, (((_2866_v8).contains(p0)) ? ((_2866_v8).get(p0)) : (p2)));
            let _2868_v10;
            _2868_v10 = new _dafny.CodePoint('i'.codePointAt(0));
            let _2869_v11;
            _2869_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2840_v3.f11,_2868_v10);
            let _2870_v12;
            let _nw407 = Array((new BigNumber(16)).toNumber());
            _nw407[(_dafny.ZERO).toNumber()] = new BigNumber((_2865_v7).length);
            _nw407[(_dafny.ONE).toNumber()] = p1;
            _nw407[(new BigNumber(2)).toNumber()] = p1;
            _nw407[(new BigNumber(3)).toNumber()] = p1;
            _nw407[(new BigNumber(4)).toNumber()] = p1;
            _nw407[(new BigNumber(5)).toNumber()] = _this.f17;
            _nw407[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p1);
            _nw407[(new BigNumber(7)).toNumber()] = new BigNumber(522);
            _nw407[(new BigNumber(8)).toNumber()] = p0;
            _nw407[(new BigNumber(9)).toNumber()] = p0;
            _nw407[(new BigNumber(10)).toNumber()] = new BigNumber(((_2867_v9).Union(_module.__default.fm62(_this.f16, _this.f16, _this.f11, !(_this.f16), globalState))).cardinality());
            _nw407[(new BigNumber(11)).toNumber()] = new BigNumber((_2869_v11).length);
            _nw407[(new BigNumber(12)).toNumber()] = p0;
            _nw407[(new BigNumber(13)).toNumber()] = p0;
            _nw407[(new BigNumber(14)).toNumber()] = (_this.f17).plus(_this.f17);
            _nw407[(new BigNumber(15)).toNumber()] = new BigNumber(205);
            _2870_v12 = _nw407;
            let _index391 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2870_v12).length));
            (_2870_v12)[_index391] = _this.f17;
            let _2871_v13;
            _2871_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2840_v3.f11,p0);
            let _2872_v14;
            _2872_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_2871_v13);
            let _2873_v15;
            _2873_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f17,_2840_v3.f11);
            let _2874_v16;
            _2874_v16 = _dafny.Map.Empty.slice().updateUnsafe((((_2872_v14).contains(false)) ? ((_2872_v14).get(false)) : (_2871_v13)),_2873_v15);
            let _index392 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2870_v12).length));
            (_2870_v12)[_index392] = new BigNumber((_2874_v16).length);
            let _2875_v17;
            let _nw408 = new _module.C2();
            _nw408.__ctor(_this.f17, _2839_v2, _this.f17, _this.f17, (_2836_v0)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_2836_v0).length))]);
            _2875_v17 = _nw408;
            let _2876_v18;
            _2876_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2840_v3.f11,_2875_v17);
            let _2877_v19;
            let _nw409 = Array((new BigNumber(22)).toNumber());
            _nw409[(_dafny.ZERO).toNumber()] = _2875_v17;
            _nw409[(_dafny.ONE).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(2)).toNumber()] = (((_2876_v18).contains(_2840_v3.f11)) ? ((_2876_v18).get(_2840_v3.f11)) : (_2875_v17));
            _nw409[(new BigNumber(3)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(4)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(5)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(6)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(7)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(8)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(9)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(10)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(11)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(12)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(13)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(14)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(15)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(16)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(17)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(18)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(19)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(20)).toNumber()] = _2875_v17;
            _nw409[(new BigNumber(21)).toNumber()] = _2875_v17;
            _2877_v19 = _nw409;
            let _index393 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_2877_v19).length));
            (_2877_v19)[_index393] = _2875_v17;
            let _index394 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_2877_v19).length));
            (_2877_v19)[_index394] = _2875_v17;
          }
        }
      }
      let _2878_v20;
      _2878_v20 = _dafny.MultiSet.fromElements(new BigNumber(54));
      let _2879_v21;
      _2879_v21 = _dafny.Seq.of(_2878_v20);
      let _2880_v22;
      _2880_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f11);
      let _2881_v23;
      _2881_v23 = _dafny.Seq.of(new BigNumber(-518), p1);
      let _2882_v24;
      _2882_v24 = _dafny.Seq.of(new BigNumber((_2880_v22).length), (_dafny.ZERO).minus(new BigNumber((_2881_v23).length)), (_dafny.ZERO).minus(new BigNumber(-60)));
      let _2883_v25;
      _2883_v25 = _dafny.Set.fromElements(_this.f11, _2840_v3.f11);
      let _2884_v26;
      _2884_v26 = _dafny.Seq.of(_2883_v25);
      let _2885_v27;
      _2885_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f17);
      let _2886_v28;
      _2886_v28 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(p0, new BigNumber((p2).length))],_dafny.Seq.update(_2879_v21, _module.__default.safeIndex(_this.f17, new BigNumber((_2879_v21).length)), _dafny.MultiSet.FromArray(_2882_v24)));
      let _2887_v29;
      _2887_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2879_v21);
      let _2888_v30;
      let _nw410 = Array((new BigNumber(24)).toNumber());
      _nw410[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_dafny.MultiSet.fromElements(p1));
      _nw410[(_dafny.ONE).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(2)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2879_v21, _2879_v21);
      _nw410[(new BigNumber(4)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(5)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_2879_v21, _module.__default.safeIndex(_this.f17, new BigNumber((_2879_v21).length)), _dafny.MultiSet.FromArray(_2882_v24));
      _nw410[(new BigNumber(7)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(8)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(9)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_2879_v21, _dafny.Seq.update(_2879_v21, _module.__default.safeIndex(new BigNumber(((_2884_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2885_v27).length)), new BigNumber((_2884_v26).length))]).length), new BigNumber((_2879_v21).length)), _2878_v20));
      _nw410[(new BigNumber(11)).toNumber()] = (((_2886_v28).contains(_this.f16)) ? ((_2886_v28).get(_this.f16)) : (_2879_v21));
      _nw410[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), ((_2889_v24) => function (_2890_i2) {
        return _dafny.MultiSet.FromArray(_2889_v24);
      })(_2882_v24));
      _nw410[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_2878_v20, (_module.D1.create_DC5(_dafny.MultiSet.fromElements(new BigNumber((_2839_v2).length)), false)).dtor_cf10), _2879_v21);
      _nw410[(new BigNumber(14)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(15)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_2879_v21, _dafny.Seq.of(_2878_v20, _2878_v20));
      _nw410[(new BigNumber(17)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(18)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(19)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat((((_2887_v29).contains(p0)) ? ((_2887_v29).get(p0)) : (_2879_v21)), _2879_v21);
      _nw410[(new BigNumber(21)).toNumber()] = _module.__default.fm63(globalState);
      _nw410[(new BigNumber(22)).toNumber()] = _2879_v21;
      _nw410[(new BigNumber(23)).toNumber()] = _2879_v21;
      _2888_v30 = _nw410;
      let _index395 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2888_v30).length));
      (_2888_v30)[_index395] = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(p0), _2878_v20), _dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), ((_2891_v20) => function (_2892_i3) {
        return _2891_v20;
      })(_2878_v20)));
      let _2893_v31;
      _2893_v31 = new _dafny.CodePoint('y'.codePointAt(0));
      let _index396 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2888_v30).length));
      let _rhs416 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-703)), ((_2894_v20, _2895_v24) => function (_2896_i4) {
        return ((_module.D1.create_DC5(_2894_v20, _this.f16)).dtor_cf10).Difference(_dafny.MultiSet.FromArray(_2895_v24));
      })(_2878_v20, _2882_v24));
      let _rhs417 = false;
      let _rhs418 = _2893_v31;
      let _rhs419 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), function (_2897_i5) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _2839_v2);
      let _lhs270 = _2888_v30;
      let _lhs271 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_2888_v30).length));
      let _lhs272 = globalState;
      _lhs270[_lhs271] = _rhs416;
      _lhs272.f8 = _rhs417;
      _2893_v31 = _rhs418;
      _2839_v2 = _rhs419;
      let _2898_i6;
      _2898_i6 = _dafny.ZERO;
      L14: {
        while (_this.f16) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2898_i6)) {
              break L14;
            }
            _2898_i6 = (_2898_i6).plus(_dafny.ONE);
            if ((_2883_v25).IsProperSubsetOf(_2883_v25)) {
              _2836_v0 = _2836_v0;
              let _2899_v32;
              let _nw411 = new _module.C2();
              _nw411.__ctor(p1, _dafny.Seq.Concat(_2839_v2, _2839_v2), (_dafny.ZERO).minus(p0), p0, _2840_v3.f11);
              _2899_v32 = _nw411;
              r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm2(globalState)));
              let _2900_v34;
              _2900_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f17,_2840_v3.f11);
              let _2901_v35;
              _2901_v35 = _dafny.Map.Empty.slice().updateUnsafe((_2900_v34).update(p1, _2840_v3.f11),_2882_v24);
              let _2902_v36;
              _2902_v36 = function () {
                let _coll105 = new _dafny.Map();
                for (const _compr_105 of (_2901_v35).Keys.Elements) {
                  let _2903_v33 = _compr_105;
                  if ((_2901_v35).contains(_2903_v33)) {
                    _coll105.push([_2903_v33,p2]);
                  }
                }
                return _coll105;
              }();
              _2902_v36 = _2902_v36;
              (_this).f17 = p0;
            } else {
              let _2904_v37;
              let _nw412 = new _module.C1();
              _nw412.__ctor(_this.f17);
              _2904_v37 = _nw412;
              (_this).f17 = (_2904_v37).fm34(globalState);
              r1 = (_this.f17).plus(p0);
              (_this).f11 = _2840_v3.f11;
              (_this).f17 = (new BigNumber((_2882_v24).length)).multipliedBy(p0);
            }
            let _2905_v38;
            _2905_v38 = _dafny.Seq.of(_2839_v2);
            let _source42 = _module.__default.fm43(p0, ((_this.f16) ? (_this.f16) : (_2840_v3.f11)), p0, (_2905_v38)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_2905_v38).length))], globalState);
            if (_source42.is_DC17) {
              let _2906___mcc_h10 = (_source42).cf30;
              let _2907___mcc_h11 = (_source42).cf31;
              let _2908_cf31 = _2907___mcc_h11;
              let _2909_cf30 = _2906___mcc_h10;
              let _2910_v39;
              _2910_v39 = _2882_v24;
              let _2911_v40;
              _2911_v40 = _dafny.Set.fromElements(_2910_v39);
              let _2912_v41;
              _2912_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f17,false);
              let _2913_v42;
              let _nw413 = new _module.C9();
              _nw413.__ctor(_2911_v40, p1, _this.f17, (p1).plus(p0), (p1).isLessThanOrEqualTo(new BigNumber((_2912_v41).length)), _module.__default.fm36(p1, p0, _2893_v31, globalState));
              _2913_v42 = _nw413;
              _2913_v42 = _2913_v42;
              _2893_v31 = _2893_v31;
              r1 = p1;
              let _2914_v43;
              let _2915_v44;
              let _out43;
              let _out44;
              let _outcollector20 = _module.__default.m0(globalState);
              _out43 = _outcollector20[0];
              _out44 = _outcollector20[1];
              _2914_v43 = _out43;
              _2915_v44 = _out44;
            } else if (_source42.is_DC18) {
              let _2916___mcc_h12 = (_source42).cf32;
              let _2917_cf32 = _2916___mcc_h12;
              (globalState).f5 = _dafny.Seq.Concat(_2882_v24, _dafny.Seq.Concat(_2882_v24, _dafny.Seq.of(p1)));
              let _2918_v45;
              _2918_v45 = _dafny.MultiSet.fromElements(_2840_v3.f11, _this.f11);
              (_this).f17 = (((new BigNumber(131)).isLessThanOrEqualTo((((_2918_v45).contains(_2840_v3.f11)) ? ((_2918_v45).get(_2840_v3.f11)) : (p0)))) ? (p1) : (p1));
              let _2919_v46;
              let _nw414 = Array((new BigNumber(4)).toNumber());
              _nw414[(_dafny.ZERO).toNumber()] = p0;
              _nw414[(_dafny.ONE).toNumber()] = p1;
              _nw414[(new BigNumber(2)).toNumber()] = _this.f17;
              _nw414[(new BigNumber(3)).toNumber()] = p1;
              _2919_v46 = _nw414;
              let _index397 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              (_2919_v46)[_index397] = _this.f17;
              let _index398 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              (_2919_v46)[_index398] = (_module.__default.safeDivisionInt(p1, p0)).plus(p1);
              let _2920_v47;
              _2920_v47 = _module.D7.create_DC21((_dafny.MultiSet.fromElements(false, !(_2840_v3.f11), _this.f11)).update(false, _module.__default.abs(_this.f17)));
              let _2921_v48;
              _2921_v48 = _module.D7.create_DC24(_2920_v47);
              let _index399 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              let _index400 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              let _rhs420 = p0;
              let _rhs421 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(p2, p2), _dafny.Seq.Concat(p2, p2))).length);
              let _rhs422 = _2921_v48;
              let _rhs423 = p0;
              let _rhs424 = p0;
              let _lhs273 = _2919_v46;
              let _lhs274 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              let _lhs275 = _this;
              let _lhs276 = _2919_v46;
              let _lhs277 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2919_v46).length));
              _lhs273[_lhs274] = _rhs420;
              _lhs275.f17 = _rhs421;
              _2921_v48 = _rhs422;
              r1 = _rhs423;
              _lhs276[_lhs277] = _rhs424;
            } else if (_source42.is_DC19) {
              let _2922___mcc_h13 = (_source42).cf33;
              let _2923___mcc_h14 = (_source42).cf34;
              let _2924___mcc_h15 = (_source42).cf35;
              let _2925___mcc_h16 = (_source42).cf36;
              let _2926___mcc_h17 = (_source42).cf37;
              let _2927_cf37 = _2926___mcc_h17;
              let _2928_cf36 = _2925___mcc_h16;
              let _2929_cf35 = _2924___mcc_h15;
              let _2930_cf34 = _2923___mcc_h14;
              let _2931_cf33 = _2922___mcc_h13;
              let _2932_v49;
              _2932_v49 = _dafny.Seq.of(_2836_v0, _2836_v0, _2836_v0, _2836_v0, _2836_v0);
              let _2933_v50;
              _2933_v50 = _module.D4.create_DC12(_2928_cf36, new BigNumber((_2839_v2).length), p0);
              let _2934_v51;
              let _nw415 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
              _2934_v51 = _nw415;
              let _2935_v52;
              let _nw416 = Array((new BigNumber(26)).toNumber());
              _nw416[(_dafny.ZERO).toNumber()] = _2934_v51;
              _nw416[(_dafny.ONE).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(2)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(3)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(4)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(5)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(6)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(7)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(8)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(9)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(10)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(11)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(12)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(13)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(14)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(15)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(16)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(17)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(18)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(19)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(20)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(21)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(22)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(23)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(24)).toNumber()] = _2934_v51;
              _nw416[(new BigNumber(25)).toNumber()] = _2934_v51;
              _2935_v52 = _nw416;
              let _2936_v53;
              _2936_v53 = _module.D7.create_DC22(_this.f17, _2935_v52);
              let _2937_v54;
              _2937_v54 = _dafny.MultiSet.fromElements(_2929_cf35, _2928_cf36);
              let _2938_v55;
              _2938_v55 = _module.D7.create_DC21(_2937_v54);
              let _2939_v56;
              _2939_v56 = _dafny.Set.fromElements(_module.D7.create_DC21(_dafny.MultiSet.fromElements(_this.f11)), _2938_v55, _2938_v55);
              let _2940_v57;
              _2940_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2939_v56,_2893_v31);
              let _2941_v58;
              _2941_v58 = _2882_v24;
              let _2942_v59;
              _2942_v59 = _dafny.Set.fromElements(_2941_v58, _2941_v58, _2941_v58, _2941_v58);
              let _2943_v60;
              let _nw417 = new _module.C9();
              _nw417.__ctor(_2942_v59, p0, new BigNumber((_2937_v54).cardinality()), p1, _2840_v3.f11, _2839_v2);
              _2943_v60 = _nw417;
              let _2944_v61;
              _2944_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2943_v60,_2930_cf34);
              let _2945_v62;
              _2945_v62 = _dafny.Set.fromElements(new BigNumber((_2940_v57).length), new BigNumber((_2944_v61).length));
              let _pat_let_tv112 = p1;
              let _pat_let_tv113 = p0;
              let _pat_let_tv114 = p1;
              let _2946_v63;
              let _nw418 = Array((new BigNumber(14)).toNumber());
              _nw418[(_dafny.ZERO).toNumber()] = _2933_v50;
              _nw418[(_dafny.ONE).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(2)).toNumber()] = function (_pat_let71_0) {
                return function (_2949_dt__update__tmp_h1) {
                  return function (_pat_let74_0) {
                    return function (_2950_dt__update_hcf23_h0) {
                      return _module.D4.create_DC12((_2949_dt__update__tmp_h1).dtor_cf21, (_2949_dt__update__tmp_h1).dtor_cf22, _2950_dt__update_hcf23_h0);
                    }(_pat_let74_0);
                  }(new BigNumber((_dafny.MultiSet.fromElements(_pat_let_tv112, _pat_let_tv113)).cardinality()));
                }(_pat_let71_0);
              }(function (_pat_let72_0) {
                return function (_2947_dt__update__tmp_h0) {
                  return function (_pat_let73_0) {
                    return function (_2948_dt__update_hcf22_h0) {
                      return _module.D4.create_DC12((_2947_dt__update__tmp_h0).dtor_cf21, _2948_dt__update_hcf22_h0, (_2947_dt__update__tmp_h0).dtor_cf23);
                    }(_pat_let73_0);
                  }(_this.f17);
                }(_pat_let72_0);
              }(_2933_v50));
              _nw418[(new BigNumber(3)).toNumber()] = ((_2928_cf36) ? (_2933_v50) : (_2933_v50));
              _nw418[(new BigNumber(4)).toNumber()] = _module.D4.create_DC12(_this.f16, _this.f17, _this.f17);
              _nw418[(new BigNumber(5)).toNumber()] = function (_pat_let75_0) {
                return function (_2951_dt__update__tmp_h2) {
                  return function (_pat_let76_0) {
                    return function (_2952_dt__update_hcf23_h1) {
                      return _module.D4.create_DC12((_2951_dt__update__tmp_h2).dtor_cf21, (_2951_dt__update__tmp_h2).dtor_cf22, _2952_dt__update_hcf23_h1);
                    }(_pat_let76_0);
                  }(_pat_let_tv114);
                }(_pat_let75_0);
              }(_2933_v50);
              _nw418[(new BigNumber(6)).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(7)).toNumber()] = _module.D4.create_DC12(_2840_v3.f11, p0, p0);
              _nw418[(new BigNumber(8)).toNumber()] = _module.D4.create_DC12(_this.f11, (_module.D9.create_DC27(_2840_v3.f11, _2936_v53, p1, _2931_cf33)).dtor_cf51, new BigNumber((_2945_v62).length));
              _nw418[(new BigNumber(9)).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(10)).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(11)).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(12)).toNumber()] = _2933_v50;
              _nw418[(new BigNumber(13)).toNumber()] = _2933_v50;
              _2946_v63 = _nw418;
              let _2953_v64;
              _2953_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2943_v60.f22,_2928_cf36);
              let _2954_v65;
              _2954_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2953_v64,_2930_cf34);
              let _2955_v66;
              _2955_v66 = _module.D20.create_DC52(_2954_v65);
              let _rhs425 = new BigNumber((((_2955_v66).dtor_cf79).Merge(_2954_v65)).length);
              let _rhs426 = _dafny.Seq.Concat(_dafny.Seq.of(_2836_v0), _dafny.Seq.of(_2836_v0, _2836_v0, _2836_v0, _2836_v0, _2836_v0));
              let _rhs427 = _2946_v63;
              r1 = _rhs425;
              _2932_v49 = _rhs426;
              _2946_v63 = _rhs427;
              let _2956_v67;
              let _nw419 = new _module.C5();
              _nw419.__ctor(_2943_v60.f22);
              _2956_v67 = _nw419;
              let _index401 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_2836_v0).length));
              (_2836_v0)[_index401] = _2930_cf34;
              let _2957_v68;
              let _init78 = ((_2958_p1, _2959_v31) => function (_2960_i7) {
                return _dafny.Seq.of(_2958_p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_2961_v31) => function (_2962_i8) {
                  return _2961_v31;
                })(_2959_v31))).length));
              })(p1, _2893_v31);
              let _nw420 = Array((new BigNumber(9)).toNumber());
              for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw420.length); _i0_78++) {
                _nw420[_i0_78] = _init78(new BigNumber(_i0_78));
              }
              _2957_v68 = _nw420;
              let _2963_v69;
              _2963_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2943_v60.f22,p2);
              let _index402 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_2836_v0).length));
              let _rhs428 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(p2, _2927_cf37), (((_2963_v69).contains(_this.f17)) ? ((_2963_v69).get(_this.f17)) : (p2)));
              let _rhs429 = (_2927_cf37)[_module.__default.safeIndex(p0, new BigNumber((_2927_cf37).length))];
              let _rhs430 = _2957_v68;
              let _rhs431 = !(_2928_cf36);
              let _rhs432 = _2943_v60.f22;
              let _lhs278 = _2840_v3;
              let _lhs279 = _2836_v0;
              let _lhs280 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_2836_v0).length));
              let _lhs281 = _this;
              let _lhs282 = _this;
              _lhs278.f11 = _rhs428;
              _lhs279[_lhs280] = _rhs429;
              _2957_v68 = _rhs430;
              _lhs281.f11 = _rhs431;
              _lhs282.f17 = _rhs432;
              let _2964_v70;
              _2964_v70 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Set.fromElements(_2840_v3.f11));
              _2964_v70 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2883_v25).Intersect(_dafny.Set.fromElements(_this.f11)));
            } else if (_source42.is_DC16) {
              let _2965___mcc_h18 = (_source42).cf29;
              let _2966_cf29 = _2965___mcc_h18;
              _2839_v2 = _2839_v2;
              (_2840_v3).f11 = _this.f16;
              let _2967_v71;
              _2967_v71 = _module.D1.create_DC5(_2878_v20, _2840_v3.f11);
              let _2968_v72;
              _2968_v72 = _module.D1.create_DC6(_2967_v71);
              let _2969_v73;
              let _nw421 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
              _2969_v73 = _nw421;
              let _2970_v74;
              _2970_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _2971_v75;
              _2971_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2893_v31,_module.__default.fm50(globalState));
              let _2972_v76;
              _2972_v76 = _dafny.Set.fromElements(p0, new BigNumber((_2885_v27).length), new BigNumber((_2878_v20).cardinality()), _this.f17);
              let _index403 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2969_v73).length));
              (_2969_v73)[_index403] = ((((_2970_v74).contains(_this.f17)) ? ((_2970_v74).get(_this.f17)) : (new BigNumber(((((_2971_v75).contains(_2893_v31)) ? ((_2971_v75).get(_2893_v31)) : (_2972_v76))).length)))).plus(_this.f17);
              let _2973_v77;
              _2973_v77 = _module.D21.create_DC55(_2838_v1);
              let _2974_v78;
              let _nw422 = new _module.C12();
              _nw422.__ctor((_2973_v77).dtor_cf81, _this.f17, false, p0, new BigNumber(-840));
              _2974_v78 = _nw422;
              let _2975_v79;
              _2975_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_dafny.Map.Empty.slice().updateUnsafe(p0,(_2974_v78).f19));
              let _2976_v80;
              let _nw423 = Array((new BigNumber(3)).toNumber());
              _nw423[(_dafny.ZERO).toNumber()] = (((_2975_v79).contains(true)) ? ((_2975_v79).get(true)) : ((_2970_v74).update((_2974_v78).f19, _module.__default.fm2(globalState))));
              _nw423[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f17,new BigNumber(352));
              _nw423[(new BigNumber(2)).toNumber()] = _2970_v74;
              _2976_v80 = _nw423;
              let _index404 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2976_v80).length));
              (_2976_v80)[_index404] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2839_v2).length),(_2974_v78).f19);
              let _2977_v81;
              _2977_v81 = _2974_v78;
              let _index405 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2969_v73).length));
              let _index406 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2976_v80).length));
              let _rhs433 = _2968_v72;
              let _rhs434 = p0;
              let _rhs435 = (_2977_v81);
              let _rhs436 = ((_2970_v74).Merge(_module.__default.fm38(_2893_v31, _2893_v31, globalState))).Merge((_2970_v74).update(p0, p1));
              let _lhs283 = _2969_v73;
              let _lhs284 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2969_v73).length));
              let _lhs285 = _2976_v80;
              let _lhs286 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2976_v80).length));
              _2968_v72 = _rhs433;
              _lhs283[_lhs284] = _rhs434;
              _2974_v78 = _rhs435;
              _lhs285[_lhs286] = _rhs436;
              let _index407 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2969_v73).length));
              (_2969_v73)[_index407] = (_dafny.ZERO).minus(new BigNumber(((_2878_v20).Difference((_2878_v20).Difference(_2878_v20))).cardinality()));
            } else {
              let _2978___mcc_h19 = (_source42).cf38;
              let _2979_cf38 = _2978___mcc_h19;
              let _index408 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_2836_v0).length));
              (_2836_v0)[_index408] = _2840_v3.f11;
              let _index409 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_2836_v0).length));
              (_2836_v0)[_index409] = _2840_v3.f11;
              let _2980_v82;
              let _nw424 = Array((new BigNumber(5)).toNumber());
              _nw424[(_dafny.ZERO).toNumber()] = ((_this.f16) ? (new BigNumber(722)) : (_this.f17));
              _nw424[(_dafny.ONE).toNumber()] = (p0).minus(p1);
              _nw424[(new BigNumber(2)).toNumber()] = p1;
              _nw424[(new BigNumber(3)).toNumber()] = p1;
              _nw424[(new BigNumber(4)).toNumber()] = new BigNumber(805);
              _2980_v82 = _nw424;
              let _index410 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_2980_v82).length));
              (_2980_v82)[_index410] = p0;
              let _index411 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_2980_v82).length));
              (_2980_v82)[_index411] = p1;
              (globalState).f8 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(p2, _module.__default.safeIndex(new BigNumber((_2839_v2).length), new BigNumber((p2).length)), _this.f11), _module.__default.fm39((p2)[_module.__default.safeIndex(_this.f17, new BigNumber((p2).length))], _this.f11, globalState));
              let _index412 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_2980_v82).length));
              (_2980_v82)[_index412] = (((_dafny.ZERO).minus((_2980_v82)[_module.__default.safeIndex(new BigNumber(759), new BigNumber((_2980_v82).length))])).minus(p1)).multipliedBy((_2980_v82)[_module.__default.safeIndex(new BigNumber(759), new BigNumber((_2980_v82).length))]);
            }
            if (!(_2840_v3.f11) || (_2840_v3.f11)) {
              let _2981_v83;
              _2981_v83 = _dafny.Set.fromElements(new BigNumber(549));
              let _2982_v84;
              _2982_v84 = _dafny.Set.fromElements(_2981_v83, _2981_v83, _dafny.Set.fromElements(new BigNumber((p2).length)), _2981_v83, _module.__default.fm50(globalState));
              (globalState).f8 = ((_2982_v84).Intersect(_dafny.Set.fromElements(function () {
                let _coll106 = new _dafny.Set();
                for (const _compr_106 of _dafny.IntegerRange(new BigNumber(-143), new BigNumber(279))) {
                  let _2983_v85 = _compr_106;
                  if (((new BigNumber(-143)).isLessThanOrEqualTo(_2983_v85)) && ((_2983_v85).isLessThan(new BigNumber(279)))) {
                    _coll106.add((_2983_v85).multipliedBy(new BigNumber(586)));
                  }
                }
                return _coll106;
              }()))).IsSubsetOf(_2982_v84);
              let _2984_v86;
              let _nw425 = new _module.C0();
              _nw425.__ctor(p1, _2840_v3.f11, p1, p0);
              _2984_v86 = _nw425;
              let _index413 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_2836_v0).length));
              (_2836_v0)[_index413] = !(!(_this.f16) || (_2840_v3.f11));
              let _index414 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_2836_v0).length));
              (_2836_v0)[_index414] = _this.f16;
              let _2985_v87;
              _2985_v87 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2984_v86).f28),_this.f17);
              let _rhs437 = true;
              let _rhs438 = _2985_v87;
              let _rhs439 = _this.f17;
              let _rhs440 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2839_v2, _2839_v2), _2839_v2);
              let _lhs287 = _2984_v86;
              _lhs287.f29 = _rhs437;
              _2985_v87 = _rhs438;
              r1 = _rhs439;
              _2839_v2 = _rhs440;
              let _2986_v88;
              _2986_v88 = _dafny.MultiSet.fromElements(_2880_v22);
              let _2987_v89;
              _2987_v89 = (_2986_v88).Difference(_2986_v88);
              _2987_v89 = _2987_v89;
            } else {
              _2893_v31 = _2893_v31;
              let _2988_v90;
              let _nw426 = Array((new BigNumber(15)).toNumber()).fill([]);
              _2988_v90 = _nw426;
              let _2989_v91;
              _2989_v91 = _module.D7.create_DC23(_2988_v90, p2, _this.f16, new BigNumber((_2881_v23).length));
              let _2990_v92;
              let _nw427 = Array((new BigNumber(12)).toNumber());
              _nw427[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_this.f17, _this.f17);
              _nw427[(_dafny.ONE).toNumber()] = p1;
              _nw427[(new BigNumber(2)).toNumber()] = p1;
              _nw427[(new BigNumber(3)).toNumber()] = _this.f17;
              _nw427[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0);
              _nw427[(new BigNumber(5)).toNumber()] = (new BigNumber((p2).length)).multipliedBy((_2989_v91).dtor_cf45);
              _nw427[(new BigNumber(6)).toNumber()] = p0;
              _nw427[(new BigNumber(7)).toNumber()] = p1;
              _nw427[(new BigNumber(8)).toNumber()] = p1;
              _nw427[(new BigNumber(9)).toNumber()] = new BigNumber((_2839_v2).length);
              _nw427[(new BigNumber(10)).toNumber()] = p0;
              _nw427[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, _this.f17));
              _2990_v92 = _nw427;
              let _index415 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length));
              (_2990_v92)[_index415] = new BigNumber((_2839_v2).length);
              let _2991_v93;
              let _nw428 = new _module.C6();
              _nw428.__ctor();
              _2991_v93 = _nw428;
              let _2992_v94;
              _2992_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2991_v93,_dafny.Map.Empty.slice().updateUnsafe(_2840_v3.f11,p1));
              let _index416 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length));
              let _rhs441 = (_2885_v27).Merge((((_2992_v94).contains(_2991_v93)) ? ((_2992_v94).get(_2991_v93)) : (_2885_v27)));
              let _rhs442 = _2836_v0;
              let _rhs443 = (p0).plus(p0);
              let _lhs288 = _2990_v92;
              let _lhs289 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length));
              _2885_v27 = _rhs441;
              _2836_v0 = _rhs442;
              _lhs288[_lhs289] = _rhs443;
              r1 = (_dafny.ZERO).minus((_2990_v92)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length))]);
              let _2993_v95;
              _2993_v95 = _dafny.MultiSet.fromElements(_2840_v3.f11, _2840_v3.f11, _this.f16);
              let _2994_v96;
              _2994_v96 = _dafny.Set.fromElements(((_2990_v92)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length))]).minus((_2990_v92)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length))]), p1, (((_2885_v27).contains((_this).fm6(_this.f16, (_2990_v92)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length))], _this.f17, globalState))) ? ((_2885_v27).get((_this).fm6(_this.f16, (_2990_v92)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length))], _this.f17, globalState))) : (new BigNumber((_2993_v95).cardinality()))));
              let _index417 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length));
              let _rhs444 = !(_2840_v3.f11);
              let _rhs445 = (_dafny.ZERO).minus(_this.f17);
              let _rhs446 = (_2994_v96).Union(_2994_v96);
              let _lhs290 = _2840_v3;
              let _lhs291 = _2990_v92;
              let _lhs292 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2990_v92).length));
              _lhs290.f11 = _rhs444;
              _lhs291[_lhs292] = _rhs445;
              _2994_v96 = _rhs446;
              let _2995_v97;
              _2995_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2893_v31,(_2883_v25).Union(_2883_v25));
              _2995_v97 = (_2995_v97).update(_2893_v31, _2883_v25);
            }
            let _2996_v98;
            _2996_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm6(_this.f11, p1, p0, globalState),p2);
            let _2997_v99;
            let _nw429 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
            _2997_v99 = _nw429;
            let _2998_v100;
            _2998_v100 = _dafny.Map.Empty.slice().updateUnsafe(_2997_v99,!(_this.f16));
            let _2999_v101;
            _2999_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2996_v98).contains(_this.f11)) ? ((_2996_v98).get(_this.f11)) : (p2))).length),_2998_v100);
            _2999_v101 = (_2999_v101).update(p0, _dafny.Map.Empty.slice().updateUnsafe(_2997_v99,_2840_v3.f11));
          }
        }
      }
      let _3000_v105;
      let _init79 = ((_3001_v27, _3002_p0, _3003_v3) => function (_3004_i9) {
        return (function () {
          let _coll107 = new _dafny.Map();
          for (const _compr_107 of _dafny.IntegerRange(new BigNumber(5), new BigNumber(-382))) {
            let _3005_v102 = _compr_107;
            if (((new BigNumber(5)).isLessThanOrEqualTo(_3005_v102)) && ((_3005_v102).isLessThan(new BigNumber(-382)))) {
              _coll107.push([(_3005_v102).minus(new BigNumber(94)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(_3001_v27))).length)]);
            }
          }
          return _coll107;
        }()).Merge(function () {
          let _coll108 = new _dafny.Map();
          for (const _compr_108 of (function () {
            let _coll109 = new _dafny.Map();
            for (const _compr_109 of _dafny.IntegerRange(new BigNumber(897), new BigNumber(68))) {
              let _3006_v104 = _compr_109;
              if (((new BigNumber(897)).isLessThanOrEqualTo(_3006_v104)) && ((_3006_v104).isLessThan(new BigNumber(68)))) {
                _coll109.push([_module.__default.safeModuloInt(_3006_v104, _3002_p0),_3003_v3.f11]);
              }
            }
            return _coll109;
          }()).Keys.Elements) {
            let _3007_v103 = _compr_108;
            if ((function () {
              let _coll110 = new _dafny.Map();
              for (const _compr_110 of _dafny.IntegerRange(new BigNumber(897), new BigNumber(68))) {
                let _3006_v104 = _compr_110;
                if (((new BigNumber(897)).isLessThanOrEqualTo(_3006_v104)) && ((_3006_v104).isLessThan(new BigNumber(68)))) {
                  _coll110.push([_module.__default.safeModuloInt(_3006_v104, _3002_p0),_3003_v3.f11]);
                }
              }
              return _coll110;
            }()).contains(_3007_v103)) {
              _coll108.push([_module.__default.safeDivisionInt(_3007_v103, _this.f17),new BigNumber((_dafny.Seq.UnicodeFromString("sfj")).length)]);
            }
          }
          return _coll108;
        }());
      })(_2885_v27, p0, _2840_v3);
      let _nw430 = Array((new BigNumber(25)).toNumber());
      for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw430.length); _i0_79++) {
        _nw430[_i0_79] = _init79(new BigNumber(_i0_79));
      }
      _3000_v105 = _nw430;
      let _3008_v106;
      let _nw431 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _3008_v106 = _nw431;
      let _3009_v107;
      _3009_v107 = _dafny.Map.Empty.slice().updateUnsafe(_3008_v106,p1);
      let _pat_let_tv115 = p1;
      let _pat_let_tv116 = _2840_v3;
      let _pat_let_tv117 = p0;
      let _rhs447 = _3000_v105;
      let _rhs448 = ((_this.f16) ? (((_2840_v3.f11) ? ((_dafny.ZERO).minus(p1)) : (_this.f17))) : (p0));
      let _rhs449 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nosncc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(548)), ((_3010_v31) => function (_3011_i10) {
        return _3010_v31;
      })(_2893_v31))), _module.__default.safeIndex(new BigNumber((_3009_v107).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nosncc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(548)), ((_3012_v31) => function (_3013_i10) {
        return _3012_v31;
      })(_2893_v31)))).length)), _2893_v31);
      let _rhs450 = (_dafny.ZERO).minus(function (_source43) {
        if (_source43.is_DC41) {
          let _3014___mcc_h20 = (_source43).cf65;
          let _3015_cf65 = _3014___mcc_h20;
          return _pat_let_tv115;
        } else if (_source43.is_DC40) {
          let _3016___mcc_h21 = (_source43).cf64;
          let _3017_cf64 = _3016___mcc_h21;
          return new BigNumber((_dafny.MultiSet.fromElements(_pat_let_tv116.f11)).cardinality());
        } else {
          let _3018___mcc_h22 = (_source43).cf66;
          let _3019_cf66 = _3018___mcc_h22;
          return _pat_let_tv117;
        }
      }(_module.D15.create_DC41(_this.f11)));
      let _lhs293 = _this;
      _3000_v105 = _rhs447;
      _lhs293.f17 = _rhs448;
      _2839_v2 = _rhs449;
      r1 = _rhs450;
      let _3020_v108;
      let _nw432 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _3020_v108 = _nw432;
      let _index418 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_3020_v108).length));
      (_3020_v108)[_index418] = _module.__default.fm25(_this.f17, _this.f17, globalState);
      let _index419 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_3020_v108).length));
      let _rhs451 = _module.__default.fm2(globalState);
      let _rhs452 = _2836_v0;
      let _rhs453 = _2893_v31;
      let _rhs454 = _2840_v3.f11;
      let _lhs294 = _3020_v108;
      let _lhs295 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_3020_v108).length));
      let _lhs296 = globalState;
      r1 = _rhs451;
      _2836_v0 = _rhs452;
      _lhs294[_lhs295] = _rhs453;
      _lhs296.f8 = _rhs454;
      let _3021_v109;
      _3021_v109 = _dafny.Map.Empty.slice().updateUnsafe(_2882_v24,p0);
      let _3022_v110;
      _3022_v110 = _module.D0.create_DC2((p0).multipliedBy(new BigNumber((_3021_v109).length)), _3008_v106);
      r0 = _3022_v110;
      let _3023_v111;
      _3023_v111 = _dafny.Map.Empty.slice().updateUnsafe(_2839_v2,p2);
      r1 = new BigNumber((_dafny.Seq.Concat(p2, (((_3023_v111).contains(_dafny.Seq.UnicodeFromString("lqst"))) ? ((_3023_v111).get(_dafny.Seq.UnicodeFromString("lqst"))) : (_dafny.Seq.of(_this.f11, _this.f16))))).length);
      let _3024_v112;
      let _nw433 = new _module.C5();
      _nw433.__ctor(new BigNumber((_2883_v25).length));
      _3024_v112 = _nw433;
      let _3025_v113;
      _3025_v113 = _dafny.Set.fromElements(_3024_v112);
      let _3026_v114;
      _3026_v114 = _dafny.Map.Empty.slice().updateUnsafe(true,_2885_v27);
      let _3027_v115;
      _3027_v115 = _dafny.Map.Empty.slice().updateUnsafe(_3025_v113,(_3026_v114).Merge(_3026_v114));
      r2 = (((_3027_v115).contains(_3025_v113)) ? ((_3027_v115).get(_3025_v113)) : (_3026_v114));
      return [r0, r1, r2];
    }
  };

  $module.C16 = class C16 {
    constructor () {
      this._tname = "_module.C16";
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this._f14 = false;
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f14, f15, f11, f12, f13) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1((_this).f13, _dafny.Seq.UnicodeFromString("agcumto"), _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f15), (_this).f12), _this.f11);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return !((_this).f14);
    };
    fm7(globalState) {
      let _this = this;
      return !(!(new BigNumber(917)).isEqualTo(_module.__default.safeModuloInt((_this).f13, (_this).f15)));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      if (_this.f11) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_3028_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_3029_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }));
      } else {
        return _dafny.Seq.UnicodeFromString("qcgbxqc");
      }
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return ((_this).f13).isLessThan((_this).f15);
    };
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = _module.D0.Default();
      let r3 = _dafny.Map.Empty;
      let _3030_v0;
      let _init80 = function (_3031_i0) {
        return (_3031_i0).minus((_this).f15);
      };
      let _nw434 = Array((new BigNumber(7)).toNumber());
      for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw434.length); _i0_80++) {
        _nw434[_i0_80] = _init80(new BigNumber(_i0_80));
      }
      _3030_v0 = _nw434;
      let _3032_v1;
      let _nw435 = Array((new BigNumber(11)).toNumber());
      _nw435[(_dafny.ZERO).toNumber()] = _3030_v0;
      _nw435[(_dafny.ONE).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(2)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(3)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(4)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(5)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(6)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(7)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(8)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(9)).toNumber()] = _3030_v0;
      _nw435[(new BigNumber(10)).toNumber()] = _3030_v0;
      _3032_v1 = _nw435;
      let _3033_v2;
      _3033_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-465),true);
      let _3034_v3;
      let _nw436 = new _module.C2();
      _nw436.__ctor((_this).f15, p0, (_this).f12, new BigNumber((_3033_v2).length), true);
      _3034_v3 = _nw436;
      let _3035_v4;
      _3035_v4 = _dafny.Seq.of(_3034_v3);
      let _3036_v5;
      _3036_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3035_v4).length),(_3034_v3).f13);
      let _3037_v6;
      _3037_v6 = _module.D7.create_DC23(_3032_v1, _dafny.Seq.of(_this.f11), (_this).f14, new BigNumber((_3036_v5).length));
      let _3038_v7;
      let _nw437 = new _module.C11();
      _nw437.__ctor((_this).f15, (_3037_v6).dtor_cf45);
      _3038_v7 = _nw437;
      let _3039_v8;
      let _nw438 = new _module.C14();
      _nw438.__ctor(!((_this).f14));
      _3039_v8 = _nw438;
      let _3040_v9;
      _3040_v9 = _dafny.MultiSet.fromElements(_3039_v8);
      let _3041_v10;
      _3041_v10 = _dafny.Set.fromElements(_3040_v9);
      let _3042_i1;
      _3042_i1 = _dafny.ZERO;
      L15: {
        while ((_3041_v10).IsSubsetOf(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_3039_v8)))) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3042_i1)) {
              break L15;
            }
            _3042_i1 = (_3042_i1).plus(_dafny.ONE);
            (_this).f11 = _this.f11;
            let _3043_v11;
            _3043_v11 = new BigNumber(963);
            _3043_v11 = (_3034_v3).f13;
            _3043_v11 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f13));
            if (_this.f11) {
              let _index420 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_3030_v0).length));
              (_3030_v0)[_index420] = _module.__default.safeModuloInt((_3034_v3).f12, (_3034_v3).f12);
              let _index421 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_3030_v0).length));
              (_3030_v0)[_index421] = (_3038_v7).fm14(globalState);
              let _3044_v12;
              _3044_v12 = _dafny.Seq.of((_this).f14);
              _3044_v12 = _3044_v12;
              let _3045_v13;
              _3045_v13 = _dafny.Set.fromElements(new BigNumber(838), (_this).f13);
              let _3046_v14;
              _3046_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(((_3034_v3).f12).isEqualTo(new BigNumber((_3045_v13).length))),(_3034_v3).f20);
              _3046_v14 = (_3046_v14).update((_this).f14, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-176)), function (_3047_i2) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              }), _module.__default.safeIndex((_3034_v3).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-176)), function (_3048_i2) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              })).length)), (p0)[_module.__default.safeIndex(new BigNumber((_3044_v12).length), new BigNumber((p0).length))]));
              let _3049_v15;
              let _nw439 = Array((new BigNumber(3)).toNumber()).fill([]);
              _3049_v15 = _nw439;
              let _3050_v16;
              let _init81 = function (_3051_i3) {
                return _this.f11;
              };
              let _nw440 = Array((new BigNumber(20)).toNumber());
              for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw440.length); _i0_81++) {
                _nw440[_i0_81] = _init81(new BigNumber(_i0_81));
              }
              _3050_v16 = _nw440;
              let _index422 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_3049_v15).length));
              (_3049_v15)[_index422] = _3050_v16;
              let _index423 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_3030_v0).length));
              let _index424 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_3049_v15).length));
              let _rhs455 = _this.f11;
              let _rhs456 = (_3034_v3).f13;
              let _rhs457 = _3050_v16;
              let _rhs458 = _this.f11;
              let _lhs297 = globalState;
              let _lhs298 = _3030_v0;
              let _lhs299 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_3030_v0).length));
              let _lhs300 = _3049_v15;
              let _lhs301 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_3049_v15).length));
              let _lhs302 = globalState;
              _lhs297.f8 = _rhs455;
              _lhs298[_lhs299] = _rhs456;
              _lhs300[_lhs301] = _rhs457;
              _lhs302.f8 = _rhs458;
              let _3052_v17;
              _3052_v17 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),(_3034_v3).f12);
              let _3053_v18;
              _3053_v18 = new _dafny.CodePoint('n'.codePointAt(0));
              _3052_v17 = (_3052_v17).update(_3053_v18, (_3034_v3).f12);
            } else {
              let _3054_v19;
              let _init82 = ((_3055_v3, _3056_v2) => function (_3057_i4) {
                return _dafny.Seq.of((_3055_v3).f12, (_this).f13, new BigNumber((_3056_v2).length));
              })(_3034_v3, _3033_v2);
              let _nw441 = Array((new BigNumber(8)).toNumber());
              for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw441.length); _i0_82++) {
                _nw441[_i0_82] = _init82(new BigNumber(_i0_82));
              }
              _3054_v19 = _nw441;
              let _3058_v20;
              _3058_v20 = _dafny.Seq.of(new BigNumber(((_3034_v3).f20).length));
              let _index425 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_3054_v19).length));
              (_3054_v19)[_index425] = _dafny.Seq.Concat(_3058_v20, _3058_v20);
              let _index426 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_3054_v19).length));
              (_3054_v19)[_index426] = _3058_v20;
              let _3059_v21;
              _3059_v21 = new _dafny.CodePoint('w'.codePointAt(0));
              let _3060_v22;
              let _nw442 = new _module.C5();
              _nw442.__ctor((((_this).f14) ? ((_3034_v3).f12) : (new BigNumber((_dafny.Seq.of(_3059_v21, _3059_v21, _3059_v21)).length))));
              _3060_v22 = _nw442;
              let _index427 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_3030_v0).length));
              (_3030_v0)[_index427] = new BigNumber(709);
              let _index428 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_3030_v0).length));
              (_3030_v0)[_index428] = (_dafny.ZERO).minus(((_this).f12).minus(((_this.f11) ? ((_this).f13) : (new BigNumber((_dafny.MultiSet.fromElements(_3043_v11, (_3060_v22).f24)).cardinality())))));
              let _3061_v23;
              let _nw443 = Array((new BigNumber(7)).toNumber()).fill(false);
              _3061_v23 = _nw443;
              let _3062_v24;
              _3062_v24 = _dafny.MultiSet.fromElements(_3061_v23);
              let _3063_v25;
              let _nw444 = new _module.C12();
              _nw444.__ctor(_3062_v24, (_this).f12, _this.f11, (_3030_v0)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_3030_v0).length))], (_3034_v3).f12);
              _3063_v25 = _nw444;
              let _3064_v26;
              _3064_v26 = ((_this.f11) ? (_3063_v25) : (_3063_v25));
              let _3065_v27;
              let _nw445 = new _module.C10();
              _nw445.__ctor((_this).f14);
              _3065_v27 = _nw445;
              let _3066_v28;
              _3066_v28 = _dafny.Map.Empty.slice().updateUnsafe(_3065_v27,_3063_v25);
              _3064_v26 = (((_3066_v28).contains(_3065_v27)) ? ((_3066_v28).get(_3065_v27)) : (_3063_v25));
              let _3067_v29;
              _3067_v29 = _module.D6.create_DC18(_this.f11);
              let _3068_v30;
              let _nw446 = Array((_dafny.ONE).toNumber());
              _nw446[(_dafny.ZERO).toNumber()] = _3067_v29;
              _3068_v30 = _nw446;
              let _3069_v31;
              _3069_v31 = _dafny.MultiSet.fromElements(_3068_v30);
              let _3070_v32;
              _3070_v32 = _module.D11.create_DC30((_3069_v31).Difference(_3069_v31));
              let _3071_v33;
              let _nw447 = Array((new BigNumber(7)).toNumber()).fill([]);
              _3071_v33 = _nw447;
              let _index429 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_3061_v23).length));
              (_3061_v23)[_index429] = _this.f11;
              let _3072_v34;
              _3072_v34 = _dafny.Map.Empty.slice().updateUnsafe(_3061_v23,(_dafny.ZERO).minus((_this).f13));
              let _3073_v35;
              _3073_v35 = _dafny.Set.fromElements((_this).f13, (_3034_v3).f13);
              let _3074_v37;
              _3074_v37 = _dafny.MultiSet.fromElements(_3059_v21);
              let _3075_v38;
              _3075_v38 = _dafny.MultiSet.fromElements((_this).f15, new BigNumber((_3033_v2).length), (_this).f13, (((_3074_v37).contains(_3059_v21)) ? ((_3074_v37).get(_3059_v21)) : ((_this).f12)), (_3038_v7).fm14(globalState));
              let _3076_v39;
              let _nw448 = new _module.C15();
              _nw448.__ctor(_this.f11, (_3034_v3).f12, _this.f11);
              _3076_v39 = _nw448;
              let _3077_v40;
              _3077_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_3076_v39);
              let _index430 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_3030_v0).length));
              let _index431 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_3061_v23).length));
              let _rhs459 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_3054_v19)[_module.__default.safeIndex(new BigNumber(707), new BigNumber((_3054_v19).length))], _module.__default.fm0(globalState)), (_3054_v19)[_module.__default.safeIndex(new BigNumber(707), new BigNumber((_3054_v19).length))]);
              let _rhs460 = _3070_v32;
              let _rhs461 = _3071_v33;
              let _rhs462 = new BigNumber(((_3072_v34).Merge(_3072_v34)).length);
              let _rhs463 = (_3039_v8).fm6((_3073_v35).equals(function () {
                let _coll111 = new _dafny.Set();
                for (const _compr_111 of _dafny.IntegerRange(new BigNumber(-330), new BigNumber(581))) {
                  let _3078_v36 = _compr_111;
                  if (((new BigNumber(-330)).isLessThanOrEqualTo(_3078_v36)) && ((_3078_v36).isLessThan(new BigNumber(581)))) {
                    _coll111.add(_module.__default.safeDivisionInt(_3078_v36, (_this).f12));
                  }
                }
                return _coll111;
              }()), (_3034_v3).f13, (((_3075_v38).contains((_dafny.ZERO).minus(new BigNumber((_3077_v40).length)))) ? ((_3075_v38).get((_dafny.ZERO).minus(new BigNumber((_3077_v40).length)))) : ((_3034_v3).f13)), globalState);
              let _lhs303 = globalState;
              let _lhs304 = _3030_v0;
              let _lhs305 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_3030_v0).length));
              let _lhs306 = _3061_v23;
              let _lhs307 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_3061_v23).length));
              _lhs303.f8 = _rhs459;
              _3070_v32 = _rhs460;
              _3071_v33 = _rhs461;
              _lhs304[_lhs305] = _rhs462;
              _lhs306[_lhs307] = _rhs463;
            }
          }
        }
      }
      let _hi13 = (_3034_v3).f13;
      for (let _3079_i5 = (_this).f12; _3079_i5.isLessThan(_hi13); _3079_i5 = _3079_i5.plus(_dafny.ONE)) {
        let _index432 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
        (_3030_v0)[_index432] = (_this).f12;
        let _index433 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
        (_3030_v0)[_index433] = (_this).f15;
        let _3080_v41;
        let _nw449 = Array((new BigNumber(15)).toNumber()).fill(false);
        _3080_v41 = _nw449;
        _3080_v41 = _3080_v41;
        if ((_3034_v3).fm7(globalState)) {
          let _3081_v42;
          let _init83 = function (_3082_i6) {
            return (_3082_i6).plus((_this).f13);
          };
          let _nw450 = Array((new BigNumber(13)).toNumber());
          for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw450.length); _i0_83++) {
            _nw450[_i0_83] = _init83(new BigNumber(_i0_83));
          }
          _3081_v42 = _nw450;
          _3081_v42 = _3081_v42;
          let _3083_v43;
          _3083_v43 = _dafny.Set.fromElements((_3034_v3).f13);
          let _3084_v44;
          _3084_v44 = _dafny.Map.Empty.slice().updateUnsafe(_3081_v42,_3083_v43);
          _3083_v43 = (((_3084_v44).contains((((_this).f14) ? (_3081_v42) : (_3081_v42)))) ? ((_3084_v44).get((((_this).f14) ? (_3081_v42) : (_3081_v42)))) : ((_3083_v43).Union(_3083_v43)));
          let _3085_v45;
          _3085_v45 = _dafny.Set.fromElements(_this.f11);
          let _3086_v46;
          _3086_v46 = _module.D19.create_DC51(new BigNumber((_3085_v45).length), (_3034_v3).f13);
          let _3087_v47;
          _3087_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_3086_v46);
          _3087_v47 = (_3087_v47).update((_this).f14, _module.D19.create_DC51((_this).f13, (_3034_v3).f13));
          let _3088_v48;
          _3088_v48 = _dafny.Seq.UnicodeFromString("sutgkklap");
          let _3089_v49;
          _3089_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(596)), function (_3090_i7) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length));
          let _3091_v50;
          _3091_v50 = new _dafny.CodePoint('p'.codePointAt(0));
          let _index434 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _rhs464 = _module.__default.safeDivisionInt(((_this).f12).multipliedBy((_this).f12), (_3034_v3).f12);
          let _rhs465 = (_module.__default.fm47(_this.f11, false, _3089_v49, (_3034_v3).f12, globalState)).Merge(_3033_v2);
          let _rhs466 = (_this).f14;
          let _rhs467 = !(((_3030_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length))]).plus((_this).f15)).isEqualTo((((_this).f14) ? (_module.__default.fm2(globalState)) : ((_3034_v3).f13)));
          let _rhs468 = _dafny.Seq.Concat(_3088_v48, _dafny.Seq.update(_dafny.Seq.Concat((_3034_v3).f20, _dafny.Seq.UnicodeFromString("ggglnxi")), _module.__default.safeIndex((_3034_v3).f13, new BigNumber((_dafny.Seq.Concat((_3034_v3).f20, _dafny.Seq.UnicodeFromString("ggglnxi"))).length)), _3091_v50));
          let _lhs308 = _3030_v0;
          let _lhs309 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          _lhs308[_lhs309] = _rhs464;
          r3 = _rhs465;
          r0 = _rhs466;
          r0 = _rhs467;
          _3088_v48 = _rhs468;
          let _3092_v51;
          _3092_v51 = _dafny.Seq.of(_this.f11, (_this).f14, (_this).f14, true, _this.f11);
          (_this).f11 = !_dafny.Seq.contains(_dafny.Seq.Concat(_3092_v51, _dafny.Seq.update(_dafny.Seq.of(_this.f11, _this.f11), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.of(_this.f11, _this.f11)).length)), true)), (_3089_v49).contains(_this.f11));
        } else {
          let _index435 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          (_3080_v41)[_index435] = (((_this).fm9((_this).f14, (_this).fm7(globalState), globalState)) ? ((_this).f14) : (!(_this.f11)));
          let _index436 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          (_3080_v41)[_index436] = _this.f11;
          let _3093_v52;
          let _nw451 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
          _3093_v52 = _nw451;
          let _3094_v53;
          _3094_v53 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState));
          let _index437 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_3093_v52).length));
          (_3093_v52)[_index437] = _3094_v53;
          let _3095_v54;
          let _nw452 = Array((new BigNumber(5)).toNumber()).fill(_module.D0.Default());
          _3095_v54 = _nw452;
          let _index438 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_3095_v54).length));
          (_3095_v54)[_index438] = (((_3080_v41)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length))]) ? (p1) : (p1));
          let _3096_v55;
          _3096_v55 = _dafny.Seq.of(_this.f11, true, (_this).f14);
          let _3097_v56;
          _3097_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3096_v55).length),_3094_v53);
          let _index439 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          let _index440 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_3093_v52).length));
          let _index441 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          let _index442 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_3095_v54).length));
          let _rhs469 = (_this).f14;
          let _rhs470 = (((_3097_v56).contains(_module.__default.safeModuloInt(new BigNumber(777), (_this).f12))) ? ((_3097_v56).get(_module.__default.safeModuloInt(new BigNumber(777), (_this).f12))) : ((_3094_v53).update((_3034_v3).f12, _module.__default.abs(new BigNumber((_3096_v55).length)))));
          let _rhs471 = (_this).f14;
          let _rhs472 = p1;
          let _lhs310 = _3080_v41;
          let _lhs311 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          let _lhs312 = _3093_v52;
          let _lhs313 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_3093_v52).length));
          let _lhs314 = _3080_v41;
          let _lhs315 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length));
          let _lhs316 = _3095_v54;
          let _lhs317 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_3095_v54).length));
          _lhs310[_lhs311] = _rhs469;
          _lhs312[_lhs313] = _rhs470;
          _lhs314[_lhs315] = _rhs471;
          _lhs316[_lhs317] = _rhs472;
          let _3098_v57;
          let _init84 = function (_3099_i8) {
            return false;
          };
          let _nw453 = Array((new BigNumber(2)).toNumber());
          for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw453.length); _i0_84++) {
            _nw453[_i0_84] = _init84(new BigNumber(_i0_84));
          }
          _3098_v57 = _nw453;
          let _3100_v58;
          _3100_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14);
          _3098_v57 = (((((_3100_v58).contains((_3080_v41)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length))])) ? ((_3100_v58).get((_3080_v41)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length))])) : (_this.f11))) ? (_3098_v57) : (_3098_v57));
          let _3101_v59;
          let _nw454 = new _module.C3();
          _nw454.__ctor((_3034_v3).f13, new BigNumber((_dafny.Seq.Concat(_3096_v55, _3096_v55)).length));
          _3101_v59 = _nw454;
          let _3102_v60;
          _3102_v60 = _dafny.Seq.of(_3096_v55);
          let _3103_v61;
          let _nw455 = new _module.C10();
          _nw455.__ctor(_this.f11);
          _3103_v61 = _nw455;
          let _3104_v62;
          _3104_v62 = _module.D6.create_DC19((_this).f14, false, false, _this.f11, (_3102_v60)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3103_v61,(_this).f14)).length), new BigNumber((_3102_v60).length))]);
          let _3105_v63;
          _3105_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_3104_v62);
          let _nw456 = new _module.C7();
          _nw456.__ctor(new BigNumber(((_3105_v63).Merge(_dafny.Map.Empty.slice().updateUnsafe((_3080_v41)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_3080_v41).length))],_3104_v62))).length), _module.__default.safeModuloInt((((_3094_v53).contains((_this).f12)) ? ((_3094_v53).get((_this).f12)) : ((_3034_v3).f12)), (_3034_v3).f12));
          _3101_v59 = _nw456;
          let _index443 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          (_3030_v0)[_index443] = ((((_3093_v52)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_3093_v52).length))]).contains((_3034_v3).f13)) ? (((_3093_v52)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_3093_v52).length))]).get((_3034_v3).f13)) : (new BigNumber((_dafny.Seq.UnicodeFromString("bthuoibuq")).length)));
        }
        if (_this.f11) {
          let _index444 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3080_v41).length));
          (_3080_v41)[_index444] = ((_3034_v3).f12).isEqualTo((_this).f13);
          let _3106_v64;
          _3106_v64 = _dafny.Set.fromElements(_this.f11, _this.f11);
          let _3107_v65;
          _3107_v65 = _dafny.Map.Empty.slice().updateUnsafe(_3033_v2,(_this).f14);
          let _3108_v66;
          _3108_v66 = _module.D20.create_DC52(_3107_v65);
          let _3109_v67;
          _3109_v67 = _module.D20.create_DC54(_3108_v66);
          let _3110_v68;
          _3110_v68 = _dafny.MultiSet.fromElements(_3109_v67);
          let _3111_v69;
          _3111_v69 = _dafny.Seq.of((_this).f12, (_3034_v3).f12, new BigNumber(572), new BigNumber((_3110_v68).cardinality()), (_3034_v3).f13);
          let _3112_v70;
          _3112_v70 = _dafny.Seq.of((_this).f14, (_this).f14);
          let _3113_v71;
          _3113_v71 = _dafny.Set.fromElements(_3112_v70, _3112_v70);
          let _index445 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _index446 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3080_v41).length));
          let _index447 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _index448 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _rhs473 = (_3034_v3).f13;
          let _rhs474 = (_this).f14;
          let _rhs475 = (new BigNumber(((_3106_v64).Difference(_dafny.Set.fromElements(_this.f11, true))).length)).minus((_3034_v3).f13);
          let _rhs476 = (((((_this).f14) ? (_this.f11) : (_this.f11))) ? (_module.__default.safeDivisionInt((_3034_v3).f13, (_3111_v69)[_module.__default.safeIndex((_this).f15, new BigNumber((_3111_v69).length))])) : (new BigNumber(847)));
          let _rhs477 = !(_3113_v71).contains(_3112_v70);
          let _lhs318 = _3030_v0;
          let _lhs319 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _lhs320 = _3080_v41;
          let _lhs321 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3080_v41).length));
          let _lhs322 = _3030_v0;
          let _lhs323 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _lhs324 = _3030_v0;
          let _lhs325 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          let _lhs326 = globalState;
          _lhs318[_lhs319] = _rhs473;
          _lhs320[_lhs321] = _rhs474;
          _lhs322[_lhs323] = _rhs475;
          _lhs324[_lhs325] = _rhs476;
          _lhs326.f8 = _rhs477;
          let _3114_v72;
          let _nw457 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _3114_v72 = _nw457;
          let _index449 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_3032_v1).length));
          (_3032_v1)[_index449] = _3114_v72;
          let _index450 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_3032_v1).length));
          (_3032_v1)[_index450] = _3114_v72;
          _3036_v5 = (_3036_v5).update((_this).f12, ((_this.f11) ? ((_this).f13) : (new BigNumber((p0).length))));
          let _index451 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          (_3030_v0)[_index451] = _module.__default.safeModuloInt((_3034_v3).f13, new BigNumber(-80));
          let _index452 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3080_v41).length));
          (_3080_v41)[_index452] = (_3080_v41)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_3080_v41).length))];
        } else {
          let _3115_v73;
          _3115_v73 = _dafny.MultiSet.fromElements((_3034_v3).f12);
          let _3116_v74;
          let _nw458 = Array((new BigNumber(2)).toNumber());
          _nw458[(_dafny.ZERO).toNumber()] = new BigNumber((_3115_v73).cardinality());
          _nw458[(_dafny.ONE).toNumber()] = (_3034_v3).f13;
          _3116_v74 = _nw458;
          let _3117_v75;
          _3117_v75 = _dafny.MultiSet.fromElements((_module.D0.create_DC2((_this).f12, _3116_v74)).dtor_cf6, _3116_v74);
          let _index453 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length));
          (_3030_v0)[_index453] = (_dafny.ZERO).minus((((_3117_v75).contains(_3116_v74)) ? ((_3117_v75).get(_3116_v74)) : ((((_3115_v73).contains((_3034_v3).f13)) ? ((_3115_v73).get((_3034_v3).f13)) : ((_3034_v3).f12)))));
          let _3118_v76;
          _3118_v76 = _dafny.Seq.of((_3034_v3).f20, (_3034_v3).f20);
          let _3119_v78;
          _3119_v78 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),(_3030_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_3030_v0).length))]);
          let _3120_v79;
          _3120_v79 = _dafny.Seq.of(new BigNumber((_3119_v78).length));
          let _3121_v80;
          _3121_v80 = _dafny.Seq.of(_this.f11);
          let _3122_v81;
          _3122_v81 = _dafny.Set.fromElements(!((_3121_v80)[_module.__default.safeIndex((_3034_v3).f12, new BigNumber((_3121_v80).length))]), _this.f11);
          let _3123_v82;
          _3123_v82 = _dafny.Map.Empty.slice().updateUnsafe((_3118_v76)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll112 = new _dafny.Map();
            for (const _compr_112 of (_3120_v79).Elements) {
              let _3124_v77 = _compr_112;
              if (_dafny.Seq.contains(_3120_v79, _3124_v77)) {
                _coll112.push([_module.__default.safeModuloInt(_3124_v77, (_this).f15),new BigNumber((_dafny.Set.fromElements(_this.f11, _this.f11, _this.f11)).length)]);
              }
            }
            return _coll112;
          }()).length), new BigNumber((_3118_v76).length))],_3122_v81);
          _3123_v82 = (_3123_v82).Merge(_3123_v82);
          let _3125_v83;
          let _init85 = ((_3126_p0) => function (_3127_i9) {
            return _3126_p0;
          })(p0);
          let _nw459 = Array((new BigNumber(24)).toNumber());
          for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw459.length); _i0_85++) {
            _nw459[_i0_85] = _init85(new BigNumber(_i0_85));
          }
          _3125_v83 = _nw459;
          let _index454 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_3125_v83).length));
          (_3125_v83)[_index454] = _dafny.Seq.Concat((_3034_v3).f20, p0);
          let _index455 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_3125_v83).length));
          (_3125_v83)[_index455] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), function (_3128_i10) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          });
          _3039_v8 = _3039_v8;
          let _3129_v84;
          _3129_v84 = _module.D7.create_DC22((_3034_v3).f13, _3032_v1);
          let _pat_let_tv118 = _3030_v0;
          let _pat_let_tv119 = _3030_v0;
          let _3130_v85;
          let _nw460 = Array((new BigNumber(15)).toNumber());
          _nw460[(_dafny.ZERO).toNumber()] = function (_pat_let77_0) {
            return function (_3131_dt__update__tmp_h1) {
              return function (_pat_let78_0) {
                return function (_3132_dt__update_hcf40_h0) {
                  return _module.D7.create_DC22(_3132_dt__update_hcf40_h0, (_3131_dt__update__tmp_h1).dtor_cf41);
                }(_pat_let78_0);
              }((_pat_let_tv119)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_pat_let_tv118).length))]);
            }(_pat_let77_0);
          }(_3129_v84);
          _nw460[(_dafny.ONE).toNumber()] = function (_pat_let79_0) {
            return function (_3133_dt__update__tmp_h2) {
              return function (_pat_let80_0) {
                return function (_3134_dt__update_hcf40_h1) {
                  return _module.D7.create_DC22(_3134_dt__update_hcf40_h1, (_3133_dt__update__tmp_h2).dtor_cf41);
                }(_pat_let80_0);
              }((_this).f12);
            }(_pat_let79_0);
          }(_module.D7.create_DC22((_3034_v3).f13, _3032_v1));
          _nw460[(new BigNumber(2)).toNumber()] = _module.D7.create_DC22((_3038_v7).fm14(globalState), _3032_v1);
          _nw460[(new BigNumber(3)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(4)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(5)).toNumber()] = _module.D7.create_DC22((_this).f12, _3032_v1);
          _nw460[(new BigNumber(6)).toNumber()] = _module.D7.create_DC22((_3034_v3).f13, _3032_v1);
          _nw460[(new BigNumber(7)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(8)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(9)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(10)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(11)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(12)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(13)).toNumber()] = _3129_v84;
          _nw460[(new BigNumber(14)).toNumber()] = _3129_v84;
          _3130_v85 = _nw460;
          let _3135_v86;
          _3135_v86 = _dafny.MultiSet.fromElements(_3130_v85);
          _3135_v86 = _3135_v86;
        }
      }
      let _3136_i11;
      _3136_i11 = _dafny.ZERO;
      L16: {
        while (!(_this.f11) || ((_this).f14)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3136_i11)) {
              break L16;
            }
            _3136_i11 = (_3136_i11).plus(_dafny.ONE);
            let _3137_v87;
            let _nw461 = Array((new BigNumber(2)).toNumber());
            _nw461[(_dafny.ZERO).toNumber()] = (_this).f14;
            _nw461[(_dafny.ONE).toNumber()] = _this.f11;
            _3137_v87 = _nw461;
            let _index456 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_3137_v87).length));
            (_3137_v87)[_index456] = false;
            let _3138_v88;
            _3138_v88 = _module.D0.create_DC1((_this).f15, (_3034_v3).f20, (_3034_v3).f12, true);
            let _index457 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_3137_v87).length));
            (_3137_v87)[_index457] = _dafny.Seq.IsPrefixOf(((_this.f11) ? ((_3034_v3).f20) : ((_3034_v3).f20)), (_3138_v88).dtor_cf2);
            let _index458 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_3032_v1).length));
            (_3032_v1)[_index458] = _3030_v0;
            let _index459 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_3032_v1).length));
            let _init86 = ((_3139_v3) => function (_3140_i12) {
              return _module.__default.safeModuloInt(_3140_i12, (_3139_v3).f13);
            })(_3034_v3);
            let _nw462 = Array((new BigNumber(26)).toNumber());
            for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw462.length); _i0_86++) {
              _nw462[_i0_86] = _init86(new BigNumber(_i0_86));
            }
            (_3032_v1)[_index459] = _nw462;
            let _3141_v89;
            let _nw463 = new _module.C3();
            _nw463.__ctor(((_3034_v3).f13).minus((_this).f12), new BigNumber(743));
            _3141_v89 = _nw463;
            let _3142_v90;
            _3142_v90 = _module.D20.create_DC53();
            _3142_v90 = (((_this).f14) ? (_3142_v90) : (_3142_v90));
          }
        }
      }
      let _3143_v91;
      _3143_v91 = _dafny.Seq.of((_3034_v3).f13);
      let _3144_v92;
      let _nw464 = new _module.C8();
      _nw464.__ctor(_3143_v91);
      _3144_v92 = _nw464;
      let _3145_v93;
      _3145_v93 = new BigNumber(736);
      let _3146_v94;
      _3146_v94 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_3034_v3).f12),_3030_v0);
      _3145_v93 = ((_dafny.ZERO).minus((_3034_v3).f13)).plus(new BigNumber((_3146_v94).length));
      r0 = _this.f11;
      let _3147_v95;
      _3147_v95 = _dafny.MultiSet.fromElements(!(_this.f11), (_3039_v8).fm6((_this).f14, (_3038_v7).fm14(globalState), (_this).f12, globalState), _this.f11);
      r1 = _3147_v95;
      let _3148_v96;
      let _nw465 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
      _3148_v96 = _nw465;
      r2 = _module.D0.create_DC0(_3148_v96);
      r3 = ((_3033_v2).Merge(_3033_v2)).Merge((((_this).f14) ? (_3033_v2) : (_3033_v2)));
      return [r0, r1, r2, r3];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C17 = class C17 {
    constructor () {
      this._tname = "_module.C17";
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.D0.create_DC1(new BigNumber(((_dafny.Set.fromElements(new BigNumber(-82))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("vldiolxvg")).length), new BigNumber((_dafny.Seq.of(new BigNumber(262), new BigNumber(497))).length))).cardinality()), new BigNumber(390), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-67),new BigNumber(526))).length), new BigNumber(122), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber(725)),new BigNumber(-479))).length), new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f11, _this.f11))).cardinality()))), new BigNumber(36))).length))).length)))).length), _dafny.Seq.UnicodeFromString("xnpq"), (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ahuf")).length), new BigNumber((function () {
  let _coll113 = new _dafny.Map();
  for (const _compr_113 of (function () {
    let _coll114 = new _dafny.Map();
    for (const _compr_114 of _dafny.IntegerRange(new BigNumber(-642), new BigNumber(865))) {
      let _3149_v1 = _compr_114;
      if (((new BigNumber(-642)).isLessThanOrEqualTo(_3149_v1)) && ((_3149_v1).isLessThan(new BigNumber(865)))) {
        _coll114.push([(_3149_v1).plus(new BigNumber(504)),_dafny.Seq.of(_this.f11)]);
      }
    }
    return _coll114;
  }()).Keys.Elements) {
    let _3150_v0 = _compr_113;
    if ((function () {
      let _coll115 = new _dafny.Map();
      for (const _compr_115 of _dafny.IntegerRange(new BigNumber(-642), new BigNumber(865))) {
        let _3149_v1 = _compr_115;
        if (((new BigNumber(-642)).isLessThanOrEqualTo(_3149_v1)) && ((_3149_v1).isLessThan(new BigNumber(865)))) {
          _coll115.push([(_3149_v1).plus(new BigNumber(504)),_dafny.Seq.of(_this.f11)]);
        }
      }
      return _coll115;
    }()).contains(_3150_v0)) {
      _coll113.push([(_3150_v0).plus(new BigNumber(770)),new BigNumber(-852)]);
    }
  }
  return _coll113;
}()).length)))), _this.f11);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f11;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = false;
      let _3151_v0;
      _3151_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      if (((((_3151_v0).contains(_this.f11)) ? ((_3151_v0).get(_this.f11)) : (p2))) && (p0)) {
        let _3152_v1;
        _3152_v1 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p1, new BigNumber(-950)),p2);
        _3152_v1 = (_3152_v1).update(p1, p2);
        let _3153_v2;
        _3153_v2 = _dafny.Seq.of(p1);
        let _3154_v3;
        _3154_v3 = _dafny.Seq.UnicodeFromString("kgsgelena");
        let _3155_v4;
        let _nw466 = new _module.C2();
        _nw466.__ctor(new BigNumber((_dafny.Seq.Concat(_3153_v2, _3153_v2)).length), _3154_v3, p1, new BigNumber((_3153_v2).length), _this.f11);
        _3155_v4 = _nw466;
        let _3156_v5;
        let _nw467 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _3156_v5 = _nw467;
        let _index460 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_3156_v5).length));
        (_3156_v5)[_index460] = _dafny.Seq.of(p0, (_3155_v4).fm7(globalState), _this.f11);
        let _3157_v6;
        _3157_v6 = _dafny.Seq.of(false, (_3155_v4).fm33((_3155_v4).f26, globalState));
        let _index461 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_3156_v5).length));
        (_3156_v5)[_index461] = _dafny.Seq.Concat(_3157_v6, _dafny.Seq.Concat(_3157_v6, _dafny.Seq.of(p2, _this.f11)));
        let _3158_v7;
        _3158_v7 = new _dafny.CodePoint('v'.codePointAt(0));
        let _3159_v8;
        _3159_v8 = _dafny.Map.Empty.slice().updateUnsafe(_3158_v7,p0);
        let _3160_v9;
        _3160_v9 = _dafny.Set.fromElements(new BigNumber(157), p1);
        let _rhs478 = (_3155_v4).fm8((_3155_v4).f26, _3159_v8, _3160_v9, globalState);
        let _rhs479 = _dafny.Seq.IsProperPrefixOf(_3153_v2, _dafny.Seq.Concat(_dafny.Seq.of((_3155_v4).f26), _3153_v2));
        let _lhs327 = globalState;
        _3154_v3 = _rhs478;
        _lhs327.f8 = _rhs479;
        r0 = p0;
      } else {
        let _3161_v10;
        _3161_v10 = _module.D7.create_DC21(_dafny.MultiSet.fromElements((_this).fm6(_this.f11, p1, p1, globalState), p0));
        let _3162_v11;
        _3162_v11 = _dafny.Map.Empty.slice().updateUnsafe(_3161_v10,!(p0) || (p0));
        let _3163_v12;
        let _nw468 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _3163_v12 = _nw468;
        let _3164_v13;
        _3164_v13 = _module.D0.create_DC0(_3163_v12);
        let _3165_v14;
        _3165_v14 = _dafny.Map.Empty.slice().updateUnsafe(_3164_v13,p0);
        let _3166_v15;
        _3166_v15 = _dafny.Seq.of(_3165_v14, _3165_v14, _dafny.Map.Empty.slice().updateUnsafe(_3164_v13,_this.f11));
        let _3167_v16;
        _3167_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_3166_v15)[_module.__default.safeIndex(p1, new BigNumber((_3166_v15).length))]).length));
        let _3168_v17;
        let _nw469 = new _module.C4();
        _nw469.__ctor(p1, p0);
        _3168_v17 = _nw469;
        let _3169_v19;
        _3169_v19 = _dafny.Map.Empty.slice().updateUnsafe((_3168_v17).f25,p1);
        let _3170_v20;
        _3170_v20 = _dafny.MultiSet.fromElements(true);
        let _3171_v21;
        _3171_v21 = _dafny.Seq.UnicodeFromString("w");
        let _3172_v22;
        let _nw470 = Array((new BigNumber(22)).toNumber());
        _nw470[(_dafny.ZERO).toNumber()] = new BigNumber(-444);
        _nw470[(_dafny.ONE).toNumber()] = p1;
        _nw470[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_3168_v17)).length);
        _nw470[(new BigNumber(3)).toNumber()] = new BigNumber(374);
        _nw470[(new BigNumber(4)).toNumber()] = p1;
        _nw470[(new BigNumber(5)).toNumber()] = new BigNumber((function () {
          let _coll116 = new _dafny.Map();
          for (const _compr_116 of (_3169_v19).Keys.Elements) {
            let _3173_v18 = _compr_116;
            if ((_3169_v19).contains(_3173_v18)) {
              _coll116.push([_module.__default.safeDivisionInt(_3173_v18, (_3168_v17).f25),p0]);
            }
          }
          return _coll116;
        }()).length);
        _nw470[(new BigNumber(6)).toNumber()] = p1;
        _nw470[(new BigNumber(7)).toNumber()] = new BigNumber(287);
        _nw470[(new BigNumber(8)).toNumber()] = new BigNumber((_3170_v20).cardinality());
        _nw470[(new BigNumber(9)).toNumber()] = (_3168_v17).f25;
        _nw470[(new BigNumber(10)).toNumber()] = p1;
        _nw470[(new BigNumber(11)).toNumber()] = (_3168_v17).f25;
        _nw470[(new BigNumber(12)).toNumber()] = p1;
        _nw470[(new BigNumber(13)).toNumber()] = (_3168_v17).f25;
        _nw470[(new BigNumber(14)).toNumber()] = new BigNumber((_3171_v21).length);
        _nw470[(new BigNumber(15)).toNumber()] = p1;
        _nw470[(new BigNumber(16)).toNumber()] = p1;
        _nw470[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_3168_v17).f25);
        _nw470[(new BigNumber(18)).toNumber()] = (_3168_v17).f25;
        _nw470[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw470[(new BigNumber(20)).toNumber()] = p1;
        _nw470[(new BigNumber(21)).toNumber()] = (_3168_v17).f25;
        _3172_v22 = _nw470;
        let _3174_v23;
        let _nw471 = Array((new BigNumber(4)).toNumber());
        _nw471[(_dafny.ZERO).toNumber()] = _3172_v22;
        _nw471[(_dafny.ONE).toNumber()] = _3172_v22;
        _nw471[(new BigNumber(2)).toNumber()] = _3172_v22;
        _nw471[(new BigNumber(3)).toNumber()] = _3172_v22;
        _3174_v23 = _nw471;
        let _3175_v24;
        _3175_v24 = _module.D7.create_DC22(new BigNumber((_dafny.MultiSet.fromElements(p1, p1, new BigNumber(-252), p1)).cardinality()), _3174_v23);
        let _3176_v25;
        _3176_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
        let _rhs480 = (((_3167_v16).contains((_3151_v0).contains(p2))) ? ((_3167_v16).get((_3151_v0).contains(p2))) : (p1));
        let _rhs481 = (_3162_v11).Merge(_3162_v11);
        let _rhs482 = (_module.D9.create_DC27(p0, _3175_v24, new BigNumber((_3176_v25).length), _this.f11)).dtor_cf51;
        let _rhs483 = (_3168_v17).f25;
        r1 = _rhs480;
        _3162_v11 = _rhs481;
        r1 = _rhs482;
        r1 = _rhs483;
        let _3177_v26;
        _3177_v26 = new _dafny.CodePoint('m'.codePointAt(0));
        let _3178_v27;
        _3178_v27 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_this.f11) : (true)),_3177_v26);
        r2 = (((_3178_v27).contains(p0)) ? ((_3178_v27).get(p0)) : (_module.__default.fm45(p1, globalState)));
        let _3179_v28;
        _3179_v28 = _dafny.Seq.of((_3168_v17).f25, p1);
        let _rhs484 = p0;
        let _rhs485 = _dafny.Seq.Concat(((p2) ? (_3179_v28) : (_3179_v28)), _dafny.Seq.Concat(_3179_v28, _dafny.Seq.of(p1, (_3168_v17).f25)));
        let _rhs486 = p1;
        let _rhs487 = ((_3168_v17).f25).plus((_dafny.ZERO).minus(new BigNumber(((_3170_v20).Union(_3170_v20)).cardinality())));
        let _rhs488 = (_3179_v28)[_module.__default.safeIndex((_3168_v17).f25, new BigNumber((_3179_v28).length))];
        let _lhs328 = globalState;
        let _lhs329 = globalState;
        _lhs328.f8 = _rhs484;
        _lhs329.f5 = _rhs485;
        r1 = _rhs486;
        r1 = _rhs487;
        r1 = _rhs488;
        let _3180_v29;
        _3180_v29 = _module.D0.create_DC2(p1, _3172_v22);
        r1 = (((false) ? ((_3180_v29).dtor_cf5) : (p1))).minus((_3168_v17).f25);
        let _index462 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_3172_v22).length));
        (_3172_v22)[_index462] = _module.__default.safeModuloInt((((_3167_v16).contains(true)) ? ((_3167_v16).get(true)) : (p1)), (_dafny.ZERO).minus(p1));
        let _3181_v30;
        _3181_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p0)).length),p0);
        let _3182_v31;
        _3182_v31 = _dafny.Map.Empty.slice().updateUnsafe(_3181_v30,p2);
        let _3183_v32;
        _3183_v32 = _module.D20.create_DC52(_3182_v31);
        let _3184_v33;
        _3184_v33 = _dafny.Set.fromElements(p0);
        let _3185_v34;
        _3185_v34 = _dafny.Map.Empty.slice().updateUnsafe(_3183_v32,(_3184_v33).IsSubsetOf(_3184_v33));
        let _3186_v35;
        _3186_v35 = _dafny.Set.fromElements((_3168_v17).f25, (_3168_v17).f25);
        let _3187_v36;
        _3187_v36 = _module.D17.create_DC45(new BigNumber((_3179_v28).length), _3186_v35);
        let _3188_v37;
        let _nw472 = new _module.C11();
        _nw472.__ctor(_module.__default.fm2(globalState), _module.__default.fm2(globalState));
        _3188_v37 = _nw472;
        let _3189_v38;
        _3189_v38 = _dafny.Seq.of(_3188_v37);
        let _3190_v39;
        _3190_v39 = _dafny.MultiSet.fromElements(new BigNumber(937), _module.__default.safeModuloInt((((_3169_v19).contains((_3168_v17).f25)) ? ((_3169_v19).get((_3168_v17).f25)) : (p1)), (_3168_v17).f25), ((_3168_v17).f25).multipliedBy((_3187_v36).dtor_cf69), new BigNumber((_3189_v38).length));
        let _index463 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_3172_v22).length));
        let _rhs489 = (_3188_v37).fm14(globalState);
        let _rhs490 = _this.f11;
        let _rhs491 = _3185_v34;
        let _rhs492 = ((_this.f11) ? (_module.__default.fm12((_3168_v17).f25, (_3168_v17).f25, p1, p2, globalState)) : ((_3190_v39).Difference(_3190_v39)));
        let _rhs493 = !(p2);
        let _lhs330 = _3172_v22;
        let _lhs331 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_3172_v22).length));
        let _lhs332 = globalState;
        _lhs330[_lhs331] = _rhs489;
        r0 = _rhs490;
        _3185_v34 = _rhs491;
        _3190_v39 = _rhs492;
        _lhs332.f8 = _rhs493;
      }
      let _3191_v40;
      _3191_v40 = _dafny.Set.fromElements(p1, new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality()), p1, p1);
      let _3192_v41;
      let _nw473 = Array((new BigNumber(24)).toNumber());
      _3192_v41 = _nw473;
      let _3193_v42;
      _3193_v42 = _dafny.Seq.UnicodeFromString("od");
      let _rhs494 = (_3191_v40).Intersect(_3191_v40);
      let _rhs495 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_3193_v42).length), p1), p1));
      let _rhs496 = ((new BigNumber(814)).isEqualTo(new BigNumber(-31))) === (p2);
      let _rhs497 = _3192_v41;
      let _lhs333 = globalState;
      _3191_v40 = _rhs494;
      r1 = _rhs495;
      _lhs333.f8 = _rhs496;
      _3192_v41 = _rhs497;
      let _hi14 = _module.__default.safeModuloInt(p1, p1);
      for (let _3194_i0 = p1; _3194_i0.isLessThan(_hi14); _3194_i0 = _3194_i0.plus(_dafny.ONE)) {
        r1 = _3194_i0;
        if ((_3194_i0).isLessThanOrEqualTo(p1)) {
          let _index464 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((p3).length));
          (p3)[_index464] = p2;
          let _index465 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((p3).length));
          (p3)[_index465] = false;
          let _3195_v43;
          let _3196_v44;
          let _3197_v45;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector21 = (_this).m2(globalState);
          _out45 = _outcollector21[0];
          _out46 = _outcollector21[1];
          _out47 = _outcollector21[2];
          _3195_v43 = _out45;
          _3196_v44 = _out46;
          _3197_v45 = _out47;
          let _3198_v46;
          let _3199_v47;
          let _3200_v48;
          let _out48;
          let _out49;
          let _out50;
          let _outcollector22 = (_this).m2(globalState);
          _out48 = _outcollector22[0];
          _out49 = _outcollector22[1];
          _out50 = _outcollector22[2];
          _3198_v46 = _out48;
          _3199_v47 = _out49;
          _3200_v48 = _out50;
          let _3201_v49;
          _3201_v49 = new _dafny.CodePoint('e'.codePointAt(0));
          let _3202_v50;
          _3202_v50 = _dafny.Set.fromElements(_3201_v49, _3201_v49);
          _3202_v50 = _3202_v50;
          let _3203_v51;
          _3203_v51 = _module.D9.create_DC26(_3201_v49);
          _3203_v51 = _3203_v51;
        } else {
          r1 = _3194_i0;
          let _3204_v52;
          _3204_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_3194_i0);
          _3204_v52 = (_3204_v52).update(p0, _module.__default.safeDivisionInt(_3194_i0, _3194_i0));
          let _index466 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((p3).length));
          (p3)[_index466] = (p2) && (!(p0));
          let _3205_v53;
          _3205_v53 = new _dafny.CodePoint('x'.codePointAt(0));
          let _index467 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((p3).length));
          (p3)[_index467] = _dafny.Seq.contains(_3193_v42, _3205_v53);
          let _3206_v54;
          _3206_v54 = _dafny.Seq.of(_3193_v42, _dafny.Seq.UnicodeFromString("fhlhneom"), _dafny.Seq.UnicodeFromString("wktwjoaxr"));
          let _index468 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((p3).length));
          let _index469 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((p3).length));
          let _rhs498 = _this.f11;
          let _rhs499 = _dafny.areEqual(_3206_v54, ((_this.f11) ? (_3206_v54) : (_3206_v54)));
          let _rhs500 = (((_this.f11) || (_this.f11)) ? (p2) : (p0));
          let _lhs334 = globalState;
          let _lhs335 = p3;
          let _lhs336 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((p3).length));
          let _lhs337 = p3;
          let _lhs338 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((p3).length));
          _lhs334.f8 = _rhs498;
          _lhs335[_lhs336] = _rhs499;
          _lhs337[_lhs338] = _rhs500;
          let _3207_v55;
          _3207_v55 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.fm2(globalState)), p1);
          let _3208_v56;
          _3208_v56 = _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_3193_v42, _dafny.Seq.UnicodeFromString("hhfpf")), p2, _module.__default.fm1(new BigNumber((_3207_v55).length), true, _3194_i0, globalState));
          let _3209_v57;
          _3209_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_3208_v56);
          _3208_v56 = (_3208_v56).Union((((_3209_v57).contains((p3)[_module.__default.safeIndex(new BigNumber(903), new BigNumber((p3).length))])) ? ((_3209_v57).get((p3)[_module.__default.safeIndex(new BigNumber(903), new BigNumber((p3).length))])) : (_3208_v56)));
          r1 = (_3194_i0).minus(_3194_i0);
        }
        let _3210_v58;
        _3210_v58 = _3191_v40;
        _3210_v58 = _3210_v58;
        let _3211_v59;
        _3211_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _3212_v60;
        _3212_v60 = _dafny.Set.fromElements(((p0) ? (p0) : (p2)), p0, (((_3211_v59).contains(p1)) ? ((_3211_v59).get(p1)) : (_this.f11)));
        let _3213_v61;
        _3213_v61 = _module.D0.create_DC3(p1, _this.f11);
        let _3214_v62;
        _3214_v62 = _dafny.MultiSet.fromElements(p2, _this.f11, p0, p0);
        let _3215_v63;
        _3215_v63 = new _dafny.CodePoint('q'.codePointAt(0));
        _3212_v60 = _module.__default.fm4(_3194_i0, _3213_v61, (((_3214_v62).contains(_this.f11)) ? ((_3214_v62).get(_this.f11)) : (_3194_i0)), _3215_v63, globalState);
      }
      let _3216_v64;
      _3216_v64 = _module.D20.create_DC53();
      let _3217_v65;
      let _nw474 = Array((new BigNumber(5)).toNumber());
      _nw474[(_dafny.ZERO).toNumber()] = _3216_v64;
      _nw474[(_dafny.ONE).toNumber()] = _3216_v64;
      _nw474[(new BigNumber(2)).toNumber()] = _3216_v64;
      _nw474[(new BigNumber(3)).toNumber()] = _3216_v64;
      _nw474[(new BigNumber(4)).toNumber()] = _3216_v64;
      _3217_v65 = _nw474;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_3217_v65).length))) {
        let _3218_i1 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_3218_i1)) && ((_3218_i1).isLessThan(new BigNumber((_3217_v65).length))))) {
          (_3217_v65)[(_3218_i1)] = _3216_v64;
        }
      }
      let _3219_v66;
      _3219_v66 = _dafny.Set.fromElements(p2);
      let _3220_v67;
      _3220_v67 = _dafny.Map.Empty.slice().updateUnsafe(_3219_v66,p2);
      let _3221_v68;
      _3221_v68 = new _dafny.CodePoint('y'.codePointAt(0));
      let _3222_v69;
      _3222_v69 = _dafny.Seq.of(p1, p1);
      let _3223_v70;
      _3223_v70 = _dafny.Map.Empty.slice().updateUnsafe(p2,_3222_v69);
      let _3224_v71;
      _3224_v71 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), ((_3225_v68) => function (_3226_i2) {
        return _3225_v68;
      })(_3221_v68)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), ((_3227_v68) => function (_3228_i2) {
        return _3227_v68;
      })(_3221_v68))).length)), _3221_v68)).length));
      let _3229_v72;
      let _nw475 = Array((new BigNumber(23)).toNumber());
      _nw475[(_dafny.ZERO).toNumber()] = new BigNumber(834);
      _nw475[(_dafny.ONE).toNumber()] = p1;
      _nw475[(new BigNumber(2)).toNumber()] = p1;
      _nw475[(new BigNumber(3)).toNumber()] = ((_module.__default.fm1(p1, p0, p1, globalState)) ? (p1) : (p1));
      _nw475[(new BigNumber(4)).toNumber()] = new BigNumber(969);
      _nw475[(new BigNumber(5)).toNumber()] = p1;
      _nw475[(new BigNumber(6)).toNumber()] = p1;
      _nw475[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_3193_v42).length), (_dafny.ZERO).minus(p1));
      _nw475[(new BigNumber(8)).toNumber()] = p1;
      _nw475[(new BigNumber(9)).toNumber()] = (p1).minus(_module.__default.fm2(globalState));
      _nw475[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_3220_v67).length));
      _nw475[(new BigNumber(11)).toNumber()] = p1;
      _nw475[(new BigNumber(12)).toNumber()] = ((p0) ? (p1) : (p1));
      _nw475[(new BigNumber(13)).toNumber()] = new BigNumber(727);
      _nw475[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.of(_3221_v68, _3221_v68)).length);
      _nw475[(new BigNumber(15)).toNumber()] = new BigNumber((_3193_v42).length);
      _nw475[(new BigNumber(16)).toNumber()] = new BigNumber((_3223_v70).length);
      _nw475[(new BigNumber(17)).toNumber()] = (p1).minus(p1);
      _nw475[(new BigNumber(18)).toNumber()] = p1;
      _nw475[(new BigNumber(19)).toNumber()] = (p1).plus(p1);
      _nw475[(new BigNumber(20)).toNumber()] = new BigNumber((_3193_v42).length);
      _nw475[(new BigNumber(21)).toNumber()] = p1;
      _nw475[(new BigNumber(22)).toNumber()] = new BigNumber((_3224_v71).length);
      _3229_v72 = _nw475;
      _3229_v72 = _3229_v72;
      (_this).f11 = (_3191_v40).IsSubsetOf(_3191_v40);
      r0 = (p1).isLessThan(p1);
      r1 = p1;
      r2 = _3221_v68;
      let _3230_v73;
      let _nw476 = new _module.C4();
      _nw476.__ctor(p1, _this.f11);
      _3230_v73 = _nw476;
      let _3231_v74;
      _3231_v74 = _dafny.Set.fromElements(_3230_v73, _3230_v73);
      let _3232_v75;
      _3232_v75 = _dafny.Seq.of(_dafny.Set.fromElements(_3230_v73));
      r3 = (_3231_v74).IsDisjointFrom((_3232_v75)[_module.__default.safeIndex(p1, new BigNumber((_3232_v75).length))]);
      return [r0, r1, r2, r3];
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _3233_v0;
      _3233_v0 = new BigNumber(585);
      let _hi15 = _3233_v0;
      for (let _3234_i0 = _3233_v0; _3234_i0.isLessThan(_hi15); _3234_i0 = _3234_i0.plus(_dafny.ONE)) {
        (globalState).f8 = _this.f11;
        let _3235_v1;
        _3235_v1 = new _dafny.CodePoint('b'.codePointAt(0));
        let _3236_v2;
        _3236_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,new _dafny.CodePoint('p'.codePointAt(0)));
        let _3237_v3;
        _3237_v3 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_3235_v1), _3236_v2, _3236_v2, _3236_v2);
        let _3238_v4;
        _3238_v4 = _dafny.MultiSet.fromElements(((_3236_v2).update(_this.f11, _3235_v1)).update(_this.f11, _3235_v1), _3236_v2, _3236_v2, (_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_3235_v1)).update(_this.f11, _3235_v1), _3236_v2);
        (globalState).f8 = !(_3237_v3).equals((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_3235_v1))).Union(function () {
          let _coll117 = new _dafny.Set();
          for (const _compr_117 of (_3238_v4).Elements) {
            let _3239_v5 = _compr_117;
            if ((_3238_v4).contains(_3239_v5)) {
              _coll117.add(_3239_v5);
            }
          }
          return _coll117;
        }()));
        let _3240_v6;
        let _3241_v7;
        let _out51;
        let _out52;
        let _outcollector23 = _module.__default.m0(globalState);
        _out51 = _outcollector23[0];
        _out52 = _outcollector23[1];
        _3240_v6 = _out51;
        _3241_v7 = _out52;
        let _rhs501 = _3241_v7;
        let _rhs502 = (_3234_i0).minus(_3233_v0);
        let _lhs339 = globalState;
        _lhs339.f8 = _rhs501;
        r2 = _rhs502;
      }
      let _3242_v8;
      _3242_v8 = new _dafny.CodePoint('k'.codePointAt(0));
      let _3243_v9;
      _3243_v9 = _dafny.MultiSet.fromElements(_3242_v8);
      r2 = _module.__default.safeModuloInt(_module.__default.fm2(globalState), new BigNumber(((_3243_v9)).cardinality()));
      let _3244_i1;
      _3244_i1 = _dafny.ZERO;
      L17: {
        while (_this.f11) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3244_i1)) {
              break L17;
            }
            _3244_i1 = (_3244_i1).plus(_dafny.ONE);
            _3233_v0 = _3233_v0;
            if (_this.f11) {
              r1 = (_this.f11) === (!(_3233_v0).isEqualTo(new BigNumber(465)));
              let _3245_v10;
              _3245_v10 = _module.D19.create_DC51(_3233_v0, _3233_v0);
              let _3246_v11;
              let _nw477 = new _module.C7();
              _nw477.__ctor(((true) ? ((_3245_v10).dtor_cf77) : (_3233_v0)), (_3233_v0).plus(_3233_v0));
              _3246_v11 = _nw477;
              let _3247_v12;
              let _nw478 = Array((new BigNumber(22)).toNumber()).fill(false);
              _3247_v12 = _nw478;
              let _index470 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_3247_v12).length));
              (_3247_v12)[_index470] = _this.f11;
              let _3248_v13;
              _3248_v13 = _dafny.Seq.of(_3247_v12);
              let _3249_v14;
              _3249_v14 = _dafny.Map.Empty.slice().updateUnsafe(_3233_v0,(new BigNumber((_3248_v13).length)).multipliedBy(_3233_v0));
              let _3250_v15;
              _3250_v15 = _dafny.Seq.of(_this.f11);
              let _3251_v16;
              _3251_v16 = _dafny.Seq.of(_dafny.Seq.Concat(_3250_v15, _3250_v15), _dafny.Seq.Concat(_3250_v15, _3250_v15), _3250_v15, _3250_v15, _3250_v15);
              let _3252_v17;
              _3252_v17 = _dafny.Seq.of(_3251_v16);
              let _index471 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_3247_v12).length));
              let _rhs503 = _this.f11;
              let _rhs504 = _3249_v14;
              let _rhs505 = _dafny.Seq.Concat(_dafny.Seq.Concat(_3251_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_3253_v15) => function (_3254_i2) {
                return _3253_v15;
              })(_3250_v15))), (_3252_v17)[_module.__default.safeIndex(_3233_v0, new BigNumber((_3252_v17).length))]);
              let _lhs340 = _3247_v12;
              let _lhs341 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_3247_v12).length));
              _lhs340[_lhs341] = _rhs503;
              _3249_v14 = _rhs504;
              _3251_v16 = _rhs505;
              _3233_v0 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_3233_v0), new BigNumber(424));
              let _3255_v18;
              let _init87 = function (_3256_i3) {
                return _module.D18.create_DC48();
              };
              let _nw479 = Array((new BigNumber(16)).toNumber());
              for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw479.length); _i0_87++) {
                _nw479[_i0_87] = _init87(new BigNumber(_i0_87));
              }
              _3255_v18 = _nw479;
              let _index472 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_3255_v18).length));
              (_3255_v18)[_index472] = _module.__default.fm64(globalState);
              let _index473 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_3255_v18).length));
              (_3255_v18)[_index473] = _module.__default.fm64(globalState);
            } else {
              let _3257_v19;
              _3257_v19 = _dafny.Set.fromElements(new BigNumber(-628), _3233_v0, _3233_v0, _3233_v0, _3233_v0);
              let _3258_v20;
              _3258_v20 = _dafny.Map.Empty.slice().updateUnsafe(_3257_v19,_3242_v8);
              let _3259_v21;
              _3259_v21 = _dafny.Map.Empty.slice().updateUnsafe(_3233_v0,_this.f11);
              (globalState).f8 = !(_module.__default.fm1(new BigNumber(((_3258_v20).Merge(_3258_v20)).length), ((_this.f11) ? ((((_3259_v21).contains(_3233_v0)) ? ((_3259_v21).get(_3233_v0)) : ((_this).fm6(_this.f11, _3233_v0, (_dafny.ZERO).minus(_3233_v0), globalState)))) : (_this.f11)), _3233_v0, globalState));
              let _3260_v22;
              _3260_v22 = _dafny.Seq.UnicodeFromString("eintml");
              let _3261_v23;
              _3261_v23 = _dafny.Seq.of(_dafny.Seq.Concat(_3260_v22, _3260_v22));
              _3260_v22 = (_3261_v23)[_module.__default.safeIndex(_3233_v0, new BigNumber((_3261_v23).length))];
              let _3262_v24;
              let _nw480 = new _module.C3();
              _nw480.__ctor(_3233_v0, _module.__default.safeModuloInt(_3233_v0, _3233_v0));
              _3262_v24 = _nw480;
              let _3263_v25;
              _3263_v25 = _dafny.Seq.of(_3257_v19);
              let _rhs506 = _module.__default.safeDivisionInt(_3233_v0, _3233_v0);
              let _rhs507 = _3260_v22;
              let _rhs508 = _this.f11;
              let _rhs509 = _3262_v24;
              let _rhs510 = _dafny.Seq.update(_dafny.Seq.Concat(_3263_v25, _3263_v25), _module.__default.safeIndex(_3233_v0, new BigNumber((_dafny.Seq.Concat(_3263_v25, _3263_v25)).length)), _3257_v19);
              r2 = _rhs506;
              _3260_v22 = _rhs507;
              r0 = _rhs508;
              _3262_v24 = _rhs509;
              _3263_v25 = _rhs510;
              let _3264_v26;
              _3264_v26 = _dafny.Seq.of(new BigNumber(585), (_dafny.ZERO).minus(_3233_v0));
              let _3265_v27;
              let _nw481 = new _module.C5();
              _nw481.__ctor((_3264_v26)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3242_v8,_3233_v0)).length), new BigNumber((_3264_v26).length))]);
              _3265_v27 = _nw481;
              _3242_v8 = _module.__default.fm25((_3265_v27).f24, _3233_v0, globalState);
            }
            _3233_v0 = _3233_v0;
            (globalState).f8 = _this.f11;
          }
        }
      }
      let _3266_v28;
      _3266_v28 = _dafny.Set.fromElements(_3233_v0);
      let _3267_v29;
      _3267_v29 = _3266_v28;
      let _3268_v30;
      let _nw482 = Array((_dafny.ONE).toNumber()).fill(false);
      _3268_v30 = _nw482;
      let _3269_v31;
      _3269_v31 = _dafny.Set.fromElements(_3268_v30);
      let _3270_v32;
      _3270_v32 = _dafny.MultiSet.fromElements(new BigNumber((_3269_v31).length));
      let _3271_v33;
      _3271_v33 = _module.D5.create_DC15(_3233_v0, _3233_v0, new BigNumber((_dafny.Set.fromElements(_3233_v0)).length));
      let _3272_v34;
      _3272_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_3270_v32).cardinality())),_3271_v33);
      let _3273_v35;
      _3273_v35 = _dafny.Seq.UnicodeFromString("idy");
      let _3274_v36;
      _3274_v36 = _dafny.Map.Empty.slice().updateUnsafe(_3233_v0,_dafny.Seq.UnicodeFromString("roaishuns"));
      _3267_v29 = _module.__default.fm65(_module.D6.create_DC16(_3272_v34), _3273_v35, _3274_v36, globalState);
      let _3275_i4;
      _3275_i4 = _dafny.ZERO;
      L18: {
        while (_module.__default.fm1(_3233_v0, _this.f11, _3233_v0, globalState)) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3275_i4)) {
              break L18;
            }
            _3275_i4 = (_3275_i4).plus(_dafny.ONE);
            let _3276_v37;
            _3276_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("kupca"),_3268_v30);
            let _3277_v38;
            _3277_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_3278_v8) => function (_3279_i5) {
              return _3278_v8;
            })(_3242_v8)),_3276_v37);
            let _3280_v39;
            _3280_v39 = _module.D24.create_DC60(_3276_v37);
            let _3281_v40;
            _3281_v40 = _dafny.Seq.of(_3276_v37, _3276_v37, _3276_v37);
            let _3282_v41;
            let _nw483 = Array((new BigNumber(29)).toNumber());
            _nw483[(_dafny.ZERO).toNumber()] = _3276_v37;
            _nw483[(_dafny.ONE).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(2)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(3)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(4)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(5)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(6)).toNumber()] = (_3276_v37).Merge(_3276_v37);
            _nw483[(new BigNumber(7)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(8)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(9)).toNumber()] = (((_3277_v38).contains(_module.__default.fm36(_3233_v0, _3233_v0, _3242_v8, globalState))) ? ((_3277_v38).get(_module.__default.fm36(_3233_v0, _3233_v0, _3242_v8, globalState))) : (_3276_v37));
            _nw483[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_3273_v35,_3268_v30)).Merge(_3276_v37);
            _nw483[(new BigNumber(11)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(12)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(13)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(14)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(15)).toNumber()] = (_3276_v37).Merge(_3276_v37);
            _nw483[(new BigNumber(16)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(17)).toNumber()] = (_3276_v37).Merge(_3276_v37);
            _nw483[(new BigNumber(18)).toNumber()] = (_3280_v39).dtor_cf88;
            _nw483[(new BigNumber(19)).toNumber()] = ((!(_this.f11)) ? ((_3280_v39).dtor_cf88) : (_3276_v37));
            _nw483[(new BigNumber(20)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(21)).toNumber()] = (_3276_v37).Merge(_3276_v37);
            _nw483[(new BigNumber(22)).toNumber()] = (_3281_v40)[_module.__default.safeIndex(_3233_v0, new BigNumber((_3281_v40).length))];
            _nw483[(new BigNumber(23)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_3273_v35,_3268_v30);
            _nw483[(new BigNumber(25)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(26)).toNumber()] = _3276_v37;
            _nw483[(new BigNumber(27)).toNumber()] = (_3276_v37).Merge(_3276_v37);
            _nw483[(new BigNumber(28)).toNumber()] = ((_this.f11) ? (_3276_v37) : (_3276_v37));
            _3282_v41 = _nw483;
            let _index474 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_3282_v41).length));
            (_3282_v41)[_index474] = _3276_v37;
            let _index475 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_3282_v41).length));
            (_3282_v41)[_index475] = _3276_v37;
            let _3283_v42;
            _3283_v42 = _dafny.MultiSet.fromElements(_this.f11, _this.f11);
            let _3284_v43;
            _3284_v43 = _dafny.Seq.of(_this.f11, (_dafny.MultiSet.fromElements(_this.f11)).IsSubsetOf(_3283_v42), _this.f11, false);
            _3284_v43 = _dafny.Seq.Concat(_3284_v43, _3284_v43);
            if (_this.f11) {
              let _3285_v44;
              let _nw484 = new _module.C14();
              _nw484.__ctor(_this.f11);
              _3285_v44 = _nw484;
              (globalState).f8 = _this.f11;
              let _3286_v45;
              let _nw485 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
              _3286_v45 = _nw485;
              let _3287_v46;
              let _nw486 = new _module.C15();
              _nw486.__ctor(_this.f11, _3233_v0, _this.f11);
              _3287_v46 = _nw486;
              let _3288_v47;
              _3288_v47 = _dafny.Map.Empty.slice().updateUnsafe(_3287_v46,_3283_v42);
              let _index476 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_3286_v45).length));
              (_3286_v45)[_index476] = _3288_v47;
              let _index477 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_3286_v45).length));
              (_3286_v45)[_index477] = _3288_v47;
              let _3289_v48;
              let _nw487 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _3289_v48 = _nw487;
              let _3290_v49;
              _3290_v49 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_3289_v48,_3287_v46.f17)).contains(_3289_v48),_3233_v0);
              _3290_v49 = (_3290_v49).update((_this.f11) === ((_this).fm6(_module.__default.fm1(_3233_v0, _module.__default.fm1(_3233_v0, true, _3287_v46.f17, globalState), _3233_v0, globalState), _3233_v0, new BigNumber((_3270_v32).cardinality()), globalState)), (_dafny.ZERO).minus(_3233_v0));
              (_this).f11 = _this.f11;
            } else {
              let _3291_v50;
              let _nw488 = new _module.C6();
              _nw488.__ctor();
              _3291_v50 = _nw488;
              _3291_v50 = _3291_v50;
              let _rhs511 = _this.f11;
              let _rhs512 = _3233_v0;
              let _lhs342 = _this;
              _lhs342.f11 = _rhs511;
              _3233_v0 = _rhs512;
              let _3292_v51;
              _3292_v51 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),_3233_v0);
              let _3293_v52;
              _3293_v52 = _dafny.Map.Empty.slice().updateUnsafe(_3233_v0,_this.f11);
              let _3294_v53;
              _3294_v53 = _dafny.Set.fromElements(_this.f11);
              let _3295_v54;
              let _nw489 = Array((new BigNumber(20)).toNumber());
              _nw489[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_3273_v35, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_3296_v8) => function (_3297_i6) {
                return _3296_v8;
              })(_3242_v8)), _module.__default.safeIndex(_3233_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_3298_v8) => function (_3299_i6) {
                return _3298_v8;
              })(_3242_v8))).length)), _3242_v8))).length);
              _nw489[(_dafny.ONE).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(2)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(3)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((((_3270_v32).contains(_3233_v0)) ? ((_3270_v32).get(_3233_v0)) : (_3233_v0)), _3233_v0);
              _nw489[(new BigNumber(5)).toNumber()] = new BigNumber((_3283_v42).cardinality());
              _nw489[(new BigNumber(6)).toNumber()] = (new BigNumber(396)).plus(_3233_v0);
              _nw489[(new BigNumber(7)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(8)).toNumber()] = _module.__default.fm2(globalState);
              _nw489[(new BigNumber(9)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(10)).toNumber()] = (((_3292_v51).contains((((_3293_v52).contains(_3233_v0)) ? ((_3293_v52).get(_3233_v0)) : (!(_this.f11))))) ? ((_3292_v51).get((((_3293_v52).contains(_3233_v0)) ? ((_3293_v52).get(_3233_v0)) : (!(_this.f11))))) : (_3233_v0));
              _nw489[(new BigNumber(11)).toNumber()] = _module.__default.fm2(globalState);
              _nw489[(new BigNumber(12)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm0(globalState)).length);
              _nw489[(new BigNumber(14)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(15)).toNumber()] = ((_dafny.ZERO).minus(_3233_v0)).minus(new BigNumber((_3294_v53).length));
              _nw489[(new BigNumber(16)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm36(_3233_v0, _3233_v0, new _dafny.CodePoint('c'.codePointAt(0)), globalState), _3273_v35)).length);
              _nw489[(new BigNumber(18)).toNumber()] = _3233_v0;
              _nw489[(new BigNumber(19)).toNumber()] = new BigNumber((_3270_v32).cardinality());
              _3295_v54 = _nw489;
              let _index478 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_3295_v54).length));
              (_3295_v54)[_index478] = _3233_v0;
              let _index479 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_3295_v54).length));
              (_3295_v54)[_index479] = _3233_v0;
              (globalState).f8 = (new BigNumber((_3292_v51).length)).isLessThan(_3233_v0);
              let _3300_v55;
              _3300_v55 = _dafny.Seq.of((new BigNumber(181)).multipliedBy(new BigNumber((_3292_v51).length)), (_3295_v54)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_3295_v54).length))], (_3295_v54)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_3295_v54).length))]);
              (globalState).f5 = _3300_v55;
            }
            _3233_v0 = _3233_v0;
          }
        }
      }
      let _3301_i7;
      _3301_i7 = _dafny.ZERO;
      L19: {
        while ((_3233_v0).isLessThan(_3233_v0)) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_3301_i7)) {
              break L19;
            }
            _3301_i7 = (_3301_i7).plus(_dafny.ONE);
            let _3302_v56;
            _3302_v56 = _dafny.Seq.of(!(_this.f11));
            let _3303_v57;
            let _nw490 = new _module.C16();
            _nw490.__ctor(_this.f11, _3233_v0, false, _3233_v0, _3233_v0);
            _3303_v57 = _nw490;
            let _3304_v58;
            _3304_v58 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3303_v57,_this.f11)).length));
            let _rhs513 = _dafny.Seq.IsPrefixOf(_3302_v56, _dafny.Seq.of(true));
            let _rhs514 = ((_3233_v0).plus(_3233_v0)).isEqualTo(((_dafny.ZERO).minus(new BigNumber((_3304_v58).length))).multipliedBy((_3303_v57).f15));
            let _rhs515 = _this.f11;
            let _rhs516 = (new BigNumber((_3304_v58).length)).minus(_3233_v0);
            let _lhs343 = globalState;
            r1 = _rhs513;
            r1 = _rhs514;
            _lhs343.f8 = _rhs515;
            r2 = _rhs516;
            let _3305_v59;
            _3305_v59 = _dafny.Seq.of(_3273_v35);
            _3274_v36 = (_3274_v36).update(_3233_v0, (_3305_v59)[_module.__default.safeIndex(_3233_v0, new BigNumber((_3305_v59).length))]);
            let _3306_v60;
            _3306_v60 = _dafny.Map.Empty.slice().updateUnsafe(_3233_v0,_3233_v0);
            let _3307_v61;
            _3307_v61 = _dafny.Seq.update(_3304_v58, _module.__default.safeIndex(_3233_v0, new BigNumber((_3304_v58).length)), new BigNumber((_3306_v60).length));
            let _3308_v62;
            _3308_v62 = _dafny.Set.fromElements(_3307_v61, _3307_v61, _3307_v61);
            let _3309_v63;
            let _nw491 = new _module.C9();
            _nw491.__ctor(_3308_v62, new BigNumber(908), _3233_v0, new BigNumber((_3273_v35).length), (_3303_v57).f14, _3273_v35);
            _3309_v63 = _nw491;
            let _3310_v64;
            _3310_v64 = _dafny.Set.fromElements(_3309_v63, _3309_v63, _3309_v63, _3309_v63);
            (_this).f11 = (_3310_v64).equals((_3310_v64).Union(_3310_v64));
            let _3311_v65;
            let _nw492 = new _module.C6();
            _nw492.__ctor();
            _3311_v65 = _nw492;
          }
        }
      }
      r0 = (_module.__default.safeModuloInt(_3233_v0, new BigNumber((_3273_v35).length))).isLessThan(_3233_v0);
      r1 = (new BigNumber(905)).isLessThanOrEqualTo((_3233_v0).minus(_3233_v0));
      r2 = (_module.__default.fm2(globalState)).minus(_3233_v0);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
