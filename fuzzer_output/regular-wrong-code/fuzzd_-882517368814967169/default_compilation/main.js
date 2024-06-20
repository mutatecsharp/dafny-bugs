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
    static fm2(globalState) {
      if (true) {
        return _dafny.Seq.UnicodeFromString("hwfd");
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), function (_0_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        });
      }
    };
    static fm3(p0, p1, p2, globalState) {
      return _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(844)), function (_1_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }), new _dafny.CodePoint('o'.codePointAt(0)));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1(true, !(false) || (true), new _dafny.CodePoint('t'.codePointAt(0)), (new BigNumber((_dafny.Set.fromElements(true, false)).length)).minus(new BigNumber(31)));
    };
    static fm5(globalState) {
      return (_dafny.Set.fromElements(false, true, true, false)).IsProperSubsetOf(_dafny.Set.fromElements(false, false, false));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), function (_2_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      });
    };
    static fm13(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber(687)).multipliedBy(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).length))),true);
    };
    static fm15(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bkklhmuw"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), function (_3_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })), (_module.D6.create_DC11(_dafny.Seq.UnicodeFromString("rw"))).dtor_cf21);
    };
    static fm16(globalState) {
      return _dafny.Seq.of((new BigNumber((_dafny.Seq.of(new BigNumber(834), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-435))).length))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('j'.codePointAt(0)))).length))));
    };
    static fm17(p0, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))) : (_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)))))).Union((_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0))))));
    };
    static fm18(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("imlqi")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length), new BigNumber(881), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber(-937))).length))).Intersect(function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(104), new BigNumber(923))) {
          let _4_v0 = _compr_0;
          if (((new BigNumber(104)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(923)))) {
            _coll0.add(_module.__default.safeModuloInt(_4_v0, new BigNumber(703)));
          }
        }
        return _coll0;
      }())).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, !(true), false, true)).cardinality()))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(500)))).Union(_dafny.MultiSet.FromArray(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_5_i0) {
        return new BigNumber(-557);
      })) : (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(996),true)).length))))));
    };
    static fm20(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),true));
    };
    static fm21(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ay"), _dafny.Seq.UnicodeFromString("poxqgk")),_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("xqxl"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-618)), function (_6_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), function (_7_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_8_i1) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))));
    };
    static fm23(globalState) {
      return new _dafny.CodePoint('j'.codePointAt(0));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      if (true) {
        return _module.D0.create_DC0(true);
      } else {
        return _module.D0.create_DC0(true);
      }
    };
    static fm26(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D7.create_DC14(_dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), function (_9_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("lvtm")).length);
}));
      if (_source0.is_DC15) {
        let _10___mcc_h0 = (_source0).cf27;
        let _11___mcc_h1 = (_source0).cf28;
        let _12_cf28 = _11___mcc_h1;
        let _13_cf27 = _10___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(true,true);
      } else {
        let _14___mcc_h2 = (_source0).cf26;
        let _15_cf26 = _14___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(!(!(true)),!(false));
      }
    };
    static fm27(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("khmjbaxct")).length),(new BigNumber(696)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)))).cardinality())));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("vxkrw");
    };
    static fm29(globalState) {
      return _module.D5.create_DC10(!((!(false)) && (true)), new BigNumber(117), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rbplxx"),function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(801), new BigNumber(233))) {
    let _16_v0 = _compr_1;
    if (((new BigNumber(801)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(233)))) {
      _coll1.push([_module.__default.safeDivisionInt(_16_v0, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())),true]);
    }
  }
  return _coll1;
}()), !(false));
    };
    static fm30(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("acqcse"))).Difference(_dafny.MultiSet.fromElements((_module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-477)), function (_17_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}))).dtor_cf21, _dafny.Seq.UnicodeFromString("oc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), function (_18_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("teyeal"), _dafny.Seq.UnicodeFromString("jcehpgj")));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(382))).Union((_dafny.Set.fromElements(new BigNumber(962), (_dafny.ZERO).minus(new BigNumber(-593)))).Intersect(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.MultiSet.fromElements(new BigNumber(883))).Elements) {
          let _19_v0 = _compr_2;
          if ((_dafny.MultiSet.fromElements(new BigNumber(883))).contains(_19_v0)) {
            _coll2.add((_19_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).length)));
          }
        }
        return _coll2;
      }()));
    };
    static fm32(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-334))).length),new BigNumber(564))).length)))), _dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-49)), function (_20_i0) {
  return new BigNumber(913);
}))), _dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-550))).length), new BigNumber((_dafny.Seq.of(false, true)).length))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ulsmcuwk")).length))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(150)), function (_21_i1) {
  return _dafny.Set.fromElements(false, true);
})).length),new BigNumber(80))).length), new BigNumber(561)))).cardinality()), new BigNumber(374), new BigNumber(-739), new BigNumber((_dafny.Seq.of(true, false)).length))))));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber(25)))), _dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-788)), function (_22_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(621), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(909)), function (_23_i1) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(782), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(750))).length))).length))))), _dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber(134))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(300))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(-30), new BigNumber(30), new BigNumber((_dafny.Seq.UnicodeFromString("fpgtqpqa")).length), new BigNumber(159)))), _dafny.Seq.of(_module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ppysr")).length), new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber(636), new BigNumber((_dafny.Seq.UnicodeFromString("gvfexkl")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), function (_24_i2) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qmdcartnm")).length))), _module.D7.create_DC14(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-803)), function (_25_i3) {
  return new BigNumber((function () {
    let _coll3 = new _dafny.Set();
    for (const _compr_3 of _dafny.IntegerRange(new BigNumber(577), new BigNumber(-231))) {
      let _26_v0 = _compr_3;
      if (((new BigNumber(577)).isLessThanOrEqualTo(_26_v0)) && ((_26_v0).isLessThan(new BigNumber(-231)))) {
        _coll3.add((_26_v0).plus(new BigNumber(8)));
      }
    }
    return _coll3;
  }()).length);
})), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(-929))), _module.D7.create_DC14(_dafny.Seq.of(new BigNumber(612)))))));
    };
    static fm34(p0, globalState) {
      return (((!(true)) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ckmfb"),false)) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pfwhyaky"),false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), function (_27_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hvc"),!(true))));
    };
    static fm35(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).Difference((_dafny.MultiSet.fromElements(new BigNumber(257), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length))).cardinality()), new BigNumber(482))))).Union((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-548), new BigNumber(152))) {
          let _28_v0 = _compr_4;
          if (((new BigNumber(-548)).isLessThanOrEqualTo(_28_v0)) && ((_28_v0).isLessThan(new BigNumber(152)))) {
            _coll4.push([(_28_v0).minus(new BigNumber(896)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yxd")).length))).cardinality())]);
          }
        }
        return _coll4;
      }()).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(395)))).cardinality()))));
    };
    static fm36(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).cardinality())), new BigNumber(861)));
    };
    static fm37(p0, globalState) {
      return new _dafny.CodePoint('x'.codePointAt(0));
    };
    static fm38(p0, globalState) {
      return (_dafny.MultiSet.fromElements(true)).Intersect(_dafny.MultiSet.fromElements(false, false, !(false)));
    };
    static fm39(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false, false)).length)),_dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))).cardinality()),_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true), true), _dafny.MultiSet.fromElements(true)), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true)), _dafny.MultiSet.fromElements(false, true, true, true)), _dafny.Seq.of(_dafny.MultiSet.fromElements(false))));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return (_module.D14.create_DC26(_module.D0.create_DC0(!(false)), _dafny.Set.fromElements(false, true, false, false, false))).dtor_cf44;
    };
    static fm42(p0, p1, p2, globalState) {
      return _module.D6.create_DC11(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cxwyhsabm"), _dafny.Seq.UnicodeFromString("gojgfnrgd")));
    };
    static fm43(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_dafny.MultiSet.fromElements(new BigNumber(212)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_dafny.MultiSet.fromElements(new BigNumber(383), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("lu")).length))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), function (_29_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length), new BigNumber(-327))));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-445), new BigNumber(-810), new BigNumber(358)),true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(777)),!(true))).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(381)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), function (_30_i0) {
          return new BigNumber((_dafny.Seq.of(true)).length);
        }))).Elements) {
          let _31_v0 = _compr_5;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(381)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), function (_30_i0) {
            return new BigNumber((_dafny.Seq.of(true)).length);
          }))).contains(_31_v0)) {
            _coll5.push([_31_v0,false]);
          }
        }
        return _coll5;
      }()));
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return _module.D11.create_DC22(!(!(false) || (true)));
    };
    static fm46(p0, globalState) {
      return _module.D11.create_DC21(((true) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (new _dafny.CodePoint('e'.codePointAt(0)))), new BigNumber((_dafny.Seq.UnicodeFromString("pldd")).length), new BigNumber(-615));
    };
    static fm47(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-131))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(391))));
    };
    static fm48(p0, globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(788))).cardinality()));
    };
    static fm49(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_module.D5.create_DC10(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))))).cardinality()), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("afr"),function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(249), new BigNumber(-663))) {
    let _32_v0 = _compr_6;
    if (((new BigNumber(249)).isLessThanOrEqualTo(_32_v0)) && ((_32_v0).isLessThan(new BigNumber(-663)))) {
      _coll6.push([(_32_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ahorcb")).length)),false]);
    }
  }
  return _coll6;
}()), false)).dtor_cf17));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D14.create_DC27(!(true));
      if (_source1.is_DC26) {
        let _33___mcc_h0 = (_source1).cf43;
        let _34___mcc_h1 = (_source1).cf44;
        let _35_cf44 = _34___mcc_h1;
        let _36_cf43 = _33___mcc_h0;
        return _module.D7.create_DC15(new BigNumber(85), new _dafny.CodePoint('d'.codePointAt(0)));
      } else if (_source1.is_DC27) {
        let _37___mcc_h2 = (_source1).cf45;
        let _38_cf45 = _37___mcc_h2;
        return _module.D7.create_DC15(new BigNumber((_dafny.Set.fromElements(_38_cf45, _38_cf45, _38_cf45, _38_cf45, _38_cf45)).length), new _dafny.CodePoint('b'.codePointAt(0)));
      } else if (_source1.is_DC28) {
        return _module.D7.create_DC15(new BigNumber(404), new _dafny.CodePoint('n'.codePointAt(0)));
      } else {
        let _39___mcc_h3 = (_source1).cf42;
        let _40_cf42 = _39___mcc_h3;
        return _module.D7.create_DC15(new BigNumber(717), new _dafny.CodePoint('m'.codePointAt(0)));
      }
    };
    static fm51(p0, p1, p2, globalState) {
      return _module.D14.create_DC28();
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_module.D6.create_DC11((_module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_41_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}))).dtor_cf21), _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("lqncemshr")))).Intersect(_dafny.Set.fromElements(_module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(718)), function (_42_i1) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})), _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("ghyhpuqt"))));
    };
    static fm53(p0, p1, globalState) {
      return _module.D14.create_DC26(_module.D0.create_DC0(true), _dafny.Set.fromElements(false));
    };
    static fm54(p0, p1, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-926)), function (_43_i0) {
        return new BigNumber(-743);
      }), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(845), new BigNumber(377))) {
          let _44_v0 = _compr_7;
          if (((new BigNumber(845)).isLessThanOrEqualTo(_44_v0)) && ((_44_v0).isLessThan(new BigNumber(377)))) {
            _coll7.push([_module.__default.safeModuloInt(_44_v0, new BigNumber(674)),false]);
          }
        }
        return _coll7;
      }()).length))).length), new BigNumber(((_module.D24.create_DC48(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber(-110))).length)))).dtor_cf67).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(893))).cardinality()), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length)))).cardinality()))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), function (_45_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(426),!(!(false)))).length);
      })))).length);
    };
    static fm55(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false)),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true));
    };
    static fm56(p0, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(747)), _dafny.MultiSet.fromElements(new BigNumber(-652)))).Union((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(745), new BigNumber(760)))).Union(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-439)))));
    };
    static fm57(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),true));
    };
    static m0(p0, p1, globalState) {
      let r0 = [];
      let _46_v0;
      _46_v0 = new BigNumber(-220);
      let _47_v1;
      let _nw0 = new _module.C2();
      _nw0.__ctor(_46_v0);
      _47_v1 = _nw0;
      let _48_v2;
      _48_v2 = true;
      _48_v2 = _module.__default.fm5(globalState);
      let _49_v3;
      let _nw1 = new _module.C2();
      _nw1.__ctor(_46_v0);
      _49_v3 = _nw1;
      let _50_v4;
      _50_v4 = _49_v3;
      let _source2 = _50_v4;
      let _51___mcc_h0 = _source2;
      let _52_cf53 = _51___mcc_h0;
      let _53_v5;
      _53_v5 = _dafny.Seq.UnicodeFromString("dxgtgy");
      let _54_v6;
      let _nw2 = Array((new BigNumber(25)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = (_50_v4);
      _nw2[(_dafny.ONE).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(2)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(3)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(4)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(5)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(6)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(7)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(8)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(9)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(10)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(11)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(12)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(13)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(14)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(15)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(16)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(17)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(18)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(19)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(20)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(21)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(22)).toNumber()] = _52_cf53;
      _nw2[(new BigNumber(23)).toNumber()] = _49_v3;
      _nw2[(new BigNumber(24)).toNumber()] = _49_v3;
      _54_v6 = _nw2;
      let _55_v7;
      let _init0 = ((_56_v0) => function (_57_i0) {
        return (_57_i0).minus(_56_v0);
      })(_46_v0);
      let _nw3 = Array((new BigNumber(9)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
        _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _55_v7 = _nw3;
      let _58_v8;
      _58_v8 = _dafny.Map.Empty.slice().updateUnsafe(_55_v7,_54_v6);
      let _59_v9;
      let _nw4 = Array((new BigNumber(24)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = _54_v6;
      _nw4[(_dafny.ONE).toNumber()] = _54_v6;
      _nw4[(new BigNumber(2)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(3)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(4)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(5)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(6)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(7)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(8)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(9)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(10)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(11)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(12)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(13)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(14)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(15)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(16)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(17)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(18)).toNumber()] = (((_58_v8).contains(_55_v7)) ? ((_58_v8).get(_55_v7)) : (_54_v6));
      _nw4[(new BigNumber(19)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(20)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(21)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(22)).toNumber()] = _54_v6;
      _nw4[(new BigNumber(23)).toNumber()] = _54_v6;
      _59_v9 = _nw4;
      let _60_v10;
      _60_v10 = _dafny.Map.Empty.slice().updateUnsafe(_53_v5,_59_v9);
      _60_v10 = (_60_v10).update(_53_v5, _59_v9);
      if (!(_48_v2)) {
        let _61_v11;
        let _nw5 = Array((new BigNumber(16)).toNumber()).fill(false);
        _61_v11 = _nw5;
        let _62_v12;
        _62_v12 = _dafny.Seq.of(_61_v11, _61_v11);
        let _63_v13;
        _63_v13 = _dafny.MultiSet.fromElements(new BigNumber(-971), _46_v0);
        let _64_v14;
        _64_v14 = _dafny.Set.fromElements(!_dafny.areEqual(_62_v12, _62_v12), !((_63_v13).IsDisjointFrom(_63_v13)), (_47_v1).fm8(_53_v5, p0, globalState), p1);
        let _65_v15;
        let _nw6 = Array((new BigNumber(13)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = _50_v4;
        _nw6[(_dafny.ONE).toNumber()] = _50_v4;
        _nw6[(new BigNumber(2)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(3)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(4)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(5)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(6)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(7)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(8)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(9)).toNumber()] = _52_cf53;
        _nw6[(new BigNumber(10)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(11)).toNumber()] = _50_v4;
        _nw6[(new BigNumber(12)).toNumber()] = _50_v4;
        _65_v15 = _nw6;
        let _rhs0 = _46_v0;
        let _rhs1 = _64_v14;
        let _rhs2 = _65_v15;
        _46_v0 = _rhs0;
        _64_v14 = _rhs1;
        _65_v15 = _rhs2;
        let _66_v16;
        _66_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_53_v5).length),!(_48_v2));
        let _67_v17;
        _67_v17 = _dafny.Map.Empty.slice().updateUnsafe(_53_v5,_66_v16);
        let _68_v18;
        _68_v18 = _module.D5.create_DC10(p1, _46_v0, _67_v17, p0);
        let _69_v19;
        _69_v19 = _dafny.Map.Empty.slice().updateUnsafe(true,_68_v18);
        _69_v19 = ((p0) ? (_69_v19) : (_69_v19));
        let _index0 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_61_v11).length));
        (_61_v11)[_index0] = p1;
        let _index1 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_61_v11).length));
        (_61_v11)[_index1] = p0;
        let _70_v20;
        let _nw7 = new _module.C7();
        _nw7.__ctor();
        _70_v20 = _nw7;
        let _71_v21;
        _71_v21 = _module.D20.create_DC38(_70_v20);
        _70_v20 = (_71_v21).dtor_cf57;
        let _72_v22;
        _72_v22 = _module.D11.create_DC22(p1);
        let _73_v23;
        _73_v23 = _dafny.Map.Empty.slice().updateUnsafe(_72_v22,_46_v0);
        let _74_v24;
        _74_v24 = _dafny.Set.fromElements(_73_v23, _73_v23, _73_v23);
        let _index2 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_55_v7).length));
        (_55_v7)[_index2] = new BigNumber((_74_v24).length);
        let _75_v25;
        _75_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-954)), ((_76_v11) => function (_77_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),(_76_v11)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_76_v11).length))]);
        })(_61_v11)),(_49_v3).fm1((_61_v11)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_61_v11).length))], globalState));
        let _78_v26;
        _78_v26 = new _dafny.CodePoint('l'.codePointAt(0));
        let _79_v27;
        _79_v27 = _dafny.Map.Empty.slice().updateUnsafe(_78_v26,p1);
        let _80_v28;
        _80_v28 = _dafny.Seq.of(p0);
        let _81_v29;
        _81_v29 = _dafny.Seq.of(_79_v27, _module.__default.fm49(_80_v28, !(p0), _46_v0, globalState));
        let _index3 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_55_v7).length));
        (_55_v7)[_index3] = (((_75_v25).contains(_81_v29)) ? ((_75_v25).get(_81_v29)) : (_46_v0));
      } else {
        _46_v0 = (_dafny.ZERO).minus(_46_v0);
        let _82_v30;
        let _nw8 = new _module.C3();
        _nw8.__ctor(((_dafny.ZERO).minus(_46_v0)).plus(_46_v0));
        _82_v30 = _nw8;
        let _83_v31;
        let _nw9 = new _module.C4();
        _nw9.__ctor(_46_v0);
        _83_v31 = _nw9;
        let _84_v32;
        _84_v32 = new _dafny.CodePoint('u'.codePointAt(0));
        let _85_v33;
        let _nw10 = Array((_dafny.ONE).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _84_v32;
        _85_v33 = _nw10;
        let _index4 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_85_v33).length));
        (_85_v33)[_index4] = _84_v32;
        let _index5 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_85_v33).length));
        (_85_v33)[_index5] = _84_v32;
        let _86_v34;
        let _nw11 = Array((new BigNumber(16)).toNumber()).fill(false);
        _86_v34 = _nw11;
        let _87_v35;
        _87_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_46_v0);
        let _88_v36;
        _88_v36 = _dafny.MultiSet.fromElements(_48_v2);
        let _89_v37;
        _89_v37 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,_47_v1);
        let _90_v38;
        _90_v38 = _dafny.Map.Empty.slice().updateUnsafe(_46_v0,_46_v0);
        let _91_v39;
        _91_v39 = _dafny.MultiSet.fromElements(_90_v38);
        let _nw12 = Array((new BigNumber(18)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = p0;
        _nw12[(_dafny.ONE).toNumber()] = (_87_v35).equals(_87_v35);
        _nw12[(new BigNumber(2)).toNumber()] = p0;
        _nw12[(new BigNumber(3)).toNumber()] = !((_88_v36).IsDisjointFrom(_88_v36));
        _nw12[(new BigNumber(4)).toNumber()] = p0;
        _nw12[(new BigNumber(5)).toNumber()] = (_89_v37).contains(p1);
        _nw12[(new BigNumber(6)).toNumber()] = (_48_v2) === (_48_v2);
        _nw12[(new BigNumber(7)).toNumber()] = p0;
        _nw12[(new BigNumber(8)).toNumber()] = _48_v2;
        _nw12[(new BigNumber(9)).toNumber()] = ((p1) ? (p0) : (_48_v2));
        _nw12[(new BigNumber(10)).toNumber()] = ((_48_v2) ? (_48_v2) : (p1));
        _nw12[(new BigNumber(11)).toNumber()] = _48_v2;
        _nw12[(new BigNumber(12)).toNumber()] = p1;
        _nw12[(new BigNumber(13)).toNumber()] = _module.__default.fm3(_48_v2, new BigNumber((_dafny.Seq.of(_82_v30)).length), _46_v0, globalState);
        _nw12[(new BigNumber(14)).toNumber()] = p0;
        _nw12[(new BigNumber(15)).toNumber()] = p1;
        _nw12[(new BigNumber(16)).toNumber()] = (_91_v39).IsDisjointFrom(_91_v39);
        _nw12[(new BigNumber(17)).toNumber()] = (_88_v36).IsSubsetOf(_88_v36);
        _86_v34 = _nw12;
      }
      let _92_v40;
      _92_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
      let _93_v41;
      _93_v41 = _92_v40;
      let _94_v42;
      _94_v42 = _dafny.Seq.of(_46_v0);
      let _95_v43;
      _95_v43 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_94_v42, _module.__default.safeIndex(new BigNumber((_53_v5).length), new BigNumber((_94_v42).length)), _46_v0)).length));
      let _96_v44;
      _96_v44 = _dafny.Seq.of((_92_v40).update(!(true), p1), _92_v40);
      let _97_v45;
      _97_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),p0);
      let _98_v46;
      let _nw13 = Array((new BigNumber(25)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _93_v41;
      _nw13[(_dafny.ONE).toNumber()] = _93_v41;
      _nw13[(new BigNumber(2)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(3)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(4)).toNumber()] = _92_v40;
      _nw13[(new BigNumber(5)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(6)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(7)).toNumber()] = _92_v40;
      _nw13[(new BigNumber(8)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(9)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(10)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(11)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(12)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(13)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(14)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(15)).toNumber()] = _module.__default.fm57(_95_v43, globalState);
      _nw13[(new BigNumber(16)).toNumber()] = _module.__default.fm57(_95_v43, globalState);
      _nw13[(new BigNumber(17)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(18)).toNumber()] = (_96_v44)[_module.__default.safeIndex(_46_v0, new BigNumber((_96_v44).length))];
      _nw13[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_97_v45).contains(_dafny.Seq.of(_48_v2, p0))) ? ((_97_v45).get(_dafny.Seq.of(_48_v2, p0))) : (!(!(_48_v2)))),p1);
      _nw13[(new BigNumber(20)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(21)).toNumber()] = _92_v40;
      _nw13[(new BigNumber(22)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(23)).toNumber()] = _93_v41;
      _nw13[(new BigNumber(24)).toNumber()] = _93_v41;
      _98_v46 = _nw13;
      let _index6 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_98_v46).length));
      (_98_v46)[_index6] = _93_v41;
      let _99_v47;
      let _nw14 = new _module.C7();
      _nw14.__ctor();
      _99_v47 = _nw14;
      let _100_v48;
      let _nw15 = new _module.C1();
      _nw15.__ctor(_53_v5);
      _100_v48 = _nw15;
      let _101_v49;
      _101_v49 = _dafny.Seq.of(_100_v48);
      let _102_v50;
      _102_v50 = _dafny.Set.fromElements(p0, p1, !(_dafny.Seq.contains(_101_v49, (_101_v49)[_module.__default.safeIndex((_47_v1).fm1(!(p0), globalState), new BigNumber((_101_v49).length))])), p1);
      let _index7 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_98_v46).length));
      let _rhs3 = _module.__default.fm57(_95_v43, globalState);
      let _rhs4 = _46_v0;
      let _rhs5 = _99_v47;
      let _rhs6 = (_102_v50).Difference((_module.__default.fm41(_94_v42, p0, p1, _48_v2, globalState)).Union(_dafny.Set.fromElements(!(_48_v2))));
      let _lhs0 = _98_v46;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_98_v46).length));
      _lhs0[_lhs1] = _rhs3;
      _46_v0 = _rhs4;
      _99_v47 = _rhs5;
      _102_v50 = _rhs6;
      _46_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt(_46_v0, new BigNumber(-418)), _46_v0));
      let _103_v51;
      let _nw16 = new _module.C3();
      _nw16.__ctor(new BigNumber(735));
      _103_v51 = _nw16;
      let _104_v52;
      _104_v52 = _103_v51;
      let _source3 = _104_v52;
      let _105___mcc_h1 = _source3;
      let _106_cf33 = _105___mcc_h1;
      let _107_v53;
      let _nw17 = new _module.C6();
      _nw17.__ctor();
      _107_v53 = _nw17;
      let _108_v54;
      let _nw18 = Array((new BigNumber(6)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = (_103_v51.f2).plus(_103_v51.f2);
      _nw18[(_dafny.ONE).toNumber()] = _103_v51.f2;
      _nw18[(new BigNumber(2)).toNumber()] = (_46_v0).multipliedBy(_106_cf33.f2);
      _nw18[(new BigNumber(3)).toNumber()] = (new BigNumber(997)).minus(new BigNumber(-856));
      _nw18[(new BigNumber(4)).toNumber()] = new BigNumber(721);
      _nw18[(new BigNumber(5)).toNumber()] = _46_v0;
      _108_v54 = _nw18;
      let _index8 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_108_v54).length));
      (_108_v54)[_index8] = (_103_v51.f2).minus(_103_v51.f2);
      let _index9 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_108_v54).length));
      (_108_v54)[_index9] = _106_cf33.f2;
      let _109_v55;
      let _init1 = function (_110_i2) {
        return true;
      };
      let _nw19 = Array((_dafny.ONE).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw19.length); _i0_1++) {
        _nw19[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _109_v55 = _nw19;
      let _index10 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_109_v55).length));
      (_109_v55)[_index10] = p0;
      let _index11 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_109_v55).length));
      (_109_v55)[_index11] = p0;
      let _index12 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_109_v55).length));
      (_109_v55)[_index12] = p0;
      let _111_v56;
      let _init2 = ((_112_p0) => function (_113_i3) {
        return _112_p0;
      })(p0);
      let _nw20 = Array((new BigNumber(21)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
        _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _111_v56 = _nw20;
      let _pat_let_tv0 = _111_v56;
      let _pat_let_tv1 = _111_v56;
      let _source4 = function (_pat_let0_0) {
        return function (_116_dt__update__tmp_h3) {
          return function (_pat_let3_0) {
            return function (_117_dt__update_hcf35_h1) {
              return _module.D11.create_DC20(_117_dt__update_hcf35_h1);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let0_0);
      }(function (_pat_let1_0) {
        return function (_114_dt__update__tmp_h2) {
          return function (_pat_let2_0) {
            return function (_115_dt__update_hcf35_h0) {
              return _module.D11.create_DC20(_115_dt__update_hcf35_h0);
            }(_pat_let2_0);
          }(_pat_let_tv0);
        }(_pat_let1_0);
      }(_module.D11.create_DC20(_111_v56)));
      if (_source4.is_DC21) {
        let _118___mcc_h2 = (_source4).cf36;
        let _119___mcc_h3 = (_source4).cf37;
        let _120___mcc_h4 = (_source4).cf38;
        let _121_cf38 = _120___mcc_h4;
        let _122_cf37 = _119___mcc_h3;
        let _123_cf36 = _118___mcc_h2;
        _48_v2 = p0;
        let _nw21 = new _module.C3();
        _nw21.__ctor(_122_cf37);
        _103_v51 = _nw21;
        let _124_v57;
        _124_v57 = _dafny.Seq.of(_103_v51.f2);
        let _125_v58;
        _125_v58 = _dafny.Seq.of(_48_v2, _48_v2);
        let _126_v59;
        _126_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_124_v57).length),_125_v58);
        let _127_v60;
        _127_v60 = _dafny.Map.Empty.slice().updateUnsafe(_103_v51.f2,_122_cf37);
        let _128_v61;
        _128_v61 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(_122_cf37)), _103_v51.f2, _46_v0, new BigNumber((_126_v59).length), (_dafny.ZERO).minus((((_127_v60).contains(_103_v51.f2)) ? ((_127_v60).get(_103_v51.f2)) : (_103_v51.f2))));
        let _129_v62;
        _129_v62 = _dafny.Seq.of(_128_v61, (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_122_cf37))).Intersect(_128_v61), _128_v61);
        _129_v62 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(722)), ((_130_cf37, _131_cf38) => function (_132_i4) {
          return _dafny.MultiSet.fromElements(_130_cf37, _131_cf38);
        })(_122_cf37, _121_cf38));
        (_103_v51).f2 = (((_47_v1).fm1(_48_v2, globalState)).minus(_46_v0)).plus(_module.__default.safeModuloInt(_121_cf38, (((_128_v61).contains(_103_v51.f2)) ? ((_128_v61).get(_103_v51.f2)) : (_121_cf38))));
      } else if (_source4.is_DC22) {
        let _133___mcc_h5 = (_source4).cf39;
        let _134_cf39 = _133___mcc_h5;
        if (p0) {
          _48_v2 = !(p0);
          _134_cf39 = !(_module.__default.fm5(globalState));
          _48_v2 = p1;
          let _135_v63;
          let _nw22 = Array((new BigNumber(4)).toNumber()).fill([]);
          _135_v63 = _nw22;
          let _136_v64;
          let _nw23 = Array((new BigNumber(2)).toNumber()).fill([]);
          _136_v64 = _nw23;
          let _index13 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_135_v63).length));
          (_135_v63)[_index13] = _136_v64;
          let _index14 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_135_v63).length));
          let _nw24 = Array((new BigNumber(2)).toNumber()).fill([]);
          (_135_v63)[_index14] = _nw24;
          (_103_v51).f2 = (_dafny.ZERO).minus(_103_v51.f2);
        } else {
          (_103_v51).f2 = ((_103_v51.f2).multipliedBy(_103_v51.f2)).plus(_103_v51.f2);
          _48_v2 = p0;
          (_103_v51).f2 = _46_v0;
          let _index15 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_111_v56).length));
          (_111_v56)[_index15] = _48_v2;
          let _137_v65;
          _137_v65 = _dafny.Seq.of(_134_cf39, ((p0) ? (_48_v2) : (_48_v2)), !(true) || (p1));
          let _138_v66;
          _138_v66 = _dafny.Seq.UnicodeFromString("pekaqfyha");
          let _index16 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_111_v56).length));
          (_111_v56)[_index16] = (_137_v65)[_module.__default.safeIndex((_module.__default.fm54(_134_cf39, _138_v66, globalState)).plus(_46_v0), new BigNumber((_137_v65).length))];
          let _139_v67;
          let _nw25 = new _module.C6();
          _nw25.__ctor();
          _139_v67 = _nw25;
          _139_v67 = _139_v67;
        }
        let _140_v68;
        _140_v68 = _dafny.Seq.UnicodeFromString("bkantehx");
        let _141_v69;
        let _nw26 = new _module.C7();
        _nw26.__ctor();
        _141_v69 = _nw26;
        let _142_v70;
        _142_v70 = _module.D20.create_DC38(_141_v69);
        let _rhs7 = _dafny.Seq.UnicodeFromString("bjadvjecw");
        let _rhs8 = _142_v70;
        let _rhs9 = _46_v0;
        let _lhs2 = _103_v51;
        _140_v68 = _rhs7;
        _142_v70 = _rhs8;
        _lhs2.f2 = _rhs9;
        let _143_v71;
        _143_v71 = new _dafny.CodePoint('b'.codePointAt(0));
        let _144_v72;
        _144_v72 = _dafny.Seq.of(_dafny.Seq.Concat(_140_v68, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("s"), _module.__default.safeIndex(_103_v51.f2, new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)), _143_v71)));
        _144_v72 = _144_v72;
        let _145_v73;
        let _nw27 = Array((new BigNumber(3)).toNumber());
        _145_v73 = _nw27;
        let _146_v74;
        let _nw28 = new _module.C3();
        _nw28.__ctor(_103_v51.f2);
        _146_v74 = _nw28;
        let _147_v75;
        _147_v75 = _dafny.Seq.of(_146_v74, _146_v74);
        let _index17 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_145_v73).length));
        (_145_v73)[_index17] = (_147_v75)[_module.__default.safeIndex(new BigNumber((_140_v68).length), new BigNumber((_147_v75).length))];
        let _148_v76;
        _148_v76 = _dafny.Set.fromElements(_141_v69);
        let _index18 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_145_v73).length));
        let _nw29 = new _module.C3();
        _nw29.__ctor((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_148_v76).length),_46_v0)).update(_46_v0, _103_v51.f2)).length)).multipliedBy(_103_v51.f2));
        (_145_v73)[_index18] = _nw29;
      } else {
        let _149___mcc_h6 = (_source4).cf35;
        let _150_cf35 = _149___mcc_h6;
        let _151_v77;
        _151_v77 = new _dafny.CodePoint('d'.codePointAt(0));
        let _152_v78;
        _152_v78 = _dafny.Seq.of(_151_v77);
        _151_v77 = (_152_v78)[_module.__default.safeIndex(_103_v51.f2, new BigNumber((_152_v78).length))];
        let _153_v79;
        let _nw30 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _153_v79 = _nw30;
        let _154_v80;
        _154_v80 = _dafny.Seq.of(_46_v0);
        let _155_v81;
        _155_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_154_v80)).cardinality()),_48_v2);
        let _index19 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_153_v79).length));
        (_153_v79)[_index19] = _155_v81;
        let _index20 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_153_v79).length));
        (_153_v79)[_index20] = _155_v81;
        let _156_v82;
        _156_v82 = _dafny.MultiSet.fromElements(_48_v2);
        let _157_v83;
        _157_v83 = _dafny.Map.Empty.slice().updateUnsafe(_46_v0,new BigNumber((_156_v82).cardinality()));
        let _158_v84;
        _158_v84 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_103_v51.f2),_46_v0)).Merge(_157_v83),_150_cf35);
        _158_v84 = _158_v84;
        let _159_v85;
        _159_v85 = _dafny.Map.Empty.slice().updateUnsafe(_48_v2,new BigNumber(622));
        let _160_v86;
        _160_v86 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC27(true),(_dafny.ZERO).minus(new BigNumber((_159_v85).length)));
        let _161_v87;
        let _nw31 = new _module.C2();
        _nw31.__ctor(new BigNumber(((_160_v86).Merge(_160_v86)).length));
        _161_v87 = _nw31;
      }
      let _162_v89;
      _162_v89 = function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-646), new BigNumber(248))) {
          let _163_v88 = _compr_8;
          if (((new BigNumber(-646)).isLessThanOrEqualTo(_163_v88)) && ((_163_v88).isLessThan(new BigNumber(248)))) {
            _coll8.add((_163_v88).multipliedBy(new BigNumber((_dafny.Seq.of(p1, p0, p1, p0)).length)));
          }
        }
        return _coll8;
      }();
      let _pat_let_tv2 = p0;
      _48_v2 = function (_source5) {
        let _164___mcc_h7 = _source5;
        let _165_cf48 = _164___mcc_h7;
        return _pat_let_tv2;
      }(_162_v89);
      let _166_v90;
      let _nw32 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _166_v90 = _nw32;
      r0 = _166_v90;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _167_globalState;
      let _nw33 = new _module.GlobalState();
      _nw33.__ctor();
      _167_globalState = _nw33;
      let _168_v0;
      _168_v0 = true;
      let _169_v1;
      let _out0;
      _out0 = _module.__default.m0(_168_v0, (_168_v0) || (true), _167_globalState);
      _169_v1 = _out0;
      _168_v0 = !(_168_v0);
      if (_168_v0) {
        let _170_v2;
        _170_v2 = new _dafny.CodePoint('y'.codePointAt(0));
        _170_v2 = _170_v2;
        let _171_v3;
        _171_v3 = _dafny.Seq.of(_170_v2);
        let _172_v4;
        let _init3 = function (_173_i0) {
          return !(true);
        };
        let _nw34 = Array((new BigNumber(10)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw34.length); _i0_3++) {
          _nw34[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _172_v4 = _nw34;
        let _174_v5;
        let _nw35 = new _module.C5();
        _nw35.__ctor(new BigNumber((_171_v3).length), _172_v4);
        _174_v5 = _nw35;
        let _175_v6;
        _175_v6 = _dafny.Seq.of((_174_v5).f0, new BigNumber((_171_v3).length));
        let _176_v7;
        _176_v7 = _dafny.MultiSet.fromElements((_174_v5).f0);
        let _rhs10 = ((_dafny.MultiSet.FromArray(_dafny.Seq.update(_175_v6, _module.__default.safeIndex((_174_v5).f0, new BigNumber((_175_v6).length)), (_174_v5).f0))).Difference(_176_v7)).IsProperSubsetOf(_176_v7);
        let _rhs11 = _168_v0;
        let _rhs12 = _168_v0;
        let _rhs13 = !_dafny.Seq.contains(_171_v3, _170_v2);
        _168_v0 = _rhs10;
        _168_v0 = _rhs11;
        _168_v0 = _rhs12;
        _168_v0 = _rhs13;
        let _177_v8;
        _177_v8 = _module.D4.create_DC5(_dafny.Set.fromElements(_168_v0));
        let _178_v9;
        _178_v9 = new BigNumber(-538);
        let _179_v10;
        _179_v10 = _dafny.Seq.of(_168_v0, _168_v0, _168_v0);
        let _rhs14 = ((_168_v0) ? (_168_v0) : ((_179_v10)[_module.__default.safeIndex(new BigNumber((_179_v10).length), new BigNumber((_179_v10).length))]));
        let _rhs15 = _177_v8;
        let _rhs16 = (_dafny.ZERO).minus((_178_v9).minus(_178_v9));
        _168_v0 = _rhs14;
        _177_v8 = _rhs15;
        _178_v9 = _rhs16;
        let _180_v11;
        let _out1;
        _out1 = _module.__default.m0(_168_v0, _module.__default.fm3((_174_v5).fm6(_168_v0, _167_globalState), _178_v9, (_174_v5).f0, _167_globalState), _167_globalState);
        _180_v11 = _out1;
      } else {
        let _181_v12;
        _181_v12 = new BigNumber(229);
        let _182_v13;
        _182_v13 = _dafny.Map.Empty.slice().updateUnsafe(_181_v12,new BigNumber(-100));
        let _183_v14;
        _183_v14 = _dafny.MultiSet.fromElements(_182_v13, _182_v13, _182_v13);
        let _184_v15;
        _184_v15 = _dafny.Seq.of(new BigNumber((_183_v14).cardinality()));
        _184_v15 = _184_v15;
        let _185_v16;
        _185_v16 = _dafny.Seq.UnicodeFromString("t");
        let _186_v17;
        _186_v17 = new _dafny.CodePoint('a'.codePointAt(0));
        let _187_v18;
        let _nw36 = new _module.C1();
        _nw36.__ctor(_dafny.Seq.update(_185_v16, _module.__default.safeIndex(new BigNumber((_184_v15).length), new BigNumber((_185_v16).length)), _186_v17));
        _187_v18 = _nw36;
        _168_v0 = (_module.D8.create_DC17(_168_v0, _187_v18, _181_v12)).dtor_cf30;
        let _188_v19;
        _188_v19 = _dafny.Seq.of(_168_v0, _168_v0, _168_v0);
        let _189_v20;
        let _nw37 = Array((new BigNumber(3)).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = _168_v0;
        _nw37[(_dafny.ONE).toNumber()] = (_188_v19)[_module.__default.safeIndex(new BigNumber(((_187_v18).f3).length), new BigNumber((_188_v19).length))];
        _nw37[(new BigNumber(2)).toNumber()] = _168_v0;
        _189_v20 = _nw37;
        let _index21 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_189_v20).length));
        (_189_v20)[_index21] = !(_168_v0);
        let _index22 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_189_v20).length));
        (_189_v20)[_index22] = _168_v0;
        let _190_v21;
        _190_v21 = _module.D6.create_DC11((_187_v18).f3);
        let _191_v22;
        _191_v22 = _dafny.Map.Empty.slice().updateUnsafe(_190_v21,_181_v12);
        _191_v22 = (_191_v22).update(_190_v21, _module.__default.safeModuloInt(_181_v12, (_dafny.ZERO).minus(_181_v12)));
        let _192_v23;
        let _nw38 = new _module.C3();
        _nw38.__ctor(_181_v12);
        _192_v23 = _nw38;
        let _193_v24;
        _193_v24 = _dafny.Seq.of(_192_v23, _192_v23);
        let _194_v25;
        let _nw39 = Array((new BigNumber(25)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = _192_v23;
        _nw39[(_dafny.ONE).toNumber()] = _192_v23;
        _nw39[(new BigNumber(2)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(3)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(4)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(5)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(6)).toNumber()] = (_193_v24)[_module.__default.safeIndex(_181_v12, new BigNumber((_193_v24).length))];
        _nw39[(new BigNumber(7)).toNumber()] = ((_168_v0) ? (_192_v23) : (_192_v23));
        _nw39[(new BigNumber(8)).toNumber()] = ((true) ? (_192_v23) : (_192_v23));
        _nw39[(new BigNumber(9)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(10)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(11)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(12)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(13)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(14)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(15)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(16)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(17)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(18)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(19)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(20)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(21)).toNumber()] = (((_192_v23).fm8(_185_v16, (_189_v20)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_189_v20).length))], _167_globalState)) ? (_192_v23) : (_192_v23));
        _nw39[(new BigNumber(22)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(23)).toNumber()] = _192_v23;
        _nw39[(new BigNumber(24)).toNumber()] = _192_v23;
        _194_v25 = _nw39;
        _194_v25 = _194_v25;
      }
      let _195_v27;
      _195_v27 = function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(691), new BigNumber(586))) {
          let _196_v26 = _compr_9;
          if (((new BigNumber(691)).isLessThanOrEqualTo(_196_v26)) && ((_196_v26).isLessThan(new BigNumber(586)))) {
            _coll9.add((_196_v26).plus(new BigNumber(462)));
          }
        }
        return _coll9;
      }();
      let _pat_let_tv3 = _168_v0;
      if (function (_source6) {
        let _197___mcc_h0 = _source6;
        let _198_cf48 = _197___mcc_h0;
        return _pat_let_tv3;
      }(_195_v27)) {
        let _199_v28;
        _199_v28 = new BigNumber(269);
        let _200_v29;
        _200_v29 = _dafny.Seq.of((_199_v28).multipliedBy(new BigNumber(-921)));
        _200_v29 = _200_v29;
        let _201_v30;
        _201_v30 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,!(false));
        let _202_v31;
        _202_v31 = _dafny.MultiSet.fromElements((((_201_v30).contains(_168_v0)) ? ((_201_v30).get(_168_v0)) : (_168_v0)));
        if ((_202_v31).IsProperSubsetOf((_202_v31).Difference(_202_v31))) {
          let _203_v32;
          let _nw40 = Array((new BigNumber(25)).toNumber()).fill([]);
          _203_v32 = _nw40;
          let _index23 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length));
          (_203_v32)[_index23] = _169_v1;
          let _index24 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length));
          (_203_v32)[_index24] = _169_v1;
          _168_v0 = _168_v0;
          let _204_v33;
          let _nw41 = Array((new BigNumber(9)).toNumber()).fill(false);
          _204_v33 = _nw41;
          let _index25 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_204_v33).length));
          (_204_v33)[_index25] = _module.__default.fm5(_167_globalState);
          let _205_v34;
          _205_v34 = _dafny.Seq.of(_module.__default.fm5(_167_globalState), _168_v0, _168_v0);
          let _206_v35;
          _206_v35 = _dafny.Seq.UnicodeFromString("qgfsf");
          let _index26 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_204_v33).length));
          let _rhs17 = (_module.__default.safeModuloInt(new BigNumber((_205_v34).length), new BigNumber((_206_v35).length))).isEqualTo((_dafny.ZERO).minus(new BigNumber(((_202_v31).Intersect(_202_v31)).cardinality())));
          let _rhs18 = _168_v0;
          let _lhs3 = _204_v33;
          let _lhs4 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_204_v33).length));
          _168_v0 = _rhs17;
          _lhs3[_lhs4] = _rhs18;
          let _207_v36;
          _207_v36 = _dafny.Map.Empty.slice().updateUnsafe(_199_v28,_199_v28);
          let _208_v37;
          _208_v37 = _dafny.Seq.of(_207_v36);
          let _arr0 = (_203_v32)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length))];
          let _index27 = _module.__default.safeIndex(new BigNumber(882), new BigNumber(((_203_v32)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length))]).length));
          _arr0[_index27] = (new BigNumber((_200_v29).length)).multipliedBy(new BigNumber((_208_v37).length));
          let _209_v38;
          let _nw42 = new _module.C6();
          _nw42.__ctor();
          _209_v38 = _nw42;
          let _210_v39;
          _210_v39 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_209_v38,_199_v28));
          let _arr1 = (_203_v32)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length))];
          let _index28 = _module.__default.safeIndex(new BigNumber(882), new BigNumber(((_203_v32)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_203_v32).length))]).length));
          _arr1[_index28] = new BigNumber((_210_v39).length);
          _168_v0 = (_204_v33)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_204_v33).length))];
        } else {
          let _211_v40;
          _211_v40 = _dafny.Seq.UnicodeFromString("cpimjma");
          _211_v40 = _211_v40;
          let _212_v41;
          let _nw43 = new _module.C1();
          _nw43.__ctor(_211_v40);
          _212_v41 = _nw43;
          _211_v40 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_213_i1) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          });
          let _214_v42;
          let _init4 = function (_215_i2) {
            return !(false);
          };
          let _nw44 = Array((new BigNumber(27)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw44.length); _i0_4++) {
            _nw44[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _214_v42 = _nw44;
          let _index29 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_214_v42).length));
          (_214_v42)[_index29] = _168_v0;
          let _index30 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_214_v42).length));
          (_214_v42)[_index30] = (_199_v28).isLessThan(_199_v28);
          let _index31 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_169_v1).length));
          (_169_v1)[_index31] = _199_v28;
          let _216_v43;
          let _nw45 = Array((new BigNumber(17)).toNumber());
          _216_v43 = _nw45;
          let _index32 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_169_v1).length));
          let _rhs19 = _199_v28;
          let _rhs20 = (_199_v28).minus(_module.__default.safeDivisionInt(_199_v28, (_212_v41).fm1(_168_v0, _167_globalState)));
          let _rhs21 = _216_v43;
          let _lhs5 = _169_v1;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_169_v1).length));
          _199_v28 = _rhs19;
          _lhs5[_lhs6] = _rhs20;
          _216_v43 = _rhs21;
        }
        let _217_v44;
        let _out2;
        _out2 = _module.__default.m0((_202_v31).IsDisjointFrom(_202_v31), _168_v0, _167_globalState);
        _217_v44 = _out2;
        let _218_v45;
        let _out3;
        _out3 = _module.__default.m0(_168_v0, _168_v0, _167_globalState);
        _218_v45 = _out3;
        let _index33 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_218_v45).length));
        (_218_v45)[_index33] = new BigNumber(-497);
        let _219_v46;
        let _nw46 = new _module.C6();
        _nw46.__ctor();
        _219_v46 = _nw46;
        let _220_v47;
        _220_v47 = _dafny.Map.Empty.slice().updateUnsafe(_219_v46,_199_v28);
        let _221_v48;
        _221_v48 = _dafny.Map.Empty.slice().updateUnsafe(_199_v28,_169_v1);
        let _index34 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_218_v45).length));
        (_218_v45)[_index34] = _module.__default.safeModuloInt((((_220_v47).contains(_219_v46)) ? ((_220_v47).get(_219_v46)) : (new BigNumber(-686))), new BigNumber((_221_v48).length));
      } else {
        let _222_v49;
        _222_v49 = _dafny.Seq.UnicodeFromString("htengavdr");
        let _223_v50;
        _223_v50 = new BigNumber(-705);
        let _224_v51;
        _224_v51 = new _dafny.CodePoint('q'.codePointAt(0));
        _222_v49 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("wafife"), _module.__default.safeIndex(_223_v50, new BigNumber((_dafny.Seq.UnicodeFromString("wafife")).length)), _224_v51);
        let _225_v52;
        _225_v52 = _dafny.Seq.of(_223_v50);
        let _226_v53;
        _226_v53 = _dafny.Set.fromElements(_168_v0);
        let _227_v54;
        _227_v54 = _dafny.Map.Empty.slice().updateUnsafe((_225_v52)[_module.__default.safeIndex(new BigNumber((_226_v53).length), new BigNumber((_225_v52).length))],_169_v1);
        let _228_v55;
        let _nw47 = Array((new BigNumber(10)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = (((_227_v54).contains(new BigNumber((_dafny.Seq.of(_223_v50, _223_v50)).length))) ? ((_227_v54).get(new BigNumber((_dafny.Seq.of(_223_v50, _223_v50)).length))) : (_169_v1));
        _nw47[(_dafny.ONE).toNumber()] = _169_v1;
        _nw47[(new BigNumber(2)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(3)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(4)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(5)).toNumber()] = ((_168_v0) ? (_169_v1) : (_169_v1));
        _nw47[(new BigNumber(6)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(7)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(8)).toNumber()] = _169_v1;
        _nw47[(new BigNumber(9)).toNumber()] = _169_v1;
        _228_v55 = _nw47;
        let _229_v56;
        _229_v56 = _dafny.Seq.of(_228_v55, ((_168_v0) ? (_228_v55) : (_228_v55)));
        _228_v55 = (_229_v56)[_module.__default.safeIndex(_223_v50, new BigNumber((_229_v56).length))];
        let _230_v57;
        _230_v57 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(464)), ((_231_v51) => function (_232_i3) {
          return _231_v51;
        })(_224_v51)), _222_v49);
        let _233_v58;
        let _nw48 = new _module.C0();
        _nw48.__ctor(!(_dafny.MultiSet.fromElements(_222_v49, _222_v49)).equals(_230_v57), _168_v0);
        _233_v58 = _nw48;
        let _index35 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_169_v1).length));
        (_169_v1)[_index35] = _223_v50;
        let _index36 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_169_v1).length));
        (_169_v1)[_index36] = _223_v50;
        let _234_v59;
        let _nw49 = new _module.C7();
        _nw49.__ctor();
        _234_v59 = _nw49;
      }
      if (_168_v0) {
        let _235_v60;
        _235_v60 = new BigNumber(823);
        let _236_v61;
        _236_v61 = _dafny.Seq.of(_235_v60);
        let _237_v62;
        _237_v62 = _dafny.Seq.of(_dafny.Seq.update(_236_v61, _module.__default.safeIndex(_235_v60, new BigNumber((_236_v61).length)), _235_v60));
        let _238_v63;
        _238_v63 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,new BigNumber((_236_v61).length));
        let _239_v64;
        _239_v64 = _dafny.Seq.UnicodeFromString("tjwbxhkeo");
        let _240_v65;
        _240_v65 = _dafny.Set.fromElements(_235_v60, new BigNumber((_237_v62).length), (((_238_v63).contains(_168_v0)) ? ((_238_v63).get(_168_v0)) : (_235_v60)), _235_v60, ((_168_v0) ? (_235_v60) : (new BigNumber((_239_v64).length))));
        _240_v65 = _240_v65;
        let _241_v66;
        let _init5 = ((_242_v0) => function (_243_i4) {
          return _242_v0;
        })(_168_v0);
        let _nw50 = Array((new BigNumber(17)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw50.length); _i0_5++) {
          _nw50[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _241_v66 = _nw50;
        let _index37 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length));
        (_241_v66)[_index37] = _168_v0;
        let _index38 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_241_v66).length));
        (_241_v66)[_index38] = _168_v0;
        let _index39 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length));
        let _index40 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_241_v66).length));
        let _rhs22 = (_168_v0) && (_168_v0);
        let _rhs23 = _168_v0;
        let _lhs7 = _241_v66;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length));
        let _lhs9 = _241_v66;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_241_v66).length));
        _lhs7[_lhs8] = _rhs22;
        _lhs9[_lhs10] = _rhs23;
        let _244_v67;
        let _out4;
        _out4 = _module.__default.m0(_168_v0, (_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))], _167_globalState);
        _244_v67 = _out4;
        let _245_v68;
        let _nw51 = new _module.C4();
        _nw51.__ctor(_235_v60);
        _245_v68 = _nw51;
        let _246_v69;
        _246_v69 = _dafny.Map.Empty.slice().updateUnsafe(_245_v68,(_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))]);
        let _247_v70;
        _247_v70 = _dafny.Seq.of(_246_v69);
        let _248_v71;
        _248_v71 = _dafny.Map.Empty.slice().updateUnsafe((_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))],_246_v69);
        let _249_v72;
        let _nw52 = Array((new BigNumber(26)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = (_246_v69).Merge(_246_v69);
        _nw52[(_dafny.ONE).toNumber()] = (_246_v69).Merge(_246_v69);
        _nw52[(new BigNumber(2)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(3)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(4)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(5)).toNumber()] = (_246_v69).Merge(_246_v69);
        _nw52[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_245_v68,(_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))])).Merge(_246_v69);
        _nw52[(new BigNumber(7)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(8)).toNumber()] = (_246_v69).Merge(_dafny.Map.Empty.slice().updateUnsafe(_245_v68,(_245_v68).fm8(_239_v64, (_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))], _167_globalState)));
        _nw52[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_245_v68,_168_v0);
        _nw52[(new BigNumber(10)).toNumber()] = (_247_v70)[_module.__default.safeIndex(_235_v60, new BigNumber((_247_v70).length))];
        _nw52[(new BigNumber(11)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(12)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(13)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(14)).toNumber()] = (((_248_v71).contains(_168_v0)) ? ((_248_v71).get(_168_v0)) : (_246_v69));
        _nw52[(new BigNumber(15)).toNumber()] = (_246_v69).Merge(_246_v69);
        _nw52[(new BigNumber(16)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(17)).toNumber()] = (_module.D18.create_DC33(_246_v69)).dtor_cf50;
        _nw52[(new BigNumber(18)).toNumber()] = (_246_v69).Merge((_dafny.Map.Empty.slice().updateUnsafe(_245_v68,_168_v0)).update(_245_v68, _168_v0));
        _nw52[(new BigNumber(19)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(20)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(21)).toNumber()] = (_246_v69).Merge(_246_v69);
        _nw52[(new BigNumber(22)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_245_v68,_168_v0);
        _nw52[(new BigNumber(24)).toNumber()] = _246_v69;
        _nw52[(new BigNumber(25)).toNumber()] = ((_246_v69).update(_245_v68, false)).Merge(_246_v69);
        _249_v72 = _nw52;
        _249_v72 = _249_v72;
        if ((_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))]) {
          let _250_v73;
          let _nw53 = Array((new BigNumber(22)).toNumber());
          _250_v73 = _nw53;
          let _251_v74;
          _251_v74 = _dafny.MultiSet.fromElements(_235_v60, _235_v60);
          let _252_v75;
          _252_v75 = _dafny.MultiSet.fromElements(_251_v74);
          let _253_v76;
          _253_v76 = _dafny.MultiSet.fromElements(new BigNumber((_252_v75).cardinality()), new BigNumber((_239_v64).length));
          let _254_v77;
          let _nw54 = new _module.C4();
          _nw54.__ctor(_235_v60);
          _254_v77 = _nw54;
          let _255_v78;
          _255_v78 = _dafny.MultiSet.fromElements(_254_v77);
          let _256_v79;
          let _nw55 = new _module.C5();
          _nw55.__ctor((((_251_v74).contains((_245_v68).fm1((_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))], _167_globalState))) ? ((_251_v74).get((_245_v68).fm1((_241_v66)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length))], _167_globalState))) : (new BigNumber((_255_v78).cardinality()))), _241_v66);
          _256_v79 = _nw55;
          let _index41 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_250_v73).length));
          (_250_v73)[_index41] = (_256_v79);
          let _257_v80;
          _257_v80 = _dafny.Map.Empty.slice().updateUnsafe(_235_v60,_254_v77);
          let _258_v81;
          _258_v81 = _dafny.Map.Empty.slice().updateUnsafe(_235_v60,_257_v80);
          let _259_v82;
          _259_v82 = _dafny.Set.fromElements((_257_v80).Merge((((_258_v81).contains(_235_v60)) ? ((_258_v81).get(_235_v60)) : (_257_v80))), _dafny.Map.Empty.slice().updateUnsafe(_235_v60,_254_v77), (_257_v80).update(_235_v60, _254_v77), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_235_v60),_254_v77), _dafny.Map.Empty.slice().updateUnsafe(_235_v60,_254_v77));
          let _260_v83;
          _260_v83 = _dafny.Seq.of(_256_v79);
          let _261_v84;
          _261_v84 = _dafny.Map.Empty.slice().updateUnsafe(_236_v61,(_260_v83)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-19)), new BigNumber((_260_v83).length))]);
          let _index42 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_250_v73).length));
          let _rhs24 = (((_261_v84).contains(_236_v61)) ? ((_261_v84).get(_236_v61)) : (_256_v79));
          let _rhs25 = _259_v82;
          let _lhs11 = _250_v73;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_250_v73).length));
          _lhs11[_lhs12] = _rhs24;
          _259_v82 = _rhs25;
          let _262_v85;
          let _nw56 = new _module.C1();
          _nw56.__ctor(_239_v64);
          _262_v85 = _nw56;
          let _263_v86;
          _263_v86 = _dafny.Set.fromElements(_262_v85, _262_v85);
          let _264_v87;
          _264_v87 = _dafny.Set.fromElements((_263_v86).Intersect(_263_v86));
          _264_v87 = (_264_v87).Difference(_264_v87);
          let _265_v88;
          let _nw57 = Array((new BigNumber(11)).toNumber()).fill([]);
          _265_v88 = _nw57;
          let _266_v89;
          let _nw58 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
          _266_v89 = _nw58;
          let _index43 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_265_v88).length));
          (_265_v88)[_index43] = _266_v89;
          let _index44 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_265_v88).length));
          (_265_v88)[_index44] = _266_v89;
          _168_v0 = _168_v0;
          let _267_v90;
          let _nw59 = Array((new BigNumber(18)).toNumber()).fill([]);
          _267_v90 = _nw59;
          let _index45 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_267_v90).length));
          (_267_v90)[_index45] = _241_v66;
          let _268_v91;
          _268_v91 = _dafny.MultiSet.fromElements(_169_v1);
          let _index46 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length));
          let _index47 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_267_v90).length));
          let _rhs26 = ((_268_v91).Union(_268_v91)).IsSubsetOf(_268_v91);
          let _rhs27 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uookla"), (_262_v85).f3);
          let _rhs28 = _235_v60;
          let _rhs29 = _241_v66;
          let _lhs13 = _241_v66;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_241_v66).length));
          let _lhs15 = _267_v90;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_267_v90).length));
          _lhs13[_lhs14] = _rhs26;
          _239_v64 = _rhs27;
          _235_v60 = _rhs28;
          _lhs15[_lhs16] = _rhs29;
        } else {
          let _269_v92;
          _269_v92 = _dafny.MultiSet.fromElements(_235_v60, _235_v60);
          _235_v60 = ((_235_v60).minus(new BigNumber((_269_v92).cardinality()))).multipliedBy(_235_v60);
          let _270_v93;
          _270_v93 = new _dafny.CodePoint('y'.codePointAt(0));
          _270_v93 = _270_v93;
          _270_v93 = _270_v93;
          _235_v60 = _235_v60;
          _235_v60 = _235_v60;
        }
      } else {
        let _271_v94;
        _271_v94 = _dafny.Seq.UnicodeFromString("dlkwbx");
        _271_v94 = _dafny.Seq.Concat(_271_v94, _dafny.Seq.Concat(_271_v94, _271_v94));
        let _272_v95;
        _272_v95 = new BigNumber(785);
        let _273_v96;
        _273_v96 = new _dafny.CodePoint('y'.codePointAt(0));
        let _274_v97;
        _274_v97 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,_dafny.MultiSet.fromElements(_dafny.Seq.update(_271_v94, _module.__default.safeIndex(_272_v95, new BigNumber((_271_v94).length)), _273_v96), _271_v94));
        let _275_v98;
        _275_v98 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("rllwbgim"));
        _274_v97 = (_274_v97).update(true, (_275_v98).Intersect(_275_v98));
        let _276_v99;
        _276_v99 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,_168_v0);
        let _277_v100;
        _277_v100 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,_dafny.Seq.of((((_276_v99).contains(false)) ? ((_276_v99).get(false)) : (!(false)))));
        let _278_v101;
        _278_v101 = _dafny.Seq.of(_168_v0, _168_v0);
        _277_v100 = (_277_v100).update(_168_v0, _dafny.Seq.Concat(_278_v101, _module.__default.fm16(_167_globalState)));
        let _279_v102;
        _279_v102 = _module.D6.create_DC11(_271_v94);
        let _280_v103;
        _280_v103 = _dafny.Map.Empty.slice().updateUnsafe(_279_v102,_272_v95);
        let _281_v104;
        _281_v104 = _dafny.Set.fromElements(_280_v103);
        let _282_v105;
        _282_v105 = _281_v104;
        let _283_v106;
        let _nw60 = Array((new BigNumber(16)).toNumber());
        _283_v106 = _nw60;
        let _284_v107;
        _284_v107 = _dafny.Map.Empty.slice().updateUnsafe(_282_v105,_283_v106);
        _284_v107 = (_284_v107).update(_282_v105, _283_v106);
        let _index48 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_169_v1).length));
        (_169_v1)[_index48] = _272_v95;
        let _index49 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_169_v1).length));
        (_169_v1)[_index49] = _module.__default.safeModuloInt(_272_v95, (_dafny.ZERO).minus(_272_v95));
      }
      _169_v1 = _169_v1;
      if (_168_v0) {
        let _285_v108;
        _285_v108 = _dafny.Seq.UnicodeFromString("roxp");
        let _286_v109;
        _286_v109 = _module.D6.create_DC11(_285_v108);
        let _pat_let_tv4 = _285_v108;
        let _287_v110;
        _287_v110 = _dafny.Set.fromElements(function (_pat_let4_0) {
          return function (_288_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_289_dt__update_hcf21_h0) {
                return _module.D6.create_DC11(_289_dt__update_hcf21_h0);
              }(_pat_let5_0);
            }(_pat_let_tv4);
          }(_pat_let4_0);
        }(_286_v109), _286_v109);
        let _290_v113;
        _290_v113 = new BigNumber(-20);
        let _291_v114;
        _291_v114 = _dafny.Seq.of(_module.__default.fm5(_167_globalState));
        let _292_v115;
        _292_v115 = _dafny.MultiSet.fromElements(_286_v109, _module.D6.create_DC11(_285_v108));
        let _293_v116;
        _293_v116 = _module.D20.create_DC36(_292_v115);
        let _294_v119;
        _294_v119 = _dafny.Map.Empty.slice().updateUnsafe(_286_v109,_168_v0);
        let _295_v121;
        let _nw61 = Array((new BigNumber(14)).toNumber());
        _nw61[(_dafny.ZERO).toNumber()] = _287_v110;
        _nw61[(_dafny.ONE).toNumber()] = function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_287_v110).Elements) {
              let _296_v111 = _compr_11;
              if ((_287_v110).contains(_296_v111)) {
                _coll11.add(_296_v111);
              }
            }
            return _coll11;
          }()).Elements) {
            let _297_v112 = _compr_10;
            if ((function () {
              let _coll12 = new _dafny.Set();
              for (const _compr_12 of (_287_v110).Elements) {
                let _298_v111 = _compr_12;
                if ((_287_v110).contains(_298_v111)) {
                  _coll12.add(_298_v111);
                }
              }
              return _coll12;
            }()).contains(_297_v112)) {
              _coll10.add(_297_v112);
            }
          }
          return _coll10;
        }();
        _nw61[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_286_v109, _286_v109, _286_v109, _286_v109, _286_v109);
        _nw61[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_286_v109, _286_v109);
        _nw61[(new BigNumber(4)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(5)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(6)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(7)).toNumber()] = _module.__default.fm52(_168_v0, _290_v113, new BigNumber((_module.__default.fm16(_167_globalState)).length), (_291_v114)[_module.__default.safeIndex(_290_v113, new BigNumber((_291_v114).length))], _167_globalState);
        _nw61[(new BigNumber(8)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(9)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(10)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(11)).toNumber()] = (function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of ((_293_v116).dtor_cf54).Elements) {
            let _299_v117 = _compr_13;
            if (((_293_v116).dtor_cf54).contains(_299_v117)) {
              _coll13.add(_299_v117);
            }
          }
          return _coll13;
        }()).Union(function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_286_v109,_290_v113)).Keys.Elements) {
            let _300_v118 = _compr_14;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_286_v109,_290_v113)).contains(_300_v118)) {
              _coll14.add(_300_v118);
            }
          }
          return _coll14;
        }());
        _nw61[(new BigNumber(12)).toNumber()] = _287_v110;
        _nw61[(new BigNumber(13)).toNumber()] = function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of (_294_v119).Keys.Elements) {
            let _301_v120 = _compr_15;
            if ((_294_v119).contains(_301_v120)) {
              _coll15.add(_301_v120);
            }
          }
          return _coll15;
        }();
        _295_v121 = _nw61;
        let _302_v122;
        _302_v122 = _dafny.Set.fromElements(!(_168_v0), _168_v0);
        _295_v121 = (((_302_v122).IsDisjointFrom(_302_v122)) ? (_295_v121) : (_295_v121));
        let _rhs30 = _dafny.Seq.Concat(_285_v108, _285_v108);
        let _rhs31 = !_dafny.areEqual(_285_v108, _285_v108);
        let _rhs32 = !(function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(492), new BigNumber(-52))) {
            let _303_v123 = _compr_16;
            if (((new BigNumber(492)).isLessThanOrEqualTo(_303_v123)) && ((_303_v123).isLessThan(new BigNumber(-52)))) {
              _coll16.push([(_303_v123).plus(_290_v113),_290_v113]);
            }
          }
          return _coll16;
        }()).contains(_290_v113);
        let _rhs33 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qgybp"), _dafny.Seq.UnicodeFromString("ubiaagdx")), _dafny.Seq.Concat(_285_v108, _285_v108));
        _285_v108 = _rhs30;
        _168_v0 = _rhs31;
        _168_v0 = _rhs32;
        _285_v108 = _rhs33;
        _168_v0 = _168_v0;
        let _304_v124;
        let _nw62 = Array((new BigNumber(4)).toNumber()).fill(false);
        _304_v124 = _nw62;
        let _305_v125;
        _305_v125 = _dafny.Map.Empty.slice().updateUnsafe(_290_v113,_304_v124);
        let _306_v126;
        let _nw63 = Array((new BigNumber(8)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = _304_v124;
        _nw63[(_dafny.ONE).toNumber()] = _304_v124;
        _nw63[(new BigNumber(2)).toNumber()] = _304_v124;
        _nw63[(new BigNumber(3)).toNumber()] = _304_v124;
        _nw63[(new BigNumber(4)).toNumber()] = _304_v124;
        _nw63[(new BigNumber(5)).toNumber()] = ((_168_v0) ? (_304_v124) : (_304_v124));
        _nw63[(new BigNumber(6)).toNumber()] = _304_v124;
        _nw63[(new BigNumber(7)).toNumber()] = (((_305_v125).contains(_290_v113)) ? ((_305_v125).get(_290_v113)) : (_304_v124));
        _306_v126 = _nw63;
        let _index50 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_306_v126).length));
        (_306_v126)[_index50] = _304_v124;
        let _307_v127;
        _307_v127 = _module.D0.create_DC0(_168_v0);
        let _308_v128;
        _308_v128 = _module.D14.create_DC26(_307_v127, _302_v122);
        let _index51 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_306_v126).length));
        let _rhs34 = _285_v108;
        let _rhs35 = _304_v124;
        let _rhs36 = _module.__default.fm53(_168_v0, _168_v0, _167_globalState);
        let _lhs17 = _306_v126;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_306_v126).length));
        _285_v108 = _rhs34;
        _lhs17[_lhs18] = _rhs35;
        _308_v128 = _rhs36;
        let _index52 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_169_v1).length));
        (_169_v1)[_index52] = _290_v113;
        let _309_v129;
        _309_v129 = _dafny.MultiSet.fromElements(_169_v1, _169_v1);
        let _index53 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_169_v1).length));
        let _rhs37 = new BigNumber((_309_v129).cardinality());
        let _rhs38 = _module.__default.fm54(_168_v0, _dafny.Seq.Concat(_285_v108, _dafny.Seq.UnicodeFromString("ef")), _167_globalState);
        let _rhs39 = _285_v108;
        let _rhs40 = _169_v1;
        let _lhs19 = _169_v1;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_169_v1).length));
        _lhs19[_lhs20] = _rhs37;
        _290_v113 = _rhs38;
        _285_v108 = _rhs39;
        _169_v1 = _rhs40;
      } else {
        _168_v0 = false;
        let _310_v130;
        _310_v130 = new BigNumber(627);
        let _index54 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_169_v1).length));
        (_169_v1)[_index54] = _310_v130;
        let _311_v131;
        _311_v131 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length));
        let _index55 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_169_v1).length));
        (_169_v1)[_index55] = ((_310_v130).minus(new BigNumber((_311_v131).length))).plus(new BigNumber(-306));
        _168_v0 = _168_v0;
        _168_v0 = _168_v0;
        let _index56 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_169_v1).length));
        (_169_v1)[_index56] = _module.__default.safeDivisionInt(_310_v130, _310_v130);
      }
      let _312_v132;
      _312_v132 = new BigNumber(786);
      let _313_v133;
      _313_v133 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,(_dafny.ZERO).minus(_312_v132));
      _313_v133 = (_313_v133).update(_168_v0, _312_v132);
      let _314_v134;
      _314_v134 = _dafny.Seq.of(_312_v132, new BigNumber(390));
      let _315_v135;
      _315_v135 = _dafny.Map.Empty.slice().updateUnsafe(_312_v132,_168_v0);
      let _316_v136;
      _316_v136 = _dafny.Seq.UnicodeFromString("bmigsl");
      let _317_v137;
      let _nw64 = new _module.C2();
      _nw64.__ctor(_module.__default.fm54(_168_v0, _316_v136, _167_globalState));
      _317_v137 = _nw64;
      let _318_v138;
      _318_v138 = _dafny.Seq.of(_317_v137, _317_v137, _317_v137, _317_v137, _317_v137);
      let _hi0 = _module.__default.safeDivisionInt(new BigNumber((_315_v135).length), new BigNumber((_318_v138).length));
      for (let _319_i5 = (_314_v134)[_module.__default.safeIndex(_312_v132, new BigNumber((_314_v134).length))]; _319_i5.isLessThan(_hi0); _319_i5 = _319_i5.plus(_dafny.ONE)) {
        let _320_v139;
        let _nw65 = new _module.C6();
        _nw65.__ctor();
        _320_v139 = _nw65;
        let _321_v140;
        _321_v140 = _dafny.Set.fromElements(_312_v132);
        let _322_v141;
        _322_v141 = new _dafny.CodePoint('s'.codePointAt(0));
        let _323_v142;
        _323_v142 = _module.D0.create_DC0(_168_v0);
        let _324_v143;
        _324_v143 = _dafny.Set.fromElements(_168_v0, _168_v0);
        let _325_v144;
        _325_v144 = _module.D14.create_DC26(_323_v142, _324_v143);
        let _326_v145;
        _326_v145 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,_168_v0);
        let _327_v146;
        let _nw66 = Array((new BigNumber(17)).toNumber());
        _nw66[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(_319_i5));
        _nw66[(_dafny.ONE).toNumber()] = _314_v134;
        _nw66[(new BigNumber(2)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(3)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(4)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(5)).toNumber()] = _module.__default.fm36(_168_v0, _322_v141, _312_v132, _167_globalState);
        _nw66[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_312_v132, new BigNumber(((_325_v144).dtor_cf44).length)), _module.__default.safeIndex(_312_v132, new BigNumber((_dafny.Seq.of(_312_v132, new BigNumber(((_325_v144).dtor_cf44).length))).length)), _319_i5);
        _nw66[(new BigNumber(7)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(8)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(9)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(10)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(11)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(12)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(13)).toNumber()] = _314_v134;
        _nw66[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_314_v134, _module.__default.safeIndex(new BigNumber((_326_v145).length), new BigNumber((_314_v134).length)), _319_i5);
        _nw66[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_314_v134, _module.__default.safeIndex(_312_v132, new BigNumber((_314_v134).length)), _319_i5);
        _nw66[(new BigNumber(16)).toNumber()] = _314_v134;
        _327_v146 = _nw66;
        let _328_v147;
        let _329_v148;
        let _330_v149;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector0 = (_317_v137).m1(new BigNumber(((_321_v140).Union(_321_v140)).length), _327_v146, _168_v0, _dafny.Seq.IsProperPrefixOf(_314_v134, _314_v134), _167_globalState);
        _out5 = _outcollector0[0];
        _out6 = _outcollector0[1];
        _out7 = _outcollector0[2];
        _328_v147 = _out5;
        _329_v148 = _out6;
        _330_v149 = _out7;
        _328_v147 = _module.__default.fm5(_167_globalState);
        let _331_v150;
        let _nw67 = new _module.C4();
        _nw67.__ctor(_319_i5);
        _331_v150 = _nw67;
        let _332_v151;
        _332_v151 = _module.D21.create_DC39(_331_v150);
        let _333_v152;
        _333_v152 = _dafny.Seq.of((_332_v151).dtor_cf58, _331_v150, _331_v150, _331_v150);
        _331_v150 = (_333_v152)[_module.__default.safeIndex(_312_v132, new BigNumber((_333_v152).length))];
      }
      if (((_168_v0) ? (_168_v0) : (_168_v0))) {
        let _334_v153;
        let _335_v154;
        let _out8;
        let _out9;
        let _outcollector1 = (_317_v137).m3(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_314_v134, _dafny.Seq.of(new BigNumber(553))), _module.__default.safeIndex(_312_v132, new BigNumber((_dafny.Seq.Concat(_314_v134, _dafny.Seq.of(new BigNumber(553)))).length)), _312_v132)), _169_v1, _168_v0, !(_168_v0), _167_globalState);
        _out8 = _outcollector1[0];
        _out9 = _outcollector1[1];
        _334_v153 = _out8;
        _335_v154 = _out9;
        let _336_v155;
        let _nw68 = new _module.C4();
        _nw68.__ctor(_334_v153);
        _336_v155 = _nw68;
        let _337_v156;
        let _nw69 = Array((new BigNumber(2)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = _336_v155;
        _nw69[(_dafny.ONE).toNumber()] = _336_v155;
        _337_v156 = _nw69;
        _337_v156 = _337_v156;
        let _338_v157;
        let _nw70 = new _module.C7();
        _nw70.__ctor();
        _338_v157 = _nw70;
        let _339_v158;
        _339_v158 = _module.D20.create_DC38(_338_v157);
        let _340_v159;
        _340_v159 = _dafny.Set.fromElements(((_168_v0) ? (_339_v158) : (_module.D20.create_DC38(_338_v157))));
        let _341_v160;
        _341_v160 = _dafny.MultiSet.fromElements(_168_v0, _168_v0);
        let _pat_let_tv5 = _340_v159;
        let _rhs41 = !(_168_v0) || (!((_341_v160).IsSubsetOf(_341_v160)));
        let _rhs42 = (function (_pat_let6_0) {
          return function (_342_dt__update__tmp_h1) {
            return function (_pat_let7_0) {
              return function (_343_dt__update_hcf62_h0) {
                return _module.D22.create_DC43(_343_dt__update_hcf62_h0);
              }(_pat_let7_0);
            }(_pat_let_tv5);
          }(_pat_let6_0);
        }(_module.D22.create_DC43(_340_v159))).dtor_cf62;
        let _rhs43 = _168_v0;
        _168_v0 = _rhs41;
        _340_v159 = _rhs42;
        _168_v0 = _rhs43;
        _334_v153 = _module.__default.safeDivisionInt(_335_v154, _312_v132);
        let _344_v161;
        _344_v161 = _dafny.Seq.of(_168_v0, !(_168_v0));
        let _345_v162;
        _345_v162 = _module.D22.create_DC44(_module.__default.fm41(_dafny.Seq.Create(_module.__default.abs(new BigNumber(930)), ((_346_v160) => function (_347_i6) {
  return new BigNumber((_346_v160).cardinality());
})(_341_v160)), _168_v0, (_344_v161)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_344_v161).length))], _168_v0, _167_globalState));
        _345_v162 = _345_v162;
      } else {
        let _348_v163;
        let _nw71 = new _module.C2();
        _nw71.__ctor(_312_v132);
        _348_v163 = _nw71;
        _348_v163 = _348_v163;
        let _349_v164;
        _349_v164 = _dafny.Seq.of(_168_v0);
        _313_v133 = (_313_v133).update(((_168_v0) ? (_168_v0) : (true)), _module.__default.safeModuloInt(new BigNumber((_349_v164).length), _312_v132));
        let _350_v165;
        _350_v165 = _dafny.MultiSet.fromElements(_168_v0, _168_v0);
        let _351_v166;
        let _nw72 = Array((new BigNumber(23)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _168_v0;
        _nw72[(_dafny.ONE).toNumber()] = _168_v0;
        _nw72[(new BigNumber(2)).toNumber()] = (_module.D5.create_DC9(_168_v0, _312_v132, _350_v165)).dtor_cf14;
        _nw72[(new BigNumber(3)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(4)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(5)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(6)).toNumber()] = !(_168_v0);
        _nw72[(new BigNumber(7)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(8)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(9)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(10)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(11)).toNumber()] = (_349_v164)[_module.__default.safeIndex(_312_v132, new BigNumber((_349_v164).length))];
        _nw72[(new BigNumber(12)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(13)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(14)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(15)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(16)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(17)).toNumber()] = true;
        _nw72[(new BigNumber(18)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(19)).toNumber()] = true;
        _nw72[(new BigNumber(20)).toNumber()] = !(_168_v0);
        _nw72[(new BigNumber(21)).toNumber()] = _168_v0;
        _nw72[(new BigNumber(22)).toNumber()] = _168_v0;
        _351_v166 = _nw72;
        let _352_v167;
        _352_v167 = _dafny.Map.Empty.slice().updateUnsafe(_168_v0,_168_v0);
        let _353_v168;
        _353_v168 = _dafny.Map.Empty.slice().updateUnsafe(_351_v166,!(((_module.__default.fm26(_312_v132, _168_v0, _168_v0, _module.__default.fm5(_167_globalState), _167_globalState)).update(_168_v0, _168_v0)).equals(_352_v167)));
        _353_v168 = (_353_v168).update(_351_v166, _168_v0);
        if (_168_v0) {
          let _index57 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_351_v166).length));
          (_351_v166)[_index57] = _168_v0;
          let _index58 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_351_v166).length));
          (_351_v166)[_index58] = _168_v0;
          _168_v0 = false;
          let _index59 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_351_v166).length));
          let _rhs44 = !((new BigNumber((_dafny.Seq.UnicodeFromString("bn")).length)).multipliedBy(_312_v132)).isEqualTo(_348_v163.f2);
          let _rhs45 = _348_v163.f2;
          let _rhs46 = false;
          let _lhs21 = _351_v166;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_351_v166).length));
          let _lhs23 = _348_v163;
          _lhs21[_lhs22] = _rhs44;
          _lhs23.f2 = _rhs45;
          _168_v0 = _rhs46;
          let _354_v169;
          _354_v169 = _module.D6.create_DC12(_168_v0, (_351_v166)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_351_v166).length))]);
          _354_v169 = _354_v169;
          let _355_v170;
          let _356_v171;
          let _out10;
          let _out11;
          let _outcollector2 = (_317_v137).m6(_348_v163.f2, _167_globalState);
          _out10 = _outcollector2[0];
          _out11 = _outcollector2[1];
          _355_v170 = _out10;
          _356_v171 = _out11;
        } else {
          let _357_v172;
          _357_v172 = _module.D6.create_DC13(true, _312_v132);
          _168_v0 = (_357_v172).dtor_cf24;
          let _358_v173;
          _358_v173 = _dafny.Map.Empty.slice().updateUnsafe(_314_v134,(_dafny.ZERO).minus(_348_v163.f2));
          let _359_v174;
          _359_v174 = _358_v173;
          _359_v174 = _358_v173;
          let _360_v175;
          _360_v175 = _dafny.Map.Empty.slice().updateUnsafe((_348_v163.f2).multipliedBy(_348_v163.f2),_314_v134);
          _314_v134 = (((_360_v175).contains(_312_v132)) ? ((_360_v175).get(_312_v132)) : (_314_v134));
          let _361_v176;
          let _init6 = ((_362_v134) => function (_363_i7) {
            return _362_v134;
          })(_314_v134);
          let _nw73 = Array((new BigNumber(14)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw73.length); _i0_6++) {
            _nw73[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _361_v176 = _nw73;
          let _364_v177;
          let _365_v178;
          let _366_v179;
          let _out12;
          let _out13;
          let _out14;
          let _outcollector3 = (_348_v163).m1(_348_v163.f2, _361_v176, _168_v0, _168_v0, _167_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _364_v177 = _out12;
          _365_v178 = _out13;
          _366_v179 = _out14;
          let _367_v180;
          let _nw74 = new _module.C1();
          _nw74.__ctor(_dafny.Seq.Concat(_316_v136, _dafny.Seq.UnicodeFromString("fifwva")));
          _367_v180 = _nw74;
          _367_v180 = _367_v180;
        }
        let _368_v181;
        let _nw75 = new _module.C0();
        _nw75.__ctor((_348_v163.f2).isLessThan((_dafny.ZERO).minus(_312_v132)), (_349_v164)[_module.__default.safeIndex(_348_v163.f2, new BigNumber((_349_v164).length))]);
        _368_v181 = _nw75;
      }
      let _369_v182;
      let _370_v183;
      let _out15;
      let _out16;
      let _outcollector4 = (_317_v137).m6(_312_v132, _167_globalState);
      _out15 = _outcollector4[0];
      _out16 = _outcollector4[1];
      _369_v182 = _out15;
      _370_v183 = _out16;
      let _hi1 = _370_v183;
      for (let _371_i8 = _370_v183; _371_i8.isLessThan(_hi1); _371_i8 = _371_i8.plus(_dafny.ONE)) {
        let _372_v184;
        _372_v184 = new _dafny.CodePoint('j'.codePointAt(0));
        let _373_v185;
        _373_v185 = _module.D7.create_DC15(_312_v132, _372_v184);
        let _374_v186;
        let _nw76 = new _module.C6();
        _nw76.__ctor();
        _374_v186 = _nw76;
        let _pat_let_tv6 = _372_v184;
        let _375_v187;
        _375_v187 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let8_0) {
          return function (_376_dt__update__tmp_h2) {
            return function (_pat_let9_0) {
              return function (_377_dt__update_hcf28_h0) {
                return _module.D7.create_DC15((_376_dt__update__tmp_h2).dtor_cf27, _377_dt__update_hcf28_h0);
              }(_pat_let9_0);
            }(_pat_let_tv6);
          }(_pat_let8_0);
        }(_373_v185)).dtor_cf27,_374_v186);
        _375_v187 = (_375_v187).update((_dafny.ZERO).minus(((true) ? (_370_v183) : (_371_i8))), _374_v186);
        let _378_v188;
        _378_v188 = _dafny.Map.Empty.slice().updateUnsafe(_316_v136,_dafny.Map.Empty.slice().updateUnsafe(_371_i8,_369_v182));
        let _379_v189;
        _379_v189 = _dafny.Seq.of(_369_v182, false, _168_v0, false);
        let _380_v190;
        _380_v190 = _module.D5.create_DC10(_168_v0, _370_v183, (_378_v188).Merge(_378_v188), !((_379_v189)[_module.__default.safeIndex((_314_v134)[_module.__default.safeIndex(_370_v183, new BigNumber((_314_v134).length))], new BigNumber((_379_v189).length))]) || (_168_v0));
        let _source7 = _380_v190;
        if (_source7.is_DC9) {
          let _381___mcc_h1 = (_source7).cf14;
          let _382___mcc_h2 = (_source7).cf15;
          let _383___mcc_h3 = (_source7).cf16;
          let _384_cf16 = _383___mcc_h3;
          let _385_cf15 = _382___mcc_h2;
          let _386_cf14 = _381___mcc_h1;
          _168_v0 = false;
          let _387_v191;
          let _out17;
          _out17 = (_317_v137).m2(_312_v132, new BigNumber(597), !(_168_v0) || (!(_369_v182)), _167_globalState);
          _387_v191 = _out17;
          _370_v183 = (((_369_v182) ? (_312_v132) : (_370_v183))).minus((_374_v186).fm1(true, _167_globalState));
          _168_v0 = (_379_v189)[_module.__default.safeIndex(_385_cf15, new BigNumber((_379_v189).length))];
        } else if (_source7.is_DC10) {
          let _388___mcc_h4 = (_source7).cf17;
          let _389___mcc_h5 = (_source7).cf18;
          let _390___mcc_h6 = (_source7).cf19;
          let _391___mcc_h7 = (_source7).cf20;
          let _392_cf20 = _391___mcc_h7;
          let _393_cf19 = _390___mcc_h6;
          let _394_cf18 = _389___mcc_h5;
          let _395_cf17 = _388___mcc_h4;
          let _396_v192;
          let _out18;
          _out18 = (_317_v137).m2(_312_v132, new BigNumber(223), _369_v182, _167_globalState);
          _396_v192 = _out18;
          _315_v135 = (_315_v135).update((_371_i8).minus(_394_cf18), ((_395_cf17) ? (_168_v0) : (_395_cf17)));
          _317_v137 = _317_v137;
          _317_v137 = _317_v137;
        } else {
          let _397___mcc_h8 = (_source7).cf13;
          let _398_cf13 = _397___mcc_h8;
          let _399_v193;
          let _nw77 = new _module.C0();
          _nw77.__ctor(true, _168_v0);
          _399_v193 = _nw77;
          let _index60 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_169_v1).length));
          (_169_v1)[_index60] = _371_i8;
          let _index61 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_169_v1).length));
          (_169_v1)[_index61] = _module.__default.safeModuloInt(_370_v183, new BigNumber(-579));
          _369_v182 = false;
          let _400_v194;
          _400_v194 = _dafny.Set.fromElements(_399_v193.f4, _399_v193.f4, _399_v193.f4, _399_v193.f4);
          _400_v194 = (_400_v194).Intersect(_400_v194);
        }
        let _401_v195;
        let _nw78 = new _module.C4();
        _nw78.__ctor((((_313_v133).contains(false)) ? ((_313_v133).get(false)) : (new BigNumber(988))));
        _401_v195 = _nw78;
        let _nw79 = new _module.C4();
        _nw79.__ctor(_370_v183);
        _401_v195 = _nw79;
        let _402_v196;
        let _nw80 = new _module.C7();
        _nw80.__ctor();
        _402_v196 = _nw80;
        _402_v196 = _402_v196;
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_169_v1).length))) {
        let _403_i9 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_403_i9)) && ((_403_i9).isLessThan(new BigNumber((_169_v1).length))))) {
          (_169_v1)[(_403_i9)] = (_403_i9).plus(_312_v132);
        }
      }
      if ((_module.__default.safeModuloInt(_370_v183, _370_v183)).isLessThanOrEqualTo(_370_v183)) {
        if (_168_v0) {
          let _404_v197;
          let _nw81 = new _module.C6();
          _nw81.__ctor();
          _404_v197 = _nw81;
          let _405_v198;
          _405_v198 = _404_v197;
          let _406_v199;
          let _nw82 = Array((new BigNumber(15)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = _405_v198;
          _nw82[(_dafny.ONE).toNumber()] = _405_v198;
          _nw82[(new BigNumber(2)).toNumber()] = _404_v197;
          _nw82[(new BigNumber(3)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(4)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(5)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(6)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(7)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(8)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(9)).toNumber()] = _404_v197;
          _nw82[(new BigNumber(10)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(11)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(12)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(13)).toNumber()] = _405_v198;
          _nw82[(new BigNumber(14)).toNumber()] = _405_v198;
          _406_v199 = _nw82;
          _406_v199 = _406_v199;
          let _407_v200;
          _407_v200 = _dafny.Set.fromElements(_369_v182);
          _168_v0 = (_407_v200).IsDisjointFrom((_407_v200).Difference(_407_v200));
          let _408_v201;
          _408_v201 = _module.D6.create_DC11(_316_v136);
          _316_v136 = _dafny.Seq.Concat(_316_v136, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bhskj"), (_408_v201).dtor_cf21));
          let _409_v202;
          let _nw83 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
          _409_v202 = _nw83;
          _409_v202 = _409_v202;
          let _410_v203;
          let _nw84 = Array((new BigNumber(23)).toNumber()).fill([]);
          _410_v203 = _nw84;
          let _index62 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_410_v203).length));
          (_410_v203)[_index62] = _169_v1;
          let _index63 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_410_v203).length));
          (_410_v203)[_index63] = ((true) ? (_169_v1) : (_169_v1));
        } else {
          _369_v182 = !(((_317_v137).fm1(false, _167_globalState)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_312_v132, new BigNumber(785))));
          let _411_v204;
          let _nw85 = Array((new BigNumber(29)).toNumber()).fill([]);
          _411_v204 = _nw85;
          let _412_v205;
          let _nw86 = Array((new BigNumber(2)).toNumber()).fill([]);
          _412_v205 = _nw86;
          let _index64 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_411_v204).length));
          (_411_v204)[_index64] = _412_v205;
          let _413_v206;
          let _nw87 = new _module.C7();
          _nw87.__ctor();
          _413_v206 = _nw87;
          let _414_v207;
          _414_v207 = _module.D20.create_DC38(_413_v206);
          let _415_v208;
          _415_v208 = _dafny.Set.fromElements(_414_v207, _414_v207, _414_v207);
          let _416_v209;
          _416_v209 = _module.D22.create_DC43(_415_v208);
          let _417_v210;
          _417_v210 = _dafny.Seq.of(_416_v209);
          let _418_v211;
          let _nw88 = new _module.C4();
          _nw88.__ctor(_312_v132);
          _418_v211 = _nw88;
          let _419_v212;
          _419_v212 = _dafny.Map.Empty.slice().updateUnsafe((_314_v134)[_module.__default.safeIndex(_312_v132, new BigNumber((_314_v134).length))],_418_v211);
          let _420_v213;
          _420_v213 = _dafny.Seq.of(_418_v211, _418_v211, _418_v211);
          let _421_v214;
          let _nw89 = Array((new BigNumber(29)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _418_v211;
          _nw89[(_dafny.ONE).toNumber()] = _418_v211;
          _nw89[(new BigNumber(2)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(3)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(4)).toNumber()] = ((_168_v0) ? (_418_v211) : (_418_v211));
          _nw89[(new BigNumber(5)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(6)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(7)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(8)).toNumber()] = (((_419_v212).contains(_312_v132)) ? ((_419_v212).get(_312_v132)) : (_418_v211));
          _nw89[(new BigNumber(9)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(10)).toNumber()] = (_420_v213)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("orqkbxy")).length), new BigNumber((_420_v213).length))];
          _nw89[(new BigNumber(11)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(12)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(13)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(14)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(15)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(16)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(17)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(18)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(19)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(20)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(21)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(22)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(23)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(24)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(25)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(26)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(27)).toNumber()] = _418_v211;
          _nw89[(new BigNumber(28)).toNumber()] = _418_v211;
          _421_v214 = _nw89;
          let _index65 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_421_v214).length));
          (_421_v214)[_index65] = _418_v211;
          let _422_v215;
          _422_v215 = new _dafny.CodePoint('r'.codePointAt(0));
          let _423_v216;
          _423_v216 = _dafny.MultiSet.fromElements(_369_v182);
          let _424_v217;
          _424_v217 = _dafny.Map.Empty.slice().updateUnsafe(_423_v216,_369_v182);
          let _425_v218;
          _425_v218 = _module.D21.create_DC39(_418_v211);
          let _index66 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_411_v204).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_421_v214).length));
          let _rhs47 = _412_v205;
          let _rhs48 = new BigNumber(((_module.__default.fm55(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), function (_426_i10) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }), _422_v215, _167_globalState)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_423_v216,true)).Merge(_424_v217))).length);
          let _rhs49 = _module.__default.fm54(_168_v0, _dafny.Seq.UnicodeFromString("kmnormpf"), _167_globalState);
          let _rhs50 = _417_v210;
          let _rhs51 = ((_369_v182) ? (_418_v211) : ((_425_v218).dtor_cf58));
          let _lhs24 = _411_v204;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_411_v204).length));
          let _lhs26 = _421_v214;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_421_v214).length));
          _lhs24[_lhs25] = _rhs47;
          _370_v183 = _rhs48;
          _312_v132 = _rhs49;
          _417_v210 = _rhs50;
          _lhs26[_lhs27] = _rhs51;
          _370_v183 = (_418_v211).fm1(_168_v0, _167_globalState);
          let _427_v219;
          _427_v219 = _dafny.Map.Empty.slice().updateUnsafe(_169_v1,_369_v182);
          let _rhs52 = _369_v182;
          let _rhs53 = _369_v182;
          let _rhs54 = ((_427_v219).Merge(_427_v219)).Merge(_427_v219);
          _168_v0 = _rhs52;
          _168_v0 = _rhs53;
          _427_v219 = _rhs54;
          let _428_v220;
          let _init7 = ((_429_v183, _430_v135, _431_v0) => function (_432_i11) {
            return (((_430_v135).contains(_429_v183)) ? ((_430_v135).get(_429_v183)) : (_431_v0));
          })(_370_v183, _315_v135, _168_v0);
          let _nw90 = Array((new BigNumber(12)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw90.length); _i0_7++) {
            _nw90[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _428_v220 = _nw90;
          let _index68 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_428_v220).length));
          (_428_v220)[_index68] = _168_v0;
          let _index69 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_428_v220).length));
          (_428_v220)[_index69] = _369_v182;
        }
        let _433_v221;
        let _nw91 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
        _433_v221 = _nw91;
        let _434_v222;
        let _nw92 = new _module.C4();
        _nw92.__ctor(_370_v183);
        _434_v222 = _nw92;
        let _435_v223;
        _435_v223 = _434_v222;
        let _index70 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_433_v221).length));
        (_433_v221)[_index70] = _dafny.MultiSet.fromElements(_435_v223, _435_v223);
        let _436_v224;
        _436_v224 = _dafny.Set.fromElements(_434_v222.f2, _434_v222.f2, new BigNumber(658), _434_v222.f2, (_370_v183).multipliedBy((_dafny.ZERO).minus(new BigNumber((_module.__default.fm38(_168_v0, _167_globalState)).cardinality()))));
        let _437_v225;
        _437_v225 = _dafny.MultiSet.fromElements(_435_v223);
        let _index71 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_433_v221).length));
        let _rhs55 = new BigNumber((_436_v224).length);
        let _rhs56 = _437_v225;
        let _lhs28 = _433_v221;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_433_v221).length));
        _370_v183 = _rhs55;
        _lhs28[_lhs29] = _rhs56;
        let _438_v226;
        _438_v226 = _module.D21.create_DC41(_369_v182);
        let _439_v227;
        let _nw93 = new _module.C3();
        _nw93.__ctor((_314_v134)[_module.__default.safeIndex(new BigNumber((_module.__default.fm36(_369_v182, new _dafny.CodePoint('i'.codePointAt(0)), _312_v132, _167_globalState)).length), new BigNumber((_314_v134).length))]);
        _439_v227 = _nw93;
        let _440_v228;
        _440_v228 = _dafny.Map.Empty.slice().updateUnsafe((_438_v226).dtor_cf60,_439_v227);
        let _441_v229;
        _441_v229 = _dafny.Seq.of(_439_v227);
        _440_v228 = (_440_v228).update(((_168_v0) ? (_168_v0) : (_168_v0)), (((_440_v228).contains(false)) ? ((_440_v228).get(false)) : ((_441_v229)[_module.__default.safeIndex(_312_v132, new BigNumber((_441_v229).length))])));
        let _442_v230;
        _442_v230 = new _dafny.CodePoint('w'.codePointAt(0));
        let _443_v231;
        _443_v231 = _module.D20.create_DC37(_442_v230, !(_168_v0));
        let _444_v232;
        _444_v232 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_312_v132)),_443_v231);
        let _445_v233;
        _445_v233 = _dafny.MultiSet.fromElements(_316_v136);
        let _pat_let_tv7 = _369_v182;
        _444_v232 = (_444_v232).update((((_445_v233).contains(_316_v136)) ? ((_445_v233).get(_316_v136)) : (_434_v222.f2)), function (_pat_let10_0) {
          return function (_446_dt__update__tmp_h4) {
            return function (_pat_let11_0) {
              return function (_447_dt__update_hcf56_h0) {
                return _module.D20.create_DC37((_446_dt__update__tmp_h4).dtor_cf55, _447_dt__update_hcf56_h0);
              }(_pat_let11_0);
            }(_pat_let_tv7);
          }(_pat_let10_0);
        }(_443_v231));
        let _448_v234;
        let _nw94 = Array((new BigNumber(9)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = _369_v182;
        _nw94[(_dafny.ONE).toNumber()] = _168_v0;
        _nw94[(new BigNumber(2)).toNumber()] = _168_v0;
        _nw94[(new BigNumber(3)).toNumber()] = _168_v0;
        _nw94[(new BigNumber(4)).toNumber()] = _168_v0;
        _nw94[(new BigNumber(5)).toNumber()] = _369_v182;
        _nw94[(new BigNumber(6)).toNumber()] = !(_369_v182);
        _nw94[(new BigNumber(7)).toNumber()] = _369_v182;
        _nw94[(new BigNumber(8)).toNumber()] = _369_v182;
        _448_v234 = _nw94;
        let _449_v235;
        _449_v235 = _dafny.Map.Empty.slice().updateUnsafe(_369_v182,_448_v234);
        let _450_v236;
        let _nw95 = Array((new BigNumber(16)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = _448_v234;
        _nw95[(_dafny.ONE).toNumber()] = _448_v234;
        _nw95[(new BigNumber(2)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(3)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(4)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(5)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(6)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(7)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(8)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(9)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(10)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(11)).toNumber()] = (((_449_v235).contains(true)) ? ((_449_v235).get(true)) : (_448_v234));
        _nw95[(new BigNumber(12)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(13)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(14)).toNumber()] = _448_v234;
        _nw95[(new BigNumber(15)).toNumber()] = ((!(_168_v0)) ? (_448_v234) : (_448_v234));
        _450_v236 = _nw95;
        let _index72 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_450_v236).length));
        (_450_v236)[_index72] = _448_v234;
        let _index73 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_450_v236).length));
        (_450_v236)[_index73] = _448_v234;
      } else {
        let _451_v237;
        let _nw96 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _451_v237 = _nw96;
        let _452_v238;
        let _nw97 = Array((new BigNumber(9)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _451_v237;
        _nw97[(_dafny.ONE).toNumber()] = _451_v237;
        _nw97[(new BigNumber(2)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(3)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(4)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(5)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(6)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(7)).toNumber()] = _451_v237;
        _nw97[(new BigNumber(8)).toNumber()] = _451_v237;
        _452_v238 = _nw97;
        let _index74 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_452_v238).length));
        (_452_v238)[_index74] = _451_v237;
        let _453_v239;
        _453_v239 = _dafny.Seq.of(_451_v237);
        let _index75 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_452_v238).length));
        (_452_v238)[_index75] = (_453_v239)[_module.__default.safeIndex(new BigNumber((_316_v136).length), new BigNumber((_453_v239).length))];
        _370_v183 = (_dafny.ZERO).minus(((new BigNumber((_314_v134).length)).multipliedBy(_370_v183)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm56(_168_v0, _167_globalState)).cardinality()), _370_v183, new BigNumber((_316_v136).length), _312_v132)).cardinality())));
        let _454_v240;
        _454_v240 = new _dafny.CodePoint('a'.codePointAt(0));
        _454_v240 = _454_v240;
        _454_v240 = _454_v240;
        let _index76 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_169_v1).length));
        (_169_v1)[_index76] = new BigNumber(390);
        let _455_v241;
        _455_v241 = _dafny.MultiSet.fromElements(_454_v240, _454_v240);
        let _index77 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_169_v1).length));
        let _rhs57 = (_dafny.ZERO).minus(_370_v183);
        let _rhs58 = !(_455_v241).contains(_454_v240);
        let _rhs59 = (_dafny.ZERO).minus(_370_v183);
        let _lhs30 = _169_v1;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_169_v1).length));
        _312_v132 = _rhs57;
        _168_v0 = _rhs58;
        _lhs30[_lhs31] = _rhs59;
      }
      _312_v132 = _370_v183;
      let _456_v242;
      let _nw98 = Array((new BigNumber(28)).toNumber()).fill(false);
      _456_v242 = _nw98;
      let _index78 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_456_v242).length));
      (_456_v242)[_index78] = _module.__default.fm3(_168_v0, _312_v132, _312_v132, _167_globalState);
      let _index79 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_456_v242).length));
      (_456_v242)[_index79] = !_dafny.areEqual(_316_v136, _316_v136);
      process.stdout.write(_dafny.toString(_168_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_195_v27)).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_312_v132));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_313_v133).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(786)).updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_314_v134, _dafny.Seq.of(new BigNumber(786), new BigNumber(390)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v135).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(786),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_316_v136).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_318_v138).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_369_v182));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_370_v183));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_456_v242)[new BigNumber(4)]));
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
    get dtor_cf4() { return this.cf4; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC2(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
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
    static create_DC3(cf6) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf7) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7;
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf9, cf10, cf11) {
      let $dt = new D4(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D4(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D4(2);
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
        return "D4.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC6(_dafny.ZERO, [], _dafny.ZERO);
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
    static create_DC9(cf14, cf15, cf16) {
      let $dt = new D5(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC10(cf17, cf18, cf19, cf20) {
      let $dt = new D5(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D5(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
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
        return "D5.DC9" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC9(false, _dafny.ZERO, _dafny.MultiSet.Empty);
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
    static create_DC12(cf22, cf23) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC13(cf24, cf25) {
      let $dt = new D6(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC11(cf21) {
      let $dt = new D6(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC11" + "(" + this.cf21.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf22 === other.cf22 && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC12(false, false);
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
    static create_DC15(cf27, cf28) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC14(cf26) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC17(cf30, cf31, cf32) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D8(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf29 === other.cf29;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC17(false, null, _dafny.ZERO);
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
    static create_DC18(cf33) {
      let $dt = new D9(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33;
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
    static create_DC19(cf34) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf36, cf37, cf38) {
      let $dt = new D11(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC22(cf39) {
      let $dt = new D11(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC20(cf35) {
      let $dt = new D11(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC21" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC22" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC20" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf35 === other.cf35;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC21(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC23(cf40) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC23" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf41) {
      let $dt = new D13(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC24" + "(" + _dafny.toString(this.cf41) + ")";
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
      return _dafny.Set.Empty;
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
    static create_DC26(cf43, cf44) {
      let $dt = new D14(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC27(cf45) {
      let $dt = new D14(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC28() {
      let $dt = new D14(2);
      return $dt;
    }
    static create_DC25(cf42) {
      let $dt = new D14(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC26" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC27" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC28";
      } else if (this.$tag === 3) {
        return "D14.DC25" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf45 === other.cf45;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf42 === other.cf42;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC26(_module.D0.Default(), _dafny.Set.Empty);
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
    static create_DC30(cf47) {
      let $dt = new D15(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC29(cf46) {
      let $dt = new D15(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC30" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC29" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC30(false);
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
    static create_DC31(cf48) {
      let $dt = new D16(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC31" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48);
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
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf49) {
      let $dt = new D17(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC32" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
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
    static create_DC34(cf51, cf52) {
      let $dt = new D18(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC33(cf50) {
      let $dt = new D18(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC34" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC33" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf51 === other.cf51 && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC34(false, false);
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
    static create_DC35(cf53) {
      let $dt = new D19(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC35" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53;
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf55, cf56) {
      let $dt = new D20(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC38(cf57) {
      let $dt = new D20(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC36(cf54) {
      let $dt = new D20(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC37" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC38" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC36" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf57 === other.cf57;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC37(new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC40(cf59) {
      let $dt = new D21(0);
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC41(cf60) {
      let $dt = new D21(1);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC39(cf58) {
      let $dt = new D21(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC42(cf61) {
      let $dt = new D21(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get is_DC42() { return this.$tag === 3; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC40" + "(" + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC41" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC39" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC42" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf60 === other.cf60;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf58 === other.cf58;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC40(null);
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
    static create_DC44(cf63) {
      let $dt = new D22(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC45(cf64) {
      let $dt = new D22(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC43(cf62) {
      let $dt = new D22(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC46(cf65) {
      let $dt = new D22(3);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get is_DC46() { return this.$tag === 3; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC44" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC45" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC43" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC46" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC44(_dafny.Set.Empty);
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
    static create_DC47(cf66) {
      let $dt = new D23(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC47" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf66, other.cf66);
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
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49(cf68, cf69, cf70) {
      let $dt = new D24(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC48(cf67) {
      let $dt = new D24(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC50(cf71) {
      let $dt = new D24(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC49" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC48" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC50" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC49(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
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
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f4 = false;
      this.f5 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f4, f5) {
      let _this = this;
      (_this).f4 = f4;
      (_this).f5 = f5;
      return;
    }
    fm14(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f4;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f3 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f3) {
      let _this = this;
      (_this)._f3 = f3;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_457_i0) {
        return new BigNumber(((_this).f3).length);
      }),new BigNumber(((_this).f3).length))).Merge(((false) ? (function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), function (_458_i1) {
          return new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality());
        }),new BigNumber(352))).Keys.Elements) {
          let _459_v0 = _compr_17;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), function (_458_i1) {
            return new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality());
          }),new BigNumber(352))).contains(_459_v0)) {
            _coll17.push([_459_v0,new BigNumber(894)]);
          }
        }
        return _coll17;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(266)),new BigNumber(-986)))));
    };
    fm1(p0, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true))).length)).plus(new BigNumber((((!(false)) ? (_dafny.Seq.of(false, false)) : (_dafny.Seq.of(!(!(false)), false)))).length));
    };
    fm11(globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(true)).IsDisjointFrom((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(false)));
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-906)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-299)), new BigNumber((_dafny.Seq.UnicodeFromString("ob")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), function (_460_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)));
    };
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _461_v0;
      let _nw99 = Array((new BigNumber(29)).toNumber()).fill([]);
      _461_v0 = _nw99;
      let _462_v1;
      let _nw100 = Array((new BigNumber(19)).toNumber()).fill(false);
      _462_v1 = _nw100;
      let _index80 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_461_v0).length));
      (_461_v0)[_index80] = _462_v1;
      let _index81 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_461_v0).length));
      (_461_v0)[_index81] = _462_v1;
      let _463_v2;
      _463_v2 = new _dafny.CodePoint('t'.codePointAt(0));
      let _464_v3;
      _464_v3 = new BigNumber(469);
      let _465_v4;
      _465_v4 = _module.D0.create_DC1(p3, p3, _463_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f3,_464_v3)).length));
      let _pat_let_tv8 = _464_v3;
      let _pat_let_tv9 = _464_v3;
      r0 = function (_source8) {
        if (_source8.is_DC1) {
          let _466___mcc_h0 = (_source8).cf1;
          let _467___mcc_h1 = (_source8).cf2;
          let _468___mcc_h2 = (_source8).cf3;
          let _469___mcc_h3 = (_source8).cf4;
          let _470_cf4 = _469___mcc_h3;
          let _471_cf3 = _468___mcc_h2;
          let _472_cf2 = _467___mcc_h1;
          let _473_cf1 = _466___mcc_h0;
          return _pat_let_tv8;
        } else {
          let _474___mcc_h4 = (_source8).cf0;
          let _475_cf0 = _474___mcc_h4;
          return _pat_let_tv9;
        }
      }(_465_v4);
      let _476_v5;
      _476_v5 = false;
      let _477_v6;
      let _init8 = ((_478_p3, _479_v5) => function (_480_i0) {
        return (_dafny.MultiSet.fromElements(_478_p3, _479_v5)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_478_p3)));
      })(p3, _476_v5);
      let _nw101 = Array((new BigNumber(27)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw101.length); _i0_8++) {
        _nw101[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _477_v6 = _nw101;
      let _481_v7;
      _481_v7 = _dafny.MultiSet.fromElements(p2, _476_v5);
      let _index82 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_477_v6).length));
      (_477_v6)[_index82] = _481_v7;
      let _482_v8;
      _482_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm13(globalState));
      let _483_v9;
      _483_v9 = _dafny.Map.Empty.slice().updateUnsafe(_464_v3,p3);
      let _pat_let_tv10 = p3;
      let _index83 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_477_v6).length));
      let _rhs60 = _463_v2;
      let _rhs61 = function (_source9) {
        if (_source9.is_DC1) {
          let _484___mcc_h5 = (_source9).cf1;
          let _485___mcc_h6 = (_source9).cf2;
          let _486___mcc_h7 = (_source9).cf3;
          let _487___mcc_h8 = (_source9).cf4;
          let _488_cf4 = _487___mcc_h8;
          let _489_cf3 = _486___mcc_h7;
          let _490_cf2 = _485___mcc_h6;
          let _491_cf1 = _484___mcc_h5;
          return _pat_let_tv10;
        } else {
          let _492___mcc_h9 = (_source9).cf0;
          let _493_cf0 = _492___mcc_h9;
          return _493_cf0;
        }
      }(_module.D0.create_DC1(p3, _476_v5, _463_v2, _464_v3));
      let _rhs62 = (_481_v7).Union(_481_v7);
      let _rhs63 = (new BigNumber((_dafny.Seq.of((((_482_v8).contains(_476_v5)) ? ((_482_v8).get(_476_v5)) : (_483_v9)))).length)).multipliedBy((_dafny.ZERO).minus(_464_v3));
      let _lhs32 = _477_v6;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_477_v6).length));
      _463_v2 = _rhs60;
      _476_v5 = _rhs61;
      _lhs32[_lhs33] = _rhs62;
      r1 = _rhs63;
      let _494_v10;
      _494_v10 = _dafny.Seq.of(new BigNumber(18), _464_v3, _module.__default.safeDivisionInt(_464_v3, _464_v3));
      let _495_v11;
      _495_v11 = _dafny.Seq.of(_476_v5, true);
      _494_v10 = _dafny.Seq.Concat(_dafny.Seq.update((_this).fm12(p3, _476_v5, (_this).fm1((_495_v11)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,_464_v3)).length), new BigNumber((_495_v11).length))], globalState), _476_v5, globalState), _module.__default.safeIndex(_464_v3, new BigNumber(((_this).fm12(p3, _476_v5, (_this).fm1((_495_v11)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,_464_v3)).length), new BigNumber((_495_v11).length))], globalState), _476_v5, globalState)).length)), _464_v3), (_this).fm12((_this).fm11(globalState), p3, _464_v3, _476_v5, globalState));
      let _496_v12;
      let _nw102 = new _module.C0();
      _nw102.__ctor((_464_v3).isLessThanOrEqualTo(_464_v3), p2);
      _496_v12 = _nw102;
      if (p2) {
        let _497_v13;
        _497_v13 = _dafny.Map.Empty.slice().updateUnsafe(_476_v5,true);
        let _source10 = (_497_v13).Merge(_497_v13);
        let _498___mcc_h10 = _source10;
        let _499_cf5 = _498___mcc_h10;
        let _500_v14;
        _500_v14 = _dafny.Map.Empty.slice().updateUnsafe(_464_v3,_464_v3);
        _464_v3 = _module.__default.safeModuloInt(_464_v3, (((_500_v14).contains(_464_v3)) ? ((_500_v14).get(_464_v3)) : (new BigNumber(583))));
        let _501_v15;
        _501_v15 = _dafny.Map.Empty.slice().updateUnsafe(_463_v2,_dafny.Seq.UnicodeFromString("rmd"));
        _501_v15 = (_501_v15).update(_463_v2, _dafny.Seq.Concat((_this).f3, (_this).f3));
        let _502_v16;
        _502_v16 = _dafny.Set.fromElements((_461_v0)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_461_v0).length))]);
        let _503_v17;
        _503_v17 = _dafny.Map.Empty.slice().updateUnsafe(_496_v12.f5,new BigNumber((_502_v16).length));
        let _504_v18;
        _504_v18 = (_497_v13).update(_496_v12.f5, p2);
        let _505_v19;
        _505_v19 = _dafny.Seq.of(_499_cf5, _504_v18);
        let _index84 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((p1).length));
        (p1)[_index84] = (new BigNumber((_503_v17).length)).plus(new BigNumber((_505_v19).length));
        let _index85 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((p1).length));
        (p1)[_index85] = (_464_v3).minus((_494_v10)[_module.__default.safeIndex(_464_v3, new BigNumber((_494_v10).length))]);
        r0 = _464_v3;
        let _506_v20;
        let _nw103 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _506_v20 = _nw103;
        let _507_v21;
        _507_v21 = _dafny.Seq.of(p1, p1, p1);
        let _508_v22;
        let _nw104 = Array((new BigNumber(14)).toNumber());
        _nw104[(_dafny.ZERO).toNumber()] = _506_v20;
        _nw104[(_dafny.ONE).toNumber()] = _506_v20;
        _nw104[(new BigNumber(2)).toNumber()] = p1;
        _nw104[(new BigNumber(3)).toNumber()] = (_507_v21)[_module.__default.safeIndex(_464_v3, new BigNumber((_507_v21).length))];
        _nw104[(new BigNumber(4)).toNumber()] = p1;
        _nw104[(new BigNumber(5)).toNumber()] = p1;
        _nw104[(new BigNumber(6)).toNumber()] = _506_v20;
        _nw104[(new BigNumber(7)).toNumber()] = p1;
        _nw104[(new BigNumber(8)).toNumber()] = p1;
        _nw104[(new BigNumber(9)).toNumber()] = _506_v20;
        _nw104[(new BigNumber(10)).toNumber()] = _506_v20;
        _nw104[(new BigNumber(11)).toNumber()] = _506_v20;
        _nw104[(new BigNumber(12)).toNumber()] = _506_v20;
        _nw104[(new BigNumber(13)).toNumber()] = _506_v20;
        _508_v22 = _nw104;
        let _509_v23;
        _509_v23 = _508_v22;
        let _rhs64 = _461_v0;
        let _rhs65 = _506_v20;
        let _rhs66 = (_509_v23);
        let _rhs67 = _dafny.Seq.Concat(_494_v10, _494_v10);
        _461_v0 = _rhs64;
        _506_v20 = _rhs65;
        _508_v22 = _rhs66;
        _494_v10 = _rhs67;
        let _510_v24;
        _510_v24 = _dafny.Map.Empty.slice().updateUnsafe(_464_v3,(_461_v0)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_461_v0).length))]);
        let _index86 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_461_v0).length));
        (_461_v0)[_index86] = (((_510_v24).contains(_464_v3)) ? ((_510_v24).get(_464_v3)) : (_462_v1));
        _476_v5 = !(((_464_v3).multipliedBy(new BigNumber(((_this).f3).length))).isLessThan(_module.__default.safeDivisionInt(_464_v3, _464_v3)));
        _476_v5 = !(!((_465_v4).dtor_cf2));
      } else {
        (_496_v12).f4 = !(_476_v5);
        (_496_v12).f5 = !(_496_v12.f5);
        r0 = _464_v3;
        (_496_v12).f4 = !(_496_v12.f4);
        (_496_v12).f4 = p3;
      }
      r0 = _module.__default.safeModuloInt(_464_v3, (_464_v3).plus(_464_v3));
      r1 = (_464_v3).multipliedBy(new BigNumber(((_this).f3).length));
      return [r0, r1];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _511_v0;
      _511_v0 = _dafny.MultiSet.fromElements(p2);
      if (!((_511_v0).IsProperSubsetOf(((_511_v0).update(p3, _module.__default.abs(p0))).Difference(_511_v0)))) {
        let _512_v1;
        _512_v1 = _dafny.Seq.of(p0);
        let _513_v2;
        _513_v2 = new _dafny.CodePoint('a'.codePointAt(0));
        let _514_v3;
        _514_v3 = _dafny.Map.Empty.slice().updateUnsafe(_513_v2,new _dafny.CodePoint('q'.codePointAt(0)));
        let _515_v4;
        _515_v4 = _dafny.MultiSet.fromElements(new BigNumber((_514_v3).length));
        r2 = (_512_v1)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((function () {
          let _coll18 = new _dafny.Set();
          for (const _compr_18 of (_515_v4).Elements) {
            let _516_v5 = _compr_18;
            if ((_515_v4).contains(_516_v5)) {
              _coll18.add((_516_v5).plus(new BigNumber(315)));
            }
          }
          return _coll18;
        }()).length), p0), new BigNumber((_512_v1).length))];
        let _517_v6;
        _517_v6 = _dafny.Set.fromElements(p2);
        let _518_v7;
        _518_v7 = _module.D4.create_DC5(_517_v6);
        let _519_v8;
        let _nw105 = new _module.C0();
        _nw105.__ctor(p3, ((_518_v7).dtor_cf8).equals(_517_v6));
        _519_v8 = _nw105;
        r2 = ((_519_v8.f4) ? (p0) : ((p0).multipliedBy(p0)));
        let _520_v9;
        let _nw106 = new _module.C0();
        _nw106.__ctor(_519_v8.f5, _519_v8.f4);
        _520_v9 = _nw106;
        let _521_v10;
        _521_v10 = _dafny.Seq.UnicodeFromString("gragg");
        _521_v10 = _module.__default.fm15(globalState);
      } else {
        let _522_v11;
        let _nw107 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _522_v11 = _nw107;
        let _523_v12;
        _523_v12 = _dafny.Seq.of(_522_v11, _522_v11);
        let _524_v13;
        _524_v13 = _module.D4.create_DC6(_module.__default.safeDivisionInt(p0, p0), (_523_v12)[_module.__default.safeIndex(new BigNumber(((_this).f3).length), new BigNumber((_523_v12).length))], (p0).minus(p0));
        let _source11 = _524_v13;
        if (_source11.is_DC6) {
          let _525___mcc_h0 = (_source11).cf9;
          let _526___mcc_h1 = (_source11).cf10;
          let _527___mcc_h2 = (_source11).cf11;
          let _528_cf11 = _527___mcc_h2;
          let _529_cf10 = _526___mcc_h1;
          let _530_cf9 = _525___mcc_h0;
          let _531_v14;
          _531_v14 = new _dafny.CodePoint('q'.codePointAt(0));
          let _532_v15;
          _532_v15 = _module.D0.create_DC1(p3, p2, _531_v14, new BigNumber(971));
          let _index87 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_522_v11).length));
          (_522_v11)[_index87] = ((p2) ? ((_532_v15).dtor_cf4) : (_530_cf9));
          let _533_v16;
          _533_v16 = _dafny.Map.Empty.slice().updateUnsafe(_528_cf11,new BigNumber(((_this).f3).length));
          let _534_v17;
          _534_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _535_v18;
          _535_v18 = _dafny.Seq.of(p2, p2, (_this).fm11(globalState), p3);
          let _index88 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_522_v11).length));
          (_522_v11)[_index88] = ((!(p3) || (!(false))) ? ((((_533_v16).contains((_dafny.ZERO).minus(new BigNumber((_534_v17).length)))) ? ((_533_v16).get((_dafny.ZERO).minus(new BigNumber((_534_v17).length)))) : (_530_cf9))) : (new BigNumber((_dafny.Seq.Concat(_535_v18, _535_v18)).length)));
          let _536_v19;
          let _nw108 = Array((new BigNumber(13)).toNumber());
          _nw108[(_dafny.ZERO).toNumber()] = p2;
          _nw108[(_dafny.ONE).toNumber()] = p2;
          _nw108[(new BigNumber(2)).toNumber()] = p3;
          _nw108[(new BigNumber(3)).toNumber()] = p3;
          _nw108[(new BigNumber(4)).toNumber()] = p2;
          _nw108[(new BigNumber(5)).toNumber()] = p3;
          _nw108[(new BigNumber(6)).toNumber()] = p2;
          _nw108[(new BigNumber(7)).toNumber()] = p2;
          _nw108[(new BigNumber(8)).toNumber()] = (_535_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber(254), _530_cf9, new BigNumber((_dafny.Seq.of(new BigNumber((_511_v0).cardinality()))).length))).length), new BigNumber((_535_v18).length))];
          _nw108[(new BigNumber(9)).toNumber()] = p2;
          _nw108[(new BigNumber(10)).toNumber()] = p2;
          _nw108[(new BigNumber(11)).toNumber()] = p3;
          _nw108[(new BigNumber(12)).toNumber()] = p3;
          _536_v19 = _nw108;
          let _index89 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_536_v19).length));
          (_536_v19)[_index89] = (_this).fm11(globalState);
          let _index90 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_536_v19).length));
          (_536_v19)[_index90] = false;
          let _537_v20;
          _537_v20 = _dafny.Seq.of((((_533_v16).contains(p0)) ? ((_533_v16).get(p0)) : ((_522_v11)[_module.__default.safeIndex(new BigNumber(620), new BigNumber((_522_v11).length))])));
          _537_v20 = _537_v20;
          let _index91 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_536_v19).length));
          (_536_v19)[_index91] = (_535_v18)[_module.__default.safeIndex((_522_v11)[_module.__default.safeIndex(new BigNumber(620), new BigNumber((_522_v11).length))], new BigNumber((_535_v18).length))];
        } else if (_source11.is_DC5) {
          let _538___mcc_h3 = (_source11).cf8;
          let _539_cf8 = _538___mcc_h3;
          let _540_v21;
          let _nw109 = new _module.C0();
          _nw109.__ctor(p2, p2);
          _540_v21 = _nw109;
          let _541_v22;
          _541_v22 = _dafny.Seq.UnicodeFromString("n");
          let _542_v23;
          let _init9 = function (_543_i0) {
            return true;
          };
          let _nw110 = Array((new BigNumber(5)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw110.length); _i0_9++) {
            _nw110[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _542_v23 = _nw110;
          let _544_v24;
          let _nw111 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Set.Empty);
          _544_v24 = _nw111;
          let _545_v25;
          _545_v25 = new _dafny.CodePoint('n'.codePointAt(0));
          let _546_v26;
          _546_v26 = _module.D0.create_DC1(_540_v21.f5, p3, _545_v25, p0);
          let _547_v27;
          _547_v27 = _dafny.Set.fromElements(_546_v26, _module.D0.create_DC1(_540_v21.f5, p3, _545_v25, p0), _546_v26, _546_v26);
          let _index92 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_544_v24).length));
          (_544_v24)[_index92] = _547_v27;
          let _index93 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_544_v24).length));
          let _rhs68 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(979)), ((_548_v25) => function (_549_i1) {
            return _548_v25;
          })(_545_v25)), (_this).f3);
          let _rhs69 = _542_v23;
          let _rhs70 = !(((new BigNumber((_541_v22).length)).multipliedBy(p0)).isLessThan(p0));
          let _rhs71 = _547_v27;
          let _rhs72 = _540_v21.f5;
          let _lhs34 = _540_v21;
          let _lhs35 = _544_v24;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_544_v24).length));
          _541_v22 = _rhs68;
          _542_v23 = _rhs69;
          _lhs34.f4 = _rhs70;
          _lhs35[_lhs36] = _rhs71;
          r0 = _rhs72;
          let _550_v28;
          let _nw112 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _550_v28 = _nw112;
          _550_v28 = _550_v28;
          r0 = false;
        } else {
          let _551___mcc_h4 = (_source11).cf12;
          let _552_cf12 = _551___mcc_h4;
          let _553_v29;
          _553_v29 = _dafny.Seq.of(p0, _module.__default.safeModuloInt(p0, (_this).fm1(p2, globalState)), _module.__default.safeModuloInt(p0, p0), p0, p0);
          _553_v29 = _553_v29;
          r0 = !(p2);
          let _554_v30;
          let _nw113 = new _module.C0();
          _nw113.__ctor(p2, !(p2));
          _554_v30 = _nw113;
          let _555_v31;
          let _nw114 = Array((new BigNumber(21)).toNumber()).fill(false);
          _555_v31 = _nw114;
          let _index94 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_555_v31).length));
          (_555_v31)[_index94] = _554_v30.f5;
          let _index95 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_555_v31).length));
          (_555_v31)[_index95] = p3;
        }
        r0 = !(p3);
        let _556_v32;
        _556_v32 = _dafny.Seq.UnicodeFromString("nutlbe");
        _556_v32 = _dafny.Seq.Concat((_this).f3, _dafny.Seq.UnicodeFromString("xdn"));
        r2 = p0;
        let _557_v33;
        let _nw115 = Array((new BigNumber(8)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = p2;
        _nw115[(_dafny.ONE).toNumber()] = p2;
        _nw115[(new BigNumber(2)).toNumber()] = false;
        _nw115[(new BigNumber(3)).toNumber()] = (_this).fm11(globalState);
        _nw115[(new BigNumber(4)).toNumber()] = p2;
        _nw115[(new BigNumber(5)).toNumber()] = ((p2) ? (false) : (p2));
        _nw115[(new BigNumber(6)).toNumber()] = p3;
        _nw115[(new BigNumber(7)).toNumber()] = p3;
        _557_v33 = _nw115;
        let _index96 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_557_v33).length));
        (_557_v33)[_index96] = p3;
        let _558_v34;
        _558_v34 = _dafny.Seq.of(p2, p2);
        let _559_v35;
        let _nw116 = new _module.C0();
        _nw116.__ctor(false, p2);
        _559_v35 = _nw116;
        let _560_v36;
        _560_v36 = _dafny.Seq.of(_559_v35);
        let _index97 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_557_v33).length));
        (_557_v33)[_index97] = !(((p2) ? ((_558_v34)[_module.__default.safeIndex(p0, new BigNumber((_558_v34).length))]) : (_dafny.Seq.IsProperPrefixOf(_560_v36, _dafny.Seq.of(_559_v35, _559_v35)))));
      }
      let _561_v37;
      _561_v37 = _dafny.MultiSet.fromElements(p0);
      let _hi2 = new BigNumber((_561_v37).cardinality());
      for (let _562_i2 = p0; _562_i2.isLessThan(_hi2); _562_i2 = _562_i2.plus(_dafny.ONE)) {
        let _563_v38;
        _563_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,_562_i2);
        let _564_v39;
        let _nw117 = Array((new BigNumber(12)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = new BigNumber(-17);
        _nw117[(_dafny.ONE).toNumber()] = _562_i2;
        _nw117[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.of(p2)).length);
        _nw117[(new BigNumber(3)).toNumber()] = _562_i2;
        _nw117[(new BigNumber(4)).toNumber()] = p0;
        _nw117[(new BigNumber(5)).toNumber()] = _562_i2;
        _nw117[(new BigNumber(6)).toNumber()] = new BigNumber((_563_v38).length);
        _nw117[(new BigNumber(7)).toNumber()] = p0;
        _nw117[(new BigNumber(8)).toNumber()] = _562_i2;
        _nw117[(new BigNumber(9)).toNumber()] = _562_i2;
        _nw117[(new BigNumber(10)).toNumber()] = _562_i2;
        _nw117[(new BigNumber(11)).toNumber()] = _562_i2;
        _564_v39 = _nw117;
        let _565_v40;
        _565_v40 = _dafny.MultiSet.fromElements(_564_v39, _564_v39, _564_v39);
        let _566_v41;
        _566_v41 = _module.D5.create_DC8(_565_v40);
        let _567_v42;
        let _nw118 = Array((new BigNumber(21)).toNumber());
        _nw118[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_564_v39, _564_v39);
        _nw118[(_dafny.ONE).toNumber()] = _565_v40;
        _nw118[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_564_v39, _564_v39, _564_v39);
        _nw118[(new BigNumber(3)).toNumber()] = ((p3) ? (_565_v40) : (_565_v40));
        _nw118[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_564_v39, _564_v39, _564_v39, _564_v39);
        _nw118[(new BigNumber(5)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(_564_v39)).Intersect(_565_v40);
        _nw118[(new BigNumber(7)).toNumber()] = (_565_v40).Difference(_565_v40);
        _nw118[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(_564_v39, _564_v39, _564_v39);
        _nw118[(new BigNumber(9)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(10)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_564_v39);
        _nw118[(new BigNumber(12)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(13)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(14)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(15)).toNumber()] = (_566_v41).dtor_cf13;
        _nw118[(new BigNumber(16)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(17)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(18)).toNumber()] = _565_v40;
        _nw118[(new BigNumber(19)).toNumber()] = ((_565_v40).update(_564_v39, _module.__default.abs(p0))).Difference(_565_v40);
        _nw118[(new BigNumber(20)).toNumber()] = _565_v40;
        _567_v42 = _nw118;
        let _index98 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_567_v42).length));
        (_567_v42)[_index98] = _565_v40;
        let _index99 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_567_v42).length));
        (_567_v42)[_index99] = _565_v40;
        let _568_v43;
        _568_v43 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_511_v0).Intersect(_511_v0));
        _568_v43 = (_568_v43).Merge((_568_v43).update(p2, _dafny.MultiSet.FromArray(_dafny.Seq.update(_module.__default.fm16(globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm16(globalState)).length)), p3))));
        let _569_v44;
        _569_v44 = new _dafny.CodePoint('h'.codePointAt(0));
        let _570_v45;
        _570_v45 = _module.D0.create_DC1(p2, p2, _569_v44, p0);
        let _571_v46;
        _571_v46 = _dafny.MultiSet.fromElements(_570_v45);
        let _572_v47;
        _572_v47 = _dafny.MultiSet.fromElements(_569_v44, _569_v44);
        let _573_v48;
        _573_v48 = _dafny.Set.fromElements(((p3) ? (!(p2)) : (p3)), (_572_v47).IsProperSubsetOf(_module.__default.fm17((_this).fm11(globalState), globalState)), true, p2);
        let _574_v49;
        _574_v49 = _module.D5.create_DC9(p2, new BigNumber(-599), _dafny.MultiSet.fromElements(p2, p3));
        let _pat_let_tv11 = p2;
        let _rhs73 = _dafny.MultiSet.fromElements(function (_pat_let12_0) {
          return function (_575_dt__update__tmp_h0) {
            return function (_pat_let13_0) {
              return function (_576_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_575_dt__update__tmp_h0).dtor_cf1, _576_dt__update_hcf2_h0, (_575_dt__update__tmp_h0).dtor_cf3, (_575_dt__update__tmp_h0).dtor_cf4);
              }(_pat_let13_0);
            }(_pat_let_tv11);
          }(_pat_let12_0);
        }(_570_v45));
        let _rhs74 = _573_v48;
        let _rhs75 = _562_i2;
        let _rhs76 = ((_this).f3)[_module.__default.safeIndex(((p3) ? ((_574_v49).dtor_cf15) : (_562_i2)), new BigNumber(((_this).f3).length))];
        let _rhs77 = (_dafny.MultiSet.FromArray(_dafny.Seq.of(_562_i2))).IsDisjointFrom(_561_v37);
        _571_v46 = _rhs73;
        _573_v48 = _rhs74;
        r2 = _rhs75;
        _569_v44 = _rhs76;
        r0 = _rhs77;
        let _577_v50;
        _577_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        r2 = (((_this).fm11(globalState)) ? (new BigNumber((_577_v50).length)) : (_562_i2));
      }
      let _578_v51;
      let _nw119 = Array((new BigNumber(13)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = p0;
      _nw119[(_dafny.ONE).toNumber()] = new BigNumber(962);
      _nw119[(new BigNumber(2)).toNumber()] = new BigNumber(((_this).f3).length);
      _nw119[(new BigNumber(3)).toNumber()] = new BigNumber(((_this).f3).length);
      _nw119[(new BigNumber(4)).toNumber()] = p0;
      _nw119[(new BigNumber(5)).toNumber()] = p0;
      _nw119[(new BigNumber(6)).toNumber()] = p0;
      _nw119[(new BigNumber(7)).toNumber()] = p0;
      _nw119[(new BigNumber(8)).toNumber()] = p0;
      _nw119[(new BigNumber(9)).toNumber()] = p0;
      _nw119[(new BigNumber(10)).toNumber()] = p0;
      _nw119[(new BigNumber(11)).toNumber()] = p0;
      _nw119[(new BigNumber(12)).toNumber()] = p0;
      _578_v51 = _nw119;
      let _pat_let_tv12 = p0;
      let _579_v52;
      _579_v52 = _dafny.Seq.of((function (_pat_let14_0) {
        return function (_580_dt__update__tmp_h1) {
          return function (_pat_let15_0) {
            return function (_581_dt__update_hcf9_h0) {
              return _module.D4.create_DC6(_581_dt__update_hcf9_h0, (_580_dt__update__tmp_h1).dtor_cf10, (_580_dt__update__tmp_h1).dtor_cf11);
            }(_pat_let15_0);
          }(_pat_let_tv12);
        }(_pat_let14_0);
      }(_module.D4.create_DC6(new BigNumber((_511_v0).cardinality()), _578_v51, p0))).dtor_cf9, (_dafny.ZERO).minus(p0), p0);
      let _582_v53;
      let _nw120 = new _module.C0();
      _nw120.__ctor(p2, p2);
      _582_v53 = _nw120;
      let _583_v54;
      _583_v54 = _dafny.Set.fromElements(_582_v53, _582_v53);
      let _584_v55;
      _584_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(_582_v53));
      let _585_v56;
      _585_v56 = _dafny.Set.fromElements(p3);
      let _rhs78 = (((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).minus(p0)).isLessThanOrEqualTo(new BigNumber((_579_v52).length));
      let _rhs79 = p0;
      let _rhs80 = ((((_584_v55).contains(p0)) ? ((_584_v55).get(p0)) : (_583_v54))).IsProperSubsetOf(_583_v54);
      let _rhs81 = new BigNumber((((_585_v56).Union(_585_v56)).Union(_585_v56)).length);
      r0 = _rhs78;
      r2 = _rhs79;
      r0 = _rhs80;
      r2 = _rhs81;
      r0 = _582_v53.f5;
      let _rhs82 = p0;
      let _rhs83 = _582_v53.f4;
      let _rhs84 = !((p0).multipliedBy(p0)).isEqualTo((_this).fm1(p3, globalState));
      r2 = _rhs82;
      r0 = _rhs83;
      r0 = _rhs84;
      let _586_v57;
      _586_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
      let _587_v58;
      _587_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_586_v57);
      let _source12 = _module.D5.create_DC10(p2, new BigNumber((_dafny.Seq.UnicodeFromString("tt")).length), _587_v58, true);
      if (_source12.is_DC9) {
        let _588___mcc_h5 = (_source12).cf14;
        let _589___mcc_h6 = (_source12).cf15;
        let _590___mcc_h7 = (_source12).cf16;
        let _591_cf16 = _590___mcc_h7;
        let _592_cf15 = _589___mcc_h6;
        let _593_cf14 = _588___mcc_h5;
        let _594_v59;
        _594_v59 = _module.D4.create_DC6(new BigNumber((_dafny.Set.fromElements(p2)).length), _578_v51, new BigNumber((_dafny.Seq.of((_579_v52)[_module.__default.safeIndex(_592_cf15, new BigNumber((_579_v52).length))])).length));
        let _595_v60;
        _595_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_594_v59);
        let _596_v61;
        _596_v61 = _dafny.Map.Empty.slice().updateUnsafe(_595_v60,p0);
        _596_v61 = (_596_v61).Merge(_596_v61);
        let _597_v62;
        _597_v62 = _dafny.Seq.of(!(p3), _582_v53.f4, ((p3) ? (p2) : (_582_v53.f5)));
        let _rhs85 = _dafny.Seq.Concat(_597_v62, _597_v62);
        let _rhs86 = new BigNumber(((_this).f3).length);
        _597_v62 = _rhs85;
        _592_cf15 = _rhs86;
        let _598_v63;
        _598_v63 = _dafny.Seq.of(_511_v0);
        _511_v0 = (_511_v0).Union((_dafny.MultiSet.fromElements(_582_v53.f4, false)).Difference((_598_v63)[_module.__default.safeIndex(_592_cf15, new BigNumber((_598_v63).length))]));
        let _599_v64;
        let _nw121 = Array((new BigNumber(23)).toNumber()).fill(false);
        _599_v64 = _nw121;
        let _600_v65;
        _600_v65 = _dafny.MultiSet.fromElements(_582_v53, _582_v53);
        let _601_v66;
        _601_v66 = _dafny.Map.Empty.slice().updateUnsafe(_599_v64,_600_v65);
        (_582_v53).f4 = (_597_v62)[_module.__default.safeIndex(((_593_cf14) ? (p0) : (new BigNumber((_601_v66).length))), new BigNumber((_597_v62).length))];
      } else if (_source12.is_DC10) {
        let _602___mcc_h8 = (_source12).cf17;
        let _603___mcc_h9 = (_source12).cf18;
        let _604___mcc_h10 = (_source12).cf19;
        let _605___mcc_h11 = (_source12).cf20;
        let _606_cf20 = _605___mcc_h11;
        let _607_cf19 = _604___mcc_h10;
        let _608_cf18 = _603___mcc_h9;
        let _609_cf17 = _602___mcc_h8;
        let _610_v67;
        _610_v67 = _dafny.Seq.UnicodeFromString("wwqob");
        let _index100 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_578_v51).length));
        (_578_v51)[_index100] = (_579_v52)[_module.__default.safeIndex(new BigNumber((_585_v56).length), new BigNumber((_579_v52).length))];
        let _index101 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_578_v51).length));
        let _rhs87 = _dafny.Seq.Concat(_610_v67, _dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), function (_611_i3) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }));
        let _rhs88 = (new BigNumber(827)).plus(_608_cf18);
        let _rhs89 = _608_cf18;
        let _lhs37 = _578_v51;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_578_v51).length));
        _610_v67 = _rhs87;
        _lhs37[_lhs38] = _rhs88;
        _608_cf18 = _rhs89;
        _606_cf20 = _609_cf17;
        let _index102 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_578_v51).length));
        (_578_v51)[_index102] = (_578_v51)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_578_v51).length))];
        (_582_v53).f5 = _582_v53.f5;
      } else {
        let _612___mcc_h12 = (_source12).cf13;
        let _613_cf13 = _612___mcc_h12;
        let _614_v68;
        _614_v68 = _dafny.Set.fromElements(p0, new BigNumber(-983));
        _511_v0 = (_511_v0).update((_614_v68).IsProperSubsetOf(_module.__default.fm18(_582_v53.f4, p3, _511_v0, globalState)), _module.__default.abs(p0));
        let _615_v69;
        _615_v69 = _dafny.Map.Empty.slice().updateUnsafe(p3,_582_v53.f5);
        let _616_v70;
        _616_v70 = _dafny.MultiSet.fromElements(_615_v69, _615_v69);
        r2 = ((new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_616_v70).cardinality())))).cardinality())).minus((_579_v52)[_module.__default.safeIndex(p0, new BigNumber((_579_v52).length))])).plus(p0);
        let _617_v71;
        let _nw122 = new _module.C0();
        _nw122.__ctor(_582_v53.f5, !(_582_v53.f5));
        _617_v71 = _nw122;
        _511_v0 = _dafny.MultiSet.fromElements(!(false) || (_617_v71.f4));
      }
      r0 = p3;
      let _618_v72;
      _618_v72 = _dafny.Set.fromElements(_561_v37, _561_v37);
      let _619_v73;
      _619_v73 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
      let _620_v74;
      let _nw123 = Array((new BigNumber(5)).toNumber());
      _nw123[(_dafny.ZERO).toNumber()] = _619_v73;
      _nw123[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_582_v53.f5,_582_v53.f4);
      _nw123[(new BigNumber(2)).toNumber()] = _619_v73;
      _nw123[(new BigNumber(3)).toNumber()] = _619_v73;
      _nw123[(new BigNumber(4)).toNumber()] = _619_v73;
      _620_v74 = _nw123;
      let _621_v75;
      _621_v75 = _dafny.Map.Empty.slice().updateUnsafe(_618_v72,_620_v74);
      r1 = (((_621_v75).contains(_dafny.Set.fromElements(_561_v37, _dafny.MultiSet.fromElements(p0), _module.__default.fm19(p0, p0, _582_v53.f4, p0, globalState)))) ? ((_621_v75).get(_dafny.Set.fromElements(_561_v37, _dafny.MultiSet.fromElements(p0), _module.__default.fm19(p0, p0, _582_v53.f4, p0, globalState)))) : (_620_v74));
      r2 = (_dafny.ZERO).minus(((_this).fm1(p2, globalState)).minus((_579_v52)[_module.__default.safeIndex(p0, new BigNumber((_579_v52).length))]));
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _622_v0;
      let _nw124 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _622_v0 = _nw124;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_622_v0).length))) {
        let _623_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_623_i0)) && ((_623_i0).isLessThan(new BigNumber((_622_v0).length))))) {
          (_622_v0)[(_623_i0)] = (_this).f3;
        }
      }
      let _624_i1;
      _624_i1 = _dafny.ZERO;
      L0: {
        while (p2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_624_i1)) {
              break L0;
            }
            _624_i1 = (_624_i1).plus(_dafny.ONE);
            let _625_v1;
            _625_v1 = true;
            _625_v1 = (_this).fm11(globalState);
            let _626_v2;
            let _nw125 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
            _626_v2 = _nw125;
            let _627_v3;
            let _nw126 = Array((new BigNumber(5)).toNumber()).fill(_module.D0.Default());
            _627_v3 = _nw126;
            let _628_v4;
            _628_v4 = _dafny.Map.Empty.slice().updateUnsafe(_627_v3,_625_v1);
            let _index103 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_626_v2).length));
            (_626_v2)[_index103] = (_628_v4).Merge(_628_v4);
            let _index104 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_626_v2).length));
            (_626_v2)[_index104] = _628_v4;
            let _629_v5;
            _629_v5 = new _dafny.CodePoint('h'.codePointAt(0));
            _629_v5 = ((_this).f3)[_module.__default.safeIndex(p0, new BigNumber(((_this).f3).length))];
            let _630_v6;
            _630_v6 = _dafny.Set.fromElements(p2, _625_v1);
            let _631_v7;
            _631_v7 = _dafny.Map.Empty.slice().updateUnsafe((_630_v6).Union(_630_v6),p1);
            _631_v7 = (_631_v7).Merge((_631_v7).Merge(_631_v7));
          }
        }
      }
      let _632_v8;
      _632_v8 = _dafny.Seq.of(true, p2);
      let _633_v9;
      _633_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),(_632_v8)[_module.__default.safeIndex(p1, new BigNumber((_632_v8).length))]);
      let _634_v10;
      _634_v10 = _dafny.Seq.of(_633_v9, _633_v9, _633_v9, _633_v9, _633_v9);
      _634_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
      let _635_v11;
      _635_v11 = false;
      _635_v11 = _635_v11;
      _635_v11 = _635_v11;
      let _636_v12;
      _636_v12 = _dafny.Set.fromElements(p0);
      _636_v12 = _636_v12;
      r0 = (_this).fm1(p2, globalState);
      return r0;
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f2 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T2, _module.T1];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    __ctor(f2) {
      let _this = this;
      (_this)._f2 = f2;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f2),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_637_i0) {
        return new BigNumber((_dafny.Seq.of(_this.f2, _this.f2)).length);
      }))).cardinality()))).Merge(function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f2),true)).Keys.Elements) {
          let _638_v0 = _compr_19;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f2),true)).contains(_638_v0)) {
            _coll19.push([_638_v0,_this.f2]);
          }
        }
        return _coll19;
      }())).Merge((function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length), new BigNumber((_dafny.Set.fromElements(true, true, true, !(false), false)).length)))).Elements) {
          let _639_v1 = _compr_20;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length), new BigNumber((_dafny.Set.fromElements(true, true, true, !(false), false)).length))), _639_v1)) {
            _coll20.push([_639_v1,new BigNumber((function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
                let _640_v2 = _compr_21;
                if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))).contains(_640_v2)) {
                  _coll21.push([_640_v2,_this.f2]);
                }
              }
              return _coll21;
            }()).length)]);
          }
        }
        return _coll20;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f2)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,_this.f2)).length))));
    };
    fm1(p0, globalState) {
      let _this = this;
      return _this.f2;
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC10(true, _this.f2, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), function (_641_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
}),function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of _dafny.IntegerRange(new BigNumber(934), new BigNumber(570))) {
    let _642_v0 = _compr_22;
    if (((new BigNumber(934)).isLessThanOrEqualTo(_642_v0)) && ((_642_v0).isLessThan(new BigNumber(570)))) {
      _coll22.push([_module.__default.safeModuloInt(_642_v0, _this.f2),false]);
    }
  }
  return _coll22;
}()), true)).dtor_cf18,!(true))).length), _this.f2)).Union(_dafny.Set.fromElements(_this.f2, _this.f2, new BigNumber((_dafny.Seq.UnicodeFromString("cawq")).length), _this.f2, _this.f2))).equals((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ydjxbae")).length))).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_this.f2)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(911))).length))));
    };
    fm24(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC14(_dafny.Seq.of(_this.f2, new BigNumber((_dafny.Seq.of(new BigNumber(-59))).length)))).dtor_cf26,_this.f2);
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _643_v0;
      _643_v0 = _dafny.Set.fromElements(_this.f2);
      let _644_v1;
      _644_v1 = _dafny.Map.Empty.slice().updateUnsafe(p3,_643_v0);
      let _645_v2;
      _645_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
      let _646_v3;
      _646_v3 = _dafny.Seq.UnicodeFromString("rftn");
      let _647_v4;
      let _nw127 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _647_v4 = _nw127;
      let _648_v5;
      let _nw128 = Array((new BigNumber(18)).toNumber());
      _nw128[(_dafny.ZERO).toNumber()] = p3;
      _nw128[(_dafny.ONE).toNumber()] = p3;
      _nw128[(new BigNumber(2)).toNumber()] = p3;
      _nw128[(new BigNumber(3)).toNumber()] = p2;
      _nw128[(new BigNumber(4)).toNumber()] = p3;
      _nw128[(new BigNumber(5)).toNumber()] = false;
      _nw128[(new BigNumber(6)).toNumber()] = p3;
      _nw128[(new BigNumber(7)).toNumber()] = p2;
      _nw128[(new BigNumber(8)).toNumber()] = p3;
      _nw128[(new BigNumber(9)).toNumber()] = p2;
      _nw128[(new BigNumber(10)).toNumber()] = p3;
      _nw128[(new BigNumber(11)).toNumber()] = p3;
      _nw128[(new BigNumber(12)).toNumber()] = true;
      _nw128[(new BigNumber(13)).toNumber()] = p3;
      _nw128[(new BigNumber(14)).toNumber()] = p2;
      _nw128[(new BigNumber(15)).toNumber()] = p3;
      _nw128[(new BigNumber(16)).toNumber()] = p2;
      _nw128[(new BigNumber(17)).toNumber()] = p2;
      _648_v5 = _nw128;
      let _649_v6;
      _649_v6 = _dafny.Map.Empty.slice().updateUnsafe(_648_v5,p0);
      let _650_v7;
      _650_v7 = _module.D4.create_DC6((_this).fm1(p3, globalState), _647_v4, new BigNumber((_649_v6).length));
      let _651_v8;
      _651_v8 = _module.D4.create_DC7(_650_v7);
      let _652_v9;
      _652_v9 = _dafny.Seq.of(_651_v8);
      let _653_v10;
      _653_v10 = _dafny.Seq.of(p3, p2, p3, p3, p2);
      let _654_v11;
      _654_v11 = new _dafny.CodePoint('e'.codePointAt(0));
      let _655_v12;
      let _nw129 = Array((new BigNumber(16)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((((_644_v1).contains(p3)) ? ((_644_v1).get(p3)) : (_643_v0))).length), new BigNumber((_645_v2).length));
      _nw129[(_dafny.ONE).toNumber()] = new BigNumber((_646_v3).length);
      _nw129[(new BigNumber(2)).toNumber()] = p0;
      _nw129[(new BigNumber(3)).toNumber()] = new BigNumber((_652_v9).length);
      _nw129[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_this.f2, _this.f2);
      _nw129[(new BigNumber(5)).toNumber()] = new BigNumber((_653_v10).length);
      _nw129[(new BigNumber(6)).toNumber()] = p0;
      _nw129[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_654_v11)).cardinality());
      _nw129[(new BigNumber(8)).toNumber()] = (_this).fm1(false, globalState);
      _nw129[(new BigNumber(9)).toNumber()] = _this.f2;
      _nw129[(new BigNumber(10)).toNumber()] = p0;
      _nw129[(new BigNumber(11)).toNumber()] = _this.f2;
      _nw129[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_this.f2);
      _nw129[(new BigNumber(13)).toNumber()] = p0;
      _nw129[(new BigNumber(14)).toNumber()] = (new BigNumber(378)).plus(new BigNumber(236));
      _nw129[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_653_v10)).cardinality());
      _655_v12 = _nw129;
      let _index105 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
      (_647_v4)[_index105] = new BigNumber((_646_v3).length);
      let _index106 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
      let _rhs90 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_646_v3).length), (_dafny.ZERO).minus((_this).fm1(p2, globalState))))).minus(_this.f2);
      let _rhs91 = (_dafny.ZERO).minus(_this.f2);
      let _lhs39 = _647_v4;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
      r2 = _rhs90;
      _lhs39[_lhs40] = _rhs91;
      if ((_this.f2).isLessThanOrEqualTo((_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))])) {
        let _656_v13;
        _656_v13 = (_this).fm0(globalState);
        let _source13 = _656_v13;
        let _657___mcc_h0 = _source13;
        let _658_cf6 = _657___mcc_h0;
        let _index107 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        (_647_v4)[_index107] = (_dafny.ZERO).minus((_this).fm1(p3, globalState));
        let _index108 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        (_647_v4)[_index108] = p0;
        r0 = p2;
        let _index109 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        (_647_v4)[_index109] = _module.__default.safeModuloInt(p0, (_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))]);
        let _659_v14;
        _659_v14 = _dafny.Set.fromElements(p2);
        _659_v14 = (((p0).isEqualTo(p0)) ? ((_659_v14).Intersect(_659_v14)) : (_659_v14));
        let _660_v15;
        _660_v15 = _dafny.Seq.of(new BigNumber(293), (_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))]);
        let _661_v16;
        _661_v16 = _dafny.Map.Empty.slice().updateUnsafe(_660_v15,new BigNumber((_646_v3).length));
        let _662_v17;
        _662_v17 = _dafny.Seq.of(_656_v13, _661_v16);
        r0 = (p3) && (!_dafny.areEqual(_662_v17, _662_v17));
        let _663_v18;
        _663_v18 = _module.D4.create_DC6(p0, _655_v12, _this.f2);
        let _664_v19;
        _664_v19 = _dafny.Map.Empty.slice().updateUnsafe((_663_v18).dtor_cf10,p3);
        let _665_v20;
        let _nw130 = new _module.C0();
        _nw130.__ctor(p3, (((_664_v19).contains(_655_v12)) ? ((_664_v19).get(_655_v12)) : (p3)));
        _665_v20 = _nw130;
        (_this).f2 = (_module.__default.safeModuloInt(p0, p0)).plus((_this.f2).multipliedBy(_this.f2));
      } else {
        let _666_v21;
        _666_v21 = _module.D7.create_DC15((_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))], _654_v11);
        let _667_v22;
        _667_v22 = _dafny.Map.Empty.slice().updateUnsafe(_666_v21,p2);
        let _668_v23;
        _668_v23 = _dafny.MultiSet.fromElements(_654_v11);
        let _669_v24;
        _669_v24 = _dafny.MultiSet.fromElements(new BigNumber((_645_v2).length));
        let _670_v25;
        _670_v25 = _dafny.Seq.of(_this.f2, p0);
        let _index110 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        let _rhs92 = (new BigNumber((_667_v22).length)).isLessThan(_module.__default.safeModuloInt(p0, _this.f2));
        let _rhs93 = _module.__default.safeDivisionInt((_this).fm1(p3, globalState), p0);
        let _rhs94 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber((_653_v10).length)))).multipliedBy(new BigNumber((_668_v23).cardinality()));
        let _rhs95 = _module.__default.safeModuloInt((_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))], _this.f2);
        let _rhs96 = new BigNumber(((_669_v24).Union((_dafny.MultiSet.FromArray(_670_v25)).Union(_669_v24))).cardinality());
        let _lhs41 = _this;
        let _lhs42 = _this;
        let _lhs43 = _647_v4;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        r0 = _rhs92;
        _lhs41.f2 = _rhs93;
        r2 = _rhs94;
        _lhs42.f2 = _rhs95;
        _lhs43[_lhs44] = _rhs96;
        let _index111 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        (_647_v4)[_index111] = new BigNumber(814);
        _653_v10 = _653_v10;
        let _671_v26;
        _671_v26 = _dafny.MultiSet.fromElements(p2);
        let _index112 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length));
        (_647_v4)[_index112] = _module.__default.safeDivisionInt(p0, new BigNumber((((p3) ? (_671_v26) : (_dafny.MultiSet.fromElements(p3)))).cardinality()));
        let _672_v27;
        _672_v27 = _dafny.Map.Empty.slice().updateUnsafe(_646_v3,_dafny.Map.Empty.slice().updateUnsafe(_this.f2,false));
        let _index113 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_648_v5).length));
        (_648_v5)[_index113] = (_module.D5.create_DC10(p2, new BigNumber((_643_v0).length), _672_v27, p2)).dtor_cf20;
        let _index114 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_648_v5).length));
        (_648_v5)[_index114] = p2;
      }
      _648_v5 = _648_v5;
      let _673_v28;
      _673_v28 = _dafny.Seq.of(new BigNumber(99));
      let _pat_let_tv13 = _654_v11;
      let _pat_let_tv14 = p3;
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = _647_v4;
      let _pat_let_tv17 = _647_v4;
      r2 = function (_source14) {
        if (_source14.is_DC1) {
          let _674___mcc_h1 = (_source14).cf1;
          let _675___mcc_h2 = (_source14).cf2;
          let _676___mcc_h3 = (_source14).cf3;
          let _677___mcc_h4 = (_source14).cf4;
          let _678_cf4 = _677___mcc_h4;
          let _679_cf3 = _676___mcc_h3;
          let _680_cf2 = _675___mcc_h2;
          let _681_cf1 = _674___mcc_h1;
          return (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_678_cf4,_pat_let_tv13)).length)));
        } else {
          let _682___mcc_h5 = (_source14).cf0;
          let _683_cf0 = _682___mcc_h5;
          return _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv14,_pat_let_tv15)).length), (_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_pat_let_tv16).length))]);
        }
      }(_module.__default.fm25(new BigNumber((_673_v28).length), (_this).fm8(_646_v3, p2, globalState), _dafny.MultiSet.fromElements(new BigNumber(245), _this.f2), _dafny.Seq.of(p2), globalState));
      let _684_v29;
      _684_v29 = _module.D7.create_DC14(_673_v28);
      let _685_v30;
      _685_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,p0);
      let _686_v31;
      let _nw131 = Array((new BigNumber(27)).toNumber());
      _nw131[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_673_v28, _673_v28);
      _nw131[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f2), _673_v28);
      _nw131[(new BigNumber(2)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(3)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(4)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_this.f2);
      _nw131[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_687_p0) => function (_688_i1) {
        return _687_p0;
      })(p0));
      _nw131[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_this.f2, (_this).fm1(p2, globalState));
      _nw131[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_673_v28, _dafny.Seq.of(p0));
      _nw131[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_673_v28, _673_v28);
      _nw131[(new BigNumber(10)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), ((_689_p0) => function (_690_i2) {
        return _689_p0;
      })(p0));
      _nw131[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_673_v28, _673_v28);
      _nw131[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f2, new BigNumber((_643_v0).length), (_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))]), _673_v28);
      _nw131[(new BigNumber(14)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f2), _dafny.Seq.of(_this.f2, (_673_v28)[_module.__default.safeIndex(_this.f2, new BigNumber((_673_v28).length))]));
      _nw131[(new BigNumber(16)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(17)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_691_v3) => function (_692_i3) {
        return new BigNumber((_691_v3).length);
      })(_646_v3));
      _nw131[(new BigNumber(19)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(20)).toNumber()] = (_684_v29).dtor_cf26;
      _nw131[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_647_v4)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_647_v4).length))], new BigNumber((_685_v30).length)), _dafny.Seq.update(_673_v28, _module.__default.safeIndex(p0, new BigNumber((_673_v28).length)), _this.f2));
      _nw131[(new BigNumber(22)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(23)).toNumber()] = _673_v28;
      _nw131[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), function (_693_i4) {
        return new BigNumber(334);
      });
      _nw131[(new BigNumber(25)).toNumber()] = _dafny.Seq.of(_this.f2);
      _nw131[(new BigNumber(26)).toNumber()] = _673_v28;
      _686_v31 = _nw131;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_686_v31).length))) {
        let _694_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_694_i0)) && ((_694_i0).isLessThan(new BigNumber((_686_v31).length))))) {
          (_686_v31)[(_694_i0)] = _673_v28;
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_647_v4).length))) {
        let _695_i5 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_695_i5)) && ((_695_i5).isLessThan(new BigNumber((_647_v4).length))))) {
          (_647_v4)[(_695_i5)] = (_695_i5).plus(p0);
        }
      }
      r0 = (!(p2) || (false)) || (p3);
      let _696_v32;
      let _init10 = ((_697_p3) => function (_698_i6) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_697_p3,_697_p3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_697_p3));
      })(p3);
      let _nw132 = Array((_dafny.ONE).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw132.length); _i0_10++) {
        _nw132[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _696_v32 = _nw132;
      r1 = _696_v32;
      r2 = _module.__default.safeDivisionInt(p0, (_this.f2).minus(p0));
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (p2) {
        let _699_v0;
        let _nw133 = Array((new BigNumber(24)).toNumber()).fill([]);
        _699_v0 = _nw133;
        let _700_v1;
        _700_v1 = _699_v0;
        _700_v1 = _699_v0;
        let _701_v2;
        _701_v2 = _dafny.Seq.of(p2);
        let _702_v3;
        _702_v3 = _dafny.Seq.of(p2, (_701_v2)[_module.__default.safeIndex(p1, new BigNumber((_701_v2).length))]);
        let _703_v4;
        _703_v4 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kvtp")).length), p1, (_this.f2).plus(new BigNumber(514)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p2), _702_v3)).length));
        let _704_v5;
        _704_v5 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        (_this).f2 = (((_703_v4).contains(_this.f2)) ? ((_703_v4).get(_this.f2)) : (new BigNumber(((_704_v5).Merge(_704_v5)).length)));
        let _705_v6;
        _705_v6 = _dafny.Seq.UnicodeFromString("h");
        let _706_v7;
        _706_v7 = _dafny.Seq.of(_705_v6, _705_v6);
        _706_v7 = _dafny.Seq.Concat(_706_v7, _dafny.Seq.of(_705_v6, _705_v6));
        let _707_v8;
        let _nw134 = new _module.C1();
        _nw134.__ctor(_dafny.Seq.UnicodeFromString("bawn"));
        _707_v8 = _nw134;
        if ((new BigNumber((function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(-299), new BigNumber(388))) {
            let _708_v9 = _compr_23;
            if (((new BigNumber(-299)).isLessThanOrEqualTo(_708_v9)) && ((_708_v9).isLessThan(new BigNumber(388)))) {
              _coll23.add((_708_v9).plus(_this.f2));
            }
          }
          return _coll23;
        }()).length)).isLessThanOrEqualTo(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_702_v3).length)), _this.f2))) {
          let _709_v10;
          let _nw135 = new _module.C1();
          _nw135.__ctor(_dafny.Seq.UnicodeFromString("qlxtt"));
          _709_v10 = _nw135;
          let _710_v11;
          _710_v11 = _dafny.Map.Empty.slice().updateUnsafe(_705_v6,_dafny.Map.Empty.slice().updateUnsafe(p1,p2));
          let _711_v12;
          let _nw136 = Array((_dafny.ONE).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = _this.f2;
          _711_v12 = _nw136;
          let _712_v13;
          _712_v13 = _dafny.Set.fromElements(_711_v12);
          let _713_v14;
          let _nw137 = new _module.C0();
          _nw137.__ctor(!(p2) || ((_module.D5.create_DC10(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(48)), ((_714_v6) => function (_715_i0) {
  return (_714_v6)[_module.__default.safeIndex(_this.f2, new BigNumber((_714_v6).length))];
})(_705_v6))).length), _710_v11, p2)).dtor_cf17), (_712_v13).equals(_712_v13));
          _713_v14 = _nw137;
          let _716_v15;
          _716_v15 = _module.D0.create_DC0(!(!(p2)));
          let _717_v16;
          _717_v16 = new _dafny.CodePoint('c'.codePointAt(0));
          let _718_v17;
          let _nw138 = Array((new BigNumber(5)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _716_v15;
          _nw138[(_dafny.ONE).toNumber()] = _module.__default.fm25(new BigNumber((_module.__default.fm26(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_717_v16,_dafny.Map.Empty.slice().updateUnsafe(_713_v14.f5,p2))).length), _713_v14.f4, _713_v14.f5, _713_v14.f5, globalState)).length), _713_v14.f4, _703_v4, _702_v3, globalState);
          _nw138[(new BigNumber(2)).toNumber()] = _module.D0.create_DC0(_713_v14.f5);
          _nw138[(new BigNumber(3)).toNumber()] = _716_v15;
          _nw138[(new BigNumber(4)).toNumber()] = _716_v15;
          _718_v17 = _nw138;
          let _index115 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_718_v17).length));
          (_718_v17)[_index115] = _716_v15;
          let _index116 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_718_v17).length));
          (_718_v17)[_index116] = _module.__default.fm25(p0, !(_713_v14.f4), _703_v4, _dafny.Seq.update(_702_v3, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_702_v3).length)), !(true)), globalState);
          let _index117 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length));
          (_711_v12)[_index117] = (_this.f2).plus(p0);
          let _index118 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length));
          (_711_v12)[_index118] = (new BigNumber(-544)).plus(p0);
          let _719_v18;
          _719_v18 = _dafny.Set.fromElements(new BigNumber((_705_v6).length), new BigNumber((_702_v3).length));
          let _720_v19;
          _720_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,((_dafny.ZERO).minus(new BigNumber((_719_v18).length))).minus((_711_v12)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length))]));
          let _721_v20;
          _721_v20 = _dafny.Seq.of((_module.__default.fm27(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,!(false))).length), p2, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_720_v19).length),p1)));
          let _index119 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length));
          let _rhs97 = (_721_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p2)).length), new BigNumber((_721_v20).length))];
          let _rhs98 = !(_713_v14.f5);
          let _rhs99 = (_711_v12)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length))];
          let _rhs100 = _module.__default.safeDivisionInt(_this.f2, (new BigNumber(((_709_v10).f3).length)).multipliedBy(_this.f2));
          let _rhs101 = _720_v19;
          let _lhs45 = _713_v14;
          let _lhs46 = _711_v12;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_711_v12).length));
          _720_v19 = _rhs97;
          _lhs45.f4 = _rhs98;
          r0 = _rhs99;
          _lhs46[_lhs47] = _rhs100;
          _720_v19 = _rhs101;
        } else {
          let _722_v21;
          let _nw139 = new _module.C0();
          _nw139.__ctor(!((p0).isLessThanOrEqualTo(p1)), !(p2));
          _722_v21 = _nw139;
          let _723_v22;
          _723_v22 = new _dafny.CodePoint('i'.codePointAt(0));
          let _724_v23;
          _724_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_704_v5).length),_723_v22);
          (_722_v21).f4 = !(!_dafny.Seq.contains((_707_v8).f3, (((_724_v23).contains((_dafny.ZERO).minus(new BigNumber(((_707_v8).f3).length)))) ? ((_724_v23).get((_dafny.ZERO).minus(new BigNumber(((_707_v8).f3).length)))) : (_723_v22))));
          let _725_v24;
          let _nw140 = Array((new BigNumber(28)).toNumber()).fill(_module.D7.Default());
          _725_v24 = _nw140;
          let _726_v25;
          _726_v25 = _dafny.Map.Empty.slice().updateUnsafe(_723_v22,_725_v24);
          let _727_v26;
          _727_v26 = _dafny.MultiSet.fromElements(_725_v24, (((_726_v25).contains(new _dafny.CodePoint('s'.codePointAt(0)))) ? ((_726_v25).get(new _dafny.CodePoint('s'.codePointAt(0)))) : (_725_v24)));
          let _728_v27;
          _728_v27 = _dafny.Map.Empty.slice().updateUnsafe(_727_v26,_this.f2);
          let _729_v28;
          _729_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,_722_v21.f5);
          _728_v27 = (_728_v27).update((_727_v26).update(_725_v24, _module.__default.abs(p1)), new BigNumber((_729_v28).length));
          let _730_v29;
          _730_v29 = _module.D0.create_DC0(p2);
          let _731_v30;
          let _nw141 = Array((new BigNumber(19)).toNumber());
          _nw141[(_dafny.ZERO).toNumber()] = _722_v21.f4;
          _nw141[(_dafny.ONE).toNumber()] = (_this).fm8((_707_v8).f3, false, globalState);
          _nw141[(new BigNumber(2)).toNumber()] = p2;
          _nw141[(new BigNumber(3)).toNumber()] = _722_v21.f4;
          _nw141[(new BigNumber(4)).toNumber()] = _722_v21.f5;
          _nw141[(new BigNumber(5)).toNumber()] = _722_v21.f4;
          _nw141[(new BigNumber(6)).toNumber()] = false;
          _nw141[(new BigNumber(7)).toNumber()] = _722_v21.f5;
          _nw141[(new BigNumber(8)).toNumber()] = !(!(false));
          _nw141[(new BigNumber(9)).toNumber()] = (_722_v21.f5) || (p2);
          _nw141[(new BigNumber(10)).toNumber()] = (p2) && (_722_v21.f4);
          _nw141[(new BigNumber(11)).toNumber()] = (_730_v29).dtor_cf0;
          _nw141[(new BigNumber(12)).toNumber()] = _722_v21.f4;
          _nw141[(new BigNumber(13)).toNumber()] = true;
          _nw141[(new BigNumber(14)).toNumber()] = p2;
          _nw141[(new BigNumber(15)).toNumber()] = true;
          _nw141[(new BigNumber(16)).toNumber()] = (_703_v4).contains((((_703_v4).contains(p0)) ? ((_703_v4).get(p0)) : (p1)));
          _nw141[(new BigNumber(17)).toNumber()] = true;
          _nw141[(new BigNumber(18)).toNumber()] = _722_v21.f4;
          _731_v30 = _nw141;
          let _index120 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          (_731_v30)[_index120] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(208)), ((_732_p0) => function (_733_i1) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_732_p0)).cardinality()),_this.f2)).length);
          })(p0))).length)).isLessThan(new BigNumber(((_707_v8).f3).length));
          let _index121 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          (_731_v30)[_index121] = _dafny.Seq.IsProperPrefixOf((_707_v8).f3, (_707_v8).f3);
          let _734_v31;
          _734_v31 = _dafny.Map.Empty.slice().updateUnsafe(_731_v30,_this.f2);
          let _735_v32;
          let _nw142 = Array((new BigNumber(12)).toNumber());
          _nw142[(_dafny.ZERO).toNumber()] = p0;
          _nw142[(_dafny.ONE).toNumber()] = p0;
          _nw142[(new BigNumber(2)).toNumber()] = new BigNumber((_729_v28).length);
          _nw142[(new BigNumber(3)).toNumber()] = _this.f2;
          _nw142[(new BigNumber(4)).toNumber()] = new BigNumber(864);
          _nw142[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_this.f2);
          _nw142[(new BigNumber(6)).toNumber()] = (_707_v8).fm1(p2, globalState);
          _nw142[(new BigNumber(7)).toNumber()] = p1;
          _nw142[(new BigNumber(8)).toNumber()] = new BigNumber((_734_v31).length);
          _nw142[(new BigNumber(9)).toNumber()] = p0;
          _nw142[(new BigNumber(10)).toNumber()] = new BigNumber(321);
          _nw142[(new BigNumber(11)).toNumber()] = p0;
          _735_v32 = _nw142;
          let _736_v33;
          _736_v33 = _dafny.MultiSet.fromElements(_735_v32, _735_v32);
          let _737_v34;
          _737_v34 = _module.D5.create_DC10(_722_v21.f4, _this.f2, _dafny.Map.Empty.slice().updateUnsafe((_707_v8).f3,_729_v28), _722_v21.f5);
          let _738_v35;
          _738_v35 = _dafny.Map.Empty.slice().updateUnsafe((p0).multipliedBy(new BigNumber(87)),_dafny.Set.fromElements(_737_v34));
          let _739_v36;
          _739_v36 = _dafny.Set.fromElements(_722_v21.f4);
          let _740_v37;
          _740_v37 = _dafny.MultiSet.fromElements(!(_722_v21.f4));
          let _index122 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          let _index123 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          let _rhs102 = _722_v21.f4;
          let _rhs103 = (_736_v33).Intersect(_736_v33);
          let _rhs104 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_705_v6, _module.__default.fm28(p2, _722_v21.f4, _722_v21.f4, _739_v36, globalState)), (_707_v8).f3);
          let _rhs105 = (_738_v35).Merge(_738_v35);
          let _rhs106 = _dafny.Seq.of((_740_v37).IsProperSubsetOf(_740_v37));
          let _lhs48 = _731_v30;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          let _lhs50 = _731_v30;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_731_v30).length));
          _lhs48[_lhs49] = _rhs102;
          _736_v33 = _rhs103;
          _lhs50[_lhs51] = _rhs104;
          _738_v35 = _rhs105;
          _702_v3 = _rhs106;
        }
      } else {
        let _741_v38;
        _741_v38 = true;
        let _742_v39;
        _742_v39 = _dafny.Seq.of(p2, p2, !(p2));
        let _743_v41;
        _743_v41 = _dafny.Seq.of(p1, p0);
        _741_v38 = (_this).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), function (_744_i2) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }), (_742_v39)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (_743_v41).Elements) {
            let _745_v40 = _compr_24;
            if (_dafny.Seq.contains(_743_v41, _745_v40)) {
              _coll24.push([_module.__default.safeDivisionInt(_745_v40, (_dafny.ZERO).minus(_this.f2)),_741_v38]);
            }
          }
          return _coll24;
        }()).length), new BigNumber((_742_v39).length))], globalState);
        (_this).f2 = _module.__default.safeDivisionInt(p1, p0);
        let _746_v42;
        _746_v42 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_747_i3) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }));
        let _748_v43;
        let _nw143 = new _module.C1();
        _nw143.__ctor((_746_v42)[_module.__default.safeIndex(p1, new BigNumber((_746_v42).length))]);
        _748_v43 = _nw143;
        let _source15 = _module.__default.fm29(globalState);
        if (_source15.is_DC9) {
          let _749___mcc_h0 = (_source15).cf14;
          let _750___mcc_h1 = (_source15).cf15;
          let _751___mcc_h2 = (_source15).cf16;
          let _752_cf16 = _751___mcc_h2;
          let _753_cf15 = _750___mcc_h1;
          let _754_cf14 = _749___mcc_h0;
          let _755_v44;
          let _init11 = function (_756_i4) {
            return _module.D7.create_DC15(new BigNumber(-344), new _dafny.CodePoint('a'.codePointAt(0)));
          };
          let _nw144 = Array((new BigNumber(12)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw144.length); _i0_11++) {
            _nw144[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _755_v44 = _nw144;
          let _757_v45;
          _757_v45 = _dafny.Set.fromElements(_755_v44, _755_v44);
          _757_v45 = _757_v45;
          (_this).f2 = p0;
          _753_cf15 = new BigNumber(((_748_v43).f3).length);
          _754_cf14 = _754_cf14;
        } else if (_source15.is_DC10) {
          let _758___mcc_h3 = (_source15).cf17;
          let _759___mcc_h4 = (_source15).cf18;
          let _760___mcc_h5 = (_source15).cf19;
          let _761___mcc_h6 = (_source15).cf20;
          let _762_cf20 = _761___mcc_h6;
          let _763_cf19 = _760___mcc_h5;
          let _764_cf18 = _759___mcc_h4;
          let _765_cf17 = _758___mcc_h3;
          _765_cf17 = p2;
          let _766_v46;
          _766_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _767_v47;
          _767_v47 = _module.D5.create_DC10(_765_cf17, new BigNumber(637), _dafny.Map.Empty.slice().updateUnsafe((_748_v43).f3,_766_v46), _765_cf17);
          let _768_v48;
          _768_v48 = _dafny.Set.fromElements((_767_v47).dtor_cf18);
          let _769_v49;
          let _nw145 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _769_v49 = _nw145;
          let _index124 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_769_v49).length));
          (_769_v49)[_index124] = _this.f2;
          let _770_v50;
          _770_v50 = _dafny.Map.Empty.slice().updateUnsafe(_769_v49,p2);
          let _index125 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_769_v49).length));
          let _rhs107 = (_dafny.Set.fromElements(p1)).Intersect(_dafny.Set.fromElements(p0));
          let _rhs108 = new BigNumber((_dafny.Seq.Concat((_748_v43).f3, (_748_v43).f3)).length);
          let _rhs109 = _764_cf18;
          let _rhs110 = _769_v49;
          let _rhs111 = (((_770_v50).contains(_769_v49)) ? ((_770_v50).get(_769_v49)) : ((_742_v39)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_742_v39).length))]));
          let _lhs52 = _this;
          let _lhs53 = _769_v49;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_769_v49).length));
          _768_v48 = _rhs107;
          _lhs52.f2 = _rhs108;
          _lhs53[_lhs54] = _rhs109;
          _769_v49 = _rhs110;
          _741_v38 = _rhs111;
          r0 = _this.f2;
          let _771_v51;
          _771_v51 = _dafny.MultiSet.fromElements((_748_v43).f3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(48)), function (_772_i5) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }));
          let _773_v52;
          _773_v52 = _dafny.Map.Empty.slice().updateUnsafe(_771_v51,(new BigNumber(48)).isLessThanOrEqualTo(p0));
          _773_v52 = (_773_v52).update(_module.__default.fm30(_765_cf17, globalState), (_this.f2).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("okwhp")).length)));
        } else {
          let _774___mcc_h7 = (_source15).cf13;
          let _775_cf13 = _774___mcc_h7;
          _741_v38 = _741_v38;
          let _776_v53;
          let _nw146 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _776_v53 = _nw146;
          let _777_v54;
          let _778_v55;
          let _779_v56;
          let _out19;
          let _out20;
          let _out21;
          let _outcollector5 = (_748_v43).m1((_748_v43).fm1(false, globalState), _776_v53, (_742_v39)[_module.__default.safeIndex(p0, new BigNumber((_742_v39).length))], p2, globalState);
          _out19 = _outcollector5[0];
          _out20 = _outcollector5[1];
          _out21 = _outcollector5[2];
          _777_v54 = _out19;
          _778_v55 = _out20;
          _779_v56 = _out21;
          let _780_v57;
          _780_v57 = _module.D0.create_DC0((p0).isLessThan(_this.f2));
          let _rhs112 = new BigNumber(-629);
          let _rhs113 = (_748_v43).fm1(_777_v54, globalState);
          let _rhs114 = _780_v57;
          let _rhs115 = (_748_v43).fm11(globalState);
          let _lhs55 = _this;
          r0 = _rhs112;
          _lhs55.f2 = _rhs113;
          _780_v57 = _rhs114;
          _741_v38 = _rhs115;
          let _781_v58;
          _781_v58 = _dafny.Set.fromElements(_this.f2, _779_v56);
          let _782_v59;
          _782_v59 = _dafny.Map.Empty.slice().updateUnsafe(_741_v38,_779_v56);
          _741_v38 = (_dafny.Set.fromElements(new BigNumber(((_748_v43).f3).length))).IsSubsetOf((_781_v58).Intersect(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_782_v59).length)), _this.f2, (_dafny.ZERO).minus(_779_v56), _779_v56, new BigNumber(-225))));
        }
        let _783_v60;
        _783_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(531),_741_v38);
        let _784_v61;
        _784_v61 = _dafny.Map.Empty.slice().updateUnsafe((_748_v43).f3,_783_v60);
        let _785_v62;
        _785_v62 = _module.D5.create_DC10(_741_v38, p1, _784_v61, p2);
        let _786_v63;
        _786_v63 = _dafny.MultiSet.fromElements(p2, p2, _741_v38, _741_v38);
        let _787_v64;
        _787_v64 = _dafny.Set.fromElements(_this.f2);
        let _788_v65;
        _788_v65 = _dafny.Map.Empty.slice().updateUnsafe(_741_v38,p2);
        let _789_v66;
        let _nw147 = Array((new BigNumber(25)).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = p1;
        _nw147[(_dafny.ONE).toNumber()] = _this.f2;
        _nw147[(new BigNumber(2)).toNumber()] = ((p2) ? (new BigNumber((_dafny.Set.fromElements(new BigNumber(-236))).length)) : (p1));
        _nw147[(new BigNumber(3)).toNumber()] = _this.f2;
        _nw147[(new BigNumber(4)).toNumber()] = p1;
        _nw147[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt((_785_v62).dtor_cf18, _this.f2);
        _nw147[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_this.f2);
        _nw147[(new BigNumber(7)).toNumber()] = p1;
        _nw147[(new BigNumber(8)).toNumber()] = ((p2) ? (p1) : ((_dafny.ZERO).minus((_dafny.ZERO).minus((((_786_v63).contains(_741_v38)) ? ((_786_v63).get(_741_v38)) : (_this.f2))))));
        _nw147[(new BigNumber(9)).toNumber()] = (new BigNumber((_dafny.Seq.of(_this.f2, new BigNumber(619))).length)).minus(p1);
        _nw147[(new BigNumber(10)).toNumber()] = (new BigNumber(241)).plus(new BigNumber(-534));
        _nw147[(new BigNumber(11)).toNumber()] = _this.f2;
        _nw147[(new BigNumber(12)).toNumber()] = ((_741_v38) ? ((_dafny.ZERO).minus(p0)) : (new BigNumber((_dafny.Seq.UnicodeFromString("uyp")).length)));
        _nw147[(new BigNumber(13)).toNumber()] = new BigNumber(276);
        _nw147[(new BigNumber(14)).toNumber()] = p0;
        _nw147[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_743_v41).length), new BigNumber((_742_v39).length));
        _nw147[(new BigNumber(16)).toNumber()] = ((_741_v38) ? (new BigNumber(308)) : (p1));
        _nw147[(new BigNumber(17)).toNumber()] = _this.f2;
        _nw147[(new BigNumber(18)).toNumber()] = new BigNumber((_787_v64).length);
        _nw147[(new BigNumber(19)).toNumber()] = (new BigNumber((_788_v65).length)).minus(p0);
        _nw147[(new BigNumber(20)).toNumber()] = new BigNumber((_786_v63).cardinality());
        _nw147[(new BigNumber(21)).toNumber()] = _this.f2;
        _nw147[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), p0);
        _nw147[(new BigNumber(23)).toNumber()] = new BigNumber(((_748_v43).f3).length);
        _nw147[(new BigNumber(24)).toNumber()] = p0;
        _789_v66 = _nw147;
        let _index126 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_789_v66).length));
        (_789_v66)[_index126] = _module.__default.safeModuloInt(p0, new BigNumber((_743_v41).length));
        let _index127 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_789_v66).length));
        (_789_v66)[_index127] = p0;
      }
      let _790_v67;
      _790_v67 = _dafny.Seq.UnicodeFromString("hhsiw");
      _790_v67 = _790_v67;
      let _791_v68;
      _791_v68 = new _dafny.CodePoint('d'.codePointAt(0));
      let _792_v69;
      _792_v69 = _dafny.Map.Empty.slice().updateUnsafe(_791_v68,_dafny.Seq.Concat(_790_v67, _790_v67));
      _792_v69 = (_792_v69).update(_791_v68, _790_v67);
      let _793_v70;
      _793_v70 = _dafny.Seq.of(p2, p2, p2);
      let _794_v71;
      _794_v71 = _dafny.Seq.of(_790_v67, _790_v67);
      let _795_v73;
      _795_v73 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _796_v74;
      _796_v74 = _dafny.Seq.of(new BigNumber((_795_v73).length));
      let _797_v75;
      _797_v75 = _dafny.Map.Empty.slice().updateUnsafe(_796_v74,true);
      let _798_v76;
      let _nw148 = Array((new BigNumber(19)).toNumber());
      _nw148[(_dafny.ZERO).toNumber()] = (_793_v70)[_module.__default.safeIndex(p1, new BigNumber((_793_v70).length))];
      _nw148[(_dafny.ONE).toNumber()] = p2;
      _nw148[(new BigNumber(2)).toNumber()] = p2;
      _nw148[(new BigNumber(3)).toNumber()] = p2;
      _nw148[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsPrefixOf((_794_v71)[_module.__default.safeIndex(_this.f2, new BigNumber((_794_v71).length))], _790_v67);
      _nw148[(new BigNumber(5)).toNumber()] = (_this).fm8(_790_v67, p2, globalState);
      _nw148[(new BigNumber(6)).toNumber()] = (_this).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-707)), ((_799_v68) => function (_800_i6) {
        return _799_v68;
      })(_791_v68)), p2, globalState);
      _nw148[(new BigNumber(7)).toNumber()] = p2;
      _nw148[(new BigNumber(8)).toNumber()] = p2;
      _nw148[(new BigNumber(9)).toNumber()] = !(_this.f2).isEqualTo(_this.f2);
      _nw148[(new BigNumber(10)).toNumber()] = p2;
      _nw148[(new BigNumber(11)).toNumber()] = (p1).isLessThanOrEqualTo(new BigNumber((function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of (_797_v75).Keys.Elements) {
          let _801_v72 = _compr_25;
          if ((_797_v75).contains(_801_v72)) {
            _coll25.push([_801_v72,p2]);
          }
        }
        return _coll25;
      }()).length));
      _nw148[(new BigNumber(12)).toNumber()] = p2;
      _nw148[(new BigNumber(13)).toNumber()] = (_793_v70)[_module.__default.safeIndex(_this.f2, new BigNumber((_793_v70).length))];
      _nw148[(new BigNumber(14)).toNumber()] = (new BigNumber(684)).isEqualTo(_this.f2);
      _nw148[(new BigNumber(15)).toNumber()] = (p2) || (p2);
      _nw148[(new BigNumber(16)).toNumber()] = p2;
      _nw148[(new BigNumber(17)).toNumber()] = p2;
      _nw148[(new BigNumber(18)).toNumber()] = (p2) === (p2);
      _798_v76 = _nw148;
      let _index128 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
      (_798_v76)[_index128] = p2;
      let _index129 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
      (_798_v76)[_index129] = p2;
      let _802_i7;
      _802_i7 = _dafny.ZERO;
      L1: {
        while (!(true) || (false)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_802_i7)) {
              break L1;
            }
            _802_i7 = (_802_i7).plus(_dafny.ONE);
            let _803_v77;
            _803_v77 = _dafny.Map.Empty.slice().updateUnsafe((((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]) ? (_791_v68) : (_791_v68)),_dafny.Map.Empty.slice().updateUnsafe(_this.f2,p0));
            let _804_v78;
            _804_v78 = _dafny.MultiSet.fromElements(p2);
            let _805_v79;
            _805_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,new BigNumber((_804_v78).cardinality()));
            _803_v77 = (_803_v77).update(_791_v68, _805_v79);
            if (((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]) && ((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))])) {
              r0 = _module.__default.safeModuloInt(_this.f2, p1);
              r0 = (_dafny.ZERO).minus(p0);
              let _index130 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              (_798_v76)[_index130] = p2;
              let _806_v81;
              _806_v81 = _dafny.Set.fromElements(p1, _this.f2);
              let _index131 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              (_798_v76)[_index131] = ((function () {
                let _coll26 = new _dafny.Set();
                for (const _compr_26 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(-727))) {
                  let _807_v80 = _compr_26;
                  if (((_dafny.ZERO).isLessThanOrEqualTo(_807_v80)) && ((_807_v80).isLessThan(new BigNumber(-727)))) {
                    _coll26.add((_807_v80).plus(p0));
                  }
                }
                return _coll26;
              }()).Union(_806_v81)).IsSubsetOf(_module.__default.fm31(!((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]), p0, p2, p0, globalState));
              r0 = new BigNumber(-962);
            } else {
              _804_v78 = _804_v78;
              let _808_v82;
              _808_v82 = _dafny.MultiSet.fromElements(_790_v67);
              let _809_v83;
              _809_v83 = _module.D0.create_DC1((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))], (_808_v82).IsDisjointFrom(_808_v82), _791_v68, _this.f2);
              let _index132 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              let _index133 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              let _rhs116 = _809_v83;
              let _rhs117 = !(!((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]));
              let _rhs118 = new _dafny.CodePoint('a'.codePointAt(0));
              let _rhs119 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_790_v67, _790_v67), _790_v67);
              let _lhs56 = _798_v76;
              let _lhs57 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              let _lhs58 = _798_v76;
              let _lhs59 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              _809_v83 = _rhs116;
              _lhs56[_lhs57] = _rhs117;
              _791_v68 = _rhs118;
              _lhs58[_lhs59] = _rhs119;
              let _810_v84;
              _810_v84 = _module.D7.create_DC14(_796_v74);
              let _811_v85;
              _811_v85 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.FromArray(_module.__default.fm32(globalState))).Intersect(_dafny.MultiSet.fromElements(_810_v84)),new BigNumber((_790_v67).length));
              let _812_v86;
              _812_v86 = _dafny.Map.Empty.slice().updateUnsafe((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))],p1);
              let _813_v87;
              _813_v87 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(p0, new BigNumber(264), new BigNumber((_790_v67).length), (((_805_v79).contains((_this).fm1(p2, globalState))) ? ((_805_v79).get((_this).fm1(p2, globalState))) : (new BigNumber((_812_v86).length))), new BigNumber(462))).length));
              let _814_v88;
              let _nw149 = new _module.C1();
              _nw149.__ctor(_790_v67);
              _814_v88 = _nw149;
              let _815_v89;
              _815_v89 = _dafny.Map.Empty.slice().updateUnsafe(p2,_814_v88);
              let _816_v90;
              _816_v90 = _module.D8.create_DC16(_814_v88);
              _811_v85 = (_811_v85).update(_module.__default.fm33(_791_v68, true, new BigNumber((_813_v87).length), p1, globalState), new BigNumber((_dafny.Seq.of((_815_v89).update(p2, (_816_v90).dtor_cf29), _815_v89, _815_v89, (_815_v89).update((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))], _814_v88), _815_v89)).length));
              r0 = _this.f2;
              let _index134 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              (_798_v76)[_index134] = (_809_v83).dtor_cf1;
            }
            if (p2) {
              let _817_v91;
              _817_v91 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]);
              let _818_v92;
              _818_v92 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_791_v68);
              let _819_v93;
              _819_v93 = _dafny.MultiSet.fromElements(p0, new BigNumber(((_818_v92).update(_this.f2, _791_v68)).length), _this.f2);
              _817_v91 = (_817_v91).update((((_819_v93).contains(_this.f2)) ? ((_819_v93).get(_this.f2)) : (p1)), (_819_v93).IsProperSubsetOf(_819_v93));
              let _820_v94;
              _820_v94 = _dafny.Map.Empty.slice().updateUnsafe((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))],_798_v76);
              let _821_v95;
              _821_v95 = _dafny.Set.fromElements((new BigNumber((_dafny.Seq.UnicodeFromString("ssamx")).length)).plus(new BigNumber((_820_v94).length)));
              _821_v95 = (((p2) ? (_821_v95) : (function () {
                let _coll27 = new _dafny.Set();
                for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-95), new BigNumber(219))) {
                  let _822_v96 = _compr_27;
                  if (((new BigNumber(-95)).isLessThanOrEqualTo(_822_v96)) && ((_822_v96).isLessThan(new BigNumber(219)))) {
                    _coll27.add((_822_v96).multipliedBy(new BigNumber(685)));
                  }
                }
                return _coll27;
              }()))).Union((_dafny.Set.fromElements(p0, p0)).Union(_821_v95));
              (_this).f2 = _this.f2;
              (_this).f2 = _this.f2;
              let _rhs120 = _790_v67;
              let _rhs121 = p1;
              let _rhs122 = (_this).fm1((false) && ((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]), globalState);
              _790_v67 = _rhs120;
              r0 = _rhs121;
              r0 = _rhs122;
            } else {
              let _823_v97;
              let _nw150 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
              _823_v97 = _nw150;
              let _824_v98;
              let _nw151 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
              _824_v98 = _nw151;
              let _825_v99;
              _825_v99 = _dafny.Seq.of(_824_v98, _824_v98);
              let _index135 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_823_v97).length));
              (_823_v97)[_index135] = _825_v99;
              let _index136 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_823_v97).length));
              (_823_v97)[_index136] = _825_v99;
              let _826_v100;
              _826_v100 = _module.D6.create_DC13((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))], p1);
              let _index137 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
              (_798_v76)[_index137] = (((false) ? (_826_v100) : (_826_v100))).dtor_cf24;
              let _827_v101;
              _827_v101 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p0, p1),_824_v98);
              let _828_v102;
              _828_v102 = _dafny.Map.Empty.slice().updateUnsafe((_793_v70)[_module.__default.safeIndex(_this.f2, new BigNumber((_793_v70).length))],_827_v101);
              _827_v101 = (((_828_v102).contains((_this).fm8(_790_v67, p2, globalState))) ? ((_828_v102).get((_this).fm8(_790_v67, p2, globalState))) : (_827_v101));
              _790_v67 = _dafny.Seq.UnicodeFromString("weab");
              let _829_v103;
              let _nw152 = Array((new BigNumber(24)).toNumber()).fill([]);
              _829_v103 = _nw152;
              let _830_v104;
              _830_v104 = _dafny.Map.Empty.slice().updateUnsafe((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))],_824_v98);
              let _nw153 = Array((new BigNumber(23)).toNumber());
              _nw153[(_dafny.ZERO).toNumber()] = _824_v98;
              _nw153[(_dafny.ONE).toNumber()] = _824_v98;
              _nw153[(new BigNumber(2)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(3)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(4)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(5)).toNumber()] = (((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]) ? (_824_v98) : (_824_v98));
              _nw153[(new BigNumber(6)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(7)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(8)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(9)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(10)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(11)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(12)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(13)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(14)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(15)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(16)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(17)).toNumber()] = (((_830_v104).contains((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))])) ? ((_830_v104).get((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))])) : (_824_v98));
              _nw153[(new BigNumber(18)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(19)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(20)).toNumber()] = (((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]) ? (_824_v98) : (_824_v98));
              _nw153[(new BigNumber(21)).toNumber()] = _824_v98;
              _nw153[(new BigNumber(22)).toNumber()] = _824_v98;
              _829_v103 = _nw153;
            }
            r0 = p1;
          }
        }
      }
      let _831_v105;
      let _nw154 = new _module.C1();
      _nw154.__ctor(_790_v67);
      _831_v105 = _nw154;
      let _832_v106;
      _832_v106 = _module.D8.create_DC16(_831_v105);
      let _source16 = _832_v106;
      if (_source16.is_DC17) {
        let _833___mcc_h8 = (_source16).cf30;
        let _834___mcc_h9 = (_source16).cf31;
        let _835___mcc_h10 = (_source16).cf32;
        let _836_cf32 = _835___mcc_h10;
        let _837_cf31 = _834___mcc_h9;
        let _838_cf30 = _833___mcc_h8;
        let _839_v107;
        _839_v107 = _795_v73;
        let _840_v108;
        _840_v108 = _dafny.Map.Empty.slice().updateUnsafe(_839_v107,p2);
        let _841_v109;
        _841_v109 = _dafny.Set.fromElements((_840_v108).Merge((_840_v108).update(_839_v107, !(p2))), (_840_v108).update(_839_v107, _838_cf30));
        let _842_v110;
        _842_v110 = _dafny.Map.Empty.slice().updateUnsafe(_840_v108,p1);
        _841_v109 = (function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of (_842_v110).Keys.Elements) {
            let _843_v111 = _compr_28;
            if ((_842_v110).contains(_843_v111)) {
              _coll28.add(_843_v111);
            }
          }
          return _coll28;
        }()).Difference(_dafny.Set.fromElements(_840_v108, _840_v108));
        let _844_v112;
        let _init12 = ((_845_v68) => function (_846_i8) {
          return _module.D7.create_DC15(_this.f2, _845_v68);
        })(_791_v68);
        let _nw155 = Array((new BigNumber(17)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw155.length); _i0_12++) {
          _nw155[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _844_v112 = _nw155;
        _844_v112 = ((p2) ? (_844_v112) : (_844_v112));
        _838_cf30 = _838_cf30;
        let _847_v113;
        let _nw156 = new _module.C1();
        _nw156.__ctor(_790_v67);
        _847_v113 = _nw156;
      } else {
        let _848___mcc_h11 = (_source16).cf29;
        let _849_cf29 = _848___mcc_h11;
        let _index138 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        (_798_v76)[_index138] = ((p1).isLessThan(p1)) && (!((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]));
        let _index139 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        (_798_v76)[_index139] = p2;
        let _850_v114;
        _850_v114 = _dafny.Set.fromElements(_this.f2);
        let _851_v115;
        let _nw157 = Array((new BigNumber(2)).toNumber());
        _nw157[(_dafny.ZERO).toNumber()] = new BigNumber(600);
        _nw157[(_dafny.ONE).toNumber()] = new BigNumber((_850_v114).length);
        _851_v115 = _nw157;
        let _852_v116;
        _852_v116 = _module.D4.create_DC6(p1, _851_v115, p1);
        let _source17 = _852_v116;
        if (_source17.is_DC6) {
          let _853___mcc_h12 = (_source17).cf9;
          let _854___mcc_h13 = (_source17).cf10;
          let _855___mcc_h14 = (_source17).cf11;
          let _856_cf11 = _855___mcc_h14;
          let _857_cf10 = _854___mcc_h13;
          let _858_cf9 = _853___mcc_h12;
          let _index140 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
          (_798_v76)[_index140] = (_this).fm8(_dafny.Seq.UnicodeFromString("yjjaoweep"), (_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))], globalState);
          let _859_v117;
          _859_v117 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))]);
          let _860_v118;
          let _nw158 = new _module.C0();
          _nw158.__ctor(p2, (((_859_v117).contains(_this.f2)) ? ((_859_v117).get(_this.f2)) : ((_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))])));
          _860_v118 = _nw158;
          let _861_v119;
          _861_v119 = _dafny.Map.Empty.slice().updateUnsafe(_860_v118,_795_v73);
          let _862_v120;
          _862_v120 = _dafny.Map.Empty.slice().updateUnsafe(_795_v73,_796_v74);
          _858_cf9 = (_dafny.ZERO).minus((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((((_861_v119).contains(_860_v118)) ? ((_861_v119).get(_860_v118)) : (_795_v73)),_796_v74)).Merge(_862_v120)).length)).minus((_dafny.ZERO).minus(_this.f2)));
          _859_v117 = (_859_v117).update(_module.__default.safeDivisionInt(p0, p0), _860_v118.f5);
          let _863_v121;
          _863_v121 = _dafny.Set.fromElements(_860_v118.f5, (_831_v105).fm11(globalState));
          let _864_v122;
          _864_v122 = _dafny.Map.Empty.slice().updateUnsafe((_860_v118).fm14((_831_v105).fm11(globalState), p2, _850_v114, globalState),_module.D6.create_DC11(_module.__default.fm28(true, p2, _860_v118.f4, _863_v121, globalState)));
          let _865_v124;
          _865_v124 = _module.D6.create_DC11((_831_v105).f3);
          _864_v122 = (_864_v122).update((function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of _dafny.IntegerRange(new BigNumber(461), new BigNumber(893))) {
              let _866_v123 = _compr_29;
              if (((new BigNumber(461)).isLessThanOrEqualTo(_866_v123)) && ((_866_v123).isLessThan(new BigNumber(893)))) {
                _coll29.add((_866_v123).minus(p0));
              }
            }
            return _coll29;
          }()).IsProperSubsetOf(_850_v114), _865_v124);
        } else if (_source17.is_DC5) {
          let _867___mcc_h15 = (_source17).cf8;
          let _868_cf8 = _867___mcc_h15;
          let _index141 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
          (_798_v76)[_index141] = p2;
          let _nw159 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _851_v115 = _nw159;
          (_this).f2 = new BigNumber((_dafny.Seq.UnicodeFromString("qnoyfxqrs")).length);
          let _869_v125;
          _869_v125 = _dafny.MultiSet.fromElements(_851_v115, _851_v115);
          let _870_v126;
          _870_v126 = _module.D5.create_DC8(_869_v125);
          let _871_v127;
          _871_v127 = _dafny.Map.Empty.slice().updateUnsafe(p1,_870_v126);
          let _872_v128;
          _872_v128 = _dafny.Seq.of(_871_v127);
          r0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_871_v127), _dafny.Seq.update(_872_v128, _module.__default.safeIndex(p1, new BigNumber((_872_v128).length)), _871_v127))).length);
        } else {
          let _873___mcc_h16 = (_source17).cf12;
          let _874_cf12 = _873___mcc_h16;
          let _875_v129;
          _875_v129 = _dafny.Map.Empty.slice().updateUnsafe(_791_v68,(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,p0)).length)).plus(_this.f2));
          _875_v129 = (_875_v129).update(_791_v68, _this.f2);
          let _index142 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
          (_798_v76)[_index142] = (_831_v105).fm11(globalState);
          (_this).f2 = p0;
          _849_cf29 = _849_cf29;
        }
        let _876_v130;
        _876_v130 = _dafny.Map.Empty.slice().updateUnsafe((_831_v105).f3,p0);
        let _index143 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        let _index144 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        let _rhs123 = _832_v106;
        let _rhs124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yac"),_this.f2);
        let _rhs125 = (p0).isEqualTo(p1);
        let _rhs126 = _851_v115;
        let _rhs127 = (_798_v76)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length))];
        let _lhs60 = _798_v76;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        let _lhs62 = _798_v76;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_798_v76).length));
        _832_v106 = _rhs123;
        _876_v130 = _rhs124;
        _lhs60[_lhs61] = _rhs125;
        _851_v115 = _rhs126;
        _lhs62[_lhs63] = _rhs127;
      }
      let _877_v131;
      _877_v131 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), ((_878_v68) => function (_879_i9) {
        return _878_v68;
      })(_791_v68))).length), _this.f2));
      let _880_v132;
      let _init13 = ((_881_p0) => function (_882_i11) {
        return _module.__default.safeDivisionInt(_882_i11, _881_p0);
      })(p0);
      let _nw160 = Array((new BigNumber(21)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw160.length); _i0_13++) {
        _nw160[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _880_v132 = _nw160;
      r0 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_877_v131).contains(p2)) ? ((_877_v131).get(p2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(482)), function (_883_i10) {
        return new BigNumber(73);
      })))).length),_880_v132)).length)), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f2)).cardinality()),(_831_v105).fm11(globalState))).length))).length));
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      if (p3) {
        let _884_v0;
        _884_v0 = true;
        _884_v0 = _884_v0;
        let _885_v1;
        _885_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f2);
        let _886_v2;
        _886_v2 = _dafny.Seq.of(p3);
        _885_v1 = (_885_v1).update(p1, (_this.f2).minus(new BigNumber((_dafny.MultiSet.FromArray(_886_v2)).cardinality())));
        let _887_v3;
        _887_v3 = new _dafny.CodePoint('m'.codePointAt(0));
        let _888_v4;
        _888_v4 = _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), _887_v3);
        let _889_v5;
        _889_v5 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_888_v4)[_module.__default.safeIndex(_this.f2, new BigNumber((_888_v4).length))]);
        let _890_v6;
        _890_v6 = _dafny.Set.fromElements(_889_v5, _889_v5);
        let _891_v7;
        _891_v7 = _dafny.Set.fromElements(_884_v0, p3);
        let _892_v8;
        _892_v8 = _dafny.Map.Empty.slice().updateUnsafe(_889_v5,_891_v7);
        let _893_v9;
        _893_v9 = _dafny.Map.Empty.slice().updateUnsafe(_890_v6,_892_v8);
        _893_v9 = (_893_v9).update((_890_v6).Difference(_dafny.Set.fromElements(_889_v5, _889_v5, _dafny.Map.Empty.slice().updateUnsafe(false,_887_v3), _889_v5, _889_v5)), _dafny.Map.Empty.slice().updateUnsafe(_889_v5,_891_v7));
        let _894_v10;
        _894_v10 = _dafny.Set.fromElements(_this.f2, _this.f2);
        if ((_894_v10).IsSubsetOf(_894_v10)) {
          _886_v2 = _dafny.Seq.of(p2, _884_v0, p3, p3, !(p3));
          let _895_v11;
          let _nw161 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _895_v11 = _nw161;
          let _896_v12;
          _896_v12 = _dafny.Seq.of(_888_v4, _888_v4, _888_v4);
          let _index145 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_895_v11).length));
          (_895_v11)[_index145] = _dafny.Seq.Concat((_896_v12)[_module.__default.safeIndex(_this.f2, new BigNumber((_896_v12).length))], _888_v4);
          let _897_v13;
          let _nw162 = Array((new BigNumber(26)).toNumber()).fill(false);
          _897_v13 = _nw162;
          let _898_v14;
          _898_v14 = _dafny.MultiSet.fromElements(p2);
          let _899_v15;
          _899_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm1(_884_v0, globalState),new BigNumber((_898_v14).cardinality()));
          let _900_v16;
          _900_v16 = _dafny.MultiSet.fromElements(_899_v15, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p2, _884_v0)).length),_this.f2));
          let _index146 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length));
          (_897_v13)[_index146] = (_900_v16).IsProperSubsetOf(_900_v16);
          let _901_v17;
          _901_v17 = _dafny.Seq.of((((_898_v14).contains(p2)) ? ((_898_v14).get(p2)) : (_this.f2)), _this.f2, _this.f2, new BigNumber(256), _this.f2);
          let _index147 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_895_v11).length));
          let _index148 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length));
          let _rhs128 = _dafny.Seq.IsPrefixOf(_901_v17, _901_v17);
          let _rhs129 = _888_v4;
          let _rhs130 = (_this.f2).minus((_dafny.ZERO).minus(_this.f2));
          let _rhs131 = (_this).fm1(((true) ? (false) : (p3)), globalState);
          let _rhs132 = _884_v0;
          let _lhs64 = _895_v11;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_895_v11).length));
          let _lhs66 = _this;
          let _lhs67 = _897_v13;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length));
          _884_v0 = _rhs128;
          _lhs64[_lhs65] = _rhs129;
          _lhs66.f2 = _rhs130;
          r0 = _rhs131;
          _lhs67[_lhs68] = _rhs132;
          let _902_v18;
          let _nw163 = new _module.C0();
          _nw163.__ctor((_this).fm8(_module.__default.fm28((_897_v13)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length))], p3, p2, _891_v7, globalState), p2, globalState), (_886_v2)[_module.__default.safeIndex(_this.f2, new BigNumber((_886_v2).length))]);
          _902_v18 = _nw163;
          _897_v13 = _897_v13;
          let _index149 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length));
          (_897_v13)[_index149] = (((_897_v13)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length))]) ? ((_897_v13)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_897_v13).length))]) : ((_this.f2).isLessThan(_this.f2)));
        } else {
          let _903_v19;
          let _nw164 = Array((new BigNumber(26)).toNumber());
          _903_v19 = _nw164;
          let _904_v20;
          let _nw165 = new _module.C0();
          _nw165.__ctor(p2, (_886_v2)[_module.__default.safeIndex(new BigNumber((_888_v4).length), new BigNumber((_886_v2).length))]);
          _904_v20 = _nw165;
          let _index150 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_903_v19).length));
          (_903_v19)[_index150] = _904_v20;
          let _index151 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_903_v19).length));
          (_903_v19)[_index151] = ((_904_v20.f4) ? (_904_v20) : (_904_v20));
          r1 = (_dafny.ZERO).minus(_this.f2);
          let _905_v22;
          _905_v22 = _module.D5.create_DC10((_904_v20).fm14(_884_v0, _904_v20.f4, _894_v10, globalState), _this.f2, _dafny.Map.Empty.slice().updateUnsafe(_888_v4,function () {
  let _coll30 = new _dafny.Map();
  for (const _compr_30 of _dafny.IntegerRange(new BigNumber(338), new BigNumber(202))) {
    let _906_v21 = _compr_30;
    if (((new BigNumber(338)).isLessThanOrEqualTo(_906_v21)) && ((_906_v21).isLessThan(new BigNumber(202)))) {
      _coll30.push([_module.__default.safeModuloInt(_906_v21, (((p0).contains(_this.f2)) ? ((p0).get(_this.f2)) : (new BigNumber(198)))),_904_v20.f5]);
    }
  }
  return _coll30;
}()), p2);
          _905_v22 = _905_v22;
          let _907_v23;
          let _nw166 = new _module.C0();
          _nw166.__ctor(_884_v0, _884_v0);
          _907_v23 = _nw166;
          (_907_v23).f4 = (_907_v23).fm14(_907_v23.f5, (_dafny.MultiSet.fromElements(p2, _884_v0)).IsSubsetOf(_dafny.MultiSet.FromArray(_886_v2)), (_894_v10).Difference(_894_v10), globalState);
        }
        let _908_v24;
        _908_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f2),_884_v0);
        let _909_v25;
        _909_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hrjp"),_908_v24);
        let _910_v26;
        _910_v26 = _module.D5.create_DC10(!(p2), _this.f2, _909_v25, p2);
        _884_v0 = ((p2) ? (((_910_v26).dtor_cf18).isLessThanOrEqualTo(new BigNumber(876))) : (p2));
      } else {
        let _911_v27;
        _911_v27 = true;
        _911_v27 = !(p3) || (_911_v27);
        let _912_v28;
        _912_v28 = new _dafny.CodePoint('v'.codePointAt(0));
        let _913_v29;
        _913_v29 = _dafny.MultiSet.fromElements(_912_v28, _912_v28);
        let _914_v30;
        _914_v30 = _dafny.Seq.UnicodeFromString("h");
        let _915_v31;
        let _nw167 = Array((new BigNumber(7)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("gmqami");
        _nw167[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_914_v30, _914_v30);
        _nw167[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("tuqtafex");
        _nw167[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("g"), _914_v30);
        _nw167[(new BigNumber(4)).toNumber()] = _914_v30;
        _nw167[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_914_v30, _914_v30);
        _nw167[(new BigNumber(6)).toNumber()] = _914_v30;
        _915_v31 = _nw167;
        let _index152 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_915_v31).length));
        (_915_v31)[_index152] = _914_v30;
        let _916_v32;
        let _init14 = ((_917_p3) => function (_918_i0) {
          return _917_p3;
        })(p3);
        let _nw168 = Array((new BigNumber(20)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw168.length); _i0_14++) {
          _nw168[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _916_v32 = _nw168;
        let _919_v33;
        _919_v33 = _dafny.Map.Empty.slice().updateUnsafe(_916_v32,_dafny.MultiSet.fromElements(_this.f2, _this.f2));
        let _920_v34;
        _920_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_913_v29).Intersect(_913_v29));
        let _index153 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_915_v31).length));
        let _rhs133 = (((_920_v34).contains(_this.f2)) ? ((_920_v34).get(_this.f2)) : ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), _912_v28)).update(new _dafny.CodePoint('d'.codePointAt(0)), _module.__default.abs(_this.f2))));
        let _rhs134 = p2;
        let _rhs135 = _914_v30;
        let _rhs136 = _919_v33;
        let _rhs137 = (_this).fm1(false, globalState);
        let _lhs69 = _915_v31;
        let _lhs70 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_915_v31).length));
        let _lhs71 = _this;
        _913_v29 = _rhs133;
        _911_v27 = _rhs134;
        _lhs69[_lhs70] = _rhs135;
        _919_v33 = _rhs136;
        _lhs71.f2 = _rhs137;
        let _921_v35;
        _921_v35 = _dafny.Set.fromElements(_this.f2);
        let _922_v36;
        _922_v36 = _dafny.MultiSet.fromElements(!(_911_v27));
        let _923_v37;
        _923_v37 = _dafny.Seq.of(_922_v36, _922_v36, _922_v36, _922_v36);
        let _924_v38;
        _924_v38 = _dafny.Map.Empty.slice().updateUnsafe(_921_v35,(_923_v37)[_module.__default.safeIndex(_this.f2, new BigNumber((_923_v37).length))]);
        _924_v38 = (_924_v38).update(_921_v35, _dafny.MultiSet.fromElements((_this).fm8((_915_v31)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_915_v31).length))], p3, globalState)));
        let _925_v39;
        let _nw169 = Array((new BigNumber(21)).toNumber()).fill([]);
        _925_v39 = _nw169;
        let _926_v40;
        let _nw170 = Array((new BigNumber(3)).toNumber()).fill([]);
        _926_v40 = _nw170;
        let _index154 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_925_v39).length));
        (_925_v39)[_index154] = _926_v40;
        let _index155 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_916_v32).length));
        (_916_v32)[_index155] = p3;
        let _index156 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((p1).length));
        (p1)[_index156] = _this.f2;
        let _927_v41;
        _927_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8((_915_v31)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_915_v31).length))], _911_v27, globalState),_926_v40);
        let _index157 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_925_v39).length));
        let _index158 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_916_v32).length));
        let _index159 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((p1).length));
        let _rhs138 = (_this.f2).minus(_this.f2);
        let _rhs139 = (((_927_v41).contains(_911_v27)) ? ((_927_v41).get(_911_v27)) : (_926_v40));
        let _rhs140 = _911_v27;
        let _rhs141 = _this.f2;
        let _lhs72 = _this;
        let _lhs73 = _925_v39;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_925_v39).length));
        let _lhs75 = _916_v32;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_916_v32).length));
        let _lhs77 = p1;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((p1).length));
        _lhs72.f2 = _rhs138;
        _lhs73[_lhs74] = _rhs139;
        _lhs75[_lhs76] = _rhs140;
        _lhs77[_lhs78] = _rhs141;
        let _928_v42;
        let _nw171 = new _module.C1();
        _nw171.__ctor(_914_v30);
        _928_v42 = _nw171;
        let _929_v43;
        _929_v43 = _dafny.Map.Empty.slice().updateUnsafe(_912_v28,_928_v42);
        _929_v43 = (_929_v43).update(_912_v28, _928_v42);
      }
      let _930_v44;
      let _nw172 = Array((new BigNumber(7)).toNumber()).fill(false);
      _930_v44 = _nw172;
      let _index160 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      (_930_v44)[_index160] = !(p3);
      let _931_v45;
      _931_v45 = _module.D6.create_DC13(p2, _this.f2);
      let _932_v46;
      _932_v46 = _dafny.Seq.UnicodeFromString("hgoffmr");
      let _index161 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      let _rhs142 = ((!((p3) || ((_931_v45).dtor_cf24))) ? (_dafny.Seq.IsPrefixOf(_932_v46, _dafny.Seq.UnicodeFromString("rabeglte"))) : (p2));
      let _rhs143 = _this.f2;
      let _lhs79 = _930_v44;
      let _lhs80 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      _lhs79[_lhs80] = _rhs142;
      r1 = _rhs143;
      let _933_v47;
      _933_v47 = _dafny.Seq.of(p3);
      let _934_v48;
      _934_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_933_v47);
      let _935_i1;
      _935_i1 = _dafny.ZERO;
      L2: {
        while (!_dafny.areEqual((((_934_v48).contains(_this.f2)) ? ((_934_v48).get(_this.f2)) : (_933_v47)), _dafny.Seq.Concat(_dafny.Seq.of(p3, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]), _933_v47))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_935_i1)) {
              break L2;
            }
            _935_i1 = (_935_i1).plus(_dafny.ONE);
            let _936_v49;
            _936_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_this.f2);
            let _index162 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
            (p1)[_index162] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_936_v49,(_dafny.ZERO).minus(_this.f2))).length);
            let _937_v50;
            _937_v50 = new _dafny.CodePoint('v'.codePointAt(0));
            let _938_v51;
            _938_v51 = _dafny.Map.Empty.slice().updateUnsafe(_937_v50,_this.f2);
            let _index163 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
            (p1)[_index163] = new BigNumber((_938_v51).length);
            if (((!(p3)) ? ((_module.__default.fm34(p3, globalState)).contains(_dafny.Seq.UnicodeFromString("tfbw"))) : (p2))) {
              let _index164 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
              (_930_v44)[_index164] = (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))];
              r0 = ((p1)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length))]).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length))])));
              let _939_v52;
              let _nw173 = Array((new BigNumber(25)).toNumber()).fill([]);
              _939_v52 = _nw173;
              let _940_v53;
              _940_v53 = _dafny.MultiSet.fromElements(_937_v50, _937_v50, _937_v50, _937_v50, _937_v50);
              let _index165 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
              let _index166 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
              let _rhs144 = ((!(p2)) ? (_939_v52) : (((p3) ? (_939_v52) : (_939_v52))));
              let _rhs145 = _dafny.MultiSet.FromArray(_932_v46);
              let _rhs146 = (_this).fm1(p3, globalState);
              let _rhs147 = _this.f2;
              let _lhs81 = p1;
              let _lhs82 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
              let _lhs83 = p1;
              let _lhs84 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length));
              _939_v52 = _rhs144;
              _940_v53 = _rhs145;
              _lhs81[_lhs82] = _rhs146;
              _lhs83[_lhs84] = _rhs147;
              r0 = (p1)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length))];
              let _941_v54;
              _941_v54 = _dafny.Map.Empty.slice().updateUnsafe((_936_v49).Merge(_936_v49),_933_v47);
              _941_v54 = (_941_v54).update(_936_v49, _dafny.Seq.Concat(_933_v47, _dafny.Seq.of(p3)));
            } else {
              _930_v44 = _930_v44;
              let _942_v55;
              let _nw174 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
              _942_v55 = _nw174;
              let _943_v56;
              _943_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_this.f2, new BigNumber((p0).cardinality())),((p2) ? (_942_v55) : (_942_v55)));
              let _944_v57;
              _944_v57 = _dafny.Map.Empty.slice().updateUnsafe(p3,_932_v46);
              let _945_v58;
              _945_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(170)), ((_946_v50) => function (_947_i2) {
                return _946_v50;
              })(_937_v50)), p2, globalState),new BigNumber(((((_944_v57).contains((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))])) ? ((_944_v57).get((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))])) : (_932_v46))).length));
              _942_v55 = (((_943_v56).contains(new BigNumber(((_945_v58).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_module.__default.fm28(p2, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], p3, _dafny.Set.fromElements(p2, true), globalState), (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], globalState),new BigNumber(766)))).length))) ? ((_943_v56).get(new BigNumber(((_945_v58).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_module.__default.fm28(p2, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], p3, _dafny.Set.fromElements(p2, true), globalState), (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], globalState),new BigNumber(766)))).length))) : (((p3) ? (_942_v55) : (_942_v55))));
              let _948_v59;
              let _nw175 = new _module.C0();
              _nw175.__ctor(((p1)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length))]).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), ((_949_v50) => function (_950_i3) {
                return _949_v50;
              })(_937_v50))).length)), (_this).fm8(_dafny.Seq.update(_932_v46, _module.__default.safeIndex(_this.f2, new BigNumber((_932_v46).length)), _937_v50), false, globalState));
              _948_v59 = _nw175;
              (_948_v59).f4 = (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))];
              let _951_v60;
              _951_v60 = _dafny.Map.Empty.slice().updateUnsafe(p2,_948_v59.f5);
              (_948_v59).f5 = (((_951_v60).contains(!((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]))) ? ((_951_v60).get(!((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]))) : (p3));
            }
            let _952_v61;
            let _nw176 = new _module.C1();
            _nw176.__ctor(_dafny.Seq.update((((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]) ? (_932_v46) : (_dafny.Seq.UnicodeFromString("dvbp"))), _module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((p1).length))], new BigNumber(((((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]) ? (_932_v46) : (_dafny.Seq.UnicodeFromString("dvbp")))).length)), _937_v50));
            _952_v61 = _nw176;
            let _953_v62;
            let _nw177 = new _module.C0();
            _nw177.__ctor((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], (new BigNumber(940)).isLessThanOrEqualTo(new BigNumber(-802)));
            _953_v62 = _nw177;
          }
        }
      }
      let _954_v63;
      _954_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this.f2).plus(_this.f2),_this.f2);
      let _955_v64;
      _955_v64 = _dafny.MultiSet.fromElements((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], p2);
      let _index167 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length));
      (p1)[_index167] = (new BigNumber(125)).plus(new BigNumber((_955_v64).cardinality()));
      let _956_v65;
      _956_v65 = _module.D7.create_DC15(new BigNumber((_dafny.Set.fromElements((_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], true)).length), new _dafny.CodePoint('n'.codePointAt(0)));
      let _957_v66;
      _957_v66 = new _dafny.CodePoint('n'.codePointAt(0));
      let _958_v67;
      _958_v67 = _module.D0.create_DC1(!(p2), (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], _957_v66, _this.f2);
      let _pat_let_tv18 = _930_v44;
      let _pat_let_tv19 = _930_v44;
      let _index168 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length));
      let _index169 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      let _rhs148 = _module.__default.fm28((_933_v47)[_module.__default.safeIndex(_this.f2, new BigNumber((_933_v47).length))], true, p3, _dafny.Set.fromElements(false, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))]), globalState);
      let _rhs149 = _954_v63;
      let _rhs150 = (_this.f2).plus(_this.f2);
      let _rhs151 = (_956_v65).dtor_cf27;
      let _rhs152 = function (_source18) {
        if (_source18.is_DC1) {
          let _959___mcc_h0 = (_source18).cf1;
          let _960___mcc_h1 = (_source18).cf2;
          let _961___mcc_h2 = (_source18).cf3;
          let _962___mcc_h3 = (_source18).cf4;
          let _963_cf4 = _962___mcc_h3;
          let _964_cf3 = _961___mcc_h2;
          let _965_cf2 = _960___mcc_h1;
          let _966_cf1 = _959___mcc_h0;
          return (_pat_let_tv19)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_pat_let_tv18).length))];
        } else {
          let _967___mcc_h4 = (_source18).cf0;
          let _968_cf0 = _967___mcc_h4;
          return (new BigNumber(172)).isLessThanOrEqualTo(_this.f2);
        }
      }(_958_v67);
      let _lhs85 = _this;
      let _lhs86 = p1;
      let _lhs87 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length));
      let _lhs88 = _930_v44;
      let _lhs89 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      _932_v46 = _rhs148;
      _954_v63 = _rhs149;
      _lhs85.f2 = _rhs150;
      _lhs86[_lhs87] = _rhs151;
      _lhs88[_lhs89] = _rhs152;
      let _969_v68;
      _969_v68 = _dafny.Set.fromElements(p2, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], true);
      let _hi3 = (p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))];
      for (let _970_i4 = ((p2) ? ((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))]) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_969_v68).length),(p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))])).length))); _970_i4.isLessThan(_hi3); _970_i4 = _970_i4.plus(_dafny.ONE)) {
        let _pat_let_tv20 = _932_v46;
        let _pat_let_tv21 = p2;
        let _971_v69;
        _971_v69 = _dafny.Set.fromElements(_957_v66, (function (_pat_let16_0) {
          return function (_972_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_973_dt__update_hcf4_h0) {
                return function (_pat_let18_0) {
                  return function (_974_dt__update_hcf2_h0) {
                    return _module.D0.create_DC1((_972_dt__update__tmp_h0).dtor_cf1, _974_dt__update_hcf2_h0, (_972_dt__update__tmp_h0).dtor_cf3, _973_dt__update_hcf4_h0);
                  }(_pat_let18_0);
                }(_pat_let_tv21);
              }(_pat_let17_0);
            }(new BigNumber((_pat_let_tv20).length));
          }(_pat_let16_0);
        }(_958_v67)).dtor_cf3);
        let _index170 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
        (_930_v44)[_index170] = !((((p0).IsProperSubsetOf(_module.__default.fm35(_957_v66, (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))], _dafny.MultiSet.fromElements(p3), globalState))) ? ((_971_v69).IsDisjointFrom(_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0))))) : (p3)));
        let _index171 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length));
        let _index172 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
        let _rhs153 = (_this).fm1(p3, globalState);
        let _rhs154 = (_930_v44)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length))];
        let _lhs90 = p1;
        let _lhs91 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length));
        let _lhs92 = _930_v44;
        let _lhs93 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
        _lhs90[_lhs91] = _rhs153;
        _lhs92[_lhs93] = _rhs154;
        _955_v64 = _955_v64;
        let _index173 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
        (_930_v44)[_index173] = !(new BigNumber((p0).cardinality())).isEqualTo(_this.f2);
      }
      let _975_v70;
      _975_v70 = _dafny.MultiSet.fromElements((((_954_v63).contains(new BigNumber(215))) ? ((_954_v63).get(new BigNumber(215))) : (new BigNumber(405))), (_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))]), ((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))]).plus((_dafny.ZERO).minus(_this.f2)));
      let _index174 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      let _rhs155 = p2;
      let _rhs156 = (p0).Intersect(_975_v70);
      let _lhs94 = _930_v44;
      let _lhs95 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_930_v44).length));
      _lhs94[_lhs95] = _rhs155;
      _975_v70 = _rhs156;
      let _976_v71;
      _976_v71 = _dafny.Set.fromElements((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))]);
      let _977_v72;
      _977_v72 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))],p2);
      let _978_v73;
      _978_v73 = _dafny.Map.Empty.slice().updateUnsafe(_932_v46,_977_v72);
      let _979_v74;
      _979_v74 = _module.D5.create_DC10(p3, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_976_v71).length))), _978_v73, true);
      r0 = (_979_v74).dtor_cf18;
      let _980_v75;
      _980_v75 = _dafny.Seq.of(_976_v71);
      r1 = new BigNumber(((_980_v75)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((p1).length))], new BigNumber((_980_v75).length))]).length);
      return [r0, r1];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _981_v0;
      let _init15 = function (_982_i0) {
        return _module.__default.safeDivisionInt(_982_i0, _this.f2);
      };
      let _nw178 = Array((new BigNumber(29)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw178.length); _i0_15++) {
        _nw178[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _981_v0 = _nw178;
      let _983_v1;
      _983_v1 = _dafny.Set.fromElements(p0);
      let _984_v2;
      _984_v2 = _dafny.Seq.of(new BigNumber(603));
      let _index175 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length));
      (_981_v0)[_index175] = new BigNumber(((_983_v1).Union(function () {
        let _coll31 = new _dafny.Set();
        for (const _compr_31 of (_984_v2).Elements) {
          let _985_v3 = _compr_31;
          if (_dafny.Seq.contains(_984_v2, _985_v3)) {
            _coll31.add(_module.__default.safeDivisionInt(_985_v3, new BigNumber(916)));
          }
        }
        return _coll31;
      }())).length);
      let _index176 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length));
      (_981_v0)[_index176] = (new BigNumber(((_module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), function (_986_i1) {
  return new _dafny.CodePoint('d'.codePointAt(0));
}))).dtor_cf21).length)).minus(p0);
      let _987_v4;
      let _nw179 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _987_v4 = _nw179;
      let _988_v5;
      _988_v5 = _dafny.Map.Empty.slice().updateUnsafe(_987_v4,_981_v0);
      _988_v5 = (_988_v5).update(_987_v4, _981_v0);
      let _989_v6;
      _989_v6 = true;
      let _rhs157 = _989_v6;
      let _rhs158 = !(_989_v6) || (_989_v6);
      let _rhs159 = _989_v6;
      r0 = _rhs157;
      r0 = _rhs158;
      r0 = _rhs159;
      let _990_v7;
      _990_v7 = _dafny.Map.Empty.slice().updateUnsafe(_989_v6,_989_v6);
      let _991_v8;
      _991_v8 = _dafny.Seq.of(_989_v6);
      let _992_v9;
      let _nw180 = Array((new BigNumber(5)).toNumber());
      _nw180[(_dafny.ZERO).toNumber()] = _989_v6;
      _nw180[(_dafny.ONE).toNumber()] = _989_v6;
      _nw180[(new BigNumber(2)).toNumber()] = (_991_v8)[_module.__default.safeIndex((_981_v0)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length))], new BigNumber((_991_v8).length))];
      _nw180[(new BigNumber(3)).toNumber()] = _989_v6;
      _nw180[(new BigNumber(4)).toNumber()] = _989_v6;
      _992_v9 = _nw180;
      let _993_v10;
      _993_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_990_v7).length), (_this).fm1(_989_v6, globalState)),_992_v9);
      let _994_v11;
      _994_v11 = _dafny.Map.Empty.slice().updateUnsafe(_989_v6,p0);
      let _995_v12;
      _995_v12 = _dafny.Seq.of(_994_v11);
      _993_v10 = (_993_v10).update(new BigNumber((_995_v12).length), _992_v9);
      r0 = _989_v6;
      let _996_i2;
      _996_i2 = _dafny.ZERO;
      L3: {
        while (_989_v6) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_996_i2)) {
              break L3;
            }
            _996_i2 = (_996_i2).plus(_dafny.ONE);
            let _997_v13;
            _997_v13 = new _dafny.CodePoint('y'.codePointAt(0));
            let _998_v14;
            _998_v14 = _dafny.Seq.UnicodeFromString("mxs");
            let _rhs160 = (((_990_v7).contains(_989_v6)) ? ((_990_v7).get(_989_v6)) : (_989_v6));
            let _rhs161 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_984_v2, _984_v2), _dafny.Seq.Concat(_984_v2, _module.__default.fm36(_989_v6, _997_v13, new BigNumber((_dafny.Set.fromElements(_989_v6, (_this).fm8(_998_v14, _989_v6, globalState))).length), globalState))));
            _989_v6 = _rhs160;
            _989_v6 = _rhs161;
            if (!(_989_v6)) {
              _997_v13 = _997_v13;
              _989_v6 = _989_v6;
              _989_v6 = _989_v6;
              let _999_v15;
              _999_v15 = _dafny.Set.fromElements(!(_989_v6), _989_v6, (p0).isEqualTo(_this.f2));
              _999_v15 = (_999_v15).Union(_999_v15);
              _989_v6 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(675)), _dafny.Seq.update(_984_v2, _module.__default.safeIndex(p0, new BigNumber((_984_v2).length)), new BigNumber((_998_v14).length))), _dafny.Seq.Concat(_dafny.Seq.of((_981_v0)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-907)), ((_1000_v8) => function (_1001_i3) {
                return new BigNumber((_dafny.MultiSet.FromArray(_1000_v8)).cardinality());
              })(_991_v8))));
            } else {
              let _index177 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_992_v9).length));
              (_992_v9)[_index177] = _989_v6;
              let _1002_v16;
              _1002_v16 = _dafny.Set.fromElements(_997_v13, _module.__default.fm37((_this).fm1(_989_v6, globalState), globalState));
              let _index178 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_992_v9).length));
              (_992_v9)[_index178] = !((_1002_v16).equals(_1002_v16)) || ((_dafny.MultiSet.FromArray(_991_v8)).IsDisjointFrom(_dafny.MultiSet.fromElements(true)));
              (_this).f2 = (_981_v0)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length))];
              let _1003_v17;
              _1003_v17 = _module.D0.create_DC0(!((_992_v9)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_992_v9).length))]));
              _989_v6 = !_dafny.Seq.contains(_984_v2, (_this).fm1((_1003_v17).dtor_cf0, globalState));
              let _1004_v18;
              let _out22;
              _out22 = _module.__default.m0(_989_v6, (_992_v9)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_992_v9).length))], globalState);
              _1004_v18 = _out22;
              let _1005_v19;
              let _nw181 = new _module.C1();
              _nw181.__ctor(_998_v14);
              _1005_v19 = _nw181;
              let _1006_v20;
              _1006_v20 = _module.D8.create_DC17(_989_v6, _1005_v19, p0);
              let _1007_v21;
              let _nw182 = Array((_dafny.ONE).toNumber());
              _nw182[(_dafny.ZERO).toNumber()] = _1006_v20;
              _1007_v21 = _nw182;
              let _index179 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_1007_v21).length));
              (_1007_v21)[_index179] = _1006_v20;
              let _index180 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_1007_v21).length));
              (_1007_v21)[_index180] = _1006_v20;
            }
            let _1008_v22;
            let _nw183 = new _module.C0();
            _nw183.__ctor(_989_v6, _989_v6);
            _1008_v22 = _nw183;
            let _1009_v23;
            let _nw184 = new _module.C1();
            _nw184.__ctor(_998_v14);
            _1009_v23 = _nw184;
            let _1010_v24;
            _1010_v24 = _module.D8.create_DC16(_1009_v23);
            let _source19 = ((_989_v6) ? (_1010_v24) : (_1010_v24));
            if (_source19.is_DC17) {
              let _1011___mcc_h0 = (_source19).cf30;
              let _1012___mcc_h1 = (_source19).cf31;
              let _1013___mcc_h2 = (_source19).cf32;
              let _1014_cf32 = _1013___mcc_h2;
              let _1015_cf31 = _1012___mcc_h1;
              let _1016_cf30 = _1011___mcc_h0;
              let _1017_v25;
              let _nw185 = Array((new BigNumber(7)).toNumber()).fill([]);
              _1017_v25 = _nw185;
              _1017_v25 = _1017_v25;
              let _1018_v26;
              _1018_v26 = _dafny.Set.fromElements(_1016_cf30, !(_1008_v22.f5), _1008_v22.f5);
              let _1019_v27;
              _1019_v27 = _module.D4.create_DC5(_1018_v26);
              let _1020_v28;
              _1020_v28 = _module.D4.create_DC7(_1019_v27);
              _1020_v28 = _module.D4.create_DC7(_module.D4.create_DC5(_dafny.Set.fromElements(true, false, _1016_cf30, _1008_v22.f4, _1008_v22.f4)));
              _998_v14 = _dafny.Seq.UnicodeFromString("cfo");
              let _1021_v29;
              _1021_v29 = _dafny.MultiSet.fromElements(_1008_v22.f5);
              _1021_v29 = ((_1021_v29).Difference(_1021_v29)).Intersect(_1021_v29);
            } else {
              let _1022___mcc_h3 = (_source19).cf29;
              let _1023_cf29 = _1022___mcc_h3;
              let _index181 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length));
              (_981_v0)[_index181] = ((_981_v0)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length))]).minus((_dafny.ZERO).minus((_this.f2).multipliedBy((_this).fm1(_1008_v22.f5, globalState))));
              let _1024_v31;
              _1024_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1009_v23).f3,(_981_v0)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_981_v0).length))]);
              let _1025_v32;
              _1025_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1008_v22.f4);
              let _1026_v33;
              _1026_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1009_v23).f3,_1025_v32);
              let _1027_v34;
              _1027_v34 = _dafny.Seq.of(function () {
                let _coll32 = new _dafny.Map();
                for (const _compr_32 of (_1024_v31).Keys.Elements) {
                  let _1028_v30 = _compr_32;
                  if ((_1024_v31).contains(_1028_v30)) {
                    _coll32.push([_1028_v30,_dafny.Map.Empty.slice().updateUnsafe(p0,!(_1008_v22.f5))]);
                  }
                }
                return _coll32;
              }(), _1026_v33);
              let _1029_v35;
              _1029_v35 = _module.D5.create_DC10((_989_v6) && (false), (_dafny.ZERO).minus(_this.f2), (_1027_v34)[_module.__default.safeIndex(p0, new BigNumber((_1027_v34).length))], _1008_v22.f5);
              let _rhs162 = _this.f2;
              let _rhs163 = _1029_v35;
              r1 = _rhs162;
              _1029_v35 = _rhs163;
              let _1030_v36;
              let _nw186 = new _module.C1();
              _nw186.__ctor(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("r"), _module.__default.safeIndex(new BigNumber((_991_v8).length), new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)), _997_v13), _dafny.Seq.UnicodeFromString("log")));
              _1030_v36 = _nw186;
              let _1031_v37;
              _1031_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1008_v22,_1008_v22.f4);
              _1031_v37 = (_1031_v37).update(_1008_v22, _1008_v22.f5);
            }
          }
        }
      }
      r0 = !(_989_v6);
      let _1032_v38;
      _1032_v38 = _dafny.Seq.UnicodeFromString("nbprgidjf");
      r1 = new BigNumber((((_989_v6) ? (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), function (_1033_i4) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _1032_v38)) : (_1032_v38))).length);
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f2 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    __ctor(f2) {
      let _this = this;
      (_this)._f2 = f2;
      return;
    }
    fm8(p0, p1, globalState) {
      let _this = this;
      return !((_dafny.Set.fromElements(!(false), false)).Union(_dafny.Set.fromElements(!(true)))).contains(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("haphhjf"), _dafny.Seq.UnicodeFromString("jktmaskl")));
    };
    fm0(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(((!(false)) ? (_dafny.Seq.of(_this.f2)) : (_dafny.Seq.of(_this.f2))),_this.f2);
    };
    fm1(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_dafny.ZERO).minus(((!(false)) ? (_this.f2) : (_this.f2))))).plus(_module.__default.safeModuloInt(_this.f2, _this.f2));
    };
    fm10(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.UnicodeFromString("tues")) : (_dafny.Seq.UnicodeFromString("kchnvrhfl"))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), function (_1034_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }));
    };
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _1035_v0;
      _1035_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,p3);
      _1035_v0 = _1035_v0;
      let _1036_v1;
      _1036_v1 = _dafny.Seq.of(_this.f2);
      (_this).f2 = new BigNumber((_dafny.Seq.Concat(_1036_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(826)), function (_1037_i0) {
        return _this.f2;
      }))).length);
      let _1038_v2;
      _1038_v2 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1039_v3;
      _1039_v3 = _module.D0.create_DC1(true, !(p3), _1038_v2, new BigNumber(87));
      let _1040_i1;
      _1040_i1 = _dafny.ZERO;
      L4: {
        while ((((_this).fm1(p2, globalState)).isEqualTo(_this.f2)) || (((p2) ? (p2) : ((_1039_v3).dtor_cf1)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1040_i1)) {
              break L4;
            }
            _1040_i1 = (_1040_i1).plus(_dafny.ONE);
            let _1041_v4;
            let _init16 = ((_1042_v2) => function (_1043_i2) {
              return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_1044_v2) => function (_1045_i3) {
                return _1044_v2;
              })(_1042_v2)), _dafny.Seq.UnicodeFromString("lap"));
            })(_1038_v2);
            let _nw187 = Array((new BigNumber(24)).toNumber());
            for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw187.length); _i0_16++) {
              _nw187[_i0_16] = _init16(new BigNumber(_i0_16));
            }
            _1041_v4 = _nw187;
            let _1046_v5;
            _1046_v5 = _dafny.Seq.UnicodeFromString("amusabog");
            let _index182 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1041_v4).length));
            (_1041_v4)[_index182] = _1046_v5;
            let _index183 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1041_v4).length));
            (_1041_v4)[_index183] = _1046_v5;
            let _1047_v6;
            _1047_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus((_this).fm1(p2, globalState)));
            let _1048_v7;
            _1048_v7 = _dafny.Seq.of(_1047_v6);
            let _1049_v8;
            _1049_v8 = _dafny.Set.fromElements(p2, p3);
            let _1050_v9;
            _1050_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1049_v8).length),_1047_v6);
            let _1051_v10;
            let _nw188 = Array((new BigNumber(19)).toNumber());
            _nw188[(_dafny.ZERO).toNumber()] = (_1048_v7)[_module.__default.safeIndex(_this.f2, new BigNumber((_1048_v7).length))];
            _nw188[(_dafny.ONE).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f2);
            _nw188[(new BigNumber(3)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(4)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(5)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(6)).toNumber()] = (_1047_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_this.f2));
            _nw188[(new BigNumber(7)).toNumber()] = (_1047_v6).update(p1, _this.f2);
            _nw188[(new BigNumber(8)).toNumber()] = (_1047_v6).Merge(_1047_v6);
            _nw188[(new BigNumber(9)).toNumber()] = (_1047_v6).Merge(_1047_v6);
            _nw188[(new BigNumber(10)).toNumber()] = (_1047_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_this.f2));
            _nw188[(new BigNumber(11)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(12)).toNumber()] = ((_1047_v6).update(p1, new BigNumber(((_1041_v4)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1041_v4).length))]).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1046_v5).length)));
            _nw188[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f2);
            _nw188[(new BigNumber(14)).toNumber()] = (_1047_v6).Merge(_1047_v6);
            _nw188[(new BigNumber(15)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(16)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(17)).toNumber()] = _1047_v6;
            _nw188[(new BigNumber(18)).toNumber()] = (((_1050_v9).contains(_this.f2)) ? ((_1050_v9).get(_this.f2)) : (_1047_v6));
            _1051_v10 = _nw188;
            let _index184 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1051_v10).length));
            (_1051_v10)[_index184] = _1047_v6;
            let _index185 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1051_v10).length));
            (_1051_v10)[_index185] = _1047_v6;
            r0 = new BigNumber((_dafny.Seq.Concat(_1046_v5, _dafny.Seq.UnicodeFromString("l"))).length);
            let _1052_v11;
            _1052_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_this.f2);
            let _index186 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((p1).length));
            (p1)[_index186] = (_1036_v1)[_module.__default.safeIndex(new BigNumber((_1052_v11).length), new BigNumber((_1036_v1).length))];
            let _index187 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((p1).length));
            (p1)[_index187] = _this.f2;
          }
        }
      }
      let _1053_v12;
      _1053_v12 = _dafny.Set.fromElements(_this.f2);
      (_this).f2 = new BigNumber((_1053_v12).length);
      let _1054_v13;
      _1054_v13 = false;
      let _1055_v14;
      _1055_v14 = _dafny.Seq.UnicodeFromString("unuhbnudi");
      let _1056_v15;
      let _nw189 = Array((new BigNumber(19)).toNumber()).fill([]);
      _1056_v15 = _nw189;
      let _1057_v16;
      _1057_v16 = _dafny.Set.fromElements(_dafny.Set.fromElements(p3));
      let _rhs164 = _1054_v13;
      let _rhs165 = (_1057_v16).IsSubsetOf(_1057_v16);
      let _rhs166 = _dafny.Seq.UnicodeFromString("mlb");
      let _rhs167 = _1056_v15;
      _1054_v13 = _rhs164;
      _1054_v13 = _rhs165;
      _1055_v14 = _rhs166;
      _1056_v15 = _rhs167;
      let _1058_v17;
      let _nw190 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1058_v17 = _nw190;
      let _index188 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1058_v17).length));
      (_1058_v17)[_index188] = _1054_v13;
      let _1059_v18;
      _1059_v18 = _module.D0.create_DC0(p2);
      let _pat_let_tv22 = p2;
      let _pat_let_tv23 = _1055_v14;
      let _pat_let_tv24 = _1059_v18;
      let _pat_let_tv25 = globalState;
      let _1060_v19;
      _1060_v19 = _dafny.MultiSet.fromElements((function (_pat_let19_0) {
        return function (_1063_dt__update__tmp_h1) {
          return function (_pat_let22_0) {
            return function (_1064_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_1064_dt__update_hcf0_h1);
            }(_pat_let22_0);
          }((_this).fm8(_pat_let_tv23, (_pat_let_tv24).dtor_cf0, _pat_let_tv25));
        }(_pat_let19_0);
      }(function (_pat_let20_0) {
        return function (_1061_dt__update__tmp_h0) {
          return function (_pat_let21_0) {
            return function (_1062_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_1062_dt__update_hcf0_h0);
            }(_pat_let21_0);
          }(_pat_let_tv22);
        }(_pat_let20_0);
      }(_1059_v18))).dtor_cf0);
      let _index189 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1058_v17).length));
      (_1058_v17)[_index189] = !(_1060_v19).contains(false);
      r0 = (_dafny.ZERO).minus((_this.f2).multipliedBy(_this.f2));
      let _1065_v20;
      _1065_v20 = _dafny.Seq.of(true);
      r1 = (_this).fm1((_1065_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1055_v14, _module.__default.safeIndex(_this.f2, new BigNumber((_1055_v14).length)), _1038_v2)).length), new BigNumber((_1065_v20).length))], globalState);
      return [r0, r1];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      if (p2) {
        if (p2) {
          let _1066_v0;
          _1066_v0 = _dafny.MultiSet.fromElements(new BigNumber(-215), (_this.f2).multipliedBy(new BigNumber(446)), (p0).minus(p0), p0);
          let _1067_v1;
          _1067_v1 = _dafny.Seq.of(_1066_v0, _dafny.MultiSet.fromElements(p0, p0));
          _1066_v0 = (_1066_v0).Intersect((_1067_v1)[_module.__default.safeIndex(_this.f2, new BigNumber((_1067_v1).length))]);
          (_this).f2 = _this.f2;
          let _1068_v2;
          _1068_v2 = _dafny.Set.fromElements(p0);
          let _1069_v3;
          _1069_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1068_v2,p0);
          _1069_v3 = (_1069_v3).update(_1068_v2, _this.f2);
          let _1070_v4;
          _1070_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),p2);
          let _1071_v5;
          let _nw191 = Array((new BigNumber(6)).toNumber());
          _nw191[(_dafny.ZERO).toNumber()] = ((((_1070_v4).contains(p3)) ? ((_1070_v4).get(p3)) : (p3))) || (p2);
          _nw191[(_dafny.ONE).toNumber()] = (p0).isEqualTo(p0);
          _nw191[(new BigNumber(2)).toNumber()] = true;
          _nw191[(new BigNumber(3)).toNumber()] = (p0).isEqualTo(p0);
          _nw191[(new BigNumber(4)).toNumber()] = p3;
          _nw191[(new BigNumber(5)).toNumber()] = !(p2);
          _1071_v5 = _nw191;
          let _index190 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1071_v5).length));
          (_1071_v5)[_index190] = p2;
          let _1072_v6;
          _1072_v6 = _dafny.Seq.UnicodeFromString("ecarnc");
          let _index191 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1071_v5).length));
          (_1071_v5)[_index191] = (_this).fm8(_1072_v6, p3, globalState);
          let _1073_v7;
          _1073_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1072_v6,_this.f2);
          _1073_v7 = (_1073_v7).update(_1072_v6, _module.__default.safeDivisionInt(p0, p0));
        } else {
          r0 = p3;
          let _1074_v8;
          let _nw192 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1074_v8 = _nw192;
          let _1075_v9;
          let _nw193 = Array((new BigNumber(5)).toNumber());
          _nw193[(_dafny.ZERO).toNumber()] = _1074_v8;
          _nw193[(_dafny.ONE).toNumber()] = _1074_v8;
          _nw193[(new BigNumber(2)).toNumber()] = _1074_v8;
          _nw193[(new BigNumber(3)).toNumber()] = _1074_v8;
          _nw193[(new BigNumber(4)).toNumber()] = _1074_v8;
          _1075_v9 = _nw193;
          let _index192 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_1075_v9).length));
          (_1075_v9)[_index192] = _1074_v8;
          let _index193 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_1075_v9).length));
          (_1075_v9)[_index193] = _1074_v8;
          r2 = (((p3) ? (_this.f2) : (p0))).multipliedBy(_module.__default.safeModuloInt(p0, _this.f2));
          let _1076_v10;
          _1076_v10 = _dafny.Seq.UnicodeFromString("k");
          r0 = (_this).fm8(_1076_v10, false, globalState);
          let _1077_v11;
          _1077_v11 = new _dafny.CodePoint('n'.codePointAt(0));
          _1077_v11 = _1077_v11;
        }
        let _1078_v12;
        let _init17 = function (_1079_i0) {
          return (_1079_i0).multipliedBy(_this.f2);
        };
        let _nw194 = Array((new BigNumber(25)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw194.length); _i0_17++) {
          _nw194[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _1078_v12 = _nw194;
        _1078_v12 = _1078_v12;
        if (p2) {
          let _1080_v13;
          let _nw195 = Array((_dafny.ONE).toNumber()).fill(false);
          _1080_v13 = _nw195;
          let _1081_v14;
          _1081_v14 = _dafny.MultiSet.fromElements(p0);
          let _index194 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length));
          (_1080_v13)[_index194] = (_1081_v14).IsDisjointFrom(_1081_v14);
          let _index195 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length));
          (_1080_v13)[_index195] = p3;
          let _1082_v15;
          _1082_v15 = _dafny.MultiSet.fromElements(p3);
          let _1083_v16;
          let _nw196 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1083_v16 = _nw196;
          let _1084_v17;
          _1084_v17 = _dafny.Map.Empty.slice().updateUnsafe((_1082_v15).Union(_1082_v15),_1083_v16);
          _1084_v17 = _1084_v17;
          let _1085_v18;
          _1085_v18 = _dafny.Seq.UnicodeFromString("tgfihkea");
          let _index196 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length));
          (_1080_v13)[_index196] = (_this).fm8(_dafny.Seq.Concat(_1085_v18, _1085_v18), (_1080_v13)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length))], globalState);
          let _1086_v19;
          let _nw197 = new _module.C0();
          _nw197.__ctor(p2, !((_1080_v13)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length))]));
          _1086_v19 = _nw197;
          let _index197 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1080_v13).length));
          (_1080_v13)[_index197] = !(!(!(p3) || (p2)));
        } else {
          let _1087_v20;
          _1087_v20 = _dafny.Seq.of(p3);
          let _1088_v21;
          _1088_v21 = _module.D0.create_DC1(p3, p2, new _dafny.CodePoint('u'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("tplkvlen")).length));
          let _1089_v22;
          _1089_v22 = new _dafny.CodePoint('h'.codePointAt(0));
          let _pat_let_tv26 = _1089_v22;
          let _pat_let_tv27 = p2;
          let _1090_v23;
          _1090_v23 = _dafny.Set.fromElements(function (_pat_let23_0) {
            return function (_1091_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_1092_dt__update_hcf3_h0) {
                  return function (_pat_let25_0) {
                    return function (_1093_dt__update_hcf2_h0) {
                      return _module.D0.create_DC1((_1091_dt__update__tmp_h0).dtor_cf1, _1093_dt__update_hcf2_h0, _1092_dt__update_hcf3_h0, (_1091_dt__update__tmp_h0).dtor_cf4);
                    }(_pat_let25_0);
                  }(_pat_let_tv27);
                }(_pat_let24_0);
              }(_pat_let_tv26);
            }(_pat_let23_0);
          }(_1088_v21), _1088_v21, _1088_v21);
          let _1094_v24;
          let _nw198 = Array((new BigNumber(14)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = !(p2);
          _nw198[(_dafny.ONE).toNumber()] = (p0).isLessThan(new BigNumber(175));
          _nw198[(new BigNumber(2)).toNumber()] = (_this.f2).isEqualTo(_this.f2);
          _nw198[(new BigNumber(3)).toNumber()] = p2;
          _nw198[(new BigNumber(4)).toNumber()] = p3;
          _nw198[(new BigNumber(5)).toNumber()] = p3;
          _nw198[(new BigNumber(6)).toNumber()] = p3;
          _nw198[(new BigNumber(7)).toNumber()] = (_1087_v20)[_module.__default.safeIndex(p0, new BigNumber((_1087_v20).length))];
          _nw198[(new BigNumber(8)).toNumber()] = p2;
          _nw198[(new BigNumber(9)).toNumber()] = p3;
          _nw198[(new BigNumber(10)).toNumber()] = p3;
          _nw198[(new BigNumber(11)).toNumber()] = (_dafny.Set.fromElements(_1088_v21, _1088_v21)).IsSubsetOf(_1090_v23);
          _nw198[(new BigNumber(12)).toNumber()] = p2;
          _nw198[(new BigNumber(13)).toNumber()] = p2;
          _1094_v24 = _nw198;
          let _index198 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1094_v24).length));
          (_1094_v24)[_index198] = !(!((_this).fm8(_dafny.Seq.UnicodeFromString("ihkkeqm"), p2, globalState)));
          let _index199 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1094_v24).length));
          (_1094_v24)[_index199] = p3;
          let _1095_v25;
          _1095_v25 = _dafny.Seq.UnicodeFromString("cwutipmey");
          _1087_v20 = _dafny.Seq.of((p0).isLessThanOrEqualTo(new BigNumber((_1095_v25).length)), p3, !(p3));
          let _1096_v26;
          let _nw199 = new _module.C0();
          _nw199.__ctor(p3, p2);
          _1096_v26 = _nw199;
          let _1097_v27;
          _1097_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_this.f2);
          let _1098_v28;
          _1098_v28 = _dafny.Set.fromElements(p0);
          let _1099_v29;
          _1099_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1098_v28);
          let _1100_v30;
          _1100_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.update(_1095_v25, _module.__default.safeIndex(new BigNumber((_1099_v29).length), new BigNumber((_1095_v25).length)), _1089_v22));
          _1097_v27 = (_1097_v27).update((p0).multipliedBy(_this.f2), new BigNumber((_1100_v30).length));
          r2 = ((_1096_v26.f4) ? ((p0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_1101_p0) => function (_1102_i1) {
            return _1101_p0;
          })(p0))).length))) : (_this.f2));
        }
        let _1103_v31;
        _1103_v31 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
        r0 = (((_1103_v31).contains(p3)) ? ((_1103_v31).get(p3)) : (p2));
        let _1104_v32;
        let _init18 = ((_1105_v31) => function (_1106_i2) {
          return _1105_v31;
        })(_1103_v31);
        let _nw200 = Array((new BigNumber(29)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw200.length); _i0_18++) {
          _nw200[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _1104_v32 = _nw200;
        let _1107_v33;
        _1107_v33 = _1103_v31;
        let _index200 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1104_v32).length));
        (_1104_v32)[_index200] = _1107_v33;
        let _index201 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1104_v32).length));
        (_1104_v32)[_index201] = _1107_v33;
      } else {
        r2 = _this.f2;
        let _1108_v34;
        _1108_v34 = _dafny.Seq.of(p3, p3, p2, p2);
        let _1109_v35;
        _1109_v35 = _dafny.Seq.UnicodeFromString("pdbarimii");
        let _1110_v36;
        _1110_v36 = _module.D6.create_DC11(_1109_v35);
        r2 = ((!((_1108_v34)[_module.__default.safeIndex(_this.f2, new BigNumber((_1108_v34).length))])) ? ((new BigNumber(121)).minus(new BigNumber(((_1110_v36).dtor_cf21).length))) : (_this.f2));
        let _1111_v37;
        _1111_v37 = _dafny.Set.fromElements(_this.f2);
        let _1112_v38;
        _1112_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1111_v37).Difference(_dafny.Set.fromElements(new BigNumber(511), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p3),true)).length), new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()))),p2);
        _1112_v38 = _1112_v38;
        let _1113_v39;
        let _nw201 = new _module.C0();
        _nw201.__ctor(false, p2);
        _1113_v39 = _nw201;
        r0 = p3;
      }
      let _1114_v40;
      _1114_v40 = _dafny.Seq.of(p3);
      if ((new BigNumber((_dafny.Seq.Concat(_1114_v40, _1114_v40)).length)).isEqualTo(p0)) {
        let _1115_v42;
        _1115_v42 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _1116_v43;
        _1116_v43 = _dafny.Seq.of(_1115_v42, _1115_v42);
        let _1117_v44;
        _1117_v44 = _dafny.Seq.of(_1116_v43, _1116_v43, _1116_v43);
        let _1118_v45;
        _1118_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p3,p2),p2);
        let _1119_v46;
        _1119_v46 = _dafny.Set.fromElements(_this.f2, _this.f2);
        let _1120_v47;
        _1120_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1119_v46,_1118_v45);
        let _1121_v48;
        _1121_v48 = _dafny.MultiSet.fromElements(false, p2);
        let _1122_v51;
        _1122_v51 = _dafny.Seq.UnicodeFromString("gpieicl");
        let _1123_v52;
        _1123_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,_1122_v51);
        let _1124_v54;
        let _nw202 = Array((new BigNumber(24)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = function () {
          let _coll33 = new _dafny.Map();
          for (const _compr_33 of ((_1117_v44)[_module.__default.safeIndex(p0, new BigNumber((_1117_v44).length))]).Elements) {
            let _1125_v41 = _compr_33;
            if (_dafny.Seq.contains((_1117_v44)[_module.__default.safeIndex(p0, new BigNumber((_1117_v44).length))], _1125_v41)) {
              _coll33.push([_1125_v41,p2]);
            }
          }
          return _coll33;
        }();
        _nw202[(_dafny.ONE).toNumber()] = (_1118_v45).update(_1115_v42, p2);
        _nw202[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,(_1114_v40)[_module.__default.safeIndex(p0, new BigNumber((_1114_v40).length))]);
        _nw202[(new BigNumber(3)).toNumber()] = (_1118_v45).update(_1115_v42, p3);
        _nw202[(new BigNumber(4)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(5)).toNumber()] = (((_1120_v47).contains(_module.__default.fm18(p2, p2, _1121_v48, globalState))) ? ((_1120_v47).get(_module.__default.fm18(p2, p2, _1121_v48, globalState))) : (function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of (_1116_v43).Elements) {
            let _1126_v49 = _compr_34;
            if (_dafny.Seq.contains(_1116_v43, _1126_v49)) {
              _coll34.push([_1126_v49,p3]);
            }
          }
          return _coll34;
        }()));
        _nw202[(new BigNumber(6)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(7)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,p2);
        _nw202[(new BigNumber(9)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(10)).toNumber()] = (_1118_v45).update(_1115_v42, p3);
        _nw202[(new BigNumber(11)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(12)).toNumber()] = (_1118_v45).Merge(_module.__default.fm20(globalState));
        _nw202[(new BigNumber(13)).toNumber()] = function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_1123_v52).Keys.Elements) {
            let _1127_v50 = _compr_35;
            if ((_1123_v52).contains(_1127_v50)) {
              _coll35.push([_1127_v50,p3]);
            }
          }
          return _coll35;
        }();
        _nw202[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,p3);
        _nw202[(new BigNumber(15)).toNumber()] = (_module.__default.fm20(globalState)).Merge(function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_1118_v45).Keys.Elements) {
            let _1128_v53 = _compr_36;
            if ((_1118_v45).contains(_1128_v53)) {
              _coll36.push([_1128_v53,p3]);
            }
          }
          return _coll36;
        }());
        _nw202[(new BigNumber(16)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(17)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(18)).toNumber()] = (_1118_v45).Merge((_1118_v45).update(_1115_v42, p3));
        _nw202[(new BigNumber(19)).toNumber()] = _1118_v45;
        _nw202[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,true);
        _nw202[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,(((_1115_v42).contains(p3)) ? ((_1115_v42).get(p3)) : (p3)));
        _nw202[(new BigNumber(22)).toNumber()] = (_1118_v45).update(_1115_v42, p2);
        _nw202[(new BigNumber(23)).toNumber()] = _1118_v45;
        _1124_v54 = _nw202;
        let _index202 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1124_v54).length));
        (_1124_v54)[_index202] = _dafny.Map.Empty.slice().updateUnsafe(_1115_v42,(_1114_v40)[_module.__default.safeIndex(new BigNumber((_1122_v51).length), new BigNumber((_1114_v40).length))]);
        let _1129_v56;
        _1129_v56 = _dafny.Set.fromElements(_1115_v42, _module.__default.fm21(globalState));
        let _index203 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1124_v54).length));
        (_1124_v54)[_index203] = function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of ((_1129_v56).Intersect(_1129_v56)).Elements) {
            let _1130_v55 = _compr_37;
            if (((_1129_v56).Intersect(_1129_v56)).contains(_1130_v55)) {
              _coll37.push([_1130_v55,(((_1115_v42).contains(p3)) ? ((_1115_v42).get(p3)) : (p3))]);
            }
          }
          return _coll37;
        }();
        if (p3) {
          (_this).f2 = p0;
          let _1131_v57;
          let _nw203 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1131_v57 = _nw203;
          let _1132_v58;
          _1132_v58 = _dafny.Seq.of(_1131_v57, _1131_v57);
          r0 = _dafny.Seq.IsPrefixOf(_1132_v58, _1132_v58);
          r0 = !(p0).isEqualTo((p0).plus(p0));
          let _index204 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1131_v57).length));
          (_1131_v57)[_index204] = (!(p3)) === (p3);
          let _index205 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1131_v57).length));
          (_1131_v57)[_index205] = p2;
          let _1133_v59;
          _1133_v59 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          let _1134_v60;
          _1134_v60 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm15(globalState));
          let _1135_v61;
          _1135_v61 = new _dafny.CodePoint('t'.codePointAt(0));
          _1133_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this.f2).isLessThan(new BigNumber((_1134_v60).length)),new BigNumber((_dafny.Seq.Concat(_1122_v51, _dafny.Seq.update(_1122_v51, _module.__default.safeIndex((_dafny.ZERO).minus(_this.f2), new BigNumber((_1122_v51).length)), _1135_v61))).length));
        } else {
          let _1136_v62;
          let _init19 = ((_1137_v51) => function (_1138_i3) {
            return (_dafny.MultiSet.fromElements(_1137_v51)).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), function (_1139_i4) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })));
          })(_1122_v51);
          let _nw204 = Array((new BigNumber(13)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw204.length); _i0_19++) {
            _nw204[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _1136_v62 = _nw204;
          let _1140_v63;
          _1140_v63 = _dafny.MultiSet.fromElements(_1122_v51, _1122_v51, _1122_v51);
          let _index206 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1136_v62).length));
          (_1136_v62)[_index206] = (_module.__default.fm22(p2, _this.f2, p2, p0, globalState)).Difference(_1140_v63);
          let _1141_v64;
          let _nw205 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1141_v64 = _nw205;
          let _1142_v65;
          _1142_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1141_v64,((p2) ? (_1140_v63) : (_1140_v63)));
          let _index207 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1136_v62).length));
          (_1136_v62)[_index207] = (((_1142_v65).contains(_1141_v64)) ? ((_1142_v65).get(_1141_v64)) : (_1140_v63));
          let _1143_v66;
          let _nw206 = new _module.C1();
          _nw206.__ctor(_dafny.Seq.Concat(_1122_v51, _1122_v51));
          _1143_v66 = _nw206;
          let _1144_v67;
          _1144_v67 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1145_v68;
          _1145_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _1146_v69;
          _1146_v69 = _module.D5.create_DC9(false, p0, _1121_v48);
          let _1147_v70;
          _1147_v70 = _module.D6.create_DC13(p3, p0);
          let _1148_v71;
          _1148_v71 = _dafny.Seq.of(_1147_v70, _1147_v70);
          let _1149_v72;
          _1149_v72 = _dafny.Seq.of(_this.f2, new BigNumber((_1148_v71).length));
          let _1150_v73;
          _1150_v73 = _module.D4.create_DC6((_1143_v66).fm1(p2, globalState), _1141_v64, new BigNumber((_1149_v72).length));
          let _1151_v74;
          let _nw207 = Array((new BigNumber(24)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = p3;
          _nw207[(_dafny.ONE).toNumber()] = p2;
          _nw207[(new BigNumber(2)).toNumber()] = p2;
          _nw207[(new BigNumber(3)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), ((_1152_v67) => function (_1153_i5) {
            return _1152_v67;
          })(_1144_v67)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), ((_1154_v67) => function (_1155_i5) {
            return _1154_v67;
          })(_1144_v67))).length)), _1144_v67), _1144_v67);
          _nw207[(new BigNumber(4)).toNumber()] = p3;
          _nw207[(new BigNumber(5)).toNumber()] = p2;
          _nw207[(new BigNumber(6)).toNumber()] = (((_1145_v68).contains(_this.f2)) ? ((_1145_v68).get(_this.f2)) : (!(true)));
          _nw207[(new BigNumber(7)).toNumber()] = (_1143_v66).fm11(globalState);
          _nw207[(new BigNumber(8)).toNumber()] = p2;
          _nw207[(new BigNumber(9)).toNumber()] = (_this.f2).isLessThanOrEqualTo(_this.f2);
          _nw207[(new BigNumber(10)).toNumber()] = p2;
          _nw207[(new BigNumber(11)).toNumber()] = (_1146_v69).dtor_cf14;
          _nw207[(new BigNumber(12)).toNumber()] = p3;
          _nw207[(new BigNumber(13)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
          _nw207[(new BigNumber(14)).toNumber()] = p2;
          _nw207[(new BigNumber(15)).toNumber()] = ((p3) ? (!(p3)) : (p2));
          _nw207[(new BigNumber(16)).toNumber()] = p2;
          _nw207[(new BigNumber(17)).toNumber()] = (_1114_v40)[_module.__default.safeIndex((_1150_v73).dtor_cf11, new BigNumber((_1114_v40).length))];
          _nw207[(new BigNumber(18)).toNumber()] = p3;
          _nw207[(new BigNumber(19)).toNumber()] = p3;
          _nw207[(new BigNumber(20)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(p0, _this.f2), _1149_v72));
          _nw207[(new BigNumber(21)).toNumber()] = p2;
          _nw207[(new BigNumber(22)).toNumber()] = true;
          _nw207[(new BigNumber(23)).toNumber()] = p3;
          _1151_v74 = _nw207;
          let _index208 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1151_v74).length));
          (_1151_v74)[_index208] = p3;
          let _index209 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1151_v74).length));
          (_1151_v74)[_index209] = p3;
          _1114_v40 = _1114_v40;
          _1141_v64 = _1141_v64;
        }
        let _1156_v75;
        _1156_v75 = new _dafny.CodePoint('m'.codePointAt(0));
        _1156_v75 = _1156_v75;
        r0 = p2;
        let _1157_v76;
        _1157_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(879),_module.__default.fm19(_this.f2, p0, false, _this.f2, globalState));
        let _1158_v77;
        let _nw208 = new _module.C0();
        _nw208.__ctor(!(new BigNumber(-623)).isEqualTo(new BigNumber((_1157_v76).length)), p2);
        _1158_v77 = _nw208;
      } else {
        r0 = p2;
        let _1159_v78;
        _1159_v78 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1160_v79;
        _1160_v79 = _dafny.Seq.of(_1159_v78);
        let _1161_v80;
        let _nw209 = Array((new BigNumber(10)).toNumber());
        _nw209[(_dafny.ZERO).toNumber()] = _1159_v78;
        _nw209[(_dafny.ONE).toNumber()] = _module.__default.fm23(globalState);
        _nw209[(new BigNumber(2)).toNumber()] = (_1160_v79)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("wa")).length), new BigNumber((_1160_v79).length))];
        _nw209[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw209[(new BigNumber(4)).toNumber()] = ((p2) ? (_1159_v78) : (_1159_v78));
        _nw209[(new BigNumber(5)).toNumber()] = _1159_v78;
        _nw209[(new BigNumber(6)).toNumber()] = _1159_v78;
        _nw209[(new BigNumber(7)).toNumber()] = _1159_v78;
        _nw209[(new BigNumber(8)).toNumber()] = _1159_v78;
        _nw209[(new BigNumber(9)).toNumber()] = _1159_v78;
        _1161_v80 = _nw209;
        let _1162_v81;
        _1162_v81 = _dafny.Set.fromElements(p0, p0, new BigNumber((_1160_v79).length), p0);
        let _rhs168 = _1159_v78;
        let _rhs169 = _1161_v80;
        let _rhs170 = _1162_v81;
        _1159_v78 = _rhs168;
        _1161_v80 = _rhs169;
        _1162_v81 = _rhs170;
        let _1163_v82;
        _1163_v82 = _module.D0.create_DC0(p2);
        r0 = (_1163_v82).dtor_cf0;
        if ((p0).isLessThan(((p3) ? (_this.f2) : (_this.f2)))) {
          let _1164_v83;
          let _nw210 = new _module.C1();
          _nw210.__ctor(_1160_v79);
          _1164_v83 = _nw210;
          let _1165_v84;
          let _nw211 = Array((new BigNumber(26)).toNumber()).fill([]);
          _1165_v84 = _nw211;
          let _1166_v85;
          _1166_v85 = _1165_v84;
          _1166_v85 = _1166_v85;
          let _1167_v86;
          _1167_v86 = _module.D6.create_DC11(_dafny.Seq.update(_1160_v79, _module.__default.safeIndex(_this.f2, new BigNumber((_1160_v79).length)), _1159_v78));
          let _1168_v87;
          _1168_v87 = _module.D6.create_DC11((_1167_v86).dtor_cf21);
          let _1169_v88;
          _1169_v88 = _dafny.Seq.of(p0);
          let _rhs171 = (_this).fm8(_1160_v79, p2, globalState);
          let _rhs172 = _1168_v87;
          let _rhs173 = _module.__default.safeDivisionInt((_1169_v88)[_module.__default.safeIndex((_1164_v83).fm1(p2, globalState), new BigNumber((_1169_v88).length))], _module.__default.safeDivisionInt(p0, _this.f2));
          let _lhs96 = _this;
          r0 = _rhs171;
          _1168_v87 = _rhs172;
          _lhs96.f2 = _rhs173;
          (_this).f2 = new BigNumber((_dafny.Seq.update((_this).fm10(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update((_1164_v83).f3, _module.__default.safeIndex(p0, new BigNumber(((_1164_v83).f3).length)), new _dafny.CodePoint('r'.codePointAt(0)))).length), new BigNumber(-521)), globalState), _module.__default.safeIndex(new BigNumber(((_1164_v83).f3).length), new BigNumber(((_this).fm10(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update((_1164_v83).f3, _module.__default.safeIndex(p0, new BigNumber(((_1164_v83).f3).length)), new _dafny.CodePoint('r'.codePointAt(0)))).length), new BigNumber(-521)), globalState)).length)), _1159_v78)).length);
          let _1170_v89;
          _1170_v89 = _dafny.Map.Empty.slice().updateUnsafe((_1164_v83).fm1(true, globalState),new BigNumber(953));
          let _1171_v90;
          _1171_v90 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p2));
          let _1172_v91;
          let _nw212 = Array((new BigNumber(24)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = p0;
          _nw212[(_dafny.ONE).toNumber()] = (p0).minus(p0);
          _nw212[(new BigNumber(2)).toNumber()] = (((_1170_v89).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p0)).length)))) ? ((_1170_v89).get((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p0)).length)))) : (new BigNumber(593)));
          _nw212[(new BigNumber(3)).toNumber()] = _this.f2;
          _nw212[(new BigNumber(4)).toNumber()] = p0;
          _nw212[(new BigNumber(5)).toNumber()] = p0;
          _nw212[(new BigNumber(6)).toNumber()] = p0;
          _nw212[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw212[(new BigNumber(8)).toNumber()] = (_1164_v83).fm1(p3, globalState);
          _nw212[(new BigNumber(9)).toNumber()] = p0;
          _nw212[(new BigNumber(10)).toNumber()] = (p0).plus(new BigNumber(144));
          _nw212[(new BigNumber(11)).toNumber()] = (_this.f2).minus(p0);
          _nw212[(new BigNumber(12)).toNumber()] = (_this.f2).multipliedBy(_this.f2);
          _nw212[(new BigNumber(13)).toNumber()] = _this.f2;
          _nw212[(new BigNumber(14)).toNumber()] = p0;
          _nw212[(new BigNumber(15)).toNumber()] = p0;
          _nw212[(new BigNumber(16)).toNumber()] = p0;
          _nw212[(new BigNumber(17)).toNumber()] = new BigNumber((_1160_v79).length);
          _nw212[(new BigNumber(18)).toNumber()] = new BigNumber((_1171_v90).length);
          _nw212[(new BigNumber(19)).toNumber()] = p0;
          _nw212[(new BigNumber(20)).toNumber()] = (_this).fm1(p3, globalState);
          _nw212[(new BigNumber(21)).toNumber()] = _this.f2;
          _nw212[(new BigNumber(22)).toNumber()] = _this.f2;
          _nw212[(new BigNumber(23)).toNumber()] = p0;
          _1172_v91 = _nw212;
          let _index210 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1172_v91).length));
          (_1172_v91)[_index210] = _this.f2;
          let _index211 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1172_v91).length));
          let _rhs174 = ((_this).fm1(p2, globalState)).multipliedBy(_module.__default.safeModuloInt(_this.f2, new BigNumber((_dafny.Seq.of(_1159_v78, _1159_v78)).length)));
          let _rhs175 = (_1169_v88)[_module.__default.safeIndex(p0, new BigNumber((_1169_v88).length))];
          let _rhs176 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1160_v79, _dafny.Seq.UnicodeFromString("cpr")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(451)), ((_1173_v78) => function (_1174_i6) {
            return _1173_v78;
          })(_1159_v78)));
          let _rhs177 = ((new BigNumber(529)).plus(p0)).multipliedBy(new BigNumber(458));
          let _lhs97 = _1172_v91;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1172_v91).length));
          r2 = _rhs174;
          r2 = _rhs175;
          _1160_v79 = _rhs176;
          _lhs97[_lhs98] = _rhs177;
        } else {
          let _1175_v92;
          _1175_v92 = _dafny.MultiSet.fromElements((p3) && (p3));
          (_this).f2 = new BigNumber((_1175_v92).cardinality());
          _1160_v79 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(745)), ((_1176_v78) => function (_1177_i7) {
            return _1176_v78;
          })(_1159_v78));
          _1175_v92 = (_1175_v92).Difference(_1175_v92);
          let _1178_v93;
          let _nw213 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1178_v93 = _nw213;
          let _index212 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1178_v93).length));
          (_1178_v93)[_index212] = new BigNumber((_1160_v79).length);
          let _index213 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1178_v93).length));
          (_1178_v93)[_index213] = p0;
          let _1179_v94;
          let _nw214 = new _module.C1();
          _nw214.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), ((_1180_v78) => function (_1181_i8) {
            return _1180_v78;
          })(_1159_v78)));
          _1179_v94 = _nw214;
        }
        let _1182_v95;
        _1182_v95 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), ((_1183_p0) => function (_1184_i9) {
          return _1183_p0;
        })(p0)),p0);
        let _1185_v96;
        _1185_v96 = _1182_v95;
        let _1186_v97;
        _1186_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1185_v96,new BigNumber(30));
        r0 = !(((p0).plus(p0)).isLessThan((((_1186_v97).contains(_1185_v96)) ? ((_1186_v97).get(_1185_v96)) : (p0))));
      }
      r2 = (_dafny.ZERO).minus(_this.f2);
      let _1187_v98;
      _1187_v98 = _dafny.Seq.UnicodeFromString("brhukyi");
      let _hi4 = (_this.f2).plus(_this.f2);
      for (let _1188_i10 = _module.__default.safeModuloInt(_this.f2, new BigNumber((_1187_v98).length)); _1188_i10.isLessThan(_hi4); _1188_i10 = _1188_i10.plus(_dafny.ONE)) {
        if (!(p3)) {
          let _1189_v99;
          let _nw215 = new _module.C2();
          _nw215.__ctor((_this).fm1(p2, globalState));
          _1189_v99 = _nw215;
          let _1190_v100;
          _1190_v100 = _1189_v99;
          let _rhs178 = _module.__default.safeDivisionInt(p0, _this.f2);
          let _rhs179 = (_1190_v100);
          let _rhs180 = (_this).fm8(_1187_v98, p3, globalState);
          r2 = _rhs178;
          _1189_v99 = _rhs179;
          r0 = _rhs180;
          let _1191_v101;
          _1191_v101 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1192_v102;
          _1192_v102 = _dafny.MultiSet.fromElements(p2, p2, !(p2), true, p3);
          let _1193_v103;
          _1193_v103 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1194_v104;
          _1194_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1193_v103,_1192_v102);
          let _1195_v105;
          _1195_v105 = _module.D5.create_DC9(true, new BigNumber(571), _1192_v102);
          let _1196_v106;
          let _nw216 = Array((new BigNumber(28)).toNumber());
          _nw216[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(p2, p3);
          _nw216[(_dafny.ONE).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(2)).toNumber()] = _module.__default.fm38(p3, globalState);
          _nw216[(new BigNumber(3)).toNumber()] = ((p2) ? (_module.__default.fm38(p2, globalState)) : (_dafny.MultiSet.fromElements(p2, p3)));
          _nw216[(new BigNumber(4)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(5)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(6)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(7)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(8)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(9)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1114_v40, _1114_v40));
          _nw216[(new BigNumber(11)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(12)).toNumber()] = (((_1189_v99).fm8(_1187_v98, p2, globalState)) ? (_1192_v102) : (_dafny.MultiSet.FromArray(_1114_v40)));
          _nw216[(new BigNumber(13)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(14)).toNumber()] = (_1192_v102).update(p3, _module.__default.abs(p0));
          _nw216[(new BigNumber(15)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(16)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(17)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(18)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(19)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(20)).toNumber()] = ((((_1194_v104).contains(_1193_v103)) ? ((_1194_v104).get(_1193_v103)) : (_1192_v102))).Difference(_1192_v102);
          _nw216[(new BigNumber(21)).toNumber()] = (_1195_v105).dtor_cf16;
          _nw216[(new BigNumber(22)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(23)).toNumber()] = (_1192_v102).Union(_1192_v102);
          _nw216[(new BigNumber(24)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(25)).toNumber()] = _1192_v102;
          _nw216[(new BigNumber(26)).toNumber()] = (_dafny.MultiSet.fromElements((_this).fm8(_1187_v98, p2, globalState))).Intersect(_1192_v102);
          _nw216[(new BigNumber(27)).toNumber()] = _dafny.MultiSet.fromElements(p2, p3);
          _1196_v106 = _nw216;
          let _index214 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1196_v106).length));
          (_1196_v106)[_index214] = _dafny.MultiSet.fromElements(p3, p3);
          let _index215 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1196_v106).length));
          let _rhs181 = ((p2) ? (_1191_v101) : (_1191_v101));
          let _rhs182 = true;
          let _rhs183 = _1192_v102;
          let _lhs99 = _1196_v106;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1196_v106).length));
          _1191_v101 = _rhs181;
          r0 = _rhs182;
          _lhs99[_lhs100] = _rhs183;
          let _1197_v107;
          let _nw217 = new _module.C1();
          _nw217.__ctor((_module.D6.create_DC11(_1187_v98)).dtor_cf21);
          _1197_v107 = _nw217;
          let _1198_v108;
          _1198_v108 = _module.D8.create_DC17(p2, _1197_v107, _this.f2);
          let _1199_v109;
          let _nw218 = Array((new BigNumber(20)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = p2;
          _nw218[(_dafny.ONE).toNumber()] = p3;
          _nw218[(new BigNumber(2)).toNumber()] = p3;
          _nw218[(new BigNumber(3)).toNumber()] = p2;
          _nw218[(new BigNumber(4)).toNumber()] = true;
          _nw218[(new BigNumber(5)).toNumber()] = true;
          _nw218[(new BigNumber(6)).toNumber()] = !(false);
          _nw218[(new BigNumber(7)).toNumber()] = p2;
          _nw218[(new BigNumber(8)).toNumber()] = !((_1198_v108).dtor_cf30);
          _nw218[(new BigNumber(9)).toNumber()] = p2;
          _nw218[(new BigNumber(10)).toNumber()] = p2;
          _nw218[(new BigNumber(11)).toNumber()] = p2;
          _nw218[(new BigNumber(12)).toNumber()] = p3;
          _nw218[(new BigNumber(13)).toNumber()] = p3;
          _nw218[(new BigNumber(14)).toNumber()] = p3;
          _nw218[(new BigNumber(15)).toNumber()] = false;
          _nw218[(new BigNumber(16)).toNumber()] = p3;
          _nw218[(new BigNumber(17)).toNumber()] = p2;
          _nw218[(new BigNumber(18)).toNumber()] = (_1197_v107).fm11(globalState);
          _nw218[(new BigNumber(19)).toNumber()] = false;
          _1199_v109 = _nw218;
          let _1200_v110;
          _1200_v110 = _dafny.Set.fromElements(_1199_v109);
          _1200_v110 = (_1200_v110).Intersect(_1200_v110);
          let _1201_v111;
          let _nw219 = new _module.C1();
          _nw219.__ctor(_dafny.Seq.Concat((_1197_v107).f3, _1187_v98));
          _1201_v111 = _nw219;
          _1191_v101 = _1191_v101;
        } else {
          _1114_v40 = _1114_v40;
          let _1202_v112;
          _1202_v112 = _dafny.MultiSet.fromElements(!(p2), p2);
          _1202_v112 = _1202_v112;
          let _1203_v113;
          _1203_v113 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.safeDivisionInt(_this.f2, (_dafny.ZERO).minus(new BigNumber((_1187_v98).length))));
          let _1204_v114;
          _1204_v114 = _dafny.Set.fromElements(p3, p3, true);
          _1203_v113 = (_1203_v113).update((_dafny.ZERO).minus((_this).fm1(p3, globalState)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1204_v114).length),p3)).length));
          let _1205_v115;
          let _init20 = ((_1206_v98) => function (_1207_i11) {
            return _1206_v98;
          })(_1187_v98);
          let _nw220 = Array((new BigNumber(2)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw220.length); _i0_20++) {
            _nw220[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _1205_v115 = _nw220;
          let _index216 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1205_v115).length));
          (_1205_v115)[_index216] = _dafny.Seq.UnicodeFromString("tybfst");
          let _index217 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1205_v115).length));
          (_1205_v115)[_index217] = _1187_v98;
          let _1208_v116;
          _1208_v116 = _module.D0.create_DC0(p3);
          let _1209_v117;
          _1209_v117 = _dafny.Seq.of(_1208_v116);
          _1209_v117 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), ((_1210_v116) => function (_1211_i12) {
            return _1210_v116;
          })(_1208_v116));
        }
        let _1212_v118;
        _1212_v118 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1213_v119;
        _1213_v119 = _dafny.MultiSet.fromElements(_1212_v118);
        _1213_v119 = _1213_v119;
        let _1214_v120;
        let _nw221 = Array((new BigNumber(16)).toNumber());
        _nw221[(_dafny.ZERO).toNumber()] = (_this).fm8(_1187_v98, p3, globalState);
        _nw221[(_dafny.ONE).toNumber()] = p3;
        _nw221[(new BigNumber(2)).toNumber()] = p2;
        _nw221[(new BigNumber(3)).toNumber()] = p2;
        _nw221[(new BigNumber(4)).toNumber()] = p2;
        _nw221[(new BigNumber(5)).toNumber()] = p3;
        _nw221[(new BigNumber(6)).toNumber()] = p2;
        _nw221[(new BigNumber(7)).toNumber()] = p3;
        _nw221[(new BigNumber(8)).toNumber()] = p3;
        _nw221[(new BigNumber(9)).toNumber()] = p2;
        _nw221[(new BigNumber(10)).toNumber()] = p2;
        _nw221[(new BigNumber(11)).toNumber()] = p2;
        _nw221[(new BigNumber(12)).toNumber()] = p2;
        _nw221[(new BigNumber(13)).toNumber()] = false;
        _nw221[(new BigNumber(14)).toNumber()] = p2;
        _nw221[(new BigNumber(15)).toNumber()] = p2;
        _1214_v120 = _nw221;
        let _1215_v121;
        _1215_v121 = _dafny.Set.fromElements(_1214_v120);
        let _1216_v122;
        _1216_v122 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1188_i10);
        let _1217_v123;
        _1217_v123 = _dafny.Set.fromElements(p2);
        let _1218_v124;
        _1218_v124 = _dafny.Seq.of(_1217_v123);
        let _1219_v125;
        let _nw222 = Array((new BigNumber(12)).toNumber());
        _nw222[(_dafny.ZERO).toNumber()] = _1216_v122;
        _nw222[(_dafny.ONE).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(2)).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(3)).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(4)).toNumber()] = (_1216_v122).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1188_i10));
        _nw222[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(_1188_i10));
        _nw222[(new BigNumber(6)).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(7)).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(8)).toNumber()] = _module.__default.fm27(_1188_i10, p2, globalState);
        _nw222[(new BigNumber(9)).toNumber()] = (_1216_v122).update((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1218_v124).length))), p0);
        _nw222[(new BigNumber(10)).toNumber()] = _1216_v122;
        _nw222[(new BigNumber(11)).toNumber()] = _1216_v122;
        _1219_v125 = _nw222;
        let _rhs184 = (_1215_v121).Difference(_1215_v121);
        let _rhs185 = !(!(!((p2) && (false))));
        let _rhs186 = _1219_v125;
        let _rhs187 = p3;
        let _rhs188 = (_dafny.ZERO).minus((_this).fm1(p2, globalState));
        let _lhs101 = _this;
        _1215_v121 = _rhs184;
        r0 = _rhs185;
        _1219_v125 = _rhs186;
        r0 = _rhs187;
        _lhs101.f2 = _rhs188;
        r0 = p2;
      }
      let _1220_v126;
      let _nw223 = Array((new BigNumber(19)).toNumber()).fill(false);
      _1220_v126 = _nw223;
      _1220_v126 = _1220_v126;
      let _1221_v127;
      let _nw224 = Array((new BigNumber(2)).toNumber());
      _nw224[(_dafny.ZERO).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("uokfbik")).length)).minus(p0);
      _nw224[(_dafny.ONE).toNumber()] = p0;
      _1221_v127 = _nw224;
      _1221_v127 = _1221_v127;
      r0 = ((p3) ? ((p2) || (p2)) : (p3));
      let _1222_v128;
      _1222_v128 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
      let _1223_v129;
      _1223_v129 = _dafny.Seq.of(_1222_v128);
      let _1224_v130;
      _1224_v130 = _dafny.Map.Empty.slice().updateUnsafe(true,p3);
      let _1225_v131;
      _1225_v131 = _dafny.Seq.of(_this.f2, _this.f2, p0, new BigNumber(-887));
      let _1226_v132;
      _1226_v132 = _dafny.MultiSet.fromElements(_1225_v131, _dafny.Seq.of(p0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), function (_1227_i13) {
        return new BigNumber(-392);
      }));
      let _1228_v133;
      _1228_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1226_v132,_1222_v128);
      let _1229_v134;
      let _nw225 = Array((new BigNumber(18)).toNumber());
      _nw225[(_dafny.ZERO).toNumber()] = _1222_v128;
      _nw225[(_dafny.ONE).toNumber()] = (_1222_v128).Merge((_1223_v129)[_module.__default.safeIndex(_this.f2, new BigNumber((_1223_v129).length))]);
      _nw225[(new BigNumber(2)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p3,p2)).update(p3, p3);
      _nw225[(new BigNumber(4)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
      _nw225[(new BigNumber(6)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(7)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(8)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(9)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(10)).toNumber()] = (_1222_v128).Merge(_1222_v128);
      _nw225[(new BigNumber(11)).toNumber()] = (_1222_v128).Merge(_1222_v128);
      _nw225[(new BigNumber(12)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
      _nw225[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,!(p3));
      _nw225[(new BigNumber(15)).toNumber()] = _1222_v128;
      _nw225[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).Merge((_1224_v130));
      _nw225[(new BigNumber(17)).toNumber()] = (_1222_v128).Merge((((_1228_v133).contains(_1226_v132)) ? ((_1228_v133).get(_1226_v132)) : (_1222_v128)));
      _1229_v134 = _nw225;
      r1 = _1229_v134;
      let _1230_v135;
      _1230_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1220_v126,_1187_v98);
      let _1231_v136;
      _1231_v136 = new _dafny.CodePoint('p'.codePointAt(0));
      r2 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat((((_1230_v135).contains(_1220_v126)) ? ((_1230_v135).get(_1220_v126)) : (_1187_v98)), _dafny.Seq.Concat(_1187_v98, _1187_v98)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat((((_1230_v135).contains(_1220_v126)) ? ((_1230_v135).get(_1220_v126)) : (_1187_v98)), _dafny.Seq.Concat(_1187_v98, _1187_v98))).length)), _1231_v136)).length);
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1232_v0;
      let _init21 = ((_1233_p2) => function (_1234_i0) {
        return _1233_p2;
      })(p2);
      let _nw226 = Array((new BigNumber(7)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw226.length); _i0_21++) {
        _nw226[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _1232_v0 = _nw226;
      let _index218 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
      (_1232_v0)[_index218] = p2;
      let _1235_v1;
      _1235_v1 = true;
      let _1236_v2;
      _1236_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1235_v1);
      let _1237_v4;
      _1237_v4 = _dafny.Seq.UnicodeFromString("geclc");
      let _index219 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
      let _rhs189 = _this.f2;
      let _rhs190 = (new BigNumber(((_1236_v2).Merge(((function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_1238_i1) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-31)), function (_1239_i2) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          })).length);
        })).Elements) {
          let _1240_v3 = _compr_38;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_1238_i1) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-31)), function (_1239_i2) {
              return new _dafny.CodePoint('l'.codePointAt(0));
            })).length);
          }), _1240_v3)) {
            _coll38.push([_module.__default.safeDivisionInt(_1240_v3, new BigNumber((_dafny.Seq.UnicodeFromString("oner")).length)),p2]);
          }
        }
        return _coll38;
      }()).update(p0, p2)).update(p0, p2))).length)).isLessThan(_this.f2);
      let _rhs191 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1237_v4, _1237_v4), _1237_v4);
      let _rhs192 = (_dafny.ZERO).minus(p1);
      let _lhs102 = _1232_v0;
      let _lhs103 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
      let _lhs104 = _this;
      r0 = _rhs189;
      _lhs102[_lhs103] = _rhs190;
      _1235_v1 = _rhs191;
      _lhs104.f2 = _rhs192;
      let _1241_v5;
      _1241_v5 = _dafny.Set.fromElements(_this.f2, _this.f2);
      let _1242_v6;
      _1242_v6 = _dafny.MultiSet.fromElements(_1241_v5);
      let _1243_i3;
      _1243_i3 = _dafny.ZERO;
      L5: {
        while (((((_1242_v6).contains(_1241_v5)) ? ((_1242_v6).get(_1241_v5)) : (p0))).isLessThanOrEqualTo((_this).fm1((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], globalState))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1243_i3)) {
              break L5;
            }
            _1243_i3 = (_1243_i3).plus(_dafny.ONE);
            _1237_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1237_v4, _1237_v4), _1237_v4);
            if (p2) {
              let _1244_v7;
              _1244_v7 = _dafny.Set.fromElements((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]);
              _1244_v7 = _1244_v7;
              let _1245_v8;
              let _nw227 = new _module.C1();
              _nw227.__ctor(_dafny.Seq.UnicodeFromString("oxtqqto"));
              _1245_v8 = _nw227;
              let _1246_v9;
              _1246_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p2, _1235_v1)).length),_this.f2);
              let _1247_v11;
              _1247_v11 = _dafny.Seq.of(_1246_v9);
              let _1248_v12;
              let _nw228 = Array((new BigNumber(9)).toNumber());
              _nw228[(_dafny.ZERO).toNumber()] = _1246_v9;
              _nw228[(_dafny.ONE).toNumber()] = _1246_v9;
              _nw228[(new BigNumber(2)).toNumber()] = _1246_v9;
              _nw228[(new BigNumber(3)).toNumber()] = (_1246_v9).Merge(_1246_v9);
              _nw228[(new BigNumber(4)).toNumber()] = _1246_v9;
              _nw228[(new BigNumber(5)).toNumber()] = _1246_v9;
              _nw228[(new BigNumber(6)).toNumber()] = function () {
                let _coll39 = new _dafny.Map();
                for (const _compr_39 of _dafny.IntegerRange(new BigNumber(216), new BigNumber(603))) {
                  let _1249_v10 = _compr_39;
                  if (((new BigNumber(216)).isLessThanOrEqualTo(_1249_v10)) && ((_1249_v10).isLessThan(new BigNumber(603)))) {
                    _coll39.push([(_1249_v10).plus(new BigNumber((_dafny.Seq.of(p1)).length)),p1]);
                  }
                }
                return _coll39;
              }();
              _nw228[(new BigNumber(7)).toNumber()] = (_1247_v11)[_module.__default.safeIndex(p0, new BigNumber((_1247_v11).length))];
              _nw228[(new BigNumber(8)).toNumber()] = _1246_v9;
              _1248_v12 = _nw228;
              let _index220 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1248_v12).length));
              (_1248_v12)[_index220] = _1246_v9;
              let _index221 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1248_v12).length));
              let _rhs193 = _1246_v9;
              let _rhs194 = _this.f2;
              let _lhs105 = _1248_v12;
              let _lhs106 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1248_v12).length));
              let _lhs107 = _this;
              _lhs105[_lhs106] = _rhs193;
              _lhs107.f2 = _rhs194;
              let _1250_v13;
              _1250_v13 = _dafny.Seq.of(new BigNumber(690), p0, new BigNumber((_dafny.Seq.Concat((_1245_v8).f3, _1237_v4)).length));
              _1250_v13 = _dafny.Seq.Concat(_1250_v13, _1250_v13);
              let _1251_v14;
              let _nw229 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _1251_v14 = _nw229;
              _1251_v14 = _1251_v14;
            } else {
              let _1252_v15;
              let _init22 = ((_1253_v0) => function (_1254_i4) {
                return _dafny.Map.Empty.slice().updateUnsafe((_1253_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1253_v0).length))],true);
              })(_1232_v0);
              let _nw230 = Array((new BigNumber(7)).toNumber());
              for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw230.length); _i0_22++) {
                _nw230[_i0_22] = _init22(new BigNumber(_i0_22));
              }
              _1252_v15 = _nw230;
              let _1255_v16;
              _1255_v16 = _dafny.Seq.of(_1252_v15);
              _1252_v15 = (_1255_v16)[_module.__default.safeIndex(_this.f2, new BigNumber((_1255_v16).length))];
              let _1256_v17;
              let _init23 = function (_1257_i5) {
                return (_1257_i5).minus(new BigNumber(771));
              };
              let _nw231 = Array((new BigNumber(14)).toNumber());
              for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw231.length); _i0_23++) {
                _nw231[_i0_23] = _init23(new BigNumber(_i0_23));
              }
              _1256_v17 = _nw231;
              let _1258_v18;
              _1258_v18 = _dafny.Seq.of(_1256_v17);
              let _1259_v19;
              _1259_v19 = new _dafny.CodePoint('i'.codePointAt(0));
              let _1260_v20;
              _1260_v20 = _module.D4.create_DC6(p1, _1256_v17, p0);
              let _1261_v21;
              let _nw232 = Array((new BigNumber(20)).toNumber());
              _nw232[(_dafny.ZERO).toNumber()] = _1256_v17;
              _nw232[(_dafny.ONE).toNumber()] = ((p2) ? (_1256_v17) : (_1256_v17));
              _nw232[(new BigNumber(2)).toNumber()] = (((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]) ? (_1256_v17) : ((_1258_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("udwhy")).length), new BigNumber((_1258_v18).length))]));
              _nw232[(new BigNumber(3)).toNumber()] = (_1258_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1237_v4, _module.__default.safeIndex(p1, new BigNumber((_1237_v4).length)), _1259_v19), _module.__default.safeIndex(_this.f2, new BigNumber((_dafny.Seq.update(_1237_v4, _module.__default.safeIndex(p1, new BigNumber((_1237_v4).length)), _1259_v19)).length)), _1259_v19)).length), new BigNumber((_1258_v18).length))];
              _nw232[(new BigNumber(4)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(5)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(6)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(7)).toNumber()] = ((!(true)) ? (_1256_v17) : (_1256_v17));
              _nw232[(new BigNumber(8)).toNumber()] = (_1260_v20).dtor_cf10;
              _nw232[(new BigNumber(9)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(10)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(11)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(12)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(13)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(14)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(15)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(16)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(17)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(18)).toNumber()] = _1256_v17;
              _nw232[(new BigNumber(19)).toNumber()] = _1256_v17;
              _1261_v21 = _nw232;
              let _index222 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1261_v21).length));
              (_1261_v21)[_index222] = ((_1235_v1) ? (_1256_v17) : (_1256_v17));
              let _index223 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1261_v21).length));
              (_1261_v21)[_index223] = (_1260_v20).dtor_cf10;
              let _1262_v22;
              _1262_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1237_v4);
              let _1263_v23;
              _1263_v23 = _dafny.Seq.of(_this.f2, _this.f2);
              let _1264_v24;
              _1264_v24 = _dafny.Seq.of(_1235_v1, _1235_v1);
              let _1265_v25;
              _1265_v25 = _dafny.Seq.of((_1263_v23)[_module.__default.safeIndex(new BigNumber((_1264_v24).length), new BigNumber((_1263_v23).length))]);
              let _1266_v26;
              _1266_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
              _1237_v4 = (((_1262_v22).contains(p1)) ? ((_1262_v22).get(p1)) : (_dafny.Seq.update(_1237_v4, _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1265_v25, _1265_v25, _dafny.Seq.of(new BigNumber((_1266_v26).length)), _1265_v25)).length), new BigNumber((_1237_v4).length)), _1259_v19)));
              let _1267_v27;
              let _init24 = ((_1268_v23, _1269_p1) => function (_1270_i6) {
                return _dafny.Seq.update(_dafny.Seq.Concat(_1268_v23, _dafny.Seq.update(_1268_v23, _module.__default.safeIndex(_1269_p1, new BigNumber((_1268_v23).length)), _this.f2)), _module.__default.safeIndex(new BigNumber((_1268_v23).length), new BigNumber((_dafny.Seq.Concat(_1268_v23, _dafny.Seq.update(_1268_v23, _module.__default.safeIndex(_1269_p1, new BigNumber((_1268_v23).length)), _this.f2))).length)), _this.f2);
              })(_1263_v23, p1);
              let _nw233 = Array((new BigNumber(19)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw233.length); _i0_24++) {
                _nw233[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _1267_v27 = _nw233;
              let _index224 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1267_v27).length));
              (_1267_v27)[_index224] = _1263_v23;
              let _1271_v28;
              _1271_v28 = _dafny.Set.fromElements(false);
              let _index225 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1267_v27).length));
              (_1267_v27)[_index225] = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_1271_v28, _dafny.Set.fromElements(p2))).length));
              let _1272_v29;
              let _nw234 = new _module.C2();
              _nw234.__ctor(_this.f2);
              _1272_v29 = _nw234;
            }
            let _index226 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
            (_1232_v0)[_index226] = p2;
            let _index227 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
            (_1232_v0)[_index227] = _1235_v1;
          }
        }
      }
      if (_1235_v1) {
        _1235_v1 = _1235_v1;
        r0 = (_dafny.ZERO).minus(_this.f2);
        r0 = p0;
        (_this).f2 = new BigNumber((_1237_v4).length);
        let _1273_v30;
        _1273_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))],new BigNumber(-750));
        let _index228 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        (_1232_v0)[_index228] = (new BigNumber((_1273_v30).length)).isLessThan(new BigNumber(-831));
      } else {
        let _1274_v31;
        _1274_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v1,(_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]);
        let _1275_v32;
        _1275_v32 = _dafny.Seq.of(new BigNumber(-366), _this.f2, _module.__default.safeModuloInt(new BigNumber((_1274_v31).length), new BigNumber((_1236_v2).length)), p0);
        _1275_v32 = _1275_v32;
        r0 = (_this).fm1(p2, globalState);
        let _1276_v33;
        _1276_v33 = _dafny.Set.fromElements(p2, _1235_v1);
        let _1277_v34;
        _1277_v34 = _dafny.Seq.of(_dafny.Set.fromElements(true, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]), _1276_v33, _1276_v33, _1276_v33);
        let _1278_v35;
        _1278_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]) ? (p0) : (new BigNumber(((_1277_v34)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1277_v34).length))]).length))));
        _1278_v35 = (_1278_v35).update(p0, p0);
        let _1279_v36;
        _1279_v36 = _dafny.MultiSet.fromElements((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], ((true) ? (p2) : (_1235_v1)), _1235_v1);
        let _rhs195 = (((_1279_v36).contains((_1241_v5).IsSubsetOf(_1241_v5))) ? ((_1279_v36).get((_1241_v5).IsSubsetOf(_1241_v5))) : (((_dafny.ZERO).minus(new BigNumber((_1237_v4).length))).multipliedBy(p1)));
        let _rhs196 = _module.__default.safeModuloInt(new BigNumber((_1276_v33).length), (_this.f2).multipliedBy(p0));
        let _lhs108 = _this;
        let _lhs109 = _this;
        _lhs108.f2 = _rhs195;
        _lhs109.f2 = _rhs196;
        let _1280_v37;
        let _init25 = ((_1281_p0) => function (_1282_i7) {
          return (_1282_i7).plus(_1281_p0);
        })(p0);
        let _nw235 = Array((new BigNumber(25)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw235.length); _i0_25++) {
          _nw235[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1280_v37 = _nw235;
        let _1283_v38;
        _1283_v38 = _dafny.MultiSet.fromElements(new BigNumber(993), _this.f2);
        let _1284_v39;
        _1284_v39 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index229 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _rhs197 = _1235_v1;
        let _rhs198 = (_this).fm8(_dafny.Seq.update(_1237_v4, _module.__default.safeIndex(p1, new BigNumber((_1237_v4).length)), _1284_v39), (_1279_v36).IsProperSubsetOf(_1279_v36), globalState);
        let _rhs199 = _1280_v37;
        let _rhs200 = ((_1283_v38).Union(_1283_v38)).Union((_1283_v38).Union(_1283_v38));
        let _rhs201 = _dafny.Seq.Concat(_1275_v32, _1275_v32);
        let _lhs110 = _1232_v0;
        let _lhs111 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        _lhs110[_lhs111] = _rhs197;
        _1235_v1 = _rhs198;
        _1280_v37 = _rhs199;
        _1283_v38 = _rhs200;
        _1275_v32 = _rhs201;
      }
      let _1285_v40;
      _1285_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(486));
      let _1286_v41;
      _1286_v41 = _dafny.Set.fromElements(_1285_v40);
      let _hi5 = (p1).minus(p1);
      for (let _1287_i8 = new BigNumber(((_dafny.Set.fromElements((_1285_v40).update(p2, p1), _1285_v40, _1285_v40)).Difference(_1286_v41)).length); _1287_i8.isLessThan(_hi5); _1287_i8 = _1287_i8.plus(_dafny.ONE)) {
        let _1288_v42;
        _1288_v42 = _dafny.Set.fromElements((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]);
        if ((_1288_v42).IsSubsetOf(_1288_v42)) {
          let _1289_v43;
          let _nw236 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1289_v43 = _nw236;
          let _1290_v44;
          _1290_v44 = _dafny.Set.fromElements(_1289_v43, _1289_v43, _1289_v43);
          _1290_v44 = (((false) ? (_dafny.Set.fromElements(_1289_v43)) : (_1290_v44))).Intersect((_1290_v44).Union(_1290_v44));
          let _1291_v45;
          let _init26 = ((_1292_v2) => function (_1293_i9) {
            return _1292_v2;
          })(_1236_v2);
          let _nw237 = Array((new BigNumber(10)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw237.length); _i0_26++) {
            _nw237[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1291_v45 = _nw237;
          let _index230 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1291_v45).length));
          (_1291_v45)[_index230] = _1236_v2;
          let _1294_v46;
          _1294_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v4,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1241_v5).length),(_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]));
          let _1295_v47;
          _1295_v47 = _module.D5.create_DC10(_1235_v1, p0, _1294_v46, _1235_v1);
          let _1296_v48;
          _1296_v48 = _dafny.Seq.of(_1235_v1);
          let _1297_v49;
          _1297_v49 = _dafny.MultiSet.fromElements(_1235_v1, (_1296_v48)[_module.__default.safeIndex(new BigNumber((_1296_v48).length), new BigNumber((_1296_v48).length))]);
          let _1298_v50;
          _1298_v50 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1299_v51;
          _1299_v51 = _module.D0.create_DC1(p2, p2, _1298_v50, p1);
          let _index231 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1291_v45).length));
          let _rhs202 = new BigNumber(-912);
          let _rhs203 = (((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]) ? (((_dafny.Map.Empty.slice().updateUnsafe(_1235_v1,new BigNumber((_1297_v49).cardinality()))).update((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], new BigNumber((_1241_v5).length))).Merge(_1285_v40)) : (_1285_v40));
          let _rhs204 = ((_1236_v2).Merge(_1236_v2)).update((_this).fm1((_this).fm8(_dafny.Seq.UnicodeFromString("o"), false, globalState), globalState), (_1299_v51).dtor_cf1);
          let _rhs205 = _1295_v47;
          let _lhs112 = _this;
          let _lhs113 = _1291_v45;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1291_v45).length));
          _lhs112.f2 = _rhs202;
          _1285_v40 = _rhs203;
          _lhs113[_lhs114] = _rhs204;
          _1295_v47 = _rhs205;
          let _1300_v52;
          let _nw238 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1300_v52 = _nw238;
          let _index232 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_1300_v52).length));
          (_1300_v52)[_index232] = _1237_v4;
          let _index233 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_1300_v52).length));
          (_1300_v52)[_index233] = _1237_v4;
          let _1301_v53;
          let _nw239 = new _module.C1();
          _nw239.__ctor(_1237_v4);
          _1301_v53 = _nw239;
          _1301_v53 = _1301_v53;
          let _1302_v54;
          _1302_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1296_v48);
          let _1303_v55;
          _1303_v55 = _dafny.Seq.of(_dafny.Seq.update(_1296_v48, _module.__default.safeIndex(_this.f2, new BigNumber((_1296_v48).length)), false), _dafny.Seq.update(_dafny.Seq.update(_1296_v48, _module.__default.safeIndex(p1, new BigNumber((_1296_v48).length)), _1235_v1), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(_1296_v48, _module.__default.safeIndex(p1, new BigNumber((_1296_v48).length)), _1235_v1)).length)), !(true)), _1296_v48, _dafny.Seq.Concat((((_1302_v54).contains(new BigNumber(-989))) ? ((_1302_v54).get(new BigNumber(-989))) : (_1296_v48)), _dafny.Seq.update(_1296_v48, _module.__default.safeIndex(p0, new BigNumber((_1296_v48).length)), false)), _1296_v48);
          let _1304_v56;
          _1304_v56 = _dafny.MultiSet.fromElements(new BigNumber((_1288_v42).length), _this.f2);
          let _1305_v57;
          _1305_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v1,_1235_v1);
          let _1306_v58;
          _1306_v58 = _module.D8.create_DC17((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], _1301_v53, new BigNumber((_1305_v57).length));
          _1296_v48 = (_1303_v55)[_module.__default.safeIndex((((_1304_v56).contains((_1306_v58).dtor_cf32)) ? ((_1304_v56).get((_1306_v58).dtor_cf32)) : (p0)), new BigNumber((_1303_v55).length))];
        } else {
          let _1307_v59;
          _1307_v59 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qsrqrrepm"));
          let _index234 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
          (_1232_v0)[_index234] = (_module.__default.fm30(_1235_v1, globalState)).IsSubsetOf((_1307_v59).Union(_1307_v59));
          let _1308_v60;
          let _init27 = ((_1309_v4) => function (_1310_i10) {
            return (_1310_i10).minus(new BigNumber((_1309_v4).length));
          })(_1237_v4);
          let _nw240 = Array((new BigNumber(23)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw240.length); _i0_27++) {
            _nw240[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1308_v60 = _nw240;
          _1308_v60 = _1308_v60;
          let _1311_v61;
          _1311_v61 = _dafny.MultiSet.fromElements(_this.f2);
          let _1312_v62;
          _1312_v62 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1313_v63;
          _1313_v63 = _dafny.MultiSet.fromElements(_1235_v1);
          let _1314_v64;
          _1314_v64 = _module.D0.create_DC1((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], _1235_v1, _1312_v62, p0);
          let _1315_v65;
          _1315_v65 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),_dafny.MultiSet.fromElements((_1314_v64).dtor_cf4));
          let _1316_v66;
          _1316_v66 = _1311_v61;
          let _1317_v67;
          _1317_v67 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f2)));
          let _1318_v68;
          let _nw241 = Array((new BigNumber(25)).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = (_1311_v61).Intersect(_1311_v61);
          _nw241[(_dafny.ONE).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(_this.f2, p1)).Difference(_1311_v61);
          _nw241[(new BigNumber(3)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(4)).toNumber()] = (_module.__default.fm35(_1312_v62, _1235_v1, _1313_v63, globalState)).Intersect(_1311_v61);
          _nw241[(new BigNumber(5)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_1319_i11) {
            return new BigNumber(924);
          }))).Union(_1311_v61);
          _nw241[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(p0, p0, p0, p0, _this.f2);
          _nw241[(new BigNumber(8)).toNumber()] = _module.__default.fm19(new BigNumber(712), p0, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], _this.f2, globalState);
          _nw241[(new BigNumber(9)).toNumber()] = (_1311_v61).Difference((((_1315_v65).contains(_1312_v62)) ? ((_1315_v65).get(_1312_v62)) : (_dafny.MultiSet.fromElements(new BigNumber((_1236_v2).length), p1))));
          _nw241[(new BigNumber(10)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(11)).toNumber()] = (_1311_v61).Intersect(_1311_v61);
          _nw241[(new BigNumber(12)).toNumber()] = ((_1311_v61).update(p1, _module.__default.abs(p1))).Intersect((_1316_v66));
          _nw241[(new BigNumber(13)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(14)).toNumber()] = (_1317_v67)[_module.__default.safeIndex(p0, new BigNumber((_1317_v67).length))];
          _nw241[(new BigNumber(15)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(16)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(17)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(18)).toNumber()] = ((true) ? (_dafny.MultiSet.fromElements((_this).fm1(p2, globalState))) : (_1311_v61));
          _nw241[(new BigNumber(19)).toNumber()] = _module.__default.fm35(_1312_v62, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], _1313_v63, globalState);
          _nw241[(new BigNumber(20)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(21)).toNumber()] = (_1311_v61).update(_1287_i8, _module.__default.abs(new BigNumber(55)));
          _nw241[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements(_1287_i8);
          _nw241[(new BigNumber(23)).toNumber()] = _1311_v61;
          _nw241[(new BigNumber(24)).toNumber()] = _1311_v61;
          _1318_v68 = _nw241;
          let _index235 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1318_v68).length));
          (_1318_v68)[_index235] = (_1311_v61).Difference(_1311_v61);
          let _index236 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1318_v68).length));
          (_1318_v68)[_index236] = _1311_v61;
          _1312_v62 = _1312_v62;
          let _index237 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
          (_1232_v0)[_index237] = (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))];
        }
        let _1320_v69;
        _1320_v69 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1321_v71;
        let _init28 = ((_1322_p0, _1323_p1) => function (_1324_i12) {
          return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(895), new BigNumber(618))) {
              let _1325_v70 = _compr_40;
              if (((new BigNumber(895)).isLessThanOrEqualTo(_1325_v70)) && ((_1325_v70).isLessThan(new BigNumber(618)))) {
                _coll40.push([(_1325_v70).multipliedBy(new BigNumber(190)),_dafny.Map.Empty.slice().updateUnsafe(_1323_p1,_this.f2)]);
              }
            }
            return _coll40;
          }()).length),_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1323_p1),_1322_p0))).length),_1322_p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f2,_this.f2));
        })(p0, p1);
        let _nw242 = Array((new BigNumber(25)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw242.length); _i0_28++) {
          _nw242[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1321_v71 = _nw242;
        let _1326_v73;
        _1326_v73 = _dafny.MultiSet.fromElements(new BigNumber((_1237_v4).length));
        let _1327_v74;
        _1327_v74 = _dafny.Map.Empty.slice().updateUnsafe((((_1326_v73).contains(_1287_i8)) ? ((_1326_v73).get(_1287_i8)) : (new BigNumber(47))),_1287_i8);
        let _1328_v75;
        _1328_v75 = _dafny.Seq.of(_1327_v74);
        let _index238 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1321_v71).length));
        (_1321_v71)[_index238] = function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of ((_1328_v75)[_module.__default.safeIndex(_this.f2, new BigNumber((_1328_v75).length))]).Keys.Elements) {
            let _1329_v72 = _compr_41;
            if (((_1328_v75)[_module.__default.safeIndex(_this.f2, new BigNumber((_1328_v75).length))]).contains(_1329_v72)) {
              _coll41.push([(_1329_v72).plus((_dafny.ZERO).minus(_this.f2)),p1]);
            }
          }
          return _coll41;
        }();
        let _1330_v76;
        _1330_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))],_1235_v1);
        let _index239 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1321_v71).length));
        let _index240 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _rhs206 = _module.__default.safeDivisionInt(new BigNumber((_1330_v76).length), p1);
        let _rhs207 = new _dafny.CodePoint('u'.codePointAt(0));
        let _rhs208 = function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(105), new BigNumber(303))) {
            let _1331_v77 = _compr_42;
            if (((new BigNumber(105)).isLessThanOrEqualTo(_1331_v77)) && ((_1331_v77).isLessThan(new BigNumber(303)))) {
              _coll42.push([(_1331_v77).multipliedBy(new BigNumber(((_1330_v76).update(p2, _1235_v1)).length)),_1287_i8]);
            }
          }
          return _coll42;
        }();
        let _rhs209 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_1332_i13) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(169))).length));
        })).length)).isEqualTo(_this.f2);
        let _rhs210 = false;
        let _lhs115 = _1321_v71;
        let _lhs116 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1321_v71).length));
        let _lhs117 = _1232_v0;
        let _lhs118 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        r0 = _rhs206;
        _1320_v69 = _rhs207;
        _lhs115[_lhs116] = _rhs208;
        _lhs117[_lhs118] = _rhs209;
        _1235_v1 = _rhs210;
        let _1333_v78;
        _1333_v78 = _dafny.Seq.of(_1235_v1);
        let _index241 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        (_1232_v0)[_index241] = ((p2) ? ((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]) : ((_1333_v78)[_module.__default.safeIndex((((_1285_v40).contains((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))])) ? ((_1285_v40).get((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))])) : (_1287_i8)), new BigNumber((_1333_v78).length))]));
        let _1334_v79;
        _1334_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v4,new BigNumber(((_1288_v42).Intersect(_1288_v42)).length));
        _1334_v79 = (_1334_v79).update(_1237_v4, new BigNumber(255));
      }
      (_this).f2 = p1;
      let _1335_v80;
      _1335_v80 = _dafny.Seq.of(p0, _this.f2, p1);
      let _1336_v81;
      _1336_v81 = _dafny.Seq.of(_1335_v80);
      if (_dafny.Seq.IsProperPrefixOf(_1336_v81, _1336_v81)) {
        _1235_v1 = p2;
        let _1337_v82;
        _1337_v82 = new _dafny.CodePoint('k'.codePointAt(0));
        r0 = new BigNumber((_dafny.Seq.update(_1237_v4, _module.__default.safeIndex(_this.f2, new BigNumber((_1237_v4).length)), _1337_v82)).length);
        let _1338_v83;
        let _nw243 = Array((new BigNumber(13)).toNumber()).fill([]);
        _1338_v83 = _nw243;
        let _1339_v84;
        _1339_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_1237_v4, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], globalState),_1235_v1);
        let _1340_v85;
        _1340_v85 = _1339_v84;
        let _1341_v86;
        _1341_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1340_v85);
        let _1342_v87;
        _1342_v87 = _dafny.Seq.of(_1341_v86);
        let _1343_v88;
        let _nw244 = Array((new BigNumber(10)).toNumber());
        _nw244[(_dafny.ZERO).toNumber()] = (_1342_v87)[_module.__default.safeIndex(p1, new BigNumber((_1342_v87).length))];
        _nw244[(_dafny.ONE).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(2)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(3)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(4)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(5)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(6)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(7)).toNumber()] = _module.__default.fm39(globalState);
        _nw244[(new BigNumber(8)).toNumber()] = _1341_v86;
        _nw244[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(449),_1339_v84);
        _1343_v88 = _nw244;
        let _index242 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1338_v83).length));
        (_1338_v83)[_index242] = _1343_v88;
        let _1344_v89;
        _1344_v89 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1241_v5);
        let _index243 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1338_v83).length));
        let _index244 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _index245 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _index246 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _rhs211 = _1343_v88;
        let _rhs212 = (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))];
        let _rhs213 = ((_1241_v5).Intersect(_1241_v5)).IsProperSubsetOf(((((_1344_v89).contains(_1235_v1)) ? ((_1344_v89).get(_1235_v1)) : (function () {
          let _coll43 = new _dafny.Set();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(56), new BigNumber(123))) {
            let _1345_v90 = _compr_43;
            if (((new BigNumber(56)).isLessThanOrEqualTo(_1345_v90)) && ((_1345_v90).isLessThan(new BigNumber(123)))) {
              _coll43.add(_module.__default.safeDivisionInt(_1345_v90, new BigNumber((_1237_v4).length)));
            }
          }
          return _coll43;
        }()))).Difference(_1241_v5));
        let _rhs214 = !(!(_1236_v2).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f2))));
        let _lhs119 = _1338_v83;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1338_v83).length));
        let _lhs121 = _1232_v0;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _lhs123 = _1232_v0;
        let _lhs124 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        let _lhs125 = _1232_v0;
        let _lhs126 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        _lhs119[_lhs120] = _rhs211;
        _lhs121[_lhs122] = _rhs212;
        _lhs123[_lhs124] = _rhs213;
        _lhs125[_lhs126] = _rhs214;
        _1235_v1 = (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))];
        let _1346_v91;
        let _nw245 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1346_v91 = _nw245;
        let _1347_v92;
        _1347_v92 = _dafny.MultiSet.fromElements(_this.f2);
        let _1348_v93;
        _1348_v93 = _1347_v92;
        let _index247 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1346_v91).length));
        (_1346_v91)[_index247] = _1348_v93;
        let _index248 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1346_v91).length));
        (_1346_v91)[_index248] = _1347_v92;
      } else {
        if ((((p2) ? (p0) : (p0))).isLessThanOrEqualTo((p0).multipliedBy(p0))) {
          r0 = p1;
          let _1349_v94;
          let _nw246 = new _module.C0();
          _nw246.__ctor(p2, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))]);
          _1349_v94 = _nw246;
          let _1350_v95;
          _1350_v95 = _dafny.Seq.of(_1237_v4);
          let _1351_v96;
          _1351_v96 = _dafny.Set.fromElements(_1237_v4, (_1350_v95)[_module.__default.safeIndex(_this.f2, new BigNumber((_1350_v95).length))]);
          _1351_v96 = _1351_v96;
          let _1352_v97;
          _1352_v97 = _dafny.MultiSet.fromElements(_1349_v94.f4);
          let _1353_v98;
          _1353_v98 = _dafny.Seq.of(_1352_v97);
          let _1354_v99;
          let _nw247 = Array((new BigNumber(15)).toNumber());
          _nw247[(_dafny.ZERO).toNumber()] = new BigNumber(746);
          _nw247[(_dafny.ONE).toNumber()] = p0;
          _nw247[(new BigNumber(2)).toNumber()] = p0;
          _nw247[(new BigNumber(3)).toNumber()] = new BigNumber(-554);
          _nw247[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw247[(new BigNumber(5)).toNumber()] = p1;
          _nw247[(new BigNumber(6)).toNumber()] = _this.f2;
          _nw247[(new BigNumber(7)).toNumber()] = new BigNumber(203);
          _nw247[(new BigNumber(8)).toNumber()] = p0;
          _nw247[(new BigNumber(9)).toNumber()] = _this.f2;
          _nw247[(new BigNumber(10)).toNumber()] = p0;
          _nw247[(new BigNumber(11)).toNumber()] = _this.f2;
          _nw247[(new BigNumber(12)).toNumber()] = _this.f2;
          _nw247[(new BigNumber(13)).toNumber()] = p0;
          _nw247[(new BigNumber(14)).toNumber()] = new BigNumber(-538);
          _1354_v99 = _nw247;
          let _1355_v100;
          _1355_v100 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pcgsm"),_1354_v99);
          let _1356_v101;
          _1356_v101 = new _dafny.CodePoint('q'.codePointAt(0));
          r0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1353_v98, _1353_v98), _dafny.Seq.Concat(_module.__default.fm40(new BigNumber((_1355_v100).length), _1356_v101, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], p1, globalState), _1353_v98))).length);
          (_1349_v94).f4 = (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))];
        } else {
          _1235_v1 = !(((_this).fm1((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], globalState)).isLessThan(new BigNumber(861)));
          _1241_v5 = _1241_v5;
          r0 = (_module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(p0))).plus((_dafny.ZERO).minus(_this.f2));
          let _index249 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
          (_1232_v0)[_index249] = p2;
          _1235_v1 = p2;
        }
        let _1357_v102;
        _1357_v102 = _dafny.Seq.of(_1237_v4);
        let _1358_v104;
        _1358_v104 = _dafny.Set.fromElements(_1237_v4);
        let _index250 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length));
        (_1232_v0)[_index250] = (function () {
          let _coll44 = new _dafny.Set();
          for (const _compr_44 of (_1357_v102).Elements) {
            let _1359_v103 = _compr_44;
            if (_dafny.Seq.contains(_1357_v102, _1359_v103)) {
              _coll44.add(_1359_v103);
            }
          }
          return _coll44;
        }()).equals(_1358_v104);
        let _1360_v105;
        _1360_v105 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, new BigNumber(167)),new BigNumber((_1335_v80).length));
        let _1361_v106;
        _1361_v106 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1360_v105);
        let _1362_v107;
        _1362_v107 = _dafny.MultiSet.fromElements((_1361_v106).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(685),_1360_v105)), _1361_v106);
        let _pat_let_tv28 = _1335_v80;
        _1362_v107 = _dafny.MultiSet.fromElements(_1361_v106, _1361_v106, function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(517), new BigNumber(472))) {
            let _1363_v108 = _compr_45;
            if (((new BigNumber(517)).isLessThanOrEqualTo(_1363_v108)) && ((_1363_v108).isLessThan(new BigNumber(472)))) {
              _coll45.push([(_1363_v108).multipliedBy((_1335_v80)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(p2))).length), new BigNumber((_1335_v80).length))]),function (_pat_let26_0) {
                return function (_1364_dt__update__tmp_h1) {
                  return function (_pat_let27_0) {
                    return function (_1365_dt__update_hcf6_h0) {
                      return _1365_dt__update_hcf6_h0;
                    }(_pat_let27_0);
                  }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv28,_this.f2));
                }(_pat_let26_0);
              }(_1360_v105)]);
            }
          }
          return _coll45;
        }(), _1361_v106, _1361_v106);
        let _1366_v109;
        let _nw248 = Array((new BigNumber(6)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = p0;
        _nw248[(_dafny.ONE).toNumber()] = _this.f2;
        _nw248[(new BigNumber(2)).toNumber()] = (_this).fm1((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], globalState);
        _nw248[(new BigNumber(3)).toNumber()] = ((_module.D6.create_DC13(p2, _this.f2)).dtor_cf25).plus(new BigNumber(825));
        _nw248[(new BigNumber(4)).toNumber()] = p1;
        _nw248[(new BigNumber(5)).toNumber()] = (((_1285_v40).contains(p2)) ? ((_1285_v40).get(p2)) : (p1));
        _1366_v109 = _nw248;
        let _index251 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1366_v109).length));
        (_1366_v109)[_index251] = p1;
        let _index252 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_1366_v109).length));
        (_1366_v109)[_index252] = ((!(p2)) ? (p0) : (p0));
        let _1367_v110;
        _1367_v110 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1368_v111;
        let _nw249 = Array((new BigNumber(9)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = _1367_v110;
        _nw249[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw249[(new BigNumber(2)).toNumber()] = _module.__default.fm37(p1, globalState);
        _nw249[(new BigNumber(3)).toNumber()] = _1367_v110;
        _nw249[(new BigNumber(4)).toNumber()] = _1367_v110;
        _nw249[(new BigNumber(5)).toNumber()] = _1367_v110;
        _nw249[(new BigNumber(6)).toNumber()] = (_module.D0.create_DC1(p2, (_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], _1367_v110, _this.f2)).dtor_cf3;
        _nw249[(new BigNumber(7)).toNumber()] = _1367_v110;
        _nw249[(new BigNumber(8)).toNumber()] = _1367_v110;
        _1368_v111 = _nw249;
        let _index253 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_1368_v111).length));
        (_1368_v111)[_index253] = new _dafny.CodePoint('b'.codePointAt(0));
        let _index254 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_1368_v111).length));
        (_1368_v111)[_index254] = _1367_v110;
      }
      r0 = (p0).multipliedBy((_1335_v80)[_module.__default.safeIndex((_this).fm1((_1232_v0)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1232_v0).length))], globalState), new BigNumber((_1335_v80).length))]);
      return r0;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let _1369_v0;
      let _nw250 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1369_v0 = _nw250;
      _1369_v0 = _1369_v0;
      let _1370_v1;
      _1370_v1 = _dafny.Seq.of(_this.f2);
      let _hi6 = (_1370_v1)[_module.__default.safeIndex(p0, new BigNumber((_1370_v1).length))];
      for (let _1371_i0 = p0; _1371_i0.isLessThan(_hi6); _1371_i0 = _1371_i0.plus(_dafny.ONE)) {
        let _1372_v2;
        _1372_v2 = true;
        if (((_1372_v2) ? (_1372_v2) : ((_1371_i0).isLessThanOrEqualTo(_this.f2)))) {
          let _1373_v3;
          _1373_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v2,_1372_v2);
          let _1374_v4;
          _1374_v4 = _dafny.Seq.of(_1372_v2, !((((_1373_v3).contains(_1372_v2)) ? ((_1373_v3).get(_1372_v2)) : (_1372_v2))), _1372_v2);
          _1373_v3 = (_1373_v3).update(_1372_v2, (_1374_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f2), new BigNumber((_1374_v4).length))]);
          (_this).f2 = p0;
          let _1375_v5;
          let _nw251 = new _module.C0();
          _nw251.__ctor(_1372_v2, _1372_v2);
          _1375_v5 = _nw251;
          let _1376_v6;
          _1376_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1375_v5.f5,(_this).fm1(!(_1375_v5.f4), globalState));
          (_this).f2 = (_dafny.ZERO).minus((p0).plus(new BigNumber(((_1376_v6).Merge(_1376_v6)).length)));
          let _1377_v7;
          let _init29 = ((_1378_v4) => function (_1379_i1) {
            return _module.__default.safeModuloInt(_1379_i1, new BigNumber((_1378_v4).length));
          })(_1374_v4);
          let _nw252 = Array((new BigNumber(15)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw252.length); _i0_29++) {
            _nw252[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1377_v7 = _nw252;
          let _1380_v8;
          _1380_v8 = _module.D4.create_DC6((_dafny.ZERO).minus((_this).fm1(_1375_v5.f5, globalState)), _1377_v7, p0);
          let _index255 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1377_v7).length));
          (_1377_v7)[_index255] = (_1380_v8).dtor_cf9;
          let _1381_v9;
          let _nw253 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _1381_v9 = _nw253;
          let _index256 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1381_v9).length));
          (_1381_v9)[_index256] = _1376_v6;
          let _1382_v10;
          _1382_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1371_i0,_1375_v5.f5);
          let _1383_v11;
          _1383_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f2,_1371_i0),_dafny.Seq.Create(_module.__default.abs(new BigNumber(505)), function (_1384_i2) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }));
          let _1385_v12;
          _1385_v12 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1386_v13;
          _1386_v13 = _module.D7.create_DC15(_this.f2, _1385_v12);
          let _1387_v14;
          _1387_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1383_v11).length),_1386_v13);
          let _index257 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1377_v7).length));
          let _index258 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1381_v9).length));
          let _rhs215 = _this.f2;
          let _rhs216 = ((_1375_v5.f5) ? ((new BigNumber((_1382_v10).length)).minus(new BigNumber(-727))) : (new BigNumber(((_1387_v14).Merge(function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(183), new BigNumber(417))) {
              let _1388_v15 = _compr_46;
              if (((new BigNumber(183)).isLessThanOrEqualTo(_1388_v15)) && ((_1388_v15).isLessThan(new BigNumber(417)))) {
                _coll46.push([(_1388_v15).plus(p0),_1386_v13]);
              }
            }
            return _coll46;
          }())).length)));
          let _rhs217 = _module.__default.safeModuloInt(p0, _this.f2);
          let _rhs218 = ((_1376_v6).Merge(_1376_v6)).Merge(_1376_v6);
          let _rhs219 = _dafny.Seq.Concat(_1370_v1, _dafny.Seq.Concat(_1370_v1, _1370_v1));
          let _lhs127 = _this;
          let _lhs128 = _1377_v7;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1377_v7).length));
          let _lhs130 = _1381_v9;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1381_v9).length));
          r0 = _rhs215;
          _lhs127.f2 = _rhs216;
          _lhs128[_lhs129] = _rhs217;
          _lhs130[_lhs131] = _rhs218;
          _1370_v1 = _rhs219;
        } else {
          let _1389_v16;
          let _nw254 = Array((new BigNumber(18)).toNumber()).fill([]);
          _1389_v16 = _nw254;
          let _1390_v17;
          let _nw255 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1390_v17 = _nw255;
          let _index259 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1389_v16).length));
          (_1389_v16)[_index259] = _1390_v17;
          let _index260 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1389_v16).length));
          (_1389_v16)[_index260] = _1390_v17;
          let _1391_v18;
          let _nw256 = new _module.C1();
          _nw256.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-92)), function (_1392_i3) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }));
          _1391_v18 = _nw256;
          _1391_v18 = _1391_v18;
          let _1393_v19;
          _1393_v19 = _dafny.Set.fromElements(_1371_i0);
          let _1394_v20;
          _1394_v20 = _dafny.Seq.of(_1393_v19);
          let _1395_v21;
          _1395_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_1394_v20, _1394_v20),_1372_v2);
          _1395_v21 = (_1395_v21).update(!(_1372_v2), _1372_v2);
          let _1396_v22;
          _1396_v22 = _dafny.Seq.UnicodeFromString("hwqgaeod");
          _1396_v22 = _1396_v22;
          let _1397_v23;
          _1397_v23 = _dafny.Seq.of(_1372_v2);
          let _index261 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1390_v17).length));
          (_1390_v17)[_index261] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(_1372_v2), _1372_v2), _1397_v23)).length);
          let _index262 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1390_v17).length));
          (_1390_v17)[_index262] = ((_1372_v2) ? ((_1371_i0).multipliedBy(new BigNumber(867))) : (_this.f2));
        }
        let _1398_v24;
        _1398_v24 = new _dafny.CodePoint('s'.codePointAt(0));
        _1398_v24 = _1398_v24;
        let _1399_v25;
        _1399_v25 = _dafny.Seq.UnicodeFromString("ednpivdn");
        let _1400_v26;
        _1400_v26 = _dafny.Seq.of(_1399_v25, _1399_v25);
        _1400_v26 = _1400_v26;
        (_this).f2 = _1371_i0;
      }
      let _1401_v27;
      _1401_v27 = _dafny.Seq.UnicodeFromString("wosyavxm");
      let _1402_v28;
      _1402_v28 = _module.D6.create_DC11(_1401_v27);
      r0 = (((_1370_v1)[_module.__default.safeIndex(p0, new BigNumber((_1370_v1).length))]).minus(p0)).minus((_dafny.ZERO).minus(new BigNumber(((_1402_v28).dtor_cf21).length)));
      let _1403_v29;
      _1403_v29 = true;
      r0 = ((_1403_v29) ? ((_this.f2).plus(_this.f2)) : (_this.f2));
      let _1404_v30;
      _1404_v30 = _dafny.Set.fromElements((_1370_v1)[_module.__default.safeIndex(_this.f2, new BigNumber((_1370_v1).length))], new BigNumber(-226), p0);
      let _1405_v31;
      _1405_v31 = _dafny.Seq.of(_1403_v29);
      let _1406_v32;
      _1406_v32 = _dafny.MultiSet.fromElements(new BigNumber((_1405_v31).length), _this.f2);
      let _1407_v33;
      _1407_v33 = _dafny.Set.fromElements(false, true);
      let _1408_v34;
      _1408_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1406_v32).cardinality()),_1407_v33);
      _1404_v30 = ((_1404_v30).Difference(function () {
        let _coll47 = new _dafny.Set();
        for (const _compr_47 of (_module.__default.fm31(_1403_v29, new BigNumber((_1408_v34).length), _1403_v29, new BigNumber((function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of (function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(-74))) {
              let _1409_v36 = _compr_49;
              if (((new BigNumber(349)).isLessThanOrEqualTo(_1409_v36)) && ((_1409_v36).isLessThan(new BigNumber(-74)))) {
                _coll49.push([_module.__default.safeDivisionInt(_1409_v36, _this.f2),new BigNumber((_1406_v32).cardinality())]);
              }
            }
            return _coll49;
          }()).Keys.Elements) {
            let _1410_v35 = _compr_48;
            if ((function () {
              let _coll50 = new _dafny.Map();
              for (const _compr_50 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(-74))) {
                let _1409_v36 = _compr_50;
                if (((new BigNumber(349)).isLessThanOrEqualTo(_1409_v36)) && ((_1409_v36).isLessThan(new BigNumber(-74)))) {
                  _coll50.push([_module.__default.safeDivisionInt(_1409_v36, _this.f2),new BigNumber((_1406_v32).cardinality())]);
                }
              }
              return _coll50;
            }()).contains(_1410_v35)) {
              _coll48.push([_module.__default.safeDivisionInt(_1410_v35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((function () {
                let _coll51 = new _dafny.Map();
                for (const _compr_51 of _dafny.IntegerRange(new BigNumber(443), new BigNumber(563))) {
                  let _1411_v37 = _compr_51;
                  if (((new BigNumber(443)).isLessThanOrEqualTo(_1411_v37)) && ((_1411_v37).isLessThan(new BigNumber(563)))) {
                    _coll51.push([(_1411_v37).multipliedBy(_this.f2),p0]);
                  }
                }
                return _coll51;
              }()).length))).length)),_1405_v31]);
            }
          }
          return _coll48;
        }()).length), globalState)).Elements) {
          let _1412_v38 = _compr_47;
          if ((_module.__default.fm31(_1403_v29, new BigNumber((_1408_v34).length), _1403_v29, new BigNumber((function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (function () {
              let _coll53 = new _dafny.Map();
              for (const _compr_53 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(-74))) {
                let _1409_v36 = _compr_53;
                if (((new BigNumber(349)).isLessThanOrEqualTo(_1409_v36)) && ((_1409_v36).isLessThan(new BigNumber(-74)))) {
                  _coll53.push([_module.__default.safeDivisionInt(_1409_v36, _this.f2),new BigNumber((_1406_v32).cardinality())]);
                }
              }
              return _coll53;
            }()).Keys.Elements) {
              let _1410_v35 = _compr_52;
              if ((function () {
                let _coll54 = new _dafny.Map();
                for (const _compr_54 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(-74))) {
                  let _1409_v36 = _compr_54;
                  if (((new BigNumber(349)).isLessThanOrEqualTo(_1409_v36)) && ((_1409_v36).isLessThan(new BigNumber(-74)))) {
                    _coll54.push([_module.__default.safeDivisionInt(_1409_v36, _this.f2),new BigNumber((_1406_v32).cardinality())]);
                  }
                }
                return _coll54;
              }()).contains(_1410_v35)) {
                _coll52.push([_module.__default.safeDivisionInt(_1410_v35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((function () {
                  let _coll55 = new _dafny.Map();
                  for (const _compr_55 of _dafny.IntegerRange(new BigNumber(443), new BigNumber(563))) {
                    let _1411_v37 = _compr_55;
                    if (((new BigNumber(443)).isLessThanOrEqualTo(_1411_v37)) && ((_1411_v37).isLessThan(new BigNumber(563)))) {
                      _coll55.push([(_1411_v37).multipliedBy(_this.f2),p0]);
                    }
                  }
                  return _coll55;
                }()).length))).length)),_1405_v31]);
              }
            }
            return _coll52;
          }()).length), globalState)).contains(_1412_v38)) {
            _coll47.add((_1412_v38).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(656),new BigNumber((_dafny.Set.fromElements(true, true)).length))).length))).length)));
          }
        }
        return _coll47;
      }())).Difference(((_1403_v29) ? (_dafny.Set.fromElements(new BigNumber((_1401_v27).length))) : (_1404_v30)));
      _1408_v34 = (_1408_v34).update(_this.f2, _1407_v33);
      let _1413_v39;
      let _nw257 = new _module.C1();
      _nw257.__ctor(_1401_v27);
      _1413_v39 = _nw257;
      let _1414_v40;
      _1414_v40 = _module.D8.create_DC16(_1413_v39);
      let _1415_v41;
      _1415_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1414_v40,p0);
      r0 = (((_1415_v41).contains(_1414_v40)) ? ((_1415_v41).get(_1414_v40)) : (p0));
      r1 = ((new BigNumber((_dafny.MultiSet.fromElements(_1403_v29, true)).cardinality())).multipliedBy((_dafny.ZERO).minus(p0))).isLessThanOrEqualTo(_this.f2);
      r2 = _1403_v29;
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f2 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1, _module.T2];
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    set f2(value) {
      let _this = this;
      _this._f2 = value;
    };
    __ctor(f2) {
      let _this = this;
      (_this)._f2 = f2;
      return;
    }
    fm0(globalState) {
      let _this = this;
      let _source20 = _module.D0.create_DC1(false, true, new _dafny.CodePoint('l'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("mephw")).length));
      if (_source20.is_DC1) {
        let _1416___mcc_h0 = (_source20).cf1;
        let _1417___mcc_h1 = (_source20).cf2;
        let _1418___mcc_h2 = (_source20).cf3;
        let _1419___mcc_h3 = (_source20).cf4;
        let _1420_cf4 = _1419___mcc_h3;
        let _1421_cf3 = _1418___mcc_h2;
        let _1422_cf2 = _1417___mcc_h1;
        let _1423_cf1 = _1416___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1420_cf4, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), ((_1424_cf3) => function (_1425_i0) {
          return _1424_cf3;
        })(_1421_cf3))).length), _this.f2, _1420_cf4, _1420_cf4),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(837)), function (_1426_i1) {
          return _this.f2;
        })).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1420_cf4, _1420_cf4),_this.f2)));
      } else {
        let _1427___mcc_h4 = (_source20).cf0;
        let _1428_cf0 = _1427___mcc_h4;
        return function () {
          let _coll56 = new _dafny.Map();
          for (const _compr_56 of (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(530)), function (_1429_i2) {
            return _this.f2;
          }), _dafny.Seq.of(_this.f2, _this.f2, new BigNumber((_dafny.MultiSet.fromElements(_this.f2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_dafny.ZERO).minus(_this.f2))).length))).cardinality())))).Elements) {
            let _1430_v0 = _compr_56;
            if ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(530)), function (_1429_i2) {
              return _this.f2;
            }), _dafny.Seq.of(_this.f2, _this.f2, new BigNumber((_dafny.MultiSet.fromElements(_this.f2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_dafny.ZERO).minus(_this.f2))).length))).cardinality())))).contains(_1430_v0)) {
              _coll56.push([_1430_v0,new BigNumber(802)]);
            }
          }
          return _coll56;
        }();
      }
    };
    fm1(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber(929), _this.f2))).length);
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("u")).length))).Union(_dafny.Set.fromElements(_this.f2, _this.f2))).IsDisjointFrom(_dafny.Set.fromElements(_this.f2, _this.f2));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _1431_v0;
      _1431_v0 = _dafny.Map.Empty.slice().updateUnsafe(true,!(p2));
      let _1432_v1;
      _1432_v1 = _1431_v0;
      let _source21 = _1432_v1;
      let _1433___mcc_h0 = _source21;
      let _1434_cf5 = _1433___mcc_h0;
      let _1435_v2;
      _1435_v2 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1436_v3;
      _1436_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1435_v2,(_this.f2).plus(_this.f2));
      _1436_v3 = (_1436_v3).update(_1435_v2, p0);
      let _1437_v4;
      _1437_v4 = _dafny.Set.fromElements(p0, new BigNumber((_module.__default.fm9(p2, p2, p2, _this.f2, globalState)).length));
      _1437_v4 = _1437_v4;
      (_this).f2 = p0;
      let _1438_v5;
      let _nw258 = new _module.C2();
      _nw258.__ctor(_this.f2);
      _1438_v5 = _nw258;
      let _1439_v6;
      _1439_v6 = _dafny.Seq.UnicodeFromString("nuymyuim");
      let _1440_v7;
      _1440_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1441_v8;
      _1441_v8 = _module.D5.create_DC10(p3, new BigNumber((_1439_v6).length), _dafny.Map.Empty.slice().updateUnsafe(_1439_v6,_1440_v7), p3);
      let _1442_v9;
      _1442_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1441_v8, _1441_v8),p3);
      let _1443_v10;
      let _nw259 = new _module.C2();
      _nw259.__ctor((_this).fm1(p3, globalState));
      _1443_v10 = _nw259;
      let _1444_v11;
      _1444_v11 = _dafny.MultiSet.fromElements(_1443_v10);
      let _1445_v12;
      let _nw260 = Array((new BigNumber(4)).toNumber());
      _nw260[(_dafny.ZERO).toNumber()] = p3;
      _nw260[(_dafny.ONE).toNumber()] = !(new BigNumber((_1442_v9).length)).isEqualTo(new BigNumber((_1444_v11).cardinality()));
      _nw260[(new BigNumber(2)).toNumber()] = p2;
      _nw260[(new BigNumber(3)).toNumber()] = false;
      _1445_v12 = _nw260;
      _1445_v12 = _1445_v12;
      let _1446_v13;
      let _nw261 = new _module.C2();
      _nw261.__ctor(p0);
      _1446_v13 = _nw261;
      let _1447_v14;
      let _nw262 = Array((new BigNumber(6)).toNumber());
      _nw262[(_dafny.ZERO).toNumber()] = p0;
      _nw262[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(_1439_v6, _module.__default.safeIndex(p0, new BigNumber((_1439_v6).length)), (_1439_v6)[_module.__default.safeIndex(_this.f2, new BigNumber((_1439_v6).length))])).length);
      _nw262[(new BigNumber(2)).toNumber()] = _this.f2;
      _nw262[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_this.f2);
      _nw262[(new BigNumber(4)).toNumber()] = _this.f2;
      _nw262[(new BigNumber(5)).toNumber()] = _this.f2;
      _1447_v14 = _nw262;
      let _1448_v15;
      _1448_v15 = _dafny.Seq.of(_this.f2, p0);
      let _rhs220 = _1447_v14;
      let _rhs221 = p3;
      let _rhs222 = (p0).plus((_1448_v15)[_module.__default.safeIndex(_this.f2, new BigNumber((_1448_v15).length))]);
      _1447_v14 = _rhs220;
      r0 = _rhs221;
      r2 = _rhs222;
      let _1449_v16;
      _1449_v16 = new _dafny.CodePoint('r'.codePointAt(0));
      if (!(_dafny.Seq.contains(_1439_v6, _1449_v16))) {
        let _1450_v17;
        let _nw263 = new _module.C1();
        _nw263.__ctor(_1439_v6);
        _1450_v17 = _nw263;
        let _pat_let_tv29 = _1450_v17;
        let _source22 = function (_pat_let28_0) {
          return function (_1451_dt__update__tmp_h0) {
            return function (_pat_let29_0) {
              return function (_1452_dt__update_hcf29_h0) {
                return _module.D8.create_DC16(_1452_dt__update_hcf29_h0);
              }(_pat_let29_0);
            }(_pat_let_tv29);
          }(_pat_let28_0);
        }(_module.D8.create_DC16(_1450_v17));
        if (_source22.is_DC17) {
          let _1453___mcc_h1 = (_source22).cf30;
          let _1454___mcc_h2 = (_source22).cf31;
          let _1455___mcc_h3 = (_source22).cf32;
          let _1456_cf32 = _1455___mcc_h3;
          let _1457_cf31 = _1454___mcc_h2;
          let _1458_cf30 = _1453___mcc_h1;
          _1456_cf32 = (p0).multipliedBy((_this.f2).multipliedBy(_1456_cf32));
          let _1459_v18;
          let _nw264 = new _module.C3();
          _nw264.__ctor(p0);
          _1459_v18 = _nw264;
          (_this).f2 = _module.__default.safeDivisionInt(p0, _this.f2);
          let _1460_v19;
          let _nw265 = new _module.C3();
          _nw265.__ctor(_1456_cf32);
          _1460_v19 = _nw265;
        } else {
          let _1461___mcc_h4 = (_source22).cf29;
          let _1462_cf29 = _1461___mcc_h4;
          let _1463_v20;
          _1463_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_1450_v17).fm1(p2, globalState));
          let _rhs223 = (_1463_v20).Merge((_1463_v20).Merge(_1463_v20));
          let _rhs224 = _1445_v12;
          let _rhs225 = (_1440_v7).Merge(_1440_v7);
          let _rhs226 = (_this.f2).minus(_module.__default.safeDivisionInt(_this.f2, _this.f2));
          _1463_v20 = _rhs223;
          _1445_v12 = _rhs224;
          _1440_v7 = _rhs225;
          r2 = _rhs226;
          r2 = _this.f2;
          (_this).f2 = _this.f2;
          r2 = (p0).plus(_this.f2);
        }
        (_this).f2 = (_dafny.ZERO).minus(_this.f2);
        let _1464_v21;
        _1464_v21 = _module.D6.create_DC12(p2, p3);
        let _1465_v22;
        _1465_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1464_v21).dtor_cf23,_this.f2);
        let _rhs227 = _dafny.Seq.IsPrefixOf(_1439_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_1466_v16) => function (_1467_i0) {
          return _1466_v16;
        })(_1449_v16)));
        let _rhs228 = _module.__default.safeModuloInt(_this.f2, (((_1465_v22).contains(p3)) ? ((_1465_v22).get(p3)) : (_this.f2)));
        r0 = _rhs227;
        r2 = _rhs228;
        let _1468_v23;
        let _nw266 = Array((new BigNumber(10)).toNumber());
        _nw266[(_dafny.ZERO).toNumber()] = _1447_v14;
        _nw266[(_dafny.ONE).toNumber()] = ((p3) ? (_1447_v14) : (_1447_v14));
        _nw266[(new BigNumber(2)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(3)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(4)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(5)).toNumber()] = ((p2) ? (_1447_v14) : (_1447_v14));
        _nw266[(new BigNumber(6)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(7)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(8)).toNumber()] = _1447_v14;
        _nw266[(new BigNumber(9)).toNumber()] = ((p2) ? (_1447_v14) : (_1447_v14));
        _1468_v23 = _nw266;
        _1468_v23 = _1468_v23;
        let _index263 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_1447_v14).length));
        (_1447_v14)[_index263] = _module.__default.safeDivisionInt(p0, new BigNumber((_1448_v15).length));
        let _index264 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_1447_v14).length));
        (_1447_v14)[_index264] = _this.f2;
      } else {
        let _1469_v24;
        _1469_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f2)),p0);
        _1469_v24 = (_1469_v24).update(p0, (p0).multipliedBy(new BigNumber(-273)));
        (_this).f2 = (_1446_v13).fm1(p2, globalState);
        let _1470_v25;
        _1470_v25 = _dafny.Seq.of(p2);
        let _1471_v26;
        _1471_v26 = _dafny.Seq.of(_dafny.Seq.of(p3));
        let _1472_v27;
        _1472_v27 = _dafny.Seq.of(_dafny.Seq.Concat(_1470_v25, (_1471_v26)[_module.__default.safeIndex(p0, new BigNumber((_1471_v26).length))]), _1470_v25, _dafny.Seq.Concat(_1470_v25, _1470_v25), _1470_v25, _dafny.Seq.of(p3, p2, p3));
        _1471_v26 = _dafny.Seq.Concat(_1472_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), ((_1473_v25) => function (_1474_i1) {
          return _1473_v25;
        })(_1470_v25)));
        let _1475_v28;
        _1475_v28 = _dafny.Set.fromElements(p2);
        let _rhs229 = _dafny.Seq.Concat(_1448_v15, _dafny.Seq.of(_this.f2, _this.f2));
        let _rhs230 = ((p2) ? (new BigNumber(((_1475_v28).Intersect(_1475_v28)).length)) : (new BigNumber(967)));
        let _rhs231 = p2;
        let _rhs232 = new BigNumber(-227);
        _1448_v15 = _rhs229;
        r2 = _rhs230;
        r0 = _rhs231;
        r2 = _rhs232;
        let _1476_v29;
        _1476_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v14,p0);
        _1476_v29 = (_1476_v29).update(_1447_v14, _this.f2);
      }
      let _1477_v30;
      _1477_v30 = _dafny.Seq.of(p2);
      let _1478_v31;
      _1478_v31 = _dafny.Set.fromElements(p3, p2, p2, !(p2), p2);
      _1439_v6 = _dafny.Seq.Concat(_module.__default.fm28((_1477_v30)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_1477_v30).length))], p3, p2, _1478_v31, globalState), _dafny.Seq.update(_module.__default.fm28(true, !(p3), p2, _module.__default.fm41(_1448_v15, p3, true, p2, globalState), globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm28(true, !(p3), p2, _module.__default.fm41(_1448_v15, p3, true, p2, globalState), globalState)).length)), _1449_v16));
      r0 = p3;
      let _nw267 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
      r1 = _nw267;
      r2 = _this.f2;
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1479_v0;
      _1479_v0 = _dafny.Seq.UnicodeFromString("uxrn");
      let _1480_v1;
      _1480_v1 = _dafny.Seq.of(p2);
      let _1481_v2;
      _1481_v2 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-166)));
      let _1482_v3;
      _1482_v3 = _dafny.Seq.of(_1481_v2);
      let _1483_v4;
      _1483_v4 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1484_v5;
      let _nw268 = Array((new BigNumber(25)).toNumber());
      _nw268[(_dafny.ZERO).toNumber()] = (p2) === (p2);
      _nw268[(_dafny.ONE).toNumber()] = p2;
      _nw268[(new BigNumber(2)).toNumber()] = (p0).isLessThan(p1);
      _nw268[(new BigNumber(3)).toNumber()] = p2;
      _nw268[(new BigNumber(4)).toNumber()] = !(p2);
      _nw268[(new BigNumber(5)).toNumber()] = !(p2);
      _nw268[(new BigNumber(6)).toNumber()] = p2;
      _nw268[(new BigNumber(7)).toNumber()] = ((p2) ? (p2) : (p2));
      _nw268[(new BigNumber(8)).toNumber()] = p2;
      _nw268[(new BigNumber(9)).toNumber()] = p2;
      _nw268[(new BigNumber(10)).toNumber()] = p2;
      _nw268[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_1479_v0, _1479_v0);
      _nw268[(new BigNumber(12)).toNumber()] = p2;
      _nw268[(new BigNumber(13)).toNumber()] = p2;
      _nw268[(new BigNumber(14)).toNumber()] = !_dafny.areEqual(_1480_v1, _1480_v1);
      _nw268[(new BigNumber(15)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_1481_v2), _1482_v3);
      _nw268[(new BigNumber(16)).toNumber()] = p2;
      _nw268[(new BigNumber(17)).toNumber()] = p2;
      _nw268[(new BigNumber(18)).toNumber()] = p2;
      _nw268[(new BigNumber(19)).toNumber()] = p2;
      _nw268[(new BigNumber(20)).toNumber()] = p2;
      _nw268[(new BigNumber(21)).toNumber()] = p2;
      _nw268[(new BigNumber(22)).toNumber()] = _dafny.Seq.contains(_1479_v0, _1483_v4);
      _nw268[(new BigNumber(23)).toNumber()] = p2;
      _nw268[(new BigNumber(24)).toNumber()] = p2;
      _1484_v5 = _nw268;
      let _1485_v6;
      _1485_v6 = _dafny.Set.fromElements(p1, p1, p1);
      let _1486_v7;
      _1486_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1483_v4,_1485_v6);
      let _1487_v8;
      _1487_v8 = _dafny.Seq.of(_1485_v6);
      let _index265 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
      (_1484_v5)[_index265] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of((((_1486_v7).contains(_1483_v4)) ? ((_1486_v7).get(_1483_v4)) : (_dafny.Set.fromElements(new BigNumber((_1479_v0).length), new BigNumber((_1480_v1).length)))), _dafny.Set.fromElements(_this.f2)), _1487_v8);
      let _1488_v9;
      let _nw269 = new _module.C2();
      _nw269.__ctor(p1);
      _1488_v9 = _nw269;
      let _1489_v10;
      _1489_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1488_v9);
      let _1490_v11;
      _1490_v11 = _dafny.Seq.of(p0);
      let _index266 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
      (_1484_v5)[_index266] = (_this).fm8(_1479_v0, !(!((new BigNumber((_1489_v10).length)).isLessThanOrEqualTo(new BigNumber((_1490_v11).length)))), globalState);
      let _1491_v12;
      let _nw270 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _1491_v12 = _nw270;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1491_v12).length))) {
        let _1492_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1492_i0)) && ((_1492_i0).isLessThan(new BigNumber((_1491_v12).length))))) {
          (_1491_v12)[(_1492_i0)] = _module.__default.safeModuloInt(_1492_i0, p1);
        }
      }
      r0 = p1;
      let _hi7 = new BigNumber((_dafny.Seq.UnicodeFromString("qsihpmbty")).length);
      for (let _1493_i1 = p0; _1493_i1.isLessThan(_hi7); _1493_i1 = _1493_i1.plus(_dafny.ONE)) {
        let _1494_v13;
        _1494_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1488_v9,_1493_i1);
        _1494_v13 = (_1494_v13).Merge((_1494_v13).Merge(_1494_v13));
        let _1495_v14;
        let _nw271 = new _module.C1();
        _nw271.__ctor(_1479_v0);
        _1495_v14 = _nw271;
        let _1496_v15;
        _1496_v15 = _module.D8.create_DC17(false, _1495_v14, _1493_i1);
        let _1497_v16;
        _1497_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))],_1483_v4);
        let _pat_let_tv30 = _1495_v14;
        let _pat_let_tv31 = _1497_v16;
        let _pat_let_tv32 = p1;
        let _source23 = function (_pat_let30_0) {
          return function (_1498_dt__update__tmp_h0) {
            return function (_pat_let31_0) {
              return function (_1499_dt__update_hcf31_h0) {
                return function (_pat_let32_0) {
                  return function (_1500_dt__update_hcf30_h0) {
                    return _module.D8.create_DC17(_1500_dt__update_hcf30_h0, _1499_dt__update_hcf31_h0, (_1498_dt__update__tmp_h0).dtor_cf32);
                  }(_pat_let32_0);
                }(!(new BigNumber((_pat_let_tv31).length)).isEqualTo(_pat_let_tv32));
              }(_pat_let31_0);
            }(_pat_let_tv30);
          }(_pat_let30_0);
        }(_1496_v15);
        if (_source23.is_DC17) {
          let _1501___mcc_h0 = (_source23).cf30;
          let _1502___mcc_h1 = (_source23).cf31;
          let _1503___mcc_h2 = (_source23).cf32;
          let _1504_cf32 = _1503___mcc_h2;
          let _1505_cf31 = _1502___mcc_h1;
          let _1506_cf30 = _1501___mcc_h0;
          let _rhs233 = _1504_cf32;
          let _rhs234 = _module.__default.safeDivisionInt(_this.f2, new BigNumber(401));
          r0 = _rhs233;
          _1504_cf32 = _rhs234;
          _1504_cf32 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("mqghm")).length), new BigNumber(104)), _module.__default.safeModuloInt(_1504_cf32, _1493_i1)));
          let _1507_v17;
          _1507_v17 = _module.D11.create_DC20(_1484_v5);
          let _1508_v18;
          _1508_v18 = _dafny.Seq.of(_1484_v5);
          let _1509_v19;
          let _nw272 = Array((new BigNumber(22)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = _1484_v5;
          _nw272[(_dafny.ONE).toNumber()] = (_1507_v17).dtor_cf35;
          _nw272[(new BigNumber(2)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(3)).toNumber()] = ((_1506_cf30) ? (_1484_v5) : (_1484_v5));
          _nw272[(new BigNumber(4)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(5)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(6)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(7)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(8)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(9)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(10)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(11)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(12)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(13)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(14)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(15)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(16)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(17)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(18)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(19)).toNumber()] = _1484_v5;
          _nw272[(new BigNumber(20)).toNumber()] = (_1508_v18)[_module.__default.safeIndex(new BigNumber((_1481_v2).cardinality()), new BigNumber((_1508_v18).length))];
          _nw272[(new BigNumber(21)).toNumber()] = _1484_v5;
          _1509_v19 = _nw272;
          let _index267 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1509_v19).length));
          (_1509_v19)[_index267] = _1484_v5;
          let _1510_v20;
          _1510_v20 = _dafny.Set.fromElements(_1506_cf30, (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]);
          let _1511_v21;
          _1511_v21 = _dafny.MultiSet.fromElements(_1510_v20);
          let _index268 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1509_v19).length));
          let _nw273 = Array((new BigNumber(11)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = ((p2) ? (_1506_cf30) : (p2));
          _nw273[(_dafny.ONE).toNumber()] = !((_1511_v21).IsProperSubsetOf(_1511_v21));
          _nw273[(new BigNumber(2)).toNumber()] = (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))];
          _nw273[(new BigNumber(3)).toNumber()] = _1506_cf30;
          _nw273[(new BigNumber(4)).toNumber()] = (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))];
          _nw273[(new BigNumber(5)).toNumber()] = !((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]);
          _nw273[(new BigNumber(6)).toNumber()] = true;
          _nw273[(new BigNumber(7)).toNumber()] = _1506_cf30;
          _nw273[(new BigNumber(8)).toNumber()] = (_1481_v2).IsSubsetOf(_1481_v2);
          _nw273[(new BigNumber(9)).toNumber()] = false;
          _nw273[(new BigNumber(10)).toNumber()] = p2;
          (_1509_v19)[_index268] = _nw273;
          _1506_cf30 = !((((_1481_v2).contains(p1)) ? ((_1481_v2).get(p1)) : (_1493_i1))).isEqualTo(p1);
        } else {
          let _1512___mcc_h3 = (_source23).cf29;
          let _1513_cf29 = _1512___mcc_h3;
          let _1514_v22;
          _1514_v22 = _module.D6.create_DC13((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], new BigNumber(763));
          let _1515_v23;
          _1515_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1514_v22,_1484_v5);
          _1484_v5 = (((_1515_v23).contains(_1514_v22)) ? ((_1515_v23).get(_1514_v22)) : (_1484_v5));
          r0 = p0;
          _1485_v6 = function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of _dafny.IntegerRange(new BigNumber(96), new BigNumber(-363))) {
              let _1516_v24 = _compr_57;
              if (((new BigNumber(96)).isLessThanOrEqualTo(_1516_v24)) && ((_1516_v24).isLessThan(new BigNumber(-363)))) {
                _coll57.add((_1516_v24).plus(_1493_i1));
              }
            }
            return _coll57;
          }();
          (_this).f2 = p1;
        }
        _1479_v0 = _1479_v0;
        if ((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]) {
          let _index269 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
          (_1484_v5)[_index269] = p2;
          let _1517_v25;
          let _nw274 = new _module.C1();
          _nw274.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_1518_v4) => function (_1519_i2) {
            return _1518_v4;
          })(_1483_v4)));
          _1517_v25 = _nw274;
          let _1520_v26;
          _1520_v26 = _dafny.Seq.of(_1484_v5);
          _1520_v26 = _1520_v26;
          let _1521_v27;
          _1521_v27 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(542));
          let _1522_v28;
          _1522_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1521_v27,(_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]);
          let _1523_v29;
          _1523_v29 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1491_v12);
          let _1524_v30;
          let _nw275 = Array((new BigNumber(24)).toNumber());
          _nw275[(_dafny.ZERO).toNumber()] = _1490_v11;
          _nw275[(_dafny.ONE).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(2)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(3)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(4)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(p0, new BigNumber((_1480_v1).length), p0), _module.__default.safeIndex(new BigNumber((_1522_v28).length), new BigNumber((_dafny.Seq.of(p0, new BigNumber((_1480_v1).length), p0)).length)), new BigNumber(-996));
          _nw275[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(new BigNumber(715));
          _nw275[(new BigNumber(7)).toNumber()] = (_1495_v14).fm12((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], new BigNumber((_1523_v29).length), (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], globalState);
          _nw275[(new BigNumber(8)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(9)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(10)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_1493_i1);
          _nw275[(new BigNumber(12)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(p1, new BigNumber(((_1495_v14).f3).length));
          _nw275[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_1525_i3) {
            return new BigNumber(451);
          });
          _nw275[(new BigNumber(15)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(16)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(17)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), ((_1526_p0) => function (_1527_i4) {
            return _1526_p0;
          })(p0));
          _nw275[(new BigNumber(19)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(20)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(21)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(22)).toNumber()] = _1490_v11;
          _nw275[(new BigNumber(23)).toNumber()] = _1490_v11;
          _1524_v30 = _nw275;
          let _1528_v31;
          let _1529_v32;
          let _1530_v33;
          let _out23;
          let _out24;
          let _out25;
          let _outcollector6 = (_1517_v25).m1(new BigNumber((_dafny.Seq.UnicodeFromString("ck")).length), _1524_v30, !(!((_this.f2).isLessThan(_1493_i1))), false, globalState);
          _out23 = _outcollector6[0];
          _out24 = _outcollector6[1];
          _out25 = _outcollector6[2];
          _1528_v31 = _out23;
          _1529_v32 = _out24;
          _1530_v33 = _out25;
          r0 = _1493_i1;
        } else {
          let _1531_v34;
          _1531_v34 = _dafny.Set.fromElements((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], p2);
          let _1532_v35;
          _1532_v35 = _module.D4.create_DC5(((true) ? (_1531_v34) : (_1531_v34)));
          let _1533_v38;
          _1533_v38 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_1485_v6).Elements) {
              let _1534_v37 = _compr_58;
              if ((_1485_v6).contains(_1534_v37)) {
                _coll58.push([(_1534_v37).plus(new BigNumber(445)),(_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]]);
              }
            }
            return _coll58;
          }(),_this.f2);
          let _1535_v39;
          _1535_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,(_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))]);
          let _index270 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
          let _rhs235 = (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))];
          let _rhs236 = new BigNumber((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (((_1533_v38).update(_1535_v39, p1)).Merge(_1533_v38)).Keys.Elements) {
              let _1536_v36 = _compr_59;
              if ((((_1533_v38).update(_1535_v39, p1)).Merge(_1533_v38)).contains(_1536_v36)) {
                _coll59.push([_1536_v36,_module.__default.safeDivisionInt(_this.f2, p0)]);
              }
            }
            return _coll59;
          }()).length);
          let _rhs237 = _1532_v35;
          let _rhs238 = _1483_v4;
          let _lhs132 = _1484_v5;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
          _lhs132[_lhs133] = _rhs235;
          r0 = _rhs236;
          _1532_v35 = _rhs237;
          _1483_v4 = _rhs238;
          _1490_v11 = _dafny.Seq.Concat(_dafny.Seq.update(_1490_v11, _module.__default.safeIndex(p0, new BigNumber((_1490_v11).length)), p1), _1490_v11);
          (_this).f2 = ((_1488_v9).fm1((_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], globalState)).multipliedBy(p1);
          r0 = _this.f2;
          let _index271 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
          (_1484_v5)[_index271] = p2;
        }
      }
      let _1537_v40;
      _1537_v40 = _dafny.MultiSet.fromElements(p2, p2);
      _1537_v40 = (_1537_v40).Union(_1537_v40);
      let _index272 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length));
      (_1484_v5)[_index272] = (_this).fm8(_dafny.Seq.Concat(_1479_v0, _1479_v0), (_1484_v5)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1484_v5).length))], globalState);
      let _1538_v41;
      _1538_v41 = _dafny.Set.fromElements(_1484_v5, _1484_v5);
      r0 = _module.__default.safeDivisionInt(_this.f2, _module.__default.safeModuloInt(new BigNumber((_1538_v41).length), new BigNumber((_dafny.Seq.of(_this.f2)).length)));
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _1539_v0;
      _1539_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p3));
      _1539_v0 = (_1539_v0).update(p3, p3);
      let _1540_i0;
      _1540_i0 = _dafny.ZERO;
      L6: {
        while (p3) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1540_i0)) {
              break L6;
            }
            _1540_i0 = (_1540_i0).plus(_dafny.ONE);
            let _1541_v1;
            _1541_v1 = new _dafny.CodePoint('q'.codePointAt(0));
            let _1542_v2;
            _1542_v2 = _dafny.Seq.of((new BigNumber(391)).minus(new BigNumber(438)), (_this.f2).multipliedBy(_this.f2), new BigNumber(-19), (_this.f2).plus(_this.f2));
            let _1543_v3;
            _1543_v3 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1541_v1);
            let _1544_v4;
            _1544_v4 = _dafny.Seq.UnicodeFromString("kwftu");
            let _rhs239 = (((_1543_v3).contains((_this).fm8(_1544_v4, true, globalState))) ? ((_1543_v3).get((_this).fm8(_1544_v4, true, globalState))) : (new _dafny.CodePoint('k'.codePointAt(0))));
            let _rhs240 = _dafny.Seq.Concat(_1542_v2, _1542_v2);
            let _rhs241 = (_this).fm1(_dafny.Seq.IsPrefixOf(_1544_v4, _1544_v4), globalState);
            _1541_v1 = _rhs239;
            _1542_v2 = _rhs240;
            r0 = _rhs241;
            let _1545_v5;
            _1545_v5 = true;
            _1545_v5 = p3;
            let _1546_v6;
            let _out26;
            _out26 = _module.__default.m0(p2, _1545_v5, globalState);
            _1546_v6 = _out26;
            _1542_v2 = _1542_v2;
          }
        }
      }
      if (p3) {
        let _1547_v7;
        _1547_v7 = true;
        let _1548_v8;
        _1548_v8 = _dafny.Seq.UnicodeFromString("yauevfig");
        _1547_v7 = (_this).fm8(_1548_v8, ((_1547_v7) ? (p3) : (p2)), globalState);
        r1 = _module.__default.safeModuloInt(_this.f2, new BigNumber((_1548_v8).length));
        let _1549_v9;
        _1549_v9 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1550_v10;
        let _init30 = function (_1551_i1) {
          return _module.D7.create_DC14(_dafny.Seq.of(_this.f2, _this.f2, new BigNumber((_dafny.Set.fromElements(_this.f2, _this.f2)).length), (_dafny.ZERO).minus(_this.f2), _this.f2));
        };
        let _nw276 = Array((new BigNumber(10)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw276.length); _i0_30++) {
          _nw276[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1550_v10 = _nw276;
        let _1552_v11;
        _1552_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1549_v9,_1550_v10);
        let _rhs242 = (_1552_v11).Merge(_1552_v11);
        let _rhs243 = _this.f2;
        let _rhs244 = p2;
        let _rhs245 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f2));
        _1552_v11 = _rhs242;
        r0 = _rhs243;
        _1547_v7 = _rhs244;
        r1 = _rhs245;
        let _1553_v12;
        let _nw277 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1553_v12 = _nw277;
        let _index273 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1553_v12).length));
        (_1553_v12)[_index273] = p2;
        let _1554_v13;
        let _nw278 = new _module.C2();
        _nw278.__ctor(new BigNumber(693));
        _1554_v13 = _nw278;
        let _1555_v14;
        _1555_v14 = _dafny.MultiSet.fromElements(_1554_v13);
        let _index274 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1553_v12).length));
        (_1553_v12)[_index274] = (_1555_v14).equals(_1555_v14);
        let _1556_v15;
        let _nw279 = new _module.C1();
        _nw279.__ctor(((p2) ? (_1548_v8) : (_1548_v8)));
        _1556_v15 = _nw279;
      } else {
        let _1557_v16;
        let _nw280 = Array((new BigNumber(8)).toNumber()).fill([]);
        _1557_v16 = _nw280;
        _1557_v16 = ((!(true) || (p2)) ? (_1557_v16) : (_1557_v16));
        let _1558_v17;
        _1558_v17 = _dafny.MultiSet.fromElements(p1);
        let _1559_v18;
        _1559_v18 = _module.D5.create_DC8(_1558_v17);
        let _1560_v19;
        _1560_v19 = true;
        let _1561_v20;
        _1561_v20 = _dafny.Seq.of(_1560_v19);
        let _1562_v21;
        let _nw281 = new _module.C3();
        _nw281.__ctor(_this.f2);
        _1562_v21 = _nw281;
        let _1563_v22;
        _1563_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1561_v20,_1562_v21);
        let _1564_v23;
        _1564_v23 = _dafny.Map.Empty.slice().updateUnsafe(_this.f2,new BigNumber(479));
        let _1565_v24;
        _1565_v24 = _dafny.MultiSet.fromElements(_this.f2, _module.__default.safeModuloInt(_this.f2, (((_1564_v23).contains(new BigNumber(61))) ? ((_1564_v23).get(new BigNumber(61))) : (_this.f2))));
        let _1566_v25;
        _1566_v25 = _module.D0.create_DC0(p2);
        let _1567_v26;
        _1567_v26 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("hweyh"));
        let _1568_v27;
        _1568_v27 = _dafny.Seq.UnicodeFromString("ljkq");
        let _1569_v28;
        _1569_v28 = _dafny.Seq.of(_this.f2);
        let _rhs246 = _1559_v18;
        let _rhs247 = (_1566_v25).dtor_cf0;
        let _rhs248 = _1563_v22;
        let _rhs249 = new BigNumber((_dafny.Seq.Concat(_1567_v26, _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-522)), function (_1570_i2) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }), _1568_v27))).length);
        let _rhs250 = (_dafny.MultiSet.FromArray(_1569_v28)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), function (_1571_i3) {
          return _this.f2;
        })));
        _1559_v18 = _rhs246;
        _1560_v19 = _rhs247;
        _1563_v22 = _rhs248;
        r1 = _rhs249;
        _1565_v24 = _rhs250;
        let _1572_v29;
        let _nw282 = new _module.C3();
        _nw282.__ctor(_module.__default.safeDivisionInt(new BigNumber(-329), (_dafny.ZERO).minus(_this.f2)));
        _1572_v29 = _nw282;
        let _1573_v30;
        let _nw283 = Array((new BigNumber(11)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = _1568_v27;
        _nw283[(_dafny.ONE).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(2)).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(3)).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-663)), function (_1574_i4) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _1568_v27);
        _nw283[(new BigNumber(5)).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gjfrs"), _1568_v27);
        _nw283[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1568_v27, _1568_v27);
        _nw283[(new BigNumber(8)).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(9)).toNumber()] = _1568_v27;
        _nw283[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("boyqhdgay");
        _1573_v30 = _nw283;
        let _index275 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1573_v30).length));
        (_1573_v30)[_index275] = _1568_v27;
        let _index276 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1573_v30).length));
        (_1573_v30)[_index276] = _1568_v27;
        _1561_v20 = _dafny.Seq.Concat(_1561_v20, _dafny.Seq.Concat(_1561_v20, _1561_v20));
      }
      let _hi8 = _this.f2;
      for (let _1575_i5 = ((p3) ? (_this.f2) : (_this.f2)); _1575_i5.isLessThan(_hi8); _1575_i5 = _1575_i5.plus(_dafny.ONE)) {
        let _1576_v31;
        _1576_v31 = _dafny.Seq.UnicodeFromString("trytchnh");
        let _1577_v32;
        let _nw284 = new _module.C1();
        _nw284.__ctor(_1576_v31);
        _1577_v32 = _nw284;
        let _1578_v33;
        _1578_v33 = _module.D8.create_DC16(_1577_v32);
        _1578_v33 = _1578_v33;
        r1 = _1575_i5;
        let _1579_v34;
        let _init31 = ((_1580_p2) => function (_1581_i6) {
          return _1580_p2;
        })(p2);
        let _nw285 = Array((new BigNumber(21)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw285.length); _i0_31++) {
          _nw285[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1579_v34 = _nw285;
        let _1582_v35;
        _1582_v35 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1579_v34);
        _1582_v35 = _1582_v35;
        let _index277 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1579_v34).length));
        (_1579_v34)[_index277] = (_this).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(507)), function (_1583_i7) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }), !(p2), globalState);
        let _index278 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1579_v34).length));
        (_1579_v34)[_index278] = p3;
      }
      let _1584_v36;
      let _nw286 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
      _1584_v36 = _nw286;
      _1584_v36 = _1584_v36;
      let _1585_v37;
      let _nw287 = new _module.C0();
      _nw287.__ctor(p3, !(p2));
      _1585_v37 = _nw287;
      r0 = _module.__default.safeModuloInt(new BigNumber(837), _this.f2);
      r1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f2, (_dafny.ZERO).minus(_this.f2)));
      return [r0, r1];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this.f1 = [];
      this._f0 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f0, f1) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_1586_i0) {
        return new BigNumber((_dafny.Seq.of((_this).f0)).length);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_1587_i1) {
        return (_this).f0;
      })),(_dafny.ZERO).minus((_this).f0));
    };
    fm1(p0, globalState) {
      let _this = this;
      return new BigNumber(426);
    };
    fm6(p0, globalState) {
      let _this = this;
      return true;
    };
    fm7(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('m'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()),new _dafny.CodePoint('m'.codePointAt(0)))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('b'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('e'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('c'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('j'.codePointAt(0))))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f0,new _dafny.CodePoint('i'.codePointAt(0)))));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      r0 = p2;
      let _1588_v0;
      _1588_v0 = _dafny.Seq.UnicodeFromString("gtifm");
      let _hi9 = (_this).f0;
      for (let _1589_i0 = new BigNumber((_1588_v0).length); _1589_i0.isLessThan(_hi9); _1589_i0 = _1589_i0.plus(_dafny.ONE)) {
        if (p2) {
          let _1590_v1;
          _1590_v1 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1589_i0,p3)).length));
          let _1591_v2;
          _1591_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
          let _1592_v3;
          _1592_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber((_1591_v2).length))).IsSubsetOf(_1590_v1),(_this).f0);
          let _rhs251 = new BigNumber(-735);
          let _rhs252 = p0;
          let _rhs253 = (_dafny.ZERO).minus((_this).f0);
          let _rhs254 = _1592_v3;
          let _rhs255 = !((_this).f0).isEqualTo(p0);
          r2 = _rhs251;
          r2 = _rhs252;
          r2 = _rhs253;
          _1592_v3 = _rhs254;
          r0 = _rhs255;
          let _arr2 = _this.f1;
          let _index279 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_this.f1).length));
          _arr2[_index279] = p2;
          let _arr3 = _this.f1;
          let _index280 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_this.f1).length));
          _arr3[_index280] = (_1589_i0).isLessThan(new BigNumber((_1591_v2).length));
          let _1593_v4;
          _1593_v4 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1591_v2);
          let _1594_v5;
          _1594_v5 = _1591_v2;
          _1593_v4 = (_1593_v4).update(p2, (_1594_v5));
          let _arr4 = _this.f1;
          let _index281 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_this.f1).length));
          _arr4[_index281] = !(true);
          let _1595_v6;
          let _nw288 = new _module.C2();
          _nw288.__ctor(_1589_i0);
          _1595_v6 = _nw288;
        } else {
          let _1596_v7;
          _1596_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f1);
          _1596_v7 = (_1596_v7).update(p2, _this.f1);
          let _arr5 = _this.f1;
          let _index282 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length));
          _arr5[_index282] = !(p3) || (p2);
          let _1597_v8;
          _1597_v8 = _dafny.MultiSet.fromElements(_this.f1, _this.f1);
          let _arr6 = _this.f1;
          let _index283 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length));
          let _rhs256 = ((_1597_v8).Union(_1597_v8)).IsDisjointFrom((_1597_v8).Intersect(_1597_v8));
          let _rhs257 = _1588_v0;
          let _lhs134 = _this.f1;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length));
          _lhs134[_lhs135] = _rhs256;
          _1588_v0 = _rhs257;
          r0 = (_this.f1)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length))];
          let _1598_v9;
          let _nw289 = new _module.C0();
          _nw289.__ctor((_this.f1)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length))], p3);
          _1598_v9 = _nw289;
          let _1599_v10;
          let _init32 = ((_1600_i0) => function (_1601_i1) {
            return (_1601_i1).multipliedBy(_1600_i0);
          })(_1589_i0);
          let _nw290 = Array((new BigNumber(24)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw290.length); _i0_32++) {
            _nw290[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1599_v10 = _nw290;
          let _index284 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1599_v10).length));
          (_1599_v10)[_index284] = _module.__default.safeDivisionInt(new BigNumber(-535), _1589_i0);
          let _1602_v11;
          _1602_v11 = _dafny.Set.fromElements(_1598_v9.f5);
          let _1603_v12;
          _1603_v12 = _dafny.MultiSet.fromElements(_1602_v11);
          let _index285 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_1599_v10).length));
          (_1599_v10)[_index285] = new BigNumber(((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_1598_v9.f5, (_this.f1)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_this.f1).length))]))).Difference(_1603_v12)).cardinality());
        }
        let _1604_v13;
        _1604_v13 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1605_v14;
        let _nw291 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _1605_v14 = _nw291;
        let _1606_v15;
        _1606_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
        let _1607_v16;
        _1607_v16 = _module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-155)), ((_1608_v0) => function (_1609_i2) {
  return (_1608_v0)[_module.__default.safeIndex((_this).f0, new BigNumber((_1608_v0).length))];
})(_1588_v0)));
        let _pat_let_tv33 = _1588_v0;
        let _rhs258 = new _dafny.CodePoint('u'.codePointAt(0));
        let _rhs259 = _1605_v14;
        let _rhs260 = !(p3) || (!(p2));
        let _rhs261 = _1606_v15;
        let _rhs262 = ((p2) ? (_module.D6.create_DC11(_1588_v0)) : (function (_pat_let33_0) {
          return function (_1610_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_1611_dt__update_hcf21_h0) {
                return _module.D6.create_DC11(_1611_dt__update_hcf21_h0);
              }(_pat_let34_0);
            }(_pat_let_tv33);
          }(_pat_let33_0);
        }(_module.__default.fm42(_1589_i0, p3, p2, globalState))));
        _1604_v13 = _rhs258;
        _1605_v14 = _rhs259;
        r0 = _rhs260;
        _1606_v15 = _rhs261;
        _1607_v16 = _rhs262;
        let _1612_v17;
        let _nw292 = new _module.C1();
        _nw292.__ctor(_1588_v0);
        _1612_v17 = _nw292;
        let _1613_v18;
        _1613_v18 = _module.D8.create_DC16(_1612_v17);
        let _rhs263 = _1613_v18;
        let _rhs264 = (p0).plus((_dafny.ZERO).minus(_1589_i0));
        let _rhs265 = ((p2) ? (false) : (p2));
        _1613_v18 = _rhs263;
        r2 = _rhs264;
        r0 = _rhs265;
        let _1614_v19;
        let _nw293 = new _module.C3();
        _nw293.__ctor(_1589_i0);
        _1614_v19 = _nw293;
      }
      let _1615_v20;
      _1615_v20 = _dafny.Seq.of(p3);
      r2 = _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Seq.Concat(_1615_v20, _1615_v20)).length));
      let _1616_i3;
      _1616_i3 = _dafny.ZERO;
      L7: {
        while (p3) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1616_i3)) {
              break L7;
            }
            _1616_i3 = (_1616_i3).plus(_dafny.ONE);
            _1588_v0 = _1588_v0;
            r0 = !(!(p2) || (p3));
            let _1617_v21;
            _1617_v21 = _dafny.MultiSet.fromElements(new BigNumber(660), p0);
            if ((_module.__default.safeDivisionInt(p0, (_this).f0)).isLessThanOrEqualTo(new BigNumber((_1617_v21).cardinality()))) {
              let _1618_v22;
              _1618_v22 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1588_v0);
              let _1619_v23;
              _1619_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((((_1618_v22).contains(false)) ? ((_1618_v22).get(false)) : (_1588_v0)), _module.__default.safeIndex((_this).f0, new BigNumber(((((_1618_v22).contains(false)) ? ((_1618_v22).get(false)) : (_1588_v0))).length)), new _dafny.CodePoint('p'.codePointAt(0))),_1588_v0);
              let _1620_v24;
              _1620_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("darjgmy")).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-130)), function (_1621_i4) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              }));
              _1619_v23 = (_1619_v23).update((((_1620_v24).contains(p0)) ? ((_1620_v24).get(p0)) : (_1588_v0)), _dafny.Seq.UnicodeFromString("min"));
              _1588_v0 = _1588_v0;
              r2 = _module.__default.safeModuloInt(new BigNumber(205), (_this).f0);
              let _1622_v25;
              _1622_v25 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f0)).update((_this).f0, (_this).f0),_dafny.Seq.UnicodeFromString("ja"));
              let _1623_v26;
              _1623_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f0);
              _1622_v25 = (_1622_v25).update(_1623_v26, _dafny.Seq.UnicodeFromString("sxeuxp"));
              let _rhs266 = _module.__default.safeModuloInt(p0, p0);
              let _rhs267 = p3;
              let _rhs268 = p3;
              let _rhs269 = (_this).fm6(true, globalState);
              r2 = _rhs266;
              r0 = _rhs267;
              r0 = _rhs268;
              r0 = _rhs269;
            } else {
              r2 = (_this).f0;
              let _1624_v27;
              let _nw294 = new _module.C3();
              _nw294.__ctor((_this).f0);
              _1624_v27 = _nw294;
              let _1625_v28;
              let _nw295 = Array((new BigNumber(8)).toNumber());
              _nw295[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).f0, p0);
              _nw295[(_dafny.ONE).toNumber()] = ((_this).f0).minus((_this).f0);
              _nw295[(new BigNumber(2)).toNumber()] = p0;
              _nw295[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).f0);
              _nw295[(new BigNumber(4)).toNumber()] = (_this).f0;
              _nw295[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
              _nw295[(new BigNumber(6)).toNumber()] = (_this).f0;
              _nw295[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_this).f0, (((_1617_v21).contains(p0)) ? ((_1617_v21).get(p0)) : (p0)));
              _1625_v28 = _nw295;
              let _index286 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1625_v28).length));
              (_1625_v28)[_index286] = new BigNumber(842);
              let _index287 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1625_v28).length));
              (_1625_v28)[_index287] = p0;
              let _1626_v29;
              let _init33 = function (_1627_i5) {
                return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(33),new BigNumber(753));
              };
              let _nw296 = Array((new BigNumber(2)).toNumber());
              for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw296.length); _i0_33++) {
                _nw296[_i0_33] = _init33(new BigNumber(_i0_33));
              }
              _1626_v29 = _nw296;
              let _1628_v30;
              _1628_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1615_v20)[_module.__default.safeIndex((_this).f0, new BigNumber((_1615_v20).length))],p3);
              let _1629_v31;
              _1629_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1626_v29),_1628_v30);
              _1629_v31 = (_1629_v31).update(_1626_v29, _1628_v30);
              let _1630_v32;
              _1630_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _1631_v33;
              _1631_v33 = _dafny.Seq.of((_this).f0);
              r0 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_1617_v21).contains(new BigNumber((_1630_v32).length))) ? ((_1617_v21).get(new BigNumber((_1630_v32).length))) : ((_dafny.ZERO).minus(p0))),!(p3))).length)).isLessThanOrEqualTo(new BigNumber((_1631_v33).length));
            }
            let _arr7 = _this.f1;
            let _index288 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_this.f1).length));
            _arr7[_index288] = (p0).isEqualTo(p0);
            let _arr8 = _this.f1;
            let _index289 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_this.f1).length));
            _arr8[_index289] = p2;
          }
        }
      }
      let _1632_v34;
      let _nw297 = new _module.C2();
      _nw297.__ctor(_module.__default.safeModuloInt((_this).f0, new BigNumber((_1588_v0).length)));
      _1632_v34 = _nw297;
      let _1633_v35;
      _1633_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1588_v0).length));
      let _1634_v36;
      _1634_v36 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f0),p2);
      let _1635_v37;
      _1635_v37 = _dafny.Seq.of((_1632_v34).fm1(true, globalState), new BigNumber((_1634_v36).length), p0);
      let _1636_v38;
      _1636_v38 = _dafny.MultiSet.fromElements(p2);
      let _1637_v39;
      _1637_v39 = _dafny.Set.fromElements(p3, p2);
      let _1638_v40;
      let _nw298 = Array((new BigNumber(22)).toNumber());
      _nw298[(_dafny.ZERO).toNumber()] = (_this).fm1(p3, globalState);
      _nw298[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-749)), function (_1639_i6) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("dhwnmp"))).length);
      _nw298[(new BigNumber(2)).toNumber()] = p0;
      _nw298[(new BigNumber(3)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(4)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(5)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(6)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(7)).toNumber()] = (new BigNumber((_1633_v35).length)).plus(new BigNumber((_1635_v37).length));
      _nw298[(new BigNumber(8)).toNumber()] = (((_1636_v38).contains(p3)) ? ((_1636_v38).get(p3)) : (new BigNumber((_1637_v39).length)));
      _nw298[(new BigNumber(9)).toNumber()] = ((p2) ? ((_1632_v34).fm1(p3, globalState)) : (p0));
      _nw298[(new BigNumber(10)).toNumber()] = (p0).minus((_dafny.ZERO).minus(new BigNumber(-8)));
      _nw298[(new BigNumber(11)).toNumber()] = p0;
      _nw298[(new BigNumber(12)).toNumber()] = ((_dafny.ZERO).minus((_this).f0)).multipliedBy((_dafny.ZERO).minus(p0));
      _nw298[(new BigNumber(13)).toNumber()] = p0;
      _nw298[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1588_v0).length), (_this).f0);
      _nw298[(new BigNumber(15)).toNumber()] = new BigNumber((_1588_v0).length);
      _nw298[(new BigNumber(16)).toNumber()] = new BigNumber((_1588_v0).length);
      _nw298[(new BigNumber(17)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(18)).toNumber()] = (p0).minus((_this).f0);
      _nw298[(new BigNumber(19)).toNumber()] = (_this).f0;
      _nw298[(new BigNumber(20)).toNumber()] = new BigNumber((_1588_v0).length);
      _nw298[(new BigNumber(21)).toNumber()] = new BigNumber((_1635_v37).length);
      _1638_v40 = _nw298;
      _1638_v40 = _1638_v40;
      let _1640_v41;
      _1640_v41 = _dafny.Set.fromElements(p0);
      r0 = !(_1640_v41).equals((_1640_v41).Intersect(_1640_v41));
      let _1641_v42;
      _1641_v42 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
      let _1642_v43;
      _1642_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1588_v0).length),_1641_v42);
      let _nw299 = Array((new BigNumber(24)).toNumber());
      _nw299[(_dafny.ZERO).toNumber()] = _1641_v42;
      _nw299[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,p3);
      _nw299[(new BigNumber(2)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(3)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(4)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(5)).toNumber()] = ((p2) ? (_1641_v42) : (_dafny.Map.Empty.slice().updateUnsafe(p3,true)));
      _nw299[(new BigNumber(6)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(7)).toNumber()] = ((p2) ? (_dafny.Map.Empty.slice().updateUnsafe(p3,p2)) : ((_1641_v42).update(p2, (_this).fm6(false, globalState))));
      _nw299[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(p3),p2);
      _nw299[(new BigNumber(9)).toNumber()] = (_1641_v42).Merge(_1641_v42);
      _nw299[(new BigNumber(10)).toNumber()] = (_1641_v42).update(p3, !(p2));
      _nw299[(new BigNumber(11)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(12)).toNumber()] = _module.__default.fm21(globalState);
      _nw299[(new BigNumber(13)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(14)).toNumber()] = ((((_1642_v43).contains(p0)) ? ((_1642_v43).get(p0)) : (_1641_v42))).Merge(_module.__default.fm21(globalState));
      _nw299[(new BigNumber(15)).toNumber()] = (_1641_v42).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,p3));
      _nw299[(new BigNumber(16)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(17)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(18)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(19)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(20)).toNumber()] = (_1641_v42).Merge(_1641_v42);
      _nw299[(new BigNumber(21)).toNumber()] = _1641_v42;
      _nw299[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
      _nw299[(new BigNumber(23)).toNumber()] = _1641_v42;
      r1 = _nw299;
      let _1643_v44;
      _1643_v44 = _dafny.Map.Empty.slice().updateUnsafe((_module.D6.create_DC12(p2, p3)).dtor_cf22,(_dafny.ZERO).minus((_this).f0));
      r2 = (((_1643_v44).contains(!(p3))) ? ((_1643_v44).get(!(p3))) : (((_this).f0).plus((_this).f0)));
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1644_v0;
      _1644_v0 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1645_v1;
      let _nw300 = Array((new BigNumber(2)).toNumber());
      _nw300[(_dafny.ZERO).toNumber()] = _1644_v0;
      _nw300[(_dafny.ONE).toNumber()] = ((true) ? (_1644_v0) : (_1644_v0));
      _1645_v1 = _nw300;
      let _index290 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
      (_1645_v1)[_index290] = _1644_v0;
      let _1646_v2;
      let _init34 = function (_1647_i0) {
        return _dafny.Seq.UnicodeFromString("tfbmfbdi");
      };
      let _nw301 = Array((new BigNumber(2)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw301.length); _i0_34++) {
        _nw301[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1646_v2 = _nw301;
      let _1648_v3;
      _1648_v3 = _dafny.Seq.UnicodeFromString("bu");
      let _index291 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
      (_1646_v2)[_index291] = _1648_v3;
      let _1649_v4;
      _1649_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f0);
      let _index292 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
      let _index293 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
      let _rhs270 = new _dafny.CodePoint('a'.codePointAt(0));
      let _rhs271 = _dafny.Seq.Concat(_1648_v3, _dafny.Seq.Concat(_dafny.Seq.update(_1648_v3, _module.__default.safeIndex(new BigNumber((_1648_v3).length), new BigNumber((_1648_v3).length)), _1644_v0), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_1650_v0) => function (_1651_i1) {
        return _1650_v0;
      })(_1644_v0)), _module.__default.safeIndex((((_1649_v4).contains(p2)) ? ((_1649_v4).get(p2)) : (p1)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_1652_v0) => function (_1653_i1) {
        return _1652_v0;
      })(_1644_v0))).length)), _1644_v0)));
      let _lhs136 = _1645_v1;
      let _lhs137 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
      let _lhs138 = _1646_v2;
      let _lhs139 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
      _lhs136[_lhs137] = _rhs270;
      _lhs138[_lhs139] = _rhs271;
      let _1654_v5;
      _1654_v5 = _module.D6.create_DC11((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]);
      let _1655_v6;
      _1655_v6 = _dafny.Seq.of((_1654_v5).dtor_cf21, (_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]);
      let _rhs272 = _1648_v3;
      let _rhs273 = new BigNumber((_dafny.Set.fromElements((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))], _dafny.Seq.update(_1648_v3, _module.__default.safeIndex(p0, new BigNumber((_1648_v3).length)), (_1645_v1)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length))]), (_1655_v6)[_module.__default.safeIndex(p1, new BigNumber((_1655_v6).length))], (_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))])).length);
      _1648_v3 = _rhs272;
      r0 = _rhs273;
      let _1656_v7;
      _1656_v7 = _dafny.Seq.of(_this.f1);
      let _1657_v8;
      _1657_v8 = _module.D11.create_DC20((_1656_v7)[_module.__default.safeIndex(p1, new BigNumber((_1656_v7).length))]);
      let _1658_v9;
      _1658_v9 = _dafny.Seq.of(_1657_v8, _1657_v8);
      let _1659_v10;
      _1659_v10 = _dafny.Seq.of(_1658_v9);
      if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat((_1659_v10)[_module.__default.safeIndex((_this).f0, new BigNumber((_1659_v10).length))], _1658_v9), _module.__default.safeIndex((_this).f0, new BigNumber((_dafny.Seq.Concat((_1659_v10)[_module.__default.safeIndex((_this).f0, new BigNumber((_1659_v10).length))], _1658_v9)).length)), _1657_v8), _1657_v8)) {
        r0 = (p0).multipliedBy(new BigNumber(-22));
        let _1660_v11;
        let _init35 = ((_1661_p0) => function (_1662_i2) {
          return (_1662_i2).multipliedBy(_1661_p0);
        })(p0);
        let _nw302 = Array((new BigNumber(2)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw302.length); _i0_35++) {
          _nw302[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1660_v11 = _nw302;
        let _index294 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
        (_1660_v11)[_index294] = (new BigNumber(859)).plus(p0);
        let _1663_v12;
        _1663_v12 = _dafny.Set.fromElements((_dafny.ZERO).minus(p1), (_this).f0);
        let _index295 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
        (_1660_v11)[_index295] = _module.__default.safeDivisionInt((_this).f0, new BigNumber((((p2) ? (_1663_v12) : (_1663_v12))).length));
        if (p2) {
          let _1664_v13;
          _1664_v13 = _dafny.Seq.of(p0);
          let _1665_v14;
          _1665_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1644_v0,_dafny.MultiSet.fromElements((_1660_v11)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length))], (_1664_v13)[_module.__default.safeIndex((((_1649_v4).contains(p2)) ? ((_1649_v4).get(p2)) : ((_this).f0)), new BigNumber((_1664_v13).length))]));
          let _1666_v15;
          _1666_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1644_v0,p0);
          let _index296 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
          (_1660_v11)[_index296] = (new BigNumber(((_1665_v14).Merge(_module.__default.fm43(globalState))).length)).minus((((_1666_v15).contains(_1644_v0)) ? ((_1666_v15).get(_1644_v0)) : (new BigNumber((_1664_v13).length))));
          let _index297 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
          (_1645_v1)[_index297] = _module.__default.fm37(p1, globalState);
          let _index298 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
          (_1660_v11)[_index298] = (_this).f0;
          let _1667_v16;
          let _out27;
          _out27 = _module.__default.m0(p2, p2, globalState);
          _1667_v16 = _out27;
          r0 = (_1660_v11)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length))];
        } else {
          let _index299 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
          (_1660_v11)[_index299] = (_this).f0;
          let _1668_v17;
          _1668_v17 = false;
          _1668_v17 = _1668_v17;
          let _index300 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1660_v11).length));
          (_1660_v11)[_index300] = (_this).f0;
          let _1669_v18;
          let _nw303 = new _module.C3();
          _nw303.__ctor(p0);
          _1669_v18 = _nw303;
          let _1670_v19;
          _1670_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1671_v20;
          _1671_v20 = _1670_v19;
          let _1672_v21;
          _1672_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v20,_this.f1);
          _1672_v21 = (_1672_v21).update(_1671_v20, _this.f1);
        }
        let _1673_v22;
        _1673_v22 = _dafny.Set.fromElements(!(p2), p2, p2);
        let _1674_v23;
        _1674_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1673_v22,_dafny.Seq.UnicodeFromString("qgpvdxtv"));
        _1674_v23 = (_1674_v23).update(_1673_v22, _1648_v3);
        let _1675_v24;
        let _nw304 = Array((new BigNumber(16)).toNumber()).fill([]);
        _1675_v24 = _nw304;
        let _pat_let_tv34 = _1648_v3;
        let _1676_v25;
        _1676_v25 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let35_0) {
          return function (_1677_dt__update__tmp_h0) {
            return function (_pat_let36_0) {
              return function (_1678_dt__update_hcf21_h0) {
                return _module.D6.create_DC11(_1678_dt__update_hcf21_h0);
              }(_pat_let36_0);
            }(_pat_let_tv34);
          }(_pat_let35_0);
        }(_1654_v5)).dtor_cf21,_1675_v24);
        _1676_v25 = (_1676_v25).update(_dafny.Seq.update(_1648_v3, _module.__default.safeIndex(new BigNumber(-940), new BigNumber((_1648_v3).length)), _1644_v0), _1675_v24);
      } else {
        let _1679_v26;
        let _nw305 = new _module.C0();
        _nw305.__ctor((_this).fm6(p2, globalState), p2);
        _1679_v26 = _nw305;
        r0 = new BigNumber(183);
        let _1680_v27;
        _1680_v27 = _dafny.Seq.of(new BigNumber((_1648_v3).length), (_this).f0, (_this).f0);
        let _1681_v28;
        let _nw306 = new _module.C3();
        _nw306.__ctor((_1680_v27)[_module.__default.safeIndex((_this).f0, new BigNumber((_1680_v27).length))]);
        _1681_v28 = _nw306;
        let _arr9 = _this.f1;
        let _index301 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_this.f1).length));
        _arr9[_index301] = _1679_v26.f5;
        let _arr10 = _this.f1;
        let _index302 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_this.f1).length));
        _arr10[_index302] = p2;
        if (_1679_v26.f4) {
          let _1682_v29;
          let _nw307 = new _module.C3();
          _nw307.__ctor(new BigNumber((_dafny.Seq.Concat(_1648_v3, _1648_v3)).length));
          _1682_v29 = _nw307;
          _1682_v29 = ((_1679_v26.f5) ? (_1682_v29) : (_1682_v29));
          r0 = (_this).f0;
          let _1683_v30;
          _1683_v30 = _dafny.MultiSet.fromElements(false);
          r0 = ((p1).multipliedBy(new BigNumber((_1683_v30).cardinality()))).minus((_this).f0);
          let _1684_v31;
          let _nw308 = Array((new BigNumber(14)).toNumber());
          _1684_v31 = _nw308;
          let _1685_v32;
          let _nw309 = new _module.C4();
          _nw309.__ctor(new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length));
          _1685_v32 = _nw309;
          let _index303 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1684_v31).length));
          (_1684_v31)[_index303] = _1685_v32;
          let _1686_v33;
          let _nw310 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _1686_v33 = _nw310;
          let _1687_v34;
          _1687_v34 = _dafny.Seq.of(_1686_v33, _1686_v33);
          let _1688_v35;
          let _nw311 = Array((new BigNumber(16)).toNumber());
          _nw311[(_dafny.ZERO).toNumber()] = _1686_v33;
          _nw311[(_dafny.ONE).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(2)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(3)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(4)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(5)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(6)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(7)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(8)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(9)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(10)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(11)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(12)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(13)).toNumber()] = (_1687_v34)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1687_v34).length))];
          _nw311[(new BigNumber(14)).toNumber()] = _1686_v33;
          _nw311[(new BigNumber(15)).toNumber()] = _1686_v33;
          _1688_v35 = _nw311;
          let _1689_v36;
          _1689_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v35,_1680_v27);
          let _1690_v37;
          _1690_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f0);
          let _1691_v38;
          _1691_v38 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1654_v5,new BigNumber((_1690_v37).length)));
          let _1692_v39;
          _1692_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC11(_1648_v3),p1);
          let _1693_v40;
          _1693_v40 = _module.D11.create_DC21(_1644_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), function (_1694_i3) {
  return (_this).f0;
})).length), p0);
          let _index304 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1684_v31).length));
          let _rhs274 = _dafny.Seq.Concat(_module.__default.fm36(_1679_v26.f4, new _dafny.CodePoint('i'.codePointAt(0)), (_this).fm1(_1679_v26.f5, globalState), globalState), (((_1689_v36).contains(_1688_v35)) ? ((_1689_v36).get(_1688_v35)) : (_1680_v27)));
          let _rhs275 = _module.__default.safeModuloInt(p1, (_this).f0);
          let _rhs276 = (((_1679_v26.f5) ? (_1691_v38) : ((_dafny.Set.fromElements(_1692_v39, _dafny.Map.Empty.slice().updateUnsafe(_1654_v5,_1685_v32.f2)))))).IsSubsetOf(_1691_v38);
          let _rhs277 = _1685_v32;
          let _rhs278 = (p1).minus((_1693_v40).dtor_cf38);
          let _lhs140 = _1679_v26;
          let _lhs141 = _1684_v31;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1684_v31).length));
          let _lhs143 = _1685_v32;
          _1680_v27 = _rhs274;
          r0 = _rhs275;
          _lhs140.f5 = _rhs276;
          _lhs141[_lhs142] = _rhs277;
          _lhs143.f2 = _rhs278;
          let _1695_v41;
          _1695_v41 = _dafny.Set.fromElements(_1679_v26.f5);
          (_1685_v32).f2 = new BigNumber((_1695_v41).length);
        } else {
          let _1696_v42;
          let _nw312 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _1696_v42 = _nw312;
          let _1697_v43;
          _1697_v43 = _dafny.Seq.of(_1696_v42, _1696_v42);
          r0 = _module.__default.safeDivisionInt(p1, ((_1679_v26.f5) ? (p0) : (new BigNumber((_1697_v43).length))));
          let _1698_v44;
          _1698_v44 = _dafny.Set.fromElements(_1679_v26.f4);
          let _1699_v45;
          _1699_v45 = _module.D4.create_DC5(_1698_v44);
          let _index305 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
          let _index306 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
          let _rhs279 = (_this).fm1(_1679_v26.f5, globalState);
          let _rhs280 = (_1648_v3)[_module.__default.safeIndex((p0).multipliedBy(new BigNumber(144)), new BigNumber((_1648_v3).length))];
          let _rhs281 = new BigNumber(((_dafny.Set.fromElements(_1679_v26.f4)).Union((_1699_v45).dtor_cf8)).length);
          let _rhs282 = (_1655_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f0)).length), new BigNumber((_1655_v6).length))];
          let _lhs144 = _1645_v1;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
          let _lhs146 = _1646_v2;
          let _lhs147 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
          r0 = _rhs279;
          _lhs144[_lhs145] = _rhs280;
          r0 = _rhs281;
          _lhs146[_lhs147] = _rhs282;
          r0 = (p0).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1680_v27,_1679_v26.f5)).Merge(function () {
            let _coll60 = new _dafny.Map();
            for (const _compr_60 of (_module.__default.fm44(p0, (_this).f0, (_this).f0, _1679_v26.f5, globalState)).Keys.Elements) {
              let _1700_v46 = _compr_60;
              if ((_module.__default.fm44(p0, (_this).f0, (_this).f0, _1679_v26.f5, globalState)).contains(_1700_v46)) {
                _coll60.push([_1700_v46,false]);
              }
            }
            return _coll60;
          }())).length));
          let _1701_v47;
          _1701_v47 = _dafny.Set.fromElements(p1);
          let _1702_v48;
          _1702_v48 = _dafny.Map.Empty.slice().updateUnsafe(((_1679_v26.f4) ? (_1679_v26.f4) : (_1679_v26.f5)),(_dafny.Set.fromElements(p0, new BigNumber(418))).Union(_1701_v47));
          _1702_v48 = (_1702_v48).update(_1679_v26.f4, _1701_v47);
          let _1703_v49;
          let _nw313 = new _module.C0();
          _nw313.__ctor(!((_this.f1)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_this.f1).length))]), _1679_v26.f5);
          _1703_v49 = _nw313;
        }
      }
      let _hi10 = new BigNumber((_dafny.Seq.UnicodeFromString("ltbe")).length);
      for (let _1704_i4 = (_dafny.ZERO).minus((p1).minus(new BigNumber((_1648_v3).length))); _1704_i4.isLessThan(_hi10); _1704_i4 = _1704_i4.plus(_dafny.ONE)) {
        let _1705_v50;
        let _nw314 = new _module.C4();
        _nw314.__ctor(new BigNumber(-578));
        _1705_v50 = _nw314;
        let _1706_v51;
        _1706_v51 = _dafny.Set.fromElements(p0, p0);
        let _1707_v53;
        _1707_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm6(p2, globalState),(_1706_v51).IsDisjointFrom(function () {
          let _coll61 = new _dafny.Set();
          for (const _compr_61 of _dafny.IntegerRange(new BigNumber(775), new BigNumber(402))) {
            let _1708_v52 = _compr_61;
            if (((new BigNumber(775)).isLessThanOrEqualTo(_1708_v52)) && ((_1708_v52).isLessThan(new BigNumber(402)))) {
              _coll61.add(_module.__default.safeDivisionInt(_1708_v52, p0));
            }
          }
          return _coll61;
        }()));
        _1707_v53 = (_1707_v53).update((p2) === (p2), p2);
        let _1709_v54;
        let _nw315 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1709_v54 = _nw315;
        _1709_v54 = _1709_v54;
        if (p2) {
          let _1710_v55;
          _1710_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1644_v0,_1654_v5);
          let _1711_v56;
          _1711_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1710_v55,_1645_v1);
          r0 = new BigNumber((_1711_v56).length);
          let _1712_v57;
          _1712_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1707_v53).length),(_1704_i4).isLessThan((_dafny.ZERO).minus(p1)));
          _1712_v57 = ((_1712_v57).Merge(_1712_v57)).Merge(_1712_v57);
          let _1713_v58;
          _1713_v58 = _dafny.Set.fromElements(p2, p2, !(p2));
          let _arr11 = _this.f1;
          let _index307 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length));
          _arr11[_index307] = (new BigNumber((_1713_v58).length)).isLessThanOrEqualTo(_1704_i4);
          let _arr12 = _this.f1;
          let _index308 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length));
          _arr12[_index308] = p2;
          let _arr13 = _this.f1;
          let _index309 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length));
          _arr13[_index309] = !(p2);
          let _arr14 = _this.f1;
          let _index310 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length));
          let _rhs283 = (_this.f1)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length))];
          let _rhs284 = (((_this.f1)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length))]) ? ((_1705_v50).fm1((_this.f1)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length))], globalState)) : (_module.__default.safeDivisionInt(p0, _1704_i4)));
          let _lhs148 = _this.f1;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_this.f1).length));
          _lhs148[_lhs149] = _rhs283;
          r0 = _rhs284;
        } else {
          let _1714_v59;
          _1714_v59 = false;
          let _1715_v60;
          _1715_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _1716_v61;
          _1716_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_1717_v0) => function (_1718_i5) {
            return _1717_v0;
          })(_1644_v0)),_1715_v60);
          let _1719_v62;
          _1719_v62 = _module.D5.create_DC10(false, p1, _1716_v61, p2);
          _1714_v59 = (_1719_v62).dtor_cf17;
          let _1720_v63;
          let _nw316 = new _module.C3();
          _nw316.__ctor(p1);
          _1720_v63 = _nw316;
          let _1721_v64;
          _1721_v64 = _dafny.Seq.of(_1704_i4);
          r0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1720_v63).fm10(_1704_i4, globalState),_1721_v64)).length);
          r0 = new BigNumber(-570);
          let _1722_v65;
          let _nw317 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1722_v65 = _nw317;
          let _1723_v66;
          let _nw318 = Array((new BigNumber(8)).toNumber());
          _nw318[(_dafny.ZERO).toNumber()] = _module.D5.create_DC10(p2, p1, _1716_v61, _1714_v59);
          _nw318[(_dafny.ONE).toNumber()] = _1719_v62;
          _nw318[(new BigNumber(2)).toNumber()] = _module.D5.create_DC10(p2, _1704_i4, _1716_v61, _1714_v59);
          _nw318[(new BigNumber(3)).toNumber()] = _1719_v62;
          _nw318[(new BigNumber(4)).toNumber()] = _module.D5.create_DC10(p2, new BigNumber((_dafny.Seq.of(p1, _1704_i4)).length), _1716_v61, p2);
          _nw318[(new BigNumber(5)).toNumber()] = _1719_v62;
          _nw318[(new BigNumber(6)).toNumber()] = _1719_v62;
          _nw318[(new BigNumber(7)).toNumber()] = _1719_v62;
          _1723_v66 = _nw318;
          let _index311 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1722_v65).length));
          (_1722_v65)[_index311] = _1723_v66;
          let _index312 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1722_v65).length));
          (_1722_v65)[_index312] = _1723_v66;
        }
      }
      let _hi11 = p1;
      for (let _1724_i6 = (new BigNumber(58)).plus((_this).f0); _1724_i6.isLessThan(_hi11); _1724_i6 = _1724_i6.plus(_dafny.ONE)) {
        let _pat_let_tv35 = _1644_v0;
        let _1725_v67;
        _1725_v67 = _module.D0.create_DC0((function (_pat_let37_0) {
  return function (_1726_dt__update__tmp_h2) {
    return function (_pat_let38_0) {
      return function (_1727_dt__update_hcf3_h0) {
        return _module.D0.create_DC1((_1726_dt__update__tmp_h2).dtor_cf1, (_1726_dt__update__tmp_h2).dtor_cf2, _1727_dt__update_hcf3_h0, (_1726_dt__update__tmp_h2).dtor_cf4);
      }(_pat_let38_0);
    }(_pat_let_tv35);
  }(_pat_let37_0);
}(_module.D0.create_DC1(!(p2), p2, (_1645_v1)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length))], p0))).dtor_cf2);
        _1725_v67 = _1725_v67;
        let _1728_v68;
        _1728_v68 = false;
        _1728_v68 = p2;
        let _1729_v69;
        let _nw319 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
        _1729_v69 = _nw319;
        let _1730_v70;
        _1730_v70 = _dafny.Seq.of(_1728_v68, _1728_v68);
        let _1731_v71;
        _1731_v71 = _dafny.Seq.of(_1730_v70);
        let _index313 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1729_v69).length));
        (_1729_v69)[_index313] = _1731_v71;
        let _index314 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1729_v69).length));
        (_1729_v69)[_index314] = _1731_v71;
        r0 = p0;
      }
      let _1732_i7;
      _1732_i7 = _dafny.ZERO;
      L8: {
        while (!(p2)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1732_i7)) {
              break L8;
            }
            _1732_i7 = (_1732_i7).plus(_dafny.ONE);
            r0 = p1;
            if (((_this).f0).isEqualTo(new BigNumber((_1648_v3).length))) {
              let _1733_v72;
              let _init36 = ((_1734_p0) => function (_1735_i8) {
                return _module.__default.safeModuloInt(_1735_i8, (_dafny.ZERO).minus(_1734_p0));
              })(p0);
              let _nw320 = Array((new BigNumber(16)).toNumber());
              for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw320.length); _i0_36++) {
                _nw320[_i0_36] = _init36(new BigNumber(_i0_36));
              }
              _1733_v72 = _nw320;
              let _1736_v73;
              _1736_v73 = _module.D4.create_DC6((_this).f0, _1733_v72, p0);
              _1733_v72 = (_1736_v73).dtor_cf10;
              let _1737_v74;
              _1737_v74 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p0).isEqualTo(new BigNumber(331)));
              _1737_v74 = (_1737_v74).update(p2, !(_1737_v74).equals(_1737_v74));
              let _1738_v75;
              _1738_v75 = _dafny.MultiSet.fromElements(p0, new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length));
              let _1739_v76;
              _1739_v76 = (_1738_v75).Difference(_dafny.MultiSet.fromElements(new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length)));
              _1739_v76 = _1739_v76;
              let _1740_v77;
              let _nw321 = new _module.C2();
              _nw321.__ctor(p0);
              _1740_v77 = _nw321;
              let _1741_v78;
              _1741_v78 = _dafny.Seq.of(new BigNumber(205), p1);
              _1741_v78 = _1741_v78;
            } else {
              let _1742_v79;
              _1742_v79 = false;
              let _1743_v80;
              _1743_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _1744_v81;
              _1744_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v80,p2);
              let _1745_v82;
              _1745_v82 = _dafny.Seq.of(p2);
              let _1746_v83;
              _1746_v83 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),p0);
              let _rhs285 = p2;
              let _rhs286 = (_1649_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_1745_v82)[_module.__default.safeIndex(p1, new BigNumber((_1745_v82).length))]),new BigNumber((_1746_v83).length)));
              let _rhs287 = (_dafny.Map.Empty.slice().updateUnsafe(_1743_v80,_1742_v79)).Merge(_1744_v81);
              let _rhs288 = _1742_v79;
              let _rhs289 = ((new BigNumber((_1745_v82).length)).plus(p0)).minus(p0);
              _1742_v79 = _rhs285;
              _1649_v4 = _rhs286;
              _1744_v81 = _rhs287;
              _1742_v79 = _rhs288;
              r0 = _rhs289;
              let _1747_v84;
              _1747_v84 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Concat(_1648_v3, _dafny.Seq.update((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))], _module.__default.safeIndex(new BigNumber((_1743_v80).length), new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length)), _1644_v0))).length));
              _1747_v84 = (_1747_v84).Intersect(_1747_v84);
              _1742_v79 = ((_dafny.ZERO).minus(p0)).isLessThan(new BigNumber(-473));
              _1742_v79 = (new BigNumber(532)).isLessThan(new BigNumber(343));
              let _index315 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
              (_1646_v2)[_index315] = _1648_v3;
            }
            let _source24 = _1657_v8;
            if (_source24.is_DC21) {
              let _1748___mcc_h0 = (_source24).cf36;
              let _1749___mcc_h1 = (_source24).cf37;
              let _1750___mcc_h2 = (_source24).cf38;
              let _1751_cf38 = _1750___mcc_h2;
              let _1752_cf37 = _1749___mcc_h1;
              let _1753_cf36 = _1748___mcc_h0;
              let _index316 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length));
              (_1646_v2)[_index316] = _1648_v3;
              let _index317 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
              (_1645_v1)[_index317] = ((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))])[_module.__default.safeIndex(_1752_cf37, new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length))];
              r0 = (_dafny.ZERO).minus(_1751_cf38);
              let _1754_v85;
              _1754_v85 = true;
              let _1755_v86;
              _1755_v86 = _dafny.MultiSet.fromElements(new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length), new BigNumber(717), p1);
              let _1756_v87;
              _1756_v87 = _dafny.MultiSet.fromElements(_1755_v86);
              _1754_v85 = (_1756_v87).equals(_1756_v87);
            } else if (_source24.is_DC22) {
              let _1757___mcc_h3 = (_source24).cf39;
              let _1758_cf39 = _1757___mcc_h3;
              let _index318 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1645_v1).length));
              (_1645_v1)[_index318] = (_1648_v3)[_module.__default.safeIndex(new BigNumber((_1648_v3).length), new BigNumber((_1648_v3).length))];
              r0 = p0;
              let _1759_v88;
              _1759_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f0,p2);
              _1758_cf39 = ((_1759_v88).Merge(_1759_v88)).contains(((!(p2)) ? ((_this).fm1(!(p2), globalState)) : ((_this).f0)));
              _1759_v88 = _1759_v88;
            } else {
              let _1760___mcc_h4 = (_source24).cf35;
              let _1761_cf35 = _1760___mcc_h4;
              (_this).f1 = _1761_cf35;
              r0 = (_this).f0;
              let _1762_v89;
              _1762_v89 = _dafny.Seq.of(new BigNumber(((_1646_v2)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1646_v2).length))]).length));
              let _1763_v90;
              _1763_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1762_v89).length),(_this).fm1((_this).fm6(p2, globalState), globalState));
              let _1764_v91;
              _1764_v91 = _dafny.Seq.of(p1, new BigNumber((_1763_v90).length), p1, (_this).f0);
              r0 = _module.__default.safeDivisionInt(p1, (_1764_v91)[_module.__default.safeIndex(p1, new BigNumber((_1764_v91).length))]);
              let _1765_v92;
              _1765_v92 = _dafny.Set.fromElements(p0);
              let _rhs290 = ((_this).f0).plus(new BigNumber((_1765_v92).length));
              let _rhs291 = p1;
              r0 = _rhs290;
              r0 = _rhs291;
            }
            let _1766_v93;
            _1766_v93 = false;
            _1766_v93 = !(!(true));
          }
        }
      }
      let _1767_v94;
      _1767_v94 = _dafny.MultiSet.fromElements((_this).f0, p1, p0);
      let _1768_v95;
      _1768_v95 = _dafny.Set.fromElements(_1644_v0);
      r0 = (((_1767_v94).contains((p1).multipliedBy(new BigNumber((_1768_v95).length)))) ? ((_1767_v94).get((p1).multipliedBy(new BigNumber((_1768_v95).length)))) : (new BigNumber((_1648_v3).length)));
      return r0;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), function (_1769_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(469),true)).length);
      }),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-916),false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(818))).length))).length),new BigNumber(937))).length))).length)),new BigNumber((_dafny.Seq.UnicodeFromString("mxumvgapi")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(77)), function (_1770_i1) {
        return new BigNumber(502);
      }),new BigNumber((_dafny.Seq.UnicodeFromString("kqtm")).length)));
    };
    fm1(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("up"), _dafny.Seq.UnicodeFromString("fgl")), _dafny.Seq.UnicodeFromString("lmdmkvq"))).length);
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _1771_v0;
      _1771_v0 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber(689)).isEqualTo(p0),p0);
      let _1772_v1;
      _1772_v1 = _dafny.Seq.of(p3);
      let _1773_v2;
      _1773_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1774_v3;
      _1774_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1772_v1).length),new BigNumber((_1773_v2).length));
      _1771_v0 = (_1771_v0).update(!(_1774_v3).contains(p0), p0);
      let _1775_v4;
      let _nw322 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _1775_v4 = _nw322;
      let _nw323 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1775_v4 = _nw323;
      if (p3) {
        r0 = (p0).isLessThanOrEqualTo(p0);
        let _1776_v5;
        let _nw324 = Array((new BigNumber(4)).toNumber()).fill([]);
        _1776_v5 = _nw324;
        let _1777_v6;
        let _nw325 = Array((new BigNumber(2)).toNumber()).fill(false);
        _1777_v6 = _nw325;
        let _1778_v7;
        _1778_v7 = _dafny.MultiSet.fromElements(p2);
        let _index319 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
        (_1777_v6)[_index319] = (_1778_v7).IsSubsetOf(_1778_v7);
        let _1779_v8;
        _1779_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1776_v5);
        let _index320 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
        let _rhs292 = (((_1779_v8).contains(p2)) ? ((_1779_v8).get(p2)) : (_1776_v5));
        let _rhs293 = _module.__default.fm5(globalState);
        let _lhs150 = _1777_v6;
        let _lhs151 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
        _1776_v5 = _rhs292;
        _lhs150[_lhs151] = _rhs293;
        let _1780_v9;
        _1780_v9 = _dafny.Seq.of(p0, p0, p0, p0);
        let _1781_v10;
        _1781_v10 = _dafny.MultiSet.fromElements(_1780_v9);
        let _1782_v11;
        _1782_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1781_v10,false);
        _1782_v11 = (_1782_v11).update(_1781_v10, _module.__default.fm5(globalState));
        if (p3) {
          r0 = !(false);
          let _1783_v12;
          let _nw326 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1783_v12 = _nw326;
          let _1784_v13;
          _1784_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("vxtfidplh"));
          let _1785_v14;
          _1785_v14 = _dafny.Seq.UnicodeFromString("to");
          let _index321 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_1783_v12).length));
          (_1783_v12)[_index321] = _dafny.Seq.Concat((((_1784_v13).contains(p0)) ? ((_1784_v13).get(p0)) : (_dafny.Seq.UnicodeFromString("o"))), _1785_v14);
          let _1786_v15;
          let _nw327 = Array((new BigNumber(7)).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = _1775_v4;
          _nw327[(_dafny.ONE).toNumber()] = _1775_v4;
          _nw327[(new BigNumber(2)).toNumber()] = _1775_v4;
          _nw327[(new BigNumber(3)).toNumber()] = _1775_v4;
          _nw327[(new BigNumber(4)).toNumber()] = _1775_v4;
          _nw327[(new BigNumber(5)).toNumber()] = ((p2) ? (_1775_v4) : (_1775_v4));
          _nw327[(new BigNumber(6)).toNumber()] = _1775_v4;
          _1786_v15 = _nw327;
          let _1787_v16;
          _1787_v16 = _dafny.Seq.of(_1786_v15, _1786_v15);
          let _index322 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          let _index323 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_1783_v12).length));
          let _rhs294 = p2;
          let _rhs295 = _1785_v14;
          let _rhs296 = (_1787_v16)[_module.__default.safeIndex(p0, new BigNumber((_1787_v16).length))];
          let _lhs152 = _1777_v6;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          let _lhs154 = _1783_v12;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_1783_v12).length));
          _lhs152[_lhs153] = _rhs294;
          _lhs154[_lhs155] = _rhs295;
          _1786_v15 = _rhs296;
          r2 = (_dafny.ZERO).minus((_1780_v9)[_module.__default.safeIndex((p0).minus((_dafny.ZERO).minus(new BigNumber(((_1783_v12)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_1783_v12).length))]).length))), new BigNumber((_1780_v9).length))]);
          let _1788_v17;
          _1788_v17 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1789_v18;
          _1789_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1778_v7,!((_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))]));
          let _index324 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          let _rhs297 = false;
          let _rhs298 = (_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))];
          let _rhs299 = _1788_v17;
          let _rhs300 = _1789_v18;
          let _lhs156 = _1777_v6;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          r0 = _rhs297;
          _lhs156[_lhs157] = _rhs298;
          _1788_v17 = _rhs299;
          _1789_v18 = _rhs300;
          let _1790_v19;
          let _nw328 = new _module.C3();
          _nw328.__ctor((p0).multipliedBy(p0));
          _1790_v19 = _nw328;
        } else {
          let _1791_v20;
          _1791_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1774_v3,_1775_v4);
          let _index325 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          (_1777_v6)[_index325] = (p0).isLessThanOrEqualTo(new BigNumber((_1791_v20).length));
          let _rhs301 = _1772_v1;
          let _rhs302 = (p0).plus((_dafny.ZERO).minus(p0));
          _1772_v1 = _rhs301;
          r2 = _rhs302;
          let _1792_v21;
          let _init37 = ((_1793_p2) => function (_1794_i0) {
            return ((_1793_p2) ? (new _dafny.CodePoint('c'.codePointAt(0))) : (new _dafny.CodePoint('u'.codePointAt(0))));
          })(p2);
          let _nw329 = Array((new BigNumber(17)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw329.length); _i0_37++) {
            _nw329[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1792_v21 = _nw329;
          let _1795_v22;
          _1795_v22 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index326 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1792_v21).length));
          (_1792_v21)[_index326] = _1795_v22;
          let _index327 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1792_v21).length));
          (_1792_v21)[_index327] = _1795_v22;
          let _1796_v23;
          _1796_v23 = _dafny.MultiSet.fromElements(p0, p0, p0, p0);
          let _1797_v24;
          _1797_v24 = _1796_v23;
          let _1798_v25;
          let _nw330 = Array((new BigNumber(13)).toNumber());
          _nw330[(_dafny.ZERO).toNumber()] = _module.__default.fm19(p0, p0, (_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))], p0, globalState);
          _nw330[(_dafny.ONE).toNumber()] = _1796_v23;
          _nw330[(new BigNumber(2)).toNumber()] = _1796_v23;
          _nw330[(new BigNumber(3)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(4)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(5)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(6)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(7)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(8)).toNumber()] = _1796_v23;
          _nw330[(new BigNumber(9)).toNumber()] = _1796_v23;
          _nw330[(new BigNumber(10)).toNumber()] = _1797_v24;
          _nw330[(new BigNumber(11)).toNumber()] = _1796_v23;
          _nw330[(new BigNumber(12)).toNumber()] = ((p3) ? (_1797_v24) : (_1797_v24));
          _1798_v25 = _nw330;
          let _index328 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_1798_v25).length));
          (_1798_v25)[_index328] = _1797_v24;
          let _index329 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_1798_v25).length));
          (_1798_v25)[_index329] = _1796_v23;
          r0 = p2;
        }
        if ((_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))]) {
          let _1799_v26;
          let _nw331 = new _module.C4();
          _nw331.__ctor((_dafny.ZERO).minus(p0));
          _1799_v26 = _nw331;
          let _1800_v27;
          let _nw332 = new _module.C5();
          _nw332.__ctor(p0, _1777_v6);
          _1800_v27 = _nw332;
          let _1801_v28;
          _1801_v28 = _module.D14.create_DC25(_1776_v5);
          _1776_v5 = (_1801_v28).dtor_cf42;
          let _1802_v29;
          let _nw333 = new _module.C4();
          _nw333.__ctor((_1800_v27).f0);
          _1802_v29 = _nw333;
          let _1803_v30;
          _1803_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1773_v2).Merge(_1773_v2),p2);
          let _index330 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          let _rhs303 = _1775_v4;
          let _rhs304 = (_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))];
          let _rhs305 = (_dafny.ZERO).minus((_1799_v26).fm1(p3, globalState));
          let _rhs306 = _1803_v30;
          let _rhs307 = (_1777_v6)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length))];
          let _lhs158 = _1777_v6;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          _1775_v4 = _rhs303;
          r0 = _rhs304;
          r2 = _rhs305;
          _1803_v30 = _rhs306;
          _lhs158[_lhs159] = _rhs307;
        } else {
          let _index331 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          (_1777_v6)[_index331] = !(!(!(p2)));
          r2 = p0;
          let _index332 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1777_v6).length));
          (_1777_v6)[_index332] = p2;
          let _1804_v31;
          let _1805_v32;
          let _1806_v33;
          let _1807_v34;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector7 = (_this).m4(true, (_module.__default.fm42(p0, p2, p3, globalState)).dtor_cf21, p0, p3, globalState);
          _out28 = _outcollector7[0];
          _out29 = _outcollector7[1];
          _out30 = _outcollector7[2];
          _out31 = _outcollector7[3];
          _1804_v31 = _out28;
          _1805_v32 = _out29;
          _1806_v33 = _out30;
          _1807_v34 = _out31;
          let _1808_v35;
          _1808_v35 = _dafny.Seq.UnicodeFromString("bevnw");
          _1808_v35 = _1808_v35;
        }
      } else {
        r0 = (p0).isLessThanOrEqualTo(p0);
        let _1809_v36;
        let _nw334 = Array((new BigNumber(2)).toNumber());
        _nw334[(_dafny.ZERO).toNumber()] = true;
        _nw334[(_dafny.ONE).toNumber()] = p2;
        _1809_v36 = _nw334;
        let _index333 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length));
        (_1809_v36)[_index333] = p3;
        let _index334 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length));
        (_1809_v36)[_index334] = p3;
        r2 = p0;
        let _1810_v37;
        _1810_v37 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1811_v38;
        _1811_v38 = _module.D0.create_DC1(!((_1809_v36)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length))]), (_1809_v36)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length))], _1810_v37, (p0).plus(p0));
        _1811_v38 = _1811_v38;
        let _1812_v39;
        _1812_v39 = _dafny.Map.Empty.slice().updateUnsafe(true,_1774_v3);
        let _1813_v40;
        _1813_v40 = _dafny.Map.Empty.slice().updateUnsafe((_1809_v36)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length))],(_1774_v3).Merge((((_1812_v39).contains((_1809_v36)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length))])) ? ((_1812_v39).get((_1809_v36)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_1809_v36).length))])) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm38(p3, globalState)).cardinality()),p0)))));
        _1813_v40 = (_1813_v40).update(_module.__default.fm5(globalState), _1774_v3);
      }
      let _hi12 = new BigNumber(575);
      for (let _1814_i1 = (_dafny.ZERO).minus(p0); _1814_i1.isLessThan(_hi12); _1814_i1 = _1814_i1.plus(_dafny.ONE)) {
        let _1815_v41;
        _1815_v41 = new _dafny.CodePoint('r'.codePointAt(0));
        let _rhs308 = !(!((_module.__default.fm45(p0, true, p2, _1814_i1, globalState)).dtor_cf39));
        let _rhs309 = !(!((true) === (p3)));
        let _rhs310 = _1815_v41;
        let _rhs311 = _1814_i1;
        r0 = _rhs308;
        r0 = _rhs309;
        _1815_v41 = _rhs310;
        r2 = _rhs311;
        r0 = !(p2);
        r2 = p0;
        let _1816_v42;
        _1816_v42 = _dafny.Set.fromElements(p2);
        let _1817_v43;
        _1817_v43 = _module.D0.create_DC0(p3);
        let _1818_v44;
        _1818_v44 = _module.D14.create_DC26(_1817_v43, _1816_v42);
        let _1819_v45;
        _1819_v45 = _dafny.Set.fromElements((_1816_v42).Intersect(_1816_v42), (_1818_v44).dtor_cf44, _1816_v42, _dafny.Set.fromElements(false));
        _1819_v45 = _1819_v45;
      }
      let _1820_i2;
      _1820_i2 = _dafny.ZERO;
      L9: {
        while (_module.__default.fm5(globalState)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1820_i2)) {
              break L9;
            }
            _1820_i2 = (_1820_i2).plus(_dafny.ONE);
            r2 = p0;
            let _1821_v46;
            _1821_v46 = _dafny.Seq.of(new BigNumber(-589));
            r0 = ((_1821_v46)[_module.__default.safeIndex(p0, new BigNumber((_1821_v46).length))]).isEqualTo(p0);
            r2 = _module.__default.safeDivisionInt(((_this).fm1(p2, globalState)).multipliedBy(p0), p0);
            let _1822_v47;
            _1822_v47 = _dafny.MultiSet.fromElements(_1775_v4, _1775_v4);
            let _1823_v48;
            let _nw335 = new _module.C0();
            _nw335.__ctor((_dafny.MultiSet.fromElements(_1775_v4, _1775_v4)).IsProperSubsetOf(_1822_v47), p2);
            _1823_v48 = _nw335;
          }
        }
      }
      if ((p0).isLessThanOrEqualTo(p0)) {
        let _1824_v49;
        _1824_v49 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm5(globalState)),!(p3));
        let _1825_v50;
        _1825_v50 = _dafny.Seq.UnicodeFromString("y");
        r0 = ((p2) ? ((((_1824_v49).contains(!(!(false)))) ? ((_1824_v49).get(!(!(false)))) : (p3))) : (_dafny.Seq.IsPrefixOf(_1825_v50, _1825_v50)));
        let _1826_v51;
        _1826_v51 = _dafny.Set.fromElements(false, p3);
        r0 = ((_1826_v51).Union(_1826_v51)).IsSubsetOf((_dafny.Set.fromElements(p3, p2, p2)).Difference(_dafny.Set.fromElements(true, p2, p2)));
        if (!(_dafny.areEqual(_1772_v1, _1772_v1)) || (!(p3))) {
          let _1827_v52;
          let _nw336 = new _module.C1();
          _nw336.__ctor(_1825_v50);
          _1827_v52 = _nw336;
          _1827_v52 = _1827_v52;
          r2 = p0;
          let _1828_v53;
          let _1829_v54;
          let _1830_v55;
          let _1831_v56;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector8 = (_this).m4(p2, _1825_v50, (_this).fm1(p3, globalState), !(((p3) ? (false) : (p3))), globalState);
          _out32 = _outcollector8[0];
          _out33 = _outcollector8[1];
          _out34 = _outcollector8[2];
          _out35 = _outcollector8[3];
          _1828_v53 = _out32;
          _1829_v54 = _out33;
          _1830_v55 = _out34;
          _1831_v56 = _out35;
          _1831_v56 = _1831_v56;
          let _1832_v57;
          _1832_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1772_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1831_v56,_1775_v4)).length));
          let _1833_v58;
          _1833_v58 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1834_v59;
          _1834_v59 = _dafny.Set.fromElements(p0, new BigNumber((_1772_v1).length), _1830_v55);
          let _1835_v60;
          let _nw337 = Array((new BigNumber(25)).toNumber());
          _nw337[(_dafny.ZERO).toNumber()] = (_1832_v57).equals(_1832_v57);
          _nw337[(_dafny.ONE).toNumber()] = !(p2) || (p3);
          _nw337[(new BigNumber(2)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(3)).toNumber()] = ((p3) ? (p2) : (_1831_v56));
          _nw337[(new BigNumber(4)).toNumber()] = (!(p2)) === (false);
          _nw337[(new BigNumber(5)).toNumber()] = ((!(p2)) ? (p2) : (_module.__default.fm5(globalState)));
          _nw337[(new BigNumber(6)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(7)).toNumber()] = !(_1831_v56);
          _nw337[(new BigNumber(8)).toNumber()] = p2;
          _nw337[(new BigNumber(9)).toNumber()] = (((_1824_v49).contains(p2)) ? ((_1824_v49).get(p2)) : (p3));
          _nw337[(new BigNumber(10)).toNumber()] = p2;
          _nw337[(new BigNumber(11)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(12)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(13)).toNumber()] = p3;
          _nw337[(new BigNumber(14)).toNumber()] = p3;
          _nw337[(new BigNumber(15)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(16)).toNumber()] = p3;
          _nw337[(new BigNumber(17)).toNumber()] = p3;
          _nw337[(new BigNumber(18)).toNumber()] = !_dafny.Seq.contains(_1825_v50, _1833_v58);
          _nw337[(new BigNumber(19)).toNumber()] = (_dafny.Set.fromElements(_1830_v55)).IsDisjointFrom(_1834_v59);
          _nw337[(new BigNumber(20)).toNumber()] = p2;
          _nw337[(new BigNumber(21)).toNumber()] = !(p2) || (p2);
          _nw337[(new BigNumber(22)).toNumber()] = !(!_dafny.Seq.contains(_1772_v1, p3));
          _nw337[(new BigNumber(23)).toNumber()] = _1831_v56;
          _nw337[(new BigNumber(24)).toNumber()] = p3;
          _1835_v60 = _nw337;
          let _index335 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1835_v60).length));
          (_1835_v60)[_index335] = true;
          let _index336 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1835_v60).length));
          (_1835_v60)[_index336] = (_1826_v51).IsSubsetOf(_1826_v51);
        } else {
          let _index337 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index337] = _module.__default.safeModuloInt(new BigNumber(212), p0);
          let _1836_v61;
          _1836_v61 = _dafny.Seq.of(_1826_v51);
          let _index338 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index338] = ((p2) ? (((p2) ? (p0) : (p0))) : (((p3) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(404)), ((_1837_p0) => function (_1838_i3) {
            return _1837_p0;
          })(p0))).length))) : (new BigNumber(((_1836_v61)[_module.__default.safeIndex(p0, new BigNumber((_1836_v61).length))]).length)))));
          let _1839_v62;
          let _out36;
          _out36 = _module.__default.m0(p2, p2, globalState);
          _1839_v62 = _out36;
          _1774_v3 = (_1774_v3).update(_module.__default.safeModuloInt(new BigNumber(-585), (_1775_v4)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1775_v4).length))]), p0);
          let _1840_v63;
          let _nw338 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _1840_v63 = _nw338;
          let _1841_v64;
          _1841_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1840_v63,!(p3));
          _1841_v64 = (_1841_v64).update(_1840_v63, false);
          let _1842_v65;
          _1842_v65 = new _dafny.CodePoint('e'.codePointAt(0));
          _1842_v65 = _1842_v65;
        }
        r2 = ((p3) ? ((_dafny.ZERO).minus(p0)) : (p0));
        let _1843_v66;
        _1843_v66 = _dafny.MultiSet.fromElements(_module.__default.fm5(globalState));
        if ((_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), function (_1844_i4) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }), new _dafny.CodePoint('q'.codePointAt(0)))) || (((((_1843_v66).contains(p2)) ? ((_1843_v66).get(p2)) : (p0))).isLessThan(p0))) {
          let _1845_v67;
          _1845_v67 = _module.D11.create_DC21(new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber(65), p0);
          let _1846_v68;
          _1846_v68 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1847_v70;
          _1847_v70 = _dafny.Set.fromElements((_this).fm1(p3, globalState), p0, p0, p0);
          let _pat_let_tv36 = _1846_v68;
          let _1848_v71;
          let _nw339 = Array((new BigNumber(21)).toNumber());
          _nw339[(_dafny.ZERO).toNumber()] = _1845_v67;
          _nw339[(_dafny.ONE).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(2)).toNumber()] = ((p2) ? (_1845_v67) : (_1845_v67));
          _nw339[(new BigNumber(3)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(4)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(5)).toNumber()] = ((p2) ? (function (_pat_let39_0) {
            return function (_1849_dt__update__tmp_h2) {
              return function (_pat_let40_0) {
                return function (_1850_dt__update_hcf36_h0) {
                  return _module.D11.create_DC21(_1850_dt__update_hcf36_h0, (_1849_dt__update__tmp_h2).dtor_cf37, (_1849_dt__update__tmp_h2).dtor_cf38);
                }(_pat_let40_0);
              }(_pat_let_tv36);
            }(_pat_let39_0);
          }(_1845_v67)) : (_module.D11.create_DC21(_1846_v68, p0, p0)));
          _nw339[(new BigNumber(6)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(7)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(8)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(9)).toNumber()] = _module.D11.create_DC21(_1846_v68, p0, p0);
          _nw339[(new BigNumber(10)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(11)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(12)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(13)).toNumber()] = _module.__default.fm46(function () {
            let _coll62 = new _dafny.Set();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(424), new BigNumber(-760))) {
              let _1851_v69 = _compr_62;
              if (((new BigNumber(424)).isLessThanOrEqualTo(_1851_v69)) && ((_1851_v69).isLessThan(new BigNumber(-760)))) {
                _coll62.add(_module.__default.safeModuloInt(_1851_v69, (_dafny.ZERO).minus(p0)));
              }
            }
            return _coll62;
          }(), globalState);
          _nw339[(new BigNumber(14)).toNumber()] = _module.D11.create_DC21(_1846_v68, p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1772_v1).length))).length),new BigNumber(836))).length));
          _nw339[(new BigNumber(15)).toNumber()] = _module.D11.create_DC21(new _dafny.CodePoint('o'.codePointAt(0)), p0, p0);
          _nw339[(new BigNumber(16)).toNumber()] = ((p3) ? (_1845_v67) : (_1845_v67));
          _nw339[(new BigNumber(17)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(18)).toNumber()] = _1845_v67;
          _nw339[(new BigNumber(19)).toNumber()] = _module.D11.create_DC21(_1846_v68, p0, p0);
          _nw339[(new BigNumber(20)).toNumber()] = _module.__default.fm46(_1847_v70, globalState);
          _1848_v71 = _nw339;
          let _index339 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1848_v71).length));
          (_1848_v71)[_index339] = _1845_v67;
          let _index340 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index340] = p0;
          let _1852_v72;
          _1852_v72 = _module.D4.create_DC6(p0, _1775_v4, p0);
          let _index341 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1848_v71).length));
          let _index342 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1775_v4).length));
          let _rhs312 = _1845_v67;
          let _rhs313 = _module.__default.safeModuloInt((_1852_v72).dtor_cf9, new BigNumber((_1774_v3).length));
          let _lhs160 = _1848_v71;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1848_v71).length));
          let _lhs162 = _1775_v4;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1775_v4).length));
          _lhs160[_lhs161] = _rhs312;
          _lhs162[_lhs163] = _rhs313;
          let _1853_v73;
          _1853_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1773_v2).Merge(_1773_v2),new BigNumber((_1825_v50).length));
          _1853_v73 = (_1853_v73).update(_1773_v2, new BigNumber(161));
          let _1854_v74;
          let _init38 = ((_1855_v68) => function (_1856_i5) {
            return _1855_v68;
          })(_1846_v68);
          let _nw340 = Array((new BigNumber(9)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw340.length); _i0_38++) {
            _nw340[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1854_v74 = _nw340;
          let _index343 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1854_v74).length));
          (_1854_v74)[_index343] = _module.__default.fm23(globalState);
          let _index344 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1854_v74).length));
          (_1854_v74)[_index344] = _module.__default.fm37((new BigNumber((_1825_v50).length)).plus((_1775_v4)[_module.__default.safeIndex(new BigNumber(134), new BigNumber((_1775_v4).length))]), globalState);
          r0 = (p0).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_1825_v50).length)));
          let _1857_v75;
          let _nw341 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
          _1857_v75 = _nw341;
          _1857_v75 = ((!(p2)) ? (_1857_v75) : (_1857_v75));
        } else {
          let _1858_v76;
          _1858_v76 = _dafny.Seq.of(p0);
          r0 = (_module.__default.fm41(_1858_v76, p2, p2, p2, globalState)).IsSubsetOf(_1826_v51);
          let _index345 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((p1).length));
          (p1)[_index345] = _dafny.Seq.Concat(_1858_v76, _dafny.Seq.of(p0));
          let _index346 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((p1).length));
          (p1)[_index346] = _1858_v76;
          r2 = (new BigNumber(131)).minus(p0);
          _1771_v0 = (_1771_v0).update((((_1773_v2).contains(p0)) ? ((_1773_v2).get(p0)) : (p2)), p0);
          r0 = !(p2);
        }
      } else {
        r2 = (_dafny.ZERO).minus(p0);
        r2 = (_this).fm1(p2, globalState);
        let _1859_v77;
        _1859_v77 = _module.D0.create_DC0(p2);
        let _1860_v78;
        _1860_v78 = _dafny.Set.fromElements(p2);
        let _1861_v79;
        _1861_v79 = _module.D14.create_DC26(_1859_v77, _1860_v78);
        let _1862_v80;
        _1862_v80 = _dafny.Set.fromElements(_1861_v79, _1861_v79);
        let _1863_v81;
        _1863_v81 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1862_v80);
        if ((_1862_v80).IsSubsetOf((((_1863_v81).contains(!(false))) ? ((_1863_v81).get(!(false))) : (_1862_v80)))) {
          let _index347 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index347] = (p0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("mig")).length));
          let _1864_v82;
          _1864_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1775_v4,p0);
          let _1865_v83;
          _1865_v83 = _dafny.Set.fromElements(p0);
          let _index348 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length));
          let _rhs314 = ((_this).fm1(p2, globalState)).minus(p0);
          let _rhs315 = new BigNumber(((_1864_v82).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1775_v4,p0)).update(_1775_v4, new BigNumber((_1865_v83).length)))).length);
          let _rhs316 = p3;
          let _rhs317 = p0;
          let _lhs164 = _1775_v4;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length));
          r2 = _rhs314;
          _lhs164[_lhs165] = _rhs315;
          r0 = _rhs316;
          r2 = _rhs317;
          let _1866_v84;
          _1866_v84 = _dafny.Seq.UnicodeFromString("kaqatlgk");
          let _1867_v85;
          _1867_v85 = _dafny.Seq.of((_1775_v4)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length))], _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(292))).length)), ((_dafny.ZERO).minus(p0)).multipliedBy(p0), _module.__default.safeModuloInt(new BigNumber((_1866_v84).length), new BigNumber(785)));
          _1867_v85 = _dafny.Seq.of(((_1775_v4)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length))]).minus((_1775_v4)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length))]), p0, p0);
          let _1868_v86;
          let _nw342 = new _module.C0();
          _nw342.__ctor(p3, p3);
          _1868_v86 = _nw342;
          let _index349 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index349] = _module.__default.safeModuloInt(((_1868_v86.f5) ? ((_this).fm1(_1868_v86.f5, globalState)) : ((_1775_v4)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length))])), (_1775_v4)[_module.__default.safeIndex(new BigNumber(589), new BigNumber((_1775_v4).length))]);
          let _1869_v87;
          let _nw343 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _1869_v87 = _nw343;
          _1869_v87 = _1869_v87;
        } else {
          let _1870_v88;
          let _init39 = ((_1871_p3) => function (_1872_i6) {
            return _1871_p3;
          })(p3);
          let _nw344 = Array((new BigNumber(12)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw344.length); _i0_39++) {
            _nw344[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1870_v88 = _nw344;
          let _index350 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          (_1870_v88)[_index350] = p3;
          let _1873_v89;
          _1873_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1860_v78,p2);
          let _index351 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          (_1870_v88)[_index351] = !((((_1873_v89).contains((_dafny.Set.fromElements(p2)).Union(_1860_v78))) ? ((_1873_v89).get((_dafny.Set.fromElements(p2)).Union(_1860_v78))) : (!(_module.__default.fm5(globalState)))));
          let _1874_v90;
          let _init40 = function (_1875_i7) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          };
          let _nw345 = Array((new BigNumber(28)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw345.length); _i0_40++) {
            _nw345[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1874_v90 = _nw345;
          let _1876_v91;
          _1876_v91 = new _dafny.CodePoint('w'.codePointAt(0));
          let _index352 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_1874_v90).length));
          (_1874_v90)[_index352] = _1876_v91;
          let _index353 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_1874_v90).length));
          (_1874_v90)[_index353] = _1876_v91;
          let _index354 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          (_1870_v88)[_index354] = p3;
          let _1877_v92;
          _1877_v92 = _dafny.MultiSet.fromElements(!(p3));
          let _index355 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          let _index356 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          let _rhs318 = _module.__default.fm5(globalState);
          let _rhs319 = (p0).isLessThanOrEqualTo((p0).plus(new BigNumber(144)));
          let _rhs320 = (_1877_v92).equals(_1877_v92);
          let _lhs166 = _1870_v88;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          let _lhs168 = _1870_v88;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length));
          _lhs166[_lhs167] = _rhs318;
          r0 = _rhs319;
          _lhs168[_lhs169] = _rhs320;
          let _1878_v93;
          _1878_v93 = _module.D6.create_DC13(p2, p0);
          r0 = ((_1878_v93).dtor_cf24) && ((_1870_v88)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_1870_v88).length))]);
        }
        let _1879_v94;
        let _init41 = ((_1880_p0) => function (_1881_i8) {
          return (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1880_p0,new _dafny.CodePoint('o'.codePointAt(0)))).length), _1880_p0, _1880_p0, _1880_p0, _1880_p0)).Intersect(_dafny.Set.fromElements(_1880_p0));
        })(p0);
        let _nw346 = Array((new BigNumber(22)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw346.length); _i0_41++) {
          _nw346[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1879_v94 = _nw346;
        _1879_v94 = _1879_v94;
        if (p2) {
          r0 = _module.__default.fm5(globalState);
          r2 = p0;
          let _1882_v97;
          _1882_v97 = _module.D14.create_DC27(false);
          let _1883_v98;
          _1883_v98 = _dafny.Map.Empty.slice().updateUnsafe((((_1771_v0).contains(p3)) ? ((_1771_v0).get(p3)) : (p0)),_1882_v97);
          let _1884_v100;
          let _nw347 = Array((new BigNumber(22)).toNumber());
          _nw347[(_dafny.ZERO).toNumber()] = _1774_v3;
          _nw347[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(84)), ((_1885_p0) => function (_1886_i9) {
            return _1885_p0;
          })(p0))).length),p0);
          _nw347[(new BigNumber(2)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(3)).toNumber()] = function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of _dafny.IntegerRange(new BigNumber(696), new BigNumber(369))) {
              let _1887_v95 = _compr_63;
              if (((new BigNumber(696)).isLessThanOrEqualTo(_1887_v95)) && ((_1887_v95).isLessThan(new BigNumber(369)))) {
                _coll63.push([_module.__default.safeDivisionInt(_1887_v95, new BigNumber((_1860_v78).length)),p0]);
              }
            }
            return _coll63;
          }();
          _nw347[(new BigNumber(4)).toNumber()] = (_1774_v3).update(new BigNumber(-183), p0);
          _nw347[(new BigNumber(5)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(6)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(7)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(8)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(9)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(10)).toNumber()] = (_1774_v3).update(p0, p0);
          _nw347[(new BigNumber(11)).toNumber()] = function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(427), new BigNumber(-816))) {
              let _1888_v96 = _compr_64;
              if (((new BigNumber(427)).isLessThanOrEqualTo(_1888_v96)) && ((_1888_v96).isLessThan(new BigNumber(-816)))) {
                _coll64.push([(_1888_v96).multipliedBy((_dafny.ZERO).minus(p0)),p0]);
              }
            }
            return _coll64;
          }();
          _nw347[(new BigNumber(12)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(13)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(14)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(15)).toNumber()] = _module.__default.fm27(p0, p3, globalState);
          _nw347[(new BigNumber(16)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1883_v98).length),new BigNumber((_1860_v78).length));
          _nw347[(new BigNumber(18)).toNumber()] = function () {
            let _coll65 = new _dafny.Map();
            for (const _compr_65 of _dafny.IntegerRange(new BigNumber(876), new BigNumber(548))) {
              let _1889_v99 = _compr_65;
              if (((new BigNumber(876)).isLessThanOrEqualTo(_1889_v99)) && ((_1889_v99).isLessThan(new BigNumber(548)))) {
                _coll65.push([(_1889_v99).plus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).length)),new BigNumber((_dafny.Seq.UnicodeFromString("gg")).length)]);
              }
            }
            return _coll65;
          }();
          _nw347[(new BigNumber(19)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(20)).toNumber()] = _1774_v3;
          _nw347[(new BigNumber(21)).toNumber()] = _1774_v3;
          _1884_v100 = _nw347;
          let _1890_v101;
          _1890_v101 = _1884_v100;
          _1890_v101 = _1890_v101;
          let _index357 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index357] = p0;
          let _index358 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index358] = p0;
          let _1891_v102;
          _1891_v102 = _dafny.Set.fromElements((_1775_v4)[_module.__default.safeIndex(new BigNumber(453), new BigNumber((_1775_v4).length))]);
          let _index359 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index359] = _module.__default.safeDivisionInt(p0, new BigNumber((_dafny.MultiSet.fromElements(_1891_v102, _1891_v102)).cardinality()));
        } else {
          let _index360 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index360] = ((p2) ? (p0) : (new BigNumber(835)));
          let _1892_v103;
          _1892_v103 = _dafny.Seq.of(p0);
          let _1893_v104;
          _1893_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1892_v103,p3);
          let _index361 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index361] = new BigNumber((_1893_v104).length);
          let _index362 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1775_v4).length));
          (_1775_v4)[_index362] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, (_1775_v4)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_1775_v4).length))]));
          let _1894_v105;
          let _nw348 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1894_v105 = _nw348;
          let _init42 = ((_1895_v4) => function (_1896_i10) {
            return (_1896_i10).minus((_1895_v4)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_1895_v4).length))]);
          })(_1775_v4);
          let _nw349 = Array((new BigNumber(14)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw349.length); _i0_42++) {
            _nw349[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1894_v105 = _nw349;
          let _1897_v106;
          _1897_v106 = _dafny.MultiSet.fromElements(p3, !(p2));
          _1897_v106 = _module.__default.fm38((_dafny.Set.fromElements(new BigNumber((_1897_v106).cardinality()))).IsSubsetOf(_dafny.Set.fromElements(p0)), globalState);
          r0 = p3;
        }
      }
      let _1898_v107;
      _1898_v107 = _module.D6.create_DC13(false, p0);
      let _pat_let_tv37 = p2;
      let _pat_let_tv38 = p2;
      r0 = !(function (_source25) {
        if (_source25.is_DC12) {
          let _1899___mcc_h0 = (_source25).cf22;
          let _1900___mcc_h1 = (_source25).cf23;
          let _1901_cf23 = _1900___mcc_h1;
          let _1902_cf22 = _1899___mcc_h0;
          return _pat_let_tv37;
        } else if (_source25.is_DC13) {
          let _1903___mcc_h2 = (_source25).cf24;
          let _1904___mcc_h3 = (_source25).cf25;
          let _1905_cf25 = _1904___mcc_h3;
          let _1906_cf24 = _1903___mcc_h2;
          return !(_1906_cf24);
        } else {
          let _1907___mcc_h4 = (_source25).cf21;
          let _1908_cf21 = _1907___mcc_h4;
          return _pat_let_tv38;
        }
      }(_1898_v107));
      let _nw350 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
      r1 = _nw350;
      r2 = _module.__default.safeModuloInt(p0, (_this).fm1(p2, globalState));
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1909_v0;
      _1909_v0 = false;
      let _1910_v1;
      let _nw351 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1910_v1 = _nw351;
      let _index363 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
      (_1910_v1)[_index363] = new BigNumber(-432);
      let _index364 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
      let _rhs321 = (p1).isEqualTo((p0).minus(new BigNumber(189)));
      let _rhs322 = p0;
      let _lhs170 = _1910_v1;
      let _lhs171 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
      _1909_v0 = _rhs321;
      _lhs170[_lhs171] = _rhs322;
      let _1911_i0;
      _1911_i0 = _dafny.ZERO;
      L10: {
        while (_1909_v0) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1911_i0)) {
              break L10;
            }
            _1911_i0 = (_1911_i0).plus(_dafny.ONE);
            _1910_v1 = _1910_v1;
            let _1912_v2;
            _1912_v2 = _dafny.MultiSet.fromElements(_1909_v0);
            if (((_dafny.MultiSet.fromElements(p2)).Intersect(_1912_v2)).IsSubsetOf((_1912_v2).Union(_1912_v2))) {
              let _1913_v3;
              _1913_v3 = new _dafny.CodePoint('h'.codePointAt(0));
              let _1914_v4;
              _1914_v4 = _dafny.MultiSet.fromElements(_1913_v3);
              let _1915_v5;
              _1915_v5 = _dafny.Seq.of(new BigNumber((_1914_v4).cardinality()), (_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))]);
              let _1916_v6;
              _1916_v6 = _dafny.Seq.UnicodeFromString("dsqqc");
              let _1917_v7;
              _1917_v7 = _dafny.Seq.of(_1916_v6);
              let _1918_v8;
              _1918_v8 = _dafny.Seq.of(_1916_v6, (_1917_v7)[_module.__default.safeIndex((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], new BigNumber((_1917_v7).length))]);
              let _1919_v9;
              _1919_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1917_v7)[_module.__default.safeIndex(p1, new BigNumber((_1917_v7).length))]);
              let _1920_v10;
              _1920_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1915_v5)[_module.__default.safeIndex(new BigNumber((_1919_v9).length), new BigNumber((_1915_v5).length))],new BigNumber((_1916_v6).length));
              let _1921_v11;
              _1921_v11 = _dafny.Set.fromElements(p2, _1909_v0, _1909_v0);
              let _1922_v12;
              _1922_v12 = _dafny.Seq.of(_1921_v11, (_1921_v11).Difference(_1921_v11));
              let _index365 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
              let _rhs323 = _dafny.Map.Empty.slice().updateUnsafe((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))],p0);
              let _rhs324 = _1922_v12;
              let _rhs325 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm47((_module.D11.create_DC21(_1913_v3, p1, p0)).dtor_cf36, p1, _module.__default.safeDivisionInt((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], p0), globalState)).length));
              let _rhs326 = _1909_v0;
              let _rhs327 = p2;
              let _lhs172 = _1910_v1;
              let _lhs173 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
              _1920_v10 = _rhs323;
              _1922_v12 = _rhs324;
              _lhs172[_lhs173] = _rhs325;
              _1909_v0 = _rhs326;
              _1909_v0 = _rhs327;
              let _1923_v13;
              _1923_v13 = _dafny.MultiSet.fromElements(_1910_v1);
              let _1924_v14;
              _1924_v14 = _module.D5.create_DC8(_1923_v13);
              let _1925_v15;
              _1925_v15 = _dafny.Set.fromElements(_1924_v14, _1924_v14);
              _1925_v15 = _1925_v15;
              let _1926_v16;
              _1926_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
              _1926_v16 = (_1926_v16).update(_1909_v0, !(_1920_v10).contains(p0));
              r0 = (_dafny.ZERO).minus((((_1912_v2).contains((p2) && (_1909_v0))) ? ((_1912_v2).get((p2) && (_1909_v0))) : ((_dafny.ZERO).minus(p1))));
              let _1927_v17;
              let _nw352 = Array((new BigNumber(13)).toNumber()).fill(false);
              _1927_v17 = _nw352;
              let _index366 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1927_v17).length));
              (_1927_v17)[_index366] = _1909_v0;
              let _index367 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1927_v17).length));
              (_1927_v17)[_index367] = p2;
            } else {
              let _1928_v18;
              let _nw353 = new _module.C4();
              _nw353.__ctor(p1);
              _1928_v18 = _nw353;
              let _1929_v19;
              _1929_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1928_v18,((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))]).plus((_dafny.ZERO).minus(_1928_v18.f2)));
              let _1930_v20;
              _1930_v20 = _dafny.Seq.UnicodeFromString("pftld");
              let _index368 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
              (_1910_v1)[_index368] = (((_1929_v19).contains(_1928_v18)) ? ((_1929_v19).get(_1928_v18)) : (_module.__default.safeModuloInt(_1928_v18.f2, new BigNumber((_1930_v20).length))));
              let _1931_v21;
              _1931_v21 = _dafny.Seq.of(_1912_v2, _1912_v2);
              let _1932_v22;
              _1932_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1909_v0,new BigNumber(76));
              let _1933_v23;
              _1933_v23 = _module.D5.create_DC9(!(_1909_v0), new BigNumber((_1930_v20).length), _1912_v2);
              let _1934_v24;
              _1934_v24 = _dafny.Set.fromElements(_1909_v0);
              let _1935_v25;
              _1935_v25 = _module.D4.create_DC5(_1934_v24);
              let _1936_v26;
              _1936_v26 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(((_1931_v21)[_module.__default.safeIndex(_1928_v18.f2, new BigNumber((_1931_v21).length))]).cardinality())), p0, new BigNumber((_1932_v22).length), (_1933_v23).dtor_cf15, new BigNumber(((_1935_v25).dtor_cf8).length));
              _1909_v0 = (_1936_v26).IsProperSubsetOf(_1936_v26);
              let _index369 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
              (_1910_v1)[_index369] = ((p2) ? (p0) : (new BigNumber(339)));
              let _1937_v27;
              _1937_v27 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
              let _1938_v28;
              _1938_v28 = _dafny.Seq.of(!(p2), _1909_v0, _1909_v0);
              let _1939_v29;
              let _nw354 = new _module.C0();
              _nw354.__ctor(((p2) ? ((((_1937_v27).contains(_1909_v0)) ? ((_1937_v27).get(_1909_v0)) : (_1909_v0))) : (!(_1909_v0))), (_1938_v28)[_module.__default.safeIndex(p1, new BigNumber((_1938_v28).length))]);
              _1939_v29 = _nw354;
              (_1939_v29).f4 = !(_module.__default.safeModuloInt((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], p0)).isEqualTo(p1);
            }
            let _1940_v30;
            _1940_v30 = _dafny.Seq.UnicodeFromString("nfqqmub");
            let _1941_v31;
            _1941_v31 = _dafny.MultiSet.fromElements((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], p1);
            let _1942_v32;
            _1942_v32 = _dafny.Seq.of((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], p1, new BigNumber((_1941_v31).cardinality()));
            let _1943_v33;
            _1943_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1909_v0,_1942_v32);
            let _1944_v34;
            _1944_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1912_v2).cardinality()),p2);
            let _1945_v35;
            _1945_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1940_v30,_1944_v34);
            let _1946_v36;
            _1946_v36 = _module.D5.create_DC10(_1909_v0, p1, _1945_v35, _1909_v0);
            let _1947_v37;
            let _nw355 = new _module.C1();
            _nw355.__ctor(_dafny.Seq.update(_1940_v30, _module.__default.safeIndex(new BigNumber((_1943_v33).length), new BigNumber((_1940_v30).length)), _module.__default.fm37((_1946_v36).dtor_cf18, globalState)));
            _1947_v37 = _nw355;
            let _index370 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
            (_1910_v1)[_index370] = ((_1909_v0) ? ((_dafny.ZERO).minus(p1)) : (p1));
          }
        }
      }
      let _1948_v38;
      let _init43 = function (_1949_i1) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jcqb"), _dafny.Seq.UnicodeFromString("cxvrnpqcq"));
      };
      let _nw356 = Array((new BigNumber(23)).toNumber());
      for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw356.length); _i0_43++) {
        _nw356[_i0_43] = _init43(new BigNumber(_i0_43));
      }
      _1948_v38 = _nw356;
      let _1950_v39;
      _1950_v39 = _dafny.Seq.UnicodeFromString("paxuh");
      let _index371 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1948_v38).length));
      (_1948_v38)[_index371] = _1950_v39;
      let _index372 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1948_v38).length));
      (_1948_v38)[_index372] = _dafny.Seq.Concat(_1950_v39, _1950_v39);
      let _index373 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length));
      (_1910_v1)[_index373] = p0;
      let _1951_v40;
      _1951_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1950_v39).length),new BigNumber((_dafny.MultiSet.fromElements(p2, p2)).cardinality()));
      let _1952_v41;
      let _nw357 = new _module.C4();
      _nw357.__ctor((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))]);
      _1952_v41 = _nw357;
      let _1953_v42;
      _1953_v42 = _dafny.Seq.of((_1910_v1)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1910_v1).length))], (((_1951_v40).contains(new BigNumber((_dafny.Seq.of(_1952_v41, _1952_v41)).length))) ? ((_1951_v40).get(new BigNumber((_dafny.Seq.of(_1952_v41, _1952_v41)).length))) : (p1)));
      let _1954_v43;
      let _nw358 = new _module.C2();
      _nw358.__ctor(p1);
      _1954_v43 = _nw358;
      let _1955_v44;
      _1955_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1909_v0,_1954_v43);
      let _1956_v45;
      _1956_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1909_v0,(_1953_v42)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1955_v44).length)), new BigNumber((_1953_v42).length))]);
      let _1957_v46;
      _1957_v46 = _dafny.Seq.of(true);
      _1956_v45 = (_1956_v45).update((_1957_v46)[_module.__default.safeIndex(p1, new BigNumber((_1957_v46).length))], p1);
      let _1958_v47;
      let _1959_v48;
      let _out37;
      let _out38;
      let _outcollector9 = (_1954_v43).m6(new BigNumber(((_1948_v38)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1948_v38).length))]).length), globalState);
      _out37 = _outcollector9[0];
      _out38 = _outcollector9[1];
      _1958_v47 = _out37;
      _1959_v48 = _out38;
      r0 = p1;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _1960_v0;
      let _nw359 = new _module.C1();
      _nw359.__ctor(_dafny.Seq.UnicodeFromString("ui"));
      _1960_v0 = _nw359;
      let _1961_v1;
      _1961_v1 = _module.D8.create_DC17(p2, _1960_v0, (_dafny.ZERO).minus(new BigNumber((p0).cardinality())));
      let _source26 = _1961_v1;
      if (_source26.is_DC17) {
        let _1962___mcc_h0 = (_source26).cf30;
        let _1963___mcc_h1 = (_source26).cf31;
        let _1964___mcc_h2 = (_source26).cf32;
        let _1965_cf32 = _1964___mcc_h2;
        let _1966_cf31 = _1963___mcc_h1;
        let _1967_cf30 = _1962___mcc_h0;
        let _1968_v2;
        _1968_v2 = _module.D4.create_DC6(new BigNumber(862), p1, _1965_cf32);
        let _1969_v3;
        _1969_v3 = _module.D4.create_DC7(_1968_v2);
        let _1970_v4;
        _1970_v4 = _dafny.Seq.of(_1968_v2);
        let _1971_v5;
        _1971_v5 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),p3);
        let _1972_v6;
        _1972_v6 = _module.D4.create_DC7((_1970_v4)[_module.__default.safeIndex(new BigNumber((_1971_v5).length), new BigNumber((_1970_v4).length))]);
        let _1973_v7;
        _1973_v7 = _module.D4.create_DC7((_1972_v6).dtor_cf12);
        let _source27 = _1973_v7;
        if (_source27.is_DC6) {
          let _1974___mcc_h4 = (_source27).cf9;
          let _1975___mcc_h5 = (_source27).cf10;
          let _1976___mcc_h6 = (_source27).cf11;
          let _1977_cf11 = _1976___mcc_h6;
          let _1978_cf10 = _1975___mcc_h5;
          let _1979_cf9 = _1974___mcc_h4;
          let _1980_v8;
          _1980_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1979_cf9).isLessThanOrEqualTo(_1977_cf11),new _dafny.CodePoint('u'.codePointAt(0)));
          let _rhs328 = p3;
          let _rhs329 = _1980_v8;
          _1967_cf30 = _rhs328;
          _1980_v8 = _rhs329;
          let _1981_v9;
          _1981_v9 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1982_v10;
          let _init44 = ((_1983_p3) => function (_1984_i0) {
            return !(_1983_p3);
          })(p3);
          let _nw360 = Array((new BigNumber(17)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw360.length); _i0_44++) {
            _nw360[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1982_v10 = _nw360;
          let _1985_v11;
          let _nw361 = new _module.C5();
          _nw361.__ctor(((_1967_cf30) ? (new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("wywgutya"), _module.__default.safeIndex(_1979_cf9, new BigNumber((_dafny.Seq.UnicodeFromString("wywgutya")).length)), _1981_v9)).length)) : ((_dafny.ZERO).minus(_1979_cf9))), _1982_v10);
          _1985_v11 = _nw361;
          let _1986_v12;
          let _init45 = ((_1987_v0) => function (_1988_i1) {
            return ((_module.D15.create_DC29(_dafny.Map.Empty.slice().updateUnsafe((_1987_v0).f3,_dafny.Seq.UnicodeFromString("luwy")))).dtor_cf46).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1987_v0).f3,_dafny.Seq.UnicodeFromString("avaemdf")));
          })(_1960_v0);
          let _nw362 = Array((new BigNumber(21)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw362.length); _i0_45++) {
            _nw362[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1986_v12 = _nw362;
          let _1989_v13;
          _1989_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1966_cf31).f3,_dafny.Seq.UnicodeFromString("qgu"));
          let _index374 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1986_v12).length));
          (_1986_v12)[_index374] = _1989_v13;
          let _index375 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1986_v12).length));
          (_1986_v12)[_index375] = _1989_v13;
          _1967_cf30 = p3;
        } else if (_source27.is_DC5) {
          let _1990___mcc_h7 = (_source27).cf8;
          let _1991_cf8 = _1990___mcc_h7;
          r0 = _1965_cf32;
          let _1992_v14;
          let _nw363 = new _module.C4();
          _nw363.__ctor(_1965_cf32);
          _1992_v14 = _nw363;
          let _1993_v15;
          _1993_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1992_v14,_1992_v14.f2);
          let _1994_v16;
          _1994_v16 = _dafny.Seq.of((_dafny.ZERO).minus(_1965_cf32), _1965_cf32, new BigNumber(((_1993_v15).update(_1992_v14, _1992_v14.f2)).length), new BigNumber(232), _1965_cf32);
          _1965_cf32 = (((new BigNumber((_1994_v16).length)).isEqualTo(_1992_v14.f2)) ? (_1992_v14.f2) : (new BigNumber(((_1960_v0).f3).length)));
          let _1995_v17;
          _1995_v17 = _dafny.Set.fromElements((_1960_v0).fm1(_1967_cf30, globalState));
          let _rhs330 = (_1965_cf32).isEqualTo(_1965_cf32);
          let _rhs331 = (_1995_v17).Intersect(_1995_v17);
          _1967_cf30 = _rhs330;
          _1995_v17 = _rhs331;
          let _1996_v18;
          _1996_v18 = _dafny.Seq.of(p3, _1967_cf30, !(p3), _1967_cf30);
          let _1997_v19;
          _1997_v19 = _dafny.Seq.of(_1996_v18, _module.__default.fm16(globalState));
          let _rhs332 = _1992_v14.f2;
          let _rhs333 = (_dafny.ZERO).minus(_1965_cf32);
          let _rhs334 = _dafny.Seq.Concat(_1997_v19, _1997_v19);
          let _rhs335 = p2;
          let _rhs336 = false;
          let _lhs174 = _1992_v14;
          r0 = _rhs332;
          _lhs174.f2 = _rhs333;
          _1997_v19 = _rhs334;
          _1967_cf30 = _rhs335;
          _1967_cf30 = _rhs336;
        } else {
          let _1998___mcc_h8 = (_source27).cf12;
          let _1999_cf12 = _1998___mcc_h8;
          let _2000_v20;
          _2000_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1965_cf32,((p3) ? (_1965_cf32) : (_1965_cf32)));
          _2000_v20 = (_2000_v20).update(new BigNumber((_module.__default.fm38(p3, globalState)).cardinality()), _1965_cf32);
          let _2001_v21;
          let _init46 = ((_2002_p2) => function (_2003_i2) {
            return !(_2002_p2);
          })(p2);
          let _nw364 = Array((new BigNumber(26)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw364.length); _i0_46++) {
            _nw364[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _2001_v21 = _nw364;
          _2001_v21 = _2001_v21;
          let _2004_v22;
          _2004_v22 = _dafny.Seq.of(p3, p2);
          let _2005_v23;
          let _nw365 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _2005_v23 = _nw365;
          let _2006_v24;
          let _2007_v25;
          let _2008_v26;
          let _out39;
          let _out40;
          let _out41;
          let _outcollector10 = (_1960_v0).m1((new BigNumber((_dafny.Seq.update(_2004_v22, _module.__default.safeIndex(new BigNumber(552), new BigNumber((_2004_v22).length)), p2)).length)).minus(new BigNumber(-947)), _2005_v23, p3, p3, globalState);
          _out39 = _outcollector10[0];
          _out40 = _outcollector10[1];
          _out41 = _outcollector10[2];
          _2006_v24 = _out39;
          _2007_v25 = _out40;
          _2008_v26 = _out41;
          let _2009_v27;
          _2009_v27 = _dafny.Set.fromElements(_1967_cf30);
          let _2010_v28;
          _2010_v28 = _dafny.Seq.of(_dafny.Set.fromElements(_2006_v24), _2009_v27, _2009_v27, _2009_v27, _2009_v27);
          let _2011_v29;
          _2011_v29 = _dafny.Seq.UnicodeFromString("cbfxcksb");
          let _rhs337 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(p2), _2009_v27), _2010_v28), _module.__default.safeIndex(_1965_cf32, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(p2), _2009_v27), _2010_v28)).length)), _2009_v27);
          let _rhs338 = _module.__default.fm15(globalState);
          let _rhs339 = (new BigNumber(((_1960_v0).f3).length)).multipliedBy(_1965_cf32);
          _2010_v28 = _rhs337;
          _2011_v29 = _rhs338;
          r0 = _rhs339;
        }
        r0 = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_1967_cf30)).length)))).length);
        let _2012_v30;
        let _nw366 = Array((new BigNumber(3)).toNumber()).fill(false);
        _2012_v30 = _nw366;
        let _index376 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2012_v30).length));
        (_2012_v30)[_index376] = p3;
        let _index377 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2012_v30).length));
        (_2012_v30)[_index377] = p2;
        r1 = new BigNumber(-385);
      } else {
        let _2013___mcc_h3 = (_source26).cf29;
        let _2014_cf29 = _2013___mcc_h3;
        let _2015_v31;
        let _nw367 = Array((new BigNumber(21)).toNumber()).fill(false);
        _2015_v31 = _nw367;
        let _2016_v32;
        _2016_v32 = new BigNumber(246);
        let _index378 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2015_v31).length));
        (_2015_v31)[_index378] = (new BigNumber(((_1960_v0).f3).length)).isLessThan(_2016_v32);
        let _2017_v33;
        _2017_v33 = true;
        let _2018_v34;
        _2018_v34 = _dafny.Map.Empty.slice().updateUnsafe((_2016_v32).isLessThanOrEqualTo(_2016_v32),!(p3));
        let _2019_v35;
        _2019_v35 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2020_v36;
        _2020_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2019_v35,new BigNumber(-702));
        let _2021_v37;
        _2021_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2020_v36).length),p3);
        let _index379 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2015_v31).length));
        let _rhs340 = (((_2018_v34).contains(p3)) ? ((_2018_v34).get(p3)) : ((((_2021_v37).contains(_2016_v32)) ? ((_2021_v37).get(_2016_v32)) : (p2))));
        let _rhs341 = _2017_v33;
        let _lhs175 = _2015_v31;
        let _lhs176 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2015_v31).length));
        _lhs175[_lhs176] = _rhs340;
        _2017_v33 = _rhs341;
        let _index380 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2015_v31).length));
        (_2015_v31)[_index380] = !((p0).IsDisjointFrom(p0));
        let _2022_v38;
        _2022_v38 = _dafny.Seq.of(p3, p2);
        let _2023_v39;
        let _out42;
        _out42 = _module.__default.m0((new BigNumber((_2022_v38).length)).isLessThanOrEqualTo((_this).fm1(_module.__default.fm5(globalState), globalState)), _dafny.Seq.IsPrefixOf((_2014_cf29).f3, _dafny.Seq.UnicodeFromString("prkosnxd")), globalState);
        _2023_v39 = _out42;
        _2017_v33 = _2017_v33;
      }
      let _2024_v40;
      _2024_v40 = true;
      let _2025_v41;
      let _init47 = function (_2026_i3) {
        return !(true);
      };
      let _nw368 = Array((new BigNumber(4)).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw368.length); _i0_47++) {
        _nw368[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _2025_v41 = _nw368;
      let _2027_v42;
      _2027_v42 = _dafny.Seq.of(_2025_v41);
      let _2028_v43;
      _2028_v43 = _module.D14.create_DC27(p3);
      let _pat_let_tv39 = _1960_v0;
      let _pat_let_tv40 = _2024_v40;
      let _pat_let_tv41 = p3;
      let _pat_let_tv42 = _1960_v0;
      let _pat_let_tv43 = p2;
      let _pat_let_tv44 = p3;
      let _pat_let_tv45 = p2;
      let _rhs342 = function (_source28) {
        if (_source28.is_DC26) {
          let _2029___mcc_h9 = (_source28).cf43;
          let _2030___mcc_h10 = (_source28).cf44;
          let _2031_cf44 = _2030___mcc_h10;
          let _2032_cf43 = _2029___mcc_h9;
          return (new BigNumber((_dafny.Set.fromElements(new BigNumber(((_pat_let_tv39).f3).length))).length)).isLessThanOrEqualTo((_module.D0.create_DC1(_pat_let_tv40, _pat_let_tv41, new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber(((_pat_let_tv42).f3).length))).dtor_cf4);
        } else if (_source28.is_DC27) {
          let _2033___mcc_h11 = (_source28).cf45;
          let _2034_cf45 = _2033___mcc_h11;
          return (new BigNumber((_dafny.Seq.of(_2034_cf45, _pat_let_tv43)).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(_2034_cf45)).length));
        } else if (_source28.is_DC28) {
          return _pat_let_tv44;
        } else {
          let _2035___mcc_h12 = (_source28).cf42;
          let _2036_cf42 = _2035___mcc_h12;
          return _pat_let_tv45;
        }
      }(_2028_v43);
      let _rhs343 = _dafny.Seq.Concat(_dafny.Seq.of(_2025_v41, _2025_v41, _2025_v41, _2025_v41), _2027_v42);
      _2024_v40 = _rhs342;
      _2027_v42 = _rhs343;
      let _2037_v44;
      _2037_v44 = _dafny.Set.fromElements(p1, p1, p1);
      let _2038_v45;
      _2038_v45 = _dafny.Seq.of(p3);
      let _2039_v46;
      _2039_v46 = _dafny.Map.Empty.slice().updateUnsafe(_2038_v45,_2024_v40);
      let _2040_v47;
      _2040_v47 = new BigNumber(420);
      let _2041_v48;
      _2041_v48 = _dafny.Set.fromElements(new BigNumber(((_2039_v46).Merge(_2039_v46)).length), _2040_v47);
      let _2042_v49;
      _2042_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p2, !(true)),_2037_v44);
      let _2043_v50;
      _2043_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1960_v0).f3).length),_2040_v47);
      let _2044_v51;
      _2044_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2025_v41,(((_2043_v50).contains(_2040_v47)) ? ((_2043_v50).get(_2040_v47)) : (_2040_v47)));
      let _rhs344 = _module.__default.safeModuloInt(_2040_v47, _2040_v47);
      let _rhs345 = ((((_2042_v49).contains((_dafny.MultiSet.FromArray(_2038_v45)).update(p3, _module.__default.abs(_2040_v47)))) ? ((_2042_v49).get((_dafny.MultiSet.FromArray(_2038_v45)).update(p3, _module.__default.abs(_2040_v47)))) : (_2037_v44))).Union(_2037_v44);
      let _rhs346 = (_module.__default.fm48(_2024_v40, globalState));
      let _rhs347 = new BigNumber((_2044_v51).length);
      r0 = _rhs344;
      _2037_v44 = _rhs345;
      _2041_v48 = _rhs346;
      r0 = _rhs347;
      let _2045_v52;
      _2045_v52 = _dafny.Seq.of(_2040_v47);
      let _2046_v53;
      _2046_v53 = _module.D7.create_DC14(_dafny.Seq.Concat(_dafny.Seq.of(_2040_v47), _2045_v52));
      let _source29 = _2046_v53;
      if (_source29.is_DC15) {
        let _2047___mcc_h13 = (_source29).cf27;
        let _2048___mcc_h14 = (_source29).cf28;
        let _2049_cf28 = _2048___mcc_h14;
        let _2050_cf27 = _2047___mcc_h13;
        let _index381 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length));
        (_2025_v41)[_index381] = _2024_v40;
        let _index382 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length));
        (_2025_v41)[_index382] = _2024_v40;
        let _2051_v54;
        let _nw369 = new _module.C4();
        _nw369.__ctor(_2050_cf27);
        _2051_v54 = _nw369;
        let _2052_v55;
        let _init48 = ((_2053_p2) => function (_2054_i4) {
          return _2053_p2;
        })(p2);
        let _nw370 = Array((new BigNumber(27)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw370.length); _i0_48++) {
          _nw370[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _2052_v55 = _nw370;
        let _2055_v56;
        let _nw371 = Array((new BigNumber(4)).toNumber());
        _nw371[(_dafny.ZERO).toNumber()] = _2052_v55;
        _nw371[(_dafny.ONE).toNumber()] = _2052_v55;
        _nw371[(new BigNumber(2)).toNumber()] = _2052_v55;
        _nw371[(new BigNumber(3)).toNumber()] = _2052_v55;
        _2055_v56 = _nw371;
        let _2056_v57;
        _2056_v57 = _module.D14.create_DC25(_2055_v56);
        let _pat_let_tv46 = _2055_v56;
        let _source30 = function (_pat_let41_0) {
          return function (_2057_dt__update__tmp_h0) {
            return function (_pat_let42_0) {
              return function (_2058_dt__update_hcf42_h0) {
                return _module.D14.create_DC25(_2058_dt__update_hcf42_h0);
              }(_pat_let42_0);
            }(_pat_let_tv46);
          }(_pat_let41_0);
        }(_2056_v57);
        if (_source30.is_DC26) {
          let _2059___mcc_h16 = (_source30).cf43;
          let _2060___mcc_h17 = (_source30).cf44;
          let _2061_cf44 = _2060___mcc_h17;
          let _2062_cf43 = _2059___mcc_h16;
          _2040_v47 = (_module.__default.safeModuloInt(_2050_cf27, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2040_v47,p3)).length))).minus(_2040_v47);
          r1 = (_2040_v47).minus((_this).fm1((_2025_v41)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length))], globalState));
          let _2063_v58;
          _2063_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          _2063_v58 = (_2063_v58).update(p1, (_2025_v41)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length))]);
          let _index383 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length));
          (_2025_v41)[_index383] = p3;
        } else if (_source30.is_DC27) {
          let _2064___mcc_h18 = (_source30).cf45;
          let _2065_cf45 = _2064___mcc_h18;
          _2040_v47 = _module.__default.safeDivisionInt(_2050_cf27, _2050_cf27);
          let _2066_v59;
          let _nw372 = new _module.C1();
          _nw372.__ctor(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-916)), ((_2067_cf28) => function (_2068_i5) {
            return _2067_cf28;
          })(_2049_cf28)), _dafny.Seq.UnicodeFromString("du")));
          _2066_v59 = _nw372;
          r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_2040_v47).minus(_2050_cf27)), (_dafny.ZERO).minus(new BigNumber((_2027_v42).length)));
          let _2069_v60;
          let _nw373 = Array((new BigNumber(6)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = _2040_v47;
          _nw373[(_dafny.ONE).toNumber()] = _2040_v47;
          _nw373[(new BigNumber(2)).toNumber()] = new BigNumber(-630);
          _nw373[(new BigNumber(3)).toNumber()] = _2050_cf27;
          _nw373[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_2049_cf28, _2049_cf28, _2049_cf28)).length);
          _nw373[(new BigNumber(5)).toNumber()] = new BigNumber(358);
          _2069_v60 = _nw373;
          _2069_v60 = p1;
        } else if (_source30.is_DC28) {
          _2024_v40 = (_2038_v45)[_module.__default.safeIndex(new BigNumber(((_1960_v0).f3).length), new BigNumber((_2038_v45).length))];
          let _2070_v61;
          _2070_v61 = p0;
          let _2071_v62;
          _2071_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2070_v61,_2049_cf28);
          let _2072_v63;
          _2072_v63 = _dafny.Set.fromElements(p3, true, _module.__default.fm5(globalState), p3);
          let _2073_v64;
          _2073_v64 = _module.D4.create_DC5(_2072_v63);
          let _2074_v65;
          _2074_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v62,_2073_v64);
          _2074_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v62,_2073_v64);
          let _2075_v66;
          let _nw374 = new _module.C1();
          _nw374.__ctor((_1960_v0).f3);
          _2075_v66 = _nw374;
          let _index384 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length));
          (_2025_v41)[_index384] = (_2025_v41)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length))];
        } else {
          let _2076___mcc_h19 = (_source30).cf42;
          let _2077_cf42 = _2076___mcc_h19;
          let _2078_v67;
          let _nw375 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _2078_v67 = _nw375;
          let _2079_v68;
          let _2080_v69;
          let _2081_v70;
          let _out43;
          let _out44;
          let _out45;
          let _outcollector11 = (_2051_v54).m1(_2050_cf27, _2078_v67, !(_dafny.Seq.IsPrefixOf(_2038_v45, _2038_v45)), p2, globalState);
          _out43 = _outcollector11[0];
          _out44 = _outcollector11[1];
          _out45 = _outcollector11[2];
          _2079_v68 = _out43;
          _2080_v69 = _out44;
          _2081_v70 = _out45;
          _2038_v45 = _2038_v45;
          let _2082_v71;
          _2082_v71 = _dafny.Seq.UnicodeFromString("vdfxyi");
          _2082_v71 = _dafny.Seq.Concat((_1960_v0).f3, _dafny.Seq.update((_1960_v0).f3, _module.__default.safeIndex(new BigNumber(-474), new BigNumber(((_1960_v0).f3).length)), _2049_cf28));
          let _2083_v72;
          let _nw376 = new _module.C3();
          _nw376.__ctor(((_this).fm1(p2, globalState)).plus(_2050_cf27));
          _2083_v72 = _nw376;
        }
        _2024_v40 = (((_2025_v41)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2025_v41).length))]) ? (p2) : (p3));
      } else {
        let _2084___mcc_h15 = (_source29).cf26;
        let _2085_cf26 = _2084___mcc_h15;
        let _2086_v73;
        _2086_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1960_v0).f3,p2);
        let _2087_v74;
        _2087_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2040_v47,_2024_v40);
        _2086_v73 = (_2086_v73).update((_1960_v0).f3, (_2024_v40) && ((((_2087_v74).contains(_2040_v47)) ? ((_2087_v74).get(_2040_v47)) : (p3))));
        let _2088_v75;
        let _nw377 = new _module.C2();
        _nw377.__ctor((_dafny.ZERO).minus(_2040_v47));
        _2088_v75 = _nw377;
        let _2089_v76;
        _2089_v76 = _dafny.Seq.UnicodeFromString("fuumxjor");
        let _2090_v77;
        _2090_v77 = new _dafny.CodePoint('x'.codePointAt(0));
        _2089_v76 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_2091_i6) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }), _dafny.Seq.update(_dafny.Seq.update(_2089_v76, _module.__default.safeIndex(new BigNumber((_2085_cf26).length), new BigNumber((_2089_v76).length)), ((_1960_v0).f3)[_module.__default.safeIndex(_2040_v47, new BigNumber(((_1960_v0).f3).length))]), _module.__default.safeIndex(_2040_v47, new BigNumber((_dafny.Seq.update(_2089_v76, _module.__default.safeIndex(new BigNumber((_2085_cf26).length), new BigNumber((_2089_v76).length)), ((_1960_v0).f3)[_module.__default.safeIndex(_2040_v47, new BigNumber(((_1960_v0).f3).length))])).length)), _2090_v77));
        _2040_v47 = _2040_v47;
      }
      let _2092_v78;
      _2092_v78 = new _dafny.CodePoint('n'.codePointAt(0));
      let _2093_v79;
      let _nw378 = Array((new BigNumber(22)).toNumber());
      _nw378[(_dafny.ZERO).toNumber()] = (_1960_v0).f3;
      _nw378[(_dafny.ONE).toNumber()] = _dafny.Seq.update((_1960_v0).f3, _module.__default.safeIndex(_2040_v47, new BigNumber(((_1960_v0).f3).length)), _2092_v78);
      _nw378[(new BigNumber(2)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(3)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("bxjtpgk");
      _nw378[(new BigNumber(5)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("bcgsqra");
      _nw378[(new BigNumber(7)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-609)), function (_2094_i7) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      });
      _nw378[(new BigNumber(9)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(10)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(11)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(12)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(365)), ((_2095_v78) => function (_2096_i8) {
        return _2095_v78;
      })(_2092_v78));
      _nw378[(new BigNumber(14)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(15)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(16)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("eoy"), _module.__default.safeIndex(_2040_v47, new BigNumber((_dafny.Seq.UnicodeFromString("eoy")).length)), new _dafny.CodePoint('l'.codePointAt(0)));
      _nw378[(new BigNumber(18)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(19)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(20)).toNumber()] = (_1960_v0).f3;
      _nw378[(new BigNumber(21)).toNumber()] = (_1960_v0).f3;
      _2093_v79 = _nw378;
      let _2097_v80;
      _2097_v80 = _dafny.Map.Empty.slice().updateUnsafe(_2040_v47,_2045_v52);
      let _2098_v81;
      let _nw379 = Array((new BigNumber(29)).toNumber());
      _nw379[(_dafny.ZERO).toNumber()] = _2045_v52;
      _nw379[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-225)), function (_2099_i9) {
        return new BigNumber(61);
      });
      _nw379[(new BigNumber(2)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(3)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(4)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(5)).toNumber()] = _module.__default.fm36(p2, _2092_v78, new BigNumber((_2041_v48).length), globalState);
      _nw379[(new BigNumber(6)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(7)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(8)).toNumber()] = _module.__default.fm36(p3, new _dafny.CodePoint('b'.codePointAt(0)), _2040_v47, globalState);
      _nw379[(new BigNumber(9)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(10)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(11)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(new BigNumber(((_1960_v0).f3).length));
      _nw379[(new BigNumber(13)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_2040_v47);
      _nw379[(new BigNumber(15)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-172)), ((_2100_v47) => function (_2101_i10) {
        return _2100_v47;
      })(_2040_v47));
      _nw379[(new BigNumber(17)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(new BigNumber((_2045_v52).length));
      _nw379[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_2045_v52, _module.__default.safeIndex(_2040_v47, new BigNumber((_2045_v52).length)), new BigNumber(5));
      _nw379[(new BigNumber(20)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(21)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(22)).toNumber()] = (((_2097_v80).contains(_2040_v47)) ? ((_2097_v80).get(_2040_v47)) : (_2045_v52));
      _nw379[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_2040_v47, _2040_v47);
      _nw379[(new BigNumber(24)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(25)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(26)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), ((_2102_v50) => function (_2103_i11) {
        return new BigNumber((_2102_v50).length);
      })(_2043_v50));
      _nw379[(new BigNumber(27)).toNumber()] = _2045_v52;
      _nw379[(new BigNumber(28)).toNumber()] = _dafny.Seq.of((_1960_v0).fm1(true, globalState));
      _2098_v81 = _nw379;
      let _2104_v82;
      _2104_v82 = _dafny.Map.Empty.slice().updateUnsafe(_2093_v79,_2098_v81);
      let _2105_v83;
      let _2106_v84;
      let _2107_v85;
      let _out46;
      let _out47;
      let _out48;
      let _outcollector12 = (_1960_v0).m1(_2040_v47, (((_2104_v82).contains(_2093_v79)) ? ((_2104_v82).get(_2093_v79)) : (_2098_v81)), p2, p2, globalState);
      _out46 = _outcollector12[0];
      _out47 = _outcollector12[1];
      _out48 = _outcollector12[2];
      _2105_v83 = _out46;
      _2106_v84 = _out47;
      _2107_v85 = _out48;
      let _2108_v86;
      _2108_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2025_v41,((_1960_v0).f3)[_module.__default.safeIndex(new BigNumber(-578), new BigNumber(((_1960_v0).f3).length))]);
      let _2109_v87;
      _2109_v87 = _dafny.Map.Empty.slice().updateUnsafe(_2108_v86,_dafny.Map.Empty.slice().updateUnsafe(_2038_v45,_2040_v47));
      let _2110_v88;
      _2110_v88 = _dafny.Map.Empty.slice().updateUnsafe(_2038_v45,_2040_v47);
      _2109_v87 = (_2109_v87).update((_2108_v86).Merge(_2108_v86), _2110_v88);
      let _2111_v89;
      _2111_v89 = _dafny.Set.fromElements(_2046_v53);
      r0 = new BigNumber((((_2111_v89).Union(_dafny.Set.fromElements(_2046_v53))).Union(_2111_v89)).length);
      let _2112_v90;
      _2112_v90 = _dafny.Seq.of(_2041_v48, _2041_v48);
      r1 = (new BigNumber((_dafny.Seq.update((_1960_v0).f3, _module.__default.safeIndex(new BigNumber((_2112_v90).length), new BigNumber(((_1960_v0).f3).length)), _2092_v78)).length)).multipliedBy((((_2038_v45)[_module.__default.safeIndex(_2040_v47, new BigNumber((_2038_v45).length))]) ? ((_dafny.ZERO).minus(_2107_v85)) : ((_this).fm1(p2, globalState))));
      return [r0, r1];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      if (p3) {
        let _2113_v0;
        let _nw380 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _2113_v0 = _nw380;
        let _index385 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2113_v0).length));
        (_2113_v0)[_index385] = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
        let _2114_v1;
        _2114_v1 = _dafny.Seq.of(p3, !(p3), !(p3), p0);
        let _index386 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_2113_v0).length));
        (_2113_v0)[_index386] = new BigNumber((_dafny.Seq.Concat(_2114_v1, _dafny.Seq.of(p3))).length);
        r3 = p3;
        let _2115_v2;
        _2115_v2 = _dafny.Seq.UnicodeFromString("iawbfjbim");
        let _2116_v3;
        _2116_v3 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2117_v4;
        _2117_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2116_v3,p0)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), ((_2118_v3) => function (_2119_i0) {
          return _2118_v3;
        })(_2116_v3)));
        _2115_v2 = _dafny.Seq.Concat(_2115_v2, (((_2117_v4).contains((_2113_v0)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_2113_v0).length))])) ? ((_2117_v4).get((_2113_v0)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_2113_v0).length))])) : (_2115_v2)));
        let _2120_v5;
        let _nw381 = Array((new BigNumber(25)).toNumber()).fill(false);
        _2120_v5 = _nw381;
        let _index387 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_2120_v5).length));
        (_2120_v5)[_index387] = !(!(true));
        let _2121_v6;
        _2121_v6 = _dafny.MultiSet.fromElements(new BigNumber(8));
        let _2122_v7;
        _2122_v7 = _dafny.MultiSet.fromElements(_2121_v6);
        let _2123_v8;
        _2123_v8 = _dafny.Seq.of(p2);
        let _index388 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_2120_v5).length));
        (_2120_v5)[_index388] = ((p0) ? (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ffeyfgi"), p1)) : ((_2122_v7).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_2123_v8)))));
        let _2124_v9;
        let _nw382 = Array((new BigNumber(3)).toNumber()).fill([]);
        _2124_v9 = _nw382;
        let _index389 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_2124_v9).length));
        (_2124_v9)[_index389] = _2113_v0;
        let _index390 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_2124_v9).length));
        (_2124_v9)[_index390] = _2113_v0;
      } else {
        let _2125_v10;
        _2125_v10 = _dafny.Set.fromElements(p2);
        let _2126_v11;
        _2126_v11 = _dafny.MultiSet.fromElements(p0);
        let _2127_v12;
        _2127_v12 = _module.D5.create_DC9(p0, p2, _2126_v11);
        let _2128_v13;
        let _nw383 = Array((new BigNumber(22)).toNumber());
        _nw383[(_dafny.ZERO).toNumber()] = p0;
        _nw383[(_dafny.ONE).toNumber()] = p3;
        _nw383[(new BigNumber(2)).toNumber()] = p3;
        _nw383[(new BigNumber(3)).toNumber()] = p3;
        _nw383[(new BigNumber(4)).toNumber()] = p0;
        _nw383[(new BigNumber(5)).toNumber()] = p3;
        _nw383[(new BigNumber(6)).toNumber()] = p3;
        _nw383[(new BigNumber(7)).toNumber()] = _module.__default.fm5(globalState);
        _nw383[(new BigNumber(8)).toNumber()] = p3;
        _nw383[(new BigNumber(9)).toNumber()] = p0;
        _nw383[(new BigNumber(10)).toNumber()] = p3;
        _nw383[(new BigNumber(11)).toNumber()] = p0;
        _nw383[(new BigNumber(12)).toNumber()] = p3;
        _nw383[(new BigNumber(13)).toNumber()] = p3;
        _nw383[(new BigNumber(14)).toNumber()] = p0;
        _nw383[(new BigNumber(15)).toNumber()] = p3;
        _nw383[(new BigNumber(16)).toNumber()] = p0;
        _nw383[(new BigNumber(17)).toNumber()] = p0;
        _nw383[(new BigNumber(18)).toNumber()] = !((_2127_v12).dtor_cf14);
        _nw383[(new BigNumber(19)).toNumber()] = p3;
        _nw383[(new BigNumber(20)).toNumber()] = p3;
        _nw383[(new BigNumber(21)).toNumber()] = _module.__default.fm5(globalState);
        _2128_v13 = _nw383;
        let _2129_v14;
        let _nw384 = new _module.C5();
        _nw384.__ctor(new BigNumber((_2125_v10).length), _2128_v13);
        _2129_v14 = _nw384;
        let _2130_v15;
        let _nw385 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _2130_v15 = _nw385;
        let _2131_v16;
        _2131_v16 = _dafny.Seq.UnicodeFromString("yghybw");
        let _rhs348 = ((p0) ? (_2130_v15) : (_2130_v15));
        let _rhs349 = p2;
        let _rhs350 = _2131_v16;
        _2130_v15 = _rhs348;
        r2 = _rhs349;
        _2131_v16 = _rhs350;
        let _index391 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length));
        (_2130_v15)[_index391] = p2;
        let _index392 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length));
        let _rhs351 = _2129_v14.f1;
        let _rhs352 = p3;
        let _rhs353 = p2;
        let _lhs177 = _2129_v14;
        let _lhs178 = _2130_v15;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length));
        _lhs177.f1 = _rhs351;
        r3 = _rhs352;
        _lhs178[_lhs179] = _rhs353;
        if (!(true) || (!(p0) || (p3))) {
          let _index393 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length));
          (_2130_v15)[_index393] = (_2129_v14).f0;
          let _index394 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length));
          (_2130_v15)[_index394] = (_2129_v14).f0;
          let _2132_v17;
          let _nw386 = new _module.C2();
          _nw386.__ctor(p2);
          _2132_v17 = _nw386;
          let _2133_v18;
          _2133_v18 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2131_v16);
          r2 = (new BigNumber((_dafny.Seq.Concat(_module.__default.fm15(globalState), (((_2133_v18).contains(!(!(p0)))) ? ((_2133_v18).get(!(!(p0)))) : (p1)))).length)).multipliedBy(p2);
          r3 = p0;
        } else {
          r3 = p3;
          let _2134_v19;
          let _nw387 = new _module.C3();
          _nw387.__ctor((_2130_v15)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_2130_v15).length))]);
          _2134_v19 = _nw387;
          let _2135_v20;
          _2135_v20 = new _dafny.CodePoint('s'.codePointAt(0));
          _2135_v20 = _2135_v20;
          let _2136_v21;
          _2136_v21 = _dafny.Seq.of(p0);
          let _2137_v22;
          _2137_v22 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
          r3 = (_2136_v21)[_module.__default.safeIndex(new BigNumber((((p0) ? (_2137_v22) : (_2137_v22))).length), new BigNumber((_2136_v21).length))];
          _2130_v15 = _2130_v15;
        }
        r3 = ((p3) ? (p0) : (p3));
      }
      if (p0) {
        let _2138_v23;
        let _nw388 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _2138_v23 = _nw388;
        let _index395 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
        (_2138_v23)[_index395] = (new BigNumber(890)).plus(p2);
        let _index396 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
        (_2138_v23)[_index396] = p2;
        let _index397 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
        (_2138_v23)[_index397] = p2;
        let _2139_v24;
        _2139_v24 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
        let _index398 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
        (_2138_v23)[_index398] = new BigNumber(((_2139_v24).Merge(_2139_v24)).length);
        let _2140_v25;
        let _nw389 = new _module.C2();
        _nw389.__ctor((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]);
        _2140_v25 = _nw389;
        let _2141_v26;
        _2141_v26 = _module.D6.create_DC11(p1);
        let _source31 = _2141_v26;
        if (_source31.is_DC12) {
          let _2142___mcc_h0 = (_source31).cf22;
          let _2143___mcc_h1 = (_source31).cf23;
          let _2144_cf23 = _2143___mcc_h1;
          let _2145_cf22 = _2142___mcc_h0;
          let _2146_v27;
          _2146_v27 = _dafny.Seq.of(_2145_cf22, _2145_cf22);
          let _2147_v28;
          let _nw390 = Array((new BigNumber(6)).toNumber());
          _nw390[(_dafny.ZERO).toNumber()] = _module.__default.fm5(globalState);
          _nw390[(_dafny.ONE).toNumber()] = _2145_cf22;
          _nw390[(new BigNumber(2)).toNumber()] = _2144_cf23;
          _nw390[(new BigNumber(3)).toNumber()] = p0;
          _nw390[(new BigNumber(4)).toNumber()] = (_2146_v27)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_2144_cf23)).length), new BigNumber((_2146_v27).length))];
          _nw390[(new BigNumber(5)).toNumber()] = _2144_cf23;
          _2147_v28 = _nw390;
          let _2148_v29;
          let _nw391 = new _module.C5();
          _nw391.__ctor(p2, ((p3) ? (_2147_v28) : (_2147_v28)));
          _2148_v29 = _nw391;
          let _arr15 = _2148_v29.f1;
          let _index399 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_2148_v29.f1).length));
          _arr15[_index399] = _2144_cf23;
          let _2149_v30;
          _2149_v30 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(151)).minus(new BigNumber(32)),p0);
          let _arr16 = _2148_v29.f1;
          let _index400 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_2148_v29.f1).length));
          _arr16[_index400] = (((_2149_v30).contains((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))])) ? ((_2149_v30).get((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))])) : (_2144_cf23));
          _2144_cf23 = (_2148_v29.f1)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_2148_v29.f1).length))];
          let _2150_v31;
          _2150_v31 = new _dafny.CodePoint('l'.codePointAt(0));
          _2150_v31 = _2150_v31;
        } else if (_source31.is_DC13) {
          let _2151___mcc_h2 = (_source31).cf24;
          let _2152___mcc_h3 = (_source31).cf25;
          let _2153_cf25 = _2152___mcc_h3;
          let _2154_cf24 = _2151___mcc_h2;
          let _2155_v32;
          _2155_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm16(globalState),p3);
          r1 = _module.__default.safeDivisionInt(new BigNumber((_2155_v32).length), _module.__default.safeModuloInt((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))], new BigNumber(908)));
          let _index401 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
          (_2138_v23)[_index401] = (_2140_v25).fm1(p0, globalState);
          let _2156_v33;
          let _nw392 = new _module.C1();
          _nw392.__ctor(p1);
          _2156_v33 = _nw392;
          _2156_v33 = _2156_v33;
          r3 = true;
        } else {
          let _2157___mcc_h4 = (_source31).cf21;
          let _2158_cf21 = _2157___mcc_h4;
          let _rhs354 = new BigNumber(853);
          let _rhs355 = ((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]).isLessThan(((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]).multipliedBy((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]));
          let _rhs356 = _2138_v23;
          let _rhs357 = (((p0) ? ((_this).fm1(p0, globalState)) : ((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber((_2158_cf21).length), p2));
          let _rhs358 = p0;
          r2 = _rhs354;
          r3 = _rhs355;
          _2138_v23 = _rhs356;
          r3 = _rhs357;
          r3 = _rhs358;
          let _index402 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length));
          (_2138_v23)[_index402] = ((p0) ? ((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]) : ((_2138_v23)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2138_v23).length))]));
          let _2159_v34;
          _2159_v34 = _dafny.MultiSet.fromElements(p2, p2);
          let _2160_v35;
          _2160_v35 = _2159_v34;
          r3 = (_2159_v34).IsProperSubsetOf(((_2160_v35)).Difference(_2159_v34));
          let _2161_v36;
          let _init49 = ((_2162_p0) => function (_2163_i1) {
            return _2162_p0;
          })(p0);
          let _nw393 = Array((new BigNumber(29)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw393.length); _i0_49++) {
            _nw393[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _2161_v36 = _nw393;
          let _index403 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2161_v36).length));
          (_2161_v36)[_index403] = p0;
          let _index404 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2161_v36).length));
          (_2161_v36)[_index404] = p3;
        }
      } else {
        if (p3) {
          r3 = !(p3);
          let _2164_v37;
          _2164_v37 = _dafny.Seq.of(p2);
          let _2165_v38;
          let _nw394 = Array((new BigNumber(7)).toNumber());
          _nw394[(_dafny.ZERO).toNumber()] = p2;
          _nw394[(_dafny.ONE).toNumber()] = p2;
          _nw394[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw394[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).fm1(p3, globalState), p2);
          _nw394[(new BigNumber(4)).toNumber()] = p2;
          _nw394[(new BigNumber(5)).toNumber()] = p2;
          _nw394[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2164_v37, _2164_v37)).length);
          _2165_v38 = _nw394;
          let _index405 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2165_v38).length));
          (_2165_v38)[_index405] = new BigNumber((_dafny.Seq.Concat(_2164_v37, _2164_v37)).length);
          let _index406 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_2165_v38).length));
          (_2165_v38)[_index406] = _module.__default.safeModuloInt(p2, new BigNumber((_module.__default.fm41(_2164_v37, p0, !(_module.__default.fm5(globalState)), p0, globalState)).length));
          let _2166_v39;
          _2166_v39 = _dafny.Map.Empty.slice().updateUnsafe((_2165_v38)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_2165_v38).length))],_dafny.Seq.of(p2));
          _2166_v39 = (_2166_v39).update((_dafny.ZERO).minus((_2165_v38)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_2165_v38).length))]), _2164_v37);
          let _2167_v40;
          let _nw395 = new _module.C1();
          _nw395.__ctor(p1);
          _2167_v40 = _nw395;
          let _2168_v41;
          let _nw396 = new _module.C4();
          _nw396.__ctor(((_module.__default.fm5(globalState)) ? (new BigNumber(795)) : ((_module.D8.create_DC17(p3, _2167_v40, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ycbfarff"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("ycbfarff")).length)), new _dafny.CodePoint('a'.codePointAt(0)))).length))).dtor_cf32)));
          _2168_v41 = _nw396;
          _2168_v41 = _2168_v41;
          let _2169_v42;
          let _nw397 = Array((_dafny.ONE).toNumber());
          _nw397[(_dafny.ZERO).toNumber()] = true;
          _2169_v42 = _nw397;
          let _2170_v43;
          _2170_v43 = _dafny.Set.fromElements(_2169_v42);
          let _2171_v44;
          let _nw398 = new _module.C0();
          _nw398.__ctor((_2170_v43).IsProperSubsetOf(_2170_v43), p0);
          _2171_v44 = _nw398;
        } else {
          let _2172_v45;
          _2172_v45 = _dafny.Map.Empty.slice().updateUnsafe(p2,(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(331),(_dafny.ZERO).minus(p2))).length)).minus(p2));
          let _2173_v46;
          _2173_v46 = _dafny.Seq.of(p0);
          let _2174_v47;
          let _nw399 = new _module.C3();
          _nw399.__ctor(p2);
          _2174_v47 = _nw399;
          let _2175_v48;
          _2175_v48 = _dafny.MultiSet.fromElements(p0);
          let _2176_v49;
          _2176_v49 = _module.D5.create_DC9(p0, new BigNumber((_dafny.Set.fromElements(_2174_v47, _2174_v47)).length), _2175_v48);
          let _2177_v50;
          _2177_v50 = _dafny.Seq.of((_2173_v46)[_module.__default.safeIndex(p2, new BigNumber((_2173_v46).length))], p0, (_2176_v49).dtor_cf14, p3);
          r1 = (_dafny.ZERO).minus((((_2172_v45).contains((p2).multipliedBy(p2))) ? ((_2172_v45).get((p2).multipliedBy(p2))) : (new BigNumber((_2173_v46).length))));
          let _2178_v51;
          _2178_v51 = _dafny.Set.fromElements(p2);
          let _2179_v52;
          _2179_v52 = _dafny.MultiSet.fromElements(p2, (_dafny.ZERO).minus(p2));
          let _2180_v53;
          let _nw400 = Array((new BigNumber(25)).toNumber());
          _nw400[(_dafny.ZERO).toNumber()] = p2;
          _nw400[(_dafny.ONE).toNumber()] = (((_2172_v45).contains(p2)) ? ((_2172_v45).get(p2)) : (new BigNumber((_2178_v51).length)));
          _nw400[(new BigNumber(2)).toNumber()] = p2;
          _nw400[(new BigNumber(3)).toNumber()] = p2;
          _nw400[(new BigNumber(4)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(5)).toNumber()] = (((_2179_v52).contains(new BigNumber((_module.__default.fm15(globalState)).length))) ? ((_2179_v52).get(new BigNumber((_module.__default.fm15(globalState)).length))) : (p2));
          _nw400[(new BigNumber(6)).toNumber()] = new BigNumber((_2175_v48).cardinality());
          _nw400[(new BigNumber(7)).toNumber()] = p2;
          _nw400[(new BigNumber(8)).toNumber()] = p2;
          _nw400[(new BigNumber(9)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw400[(new BigNumber(11)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(12)).toNumber()] = (_this).fm1(p0, globalState);
          _nw400[(new BigNumber(13)).toNumber()] = new BigNumber((p1).length);
          _nw400[(new BigNumber(14)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(15)).toNumber()] = p2;
          _nw400[(new BigNumber(16)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(17)).toNumber()] = new BigNumber(267);
          _nw400[(new BigNumber(18)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(19)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(20)).toNumber()] = p2;
          _nw400[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw400[(new BigNumber(22)).toNumber()] = _2174_v47.f2;
          _nw400[(new BigNumber(23)).toNumber()] = new BigNumber((p1).length);
          _nw400[(new BigNumber(24)).toNumber()] = (_2174_v47).fm1(p0, globalState);
          _2180_v53 = _nw400;
          r1 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p0),_2180_v53)).length);
          let _2181_v54;
          _2181_v54 = _dafny.Seq.of(_module.__default.fm16(globalState));
          let _2182_v55;
          let _nw401 = Array((new BigNumber(8)).toNumber());
          _nw401[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p0), _dafny.Seq.update(_2173_v46, _module.__default.safeIndex(_2174_v47.f2, new BigNumber((_2173_v46).length)), p0));
          _nw401[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2173_v46, _2177_v50);
          _nw401[(new BigNumber(2)).toNumber()] = (_2181_v54)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_2181_v54).length))];
          _nw401[(new BigNumber(3)).toNumber()] = _2177_v50;
          _nw401[(new BigNumber(4)).toNumber()] = _2173_v46;
          _nw401[(new BigNumber(5)).toNumber()] = _2177_v50;
          _nw401[(new BigNumber(6)).toNumber()] = _2173_v46;
          _nw401[(new BigNumber(7)).toNumber()] = _2177_v50;
          _2182_v55 = _nw401;
          let _index407 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2182_v55).length));
          (_2182_v55)[_index407] = _module.__default.fm16(globalState);
          let _index408 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2182_v55).length));
          (_2182_v55)[_index408] = _2173_v46;
          r3 = p0;
          let _2183_v56;
          _2183_v56 = _dafny.Seq.UnicodeFromString("ndgceyxxq");
          let _index409 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2180_v53).length));
          (_2180_v53)[_index409] = (_this).fm1(p0, globalState);
          let _index410 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2180_v53).length));
          let _rhs359 = _dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("sxrkeyvv"));
          let _rhs360 = p3;
          let _rhs361 = _2174_v47.f2;
          let _rhs362 = _2180_v53;
          let _rhs363 = (_this).fm1((_2174_v47).fm8(_dafny.Seq.UnicodeFromString("wtikft"), p3, globalState), globalState);
          let _lhs180 = _2180_v53;
          let _lhs181 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_2180_v53).length));
          let _lhs182 = _2174_v47;
          _2183_v56 = _rhs359;
          r3 = _rhs360;
          _lhs180[_lhs181] = _rhs361;
          _2180_v53 = _rhs362;
          _lhs182.f2 = _rhs363;
        }
        let _2184_v57;
        let _nw402 = Array((new BigNumber(9)).toNumber()).fill(_module.D4.Default());
        _2184_v57 = _nw402;
        let _2185_v58;
        _2185_v58 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
        let _2186_v59;
        _2186_v59 = _dafny.Set.fromElements(new BigNumber((_2185_v58).length), p2);
        let _2187_v60;
        _2187_v60 = _dafny.Seq.of(new BigNumber(410));
        let _2188_v61;
        let _nw403 = Array((new BigNumber(20)).toNumber());
        _nw403[(_dafny.ZERO).toNumber()] = p2;
        _nw403[(_dafny.ONE).toNumber()] = p2;
        _nw403[(new BigNumber(2)).toNumber()] = p2;
        _nw403[(new BigNumber(3)).toNumber()] = new BigNumber((_2186_v59).length);
        _nw403[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw403[(new BigNumber(5)).toNumber()] = p2;
        _nw403[(new BigNumber(6)).toNumber()] = p2;
        _nw403[(new BigNumber(7)).toNumber()] = p2;
        _nw403[(new BigNumber(8)).toNumber()] = p2;
        _nw403[(new BigNumber(9)).toNumber()] = new BigNumber((_2187_v60).length);
        _nw403[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((p1).length));
        _nw403[(new BigNumber(11)).toNumber()] = p2;
        _nw403[(new BigNumber(12)).toNumber()] = new BigNumber(934);
        _nw403[(new BigNumber(13)).toNumber()] = p2;
        _nw403[(new BigNumber(14)).toNumber()] = new BigNumber((_2187_v60).length);
        _nw403[(new BigNumber(15)).toNumber()] = p2;
        _nw403[(new BigNumber(16)).toNumber()] = p2;
        _nw403[(new BigNumber(17)).toNumber()] = new BigNumber(-472);
        _nw403[(new BigNumber(18)).toNumber()] = p2;
        _nw403[(new BigNumber(19)).toNumber()] = new BigNumber(-928);
        _2188_v61 = _nw403;
        let _2189_v62;
        _2189_v62 = _module.D4.create_DC6(p2, _2188_v61, new BigNumber(994));
        let _index411 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2184_v57).length));
        (_2184_v57)[_index411] = _2189_v62;
        let _2190_v63;
        let _nw404 = Array((new BigNumber(26)).toNumber()).fill([]);
        _2190_v63 = _nw404;
        let _2191_v64;
        let _nw405 = Array((new BigNumber(4)).toNumber()).fill(false);
        _2191_v64 = _nw405;
        let _index412 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_2190_v63).length));
        (_2190_v63)[_index412] = _2191_v64;
        let _2192_v65;
        _2192_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,new _dafny.CodePoint('n'.codePointAt(0)));
        let _2193_v66;
        _2193_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2192_v65,((true) ? (p0) : (p3)));
        let _2194_v68;
        _2194_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("fao"));
        let _2195_v69;
        _2195_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_2194_v68).contains(p0)) ? ((_2194_v68).get(p0)) : (p1)));
        let _index413 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2184_v57).length));
        let _index414 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_2190_v63).length));
        let _rhs364 = _module.D4.create_DC6(new BigNumber((_2187_v60).length), _2188_v61, p2);
        let _rhs365 = _2191_v64;
        let _rhs366 = (((_2193_v66).contains((_2192_v65).Merge(function () {
          let _coll67 = new _dafny.Map();
          for (const _compr_67 of (_2195_v69).Keys.Elements) {
            let _2197_v67 = _compr_67;
            if ((_2195_v69).contains(_2197_v67)) {
              _coll67.push([_2197_v67,new _dafny.CodePoint('o'.codePointAt(0))]);
            }
          }
          return _coll67;
        }()))) ? ((_2193_v66).get((_2192_v65).Merge(function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_2195_v69).Keys.Elements) {
            let _2196_v67 = _compr_66;
            if ((_2195_v69).contains(_2196_v67)) {
              _coll66.push([_2196_v67,new _dafny.CodePoint('o'.codePointAt(0))]);
            }
          }
          return _coll66;
        }()))) : (p3));
        let _rhs367 = new BigNumber(206);
        let _lhs183 = _2184_v57;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2184_v57).length));
        let _lhs185 = _2190_v63;
        let _lhs186 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_2190_v63).length));
        _lhs183[_lhs184] = _rhs364;
        _lhs185[_lhs186] = _rhs365;
        r3 = _rhs366;
        r2 = _rhs367;
        let _2198_v70;
        _2198_v70 = new _dafny.CodePoint('l'.codePointAt(0));
        _2198_v70 = _2198_v70;
        let _2199_v71;
        _2199_v71 = _dafny.MultiSet.fromElements(new BigNumber((p1).length));
        let _2200_v72;
        _2200_v72 = _dafny.Seq.of(_2199_v71);
        let _2201_v73;
        _2201_v73 = _dafny.Map.Empty.slice().updateUnsafe((_2190_v63)[_module.__default.safeIndex(new BigNumber(439), new BigNumber((_2190_v63).length))],!((_2199_v71).IsSubsetOf((_2200_v72)[_module.__default.safeIndex(p2, new BigNumber((_2200_v72).length))])));
        _2201_v73 = _2201_v73;
        r2 = (p2).multipliedBy(new BigNumber(119));
      }
      let _2202_v74;
      _2202_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(66),p3);
      let _2203_v75;
      _2203_v75 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2202_v74);
      let _2204_v76;
      _2204_v76 = _module.D5.create_DC10(false, p2, (_2203_v75).update(p1, _2202_v74), p0);
      let _hi13 = (_2204_v76).dtor_cf18;
      for (let _2205_i2 = p2; _2205_i2.isLessThan(_hi13); _2205_i2 = _2205_i2.plus(_dafny.ONE)) {
        r3 = p0;
        let _2206_v77;
        let _nw406 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2206_v77 = _nw406;
        _2206_v77 = _2206_v77;
        if (!(p3)) {
          let _2207_v78;
          _2207_v78 = _dafny.Set.fromElements(_2205_i2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_2208_i3) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length), _2205_i2);
          let _rhs368 = _module.__default.safeModuloInt((_dafny.ZERO).minus(p2), new BigNumber((_2207_v78).length));
          let _rhs369 = (((p0) ? (_2205_i2) : (p2))).plus(new BigNumber(-664));
          r1 = _rhs368;
          r2 = _rhs369;
          r2 = _2205_i2;
          r3 = false;
          r2 = p2;
          let _2209_v79;
          _2209_v79 = _dafny.MultiSet.fromElements(p0, p3);
          let _2210_v80;
          _2210_v80 = _dafny.Map.Empty.slice().updateUnsafe((_2205_i2).plus((((_2209_v79).contains(p0)) ? ((_2209_v79).get(p0)) : (_2205_i2))),p2);
          let _2211_v81;
          _2211_v81 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2209_v79);
          _2210_v80 = (_2210_v80).update(((p3) ? (p2) : (_2205_i2)), new BigNumber((((_module.__default.fm5(globalState)) ? (_2209_v79) : ((((_2211_v81).contains(p3)) ? ((_2211_v81).get(p3)) : (_2209_v79))))).cardinality()));
        } else {
          let _2212_v82;
          _2212_v82 = _dafny.Seq.of(new BigNumber((p1).length), _2205_i2);
          let _2213_v83;
          _2213_v83 = _dafny.Map.Empty.slice().updateUnsafe(_2212_v82,p0);
          let _2214_v84;
          _2214_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("c"),!_dafny.areEqual(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-663)), new BigNumber((_2213_v83).length), (_this).fm1(true, globalState)), _2212_v82));
          _2214_v84 = (_2214_v84).update(p1, p3);
          let _index415 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length));
          (_2206_v77)[_index415] = !(p0);
          let _index416 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length));
          (_2206_v77)[_index416] = !(p3) || (p3);
          let _2215_v85;
          let _nw407 = new _module.C1();
          _nw407.__ctor(_dafny.Seq.Concat(p1, p1));
          _2215_v85 = _nw407;
          let _2216_v86;
          let _nw408 = Array((new BigNumber(26)).toNumber());
          _nw408[(_dafny.ZERO).toNumber()] = false;
          _nw408[(_dafny.ONE).toNumber()] = p3;
          _nw408[(new BigNumber(2)).toNumber()] = p0;
          _nw408[(new BigNumber(3)).toNumber()] = p0;
          _nw408[(new BigNumber(4)).toNumber()] = p3;
          _nw408[(new BigNumber(5)).toNumber()] = !(_module.__default.fm5(globalState));
          _nw408[(new BigNumber(6)).toNumber()] = p3;
          _nw408[(new BigNumber(7)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(8)).toNumber()] = p3;
          _nw408[(new BigNumber(9)).toNumber()] = p0;
          _nw408[(new BigNumber(10)).toNumber()] = p3;
          _nw408[(new BigNumber(11)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(12)).toNumber()] = p0;
          _nw408[(new BigNumber(13)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(14)).toNumber()] = true;
          _nw408[(new BigNumber(15)).toNumber()] = p3;
          _nw408[(new BigNumber(16)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(17)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(18)).toNumber()] = p0;
          _nw408[(new BigNumber(19)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(20)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(21)).toNumber()] = p3;
          _nw408[(new BigNumber(22)).toNumber()] = p3;
          _nw408[(new BigNumber(23)).toNumber()] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
          _nw408[(new BigNumber(24)).toNumber()] = p0;
          _nw408[(new BigNumber(25)).toNumber()] = p3;
          _2216_v86 = _nw408;
          let _2217_v87;
          _2217_v87 = _dafny.Map.Empty.slice().updateUnsafe(_2216_v86,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pgdsyug"), (_2215_v85).f3));
          let _2218_v88;
          _2218_v88 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index417 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length));
          let _rhs370 = new BigNumber((_dafny.Seq.of(((false) ? (_2218_v88) : (_2218_v88)))).length);
          let _rhs371 = _2205_i2;
          let _rhs372 = p2;
          let _rhs373 = ((!((_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))])) ? ((_2217_v87).Merge(_2217_v87)) : (_2217_v87));
          let _rhs374 = p3;
          let _lhs187 = _2206_v77;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length));
          r2 = _rhs370;
          r2 = _rhs371;
          r2 = _rhs372;
          _2217_v87 = _rhs373;
          _lhs187[_lhs188] = _rhs374;
          let _index418 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length));
          (_2206_v77)[_index418] = (_2206_v77)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_2206_v77).length))];
        }
        r1 = (_this).fm1(p3, globalState);
      }
      let _2219_v89;
      _2219_v89 = _module.D15.create_DC30(true);
      r1 = (_this).fm1((_2219_v89).dtor_cf47, globalState);
      let _2220_v90;
      let _nw409 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _2220_v90 = _nw409;
      let _index419 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length));
      (_2220_v90)[_index419] = _module.__default.safeModuloInt(new BigNumber(138), new BigNumber(795));
      let _2221_v91;
      _2221_v91 = _dafny.Seq.of(new BigNumber((p1).length));
      let _2222_v92;
      _2222_v92 = _dafny.Map.Empty.slice().updateUnsafe(_2221_v91,p3);
      let _index420 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length));
      (_2220_v90)[_index420] = (new BigNumber((_2222_v92).length)).minus(p2);
      let _2223_v93;
      let _nw410 = Array((new BigNumber(13)).toNumber()).fill([]);
      _2223_v93 = _nw410;
      let _2224_v95;
      _2224_v95 = _dafny.Seq.of(p1, _dafny.Seq.UnicodeFromString("sdflwo"));
      let _2225_v96;
      _2225_v96 = _module.D15.create_DC29(function () {
  let _coll68 = new _dafny.Map();
  for (const _compr_68 of (_dafny.Seq.update(_2224_v95, _module.__default.safeIndex(p2, new BigNumber((_2224_v95).length)), (_2224_v95)[_module.__default.safeIndex((_2220_v90)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length))], new BigNumber((_2224_v95).length))])).Elements) {
    let _2226_v94 = _compr_68;
    if (_dafny.Seq.contains(_dafny.Seq.update(_2224_v95, _module.__default.safeIndex(p2, new BigNumber((_2224_v95).length)), (_2224_v95)[_module.__default.safeIndex((_2220_v90)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length))], new BigNumber((_2224_v95).length))]), _2226_v94)) {
      _coll68.push([_2226_v94,p1]);
    }
  }
  return _coll68;
}());
      let _pat_let_tv47 = _2220_v90;
      let _pat_let_tv48 = _2220_v90;
      let _pat_let_tv49 = _2220_v90;
      let _pat_let_tv50 = _2220_v90;
      let _pat_let_tv51 = _2220_v90;
      let _pat_let_tv52 = _2220_v90;
      let _index421 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length));
      let _rhs375 = _2223_v93;
      let _rhs376 = _module.__default.safeDivisionInt((_2220_v90)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length))], p2);
      let _rhs377 = function (_source32) {
        if (_source32.is_DC30) {
          let _2227___mcc_h5 = (_source32).cf47;
          let _2228_cf47 = _2227___mcc_h5;
          return _module.__default.safeDivisionInt((_pat_let_tv48)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_pat_let_tv47).length))], (_pat_let_tv50)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_pat_let_tv49).length))]);
        } else {
          let _2229___mcc_h6 = (_source32).cf46;
          let _2230_cf46 = _2229___mcc_h6;
          return (_pat_let_tv52)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_pat_let_tv51).length))];
        }
      }(_2225_v96);
      let _rhs378 = new BigNumber((p1).length);
      let _lhs189 = _2220_v90;
      let _lhs190 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length));
      _2223_v93 = _rhs375;
      r1 = _rhs376;
      _lhs189[_lhs190] = _rhs377;
      r1 = _rhs378;
      r0 = _dafny.Seq.Concat(((!(p3)) ? (_2224_v95) : (_dafny.Seq.of(p1, p1, p1, p1, p1))), _dafny.Seq.Concat(_2224_v95, _2224_v95));
      let _2231_v97;
      _2231_v97 = _dafny.MultiSet.fromElements(p0);
      r1 = (((_2231_v97).IsProperSubsetOf(_2231_v97)) ? (new BigNumber((p1).length)) : (p2));
      r2 = (_2221_v91)[_module.__default.safeIndex((_2220_v90)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length))], new BigNumber((_2221_v91).length))];
      let _2232_v98;
      _2232_v98 = new _dafny.CodePoint('b'.codePointAt(0));
      let _2233_v99;
      _2233_v99 = _module.D7.create_DC15((_2220_v90)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_2220_v90).length))], _2232_v98);
      let _pat_let_tv53 = p3;
      let _pat_let_tv54 = _2220_v90;
      let _pat_let_tv55 = _2220_v90;
      r3 = function (_source33) {
        if (_source33.is_DC15) {
          let _2234___mcc_h7 = (_source33).cf27;
          let _2235___mcc_h8 = (_source33).cf28;
          let _2236_cf28 = _2235___mcc_h8;
          let _2237_cf27 = _2234___mcc_h7;
          return _pat_let_tv53;
        } else {
          let _2238___mcc_h9 = (_source33).cf26;
          let _2239_cf26 = _2238___mcc_h9;
          return false;
        }
      }(function (_pat_let43_0) {
        return function (_2240_dt__update__tmp_h0) {
          return function (_pat_let44_0) {
            return function (_2241_dt__update_hcf27_h0) {
              return _module.D7.create_DC15(_2241_dt__update_hcf27_h0, (_2240_dt__update__tmp_h0).dtor_cf28);
            }(_pat_let44_0);
          }((_pat_let_tv55)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_pat_let_tv54).length))]);
        }(_pat_let43_0);
      }(_2233_v99));
      return [r0, r1, r2, r3];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_2242_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality());
      }),new BigNumber(711))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), function (_2243_i1) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("cpdthucj")).length);
      }),new BigNumber(618))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(280), new BigNumber((function () {
        let _coll69 = new _dafny.Map();
        for (const _compr_69 of _dafny.IntegerRange(new BigNumber(-18), new BigNumber(784))) {
          let _2244_v0 = _compr_69;
          if (((new BigNumber(-18)).isLessThanOrEqualTo(_2244_v0)) && ((_2244_v0).isLessThan(new BigNumber(784)))) {
            _coll69.push([_module.__default.safeDivisionInt(_2244_v0, new BigNumber((_dafny.Seq.of(false)).length)),false]);
          }
        }
        return _coll69;
      }()).length), new BigNumber(438), new BigNumber(-11), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), function (_2245_i2) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-503)), function (_2246_i3) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))));
    };
    fm1(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber(((_dafny.Set.fromElements(!(!(false)), false, false)).Intersect(_dafny.Set.fromElements(true, (_module.D0.create_DC0(true)).dtor_cf0))).length), new BigNumber(312));
    };
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _2247_v0;
      _2247_v0 = _dafny.Seq.UnicodeFromString("hsn");
      _2247_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2247_v0, _module.__default.fm2(globalState)), _2247_v0);
      let _2248_v1;
      _2248_v1 = new BigNumber(583);
      r1 = (_dafny.ZERO).minus(_2248_v1);
      let _2249_v2;
      _2249_v2 = true;
      _2249_v2 = _2249_v2;
      _2249_v2 = _2249_v2;
      r0 = _2248_v1;
      let _2250_v3;
      _2250_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v2,p3);
      let _2251_v4;
      _2251_v4 = _dafny.Seq.of(_2249_v2);
      let _2252_v5;
      _2252_v5 = _dafny.Seq.of(p2, (((_2250_v3).contains(true)) ? ((_2250_v3).get(true)) : ((_2251_v4)[_module.__default.safeIndex((_this).fm1(p3, globalState), new BigNumber((_2251_v4).length))])), p2);
      let _2253_v6;
      let _nw411 = Array((new BigNumber(22)).toNumber());
      _nw411[(_dafny.ZERO).toNumber()] = (p2) === (_2249_v2);
      _nw411[(_dafny.ONE).toNumber()] = p3;
      _nw411[(new BigNumber(2)).toNumber()] = _2249_v2;
      _nw411[(new BigNumber(3)).toNumber()] = true;
      _nw411[(new BigNumber(4)).toNumber()] = _2249_v2;
      _nw411[(new BigNumber(5)).toNumber()] = (_2251_v4)[_module.__default.safeIndex(_2248_v1, new BigNumber((_2251_v4).length))];
      _nw411[(new BigNumber(6)).toNumber()] = _2249_v2;
      _nw411[(new BigNumber(7)).toNumber()] = !(_2249_v2);
      _nw411[(new BigNumber(8)).toNumber()] = p2;
      _nw411[(new BigNumber(9)).toNumber()] = !(_2249_v2);
      _nw411[(new BigNumber(10)).toNumber()] = true;
      _nw411[(new BigNumber(11)).toNumber()] = !(p3);
      _nw411[(new BigNumber(12)).toNumber()] = (_2249_v2) && (p2);
      _nw411[(new BigNumber(13)).toNumber()] = _2249_v2;
      _nw411[(new BigNumber(14)).toNumber()] = p2;
      _nw411[(new BigNumber(15)).toNumber()] = _module.__default.fm3(true, _2248_v1, _2248_v1, globalState);
      _nw411[(new BigNumber(16)).toNumber()] = !_dafny.areEqual(_2251_v4, _2251_v4);
      _nw411[(new BigNumber(17)).toNumber()] = (_2248_v1).isLessThan(_2248_v1);
      _nw411[(new BigNumber(18)).toNumber()] = !((((((_2250_v3).contains(p3)) ? ((_2250_v3).get(p3)) : ((((_2250_v3).contains(p2)) ? ((_2250_v3).get(p2)) : (p2))))) ? (p2) : (false)));
      _nw411[(new BigNumber(19)).toNumber()] = _2249_v2;
      _nw411[(new BigNumber(20)).toNumber()] = p2;
      _nw411[(new BigNumber(21)).toNumber()] = _2249_v2;
      _2253_v6 = _nw411;
      let _2254_v7;
      _2254_v7 = _dafny.Seq.of(_2248_v1, new BigNumber((p0).cardinality()), _2248_v1);
      let _2255_v8;
      _2255_v8 = _dafny.Set.fromElements(_2248_v1, new BigNumber(799));
      let _2256_v9;
      _2256_v9 = _dafny.Set.fromElements(_2255_v8);
      let _2257_v10;
      _2257_v10 = _dafny.Seq.of(_2255_v8, _2255_v8);
      let _2258_v12;
      _2258_v12 = _dafny.Seq.of(_2256_v9, _dafny.Set.fromElements(_2255_v8), function () {
        let _coll70 = new _dafny.Set();
        for (const _compr_70 of (_2257_v10).Elements) {
          let _2259_v11 = _compr_70;
          if (_dafny.Seq.contains(_2257_v10, _2259_v11)) {
            _coll70.add(_2259_v11);
          }
        }
        return _coll70;
      }(), _2256_v9, _2256_v9);
      let _2260_v13;
      _2260_v13 = new _dafny.CodePoint('t'.codePointAt(0));
      let _2261_v14;
      _2261_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2260_v13,p3);
      let _2262_v15;
      _2262_v15 = _module.D0.create_DC1(_2249_v2, p3, _2260_v13, (_dafny.ZERO).minus(_2248_v1));
      let _2263_v16;
      _2263_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2262_v15,_2249_v2);
      let _nw412 = Array((new BigNumber(24)).toNumber());
      _nw412[(_dafny.ZERO).toNumber()] = (p2) || (p2);
      _nw412[(_dafny.ONE).toNumber()] = !(_2249_v2) || (_2249_v2);
      _nw412[(new BigNumber(2)).toNumber()] = p2;
      _nw412[(new BigNumber(3)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(4)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(5)).toNumber()] = p3;
      _nw412[(new BigNumber(6)).toNumber()] = (((_2250_v3).contains(true)) ? ((_2250_v3).get(true)) : (p2));
      _nw412[(new BigNumber(7)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(8)).toNumber()] = true;
      _nw412[(new BigNumber(9)).toNumber()] = p3;
      _nw412[(new BigNumber(10)).toNumber()] = false;
      _nw412[(new BigNumber(11)).toNumber()] = p2;
      _nw412[(new BigNumber(12)).toNumber()] = true;
      _nw412[(new BigNumber(13)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(14)).toNumber()] = false;
      _nw412[(new BigNumber(15)).toNumber()] = !_dafny.areEqual(_2254_v7, _2254_v7);
      _nw412[(new BigNumber(16)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(17)).toNumber()] = _2249_v2;
      _nw412[(new BigNumber(18)).toNumber()] = ((_2258_v12)[_module.__default.safeIndex(new BigNumber((_2261_v14).length), new BigNumber((_2258_v12).length))]).IsProperSubsetOf(_dafny.Set.fromElements(_2255_v8));
      _nw412[(new BigNumber(19)).toNumber()] = !(_2263_v16).contains(_module.D0.create_DC1(_2249_v2, _2249_v2, _2260_v13, _2248_v1));
      _nw412[(new BigNumber(20)).toNumber()] = (_2255_v8).IsSubsetOf(_2255_v8);
      _nw412[(new BigNumber(21)).toNumber()] = p2;
      _nw412[(new BigNumber(22)).toNumber()] = p3;
      _nw412[(new BigNumber(23)).toNumber()] = p3;
      _2253_v6 = _nw412;
      r0 = (_2248_v1).multipliedBy(_2248_v1);
      let _2264_v17;
      _2264_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2253_v6,new _dafny.CodePoint('r'.codePointAt(0)));
      let _2265_v18;
      _2265_v18 = _dafny.MultiSet.fromElements(_2260_v13, _2260_v13, (((_2264_v17).contains(_2253_v6)) ? ((_2264_v17).get(_2253_v6)) : (_2260_v13)));
      r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_2265_v18).cardinality())), _module.__default.safeModuloInt((_this).fm1(_2249_v2, globalState), _2248_v1));
      return [r0, r1];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _2266_v1;
      _2266_v1 = _dafny.MultiSet.fromElements(p0);
      if ((function () {
        let _coll71 = new _dafny.Map();
        for (const _compr_71 of (_2266_v1).Elements) {
          let _2267_v0 = _compr_71;
          if ((_2266_v1).contains(_2267_v0)) {
            _coll71.push([(_2267_v0).multipliedBy(new BigNumber(((_2266_v1).update(p0, _module.__default.abs(p0))).cardinality())),p0]);
          }
        }
        return _coll71;
      }()).contains(p0)) {
        let _2268_v2;
        _2268_v2 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2269_v3;
        _2269_v3 = _dafny.Set.fromElements(p3, p2, p2);
        let _2270_v4;
        _2270_v4 = _module.D0.create_DC1(p2, p3, _2268_v2, new BigNumber((_2269_v3).length));
        let _2271_v5;
        _2271_v5 = _dafny.MultiSet.fromElements(_2270_v4, _2270_v4);
        let _2272_v6;
        _2272_v6 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_2271_v5).update(_2270_v4, _module.__default.abs(p0)));
        let _2273_v7;
        _2273_v7 = _dafny.Seq.of(p0, p0);
        let _2274_v8;
        _2274_v8 = _dafny.Seq.of(_module.__default.fm4(p0, new BigNumber((_2273_v7).length), p2, p2, globalState));
        _2272_v6 = (_2272_v6).update(p2, _dafny.MultiSet.FromArray(_2274_v8));
        let _2275_v9;
        let _out49;
        _out49 = _module.__default.m0((!(p3)) === (p2), p3, globalState);
        _2275_v9 = _out49;
        r0 = p3;
        let _2276_v10;
        let _nw413 = new _module.C0();
        _nw413.__ctor(p3, p3);
        _2276_v10 = _nw413;
        let _2277_v11;
        let _nw414 = Array((new BigNumber(2)).toNumber()).fill(false);
        _2277_v11 = _nw414;
        let _index422 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2277_v11).length));
        (_2277_v11)[_index422] = p2;
        let _index423 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2277_v11).length));
        (_2277_v11)[_index423] = p2;
      } else {
        let _2278_v12;
        _2278_v12 = _dafny.Seq.UnicodeFromString("acq");
        let _2279_v13;
        _2279_v13 = _dafny.Set.fromElements(p3);
        let _2280_v14;
        _2280_v14 = _dafny.Map.Empty.slice().updateUnsafe(((p3) ? (_2278_v12) : (_2278_v12)),(_dafny.Set.fromElements(p2)).Intersect(_2279_v13));
        let _2281_v15;
        _2281_v15 = new _dafny.CodePoint('c'.codePointAt(0));
        _2280_v14 = (_2280_v14).update(_dafny.Seq.update(_2278_v12, _module.__default.safeIndex((_this).fm1(p2, globalState), new BigNumber((_2278_v12).length)), _2281_v15), _2279_v13);
        let _2282_v16;
        _2282_v16 = _dafny.Seq.of(p0);
        let _2283_v17;
        _2283_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2266_v1,p3);
        let _2284_v18;
        _2284_v18 = _dafny.Set.fromElements((_this).fm1(p3, globalState), p0, (_dafny.ZERO).minus(new BigNumber((_2282_v16).length)), new BigNumber((_2283_v17).length));
        r2 = _module.__default.safeModuloInt(p0, new BigNumber(((_2284_v18).Difference(_2284_v18)).length));
        let _2285_v19;
        _2285_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _2286_v20;
        _2286_v20 = _module.D6.create_DC13(p2, p0);
        let _2287_v21;
        let _nw415 = Array((new BigNumber(20)).toNumber());
        _nw415[(_dafny.ZERO).toNumber()] = p3;
        _nw415[(_dafny.ONE).toNumber()] = p3;
        _nw415[(new BigNumber(2)).toNumber()] = (((_2285_v19).contains(!(p2))) ? ((_2285_v19).get(!(p2))) : (p3));
        _nw415[(new BigNumber(3)).toNumber()] = p2;
        _nw415[(new BigNumber(4)).toNumber()] = _module.__default.fm5(globalState);
        _nw415[(new BigNumber(5)).toNumber()] = (((_2285_v19).contains(true)) ? ((_2285_v19).get(true)) : (p3));
        _nw415[(new BigNumber(6)).toNumber()] = p2;
        _nw415[(new BigNumber(7)).toNumber()] = p2;
        _nw415[(new BigNumber(8)).toNumber()] = !(p0).isEqualTo(p0);
        _nw415[(new BigNumber(9)).toNumber()] = p2;
        _nw415[(new BigNumber(10)).toNumber()] = !(p2);
        _nw415[(new BigNumber(11)).toNumber()] = p3;
        _nw415[(new BigNumber(12)).toNumber()] = false;
        _nw415[(new BigNumber(13)).toNumber()] = p3;
        _nw415[(new BigNumber(14)).toNumber()] = ((p2) ? (p3) : (!(p2)));
        _nw415[(new BigNumber(15)).toNumber()] = p3;
        _nw415[(new BigNumber(16)).toNumber()] = !(p2);
        _nw415[(new BigNumber(17)).toNumber()] = _module.__default.fm3(p3, p0, p0, globalState);
        _nw415[(new BigNumber(18)).toNumber()] = p3;
        _nw415[(new BigNumber(19)).toNumber()] = (_2286_v20).dtor_cf24;
        _2287_v21 = _nw415;
        let _2288_v22;
        _2288_v22 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.Concat(_2278_v12, _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), _2281_v15, _2281_v15)));
        let _rhs379 = _2287_v21;
        let _rhs380 = (_dafny.ZERO).minus(new BigNumber((_2288_v22).length));
        _2287_v21 = _rhs379;
        r2 = _rhs380;
        let _2289_v23;
        let _nw416 = new _module.C1();
        _nw416.__ctor(_2278_v12);
        _2289_v23 = _nw416;
        let _2290_v24;
        _2290_v24 = _dafny.MultiSet.fromElements(_2289_v23, _2289_v23);
        let _2291_v25;
        _2291_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2290_v24, _2290_v24, _2290_v24),_dafny.Seq.Concat(_2282_v16, _2282_v16));
        _2291_v25 = (_2291_v25).update(_dafny.Set.fromElements((_2290_v24).update(_2289_v23, _module.__default.abs(p0))), _2282_v16);
        if (p3) {
          r2 = p0;
          let _2292_v26;
          let _nw417 = new _module.C4();
          _nw417.__ctor(p0);
          _2292_v26 = _nw417;
          let _2293_v27;
          _2293_v27 = _dafny.Set.fromElements(_2292_v26);
          let _2294_v28;
          _2294_v28 = _2293_v27;
          let _2295_v29;
          _2295_v29 = _dafny.Seq.of(_2293_v27, (_dafny.Set.fromElements(_2292_v26)).Union(_2293_v27), (_2294_v28));
          _2295_v29 = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(_2292_v26, _2292_v26, _2292_v26, _2292_v26), _2293_v27), _2295_v29);
          r0 = (((_2285_v19).contains(_module.__default.fm5(globalState))) ? ((_2285_v19).get(_module.__default.fm5(globalState))) : (true));
          let _2296_v30;
          _2296_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2287_v21,_2278_v12);
          _2296_v30 = _2296_v30;
          let _2297_v31;
          _2297_v31 = _module.D14.create_DC28();
          let _2298_v32;
          _2298_v32 = _dafny.Seq.of(_2297_v31, _2297_v31);
          let _index424 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2287_v21).length));
          (_2287_v21)[_index424] = p3;
          let _2299_v33;
          let _nw418 = Array((new BigNumber(28)).toNumber()).fill([]);
          _2299_v33 = _nw418;
          let _index425 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2299_v33).length));
          (_2299_v33)[_index425] = _2287_v21;
          let _2300_v34;
          _2300_v34 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(321)), ((_2301_v31) => function (_2302_i0) {
            return _2301_v31;
          })(_2297_v31)), _2298_v32));
          let _2303_v35;
          _2303_v35 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
          let _2304_v36;
          _2304_v36 = _dafny.Map.Empty.slice().updateUnsafe((_2303_v35).contains(p3),_2287_v21);
          let _index426 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2287_v21).length));
          let _index427 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2299_v33).length));
          let _rhs381 = (_2300_v34)[_module.__default.safeIndex(p0, new BigNumber((_2300_v34).length))];
          let _rhs382 = p3;
          let _rhs383 = (((_2304_v36).contains(p2)) ? ((_2304_v36).get(p2)) : (_2287_v21));
          let _rhs384 = !(p3);
          let _lhs191 = _2287_v21;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_2287_v21).length));
          let _lhs193 = _2299_v33;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2299_v33).length));
          _2298_v32 = _rhs381;
          _lhs191[_lhs192] = _rhs382;
          _lhs193[_lhs194] = _rhs383;
          r0 = _rhs384;
        } else {
          let _2305_v37;
          _2305_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.safeModuloInt(p0, p0));
          _2305_v37 = (_2305_v37).update(new BigNumber(241), p0);
          let _2306_v38;
          _2306_v38 = _dafny.MultiSet.fromElements(p3);
          let _index428 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          (_2287_v21)[_index428] = (_module.__default.fm38(p2, globalState)).IsSubsetOf(_2306_v38);
          let _index429 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          (_2287_v21)[_index429] = p2;
          let _2307_v39;
          let _nw419 = new _module.C4();
          _nw419.__ctor(p0);
          _2307_v39 = _nw419;
          let _2308_v40;
          _2308_v40 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_2309_v15) => function (_2310_i1) {
            return _2309_v15;
          })(_2281_v15)), (_2289_v23).f3, _2278_v12, (_2289_v23).f3, (_2289_v23).f3);
          let _index430 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          let _index431 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          let _rhs385 = _2281_v15;
          let _rhs386 = (_2307_v39).fm8(_dafny.Seq.UnicodeFromString("pinmtrsal"), p3, globalState);
          let _rhs387 = !((_2287_v21)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length))]);
          let _rhs388 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_2311_v23, _2312_p0, _2313_v15) => function (_2314_i2) {
            return _dafny.Seq.update((_module.D6.create_DC11((_2311_v23).f3)).dtor_cf21, _module.__default.safeIndex(_2312_p0, new BigNumber(((_module.D6.create_DC11((_2311_v23).f3)).dtor_cf21).length)), _2313_v15);
          })(_2289_v23, p0, _2281_v15));
          let _lhs195 = _2287_v21;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          let _lhs197 = _2287_v21;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_2287_v21).length));
          _2281_v15 = _rhs385;
          _lhs195[_lhs196] = _rhs386;
          _lhs197[_lhs198] = _rhs387;
          _2308_v40 = _rhs388;
          let _2315_v41;
          _2315_v41 = _dafny.Seq.of(((_2306_v38).update((((_2285_v19).contains(p2)) ? ((_2285_v19).get(p2)) : (p2)), _module.__default.abs(p0))).Union((_2306_v38).update(p3, _module.__default.abs(new BigNumber(669)))));
          _2315_v41 = _2315_v41;
        }
      }
      let _2316_v42;
      _2316_v42 = new _dafny.CodePoint('u'.codePointAt(0));
      let _2317_v43;
      _2317_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      if (!_dafny.Seq.contains(_dafny.Seq.update(_module.__default.fm9(p3, !(p2), true, (_this).fm1((((_2317_v43).contains(p3)) ? ((_2317_v43).get(p3)) : (false)), globalState), globalState), _module.__default.safeIndex((_this).fm1(p3, globalState), new BigNumber((_module.__default.fm9(p3, !(p2), true, (_this).fm1((((_2317_v43).contains(p3)) ? ((_2317_v43).get(p3)) : (false)), globalState), globalState)).length)), _2316_v42), _2316_v42)) {
        let _2318_v44;
        _2318_v44 = _dafny.Seq.UnicodeFromString("tlysn");
        let _2319_v45;
        _2319_v45 = _module.D6.create_DC11(_2318_v44);
        let _2320_v46;
        _2320_v46 = _module.D5.create_DC9(p2, p0, _module.__default.fm38(p3, globalState));
        let _2321_v47;
        _2321_v47 = _dafny.Seq.of(p3, !_dafny.areEqual(_2318_v44, (_2319_v45).dtor_cf21), (_2320_v46).dtor_cf14);
        if ((_2321_v47)[_module.__default.safeIndex(p0, new BigNumber((_2321_v47).length))]) {
          r2 = ((new BigNumber((_dafny.Seq.of(p3)).length)).minus(p0)).plus(_module.__default.safeModuloInt(p0, p0));
          let _2322_v48;
          let _out50;
          _out50 = _module.__default.m0(p2, p3, globalState);
          _2322_v48 = _out50;
          let _2323_v49;
          _2323_v49 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
          let _2324_v50;
          _2324_v50 = _dafny.Seq.of(p0);
          let _2325_v51;
          _2325_v51 = _dafny.Map.Empty.slice().updateUnsafe((((_2323_v49).contains(p2)) ? ((_2323_v49).get(p2)) : (p0)),((p3) ? (_dafny.Seq.update(_2324_v50, _module.__default.safeIndex(p0, new BigNumber((_2324_v50).length)), (_this).fm1(p2, globalState))) : (_2324_v50)));
          _2325_v51 = (_2325_v51).update(p0, _dafny.Seq.Concat(_2324_v50, _2324_v50));
          r0 = p2;
          r0 = !(p3) || (p2);
        } else {
          r0 = (p3) === (p3);
          let _2326_v52;
          let _nw420 = new _module.C3();
          _nw420.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qiesjmhyn"), _2318_v44)).length));
          _2326_v52 = _nw420;
          r0 = !(p3);
          let _2327_v53;
          _2327_v53 = _dafny.Map.Empty.slice().updateUnsafe((p0).isLessThanOrEqualTo(p0),p0);
          _2327_v53 = (_2327_v53).update(!(p2), ((true) ? (p0) : (p0)));
          r2 = p0;
        }
        let _2328_v54;
        let _nw421 = Array((new BigNumber(14)).toNumber());
        _nw421[(_dafny.ZERO).toNumber()] = false;
        _nw421[(_dafny.ONE).toNumber()] = p2;
        _nw421[(new BigNumber(2)).toNumber()] = p3;
        _nw421[(new BigNumber(3)).toNumber()] = p2;
        _nw421[(new BigNumber(4)).toNumber()] = p2;
        _nw421[(new BigNumber(5)).toNumber()] = _module.__default.fm3(p2, p0, p0, globalState);
        _nw421[(new BigNumber(6)).toNumber()] = false;
        _nw421[(new BigNumber(7)).toNumber()] = _module.__default.fm3(p3, p0, new BigNumber(967), globalState);
        _nw421[(new BigNumber(8)).toNumber()] = _module.__default.fm3(false, p0, p0, globalState);
        _nw421[(new BigNumber(9)).toNumber()] = p2;
        _nw421[(new BigNumber(10)).toNumber()] = p3;
        _nw421[(new BigNumber(11)).toNumber()] = p3;
        _nw421[(new BigNumber(12)).toNumber()] = p3;
        _nw421[(new BigNumber(13)).toNumber()] = p2;
        _2328_v54 = _nw421;
        let _2329_v55;
        _2329_v55 = _dafny.Set.fromElements(_2328_v54, _2328_v54, _2328_v54, _2328_v54, _2328_v54);
        let _rhs389 = !(!(p2));
        let _rhs390 = _module.__default.safeDivisionInt((_this).fm1(p2, globalState), p0);
        let _rhs391 = (_dafny.Set.fromElements(_2328_v54)).IsSubsetOf(_2329_v55);
        r0 = _rhs389;
        r2 = _rhs390;
        r0 = _rhs391;
        let _index432 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length));
        (_2328_v54)[_index432] = p3;
        let _2330_v56;
        _2330_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2328_v54,p3);
        let _index433 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length));
        let _rhs392 = p3;
        let _rhs393 = _2330_v56;
        let _rhs394 = p0;
        let _lhs199 = _2328_v54;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length));
        _lhs199[_lhs200] = _rhs392;
        _2330_v56 = _rhs393;
        r2 = _rhs394;
        let _2331_v57;
        _2331_v57 = _dafny.MultiSet.fromElements((_2328_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length))], !(true), !((((_2317_v43).contains((_2328_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length))])) ? ((_2317_v43).get((_2328_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_2328_v54).length))])) : (p2))));
        _2331_v57 = (_dafny.MultiSet.fromElements(p3)).Intersect(_2331_v57);
        let _2332_v58;
        let _nw422 = new _module.C1();
        _nw422.__ctor(_dafny.Seq.update(_2318_v44, _module.__default.safeIndex(p0, new BigNumber((_2318_v44).length)), _2316_v42));
        _2332_v58 = _nw422;
        _2332_v58 = _2332_v58;
      } else {
        r0 = !(p2) || (p2);
        let _2333_v59;
        let _nw423 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2333_v59 = _nw423;
        let _2334_v60;
        _2334_v60 = _dafny.Seq.of(true);
        let _index434 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2333_v59).length));
        (_2333_v59)[_index434] = _module.__default.fm37(new BigNumber((_module.__default.fm49(_2334_v60, p2, p0, globalState)).length), globalState);
        let _2335_v61;
        _2335_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2316_v42,_dafny.Seq.UnicodeFromString("qnctyt"));
        let _2336_v62;
        _2336_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), ((_2337_v42) => function (_2338_i3) {
          return _2337_v42;
        })(_2316_v42))).length),p0);
        let _2339_v63;
        let _nw424 = new _module.C2();
        _nw424.__ctor((_dafny.ZERO).minus(p0));
        _2339_v63 = _nw424;
        let _2340_v64;
        _2340_v64 = _dafny.Set.fromElements(_2339_v63, _2339_v63, _2339_v63);
        let _2341_v65;
        _2341_v65 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2340_v64);
        let _2342_v66;
        _2342_v66 = _dafny.Seq.of(p0, new BigNumber(436), (_this).fm1(true, globalState), new BigNumber((_2341_v65).length));
        let _2343_v67;
        _2343_v67 = _dafny.Seq.of(new BigNumber((_2336_v62).length), new BigNumber((_module.__default.fm15(globalState)).length), (_2342_v66)[_module.__default.safeIndex(_2339_v63.f2, new BigNumber((_2342_v66).length))]);
        let _2344_v68;
        _2344_v68 = _module.D11.create_DC21(_2316_v42, _2339_v63.f2, _2339_v63.f2);
        let _2345_v69;
        _2345_v69 = _module.D11.create_DC22(p2);
        let _2346_v70;
        _2346_v70 = _dafny.MultiSet.fromElements(_2345_v69, _module.D11.create_DC22(p3), _2345_v69);
        let _index435 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2333_v59).length));
        let _rhs395 = (_2344_v68).dtor_cf36;
        let _rhs396 = _2335_v61;
        let _rhs397 = _dafny.Seq.Concat(_2342_v66, _2343_v67);
        let _rhs398 = !((_2346_v70).IsSubsetOf(_2346_v70));
        let _lhs201 = _2333_v59;
        let _lhs202 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2333_v59).length));
        _lhs201[_lhs202] = _rhs395;
        _2335_v61 = _rhs396;
        _2342_v66 = _rhs397;
        r0 = _rhs398;
        let _2347_v71;
        _2347_v71 = _dafny.Seq.of(_2345_v69);
        _2347_v71 = _2347_v71;
        r2 = p0;
        let _source34 = _2345_v69;
        if (_source34.is_DC21) {
          let _2348___mcc_h0 = (_source34).cf36;
          let _2349___mcc_h1 = (_source34).cf37;
          let _2350___mcc_h2 = (_source34).cf38;
          let _2351_cf38 = _2350___mcc_h2;
          let _2352_cf37 = _2349___mcc_h1;
          let _2353_cf36 = _2348___mcc_h0;
          let _2354_v72;
          let _init50 = ((_2355_cf37) => function (_2356_i4) {
            return _dafny.Set.fromElements(new BigNumber(-495), _2355_cf37, _2355_cf37);
          })(_2352_cf37);
          let _nw425 = Array((new BigNumber(7)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw425.length); _i0_50++) {
            _nw425[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _2354_v72 = _nw425;
          let _2357_v74;
          _2357_v74 = _dafny.Set.fromElements(new BigNumber((_2317_v43).length));
          let _index436 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_2354_v72).length));
          (_2354_v72)[_index436] = (function () {
            let _coll72 = new _dafny.Set();
            for (const _compr_72 of _dafny.IntegerRange(new BigNumber(-245), new BigNumber(-526))) {
              let _2358_v73 = _compr_72;
              if (((new BigNumber(-245)).isLessThanOrEqualTo(_2358_v73)) && ((_2358_v73).isLessThan(new BigNumber(-526)))) {
                _coll72.add((_2358_v73).minus(_2339_v63.f2));
              }
            }
            return _coll72;
          }()).Union(_2357_v74);
          let _index437 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_2354_v72).length));
          (_2354_v72)[_index437] = _2357_v74;
          r0 = p2;
          _2352_cf37 = _2351_cf38;
          let _2359_v75;
          _2359_v75 = _dafny.Seq.UnicodeFromString("ogsgwnlyp");
          let _2360_v76;
          _2360_v76 = _dafny.Set.fromElements(_module.__default.fm5(globalState));
          let _2361_v77;
          let _nw426 = Array((new BigNumber(29)).toNumber());
          _nw426[(_dafny.ZERO).toNumber()] = _2359_v75;
          _nw426[(_dafny.ONE).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(2)).toNumber()] = _module.__default.fm28(p2, p2, p2, _2360_v76, globalState);
          _nw426[(new BigNumber(3)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(4)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(5)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(6)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), ((_2362_cf36) => function (_2363_i5) {
            return _2362_cf36;
          })(_2353_cf36));
          _nw426[(new BigNumber(8)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("myypb");
          _nw426[(new BigNumber(10)).toNumber()] = ((p2) ? (_2359_v75) : (_2359_v75));
          _nw426[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_2359_v75, _2359_v75);
          _nw426[(new BigNumber(12)).toNumber()] = ((p2) ? (_2359_v75) : (_2359_v75));
          _nw426[(new BigNumber(13)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(14)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_2359_v75, _2359_v75);
          _nw426[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), ((_2364_cf36) => function (_2365_i6) {
            return _2364_cf36;
          })(_2353_cf36));
          _nw426[(new BigNumber(17)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(18)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(19)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(20)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(21)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_2359_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), ((_2366_cf36) => function (_2367_i7) {
            return _2366_cf36;
          })(_2353_cf36)));
          _nw426[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_2359_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), ((_2368_v75, _2369_v43) => function (_2370_i8) {
            return (_2368_v75)[_module.__default.safeIndex(new BigNumber((_2369_v43).length), new BigNumber((_2368_v75).length))];
          })(_2359_v75, _2317_v43)));
          _nw426[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kcnc"), _dafny.Seq.UnicodeFromString("s"));
          _nw426[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_2359_v75, _2359_v75);
          _nw426[(new BigNumber(26)).toNumber()] = _2359_v75;
          _nw426[(new BigNumber(27)).toNumber()] = ((p2) ? (_2359_v75) : (_2359_v75));
          _nw426[(new BigNumber(28)).toNumber()] = _2359_v75;
          _2361_v77 = _nw426;
          let _index438 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_2361_v77).length));
          (_2361_v77)[_index438] = _2359_v75;
          let _index439 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_2361_v77).length));
          (_2361_v77)[_index439] = _2359_v75;
        } else if (_source34.is_DC22) {
          let _2371___mcc_h3 = (_source34).cf39;
          let _2372_cf39 = _2371___mcc_h3;
          let _2373_v78;
          _2373_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2339_v63.f2,true);
          let _2374_v79;
          _2374_v79 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cktdo"),_2373_v78);
          let _2375_v80;
          _2375_v80 = _module.D5.create_DC10(p2, p0, _2374_v79, !(p3));
          let _rhs399 = _module.__default.fm29(globalState);
          let _rhs400 = !((true) === (p3));
          _2375_v80 = _rhs399;
          r0 = _rhs400;
          let _2376_v81;
          let _init51 = ((_2377_v66) => function (_2378_i9) {
            return _2377_v66;
          })(_2342_v66);
          let _nw427 = Array((new BigNumber(28)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw427.length); _i0_51++) {
            _nw427[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _2376_v81 = _nw427;
          _2376_v81 = _2376_v81;
          let _2379_v82;
          _2379_v82 = _dafny.Seq.UnicodeFromString("odck");
          let _2380_v83;
          let _nw428 = new _module.C1();
          _nw428.__ctor(_2379_v82);
          _2380_v83 = _nw428;
          let _rhs401 = _module.__default.safeModuloInt(p0, _2339_v63.f2);
          let _rhs402 = _2380_v83;
          let _rhs403 = new _dafny.CodePoint('j'.codePointAt(0));
          let _lhs203 = _2339_v63;
          _lhs203.f2 = _rhs401;
          _2380_v83 = _rhs402;
          _2316_v42 = _rhs403;
          let _2381_v84;
          _2381_v84 = _module.D6.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-691)), ((_2382_v59) => function (_2383_i10) {
  return (_2382_v59)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_2382_v59).length))];
})(_2333_v59)));
          _2379_v82 = (_2381_v84).dtor_cf21;
        } else {
          let _2384___mcc_h4 = (_source34).cf35;
          let _2385_cf35 = _2384___mcc_h4;
          r0 = p3;
          let _2386_v85;
          let _nw429 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2386_v85 = _nw429;
          let _2387_v86;
          _2387_v86 = _dafny.Seq.UnicodeFromString("gcqafja");
          let _index440 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_2386_v85).length));
          (_2386_v85)[_index440] = _2387_v86;
          let _index441 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_2386_v85).length));
          (_2386_v85)[_index441] = ((p2) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), ((_2388_v59) => function (_2389_i11) {
            return (_2388_v59)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_2388_v59).length))];
          })(_2333_v59))) : (_2387_v86));
          let _2390_v87;
          let _nw430 = Array((new BigNumber(14)).toNumber());
          _nw430[(_dafny.ZERO).toNumber()] = _2339_v63.f2;
          _nw430[(_dafny.ONE).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(2)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(3)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(4)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(5)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(6)).toNumber()] = p0;
          _nw430[(new BigNumber(7)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(8)).toNumber()] = new BigNumber(-400);
          _nw430[(new BigNumber(9)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(10)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(11)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(12)).toNumber()] = _2339_v63.f2;
          _nw430[(new BigNumber(13)).toNumber()] = new BigNumber(((_2386_v85)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_2386_v85).length))]).length);
          _2390_v87 = _nw430;
          let _2391_v88;
          _2391_v88 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2390_v87);
          let _2392_v89;
          let _nw431 = new _module.C1();
          _nw431.__ctor((_2386_v85)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_2386_v85).length))]);
          _2392_v89 = _nw431;
          let _2393_v90;
          _2393_v90 = _dafny.MultiSet.fromElements(_2392_v89, _2392_v89);
          let _2394_v91;
          _2394_v91 = _dafny.Set.fromElements(new BigNumber((_2391_v88).length), p0, new BigNumber((_2393_v90).cardinality()), (_dafny.ZERO).minus(p0), new BigNumber(-947));
          r0 = !((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), ((_2395_v42) => function (_2396_i12) {
            return _2395_v42;
          })(_2316_v42)), _2387_v86)) === ((_2339_v63.f2).isEqualTo(new BigNumber((_2394_v91).length))));
          (_2339_v63).f2 = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(_2339_v63.f2));
        }
      }
      let _2397_v92;
      let _nw432 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2397_v92 = _nw432;
      let _rhs404 = new BigNumber((function () {
        let _coll73 = new _dafny.Set();
        for (const _compr_73 of _dafny.IntegerRange(new BigNumber(182), new BigNumber(266))) {
          let _2398_v93 = _compr_73;
          if (((new BigNumber(182)).isLessThanOrEqualTo(_2398_v93)) && ((_2398_v93).isLessThan(new BigNumber(266)))) {
            _coll73.add((_2398_v93).minus(p0));
          }
        }
        return _coll73;
      }()).length);
      let _rhs405 = _2397_v92;
      r2 = _rhs404;
      _2397_v92 = _rhs405;
      let _2399_v94;
      let _nw433 = new _module.C6();
      _nw433.__ctor();
      _2399_v94 = _nw433;
      let _2400_v95;
      _2400_v95 = _dafny.Seq.of(_2316_v42);
      let _2401_v96;
      _2401_v96 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.update(_2400_v95, _module.__default.safeIndex(p0, new BigNumber((_2400_v95).length)), _2316_v42)).length));
      let _2402_v97;
      _2402_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2399_v94,_2401_v96);
      let _2403_v98;
      _2403_v98 = _dafny.Set.fromElements(p0, p0);
      let _2404_v99;
      _2404_v99 = _dafny.Seq.of(true);
      let _2405_v100;
      _2405_v100 = _dafny.MultiSet.fromElements(p2);
      let _2406_v101;
      let _nw434 = Array((new BigNumber(14)).toNumber());
      _nw434[(_dafny.ZERO).toNumber()] = (_module.D6.create_DC13(p3, p0)).dtor_cf25;
      _nw434[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((p3) ? ((_dafny.ZERO).minus(p0)) : (p0))));
      _nw434[(new BigNumber(2)).toNumber()] = p0;
      _nw434[(new BigNumber(3)).toNumber()] = new BigNumber(((((_2402_v97).contains(_2399_v94)) ? ((_2402_v97).get(_2399_v94)) : (_2401_v96))).length);
      _nw434[(new BigNumber(4)).toNumber()] = new BigNumber(((_2403_v98).Intersect(_2403_v98)).length);
      _nw434[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2404_v99).length), p0);
      _nw434[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
      _nw434[(new BigNumber(7)).toNumber()] = p0;
      _nw434[(new BigNumber(8)).toNumber()] = new BigNumber((_2266_v1).cardinality());
      _nw434[(new BigNumber(9)).toNumber()] = new BigNumber((_2405_v100).cardinality());
      _nw434[(new BigNumber(10)).toNumber()] = p0;
      _nw434[(new BigNumber(11)).toNumber()] = p0;
      _nw434[(new BigNumber(12)).toNumber()] = p0;
      _nw434[(new BigNumber(13)).toNumber()] = p0;
      _2406_v101 = _nw434;
      let _2407_v102;
      _2407_v102 = _dafny.Seq.of((_this).fm1(p2, globalState), p0, p0, (_dafny.ZERO).minus(p0), new BigNumber((_2404_v99).length));
      let _index442 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_2406_v101).length));
      (_2406_v101)[_index442] = (_dafny.ZERO).minus((_module.__default.fm50(!(p2), !(!(p3)), p2, (_2407_v102)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2317_v43).length)), new BigNumber((_2407_v102).length))], globalState)).dtor_cf27);
      let _2408_v103;
      _2408_v103 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      let _index443 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_2406_v101).length));
      (_2406_v101)[_index443] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_2408_v103).contains(p3)) ? ((_2408_v103).get(p3)) : (p0)), new BigNumber((_2407_v102).length))));
      let _2409_v104;
      let _nw435 = Array((new BigNumber(11)).toNumber()).fill(false);
      _2409_v104 = _nw435;
      let _index444 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_2409_v104).length));
      (_2409_v104)[_index444] = p3;
      let _2410_v105;
      _2410_v105 = _dafny.Map.Empty.slice().updateUnsafe((_2406_v101)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_2406_v101).length))],_2405_v100);
      let _index445 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_2409_v104).length));
      (_2409_v104)[_index445] = (_2405_v100).IsProperSubsetOf((((_2410_v105).contains(p0)) ? ((_2410_v105).get(p0)) : (_2405_v100)));
      let _index446 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_2406_v101).length));
      (_2406_v101)[_index446] = new BigNumber((_2404_v99).length);
      r0 = p3;
      let _2411_v106;
      let _nw436 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
      _2411_v106 = _nw436;
      r1 = _2411_v106;
      r2 = (p0).multipliedBy(p0);
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2412_v0;
      _2412_v0 = true;
      _2412_v0 = p2;
      if (false) {
        let _2413_v1;
        _2413_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        _2413_v1 = (_2413_v1).update(_module.__default.safeDivisionInt(new BigNumber(-409), p0), _2412_v0);
        let _2414_v2;
        let _nw437 = new _module.C3();
        _nw437.__ctor(((_dafny.ZERO).minus(p1)).plus(p0));
        _2414_v2 = _nw437;
        _2414_v2 = _2414_v2;
        let _2415_v3;
        _2415_v3 = _dafny.MultiSet.fromElements(!(true), p2, _2412_v0, !(!(_2412_v0)));
        let _2416_v4;
        _2416_v4 = _dafny.Seq.of((_2415_v3).update(true, _module.__default.abs(new BigNumber(525))));
        _2416_v4 = _dafny.Seq.of((_2415_v3).Union((_2416_v4)[_module.__default.safeIndex(p0, new BigNumber((_2416_v4).length))]));
        let _2417_v5;
        _2417_v5 = new _dafny.CodePoint('r'.codePointAt(0));
        let _2418_v6;
        _2418_v6 = _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)), _2417_v5);
        let _2419_v7;
        _2419_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2418_v6, _dafny.Seq.update(_dafny.Seq.of(_2417_v5, new _dafny.CodePoint('k'.codePointAt(0))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_2417_v5, new _dafny.CodePoint('k'.codePointAt(0)))).length)), _2417_v5)),_2412_v0);
        if ((((_2419_v7).contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), ((_2424_v5) => function (_2425_i0) {
          return _2424_v5;
        })(_2417_v5)), _module.__default.safeIndex(new BigNumber((_2418_v6).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), ((_2426_v5) => function (_2427_i0) {
          return _2426_v5;
        })(_2417_v5))).length)), _2417_v5))) ? ((_2419_v7).get(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), ((_2420_v5) => function (_2421_i0) {
          return _2420_v5;
        })(_2417_v5)), _module.__default.safeIndex(new BigNumber((_2418_v6).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), ((_2422_v5) => function (_2423_i0) {
          return _2422_v5;
        })(_2417_v5))).length)), _2417_v5))) : (!(p2)))) {
          let _2428_v8;
          _2428_v8 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(227)), ((_2429_v5) => function (_2430_i1) {
            return _2429_v5;
          })(_2417_v5)), _2418_v6, _2418_v6, _2418_v6);
          let _2431_v9;
          let _nw438 = Array((new BigNumber(5)).toNumber());
          _nw438[(_dafny.ZERO).toNumber()] = _2414_v2.f2;
          _nw438[(_dafny.ONE).toNumber()] = new BigNumber(-808);
          _nw438[(new BigNumber(2)).toNumber()] = _2414_v2.f2;
          _nw438[(new BigNumber(3)).toNumber()] = new BigNumber((_2428_v8).length);
          _nw438[(new BigNumber(4)).toNumber()] = p0;
          _2431_v9 = _nw438;
          let _index447 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2431_v9).length));
          (_2431_v9)[_index447] = p0;
          let _index448 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2431_v9).length));
          (_2431_v9)[_index448] = (_2414_v2.f2).plus(p1);
          let _2432_v10;
          _2432_v10 = _dafny.Seq.of(p1);
          let _2433_v11;
          _2433_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2432_v10,_2414_v2.f2);
          _2433_v11 = _2433_v11;
          let _2434_v12;
          _2434_v12 = _dafny.Set.fromElements(_2412_v0, p2);
          let _rhs406 = !(_module.__default.fm3(p2, ((p2) ? (new BigNumber((_dafny.Seq.of(false, p2)).length)) : ((_2431_v9)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_2431_v9).length))])), new BigNumber((_module.__default.fm15(globalState)).length), globalState));
          let _rhs407 = (new BigNumber((_2434_v12).length)).plus((_dafny.ZERO).minus(_2414_v2.f2));
          let _rhs408 = _2412_v0;
          let _rhs409 = _2418_v6;
          let _lhs204 = _2414_v2;
          _2412_v0 = _rhs406;
          _lhs204.f2 = _rhs407;
          _2412_v0 = _rhs408;
          _2418_v6 = _rhs409;
          _2417_v5 = _2417_v5;
          _2412_v0 = _2412_v0;
        } else {
          _2412_v0 = (_2414_v2.f2).isLessThanOrEqualTo(new BigNumber((_2413_v1).length));
          let _2435_v13;
          let _nw439 = new _module.C4();
          _nw439.__ctor(_module.__default.safeModuloInt(p1, new BigNumber((_2415_v3).cardinality())));
          _2435_v13 = _nw439;
          let _2436_v14;
          _2436_v14 = _dafny.Seq.of(p2, p2);
          let _2437_v15;
          _2437_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2436_v14,p1);
          _2437_v15 = (_2437_v15).update(_2436_v14, (_dafny.ZERO).minus((_2435_v13).fm1(_2412_v0, globalState)));
          let _2438_v16;
          let _nw440 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2438_v16 = _nw440;
          let _init52 = ((_2439_v6) => function (_2440_i2) {
            return _2439_v6;
          })(_2418_v6);
          let _nw441 = Array((new BigNumber(9)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw441.length); _i0_52++) {
            _nw441[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _2438_v16 = _nw441;
          let _2441_v17;
          _2441_v17 = _dafny.Seq.of(_2414_v2.f2, _2414_v2.f2);
          let _2442_v18;
          _2442_v18 = _dafny.MultiSet.fromElements(_2441_v17);
          let _2443_v19;
          _2443_v19 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_2412_v0)).IsProperSubsetOf(_2415_v3),_2442_v18);
          _2443_v19 = (_2443_v19).update((_2414_v2.f2).isLessThan(new BigNumber(-163)), _dafny.MultiSet.fromElements(_2441_v17, _2441_v17, _2441_v17));
        }
        if (!((p2) === ((_2414_v2).fm8(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("fvqbna"), _module.__default.safeIndex(_2414_v2.f2, new BigNumber((_dafny.Seq.UnicodeFromString("fvqbna")).length)), new _dafny.CodePoint('k'.codePointAt(0))), _2412_v0, globalState)))) {
          let _2444_v20;
          _2444_v20 = _dafny.MultiSet.fromElements(_2414_v2.f2);
          let _2445_v21;
          _2445_v21 = _dafny.Seq.of(_2418_v6);
          let _2446_v22;
          _2446_v22 = _dafny.Seq.of(_2412_v0, p2);
          let _2447_v23;
          _2447_v23 = _module.D6.create_DC12((_2446_v22)[_module.__default.safeIndex(p1, new BigNumber((_2446_v22).length))], _2412_v0);
          let _2448_v24;
          _2448_v24 = _dafny.Set.fromElements(_2414_v2.f2, _2414_v2.f2);
          let _2449_v25;
          _2449_v25 = _module.D11.create_DC21(_2417_v5, _2414_v2.f2, p0);
          let _2450_v26;
          let _nw442 = Array((new BigNumber(10)).toNumber());
          _nw442[(_dafny.ZERO).toNumber()] = _2412_v0;
          _nw442[(_dafny.ONE).toNumber()] = p2;
          _nw442[(new BigNumber(2)).toNumber()] = _2412_v0;
          _nw442[(new BigNumber(3)).toNumber()] = !(p2);
          _nw442[(new BigNumber(4)).toNumber()] = _2412_v0;
          _nw442[(new BigNumber(5)).toNumber()] = _2412_v0;
          _nw442[(new BigNumber(6)).toNumber()] = _2412_v0;
          _nw442[(new BigNumber(7)).toNumber()] = true;
          _nw442[(new BigNumber(8)).toNumber()] = p2;
          _nw442[(new BigNumber(9)).toNumber()] = p2;
          _2450_v26 = _nw442;
          let _2451_v27;
          _2451_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2450_v26,_2414_v2.f2);
          let _2452_v28;
          let _nw443 = new _module.C1();
          _nw443.__ctor(_2418_v6);
          _2452_v28 = _nw443;
          let _2453_v29;
          _2453_v29 = _module.D8.create_DC17(_2412_v0, _2452_v28, p1);
          let _2454_v30;
          _2454_v30 = _dafny.Seq.of(p0);
          let _2455_v31;
          let _nw444 = Array((new BigNumber(26)).toNumber());
          _nw444[(_dafny.ZERO).toNumber()] = (_2414_v2.f2).plus(p1);
          _nw444[(_dafny.ONE).toNumber()] = new BigNumber((_2444_v20).cardinality());
          _nw444[(new BigNumber(2)).toNumber()] = (_2414_v2.f2).plus(_2414_v2.f2);
          _nw444[(new BigNumber(3)).toNumber()] = new BigNumber(-705);
          _nw444[(new BigNumber(4)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(5)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_2445_v21)[_module.__default.safeIndex((_this).fm1(p2, globalState), new BigNumber((_2445_v21).length))], _2418_v6)).length);
          _nw444[(new BigNumber(7)).toNumber()] = new BigNumber((((p2) ? (_2418_v6) : (_2418_v6))).length);
          _nw444[(new BigNumber(8)).toNumber()] = p1;
          _nw444[(new BigNumber(9)).toNumber()] = ((p2) ? (_2414_v2.f2) : (p0));
          _nw444[(new BigNumber(10)).toNumber()] = p0;
          _nw444[(new BigNumber(11)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(12)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(13)).toNumber()] = (((_2447_v23).dtor_cf23) ? (_2414_v2.f2) : (_2414_v2.f2));
          _nw444[(new BigNumber(14)).toNumber()] = (new BigNumber((_2446_v22).length)).minus(_2414_v2.f2);
          _nw444[(new BigNumber(15)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(16)).toNumber()] = (_2414_v2.f2).multipliedBy(_2414_v2.f2);
          _nw444[(new BigNumber(17)).toNumber()] = _2414_v2.f2;
          _nw444[(new BigNumber(18)).toNumber()] = (p0).multipliedBy(new BigNumber((_2448_v24).length));
          _nw444[(new BigNumber(19)).toNumber()] = (_2449_v25).dtor_cf37;
          _nw444[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt(_2414_v2.f2, p0);
          _nw444[(new BigNumber(21)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_2451_v27).length));
          _nw444[(new BigNumber(22)).toNumber()] = (_2453_v29).dtor_cf32;
          _nw444[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(_2414_v2.f2, p0);
          _nw444[(new BigNumber(24)).toNumber()] = p0;
          _nw444[(new BigNumber(25)).toNumber()] = new BigNumber((_2454_v30).length);
          _2455_v31 = _nw444;
          let _index449 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_2455_v31).length));
          (_2455_v31)[_index449] = p0;
          let _2456_v32;
          _2456_v32 = _dafny.Set.fromElements(_2450_v26);
          let _index450 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_2455_v31).length));
          (_2455_v31)[_index450] = (new BigNumber((_2456_v32).length)).multipliedBy((_dafny.ZERO).minus(_2414_v2.f2));
          let _2457_v33;
          let _nw445 = new _module.C2();
          _nw445.__ctor(_2414_v2.f2);
          _2457_v33 = _nw445;
          (_2414_v2).f2 = p1;
          let _2458_v34;
          let _nw446 = Array((new BigNumber(25)).toNumber());
          _nw446[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), ((_2459_v2) => function (_2460_i3) {
            return _2459_v2.f2;
          })(_2414_v2));
          _nw446[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), ((_2461_p1) => function (_2462_i4) {
            return _2461_p1;
          })(p1));
          _nw446[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_2414_v2.f2);
          _nw446[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-633)), ((_2463_p0) => function (_2464_i5) {
            return _2463_p0;
          })(p0)), _module.__default.safeIndex(_2414_v2.f2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-633)), ((_2465_p0) => function (_2466_i5) {
            return _2465_p0;
          })(p0))).length)), new BigNumber((_2413_v1).length));
          _nw446[(new BigNumber(4)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(new BigNumber(((_2452_v28).f3).length), p0);
          _nw446[(new BigNumber(6)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(7)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(8)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(9)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(new BigNumber((_2454_v30).length), new BigNumber((_2454_v30).length));
          _nw446[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), ((_2467_p1) => function (_2468_i6) {
            return _2467_p1;
          })(p1));
          _nw446[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(new BigNumber((_2415_v3).cardinality()));
          _nw446[(new BigNumber(13)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(14)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(15)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(16)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(17)).toNumber()] = _2454_v30;
          _nw446[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_2414_v2.f2, _2414_v2.f2);
          _nw446[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_2414_v2.f2);
          _nw446[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_2414_v2.f2);
          _nw446[(new BigNumber(21)).toNumber()] = _dafny.Seq.of((_2455_v31)[_module.__default.safeIndex(new BigNumber(276), new BigNumber((_2455_v31).length))], (_2457_v33).fm1(p2, globalState));
          _nw446[(new BigNumber(22)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(54)), ((_2469_v2) => function (_2470_i7) {
            return _2469_v2.f2;
          })(_2414_v2));
          _nw446[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(383)), ((_2471_v2) => function (_2472_i8) {
            return _2471_v2.f2;
          })(_2414_v2));
          _nw446[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), ((_2473_v2) => function (_2474_i9) {
            return _2473_v2.f2;
          })(_2414_v2));
          _2458_v34 = _nw446;
          let _2475_v35;
          let _2476_v36;
          let _2477_v37;
          let _out51;
          let _out52;
          let _out53;
          let _outcollector13 = (_2452_v28).m1((new BigNumber((_2454_v30).length)).plus(p1), _2458_v34, p2, p2, globalState);
          _out51 = _outcollector13[0];
          _out52 = _outcollector13[1];
          _out53 = _outcollector13[2];
          _2475_v35 = _out51;
          _2476_v36 = _out52;
          _2477_v37 = _out53;
          let _2478_v38;
          _2478_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.safeDivisionInt((_2414_v2).fm1((_2452_v28).fm11(globalState), globalState), _2414_v2.f2));
          _2478_v38 = (_2478_v38).update(_2412_v0, _2477_v37);
        } else {
          _2412_v0 = !(p2);
          _2412_v0 = _2412_v0;
          r0 = _2414_v2.f2;
          let _2479_v39;
          _2479_v39 = _dafny.Seq.of(p0, _2414_v2.f2, _2414_v2.f2, p0);
          let _2480_v40;
          _2480_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2417_v5,new BigNumber((_2479_v39).length));
          _2480_v40 = ((_dafny.Map.Empty.slice().updateUnsafe(_2417_v5,p1)).Merge(_2480_v40)).update(_2417_v5, p0);
          let _2481_v41;
          let _nw447 = Array((new BigNumber(11)).toNumber()).fill([]);
          _2481_v41 = _nw447;
          let _2482_v43;
          _2482_v43 = _dafny.Set.fromElements(_2414_v2.f2, (_2414_v2).fm1(true, globalState), p1);
          let _rhs410 = new BigNumber(((function () {
            let _coll74 = new _dafny.Set();
            for (const _compr_74 of _dafny.IntegerRange(new BigNumber(329), new BigNumber(686))) {
              let _2483_v42 = _compr_74;
              if (((new BigNumber(329)).isLessThanOrEqualTo(_2483_v42)) && ((_2483_v42).isLessThan(new BigNumber(686)))) {
                _coll74.add((_2483_v42).minus(new BigNumber(270)));
              }
            }
            return _coll74;
          }()).Union(_2482_v43)).length);
          let _rhs411 = _2481_v41;
          let _rhs412 = p0;
          r0 = _rhs410;
          _2481_v41 = _rhs411;
          r0 = _rhs412;
        }
      } else {
        let _2484_v44;
        _2484_v44 = _dafny.Seq.UnicodeFromString("jy");
        let _2485_v45;
        _2485_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2412_v0,p0);
        let _2486_v47;
        _2486_v47 = _dafny.Seq.of(new BigNumber((_2484_v44).length));
        let _2487_v48;
        _2487_v48 = _dafny.Seq.of(_2412_v0, true);
        let _2488_v49;
        _2488_v49 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2486_v47)[_module.__default.safeIndex(new BigNumber((_2487_v48).length), new BigNumber((_2486_v47).length))]);
        let _2489_v50;
        _2489_v50 = _dafny.Set.fromElements(p0, new BigNumber((_2488_v49).length));
        let _2490_v51;
        _2490_v51 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_2484_v44).length), p0, (((_2485_v45).contains(_2412_v0)) ? ((_2485_v45).get(_2412_v0)) : (p1))), function () {
          let _coll75 = new _dafny.Set();
          for (const _compr_75 of _dafny.IntegerRange(new BigNumber(345), new BigNumber(363))) {
            let _2491_v46 = _compr_75;
            if (((new BigNumber(345)).isLessThanOrEqualTo(_2491_v46)) && ((_2491_v46).isLessThan(new BigNumber(363)))) {
              _coll75.add((_2491_v46).plus(p1));
            }
          }
          return _coll75;
        }(), _2489_v50);
        let _2492_v52;
        _2492_v52 = _module.D7.create_DC14(_2486_v47);
        let _2493_v53;
        _2493_v53 = _dafny.Seq.of(_2492_v52, _2492_v52);
        let _2494_v54;
        let _init53 = ((_2495_p2, _2496_p1, _2497_v44) => function (_2498_i10) {
          return (_2498_i10).plus((_module.D5.create_DC10(_2495_p2, _2496_p1, _dafny.Map.Empty.slice().updateUnsafe(_2497_v44,_dafny.Map.Empty.slice().updateUnsafe(_2496_p1,_2495_p2)), _2495_p2)).dtor_cf18);
        })(p2, p1, _2484_v44);
        let _nw448 = Array((new BigNumber(5)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw448.length); _i0_53++) {
          _nw448[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2494_v54 = _nw448;
        let _2499_v55;
        _2499_v55 = _dafny.Seq.of(_2494_v54);
        let _2500_v56;
        _2500_v56 = _dafny.MultiSet.fromElements(_2412_v0, false, p2, _2412_v0);
        let _2501_v57;
        let _nw449 = Array((new BigNumber(27)).toNumber());
        _nw449[(_dafny.ZERO).toNumber()] = (new BigNumber((_2490_v51).length)).isLessThan(p0);
        _nw449[(_dafny.ONE).toNumber()] = ((_this).fm1(_2412_v0, globalState)).isLessThan(new BigNumber((_2493_v53).length));
        _nw449[(new BigNumber(2)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(3)).toNumber()] = (_2412_v0) === (p2);
        _nw449[(new BigNumber(4)).toNumber()] = (true) && (!(false));
        _nw449[(new BigNumber(5)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(6)).toNumber()] = p2;
        _nw449[(new BigNumber(7)).toNumber()] = (new BigNumber(-180)).isEqualTo(new BigNumber((_2499_v55).length));
        _nw449[(new BigNumber(8)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(9)).toNumber()] = !(p2);
        _nw449[(new BigNumber(10)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(11)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(12)).toNumber()] = (_2412_v0) && (false);
        _nw449[(new BigNumber(13)).toNumber()] = p2;
        _nw449[(new BigNumber(14)).toNumber()] = (_2500_v56).IsSubsetOf(_2500_v56);
        _nw449[(new BigNumber(15)).toNumber()] = p2;
        _nw449[(new BigNumber(16)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(17)).toNumber()] = _module.__default.fm3(_2412_v0, p1, p1, globalState);
        _nw449[(new BigNumber(18)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.of((_this).fm1(_2412_v0, globalState), (_2486_v47)[_module.__default.safeIndex(new BigNumber((_2484_v44).length), new BigNumber((_2486_v47).length))]), p0);
        _nw449[(new BigNumber(19)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
        _nw449[(new BigNumber(20)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(21)).toNumber()] = false;
        _nw449[(new BigNumber(22)).toNumber()] = !(p2) || (_2412_v0);
        _nw449[(new BigNumber(23)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(24)).toNumber()] = p2;
        _nw449[(new BigNumber(25)).toNumber()] = _2412_v0;
        _nw449[(new BigNumber(26)).toNumber()] = _dafny.areEqual(_2484_v44, _dafny.Seq.UnicodeFromString("krmp"));
        _2501_v57 = _nw449;
        let _index451 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
        (_2501_v57)[_index451] = _module.__default.fm5(globalState);
        let _index452 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
        let _rhs413 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2484_v44, _2484_v44), _2484_v44);
        let _rhs414 = p0;
        let _rhs415 = _2412_v0;
        let _lhs205 = _2501_v57;
        let _lhs206 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
        _2484_v44 = _rhs413;
        r0 = _rhs414;
        _lhs205[_lhs206] = _rhs415;
        if (!((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))])) {
          let _2502_v58;
          let _nw450 = Array((new BigNumber(14)).toNumber()).fill([]);
          _2502_v58 = _nw450;
          let _index453 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2502_v58).length));
          (_2502_v58)[_index453] = _2494_v54;
          let _index454 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2502_v58).length));
          (_2502_v58)[_index454] = _2494_v54;
          let _2503_v59;
          _2503_v59 = _dafny.Map.Empty.slice().updateUnsafe((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))],new _dafny.CodePoint('r'.codePointAt(0)));
          let _2504_v60;
          _2504_v60 = _dafny.Set.fromElements(_2503_v59);
          let _index455 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _rhs416 = p0;
          let _rhs417 = (_2504_v60).Union((_dafny.Set.fromElements(_2503_v59)).Union(_2504_v60));
          let _rhs418 = _2412_v0;
          let _rhs419 = (p1).isEqualTo(new BigNumber((_dafny.Set.fromElements((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))])).length));
          let _lhs207 = _2501_v57;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          r0 = _rhs416;
          _2504_v60 = _rhs417;
          _2412_v0 = _rhs418;
          _lhs207[_lhs208] = _rhs419;
          let _2505_v61;
          _2505_v61 = new _dafny.CodePoint('b'.codePointAt(0));
          let _2506_v62;
          _2506_v62 = _module.D5.create_DC9((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], new BigNumber(69), _2500_v56);
          let _index456 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _rhs420 = (p2) && ((_2487_v48)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2487_v48).length))]);
          let _rhs421 = (_dafny.ZERO).minus((((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))]) ? ((_2506_v62).dtor_cf15) : (p0)));
          let _rhs422 = _2505_v61;
          let _lhs209 = _2501_v57;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          _lhs209[_lhs210] = _rhs420;
          r0 = _rhs421;
          _2505_v61 = _rhs422;
          let _2507_v63;
          let _nw451 = new _module.C6();
          _nw451.__ctor();
          _2507_v63 = _nw451;
          let _index457 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _rhs423 = _2507_v63;
          let _rhs424 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ongtjwil"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(959)), ((_2508_v61) => function (_2509_i11) {
            return _2508_v61;
          })(_2505_v61)));
          let _rhs425 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2487_v48, _2487_v48), _dafny.Seq.of(_2412_v0, false));
          let _rhs426 = (_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))];
          let _lhs211 = _2501_v57;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          _2507_v63 = _rhs423;
          _2484_v44 = _rhs424;
          _2487_v48 = _rhs425;
          _lhs211[_lhs212] = _rhs426;
          let _2510_v64;
          let _nw452 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
          _2510_v64 = _nw452;
          let _index458 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2510_v64).length));
          (_2510_v64)[_index458] = _2489_v50;
          let _index459 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _index460 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2510_v64).length));
          let _rhs427 = !((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))]);
          let _rhs428 = (p1).isLessThanOrEqualTo(((_2412_v0) ? (p1) : (new BigNumber((_2503_v59).length))));
          let _rhs429 = _2501_v57;
          let _rhs430 = _dafny.Set.fromElements(new BigNumber(353), p1, (p0).plus(new BigNumber(573)));
          let _rhs431 = _2412_v0;
          let _lhs213 = _2501_v57;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _lhs215 = _2510_v64;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2510_v64).length));
          _2412_v0 = _rhs427;
          _lhs213[_lhs214] = _rhs428;
          _2501_v57 = _rhs429;
          _lhs215[_lhs216] = _rhs430;
          _2412_v0 = _rhs431;
        } else {
          let _2511_v65;
          _2511_v65 = _module.D6.create_DC13((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], (_this).fm1(_2412_v0, globalState));
          _2412_v0 = (_2511_v65).dtor_cf24;
          let _2512_v66;
          _2512_v66 = new _dafny.CodePoint('w'.codePointAt(0));
          _2485_v45 = _module.__default.fm47(_2512_v66, new BigNumber((_dafny.Seq.update(_2487_v48, _module.__default.safeIndex(p0, new BigNumber((_2487_v48).length)), !((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))]))).length), p1, globalState);
          let _index461 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          (_2501_v57)[_index461] = ((_2500_v56).Difference(_2500_v56)).IsProperSubsetOf(_2500_v56);
          let _2513_v67;
          _2513_v67 = _dafny.Set.fromElements(_2412_v0);
          let _index462 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          let _rhs432 = _dafny.Seq.Concat(_2484_v44, _dafny.Seq.Concat(_2484_v44, _2484_v44));
          let _rhs433 = ((_dafny.Set.fromElements(true, _2412_v0, (_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], _2412_v0, p2)).Intersect(_2513_v67)).IsDisjointFrom(_2513_v67);
          let _rhs434 = !(((p0).multipliedBy(new BigNumber((_module.__default.fm49(_2487_v48, (_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], p0, globalState)).length))).isLessThan(new BigNumber((function () {
            let _coll76 = new _dafny.Set();
            for (const _compr_76 of _dafny.IntegerRange(new BigNumber(879), new BigNumber(273))) {
              let _2514_v68 = _compr_76;
              if (((new BigNumber(879)).isLessThanOrEqualTo(_2514_v68)) && ((_2514_v68).isLessThan(new BigNumber(273)))) {
                _coll76.add(_module.__default.safeDivisionInt(_2514_v68, p1));
              }
            }
            return _coll76;
          }()).length)));
          let _lhs217 = _2501_v57;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
          _2484_v44 = _rhs432;
          _2412_v0 = _rhs433;
          _lhs217[_lhs218] = _rhs434;
          let _2515_v69;
          _2515_v69 = _module.D14.create_DC28();
          let _2516_v70;
          _2516_v70 = _dafny.Seq.of(_2515_v69, _module.__default.fm51(_2484_v44, false, _2412_v0, globalState));
          let _rhs435 = (_2412_v0) || (_module.__default.fm5(globalState));
          let _rhs436 = _2516_v70;
          _2412_v0 = _rhs435;
          _2516_v70 = _rhs436;
        }
        let _2517_v71;
        _2517_v71 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2518_v72;
        _2518_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm36((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], _2517_v71, p0, globalState)).length),p2);
        let _2519_v73;
        let _init54 = ((_2520_v49) => function (_2521_i12) {
          return _2520_v49;
        })(_2488_v49);
        let _nw453 = Array((new BigNumber(16)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw453.length); _i0_54++) {
          _nw453[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2519_v73 = _nw453;
        let _2522_v74;
        _2522_v74 = _2519_v73;
        let _2523_v75;
        _2523_v75 = _dafny.MultiSet.fromElements(_2522_v74, _2522_v74, _2522_v74, _2522_v74);
        let _2524_v76;
        _2524_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2523_v75);
        let _2525_v77;
        _2525_v77 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2518_v72).length),_dafny.MultiSet.fromElements(_2522_v74))).Merge(_2524_v76),_2412_v0);
        _2525_v77 = (_2525_v77).update(_dafny.Map.Empty.slice().updateUnsafe(p0,_2523_v75), !(p2) || (_2412_v0));
        let _2526_v78;
        let _nw454 = new _module.C0();
        _nw454.__ctor((_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))], (_2501_v57)[_module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length))]);
        _2526_v78 = _nw454;
        let _2527_v79;
        _2527_v79 = _dafny.MultiSet.fromElements(p0);
        let _index463 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
        let _rhs437 = _dafny.Seq.Concat(_2484_v44, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_2528_v71) => function (_2529_i13) {
          return _2528_v71;
        })(_2517_v71)), _dafny.Seq.UnicodeFromString("fgyncltxh")));
        let _rhs438 = (p1).isLessThanOrEqualTo(new BigNumber((_2527_v79).cardinality()));
        let _lhs219 = _2501_v57;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2501_v57).length));
        _2484_v44 = _rhs437;
        _lhs219[_lhs220] = _rhs438;
      }
      let _2530_v80;
      _2530_v80 = _dafny.Seq.of(true, _module.__default.fm3(false, p0, p1, globalState), _2412_v0, p2);
      let _2531_v81;
      _2531_v81 = _dafny.MultiSet.fromElements(!(_2412_v0));
      let _2532_v82;
      let _nw455 = Array((new BigNumber(10)).toNumber());
      _nw455[(_dafny.ZERO).toNumber()] = _2412_v0;
      _nw455[(_dafny.ONE).toNumber()] = true;
      _nw455[(new BigNumber(2)).toNumber()] = !(!(_2412_v0));
      _nw455[(new BigNumber(3)).toNumber()] = _2412_v0;
      _nw455[(new BigNumber(4)).toNumber()] = (_2530_v80)[_module.__default.safeIndex(new BigNumber((_2531_v81).cardinality()), new BigNumber((_2530_v80).length))];
      _nw455[(new BigNumber(5)).toNumber()] = p2;
      _nw455[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2530_v80, _2530_v80);
      _nw455[(new BigNumber(7)).toNumber()] = _2412_v0;
      _nw455[(new BigNumber(8)).toNumber()] = p2;
      _nw455[(new BigNumber(9)).toNumber()] = !((p2) && (p2));
      _2532_v82 = _nw455;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2532_v82).length))) {
        let _2533_i14 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2533_i14)) && ((_2533_i14).isLessThan(new BigNumber((_2532_v82).length))))) {
          (_2532_v82)[(_2533_i14)] = (new BigNumber(900)).isLessThan(_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.of(p0)).length)));
        }
      }
      let _2534_v83;
      _2534_v83 = _module.D6.create_DC13(!(_2412_v0), p1);
      let _pat_let_tv56 = p0;
      let _pat_let_tv57 = p1;
      let _pat_let_tv58 = p1;
      r0 = function (_source35) {
        if (_source35.is_DC12) {
          let _2535___mcc_h0 = (_source35).cf22;
          let _2536___mcc_h1 = (_source35).cf23;
          let _2537_cf23 = _2536___mcc_h1;
          let _2538_cf22 = _2535___mcc_h0;
          return _pat_let_tv56;
        } else if (_source35.is_DC13) {
          let _2539___mcc_h2 = (_source35).cf24;
          let _2540___mcc_h3 = (_source35).cf25;
          let _2541_cf25 = _2540___mcc_h3;
          let _2542_cf24 = _2539___mcc_h2;
          return _pat_let_tv57;
        } else {
          let _2543___mcc_h4 = (_source35).cf21;
          let _2544_cf21 = _2543___mcc_h4;
          return _pat_let_tv58;
        }
      }(_2534_v83);
      let _2545_v84;
      let _nw456 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _2545_v84 = _nw456;
      let _index464 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2545_v84).length));
      (_2545_v84)[_index464] = _module.__default.safeModuloInt(new BigNumber(308), p1);
      let _index465 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2545_v84).length));
      (_2545_v84)[_index465] = p1;
      let _2546_v85;
      _2546_v85 = _module.D4.create_DC6(p0, _2545_v84, p0);
      _2546_v85 = _2546_v85;
      r0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_2531_v81).cardinality())), (p0).plus(new BigNumber((_dafny.Seq.of(p0)).length)));
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
