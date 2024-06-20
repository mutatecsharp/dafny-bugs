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
      return (new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(985), new BigNumber(907))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(985)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(907)))) {
            _coll0.push([(_0_v0).minus(new BigNumber(-431)),new _dafny.CodePoint('k'.codePointAt(0))]);
          }
        }
        return _coll0;
      }()).length)).minus((_dafny.ZERO).minus((new BigNumber(-875)).plus(new BigNumber(32))));
    };
    static fm3(p0, globalState) {
      return (_dafny.Seq.IsPrefixOf(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-909)), new BigNumber((_dafny.Seq.of(false, true)).length), new BigNumber(-14), new BigNumber(948), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))), _dafny.Seq.of(new BigNumber(502), new BigNumber(581), new BigNumber(-87), new BigNumber(242)))) && ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length))).length)).isEqualTo(new BigNumber(-105)));
    };
    static fm5(globalState) {
      return new _dafny.CodePoint('t'.codePointAt(0));
    };
    static fm13(p0, globalState) {
      let _source0 = _module.D7.create_DC16(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), function (_1_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
})).length));
      if (_source0.is_DC16) {
        let _2___mcc_h0 = (_source0).cf29;
        let _3_cf29 = _2___mcc_h0;
        if (true) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }
      } else if (_source0.is_DC15) {
        let _4___mcc_h1 = (_source0).cf28;
        let _5_cf28 = _4___mcc_h1;
        return new _dafny.CodePoint('i'.codePointAt(0));
      } else {
        let _6___mcc_h2 = (_source0).cf30;
        let _7_cf30 = _6___mcc_h2;
        return new _dafny.CodePoint('q'.codePointAt(0));
      }
    };
    static fm17(globalState) {
      return true;
    };
    static fm18(p0, p1, p2, p3, globalState) {
      if (!(!((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("og"))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sple")))))) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }
    };
    static fm19(globalState) {
      return ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), function (_8_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length))), (_dafny.ZERO).minus(new BigNumber(-377)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(218), new BigNumber((_dafny.Seq.UnicodeFromString("aqrya")).length))).length))))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(-649), new BigNumber((_dafny.Seq.UnicodeFromString("mbapixuot")).length), new BigNumber(394))).length), new BigNumber(923), new BigNumber(695)))));
    };
    static fm20(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ofvstaor"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rfjmdefgm"), _dafny.Seq.UnicodeFromString("r")));
    };
    static fm21(globalState) {
      return _module.D1.create_DC1(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _dafny.Seq.UnicodeFromString("sflqu")));
    };
    static fm22(p0, p1, globalState) {
      let _source1 = ((false) ? (_module.D6.create_DC11(_dafny.MultiSet.fromElements(new BigNumber(344)))) : (_module.D6.create_DC11(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-70))))));
      if (_source1.is_DC12) {
        let _9___mcc_h0 = (_source1).cf18;
        let _10___mcc_h1 = (_source1).cf19;
        let _11___mcc_h2 = (_source1).cf20;
        let _12___mcc_h3 = (_source1).cf21;
        let _13_cf21 = _12___mcc_h3;
        let _14_cf20 = _11___mcc_h2;
        let _15_cf19 = _10___mcc_h1;
        let _16_cf18 = _9___mcc_h0;
        return _dafny.Set.fromElements(_15_cf19);
      } else if (_source1.is_DC13) {
        let _17___mcc_h4 = (_source1).cf22;
        let _18___mcc_h5 = (_source1).cf23;
        let _19___mcc_h6 = (_source1).cf24;
        let _20___mcc_h7 = (_source1).cf25;
        let _21___mcc_h8 = (_source1).cf26;
        let _22_cf26 = _21___mcc_h8;
        let _23_cf25 = _20___mcc_h7;
        let _24_cf24 = _19___mcc_h6;
        let _25_cf23 = _18___mcc_h5;
        let _26_cf22 = _17___mcc_h4;
        return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("yyjhaqfxv"));
      } else if (_source1.is_DC11) {
        let _27___mcc_h9 = (_source1).cf17;
        let _28_cf17 = _27___mcc_h9;
        return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ycub"), _dafny.Seq.UnicodeFromString("dvtlsfkq"));
      } else {
        let _29___mcc_h10 = (_source1).cf27;
        let _30_cf27 = _29___mcc_h10;
        return function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vecee"),true)).Keys.Elements) {
            let _31_v0 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vecee"),true)).contains(_31_v0)) {
              _coll1.add(_31_v0);
            }
          }
          return _coll1;
        }();
      }
    };
    static fm23(p0, globalState) {
      return ((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(((_module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length)), _dafny.Set.fromElements(true, false), _dafny.MultiSet.fromElements(new BigNumber(296)), _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("lqbbxmxcd")))).dtor_cf34).cardinality())))).Elements) {
          let _32_v0 = _compr_2;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(((_module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length)), _dafny.Set.fromElements(true, false), _dafny.MultiSet.fromElements(new BigNumber(296)), _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("lqbbxmxcd")))).dtor_cf34).cardinality())))).contains(_32_v0)) {
            _coll2.push([_module.__default.safeModuloInt(_32_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(576), new BigNumber(33))).cardinality())),new _dafny.CodePoint('f'.codePointAt(0))]);
          }
        }
        return _coll2;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-86),new _dafny.CodePoint('f'.codePointAt(0))))).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(137),new _dafny.CodePoint('n'.codePointAt(0)))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-424))).length),new _dafny.CodePoint('w'.codePointAt(0))))));
    };
    static fm24(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(137),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("gfpmde")).length),true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), function (_33_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length),new BigNumber((_dafny.Seq.UnicodeFromString("qtnakf")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bug")).length)),new BigNumber(798)));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-151)), function (_34_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("puihdeaw")).length),!(!(false)));
      });
    };
    static fm26(p0, p1, p2, globalState) {
      return ((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), function (_35_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }))).Elements) {
          let _36_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), function (_35_i0) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })), _36_v0)) {
            _coll3.push([_36_v0,!(true)]);
          }
        }
        return _coll3;
      }()).Merge(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tio"),_dafny.Set.fromElements(true))).Keys.Elements) {
          let _37_v1 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tio"),_dafny.Set.fromElements(true))).contains(_37_v1)) {
            _coll4.push([_37_v1,false]);
          }
        }
        return _coll4;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe((_module.D9.create_DC22(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()),true)).length), !(true), _dafny.Set.fromElements(!(true)), _dafny.Seq.UnicodeFromString("gc"))).dtor_cf41,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vgp"),!(false))));
    };
    static fm27(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), function (_38_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber(-901)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-458)), function (_39_i1) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(((_module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(104)), _dafny.Set.fromElements(false), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(445))).length), new BigNumber((_dafny.Seq.UnicodeFromString("dixgeiij")).length)), _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("lqypklsa")))).dtor_cf32).length))).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true, true)).length)))).length));
      }));
    };
    static fm28(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(false))).Intersect((_dafny.MultiSet.fromElements(false, false, true)).Union(_dafny.MultiSet.fromElements((_module.D6.create_DC12(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(186)), function (_40_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
}), new BigNumber(574), true)).dtor_cf18, true, false, false, true)));
    };
    static fm29(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(false)));
    };
    static fm30(p0, globalState) {
      return false;
    };
    static fm31(globalState) {
      return function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-335)), function (_41_i0) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        })), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("aavabqni")))).Elements) {
          let _42_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-335)), function (_41_i0) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("aavabqni"))), _42_v0)) {
            _coll5.push([_42_v0,_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('n'.codePointAt(0)))]);
          }
        }
        return _coll5;
      }();
    };
    static fm32(p0, globalState) {
      let _source2 = _module.D16.create_DC41(true);
      if (_source2.is_DC41) {
        let _43___mcc_h0 = (_source2).cf84;
        let _44_cf84 = _43___mcc_h0;
        return _dafny.Seq.UnicodeFromString("sqsbuy");
      } else if (_source2.is_DC42) {
        let _45___mcc_h1 = (_source2).cf85;
        let _46_cf85 = _45___mcc_h1;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(660)), function (_47_i0) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        });
      } else if (_source2.is_DC40) {
        let _48___mcc_h2 = (_source2).cf83;
        let _49_cf83 = _48___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sidmalhfq"), _dafny.Seq.UnicodeFromString("nga"));
      } else {
        let _50___mcc_h3 = (_source2).cf86;
        let _51_cf86 = _50___mcc_h3;
        return _dafny.Seq.UnicodeFromString("bnryd");
      }
    };
    static fm33(p0, p1, globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm34(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("pdggpy")).length)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(859)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-621)), function (_52_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-560)), function (_53_i1) {
          return new BigNumber(497);
        })).length);
      })));
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((true) === (false),_dafny.Seq.UnicodeFromString("coubrsv"));
    };
    static fm36(p0, p1, p2, globalState) {
      if ((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(789), new BigNumber(442))) {
          let _54_v1 = _compr_6;
          if (((new BigNumber(789)).isLessThanOrEqualTo(_54_v1)) && ((_54_v1).isLessThan(new BigNumber(442)))) {
            _coll6.push([(_54_v1).multipliedBy(new BigNumber(-848)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))).length), new BigNumber(12)))).cardinality())]);
          }
        }
        return _coll6;
      }()).length))).IsSubsetOf(function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hvfvypqh")).length),_dafny.Seq.UnicodeFromString("weervqcj"))).Keys.Elements) {
          let _55_v0 = _compr_7;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hvfvypqh")).length),_dafny.Seq.UnicodeFromString("weervqcj"))).contains(_55_v0)) {
            _coll7.add(_module.__default.safeModuloInt(_55_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lhu")).length))).cardinality())));
          }
        }
        return _coll7;
      }())) {
        return (_dafny.Set.fromElements(true, !(false), false)).Union(_dafny.Set.fromElements(true));
      } else {
        return _dafny.Set.fromElements(true);
      }
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC7(),_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ti")).length), new BigNumber(332)))).length), (new BigNumber(-570)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("wmlctwlx")).length)), false);
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(201), new BigNumber(-244))) {
          let _56_v0 = _compr_8;
          if (((new BigNumber(201)).isLessThanOrEqualTo(_56_v0)) && ((_56_v0).isLessThan(new BigNumber(-244)))) {
            _coll8.add((_56_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-467), new BigNumber(716))).length))));
          }
        }
        return _coll8;
      }()).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(289), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-587),new _dafny.CodePoint('q'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(969)), function (_57_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length), new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(371), new BigNumber(582))) {
          let _58_v1 = _compr_9;
          if (((new BigNumber(371)).isLessThanOrEqualTo(_58_v1)) && ((_58_v1).isLessThan(new BigNumber(582)))) {
            _coll9.add((_58_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(704))).length)));
          }
        }
        return _coll9;
      }()).length))).length),false)).length))));
    };
    static fm39(p0, p1, globalState) {
      return _module.D6.create_DC13((new BigNumber(-93)).isEqualTo(new BigNumber(37)), false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), function (_59_i0) {
  return new BigNumber(-741);
})).length)), new BigNumber(232), !(!(true) || (!(false))));
    };
    static fm40(globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))).Union(_dafny.MultiSet.fromElements(true, !(false)));
    };
    static fm41(p0, p1, globalState) {
      return ((_module.D21.create_DC54(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fdkkibdc")).length))))).dtor_cf105).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(958))).length), new BigNumber((_dafny.Seq.UnicodeFromString("dg")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(new BigNumber((function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(272), new BigNumber(733))) {
          let _60_v0 = _compr_10;
          if (((new BigNumber(272)).isLessThanOrEqualTo(_60_v0)) && ((_60_v0).isLessThan(new BigNumber(733)))) {
            _coll10.add(_module.__default.safeDivisionInt(_60_v0, new BigNumber(509)));
          }
        }
        return _coll10;
      }()).length), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("win"))).length)))));
    };
    static fm44(p0, globalState) {
      let _source3 = ((true) ? (_module.D21.create_DC56(true)) : (_module.D21.create_DC56(false)));
      if (_source3.is_DC55) {
        let _61___mcc_h0 = (_source3).cf106;
        let _62___mcc_h1 = (_source3).cf107;
        let _63___mcc_h2 = (_source3).cf108;
        let _64_cf108 = _63___mcc_h2;
        let _65_cf107 = _62___mcc_h1;
        let _66_cf106 = _61___mcc_h0;
        return (function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_66_cf106),_66_cf106)).Keys.Elements) {
              let _67_v1 = _compr_12;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_66_cf106),_66_cf106)).contains(_67_v1)) {
                _coll12.add(_67_v1);
              }
            }
            return _coll12;
          }()).Elements) {
            let _68_v0 = _compr_11;
            if ((function () {
              let _coll13 = new _dafny.Set();
              for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_66_cf106),_66_cf106)).Keys.Elements) {
                let _69_v1 = _compr_13;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_66_cf106),_66_cf106)).contains(_69_v1)) {
                  _coll13.add(_69_v1);
                }
              }
              return _coll13;
            }()).contains(_68_v0)) {
              _coll11.push([_68_v0,(_dafny.ZERO).minus(_64_cf108)]);
            }
          }
          return _coll11;
        }()).Merge((_module.D22.create_DC58(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of (_dafny.Seq.of(_dafny.Set.fromElements(!(true)))).Elements) {
    let _70_v2 = _compr_14;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(!(true))), _70_v2)) {
      _coll14.push([_70_v2,_64_cf108]);
    }
  }
  return _coll14;
}())).dtor_cf112);
      } else if (_source3.is_DC56) {
        let _71___mcc_h3 = (_source3).cf109;
        let _72_cf109 = _71___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_72_cf109, _72_cf109),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()));
      } else if (_source3.is_DC57) {
        let _73___mcc_h4 = (_source3).cf110;
        let _74___mcc_h5 = (_source3).cf111;
        let _75_cf111 = _74___mcc_h5;
        let _76_cf110 = _73___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_76_cf110),new BigNumber((_dafny.MultiSet.fromElements(_76_cf110)).cardinality()));
      } else {
        let _77___mcc_h6 = (_source3).cf105;
        let _78_cf105 = _77___mcc_h6;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),new BigNumber(345));
      }
    };
    static fm45(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("k"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(483)), function (_79_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(797)), function (_80_i1) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("wenrpms")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("shmvcrb")));
    };
    static fm46(p0, p1, p2, globalState) {
      return _module.D9.create_DC22(((false) ? (false) : (true)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-266)), _dafny.Seq.of(new BigNumber(997), new BigNumber(737)))).length)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.UnicodeFromString("uxklvq")).length))).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(264)), function (_81_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})).length)), _dafny.Set.fromElements(false, false), _dafny.Seq.UnicodeFromString("vot"));
    };
    static fm47(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(167), (_dafny.ZERO).minus(new BigNumber(-372)))).Intersect((_dafny.Set.fromElements(new BigNumber(-541))).Difference(function () {
        let _coll15 = new _dafny.Set();
        for (const _compr_15 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-279)), function (_82_i0) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })).length)))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_83_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length))).Elements) {
          let _84_v0 = _compr_15;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-279)), function (_82_i0) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length)))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_83_i1) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length))).contains(_84_v0)) {
            _coll15.add((_84_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("thfmph")).length)));
          }
        }
        return _coll15;
      }()));
    };
    static fm48(globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), function (_85_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("utp")))).Intersect((function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("i")))).Elements) {
          let _86_v0 = _compr_16;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("i")))).contains(_86_v0)) {
            _coll16.add(_86_v0);
          }
        }
        return _coll16;
      }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ogjtqc"), _dafny.Seq.UnicodeFromString("hyaujykjq"))));
    };
    static fm49(p0, p1, globalState) {
      return _module.D3.create_DC8(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("idqyyo")).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())))).cardinality()));
    };
    static fm50(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(956))).cardinality())).plus(new BigNumber(415)),!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("v"), new _dafny.CodePoint('s'.codePointAt(0))));
    };
    static fm51(globalState) {
      return _module.D9.create_DC23(((true) ? (new BigNumber(416)) : (new BigNumber(-787))), _module.__default.safeDivisionInt(new BigNumber(371), new BigNumber(820)), new BigNumber(((_dafny.MultiSet.fromElements(true, false, !(true))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, !(!(true)), true, false)))).cardinality()), (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-494)), function (_87_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})).length)).multipliedBy(new BigNumber(839)), new BigNumber(109));
    };
    static fm54(p0, globalState) {
      return ((_dafny.Set.fromElements(!(!(false)))).Intersect(_dafny.Set.fromElements(false, true))).Difference(_dafny.Set.fromElements(true));
    };
    static fm55(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(true);
    };
    static fm56(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bdebqr")).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("hixovsg")).length))).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_88_i0) {
        return new BigNumber(-869);
      })).length), new BigNumber((_dafny.MultiSet.fromElements(false, false, false)).cardinality())))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(941)));
    };
    static fm57(globalState) {
      let _source4 = _module.D21.create_DC56(true);
      if (_source4.is_DC55) {
        let _89___mcc_h0 = (_source4).cf106;
        let _90___mcc_h1 = (_source4).cf107;
        let _91___mcc_h2 = (_source4).cf108;
        let _92_cf108 = _91___mcc_h2;
        let _93_cf107 = _90___mcc_h1;
        let _94_cf106 = _89___mcc_h0;
        return _module.D12.create_DC32(_94_cf106, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length), false, _94_cf106, _92_cf108);
      } else if (_source4.is_DC56) {
        let _95___mcc_h3 = (_source4).cf109;
        let _96_cf109 = _95___mcc_h3;
        return _module.D12.create_DC32(_96_cf109, new BigNumber(649), _96_cf109, _96_cf109, new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_96_cf109,_96_cf109), _dafny.Map.Empty.slice().updateUnsafe(_96_cf109,_96_cf109), _dafny.Map.Empty.slice().updateUnsafe(_96_cf109,_96_cf109), _dafny.Map.Empty.slice().updateUnsafe(_96_cf109,_96_cf109), _dafny.Map.Empty.slice().updateUnsafe(_96_cf109,_96_cf109))).cardinality()));
      } else if (_source4.is_DC57) {
        let _97___mcc_h4 = (_source4).cf110;
        let _98___mcc_h5 = (_source4).cf111;
        let _99_cf111 = _98___mcc_h5;
        let _100_cf110 = _97___mcc_h4;
        return _module.D12.create_DC32(!(_100_cf110), new BigNumber(((_module.D22.create_DC60(new BigNumber(-656), _dafny.Seq.of(false, !(_100_cf110)), false)).dtor_cf116).length), _100_cf110, _100_cf110, new BigNumber(559));
      } else {
        let _101___mcc_h6 = (_source4).cf105;
        let _102_cf105 = _101___mcc_h6;
        return _module.D12.create_DC32(false, new BigNumber(-50), false, true, new BigNumber((_dafny.Set.fromElements(new BigNumber(698), new BigNumber(259))).length));
      }
    };
    static fm58(globalState) {
      if ((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-639)), function (_103_i0) {
        return _module.D22.create_DC59(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), function (_104_i1) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length), false);
      }),_dafny.Seq.of(new BigNumber(668)))).length))).cardinality())).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber(325))).length))) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      } else if (!(false)) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }
    };
    static fm59(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),true);
    };
    static fm60(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(182)), function (_105_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      });
    };
    static fm61(p0, p1, p2, globalState) {
      let _source5 = _module.D6.create_DC11(_dafny.MultiSet.fromElements(new BigNumber(-259)));
      if (_source5.is_DC12) {
        let _106___mcc_h0 = (_source5).cf18;
        let _107___mcc_h1 = (_source5).cf19;
        let _108___mcc_h2 = (_source5).cf20;
        let _109___mcc_h3 = (_source5).cf21;
        let _110_cf21 = _109___mcc_h3;
        let _111_cf20 = _108___mcc_h2;
        let _112_cf19 = _107___mcc_h1;
        let _113_cf18 = _106___mcc_h0;
        return _dafny.Seq.of(!(_110_cf21), true, _110_cf21, false, false);
      } else if (_source5.is_DC13) {
        let _114___mcc_h4 = (_source5).cf22;
        let _115___mcc_h5 = (_source5).cf23;
        let _116___mcc_h6 = (_source5).cf24;
        let _117___mcc_h7 = (_source5).cf25;
        let _118___mcc_h8 = (_source5).cf26;
        let _119_cf26 = _118___mcc_h8;
        let _120_cf25 = _117___mcc_h7;
        let _121_cf24 = _116___mcc_h6;
        let _122_cf23 = _115___mcc_h5;
        let _123_cf22 = _114___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.of(_123_cf22), _dafny.Seq.of(_119_cf26));
      } else if (_source5.is_DC11) {
        let _124___mcc_h9 = (_source5).cf17;
        let _125_cf17 = _124___mcc_h9;
        return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true));
      } else {
        let _126___mcc_h10 = (_source5).cf27;
        let _127_cf27 = _126___mcc_h10;
        return _dafny.Seq.of(!(true));
      }
    };
    static fm62(p0, p1, globalState) {
      return _module.D1.create_DC1(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cb"), _dafny.Seq.UnicodeFromString("rjfj")));
    };
    static fm63(p0, p1, globalState) {
      if (!(false)) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), function (_128_i0) {
          return new BigNumber(489);
        }), _dafny.Seq.of(new BigNumber(31), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(449),true)).length))));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_129_i1) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("gqn")).length);
        }), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())));
      }
    };
    static fm64(globalState) {
      return (_dafny.Set.fromElements(new BigNumber(787))).Intersect((_dafny.Set.fromElements(new BigNumber(350), new BigNumber(-91), new BigNumber((_dafny.Seq.UnicodeFromString("w")).length))).Intersect((_module.D19.create_DC49(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_130_i0) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).cardinality())))).dtor_cf93));
    };
    static fm65(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(new BigNumber(((_module.D19.create_DC50(true, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(107))).cardinality()), _dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_131_i0) {
  return new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((function () {
    let _coll17 = new _dafny.Set();
    for (const _compr_17 of (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(649))).length))).Elements) {
      let _132_v0 = _compr_17;
      if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(649))).length)), _132_v0)) {
        _coll17.add((_132_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("xfi")).length)));
      }
    }
    return _coll17;
  }()).length)), new BigNumber((_dafny.Set.fromElements(new BigNumber(74))).length))).cardinality());
}))).dtor_cf96).length), new BigNumber(719), new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(false, true), _dafny.Seq.of(true))).length)), (_module.D11.create_DC27(true, new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("ycrugglv")).length), new BigNumber((_dafny.Seq.UnicodeFromString("nokerh")).length), _dafny.Seq.of(new BigNumber(453)))).dtor_cf55, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ggmughh")).length), new BigNumber((_dafny.Seq.of((_module.D3.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()), false)).dtor_cf13, true)).length), new BigNumber((function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(210), new BigNumber(31))) {
          let _133_v1 = _compr_18;
          if (((new BigNumber(210)).isLessThanOrEqualTo(_133_v1)) && ((_133_v1).isLessThan(new BigNumber(31)))) {
            _coll18.add((_133_v1).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jrgri")).length)));
          }
        }
        return _coll18;
      }()).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))), new BigNumber((_dafny.Set.fromElements(true, true, false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false)).length))), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(293)))).Elements) {
            let _134_v2 = _compr_20;
            if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(293)))).contains(_134_v2)) {
              _coll20.push([_134_v2,true]);
            }
          }
          return _coll20;
        }()).Keys.Elements) {
          let _135_v3 = _compr_19;
          if ((function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(293)))).Elements) {
              let _134_v2 = _compr_21;
              if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(293)))).contains(_134_v2)) {
                _coll21.push([_134_v2,true]);
              }
            }
            return _coll21;
          }()).contains(_135_v3)) {
            _coll19.add(_135_v3);
          }
        }
        return _coll19;
      }()).length), new BigNumber(-985), (_dafny.ZERO).minus(new BigNumber(-569)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), function (_136_i1) {
        return new BigNumber(106);
      })).length), new BigNumber((_dafny.Set.fromElements(true)).length))).cardinality())), ((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-157)), function (_137_i2) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('m'.codePointAt(0)))).length))).length))).length);
      })) : (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length), new BigNumber((function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("glco"),new BigNumber(914))).Keys.Elements) {
          let _138_v4 = _compr_22;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("glco"),new BigNumber(914))).contains(_138_v4)) {
            _coll22.push([_138_v4,false]);
          }
        }
        return _coll22;
      }()).length), new BigNumber((_dafny.Seq.of(new BigNumber(751))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vqoyap")).length))).length)))));
    };
    static fm66(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nnwdhunb"), _dafny.Seq.UnicodeFromString("dpcfmoy"))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ehco"), _dafny.Seq.UnicodeFromString("eo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), function (_139_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(487)), function (_140_i1) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("bj")))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("m"), _dafny.Seq.UnicodeFromString("bs"), _dafny.Seq.UnicodeFromString("rlwyon")));
    };
    static fm67(p0, globalState) {
      return _module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(693)), function (_141_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length))), (_dafny.Set.fromElements(true, true)).Difference(_dafny.Set.fromElements(false, false)), (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-601)), new BigNumber(-799))).Union(_dafny.MultiSet.fromElements(new BigNumber(-453))), _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("n")));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.MultiSet.Empty;
      let _142_v0;
      let _nw0 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _142_v0 = _nw0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_142_v0).length))) {
        let _143_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_143_i0)) && ((_143_i0).isLessThan(new BigNumber((_142_v0).length))))) {
          (_142_v0)[(_143_i0)] = (_143_i0).plus(new BigNumber(605));
        }
      }
      let _144_i1;
      _144_i1 = _dafny.ZERO;
      L0: {
        while ((p1) || (p1)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_144_i1)) {
              break L0;
            }
            _144_i1 = (_144_i1).plus(_dafny.ONE);
            (globalState).f7 = new BigNumber(-8);
            _142_v0 = _142_v0;
            let _145_v1;
            _145_v1 = new BigNumber(859);
            let _146_v2;
            _146_v2 = _dafny.Seq.of(_145_v1);
            let _147_v3;
            _147_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,_146_v2);
            let _148_v4;
            let _nw1 = Array((new BigNumber(2)).toNumber());
            _nw1[(_dafny.ZERO).toNumber()] = _147_v3;
            _nw1[(_dafny.ONE).toNumber()] = _147_v3;
            _148_v4 = _nw1;
            let _index0 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_148_v4).length));
            (_148_v4)[_index0] = _147_v3;
            let _index1 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_148_v4).length));
            (_148_v4)[_index1] = _147_v3;
            let _149_v5;
            _149_v5 = _dafny.Map.Empty.slice().updateUnsafe(_145_v1,p1);
            let _150_v6;
            let _nw2 = Array((new BigNumber(8)).toNumber());
            _nw2[(_dafny.ZERO).toNumber()] = p1;
            _nw2[(_dafny.ONE).toNumber()] = (((_149_v5).contains(_145_v1)) ? ((_149_v5).get(_145_v1)) : (p1));
            _nw2[(new BigNumber(2)).toNumber()] = false;
            _nw2[(new BigNumber(3)).toNumber()] = p1;
            _nw2[(new BigNumber(4)).toNumber()] = p1;
            _nw2[(new BigNumber(5)).toNumber()] = p1;
            _nw2[(new BigNumber(6)).toNumber()] = p1;
            _nw2[(new BigNumber(7)).toNumber()] = p1;
            _150_v6 = _nw2;
            let _151_v7;
            _151_v7 = _dafny.Set.fromElements(_145_v1);
            let _152_v8;
            _152_v8 = _module.D1.create_DC2(_145_v1, _145_v1, p1);
            let _153_v9;
            let _nw3 = new _module.C5();
            _nw3.__ctor(p1, _150_v6, _151_v7, _152_v8, p1);
            _153_v9 = _nw3;
            let _154_v10;
            _154_v10 = _dafny.Seq.of(_153_v9);
            let _155_v11;
            _155_v11 = _dafny.Map.Empty.slice().updateUnsafe(_154_v10,_145_v1);
            let _156_v12;
            _156_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,(new BigNumber(499)).isLessThan(new BigNumber((_155_v11).length)));
            _156_v12 = (_156_v12).update(p1, (((_156_v12).contains(p1)) ? ((_156_v12).get(p1)) : (p1)));
          }
        }
      }
      let _157_v13;
      let _nw4 = Array((new BigNumber(12)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = p1;
      _nw4[(_dafny.ONE).toNumber()] = p1;
      _nw4[(new BigNumber(2)).toNumber()] = p1;
      _nw4[(new BigNumber(3)).toNumber()] = true;
      _nw4[(new BigNumber(4)).toNumber()] = p1;
      _nw4[(new BigNumber(5)).toNumber()] = p1;
      _nw4[(new BigNumber(6)).toNumber()] = p1;
      _nw4[(new BigNumber(7)).toNumber()] = p1;
      _nw4[(new BigNumber(8)).toNumber()] = p1;
      _nw4[(new BigNumber(9)).toNumber()] = p1;
      _nw4[(new BigNumber(10)).toNumber()] = p1;
      _nw4[(new BigNumber(11)).toNumber()] = p1;
      _157_v13 = _nw4;
      let _158_v14;
      _158_v14 = new BigNumber(-466);
      let _159_v15;
      _159_v15 = _dafny.Set.fromElements(_158_v14, _158_v14);
      let _160_v16;
      _160_v16 = _dafny.Seq.of(p1, p1);
      let _161_v17;
      _161_v17 = _module.D1.create_DC2(new BigNumber((_160_v16).length), _158_v14, p1);
      let _162_v18;
      let _nw5 = new _module.C5();
      _nw5.__ctor(p1, _157_v13, _159_v15, _161_v17, p1);
      _162_v18 = _nw5;
      let _163_v19;
      _163_v19 = _dafny.Seq.of(_162_v18);
      let _164_v20;
      _164_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_163_v19)[_module.__default.safeIndex(_158_v14, new BigNumber((_163_v19).length))]);
      let _165_v21;
      let _nw6 = new _module.C5();
      _nw6.__ctor(p1, _157_v13, _159_v15, _161_v17, p1);
      _165_v21 = _nw6;
      let _166_v22;
      _166_v22 = _dafny.Set.fromElements(_165_v21);
      let _hi0 = new BigNumber((((_162_v18.f24) ? (_166_v22) : (_166_v22))).length);
      for (let _167_i2 = new BigNumber((_164_v20).length); _167_i2.isLessThan(_hi0); _167_i2 = _167_i2.plus(_dafny.ONE)) {
        let _168_v23;
        _168_v23 = _dafny.Map.Empty.slice().updateUnsafe(_162_v18.f24,false);
        let _169_v24;
        _169_v24 = _dafny.Seq.of(_168_v23);
        let _170_v25;
        _170_v25 = _dafny.Map.Empty.slice().updateUnsafe((_165_v21).f13,_dafny.Seq.of(_168_v23, _168_v23));
        let _171_v26;
        _171_v26 = _dafny.Set.fromElements(_160_v16, _dafny.Seq.of(false, (_165_v21).f13), _160_v16);
        let _172_v27;
        _172_v27 = _dafny.Set.fromElements(_dafny.Seq.of((_162_v18).fm52(_167_i2, _171_v26, globalState), _162_v18.f24));
        let _173_v28;
        let _nw7 = new _module.C0();
        _nw7.__ctor(!_dafny.areEqual(_169_v24, (((_170_v25).contains(_162_v18.f24)) ? ((_170_v25).get(_162_v18.f24)) : (_169_v24))), new _dafny.CodePoint('b'.codePointAt(0)), _165_v21.f12, (_162_v18).fm52(_158_v14, _172_v27, globalState));
        _173_v28 = _nw7;
        let _174_v29;
        _174_v29 = _dafny.Seq.of(_173_v28);
        _173_v28 = (_174_v29)[_module.__default.safeIndex((_173_v28).fm16((_165_v21).f13, _167_i2, globalState), new BigNumber((_174_v29).length))];
        let _175_v30;
        let _nw8 = Array((new BigNumber(2)).toNumber()).fill([]);
        _175_v30 = _nw8;
        let _176_v31;
        let _nw9 = Array((new BigNumber(24)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = _175_v30;
        _nw9[(_dafny.ONE).toNumber()] = _175_v30;
        _nw9[(new BigNumber(2)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(3)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(4)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(5)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(6)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(7)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(8)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(9)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(10)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(11)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(12)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(13)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(14)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(15)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(16)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(17)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(18)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(19)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(20)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(21)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(22)).toNumber()] = _175_v30;
        _nw9[(new BigNumber(23)).toNumber()] = _175_v30;
        _176_v31 = _nw9;
        let _index2 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_176_v31).length));
        (_176_v31)[_index2] = _175_v30;
        let _index3 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_176_v31).length));
        (_176_v31)[_index3] = _175_v30;
        let _177_v32;
        _177_v32 = _dafny.Seq.UnicodeFromString("qrcl");
        _177_v32 = _module.__default.fm20(globalState);
        (globalState).f7 = _158_v14;
      }
      (_162_v18).f24 = ((p1) ? (!(_162_v18.f24)) : ((p1) && (_162_v18.f24)));
      let _178_v33;
      _178_v33 = _dafny.Seq.UnicodeFromString("ulnyf");
      let _179_v34;
      _179_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_178_v33).length),_158_v14);
      (_162_v18).f24 = _module.__default.fm30((_dafny.ZERO).minus(new BigNumber((((_162_v18.f24) ? (_179_v34) : (_179_v34))).length)), globalState);
      let _180_v35;
      _180_v35 = _module.D1.create_DC1(_178_v33);
      let _181_v36;
      let _nw10 = new _module.C3();
      _nw10.__ctor((_165_v21).f13, _180_v35, _162_v18.f24, _157_v13, _159_v15, _161_v17, ((p1) ? (_162_v18.f24) : (p1)));
      _181_v36 = _nw10;
      let _pat_let_tv0 = _181_v36;
      let _pat_let_tv1 = _162_v18;
      let _pat_let_tv2 = _165_v21;
      let _pat_let_tv3 = _181_v36;
      let _pat_let_tv4 = _162_v18;
      let _pat_let_tv5 = _181_v36;
      let _pat_let_tv6 = _160_v16;
      let _pat_let_tv7 = _158_v14;
      let _pat_let_tv8 = _160_v16;
      r0 = function (_source6) {
        if (_source6.is_DC30) {
          let _182___mcc_h0 = (_source6).cf58;
          let _183___mcc_h1 = (_source6).cf59;
          let _184_cf59 = _183___mcc_h1;
          let _185_cf58 = _182___mcc_h0;
          let _186_dt__update__tmp_h0 = _dafny.MultiSet.fromElements((_pat_let_tv0).f18, _pat_let_tv1.f24);
          let _187_dt__update_hcf0_h0 = _dafny.MultiSet.FromArray(_dafny.Seq.of((_pat_let_tv2).f13));
          return _187_dt__update_hcf0_h0;
        } else if (_source6.is_DC31) {
          let _188___mcc_h2 = (_source6).cf60;
          let _189___mcc_h3 = (_source6).cf61;
          let _190___mcc_h4 = (_source6).cf62;
          let _191___mcc_h5 = (_source6).cf63;
          let _192___mcc_h6 = (_source6).cf64;
          let _193_cf64 = _192___mcc_h6;
          let _194_cf63 = _191___mcc_h5;
          let _195_cf62 = _190___mcc_h4;
          let _196_cf61 = _189___mcc_h3;
          let _197_cf60 = _188___mcc_h2;
          return _dafny.MultiSet.fromElements((_pat_let_tv3).f18, _197_cf60);
        } else if (_source6.is_DC32) {
          let _198___mcc_h7 = (_source6).cf65;
          let _199___mcc_h8 = (_source6).cf66;
          let _200___mcc_h9 = (_source6).cf67;
          let _201___mcc_h10 = (_source6).cf68;
          let _202___mcc_h11 = (_source6).cf69;
          let _203_cf69 = _202___mcc_h11;
          let _204_cf68 = _201___mcc_h10;
          let _205_cf67 = _200___mcc_h9;
          let _206_cf66 = _199___mcc_h8;
          let _207_cf65 = _198___mcc_h7;
          return _dafny.MultiSet.fromElements(!(true), _205_cf67, _pat_let_tv4.f24);
        } else {
          let _208___mcc_h12 = (_source6).cf57;
          let _209_cf57 = _208___mcc_h12;
          return _dafny.MultiSet.fromElements((_pat_let_tv5).f18, (_pat_let_tv6)[_module.__default.safeIndex(_pat_let_tv7, new BigNumber((_pat_let_tv8).length))]);
        }
      }(_module.D12.create_DC32(p1, new BigNumber((_178_v33).length), (_181_v36).f18, _162_v18.f24, _158_v14));
      return r0;
    }
    static Main(__noArgsParameter) {
      let _210_v1;
      let _init0 = function (_211_i0) {
        return function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(633),new BigNumber(((_dafny.MultiSet.fromElements(false))).cardinality()))).Keys.Elements) {
            let _212_v0 = _compr_23;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(633),new BigNumber(((_dafny.MultiSet.fromElements(false))).cardinality()))).contains(_212_v0)) {
              _coll23.push([_module.__default.safeDivisionInt(_212_v0, new BigNumber(805)),false]);
            }
          }
          return _coll23;
        }();
      };
      let _nw11 = Array((new BigNumber(2)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw11.length); _i0_0++) {
        _nw11[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _210_v1 = _nw11;
      let _213_v2;
      _213_v2 = true;
      let _214_v3;
      _214_v3 = _dafny.Seq.of(_213_v2);
      let _215_v4;
      _215_v4 = _dafny.Seq.UnicodeFromString("dgxao");
      let _216_v5;
      let _nw12 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _216_v5 = _nw12;
      let _217_v6;
      _217_v6 = _dafny.MultiSet.fromElements(_216_v5, _216_v5);
      let _218_v7;
      _218_v7 = _dafny.Map.Empty.slice().updateUnsafe(_213_v2,_215_v4);
      let _219_v8;
      _219_v8 = new _dafny.CodePoint('g'.codePointAt(0));
      let _220_v9;
      _220_v9 = new BigNumber(336);
      let _221_v10;
      _221_v10 = _dafny.Map.Empty.slice().updateUnsafe(_220_v9,_215_v4);
      let _222_v11;
      _222_v11 = _module.D1.create_DC1(_215_v4);
      let _223_v12;
      let _nw13 = Array((new BigNumber(28)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = (((_218_v7).contains(_213_v2)) ? ((_218_v7).get(_213_v2)) : (_215_v4));
      _nw13[(_dafny.ONE).toNumber()] = _215_v4;
      _nw13[(new BigNumber(2)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(3)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(4)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(5)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(6)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(7)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(8)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(9)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(10)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("yauffnrew"), _module.__default.safeIndex(new BigNumber(752), new BigNumber((_dafny.Seq.UnicodeFromString("yauffnrew")).length)), _219_v8);
      _nw13[(new BigNumber(12)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(13)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(14)).toNumber()] = (((_221_v10).contains(_220_v9)) ? ((_221_v10).get(_220_v9)) : (_215_v4));
      _nw13[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_224_v8) => function (_225_i1) {
        return _224_v8;
      })(_219_v8));
      _nw13[(new BigNumber(16)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(17)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(18)).toNumber()] = (_222_v11).dtor_cf1;
      _nw13[(new BigNumber(19)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("tqk");
      _nw13[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("dgjly");
      _nw13[(new BigNumber(22)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), ((_226_v8) => function (_227_i2) {
        return _226_v8;
      })(_219_v8));
      _nw13[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("dffircy");
      _nw13[(new BigNumber(25)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(26)).toNumber()] = _215_v4;
      _nw13[(new BigNumber(27)).toNumber()] = _215_v4;
      _223_v12 = _nw13;
      let _228_globalState;
      let _nw14 = new _module.GlobalState();
      _nw14.__ctor(_210_v1, true, _214_v3, _215_v4, true, false, new BigNumber(142), new BigNumber(530), _dafny.Seq.Concat(_215_v4, _215_v4), _217_v6, _223_v12);
      _228_globalState = _nw14;
      let _229_v13;
      let _out0;
      _out0 = _module.__default.m0(new _dafny.CodePoint('g'.codePointAt(0)), _213_v2, _228_globalState);
      _229_v13 = _out0;
      let _230_v14;
      let _nw15 = Array((new BigNumber(6)).toNumber());
      _nw15[(_dafny.ZERO).toNumber()] = !(_213_v2);
      _nw15[(_dafny.ONE).toNumber()] = !(_213_v2);
      _nw15[(new BigNumber(2)).toNumber()] = _213_v2;
      _nw15[(new BigNumber(3)).toNumber()] = !(_213_v2);
      _nw15[(new BigNumber(4)).toNumber()] = true;
      _nw15[(new BigNumber(5)).toNumber()] = _213_v2;
      _230_v14 = _nw15;
      let _231_v15;
      _231_v15 = _module.D2.create_DC3(_230_v14);
      let _232_v16;
      _232_v16 = _dafny.Seq.of(_230_v14);
      let _233_v17;
      _233_v17 = _dafny.Set.fromElements(_213_v2);
      let _234_v18;
      _234_v18 = _dafny.MultiSet.fromElements(new BigNumber((_233_v17).length));
      let _235_v19;
      _235_v19 = _dafny.Map.Empty.slice().updateUnsafe(_234_v18,_213_v2);
      let _236_v20;
      _236_v20 = _module.D2.create_DC4(_216_v5, _213_v2, _220_v9, false, _230_v14);
      let _237_v21;
      let _nw16 = Array((new BigNumber(29)).toNumber());
      _nw16[(_dafny.ZERO).toNumber()] = ((_213_v2) ? (_230_v14) : (_230_v14));
      _nw16[(_dafny.ONE).toNumber()] = _230_v14;
      _nw16[(new BigNumber(2)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(3)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(4)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(5)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(6)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(7)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(8)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(9)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(10)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(11)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(12)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(13)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(14)).toNumber()] = ((_213_v2) ? (_230_v14) : (_230_v14));
      _nw16[(new BigNumber(15)).toNumber()] = (_231_v15).dtor_cf5;
      _nw16[(new BigNumber(16)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(17)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(18)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(19)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(20)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(21)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(22)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(23)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(24)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(25)).toNumber()] = (_232_v16)[_module.__default.safeIndex(new BigNumber((_235_v19).length), new BigNumber((_232_v16).length))];
      _nw16[(new BigNumber(26)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(27)).toNumber()] = _230_v14;
      _nw16[(new BigNumber(28)).toNumber()] = (_236_v20).dtor_cf10;
      _237_v21 = _nw16;
      let _index4 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length));
      (_237_v21)[_index4] = _230_v14;
      let _index5 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length));
      let _rhs0 = _230_v14;
      let _rhs1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("egift")).length), new BigNumber(294));
      let _lhs0 = _237_v21;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length));
      let _lhs2 = _228_globalState;
      _lhs0[_lhs1] = _rhs0;
      _lhs2.f7 = _rhs1;
      let _index6 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
      (_216_v5)[_index6] = _220_v9;
      let _index7 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
      (_216_v5)[_index7] = _220_v9;
      let _238_v22;
      _238_v22 = _dafny.Map.Empty.slice().updateUnsafe(_213_v2,_214_v3);
      let _hi1 = (_220_v9).multipliedBy(_module.__default.fm0((((_238_v22).contains(_213_v2)) ? ((_238_v22).get(_213_v2)) : (_214_v3)), _213_v2, _213_v2, (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _228_globalState));
      for (let _239_i3 = new BigNumber((_233_v17).length); _239_i3.isLessThan(_hi1); _239_i3 = _239_i3.plus(_dafny.ONE)) {
        _213_v2 = _213_v2;
        let _source7 = _module.D1.create_DC1(_dafny.Seq.Concat(_215_v4, _215_v4));
        if (_source7.is_DC2) {
          let _240___mcc_h0 = (_source7).cf2;
          let _241___mcc_h1 = (_source7).cf3;
          let _242___mcc_h2 = (_source7).cf4;
          let _243_cf4 = _242___mcc_h2;
          let _244_cf3 = _241___mcc_h1;
          let _245_cf2 = _240___mcc_h0;
          let _246_v23;
          _246_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_215_v4).length),_230_v14);
          let _247_v24;
          _247_v24 = _module.D1.create_DC2(new BigNumber(87), new BigNumber((_215_v4).length), _243_cf4);
          let _248_v25;
          let _nw17 = new _module.C2();
          _nw17.__ctor(_244_cf3, !(!(_243_cf4)), _222_v11, _213_v2, ((_243_cf4) ? (_230_v14) : ((((_246_v23).contains(_245_cf2)) ? ((_246_v23).get(_245_cf2)) : ((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])))), _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_243_cf4, _243_cf4)).length)), _247_v24, _243_cf4);
          _248_v25 = _nw17;
          let _249_v26;
          _249_v26 = _dafny.Map.Empty.slice().updateUnsafe(_243_cf4,_220_v9);
          _249_v26 = (_249_v26).update(!((_248_v25).f22), (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]);
          _244_cf3 = (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))];
          _216_v5 = _216_v5;
        } else {
          let _250___mcc_h3 = (_source7).cf1;
          let _251_cf1 = _250___mcc_h3;
          let _252_v27;
          let _nw18 = new _module.C5();
          _nw18.__ctor(_213_v2, _230_v14, _module.__default.fm64(_228_globalState), _module.__default.fm37((_dafny.ZERO).minus((_dafny.ZERO).minus((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))])), _239_i3, (_dafny.ZERO).minus(_220_v9), (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _228_globalState), (_239_i3).isLessThan((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]));
          _252_v27 = _nw18;
          let _253_v29;
          _253_v29 = _module.D1.create_DC2(_239_i3, new BigNumber((_215_v4).length), _213_v2);
          let _254_v30;
          let _nw19 = new _module.C2();
          _nw19.__ctor(_239_i3, _252_v27.f24, _module.D1.create_DC1(_251_cf1), _213_v2, (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))], function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(536), new BigNumber(571))) {
              let _255_v28 = _compr_24;
              if (((new BigNumber(536)).isLessThanOrEqualTo(_255_v28)) && ((_255_v28).isLessThan(new BigNumber(571)))) {
                _coll24.add((_255_v28).plus(_220_v9));
              }
            }
            return _coll24;
          }(), _253_v29, _252_v27.f24);
          _254_v30 = _nw19;
          let _256_v31;
          _256_v31 = _dafny.Map.Empty.slice().updateUnsafe(_252_v27.f24,_254_v30);
          let _257_v32;
          _257_v32 = _dafny.Seq.of(new BigNumber((_256_v31).length), (_254_v30).f21);
          let _258_v33;
          let _nw20 = Array((new BigNumber(5)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = (_module.D11.create_DC27(true, _219_v8, _220_v9, (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _257_v32)).dtor_cf52;
          _nw20[(_dafny.ONE).toNumber()] = _219_v8;
          _nw20[(new BigNumber(2)).toNumber()] = (_215_v4)[_module.__default.safeIndex(new BigNumber(-56), new BigNumber((_215_v4).length))];
          _nw20[(new BigNumber(3)).toNumber()] = _219_v8;
          _nw20[(new BigNumber(4)).toNumber()] = _219_v8;
          _258_v33 = _nw20;
          _258_v33 = _258_v33;
          let _259_v34;
          let _init1 = ((_260_v30, _261_v32) => function (_262_i4) {
            return _dafny.MultiSet.fromElements(_module.D19.create_DC50((_260_v30).f22, new BigNumber(-629), _261_v32), _module.D19.create_DC50((_260_v30).f22, (_260_v30).f21, _261_v32));
          })(_254_v30, _257_v32);
          let _nw21 = Array((_dafny.ONE).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw21.length); _i0_1++) {
            _nw21[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _259_v34 = _nw21;
          let _index8 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length));
          let _rhs2 = _216_v5;
          let _rhs3 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
          let _rhs4 = (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))];
          let _rhs5 = _259_v34;
          let _lhs3 = _237_v21;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length));
          let _lhs5 = _228_globalState;
          _216_v5 = _rhs2;
          _lhs3[_lhs4] = _rhs3;
          _lhs5.f7 = _rhs4;
          _259_v34 = _rhs5;
          let _263_v35;
          let _out1;
          _out1 = (_254_v30).m6((_254_v30).f22, _228_globalState);
          _263_v35 = _out1;
        }
        let _264_v36;
        _264_v36 = _dafny.Set.fromElements(_220_v9);
        let _265_v37;
        _265_v37 = _module.D19.create_DC49(_264_v36);
        let _pat_let_tv9 = _264_v36;
        let _arr0 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
        let _index9 = _module.__default.safeIndex(new BigNumber(315), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
        _arr0[_index9] = ((function (_pat_let0_0) {
          return function (_266_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_267_dt__update_hcf93_h0) {
                return _module.D19.create_DC49(_267_dt__update_hcf93_h0);
              }(_pat_let1_0);
            }(_pat_let_tv9);
          }(_pat_let0_0);
        }(_265_v37)).dtor_cf93).IsDisjointFrom(_264_v36);
        let _index10 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        let _arr1 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
        let _index11 = _module.__default.safeIndex(new BigNumber(315), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
        let _rhs6 = _220_v9;
        let _rhs7 = ((_213_v2) ? (_220_v9) : (new BigNumber(-487)));
        let _rhs8 = !(true);
        let _rhs9 = false;
        let _lhs6 = _228_globalState;
        let _lhs7 = _216_v5;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        let _lhs9 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
        let _lhs10 = _module.__default.safeIndex(new BigNumber(315), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
        _lhs6.f7 = _rhs6;
        _lhs7[_lhs8] = _rhs7;
        _213_v2 = _rhs8;
        _lhs9[_lhs10] = _rhs9;
        let _index12 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        (_216_v5)[_index12] = _220_v9;
      }
      let _268_i5;
      _268_i5 = _dafny.ZERO;
      L1: {
        while (false) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_268_i5)) {
              break L1;
            }
            _268_i5 = (_268_i5).plus(_dafny.ONE);
            let _arr2 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
            let _index13 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
            _arr2[_index13] = _213_v2;
            let _arr3 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
            let _index14 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
            _arr3[_index14] = _213_v2;
            if (_213_v2) {
              let _269_v38;
              _269_v38 = _dafny.Map.Empty.slice().updateUnsafe((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))],_220_v9);
              let _270_v40;
              _270_v40 = _module.D1.create_DC2(_220_v9, new BigNumber((_dafny.MultiSet.fromElements(_214_v3)).cardinality()), !(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))]));
              let _271_v41;
              let _nw22 = new _module.C6();
              _nw22.__ctor(_222_v11, _module.__default.fm17(_228_globalState), _230_v14, function () {
                let _coll25 = new _dafny.Set();
                for (const _compr_25 of (_269_v38).Keys.Elements) {
                  let _272_v39 = _compr_25;
                  if ((_269_v38).contains(_272_v39)) {
                    _coll25.add((_272_v39).multipliedBy((_dafny.ZERO).minus(new BigNumber(-879))));
                  }
                }
                return _coll25;
              }(), _270_v40, ((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))]);
              _271_v41 = _nw22;
              _220_v9 = new BigNumber((_233_v17).length);
              let _273_v42;
              _273_v42 = _dafny.MultiSet.fromElements(_230_v14, _230_v14);
              let _arr4 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
              let _index15 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
              _arr4[_index15] = ((_273_v42).Difference(_273_v42)).IsSubsetOf(_273_v42);
              let _274_v43;
              let _nw23 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
              _274_v43 = _nw23;
              _274_v43 = _274_v43;
              let _275_v44;
              _275_v44 = _dafny.Map.Empty.slice().updateUnsafe(_219_v8,_220_v9);
              let _276_v45;
              _276_v45 = _dafny.Seq.of(_275_v44);
              _276_v45 = _276_v45;
            } else {
              _215_v4 = _dafny.Seq.Concat(_215_v4, _215_v4);
              let _277_v46;
              let _nw24 = new _module.C4();
              _nw24.__ctor();
              _277_v46 = _nw24;
              let _278_v47;
              _278_v47 = _dafny.Map.Empty.slice().updateUnsafe(_220_v9,!(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))]));
              let _279_v48;
              _279_v48 = _dafny.Map.Empty.slice().updateUnsafe(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))],_278_v47);
              let _280_v49;
              _280_v49 = _dafny.Map.Empty.slice().updateUnsafe(_277_v46,new BigNumber(((((_279_v48).contains(_213_v2)) ? ((_279_v48).get(_213_v2)) : (_dafny.Map.Empty.slice().updateUnsafe(_220_v9,((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))])))).length));
              let _281_v50;
              _281_v50 = _module.D1.create_DC2((_dafny.ZERO).minus((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_220_v9))).cardinality()), _213_v2);
              let _282_v51;
              let _nw25 = new _module.C0();
              _nw25.__ctor(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))], (_215_v4)[_module.__default.safeIndex((((_280_v49).contains(_277_v46)) ? ((_280_v49).get(_277_v46)) : (new BigNumber(86))), new BigNumber((_215_v4).length))], _281_v50, (((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))]) === (true));
              _282_v51 = _nw25;
              _282_v51 = _282_v51;
              let _283_v52;
              _283_v52 = _dafny.Map.Empty.slice().updateUnsafe(_220_v9,new BigNumber(-643));
              let _284_v53;
              _284_v53 = _dafny.Seq.of((_dafny.ZERO).minus((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]), new BigNumber((_283_v52).length));
              let _285_v54;
              _285_v54 = _dafny.Set.fromElements((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]);
              let _286_v55;
              let _nw26 = new _module.C3();
              _nw26.__ctor(_282_v51.f19, _222_v11, _module.__default.fm30(new BigNumber((_284_v53).length), _228_globalState), _230_v14, _285_v54, _281_v50, ((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))])[_module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length))]);
              _286_v55 = _nw26;
              let _287_v56;
              _287_v56 = _dafny.Seq.of(_286_v55, _286_v55, _286_v55);
              _286_v55 = (_287_v56)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_287_v56).length))];
              let _288_v57;
              _288_v57 = _dafny.MultiSet.fromElements(_213_v2);
              (_282_v51).f19 = ((_288_v57).Difference(_288_v57)).IsDisjointFrom((_288_v57).Intersect(_288_v57));
              (_228_globalState).f2 = _214_v3;
            }
            let _arr5 = (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))];
            let _index16 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))]).length));
            _arr5[_index16] = !(!(_module.__default.fm3((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _228_globalState)));
            _213_v2 = false;
          }
        }
      }
      let _289_v58;
      _289_v58 = _dafny.Map.Empty.slice().updateUnsafe(_213_v2,_213_v2);
      let _index17 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
      (_230_v14)[_index17] = (((_289_v58).contains(_213_v2)) ? ((_289_v58).get(_213_v2)) : (_213_v2));
      let _index18 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
      (_230_v14)[_index18] = ((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]).isEqualTo(new BigNumber((_dafny.Seq.of(_213_v2)).length));
      _289_v58 = (_289_v58).update(_213_v2, ((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]) === (_213_v2));
      let _hi2 = _220_v9;
      for (let _290_i6 = _220_v9; _290_i6.isLessThan(_hi2); _290_i6 = _290_i6.plus(_dafny.ONE)) {
        _213_v2 = _dafny.Seq.IsPrefixOf(_232_v16, _232_v16);
        let _291_v59;
        let _out2;
        _out2 = _module.__default.m0(_219_v8, _module.__default.fm17(_228_globalState), _228_globalState);
        _291_v59 = _out2;
        let _292_v60;
        let _nw27 = Array((new BigNumber(22)).toNumber()).fill([]);
        _292_v60 = _nw27;
        let _index19 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_292_v60).length));
        (_292_v60)[_index19] = _216_v5;
        let _index20 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_292_v60).length));
        (_292_v60)[_index20] = _216_v5;
        let _index21 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_292_v60).length));
        (_292_v60)[_index21] = _216_v5;
      }
      let _293_v61;
      _293_v61 = _module.D19.create_DC50((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _dafny.Seq.of(new BigNumber((_module.__default.fm24(new BigNumber(997), _228_globalState)).length), _220_v9));
      let _294_v62;
      let _out3;
      _out3 = _module.__default.m0(_219_v8, (_293_v61).dtor_cf94, _228_globalState);
      _294_v62 = _out3;
      let _index22 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
      let _rhs10 = _module.__default.fm0(((_module.__default.fm17(_228_globalState)) ? (_214_v3) : (_214_v3)), (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], !((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]), _module.__default.fm0(_214_v3, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], new BigNumber(-259), _228_globalState), _228_globalState);
      let _rhs11 = _module.__default.fm3((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _228_globalState);
      let _rhs12 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))],_213_v2)).length);
      let _lhs11 = _228_globalState;
      let _lhs12 = _230_v14;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
      _lhs11.f7 = _rhs10;
      _lhs12[_lhs13] = _rhs11;
      _220_v9 = _rhs12;
      let _295_v63;
      let _init2 = ((_296_v2, _297_v14) => function (_298_i8) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(!(_296_v2)),(_297_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_297_v14).length))]);
      })(_213_v2, _230_v14);
      let _nw28 = Array((new BigNumber(12)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw28.length); _i0_2++) {
        _nw28[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _295_v63 = _nw28;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_295_v63).length))) {
        let _299_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_299_i7)) && ((_299_i7).isLessThan(new BigNumber((_295_v63).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_295_v63, (_299_i7).toNumber(), _dafny.Map.Empty.slice().updateUnsafe((_233_v17).equals(_233_v17),(((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]) ? (false) : (_213_v2)))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      (_228_globalState).f7 = ((_dafny.ZERO).minus((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))])).multipliedBy(((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]).multipliedBy((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]));
      let _ingredients1 = [];
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_230_v14).length))) {
        let _300_i9 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_300_i9)) && ((_300_i9).isLessThan(new BigNumber((_230_v14).length))))) {
          _ingredients1.push(_dafny.Tuple.of(_230_v14, (_300_i9).toNumber(), ((function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_301_v5) => function (_302_i10) {
              return (_dafny.ZERO).minus((_301_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_301_v5).length))]);
            })(_216_v5))).Elements) {
              let _303_v64 = _compr_26;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_304_v5) => function (_302_i10) {
                return (_dafny.ZERO).minus((_304_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_304_v5).length))]);
              })(_216_v5)), _303_v64)) {
                _coll26.add((_303_v64).plus(new BigNumber((_dafny.Seq.of(new BigNumber(51))).length)));
              }
            }
            return _coll26;
          }()).Union(_dafny.Set.fromElements(new BigNumber(-321)))).IsProperSubsetOf((_dafny.Set.fromElements(new BigNumber((_289_v58).length))).Union(_dafny.Set.fromElements(_220_v9)))));
        }
      }
      for (const _tup1 of _ingredients1) {
        _tup1[0][_tup1[1]] = _tup1[2];
      }
      let _305_v65;
      _305_v65 = _dafny.Map.Empty.slice().updateUnsafe(_213_v2,_220_v9);
      let _hi3 = new BigNumber((_dafny.Seq.of((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))])).length);
      for (let _306_i11 = _module.__default.safeModuloInt((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], new BigNumber((_305_v65).length)); _306_i11.isLessThan(_hi3); _306_i11 = _306_i11.plus(_dafny.ONE)) {
        let _307_v67;
        _307_v67 = _dafny.Seq.of(_220_v9, (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]);
        let _308_v68;
        _308_v68 = _module.D13.create_DC34(_220_v9, _module.__default.fm65(_213_v2, _307_v67, _219_v8, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], _228_globalState));
        let _309_v69;
        _309_v69 = _module.D1.create_DC2(_220_v9, (_308_v68).dtor_cf71, _213_v2);
        let _310_v70;
        let _nw29 = new _module.C6();
        _nw29.__ctor(_222_v11, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], (_237_v21)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_237_v21).length))], function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of _dafny.IntegerRange(new BigNumber(163), new BigNumber(958))) {
            let _311_v66 = _compr_27;
            if (((new BigNumber(163)).isLessThanOrEqualTo(_311_v66)) && ((_311_v66).isLessThan(new BigNumber(958)))) {
              _coll27.add(_module.__default.safeModuloInt(_311_v66, _306_i11));
            }
          }
          return _coll27;
        }(), _309_v69, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]);
        _310_v70 = _nw29;
        let _312_v71;
        _312_v71 = _dafny.Map.Empty.slice().updateUnsafe(_306_i11,_310_v70);
        let _pat_let_tv10 = _215_v4;
        let _313_v72;
        let _nw30 = new _module.C3();
        _nw30.__ctor(!((_312_v71).equals(_312_v71)), function (_pat_let2_0) {
          return function (_314_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_315_dt__update_hcf1_h0) {
                return _module.D1.create_DC1(_315_dt__update_hcf1_h0);
              }(_pat_let3_0);
            }(_pat_let_tv10);
          }(_pat_let2_0);
        }(_module.D1.create_DC1(_215_v4)), (_310_v70).f13, (_310_v70).f14, (_310_v70).f15, _309_v69, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]);
        _313_v72 = _nw30;
        _313_v72 = (((_310_v70).f17) ? (_313_v72) : (_313_v72));
        _234_v18 = _234_v18;
        _213_v2 = (_220_v9).isLessThan(_220_v9);
        let _316_v73;
        _316_v73 = _dafny.Seq.of(_305_v65);
        let _317_v74;
        let _nw31 = new _module.C8();
        _nw31.__ctor((_310_v70).f17);
        _317_v74 = _nw31;
        let _318_v75;
        _318_v75 = _dafny.Set.fromElements(_317_v74);
        let _319_v76;
        let _nw32 = Array((new BigNumber(16)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = (_305_v65).Merge((_305_v65).update((_313_v72).f13, _306_i11));
        _nw32[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_306_i11);
        _nw32[(new BigNumber(2)).toNumber()] = (_310_v70).fm10((_310_v70).f13, (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], _dafny.Seq.of((_310_v70).f17), _228_globalState);
        _nw32[(new BigNumber(3)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(4)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(5)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(6)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(7)).toNumber()] = (_305_v65).Merge(_305_v65);
        _nw32[(new BigNumber(8)).toNumber()] = (_305_v65).Merge((_316_v73)[_module.__default.safeIndex((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], new BigNumber((_316_v73).length))]);
        _nw32[(new BigNumber(9)).toNumber()] = (_305_v65).Merge(_305_v65);
        _nw32[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_310_v70).f17,new BigNumber((_318_v75).length))).Merge((_module.__default.fm67((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], _228_globalState)).dtor_cf32);
        _nw32[(new BigNumber(11)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(12)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(13)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(14)).toNumber()] = _305_v65;
        _nw32[(new BigNumber(15)).toNumber()] = _305_v65;
        _319_v76 = _nw32;
        _319_v76 = _319_v76;
      }
      if (_213_v2) {
        let _320_v77;
        let _out4;
        _out4 = _module.__default.m0((_215_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("oxurkuj")).length), new BigNumber((_215_v4).length))], (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], _228_globalState);
        _320_v77 = _out4;
        (_228_globalState).f7 = (_220_v9).multipliedBy(_220_v9);
        let _321_v78;
        let _out5;
        _out5 = _module.__default.m0(_219_v8, _213_v2, _228_globalState);
        _321_v78 = _out5;
        let _322_v79;
        _322_v79 = _module.D16.create_DC41(!(_213_v2));
        _213_v2 = (_322_v79).dtor_cf84;
        let _index23 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
        (_230_v14)[_index23] = ((((_dafny.ZERO).minus(_220_v9)).isLessThan(_220_v9)) ? (_213_v2) : (!(!_dafny.areEqual(_215_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-115)), ((_323_v8) => function (_324_i12) {
          return _323_v8;
        })(_219_v8))))));
      } else {
        let _325_v80;
        _325_v80 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))],_233_v17),((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]) || ((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]));
        let _326_v81;
        _326_v81 = _dafny.Map.Empty.slice().updateUnsafe((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))],_233_v17);
        _325_v80 = (_325_v80).update(_326_v81, _213_v2);
        _215_v4 = _dafny.Seq.Concat(_215_v4, _dafny.Seq.update(((_213_v2) ? (_215_v4) : (_215_v4)), _module.__default.safeIndex((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], new BigNumber((((_213_v2) ? (_215_v4) : (_215_v4))).length)), _219_v8));
        let _index24 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
        let _rhs13 = false;
        let _rhs14 = (_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))];
        let _lhs14 = _230_v14;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length));
        _213_v2 = _rhs13;
        _lhs14[_lhs15] = _rhs14;
        let _327_v82;
        _327_v82 = _dafny.Map.Empty.slice().updateUnsafe(_219_v8,(_234_v18).IsProperSubsetOf(_234_v18));
        _213_v2 = (((_327_v82).contains(new _dafny.CodePoint('x'.codePointAt(0)))) ? ((_327_v82).get(new _dafny.CodePoint('x'.codePointAt(0)))) : (_213_v2));
        let _328_v83;
        let _nw33 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
        _328_v83 = _nw33;
        let _329_v84;
        _329_v84 = _dafny.MultiSet.fromElements(_215_v4, _215_v4);
        let _index25 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_328_v83).length));
        (_328_v83)[_index25] = _329_v84;
        let _index26 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_328_v83).length));
        (_328_v83)[_index26] = _329_v84;
      }
      let _330_v85;
      _330_v85 = _dafny.Map.Empty.slice().updateUnsafe(_220_v9,(_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]);
      let _331_v86;
      _331_v86 = _dafny.Seq.of((_dafny.ZERO).minus(_220_v9), _220_v9, _220_v9, new BigNumber((_330_v85).length));
      let _hi4 = (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))];
      for (let _332_i13 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_331_v86, _dafny.Seq.of((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]))).length)); _332_i13.isLessThan(_hi4); _332_i13 = _332_i13.plus(_dafny.ONE)) {
        let _333_v87;
        let _nw34 = Array((new BigNumber(4)).toNumber()).fill(_module.D1.Default());
        _333_v87 = _nw34;
        let _334_v88;
        let _nw35 = new _module.C9();
        _nw35.__ctor();
        _334_v88 = _nw35;
        let _335_v89;
        _335_v89 = _dafny.Map.Empty.slice().updateUnsafe((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))],_333_v87);
        let _336_v90;
        _336_v90 = _module.D1.create_DC2(_332_i13, (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))], _213_v2);
        let _337_v91;
        let _nw36 = new _module.C0();
        _nw36.__ctor((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))], _module.__default.fm58(_228_globalState), _336_v90, _213_v2);
        _337_v91 = _nw36;
        let _338_v92;
        _338_v92 = _dafny.Set.fromElements(_337_v91, _337_v91);
        let _339_v93;
        _339_v93 = _dafny.Map.Empty.slice().updateUnsafe(_334_v88,(((_335_v89).contains(new BigNumber(((_module.D20.create_DC52(_338_v92)).dtor_cf99).length))) ? ((_335_v89).get(new BigNumber(((_module.D20.create_DC52(_338_v92)).dtor_cf99).length))) : (_333_v87)));
        let _index27 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        let _rhs15 = (_220_v9).plus(_220_v9);
        let _rhs16 = !((_230_v14)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_230_v14).length))]);
        let _rhs17 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_332_i13));
        let _rhs18 = (((_339_v93).contains(_334_v88)) ? ((_339_v93).get(_334_v88)) : (_333_v87));
        let _lhs16 = _216_v5;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        _lhs16[_lhs17] = _rhs15;
        _213_v2 = _rhs16;
        _220_v9 = _rhs17;
        _333_v87 = _rhs18;
        _331_v86 = _331_v86;
        _220_v9 = ((_module.__default.fm17(_228_globalState)) ? (_220_v9) : ((_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))]));
        let _340_v94;
        let _nw37 = new _module.C4();
        _nw37.__ctor();
        _340_v94 = _nw37;
        let _index28 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        let _rhs19 = (_216_v5)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length))];
        let _rhs20 = _332_i13;
        let _rhs21 = _216_v5;
        let _rhs22 = ((_337_v91.f19) ? (_340_v94) : (_340_v94));
        let _lhs18 = _228_globalState;
        let _lhs19 = _216_v5;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_216_v5).length));
        _lhs18.f7 = _rhs19;
        _lhs19[_lhs20] = _rhs20;
        _216_v5 = _rhs21;
        _340_v94 = _rhs22;
      }
      process.stdout.write(_dafny.toString(((_210_v1)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_210_v1)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_213_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_214_v3, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_215_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_217_v6).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_218_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("dgxao")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_219_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_220_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_221_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(336),_dafny.Seq.UnicodeFromString("dgxao")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_222_v11).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_223_v12)[new BigNumber(15)], _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(21)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(22)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_223_v12)[new BigNumber(23)], _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(24)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(25)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(26)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_223_v12)[new BigNumber(27)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_228_globalState).f0)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_228_globalState).f0)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_228_globalState.f2, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_228_globalState).f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_228_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_228_globalState).f9).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_228_globalState).f10)[new BigNumber(15)], _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(21)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(22)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_228_globalState).f10)[new BigNumber(23)], _dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(24)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(25)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(26)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_228_globalState).f10)[new BigNumber(27)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_229_v13)).equals(_dafny.MultiSet.fromElements(false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_231_v15).dtor_cf5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_232_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v17).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v18).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_dafny.ONE),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v20).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v20).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v20).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_236_v20).dtor_cf10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(11)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(12)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(13)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(14)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(15)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(16)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(17)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(18)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(19)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(20)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(21)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(22)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(23)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(24)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(25)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(26)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(27)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_237_v21)[new BigNumber(28)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v22).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_268_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v58).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_293_v61).dtor_cf94));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_293_v61).dtor_cf95));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_293_v61).dtor_cf96, _dafny.Seq.of(new BigNumber(3), new BigNumber(336)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_294_v62)).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_295_v63)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_305_v65).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_330_v85).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_331_v86, _dafny.Seq.of(new BigNumber(-1), _dafny.ONE, _dafny.ONE, _dafny.ONE))));
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
      return _dafny.MultiSet.Empty;
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
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + this.cf1.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC4(cf6, cf7, cf8, cf9, cf10) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D2(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf5 === other.cf5;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4([], false, _dafny.ZERO, false, []);
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
    static create_DC6(cf12, cf13) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC7() {
      let $dt = new D3(1);
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D3(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.ZERO, false);
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
    static create_DC9(cf15) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15;
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
    static create_DC10(cf16) {
      let $dt = new D5(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf18, cf19, cf20, cf21) {
      let $dt = new D6(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC13(cf22, cf23, cf24, cf25, cf26) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D6(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC14(cf27) {
      let $dt = new D6(3);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf18) + ", " + this.cf19.toVerbatimString(true) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC12(false, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false);
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
    static create_DC16(cf29) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D7(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC17(cf30) {
      let $dt = new D7(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16(_dafny.ZERO);
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
    static create_DC19(cf32, cf33, cf34, cf35) {
      let $dt = new D8(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC18(cf31) {
      let $dt = new D8(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC19(_dafny.Map.Empty, _dafny.Set.Empty, _dafny.MultiSet.Empty, _module.D1.Default());
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
    static create_DC21() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC22(cf37, cf38, cf39, cf40, cf41) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf42, cf43, cf44, cf45, cf46) {
      let $dt = new D9(2);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC20(cf36) {
      let $dt = new D9(3);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + this.cf41.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf36) + ")";
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
        return other.$tag === 1 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC21();
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
    static create_DC24(cf47) {
      let $dt = new D10(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf47) + ")";
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf49, cf50) {
      let $dt = new D11(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC27(cf51, cf52, cf53, cf54, cf55) {
      let $dt = new D11(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC25(cf48) {
      let $dt = new D11(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC28(cf56) {
      let $dt = new D11(3);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
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
        return "D11.DC26" + "(" + _dafny.toString(this.cf49) + ", " + this.cf50.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC26([], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC30(cf58, cf59) {
      let $dt = new D12(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC31(cf60, cf61, cf62, cf63, cf64) {
      let $dt = new D12(1);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC32(cf65, cf66, cf67, cf68, cf69) {
      let $dt = new D12(2);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC29(cf57) {
      let $dt = new D12(3);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get is_DC29() { return this.$tag === 3; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf57) + ")";
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
        return other.$tag === 1 && this.cf60 === other.cf60 && this.cf61 === other.cf61 && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63) && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66) && this.cf67 === other.cf67 && this.cf68 === other.cf68 && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC30([], false);
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
    static create_DC34(cf71, cf72) {
      let $dt = new D13(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC33(cf70) {
      let $dt = new D13(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC35(cf73) {
      let $dt = new D13(2);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71) && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf70 === other.cf70;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC34(_dafny.ZERO, _dafny.Seq.of());
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
    static create_DC37(cf75, cf76) {
      let $dt = new D14(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC36(cf74) {
      let $dt = new D14(1);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC37" + "(" + this.cf75.toVerbatimString(true) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC37(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC39(cf78, cf79, cf80, cf81, cf82) {
      let $dt = new D15(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC38(cf77) {
      let $dt = new D15(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + this.cf80.toVerbatimString(true) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79) && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC39(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC41(cf84) {
      let $dt = new D16(0);
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC42(cf85) {
      let $dt = new D16(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC40(cf83) {
      let $dt = new D16(2);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC43(cf86) {
      let $dt = new D16(3);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf84 === other.cf84;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf83 === other.cf83;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf86, other.cf86);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC41(false);
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
    static create_DC45(cf88) {
      let $dt = new D17(0);
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC44(cf87) {
      let $dt = new D17(1);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf87 === other.cf87;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC45(false);
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
    static create_DC47(cf90, cf91) {
      let $dt = new D18(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC46(cf89) {
      let $dt = new D18(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC46" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf90, other.cf90) && this.cf91 === other.cf91;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC47(_dafny.ZERO, []);
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
    static create_DC49(cf93) {
      let $dt = new D19(0);
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC50(cf94, cf95, cf96) {
      let $dt = new D19(1);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC51(cf97, cf98) {
      let $dt = new D19(2);
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC48(cf92) {
      let $dt = new D19(3);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC48() { return this.$tag === 3; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC49" + "(" + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC50" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf93, other.cf93);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf94 === other.cf94 && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf97, other.cf97) && this.cf98 === other.cf98;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf92 === other.cf92;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC49(_dafny.Set.Empty);
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
    static create_DC53(cf100, cf101, cf102, cf103, cf104) {
      let $dt = new D20(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      return $dt;
    }
    static create_DC52(cf99) {
      let $dt = new D20(1);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC53" + "(" + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf100 === other.cf100 && this.cf101 === other.cf101 && _dafny.areEqual(this.cf102, other.cf102) && this.cf103 === other.cf103 && this.cf104 === other.cf104;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC53(null, false, _dafny.ZERO, false, false);
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
    static create_DC55(cf106, cf107, cf108) {
      let $dt = new D21(0);
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      return $dt;
    }
    static create_DC56(cf109) {
      let $dt = new D21(1);
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC57(cf110, cf111) {
      let $dt = new D21(2);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC54(cf105) {
      let $dt = new D21(3);
      $dt.cf105 = cf105;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf105() { return this.cf105; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf105) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf106 === other.cf106 && _dafny.areEqual(this.cf107, other.cf107) && _dafny.areEqual(this.cf108, other.cf108);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf109 === other.cf109;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf110 === other.cf110 && _dafny.areEqual(this.cf111, other.cf111);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf105, other.cf105);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC55(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC59(cf113, cf114) {
      let $dt = new D22(0);
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      return $dt;
    }
    static create_DC60(cf115, cf116, cf117) {
      let $dt = new D22(1);
      $dt.cf115 = cf115;
      $dt.cf116 = cf116;
      $dt.cf117 = cf117;
      return $dt;
    }
    static create_DC58(cf112) {
      let $dt = new D22(2);
      $dt.cf112 = cf112;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get is_DC58() { return this.$tag === 2; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf112() { return this.cf112; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC59" + "(" + _dafny.toString(this.cf113) + ", " + _dafny.toString(this.cf114) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC60" + "(" + _dafny.toString(this.cf115) + ", " + _dafny.toString(this.cf116) + ", " + _dafny.toString(this.cf117) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC58" + "(" + _dafny.toString(this.cf112) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf113, other.cf113) && this.cf114 === other.cf114;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf115, other.cf115) && _dafny.areEqual(this.cf116, other.cf116) && this.cf117 === other.cf117;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf112, other.cf112);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC59(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
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

  $module.T3 = class T3 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.Seq.of();
      this.f7 = _dafny.ZERO;
      this._f0 = [];
      this._f1 = false;
      this._f3 = _dafny.Seq.UnicodeFromString("");
      this._f4 = false;
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = _dafny.MultiSet.Empty;
      this._f10 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = _module.D1.Default();
      this._f13 = false;
      this.f19 = false;
      this._f20 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f19, f20, f12, f13) {
      let _this = this;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm16(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((function () {
        let _coll28 = new _dafny.Set();
        for (const _compr_28 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bxjtdswbn"), _dafny.Seq.UnicodeFromString("x"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(560)), function (_341_i0) {
          return (_this).f20;
        }))).Elements) {
          let _342_v0 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bxjtdswbn"), _dafny.Seq.UnicodeFromString("x"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(560)), function (_341_i0) {
            return (_this).f20;
          })), _342_v0)) {
            _coll28.add(_342_v0);
          }
        }
        return _coll28;
      }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vlirfxd")))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ivnwymujb")).length)));
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = _module.D1.Default();
      this._f13 = false;
      this.f23 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f23, f12, f13) {
      let _this = this;
      (_this).f23 = f23;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm8(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), _this.f23, _this.f23, _this.f23)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f13),(_this).f13)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cug")).length)));
    };
    fm9(p0, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f13, (_this).f13, (_this).f13)).length)),new BigNumber(454))).length)).plus(new BigNumber(-762));
    };
    fm42(globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(new BigNumber(97), new BigNumber(693), _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(-83), new BigNumber(-824))) {
          let _343_v0 = _compr_29;
          if (((new BigNumber(-83)).isLessThanOrEqualTo(_343_v0)) && ((_343_v0).isLessThan(new BigNumber(-824)))) {
            _coll29.push([(_343_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("lx")).length))),(_this).f13]);
          }
        }
        return _coll29;
      }()).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f13)).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f13)).cardinality()))).length))), _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of (_dafny.Set.fromElements(new BigNumber(436), new BigNumber(724))).Elements) {
          let _344_v1 = _compr_30;
          if ((_dafny.Set.fromElements(new BigNumber(436), new BigNumber(724))).contains(_344_v1)) {
            _coll30.push([(_344_v1).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(931),new BigNumber(668))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC15(_dafny.Seq.of(new BigNumber(477))),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_this).f13)).length))).length))).length)]);
          }
        }
        return _coll30;
      }())).length), new BigNumber((_dafny.Seq.UnicodeFromString("iik")).length)), _module.__default.safeModuloInt(new BigNumber(-784), new BigNumber(-992)));
    };
    fm43(globalState) {
      let _this = this;
      let _source8 = _module.D3.create_DC6(new BigNumber(209), (_this).f13);
      if (_source8.is_DC6) {
        let _345___mcc_h0 = (_source8).cf12;
        let _346___mcc_h1 = (_source8).cf13;
        let _347_cf13 = _346___mcc_h1;
        let _348_cf12 = _345___mcc_h0;
        return _347_cf13;
      } else if (_source8.is_DC7) {
        return (_this).f13;
      } else if (_source8.is_DC8) {
        let _349___mcc_h2 = (_source8).cf14;
        let _350_cf14 = _349___mcc_h2;
        return (_this).f13;
      } else {
        let _351___mcc_h3 = (_source8).cf11;
        let _352_cf11 = _351___mcc_h3;
        return true;
      }
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f12 = _module.D1.Default();
      this._f16 = _module.D1.Default();
      this._f14 = [];
      this._f15 = _dafny.Set.Empty;
      this._f13 = false;
      this._f17 = false;
      this._f21 = _dafny.ZERO;
      this._f22 = false;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f21, f22, f16, f17, f14, f15, f12, f13) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(!((_this).f22),(_this).f21)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f22,new BigNumber((_dafny.Set.fromElements((_this).f17)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f21)));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return (_this).f21;
    };
    fm9(p0, globalState) {
      let _this = this;
      return (_this).f21;
    };
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _353_i0;
      _353_i0 = _dafny.ZERO;
      L2: {
        while ((_this).f22) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_353_i0)) {
              break L2;
            }
            _353_i0 = (_353_i0).plus(_dafny.ONE);
            let _354_v0;
            let _nw38 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
            _354_v0 = _nw38;
            let _355_v1;
            _355_v1 = _dafny.Seq.of((_this).f17);
            let _356_v2;
            _356_v2 = new _dafny.CodePoint('k'.codePointAt(0));
            let _357_v3;
            _357_v3 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), _356_v2, _356_v2);
            let _source9 = _module.D2.create_DC4(_354_v0, (_this).f22, _module.__default.fm0(_355_v1, false, (_this).f13, new BigNumber((_357_v3).cardinality()), globalState), (_this).f22, (_this).f14);
            if (_source9.is_DC4) {
              let _358___mcc_h0 = (_source9).cf6;
              let _359___mcc_h1 = (_source9).cf7;
              let _360___mcc_h2 = (_source9).cf8;
              let _361___mcc_h3 = (_source9).cf9;
              let _362___mcc_h4 = (_source9).cf10;
              let _363_cf10 = _362___mcc_h4;
              let _364_cf9 = _361___mcc_h3;
              let _365_cf8 = _360___mcc_h2;
              let _366_cf7 = _359___mcc_h1;
              let _367_cf6 = _358___mcc_h0;
              let _368_v4;
              let _nw39 = new _module.C0();
              _nw39.__ctor(_module.__default.fm30((_this).f21, globalState), new _dafny.CodePoint('p'.codePointAt(0)), _this.f12, _364_cf9);
              _368_v4 = _nw39;
              let _369_v5;
              _369_v5 = _dafny.MultiSet.fromElements((_this).f21, new BigNumber((_355_v1).length));
              let _370_v6;
              _370_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC11(_369_v5),_356_v2);
              let _371_v7;
              _371_v7 = _module.D6.create_DC11(_369_v5);
              let _372_v8;
              _372_v8 = _dafny.Seq.of((_this).f21);
              let _373_v9;
              _373_v9 = _dafny.Seq.of(_372_v8);
              let _374_v10;
              _374_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_372_v8, (_373_v9)[_module.__default.safeIndex(_365_cf8, new BigNumber((_373_v9).length))]),_365_cf8);
              let _rhs23 = _364_cf9;
              let _rhs24 = _365_cf8;
              let _rhs25 = (_368_v4).fm16(_366_cf7, new BigNumber(((_370_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_371_v7,(_368_v4).f20))).length), globalState);
              let _rhs26 = _368_v4;
              let _rhs27 = (((_374_v10).contains(_dafny.Seq.Concat(_372_v8, _372_v8))) ? ((_374_v10).get(_dafny.Seq.Concat(_372_v8, _372_v8))) : ((_this).f21));
              let _lhs21 = globalState;
              let _lhs22 = globalState;
              _366_cf7 = _rhs23;
              _lhs21.f7 = _rhs24;
              _lhs22.f7 = _rhs25;
              _368_v4 = _rhs26;
              r0 = _rhs27;
              let _375_v11;
              _375_v11 = _module.D2.create_DC3(_363_cf10);
              let _376_v12;
              _376_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f22) ? (true) : ((_this).f22)),_375_v11);
              _376_v12 = (_376_v12).update((_355_v1)[_module.__default.safeIndex((_this).f21, new BigNumber((_355_v1).length))], _375_v11);
              let _377_v13;
              _377_v13 = _dafny.Map.Empty.slice().updateUnsafe((_372_v8)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), function (_378_i1) {
                return (_dafny.ZERO).minus((_this).f21);
              })).length), new BigNumber((_372_v8).length))],_356_v2);
              let _379_v14;
              _379_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_356_v2);
              _377_v13 = (_377_v13).update((_dafny.ZERO).minus((_dafny.ZERO).minus(((_366_cf7) ? ((_this).f21) : (_365_cf8)))), (((_379_v14).contains((_this).f17)) ? ((_379_v14).get((_this).f17)) : ((_368_v4).f20)));
              let _380_v15;
              _380_v15 = _dafny.Seq.of((_368_v4).f20);
              _380_v15 = _380_v15;
            } else {
              let _381___mcc_h5 = (_source9).cf5;
              let _382_cf5 = _381___mcc_h5;
              let _index29 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_354_v0).length));
              (_354_v0)[_index29] = (_this).f21;
              let _index30 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_354_v0).length));
              (_354_v0)[_index30] = (_this).f21;
              let _383_v16;
              _383_v16 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
              let _rhs28 = ((_354_v0)[_module.__default.safeIndex(new BigNumber(639), new BigNumber((_354_v0).length))]).plus(new BigNumber(((_383_v16).Merge(_383_v16)).length));
              let _rhs29 = (_354_v0)[_module.__default.safeIndex(new BigNumber(639), new BigNumber((_354_v0).length))];
              let _lhs23 = globalState;
              let _lhs24 = globalState;
              _lhs23.f7 = _rhs28;
              _lhs24.f7 = _rhs29;
              let _384_v17;
              _384_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_356_v2);
              let _385_v18;
              _385_v18 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f21).multipliedBy(new BigNumber(((_this).f15).length)),_384_v17);
              _385_v18 = _385_v18;
              let _386_v19;
              let _out6;
              _out6 = _module.__default.m0(_356_v2, !(false), globalState);
              _386_v19 = _out6;
            }
            (globalState).f2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_355_v1, _355_v1), _355_v1);
            _356_v2 = _356_v2;
            let _387_v20;
            let _nw40 = new _module.C0();
            _nw40.__ctor((_this).f17, _356_v2, _this.f12, (_this).f17);
            _387_v20 = _nw40;
          }
        }
      }
      let _388_v21;
      _388_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(66),(_this).f13);
      let _389_v22;
      _389_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_388_v21).length),!(false));
      let _390_v23;
      _390_v23 = _module.D6.create_DC12((_this).f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), function (_391_i2) {
  return new _dafny.CodePoint('g'.codePointAt(0));
}), (_this).f21, p0);
      let _392_v24;
      _392_v24 = _dafny.Seq.UnicodeFromString("mpspf");
      let _pat_let_tv11 = p0;
      let _pat_let_tv12 = _392_v24;
      let _pat_let_tv13 = p0;
      let _pat_let_tv14 = _392_v24;
      let _393_v25;
      let _nw41 = Array((new BigNumber(2)).toNumber());
      _nw41[(_dafny.ZERO).toNumber()] = _388_v21;
      _nw41[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(function (_pat_let4_0) {
        return function (_398_dt__update__tmp_h1) {
          return function (_pat_let9_0) {
            return function (_399_dt__update_hcf20_h1) {
              return _module.D6.create_DC12((_398_dt__update__tmp_h1).dtor_cf18, (_398_dt__update__tmp_h1).dtor_cf19, _399_dt__update_hcf20_h1, (_398_dt__update__tmp_h1).dtor_cf21);
            }(_pat_let9_0);
          }((_this).f21);
        }(_pat_let4_0);
      }(function (_pat_let5_0) {
        return function (_394_dt__update__tmp_h0) {
          return function (_pat_let6_0) {
            return function (_395_dt__update_hcf20_h0) {
              return function (_pat_let7_0) {
                return function (_396_dt__update_hcf18_h0) {
                  return function (_pat_let8_0) {
                    return function (_397_dt__update_hcf19_h0) {
                      return _module.D6.create_DC12(_396_dt__update_hcf18_h0, _397_dt__update_hcf19_h0, _395_dt__update_hcf20_h0, (_394_dt__update__tmp_h0).dtor_cf21);
                    }(_pat_let8_0);
                  }(_pat_let_tv14);
                }(_pat_let7_0);
              }(false);
            }(_pat_let6_0);
          }((_module.D6.create_DC12(_pat_let_tv11, _pat_let_tv12, (_this).f21, !(_pat_let_tv13))).dtor_cf20);
        }(_pat_let5_0);
      }(_390_v23))).dtor_cf21);
      _393_v25 = _nw41;
      let _400_v26;
      _400_v26 = new _dafny.CodePoint('u'.codePointAt(0));
      let _401_v27;
      let _nw42 = new _module.C0();
      _nw42.__ctor((_this).f17, _400_v26, _this.f12, p0);
      _401_v27 = _nw42;
      let _402_v28;
      _402_v28 = _dafny.Seq.of((_this).f15, (_this).f15, (_this).f15);
      let _403_v29;
      _403_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0);
      let _404_v30;
      _404_v30 = _dafny.MultiSet.fromElements((_this).f21, new BigNumber((_403_v29).length));
      let _405_v31;
      _405_v31 = _module.D6.create_DC11(_404_v30);
      let _406_v32;
      _406_v32 = _dafny.Map.Empty.slice().updateUnsafe(_401_v27,(_dafny.ZERO).minus(new BigNumber(((_402_v28)[_module.__default.safeIndex(new BigNumber(((_405_v31).dtor_cf17).cardinality()), new BigNumber((_402_v28).length))]).length)));
      let _407_v33;
      _407_v33 = _dafny.Seq.of(_406_v32);
      let _408_v34;
      _408_v34 = _dafny.Seq.of((_this).f21, (_this).f21);
      let _409_v35;
      _409_v35 = _dafny.Seq.of(new BigNumber((_408_v34).length), (_this).f21);
      let _410_v36;
      _410_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_409_v35).length),_392_v24);
      let _rhs30 = _393_v25;
      let _rhs31 = ((_dafny.Map.Empty.slice().updateUnsafe(_401_v27,(_this).f21)).Merge((_407_v33)[_module.__default.safeIndex((_this).f21, new BigNumber((_407_v33).length))])).Merge((_406_v32).Merge(_406_v32));
      let _rhs32 = new BigNumber(((((_410_v36).contains((_this).f21)) ? ((_410_v36).get((_this).f21)) : (_dafny.Seq.UnicodeFromString("uxadk")))).length);
      let _lhs25 = globalState;
      _393_v25 = _rhs30;
      _406_v32 = _rhs31;
      _lhs25.f7 = _rhs32;
      (globalState).f7 = (_this).f21;
      _389_v22 = (_389_v22).update(_module.__default.safeDivisionInt((_this).fm9(_400_v26, globalState), (_this).f21), (_this).f17);
      let _411_v37;
      let _nw43 = new _module.C0();
      _nw43.__ctor(!((_this).f17), _400_v26, _this.f12, (_401_v27).f13);
      _411_v37 = _nw43;
      if ((_this).f22) {
        let _412_v38;
        _412_v38 = _module.D3.create_DC8((_this).f21);
        let _413_v39;
        _413_v39 = _dafny.Map.Empty.slice().updateUnsafe((_401_v27).f13,(_this).f21);
        let _414_v40;
        _414_v40 = _dafny.Map.Empty.slice().updateUnsafe((_401_v27).f13,_400_v26);
        let _415_v41;
        _415_v41 = _dafny.MultiSet.fromElements((_this).f22, (_401_v27).f13);
        let _416_v43;
        _416_v43 = _dafny.Map.Empty.slice().updateUnsafe((_411_v37).f20,(_this).f21);
        let _417_v44;
        _417_v44 = _dafny.Seq.of(function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of (_416_v43).Keys.Elements) {
            let _418_v42 = _compr_31;
            if ((_416_v43).contains(_418_v42)) {
              _coll31.push([_418_v42,(_this).f21]);
            }
          }
          return _coll31;
        }());
        let _419_v45;
        _419_v45 = _dafny.Set.fromElements(false, (_401_v27).f13);
        let _420_v46;
        _420_v46 = _dafny.Map.Empty.slice().updateUnsafe(_403_v29,_419_v45);
        let _421_v47;
        _421_v47 = _dafny.Seq.of(_411_v37.f19);
        let _422_v48;
        let _nw44 = Array((new BigNumber(23)).toNumber());
        _nw44[(_dafny.ZERO).toNumber()] = (function (_pat_let10_0) {
          return function (_423_dt__update__tmp_h2) {
            return function (_pat_let11_0) {
              return function (_424_dt__update_hcf14_h0) {
                return _module.D3.create_DC8(_424_dt__update_hcf14_h0);
              }(_pat_let11_0);
            }(new BigNumber(381));
          }(_pat_let10_0);
        }(_412_v38)).dtor_cf14;
        _nw44[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus((_this).f21)).plus((_this).f21);
        _nw44[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_408_v34, _module.__default.safeIndex((_this).f21, new BigNumber((_408_v34).length)), (_this).f21), _module.__default.safeIndex((_this).f21, new BigNumber((_dafny.Seq.update(_408_v34, _module.__default.safeIndex((_this).f21, new BigNumber((_408_v34).length)), (_this).f21)).length)), (_this).f21)).length);
        _nw44[(new BigNumber(3)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f21);
        _nw44[(new BigNumber(5)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_this).f21, (_this).f21);
        _nw44[(new BigNumber(7)).toNumber()] = (((_413_v39).contains((_401_v27).f13)) ? ((_413_v39).get((_401_v27).f13)) : ((_this).f21));
        _nw44[(new BigNumber(8)).toNumber()] = new BigNumber(670);
        _nw44[(new BigNumber(9)).toNumber()] = new BigNumber((_392_v24).length);
        _nw44[(new BigNumber(10)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(11)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(12)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(13)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_392_v24,_414_v40)).Merge(_module.__default.fm31(globalState))).length);
        _nw44[(new BigNumber(14)).toNumber()] = (_this).f21;
        _nw44[(new BigNumber(15)).toNumber()] = (new BigNumber((_415_v41).cardinality())).multipliedBy(new BigNumber((_417_v44).length));
        _nw44[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt((_this).f21, (_this).f21);
        _nw44[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_392_v24).length), (_dafny.ZERO).minus((_this).f21));
        _nw44[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt((_this).f21, (_this).f21);
        _nw44[(new BigNumber(19)).toNumber()] = (new BigNumber((_dafny.Seq.of((((_420_v46).contains(_403_v29)) ? ((_420_v46).get(_403_v29)) : (_dafny.Set.fromElements(_411_v37.f19))))).length)).multipliedBy((_this).f21);
        _nw44[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt((_module.D3.create_DC6((_dafny.ZERO).minus((_this).f21), (_401_v27).f13)).dtor_cf12, new BigNumber(703));
        _nw44[(new BigNumber(21)).toNumber()] = new BigNumber(((_this).f15).length);
        _nw44[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_421_v47, _421_v47)).length);
        _422_v48 = _nw44;
        let _425_v49;
        _425_v49 = _dafny.Map.Empty.slice().updateUnsafe(_400_v26,_411_v37.f19);
        let _index31 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_422_v48).length));
        (_422_v48)[_index31] = _module.__default.safeDivisionInt(new BigNumber((_425_v49).length), (_this).f21);
        let _index32 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_422_v48).length));
        (_422_v48)[_index32] = (_dafny.ZERO).minus((_this).f21);
        let _426_v50;
        let _nw45 = new _module.C0();
        _nw45.__ctor((_401_v27).f13, (_411_v37).f20, _401_v27.f12, (_this).f17);
        _426_v50 = _nw45;
        let _427_v51;
        let _nw46 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _427_v51 = _nw46;
        let _index33 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_427_v51).length));
        (_427_v51)[_index33] = _392_v24;
        let _index34 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_427_v51).length));
        (_427_v51)[_index34] = _dafny.Seq.Concat(_392_v24, ((false) ? ((((_410_v36).contains((_422_v48)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_422_v48).length))])) ? ((_410_v36).get((_422_v48)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_422_v48).length))])) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(611)), ((_428_v37) => function (_429_i3) {
          return (_428_v37).f20;
        })(_411_v37))))) : (_392_v24)));
        _422_v48 = _422_v48;
        (globalState).f7 = (_422_v48)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_422_v48).length))];
      } else {
        let _430_v52;
        _430_v52 = _dafny.MultiSet.fromElements((_402_v28)[_module.__default.safeIndex((_this).f21, new BigNumber((_402_v28).length))]);
        let _431_v53;
        _431_v53 = _dafny.Seq.of((_401_v27).f13, (_this).f22);
        let _432_v54;
        _432_v54 = _dafny.Map.Empty.slice().updateUnsafe((_401_v27).f13,_430_v52);
        _430_v52 = (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_431_v53).length)))).Intersect((((_432_v54).contains(!(p0))) ? ((_432_v54).get(!(p0))) : (_430_v52)));
        let _433_v56;
        _433_v56 = _dafny.Map.Empty.slice().updateUnsafe((function () {
          let _coll32 = new _dafny.Set();
          for (const _compr_32 of (_dafny.MultiSet.fromElements((_this).f21)).Elements) {
            let _434_v55 = _compr_32;
            if ((_dafny.MultiSet.fromElements((_this).f21)).contains(_434_v55)) {
              _coll32.add(_module.__default.safeModuloInt(_434_v55, new BigNumber(781)));
            }
          }
          return _coll32;
        }()).Union((_this).f15),_dafny.Map.Empty.slice().updateUnsafe(_392_v24,(_411_v37).f20));
        let _435_v57;
        _435_v57 = _dafny.Map.Empty.slice().updateUnsafe(_392_v24,_400_v26);
        _433_v56 = (_433_v56).update((_this).f15, _435_v57);
        let _436_v58;
        let _nw47 = Array((new BigNumber(19)).toNumber()).fill([]);
        _436_v58 = _nw47;
        let _437_v59;
        _437_v59 = _dafny.Map.Empty.slice().updateUnsafe(_431_v53,(_this).f21);
        let _438_v60;
        _438_v60 = _module.D7.create_DC15(_dafny.Seq.of((((_437_v59).contains(_431_v53)) ? ((_437_v59).get(_431_v53)) : ((_this).f21))));
        let _439_v61;
        let _nw48 = Array((new BigNumber(16)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _408_v34;
        _nw48[(_dafny.ONE).toNumber()] = _408_v34;
        _nw48[(new BigNumber(2)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(3)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(389)), ((_440_v37) => function (_441_i4) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_440_v37).f20,(_440_v37).f20),true)).length);
        })(_411_v37));
        _nw48[(new BigNumber(5)).toNumber()] = _408_v34;
        _nw48[(new BigNumber(6)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(7)).toNumber()] = (_438_v60).dtor_cf28;
        _nw48[(new BigNumber(8)).toNumber()] = _408_v34;
        _nw48[(new BigNumber(9)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(10)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(11)).toNumber()] = _409_v35;
        _nw48[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(954)), function (_442_i5) {
          return (_this).f21;
        });
        _nw48[(new BigNumber(13)).toNumber()] = _408_v34;
        _nw48[(new BigNumber(14)).toNumber()] = _408_v34;
        _nw48[(new BigNumber(15)).toNumber()] = _408_v34;
        _439_v61 = _nw48;
        let _index35 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_436_v58).length));
        (_436_v58)[_index35] = _439_v61;
        let _index36 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_436_v58).length));
        (_436_v58)[_index36] = _439_v61;
        (_411_v37).f19 = ((((_this).f17) ? (new BigNumber((_392_v24).length)) : ((_this).f21))).isLessThanOrEqualTo((((_404_v30).contains((_this).f21)) ? ((_404_v30).get((_this).f21)) : ((_this).f21)));
        if (!(_dafny.areEqual((((_this.f12).dtor_cf4) ? (_392_v24) : (_dafny.Seq.UnicodeFromString("agacmtia"))), _392_v24))) {
          let _443_v62;
          _443_v62 = _dafny.Seq.of((_this).f14);
          let _444_v63;
          _444_v63 = _dafny.Set.fromElements(false, _411_v37.f19, true);
          let _445_v64;
          _445_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_443_v62)[_module.__default.safeIndex(new BigNumber((_444_v63).length), new BigNumber((_443_v62).length))]);
          let _446_v65;
          _446_v65 = _module.D6.create_DC11(_404_v30);
          let _447_v66;
          _447_v66 = _module.D6.create_DC14(_446_v65);
          let _448_v67;
          let _nw49 = Array((new BigNumber(20)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = (_this).f14;
          _nw49[(_dafny.ONE).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(2)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(3)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(4)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(5)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(7)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(8)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(9)).toNumber()] = (((_445_v64).contains((_this).f21)) ? ((_445_v64).get((_this).f21)) : ((_this).f14));
          _nw49[(new BigNumber(10)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(11)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(13)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(14)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(15)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(16)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(17)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(18)).toNumber()] = (_this).f14;
          _nw49[(new BigNumber(19)).toNumber()] = (_443_v62)[_module.__default.safeIndex(new BigNumber((_module.__default.fm32(_447_v66, globalState)).length), new BigNumber((_443_v62).length))];
          _448_v67 = _nw49;
          let _449_v68;
          _449_v68 = _448_v67;
          _449_v68 = _449_v68;
          (_411_v37).f19 = false;
          (globalState).f7 = ((_this).f21).plus((_this).f21);
          let _450_v69;
          let _init3 = function (_451_i6) {
            return _module.__default.safeDivisionInt(_451_i6, (_this).f21);
          };
          let _nw50 = Array((new BigNumber(17)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw50.length); _i0_3++) {
            _nw50[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _450_v69 = _nw50;
          _450_v69 = (((_this).f22) ? (_450_v69) : (_450_v69));
          let _452_v70;
          _452_v70 = _dafny.Seq.of(_390_v23);
          let _453_v71;
          _453_v71 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(541)), ((_454_v23) => function (_455_i7) {
            return _454_v23;
          })(_390_v23)), _452_v70);
          let _456_v72;
          _456_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_411_v37.f19)).length),(_453_v71)[_module.__default.safeIndex((_this).f21, new BigNumber((_453_v71).length))]);
          let _457_v73;
          let _nw51 = new _module.C0();
          _nw51.__ctor((((_this).f17) ? ((_401_v27).f13) : ((_401_v27).f13)), _400_v26, _401_v27.f12, !(_456_v72).contains(new BigNumber(693)));
          _457_v73 = _nw51;
        } else {
          let _index37 = _module.__default.safeIndex(new BigNumber(376), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index37] = (_401_v27).f13;
          let _458_v74;
          _458_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          let _459_v75;
          _459_v75 = _dafny.Set.fromElements(_411_v37.f19, p0, (_this).f17, false);
          let _460_v76;
          _460_v76 = _dafny.Seq.of(_392_v24);
          let _index38 = _module.__default.safeIndex(new BigNumber(376), new BigNumber(((_this).f14).length));
          let _rhs33 = (_401_v27).f13;
          let _rhs34 = (_459_v75).equals((_459_v75).Union(_459_v75));
          let _rhs35 = _458_v74;
          let _rhs36 = (((_this).f22) ? ((_431_v53)[_module.__default.safeIndex((_this).f21, new BigNumber((_431_v53).length))]) : (!((new BigNumber((_460_v76).length)).isLessThanOrEqualTo((_this).f21))));
          let _lhs26 = _411_v37;
          let _lhs27 = (_this).f14;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(376), new BigNumber(((_this).f14).length));
          let _lhs29 = _411_v37;
          _lhs26.f19 = _rhs33;
          _lhs27[_lhs28] = _rhs34;
          _458_v74 = _rhs35;
          _lhs29.f19 = _rhs36;
          _392_v24 = _392_v24;
          let _461_v77;
          let _init4 = function (_462_i8) {
            return (_462_i8).minus((_this).f21);
          };
          let _nw52 = Array((new BigNumber(21)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw52.length); _i0_4++) {
            _nw52[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _461_v77 = _nw52;
          _461_v77 = _461_v77;
          let _463_v78;
          let _init5 = ((_464_v37) => function (_465_i9) {
            return (_464_v37).f20;
          })(_411_v37);
          let _nw53 = Array((new BigNumber(5)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw53.length); _i0_5++) {
            _nw53[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _463_v78 = _nw53;
          let _index39 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_463_v78).length));
          (_463_v78)[_index39] = (_411_v37).f20;
          let _index40 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_463_v78).length));
          let _rhs37 = (_411_v37).f20;
          let _rhs38 = _431_v53;
          let _rhs39 = (_411_v37).f20;
          let _lhs30 = _463_v78;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_463_v78).length));
          let _lhs32 = globalState;
          _lhs30[_lhs31] = _rhs37;
          _lhs32.f2 = _rhs38;
          _400_v26 = _rhs39;
          (globalState).f7 = (_dafny.ZERO).minus((_this).f21);
        }
      }
      r0 = new BigNumber(139);
      return r0;
    }
    m7(globalState) {
      let _this = this;
      let _466_v0;
      _466_v0 = _dafny.Seq.UnicodeFromString("rxto");
      let _467_v1;
      let _init6 = function (_468_i0) {
        return (_468_i0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), function (_469_i1) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })).length));
      };
      let _nw54 = Array((new BigNumber(9)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw54.length); _i0_6++) {
        _nw54[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _467_v1 = _nw54;
      let _470_v2;
      _470_v2 = _dafny.Seq.of(_467_v1, _467_v1, _467_v1);
      let _471_v3;
      _471_v3 = _module.D6.create_DC12((_this).f17, _466_v0, new BigNumber((_470_v2).length), (_this).f22);
      let _source10 = _471_v3;
      if (_source10.is_DC12) {
        let _472___mcc_h0 = (_source10).cf18;
        let _473___mcc_h1 = (_source10).cf19;
        let _474___mcc_h2 = (_source10).cf20;
        let _475___mcc_h3 = (_source10).cf21;
        let _476_cf21 = _475___mcc_h3;
        let _477_cf20 = _474___mcc_h2;
        let _478_cf19 = _473___mcc_h1;
        let _479_cf18 = _472___mcc_h0;
        let _index41 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length));
        (_467_v1)[_index41] = (_this).f21;
        let _index42 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length));
        (_467_v1)[_index42] = _477_cf20;
        let _480_v4;
        _480_v4 = _dafny.Map.Empty.slice().updateUnsafe(_477_cf20,(_this).f14);
        let _481_v5;
        _481_v5 = _dafny.Map.Empty.slice().updateUnsafe(_476_cf21,new BigNumber((_480_v4).length));
        let _482_v6;
        _482_v6 = _dafny.Map.Empty.slice().updateUnsafe(_481_v5,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(877))).cardinality()));
        let _483_v7;
        _483_v7 = _dafny.Map.Empty.slice().updateUnsafe(_479_cf18,false);
        let _484_v8;
        _484_v8 = _dafny.Seq.of((_this).f21, (_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))], _477_cf20, (_this).f21, _477_cf20);
        let _485_v9;
        _485_v9 = _dafny.MultiSet.fromElements(_477_cf20);
        let _486_v10;
        let _nw55 = Array((new BigNumber(18)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))];
        _nw55[(_dafny.ONE).toNumber()] = new BigNumber((_478_cf19).length);
        _nw55[(new BigNumber(2)).toNumber()] = ((true) ? ((_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))]) : ((((_482_v6).contains(_481_v5)) ? ((_482_v6).get(_481_v5)) : ((_this).f21))));
        _nw55[(new BigNumber(3)).toNumber()] = ((_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))]).plus(new BigNumber((_483_v7).length));
        _nw55[(new BigNumber(4)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))];
        _nw55[(new BigNumber(5)).toNumber()] = _477_cf20;
        _nw55[(new BigNumber(6)).toNumber()] = (_477_cf20).minus((_this).f21);
        _nw55[(new BigNumber(7)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))];
        _nw55[(new BigNumber(8)).toNumber()] = (_this).f21;
        _nw55[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_484_v8, _module.__default.safeIndex((_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))], new BigNumber((_484_v8).length)), _477_cf20), _484_v8)).length);
        _nw55[(new BigNumber(10)).toNumber()] = (_477_cf20).multipliedBy(new BigNumber((_dafny.Seq.update(_484_v8, _module.__default.safeIndex(_477_cf20, new BigNumber((_484_v8).length)), _477_cf20)).length));
        _nw55[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(_477_cf20, _477_cf20);
        _nw55[(new BigNumber(12)).toNumber()] = ((true) ? (new BigNumber(547)) : (_477_cf20));
        _nw55[(new BigNumber(13)).toNumber()] = new BigNumber((_485_v9).cardinality());
        _nw55[(new BigNumber(14)).toNumber()] = (_484_v8)[_module.__default.safeIndex(_477_cf20, new BigNumber((_484_v8).length))];
        _nw55[(new BigNumber(15)).toNumber()] = new BigNumber(865);
        _nw55[(new BigNumber(16)).toNumber()] = _477_cf20;
        _nw55[(new BigNumber(17)).toNumber()] = (_this).f21;
        _486_v10 = _nw55;
        _486_v10 = (_module.D2.create_DC4(_486_v10, true, (_this).f21, (_this).f22, (_this).f14)).dtor_cf6;
        (globalState).f7 = new BigNumber(527);
        let _487_v11;
        _487_v11 = _dafny.Seq.of(!((_this).f13), _479_cf18);
        let _488_v12;
        let _nw56 = new _module.C0();
        _nw56.__ctor((((_this).f22) ? ((_487_v11)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f21), new BigNumber((_487_v11).length))]) : (false)), (_466_v0)[_module.__default.safeIndex((_467_v1)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_467_v1).length))], new BigNumber((_466_v0).length))], _this.f12, _479_cf18);
        _488_v12 = _nw56;
      } else if (_source10.is_DC13) {
        let _489___mcc_h4 = (_source10).cf22;
        let _490___mcc_h5 = (_source10).cf23;
        let _491___mcc_h6 = (_source10).cf24;
        let _492___mcc_h7 = (_source10).cf25;
        let _493___mcc_h8 = (_source10).cf26;
        let _494_cf26 = _493___mcc_h8;
        let _495_cf25 = _492___mcc_h7;
        let _496_cf24 = _491___mcc_h6;
        let _497_cf23 = _490___mcc_h5;
        let _498_cf22 = _489___mcc_h4;
        _497_cf23 = (_this).f22;
        let _499_v13;
        _499_v13 = _dafny.Set.fromElements((_this).f17);
        let _500_v14;
        let _out7;
        _out7 = (_this).m13((_this).f13, _module.__default.fm33((_this).f13, _495_cf25, globalState), _499_v13, (_dafny.ZERO).minus(_496_cf24), globalState);
        _500_v14 = _out7;
        let _index43 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_467_v1).length));
        (_467_v1)[_index43] = _496_cf24;
        let _501_v15;
        _501_v15 = _dafny.Seq.of((_this).f17);
        let _index44 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_467_v1).length));
        (_467_v1)[_index44] = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!((((_501_v15)[_module.__default.safeIndex((_this).f21, new BigNumber((_501_v15).length))]) ? ((_this).f22) : (!(_497_cf23)))), (_495_cf25).isEqualTo((_this).f21), (_494_cf26) || (_500_v14))).length));
        _500_v14 = (_497_cf23) === ((_this).f13);
      } else if (_source10.is_DC11) {
        let _502___mcc_h9 = (_source10).cf17;
        let _503_cf17 = _502___mcc_h9;
        let _504_v16;
        _504_v16 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21);
        (globalState).f7 = (((_504_v16).contains((_this).f22)) ? ((_504_v16).get((_this).f22)) : ((_this).f21));
        let _505_v17;
        _505_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_this.f12);
        let _506_v18;
        _506_v18 = new _dafny.CodePoint('c'.codePointAt(0));
        let _507_v19;
        let _nw57 = new _module.C0();
        _nw57.__ctor(!(_505_v17).equals(_505_v17), _506_v18, _module.D1.create_DC2((_this).f21, (_this).f21, true), (_this).f17);
        _507_v19 = _nw57;
        let _508_v20;
        _508_v20 = _module.D6.create_DC11(_503_cf17);
        let _source11 = _508_v20;
        if (_source11.is_DC12) {
          let _509___mcc_h11 = (_source11).cf18;
          let _510___mcc_h12 = (_source11).cf19;
          let _511___mcc_h13 = (_source11).cf20;
          let _512___mcc_h14 = (_source11).cf21;
          let _513_cf21 = _512___mcc_h14;
          let _514_cf20 = _511___mcc_h13;
          let _515_cf19 = _510___mcc_h12;
          let _516_cf18 = _509___mcc_h11;
          let _517_v21;
          let _init7 = function (_518_i2) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          };
          let _nw58 = Array((new BigNumber(5)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw58.length); _i0_7++) {
            _nw58[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _517_v21 = _nw58;
          let _index45 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_517_v21).length));
          (_517_v21)[_index45] = _506_v18;
          let _index46 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_517_v21).length));
          (_517_v21)[_index46] = new _dafny.CodePoint('v'.codePointAt(0));
          let _519_v22;
          _519_v22 = _dafny.MultiSet.fromElements((_this).f13, _516_cf18);
          let _520_v23;
          _520_v23 = _dafny.Seq.of(_516_cf18, _513_cf21);
          let _521_v24;
          _521_v24 = _dafny.Seq.of(_514_cf20, new BigNumber((_466_v0).length));
          let _522_v25;
          _522_v25 = _module.D6.create_DC14(_module.D6.create_DC11(_dafny.MultiSet.FromArray(_521_v24)));
          let _523_v26;
          _523_v26 = _module.D6.create_DC14(_522_v25);
          let _524_v27;
          _524_v27 = _dafny.Map.Empty.slice().updateUnsafe(((((_519_v22).contains((_this).f17)) ? ((_519_v22).get((_this).f17)) : (_module.__default.fm0(_520_v23, (_this).f13, (_this).f22, _514_cf20, globalState)))).plus(new BigNumber((_521_v24).length)),(new BigNumber((_dafny.Seq.update(_module.__default.fm34(_514_cf20, globalState), _module.__default.safeIndex((_this).f21, new BigNumber((_module.__default.fm34(_514_cf20, globalState)).length)), (_this).f21)).length)).minus(new BigNumber((_module.__default.fm32(_523_v26, globalState)).length)));
          _524_v27 = (_524_v27).update(_514_cf20, new BigNumber((_515_cf19).length));
          let _525_v28;
          let _nw59 = Array((new BigNumber(12)).toNumber()).fill(_module.D7.Default());
          _525_v28 = _nw59;
          let _526_v29;
          _526_v29 = _module.D7.create_DC16(_514_cf20);
          let _527_v30;
          _527_v30 = _module.D7.create_DC17(_526_v29);
          let _528_v31;
          _528_v31 = _module.D7.create_DC17(_527_v30);
          let _index47 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_525_v28).length));
          (_525_v28)[_index47] = _528_v31;
          let _529_v32;
          _529_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_515_cf19);
          let _530_v33;
          _530_v33 = _dafny.Seq.of(_466_v0, _466_v0, _515_cf19, _dafny.Seq.Concat(_466_v0, _515_cf19), _466_v0);
          let _index48 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_525_v28).length));
          let _rhs40 = _514_cf20;
          let _rhs41 = new BigNumber(((_530_v33)[_module.__default.safeIndex((_521_v24)[_module.__default.safeIndex(new BigNumber(391), new BigNumber((_521_v24).length))], new BigNumber((_530_v33).length))]).length);
          let _rhs42 = _528_v31;
          let _rhs43 = _module.__default.fm35((_471_v3).dtor_cf20, ((_516_cf18) ? (new BigNumber((_520_v23).length)) : (_514_cf20)), (_507_v19).f20, !(_module.__default.fm30((_dafny.ZERO).minus((_this).f21), globalState)), globalState);
          let _lhs33 = globalState;
          let _lhs34 = globalState;
          let _lhs35 = _525_v28;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_525_v28).length));
          _lhs33.f7 = _rhs40;
          _lhs34.f7 = _rhs41;
          _lhs35[_lhs36] = _rhs42;
          _529_v32 = _rhs43;
          _523_v26 = _523_v26;
        } else if (_source11.is_DC13) {
          let _531___mcc_h15 = (_source11).cf22;
          let _532___mcc_h16 = (_source11).cf23;
          let _533___mcc_h17 = (_source11).cf24;
          let _534___mcc_h18 = (_source11).cf25;
          let _535___mcc_h19 = (_source11).cf26;
          let _536_cf26 = _535___mcc_h19;
          let _537_cf25 = _534___mcc_h18;
          let _538_cf24 = _533___mcc_h17;
          let _539_cf23 = _532___mcc_h16;
          let _540_cf22 = _531___mcc_h15;
          let _541_v34;
          let _init8 = ((_542_cf23, _543_cf25) => function (_544_i3) {
            return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_542_cf23,true)).length)), _dafny.Seq.of((_this).f21, (_this).f21, _543_cf25));
          })(_539_cf23, _537_cf25);
          let _nw60 = Array((new BigNumber(14)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw60.length); _i0_8++) {
            _nw60[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _541_v34 = _nw60;
          let _545_v35;
          _545_v35 = _dafny.Seq.of((_this).f21);
          let _index49 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_541_v34).length));
          (_541_v34)[_index49] = _545_v35;
          let _index50 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_541_v34).length));
          (_541_v34)[_index50] = _545_v35;
          let _546_v36;
          _546_v36 = _dafny.Map.Empty.slice().updateUnsafe(_537_cf25,!((_this).f17));
          let _547_v37;
          _547_v37 = _dafny.Seq.of((_this).f13, (((_546_v36).contains(_538_cf24)) ? ((_546_v36).get(_538_cf24)) : (_507_v19.f19)), _536_cf26, (_this).f22, _536_cf26);
          (_507_v19).f19 = (new BigNumber(((_this).fm10(_539_cf23, _536_cf26, _547_v37, globalState)).length)).isLessThan(_538_cf24);
          _539_cf23 = ((_this).f21).isEqualTo(_537_cf25);
          let _548_v38;
          let _nw61 = Array((new BigNumber(29)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = (_this).f14;
          _nw61[(_dafny.ONE).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(2)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(3)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(4)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(5)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(7)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(8)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(9)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(10)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(11)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(12)).toNumber()] = (_module.D2.create_DC4(_467_v1, (_this).f22, new BigNumber((_466_v0).length), _507_v19.f19, (_this).f14)).dtor_cf10;
          _nw61[(new BigNumber(13)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(14)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(15)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(16)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(17)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(18)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(19)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(20)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(21)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(22)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(23)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(24)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(25)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(26)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(27)).toNumber()] = (_this).f14;
          _nw61[(new BigNumber(28)).toNumber()] = (_this).f14;
          _548_v38 = _nw61;
          let _549_v39;
          _549_v39 = _548_v38;
          let _550_v40;
          let _nw62 = Array((new BigNumber(27)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _549_v39;
          _nw62[(_dafny.ONE).toNumber()] = _549_v39;
          _nw62[(new BigNumber(2)).toNumber()] = _548_v38;
          _nw62[(new BigNumber(3)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(4)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(5)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(6)).toNumber()] = _548_v38;
          _nw62[(new BigNumber(7)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(8)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(9)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(10)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(11)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(12)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(13)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(14)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(15)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(16)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(17)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(18)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(19)).toNumber()] = _548_v38;
          _nw62[(new BigNumber(20)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(21)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(22)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(23)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(24)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(25)).toNumber()] = _549_v39;
          _nw62[(new BigNumber(26)).toNumber()] = _549_v39;
          _550_v40 = _nw62;
          let _index51 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_550_v40).length));
          (_550_v40)[_index51] = _549_v39;
          let _index52 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_550_v40).length));
          (_550_v40)[_index52] = _548_v38;
        } else if (_source11.is_DC11) {
          let _551___mcc_h20 = (_source11).cf17;
          let _552_cf17 = _551___mcc_h20;
          let _index53 = _module.__default.safeIndex(new BigNumber(312), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index53] = (_this).f22;
          let _index54 = _module.__default.safeIndex(new BigNumber(312), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index54] = _507_v19.f19;
          (globalState).f7 = (_this).f21;
          let _553_v41;
          _553_v41 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_466_v0, _466_v0), _466_v0, _466_v0, _466_v0, _466_v0);
          let _554_v42;
          _554_v42 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(658)).isLessThanOrEqualTo((_this).f21),_467_v1);
          let _555_v43;
          _555_v43 = _dafny.Set.fromElements((_this).f17);
          let _556_v44;
          _556_v44 = _dafny.Seq.of(_dafny.Set.fromElements(false), _555_v43);
          let _557_v45;
          _557_v45 = _dafny.MultiSet.fromElements(((_556_v44)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f21)).length), new BigNumber((_556_v44).length))]).IsDisjointFrom(_555_v43), _507_v19.f19, ((_this).f15).IsSubsetOf((_this).f15), _507_v19.f19, _module.__default.fm30((_this).f21, globalState));
          let _558_v46;
          _558_v46 = _dafny.Seq.of(_557_v45, _557_v45, (_557_v45).Difference(_557_v45));
          let _index55 = _module.__default.safeIndex(new BigNumber(312), new BigNumber(((_this).f14).length));
          let _rhs44 = false;
          let _rhs45 = (((_507_v19.f19) ? (_553_v41) : ((_553_v41).update(_466_v0, _module.__default.abs(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f21))).cardinality())))))).update((_this.f16).dtor_cf1, _module.__default.abs((_this).f21));
          let _rhs46 = _554_v42;
          let _rhs47 = (_558_v46)[_module.__default.safeIndex((_this).f21, new BigNumber((_558_v46).length))];
          let _rhs48 = !(!(function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(391), new BigNumber(894))) {
              let _559_v47 = _compr_33;
              if (((new BigNumber(391)).isLessThanOrEqualTo(_559_v47)) && ((_559_v47).isLessThan(new BigNumber(894)))) {
                _coll33.push([(_559_v47).multipliedBy((_this).f21),_557_v45]);
              }
            }
            return _coll33;
          }()).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,_557_v45))) || ((_this).f13);
          let _lhs37 = (_this).f14;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(312), new BigNumber(((_this).f14).length));
          let _lhs39 = _507_v19;
          _lhs37[_lhs38] = _rhs44;
          _553_v41 = _rhs45;
          _554_v42 = _rhs46;
          _557_v45 = _rhs47;
          _lhs39.f19 = _rhs48;
          let _560_v48;
          _560_v48 = _dafny.Seq.of((_this).f17, (_this).f13);
          let _index56 = _module.__default.safeIndex(new BigNumber(312), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index56] = ((_module.__default.fm36((_560_v48)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_560_v48).length))], (_507_v19).f20, _module.__default.fm30((_this).f21, globalState), globalState)).equals(_555_v43)) && (true);
        } else {
          let _561___mcc_h21 = (_source11).cf27;
          let _562_cf27 = _561___mcc_h21;
          let _563_v49;
          let _nw63 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
          _563_v49 = _nw63;
          let _rhs49 = _563_v49;
          let _rhs50 = _dafny.Seq.UnicodeFromString("mxwbndl");
          let _rhs51 = _507_v19.f19;
          let _lhs40 = _507_v19;
          _563_v49 = _rhs49;
          _466_v0 = _rhs50;
          _lhs40.f19 = _rhs51;
          let _564_v50;
          _564_v50 = _dafny.Seq.of((_this).f21);
          let _565_v51;
          _565_v51 = _dafny.Map.Empty.slice().updateUnsafe(_466_v0,new BigNumber((_dafny.Seq.Concat(_564_v50, _564_v50)).length));
          _565_v51 = (_565_v51).update(_466_v0, (_this).f21);
          (_507_v19).f19 = _507_v19.f19;
          let _566_v52;
          let _nw64 = new _module.C0();
          _nw64.__ctor((_this).f22, (_507_v19).f20, _module.__default.fm37((_this).f21, (_this).f21, (_this).f21, (_507_v19).fm16((_this).f22, (_this).f21, globalState), globalState), _507_v19.f19);
          _566_v52 = _nw64;
          let _567_v53;
          _567_v53 = _dafny.MultiSet.fromElements(_566_v52, _566_v52, _566_v52);
          _567_v53 = (_567_v53).Intersect(_567_v53);
        }
        let _568_v54;
        let _init9 = ((_569_v18) => function (_570_i4) {
          return _569_v18;
        })(_506_v18);
        let _nw65 = Array((new BigNumber(14)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw65.length); _i0_9++) {
          _nw65[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _568_v54 = _nw65;
        let _571_v55;
        _571_v55 = _dafny.MultiSet.fromElements(_568_v54);
        let _572_v56;
        _572_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_571_v55);
        let _573_v57;
        _573_v57 = _dafny.MultiSet.fromElements((_this).f13, !((_this).f13));
        if (!(_573_v57).contains(((((_572_v56).contains((_this).f14)) ? ((_572_v56).get((_this).f14)) : (_571_v55))).contains(_568_v54))) {
          let _574_v58;
          let _nw66 = new _module.C0();
          _nw66.__ctor(!(_507_v19.f19) || ((_this).f13), (_507_v19).f20, _module.D1.create_DC2((_this).f21, new BigNumber(609), _module.__default.fm30((_this).f21, globalState)), _507_v19.f19);
          _574_v58 = _nw66;
          _506_v18 = (_507_v19).f20;
          _466_v0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("avvykq"), _dafny.Seq.Concat(_466_v0, _466_v0));
          let _575_v59;
          _575_v59 = _dafny.Seq.of((_this).f21);
          let _576_v60;
          _576_v60 = _dafny.Map.Empty.slice().updateUnsafe(_466_v0,_575_v59);
          let _577_v62;
          _577_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of (_576_v60).Keys.Elements) {
              let _578_v61 = _compr_34;
              if ((_576_v60).contains(_578_v61)) {
                _coll34.add(_578_v61);
              }
            }
            return _coll34;
          }());
          let _579_v63;
          _579_v63 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("mbugy"), _466_v0, _466_v0, _dafny.Seq.UnicodeFromString("fw"), _466_v0);
          _577_v62 = (_577_v62).update((_this).f21, _579_v63);
          let _rhs52 = (_this).f21;
          let _rhs53 = new BigNumber(((_this).f15).length);
          let _lhs41 = globalState;
          let _lhs42 = globalState;
          _lhs41.f7 = _rhs52;
          _lhs42.f7 = _rhs53;
        } else {
          let _580_v64;
          _580_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f21);
          let _581_v65;
          _581_v65 = _dafny.Map.Empty.slice().updateUnsafe(_507_v19.f19,_580_v64);
          _581_v65 = _581_v65;
          let _index57 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length));
          (_467_v1)[_index57] = (_this).f21;
          let _index58 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_467_v1).length));
          (_467_v1)[_index58] = (_this).f21;
          let _index59 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_467_v1).length));
          let _rhs54 = (_this).f21;
          let _rhs55 = (((_this).f21).minus((_this).f21)).plus((_dafny.ZERO).minus((_this).f21));
          let _lhs43 = _467_v1;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length));
          let _lhs45 = _467_v1;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_467_v1).length));
          _lhs43[_lhs44] = _rhs54;
          _lhs45[_lhs46] = _rhs55;
          let _582_v66;
          _582_v66 = _dafny.Seq.of((_this).f17);
          let _583_v67;
          _583_v67 = _dafny.Set.fromElements((_582_v66)[_module.__default.safeIndex((_this).f21, new BigNumber((_582_v66).length))], true);
          _583_v67 = (_dafny.Set.fromElements(_507_v19.f19)).Intersect(_583_v67);
          let _584_v68;
          _584_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f15);
          let _585_v69;
          let _nw67 = Array((new BigNumber(24)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update(_466_v0, _module.__default.safeIndex((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], new BigNumber((_466_v0).length)), (_507_v19).f20)).length);
          _nw67[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(2)).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(3)).toNumber()] = new BigNumber((_504_v16).length);
          _nw67[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f21);
          _nw67[(new BigNumber(5)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(6)).toNumber()] = new BigNumber(256);
          _nw67[(new BigNumber(7)).toNumber()] = new BigNumber(565);
          _nw67[(new BigNumber(8)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(9)).toNumber()] = new BigNumber(780);
          _nw67[(new BigNumber(10)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(11)).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(12)).toNumber()] = _module.__default.fm0(_582_v66, _507_v19.f19, (_this).f22, (_this).fm8(_module.__default.fm30(new BigNumber((_582_v66).length), globalState), _this.f16, globalState), globalState);
          _nw67[(new BigNumber(13)).toNumber()] = new BigNumber((_584_v68).length);
          _nw67[(new BigNumber(14)).toNumber()] = (((_504_v16).contains(false)) ? ((_504_v16).get(false)) : ((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))]));
          _nw67[(new BigNumber(15)).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(16)).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(17)).toNumber()] = (_this).f21;
          _nw67[(new BigNumber(18)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(19)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(20)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(21)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(22)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _nw67[(new BigNumber(23)).toNumber()] = (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))];
          _585_v69 = _nw67;
          let _586_v70;
          _586_v70 = _module.D2.create_DC4(_585_v69, ((_this).f21).isEqualTo(new BigNumber((_583_v67).length)), (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], _module.__default.fm30((_this).f21, globalState), (_this).f14);
          let _587_v71;
          _587_v71 = _dafny.Seq.of((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))]);
          let _588_v72;
          _588_v72 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("txsjiq"),(_this).f22);
          let _589_v73;
          _589_v73 = _dafny.Seq.of(_588_v72);
          let _590_v74;
          _590_v74 = _dafny.Seq.of(_588_v72, (_589_v73)[_module.__default.safeIndex((_this).f21, new BigNumber((_589_v73).length))], _588_v72);
          let _591_v75;
          _591_v75 = _dafny.Map.Empty.slice().updateUnsafe(_585_v69,(_503_cf17).IsDisjointFrom((_dafny.MultiSet.FromArray(_587_v71)).update((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], _module.__default.abs(new BigNumber(((_590_v74)[_module.__default.safeIndex((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], new BigNumber((_590_v74).length))]).length)))));
          let _rhs56 = _module.D2.create_DC4(_585_v69, (_this).f22, (_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], (((_this).f22) ? ((_582_v66)[_module.__default.safeIndex((_467_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_467_v1).length))], new BigNumber((_582_v66).length))]) : (_507_v19.f19)), (_this).f14);
          let _rhs57 = ((_dafny.areEqual(_module.__default.fm38(false, _507_v19.f19, _507_v19.f19, (_this).f22, globalState), _582_v66)) ? (_591_v75) : ((_591_v75).Merge(_591_v75)));
          _586_v70 = _rhs56;
          _591_v75 = _rhs57;
          (globalState).f2 = _module.__default.fm38((_this).f22, (_this).f17, false, (_this).f22, globalState);
        }
      } else {
        let _592___mcc_h10 = (_source10).cf27;
        let _593_cf27 = _592___mcc_h10;
        let _594_v76;
        _594_v76 = _dafny.Seq.of(_this.f12);
        _594_v76 = _594_v76;
        let _595_v77;
        let _nw68 = Array((new BigNumber(7)).toNumber()).fill([]);
        _595_v77 = _nw68;
        let _index61 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length));
        (_595_v77)[_index61] = _467_v1;
        let _index62 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length));
        (_595_v77)[_index62] = _467_v1;
        let _index63 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index63] = (_this).f22;
        let _index64 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index64] = (_this).f17;
        if (!(!((_this).f22))) {
          let _596_v78;
          let _nw69 = Array((new BigNumber(11)).toNumber());
          _596_v78 = _nw69;
          let _597_v79;
          _597_v79 = new _dafny.CodePoint('b'.codePointAt(0));
          let _598_v80;
          let _nw70 = new _module.C0();
          _nw70.__ctor(false, _597_v79, _this.f12, (_this).f22);
          _598_v80 = _nw70;
          let _index65 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_596_v78).length));
          (_596_v78)[_index65] = _598_v80;
          let _index66 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_596_v78).length));
          let _rhs58 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))];
          let _rhs59 = _598_v80;
          let _lhs47 = (_this).f14;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _lhs49 = _596_v78;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_596_v78).length));
          _lhs47[_lhs48] = _rhs58;
          _lhs49[_lhs50] = _rhs59;
          (_598_v80).f19 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), ((_599_v80) => function (_600_i5) {
            return (_599_v80).f20;
          })(_598_v80)), (_598_v80).f20);
          let _601_v81;
          _601_v81 = _dafny.Set.fromElements(((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))], (_this).f13);
          (_598_v80).f19 = (((((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))]) ? (new BigNumber((_601_v81).length)) : (new BigNumber((_466_v0).length)))).isLessThan((_this).f21);
          let _602_v82;
          _602_v82 = _dafny.Seq.of(new BigNumber((_466_v0).length));
          let _603_v83;
          _603_v83 = _dafny.Map.Empty.slice().updateUnsafe(_602_v82,_598_v80.f19);
          let _604_v84;
          _604_v84 = _dafny.Seq.of(_602_v82);
          _603_v83 = (_603_v83).update((_604_v84)[_module.__default.safeIndex((_this).f21, new BigNumber((_604_v84).length))], (_this).f22);
          let _605_v85;
          let _nw71 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _605_v85 = _nw71;
          let _index68 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_605_v85).length));
          (_605_v85)[_index68] = (_598_v80).f20;
          let _606_v86;
          _606_v86 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), ((_607_v79) => function (_608_i6) {
            return _607_v79;
          })(_597_v79)));
          let _index69 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_605_v85).length));
          let _index70 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _rhs60 = (_this).fm9(new _dafny.CodePoint('y'.codePointAt(0)), globalState);
          let _rhs61 = _597_v79;
          let _rhs62 = false;
          let _rhs63 = (_606_v86).IsSubsetOf(_dafny.Set.fromElements(_466_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), function (_609_i7) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })));
          let _rhs64 = !((_this).f22);
          let _lhs51 = globalState;
          let _lhs52 = _605_v85;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_605_v85).length));
          let _lhs54 = (_this).f14;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _lhs56 = _598_v80;
          let _lhs57 = _598_v80;
          _lhs51.f7 = _rhs60;
          _lhs52[_lhs53] = _rhs61;
          _lhs54[_lhs55] = _rhs62;
          _lhs56.f19 = _rhs63;
          _lhs57.f19 = _rhs64;
        } else {
          let _610_v87;
          _610_v87 = _dafny.Seq.of((_this).f21);
          let _611_v88;
          _611_v88 = _module.D7.create_DC15(_610_v87);
          let _pat_let_tv15 = globalState;
          _611_v88 = ((((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))]) ? (function (_pat_let12_0) {
            return function (_612_dt__update__tmp_h3) {
              return function (_pat_let13_0) {
                return function (_613_dt__update_hcf28_h0) {
                  return _module.D7.create_DC15(_613_dt__update_hcf28_h0);
                }(_pat_let13_0);
              }(_module.__default.fm34((_this).f21, _pat_let_tv15));
            }(_pat_let12_0);
          }(_611_v88)) : (_611_v88));
          let _614_v89;
          _614_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f13);
          let _615_v90;
          _615_v90 = _module.D6.create_DC13((_this).f17, (_this).f22, new BigNumber((_466_v0).length), (_this).f21, (((_614_v89).contains((_this).f13)) ? ((_614_v89).get((_this).f13)) : ((_this).f17)));
          let _index71 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index71] = (_615_v90).dtor_cf26;
          let _616_v91;
          let _nw72 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          _616_v91 = _nw72;
          let _617_v92;
          _617_v92 = _dafny.Seq.of(true, (_this).f17, true, ((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))], ((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))]);
          let _index72 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_616_v91).length));
          (_616_v91)[_index72] = _dafny.Seq.Concat(_617_v92, _617_v92);
          let _index73 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_616_v91).length));
          let _index74 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _index75 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _rhs65 = _617_v92;
          let _rhs66 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))];
          let _rhs67 = ((_module.__default.fm30((_this).f21, globalState)) ? (_466_v0) : (_466_v0));
          let _rhs68 = !(!((_this).f13));
          let _lhs58 = _616_v91;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_616_v91).length));
          let _lhs60 = (_this).f14;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _lhs62 = (_this).f14;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          _lhs58[_lhs59] = _rhs65;
          _lhs60[_lhs61] = _rhs66;
          _466_v0 = _rhs67;
          _lhs62[_lhs63] = _rhs68;
          let _618_v93;
          _618_v93 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))],(_this).f21);
          let _arr6 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _index76 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          _arr6[_index76] = new BigNumber((_618_v93).length);
          let _index77 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_467_v1).length));
          (_467_v1)[_index77] = (_610_v87)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_610_v87).length))];
          let _arr7 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _index78 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          _arr7[_index78] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f17)).length);
          let _619_v94;
          _619_v94 = new _dafny.CodePoint('x'.codePointAt(0));
          let _620_v95;
          _620_v95 = _module.D6.create_DC14(_module.__default.fm39(false, ((_this).f14)[_module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length))], globalState));
          let _621_v96;
          let _nw73 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _621_v96 = _nw73;
          let _622_v97;
          _622_v97 = _dafny.Seq.of(_621_v96, _621_v96, _621_v96);
          let _arr8 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _index79 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          let _index80 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _index81 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_467_v1).length));
          let _arr9 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _index82 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          let _rhs69 = (_this).f21;
          let _rhs70 = _dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.update(_466_v0, _module.__default.safeIndex((_this).f21, new BigNumber((_466_v0).length)), _619_v94), _466_v0, _module.__default.fm32(_620_v95, globalState)), _466_v0);
          let _rhs71 = (new BigNumber((_610_v87).length)).plus(((_this).f21).plus((_this).f21));
          let _rhs72 = new BigNumber((_622_v97).length);
          let _lhs64 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _lhs65 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          let _lhs66 = (_this).f14;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(801), new BigNumber(((_this).f14).length));
          let _lhs68 = _467_v1;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_467_v1).length));
          let _lhs70 = (_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))];
          let _lhs71 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length));
          _lhs64[_lhs65] = _rhs69;
          _lhs66[_lhs67] = _rhs70;
          _lhs68[_lhs69] = _rhs71;
          _lhs70[_lhs71] = _rhs72;
          let _index83 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_467_v1).length));
          (_467_v1)[_index83] = _module.__default.safeDivisionInt(new BigNumber(-317), ((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))])[_module.__default.safeIndex(new BigNumber(181), new BigNumber(((_595_v77)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_595_v77).length))]).length))]);
        }
      }
      let _623_v98;
      _623_v98 = false;
      let _624_v99;
      _624_v99 = _dafny.Seq.of(false, (_this).f13, _623_v98, _623_v98, (_this).f13);
      let _625_v100;
      _625_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_624_v99);
      let _626_v101;
      _626_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_624_v99).length),(_this).f22);
      _623_v98 = _dafny.areEqual(_dafny.Seq.update((((_625_v100).contains(new BigNumber((_626_v101).length))) ? ((_625_v100).get(new BigNumber((_626_v101).length))) : (_624_v99)), _module.__default.safeIndex((_this).f21, new BigNumber(((((_625_v100).contains(new BigNumber((_626_v101).length))) ? ((_625_v100).get(new BigNumber((_626_v101).length))) : (_624_v99))).length)), (_this).f22), _dafny.Seq.Concat(_624_v99, _624_v99));
      let _627_v102;
      _627_v102 = _dafny.MultiSet.fromElements((_this).f17);
      let _628_v103;
      _628_v103 = _627_v102;
      let _629_v104;
      let _nw74 = Array((new BigNumber(16)).toNumber());
      _nw74[(_dafny.ZERO).toNumber()] = _627_v102;
      _nw74[(_dafny.ONE).toNumber()] = _627_v102;
      _nw74[(new BigNumber(2)).toNumber()] = (_627_v102).Difference((_dafny.MultiSet.fromElements((_this).f22, true)).update(true, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), function (_630_i9) {
        return (_this).f21;
      })).length))));
      _nw74[(new BigNumber(3)).toNumber()] = _627_v102;
      _nw74[(new BigNumber(4)).toNumber()] = (_627_v102).Intersect(_627_v102);
      _nw74[(new BigNumber(5)).toNumber()] = (_628_v103);
      _nw74[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f13));
      _nw74[(new BigNumber(7)).toNumber()] = _627_v102;
      _nw74[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_this).f22);
      _nw74[(new BigNumber(9)).toNumber()] = (_module.__default.fm40(globalState)).Intersect(_627_v102);
      _nw74[(new BigNumber(10)).toNumber()] = _627_v102;
      _nw74[(new BigNumber(11)).toNumber()] = (_627_v102).Difference(_627_v102);
      _nw74[(new BigNumber(12)).toNumber()] = _627_v102;
      _nw74[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_624_v99);
      _nw74[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f22)).Union(_627_v102);
      _nw74[(new BigNumber(15)).toNumber()] = (_627_v102).Difference(_dafny.MultiSet.fromElements((_this).f13));
      _629_v104 = _nw74;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_629_v104).length))) {
        let _631_i8 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_631_i8)) && ((_631_i8).isLessThan(new BigNumber((_629_v104).length))))) {
          (_629_v104)[(_631_i8)] = (_627_v102).Union(_627_v102);
        }
      }
      let _632_v105;
      let _init10 = function (_633_i10) {
        return (_this.f16).dtor_cf1;
      };
      let _nw75 = Array((new BigNumber(29)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw75.length); _i0_10++) {
        _nw75[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _632_v105 = _nw75;
      let _index84 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_632_v105).length));
      (_632_v105)[_index84] = ((_623_v98) ? (_466_v0) : (_466_v0));
      let _634_v106;
      _634_v106 = _module.D6.create_DC14(_module.D6.create_DC13((_this).f17, (_this).f13, (_this).f21, (_this).f21, _623_v98));
      let _635_v107;
      _635_v107 = _module.D6.create_DC14(_634_v106);
      let _636_v108;
      _636_v108 = _module.D6.create_DC14(_634_v106);
      let _pat_let_tv16 = _466_v0;
      let _pat_let_tv17 = _466_v0;
      let _index85 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_632_v105).length));
      (_632_v105)[_index85] = function (_source12) {
        if (_source12.is_DC12) {
          let _637___mcc_h22 = (_source12).cf18;
          let _638___mcc_h23 = (_source12).cf19;
          let _639___mcc_h24 = (_source12).cf20;
          let _640___mcc_h25 = (_source12).cf21;
          let _641_cf21 = _640___mcc_h25;
          let _642_cf20 = _639___mcc_h24;
          let _643_cf19 = _638___mcc_h23;
          let _644_cf18 = _637___mcc_h22;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(86)), function (_645_i11) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          });
        } else if (_source12.is_DC13) {
          let _646___mcc_h26 = (_source12).cf22;
          let _647___mcc_h27 = (_source12).cf23;
          let _648___mcc_h28 = (_source12).cf24;
          let _649___mcc_h29 = (_source12).cf25;
          let _650___mcc_h30 = (_source12).cf26;
          let _651_cf26 = _650___mcc_h30;
          let _652_cf25 = _649___mcc_h29;
          let _653_cf24 = _648___mcc_h28;
          let _654_cf23 = _647___mcc_h27;
          let _655_cf22 = _646___mcc_h26;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-521)), function (_656_i12) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          });
        } else if (_source12.is_DC11) {
          let _657___mcc_h31 = (_source12).cf17;
          let _658_cf17 = _657___mcc_h31;
          return _pat_let_tv16;
        } else {
          let _659___mcc_h32 = (_source12).cf27;
          let _660_cf27 = _659___mcc_h32;
          return _pat_let_tv17;
        }
      }(_636_v108);
      let _661_v109;
      let _nw76 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
      _661_v109 = _nw76;
      let _662_v110;
      _662_v110 = new _dafny.CodePoint('u'.codePointAt(0));
      let _663_v111;
      _663_v111 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f15);
      let _index86 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_661_v109).length));
      (_661_v109)[_index86] = (_module.__default.fm41((_this).f21, _662_v110, globalState)).Merge(_663_v111);
      let _664_v112;
      _664_v112 = _dafny.Seq.of(_466_v0, _466_v0);
      let _665_v113;
      _665_v113 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), ((_666_v110) => function (_667_i13) {
        return _666_v110;
      })(_662_v110)), (_664_v112)[_module.__default.safeIndex((_this).f21, new BigNumber((_664_v112).length))]);
      let _668_v114;
      _668_v114 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,true);
      let _669_v115;
      _669_v115 = _module.D3.create_DC5(_module.__default.fm33((_this).f22, new BigNumber((_668_v114).length), globalState));
      let _pat_let_tv18 = _663_v111;
      let _pat_let_tv19 = _663_v111;
      let _pat_let_tv20 = _663_v111;
      let _pat_let_tv21 = _663_v111;
      let _pat_let_tv22 = _663_v111;
      let _pat_let_tv23 = _663_v111;
      let _index87 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_661_v109).length));
      let _rhs73 = function (_source13) {
        let _670___mcc_h33 = _source13;
        let _671_cf16 = _670___mcc_h33;
        return (_this).f22;
      }(_665_v113);
      let _rhs74 = function (_source14) {
        if (_source14.is_DC6) {
          let _672___mcc_h34 = (_source14).cf12;
          let _673___mcc_h35 = (_source14).cf13;
          let _674_cf13 = _673___mcc_h35;
          let _675_cf12 = _672___mcc_h34;
          return (_pat_let_tv18).Merge(_pat_let_tv19);
        } else if (_source14.is_DC7) {
          return _pat_let_tv20;
        } else if (_source14.is_DC8) {
          let _676___mcc_h36 = (_source14).cf14;
          let _677_cf14 = _676___mcc_h36;
          return _pat_let_tv21;
        } else {
          let _678___mcc_h37 = (_source14).cf11;
          let _679_cf11 = _678___mcc_h37;
          if ((_this).f17) {
            return _pat_let_tv22;
          } else {
            return _pat_let_tv23;
          }
        }
      }(_669_v115);
      let _rhs75 = (_this).f21;
      let _rhs76 = (_this).f21;
      let _lhs72 = _661_v109;
      let _lhs73 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_661_v109).length));
      let _lhs74 = globalState;
      let _lhs75 = globalState;
      _623_v98 = _rhs73;
      _lhs72[_lhs73] = _rhs74;
      _lhs74.f7 = _rhs75;
      _lhs75.f7 = _rhs76;
      let _680_i14;
      _680_i14 = _dafny.ZERO;
      L3: {
        while ((_this).f17) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_680_i14)) {
              break L3;
            }
            _680_i14 = (_680_i14).plus(_dafny.ONE);
            let _681_v116;
            _681_v116 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f13) ? ((_this).f22) : (!(!((((_626_v101).contains((_this).f21)) ? ((_626_v101).get((_this).f21)) : (_623_v98)))))),_dafny.Seq.UnicodeFromString("gr"));
            _681_v116 = (_681_v116).update((_624_v99)[_module.__default.safeIndex((_this).f21, new BigNumber((_624_v99).length))], _dafny.Seq.Concat((_632_v105)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_632_v105).length))], (_632_v105)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_632_v105).length))]));
            let _index88 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length));
            (_467_v1)[_index88] = (_this).f21;
            let _index89 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length));
            (_467_v1)[_index89] = (_this).f21;
            _623_v98 = !(((_623_v98) ? ((_467_v1)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length))]) : (new BigNumber((_dafny.Seq.UnicodeFromString("av")).length)))).isEqualTo(new BigNumber(540));
            if ((_this).f17) {
              let _pat_let_tv24 = _662_v110;
              let _pat_let_tv25 = _662_v110;
              _669_v115 = function (_pat_let14_0) {
                return function (_684_dt__update__tmp_h5) {
                  return function (_pat_let17_0) {
                    return function (_685_dt__update_hcf11_h1) {
                      return _module.D3.create_DC5(_685_dt__update_hcf11_h1);
                    }(_pat_let17_0);
                  }(_pat_let_tv25);
                }(_pat_let14_0);
              }(function (_pat_let15_0) {
                return function (_682_dt__update__tmp_h4) {
                  return function (_pat_let16_0) {
                    return function (_683_dt__update_hcf11_h0) {
                      return _module.D3.create_DC5(_683_dt__update_hcf11_h0);
                    }(_pat_let16_0);
                  }(_pat_let_tv24);
                }(_pat_let15_0);
              }(_669_v115));
              let _index90 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index90] = false;
              let _index91 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index91] = !((_this).f13);
              let _index92 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index92] = (_this).f17;
              let _686_v117;
              _686_v117 = _dafny.Seq.of((_this).f21);
              let _687_v118;
              _687_v118 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_686_v117).length),_626_v101);
              _687_v118 = (_687_v118).update((_this).f21, (_626_v101).Merge(_626_v101));
              _623_v98 = !((_this).f13);
            } else {
              let _index93 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length));
              (_467_v1)[_index93] = (_this).f21;
              _626_v101 = (_626_v101).update((_467_v1)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length))], (_this).f22);
              let _688_v119;
              let _nw77 = new _module.C0();
              _nw77.__ctor(!((_this).f13), _662_v110, _this.f12, ((!((_this).f22)) ? (!(true)) : ((_this).f22)));
              _688_v119 = _nw77;
              _623_v98 = (_this).f17;
              let _689_v120;
              _689_v120 = _dafny.MultiSet.fromElements(new BigNumber(449), (_467_v1)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_467_v1).length))], new BigNumber(((_this).f15).length), (_this).f21);
              let _690_v121;
              let _nw78 = Array((new BigNumber(21)).toNumber());
              _nw78[(_dafny.ZERO).toNumber()] = _662_v110;
              _nw78[(_dafny.ONE).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(2)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(3)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(4)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(5)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
              _nw78[(new BigNumber(7)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(8)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(9)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(10)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(11)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(12)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
              _nw78[(new BigNumber(14)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(15)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(16)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(17)).toNumber()] = (_688_v119).f20;
              _nw78[(new BigNumber(18)).toNumber()] = _662_v110;
              _nw78[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
              _nw78[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
              _690_v121 = _nw78;
              let _rhs77 = ((_623_v98) ? (_dafny.Seq.UnicodeFromString("pwfti")) : (_dafny.Seq.Concat(_466_v0, _466_v0)));
              let _rhs78 = !((_689_v120).IsProperSubsetOf((_689_v120).update(new BigNumber(365), _module.__default.abs(new BigNumber((_dafny.MultiSet.fromElements(_690_v121)).cardinality())))));
              let _lhs76 = _688_v119;
              _466_v0 = _rhs77;
              _lhs76.f19 = _rhs78;
            }
          }
        }
      }
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _691_v0;
      let _nw79 = Array((new BigNumber(3)).toNumber());
      _nw79[(_dafny.ZERO).toNumber()] = (_this).f21;
      _nw79[(_dafny.ONE).toNumber()] = (_this).f21;
      _nw79[(new BigNumber(2)).toNumber()] = (_this).f21;
      _691_v0 = _nw79;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_691_v0).length))) {
        let _692_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_692_i0)) && ((_692_i0).isLessThan(new BigNumber((_691_v0).length))))) {
          (_691_v0)[(_692_i0)] = (_692_i0).multipliedBy((p0).plus(p0));
        }
      }
      (globalState).f7 = ((_this).f21).plus((_dafny.ZERO).minus(p0));
      let _693_v1;
      _693_v1 = new _dafny.CodePoint('m'.codePointAt(0));
      let _694_v2;
      let _nw80 = new _module.C0();
      _nw80.__ctor((_this).f22, _693_v1, (((_this).f22) ? (_this.f12) : (_module.D1.create_DC2((_this).f21, (_this).fm8((_this).f22, _this.f16, globalState), (_this).f13))), (_this).f13);
      _694_v2 = _nw80;
      let _695_v3;
      let _init11 = ((_696_v2) => function (_697_i1) {
        return _dafny.Seq.of(_696_v2.f19, (_this).f17);
      })(_694_v2);
      let _nw81 = Array((new BigNumber(18)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw81.length); _i0_11++) {
        _nw81[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _695_v3 = _nw81;
      let _698_v4;
      _698_v4 = _module.D8.create_DC18(_695_v3);
      let _699_v5;
      _699_v5 = _dafny.Map.Empty.slice().updateUnsafe((_698_v4).dtor_cf31,_dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), ((_700_p0) => function (_701_i2) {
        return _700_p0;
      })(p0)));
      let _702_v6;
      _702_v6 = _dafny.Seq.of(p0);
      let _703_v7;
      _703_v7 = _dafny.Seq.of(p0, (_this).f21, (_this).f21, new BigNumber((_702_v6).length), p0);
      _699_v5 = (_699_v5).update(_695_v3, _702_v6);
      if ((_this).f17) {
        let _704_v8;
        _704_v8 = _dafny.Seq.of(_694_v2.f19, _694_v2.f19, (_this).f13);
        let _index94 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_695_v3).length));
        (_695_v3)[_index94] = _704_v8;
        let _705_v9;
        let _init12 = ((_706_v2, _707_v1) => function (_708_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe((_706_v2).f20,_dafny.Seq.of(_707_v1));
        })(_694_v2, _693_v1);
        let _nw82 = Array((new BigNumber(9)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw82.length); _i0_12++) {
          _nw82[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _705_v9 = _nw82;
        let _709_v10;
        _709_v10 = _dafny.Seq.of((_694_v2).f20, _module.__default.fm33(_694_v2.f19, p0, globalState));
        let _710_v11;
        _710_v11 = _dafny.Map.Empty.slice().updateUnsafe(_693_v1,_709_v10);
        let _index95 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_705_v9).length));
        (_705_v9)[_index95] = _710_v11;
        let _index96 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_695_v3).length));
        let _index97 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_705_v9).length));
        let _rhs79 = _module.__default.fm0(_dafny.Seq.update(_dafny.Seq.update(_704_v8, _module.__default.safeIndex((_this).f21, new BigNumber((_704_v8).length)), _694_v2.f19), _module.__default.safeIndex((_this).f21, new BigNumber((_dafny.Seq.update(_704_v8, _module.__default.safeIndex((_this).f21, new BigNumber((_704_v8).length)), _694_v2.f19)).length)), (_this).f17), !(((!((_this).f13)) ? ((_this).f17) : (_694_v2.f19))), ((_694_v2.f19) ? (_module.__default.fm30((_702_v6)[_module.__default.safeIndex(p0, new BigNumber((_702_v6).length))], globalState)) : (_694_v2.f19)), _module.__default.safeModuloInt(p0, p0), globalState);
        let _rhs80 = _dafny.Seq.Concat(_704_v8, _704_v8);
        let _rhs81 = _710_v11;
        let _lhs77 = globalState;
        let _lhs78 = _695_v3;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_695_v3).length));
        let _lhs80 = _705_v9;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_705_v9).length));
        _lhs77.f7 = _rhs79;
        _lhs78[_lhs79] = _rhs80;
        _lhs80[_lhs81] = _rhs81;
        (globalState).f7 = (_this).fm9((_694_v2).f20, globalState);
        let _711_v12;
        _711_v12 = _dafny.MultiSet.fromElements((_this).f21, (_this).f21);
        let _712_v13;
        _712_v13 = _dafny.MultiSet.fromElements((_dafny.MultiSet.fromElements(new BigNumber(((_695_v3)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_695_v3).length))]).length), new BigNumber((_dafny.MultiSet.FromArray((_695_v3)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_695_v3).length))])).cardinality()), (_this).f21)).Intersect((_711_v12).update(p0, _module.__default.abs(new BigNumber((_709_v10).length)))));
        (globalState).f7 = (((_712_v13).contains(_711_v12)) ? ((_712_v13).get(_711_v12)) : ((_this).f21));
        let _713_v14;
        _713_v14 = _dafny.Set.fromElements(_dafny.Seq.Concat(_709_v10, _709_v10), _dafny.Seq.UnicodeFromString("wk"), _709_v10, _dafny.Seq.Concat(_709_v10, _709_v10));
        _713_v14 = _713_v14;
        (_694_v2).f19 = (_this).f22;
      } else {
        let _714_v15;
        _714_v15 = _dafny.MultiSet.fromElements((_this).f22, _694_v2.f19);
        let _715_v16;
        _715_v16 = _dafny.Map.Empty.slice().updateUnsafe(_714_v15,_694_v2.f19);
        let _716_v17;
        _716_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_dafny.Map.Empty.slice().updateUnsafe(_714_v15,_694_v2.f19));
        _715_v16 = (((_716_v17).contains((_this).f21)) ? ((_716_v17).get((_this).f21)) : (_715_v16));
        if ((_this).f13) {
          let _717_v18;
          _717_v18 = _dafny.Seq.UnicodeFromString("b");
          let _718_v19;
          _718_v19 = _module.D6.create_DC13((_this).f17, false, (_this).f21, (_this).f21, _694_v2.f19);
          let _719_v20;
          _719_v20 = _dafny.Set.fromElements(_694_v2.f19);
          let _720_v21;
          let _nw83 = new _module.C1();
          _nw83.__ctor((_694_v2).f20, _this.f12, (_this).f17);
          _720_v21 = _nw83;
          let _721_v22;
          _721_v22 = _dafny.MultiSet.fromElements(_720_v21);
          let _722_v23;
          let _nw84 = Array((new BigNumber(22)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _694_v2.f19;
          _nw84[(_dafny.ONE).toNumber()] = (_this).f13;
          _nw84[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_717_v18, _module.__default.fm32(_module.D6.create_DC14(_718_v19), globalState));
          _nw84[(new BigNumber(3)).toNumber()] = _694_v2.f19;
          _nw84[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_694_v2.f19, (_this).f22), _module.__default.fm38((_this).f13, (_this).f22, _694_v2.f19, (_this).f17, globalState));
          _nw84[(new BigNumber(5)).toNumber()] = (_this).f22;
          _nw84[(new BigNumber(6)).toNumber()] = _module.__default.fm30(new BigNumber(-545), globalState);
          _nw84[(new BigNumber(7)).toNumber()] = _module.__default.fm30(p0, globalState);
          _nw84[(new BigNumber(8)).toNumber()] = _694_v2.f19;
          _nw84[(new BigNumber(9)).toNumber()] = (_this).f13;
          _nw84[(new BigNumber(10)).toNumber()] = (((_this).f13) ? (!(_694_v2.f19)) : (true));
          _nw84[(new BigNumber(11)).toNumber()] = ((_702_v6)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_702_v6).length))]).isLessThanOrEqualTo(p0);
          _nw84[(new BigNumber(12)).toNumber()] = ((_this).f21).isLessThanOrEqualTo(new BigNumber(274));
          _nw84[(new BigNumber(13)).toNumber()] = (_this).f17;
          _nw84[(new BigNumber(14)).toNumber()] = (_719_v20).IsSubsetOf(_719_v20);
          _nw84[(new BigNumber(15)).toNumber()] = (_this).f13;
          _nw84[(new BigNumber(16)).toNumber()] = (_this).f17;
          _nw84[(new BigNumber(17)).toNumber()] = (_721_v22).IsSubsetOf(_721_v22);
          _nw84[(new BigNumber(18)).toNumber()] = (p0).isLessThan((_this).f21);
          _nw84[(new BigNumber(19)).toNumber()] = (new BigNumber((_703_v7).length)).isLessThanOrEqualTo((_this).f21);
          _nw84[(new BigNumber(20)).toNumber()] = _694_v2.f19;
          _nw84[(new BigNumber(21)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_717_v18, _717_v18);
          _722_v23 = _nw84;
          _722_v23 = _722_v23;
          let _723_v24;
          let _nw85 = Array((new BigNumber(23)).toNumber()).fill([]);
          _723_v24 = _nw85;
          let _724_v25;
          let _nw86 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _724_v25 = _nw86;
          let _index98 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_723_v24).length));
          (_723_v24)[_index98] = _724_v25;
          let _index99 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_723_v24).length));
          (_723_v24)[_index99] = _724_v25;
          (_694_v2).f19 = ((_694_v2.f19) ? (!((_720_v21).f13)) : (_module.__default.fm30(p0, globalState)));
          let _index100 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_722_v23).length));
          (_722_v23)[_index100] = !((_this).f13);
          let _index101 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_722_v23).length));
          (_722_v23)[_index101] = (_this).f17;
          let _725_v26;
          _725_v26 = _dafny.Map.Empty.slice().updateUnsafe(false,(_720_v21).f13);
          _702_v6 = _dafny.Seq.update(_dafny.Seq.Concat(_702_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-120)), ((_726_p0) => function (_727_i4) {
            return _726_p0;
          })(p0))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_702_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-120)), ((_728_p0) => function (_729_i4) {
            return _728_p0;
          })(p0)))).length)), new BigNumber((_725_v26).length));
        } else {
          let _730_v27;
          _730_v27 = _dafny.Seq.of(true, (_this).f22, (_this).f22);
          (globalState).f7 = _module.__default.fm0(_730_v27, false, (_this).f22, (p0).minus((_this).f21), globalState);
          let _731_v28;
          let _nw87 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
          _731_v28 = _nw87;
          let _732_v29;
          _732_v29 = _dafny.Set.fromElements(_694_v2.f19);
          let _733_v30;
          _733_v30 = _module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(!(_694_v2.f19),(_this).f21), _732_v29, _dafny.MultiSet.fromElements(new BigNumber(680)), _this.f16);
          let _734_v31;
          _734_v31 = _dafny.Seq.of(_732_v29, (_733_v30).dtor_cf33, _732_v29, _732_v29, _732_v29);
          let _index102 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_731_v28).length));
          (_731_v28)[_index102] = _734_v31;
          let _index103 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_731_v28).length));
          (_731_v28)[_index103] = _734_v31;
          let _735_v32;
          let _out8;
          _out8 = _module.__default.m0((_694_v2).f20, (_this).f22, globalState);
          _735_v32 = _out8;
          (globalState).f7 = (_dafny.ZERO).minus(((p0).multipliedBy(p0)).multipliedBy(p0));
          (globalState).f7 = (_dafny.ZERO).minus(((!((_this).f22) || (_694_v2.f19)) ? (p0) : (p0)));
        }
        let _736_v33;
        _736_v33 = _module.D6.create_DC11(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_702_v6, _702_v6)));
        _736_v33 = _736_v33;
        r0 = _694_v2.f19;
        (globalState).f7 = ((_this).f21).minus((_dafny.ZERO).minus(p0));
      }
      (globalState).f7 = (new BigNumber(-221)).minus(p0);
      r0 = (_this).f22;
      return r0;
    }
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _737_v0;
      let _nw88 = Array((new BigNumber(4)).toNumber()).fill(false);
      _737_v0 = _nw88;
      _737_v0 = _737_v0;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_737_v0).length))) {
        let _738_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_738_i0)) && ((_738_i0).isLessThan(new BigNumber((_737_v0).length))))) {
          (_737_v0)[(_738_i0)] = (_this).f17;
        }
      }
      let _739_v1;
      _739_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
      let _740_v2;
      _740_v2 = _dafny.Seq.of(new BigNumber((_739_v1).length), (_this).f21, (_this).f21);
      _740_v2 = _740_v2;
      let _741_v3;
      _741_v3 = _dafny.Map.Empty.slice().updateUnsafe(!((new BigNumber(908)).isEqualTo(p3)),(_this).f22);
      _741_v3 = (_741_v3).update(((_this).f17) === (p0), (_this).f17);
      if ((_this).f22) {
        (globalState).f7 = (_this).fm8((_this).f17, _this.f16, globalState);
        let _742_v4;
        let _nw89 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _742_v4 = _nw89;
        _742_v4 = _742_v4;
        let _743_v5;
        _743_v5 = _module.D3.create_DC5(p1);
        let _744_v6;
        let _nw90 = new _module.C1();
        _nw90.__ctor((_743_v5).dtor_cf11, _module.D1.create_DC2(new BigNumber(-728), p3, (_this).f17), true);
        _744_v6 = _nw90;
        let _745_v7;
        let _out9;
        _out9 = _module.__default.m0(_744_v6.f23, (_this).f22, globalState);
        _745_v7 = _out9;
        r0 = _module.__default.fm30(p3, globalState);
      } else {
        let _746_v8;
        _746_v8 = _dafny.Seq.of(p0);
        (globalState).f7 = ((_dafny.ZERO).minus((_this).f21)).minus(_module.__default.fm0(_746_v8, (_this).f17, (_this).f13, (_this).f21, globalState));
        let _747_v9;
        _747_v9 = _dafny.Seq.UnicodeFromString("rkdajmk");
        (globalState).f7 = new BigNumber((_dafny.Seq.Concat(_747_v9, _dafny.Seq.UnicodeFromString("lvxjwypfe"))).length);
        let _748_v10;
        _748_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_737_v0);
        _748_v10 = (_748_v10).update((_this).f21, (_this).f14);
        (globalState).f7 = p3;
        let _749_v11;
        let _nw91 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _749_v11 = _nw91;
        let _750_v12;
        _750_v12 = _dafny.Map.Empty.slice().updateUnsafe(_737_v0,_749_v11);
        _750_v12 = (_750_v12).Merge(_dafny.Map.Empty.slice().updateUnsafe(_737_v0,_749_v11));
      }
      r0 = !((_this).f13);
      r0 = (_this).f22;
      return r0;
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = _module.D1.Default();
      this._f16 = _module.D1.Default();
      this._f14 = [];
      this._f15 = _dafny.Set.Empty;
      this._f13 = false;
      this._f17 = false;
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T3, _module.T0, _module.T1, _module.T2];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f18, f16, f17, f14, f15, f12, f13) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),new BigNumber((_dafny.Seq.UnicodeFromString("tfdhco")).length));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((new BigNumber(144)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)), new BigNumber(150));
    };
    fm9(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(874)).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f17)).length))).cardinality()))))).multipliedBy(((false) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17)).length)) : (new BigNumber(537))));
    };
    fm15(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_module.__default.safeModuloInt(new BigNumber(981), new BigNumber(611))).minus(_module.__default.safeDivisionInt(new BigNumber(-191), new BigNumber(434))));
    };
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if ((_this.f12).dtor_cf4) {
        let _751_v0;
        let _init13 = ((_752_p0) => function (_753_i0) {
          return (_753_i0).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f13, _752_p0, (_this).f18, (_this).f17)).cardinality()));
        })(p0);
        let _nw92 = Array((new BigNumber(15)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw92.length); _i0_13++) {
          _nw92[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _751_v0 = _nw92;
        let _754_v1;
        _754_v1 = new BigNumber(673);
        let _index104 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length));
        (_751_v0)[_index104] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_754_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_755_p0) => function (_756_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_dafny.Seq.of(_755_p0)).length));
        })(p0))).length)));
        let _757_v2;
        _757_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(true) || ((_this).f13),!((_this).f17));
        let _758_v3;
        _758_v3 = _dafny.Seq.of(new BigNumber(111));
        let _index105 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length));
        let _rhs82 = (_module.__default.safeModuloInt(new BigNumber((_758_v3).length), _754_v1)).plus((_758_v3)[_module.__default.safeIndex((_dafny.ZERO).minus(_754_v1), new BigNumber((_758_v3).length))]);
        let _rhs83 = _757_v2;
        let _rhs84 = _754_v1;
        let _lhs82 = _751_v0;
        let _lhs83 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length));
        _lhs82[_lhs83] = _rhs82;
        _757_v2 = _rhs83;
        _754_v1 = _rhs84;
        let _759_v4;
        _759_v4 = true;
        let _760_v5;
        _760_v5 = _dafny.Seq.UnicodeFromString("db");
        _759_v4 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("rgxktld"), _760_v5);
        let _761_v6;
        _761_v6 = _dafny.Map.Empty.slice().updateUnsafe((_751_v0)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length))],((_751_v0)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length))]).minus(_754_v1));
        let _762_v7;
        _762_v7 = _dafny.Seq.of((_this).f17, (_this).f13);
        _761_v6 = (_761_v6).update((_751_v0)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length))], _module.__default.safeModuloInt(_module.__default.fm0(_762_v7, !(true), (_this).f17, _754_v1, globalState), (_751_v0)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length))]));
        let _763_v8;
        _763_v8 = new _dafny.CodePoint('w'.codePointAt(0));
        let _764_v9;
        let _nw93 = new _module.C0();
        _nw93.__ctor((_this).f13, _763_v8, _this.f12, (_this).f18);
        _764_v9 = _nw93;
        let _765_v10;
        _765_v10 = _dafny.Seq.of(_764_v9, _764_v9);
        let _766_v11;
        _766_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,new BigNumber((_765_v10).length));
        let _index106 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_751_v0).length));
        (_751_v0)[_index106] = (_758_v3)[_module.__default.safeIndex(new BigNumber((_766_v11).length), new BigNumber((_758_v3).length))];
        r0 = _754_v1;
      } else {
        let _767_v12;
        _767_v12 = new BigNumber(352);
        let _768_v13;
        _768_v13 = _dafny.MultiSet.fromElements(p0);
        let _769_v14;
        _769_v14 = _dafny.Seq.of(_768_v13);
        let _770_v15;
        _770_v15 = _dafny.Map.Empty.slice().updateUnsafe(_767_v12,new BigNumber((_769_v14).length));
        (globalState).f7 = _module.__default.safeDivisionInt(new BigNumber((_770_v15).length), new BigNumber(6));
        r0 = _767_v12;
        if ((_this).f18) {
          let _rhs85 = new BigNumber(((_this).f15).length);
          let _rhs86 = (_dafny.ZERO).minus(new BigNumber(-213));
          let _lhs84 = globalState;
          _lhs84.f7 = _rhs85;
          _767_v12 = _rhs86;
          let _771_v16;
          _771_v16 = _dafny.Seq.of((_this).f14);
          _767_v12 = new BigNumber((_771_v16).length);
          let _772_v17;
          _772_v17 = true;
          _772_v17 = _module.__default.fm17(globalState);
          let _773_v18;
          _773_v18 = _dafny.Seq.of(_767_v12);
          let _774_v19;
          _774_v19 = _dafny.MultiSet.fromElements(new BigNumber((_773_v18).length));
          let _775_v20;
          _775_v20 = _dafny.Seq.UnicodeFromString("j");
          let _776_v21;
          let _nw94 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _776_v21 = _nw94;
          let _777_v22;
          _777_v22 = _module.D2.create_DC4(_776_v21, (_this).f18, new BigNumber(945), _module.__default.fm17(globalState), (_this).f14);
          let _778_v23;
          let _nw95 = Array((new BigNumber(17)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = new BigNumber((_773_v18).length);
          _nw95[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_767_v12);
          _nw95[(new BigNumber(2)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(3)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_767_v12, _767_v12);
          _nw95[(new BigNumber(5)).toNumber()] = new BigNumber((_770_v15).length);
          _nw95[(new BigNumber(6)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(7)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(8)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_767_v12, _767_v12);
          _nw95[(new BigNumber(10)).toNumber()] = new BigNumber((_774_v19).cardinality());
          _nw95[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(_767_v12, _767_v12);
          _nw95[(new BigNumber(12)).toNumber()] = new BigNumber((_775_v20).length);
          _nw95[(new BigNumber(13)).toNumber()] = (_777_v22).dtor_cf8;
          _nw95[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_773_v18, _module.__default.safeIndex(new BigNumber((_775_v20).length), new BigNumber((_773_v18).length)), new BigNumber(263))).length);
          _nw95[(new BigNumber(15)).toNumber()] = _767_v12;
          _nw95[(new BigNumber(16)).toNumber()] = _767_v12;
          _778_v23 = _nw95;
          let _index107 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_776_v21).length));
          (_776_v21)[_index107] = new BigNumber(680);
          let _index108 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_776_v21).length));
          let _rhs87 = _767_v12;
          let _rhs88 = _767_v12;
          let _lhs85 = _776_v21;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_776_v21).length));
          _lhs85[_lhs86] = _rhs87;
          _767_v12 = _rhs88;
          let _779_v24;
          _779_v24 = new _dafny.CodePoint('l'.codePointAt(0));
          let _780_v25;
          _780_v25 = _module.D3.create_DC5(new _dafny.CodePoint('s'.codePointAt(0)));
          _779_v24 = (_780_v25).dtor_cf11;
        } else {
          let _781_v26;
          _781_v26 = new _dafny.CodePoint('w'.codePointAt(0));
          let _782_v27;
          _782_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_781_v26);
          let _783_v28;
          _783_v28 = _dafny.Seq.UnicodeFromString("fshc");
          let _784_v29;
          _784_v29 = _dafny.Seq.of((_this).f13);
          let _785_v30;
          let _nw96 = Array((new BigNumber(25)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _781_v26;
          _nw96[(_dafny.ONE).toNumber()] = _781_v26;
          _nw96[(new BigNumber(2)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(3)).toNumber()] = ((!((_this).f13)) ? (_781_v26) : ((((_782_v27).contains(p0)) ? ((_782_v27).get(p0)) : (_781_v26))));
          _nw96[(new BigNumber(4)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(5)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(6)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw96[(new BigNumber(8)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
          _nw96[(new BigNumber(10)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw96[(new BigNumber(12)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(13)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(14)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw96[(new BigNumber(16)).toNumber()] = _module.__default.fm18(new BigNumber((_dafny.Seq.UnicodeFromString("ukupbqlom")).length), _783_v28, (_this).f17, (_this).fm10((_this).f13, (_this).f13, _784_v29, globalState), globalState);
          _nw96[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _nw96[(new BigNumber(18)).toNumber()] = ((!((_this).f18)) ? (_781_v26) : (_781_v26));
          _nw96[(new BigNumber(19)).toNumber()] = (_783_v28)[_module.__default.safeIndex((_this).fm9(_781_v26, globalState), new BigNumber((_783_v28).length))];
          _nw96[(new BigNumber(20)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(21)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(22)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(23)).toNumber()] = _781_v26;
          _nw96[(new BigNumber(24)).toNumber()] = _781_v26;
          _785_v30 = _nw96;
          _785_v30 = _785_v30;
          let _786_v31;
          _786_v31 = false;
          _786_v31 = ((_767_v12).plus(_767_v12)).isLessThanOrEqualTo(_767_v12);
          _786_v31 = _786_v31;
          let _787_v32;
          let _nw97 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _787_v32 = _nw97;
          let _788_v33;
          _788_v33 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lrwaorxc")).length));
          let _789_v34;
          _789_v34 = _dafny.Seq.of(_788_v33);
          let _index109 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_787_v32).length));
          (_787_v32)[_index109] = (((_this).f17) ? (_789_v34) : (_dafny.Seq.of(_788_v33, _788_v33)));
          let _index110 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_787_v32).length));
          (_787_v32)[_index110] = _789_v34;
          let _790_v35;
          let _nw98 = Array((new BigNumber(28)).toNumber()).fill([]);
          _790_v35 = _nw98;
          _790_v35 = _790_v35;
        }
        let _791_v36;
        _791_v36 = new _dafny.CodePoint('g'.codePointAt(0));
        let _792_v37;
        let _nw99 = new _module.C0();
        _nw99.__ctor((_767_v12).isLessThan(_767_v12), _791_v36, _module.D1.create_DC2(_767_v12, _767_v12, p0), p0);
        _792_v37 = _nw99;
        let _793_v38;
        _793_v38 = _dafny.Seq.of((_this).f14);
        let _794_v39;
        _794_v39 = _dafny.Seq.UnicodeFromString("b");
        let _795_v40;
        let _nw100 = Array((new BigNumber(14)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw100[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(3)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(4)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(6)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(11)).toNumber()] = (_this).f14;
        _nw100[(new BigNumber(12)).toNumber()] = (_793_v38)[_module.__default.safeIndex(new BigNumber((_794_v39).length), new BigNumber((_793_v38).length))];
        _nw100[(new BigNumber(13)).toNumber()] = (_this).f14;
        _795_v40 = _nw100;
        let _796_v41;
        _796_v41 = _795_v40;
        let _797_v42;
        _797_v42 = _dafny.Seq.of(_795_v40, _795_v40, (_796_v41), _795_v40);
        let _798_v43;
        let _nw101 = Array((new BigNumber(14)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = _795_v40;
        _nw101[(_dafny.ONE).toNumber()] = _795_v40;
        _nw101[(new BigNumber(2)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(3)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(4)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(5)).toNumber()] = (_797_v42)[_module.__default.safeIndex(_767_v12, new BigNumber((_797_v42).length))];
        _nw101[(new BigNumber(6)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(7)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(8)).toNumber()] = (_797_v42)[_module.__default.safeIndex(_767_v12, new BigNumber((_797_v42).length))];
        _nw101[(new BigNumber(9)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(10)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(11)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(12)).toNumber()] = _795_v40;
        _nw101[(new BigNumber(13)).toNumber()] = (_797_v42)[_module.__default.safeIndex(_767_v12, new BigNumber((_797_v42).length))];
        _798_v43 = _nw101;
        let _index111 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_798_v43).length));
        (_798_v43)[_index111] = _795_v40;
        let _index112 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_798_v43).length));
        (_798_v43)[_index112] = _795_v40;
      }
      let _799_v44;
      let _init14 = ((_800_p0) => function (_801_i2) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f18), _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f13))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_800_p0,new BigNumber((_dafny.Seq.of((_this).f18)).length))).length)))).length),!(_800_p0));
      })(p0);
      let _nw102 = Array((new BigNumber(4)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw102.length); _i0_14++) {
        _nw102[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _799_v44 = _nw102;
      let _802_v45;
      _802_v45 = new BigNumber(10);
      let _803_v46;
      _803_v46 = _dafny.Map.Empty.slice().updateUnsafe(_799_v44,_802_v45);
      _803_v46 = (_803_v46).update(_799_v44, _module.__default.safeModuloInt(new BigNumber(771), new BigNumber(190)));
      if (_module.__default.fm17(globalState)) {
        let _804_v47;
        let _nw103 = Array((new BigNumber(22)).toNumber()).fill([]);
        _804_v47 = _nw103;
        let _805_v48;
        _805_v48 = _804_v47;
        let _806_v49;
        let _nw104 = Array((new BigNumber(4)).toNumber());
        _nw104[(_dafny.ZERO).toNumber()] = _804_v47;
        _nw104[(_dafny.ONE).toNumber()] = _805_v48;
        _nw104[(new BigNumber(2)).toNumber()] = _805_v48;
        _nw104[(new BigNumber(3)).toNumber()] = _804_v47;
        _806_v49 = _nw104;
        let _rhs89 = _806_v49;
        let _rhs90 = _802_v45;
        let _lhs87 = globalState;
        _806_v49 = _rhs89;
        _lhs87.f7 = _rhs90;
        let _807_v50;
        _807_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_802_v45);
        let _808_v51;
        _808_v51 = _dafny.Seq.UnicodeFromString("pwgijra");
        let _809_v52;
        _809_v52 = _dafny.MultiSet.fromElements((_this).f13);
        let _810_v53;
        _810_v53 = _dafny.Seq.of(_809_v52, _809_v52, _809_v52);
        let _811_v54;
        let _nw105 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _811_v54 = _nw105;
        let _812_v55;
        _812_v55 = _module.D3.create_DC7();
        let _813_v56;
        _813_v56 = _dafny.Map.Empty.slice().updateUnsafe(_811_v54,_812_v55);
        let _814_v57;
        _814_v57 = _dafny.Seq.of(_802_v45);
        let _815_v58;
        _815_v58 = new _dafny.CodePoint('p'.codePointAt(0));
        let _816_v59;
        _816_v59 = _dafny.Seq.of(_811_v54);
        let _817_v60;
        _817_v60 = _module.D3.create_DC8(new BigNumber(698));
        let _818_v61;
        let _nw106 = Array((new BigNumber(26)).toNumber());
        _nw106[(_dafny.ZERO).toNumber()] = _802_v45;
        _nw106[(_dafny.ONE).toNumber()] = new BigNumber(268);
        _nw106[(new BigNumber(2)).toNumber()] = (_802_v45).multipliedBy(_802_v45);
        _nw106[(new BigNumber(3)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(4)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(5)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_802_v45, _802_v45);
        _nw106[(new BigNumber(7)).toNumber()] = (((_807_v50).contains((_this).f17)) ? ((_807_v50).get((_this).f17)) : (new BigNumber((_808_v51).length)));
        _nw106[(new BigNumber(8)).toNumber()] = new BigNumber(((_810_v53)[_module.__default.safeIndex(new BigNumber((_module.__default.fm19(globalState)).cardinality()), new BigNumber((_810_v53).length))]).cardinality());
        _nw106[(new BigNumber(9)).toNumber()] = new BigNumber((_813_v56).length);
        _nw106[(new BigNumber(10)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(11)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(12)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(13)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(14)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(15)).toNumber()] = (_814_v57)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_module.__default.fm20(globalState), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), function (_819_i3) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })).length), new BigNumber((_module.__default.fm20(globalState)).length)), _815_v58)).length), new BigNumber((_814_v57).length))];
        _nw106[(new BigNumber(16)).toNumber()] = (_this).fm8((_this).f18, _module.__default.fm21(globalState), globalState);
        _nw106[(new BigNumber(17)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(18)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(19)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(20)).toNumber()] = (_this).fm8(_module.__default.fm17(globalState), _this.f16, globalState);
        _nw106[(new BigNumber(21)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(22)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(_802_v45)).Union(_dafny.MultiSet.fromElements(_802_v45))).cardinality());
        _nw106[(new BigNumber(23)).toNumber()] = _module.__default.safeDivisionInt(_802_v45, new BigNumber((_816_v59).length));
        _nw106[(new BigNumber(24)).toNumber()] = _802_v45;
        _nw106[(new BigNumber(25)).toNumber()] = (_817_v60).dtor_cf14;
        _818_v61 = _nw106;
        let _index113 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_811_v54).length));
        (_811_v54)[_index113] = ((((_807_v50).contains((_this).f17)) ? ((_807_v50).get((_this).f17)) : (new BigNumber((_814_v57).length)))).minus(_802_v45);
        let _index114 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_811_v54).length));
        (_811_v54)[_index114] = (_802_v45).minus(((p0) ? (_802_v45) : (new BigNumber((_module.__default.fm20(globalState)).length))));
        (globalState).f7 = (_dafny.ZERO).minus((_811_v54)[_module.__default.safeIndex(new BigNumber(374), new BigNumber((_811_v54).length))]);
        (globalState).f7 = _802_v45;
        let _820_v63;
        let _init15 = ((_821_p0, _822_v58, _823_v51) => function (_824_i4) {
          return ((_821_p0) ? (function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_dafny.Map.Empty.slice().updateUnsafe((_823_v51)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_823_v51).length))],!((_this).f17))).Keys.Elements) {
              let _825_v62 = _compr_35;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_823_v51)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_823_v51).length))],!((_this).f17))).contains(_825_v62)) {
                _coll35.add(_825_v62);
              }
            }
            return _coll35;
          }()) : (_dafny.Set.fromElements(_822_v58, _822_v58)));
        })(p0, _815_v58, _808_v51);
        let _nw107 = Array((new BigNumber(19)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw107.length); _i0_15++) {
          _nw107[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _820_v63 = _nw107;
        let _826_v64;
        _826_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(494),_dafny.Seq.UnicodeFromString("uh"));
        let _827_v65;
        _827_v65 = _dafny.Set.fromElements((((_826_v64).contains(_802_v45)) ? ((_826_v64).get(_802_v45)) : (_808_v51)), _dafny.Seq.UnicodeFromString("dgxndwfh"), _808_v51);
        let _828_v66;
        _828_v66 = _827_v65;
        let _829_v67;
        _829_v67 = _dafny.Set.fromElements(!(_module.__default.fm17(globalState)), true, (_this).f13, (_this).f13);
        let _rhs91 = _820_v63;
        let _rhs92 = _dafny.Seq.Concat(_808_v51, _dafny.Seq.Concat(_808_v51, _dafny.Seq.UnicodeFromString("wxpgj")));
        let _rhs93 = (_828_v66);
        let _rhs94 = _811_v54;
        let _rhs95 = _module.__default.safeModuloInt((_802_v45).minus(new BigNumber((_829_v67).length)), (_811_v54)[_module.__default.safeIndex(new BigNumber(374), new BigNumber((_811_v54).length))]);
        _820_v63 = _rhs91;
        _808_v51 = _rhs92;
        _827_v65 = _rhs93;
        _818_v61 = _rhs94;
        _802_v45 = _rhs95;
      } else {
        let _830_v69;
        _830_v69 = _dafny.Seq.of(_802_v45, _802_v45);
        let _831_v70;
        _831_v70 = _dafny.MultiSet.fromElements(new BigNumber((_830_v69).length), (_dafny.ZERO).minus(_802_v45));
        let _832_v71;
        _832_v71 = _dafny.Seq.of(_831_v70);
        let _833_v73;
        _833_v73 = _module.D6.create_DC11(_831_v70);
        let _834_v74;
        _834_v74 = _dafny.MultiSet.fromElements((_833_v73).dtor_cf17);
        let _835_v75;
        _835_v75 = _dafny.Map.Empty.slice().updateUnsafe((function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_832_v71).Elements) {
            let _836_v68 = _compr_36;
            if (_dafny.Seq.contains(_832_v71, _836_v68)) {
              _coll36.push([_836_v68,true]);
            }
          }
          return _coll36;
        }()).Merge(function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_834_v74).Elements) {
            let _837_v72 = _compr_37;
            if ((_834_v74).contains(_837_v72)) {
              _coll37.push([_837_v72,(_this).f13]);
            }
          }
          return _coll37;
        }()),false);
        let _838_v76;
        _838_v76 = _dafny.Map.Empty.slice().updateUnsafe(_831_v70,(_this).f13);
        _835_v75 = (_835_v75).update(_838_v76, p0);
        let _839_v77;
        _839_v77 = _module.D2.create_DC3((_this).f14);
        let _source15 = _839_v77;
        if (_source15.is_DC4) {
          let _840___mcc_h0 = (_source15).cf6;
          let _841___mcc_h1 = (_source15).cf7;
          let _842___mcc_h2 = (_source15).cf8;
          let _843___mcc_h3 = (_source15).cf9;
          let _844___mcc_h4 = (_source15).cf10;
          let _845_cf10 = _844___mcc_h4;
          let _846_cf9 = _843___mcc_h3;
          let _847_cf8 = _842___mcc_h2;
          let _848_cf7 = _841___mcc_h1;
          let _849_cf6 = _840___mcc_h0;
          let _850_v78;
          _850_v78 = _dafny.Seq.UnicodeFromString("fkeo");
          let _851_v79;
          _851_v79 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("gn"), _850_v78);
          let _852_v80;
          _852_v80 = _851_v79;
          let _853_v81;
          _853_v81 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_852_v80, _852_v80, _module.__default.fm22(_846_cf9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-793)), function (_854_i5) {
            return new BigNumber(-328);
          }), globalState), _852_v80, _851_v79),p0);
          let _855_v82;
          _855_v82 = _dafny.MultiSet.fromElements(_852_v80, _852_v80);
          let _856_v83;
          _856_v83 = _dafny.Seq.of(_846_cf9);
          _853_v81 = (_853_v81).update((_855_v82).Intersect(_855_v82), (_856_v83)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_847_cf8)).cardinality()), new BigNumber((_856_v83).length))]);
          let _857_v84;
          _857_v84 = new _dafny.CodePoint('b'.codePointAt(0));
          let _858_v85;
          _858_v85 = _dafny.Set.fromElements(_849_cf6, _849_cf6);
          let _859_v86;
          let _nw108 = new _module.C0();
          _nw108.__ctor((_this).f18, _857_v84, _this.f12, !(_858_v85).contains(_849_cf6));
          _859_v86 = _nw108;
          let _860_v87;
          _860_v87 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("x")), _850_v78, _850_v78, _dafny.Seq.Concat((_this.f16).dtor_cf1, _dafny.Seq.UnicodeFromString("te")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_861_v86) => function (_862_i6) {
            return (_module.D3.create_DC5((_861_v86).f20)).dtor_cf11;
          })(_859_v86)));
          let _863_v88;
          _863_v88 = _dafny.Map.Empty.slice().updateUnsafe(_847_cf8,_860_v87);
          _860_v87 = (((_863_v88).contains(_802_v45)) ? ((_863_v88).get(_802_v45)) : (_860_v87));
          _846_cf9 = _848_cf7;
        } else {
          let _864___mcc_h5 = (_source15).cf5;
          let _865_cf5 = _864___mcc_h5;
          let _rhs96 = _802_v45;
          let _rhs97 = (_dafny.ZERO).minus(_802_v45);
          let _lhs88 = globalState;
          _lhs88.f7 = _rhs96;
          r0 = _rhs97;
          let _866_v89;
          _866_v89 = _dafny.Seq.of((_this).f13);
          (globalState).f2 = _866_v89;
          let _867_v90;
          _867_v90 = true;
          let _868_v91;
          _868_v91 = new _dafny.CodePoint('x'.codePointAt(0));
          let _869_v92;
          let _init16 = function (_870_i7) {
            return _this.f12;
          };
          let _nw109 = Array((new BigNumber(29)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw109.length); _i0_16++) {
            _nw109[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _869_v92 = _nw109;
          let _index115 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_869_v92).length));
          (_869_v92)[_index115] = _this.f12;
          let _871_v93;
          let _init17 = ((_872_v45) => function (_873_i8) {
            return (_873_i8).plus(_872_v45);
          })(_802_v45);
          let _nw110 = Array((new BigNumber(28)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw110.length); _i0_17++) {
            _nw110[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _871_v93 = _nw110;
          let _index116 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_871_v93).length));
          (_871_v93)[_index116] = ((_dafny.ZERO).minus(_802_v45)).multipliedBy(_802_v45);
          let _pat_let_tv26 = _802_v45;
          let _index117 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_869_v92).length));
          let _index118 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_871_v93).length));
          let _rhs98 = _867_v90;
          let _rhs99 = _802_v45;
          let _rhs100 = _868_v91;
          let _rhs101 = function (_pat_let18_0) {
            return function (_874_dt__update__tmp_h1) {
              return function (_pat_let19_0) {
                return function (_875_dt__update_hcf4_h0) {
                  return function (_pat_let20_0) {
                    return function (_876_dt__update_hcf2_h0) {
                      return _module.D1.create_DC2(_876_dt__update_hcf2_h0, (_874_dt__update__tmp_h1).dtor_cf3, _875_dt__update_hcf4_h0);
                    }(_pat_let20_0);
                  }(_pat_let_tv26);
                }(_pat_let19_0);
              }((_this).f17);
            }(_pat_let18_0);
          }(_this.f12);
          let _rhs102 = _802_v45;
          let _lhs89 = _869_v92;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_869_v92).length));
          let _lhs91 = _871_v93;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_871_v93).length));
          _867_v90 = _rhs98;
          _802_v45 = _rhs99;
          _868_v91 = _rhs100;
          _lhs89[_lhs90] = _rhs101;
          _lhs91[_lhs92] = _rhs102;
          (globalState).f7 = new BigNumber(785);
        }
        let _877_v94;
        let _nw111 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _877_v94 = _nw111;
        let _index119 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_877_v94).length));
        (_877_v94)[_index119] = (_802_v45).plus(_802_v45);
        let _index120 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_877_v94).length));
        (_877_v94)[_index120] = (new BigNumber(520)).multipliedBy(_802_v45);
        let _878_v95;
        _878_v95 = new _dafny.CodePoint('t'.codePointAt(0));
        let _879_v96;
        let _nw112 = new _module.C0();
        _nw112.__ctor((_this).f13, _878_v95, _this.f12, _module.__default.fm17(globalState));
        _879_v96 = _nw112;
        _879_v96 = _879_v96;
      }
      (_this).f12 = function (_pat_let21_0) {
        return function (_880_dt__update__tmp_h2) {
          return function (_pat_let22_0) {
            return function (_881_dt__update_hcf4_h1) {
              return _module.D1.create_DC2((_880_dt__update__tmp_h2).dtor_cf2, (_880_dt__update__tmp_h2).dtor_cf3, _881_dt__update_hcf4_h1);
            }(_pat_let22_0);
          }((_this).f13);
        }(_pat_let21_0);
      }(_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f13)).length), _802_v45, (_this).f13));
      if (!((_802_v45).isLessThan(_802_v45))) {
        let _882_v97;
        _882_v97 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f18);
        let _883_v98;
        let _nw113 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _883_v98 = _nw113;
        let _884_v99;
        _884_v99 = _dafny.MultiSet.fromElements(_883_v98);
        let _885_v100;
        _885_v100 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f17) ? (_882_v97) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f18))),_module.__default.safeModuloInt(new BigNumber((_884_v99).cardinality()), new BigNumber(527)));
        let _886_v101;
        _886_v101 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23((_this).f13, globalState),_885_v100);
        let _887_v102;
        _887_v102 = new _dafny.CodePoint('h'.codePointAt(0));
        let _888_v103;
        _888_v103 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,_887_v102);
        let _889_v105;
        _889_v105 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,(_this).f17);
        _885_v100 = (((_886_v101).contains((_888_v103).Merge(function () {
          let _coll40 = new _dafny.Map();
          for (const _compr_40 of (_889_v105).Keys.Elements) {
            let _892_v104 = _compr_40;
            if ((_889_v105).contains(_892_v104)) {
              _coll40.push([(_892_v104).minus(new BigNumber((_dafny.Seq.of((_this).f13, (_this).f17)).length)),_887_v102]);
            }
          }
          return _coll40;
        }()))) ? ((_886_v101).get((_888_v103).Merge(function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of (_889_v105).Keys.Elements) {
            let _890_v104 = _compr_38;
            if ((_889_v105).contains(_890_v104)) {
              _coll38.push([(_890_v104).minus(new BigNumber((_dafny.Seq.of((_this).f13, (_this).f17)).length)),_887_v102]);
            }
          }
          return _coll38;
        }()))) : (function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of (_885_v100).Keys.Elements) {
            let _891_v106 = _compr_39;
            if ((_885_v100).contains(_891_v106)) {
              _coll39.push([_891_v106,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_887_v102,(_this).f13)).length)]);
            }
          }
          return _coll39;
        }()));
        let _893_v107;
        let _nw114 = new _module.C0();
        _nw114.__ctor((_this).f18, _887_v102, _this.f12, (_this).f17);
        _893_v107 = _nw114;
        let _894_v108;
        _894_v108 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,_802_v45);
        let _895_v109;
        _895_v109 = _dafny.Seq.of((_this).f13, _893_v107.f19);
        let _rhs103 = _module.__default.fm17(globalState);
        let _rhs104 = _802_v45;
        let _rhs105 = _module.__default.fm24(_802_v45, globalState);
        let _rhs106 = (_module.__default.fm0(_895_v109, (_this).f17, (_this).f18, _802_v45, globalState)).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), function (_896_i9) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("e"))).length));
        let _lhs93 = _893_v107;
        _lhs93.f19 = _rhs103;
        _802_v45 = _rhs104;
        _894_v108 = _rhs105;
        r0 = _rhs106;
        (_893_v107).f19 = _893_v107.f19;
        _889_v105 = (_889_v105).update(_802_v45, (_this).f17);
      } else {
        let _index121 = _module.__default.safeIndex(new BigNumber(83), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index121] = p0;
        let _index122 = _module.__default.safeIndex(new BigNumber(83), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index122] = p0;
        let _897_v110;
        _897_v110 = new _dafny.CodePoint('w'.codePointAt(0));
        let _898_v111;
        let _nw115 = new _module.C0();
        _nw115.__ctor((!(((_this).f14)[_module.__default.safeIndex(new BigNumber(83), new BigNumber(((_this).f14).length))])) || ((_this).f18), _897_v110, _this.f12, (_this).f18);
        _898_v111 = _nw115;
        let _899_v112;
        _899_v112 = _dafny.MultiSet.fromElements(true);
        let _900_v113;
        _900_v113 = _dafny.Seq.of(_899_v112);
        _900_v113 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_901_v112) => function (_902_i10) {
          return _901_v112;
        })(_899_v112));
        let _903_v114;
        _903_v114 = _dafny.MultiSet.fromElements(_802_v45, _802_v45);
        let _904_v115;
        _904_v115 = _dafny.Seq.of((_this).f17, (_this).f17);
        let _905_v116;
        _905_v116 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f17);
        let _906_v117;
        _906_v117 = _dafny.Set.fromElements(_module.__default.fm17(globalState), (((_905_v116).contains((_this).f17)) ? ((_905_v116).get((_this).f17)) : (true)));
        let _907_v118;
        _907_v118 = _dafny.Map.Empty.slice().updateUnsafe(_898_v111.f19,_802_v45);
        let _908_v119;
        let _nw116 = Array((new BigNumber(26)).toNumber());
        _nw116[(_dafny.ZERO).toNumber()] = _802_v45;
        _nw116[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_802_v45);
        _nw116[(new BigNumber(2)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(3)).toNumber()] = (((_903_v114).contains(_802_v45)) ? ((_903_v114).get(_802_v45)) : (new BigNumber((_dafny.MultiSet.fromElements((_this).f13)).cardinality())));
        _nw116[(new BigNumber(4)).toNumber()] = new BigNumber((_904_v115).length);
        _nw116[(new BigNumber(5)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(6)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(7)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(8)).toNumber()] = new BigNumber((_906_v117).length);
        _nw116[(new BigNumber(9)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(10)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(11)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(12)).toNumber()] = new BigNumber(-840);
        _nw116[(new BigNumber(13)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(14)).toNumber()] = new BigNumber((_907_v118).length);
        _nw116[(new BigNumber(15)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(16)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(17)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(18)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(19)).toNumber()] = new BigNumber(787);
        _nw116[(new BigNumber(20)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(21)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(22)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(23)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(24)).toNumber()] = _802_v45;
        _nw116[(new BigNumber(25)).toNumber()] = _802_v45;
        _908_v119 = _nw116;
        let _909_v120;
        let _init18 = function (_910_i11) {
          return (_this).f13;
        };
        let _nw117 = Array((new BigNumber(20)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw117.length); _i0_18++) {
          _nw117[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _909_v120 = _nw117;
        let _911_v121;
        _911_v121 = _module.D2.create_DC4(_908_v119, (_this).f17, (_dafny.ZERO).minus(_802_v45), _898_v111.f19, _909_v120);
        (_898_v111).f19 = (((p0) && (_898_v111.f19)) ? ((_911_v121).dtor_cf7) : (true));
        _802_v45 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_802_v45, _802_v45)));
      }
      if (p0) {
        let _912_v122;
        _912_v122 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
        r0 = (_dafny.ZERO).minus((((_this).f18) ? (_802_v45) : (_module.__default.safeDivisionInt(new BigNumber((_912_v122).length), _802_v45))));
        let _913_v123;
        _913_v123 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,(_this).f17);
        let _914_v125;
        _914_v125 = _dafny.Seq.of(_913_v123, (_913_v123).Merge(_913_v123), _913_v123, (function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of _dafny.IntegerRange(new BigNumber(627), new BigNumber(879))) {
            let _915_v124 = _compr_41;
            if (((new BigNumber(627)).isLessThanOrEqualTo(_915_v124)) && ((_915_v124).isLessThan(new BigNumber(879)))) {
              _coll41.push([_module.__default.safeDivisionInt(_915_v124, _802_v45),!((_this).f13)]);
            }
          }
          return _coll41;
        }()).update(_802_v45, (_this).f18));
        let _916_v126;
        _916_v126 = _dafny.Seq.UnicodeFromString("fonto");
        _914_v125 = _module.__default.fm25(_916_v126, p0, (_dafny.ZERO).minus(_802_v45), globalState);
        let _917_v127;
        _917_v127 = new _dafny.CodePoint('a'.codePointAt(0));
        let _pat_let_tv27 = _802_v45;
        let _pat_let_tv28 = _802_v45;
        let _pat_let_tv29 = _802_v45;
        let _918_v128;
        let _nw118 = new _module.C0();
        _nw118.__ctor(p0, _917_v127, function (_pat_let23_0) {
          return function (_921_dt__update__tmp_h4) {
            return function (_pat_let26_0) {
              return function (_922_dt__update_hcf3_h0) {
                return function (_pat_let27_0) {
                  return function (_923_dt__update_hcf2_h2) {
                    return _module.D1.create_DC2(_923_dt__update_hcf2_h2, _922_dt__update_hcf3_h0, (_921_dt__update__tmp_h4).dtor_cf4);
                  }(_pat_let27_0);
                }(new BigNumber(-76));
              }(_pat_let26_0);
            }((_module.D1.create_DC2(_pat_let_tv28, _pat_let_tv29, (_this).f13)).dtor_cf3);
          }(_pat_let23_0);
        }(function (_pat_let24_0) {
          return function (_919_dt__update__tmp_h3) {
            return function (_pat_let25_0) {
              return function (_920_dt__update_hcf2_h1) {
                return _module.D1.create_DC2(_920_dt__update_hcf2_h1, (_919_dt__update__tmp_h3).dtor_cf3, (_919_dt__update__tmp_h3).dtor_cf4);
              }(_pat_let25_0);
            }(_pat_let_tv27);
          }(_pat_let24_0);
        }(_this.f12)), (_this).f13);
        _918_v128 = _nw118;
        let _924_v129;
        _924_v129 = _dafny.Seq.of(!(p0));
        let _925_v130;
        let _nw119 = Array((new BigNumber(18)).toNumber());
        _nw119[(_dafny.ZERO).toNumber()] = ((p0) ? (_802_v45) : (new BigNumber((_924_v129).length)));
        _nw119[(_dafny.ONE).toNumber()] = new BigNumber(923);
        _nw119[(new BigNumber(2)).toNumber()] = new BigNumber((_916_v126).length);
        _nw119[(new BigNumber(3)).toNumber()] = new BigNumber(581);
        _nw119[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(_802_v45, _802_v45);
        _nw119[(new BigNumber(5)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(6)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(7)).toNumber()] = new BigNumber((_module.__default.fm26(_802_v45, _802_v45, p0, globalState)).length);
        _nw119[(new BigNumber(8)).toNumber()] = ((p0) ? (_802_v45) : (_802_v45));
        _nw119[(new BigNumber(9)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(10)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(11)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(12)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(13)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(14)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(15)).toNumber()] = _802_v45;
        _nw119[(new BigNumber(16)).toNumber()] = (_802_v45).multipliedBy(_802_v45);
        _nw119[(new BigNumber(17)).toNumber()] = _802_v45;
        _925_v130 = _nw119;
        let _index123 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_925_v130).length));
        (_925_v130)[_index123] = _802_v45;
        let _index124 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_925_v130).length));
        (_925_v130)[_index124] = _802_v45;
        let _926_v131;
        let _nw120 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
        _926_v131 = _nw120;
        let _927_v132;
        _927_v132 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,(_925_v130)[_module.__default.safeIndex(new BigNumber(639), new BigNumber((_925_v130).length))]);
        let _928_v133;
        _928_v133 = _dafny.Seq.of(new BigNumber((_927_v132).length));
        let _929_v134;
        _929_v134 = _dafny.Set.fromElements(_928_v133);
        let _index125 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_926_v131).length));
        (_926_v131)[_index125] = (_929_v134).Union(_929_v134);
        let _index126 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_926_v131).length));
        (_926_v131)[_index126] = _929_v134;
      } else {
        let _930_v135;
        let _nw121 = Array((new BigNumber(28)).toNumber()).fill([]);
        _930_v135 = _nw121;
        let _931_v136;
        let _nw122 = Array((_dafny.ONE).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = _this.f12;
        _931_v136 = _nw122;
        let _index127 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_930_v135).length));
        (_930_v135)[_index127] = _931_v136;
        let _index128 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_930_v135).length));
        (_930_v135)[_index128] = _931_v136;
        let _932_v137;
        _932_v137 = _module.D6.create_DC12(p0, _module.__default.fm20(globalState), _802_v45, (_this).f13);
        let _source16 = _932_v137;
        if (_source16.is_DC12) {
          let _933___mcc_h6 = (_source16).cf18;
          let _934___mcc_h7 = (_source16).cf19;
          let _935___mcc_h8 = (_source16).cf20;
          let _936___mcc_h9 = (_source16).cf21;
          let _937_cf21 = _936___mcc_h9;
          let _938_cf20 = _935___mcc_h8;
          let _939_cf19 = _934___mcc_h7;
          let _940_cf18 = _933___mcc_h6;
          let _941_v138;
          let _init19 = ((_942_cf19) => function (_943_i12) {
            return _942_cf19;
          })(_939_cf19);
          let _nw123 = Array((new BigNumber(3)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw123.length); _i0_19++) {
            _nw123[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _941_v138 = _nw123;
          let _index129 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_941_v138).length));
          (_941_v138)[_index129] = _939_cf19;
          let _index130 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_941_v138).length));
          (_941_v138)[_index130] = _939_cf19;
          (_this).m12(globalState);
          let _944_v139;
          _944_v139 = new _dafny.CodePoint('b'.codePointAt(0));
          let _945_v140;
          let _nw124 = new _module.C0();
          _nw124.__ctor((_this).f18, _944_v139, _this.f12, (_802_v45).isLessThan(_802_v45));
          _945_v140 = _nw124;
          let _946_v141;
          let _nw125 = new _module.C0();
          _nw125.__ctor(_940_cf18, _944_v139, _this.f12, _940_cf18);
          _946_v141 = _nw125;
        } else if (_source16.is_DC13) {
          let _947___mcc_h10 = (_source16).cf22;
          let _948___mcc_h11 = (_source16).cf23;
          let _949___mcc_h12 = (_source16).cf24;
          let _950___mcc_h13 = (_source16).cf25;
          let _951___mcc_h14 = (_source16).cf26;
          let _952_cf26 = _951___mcc_h14;
          let _953_cf25 = _950___mcc_h13;
          let _954_cf24 = _949___mcc_h12;
          let _955_cf23 = _948___mcc_h11;
          let _956_cf22 = _947___mcc_h10;
          let _957_v142;
          let _nw126 = Array((new BigNumber(12)).toNumber()).fill(false);
          _957_v142 = _nw126;
          _957_v142 = _957_v142;
          let _958_v143;
          _958_v143 = _dafny.Seq.of(_954_cf24);
          _958_v143 = _dafny.Seq.Concat(_dafny.Seq.of(_802_v45, _953_cf25), _958_v143);
          let _959_v144;
          _959_v144 = new _dafny.CodePoint('t'.codePointAt(0));
          let _960_v145;
          _960_v145 = _dafny.Seq.UnicodeFromString("tpil");
          _952_cf26 = !_dafny.Seq.contains(_dafny.Seq.Concat(_960_v145, _960_v145), _959_v144);
          _952_cf26 = ((!((_this).f17)) ? ((_this).f18) : ((_802_v45).isLessThanOrEqualTo(_802_v45)));
        } else if (_source16.is_DC11) {
          let _961___mcc_h15 = (_source16).cf17;
          let _962_cf17 = _961___mcc_h15;
          let _963_v146;
          let _nw127 = Array((new BigNumber(28)).toNumber()).fill([]);
          _963_v146 = _nw127;
          let _964_v147;
          let _nw128 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _964_v147 = _nw128;
          let _index131 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_963_v146).length));
          (_963_v146)[_index131] = _964_v147;
          let _index132 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_963_v146).length));
          (_963_v146)[_index132] = _964_v147;
          let _index133 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_964_v147).length));
          (_964_v147)[_index133] = _802_v45;
          let _965_v148;
          let _nw129 = Array((new BigNumber(20)).toNumber()).fill(_module.D6.Default());
          _965_v148 = _nw129;
          let _966_v149;
          _966_v149 = _dafny.Seq.of(_965_v148);
          let _967_v150;
          _967_v150 = _dafny.Seq.of(_966_v149, _966_v149, _966_v149);
          let _index134 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_964_v147).length));
          (_964_v147)[_index134] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_802_v45, new BigNumber(((_967_v150)[_module.__default.safeIndex(_802_v45, new BigNumber((_967_v150).length))]).length)), _802_v45);
          let _968_v151;
          _968_v151 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _969_v152;
          _969_v152 = _dafny.Map.Empty.slice().updateUnsafe(true,(_964_v147)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_964_v147).length))]);
          let _index135 = _module.__default.safeIndex(new BigNumber(711), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index135] = ((((_969_v152).contains((_this).f18)) ? ((_969_v152).get((_this).f18)) : (_802_v45))).isLessThan((_dafny.ZERO).minus(new BigNumber((_968_v151).length)));
          let _index136 = _module.__default.safeIndex(new BigNumber(711), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index136] = p0;
          let _970_v153;
          _970_v153 = _dafny.Seq.UnicodeFromString("x");
          let _971_v154;
          _971_v154 = _dafny.Map.Empty.slice().updateUnsafe(_970_v153,_964_v147);
          (globalState).f7 = new BigNumber(((_971_v154).update(_dafny.Seq.Concat(_970_v153, _970_v153), (_963_v146)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_963_v146).length))])).length);
        } else {
          let _972___mcc_h16 = (_source16).cf27;
          let _973_cf27 = _972___mcc_h16;
          (globalState).f7 = _802_v45;
          let _974_v155;
          _974_v155 = _dafny.Seq.of((_this).f14, (_this).f14);
          let _975_v156;
          _975_v156 = _dafny.Seq.UnicodeFromString("xwoa");
          let _976_v157;
          _976_v157 = _dafny.Map.Empty.slice().updateUnsafe((_974_v155)[_module.__default.safeIndex(new BigNumber((_975_v156).length), new BigNumber((_974_v155).length))],_dafny.Seq.of((_dafny.ZERO).minus(_802_v45), _802_v45));
          _976_v157 = (_976_v157).update((_this).f14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(315)), function (_977_i13) {
            return new BigNumber(447);
          }));
          let _978_v158;
          _978_v158 = true;
          let _979_v159;
          let _init20 = ((_980_v45) => function (_981_i14) {
            return _dafny.Seq.of(_980_v45);
          })(_802_v45);
          let _nw130 = Array((new BigNumber(26)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw130.length); _i0_20++) {
            _nw130[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _979_v159 = _nw130;
          let _982_v160;
          _982_v160 = _dafny.Seq.of(_802_v45);
          let _index137 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_979_v159).length));
          (_979_v159)[_index137] = _982_v160;
          let _983_v161;
          _983_v161 = _dafny.Seq.of(_module.__default.fm17(globalState));
          let _984_v162;
          _984_v162 = _dafny.Map.Empty.slice().updateUnsafe(_802_v45,p0);
          let _985_v163;
          _985_v163 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_983_v161, _module.__default.safeIndex(new BigNumber((_984_v162).length), new BigNumber((_983_v161).length)), true));
          let _index138 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_979_v159).length));
          let _rhs107 = ((_this).f17) && (!(p0));
          let _rhs108 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_802_v45,p0)).length), new BigNumber((_985_v163).cardinality()), _802_v45, _802_v45), _dafny.Seq.update(_module.__default.fm27(globalState), _module.__default.safeIndex(_802_v45, new BigNumber((_module.__default.fm27(globalState)).length)), _802_v45));
          let _rhs109 = ((_802_v45).minus(_802_v45)).minus(_module.__default.safeDivisionInt(_802_v45, _802_v45));
          let _lhs94 = _979_v159;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_979_v159).length));
          _978_v158 = _rhs107;
          _lhs94[_lhs95] = _rhs108;
          _802_v45 = _rhs109;
          let _986_v164;
          let _nw131 = new _module.C0();
          _nw131.__ctor(p0, new _dafny.CodePoint('n'.codePointAt(0)), _this.f12, _978_v158);
          _986_v164 = _nw131;
          let _987_v165;
          _987_v165 = _dafny.Map.Empty.slice().updateUnsafe(_986_v164,_975_v156);
          _987_v165 = (_987_v165).update(_986_v164, _dafny.Seq.Concat(_975_v156, _dafny.Seq.UnicodeFromString("e")));
        }
        let _988_v166;
        _988_v166 = _dafny.MultiSet.fromElements((_this).f13, (_this).f18);
        let _index139 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index139] = (_this).f17;
        let _989_v167;
        _989_v167 = false;
        let _990_v168;
        _990_v168 = new _dafny.CodePoint('b'.codePointAt(0));
        let _991_v169;
        _991_v169 = _dafny.Map.Empty.slice().updateUnsafe(_990_v168,(_this).f18);
        let _index140 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
        let _rhs110 = _module.__default.fm28((((_991_v169).contains(_990_v168)) ? ((_991_v169).get(_990_v168)) : ((_this).f18)), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ybya")).length), _802_v45)), !((_this).f17) || (!(p0)), globalState);
        let _rhs111 = (_this).f13;
        let _rhs112 = _module.__default.fm17(globalState);
        let _rhs113 = ((_802_v45).plus(_802_v45)).isLessThan(_802_v45);
        let _rhs114 = (_this).f13;
        let _lhs96 = (_this).f14;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
        _988_v166 = _rhs110;
        _lhs96[_lhs97] = _rhs111;
        _989_v167 = _rhs112;
        _989_v167 = _rhs113;
        _989_v167 = _rhs114;
        if (((_this).f14)[_module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length))]) {
          let _992_v170;
          let _init21 = ((_993_v45) => function (_994_i15) {
            return (_994_i15).plus(_993_v45);
          })(_802_v45);
          let _nw132 = Array((new BigNumber(25)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw132.length); _i0_21++) {
            _nw132[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _992_v170 = _nw132;
          let _index141 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_992_v170).length));
          (_992_v170)[_index141] = ((_this).fm15(globalState)).minus((_this).fm9(_990_v168, globalState));
          let _index142 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_992_v170).length));
          (_992_v170)[_index142] = _802_v45;
          let _index143 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index143] = _989_v167;
          let _995_v171;
          _995_v171 = _dafny.Seq.UnicodeFromString("nsq");
          _995_v171 = _dafny.Seq.update(_995_v171, _module.__default.safeIndex(_802_v45, new BigNumber((_995_v171).length)), _990_v168);
          let _996_v172;
          _996_v172 = _module.D6.create_DC13(p0, p0, _802_v45, _802_v45, _module.__default.fm17(globalState));
          let _index144 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
          let _rhs115 = ((p0) ? (_989_v167) : (((_992_v170)[_module.__default.safeIndex(new BigNumber(362), new BigNumber((_992_v170).length))]).isLessThan(_802_v45)));
          let _rhs116 = ((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(_992_v170, _992_v170)).length)).multipliedBy(_802_v45)))).plus((_996_v172).dtor_cf25);
          let _lhs98 = (_this).f14;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(413), new BigNumber(((_this).f14).length));
          let _lhs100 = globalState;
          _lhs98[_lhs99] = _rhs115;
          _lhs100.f7 = _rhs116;
          let _997_v173;
          let _nw133 = Array((new BigNumber(23)).toNumber()).fill(false);
          _997_v173 = _nw133;
          let _998_v174;
          _998_v174 = _dafny.Map.Empty.slice().updateUnsafe(_989_v167,_997_v173);
          _998_v174 = ((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_997_v173)).Merge(_998_v174)).Merge(_998_v174);
        } else {
          let _999_v175;
          _999_v175 = _dafny.Seq.UnicodeFromString("cmyb");
          let _pat_let_tv30 = _999_v175;
          let _1000_v176;
          _1000_v176 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29((function (_pat_let28_0) {
            return function (_1001_dt__update__tmp_h5) {
              return function (_pat_let29_0) {
                return function (_1002_dt__update_hcf19_h0) {
                  return function (_pat_let30_0) {
                    return function (_1003_dt__update_hcf21_h0) {
                      return _module.D6.create_DC12((_1001_dt__update__tmp_h5).dtor_cf18, _1002_dt__update_hcf19_h0, (_1001_dt__update__tmp_h5).dtor_cf20, _1003_dt__update_hcf21_h0);
                    }(_pat_let30_0);
                  }((_this).f17);
                }(_pat_let29_0);
              }(_pat_let_tv30);
            }(_pat_let28_0);
          }(_932_v137)).dtor_cf20, globalState),_802_v45);
          let _1004_v177;
          _1004_v177 = _dafny.Seq.of(_989_v167, !(p0), true, !(p0));
          _1000_v176 = (_1000_v176).update(_1004_v177, _802_v45);
          let _1005_v178;
          _1005_v178 = _dafny.Map.Empty.slice().updateUnsafe(_989_v167,_999_v175);
          let _1006_v179;
          let _nw134 = Array((new BigNumber(13)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_999_v175, _999_v175);
          _nw134[(_dafny.ONE).toNumber()] = _999_v175;
          _nw134[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_999_v175, _999_v175);
          _nw134[(new BigNumber(3)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(4)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(5)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(6)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(7)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(8)).toNumber()] = (((_1005_v178).contains(p0)) ? ((_1005_v178).get(p0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_1007_v168) => function (_1008_i16) {
            return _1007_v168;
          })(_990_v168))));
          _nw134[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_999_v175, _dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), ((_1009_v168) => function (_1010_i17) {
            return _1009_v168;
          })(_990_v168)));
          _nw134[(new BigNumber(10)).toNumber()] = (((_this).f13) ? (_999_v175) : (_dafny.Seq.UnicodeFromString("plvqc")));
          _nw134[(new BigNumber(11)).toNumber()] = _999_v175;
          _nw134[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_999_v175, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-228)), ((_1011_v168) => function (_1012_i18) {
            return _1011_v168;
          })(_990_v168)));
          _1006_v179 = _nw134;
          let _index145 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1006_v179).length));
          (_1006_v179)[_index145] = _dafny.Seq.UnicodeFromString("yhhsck");
          let _index146 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1006_v179).length));
          let _rhs117 = _990_v168;
          let _rhs118 = _999_v175;
          let _rhs119 = _999_v175;
          let _lhs101 = _1006_v179;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1006_v179).length));
          _990_v168 = _rhs117;
          _999_v175 = _rhs118;
          _lhs101[_lhs102] = _rhs119;
          (globalState).f7 = _module.__default.safeModuloInt((new BigNumber(((_1006_v179)[_module.__default.safeIndex(new BigNumber(367), new BigNumber((_1006_v179).length))]).length)).multipliedBy(_802_v45), _802_v45);
          _989_v167 = _989_v167;
          _989_v167 = false;
        }
        let _1013_v180;
        _1013_v180 = _dafny.MultiSet.fromElements(_802_v45, new BigNumber(882), _802_v45, _802_v45);
        let _1014_v181;
        _1014_v181 = _dafny.Seq.of(_802_v45, (((_1013_v180).contains(_802_v45)) ? ((_1013_v180).get(_802_v45)) : (_802_v45)), _802_v45);
        _1014_v181 = _dafny.Seq.update(_module.__default.fm27(globalState), _module.__default.safeIndex((_1014_v181)[_module.__default.safeIndex(_802_v45, new BigNumber((_1014_v181).length))], new BigNumber((_module.__default.fm27(globalState)).length)), _802_v45);
      }
      let _1015_v182;
      _1015_v182 = _dafny.Seq.of((_this).f17);
      r0 = _module.__default.fm0(_1015_v182, (_this).f13, (_this).f13, _802_v45, globalState);
      return r0;
    }
    m7(globalState) {
      let _this = this;
      if ((_this).f17) {
        let _1016_v0;
        _1016_v0 = new BigNumber(554);
        let _1017_v1;
        _1017_v1 = _module.D6.create_DC13(!((_this).f17), !((new BigNumber(383)).isLessThanOrEqualTo(new BigNumber(272))), _1016_v0, _1016_v0, (_this).f17);
        _1017_v1 = _1017_v1;
        let _1018_v2;
        _1018_v2 = _dafny.Seq.of(new BigNumber(((_this).fm10((_this.f12).dtor_cf4, _module.__default.fm17(globalState), _dafny.Seq.update(_dafny.Seq.of((_this).f13, (_this).f13), _module.__default.safeIndex(_1016_v0, new BigNumber((_dafny.Seq.of((_this).f13, (_this).f13)).length)), (_this).f13), globalState)).length));
        let _1019_v3;
        _1019_v3 = _dafny.Seq.of((_this).f18);
        let _1020_v4;
        _1020_v4 = _dafny.Seq.UnicodeFromString("xdavfw");
        let _1021_v5;
        _1021_v5 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1022_v6;
        _1022_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1020_v4,_1021_v5);
        let _1023_v7;
        let _nw135 = Array((new BigNumber(17)).toNumber());
        _nw135[(_dafny.ZERO).toNumber()] = (((_this).f17) ? (_1016_v0) : (_1016_v0));
        _nw135[(_dafny.ONE).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(2)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(3)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(4)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(5)).toNumber()] = new BigNumber((_1018_v2).length);
        _nw135[(new BigNumber(6)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(7)).toNumber()] = (_1016_v0).plus(_1016_v0);
        _nw135[(new BigNumber(8)).toNumber()] = new BigNumber((_1019_v3).length);
        _nw135[(new BigNumber(9)).toNumber()] = (_1018_v2)[_module.__default.safeIndex(_1016_v0, new BigNumber((_1018_v2).length))];
        _nw135[(new BigNumber(10)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(11)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(12)).toNumber()] = new BigNumber(-137);
        _nw135[(new BigNumber(13)).toNumber()] = new BigNumber((_1022_v6).length);
        _nw135[(new BigNumber(14)).toNumber()] = _1016_v0;
        _nw135[(new BigNumber(15)).toNumber()] = new BigNumber(440);
        _nw135[(new BigNumber(16)).toNumber()] = (_1016_v0).minus(_1016_v0);
        _1023_v7 = _nw135;
        let _index147 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1023_v7).length));
        (_1023_v7)[_index147] = new BigNumber(((_module.D1.create_DC1(_1020_v4)).dtor_cf1).length);
        let _index148 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1023_v7).length));
        (_1023_v7)[_index148] = _1016_v0;
        let _1024_v8;
        _1024_v8 = _dafny.MultiSet.fromElements(_1021_v5, _1021_v5, _1021_v5);
        let _index149 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1023_v7).length));
        let _index150 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1023_v7).length));
        let _rhs120 = _1016_v0;
        let _rhs121 = (new BigNumber((_dafny.Seq.Concat(_1020_v4, _1020_v4)).length)).plus(_1016_v0);
        let _rhs122 = (_dafny.MultiSet.fromElements(_1021_v5)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1020_v4, _dafny.Seq.of(_1021_v5, _1021_v5, _1021_v5))));
        let _lhs103 = _1023_v7;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_1023_v7).length));
        let _lhs105 = _1023_v7;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1023_v7).length));
        _lhs103[_lhs104] = _rhs120;
        _lhs105[_lhs106] = _rhs121;
        _1024_v8 = _rhs122;
        (globalState).f7 = (_1023_v7)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_1023_v7).length))];
        let _1025_v9;
        _1025_v9 = true;
        _1025_v9 = (_this).f13;
        let _1026_v10;
        _1026_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1023_v7)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_1023_v7).length))],new BigNumber((_1020_v4).length));
        let _1027_v11;
        _1027_v11 = _dafny.Seq.of(_1020_v4, _1020_v4);
        _1026_v10 = (_1026_v10).update(new BigNumber(-212), new BigNumber((_1027_v11).length));
      } else {
        let _1028_v12;
        _1028_v12 = _dafny.Seq.of((_this).f13);
        let _1029_v13;
        let _nw136 = Array((new BigNumber(3)).toNumber());
        _nw136[(_dafny.ZERO).toNumber()] = ((_this).f17) === ((_this).f18);
        _nw136[(_dafny.ONE).toNumber()] = (_this).f13;
        _nw136[(new BigNumber(2)).toNumber()] = (_1028_v12)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1028_v12).length))];
        _1029_v13 = _nw136;
        let _1030_v14;
        _1030_v14 = false;
        let _1031_v15;
        let _init22 = function (_1032_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        };
        let _nw137 = Array((new BigNumber(23)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw137.length); _i0_22++) {
          _nw137[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _1031_v15 = _nw137;
        let _1033_v16;
        _1033_v16 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1034_v17;
        _1034_v17 = _dafny.Seq.of(_1033_v16, _1033_v16);
        let _1035_v18;
        _1035_v18 = new BigNumber(677);
        let _index151 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_1031_v15).length));
        (_1031_v15)[_index151] = (_1034_v17)[_module.__default.safeIndex(_1035_v18, new BigNumber((_1034_v17).length))];
        let _1036_v19;
        _1036_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(709),!((!(_1030_v14)) || ((_this).f18)));
        let _index152 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_1031_v15).length));
        let _rhs123 = _dafny.Seq.Concat(_dafny.Seq.of(_1030_v14), _dafny.Seq.Concat(_dafny.Seq.of(_1030_v14), _module.__default.fm29(_1035_v18, globalState)));
        let _rhs124 = _1029_v13;
        let _rhs125 = (((_1036_v19).contains(_1035_v18)) ? ((_1036_v19).get(_1035_v18)) : ((_this).f17));
        let _rhs126 = _1033_v16;
        let _lhs107 = globalState;
        let _lhs108 = _1031_v15;
        let _lhs109 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_1031_v15).length));
        _lhs107.f2 = _rhs123;
        _1029_v13 = _rhs124;
        _1030_v14 = _rhs125;
        _lhs108[_lhs109] = _rhs126;
        let _1037_v20;
        _1037_v20 = _dafny.Set.fromElements((_this).f13);
        (globalState).f7 = (_module.__default.safeModuloInt((_this).fm9(_1033_v16, globalState), new BigNumber(757))).multipliedBy(new BigNumber((_1037_v20).length));
        let _1038_v21;
        _1038_v21 = _dafny.MultiSet.fromElements(_1035_v18);
        let _1039_v22;
        _1039_v22 = _dafny.Seq.of(_1035_v18);
        let _1040_v23;
        let _nw138 = Array((new BigNumber(6)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_1041_v16) => function (_1042_i1) {
          return _1041_v16;
        })(_1033_v16))).length);
        _nw138[(_dafny.ONE).toNumber()] = new BigNumber((_1034_v17).length);
        _nw138[(new BigNumber(2)).toNumber()] = _1035_v18;
        _nw138[(new BigNumber(3)).toNumber()] = _1035_v18;
        _nw138[(new BigNumber(4)).toNumber()] = new BigNumber((_1038_v21).cardinality());
        _nw138[(new BigNumber(5)).toNumber()] = (_1039_v22)[_module.__default.safeIndex(new BigNumber((_1034_v17).length), new BigNumber((_1039_v22).length))];
        _1040_v23 = _nw138;
        let _1043_v24;
        _1043_v24 = _module.D2.create_DC4(_1040_v23, _module.__default.fm17(globalState), (_1035_v18).plus(new BigNumber((_1038_v21).cardinality())), (_dafny.Map.Empty.slice().updateUnsafe(_1035_v18,_1035_v18)).contains(_1035_v18), _1029_v13);
        let _source17 = _1043_v24;
        if (_source17.is_DC4) {
          let _1044___mcc_h0 = (_source17).cf6;
          let _1045___mcc_h1 = (_source17).cf7;
          let _1046___mcc_h2 = (_source17).cf8;
          let _1047___mcc_h3 = (_source17).cf9;
          let _1048___mcc_h4 = (_source17).cf10;
          let _1049_cf10 = _1048___mcc_h4;
          let _1050_cf9 = _1047___mcc_h3;
          let _1051_cf8 = _1046___mcc_h2;
          let _1052_cf7 = _1045___mcc_h1;
          let _1053_cf6 = _1044___mcc_h0;
          (globalState).f7 = _1035_v18;
          let _1054_v25;
          let _nw139 = Array((new BigNumber(24)).toNumber()).fill([]);
          _1054_v25 = _nw139;
          let _index153 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1054_v25).length));
          (_1054_v25)[_index153] = _1029_v13;
          let _index154 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1054_v25).length));
          (_1054_v25)[_index154] = _1049_cf10;
          let _1055_v26;
          _1055_v26 = _dafny.Seq.of(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), ((_1056_v15) => function (_1057_i2) {
            return (_1056_v15)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_1056_v15).length))];
          })(_1031_v15))));
          let _1058_v27;
          _1058_v27 = _module.D3.create_DC7();
          let _1059_v28;
          _1059_v28 = _dafny.Set.fromElements(_1058_v27);
          let _rhs127 = _1035_v18;
          let _rhs128 = ((_1035_v18).plus(new BigNumber((_1055_v26).length))).plus(_1035_v18);
          let _rhs129 = new BigNumber((_1059_v28).length);
          let _lhs110 = globalState;
          _1051_cf8 = _rhs127;
          _1051_cf8 = _rhs128;
          _lhs110.f7 = _rhs129;
          (globalState).f7 = _1035_v18;
        } else {
          let _1060___mcc_h5 = (_source17).cf5;
          let _1061_cf5 = _1060___mcc_h5;
          (_this).m11(_1029_v13, _1030_v14, (_1035_v18).isLessThanOrEqualTo(new BigNumber((_1028_v12).length)), globalState);
          _1061_cf5 = (_this).f14;
          (globalState).f7 = (_1035_v18).plus(_module.__default.safeModuloInt((_dafny.ZERO).minus(_1035_v18), _1035_v18));
          _1061_cf5 = (_this).f14;
        }
        let _1062_v29;
        let _nw140 = Array((new BigNumber(14)).toNumber());
        _1062_v29 = _nw140;
        let _index155 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_1040_v23).length));
        (_1040_v23)[_index155] = _1035_v18;
        let _index156 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_1040_v23).length));
        let _rhs130 = _1062_v29;
        let _rhs131 = (_this).f14;
        let _rhs132 = _1035_v18;
        let _rhs133 = (_1039_v22)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_1035_v18, _1035_v18), new BigNumber((_1039_v22).length))];
        let _lhs111 = _1040_v23;
        let _lhs112 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_1040_v23).length));
        let _lhs113 = globalState;
        _1062_v29 = _rhs130;
        _1029_v13 = _rhs131;
        _lhs111[_lhs112] = _rhs132;
        _lhs113.f7 = _rhs133;
        _1030_v14 = (_this).f18;
      }
      let _1063_v30;
      _1063_v30 = false;
      _1063_v30 = true;
      let _1064_v31;
      _1064_v31 = _module.D6.create_DC11(_module.__default.fm19(globalState));
      _1064_v31 = _1064_v31;
      let _source18 = _this.f12;
      if (_source18.is_DC2) {
        let _1065___mcc_h6 = (_source18).cf2;
        let _1066___mcc_h7 = (_source18).cf3;
        let _1067___mcc_h8 = (_source18).cf4;
        let _1068_cf4 = _1067___mcc_h8;
        let _1069_cf3 = _1066___mcc_h7;
        let _1070_cf2 = _1065___mcc_h6;
        let _1071_v32;
        _1071_v32 = new _dafny.CodePoint('f'.codePointAt(0));
        _1071_v32 = _1071_v32;
        let _1072_v33;
        _1072_v33 = _dafny.Seq.of(_1069_cf3);
        _1063_v30 = !_dafny.areEqual(_dafny.Seq.Concat(_1072_v33, _1072_v33), _1072_v33);
        let _1073_v34;
        let _nw141 = new _module.C2();
        _nw141.__ctor(new BigNumber(-855), _1068_cf4, _this.f16, _1068_cf4, (_this).f14, (_this).f15, _this.f12, _1063_v30);
        _1073_v34 = _nw141;
        let _1074_v35;
        let _init23 = ((_1075_cf2) => function (_1076_i3) {
          return _module.__default.safeDivisionInt(_1076_i3, _1075_cf2);
        })(_1070_cf2);
        let _nw142 = Array((new BigNumber(21)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw142.length); _i0_23++) {
          _nw142[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1074_v35 = _nw142;
        let _1077_v36;
        _1077_v36 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f17) ? (_1073_v34) : (_1073_v34)),_1074_v35);
        (globalState).f7 = (_dafny.ZERO).minus(new BigNumber((_1077_v36).length));
        let _1078_v37;
        _1078_v37 = _dafny.Seq.of(((_1068_cf4) ? ((_this).f13) : (_1068_cf4)));
        _1063_v30 = !(!((_1078_v37)[_module.__default.safeIndex(_1070_cf2, new BigNumber((_1078_v37).length))]));
      } else {
        let _1079___mcc_h9 = (_source18).cf1;
        let _1080_cf1 = _1079___mcc_h9;
        _1063_v30 = ((!(false)) ? (_1063_v30) : (_module.__default.fm30(new BigNumber(38), globalState)));
        let _1081_v38;
        _1081_v38 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1082_v39;
        let _nw143 = new _module.C0();
        _nw143.__ctor((_this).f17, _1081_v38, _this.f12, (((_this).f13) ? ((_this).f13) : (_1063_v30)));
        _1082_v39 = _nw143;
        let _1083_v40;
        _1083_v40 = new BigNumber(872);
        let _1084_v41;
        _1084_v41 = _dafny.Set.fromElements((_this).f13, !(!((_this).f13)));
        let _1085_v42;
        let _nw144 = new _module.C2();
        _nw144.__ctor((_dafny.ZERO).minus(_1083_v40), (_dafny.Set.fromElements((_this).f17, !(false))).IsSubsetOf(_1084_v41), _module.D1.create_DC1(_1080_cf1), _dafny.Seq.IsPrefixOf(_1080_cf1, _1080_cf1), (_this).f14, (_this).f15, _this.f12, _1063_v30);
        _1085_v42 = _nw144;
        let _1086_v43;
        _1086_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1085_v42).f21,_1063_v30);
        let _1087_v44;
        _1087_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1086_v43,(_1083_v40).plus(_1083_v40));
        let _1088_v45;
        _1088_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f13);
        _1087_v44 = (_1087_v44).update(_1086_v43, (new BigNumber((_1088_v45).length)).minus(_1083_v40));
      }
      let _1089_v46;
      _1089_v46 = new BigNumber(-61);
      let _1090_v47;
      _1090_v47 = _dafny.Seq.UnicodeFromString("hu");
      let _1091_v48;
      _1091_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1089_v46,_1090_v47);
      let _hi5 = _module.__default.safeModuloInt(new BigNumber((_1091_v48).length), _1089_v46);
      for (let _1092_i4 = (_dafny.ZERO).minus(_1089_v46); _1092_i4.isLessThan(_hi5); _1092_i4 = _1092_i4.plus(_dafny.ONE)) {
        let _1093_v49;
        _1093_v49 = _dafny.MultiSet.fromElements((_this).f14);
        let _1094_v50;
        _1094_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1093_v49,(_this).f14);
        _1094_v50 = (_1094_v50).update(_1093_v49, (_this).f14);
        let _1095_v51;
        let _init24 = function (_1096_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f17,new BigNumber(323));
        };
        let _nw145 = Array((new BigNumber(3)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw145.length); _i0_24++) {
          _nw145[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1095_v51 = _nw145;
        let _1097_v53;
        _1097_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1095_v51,_dafny.Map.Empty.slice().updateUnsafe((function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f15).length),(_this).f17)).Keys.Elements) {
            let _1098_v52 = _compr_42;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f15).length),(_this).f17)).contains(_1098_v52)) {
              _coll42.push([(_1098_v52).multipliedBy(new BigNumber(-276)),false]);
            }
          }
          return _coll42;
        }()).update(_1089_v46, (_this).f13),_1090_v47));
        let _1099_v54;
        _1099_v54 = _dafny.MultiSet.fromElements(new BigNumber(-748));
        let _1100_v55;
        _1100_v55 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1099_v54).cardinality())),_1063_v30);
        let _1101_v57;
        _1101_v57 = _dafny.Seq.of(_1063_v30);
        let _1102_v58;
        _1102_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1100_v55,_1101_v57);
        _1097_v53 = (_1097_v53).update(_1095_v51, (_dafny.Map.Empty.slice().updateUnsafe(_1100_v55,_1090_v47)).Merge((function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of (_1102_v58).Keys.Elements) {
            let _1103_v56 = _compr_43;
            if ((_1102_v58).contains(_1103_v56)) {
              _coll43.push([_1103_v56,_1090_v47]);
            }
          }
          return _coll43;
        }()).update(_1100_v55, _1090_v47)));
        let _1104_v59;
        _1104_v59 = _dafny.Set.fromElements(_1063_v30, (_this).f13);
        _1063_v30 = ((_dafny.ZERO).minus(_1089_v46)).isLessThanOrEqualTo(new BigNumber((_1104_v59).length));
        _1063_v30 = (_this).f17;
      }
      let _1105_v60;
      _1105_v60 = _module.D6.create_DC12(false, _1090_v47, new BigNumber(629), _1063_v30);
      let _index157 = _module.__default.safeIndex(new BigNumber(901), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index157] = (_1105_v60).dtor_cf21;
      let _index158 = _module.__default.safeIndex(new BigNumber(901), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index158] = (_dafny.Set.fromElements(_1089_v46)).IsDisjointFrom((_module.D9.create_DC20((_this).f15)).dtor_cf36);
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _source19 = _this.f16;
      if (_source19.is_DC2) {
        let _1106___mcc_h0 = (_source19).cf2;
        let _1107___mcc_h1 = (_source19).cf3;
        let _1108___mcc_h2 = (_source19).cf4;
        let _1109_cf4 = _1108___mcc_h2;
        let _1110_cf3 = _1107___mcc_h1;
        let _1111_cf2 = _1106___mcc_h0;
        let _1112_v0;
        _1112_v0 = _dafny.Seq.UnicodeFromString("havb");
        let _1113_v1;
        _1113_v1 = _dafny.MultiSet.fromElements(_1112_v0, _1112_v0);
        _1109_cf4 = !((_1113_v1).IsProperSubsetOf(_1113_v1));
        let _1114_v2;
        let _nw146 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1114_v2 = _nw146;
        _1114_v2 = _1114_v2;
        let _1115_v3;
        _1115_v3 = _dafny.Set.fromElements(_1114_v2);
        _1111_cf2 = new BigNumber((_1115_v3).length);
        let _1116_v4;
        _1116_v4 = _dafny.MultiSet.fromElements((_1111_cf2).isEqualTo((_this).fm8((_this).f17, _this.f16, globalState)), (_this).f17);
        let _1117_v5;
        _1117_v5 = _dafny.Seq.of(_1109_cf4, _1109_cf4);
        let _1118_v6;
        _1118_v6 = _dafny.Seq.of(_1117_v5, _1117_v5);
        let _rhs134 = ((true) ? (_1116_v4) : ((_dafny.MultiSet.FromArray((_1118_v6)[_module.__default.safeIndex(_1110_cf3, new BigNumber((_1118_v6).length))])).Difference(_1116_v4)));
        let _rhs135 = (_this).f13;
        let _rhs136 = true;
        _1116_v4 = _rhs134;
        _1109_cf4 = _rhs135;
        _1109_cf4 = _rhs136;
      } else {
        let _1119___mcc_h3 = (_source19).cf1;
        let _1120_cf1 = _1119___mcc_h3;
        let _1121_v7;
        let _init25 = ((_1122_p0) => function (_1123_i0) {
          return (_1123_i0).plus(_1122_p0);
        })(p0);
        let _nw147 = Array((new BigNumber(4)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw147.length); _i0_25++) {
          _nw147[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1121_v7 = _nw147;
        let _index159 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1121_v7).length));
        (_1121_v7)[_index159] = p0;
        let _index160 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1121_v7).length));
        (_1121_v7)[_index160] = p0;
        let _1124_v8;
        _1124_v8 = _dafny.Set.fromElements((_this).f17, _dafny.areEqual(_1120_cf1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), function (_1125_i1) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })), (_this).f13, (_this).f18);
        _1124_v8 = _1124_v8;
        _1120_cf1 = _1120_cf1;
        let _1126_v9;
        _1126_v9 = new _dafny.CodePoint('h'.codePointAt(0));
        _1126_v9 = _1126_v9;
      }
      let _1127_v10;
      let _nw148 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1127_v10 = _nw148;
      _1127_v10 = (((p0).isLessThan(p0)) ? (_1127_v10) : (_1127_v10));
      let _1128_v11;
      let _init26 = ((_1129_p0) => function (_1130_i3) {
        return (_1130_i3).multipliedBy(_1129_p0);
      })(p0);
      let _nw149 = Array((new BigNumber(23)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw149.length); _i0_26++) {
        _nw149[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1128_v11 = _nw149;
      let _1131_v12;
      _1131_v12 = _module.D2.create_DC4(_1128_v11, (_this).f17, (_dafny.ZERO).minus(p0), true, (_this).f14);
      let _1132_v13;
      _1132_v13 = _dafny.MultiSet.fromElements((_this).f17);
      let _1133_v14;
      _1133_v14 = _1132_v13;
      let _1134_v15;
      _1134_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_this).f18));
      let _1135_v16;
      _1135_v16 = _module.D3.create_DC6((_dafny.ZERO).minus(p0), (_this.f12).dtor_cf4);
      let _1136_v17;
      let _nw150 = Array((new BigNumber(14)).toNumber());
      _nw150[(_dafny.ZERO).toNumber()] = _1131_v12;
      _nw150[(_dafny.ONE).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(2)).toNumber()] = _module.D2.create_DC4(_1128_v11, (_this).f17, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1133_v14)).length)), (((_1134_v15).contains(p0)) ? ((_1134_v15).get(p0)) : (false)), (_this).f14);
      _nw150[(new BigNumber(3)).toNumber()] = _module.D2.create_DC4(_1128_v11, (_this).f18, p0, (_this).f17, (_this).f14);
      _nw150[(new BigNumber(4)).toNumber()] = _module.D2.create_DC4(_1128_v11, (_this).f18, p0, (_1135_v16).dtor_cf13, (_this).f14);
      _nw150[(new BigNumber(5)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(6)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(7)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(8)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(9)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(10)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(11)).toNumber()] = _1131_v12;
      _nw150[(new BigNumber(12)).toNumber()] = _module.D2.create_DC4(_1128_v11, (_this).f13, p0, (_this).f13, (_1131_v12).dtor_cf10);
      _nw150[(new BigNumber(13)).toNumber()] = _1131_v12;
      _1136_v17 = _nw150;
      let _1137_v18;
      _1137_v18 = _dafny.Set.fromElements(_1136_v17);
      let _1138_i2;
      _1138_i2 = _dafny.ZERO;
      L4: {
        while ((_dafny.Set.fromElements(_1136_v17)).IsSubsetOf(_1137_v18)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1138_i2)) {
              break L4;
            }
            _1138_i2 = (_1138_i2).plus(_dafny.ONE);
            _1134_v15 = (_1134_v15).update(p0, (_this).f17);
            let _1139_v19;
            _1139_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _1140_v20;
            _1140_v20 = _dafny.Seq.of((_this).f17);
            _1139_v19 = (_1139_v19).update(_module.__default.fm0(_1140_v20, (_this).f18, (_this).f13, p0, globalState), new BigNumber(114));
            r0 = (_this).f18;
            let _1141_v21;
            _1141_v21 = _dafny.MultiSet.fromElements(p0, p0);
            r0 = ((!(_1141_v21).equals(_1141_v21)) ? ((_this).f18) : (((_this).f15).IsSubsetOf((_this).f15)));
          }
        }
      }
      let _1142_v22;
      _1142_v22 = _dafny.Seq.of(!(false));
      let _1143_v23;
      _1143_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-575),p0);
      let _1144_v24;
      _1144_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC23(p0, p0, (_this).fm15(globalState), _module.__default.fm0(_1142_v22, (_this).f17, false, new BigNumber((_1143_v23).length), globalState), (_dafny.ZERO).minus(p0)),p0);
      let _1145_v25;
      _1145_v25 = _dafny.Seq.of(new BigNumber((_1144_v24).length), p0);
      let _1146_v26;
      _1146_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1145_v25,p0);
      let _index161 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_1128_v11).length));
      (_1128_v11)[_index161] = new BigNumber((_1146_v26).length);
      let _index162 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_1128_v11).length));
      (_1128_v11)[_index162] = p0;
      let _1147_v27;
      _1147_v27 = _dafny.Seq.UnicodeFromString("itq");
      let _1148_v28;
      _1148_v28 = _module.D6.create_DC12((_this).f13, _1147_v27, p0, (_this).f18);
      let _1149_v29;
      let _nw151 = new _module.C2();
      _nw151.__ctor(_module.__default.safeModuloInt(p0, (_1128_v11)[_module.__default.safeIndex(new BigNumber(895), new BigNumber((_1128_v11).length))]), (_this).f13, _this.f16, (((_1148_v28).dtor_cf21) ? ((_this).f13) : ((_this).f13)), (_this).f14, ((_this).f15).Union((_this).f15), _this.f12, true);
      _1149_v29 = _nw151;
      _1147_v27 = _dafny.Seq.Concat(_1147_v27, _1147_v27);
      r0 = (_this).f17;
      return r0;
    }
    m11(p0, p1, p2, globalState) {
      let _this = this;
      let _1150_v0;
      _1150_v0 = new BigNumber(685);
      if (_module.__default.fm30(_1150_v0, globalState)) {
        let _1151_v1;
        _1151_v1 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(_1150_v0, globalState),p1);
        let _1152_v2;
        _1152_v2 = _dafny.Seq.of(true, false, p2, (_this).f18, true);
        _1151_v1 = (_1151_v1).update((_1152_v2)[_module.__default.safeIndex(_1150_v0, new BigNumber((_1152_v2).length))], (_this).f18);
        let _1153_v3;
        _1153_v3 = _dafny.Seq.of(new BigNumber(-44));
        let _1154_v4;
        _1154_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1150_v0);
        let _1155_v5;
        _1155_v5 = _dafny.Set.fromElements(p2);
        let _1156_v6;
        _1156_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1155_v5,_1150_v0);
        let _1157_v7;
        _1157_v7 = _dafny.MultiSet.fromElements(false);
        let _1158_v8;
        _1158_v8 = _dafny.Seq.UnicodeFromString("jlbyvmu");
        let _1159_v9;
        let _init27 = ((_1160_v0) => function (_1161_i0) {
          return _module.__default.safeDivisionInt(_1161_i0, _1160_v0);
        })(_1150_v0);
        let _nw152 = Array((new BigNumber(9)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw152.length); _i0_27++) {
          _nw152[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1159_v9 = _nw152;
        let _1162_v10;
        _1162_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v0,_1150_v0);
        let _1163_v11;
        _1163_v11 = _dafny.Seq.of(_1154_v4);
        let _1164_v12;
        _1164_v12 = _module.D9.create_DC23(_1150_v0, new BigNumber((_dafny.Seq.UnicodeFromString("f")).length), _1150_v0, new BigNumber((_1163_v11).length), _1150_v0);
        let _1165_v13;
        let _nw153 = Array((new BigNumber(24)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = ((p2) ? ((_1153_v3)[_module.__default.safeIndex((((_1154_v4).contains((_this).f17)) ? ((_1154_v4).get((_this).f17)) : (_1150_v0)), new BigNumber((_1153_v3).length))]) : (new BigNumber(272)));
        _nw153[(_dafny.ONE).toNumber()] = new BigNumber(((_module.__default.fm44((_this).f17, globalState)).Merge(_1156_v6)).length);
        _nw153[(new BigNumber(2)).toNumber()] = _module.__default.fm0(_1152_v2, p1, false, _1150_v0, globalState);
        _nw153[(new BigNumber(3)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(4)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(5)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(6)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(7)).toNumber()] = (_1150_v0).multipliedBy(_1150_v0);
        _nw153[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_1150_v0, _1150_v0);
        _nw153[(new BigNumber(9)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(305)).plus(new BigNumber(((_this).f15).length)));
        _nw153[(new BigNumber(11)).toNumber()] = new BigNumber(770);
        _nw153[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1157_v7).cardinality()), (_dafny.ZERO).minus(_1150_v0));
        _nw153[(new BigNumber(13)).toNumber()] = _module.__default.fm0(_1152_v2, (_this).f13, false, _module.__default.fm0(_1152_v2, (_this).f13, _module.__default.fm30(new BigNumber((_1158_v8).length), globalState), _1150_v0, globalState), globalState);
        _nw153[(new BigNumber(14)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(15)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_module.D2.create_DC4(_1159_v9, !((_this).f17), _1150_v0, (_this).f18, (_this).f14)).dtor_cf8)).length);
        _nw153[(new BigNumber(17)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(18)).toNumber()] = (_1150_v0).minus(new BigNumber(998));
        _nw153[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_1150_v0, _1150_v0, new BigNumber((_1162_v10).length), new BigNumber(-793), new BigNumber(((_this).f15).length))).cardinality());
        _nw153[(new BigNumber(20)).toNumber()] = (_1150_v0).multipliedBy(_1150_v0);
        _nw153[(new BigNumber(21)).toNumber()] = (((_1162_v10).contains((_1164_v12).dtor_cf45)) ? ((_1162_v10).get((_1164_v12).dtor_cf45)) : (new BigNumber((_1158_v8).length)));
        _nw153[(new BigNumber(22)).toNumber()] = _1150_v0;
        _nw153[(new BigNumber(23)).toNumber()] = new BigNumber(630);
        _1165_v13 = _nw153;
        let _index163 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length));
        (_1165_v13)[_index163] = _1150_v0;
        let _index164 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length));
        (_1165_v13)[_index164] = ((_this).fm15(globalState)).multipliedBy(_1150_v0);
        let _1166_v14;
        let _nw154 = new _module.C1();
        _nw154.__ctor(new _dafny.CodePoint('j'.codePointAt(0)), _this.f12, (_this).f17);
        _1166_v14 = _nw154;
        let _1167_v15;
        let _nw155 = Array((new BigNumber(18)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _1166_v14;
        _nw155[(_dafny.ONE).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(2)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(3)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(4)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(5)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(6)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(7)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(8)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(9)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(10)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(11)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(12)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(13)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(14)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(15)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(16)).toNumber()] = _1166_v14;
        _nw155[(new BigNumber(17)).toNumber()] = _1166_v14;
        _1167_v15 = _nw155;
        let _1168_v16;
        _1168_v16 = _dafny.MultiSet.fromElements(_1167_v15, _1167_v15);
        let _1169_v17;
        _1169_v17 = false;
        let _rhs137 = (_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))];
        let _rhs138 = _1168_v16;
        let _rhs139 = ((_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))]).isLessThan((_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))]);
        let _rhs140 = (new BigNumber((_1157_v7).cardinality())).isEqualTo((_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))]);
        let _rhs141 = p1;
        let _lhs114 = globalState;
        _lhs114.f7 = _rhs137;
        _1168_v16 = _rhs138;
        _1169_v17 = _rhs139;
        _1169_v17 = _rhs140;
        _1169_v17 = _rhs141;
        if ((_this).f13) {
          let _index165 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length));
          (_1165_v13)[_index165] = new BigNumber(-903);
          (globalState).f7 = (_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))];
          _1169_v17 = p2;
          let _1170_v18;
          _1170_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f15).length),p2);
          let _rhs142 = (((((_1170_v18).contains(new BigNumber((_1157_v7).cardinality()))) ? ((_1170_v18).get(new BigNumber((_1157_v7).cardinality()))) : ((_this).f13))) ? (((p2) ? (_1166_v14.f23) : (_1166_v14.f23))) : (_1166_v14.f23));
          let _rhs143 = _1166_v14.f23;
          let _lhs115 = _1166_v14;
          let _lhs116 = _1166_v14;
          _lhs115.f23 = _rhs142;
          _lhs116.f23 = _rhs143;
          let _1171_v19;
          let _nw156 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1171_v19 = _nw156;
          _1171_v19 = p0;
        } else {
          let _1172_v20;
          _1172_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))],p0);
          _1172_v20 = (_1172_v20).update(new BigNumber((_1152_v2).length), (_this).f14);
          let _1173_v21;
          _1173_v21 = _dafny.Set.fromElements(_1158_v8);
          let _1174_v22;
          _1174_v22 = _1173_v21;
          let _pat_let_tv31 = _1165_v13;
          let _pat_let_tv32 = _1165_v13;
          _1174_v22 = (((function (_pat_let31_0) {
            return function (_1175_dt__update__tmp_h0) {
              return function (_pat_let32_0) {
                return function (_1176_dt__update_hcf2_h0) {
                  return _module.D1.create_DC2(_1176_dt__update_hcf2_h0, (_1175_dt__update__tmp_h0).dtor_cf3, (_1175_dt__update__tmp_h0).dtor_cf4);
                }(_pat_let32_0);
              }((_pat_let_tv32)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_pat_let_tv31).length))]);
            }(_pat_let31_0);
          }(_this.f12)).dtor_cf4) ? (_1173_v21) : (_1174_v22));
          let _1177_v23;
          _1177_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1166_v14,_1150_v0);
          _1177_v23 = (_1177_v23).update(_1166_v14, (_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))]);
          let _1178_v24;
          let _nw157 = Array((new BigNumber(20)).toNumber());
          _nw157[(_dafny.ZERO).toNumber()] = _1159_v9;
          _nw157[(_dafny.ONE).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(2)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(3)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(4)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(5)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(6)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(7)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(8)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(9)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(10)).toNumber()] = (_module.D2.create_DC4(_1159_v9, (_this).f18, (_this).fm8((_this).f13, _this.f16, globalState), !((_this).f18), (_this).f14)).dtor_cf6;
          _nw157[(new BigNumber(11)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(12)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(13)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(14)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(15)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(16)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(17)).toNumber()] = _1165_v13;
          _nw157[(new BigNumber(18)).toNumber()] = _1159_v9;
          _nw157[(new BigNumber(19)).toNumber()] = _1165_v13;
          _1178_v24 = _nw157;
          let _index166 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1178_v24).length));
          (_1178_v24)[_index166] = _1159_v9;
          let _1179_v25;
          _1179_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v0,_1165_v13);
          let _index167 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1178_v24).length));
          (_1178_v24)[_index167] = (((_1179_v25).contains(_1150_v0)) ? ((_1179_v25).get(_1150_v0)) : ((_module.D2.create_DC4(_1165_v13, (_this).f13, new BigNumber(612), (_this).f13, (_this).f14)).dtor_cf6));
          let _1180_v26;
          _1180_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_dafny.Set.fromElements((_1165_v13)[_module.__default.safeIndex(new BigNumber(465), new BigNumber((_1165_v13).length))]));
          let _1181_v27;
          _1181_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1166_v14).fm43(globalState),(((_1180_v26).contains((_this).f17)) ? ((_1180_v26).get((_this).f17)) : ((_this).f15)));
          _1180_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f15);
        }
        (_this).m12(globalState);
      } else {
        let _1182_v28;
        _1182_v28 = _dafny.Seq.UnicodeFromString("lxk");
        let _1183_v29;
        _1183_v29 = _dafny.Seq.of(_1150_v0);
        let _1184_v30;
        _1184_v30 = _module.D3.create_DC5(new _dafny.CodePoint('c'.codePointAt(0)));
        let _1185_v31;
        _1185_v31 = _dafny.MultiSet.fromElements((_1184_v30).dtor_cf11);
        let _1186_v32;
        _1186_v32 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1187_v33;
        _1187_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(997),_module.D7.create_DC17(_module.D7.create_DC15(_1183_v29)));
        let _1188_v34;
        _1188_v34 = _dafny.MultiSet.fromElements(_1150_v0, _1150_v0);
        let _1189_v35;
        _1189_v35 = _module.D7.create_DC16(_1150_v0);
        let _1190_v36;
        _1190_v36 = _module.D7.create_DC17(_1189_v35);
        let _1191_v37;
        let _nw158 = Array((new BigNumber(20)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = _1150_v0;
        _nw158[(_dafny.ONE).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(642), _1150_v0);
        _nw158[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_1150_v0, _1150_v0);
        _nw158[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_1150_v0);
        _nw158[(new BigNumber(5)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(6)).toNumber()] = new BigNumber((_1182_v28).length);
        _nw158[(new BigNumber(7)).toNumber()] = (new BigNumber((_1183_v29).length)).minus((_dafny.ZERO).minus(_1150_v0));
        _nw158[(new BigNumber(8)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(9)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(10)).toNumber()] = ((p1) ? ((((_1185_v31).contains(_1186_v32)) ? ((_1185_v31).get(_1186_v32)) : (_1150_v0))) : ((_this).fm9(_1186_v32, globalState)));
        _nw158[(new BigNumber(11)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(12)).toNumber()] = new BigNumber(((_this).f15).length);
        _nw158[(new BigNumber(13)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(14)).toNumber()] = (_1150_v0).plus(_1150_v0);
        _nw158[(new BigNumber(15)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(16)).toNumber()] = new BigNumber(-151);
        _nw158[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.update(_1182_v28, _module.__default.safeIndex(new BigNumber(((_1187_v33).update(new BigNumber((_1188_v34).cardinality()), _1190_v36)).length), new BigNumber((_1182_v28).length)), _1186_v32)).length);
        _nw158[(new BigNumber(18)).toNumber()] = _1150_v0;
        _nw158[(new BigNumber(19)).toNumber()] = (_1150_v0).plus(new BigNumber((_1182_v28).length));
        _1191_v37 = _nw158;
        let _index168 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length));
        (_1191_v37)[_index168] = _1150_v0;
        let _index169 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length));
        (_1191_v37)[_index169] = _1150_v0;
        _1150_v0 = new BigNumber((_module.__default.fm20(globalState)).length);
        let _1192_v38;
        let _nw159 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1192_v38 = _nw159;
        let _index170 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1192_v38).length));
        (_1192_v38)[_index170] = _1186_v32;
        let _index171 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1192_v38).length));
        (_1192_v38)[_index171] = _1186_v32;
        let _1193_v39;
        let _nw160 = Array((new BigNumber(11)).toNumber()).fill([]);
        _1193_v39 = _nw160;
        let _1194_v40;
        _1194_v40 = _1193_v39;
        let _source20 = _1194_v40;
        let _1195___mcc_h0 = _source20;
        let _1196_cf15 = _1195___mcc_h0;
        let _1197_v41;
        _1197_v41 = _dafny.Set.fromElements((_this).f18);
        _1150_v0 = new BigNumber(((_1188_v34).Union((_dafny.MultiSet.fromElements(new BigNumber((_1197_v41).length))).Intersect(_dafny.MultiSet.fromElements((_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))])))).cardinality());
        let _index172 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((p0).length));
        (p0)[_index172] = !(p1);
        let _index173 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((p0).length));
        (p0)[_index173] = p1;
        let _1198_v42;
        let _init28 = function (_1199_i1) {
          return false;
        };
        let _nw161 = Array((new BigNumber(11)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw161.length); _i0_28++) {
          _nw161[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1198_v42 = _nw161;
        let _1200_v43;
        let _nw162 = Array((_dafny.ONE).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _1191_v37;
        _1200_v43 = _nw162;
        let _index174 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1200_v43).length));
        (_1200_v43)[_index174] = _1191_v37;
        let _1201_v44;
        _1201_v44 = _dafny.MultiSet.fromElements((_this).f14);
        let _1202_v45;
        _1202_v45 = _dafny.MultiSet.fromElements(_1198_v42);
        let _1203_v46;
        _1203_v46 = _dafny.MultiSet.fromElements(_1183_v29);
        let _1204_v47;
        let _nw163 = Array((new BigNumber(5)).toNumber()).fill(_module.D3.Default());
        _1204_v47 = _nw163;
        let _1205_v48;
        _1205_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1204_v47,p1);
        let _1206_v49;
        _1206_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v48,_1198_v42);
        let _1207_v50;
        _1207_v50 = _dafny.Seq.of(_1191_v37, ((p1) ? (_1191_v37) : (_1191_v37)));
        let _index175 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((p0).length));
        let _index176 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1200_v43).length));
        let _rhs144 = _module.__default.safeModuloInt((_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))], (_dafny.ZERO).minus(new BigNumber(((_1202_v45)).cardinality())));
        let _rhs145 = (_1203_v46).IsDisjointFrom(_1203_v46);
        let _rhs146 = (((_1206_v49).contains(_1205_v48)) ? ((_1206_v49).get(_1205_v48)) : ((_this).f14));
        let _rhs147 = (_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))];
        let _rhs148 = (_1207_v50)[_module.__default.safeIndex(_1150_v0, new BigNumber((_1207_v50).length))];
        let _lhs117 = p0;
        let _lhs118 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((p0).length));
        let _lhs119 = globalState;
        let _lhs120 = _1200_v43;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1200_v43).length));
        _1150_v0 = _rhs144;
        _lhs117[_lhs118] = _rhs145;
        _1198_v42 = _rhs146;
        _lhs119.f7 = _rhs147;
        _lhs120[_lhs121] = _rhs148;
        let _1208_v51;
        _1208_v51 = _dafny.Seq.of(p1, (_this).f13);
        let _1209_v52;
        _1209_v52 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f18) ? (_1198_v42) : (_1198_v42)),(_dafny.ZERO).minus(new BigNumber((_1208_v51).length)));
        _1209_v52 = (_1209_v52).update(_1198_v42, _module.__default.safeModuloInt(_1150_v0, (_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))]));
        if (p1) {
          let _1210_v53;
          _1210_v53 = false;
          _1210_v53 = p2;
          _1182_v28 = _module.__default.fm20(globalState);
          _1186_v32 = _1186_v32;
          let _index177 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length));
          (_1191_v37)[_index177] = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber(41), _1150_v0)).multipliedBy((_dafny.ZERO).minus(_1150_v0)));
          _1210_v53 = ((((_this).f18) ? (_1188_v34) : (_dafny.MultiSet.fromElements(new BigNumber(670))))).IsDisjointFrom(_module.__default.fm19(globalState));
        } else {
          let _1211_v54;
          _1211_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC11(_1188_v34),false);
          let _1212_v55;
          let _nw164 = new _module.C2();
          _nw164.__ctor(_1150_v0, (_this).f18, _this.f16, (((_1211_v54).contains(_module.D6.create_DC11(_1188_v34))) ? ((_1211_v54).get(_module.D6.create_DC11(_1188_v34))) : (p1)), p0, ((_this).f15).Difference((_this).f15), _this.f12, ((_dafny.ZERO).minus(_1150_v0)).isEqualTo((_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))]));
          _1212_v55 = _nw164;
          let _1213_v56;
          _1213_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1191_v37,(_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))]);
          _1213_v56 = (_1213_v56).update(_1191_v37, _module.__default.safeDivisionInt((_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))], _1150_v0));
          let _1214_v57;
          _1214_v57 = _dafny.Map.Empty.slice().updateUnsafe((_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))],(_1212_v55).f22);
          _1214_v57 = _1214_v57;
          let _1215_v58;
          _1215_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,(_1212_v55).f21);
          _1215_v58 = (_1215_v58).update((_this).f13, (_1212_v55).f21);
          (globalState).f7 = _module.__default.safeDivisionInt(_1150_v0, (_1191_v37)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_1191_v37).length))]);
        }
      }
      let _1216_v59;
      _1216_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v0,new _dafny.CodePoint('h'.codePointAt(0)));
      let _1217_v60;
      _1217_v60 = _dafny.Seq.of(_1150_v0, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f17, (_this).f13, (_this).f13, (_this).f18))).cardinality()));
      let _1218_v61;
      let _nw165 = Array((new BigNumber(9)).toNumber());
      _nw165[(_dafny.ZERO).toNumber()] = _1150_v0;
      _nw165[(_dafny.ONE).toNumber()] = new BigNumber(((_1216_v59).Merge(_1216_v59)).length);
      _nw165[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_1217_v60)[_module.__default.safeIndex(_1150_v0, new BigNumber((_1217_v60).length))]);
      _nw165[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_1150_v0)));
      _nw165[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_dafny.Seq.of((_this).f13), true, (_this).f17, new BigNumber(384), globalState);
      _nw165[(new BigNumber(5)).toNumber()] = _1150_v0;
      _nw165[(new BigNumber(6)).toNumber()] = _1150_v0;
      _nw165[(new BigNumber(7)).toNumber()] = _1150_v0;
      _nw165[(new BigNumber(8)).toNumber()] = _1150_v0;
      _1218_v61 = _nw165;
      let _index178 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
      (_1218_v61)[_index178] = (_1150_v0).plus(_1150_v0);
      let _1219_v62;
      _1219_v62 = _dafny.MultiSet.fromElements((_this).f14);
      let _1220_v63;
      _1220_v63 = true;
      let _1221_v64;
      _1221_v64 = _dafny.Seq.UnicodeFromString("paa");
      let _index179 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_1218_v61).length));
      (_1218_v61)[_index179] = new BigNumber((_1221_v64).length);
      let _index180 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
      let _index181 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_1218_v61).length));
      let _rhs149 = _1150_v0;
      let _rhs150 = _1219_v62;
      let _rhs151 = p1;
      let _rhs152 = false;
      let _rhs153 = _1150_v0;
      let _lhs122 = _1218_v61;
      let _lhs123 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
      let _lhs124 = _1218_v61;
      let _lhs125 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_1218_v61).length));
      _lhs122[_lhs123] = _rhs149;
      _1219_v62 = _rhs150;
      _1220_v63 = _rhs151;
      _1220_v63 = _rhs152;
      _lhs124[_lhs125] = _rhs153;
      let _1222_v65;
      _1222_v65 = _dafny.Seq.of(_1220_v63);
      let _1223_v66;
      _1223_v66 = _dafny.Map.Empty.slice().updateUnsafe((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))],(_1222_v65)[_module.__default.safeIndex(_1150_v0, new BigNumber((_1222_v65).length))]);
      if ((((_1223_v66).contains(_1150_v0)) ? ((_1223_v66).get(_1150_v0)) : (_1220_v63))) {
        (globalState).f7 = (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))];
        let _index182 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
        (_1218_v61)[_index182] = new BigNumber((_1217_v60).length);
        let _index183 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
        (_1218_v61)[_index183] = _module.__default.safeDivisionInt((_this).fm15(globalState), (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]);
        let _1224_v67;
        _1224_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(globalState),_1220_v63);
        _1224_v67 = (_1224_v67).update(_this.f16, (new BigNumber(810)).isLessThan(_1150_v0));
        (globalState).f7 = _module.__default.safeModuloInt((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], _1150_v0);
      } else {
        let _1225_v68;
        _1225_v68 = _dafny.Seq.of(_1221_v64);
        let _1226_v69;
        let _nw166 = Array((new BigNumber(3)).toNumber());
        _nw166[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1225_v68, _1225_v68);
        _nw166[(_dafny.ONE).toNumber()] = _1225_v68;
        _nw166[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_1221_v64);
        _1226_v69 = _nw166;
        let _index184 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1226_v69).length));
        (_1226_v69)[_index184] = _1225_v68;
        let _1227_v70;
        _1227_v70 = _module.D6.create_DC11(_dafny.MultiSet.fromElements((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]));
        let _1228_v71;
        _1228_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1227_v70,_1221_v64);
        let _1229_v72;
        _1229_v72 = _module.D6.create_DC12((_this).f13, (((_1228_v71).contains(_1227_v70)) ? ((_1228_v71).get(_1227_v70)) : (_1221_v64)), (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], _1220_v63);
        let _index185 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_1226_v69).length));
        (_1226_v69)[_index185] = _module.__default.fm45(_1229_v72, _1150_v0, globalState);
        let _1230_v73;
        _1230_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(332),new BigNumber(-981));
        let _index186 = _module.__default.safeIndex(new BigNumber(222), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index186] = !(_module.__default.fm30(new BigNumber((_1230_v73).length), globalState));
        let _1231_v74;
        _1231_v74 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1232_v75;
        _1232_v75 = _module.D9.create_DC22((_this).f17, (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], (_this).f18, _module.__default.fm36(p2, new _dafny.CodePoint('k'.codePointAt(0)), p1, globalState), _dafny.Seq.update(_1221_v64, _module.__default.safeIndex(_1150_v0, new BigNumber((_1221_v64).length)), _1231_v74));
        let _index187 = _module.__default.safeIndex(new BigNumber(222), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index187] = (_1232_v75).dtor_cf37;
        _1150_v0 = (_1150_v0).plus(new BigNumber(398));
        let _1233_v76;
        _1233_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1220_v63,_1150_v0);
        let _1234_v77;
        _1234_v77 = _dafny.MultiSet.fromElements((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], new BigNumber(32), new BigNumber((_1233_v76).length), (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]);
        let _1235_v78;
        _1235_v78 = _dafny.MultiSet.fromElements(new BigNumber((_1234_v77).cardinality()));
        let _1236_v79;
        _1236_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(121),_1235_v78);
        let _1237_v80;
        _1237_v80 = _dafny.Seq.of(_1233_v76, _1233_v76, _1233_v76, _1233_v76);
        let _1238_v81;
        _1238_v81 = _dafny.MultiSet.fromElements(_1233_v76, _1233_v76, _1233_v76, (_1237_v80)[_module.__default.safeIndex((_1229_v72).dtor_cf20, new BigNumber((_1237_v80).length))]);
        let _index188 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
        let _rhs154 = (_dafny.ZERO).minus(((_this).fm8(_1220_v63, _this.f16, globalState)).minus(new BigNumber(((_1236_v79).Merge((_1236_v79).update((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], _1235_v78))).length)));
        let _rhs155 = ((_module.__default.fm0(_1222_v65, (_this).f18, (_this).f18, new BigNumber((_1238_v81).cardinality()), globalState)).plus(_1150_v0)).multipliedBy((_dafny.ZERO).minus((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]));
        let _lhs126 = globalState;
        let _lhs127 = _1218_v61;
        let _lhs128 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length));
        _lhs126.f7 = _rhs154;
        _lhs127[_lhs128] = _rhs155;
        let _1239_v82;
        _1239_v82 = _module.D6.create_DC13(((_this).f14)[_module.__default.safeIndex(new BigNumber(222), new BigNumber(((_this).f14).length))], (_this).f18, (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], p2);
        let _1240_v83;
        _1240_v83 = _module.D6.create_DC14(_1239_v82);
        _1221_v64 = _dafny.Seq.Concat(((((_this).f14)[_module.__default.safeIndex(new BigNumber(222), new BigNumber(((_this).f14).length))]) ? (_module.__default.fm32(_1240_v83, globalState)) : (_1221_v64)), (_1225_v68)[_module.__default.safeIndex(_1150_v0, new BigNumber((_1225_v68).length))]);
      }
      let _1241_v84;
      _1241_v84 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).fm8(false, _this.f16, globalState));
      let _1242_v85;
      _1242_v85 = _dafny.Seq.of((_1241_v84).Merge(_1241_v84));
      let _1243_v86;
      _1243_v86 = _dafny.Map.Empty.slice().updateUnsafe((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))],_dafny.Seq.of(_1241_v84));
      _1242_v85 = _dafny.Seq.Concat((((_1243_v86).contains(_1150_v0)) ? ((_1243_v86).get(_1150_v0)) : (_dafny.Seq.of((_1242_v85)[_module.__default.safeIndex((_dafny.ZERO).minus((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]), new BigNumber((_1242_v85).length))]))), _1242_v85);
      let _1244_i2;
      _1244_i2 = _dafny.ZERO;
      L5: {
        while (p1) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1244_i2)) {
              break L5;
            }
            _1244_i2 = (_1244_i2).plus(_dafny.ONE);
            let _index189 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((p0).length));
            (p0)[_index189] = false;
            let _index190 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((p0).length));
            (p0)[_index190] = p2;
            let _1245_v87;
            _1245_v87 = _dafny.MultiSet.fromElements((_this).f18, true);
            let _index191 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((p0).length));
            let _index192 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((p0).length));
            let _rhs156 = (_this).f13;
            let _rhs157 = true;
            let _rhs158 = (_this).f18;
            let _rhs159 = (_dafny.Seq.IsPrefixOf(_1217_v60, _dafny.Seq.update(_module.__default.fm34((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], globalState), _module.__default.safeIndex(_1150_v0, new BigNumber((_module.__default.fm34((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], globalState)).length)), (_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]))) && ((new BigNumber((_1245_v87).cardinality())).isLessThanOrEqualTo(new BigNumber((_1221_v64).length)));
            let _rhs160 = (_1217_v60)[_module.__default.safeIndex((((_this).f13) ? (new BigNumber(250)) : (_1150_v0)), new BigNumber((_1217_v60).length))];
            let _lhs129 = p0;
            let _lhs130 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((p0).length));
            let _lhs131 = p0;
            let _lhs132 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((p0).length));
            _lhs129[_lhs130] = _rhs156;
            _lhs131[_lhs132] = _rhs157;
            _1220_v63 = _rhs158;
            _1220_v63 = _rhs159;
            _1150_v0 = _rhs160;
            let _1246_v88;
            let _init29 = ((_1247_v0) => function (_1248_i3) {
              return _module.D6.create_DC11(_dafny.MultiSet.fromElements(new BigNumber(313), _1247_v0));
            })(_1150_v0);
            let _nw167 = Array((new BigNumber(6)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw167.length); _i0_29++) {
              _nw167[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _1246_v88 = _nw167;
            _1246_v88 = _1246_v88;
            (globalState).f2 = _1222_v65;
            (globalState).f7 = (new BigNumber(-41)).multipliedBy(((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))]).multipliedBy(_1150_v0));
          }
        }
      }
      let _1249_v89;
      _1249_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v0,_1221_v64);
      _1249_v89 = (_1249_v89).update((_1218_v61)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_1218_v61).length))], _1221_v64);
      return;
    }
    m12(globalState) {
      let _this = this;
      let _1250_v0;
      _1250_v0 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1251_v1;
      let _nw168 = new _module.C0();
      _nw168.__ctor(true, (((_this).f13) ? (new _dafny.CodePoint('p'.codePointAt(0))) : (_1250_v0)), _this.f12, (_this).f18);
      _1251_v1 = _nw168;
      let _1252_v2;
      _1252_v2 = new BigNumber(385);
      (globalState).f7 = _1252_v2;
      (_1251_v1).f19 = (_this).f18;
      let _1253_v3;
      _1253_v3 = _dafny.Seq.UnicodeFromString("j");
      let _1254_v4;
      _1254_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1253_v3,(_this).f18);
      _1254_v4 = _1254_v4;
      _1253_v3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), ((_1255_v0) => function (_1256_i0) {
        return _1255_v0;
      })(_1250_v0));
      let _hi6 = _1252_v2;
      for (let _1257_i1 = _1252_v2; _1257_i1.isLessThan(_hi6); _1257_i1 = _1257_i1.plus(_dafny.ONE)) {
        (globalState).f7 = _module.__default.safeModuloInt(_1257_i1, new BigNumber(-145));
        let _1258_v5;
        let _nw169 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _1258_v5 = _nw169;
        let _1259_v6;
        _1259_v6 = _dafny.Set.fromElements(_1251_v1.f19, _1251_v1.f19);
        let _1260_v7;
        _1260_v7 = _dafny.Seq.of(new BigNumber((_1259_v6).length), _1252_v2, _1252_v2, _1252_v2);
        let _1261_v8;
        _1261_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1258_v5,_1260_v7);
        let _1262_v9;
        _1262_v9 = _dafny.Seq.of(_1252_v2, new BigNumber((_1261_v8).length));
        let _1263_v10;
        _1263_v10 = _module.D9.create_DC22(_1251_v1.f19, (_1260_v7)[_module.__default.safeIndex(_1252_v2, new BigNumber((_1260_v7).length))], (_this).f13, _1259_v6, _1253_v3);
        let _1264_v11;
        _1264_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1252_v2).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f18), _module.__default.safeIndex(_1257_i1, new BigNumber((_dafny.Seq.of((_this).f18)).length)), (_this).f13)).length)),(_1263_v10).dtor_cf38);
        _1264_v11 = _1264_v11;
        (globalState).f7 = _1257_i1;
        if (((function () {
          let _coll44 = new _dafny.Set();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(892), new BigNumber(210))) {
            let _1265_v12 = _compr_44;
            if (((new BigNumber(892)).isLessThanOrEqualTo(_1265_v12)) && ((_1265_v12).isLessThan(new BigNumber(210)))) {
              _coll44.add((_1265_v12).multipliedBy(_1252_v2));
            }
          }
          return _coll44;
        }()).Difference((_this).f15)).equals(((_this).f15).Difference((_this).f15))) {
          let _1266_v13;
          _1266_v13 = _module.D3.create_DC6(_1257_i1, _1251_v1.f19);
          let _1267_v14;
          let _nw170 = Array((new BigNumber(29)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = (_1251_v1).f20;
          _nw170[(_dafny.ONE).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(2)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(3)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(4)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(5)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(6)).toNumber()] = _1250_v0;
          _nw170[(new BigNumber(7)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(8)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw170[(new BigNumber(10)).toNumber()] = _1250_v0;
          _nw170[(new BigNumber(11)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(12)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(13)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(14)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(15)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(16)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(17)).toNumber()] = _module.__default.fm18(_1252_v2, _1253_v3, (_this).f17, _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_dafny.ZERO).minus(_1257_i1)), globalState);
          _nw170[(new BigNumber(18)).toNumber()] = _module.__default.fm33((_1266_v13).dtor_cf13, _1257_i1, globalState);
          _nw170[(new BigNumber(19)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw170[(new BigNumber(21)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(22)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(23)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(24)).toNumber()] = _1250_v0;
          _nw170[(new BigNumber(25)).toNumber()] = (_1253_v3)[_module.__default.safeIndex(_1257_i1, new BigNumber((_1253_v3).length))];
          _nw170[(new BigNumber(26)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(27)).toNumber()] = (_1251_v1).f20;
          _nw170[(new BigNumber(28)).toNumber()] = (_1251_v1).f20;
          _1267_v14 = _nw170;
          let _1268_v15;
          _1268_v15 = _module.D3.create_DC5((_1251_v1).f20);
          let _index193 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1267_v14).length));
          (_1267_v14)[_index193] = (_1268_v15).dtor_cf11;
          let _1269_v16;
          _1269_v16 = _dafny.Seq.of(_1251_v1.f19, !((_this).f17), _module.__default.fm30(_1252_v2, globalState));
          let _1270_v17;
          _1270_v17 = _dafny.MultiSet.fromElements((_this).f17);
          let _1271_v18;
          _1271_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1251_v1.f19,new BigNumber((_1270_v17).cardinality()));
          let _1272_v19;
          _1272_v19 = _dafny.Seq.of((_this).f18, (_this).f13, !(_1251_v1.f19), (_this).f13, (_1269_v16)[_module.__default.safeIndex(new BigNumber((_1271_v18).length), new BigNumber((_1269_v16).length))]);
          let _1273_v20;
          _1273_v20 = _dafny.Seq.of(_1269_v16, _1272_v19);
          let _index194 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1267_v14).length));
          let _rhs161 = _dafny.Seq.Concat(_1269_v16, (_1273_v20)[_module.__default.safeIndex(_1257_i1, new BigNumber((_1273_v20).length))]);
          let _rhs162 = (_1263_v10).dtor_cf39;
          let _rhs163 = _1250_v0;
          let _rhs164 = (_1251_v1).f20;
          let _lhs133 = globalState;
          let _lhs134 = _1251_v1;
          let _lhs135 = _1267_v14;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1267_v14).length));
          _lhs133.f2 = _rhs161;
          _lhs134.f19 = _rhs162;
          _lhs135[_lhs136] = _rhs163;
          _1250_v0 = _rhs164;
          (_1251_v1).f19 = (_this).f17;
          let _1274_v21;
          _1274_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1257_i1,_1251_v1.f19);
          (_1251_v1).f19 = (((_1274_v21).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1272_v19).length),_1257_i1)).length))) ? ((_1274_v21).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1272_v19).length),_1257_i1)).length))) : (_dafny.areEqual(_1260_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_1275_v3) => function (_1276_i2) {
            return new BigNumber((_1275_v3).length);
          })(_1253_v3)))));
          let _index195 = _module.__default.safeIndex(new BigNumber(78), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index195] = _1251_v1.f19;
          let _index196 = _module.__default.safeIndex(new BigNumber(78), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index196] = !(!((_1252_v2).isLessThanOrEqualTo(new BigNumber(509)))) || ((_1269_v16)[_module.__default.safeIndex(_1252_v2, new BigNumber((_1269_v16).length))]);
          (_1251_v1).f19 = ((_this).f18) && (_1251_v1.f19);
        } else {
          (_1251_v1).f19 = (_this.f12).dtor_cf4;
          (globalState).f7 = ((_1257_i1).plus(new BigNumber((_1253_v3).length))).multipliedBy(_1252_v2);
          let _1277_v22;
          _1277_v22 = _dafny.Seq.of(_1253_v3);
          let _rhs165 = _dafny.Seq.Concat(_1253_v3, _1253_v3);
          let _rhs166 = _1257_i1;
          let _rhs167 = _dafny.Seq.Concat(_1277_v22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-342)), ((_1278_v3) => function (_1279_i3) {
            return _1278_v3;
          })(_1253_v3)));
          let _lhs137 = globalState;
          _1253_v3 = _rhs165;
          _lhs137.f7 = _rhs166;
          _1277_v22 = _rhs167;
          (_1251_v1).f19 = _1251_v1.f19;
          _1253_v3 = _1253_v3;
        }
      }
      return;
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm14(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(963)), function (_1280_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }));
    };
    m10(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.Seq.of();
      let _1281_v0;
      _1281_v0 = _dafny.Seq.UnicodeFromString("gwx");
      _1281_v0 = _1281_v0;
      let _1282_v1;
      let _nw171 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1282_v1 = _nw171;
      let _1283_v2;
      _1283_v2 = _dafny.Seq.of(_1282_v1);
      let _hi7 = new BigNumber((_1283_v2).length);
      for (let _1284_i0 = ((p0) ? (new BigNumber(450)) : (new BigNumber(-597))); _1284_i0.isLessThan(_hi7); _1284_i0 = _1284_i0.plus(_dafny.ONE)) {
        let _1285_v3;
        _1285_v3 = _dafny.Seq.of(_1284_i0);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(-808)), _1285_v3)) {
          let _1286_v4;
          _1286_v4 = false;
          _1286_v4 = _1286_v4;
          let _1287_v5;
          let _nw172 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1287_v5 = _nw172;
          _1287_v5 = _1287_v5;
          let _1288_v6;
          _1288_v6 = _module.D1.create_DC2(new BigNumber(130), _1284_i0, p0);
          let _1289_v7;
          let _nw173 = new _module.C1();
          _nw173.__ctor(p2, _1288_v6, p0);
          _1289_v7 = _nw173;
          let _index197 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1287_v5).length));
          (_1287_v5)[_index197] = !(!(!_dafny.Seq.contains(_dafny.Seq.of(!(_1286_v4)), p0)));
          let _1290_v8;
          _1290_v8 = _dafny.Set.fromElements(_1286_v4);
          let _index198 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1287_v5).length));
          let _rhs168 = (new BigNumber((_1281_v0).length)).isLessThanOrEqualTo(new BigNumber((_1290_v8).length));
          let _rhs169 = _1281_v0;
          let _rhs170 = (_1285_v3)[_module.__default.safeIndex(_1284_i0, new BigNumber((_1285_v3).length))];
          let _lhs138 = _1287_v5;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1287_v5).length));
          let _lhs140 = globalState;
          _lhs138[_lhs139] = _rhs168;
          _1281_v0 = _rhs169;
          _lhs140.f7 = _rhs170;
          let _index199 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1287_v5).length));
          (_1287_v5)[_index199] = !(p0);
        } else {
          let _1291_v9;
          _1291_v9 = _dafny.Seq.of(p0);
          let _1292_v10;
          _1292_v10 = _module.D1.create_DC1(_1281_v0);
          let _1293_v11;
          let _nw174 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1293_v11 = _nw174;
          let _1294_v12;
          _1294_v12 = _dafny.Set.fromElements(_1284_i0);
          let _pat_let_tv33 = _1281_v0;
          let _1295_v13;
          let _nw175 = new _module.C2();
          _nw175.__ctor(_1284_i0, (_1291_v9)[_module.__default.safeIndex(_1284_i0, new BigNumber((_1291_v9).length))], function (_pat_let33_0) {
            return function (_1296_dt__update__tmp_h0) {
              return function (_pat_let34_0) {
                return function (_1297_dt__update_hcf1_h0) {
                  return _module.D1.create_DC1(_1297_dt__update_hcf1_h0);
                }(_pat_let34_0);
              }(_pat_let_tv33);
            }(_pat_let33_0);
          }(_1292_v10), p0, _1293_v11, _1294_v12, _module.D1.create_DC2(_1284_i0, _1284_i0, p0), !(p0) || (p0));
          _1295_v13 = _nw175;
          let _1298_v14;
          _1298_v14 = _dafny.Set.fromElements(true);
          let _1299_v15;
          _1299_v15 = _module.D9.create_DC22(p0, (_1295_v13).f21, p0, _1298_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), ((_1300_p2) => function (_1301_i1) {
  return _1300_p2;
})(p2)));
          let _1302_v16;
          _1302_v16 = _dafny.Set.fromElements(_1299_v15, _1299_v15, _module.D9.create_DC22(p0, (_1295_v13).f21, p0, _1298_v14, _dafny.Seq.UnicodeFromString("m")));
          let _1303_v17;
          _1303_v17 = _module.D3.create_DC5(p2);
          let _1304_v18;
          let _nw176 = Array((new BigNumber(16)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = p2;
          _nw176[(_dafny.ONE).toNumber()] = (_1303_v17).dtor_cf11;
          _nw176[(new BigNumber(2)).toNumber()] = p2;
          _nw176[(new BigNumber(3)).toNumber()] = p2;
          _nw176[(new BigNumber(4)).toNumber()] = p2;
          _nw176[(new BigNumber(5)).toNumber()] = p2;
          _nw176[(new BigNumber(6)).toNumber()] = p2;
          _nw176[(new BigNumber(7)).toNumber()] = p2;
          _nw176[(new BigNumber(8)).toNumber()] = p2;
          _nw176[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw176[(new BigNumber(10)).toNumber()] = p2;
          _nw176[(new BigNumber(11)).toNumber()] = p2;
          _nw176[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw176[(new BigNumber(13)).toNumber()] = p2;
          _nw176[(new BigNumber(14)).toNumber()] = p2;
          _nw176[(new BigNumber(15)).toNumber()] = p2;
          _1304_v18 = _nw176;
          let _1305_v19;
          _1305_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1304_v18,new BigNumber((_1285_v3).length));
          let _1306_v20;
          _1306_v20 = _dafny.MultiSet.fromElements(new BigNumber((_1302_v16).length), (((_1305_v19).contains(_1304_v18)) ? ((_1305_v19).get(_1304_v18)) : (_1284_i0)));
          let _1307_v21;
          _1307_v21 = _module.D7.create_DC16((_1295_v13).f21);
          let _1308_v22;
          _1308_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1306_v20);
          let _1309_v23;
          _1309_v23 = _module.D6.create_DC11(_dafny.MultiSet.fromElements(_1284_i0, new BigNumber((_1281_v0).length), (_1295_v13).f21));
          let _1310_v24;
          let _nw177 = Array((new BigNumber(22)).toNumber());
          _nw177[(_dafny.ZERO).toNumber()] = _1306_v20;
          _nw177[(_dafny.ONE).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_1284_i0, new BigNumber((_1298_v14).length)));
          _nw177[(new BigNumber(2)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(3)).toNumber()] = (_1306_v20).Intersect(_1306_v20);
          _nw177[(new BigNumber(4)).toNumber()] = (_1306_v20).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(981), _1284_i0, new BigNumber((_dafny.Seq.of(!((_1295_v13).f22))).length))).length), (_1307_v21).dtor_cf29));
          _nw177[(new BigNumber(5)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(6)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(7)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(8)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(9)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(10)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(_1284_i0)).Difference(_1306_v20);
          _nw177[(new BigNumber(12)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(13)).toNumber()] = (((_1308_v22).contains((_1295_v13).f22)) ? ((_1308_v22).get((_1295_v13).f22)) : (_1306_v20));
          _nw177[(new BigNumber(14)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(15)).toNumber()] = (_1306_v20).Union((_1306_v20).update(new BigNumber((_1306_v20).cardinality()), _module.__default.abs(new BigNumber(473))));
          _nw177[(new BigNumber(16)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(17)).toNumber()] = (_1306_v20).Intersect(_1306_v20);
          _nw177[(new BigNumber(18)).toNumber()] = (_1306_v20).Intersect(_1306_v20);
          _nw177[(new BigNumber(19)).toNumber()] = _1306_v20;
          _nw177[(new BigNumber(20)).toNumber()] = _module.__default.fm19(globalState);
          _nw177[(new BigNumber(21)).toNumber()] = ((_1309_v23).dtor_cf17).Difference(_dafny.MultiSet.FromArray(_1285_v3));
          _1310_v24 = _nw177;
          let _index200 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_1310_v24).length));
          (_1310_v24)[_index200] = _1306_v20;
          let _index201 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_1310_v24).length));
          (_1310_v24)[_index201] = _1306_v20;
          let _1311_v25;
          _1311_v25 = false;
          let _1312_v26;
          let _nw178 = Array((new BigNumber(28)).toNumber()).fill([]);
          _1312_v26 = _nw178;
          let _1313_v27;
          let _nw179 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
          _1313_v27 = _nw179;
          let _index202 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1312_v26).length));
          (_1312_v26)[_index202] = _1313_v27;
          let _index203 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          (_1293_v11)[_index203] = (_1295_v13).f22;
          let _1314_v28;
          _1314_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1284_i0).minus(new BigNumber(-417)),_1313_v27);
          let _index204 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1312_v26).length));
          let _index205 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          let _rhs171 = _dafny.Seq.contains(_1281_v0, p2);
          let _rhs172 = (((_1314_v28).contains(new BigNumber(183))) ? ((_1314_v28).get(new BigNumber(183))) : (_1313_v27));
          let _rhs173 = false;
          let _rhs174 = _1282_v1;
          let _lhs141 = _1312_v26;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1312_v26).length));
          let _lhs143 = _1293_v11;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          _1311_v25 = _rhs171;
          _lhs141[_lhs142] = _rhs172;
          _lhs143[_lhs144] = _rhs173;
          _1282_v1 = _rhs174;
          let _1315_v29;
          _1315_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1295_v13).f22,(_1293_v11)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length))]);
          let _index206 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          let _index207 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          let _rhs175 = (_1293_v11)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length))];
          let _rhs176 = _1306_v20;
          let _rhs177 = !(((_1295_v13).f21).isLessThanOrEqualTo((_1295_v13).f21));
          let _rhs178 = !((((_1295_v13).f22) ? (!_dafny.Seq.contains(_1285_v3, _1284_i0)) : ((((_1315_v29).contains(_1311_v25)) ? ((_1315_v29).get(_1311_v25)) : ((_1293_v11)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length))])))));
          let _lhs145 = _1293_v11;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          let _lhs147 = _1293_v11;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length));
          _lhs145[_lhs146] = _rhs175;
          _1306_v20 = _rhs176;
          _lhs147[_lhs148] = _rhs177;
          _1311_v25 = _rhs178;
          let _1316_v30;
          _1316_v30 = _dafny.MultiSet.fromElements((_1295_v13).f22);
          let _1317_v31;
          _1317_v31 = _module.D1.create_DC2(new BigNumber(236), (_1295_v13).f21, (_1295_v13).f22);
          let _1318_v32;
          _1318_v32 = _dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), p2, p2);
          let _pat_let_tv34 = _1281_v0;
          let _1319_v33;
          let _nw180 = new _module.C3();
          _nw180.__ctor((_1316_v30).IsDisjointFrom(_1316_v30), function (_pat_let35_0) {
            return function (_1320_dt__update__tmp_h1) {
              return function (_pat_let36_0) {
                return function (_1321_dt__update_hcf1_h1) {
                  return _module.D1.create_DC1(_1321_dt__update_hcf1_h1);
                }(_pat_let36_0);
              }(_pat_let_tv34);
            }(_pat_let35_0);
          }(_1292_v10), (_1295_v13).f22, _1293_v11, _1294_v12, _1317_v31, (_dafny.Set.fromElements(_module.__default.fm33((_1293_v11)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_1293_v11).length))], (_1295_v13).f21, globalState))).IsSubsetOf(_1318_v32));
          _1319_v33 = _nw180;
        }
        (globalState).f7 = ((p0) ? (_1284_i0) : (_1284_i0));
        let _1322_v34;
        _1322_v34 = _module.D6.create_DC12(p0, _1281_v0, _1284_i0, p0);
        let _1323_v35;
        _1323_v35 = _dafny.Set.fromElements(new BigNumber(235), _1284_i0);
        let _1324_v36;
        let _nw181 = Array((new BigNumber(4)).toNumber());
        _nw181[(_dafny.ZERO).toNumber()] = p0;
        _nw181[(_dafny.ONE).toNumber()] = (_1322_v34).dtor_cf21;
        _nw181[(new BigNumber(2)).toNumber()] = p0;
        _nw181[(new BigNumber(3)).toNumber()] = (_1323_v35).IsProperSubsetOf(_1323_v35);
        _1324_v36 = _nw181;
        let _index208 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1324_v36).length));
        (_1324_v36)[_index208] = !(p0);
        let _index209 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1324_v36).length));
        (_1324_v36)[_index209] = !((p0) || (p0));
        let _1325_v37;
        _1325_v37 = _module.D1.create_DC2(_1284_i0, _1284_i0, (_1324_v36)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_1324_v36).length))]);
        let _1326_v38;
        let _nw182 = new _module.C1();
        _nw182.__ctor(p2, _1325_v37, p0);
        _1326_v38 = _nw182;
      }
      let _1327_v39;
      _1327_v39 = _dafny.Set.fromElements(p0);
      let _1328_v40;
      _1328_v40 = new BigNumber(706);
      let _rhs179 = _1327_v39;
      let _rhs180 = _1328_v40;
      let _rhs181 = _1328_v40;
      let _lhs149 = globalState;
      let _lhs150 = globalState;
      _1327_v39 = _rhs179;
      _lhs149.f7 = _rhs180;
      _lhs150.f7 = _rhs181;
      let _1329_v41;
      _1329_v41 = _module.D1.create_DC1(_1281_v0);
      let _1330_v42;
      let _init30 = ((_1331_p0) => function (_1332_i2) {
        return _1331_p0;
      })(p0);
      let _nw183 = Array((new BigNumber(4)).toNumber());
      for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw183.length); _i0_30++) {
        _nw183[_i0_30] = _init30(new BigNumber(_i0_30));
      }
      _1330_v42 = _nw183;
      let _1333_v43;
      _1333_v43 = _dafny.Set.fromElements(_1328_v40, _1328_v40);
      let _1334_v44;
      _1334_v44 = _module.D1.create_DC2(_1328_v40, _1328_v40, !(p0));
      let _1335_v45;
      let _nw184 = new _module.C3();
      _nw184.__ctor(!(_1327_v39).contains(p0), _1329_v41, ((p0) ? (p0) : (!(p0))), _1330_v42, _1333_v43, _1334_v44, true);
      _1335_v45 = _nw184;
      let _1336_v46;
      _1336_v46 = true;
      _1336_v46 = p0;
      let _hi8 = new BigNumber(881);
      for (let _1337_i3 = _1328_v40; _1337_i3.isLessThan(_hi8); _1337_i3 = _1337_i3.plus(_dafny.ONE)) {
        (globalState).f7 = _1337_i3;
        let _1338_v47;
        _1338_v47 = new _dafny.CodePoint('f'.codePointAt(0));
        _1338_v47 = p2;
        let _1339_v48;
        let _nw185 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1339_v48 = _nw185;
        let _rhs182 = _1339_v48;
        let _rhs183 = p0;
        let _rhs184 = _dafny.Seq.IsProperPrefixOf(_1281_v0, _1281_v0);
        _1339_v48 = _rhs182;
        _1336_v46 = _rhs183;
        _1336_v46 = _rhs184;
        if ((_1335_v45).f18) {
          let _1340_v49;
          _1340_v49 = _dafny.Set.fromElements(p2, p2);
          _1340_v49 = _1340_v49;
          (globalState).f7 = _1328_v40;
          _1336_v46 = _1336_v46;
          let _1341_v50;
          _1341_v50 = _dafny.Seq.of(_1336_v46, false);
          let _1342_v51;
          _1342_v51 = _dafny.Seq.of(!(_1328_v40).isEqualTo(new BigNumber((_1341_v50).length)), (_1335_v45).f18, false);
          (globalState).f2 = _1342_v51;
          let _1343_v52;
          _1343_v52 = _dafny.MultiSet.fromElements(_1282_v1);
          let _1344_v53;
          let _nw186 = new _module.C2();
          _nw186.__ctor(_1337_i3, !(!((_1335_v45).f18)) || ((_1335_v45).f18), _1329_v41, (_1342_v51)[_module.__default.safeIndex(_1337_i3, new BigNumber((_1342_v51).length))], _1330_v42, _dafny.Set.fromElements(_1328_v40), _1334_v44, (_1343_v52).IsProperSubsetOf(_1343_v52));
          _1344_v53 = _nw186;
        } else {
          (globalState).f7 = new BigNumber(75);
          _1336_v46 = (_1335_v45).f18;
          let _1345_v54;
          let _nw187 = new _module.C0();
          _nw187.__ctor((_1335_v45).f18, p2, _module.D1.create_DC2(_1328_v40, _1337_i3, (_1335_v45).f18), !(_module.__default.fm17(globalState)));
          _1345_v54 = _nw187;
          let _1346_v55;
          let _nw188 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
          _1346_v55 = _nw188;
          _1346_v55 = _1346_v55;
          let _1347_v56;
          _1347_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1345_v54.f19,_1328_v40);
          let _1348_v57;
          _1348_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1347_v56,new BigNumber(-568));
          let _1349_v58;
          _1349_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1348_v57,_1336_v46);
          _1349_v58 = (_1349_v58).update(_1348_v57, (_1335_v45).f18);
        }
      }
      r0 = p1;
      r1 = _1283_v2;
      return [r0, r1];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f12 = _module.D1.Default();
      this._f14 = [];
      this._f15 = _dafny.Set.Empty;
      this._f13 = false;
      this.f24 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    __ctor(f24, f14, f15, f12, f13) {
      let _this = this;
      (_this).f24 = f24;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-193),new BigNumber(525))).length)), _dafny.Set.fromElements(false, _this.f24, (_module.D12.create_DC32(_this.f24, new BigNumber((_dafny.Seq.UnicodeFromString("wbv")).length), false, _this.f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber(732))).length))).dtor_cf67, _this.f24, _this.f24), _dafny.MultiSet.fromElements(new BigNumber(144), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-496), new BigNumber(687))).cardinality())), new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber(-869)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(342),_this.f24)).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f13, _this.f24)).cardinality())), _module.D1.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_1350_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})))).dtor_cf32;
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((((_this).f15).Difference((_this).f15)).length);
    };
    fm9(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this.f24) ? (_module.D12.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f13),(_this).f13))) : (_module.D12.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),_this.f24)))),new BigNumber(((_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D7.create_DC17(_module.D7.create_DC17(_module.D7.create_DC17(_module.D7.create_DC16(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(460)), function (_1351_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length))))), _module.D7.create_DC17(_module.D7.create_DC15(_dafny.Seq.of(new BigNumber(324))))))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D7.create_DC17(_module.D7.create_DC16(new BigNumber((_dafny.Set.fromElements(_this.f24, (_this).f13)).length))), _module.D7.create_DC17(_module.D7.create_DC15(_dafny.Seq.of(new BigNumber(-703), new BigNumber(177)))), _module.D7.create_DC17(_module.D7.create_DC17(_module.D7.create_DC17(_module.D7.create_DC17(_module.D7.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-543)), function (_1352_i1) {
  return new BigNumber(387);
}))))))), _dafny.Seq.of(_module.D7.create_DC17(_module.D7.create_DC16((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)))).length))).length))))))))).cardinality()))).length);
    };
    fm52(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(731)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber(459), new BigNumber((_dafny.Seq.of(new BigNumber(-683))).length)));
    };
    fm53(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((new BigNumber(347)).plus((new BigNumber(111)).plus(new BigNumber(19))));
    };
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _index210 = _module.__default.safeIndex(new BigNumber(855), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index210] = false;
      let _1353_v0;
      _1353_v0 = _dafny.Seq.UnicodeFromString("av");
      let _1354_v1;
      _1354_v1 = _dafny.Seq.of(false, (_this).f13);
      let _1355_v2;
      _1355_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,p0);
      let _1356_v3;
      let _nw189 = Array((new BigNumber(14)).toNumber());
      _nw189[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw189[(_dafny.ONE).toNumber()] = (_this).fm8(_this.f24, _module.D1.create_DC1(_1353_v0), globalState);
      _nw189[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
      _nw189[(new BigNumber(3)).toNumber()] = (new BigNumber((_1354_v1).length)).multipliedBy(new BigNumber((_1355_v2).length));
      _nw189[(new BigNumber(4)).toNumber()] = (p0).multipliedBy(p0);
      _nw189[(new BigNumber(5)).toNumber()] = p0;
      _nw189[(new BigNumber(6)).toNumber()] = new BigNumber(370);
      _nw189[(new BigNumber(7)).toNumber()] = p0;
      _nw189[(new BigNumber(8)).toNumber()] = p0;
      _nw189[(new BigNumber(9)).toNumber()] = p0;
      _nw189[(new BigNumber(10)).toNumber()] = p0;
      _nw189[(new BigNumber(11)).toNumber()] = p0;
      _nw189[(new BigNumber(12)).toNumber()] = p0;
      _nw189[(new BigNumber(13)).toNumber()] = p0;
      _1356_v3 = _nw189;
      let _index211 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      (_1356_v3)[_index211] = ((_this.f24) ? (p0) : (new BigNumber((_1355_v2).length)));
      let _1357_v4;
      let _nw190 = Array((new BigNumber(12)).toNumber()).fill(_module.D12.Default());
      _1357_v4 = _nw190;
      let _index212 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1357_v4).length));
      (_1357_v4)[_index212] = _module.D12.create_DC30((_this).f14, (_this).f13);
      let _1358_v5;
      _1358_v5 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1359_v6;
      _1359_v6 = _dafny.MultiSet.fromElements(_1358_v5);
      let _1360_v7;
      _1360_v7 = _dafny.MultiSet.fromElements(false, false, _this.f24, (_module.D12.create_DC31((_this).f13, _1356_v3, (_this).f13, new _dafny.CodePoint('l'.codePointAt(0)), _1359_v6)).dtor_cf62);
      let _1361_v8;
      _1361_v8 = _dafny.MultiSet.fromElements(_1356_v3);
      let _1362_v9;
      _1362_v9 = _dafny.Seq.of(_1361_v8, _1361_v8);
      let _index213 = _module.__default.safeIndex(new BigNumber(855), new BigNumber(((_this).f14).length));
      let _index214 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      let _index215 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1357_v4).length));
      let _rhs185 = p0;
      let _rhs186 = !(!(!(false) || ((((_this).f13) ? (_this.f24) : (_this.f24)))));
      let _rhs187 = _module.__default.safeModuloInt(new BigNumber((_1354_v1).length), new BigNumber((((_this.f24) ? (_1360_v7) : (_1360_v7))).cardinality()));
      let _rhs188 = _module.D12.create_DC30((((_this).f13) ? ((_this).f14) : ((_this).f14)), ((_1362_v9)[_module.__default.safeIndex(p0, new BigNumber((_1362_v9).length))]).IsProperSubsetOf(_1361_v8));
      let _lhs151 = globalState;
      let _lhs152 = (_this).f14;
      let _lhs153 = _module.__default.safeIndex(new BigNumber(855), new BigNumber(((_this).f14).length));
      let _lhs154 = _1356_v3;
      let _lhs155 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      let _lhs156 = _1357_v4;
      let _lhs157 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1357_v4).length));
      _lhs151.f7 = _rhs185;
      _lhs152[_lhs153] = _rhs186;
      _lhs154[_lhs155] = _rhs187;
      _lhs156[_lhs157] = _rhs188;
      let _1363_v10;
      _1363_v10 = _dafny.Seq.of(p0);
      let _1364_v11;
      _1364_v11 = _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("rkrn"));
      let _1365_v12;
      let _nw191 = Array((new BigNumber(13)).toNumber()).fill(false);
      _1365_v12 = _nw191;
      let _pat_let_tv35 = _1358_v5;
      let _1366_v13;
      let _nw192 = new _module.C2();
      _nw192.__ctor((new BigNumber((_1363_v10).length)).plus((_this.f12).dtor_cf3), _this.f24, function (_pat_let37_0) {
        return function (_1367_dt__update__tmp_h0) {
          return function (_pat_let38_0) {
            return function (_1370_dt__update_hcf1_h0) {
              return _module.D1.create_DC1(_1370_dt__update_hcf1_h0);
            }(_pat_let38_0);
          }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_1368_v5) => function (_1369_i0) {
            return _1368_v5;
          })(_pat_let_tv35)));
        }(_pat_let37_0);
      }(_1364_v11), !(((_1356_v3)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length))]).isLessThan((_dafny.ZERO).minus(new BigNumber((_1353_v0).length)))), _1365_v12, (_this).f15, _this.f12, (_this).f13);
      _1366_v13 = _nw192;
      let _1371_v14;
      _1371_v14 = _module.D9.create_DC22((_this).f13, (_1356_v3)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length))], (_this).f13, _module.__default.fm54(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(623)), ((_1372_v0, _1373_v13) => function (_1374_i1) {
  return new BigNumber((_dafny.Seq.update(_1372_v0, _module.__default.safeIndex((_1373_v13).f21, new BigNumber((_1372_v0).length)), (_1372_v0)[_module.__default.safeIndex((_1373_v13).f21, new BigNumber((_1372_v0).length))])).length);
})(_1353_v0, _1366_v13))).length), globalState), _1353_v0);
      r0 = (_1371_v14).dtor_cf37;
      let _index216 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      (_1356_v3)[_index216] = p0;
      let _index217 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      (_1356_v3)[_index217] = (_1356_v3)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length))];
      let _1375_v15;
      _1375_v15 = _dafny.Seq.of(_1365_v12, _1365_v12);
      let _1376_v16;
      _1376_v16 = _dafny.Set.fromElements(_1365_v12, (_1375_v15)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_1375_v15).length))], _1365_v12, _1365_v12, _1365_v12);
      let _index218 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      let _rhs189 = _1376_v16;
      let _rhs190 = _module.__default.safeModuloInt(new BigNumber((_1354_v1).length), (_1356_v3)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length))]);
      let _rhs191 = (_1356_v3)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length))];
      let _lhs158 = globalState;
      let _lhs159 = _1356_v3;
      let _lhs160 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1356_v3).length));
      _1376_v16 = _rhs189;
      _lhs158.f7 = _rhs190;
      _lhs159[_lhs160] = _rhs191;
      r0 = (_1366_v13).f22;
      return r0;
    }
    m14(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1377_v0;
      let _out10;
      _out10 = (_this).m15(globalState);
      _1377_v0 = _out10;
      let _1378_i0;
      _1378_i0 = _dafny.ZERO;
      L6: {
        while (_this.f24) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1378_i0)) {
              break L6;
            }
            _1378_i0 = (_1378_i0).plus(_dafny.ONE);
            let _1379_v1;
            _1379_v1 = new _dafny.CodePoint('e'.codePointAt(0));
            let _1380_v2;
            _1380_v2 = _module.D3.create_DC5(_1379_v1);
            let _index219 = _module.__default.safeIndex(new BigNumber(698), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index219] = _this.f24;
            let _1381_v3;
            _1381_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v1,new BigNumber(78));
            let _1382_v4;
            _1382_v4 = _dafny.Set.fromElements(_1381_v3);
            let _1383_v5;
            let _nw193 = Array((_dafny.ONE).toNumber());
            _nw193[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm55((_this).fm9(_1379_v1, globalState), _1382_v4, p0, globalState)).cardinality());
            _1383_v5 = _nw193;
            let _index220 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1383_v5).length));
            (_1383_v5)[_index220] = _1377_v0;
            let _index221 = _module.__default.safeIndex(new BigNumber(698), new BigNumber(((_this).f14).length));
            let _index222 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1383_v5).length));
            let _rhs192 = _1380_v2;
            let _rhs193 = ((p0) ? (true) : ((_this).f13));
            let _rhs194 = _1377_v0;
            let _lhs161 = (_this).f14;
            let _lhs162 = _module.__default.safeIndex(new BigNumber(698), new BigNumber(((_this).f14).length));
            let _lhs163 = _1383_v5;
            let _lhs164 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1383_v5).length));
            _1380_v2 = _rhs192;
            _lhs161[_lhs162] = _rhs193;
            _lhs163[_lhs164] = _rhs194;
            let _index223 = _module.__default.safeIndex(new BigNumber(698), new BigNumber(((_this).f14).length));
            ((_this).f14)[_index223] = (_this).f13;
            let _1384_v6;
            _1384_v6 = _module.D11.create_DC26(_1383_v5, p3);
            _1383_v5 = (((p0) ? (_1384_v6) : (_1384_v6))).dtor_cf49;
            r0 = (_1383_v5)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_1383_v5).length))];
          }
        }
      }
      let _1385_v7;
      _1385_v7 = _dafny.Set.fromElements(p2);
      let _1386_v8;
      _1386_v8 = (_1385_v7);
      let _source21 = _1386_v8;
      let _1387___mcc_h0 = _source21;
      let _1388_cf16 = _1387___mcc_h0;
      let _1389_v9;
      _1389_v9 = _dafny.MultiSet.fromElements(_1377_v0, _1377_v0, _1377_v0);
      let _1390_v10;
      let _init31 = ((_1391_v0) => function (_1392_i1) {
        return (_1392_i1).multipliedBy(_1391_v0);
      })(_1377_v0);
      let _nw194 = Array((new BigNumber(6)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw194.length); _i0_31++) {
        _nw194[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1390_v10 = _nw194;
      let _1393_v11;
      _1393_v11 = _module.D11.create_DC26(_1390_v10, p3);
      let _1394_v12;
      _1394_v12 = _dafny.Seq.of(_1393_v11, _module.D11.create_DC26(_1390_v10, p3));
      _1389_v9 = ((_this.f24) ? (_dafny.MultiSet.fromElements(_1377_v0, _1377_v0)) : (_module.__default.fm56(new BigNumber((_1394_v12).length), (_this).f15, globalState)));
      let _index224 = _module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index224] = (_this).f13;
      let _index225 = _module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index225] = p0;
      let _1395_v13;
      _1395_v13 = _dafny.Set.fromElements(_module.D1.create_DC1(p3));
      if ((_1395_v13).IsDisjointFrom(_1395_v13)) {
        let _1396_v14;
        _1396_v14 = _module.D3.create_DC6(_1377_v0, _this.f24);
        let _1397_v15;
        _1397_v15 = _module.D7.create_DC16((_1396_v14).dtor_cf12);
        (globalState).f7 = (_1397_v15).dtor_cf29;
        let _1398_v16;
        _1398_v16 = _dafny.Set.fromElements(false, _this.f24, p0);
        let _1399_v17;
        _1399_v17 = _module.D9.create_DC22((_this).f13, _1377_v0, false, _1398_v16, p2);
        let _1400_v18;
        _1400_v18 = _dafny.Map.Empty.slice().updateUnsafe((_1377_v0).isEqualTo(_1377_v0),((_this).f13) && ((_1399_v17).dtor_cf37));
        _1400_v18 = (_1400_v18).Merge(_1400_v18);
        let _index226 = _module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index226] = (_module.__default.fm57(globalState)).dtor_cf65;
        (_this).f24 = !(!(!(true) || (((p0) ? (((_this).f14)[_module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length))]) : ((_this).f13)))));
        (globalState).f7 = _module.__default.safeModuloInt(_1377_v0, _1377_v0);
      } else {
        let _index227 = _module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index227] = false;
        let _1401_v19;
        let _out11;
        _out11 = _module.__default.m0(new _dafny.CodePoint('j'.codePointAt(0)), p0, globalState);
        _1401_v19 = _out11;
        let _1402_v20;
        _1402_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1377_v0,_1390_v10);
        let _index228 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length));
        (_1390_v10)[_index228] = new BigNumber((_1402_v20).length);
        let _1403_v21;
        _1403_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,p0);
        let _1404_v22;
        _1404_v22 = _dafny.Seq.of(new BigNumber((_1403_v21).length), new BigNumber((_dafny.Seq.of(_1390_v10, _1390_v10)).length));
        let _index229 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length));
        (_1390_v10)[_index229] = (new BigNumber((_1404_v22).length)).plus(_1377_v0);
        let _index230 = _module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index230] = !((_this).f13);
        let _1405_v23;
        let _nw195 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1405_v23 = _nw195;
        let _1406_v24;
        _1406_v24 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index231 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1405_v23).length));
        (_1405_v23)[_index231] = _1406_v24;
        let _1407_v25;
        _1407_v25 = _module.D12.create_DC32(!(_this.f24) || (p0), (_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))], ((_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))]).isLessThan((_dafny.ZERO).minus(new BigNumber((_1404_v22).length))), (_this).f13, (_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))]);
        let _1408_v26;
        _1408_v26 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm58(globalState),(_this).f13);
        let _1409_v27;
        let _nw196 = Array((new BigNumber(6)).toNumber());
        _nw196[(_dafny.ZERO).toNumber()] = _1408_v26;
        _nw196[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1406_v24,false);
        _nw196[(new BigNumber(2)).toNumber()] = (_1408_v26).Merge(_1408_v26);
        _nw196[(new BigNumber(3)).toNumber()] = (_module.__default.fm59(_this.f24, p0, (_this).f13, _1406_v24, globalState)).Merge(_1408_v26);
        _nw196[(new BigNumber(4)).toNumber()] = _1408_v26;
        _nw196[(new BigNumber(5)).toNumber()] = _module.__default.fm59(p0, true, p0, _1406_v24, globalState);
        _1409_v27 = _nw196;
        let _1410_v28;
        _1410_v28 = _dafny.Seq.of((_this).f13, (_this).f13, (_this).f13, false, (_this).f13);
        let _index232 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1405_v23).length));
        let _rhs195 = (((new BigNumber((_1410_v28).length)).isLessThan(_1377_v0)) ? (_1406_v24) : (_1406_v24));
        let _rhs196 = _dafny.Seq.of((_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))], (_1377_v0).minus((_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))]), _module.__default.safeDivisionInt(_1377_v0, _1377_v0));
        let _rhs197 = _module.__default.fm57(globalState);
        let _rhs198 = (_1390_v10)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1390_v10).length))];
        let _rhs199 = _1409_v27;
        let _lhs165 = _1405_v23;
        let _lhs166 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1405_v23).length));
        _lhs165[_lhs166] = _rhs195;
        _1404_v22 = _rhs196;
        _1407_v25 = _rhs197;
        r0 = _rhs198;
        _1409_v27 = _rhs199;
      }
      let _1411_v29;
      _1411_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1390_v10,((_this).f14)[_module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length))]);
      let _1412_v30;
      _1412_v30 = _module.D14.create_DC36((_1411_v29).update(_1390_v10, (_this).f13));
      _1411_v29 = (_dafny.Map.Empty.slice().updateUnsafe(_1390_v10,((_this).f14)[_module.__default.safeIndex(new BigNumber(761), new BigNumber(((_this).f14).length))])).Merge((_1412_v30).dtor_cf74);
      (_this).f24 = (!(p0)) || (p0);
      let _1413_v31;
      let _init32 = function (_1414_i2) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      };
      let _nw197 = Array((_dafny.ONE).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw197.length); _i0_32++) {
        _nw197[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1413_v31 = _nw197;
      let _1415_v32;
      _1415_v32 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1416_v33;
      _1416_v33 = _dafny.Seq.of(_1377_v0, _1377_v0, _1377_v0, _1377_v0, _1377_v0);
      let _1417_v34;
      _1417_v34 = _module.D11.create_DC27(p0, _1415_v32, _1377_v0, new BigNumber((p1).length), _1416_v33);
      let _nw198 = Array((new BigNumber(28)).toNumber());
      _nw198[(_dafny.ZERO).toNumber()] = _1415_v32;
      _nw198[(_dafny.ONE).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(2)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(3)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(4)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(5)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(6)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((p1).length))];
      _nw198[(new BigNumber(7)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(8)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(9)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(10)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(11)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(12)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(13)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(14)).toNumber()] = (_1417_v34).dtor_cf52;
      _nw198[(new BigNumber(15)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(16)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(17)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(18)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(19)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(20)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(21)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(22)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
      _nw198[(new BigNumber(23)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(24)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(25)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(26)).toNumber()] = _1415_v32;
      _nw198[(new BigNumber(27)).toNumber()] = _1415_v32;
      _1413_v31 = _nw198;
      let _1418_v35;
      _1418_v35 = _dafny.Seq.of(p0, p0);
      (globalState).f2 = _1418_v35;
      r0 = _1377_v0;
      return r0;
    }
    m15(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1419_v0;
      _1419_v0 = _dafny.MultiSet.fromElements((_this).f13);
      let _1420_v1;
      _1420_v1 = _dafny.Seq.of(_1419_v0);
      let _1421_v2;
      _1421_v2 = new BigNumber(-765);
      let _1422_v3;
      _1422_v3 = (_1420_v1)[_module.__default.safeIndex(_1421_v2, new BigNumber((_1420_v1).length))];
      let _source22 = _1422_v3;
      let _1423___mcc_h0 = _source22;
      let _1424_cf0 = _1423___mcc_h0;
      let _1425_v4;
      _1425_v4 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1426_v5;
      _1426_v5 = _module.D3.create_DC5(_1425_v4);
      let _source23 = _1426_v5;
      if (_source23.is_DC6) {
        let _1427___mcc_h1 = (_source23).cf12;
        let _1428___mcc_h2 = (_source23).cf13;
        let _1429_cf13 = _1428___mcc_h2;
        let _1430_cf12 = _1427___mcc_h1;
        let _1431_v7;
        _1431_v7 = _dafny.Seq.UnicodeFromString("cdytr");
        let _1432_v8;
        _1432_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v7,_1425_v4);
        _1421_v2 = (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of ((_1432_v8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1431_v7,_1425_v4))).Keys.Elements) {
            let _1433_v6 = _compr_45;
            if (((_1432_v8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1431_v7,_1425_v4))).contains(_1433_v6)) {
              _coll45.push([_1433_v6,(new BigNumber((_dafny.Seq.of(_this.f24, (_this).f13)).length)).multipliedBy(_1421_v2)]);
            }
          }
          return _coll45;
        }()).length));
        let _1434_v9;
        let _init33 = function (_1435_i0) {
          return (_this).f13;
        };
        let _nw199 = Array((new BigNumber(7)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw199.length); _i0_33++) {
          _nw199[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1434_v9 = _nw199;
        _1434_v9 = ((_this.f24) ? (_1434_v9) : ((_this).f14));
        let _1436_v10;
        _1436_v10 = _dafny.Seq.of(_1429_cf13);
        let _1437_v11;
        _1437_v11 = _dafny.Set.fromElements(_1436_v10);
        let _index233 = _module.__default.safeIndex(new BigNumber(319), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index233] = (_this).fm52((_dafny.ZERO).minus((_this).fm53(_1430_cf12, globalState)), _1437_v11, globalState);
        let _1438_v12;
        _1438_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
        let _1439_v13;
        _1439_v13 = _dafny.MultiSet.fromElements(_1431_v7);
        let _1440_v14;
        _1440_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1430_cf12,_1429_cf13);
        let _1441_v16;
        _1441_v16 = _dafny.Set.fromElements(function () {
          let _coll46 = new _dafny.Set();
          for (const _compr_46 of _dafny.IntegerRange(new BigNumber(996), new BigNumber(-82))) {
            let _1442_v15 = _compr_46;
            if (((new BigNumber(996)).isLessThanOrEqualTo(_1442_v15)) && ((_1442_v15).isLessThan(new BigNumber(-82)))) {
              _coll46.add((_1442_v15).multipliedBy(_1421_v2));
            }
          }
          return _coll46;
        }());
        let _index234 = _module.__default.safeIndex(new BigNumber(319), new BigNumber(((_this).f14).length));
        let _rhs200 = !((((_1438_v12).contains((_this).fm52(_1421_v2, _1437_v11, globalState))) ? ((_1438_v12).get((_this).fm52(_1421_v2, _1437_v11, globalState))) : ((_this).f13)));
        let _rhs201 = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kv"), _dafny.Seq.UnicodeFromString("nolgwiuyb")), _1425_v4);
        let _rhs202 = (((_1439_v13).contains(_dafny.Seq.Concat(_1431_v7, _module.__default.fm60(_this.f24, (_this).f13, _1440_v14, new BigNumber((_1441_v16).length), globalState)))) ? ((_1439_v13).get(_dafny.Seq.Concat(_1431_v7, _module.__default.fm60(_this.f24, (_this).f13, _1440_v14, new BigNumber((_1441_v16).length), globalState)))) : (_1421_v2));
        let _lhs167 = (_this).f14;
        let _lhs168 = _module.__default.safeIndex(new BigNumber(319), new BigNumber(((_this).f14).length));
        let _lhs169 = globalState;
        _1429_cf13 = _rhs200;
        _lhs167[_lhs168] = _rhs201;
        _lhs169.f7 = _rhs202;
        let _1443_v17;
        let _nw200 = new _module.C4();
        _nw200.__ctor();
        _1443_v17 = _nw200;
        let _1444_v18;
        _1444_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v17,_1421_v2);
        let _1445_v19;
        _1445_v19 = _dafny.Seq.of(new BigNumber((_1444_v18).length));
        let _1446_v20;
        _1446_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,(_1445_v19)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((_1445_v19).length))]);
        _1446_v20 = _1446_v20;
      } else if (_source23.is_DC7) {
        let _1447_v21;
        _1447_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(937),(_dafny.ZERO).minus(new BigNumber(-550)));
        let _1448_v22;
        _1448_v22 = _dafny.Set.fromElements(_module.__default.fm61(_this.f24, new BigNumber((_1447_v21).length), _dafny.MultiSet.fromElements(_1425_v4, _1425_v4), globalState));
        let _1449_v23;
        _1449_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm52((_dafny.ZERO).minus(_1421_v2), _1448_v22, globalState),(_this).f14);
        let _1450_v24;
        let _nw201 = new _module.C2();
        _nw201.__ctor(_1421_v2, ((true) ? (_this.f24) : (_this.f24)), _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("uy")), (_this).f13, (((_1449_v23).contains((_this).f13)) ? ((_1449_v23).get((_this).f13)) : ((_this).f14)), (_this).f15, _this.f12, _this.f24);
        _1450_v24 = _nw201;
        let _1451_v25;
        let _nw202 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _1451_v25 = _nw202;
        let _1452_v26;
        _1452_v26 = _module.D8.create_DC18(_1451_v25);
        _1452_v26 = _1452_v26;
        let _index235 = _module.__default.safeIndex(new BigNumber(101), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index235] = (_this).f13;
        let _index236 = _module.__default.safeIndex(new BigNumber(101), new BigNumber(((_this).f14).length));
        let _rhs203 = !(_this.f24);
        let _rhs204 = _this.f24;
        let _rhs205 = false;
        let _lhs170 = (_this).f14;
        let _lhs171 = _module.__default.safeIndex(new BigNumber(101), new BigNumber(((_this).f14).length));
        let _lhs172 = _this;
        let _lhs173 = _this;
        _lhs170[_lhs171] = _rhs203;
        _lhs172.f24 = _rhs204;
        _lhs173.f24 = _rhs205;
        let _1453_v27;
        _1453_v27 = _dafny.MultiSet.fromElements(_1421_v2, _1421_v2);
        (globalState).f7 = ((((_1453_v27).contains((_1450_v24).f21)) ? ((_1453_v27).get((_1450_v24).f21)) : (_1421_v2))).multipliedBy(_1421_v2);
      } else if (_source23.is_DC8) {
        let _1454___mcc_h3 = (_source23).cf14;
        let _1455_cf14 = _1454___mcc_h3;
        let _1456_v28;
        _1456_v28 = _dafny.Seq.of(_this.f24, _this.f24);
        let _1457_v29;
        _1457_v29 = _module.D15.create_DC38(_1456_v28);
        let _1458_v30;
        _1458_v30 = _dafny.Set.fromElements((_1457_v29).dtor_cf77, _dafny.Seq.of(_this.f24));
        let _1459_v31;
        _1459_v31 = _dafny.Seq.of(_1458_v30, _dafny.Set.fromElements(_1456_v28, _1456_v28, _1456_v28));
        let _1460_v32;
        _1460_v32 = _dafny.Set.fromElements(_this.f24);
        let _1461_v33;
        _1461_v33 = _dafny.Seq.of(_1460_v32);
        (_this).f24 = !((_this).fm52(_1421_v2, (_1459_v31)[_module.__default.safeIndex(new BigNumber(((_1461_v33)[_module.__default.safeIndex(_1421_v2, new BigNumber((_1461_v33).length))]).length), new BigNumber((_1459_v31).length))], globalState));
        let _1462_v34;
        _1462_v34 = _dafny.Seq.of(_1421_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true, (_this).f13), _1456_v28)).length), _1455_cf14);
        (globalState).f7 = (_1462_v34)[_module.__default.safeIndex((_dafny.ZERO).minus(_1455_cf14), new BigNumber((_1462_v34).length))];
        _1421_v2 = _1455_cf14;
        let _rhs206 = (((_this).f15).IsProperSubsetOf((_this).f15)) || ((_this).f13);
        let _rhs207 = _module.__default.safeModuloInt(_1421_v2, _1455_cf14);
        let _rhs208 = _1456_v28;
        let _lhs174 = _this;
        let _lhs175 = globalState;
        let _lhs176 = globalState;
        _lhs174.f24 = _rhs206;
        _lhs175.f7 = _rhs207;
        _lhs176.f2 = _rhs208;
      } else {
        let _1463___mcc_h4 = (_source23).cf11;
        let _1464_cf11 = _1463___mcc_h4;
        let _1465_v35;
        _1465_v35 = _dafny.Seq.UnicodeFromString("dffjrx");
        let _1466_v36;
        _1466_v36 = _dafny.Seq.of(_1421_v2);
        let _1467_v37;
        _1467_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1466_v36)).cardinality()),_1425_v4);
        let _1468_v38;
        _1468_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1467_v37).length),(_this).f13);
        let _1469_v39;
        _1469_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_1421_v2);
        let _1470_v40;
        _1470_v40 = _module.D6.create_DC13(_this.f24, _this.f24, _1421_v2, new BigNumber((_dafny.Seq.of((_this).f13)).length), !((_this).f13));
        let _1471_v41;
        let _nw203 = new _module.C3();
        _nw203.__ctor((_this).f13, _module.__default.fm62(_this.f24, _1465_v35, globalState), (_this).f13, (_this).f14, (_this).f15, _this.f12, _this.f24);
        _1471_v41 = _nw203;
        let _1472_v42;
        _1472_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1471_v41,(_this).f13);
        let _1473_v43;
        _1473_v43 = _dafny.MultiSet.fromElements(_1421_v2);
        let _1474_v44;
        _1474_v44 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.of((_1471_v41).f13, true, (_this).f13, (_this).f13), _module.__default.safeIndex(new BigNumber((_1473_v43).cardinality()), new BigNumber((_dafny.Seq.of((_1471_v41).f13, true, (_this).f13, (_this).f13)).length)), (_this).f13));
        let _1475_v45;
        _1475_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).fm52(_1421_v2, _1474_v44, globalState));
        let _1476_v46;
        let _nw204 = Array((new BigNumber(28)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = (new BigNumber((_1465_v35).length)).plus(_1421_v2);
        _nw204[(_dafny.ONE).toNumber()] = new BigNumber(((_1468_v38).Merge(_1468_v38)).length);
        _nw204[(new BigNumber(2)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(3)).toNumber()] = (((_1469_v39).contains(_this.f24)) ? ((_1469_v39).get(_this.f24)) : (_1421_v2));
        _nw204[(new BigNumber(4)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(5)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(6)).toNumber()] = (new BigNumber(428)).plus((_this).fm8((_this).f13, _module.D1.create_DC1(_1465_v35), globalState));
        _nw204[(new BigNumber(7)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(8)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(9)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(10)).toNumber()] = ((_dafny.ZERO).minus((_1470_v40).dtor_cf25)).plus(_1421_v2);
        _nw204[(new BigNumber(11)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_1421_v2).plus(_1421_v2));
        _nw204[(new BigNumber(13)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(14)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(15)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(16)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_1421_v2)).length);
        _nw204[(new BigNumber(18)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(19)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(20)).toNumber()] = (((_this).f13) ? (_1421_v2) : (_1421_v2));
        _nw204[(new BigNumber(21)).toNumber()] = new BigNumber((_1472_v42).length);
        _nw204[(new BigNumber(22)).toNumber()] = new BigNumber(((_1475_v45).Merge(_1475_v45)).length);
        _nw204[(new BigNumber(23)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(24)).toNumber()] = _1421_v2;
        _nw204[(new BigNumber(25)).toNumber()] = (new BigNumber(105)).minus(_1421_v2);
        _nw204[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus(_1421_v2);
        _nw204[(new BigNumber(27)).toNumber()] = _1421_v2;
        _1476_v46 = _nw204;
        let _index237 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1476_v46).length));
        (_1476_v46)[_index237] = _1421_v2;
        let _index238 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1476_v46).length));
        (_1476_v46)[_index238] = (new BigNumber((_1424_cf0).cardinality())).plus(_1421_v2);
        _1473_v43 = _1473_v43;
        let _index239 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1476_v46).length));
        (_1476_v46)[_index239] = (_dafny.ZERO).minus((_1476_v46)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_1476_v46).length))]);
        let _1477_v47;
        _1477_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1476_v46)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_1476_v46).length))],(_dafny.ZERO).minus(_1421_v2));
        let _1478_v48;
        _1478_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1464_cf11,(_1466_v36)[_module.__default.safeIndex(new BigNumber((_1477_v47).length), new BigNumber((_1466_v36).length))]);
        _1478_v48 = (_1478_v48).update(_1425_v4, _1421_v2);
      }
      let _index240 = _module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index240] = _this.f24;
      let _1479_v49;
      _1479_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1421_v2);
      let _1480_v50;
      _1480_v50 = _module.D6.create_DC13((_this).f13, (_this).f13, _1421_v2, new BigNumber((_1479_v49).length), !((_this).f13));
      let _1481_v51;
      _1481_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f24);
      let _1482_v52;
      _1482_v52 = _dafny.Seq.UnicodeFromString("hmg");
      let _pat_let_tv36 = _1482_v52;
      let _pat_let_tv37 = _1481_v51;
      let _pat_let_tv38 = _1481_v51;
      let _1483_v53;
      _1483_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1421_v2,(function (_pat_let39_0) {
        return function (_1484_dt__update__tmp_h0) {
          return function (_pat_let40_0) {
            return function (_1485_dt__update_hcf24_h0) {
              return function (_pat_let41_0) {
                return function (_1486_dt__update_hcf22_h0) {
                  return _module.D6.create_DC13(_1486_dt__update_hcf22_h0, (_1484_dt__update__tmp_h0).dtor_cf23, _1485_dt__update_hcf24_h0, (_1484_dt__update__tmp_h0).dtor_cf25, (_1484_dt__update__tmp_h0).dtor_cf26);
                }(_pat_let41_0);
              }((((_pat_let_tv38).contains(_this.f24)) ? ((_pat_let_tv37).get(_this.f24)) : (!(!((_this).f13)))));
            }(_pat_let40_0);
          }(new BigNumber((_pat_let_tv36).length));
        }(_pat_let39_0);
      }(_1480_v50)).dtor_cf26);
      let _index241 = _module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index241] = (((_1483_v53).contains((_1421_v2).minus(_1421_v2))) ? ((_1483_v53).get((_1421_v2).minus(_1421_v2))) : (!(false)));
      (_this).f24 = (((_this).f13) ? (!((_this).f13)) : (_this.f24));
      let _1487_v54;
      _1487_v54 = _dafny.Seq.of(((_this).f14)[_module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length))]);
      let _1488_v55;
      _1488_v55 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_1482_v52, _module.__default.safeIndex(_1421_v2, new BigNumber((_1482_v52).length)), _1425_v4)).length), _1421_v2);
      if ((_1488_v55).IsProperSubsetOf((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), ((_1489_v2) => function (_1490_i1) {
        return _1489_v2;
      })(_1421_v2)))).Intersect(_module.__default.fm56(new BigNumber((_1487_v54).length), (_this).f15, globalState)))) {
        let _1491_v56;
        _1491_v56 = _dafny.Seq.of(_1421_v2, new BigNumber(703), _1421_v2, _1421_v2, _1421_v2);
        _1491_v56 = _1491_v56;
        let _1492_v57;
        let _init34 = ((_1493_v2) => function (_1494_i2) {
          return (_1494_i2).plus(_1493_v2);
        })(_1421_v2);
        let _nw205 = Array((new BigNumber(4)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw205.length); _i0_34++) {
          _nw205[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1492_v57 = _nw205;
        let _index242 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1492_v57).length));
        (_1492_v57)[_index242] = (_1421_v2).minus(new BigNumber((_1479_v49).length));
        let _index243 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1492_v57).length));
        (_1492_v57)[_index243] = (_1421_v2).multipliedBy(_1421_v2);
        (globalState).f7 = new BigNumber(((_this).f15).length);
        _1482_v52 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(483)), function (_1495_i3) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        });
        let _index244 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_1492_v57).length));
        (_1492_v57)[_index244] = (_1492_v57)[_module.__default.safeIndex(new BigNumber(548), new BigNumber((_1492_v57).length))];
      } else {
        let _index245 = _module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index245] = ((_this).f14)[_module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length))];
        let _1496_v58;
        _1496_v58 = _dafny.Seq.of(_1421_v2);
        let _index246 = _module.__default.safeIndex(new BigNumber(234), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index246] = !((new BigNumber((_1496_v58).length)).isLessThanOrEqualTo(_1421_v2));
        r0 = _1421_v2;
        (globalState).f7 = _1421_v2;
        let _1497_v59;
        _1497_v59 = _module.D7.create_DC16(_1421_v2);
        _1497_v59 = _1497_v59;
      }
      let _1498_v60;
      _1498_v60 = _dafny.Seq.UnicodeFromString("jnxuox");
      let _1499_v61;
      let _nw206 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1499_v61 = _nw206;
      let _1500_v62;
      _1500_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1421_v2, new BigNumber((_1498_v60).length)),_1499_v61);
      let _1501_v63;
      _1501_v63 = _dafny.Map.Empty.slice().updateUnsafe(false,_1421_v2);
      let _1502_v64;
      _1502_v64 = _module.D1.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_1503_i4) {
  return new _dafny.CodePoint('d'.codePointAt(0));
}));
      let _1504_v65;
      let _nw207 = new _module.C2();
      _nw207.__ctor(new BigNumber((_1501_v63).length), _this.f24, _1502_v64, !(true), (_this).f14, (_this).f15, _this.f12, _this.f24);
      _1504_v65 = _nw207;
      let _1505_v66;
      _1505_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v65,_dafny.Seq.of((_this).f13));
      let _1506_v67;
      _1506_v67 = _dafny.Seq.of(_this.f24);
      let _1507_v68;
      _1507_v68 = _dafny.Set.fromElements((((_1505_v66).contains(_1504_v65)) ? ((_1505_v66).get(_1504_v65)) : (_1506_v67)));
      let _1508_v69;
      _1508_v69 = _module.D6.create_DC14(_module.D6.create_DC13((_this).f13, (_this).fm52(_1421_v2, _1507_v68, globalState), new BigNumber((_dafny.MultiSet.FromArray(_1506_v67)).cardinality()), _1421_v2, (_1504_v65).f22));
      let _1509_v70;
      _1509_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1421_v2,_1508_v69);
      _1500_v62 = (_1500_v62).update(new BigNumber((_1509_v70).length), _1499_v61);
      let _hi9 = _1421_v2;
      for (let _1510_i5 = _1421_v2; _1510_i5.isLessThan(_hi9); _1510_i5 = _1510_i5.plus(_dafny.ONE)) {
        let _index247 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1499_v61).length));
        (_1499_v61)[_index247] = _1510_i5;
        let _1511_v71;
        _1511_v71 = new _dafny.CodePoint('t'.codePointAt(0));
        let _index248 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1499_v61).length));
        let _rhs209 = !(_dafny.Seq.IsProperPrefixOf(_1498_v60, _dafny.Seq.update(_1498_v60, _module.__default.safeIndex((_1504_v65).f21, new BigNumber((_1498_v60).length)), _1511_v71))) || ((_this).fm52(_1421_v2, _1507_v68, globalState));
        let _rhs210 = _1499_v61;
        let _rhs211 = _1498_v60;
        let _rhs212 = (_1504_v65).f22;
        let _rhs213 = (_1504_v65).f21;
        let _lhs177 = _this;
        let _lhs178 = _this;
        let _lhs179 = _1499_v61;
        let _lhs180 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_1499_v61).length));
        _lhs177.f24 = _rhs209;
        _1499_v61 = _rhs210;
        _1498_v60 = _rhs211;
        _lhs178.f24 = _rhs212;
        _lhs179[_lhs180] = _rhs213;
        let _1512_v72;
        let _nw208 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1512_v72 = _nw208;
        let _1513_v73;
        _1513_v73 = _module.D11.create_DC26(_1512_v72, _1498_v60);
        let _1514_v74;
        let _nw209 = Array((new BigNumber(8)).toNumber());
        _nw209[(_dafny.ZERO).toNumber()] = _1512_v72;
        _nw209[(_dafny.ONE).toNumber()] = _1512_v72;
        _nw209[(new BigNumber(2)).toNumber()] = _1512_v72;
        _nw209[(new BigNumber(3)).toNumber()] = _1512_v72;
        _nw209[(new BigNumber(4)).toNumber()] = _1512_v72;
        _nw209[(new BigNumber(5)).toNumber()] = (_1513_v73).dtor_cf49;
        _nw209[(new BigNumber(6)).toNumber()] = _1512_v72;
        _nw209[(new BigNumber(7)).toNumber()] = _1512_v72;
        _1514_v74 = _nw209;
        let _index249 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1514_v74).length));
        (_1514_v74)[_index249] = _1512_v72;
        let _1515_v75;
        let _nw210 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1515_v75 = _nw210;
        let _1516_v76;
        _1516_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1504_v65).f21,new BigNumber(501));
        let _index250 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1514_v74).length));
        let _rhs214 = _this.f24;
        let _rhs215 = _1512_v72;
        let _rhs216 = ((((_dafny.ZERO).minus(new BigNumber((_1516_v76).length))).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,true)).length))) ? (_1515_v75) : (_1515_v75));
        let _lhs181 = _this;
        let _lhs182 = _1514_v74;
        let _lhs183 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1514_v74).length));
        _lhs181.f24 = _rhs214;
        _lhs182[_lhs183] = _rhs215;
        _1515_v75 = _rhs216;
        let _1517_v77;
        _1517_v77 = _module.D2.create_DC4((_1514_v74)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_1514_v74).length))], (_this).f13, new BigNumber((_dafny.Seq.UnicodeFromString("iqftrwd")).length), true, (_this).f14);
        let _1518_v78;
        _1518_v78 = _dafny.Map.Empty.slice().updateUnsafe(((!(_this.f24)) ? (_1511_v71) : (_1511_v71)),(_1517_v77).dtor_cf7);
        _1518_v78 = (_1518_v78).update((((_this).f13) ? (_1511_v71) : (_1511_v71)), (_1504_v65).f22);
        (globalState).f7 = _1510_i5;
      }
      let _1519_i6;
      _1519_i6 = _dafny.ZERO;
      L7: {
        while (!((_1504_v65).f22)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1519_i6)) {
              break L7;
            }
            _1519_i6 = (_1519_i6).plus(_dafny.ONE);
            if (true) {
              let _index251 = _module.__default.safeIndex(new BigNumber(866), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index251] = (_1504_v65).f22;
              let _1520_v79;
              _1520_v79 = _dafny.Set.fromElements(_1499_v61, _1499_v61, _1499_v61);
              let _index252 = _module.__default.safeIndex(new BigNumber(866), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index252] = (_dafny.Set.fromElements(_1499_v61)).equals(_1520_v79);
              let _1521_v80;
              _1521_v80 = new _dafny.CodePoint('x'.codePointAt(0));
              let _1522_v81;
              let _nw211 = new _module.C0();
              _nw211.__ctor(_this.f24, _1521_v80, _this.f12, (_1504_v65).f22);
              _1522_v81 = _nw211;
              let _1523_v82;
              _1523_v82 = _module.D16.create_DC40(_1522_v81);
              _1522_v81 = (_1523_v82).dtor_cf83;
              (globalState).f7 = (_1504_v65).f21;
              let _1524_v83;
              let _nw212 = Array((new BigNumber(16)).toNumber()).fill([]);
              _1524_v83 = _nw212;
              let _1525_v84;
              let _nw213 = Array((new BigNumber(20)).toNumber()).fill(false);
              _1525_v84 = _nw213;
              let _index253 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1524_v83).length));
              (_1524_v83)[_index253] = _1525_v84;
              let _index254 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1524_v83).length));
              (_1524_v83)[_index254] = _1525_v84;
              _1509_v70 = (_1509_v70).update((_1504_v65).f21, _1508_v69);
            } else {
              let _1526_v85;
              _1526_v85 = _dafny.MultiSet.fromElements((_this).f14, (_this).f14);
              (globalState).f7 = _module.__default.safeModuloInt(new BigNumber(822), (((_1526_v85).contains((_this).f14)) ? ((_1526_v85).get((_this).f14)) : ((_1504_v65).f21)));
              let _1527_v86;
              _1527_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1498_v60,(_this).f13);
              let _rhs217 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), function (_1528_i7) {
                return new _dafny.CodePoint('i'.codePointAt(0));
              })).length)).minus(new BigNumber((_1527_v86).length));
              let _rhs218 = (((_1504_v65).f21).plus(_1421_v2)).minus((_dafny.ZERO).minus((_1504_v65).f21));
              let _rhs219 = (_this).f13;
              let _rhs220 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_1504_v65).f21, _1421_v2));
              let _rhs221 = false;
              let _lhs184 = globalState;
              let _lhs185 = _this;
              let _lhs186 = _this;
              r0 = _rhs217;
              _lhs184.f7 = _rhs218;
              _lhs185.f24 = _rhs219;
              r0 = _rhs220;
              _lhs186.f24 = _rhs221;
              (globalState).f7 = new BigNumber((((!((_1504_v65).f22)) ? (_1509_v70) : (((_this.f24) ? (_1509_v70) : (function () {
                let _coll47 = new _dafny.Map();
                for (const _compr_47 of _dafny.IntegerRange(new BigNumber(155), new BigNumber(-903))) {
                  let _1529_v87 = _compr_47;
                  if (((new BigNumber(155)).isLessThanOrEqualTo(_1529_v87)) && ((_1529_v87).isLessThan(new BigNumber(-903)))) {
                    _coll47.push([_module.__default.safeDivisionInt(_1529_v87, _1421_v2),_1508_v69]);
                  }
                }
                return _coll47;
              }()))))).length);
              let _index255 = _module.__default.safeIndex(new BigNumber(941), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index255] = _this.f24;
              let _1530_v88;
              _1530_v88 = _dafny.Seq.of(_1498_v60, _1498_v60, _1498_v60);
              let _index256 = _module.__default.safeIndex(new BigNumber(941), new BigNumber(((_this).f14).length));
              let _rhs222 = (_1530_v88)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1421_v2, (_1504_v65).f21), new BigNumber((_1530_v88).length))];
              let _rhs223 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_1504_v65).f21, (_1504_v65).f21))).isLessThanOrEqualTo((_1504_v65).f21);
              let _lhs187 = (_this).f14;
              let _lhs188 = _module.__default.safeIndex(new BigNumber(941), new BigNumber(((_this).f14).length));
              _1498_v60 = _rhs222;
              _lhs187[_lhs188] = _rhs223;
              (_this).f24 = ((_1504_v65).f21).isEqualTo((_1504_v65).f21);
            }
            (_this).f24 = (_this).fm52((_1504_v65).f21, _1507_v68, globalState);
            let _1531_v89;
            _1531_v89 = _dafny.Seq.of(_1421_v2, _1421_v2, (_1504_v65).f21, (_1504_v65).f21, (_1504_v65).f21);
            let _1532_v90;
            _1532_v90 = _dafny.MultiSet.fromElements(_1531_v89);
            if ((((_this.f24) ? (_dafny.MultiSet.fromElements(_1531_v89)) : (_1532_v90))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(99)), ((_1533_v65) => function (_1534_i8) {
              return (_1533_v65).f21;
            })(_1504_v65)), _module.__default.fm63((_1504_v65).f22, (_1504_v65).f22, globalState)))) {
              let _1535_v91;
              _1535_v91 = _dafny.Map.Empty.slice().updateUnsafe((_1504_v65).f22,(_this).f13);
              _1535_v91 = (_1535_v91).update((_1504_v65).f22, (_1504_v65).f22);
              let _1536_v92;
              _1536_v92 = new _dafny.CodePoint('t'.codePointAt(0));
              let _1537_v93;
              let _nw214 = new _module.C1();
              _nw214.__ctor(_1536_v92, _module.D1.create_DC2((_1504_v65).f21, (_1504_v65).f21, (_1504_v65).f22), (_this).f13);
              _1537_v93 = _nw214;
              let _1538_v94;
              _1538_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v93,(_1504_v65).f22);
              let _1539_v95;
              _1539_v95 = _dafny.Map.Empty.slice().updateUnsafe((_1538_v94).Merge(_1538_v94),(_1504_v65).f21);
              let _rhs224 = (_1537_v93).fm43(globalState);
              let _rhs225 = (_1539_v95).Merge(_1539_v95);
              let _lhs189 = _this;
              _lhs189.f24 = _rhs224;
              _1539_v95 = _rhs225;
              (globalState).f7 = (_1504_v65).f21;
              (_this).f24 = (_this.f24) || ((_this).f13);
              let _1540_v96;
              _1540_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(502),(_1504_v65).f21);
              let _1541_v97;
              _1541_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1498_v60,_1540_v96);
              _1541_v97 = (_1541_v97).update(_1498_v60, _1540_v96);
            } else {
              let _1542_v98;
              _1542_v98 = _dafny.MultiSet.fromElements(_1506_v67, _1506_v67);
              let _index257 = _module.__default.safeIndex(new BigNumber(476), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index257] = (_1542_v98).equals(_dafny.MultiSet.fromElements(_1506_v67));
              let _index258 = _module.__default.safeIndex(new BigNumber(476), new BigNumber(((_this).f14).length));
              ((_this).f14)[_index258] = (_this.f24) && ((_this).fm52(new BigNumber((_dafny.MultiSet.fromElements((_1504_v65).f22)).cardinality()), _1507_v68, globalState));
              r0 = _1421_v2;
              (globalState).f7 = (_1504_v65).f21;
              let _1543_v99;
              _1543_v99 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt((_1504_v65).f21, (_1504_v65).f21));
              _1421_v2 = new BigNumber((_1543_v99).cardinality());
              _1498_v60 = (((_1419_v0).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f13))) ? (_1498_v60) : (_1498_v60));
            }
            let _1544_v100;
            _1544_v100 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f24);
            _1544_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_1419_v0).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f13)));
          }
        }
      }
      (globalState).f7 = (new BigNumber((_module.__default.fm63(true, _this.f24, globalState)).length)).plus(_1421_v2);
      let _1545_v101;
      _1545_v101 = _module.D11.create_DC27(false, new _dafny.CodePoint('q'.codePointAt(0)), new BigNumber(890), _1421_v2, _dafny.Seq.of(new BigNumber(118)));
      let _1546_i9;
      _1546_i9 = _dafny.ZERO;
      L8: {
        while ((!(_1419_v0).equals(_1419_v0)) && ((_1545_v101).dtor_cf51)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1546_i9)) {
              break L8;
            }
            _1546_i9 = (_1546_i9).plus(_dafny.ONE);
            let _1547_v102;
            let _nw215 = Array((new BigNumber(14)).toNumber()).fill(false);
            _1547_v102 = _nw215;
            let _init35 = ((_1548_v101) => function (_1549_i10) {
              return (_1548_v101).dtor_cf51;
            })(_1545_v101);
            let _nw216 = Array((_dafny.ONE).toNumber());
            for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw216.length); _i0_35++) {
              _nw216[_i0_35] = _init35(new BigNumber(_i0_35));
            }
            _1547_v102 = _nw216;
            _1498_v60 = (((_1504_v65).f22) ? (_1498_v60) : (_1498_v60));
            let _1550_v103;
            _1550_v103 = new _dafny.CodePoint('i'.codePointAt(0));
            let _1551_v104;
            let _nw217 = new _module.C0();
            _nw217.__ctor((_1504_v65).f22, _1550_v103, _this.f12, (_1506_v67)[_module.__default.safeIndex(new BigNumber(-991), new BigNumber((_1506_v67).length))]);
            _1551_v104 = _nw217;
            let _1552_v105;
            _1552_v105 = _dafny.Map.Empty.slice().updateUnsafe((_1504_v65).f21,(_dafny.Set.fromElements((_this).fm53(_1421_v2, globalState))).Union((_this).f15));
            _1552_v105 = (_1552_v105).update((_1504_v65).f21, (_module.__default.fm64(globalState)).Difference(_dafny.Set.fromElements(_1421_v2)));
          }
        }
      }
      let _1553_v106;
      _1553_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(134),(_1504_v65).f21);
      r0 = new BigNumber(((_1553_v106).Merge(_1553_v106)).length);
      return r0;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f12 = _module.D1.Default();
      this._f16 = _module.D1.Default();
      this._f14 = [];
      this._f15 = _dafny.Set.Empty;
      this._f13 = false;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f16, f17, f14, f15, f12, f13) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_this).f13, (_this).f17)).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f17)),new BigNumber((_dafny.Seq.of(new BigNumber(705), (_dafny.ZERO).minus(new BigNumber(-805)))).length));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      if ((_this).f17) {
        return new BigNumber(-683);
      } else {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_dafny.Seq.of((_this).f17, (_this).f13))).length))))).cardinality()));
      }
    };
    fm9(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(711)), function (_1554_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length))).multipliedBy(new BigNumber((function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of _dafny.IntegerRange(new BigNumber(-200), new BigNumber(231))) {
          let _1555_v0 = _compr_48;
          if (((new BigNumber(-200)).isLessThanOrEqualTo(_1555_v0)) && ((_1555_v0).isLessThan(new BigNumber(231)))) {
            _coll48.push([_module.__default.safeDivisionInt(_1555_v0, new BigNumber(374)),new BigNumber(805)]);
          }
        }
        return _coll48;
      }()).length)), new BigNumber(((_dafny.Set.fromElements((_this).f13, !((_this).f17))).Intersect(_dafny.Set.fromElements((_this).f13))).length)));
    };
    fm11(globalState) {
      let _this = this;
      return (_this).f13;
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return (_this).f17;
    };
    m6(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1556_v0;
      _1556_v0 = new BigNumber(622);
      let _hi10 = (_1556_v0).minus(_1556_v0);
      for (let _1557_i0 = (_1556_v0).minus(_1556_v0); _1557_i0.isLessThan(_hi10); _1557_i0 = _1557_i0.plus(_dafny.ONE)) {
        let _1558_v1;
        let _nw218 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
        _1558_v1 = _nw218;
        let _1559_v2;
        _1559_v2 = _dafny.Set.fromElements(p0, (_this).f17);
        let _index259 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_1558_v1).length));
        (_1558_v1)[_index259] = (_1559_v2).Union(_1559_v2);
        let _index260 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_1558_v1).length));
        (_1558_v1)[_index260] = _1559_v2;
        let _1560_v3;
        _1560_v3 = false;
        let _1561_v4;
        let _nw219 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1561_v4 = _nw219;
        let _1562_v5;
        _1562_v5 = _dafny.MultiSet.fromElements((_this).f13);
        let _index261 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length));
        (_1561_v4)[_index261] = new BigNumber((_1562_v5).cardinality());
        let _1563_v6;
        _1563_v6 = _dafny.Seq.UnicodeFromString("pjvi");
        let _1564_v7;
        _1564_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1563_v6).length),(_this).f13);
        let _1565_v8;
        _1565_v8 = _dafny.Seq.of(new BigNumber(((_1558_v1)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_1558_v1).length))]).length));
        let _1566_v9;
        _1566_v9 = _dafny.MultiSet.fromElements(_1557_i0);
        let _1567_v10;
        _1567_v10 = _dafny.Seq.of(!((((_1564_v7).contains(_1556_v0)) ? ((_1564_v7).get(_1556_v0)) : (p0))), (_1566_v9).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1565_v8)));
        let _1568_v11;
        _1568_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(_1563_v6, globalState),(_this).f13);
        let _index262 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length));
        let _rhs226 = (_1567_v10)[_module.__default.safeIndex(_1557_i0, new BigNumber((_1567_v10).length))];
        let _rhs227 = new BigNumber((_dafny.Seq.UnicodeFromString("sn")).length);
        let _rhs228 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1556_v0));
        let _rhs229 = _1556_v0;
        let _rhs230 = !(p0) || ((_this).fm12(_1568_v11, _this.f12, globalState));
        let _lhs190 = globalState;
        let _lhs191 = _1561_v4;
        let _lhs192 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length));
        _1560_v3 = _rhs226;
        r0 = _rhs227;
        _lhs190.f7 = _rhs228;
        _lhs191[_lhs192] = _rhs229;
        _1560_v3 = _rhs230;
        let _1569_v12;
        _1569_v12 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1570_v13;
        _1570_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), function (_1571_i1) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), function (_1572_i2) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length);
        })).length),_1569_v12);
        let _1573_v14;
        _1573_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1561_v4)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length))]);
        let _1574_v15;
        _1574_v15 = _dafny.Set.fromElements(_1561_v4);
        let _1575_v16;
        _1575_v16 = _module.D2.create_DC4(_1561_v4, p0, _1557_i0, _1560_v3, (_this).f14);
        let _1576_v17;
        let _nw220 = Array((new BigNumber(17)).toNumber());
        _nw220[(_dafny.ZERO).toNumber()] = _1560_v3;
        _nw220[(_dafny.ONE).toNumber()] = _1560_v3;
        _nw220[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_1563_v6, _1563_v6);
        _nw220[(new BigNumber(3)).toNumber()] = p0;
        _nw220[(new BigNumber(4)).toNumber()] = (_1560_v3) || (_1560_v3);
        _nw220[(new BigNumber(5)).toNumber()] = _1560_v3;
        _nw220[(new BigNumber(6)).toNumber()] = _1560_v3;
        _nw220[(new BigNumber(7)).toNumber()] = p0;
        _nw220[(new BigNumber(8)).toNumber()] = (false) && (_1560_v3);
        _nw220[(new BigNumber(9)).toNumber()] = (_this).f17;
        _nw220[(new BigNumber(10)).toNumber()] = (_this).f13;
        _nw220[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_1570_v13).length), (_dafny.ZERO).minus((_1565_v8)[_module.__default.safeIndex((_1561_v4)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length))], new BigNumber((_1565_v8).length))]))).IsDisjointFrom(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1557_i0)));
        _nw220[(new BigNumber(12)).toNumber()] = _1560_v3;
        _nw220[(new BigNumber(13)).toNumber()] = (_1556_v0).isLessThanOrEqualTo((((_1573_v14).contains((_this).f17)) ? ((_1573_v14).get((_this).f17)) : (_module.__default.fm0(_dafny.Seq.update(_1567_v10, _module.__default.safeIndex((_1561_v4)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length))], new BigNumber((_1567_v10).length)), (_this).f17), (_1567_v10)[_module.__default.safeIndex((_1561_v4)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length))], new BigNumber((_1567_v10).length))], (_this).f13, new BigNumber((_1574_v15).length), globalState))));
        _nw220[(new BigNumber(14)).toNumber()] = (_this).f13;
        _nw220[(new BigNumber(15)).toNumber()] = !((_1567_v10)[_module.__default.safeIndex(new BigNumber(980), new BigNumber((_1567_v10).length))]);
        _nw220[(new BigNumber(16)).toNumber()] = (_1575_v16).dtor_cf9;
        _1576_v17 = _nw220;
        _1576_v17 = (_this).f14;
        let _index263 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length));
        (_1561_v4)[_index263] = (_1561_v4)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_1561_v4).length))];
      }
      let _1577_v18;
      _1577_v18 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _1578_v19;
      _1578_v19 = _dafny.Seq.of((((_1577_v18).contains(p0)) ? ((_1577_v18).get(p0)) : (!(false))), (_this).f13);
      if (_dafny.Seq.IsPrefixOf(_1578_v19, _dafny.Seq.Concat(_1578_v19, _dafny.Seq.of((_this).f17)))) {
        let _1579_v20;
        _1579_v20 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1580_v21;
        _1580_v21 = _dafny.Seq.of(_1556_v0, _1556_v0, _1556_v0, _1556_v0);
        let _1581_v22;
        _1581_v22 = _dafny.Seq.of((_1580_v21)[_module.__default.safeIndex(_1556_v0, new BigNumber((_1580_v21).length))], _1556_v0);
        let _1582_v23;
        let _nw221 = new _module.C1();
        _nw221.__ctor(_1579_v20, _this.f12, !_dafny.areEqual(_1581_v22, _1580_v21));
        _1582_v23 = _nw221;
        _1556_v0 = (_1556_v0).plus(_1556_v0);
        let _index264 = _module.__default.safeIndex(new BigNumber(608), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index264] = (_this).f17;
        let _1583_v24;
        _1583_v24 = _dafny.Seq.UnicodeFromString("mttwbblrp");
        let _1584_v25;
        _1584_v25 = _dafny.Map.Empty.slice().updateUnsafe(true,_1556_v0);
        let _index265 = _module.__default.safeIndex(new BigNumber(608), new BigNumber(((_this).f14).length));
        let _rhs231 = _module.__default.fm0(_dafny.Seq.Concat(_1578_v19, _1578_v19), !((_this).f17), (_this).f13, _1556_v0, globalState);
        let _rhs232 = (_this).f13;
        let _rhs233 = _module.__default.fm18((_dafny.ZERO).minus((_dafny.ZERO).minus(_1556_v0)), _1583_v24, (_this).f17, (((_this).f13) ? (_1584_v25) : ((_this).fm10((_this).f17, (_this).f13, _1578_v19, globalState))), globalState);
        let _rhs234 = _1556_v0;
        let _lhs193 = globalState;
        let _lhs194 = (_this).f14;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(608), new BigNumber(((_this).f14).length));
        let _lhs196 = _1582_v23;
        _lhs193.f7 = _rhs231;
        _lhs194[_lhs195] = _rhs232;
        _lhs196.f23 = _rhs233;
        r0 = _rhs234;
        let _1585_v26;
        let _nw222 = Array((new BigNumber(4)).toNumber());
        _nw222[(_dafny.ZERO).toNumber()] = _1556_v0;
        _nw222[(_dafny.ONE).toNumber()] = _1556_v0;
        _nw222[(new BigNumber(2)).toNumber()] = _1556_v0;
        _nw222[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_1556_v0);
        _1585_v26 = _nw222;
        let _1586_v27;
        let _nw223 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1586_v27 = _nw223;
        let _1587_v28;
        _1587_v28 = _module.D2.create_DC4(_1585_v26, (_1578_v19)[_module.__default.safeIndex(_1556_v0, new BigNumber((_1578_v19).length))], _1556_v0, (_this).f13, _1586_v27);
        let _pat_let_tv39 = _1586_v27;
        let _pat_let_tv40 = _1556_v0;
        let _pat_let_tv41 = _1585_v26;
        _1584_v25 = (_1584_v25).update((function (_pat_let42_0) {
          return function (_1588_dt__update__tmp_h0) {
            return function (_pat_let43_0) {
              return function (_1589_dt__update_hcf10_h0) {
                return function (_pat_let44_0) {
                  return function (_1590_dt__update_hcf7_h0) {
                    return function (_pat_let45_0) {
                      return function (_1591_dt__update_hcf8_h0) {
                        return function (_pat_let46_0) {
                          return function (_1592_dt__update_hcf6_h0) {
                            return _module.D2.create_DC4(_1592_dt__update_hcf6_h0, _1590_dt__update_hcf7_h0, _1591_dt__update_hcf8_h0, (_1588_dt__update__tmp_h0).dtor_cf9, _1589_dt__update_hcf10_h0);
                          }(_pat_let46_0);
                        }(_pat_let_tv41);
                      }(_pat_let45_0);
                    }(_pat_let_tv40);
                  }(_pat_let44_0);
                }(((_this).f14)[_module.__default.safeIndex(new BigNumber(608), new BigNumber(((_this).f14).length))]);
              }(_pat_let43_0);
            }(_pat_let_tv39);
          }(_pat_let42_0);
        }(_1587_v28)).dtor_cf9, _module.__default.safeDivisionInt(_1556_v0, _1556_v0));
        (globalState).f7 = (_1556_v0).plus(_1556_v0);
      } else {
        let _1593_v29;
        let _nw224 = new _module.C4();
        _nw224.__ctor();
        _1593_v29 = _nw224;
        let _nw225 = new _module.C4();
        _nw225.__ctor();
        _1593_v29 = _nw225;
        let _1594_v30;
        _1594_v30 = _dafny.Seq.of((_1556_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tfumdn")).length)), _1556_v0);
        _1594_v30 = _dafny.Seq.Concat(_1594_v30, _1594_v30);
        (globalState).f7 = _1556_v0;
        r0 = (_1556_v0).minus((_this.f12).dtor_cf3);
        let _1595_v31;
        _1595_v31 = _module.D2.create_DC3((_this).f14);
        let _1596_v32;
        let _init36 = function (_1597_i3) {
          return (_1597_i3).plus(new BigNumber(118));
        };
        let _nw226 = Array((new BigNumber(10)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw226.length); _i0_36++) {
          _nw226[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1596_v32 = _nw226;
        let _1598_v33;
        _1598_v33 = _module.D2.create_DC4(_1596_v32, (_this).f17, _1556_v0, false, (_this).f14);
        let _1599_v34;
        _1599_v34 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17),(_this).f14);
        let _1600_v35;
        let _nw227 = Array((new BigNumber(17)).toNumber());
        _nw227[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw227[(_dafny.ONE).toNumber()] = (_1595_v31).dtor_cf5;
        _nw227[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(3)).toNumber()] = (_1598_v33).dtor_cf10;
        _nw227[(new BigNumber(4)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(6)).toNumber()] = (_1598_v33).dtor_cf10;
        _nw227[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(11)).toNumber()] = (((_1599_v34).contains(true)) ? ((_1599_v34).get(true)) : ((_this).f14));
        _nw227[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(13)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(14)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(15)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(16)).toNumber()] = (_this).f14;
        _1600_v35 = _nw227;
        let _index266 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_1600_v35).length));
        (_1600_v35)[_index266] = (_this).f14;
        let _index267 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_1600_v35).length));
        (_1600_v35)[_index267] = (_this).f14;
      }
      let _1601_v36;
      let _nw228 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
      _1601_v36 = _nw228;
      let _nw229 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
      _1601_v36 = _nw229;
      let _1602_v37;
      let _nw230 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _1602_v37 = _nw230;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1602_v37).length))) {
        let _1603_i4 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1603_i4)) && ((_1603_i4).isLessThan(new BigNumber((_1602_v37).length))))) {
          (_1602_v37)[(_1603_i4)] = _module.__default.safeModuloInt(_1603_i4, _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_1556_v0)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dscsy")).length)));
        }
      }
      let _hi11 = _1556_v0;
      for (let _1604_i5 = _1556_v0; _1604_i5.isLessThan(_hi11); _1604_i5 = _1604_i5.plus(_dafny.ONE)) {
        let _1605_v38;
        _1605_v38 = false;
        _1605_v38 = !(p0);
        let _1606_v39;
        _1606_v39 = _dafny.MultiSet.fromElements(_1556_v0);
        let _1607_v40;
        _1607_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1606_v39);
        (globalState).f2 = _dafny.Seq.update(_1578_v19, _module.__default.safeIndex(new BigNumber(((((_1607_v40).contains(!(p0))) ? ((_1607_v40).get(!(p0))) : (_1606_v39))).cardinality()), new BigNumber((_1578_v19).length)), ((true) ? (_1605_v38) : ((_this).f13)));
        let _1608_v41;
        _1608_v41 = _dafny.Seq.UnicodeFromString("fbqky");
        let _1609_v42;
        _1609_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-435),p0);
        let _rhs235 = _module.__default.safeModuloInt(_1556_v0, _1604_i5);
        let _rhs236 = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-811)), function (_1610_i6) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), _1608_v41);
        let _rhs237 = (((((_1609_v42).contains(_1604_i5)) ? ((_1609_v42).get(_1604_i5)) : (_1605_v38))) ? (_1577_v18) : ((_1577_v18).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f17))));
        _1556_v0 = _rhs235;
        _1605_v38 = _rhs236;
        _1577_v18 = _rhs237;
        let _1611_v43;
        _1611_v43 = _module.D2.create_DC3((_this).f14);
        _1611_v43 = function (_pat_let47_0) {
          return function (_1612_dt__update__tmp_h1) {
            return function (_pat_let48_0) {
              return function (_1613_dt__update_hcf5_h0) {
                return _module.D2.create_DC3(_1613_dt__update_hcf5_h0);
              }(_pat_let48_0);
            }((_this).f14);
          }(_pat_let47_0);
        }(_1611_v43);
      }
      let _1614_v44;
      _1614_v44 = _dafny.Seq.UnicodeFromString("axb");
      if (((_1556_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), ((_1615_v0) => function (_1616_i7) {
        return _1615_v0;
      })(_1556_v0))).length))).isLessThan(new BigNumber((_dafny.Seq.Concat(_1614_v44, _dafny.Seq.UnicodeFromString("gsllko"))).length))) {
        let _1617_v45;
        _1617_v45 = new _dafny.CodePoint('h'.codePointAt(0));
        _1617_v45 = _1617_v45;
        let _1618_v46;
        _1618_v46 = _dafny.MultiSet.fromElements(_1602_v37);
        let _1619_v47;
        _1619_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1617_v45,(_this).f13);
        let _1620_v48;
        let _nw231 = new _module.C0();
        _nw231.__ctor((_1618_v46).IsSubsetOf(_1618_v46), (((_this).f13) ? (_1617_v45) : (_1617_v45)), _this.f12, (_this).fm12(_1619_v47, _this.f12, globalState));
        _1620_v48 = _nw231;
        let _1621_v49;
        let _init37 = ((_1622_v0) => function (_1623_i8) {
          return _module.D7.create_DC16(_1622_v0);
        })(_1556_v0);
        let _nw232 = Array((new BigNumber(17)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw232.length); _i0_37++) {
          _nw232[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1621_v49 = _nw232;
        let _1624_v50;
        _1624_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm46(_1556_v0, (((_1577_v18).contains((_this).f17)) ? ((_1577_v18).get((_this).f17)) : ((_this.f12).dtor_cf4)), (_this).f17, globalState),_1621_v49);
        let _index268 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length));
        (_1602_v37)[_index268] = new BigNumber((_1624_v50).length);
        let _index269 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length));
        (_1602_v37)[_index269] = new BigNumber((_module.__default.fm47(p0, globalState)).length);
        let _1625_v51;
        _1625_v51 = _dafny.Map.Empty.slice().updateUnsafe((_1602_v37)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length))],_1614_v44);
        let _1626_v52;
        let _nw233 = Array((new BigNumber(23)).toNumber());
        _nw233[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("wnydrhnob");
        _nw233[(_dafny.ONE).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-269)), ((_1627_v48) => function (_1628_i9) {
          return (_1627_v48).f20;
        })(_1620_v48));
        _nw233[(new BigNumber(3)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(4)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fh"), _1614_v44);
        _nw233[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("orofsd");
        _nw233[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1614_v44, _1614_v44), _module.__default.safeIndex((_1602_v37)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length))], new BigNumber((_dafny.Seq.Concat(_1614_v44, _1614_v44)).length)), new _dafny.CodePoint('f'.codePointAt(0)));
        _nw233[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ftv"), _1614_v44);
        _nw233[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1614_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), function (_1629_i10) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }));
        _nw233[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("lgqso");
        _nw233[(new BigNumber(11)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1614_v44, _dafny.Seq.update(_1614_v44, _module.__default.safeIndex(_1556_v0, new BigNumber((_1614_v44).length)), new _dafny.CodePoint('f'.codePointAt(0))));
        _nw233[(new BigNumber(13)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(14)).toNumber()] = (((_1625_v51).contains((_dafny.ZERO).minus((_1602_v37)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length))]))) ? ((_1625_v51).get((_dafny.ZERO).minus((_1602_v37)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_1602_v37).length))]))) : (_1614_v44));
        _nw233[(new BigNumber(15)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-534)), ((_1630_v48) => function (_1631_i11) {
          return (_1630_v48).f20;
        })(_1620_v48));
        _nw233[(new BigNumber(17)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(18)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(19)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(20)).toNumber()] = _1614_v44;
        _nw233[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_1614_v44, _dafny.Seq.UnicodeFromString("xbclijk"));
        _nw233[(new BigNumber(22)).toNumber()] = _1614_v44;
        _1626_v52 = _nw233;
        _1626_v52 = _1626_v52;
        let _1632_v53;
        _1632_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0);
        let _1633_v54;
        let _nw234 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1633_v54 = _nw234;
        let _1634_v55;
        _1634_v55 = _module.D2.create_DC4(_1633_v54, _1620_v48.f19, _1556_v0, _1620_v48.f19, (_this).f14);
        let _pat_let_tv42 = _1602_v37;
        let _pat_let_tv43 = _1602_v37;
        _1632_v53 = (_1632_v53).update((_this).f14, (function (_pat_let49_0) {
          return function (_1635_dt__update__tmp_h2) {
            return function (_pat_let50_0) {
              return function (_1636_dt__update_hcf8_h1) {
                return function (_pat_let51_0) {
                  return function (_1637_dt__update_hcf9_h0) {
                    return _module.D2.create_DC4((_1635_dt__update__tmp_h2).dtor_cf6, (_1635_dt__update__tmp_h2).dtor_cf7, _1636_dt__update_hcf8_h1, _1637_dt__update_hcf9_h0, (_1635_dt__update__tmp_h2).dtor_cf10);
                  }(_pat_let51_0);
                }((_this).f17);
              }(_pat_let50_0);
            }((_pat_let_tv43)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_pat_let_tv42).length))]);
          }(_pat_let49_0);
        }(_1634_v55)).dtor_cf9);
      } else {
        let _1638_v56;
        _1638_v56 = false;
        _1638_v56 = _module.__default.fm17(globalState);
        r0 = _1556_v0;
        let _1639_v57;
        _1639_v57 = new _dafny.CodePoint('q'.codePointAt(0));
        let _rhs238 = new BigNumber(486);
        let _rhs239 = _1639_v57;
        r0 = _rhs238;
        _1639_v57 = _rhs239;
        let _1640_v58;
        _1640_v58 = _dafny.Seq.of(_1556_v0);
        if ((function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of (_dafny.Map.Empty.slice().updateUnsafe((_1640_v58)[_module.__default.safeIndex(_1556_v0, new BigNumber((_1640_v58).length))],p0)).Keys.Elements) {
            let _1641_v59 = _compr_49;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_1640_v58)[_module.__default.safeIndex(_1556_v0, new BigNumber((_1640_v58).length))],p0)).contains(_1641_v59)) {
              _coll49.add((_1641_v59).plus(new BigNumber(-552)));
            }
          }
          return _coll49;
        }()).IsDisjointFrom((_this).f15)) {
          let _1642_v60;
          _1642_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(309),_1556_v0);
          let _1643_v61;
          _1643_v61 = _dafny.MultiSet.fromElements((((_1642_v60).contains((_dafny.ZERO).minus(_1556_v0))) ? ((_1642_v60).get((_dafny.ZERO).minus(_1556_v0))) : (_1556_v0)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_module.__default.fm27(globalState), _dafny.Seq.of(new BigNumber(126), _1556_v0))).length)));
          _1643_v61 = (_1643_v61).Union(_1643_v61);
          let _index270 = _module.__default.safeIndex(new BigNumber(981), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index270] = (_this).f13;
          let _index271 = _module.__default.safeIndex(new BigNumber(981), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index271] = (_this).f13;
          let _index272 = _module.__default.safeIndex(new BigNumber(981), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index272] = !(!(((_this).f14)[_module.__default.safeIndex(new BigNumber(981), new BigNumber(((_this).f14).length))]));
          let _index273 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1602_v37).length));
          (_1602_v37)[_index273] = _1556_v0;
          let _index274 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1602_v37).length));
          (_1602_v37)[_index274] = _1556_v0;
          (globalState).f7 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1614_v44, _dafny.Seq.update(_1614_v44, _module.__default.safeIndex(_1556_v0, new BigNumber((_1614_v44).length)), _1639_v57)), _module.__default.safeIndex(_1556_v0, new BigNumber((_dafny.Seq.Concat(_1614_v44, _dafny.Seq.update(_1614_v44, _module.__default.safeIndex(_1556_v0, new BigNumber((_1614_v44).length)), _1639_v57))).length)), _1639_v57)).length);
        } else {
          let _index275 = _module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index275] = !(p0) || (p0);
          let _index276 = _module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index276] = (_this).f13;
          (globalState).f7 = (_this).fm8((_this).f17, _this.f16, globalState);
          let _1644_v62;
          _1644_v62 = _dafny.Seq.of(_1614_v44);
          let _1645_v63;
          _1645_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm38((_this).f13, ((_this).f14)[_module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length))], false, ((_this).f14)[_module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length))], globalState),!_dafny.Seq.contains(_dafny.Seq.update(_1644_v62, _module.__default.safeIndex(_1556_v0, new BigNumber((_1644_v62).length)), _1614_v44), _1614_v44));
          let _1646_v64;
          _1646_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1578_v19).length),!(((_this).f14)[_module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length))]));
          let _1647_v65;
          _1647_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1646_v64).update(new BigNumber((_1578_v19).length), (_this).f13)).length),(_this).f13);
          _1645_v63 = (_1645_v63).update(_dafny.Seq.Concat(_1578_v19, _1578_v19), (((_this).f13) ? (_1638_v56) : ((((_1647_v65).contains(_1556_v0)) ? ((_1647_v65).get(_1556_v0)) : (((_this).f14)[_module.__default.safeIndex(new BigNumber(411), new BigNumber(((_this).f14).length))])))));
          let _index277 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1602_v37).length));
          (_1602_v37)[_index277] = ((_dafny.ZERO).minus(new BigNumber((_module.__default.fm20(globalState)).length))).plus(_1556_v0);
          let _index278 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1602_v37).length));
          let _rhs240 = _module.__default.safeDivisionInt((_this).fm9(new _dafny.CodePoint('g'.codePointAt(0)), globalState), _1556_v0);
          let _rhs241 = _1639_v57;
          let _rhs242 = (_1556_v0).minus(new BigNumber(-984));
          let _lhs197 = _1602_v37;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1602_v37).length));
          _lhs197[_lhs198] = _rhs240;
          _1639_v57 = _rhs241;
          _1556_v0 = _rhs242;
          let _1648_v66;
          _1648_v66 = _dafny.Set.fromElements((_dafny.ZERO).minus((_1602_v37)[_module.__default.safeIndex(new BigNumber(584), new BigNumber((_1602_v37).length))]));
          _1648_v66 = (_this).f15;
        }
        if (_1638_v56) {
          let _1649_v67;
          _1649_v67 = _module.D9.create_DC20((_this).f15);
          let _1650_v68;
          _1650_v68 = _dafny.MultiSet.fromElements(_1649_v67, _1649_v67);
          let _index279 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length));
          (_1602_v37)[_index279] = _1556_v0;
          let _index280 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length));
          let _rhs243 = ((_1650_v68).Intersect(_1650_v68)).Difference(_1650_v68);
          let _rhs244 = _1556_v0;
          let _rhs245 = true;
          let _rhs246 = _1578_v19;
          let _lhs199 = _1602_v37;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length));
          let _lhs201 = globalState;
          _1650_v68 = _rhs243;
          _lhs199[_lhs200] = _rhs244;
          _1638_v56 = _rhs245;
          _lhs201.f2 = _rhs246;
          let _index281 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index281] = true;
          let _1651_v69;
          let _nw235 = new _module.C3();
          _nw235.__ctor((_this).f13, _this.f16, _1638_v56, (_this).f14, (_this).f15, _module.D1.create_DC2(_1556_v0, (_1602_v37)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length))], (_this).f17), _1638_v56);
          _1651_v69 = _nw235;
          let _index282 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index282] = _module.__default.fm30(((_1602_v37)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length))]).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1651_v69)).length)), globalState);
          let _1652_v70;
          _1652_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1602_v37)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length))],(_1602_v37)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length))]);
          let _1653_v71;
          _1653_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1652_v70,_module.__default.fm0(_1578_v19, !(true), p0, (_1602_v37)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1602_v37).length))], globalState));
          let _1654_v72;
          _1654_v72 = _dafny.MultiSet.fromElements(p0, _1638_v56, p0);
          let _index283 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          let _index284 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          let _rhs247 = _module.__default.fm38(!(false) || ((_this).f13), ((((_this).f14)[_module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length))]) ? ((_this).f17) : (_1638_v56)), _1638_v56, _1638_v56, globalState);
          let _rhs248 = (((_this).f14)[_module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length))]) && ((_1651_v69).f18);
          let _rhs249 = (_module.D11.create_DC25(_1653_v71)).dtor_cf48;
          let _rhs250 = (_1577_v18).Merge((_1577_v18).update((_this).f13, (_this).f17));
          let _rhs251 = ((_1654_v72).Intersect(_dafny.MultiSet.FromArray(_1578_v19))).equals(_1654_v72);
          let _lhs202 = globalState;
          let _lhs203 = (_this).f14;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          let _lhs205 = (_this).f14;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(181), new BigNumber(((_this).f14).length));
          _lhs202.f2 = _rhs247;
          _lhs203[_lhs204] = _rhs248;
          _1653_v71 = _rhs249;
          _1577_v18 = _rhs250;
          _lhs205[_lhs206] = _rhs251;
          let _1655_v73;
          _1655_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((_this).f15).Union((_this).f15)).length),((_this).f13) && ((_1651_v69).f18));
          _1655_v73 = (_1655_v73).Merge(_1655_v73);
          _1655_v73 = (_1655_v73).update(_module.__default.safeModuloInt(_1556_v0, new BigNumber(671)), true);
        } else {
          let _1656_v74;
          _1656_v74 = _module.D11.create_DC26(_1602_v37, _1614_v44);
          let _1657_v75;
          _1657_v75 = _module.D11.create_DC28(_1656_v74);
          let _1658_v76;
          let _nw236 = Array((new BigNumber(9)).toNumber());
          _nw236[(_dafny.ZERO).toNumber()] = _1657_v75;
          _nw236[(_dafny.ONE).toNumber()] = _1657_v75;
          _nw236[(new BigNumber(2)).toNumber()] = _1657_v75;
          _nw236[(new BigNumber(3)).toNumber()] = _1657_v75;
          _nw236[(new BigNumber(4)).toNumber()] = _1657_v75;
          _nw236[(new BigNumber(5)).toNumber()] = _module.D11.create_DC28(_1656_v74);
          _nw236[(new BigNumber(6)).toNumber()] = _1657_v75;
          _nw236[(new BigNumber(7)).toNumber()] = _module.D11.create_DC28(_1656_v74);
          _nw236[(new BigNumber(8)).toNumber()] = _module.D11.create_DC28(_1656_v74);
          _1658_v76 = _nw236;
          let _1659_v77;
          _1659_v77 = _dafny.Seq.of(_1658_v76);
          _1659_v77 = _1659_v77;
          _1638_v56 = ((_this).f15).IsDisjointFrom(_module.__default.fm47((_this).f13, globalState));
          _1602_v37 = _1602_v37;
          let _1660_v78;
          _1660_v78 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1556_v0).multipliedBy(_1556_v0));
          _1660_v78 = (_1660_v78).update((_this).f17, _1556_v0);
          let _1661_v79;
          _1661_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1556_v0,_1639_v57);
          let _1662_v80;
          _1662_v80 = _module.D3.create_DC5(_1639_v57);
          _1661_v79 = (_1661_v79).update(_1556_v0, (_1662_v80).dtor_cf11);
        }
      }
      r0 = new BigNumber(411);
      return r0;
    }
    m7(globalState) {
      let _this = this;
      let _1663_i0;
      _1663_i0 = _dafny.ZERO;
      L9: {
        while ((_this).f13) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1663_i0)) {
              break L9;
            }
            _1663_i0 = (_1663_i0).plus(_dafny.ONE);
            let _1664_v0;
            _1664_v0 = new _dafny.CodePoint('c'.codePointAt(0));
            let _1665_v1;
            let _nw237 = new _module.C0();
            _nw237.__ctor(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("erlwcj"), _1664_v0), _1664_v0, _this.f12, (_this).f13);
            _1665_v1 = _nw237;
            let _1666_v2;
            _1666_v2 = new BigNumber(697);
            let _1667_v3;
            let _nw238 = new _module.C1();
            _nw238.__ctor(_1664_v0, _module.D1.create_DC2(_1666_v2, _1666_v2, (_this).f13), (_this).f17);
            _1667_v3 = _nw238;
            let _1668_v4;
            _1668_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(_1665_v1.f19),_1667_v3);
            let _1669_v5;
            _1669_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1664_v0,_1665_v1.f19);
            let _1670_v6;
            _1670_v6 = _dafny.Seq.of((_this).f17);
            _1668_v4 = (_1668_v4).update((_this).fm12(_1669_v5, _module.D1.create_DC2(new BigNumber((_1670_v6).length), _1666_v2, (_this).f17), globalState), _1667_v3);
            let _1671_v7;
            _1671_v7 = _module.D3.create_DC5(_module.__default.fm33(_1665_v1.f19, new BigNumber(4), globalState));
            _1671_v7 = _1671_v7;
            let _1672_v8;
            let _nw239 = new _module.C3();
            _nw239.__ctor((_this).f17, _module.D1.create_DC1(_dafny.Seq.UnicodeFromString("yitlqi")), (_this).fm12(_1669_v5, _module.D1.create_DC2(_1666_v2, new BigNumber(219), (_this).f17), globalState), (_this).f14, (_this).f15, _module.D1.create_DC2(_1666_v2, new BigNumber(478), false), _1665_v1.f19);
            _1672_v8 = _nw239;
          }
        }
      }
      let _1673_v9;
      _1673_v9 = new BigNumber(993);
      (globalState).f7 = _1673_v9;
      _1673_v9 = _1673_v9;
      let _1674_v10;
      _1674_v10 = _dafny.Seq.of(true);
      let _1675_v11;
      _1675_v11 = _dafny.Seq.of(_dafny.Seq.Concat(_1674_v10, _1674_v10));
      (globalState).f2 = (_1675_v11)[_module.__default.safeIndex(_1673_v9, new BigNumber((_1675_v11).length))];
      let _1676_v12;
      _1676_v12 = _dafny.Seq.UnicodeFromString("r");
      let _1677_v13;
      _1677_v13 = _dafny.Set.fromElements(_1676_v12, _1676_v12);
      let _1678_v14;
      _1678_v14 = (((_this).f13) ? (_module.__default.fm48(globalState)) : (_1677_v13));
      _1678_v14 = _1678_v14;
      let _1679_v15;
      let _nw240 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
      _1679_v15 = _nw240;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1679_v15).length))) {
        let _1680_i1 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1680_i1)) && ((_1680_i1).isLessThan(new BigNumber((_1679_v15).length))))) {
          (_1679_v15)[(_1680_i1)] = ((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1673_v9)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-165)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1673_v9)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1673_v9)));
        }
      }
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _1681_i0;
      _1681_i0 = _dafny.ZERO;
      L10: {
        while ((_this).f17) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1681_i0)) {
              break L10;
            }
            _1681_i0 = (_1681_i0).plus(_dafny.ONE);
            let _1682_v0;
            _1682_v0 = _dafny.Seq.of((_this).f17, (_this).f13);
            let _1683_v1;
            _1683_v1 = new _dafny.CodePoint('o'.codePointAt(0));
            let _1684_v2;
            _1684_v2 = _module.D3.create_DC5(_1683_v1);
            let _1685_v3;
            _1685_v3 = _dafny.MultiSet.fromElements((_this).f17, (_1682_v0)[_module.__default.safeIndex(p0, new BigNumber((_1682_v0).length))], (_this).fm12(_dafny.Map.Empty.slice().updateUnsafe((_1684_v2).dtor_cf11,(_this).f17), _module.D1.create_DC2(p0, p0, (_this).f13), globalState));
            let _1686_v4;
            _1686_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1683_v1,p0);
            let _1687_v5;
            let _init38 = function (_1688_i1) {
              return _module.__default.safeModuloInt(_1688_i1, new BigNumber((_dafny.Set.fromElements((_this).f17, !((_this).f13), (_this).f17, (_this).f13)).length));
            };
            let _nw241 = Array((new BigNumber(20)).toNumber());
            for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw241.length); _i0_38++) {
              _nw241[_i0_38] = _init38(new BigNumber(_i0_38));
            }
            _1687_v5 = _nw241;
            let _1689_v6;
            _1689_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v4,_1687_v5);
            let _1690_v7;
            _1690_v7 = _dafny.Seq.of(p0, p0);
            let _1691_v8;
            _1691_v8 = _dafny.Set.fromElements((_this).f17, (_this).f17);
            let _1692_v9;
            _1692_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v8,_dafny.Set.fromElements(p0));
            let _1693_v10;
            let _nw242 = Array((new BigNumber(13)).toNumber());
            _nw242[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1685_v3).cardinality()));
            _nw242[(_dafny.ONE).toNumber()] = new BigNumber((((false) ? (_1689_v6) : (_1689_v6))).length);
            _nw242[(new BigNumber(2)).toNumber()] = p0;
            _nw242[(new BigNumber(3)).toNumber()] = (p0).multipliedBy(p0);
            _nw242[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_1690_v7)[_module.__default.safeIndex(p0, new BigNumber((_1690_v7).length))], p0)).length);
            _nw242[(new BigNumber(5)).toNumber()] = (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(740)), function (_1694_i2) {
              return new BigNumber(811);
            }))).cardinality())).multipliedBy(p0);
            _nw242[(new BigNumber(6)).toNumber()] = new BigNumber((_1691_v8).length);
            _nw242[(new BigNumber(7)).toNumber()] = p0;
            _nw242[(new BigNumber(8)).toNumber()] = new BigNumber(-766);
            _nw242[(new BigNumber(9)).toNumber()] = p0;
            _nw242[(new BigNumber(10)).toNumber()] = (_1690_v7)[_module.__default.safeIndex(p0, new BigNumber((_1690_v7).length))];
            _nw242[(new BigNumber(11)).toNumber()] = p0;
            _nw242[(new BigNumber(12)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1691_v8,(_this).f15)).Merge(_1692_v9)).length);
            _1693_v10 = _nw242;
            let _index285 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length));
            (_1687_v5)[_index285] = p0;
            let _1695_v11;
            _1695_v11 = _dafny.Seq.of(_1691_v8, _1691_v8);
            let _index286 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length));
            (_1687_v5)[_index286] = (_module.D9.create_DC22((_this).f17, p0, false, (_1695_v11)[_module.__default.safeIndex(p0, new BigNumber((_1695_v11).length))], _dafny.Seq.UnicodeFromString("apxjf"))).dtor_cf38;
            _1685_v3 = (_dafny.MultiSet.FromArray(_module.__default.fm38((_this).f17, (_this).f13, (_this).f17, _module.__default.fm30(new BigNumber(596), globalState), globalState))).update(false, _module.__default.abs(p0));
            let _1696_v12;
            let _nw243 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
            _1696_v12 = _nw243;
            let _index287 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length));
            let _rhs252 = p0;
            let _rhs253 = _1696_v12;
            let _rhs254 = false;
            let _lhs207 = _1687_v5;
            let _lhs208 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length));
            _lhs207[_lhs208] = _rhs252;
            _1696_v12 = _rhs253;
            r0 = _rhs254;
            let _1697_v13;
            _1697_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(814));
            let _index288 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length));
            (_1687_v5)[_index288] = _module.__default.safeDivisionInt((((_1697_v13).contains((_1687_v5)[_module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length))])) ? ((_1697_v13).get((_1687_v5)[_module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length))])) : ((_1687_v5)[_module.__default.safeIndex(new BigNumber(801), new BigNumber((_1687_v5).length))])), (_1690_v7)[_module.__default.safeIndex(p0, new BigNumber((_1690_v7).length))]);
          }
        }
      }
      let _1698_v14;
      let _nw244 = Array((new BigNumber(26)).toNumber()).fill([]);
      _1698_v14 = _nw244;
      let _1699_v15;
      let _init39 = function (_1700_i3) {
        return (_1700_i3).plus(new BigNumber(683));
      };
      let _nw245 = Array((new BigNumber(19)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw245.length); _i0_39++) {
        _nw245[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1699_v15 = _nw245;
      let _index289 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length));
      (_1698_v14)[_index289] = _1699_v15;
      let _1701_v16;
      _1701_v16 = _dafny.MultiSet.fromElements(_this.f12);
      let _1702_v18;
      _1702_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,true);
      let _index290 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length));
      (_1698_v14)[_index290] = (((function () {
        let _coll50 = new _dafny.Set();
        for (const _compr_50 of ((_1701_v16).update(_this.f12, _module.__default.abs(p0))).Elements) {
          let _1703_v17 = _compr_50;
          if (((_1701_v16).update(_this.f12, _module.__default.abs(p0))).contains(_1703_v17)) {
            _coll50.add(_1703_v17);
          }
        }
        return _coll50;
      }()).equals(function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of (_1702_v18).Keys.Elements) {
          let _1704_v19 = _compr_51;
          if ((_1702_v18).contains(_1704_v19)) {
            _coll51.add(_1704_v19);
          }
        }
        return _coll51;
      }())) ? (_1699_v15) : (_1699_v15));
      let _1705_v20;
      _1705_v20 = _module.D7.create_DC16(p0);
      let _1706_v21;
      _1706_v21 = _module.D7.create_DC17(_1705_v20);
      let _pat_let_tv44 = _1705_v20;
      let _source24 = function (_pat_let52_0) {
        return function (_1707_dt__update__tmp_h0) {
          return function (_pat_let53_0) {
            return function (_1708_dt__update_hcf30_h0) {
              return _module.D7.create_DC17(_1708_dt__update_hcf30_h0);
            }(_pat_let53_0);
          }(_pat_let_tv44);
        }(_pat_let52_0);
      }(_1706_v21);
      if (_source24.is_DC16) {
        let _1709___mcc_h0 = (_source24).cf29;
        let _1710_cf29 = _1709___mcc_h0;
        _1699_v15 = _1699_v15;
        let _1711_v22;
        _1711_v22 = _dafny.Seq.of((_this).f13);
        let _1712_v23;
        _1712_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
        let _1713_v24;
        _1713_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_1712_v23).contains(_1710_cf29)) ? ((_1712_v23).get(_1710_cf29)) : ((_this).f17)),p0);
        let _1714_v25;
        _1714_v25 = _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)));
        let _1715_v26;
        _1715_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1714_v25);
        let _index291 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length));
        let _nw246 = Array((new BigNumber(7)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_1710_cf29, (_dafny.ZERO).minus(_module.__default.fm0(_1711_v22, !((_this).f13), (_this).f17, (((_1713_v24).contains((_this).f17)) ? ((_1713_v24).get((_this).f17)) : (_1710_cf29)), globalState)));
        _nw246[(_dafny.ONE).toNumber()] = _1710_cf29;
        _nw246[(new BigNumber(2)).toNumber()] = new BigNumber(((_1715_v26).Merge(_1715_v26)).length);
        _nw246[(new BigNumber(3)).toNumber()] = p0;
        _nw246[(new BigNumber(4)).toNumber()] = p0;
        _nw246[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f15).length));
        _nw246[(new BigNumber(6)).toNumber()] = _1710_cf29;
        (_1698_v14)[_index291] = _nw246;
        let _arr10 = (_1698_v14)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length))];
        let _index292 = _module.__default.safeIndex(new BigNumber(281), new BigNumber(((_1698_v14)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length))]).length));
        _arr10[_index292] = _1710_cf29;
        let _arr11 = (_1698_v14)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length))];
        let _index293 = _module.__default.safeIndex(new BigNumber(281), new BigNumber(((_1698_v14)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_1698_v14).length))]).length));
        _arr11[_index293] = new BigNumber(943);
        _1714_v25 = _dafny.Seq.UnicodeFromString("g");
      } else if (_source24.is_DC15) {
        let _1716___mcc_h1 = (_source24).cf28;
        let _1717_cf28 = _1716___mcc_h1;
        let _1718_v27;
        _1718_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!((_this).f17));
        _1718_v27 = (_1718_v27).update((_this).f15, (_this).f13);
        let _1719_v28;
        let _init40 = function (_1720_i4) {
          return (_this).f17;
        };
        let _nw247 = Array((new BigNumber(11)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw247.length); _i0_40++) {
          _nw247[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1719_v28 = _nw247;
        let _rhs255 = (_this).f17;
        let _rhs256 = _1719_v28;
        r0 = _rhs255;
        _1719_v28 = _rhs256;
        let _1721_v29;
        _1721_v29 = new _dafny.CodePoint('p'.codePointAt(0));
        let _1722_v30;
        _1722_v30 = _dafny.Seq.of(_1721_v29, new _dafny.CodePoint('k'.codePointAt(0)), _1721_v29);
        let _1723_v31;
        let _nw248 = new _module.C0();
        _nw248.__ctor((_this).f13, (_1722_v30)[_module.__default.safeIndex(p0, new BigNumber((_1722_v30).length))], _this.f12, true);
        _1723_v31 = _nw248;
        let _1724_v32;
        _1724_v32 = _dafny.Seq.of(_1723_v31, _1723_v31);
        let _1725_v33;
        let _nw249 = Array((new BigNumber(14)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = (_1724_v32)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1724_v32).length))];
        _nw249[(_dafny.ONE).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(2)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(3)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(4)).toNumber()] = (((_this).f13) ? (_1723_v31) : (_1723_v31));
        _nw249[(new BigNumber(5)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(6)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(7)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(8)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(9)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(10)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(11)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(12)).toNumber()] = _1723_v31;
        _nw249[(new BigNumber(13)).toNumber()] = _1723_v31;
        _1725_v33 = _nw249;
        let _index294 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1725_v33).length));
        (_1725_v33)[_index294] = _1723_v31;
        let _index295 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
        (_1699_v15)[_index295] = p0;
        let _1726_v34;
        let _nw250 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
        _1726_v34 = _nw250;
        let _1727_v35;
        _1727_v35 = _dafny.Seq.of(_1726_v34);
        let _1728_v36;
        _1728_v36 = _module.D8.create_DC18((_1727_v35)[_module.__default.safeIndex(p0, new BigNumber((_1727_v35).length))]);
        let _index296 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1725_v33).length));
        let _index297 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
        let _rhs257 = new BigNumber(883);
        let _rhs258 = _1723_v31;
        let _rhs259 = (_dafny.ZERO).minus((((_this).f13) ? ((p0).plus(new BigNumber(751))) : ((p0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ffp")).length)))));
        let _rhs260 = _module.D8.create_DC18(_1726_v34);
        let _rhs261 = _1723_v31.f19;
        let _lhs209 = globalState;
        let _lhs210 = _1725_v33;
        let _lhs211 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1725_v33).length));
        let _lhs212 = _1699_v15;
        let _lhs213 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
        let _lhs214 = _1723_v31;
        _lhs209.f7 = _rhs257;
        _lhs210[_lhs211] = _rhs258;
        _lhs212[_lhs213] = _rhs259;
        _1728_v36 = _rhs260;
        _lhs214.f19 = _rhs261;
        if (_1723_v31.f19) {
          let _1729_v37;
          _1729_v37 = _dafny.Seq.of(_1723_v31.f19, _1723_v31.f19);
          let _1730_v38;
          _1730_v38 = _module.D3.create_DC8((_dafny.ZERO).minus(_module.__default.fm0(_1729_v37, _1723_v31.f19, !(_1723_v31.f19), (_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))], globalState)));
          let _1731_v39;
          _1731_v39 = _dafny.Set.fromElements((_1723_v31).f20, (_1723_v31).f20);
          let _1732_v40;
          _1732_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1721_v29,_module.__default.fm17(globalState));
          let _pat_let_tv45 = _1699_v15;
          let _pat_let_tv46 = _1699_v15;
          let _pat_let_tv47 = _1699_v15;
          let _pat_let_tv48 = _1699_v15;
          let _index298 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
          let _rhs262 = _module.__default.fm49(_this.f16, _1731_v39, globalState);
          let _rhs263 = (_1723_v31).fm16((_this).fm12(_1732_v40, function (_pat_let54_0) {
            return function (_1733_dt__update__tmp_h1) {
              return function (_pat_let55_0) {
                return function (_1734_dt__update_hcf3_h0) {
                  return function (_pat_let56_0) {
                    return function (_1735_dt__update_hcf2_h0) {
                      return _module.D1.create_DC2(_1735_dt__update_hcf2_h0, _1734_dt__update_hcf3_h0, (_1733_dt__update__tmp_h1).dtor_cf4);
                    }(_pat_let56_0);
                  }((_pat_let_tv48)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_pat_let_tv47).length))]);
                }(_pat_let55_0);
              }((_pat_let_tv46)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_pat_let_tv45).length))]);
            }(_pat_let54_0);
          }(_this.f12), globalState), p0, globalState);
          let _rhs264 = _1723_v31.f19;
          let _lhs215 = _1699_v15;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
          let _lhs217 = _1723_v31;
          _1730_v38 = _rhs262;
          _lhs215[_lhs216] = _rhs263;
          _lhs217.f19 = _rhs264;
          let _index299 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index299] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((((_1723_v31.f19) ? (_module.__default.fm17(globalState)) : (_1723_v31.f19))) ? ((_1717_cf28)[_module.__default.safeIndex((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))], new BigNumber((_1717_cf28).length))]) : (_module.__default.fm0(_dafny.Seq.of(_1723_v31.f19, false, (_1729_v37)[_module.__default.safeIndex(new BigNumber((_1729_v37).length), new BigNumber((_1729_v37).length))]), (_this).f17, (_this).f17, (_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))], globalState)))));
          let _1736_v41;
          _1736_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1729_v37);
          _1736_v41 = (_1736_v41).update(new BigNumber((function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of _dafny.IntegerRange(new BigNumber(736), new BigNumber(816))) {
              let _1737_v42 = _compr_52;
              if (((new BigNumber(736)).isLessThanOrEqualTo(_1737_v42)) && ((_1737_v42).isLessThan(new BigNumber(816)))) {
                _coll52.add(_module.__default.safeModuloInt(_1737_v42, new BigNumber((_dafny.Seq.of(_1723_v31.f19)).length)));
              }
            }
            return _coll52;
          }()).length), _dafny.Seq.Concat(_1729_v37, _1729_v37));
          let _1738_v43;
          _1738_v43 = _dafny.MultiSet.fromElements((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))]);
          let _1739_v44;
          _1739_v44 = _module.D6.create_DC11(_1738_v43);
          let _1740_v45;
          _1740_v45 = _module.D6.create_DC14(_1739_v44);
          let _1741_v46;
          let _nw251 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
          _1741_v46 = _nw251;
          let _index300 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1741_v46).length));
          (_1741_v46)[_index300] = _module.__default.fm50(globalState);
          let _1742_v47;
          _1742_v47 = _dafny.MultiSet.fromElements(((!(_1723_v31.f19)) ? (_1723_v31.f19) : ((((_1732_v40).contains(_1721_v29)) ? ((_1732_v40).get(_1721_v29)) : ((_this).f17)))), false, false, !(((_this).f17) && ((_this).f17)), _dafny.Seq.IsProperPrefixOf(_1722_v30, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ido"), _module.__default.safeIndex((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))], new BigNumber((_dafny.Seq.UnicodeFromString("ido")).length)), (_1723_v31).f20)));
          let _1743_v48;
          _1743_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
          let _index301 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1741_v46).length));
          let _rhs265 = (_1723_v31).f20;
          let _rhs266 = _module.D6.create_DC14(_1739_v44);
          let _rhs267 = ((_1743_v48).update(new BigNumber(376), _1723_v31.f19)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))],_1723_v31.f19));
          let _rhs268 = ((((_this).f15).IsProperSubsetOf((_this).f15)) ? (_1742_v47) : (_1742_v47));
          let _lhs218 = _1741_v46;
          let _lhs219 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1741_v46).length));
          _1721_v29 = _rhs265;
          _1740_v45 = _rhs266;
          _lhs218[_lhs219] = _rhs267;
          _1742_v47 = _rhs268;
          let _rhs269 = (_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))];
          let _rhs270 = _dafny.Map.Empty.slice().updateUnsafe(_1721_v29,(_this).f13);
          let _rhs271 = ((!(false)) ? (_1740_v45) : (_1740_v45));
          let _rhs272 = (_1729_v37)[_module.__default.safeIndex((((_this).f13) ? ((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))]) : ((_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))])), new BigNumber((_1729_v37).length))];
          let _lhs220 = globalState;
          _lhs220.f7 = _rhs269;
          _1732_v40 = _rhs270;
          _1740_v45 = _rhs271;
          r0 = _rhs272;
        } else {
          let _1744_v49;
          let _nw252 = new _module.C1();
          _nw252.__ctor((_1723_v31).f20, _this.f12, !((_this).f13));
          _1744_v49 = _nw252;
          let _1745_v50;
          _1745_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1699_v15)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((_1699_v15).length))]);
          _1745_v50 = _1745_v50;
          (_1723_v31).f19 = _1723_v31.f19;
          let _1746_v51;
          _1746_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1723_v31.f19,p0);
          _1746_v51 = (_1746_v51).update((_this).f13, p0);
          (_1723_v31).f19 = !((_this).f17) || (_dafny.areEqual(_1722_v30, _1722_v30));
        }
      } else {
        let _1747___mcc_h2 = (_source24).cf30;
        let _1748_cf30 = _1747___mcc_h2;
        let _index302 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index302] = (_this).f17;
        let _index303 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        ((_this).f14)[_index303] = (_this).f17;
        if (((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]) {
          let _1749_v52;
          let _nw253 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1749_v52 = _nw253;
          let _1750_v53;
          _1750_v53 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index304 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1749_v52).length));
          (_1749_v52)[_index304] = _1750_v53;
          let _1751_v54;
          _1751_v54 = _dafny.Seq.of((_this).f17);
          let _index305 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1749_v52).length));
          let _rhs273 = (p0).isLessThan(_module.__default.fm0(_1751_v54, ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], p0, globalState));
          let _rhs274 = _1750_v53;
          let _rhs275 = _module.__default.safeDivisionInt(p0, p0);
          let _lhs221 = _1749_v52;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1749_v52).length));
          let _lhs223 = globalState;
          r0 = _rhs273;
          _lhs221[_lhs222] = _rhs274;
          _lhs223.f7 = _rhs275;
          (globalState).f7 = (_module.__default.safeModuloInt(new BigNumber(832), p0)).multipliedBy(p0);
          let _1752_v55;
          _1752_v55 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(_this).f17);
          let _1753_v56;
          _1753_v56 = _dafny.Seq.UnicodeFromString("uwc");
          let _1754_v57;
          let _nw254 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1754_v57 = _nw254;
          let _pat_let_tv49 = _1753_v56;
          let _1755_v58;
          let _nw255 = new _module.C2();
          _nw255.__ctor(p0, (((_1752_v55).contains(new BigNumber(-780))) ? ((_1752_v55).get(new BigNumber(-780))) : ((_this).f13)), function (_pat_let57_0) {
            return function (_1756_dt__update__tmp_h2) {
              return function (_pat_let58_0) {
                return function (_1757_dt__update_hcf1_h0) {
                  return _module.D1.create_DC1(_1757_dt__update_hcf1_h0);
                }(_pat_let58_0);
              }(_pat_let_tv49);
            }(_pat_let57_0);
          }(_this.f16), (_this).f17, _1754_v57, _module.__default.fm47(!((_this).f13), globalState), _this.f12, ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]);
          _1755_v58 = _nw255;
          let _1758_v59;
          let _nw256 = new _module.C3();
          _nw256.__ctor((_this).f17, ((((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]) ? (_this.f16) : (_module.D1.create_DC1(_1753_v56))), (_this).f13, _1754_v57, ((_this).f15).Difference((_this).f15), _module.D1.create_DC2(p0, (_1755_v58).f21, ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]), (_this).f13);
          _1758_v59 = _nw256;
          let _index306 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index306] = !(p0).isEqualTo(p0);
        } else {
          let _index307 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index307] = ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))];
          let _1759_v60;
          _1759_v60 = _dafny.MultiSet.fromElements((_this).f13, ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]);
          _1759_v60 = _module.__default.fm40(globalState);
          let _1760_v61;
          _1760_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(5),((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]);
          let _1761_v62;
          _1761_v62 = _dafny.Seq.of((_this).f13);
          _1760_v61 = (_1760_v61).update(p0, (_1761_v62)[_module.__default.safeIndex(p0, new BigNumber((_1761_v62).length))]);
          let _1762_v63;
          _1762_v63 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), (_dafny.ZERO).minus(p0));
          let _1763_v64;
          _1763_v64 = _module.D6.create_DC11(_1762_v63);
          let _1764_v65;
          _1764_v65 = _module.D6.create_DC14(_module.D6.create_DC11(((_1763_v64).dtor_cf17).update(p0, _module.__default.abs(_module.__default.fm0(_dafny.Seq.of(!((_this).f17)), (_this).f13, false, p0, globalState)))));
          let _1765_v66;
          let _out12;
          _out12 = _module.__default.m0(_module.__default.fm13(_module.__default.fm32(_1764_v65, globalState), globalState), (_1761_v62)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((_1761_v62).length))], globalState);
          _1765_v66 = _out12;
          let _1766_v67;
          _1766_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1761_v62,!((_this).f13));
          let _1767_v68;
          _1767_v68 = _dafny.Seq.UnicodeFromString("u");
          let _1768_v69;
          _1768_v69 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1769_v70;
          _1769_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1768_v69,((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]);
          let _1770_v71;
          _1770_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1767_v68).length),_1768_v69);
          let _1771_v72;
          _1771_v72 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_1767_v68, _module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_dafny.Seq.update(_1761_v62, _module.__default.safeIndex(_module.__default.fm0(_1761_v62, (_this).f17, ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], p0, globalState), new BigNumber((_1761_v62).length)), (_this).f17), (_this).f17, (_this).f13, p0, globalState),p0)).update(p0, p0)).length), new BigNumber((_1767_v68).length)), _module.__default.fm33((((_1769_v70).contains((((_1770_v71).contains(new BigNumber((_1767_v68).length))) ? ((_1770_v71).get(new BigNumber((_1767_v68).length))) : (_1768_v69)))) ? ((_1769_v70).get((((_1770_v71).contains(new BigNumber((_1767_v68).length))) ? ((_1770_v71).get(new BigNumber((_1767_v68).length))) : (_1768_v69)))) : (_module.__default.fm17(globalState))), new BigNumber((_dafny.Seq.UnicodeFromString("jmgntc")).length), globalState)), _1767_v68, _1767_v68);
          let _1772_v73;
          _1772_v73 = _module.D12.create_DC29(_1766_v67);
          let _index308 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          let _rhs276 = (((false) ? (new BigNumber(148)) : (p0))).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus((((_1771_v72).contains(_dafny.Seq.UnicodeFromString("qxcd"))) ? ((_1771_v72).get(_dafny.Seq.UnicodeFromString("qxcd"))) : ((_dafny.ZERO).minus(p0))))));
          let _rhs277 = p0;
          let _rhs278 = (_1772_v73).dtor_cf57;
          let _lhs224 = (_this).f14;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          let _lhs226 = globalState;
          _lhs224[_lhs225] = _rhs276;
          _lhs226.f7 = _rhs277;
          _1766_v67 = _rhs278;
        }
        let _1773_v74;
        _1773_v74 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))]);
        let _1774_v75;
        _1774_v75 = _dafny.Seq.of(true);
        let _index309 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        let _index310 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        let _rhs279 = new BigNumber(243);
        let _rhs280 = true;
        let _rhs281 = (((_1773_v74).contains(p0)) ? ((_1773_v74).get(p0)) : (_dafny.Seq.IsPrefixOf(_1774_v75, _1774_v75)));
        let _lhs227 = globalState;
        let _lhs228 = (_this).f14;
        let _lhs229 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        let _lhs230 = (_this).f14;
        let _lhs231 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
        _lhs227.f7 = _rhs279;
        _lhs228[_lhs229] = _rhs280;
        _lhs230[_lhs231] = _rhs281;
        let _source25 = _module.__default.fm51(globalState);
        if (_source25.is_DC21) {
          let _index311 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index311] = p0;
          let _index312 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index312] = p0;
          let _1775_v76;
          _1775_v76 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1776_v77;
          let _out13;
          _out13 = _module.__default.m0(_1775_v76, false, globalState);
          _1776_v77 = _out13;
          (globalState).f7 = (_dafny.ZERO).minus(p0);
          let _index313 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index313] = !((_this).f17);
        } else if (_source25.is_DC22) {
          let _1777___mcc_h3 = (_source25).cf37;
          let _1778___mcc_h4 = (_source25).cf38;
          let _1779___mcc_h5 = (_source25).cf39;
          let _1780___mcc_h6 = (_source25).cf40;
          let _1781___mcc_h7 = (_source25).cf41;
          let _1782_cf41 = _1781___mcc_h7;
          let _1783_cf40 = _1780___mcc_h6;
          let _1784_cf39 = _1779___mcc_h5;
          let _1785_cf38 = _1778___mcc_h4;
          let _1786_cf37 = _1777___mcc_h3;
          let _index314 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index314] = (_dafny.ZERO).minus(_1785_cf38);
          let _index315 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index315] = new BigNumber(879);
          (globalState).f2 = _dafny.Seq.Concat(_1774_v75, _dafny.Seq.of(_1786_cf37, (_this).f17, _module.__default.fm30((_1699_v15)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_1699_v15).length))], globalState)));
          let _1787_v78;
          let _nw257 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1787_v78 = _nw257;
          let _rhs282 = _1787_v78;
          let _rhs283 = p0;
          let _lhs232 = globalState;
          _1787_v78 = _rhs282;
          _lhs232.f7 = _rhs283;
          r0 = true;
        } else if (_source25.is_DC23) {
          let _1788___mcc_h8 = (_source25).cf42;
          let _1789___mcc_h9 = (_source25).cf43;
          let _1790___mcc_h10 = (_source25).cf44;
          let _1791___mcc_h11 = (_source25).cf45;
          let _1792___mcc_h12 = (_source25).cf46;
          let _1793_cf46 = _1792___mcc_h12;
          let _1794_cf45 = _1791___mcc_h11;
          let _1795_cf44 = _1790___mcc_h10;
          let _1796_cf43 = _1789___mcc_h9;
          let _1797_cf42 = _1788___mcc_h8;
          let _index316 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index316] = (((_this).f13) ? (true) : ((_this).f17));
          let _1798_v79;
          _1798_v79 = _module.D6.create_DC13(!((_this).f17), ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], new BigNumber(556), _1793_cf46, (_this).f17);
          _1798_v79 = _1798_v79;
          let _index317 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index317] = ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))];
          let _1799_v80;
          let _init41 = function (_1800_i5) {
            return ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))];
          };
          let _nw258 = Array((new BigNumber(4)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw258.length); _i0_41++) {
            _nw258[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1799_v80 = _nw258;
          let _1801_v81;
          _1801_v81 = _dafny.MultiSet.fromElements(_1799_v80);
          let _1802_v82;
          _1802_v82 = _1801_v81;
          let _1803_v83;
          _1803_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1801_v81,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("f"), _dafny.Seq.UnicodeFromString("uothuadbe"))).length));
          let _1804_v84;
          _1804_v84 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1805_v85;
          let _nw259 = new _module.C1();
          _nw259.__ctor(_1804_v84, _this.f12, (_this).f17);
          _1805_v85 = _nw259;
          let _1806_v86;
          _1806_v86 = _dafny.Seq.of(_1805_v85);
          _1803_v83 = (_1803_v83).update(_1802_v82, new BigNumber((_1806_v86).length));
        } else {
          let _1807___mcc_h13 = (_source25).cf36;
          let _1808_cf36 = _1807___mcc_h13;
          let _index318 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index318] = (((_this).f15).Union(_dafny.Set.fromElements(p0))).IsDisjointFrom((_this).f15);
          let _index319 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index319] = p0;
          let _index320 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index320] = p0;
          let _1809_v87;
          _1809_v87 = _dafny.MultiSet.fromElements(p0);
          let _1810_v88;
          _1810_v88 = _dafny.MultiSet.fromElements((_1699_v15)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length))], (((_1809_v87).contains((_dafny.ZERO).minus(p0))) ? ((_1809_v87).get((_dafny.ZERO).minus(p0))) : (p0)));
          let _1811_v89;
          _1811_v89 = _dafny.Set.fromElements((_1809_v87).IsProperSubsetOf(_1809_v87));
          let _1812_v90;
          _1812_v90 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements((_this).f13));
          let _1813_v91;
          _1813_v91 = _dafny.Seq.UnicodeFromString("kxckk");
          let _1814_v92;
          _1814_v92 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1815_v93;
          _1815_v93 = _dafny.Seq.of(p0);
          let _rhs284 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))];
          let _rhs285 = (((_1812_v90).contains(new BigNumber((_dafny.Seq.update(_1813_v91, _module.__default.safeIndex(p0, new BigNumber((_1813_v91).length)), _1814_v92)).length))) ? ((_1812_v90).get(new BigNumber((_dafny.Seq.update(_1813_v91, _module.__default.safeIndex(p0, new BigNumber((_1813_v91).length)), _1814_v92)).length))) : (_dafny.Set.fromElements(((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))])));
          let _rhs286 = ((_dafny.MultiSet.FromArray(_1815_v93)).Intersect((_1810_v88).update(new BigNumber(901), _module.__default.abs((_1699_v15)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length))])))).Difference(_1809_v87);
          r0 = _rhs284;
          _1811_v89 = _rhs285;
          _1810_v88 = _rhs286;
          let _index321 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length));
          (_1699_v15)[_index321] = _module.__default.fm0(_dafny.Seq.of((_this).f13), ((_this).f14)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f14).length))], true, (_1699_v15)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1699_v15).length))], globalState);
        }
      }
      r0 = ((_this).f13) || ((_this).f17);
      let _index322 = _module.__default.safeIndex(new BigNumber(701), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index322] = !((_this).f17);
      let _index323 = _module.__default.safeIndex(new BigNumber(701), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index323] = (_this).f17;
      let _index324 = _module.__default.safeIndex(new BigNumber(701), new BigNumber(((_this).f14).length));
      ((_this).f14)[_index324] = ((_this).f14)[_module.__default.safeIndex(new BigNumber(701), new BigNumber(((_this).f14).length))];
      r0 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(701), new BigNumber(((_this).f14).length))];
      return r0;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _1816_v0;
      _1816_v0 = _dafny.Seq.of((_this).f17);
      (globalState).f7 = _module.__default.fm0(_1816_v0, (_this).f17, (_this).f17, p0, globalState);
      let _1817_i0;
      _1817_i0 = _dafny.ZERO;
      L11: {
        while (!((_this).f17)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1817_i0)) {
              break L11;
            }
            _1817_i0 = (_1817_i0).plus(_dafny.ONE);
            let _1818_v1;
            _1818_v1 = _dafny.MultiSet.fromElements((_this).f17);
            r0 = ((_1818_v1).Union(_1818_v1)).IsSubsetOf(_1818_v1);
            (globalState).f7 = p0;
            let _1819_v2;
            let _nw260 = Array((new BigNumber(2)).toNumber()).fill([]);
            _1819_v2 = _nw260;
            let _1820_v3;
            let _nw261 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
            _1820_v3 = _nw261;
            let _index325 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1819_v2).length));
            (_1819_v2)[_index325] = _1820_v3;
            let _1821_v4;
            let _nw262 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
            _1821_v4 = _nw262;
            let _1822_v5;
            _1822_v5 = new _dafny.CodePoint('h'.codePointAt(0));
            let _index326 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1821_v4).length));
            (_1821_v4)[_index326] = (_this).fm9(_1822_v5, globalState);
            let _1823_v6;
            _1823_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
            let _1824_v7;
            _1824_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(globalState),(_dafny.MultiSet.fromElements(_1823_v6, _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17), _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f13))).IsSubsetOf(_dafny.MultiSet.fromElements(_1823_v6)));
            let _1825_v8;
            _1825_v8 = _module.D2.create_DC4(_1821_v4, (_this).f17, p0, (_this).f13, (_this).f14);
            let _index327 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1819_v2).length));
            let _index328 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1821_v4).length));
            let _rhs287 = (_this).f17;
            let _rhs288 = (((_1824_v7).contains((_this).f17)) ? ((_1824_v7).get((_this).f17)) : ((_this).f13));
            let _rhs289 = _1820_v3;
            let _rhs290 = (_1825_v8).dtor_cf8;
            let _rhs291 = (p0).multipliedBy(new BigNumber(2));
            let _lhs233 = _1819_v2;
            let _lhs234 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1819_v2).length));
            let _lhs235 = _1821_v4;
            let _lhs236 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1821_v4).length));
            let _lhs237 = globalState;
            r0 = _rhs287;
            r0 = _rhs288;
            _lhs233[_lhs234] = _rhs289;
            _lhs235[_lhs236] = _rhs290;
            _lhs237.f7 = _rhs291;
            let _index329 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1821_v4).length));
            (_1821_v4)[_index329] = (_1821_v4)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_1821_v4).length))];
          }
        }
      }
      let _1826_v9;
      _1826_v9 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1827_v10;
      _1827_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1826_v9,new _dafny.CodePoint('i'.codePointAt(0)));
      _1827_v10 = (_1827_v10).update(_1826_v9, _1826_v9);
      let _1828_v11;
      let _nw263 = Array((new BigNumber(20)).toNumber()).fill([]);
      _1828_v11 = _nw263;
      let _1829_v12;
      _1829_v12 = _1828_v11;
      let _1830_v13;
      _1830_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1829_v12),!(false));
      let _1831_v14;
      _1831_v14 = _dafny.Set.fromElements((_this).f17);
      _1830_v13 = (_1830_v13).update(_1828_v11, (_1831_v14).equals(_1831_v14));
      (globalState).f7 = (_this).fm9(_1826_v9, globalState);
      let _hi12 = p0;
      for (let _1832_i1 = p0; _1832_i1.isLessThan(_hi12); _1832_i1 = _1832_i1.plus(_dafny.ONE)) {
        let _1833_v15;
        _1833_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm11(globalState),_1832_i1);
        _1833_v15 = (_1833_v15).update((_this).f13, _1832_i1);
        (globalState).f7 = _1832_i1;
        let _1834_v16;
        _1834_v16 = _module.D7.create_DC16(_1832_i1);
        let _1835_v17;
        _1835_v17 = _module.D9.create_DC23(_module.__default.safeDivisionInt(p0, p0), p0, ((_1834_v16).dtor_cf29).minus(new BigNumber((_dafny.Set.fromElements(p0)).length)), _1832_i1, _module.__default.safeDivisionInt(_module.__default.fm0(_1816_v0, (_this).f17, !(false), p0, globalState), _1832_i1));
        let _1836_v18;
        let _nw264 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1836_v18 = _nw264;
        let _1837_v19;
        let _nw265 = Array((new BigNumber(27)).toNumber());
        _nw265[(_dafny.ZERO).toNumber()] = _1826_v9;
        _nw265[(_dafny.ONE).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(2)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw265[(new BigNumber(4)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
        _nw265[(new BigNumber(6)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(7)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(8)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(9)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(10)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(11)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(12)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(13)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(14)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(15)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(16)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(17)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(18)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(19)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw265[(new BigNumber(21)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(22)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(23)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(24)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(25)).toNumber()] = _1826_v9;
        _nw265[(new BigNumber(26)).toNumber()] = _1826_v9;
        _1837_v19 = _nw265;
        let _index330 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1836_v18).length));
        (_1836_v18)[_index330] = (_module.D13.create_DC33(_1837_v19)).dtor_cf70;
        let _1838_v20;
        _1838_v20 = _dafny.MultiSet.fromElements(true, (_this).f17, (_this).f13);
        let _1839_v21;
        _1839_v21 = _dafny.Set.fromElements(_1838_v20, _1838_v20);
        let _1840_v22;
        _1840_v22 = _module.D6.create_DC11((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs(_1832_i1)));
        let _1841_v23;
        _1841_v23 = _module.D6.create_DC14(_1840_v22);
        let _index331 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1836_v18).length));
        let _rhs292 = _1835_v17;
        let _rhs293 = _1837_v19;
        let _rhs294 = _1826_v9;
        let _rhs295 = (_1839_v21).Difference(_1839_v21);
        let _rhs296 = _module.__default.fm30((((_this).f17) ? (new BigNumber((_module.__default.fm32(_module.D6.create_DC14(_1840_v22), globalState)).length)) : (_1832_i1)), globalState);
        let _lhs238 = _1836_v18;
        let _lhs239 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1836_v18).length));
        _1835_v17 = _rhs292;
        _lhs238[_lhs239] = _rhs293;
        _1826_v9 = _rhs294;
        _1839_v21 = _rhs295;
        r0 = _rhs296;
        let _1842_v24;
        let _nw266 = new _module.C4();
        _nw266.__ctor();
        _1842_v24 = _nw266;
      }
      r0 = ((_this).f15).IsProperSubsetOf((_this).f15);
      return r0;
    }
    m9(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1843_v0;
      _1843_v0 = _dafny.Seq.of((_this).f17);
      let _1844_v1;
      _1844_v1 = _dafny.Seq.of(p1, p1);
      let _1845_v2;
      _1845_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_1844_v1).length));
      let _1846_v3;
      _1846_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1845_v2).update((_this).f13, p1));
      let _1847_i0;
      _1847_i0 = _dafny.ZERO;
      L12: {
        while ((new BigNumber((_dafny.Seq.Concat(_1843_v0, _1843_v0)).length)).isEqualTo(new BigNumber((_1846_v3).length))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1847_i0)) {
              break L12;
            }
            _1847_i0 = (_1847_i0).plus(_dafny.ONE);
            let _1848_v4;
            let _nw267 = new _module.C5();
            _nw267.__ctor(!((_this).f13), (_this).f14, (_this).f15, _module.D1.create_DC2(p1, new BigNumber((_dafny.MultiSet.fromElements(p3)).cardinality()), (_this).f17), p2);
            _1848_v4 = _nw267;
            _1848_v4 = _1848_v4;
            let _1849_v5;
            let _nw268 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1849_v5 = _nw268;
            let _1850_v6;
            _1850_v6 = _dafny.Map.Empty.slice().updateUnsafe((_1848_v4).f13,_1849_v5);
            (globalState).f7 = (p1).multipliedBy((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1850_v6).length), p1)));
            let _1851_v7;
            _1851_v7 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("thmbvav"));
            let _1852_v8;
            _1852_v8 = _dafny.Seq.UnicodeFromString("vhrtlm");
            _1851_v7 = _dafny.Seq.Concat(_1851_v7, _dafny.Seq.update(_1851_v7, _module.__default.safeIndex(p1, new BigNumber((_1851_v7).length)), _1852_v8));
            r1 = (p1).isEqualTo(p1);
          }
        }
      }
      if (!(function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-360)), ((_1853_v0) => function (_1854_i1) {
          return new BigNumber((_1853_v0).length);
        })(_1843_v0))).Elements) {
          let _1855_v9 = _compr_53;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-360)), ((_1856_v0) => function (_1854_i1) {
            return new BigNumber((_1856_v0).length);
          })(_1843_v0)), _1855_v9)) {
            _coll53.push([_module.__default.safeDivisionInt(_1855_v9, p1),p0]);
          }
        }
        return _coll53;
      }()).contains(p1)) {
        r1 = false;
        if (!(p3) || ((_this).f17)) {
          let _1857_v10;
          let _init42 = ((_1858_p1) => function (_1859_i2) {
            return _module.__default.safeDivisionInt(_1859_i2, _1858_p1);
          })(p1);
          let _nw269 = Array((new BigNumber(24)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw269.length); _i0_42++) {
            _nw269[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1857_v10 = _nw269;
          let _index332 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1857_v10).length));
          (_1857_v10)[_index332] = p1;
          let _index333 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1857_v10).length));
          (_1857_v10)[_index333] = (p1).minus(_module.__default.safeDivisionInt(p1, p1));
          let _1860_v11;
          _1860_v11 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1861_v12;
          _1861_v12 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), _1860_v11, new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)));
          _1861_v12 = _1861_v12;
          r1 = false;
          let _1862_v13;
          let _init43 = ((_1863_p1) => function (_1864_i3) {
            return _module.D7.create_DC15(_dafny.Seq.of(_1863_p1));
          })(p1);
          let _nw270 = Array((new BigNumber(24)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw270.length); _i0_43++) {
            _nw270[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1862_v13 = _nw270;
          let _1865_v14;
          _1865_v14 = _dafny.Set.fromElements(_1862_v13, _1862_v13, _1862_v13);
          let _index334 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1857_v10).length));
          let _rhs297 = (_1857_v10)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_1857_v10).length))];
          let _rhs298 = !(_1865_v14).equals(_1865_v14);
          let _lhs240 = _1857_v10;
          let _lhs241 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1857_v10).length));
          _lhs240[_lhs241] = _rhs297;
          r1 = _rhs298;
          let _1866_v15;
          let _init44 = ((_1867_p1) => function (_1868_i4) {
            return (new BigNumber(545)).isLessThanOrEqualTo(_1867_p1);
          })(p1);
          let _nw271 = Array((new BigNumber(3)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw271.length); _i0_44++) {
            _nw271[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1866_v15 = _nw271;
          let _1869_v16;
          _1869_v16 = _dafny.Seq.UnicodeFromString("fxdirlqrk");
          let _1870_v17;
          let _nw272 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1870_v17 = _nw272;
          let _index335 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1870_v17).length));
          (_1870_v17)[_index335] = _1866_v15;
          let _index336 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1870_v17).length));
          let _rhs299 = (_module.__default.safeDivisionInt(p1, p1)).isEqualTo(p1);
          let _rhs300 = _1857_v10;
          let _rhs301 = _1866_v15;
          let _rhs302 = _1869_v16;
          let _rhs303 = (_this).f14;
          let _lhs242 = _1870_v17;
          let _lhs243 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1870_v17).length));
          r1 = _rhs299;
          _1857_v10 = _rhs300;
          _1866_v15 = _rhs301;
          _1869_v16 = _rhs302;
          _lhs242[_lhs243] = _rhs303;
        } else {
          r0 = p1;
          let _1871_v18;
          _1871_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p1, p1),(p1).plus(p1));
          _1871_v18 = _1871_v18;
          let _1872_v19;
          _1872_v19 = _dafny.MultiSet.fromElements(p2, (_this).f13);
          let _1873_v20;
          _1873_v20 = _module.D16.create_DC42(_1872_v19);
          _1873_v20 = _1873_v20;
          r1 = (_this).f13;
          let _1874_v21;
          _1874_v21 = _dafny.Seq.UnicodeFromString("olfi");
          let _1875_v22;
          _1875_v22 = _dafny.Seq.of(_1874_v21, _1874_v21);
          let _1876_v23;
          _1876_v23 = _module.D6.create_DC12((_this).f13, _1874_v21, p1, p3);
          let _1877_v24;
          let _nw273 = Array((new BigNumber(16)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = _1875_v22;
          _nw273[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1874_v21), _1875_v22);
          _nw273[(new BigNumber(2)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(3)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(4)).toNumber()] = _module.__default.fm45(_1876_v23, p1, globalState);
          _nw273[(new BigNumber(5)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(6)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(7)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(8)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1875_v22, _1875_v22);
          _nw273[(new BigNumber(10)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(11)).toNumber()] = _1875_v22;
          _nw273[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-12)), function (_1878_i5) {
            return _dafny.Seq.UnicodeFromString("msjf");
          });
          _nw273[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1875_v22, _module.__default.safeIndex(new BigNumber((_1874_v21).length), new BigNumber((_1875_v22).length)), _dafny.Seq.UnicodeFromString("vqwt")), _1875_v22);
          _nw273[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1874_v21), _dafny.Seq.Create(_module.__default.abs(new BigNumber(721)), function (_1879_i6) {
            return _dafny.Seq.UnicodeFromString("qpiiluosn");
          }));
          _nw273[(new BigNumber(15)).toNumber()] = _module.__default.fm45(_1876_v23, new BigNumber((_dafny.Seq.of(_1871_v18, _module.__default.fm24(p1, globalState))).length), globalState);
          _1877_v24 = _nw273;
          let _index337 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1877_v24).length));
          (_1877_v24)[_index337] = _1875_v22;
          let _index338 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1877_v24).length));
          (_1877_v24)[_index338] = _1875_v22;
        }
        let _1880_v25;
        _1880_v25 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1881_v26;
        let _nw274 = new _module.C1();
        _nw274.__ctor(_1880_v25, _this.f12, ((_this).fm9(_1880_v25, globalState)).isLessThanOrEqualTo(p1));
        _1881_v26 = _nw274;
        r1 = (_this).f17;
        let _1882_v27;
        _1882_v27 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm17(globalState));
        let _1883_v28;
        _1883_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f17);
        let _1884_v29;
        _1884_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1881_v26.f23,p3);
        let _1885_v30;
        _1885_v30 = _module.D13.create_DC34(p1, _module.__default.fm65(p3, _dafny.Seq.of(new BigNumber((_1884_v29).length)), _1880_v25, !(p2), globalState));
        let _pat_let_tv50 = p1;
        let _1886_v31;
        let _nw275 = Array((new BigNumber(29)).toNumber());
        _nw275[(_dafny.ZERO).toNumber()] = new BigNumber(105);
        _nw275[(_dafny.ONE).toNumber()] = p1;
        _nw275[(new BigNumber(2)).toNumber()] = p1;
        _nw275[(new BigNumber(3)).toNumber()] = p1;
        _nw275[(new BigNumber(4)).toNumber()] = p1;
        _nw275[(new BigNumber(5)).toNumber()] = p1;
        _nw275[(new BigNumber(6)).toNumber()] = p1;
        _nw275[(new BigNumber(7)).toNumber()] = p1;
        _nw275[(new BigNumber(8)).toNumber()] = p1;
        _nw275[(new BigNumber(9)).toNumber()] = (new BigNumber((_1882_v27).length)).multipliedBy(p1);
        _nw275[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((p1).plus(p1));
        _nw275[(new BigNumber(11)).toNumber()] = p1;
        _nw275[(new BigNumber(12)).toNumber()] = new BigNumber(916);
        _nw275[(new BigNumber(13)).toNumber()] = p1;
        _nw275[(new BigNumber(14)).toNumber()] = new BigNumber((_1883_v28).length);
        _nw275[(new BigNumber(15)).toNumber()] = p1;
        _nw275[(new BigNumber(16)).toNumber()] = new BigNumber(99);
        _nw275[(new BigNumber(17)).toNumber()] = p1;
        _nw275[(new BigNumber(18)).toNumber()] = p1;
        _nw275[(new BigNumber(19)).toNumber()] = p1;
        _nw275[(new BigNumber(20)).toNumber()] = p1;
        _nw275[(new BigNumber(21)).toNumber()] = new BigNumber(-859);
        _nw275[(new BigNumber(22)).toNumber()] = p1;
        _nw275[(new BigNumber(23)).toNumber()] = p1;
        _nw275[(new BigNumber(24)).toNumber()] = new BigNumber(-842);
        _nw275[(new BigNumber(25)).toNumber()] = _module.__default.fm0(_1843_v0, p3, !(p3), p1, globalState);
        _nw275[(new BigNumber(26)).toNumber()] = p1;
        _nw275[(new BigNumber(27)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
        _nw275[(new BigNumber(28)).toNumber()] = (function (_pat_let59_0) {
          return function (_1887_dt__update__tmp_h0) {
            return function (_pat_let60_0) {
              return function (_1888_dt__update_hcf71_h0) {
                return _module.D13.create_DC34(_1888_dt__update_hcf71_h0, (_1887_dt__update__tmp_h0).dtor_cf72);
              }(_pat_let60_0);
            }(_pat_let_tv50);
          }(_pat_let59_0);
        }(_1885_v30)).dtor_cf71;
        _1886_v31 = _nw275;
        _1886_v31 = _1886_v31;
      } else {
        r1 = (_this).f17;
        r1 = (_this).f17;
        let _1889_v32;
        _1889_v32 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1890_v33;
        _1890_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8((_this).f17, _this.f16, globalState),_1889_v32);
        let _1891_v34;
        _1891_v34 = _dafny.MultiSet.fromElements(p1, p1);
        _1890_v33 = (_1890_v33).update(new BigNumber((_1891_v34).cardinality()), _1889_v32);
        let _1892_v35;
        _1892_v35 = _dafny.Seq.UnicodeFromString("vquk");
        let _1893_v36;
        _1893_v36 = _dafny.Seq.of(_1892_v35);
        let _1894_v37;
        _1894_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
        let _1895_v38;
        let _nw276 = Array((new BigNumber(8)).toNumber());
        _nw276[(_dafny.ZERO).toNumber()] = new BigNumber(466);
        _nw276[(_dafny.ONE).toNumber()] = new BigNumber((_1894_v37).length);
        _nw276[(new BigNumber(2)).toNumber()] = p1;
        _nw276[(new BigNumber(3)).toNumber()] = p1;
        _nw276[(new BigNumber(4)).toNumber()] = p1;
        _nw276[(new BigNumber(5)).toNumber()] = p1;
        _nw276[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1892_v35).length));
        _nw276[(new BigNumber(7)).toNumber()] = new BigNumber(-758);
        _1895_v38 = _nw276;
        let _1896_v39;
        _1896_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1893_v36)[_module.__default.safeIndex(p1, new BigNumber((_1893_v36).length))],_dafny.Map.Empty.slice().updateUnsafe(_1895_v38,p1));
        _1896_v39 = (_1896_v39).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(937)), ((_1897_v32) => function (_1898_i7) {
          return _1897_v32;
        })(_1889_v32)), _dafny.Map.Empty.slice().updateUnsafe(_1895_v38,new BigNumber((_1892_v35).length)));
        let _pat_let_tv51 = p1;
        let _1899_v40;
        let _nw277 = new _module.C2();
        _nw277.__ctor(_module.__default.safeDivisionInt(p1, new BigNumber(18)), (p3) || (!(p0)), _this.f16, p2, (_this).f14, ((_this).f15).Intersect((_this).f15), function (_pat_let61_0) {
          return function (_1900_dt__update__tmp_h1) {
            return function (_pat_let62_0) {
              return function (_1901_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_1901_dt__update_hcf2_h0, (_1900_dt__update__tmp_h1).dtor_cf3, (_1900_dt__update__tmp_h1).dtor_cf4);
              }(_pat_let62_0);
            }(_pat_let_tv51);
          }(_pat_let61_0);
        }(_module.D1.create_DC2(p1, p1, !((_this).f13))), (_this).f13);
        _1899_v40 = _nw277;
      }
      r1 = (_this).f17;
      let _1902_v41;
      _1902_v41 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1903_v42;
      _1903_v42 = _dafny.Set.fromElements(_module.D3.create_DC5(_1902_v41));
      let _1904_v43;
      _1904_v43 = _dafny.MultiSet.fromElements(p1, p1);
      let _pat_let_tv52 = _1903_v42;
      let _pat_let_tv53 = _1902_v41;
      let _pat_let_tv54 = _1902_v41;
      let _pat_let_tv55 = _1902_v41;
      let _pat_let_tv56 = _1902_v41;
      let _pat_let_tv57 = _1902_v41;
      let _pat_let_tv58 = _1903_v42;
      let _rhs304 = p1;
      let _rhs305 = new BigNumber(209);
      let _rhs306 = function (_source26) {
        if (_source26.is_DC12) {
          let _1905___mcc_h0 = (_source26).cf18;
          let _1906___mcc_h1 = (_source26).cf19;
          let _1907___mcc_h2 = (_source26).cf20;
          let _1908___mcc_h3 = (_source26).cf21;
          let _1909_cf21 = _1908___mcc_h3;
          let _1910_cf20 = _1907___mcc_h2;
          let _1911_cf19 = _1906___mcc_h1;
          let _1912_cf18 = _1905___mcc_h0;
          return _pat_let_tv52;
        } else if (_source26.is_DC13) {
          let _1913___mcc_h4 = (_source26).cf22;
          let _1914___mcc_h5 = (_source26).cf23;
          let _1915___mcc_h6 = (_source26).cf24;
          let _1916___mcc_h7 = (_source26).cf25;
          let _1917___mcc_h8 = (_source26).cf26;
          let _1918_cf26 = _1917___mcc_h8;
          let _1919_cf25 = _1916___mcc_h7;
          let _1920_cf24 = _1915___mcc_h6;
          let _1921_cf23 = _1914___mcc_h5;
          let _1922_cf22 = _1913___mcc_h4;
          return _dafny.Set.fromElements(_module.D3.create_DC5(_pat_let_tv53), _module.D3.create_DC5(_pat_let_tv54));
        } else if (_source26.is_DC11) {
          let _1923___mcc_h9 = (_source26).cf17;
          let _1924_cf17 = _1923___mcc_h9;
          return _dafny.Set.fromElements(_module.D3.create_DC5(_pat_let_tv55), _module.D3.create_DC5(_pat_let_tv56), _module.D3.create_DC5(_pat_let_tv57));
        } else {
          let _1925___mcc_h10 = (_source26).cf27;
          let _1926_cf27 = _1925___mcc_h10;
          return _pat_let_tv58;
        }
      }(_module.D6.create_DC11(_1904_v43));
      let _rhs307 = p0;
      r0 = _rhs304;
      r0 = _rhs305;
      _1903_v42 = _rhs306;
      r1 = _rhs307;
      r1 = p3;
      if ((p3) || (!((_this).f17))) {
        let _1927_v44;
        _1927_v44 = _dafny.Seq.UnicodeFromString("wypej");
        let _1928_v45;
        let _nw278 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _1928_v45 = _nw278;
        let _index339 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length));
        (_1928_v45)[_index339] = new BigNumber((_dafny.MultiSet.fromElements((_this).f13, p3)).cardinality());
        let _index340 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length));
        let _rhs308 = _module.__default.safeDivisionInt((((_this).f17) ? (p1) : (new BigNumber(36))), p1);
        let _rhs309 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1927_v44, _1927_v44), _1927_v44);
        let _rhs310 = _1902_v41;
        let _rhs311 = (p1).isEqualTo(p1);
        let _rhs312 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _1843_v0), _1843_v0)).length);
        let _lhs244 = _1928_v45;
        let _lhs245 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length));
        r0 = _rhs308;
        _1927_v44 = _rhs309;
        _1902_v41 = _rhs310;
        r1 = _rhs311;
        _lhs244[_lhs245] = _rhs312;
        let _1929_v46;
        let _nw279 = new _module.C1();
        _nw279.__ctor(_1902_v41, _this.f12, p2);
        _1929_v46 = _nw279;
        let _1930_v47;
        _1930_v47 = _dafny.Seq.of(_1929_v46);
        let _1931_v48;
        let _nw280 = Array((new BigNumber(13)).toNumber());
        _nw280[(_dafny.ZERO).toNumber()] = _1929_v46;
        _nw280[(_dafny.ONE).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(2)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(3)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(4)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(5)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(6)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(7)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(8)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(9)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(10)).toNumber()] = (_1930_v47)[_module.__default.safeIndex(p1, new BigNumber((_1930_v47).length))];
        _nw280[(new BigNumber(11)).toNumber()] = _1929_v46;
        _nw280[(new BigNumber(12)).toNumber()] = _1929_v46;
        _1931_v48 = _nw280;
        let _1932_v49;
        _1932_v49 = _module.D17.create_DC44(_1931_v48);
        let _1933_v50;
        _1933_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_dafny.Seq.of(p3), (_this).f13, p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(527)), ((_1934_v41) => function (_1935_i8) {
          return _1934_v41;
        })(_1902_v41))).length), globalState),_1931_v48);
        let _1936_v51;
        let _nw281 = Array((new BigNumber(14)).toNumber());
        _nw281[(_dafny.ZERO).toNumber()] = _1931_v48;
        _nw281[(_dafny.ONE).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(2)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(3)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(4)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(5)).toNumber()] = (_1932_v49).dtor_cf87;
        _nw281[(new BigNumber(6)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(7)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(8)).toNumber()] = (((_1933_v50).contains((_1928_v45)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length))])) ? ((_1933_v50).get((_1928_v45)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length))])) : (_1931_v48));
        _nw281[(new BigNumber(9)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(10)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(11)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(12)).toNumber()] = _1931_v48;
        _nw281[(new BigNumber(13)).toNumber()] = _1931_v48;
        _1936_v51 = _nw281;
        let _index341 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1936_v51).length));
        (_1936_v51)[_index341] = _1931_v48;
        let _1937_v52;
        _1937_v52 = _dafny.Seq.of(((false) ? (_1931_v48) : (_1931_v48)), _1931_v48, _1931_v48, _1931_v48, _1931_v48);
        let _index342 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1936_v51).length));
        (_1936_v51)[_index342] = (_1937_v52)[_module.__default.safeIndex(p1, new BigNumber((_1937_v52).length))];
        let _index343 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1928_v45).length));
        (_1928_v45)[_index343] = (_dafny.ZERO).minus(p1);
        r1 = (_this).f17;
        (globalState).f7 = p1;
      } else {
        let _1938_v53;
        _1938_v53 = _dafny.Seq.UnicodeFromString("uvv");
        (globalState).f7 = new BigNumber((((p0) ? (_1938_v53) : (_1938_v53))).length);
        let _1939_v54;
        _1939_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f17);
        let _1940_v55;
        let _nw282 = new _module.C2();
        _nw282.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("bndxiw"))).length), p0, _this.f16, (((_1939_v54).contains(p1)) ? ((_1939_v54).get(p1)) : ((_this).f13)), (_this).f14, (_this).f15, _this.f12, (_this).f17);
        _1940_v55 = _nw282;
        _1940_v55 = _1940_v55;
        let _1941_v56;
        _1941_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1940_v55,(_this).f17);
        let _1942_v57;
        _1942_v57 = _module.D6.create_DC12(!((_1940_v55).f21).isEqualTo((_1940_v55).f21), _dafny.Seq.Concat(_1938_v53, _dafny.Seq.update(_1938_v53, _module.__default.safeIndex((_1940_v55).f21, new BigNumber((_1938_v53).length)), _1902_v41)), new BigNumber(((_1941_v56).Merge(_1941_v56)).length), p3);
        _1942_v57 = _module.__default.fm39(!(p3) || (false), p3, globalState);
        if (p0) {
          let _1943_v58;
          let _init45 = ((_1944_v41) => function (_1945_i9) {
            return _1944_v41;
          })(_1902_v41);
          let _nw283 = Array((new BigNumber(17)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw283.length); _i0_45++) {
            _nw283[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1943_v58 = _nw283;
          let _1946_v59;
          _1946_v59 = _module.D13.create_DC33(_1943_v58);
          let _1947_v60;
          _1947_v60 = _module.D13.create_DC35(_1946_v59);
          let _1948_v61;
          _1948_v61 = _module.D13.create_DC35((_1947_v60).dtor_cf73);
          _1948_v61 = _1948_v61;
          let _index344 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index344] = (_this).f13;
          let _1949_v62;
          _1949_v62 = _dafny.MultiSet.fromElements((_1940_v55).f22, !((_this).f13), !((_1940_v55).f22));
          let _1950_v63;
          _1950_v63 = _module.D6.create_DC13(p2, p2, (_1940_v55).f21, new BigNumber((_1949_v62).cardinality()), !(p0));
          let _index345 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f14).length));
          let _rhs313 = (_1950_v63).dtor_cf23;
          let _rhs314 = _1940_v55;
          let _rhs315 = (_1940_v55).f22;
          let _lhs246 = (_this).f14;
          let _lhs247 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f14).length));
          r1 = _rhs313;
          _1940_v55 = _rhs314;
          _lhs246[_lhs247] = _rhs315;
          let _1951_v64;
          let _init46 = ((_1952_v55) => function (_1953_i10) {
            return (_1953_i10).plus(new BigNumber((_dafny.Seq.of((_1952_v55).f22)).length));
          })(_1940_v55);
          let _nw284 = Array((new BigNumber(6)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw284.length); _i0_46++) {
            _nw284[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1951_v64 = _nw284;
          _1951_v64 = _1951_v64;
          let _1954_v65;
          _1954_v65 = _dafny.Map.Empty.slice().updateUnsafe((_1940_v55).f21,(_1940_v55).f21);
          let _1955_v66;
          _1955_v66 = _module.D11.create_DC27((_this).f13, _1902_v41, (((_1954_v65).contains((_1940_v55).f21)) ? ((_1954_v65).get((_1940_v55).f21)) : ((_dafny.ZERO).minus((_1940_v55).f21))), (_1940_v55).f21, _1844_v1);
          _1955_v66 = _1955_v66;
          let _1956_v67;
          let _nw285 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1956_v67 = _nw285;
          _1956_v67 = _1956_v67;
        } else {
          let _index346 = _module.__default.safeIndex(new BigNumber(678), new BigNumber(((_this).f14).length));
          ((_this).f14)[_index346] = false;
          let _1957_v68;
          let _nw286 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1957_v68 = _nw286;
          let _1958_v69;
          let _nw287 = Array((new BigNumber(15)).toNumber());
          _nw287[(_dafny.ZERO).toNumber()] = _1904_v43;
          _nw287[(_dafny.ONE).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray(_1844_v1)).Intersect(_1904_v43);
          _nw287[(new BigNumber(3)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(4)).toNumber()] = (_1904_v43).Intersect(_module.__default.fm19(globalState));
          _nw287[(new BigNumber(5)).toNumber()] = (_1904_v43).Union(_1904_v43);
          _nw287[(new BigNumber(6)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(7)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(8)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(9)).toNumber()] = (_1904_v43).Union(_1904_v43);
          _nw287[(new BigNumber(10)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(11)).toNumber()] = (_1904_v43).Union(_dafny.MultiSet.FromArray(_module.__default.fm27(globalState)));
          _nw287[(new BigNumber(12)).toNumber()] = _1904_v43;
          _nw287[(new BigNumber(13)).toNumber()] = (_1904_v43).update((_1940_v55).f21, _module.__default.abs((_1940_v55).f21));
          _nw287[(new BigNumber(14)).toNumber()] = (_1904_v43).update((_dafny.ZERO).minus(p1), _module.__default.abs((_dafny.ZERO).minus((_1940_v55).f21)));
          _1958_v69 = _nw287;
          let _index347 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1958_v69).length));
          (_1958_v69)[_index347] = (_module.__default.fm56(p1, (_this).f15, globalState)).Intersect(_1904_v43);
          let _1959_v70;
          _1959_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1940_v55).f21,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_1843_v0)).cardinality())));
          let _index348 = _module.__default.safeIndex(new BigNumber(678), new BigNumber(((_this).f14).length));
          let _index349 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1958_v69).length));
          let _rhs316 = _module.__default.safeDivisionInt(((_1940_v55).f21).multipliedBy((_1940_v55).f21), new BigNumber(721));
          let _rhs317 = _module.__default.fm30((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1959_v70).update((_1940_v55).f21, new BigNumber(217))).length),p1)).length)).multipliedBy((_1940_v55).f21), globalState);
          let _rhs318 = (_this).f14;
          let _rhs319 = (_1904_v43).Union(_1904_v43);
          let _rhs320 = _1902_v41;
          let _lhs248 = globalState;
          let _lhs249 = (_this).f14;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(678), new BigNumber(((_this).f14).length));
          let _lhs251 = _1958_v69;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1958_v69).length));
          _lhs248.f7 = _rhs316;
          _lhs249[_lhs250] = _rhs317;
          _1957_v68 = _rhs318;
          _lhs251[_lhs252] = _rhs319;
          _1902_v41 = _rhs320;
          r1 = (((_1940_v55).f21).multipliedBy((_dafny.ZERO).minus((_1940_v55).fm9(_module.__default.fm18((_1940_v55).f21, _1938_v53, p2, _1845_v2, globalState), globalState)))).isLessThan(p1);
          r0 = ((p2) ? ((_dafny.ZERO).minus((_1940_v55).f21)) : ((_1940_v55).f21));
          (globalState).f7 = (_1940_v55).f21;
          let _1960_v71;
          let _init47 = ((_1961_v55) => function (_1962_i11) {
            return _module.__default.safeModuloInt(_1962_i11, (_1961_v55).f21);
          })(_1940_v55);
          let _nw288 = Array((new BigNumber(20)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw288.length); _i0_47++) {
            _nw288[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1960_v71 = _nw288;
          let _index350 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1960_v71).length));
          (_1960_v71)[_index350] = p1;
          let _1963_v72;
          _1963_v72 = _dafny.Map.Empty.slice().updateUnsafe(!(((_this).f14)[_module.__default.safeIndex(new BigNumber(678), new BigNumber(((_this).f14).length))]),_1845_v2);
          let _index351 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1960_v71).length));
          let _rhs321 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).fm9(_1902_v41, globalState), (_1940_v55).f21)), (p1).multipliedBy((_1940_v55).f21));
          let _rhs322 = _1957_v68;
          let _rhs323 = (_1940_v55).f21;
          let _rhs324 = ((((_this).f13) ? (_1963_v72) : (_1963_v72))).update((_1940_v55).f22, (_1845_v2).update((_this).f17, new BigNumber((_1844_v1).length)));
          let _rhs325 = (_1940_v55).f22;
          let _lhs253 = globalState;
          let _lhs254 = _1960_v71;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_1960_v71).length));
          _lhs253.f7 = _rhs321;
          _1957_v68 = _rhs322;
          _lhs254[_lhs255] = _rhs323;
          _1963_v72 = _rhs324;
          r1 = _rhs325;
        }
        (globalState).f7 = (_module.__default.safeDivisionInt((((_1845_v2).contains(p2)) ? ((_1845_v2).get(p2)) : (p1)), (_1940_v55).f21)).minus(p1);
      }
      r0 = p1;
      r1 = !(new BigNumber((_module.__default.fm56(new BigNumber(414), (_this).f15, globalState)).cardinality())).isEqualTo(new BigNumber(756));
      return [r0, r1];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      if (!(!(true))) {
        return (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)))).length))).length))).multipliedBy(new BigNumber((_dafny.Seq.of(true)).length)));
      } else {
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("pkq")).length))).length)).multipliedBy(new BigNumber(959));
      }
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("t"), _dafny.Seq.UnicodeFromString("sgoyyjn"))).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-407)), function (_1964_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)));
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1965_v0;
      let _init48 = ((_1966_p3) => function (_1967_i0) {
        return _module.__default.safeDivisionInt(_1967_i0, _1966_p3);
      })(p3);
      let _nw289 = Array((new BigNumber(2)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw289.length); _i0_48++) {
        _nw289[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1965_v0 = _nw289;
      let _1968_v1;
      _1968_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1965_v0,p0);
      _1968_v1 = (_1968_v1).update(_1965_v0, p0);
      let _1969_i1;
      _1969_i1 = _dafny.ZERO;
      L13: {
        while (!(p0)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1969_i1)) {
              break L13;
            }
            _1969_i1 = (_1969_i1).plus(_dafny.ONE);
            let _1970_v2;
            _1970_v2 = new _dafny.CodePoint('j'.codePointAt(0));
            let _1971_v3;
            _1971_v3 = _dafny.Seq.UnicodeFromString("yqxdb");
            _1970_v2 = (_1971_v3)[_module.__default.safeIndex(p2, new BigNumber((_1971_v3).length))];
            let _1972_v4;
            let _nw290 = Array((new BigNumber(3)).toNumber()).fill(false);
            _1972_v4 = _nw290;
            let _1973_v5;
            _1973_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
            let _1974_v6;
            _1974_v6 = _dafny.Set.fromElements(p0, p0);
            let _index352 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1972_v4).length));
            (_1972_v4)[_index352] = (_this).fm7(p2, _1973_v5, _1971_v3, _1974_v6, globalState);
            let _1975_v7;
            _1975_v7 = _module.D2.create_DC4(_1965_v0, p0, p3, p0, _1972_v4);
            let _index353 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1972_v4).length));
            (_1972_v4)[_index353] = (_1975_v7).dtor_cf9;
            _1971_v3 = _1971_v3;
            let _index354 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1965_v0).length));
            (_1965_v0)[_index354] = p3;
            let _index355 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1965_v0).length));
            (_1965_v0)[_index355] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(901), new BigNumber((_dafny.Seq.UnicodeFromString("vtl")).length)), p3);
          }
        }
      }
      let _1976_v8;
      let _nw291 = Array((new BigNumber(14)).toNumber()).fill(false);
      _1976_v8 = _nw291;
      let _1977_v9;
      _1977_v9 = _module.D2.create_DC4(_1965_v0, !(p0), p3, p0, _1976_v8);
      let _source27 = _1977_v9;
      if (_source27.is_DC4) {
        let _1978___mcc_h0 = (_source27).cf6;
        let _1979___mcc_h1 = (_source27).cf7;
        let _1980___mcc_h2 = (_source27).cf8;
        let _1981___mcc_h3 = (_source27).cf9;
        let _1982___mcc_h4 = (_source27).cf10;
        let _1983_cf10 = _1982___mcc_h4;
        let _1984_cf9 = _1981___mcc_h3;
        let _1985_cf8 = _1980___mcc_h2;
        let _1986_cf7 = _1979___mcc_h1;
        let _1987_cf6 = _1978___mcc_h0;
        _1985_cf8 = p3;
        let _1988_v10;
        _1988_v10 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1989_v11;
        let _out14;
        _out14 = _module.__default.m0(_1988_v10, p0, globalState);
        _1989_v11 = _out14;
        let _1990_v12;
        _1990_v12 = _dafny.Seq.UnicodeFromString("c");
        let _1991_v13;
        _1991_v13 = _module.D1.create_DC1(_1990_v12);
        let _1992_v14;
        _1992_v14 = _dafny.Set.fromElements(p2);
        let _1993_v15;
        _1993_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1986_cf7,p3);
        let _1994_v16;
        _1994_v16 = _module.D1.create_DC2(new BigNumber((_1993_v15).length), _1985_cf8, false);
        let _pat_let_tv59 = _1990_v12;
        let _1995_v17;
        let _nw292 = new _module.C6();
        _nw292.__ctor(function (_pat_let63_0) {
          return function (_1996_dt__update__tmp_h0) {
            return function (_pat_let64_0) {
              return function (_1997_dt__update_hcf1_h0) {
                return _module.D1.create_DC1(_1997_dt__update_hcf1_h0);
              }(_pat_let64_0);
            }(_pat_let_tv59);
          }(_pat_let63_0);
        }(_1991_v13), _1984_cf9, _1976_v8, _1992_v14, _1994_v16, true);
        _1995_v17 = _nw292;
        let _index356 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1987_cf6).length));
        (_1987_cf6)[_index356] = _1985_cf8;
        let _1998_v18;
        _1998_v18 = _dafny.Seq.of(p1);
        let _index357 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1987_cf6).length));
        (_1987_cf6)[_index357] = _module.__default.fm0((_1998_v18)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1998_v18).length))], _1984_cf9, _1986_cf7, (_dafny.ZERO).minus(p2), globalState);
      } else {
        let _1999___mcc_h5 = (_source27).cf5;
        let _2000_cf5 = _1999___mcc_h5;
        let _2001_v19;
        _2001_v19 = _dafny.Seq.UnicodeFromString("pliy");
        let _2002_v20;
        _2002_v20 = _dafny.MultiSet.fromElements(_2001_v19, _2001_v19);
        let _2003_v21;
        _2003_v21 = _dafny.MultiSet.fromElements(p0);
        let _2004_v22;
        _2004_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2002_v20,_2003_v21);
        let _2005_v23;
        _2005_v23 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2006_v24;
        _2006_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v23,true);
        let _2007_v25;
        _2007_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2006_v24).length),p0);
        let _2008_v26;
        _2008_v26 = _dafny.Set.fromElements(new BigNumber((_2007_v25).length), p3);
        let _2009_v27;
        let _nw293 = new _module.C5();
        _nw293.__ctor(p0, _2000_cf5, _2008_v26, _module.D1.create_DC2(new BigNumber(-499), new BigNumber((_dafny.Set.fromElements(_2005_v23)).length), p0), p0);
        _2009_v27 = _nw293;
        let _2010_v28;
        _2010_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2009_v27);
        _2004_v22 = (_2004_v22).update(_module.__default.fm66(p3, globalState), (_2003_v21).update(p0, _module.__default.abs(new BigNumber((_2010_v28).length))));
        let _2011_v29;
        let _nw294 = Array((new BigNumber(13)).toNumber());
        _2011_v29 = _nw294;
        let _2012_v30;
        _2012_v30 = _module.D17.create_DC44(_2011_v29);
        let _source28 = _2012_v30;
        if (_source28.is_DC45) {
          let _2013___mcc_h6 = (_source28).cf88;
          let _2014_cf88 = _2013___mcc_h6;
          let _2015_v31;
          _2015_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2014_cf88,p3);
          let _2016_v32;
          _2016_v32 = _dafny.Seq.of(p3);
          let _2017_v33;
          _2017_v33 = _dafny.Set.fromElements(_2016_v32);
          let _2018_v34;
          _2018_v34 = _dafny.Seq.of(_2017_v33, _2017_v33, _2017_v33);
          let _rhs326 = (_this).fm6(_2001_v19, globalState);
          let _rhs327 = new BigNumber(((((_module.__default.fm0(p1, (_2009_v27).f13, p0, (((_2003_v21).contains(!((_2009_v27).f13))) ? ((_2003_v21).get(!((_2009_v27).f13))) : (p2)), globalState)).isEqualTo(new BigNumber((_2015_v31).length))) ? (_2001_v19) : (_2001_v19))).length);
          let _rhs328 = ((_2017_v33).Intersect(_2017_v33)).IsSubsetOf(((_2014_cf88) ? ((_2018_v34)[_module.__default.safeIndex(p2, new BigNumber((_2018_v34).length))]) : (_2017_v33)));
          let _lhs256 = globalState;
          let _lhs257 = globalState;
          _lhs256.f7 = _rhs326;
          _lhs257.f7 = _rhs327;
          r2 = _rhs328;
          r2 = p0;
          let _index358 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_2009_v27).f14).length));
          ((_2009_v27).f14)[_index358] = (_2009_v27).f13;
          let _index359 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_2009_v27).f14).length));
          ((_2009_v27).f14)[_index359] = _2014_cf88;
          let _2019_v35;
          _2019_v35 = _dafny.Seq.of(_2001_v19);
          let _2020_v36;
          _2020_v36 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC16(p3),_2019_v35);
          _2020_v36 = (_2020_v36).update(_module.D7.create_DC16(p3), _dafny.Seq.Concat(_2019_v35, _dafny.Seq.of(_2001_v19)));
        } else {
          let _2021___mcc_h7 = (_source28).cf87;
          let _2022_cf87 = _2021___mcc_h7;
          let _2023_v37;
          _2023_v37 = _dafny.MultiSet.fromElements(_2000_cf5);
          let _2024_v38;
          _2024_v38 = _2023_v37;
          let _2025_v39;
          _2025_v39 = _dafny.Seq.of(_2024_v38, _2024_v38);
          let _2026_v40;
          _2026_v40 = _dafny.Set.fromElements(_2001_v19);
          let _2027_v41;
          _2027_v41 = _dafny.Set.fromElements(((_dafny.ZERO).minus(new BigNumber((_2025_v39).length))).isLessThan(new BigNumber(488)), (p2).isLessThan(new BigNumber((_2026_v40).length)));
          _2027_v41 = _2027_v41;
          let _2028_v42;
          _2028_v42 = _dafny.MultiSet.fromElements(p2);
          let _2029_v43;
          let _nw295 = new _module.C0();
          _nw295.__ctor(true, _2005_v23, _2009_v27.f12, (_2028_v42).IsProperSubsetOf(_2028_v42));
          _2029_v43 = _nw295;
          let _2030_v44;
          _2030_v44 = _dafny.Seq.of(_1965_v0);
          _2030_v44 = _dafny.Seq.Concat(_2030_v44, _2030_v44);
          let _2031_v45;
          _2031_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2029_v43.f19,(_2009_v27).f13);
          _2031_v45 = (_2031_v45).update(p0, (_2009_v27).f13);
        }
        r2 = (_2009_v27).f13;
        let _2032_v46;
        let _nw296 = Array((new BigNumber(15)).toNumber()).fill([]);
        _2032_v46 = _nw296;
        let _index360 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2032_v46).length));
        (_2032_v46)[_index360] = _1976_v8;
        let _index361 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2032_v46).length));
        let _rhs329 = _2005_v23;
        let _rhs330 = _1976_v8;
        let _lhs258 = _2032_v46;
        let _lhs259 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2032_v46).length));
        _2005_v23 = _rhs329;
        _lhs258[_lhs259] = _rhs330;
      }
      let _2033_v47;
      let _nw297 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
      _2033_v47 = _nw297;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2033_v47).length))) {
        let _2034_i2 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2034_i2)) && ((_2034_i2).isLessThan(new BigNumber((_2033_v47).length))))) {
          (_2033_v47)[(_2034_i2)] = ((_dafny.Map.Empty.slice().updateUnsafe(p0,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p0))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p0)));
        }
      }
      let _2035_i3;
      _2035_i3 = _dafny.ZERO;
      L14: {
        while ((p3).isLessThanOrEqualTo(p3)) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2035_i3)) {
              break L14;
            }
            _2035_i3 = (_2035_i3).plus(_dafny.ONE);
            let _2036_v48;
            _2036_v48 = _dafny.MultiSet.fromElements(p3);
            _2036_v48 = _2036_v48;
            let _2037_v49;
            _2037_v49 = _dafny.Seq.UnicodeFromString("lvqcjon");
            let _2038_v50;
            _2038_v50 = new _dafny.CodePoint('w'.codePointAt(0));
            r1 = (new BigNumber(521)).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2037_v49, _module.__default.safeIndex(p2, new BigNumber((_2037_v49).length)), _2038_v50), _2037_v49)).length));
            (globalState).f7 = (_dafny.ZERO).minus(p2);
            let _2039_v51;
            _2039_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p3, p3),!(true));
            r2 = (((_2039_v51).contains(p2)) ? ((_2039_v51).get(p2)) : (p0));
          }
        }
      }
      let _2040_v52;
      _2040_v52 = _dafny.Seq.UnicodeFromString("rxfde");
      let _hi13 = new BigNumber((_2040_v52).length);
      for (let _2041_i4 = (p3).multipliedBy(new BigNumber(325)); _2041_i4.isLessThan(_hi13); _2041_i4 = _2041_i4.plus(_dafny.ONE)) {
        r0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(2)), function (_2042_i5) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })).length);
        let _2043_v53;
        _2043_v53 = _module.D15.create_DC39(p3, new _dafny.CodePoint('h'.codePointAt(0)), _dafny.Seq.UnicodeFromString("uny"), p3, new BigNumber(80));
        let _2044_v54;
        let _out15;
        _out15 = _module.__default.m0((_2043_v53).dtor_cf79, !(p0), globalState);
        _2044_v54 = _out15;
        (globalState).f2 = p1;
        r0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p2, _2041_i4));
      }
      let _2045_v55;
      _2045_v55 = _dafny.Set.fromElements(p0);
      r0 = ((p3).minus(new BigNumber((_2045_v55).length))).multipliedBy(p3);
      r1 = p2;
      r2 = p0;
      return [r0, r1, r2];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this.f11 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11) {
      let _this = this;
      (_this).f11 = f11;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f11;
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _2046_v0;
      _2046_v0 = _dafny.Seq.of(p2, _this.f11);
      let _hi14 = p1;
      for (let _2047_i0 = (_module.__default.fm0(_2046_v0, _this.f11, _this.f11, (_dafny.ZERO).minus(p0), globalState)).multipliedBy(new BigNumber(-964)); _2047_i0.isLessThan(_hi14); _2047_i0 = _2047_i0.plus(_dafny.ONE)) {
        (_this).f11 = _this.f11;
        let _2048_v1;
        _2048_v1 = _dafny.MultiSet.fromElements(p1);
        let _2049_v2;
        _2049_v2 = _dafny.Seq.UnicodeFromString("nmecnbdo");
        let _2050_v3;
        _2050_v3 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).fm4(_2046_v0, new BigNumber((_2048_v1).cardinality()), _2049_v2, _this.f11, globalState), p2),p1);
        _2050_v3 = _2050_v3;
        (globalState).f7 = p0;
        if (_this.f11) {
          (globalState).f7 = p1;
          let _2051_v4;
          let _nw298 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _2051_v4 = _nw298;
          _2051_v4 = _2051_v4;
          let _2052_v5;
          _2052_v5 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_2049_v2, _2049_v2),_module.__default.safeModuloInt(new BigNumber((_2046_v0).length), (_dafny.ZERO).minus(_2047_i0)));
          _2052_v5 = (_2052_v5).update(p2, (p1).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("npixpjow"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("npixpjow")).length)), _module.__default.fm5(globalState))).length)));
          r0 = p2;
          let _2053_v6;
          let _nw299 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _2053_v6 = _nw299;
          let _2054_v7;
          _2054_v7 = _dafny.Seq.of(_2046_v0);
          let _index362 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_2053_v6).length));
          (_2053_v6)[_index362] = _2054_v7;
          let _2055_v8;
          _2055_v8 = _dafny.Seq.of(p0);
          let _index363 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_2053_v6).length));
          (_2053_v6)[_index363] = _dafny.Seq.Concat(_2054_v7, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_2046_v0), _2054_v7), _module.__default.safeIndex((_2055_v8)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2055_v8).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_2046_v0), _2054_v7)).length)), _2046_v0));
        } else {
          let _2056_v9;
          _2056_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2047_i0);
          let _2057_v10;
          _2057_v10 = _module.D6.create_DC12(_this.f11, _2049_v2, _module.__default.fm0(_2046_v0, _this.f11, _this.f11, p0, globalState), _this.f11);
          let _2058_v11;
          _2058_v11 = _module.D1.create_DC1(_2049_v2);
          let _2059_v12;
          let _nw300 = Array((new BigNumber(17)).toNumber());
          _nw300[(_dafny.ZERO).toNumber()] = p2;
          _nw300[(_dafny.ONE).toNumber()] = p2;
          _nw300[(new BigNumber(2)).toNumber()] = true;
          _nw300[(new BigNumber(3)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(4)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(5)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(6)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(7)).toNumber()] = !(_this.f11);
          _nw300[(new BigNumber(8)).toNumber()] = p2;
          _nw300[(new BigNumber(9)).toNumber()] = p2;
          _nw300[(new BigNumber(10)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(11)).toNumber()] = p2;
          _nw300[(new BigNumber(12)).toNumber()] = p2;
          _nw300[(new BigNumber(13)).toNumber()] = _this.f11;
          _nw300[(new BigNumber(14)).toNumber()] = p2;
          _nw300[(new BigNumber(15)).toNumber()] = !(true);
          _nw300[(new BigNumber(16)).toNumber()] = p2;
          _2059_v12 = _nw300;
          let _2060_v13;
          _2060_v13 = _dafny.Set.fromElements(p0);
          let _pat_let_tv60 = p2;
          let _2061_v14;
          let _nw301 = new _module.C2();
          _nw301.__ctor((_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(_2049_v2)).length)).multipliedBy(new BigNumber((_2056_v9).length))), (((_2057_v10).dtor_cf21) ? (_this.f11) : (_this.f11)), _2058_v11, _module.__default.fm30(p1, globalState), _2059_v12, (_2060_v13).Difference(_dafny.Set.fromElements(p1)), function (_pat_let65_0) {
            return function (_2062_dt__update__tmp_h0) {
              return function (_pat_let66_0) {
                return function (_2063_dt__update_hcf4_h0) {
                  return function (_pat_let67_0) {
                    return function (_2064_dt__update_hcf3_h0) {
                      return _module.D1.create_DC2((_2062_dt__update__tmp_h0).dtor_cf2, _2064_dt__update_hcf3_h0, _2063_dt__update_hcf4_h0);
                    }(_pat_let67_0);
                  }(_2047_i0);
                }(_pat_let66_0);
              }(_pat_let_tv60);
            }(_pat_let65_0);
          }(_module.__default.fm37(p1, p0, new BigNumber((_2049_v2).length), new BigNumber((_dafny.Seq.UnicodeFromString("knl")).length), globalState)), p2);
          _2061_v14 = _nw301;
          (_this).f11 = !(_this.f11);
          _2059_v12 = _2059_v12;
          (globalState).f7 = _module.__default.safeDivisionInt(new BigNumber(650), (p1).plus((_2061_v14).f21));
          let _2065_v15;
          _2065_v15 = _dafny.Set.fromElements(p2);
          let _2066_v16;
          let _nw302 = new _module.C2();
          _nw302.__ctor((_dafny.ZERO).minus(new BigNumber((((p2) ? (_2049_v2) : (_dafny.Seq.UnicodeFromString("le")))).length)), (_this).fm4(_2046_v0, new BigNumber((_dafny.Seq.UnicodeFromString("ajkrilfy")).length), _dafny.Seq.UnicodeFromString("dli"), !((_2061_v14).f22), globalState), _module.D1.create_DC1(_2049_v2), !(_module.__default.fm17(globalState)), _2059_v12, (_2060_v13).Intersect(_2060_v13), _module.D1.create_DC2(_module.__default.fm0(_dafny.Seq.of(_this.f11, (_2061_v14).f22), p2, _this.f11, new BigNumber((_2065_v15).length), globalState), (_2061_v14).f21, (_2061_v14).f22), (_2061_v14).f22);
          _2066_v16 = _nw302;
        }
      }
      let _2067_v17;
      _2067_v17 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_this.f11, p2, _this.f11), _dafny.MultiSet.fromElements(_this.f11, _this.f11));
      let _2068_v18;
      _2068_v18 = _dafny.MultiSet.fromElements(p2);
      let _2069_v19;
      _2069_v19 = _dafny.Set.fromElements(p1, p0);
      let _2070_v20;
      _2070_v20 = _dafny.Seq.of(_dafny.Seq.of((((_2067_v17).contains(_2068_v18)) ? ((_2067_v17).get(_2068_v18)) : (p0)), _module.__default.fm0(_2046_v0, _this.f11, p2, new BigNumber((_2069_v19).length), globalState), p0, p1, p0));
      let _2071_v21;
      _2071_v21 = _dafny.Seq.of(p0, new BigNumber((_2046_v0).length));
      let _2072_v22;
      _2072_v22 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2073_v23;
      _2073_v23 = _dafny.Seq.UnicodeFromString("vrdiples");
      let _2074_v24;
      _2074_v24 = _module.D1.create_DC2(new BigNumber((_2073_v23).length), p1, _this.f11);
      let _2075_v25;
      let _nw303 = new _module.C1();
      _nw303.__ctor(_2072_v22, _2074_v24, _this.f11);
      _2075_v25 = _nw303;
      let _2076_v26;
      _2076_v26 = _dafny.MultiSet.fromElements(_2075_v25, _2075_v25);
      let _2077_v27;
      let _nw304 = Array((new BigNumber(25)).toNumber());
      _nw304[(_dafny.ZERO).toNumber()] = new BigNumber(791);
      _nw304[(_dafny.ONE).toNumber()] = p1;
      _nw304[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("w")).length);
      _nw304[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.of(p0)).length);
      _nw304[(new BigNumber(4)).toNumber()] = new BigNumber((_2070_v20).length);
      _nw304[(new BigNumber(5)).toNumber()] = p1;
      _nw304[(new BigNumber(6)).toNumber()] = p0;
      _nw304[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw304[(new BigNumber(8)).toNumber()] = p1;
      _nw304[(new BigNumber(9)).toNumber()] = new BigNumber(341);
      _nw304[(new BigNumber(10)).toNumber()] = p0;
      _nw304[(new BigNumber(11)).toNumber()] = p1;
      _nw304[(new BigNumber(12)).toNumber()] = p1;
      _nw304[(new BigNumber(13)).toNumber()] = p1;
      _nw304[(new BigNumber(14)).toNumber()] = p0;
      _nw304[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p1), p1)).length);
      _nw304[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_2046_v0, p2, _this.f11, p0, globalState);
      _nw304[(new BigNumber(17)).toNumber()] = p1;
      _nw304[(new BigNumber(18)).toNumber()] = p0;
      _nw304[(new BigNumber(19)).toNumber()] = p1;
      _nw304[(new BigNumber(20)).toNumber()] = p1;
      _nw304[(new BigNumber(21)).toNumber()] = p1;
      _nw304[(new BigNumber(22)).toNumber()] = (_2071_v21)[_module.__default.safeIndex(new BigNumber((_2076_v26).cardinality()), new BigNumber((_2071_v21).length))];
      _nw304[(new BigNumber(23)).toNumber()] = p1;
      _nw304[(new BigNumber(24)).toNumber()] = p0;
      _2077_v27 = _nw304;
      let _2078_v28;
      _2078_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2077_v27);
      let _hi15 = new BigNumber((_dafny.Seq.Concat(_2071_v21, _2071_v21)).length);
      for (let _2079_i1 = new BigNumber((((_this.f11) ? (_2078_v28) : (_dafny.Map.Empty.slice().updateUnsafe((((_2068_v18).contains(p2)) ? ((_2068_v18).get(p2)) : (p1)),_2077_v27)))).length); _2079_i1.isLessThan(_hi15); _2079_i1 = _2079_i1.plus(_dafny.ONE)) {
        (globalState).f7 = _2079_i1;
        (globalState).f7 = p1;
        let _2080_v29;
        _2080_v29 = _dafny.MultiSet.fromElements(_2072_v22, _2072_v22, ((_this.f11) ? (_2075_v25.f23) : (_2075_v25.f23)), new _dafny.CodePoint('f'.codePointAt(0)));
        let _2081_v30;
        _2081_v30 = _dafny.Seq.of(_2080_v29);
        _2080_v29 = (_2081_v30)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_this.f11)).length))), new BigNumber(980))), new BigNumber((_2081_v30).length))];
        let _2082_v31;
        _2082_v31 = _dafny.MultiSet.fromElements(new BigNumber(-531));
        (_this).f11 = (_this.f11) === ((_2082_v31).equals(_2082_v31));
      }
      let _2083_v32;
      _2083_v32 = _dafny.Set.fromElements(_this.f11, _this.f11);
      let _2084_v33;
      _2084_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2083_v32,_dafny.Seq.Concat(_2071_v21, _2071_v21));
      _2084_v33 = (_2084_v33).update(_2083_v32, _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(p0)).length), p0, new BigNumber(128), (_2071_v21)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_2071_v21).length))], p0));
      (_this).f11 = false;
      let _2085_v34;
      _2085_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),p2);
      let _2086_v35;
      _2086_v35 = _dafny.Seq.of(_2085_v34, _2085_v34, _2085_v34);
      let _2087_v36;
      _2087_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2086_v35);
      _2086_v35 = (((_2087_v36).contains(p1)) ? ((_2087_v36).get(p1)) : (_dafny.Seq.update(_dafny.Seq.of(_2085_v34, _2085_v34), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(_2085_v34, _2085_v34)).length)), _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f11))));
      let _2088_v37;
      _2088_v37 = _module.D1.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), ((_2089_v22) => function (_2090_i2) {
  return _2089_v22;
})(_2072_v22)));
      let _2091_v38;
      let _init49 = ((_2092_p2) => function (_2093_i3) {
        return _2092_p2;
      })(p2);
      let _nw305 = Array((new BigNumber(29)).toNumber());
      for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw305.length); _i0_49++) {
        _nw305[_i0_49] = _init49(new BigNumber(_i0_49));
      }
      _2091_v38 = _nw305;
      let _2094_v39;
      let _nw306 = new _module.C3();
      _nw306.__ctor(true, _module.__default.fm62(!(false), _2073_v23, globalState), p2, _2091_v38, _2069_v19, _2074_v24, _this.f11);
      _2094_v39 = _nw306;
      let _2095_v40;
      _2095_v40 = _dafny.MultiSet.fromElements(_2094_v39);
      let _2096_v41;
      _2096_v41 = _dafny.Seq.of(_2094_v39);
      let _2097_v42;
      _2097_v42 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_2096_v41), _dafny.MultiSet.FromArray(_dafny.Seq.of(_2094_v39)));
      let _2098_v43;
      let _nw307 = new _module.C2();
      _nw307.__ctor(p1, p2, _2088_v37, (_2097_v42).contains(_2095_v40), _2091_v38, _dafny.Set.fromElements(p0), _2074_v24, (_2094_v39).f18);
      _2098_v43 = _nw307;
      r0 = (_2098_v43).f22;
      return r0;
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _module.D2.Default();
      if (true) {
        (globalState).f7 = new BigNumber(56);
        let _2099_v0;
        _2099_v0 = _dafny.Seq.UnicodeFromString("yhesaykny");
        let _2100_v1;
        let _init50 = function (_2101_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_this.f11, _this.f11, _this.f11, _this.f11),_this.f11);
        };
        let _nw308 = Array((new BigNumber(23)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw308.length); _i0_50++) {
          _nw308[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _2100_v1 = _nw308;
        let _2102_v2;
        _2102_v2 = _dafny.Set.fromElements(_this.f11, !(_this.f11), !(_this.f11));
        let _2103_v3;
        _2103_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2102_v2,_this.f11);
        let _index364 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_2100_v1).length));
        (_2100_v1)[_index364] = _2103_v3;
        let _2104_v5;
        _2104_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2102_v2,new BigNumber(153));
        let _index365 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_2100_v1).length));
        let _rhs331 = _this.f11;
        let _rhs332 = _2099_v0;
        let _rhs333 = ((_dafny.Map.Empty.slice().updateUnsafe(_2102_v2,_this.f11)).Merge(function () {
          let _coll54 = new _dafny.Map();
          for (const _compr_54 of (_2104_v5).Keys.Elements) {
            let _2105_v4 = _compr_54;
            if ((_2104_v5).contains(_2105_v4)) {
              _coll54.push([_2105_v4,!(_this.f11)]);
            }
          }
          return _coll54;
        }())).Merge((_2103_v3).update(_2102_v2, _this.f11));
        let _rhs334 = true;
        let _rhs335 = false;
        let _lhs260 = _this;
        let _lhs261 = _2100_v1;
        let _lhs262 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_2100_v1).length));
        let _lhs263 = _this;
        let _lhs264 = _this;
        _lhs260.f11 = _rhs331;
        _2099_v0 = _rhs332;
        _lhs261[_lhs262] = _rhs333;
        _lhs263.f11 = _rhs334;
        _lhs264.f11 = _rhs335;
        let _2106_v6;
        _2106_v6 = _dafny.Seq.of(_this.f11);
        let _2107_v7;
        _2107_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2106_v6)[_module.__default.safeIndex(p0, new BigNumber((_2106_v6).length))],_this.f11);
        let _2108_v8;
        _2108_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,p0);
        let _2109_v9;
        _2109_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(new BigNumber((_2107_v7).length), _dafny.Seq.UnicodeFromString("narktrxe"), _this.f11, _2108_v8, globalState),_this.f11);
        let _2110_v10;
        _2110_v10 = new _dafny.CodePoint('k'.codePointAt(0));
        _2109_v9 = (_dafny.Map.Empty.slice().updateUnsafe(_2110_v10,_this.f11)).Merge(_2109_v9);
        let _2111_v11;
        _2111_v11 = _module.D6.create_DC12(_this.f11, _dafny.Seq.update(_dafny.Seq.Concat(_2099_v0, _2099_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_2099_v0, _2099_v0)).length)), _2110_v10), p0, _this.f11);
        let _source29 = _2111_v11;
        if (_source29.is_DC12) {
          let _2112___mcc_h0 = (_source29).cf18;
          let _2113___mcc_h1 = (_source29).cf19;
          let _2114___mcc_h2 = (_source29).cf20;
          let _2115___mcc_h3 = (_source29).cf21;
          let _2116_cf21 = _2115___mcc_h3;
          let _2117_cf20 = _2114___mcc_h2;
          let _2118_cf19 = _2113___mcc_h1;
          let _2119_cf18 = _2112___mcc_h0;
          let _2120_v12;
          let _init51 = ((_2121_p0) => function (_2122_i1) {
            return _module.__default.safeModuloInt(_2122_i1, (_dafny.ZERO).minus(_2121_p0));
          })(p0);
          let _nw309 = Array((new BigNumber(29)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw309.length); _i0_51++) {
            _nw309[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _2120_v12 = _nw309;
          _2120_v12 = _2120_v12;
          let _index366 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2120_v12).length));
          (_2120_v12)[_index366] = _2117_cf20;
          let _index367 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2120_v12).length));
          (_2120_v12)[_index367] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2099_v0).length), _2117_cf20));
          let _2123_v13;
          let _nw310 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
          _2123_v13 = _nw310;
          let _nw311 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
          _2123_v13 = _nw311;
          let _2124_v14;
          let _nw312 = Array((new BigNumber(24)).toNumber()).fill([]);
          _2124_v14 = _nw312;
          _2124_v14 = _2124_v14;
        } else if (_source29.is_DC13) {
          let _2125___mcc_h4 = (_source29).cf22;
          let _2126___mcc_h5 = (_source29).cf23;
          let _2127___mcc_h6 = (_source29).cf24;
          let _2128___mcc_h7 = (_source29).cf25;
          let _2129___mcc_h8 = (_source29).cf26;
          let _2130_cf26 = _2129___mcc_h8;
          let _2131_cf25 = _2128___mcc_h7;
          let _2132_cf24 = _2127___mcc_h6;
          let _2133_cf23 = _2126___mcc_h5;
          let _2134_cf22 = _2125___mcc_h4;
          let _2135_v15;
          _2135_v15 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_2131_cf25, p0), ((_dafny.ZERO).minus(_2131_cf25)).plus(_2131_cf25), _2132_cf24, (p0).minus(p0));
          _2135_v15 = _2135_v15;
          let _2136_v16;
          _2136_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2106_v6,_2134_cf22);
          let _2137_v17;
          _2137_v17 = _module.D12.create_DC29(_2136_v16);
          let _2138_v18;
          _2138_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2137_v17,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("swnt"), _module.__default.safeIndex(_2131_cf25, new BigNumber((_dafny.Seq.UnicodeFromString("swnt")).length)), _2110_v10));
          _2138_v18 = (_2138_v18).update(_2137_v17, ((_2130_cf26) ? (_dafny.Seq.UnicodeFromString("dii")) : (_2099_v0)));
          let _2139_v19;
          let _nw313 = new _module.C7();
          _nw313.__ctor();
          _2139_v19 = _nw313;
          let _2140_v20;
          _2140_v20 = _dafny.Map.Empty.slice().updateUnsafe((_2132_cf24).plus(_2131_cf25),_module.__default.safeModuloInt(p0, p0));
          _2140_v20 = (_2140_v20).update(p0, p0);
        } else if (_source29.is_DC11) {
          let _2141___mcc_h9 = (_source29).cf17;
          let _2142_cf17 = _2141___mcc_h9;
          let _2143_v22;
          _2143_v22 = _dafny.Seq.of(p0);
          let _2144_v23;
          _2144_v23 = _dafny.MultiSet.fromElements(function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of (_2143_v22).Elements) {
              let _2145_v21 = _compr_55;
              if (_dafny.Seq.contains(_2143_v22, _2145_v21)) {
                _coll55.push([(_2145_v21).plus(p0),_this.f11]);
              }
            }
            return _coll55;
          }());
          let _2146_v24;
          _2146_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f11);
          let _2147_v25;
          let _out16;
          _out16 = (_this).m2(new BigNumber(((_2144_v23).Difference(_dafny.MultiSet.fromElements(_2146_v24))).cardinality()), new BigNumber((_2099_v0).length), !(_this.f11), globalState);
          _2147_v25 = _out16;
          let _2148_v26;
          let _nw314 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2148_v26 = _nw314;
          let _2149_v27;
          let _nw315 = Array((new BigNumber(3)).toNumber());
          _nw315[(_dafny.ZERO).toNumber()] = p0;
          _nw315[(_dafny.ONE).toNumber()] = p0;
          _nw315[(new BigNumber(2)).toNumber()] = p0;
          _2149_v27 = _nw315;
          let _index368 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2148_v26).length));
          (_2148_v26)[_index368] = _2149_v27;
          let _index369 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2148_v26).length));
          let _rhs336 = !(_2147_v25);
          let _rhs337 = _2149_v27;
          let _lhs265 = _this;
          let _lhs266 = _2148_v26;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2148_v26).length));
          _lhs265.f11 = _rhs336;
          _lhs266[_lhs267] = _rhs337;
          _2143_v22 = _dafny.Seq.of(new BigNumber(533), new BigNumber(549));
          let _2150_v28;
          let _init52 = ((_2151_v0) => function (_2152_i2) {
            return _dafny.Set.fromElements(new BigNumber((_2151_v0).length));
          })(_2099_v0);
          let _nw316 = Array((new BigNumber(5)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw316.length); _i0_52++) {
            _nw316[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _2150_v28 = _nw316;
          _2150_v28 = _2150_v28;
        } else {
          let _2153___mcc_h10 = (_source29).cf27;
          let _2154_cf27 = _2153___mcc_h10;
          let _2155_v29;
          _2155_v29 = _module.D1.create_DC1(_dafny.Seq.update(_2099_v0, _module.__default.safeIndex(new BigNumber(-526), new BigNumber((_2099_v0).length)), new _dafny.CodePoint('s'.codePointAt(0))));
          let _2156_v30;
          let _init53 = function (_2157_i3) {
            return !(_this.f11);
          };
          let _nw317 = Array((new BigNumber(9)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw317.length); _i0_53++) {
            _nw317[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _2156_v30 = _nw317;
          let _2158_v31;
          _2158_v31 = _dafny.Set.fromElements(p0);
          let _2159_v32;
          let _nw318 = new _module.C6();
          _nw318.__ctor(_2155_v29, _this.f11, _2156_v30, _2158_v31, _module.D1.create_DC2(p0, p0, _this.f11), _this.f11);
          _2159_v32 = _nw318;
          let _2160_v33;
          _2160_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2159_v32,_2107_v7);
          let _2161_v34;
          _2161_v34 = _dafny.MultiSet.fromElements(_this.f11, (_2159_v32).fm11(globalState));
          let _2162_v35;
          _2162_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2161_v34).cardinality()),p0);
          let _2163_v36;
          _2163_v36 = _dafny.Seq.of(new BigNumber((_2162_v35).length), p0, new BigNumber(352), p0);
          let _2164_v37;
          let _nw319 = Array((new BigNumber(13)).toNumber());
          _nw319[(_dafny.ZERO).toNumber()] = new BigNumber((_2099_v0).length);
          _nw319[(_dafny.ONE).toNumber()] = new BigNumber((_2160_v33).length);
          _nw319[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p0)).length));
          _nw319[(new BigNumber(3)).toNumber()] = p0;
          _nw319[(new BigNumber(4)).toNumber()] = new BigNumber((_2163_v36).length);
          _nw319[(new BigNumber(5)).toNumber()] = p0;
          _nw319[(new BigNumber(6)).toNumber()] = (new BigNumber((_2099_v0).length)).minus(new BigNumber(983));
          _nw319[(new BigNumber(7)).toNumber()] = new BigNumber((_2099_v0).length);
          _nw319[(new BigNumber(8)).toNumber()] = p0;
          _nw319[(new BigNumber(9)).toNumber()] = p0;
          _nw319[(new BigNumber(10)).toNumber()] = p0;
          _nw319[(new BigNumber(11)).toNumber()] = new BigNumber(((_2107_v7).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f11,true))).length);
          _nw319[(new BigNumber(12)).toNumber()] = p0;
          _2164_v37 = _nw319;
          let _2165_v38;
          _2165_v38 = _module.D2.create_DC4(_2164_v37, _this.f11, p0, _this.f11, _2156_v30);
          let _index370 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_2164_v37).length));
          (_2164_v37)[_index370] = ((_this.f11) ? ((((_2162_v35).contains(p0)) ? ((_2162_v35).get(p0)) : (p0))) : ((_2165_v38).dtor_cf8));
          let _index371 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_2164_v37).length));
          (_2164_v37)[_index371] = ((p0).multipliedBy(p0)).minus(p0);
          (globalState).f7 = (_2164_v37)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_2164_v37).length))];
          (globalState).f7 = _module.__default.safeDivisionInt(new BigNumber((_2099_v0).length), (_2164_v37)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_2164_v37).length))]);
          let _index372 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_2164_v37).length));
          (_2164_v37)[_index372] = p0;
        }
        let _2166_v39;
        let _init54 = ((_2167_v0) => function (_2168_i4) {
          return _2167_v0;
        })(_2099_v0);
        let _nw320 = Array((_dafny.ONE).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw320.length); _i0_54++) {
          _nw320[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2166_v39 = _nw320;
        _2166_v39 = _2166_v39;
      } else {
        let _2169_v40;
        _2169_v40 = _dafny.Seq.UnicodeFromString("dpqp");
        let _2170_v41;
        _2170_v41 = new _dafny.CodePoint('y'.codePointAt(0));
        let _2171_v42;
        let _nw321 = Array((new BigNumber(11)).toNumber()).fill(false);
        _2171_v42 = _nw321;
        let _2172_v43;
        _2172_v43 = _module.D12.create_DC30(_2171_v42, _this.f11);
        let _2173_v44;
        _2173_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _2174_v45;
        _2174_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
        let _2175_v46;
        let _nw322 = Array((new BigNumber(27)).toNumber());
        _nw322[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("sq");
        _nw322[(_dafny.ONE).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(2)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(3)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_2169_v40, _dafny.Seq.UnicodeFromString("ugei"));
        _nw322[(new BigNumber(5)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(6)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2169_v40, _dafny.Seq.UnicodeFromString("vjrdhd"));
        _nw322[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hd"), _2169_v40);
        _nw322[(new BigNumber(9)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(10)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(11)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_2169_v40, _module.__default.safeIndex(p0, new BigNumber((_2169_v40).length)), _2170_v41);
        _nw322[(new BigNumber(13)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(14)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(15)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("scbj"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("scbj")).length)), _2170_v41);
        _nw322[(new BigNumber(17)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(18)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(19)).toNumber()] = ((_this.f11) ? (_dafny.Seq.UnicodeFromString("bbwe")) : (_2169_v40));
        _nw322[(new BigNumber(20)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(21)).toNumber()] = ((_this.f11) ? (_dafny.Seq.update(_2169_v40, _module.__default.safeIndex(p0, new BigNumber((_2169_v40).length)), _2170_v41)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), ((_2176_v41) => function (_2177_i5) {
          return _2176_v41;
        })(_2170_v41))));
        _nw322[(new BigNumber(22)).toNumber()] = _module.__default.fm60(_this.f11, !(!((_2172_v43).dtor_cf59)), _2173_v44, p0, globalState);
        _nw322[(new BigNumber(23)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ihv"), _module.__default.safeIndex(new BigNumber((_2174_v45).length), new BigNumber((_dafny.Seq.UnicodeFromString("ihv")).length)), new _dafny.CodePoint('e'.codePointAt(0)));
        _nw322[(new BigNumber(24)).toNumber()] = _2169_v40;
        _nw322[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_2169_v40, _2169_v40);
        _nw322[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_2169_v40, _dafny.Seq.UnicodeFromString("kf"));
        _2175_v46 = _nw322;
        let _index373 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2175_v46).length));
        (_2175_v46)[_index373] = _dafny.Seq.Concat(_2169_v40, _dafny.Seq.Create(_module.__default.abs(new BigNumber(927)), ((_2178_v41) => function (_2179_i6) {
          return _2178_v41;
        })(_2170_v41)));
        let _2180_v47;
        let _nw323 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2180_v47 = _nw323;
        let _2181_v48;
        _2181_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(828),_2180_v47);
        let _2182_v49;
        _2182_v49 = _dafny.MultiSet.fromElements(_2180_v47, _2180_v47, (((_2181_v48).contains(p0)) ? ((_2181_v48).get(p0)) : (_2180_v47)));
        let _2183_v50;
        _2183_v50 = _dafny.MultiSet.fromElements(new BigNumber(773));
        let _index374 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2175_v46).length));
        let _rhs338 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), ((_2184_v41) => function (_2185_i7) {
          return _2184_v41;
        })(_2170_v41));
        let _rhs339 = _2182_v49;
        let _rhs340 = (((_2173_v44).contains(p0)) ? ((_2173_v44).get(p0)) : ((_2183_v50).IsDisjointFrom(_2183_v50)));
        let _rhs341 = _this.f11;
        let _lhs268 = _2175_v46;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2175_v46).length));
        let _lhs270 = _this;
        let _lhs271 = _this;
        _lhs268[_lhs269] = _rhs338;
        _2182_v49 = _rhs339;
        _lhs270.f11 = _rhs340;
        _lhs271.f11 = _rhs341;
        let _index375 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length));
        (_2171_v42)[_index375] = ((_this.f11) ? (_this.f11) : (_this.f11));
        let _2186_v51;
        let _init55 = ((_2187_p0) => function (_2188_i8) {
          return _module.__default.safeDivisionInt(_2188_i8, _2187_p0);
        })(p0);
        let _nw324 = Array((new BigNumber(23)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw324.length); _i0_55++) {
          _nw324[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2186_v51 = _nw324;
        let _index376 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2186_v51).length));
        (_2186_v51)[_index376] = p0;
        let _2189_v52;
        _2189_v52 = _dafny.Seq.of(true, false, true);
        let _2190_v53;
        _2190_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(new BigNumber(596), globalState),_2171_v42);
        let _index377 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length));
        let _index378 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2186_v51).length));
        let _rhs342 = !(true) || ((p0).isLessThanOrEqualTo(new BigNumber((_2189_v52).length)));
        let _rhs343 = p0;
        let _rhs344 = (p0).multipliedBy(_module.__default.safeModuloInt(_module.__default.fm0(_2189_v52, _this.f11, _this.f11, p0, globalState), p0));
        let _rhs345 = (((_2190_v53).contains(_module.__default.fm30((_dafny.ZERO).minus(p0), globalState))) ? ((_2190_v53).get(_module.__default.fm30((_dafny.ZERO).minus(p0), globalState))) : (_2171_v42));
        let _rhs346 = _module.__default.fm0(_2189_v52, (_this.f11) && (_this.f11), false, new BigNumber(953), globalState);
        let _lhs272 = _2171_v42;
        let _lhs273 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length));
        let _lhs274 = globalState;
        let _lhs275 = globalState;
        let _lhs276 = _2186_v51;
        let _lhs277 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2186_v51).length));
        _lhs272[_lhs273] = _rhs342;
        _lhs274.f7 = _rhs343;
        _lhs275.f7 = _rhs344;
        _2171_v42 = _rhs345;
        _lhs276[_lhs277] = _rhs346;
        (globalState).f7 = (((new BigNumber((_2169_v40).length)).isLessThanOrEqualTo(p0)) ? ((_2186_v51)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_2186_v51).length))]) : (_module.__default.fm0(_2189_v52, _this.f11, (_2171_v42)[_module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length))], p0, globalState)));
        let _index379 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length));
        (_2171_v42)[_index379] = _this.f11;
        _2173_v44 = (_2173_v44).update(p0, (_2171_v42)[_module.__default.safeIndex(new BigNumber(228), new BigNumber((_2171_v42).length))]);
      }
      let _2191_v54;
      let _init56 = function (_2192_i10) {
        return (_dafny.MultiSet.fromElements(_this.f11)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_this.f11));
      };
      let _nw325 = Array((new BigNumber(26)).toNumber());
      for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw325.length); _i0_56++) {
        _nw325[_i0_56] = _init56(new BigNumber(_i0_56));
      }
      _2191_v54 = _nw325;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2191_v54).length))) {
        let _2193_i9 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2193_i9)) && ((_2193_i9).isLessThan(new BigNumber((_2191_v54).length))))) {
          (_2191_v54)[(_2193_i9)] = _this.f11;
        }
      }
      let _2194_v55;
      let _init57 = ((_2195_p0) => function (_2196_i12) {
        return _dafny.Set.fromElements(_2195_p0, _2195_p0);
      })(p0);
      let _nw326 = Array((new BigNumber(17)).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw326.length); _i0_57++) {
        _nw326[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _2194_v55 = _nw326;
      for (const _guard_loop_10 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2194_v55).length))) {
        let _2197_i11 = _guard_loop_10;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2197_i11)) && ((_2197_i11).isLessThan(new BigNumber((_2194_v55).length))))) {
          (_2194_v55)[(_2197_i11)] = _dafny.Set.fromElements((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(235))), p0);
        }
      }
      (_this).f11 = !(p0).isEqualTo((new BigNumber(-217)).multipliedBy(p0));
      if (true) {
        let _2198_v56;
        _2198_v56 = new _dafny.CodePoint('n'.codePointAt(0));
        let _2199_v57;
        _2199_v57 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        let _2200_v58;
        _2200_v58 = _module.D1.create_DC2((((_2199_v57).contains(_this.f11)) ? ((_2199_v57).get(_this.f11)) : (p0)), (_dafny.ZERO).minus(p0), _this.f11);
        let _2201_v59;
        let _nw327 = new _module.C0();
        _nw327.__ctor(_this.f11, _2198_v56, _2200_v58, !(_this.f11));
        _2201_v59 = _nw327;
        _2201_v59 = _2201_v59;
        (globalState).f7 = p0;
        if (_2201_v59.f19) {
          let _2202_v60;
          _2202_v60 = _dafny.Seq.of(_2201_v59.f19, _2201_v59.f19, _this.f11);
          (globalState).f7 = _module.__default.fm0(_2202_v60, _2201_v59.f19, _this.f11, (new BigNumber((_2202_v60).length)).minus(p0), globalState);
          let _2203_v61;
          _2203_v61 = _dafny.Seq.UnicodeFromString("ped");
          let _rhs347 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm30(new BigNumber((_dafny.Set.fromElements(_2203_v61, _2203_v61)).length), globalState))).length)).plus(_module.__default.fm0(_dafny.Seq.of(_2201_v59.f19, _2201_v59.f19), _this.f11, _this.f11, p0, globalState));
          let _rhs348 = p0;
          let _lhs278 = globalState;
          let _lhs279 = globalState;
          _lhs278.f7 = _rhs347;
          _lhs279.f7 = _rhs348;
          let _2204_v62;
          _2204_v62 = _dafny.Seq.of(new BigNumber(7));
          let _2205_v63;
          _2205_v63 = _module.D7.create_DC15(_2204_v62);
          let _2206_v64;
          _2206_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2201_v59.f19,(_2201_v59).f20);
          let _2207_v65;
          let _nw328 = new _module.C4();
          _nw328.__ctor();
          _2207_v65 = _nw328;
          let _2208_v66;
          _2208_v66 = _dafny.Seq.of(_2207_v65);
          let _2209_v67;
          let _nw329 = Array((new BigNumber(20)).toNumber());
          _nw329[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2204_v62, _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_this.f11, _2201_v59.f19)).cardinality()), _module.__default.fm0(_2202_v60, _this.f11, _2201_v59.f19, p0, globalState)));
          _nw329[(_dafny.ONE).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(2)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(3)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat((_2205_v63).dtor_cf28, _2204_v62);
          _nw329[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(828)), ((_2210_p0) => function (_2211_i13) {
            return _2210_p0;
          })(p0));
          _nw329[(new BigNumber(6)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(7)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_2204_v62, _2204_v62);
          _nw329[(new BigNumber(9)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(10)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(11)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), ((_2212_v59, _2213_p0) => function (_2214_i14) {
            return new BigNumber((_dafny.Set.fromElements(new BigNumber(432), new BigNumber((_dafny.Set.fromElements(_2212_v59.f19, _this.f11, _2212_v59.f19, _this.f11, _2212_v59.f19)).length), _2213_p0)).length);
          })(_2201_v59, p0)), _2204_v62);
          _nw329[(new BigNumber(13)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(p0);
          _nw329[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_2204_v62, _dafny.Seq.update(_2204_v62, _module.__default.safeIndex(new BigNumber((_2206_v64).length), new BigNumber((_2204_v62).length)), p0));
          _nw329[(new BigNumber(16)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(17)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(18)).toNumber()] = _2204_v62;
          _nw329[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(p0, new BigNumber((_2208_v66).length));
          _2209_v67 = _nw329;
          let _rhs349 = (_dafny.ZERO).minus(p0);
          let _rhs350 = _2209_v67;
          let _lhs280 = globalState;
          _lhs280.f7 = _rhs349;
          _2209_v67 = _rhs350;
          let _index380 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index380] = _2201_v59.f19;
          let _2215_v68;
          _2215_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2201_v59.f19);
          let _index381 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index381] = _dafny.areEqual(_2204_v62, _dafny.Seq.update(_dafny.Seq.of(p0, new BigNumber((_2215_v68).length), p0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(p0, new BigNumber((_2215_v68).length), p0)).length)), new BigNumber(725)));
          let _2216_v69;
          _2216_v69 = _dafny.Set.fromElements(_2201_v59.f19, true, _2201_v59.f19);
          let _2217_v70;
          _2217_v70 = _dafny.MultiSet.fromElements(_this.f11, _this.f11);
          let _index382 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          let _index383 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          let _rhs351 = !(_2216_v69).contains((_2191_v54)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length))]);
          let _rhs352 = ((_module.__default.fm40(globalState)).update(_this.f11, _module.__default.abs(new BigNumber((_2216_v69).length)))).IsSubsetOf(_2217_v70);
          let _lhs281 = _2191_v54;
          let _lhs282 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          let _lhs283 = _2191_v54;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2191_v54).length));
          _lhs281[_lhs282] = _rhs351;
          _lhs283[_lhs284] = _rhs352;
        } else {
          let _2218_v71;
          _2218_v71 = _dafny.Seq.UnicodeFromString("arsssdc");
          _2198_v56 = (_2218_v71)[_module.__default.safeIndex(p0, new BigNumber((_2218_v71).length))];
          let _2219_v72;
          _2219_v72 = _dafny.Set.fromElements(_2218_v71);
          let _2220_v73;
          _2220_v73 = _dafny.Map.Empty.slice().updateUnsafe((_2219_v72).Difference(_dafny.Set.fromElements(_2218_v71, _2218_v71)),_2201_v59.f19);
          _2220_v73 = (_2220_v73).update(_2219_v72, _2201_v59.f19);
          (_2201_v59).f19 = ((_2201_v59.f19) ? (false) : (false));
          let _index384 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index384] = false;
          let _index385 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index385] = _this.f11;
          let _2221_v74;
          _2221_v74 = _dafny.Set.fromElements(p0);
          let _2222_v75;
          _2222_v75 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_2221_v74).IsSubsetOf(_2221_v74));
          let _2223_v76;
          let _init58 = ((_2224_p0) => function (_2225_i15) {
            return (_2225_i15).minus(_2224_p0);
          })(p0);
          let _nw330 = Array((new BigNumber(8)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw330.length); _i0_58++) {
            _nw330[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2223_v76 = _nw330;
          let _2226_v77;
          _2226_v77 = _dafny.MultiSet.fromElements((_2201_v59).f20);
          let _2227_v78;
          _2227_v78 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2226_v77);
          let _2228_v79;
          _2228_v79 = _module.D12.create_DC31(_this.f11, _2223_v76, _2201_v59.f19, (_2201_v59).f20, (((_2227_v78).contains(_2201_v59.f19)) ? ((_2227_v78).get(_2201_v59.f19)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of((_2201_v59).f20)))));
          let _2229_v80;
          _2229_v80 = _dafny.Set.fromElements(true, (_2191_v54)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length))], (_2191_v54)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length))], _2201_v59.f19, true);
          let _index386 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index386] = (((_2222_v75).contains(!((_2191_v54)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length))]) || ((_2228_v79).dtor_cf62))) ? ((_2222_v75).get(!((_2191_v54)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_2191_v54).length))]) || ((_2228_v79).dtor_cf62))) : (!(_module.__default.fm30(new BigNumber((_2229_v80).length), globalState))));
        }
        let _2230_v81;
        _2230_v81 = _dafny.Seq.of(_2201_v59.f19, true);
        let _2231_v82;
        _2231_v82 = _dafny.MultiSet.fromElements(new BigNumber((_2230_v81).length));
        let _2232_v83;
        _2232_v83 = _dafny.Map.Empty.slice().updateUnsafe(_2231_v82,_2191_v54);
        let _2233_v84;
        _2233_v84 = _dafny.Set.fromElements(p0);
        let _2234_v85;
        let _nw331 = new _module.C5();
        _nw331.__ctor(_2201_v59.f19, (((_2232_v83).contains(_2231_v82)) ? ((_2232_v83).get(_2231_v82)) : (_2191_v54)), (_2233_v84).Difference(_2233_v84), _2200_v58, _2201_v59.f19);
        _2234_v85 = _nw331;
        if (_this.f11) {
          let _2235_v86;
          _2235_v86 = _dafny.Seq.UnicodeFromString("hhqingxd");
          let _index387 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index387] = (_this).fm4(_2230_v81, p0, _2235_v86, _2234_v85.f24, globalState);
          let _2236_v87;
          _2236_v87 = _dafny.Seq.of(_2235_v86, _2235_v86, _2235_v86, _2235_v86, _2235_v86);
          let _2237_v88;
          _2237_v88 = _dafny.Map.Empty.slice().updateUnsafe((_2236_v87)[_module.__default.safeIndex(p0, new BigNumber((_2236_v87).length))],_2199_v57);
          let _index388 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index388] = (_2237_v88).contains(_2235_v86);
          let _2238_v89;
          _2238_v89 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(p0));
          let _2239_v90;
          _2239_v90 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_2234_v85.f24));
          let _2240_v91;
          _2240_v91 = _dafny.Set.fromElements(false, _2234_v85.f24);
          let _2241_v92;
          _2241_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2240_v91).length),p0);
          let _2242_v93;
          _2242_v93 = _dafny.Seq.of(_2231_v82, _2231_v82);
          let _2243_v94;
          _2243_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2240_v91,new BigNumber((_2230_v81).length));
          let _2244_v95;
          let _nw332 = Array((new BigNumber(16)).toNumber());
          _nw332[(_dafny.ZERO).toNumber()] = (_2238_v89)[_module.__default.safeIndex((_2234_v85).fm9((_2201_v59).f20, globalState), new BigNumber((_2238_v89).length))];
          _nw332[(_dafny.ONE).toNumber()] = p0;
          _nw332[(new BigNumber(2)).toNumber()] = new BigNumber((_2239_v90).length);
          _nw332[(new BigNumber(3)).toNumber()] = p0;
          _nw332[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_2234_v85).fm9((_2201_v59).f20, globalState));
          _nw332[(new BigNumber(5)).toNumber()] = p0;
          _nw332[(new BigNumber(6)).toNumber()] = (p0).plus((_dafny.ZERO).minus((((_2241_v92).contains(p0)) ? ((_2241_v92).get(p0)) : (p0))));
          _nw332[(new BigNumber(7)).toNumber()] = p0;
          _nw332[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p0, p0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(p0, p0)).length)), (_dafny.ZERO).minus(p0))).length));
          _nw332[(new BigNumber(9)).toNumber()] = (p0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length));
          _nw332[(new BigNumber(10)).toNumber()] = p0;
          _nw332[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw332[(new BigNumber(12)).toNumber()] = new BigNumber((_2239_v90).length);
          _nw332[(new BigNumber(13)).toNumber()] = new BigNumber(366);
          _nw332[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((_2242_v93)[_module.__default.safeIndex(p0, new BigNumber((_2242_v93).length))]).cardinality()), new BigNumber((_module.__default.fm64(globalState)).length));
          _nw332[(new BigNumber(15)).toNumber()] = new BigNumber(((_2243_v94).Merge(_2243_v94)).length);
          _2244_v95 = _nw332;
          _2244_v95 = _2244_v95;
          (globalState).f7 = (p0).multipliedBy(p0);
          let _rhs353 = !(_this.f11);
          let _rhs354 = p0;
          let _lhs285 = _2234_v85;
          let _lhs286 = globalState;
          _lhs285.f24 = _rhs353;
          _lhs286.f7 = _rhs354;
          (_this).f11 = (_2201_v59.f19) === (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("sftaoljs"), (_2236_v87)[_module.__default.safeIndex(new BigNumber((_2240_v91).length), new BigNumber((_2236_v87).length))]));
        } else {
          let _2245_v96;
          _2245_v96 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f11);
          let _2246_v97;
          _2246_v97 = _dafny.MultiSet.fromElements((_2245_v96).Merge(_2245_v96));
          let _2247_v98;
          _2247_v98 = _dafny.Seq.of(_2245_v96, _2245_v96, _2245_v96, _dafny.Map.Empty.slice().updateUnsafe(_2201_v59.f19,_2201_v59.f19), (_2245_v96).update(_2234_v85.f24, _this.f11));
          _2246_v97 = (_dafny.MultiSet.FromArray(_2247_v98)).Difference(_2246_v97);
          let _2248_v99;
          _2248_v99 = _dafny.MultiSet.fromElements(_2191_v54, _2191_v54, _2191_v54, _2191_v54);
          _2248_v99 = _2248_v99;
          let _2249_v100;
          _2249_v100 = _dafny.Seq.UnicodeFromString("bvpukkl");
          let _2250_v102;
          _2250_v102 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ipcpel"), _2249_v100);
          let _2251_v103;
          _2251_v103 = _dafny.MultiSet.fromElements(_2249_v100);
          let _2252_v105;
          _2252_v105 = _dafny.Seq.of(_dafny.Set.fromElements(_2249_v100, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-402)), ((_2253_v59) => function (_2254_i16) {
            return (_2253_v59).f20;
          })(_2201_v59)), _2249_v100), function () {
            let _coll56 = new _dafny.Set();
            for (const _compr_56 of (_2251_v103).Elements) {
              let _2255_v104 = _compr_56;
              if ((_2251_v103).contains(_2255_v104)) {
                _coll56.add(_2255_v104);
              }
            }
            return _coll56;
          }());
          let _2256_v106;
          _2256_v106 = _dafny.Seq.of((_dafny.ZERO).minus(p0), p0);
          let _2257_v108;
          _2257_v108 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _2258_v109;
          let _nw333 = new _module.C7();
          _nw333.__ctor();
          _2258_v109 = _nw333;
          let _2259_v110;
          _2259_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v100,_2258_v109);
          let _2260_v111;
          _2260_v111 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_2257_v108).contains(p0)) ? ((_2257_v108).get(p0)) : (new BigNumber((_2259_v110).length))));
          let _2261_v112;
          let _nw334 = Array((new BigNumber(17)).toNumber());
          _nw334[(_dafny.ZERO).toNumber()] = function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_dafny.Seq.of(_2249_v100)).Elements) {
              let _2262_v101 = _compr_57;
              if (_dafny.Seq.contains(_dafny.Seq.of(_2249_v100), _2262_v101)) {
                _coll57.add(_2262_v101);
              }
            }
            return _coll57;
          }();
          _nw334[(_dafny.ONE).toNumber()] = _module.__default.fm48(globalState);
          _nw334[(new BigNumber(2)).toNumber()] = (_2250_v102).Union(_2250_v102);
          _nw334[(new BigNumber(3)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(4)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(5)).toNumber()] = ((_2252_v105)[_module.__default.safeIndex((_2256_v106)[_module.__default.safeIndex(p0, new BigNumber((_2256_v106).length))], new BigNumber((_2252_v105).length))]).Difference(_2250_v102);
          _nw334[(new BigNumber(6)).toNumber()] = function () {
            let _coll58 = new _dafny.Set();
            for (const _compr_58 of (_2251_v103).Elements) {
              let _2263_v107 = _compr_58;
              if ((_2251_v103).contains(_2263_v107)) {
                _coll58.add(_2263_v107);
              }
            }
            return _coll58;
          }();
          _nw334[(new BigNumber(7)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(8)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(9)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(10)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(11)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(12)).toNumber()] = (_2252_v105)[_module.__default.safeIndex(new BigNumber((_2257_v108).length), new BigNumber((_2252_v105).length))];
          _nw334[(new BigNumber(13)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(14)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(15)).toNumber()] = _2250_v102;
          _nw334[(new BigNumber(16)).toNumber()] = (_2250_v102).Intersect(_module.__default.fm48(globalState));
          _2261_v112 = _nw334;
          let _index389 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_2261_v112).length));
          (_2261_v112)[_index389] = _2250_v102;
          let _index390 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_2261_v112).length));
          (_2261_v112)[_index390] = (((_this).fm4(_2230_v81, p0, _2249_v100, _2201_v59.f19, globalState)) ? (_2250_v102) : ((_dafny.Set.fromElements(_2249_v100, _2249_v100)).Difference(_2250_v102)));
          let _2264_v113;
          let _nw335 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2264_v113 = _nw335;
          let _2265_v114;
          _2265_v114 = _dafny.MultiSet.fromElements(_2264_v113);
          _2265_v114 = (_2265_v114).Difference(_2265_v114);
          (_this).f11 = !(false);
        }
      } else {
        let _2266_v115;
        _2266_v115 = new _dafny.CodePoint('i'.codePointAt(0));
        let _2267_v116;
        _2267_v116 = _dafny.Map.Empty.slice().updateUnsafe((p0).isLessThanOrEqualTo(new BigNumber(680)),_2266_v115);
        _2267_v116 = _2267_v116;
        let _2268_v117;
        let _nw336 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _2268_v117 = _nw336;
        let _2269_v118;
        _2269_v118 = _dafny.Set.fromElements(_this.f11);
        let _index391 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length));
        (_2268_v117)[_index391] = (new BigNumber((_2269_v118).length)).minus(p0);
        let _2270_v119;
        _2270_v119 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2269_v118);
        let _2271_v120;
        _2271_v120 = _dafny.Seq.of(_this.f11, _this.f11);
        let _2272_v121;
        _2272_v121 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2269_v118);
        let _index392 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length));
        (_2268_v117)[_index392] = new BigNumber(((((_2270_v119).contains((_dafny.ZERO).minus(_module.__default.fm0(_2271_v120, _this.f11, _this.f11, p0, globalState)))) ? ((_2270_v119).get((_dafny.ZERO).minus(_module.__default.fm0(_2271_v120, _this.f11, _this.f11, p0, globalState)))) : (((_this.f11) ? ((((_2272_v121).contains(false)) ? ((_2272_v121).get(false)) : (_2269_v118))) : (_2269_v118))))).length);
        let _2273_v122;
        _2273_v122 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_2269_v118).length)), (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], p0);
        let _rhs355 = _2273_v122;
        let _rhs356 = ((p0).plus(p0)).minus(p0);
        let _rhs357 = _this.f11;
        let _lhs287 = globalState;
        let _lhs288 = _this;
        _2273_v122 = _rhs355;
        _lhs287.f7 = _rhs356;
        _lhs288.f11 = _rhs357;
        if (_this.f11) {
          (_this).f11 = ((_this.f11) ? (((_this.f11) ? (!(_this.f11)) : (_this.f11))) : (!(p0).isEqualTo(p0)));
          let _2274_v123;
          _2274_v123 = _dafny.Seq.UnicodeFromString("lf");
          let _index393 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length));
          (_2268_v117)[_index393] = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_module.__default.fm55(new BigNumber((_2274_v123).length), _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2266_v115,p0)), _this.f11, globalState)).cardinality())), p0, (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))])).length);
          let _2275_v124;
          let _nw337 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Set.Empty);
          _2275_v124 = _nw337;
          let _2276_v125;
          _2276_v125 = _dafny.Seq.of(_2269_v118);
          let _index394 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_2275_v124).length));
          (_2275_v124)[_index394] = (_2276_v125)[_module.__default.safeIndex(p0, new BigNumber((_2276_v125).length))];
          let _2277_v126;
          _2277_v126 = _dafny.Seq.of(_2191_v54);
          let _2278_v127;
          _2278_v127 = _dafny.Seq.of(_2274_v123, _2274_v123);
          let _2279_v128;
          _2279_v128 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f11),_dafny.Seq.of(_this.f11, _this.f11));
          let _index395 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_2275_v124).length));
          let _index396 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length));
          let _rhs358 = _2269_v118;
          let _rhs359 = (_2277_v126)[_module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber(((_2278_v127)[_module.__default.safeIndex(p0, new BigNumber((_2278_v127).length))]).length)).minus(p0)), new BigNumber((_2277_v126).length))];
          let _rhs360 = _module.__default.safeModuloInt((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))]);
          let _rhs361 = _dafny.Seq.UnicodeFromString("e");
          let _rhs362 = ((p0).plus((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))])).isLessThan(new BigNumber((_dafny.Seq.Concat((((_2279_v128).contains(_this.f11)) ? ((_2279_v128).get(_this.f11)) : (_2271_v120)), _2271_v120)).length));
          let _lhs289 = _2275_v124;
          let _lhs290 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_2275_v124).length));
          let _lhs291 = _2268_v117;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length));
          let _lhs293 = _this;
          _lhs289[_lhs290] = _rhs358;
          _2191_v54 = _rhs359;
          _lhs291[_lhs292] = _rhs360;
          _2274_v123 = _rhs361;
          _lhs293.f11 = _rhs362;
          _2268_v117 = _2268_v117;
          let _2280_v129;
          let _init59 = ((_2281_v117) => function (_2282_i17) {
            return _dafny.MultiSet.fromElements((_2281_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2281_v117).length))]);
          })(_2268_v117);
          let _nw338 = Array((new BigNumber(14)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw338.length); _i0_59++) {
            _nw338[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2280_v129 = _nw338;
          let _2283_v130;
          _2283_v130 = _dafny.MultiSet.fromElements(false, _this.f11, _this.f11);
          let _2284_v131;
          _2284_v131 = _dafny.MultiSet.fromElements(new BigNumber(4), p0, (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], (((_2283_v130).contains(_this.f11)) ? ((_2283_v130).get(_this.f11)) : ((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))])));
          let _index397 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2280_v129).length));
          (_2280_v129)[_index397] = _2284_v131;
          let _index398 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_2280_v129).length));
          (_2280_v129)[_index398] = _2284_v131;
        } else {
          let _2285_v132;
          _2285_v132 = _module.D7.create_DC16((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))]);
          let _2286_v133;
          _2286_v133 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC3(_2191_v54),_this.f11);
          let _2287_v134;
          _2287_v134 = _module.D2.create_DC3(_2191_v54);
          let _2288_v135;
          _2288_v135 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2287_v134,true), _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC3(_2191_v54),true));
          let _2289_v136;
          _2289_v136 = _dafny.MultiSet.fromElements(_2266_v115);
          let _2290_v137;
          let _nw339 = Array((new BigNumber(24)).toNumber());
          _nw339[(_dafny.ZERO).toNumber()] = _2286_v133;
          _nw339[(_dafny.ONE).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(2)).toNumber()] = (_2288_v135)[_module.__default.safeIndex(new BigNumber((_2289_v136).cardinality()), new BigNumber((_2288_v135).length))];
          _nw339[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2287_v134,_this.f11);
          _nw339[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2287_v134,false);
          _nw339[(new BigNumber(5)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(6)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(7)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(8)).toNumber()] = (_2286_v133).update(_module.D2.create_DC3(_2191_v54), _this.f11);
          _nw339[(new BigNumber(9)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(10)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(11)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(12)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2287_v134,_this.f11);
          _nw339[(new BigNumber(14)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(15)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(16)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(17)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(18)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(19)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2287_v134,_this.f11);
          _nw339[(new BigNumber(21)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(22)).toNumber()] = _2286_v133;
          _nw339[(new BigNumber(23)).toNumber()] = _2286_v133;
          _2290_v137 = _nw339;
          let _2291_v138;
          _2291_v138 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], (_2285_v132).dtor_cf29),_2290_v137);
          _2291_v138 = (_2291_v138).update((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], _2290_v137);
          let _index399 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index399] = !(_this.f11);
          let _2292_v139;
          _2292_v139 = _dafny.MultiSet.fromElements(true);
          let _2293_v140;
          _2293_v140 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2292_v139);
          let _2294_v141;
          _2294_v141 = _dafny.MultiSet.fromElements(_2273_v122);
          let _2295_v142;
          _2295_v142 = _module.D18.create_DC46(_2294_v141);
          let _2296_v143;
          _2296_v143 = _dafny.Seq.of(new BigNumber(((_2295_v142).dtor_cf89).cardinality()));
          let _index400 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2191_v54).length));
          let _rhs363 = (_dafny.ZERO).minus((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))]);
          let _rhs364 = _dafny.Set.fromElements(((((_2293_v140).contains(_this.f11)) ? ((_2293_v140).get(_this.f11)) : (_2292_v139))).IsDisjointFrom(_dafny.MultiSet.fromElements(false)), !_dafny.areEqual(_module.__default.fm63(false, (_2271_v120)[_module.__default.safeIndex((_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], new BigNumber((_2271_v120).length))], globalState), _2296_v143));
          let _rhs365 = (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))];
          let _rhs366 = _this.f11;
          let _rhs367 = _this.f11;
          let _lhs294 = globalState;
          let _lhs295 = globalState;
          let _lhs296 = _2191_v54;
          let _lhs297 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2191_v54).length));
          let _lhs298 = _this;
          _lhs294.f7 = _rhs363;
          _2269_v118 = _rhs364;
          _lhs295.f7 = _rhs365;
          _lhs296[_lhs297] = _rhs366;
          _lhs298.f11 = _rhs367;
          let _2297_v144;
          let _init60 = ((_2298_v54) => function (_2299_i18) {
            return (_2298_v54)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_2298_v54).length))];
          })(_2191_v54);
          let _nw340 = Array((new BigNumber(10)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw340.length); _i0_60++) {
            _nw340[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2297_v144 = _nw340;
          let _2300_v145;
          _2300_v145 = _module.D2.create_DC4(_2268_v117, (_2191_v54)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_2191_v54).length))], (_2268_v117)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_2268_v117).length))], false, _2297_v144);
          let _2301_v146;
          let _nw341 = Array((new BigNumber(8)).toNumber());
          _nw341[(_dafny.ZERO).toNumber()] = _2268_v117;
          _nw341[(_dafny.ONE).toNumber()] = _2268_v117;
          _nw341[(new BigNumber(2)).toNumber()] = _2268_v117;
          _nw341[(new BigNumber(3)).toNumber()] = _2268_v117;
          _nw341[(new BigNumber(4)).toNumber()] = _2268_v117;
          _nw341[(new BigNumber(5)).toNumber()] = (_2300_v145).dtor_cf6;
          _nw341[(new BigNumber(6)).toNumber()] = _2268_v117;
          _nw341[(new BigNumber(7)).toNumber()] = _2268_v117;
          _2301_v146 = _nw341;
          let _index401 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_2301_v146).length));
          (_2301_v146)[_index401] = _2268_v117;
          let _index402 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_2301_v146).length));
          let _nw342 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          (_2301_v146)[_index402] = _nw342;
          let _index403 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2191_v54).length));
          (_2191_v54)[_index403] = _dafny.Seq.IsPrefixOf(_2271_v120, _2271_v120);
          let _2302_v147;
          let _nw343 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2302_v147 = _nw343;
          let _2303_v148;
          _2303_v148 = _dafny.Seq.UnicodeFromString("psmsf");
          let _index404 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_2302_v147).length));
          (_2302_v147)[_index404] = _2303_v148;
          let _index405 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_2302_v147).length));
          (_2302_v147)[_index405] = _2303_v148;
        }
        let _2304_v149;
        let _init61 = function (_2305_i19) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f11),_this.f11)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f11),_this.f11));
        };
        let _nw344 = Array((new BigNumber(22)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw344.length); _i0_61++) {
          _nw344[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2304_v149 = _nw344;
        let _2306_v150;
        _2306_v150 = _dafny.MultiSet.fromElements(_this.f11);
        let _2307_v151;
        _2307_v151 = _dafny.Map.Empty.slice().updateUnsafe(_2306_v150,_this.f11);
        let _index406 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_2304_v149).length));
        (_2304_v149)[_index406] = _2307_v151;
        let _index407 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_2304_v149).length));
        (_2304_v149)[_index407] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_this.f11, !(_this.f11))).Union(_2306_v150),_this.f11);
      }
      let _2308_v152;
      _2308_v152 = new _dafny.CodePoint('k'.codePointAt(0));
      let _2309_v153;
      _2309_v153 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_2308_v152);
      let _2310_v154;
      _2310_v154 = _dafny.Seq.of(p0, new BigNumber((_dafny.Seq.of(_2309_v153)).length));
      let _2311_v155;
      _2311_v155 = _module.D9.create_DC23((p0).minus((_2310_v154)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_2310_v154).length))]), p0, p0, (p0).plus(p0), new BigNumber(988));
      _2311_v155 = _2311_v155;
      let _2312_v156;
      _2312_v156 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
      let _2313_v157;
      _2313_v157 = _dafny.MultiSet.fromElements(p0);
      let _2314_v158;
      _2314_v158 = _dafny.Seq.UnicodeFromString("gvrrji");
      let _2315_v159;
      _2315_v159 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_2310_v154).length)),_dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11));
      let _2316_v160;
      _2316_v160 = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,_this.f11);
      let _2317_v161;
      _2317_v161 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2314_v158).length),new BigNumber(((((_2315_v159).contains(_2310_v154)) ? ((_2315_v159).get(_2310_v154)) : (_2316_v160))).length));
      let _2318_v162;
      _2318_v162 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,_module.D1.create_DC2(p0, new BigNumber((_2312_v156).length), _this.f11)),(_2313_v157).update(new BigNumber((_2314_v158).length), _module.__default.abs((((_2317_v161).contains(p0)) ? ((_2317_v161).get(p0)) : (p0)))));
      r0 = (_2318_v162).Merge((_2318_v162).Merge(_2318_v162));
      let _2319_v163;
      _2319_v163 = _module.D2.create_DC3(_2191_v54);
      let _pat_let_tv61 = _2191_v54;
      r1 = ((!(_this.f11)) ? (function (_pat_let68_0) {
        return function (_2320_dt__update__tmp_h0) {
          return function (_pat_let69_0) {
            return function (_2321_dt__update_hcf5_h0) {
              return _module.D2.create_DC3(_2321_dt__update_hcf5_h0);
            }(_pat_let69_0);
          }(_pat_let_tv61);
        }(_pat_let68_0);
      }(_2319_v163)) : (_2319_v163));
      return [r0, r1];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      let _source30 = _module.D1.create_DC2(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(260)), function (_2322_i0) {
  return _dafny.Map.Empty.slice().updateUnsafe(true,true);
})).length), new BigNumber(-126), false);
      if (_source30.is_DC2) {
        let _2323___mcc_h0 = (_source30).cf2;
        let _2324___mcc_h1 = (_source30).cf3;
        let _2325___mcc_h2 = (_source30).cf4;
        let _2326_cf4 = _2325___mcc_h2;
        let _2327_cf3 = _2324___mcc_h1;
        let _2328_cf2 = _2323___mcc_h0;
        return (_2328_cf2).minus(_2328_cf2);
      } else {
        let _2329___mcc_h3 = (_source30).cf1;
        let _2330_cf1 = _2329___mcc_h3;
        return new BigNumber(734);
      }
    };
    fm2(globalState) {
      let _this = this;
      if ((new BigNumber((_dafny.Seq.UnicodeFromString("vahhmfdxj")).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(true)).length))) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)));
      } else {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)));
      }
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _2331_v0;
      _2331_v0 = true;
      let _2332_i0;
      _2332_i0 = _dafny.ZERO;
      L15: {
        while (_2331_v0) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2332_i0)) {
              break L15;
            }
            _2332_i0 = (_2332_i0).plus(_dafny.ONE);
            let _index408 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((p1).length));
            (p1)[_index408] = _2331_v0;
            let _index409 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((p1).length));
            (p1)[_index409] = (new BigNumber(51)).isLessThan(p0);
            r3 = p0;
            let _2333_v1;
            let _init62 = ((_2334_p1) => function (_2335_i1) {
              return (((_2334_p1)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_2334_p1).length))]) ? (_module.D1.create_DC1(_dafny.Seq.UnicodeFromString("hrlebc"))) : (_module.D1.create_DC1(_dafny.Seq.UnicodeFromString("houbdlof"))));
            })(p1);
            let _nw345 = Array((new BigNumber(17)).toNumber());
            for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw345.length); _i0_62++) {
              _nw345[_i0_62] = _init62(new BigNumber(_i0_62));
            }
            _2333_v1 = _nw345;
            let _2336_v2;
            _2336_v2 = _dafny.Seq.UnicodeFromString("rpfcdldk");
            let _2337_v3;
            _2337_v3 = _module.D1.create_DC1(_2336_v2);
            let _index410 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_2333_v1).length));
            (_2333_v1)[_index410] = _2337_v3;
            let _2338_v4;
            _2338_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2331_v0,p0);
            let _2339_v5;
            _2339_v5 = _dafny.MultiSet.fromElements((_2338_v4).update(_module.__default.fm3(p0, globalState), p0));
            let _2340_v6;
            let _init63 = ((_2341_p0) => function (_2342_i2) {
              return _module.__default.safeModuloInt(_2342_i2, _2341_p0);
            })(p0);
            let _nw346 = Array((new BigNumber(13)).toNumber());
            for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw346.length); _i0_63++) {
              _nw346[_i0_63] = _init63(new BigNumber(_i0_63));
            }
            _2340_v6 = _nw346;
            let _2343_v7;
            _2343_v7 = _dafny.Seq.of(_2340_v6);
            let _index411 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_2333_v1).length));
            let _rhs368 = _2337_v3;
            let _rhs369 = _2339_v5;
            let _rhs370 = (_2337_v3).dtor_cf1;
            let _rhs371 = (p0).isLessThanOrEqualTo((p0).plus(new BigNumber((_2336_v2).length)));
            let _rhs372 = (_2343_v7)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2343_v7).length))];
            let _lhs299 = _2333_v1;
            let _lhs300 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_2333_v1).length));
            _lhs299[_lhs300] = _rhs368;
            _2339_v5 = _rhs369;
            _2336_v2 = _rhs370;
            r2 = _rhs371;
            _2340_v6 = _rhs372;
            r0 = p0;
          }
        }
      }
      let _2344_v8;
      let _nw347 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _2344_v8 = _nw347;
      let _2345_v9;
      _2345_v9 = _module.D2.create_DC4(_2344_v8, _2331_v0, p0, _2331_v0, p1);
      let _pat_let_tv62 = p0;
      let _pat_let_tv63 = _2331_v0;
      let _pat_let_tv64 = p1;
      _2345_v9 = function (_pat_let70_0) {
        return function (_2346_dt__update__tmp_h0) {
          return function (_pat_let71_0) {
            return function (_2347_dt__update_hcf8_h0) {
              return function (_pat_let72_0) {
                return function (_2348_dt__update_hcf9_h0) {
                  return function (_pat_let73_0) {
                    return function (_2349_dt__update_hcf10_h0) {
                      return _module.D2.create_DC4((_2346_dt__update__tmp_h0).dtor_cf6, (_2346_dt__update__tmp_h0).dtor_cf7, _2347_dt__update_hcf8_h0, _2348_dt__update_hcf9_h0, _2349_dt__update_hcf10_h0);
                    }(_pat_let73_0);
                  }(_pat_let_tv64);
                }(_pat_let72_0);
              }(!(_pat_let_tv63));
            }(_pat_let71_0);
          }(_pat_let_tv62);
        }(_pat_let70_0);
      }(_2345_v9);
      let _2350_v10;
      _2350_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2331_v0,!(_2331_v0) || (_2331_v0));
      _2350_v10 = (_2350_v10).update(_2331_v0, _2331_v0);
      r2 = _2331_v0;
      let _2351_v11;
      _2351_v11 = _dafny.Seq.of(p0);
      let _hi16 = (_2351_v11)[_module.__default.safeIndex(p0, new BigNumber((_2351_v11).length))];
      for (let _2352_i3 = p0; _2352_i3.isLessThan(_hi16); _2352_i3 = _2352_i3.plus(_dafny.ONE)) {
        let _2353_v12;
        let _init64 = function (_2354_i4) {
          return _dafny.Seq.UnicodeFromString("ouwtuvdwo");
        };
        let _nw348 = Array((new BigNumber(16)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw348.length); _i0_64++) {
          _nw348[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _2353_v12 = _nw348;
        let _2355_v13;
        _2355_v13 = _dafny.Seq.UnicodeFromString("mm");
        let _2356_v14;
        _2356_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2353_v12,new BigNumber((_dafny.Seq.Concat(_2355_v13, _2355_v13)).length));
        _2356_v14 = (_2356_v14).update(_2353_v12, p0);
        let _pat_let_tv65 = _2355_v13;
        let _source31 = function (_pat_let74_0) {
          return function (_2357_dt__update__tmp_h1) {
            return function (_pat_let75_0) {
              return function (_2358_dt__update_hcf1_h0) {
                return _module.D1.create_DC1(_2358_dt__update_hcf1_h0);
              }(_pat_let75_0);
            }(_pat_let_tv65);
          }(_pat_let74_0);
        }(_module.D1.create_DC1(_2355_v13));
        if (_source31.is_DC2) {
          let _2359___mcc_h0 = (_source31).cf2;
          let _2360___mcc_h1 = (_source31).cf3;
          let _2361___mcc_h2 = (_source31).cf4;
          let _2362_cf4 = _2361___mcc_h2;
          let _2363_cf3 = _2360___mcc_h1;
          let _2364_cf2 = _2359___mcc_h0;
          r3 = _2364_cf2;
          let _2365_v15;
          _2365_v15 = _dafny.MultiSet.fromElements(new BigNumber(931), new BigNumber((_dafny.Seq.UnicodeFromString("dhfqdpmbx")).length));
          let _2366_v16;
          _2366_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2365_v15).IsSubsetOf(_2365_v15),p0);
          _2366_v16 = (_2366_v16).update(_2331_v0, _2352_i3);
          let _2367_v17;
          _2367_v17 = _module.D1.create_DC1(_2355_v13);
          let _2368_v18;
          _2368_v18 = _module.D1.create_DC2(_2364_cf2, _2363_cf3, _2331_v0);
          let _2369_v19;
          let _nw349 = new _module.C6();
          _nw349.__ctor(_2367_v17, _2362_cf4, p1, _dafny.Set.fromElements(_2363_cf3), _2368_v18, _2362_cf4);
          _2369_v19 = _nw349;
          let _2370_v20;
          _2370_v20 = _dafny.Seq.of(_2369_v19, _2369_v19, _2369_v19, _2369_v19);
          let _2371_v21;
          _2371_v21 = _dafny.Seq.of(_2331_v0);
          let _2372_v22;
          _2372_v22 = _dafny.Set.fromElements(p0, _2364_cf2, _module.__default.fm0(_2371_v21, _2331_v0, true, _2352_i3, globalState));
          let _2373_v23;
          let _nw350 = new _module.C2();
          _nw350.__ctor(new BigNumber((_2370_v20).length), _2362_cf4, _2367_v17, _2362_cf4, p1, _2372_v22, _2368_v18, _2362_cf4);
          _2373_v23 = _nw350;
          let _2374_v24;
          _2374_v24 = _dafny.Set.fromElements(_2373_v23);
          let _2375_v25;
          _2375_v25 = _dafny.Seq.of(_2374_v24);
          let _2376_v26;
          _2376_v26 = new _dafny.CodePoint('l'.codePointAt(0));
          let _2377_v27;
          let _nw351 = new _module.C0();
          _nw351.__ctor((_dafny.Set.fromElements(_2373_v23, _2373_v23)).IsProperSubsetOf((_2375_v25)[_module.__default.safeIndex((_2373_v23).f21, new BigNumber((_2375_v25).length))]), _2376_v26, _2368_v18, (_2373_v23).f22);
          _2377_v27 = _nw351;
          let _2378_v28;
          _2378_v28 = _dafny.Seq.of(_2351_v11);
          let _2379_v29;
          _2379_v29 = _module.D11.create_DC26(_2344_v8, _2355_v13);
          let _2380_v30;
          _2380_v30 = _dafny.Map.Empty.slice().updateUnsafe(!((_2352_i3).isLessThanOrEqualTo((_module.D13.create_DC34(_2352_i3, _2378_v28)).dtor_cf71)),_2379_v29);
          _2380_v30 = (_2380_v30).update(_2331_v0, _2379_v29);
        } else {
          let _2381___mcc_h3 = (_source31).cf1;
          let _2382_cf1 = _2381___mcc_h3;
          let _2383_v31;
          _2383_v31 = _dafny.Set.fromElements(_2352_i3);
          let _2384_v32;
          _2384_v32 = _module.D1.create_DC2(_2352_i3, new BigNumber(-87), _2331_v0);
          let _2385_v33;
          let _nw352 = new _module.C2();
          _nw352.__ctor(p0, _2331_v0, _module.D1.create_DC1(_2382_cf1), _2331_v0, p1, _2383_v31, _2384_v32, _2331_v0);
          _2385_v33 = _nw352;
          let _2386_v34;
          _2386_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2352_i3,_2385_v33);
          _2386_v34 = _2386_v34;
          let _2387_v35;
          let _nw353 = new _module.C4();
          _nw353.__ctor();
          _2387_v35 = _nw353;
          let _2388_v36;
          let _nw354 = Array((new BigNumber(2)).toNumber()).fill(_module.D3.Default());
          _2388_v36 = _nw354;
          let _2389_v37;
          _2389_v37 = _module.D3.create_DC8(p0);
          let _index412 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_2388_v36).length));
          (_2388_v36)[_index412] = _2389_v37;
          let _index413 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_2388_v36).length));
          (_2388_v36)[_index413] = _2389_v37;
          let _2390_v38;
          _2390_v38 = _dafny.MultiSet.fromElements(new BigNumber((_2351_v11).length), new BigNumber((_2383_v31).length), new BigNumber(-180), _2352_i3);
          let _2391_v40;
          let _nw355 = new _module.C2();
          _nw355.__ctor(_2352_i3, false, _2385_v33.f16, !(_2331_v0), p1, _2383_v31, _module.D1.create_DC2(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_2392_i5) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length), new BigNumber((_2390_v38).cardinality()), (_2385_v33).f13), _dafny.areEqual(_dafny.Seq.of(new BigNumber((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(186), new BigNumber(82))) {
              let _2393_v39 = _compr_59;
              if (((new BigNumber(186)).isLessThanOrEqualTo(_2393_v39)) && ((_2393_v39).isLessThan(new BigNumber(82)))) {
                _coll59.push([_module.__default.safeDivisionInt(_2393_v39, _2352_i3),(_module.D15.create_DC39(p0, new _dafny.CodePoint('q'.codePointAt(0)), _2355_v13, p0, _2352_i3)).dtor_cf82]);
              }
            }
            return _coll59;
          }()).length)), _2351_v11));
          _2391_v40 = _nw355;
          let _2394_v41;
          _2394_v41 = _module.D9.create_DC23(new BigNumber(570), (_2352_i3).plus(p0), (_2391_v40).f21, _2352_i3, (_2391_v40).f21);
          let _rhs373 = _2391_v40;
          let _rhs374 = _2394_v41;
          _2391_v40 = _rhs373;
          _2394_v41 = _rhs374;
        }
        let _index414 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2344_v8).length));
        (_2344_v8)[_index414] = (_dafny.ZERO).minus(_2352_i3);
        let _2395_v42;
        _2395_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2331_v0,_2355_v13);
        let _2396_v43;
        _2396_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_2395_v42).contains(!(_2331_v0))) ? ((_2395_v42).get(!(_2331_v0))) : (_2355_v13)));
        let _index415 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2344_v8).length));
        (_2344_v8)[_index415] = (_module.__default.safeModuloInt(_2352_i3, new BigNumber((_2396_v43).length))).plus(p0);
        let _2397_v44;
        _2397_v44 = new _dafny.CodePoint('a'.codePointAt(0));
        let _2398_v45;
        _2398_v45 = _module.D1.create_DC2(_2352_i3, p0, false);
        let _2399_v46;
        let _nw356 = new _module.C0();
        _nw356.__ctor(_2331_v0, _2397_v44, _2398_v45, _2331_v0);
        _2399_v46 = _nw356;
        let _2400_v47;
        _2400_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2399_v46.f19,_2399_v46);
        let _2401_v48;
        _2401_v48 = _module.D19.create_DC48(_2399_v46);
        let _2402_v49;
        let _nw357 = Array((new BigNumber(25)).toNumber());
        _nw357[(_dafny.ZERO).toNumber()] = _2399_v46;
        _nw357[(_dafny.ONE).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(2)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(3)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(4)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(5)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(6)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(7)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(8)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(9)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(10)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(11)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(12)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(13)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(14)).toNumber()] = (((_2400_v47).contains(true)) ? ((_2400_v47).get(true)) : (_2399_v46));
        _nw357[(new BigNumber(15)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(16)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(17)).toNumber()] = ((_2331_v0) ? (_2399_v46) : (_2399_v46));
        _nw357[(new BigNumber(18)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(19)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(20)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(21)).toNumber()] = (_2401_v48).dtor_cf92;
        _nw357[(new BigNumber(22)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(23)).toNumber()] = _2399_v46;
        _nw357[(new BigNumber(24)).toNumber()] = _2399_v46;
        _2402_v49 = _nw357;
        _2402_v49 = _2402_v49;
      }
      let _2403_v50;
      let _nw358 = new _module.C4();
      _nw358.__ctor();
      _2403_v50 = _nw358;
      let _2404_v51;
      _2404_v51 = _dafny.MultiSet.fromElements(p0, p0, (_dafny.ZERO).minus(p0), (_dafny.ZERO).minus(p0), p0);
      r0 = (_2351_v11)[_module.__default.safeIndex((((_2404_v51).contains(new BigNumber(653))) ? ((_2404_v51).get(new BigNumber(653))) : (p0)), new BigNumber((_2351_v11).length))];
      r1 = _2331_v0;
      r2 = (((_2350_v10).contains(true)) ? ((_2350_v10).get(true)) : ((p0).isLessThan(new BigNumber(940))));
      let _2405_v52;
      _2405_v52 = _dafny.Seq.of(_2331_v0);
      let _2406_v53;
      _2406_v53 = _dafny.Seq.UnicodeFromString("sigql");
      let _2407_v54;
      _2407_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(new BigNumber((_2406_v53).length)));
      r3 = (new BigNumber((_dafny.Seq.Concat(_2405_v52, _2405_v52)).length)).multipliedBy(((_2331_v0) ? (p0) : ((_dafny.ZERO).minus((((_2407_v54).contains(p0)) ? ((_2407_v54).get(p0)) : ((_dafny.ZERO).minus(p0)))))));
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
