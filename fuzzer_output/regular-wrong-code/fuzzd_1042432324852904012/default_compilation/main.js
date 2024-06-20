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
    static fm2(p0, p1, p2, globalState) {
      let _source0 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-702))));
      let _0___mcc_h0 = _source0;
      let _1_cf0 = _0___mcc_h0;
      return true;
    };
    static fm6(p0, p1, globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("urqk"),new BigNumber(865))).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), function (_2_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length))) {
        return true;
      } else {
        return true;
      }
    };
    static fm7(p0, p1, globalState) {
      if (!(!((false) === (true)))) {
        return _dafny.Set.fromElements(true);
      } else {
        return (_dafny.Set.fromElements(false, false, false, !(true))).Intersect(_dafny.Set.fromElements(true, (_module.D16.create_DC45(false)).dtor_cf61, false));
      }
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)), _dafny.Seq.of(false, true));
    };
    static fm9(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(592),true)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(227))).cardinality()),true))));
    };
    static fm10(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_3_i0) {
        return new BigNumber(128);
      }), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(411), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("jsxnajks")).length))).length))).length)), new BigNumber(437)));
    };
    static fm11(p0, p1, globalState) {
      if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)))).IsDisjointFrom(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), (_module.D13.create_DC35(new _dafny.CodePoint('c'.codePointAt(0)))).dtor_cf45, new _dafny.CodePoint('w'.codePointAt(0))))) {
        return (_dafny.MultiSet.fromElements(new BigNumber(405), new BigNumber(994))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality())));
      } else {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qo")).length));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xrq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), function (_4_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("jhpmme"));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),false));
    };
    static fm17(p0, globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm18(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(false));
    };
    static fm19(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_module.D11.create_DC30(false, _dafny.MultiSet.FromArray(_dafny.Seq.of(true, !(false), true, true)))).dtor_cf40))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(true))));
    };
    static fm20(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC16(_dafny.Set.fromElements(false)), _module.D7.create_DC16(_dafny.Set.fromElements(true, false)), _module.D7.create_DC16(_dafny.Set.fromElements(!(false), true, !(true), false))), _dafny.Seq.of(_module.D7.create_DC16(_dafny.Set.fromElements(false))));
    };
    static fm21(globalState) {
      return _dafny.Seq.of(!(!(true)));
    };
    static fm22(p0, globalState) {
      let _source1 = !(!(true));
      let _5___mcc_h0 = _source1;
      let _6_cf1 = _5___mcc_h0;
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qaete")).length),new BigNumber((_dafny.Seq.of(new BigNumber(687))).length));
    };
    static fm23(globalState) {
      return (_dafny.MultiSet.fromElements(!(false), false, false)).Difference(_dafny.MultiSet.fromElements(false, !(!(false)), !(false)));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC11(new BigNumber(-410), !((_dafny.Set.fromElements(false)).IsProperSubsetOf(_dafny.Set.fromElements(true, false))), (!(false)) === (!(false)));
    };
    static fm25(globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(528), new BigNumber(-787))) {
          let _7_v0 = _compr_0;
          if (((new BigNumber(528)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(-787)))) {
            _coll0.add((_7_v0).plus(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(298)))).length)));
          }
        }
        return _coll0;
      }()).Intersect(_dafny.Set.fromElements(new BigNumber(949)));
    };
    static fm26(p0, globalState) {
      if (false) {
        return _dafny.Seq.UnicodeFromString("fix");
      } else {
        return (_module.D6.create_DC13(_dafny.Seq.UnicodeFromString("w"))).dtor_cf21;
      }
    };
    static fm27(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(634))).Merge((((_module.D4.create_DC10(_dafny.MultiSet.fromElements(false), true)).dtor_cf16) ? (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(945))) : (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-925)))));
    };
    static fm28(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(277))).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, true, !(true), false, false),false)).length)));
    };
    static fm29(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-614), (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(6), new BigNumber(99))) {
          let _8_v0 = _compr_1;
          if (((new BigNumber(6)).isLessThanOrEqualTo(_8_v0)) && ((_8_v0).isLessThan(new BigNumber(99)))) {
            _coll1.push([(_8_v0).plus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())),true]);
          }
        }
        return _coll1;
      }()).length)), new BigNumber((_dafny.Seq.UnicodeFromString("tnsrwmjof")).length))).length))),true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ilip")).length),true)).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(46), new BigNumber(717))) {
          let _9_v1 = _compr_2;
          if (((new BigNumber(46)).isLessThanOrEqualTo(_9_v1)) && ((_9_v1).isLessThan(new BigNumber(717)))) {
            _coll2.push([_module.__default.safeModuloInt(_9_v1, new BigNumber(132)),(false)]);
          }
        }
        return _coll2;
      }()));
    };
    static fm30(globalState) {
      let _source2 = _module.D12.create_DC33();
      if (_source2.is_DC33) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), function (_10_i0) {
          return _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)));
        }));
      } else if (_source2.is_DC32) {
        let _11___mcc_h0 = (_source2).cf43;
        let _12_cf43 = _11___mcc_h0;
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0))));
      } else {
        let _13___mcc_h1 = (_source2).cf44;
        let _14_cf44 = _13___mcc_h1;
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))));
      }
    };
    static fm31(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0))))).Difference(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))));
    };
    static fm32(globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)))).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0))))).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0))));
    };
    static fm33(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qe"), _dafny.Seq.UnicodeFromString("bmhgftehd"), _dafny.Seq.UnicodeFromString("yatjnok"), (_module.D6.create_DC13(_dafny.Seq.UnicodeFromString("ebufvm"))).dtor_cf21, _dafny.Seq.UnicodeFromString("ytlonbky")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("lw"))), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oojfql")));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_15_i0) {
  return _dafny.Set.fromElements(!(false));
}), !(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),new BigNumber((_dafny.Seq.of(true)).length))).contains(new _dafny.CodePoint('p'.codePointAt(0))));
    };
    static fm35(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(703)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(36)), function (_16_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(945),true);
      })).length)),(_dafny.Set.fromElements(!(false))).Difference((_module.D7.create_DC16(_dafny.Set.fromElements(true))).dtor_cf24));
    };
    static fm36(globalState) {
      return _module.D8.create_DC20();
    };
    static fm37(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(730), new BigNumber(70), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("thphighp")).length)))).cardinality()), new BigNumber(-40), new BigNumber(563))).Union(_dafny.MultiSet.fromElements(new BigNumber(210), new BigNumber(532), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), function (_17_i0) {
        return new BigNumber((_dafny.Seq.of(true, !(true))).length);
      })).length), new BigNumber(561)))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(163),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(290), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(320))).length))).length))).length))).length)));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-678), new BigNumber(-417))) {
          let _18_v0 = _compr_3;
          if (((new BigNumber(-678)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(-417)))) {
            _coll3.push([(_18_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(true))).length)),!((_module.D7.create_DC17(true)).dtor_cf25))).length)),false]);
          }
        }
        return _coll3;
      }(),new BigNumber(235))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(998),new BigNumber(447))).length))).length),!(true)),new BigNumber(951))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("g")).length),false),new BigNumber(709))));
    };
    static fm39(globalState) {
      return _module.D10.create_DC25(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length)));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(457))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(619))));
    };
    static fm41(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(true), _dafny.Seq.of((_module.D16.create_DC45(!(true))).dtor_cf61))).Intersect(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Seq.of(_dafny.Seq.of(true))).Elements) {
          let _19_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(true)), _19_v0)) {
            _coll4.add(_19_v0);
          }
        }
        return _coll4;
      }());
    };
    static fm42(p0, p1, globalState) {
      if (!((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tanewc")).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(410), new BigNumber(632))) {
          let _20_v0 = _compr_5;
          if (((new BigNumber(410)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(632)))) {
            _coll5.push([(_20_v0).plus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)))).length)),new BigNumber(78)]);
          }
        }
        return _coll5;
      }())).length)))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bvsgx")).length)),false)).length))))) {
        return _module.D3.create_DC6((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of (function () {
    let _coll7 = new _dafny.Set();
    for (const _compr_7 of (_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(452))).Elements) {
      let _21_v2 = _compr_7;
      if ((_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(452))).contains(_21_v2)) {
        _coll7.add(_module.__default.safeModuloInt(_21_v2, new BigNumber(-222)));
      }
    }
    return _coll7;
  }()).Elements) {
    let _22_v1 = _compr_6;
    if ((function () {
      let _coll8 = new _dafny.Set();
      for (const _compr_8 of (_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(452))).Elements) {
        let _23_v2 = _compr_8;
        if ((_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(452))).contains(_23_v2)) {
          _coll8.add(_module.__default.safeModuloInt(_23_v2, new BigNumber(-222)));
        }
      }
      return _coll8;
    }()).contains(_22_v1)) {
      _coll6.push([_module.__default.safeModuloInt(_22_v1, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).cardinality())),true]);
    }
  }
  return _coll6;
}()).length)));
      } else {
        return _module.D3.create_DC6(new BigNumber((function () {
  let _coll9 = new _dafny.Set();
  for (const _compr_9 of (_dafny.MultiSet.fromElements(new BigNumber(162))).Elements) {
    let _24_v3 = _compr_9;
    if ((_dafny.MultiSet.fromElements(new BigNumber(162))).contains(_24_v3)) {
      _coll9.add((_24_v3).multipliedBy((_dafny.ZERO).minus(new BigNumber(-187))));
    }
  }
  return _coll9;
}()).length));
      }
    };
    static Main(__noArgsParameter) {
      let _25_v0;
      _25_v0 = _dafny.Seq.UnicodeFromString("mrqsc");
      let _26_v1;
      _26_v1 = new BigNumber(-60);
      let _27_v2;
      let _nw0 = Array((new BigNumber(9)).toNumber()).fill(false);
      _27_v2 = _nw0;
      let _28_v3;
      _28_v3 = _dafny.Map.Empty.slice().updateUnsafe(_26_v1,_27_v2);
      let _29_v4;
      _29_v4 = _dafny.Set.fromElements(_26_v1);
      let _30_v5;
      _30_v5 = _dafny.Map.Empty.slice().updateUnsafe(_26_v1,_29_v4);
      let _31_v6;
      _31_v6 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((((_30_v5).contains(_26_v1)) ? ((_30_v5).get(_26_v1)) : (_29_v4))).length));
      let _32_v7;
      _32_v7 = true;
      let _33_v8;
      _33_v8 = _dafny.Map.Empty.slice().updateUnsafe(_26_v1,_26_v1);
      let _34_v9;
      _34_v9 = _dafny.Set.fromElements(_32_v7);
      let _35_v10;
      _35_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_34_v9).length),_32_v7);
      let _36_v12;
      _36_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_32_v7),_33_v8);
      let _37_v13;
      _37_v13 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_26_v1),_33_v8);
      let _38_v16;
      _38_v16 = _dafny.Seq.of(_26_v1, _26_v1);
      let _39_v17;
      _39_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_25_v0).length),_38_v16);
      let _40_v18;
      let _nw1 = Array((new BigNumber(19)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = (_33_v8).update(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_32_v7, _32_v7, true), _module.__default.safeIndex(_26_v1, new BigNumber((_dafny.Seq.of(_32_v7, _32_v7, true)).length)), _32_v7)).length), new BigNumber((_35_v10).length));
      _nw1[(_dafny.ONE).toNumber()] = _33_v8;
      _nw1[(new BigNumber(2)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(3)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(4)).toNumber()] = function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(419), new BigNumber(570))) {
          let _41_v11 = _compr_10;
          if (((new BigNumber(419)).isLessThanOrEqualTo(_41_v11)) && ((_41_v11).isLessThan(new BigNumber(570)))) {
            _coll10.push([_module.__default.safeDivisionInt(_41_v11, _26_v1),new BigNumber(898)]);
          }
        }
        return _coll10;
      }();
      _nw1[(new BigNumber(5)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(6)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(7)).toNumber()] = (((_36_v12).contains(_32_v7)) ? ((_36_v12).get(_32_v7)) : (_33_v8));
      _nw1[(new BigNumber(8)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(9)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(10)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(11)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(12)).toNumber()] = (((_37_v13).contains(_26_v1)) ? ((_37_v13).get(_26_v1)) : (_33_v8));
      _nw1[(new BigNumber(13)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(14)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(15)).toNumber()] = (((_37_v13).contains(_26_v1)) ? ((_37_v13).get(_26_v1)) : (_33_v8));
      _nw1[(new BigNumber(16)).toNumber()] = _33_v8;
      _nw1[(new BigNumber(17)).toNumber()] = function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_29_v4).Elements) {
          let _42_v14 = _compr_11;
          if ((_29_v4).contains(_42_v14)) {
            _coll11.push([_module.__default.safeDivisionInt(_42_v14, _26_v1),_26_v1]);
          }
        }
        return _coll11;
      }();
      _nw1[(new BigNumber(18)).toNumber()] = function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_39_v17).Keys.Elements) {
          let _43_v15 = _compr_12;
          if ((_39_v17).contains(_43_v15)) {
            _coll12.push([(_43_v15).minus(_26_v1),_26_v1]);
          }
        }
        return _coll12;
      }();
      _40_v18 = _nw1;
      let _44_v19;
      _44_v19 = _dafny.Seq.of(_25_v0);
      let _45_v20;
      _45_v20 = _dafny.Map.Empty.slice().updateUnsafe(_27_v2,new BigNumber((_44_v19).length));
      let _46_v21;
      _46_v21 = new _dafny.CodePoint('m'.codePointAt(0));
      let _47_globalState;
      let _nw2 = new _module.GlobalState();
      _nw2.__ctor(new BigNumber(687), _dafny.Seq.Concat(_25_v0, _25_v0), true, false, false, new BigNumber(351), (((_28_v3).contains((((_31_v6).contains(_32_v7)) ? ((_31_v6).get(_32_v7)) : (_26_v1)))) ? ((_28_v3).get((((_31_v6).contains(_32_v7)) ? ((_31_v6).get(_32_v7)) : (_26_v1)))) : (_27_v2)), false, _40_v18, new BigNumber(825), false, new BigNumber(-252), false, new BigNumber(318), true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), function (_48_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), (_31_v6).Merge(_31_v6), new BigNumber(124), false, _dafny.Seq.Concat(_dafny.Seq.update(_25_v0, _module.__default.safeIndex(new BigNumber((_45_v20).length), new BigNumber((_25_v0).length)), _46_v21), _25_v0));
      _47_globalState = _nw2;
      let _49_v22;
      let _nw3 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _49_v22 = _nw3;
      let _50_v23;
      let _nw4 = new _module.C1();
      _nw4.__ctor(_26_v1, _49_v22, new BigNumber(542), _32_v7);
      _50_v23 = _nw4;
      _34_v9 = _34_v9;
      let _51_v24;
      _51_v24 = _dafny.Map.Empty.slice().updateUnsafe(_32_v7,_34_v9);
      _34_v9 = (((((_51_v24).contains(_32_v7)) ? ((_51_v24).get(_32_v7)) : (_34_v9))).Intersect(_34_v9)).Union((_34_v9).Difference(_34_v9));
      let _52_v25;
      let _53_v26;
      let _out0;
      let _out1;
      let _outcollector0 = (_50_v23).m0(_module.__default.safeModuloInt(_26_v1, (_dafny.ZERO).minus(_26_v1)), _32_v7, _47_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _52_v25 = _out0;
      _53_v26 = _out1;
      (_47_globalState).f3 = _32_v7;
      let _54_v27;
      let _init0 = ((_55_v10, _56_v23) => function (_57_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(_55_v10,_56_v23.f27);
      })(_35_v10, _50_v23);
      let _nw5 = Array((new BigNumber(26)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
        _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _54_v27 = _nw5;
      let _58_v28;
      _58_v28 = _dafny.Map.Empty.slice().updateUnsafe(_32_v7,_dafny.Map.Empty.slice().updateUnsafe(_50_v23.f27,_53_v26));
      let _59_v29;
      _59_v29 = _dafny.Map.Empty.slice().updateUnsafe((((_58_v28).contains(_32_v7)) ? ((_58_v28).get(_32_v7)) : (_35_v10)),_50_v23.f27);
      let _index0 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_54_v27).length));
      (_54_v27)[_index0] = (_59_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(_35_v10,_50_v23.f27));
      let _60_v30;
      let _nw6 = new _module.C5();
      _nw6.__ctor(_53_v26, _26_v1, _53_v26);
      _60_v30 = _nw6;
      let _61_v31;
      _61_v31 = _dafny.Map.Empty.slice().updateUnsafe(_53_v26,_60_v30);
      let _62_v34;
      _62_v34 = _dafny.Map.Empty.slice().updateUnsafe(_53_v26,_60_v30.f23);
      let _index1 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_54_v27).length));
      (_54_v27)[_index1] = ((!(_61_v31).contains(_32_v7)) ? ((function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of ((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of ((_59_v29).update(_35_v10, new BigNumber((_62_v34).length))).Keys.Elements) {
            let _63_v33 = _compr_14;
            if (((_59_v29).update(_35_v10, new BigNumber((_62_v34).length))).contains(_63_v33)) {
              _coll14.push([_63_v33,false]);
            }
          }
          return _coll14;
        }()).update(_35_v10, _60_v30.f23)).Keys.Elements) {
          let _64_v32 = _compr_13;
          if (((function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of ((_59_v29).update(_35_v10, new BigNumber((_62_v34).length))).Keys.Elements) {
              let _63_v33 = _compr_15;
              if (((_59_v29).update(_35_v10, new BigNumber((_62_v34).length))).contains(_63_v33)) {
                _coll15.push([_63_v33,false]);
              }
            }
            return _coll15;
          }()).update(_35_v10, _60_v30.f23)).contains(_64_v32)) {
            _coll13.push([_64_v32,_50_v23.f27]);
          }
        }
        return _coll13;
      }()).Merge(_59_v29)) : (_dafny.Map.Empty.slice().updateUnsafe(_35_v10,_50_v23.f27)));
      let _hi0 = new BigNumber((_dafny.Seq.of(_46_v21)).length);
      for (let _65_i2 = (new BigNumber(676)).minus(_50_v23.f27); _65_i2.isLessThan(_hi0); _65_i2 = _65_i2.plus(_dafny.ONE)) {
        let _66_v35;
        _66_v35 = _49_v22;
        let _67_v36;
        _67_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_60_v30.f23) ? (_module.__default.fm17(_53_v26, _47_globalState)) : (_46_v21)),_66_v35);
        let _68_v37;
        let _init1 = ((_69_v23) => function (_70_i3) {
          return _dafny.MultiSet.fromElements(_69_v23.f27, new BigNumber(195));
        })(_50_v23);
        let _nw7 = Array((new BigNumber(20)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
          _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _68_v37 = _nw7;
        let _71_v38;
        _71_v38 = _dafny.MultiSet.fromElements(_26_v1);
        let _index2 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_68_v37).length));
        (_68_v37)[_index2] = (_71_v38).Difference((_71_v38).update(_50_v23.f27, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length))));
        let _72_v39;
        let _nw8 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _72_v39 = _nw8;
        let _73_v40;
        _73_v40 = _module.D6.create_DC13(_dafny.Seq.UnicodeFromString("ip"));
        let _74_v41;
        _74_v41 = _dafny.Seq.of(_62_v34);
        let _index3 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_68_v37).length));
        let _rhs0 = !((_50_v23).fm0((_33_v8).Merge(_33_v8), (_73_v40).dtor_cf21, _dafny.Seq.Concat(_74_v41, _dafny.Seq.update(_dafny.Seq.update(_74_v41, _module.__default.safeIndex(_50_v23.f27, new BigNumber((_74_v41).length)), _62_v34), _module.__default.safeIndex(_50_v23.f27, new BigNumber((_dafny.Seq.update(_74_v41, _module.__default.safeIndex(_50_v23.f27, new BigNumber((_74_v41).length)), _62_v34)).length)), _62_v34)), _47_globalState));
        let _rhs1 = _67_v36;
        let _rhs2 = ((_71_v38).update(new BigNumber(14), _module.__default.abs(_50_v23.f27))).Union((_71_v38).Union(_71_v38));
        let _rhs3 = _53_v26;
        let _rhs4 = _72_v39;
        let _lhs0 = _47_globalState;
        let _lhs1 = _68_v37;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_68_v37).length));
        let _lhs3 = _47_globalState;
        _lhs0.f14 = _rhs0;
        _67_v36 = _rhs1;
        _lhs1[_lhs2] = _rhs2;
        _lhs3.f14 = _rhs3;
        _72_v39 = _rhs4;
        (_50_v23).f27 = _65_i2;
        _49_v22 = _49_v22;
        (_47_globalState).f3 = !(_50_v23.f27).isEqualTo(new BigNumber(214));
      }
      (_47_globalState).f17 = _50_v23.f27;
      let _75_v42;
      _75_v42 = _module.D13.create_DC36(_50_v23.f27, (_dafny.Set.fromElements(_50_v23.f27, _50_v23.f27)).IsProperSubsetOf(_29_v4), _32_v7);
      let _source3 = _75_v42;
      if (_source3.is_DC36) {
        let _76___mcc_h0 = (_source3).cf46;
        let _77___mcc_h1 = (_source3).cf47;
        let _78___mcc_h2 = (_source3).cf48;
        let _79_cf48 = _78___mcc_h2;
        let _80_cf47 = _77___mcc_h1;
        let _81_cf46 = _76___mcc_h0;
        (_47_globalState).f4 = _60_v30.f23;
        let _82_v43;
        _82_v43 = _32_v7;
        let _83_v44;
        let _nw9 = new _module.C2();
        _nw9.__ctor(_module.__default.safeDivisionInt(_50_v23.f27, _50_v23.f27), (_82_v43));
        _83_v44 = _nw9;
        let _84_v45;
        let _nw10 = new _module.C4();
        _nw10.__ctor(!(_79_cf48), _50_v23.f27, _32_v7);
        _84_v45 = _nw10;
        let _85_v46;
        _85_v46 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_25_v0).length),_26_v1)).length), _26_v1, _81_cf46);
        let _86_v47;
        _86_v47 = _dafny.Map.Empty.slice().updateUnsafe(_84_v45,_85_v46);
        let _87_v48;
        _87_v48 = _dafny.Seq.of(_53_v26);
        let _88_v49;
        _88_v49 = _dafny.Map.Empty.slice().updateUnsafe(((((_86_v47).contains(_84_v45)) ? ((_86_v47).get(_84_v45)) : (_dafny.MultiSet.fromElements(_26_v1, (_dafny.ZERO).minus(_26_v1), _50_v23.f27, _26_v1, new BigNumber((_87_v48).length))))).IsDisjointFrom(_85_v46),_25_v0);
        _81_cf46 = new BigNumber((_88_v49).length);
        (_60_v30).f23 = (((_35_v10).contains(_50_v23.f27)) ? ((_35_v10).get(_50_v23.f27)) : ((_60_v30.f23) === (_32_v7)));
      } else if (_source3.is_DC37) {
        let _89___mcc_h3 = (_source3).cf49;
        let _90___mcc_h4 = (_source3).cf50;
        let _91_cf50 = _90___mcc_h4;
        let _92_cf49 = _89___mcc_h3;
        (_47_globalState).f18 = _92_cf49;
        let _93_v50;
        _93_v50 = _dafny.Map.Empty.slice().updateUnsafe((_29_v4).Intersect(_29_v4),_dafny.Set.fromElements(!(_60_v30.f23)));
        _91_cf50 = new BigNumber((_93_v50).length);
        _46_v21 = new _dafny.CodePoint('t'.codePointAt(0));
        let _94_v51;
        _94_v51 = _dafny.Seq.of(_32_v7);
        if ((new BigNumber((_94_v51).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_module.__default.safeModuloInt(_26_v1, _91_cf50)))) {
          let _95_v52;
          _95_v52 = _dafny.Seq.of(_27_v2, _27_v2);
          _95_v52 = _dafny.Seq.update(_dafny.Seq.of(_27_v2, _27_v2, _27_v2), _module.__default.safeIndex(_module.__default.safeDivisionInt(_26_v1, _26_v1), new BigNumber((_dafny.Seq.of(_27_v2, _27_v2, _27_v2)).length)), _27_v2);
          _33_v8 = (_33_v8).update(((_60_v30.f23) ? (_50_v23.f27) : ((_dafny.ZERO).minus(new BigNumber((_33_v8).length)))), _50_v23.f27);
          let _96_v53;
          let _nw11 = new _module.C5();
          _nw11.__ctor(_60_v30.f23, _50_v23.f27, _32_v7);
          _96_v53 = _nw11;
          _96_v53 = _96_v53;
          _46_v21 = _46_v21;
          let _97_v54;
          _97_v54 = _dafny.Map.Empty.slice().updateUnsafe(_46_v21,_29_v4);
          _97_v54 = _97_v54;
        } else {
          (_47_globalState).f5 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), _46_v21), _25_v0)).length);
          _31_v6 = (_31_v6).update(_60_v30.f23, _91_cf50);
          let _98_v55;
          let _nw12 = Array((new BigNumber(6)).toNumber()).fill([]);
          _98_v55 = _nw12;
          let _99_v56;
          _99_v56 = _module.D18.create_DC50(_98_v55);
          let _100_v57;
          let _nw13 = Array((new BigNumber(17)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _98_v55;
          _nw13[(_dafny.ONE).toNumber()] = _98_v55;
          _nw13[(new BigNumber(2)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(3)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(4)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(5)).toNumber()] = (_99_v56).dtor_cf69;
          _nw13[(new BigNumber(6)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(7)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(8)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(9)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(10)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(11)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(12)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(13)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(14)).toNumber()] = ((_53_v26) ? ((_module.D18.create_DC50(_98_v55)).dtor_cf69) : (_98_v55));
          _nw13[(new BigNumber(15)).toNumber()] = _98_v55;
          _nw13[(new BigNumber(16)).toNumber()] = _98_v55;
          _100_v57 = _nw13;
          let _index4 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_100_v57).length));
          (_100_v57)[_index4] = _98_v55;
          let _index5 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_100_v57).length));
          (_100_v57)[_index5] = _98_v55;
          let _101_v58;
          _101_v58 = _dafny.Seq.of(_34_v9, _34_v9, _34_v9);
          let _102_v59;
          _102_v59 = _module.D3.create_DC7(_101_v58, _60_v30.f23);
          let _103_v60;
          let _nw14 = new _module.C4();
          _nw14.__ctor(true, _50_v23.f27, _32_v7);
          _103_v60 = _nw14;
          let _104_v61;
          _104_v61 = _dafny.Map.Empty.slice().updateUnsafe(_103_v60,_50_v23.f27);
          let _105_v62;
          _105_v62 = _module.D19.create_DC54(_104_v61);
          let _106_v63;
          let _nw15 = new _module.C3();
          _nw15.__ctor(_50_v23.f27, (_102_v59).dtor_cf12, (_dafny.ZERO).minus(_50_v23.f27), ((_105_v62).dtor_cf72).contains(_103_v60));
          _106_v63 = _nw15;
          (_60_v30).f23 = ((_34_v9).Union(_34_v9)).IsDisjointFrom((_34_v9).Union(_dafny.Set.fromElements((_103_v60).f24)));
        }
      } else {
        let _107___mcc_h5 = (_source3).cf45;
        let _108_cf45 = _107___mcc_h5;
        let _109_v64;
        let _nw16 = new _module.C2();
        _nw16.__ctor(_50_v23.f27, _32_v7);
        _109_v64 = _nw16;
        _109_v64 = _109_v64;
        let _110_v65;
        _110_v65 = _49_v22;
        let _111_v66;
        _111_v66 = _dafny.Map.Empty.slice().updateUnsafe(_60_v30.f23,(_50_v23).f28);
        let _112_v67;
        let _nw17 = Array((new BigNumber(16)).toNumber());
        _nw17[(_dafny.ZERO).toNumber()] = _49_v22;
        _nw17[(_dafny.ONE).toNumber()] = _110_v65;
        _nw17[(new BigNumber(2)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(3)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(4)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(5)).toNumber()] = ((false) ? ((((_111_v66).contains(_60_v30.f23)) ? ((_111_v66).get(_60_v30.f23)) : ((_50_v23).f28))) : ((_50_v23).f28));
        _nw17[(new BigNumber(6)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(7)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(8)).toNumber()] = (_50_v23).f28;
        _nw17[(new BigNumber(9)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(10)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(11)).toNumber()] = (_50_v23).f28;
        _nw17[(new BigNumber(12)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(13)).toNumber()] = (_50_v23).f28;
        _nw17[(new BigNumber(14)).toNumber()] = _110_v65;
        _nw17[(new BigNumber(15)).toNumber()] = _110_v65;
        _112_v67 = _nw17;
        let _index6 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_112_v67).length));
        (_112_v67)[_index6] = (_50_v23).f28;
        let _index7 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_112_v67).length));
        let _rhs5 = _110_v65;
        let _rhs6 = _dafny.Seq.Concat(_38_v16, _38_v16);
        let _rhs7 = _50_v23.f27;
        let _lhs4 = _112_v67;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_112_v67).length));
        let _lhs6 = _47_globalState;
        _lhs4[_lhs5] = _rhs5;
        _38_v16 = _rhs6;
        _lhs6.f17 = _rhs7;
        (_47_globalState).f4 = (new BigNumber((_62_v34).length)).isLessThan(_26_v1);
        let _113_v68;
        let _nw18 = new _module.C5();
        _nw18.__ctor(!(_60_v30.f23), new BigNumber(651), _32_v7);
        _113_v68 = _nw18;
      }
      (_47_globalState).f17 = _26_v1;
      (_47_globalState).f5 = _26_v1;
      let _114_v69;
      _114_v69 = _module.D17.create_DC47(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm41((_dafny.ZERO).minus(_26_v1), (((_31_v6).contains(_32_v7)) ? ((_31_v6).get(_32_v7)) : ((_50_v23).fm1(_26_v1, _26_v1, _60_v30.f23, _47_globalState))), _47_globalState)).length), (_dafny.ZERO).minus(new BigNumber((_25_v0).length))), _60_v30.f23, (_26_v1).plus(new BigNumber((_33_v8).length)));
      _114_v69 = _114_v69;
      let _115_v70;
      let _nw19 = new _module.C0();
      _nw19.__ctor();
      _115_v70 = _nw19;
      let _116_v71;
      let _nw20 = Array((new BigNumber(12)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = _115_v70;
      _nw20[(_dafny.ONE).toNumber()] = _115_v70;
      _nw20[(new BigNumber(2)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(3)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(4)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(5)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(6)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(7)).toNumber()] = ((_60_v30.f23) ? (_115_v70) : (_115_v70));
      _nw20[(new BigNumber(8)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(9)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(10)).toNumber()] = _115_v70;
      _nw20[(new BigNumber(11)).toNumber()] = _115_v70;
      _116_v71 = _nw20;
      let _117_v72;
      _117_v72 = _dafny.Seq.of(_115_v70);
      let _index8 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_116_v71).length));
      (_116_v71)[_index8] = (_117_v72)[_module.__default.safeIndex(_50_v23.f27, new BigNumber((_117_v72).length))];
      let _index9 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_116_v71).length));
      (_116_v71)[_index9] = _115_v70;
      let _118_v73;
      let _init2 = ((_119_v9) => function (_120_i5) {
        return _119_v9;
      })(_34_v9);
      let _nw21 = Array((new BigNumber(9)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw21.length); _i0_2++) {
        _nw21[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _118_v73 = _nw21;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_118_v73).length))) {
        let _121_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_121_i4)) && ((_121_i4).isLessThan(new BigNumber((_118_v73).length))))) {
          (_118_v73)[(_121_i4)] = (_module.D7.create_DC16(_34_v9)).dtor_cf24;
        }
      }
      _46_v21 = _module.__default.fm17(_53_v26, _47_globalState);
      let _122_v74;
      _122_v74 = _module.D12.create_DC33();
      let _123_i6;
      _123_i6 = _dafny.ZERO;
      L0: {
        let _pat_let_tv0 = _50_v23;
        let _pat_let_tv1 = _50_v23;
        let _pat_let_tv2 = _60_v30;
        let _pat_let_tv3 = _32_v7;
        while (function (_source4) {
          if (_source4.is_DC33) {
            return !(_pat_let_tv0.f27).isEqualTo(_pat_let_tv1.f27);
          } else if (_source4.is_DC32) {
            let _133___mcc_h6 = (_source4).cf43;
            let _134_cf43 = _133___mcc_h6;
            return _pat_let_tv2.f23;
          } else {
            let _135___mcc_h7 = (_source4).cf44;
            let _136_cf44 = _135___mcc_h7;
            return _pat_let_tv3;
          }
        }(_122_v74)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_123_i6)) {
              break L0;
            }
            _123_i6 = (_123_i6).plus(_dafny.ONE);
            let _index10 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_27_v2).length));
            (_27_v2)[_index10] = _60_v30.f23;
            let _index11 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_27_v2).length));
            (_27_v2)[_index11] = _32_v7;
            let _124_v75;
            let _nw22 = Array((new BigNumber(25)).toNumber());
            _124_v75 = _nw22;
            let _125_v76;
            _125_v76 = _dafny.Seq.of(_53_v26);
            let _126_v77;
            let _nw23 = new _module.C6();
            _nw23.__ctor(_25_v0, new BigNumber((_31_v6).length), (_125_v76)[_module.__default.safeIndex(_26_v1, new BigNumber((_125_v76).length))]);
            _126_v77 = _nw23;
            let _index12 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_124_v75).length));
            (_124_v75)[_index12] = _126_v77;
            let _index13 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_27_v2).length));
            (_27_v2)[_index13] = _53_v26;
            let _127_v78;
            _127_v78 = _dafny.MultiSet.fromElements((_126_v77).f22);
            let _128_v79;
            _128_v79 = _module.D16.create_DC43(_26_v1, _127_v78);
            let _index14 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_27_v2).length));
            let _index15 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_27_v2).length));
            let _index16 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_124_v75).length));
            let _index17 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_27_v2).length));
            let _rhs8 = _60_v30.f23;
            let _rhs9 = _53_v26;
            let _rhs10 = ((_32_v7) ? (((_60_v30).fm1((_50_v23).fm1((_128_v79).dtor_cf55, _26_v1, _53_v26, _47_globalState), new BigNumber((_module.__default.fm21(_47_globalState)).length), _60_v30.f23, _47_globalState)).minus(_26_v1)) : ((new BigNumber(((_126_v77).f22).length)).multipliedBy(_50_v23.f27)));
            let _rhs11 = ((_53_v26) ? (_126_v77) : (_126_v77));
            let _rhs12 = false;
            let _lhs7 = _27_v2;
            let _lhs8 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_27_v2).length));
            let _lhs9 = _27_v2;
            let _lhs10 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_27_v2).length));
            let _lhs11 = _47_globalState;
            let _lhs12 = _124_v75;
            let _lhs13 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_124_v75).length));
            let _lhs14 = _27_v2;
            let _lhs15 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_27_v2).length));
            _lhs7[_lhs8] = _rhs8;
            _lhs9[_lhs10] = _rhs9;
            _lhs11.f5 = _rhs10;
            _lhs12[_lhs13] = _rhs11;
            _lhs14[_lhs15] = _rhs12;
            (_47_globalState).f5 = (_50_v23.f27).plus(_26_v1);
            let _129_v80;
            _129_v80 = _module.D13.create_DC35(_46_v21);
            _129_v80 = (((_module.D17.create_DC48(_60_v30.f23, (_52_v25)[_module.__default.safeIndex((_module.__default.fm42(_53_v26, (_27_v2)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_27_v2).length))], _47_globalState)).dtor_cf10, new BigNumber((_52_v25).length))])).dtor_cf66) ? (_129_v80) : (_module.D13.create_DC35(_46_v21)));
            let _130_v81;
            let _init3 = ((_131_v2) => function (_132_i7) {
              return (_131_v2)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_131_v2).length))];
            })(_27_v2);
            let _nw24 = Array((new BigNumber(8)).toNumber());
            for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw24.length); _i0_3++) {
              _nw24[_i0_3] = _init3(new BigNumber(_i0_3));
            }
            _130_v81 = _nw24;
            _28_v3 = (_28_v3).update(_50_v23.f27, _130_v81);
          }
        }
      }
      process.stdout.write((_25_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_26_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_28_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_29_v4).equals(_dafny.Set.fromElements(new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_30_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),_dafny.Set.fromElements(new BigNumber(-60))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_31_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_32_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_33_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_34_v9).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_35_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_36_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_37_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(60),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_38_v16, _dafny.Seq.of(new BigNumber(-60), new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_39_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(5),_dafny.Seq.of(new BigNumber(-60), new BigNumber(-60))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)).updateUnsafe(new BigNumber(3),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-6),new BigNumber(898)).updateUnsafe(new BigNumber(-7),new BigNumber(898)).updateUnsafe(new BigNumber(-8),new BigNumber(898)).updateUnsafe(new BigNumber(-9),new BigNumber(898)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_40_v18)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_44_v19, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("mrqsc")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_45_v20).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_46_v21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_47_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)).updateUnsafe(new BigNumber(3),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-6),new BigNumber(898)).updateUnsafe(new BigNumber(-7),new BigNumber(898)).updateUnsafe(new BigNumber(-8),new BigNumber(898)).updateUnsafe(new BigNumber(-9),new BigNumber(898)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState.f8)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_47_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_47_globalState.f15).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_47_globalState).f16).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_47_globalState).f19).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_50_v23.f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v24).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_52_v25).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_53_v26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_54_v27)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_58_v28).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_v29).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-60),true),new BigNumber(-60)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_60_v30.f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_61_v31).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_62_v34).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_v42).dtor_cf46));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_v42).dtor_cf47));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_v42).dtor_cf48));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v69).dtor_cf63));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v69).dtor_cf64));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v69).dtor_cf65));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_117_v72).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(5)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(6)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(7)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_v73)[new BigNumber(8)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_123_i6));
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
      return _dafny.Seq.of();
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
        return other.$tag === 0 && this.cf1 === other.cf1;
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC3(cf3, cf4, cf5, cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D2(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D2(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC3(_dafny.ZERO, false, false, _dafny.ZERO, false);
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
    static create_DC6(cf10) {
      let $dt = new D3(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC7(cf11, cf12) {
      let $dt = new D3(1);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D3(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D3(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.ZERO);
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
    static create_DC10(cf15, cf16) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC11(cf17, cf18, cf19) {
      let $dt = new D4(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D4(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf14 === other.cf14;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(_dafny.MultiSet.Empty, false);
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
    static create_DC12(cf20) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf22) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC15(cf23) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC13(cf21) {
      let $dt = new D6(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC13" + "(" + this.cf21.toVerbatimString(true) + ")";
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
      return _module.D6.create_DC14(_dafny.ZERO);
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
    static create_DC17(cf25) {
      let $dt = new D7(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC18(cf26, cf27) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC16(cf24) {
      let $dt = new D7(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf24) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(false);
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
    static create_DC20() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC21(cf29) {
      let $dt = new D8(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D8(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf28) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20();
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
    static create_DC23(cf31, cf32) {
      let $dt = new D9(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC22(cf30) {
      let $dt = new D9(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC24(cf33) {
      let $dt = new D9(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.ZERO, false);
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
    static create_DC26(cf35) {
      let $dt = new D10(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC27(cf36, cf37, cf38) {
      let $dt = new D10(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf34) {
      let $dt = new D10(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(_dafny.ZERO);
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
    static create_DC29() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC30(cf40, cf41) {
      let $dt = new D11(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC28(cf39) {
      let $dt = new D11(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC31(cf42) {
      let $dt = new D11(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC29";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf42) + ")";
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
        return other.$tag === 1 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf39 === other.cf39;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC29();
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
    static create_DC33() {
      let $dt = new D12(0);
      return $dt;
    }
    static create_DC32(cf43) {
      let $dt = new D12(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC34(cf44) {
      let $dt = new D12(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf44) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33();
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
    static create_DC36(cf46, cf47, cf48) {
      let $dt = new D13(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC37(cf49, cf50) {
      let $dt = new D13(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC35(cf45) {
      let $dt = new D13(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC36(_dafny.ZERO, false, false);
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
    static create_DC39() {
      let $dt = new D14(0);
      return $dt;
    }
    static create_DC38(cf51) {
      let $dt = new D14(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC40(cf52) {
      let $dt = new D14(2);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC39";
      } else if (this.$tag === 1) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf52) + ")";
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
        return other.$tag === 1 && this.cf51 === other.cf51;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC39();
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
    static create_DC41(cf53) {
      let $dt = new D15(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC41" + "(" + _dafny.toString(this.cf53) + ")";
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC43(cf55, cf56) {
      let $dt = new D16(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC44(cf57, cf58, cf59, cf60) {
      let $dt = new D16(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC45(cf61) {
      let $dt = new D16(2);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC42(cf54) {
      let $dt = new D16(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get is_DC42() { return this.$tag === 3; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf61 === other.cf61;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC43(_dafny.ZERO, _dafny.MultiSet.Empty);
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
    static create_DC47(cf63, cf64, cf65) {
      let $dt = new D17(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC48(cf66, cf67) {
      let $dt = new D17(1);
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC49(cf68) {
      let $dt = new D17(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC46(cf62) {
      let $dt = new D17(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get is_DC46() { return this.$tag === 3; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64 && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf68 === other.cf68;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC47(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC51() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC52(cf70) {
      let $dt = new D18(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC50(cf69) {
      let $dt = new D18(2);
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC53(cf71) {
      let $dt = new D18(3);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get is_DC53() { return this.$tag === 3; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC51";
      } else if (this.$tag === 1) {
        return "D18.DC52" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC53" + "(" + _dafny.toString(this.cf71) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf69 === other.cf69;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC51();
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
    static create_DC55(cf73, cf74) {
      let $dt = new D19(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC54(cf72) {
      let $dt = new D19(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC55" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC54" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC55(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC56(cf75) {
      let $dt = new D20(0);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC56" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75);
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
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf76) {
      let $dt = new D21(0);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76);
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
          return D21.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = false;
      this.f4 = false;
      this.f5 = _dafny.ZERO;
      this.f8 = [];
      this.f14 = false;
      this.f15 = _dafny.Seq.UnicodeFromString("");
      this.f17 = _dafny.ZERO;
      this.f18 = false;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f2 = false;
      this._f6 = [];
      this._f7 = false;
      this._f9 = _dafny.ZERO;
      this._f10 = false;
      this._f11 = _dafny.ZERO;
      this._f12 = false;
      this._f13 = _dafny.ZERO;
      this._f16 = _dafny.Map.Empty;
      this._f19 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this).f18 = f18;
      (_this)._f19 = f19;
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
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f19() {
      let _this = this;
      return _this._f19;
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
    fm3(p0, p1, globalState) {
      let _this = this;
      return !(false);
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return !(!((true)));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
      this.f27 = _dafny.ZERO;
      this._f28 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f27, f28, f20, f21) {
      let _this = this;
      (_this).f27 = f27;
      (_this)._f28 = f28;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f21;
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f27;
    };
    fm16(p0, p1, globalState) {
      let _this = this;
      return ((((_this).f21) ? (_dafny.Set.fromElements((_this).f21, (_this).f21, (_this).f21, (_this).f21, (_this).f21)) : (_dafny.Set.fromElements((_this).f21, (_this).f21, (_this).f21)))).IsProperSubsetOf((_dafny.Set.fromElements((_this).f21)).Difference(_dafny.Set.fromElements((_this).f21)));
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _137_v0;
      _137_v0 = _dafny.MultiSet.fromElements(_this.f27);
      _137_v0 = _dafny.MultiSet.fromElements(_this.f27);
      let _138_i0;
      _138_i0 = _dafny.ZERO;
      L1: {
        while (!(p1) || ((_this.f27).isLessThanOrEqualTo((_this).f20))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_138_i0)) {
              break L1;
            }
            _138_i0 = (_138_i0).plus(_dafny.ONE);
            (globalState).f18 = p1;
            let _139_v1;
            _139_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28);
            let _140_v2;
            _140_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,!(p1));
            let _141_v3;
            _141_v3 = _dafny.MultiSet.fromElements((((_140_v2).contains(p1)) ? ((_140_v2).get(p1)) : ((_this).f21)));
            _139_v1 = (_139_v1).update((new BigNumber((_141_v3).cardinality())).multipliedBy((_this).f20), (_this).f28);
            (globalState).f5 = (_this.f27).plus(p0);
            (globalState).f5 = (((_this).f21) ? ((_this).f20) : (p0));
          }
        }
      }
      let _142_v4;
      _142_v4 = _dafny.MultiSet.fromElements((_this).f21);
      let _143_v5;
      _143_v5 = _dafny.Seq.of(p0, new BigNumber((_137_v0).cardinality()));
      let _144_v6;
      _144_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,_143_v5);
      if ((_this).fm16(_142_v4, (((_144_v6).contains(new BigNumber(824))) ? ((_144_v6).get(new BigNumber(824))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-961)), function (_145_i1) {
        return _this.f27;
      }))), globalState)) {
        let _146_v7;
        _146_v7 = _module.D4.create_DC11((_this).f20, (_this).f21, (_this).f21);
        let _source5 = _146_v7;
        if (_source5.is_DC10) {
          let _147___mcc_h0 = (_source5).cf15;
          let _148___mcc_h1 = (_source5).cf16;
          let _149_cf16 = _148___mcc_h1;
          let _150_cf15 = _147___mcc_h0;
          let _151_v8;
          let _nw25 = Array((new BigNumber(19)).toNumber()).fill([]);
          _151_v8 = _nw25;
          let _index18 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_151_v8).length));
          (_151_v8)[_index18] = (_this).f28;
          let _index19 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_151_v8).length));
          (_151_v8)[_index19] = (_this).f28;
          let _152_v9;
          let _nw26 = new _module.C0();
          _nw26.__ctor();
          _152_v9 = _nw26;
          (globalState).f3 = p1;
          (globalState).f14 = ((_this).f20).isLessThanOrEqualTo(_this.f27);
        } else if (_source5.is_DC11) {
          let _153___mcc_h2 = (_source5).cf17;
          let _154___mcc_h3 = (_source5).cf18;
          let _155___mcc_h4 = (_source5).cf19;
          let _156_cf19 = _155___mcc_h4;
          let _157_cf18 = _154___mcc_h3;
          let _158_cf17 = _153___mcc_h2;
          (globalState).f14 = p1;
          let _159_v10;
          let _nw27 = new _module.C0();
          _nw27.__ctor();
          _159_v10 = _nw27;
          (globalState).f18 = !(_156_cf19);
          (globalState).f18 = (_this).f21;
        } else {
          let _160___mcc_h5 = (_source5).cf14;
          let _161_cf14 = _160___mcc_h5;
          let _162_v11;
          _162_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f27,(_this).f28);
          let _163_v12;
          _163_v12 = _dafny.Seq.of((((_162_v11).contains(new BigNumber(35))) ? ((_162_v11).get(new BigNumber(35))) : ((_this).f28)));
          _163_v12 = _163_v12;
          let _index20 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_161_cf14).length));
          (_161_cf14)[_index20] = p1;
          let _index21 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_161_cf14).length));
          (_161_cf14)[_index21] = !((_this).f21) || (p1);
          let _164_v13;
          let _nw28 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _164_v13 = _nw28;
          let _165_v14;
          _165_v14 = _dafny.Seq.UnicodeFromString("mgaj");
          let _index22 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_164_v13).length));
          (_164_v13)[_index22] = _165_v14;
          let _166_v15;
          _166_v15 = new _dafny.CodePoint('k'.codePointAt(0));
          let _index23 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_164_v13).length));
          (_164_v13)[_index23] = ((true) ? (_165_v14) : (_dafny.Seq.update(_165_v14, _module.__default.safeIndex(_this.f27, new BigNumber((_165_v14).length)), _166_v15)));
          let _167_v16;
          _167_v16 = _module.D6.create_DC13(_165_v14);
          _165_v14 = (_167_v16).dtor_cf21;
        }
        let _168_v17;
        let _nw29 = Array((new BigNumber(8)).toNumber()).fill([]);
        _168_v17 = _nw29;
        _168_v17 = _168_v17;
        (globalState).f3 = p1;
        (globalState).f4 = !(false) || (p1);
        let _169_v18;
        _169_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(p0).multipliedBy(_this.f27));
        let _170_v19;
        _170_v19 = _dafny.Seq.UnicodeFromString("e");
        _169_v18 = (_169_v18).update((_this).f28, (_module.D2.create_DC3((((_137_v0).contains(p0)) ? ((_137_v0).get(p0)) : ((_dafny.ZERO).minus(new BigNumber((_170_v19).length)))), (_this).f21, !((_this).f21), (_dafny.ZERO).minus((_this).f20), p1)).dtor_cf6);
      } else {
        (globalState).f18 = (_this).f21;
        let _171_v20;
        let _nw30 = Array((new BigNumber(14)).toNumber()).fill(false);
        _171_v20 = _nw30;
        let _172_v21;
        _172_v21 = _dafny.Seq.of(_171_v20, _171_v20);
        let _173_v22;
        let _nw31 = Array((new BigNumber(7)).toNumber());
        _nw31[(_dafny.ZERO).toNumber()] = ((true) ? (_171_v20) : (_171_v20));
        _nw31[(_dafny.ONE).toNumber()] = _171_v20;
        _nw31[(new BigNumber(2)).toNumber()] = _171_v20;
        _nw31[(new BigNumber(3)).toNumber()] = _171_v20;
        _nw31[(new BigNumber(4)).toNumber()] = _171_v20;
        _nw31[(new BigNumber(5)).toNumber()] = _171_v20;
        _nw31[(new BigNumber(6)).toNumber()] = (_172_v21)[_module.__default.safeIndex(_this.f27, new BigNumber((_172_v21).length))];
        _173_v22 = _nw31;
        let _174_v23;
        _174_v23 = _dafny.Seq.of(_173_v22);
        _173_v22 = (_174_v23)[_module.__default.safeIndex(p0, new BigNumber((_174_v23).length))];
        (_this).f27 = _this.f27;
        (globalState).f17 = p0;
        let _175_v24;
        _175_v24 = new _dafny.CodePoint('j'.codePointAt(0));
        let _176_v25;
        _176_v25 = _dafny.MultiSet.fromElements(_175_v24, new _dafny.CodePoint('l'.codePointAt(0)), _175_v24, _175_v24);
        _176_v25 = _dafny.MultiSet.fromElements(_175_v24, _175_v24, _module.__default.fm17(!((_this).f21), globalState), _175_v24, _175_v24);
      }
      let _177_v26;
      _177_v26 = _dafny.Seq.UnicodeFromString("ogosrhy");
      (globalState).f15 = _dafny.Seq.Concat(_177_v26, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(531)), function (_178_i2) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _177_v26));
      let _179_v27;
      let _nw32 = Array((new BigNumber(15)).toNumber()).fill(false);
      _179_v27 = _nw32;
      let _180_v28;
      _180_v28 = _dafny.Set.fromElements((_this).f21, (_this).f21);
      let _181_v29;
      _181_v29 = _dafny.Seq.of(_180_v28, _180_v28);
      let _182_v30;
      _182_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f27);
      let _183_v31;
      _183_v31 = new _dafny.CodePoint('m'.codePointAt(0));
      let _184_v32;
      let _nw33 = Array((new BigNumber(25)).toNumber());
      _nw33[(_dafny.ZERO).toNumber()] = _180_v28;
      _nw33[(_dafny.ONE).toNumber()] = _180_v28;
      _nw33[(new BigNumber(2)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(3)).toNumber()] = (_181_v29)[_module.__default.safeIndex(_this.f27, new BigNumber((_181_v29).length))];
      _nw33[(new BigNumber(4)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(!(p1), !((_this).f21));
      _nw33[(new BigNumber(6)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(7)).toNumber()] = _module.__default.fm18(new BigNumber((_dafny.Seq.UnicodeFromString("m")).length), (((_182_v30).contains(p1)) ? ((_182_v30).get(p1)) : (new BigNumber(221))), _183_v31, globalState);
      _nw33[(new BigNumber(8)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(9)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(10)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(11)).toNumber()] = (_module.D7.create_DC16(_180_v28)).dtor_cf24;
      _nw33[(new BigNumber(12)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(p1, (_this).f21);
      _nw33[(new BigNumber(14)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(15)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(16)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(17)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(18)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(19)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(true, p1);
      _nw33[(new BigNumber(21)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(22)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(23)).toNumber()] = _180_v28;
      _nw33[(new BigNumber(24)).toNumber()] = _module.__default.fm18(p0, p0, _183_v31, globalState);
      _184_v32 = _nw33;
      let _185_v33;
      _185_v33 = _dafny.Map.Empty.slice().updateUnsafe(_179_v27,_184_v32);
      _185_v33 = (_185_v33).update(_179_v27, _184_v32);
      let _index24 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_179_v27).length));
      (_179_v27)[_index24] = p1;
      let _186_v34;
      _186_v34 = _dafny.Seq.of(p1);
      let _index25 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_179_v27).length));
      (_179_v27)[_index25] = ((_186_v34)[_module.__default.safeIndex(new BigNumber((_182_v30).length), new BigNumber((_186_v34).length))]) && ((_this).f21);
      let _187_v35;
      _187_v35 = _dafny.Seq.of((_this).f28, (_this).f28);
      let _188_v36;
      _188_v36 = _dafny.Seq.of(_187_v35);
      r0 = _dafny.Seq.Concat((_188_v36)[_module.__default.safeIndex((_143_v5)[_module.__default.safeIndex(new BigNumber((_180_v28).length), new BigNumber((_143_v5).length))], new BigNumber((_188_v36).length))], _187_v35);
      r1 = p1;
      return [r0, r1];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f20, f21) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return !(_dafny.Seq.contains(_dafny.Seq.of((_this).f20, (_this).f20), (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of((_this).f21, (_this).f21, true, (_this).f21, (_this).f21)).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), function (_189_i0) {
        return (_this).f20;
      })).length)))));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f20;
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _190_v0;
      _190_v0 = _dafny.Seq.of(true);
      let _191_v1;
      let _192_v2;
      let _193_v3;
      let _194_v4;
      let _out2;
      let _out3;
      let _out4;
      let _out5;
      let _outcollector1 = (_this).m3(_dafny.areEqual(_190_v0, _190_v0), globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _out4 = _outcollector1[2];
      _out5 = _outcollector1[3];
      _191_v1 = _out2;
      _192_v2 = _out3;
      _193_v3 = _out4;
      _194_v4 = _out5;
      let _195_v5;
      _195_v5 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_191_v1)).length), (_this).f20);
      let _196_v6;
      _196_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_195_v5).length));
      _196_v6 = (_196_v6).update(p1, (_this).f20);
      if (true) {
        let _197_v7;
        _197_v7 = _dafny.Seq.of(p0, (_dafny.ZERO).minus((_this).f20));
        _197_v7 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_198_p0) => function (_199_i0) {
          return _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f20), _198_p0);
        })(p0)), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_200_p0) => function (_201_i0) {
          return _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f20), _200_p0);
        })(p0))).length)), p0);
        let _202_v8;
        _202_v8 = new _dafny.CodePoint('v'.codePointAt(0));
        _202_v8 = _202_v8;
        let _203_v9;
        _203_v9 = _dafny.MultiSet.fromElements((_this).f21);
        _193_v3 = (_this).fm1(_191_v1, (_this).fm1(_194_v4, _194_v4, _192_v2, globalState), (_203_v9).IsSubsetOf(_203_v9), globalState);
        let _204_v10;
        let _nw34 = new _module.C0();
        _nw34.__ctor();
        _204_v10 = _nw34;
        let _205_v11;
        let _nw35 = new _module.C0();
        _nw35.__ctor();
        _205_v11 = _nw35;
      } else {
        let _206_v12;
        let _nw36 = Array((_dafny.ONE).toNumber()).fill(false);
        _206_v12 = _nw36;
        let _207_v13;
        _207_v13 = _dafny.Seq.of(_206_v12);
        let _208_v14;
        _208_v14 = _module.D3.create_DC5(_207_v13);
        let _pat_let_tv4 = _207_v13;
        let _pat_let_tv5 = _207_v13;
        let _source6 = function (_pat_let0_0) {
          return function (_211_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_212_dt__update_hcf9_h1) {
                return _module.D3.create_DC5(_212_dt__update_hcf9_h1);
              }(_pat_let3_0);
            }(_pat_let_tv5);
          }(_pat_let0_0);
        }((((_this).f21) ? (_208_v14) : (function (_pat_let1_0) {
          return function (_209_dt__update__tmp_h0) {
            return function (_pat_let2_0) {
              return function (_210_dt__update_hcf9_h0) {
                return _module.D3.create_DC5(_210_dt__update_hcf9_h0);
              }(_pat_let2_0);
            }(_pat_let_tv4);
          }(_pat_let1_0);
        }(_208_v14))));
        if (_source6.is_DC6) {
          let _213___mcc_h0 = (_source6).cf10;
          let _214_cf10 = _213___mcc_h0;
          let _215_v15;
          let _init4 = ((_216_v3) => function (_217_i1) {
            return (_217_i1).minus(_216_v3);
          })(_193_v3);
          let _nw37 = Array((new BigNumber(16)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
            _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _215_v15 = _nw37;
          let _218_v16;
          let _nw38 = new _module.C1();
          _nw38.__ctor(new BigNumber((_dafny.Set.fromElements(_192_v2)).length), _215_v15, _194_v4, true);
          _218_v16 = _nw38;
          let _219_v17;
          _219_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_218_v16);
          let _220_v18;
          _220_v18 = _dafny.Set.fromElements(_192_v2, p1, _192_v2, _192_v2);
          let _221_v19;
          _221_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f20).plus(new BigNumber((_219_v17).length)),_220_v18);
          let _222_v20;
          _222_v20 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index26 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_206_v12).length));
          (_206_v12)[_index26] = (_220_v18).IsDisjointFrom(_module.__default.fm18(_214_cf10, _194_v4, _222_v20, globalState));
          let _223_v21;
          _223_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
          let _224_v22;
          _224_v22 = _dafny.Seq.UnicodeFromString("nlhaskp");
          let _225_v23;
          _225_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_190_v0)[_module.__default.safeIndex(_193_v3, new BigNumber((_190_v0).length))]);
          let _226_v24;
          _226_v24 = _dafny.Seq.of(_module.__default.fm19(true, globalState), _225_v23);
          let _index27 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_206_v12).length));
          let _rhs13 = _221_v19;
          let _rhs14 = (_this).f20;
          let _rhs15 = (_218_v16).fm0(_223_v21, _224_v22, _226_v24, globalState);
          let _rhs16 = !(p1);
          let _lhs16 = globalState;
          let _lhs17 = _206_v12;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_206_v12).length));
          _221_v19 = _rhs13;
          _214_cf10 = _rhs14;
          _lhs16.f18 = _rhs15;
          _lhs17[_lhs18] = _rhs16;
          _190_v0 = _dafny.Seq.Concat(_190_v0, _190_v0);
          let _227_v25;
          _227_v25 = _215_v15;
          let _228_v26;
          _228_v26 = _dafny.Seq.of(_215_v15);
          let _229_v27;
          _229_v27 = _dafny.Map.Empty.slice().updateUnsafe(_227_v25,(_228_v26)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_228_v26).length))]);
          let _230_v28;
          _230_v28 = _dafny.Set.fromElements((((_229_v27).contains(_227_v25)) ? ((_229_v27).get(_227_v25)) : (_215_v15)), _215_v15, _215_v15);
          let _index28 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_206_v12).length));
          (_206_v12)[_index28] = !(!(!((_230_v28).IsSubsetOf(_230_v28))));
          (globalState).f3 = (_this).f21;
        } else if (_source6.is_DC7) {
          let _231___mcc_h1 = (_source6).cf11;
          let _232___mcc_h2 = (_source6).cf12;
          let _233_cf12 = _232___mcc_h2;
          let _234_cf11 = _231___mcc_h1;
          _191_v1 = (_this).f20;
          let _235_v29;
          let _nw39 = Array((new BigNumber(20)).toNumber());
          _235_v29 = _nw39;
          let _236_v30;
          _236_v30 = _dafny.Map.Empty.slice().updateUnsafe(_192_v2,_235_v29);
          let _237_v31;
          _237_v31 = _dafny.Seq.of(_236_v30, _236_v30, _dafny.Map.Empty.slice().updateUnsafe(false,_235_v29));
          _192_v2 = (new BigNumber(((_237_v31)[_module.__default.safeIndex(p0, new BigNumber((_237_v31).length))]).length)).isEqualTo(p0);
          let _238_v32;
          _238_v32 = _dafny.Seq.of(_194_v4);
          _193_v3 = ((_193_v3).plus((_238_v32)[_module.__default.safeIndex(_194_v4, new BigNumber((_238_v32).length))])).minus(p0);
          let _239_v33;
          _239_v33 = _dafny.Map.Empty.slice().updateUnsafe(_192_v2,p1);
          let _240_v34;
          _240_v34 = _dafny.Seq.UnicodeFromString("uj");
          let _241_v35;
          _241_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(518),_191_v1);
          let _242_v36;
          _242_v36 = p1;
          let _243_v37;
          _243_v37 = _dafny.Map.Empty.slice().updateUnsafe(_242_v36,false);
          let _244_v38;
          _244_v38 = _dafny.Map.Empty.slice().updateUnsafe(_206_v12,false);
          let _245_v39;
          let _nw40 = Array((new BigNumber(28)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = new BigNumber((_239_v33).length);
          _nw40[(_dafny.ONE).toNumber()] = _191_v1;
          _nw40[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_240_v34,new BigNumber((_241_v35).length))).length);
          _nw40[(new BigNumber(3)).toNumber()] = _193_v3;
          _nw40[(new BigNumber(4)).toNumber()] = _194_v4;
          _nw40[(new BigNumber(5)).toNumber()] = (_this).fm1(_191_v1, new BigNumber((_241_v35).length), _233_cf12, globalState);
          _nw40[(new BigNumber(6)).toNumber()] = new BigNumber(388);
          _nw40[(new BigNumber(7)).toNumber()] = _191_v1;
          _nw40[(new BigNumber(8)).toNumber()] = _193_v3;
          _nw40[(new BigNumber(9)).toNumber()] = new BigNumber(135);
          _nw40[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()), _194_v4);
          _nw40[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_240_v34).length))).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-625)), function (_246_i2) {
            return (_this).f20;
          })).length));
          _nw40[(new BigNumber(12)).toNumber()] = ((_233_cf12) ? ((_this).f20) : (new BigNumber((_243_v37).length)));
          _nw40[(new BigNumber(13)).toNumber()] = new BigNumber((_244_v38).length);
          _nw40[(new BigNumber(14)).toNumber()] = (_this).f20;
          _nw40[(new BigNumber(15)).toNumber()] = _194_v4;
          _nw40[(new BigNumber(16)).toNumber()] = (_this).f20;
          _nw40[(new BigNumber(17)).toNumber()] = _191_v1;
          _nw40[(new BigNumber(18)).toNumber()] = new BigNumber(730);
          _nw40[(new BigNumber(19)).toNumber()] = _191_v1;
          _nw40[(new BigNumber(20)).toNumber()] = p0;
          _nw40[(new BigNumber(21)).toNumber()] = (_this).f20;
          _nw40[(new BigNumber(22)).toNumber()] = (_this).f20;
          _nw40[(new BigNumber(23)).toNumber()] = _193_v3;
          _nw40[(new BigNumber(24)).toNumber()] = _193_v3;
          _nw40[(new BigNumber(25)).toNumber()] = p0;
          _nw40[(new BigNumber(26)).toNumber()] = _191_v1;
          _nw40[(new BigNumber(27)).toNumber()] = new BigNumber(525);
          _245_v39 = _nw40;
          let _index29 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_245_v39).length));
          (_245_v39)[_index29] = _193_v3;
          let _247_v40;
          _247_v40 = _dafny.Set.fromElements(p1);
          let _248_v41;
          _248_v41 = _module.D7.create_DC16(_247_v40);
          let _index30 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_245_v39).length));
          (_245_v39)[_index30] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_248_v41, _248_v41), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), ((_249_v41) => function (_250_i3) {
            return _249_v41;
          })(_248_v41)), _module.__default.fm20((_this).f20, _193_v3, _193_v3, globalState)))).length);
        } else if (_source6.is_DC5) {
          let _251___mcc_h3 = (_source6).cf9;
          let _252_cf9 = _251___mcc_h3;
          let _253_v42;
          _253_v42 = _dafny.Map.Empty.slice().updateUnsafe(_194_v4,true);
          let _254_v43;
          _254_v43 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC16(_dafny.Set.fromElements((((_253_v42).contains(new BigNumber(636))) ? ((_253_v42).get(new BigNumber(636))) : ((_this).f21)), (_this).f21)),p1);
          let _255_v45;
          _255_v45 = _dafny.Seq.UnicodeFromString("vnfdr");
          let _256_v46;
          _256_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
          let _257_v47;
          _257_v47 = _dafny.Seq.of(_256_v46);
          let _258_v48;
          _258_v48 = _dafny.Map.Empty.slice().updateUnsafe(!(_192_v2),(_this).fm0(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(926), new BigNumber(940))) {
              let _259_v44 = _compr_16;
              if (((new BigNumber(926)).isLessThanOrEqualTo(_259_v44)) && ((_259_v44).isLessThan(new BigNumber(940)))) {
                _coll16.push([(_259_v44).plus(new BigNumber((_190_v0).length)),new BigNumber((_dafny.Set.fromElements(true, p1, _192_v2)).length)]);
              }
            }
            return _coll16;
          }(), _255_v45, _257_v47, globalState));
          let _260_v49;
          _260_v49 = _dafny.Set.fromElements((((_256_v46).contains(_192_v2)) ? ((_256_v46).get(_192_v2)) : (true)));
          let _261_v50;
          _261_v50 = _module.D7.create_DC16(_260_v49);
          (globalState).f3 = (((_254_v43).contains(_261_v50)) ? ((_254_v43).get(_261_v50)) : (!(p1)));
          let _262_v51;
          _262_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_255_v45).length),(_this).f20);
          let _263_v52;
          _263_v52 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_193_v3),_257_v47);
          let _264_v53;
          _264_v53 = _dafny.MultiSet.fromElements((_this).fm0(_262_v51, _255_v45, (((_263_v52).contains(p0)) ? ((_263_v52).get(p0)) : (_257_v47)), globalState), _192_v2);
          let _265_v54;
          _265_v54 = _dafny.MultiSet.fromElements(((_this).f21) === ((_this).f21), (_264_v53).IsProperSubsetOf(_dafny.MultiSet.fromElements(!(!(p1)), (_this).f21)), p1);
          let _266_v55;
          _266_v55 = _dafny.Seq.of(_265_v54);
          _264_v53 = ((_266_v55)[_module.__default.safeIndex(_194_v4, new BigNumber((_266_v55).length))]).Difference((_dafny.MultiSet.fromElements(!(p1), true)).Difference(_264_v53));
          let _267_v56;
          _267_v56 = _dafny.Seq.of(p0, _193_v3);
          let _268_v57;
          _268_v57 = _module.D7.create_DC18(_267_v56, p0);
          let _269_v58;
          let _nw41 = Array((new BigNumber(9)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _191_v1;
          _nw41[(_dafny.ONE).toNumber()] = new BigNumber((_255_v45).length);
          _nw41[(new BigNumber(2)).toNumber()] = _191_v1;
          _nw41[(new BigNumber(3)).toNumber()] = _191_v1;
          _nw41[(new BigNumber(4)).toNumber()] = (_this).f20;
          _nw41[(new BigNumber(5)).toNumber()] = _194_v4;
          _nw41[(new BigNumber(6)).toNumber()] = p0;
          _nw41[(new BigNumber(7)).toNumber()] = (_268_v57).dtor_cf27;
          _nw41[(new BigNumber(8)).toNumber()] = (_this).f20;
          _269_v58 = _nw41;
          let _270_v59;
          _270_v59 = _269_v58;
          let _271_v60;
          _271_v60 = _dafny.MultiSet.fromElements(_270_v59, _270_v59);
          let _272_v61;
          _272_v61 = _dafny.Map.Empty.slice().updateUnsafe(_194_v4,_271_v60);
          let _273_v62;
          _273_v62 = _module.D6.create_DC14(new BigNumber(764));
          let _274_v63;
          _274_v63 = _dafny.Seq.of(_269_v58);
          let _275_v64;
          _275_v64 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),(_this).f21);
          let _276_v65;
          _276_v65 = _module.D8.create_DC19(_262_v51);
          let _277_v66;
          let _nw42 = Array((new BigNumber(26)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _271_v60;
          _nw42[(_dafny.ONE).toNumber()] = _271_v60;
          _nw42[(new BigNumber(2)).toNumber()] = (_271_v60).Difference(_271_v60);
          _nw42[(new BigNumber(3)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(4)).toNumber()] = ((((_272_v61).contains((_273_v62).dtor_cf22)) ? ((_272_v61).get((_273_v62).dtor_cf22)) : ((((_272_v61).contains(_191_v1)) ? ((_272_v61).get(_191_v1)) : (_271_v60))))).Union(_271_v60);
          _nw42[(new BigNumber(5)).toNumber()] = (_271_v60).Intersect(_dafny.MultiSet.fromElements(_270_v59));
          _nw42[(new BigNumber(6)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(7)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(8)).toNumber()] = (_271_v60).update(_270_v59, _module.__default.abs(_193_v3));
          _nw42[(new BigNumber(9)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(10)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(11)).toNumber()] = (_271_v60).Union(_271_v60);
          _nw42[(new BigNumber(12)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_274_v63);
          _nw42[(new BigNumber(14)).toNumber()] = ((_271_v60).update(_270_v59, _module.__default.abs(new BigNumber((_275_v64).length)))).Intersect(_271_v60);
          _nw42[(new BigNumber(15)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(16)).toNumber()] = (_271_v60).update(_270_v59, _module.__default.abs(_193_v3));
          _nw42[(new BigNumber(17)).toNumber()] = (_dafny.MultiSet.fromElements(_269_v58, _269_v58, _270_v59)).update(_270_v59, _module.__default.abs((((_262_v51).contains(new BigNumber(((_276_v65).dtor_cf28).length))) ? ((_262_v51).get(new BigNumber(((_276_v65).dtor_cf28).length))) : (_194_v4))));
          _nw42[(new BigNumber(18)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(19)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(20)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(21)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(22)).toNumber()] = (_271_v60).Intersect(_271_v60);
          _nw42[(new BigNumber(23)).toNumber()] = _271_v60;
          _nw42[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.fromElements(_270_v59);
          _nw42[(new BigNumber(25)).toNumber()] = _271_v60;
          _277_v66 = _nw42;
          let _index31 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_277_v66).length));
          (_277_v66)[_index31] = _271_v60;
          let _index32 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_277_v66).length));
          let _rhs17 = !(((p1) ? ((_this).f21) : (p1)));
          let _rhs18 = (_dafny.ZERO).minus((_this).f20);
          let _rhs19 = _271_v60;
          let _lhs19 = globalState;
          let _lhs20 = globalState;
          let _lhs21 = _277_v66;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_277_v66).length));
          _lhs19.f4 = _rhs17;
          _lhs20.f5 = _rhs18;
          _lhs21[_lhs22] = _rhs19;
          let _278_v67;
          _278_v67 = _dafny.Seq.of(_190_v0);
          let _279_v68;
          _279_v68 = new _dafny.CodePoint('s'.codePointAt(0));
          let _280_v69;
          _280_v69 = _dafny.Set.fromElements(_279_v68);
          let _281_v70;
          _281_v70 = _module.D9.create_DC22(_190_v0);
          let _282_v71;
          _282_v71 = _module.D9.create_DC23(_193_v3, _192_v2);
          let _283_v72;
          _283_v72 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), ((_284_v68) => function (_285_i4) {
            return _284_v68;
          })(_279_v68))).length)));
          let _286_v73;
          _286_v73 = _dafny.Map.Empty.slice().updateUnsafe(_283_v72,p1);
          let _287_v74;
          let _nw43 = Array((new BigNumber(22)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f21, _192_v2), _190_v0);
          _nw43[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_module.__default.fm21(globalState), _module.__default.safeIndex(new BigNumber(-21), new BigNumber((_module.__default.fm21(globalState)).length)), (_this).f21);
          _nw43[(new BigNumber(2)).toNumber()] = ((p1) ? (_190_v0) : (_dafny.Seq.of(p1, p1)));
          _nw43[(new BigNumber(3)).toNumber()] = (_278_v67)[_module.__default.safeIndex(new BigNumber((_280_v69).length), new BigNumber((_278_v67).length))];
          _nw43[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_190_v0, _module.__default.safeIndex(_194_v4, new BigNumber((_190_v0).length)), p1);
          _nw43[(new BigNumber(5)).toNumber()] = _190_v0;
          _nw43[(new BigNumber(6)).toNumber()] = _module.__default.fm21(globalState);
          _nw43[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(p1);
          _nw43[(new BigNumber(8)).toNumber()] = _190_v0;
          _nw43[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_190_v0, _module.__default.safeIndex(_194_v4, new BigNumber((_190_v0).length)), _192_v2);
          _nw43[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_190_v0, _190_v0), _module.__default.safeIndex(new BigNumber(112), new BigNumber((_dafny.Seq.Concat(_190_v0, _190_v0)).length)), true), _module.__default.safeIndex(_194_v4, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_190_v0, _190_v0), _module.__default.safeIndex(new BigNumber(112), new BigNumber((_dafny.Seq.Concat(_190_v0, _190_v0)).length)), true)).length)), _192_v2);
          _nw43[(new BigNumber(11)).toNumber()] = _190_v0;
          _nw43[(new BigNumber(12)).toNumber()] = _190_v0;
          _nw43[(new BigNumber(13)).toNumber()] = _190_v0;
          _nw43[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p1, _192_v2), (_281_v70).dtor_cf30), _module.__default.safeIndex(_194_v4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p1, _192_v2), (_281_v70).dtor_cf30)).length)), (_this).f21);
          _nw43[(new BigNumber(15)).toNumber()] = _module.__default.fm21(globalState);
          _nw43[(new BigNumber(16)).toNumber()] = (((_282_v71).dtor_cf32) ? (_190_v0) : (_190_v0));
          _nw43[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_this).f21);
          _nw43[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_190_v0, _dafny.Seq.of(p1, (_this).fm0(_module.__default.fm22(_193_v3, globalState), _255_v45, _dafny.Seq.of(_258_v48), globalState)));
          _nw43[(new BigNumber(19)).toNumber()] = _module.__default.fm21(globalState);
          _nw43[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_190_v0, _dafny.Seq.of((_this).f21, (((_286_v73).contains(_283_v72)) ? ((_286_v73).get(_283_v72)) : ((_this).f21)), (_190_v0)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f20), new BigNumber((_190_v0).length))], _192_v2));
          _nw43[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_190_v0, _dafny.Seq.update(_dafny.Seq.of((_this).f21, p1), _module.__default.safeIndex((_dafny.ZERO).minus(_194_v4), new BigNumber((_dafny.Seq.of((_this).f21, p1)).length)), (_this).f21));
          _287_v74 = _nw43;
          let _index33 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_287_v74).length));
          (_287_v74)[_index33] = _190_v0;
          let _index34 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_287_v74).length));
          (_287_v74)[_index34] = _dafny.Seq.Concat(_dafny.Seq.of(_192_v2), _190_v0);
        } else {
          let _288___mcc_h4 = (_source6).cf13;
          let _289_cf13 = _288___mcc_h4;
          let _290_v75;
          _290_v75 = _dafny.Seq.of((_this).f20);
          let _291_v76;
          _291_v76 = _module.D7.create_DC18(_290_v75, (_this).f20);
          let _292_v77;
          _292_v77 = _dafny.Map.Empty.slice().updateUnsafe(_192_v2,p1);
          let _293_v78;
          _293_v78 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_192_v2,p1)).update(_192_v2, p1), _292_v77, _292_v77);
          let _294_v79;
          _294_v79 = _dafny.Map.Empty.slice().updateUnsafe(_191_v1,(_this).f21);
          let _rhs20 = (_194_v4).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.update((_291_v76).dtor_cf26, _module.__default.safeIndex(new BigNumber(((_293_v78)[_module.__default.safeIndex(_194_v4, new BigNumber((_293_v78).length))]).length), new BigNumber(((_291_v76).dtor_cf26).length)), _193_v3), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update((_291_v76).dtor_cf26, _module.__default.safeIndex(new BigNumber(((_293_v78)[_module.__default.safeIndex(_194_v4, new BigNumber((_293_v78).length))]).length), new BigNumber(((_291_v76).dtor_cf26).length)), _193_v3)).length)), new BigNumber((_294_v79).length))).length));
          let _rhs21 = p0;
          let _lhs23 = globalState;
          _lhs23.f17 = _rhs20;
          _194_v4 = _rhs21;
          let _295_v80;
          _295_v80 = _dafny.Seq.UnicodeFromString("lmfshk");
          let _296_v81;
          _296_v81 = _dafny.Map.Empty.slice().updateUnsafe(_295_v80,new BigNumber(170));
          _296_v81 = (_296_v81).update(_dafny.Seq.Concat(_295_v80, _dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_297_i5) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })), _191_v1);
          _194_v4 = (_this).fm1(_194_v4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-454)), function (_298_i6) {
            return new BigNumber(840);
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(85)), ((_299_p0) => function (_300_i7) {
            return _299_p0;
          })(p0)))).length), _192_v2, globalState);
          let _301_v82;
          _301_v82 = new _dafny.CodePoint('m'.codePointAt(0));
          _301_v82 = _301_v82;
        }
        let _302_v83;
        _302_v83 = _dafny.Seq.UnicodeFromString("xoj");
        (globalState).f15 = _dafny.Seq.Concat(_302_v83, _dafny.Seq.UnicodeFromString("hsyle"));
        let _303_v84;
        _303_v84 = _dafny.Seq.of(_191_v1);
        let _304_v85;
        _304_v85 = _dafny.MultiSet.fromElements(_192_v2);
        let _305_v86;
        _305_v86 = _dafny.MultiSet.fromElements((_module.__default.fm23(globalState)).update((_this).f21, _module.__default.abs(_194_v4)), _304_v85, _304_v85);
        let _306_v87;
        let _nw44 = Array((new BigNumber(23)).toNumber());
        _nw44[(_dafny.ZERO).toNumber()] = (_this).f20;
        _nw44[(_dafny.ONE).toNumber()] = (_this).f20;
        _nw44[(new BigNumber(2)).toNumber()] = (((_this).f21) ? ((_this).f20) : ((_dafny.ZERO).minus(p0)));
        _nw44[(new BigNumber(3)).toNumber()] = _191_v1;
        _nw44[(new BigNumber(4)).toNumber()] = (_this).f20;
        _nw44[(new BigNumber(5)).toNumber()] = _191_v1;
        _nw44[(new BigNumber(6)).toNumber()] = _194_v4;
        _nw44[(new BigNumber(7)).toNumber()] = new BigNumber((_195_v5).length);
        _nw44[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_194_v4, _191_v1);
        _nw44[(new BigNumber(9)).toNumber()] = p0;
        _nw44[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_193_v3, _194_v4);
        _nw44[(new BigNumber(11)).toNumber()] = _194_v4;
        _nw44[(new BigNumber(12)).toNumber()] = _193_v3;
        _nw44[(new BigNumber(13)).toNumber()] = _194_v4;
        _nw44[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_191_v1);
        _nw44[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_194_v4);
        _nw44[(new BigNumber(16)).toNumber()] = (_303_v84)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(p1)).length), new BigNumber((_303_v84).length))];
        _nw44[(new BigNumber(17)).toNumber()] = p0;
        _nw44[(new BigNumber(18)).toNumber()] = (_this).f20;
        _nw44[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(_194_v4, (_this).f20);
        _nw44[(new BigNumber(20)).toNumber()] = ((_this).f20).multipliedBy(_194_v4);
        _nw44[(new BigNumber(21)).toNumber()] = p0;
        _nw44[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_305_v86).cardinality())));
        _306_v87 = _nw44;
        let _index35 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_306_v87).length));
        (_306_v87)[_index35] = p0;
        let _index36 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_306_v87).length));
        (_306_v87)[_index36] = ((p0).multipliedBy(new BigNumber((_302_v83).length))).minus(((_this).f20).plus((_this).fm1(_194_v4, _193_v3, (_this).f21, globalState)));
        (globalState).f5 = _191_v1;
        let _307_v89;
        _307_v89 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f20));
        let _308_v90;
        _308_v90 = _module.D8.create_DC19(function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of (_307_v89).Elements) {
    let _309_v88 = _compr_17;
    if ((_307_v89).contains(_309_v88)) {
      _coll17.push([(_309_v88).multipliedBy(new BigNumber((_302_v83).length)),_193_v3]);
    }
  }
  return _coll17;
}());
        let _310_v91;
        _310_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_307_v89).cardinality()),new BigNumber(409));
        _308_v90 = _module.D8.create_DC19(_310_v91);
      }
      let _311_v92;
      _311_v92 = _module.D4.create_DC11(_191_v1, (_190_v0)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_190_v0).length))], p1);
      let _312_v93;
      _312_v93 = _dafny.Set.fromElements(p1);
      let _pat_let_tv6 = _192_v2;
      let _pat_let_tv7 = _193_v3;
      let _313_v94;
      let _nw45 = Array((new BigNumber(28)).toNumber());
      _nw45[(_dafny.ZERO).toNumber()] = _311_v92;
      _nw45[(_dafny.ONE).toNumber()] = _311_v92;
      _nw45[(new BigNumber(2)).toNumber()] = _module.D4.create_DC11((_this).f20, p1, _192_v2);
      _nw45[(new BigNumber(3)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(4)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(5)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(6)).toNumber()] = _module.__default.fm24(_191_v1, _192_v2, _312_v93, p1, globalState);
      _nw45[(new BigNumber(7)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(8)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(9)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(10)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(11)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(12)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(13)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(14)).toNumber()] = _module.D4.create_DC11(_191_v1, (_this).f21, _192_v2);
      _nw45[(new BigNumber(15)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(16)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(17)).toNumber()] = _module.__default.fm24(new BigNumber(127), p1, _312_v93, _192_v2, globalState);
      _nw45[(new BigNumber(18)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(19)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(20)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(21)).toNumber()] = function (_pat_let4_0) {
        return function (_314_dt__update__tmp_h2) {
          return function (_pat_let5_0) {
            return function (_315_dt__update_hcf19_h0) {
              return _module.D4.create_DC11((_314_dt__update__tmp_h2).dtor_cf17, (_314_dt__update__tmp_h2).dtor_cf18, _315_dt__update_hcf19_h0);
            }(_pat_let5_0);
          }((_this).f21);
        }(_pat_let4_0);
      }(_311_v92);
      _nw45[(new BigNumber(22)).toNumber()] = function (_pat_let6_0) {
        return function (_316_dt__update__tmp_h3) {
          return function (_pat_let7_0) {
            return function (_317_dt__update_hcf18_h0) {
              return _module.D4.create_DC11((_316_dt__update__tmp_h3).dtor_cf17, _317_dt__update_hcf18_h0, (_316_dt__update__tmp_h3).dtor_cf19);
            }(_pat_let7_0);
          }(_pat_let_tv6);
        }(_pat_let6_0);
      }(_311_v92);
      _nw45[(new BigNumber(23)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(24)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(25)).toNumber()] = _311_v92;
      _nw45[(new BigNumber(26)).toNumber()] = function (_pat_let8_0) {
        return function (_318_dt__update__tmp_h4) {
          return function (_pat_let9_0) {
            return function (_319_dt__update_hcf18_h1) {
              return function (_pat_let10_0) {
                return function (_320_dt__update_hcf17_h0) {
                  return _module.D4.create_DC11(_320_dt__update_hcf17_h0, _319_dt__update_hcf18_h1, (_318_dt__update__tmp_h4).dtor_cf19);
                }(_pat_let10_0);
              }(_pat_let_tv7);
            }(_pat_let9_0);
          }((_this).f21);
        }(_pat_let8_0);
      }(_311_v92);
      _nw45[(new BigNumber(27)).toNumber()] = _311_v92;
      _313_v94 = _nw45;
      let _index37 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_313_v94).length));
      (_313_v94)[_index37] = _module.D4.create_DC11((_this).f20, !((_this).f21), (_this).f21);
      let _pat_let_tv8 = _191_v1;
      let _pat_let_tv9 = _191_v1;
      let _index38 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_313_v94).length));
      (_313_v94)[_index38] = function (_pat_let11_0) {
        return function (_321_dt__update__tmp_h5) {
          return function (_pat_let12_0) {
            return function (_322_dt__update_hcf19_h1) {
              return _module.D4.create_DC11((_321_dt__update__tmp_h5).dtor_cf17, (_321_dt__update__tmp_h5).dtor_cf18, _322_dt__update_hcf19_h1);
            }(_pat_let12_0);
          }((_pat_let_tv8).isLessThan(_pat_let_tv9));
        }(_pat_let11_0);
      }(_311_v92);
      let _323_v95;
      _323_v95 = new _dafny.CodePoint('w'.codePointAt(0));
      let _324_v96;
      _324_v96 = _dafny.Map.Empty.slice().updateUnsafe(_323_v95,p1);
      if ((_192_v2) && (!((((_324_v96).contains(_module.__default.fm17(_192_v2, globalState))) ? ((_324_v96).get(_module.__default.fm17(_192_v2, globalState))) : ((_190_v0)[_module.__default.safeIndex(_191_v1, new BigNumber((_190_v0).length))]))) || (p1))) {
        let _325_v97;
        _325_v97 = _dafny.Seq.of(_193_v3, _191_v1);
        let _326_v98;
        _326_v98 = _dafny.Map.Empty.slice().updateUnsafe(_191_v1,new BigNumber((_325_v97).length));
        _326_v98 = (_326_v98).Merge(_module.__default.fm22(new BigNumber(398), globalState));
        let _327_v99;
        _327_v99 = _dafny.Seq.UnicodeFromString("wqqgtwun");
        let _328_v100;
        _328_v100 = _dafny.Map.Empty.slice().updateUnsafe(_327_v99,_323_v95);
        let _329_v101;
        _329_v101 = _dafny.Seq.of(_325_v97);
        _325_v97 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_328_v100).length)), (_329_v101)[_module.__default.safeIndex(_194_v4, new BigNumber((_329_v101).length))]), _325_v97);
        let _330_v102;
        let _nw46 = Array((new BigNumber(4)).toNumber());
        _nw46[(_dafny.ZERO).toNumber()] = (_this).f21;
        _nw46[(_dafny.ONE).toNumber()] = false;
        _nw46[(new BigNumber(2)).toNumber()] = false;
        _nw46[(new BigNumber(3)).toNumber()] = p1;
        _330_v102 = _nw46;
        let _331_v103;
        _331_v103 = _dafny.Seq.of(_330_v102, _330_v102, _330_v102, _330_v102, _330_v102);
        let _332_v104;
        _332_v104 = _module.D3.create_DC5(_331_v103);
        _332_v104 = _332_v104;
        let _333_v105;
        _333_v105 = _dafny.Seq.of((_dafny.Set.fromElements(new BigNumber((_196_v6).length), (_this).f20)).Difference(_module.__default.fm25(globalState)), _195_v5, _195_v5);
        _333_v105 = _dafny.Seq.of(_195_v5, _195_v5, _195_v5);
        _192_v2 = p1;
      } else {
        let _334_v106;
        _334_v106 = _dafny.Seq.of(_194_v4);
        (globalState).f5 = ((_334_v106)[_module.__default.safeIndex((_this).f20, new BigNumber((_334_v106).length))]).minus(_191_v1);
        let _335_v107;
        let _nw47 = new _module.C0();
        _nw47.__ctor();
        _335_v107 = _nw47;
        let _336_v108;
        _336_v108 = _dafny.Seq.UnicodeFromString("ssmu");
        let _337_v109;
        _337_v109 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_336_v108, _module.__default.safeIndex(_194_v4, new BigNumber((_336_v108).length)), _module.__default.fm17(p1, globalState))).length), _module.__default.safeDivisionInt(_194_v4, p0), (_dafny.ZERO).minus((_this).f20), (new BigNumber((_336_v108).length)).minus(_194_v4), (_this).fm1(_191_v1, _193_v3, _192_v2, globalState));
        _337_v109 = (_337_v109).Union(_337_v109);
        let _338_v110;
        let _nw48 = new _module.C0();
        _nw48.__ctor();
        _338_v110 = _nw48;
        let _339_v111;
        _339_v111 = _module.D8.create_DC20();
        _339_v111 = _module.D8.create_DC20();
      }
      let _340_v112;
      _340_v112 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_194_v4);
      let _341_v113;
      _341_v113 = _dafny.Seq.UnicodeFromString("ac");
      _340_v112 = (_340_v112).update(((_this).f20).plus(new BigNumber((_341_v113).length)), p0);
      let _342_v114;
      let _nw49 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _342_v114 = _nw49;
      let _343_v115;
      _343_v115 = _dafny.Seq.of(_342_v114, _342_v114, _342_v114);
      r0 = ((_192_v2) ? (_343_v115) : (_343_v115));
      r1 = p1;
      return [r0, r1];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _hi1 = (_this).f20;
      for (let _344_i0 = new BigNumber((_module.__default.fm26((_this).f21, globalState)).length); _344_i0.isLessThan(_hi1); _344_i0 = _344_i0.plus(_dafny.ONE)) {
        (globalState).f3 = false;
        let _345_v0;
        _345_v0 = _dafny.Seq.UnicodeFromString("nfbohx");
        let _346_v1;
        _346_v1 = new _dafny.CodePoint('r'.codePointAt(0));
        (globalState).f15 = _dafny.Seq.Concat(_dafny.Seq.update(_345_v0, _module.__default.safeIndex(_344_i0, new BigNumber((_345_v0).length)), _346_v1), _345_v0);
        (globalState).f15 = _345_v0;
        (globalState).f3 = (_this).f21;
      }
      let _347_v2;
      _347_v2 = _dafny.Seq.UnicodeFromString("fahojwr");
      (globalState).f5 = (_this).fm1((_this).f20, new BigNumber((_347_v2).length), (_this).f21, globalState);
      let _348_v3;
      let _init5 = function (_349_i1) {
        return _module.__default.safeModuloInt(_349_i1, (_this).f20);
      };
      let _nw50 = Array((new BigNumber(4)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw50.length); _i0_5++) {
        _nw50[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _348_v3 = _nw50;
      let _index39 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
      (_348_v3)[_index39] = new BigNumber((_dafny.Seq.of((_this).f20, (_this).f20)).length);
      let _350_v4;
      _350_v4 = new _dafny.CodePoint('g'.codePointAt(0));
      let _351_v5;
      _351_v5 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_347_v2, _module.__default.safeIndex((_this).f20, new BigNumber((_347_v2).length)), _350_v4)).length));
      let _352_v6;
      _352_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_351_v5).length));
      let _index40 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
      (_348_v3)[_index40] = ((_this).f20).minus((((_352_v6).contains((_this).f21)) ? ((_352_v6).get((_this).f21)) : ((_this).f20)));
      let _353_v7;
      _353_v7 = _dafny.Map.Empty.slice().updateUnsafe(_348_v3,new BigNumber((_dafny.Seq.UnicodeFromString("wpkgyo")).length));
      let _354_v8;
      _354_v8 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)));
      let _355_v9;
      _355_v9 = _module.D7.create_DC18(_351_v5, new BigNumber((_354_v8).cardinality()));
      let _hi2 = _module.__default.safeModuloInt((((_353_v7).contains(_348_v3)) ? ((_353_v7).get(_348_v3)) : ((_355_v9).dtor_cf27)), (_this).f20);
      for (let _356_i2 = (new BigNumber((_347_v2).length)).plus((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]); _356_i2.isLessThan(_hi2); _356_i2 = _356_i2.plus(_dafny.ONE)) {
        let _357_v10;
        _357_v10 = _module.D3.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("ksptncyh")).length));
        let _358_v11;
        _358_v11 = _dafny.Map.Empty.slice().updateUnsafe(_356_i2,(_dafny.ZERO).minus(_356_i2));
        let _359_v12;
        _359_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,!((_this).f21));
        let _360_v13;
        _360_v13 = _dafny.Seq.of(_359_v12, _dafny.Map.Empty.slice().updateUnsafe(p0,false), _359_v12, _359_v12);
        let _361_v14;
        _361_v14 = _dafny.Seq.of(_359_v12, (_360_v13)[_module.__default.safeIndex((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], new BigNumber((_360_v13).length))], _359_v12, _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21));
        let _index41 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        let _rhs22 = _350_v4;
        let _rhs23 = (_357_v10).dtor_cf10;
        let _rhs24 = (_this).f21;
        let _rhs25 = !((_this).f21) || (!((_this).fm0(_358_v11, _347_v2, _361_v14, globalState)));
        let _lhs24 = _348_v3;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        let _lhs26 = globalState;
        let _lhs27 = globalState;
        _350_v4 = _rhs22;
        _lhs24[_lhs25] = _rhs23;
        _lhs26.f14 = _rhs24;
        _lhs27.f18 = _rhs25;
        let _362_v15;
        _362_v15 = _dafny.Seq.of((_this).fm0(_358_v11, _347_v2, _361_v14, globalState), p0, (_this).fm0(_358_v11, _dafny.Seq.UnicodeFromString("mgw"), _360_v13, globalState));
        let _363_v16;
        _363_v16 = _module.D4.create_DC11((_this).f20, (_this).f21, !((_this).f21));
        if ((_362_v15)[_module.__default.safeIndex((_363_v16).dtor_cf17, new BigNumber((_362_v15).length))]) {
          let _364_v17;
          let _nw51 = Array((new BigNumber(10)).toNumber()).fill(false);
          _364_v17 = _nw51;
          let _index42 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_364_v17).length));
          (_364_v17)[_index42] = (_this).fm0(function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-127), new BigNumber(39))) {
              let _365_v18 = _compr_18;
              if (((new BigNumber(-127)).isLessThanOrEqualTo(_365_v18)) && ((_365_v18).isLessThan(new BigNumber(39)))) {
                _coll18.push([(_365_v18).minus((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]),_356_i2]);
              }
            }
            return _coll18;
          }(), _347_v2, _361_v14, globalState);
          let _index43 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_364_v17).length));
          (_364_v17)[_index43] = (_this).f21;
          let _366_v19;
          _366_v19 = _module.D4.create_DC9(_364_v17);
          let _367_v20;
          _367_v20 = _dafny.Map.Empty.slice().updateUnsafe((_364_v17)[_module.__default.safeIndex(new BigNumber(120), new BigNumber((_364_v17).length))],(_366_v19).dtor_cf14);
          let _index44 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
          (_348_v3)[_index44] = (_this).fm1(_356_i2, new BigNumber((_367_v20).length), p0, globalState);
          (globalState).f17 = (new BigNumber((_347_v2).length)).plus((_this).f20);
          let _index45 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_364_v17).length));
          let _rhs26 = _dafny.Seq.contains(_dafny.Seq.of(new BigNumber(297), _356_i2), _356_i2);
          let _rhs27 = (_dafny.ZERO).minus((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
          let _lhs28 = _364_v17;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_364_v17).length));
          let _lhs30 = globalState;
          _lhs28[_lhs29] = _rhs26;
          _lhs30.f5 = _rhs27;
          let _368_v21;
          _368_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mqpn"),((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]).multipliedBy(_356_i2));
          let _369_v22;
          _369_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_347_v2).length),(_this).f21);
          let _370_v23;
          _370_v23 = _module.D9.create_DC23(_356_i2, p0);
          _368_v21 = (_368_v21).update(_module.__default.fm26(!((((_369_v22).contains((_370_v23).dtor_cf31)) ? ((_369_v22).get((_370_v23).dtor_cf31)) : (true))), globalState), (_356_i2).minus(_356_i2));
        } else {
          _351_v5 = _351_v5;
          let _index46 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
          (_348_v3)[_index46] = (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))];
          let _371_v24;
          let _nw52 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
          _371_v24 = _nw52;
          let _rhs28 = _371_v24;
          let _rhs29 = _module.__default.safeDivisionInt((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_dafny.ZERO).minus(((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]).plus(_356_i2)));
          let _rhs30 = (_dafny.ZERO).minus((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
          let _rhs31 = (((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]).plus(new BigNumber(193))).isEqualTo((_dafny.ZERO).minus((_this).f20));
          let _lhs31 = globalState;
          let _lhs32 = globalState;
          _371_v24 = _rhs28;
          r3 = _rhs29;
          _lhs31.f17 = _rhs30;
          _lhs32.f4 = _rhs31;
          let _372_v25;
          _372_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_362_v15);
          (globalState).f5 = ((_dafny.ZERO).minus((_this).f20)).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,_362_v15)).Merge(_372_v25)).length));
          let _373_v26;
          let _nw53 = Array((_dafny.ONE).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = (_this).fm0(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f20),(_this).fm1((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], p0, globalState)), _dafny.Seq.UnicodeFromString("vbcdjta"), _360_v13, globalState);
          _373_v26 = _nw53;
          let _374_v27;
          _374_v27 = _module.D8.create_DC20();
          let _rhs32 = (_this).f21;
          let _rhs33 = _373_v26;
          let _rhs34 = _module.D8.create_DC20();
          let _lhs33 = globalState;
          _lhs33.f3 = _rhs32;
          _373_v26 = _rhs33;
          _374_v27 = _rhs34;
        }
        let _rhs35 = (_this).f21;
        let _rhs36 = new BigNumber(948);
        let _lhs34 = globalState;
        let _lhs35 = globalState;
        _lhs34.f4 = _rhs35;
        _lhs35.f5 = _rhs36;
        let _375_v28;
        _375_v28 = _dafny.MultiSet.fromElements(!(p0));
        let _376_v29;
        let _nw54 = new _module.C1();
        _nw54.__ctor((_this).f20, _348_v3, _module.__default.safeModuloInt((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_this).f20), (_module.__default.fm23(globalState)).IsSubsetOf(_375_v28));
        _376_v29 = _nw54;
      }
      if ((_this).f21) {
        r0 = (_this).f20;
        let _377_v30;
        _377_v30 = _dafny.Map.Empty.slice().updateUnsafe(_350_v4,(_this).f21);
        let _378_v31;
        _378_v31 = _dafny.Map.Empty.slice().updateUnsafe(_377_v30,(_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
        _378_v31 = (_378_v31).update(_377_v30, (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
        let _379_v32;
        _379_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
        let _380_v33;
        _380_v33 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
        let _381_v34;
        _381_v34 = _dafny.Set.fromElements(p0, p0, (_this).f21, (_this).f21);
        r1 = (_381_v34).IsProperSubsetOf((_dafny.Set.fromElements((_this).fm0(_379_v32, _347_v2, _dafny.Seq.update(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,p0)), _module.__default.safeIndex((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,p0))).length)), _380_v33), globalState), p0, p0)).Union(_dafny.Set.fromElements(p0)));
        let _382_v35;
        let _nw55 = Array((new BigNumber(10)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = _350_v4;
        _nw55[(_dafny.ONE).toNumber()] = ((p0) ? (_350_v4) : (new _dafny.CodePoint('t'.codePointAt(0))));
        _nw55[(new BigNumber(2)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(3)).toNumber()] = _module.__default.fm17(p0, globalState);
        _nw55[(new BigNumber(4)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(5)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(6)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(7)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(8)).toNumber()] = _350_v4;
        _nw55[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _382_v35 = _nw55;
        let _383_v36;
        _383_v36 = _dafny.MultiSet.fromElements(true, false, (_this).f21);
        let _384_v37;
        _384_v37 = _module.D3.create_DC6(new BigNumber((_383_v36).cardinality()));
        let _385_v38;
        _385_v38 = _dafny.MultiSet.fromElements((_dafny.Set.fromElements((_this).f21, (_this).f21, (_this).f21)).Difference(_381_v34));
        let _index47 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        let _rhs37 = _382_v35;
        let _rhs38 = _384_v37;
        let _rhs39 = (_this).f20;
        let _rhs40 = ((((_this).f20).isLessThanOrEqualTo((_this).f20)) ? (new BigNumber((_347_v2).length)) : ((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]));
        let _rhs41 = _385_v38;
        let _lhs36 = _348_v3;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        _382_v35 = _rhs37;
        _384_v37 = _rhs38;
        r0 = _rhs39;
        _lhs36[_lhs37] = _rhs40;
        _385_v38 = _rhs41;
        let _386_v39;
        _386_v39 = _dafny.Seq.of(_380_v33);
        (globalState).f3 = (_this).fm0(_379_v32, _module.__default.fm26(p0, globalState), _dafny.Seq.Concat(_386_v39, _386_v39), globalState);
      } else {
        (globalState).f14 = (_this).f21;
        (globalState).f5 = (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))];
        (globalState).f17 = (_this).f20;
        let _387_v40;
        _387_v40 = _dafny.Map.Empty.slice().updateUnsafe(_347_v2,(_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
        let _index48 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        (_348_v3)[_index48] = _module.__default.safeModuloInt(new BigNumber((_387_v40).length), (_this).f20);
        let _388_v41;
        _388_v41 = p0;
        _388_v41 = _388_v41;
      }
      let _389_v42;
      _389_v42 = _dafny.Map.Empty.slice().updateUnsafe((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))],(_this).f20);
      let _390_v43;
      _390_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _391_v44;
      _391_v44 = _dafny.Seq.of(_390_v43);
      let _392_v45;
      _392_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).fm0(_389_v42, _347_v2, _391_v44, globalState));
      let _393_v46;
      _393_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_390_v43);
      if (!(((_390_v43).Merge(_392_v45)).equals(((((_393_v46).contains(new BigNumber(240))) ? ((_393_v46).get(new BigNumber(240))) : (_module.__default.fm19((_this).f21, globalState)))).Merge(_392_v45)))) {
        let _394_v47;
        _394_v47 = _dafny.Set.fromElements(_350_v4, _350_v4);
        let _395_v48;
        _395_v48 = _dafny.Seq.of(_351_v5, _351_v5);
        _392_v45 = (_392_v45).update((new BigNumber((_394_v47).length)).isLessThan(new BigNumber((_395_v48).length)), true);
        let _396_v49;
        let _init6 = function (_397_i3) {
          return (_this).f21;
        };
        let _nw56 = Array((new BigNumber(18)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw56.length); _i0_6++) {
          _nw56[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _396_v49 = _nw56;
        let _398_v50;
        _398_v50 = _module.D4.create_DC9(_396_v49);
        let _pat_let_tv10 = _396_v49;
        let _399_v51;
        _399_v51 = _dafny.MultiSet.fromElements(function (_pat_let13_0) {
          return function (_400_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_401_dt__update_hcf14_h0) {
                return _module.D4.create_DC9(_401_dt__update_hcf14_h0);
              }(_pat_let14_0);
            }(_pat_let_tv10);
          }(_pat_let13_0);
        }(_398_v50), _398_v50, _module.D4.create_DC9(_396_v49), _398_v50);
        let _402_v52;
        _402_v52 = _dafny.MultiSet.fromElements((_this).f20, (_dafny.ZERO).minus((_this).f20), (((_399_v51).contains(_398_v50)) ? ((_399_v51).get(_398_v50)) : ((_dafny.ZERO).minus((_this).f20))), new BigNumber(969));
        let _403_v53;
        _403_v53 = _dafny.Map.Empty.slice().updateUnsafe(_402_v52,_355_v9);
        let _pat_let_tv11 = _351_v5;
        _403_v53 = (_403_v53).update(_dafny.MultiSet.fromElements((_this).f20, (_this).f20, new BigNumber(-398), (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]), function (_pat_let15_0) {
          return function (_404_dt__update__tmp_h1) {
            return function (_pat_let16_0) {
              return function (_405_dt__update_hcf26_h0) {
                return _module.D7.create_DC18(_405_dt__update_hcf26_h0, (_404_dt__update__tmp_h1).dtor_cf27);
              }(_pat_let16_0);
            }(_pat_let_tv11);
          }(_pat_let15_0);
        }(_module.D7.create_DC18(_351_v5, (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))])));
        let _406_v54;
        _406_v54 = _dafny.MultiSet.fromElements((_this).f21);
        let _407_v55;
        _407_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_406_v54).cardinality()),_348_v3);
        let _408_v56;
        _408_v56 = _dafny.Set.fromElements((_this).f21);
        let _409_v57;
        let _nw57 = new _module.C1();
        _nw57.__ctor((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (((_407_v55).contains((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))])) ? ((_407_v55).get((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))])) : (_348_v3)), _module.__default.safeDivisionInt((_this).f20, (_dafny.ZERO).minus(new BigNumber((_408_v56).length))), !((_this).f21));
        _409_v57 = _nw57;
        let _rhs42 = new BigNumber(632);
        let _rhs43 = _409_v57;
        r3 = _rhs42;
        _409_v57 = _rhs43;
        let _410_v58;
        _410_v58 = _dafny.Map.Empty.slice().updateUnsafe(_396_v49,(_dafny.ZERO).minus((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]));
        _410_v58 = (_410_v58).update(_396_v49, (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
        let _index49 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length));
        (_348_v3)[_index49] = (_dafny.ZERO).minus(_409_v57.f27);
      } else {
        let _411_v59;
        let _init7 = ((_412_p0) => function (_413_i4) {
          return _412_p0;
        })(p0);
        let _nw58 = Array((new BigNumber(9)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw58.length); _i0_7++) {
          _nw58[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _411_v59 = _nw58;
        let _414_v60;
        _414_v60 = _module.D4.create_DC9(_411_v59);
        let _415_v61;
        _415_v61 = _dafny.Seq.of(!(!((_this).f20).isEqualTo((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))])), (_this).f21);
        let _rhs44 = _414_v60;
        let _rhs45 = _415_v61;
        let _rhs46 = (_this).f21;
        let _rhs47 = (_this).f20;
        let _lhs38 = globalState;
        _414_v60 = _rhs44;
        _415_v61 = _rhs45;
        r1 = _rhs46;
        _lhs38.f17 = _rhs47;
        let _416_v62;
        _416_v62 = _dafny.Map.Empty.slice().updateUnsafe((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))],p0);
        _389_v42 = (_389_v42).update(((_this).f20).multipliedBy(new BigNumber((_416_v62).length)), (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]);
        let _417_v63;
        _417_v63 = _dafny.Set.fromElements((_this).f21, (_this).f21);
        let _418_v64;
        _418_v64 = _dafny.MultiSet.fromElements(_417_v63);
        let _index50 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_411_v59).length));
        (_411_v59)[_index50] = (new BigNumber((_418_v64).cardinality())).isLessThanOrEqualTo((_this).f20);
        let _index51 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_411_v59).length));
        (_411_v59)[_index51] = ((_this).f21) && (p0);
        (globalState).f4 = (_411_v59)[_module.__default.safeIndex(new BigNumber(228), new BigNumber((_411_v59).length))];
        let _419_v65;
        _419_v65 = _dafny.Seq.of(_dafny.Seq.of((_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))], (_this).f20, (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))]), _351_v5);
        let _420_v66;
        _420_v66 = _dafny.Seq.of(_dafny.Seq.Concat(_351_v5, (_419_v65)[_module.__default.safeIndex((_this).f20, new BigNumber((_419_v65).length))]), _351_v5);
        _419_v65 = _420_v66;
      }
      r0 = (_348_v3)[_module.__default.safeIndex(new BigNumber(768), new BigNumber((_348_v3).length))];
      r1 = (_this).fm0(_389_v42, _347_v2, _391_v44, globalState);
      r2 = new BigNumber(-558);
      let _421_v67;
      _421_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_351_v5).length),p0);
      let _422_v68;
      _422_v68 = _dafny.Map.Empty.slice().updateUnsafe(_421_v67,p0);
      r3 = _module.__default.safeModuloInt(new BigNumber((_422_v68).length), (_this).f20);
      return [r0, r1, r2, r3];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _hi3 = (_this).f20;
      for (let _423_i0 = (_dafny.ZERO).minus(p0); _423_i0.isLessThan(_hi3); _423_i0 = _423_i0.plus(_dafny.ONE)) {
        let _424_v0;
        let _nw59 = Array((new BigNumber(19)).toNumber()).fill([]);
        _424_v0 = _nw59;
        let _425_v1;
        let _nw60 = Array((new BigNumber(26)).toNumber()).fill(false);
        _425_v1 = _nw60;
        let _426_v2;
        _426_v2 = _module.D3.create_DC6(_423_i0);
        let _427_v3;
        _427_v3 = _dafny.Seq.UnicodeFromString("qvejr");
        let _rhs48 = _dafny.Seq.Concat(_dafny.Seq.update(_427_v3, _module.__default.safeIndex(new BigNumber(-478), new BigNumber((_427_v3).length)), new _dafny.CodePoint('r'.codePointAt(0))), _427_v3);
        let _rhs49 = _424_v0;
        let _rhs50 = _425_v1;
        let _rhs51 = _module.D3.create_DC6(_423_i0);
        let _rhs52 = p0;
        let _lhs39 = globalState;
        let _lhs40 = globalState;
        _lhs39.f15 = _rhs48;
        _424_v0 = _rhs49;
        _425_v1 = _rhs50;
        _426_v2 = _rhs51;
        _lhs40.f5 = _rhs52;
        let _index52 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length));
        (_424_v0)[_index52] = _425_v1;
        let _index53 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length));
        (_424_v0)[_index53] = _425_v1;
        let _index54 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_425_v1).length));
        (_425_v1)[_index54] = p1;
        let _arr0 = (_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))];
        let _index55 = _module.__default.safeIndex(new BigNumber(969), new BigNumber(((_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))]).length));
        _arr0[_index55] = (_this).f21;
        let _index56 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_425_v1).length));
        let _arr1 = (_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))];
        let _index57 = _module.__default.safeIndex(new BigNumber(969), new BigNumber(((_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))]).length));
        let _rhs53 = (_this).f21;
        let _rhs54 = p1;
        let _lhs41 = _425_v1;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_425_v1).length));
        let _lhs43 = (_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))];
        let _lhs44 = _module.__default.safeIndex(new BigNumber(969), new BigNumber(((_424_v0)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_424_v0).length))]).length));
        _lhs41[_lhs42] = _rhs53;
        _lhs43[_lhs44] = _rhs54;
        let _428_v4;
        _428_v4 = _dafny.Set.fromElements((_425_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_425_v1).length))]);
        let _429_v5;
        let _nw61 = Array((new BigNumber(6)).toNumber());
        _nw61[(_dafny.ZERO).toNumber()] = _423_i0;
        _nw61[(_dafny.ONE).toNumber()] = (_this).f20;
        _nw61[(new BigNumber(2)).toNumber()] = new BigNumber(-412);
        _nw61[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_423_i0);
        _nw61[(new BigNumber(4)).toNumber()] = _423_i0;
        _nw61[(new BigNumber(5)).toNumber()] = new BigNumber((_428_v4).length);
        _429_v5 = _nw61;
        let _430_v6;
        let _nw62 = new _module.C1();
        _nw62.__ctor(p0, _429_v5, _423_i0, p1);
        _430_v6 = _nw62;
      }
      let _431_i1;
      _431_i1 = _dafny.ZERO;
      L2: {
        while ((_this).f21) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_431_i1)) {
              break L2;
            }
            _431_i1 = (_431_i1).plus(_dafny.ONE);
            let _432_v7;
            _432_v7 = _dafny.Seq.of(p0);
            let _433_v8;
            _433_v8 = _dafny.Seq.UnicodeFromString("ogxg");
            let _434_v9;
            _434_v9 = _dafny.Seq.of((_this).f21, p1);
            let _435_v10;
            let _nw63 = Array((new BigNumber(24)).toNumber());
            _nw63[(_dafny.ZERO).toNumber()] = (_432_v7)[_module.__default.safeIndex(p0, new BigNumber((_432_v7).length))];
            _nw63[(_dafny.ONE).toNumber()] = new BigNumber((_433_v8).length);
            _nw63[(new BigNumber(2)).toNumber()] = (_this).f20;
            _nw63[(new BigNumber(3)).toNumber()] = (_this).f20;
            _nw63[(new BigNumber(4)).toNumber()] = (_432_v7)[_module.__default.safeIndex((_this).f20, new BigNumber((_432_v7).length))];
            _nw63[(new BigNumber(5)).toNumber()] = (_this).f20;
            _nw63[(new BigNumber(6)).toNumber()] = p0;
            _nw63[(new BigNumber(7)).toNumber()] = new BigNumber((_433_v8).length);
            _nw63[(new BigNumber(8)).toNumber()] = new BigNumber(551);
            _nw63[(new BigNumber(9)).toNumber()] = p0;
            _nw63[(new BigNumber(10)).toNumber()] = (_this).fm1(new BigNumber(-469), (_this).f20, (_this).f21, globalState);
            _nw63[(new BigNumber(11)).toNumber()] = new BigNumber(158);
            _nw63[(new BigNumber(12)).toNumber()] = p0;
            _nw63[(new BigNumber(13)).toNumber()] = p0;
            _nw63[(new BigNumber(14)).toNumber()] = p0;
            _nw63[(new BigNumber(15)).toNumber()] = (_this).f20;
            _nw63[(new BigNumber(16)).toNumber()] = new BigNumber((_433_v8).length);
            _nw63[(new BigNumber(17)).toNumber()] = new BigNumber((_434_v9).length);
            _nw63[(new BigNumber(18)).toNumber()] = p0;
            _nw63[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Set.fromElements(true, p1, (_this).f21, (_this).f21)).length);
            _nw63[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_432_v7,p1)).length);
            _nw63[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((_this).fm1(new BigNumber((_433_v8).length), (_this).f20, p1, globalState));
            _nw63[(new BigNumber(22)).toNumber()] = p0;
            _nw63[(new BigNumber(23)).toNumber()] = p0;
            _435_v10 = _nw63;
            let _436_v11;
            let _nw64 = new _module.C1();
            _nw64.__ctor(new BigNumber(-261), _435_v10, new BigNumber(556), (_this).f21);
            _436_v11 = _nw64;
            (globalState).f14 = !(!(p1));
            if ((_434_v9)[_module.__default.safeIndex(_436_v11.f27, new BigNumber((_434_v9).length))]) {
              r1 = false;
              let _437_v12;
              _437_v12 = new _dafny.CodePoint('w'.codePointAt(0));
              _437_v12 = new _dafny.CodePoint('i'.codePointAt(0));
              (globalState).f14 = (p0).isEqualTo(p0);
              let _438_v13;
              let _init8 = ((_439_v8) => function (_440_i2) {
                return _439_v8;
              })(_433_v8);
              let _nw65 = Array((_dafny.ONE).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
                _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _438_v13 = _nw65;
              let _index58 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_438_v13).length));
              (_438_v13)[_index58] = _433_v8;
              let _index59 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_438_v13).length));
              (_438_v13)[_index59] = _433_v8;
              (globalState).f17 = (_436_v11.f27).plus(_436_v11.f27);
            } else {
              let _441_v14;
              let _442_v15;
              let _443_v16;
              let _444_v17;
              let _out6;
              let _out7;
              let _out8;
              let _out9;
              let _outcollector2 = (_this).m3(((p1) ? ((_this).f21) : ((_this).f21)), globalState);
              _out6 = _outcollector2[0];
              _out7 = _outcollector2[1];
              _out8 = _outcollector2[2];
              _out9 = _outcollector2[3];
              _441_v14 = _out6;
              _442_v15 = _out7;
              _443_v16 = _out8;
              _444_v17 = _out9;
              let _445_v18;
              _445_v18 = _dafny.Set.fromElements(_433_v8, _dafny.Seq.Concat(_433_v8, _433_v8), _433_v8, _433_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), function (_446_i3) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              }));
              _445_v18 = _445_v18;
              (globalState).f5 = _443_v16;
              let _447_v19;
              _447_v19 = _dafny.MultiSet.fromElements(p1, (_this).f21, _442_v15);
              let _448_v20;
              _448_v20 = _dafny.Seq.of(_436_v11.f27, _443_v16);
              let _449_v21;
              let _nw66 = new _module.C1();
              _nw66.__ctor(_436_v11.f27, (_436_v11).f28, (_436_v11).fm1(new BigNumber((_434_v9).length), _436_v11.f27, false, globalState), (_436_v11).fm16(_447_v19, _448_v20, globalState));
              _449_v21 = _nw66;
              let _450_v22;
              _450_v22 = _module.D7.create_DC17(!(p1));
              (globalState).f4 = (_450_v22).dtor_cf25;
            }
            if (!(!(_436_v11.f27).isEqualTo(p0)) || (true)) {
              let _index60 = _module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length));
              ((_436_v11).f28)[_index60] = _436_v11.f27;
              let _index61 = _module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length));
              ((_436_v11).f28)[_index61] = _436_v11.f27;
              let _451_v23;
              _451_v23 = _dafny.Set.fromElements(p0, p0, _436_v11.f27, _436_v11.f27, p0);
              let _452_v24;
              let _453_v25;
              let _454_v26;
              let _455_v27;
              let _out10;
              let _out11;
              let _out12;
              let _out13;
              let _outcollector3 = (_this).m3((new BigNumber((_451_v23).length)).isLessThan(_436_v11.f27), globalState);
              _out10 = _outcollector3[0];
              _out11 = _outcollector3[1];
              _out12 = _outcollector3[2];
              _out13 = _outcollector3[3];
              _452_v24 = _out10;
              _453_v25 = _out11;
              _454_v26 = _out12;
              _455_v27 = _out13;
              let _456_v28;
              let _nw67 = Array((new BigNumber(29)).toNumber()).fill(false);
              _456_v28 = _nw67;
              let _457_v29;
              _457_v29 = _dafny.MultiSet.fromElements(false);
              let _index62 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_456_v28).length));
              (_456_v28)[_index62] = (_457_v29).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f21, _453_v25));
              let _index63 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_456_v28).length));
              (_456_v28)[_index63] = (((_436_v11).f28)[_module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length))]).isLessThan(_452_v24);
              let _458_v30;
              let _nw68 = new _module.C1();
              _nw68.__ctor(_436_v11.f27, (_436_v11).f28, _module.__default.safeDivisionInt((_dafny.ZERO).minus(((_436_v11).f28)[_module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length))]), _436_v11.f27), (((_436_v11).f28)[_module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length))]).isEqualTo(new BigNumber(992)));
              _458_v30 = _nw68;
              let _459_v31;
              _459_v31 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt((_436_v11).fm1(_454_v26, new BigNumber((_451_v23).length), true, globalState), (_dafny.ZERO).minus(new BigNumber((_433_v8).length))), new BigNumber((_433_v8).length));
              let _460_v32;
              _460_v32 = _module.D10.create_DC25(_459_v31);
              let _index64 = _module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length));
              let _rhs55 = (_this).f21;
              let _rhs56 = ((_460_v32).dtor_cf34).Intersect(_459_v31);
              let _rhs57 = (_dafny.ZERO).minus(_452_v24);
              let _rhs58 = p1;
              let _lhs45 = globalState;
              let _lhs46 = (_436_v11).f28;
              let _lhs47 = _module.__default.safeIndex(new BigNumber(851), new BigNumber(((_436_v11).f28).length));
              let _lhs48 = globalState;
              _lhs45.f18 = _rhs55;
              _459_v31 = _rhs56;
              _lhs46[_lhs47] = _rhs57;
              _lhs48.f18 = _rhs58;
            } else {
              (_436_v11).f27 = new BigNumber((_434_v9).length);
              let _461_v33;
              _461_v33 = !(_dafny.Seq.IsPrefixOf(_433_v8, _433_v8));
              _461_v33 = _461_v33;
              let _462_v34;
              _462_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(150),p0);
              let _463_v35;
              _463_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f21);
              let _464_v36;
              _464_v36 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21), _463_v35, _463_v35);
              (globalState).f3 = (((_this).f21) ? (!((_this).fm0(_462_v34, _433_v8, _464_v36, globalState))) : ((_this).f21));
              let _465_v37;
              _465_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_436_v11.f27);
              _465_v37 = ((_module.__default.fm27((_this).f20, globalState)).Merge(_465_v37)).Merge(p2);
              let _466_v38;
              let _nw69 = new _module.C1();
              _nw69.__ctor((_436_v11.f27).multipliedBy(new BigNumber(687)), (_436_v11).f28, p0, (_this).f21);
              _466_v38 = _nw69;
            }
          }
        }
      }
      r1 = ((!(!((_this).f21))) ? ((_this).f21) : (p1));
      let _467_v39;
      _467_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f21);
      let _468_v40;
      _468_v40 = new _dafny.CodePoint('k'.codePointAt(0));
      let _469_v41;
      _469_v41 = _dafny.Seq.UnicodeFromString("xuposfy");
      let _470_v42;
      _470_v42 = _dafny.MultiSet.fromElements(p0);
      let _471_v43;
      _471_v43 = _dafny.Seq.of(p0);
      let _472_v44;
      _472_v44 = _dafny.Set.fromElements(new BigNumber(862), p0);
      let _473_v45;
      _473_v45 = _dafny.Map.Empty.slice().updateUnsafe(_468_v40,new BigNumber((_472_v44).length));
      let _474_v46;
      let _init9 = function (_475_i4) {
        return (_475_i4).plus((_this).f20);
      };
      let _nw70 = Array((new BigNumber(10)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw70.length); _i0_9++) {
        _nw70[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _474_v46 = _nw70;
      let _476_v47;
      let _nw71 = new _module.C1();
      _nw71.__ctor(new BigNumber((_473_v45).length), _474_v46, p0, p1);
      _476_v47 = _nw71;
      let _477_v48;
      _477_v48 = _dafny.Set.fromElements(_476_v47, _476_v47);
      let _478_v49;
      _478_v49 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f21);
      let _479_v50;
      let _nw72 = Array((new BigNumber(28)).toNumber());
      _nw72[(_dafny.ZERO).toNumber()] = new BigNumber((_469_v41).length);
      _nw72[(_dafny.ONE).toNumber()] = p0;
      _nw72[(new BigNumber(2)).toNumber()] = new BigNumber((_470_v42).cardinality());
      _nw72[(new BigNumber(3)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(4)).toNumber()] = (_471_v43)[_module.__default.safeIndex(p0, new BigNumber((_471_v43).length))];
      _nw72[(new BigNumber(5)).toNumber()] = p0;
      _nw72[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(7)).toNumber()] = p0;
      _nw72[(new BigNumber(8)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(9)).toNumber()] = new BigNumber((_467_v39).length);
      _nw72[(new BigNumber(10)).toNumber()] = p0;
      _nw72[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_477_v48).length))).length);
      _nw72[(new BigNumber(12)).toNumber()] = p0;
      _nw72[(new BigNumber(13)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(14)).toNumber()] = new BigNumber((_478_v49).length);
      _nw72[(new BigNumber(15)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(16)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(17)).toNumber()] = new BigNumber((_469_v41).length);
      _nw72[(new BigNumber(18)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(19)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(20)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(21)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(22)).toNumber()] = _476_v47.f27;
      _nw72[(new BigNumber(23)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(24)).toNumber()] = p0;
      _nw72[(new BigNumber(25)).toNumber()] = (_this).f20;
      _nw72[(new BigNumber(26)).toNumber()] = new BigNumber((_473_v45).length);
      _nw72[(new BigNumber(27)).toNumber()] = new BigNumber((_471_v43).length);
      _479_v50 = _nw72;
      let _480_v51;
      let _nw73 = new _module.C1();
      _nw73.__ctor((_module.D10.create_DC26(new BigNumber(545))).dtor_cf35, _474_v46, p0, p1);
      _480_v51 = _nw73;
      let _481_v52;
      _481_v52 = _dafny.Map.Empty.slice().updateUnsafe(_468_v40,_476_v47);
      let _482_v53;
      _482_v53 = _dafny.Seq.of(_481_v52, _481_v52, _481_v52);
      let _483_v54;
      _483_v54 = _module.D10.create_DC27(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_467_v39).length)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(new BigNumber((_467_v39).length))).length)), p0))).cardinality()), _482_v53, _module.__default.fm17((_this).f21, globalState));
      let _484_v55;
      _484_v55 = _471_v43;
      let _pat_let_tv12 = _480_v51;
      let _pat_let_tv13 = p1;
      let _pat_let_tv14 = _484_v55;
      let _pat_let_tv15 = globalState;
      let _pat_let_tv16 = globalState;
      let _source7 = function (_pat_let17_0) {
        return function (_485_dt__update__tmp_h0) {
          return function (_pat_let18_0) {
            return function (_486_dt__update_hcf36_h0) {
              return function (_pat_let19_0) {
                return function (_487_dt__update_hcf38_h0) {
                  return _module.D10.create_DC27(_486_dt__update_hcf36_h0, (_485_dt__update__tmp_h0).dtor_cf37, _487_dt__update_hcf38_h0);
                }(_pat_let19_0);
              }(_module.__default.fm17((_pat_let_tv12).fm16(_dafny.MultiSet.fromElements(true, !(_pat_let_tv13)), _pat_let_tv14, _pat_let_tv15), _pat_let_tv16));
            }(_pat_let18_0);
          }((_this).f20);
        }(_pat_let17_0);
      }(_483_v54);
      if (_source7.is_DC26) {
        let _488___mcc_h0 = (_source7).cf35;
        let _489_cf35 = _488___mcc_h0;
        (globalState).f18 = (new BigNumber(350)).isLessThanOrEqualTo(p0);
        let _490_v56;
        let _nw74 = Array((new BigNumber(12)).toNumber()).fill(false);
        _490_v56 = _nw74;
        let _491_v57;
        _491_v57 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_476_v47.f27),_476_v47.f27);
        let _492_v58;
        _492_v58 = _dafny.Seq.of(_478_v49);
        let _index65 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length));
        (_490_v56)[_index65] = (_this).fm0(_491_v57, _469_v41, _dafny.Seq.update(_492_v58, _module.__default.safeIndex(_476_v47.f27, new BigNumber((_492_v58).length)), _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1)), globalState);
        let _493_v59;
        _493_v59 = _dafny.Set.fromElements((_this).f21, p1);
        let _index66 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length));
        ((_480_v51).f28)[_index66] = (_module.__default.fm24(_476_v47.f27, !(p1), _493_v59, p1, globalState)).dtor_cf17;
        let _index67 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_490_v56).length));
        (_490_v56)[_index67] = p1;
        let _index68 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length));
        (_474_v46)[_index68] = _480_v51.f27;
        let _494_v60;
        _494_v60 = _dafny.MultiSet.fromElements(_490_v56, _490_v56);
        let _495_v61;
        _495_v61 = _dafny.Map.Empty.slice().updateUnsafe(_472_v44,_480_v51);
        let _index69 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length));
        let _index70 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length));
        let _index71 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_490_v56).length));
        let _index72 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length));
        let _rhs59 = (_this).f21;
        let _rhs60 = _489_cf35;
        let _rhs61 = !(_495_v61).equals(_495_v61);
        let _rhs62 = _480_v51.f27;
        let _rhs63 = ((_494_v60).Union(_494_v60)).Intersect(_494_v60);
        let _lhs49 = _490_v56;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length));
        let _lhs51 = (_480_v51).f28;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length));
        let _lhs53 = _490_v56;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_490_v56).length));
        let _lhs55 = _474_v46;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length));
        _lhs49[_lhs50] = _rhs59;
        _lhs51[_lhs52] = _rhs60;
        _lhs53[_lhs54] = _rhs61;
        _lhs55[_lhs56] = _rhs62;
        _494_v60 = _rhs63;
        let _496_v62;
        let _nw75 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _496_v62 = _nw75;
        let _index73 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length));
        (_496_v62)[_index73] = _469_v41;
        let _index74 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length));
        let _index75 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length));
        let _rhs64 = _469_v41;
        let _rhs65 = _module.__default.safeModuloInt(_476_v47.f27, _476_v47.f27);
        let _rhs66 = !((_this).f21);
        let _lhs57 = _496_v62;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length));
        let _lhs59 = (_480_v51).f28;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length));
        let _lhs61 = globalState;
        _lhs57[_lhs58] = _rhs64;
        _lhs59[_lhs60] = _rhs65;
        _lhs61.f3 = _rhs66;
        if ((((_478_v49).contains((_490_v56)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length))])) ? ((_478_v49).get((_490_v56)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length))])) : (p1))) {
          let _497_v64;
          let _init10 = ((_498_v44, _499_p2, _500_v57) => function (_501_i5) {
            return _dafny.Seq.of(_498_v44, function () {
              let _coll19 = new _dafny.Set();
              for (const _compr_19 of (_500_v57).Keys.Elements) {
                let _502_v63 = _compr_19;
                if ((_500_v57).contains(_502_v63)) {
                  _coll19.add((_502_v63).minus(new BigNumber((_dafny.Seq.UnicodeFromString("afyfapfi")).length)));
                }
              }
              return _coll19;
            }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ml")).length), (((_499_p2).contains((_this).f21)) ? ((_499_p2).get((_this).f21)) : ((_this).f20))));
          })(_472_v44, p2, _491_v57);
          let _nw76 = Array((new BigNumber(8)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw76.length); _i0_10++) {
            _nw76[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _497_v64 = _nw76;
          let _503_v65;
          _503_v65 = _dafny.MultiSet.fromElements(p1, true);
          let _index76 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_497_v64).length));
          (_497_v64)[_index76] = _dafny.Seq.of(_dafny.Set.fromElements((((_503_v65).contains((_this).f21)) ? ((_503_v65).get((_this).f21)) : (new BigNumber(684))), (_480_v51).fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-418)), ((_504_cf35) => function (_505_i6) {
            return _504_cf35;
          })(_489_cf35))).length), new BigNumber(-357), (_490_v56)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length))], globalState)), _472_v44);
          let _506_v66;
          _506_v66 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((p2).length), (_dafny.ZERO).minus(_489_cf35), ((_480_v51).f28)[_module.__default.safeIndex(new BigNumber(254), new BigNumber(((_480_v51).f28).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_507_v40) => function (_508_i7) {
            return _507_v40;
          })(_468_v40))).length)));
          let _509_v67;
          _509_v67 = _dafny.Seq.of(_506_v66, _506_v66, _506_v66, _dafny.Seq.of((_506_v66)[_module.__default.safeIndex(_480_v51.f27, new BigNumber((_506_v66).length))]), _506_v66);
          let _index77 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_497_v64).length));
          (_497_v64)[_index77] = _dafny.Seq.Concat(_dafny.Seq.update((_509_v67)[_module.__default.safeIndex(_480_v51.f27, new BigNumber((_509_v67).length))], _module.__default.safeIndex(_476_v47.f27, new BigNumber(((_509_v67)[_module.__default.safeIndex(_480_v51.f27, new BigNumber((_509_v67).length))]).length)), _dafny.Set.fromElements(new BigNumber(((_496_v62)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length))]).length))), _506_v66);
          let _510_v68;
          _510_v68 = _dafny.Map.Empty.slice().updateUnsafe(_490_v56,(_489_cf35).plus(new BigNumber((_472_v44).length)));
          _510_v68 = (_510_v68).update(_490_v56, _476_v47.f27);
          let _index78 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length));
          (_496_v62)[_index78] = _469_v41;
          let _511_v69;
          _511_v69 = _dafny.MultiSet.fromElements(_module.__default.fm17(p1, globalState));
          (globalState).f5 = (_476_v47.f27).minus((new BigNumber((_467_v39).length)).multipliedBy((((_511_v69).contains(_468_v40)) ? ((_511_v69).get(_468_v40)) : ((_this).f20))));
          let _512_v70;
          _512_v70 = _dafny.Seq.of(p1, !((_480_v51.f27).isLessThan(_480_v51.f27)), (_503_v65).IsSubsetOf(_503_v65));
          (globalState).f18 = (_512_v70)[_module.__default.safeIndex(new BigNumber(-639), new BigNumber((_512_v70).length))];
        } else {
          (globalState).f5 = (_471_v43)[_module.__default.safeIndex(((_474_v46)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length))]).minus(_489_cf35), new BigNumber((_471_v43).length))];
          let _513_v71;
          _513_v71 = _dafny.Seq.of((_476_v47).f28);
          let _index79 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length));
          (_490_v56)[_index79] = (new BigNumber((_513_v71).length)).isEqualTo(_480_v51.f27);
          let _index80 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length));
          (_474_v46)[_index80] = (_480_v51).fm1(_480_v51.f27, _489_cf35, !(p1), globalState);
          let _514_v72;
          _514_v72 = _module.D9.create_DC23((_474_v46)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_474_v46).length))], p1);
          let _515_v73;
          _515_v73 = _module.D9.create_DC24(_514_v72);
          let _516_v74;
          _516_v74 = _module.D9.create_DC24(_515_v73);
          _516_v74 = _516_v74;
          let _517_v75;
          let _nw77 = new _module.C1();
          _nw77.__ctor((_dafny.ZERO).minus(_476_v47.f27), (_480_v51).f28, new BigNumber((_dafny.Seq.Concat((_496_v62)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_496_v62).length))], _dafny.Seq.UnicodeFromString("ien"))).length), (_490_v56)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_490_v56).length))]);
          _517_v75 = _nw77;
        }
      } else if (_source7.is_DC27) {
        let _518___mcc_h1 = (_source7).cf36;
        let _519___mcc_h2 = (_source7).cf37;
        let _520___mcc_h3 = (_source7).cf38;
        let _521_cf38 = _520___mcc_h3;
        let _522_cf37 = _519___mcc_h2;
        let _523_cf36 = _518___mcc_h1;
        (globalState).f18 = p1;
        let _524_v76;
        let _nw78 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _524_v76 = _nw78;
        let _index81 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v76).length));
        (_524_v76)[_index81] = _dafny.Seq.Concat(_dafny.Seq.update(_469_v41, _module.__default.safeIndex(new BigNumber((_469_v41).length), new BigNumber((_469_v41).length)), new _dafny.CodePoint('a'.codePointAt(0))), _469_v41);
        let _525_v77;
        _525_v77 = _module.D6.create_DC13(_469_v41);
        let _index82 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_524_v76).length));
        (_524_v76)[_index82] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ndkkvypt"), _dafny.Seq.UnicodeFromString("dat")), (_525_v77).dtor_cf21);
        let _526_v78;
        _526_v78 = _dafny.Seq.of((_this).f21, (_this).f21, (_this).f21);
        let _527_v79;
        _527_v79 = _dafny.MultiSet.fromElements(!((_526_v78)[_module.__default.safeIndex(new BigNumber((_471_v43).length), new BigNumber((_526_v78).length))]), (_this).f21);
        let _528_v80;
        _528_v80 = _dafny.Seq.of((_480_v51).fm16(_527_v79, _dafny.Seq.of(_523_cf36, _480_v51.f27), globalState), p1);
        let _529_v81;
        _529_v81 = _dafny.Seq.of(_478_v49, (_478_v49).update((_this).f21, p1), ((!((_528_v80)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_528_v80).length))])) ? (_478_v49) : (_478_v49)));
        let _530_v82;
        _530_v82 = _module.D7.create_DC18(_471_v43, _476_v47.f27);
        let _531_v83;
        _531_v83 = _dafny.Map.Empty.slice().updateUnsafe(_530_v82,(true) && (p1));
        let _rhs67 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21), _478_v49), _529_v81), _module.__default.safeIndex((_dafny.ZERO).minus(_476_v47.f27), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21), _478_v49), _529_v81)).length)), _478_v49);
        let _rhs68 = (((_531_v83).contains(_530_v82)) ? ((_531_v83).get(_530_v82)) : (p1));
        let _rhs69 = new BigNumber((_471_v43).length);
        let _lhs62 = globalState;
        let _lhs63 = _476_v47;
        _529_v81 = _rhs67;
        _lhs62.f3 = _rhs68;
        _lhs63.f27 = _rhs69;
        let _532_v84;
        let _init11 = function (_533_i8) {
          return (_this).f21;
        };
        let _nw79 = Array((new BigNumber(17)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw79.length); _i0_11++) {
          _nw79[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _532_v84 = _nw79;
        let _index83 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_532_v84).length));
        (_532_v84)[_index83] = (_this).f21;
        let _534_v85;
        _534_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_476_v47.f27);
        let _index84 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_532_v84).length));
        (_532_v84)[_index84] = (_480_v51).fm0(_534_v85, _469_v41, _529_v81, globalState);
      } else {
        let _535___mcc_h4 = (_source7).cf34;
        let _536_cf34 = _535___mcc_h4;
        let _537_v86;
        let _init12 = function (_538_i9) {
          return (_this).f21;
        };
        let _nw80 = Array((new BigNumber(16)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw80.length); _i0_12++) {
          _nw80[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _537_v86 = _nw80;
        let _539_v87;
        _539_v87 = _dafny.MultiSet.fromElements(_537_v86, _537_v86);
        let _source8 = _module.D2.create_DC2((_dafny.MultiSet.fromElements(_537_v86)).Difference(_539_v87));
        if (_source8.is_DC3) {
          let _540___mcc_h5 = (_source8).cf3;
          let _541___mcc_h6 = (_source8).cf4;
          let _542___mcc_h7 = (_source8).cf5;
          let _543___mcc_h8 = (_source8).cf6;
          let _544___mcc_h9 = (_source8).cf7;
          let _545_cf7 = _544___mcc_h9;
          let _546_cf6 = _543___mcc_h8;
          let _547_cf5 = _542___mcc_h7;
          let _548_cf4 = _541___mcc_h6;
          let _549_cf3 = _540___mcc_h5;
          (globalState).f4 = _548_cf4;
          (_476_v47).f27 = _480_v51.f27;
          _546_cf6 = (_476_v47.f27).plus(_module.__default.safeDivisionInt(_480_v51.f27, _476_v47.f27));
          let _550_v88;
          let _nw81 = new _module.C0();
          _nw81.__ctor();
          _550_v88 = _nw81;
        } else if (_source8.is_DC2) {
          let _551___mcc_h10 = (_source8).cf2;
          let _552_cf2 = _551___mcc_h10;
          let _553_v89;
          _553_v89 = _dafny.Seq.of(p1);
          let _554_v90;
          _554_v90 = _dafny.Seq.of(_478_v49);
          let _555_v91;
          _555_v91 = _dafny.MultiSet.fromElements((_this).fm0(_dafny.Map.Empty.slice().updateUnsafe((_480_v51).fm1(new BigNumber((_dafny.MultiSet.fromElements(_476_v47.f27)).cardinality()), (_this).fm1(new BigNumber((_553_v89).length), _476_v47.f27, (_this).f21, globalState), p1, globalState),new BigNumber((_553_v89).length)), _dafny.Seq.UnicodeFromString("folerbj"), _554_v90, globalState), (_this).f21);
          let _rhs70 = ((_this).f20).multipliedBy((_dafny.ZERO).minus((((_555_v91).contains(p1)) ? ((_555_v91).get(p1)) : (_480_v51.f27))));
          let _rhs71 = p1;
          let _lhs64 = _476_v47;
          _lhs64.f27 = _rhs70;
          r1 = _rhs71;
          _478_v49 = (_478_v49).update(((p1) ? (p1) : (false)), (_this).f21);
          (globalState).f18 = !(!(p1));
          let _556_v92;
          let _nw82 = new _module.C1();
          _nw82.__ctor((_this).f20, (_476_v47).f28, _module.__default.safeDivisionInt((_this).f20, (_this).f20), p1);
          _556_v92 = _nw82;
        } else {
          let _557___mcc_h11 = (_source8).cf8;
          let _558_cf8 = _557___mcc_h11;
          let _559_v93;
          let _nw83 = new _module.C1();
          _nw83.__ctor(_480_v51.f27, _474_v46, (_480_v51.f27).multipliedBy((_this).f20), (_this).f21);
          _559_v93 = _nw83;
          let _560_v94;
          let _nw84 = Array((new BigNumber(28)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = (_476_v47).f28;
          _nw84[(_dafny.ONE).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(2)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(3)).toNumber()] = (_480_v51).f28;
          _nw84[(new BigNumber(4)).toNumber()] = (_480_v51).f28;
          _nw84[(new BigNumber(5)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(6)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(7)).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(8)).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(9)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(10)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(11)).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(12)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(13)).toNumber()] = (_480_v51).f28;
          _nw84[(new BigNumber(14)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(15)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(16)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(17)).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(18)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(19)).toNumber()] = (_559_v93).f28;
          _nw84[(new BigNumber(20)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(21)).toNumber()] = (_480_v51).f28;
          _nw84[(new BigNumber(22)).toNumber()] = _479_v50;
          _nw84[(new BigNumber(23)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(24)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(25)).toNumber()] = (_476_v47).f28;
          _nw84[(new BigNumber(26)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(27)).toNumber()] = (_480_v51).f28;
          _560_v94 = _nw84;
          let _561_v95;
          _561_v95 = _dafny.Seq.of(_560_v94);
          let _562_v96;
          let _nw85 = Array((new BigNumber(29)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = (_561_v95)[_module.__default.safeIndex((((_536_cf34).contains(_476_v47.f27)) ? ((_536_cf34).get(_476_v47.f27)) : (new BigNumber(592))), new BigNumber((_561_v95).length))];
          _nw85[(_dafny.ONE).toNumber()] = _560_v94;
          _nw85[(new BigNumber(2)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(3)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(4)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(5)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(6)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(7)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(8)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(9)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(10)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(11)).toNumber()] = ((p1) ? (_560_v94) : (_560_v94));
          _nw85[(new BigNumber(12)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(13)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(14)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(15)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(16)).toNumber()] = (_561_v95)[_module.__default.safeIndex(_476_v47.f27, new BigNumber((_561_v95).length))];
          _nw85[(new BigNumber(17)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(18)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(19)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(20)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(21)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(22)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(23)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(24)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(25)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(26)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(27)).toNumber()] = _560_v94;
          _nw85[(new BigNumber(28)).toNumber()] = _560_v94;
          _562_v96 = _nw85;
          let _index85 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_562_v96).length));
          (_562_v96)[_index85] = _560_v94;
          let _index86 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_562_v96).length));
          (_562_v96)[_index86] = _560_v94;
          let _563_v97;
          _563_v97 = _module.D7.create_DC17((p1));
          _563_v97 = _563_v97;
          let _564_v98;
          let _init13 = ((_565_v41) => function (_566_i10) {
            return _565_v41;
          })(_469_v41);
          let _nw86 = Array((new BigNumber(22)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw86.length); _i0_13++) {
            _nw86[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _564_v98 = _nw86;
          let _index87 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_564_v98).length));
          (_564_v98)[_index87] = _469_v41;
          let _index88 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_564_v98).length));
          (_564_v98)[_index88] = _469_v41;
        }
        (globalState).f5 = (_this).f20;
        if (!((_this).f21) || ((_this).f21)) {
          let _567_v99;
          let _nw87 = Array((new BigNumber(27)).toNumber()).fill([]);
          _567_v99 = _nw87;
          let _568_v100;
          let _nw88 = Array((new BigNumber(29)).toNumber()).fill([]);
          _568_v100 = _nw88;
          let _index89 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_567_v99).length));
          (_567_v99)[_index89] = _568_v100;
          let _index90 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_567_v99).length));
          (_567_v99)[_index90] = _568_v100;
          let _569_v101;
          _569_v101 = _module.D8.create_DC20();
          let _index91 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_568_v100).length));
          (_568_v100)[_index91] = (_480_v51).f28;
          let _570_v102;
          _570_v102 = _dafny.Seq.of((_480_v51).fm16(_dafny.MultiSet.fromElements(p1), _484_v55, globalState), !(p1));
          let _index92 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_568_v100).length));
          let _rhs72 = _569_v101;
          let _rhs73 = p1;
          let _rhs74 = (_476_v47).f28;
          let _rhs75 = (_this).f21;
          let _rhs76 = _570_v102;
          let _lhs65 = globalState;
          let _lhs66 = _568_v100;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_568_v100).length));
          let _lhs68 = globalState;
          _569_v101 = _rhs72;
          _lhs65.f18 = _rhs73;
          _lhs66[_lhs67] = _rhs74;
          _lhs68.f3 = _rhs75;
          _570_v102 = _rhs76;
          let _index93 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_474_v46).length));
          (_474_v46)[_index93] = (_480_v51.f27).plus(_476_v47.f27);
          let _571_v103;
          _571_v103 = _module.D7.create_DC17(!((_this).f21));
          let _index94 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_474_v46).length));
          (_474_v46)[_index94] = _module.__default.safeDivisionInt(_476_v47.f27, (_480_v51).fm1((_this).f20, _476_v47.f27, !((_571_v103).dtor_cf25), globalState));
          (globalState).f5 = _480_v51.f27;
          let _pat_let_tv17 = _482_v53;
          let _572_v104;
          _572_v104 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let20_0) {
            return function (_573_dt__update__tmp_h1) {
              return function (_pat_let21_0) {
                return function (_574_dt__update_hcf37_h0) {
                  return _module.D10.create_DC27((_573_dt__update__tmp_h1).dtor_cf36, _574_dt__update_hcf37_h0, (_573_dt__update__tmp_h1).dtor_cf38);
                }(_pat_let21_0);
              }(_pat_let_tv17);
            }(_pat_let20_0);
          }(_483_v54),!(!((_this).f21)));
          (_480_v51).f27 = (((_this).f21) ? (_476_v47.f27) : (new BigNumber((_572_v104).length)));
        } else {
          (globalState).f4 = (_this).f21;
          let _575_v105;
          let _nw89 = Array((new BigNumber(10)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _module.__default.fm28(_480_v51.f27, globalState);
          _nw89[(_dafny.ONE).toNumber()] = _471_v43;
          _nw89[(new BigNumber(2)).toNumber()] = _471_v43;
          _nw89[(new BigNumber(3)).toNumber()] = _471_v43;
          _nw89[(new BigNumber(4)).toNumber()] = _471_v43;
          _nw89[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_471_v43, _module.__default.safeIndex(_480_v51.f27, new BigNumber((_471_v43).length)), _476_v47.f27);
          _nw89[(new BigNumber(6)).toNumber()] = _471_v43;
          _nw89[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_471_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), ((_576_v51) => function (_577_i11) {
            return _576_v51.f27;
          })(_480_v51)));
          _nw89[(new BigNumber(8)).toNumber()] = _471_v43;
          _nw89[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p0, (_this).f20, _476_v47.f27);
          _575_v105 = _nw89;
          let _578_v106;
          _578_v106 = _dafny.Map.Empty.slice().updateUnsafe(_480_v51.f27,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(512)), ((_579_v40) => function (_580_i12) {
            return _579_v40;
          })(_468_v40))).length));
          let _581_v107;
          _581_v107 = _dafny.Seq.of(_478_v49, _478_v49);
          let _582_v108;
          _582_v108 = _dafny.Map.Empty.slice().updateUnsafe(!((_476_v47).fm0(_578_v106, _469_v41, _581_v107, globalState)),new _dafny.CodePoint('q'.codePointAt(0)));
          let _rhs77 = (_480_v51).fm1(_480_v51.f27, _480_v51.f27, p1, globalState);
          let _rhs78 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), _480_v51.f27);
          let _rhs79 = (((_this).f21) ? (p1) : ((_476_v47.f27).isLessThanOrEqualTo((_this).f20)));
          let _rhs80 = ((p1) ? (_575_v105) : (_575_v105));
          let _rhs81 = _582_v108;
          let _lhs69 = globalState;
          let _lhs70 = _480_v51;
          let _lhs71 = globalState;
          _lhs69.f17 = _rhs77;
          _lhs70.f27 = _rhs78;
          _lhs71.f18 = _rhs79;
          _575_v105 = _rhs80;
          _582_v108 = _rhs81;
          (globalState).f18 = p1;
          let _583_v109;
          _583_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_478_v49);
          (globalState).f5 = ((new BigNumber((_583_v109).length)).minus(new BigNumber((_module.__default.fm26(p1, globalState)).length))).minus(_476_v47.f27);
          let _584_v110;
          let _init14 = ((_585_p2, _586_p1, _587_v40) => function (_588_i13) {
            return (_585_p2).update(_586_p1, new BigNumber((_dafny.Seq.of(_587_v40, _587_v40)).length));
          })(p2, p1, _468_v40);
          let _nw90 = Array((new BigNumber(13)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw90.length); _i0_14++) {
            _nw90[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _584_v110 = _nw90;
          let _init15 = ((_589_p2) => function (_590_i14) {
            return (_589_p2).Merge(_589_p2);
          })(p2);
          let _nw91 = Array((new BigNumber(13)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw91.length); _i0_15++) {
            _nw91[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _584_v110 = _nw91;
        }
        (_480_v51).f27 = new BigNumber((_472_v44).length);
      }
      let _index95 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_474_v46).length));
      (_474_v46)[_index95] = p0;
      let _591_v111;
      _591_v111 = _module.D6.create_DC13(_469_v41);
      let _592_v112;
      let _nw92 = Array((new BigNumber(22)).toNumber());
      _nw92[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_469_v41, _module.__default.safeIndex(_480_v51.f27, new BigNumber((_469_v41).length)), _468_v40);
      _nw92[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("miga");
      _nw92[(new BigNumber(2)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(3)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(4)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(5)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(6)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_469_v41, _469_v41);
      _nw92[(new BigNumber(8)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_593_v40) => function (_594_i15) {
        return _593_v40;
      })(_468_v40));
      _nw92[(new BigNumber(10)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(11)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(12)).toNumber()] = _module.__default.fm26(true, globalState);
      _nw92[(new BigNumber(13)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(14)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(15)).toNumber()] = (_591_v111).dtor_cf21;
      _nw92[(new BigNumber(16)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(17)).toNumber()] = _469_v41;
      _nw92[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_469_v41, _469_v41);
      _nw92[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), function (_595_i16) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      });
      _nw92[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("qifnsfgu");
      _nw92[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(922)), ((_596_v40) => function (_597_i17) {
        return _596_v40;
      })(_468_v40));
      _592_v112 = _nw92;
      let _index96 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_592_v112).length));
      (_592_v112)[_index96] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bwsydd"), _469_v41);
      let _598_v113;
      let _nw93 = Array((new BigNumber(12)).toNumber()).fill(false);
      _598_v113 = _nw93;
      let _index97 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_598_v113).length));
      (_598_v113)[_index97] = (p0).isEqualTo(_480_v51.f27);
      let _599_v114;
      let _nw94 = new _module.C0();
      _nw94.__ctor();
      _599_v114 = _nw94;
      let _600_v115;
      _600_v115 = _dafny.MultiSet.fromElements(_599_v114);
      let _601_v116;
      _601_v116 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,new BigNumber((_600_v115).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(p1,_480_v51.f27));
      let _602_v117;
      _602_v117 = _dafny.Seq.of(_469_v41);
      let _603_v118;
      _603_v118 = _module.D11.create_DC28(_592_v112);
      let _604_v119;
      _604_v119 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(144),new BigNumber((_469_v41).length));
      let _605_v120;
      _605_v120 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_476_v47).fm16(_dafny.MultiSet.fromElements((_this).f21), _484_v55, globalState),p1), _478_v49, _478_v49);
      let _index98 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_474_v46).length));
      let _index99 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_592_v112).length));
      let _index100 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_598_v113).length));
      let _rhs82 = new BigNumber((_dafny.Seq.Concat(_601_v116, _dafny.Seq.Concat(_601_v116, _601_v116))).length);
      let _rhs83 = _dafny.Seq.Concat((_602_v117)[_module.__default.safeIndex(new BigNumber((_469_v41).length), new BigNumber((_602_v117).length))], _dafny.Seq.Concat(_469_v41, _469_v41));
      let _rhs84 = (_603_v118).dtor_cf39;
      let _rhs85 = (_this).fm0(((_604_v119).update((_this).f20, _476_v47.f27)).Merge(_604_v119), _469_v41, _dafny.Seq.update(_605_v120, _module.__default.safeIndex(new BigNumber(166), new BigNumber((_605_v120).length)), _478_v49), globalState);
      let _lhs72 = _474_v46;
      let _lhs73 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_474_v46).length));
      let _lhs74 = _592_v112;
      let _lhs75 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_592_v112).length));
      let _lhs76 = _598_v113;
      let _lhs77 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_598_v113).length));
      _lhs72[_lhs73] = _rhs82;
      _lhs74[_lhs75] = _rhs83;
      _592_v112 = _rhs84;
      _lhs76[_lhs77] = _rhs85;
      let _606_v122;
      _606_v122 = _dafny.Seq.of(!((_598_v113)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_598_v113).length))]), p1);
      let _607_v123;
      _607_v123 = _dafny.Map.Empty.slice().updateUnsafe((_606_v122)[_module.__default.safeIndex((_this).fm1(_476_v47.f27, new BigNumber((_dafny.MultiSet.FromArray(_602_v117)).cardinality()), p1, globalState), new BigNumber((_606_v122).length))],_604_v119);
      let _608_v124;
      let _nw95 = Array((new BigNumber(11)).toNumber());
      _nw95[(_dafny.ZERO).toNumber()] = _604_v119;
      _nw95[(_dafny.ONE).toNumber()] = _604_v119;
      _nw95[(new BigNumber(2)).toNumber()] = _604_v119;
      _nw95[(new BigNumber(3)).toNumber()] = ((p1) ? (_604_v119) : (_604_v119));
      _nw95[(new BigNumber(4)).toNumber()] = (_604_v119).update(new BigNumber(-384), _476_v47.f27);
      _nw95[(new BigNumber(5)).toNumber()] = _604_v119;
      _nw95[(new BigNumber(6)).toNumber()] = (_604_v119).Merge(_604_v119);
      _nw95[(new BigNumber(7)).toNumber()] = _604_v119;
      _nw95[(new BigNumber(8)).toNumber()] = ((false) ? (_604_v119) : (function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-508), new BigNumber(-87))) {
          let _609_v121 = _compr_20;
          if (((new BigNumber(-508)).isLessThanOrEqualTo(_609_v121)) && ((_609_v121).isLessThan(new BigNumber(-87)))) {
            _coll20.push([_module.__default.safeModuloInt(_609_v121, new BigNumber(392)),_480_v51.f27]);
          }
        }
        return _coll20;
      }()));
      _nw95[(new BigNumber(9)).toNumber()] = (_604_v119).Merge((((_607_v123).contains(false)) ? ((_607_v123).get(false)) : (_604_v119)));
      _nw95[(new BigNumber(10)).toNumber()] = _604_v119;
      _608_v124 = _nw95;
      (globalState).f8 = _608_v124;
      r0 = ((_472_v44).IsProperSubsetOf(function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of _dafny.IntegerRange(new BigNumber(72), new BigNumber(341))) {
          let _610_v125 = _compr_21;
          if (((new BigNumber(72)).isLessThanOrEqualTo(_610_v125)) && ((_610_v125).isLessThan(new BigNumber(341)))) {
            _coll21.add((_610_v125).multipliedBy((_this).f20));
          }
        }
        return _coll21;
      }())) && ((_470_v42).equals(_470_v42));
      r1 = (_472_v44).IsProperSubsetOf((_472_v44).Intersect(function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-389), new BigNumber(-64))) {
          let _611_v126 = _compr_22;
          if (((new BigNumber(-389)).isLessThanOrEqualTo(_611_v126)) && ((_611_v126).isLessThan(new BigNumber(-64)))) {
            _coll22.add((_611_v126).minus(_476_v47.f27));
          }
        }
        return _coll22;
      }()));
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
      this._f25 = _dafny.ZERO;
      this._f26 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f25, f26, f20, f21) {
      let _this = this;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f26;
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f25;
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _612_v0;
      _612_v0 = _module.D3.create_DC6((_this).f25);
      let _pat_let_tv18 = p1;
      let _pat_let_tv19 = p1;
      let _pat_let_tv20 = p1;
      let _pat_let_tv21 = p1;
      let _source9 = function (_source10) {
        if (_source10.is_DC6) {
          let _613___mcc_h6 = (_source10).cf10;
          let _614_cf10 = _613___mcc_h6;
          return _module.D4.create_DC10(_dafny.MultiSet.fromElements((false)), (_this).f21);
        } else if (_source10.is_DC7) {
          let _615___mcc_h7 = (_source10).cf11;
          let _616___mcc_h8 = (_source10).cf12;
          let _617_cf12 = _616___mcc_h8;
          let _618_cf11 = _615___mcc_h7;
          return _module.D4.create_DC10(_dafny.MultiSet.fromElements(_617_cf12, (_this).f26, _pat_let_tv18), (_this).f21);
        } else if (_source10.is_DC5) {
          let _619___mcc_h9 = (_source10).cf9;
          let _620_cf9 = _619___mcc_h9;
          return _module.D4.create_DC10(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, (_this).f26, _pat_let_tv19, false)), _pat_let_tv20);
        } else {
          let _621___mcc_h10 = (_source10).cf13;
          let _622_cf13 = _621___mcc_h10;
          return _module.D4.create_DC10(_dafny.MultiSet.fromElements(!(_pat_let_tv21)), (_this).f26);
        }
      }(_612_v0);
      if (_source9.is_DC10) {
        let _623___mcc_h0 = (_source9).cf15;
        let _624___mcc_h1 = (_source9).cf16;
        let _625_cf16 = _624___mcc_h1;
        let _626_cf15 = _623___mcc_h0;
        let _627_v1;
        _627_v1 = _dafny.Seq.UnicodeFromString("dknejhjps");
        let _628_v2;
        let _nw96 = Array((new BigNumber(16)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).f25, (_this).f20);
        _nw96[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw96[(new BigNumber(2)).toNumber()] = (_this).fm1(new BigNumber((_dafny.Seq.UnicodeFromString("nbras")).length), (_this).f20, (_this).f26, globalState);
        _nw96[(new BigNumber(3)).toNumber()] = p0;
        _nw96[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_this).f25, (_this).f20);
        _nw96[(new BigNumber(5)).toNumber()] = p0;
        _nw96[(new BigNumber(6)).toNumber()] = new BigNumber(205);
        _nw96[(new BigNumber(7)).toNumber()] = (((_this).f21) ? ((_this).f20) : ((_this).f25));
        _nw96[(new BigNumber(8)).toNumber()] = ((_this).f20).minus(new BigNumber((_627_v1).length));
        _nw96[(new BigNumber(9)).toNumber()] = (_this).f25;
        _nw96[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw96[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt((_this).f20, p0);
        _nw96[(new BigNumber(12)).toNumber()] = (_this).fm1(p0, new BigNumber(978), (_this).f21, globalState);
        _nw96[(new BigNumber(13)).toNumber()] = (_this).f25;
        _nw96[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f20, (_this).f20));
        _nw96[(new BigNumber(15)).toNumber()] = (_this).f20;
        _628_v2 = _nw96;
        let _index101 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length));
        (_628_v2)[_index101] = (_this).fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(759)), function (_629_i0) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length), (_this).f20, p1, globalState);
        let _630_v3;
        let _init16 = ((_631_p0) => function (_632_i1) {
          return _module.D4.create_DC11(_631_p0, (_this).f21, false);
        })(p0);
        let _nw97 = Array((new BigNumber(22)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw97.length); _i0_16++) {
          _nw97[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _630_v3 = _nw97;
        let _633_v4;
        _633_v4 = _628_v2;
        let _634_v6;
        _634_v6 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f21),(_this).f21);
        let _635_v7;
        _635_v7 = _dafny.Seq.of(_634_v6);
        let _636_v8;
        _636_v8 = _dafny.Seq.of((_this).f20, p0);
        let _637_v9;
        let _nw98 = Array((new BigNumber(25)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = (_633_v4);
        _nw98[(_dafny.ONE).toNumber()] = _628_v2;
        _nw98[(new BigNumber(2)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(3)).toNumber()] = (((_this).f21) ? (_628_v2) : (_628_v2));
        _nw98[(new BigNumber(4)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(5)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(6)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(7)).toNumber()] = (((_this).fm0(function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(82), new BigNumber(-571))) {
            let _638_v5 = _compr_23;
            if (((new BigNumber(82)).isLessThanOrEqualTo(_638_v5)) && ((_638_v5).isLessThan(new BigNumber(-571)))) {
              _coll23.push([_module.__default.safeModuloInt(_638_v5, new BigNumber(939)),(_this).f20]);
            }
          }
          return _coll23;
        }(), _627_v1, _dafny.Seq.update(_635_v7, _module.__default.safeIndex((_636_v8)[_module.__default.safeIndex((_this).f20, new BigNumber((_636_v8).length))], new BigNumber((_635_v7).length)), _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_625_cf16)), globalState)) ? (_628_v2) : (_628_v2));
        _nw98[(new BigNumber(8)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(9)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(10)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(11)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(12)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(13)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(14)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(15)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(16)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(17)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(18)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(19)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(20)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(21)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(22)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(23)).toNumber()] = _628_v2;
        _nw98[(new BigNumber(24)).toNumber()] = _628_v2;
        _637_v9 = _nw98;
        let _index102 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_637_v9).length));
        (_637_v9)[_index102] = _628_v2;
        let _639_v10;
        let _init17 = ((_640_v6, _641_p1) => function (_642_i2) {
          return (_640_v6).update((_this).f21, _641_p1);
        })(_634_v6, p1);
        let _nw99 = Array((new BigNumber(26)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw99.length); _i0_17++) {
          _nw99[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _639_v10 = _nw99;
        let _index103 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_639_v10).length));
        (_639_v10)[_index103] = _634_v6;
        let _643_v11;
        _643_v11 = new _dafny.CodePoint('b'.codePointAt(0));
        let _644_v12;
        _644_v12 = _dafny.MultiSet.fromElements(_643_v11);
        let _index104 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length));
        let _index105 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_637_v9).length));
        let _index106 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_639_v10).length));
        let _rhs86 = ((_dafny.ZERO).minus((p0).minus(new BigNumber((_644_v12).cardinality())))).minus(new BigNumber(-180));
        let _rhs87 = _630_v3;
        let _rhs88 = _628_v2;
        let _rhs89 = !((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_645_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
        })).length))).isEqualTo((_this).f20);
        let _rhs90 = (((_this).f21) ? ((_634_v6).update((_this).f26, (_this).f21)) : (_module.__default.fm15((_this).f20, (_this).f25, (_this).f25, new BigNumber(-399), globalState)));
        let _lhs78 = _628_v2;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length));
        let _lhs80 = _637_v9;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_637_v9).length));
        let _lhs82 = globalState;
        let _lhs83 = _639_v10;
        let _lhs84 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_639_v10).length));
        _lhs78[_lhs79] = _rhs86;
        _630_v3 = _rhs87;
        _lhs80[_lhs81] = _rhs88;
        _lhs82.f3 = _rhs89;
        _lhs83[_lhs84] = _rhs90;
        if (((_this).f20).isEqualTo((_this).fm1(new BigNumber((_636_v8).length), (_628_v2)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length))], false, globalState))) {
          (globalState).f15 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), ((_646_v11) => function (_647_i4) {
            return _646_v11;
          })(_643_v11));
          let _648_v13;
          _648_v13 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(((_this).f20).plus((_this).f25)));
          _648_v13 = _648_v13;
          let _649_v14;
          let _nw100 = Array((new BigNumber(28)).toNumber()).fill(false);
          _649_v14 = _nw100;
          let _index107 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_649_v14).length));
          (_649_v14)[_index107] = (_this).f21;
          let _650_v15;
          _650_v15 = _module.D4.create_DC10(_626_cf15, p1);
          let _index108 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_649_v14).length));
          (_649_v14)[_index108] = (_650_v15).dtor_cf16;
          let _651_v16;
          _651_v16 = _dafny.Map.Empty.slice().updateUnsafe((_628_v2)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length))],(_this).f25);
          let _652_v17;
          _652_v17 = _dafny.Map.Empty.slice().updateUnsafe(_625_cf16,new BigNumber((_627_v1).length));
          let _653_v18;
          _653_v18 = _dafny.Map.Empty.slice().updateUnsafe(_651_v16,new BigNumber((_652_v17).length));
          (globalState).f17 = (((_653_v18).contains(_651_v16)) ? ((_653_v18).get(_651_v16)) : (new BigNumber(-292)));
          let _654_v19;
          let _nw101 = new _module.C0();
          _nw101.__ctor();
          _654_v19 = _nw101;
        } else {
          let _655_v20;
          _655_v20 = _dafny.Set.fromElements(_626_cf15);
          let _656_v21;
          _656_v21 = _dafny.Seq.of(_dafny.Set.fromElements(_626_cf15, _626_cf15, _626_cf15), _dafny.Set.fromElements(_626_cf15, _dafny.MultiSet.fromElements(false, p1), _626_cf15));
          _655_v20 = (_656_v21)[_module.__default.safeIndex(((_this).f20).plus((_this).f25), new BigNumber((_656_v21).length))];
          let _657_v22;
          let _nw102 = Array((new BigNumber(6)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _633_v4;
          _nw102[(_dafny.ONE).toNumber()] = _633_v4;
          _nw102[(new BigNumber(2)).toNumber()] = _633_v4;
          _nw102[(new BigNumber(3)).toNumber()] = _633_v4;
          _nw102[(new BigNumber(4)).toNumber()] = _633_v4;
          _nw102[(new BigNumber(5)).toNumber()] = _633_v4;
          _657_v22 = _nw102;
          let _658_v23;
          _658_v23 = _dafny.Map.Empty.slice().updateUnsafe(_657_v22,p1);
          _658_v23 = (_658_v23).update(_657_v22, (_this).f26);
          let _index109 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length));
          (_628_v2)[_index109] = p0;
          let _659_v24;
          let _nw103 = new _module.C0();
          _nw103.__ctor();
          _659_v24 = _nw103;
          let _660_v25;
          _660_v25 = _dafny.Map.Empty.slice().updateUnsafe(_659_v24,(_dafny.ZERO).minus(((_this).f25).multipliedBy((_628_v2)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length))])));
          let _661_v26;
          _661_v26 = _dafny.Seq.of(_659_v24);
          _660_v25 = (_660_v25).update((_661_v26)[_module.__default.safeIndex((_628_v2)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length))], new BigNumber((_661_v26).length))], (_this).f25);
          let _662_v27;
          _662_v27 = _dafny.Seq.of((_633_v4), (_637_v9)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_637_v9).length))], (_637_v9)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_637_v9).length))]);
          let _663_v28;
          let _nw104 = new _module.C1();
          _nw104.__ctor(((_636_v8)[_module.__default.safeIndex((_this).f20, new BigNumber((_636_v8).length))]).plus(p0), (_662_v27)[_module.__default.safeIndex(new BigNumber((_module.__default.fm29((_this).f25, (_this).f26, globalState)).length), new BigNumber((_662_v27).length))], new BigNumber((_627_v1).length), (_this).f21);
          _663_v28 = _nw104;
          let _664_v29;
          _664_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_639_v10)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_639_v10).length))])).length),(_663_v28).f21);
          let _665_v30;
          _665_v30 = _module.D9.create_DC23((_628_v2)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_628_v2).length))], (_this).f26);
          let _666_v31;
          let _init18 = ((_667_v1) => function (_668_i5) {
            return _module.D6.create_DC13(_667_v1);
          })(_627_v1);
          let _nw105 = Array((new BigNumber(16)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
            _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _666_v31 = _nw105;
          let _669_v32;
          _669_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_666_v31);
          let _670_v33;
          let _nw106 = Array((new BigNumber(10)).toNumber()).fill(false);
          _670_v33 = _nw106;
          let _671_v34;
          _671_v34 = _dafny.Seq.of(_670_v33, _670_v33);
          let _672_v35;
          _672_v35 = _module.D3.create_DC5(_671_v34);
          let _673_v36;
          _673_v36 = _dafny.MultiSet.fromElements(_672_v35);
          let _674_v37;
          _674_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(((_673_v36).contains(_672_v35)) ? ((_673_v36).get(_672_v35)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), ((_675_v11) => function (_676_i6) {
            return _675_v11;
          })(_643_v11))).length))));
          let _677_v38;
          _677_v38 = _dafny.Seq.of(_674_v37);
          let _678_v39;
          let _nw107 = new _module.C1();
          _nw107.__ctor(p0, _628_v2, (_663_v28).f20, _625_cf16);
          _678_v39 = _nw107;
          let _679_v40;
          _679_v40 = _dafny.Seq.of(_678_v39);
          let _680_v41;
          let _nw108 = Array((new BigNumber(19)).toNumber());
          _nw108[(_dafny.ZERO).toNumber()] = p1;
          _nw108[(_dafny.ONE).toNumber()] = (_663_v28).f21;
          _nw108[(new BigNumber(2)).toNumber()] = (((_664_v29).contains((_663_v28).f20)) ? ((_664_v29).get((_663_v28).f20)) : (true));
          _nw108[(new BigNumber(3)).toNumber()] = (_665_v30).dtor_cf32;
          _nw108[(new BigNumber(4)).toNumber()] = (_this).f21;
          _nw108[(new BigNumber(5)).toNumber()] = (_this).f26;
          _nw108[(new BigNumber(6)).toNumber()] = (_663_v28).f21;
          _nw108[(new BigNumber(7)).toNumber()] = _625_cf16;
          _nw108[(new BigNumber(8)).toNumber()] = false;
          _nw108[(new BigNumber(9)).toNumber()] = ((_663_v28).f21) === ((_this).f21);
          _nw108[(new BigNumber(10)).toNumber()] = ((_625_cf16) ? ((_this).f21) : (_625_cf16));
          _nw108[(new BigNumber(11)).toNumber()] = !(((_625_cf16) ? ((_this).f21) : (_625_cf16)));
          _nw108[(new BigNumber(12)).toNumber()] = (_659_v24).fm4(new BigNumber((_669_v32).length), _677_v38, (_this).f20, globalState);
          _nw108[(new BigNumber(13)).toNumber()] = (_this).f21;
          _nw108[(new BigNumber(14)).toNumber()] = _dafny.Seq.contains(_679_v40, _678_v39);
          _nw108[(new BigNumber(15)).toNumber()] = (_678_v39).fm0(_674_v37, _dafny.Seq.UnicodeFromString("safjb"), _635_v7, globalState);
          _nw108[(new BigNumber(16)).toNumber()] = p1;
          _nw108[(new BigNumber(17)).toNumber()] = (_this).f21;
          _nw108[(new BigNumber(18)).toNumber()] = (_this).f26;
          _680_v41 = _nw108;
          let _index110 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_670_v33).length));
          (_670_v33)[_index110] = p1;
          let _681_v42;
          _681_v42 = _dafny.MultiSet.fromElements(new BigNumber(286), (_this).f25);
          let _682_v44;
          _682_v44 = _dafny.Set.fromElements((_this).f26, (_this).f21, (_this).f26);
          let _683_v45;
          _683_v45 = _dafny.Seq.of(new BigNumber((_682_v44).length));
          let _684_v46;
          _684_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,new BigNumber(-665));
          let _index111 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_670_v33).length));
          let _rhs91 = (((_663_v28).f21) ? (_module.__default.safeDivisionInt(new BigNumber(-584), new BigNumber((function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of (_681_v42).Elements) {
              let _685_v43 = _compr_24;
              if ((_681_v42).contains(_685_v43)) {
                _coll24.add(_module.__default.safeDivisionInt(_685_v43, new BigNumber(145)));
              }
            }
            return _coll24;
          }()).length))) : ((_this).f25));
          let _rhs92 = _663_v28;
          let _rhs93 = !((_663_v28).f21);
          let _rhs94 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_686_i7) {
            return (_this).f20;
          });
          let _rhs95 = !(!((_678_v39).fm16(_626_cf15, _683_v45, globalState)) || (p1)) || (!(_684_v46).contains((_this).f26));
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = _670_v33;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_670_v33).length));
          _lhs85.f17 = _rhs91;
          _663_v28 = _rhs92;
          _lhs86.f3 = _rhs93;
          _636_v8 = _rhs94;
          _lhs87[_lhs88] = _rhs95;
        }
        let _687_v47;
        _687_v47 = _dafny.MultiSet.fromElements(new BigNumber(512));
        let _688_v48;
        _688_v48 = _dafny.Map.Empty.slice().updateUnsafe(_687_v47,_dafny.Seq.UnicodeFromString("auwfsukpw"));
        _688_v48 = (_688_v48).update(_687_v47, _dafny.Seq.Concat(_627_v1, _627_v1));
        let _689_v49;
        _689_v49 = _dafny.Seq.of(_687_v47);
        _689_v49 = _689_v49;
      } else if (_source9.is_DC11) {
        let _690___mcc_h2 = (_source9).cf17;
        let _691___mcc_h3 = (_source9).cf18;
        let _692___mcc_h4 = (_source9).cf19;
        let _693_cf19 = _692___mcc_h4;
        let _694_cf18 = _691___mcc_h3;
        let _695_cf17 = _690___mcc_h2;
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(372)), function (_696_i8) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), function (_697_i9) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }))) {
          let _698_v50;
          _698_v50 = _dafny.Seq.UnicodeFromString("iagyloo");
          let _699_v51;
          let _init19 = function (_700_i10) {
            return _module.__default.safeDivisionInt(_700_i10, (_dafny.ZERO).minus((_this).f25));
          };
          let _nw109 = Array((new BigNumber(2)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw109.length); _i0_19++) {
            _nw109[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _699_v51 = _nw109;
          let _701_v52;
          let _nw110 = new _module.C1();
          _nw110.__ctor((new BigNumber((_698_v50).length)).plus((_this).f20), _699_v51, p0, _694_cf18);
          _701_v52 = _nw110;
          let _702_v53;
          let _nw111 = new _module.C0();
          _nw111.__ctor();
          _702_v53 = _nw111;
          let _703_v54;
          let _nw112 = Array((new BigNumber(14)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _694_cf18;
          _nw112[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw112[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw112[(new BigNumber(3)).toNumber()] = (_this).f26;
          _nw112[(new BigNumber(4)).toNumber()] = _693_cf19;
          _nw112[(new BigNumber(5)).toNumber()] = false;
          _nw112[(new BigNumber(6)).toNumber()] = (_this).f21;
          _nw112[(new BigNumber(7)).toNumber()] = (_this).f26;
          _nw112[(new BigNumber(8)).toNumber()] = _693_cf19;
          _nw112[(new BigNumber(9)).toNumber()] = (_this).f26;
          _nw112[(new BigNumber(10)).toNumber()] = _694_cf18;
          _nw112[(new BigNumber(11)).toNumber()] = !(_694_cf18);
          _nw112[(new BigNumber(12)).toNumber()] = _694_cf18;
          _nw112[(new BigNumber(13)).toNumber()] = _693_cf19;
          _703_v54 = _nw112;
          let _704_v55;
          _704_v55 = _dafny.Set.fromElements(_703_v54, _703_v54, _703_v54);
          r1 = (_704_v55).contains(_703_v54);
          let _705_v56;
          _705_v56 = new _dafny.CodePoint('j'.codePointAt(0));
          let _706_v57;
          _706_v57 = _dafny.MultiSet.fromElements(_705_v56, _705_v56);
          let _707_v58;
          _707_v58 = _dafny.MultiSet.fromElements(_706_v57);
          let _708_v59;
          _708_v59 = _dafny.Seq.of(_707_v58, _dafny.MultiSet.fromElements(_706_v57, _dafny.MultiSet.fromElements(_705_v56)));
          let _709_v60;
          _709_v60 = _dafny.Seq.of(_707_v58, _707_v58, _dafny.MultiSet.fromElements(_706_v57), (_708_v59)[_module.__default.safeIndex(_695_cf17, new BigNumber((_708_v59).length))], _module.__default.fm31(globalState));
          _693_cf19 = !((_dafny.MultiSet.FromArray(_module.__default.fm30(globalState))).equals((_708_v59)[_module.__default.safeIndex((_this).f20, new BigNumber((_708_v59).length))]));
          (globalState).f4 = (_this).f21;
        } else {
          let _710_v61;
          _710_v61 = _dafny.MultiSet.fromElements(p0, _695_cf17);
          let _711_v62;
          _711_v62 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).fm1(p0, new BigNumber((_710_v61).cardinality()), _693_cf19, globalState));
          let _712_v63;
          _712_v63 = _dafny.Map.Empty.slice().updateUnsafe(((!(true)) ? ((_this).f26) : ((_this).f21)),(_this).fm1((((_711_v62).contains((_this).f26)) ? ((_711_v62).get((_this).f26)) : (p0)), p0, (_this).f21, globalState));
          _712_v63 = (_712_v63).update((_this).f26, new BigNumber(-402));
          let _713_v64;
          _713_v64 = _dafny.Seq.UnicodeFromString("qvfhbcfbh");
          let _714_v65;
          _714_v65 = _dafny.Seq.of(_module.__default.fm26(_694_cf18, globalState), _713_v64);
          _714_v65 = _714_v65;
          let _715_v66;
          _715_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _716_v67;
          _716_v67 = _dafny.Seq.of((((_715_v66).contains(_694_cf18)) ? ((_715_v66).get(_694_cf18)) : (_693_cf19)));
          let _717_v68;
          _717_v68 = _dafny.MultiSet.fromElements(_716_v67, _716_v67);
          let _718_v69;
          _718_v69 = _dafny.Seq.of(new BigNumber(594), new BigNumber((_717_v68).cardinality()), (_this).fm1(_695_cf17, new BigNumber((_715_v66).length), _694_cf18, globalState));
          let _719_v70;
          _719_v70 = _dafny.Set.fromElements((_dafny.ZERO).minus(_695_cf17));
          let _720_v71;
          _720_v71 = _dafny.Seq.of(_719_v70, _719_v70);
          let _721_v72;
          _721_v72 = _dafny.Seq.of(_695_cf17, (_this).f25, (_module.D7.create_DC18(_718_v69, _695_cf17)).dtor_cf27, (_this).f25, (_module.D8.create_DC21(new BigNumber(((_720_v71)[_module.__default.safeIndex((_this).f20, new BigNumber((_720_v71).length))]).length))).dtor_cf29);
          let _722_v73;
          let _nw113 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _722_v73 = _nw113;
          let _723_v74;
          _723_v74 = _dafny.Seq.of(_722_v73);
          let _rhs96 = _694_cf18;
          let _rhs97 = ((!(true)) ? ((_dafny.ZERO).minus(_695_cf17)) : ((_dafny.ZERO).minus(p0)));
          let _rhs98 = _dafny.Seq.IsProperPrefixOf(_721_v72, _721_v72);
          let _rhs99 = _723_v74;
          let _lhs89 = globalState;
          let _lhs90 = globalState;
          let _lhs91 = globalState;
          _lhs89.f18 = _rhs96;
          _lhs90.f5 = _rhs97;
          _lhs91.f14 = _rhs98;
          r0 = _rhs99;
          let _724_v75;
          _724_v75 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_718_v69).length));
          let _725_v76;
          _725_v76 = _dafny.Map.Empty.slice().updateUnsafe(_724_v75,(_this).f20);
          _725_v76 = (_725_v76).update(_724_v75, p0);
          let _726_v77;
          _726_v77 = new _dafny.CodePoint('v'.codePointAt(0));
          let _727_v78;
          let _nw114 = new _module.C1();
          _nw114.__ctor(_695_cf17, _722_v73, (_this).f25, _694_cf18);
          _727_v78 = _nw114;
          let _728_v79;
          _728_v79 = _dafny.Map.Empty.slice().updateUnsafe(_726_v77,_727_v78);
          let _729_v80;
          _729_v80 = _dafny.Seq.of(_728_v79, _728_v79);
          let _730_v81;
          _730_v81 = _module.D10.create_DC27(_695_cf17, _729_v80, _726_v77);
          let _731_v82;
          _731_v82 = _dafny.MultiSet.fromElements((_730_v81).dtor_cf38);
          let _732_v83;
          _732_v83 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_713_v64), _731_v82);
          let _733_v84;
          let _nw115 = Array((new BigNumber(17)).toNumber());
          _nw115[(_dafny.ZERO).toNumber()] = (_731_v82).Union(_module.__default.fm32(globalState));
          _nw115[(_dafny.ONE).toNumber()] = _731_v82;
          _nw115[(new BigNumber(2)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(3)).toNumber()] = (_732_v83)[_module.__default.safeIndex((_this).f20, new BigNumber((_732_v83).length))];
          _nw115[(new BigNumber(4)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(5)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), _726_v77));
          _nw115[(new BigNumber(7)).toNumber()] = ((_693_cf19) ? (_dafny.MultiSet.fromElements(_726_v77, new _dafny.CodePoint('v'.codePointAt(0)))) : (_731_v82));
          _nw115[(new BigNumber(8)).toNumber()] = (_731_v82).Union((_732_v83)[_module.__default.safeIndex(_727_v78.f27, new BigNumber((_732_v83).length))]);
          _nw115[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_726_v77, _726_v77, _726_v77, _726_v77, _726_v77);
          _nw115[(new BigNumber(10)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(11)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_726_v77, _726_v77);
          _nw115[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_726_v77);
          _nw115[(new BigNumber(14)).toNumber()] = _731_v82;
          _nw115[(new BigNumber(15)).toNumber()] = (_731_v82).Difference(_731_v82);
          _nw115[(new BigNumber(16)).toNumber()] = _731_v82;
          _733_v84 = _nw115;
          _733_v84 = _733_v84;
        }
        (globalState).f17 = (_this).f25;
        let _734_v85;
        _734_v85 = _693_cf19;
        if ((_734_v85)) {
          let _735_v86;
          let _nw116 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _735_v86 = _nw116;
          let _index112 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_735_v86).length));
          (_735_v86)[_index112] = (_this).f25;
          let _736_v87;
          _736_v87 = _module.D7.create_DC17(_694_cf18);
          let _737_v88;
          _737_v88 = _dafny.Set.fromElements(false);
          let _738_v89;
          _738_v89 = _dafny.Seq.of(_737_v88, _737_v88);
          let _739_v90;
          _739_v90 = _dafny.Map.Empty.slice().updateUnsafe(_736_v87,_738_v89);
          let _740_v91;
          _740_v91 = _module.D3.create_DC7((((_739_v90).contains(_736_v87)) ? ((_739_v90).get(_736_v87)) : (_738_v89)), _693_cf19);
          let _741_v92;
          _741_v92 = _dafny.MultiSet.fromElements((_this).f20, (_this).fm1(new BigNumber(-252), (_dafny.ZERO).minus(_695_cf17), (_740_v91).dtor_cf12, globalState), (_this).f25);
          let _index113 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_735_v86).length));
          (_735_v86)[_index113] = (((_741_v92).contains((_this).f20)) ? ((_741_v92).get((_this).f20)) : ((_this).f25));
          let _742_v93;
          let _init20 = ((_743_p1) => function (_744_i11) {
            return _743_p1;
          })(p1);
          let _nw117 = Array((new BigNumber(12)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw117.length); _i0_20++) {
            _nw117[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _742_v93 = _nw117;
          let _nw118 = Array((new BigNumber(9)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = p1;
          _nw118[(_dafny.ONE).toNumber()] = !(_694_cf18) || (p1);
          _nw118[(new BigNumber(2)).toNumber()] = true;
          _nw118[(new BigNumber(3)).toNumber()] = p1;
          _nw118[(new BigNumber(4)).toNumber()] = (_this).f21;
          _nw118[(new BigNumber(5)).toNumber()] = _694_cf18;
          _nw118[(new BigNumber(6)).toNumber()] = p1;
          _nw118[(new BigNumber(7)).toNumber()] = !((!(p1)) && (_694_cf18));
          _nw118[(new BigNumber(8)).toNumber()] = true;
          _742_v93 = _nw118;
          let _745_v94;
          let _nw119 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
          _745_v94 = _nw119;
          let _746_v95;
          _746_v95 = _dafny.Seq.of(_735_v86, _735_v86);
          let _747_v96;
          _747_v96 = _dafny.Seq.UnicodeFromString("ysgax");
          let _rhs100 = _745_v94;
          let _rhs101 = new BigNumber(785);
          let _rhs102 = new BigNumber((_dafny.Seq.of((_746_v95)[_module.__default.safeIndex(p0, new BigNumber((_746_v95).length))], _735_v86)).length);
          let _rhs103 = _693_cf19;
          let _rhs104 = _module.__default.safeModuloInt(p0, (p0).minus(new BigNumber((_747_v96).length)));
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          let _lhs94 = globalState;
          _745_v94 = _rhs100;
          _lhs92.f17 = _rhs101;
          _695_cf17 = _rhs102;
          _lhs93.f14 = _rhs103;
          _lhs94.f5 = _rhs104;
          let _748_v97;
          _748_v97 = _dafny.Seq.of((_this).f21, _694_cf18);
          let _749_v98;
          _749_v98 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_748_v97);
          (globalState).f18 = (new BigNumber((_749_v98).length)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_747_v96, _module.__default.safeIndex(new BigNumber(-800), new BigNumber((_747_v96).length)), new _dafny.CodePoint('f'.codePointAt(0)))).length)));
          (globalState).f4 = (_694_cf18) || (false);
        } else {
          let _750_v99;
          let _nw120 = Array((new BigNumber(25)).toNumber()).fill(_module.D6.Default());
          _750_v99 = _nw120;
          let _751_v100;
          _751_v100 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f25, (_this).f20, new BigNumber(81), p0));
          let _752_v101;
          _752_v101 = _dafny.Seq.of((_this).f26);
          let _rhs105 = ((p1) ? (_750_v99) : (_750_v99));
          let _rhs106 = _751_v100;
          let _rhs107 = (_752_v101)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("uvtiyhvfl")).length), new BigNumber((_752_v101).length))];
          let _rhs108 = (_this).f20;
          let _rhs109 = _695_cf17;
          let _lhs95 = globalState;
          _750_v99 = _rhs105;
          _751_v100 = _rhs106;
          _lhs95.f14 = _rhs107;
          _695_cf17 = _rhs108;
          _695_cf17 = _rhs109;
          let _753_v102;
          let _init21 = function (_754_i12) {
            return (_dafny.Set.fromElements((_this).f21)).IsProperSubsetOf(_dafny.Set.fromElements((_this).f26));
          };
          let _nw121 = Array((new BigNumber(27)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw121.length); _i0_21++) {
            _nw121[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _753_v102 = _nw121;
          let _index114 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_753_v102).length));
          (_753_v102)[_index114] = (_694_cf18) === (_693_cf19);
          let _index115 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_753_v102).length));
          (_753_v102)[_index115] = (_this).f26;
          let _755_v103;
          _755_v103 = _dafny.Seq.UnicodeFromString("etpbgjoh");
          let _756_v104;
          _756_v104 = _module.D6.create_DC13(_755_v103);
          (globalState).f15 = (_756_v104).dtor_cf21;
          let _757_v105;
          _757_v105 = _dafny.Map.Empty.slice().updateUnsafe(p0,_755_v103);
          let _758_v106;
          _758_v106 = _dafny.Seq.of(_755_v103, _755_v103, _755_v103, _755_v103);
          _757_v105 = (_757_v105).update(p0, (_758_v106)[_module.__default.safeIndex(p0, new BigNumber((_758_v106).length))]);
          let _759_v107;
          _759_v107 = _module.D9.create_DC23(new BigNumber(75), p1);
          let _760_v108;
          _760_v108 = _module.D9.create_DC24(_module.D9.create_DC24(_759_v107));
          let _761_v109;
          let _nw122 = Array((new BigNumber(11)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _module.D9.create_DC24(_759_v107);
          _nw122[(_dafny.ONE).toNumber()] = _760_v108;
          _nw122[(new BigNumber(2)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(3)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(4)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(5)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(6)).toNumber()] = _module.D9.create_DC24(_module.D9.create_DC23(_695_cf17, _694_cf18));
          _nw122[(new BigNumber(7)).toNumber()] = _module.D9.create_DC24(_module.D9.create_DC24(_module.D9.create_DC23(new BigNumber(722), _694_cf18)));
          _nw122[(new BigNumber(8)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(9)).toNumber()] = _760_v108;
          _nw122[(new BigNumber(10)).toNumber()] = _760_v108;
          _761_v109 = _nw122;
          let _index116 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_761_v109).length));
          (_761_v109)[_index116] = _760_v108;
          let _762_v110;
          _762_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_753_v102)[_module.__default.safeIndex(new BigNumber(278), new BigNumber((_753_v102).length))]);
          let _763_v111;
          _763_v111 = _dafny.Seq.of(_762_v110);
          let _index117 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_761_v109).length));
          (_761_v109)[_index117] = (((_this).fm0(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber((_dafny.Seq.UnicodeFromString("kmql")).length)), _dafny.Seq.UnicodeFromString("rbvsweaf"), _763_v111, globalState)) ? (((!(p1)) ? (_760_v108) : (_760_v108))) : (_760_v108));
        }
        let _764_v112;
        let _nw123 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _764_v112 = _nw123;
        let _765_v113;
        _765_v113 = _dafny.Seq.of(p0, (_this).f20);
        let _index118 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_764_v112).length));
        (_764_v112)[_index118] = _765_v113;
        let _index119 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_764_v112).length));
        (_764_v112)[_index119] = _765_v113;
      } else {
        let _766___mcc_h5 = (_source9).cf14;
        let _767_cf14 = _766___mcc_h5;
        let _768_v114;
        _768_v114 = _dafny.Seq.of((_this).f25, p0);
        let _769_v115;
        _769_v115 = _dafny.Set.fromElements(((p1) ? (_768_v114) : (_768_v114)));
        _769_v115 = _769_v115;
        let _770_v116;
        let _nw124 = new _module.C2();
        _nw124.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("ngcig")).length), (_this).f26);
        _770_v116 = _nw124;
        let _771_v117;
        _771_v117 = new _dafny.CodePoint('a'.codePointAt(0));
        _771_v117 = _771_v117;
        let _772_v118;
        _772_v118 = _dafny.Set.fromElements(p1);
        let _773_v119;
        _773_v119 = _dafny.Seq.of(_dafny.Set.fromElements(p1, (_this).f26));
        let _774_v120;
        _774_v120 = _module.D3.create_DC7(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_772_v118, _772_v118), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.of(_772_v118, _772_v118)).length)), _772_v118), _773_v119), p1);
        let _source11 = _774_v120;
        if (_source11.is_DC6) {
          let _775___mcc_h11 = (_source11).cf10;
          let _776_cf10 = _775___mcc_h11;
          let _777_v121;
          _777_v121 = _dafny.Seq.of(p1);
          let _778_v122;
          let _init22 = function (_779_i13) {
            return _dafny.Seq.UnicodeFromString("bty");
          };
          let _nw125 = Array((new BigNumber(15)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw125.length); _i0_22++) {
            _nw125[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _778_v122 = _nw125;
          let _780_v123;
          _780_v123 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p0);
          let _781_v124;
          _781_v124 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,true);
          let _782_v125;
          _782_v125 = _dafny.Seq.of(_781_v124);
          let _783_v126;
          _783_v126 = _dafny.Seq.of(_782_v125);
          let _rhs110 = (new BigNumber(332)).minus((((_this).fm0(_780_v123, _dafny.Seq.UnicodeFromString("ivwxw"), (_783_v126)[_module.__default.safeIndex((_this).f20, new BigNumber((_783_v126).length))], globalState)) ? (p0) : (new BigNumber(-46))));
          let _rhs111 = _777_v121;
          let _rhs112 = _778_v122;
          let _rhs113 = (_this).f26;
          let _lhs96 = globalState;
          let _lhs97 = globalState;
          _lhs96.f17 = _rhs110;
          _777_v121 = _rhs111;
          _778_v122 = _rhs112;
          _lhs97.f4 = _rhs113;
          let _784_v128;
          _784_v128 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _785_v129;
          _785_v129 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_768_v114).length),(((_784_v128).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f20)).length))) ? ((_784_v128).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f20)).length))) : (p1)));
          let _786_v131;
          _786_v131 = _dafny.Seq.UnicodeFromString("kdacsl");
          let _787_v132;
          _787_v132 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll25 = new _dafny.Set();
            for (const _compr_25 of (_785_v129).Keys.Elements) {
              let _788_v130 = _compr_25;
              if ((_785_v129).contains(_788_v130)) {
                _coll25.add(_module.__default.safeModuloInt(_788_v130, new BigNumber(486)));
              }
            }
            return _coll25;
          }(),_786_v131);
          let _789_v133;
          let _nw126 = Array((new BigNumber(14)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).f20, p0);
          _nw126[(_dafny.ONE).toNumber()] = new BigNumber(((_781_v124).Merge(_781_v124)).length);
          _nw126[(new BigNumber(2)).toNumber()] = new BigNumber((_772_v118).length);
          _nw126[(new BigNumber(3)).toNumber()] = (_this).f20;
          _nw126[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw126[(new BigNumber(5)).toNumber()] = _776_cf10;
          _nw126[(new BigNumber(6)).toNumber()] = (_this).f25;
          _nw126[(new BigNumber(7)).toNumber()] = p0;
          _nw126[(new BigNumber(8)).toNumber()] = new BigNumber((_780_v123).length);
          _nw126[(new BigNumber(9)).toNumber()] = (_this).f20;
          _nw126[(new BigNumber(10)).toNumber()] = (new BigNumber((function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_787_v132).Keys.Elements) {
              let _790_v127 = _compr_26;
              if ((_787_v132).contains(_790_v127)) {
                _coll26.push([_790_v127,new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)]);
              }
            }
            return _coll26;
          }()).length)).plus((_this).f25);
          _nw126[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("pttigqy")).length);
          _nw126[(new BigNumber(12)).toNumber()] = p0;
          _nw126[(new BigNumber(13)).toNumber()] = p0;
          _789_v133 = _nw126;
          let _index120 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length));
          (_789_v133)[_index120] = _776_cf10;
          let _index121 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length));
          let _rhs114 = !((_772_v118).IsDisjointFrom(_772_v118));
          let _rhs115 = (_770_v116).fm1(_776_cf10, (_this).f20, (_this).f21, globalState);
          let _lhs98 = globalState;
          let _lhs99 = _789_v133;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length));
          _lhs98.f18 = _rhs114;
          _lhs99[_lhs100] = _rhs115;
          let _rhs116 = (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f21)).update(p1, !((_this).f26));
          let _rhs117 = (_dafny.ZERO).minus((_789_v133)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length))]);
          let _rhs118 = _786_v131;
          let _rhs119 = p1;
          let _rhs120 = !(_dafny.Map.Empty.slice().updateUnsafe(false,true)).equals(_module.__default.fm15((_789_v133)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length))], (_789_v133)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length))], new BigNumber(953), new BigNumber((_786_v131).length), globalState));
          let _lhs101 = globalState;
          let _lhs102 = globalState;
          _781_v124 = _rhs116;
          _776_cf10 = _rhs117;
          _786_v131 = _rhs118;
          _lhs101.f3 = _rhs119;
          _lhs102.f3 = _rhs120;
          let _index122 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_789_v133).length));
          (_789_v133)[_index122] = (_dafny.ZERO).minus(p0);
        } else if (_source11.is_DC7) {
          let _791___mcc_h12 = (_source11).cf11;
          let _792___mcc_h13 = (_source11).cf12;
          let _793_cf12 = _792___mcc_h13;
          let _794_cf11 = _791___mcc_h12;
          let _795_v134;
          let _nw127 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _795_v134 = _nw127;
          let _index123 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length));
          (_795_v134)[_index123] = (_this).f20;
          let _796_v135;
          let _init23 = ((_797_p0, _798_cf12) => function (_799_i14) {
            return _module.D2.create_DC3(_797_p0, (_this).f26, _798_cf12, new BigNumber(593), (_this).f26);
          })(p0, _793_cf12);
          let _nw128 = Array((new BigNumber(19)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw128.length); _i0_23++) {
            _nw128[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _796_v135 = _nw128;
          let _800_v136;
          _800_v136 = _module.D2.create_DC3((_this).f25, p1, false, (_this).f25, (_this).f21);
          let _index124 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_796_v135).length));
          (_796_v135)[_index124] = _800_v136;
          let _801_v137;
          _801_v137 = _dafny.Map.Empty.slice().updateUnsafe(p1,_770_v116);
          let _index125 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length));
          let _index126 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_796_v135).length));
          let _rhs121 = new BigNumber((_801_v137).length);
          let _rhs122 = _800_v136;
          let _rhs123 = p1;
          let _rhs124 = (_768_v114)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_768_v114).length))];
          let _lhs103 = _795_v134;
          let _lhs104 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length));
          let _lhs105 = _796_v135;
          let _lhs106 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_796_v135).length));
          let _lhs107 = globalState;
          let _lhs108 = globalState;
          _lhs103[_lhs104] = _rhs121;
          _lhs105[_lhs106] = _rhs122;
          _lhs107.f3 = _rhs123;
          _lhs108.f17 = _rhs124;
          let _802_v138;
          _802_v138 = _dafny.Seq.UnicodeFromString("qgfualrx");
          (globalState).f5 = (new BigNumber((_dafny.Seq.Concat(_802_v138, _802_v138)).length)).plus(new BigNumber((_dafny.Seq.update(((_793_cf12) ? (_dafny.Seq.of(new BigNumber(756))) : (_768_v114)), _module.__default.safeIndex((_768_v114)[_module.__default.safeIndex((_this).f20, new BigNumber((_768_v114).length))], new BigNumber((((_793_cf12) ? (_dafny.Seq.of(new BigNumber(756))) : (_768_v114))).length)), (_795_v134)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length))])).length));
          let _803_v139;
          _803_v139 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f26);
          let _804_v140;
          _804_v140 = _dafny.Map.Empty.slice().updateUnsafe(_802_v138,_803_v139);
          let _805_v141;
          _805_v141 = _dafny.Seq.of((_this).f26);
          let _806_v142;
          _806_v142 = _dafny.Map.Empty.slice().updateUnsafe(_804_v140,_dafny.Seq.IsPrefixOf(_805_v141, _805_v141));
          _806_v142 = (_806_v142).update(_dafny.Map.Empty.slice().updateUnsafe(_802_v138,_803_v139), (_this).f26);
          let _index127 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_767_cf14).length));
          (_767_cf14)[_index127] = _793_cf12;
          let _807_v143;
          _807_v143 = _dafny.Set.fromElements(_771_v117, _771_v117);
          let _index128 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_767_cf14).length));
          let _index129 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length));
          let _rhs125 = _793_cf12;
          let _rhs126 = ((_795_v134)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length))]).isLessThanOrEqualTo((_795_v134)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length))]);
          let _rhs127 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_795_v134)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length))]).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_802_v138, _802_v138), _module.__default.safeIndex(new BigNumber((_807_v143).length), new BigNumber((_dafny.Seq.Concat(_802_v138, _802_v138)).length)), _771_v117)).length))));
          let _lhs109 = globalState;
          let _lhs110 = _767_cf14;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_767_cf14).length));
          let _lhs112 = _795_v134;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_795_v134).length));
          _lhs109.f14 = _rhs125;
          _lhs110[_lhs111] = _rhs126;
          _lhs112[_lhs113] = _rhs127;
        } else if (_source11.is_DC5) {
          let _808___mcc_h14 = (_source11).cf9;
          let _809_cf9 = _808___mcc_h14;
          let _810_v144;
          _810_v144 = _dafny.Map.Empty.slice().updateUnsafe(p1,_768_v114);
          _810_v144 = (_810_v144).update(!(p1), _dafny.Seq.of(p0, p0, (_this).f25));
          (globalState).f5 = (_dafny.ZERO).minus((_this).f25);
          let _811_v145;
          _811_v145 = _dafny.Seq.UnicodeFromString("pd");
          let _812_v146;
          let _nw129 = Array((new BigNumber(2)).toNumber());
          _nw129[(_dafny.ZERO).toNumber()] = (_this).f20;
          _nw129[(_dafny.ONE).toNumber()] = (_this).f20;
          _812_v146 = _nw129;
          let _813_v147;
          _813_v147 = _dafny.Map.Empty.slice().updateUnsafe(_811_v145,_812_v146);
          let _814_v148;
          let _nw130 = new _module.C1();
          _nw130.__ctor(p0, (((_813_v147).contains(_811_v145)) ? ((_813_v147).get(_811_v145)) : (_812_v146)), _module.__default.safeDivisionInt((_this).f25, new BigNumber(656)), p1);
          _814_v148 = _nw130;
          let _815_v149;
          let _nw131 = new _module.C1();
          _nw131.__ctor(((_dafny.ZERO).minus((_this).f25)).multipliedBy(p0), (_814_v148).f28, ((_this).f25).plus(new BigNumber(316)), (_this).f26);
          _815_v149 = _nw131;
        } else {
          let _816___mcc_h15 = (_source11).cf13;
          let _817_cf13 = _816___mcc_h15;
          let _rhs128 = p0;
          let _rhs129 = true;
          let _rhs130 = (_770_v116).fm1(p0, (_this).f25, (_this).f26, globalState);
          let _lhs114 = globalState;
          let _lhs115 = globalState;
          let _lhs116 = globalState;
          _lhs114.f17 = _rhs128;
          _lhs115.f18 = _rhs129;
          _lhs116.f17 = _rhs130;
          let _818_v151;
          let _init24 = ((_819_p0) => function (_820_i15) {
            return function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of _dafny.IntegerRange(new BigNumber(439), new BigNumber(582))) {
                let _821_v150 = _compr_27;
                if (((new BigNumber(439)).isLessThanOrEqualTo(_821_v150)) && ((_821_v150).isLessThan(new BigNumber(582)))) {
                  _coll27.add((_821_v150).plus(_819_p0));
                }
              }
              return _coll27;
            }();
          })(p0);
          let _nw132 = Array((new BigNumber(3)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw132.length); _i0_24++) {
            _nw132[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _818_v151 = _nw132;
          let _822_v152;
          _822_v152 = _module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20));
          let _823_v153;
          _823_v153 = _dafny.Seq.UnicodeFromString("fvwgksww");
          let _824_v154;
          _824_v154 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),true);
          let _825_v155;
          _825_v155 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).update((_770_v116).fm0((_822_v152).dtor_cf28, _dafny.Seq.update(_823_v153, _module.__default.safeIndex((_770_v116).fm1((_this).f20, p0, p1, globalState), new BigNumber((_823_v153).length)), _771_v117), _dafny.Seq.of(_824_v154, _824_v154, _824_v154), globalState), (_this).f21),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_768_v114).length),(_this).f26));
          let _rhs131 = _818_v151;
          let _rhs132 = (_825_v155).Merge(_825_v155);
          let _rhs133 = (_this).f21;
          let _lhs117 = globalState;
          _818_v151 = _rhs131;
          _825_v155 = _rhs132;
          _lhs117.f3 = _rhs133;
          let _826_v156;
          _826_v156 = _dafny.Map.Empty.slice().updateUnsafe(_767_cf14,(_dafny.ZERO).minus((_this).f25));
          let _827_v157;
          _827_v157 = _dafny.Seq.of(_826_v156, _826_v156, (_826_v156).Merge(_826_v156), (_826_v156).Merge(_826_v156));
          _826_v156 = (_827_v157)[_module.__default.safeIndex(new BigNumber(-318), new BigNumber((_827_v157).length))];
          let _828_v158;
          _828_v158 = _module.D4.create_DC11(new BigNumber((_dafny.Seq.UnicodeFromString("rqebepvf")).length), p1, p1);
          let _829_v159;
          _829_v159 = _dafny.Map.Empty.slice().updateUnsafe(_772_v118,p0);
          let _830_v160;
          let _nw133 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _830_v160 = _nw133;
          let _831_v161;
          _831_v161 = _dafny.MultiSet.fromElements(_830_v160);
          let _832_v162;
          let _nw134 = Array((new BigNumber(12)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = (_828_v158).dtor_cf17;
          _nw134[(_dafny.ONE).toNumber()] = (_this).fm1((_this).f20, (_this).f25, p1, globalState);
          _nw134[(new BigNumber(2)).toNumber()] = p0;
          _nw134[(new BigNumber(3)).toNumber()] = ((((_829_v159).contains(_772_v118)) ? ((_829_v159).get(_772_v118)) : ((((_831_v161).contains(_830_v160)) ? ((_831_v161).get(_830_v160)) : ((_this).f25))))).multipliedBy((_this).f25);
          _nw134[(new BigNumber(4)).toNumber()] = new BigNumber(489);
          _nw134[(new BigNumber(5)).toNumber()] = p0;
          _nw134[(new BigNumber(6)).toNumber()] = p0;
          _nw134[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt((_this).f20, new BigNumber((_824_v154).length));
          _nw134[(new BigNumber(8)).toNumber()] = new BigNumber(-12);
          _nw134[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f20, (_this).f25)).length);
          _nw134[(new BigNumber(10)).toNumber()] = ((_this).f25).plus((_dafny.ZERO).minus(p0));
          _nw134[(new BigNumber(11)).toNumber()] = p0;
          _832_v162 = _nw134;
          let _index130 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_832_v162).length));
          (_832_v162)[_index130] = p0;
          let _index131 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_832_v162).length));
          (_832_v162)[_index131] = (_this).f20;
        }
      }
      let _833_v163;
      _833_v163 = _dafny.Seq.UnicodeFromString("uftkj");
      let _834_v164;
      _834_v164 = _dafny.Seq.of(_833_v163, (((_this).f21) ? (_833_v163) : (_833_v163)));
      let _835_v165;
      _835_v165 = _dafny.MultiSet.fromElements((_this).f25);
      let _rhs134 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber((_835_v165).cardinality()), (_this).f25), (_this).f20);
      let _rhs135 = (_dafny.ZERO).minus((((_this).f20).multipliedBy((_this).f25)).multipliedBy((_this).f20));
      let _rhs136 = _834_v164;
      let _rhs137 = (_this).f26;
      let _rhs138 = p0;
      let _lhs118 = globalState;
      let _lhs119 = globalState;
      let _lhs120 = globalState;
      let _lhs121 = globalState;
      _lhs118.f5 = _rhs134;
      _lhs119.f17 = _rhs135;
      _834_v164 = _rhs136;
      _lhs120.f14 = _rhs137;
      _lhs121.f5 = _rhs138;
      let _836_v166;
      _836_v166 = new _dafny.CodePoint('w'.codePointAt(0));
      let _837_v167;
      _837_v167 = _dafny.MultiSet.fromElements(p1);
      let _838_v168;
      _838_v168 = _dafny.Seq.of((_this).f21, (_this).f21);
      let _839_v169;
      _839_v169 = _dafny.Set.fromElements((_this).f26);
      let _840_v170;
      _840_v170 = _dafny.Seq.of((_this).f20, (_this).f20);
      let _841_v171;
      _841_v171 = _module.D6.create_DC13(_833_v163);
      let _842_v172;
      _842_v172 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-993),new BigNumber(((_841_v171).dtor_cf21).length));
      let _843_v174;
      _843_v174 = _module.D8.create_DC19(function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(377), new BigNumber(479))) {
    let _844_v173 = _compr_28;
    if (((new BigNumber(377)).isLessThanOrEqualTo(_844_v173)) && ((_844_v173).isLessThan(new BigNumber(479)))) {
      _coll28.push([_module.__default.safeDivisionInt(_844_v173, (_this).f20),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f20)).length)]);
    }
  }
  return _coll28;
}());
      let _845_v175;
      let _nw135 = Array((new BigNumber(27)).toNumber());
      _nw135[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw135[(_dafny.ONE).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(2)).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(3)).toNumber()] = (_this).f20;
      _nw135[(new BigNumber(4)).toNumber()] = (((_837_v167).contains((_this).f21)) ? ((_837_v167).get((_this).f21)) : (p0));
      _nw135[(new BigNumber(5)).toNumber()] = p0;
      _nw135[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw135[(new BigNumber(7)).toNumber()] = new BigNumber(489);
      _nw135[(new BigNumber(8)).toNumber()] = new BigNumber((_module.__default.fm18((_this).f20, (_this).f20, _836_v166, globalState)).length);
      _nw135[(new BigNumber(9)).toNumber()] = new BigNumber(248);
      _nw135[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(11)).toNumber()] = new BigNumber((_838_v168).length);
      _nw135[(new BigNumber(12)).toNumber()] = new BigNumber((_839_v169).length);
      _nw135[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("esaxrsb")).length);
      _nw135[(new BigNumber(14)).toNumber()] = new BigNumber(268);
      _nw135[(new BigNumber(15)).toNumber()] = p0;
      _nw135[(new BigNumber(16)).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(17)).toNumber()] = (_this).f20;
      _nw135[(new BigNumber(18)).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(19)).toNumber()] = p0;
      _nw135[(new BigNumber(20)).toNumber()] = new BigNumber((_840_v170).length);
      _nw135[(new BigNumber(21)).toNumber()] = (_this).f20;
      _nw135[(new BigNumber(22)).toNumber()] = new BigNumber(48);
      _nw135[(new BigNumber(23)).toNumber()] = (_this).f25;
      _nw135[(new BigNumber(24)).toNumber()] = new BigNumber((_842_v172).length);
      _nw135[(new BigNumber(25)).toNumber()] = new BigNumber((_838_v168).length);
      _nw135[(new BigNumber(26)).toNumber()] = new BigNumber(((_843_v174).dtor_cf28).length);
      _845_v175 = _nw135;
      let _846_v176;
      let _nw136 = new _module.C1();
      _nw136.__ctor((_this).f20, _845_v175, (_this).f20, (_this).f21);
      _846_v176 = _nw136;
      let _847_v177;
      _847_v177 = _dafny.Map.Empty.slice().updateUnsafe(_836_v166,_846_v176);
      let _848_v178;
      _848_v178 = _dafny.Seq.of(_847_v177);
      let _849_v179;
      _849_v179 = _module.D10.create_DC27(new BigNumber((_833_v163).length), _848_v178, new _dafny.CodePoint('p'.codePointAt(0)));
      let _850_v180;
      let _nw137 = Array((new BigNumber(16)).toNumber());
      _nw137[(_dafny.ZERO).toNumber()] = _836_v166;
      _nw137[(_dafny.ONE).toNumber()] = _836_v166;
      _nw137[(new BigNumber(2)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(3)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(4)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(5)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(6)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(7)).toNumber()] = (_849_v179).dtor_cf38;
      _nw137[(new BigNumber(8)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(9)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(10)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(11)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(12)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(13)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(14)).toNumber()] = _836_v166;
      _nw137[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
      _850_v180 = _nw137;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_850_v180).length))) {
        let _851_i16 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_851_i16)) && ((_851_i16).isLessThan(new BigNumber((_850_v180).length))))) {
          (_850_v180)[(_851_i16)] = (_833_v163)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_840_v170,_846_v176.f27)).length), new BigNumber((_833_v163).length))];
        }
      }
      _834_v164 = _module.__default.fm33(_846_v176.f27, (_this).f21, _dafny.Seq.Concat(_833_v163, _833_v163), globalState);
      let _852_i17;
      _852_i17 = _dafny.ZERO;
      L3: {
        while ((_this).f26) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_852_i17)) {
              break L3;
            }
            _852_i17 = (_852_i17).plus(_dafny.ONE);
            let _853_v181;
            _853_v181 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
            let _854_v182;
            _854_v182 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f26) || ((((_853_v181).contains(false)) ? ((_853_v181).get(false)) : ((_this).f21))),(new BigNumber(290)).minus((_this).f25));
            (_846_v176).f27 = (((_854_v182).contains((_this).f26)) ? ((_854_v182).get((_this).f26)) : ((((_842_v172).contains((_this).f20)) ? ((_842_v172).get((_this).f20)) : ((_this).f20))));
            (globalState).f5 = _module.__default.safeModuloInt(_846_v176.f27, _846_v176.f27);
            let _855_v183;
            _855_v183 = _dafny.Map.Empty.slice().updateUnsafe(_835_v165,p1);
            (globalState).f5 = _module.__default.safeModuloInt((_846_v176).fm1(_846_v176.f27, p0, (((_855_v183).contains(_835_v165)) ? ((_855_v183).get(_835_v165)) : ((_this).f26)), globalState), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,new BigNumber(793))).Merge(_854_v182)).length));
            let _856_v184;
            _856_v184 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f26);
            _856_v184 = (_856_v184).update(p0, (((_856_v184).contains(_846_v176.f27)) ? ((_856_v184).get(_846_v176.f27)) : ((_this).f21)));
          }
        }
      }
      (globalState).f14 = false;
      let _857_v185;
      _857_v185 = _dafny.Seq.of((_846_v176).f28);
      let _858_v186;
      _858_v186 = _dafny.Seq.of(_857_v185, _857_v185);
      r0 = (_858_v186)[_module.__default.safeIndex(((_this).f20).minus(new BigNumber(463)), new BigNumber((_858_v186).length))];
      r1 = true;
      return [r0, r1];
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _859_v0;
      _859_v0 = _dafny.Set.fromElements((_this).f26, (_this).f26);
      let _860_v1;
      _860_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_859_v0).length),new BigNumber(856));
      let _861_v2;
      _861_v2 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f20)),true);
      let _862_v3;
      _862_v3 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f20), new BigNumber((_861_v2).length), (_this).f20, (_this).f20);
      _860_v1 = (_860_v1).update(new BigNumber((_862_v3).length), (_dafny.ZERO).minus((_this).f25));
      let _863_v4;
      _863_v4 = _dafny.Seq.of((_this).f21, (_this).f26);
      if ((_863_v4)[_module.__default.safeIndex((_this).f20, new BigNumber((_863_v4).length))]) {
        let _864_v5;
        let _init25 = function (_865_i0) {
          return (((_this).f21) ? (_dafny.Seq.UnicodeFromString("rjsrayy")) : (_dafny.Seq.UnicodeFromString("asut")));
        };
        let _nw138 = Array((new BigNumber(9)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw138.length); _i0_25++) {
          _nw138[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _864_v5 = _nw138;
        let _866_v6;
        _866_v6 = _dafny.Seq.UnicodeFromString("qxyux");
        let _index132 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_864_v5).length));
        (_864_v5)[_index132] = _866_v6;
        let _867_v7;
        _867_v7 = new _dafny.CodePoint('l'.codePointAt(0));
        let _index133 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_864_v5).length));
        let _rhs139 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_868_i1) {
          return (_this).f25;
        }), _dafny.Seq.of((_this).f25))).length)).multipliedBy(new BigNumber(821)));
        let _rhs140 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_866_v6, _866_v6), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Concat(_866_v6, _866_v6)).length)), _867_v7), _866_v6);
        let _rhs141 = (_this).f25;
        let _lhs122 = globalState;
        let _lhs123 = _864_v5;
        let _lhs124 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_864_v5).length));
        let _lhs125 = globalState;
        _lhs122.f17 = _rhs139;
        _lhs123[_lhs124] = _rhs140;
        _lhs125.f17 = _rhs141;
        (globalState).f18 = (_this).f26;
        let _869_v8;
        _869_v8 = _module.D2.create_DC3((_this).f20, (_this).f26, !(!((_this).f21)), (_this).f20, (_this).f26);
        let _870_v9;
        _870_v9 = _dafny.Seq.of(_859_v0);
        let _871_v10;
        _871_v10 = _module.D3.create_DC7(_870_v9, (_this).f26);
        let _pat_let_tv22 = _870_v9;
        let _source12 = function (_pat_let22_0) {
          return function (_872_dt__update__tmp_h0) {
            return function (_pat_let23_0) {
              return function (_873_dt__update_hcf11_h0) {
                return _module.D3.create_DC7(_873_dt__update_hcf11_h0, (_872_dt__update__tmp_h0).dtor_cf12);
              }(_pat_let23_0);
            }(_pat_let_tv22);
          }(_pat_let22_0);
        }((((_869_v8).dtor_cf7) ? (_871_v10) : (_module.__default.fm34(new _dafny.CodePoint('v'.codePointAt(0)), (_this).f21, (_this).f20, (_this).f25, globalState))));
        if (_source12.is_DC6) {
          let _874___mcc_h0 = (_source12).cf10;
          let _875_cf10 = _874___mcc_h0;
          _867_v7 = _867_v7;
          let _index134 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((p1).length));
          (p1)[_index134] = (((_860_v1).contains(_875_cf10)) ? ((_860_v1).get(_875_cf10)) : (_875_cf10));
          let _876_v11;
          _876_v11 = _dafny.Set.fromElements((_this).f25);
          let _index135 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((p1).length));
          (p1)[_index135] = (_dafny.ZERO).minus(new BigNumber(((_module.__default.fm25(globalState)).Intersect(_876_v11)).length));
          let _877_v12;
          let _nw139 = Array((new BigNumber(4)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = (_this).f21;
          _nw139[(_dafny.ONE).toNumber()] = (_this).f26;
          _nw139[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw139[(new BigNumber(3)).toNumber()] = true;
          _877_v12 = _nw139;
          let _878_v13;
          _878_v13 = _dafny.Seq.of(_877_v12);
          let _879_v14;
          _879_v14 = _dafny.MultiSet.fromElements((_878_v13)[_module.__default.safeIndex((_module.D2.create_DC3((p1)[_module.__default.safeIndex(new BigNumber(342), new BigNumber((p1).length))], !((_this).f21), false, (_this).f20, (_this).f21)).dtor_cf3, new BigNumber((_878_v13).length))], _877_v12);
          (globalState).f4 = (_879_v14).IsSubsetOf(_879_v14);
          (globalState).f5 = (_this).f20;
        } else if (_source12.is_DC7) {
          let _880___mcc_h1 = (_source12).cf11;
          let _881___mcc_h2 = (_source12).cf12;
          let _882_cf12 = _881___mcc_h2;
          let _883_cf11 = _880___mcc_h1;
          let _index136 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((p1).length));
          (p1)[_index136] = (_this).f20;
          let _884_v15;
          let _nw140 = new _module.C0();
          _nw140.__ctor();
          _884_v15 = _nw140;
          let _885_v16;
          _885_v16 = _dafny.Set.fromElements((_this).f25);
          let _886_v17;
          _886_v17 = _module.D4.create_DC11((_this).f25, _882_cf12, (_this).f26);
          let _887_v18;
          _887_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_886_v17);
          let _888_v19;
          _888_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_887_v18);
          let _index137 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((p1).length));
          let _rhs142 = !((_884_v15).fm3((_this).f20, (_this).f21, globalState));
          let _rhs143 = new BigNumber(((_888_v19).Merge(_888_v19)).length);
          let _rhs144 = _884_v15;
          let _rhs145 = (_885_v16).Union((_dafny.Set.fromElements((_this).f25)).Union(_885_v16));
          let _rhs146 = _882_cf12;
          let _lhs126 = globalState;
          let _lhs127 = p1;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((p1).length));
          let _lhs129 = globalState;
          _lhs126.f18 = _rhs142;
          _lhs127[_lhs128] = _rhs143;
          _884_v15 = _rhs144;
          _885_v16 = _rhs145;
          _lhs129.f3 = _rhs146;
          (globalState).f17 = (_this).f20;
          (globalState).f17 = (new BigNumber((_dafny.Seq.UnicodeFromString("ysfj")).length)).plus((_this).f25);
          let _index138 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((p0).length));
          (p0)[_index138] = _860_v1;
          let _index139 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((p0).length));
          (p0)[_index139] = _860_v1;
        } else if (_source12.is_DC5) {
          let _889___mcc_h3 = (_source12).cf9;
          let _890_cf9 = _889___mcc_h3;
          let _index140 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((p1).length));
          (p1)[_index140] = (_this).f25;
          let _index141 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((p1).length));
          (p1)[_index141] = new BigNumber((_862_v3).length);
          (globalState).f17 = new BigNumber(-305);
          let _891_v20;
          let _nw141 = new _module.C0();
          _nw141.__ctor();
          _891_v20 = _nw141;
          let _892_v21;
          _892_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_891_v20);
          let _893_v22;
          _893_v22 = _dafny.Set.fromElements(_891_v20, (((_892_v21).contains((_this).f25)) ? ((_892_v21).get((_this).f25)) : (_891_v20)), _891_v20);
          let _rhs147 = _893_v22;
          let _rhs148 = (_this).f21;
          let _rhs149 = (((_860_v1).contains((_this).f20)) ? ((_860_v1).get((_this).f20)) : ((_this).f25));
          let _rhs150 = (_this).f25;
          let _rhs151 = (((_this).f21) ? ((_862_v3)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((p1).length))], new BigNumber((_862_v3).length))]) : (((_this).f25).minus((p1)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((p1).length))])));
          let _lhs130 = globalState;
          let _lhs131 = globalState;
          let _lhs132 = globalState;
          let _lhs133 = globalState;
          _893_v22 = _rhs147;
          _lhs130.f3 = _rhs148;
          _lhs131.f5 = _rhs149;
          _lhs132.f17 = _rhs150;
          _lhs133.f17 = _rhs151;
          (globalState).f4 = (_this).f26;
        } else {
          let _894___mcc_h4 = (_source12).cf13;
          let _895_cf13 = _894___mcc_h4;
          (globalState).f17 = (_dafny.ZERO).minus(new BigNumber((_860_v1).length));
          let _index142 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_864_v5).length));
          (_864_v5)[_index142] = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm17((_this).f26, globalState)), _dafny.Seq.Concat(_dafny.Seq.of(_867_v7, new _dafny.CodePoint('u'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-627)), ((_896_v7) => function (_897_i2) {
            return _896_v7;
          })(_867_v7))));
          let _898_v23;
          let _nw142 = Array((new BigNumber(3)).toNumber()).fill(false);
          _898_v23 = _nw142;
          let _init26 = function (_899_i3) {
            return ((_this).f25).isLessThan((_this).f20);
          };
          let _nw143 = Array((new BigNumber(21)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw143.length); _i0_26++) {
            _nw143[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _898_v23 = _nw143;
          let _index143 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p1).length));
          (p1)[_index143] = (_this).f20;
          let _index144 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p1).length));
          (p1)[_index144] = (_this).fm1((_dafny.ZERO).minus((_this).fm1((_this).f20, (_this).f25, !((_this).f26), globalState)), new BigNumber(900), false, globalState);
        }
        (globalState).f17 = (_this).f25;
        _866_v6 = (_864_v5)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_864_v5).length))];
      } else {
        let _900_v24;
        let _nw144 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _900_v24 = _nw144;
        let _901_v25;
        let _nw145 = Array((new BigNumber(15)).toNumber()).fill(false);
        _901_v25 = _nw145;
        let _index145 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_901_v25).length));
        (_901_v25)[_index145] = (_this).f26;
        let _902_v26;
        _902_v26 = _dafny.MultiSet.fromElements((_this).f21);
        let _903_v27;
        _903_v27 = _module.D11.create_DC31(_module.D11.create_DC30((_this).f26, (_902_v26).update(true, _module.__default.abs((_this).f20))));
        let _904_v28;
        _904_v28 = _dafny.Map.Empty.slice().updateUnsafe(_903_v27,new BigNumber((_dafny.Seq.of((_this).f25)).length));
        let _905_v29;
        _905_v29 = _dafny.Seq.of((_904_v28).Merge(_904_v28), _904_v28);
        let _906_v30;
        _906_v30 = _dafny.Seq.UnicodeFromString("uielcygj");
        let _index146 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_901_v25).length));
        let _rhs152 = _900_v24;
        let _rhs153 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ie"), _dafny.Seq.Concat(_906_v30, _906_v30));
        let _rhs154 = false;
        let _rhs155 = _905_v29;
        let _rhs156 = _906_v30;
        let _lhs134 = globalState;
        let _lhs135 = _901_v25;
        let _lhs136 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_901_v25).length));
        let _lhs137 = globalState;
        _900_v24 = _rhs152;
        _lhs134.f15 = _rhs153;
        _lhs135[_lhs136] = _rhs154;
        _905_v29 = _rhs155;
        _lhs137.f15 = _rhs156;
        (globalState).f4 = (_901_v25)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_901_v25).length))];
        let _index147 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((p1).length));
        (p1)[_index147] = (_this).f25;
        let _index148 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((p1).length));
        (p1)[_index148] = (_this).f25;
        let _907_v31;
        let _nw146 = Array((new BigNumber(27)).toNumber());
        _907_v31 = _nw146;
        let _908_v32;
        let _nw147 = new _module.C2();
        _nw147.__ctor(new BigNumber((_906_v30).length), (_901_v25)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_901_v25).length))]);
        _908_v32 = _nw147;
        let _index149 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_907_v31).length));
        (_907_v31)[_index149] = _908_v32;
        let _index150 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_907_v31).length));
        (_907_v31)[_index150] = _908_v32;
        let _909_v33;
        _909_v33 = _dafny.Set.fromElements(_860_v1, _dafny.Map.Empty.slice().updateUnsafe((_908_v32).f20,(_this).f20), _860_v1, _860_v1);
        let _910_v34;
        let _nw148 = new _module.C1();
        _nw148.__ctor((_this).f25, _900_v24, new BigNumber((_909_v33).length), (_this).f26);
        _910_v34 = _nw148;
      }
      let _911_v35;
      _911_v35 = new _dafny.CodePoint('b'.codePointAt(0));
      let _912_v36;
      _912_v36 = _dafny.Map.Empty.slice().updateUnsafe((!(!((_this).f26))) && ((_this).f21),_911_v35);
      _912_v36 = (_912_v36).update((_this).f21, _911_v35);
      let _913_v37;
      _913_v37 = _dafny.Seq.of(p1, p1, (((_this).f26) ? (p1) : (p1)), p1, p1);
      _913_v37 = _913_v37;
      let _914_v38;
      _914_v38 = _dafny.MultiSet.fromElements((_this).f21, (_this).f26);
      let _source13 = _module.D11.create_DC30((_this).f21, _914_v38);
      if (_source13.is_DC29) {
        let _915_v39;
        _915_v39 = _dafny.Seq.UnicodeFromString("yakhfedov");
        let _916_v40;
        _916_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,false);
        let _917_v41;
        _917_v41 = _dafny.Seq.of(_916_v40, _916_v40, _dafny.Map.Empty.slice().updateUnsafe((_this).f26,false), _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21));
        let _rhs157 = (_this).f21;
        let _rhs158 = (_this).fm0(_860_v1, _915_v39, _dafny.Seq.Concat(_917_v41, _917_v41), globalState);
        let _lhs138 = globalState;
        r0 = _rhs157;
        _lhs138.f4 = _rhs158;
        (globalState).f17 = (_this).f25;
        let _918_v42;
        let _nw149 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _918_v42 = _nw149;
        let _919_v43;
        let _nw150 = Array((new BigNumber(2)).toNumber());
        _nw150[(_dafny.ZERO).toNumber()] = _918_v42;
        _nw150[(_dafny.ONE).toNumber()] = _918_v42;
        _919_v43 = _nw150;
        let _index151 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_919_v43).length));
        (_919_v43)[_index151] = _918_v42;
        let _920_v44;
        _920_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber(495));
        let _921_v45;
        _921_v45 = _dafny.Map.Empty.slice().updateUnsafe(_920_v44,_918_v42);
        let _922_v46;
        _922_v46 = _dafny.Seq.of(_918_v42, (((_921_v45).contains(_920_v44)) ? ((_921_v45).get(_920_v44)) : (_918_v42)));
        let _index152 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_919_v43).length));
        (_919_v43)[_index152] = (_922_v46)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_915_v39).length), (_this).f20), new BigNumber((_922_v46).length))];
        let _923_v47;
        _923_v47 = _dafny.MultiSet.fromElements((((_920_v44).contains((_this).f26)) ? ((_920_v44).get((_this).f26)) : ((_this).f25)));
        if (!(_923_v47).contains((_this).f20)) {
          let _924_v48;
          _924_v48 = _dafny.MultiSet.fromElements(_911_v35);
          let _925_v49;
          _925_v49 = _module.D7.create_DC18(_dafny.Seq.of(new BigNumber((_915_v39).length)), (_this).f20);
          let _926_v50;
          let _nw151 = new _module.C1();
          _nw151.__ctor(new BigNumber((_924_v48).cardinality()), p1, new BigNumber(((_925_v49).dtor_cf26).length), !_dafny.areEqual(_915_v39, _915_v39));
          _926_v50 = _nw151;
          _911_v35 = _911_v35;
          let _927_v51;
          _927_v51 = _dafny.Set.fromElements((_this).f20, new BigNumber(438), (_this).f25, new BigNumber(-926), (_this).f20);
          let _index153 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((p1).length));
          (p1)[_index153] = (_this).f20;
          let _index154 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((p1).length));
          let _rhs159 = new BigNumber((_module.__default.fm35((_this).f26, (_this).f21, (_this).fm1((_dafny.ZERO).minus(_926_v50.f27), new BigNumber(981), (_this).f26, globalState), globalState)).length);
          let _rhs160 = (((_this).f26) === ((_863_v4)[_module.__default.safeIndex((_this).f20, new BigNumber((_863_v4).length))])) || ((_this).f21);
          let _rhs161 = (_927_v51).Union(_module.__default.fm25(globalState));
          let _rhs162 = _926_v50.f27;
          let _rhs163 = (((_926_v50).fm1(_926_v50.f27, (_this).f25, (_this).f21, globalState)).plus((_this).f25)).isLessThan((_this).f25);
          let _lhs139 = globalState;
          let _lhs140 = p1;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((p1).length));
          let _lhs142 = globalState;
          _lhs139.f5 = _rhs159;
          r0 = _rhs160;
          _927_v51 = _rhs161;
          _lhs140[_lhs141] = _rhs162;
          _lhs142.f3 = _rhs163;
          r0 = ((_this).f21) && (!_dafny.areEqual(_915_v39, _915_v39));
          let _928_v52;
          let _nw152 = new _module.C0();
          _nw152.__ctor();
          _928_v52 = _nw152;
        } else {
          let _929_v53;
          let _nw153 = Array((new BigNumber(16)).toNumber()).fill(false);
          _929_v53 = _nw153;
          let _index155 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_929_v53).length));
          (_929_v53)[_index155] = (_this).f26;
          let _index156 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_929_v53).length));
          (_929_v53)[_index156] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_915_v39, _915_v39), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Concat(_915_v39, _915_v39)).length)), _911_v35), _dafny.Seq.update(_915_v39, _module.__default.safeIndex((_this).f25, new BigNumber((_915_v39).length)), _911_v35));
          (globalState).f17 = _module.__default.safeModuloInt((new BigNumber(-536)).plus((_this).f20), (_this).f25);
          let _930_v54;
          _930_v54 = _dafny.Seq.of(_929_v53);
          _930_v54 = _930_v54;
          let _index157 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length));
          (p1)[_index157] = ((_this).f20).minus((_this).f20);
          let _index158 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length));
          let _rhs164 = (_this).f26;
          let _rhs165 = ((_this).f20).isEqualTo((_this).f20);
          let _rhs166 = _859_v0;
          let _rhs167 = (((_860_v1).contains(((_862_v3)[_module.__default.safeIndex((_this).f25, new BigNumber((_862_v3).length))]).multipliedBy((_this).f20))) ? ((_860_v1).get(((_862_v3)[_module.__default.safeIndex((_this).f25, new BigNumber((_862_v3).length))]).multipliedBy((_this).f20))) : (new BigNumber(961)));
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          let _lhs145 = p1;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length));
          _lhs143.f3 = _rhs164;
          _lhs144.f4 = _rhs165;
          _859_v0 = _rhs166;
          _lhs145[_lhs146] = _rhs167;
          _862_v3 = (((_929_v53)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_929_v53).length))]) ? (_862_v3) : (_dafny.Seq.Concat(_dafny.Seq.update(_862_v3, _module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((p1).length))], new BigNumber((_862_v3).length)), (_this).f25), _862_v3)));
        }
      } else if (_source13.is_DC30) {
        let _931___mcc_h5 = (_source13).cf40;
        let _932___mcc_h6 = (_source13).cf41;
        let _933_cf41 = _932___mcc_h6;
        let _934_cf40 = _931___mcc_h5;
        let _935_v55;
        let _init27 = function (_936_i4) {
          return (_this).f26;
        };
        let _nw154 = Array((new BigNumber(21)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw154.length); _i0_27++) {
          _nw154[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _935_v55 = _nw154;
        let _937_v56;
        _937_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_935_v55);
        let _938_v57;
        _938_v57 = _module.D8.create_DC21(new BigNumber(((_937_v56).Merge(_937_v56)).length));
        _938_v57 = _938_v57;
        let _939_v58;
        _939_v58 = _dafny.MultiSet.fromElements((_this).f25);
        _860_v1 = (_860_v1).update(new BigNumber((_939_v58).cardinality()), (_this).f25);
        if ((_this).f21) {
          let _940_v59;
          _940_v59 = _dafny.Map.Empty.slice().updateUnsafe(_911_v35,(_this).f25);
          let _index159 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length));
          (p1)[_index159] = ((_this).f20).plus((((_940_v59).contains(_911_v35)) ? ((_940_v59).get(_911_v35)) : ((_this).f20)));
          let _index160 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length));
          (p1)[_index160] = ((_this).f20).multipliedBy(((_this).f25).multipliedBy((_this).f20));
          let _941_v60;
          _941_v60 = _dafny.Set.fromElements((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))], (_this).f20, (_this).f25, (_this).f25, (_dafny.ZERO).minus((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))]));
          let _942_v61;
          _942_v61 = _dafny.Seq.UnicodeFromString("hwv");
          _941_v60 = _dafny.Set.fromElements(((((_939_v58).contains(new BigNumber((_942_v61).length))) ? ((_939_v58).get(new BigNumber((_942_v61).length))) : (new BigNumber((_860_v1).length)))).minus((_this).f20), new BigNumber(221), (_this).f20, new BigNumber(123), (_this).f20);
          let _943_v62;
          _943_v62 = _module.D8.create_DC20();
          _943_v62 = _module.__default.fm36(globalState);
          let _index161 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length));
          (p1)[_index161] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,(p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))])).length))), _module.__default.safeModuloInt((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))], new BigNumber((_863_v4).length)));
          let _944_v63;
          let _nw155 = new _module.C2();
          _nw155.__ctor(((_this).f25).multipliedBy((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))]), (((_861_v2).contains((_this).fm1((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))], (_this).f20, true, globalState))) ? ((_861_v2).get((_this).fm1((p1)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((p1).length))], (_this).f20, true, globalState))) : (_934_cf40)));
          _944_v63 = _nw155;
        } else {
          (globalState).f17 = ((_this).f25).minus(_module.__default.safeModuloInt((_this).f25, (_this).f20));
          let _945_v64;
          _945_v64 = _dafny.Set.fromElements((_this).f20, (_this).f20, (_this).f25, _module.__default.safeModuloInt((_this).f20, new BigNumber(752)), (new BigNumber(129)).plus((_this).f25));
          let _946_v65;
          let _nw156 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _946_v65 = _nw156;
          let _index162 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_946_v65).length));
          (_946_v65)[_index162] = _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f25));
          let _index163 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_946_v65).length));
          let _rhs168 = function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of _dafny.IntegerRange(new BigNumber(632), new BigNumber(-38))) {
              let _947_v66 = _compr_29;
              if (((new BigNumber(632)).isLessThanOrEqualTo(_947_v66)) && ((_947_v66).isLessThan(new BigNumber(-38)))) {
                _coll29.add(_module.__default.safeDivisionInt(_947_v66, new BigNumber((_863_v4).length)));
              }
            }
            return _coll29;
          }();
          let _rhs169 = _939_v58;
          let _rhs170 = new BigNumber(-424);
          let _lhs147 = _946_v65;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_946_v65).length));
          let _lhs149 = globalState;
          _945_v64 = _rhs168;
          _lhs147[_lhs148] = _rhs169;
          _lhs149.f17 = _rhs170;
          let _948_v67;
          _948_v67 = _dafny.Seq.of(_dafny.Seq.of((_this).f20, new BigNumber(860)));
          _948_v67 = _948_v67;
          let _949_v68;
          _949_v68 = _dafny.Map.Empty.slice().updateUnsafe(_863_v4,(_this).f20);
          _949_v68 = (_949_v68).update(_dafny.Seq.of((((_861_v2).contains(new BigNumber((_862_v3).length))) ? ((_861_v2).get(new BigNumber((_862_v3).length))) : ((_this).f21)), (_this).f26), new BigNumber((_dafny.Seq.of(_934_cf40, (_this).f26, (_this).f26, false)).length));
          let _950_v69;
          _950_v69 = _dafny.MultiSet.fromElements(_863_v4, _dafny.Seq.Concat(_863_v4, _863_v4));
          let _951_v70;
          _951_v70 = _dafny.Seq.of(_863_v4, _863_v4);
          let _952_v71;
          _952_v71 = _dafny.Seq.of(_863_v4, _863_v4, (_951_v70)[_module.__default.safeIndex((_this).f25, new BigNumber((_951_v70).length))], _863_v4);
          let _953_v72;
          _953_v72 = _module.D12.create_DC32(_952_v71);
          _950_v69 = _dafny.MultiSet.FromArray((_953_v72).dtor_cf43);
        }
        _861_v2 = (_861_v2).update((_this).f20, !(!((_this).f25).isEqualTo((_this).f25)));
      } else if (_source13.is_DC28) {
        let _954___mcc_h7 = (_source13).cf39;
        let _955_cf39 = _954___mcc_h7;
        let _956_v73;
        _956_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f25);
        let _957_v74;
        let _nw157 = new _module.C1();
        _nw157.__ctor(_module.__default.safeModuloInt((_this).f25, (_this).f20), p1, (_dafny.ZERO).minus((((_956_v73).contains((_this).f26)) ? ((_956_v73).get((_this).f26)) : ((_this).f20))), ((_this).f26) || ((_this).f26));
        _957_v74 = _nw157;
        _957_v74 = _957_v74;
        let _958_v75;
        _958_v75 = (_957_v74).f28;
        let _959_v76;
        _959_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_863_v4).length),_958_v75);
        let _960_v77;
        _960_v77 = _dafny.Seq.of(_959_v76);
        (_957_v74).f27 = (_module.__default.safeModuloInt(_957_v74.f27, (_dafny.ZERO).minus(new BigNumber(((_960_v77)[_module.__default.safeIndex((_this).f25, new BigNumber((_960_v77).length))]).length)))).plus(_module.__default.safeDivisionInt(_957_v74.f27, new BigNumber(-289)));
        let _961_v78;
        let _nw158 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
        _961_v78 = _nw158;
        let _index164 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_961_v78).length));
        (_961_v78)[_index164] = _914_v38;
        let _index165 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_961_v78).length));
        (_961_v78)[_index165] = _dafny.MultiSet.fromElements(!((_this).f21));
      } else {
        let _962___mcc_h8 = (_source13).cf42;
        let _963_cf42 = _962___mcc_h8;
        let _964_v79;
        let _nw159 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _964_v79 = _nw159;
        let _index166 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_964_v79).length));
        (_964_v79)[_index166] = _863_v4;
        let _965_v80;
        _965_v80 = _dafny.Set.fromElements((_this).f25);
        let _966_v81;
        let _nw160 = Array((new BigNumber(12)).toNumber());
        _nw160[(_dafny.ZERO).toNumber()] = _911_v35;
        _nw160[(_dafny.ONE).toNumber()] = _911_v35;
        _nw160[(new BigNumber(2)).toNumber()] = _module.__default.fm17((_this).f26, globalState);
        _nw160[(new BigNumber(3)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(4)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(5)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(6)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(7)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(8)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(9)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(10)).toNumber()] = _911_v35;
        _nw160[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _966_v81 = _nw160;
        let _967_v82;
        _967_v82 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),(_this).f25);
        let _968_v83;
        _968_v83 = _dafny.Map.Empty.slice().updateUnsafe(_966_v81,_967_v82);
        let _index167 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_964_v79).length));
        let _rhs171 = _dafny.Seq.of((_this).f26, (_this).f26, false, (_965_v80).IsProperSubsetOf(_965_v80), (_this).f26);
        let _rhs172 = (((_this).f20).multipliedBy((_this).f20)).isLessThan((_862_v3)[_module.__default.safeIndex(new BigNumber((_862_v3).length), new BigNumber((_862_v3).length))]);
        let _rhs173 = (((_this).f21) ? (new BigNumber((_968_v83).length)) : (new BigNumber(153)));
        let _rhs174 = (_this).f25;
        let _rhs175 = _module.__default.safeModuloInt((_this).f25, (_dafny.ZERO).minus((_this).f25));
        let _lhs150 = _964_v79;
        let _lhs151 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_964_v79).length));
        let _lhs152 = globalState;
        let _lhs153 = globalState;
        let _lhs154 = globalState;
        let _lhs155 = globalState;
        _lhs150[_lhs151] = _rhs171;
        _lhs152.f14 = _rhs172;
        _lhs153.f17 = _rhs173;
        _lhs154.f5 = _rhs174;
        _lhs155.f17 = _rhs175;
        let _969_v85;
        _969_v85 = _dafny.Seq.UnicodeFromString("dr");
        if (!(!((_this).fm0((function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of _dafny.IntegerRange(new BigNumber(587), new BigNumber(-160))) {
            let _970_v84 = _compr_30;
            if (((new BigNumber(587)).isLessThanOrEqualTo(_970_v84)) && ((_970_v84).isLessThan(new BigNumber(-160)))) {
              _coll30.push([(_970_v84).multipliedBy((_this).f20),(_this).f25]);
            }
          }
          return _coll30;
        }()).Merge(_860_v1), _969_v85, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-466)), function (_971_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
        }), globalState)))) {
          (globalState).f17 = new BigNumber(88);
          let _972_v86;
          _972_v86 = _module.D6.create_DC14((_this).fm1((_this).f20, (_this).f20, (_this).f21, globalState));
          let _973_v87;
          _973_v87 = _module.D2.create_DC3(new BigNumber((_dafny.Seq.update(_969_v85, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("nw")).length), new BigNumber((_969_v85).length)), _911_v35)).length), (_this).f26, (_this).f26, (_this).f20, (_this).f21);
          let _974_v88;
          _974_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f21);
          let _975_v89;
          _975_v89 = _dafny.Seq.of(_974_v88);
          let _976_v90;
          _976_v90 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_863_v4).length),(_this).f25), _dafny.Seq.UnicodeFromString("rwox"), _975_v89, globalState),new BigNumber(798));
          let _977_v91;
          let _nw161 = Array((new BigNumber(23)).toNumber());
          _nw161[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_863_v4).length));
          _nw161[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), new BigNumber(-804));
          _nw161[(new BigNumber(2)).toNumber()] = (_972_v86).dtor_cf22;
          _nw161[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).f20, (_this).f20);
          _nw161[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(5)).toNumber()] = (_this).f20;
          _nw161[(new BigNumber(6)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(7)).toNumber()] = (((_this).f21) ? ((_this).f25) : ((_this).f20));
          _nw161[(new BigNumber(8)).toNumber()] = (_this).fm1(new BigNumber((_969_v85).length), (_this).f25, (_this).f26, globalState);
          _nw161[(new BigNumber(9)).toNumber()] = (_this).f20;
          _nw161[(new BigNumber(10)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt((_this).f25, (_this).f20);
          _nw161[(new BigNumber(12)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(13)).toNumber()] = (_973_v87).dtor_cf6;
          _nw161[(new BigNumber(14)).toNumber()] = (_this).f20;
          _nw161[(new BigNumber(15)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(16)).toNumber()] = (((_976_v90).contains((_this).f26)) ? ((_976_v90).get((_this).f26)) : ((_this).f20));
          _nw161[(new BigNumber(17)).toNumber()] = (_this).f20;
          _nw161[(new BigNumber(18)).toNumber()] = new BigNumber((_974_v88).length);
          _nw161[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus((_this).f25);
          _nw161[(new BigNumber(20)).toNumber()] = (_this).f25;
          _nw161[(new BigNumber(21)).toNumber()] = (((_this).f26) ? (new BigNumber((_859_v0).length)) : ((_this).f20));
          _nw161[(new BigNumber(22)).toNumber()] = (_this).f25;
          _977_v91 = _nw161;
          let _rhs176 = !((_this).f21);
          let _rhs177 = p1;
          let _lhs156 = globalState;
          _lhs156.f18 = _rhs176;
          _977_v91 = _rhs177;
          (globalState).f3 = ((_914_v38).Union(_module.__default.fm23(globalState))).IsDisjointFrom((_914_v38).Intersect(_dafny.MultiSet.FromArray(_module.__default.fm21(globalState))));
          let _nw162 = Array((new BigNumber(2)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = (_this).f20;
          _nw162[(_dafny.ONE).toNumber()] = (_this).f20;
          _977_v91 = _nw162;
          let _978_v92;
          let _nw163 = new _module.C0();
          _nw163.__ctor();
          _978_v92 = _nw163;
        } else {
          (globalState).f5 = (_this).f25;
          _969_v85 = _dafny.Seq.UnicodeFromString("tbh");
          let _979_v93;
          _979_v93 = p1;
          _979_v93 = _979_v93;
          (globalState).f5 = (_dafny.ZERO).minus(new BigNumber((_862_v3).length));
          let _980_v94;
          let _nw164 = new _module.C2();
          _nw164.__ctor(new BigNumber(340), (_this).f21);
          _980_v94 = _nw164;
          let _981_v95;
          _981_v95 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f26) ? (_980_v94) : (_980_v94)),(_dafny.ZERO).minus(((_this).fm1((((_860_v1).contains((_this).f20)) ? ((_860_v1).get((_this).f20)) : ((_this).f20)), (_this).f25, true, globalState)).multipliedBy((_980_v94).f20)));
          _981_v95 = (_981_v95).update((((_this).f21) ? (_980_v94) : (_980_v94)), (_980_v94).f20);
        }
        let _982_v96;
        _982_v96 = _module.D8.create_DC20();
        _982_v96 = _982_v96;
        let _983_v97;
        _983_v97 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25)));
        let _index168 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((p1).length));
        (p1)[_index168] = new BigNumber(((((_this).f21) ? ((_983_v97).update((_this).f20, _module.__default.abs(new BigNumber((_861_v2).length)))) : (_983_v97))).cardinality());
        let _index169 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((p1).length));
        (p1)[_index169] = (new BigNumber(-521)).multipliedBy((_this).f25);
      }
      let _984_v98;
      let _init28 = function (_985_i6) {
        return _module.D8.create_DC20();
      };
      let _nw165 = Array((new BigNumber(16)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw165.length); _i0_28++) {
        _nw165[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _984_v98 = _nw165;
      let _986_v99;
      let _init29 = ((_987_v35) => function (_988_i7) {
        return (_module.D13.create_DC35(_987_v35)).dtor_cf45;
      })(_911_v35);
      let _nw166 = Array((new BigNumber(14)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw166.length); _i0_29++) {
        _nw166[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _986_v99 = _nw166;
      let _989_v100;
      _989_v100 = _dafny.Seq.of(_986_v99);
      let _990_v101;
      _990_v101 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f20);
      let _991_v102;
      _991_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_862_v3);
      let _992_v103;
      _992_v103 = _dafny.Seq.UnicodeFromString("cntlqc");
      let _rhs178 = _dafny.areEqual(_dafny.Seq.Concat(_989_v100, _dafny.Seq.of(_986_v99, _986_v99, _986_v99, _986_v99, _986_v99)), _989_v100);
      let _rhs179 = _module.__default.safeModuloInt((((_990_v101).contains(p1)) ? ((_990_v101).get(p1)) : (new BigNumber(862))), (_this).f20);
      let _rhs180 = _984_v98;
      let _rhs181 = ((_module.D8.create_DC21((_this).f25)).dtor_cf29).plus(_module.__default.safeModuloInt(new BigNumber((_991_v102).length), (_dafny.ZERO).minus((_this).f20)));
      let _rhs182 = !((_this).f26) || ((new BigNumber((_992_v103).length)).isEqualTo((_this).f20));
      let _lhs157 = globalState;
      let _lhs158 = globalState;
      let _lhs159 = globalState;
      _lhs157.f14 = _rhs178;
      _lhs158.f5 = _rhs179;
      _984_v98 = _rhs180;
      _lhs159.f5 = _rhs181;
      r0 = _rhs182;
      let _993_v104;
      _993_v104 = _module.D8.create_DC19(_860_v1);
      let _994_v105;
      _994_v105 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f21);
      let _995_v106;
      _995_v106 = _dafny.Seq.of(_994_v105);
      r0 = (_this).fm0((_993_v104).dtor_cf28, _992_v103, _995_v106, globalState);
      return r0;
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f24, f20, f21) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f24;
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Set.fromElements((_this).f24)).length)).multipliedBy((_this).f20)).multipliedBy((_module.D4.create_DC11((_this).f20, (_this).f21, (_this).f24)).dtor_cf17);
    };
    fm13(p0, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f21)).length)).plus((_this).f20)).isEqualTo((_dafny.ZERO).minus((_this).f20));
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f20, (_this).f20));
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _996_v0;
      let _nw167 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
      _996_v0 = _nw167;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_996_v0).length))) {
        let _997_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_997_i0)) && ((_997_i0).isLessThan(new BigNumber((_996_v0).length))))) {
          (_996_v0)[(_997_i0)] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f21, (_this).f21), _dafny.Seq.Concat(_dafny.Seq.of((_this).f24, p1), _dafny.Seq.of((_this).f24)));
        }
      }
      let _998_v1;
      _998_v1 = _module.D4.create_DC10(_dafny.MultiSet.fromElements((_this).f21), (_this).f24);
      (globalState).f5 = new BigNumber((function (_source14) {
        if (_source14.is_DC10) {
          let _999___mcc_h0 = (_source14).cf15;
          let _1000___mcc_h1 = (_source14).cf16;
          let _1001_cf16 = _1000___mcc_h1;
          let _1002_cf15 = _999___mcc_h0;
          return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("lqfaswbj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-778)), function (_1003_i1) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), function (_1004_i2) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }));
        } else if (_source14.is_DC11) {
          let _1005___mcc_h2 = (_source14).cf17;
          let _1006___mcc_h3 = (_source14).cf18;
          let _1007___mcc_h4 = (_source14).cf19;
          let _1008_cf19 = _1007___mcc_h4;
          let _1009_cf18 = _1006___mcc_h3;
          let _1010_cf17 = _1005___mcc_h2;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_1011_i3) {
            return _dafny.Seq.UnicodeFromString("tlweiv");
          });
        } else {
          let _1012___mcc_h5 = (_source14).cf14;
          let _1013_cf14 = _1012___mcc_h5;
          return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cnq"));
        }
      }(_998_v1)).length);
      let _1014_v2;
      _1014_v2 = _module.D10.create_DC26(p0);
      let _1015_v3;
      _1015_v3 = _dafny.Set.fromElements(_1014_v2);
      let _1016_v4;
      _1016_v4 = _dafny.Seq.of(!(false));
      let _1017_v5;
      let _nw168 = Array((new BigNumber(9)).toNumber());
      _nw168[(_dafny.ZERO).toNumber()] = (_this).f20;
      _nw168[(_dafny.ONE).toNumber()] = new BigNumber((_1016_v4).length);
      _nw168[(new BigNumber(2)).toNumber()] = (_this).f20;
      _nw168[(new BigNumber(3)).toNumber()] = (_this).f20;
      _nw168[(new BigNumber(4)).toNumber()] = (_this).f20;
      _nw168[(new BigNumber(5)).toNumber()] = p0;
      _nw168[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw168[(new BigNumber(7)).toNumber()] = p0;
      _nw168[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Set.fromElements(p0, (_this).f20)).length);
      _1017_v5 = _nw168;
      let _1018_v6;
      let _nw169 = new _module.C1();
      _nw169.__ctor(_module.__default.safeModuloInt(new BigNumber(749), new BigNumber((_1015_v3).length)), _1017_v5, p0, ((p1) ? ((_this).f24) : ((_this).f24)));
      _1018_v6 = _nw169;
      _1018_v6 = _1018_v6;
      if (false) {
        let _1019_v7;
        _1019_v7 = _dafny.Seq.UnicodeFromString("qxdrfss");
        let _1020_v8;
        _1020_v8 = _module.D6.create_DC13(_dafny.Seq.Concat(_module.__default.fm26((_this).f21, globalState), _1019_v7));
        let _source15 = _1020_v8;
        if (_source15.is_DC14) {
          let _1021___mcc_h6 = (_source15).cf22;
          let _1022_cf22 = _1021___mcc_h6;
          let _1023_v9;
          let _nw170 = new _module.C1();
          _nw170.__ctor((_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f20), (_1018_v6).f20)), _1017_v5, new BigNumber((_1019_v7).length), (_1018_v6).f21);
          _1023_v9 = _nw170;
          (globalState).f5 = (_dafny.ZERO).minus((_1018_v6).f20);
          let _1024_v10;
          let _1025_v11;
          let _out14;
          let _out15;
          let _outcollector4 = (_this).m1(_1022_cf22, (_1018_v6).f21, _1018_v6, globalState);
          _out14 = _outcollector4[0];
          _out15 = _outcollector4[1];
          _1024_v10 = _out14;
          _1025_v11 = _out15;
          (_1023_v9).f27 = (_1018_v6).f20;
        } else if (_source15.is_DC15) {
          let _1026___mcc_h7 = (_source15).cf23;
          let _1027_cf23 = _1026___mcc_h7;
          let _1028_v12;
          _1028_v12 = _dafny.MultiSet.fromElements((_this).f20);
          let _1029_v13;
          _1029_v13 = _dafny.Seq.of((_1018_v6).f20, new BigNumber((_1019_v7).length), new BigNumber((_1019_v7).length), p0, (_1018_v6).f20);
          let _1030_v14;
          _1030_v14 = _dafny.MultiSet.fromElements(p1);
          let _1031_v15;
          _1031_v15 = _dafny.Seq.of((_1018_v6).f20, (_this).f20, (_this).fm14((_1018_v6).f21, (_1018_v6).f20, (((_1028_v12).contains((_1029_v13)[_module.__default.safeIndex((((_1028_v12).contains(new BigNumber((_1030_v14).cardinality()))) ? ((_1028_v12).get(new BigNumber((_1030_v14).cardinality()))) : ((_1018_v6).f20)), new BigNumber((_1029_v13).length))])) ? ((_1028_v12).get((_1029_v13)[_module.__default.safeIndex((((_1028_v12).contains(new BigNumber((_1030_v14).cardinality()))) ? ((_1028_v12).get(new BigNumber((_1030_v14).cardinality()))) : ((_1018_v6).f20)), new BigNumber((_1029_v13).length))])) : ((_this).f20)), p1, globalState));
          _1031_v15 = _1031_v15;
          let _1032_v16;
          let _nw171 = new _module.C0();
          _nw171.__ctor();
          _1032_v16 = _nw171;
          (globalState).f3 = !((_this).f21) || (((_1018_v6).f21) && ((_module.D13.create_DC36(p0, !(true), (_this).f24)).dtor_cf48));
          let _1033_v17;
          let _nw172 = new _module.C2();
          _nw172.__ctor(_module.__default.safeModuloInt((_1018_v6).f20, (_1018_v6).f20), !(p0).isEqualTo(new BigNumber((_1019_v7).length)));
          _1033_v17 = _nw172;
        } else {
          let _1034___mcc_h8 = (_source15).cf21;
          let _1035_cf21 = _1034___mcc_h8;
          let _1036_v18;
          let _init30 = function (_1037_i4) {
            return (_this).f21;
          };
          let _nw173 = Array((new BigNumber(22)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw173.length); _i0_30++) {
            _nw173[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1036_v18 = _nw173;
          let _index170 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1036_v18).length));
          (_1036_v18)[_index170] = false;
          let _index171 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1036_v18).length));
          let _rhs183 = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_1038_i5) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }), _1035_cf21);
          let _rhs184 = !((_this).f24);
          let _lhs160 = _1036_v18;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1036_v18).length));
          let _lhs162 = globalState;
          _lhs160[_lhs161] = _rhs183;
          _lhs162.f14 = _rhs184;
          _1017_v5 = _1017_v5;
          let _1039_v19;
          _1039_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_1018_v6).f21) === ((_this).f24),p0);
          _1039_v19 = (_1039_v19).update((_1018_v6).f21, p0);
          let _1040_v20;
          _1040_v20 = _dafny.Seq.of(p0, new BigNumber(478), new BigNumber(-570));
          let _index172 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1036_v18).length));
          let _rhs185 = (_1018_v6).f21;
          let _rhs186 = new BigNumber((_1040_v20).length);
          let _rhs187 = (p0).plus(new BigNumber((_1040_v20).length));
          let _rhs188 = (_this).f21;
          let _rhs189 = (_1018_v6).f21;
          let _lhs163 = globalState;
          let _lhs164 = globalState;
          let _lhs165 = globalState;
          let _lhs166 = _1036_v18;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1036_v18).length));
          let _lhs168 = globalState;
          _lhs163.f18 = _rhs185;
          _lhs164.f5 = _rhs186;
          _lhs165.f5 = _rhs187;
          _lhs166[_lhs167] = _rhs188;
          _lhs168.f14 = _rhs189;
        }
        let _1041_v21;
        let _init31 = ((_1042_v7) => function (_1043_i6) {
          return _1042_v7;
        })(_1019_v7);
        let _nw174 = Array((new BigNumber(20)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw174.length); _i0_31++) {
          _nw174[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1041_v21 = _nw174;
        let _index173 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1041_v21).length));
        (_1041_v21)[_index173] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), function (_1044_i7) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        });
        let _1045_v22;
        _1045_v22 = _dafny.Seq.of(p0, p0, (_1018_v6).f20);
        let _1046_v23;
        _1046_v23 = _dafny.Seq.of(_1045_v22, _1045_v22, _1045_v22);
        let _1047_v24;
        _1047_v24 = _module.D9.create_DC22(_1016_v4);
        let _1048_v25;
        _1048_v25 = _module.D9.create_DC24(_1047_v24);
        let _1049_v26;
        _1049_v26 = _module.D9.create_DC24(_1048_v25);
        let _1050_v27;
        _1050_v27 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1051_v28;
        _1051_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f20,(_1018_v6).f20);
        let _1052_v29;
        _1052_v29 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21);
        let _1053_v30;
        _1053_v30 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1018_v6).f21), _1052_v29);
        let _1054_v31;
        let _nw175 = new _module.C3();
        _nw175.__ctor(new BigNumber(382), (_this).fm0(_1051_v28, _1019_v7, _1053_v30, globalState), (_1018_v6).f20, (_this).fm13((_1018_v6).f20, globalState));
        _1054_v31 = _nw175;
        let _1055_v32;
        _1055_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v31,_1046_v23);
        let _index174 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1041_v21).length));
        let _rhs190 = _dafny.Seq.update(_1019_v7, _module.__default.safeIndex((_1018_v6).f20, new BigNumber((_1019_v7).length)), _1050_v27);
        let _rhs191 = _module.__default.safeModuloInt(new BigNumber(657), p0);
        let _rhs192 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1046_v23, (((_1055_v32).contains(_1054_v31)) ? ((_1055_v32).get(_1054_v31)) : (_1046_v23))), _1046_v23);
        let _rhs193 = _1049_v26;
        let _rhs194 = _1041_v21;
        let _lhs169 = _1041_v21;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_1041_v21).length));
        let _lhs171 = globalState;
        _lhs169[_lhs170] = _rhs190;
        _lhs171.f5 = _rhs191;
        _1046_v23 = _rhs192;
        _1049_v26 = _rhs193;
        _1041_v21 = _rhs194;
        let _1056_v33;
        let _nw176 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1056_v33 = _nw176;
        let _index175 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1056_v33).length));
        (_1056_v33)[_index175] = (function () {
          let _coll31 = new _dafny.Set();
          for (const _compr_31 of _dafny.IntegerRange(new BigNumber(841), new BigNumber(77))) {
            let _1057_v34 = _compr_31;
            if (((new BigNumber(841)).isLessThanOrEqualTo(_1057_v34)) && ((_1057_v34).isLessThan(new BigNumber(77)))) {
              _coll31.add((_1057_v34).multipliedBy((_1018_v6).f20));
            }
          }
          return _coll31;
        }()).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(987)));
        let _index176 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1056_v33).length));
        (_1056_v33)[_index176] = (new BigNumber((_1019_v7).length)).isLessThan(new BigNumber((_1019_v7).length));
        let _1058_v35;
        _1058_v35 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm34(_1050_v27, (_this).f21, (_1018_v6).f20, p0, globalState)).dtor_cf12,new BigNumber(-844));
        let _1059_v36;
        _1059_v36 = _dafny.MultiSet.fromElements(p0);
        let _1060_v37;
        _1060_v37 = _dafny.Seq.of(_1058_v35, (_dafny.Map.Empty.slice().updateUnsafe((_1056_v33)[_module.__default.safeIndex(new BigNumber(620), new BigNumber((_1056_v33).length))],new BigNumber((_1059_v36).cardinality()))).Merge(_1058_v35), (((_this).f21) ? (_dafny.Map.Empty.slice().updateUnsafe((_1056_v33)[_module.__default.safeIndex(new BigNumber(620), new BigNumber((_1056_v33).length))],p0)) : (_1058_v35)), (_1058_v35).Merge(_module.__default.fm27((_1018_v6).f20, globalState)), _1058_v35);
        _1060_v37 = _1060_v37;
        let _1061_v38;
        let _nw177 = new _module.C1();
        _nw177.__ctor(p0, _1017_v5, new BigNumber(342), (_1054_v31).f26);
        _1061_v38 = _nw177;
        let _1062_v39;
        _1062_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f20,_1061_v38);
        let _1063_v40;
        let _nw178 = new _module.C1();
        _nw178.__ctor((_dafny.ZERO).minus(p0), _1017_v5, new BigNumber((_1062_v39).length), !(_1061_v38.f27).isEqualTo((_1018_v6).f20));
        _1063_v40 = _nw178;
      } else {
        (globalState).f18 = (_1018_v6).f21;
        let _rhs195 = p0;
        let _rhs196 = (_1018_v6).f20;
        let _rhs197 = (p1) && (((_1018_v6).f20).isLessThanOrEqualTo(p0));
        let _lhs172 = globalState;
        let _lhs173 = globalState;
        let _lhs174 = globalState;
        _lhs172.f17 = _rhs195;
        _lhs173.f17 = _rhs196;
        _lhs174.f14 = _rhs197;
        let _1064_v41;
        _1064_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_1018_v6).f21));
        let _1065_v42;
        _1065_v42 = _dafny.Map.Empty.slice().updateUnsafe(!((((_1064_v41).contains((_this).f20)) ? ((_1064_v41).get((_this).f20)) : ((_this).f21))),new BigNumber(-741));
        _1065_v42 = (_1065_v42).update((((_1018_v6).f21) ? (p1) : ((_1018_v6).f21)), (_this).f20);
        _1017_v5 = _1017_v5;
        (globalState).f17 = _module.__default.safeModuloInt((_this).f20, (_1018_v6).f20);
      }
      let _1066_v43;
      _1066_v43 = _dafny.Set.fromElements((_this).f21, p1);
      let _1067_v44;
      _1067_v44 = _dafny.Set.fromElements(new BigNumber((_1066_v43).length));
      let _1068_v45;
      _1068_v45 = _dafny.MultiSet.fromElements(new BigNumber((_1067_v44).length));
      if ((_1068_v45).IsProperSubsetOf((_1068_v45).Difference(_1068_v45))) {
        let _1069_v46;
        _1069_v46 = _module.D13.create_DC36((_this).f20, !(!(((_1018_v6).f21) && ((_1018_v6).f21))), (_1016_v4)[_module.__default.safeIndex((_1018_v6).f20, new BigNumber((_1016_v4).length))]);
        let _1070_v47;
        _1070_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
        let _1071_v48;
        _1071_v48 = _dafny.Seq.UnicodeFromString("hadgdwudf");
        let _rhs198 = (_module.__default.safeModuloInt((_1018_v6).f20, (_this).f20)).plus((((_1070_v47).contains(p0)) ? ((_1070_v47).get(p0)) : ((_this).f20)));
        let _rhs199 = new BigNumber((_1071_v48).length);
        let _rhs200 = _1069_v46;
        let _lhs175 = globalState;
        let _lhs176 = globalState;
        _lhs175.f5 = _rhs198;
        _lhs176.f5 = _rhs199;
        _1069_v46 = _rhs200;
        let _1072_v49;
        _1072_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1071_v48).length),_1018_v6);
        if ((new BigNumber((_1072_v49).length)).isLessThan((_1018_v6).f20)) {
          _1017_v5 = _1017_v5;
          (globalState).f5 = (_1018_v6).f20;
          let _1073_v50;
          _1073_v50 = _dafny.MultiSet.fromElements(p1);
          let _1074_v51;
          _1074_v51 = _dafny.Seq.of((((_1073_v50).contains((_this).f21)) ? ((_1073_v50).get((_this).f21)) : ((_1018_v6).fm1((_1018_v6).f20, new BigNumber((_dafny.Set.fromElements((_this).f20, p0, new BigNumber(809))).length), (_1018_v6).f21, globalState))), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_1075_i8) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          })).length));
          let _1076_v52;
          _1076_v52 = new _dafny.CodePoint('v'.codePointAt(0));
          (globalState).f18 = _dafny.Seq.IsPrefixOf(_1071_v48, (((_1018_v6).f21) ? (_1071_v48) : (_dafny.Seq.update(_1071_v48, _module.__default.safeIndex(new BigNumber((_1074_v51).length), new BigNumber((_1071_v48).length)), _1076_v52))));
          let _1077_v53;
          let _nw179 = Array((_dafny.ONE).toNumber()).fill(false);
          _1077_v53 = _nw179;
          let _1078_v56;
          _1078_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37(new BigNumber((_1071_v48).length), globalState),new BigNumber((function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), ((_1079_v52) => function (_1080_i9) {
              return _1079_v52;
            })(_1076_v52))).Elements) {
              let _1081_v55 = _compr_32;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), ((_1082_v52) => function (_1080_i9) {
                return _1082_v52;
              })(_1076_v52)), _1081_v55)) {
                _coll32.push([_1081_v55,(_1018_v6).f20]);
              }
            }
            return _coll32;
          }()).length));
          let _index177 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1077_v53).length));
          (_1077_v53)[_index177] = (function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (_1078_v56).Keys.Elements) {
              let _1083_v54 = _compr_33;
              if ((_1078_v56).contains(_1083_v54)) {
                _coll33.push([_1083_v54,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_1084_v52) => function (_1085_i10) {
                  return _1084_v52;
                })(_1076_v52))).length)]);
              }
            }
            return _coll33;
          }()).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f20),(_this).f20));
          let _index178 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1077_v53).length));
          (_1077_v53)[_index178] = (_1018_v6).f21;
          let _index179 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_1077_v53).length));
          (_1077_v53)[_index179] = !((_1077_v53)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_1077_v53).length))]) || ((p0).isEqualTo(new BigNumber(208)));
        } else {
          let _1086_v57;
          let _nw180 = new _module.C3();
          _nw180.__ctor(((_1018_v6).f20).minus((_1018_v6).f20), ((_1018_v6).f20).isLessThanOrEqualTo(p0), new BigNumber(582), ((_1018_v6).f21) === ((_1018_v6).f21));
          _1086_v57 = _nw180;
          let _1087_v58;
          _1087_v58 = _dafny.Seq.of(new BigNumber(251));
          let _1088_v59;
          _1088_v59 = _dafny.MultiSet.fromElements(_1067_v44);
          let _1089_v60;
          _1089_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_1086_v57).f25);
          let _1090_v61;
          _1090_v61 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1091_v62;
          _1091_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v60,_1090_v61);
          let _1092_v63;
          let _nw181 = Array((new BigNumber(11)).toNumber());
          _nw181[(_dafny.ZERO).toNumber()] = _1087_v58;
          _nw181[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_1087_v58, _module.__default.safeIndex((((_1088_v59).contains(_1067_v44)) ? ((_1088_v59).get(_1067_v44)) : (new BigNumber((_1071_v48).length))), new BigNumber((_1087_v58).length)), new BigNumber((_1071_v48).length)), _module.__default.safeIndex(new BigNumber((_1091_v62).length), new BigNumber((_dafny.Seq.update(_1087_v58, _module.__default.safeIndex((((_1088_v59).contains(_1067_v44)) ? ((_1088_v59).get(_1067_v44)) : (new BigNumber((_1071_v48).length))), new BigNumber((_1087_v58).length)), new BigNumber((_1071_v48).length))).length)), (_1018_v6).f20);
          _nw181[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_1093_i11) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("gxcytgdfb")).length);
          }), _1087_v58);
          _nw181[(new BigNumber(3)).toNumber()] = _dafny.Seq.of((_1018_v6).f20);
          _nw181[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), ((_1094_p0) => function (_1095_i12) {
            return _1094_p0;
          })(p0));
          _nw181[(new BigNumber(5)).toNumber()] = _module.__default.fm28((_1018_v6).f20, globalState);
          _nw181[(new BigNumber(6)).toNumber()] = _1087_v58;
          _nw181[(new BigNumber(7)).toNumber()] = _1087_v58;
          _nw181[(new BigNumber(8)).toNumber()] = _1087_v58;
          _nw181[(new BigNumber(9)).toNumber()] = _1087_v58;
          _nw181[(new BigNumber(10)).toNumber()] = _1087_v58;
          _1092_v63 = _nw181;
          let _index180 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1092_v63).length));
          (_1092_v63)[_index180] = _1087_v58;
          let _1096_v64;
          _1096_v64 = _dafny.MultiSet.fromElements(true, (_this).f24);
          let _index181 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1092_v63).length));
          let _rhs201 = (_dafny.ZERO).minus(((new BigNumber((_1096_v64).cardinality())).plus(new BigNumber(908))).plus((_1018_v6).f20));
          let _rhs202 = _dafny.Seq.Concat(_dafny.Seq.update(_1087_v58, _module.__default.safeIndex(p0, new BigNumber((_1087_v58).length)), new BigNumber((_1016_v4).length)), _module.__default.fm28((_1018_v6).f20, globalState));
          let _lhs177 = globalState;
          let _lhs178 = _1092_v63;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1092_v63).length));
          _lhs177.f17 = _rhs201;
          _lhs178[_lhs179] = _rhs202;
          (globalState).f18 = !((_1018_v6).f21) || (!(!((_1018_v6).f21)));
          _1017_v5 = _1017_v5;
          let _1097_v65;
          let _nw182 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1097_v65 = _nw182;
          let _1098_v66;
          _1098_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1097_v65,(_1018_v6).f21);
          let _1099_v67;
          _1099_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f20,p1),(_1086_v57).f25);
          let _1100_v68;
          _1100_v68 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f20,(_this).f21);
          (globalState).f17 = new BigNumber((((((_this).f21) ? (_module.__default.fm38((_dafny.ZERO).minus((_1086_v57).f25), (((_1098_v66).contains(_1097_v65)) ? ((_1098_v66).get(_1097_v65)) : (p1)), new BigNumber((_dafny.Seq.UnicodeFromString("qrgiiat")).length), (_1018_v6).f21, globalState)) : (_1099_v67))).update(_1100_v68, (p0).minus((_1086_v57).f25))).length);
        }
        (globalState).f5 = (_1018_v6).f20;
        let _1101_v69;
        _1101_v69 = _dafny.Seq.of(_1071_v48);
        let _1102_v70;
        let _nw183 = Array((new BigNumber(13)).toNumber());
        _nw183[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-767)), function (_1103_i13) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        });
        _nw183[(_dafny.ONE).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(2)).toNumber()] = _module.__default.fm26((_1018_v6).f21, globalState);
        _nw183[(new BigNumber(3)).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm26((_this).f24, globalState), _1071_v48);
        _nw183[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1071_v48, _dafny.Seq.UnicodeFromString("gntqqvui"));
        _nw183[(new BigNumber(6)).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat((_1101_v69)[_module.__default.safeIndex((_this).f20, new BigNumber((_1101_v69).length))], _1071_v48);
        _nw183[(new BigNumber(8)).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(9)).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("pxfbmf");
        _nw183[(new BigNumber(11)).toNumber()] = _1071_v48;
        _nw183[(new BigNumber(12)).toNumber()] = _1071_v48;
        _1102_v70 = _nw183;
        _1102_v70 = _1102_v70;
        if (!(true)) {
          (globalState).f4 = (_1018_v6).f21;
          (globalState).f15 = _dafny.Seq.UnicodeFromString("yhpdkduv");
          let _1104_v71;
          _1104_v71 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1105_v72;
          let _init32 = function (_1106_i14) {
            return (_this).f21;
          };
          let _nw184 = Array((new BigNumber(16)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw184.length); _i0_32++) {
            _nw184[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1105_v72 = _nw184;
          let _index182 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1105_v72).length));
          (_1105_v72)[_index182] = (_1018_v6).f21;
          let _1107_v73;
          _1107_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_this).f21);
          let _1108_v74;
          _1108_v74 = _dafny.Map.Empty.slice().updateUnsafe(!(true),((_1107_v73).update((_this).f24, (_1018_v6).f21)).Merge(_1107_v73));
          let _1109_v75;
          let _nw185 = Array((new BigNumber(23)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = _1017_v5;
          _nw185[(_dafny.ONE).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(2)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(3)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(4)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(5)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(6)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(7)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(8)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(9)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(10)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(11)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(12)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(13)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(14)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(15)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(16)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(17)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(18)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(19)).toNumber()] = _1017_v5;
          _nw185[(new BigNumber(20)).toNumber()] = ((true) ? (_1017_v5) : (_1017_v5));
          _nw185[(new BigNumber(21)).toNumber()] = (((_this).f24) ? (_1017_v5) : (_1017_v5));
          _nw185[(new BigNumber(22)).toNumber()] = (((_1018_v6).f21) ? (_1017_v5) : (_1017_v5));
          _1109_v75 = _nw185;
          let _index183 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length));
          (_1109_v75)[_index183] = _1017_v5;
          let _index184 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1105_v72).length));
          let _index185 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length));
          let _rhs203 = _1104_v71;
          let _rhs204 = (_1018_v6).f21;
          let _rhs205 = (_1018_v6).f21;
          let _rhs206 = _1108_v74;
          let _rhs207 = _1017_v5;
          let _lhs180 = globalState;
          let _lhs181 = _1105_v72;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1105_v72).length));
          let _lhs183 = _1109_v75;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length));
          _1104_v71 = _rhs203;
          _lhs180.f14 = _rhs204;
          _lhs181[_lhs182] = _rhs205;
          _1108_v74 = _rhs206;
          _lhs183[_lhs184] = _rhs207;
          let _1110_v76;
          _1110_v76 = _dafny.MultiSet.fromElements(p1);
          let _1111_v77;
          _1111_v77 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_1018_v6).f20);
          let _1112_v78;
          _1112_v78 = _dafny.Seq.of((_this).f20, (_1018_v6).f20, (_this).f20, new BigNumber((_1111_v77).length));
          let _rhs208 = (((_1110_v76).update((_1018_v6).f21, _module.__default.abs(new BigNumber((_dafny.Seq.of((((_1068_v45).contains((_1018_v6).f20)) ? ((_1068_v45).get((_1018_v6).f20)) : ((_this).f20)), (_this).f20, new BigNumber((_1112_v78).length), (_1018_v6).f20)).length)))).Union(_dafny.MultiSet.FromArray(_1016_v4))).IsProperSubsetOf(_1110_v76);
          let _rhs209 = _module.__default.safeDivisionInt((_1018_v6).f20, new BigNumber((_dafny.Seq.Concat((_1101_v69)[_module.__default.safeIndex((_1018_v6).f20, new BigNumber((_1101_v69).length))], _dafny.Seq.UnicodeFromString("jwncqvclf"))).length));
          let _rhs210 = (_1018_v6).f20;
          let _lhs185 = globalState;
          let _lhs186 = globalState;
          let _lhs187 = globalState;
          _lhs185.f4 = _rhs208;
          _lhs186.f5 = _rhs209;
          _lhs187.f17 = _rhs210;
          let _1113_v79;
          _1113_v79 = _dafny.Seq.of(_1068_v45, (_1068_v45).Union(_1068_v45));
          let _arr2 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _index186 = _module.__default.safeIndex(new BigNumber(754), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          _arr2[_index186] = _module.__default.safeModuloInt((_this).f20, (_1018_v6).f20);
          let _index187 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1102_v70).length));
          (_1102_v70)[_index187] = _dafny.Seq.UnicodeFromString("u");
          let _arr3 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _index188 = _module.__default.safeIndex(new BigNumber(400), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          _arr3[_index188] = ((_dafny.ZERO).minus((_1018_v6).f20)).multipliedBy(p0);
          let _arr4 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _index189 = _module.__default.safeIndex(new BigNumber(754), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          let _index190 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1102_v70).length));
          let _arr5 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _index191 = _module.__default.safeIndex(new BigNumber(400), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          let _rhs211 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(185)), ((_1114_v45) => function (_1115_i15) {
            return _1114_v45;
          })(_1068_v45));
          let _rhs212 = (_dafny.ZERO).minus((_this).fm14((_1110_v76).IsSubsetOf(_dafny.MultiSet.fromElements(!((_1018_v6).f21), (_1018_v6).f21, !((_1018_v6).f21), (_1018_v6).f21)), (_1018_v6).f20, (_1018_v6).f20, (_1018_v6).f21, globalState));
          let _rhs213 = (_1018_v6).f20;
          let _rhs214 = _dafny.Seq.Concat(_1071_v48, _1071_v48);
          let _rhs215 = ((((_1110_v76).contains((_1018_v6).f21)) ? ((_1110_v76).get((_1018_v6).f21)) : ((_dafny.ZERO).minus((_1018_v6).f20)))).plus(_module.__default.safeDivisionInt((_1018_v6).f20, (_1018_v6).f20));
          let _lhs188 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _lhs189 = _module.__default.safeIndex(new BigNumber(754), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          let _lhs190 = globalState;
          let _lhs191 = _1102_v70;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1102_v70).length));
          let _lhs193 = (_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))];
          let _lhs194 = _module.__default.safeIndex(new BigNumber(400), new BigNumber(((_1109_v75)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_1109_v75).length))]).length));
          _1113_v79 = _rhs211;
          _lhs188[_lhs189] = _rhs212;
          _lhs190.f5 = _rhs213;
          _lhs191[_lhs192] = _rhs214;
          _lhs193[_lhs194] = _rhs215;
        } else {
          (globalState).f17 = (_1018_v6).fm1((_1018_v6).f20, p0, (_this).f21, globalState);
          let _1116_v80;
          _1116_v80 = _dafny.Seq.of((_dafny.ZERO).minus(p0), (_1018_v6).f20, (_this).f20, (_1018_v6).f20);
          _1116_v80 = _dafny.Seq.Concat(_1116_v80, _1116_v80);
          let _1117_v81;
          let _nw186 = new _module.C2();
          _nw186.__ctor(_module.__default.safeDivisionInt((_1018_v6).f20, (_1018_v6).f20), (_this).fm13((_1018_v6).f20, globalState));
          _1117_v81 = _nw186;
          let _1118_v82;
          let _init33 = ((_1119_v6) => function (_1120_i16) {
            return !((_1119_v6).f21);
          })(_1018_v6);
          let _nw187 = Array((new BigNumber(15)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw187.length); _i0_33++) {
            _nw187[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1118_v82 = _nw187;
          let _1121_v83;
          _1121_v83 = _dafny.MultiSet.fromElements(_1118_v82, _1118_v82);
          _1121_v83 = _dafny.MultiSet.fromElements(_1118_v82);
          let _index192 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1118_v82).length));
          (_1118_v82)[_index192] = (_this).f24;
          let _index193 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1118_v82).length));
          (_1118_v82)[_index193] = (_this).f24;
          let _1122_v84;
          _1122_v84 = _dafny.MultiSet.fromElements((_1071_v48)[_module.__default.safeIndex(p0, new BigNumber((_1071_v48).length))]);
          let _1123_v85;
          _1123_v85 = (_1018_v6).f21;
          let _1124_v86;
          _1124_v86 = _1116_v80;
          let _index194 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1118_v82).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1118_v82).length));
          let _rhs216 = !((new BigNumber((_1122_v84).cardinality())).isLessThanOrEqualTo((_1018_v6).f20));
          let _rhs217 = (((_1018_v6).f21) ? ((_this).f21) : ((_1123_v85)));
          let _rhs218 = (_1018_v6).f21;
          let _rhs219 = (_1018_v6).f21;
          let _rhs220 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1116_v80, (((_this).f24) ? (_1116_v80) : ((_1124_v86))))).length));
          let _lhs195 = _1118_v82;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1118_v82).length));
          let _lhs197 = globalState;
          let _lhs198 = _1118_v82;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1118_v82).length));
          let _lhs200 = globalState;
          r1 = _rhs216;
          _lhs195[_lhs196] = _rhs217;
          _lhs197.f4 = _rhs218;
          _lhs198[_lhs199] = _rhs219;
          _lhs200.f17 = _rhs220;
        }
      } else {
        (globalState).f17 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), function (_1125_i17) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("h"))).length), new BigNumber(409));
        let _1126_v88;
        _1126_v88 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_this).f21);
        let _1127_v89;
        _1127_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1126_v88,(_1018_v6).f21);
        let _1128_v90;
        _1128_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1066_v43).length),(_1018_v6).f21);
        let _1129_v91;
        _1129_v91 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1128_v90);
        let _1130_v92;
        _1130_v92 = _dafny.Map.Empty.slice().updateUnsafe((((_1129_v91).contains(p0)) ? ((_1129_v91).get(p0)) : (_1128_v90)),(_this).f24);
        (globalState).f14 = (_1130_v92).contains(function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of (_dafny.MultiSet.fromElements(new BigNumber((_1127_v89).length))).Elements) {
            let _1131_v87 = _compr_34;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_1127_v89).length))).contains(_1131_v87)) {
              _coll34.push([(_1131_v87).multipliedBy((_1018_v6).f20),(_998_v1).dtor_cf16]);
            }
          }
          return _coll34;
        }());
        if (p1) {
          let _1132_v93;
          let _nw188 = Array((new BigNumber(5)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = _1068_v45;
          _nw188[(_dafny.ONE).toNumber()] = _1068_v45;
          _nw188[(new BigNumber(2)).toNumber()] = _1068_v45;
          _nw188[(new BigNumber(3)).toNumber()] = (_1068_v45).Difference(_1068_v45);
          _nw188[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(993));
          _1132_v93 = _nw188;
          let _1133_v94;
          _1133_v94 = _dafny.Seq.of((_1018_v6).f20);
          let _index196 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1132_v93).length));
          (_1132_v93)[_index196] = (_dafny.MultiSet.FromArray(_1133_v94)).Intersect(_1068_v45);
          let _index197 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1132_v93).length));
          (_1132_v93)[_index197] = _1068_v45;
          r1 = ((_1066_v43).Difference(_1066_v43)).equals((_1066_v43).Difference(_1066_v43));
          let _1134_v95;
          _1134_v95 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1135_v96;
          let _nw189 = new _module.C1();
          _nw189.__ctor((new BigNumber(955)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), function (_1136_i18) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(((_1132_v93)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_1132_v93).length))]).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), function (_1137_i18) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length)), _1134_v95)).length))), (((_this).f21) ? (_1017_v5) : (_1017_v5)), (_1018_v6).f20, (new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), ((_1138_v95) => function (_1139_i19) {
            return _1138_v95;
          })(_1134_v95)), _module.__default.safeIndex((_1018_v6).f20, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), ((_1140_v95) => function (_1141_i19) {
            return _1140_v95;
          })(_1134_v95))).length)), _1134_v95)).length)).isLessThan((_1018_v6).f20));
          _1135_v96 = _nw189;
          let _1142_v97;
          _1142_v97 = _dafny.Seq.UnicodeFromString("d");
          let _1143_v98;
          _1143_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(453),(_dafny.ZERO).minus(new BigNumber((_1142_v97).length)));
          (globalState).f18 = (_1018_v6).fm0(_1143_v98, _1142_v97, _dafny.Seq.update(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1018_v6).f21)), _module.__default.safeIndex(new BigNumber((_1133_v94).length), new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1018_v6).f21))).length)), _1126_v88), globalState);
          (globalState).f18 = ((_1018_v6).f20).isLessThan(_1135_v96.f27);
        } else {
          let _1144_v99;
          let _nw190 = new _module.C1();
          _nw190.__ctor(p0, _1017_v5, p0, !(!(!(true) || ((_this).f21))));
          _1144_v99 = _nw190;
          let _1145_v100;
          let _nw191 = Array((new BigNumber(26)).toNumber()).fill(false);
          _1145_v100 = _nw191;
          _1145_v100 = _1145_v100;
          let _index198 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1145_v100).length));
          (_1145_v100)[_index198] = (_this).f24;
          let _index199 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1145_v100).length));
          (_1145_v100)[_index199] = p1;
          let _index200 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1017_v5).length));
          (_1017_v5)[_index200] = new BigNumber(753);
          let _1146_v101;
          _1146_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1016_v4);
          let _index201 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1017_v5).length));
          (_1017_v5)[_index201] = new BigNumber(((_1146_v101).Merge(_1146_v101)).length);
          _1067_v44 = (_1067_v44).Union((_1067_v44).Difference(_1067_v44));
        }
        (globalState).f5 = _module.__default.safeDivisionInt((_1018_v6).f20, p0);
        (globalState).f14 = (_this).f21;
      }
      let _1147_v102;
      _1147_v102 = _dafny.MultiSet.fromElements(p1);
      if ((!((_dafny.MultiSet.fromElements((_1018_v6).f21)).IsDisjointFrom(_1147_v102))) && ((_this).f24)) {
        let _1148_v103;
        _1148_v103 = _dafny.Seq.of(p0, p0, (_this).f20);
        let _1149_v105;
        _1149_v105 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f20);
        let _1150_v106;
        _1150_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_dafny.Seq.UnicodeFromString("keo"));
        let _1151_v107;
        _1151_v107 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        let _1152_v108;
        _1152_v108 = _dafny.Seq.UnicodeFromString("rupn");
        let _1153_v109;
        _1153_v109 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_1018_v6).f21);
        let _rhs221 = (_this).fm0((function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(305))) {
            let _1154_v104 = _compr_35;
            if (((new BigNumber(297)).isLessThanOrEqualTo(_1154_v104)) && ((_1154_v104).isLessThan(new BigNumber(305)))) {
              _coll35.push([(_1154_v104).multipliedBy((_1018_v6).f20),(_this).f20]);
            }
          }
          return _coll35;
        }()).Merge(_1149_v105), (((_1150_v106).contains((((_1151_v107).contains((_1018_v6).f20)) ? ((_1151_v107).get((_1018_v6).f20)) : (p1)))) ? ((_1150_v106).get((((_1151_v107).contains((_1018_v6).f20)) ? ((_1151_v107).get((_1018_v6).f20)) : (p1)))) : (_1152_v108)), _dafny.Seq.of(_1153_v109), globalState);
        let _rhs222 = _dafny.Seq.Concat(_1148_v103, _1148_v103);
        let _lhs201 = globalState;
        _lhs201.f4 = _rhs221;
        _1148_v103 = _rhs222;
        (globalState).f5 = _module.__default.safeModuloInt(p0, (_1018_v6).f20);
        (globalState).f5 = (_dafny.ZERO).minus((_1018_v6).f20);
        if (false) {
          let _1155_v110;
          _1155_v110 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,_module.__default.fm17((_this).f21, globalState));
          let _index202 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1017_v5).length));
          (_1017_v5)[_index202] = new BigNumber(((_1155_v110).Merge(_1155_v110)).length);
          let _index203 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1017_v5).length));
          let _rhs223 = _1147_v102;
          let _rhs224 = (_dafny.ZERO).minus((_this).f20);
          let _rhs225 = new BigNumber((_1147_v102).cardinality());
          let _rhs226 = _module.__default.safeModuloInt(((_1018_v6).f20).multipliedBy((_this).f20), (_this).f20);
          let _lhs202 = globalState;
          let _lhs203 = _1017_v5;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1017_v5).length));
          let _lhs205 = globalState;
          _1147_v102 = _rhs223;
          _lhs202.f17 = _rhs224;
          _lhs203[_lhs204] = _rhs225;
          _lhs205.f5 = _rhs226;
          let _1156_v111;
          _1156_v111 = _module.D10.create_DC25(_1068_v45);
          let _1157_v112;
          _1157_v112 = _dafny.Seq.of((_1156_v111).dtor_cf34);
          let _1158_v113;
          let _nw192 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1158_v113 = _nw192;
          let _1159_v114;
          _1159_v114 = _dafny.Seq.of(_1158_v113);
          let _1160_v115;
          _1160_v115 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(776)), ((_1161_v45) => function (_1162_i20) {
            return _1161_v45;
          })(_1068_v45)), _1157_v112),_dafny.Seq.Concat(_1159_v114, _1159_v114));
          let _1163_v116;
          _1163_v116 = _module.D3.create_DC5(_1159_v114);
          _1160_v115 = (_1160_v115).update(_dafny.Seq.of(_1068_v45), (_1163_v116).dtor_cf9);
          let _1164_v117;
          _1164_v117 = _dafny.Seq.of(_1016_v4, _dafny.Seq.update(_1016_v4, _module.__default.safeIndex(new BigNumber((_1153_v109).length), new BigNumber((_1016_v4).length)), (_1018_v6).f21), _1016_v4, _1016_v4, _1016_v4);
          let _1165_v118;
          _1165_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1164_v117,_dafny.Seq.UnicodeFromString("cbrjrlv"));
          let _1166_v119;
          _1166_v119 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs227 = (_this).f21;
          let _rhs228 = (_1018_v6).f21;
          let _rhs229 = _dafny.Seq.update(_dafny.Seq.Concat(_1152_v108, (((_1165_v118).contains(_1164_v117)) ? ((_1165_v118).get(_1164_v117)) : (_1152_v108))), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Concat(_1152_v108, (((_1165_v118).contains(_1164_v117)) ? ((_1165_v118).get(_1164_v117)) : (_1152_v108)))).length)), _1166_v119);
          let _lhs206 = globalState;
          let _lhs207 = globalState;
          let _lhs208 = globalState;
          _lhs206.f14 = _rhs227;
          _lhs207.f14 = _rhs228;
          _lhs208.f15 = _rhs229;
          let _index204 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1017_v5).length));
          (_1017_v5)[_index204] = (_this).fm14(true, (_1018_v6).f20, (_1018_v6).f20, !_dafny.areEqual(_1152_v108, _module.__default.fm26((_this).f21, globalState)), globalState);
          _1153_v109 = (_1153_v109).update(true, (((_1151_v107).contains(new BigNumber((_1153_v109).length))) ? ((_1151_v107).get(new BigNumber((_1153_v109).length))) : ((_1018_v6).f21)));
        } else {
          let _1167_v120;
          _1167_v120 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(93));
          (globalState).f5 = (_dafny.ZERO).minus(((((_1167_v120).contains((_1018_v6).f21)) ? ((_1167_v120).get((_1018_v6).f21)) : ((_1018_v6).f20))).minus((_1018_v6).f20));
          let _1168_v121;
          let _nw193 = Array((new BigNumber(21)).toNumber()).fill(false);
          _1168_v121 = _nw193;
          let _index205 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1168_v121).length));
          (_1168_v121)[_index205] = (_this).f21;
          let _index206 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1168_v121).length));
          (_1168_v121)[_index206] = (_1016_v4)[_module.__default.safeIndex((_1018_v6).f20, new BigNumber((_1016_v4).length))];
          let _index207 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1168_v121).length));
          (_1168_v121)[_index207] = (_this).fm13(p0, globalState);
          (globalState).f5 = (new BigNumber(((_1149_v105).Merge(_1149_v105)).length)).plus((_1018_v6).f20);
          (globalState).f17 = (_this).f20;
        }
        let _1169_v122;
        let _nw194 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1169_v122 = _nw194;
        _1169_v122 = _1169_v122;
      } else {
        _1017_v5 = (((_this).f21) ? (_1017_v5) : (_1017_v5));
        let _1170_v123;
        _1170_v123 = _dafny.Seq.UnicodeFromString("t");
        let _1171_v124;
        _1171_v124 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1170_v123).length),new BigNumber((_1016_v4).length));
        let _1172_v125;
        _1172_v125 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f20,(((_1171_v124).contains((_this).f20)) ? ((_1171_v124).get((_this).f20)) : ((_this).f20)));
        (globalState).f4 = (_1018_v6).fm0(_1172_v125, _1170_v123, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-561)), ((_1173_v6) => function (_1174_i21) {
          return _dafny.Map.Empty.slice().updateUnsafe((_1173_v6).f21,(_1173_v6).f21);
        })(_1018_v6)), globalState);
        let _1175_v126;
        let _1176_v127;
        let _out16;
        let _out17;
        let _outcollector5 = (_this).m1(new BigNumber((_1170_v123).length), (_this).f24, _1018_v6, globalState);
        _out16 = _outcollector5[0];
        _out17 = _outcollector5[1];
        _1175_v126 = _out16;
        _1176_v127 = _out17;
        let _1177_v128;
        _1177_v128 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f21);
        let _1178_v129;
        _1178_v129 = _dafny.Seq.of((_1018_v6).f20);
        let _1179_v130;
        _1179_v130 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1180_v131;
        _1180_v131 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_1018_v6).f21);
        let _1181_v132;
        _1181_v132 = _dafny.Seq.of(_1180_v131);
        let _1182_v133;
        _1182_v133 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1018_v6).fm0(_1171_v124, _1170_v123, _1181_v132, globalState));
        _1177_v128 = (_1177_v128).update(new BigNumber((_dafny.Seq.Concat(_1178_v129, _dafny.Seq.update(_1178_v129, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1179_v130)).length), new BigNumber((_1178_v129).length)), new BigNumber((_1068_v45).cardinality())))).length), (((_1182_v133).contains((_1018_v6).f21)) ? ((_1182_v133).get((_1018_v6).f21)) : ((_this).f21)));
        let _1183_v134;
        _1183_v134 = _dafny.Seq.of(_1175_v126, _1175_v126, _1175_v126, _1017_v5, _1175_v126);
        r0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1183_v134, _1183_v134), (((_this).f21) ? (_dafny.Seq.of(_1175_v126, _1175_v126)) : (_1183_v134)));
      }
      let _1184_v135;
      _1184_v135 = _dafny.Seq.of(_1017_v5);
      r0 = _dafny.Seq.Concat(_1184_v135, _1184_v135);
      let _1185_v138;
      _1185_v138 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f21);
      let _1186_v139;
      _1186_v139 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _1187_v140;
      _1187_v140 = _dafny.Seq.of(_1185_v138, _module.__default.fm15(new BigNumber(522), (_1018_v6).f20, (((_1186_v139).contains((_this).f24)) ? ((_1186_v139).get((_this).f24)) : (p0)), (_1018_v6).f20, globalState), _1185_v138, _dafny.Map.Empty.slice().updateUnsafe((_1018_v6).f21,(_this).f21), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24));
      r1 = (_this).fm0(function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of _dafny.IntegerRange(new BigNumber(702), new BigNumber(756))) {
          let _1188_v136 = _compr_36;
          if (((new BigNumber(702)).isLessThanOrEqualTo(_1188_v136)) && ((_1188_v136).isLessThan(new BigNumber(756)))) {
            _coll36.push([(_1188_v136).multipliedBy(new BigNumber((function () {
              let _coll37 = new _dafny.Map();
              for (const _compr_37 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),p1)).length)))).Elements) {
                let _1189_v137 = _compr_37;
                if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),p1)).length))), _1189_v137)) {
                  _coll37.push([(_1189_v137).minus((_1018_v6).f20),new BigNumber((_1067_v44).length)]);
                }
              }
              return _coll37;
            }()).length)),(_dafny.ZERO).minus((_1018_v6).f20)]);
          }
        }
        return _coll36;
      }(), _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), function (_1190_i22) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }), _1187_v140, globalState);
      return [r0, r1];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      (globalState).f17 = p0;
      if ((_this).f24) {
        let _1191_v0;
        _1191_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(p2).f20);
        let _1192_v1;
        _1192_v1 = _dafny.Map.Empty.slice().updateUnsafe((p2).f20,_dafny.Seq.UnicodeFromString("tyiwlx"));
        let _1193_v2;
        _1193_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21);
        let _1194_v3;
        _1194_v3 = _dafny.Seq.of(_1193_v2, _1193_v2, _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_module.D7.create_DC17((_this).f24)).dtor_cf25));
        let _1195_v4;
        let _nw195 = new _module.C3();
        _nw195.__ctor((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm0(_1191_v0, (((_1192_v1).contains((p2).f20)) ? ((_1192_v1).get((p2).f20)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-688)), function (_1196_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }))), _1194_v3, globalState),true)).length)), (_this).f21, p0, (_this).f21);
        _1195_v4 = _nw195;
        let _nw196 = new _module.C3();
        _nw196.__ctor(new BigNumber(365), (_this).f24, new BigNumber(599), (p2).f21);
        _1195_v4 = _nw196;
        if (!((_1195_v4).f26)) {
          let _1197_v5;
          let _nw197 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1197_v5 = _nw197;
          let _1198_v6;
          _1198_v6 = _dafny.Seq.of(_1197_v5);
          let _1199_v7;
          _1199_v7 = _module.D3.create_DC5(_1198_v6);
          let _pat_let_tv23 = _1198_v6;
          let _1200_v8;
          let _nw198 = Array((new BigNumber(15)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = _1199_v7;
          _nw198[(_dafny.ONE).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(2)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(3)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(4)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(5)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(6)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(7)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(8)).toNumber()] = function (_pat_let24_0) {
            return function (_1201_dt__update__tmp_h0) {
              return function (_pat_let25_0) {
                return function (_1202_dt__update_hcf9_h0) {
                  return _module.D3.create_DC5(_1202_dt__update_hcf9_h0);
                }(_pat_let25_0);
              }(_pat_let_tv23);
            }(_pat_let24_0);
          }(_1199_v7);
          _nw198[(new BigNumber(9)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(10)).toNumber()] = _module.D3.create_DC5(_1198_v6);
          _nw198[(new BigNumber(11)).toNumber()] = _module.D3.create_DC5(_1198_v6);
          _nw198[(new BigNumber(12)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(13)).toNumber()] = _1199_v7;
          _nw198[(new BigNumber(14)).toNumber()] = _1199_v7;
          _1200_v8 = _nw198;
          let _1203_v9;
          let _nw199 = Array((new BigNumber(23)).toNumber());
          _nw199[(_dafny.ZERO).toNumber()] = _1200_v8;
          _nw199[(_dafny.ONE).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(2)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(3)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(4)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(5)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(6)).toNumber()] = (_module.D14.create_DC38(_1200_v8)).dtor_cf51;
          _nw199[(new BigNumber(7)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(8)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(9)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(10)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(11)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(12)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(13)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(14)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(15)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(16)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(17)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(18)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(19)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(20)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(21)).toNumber()] = _1200_v8;
          _nw199[(new BigNumber(22)).toNumber()] = _1200_v8;
          _1203_v9 = _nw199;
          let _index208 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1203_v9).length));
          (_1203_v9)[_index208] = _1200_v8;
          let _1204_v10;
          _1204_v10 = _dafny.Seq.UnicodeFromString("dpcrmxv");
          let _1205_v11;
          _1205_v11 = _dafny.Set.fromElements((_this).f24, !(!((_1195_v4).f26)));
          let _index209 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1203_v9).length));
          let _rhs230 = _module.__default.safeModuloInt(((_1195_v4).f25).multipliedBy((_1195_v4).f25), (_1195_v4).fm1(new BigNumber((_1204_v10).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,p1)).length), false, globalState));
          let _rhs231 = ((_this).f20).isLessThanOrEqualTo((new BigNumber((_1205_v11).length)).multipliedBy((p2).f20));
          let _rhs232 = (new BigNumber((_1193_v2).length)).minus(new BigNumber((_dafny.Seq.Concat(_1204_v10, _1204_v10)).length));
          let _rhs233 = _1200_v8;
          let _lhs209 = globalState;
          let _lhs210 = globalState;
          let _lhs211 = globalState;
          let _lhs212 = _1203_v9;
          let _lhs213 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1203_v9).length));
          _lhs209.f17 = _rhs230;
          _lhs210.f4 = _rhs231;
          _lhs211.f5 = _rhs232;
          _lhs212[_lhs213] = _rhs233;
          let _1206_v12;
          _1206_v12 = _dafny.Seq.of((p2).f21, (p2).f21);
          let _1207_v13;
          _1207_v13 = _dafny.Seq.of((_1206_v12)[_module.__default.safeIndex((_1195_v4).f25, new BigNumber((_1206_v12).length))]);
          let _1208_v14;
          _1208_v14 = _dafny.Seq.of(_1206_v12, _dafny.Seq.of((p2).f21));
          let _1209_v15;
          let _nw200 = new _module.C2();
          _nw200.__ctor(new BigNumber((_1208_v14).length), !((_this).fm13((p2).f20, globalState)));
          _1209_v15 = _nw200;
          let _1210_v18;
          _1210_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm0(_1191_v0, _1204_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), ((_1211_v3, _1212_v4) => function (_1213_i1) {
            return (_1211_v3)[_module.__default.safeIndex((_1212_v4).f25, new BigNumber((_1211_v3).length))];
          })(_1194_v3, _1195_v4)), globalState));
          let _1214_v19;
          _1214_v19 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((p2).f20,(_this).f24), _1210_v18, _1210_v18);
          let _1215_v20;
          _1215_v20 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1216_v21;
          _1216_v21 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (function () {
              let _coll39 = new _dafny.Map();
              for (const _compr_39 of (_1214_v19).Elements) {
                let _1217_v17 = _compr_39;
                if ((_1214_v19).contains(_1217_v17)) {
                  _coll39.push([_1217_v17,(_this).f24]);
                }
              }
              return _coll39;
            }()).Keys.Elements) {
              let _1218_v16 = _compr_38;
              if ((function () {
                let _coll40 = new _dafny.Map();
                for (const _compr_40 of (_1214_v19).Elements) {
                  let _1217_v17 = _compr_40;
                  if ((_1214_v19).contains(_1217_v17)) {
                    _coll40.push([_1217_v17,(_this).f24]);
                  }
                }
                return _coll40;
              }()).contains(_1218_v16)) {
                _coll38.push([_1218_v16,p1]);
              }
            }
            return _coll38;
          }(),_1215_v20);
          let _1219_v22;
          _1219_v22 = (_this).f24;
          let _1220_v23;
          _1220_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1210_v18,_1219_v22);
          let _1221_v24;
          _1221_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1215_v20);
          _1216_v21 = (_1216_v21).update(_1220_v23, (((_1221_v24).contains((_1195_v4).f26)) ? ((_1221_v24).get((_1195_v4).f26)) : (_1215_v20)));
          let _1222_v26;
          _1222_v26 = _dafny.Seq.of((_1195_v4).f25);
          let _pat_let_tv24 = _1207_v13;
          let _pat_let_tv25 = _1207_v13;
          let _1223_v27;
          _1223_v27 = _module.D8.create_DC19(function () {
  let _coll41 = new _dafny.Map();
  for (const _compr_41 of ((function (_pat_let26_0) {
    return function (_1224_dt__update__tmp_h1) {
      return function (_pat_let27_0) {
        return function (_1225_dt__update_hcf0_h0) {
          return _1225_dt__update_hcf0_h0;
        }(_pat_let27_0);
      }(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_pat_let_tv24)).cardinality())));
    }(_pat_let26_0);
  }(_1222_v26))).Elements) {
    let _1226_v25 = _compr_41;
    if (_dafny.Seq.contains((function (_pat_let28_0) {
      return function (_1227_dt__update__tmp_h1) {
        return function (_pat_let29_0) {
          return function (_1228_dt__update_hcf0_h0) {
            return _1228_dt__update_hcf0_h0;
          }(_pat_let29_0);
        }(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_pat_let_tv25)).cardinality())));
      }(_pat_let28_0);
    }(_1222_v26)), _1226_v25)) {
      _coll41.push([(_1226_v25).multipliedBy(p0),(p2).f20]);
    }
  }
  return _coll41;
}());
          let _1229_v28;
          _1229_v28 = _dafny.Seq.of((_1223_v27).dtor_cf28, (_1191_v0).update((_1195_v4).f25, p0));
          (globalState).f18 = !(_module.__default.safeModuloInt((p2).f20, new BigNumber(((_1229_v28)[_module.__default.safeIndex((p2).f20, new BigNumber((_1229_v28).length))]).length))).isEqualTo((p2).f20);
          let _1230_v29;
          _1230_v29 = _module.D13.create_DC37(p1, (_this).f20);
          let _1231_v30;
          _1231_v30 = _module.D4.create_DC11((p2).f20, (_1230_v29).dtor_cf49, (_1195_v4).f26);
          let _1232_v31;
          _1232_v31 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,_1231_v30);
          (globalState).f4 = (_1206_v12)[_module.__default.safeIndex(new BigNumber(((_1232_v31).update((_1195_v4).f26, _1231_v30)).length), new BigNumber((_1206_v12).length))];
        } else {
          (globalState).f14 = (p2).f21;
          let _1233_v32;
          let _init34 = ((_1234_v4) => function (_1235_i2) {
            return (_1235_i2).multipliedBy((_1234_v4).f25);
          })(_1195_v4);
          let _nw201 = Array((new BigNumber(29)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw201.length); _i0_34++) {
            _nw201[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1233_v32 = _nw201;
          let _1236_v33;
          _1236_v33 = _dafny.Set.fromElements(_1233_v32);
          let _1237_v34;
          let _nw202 = new _module.C1();
          _nw202.__ctor(new BigNumber(675), _1233_v32, (p2).f20, (_1195_v4).f26);
          _1237_v34 = _nw202;
          let _1238_v35;
          let _nw203 = Array((new BigNumber(4)).toNumber()).fill(false);
          _1238_v35 = _nw203;
          let _1239_v36;
          _1239_v36 = _dafny.Seq.of(_1238_v35, _1238_v35);
          let _1240_v37;
          _1240_v37 = _module.D3.create_DC5(_1239_v36);
          let _1241_v38;
          _1241_v38 = _module.D3.create_DC8(_1240_v37);
          let _1242_v39;
          _1242_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v34,_1241_v38);
          let _1243_v40;
          _1243_v40 = _dafny.Set.fromElements(_1237_v34.f27, new BigNumber(-407), new BigNumber((_dafny.Seq.of(p1)).length));
          let _1244_v41;
          _1244_v41 = _dafny.Seq.of((_this).f24);
          let _1245_v42;
          _1245_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v39,(_1243_v40).equals(_dafny.Set.fromElements((_this).f20, new BigNumber((_1244_v41).length), (_1237_v34).fm1((p2).f20, new BigNumber(-317), p1, globalState))));
          let _1246_v43;
          _1246_v43 = _dafny.Seq.UnicodeFromString("klgw");
          let _rhs234 = _1236_v33;
          let _rhs235 = _1245_v42;
          let _rhs236 = (_1237_v34).fm0(_1191_v0, _1246_v43, _1194_v3, globalState);
          let _lhs214 = globalState;
          _1236_v33 = _rhs234;
          _1245_v42 = _rhs235;
          _lhs214.f4 = _rhs236;
          let _1247_v44;
          _1247_v44 = _dafny.Map.Empty.slice().updateUnsafe((p2).f20,_1191_v0);
          (globalState).f4 = (p2).fm0((((_1247_v44).contains((p2).f20)) ? ((_1247_v44).get((p2).f20)) : (_1191_v0)), _dafny.Seq.UnicodeFromString("vtqfdj"), _1194_v3, globalState);
          let _1248_v45;
          _1248_v45 = _dafny.MultiSet.fromElements(!((_this).f24));
          _1248_v45 = _dafny.MultiSet.FromArray(_1244_v41);
          (globalState).f4 = (((_1195_v4).f26) ? ((_module.__default.fm37((p2).f20, globalState)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f20))) : ((p2).f21));
        }
        let _1249_v46;
        _1249_v46 = _dafny.Seq.UnicodeFromString("tyhopkq");
        let _1250_v47;
        _1250_v47 = _module.D8.create_DC19(_1191_v0);
        let _1251_v48;
        _1251_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1249_v46,true);
        let _1252_v49;
        _1252_v49 = _dafny.Set.fromElements(!((p2).f21));
        (globalState).f4 = (new BigNumber((((_1250_v47).dtor_cf28).update(new BigNumber((_1251_v48).length), new BigNumber((_1252_v49).length))).length)).isLessThan(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lwcov"), _1249_v46)).length));
        (globalState).f15 = _dafny.Seq.Concat((((_this).fm13((_dafny.ZERO).minus((p2).f20), globalState)) ? (_1249_v46) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), function (_1253_i3) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }))), _1249_v46);
        let _1254_v50;
        let _nw204 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _1254_v50 = _nw204;
        let _1255_v51;
        let _nw205 = new _module.C1();
        _nw205.__ctor((_this).f20, (((_this).f24) ? (_1254_v50) : (_1254_v50)), (new BigNumber((_dafny.Seq.UnicodeFromString("ndkax")).length)).multipliedBy(new BigNumber(447)), (p2).f21);
        _1255_v51 = _nw205;
        _1255_v51 = _1255_v51;
      } else {
        (globalState).f17 = (_dafny.ZERO).minus((p2).f20);
        let _1256_v52;
        _1256_v52 = _dafny.Seq.UnicodeFromString("ucx");
        let _1257_v53;
        _1257_v53 = _dafny.Seq.of((_this).f20);
        let _1258_v54;
        _1258_v54 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),new BigNumber((_1257_v53).length));
        let _1259_v55;
        _1259_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
        let _1260_v56;
        _1260_v56 = _dafny.Seq.of(_1259_v55, _1259_v55, _1259_v55, _1259_v55);
        let _1261_v57;
        _1261_v57 = _dafny.Seq.of((_this).f24, !((_this).fm0(_1258_v54, _1256_v52, _1260_v56, globalState)));
        let _1262_v58;
        let _nw206 = Array((new BigNumber(8)).toNumber());
        _nw206[(_dafny.ZERO).toNumber()] = new BigNumber(-892);
        _nw206[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p1, (_this).f21, (_this).f24, !(p1), (p2).f21)).cardinality());
        _nw206[(new BigNumber(2)).toNumber()] = new BigNumber((_1256_v52).length);
        _nw206[(new BigNumber(3)).toNumber()] = (p2).f20;
        _nw206[(new BigNumber(4)).toNumber()] = new BigNumber(120);
        _nw206[(new BigNumber(5)).toNumber()] = new BigNumber((_1261_v57).length);
        _nw206[(new BigNumber(6)).toNumber()] = p0;
        _nw206[(new BigNumber(7)).toNumber()] = (_this).f20;
        _1262_v58 = _nw206;
        let _1263_v59;
        let _nw207 = new _module.C1();
        _nw207.__ctor((p2).f20, _1262_v58, (_dafny.ZERO).minus((p2).f20), (p2).f21);
        _1263_v59 = _nw207;
        _1263_v59 = _1263_v59;
        let _1264_v60;
        _1264_v60 = _dafny.Seq.of(_1261_v57);
        _1261_v57 = _dafny.Seq.Concat((_1264_v60)[_module.__default.safeIndex((_this).f20, new BigNumber((_1264_v60).length))], _dafny.Seq.of((p2).f21, (p2).f21));
        let _1265_v61;
        let _nw208 = Array((new BigNumber(21)).toNumber());
        _1265_v61 = _nw208;
        let _index210 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1265_v61).length));
        (_1265_v61)[_index210] = _1263_v59;
        let _1266_v62;
        _1266_v62 = _1263_v59;
        let _index211 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1265_v61).length));
        (_1265_v61)[_index211] = (_1266_v62);
        (globalState).f18 = (p2).f21;
      }
      let _1267_v63;
      _1267_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
      let _1268_i4;
      _1268_i4 = _dafny.ZERO;
      L4: {
        while (((((_1267_v63).contains(new BigNumber(-116))) ? ((_1267_v63).get(new BigNumber(-116))) : ((_this).f24))) && ((p2).f21)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1268_i4)) {
              break L4;
            }
            _1268_i4 = (_1268_i4).plus(_dafny.ONE);
            let _1269_v64;
            _1269_v64 = _dafny.Seq.UnicodeFromString("uqwrp");
            let _1270_v65;
            _1270_v65 = _dafny.Seq.of((p2).f20);
            let _1271_v66;
            _1271_v66 = new _dafny.CodePoint('x'.codePointAt(0));
            (globalState).f3 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.update(_1269_v64, _module.__default.safeIndex(new BigNumber((_1270_v65).length), new BigNumber((_1269_v64).length)), _1271_v66), _1269_v64), _dafny.Seq.Concat(_1269_v64, _dafny.Seq.update(_1269_v64, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1271_v66, _1271_v66, _1271_v66, _1271_v66, new _dafny.CodePoint('f'.codePointAt(0)))).length), new BigNumber((_1269_v64).length)), _1271_v66)));
            let _1272_v67;
            let _init35 = function (_1273_i5) {
              return (_1273_i5).minus((_this).f20);
            };
            let _nw209 = Array((new BigNumber(17)).toNumber());
            for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw209.length); _i0_35++) {
              _nw209[_i0_35] = _init35(new BigNumber(_i0_35));
            }
            _1272_v67 = _nw209;
            let _1274_v68;
            let _nw210 = new _module.C1();
            _nw210.__ctor((_this).f20, _1272_v67, new BigNumber(284), !((_this).f21));
            _1274_v68 = _nw210;
            if ((p2).f21) {
              (globalState).f17 = (p2).f20;
              let _1275_v70;
              _1275_v70 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,(p2).f21);
              let _1276_v71;
              _1276_v71 = _module.D2.create_DC3((p2).f20, !((_this).f24), p1, (_this).f20, (_1274_v68).fm0(function () {
  let _coll42 = new _dafny.Map();
  for (const _compr_42 of _dafny.IntegerRange(new BigNumber(260), new BigNumber(408))) {
    let _1277_v69 = _compr_42;
    if (((new BigNumber(260)).isLessThanOrEqualTo(_1277_v69)) && ((_1277_v69).isLessThan(new BigNumber(408)))) {
      _coll42.push([_module.__default.safeModuloInt(_1277_v69, new BigNumber((_1269_v64).length)),_1274_v68.f27]);
    }
  }
  return _coll42;
}(), _1269_v64, _dafny.Seq.of(_1275_v70, _1275_v70), globalState));
              let _1278_v72;
              _1278_v72 = _dafny.Seq.of(!((_1276_v71).dtor_cf7));
              let _1279_v73;
              let _nw211 = new _module.C0();
              _nw211.__ctor();
              _1279_v73 = _nw211;
              let _1280_v74;
              _1280_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1279_v73,p2);
              let _1281_v75;
              _1281_v75 = _dafny.Set.fromElements((((_1280_v74).contains(_1279_v73)) ? ((_1280_v74).get(_1279_v73)) : (p2)));
              let _1282_v76;
              let _nw212 = Array((new BigNumber(25)).toNumber());
              _nw212[(_dafny.ZERO).toNumber()] = (p2).f21;
              _nw212[(_dafny.ONE).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(2)).toNumber()] = p1;
              _nw212[(new BigNumber(3)).toNumber()] = p1;
              _nw212[(new BigNumber(4)).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(5)).toNumber()] = _dafny.Seq.IsPrefixOf(_1269_v64, _module.__default.fm26((p2).f21, globalState));
              _nw212[(new BigNumber(6)).toNumber()] = !(!(!((p2).f21)));
              _nw212[(new BigNumber(7)).toNumber()] = (_1267_v63).equals(_1267_v63);
              _nw212[(new BigNumber(8)).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(9)).toNumber()] = (_this).f24;
              _nw212[(new BigNumber(10)).toNumber()] = (_this).f21;
              _nw212[(new BigNumber(11)).toNumber()] = p1;
              _nw212[(new BigNumber(12)).toNumber()] = ((p2).f20).isLessThanOrEqualTo(p0);
              _nw212[(new BigNumber(13)).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(14)).toNumber()] = (p1) || ((_this).f24);
              _nw212[(new BigNumber(15)).toNumber()] = (_1278_v72)[_module.__default.safeIndex(new BigNumber(769), new BigNumber((_1278_v72).length))];
              _nw212[(new BigNumber(16)).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(17)).toNumber()] = !((_1281_v75).IsProperSubsetOf(_1281_v75));
              _nw212[(new BigNumber(18)).toNumber()] = (_this).f24;
              _nw212[(new BigNumber(19)).toNumber()] = (p2).f21;
              _nw212[(new BigNumber(20)).toNumber()] = (_1274_v68.f27).isLessThanOrEqualTo(new BigNumber(-465));
              _nw212[(new BigNumber(21)).toNumber()] = !((p2).f21);
              _nw212[(new BigNumber(22)).toNumber()] = (((p2).f21) ? (!((p2).f21)) : ((_this).f24));
              _nw212[(new BigNumber(23)).toNumber()] = true;
              _nw212[(new BigNumber(24)).toNumber()] = ((_this).f21) === ((p2).f21);
              _1282_v76 = _nw212;
              let _index212 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1282_v76).length));
              (_1282_v76)[_index212] = (p2).f21;
              let _index213 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1282_v76).length));
              (_1282_v76)[_index213] = p1;
              let _index214 = _module.__default.safeIndex(new BigNumber(261), new BigNumber(((_1274_v68).f28).length));
              ((_1274_v68).f28)[_index214] = p0;
              let _index215 = _module.__default.safeIndex(new BigNumber(261), new BigNumber(((_1274_v68).f28).length));
              let _rhs237 = (_dafny.ZERO).minus(_1274_v68.f27);
              let _rhs238 = _dafny.Seq.UnicodeFromString("qfdxlwohg");
              let _lhs215 = (_1274_v68).f28;
              let _lhs216 = _module.__default.safeIndex(new BigNumber(261), new BigNumber(((_1274_v68).f28).length));
              _lhs215[_lhs216] = _rhs237;
              _1269_v64 = _rhs238;
              let _index216 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1282_v76).length));
              (_1282_v76)[_index216] = (p2).f21;
              (globalState).f5 = ((p2).f20).plus(new BigNumber(86));
            } else {
              let _1283_v77;
              let _nw213 = Array((new BigNumber(11)).toNumber()).fill(false);
              _1283_v77 = _nw213;
              let _1284_v78;
              _1284_v78 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,_1283_v77);
              let _1285_v79;
              _1285_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1269_v64).length),new BigNumber((_1284_v78).length));
              let _1286_v80;
              _1286_v80 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
              let _1287_v81;
              _1287_v81 = _1270_v65;
              let _1288_v82;
              _1288_v82 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,(_1274_v68).fm16(_1286_v80, _1287_v81, globalState));
              let _1289_v83;
              _1289_v83 = _dafny.Seq.of(_1288_v82, _1288_v82, _1288_v82, _1288_v82);
              let _1290_v84;
              _1290_v84 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,_1269_v64);
              (globalState).f15 = ((!((_1274_v68).fm0(_1285_v79, _dafny.Seq.update(_1269_v64, _module.__default.safeIndex(p0, new BigNumber((_1269_v64).length)), _1271_v66), _1289_v83, globalState)) || ((p2).f21)) ? (_dafny.Seq.UnicodeFromString("tyibl")) : ((((_1290_v84).contains((p2).f21)) ? ((_1290_v84).get((p2).f21)) : (_1269_v64))));
              let _1291_v85;
              let _nw214 = new _module.C0();
              _nw214.__ctor();
              _1291_v85 = _nw214;
              (globalState).f18 = _dafny.areEqual(_1270_v65, _dafny.Seq.Concat(_1270_v65, _dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), function (_1292_i6) {
                return (_this).f20;
              })));
              let _1293_v86;
              _1293_v86 = _dafny.MultiSet.fromElements((p2).fm1((p2).f20, new BigNumber((_module.__default.fm28((p2).f20, globalState)).length), (_module.D9.create_DC23(p0, (p2).f21)).dtor_cf32, globalState), (_dafny.ZERO).minus((p2).fm1(new BigNumber((_1267_v63).length), (_dafny.ZERO).minus((p2).f20), p1, globalState)), _1274_v68.f27, new BigNumber(855));
              let _1294_v87;
              _1294_v87 = _module.D10.create_DC25(_1293_v86);
              let _1295_v88;
              let _nw215 = Array((new BigNumber(15)).toNumber()).fill([]);
              _1295_v88 = _nw215;
              let _index217 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1295_v88).length));
              (_1295_v88)[_index217] = (((_this).f21) ? (_1283_v77) : (_1283_v77));
              let _index218 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1283_v77).length));
              (_1283_v77)[_index218] = (_this).f21;
              let _1296_v89;
              _1296_v89 = _dafny.Seq.of(_1285_v79);
              let _1297_v90;
              _1297_v90 = _dafny.Map.Empty.slice().updateUnsafe((_1274_v68).f28,false);
              let _1298_v91;
              _1298_v91 = _dafny.Seq.of(_1283_v77);
              let _index219 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1295_v88).length));
              let _index220 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1283_v77).length));
              let _rhs239 = (p0).isEqualTo((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f20, (_this).f20)));
              let _rhs240 = _module.__default.fm39(globalState);
              let _rhs241 = !((((_1291_v85).fm4(new BigNumber((_1286_v80).cardinality()), _1296_v89, (_dafny.ZERO).minus(new BigNumber((_1297_v90).length)), globalState)) && ((p2).f21)) && ((_this).f24));
              let _rhs242 = (_1298_v91)[_module.__default.safeIndex((p2).f20, new BigNumber((_1298_v91).length))];
              let _rhs243 = (p2).f21;
              let _lhs217 = globalState;
              let _lhs218 = globalState;
              let _lhs219 = _1295_v88;
              let _lhs220 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1295_v88).length));
              let _lhs221 = _1283_v77;
              let _lhs222 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1283_v77).length));
              _lhs217.f18 = _rhs239;
              _1294_v87 = _rhs240;
              _lhs218.f3 = _rhs241;
              _lhs219[_lhs220] = _rhs242;
              _lhs221[_lhs222] = _rhs243;
              (globalState).f14 = (p2).f21;
            }
            let _1299_v92;
            _1299_v92 = _dafny.MultiSet.fromElements((_this).f21);
            let _1300_v93;
            _1300_v93 = _module.D9.create_DC23(new BigNumber(455), p1);
            let _1301_v94;
            let _nw216 = new _module.C0();
            _nw216.__ctor();
            _1301_v94 = _nw216;
            let _1302_v95;
            _1302_v95 = _dafny.MultiSet.fromElements(_1301_v94);
            let _1303_v96;
            _1303_v96 = _dafny.Map.Empty.slice().updateUnsafe((p2).f21,(p2).f21);
            let _1304_v97;
            _1304_v97 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((p2).f20), (p2).f20);
            let _1305_v98;
            let _nw217 = Array((new BigNumber(17)).toNumber());
            _nw217[(_dafny.ZERO).toNumber()] = false;
            _nw217[(_dafny.ONE).toNumber()] = p1;
            _nw217[(new BigNumber(2)).toNumber()] = (p2).f21;
            _nw217[(new BigNumber(3)).toNumber()] = (_1299_v92).IsSubsetOf(_1299_v92);
            _nw217[(new BigNumber(4)).toNumber()] = !((_1300_v93).dtor_cf32);
            _nw217[(new BigNumber(5)).toNumber()] = !_dafny.areEqual(_1269_v64, _1269_v64);
            _nw217[(new BigNumber(6)).toNumber()] = p1;
            _nw217[(new BigNumber(7)).toNumber()] = p1;
            _nw217[(new BigNumber(8)).toNumber()] = (_1302_v95).contains(_1301_v94);
            _nw217[(new BigNumber(9)).toNumber()] = (p2).f21;
            _nw217[(new BigNumber(10)).toNumber()] = (_1274_v68).fm0(_dafny.Map.Empty.slice().updateUnsafe(_1274_v68.f27,(p2).f20), _1269_v64, _dafny.Seq.of(_1303_v96, _1303_v96), globalState);
            _nw217[(new BigNumber(11)).toNumber()] = (p2).f21;
            _nw217[(new BigNumber(12)).toNumber()] = (_this).f21;
            _nw217[(new BigNumber(13)).toNumber()] = (_this).f24;
            _nw217[(new BigNumber(14)).toNumber()] = ((p2).f20).isLessThanOrEqualTo(new BigNumber((_1304_v97).cardinality()));
            _nw217[(new BigNumber(15)).toNumber()] = (_this).f21;
            _nw217[(new BigNumber(16)).toNumber()] = p1;
            _1305_v98 = _nw217;
            _1305_v98 = _1305_v98;
          }
        }
      }
      _1267_v63 = (_1267_v63).update(new BigNumber(605), (p2).f21);
      (globalState).f17 = (p2).f20;
      let _1306_v99;
      _1306_v99 = _dafny.Seq.of(false, (_this).f24, (_this).f24);
      let _1307_v100;
      _1307_v100 = _dafny.Set.fromElements((_this).f20, (_this).f20, new BigNumber((_1306_v99).length), p0);
      let _1308_v101;
      let _nw218 = new _module.C3();
      _nw218.__ctor(_module.__default.safeDivisionInt(p0, (_this).f20), false, (_this).f20, (_1307_v100).IsDisjointFrom(_1307_v100));
      _1308_v101 = _nw218;
      let _1309_v102;
      _1309_v102 = _dafny.MultiSet.fromElements(p1);
      let _1310_v104;
      _1310_v104 = _dafny.Seq.of(function () {
        let _coll43 = new _dafny.Map();
        for (const _compr_43 of (_1267_v63).Keys.Elements) {
          let _1311_v103 = _compr_43;
          if ((_1267_v63).contains(_1311_v103)) {
            _coll43.push([_module.__default.safeModuloInt(_1311_v103, (_1308_v101).f25),(_1308_v101).f26]);
          }
        }
        return _coll43;
      }());
      let _1312_v105;
      _1312_v105 = _dafny.Seq.UnicodeFromString("rukj");
      let _1313_v106;
      _1313_v106 = _dafny.Seq.of(new BigNumber((_1312_v105).length));
      let _1314_v107;
      let _nw219 = Array((new BigNumber(23)).toNumber());
      _nw219[(_dafny.ZERO).toNumber()] = (p2).f20;
      _nw219[(_dafny.ONE).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(2)).toNumber()] = ((_1308_v101).f25).minus(p0);
      _nw219[(new BigNumber(3)).toNumber()] = new BigNumber((_1309_v102).cardinality());
      _nw219[(new BigNumber(4)).toNumber()] = new BigNumber((_1310_v104).length);
      _nw219[(new BigNumber(5)).toNumber()] = (new BigNumber(325)).plus((_1308_v101).f25);
      _nw219[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_1308_v101).f25, (_this).fm1((p2).f20, (_1308_v101).f25, (_1308_v101).f26, globalState));
      _nw219[(new BigNumber(7)).toNumber()] = (_this).fm1(p0, (_1308_v101).f25, (_this).f24, globalState);
      _nw219[(new BigNumber(8)).toNumber()] = new BigNumber((_1312_v105).length);
      _nw219[(new BigNumber(9)).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(10)).toNumber()] = p0;
      _nw219[(new BigNumber(11)).toNumber()] = (p2).fm1(new BigNumber((_1306_v99).length), (_this).f20, (p2).f21, globalState);
      _nw219[(new BigNumber(12)).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(13)).toNumber()] = p0;
      _nw219[(new BigNumber(14)).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(15)).toNumber()] = (p2).fm1(p0, p0, (_this).f24, globalState);
      _nw219[(new BigNumber(16)).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(17)).toNumber()] = new BigNumber((_1313_v106).length);
      _nw219[(new BigNumber(18)).toNumber()] = (_this).f20;
      _nw219[(new BigNumber(19)).toNumber()] = (new BigNumber(-384)).multipliedBy((_this).f20);
      _nw219[(new BigNumber(20)).toNumber()] = (_1308_v101).f25;
      _nw219[(new BigNumber(21)).toNumber()] = (p2).f20;
      _nw219[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_this).f20);
      _1314_v107 = _nw219;
      r0 = _1314_v107;
      r1 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(500), p0), (_dafny.ZERO).minus((_this).f20));
      return [r0, r1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
      this.f23 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f23, f20, f21) {
      let _this = this;
      (_this).f23 = f23;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return !((_this).f21);
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Set.fromElements((_this).f20, (_this).f20)).length)).plus(new BigNumber(((function () {
        let _coll44 = new _dafny.Map();
        for (const _compr_44 of _dafny.IntegerRange(new BigNumber(-451), new BigNumber(-31))) {
          let _1315_v0 = _compr_44;
          if (((new BigNumber(-451)).isLessThanOrEqualTo(_1315_v0)) && ((_1315_v0).isLessThan(new BigNumber(-31)))) {
            _coll44.push([(_1315_v0).plus((_this).f20),_this.f23]);
          }
        }
        return _coll44;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f21))).length));
    };
    fm5(p0, globalState) {
      let _this = this;
      return _this.f23;
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1316_v0;
      _1316_v0 = _dafny.Seq.UnicodeFromString("gjgjhukoy");
      let _1317_v1;
      _1317_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1316_v0).length),(_this).f21);
      let _1318_v2;
      _1318_v2 = _dafny.MultiSet.fromElements(_1317_v1);
      let _1319_v3;
      _1319_v3 = _dafny.Seq.of(_module.__default.fm6(false, _1318_v2, globalState));
      let _1320_v4;
      _1320_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7((_this).f20, p0, globalState),p1);
      let _1321_v5;
      _1321_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1319_v3,(p0).plus(new BigNumber((_1320_v4).length)));
      let _1322_v6;
      _1322_v6 = _dafny.Set.fromElements((_this).fm5(!(p1), globalState), (_this).f21);
      let _1323_v7;
      _1323_v7 = _dafny.Seq.of((_this).f20, (_dafny.ZERO).minus(new BigNumber((_1322_v6).length)));
      _1321_v5 = (_1321_v5).update((((_this).f21) ? (_1319_v3) : (_module.__default.fm8((_this).f20, (_this).f21, _1323_v7, _this.f23, globalState))), new BigNumber(820));
      let _1324_v9;
      _1324_v9 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1));
      let _1325_v10;
      _1325_v10 = new _dafny.CodePoint('p'.codePointAt(0));
      let _1326_v11;
      _1326_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber(121));
      let _1327_v12;
      _1327_v12 = _dafny.Seq.of((_this).fm0(_1326_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), ((_1328_v10) => function (_1329_i0) {
        return _1328_v10;
      })(_1325_v10)), _1324_v9, globalState), _this.f23, !(_this.f23));
      let _1330_v13;
      _1330_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,p1);
      let _1331_v14;
      let _nw220 = Array((new BigNumber(22)).toNumber());
      _nw220[(_dafny.ZERO).toNumber()] = !((_this).f21);
      _nw220[(_dafny.ONE).toNumber()] = p1;
      _nw220[(new BigNumber(2)).toNumber()] = p1;
      _nw220[(new BigNumber(3)).toNumber()] = p1;
      _nw220[(new BigNumber(4)).toNumber()] = !((p1) === (_this.f23));
      _nw220[(new BigNumber(5)).toNumber()] = true;
      _nw220[(new BigNumber(6)).toNumber()] = ((_this).f20).isLessThanOrEqualTo((_this).f20);
      _nw220[(new BigNumber(7)).toNumber()] = p1;
      _nw220[(new BigNumber(8)).toNumber()] = (_this).f21;
      _nw220[(new BigNumber(9)).toNumber()] = false;
      _nw220[(new BigNumber(10)).toNumber()] = p1;
      _nw220[(new BigNumber(11)).toNumber()] = (_this).fm0(function () {
        let _coll45 = new _dafny.Map();
        for (const _compr_45 of (_1323_v7).Elements) {
          let _1332_v8 = _compr_45;
          if (_dafny.Seq.contains(_1323_v7, _1332_v8)) {
            _coll45.push([_module.__default.safeDivisionInt(_1332_v8, p0),(_this).f20]);
          }
        }
        return _coll45;
      }(), _dafny.Seq.UnicodeFromString("i"), _1324_v9, globalState);
      _nw220[(new BigNumber(12)).toNumber()] = !_dafny.Seq.contains(_1316_v0, _1325_v10);
      _nw220[(new BigNumber(13)).toNumber()] = ((_this).f20).isLessThanOrEqualTo((_this).f20);
      _nw220[(new BigNumber(14)).toNumber()] = _this.f23;
      _nw220[(new BigNumber(15)).toNumber()] = _this.f23;
      _nw220[(new BigNumber(16)).toNumber()] = (_1327_v12)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("eghu")).length), new BigNumber((_1327_v12).length))];
      _nw220[(new BigNumber(17)).toNumber()] = _this.f23;
      _nw220[(new BigNumber(18)).toNumber()] = ((_this).f20).isEqualTo(new BigNumber((_1330_v13).length));
      _nw220[(new BigNumber(19)).toNumber()] = (_this).f21;
      _nw220[(new BigNumber(20)).toNumber()] = !((_this).f21) || (_this.f23);
      _nw220[(new BigNumber(21)).toNumber()] = ((_this).f20).isLessThanOrEqualTo(p0);
      _1331_v14 = _nw220;
      let _index221 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
      (_1331_v14)[_index221] = p1;
      let _index222 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
      (_1331_v14)[_index222] = !(_this.f23);
      let _1333_i1;
      _1333_i1 = _dafny.ZERO;
      L5: {
        while ((new BigNumber((_1316_v0).length)).isLessThanOrEqualTo((_this).f20)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1333_i1)) {
              break L5;
            }
            _1333_i1 = (_1333_i1).plus(_dafny.ONE);
            let _1334_v15;
            let _init36 = ((_1335_p0) => function (_1336_i2) {
              return _module.__default.safeModuloInt(_1336_i2, _1335_p0);
            })(p0);
            let _nw221 = Array((new BigNumber(15)).toNumber());
            for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw221.length); _i0_36++) {
              _nw221[_i0_36] = _init36(new BigNumber(_i0_36));
            }
            _1334_v15 = _nw221;
            let _index223 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length));
            (_1334_v15)[_index223] = p0;
            let _1337_v16;
            _1337_v16 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20);
            let _index224 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length));
            (_1334_v15)[_index224] = (p0).minus(new BigNumber((_1337_v16).cardinality()));
            (globalState).f18 = _this.f23;
            let _1338_v17;
            _1338_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1316_v0,(_1334_v15)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length))]);
            let _index225 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length));
            (_1334_v15)[_index225] = new BigNumber((((_1338_v17).Merge(_1338_v17)).Merge(_1338_v17)).length);
            if (((_this).f20).isLessThan((_this).f20)) {
              let _1339_v18;
              let _nw222 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
              _1339_v18 = _nw222;
              let _1340_v19;
              _1340_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1339_v18,!((_this).f21) || (_this.f23));
              let _index226 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length));
              let _rhs244 = _1331_v14;
              let _rhs245 = new BigNumber(273);
              let _rhs246 = (_this).fm0(_1326_v11, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), ((_1341_v10) => function (_1342_i3) {
                return _1341_v10;
              })(_1325_v10)), _1316_v0), _1324_v9, globalState);
              let _rhs247 = (((_1340_v19).contains(_1339_v18)) ? ((_1340_v19).get(_1339_v18)) : ((_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))]));
              let _lhs223 = _1334_v15;
              let _lhs224 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length));
              let _lhs225 = globalState;
              let _lhs226 = globalState;
              _1331_v14 = _rhs244;
              _lhs223[_lhs224] = _rhs245;
              _lhs225.f18 = _rhs246;
              _lhs226.f3 = _rhs247;
              let _1343_v20;
              let _init37 = function (_1344_i4) {
                return (_this).f21;
              };
              let _nw223 = Array((new BigNumber(21)).toNumber());
              for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw223.length); _i0_37++) {
                _nw223[_i0_37] = _init37(new BigNumber(_i0_37));
              }
              _1343_v20 = _nw223;
              let _1345_v21;
              _1345_v21 = _dafny.Seq.of(_1343_v20);
              _1345_v21 = _dafny.Seq.of(_1343_v20, _1343_v20);
              (globalState).f18 = (new BigNumber((_1327_v12).length)).isLessThan(p0);
              let _index227 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
              (_1331_v14)[_index227] = (_this).f21;
              let _1346_v22;
              _1346_v22 = _dafny.Seq.of(_1331_v14, _1331_v14);
              let _1347_v23;
              _1347_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1346_v22);
              _1347_v23 = (_1347_v23).update((_this).f20, _1346_v22);
            } else {
              (globalState).f4 = (_this).f21;
              let _1348_v24;
              _1348_v24 = _dafny.MultiSet.fromElements(_1331_v14, _1331_v14);
              let _1349_v25;
              _1349_v25 = _dafny.Seq.of(_1348_v24);
              let _1350_v26;
              _1350_v26 = _dafny.Seq.of(_dafny.Seq.update(_1323_v7, _module.__default.safeIndex(p0, new BigNumber((_1323_v7).length)), (_this).f20));
              let _1351_v27;
              _1351_v27 = (_1350_v26)[_module.__default.safeIndex((_this).f20, new BigNumber((_1350_v26).length))];
              let _1352_v28;
              _1352_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1334_v15)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length))],_1348_v24);
              let _1353_v29;
              _1353_v29 = _module.D2.create_DC2(_1348_v24);
              let _1354_v30;
              _1354_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f20);
              let _1355_v31;
              let _nw224 = Array((new BigNumber(21)).toNumber());
              _nw224[(_dafny.ZERO).toNumber()] = _1348_v24;
              _nw224[(_dafny.ONE).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(2)).toNumber()] = (_1349_v25)[_module.__default.safeIndex(new BigNumber((_module.__default.fm9(new BigNumber((_1316_v0).length), _1351_v27, (_this).f20, globalState)).length), new BigNumber((_1349_v25).length))];
              _nw224[(new BigNumber(3)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(4)).toNumber()] = (_1349_v25)[_module.__default.safeIndex((_1334_v15)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1334_v15).length))], new BigNumber((_1349_v25).length))];
              _nw224[(new BigNumber(5)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(6)).toNumber()] = (((_1352_v28).contains((_this).f20)) ? ((_1352_v28).get((_this).f20)) : (_1348_v24));
              _nw224[(new BigNumber(7)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(8)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(9)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(10)).toNumber()] = (_1348_v24).update(_1331_v14, _module.__default.abs(p0));
              _nw224[(new BigNumber(11)).toNumber()] = (_1353_v29).dtor_cf2;
              _nw224[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.fromElements(_1331_v14)).Difference(_1348_v24);
              _nw224[(new BigNumber(13)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(14)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(15)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(16)).toNumber()] = (_1348_v24).update(_1331_v14, _module.__default.abs((_this).fm1((_this).f20, new BigNumber((_1354_v30).length), true, globalState)));
              _nw224[(new BigNumber(17)).toNumber()] = (_1348_v24).Intersect(_dafny.MultiSet.fromElements(_1331_v14, _1331_v14));
              _nw224[(new BigNumber(18)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(19)).toNumber()] = _1348_v24;
              _nw224[(new BigNumber(20)).toNumber()] = _1348_v24;
              _1355_v31 = _nw224;
              let _index228 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1355_v31).length));
              (_1355_v31)[_index228] = _1348_v24;
              let _index229 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1355_v31).length));
              (_1355_v31)[_index229] = (_1349_v25)[_module.__default.safeIndex((_this).f20, new BigNumber((_1349_v25).length))];
              (globalState).f3 = !((_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))]);
              (globalState).f18 = !((_this).f21);
              let _1356_v32;
              let _nw225 = new _module.C0();
              _nw225.__ctor();
              _1356_v32 = _nw225;
            }
          }
        }
      }
      if (_this.f23) {
        let _1357_v33;
        let _nw226 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _1357_v33 = _nw226;
        let _index230 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        (_1357_v33)[_index230] = (((_this).f21) ? ((_this).f20) : ((_this).f20));
        let _index231 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        (_1357_v33)[_index231] = _module.__default.safeDivisionInt((_this).f20, (_this).f20);
        let _index232 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        (_1357_v33)[_index232] = new BigNumber(((_1326_v11).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber(204)))).length);
        _1325_v10 = _1325_v10;
        let _1358_v34;
        _1358_v34 = _dafny.Seq.of(_1331_v14, _1331_v14);
        let _1359_v35;
        _1359_v35 = _module.D3.create_DC5(_1358_v34);
        let _1360_v36;
        _1360_v36 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), _1325_v10, _1325_v10);
        let _index233 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        let _rhs248 = (_this).f20;
        let _rhs249 = (((true) ? (_1359_v35) : (_1359_v35))).dtor_cf9;
        let _rhs250 = (_dafny.MultiSet.fromElements(_1325_v10)).IsSubsetOf(_1360_v36);
        let _lhs227 = _1357_v33;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        _lhs227[_lhs228] = _rhs248;
        _1358_v34 = _rhs249;
        r1 = _rhs250;
        let _index234 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
        let _index235 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        let _rhs251 = !(((_this).f20).isLessThan(new BigNumber(714)));
        let _rhs252 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), function (_1361_i5) {
          return (_this).f20;
        });
        let _rhs253 = _module.__default.safeModuloInt((p0).multipliedBy(new BigNumber((_1327_v12).length)), p0);
        let _lhs229 = _1331_v14;
        let _lhs230 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
        let _lhs231 = _1357_v33;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1357_v33).length));
        _lhs229[_lhs230] = _rhs251;
        _1323_v7 = _rhs252;
        _lhs231[_lhs232] = _rhs253;
      } else {
        let _1362_v37;
        _1362_v37 = _dafny.MultiSet.fromElements(_1331_v14);
        let _1363_v38;
        _1363_v38 = _module.D2.create_DC2((_1362_v37).Intersect((_dafny.MultiSet.fromElements(_1331_v14, _1331_v14, _1331_v14, _1331_v14, _1331_v14)).update(_1331_v14, _module.__default.abs(new BigNumber(438)))));
        _1363_v38 = _1363_v38;
        let _1364_v39;
        _1364_v39 = _1323_v7;
        let _1365_v40;
        _1365_v40 = _dafny.Set.fromElements(_1364_v39);
        let _1366_v41;
        _1366_v41 = _dafny.MultiSet.fromElements(_module.__default.fm10(true, (_this).f21, new BigNumber((_1316_v0).length), globalState));
        (globalState).f5 = new BigNumber(((_1365_v40).Intersect(function () {
          let _coll46 = new _dafny.Set();
          for (const _compr_46 of (_1366_v41).Elements) {
            let _1367_v42 = _compr_46;
            if ((_1366_v41).contains(_1367_v42)) {
              _coll46.add(_1367_v42);
            }
          }
          return _coll46;
        }())).length);
        let _1368_v43;
        let _nw227 = new _module.C0();
        _nw227.__ctor();
        _1368_v43 = _nw227;
        let _1369_v44;
        let _nw228 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1369_v44 = _nw228;
        let _index236 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1369_v44).length));
        (_1369_v44)[_index236] = p0;
        let _index237 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1369_v44).length));
        let _rhs254 = (_dafny.ZERO).minus(p0);
        let _rhs255 = (_this).f20;
        let _rhs256 = ((_this).f20).plus(_module.__default.safeModuloInt((_this).f20, new BigNumber((_1316_v0).length)));
        let _lhs233 = _1369_v44;
        let _lhs234 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1369_v44).length));
        let _lhs235 = globalState;
        let _lhs236 = globalState;
        _lhs233[_lhs234] = _rhs254;
        _lhs235.f17 = _rhs255;
        _lhs236.f17 = _rhs256;
        (globalState).f18 = !(false) || (_this.f23);
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1331_v14).length))) {
        let _1370_i6 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1370_i6)) && ((_1370_i6).isLessThan(new BigNumber((_1331_v14).length))))) {
          (_1331_v14)[(_1370_i6)] = !(p1);
        }
      }
      let _1371_v45;
      _1371_v45 = _dafny.MultiSet.fromElements(_this.f23, _this.f23);
      let _1372_v46;
      _1372_v46 = _module.D2.create_DC3(new BigNumber((_1327_v12).length), !(_1371_v45).contains(_this.f23), (_this).f21, p0, ((_this).fm5((_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))], globalState)) || ((_this).f21));
      let _source16 = _1372_v46;
      if (_source16.is_DC3) {
        let _1373___mcc_h0 = (_source16).cf3;
        let _1374___mcc_h1 = (_source16).cf4;
        let _1375___mcc_h2 = (_source16).cf5;
        let _1376___mcc_h3 = (_source16).cf6;
        let _1377___mcc_h4 = (_source16).cf7;
        let _1378_cf7 = _1377___mcc_h4;
        let _1379_cf6 = _1376___mcc_h3;
        let _1380_cf5 = _1375___mcc_h2;
        let _1381_cf4 = _1374___mcc_h1;
        let _1382_cf3 = _1373___mcc_h0;
        let _1383_v47;
        let _nw229 = Array((new BigNumber(6)).toNumber()).fill(_module.D2.Default());
        _1383_v47 = _nw229;
        let _1384_v48;
        _1384_v48 = _dafny.MultiSet.fromElements(_1331_v14, (_module.D4.create_DC9(_1331_v14)).dtor_cf14, _1331_v14);
        let _1385_v49;
        _1385_v49 = _module.D2.create_DC2(_1384_v48);
        let _index238 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1383_v47).length));
        (_1383_v47)[_index238] = _1385_v49;
        let _index239 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1383_v47).length));
        (_1383_v47)[_index239] = _1385_v49;
        let _index240 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length));
        (_1331_v14)[_index240] = (_this).f21;
        (globalState).f17 = p0;
        let _1386_v50;
        _1386_v50 = _dafny.Seq.of(_1316_v0);
        let _1387_v51;
        _1387_v51 = _dafny.Seq.of(_1316_v0, _1316_v0, _dafny.Seq.Concat(_1316_v0, _1316_v0), _1316_v0, (_1386_v50)[_module.__default.safeIndex(_1379_cf6, new BigNumber((_1386_v50).length))]);
        _1386_v50 = _1387_v51;
      } else if (_source16.is_DC2) {
        let _1388___mcc_h5 = (_source16).cf2;
        let _1389_cf2 = _1388___mcc_h5;
        (globalState).f5 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(696)), ((_1390_v10) => function (_1391_i7) {
          return _1390_v10;
        })(_1325_v10))).length);
        (_this).f23 = (_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))];
        (globalState).f17 = (_this).f20;
        let _1392_v52;
        let _init38 = ((_1393_v0) => function (_1394_i8) {
          return _dafny.Seq.Concat(_1393_v0, _1393_v0);
        })(_1316_v0);
        let _nw230 = Array((new BigNumber(20)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw230.length); _i0_38++) {
          _nw230[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1392_v52 = _nw230;
        let _index241 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1392_v52).length));
        (_1392_v52)[_index241] = _1316_v0;
        let _index242 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1392_v52).length));
        (_1392_v52)[_index242] = _1316_v0;
      } else {
        let _1395___mcc_h6 = (_source16).cf8;
        let _1396_cf8 = _1395___mcc_h6;
        let _1397_v53;
        let _nw231 = new _module.C0();
        _nw231.__ctor();
        _1397_v53 = _nw231;
        let _1398_v54;
        let _init39 = ((_1399_v10) => function (_1400_i9) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-378)), ((_1401_v10) => function (_1402_i10) {
            return _1401_v10;
          })(_1399_v10));
        })(_1325_v10);
        let _nw232 = Array((new BigNumber(22)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw232.length); _i0_39++) {
          _nw232[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1398_v54 = _nw232;
        let _index243 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1398_v54).length));
        (_1398_v54)[_index243] = _1316_v0;
        let _index244 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1398_v54).length));
        (_1398_v54)[_index244] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iby"), _1316_v0);
        if (((_this).f20).isLessThan(p0)) {
          (globalState).f17 = (_this).fm1((_this).f20, (_this).f20, ((p1) ? (p1) : ((_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))])), globalState);
          let _1403_v55;
          _1403_v55 = _dafny.Map.Empty.slice().updateUnsafe((_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))],_1398_v54);
          let _1404_v56;
          _1404_v56 = _dafny.MultiSet.fromElements((_this).f20, p0, (_this).f20);
          _1398_v54 = (((_1403_v55).contains((_1404_v56).IsProperSubsetOf(_1404_v56))) ? ((_1403_v55).get((_1404_v56).IsProperSubsetOf(_1404_v56))) : (_1398_v54));
          let _1405_v57;
          _1405_v57 = _dafny.Map.Empty.slice().updateUnsafe((_1372_v46).dtor_cf6,_1404_v56);
          _1405_v57 = (_1405_v57).update(new BigNumber(797), _module.__default.fm11(new BigNumber((_1323_v7).length), (_this).f21, globalState));
          (_this).f23 = _dafny.Seq.IsPrefixOf(_module.__default.fm12(p0, _1325_v10, (_this).f21, (((_1326_v11).contains((_this).f20)) ? ((_1326_v11).get((_this).f20)) : ((_this).f20)), globalState), _dafny.Seq.Concat((_1398_v54)[_module.__default.safeIndex(new BigNumber(730), new BigNumber((_1398_v54).length))], _dafny.Seq.UnicodeFromString("avxghk")));
          _1323_v7 = ((p1) ? (_1323_v7) : (_1323_v7));
        } else {
          let _1406_v58;
          let _nw233 = new _module.C0();
          _nw233.__ctor();
          _1406_v58 = _nw233;
          let _1407_v59;
          _1407_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1317_v1);
          let _1408_v60;
          _1408_v60 = _dafny.Seq.of(_1326_v11);
          let _1409_v61;
          _1409_v61 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC3(p0, _this.f23, (_this).f21, p0, true),(_1397_v53).fm4(p0, _1408_v60, (_this).f20, globalState));
          let _1410_v62;
          _1410_v62 = _dafny.Seq.of(_1331_v14);
          let _1411_v63;
          _1411_v63 = _module.D3.create_DC6((_this).f20);
          let _1412_v64;
          _1412_v64 = _dafny.Set.fromElements(p0, (_this).f20, new BigNumber(((_1398_v54)[_module.__default.safeIndex(new BigNumber(730), new BigNumber((_1398_v54).length))]).length));
          let _1413_v65;
          _1413_v65 = _module.D4.create_DC11(new BigNumber((_1327_v12).length), (_this).f21, p1);
          let _1414_v66;
          let _nw234 = Array((new BigNumber(25)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw234[(_dafny.ONE).toNumber()] = new BigNumber((_1407_v59).length);
          _nw234[(new BigNumber(2)).toNumber()] = p0;
          _nw234[(new BigNumber(3)).toNumber()] = p0;
          _nw234[(new BigNumber(4)).toNumber()] = new BigNumber((_1327_v12).length);
          _nw234[(new BigNumber(5)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1372_v46,(_this).f21)).Merge(_1409_v61)).length);
          _nw234[(new BigNumber(6)).toNumber()] = new BigNumber((_1410_v62).length);
          _nw234[(new BigNumber(7)).toNumber()] = p0;
          _nw234[(new BigNumber(8)).toNumber()] = (p0).multipliedBy(new BigNumber(418));
          _nw234[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(p0, (_this).f20);
          _nw234[(new BigNumber(10)).toNumber()] = ((_this).f20).minus((_1411_v63).dtor_cf10);
          _nw234[(new BigNumber(11)).toNumber()] = new BigNumber((_1412_v64).length);
          _nw234[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw234[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(((_this).f20).multipliedBy(p0));
          _nw234[(new BigNumber(14)).toNumber()] = new BigNumber((_module.__default.fm12((_this).f20, new _dafny.CodePoint('n'.codePointAt(0)), (_1397_v53).fm3(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), ((_1415_v10) => function (_1416_i11) {
            return _1415_v10;
          })(_1325_v10))).length), false, globalState), (_this).fm1(new BigNumber(-961), new BigNumber(442), (_this).f21, globalState), globalState)).length);
          _nw234[(new BigNumber(15)).toNumber()] = new BigNumber(390);
          _nw234[(new BigNumber(16)).toNumber()] = (_this).f20;
          _nw234[(new BigNumber(17)).toNumber()] = new BigNumber((_1323_v7).length);
          _nw234[(new BigNumber(18)).toNumber()] = (_this).f20;
          _nw234[(new BigNumber(19)).toNumber()] = (p0).plus((_1413_v65).dtor_cf17);
          _nw234[(new BigNumber(20)).toNumber()] = p0;
          _nw234[(new BigNumber(21)).toNumber()] = (_this).f20;
          _nw234[(new BigNumber(22)).toNumber()] = (_this).f20;
          _nw234[(new BigNumber(23)).toNumber()] = p0;
          _nw234[(new BigNumber(24)).toNumber()] = p0;
          _1414_v66 = _nw234;
          let _index245 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1414_v66).length));
          (_1414_v66)[_index245] = (_dafny.ZERO).minus(p0);
          let _1417_v67;
          _1417_v67 = _dafny.MultiSet.fromElements(new BigNumber((_1330_v13).length));
          let _1418_v68;
          _1418_v68 = _dafny.Seq.of(_dafny.MultiSet.fromElements(p0), _1417_v67, _1417_v67);
          let _index246 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1414_v66).length));
          (_1414_v66)[_index246] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(498)), ((_1419_p0) => function (_1420_i12) {
            return _dafny.MultiSet.fromElements(_1419_p0);
          })(p0)), _1418_v68)).length));
          let _1421_v69;
          let _nw235 = new _module.C0();
          _nw235.__ctor();
          _1421_v69 = _nw235;
          let _1422_v70;
          let _nw236 = new _module.C0();
          _nw236.__ctor();
          _1422_v70 = _nw236;
          _1414_v66 = _1414_v66;
        }
        (globalState).f5 = new BigNumber((function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of (_dafny.Seq.of(_1323_v7)).Elements) {
            let _1423_v71 = _compr_47;
            if (_dafny.Seq.contains(_dafny.Seq.of(_1323_v7), _1423_v71)) {
              _coll47.push([_1423_v71,p1]);
            }
          }
          return _coll47;
        }()).length);
      }
      let _1424_v72;
      _1424_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1331_v14,(_this).f21);
      let _1425_v73;
      let _nw237 = new _module.C0();
      _nw237.__ctor();
      _1425_v73 = _nw237;
      let _1426_v74;
      _1426_v74 = _dafny.Set.fromElements(_1425_v73);
      let _1427_v75;
      let _nw238 = new _module.C4();
      _nw238.__ctor(p1, new BigNumber(471), true);
      _1427_v75 = _nw238;
      let _1428_v76;
      _1428_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_1427_v75);
      let _1429_v77;
      _1429_v77 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), (_this).f20);
      let _1430_v78;
      _1430_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_1429_v77);
      let _1431_v79;
      _1431_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1427_v75,p0);
      let _1432_v80;
      _1432_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_1371_v45);
      let _1433_v81;
      let _nw239 = Array((new BigNumber(16)).toNumber());
      _nw239[(_dafny.ZERO).toNumber()] = (_this).f20;
      _nw239[(_dafny.ONE).toNumber()] = p0;
      _nw239[(new BigNumber(2)).toNumber()] = new BigNumber((_1426_v74).length);
      _nw239[(new BigNumber(3)).toNumber()] = new BigNumber(815);
      _nw239[(new BigNumber(4)).toNumber()] = p0;
      _nw239[(new BigNumber(5)).toNumber()] = new BigNumber((_1428_v76).length);
      _nw239[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw239[(new BigNumber(7)).toNumber()] = (_1427_v75).f20;
      _nw239[(new BigNumber(8)).toNumber()] = (_this).f20;
      _nw239[(new BigNumber(9)).toNumber()] = (_1427_v75).f20;
      _nw239[(new BigNumber(10)).toNumber()] = (_this).f20;
      _nw239[(new BigNumber(11)).toNumber()] = p0;
      _nw239[(new BigNumber(12)).toNumber()] = new BigNumber(((((_1430_v78).contains(new BigNumber((_1431_v79).length))) ? ((_1430_v78).get(new BigNumber((_1431_v79).length))) : (_1429_v77))).cardinality());
      _nw239[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1371_v45).cardinality()));
      _nw239[(new BigNumber(14)).toNumber()] = new BigNumber((_1326_v11).length);
      _nw239[(new BigNumber(15)).toNumber()] = new BigNumber((_1432_v80).length);
      _1433_v81 = _nw239;
      let _1434_v82;
      _1434_v82 = _dafny.Seq.of(_1433_v81, _1433_v81);
      let _1435_v83;
      _1435_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1424_v72).length),_dafny.Seq.update(_1434_v82, _module.__default.safeIndex((_1323_v7)[_module.__default.safeIndex(new BigNumber(244), new BigNumber((_1323_v7).length))], new BigNumber((_1434_v82).length)), _1433_v81));
      r0 = _dafny.Seq.Concat((((_1435_v83).contains(p0)) ? ((_1435_v83).get(p0)) : (_1434_v82)), _1434_v82);
      r1 = (_1331_v14)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1331_v14).length))];
      return [r0, r1];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f20 = _dafny.ZERO;
      this._f21 = false;
      this._f22 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    __ctor(f22, f20, f21) {
      let _this = this;
      (_this)._f22 = f22;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f21;
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_this).f20)).minus((_this).f20);
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1436_i0;
      _1436_i0 = _dafny.ZERO;
      L6: {
        while ((_this).f21) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1436_i0)) {
              break L6;
            }
            _1436_i0 = (_1436_i0).plus(_dafny.ONE);
            let _1437_v0;
            _1437_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), ((_1438_p0) => function (_1439_i1) {
              return _1438_p0;
            })(p0));
            let _1440_v1;
            _1440_v1 = _dafny.MultiSet.fromElements(p0);
            let _1441_v2;
            _1441_v2 = _dafny.Seq.of(_1440_v1, _1440_v1, _1440_v1);
            let _1442_v3;
            _1442_v3 = _dafny.MultiSet.fromElements(p0, new BigNumber((_1441_v2).length));
            let _1443_v4;
            _1443_v4 = _dafny.Seq.of((((_1442_v3).contains(p0)) ? ((_1442_v3).get(p0)) : ((_this).f20)));
            let _1444_v5;
            _1444_v5 = _dafny.Map.Empty.slice().updateUnsafe((true) === ((_this).f21),(_1443_v4));
            let _1445_v6;
            _1445_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber(-218));
            let _1446_v7;
            _1446_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
            let _1447_v8;
            _1447_v8 = (_this).f21;
            let _1448_v9;
            _1448_v9 = _dafny.Seq.of(_1446_v7, _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_1447_v8)));
            _1444_v5 = (_1444_v5).update((_this).fm0(_1445_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(787)), function (_1449_i2) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            }), _1448_v9, globalState), _1443_v4);
            let _source17 = _1437_v0;
            let _1450___mcc_h0 = _source17;
            let _1451_cf0 = _1450___mcc_h0;
            let _1452_v10;
            _1452_v10 = _dafny.Set.fromElements(p1, (_this).f21);
            (globalState).f3 = !(_dafny.Set.fromElements((_this).f21)).equals(_1452_v10);
            let _1453_v11;
            _1453_v11 = _dafny.MultiSet.fromElements((_this).f21);
            let _1454_v12;
            let _nw240 = Array((new BigNumber(8)).toNumber());
            _nw240[(_dafny.ZERO).toNumber()] = p0;
            _nw240[(_dafny.ONE).toNumber()] = new BigNumber((_1453_v11).cardinality());
            _nw240[(new BigNumber(2)).toNumber()] = (_this).fm1(p0, new BigNumber(-743), p1, globalState);
            _nw240[(new BigNumber(3)).toNumber()] = p0;
            _nw240[(new BigNumber(4)).toNumber()] = (p0).multipliedBy(p0);
            _nw240[(new BigNumber(5)).toNumber()] = new BigNumber(296);
            _nw240[(new BigNumber(6)).toNumber()] = (_1451_cf0)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f20), new BigNumber((_1451_cf0).length))];
            _nw240[(new BigNumber(7)).toNumber()] = (p0).multipliedBy(p0);
            _1454_v12 = _nw240;
            let _1455_v13;
            _1455_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v4,_1454_v12);
            _1454_v12 = (((_1455_v13).contains((((_this).f21) ? (_1437_v0) : (_1437_v0)))) ? ((_1455_v13).get((((_this).f21) ? (_1437_v0) : (_1437_v0)))) : (_1454_v12));
            let _1456_v14;
            let _init40 = ((_1457_v8) => function (_1458_i3) {
              return _1457_v8;
            })(_1447_v8);
            let _nw241 = Array((new BigNumber(28)).toNumber());
            for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw241.length); _i0_40++) {
              _nw241[_i0_40] = _init40(new BigNumber(_i0_40));
            }
            _1456_v14 = _nw241;
            let _index247 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1456_v14).length));
            (_1456_v14)[_index247] = _1447_v8;
            let _index248 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1456_v14).length));
            let _rhs257 = _1447_v8;
            let _rhs258 = new BigNumber((_1452_v10).length);
            let _rhs259 = ((((_this).f21) ? (p1) : (_1447_v8)));
            let _rhs260 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f22).length)), _1443_v4), _1451_cf0);
            let _rhs261 = p1;
            let _lhs237 = _1456_v14;
            let _lhs238 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1456_v14).length));
            let _lhs239 = globalState;
            let _lhs240 = globalState;
            let _lhs241 = globalState;
            _lhs237[_lhs238] = _rhs257;
            _lhs239.f5 = _rhs258;
            _lhs240.f14 = _rhs259;
            _1451_cf0 = _rhs260;
            _lhs241.f14 = _rhs261;
            let _index249 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1454_v12).length));
            (_1454_v12)[_index249] = _module.__default.safeDivisionInt(p0, (_this).fm1(p0, (_this).f20, p1, globalState));
            let _index250 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1454_v12).length));
            (_1454_v12)[_index250] = new BigNumber(157);
            _1446_v7 = (_1446_v7).update(!((_this).f21), p1);
            let _1459_v15;
            let _nw242 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
            _1459_v15 = _nw242;
            let _1460_v16;
            _1460_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1459_v15);
            _1460_v16 = (_1460_v16).update((_this).f21, _1459_v15);
          }
        }
      }
      let _hi4 = p0;
      for (let _1461_i4 = (_dafny.ZERO).minus(new BigNumber(-898)); _1461_i4.isLessThan(_hi4); _1461_i4 = _1461_i4.plus(_dafny.ONE)) {
        let _1462_v17;
        _1462_v17 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1463_v18;
        _1463_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1462_v17,_module.__default.fm2(!((_this).f21), (_this).f20, p1, globalState));
        let _1464_v19;
        _1464_v19 = (_this).f21;
        _1463_v18 = (_1463_v18).update(_1462_v17, _1464_v19);
        (globalState).f5 = (_this).f20;
        let _1465_v20;
        let _nw243 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _1465_v20 = _nw243;
        let _index251 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1465_v20).length));
        (_1465_v20)[_index251] = (p0).multipliedBy(new BigNumber(((_this).f22).length));
        let _index252 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1465_v20).length));
        let _rhs262 = (_dafny.ZERO).minus((_this).f20);
        let _rhs263 = (_this).f22;
        let _lhs242 = _1465_v20;
        let _lhs243 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1465_v20).length));
        let _lhs244 = globalState;
        _lhs242[_lhs243] = _rhs262;
        _lhs244.f15 = _rhs263;
        let _1466_v21;
        _1466_v21 = _dafny.Seq.of(p0);
        let _source18 = _1466_v21;
        let _1467___mcc_h1 = _source18;
        let _1468_cf0 = _1467___mcc_h1;
        (globalState).f18 = (_this).f21;
        let _1469_v22;
        let _nw244 = new _module.C0();
        _nw244.__ctor();
        _1469_v22 = _nw244;
        let _1470_v23;
        _1470_v23 = _1468_cf0;
        _1470_v23 = _1470_v23;
        let _1471_v24;
        _1471_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1464_v19,p1);
        let _1472_v25;
        _1472_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
        let _1473_v26;
        _1473_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1465_v20)[_module.__default.safeIndex(new BigNumber(383), new BigNumber((_1465_v20).length))]);
        let _1474_v27;
        _1474_v27 = _dafny.Seq.of(_1472_v25);
        _1471_v24 = (_1471_v24).update(_1464_v19, (((((_1472_v25).contains((_this).fm0(_1473_v26, (_this).f22, _1474_v27, globalState))) ? ((_1472_v25).get((_this).fm0(_1473_v26, (_this).f22, _1474_v27, globalState))) : (p1))) ? (!((_this).f21)) : ((_this).f21)));
      }
      let _1475_v28;
      let _nw245 = Array((new BigNumber(26)).toNumber()).fill([]);
      _1475_v28 = _nw245;
      let _1476_v29;
      let _init41 = ((_1477_p0) => function (_1478_i5) {
        return (_1478_i5).plus(_1477_p0);
      })(p0);
      let _nw246 = Array((new BigNumber(19)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw246.length); _i0_41++) {
        _nw246[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1476_v29 = _nw246;
      let _1479_v30;
      let _nw247 = new _module.C1();
      _nw247.__ctor((_this).f20, _1476_v29, (_this).fm1((_this).f20, p0, p1, globalState), false);
      _1479_v30 = _nw247;
      let _1480_v31;
      _1480_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1479_v30,(_this).f21);
      let _1481_v32;
      _1481_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1480_v31,_1475_v28);
      let _1482_v33;
      _1482_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1479_v30).f20);
      let _1483_v34;
      _1483_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(true),true);
      let _1484_v35;
      _1484_v35 = _dafny.Seq.of(_1483_v34);
      let _1485_v36;
      let _nw248 = Array((new BigNumber(13)).toNumber());
      _nw248[(_dafny.ZERO).toNumber()] = _1475_v28;
      _nw248[(_dafny.ONE).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(2)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(3)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(4)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(5)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(6)).toNumber()] = ((p1) ? (_1475_v28) : (_1475_v28));
      _nw248[(new BigNumber(7)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(8)).toNumber()] = (((_1481_v32).contains(_dafny.Map.Empty.slice().updateUnsafe(_1479_v30,(_this).fm0(_1482_v33, (_this).f22, _1484_v35, globalState)))) ? ((_1481_v32).get(_dafny.Map.Empty.slice().updateUnsafe(_1479_v30,(_this).fm0(_1482_v33, (_this).f22, _1484_v35, globalState)))) : (_1475_v28));
      _nw248[(new BigNumber(9)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(10)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(11)).toNumber()] = _1475_v28;
      _nw248[(new BigNumber(12)).toNumber()] = _1475_v28;
      _1485_v36 = _nw248;
      let _index253 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1485_v36).length));
      (_1485_v36)[_index253] = _1475_v28;
      let _1486_v37;
      _1486_v37 = new _dafny.CodePoint('i'.codePointAt(0));
      let _index254 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1485_v36).length));
      let _rhs264 = _1475_v28;
      let _rhs265 = ((false) ? ((new BigNumber((_module.__default.fm12((_1479_v30).f20, _1486_v37, p1, (_this).f20, globalState)).length)).plus((_1479_v30).f20)) : (_module.__default.safeModuloInt((_1479_v30).f20, p0)));
      let _rhs266 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(469)), ((_1487_v37) => function (_1488_i6) {
        return _1487_v37;
      })(_1486_v37)), ((p1) ? (_1486_v37) : (_1486_v37)));
      let _lhs245 = _1485_v36;
      let _lhs246 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1485_v36).length));
      let _lhs247 = globalState;
      let _lhs248 = globalState;
      _lhs245[_lhs246] = _rhs264;
      _lhs247.f17 = _rhs265;
      _lhs248.f4 = _rhs266;
      _1479_v30 = _1479_v30;
      let _hi5 = new BigNumber(585);
      for (let _1489_i7 = (((_1479_v30).f21) ? ((_this).f20) : ((((_1482_v33).contains((_1479_v30).f20)) ? ((_1482_v33).get((_1479_v30).f20)) : ((_this).f20)))); _1489_i7.isLessThan(_hi5); _1489_i7 = _1489_i7.plus(_dafny.ONE)) {
        let _1490_v38;
        _1490_v38 = p1;
        _1490_v38 = _1490_v38;
        let _1491_v39;
        let _nw249 = new _module.C1();
        _nw249.__ctor((_1479_v30).f20, _1476_v29, p0, (_1479_v30).f21);
        _1491_v39 = _nw249;
        _1491_v39 = _1491_v39;
        let _1492_v40;
        _1492_v40 = _dafny.Set.fromElements(new BigNumber((_1483_v34).length), (_this).f20);
        let _1493_v41;
        _1493_v41 = _dafny.Seq.of((_1492_v40).Difference(_1492_v40));
        _1493_v41 = _1493_v41;
        if ((_1479_v30).f21) {
          (_1491_v39).f27 = p0;
          let _1494_v42;
          _1494_v42 = _module.D8.create_DC21(_1491_v39.f27);
          let _1495_v43;
          _1495_v43 = _dafny.MultiSet.fromElements(((_this).f20).isLessThan((_1494_v42).dtor_cf29));
          _1495_v43 = (_1495_v43).Union(_1495_v43);
          let _1496_v44;
          _1496_v44 = _dafny.Seq.of(new BigNumber(((_this).f22).length));
          _1496_v44 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), function (_1497_i8) {
            return (_this).f20;
          }), _dafny.Seq.Concat(_1496_v44, _1496_v44));
          (globalState).f15 = (_this).f22;
          _1476_v29 = (_1491_v39).f28;
        } else {
          let _1498_v45;
          _1498_v45 = _dafny.Seq.of(_1491_v39.f27, (_1479_v30).f20);
          let _1499_v47;
          _1499_v47 = _dafny.Set.fromElements(_1492_v40, _1492_v40, _1492_v40, function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_dafny.MultiSet.FromArray(_1498_v45)).Elements) {
              let _1500_v46 = _compr_48;
              if ((_dafny.MultiSet.FromArray(_1498_v45)).contains(_1500_v46)) {
                _coll48.add((_1500_v46).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality())));
              }
            }
            return _coll48;
          }(), _1492_v40);
          let _rhs267 = (_1479_v30).f20;
          let _rhs268 = (_this).f21;
          let _rhs269 = (_1499_v47).Difference(_1499_v47);
          let _lhs249 = globalState;
          let _lhs250 = globalState;
          _lhs249.f17 = _rhs267;
          _lhs250.f4 = _rhs268;
          _1499_v47 = _rhs269;
          (globalState).f5 = p0;
          let _1501_v48;
          _1501_v48 = _module.D13.create_DC37((_1479_v30).f21, (_this).f20);
          (globalState).f17 = ((_1479_v30).f20).plus((_1501_v48).dtor_cf50);
          let _1502_v49;
          let _nw250 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1502_v49 = _nw250;
          let _1503_v50;
          _1503_v50 = _dafny.Set.fromElements((_this).f21);
          let _1504_v51;
          _1504_v51 = _dafny.Seq.of(_1503_v50);
          let _1505_v52;
          _1505_v52 = _module.D3.create_DC7(_1504_v51, p1);
          let _index255 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1502_v49).length));
          (_1502_v49)[_index255] = (_1505_v52).dtor_cf12;
          let _index256 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1502_v49).length));
          (_1502_v49)[_index256] = p1;
          let _1506_v53;
          _1506_v53 = _1476_v29;
          let _1507_v54;
          _1507_v54 = _dafny.Set.fromElements(_1506_v53);
          let _1508_v55;
          _1508_v55 = _dafny.Seq.of((_1479_v30).f21);
          let _1509_v56;
          _1509_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1508_v55);
          let _1510_v57;
          _1510_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1507_v54,(_dafny.ZERO).minus(new BigNumber(((((_1509_v56).contains((_1479_v30).f21)) ? ((_1509_v56).get((_1479_v30).f21)) : (_1508_v55))).length)));
          _1510_v57 = _1510_v57;
        }
      }
      let _1511_v58;
      _1511_v58 = _dafny.Set.fromElements(_1476_v29);
      let _hi6 = new BigNumber((_1511_v58).length);
      for (let _1512_i9 = _module.__default.safeModuloInt((_this).f20, p0); _1512_i9.isLessThan(_hi6); _1512_i9 = _1512_i9.plus(_dafny.ONE)) {
        let _1513_v59;
        _1513_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f21);
        let _1514_v60;
        _1514_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_1513_v59);
        let _1515_v61;
        _1515_v61 = _module.D16.create_DC42(_module.__default.fm40((_this).f21, (_1479_v30).f21, (_1479_v30).f20, new BigNumber((_1514_v60).length), globalState));
        (globalState).f17 = new BigNumber(((_1515_v61).dtor_cf54).length);
        (globalState).f15 = _dafny.Seq.Concat((_this).f22, (_this).f22);
        _1483_v34 = _1483_v34;
        if (!((_dafny.Seq.contains((_this).f22, _1486_v37)) === (!((_1479_v30).f21)))) {
          (globalState).f17 = (_this).f20;
          (globalState).f3 = true;
          (globalState).f4 = (_this).fm0(_1482_v33, _dafny.Seq.Concat((_this).f22, (_this).f22), _1484_v35, globalState);
          let _1516_v62;
          let _nw251 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _1516_v62 = _nw251;
          let _1517_v63;
          _1517_v63 = _dafny.Seq.of((_1479_v30).f21);
          let _index257 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1516_v62).length));
          (_1516_v62)[_index257] = _1517_v63;
          let _index258 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1516_v62).length));
          let _rhs270 = _dafny.Seq.Concat(_1517_v63, _dafny.Seq.of(p1, (_this).f21, (_1479_v30).f21, (_1479_v30).f21));
          let _rhs271 = (((_1482_v33).contains(new BigNumber((_dafny.Seq.UnicodeFromString("eopxabkhd")).length))) ? ((_1482_v33).get(new BigNumber((_dafny.Seq.UnicodeFromString("eopxabkhd")).length))) : (((p1) ? (new BigNumber(255)) : ((_1479_v30).f20))));
          let _lhs251 = _1516_v62;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1516_v62).length));
          let _lhs253 = globalState;
          _lhs251[_lhs252] = _rhs270;
          _lhs253.f5 = _rhs271;
          let _1518_v64;
          let _nw252 = Array((new BigNumber(19)).toNumber());
          _nw252[(_dafny.ZERO).toNumber()] = (_this).f22;
          _nw252[(_dafny.ONE).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(2)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(3)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(4)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("diwso");
          _nw252[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("xkkwmspk");
          _nw252[(new BigNumber(7)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(8)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_1519_v37) => function (_1520_i10) {
            return _1519_v37;
          })(_1486_v37));
          _nw252[(new BigNumber(10)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(11)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-284)), ((_1521_v37) => function (_1522_i11) {
            return _1521_v37;
          })(_1486_v37));
          _nw252[(new BigNumber(13)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(14)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(15)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(16)).toNumber()] = (_this).f22;
          _nw252[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("ikkiep");
          _nw252[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("yvy");
          _1518_v64 = _nw252;
          let _1523_v65;
          _1523_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.D11.create_DC28(_1518_v64),(_1479_v30).f21);
          let _1524_v66;
          _1524_v66 = _dafny.Seq.of(_1523_v65);
          let _1525_v67;
          _1525_v67 = _module.D2.create_DC3((_1479_v30).f20, (_this).f21, (_1479_v30).f21, (_1479_v30).f20, (_1479_v30).f21);
          let _1526_v68;
          _1526_v68 = _dafny.Seq.of(_1517_v63);
          let _1527_v69;
          let _init42 = ((_1528_p1) => function (_1529_i12) {
            return _1528_p1;
          })(p1);
          let _nw253 = Array((new BigNumber(3)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw253.length); _i0_42++) {
            _nw253[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1527_v69 = _nw253;
          let _1530_v70;
          _1530_v70 = _dafny.Set.fromElements((_1479_v30).f20, (_1479_v30).f20);
          let _1531_v71;
          _1531_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1527_v69,_1530_v70);
          let _1532_v72;
          _1532_v72 = _dafny.Set.fromElements(_dafny.Seq.of((((_1483_v34).contains((_1479_v30).f21)) ? ((_1483_v34).get((_1479_v30).f21)) : (!((_1525_v67).dtor_cf4)))), (_1526_v68)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1517_v63, _module.__default.safeIndex(new BigNumber((_1531_v71).length), new BigNumber((_1517_v63).length)), (_1479_v30).f21)).length), new BigNumber((_1526_v68).length))]);
          let _1533_v73;
          _1533_v73 = _dafny.MultiSet.fromElements(!((_1479_v30).f21));
          let _1534_v74;
          _1534_v74 = _module.D4.create_DC10(_1533_v73, (_this).f21);
          let _1535_v75;
          _1535_v75 = _dafny.Seq.of((_1479_v30).f20);
          let _rhs272 = new BigNumber(-116);
          let _rhs273 = (_1479_v30).fm1((_dafny.ZERO).minus(p0), ((_1479_v30).f20).multipliedBy((_1479_v30).f20), (_1534_v74).dtor_cf16, globalState);
          let _rhs274 = _1524_v66;
          let _rhs275 = _1532_v72;
          let _rhs276 = _module.__default.safeDivisionInt((_1479_v30).f20, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1535_v75, _module.__default.safeIndex((_this).f20, new BigNumber((_1535_v75).length)), _1512_i9), _dafny.Seq.of((_this).f20))).length));
          let _lhs254 = globalState;
          let _lhs255 = globalState;
          let _lhs256 = globalState;
          _lhs254.f17 = _rhs272;
          _lhs255.f17 = _rhs273;
          _1524_v66 = _rhs274;
          _1532_v72 = _rhs275;
          _lhs256.f17 = _rhs276;
        } else {
          _1513_v59 = (_1513_v59).update(p0, true);
          (globalState).f18 = !_dafny.Seq.contains(_dafny.Seq.Concat((_this).f22, (_this).f22), _1486_v37);
          let _1536_v76;
          let _nw254 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1536_v76 = _nw254;
          let _1537_v77;
          _1537_v77 = _module.D11.create_DC28(_1536_v76);
          _1537_v77 = _module.D11.create_DC28(_1536_v76);
          let _1538_v78;
          let _nw255 = Array((new BigNumber(22)).toNumber()).fill([]);
          _1538_v78 = _nw255;
          let _1539_v79;
          _1539_v79 = _dafny.MultiSet.fromElements((_1479_v30).f21, (_1479_v30).f21);
          let _1540_v81;
          _1540_v81 = _module.D11.create_DC30(false, _1539_v79);
          let _1541_v82;
          _1541_v82 = _dafny.Map.Empty.slice().updateUnsafe((_1479_v30).f20,_1540_v81);
          let _1542_v83;
          let _nw256 = Array((new BigNumber(13)).toNumber());
          _nw256[(_dafny.ZERO).toNumber()] = (_this).f21;
          _nw256[(_dafny.ONE).toNumber()] = (_1479_v30).f21;
          _nw256[(new BigNumber(2)).toNumber()] = p1;
          _nw256[(new BigNumber(3)).toNumber()] = (_this).fm0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1539_v79).cardinality()),(_1479_v30).f20), (_this).f22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_1543_v34) => function (_1544_i13) {
            return _1543_v34;
          })(_1483_v34)), globalState);
          _nw256[(new BigNumber(4)).toNumber()] = (_1479_v30).f21;
          _nw256[(new BigNumber(5)).toNumber()] = (_this).fm0(_1482_v33, (_this).f22, _1484_v35, globalState);
          _nw256[(new BigNumber(6)).toNumber()] = (_this).f21;
          _nw256[(new BigNumber(7)).toNumber()] = (((_1513_v59).contains(new BigNumber(((_this).f22).length))) ? ((_1513_v59).get(new BigNumber(((_this).f22).length))) : ((_1479_v30).f21));
          _nw256[(new BigNumber(8)).toNumber()] = (_1479_v30).f21;
          _nw256[(new BigNumber(9)).toNumber()] = (_this).f21;
          _nw256[(new BigNumber(10)).toNumber()] = (_1479_v30).f21;
          _nw256[(new BigNumber(11)).toNumber()] = true;
          _nw256[(new BigNumber(12)).toNumber()] = (_this).fm0(function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_1541_v82).Keys.Elements) {
              let _1545_v80 = _compr_49;
              if ((_1541_v82).contains(_1545_v80)) {
                _coll49.push([_module.__default.safeModuloInt(_1545_v80, (_1479_v30).f20),p0]);
              }
            }
            return _coll49;
          }(), (_this).f22, _1484_v35, globalState);
          _1542_v83 = _nw256;
          let _index259 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1538_v78).length));
          (_1538_v78)[_index259] = _1542_v83;
          let _index260 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1538_v78).length));
          (_1538_v78)[_index260] = _1542_v83;
          (globalState).f5 = _module.__default.safeDivisionInt(new BigNumber(787), (_this).f20);
        }
      }
      let _1546_v84;
      _1546_v84 = _dafny.Seq.of(_1476_v29);
      let _1547_v85;
      _1547_v85 = _module.D17.create_DC46(_1546_v84);
      r0 = _dafny.Seq.Concat((_1547_v85).dtor_cf62, (_1547_v85).dtor_cf62);
      let _1548_v86;
      _1548_v86 = _module.D6.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-363)), ((_1549_v37) => function (_1550_i14) {
  return _1549_v37;
})(_1486_v37)));
      r1 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("jhqis"), _dafny.Seq.Concat((_1548_v86).dtor_cf21, (_this).f22));
      return [r0, r1];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
