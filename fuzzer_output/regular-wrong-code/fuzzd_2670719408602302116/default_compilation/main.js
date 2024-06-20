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
      return _dafny.MultiSet.fromElements(!(!(!(!(!(true))))) || ((_module.D2.create_DC6((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(797)), function (_0_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length))).length)), true, !(true))).dtor_cf8), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("fqkb"), _dafny.Seq.UnicodeFromString("ocpjiygo")), !(new BigNumber(-380)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))).length))));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), function (_1_i0) {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-386)), function (_2_i1) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("xk"))).length);
      });
    };
    static fm5(globalState) {
      return _module.D0.create_DC3(((true) ? (_module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(189))) : (_module.D0.create_DC3(_module.D0.create_DC2()))));
    };
    static fm6(p0, globalState) {
      return new BigNumber(84);
    };
    static fm7(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_3_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), new _dafny.CodePoint('y'.codePointAt(0))),new _dafny.CodePoint('l'.codePointAt(0)));
    };
    static fm11(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(false), _dafny.Seq.of(!(false), true, true, true)),new _dafny.CodePoint('k'.codePointAt(0)));
    };
    static fm14(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("g");
    };
    static fm17(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("elk"), _dafny.Seq.UnicodeFromString("f")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("f"), _dafny.Seq.UnicodeFromString("jflao")));
    };
    static fm18(p0, p1, globalState) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    };
    static fm20(p0, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jos"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_4_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("ilgan")))).length);
    };
    static fm21(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,false);
    };
    static fm22(p0, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(376)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("akm"),!(true))).length)));
    };
    static fm23(p0, globalState) {
      let _source0 = _dafny.Seq.of(new BigNumber(((_module.D5.create_DC13((_module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(true,false))).dtor_cf19)).dtor_cf19).length), new BigNumber(382));
      let _5___mcc_h0 = _source0;
      let _6_cf24 = _5___mcc_h0;
      return _dafny.Seq.Concat(_dafny.Seq.of(true), (_module.D9.create_DC21(false, _dafny.Seq.of(true), true)).dtor_cf27);
    };
    static fm24(globalState) {
      return _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(216),new BigNumber(372)));
    };
    static fm25(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-715)), function (_7_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("unujxwuqx"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(37)), function (_8_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })));
    };
    static fm26(p0, p1, p2, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber(-978), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_9_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length)))).multipliedBy(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length), new BigNumber(312)));
    };
    static fm27(p0, globalState) {
      let _source1 = _module.D19.create_DC39(false, new BigNumber(106));
      if (_source1.is_DC39) {
        let _10___mcc_h0 = (_source1).cf58;
        let _11___mcc_h1 = (_source1).cf59;
        let _12_cf59 = _11___mcc_h1;
        let _13_cf58 = _10___mcc_h0;
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(506), new BigNumber((_dafny.MultiSet.fromElements(_12_cf59)).cardinality()), _12_cf59, new BigNumber((_dafny.MultiSet.fromElements(_12_cf59, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_13_cf58,_12_cf59)).length))).cardinality()), _12_cf59)).cardinality()), _12_cf59, (_dafny.ZERO).minus(_12_cf59)));
      } else if (_source1.is_DC38) {
        let _14___mcc_h2 = (_source1).cf57;
        let _15_cf57 = _14___mcc_h2;
        return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, true, true, !(false), true)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("kh")).length), new BigNumber(-878), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-273))).length))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(939))).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber(-120))).length)));
      } else {
        let _16___mcc_h3 = (_source1).cf60;
        let _17_cf60 = _16___mcc_h3;
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(687)))).Difference(_dafny.MultiSet.fromElements((_module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-72)))).length))).dtor_cf0, new BigNumber(919)));
      }
    };
    static fm28(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).Union(_dafny.Set.fromElements(new BigNumber(564), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(2), new BigNumber(844))) {
          let _18_v0 = _compr_0;
          if (((new BigNumber(2)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(844)))) {
            _coll0.add(_module.__default.safeModuloInt(_18_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).length)));
          }
        }
        return _coll0;
      }()).length), new BigNumber(277))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), function (_19_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("qthfycdk")).length);
      }),!(!(false)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("heoi")).length))).length), new BigNumber(387), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality()))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(false, false, false);
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC11(!((_dafny.MultiSet.fromElements(new BigNumber(-407))).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(708), new BigNumber(789), new BigNumber(-937)))), !(new BigNumber(210)).isEqualTo(new BigNumber(-759)), new BigNumber(842), _dafny.Seq.Concat((_module.D10.create_DC24(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-675))).length), true, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))))).dtor_cf33, _dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_20_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})));
    };
    static fm31(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(true, false)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D10.create_DC23(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sisutg")));
      if (_source2.is_DC24) {
        let _21___mcc_h0 = (_source2).cf31;
        let _22___mcc_h1 = (_source2).cf32;
        let _23___mcc_h2 = (_source2).cf33;
        let _24_cf33 = _23___mcc_h2;
        let _25_cf32 = _22___mcc_h1;
        let _26_cf31 = _21___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_25_cf32,_dafny.Map.Empty.slice().updateUnsafe(_25_cf32,_26_cf31));
      } else {
        let _27___mcc_h3 = (_source2).cf30;
        let _28_cf30 = _27___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC7(new _dafny.CodePoint('n'.codePointAt(0)), true, new BigNumber(-255))).dtor_cf10,_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(275)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-616))));
      }
    };
    static fm33(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true), !(false)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),false));
    };
    static fm34(p0, p1, globalState) {
      return _module.D0.create_DC0(new BigNumber(997));
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(false, true)).Difference(_dafny.Set.fromElements(!(true), !(true), true, false, (_module.D19.create_DC39(false, new BigNumber(831))).dtor_cf58))).Union((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(!(true), !(false))));
    };
    static fm36(p0, p1, p2, globalState) {
      return _module.D10.create_DC23(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hx"), _dafny.Seq.UnicodeFromString("kueav")));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return ((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(false))).Elements) {
          let _29_v0 = _compr_1;
          if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(false))).contains(_29_v0)) {
            _coll1.push([_29_v0,new BigNumber(737)]);
          }
        }
        return _coll1;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, true),new BigNumber(790)))).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber((_dafny.Seq.of(false)).length))).Keys.Elements) {
          let _30_v1 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber((_dafny.Seq.of(false)).length))).contains(_30_v1)) {
            _coll2.push([_30_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)]);
          }
        }
        return _coll2;
      }());
    };
    static fm38(p0, globalState) {
      return _module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC1(new BigNumber(-725), new BigNumber(483)))));
    };
    static fm39(p0, p1, p2, globalState) {
      let _source3 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(-707))).length));
      let _31___mcc_h0 = _source3;
      let _32_cf24 = _31___mcc_h0;
      return _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(48)));
    };
    static fm40(p0, globalState) {
      let _source4 = _module.D7.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("gndcte")).length),!(true)));
      if (_source4.is_DC18) {
        if (true) {
          return _module.D2.create_DC7(new _dafny.CodePoint('j'.codePointAt(0)), !(true), new BigNumber((_dafny.Seq.UnicodeFromString("mhiixh")).length));
        } else {
          return _module.D2.create_DC7(new _dafny.CodePoint('p'.codePointAt(0)), true, new BigNumber(-733));
        }
      } else {
        let _33___mcc_h0 = (_source4).cf23;
        let _34_cf23 = _33___mcc_h0;
        return _module.D2.create_DC7(new _dafny.CodePoint('v'.codePointAt(0)), false, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality())))).length));
      }
    };
    static fm41(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("fvh"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sfclqlf"), _dafny.Seq.UnicodeFromString("vxt")));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      let _source5 = ((true) ? (_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(!(true))).length)))) : (_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-332)))));
      if (_source5.is_DC27) {
        let _35___mcc_h0 = (_source5).cf36;
        let _36___mcc_h1 = (_source5).cf37;
        let _37___mcc_h2 = (_source5).cf38;
        let _38___mcc_h3 = (_source5).cf39;
        let _39_cf39 = _38___mcc_h3;
        let _40_cf38 = _37___mcc_h2;
        let _41_cf37 = _36___mcc_h1;
        let _42_cf36 = _35___mcc_h0;
        return (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fup"), _dafny.Seq.UnicodeFromString("xmp"))).Elements) {
              let _43_v1 = _compr_4;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fup"), _dafny.Seq.UnicodeFromString("xmp"))).contains(_43_v1)) {
                _coll4.push([_43_v1,true]);
              }
            }
            return _coll4;
          }()).Keys.Elements) {
            let _44_v0 = _compr_3;
            if ((function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fup"), _dafny.Seq.UnicodeFromString("xmp"))).Elements) {
                let _43_v1 = _compr_5;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fup"), _dafny.Seq.UnicodeFromString("xmp"))).contains(_43_v1)) {
                  _coll5.push([_43_v1,true]);
                }
              }
              return _coll5;
            }()).contains(_44_v0)) {
              _coll3.push([_44_v0,true]);
            }
          }
          return _coll3;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mbdnvowfp"),false));
      } else if (_source5.is_DC26) {
        let _45___mcc_h4 = (_source5).cf35;
        let _46_cf35 = _45___mcc_h4;
        return function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("c"), _dafny.Seq.UnicodeFromString("xegsurtcy"))).Elements) {
            let _47_v2 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("c"), _dafny.Seq.UnicodeFromString("xegsurtcy")), _47_v2)) {
              _coll6.push([_47_v2,true]);
            }
          }
          return _coll6;
        }();
      } else {
        let _48___mcc_h5 = (_source5).cf40;
        let _49_cf40 = _48___mcc_h5;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("s"),false)).Merge(function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rasmegb"),true)).Keys.Elements) {
            let _50_v3 = _compr_7;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rasmegb"),true)).contains(_50_v3)) {
              _coll7.push([_50_v3,true]);
            }
          }
          return _coll7;
        }());
      }
    };
    static fm43(p0, p1, p2, globalState) {
      if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(315),new _dafny.CodePoint('o'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(6))).length),new _dafny.CodePoint('s'.codePointAt(0))));
      } else {
        return (function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-535),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_51_i0) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
            })).length))).Keys.Elements) {
              let _52_v1 = _compr_9;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-535),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_51_i0) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
              })).length))).contains(_52_v1)) {
                _coll9.push([(_52_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("nkdcjuouw")).length)),(_dafny.ZERO).minus(new BigNumber(-371))]);
              }
            }
            return _coll9;
          }()).Keys.Elements) {
            let _53_v0 = _compr_8;
            if ((function () {
              let _coll10 = new _dafny.Map();
              for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-535),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_51_i0) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
              })).length))).Keys.Elements) {
                let _52_v1 = _compr_10;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-535),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_51_i0) {
                  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length);
                })).length))).contains(_52_v1)) {
                  _coll10.push([(_52_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("nkdcjuouw")).length)),(_dafny.ZERO).minus(new BigNumber(-371))]);
                }
              }
              return _coll10;
            }()).contains(_53_v0)) {
              _coll8.push([(_53_v0).multipliedBy(new BigNumber((_dafny.Seq.of(!(true))).length)),new _dafny.CodePoint('e'.codePointAt(0))]);
            }
          }
          return _coll8;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(827),new _dafny.CodePoint('b'.codePointAt(0))));
      }
    };
    static fm44(p0, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(false, true), _dafny.Seq.of(!(true)));
    };
    static fm45(p0, p1, p2, globalState) {
      return (((!(!(true))) ? (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(754))),new BigNumber(665)), _dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(202))),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(198))).length)))) : (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-239))),new BigNumber(616)), function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Set.fromElements(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))))).Elements) {
          let _54_v0 = _compr_11;
          if ((_dafny.Set.fromElements(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))))).contains(_54_v0)) {
            _coll11.push([_54_v0,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-170)), function (_55_i0) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll11;
      }())))).Difference((_dafny.Set.fromElements(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(333))).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_module.D4.create_DC11(true, false, new BigNumber(251), _dafny.Seq.UnicodeFromString("uid")))).cardinality()))).length))).length)))),function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length),new _dafny.CodePoint('t'.codePointAt(0)))).Keys.Elements) {
            let _56_v2 = _compr_13;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length),new _dafny.CodePoint('t'.codePointAt(0)))).contains(_56_v2)) {
              _coll13.push([(_56_v2).plus(new BigNumber(842)),new _dafny.CodePoint('s'.codePointAt(0))]);
            }
          }
          return _coll13;
        }())).Keys.Elements) {
          let _57_v1 = _compr_12;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(333))).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_module.D4.create_DC11(true, false, new BigNumber(251), _dafny.Seq.UnicodeFromString("uid")))).cardinality()))).length))).length)))),function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length),new _dafny.CodePoint('t'.codePointAt(0)))).Keys.Elements) {
              let _56_v2 = _compr_14;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length),new _dafny.CodePoint('t'.codePointAt(0)))).contains(_56_v2)) {
                _coll14.push([(_56_v2).plus(new BigNumber(842)),new _dafny.CodePoint('s'.codePointAt(0))]);
              }
            }
            return _coll14;
          }())).contains(_57_v1)) {
            _coll12.push([_57_v1,new BigNumber(230)]);
          }
        }
        return _coll12;
      }())).Intersect(function () {
        let _coll15 = new _dafny.Set();
        for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of (function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (function () {
              let _coll18 = new _dafny.Map();
              for (const _compr_18 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll19 = new _dafny.Map();
  for (const _compr_19 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_19;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll19.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll19;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                let _59_v5 = _compr_18;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_20;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll20.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll20;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                  _coll18.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                }
              }
              return _coll18;
            }()).Keys.Elements) {
              let _60_v4 = _compr_17;
              if ((function () {
                let _coll21 = new _dafny.Map();
                for (const _compr_21 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_22;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll22.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll22;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                  let _59_v5 = _compr_21;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll23 = new _dafny.Map();
  for (const _compr_23 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_23;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll23.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll23;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                    _coll21.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                  }
                }
                return _coll21;
              }()).contains(_60_v4)) {
                _coll17.push([_60_v4,new BigNumber((function () {
                  let _coll24 = new _dafny.Map();
                  for (const _compr_24 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
                    let _61_v7 = _compr_24;
                    if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).contains(_61_v7)) {
                      _coll24.push([_61_v7,true]);
                    }
                  }
                  return _coll24;
                }()).length)]);
              }
            }
            return _coll17;
          }()).Keys.Elements) {
            let _62_v3 = _compr_16;
            if ((function () {
              let _coll25 = new _dafny.Map();
              for (const _compr_25 of (function () {
                let _coll26 = new _dafny.Map();
                for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll27 = new _dafny.Map();
  for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_27;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll27.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll27;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                  let _59_v5 = _compr_26;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_28;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll28.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll28;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                    _coll26.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                  }
                }
                return _coll26;
              }()).Keys.Elements) {
                let _60_v4 = _compr_25;
                if ((function () {
                  let _coll29 = new _dafny.Map();
                  for (const _compr_29 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll30 = new _dafny.Map();
  for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_30;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll30.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll30;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                    let _59_v5 = _compr_29;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll31 = new _dafny.Map();
  for (const _compr_31 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_31;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll31.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll31;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                      _coll29.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                    }
                  }
                  return _coll29;
                }()).contains(_60_v4)) {
                  _coll25.push([_60_v4,new BigNumber((function () {
                    let _coll32 = new _dafny.Map();
                    for (const _compr_32 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
                      let _61_v7 = _compr_32;
                      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).contains(_61_v7)) {
                        _coll32.push([_61_v7,true]);
                      }
                    }
                    return _coll32;
                  }()).length)]);
                }
              }
              return _coll25;
            }()).contains(_62_v3)) {
              _coll16.push([_62_v3,new BigNumber(754)]);
            }
          }
          return _coll16;
        }(),new _dafny.CodePoint('m'.codePointAt(0)))).Keys.Elements) {
          let _63_v8 = _compr_15;
          if ((_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (function () {
              let _coll34 = new _dafny.Map();
              for (const _compr_34 of (function () {
                let _coll35 = new _dafny.Map();
                for (const _compr_35 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_36;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll36.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll36;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                  let _59_v5 = _compr_35;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll37 = new _dafny.Map();
  for (const _compr_37 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_37;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll37.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll37;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                    _coll35.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                  }
                }
                return _coll35;
              }()).Keys.Elements) {
                let _60_v4 = _compr_34;
                if ((function () {
                  let _coll38 = new _dafny.Map();
                  for (const _compr_38 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll39 = new _dafny.Map();
  for (const _compr_39 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_39;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll39.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll39;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                    let _59_v5 = _compr_38;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll40 = new _dafny.Map();
  for (const _compr_40 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_40;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll40.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll40;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                      _coll38.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                    }
                  }
                  return _coll38;
                }()).contains(_60_v4)) {
                  _coll34.push([_60_v4,new BigNumber((function () {
                    let _coll41 = new _dafny.Map();
                    for (const _compr_41 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
                      let _61_v7 = _compr_41;
                      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).contains(_61_v7)) {
                        _coll41.push([_61_v7,true]);
                      }
                    }
                    return _coll41;
                  }()).length)]);
                }
              }
              return _coll34;
            }()).Keys.Elements) {
              let _62_v3 = _compr_33;
              if ((function () {
                let _coll42 = new _dafny.Map();
                for (const _compr_42 of (function () {
                  let _coll43 = new _dafny.Map();
                  for (const _compr_43 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll44 = new _dafny.Map();
  for (const _compr_44 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_44;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll44.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll44;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                    let _59_v5 = _compr_43;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll45 = new _dafny.Map();
  for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_45;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll45.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll45;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                      _coll43.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                    }
                  }
                  return _coll43;
                }()).Keys.Elements) {
                  let _60_v4 = _compr_42;
                  if ((function () {
                    let _coll46 = new _dafny.Map();
                    for (const _compr_46 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll47 = new _dafny.Map();
  for (const _compr_47 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_47;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll47.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll47;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).Keys.Elements) {
                      let _59_v5 = _compr_46;
                      if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll48 = new _dafny.Map();
  for (const _compr_48 of _dafny.IntegerRange(new BigNumber(-428), new BigNumber(530))) {
    let _58_v6 = _compr_48;
    if (((new BigNumber(-428)).isLessThanOrEqualTo(_58_v6)) && ((_58_v6).isLessThan(new BigNumber(530)))) {
      _coll48.push([_module.__default.safeDivisionInt(_58_v6, new BigNumber(-206)),new BigNumber((_dafny.Seq.UnicodeFromString("daydimu")).length)]);
    }
  }
  return _coll48;
}()).length))),new _dafny.CodePoint('d'.codePointAt(0)))).contains(_59_v5)) {
                        _coll46.push([_59_v5,_dafny.Seq.UnicodeFromString("e")]);
                      }
                    }
                    return _coll46;
                  }()).contains(_60_v4)) {
                    _coll42.push([_60_v4,new BigNumber((function () {
                      let _coll49 = new _dafny.Map();
                      for (const _compr_49 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
                        let _61_v7 = _compr_49;
                        if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))).contains(_61_v7)) {
                          _coll49.push([_61_v7,true]);
                        }
                      }
                      return _coll49;
                    }()).length)]);
                  }
                }
                return _coll42;
              }()).contains(_62_v3)) {
                _coll33.push([_62_v3,new BigNumber(754)]);
              }
            }
            return _coll33;
          }(),new _dafny.CodePoint('m'.codePointAt(0)))).contains(_63_v8)) {
            _coll15.add(_63_v8);
          }
        }
        return _coll15;
      }()));
    };
    static fm46(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(424)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(!(!(true))))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(470))));
    };
    static fm47(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))).length)),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), function (_64_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("yjv")));
    };
    static fm48(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((false) ? (true) : (!(false))),_module.D7.create_DC18());
    };
    static fm49(p0, globalState) {
      return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(959)), function (_65_i0) {
        return new BigNumber(926);
      })).length)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-701))).cardinality()));
    };
    static fm50(globalState) {
      if (true) {
        return _module.D3.create_DC9();
      } else {
        return _module.D3.create_DC9();
      }
    };
    static fm51(p0, globalState) {
      return false;
    };
    static fm52(globalState) {
      let _source6 = _dafny.MultiSet.FromArray(_dafny.Seq.of(true));
      let _66___mcc_h0 = _source6;
      let _67_cf34 = _66___mcc_h0;
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(369),new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll50 = new _dafny.Set();
        for (const _compr_50 of _dafny.IntegerRange(new BigNumber(944), new BigNumber(690))) {
          let _68_v0 = _compr_50;
          if (((new BigNumber(944)).isLessThanOrEqualTo(_68_v0)) && ((_68_v0).isLessThan(new BigNumber(690)))) {
            _coll50.add(_module.__default.safeDivisionInt(_68_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.UnicodeFromString("kgpvcbka")).length))).length)));
          }
        }
        return _coll50;
      }()).length))).length));
    };
    static fm53(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kc")).length),false), function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of _dafny.IntegerRange(new BigNumber(-717), new BigNumber(130))) {
          let _69_v0 = _compr_51;
          if (((new BigNumber(-717)).isLessThanOrEqualTo(_69_v0)) && ((_69_v0).isLessThan(new BigNumber(130)))) {
            _coll51.push([(_69_v0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), function (_70_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length)),true]);
          }
        }
        return _coll51;
      }())).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll52 = new _dafny.Set();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(634), new BigNumber(678))) {
          let _71_v1 = _compr_52;
          if (((new BigNumber(634)).isLessThanOrEqualTo(_71_v1)) && ((_71_v1).isLessThan(new BigNumber(678)))) {
            _coll52.add(_module.__default.safeModuloInt(_71_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D0.create_DC0(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("oghwvum")).length)))).length)))).length)));
          }
        }
        return _coll52;
      }()).length),true)));
    };
    static fm54(p0, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(((true) ? (_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))) : (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-406)), function (_72_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })));
    };
    static m15(globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let r3 = undefined;
      let _73_v0;
      _73_v0 = _dafny.Seq.UnicodeFromString("dpud");
      let _74_v1;
      _74_v1 = new BigNumber(36);
      let _hi0 = _74_v1;
      for (let _75_i0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bheo"), _73_v0)).length); _75_i0.isLessThan(_hi0); _75_i0 = _75_i0.plus(_dafny.ONE)) {
        let _76_v2;
        _76_v2 = true;
        let _77_v3;
        let _nw0 = Array((_dafny.ONE).toNumber());
        _nw0[(_dafny.ZERO).toNumber()] = _76_v2;
        _77_v3 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_77_v3).length));
        (_77_v3)[_index0] = _76_v2;
        let _index1 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_77_v3).length));
        (_77_v3)[_index1] = _76_v2;
        (globalState).f8 = _75_i0;
        let _78_v4;
        _78_v4 = _dafny.Seq.of(_76_v2, _76_v2);
        let _79_v5;
        _79_v5 = _dafny.MultiSet.fromElements(true, _76_v2);
        let _80_v6;
        _80_v6 = _dafny.Seq.of(_74_v1, new BigNumber((_dafny.Seq.of(_74_v1)).length));
        (globalState).f3 = ((new BigNumber((_78_v4).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_74_v1, _74_v1, new BigNumber((_79_v5).cardinality()), new BigNumber((_79_v5).cardinality()))).cardinality()))).minus((_80_v6)[_module.__default.safeIndex(_74_v1, new BigNumber((_80_v6).length))]);
        _76_v2 = !(_76_v2);
      }
      (globalState).f3 = _74_v1;
      (globalState).f3 = new BigNumber((_73_v0).length);
      let _81_v7;
      _81_v7 = false;
      let _82_v9;
      _82_v9 = _dafny.Map.Empty.slice().updateUnsafe(_81_v7,function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of _dafny.IntegerRange(new BigNumber(-904), new BigNumber(973))) {
          let _83_v8 = _compr_53;
          if (((new BigNumber(-904)).isLessThanOrEqualTo(_83_v8)) && ((_83_v8).isLessThan(new BigNumber(973)))) {
            _coll53.add((_83_v8).multipliedBy(_74_v1));
          }
        }
        return _coll53;
      }());
      let _84_v10;
      _84_v10 = _dafny.Seq.of(_74_v1, _74_v1, _74_v1, new BigNumber(657));
      (globalState).f3 = (_dafny.ZERO).minus(((true) ? ((new BigNumber((_82_v9).length)).multipliedBy((_84_v10)[_module.__default.safeIndex(_74_v1, new BigNumber((_84_v10).length))])) : (_74_v1)));
      let _85_v11;
      let _nw1 = Array((new BigNumber(21)).toNumber()).fill(false);
      _85_v11 = _nw1;
      let _86_v12;
      _86_v12 = _dafny.Seq.of(_85_v11);
      _85_v11 = (_86_v12)[_module.__default.safeIndex((_dafny.ZERO).minus(_74_v1), new BigNumber((_86_v12).length))];
      let _87_v13;
      let _nw2 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _87_v13 = _nw2;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_87_v13).length))) {
        let _88_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_88_i1)) && ((_88_i1).isLessThan(new BigNumber((_87_v13).length))))) {
          (_87_v13)[(_88_i1)] = _module.__default.safeModuloInt(_88_i1, _module.__default.safeModuloInt(_74_v1, _74_v1));
        }
      }
      r0 = _74_v1;
      let _89_v14;
      _89_v14 = _dafny.Map.Empty.slice().updateUnsafe(_87_v13,_module.__default.fm54(_74_v1, globalState));
      let _90_v15;
      _90_v15 = _89_v14;
      r1 = (_90_v15);
      r2 = _81_v7;
      let _91_v16;
      let _nw3 = new _module.C0();
      _nw3.__ctor();
      _91_v16 = _nw3;
      r3 = ((_81_v7) ? (_91_v16) : (_91_v16));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _92_v0;
      _92_v0 = new BigNumber(545);
      let _93_v1;
      _93_v1 = _dafny.Seq.of(_92_v0);
      let _94_v2;
      _94_v2 = false;
      let _95_v3;
      _95_v3 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_92_v0);
      let _96_v4;
      _96_v4 = _module.D0.create_DC0(new BigNumber((_95_v3).length));
      let _97_v5;
      _97_v5 = _dafny.Seq.UnicodeFromString("cj");
      let _98_v6;
      _98_v6 = _dafny.Set.fromElements(_97_v5);
      let _99_v7;
      _99_v7 = _dafny.MultiSet.fromElements(_92_v0);
      let _100_globalState;
      let _nw4 = new _module.GlobalState();
      _nw4.__ctor(_dafny.Seq.update(_dafny.Seq.update(_93_v1, _module.__default.safeIndex((_96_v4).dtor_cf0, new BigNumber((_93_v1).length)), _92_v0), _module.__default.safeIndex(_92_v0, new BigNumber((_dafny.Seq.update(_93_v1, _module.__default.safeIndex((_96_v4).dtor_cf0, new BigNumber((_93_v1).length)), _92_v0)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_94_v2,new BigNumber((_98_v6).length))).length)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("w"), _97_v5), _99_v7, new BigNumber(742), false, false, true, new BigNumber(339), new BigNumber(344));
      _100_globalState = _nw4;
      let _101_v8;
      let _init0 = ((_102_v2) => function (_103_i0) {
        return (_103_i0).multipliedBy(new BigNumber((_dafny.Set.fromElements(_102_v2, _102_v2)).length));
      })(_94_v2);
      let _nw5 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
        _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _101_v8 = _nw5;
      let _104_v9;
      _104_v9 = _module.D0.create_DC1(_92_v0, _92_v0);
      let _rhs0 = _101_v8;
      let _rhs1 = (_104_v9).dtor_cf1;
      let _lhs0 = _100_globalState;
      _101_v8 = _rhs0;
      _lhs0.f3 = _rhs1;
      let _105_i1;
      _105_i1 = _dafny.ZERO;
      L0: {
        while (_94_v2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_105_i1)) {
              break L0;
            }
            _105_i1 = (_105_i1).plus(_dafny.ONE);
            let _index2 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length));
            (_101_v8)[_index2] = _92_v0;
            let _index3 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length));
            (_101_v8)[_index3] = _92_v0;
            let _source7 = _96_v4;
            if (_source7.is_DC1) {
              let _106___mcc_h0 = (_source7).cf1;
              let _107___mcc_h1 = (_source7).cf2;
              let _108_cf2 = _107___mcc_h1;
              let _109_cf1 = _106___mcc_h0;
              let _110_v10;
              let _nw6 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
              _110_v10 = _nw6;
              let _init1 = ((_111_cf2) => function (_112_i2) {
                return (_112_i2).minus(_111_cf2);
              })(_108_cf2);
              let _nw7 = Array((new BigNumber(14)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
                _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _110_v10 = _nw7;
              (_100_globalState).f3 = new BigNumber(716);
              (_100_globalState).f2 = _99_v7;
              let _index4 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length));
              let _rhs2 = _94_v2;
              let _rhs3 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_108_cf2, _92_v0));
              let _rhs4 = !(_92_v0).isEqualTo((_109_cf1).multipliedBy(_109_cf1));
              let _lhs1 = _100_globalState;
              let _lhs2 = _101_v8;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length));
              let _lhs4 = _100_globalState;
              _lhs1.f4 = _rhs2;
              _lhs2[_lhs3] = _rhs3;
              _lhs4.f4 = _rhs4;
            } else if (_source7.is_DC2) {
              let _113_v11;
              let _nw8 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
              _113_v11 = _nw8;
              _113_v11 = _113_v11;
              _95_v3 = (_95_v3).update(!(!(_94_v2) || (_94_v2)), (_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))]);
              let _114_v12;
              _114_v12 = _dafny.Map.Empty.slice().updateUnsafe((_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))],_92_v0);
              (_100_globalState).f0 = ((_94_v2) ? (_93_v1) : (_dafny.Seq.of(_92_v0, new BigNumber((_114_v12).length))));
              _95_v3 = _95_v3;
            } else if (_source7.is_DC0) {
              let _115___mcc_h2 = (_source7).cf0;
              let _116_cf0 = _115___mcc_h2;
              (_100_globalState).f8 = (_module.__default.safeDivisionInt(_116_cf0, (_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))])).minus((_module.D0.create_DC0(new BigNumber((_module.__default.fm0(_116_cf0, _100_globalState)).cardinality()))).dtor_cf0);
              (_100_globalState).f4 = _94_v2;
              let _117_v13;
              let _nw9 = Array((_dafny.ONE).toNumber());
              _nw9[(_dafny.ZERO).toNumber()] = _93_v1;
              _117_v13 = _nw9;
              let _118_v14;
              _118_v14 = _dafny.Seq.of(_94_v2, _94_v2, _94_v2, false, false);
              let _119_v15;
              _119_v15 = _dafny.Set.fromElements(_94_v2);
              let _120_v16;
              _120_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_119_v15).length),_116_cf0);
              let _index5 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_117_v13).length));
              (_117_v13)[_index5] = _module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_116_cf0, (_dafny.ZERO).minus(new BigNumber((_118_v14).length)), new BigNumber((_97_v5).length), new BigNumber(454), new BigNumber((_97_v5).length))).length)), new _dafny.CodePoint('f'.codePointAt(0)), !(_94_v2), _120_v16, _100_globalState);
              let _index6 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_117_v13).length));
              (_117_v13)[_index6] = _93_v1;
              _92_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(681), (_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))]));
            } else {
              let _121___mcc_h3 = (_source7).cf3;
              let _122_cf3 = _121___mcc_h3;
              (_100_globalState).f8 = (_dafny.ZERO).minus((_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))]);
              let _123_v17;
              let _nw10 = new _module.C5();
              _nw10.__ctor();
              _123_v17 = _nw10;
              let _124_v18;
              _124_v18 = _dafny.Set.fromElements((_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))], _92_v0);
              let _rhs5 = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ethjt"), (_97_v5)[_module.__default.safeIndex(_92_v0, new BigNumber((_97_v5).length))]);
              let _rhs6 = _97_v5;
              let _rhs7 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus((_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))]), (_101_v8)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_101_v8).length))])).isLessThanOrEqualTo(new BigNumber((_124_v18).length));
              let _rhs8 = _dafny.Seq.Concat(_97_v5, _97_v5);
              let _lhs5 = _100_globalState;
              _94_v2 = _rhs5;
              _97_v5 = _rhs6;
              _lhs5.f4 = _rhs7;
              _97_v5 = _rhs8;
              (_100_globalState).f4 = _94_v2;
            }
            let _125_v19;
            let _nw11 = new _module.C8();
            _nw11.__ctor();
            _125_v19 = _nw11;
            let _126_v20;
            let _init2 = ((_127_v2) => function (_128_i3) {
              return _127_v2;
            })(_94_v2);
            let _nw12 = Array((new BigNumber(22)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw12.length); _i0_2++) {
              _nw12[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _126_v20 = _nw12;
            let _index7 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_126_v20).length));
            (_126_v20)[_index7] = !(_94_v2);
            let _129_v21;
            _129_v21 = _dafny.Seq.of(_94_v2, _94_v2, false, _94_v2);
            let _130_v22;
            _130_v22 = _dafny.MultiSet.fromElements(_94_v2);
            let _131_v23;
            _131_v23 = _dafny.MultiSet.fromElements((_dafny.MultiSet.FromArray(_129_v21)).update(_94_v2, _module.__default.abs(_92_v0)), _130_v22, _130_v22);
            let _index8 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_126_v20).length));
            (_126_v20)[_index8] = !(_module.__default.fm49(_module.__default.fm49(true, _100_globalState), _100_globalState)) || ((_131_v23).IsProperSubsetOf(_131_v23));
          }
        }
      }
      let _132_v24;
      _132_v24 = _dafny.Set.fromElements(new BigNumber((_97_v5).length), _92_v0);
      if (!((_132_v24).IsSubsetOf(_132_v24))) {
        if (_94_v2) {
          let _133_v25;
          let _nw13 = Array((new BigNumber(2)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _94_v2;
          _nw13[(_dafny.ONE).toNumber()] = !(_94_v2);
          _133_v25 = _nw13;
          let _134_v26;
          _134_v26 = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,_133_v25);
          _134_v26 = _134_v26;
          let _135_v27;
          let _nw14 = new _module.C5();
          _nw14.__ctor();
          _135_v27 = _nw14;
          let _136_v28;
          let _nw15 = Array((new BigNumber(5)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _135_v27;
          _nw15[(_dafny.ONE).toNumber()] = _135_v27;
          _nw15[(new BigNumber(2)).toNumber()] = _135_v27;
          _nw15[(new BigNumber(3)).toNumber()] = _135_v27;
          _nw15[(new BigNumber(4)).toNumber()] = _135_v27;
          _136_v28 = _nw15;
          let _index9 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_136_v28).length));
          (_136_v28)[_index9] = _135_v27;
          let _index10 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_136_v28).length));
          (_136_v28)[_index10] = _135_v27;
          let _137_v29;
          let _init3 = function (_138_i4) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          };
          let _nw16 = Array((new BigNumber(26)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw16.length); _i0_3++) {
            _nw16[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _137_v29 = _nw16;
          let _139_v30;
          _139_v30 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index11 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_137_v29).length));
          (_137_v29)[_index11] = _139_v30;
          let _index12 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_137_v29).length));
          (_137_v29)[_index12] = new _dafny.CodePoint('h'.codePointAt(0));
          let _140_v31;
          let _nw17 = new _module.C5();
          _nw17.__ctor();
          _140_v31 = _nw17;
          let _141_v32;
          let _nw18 = new _module.C9();
          _nw18.__ctor(_94_v2);
          _141_v32 = _nw18;
        } else {
          let _142_v33;
          _142_v33 = _module.D23.create_DC46();
          let _143_v34;
          _143_v34 = _dafny.MultiSet.fromElements(_142_v33);
          let _144_v35;
          _144_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,_143_v34);
          let _145_v36;
          _145_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_144_v35).length),(_dafny.ZERO).minus(_92_v0));
          _145_v36 = _module.__default.fm52(_100_globalState);
          (_100_globalState).f3 = _92_v0;
          let _146_v37;
          let _init4 = ((_147_v2) => function (_148_i5) {
            return _147_v2;
          })(_94_v2);
          let _nw19 = Array((new BigNumber(6)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw19.length); _i0_4++) {
            _nw19[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _146_v37 = _nw19;
          let _index13 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_146_v37).length));
          (_146_v37)[_index13] = !(_94_v2);
          let _index14 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_146_v37).length));
          (_146_v37)[_index14] = (_92_v0).isLessThanOrEqualTo(_92_v0);
          let _149_v38;
          _149_v38 = _dafny.Map.Empty.slice().updateUnsafe(_101_v8,_146_v37);
          _149_v38 = _149_v38;
          let _index15 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_146_v37).length));
          (_146_v37)[_index15] = (new BigNumber((_97_v5).length)).isLessThan(_92_v0);
        }
        let _150_v39;
        let _151_v40;
        let _152_v41;
        let _153_v42;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m15(_100_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _150_v39 = _out0;
        _151_v40 = _out1;
        _152_v41 = _out2;
        _153_v42 = _out3;
        (_100_globalState).f4 = _152_v41;
        let _154_v43;
        _154_v43 = _module.D10.create_DC24(new BigNumber((_module.__default.fm23(_97_v5, _100_globalState)).length), _152_v41, _97_v5);
        let _155_v44;
        _155_v44 = _dafny.MultiSet.fromElements(_152_v41, (_154_v43).dtor_cf32, _dafny.Seq.IsProperPrefixOf(_97_v5, _97_v5));
        _155_v44 = ((_155_v44).update(!(_94_v2), _module.__default.abs(new BigNumber((_module.__default.fm35(_152_v41, _92_v0, _94_v2, _152_v41, _100_globalState)).length)))).Difference(_155_v44);
        let _156_v45;
        let _157_v46;
        let _158_v47;
        let _159_v48;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = _module.__default.m15(_100_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _156_v45 = _out4;
        _157_v46 = _out5;
        _158_v47 = _out6;
        _159_v48 = _out7;
      } else {
        if (_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("q"), _module.__default.fm18(_94_v2, _94_v2, _100_globalState))) {
          (_100_globalState).f8 = (_93_v1)[_module.__default.safeIndex(_92_v0, new BigNumber((_93_v1).length))];
          (_100_globalState).f8 = (_92_v0).plus(_92_v0);
          let _index16 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_101_v8).length));
          (_101_v8)[_index16] = _92_v0;
          let _index17 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_101_v8).length));
          (_101_v8)[_index17] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_92_v0), new BigNumber((_dafny.Seq.of(new BigNumber((_95_v3).length))).length));
          (_100_globalState).f0 = _93_v1;
          let _index18 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_101_v8).length));
          (_101_v8)[_index18] = _92_v0;
        } else {
          let _160_v49;
          let _nw20 = new _module.C2();
          _nw20.__ctor();
          _160_v49 = _nw20;
          let _161_v50;
          let _nw21 = new _module.C7();
          _nw21.__ctor();
          _161_v50 = _nw21;
          let _162_v51;
          _162_v51 = _dafny.Map.Empty.slice().updateUnsafe(_161_v50,(_161_v50).fm3(_93_v1, _99_v7, _94_v2, _100_globalState));
          let _163_v52;
          _163_v52 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_162_v51);
          (_100_globalState).f4 = (!(_162_v51).equals((((_163_v52).contains(false)) ? ((_163_v52).get(false)) : (_162_v51)))) === (_94_v2);
          let _164_v53;
          let _nw22 = new _module.C9();
          _nw22.__ctor(_94_v2);
          _164_v53 = _nw22;
          let _165_v54;
          _165_v54 = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,_164_v53);
          let _166_v55;
          let _nw23 = new _module.C5();
          _nw23.__ctor();
          _166_v55 = _nw23;
          let _167_v56;
          _167_v56 = _dafny.Map.Empty.slice().updateUnsafe(_166_v55,_92_v0);
          let _168_v57;
          _168_v57 = _module.D26.create_DC55(_166_v55);
          let _169_v58;
          _169_v58 = _dafny.Seq.of(_94_v2, _94_v2);
          let _170_v59;
          let _nw24 = Array((new BigNumber(5)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _165_v54;
          _nw24[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_167_v56).contains((_168_v57).dtor_cf77)) ? ((_167_v56).get((_168_v57).dtor_cf77)) : (new BigNumber((_169_v58).length))),_164_v53);
          _nw24[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,_164_v53);
          _nw24[(new BigNumber(3)).toNumber()] = _165_v54;
          _nw24[(new BigNumber(4)).toNumber()] = _165_v54;
          _170_v59 = _nw24;
          _170_v59 = _170_v59;
          _97_v5 = _97_v5;
          (_100_globalState).f8 = _92_v0;
        }
        let _171_v60;
        let _nw25 = new _module.C8();
        _nw25.__ctor();
        _171_v60 = _nw25;
        let _172_v61;
        _172_v61 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_97_v5);
        (_171_v60).m4(((_dafny.ZERO).minus(_92_v0)).isEqualTo(_92_v0), (_dafny.Map.Empty.slice().updateUnsafe(_94_v2,_97_v5)).Merge(_172_v61), new BigNumber(709), _100_globalState);
        if (_94_v2) {
          _101_v8 = ((!(_94_v2)) ? (_101_v8) : (_101_v8));
          (_100_globalState).f4 = !(_94_v2);
          (_171_v60).m4(_94_v2, _172_v61, _92_v0, _100_globalState);
          let _173_v62;
          _173_v62 = _dafny.Seq.of(_94_v2);
          _173_v62 = _dafny.Seq.Concat(_173_v62, _dafny.Seq.of(!(_94_v2), false, _94_v2));
          _101_v8 = _101_v8;
        } else {
          let _174_v63;
          _174_v63 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,false);
          let _175_v64;
          let _nw26 = Array((new BigNumber(12)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _95_v3;
          _nw26[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_94_v2,new BigNumber(-686))).update(_94_v2, _92_v0);
          _nw26[(new BigNumber(2)).toNumber()] = (_module.__default.fm46(_100_globalState)).update((((_174_v63).contains(!(_94_v2))) ? ((_174_v63).get(!(_94_v2))) : (_94_v2)), new BigNumber(379));
          _nw26[(new BigNumber(3)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(4)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(5)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(6)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_92_v0);
          _nw26[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_92_v0);
          _nw26[(new BigNumber(9)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(10)).toNumber()] = _95_v3;
          _nw26[(new BigNumber(11)).toNumber()] = _95_v3;
          _175_v64 = _nw26;
          let _176_v65;
          let _nw27 = new _module.C4();
          _nw27.__ctor(_175_v64);
          _176_v65 = _nw27;
          _176_v65 = _176_v65;
          let _177_v66;
          let _nw28 = Array((new BigNumber(6)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = ((true) ? (_93_v1) : (_93_v1));
          _nw28[(_dafny.ONE).toNumber()] = _93_v1;
          _nw28[(new BigNumber(2)).toNumber()] = _93_v1;
          _nw28[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_92_v0, _92_v0, _92_v0);
          _nw28[(new BigNumber(4)).toNumber()] = _93_v1;
          _nw28[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), ((_178_v0) => function (_179_i6) {
            return _178_v0;
          })(_92_v0));
          _177_v66 = _nw28;
          let _index19 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length));
          (_177_v66)[_index19] = _93_v1;
          let _index20 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length));
          (_177_v66)[_index20] = _dafny.Seq.Concat(_93_v1, _93_v1);
          let _180_v67;
          let _init5 = ((_181_v3) => function (_182_i7) {
            return !(_181_v3).equals(_181_v3);
          })(_95_v3);
          let _nw29 = Array((new BigNumber(21)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw29.length); _i0_5++) {
            _nw29[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _180_v67 = _nw29;
          let _index21 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_180_v67).length));
          (_180_v67)[_index21] = (new BigNumber((_97_v5).length)).isLessThanOrEqualTo(_92_v0);
          let _index22 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_180_v67).length));
          (_180_v67)[_index22] = !(_94_v2);
          (_100_globalState).f4 = !(!((_92_v0).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((((_174_v63).contains(_94_v2)) ? ((_174_v63).get(_94_v2)) : ((_180_v67)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_180_v67).length))])))).length))) || ((_180_v67)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_180_v67).length))]));
          let _183_v68;
          _183_v68 = _dafny.MultiSet.fromElements(true);
          let _184_v69;
          _184_v69 = _dafny.Seq.of(_183_v68);
          let _185_v70;
          _185_v70 = _dafny.Seq.of(_dafny.Seq.update((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))], _module.__default.safeIndex(_92_v0, new BigNumber(((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))]).length)), _92_v0));
          let _rhs9 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_93_v1, _dafny.Seq.of(((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))])[_module.__default.safeIndex(_92_v0, new BigNumber(((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))]).length))])), (_185_v70)[_module.__default.safeIndex(_92_v0, new BigNumber((_185_v70).length))]), _module.__default.safeIndex(_92_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_93_v1, _dafny.Seq.of(((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))])[_module.__default.safeIndex(_92_v0, new BigNumber(((_177_v66)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_177_v66).length))]).length))])), (_185_v70)[_module.__default.safeIndex(_92_v0, new BigNumber((_185_v70).length))])).length)), _92_v0);
          let _rhs10 = _184_v69;
          let _rhs11 = (_92_v0).plus(_92_v0);
          let _lhs6 = _100_globalState;
          _93_v1 = _rhs9;
          _184_v69 = _rhs10;
          _lhs6.f3 = _rhs11;
        }
        if (((_94_v2) ? (_94_v2) : (_94_v2))) {
          let _186_v71;
          _186_v71 = new _dafny.CodePoint('r'.codePointAt(0));
          let _187_v72;
          _187_v72 = _module.D10.create_DC24((((_99_v7).contains((_dafny.ZERO).minus(_92_v0))) ? ((_99_v7).get((_dafny.ZERO).minus(_92_v0))) : (_module.__default.fm6(new BigNumber(68), _100_globalState))), _94_v2, _dafny.Seq.update(_dafny.Seq.Concat(_97_v5, _dafny.Seq.of(_186_v71)), _module.__default.safeIndex(new BigNumber((_module.__default.fm7(_186_v71, _100_globalState)).length), new BigNumber((_dafny.Seq.Concat(_97_v5, _dafny.Seq.of(_186_v71))).length)), _186_v71));
          let _pat_let_tv0 = _94_v2;
          let _pat_let_tv1 = _92_v0;
          let _pat_let_tv2 = _92_v0;
          let _pat_let_tv3 = _94_v2;
          let _pat_let_tv4 = _92_v0;
          _187_v72 = function (_pat_let0_0) {
            return function (_193_dt__update__tmp_h2) {
              return function (_pat_let6_0) {
                return function (_194_dt__update_hcf31_h1) {
                  return _module.D10.create_DC24(_194_dt__update_hcf31_h1, (_193_dt__update__tmp_h2).dtor_cf32, (_193_dt__update__tmp_h2).dtor_cf33);
                }(_pat_let6_0);
              }(_pat_let_tv4);
            }(_pat_let0_0);
          }(function (_pat_let1_0) {
            return function (_191_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_192_dt__update_hcf32_h1) {
                  return _module.D10.create_DC24((_191_dt__update__tmp_h1).dtor_cf31, _192_dt__update_hcf32_h1, (_191_dt__update__tmp_h1).dtor_cf33);
                }(_pat_let5_0);
              }(_pat_let_tv3);
            }(_pat_let1_0);
          }(function (_pat_let2_0) {
            return function (_188_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_189_dt__update_hcf32_h0) {
                  return function (_pat_let4_0) {
                    return function (_190_dt__update_hcf31_h0) {
                      return _module.D10.create_DC24(_190_dt__update_hcf31_h0, _189_dt__update_hcf32_h0, (_188_dt__update__tmp_h0).dtor_cf33);
                    }(_pat_let4_0);
                  }(new BigNumber((_dafny.Seq.of(_pat_let_tv1, _pat_let_tv2)).length));
                }(_pat_let3_0);
              }(_pat_let_tv0);
            }(_pat_let2_0);
          }(_187_v72)));
          let _index23 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length));
          (_101_v8)[_index23] = _92_v0;
          let _index24 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length));
          (_101_v8)[_index24] = _module.__default.safeDivisionInt(_92_v0, _92_v0);
          let _195_v73;
          let _nw30 = Array((new BigNumber(5)).toNumber()).fill(_module.D2.Default());
          _195_v73 = _nw30;
          let _196_v74;
          let _nw31 = Array((new BigNumber(25)).toNumber()).fill(false);
          _196_v74 = _nw31;
          let _197_v75;
          _197_v75 = _dafny.Seq.of(_196_v74);
          let _198_v76;
          _198_v76 = _module.D2.create_DC5((_197_v75)[_module.__default.safeIndex((_101_v8)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length))], new BigNumber((_197_v75).length))]);
          let _index25 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_195_v73).length));
          (_195_v73)[_index25] = _198_v76;
          let _199_v77;
          let _nw32 = new _module.C1();
          _nw32.__ctor();
          _199_v77 = _nw32;
          let _200_v78;
          let _nw33 = Array((new BigNumber(4)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = ((_94_v2) ? (_199_v77) : (_199_v77));
          _nw33[(_dafny.ONE).toNumber()] = _199_v77;
          _nw33[(new BigNumber(2)).toNumber()] = _199_v77;
          _nw33[(new BigNumber(3)).toNumber()] = _199_v77;
          _200_v78 = _nw33;
          let _index26 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_200_v78).length));
          (_200_v78)[_index26] = _199_v77;
          let _201_v79;
          _201_v79 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), ((_202_v71) => function (_203_i8) {
            return _202_v71;
          })(_186_v71)), _97_v5);
          let _204_v80;
          let _nw34 = new _module.C7();
          _nw34.__ctor();
          _204_v80 = _nw34;
          let _205_v81;
          _205_v81 = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,_204_v80);
          let _index27 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_195_v73).length));
          let _index28 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_200_v78).length));
          let _rhs12 = _198_v76;
          let _rhs13 = _199_v77;
          let _rhs14 = _94_v2;
          let _rhs15 = _94_v2;
          let _rhs16 = ((((_201_v79).contains(_97_v5)) ? ((_201_v79).get(_97_v5)) : (_92_v0))).minus(_module.__default.safeModuloInt(new BigNumber((_205_v81).length), _92_v0));
          let _lhs7 = _195_v73;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_195_v73).length));
          let _lhs9 = _200_v78;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_200_v78).length));
          let _lhs11 = _100_globalState;
          let _lhs12 = _100_globalState;
          _lhs7[_lhs8] = _rhs12;
          _lhs9[_lhs10] = _rhs13;
          _94_v2 = _rhs14;
          _lhs11.f4 = _rhs15;
          _lhs12.f8 = _rhs16;
          let _206_v82;
          _206_v82 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_171_v60);
          let _207_v83;
          _207_v83 = _module.D0.create_DC2();
          let _208_v84;
          _208_v84 = _dafny.Map.Empty.slice().updateUnsafe((_206_v82).Merge(_206_v82),_207_v83);
          let _index29 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length));
          let _rhs17 = _94_v2;
          let _rhs18 = _module.__default.fm49(_94_v2, _100_globalState);
          let _rhs19 = _module.__default.safeDivisionInt((_101_v8)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length))], _module.__default.safeDivisionInt(new BigNumber(104), _92_v0));
          let _rhs20 = !(_94_v2);
          let _rhs21 = new BigNumber((_208_v84).length);
          let _lhs13 = _100_globalState;
          let _lhs14 = _101_v8;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_101_v8).length));
          let _lhs16 = _100_globalState;
          let _lhs17 = _100_globalState;
          _lhs13.f4 = _rhs17;
          _94_v2 = _rhs18;
          _lhs14[_lhs15] = _rhs19;
          _lhs16.f4 = _rhs20;
          _lhs17.f3 = _rhs21;
          let _209_v85;
          _209_v85 = _dafny.MultiSet.fromElements(_module.__default.fm49(_94_v2, _100_globalState));
          let _210_v86;
          _210_v86 = _dafny.Seq.of(_97_v5);
          let _211_v87;
          _211_v87 = _dafny.Map.Empty.slice().updateUnsafe((((_209_v85).contains(_94_v2)) ? ((_209_v85).get(_94_v2)) : (_92_v0)),(new BigNumber(-212)).multipliedBy(new BigNumber(((_210_v86)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("jjqbikwfu")).length), new BigNumber((_210_v86).length))]).length)));
          let _rhs22 = _94_v2;
          let _rhs23 = (((_94_v2) ? (_211_v87) : (_211_v87))).update(_92_v0, _92_v0);
          let _lhs18 = _100_globalState;
          _lhs18.f4 = _rhs22;
          _211_v87 = _rhs23;
        } else {
          let _212_v88;
          _212_v88 = _101_v8;
          let _213_v89;
          _213_v89 = _dafny.Seq.of(_101_v8);
          let _214_v90;
          _214_v90 = _dafny.MultiSet.fromElements(_94_v2);
          let _215_v91;
          let _nw35 = Array((new BigNumber(8)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = ((_94_v2) ? (_101_v8) : (_101_v8));
          _nw35[(_dafny.ONE).toNumber()] = (_212_v88);
          _nw35[(new BigNumber(2)).toNumber()] = _101_v8;
          _nw35[(new BigNumber(3)).toNumber()] = _101_v8;
          _nw35[(new BigNumber(4)).toNumber()] = _101_v8;
          _nw35[(new BigNumber(5)).toNumber()] = _101_v8;
          _nw35[(new BigNumber(6)).toNumber()] = _101_v8;
          _nw35[(new BigNumber(7)).toNumber()] = (_213_v89)[_module.__default.safeIndex(new BigNumber((_214_v90).cardinality()), new BigNumber((_213_v89).length))];
          _215_v91 = _nw35;
          let _index30 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_215_v91).length));
          (_215_v91)[_index30] = _101_v8;
          let _index31 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_215_v91).length));
          let _nw36 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          (_215_v91)[_index31] = _nw36;
          let _216_v93;
          _216_v93 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_92_v0, _92_v0, (_dafny.ZERO).minus(new BigNumber((function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_132_v24).Elements) {
              let _217_v92 = _compr_54;
              if ((_132_v24).contains(_217_v92)) {
                _coll54.add((_217_v92).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(-834)))).length)));
              }
            }
            return _coll54;
          }()).length)), new BigNumber(605))).Union(_132_v24),_94_v2);
          _216_v93 = (_216_v93).Merge(_216_v93);
          let _218_v94;
          _218_v94 = _dafny.Seq.of(_94_v2, _94_v2);
          (_100_globalState).f8 = (_92_v0).minus((((_214_v90).contains((_218_v94)[_module.__default.safeIndex(_92_v0, new BigNumber((_218_v94).length))])) ? ((_214_v90).get((_218_v94)[_module.__default.safeIndex(_92_v0, new BigNumber((_218_v94).length))])) : (_92_v0)));
          let _219_v95;
          let _nw37 = new _module.C9();
          _nw37.__ctor((false) && (_94_v2));
          _219_v95 = _nw37;
          let _220_v96;
          let _nw38 = new _module.C8();
          _nw38.__ctor();
          _220_v96 = _nw38;
          _220_v96 = _220_v96;
        }
      }
      let _221_v97;
      let _nw39 = new _module.C0();
      _nw39.__ctor();
      _221_v97 = _nw39;
      let _222_v98;
      _222_v98 = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,_92_v0);
      let _223_v99;
      _223_v99 = _module.D3.create_DC8(_222_v98);
      let _224_v100;
      _224_v100 = _dafny.Seq.of(_94_v2);
      let _225_v101;
      _225_v101 = _dafny.Set.fromElements(_94_v2);
      let _226_i9;
      _226_i9 = _dafny.ZERO;
      L1: {
        while (_module.__default.fm49((_225_v101).IsSubsetOf(_module.__default.fm35(_94_v2, _module.__default.fm26(_223_v99, _92_v0, _dafny.MultiSet.fromElements(_224_v100, _224_v100, _224_v100), _100_globalState), _94_v2, true, _100_globalState)), _100_globalState)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_226_i9)) {
              break L1;
            }
            _226_i9 = (_226_i9).plus(_dafny.ONE);
            let _227_v102;
            let _nw40 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
            _227_v102 = _nw40;
            let _228_v103;
            let _nw41 = new _module.C4();
            _nw41.__ctor(_227_v102);
            _228_v103 = _nw41;
            let _229_v104;
            _229_v104 = _dafny.Seq.of(_228_v103);
            _228_v103 = (_229_v104)[_module.__default.safeIndex(new BigNumber(208), new BigNumber((_229_v104).length))];
            (_100_globalState).f8 = ((!(_94_v2)) ? (_module.__default.safeDivisionInt(new BigNumber(-265), _92_v0)) : ((_92_v0).plus(_92_v0)));
            let _230_v105;
            let _nw42 = new _module.C7();
            _nw42.__ctor();
            _230_v105 = _nw42;
            let _231_v106;
            _231_v106 = new _dafny.CodePoint('w'.codePointAt(0));
            let _232_v107;
            _232_v107 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_230_v105,_231_v106));
            let _233_v108;
            _233_v108 = _dafny.Map.Empty.slice().updateUnsafe(_230_v105,new _dafny.CodePoint('q'.codePointAt(0)));
            _232_v107 = _dafny.Set.fromElements(_233_v108, _233_v108, _dafny.Map.Empty.slice().updateUnsafe(_230_v105,new _dafny.CodePoint('y'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(_230_v105,_231_v106));
            (_228_v103).m9(_module.__default.fm6(_92_v0, _100_globalState), _92_v0, _100_globalState);
          }
        }
      }
      let _234_v109;
      _234_v109 = _module.D3.create_DC9();
      _234_v109 = _234_v109;
      let _235_v110;
      _235_v110 = _module.D21.create_DC43(_dafny.Seq.Create(_module.__default.abs(new BigNumber(678)), ((_236_v2) => function (_237_i11) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_236_v2)).length);
})(_94_v2)), false);
      let _238_i10;
      _238_i10 = _dafny.ZERO;
      L2: {
        let _pat_let_tv5 = _92_v0;
        let _pat_let_tv6 = _92_v0;
        let _pat_let_tv7 = _94_v2;
        let _pat_let_tv8 = _94_v2;
        let _pat_let_tv9 = _100_globalState;
        while (function (_source9) {
          if (_source9.is_DC43) {
            let _288___mcc_h11 = (_source9).cf63;
            let _289___mcc_h12 = (_source9).cf64;
            let _290_cf64 = _289___mcc_h12;
            let _291_cf63 = _288___mcc_h11;
            return (_pat_let_tv5).isLessThanOrEqualTo(_pat_let_tv6);
          } else {
            let _292___mcc_h13 = (_source9).cf62;
            let _293_cf62 = _292___mcc_h13;
            return _pat_let_tv7;
          }
        }(function (_pat_let7_0) {
          return function (_294_dt__update__tmp_h3) {
            return function (_pat_let8_0) {
              return function (_295_dt__update_hcf64_h0) {
                return _module.D21.create_DC43((_294_dt__update__tmp_h3).dtor_cf63, _295_dt__update_hcf64_h0);
              }(_pat_let8_0);
            }(_module.__default.fm49(_pat_let_tv8, _pat_let_tv9));
          }(_pat_let7_0);
        }(_235_v110))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_238_i10)) {
              break L2;
            }
            _238_i10 = (_238_i10).plus(_dafny.ONE);
            if (_94_v2) {
              let _239_v111;
              let _nw43 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
              _239_v111 = _nw43;
              let _240_v112;
              let _nw44 = new _module.C4();
              _nw44.__ctor(_239_v111);
              _240_v112 = _nw44;
              (_100_globalState).f8 = ((_94_v2) ? (_92_v0) : ((_dafny.ZERO).minus(_92_v0)));
              let _241_v113;
              _241_v113 = _dafny.Map.Empty.slice().updateUnsafe(_92_v0,(_240_v112).fm3(_93_v1, _dafny.MultiSet.fromElements(_92_v0), !(_94_v2), _100_globalState));
              let _242_v114;
              _242_v114 = _dafny.Set.fromElements(_241_v113, _241_v113, _241_v113, (_241_v113).Merge(_241_v113), _241_v113);
              let _243_v115;
              let _nw45 = new _module.C2();
              _nw45.__ctor();
              _243_v115 = _nw45;
              let _244_v116;
              _244_v116 = _dafny.Seq.of(_243_v115, _243_v115);
              let _245_v117;
              _245_v117 = new _dafny.CodePoint('y'.codePointAt(0));
              let _246_v118;
              let _nw46 = new _module.C1();
              _nw46.__ctor();
              _246_v118 = _nw46;
              let _247_v119;
              _247_v119 = _dafny.Map.Empty.slice().updateUnsafe(_246_v118,new BigNumber(110));
              let _rhs24 = _dafny.Seq.IsPrefixOf(_97_v5, _module.__default.fm25(_100_globalState));
              let _rhs25 = (_dafny.Set.fromElements(_241_v113)).Intersect(_module.__default.fm53(_92_v0, _92_v0, _100_globalState));
              let _rhs26 = (_244_v116)[_module.__default.safeIndex(_92_v0, new BigNumber((_244_v116).length))];
              let _rhs27 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ukahennov"), _module.__default.safeIndex(_92_v0, new BigNumber((_dafny.Seq.UnicodeFromString("ukahennov")).length)), _245_v117), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uvcx"), _97_v5)), _module.__default.safeIndex((((_247_v119).contains(_246_v118)) ? ((_247_v119).get(_246_v118)) : ((_dafny.ZERO).minus(_92_v0))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ukahennov"), _module.__default.safeIndex(_92_v0, new BigNumber((_dafny.Seq.UnicodeFromString("ukahennov")).length)), _245_v117), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uvcx"), _97_v5))).length)), _245_v117);
              _94_v2 = _rhs24;
              _242_v114 = _rhs25;
              _243_v115 = _rhs26;
              _97_v5 = _rhs27;
              let _248_v120;
              _248_v120 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gofjuqlqk"),_94_v2);
              (_100_globalState).f3 = new BigNumber((_248_v120).length);
              _92_v0 = new BigNumber(-256);
            } else {
              let _249_v121;
              let _nw47 = Array((new BigNumber(10)).toNumber()).fill(false);
              _249_v121 = _nw47;
              let _index32 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_249_v121).length));
              (_249_v121)[_index32] = _94_v2;
              let _index33 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_249_v121).length));
              (_249_v121)[_index33] = _94_v2;
              let _250_v122;
              let _251_v123;
              let _252_v124;
              let _253_v125;
              let _out8;
              let _out9;
              let _out10;
              let _out11;
              let _outcollector2 = _module.__default.m15(_100_globalState);
              _out8 = _outcollector2[0];
              _out9 = _outcollector2[1];
              _out10 = _outcollector2[2];
              _out11 = _outcollector2[3];
              _250_v122 = _out8;
              _251_v123 = _out9;
              _252_v124 = _out10;
              _253_v125 = _out11;
              let _254_v126;
              _254_v126 = _dafny.Map.Empty.slice().updateUnsafe(_252_v124,(_249_v121)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_249_v121).length))]);
              (_100_globalState).f4 = !(_254_v126).equals(_254_v126);
              let _255_v127;
              let _256_v128;
              let _257_v129;
              let _258_v130;
              let _out12;
              let _out13;
              let _out14;
              let _out15;
              let _outcollector3 = _module.__default.m15(_100_globalState);
              _out12 = _outcollector3[0];
              _out13 = _outcollector3[1];
              _out14 = _outcollector3[2];
              _out15 = _outcollector3[3];
              _255_v127 = _out12;
              _256_v128 = _out13;
              _257_v129 = _out14;
              _258_v130 = _out15;
              let _index34 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_101_v8).length));
              (_101_v8)[_index34] = _250_v122;
              let _index35 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_101_v8).length));
              (_101_v8)[_index35] = _255_v127;
            }
            let _rhs28 = new BigNumber(-265);
            let _rhs29 = (_dafny.ZERO).minus(_92_v0);
            let _lhs19 = _100_globalState;
            let _lhs20 = _100_globalState;
            _lhs19.f8 = _rhs28;
            _lhs20.f8 = _rhs29;
            let _259_v131;
            let _nw48 = new _module.C7();
            _nw48.__ctor();
            _259_v131 = _nw48;
            let _260_v132;
            _260_v132 = _dafny.Seq.of(_259_v131);
            let _261_v133;
            _261_v133 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_94_v2);
            let _262_v134;
            _262_v134 = _dafny.MultiSet.fromElements(_94_v2);
            let _263_v135;
            let _nw49 = Array((new BigNumber(29)).toNumber());
            _nw49[(_dafny.ZERO).toNumber()] = (_132_v24).IsDisjointFrom(_132_v24);
            _nw49[(_dafny.ONE).toNumber()] = _dafny.areEqual(_224_v100, _dafny.Seq.of(_94_v2, _94_v2, _94_v2, _94_v2, _94_v2));
            _nw49[(new BigNumber(2)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(3)).toNumber()] = (((_261_v133).contains(_94_v2)) ? ((_261_v133).get(_94_v2)) : (_94_v2));
            _nw49[(new BigNumber(4)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(5)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(6)).toNumber()] = _dafny.areEqual(_224_v100, _224_v100);
            _nw49[(new BigNumber(7)).toNumber()] = false;
            _nw49[(new BigNumber(8)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(9)).toNumber()] = true;
            _nw49[(new BigNumber(10)).toNumber()] = ((_94_v2) ? (_94_v2) : (_94_v2));
            _nw49[(new BigNumber(11)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(12)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(13)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(14)).toNumber()] = (_259_v131).fm3(_dafny.Seq.of(new BigNumber((_97_v5).length)), _99_v7, _94_v2, _100_globalState);
            _nw49[(new BigNumber(15)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(16)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(17)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(18)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(19)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(20)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(21)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(22)).toNumber()] = !(_94_v2) || (_94_v2);
            _nw49[(new BigNumber(23)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(24)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(25)).toNumber()] = false;
            _nw49[(new BigNumber(26)).toNumber()] = (_dafny.MultiSet.fromElements(_94_v2)).IsDisjointFrom(_262_v134);
            _nw49[(new BigNumber(27)).toNumber()] = _94_v2;
            _nw49[(new BigNumber(28)).toNumber()] = _94_v2;
            _263_v135 = _nw49;
            let _index36 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length));
            (_263_v135)[_index36] = !(!(_94_v2) || (_94_v2));
            let _264_v137;
            let _init6 = ((_265_v2) => function (_266_i12) {
              return function () {
                let _coll55 = new _dafny.Map();
                for (const _compr_55 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)))).Elements) {
                  let _267_v136 = _compr_55;
                  if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)))).contains(_267_v136)) {
                    _coll55.push([_267_v136,_265_v2]);
                  }
                }
                return _coll55;
              }();
            })(_94_v2);
            let _nw50 = Array((new BigNumber(4)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw50.length); _i0_6++) {
              _nw50[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _264_v137 = _nw50;
            let _268_v138;
            _268_v138 = new _dafny.CodePoint('q'.codePointAt(0));
            let _269_v139;
            _269_v139 = _dafny.Map.Empty.slice().updateUnsafe(_268_v138,_94_v2);
            let _index37 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_264_v137).length));
            (_264_v137)[_index37] = (_269_v139).update(_268_v138, _94_v2);
            let _270_v140;
            _270_v140 = _module.D27.create_DC58(_269_v139);
            let _index38 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length));
            let _index39 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_264_v137).length));
            let _rhs30 = _92_v0;
            let _rhs31 = _dafny.Seq.of(((_94_v2) ? (_259_v131) : (_259_v131)), _259_v131, _259_v131, _259_v131, _259_v131);
            let _rhs32 = false;
            let _rhs33 = (_259_v131).fm3(_93_v1, _99_v7, _94_v2, _100_globalState);
            let _rhs34 = (_270_v140).dtor_cf82;
            let _lhs21 = _100_globalState;
            let _lhs22 = _263_v135;
            let _lhs23 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length));
            let _lhs24 = _264_v137;
            let _lhs25 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_264_v137).length));
            _lhs21.f8 = _rhs30;
            _260_v132 = _rhs31;
            _94_v2 = _rhs32;
            _lhs22[_lhs23] = _rhs33;
            _lhs24[_lhs25] = _rhs34;
            let _271_v141;
            _271_v141 = _module.D2.create_DC5(_263_v135);
            let _source8 = _271_v141;
            if (_source8.is_DC6) {
              let _272___mcc_h4 = (_source8).cf6;
              let _273___mcc_h5 = (_source8).cf7;
              let _274___mcc_h6 = (_source8).cf8;
              let _275_cf8 = _274___mcc_h6;
              let _276_cf7 = _273___mcc_h5;
              let _277_cf6 = _272___mcc_h4;
              _95_v3 = (_95_v3).update(_276_cf7, _277_cf6);
              (_100_globalState).f8 = _module.__default.safeDivisionInt(_92_v0, _92_v0);
              let _278_v142;
              let _nw51 = Array((new BigNumber(25)).toNumber()).fill([]);
              _278_v142 = _nw51;
              _278_v142 = _278_v142;
              let _index40 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_101_v8).length));
              (_101_v8)[_index40] = _92_v0;
              let _index41 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_101_v8).length));
              (_101_v8)[_index41] = _277_cf6;
            } else if (_source8.is_DC7) {
              let _279___mcc_h7 = (_source8).cf9;
              let _280___mcc_h8 = (_source8).cf10;
              let _281___mcc_h9 = (_source8).cf11;
              let _282_cf11 = _281___mcc_h9;
              let _283_cf10 = _280___mcc_h8;
              let _284_cf9 = _279___mcc_h7;
              _283_cf10 = (_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))];
              (_100_globalState).f4 = (_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))];
              _261_v133 = (_261_v133).update(!((_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))]) || (_283_cf10), (new BigNumber((_dafny.Seq.UnicodeFromString("dmy")).length)).isLessThanOrEqualTo(_92_v0));
              (_100_globalState).f8 = (_dafny.ZERO).minus((((_95_v3).contains(_283_cf10)) ? ((_95_v3).get(_283_cf10)) : (_92_v0)));
            } else {
              let _285___mcc_h10 = (_source8).cf5;
              let _286_cf5 = _285___mcc_h10;
              let _index42 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length));
              (_263_v135)[_index42] = (((_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))]) ? ((_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))]) : ((_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))]));
              _101_v8 = _101_v8;
              let _287_v143;
              _287_v143 = _module.D27.create_DC59(_94_v2, _263_v135, (_263_v135)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_263_v135).length))]);
              (_100_globalState).f4 = (_287_v143).dtor_cf83;
              (_100_globalState).f3 = _module.__default.safeDivisionInt(_92_v0, _92_v0);
            }
          }
        }
      }
      let _296_v144;
      let _init7 = ((_297_v2) => function (_298_i13) {
        return _297_v2;
      })(_94_v2);
      let _nw52 = Array((new BigNumber(12)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw52.length); _i0_7++) {
        _nw52[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _296_v144 = _nw52;
      let _index43 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
      (_296_v144)[_index43] = _94_v2;
      let _index44 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
      (_296_v144)[_index44] = _94_v2;
      let _hi1 = ((_94_v2) ? (_92_v0) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(250)), function (_299_i15) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length)));
      for (let _300_i14 = _92_v0; _300_i14.isLessThan(_hi1); _300_i14 = _300_i14.plus(_dafny.ONE)) {
        let _301_v145;
        _301_v145 = _dafny.Seq.of(_296_v144);
        (_100_globalState).f3 = new BigNumber((_dafny.Seq.of((_301_v145)[_module.__default.safeIndex(_92_v0, new BigNumber((_301_v145).length))], _296_v144)).length);
        let _302_v146;
        _302_v146 = new _dafny.CodePoint('y'.codePointAt(0));
        _302_v146 = new _dafny.CodePoint('l'.codePointAt(0));
        _99_v7 = _99_v7;
        if (!_dafny.Seq.contains(_module.__default.fm23(_dafny.Seq.update(_97_v5, _module.__default.safeIndex(_92_v0, new BigNumber((_97_v5).length)), _302_v146), _100_globalState), !(false))) {
          let _303_v147;
          let _nw53 = Array((new BigNumber(23)).toNumber()).fill([]);
          _303_v147 = _nw53;
          let _index45 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_303_v147).length));
          (_303_v147)[_index45] = _101_v8;
          let _index46 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_303_v147).length));
          (_303_v147)[_index46] = _101_v8;
          let _304_v148;
          let _nw54 = new _module.C5();
          _nw54.__ctor();
          _304_v148 = _nw54;
          let _305_v149;
          let _nw55 = Array((_dafny.ONE).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = _304_v148;
          _305_v149 = _nw55;
          let _index47 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_305_v149).length));
          (_305_v149)[_index47] = (((_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))]) ? (_304_v148) : (_304_v148));
          let _index48 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_305_v149).length));
          (_305_v149)[_index48] = _304_v148;
          (_100_globalState).f4 = (_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))];
          let _306_v150;
          _306_v150 = _dafny.MultiSet.fromElements(_94_v2);
          let _index49 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
          let _rhs35 = ((!((_300_i14).isLessThan(_300_i14))) ? (_92_v0) : (_300_i14));
          let _rhs36 = ((_dafny.MultiSet.FromArray(_224_v100)).Difference(_306_v150)).IsProperSubsetOf((_306_v150).update((_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))], _module.__default.abs(_300_i14)));
          let _rhs37 = (_92_v0).minus(_92_v0);
          let _rhs38 = _dafny.Seq.Concat(_dafny.Seq.Concat(_224_v100, _224_v100), _224_v100);
          let _rhs39 = new BigNumber((_dafny.Seq.update(_224_v100, _module.__default.safeIndex((_dafny.ZERO).minus(_92_v0), new BigNumber((_224_v100).length)), _94_v2)).length);
          let _lhs26 = _100_globalState;
          let _lhs27 = _296_v144;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
          let _lhs29 = _100_globalState;
          let _lhs30 = _100_globalState;
          _lhs26.f3 = _rhs35;
          _lhs27[_lhs28] = _rhs36;
          _lhs29.f3 = _rhs37;
          _224_v100 = _rhs38;
          _lhs30.f3 = _rhs39;
          let _index50 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
          (_296_v144)[_index50] = (_306_v150).IsDisjointFrom(_306_v150);
        } else {
          let _307_v151;
          let _nw56 = new _module.C8();
          _nw56.__ctor();
          _307_v151 = _nw56;
          let _308_v152;
          _308_v152 = _dafny.Seq.of(_307_v151, _307_v151);
          _307_v151 = (_308_v152)[_module.__default.safeIndex(_300_i14, new BigNumber((_308_v152).length))];
          (_100_globalState).f3 = _92_v0;
          let _index51 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
          (_296_v144)[_index51] = _module.__default.fm49(_94_v2, _100_globalState);
          let _309_v153;
          let _nw57 = new _module.C3();
          _nw57.__ctor();
          _309_v153 = _nw57;
          _309_v153 = _309_v153;
          (_100_globalState).f8 = _92_v0;
        }
      }
      let _310_v154;
      let _nw58 = new _module.C8();
      _nw58.__ctor();
      _310_v154 = _nw58;
      let _311_v155;
      let _nw59 = Array((new BigNumber(3)).toNumber());
      _nw59[(_dafny.ZERO).toNumber()] = _310_v154;
      _nw59[(_dafny.ONE).toNumber()] = _310_v154;
      _nw59[(new BigNumber(2)).toNumber()] = _310_v154;
      _311_v155 = _nw59;
      let _312_v156;
      let _nw60 = Array((new BigNumber(17)).toNumber());
      _nw60[(_dafny.ZERO).toNumber()] = _311_v155;
      _nw60[(_dafny.ONE).toNumber()] = _311_v155;
      _nw60[(new BigNumber(2)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(3)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(4)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(5)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(6)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(7)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(8)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(9)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(10)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(11)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(12)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(13)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(14)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(15)).toNumber()] = _311_v155;
      _nw60[(new BigNumber(16)).toNumber()] = _311_v155;
      _312_v156 = _nw60;
      let _index52 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_312_v156).length));
      (_312_v156)[_index52] = _311_v155;
      let _313_v157;
      _313_v157 = _dafny.Map.Empty.slice().updateUnsafe((_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))],_224_v100);
      let _314_v158;
      _314_v158 = _dafny.MultiSet.fromElements((((_313_v157).contains(_94_v2)) ? ((_313_v157).get(_94_v2)) : (_224_v100)));
      let _index53 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_312_v156).length));
      let _rhs40 = false;
      let _rhs41 = _dafny.Seq.UnicodeFromString("kmhjp");
      let _rhs42 = _module.__default.fm26(_223_v99, new BigNumber((_224_v100).length), _314_v158, _100_globalState);
      let _rhs43 = _311_v155;
      let _lhs31 = _100_globalState;
      let _lhs32 = _100_globalState;
      let _lhs33 = _312_v156;
      let _lhs34 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_312_v156).length));
      _lhs31.f4 = _rhs40;
      _97_v5 = _rhs41;
      _lhs32.f3 = _rhs42;
      _lhs33[_lhs34] = _rhs43;
      (_100_globalState).f4 = !(_94_v2);
      (_100_globalState).f3 = (_92_v0).multipliedBy(_92_v0);
      (_100_globalState).f4 = true;
      let _315_v159;
      _315_v159 = _module.D2.create_DC6(_92_v0, (_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))], (_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))]);
      let _hi2 = (_92_v0).minus((_315_v159).dtor_cf6);
      for (let _316_i16 = new BigNumber((_224_v100).length); _316_i16.isLessThan(_hi2); _316_i16 = _316_i16.plus(_dafny.ONE)) {
        if (_94_v2) {
          let _317_v160;
          _317_v160 = new _dafny.CodePoint('x'.codePointAt(0));
          let _318_v161;
          _318_v161 = _dafny.Map.Empty.slice().updateUnsafe(_317_v160,_97_v5);
          let _319_v162;
          _319_v162 = _dafny.Map.Empty.slice().updateUnsafe((((_318_v161).contains(_317_v160)) ? ((_318_v161).get(_317_v160)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-780)), ((_320_v0, _321_v5, _322_v2, _323_i16) => function (_324_i17) {
            return (_module.D16.create_DC35(_320_v0, (_321_v5)[_module.__default.safeIndex(_320_v0, new BigNumber((_321_v5).length))], new BigNumber((_dafny.Set.fromElements(_322_v2, _322_v2)).length), _323_i16, _320_v0)).dtor_cf51;
          })(_92_v0, _97_v5, _94_v2, _316_i16)))),_94_v2);
          let _325_v163;
          _325_v163 = _dafny.Seq.of(_dafny.Set.fromElements(_92_v0));
          _319_v162 = _module.__default.fm42(_316_i16, new BigNumber(198), new BigNumber(((_325_v163)[_module.__default.safeIndex(_316_i16, new BigNumber((_325_v163).length))]).length), true, _100_globalState);
          let _326_v164;
          _326_v164 = _dafny.Set.fromElements(_222_v98, (_222_v98).Merge(_222_v98));
          let _327_v165;
          _327_v165 = _dafny.Map.Empty.slice().updateUnsafe((_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))],_dafny.Set.fromElements(_222_v98));
          _326_v164 = (((_327_v165).contains(!(_94_v2) || (true))) ? ((_327_v165).get(!(_94_v2) || (true))) : (_326_v164));
          (_100_globalState).f3 = new BigNumber(359);
          let _index54 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length));
          (_296_v144)[_index54] = (_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))];
          let _328_v166;
          _328_v166 = _dafny.MultiSet.fromElements(_97_v5, _97_v5);
          (_100_globalState).f3 = _module.__default.safeModuloInt(new BigNumber((_328_v166).cardinality()), _316_i16);
        } else {
          (_100_globalState).f4 = (_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))];
          _95_v3 = (_95_v3).update(_94_v2, new BigNumber((_98_v6).length));
          (_100_globalState).f4 = _94_v2;
          let _329_v167;
          _329_v167 = _module.D13.create_DC30((_296_v144)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_296_v144).length))], _94_v2, _132_v24, _221_v97);
          (_100_globalState).f4 = !(!(_94_v2)) || ((_329_v167).dtor_cf43);
          let _330_v168;
          _330_v168 = _dafny.Map.Empty.slice().updateUnsafe(_94_v2,_97_v5);
          (_310_v154).m4(_94_v2, _330_v168, new BigNumber(239), _100_globalState);
        }
        (_100_globalState).f4 = true;
        (_100_globalState).f4 = true;
        _101_v8 = _101_v8;
      }
      let _331_v169;
      let _nw61 = new _module.C0();
      _nw61.__ctor();
      _331_v169 = _nw61;
      let _332_v170;
      _332_v170 = _dafny.MultiSet.fromElements(_97_v5);
      (_100_globalState).f4 = (_332_v170).IsProperSubsetOf((_332_v170).Union(_332_v170));
      process.stdout.write(_dafny.toString(_92_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_93_v1, _dafny.Seq.of(new BigNumber(545), new BigNumber(545), new BigNumber(545), new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_94_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_95_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_v4).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_97_v5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v6).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("cj")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_99_v7).equals(_dafny.MultiSet.fromElements(new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_100_globalState.f0, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_100_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_globalState.f2).equals(_dafny.MultiSet.fromElements(new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_100_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_100_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_100_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v9).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v9).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v24).equals(_dafny.Set.fromElements(new BigNumber(2), new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v98).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(545),new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_v99).dtor_cf12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(545),new BigNumber(545)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_224_v100, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_v101).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_226_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_235_v110).dtor_cf63, _dafny.Seq.of(_dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v110).dtor_cf64));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_238_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v144)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_313_v157).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_314_v158).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v159).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v159).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v159).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_332_v170).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("kmhjp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC2() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D0(3);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4;
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
    static create_DC6(cf6, cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC7(cf9, cf10, cf11) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO, false, false);
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
    static create_DC9() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC8(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ")";
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
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9();
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
    static create_DC11(cf14, cf15, cf16, cf17) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf13) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + this.cf17.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(false, false, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC14(cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC15() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D5(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15";
      } else if (this.$tag === 2) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, false);
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
    static create_DC16(cf22) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC17(cf23) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf23) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18();
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
    static create_DC19(cf24) {
      let $dt = new D8(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf26, cf27, cf28) {
      let $dt = new D9(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC22(cf29) {
      let $dt = new D9(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC20(cf25) {
      let $dt = new D9(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC21(false, _dafny.Seq.of(), false);
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
    static create_DC24(cf31, cf32, cf33) {
      let $dt = new D10(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC23(cf30) {
      let $dt = new D10(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + this.cf33.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC24(_dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC25(cf34) {
      let $dt = new D11(0);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf34) + ")";
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
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf36, cf37, cf38, cf39) {
      let $dt = new D12(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC26(cf35) {
      let $dt = new D12(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC28(cf40) {
      let $dt = new D12(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC27(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC30(cf42, cf43, cf44, cf45) {
      let $dt = new D13(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC29(cf41) {
      let $dt = new D13(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC31(cf46) {
      let $dt = new D13(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC30(false, false, _dafny.Set.Empty, null);
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
    static create_DC32(cf47) {
      let $dt = new D14(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf47) + ")";
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
    static create_DC33(cf48) {
      let $dt = new D15(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf48) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC35(cf50, cf51, cf52, cf53, cf54) {
      let $dt = new D16(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC34(cf49) {
      let $dt = new D16(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC34" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf49 === other.cf49;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC35(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC36(cf55) {
      let $dt = new D17(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC36" + "(" + _dafny.toString(this.cf55) + ")";
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf56) {
      let $dt = new D18(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC37" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
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
    static create_DC39(cf58, cf59) {
      let $dt = new D19(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC38(cf57) {
      let $dt = new D19(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC40(cf60) {
      let $dt = new D19(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC39" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC38" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC40" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC39(false, _dafny.ZERO);
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
    static create_DC41(cf61) {
      let $dt = new D20(0);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC41" + "(" + _dafny.toString(this.cf61) + ")";
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
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC43(cf63, cf64) {
      let $dt = new D21(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC42(cf62) {
      let $dt = new D21(1);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC43" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC42" + "(" + _dafny.toString(this.cf62) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC43(_dafny.Seq.of(), false);
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
    static create_DC44(cf65) {
      let $dt = new D22(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC44" + "(" + _dafny.toString(this.cf65) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC46() {
      let $dt = new D23(0);
      return $dt;
    }
    static create_DC47(cf67, cf68, cf69, cf70) {
      let $dt = new D23(1);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC45(cf66) {
      let $dt = new D23(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC46";
      } else if (this.$tag === 1) {
        return "D23.DC47" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC45" + "(" + _dafny.toString(this.cf66) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67) && this.cf68 === other.cf68 && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf66 === other.cf66;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC46();
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
    static create_DC49() {
      let $dt = new D24(0);
      return $dt;
    }
    static create_DC50() {
      let $dt = new D24(1);
      return $dt;
    }
    static create_DC48(cf71) {
      let $dt = new D24(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC51(cf72) {
      let $dt = new D24(3);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get is_DC51() { return this.$tag === 3; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC49";
      } else if (this.$tag === 1) {
        return "D24.DC50";
      } else if (this.$tag === 2) {
        return "D24.DC48" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC51" + "(" + _dafny.toString(this.cf72) + ")";
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
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC49();
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
    static create_DC53(cf74, cf75) {
      let $dt = new D25(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC52(cf73) {
      let $dt = new D25(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC54(cf76) {
      let $dt = new D25(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC53" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC52" + "(" + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC54" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf74, other.cf74) && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf73 === other.cf73;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC53(_dafny.Map.Empty, _dafny.ZERO);
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
    static create_DC56(cf78, cf79) {
      let $dt = new D26(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC57(cf80, cf81) {
      let $dt = new D26(1);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC55(cf77) {
      let $dt = new D26(2);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC56" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC57" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC55" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf77 === other.cf77;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC56(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC59(cf83, cf84, cf85) {
      let $dt = new D27(0);
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC58(cf82) {
      let $dt = new D27(1);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC59" + "(" + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC58" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf83 === other.cf83 && this.cf84 === other.cf84 && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC59(false, [], false);
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
    static create_DC60(cf86) {
      let $dt = new D28(0);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC60" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf86, other.cf86);
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
          return D28.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.Seq.of();
      this.f2 = _dafny.MultiSet.Empty;
      this.f3 = _dafny.ZERO;
      this.f4 = false;
      this.f8 = _dafny.ZERO;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f5 = false;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
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
    fm19(globalState) {
      let _this = this;
      return ((new BigNumber(-686)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(568)), function (_333_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length))).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(529), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).cardinality()));
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
    fm2(globalState) {
      let _this = this;
      return function () {
        let _coll56 = new _dafny.Map();
        for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-491), new BigNumber(-220))) {
          let _334_v0 = _compr_56;
          if (((new BigNumber(-491)).isLessThanOrEqualTo(_334_v0)) && ((_334_v0).isLessThan(new BigNumber(-220)))) {
            _coll56.push([(_334_v0).multipliedBy(new BigNumber(697)),new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(!(true), true))).length)]);
          }
        }
        return _coll56;
      }();
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return (!(true) || (!(true))) && (_dafny.areEqual(_dafny.Seq.of(true, !(false)), _dafny.Seq.of(true)));
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      if (false) {
        let _335_v0;
        _335_v0 = new BigNumber(847);
        let _336_v1;
        _336_v1 = _dafny.Map.Empty.slice().updateUnsafe(_335_v0,_335_v0);
        let _337_v2;
        _337_v2 = _module.D3.create_DC8(_336_v1);
        let _338_v3;
        _338_v3 = true;
        let _pat_let_tv10 = _336_v1;
        let _339_v4;
        let _nw62 = Array((new BigNumber(13)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _337_v2;
        _nw62[(_dafny.ONE).toNumber()] = ((!(_338_v3)) ? (_337_v2) : (_337_v2));
        _nw62[(new BigNumber(2)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(3)).toNumber()] = _module.D3.create_DC8(_336_v1);
        _nw62[(new BigNumber(4)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(5)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(6)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(7)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(8)).toNumber()] = _module.__default.fm24(globalState);
        _nw62[(new BigNumber(9)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(10)).toNumber()] = ((false) ? (function (_pat_let9_0) {
          return function (_340_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_341_dt__update_hcf12_h0) {
                return _module.D3.create_DC8(_341_dt__update_hcf12_h0);
              }(_pat_let10_0);
            }(_pat_let_tv10);
          }(_pat_let9_0);
        }(_337_v2)) : (_337_v2));
        _nw62[(new BigNumber(11)).toNumber()] = _337_v2;
        _nw62[(new BigNumber(12)).toNumber()] = _337_v2;
        _339_v4 = _nw62;
        let _index55 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_339_v4).length));
        (_339_v4)[_index55] = _337_v2;
        let _index56 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_339_v4).length));
        (_339_v4)[_index56] = _337_v2;
        let _342_v5;
        let _nw63 = new _module.C0();
        _nw63.__ctor();
        _342_v5 = _nw63;
        let _343_v6;
        _343_v6 = _dafny.Set.fromElements(_335_v0, _335_v0, _module.__default.safeDivisionInt(_335_v0, _335_v0));
        let _344_v7;
        _344_v7 = _dafny.Map.Empty.slice().updateUnsafe(_338_v3,_343_v6);
        _343_v6 = ((_338_v3) ? ((((_344_v7).contains(_338_v3)) ? ((_344_v7).get(_338_v3)) : (_343_v6))) : ((_343_v6).Difference(_343_v6)));
        let _345_v8;
        _345_v8 = _module.D3.create_DC9();
        let _source10 = _345_v8;
        if (_source10.is_DC9) {
          let _346_v9;
          _346_v9 = _dafny.Set.fromElements(_338_v3);
          let _347_v10;
          let _init8 = ((_348_v0, _349_v3) => function (_350_i0) {
            return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_348_v0, _348_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_349_v3)).length)), _348_v0))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_348_v0)));
          })(_335_v0, _338_v3);
          let _nw64 = Array((new BigNumber(15)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw64.length); _i0_8++) {
            _nw64[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _347_v10 = _nw64;
          let _351_v11;
          _351_v11 = _dafny.Seq.of(_335_v0, _335_v0, _335_v0, _335_v0, _335_v0);
          let _index57 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_347_v10).length));
          (_347_v10)[_index57] = _dafny.MultiSet.FromArray(_351_v11);
          let _352_v12;
          _352_v12 = _dafny.Seq.of(_338_v3);
          let _353_v13;
          _353_v13 = _dafny.MultiSet.fromElements(_335_v0, _335_v0, new BigNumber((_346_v9).length), _335_v0);
          let _354_v14;
          _354_v14 = _module.D4.create_DC10(_352_v12);
          let _355_v15;
          _355_v15 = _dafny.MultiSet.fromElements(_354_v14);
          let _356_v16;
          _356_v16 = _dafny.Map.Empty.slice().updateUnsafe(_355_v15,_335_v0);
          let _357_v17;
          _357_v17 = _dafny.Map.Empty.slice().updateUnsafe(_335_v0,_356_v16);
          let _index58 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_347_v10).length));
          let _rhs44 = new BigNumber((_352_v12).length);
          let _rhs45 = _335_v0;
          let _rhs46 = _346_v9;
          let _rhs47 = ((_353_v13).Difference(_353_v13)).Intersect((_353_v13).Intersect(_353_v13));
          let _rhs48 = _module.__default.safeModuloInt(_335_v0, new BigNumber(((((_357_v17).contains((_dafny.ZERO).minus(_335_v0))) ? ((_357_v17).get((_dafny.ZERO).minus(_335_v0))) : (_356_v16))).length));
          let _lhs35 = globalState;
          let _lhs36 = globalState;
          let _lhs37 = _347_v10;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_347_v10).length));
          let _lhs39 = globalState;
          _lhs35.f8 = _rhs44;
          _lhs36.f8 = _rhs45;
          _346_v9 = _rhs46;
          _lhs37[_lhs38] = _rhs47;
          _lhs39.f8 = _rhs48;
          let _358_v18;
          _358_v18 = _dafny.Seq.UnicodeFromString("uou");
          let _359_v19;
          _359_v19 = new _dafny.CodePoint('r'.codePointAt(0));
          let _rhs49 = _358_v18;
          let _rhs50 = !_dafny.Seq.contains(_358_v18, ((_338_v3) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (_359_v19)));
          let _lhs40 = globalState;
          _358_v18 = _rhs49;
          _lhs40.f4 = _rhs50;
          (globalState).f3 = (_342_v5).fm19(globalState);
          let _360_v20;
          _360_v20 = _module.D2.create_DC7(_359_v19, _338_v3, _335_v0);
          let _361_v21;
          _361_v21 = _module.D4.create_DC11(_338_v3, _338_v3, (_351_v11)[_module.__default.safeIndex(_335_v0, new BigNumber((_351_v11).length))], _dafny.Seq.update(_dafny.Seq.UnicodeFromString("e"), _module.__default.safeIndex(_335_v0, new BigNumber((_dafny.Seq.UnicodeFromString("e")).length)), _359_v19));
          let _362_v22;
          _362_v22 = _351_v11;
          let _363_v23;
          _363_v23 = _dafny.Map.Empty.slice().updateUnsafe(_335_v0,_343_v6);
          let _364_v24;
          _364_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_335_v0),_353_v13);
          let _365_v25;
          _365_v25 = _dafny.Seq.of(_364_v24);
          let _366_v26;
          let _nw65 = Array((new BigNumber(28)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_335_v0, _335_v0);
          _nw65[(_dafny.ONE).toNumber()] = _335_v0;
          _nw65[(new BigNumber(2)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_335_v0,_335_v0)).length);
          _nw65[(new BigNumber(4)).toNumber()] = (_335_v0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), ((_367_v19) => function (_368_i1) {
            return _367_v19;
          })(_359_v19))).length));
          _nw65[(new BigNumber(5)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(6)).toNumber()] = ((_338_v3) ? ((_360_v20).dtor_cf11) : ((_361_v21).dtor_cf16));
          _nw65[(new BigNumber(7)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(8)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(9)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(10)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(11)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(12)).toNumber()] = new BigNumber(-245);
          _nw65[(new BigNumber(13)).toNumber()] = (((_336_v1).contains((_dafny.ZERO).minus((((_336_v1).contains(_335_v0)) ? ((_336_v1).get(_335_v0)) : (_335_v0))))) ? ((_336_v1).get((_dafny.ZERO).minus((((_336_v1).contains(_335_v0)) ? ((_336_v1).get(_335_v0)) : (_335_v0))))) : (_335_v0));
          _nw65[(new BigNumber(14)).toNumber()] = ((_dafny.ZERO).minus(((((_347_v10)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_347_v10).length))]).contains(_335_v0)) ? (((_347_v10)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_347_v10).length))]).get(_335_v0)) : (_335_v0)))).minus(_335_v0);
          _nw65[(new BigNumber(15)).toNumber()] = (_335_v0).plus(_335_v0);
          _nw65[(new BigNumber(16)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(17)).toNumber()] = new BigNumber(((_362_v22)).length);
          _nw65[(new BigNumber(18)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(19)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(20)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_363_v23).length), _335_v0);
          _nw65[(new BigNumber(22)).toNumber()] = (_335_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("nalyv")).length));
          _nw65[(new BigNumber(23)).toNumber()] = ((true) ? ((_342_v5).fm19(globalState)) : (_335_v0));
          _nw65[(new BigNumber(24)).toNumber()] = _335_v0;
          _nw65[(new BigNumber(25)).toNumber()] = new BigNumber(715);
          _nw65[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-239)), ((_369_v19) => function (_370_i2) {
            return _369_v19;
          })(_359_v19))).length), new BigNumber(((_365_v25)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_365_v25).length))]).length));
          _nw65[(new BigNumber(27)).toNumber()] = (_dafny.ZERO).minus(_335_v0);
          _366_v26 = _nw65;
          let _index59 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_366_v26).length));
          (_366_v26)[_index59] = _module.__default.safeDivisionInt(_335_v0, _335_v0);
          let _index60 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_366_v26).length));
          (_366_v26)[_index60] = (_dafny.ZERO).minus((new BigNumber(865)).minus(new BigNumber((_358_v18).length)));
        } else {
          let _371___mcc_h0 = (_source10).cf12;
          let _372_cf12 = _371___mcc_h0;
          let _373_v27;
          let _nw66 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _373_v27 = _nw66;
          let _374_v28;
          _374_v28 = _dafny.Seq.UnicodeFromString("pqitorcr");
          let _index61 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_373_v27).length));
          (_373_v27)[_index61] = (new BigNumber((_374_v28).length)).plus(_335_v0);
          let _index62 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_373_v27).length));
          (_373_v27)[_index62] = _335_v0;
          let _375_v29;
          _375_v29 = _dafny.Set.fromElements(_338_v3, _338_v3, _338_v3);
          _338_v3 = (_375_v29).IsSubsetOf(_dafny.Set.fromElements(!(_338_v3)));
          (globalState).f8 = (_373_v27)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_373_v27).length))];
          let _376_v30;
          let _init9 = ((_377_v0) => function (_378_i3) {
            return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_377_v0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber(-308)));
          })(_335_v0);
          let _nw67 = Array((new BigNumber(17)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw67.length); _i0_9++) {
            _nw67[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _376_v30 = _nw67;
          _376_v30 = _376_v30;
        }
        let _379_v31;
        _379_v31 = _dafny.Seq.of(false, _338_v3);
        if ((_379_v31)[_module.__default.safeIndex(_335_v0, new BigNumber((_379_v31).length))]) {
          let _380_v32;
          _380_v32 = _dafny.Seq.UnicodeFromString("prrqlk");
          let _381_v33;
          let _nw68 = Array((new BigNumber(16)).toNumber()).fill(false);
          _381_v33 = _nw68;
          let _382_v34;
          _382_v34 = _dafny.Map.Empty.slice().updateUnsafe(_338_v3,_381_v33);
          let _383_v35;
          _383_v35 = new _dafny.CodePoint('v'.codePointAt(0));
          (globalState).f8 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm1(_335_v0, new _dafny.CodePoint('e'.codePointAt(0)), _338_v3, _336_v1, globalState), _module.__default.safeIndex(new BigNumber(-7), new BigNumber((_module.__default.fm1(_335_v0, new _dafny.CodePoint('e'.codePointAt(0)), _338_v3, _336_v1, globalState)).length)), new BigNumber((_380_v32).length)), _module.__default.fm1(new BigNumber((_382_v34).length), _383_v35, false, _336_v1, globalState))).length);
          let _384_v36;
          _384_v36 = _dafny.Map.Empty.slice().updateUnsafe(_338_v3,_338_v3);
          let _385_v37;
          _385_v37 = _dafny.Map.Empty.slice().updateUnsafe(_384_v36,_338_v3);
          let _386_v38;
          let _nw69 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _386_v38 = _nw69;
          let _index63 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_386_v38).length));
          (_386_v38)[_index63] = _383_v35;
          let _index64 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_386_v38).length));
          let _rhs51 = _module.__default.fm25(globalState);
          let _rhs52 = ((_338_v3) ? ((_335_v0).multipliedBy((((_336_v1).contains(_335_v0)) ? ((_336_v1).get(_335_v0)) : (_335_v0)))) : (_335_v0));
          let _rhs53 = (_385_v37).Merge(_385_v37);
          let _rhs54 = _383_v35;
          let _rhs55 = _module.__default.safeDivisionInt(_335_v0, (_335_v0).plus((_dafny.ZERO).minus(_335_v0)));
          let _lhs41 = globalState;
          let _lhs42 = _386_v38;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_386_v38).length));
          let _lhs44 = globalState;
          _380_v32 = _rhs51;
          _lhs41.f3 = _rhs52;
          _385_v37 = _rhs53;
          _lhs42[_lhs43] = _rhs54;
          _lhs44.f8 = _rhs55;
          let _387_v39;
          _387_v39 = _dafny.MultiSet.fromElements(_335_v0);
          let _388_v40;
          _388_v40 = _dafny.MultiSet.fromElements(!(_338_v3));
          let _389_v41;
          _389_v41 = _dafny.Seq.of((_dafny.ZERO).minus(_335_v0), new BigNumber((_387_v39).cardinality()), _335_v0, _335_v0, (((_388_v40).contains(_338_v3)) ? ((_388_v40).get(_338_v3)) : ((_dafny.ZERO).minus(_335_v0))));
          (globalState).f8 = (_389_v41)[_module.__default.safeIndex((_335_v0).multipliedBy(_335_v0), new BigNumber((_389_v41).length))];
          _338_v3 = _338_v3;
          (globalState).f8 = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_383_v35,_dafny.areEqual(_380_v32, _380_v32))).length));
        } else {
          let _390_v42;
          let _nw70 = Array((new BigNumber(24)).toNumber()).fill([]);
          _390_v42 = _nw70;
          let _391_v43;
          let _nw71 = Array((new BigNumber(11)).toNumber()).fill([]);
          _391_v43 = _nw71;
          let _index65 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_390_v42).length));
          (_390_v42)[_index65] = _391_v43;
          let _392_v44;
          _392_v44 = _module.D5.create_DC14(_335_v0, !(_338_v3));
          let _393_v45;
          _393_v45 = _dafny.Seq.of(_392_v44);
          let _394_v46;
          _394_v46 = _dafny.Set.fromElements(_338_v3, _338_v3, _338_v3, _338_v3);
          let _395_v47;
          _395_v47 = _module.D9.create_DC20(_394_v46);
          let _396_v48;
          _396_v48 = _dafny.Seq.UnicodeFromString("fw");
          let _397_v49;
          _397_v49 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_396_v48).length)));
          let _pat_let_tv11 = _394_v46;
          let _pat_let_tv12 = _397_v49;
          let _398_v50;
          let _nw72 = Array((new BigNumber(19)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _338_v3;
          _nw72[(_dafny.ONE).toNumber()] = _338_v3;
          _nw72[(new BigNumber(2)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(3)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(4)).toNumber()] = (_338_v3) || (_338_v3);
          _nw72[(new BigNumber(5)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(6)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(7)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(8)).toNumber()] = !(_338_v3);
          _nw72[(new BigNumber(9)).toNumber()] = !(_338_v3);
          _nw72[(new BigNumber(10)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(11)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(12)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(13)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(14)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(15)).toNumber()] = !(((function (_pat_let11_0) {
            return function (_399_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_400_dt__update_hcf25_h0) {
                  return _module.D9.create_DC20(_400_dt__update_hcf25_h0);
                }(_pat_let12_0);
              }(_pat_let_tv11);
            }(_pat_let11_0);
          }(_395_v47)).dtor_cf25).IsProperSubsetOf(_394_v46));
          _nw72[(new BigNumber(16)).toNumber()] = _338_v3;
          _nw72[(new BigNumber(17)).toNumber()] = (function (_pat_let13_0) {
            return function (_401_dt__update__tmp_h2) {
              return function (_pat_let14_0) {
                return function (_402_dt__update_hcf20_h0) {
                  return _module.D5.create_DC14(_402_dt__update_hcf20_h0, (_401_dt__update__tmp_h2).dtor_cf21);
                }(_pat_let14_0);
              }(new BigNumber((_pat_let_tv12).length));
            }(_pat_let13_0);
          }(_392_v44)).dtor_cf21;
          _nw72[(new BigNumber(18)).toNumber()] = _338_v3;
          _398_v50 = _nw72;
          let _index66 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length));
          (_398_v50)[_index66] = !(_338_v3);
          let _index67 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_390_v42).length));
          let _index68 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length));
          let _rhs56 = _338_v3;
          let _rhs57 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("hqv"), _396_v48);
          let _rhs58 = _391_v43;
          let _rhs59 = _dafny.Seq.Concat(_393_v45, _393_v45);
          let _rhs60 = false;
          let _lhs45 = _390_v42;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_390_v42).length));
          let _lhs47 = _398_v50;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length));
          _338_v3 = _rhs56;
          r0 = _rhs57;
          _lhs45[_lhs46] = _rhs58;
          _393_v45 = _rhs59;
          _lhs47[_lhs48] = _rhs60;
          let _403_v51;
          let _nw73 = new _module.C0();
          _nw73.__ctor();
          _403_v51 = _nw73;
          let _404_v52;
          let _nw74 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _404_v52 = _nw74;
          let _index69 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_404_v52).length));
          (_404_v52)[_index69] = _335_v0;
          let _index70 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_404_v52).length));
          let _index71 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length));
          let _rhs61 = (new BigNumber(901)).multipliedBy((_335_v0).plus(_335_v0));
          let _rhs62 = _335_v0;
          let _rhs63 = (_342_v5).fm19(globalState);
          let _rhs64 = (_398_v50)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length))];
          let _lhs49 = globalState;
          let _lhs50 = globalState;
          let _lhs51 = _404_v52;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_404_v52).length));
          let _lhs53 = _398_v50;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length));
          _lhs49.f8 = _rhs61;
          _lhs50.f3 = _rhs62;
          _lhs51[_lhs52] = _rhs63;
          _lhs53[_lhs54] = _rhs64;
          let _405_v53;
          _405_v53 = _dafny.MultiSet.fromElements(_338_v3);
          _405_v53 = _dafny.MultiSet.fromElements(_338_v3, (_398_v50)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length))], !(_338_v3), (((_379_v31)[_module.__default.safeIndex((_404_v52)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_404_v52).length))], new BigNumber((_379_v31).length))]) ? ((_398_v50)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_398_v50).length))]) : (_338_v3)));
          let _rhs65 = _338_v3;
          let _rhs66 = false;
          let _rhs67 = _404_v52;
          let _lhs55 = globalState;
          _lhs55.f4 = _rhs65;
          r0 = _rhs66;
          _404_v52 = _rhs67;
        }
      } else {
        let _406_v54;
        _406_v54 = false;
        if (_406_v54) {
          let _407_v55;
          _407_v55 = new BigNumber(-909);
          let _408_v56;
          _408_v56 = _dafny.Map.Empty.slice().updateUnsafe(_407_v55,_dafny.Seq.of(!(true)));
          let _409_v57;
          _409_v57 = _module.D4.create_DC10(_dafny.Seq.of(false, _406_v54));
          let _410_v58;
          _410_v58 = _dafny.Seq.of(_module.D4.create_DC10((((_408_v56).contains(_407_v55)) ? ((_408_v56).get(_407_v55)) : (_dafny.Seq.of(_406_v54)))), _409_v57, _409_v57, _409_v57, _409_v57);
          _410_v58 = _410_v58;
          let _411_v59;
          _411_v59 = _dafny.Seq.of(_407_v55);
          let _412_v60;
          _412_v60 = _dafny.Map.Empty.slice().updateUnsafe(_406_v54,_406_v54);
          let _413_v61;
          _413_v61 = _dafny.MultiSet.fromElements(new BigNumber(-813), _407_v55, new BigNumber((_412_v60).length));
          let _414_v62;
          _414_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_411_v59, _413_v61, _406_v54, globalState),_407_v55);
          let _415_v63;
          _415_v63 = new _dafny.CodePoint('y'.codePointAt(0));
          let _416_v64;
          _416_v64 = _dafny.Map.Empty.slice().updateUnsafe((_407_v55).minus(new BigNumber((_414_v62).length)),((_406_v54) ? (_415_v63) : (_415_v63)));
          _416_v64 = (_416_v64).update((_407_v55).multipliedBy(_407_v55), _415_v63);
          let _417_v65;
          let _init10 = ((_418_v63) => function (_419_i4) {
            return _418_v63;
          })(_415_v63);
          let _nw75 = Array((new BigNumber(26)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw75.length); _i0_10++) {
            _nw75[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _417_v65 = _nw75;
          let _index72 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_417_v65).length));
          (_417_v65)[_index72] = _415_v63;
          let _420_v66;
          let _nw76 = Array((new BigNumber(12)).toNumber()).fill(false);
          _420_v66 = _nw76;
          let _index73 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_417_v65).length));
          let _rhs68 = (_407_v55).plus(_407_v55);
          let _rhs69 = _415_v63;
          let _rhs70 = _420_v66;
          let _lhs56 = globalState;
          let _lhs57 = _417_v65;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_417_v65).length));
          _lhs56.f3 = _rhs68;
          _lhs57[_lhs58] = _rhs69;
          _420_v66 = _rhs70;
          let _421_v67;
          let _nw77 = new _module.C0();
          _nw77.__ctor();
          _421_v67 = _nw77;
          (globalState).f8 = (_411_v59)[_module.__default.safeIndex((_dafny.ZERO).minus(_407_v55), new BigNumber((_411_v59).length))];
        } else {
          let _422_v68;
          _422_v68 = _dafny.Set.fromElements(true);
          let _423_v69;
          _423_v69 = _dafny.Seq.of(_406_v54, (_dafny.Set.fromElements(true, _406_v54)).IsDisjointFrom(_422_v68));
          let _424_v70;
          _424_v70 = new BigNumber(70);
          (globalState).f4 = (_423_v69)[_module.__default.safeIndex(_424_v70, new BigNumber((_423_v69).length))];
          r0 = _406_v54;
          let _425_v71;
          let _nw78 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _425_v71 = _nw78;
          let _426_v72;
          _426_v72 = new _dafny.CodePoint('m'.codePointAt(0));
          let _index74 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_425_v71).length));
          (_425_v71)[_index74] = _426_v72;
          let _index75 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_425_v71).length));
          (_425_v71)[_index75] = _426_v72;
          let _index76 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_425_v71).length));
          (_425_v71)[_index76] = (_425_v71)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_425_v71).length))];
          let _427_v73;
          _427_v73 = _dafny.Seq.UnicodeFromString("jcrxtqw");
          _427_v73 = _427_v73;
        }
        let _428_v74;
        _428_v74 = _dafny.Seq.UnicodeFromString("ilxgf");
        let _429_v75;
        _429_v75 = new BigNumber(857);
        let _430_v76;
        _430_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_428_v74).length),_429_v75);
        let _431_v78;
        _431_v78 = new _dafny.CodePoint('u'.codePointAt(0));
        _430_v76 = (_430_v76).update(_module.__default.safeDivisionInt(new BigNumber((function () {
          let _coll57 = new _dafny.Map();
          for (const _compr_57 of _dafny.IntegerRange(new BigNumber(513), new BigNumber(938))) {
            let _432_v77 = _compr_57;
            if (((new BigNumber(513)).isLessThanOrEqualTo(_432_v77)) && ((_432_v77).isLessThan(new BigNumber(938)))) {
              _coll57.push([_module.__default.safeModuloInt(_432_v77, new BigNumber(((_module.D10.create_DC23(_dafny.MultiSet.FromArray(_dafny.Seq.of(_428_v74)))).dtor_cf30).cardinality())),_406_v54]);
            }
          }
          return _coll57;
        }()).length), _429_v75), new BigNumber((((_406_v54) ? (_dafny.Map.Empty.slice().updateUnsafe(_431_v78,_431_v78)) : (_dafny.Map.Empty.slice().updateUnsafe(_431_v78,_431_v78)))).length));
        let _433_v79;
        _433_v79 = _dafny.Seq.of(_406_v54, _406_v54);
        let _434_v80;
        _434_v80 = _dafny.Set.fromElements(_406_v54);
        let _435_v81;
        _435_v81 = _dafny.Seq.of(((_406_v54) ? (_dafny.Set.fromElements(_406_v54)) : (_434_v80)));
        let _rhs71 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_429_v75, new BigNumber(669)), (new BigNumber((_433_v79).length)).multipliedBy(_429_v75)));
        let _rhs72 = new BigNumber((_435_v81).length);
        let _lhs59 = globalState;
        let _lhs60 = globalState;
        _lhs59.f8 = _rhs71;
        _lhs60.f3 = _rhs72;
        let _436_v82;
        _436_v82 = _module.D3.create_DC8(_430_v76);
        let _437_v83;
        _437_v83 = _dafny.MultiSet.fromElements(_433_v79);
        let _rhs73 = _429_v75;
        let _rhs74 = _module.__default.fm26(_436_v82, _429_v75, _437_v83, globalState);
        let _lhs61 = globalState;
        _lhs61.f3 = _rhs73;
        _429_v75 = _rhs74;
        let _438_v84;
        _438_v84 = _dafny.Seq.of(_429_v75);
        let _439_v85;
        _439_v85 = _module.D5.create_DC15();
        let _source11 = ((!_dafny.Seq.contains(_438_v84, new BigNumber(733))) ? (_439_v85) : (_439_v85));
        if (_source11.is_DC14) {
          let _440___mcc_h1 = (_source11).cf20;
          let _441___mcc_h2 = (_source11).cf21;
          let _442_cf21 = _441___mcc_h2;
          let _443_cf20 = _440___mcc_h1;
          (globalState).f4 = _442_cf21;
          let _444_v86;
          let _nw79 = Array((new BigNumber(23)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = _443_cf20;
          _nw79[(_dafny.ONE).toNumber()] = new BigNumber(340);
          _nw79[(new BigNumber(2)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(3)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(4)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(5)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(6)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(7)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(8)).toNumber()] = _module.__default.fm26(_436_v82, _443_cf20, _437_v83, globalState);
          _nw79[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_443_cf20);
          _nw79[(new BigNumber(10)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(11)).toNumber()] = new BigNumber(78);
          _nw79[(new BigNumber(12)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(13)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(14)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(15)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(16)).toNumber()] = new BigNumber((_428_v74).length);
          _nw79[(new BigNumber(17)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(18)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(19)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(20)).toNumber()] = _429_v75;
          _nw79[(new BigNumber(21)).toNumber()] = _443_cf20;
          _nw79[(new BigNumber(22)).toNumber()] = new BigNumber(-573);
          _444_v86 = _nw79;
          let _445_v87;
          _445_v87 = _dafny.Map.Empty.slice().updateUnsafe(_429_v75,_444_v86);
          let _446_v88;
          _446_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((((_445_v87).contains(_443_cf20)) ? ((_445_v87).get(_443_cf20)) : (_444_v86)), _444_v86),!(_406_v54));
          let _447_v89;
          _447_v89 = _dafny.MultiSet.fromElements(_444_v86);
          _442_cf21 = (((_446_v88).contains((_447_v89).Union(_447_v89))) ? ((_446_v88).get((_447_v89).Union(_447_v89))) : (_406_v54));
          r0 = _406_v54;
          (globalState).f8 = _443_cf20;
        } else if (_source11.is_DC15) {
          (globalState).f3 = _429_v75;
          (globalState).f3 = _429_v75;
          let _448_v90;
          _448_v90 = _module.D2.create_DC7(_431_v78, _406_v54, _429_v75);
          _448_v90 = _module.D2.create_DC7(_431_v78, _406_v54, _429_v75);
          (globalState).f3 = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_428_v74).length)).plus(_429_v75)));
        } else {
          let _449___mcc_h3 = (_source11).cf19;
          let _450_cf19 = _449___mcc_h3;
          let _451_v91;
          _451_v91 = _dafny.Map.Empty.slice().updateUnsafe(_434_v80,_module.__default.fm26(_436_v82, new BigNumber((_428_v74).length), _437_v83, globalState));
          r0 = (_451_v91).equals(_451_v91);
          let _452_v92;
          _452_v92 = _dafny.MultiSet.fromElements(_406_v54);
          r0 = (new BigNumber((_452_v92).cardinality())).isLessThan((_429_v75).plus(new BigNumber((_428_v74).length)));
          r0 = _406_v54;
          let _453_v93;
          _453_v93 = _dafny.MultiSet.fromElements(new BigNumber(627));
          (_this).m13(_429_v75, (_453_v93).Intersect(_453_v93), !((_dafny.Set.fromElements(false)).IsDisjointFrom(_434_v80)), globalState);
        }
      }
      let _454_v94;
      _454_v94 = true;
      (globalState).f4 = _454_v94;
      let _455_v95;
      let _init11 = ((_456_v94) => function (_457_i5) {
        return _456_v94;
      })(_454_v94);
      let _nw80 = Array((new BigNumber(17)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw80.length); _i0_11++) {
        _nw80[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _455_v95 = _nw80;
      let _index77 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_455_v95).length));
      (_455_v95)[_index77] = ((_454_v94) ? (_454_v94) : (_454_v94));
      let _458_v96;
      let _nw81 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _458_v96 = _nw81;
      let _459_v97;
      let _nw82 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
      _459_v97 = _nw82;
      let _460_v98;
      _460_v98 = _dafny.Map.Empty.slice().updateUnsafe(_458_v96,_459_v97);
      let _index78 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_455_v95).length));
      (_455_v95)[_index78] = (_460_v98).equals(_460_v98);
      let _461_i6;
      _461_i6 = _dafny.ZERO;
      L3: {
        while ((_455_v95)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_455_v95).length))]) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_461_i6)) {
              break L3;
            }
            _461_i6 = (_461_i6).plus(_dafny.ONE);
            let _462_v99;
            _462_v99 = new BigNumber(-172);
            let _463_v100;
            _463_v100 = _dafny.Seq.of(_462_v99, _462_v99, _462_v99);
            let _464_v101;
            _464_v101 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_462_v99), _463_v100);
            let _465_v102;
            _465_v102 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_462_v99),_464_v101);
            let _466_v103;
            _466_v103 = _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_462_v99,_462_v99));
            let _467_v104;
            _467_v104 = _dafny.MultiSet.fromElements(_dafny.Seq.of(true));
            _465_v102 = (_465_v102).update(_module.__default.fm26(_466_v103, _462_v99, _467_v104, globalState), _464_v101);
            _462_v99 = _462_v99;
            let _468_v105;
            let _nw83 = new _module.C0();
            _nw83.__ctor();
            _468_v105 = _nw83;
            (globalState).f8 = ((_462_v99).plus(_462_v99)).plus(_462_v99);
          }
        }
      }
      let _469_v106;
      _469_v106 = new BigNumber(160);
      let _470_v107;
      _470_v107 = _dafny.Map.Empty.slice().updateUnsafe(_469_v106,_469_v106);
      (globalState).f8 = (_module.__default.safeModuloInt(_469_v106, _469_v106)).multipliedBy((new BigNumber(602)).multipliedBy(new BigNumber((_470_v107).length)));
      let _index79 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_459_v97).length));
      (_459_v97)[_index79] = _469_v106;
      let _471_v108;
      _471_v108 = _dafny.Seq.UnicodeFromString("p");
      let _472_v109;
      _472_v109 = _module.D2.create_DC6(_469_v106, true, (_455_v95)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_455_v95).length))]);
      let _index80 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_459_v97).length));
      let _rhs75 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_469_v106, (_469_v106).multipliedBy((_dafny.ZERO).minus(_469_v106))));
      let _rhs76 = (_455_v95)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_455_v95).length))];
      let _rhs77 = (_472_v109).dtor_cf8;
      let _rhs78 = _471_v108;
      let _rhs79 = false;
      let _lhs62 = _459_v97;
      let _lhs63 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_459_v97).length));
      let _lhs64 = globalState;
      _lhs62[_lhs63] = _rhs75;
      r1 = _rhs76;
      _lhs64.f4 = _rhs77;
      _471_v108 = _rhs78;
      r1 = _rhs79;
      r0 = _454_v94;
      let _pat_let_tv13 = _459_v97;
      let _pat_let_tv14 = _459_v97;
      let _pat_let_tv15 = _454_v94;
      let _pat_let_tv16 = _455_v95;
      let _pat_let_tv17 = _455_v95;
      let _pat_let_tv18 = _454_v94;
      r1 = function (_source12) {
        if (_source12.is_DC14) {
          let _473___mcc_h4 = (_source12).cf20;
          let _474___mcc_h5 = (_source12).cf21;
          let _475_cf21 = _474___mcc_h5;
          let _476_cf20 = _473___mcc_h4;
          return (_module.D2.create_DC6((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_pat_let_tv13).length))], true, _475_cf21)).dtor_cf8;
        } else if (_source12.is_DC15) {
          return (_dafny.Set.fromElements(_pat_let_tv15)).IsSubsetOf(_dafny.Set.fromElements((_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_pat_let_tv16).length))]));
        } else {
          let _477___mcc_h6 = (_source12).cf19;
          let _478_cf19 = _477___mcc_h6;
          return _pat_let_tv18;
        }
      }(_module.D5.create_DC15());
      return [r0, r1];
    }
    m13(p0, p1, p2, globalState) {
      let _this = this;
      let _479_v0;
      _479_v0 = _dafny.Seq.UnicodeFromString("eahalakud");
      if ((p1).contains((new BigNumber((_479_v0).length)).minus(new BigNumber(150)))) {
        let _480_v1;
        let _nw84 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _480_v1 = _nw84;
        let _481_v2;
        _481_v2 = _dafny.Seq.of(_480_v1);
        _481_v2 = _481_v2;
        let _index81 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length));
        (_480_v1)[_index81] = p0;
        let _index82 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length));
        (_480_v1)[_index82] = p0;
        let _index83 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length));
        (_480_v1)[_index83] = (_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))];
        let _init12 = ((_482_v1) => function (_483_i0) {
          return (_483_i0).plus((_482_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_482_v1).length))]);
        })(_480_v1);
        let _nw85 = Array((new BigNumber(24)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw85.length); _i0_12++) {
          _nw85[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _480_v1 = _nw85;
        let _484_v3;
        _484_v3 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(new BigNumber(-256)));
        let _485_v4;
        _485_v4 = _module.D9.create_DC21(false, _dafny.Seq.of(p2, p2, (_this).fm3(_484_v3, p1, p2, globalState), p2), p2);
        let _486_v5;
        _486_v5 = _dafny.Seq.of(!(p2), !((_this).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(586)), ((_487_v1) => function (_488_i1) {
          return (_487_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_487_v1).length))];
        })(_480_v1)), p1, !(false), globalState)), p2);
        let _pat_let_tv19 = p2;
        if (_dafny.areEqual((function (_pat_let15_0) {
          return function (_489_dt__update__tmp_h0) {
            return function (_pat_let16_0) {
              return function (_490_dt__update_hcf26_h0) {
                return _module.D9.create_DC21(_490_dt__update_hcf26_h0, (_489_dt__update__tmp_h0).dtor_cf27, (_489_dt__update__tmp_h0).dtor_cf28);
              }(_pat_let16_0);
            }(_pat_let_tv19);
          }(_pat_let15_0);
        }(_485_v4)).dtor_cf27, _dafny.Seq.Concat(_486_v5, _486_v5))) {
          (globalState).f3 = ((p2) ? (p0) : (new BigNumber(655)));
          _480_v1 = _480_v1;
          (globalState).f4 = p2;
          let _491_v6;
          _491_v6 = _module.D2.create_DC6(p0, (p0).isEqualTo(p0), p2);
          _491_v6 = _491_v6;
          (globalState).f8 = (_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))];
        } else {
          let _492_v7;
          let _nw86 = new _module.C0();
          _nw86.__ctor();
          _492_v7 = _nw86;
          let _493_v9;
          let _init13 = ((_494_p2) => function (_495_i2) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_494_p2,_494_p2)).Merge(function () {
              let _coll58 = new _dafny.Map();
              for (const _compr_58 of (_dafny.Set.fromElements(_494_p2, false)).Elements) {
                let _496_v8 = _compr_58;
                if ((_dafny.Set.fromElements(_494_p2, false)).contains(_496_v8)) {
                  _coll58.push([_496_v8,_494_p2]);
                }
              }
              return _coll58;
            }());
          })(p2);
          let _nw87 = Array((new BigNumber(20)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw87.length); _i0_13++) {
            _nw87[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _493_v9 = _nw87;
          let _497_v10;
          _497_v10 = p2;
          let _498_v11;
          _498_v11 = _module.D2.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,!(p2))).length), p2, p2);
          let _499_v12;
          _499_v12 = _dafny.Map.Empty.slice().updateUnsafe(_497_v10,!((_498_v11).dtor_cf8));
          let _index84 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_493_v9).length));
          (_493_v9)[_index84] = _499_v12;
          let _500_v14;
          _500_v14 = _dafny.Set.fromElements(_497_v10, _497_v10);
          let _index85 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_493_v9).length));
          (_493_v9)[_index85] = function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (function () {
              let _coll60 = new _dafny.Set();
              for (const _compr_60 of (_500_v14).Elements) {
                let _501_v15 = _compr_60;
                if ((_500_v14).contains(_501_v15)) {
                  _coll60.add(_501_v15);
                }
              }
              return _coll60;
            }()).Elements) {
              let _502_v13 = _compr_59;
              if ((function () {
                let _coll61 = new _dafny.Set();
                for (const _compr_61 of (_500_v14).Elements) {
                  let _503_v15 = _compr_61;
                  if ((_500_v14).contains(_503_v15)) {
                    _coll61.add(_503_v15);
                  }
                }
                return _coll61;
              }()).contains(_502_v13)) {
                _coll59.push([_502_v13,p2]);
              }
            }
            return _coll59;
          }();
          (globalState).f8 = (_492_v7).fm19(globalState);
          let _504_v16;
          _504_v16 = new _dafny.CodePoint('l'.codePointAt(0));
          let _505_v17;
          _505_v17 = _dafny.Map.Empty.slice().updateUnsafe(_492_v7,_504_v16);
          let _506_v18;
          _506_v18 = _dafny.Map.Empty.slice().updateUnsafe((_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))],new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-817)), ((_507_v16) => function (_508_i5) {
            return _507_v16;
          })(_504_v16))).length));
          let _509_v19;
          _509_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _510_v20;
          _510_v20 = _dafny.Map.Empty.slice().updateUnsafe(_509_v19,_504_v16);
          let _511_v21;
          let _nw88 = Array((new BigNumber(15)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = _484_v3;
          _nw88[(_dafny.ONE).toNumber()] = _484_v3;
          _nw88[(new BigNumber(2)).toNumber()] = _484_v3;
          _nw88[(new BigNumber(3)).toNumber()] = _484_v3;
          _nw88[(new BigNumber(4)).toNumber()] = ((!((_this).fm3(_484_v3, p1, p2, globalState))) ? (_484_v3) : (_484_v3));
          _nw88[(new BigNumber(5)).toNumber()] = _484_v3;
          _nw88[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(new BigNumber((_505_v17).length), (_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))]);
          _nw88[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), ((_512_v1) => function (_513_i3) {
            return (_512_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_512_v1).length))];
          })(_480_v1));
          _nw88[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(817)), function (_514_i4) {
            return new BigNumber(782);
          }), _484_v3);
          _nw88[(new BigNumber(9)).toNumber()] = _484_v3;
          _nw88[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm1((_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))], new _dafny.CodePoint('j'.codePointAt(0)), p2, _506_v18, globalState), _dafny.Seq.of((_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))], (_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))]));
          _nw88[(new BigNumber(11)).toNumber()] = _484_v3;
          _nw88[(new BigNumber(12)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus((_480_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_480_v1).length))]), new BigNumber((_510_v20).length));
          _nw88[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_484_v3, _484_v3);
          _nw88[(new BigNumber(14)).toNumber()] = _484_v3;
          _511_v21 = _nw88;
          let _index86 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_511_v21).length));
          (_511_v21)[_index86] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_515_v1) => function (_516_i6) {
            return (_515_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_515_v1).length))];
          })(_480_v1));
          let _index87 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_511_v21).length));
          (_511_v21)[_index87] = _484_v3;
          let _517_v22;
          let _nw89 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _517_v22 = _nw89;
          let _index88 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_517_v22).length));
          (_517_v22)[_index88] = _dafny.Seq.UnicodeFromString("fusvcouy");
          let _index89 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_517_v22).length));
          (_517_v22)[_index89] = _dafny.Seq.UnicodeFromString("k");
        }
      } else {
        let _518_v23;
        let _init14 = ((_519_p0) => function (_520_i7) {
          return (_520_i7).minus(_519_p0);
        })(p0);
        let _nw90 = Array((new BigNumber(2)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw90.length); _i0_14++) {
          _nw90[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _518_v23 = _nw90;
        let _521_v24;
        _521_v24 = _module.D0.create_DC3(_module.D0.create_DC1((_dafny.ZERO).minus(p0), p0));
        let _522_v25;
        _522_v25 = _dafny.Seq.of(_521_v24);
        let _index90 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
        (_518_v23)[_index90] = new BigNumber((_522_v25).length);
        let _523_v26;
        _523_v26 = _dafny.Seq.of(p2, (p2) || (true), true);
        let _index91 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
        let _rhs80 = (((p2) ? (p0) : (p0))).plus(p0);
        let _rhs81 = new BigNumber((_523_v26).length);
        let _lhs65 = globalState;
        let _lhs66 = _518_v23;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
        _lhs65.f3 = _rhs80;
        _lhs66[_lhs67] = _rhs81;
        (globalState).f4 = p2;
        let _524_v27;
        _524_v27 = _dafny.Set.fromElements(p2);
        let _525_v28;
        _525_v28 = _module.D9.create_DC20(_524_v27);
        let _526_v29;
        _526_v29 = _dafny.Seq.of(_525_v28);
        let _rhs82 = (_dafny.ZERO).minus((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]);
        let _rhs83 = _526_v29;
        let _lhs68 = globalState;
        _lhs68.f8 = _rhs82;
        _526_v29 = _rhs83;
        let _527_v30;
        _527_v30 = _dafny.Seq.of((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], p0, (_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], p0, p0);
        let _528_v31;
        _528_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p0)).length),p0);
        let _529_v32;
        _529_v32 = _module.D3.create_DC8(_528_v31);
        let _530_v33;
        _530_v33 = _dafny.MultiSet.fromElements(_523_v26);
        let _531_v34;
        let _init15 = function (_532_i8) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        };
        let _nw91 = Array((new BigNumber(9)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw91.length); _i0_15++) {
          _nw91[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _531_v34 = _nw91;
        let _533_v35;
        _533_v35 = _dafny.Map.Empty.slice().updateUnsafe(_531_v34,new BigNumber(-607));
        if ((_this).fm3(_527_v30, (p1).Union(_dafny.MultiSet.fromElements(p0, _module.__default.fm26(_529_v32, _module.__default.fm26(_529_v32, (_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], _dafny.MultiSet.FromArray(_dafny.Seq.of(_523_v26)), globalState), _530_v33, globalState), new BigNumber((_533_v35).length), p0, p0)), p2, globalState)) {
          (globalState).f4 = p2;
          let _534_v36;
          _534_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]);
          (globalState).f3 = (((_534_v36).contains(p2)) ? ((_534_v36).get(p2)) : (p0));
          _479_v0 = _479_v0;
          _479_v0 = _479_v0;
          let _535_v37;
          _535_v37 = _dafny.Set.fromElements(new BigNumber((p1).cardinality()));
          _535_v37 = (_dafny.Set.fromElements((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((function () {
            let _coll62 = new _dafny.Set();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(331), new BigNumber(658))) {
              let _536_v38 = _compr_62;
              if (((new BigNumber(331)).isLessThanOrEqualTo(_536_v38)) && ((_536_v38).isLessThan(new BigNumber(658)))) {
                _coll62.add((_536_v38).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_479_v0).length),p0)).length)));
              }
            }
            return _coll62;
          }()).length))).length), (_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))])).Union(_535_v37);
        } else {
          _521_v24 = _521_v24;
          let _537_v39;
          _537_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(469),_dafny.Seq.of(p0));
          let _538_v40;
          let _nw92 = Array((new BigNumber(9)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]);
          _nw92[(_dafny.ONE).toNumber()] = _527_v30;
          _nw92[(new BigNumber(2)).toNumber()] = _527_v30;
          _nw92[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_527_v30, _module.__default.safeIndex((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], new BigNumber((_527_v30).length)), new BigNumber(645));
          _nw92[(new BigNumber(4)).toNumber()] = _527_v30;
          _nw92[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((((_537_v39).contains(new BigNumber(182))) ? ((_537_v39).get(new BigNumber(182))) : (_527_v30)), _dafny.Seq.update(_527_v30, _module.__default.safeIndex(p0, new BigNumber((_527_v30).length)), new BigNumber((_479_v0).length)));
          _nw92[(new BigNumber(6)).toNumber()] = _527_v30;
          _nw92[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_527_v30, _527_v30);
          _nw92[(new BigNumber(8)).toNumber()] = _527_v30;
          _538_v40 = _nw92;
          let _index92 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_538_v40).length));
          (_538_v40)[_index92] = _527_v30;
          let _539_v41;
          _539_v41 = new _dafny.CodePoint('n'.codePointAt(0));
          let _540_v42;
          let _nw93 = Array((new BigNumber(4)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("eypc");
          _nw93[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_479_v0, _module.__default.safeIndex(p0, new BigNumber((_479_v0).length)), _539_v41), _module.__default.safeIndex((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], new BigNumber((_dafny.Seq.update(_479_v0, _module.__default.safeIndex(p0, new BigNumber((_479_v0).length)), _539_v41)).length)), _539_v41), _479_v0);
          _nw93[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_479_v0, _479_v0);
          _nw93[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_479_v0, _479_v0);
          _540_v42 = _nw93;
          let _541_v43;
          _541_v43 = _module.D2.create_DC7(_539_v41, p2, (_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]);
          let _index93 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_538_v40).length));
          let _index94 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
          let _rhs84 = _527_v30;
          let _rhs85 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), ((_dafny.ZERO).minus((_dafny.ZERO).minus((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]))).plus(p0));
          let _rhs86 = _540_v42;
          let _rhs87 = !(_528_v31).contains((_541_v43).dtor_cf11);
          let _lhs69 = _538_v40;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_538_v40).length));
          let _lhs71 = _518_v23;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
          let _lhs73 = globalState;
          _lhs69[_lhs70] = _rhs84;
          _lhs71[_lhs72] = _rhs85;
          _540_v42 = _rhs86;
          _lhs73.f4 = _rhs87;
          let _542_v44;
          let _out16;
          _out16 = (_this).m14(p2, (_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).equals(_dafny.Map.Empty.slice().updateUnsafe(true,p2)), p2, new BigNumber((_527_v30).length), globalState);
          _542_v44 = _out16;
          let _543_v45;
          _543_v45 = _dafny.Seq.of(_523_v26);
          let _544_v46;
          _544_v46 = _module.D5.create_DC14(_module.__default.fm26(_529_v32, (_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], _dafny.MultiSet.FromArray(_dafny.Seq.of(_523_v26)), globalState), _542_v44);
          let _545_v47;
          _545_v47 = _dafny.Map.Empty.slice().updateUnsafe(_544_v46,p0);
          let _546_v48;
          _546_v48 = _dafny.Map.Empty.slice().updateUnsafe(_543_v45,_545_v47);
          _546_v48 = (_546_v48).update(_543_v45, _545_v47);
          (globalState).f4 = !(new BigNumber(628)).isEqualTo(_module.__default.fm26(_529_v32, ((_538_v40)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_538_v40).length))])[_module.__default.safeIndex((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))], new BigNumber(((_538_v40)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_538_v40).length))]).length))], _530_v33, globalState));
        }
        let _547_v49;
        _547_v49 = _module.D0.create_DC0((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))]);
        let _index95 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
        let _rhs88 = p0;
        let _rhs89 = new BigNumber((_dafny.Seq.Concat(_479_v0, _479_v0)).length);
        let _rhs90 = p2;
        let _rhs91 = ((p2) ? (_module.D0.create_DC0((_518_v23)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length))])) : (_547_v49));
        let _rhs92 = _dafny.Seq.Concat(_527_v30, _dafny.Seq.Concat(_527_v30, _527_v30));
        let _lhs74 = _518_v23;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_518_v23).length));
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        _lhs74[_lhs75] = _rhs88;
        _lhs76.f8 = _rhs89;
        _lhs77.f4 = _rhs90;
        _547_v49 = _rhs91;
        _lhs78.f0 = _rhs92;
      }
      let _548_v50;
      _548_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      _548_v50 = (_548_v50).update(p0, ((p2) ? (p2) : (p2)));
      (globalState).f4 = p2;
      let _549_v51;
      _549_v51 = _dafny.Seq.of(new BigNumber(72), p0, p0, p0);
      let _550_v52;
      _550_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_549_v51, _module.__default.safeIndex(p0, new BigNumber((_549_v51).length)), p0),p2);
      let _551_v53;
      _551_v53 = _dafny.Map.Empty.slice().updateUnsafe(_550_v52,p2);
      if (((((_551_v53).contains(_550_v52)) ? ((_551_v53).get(_550_v52)) : (p2))) && (p2)) {
        let _552_v54;
        let _nw94 = new _module.C0();
        _nw94.__ctor();
        _552_v54 = _nw94;
        (globalState).f8 = p0;
        if ((_this).fm3(_549_v51, p1, p2, globalState)) {
          let _553_v55;
          let _nw95 = new _module.C0();
          _nw95.__ctor();
          _553_v55 = _nw95;
          let _rhs93 = _dafny.Seq.UnicodeFromString("meyetabe");
          let _rhs94 = (_dafny.ZERO).minus(p0);
          let _lhs79 = globalState;
          _479_v0 = _rhs93;
          _lhs79.f3 = _rhs94;
          let _554_v56;
          let _nw96 = new _module.C0();
          _nw96.__ctor();
          _554_v56 = _nw96;
          (globalState).f4 = (_this).fm3(_549_v51, p1, ((p2) ? (p2) : (p2)), globalState);
          let _555_v57;
          let _init16 = ((_556_p2) => function (_557_i9) {
            return _556_p2;
          })(p2);
          let _nw97 = Array((new BigNumber(3)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw97.length); _i0_16++) {
            _nw97[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _555_v57 = _nw97;
          let _558_v58;
          _558_v58 = _module.D5.create_DC14(p0, p2);
          let _index96 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_555_v57).length));
          (_555_v57)[_index96] = ((_558_v58).dtor_cf21) === (p2);
          let _index97 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_555_v57).length));
          (_555_v57)[_index97] = p2;
          let _559_v59;
          _559_v59 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24(globalState),_dafny.MultiSet.fromElements(p2, p2, p2, p2));
          let _560_v60;
          _560_v60 = _dafny.MultiSet.fromElements(p2);
          let _561_v61;
          _561_v61 = _dafny.Seq.of(_dafny.MultiSet.fromElements(p2, p2), _560_v60);
          let _562_v62;
          _562_v62 = _560_v60;
          let _563_v63;
          _563_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,_560_v60);
          let _564_v64;
          let _nw98 = Array((new BigNumber(27)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = ((_560_v60).update(p2, _module.__default.abs(p0))).Intersect(_module.__default.fm0(p0, globalState));
          _nw98[(_dafny.ONE).toNumber()] = (_561_v61)[_module.__default.safeIndex(p0, new BigNumber((_561_v61).length))];
          _nw98[(new BigNumber(2)).toNumber()] = (_560_v60).Difference(_dafny.MultiSet.fromElements(true, p2, p2, p2));
          _nw98[(new BigNumber(3)).toNumber()] = ((p2) ? (_560_v60) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(p2, !(p2)))));
          _nw98[(new BigNumber(4)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(5)).toNumber()] = (_562_v62);
          _nw98[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(p2);
          _nw98[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(p2, p2)).Difference(_560_v60);
          _nw98[(new BigNumber(8)).toNumber()] = (_560_v60).Intersect(_560_v60);
          _nw98[(new BigNumber(9)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(10)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(11)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(12)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(13)).toNumber()] = ((p2) ? ((_561_v61)[_module.__default.safeIndex(new BigNumber((_549_v51).length), new BigNumber((_561_v61).length))]) : (_560_v60));
          _nw98[(new BigNumber(14)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(15)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(16)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(17)).toNumber()] = (_560_v60).Intersect(_560_v60);
          _nw98[(new BigNumber(18)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(19)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(20)).toNumber()] = (_560_v60).update(p2, _module.__default.abs(p0));
          _nw98[(new BigNumber(21)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(22)).toNumber()] = (((_563_v63).contains(p0)) ? ((_563_v63).get(p0)) : (_560_v60));
          _nw98[(new BigNumber(23)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(24)).toNumber()] = _560_v60;
          _nw98[(new BigNumber(25)).toNumber()] = (_560_v60).Intersect(_560_v60);
          _nw98[(new BigNumber(26)).toNumber()] = (_562_v62);
          _564_v64 = _nw98;
          let _index98 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_564_v64).length));
          (_564_v64)[_index98] = _module.__default.fm0(new BigNumber((_479_v0).length), globalState);
          let _index99 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_555_v57).length));
          let _index100 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_555_v57).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_564_v64).length));
          let _rhs95 = _module.__default.fm27(true, globalState);
          let _rhs96 = p2;
          let _rhs97 = !(p2) || (p2);
          let _rhs98 = _559_v59;
          let _rhs99 = _560_v60;
          let _lhs80 = globalState;
          let _lhs81 = _555_v57;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_555_v57).length));
          let _lhs83 = _555_v57;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_555_v57).length));
          let _lhs85 = _564_v64;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_564_v64).length));
          _lhs80.f2 = _rhs95;
          _lhs81[_lhs82] = _rhs96;
          _lhs83[_lhs84] = _rhs97;
          _559_v59 = _rhs98;
          _lhs85[_lhs86] = _rhs99;
        } else {
          let _565_v65;
          let _init17 = ((_566_p0) => function (_567_i10) {
            return !(_566_p0).isEqualTo(_566_p0);
          })(p0);
          let _nw99 = Array((new BigNumber(3)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw99.length); _i0_17++) {
            _nw99[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _565_v65 = _nw99;
          let _568_v66;
          _568_v66 = new _dafny.CodePoint('h'.codePointAt(0));
          let _index102 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length));
          (_565_v65)[_index102] = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ijcbji"), _568_v66);
          let _index103 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length));
          (_565_v65)[_index103] = p2;
          let _569_v67;
          _569_v67 = _dafny.MultiSet.fromElements((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))], (_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))], p2, (_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))]);
          let _570_v68;
          _570_v68 = _module.D2.create_DC6(p0, (_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))], p2);
          let _571_v69;
          _571_v69 = _dafny.Seq.of(p2);
          let _572_v70;
          let _nw100 = Array((new BigNumber(20)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _569_v67;
          _nw100[(_dafny.ONE).toNumber()] = _569_v67;
          _nw100[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))])).Difference(_569_v67);
          _nw100[(new BigNumber(3)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.fromElements((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))])).Union(_dafny.MultiSet.fromElements((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))]));
          _nw100[(new BigNumber(5)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(p2, (_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))], false, !(p2));
          _nw100[(new BigNumber(7)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(8)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(9)).toNumber()] = (((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))]) ? (_569_v67) : (_569_v67));
          _nw100[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))]);
          _nw100[(new BigNumber(11)).toNumber()] = (_569_v67).Difference(_dafny.MultiSet.fromElements((_570_v68).dtor_cf7, p2, (_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))]));
          _nw100[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(p2);
          _nw100[(new BigNumber(13)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements((_module.D9.create_DC21(!(!((_565_v65)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_565_v65).length))])), _571_v69, !(p2))).dtor_cf28);
          _nw100[(new BigNumber(15)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(16)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.FromArray(_571_v69);
          _nw100[(new BigNumber(18)).toNumber()] = _569_v67;
          _nw100[(new BigNumber(19)).toNumber()] = _569_v67;
          _572_v70 = _nw100;
          let _index104 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_572_v70).length));
          (_572_v70)[_index104] = _569_v67;
          let _index105 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_572_v70).length));
          (_572_v70)[_index105] = (_569_v67).Union(_dafny.MultiSet.fromElements(p2));
          let _573_v71;
          let _nw101 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _573_v71 = _nw101;
          let _index106 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_573_v71).length));
          (_573_v71)[_index106] = p0;
          let _574_v72;
          _574_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-798),p0);
          let _index107 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_573_v71).length));
          (_573_v71)[_index107] = (new BigNumber(((_574_v72).Merge(_574_v72)).length)).plus(p0);
          (globalState).f8 = (p0).multipliedBy(new BigNumber(975));
          _479_v0 = _479_v0;
        }
        _552_v54 = _552_v54;
        let _575_v73;
        _575_v73 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_479_v0).length));
        (globalState).f3 = (new BigNumber(((_module.D12.create_DC26(_575_v73)).dtor_cf35).length)).plus(new BigNumber(((_548_v50).Merge(function () {
          let _coll63 = new _dafny.Map();
          for (const _compr_63 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(-919))) {
            let _576_v74 = _compr_63;
            if (((new BigNumber(958)).isLessThanOrEqualTo(_576_v74)) && ((_576_v74).isLessThan(new BigNumber(-919)))) {
              _coll63.push([(_576_v74).plus(p0),p2]);
            }
          }
          return _coll63;
        }())).length));
      } else {
        let _577_v75;
        _577_v75 = _dafny.Set.fromElements(p0, (_dafny.ZERO).minus(p0));
        if ((_577_v75).IsProperSubsetOf((_577_v75).Intersect(_module.__default.fm28(p0, globalState)))) {
          (globalState).f3 = (_dafny.ZERO).minus(p0);
          let _578_v76;
          _578_v76 = _dafny.Set.fromElements(p2, p2, p2);
          let _579_v77;
          _579_v77 = _dafny.MultiSet.fromElements(p2, true, p2, p2);
          let _580_v78;
          _580_v78 = _dafny.Seq.of(_549_v51);
          let _581_v79;
          _581_v79 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p2));
          let _582_v80;
          let _nw102 = Array((new BigNumber(16)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(p0)).isLessThan(p0);
          _nw102[(_dafny.ONE).toNumber()] = p2;
          _nw102[(new BigNumber(2)).toNumber()] = (_578_v76).IsSubsetOf(_578_v76);
          _nw102[(new BigNumber(3)).toNumber()] = p2;
          _nw102[(new BigNumber(4)).toNumber()] = p2;
          _nw102[(new BigNumber(5)).toNumber()] = (_579_v77).IsDisjointFrom(_579_v77);
          _nw102[(new BigNumber(6)).toNumber()] = p2;
          _nw102[(new BigNumber(7)).toNumber()] = true;
          _nw102[(new BigNumber(8)).toNumber()] = (_this).fm3(_549_v51, _dafny.MultiSet.FromArray((_580_v78)[_module.__default.safeIndex(p0, new BigNumber((_580_v78).length))]), p2, globalState);
          _nw102[(new BigNumber(9)).toNumber()] = p2;
          _nw102[(new BigNumber(10)).toNumber()] = p2;
          _nw102[(new BigNumber(11)).toNumber()] = (p1).IsDisjointFrom(p1);
          _nw102[(new BigNumber(12)).toNumber()] = (((_581_v79).contains(p2)) ? ((_581_v79).get(p2)) : (p2));
          _nw102[(new BigNumber(13)).toNumber()] = p2;
          _nw102[(new BigNumber(14)).toNumber()] = false;
          _nw102[(new BigNumber(15)).toNumber()] = !(!(p2) || (p2));
          _582_v80 = _nw102;
          let _index108 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_582_v80).length));
          (_582_v80)[_index108] = p2;
          let _index109 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_582_v80).length));
          (_582_v80)[_index109] = !(false);
          let _583_v81;
          let _nw103 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _583_v81 = _nw103;
          let _index110 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_583_v81).length));
          (_583_v81)[_index110] = _581_v79;
          let _index111 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_583_v81).length));
          (_583_v81)[_index111] = _581_v79;
          let _584_v82;
          let _nw104 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _584_v82 = _nw104;
          let _585_v83;
          _585_v83 = new _dafny.CodePoint('x'.codePointAt(0));
          let _index112 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_584_v82).length));
          (_584_v82)[_index112] = _585_v83;
          let _index113 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_584_v82).length));
          (_584_v82)[_index113] = _585_v83;
          (globalState).f4 = ((_579_v77).Union((_dafny.MultiSet.fromElements(p2)))).equals(_579_v77);
        } else {
          (globalState).f4 = !(_548_v50).contains(new BigNumber((_577_v75).length));
          let _586_v84;
          _586_v84 = _dafny.MultiSet.fromElements(!(p2) || (p2), p2);
          _586_v84 = _586_v84;
          let _587_v85;
          _587_v85 = new _dafny.CodePoint('v'.codePointAt(0));
          _587_v85 = _587_v85;
          (globalState).f3 = p0;
          let _588_v86;
          let _init18 = ((_589_p0) => function (_590_i11) {
            return _module.__default.safeDivisionInt(_590_i11, _589_p0);
          })(p0);
          let _nw105 = Array((new BigNumber(16)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
            _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _588_v86 = _nw105;
          let _index114 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_588_v86).length));
          (_588_v86)[_index114] = (new BigNumber(-359)).plus(p0);
          let _591_v87;
          _591_v87 = _dafny.Seq.of(!(!(true)), p2, p2, p2, p2);
          let _592_v88;
          _592_v88 = _dafny.MultiSet.fromElements(_591_v87, _591_v87);
          let _593_v89;
          _593_v89 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _594_v90;
          _594_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_593_v89).length),_591_v87);
          let _index115 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_588_v86).length));
          (_588_v86)[_index115] = (((_592_v88).contains((((_594_v90).contains(p0)) ? ((_594_v90).get(p0)) : (_591_v87)))) ? ((_592_v88).get((((_594_v90).contains(p0)) ? ((_594_v90).get(p0)) : (_591_v87)))) : (p0));
        }
        (globalState).f4 = p2;
        (globalState).f0 = _dafny.Seq.Concat(_549_v51, _549_v51);
        (globalState).f4 = (((p2) ? (p2) : (p2))) || (p2);
        let _595_v91;
        let _nw106 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _595_v91 = _nw106;
        let _index116 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_595_v91).length));
        (_595_v91)[_index116] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), function (_596_i12) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        });
        let _597_v92;
        let _init19 = ((_598_p0) => function (_599_i13) {
          return (_598_p0).isEqualTo(_598_p0);
        })(p0);
        let _nw107 = Array((new BigNumber(27)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw107.length); _i0_19++) {
          _nw107[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _597_v92 = _nw107;
        let _index117 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_597_v92).length));
        (_597_v92)[_index117] = true;
        let _index118 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_595_v91).length));
        let _index119 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_597_v92).length));
        let _rhs100 = _dafny.Seq.UnicodeFromString("miiuxq");
        let _rhs101 = _597_v92;
        let _rhs102 = false;
        let _lhs87 = _595_v91;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_595_v91).length));
        let _lhs89 = _597_v92;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_597_v92).length));
        _lhs87[_lhs88] = _rhs100;
        _597_v92 = _rhs101;
        _lhs89[_lhs90] = _rhs102;
      }
      let _600_i14;
      _600_i14 = _dafny.ZERO;
      L4: {
        while (_dafny.areEqual(_479_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_611_i15) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_600_i14)) {
              break L4;
            }
            _600_i14 = (_600_i14).plus(_dafny.ONE);
            let _601_v93;
            _601_v93 = _dafny.Seq.of(!(true));
            let _602_v94;
            _602_v94 = _dafny.MultiSet.fromElements((_601_v93)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_601_v93).length))], false);
            let _603_v95;
            _603_v95 = _602_v94;
            let _604_v96;
            let _nw108 = Array((new BigNumber(13)).toNumber());
            _nw108[(_dafny.ZERO).toNumber()] = _603_v95;
            _nw108[(_dafny.ONE).toNumber()] = _602_v94;
            _nw108[(new BigNumber(2)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(3)).toNumber()] = _module.__default.fm29(p0, p0, p2, p0, globalState);
            _nw108[(new BigNumber(4)).toNumber()] = ((!(p2)) ? (_603_v95) : (_603_v95));
            _nw108[(new BigNumber(5)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(6)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(7)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(8)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(9)).toNumber()] = _602_v94;
            _nw108[(new BigNumber(10)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(11)).toNumber()] = _603_v95;
            _nw108[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(p2);
            _604_v96 = _nw108;
            let _index120 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_604_v96).length));
            (_604_v96)[_index120] = _603_v95;
            let _index121 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_604_v96).length));
            (_604_v96)[_index121] = _603_v95;
            let _605_v97;
            _605_v97 = _dafny.Seq.of(p1);
            let _606_v98;
            let _nw109 = Array((new BigNumber(13)).toNumber());
            _nw109[(_dafny.ZERO).toNumber()] = p2;
            _nw109[(_dafny.ONE).toNumber()] = p2;
            _nw109[(new BigNumber(2)).toNumber()] = p2;
            _nw109[(new BigNumber(3)).toNumber()] = p2;
            _nw109[(new BigNumber(4)).toNumber()] = (_601_v93)[_module.__default.safeIndex(new BigNumber((_548_v50).length), new BigNumber((_601_v93).length))];
            _nw109[(new BigNumber(5)).toNumber()] = p2;
            _nw109[(new BigNumber(6)).toNumber()] = (p1).IsSubsetOf((_605_v97)[_module.__default.safeIndex(p0, new BigNumber((_605_v97).length))]);
            _nw109[(new BigNumber(7)).toNumber()] = p2;
            _nw109[(new BigNumber(8)).toNumber()] = p2;
            _nw109[(new BigNumber(9)).toNumber()] = p2;
            _nw109[(new BigNumber(10)).toNumber()] = p2;
            _nw109[(new BigNumber(11)).toNumber()] = p2;
            _nw109[(new BigNumber(12)).toNumber()] = true;
            _606_v98 = _nw109;
            let _rhs103 = p2;
            let _rhs104 = _549_v51;
            let _rhs105 = _606_v98;
            let _lhs91 = globalState;
            let _lhs92 = globalState;
            _lhs91.f4 = _rhs103;
            _lhs92.f0 = _rhs104;
            _606_v98 = _rhs105;
            let _607_v99;
            _607_v99 = _dafny.Map.Empty.slice().updateUnsafe(p2,_549_v51);
            let _608_v100;
            _608_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.update((((_607_v99).contains(p2)) ? ((_607_v99).get(p2)) : (_549_v51)), _module.__default.safeIndex(p0, new BigNumber(((((_607_v99).contains(p2)) ? ((_607_v99).get(p2)) : (_549_v51))).length)), new BigNumber(288)));
            (globalState).f0 = (((_608_v100).contains(p0)) ? ((_608_v100).get(p0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), ((_609_p0) => function (_610_i16) {
              return _609_p0;
            })(p0))));
            (globalState).f4 = ((p2) ? (p2) : (p2));
          }
        }
      }
      let _612_v101;
      let _nw110 = Array((new BigNumber(5)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = ((p2) ? ((_dafny.ZERO).minus(p0)) : (p0));
      _nw110[(_dafny.ONE).toNumber()] = p0;
      _nw110[(new BigNumber(2)).toNumber()] = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).update(p2, p2)).length)).plus(new BigNumber(720));
      _nw110[(new BigNumber(3)).toNumber()] = p0;
      _nw110[(new BigNumber(4)).toNumber()] = p0;
      _612_v101 = _nw110;
      _612_v101 = _612_v101;
      return;
    }
    m14(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f8 = (p3).minus(p3);
      let _hi3 = p3;
      for (let _613_i0 = p3; _613_i0.isLessThan(_hi3); _613_i0 = _613_i0.plus(_dafny.ONE)) {
        let _614_v0;
        let _nw111 = new _module.C0();
        _nw111.__ctor();
        _614_v0 = _nw111;
        (globalState).f4 = p0;
        (globalState).f4 = p2;
        (globalState).f3 = (_dafny.ZERO).minus(_613_i0);
      }
      if (p1) {
        let _615_v1;
        let _init20 = ((_616_p3) => function (_617_i1) {
          return _module.__default.safeDivisionInt(_617_i1, _616_p3);
        })(p3);
        let _nw112 = Array((_dafny.ONE).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw112.length); _i0_20++) {
          _nw112[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _615_v1 = _nw112;
        let _index122 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
        (_615_v1)[_index122] = new BigNumber(199);
        let _index123 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_615_v1).length));
        (_615_v1)[_index123] = p3;
        let _618_v2;
        _618_v2 = _dafny.Set.fromElements(p3);
        let _619_v3;
        _619_v3 = _dafny.Seq.of(_618_v2, _618_v2);
        let _620_v4;
        _620_v4 = _dafny.Seq.of(p0, !(p1));
        let _621_v5;
        _621_v5 = _dafny.Seq.UnicodeFromString("auprui");
        let _622_v6;
        _622_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_621_v5).length),p3);
        let _623_v7;
        _623_v7 = _module.D3.create_DC8(_622_v6);
        let _624_v8;
        _624_v8 = _dafny.Map.Empty.slice().updateUnsafe(_620_v4,_623_v7);
        let _625_v9;
        _625_v9 = _module.D4.create_DC10(_620_v4);
        let _626_v10;
        _626_v10 = _module.D4.create_DC12(_625_v9);
        let _627_v11;
        _627_v11 = _dafny.Map.Empty.slice().updateUnsafe(_626_v10,_618_v2);
        let _628_v12;
        _628_v12 = _module.D9.create_DC21(p1, _620_v4, !(p0));
        let _629_v13;
        _629_v13 = _dafny.Seq.of(p3, p3);
        let _630_v14;
        _630_v14 = _dafny.MultiSet.fromElements(p3, p3, p3, p3, p3);
        let _631_v15;
        _631_v15 = _dafny.Seq.of(_module.__default.fm31(p1, p1, p1, globalState));
        let _pat_let_tv20 = _621_v5;
        let _pat_let_tv21 = _628_v12;
        let _pat_let_tv22 = p1;
        let _pat_let_tv23 = _621_v5;
        let _pat_let_tv24 = globalState;
        let _pat_let_tv25 = _621_v5;
        let _pat_let_tv26 = _628_v12;
        let _pat_let_tv27 = p1;
        let _pat_let_tv28 = _621_v5;
        let _pat_let_tv29 = globalState;
        let _632_v16;
        let _nw113 = Array((new BigNumber(29)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = ((_619_v3)[_module.__default.safeIndex(new BigNumber((_624_v8).length), new BigNumber((_619_v3).length))]).IsSubsetOf((((_627_v11).contains(function (_pat_let19_0) {
          return function (_635_dt__update__tmp_h0) {
            return function (_pat_let20_0) {
              return function (_636_dt__update_hcf18_h0) {
                return _module.D4.create_DC12(_636_dt__update_hcf18_h0);
              }(_pat_let20_0);
            }(_module.__default.fm30(new BigNumber((_pat_let_tv25).length), _pat_let_tv26, _pat_let_tv27, new BigNumber((_pat_let_tv28).length), _pat_let_tv29));
          }(_pat_let19_0);
        }(_module.D4.create_DC12(_625_v9)))) ? ((_627_v11).get(function (_pat_let17_0) {
          return function (_633_dt__update__tmp_h1) {
            return function (_pat_let18_0) {
              return function (_634_dt__update_hcf18_h1) {
                return _module.D4.create_DC12(_634_dt__update_hcf18_h1);
              }(_pat_let18_0);
            }(_module.__default.fm30(new BigNumber((_pat_let_tv20).length), _pat_let_tv21, _pat_let_tv22, new BigNumber((_pat_let_tv23).length), _pat_let_tv24));
          }(_pat_let17_0);
        }(_module.D4.create_DC12(_625_v9)))) : (_618_v2)));
        _nw113[(_dafny.ONE).toNumber()] = p0;
        _nw113[(new BigNumber(2)).toNumber()] = p0;
        _nw113[(new BigNumber(3)).toNumber()] = p0;
        _nw113[(new BigNumber(4)).toNumber()] = p0;
        _nw113[(new BigNumber(5)).toNumber()] = p2;
        _nw113[(new BigNumber(6)).toNumber()] = !(p1);
        _nw113[(new BigNumber(7)).toNumber()] = p0;
        _nw113[(new BigNumber(8)).toNumber()] = p1;
        _nw113[(new BigNumber(9)).toNumber()] = p2;
        _nw113[(new BigNumber(10)).toNumber()] = !(p1);
        _nw113[(new BigNumber(11)).toNumber()] = !(false);
        _nw113[(new BigNumber(12)).toNumber()] = (_this).fm3(_629_v13, _630_v14, (_this).fm3(_629_v13, _630_v14, p2, globalState), globalState);
        _nw113[(new BigNumber(13)).toNumber()] = p1;
        _nw113[(new BigNumber(14)).toNumber()] = p1;
        _nw113[(new BigNumber(15)).toNumber()] = p2;
        _nw113[(new BigNumber(16)).toNumber()] = (p1) === (p1);
        _nw113[(new BigNumber(17)).toNumber()] = !(p1);
        _nw113[(new BigNumber(18)).toNumber()] = p1;
        _nw113[(new BigNumber(19)).toNumber()] = (p3).isEqualTo(new BigNumber((_631_v15).length));
        _nw113[(new BigNumber(20)).toNumber()] = p1;
        _nw113[(new BigNumber(21)).toNumber()] = p2;
        _nw113[(new BigNumber(22)).toNumber()] = p1;
        _nw113[(new BigNumber(23)).toNumber()] = (p1) && (p1);
        _nw113[(new BigNumber(24)).toNumber()] = p0;
        _nw113[(new BigNumber(25)).toNumber()] = !(p2);
        _nw113[(new BigNumber(26)).toNumber()] = p0;
        _nw113[(new BigNumber(27)).toNumber()] = _dafny.Seq.contains(_629_v13, p3);
        _nw113[(new BigNumber(28)).toNumber()] = p2;
        _632_v16 = _nw113;
        let _index124 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length));
        (_632_v16)[_index124] = (_620_v4)[_module.__default.safeIndex(p3, new BigNumber((_620_v4).length))];
        let _index125 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
        let _index126 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_615_v1).length));
        let _index127 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length));
        let _rhs106 = p3;
        let _rhs107 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(p3, p3)).length), p3);
        let _rhs108 = p3;
        let _rhs109 = !(p2);
        let _rhs110 = p1;
        let _lhs93 = globalState;
        let _lhs94 = _615_v1;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
        let _lhs96 = _615_v1;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_615_v1).length));
        let _lhs98 = _632_v16;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length));
        let _lhs100 = globalState;
        _lhs93.f8 = _rhs106;
        _lhs94[_lhs95] = _rhs107;
        _lhs96[_lhs97] = _rhs108;
        _lhs98[_lhs99] = _rhs109;
        _lhs100.f4 = _rhs110;
        let _637_v17;
        _637_v17 = _dafny.Seq.of(_621_v5, _621_v5, _621_v5, _621_v5, _621_v5);
        let _638_v18;
        _638_v18 = new _dafny.CodePoint('y'.codePointAt(0));
        let _639_v19;
        _639_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(!(p0),(_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))]));
        let _index128 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length));
        (_632_v16)[_index128] = !(_module.__default.fm32(p3, _637_v17, _638_v18, p0, globalState)).equals(_639_v19);
        let _640_v20;
        _640_v20 = _dafny.Set.fromElements(p2);
        if (((_640_v20).IsProperSubsetOf(_640_v20)) || ((_640_v20).contains(p2))) {
          _621_v5 = _dafny.Seq.UnicodeFromString("mosxmc");
          let _index129 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
          (_615_v1)[_index129] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p3, p3)), new BigNumber(837)));
          let _641_v21;
          _641_v21 = _dafny.Map.Empty.slice().updateUnsafe(_632_v16,_638_v18);
          (globalState).f8 = new BigNumber((_641_v21).length);
          let _642_v22;
          _642_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_637_v17).length),p1);
          let _643_v25;
          _643_v25 = _dafny.Set.fromElements(_618_v2, function () {
            let _coll64 = new _dafny.Set();
            for (const _compr_64 of (_642_v22).Keys.Elements) {
              let _644_v23 = _compr_64;
              if ((_642_v22).contains(_644_v23)) {
                _coll64.add(_module.__default.safeModuloInt(_644_v23, new BigNumber((_dafny.Set.fromElements(false, !(false), true, false, !(true))).length)));
              }
            }
            return _coll64;
          }(), function () {
            let _coll65 = new _dafny.Set();
            for (const _compr_65 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(451))) {
              let _645_v24 = _compr_65;
              if (((new BigNumber(280)).isLessThanOrEqualTo(_645_v24)) && ((_645_v24).isLessThan(new BigNumber(451)))) {
                _coll65.add((_645_v24).multipliedBy(p3));
              }
            }
            return _coll65;
          }());
          _643_v25 = _643_v25;
          let _646_v26;
          let _nw114 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _646_v26 = _nw114;
          let _index130 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_646_v26).length));
          (_646_v26)[_index130] = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_629_v13, _630_v14, !((_632_v16)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length))]), globalState),false);
          let _index131 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_646_v26).length));
          (_646_v26)[_index131] = _dafny.Map.Empty.slice().updateUnsafe(!(_622_v6).contains(new BigNumber(340)),p2);
        } else {
          let _index132 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
          (_615_v1)[_index132] = p3;
          let _647_v27;
          _647_v27 = _dafny.MultiSet.fromElements(_620_v4);
          let _648_v28;
          _648_v28 = _dafny.Seq.of(_647_v27, _647_v27, _647_v27);
          let _649_v29;
          _649_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm26(_623_v7, new BigNumber(390), (_648_v28)[_module.__default.safeIndex((_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))], new BigNumber((_648_v28).length))], globalState),p1);
          (globalState).f3 = ((_dafny.ZERO).minus(new BigNumber((_649_v29).length))).minus(new BigNumber(126));
          let _650_v30;
          let _nw115 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _650_v30 = _nw115;
          let _651_v31;
          _651_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, (_632_v16)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_632_v16).length))]),p1);
          let _index133 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_650_v30).length));
          (_650_v30)[_index133] = _651_v31;
          let _index134 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_650_v30).length));
          let _rhs111 = new BigNumber(393);
          let _rhs112 = (_618_v2).Difference(_dafny.Set.fromElements(new BigNumber((_621_v5).length)));
          let _rhs113 = (_module.__default.fm33(p1, p0, globalState)).Merge((function () {
            let _coll66 = new _dafny.Map();
            for (const _compr_66 of (_631_v15).Elements) {
              let _652_v32 = _compr_66;
              if (_dafny.Seq.contains(_631_v15, _652_v32)) {
                _coll66.push([_652_v32,p1]);
              }
            }
            return _coll66;
          }()).Merge(_651_v31));
          let _rhs114 = _638_v18;
          let _lhs101 = globalState;
          let _lhs102 = _650_v30;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_650_v30).length));
          _lhs101.f3 = _rhs111;
          _618_v2 = _rhs112;
          _lhs102[_lhs103] = _rhs113;
          _638_v18 = _rhs114;
          let _653_v33;
          _653_v33 = _dafny.Map.Empty.slice().updateUnsafe(_632_v16,(_dafny.ZERO).minus((_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))]));
          _653_v33 = (_653_v33).update(_632_v16, _module.__default.safeModuloInt((_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))], p3));
          let _654_v34;
          let _nw116 = new _module.C0();
          _nw116.__ctor();
          _654_v34 = _nw116;
        }
        let _655_v35;
        _655_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_640_v20).length),_638_v18);
        let _index135 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length));
        (_615_v1)[_index135] = _module.__default.safeDivisionInt(new BigNumber((_655_v35).length), (_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))]);
        let _656_v37;
        _656_v37 = _dafny.MultiSet.fromElements(p3, new BigNumber((_621_v5).length));
        let _657_v38;
        _657_v38 = _dafny.MultiSet.fromElements(_620_v4);
        let _pat_let_tv30 = p2;
        let _pat_let_tv31 = globalState;
        let _pat_let_tv32 = p2;
        let _pat_let_tv33 = globalState;
        let _658_v39;
        _658_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm26(_623_v7, new BigNumber((function () {
          let _coll67 = new _dafny.Map();
          for (const _compr_67 of (_dafny.Seq.of(_656_v37, _630_v14, function (_pat_let21_0) {
            return function (_659_dt__update__tmp_h2) {
              return function (_pat_let22_0) {
                return function (_660_dt__update_hcf22_h0) {
                  return _660_dt__update_hcf22_h0;
                }(_pat_let22_0);
              }(_module.__default.fm27(_pat_let_tv30, _pat_let_tv31));
            }(_pat_let21_0);
          }(_656_v37))).Elements) {
            let _661_v36 = _compr_67;
            if (_dafny.Seq.contains(_dafny.Seq.of(_656_v37, _630_v14, function (_pat_let23_0) {
              return function (_662_dt__update__tmp_h2) {
                return function (_pat_let24_0) {
                  return function (_663_dt__update_hcf22_h0) {
                    return _663_dt__update_hcf22_h0;
                  }(_pat_let24_0);
                }(_module.__default.fm27(_pat_let_tv32, _pat_let_tv33));
              }(_pat_let23_0);
            }(_656_v37)), _661_v36)) {
              _coll67.push([_661_v36,true]);
            }
          }
          return _coll67;
        }()).length), _657_v38, globalState));
        _658_v39 = (_658_v39).update(p2, _module.__default.safeModuloInt(p3, (_615_v1)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_615_v1).length))]));
      } else {
        let _664_v40;
        _664_v40 = _dafny.Seq.of(p3);
        let _665_v41;
        _665_v41 = _dafny.MultiSet.fromElements(p3, new BigNumber(492), p3, p3, p3);
        r0 = (_this).fm3(_664_v40, _665_v41, p1, globalState);
        let _666_v42;
        _666_v42 = _dafny.Set.fromElements(p0, p2, p1, (p1) && (p1));
        let _667_v43;
        let _nw117 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
        _667_v43 = _nw117;
        let _668_v44;
        _668_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _rhs115 = (_666_v42).Intersect(_666_v42);
        let _rhs116 = _667_v43;
        let _rhs117 = (((_665_v41).contains(_module.__default.safeModuloInt(new BigNumber((_668_v44).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(p3))))) ? ((_665_v41).get(_module.__default.safeModuloInt(new BigNumber((_668_v44).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(p3))))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,p3)).length)));
        let _lhs104 = globalState;
        _666_v42 = _rhs115;
        _667_v43 = _rhs116;
        _lhs104.f3 = _rhs117;
        (globalState).f8 = p3;
        let _669_v45;
        let _nw118 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _669_v45 = _nw118;
        let _index136 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_669_v45).length));
        (_669_v45)[_index136] = new _dafny.CodePoint('h'.codePointAt(0));
        let _670_v46;
        _670_v46 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index137 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_669_v45).length));
        (_669_v45)[_index137] = _670_v46;
        (globalState).f8 = p3;
      }
      r0 = p0;
      let _671_v47;
      _671_v47 = _dafny.Seq.UnicodeFromString("mjj");
      let _672_v48;
      let _nw119 = Array((new BigNumber(13)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = false;
      _nw119[(_dafny.ONE).toNumber()] = p1;
      _nw119[(new BigNumber(2)).toNumber()] = !(p0);
      _nw119[(new BigNumber(3)).toNumber()] = p0;
      _nw119[(new BigNumber(4)).toNumber()] = !_dafny.Seq.contains(_671_v47, new _dafny.CodePoint('x'.codePointAt(0)));
      _nw119[(new BigNumber(5)).toNumber()] = p1;
      _nw119[(new BigNumber(6)).toNumber()] = p2;
      _nw119[(new BigNumber(7)).toNumber()] = (p3).isLessThan((_module.__default.fm34(p0, p1, globalState)).dtor_cf0);
      _nw119[(new BigNumber(8)).toNumber()] = p2;
      _nw119[(new BigNumber(9)).toNumber()] = p1;
      _nw119[(new BigNumber(10)).toNumber()] = p0;
      _nw119[(new BigNumber(11)).toNumber()] = p0;
      _nw119[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("vvo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_673_i2) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }));
      _672_v48 = _nw119;
      let _index138 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_672_v48).length));
      (_672_v48)[_index138] = p1;
      let _index139 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_672_v48).length));
      (_672_v48)[_index139] = p2;
      let _674_v49;
      _674_v49 = _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_671_v47, _671_v47), (_672_v48)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_672_v48).length))], (_672_v48)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_672_v48).length))], p2);
      let _675_v50;
      _675_v50 = new _dafny.CodePoint('f'.codePointAt(0));
      let _676_v51;
      _676_v51 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p3));
      let _677_v52;
      _677_v52 = _dafny.Seq.of(new BigNumber((_676_v51).cardinality()));
      let _678_v53;
      _678_v53 = _module.D2.create_DC7(_675_v50, p1, new BigNumber((_677_v52).length));
      let _679_v54;
      _679_v54 = _module.D12.create_DC27(p3, p3, p3, p3);
      let _pat_let_tv34 = _679_v54;
      let _index140 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_672_v48).length));
      (_672_v48)[_index140] = (_674_v49)[_module.__default.safeIndex((function (_pat_let25_0) {
        return function (_682_dt__update__tmp_h4) {
          return function (_pat_let28_0) {
            return function (_683_dt__update_hcf10_h0) {
              return _module.D2.create_DC7((_682_dt__update__tmp_h4).dtor_cf9, _683_dt__update_hcf10_h0, (_682_dt__update__tmp_h4).dtor_cf11);
            }(_pat_let28_0);
          }(true);
        }(_pat_let25_0);
      }(function (_pat_let26_0) {
        return function (_680_dt__update__tmp_h3) {
          return function (_pat_let27_0) {
            return function (_681_dt__update_hcf11_h0) {
              return _module.D2.create_DC7((_680_dt__update__tmp_h3).dtor_cf9, (_680_dt__update__tmp_h3).dtor_cf10, _681_dt__update_hcf11_h0);
            }(_pat_let27_0);
          }((_pat_let_tv34).dtor_cf37);
        }(_pat_let26_0);
      }(_678_v53))).dtor_cf11, new BigNumber((_674_v49).length))];
      r0 = p1;
      return r0;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return ((function () {
        let _coll68 = new _dafny.Map();
        for (const _compr_68 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(920),false)).Keys.Elements) {
          let _684_v0 = _compr_68;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(920),false)).contains(_684_v0)) {
            _coll68.push([(_684_v0).minus(new BigNumber(851)),new BigNumber((_dafny.Seq.UnicodeFromString("jf")).length)]);
          }
        }
        return _coll68;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(441),new BigNumber(596)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-204),new BigNumber((_dafny.Seq.UnicodeFromString("ee")).length))).length),new BigNumber(986)));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      let _source13 = _module.D4.create_DC11(true, true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rqonprou")).length)), _dafny.Seq.UnicodeFromString("svtwbn"));
      if (_source13.is_DC11) {
        let _685___mcc_h0 = (_source13).cf14;
        let _686___mcc_h1 = (_source13).cf15;
        let _687___mcc_h2 = (_source13).cf16;
        let _688___mcc_h3 = (_source13).cf17;
        let _689_cf17 = _688___mcc_h3;
        let _690_cf16 = _687___mcc_h2;
        let _691_cf15 = _686___mcc_h1;
        let _692_cf14 = _685___mcc_h0;
        return false;
      } else if (_source13.is_DC10) {
        let _693___mcc_h4 = (_source13).cf13;
        let _694_cf13 = _693___mcc_h4;
        return !(new BigNumber(378)).isEqualTo(new BigNumber(708));
      } else {
        let _695___mcc_h5 = (_source13).cf18;
        let _696_cf18 = _695___mcc_h5;
        return false;
      }
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _697_v0;
      _697_v0 = new BigNumber(558);
      (globalState).f3 = _697_v0;
      let _698_v2;
      _698_v2 = _dafny.MultiSet.fromElements(_697_v0);
      if ((_dafny.MultiSet.fromElements(_697_v0, _697_v0, _697_v0)).IsProperSubsetOf((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll69 = new _dafny.Map();
        for (const _compr_69 of _dafny.IntegerRange(new BigNumber(764), new BigNumber(518))) {
          let _699_v1 = _compr_69;
          if (((new BigNumber(764)).isLessThanOrEqualTo(_699_v1)) && ((_699_v1).isLessThan(new BigNumber(518)))) {
            _coll69.push([(_699_v1).minus(_697_v0),new _dafny.CodePoint('u'.codePointAt(0))]);
          }
        }
        return _coll69;
      }()).length))).Union(_698_v2))) {
        let _700_v3;
        _700_v3 = _dafny.Seq.UnicodeFromString("xxhq");
        let _701_v4;
        _701_v4 = _dafny.Seq.of(_700_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-395)), function (_702_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }));
        let _703_v5;
        _703_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ilkioe"),(_701_v4)[_module.__default.safeIndex(new BigNumber(746), new BigNumber((_701_v4).length))]);
        let _704_v6;
        _704_v6 = true;
        _700_v3 = (((_703_v5).contains(_700_v3)) ? ((_703_v5).get(_700_v3)) : (_module.__default.fm17(_module.__default.fm18(_704_v6, _704_v6, globalState), globalState)));
        let _705_v7;
        _705_v7 = _dafny.Seq.of(_697_v0, _697_v0);
        let _706_v8;
        _706_v8 = _dafny.Seq.of((_this).fm3(_705_v7, _698_v2, _704_v6, globalState));
        let _707_v9;
        _707_v9 = _module.D2.create_DC6((((_706_v8)[_module.__default.safeIndex(new BigNumber((_700_v3).length), new BigNumber((_706_v8).length))]) ? (_697_v0) : (new BigNumber(446))), _704_v6, _704_v6);
        _707_v9 = _707_v9;
        let _708_v10;
        _708_v10 = _dafny.Set.fromElements(_704_v6);
        (globalState).f8 = _module.__default.safeDivisionInt((_697_v0).plus(_697_v0), new BigNumber((_708_v10).length));
        if (_704_v6) {
          let _709_v11;
          let _nw120 = new _module.C0();
          _nw120.__ctor();
          _709_v11 = _nw120;
          let _710_v12;
          let _nw121 = new _module.C0();
          _nw121.__ctor();
          _710_v12 = _nw121;
          let _711_v13;
          let _nw122 = Array((new BigNumber(14)).toNumber()).fill(false);
          _711_v13 = _nw122;
          let _index141 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_711_v13).length));
          (_711_v13)[_index141] = _704_v6;
          let _712_v14;
          _712_v14 = new _dafny.CodePoint('j'.codePointAt(0));
          let _713_v15;
          let _nw123 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
          _713_v15 = _nw123;
          let _714_v16;
          _714_v16 = _dafny.Seq.of(_711_v13, _711_v13);
          let _index142 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_713_v15).length));
          (_713_v15)[_index142] = _714_v16;
          let _index143 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_711_v13).length));
          let _index144 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_713_v15).length));
          let _rhs118 = _697_v0;
          let _rhs119 = new BigNumber(635);
          let _rhs120 = _704_v6;
          let _rhs121 = _712_v14;
          let _rhs122 = _714_v16;
          let _lhs105 = globalState;
          let _lhs106 = globalState;
          let _lhs107 = _711_v13;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_711_v13).length));
          let _lhs109 = _713_v15;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_713_v15).length));
          _lhs105.f8 = _rhs118;
          _lhs106.f3 = _rhs119;
          _lhs107[_lhs108] = _rhs120;
          _712_v14 = _rhs121;
          _lhs109[_lhs110] = _rhs122;
          let _index145 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_711_v13).length));
          (_711_v13)[_index145] = (_711_v13)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_711_v13).length))];
          let _715_v17;
          let _nw124 = new _module.C0();
          _nw124.__ctor();
          _715_v17 = _nw124;
        } else {
          let _716_v18;
          _716_v18 = new _dafny.CodePoint('a'.codePointAt(0));
          _700_v3 = _dafny.Seq.update(_dafny.Seq.update(_700_v3, _module.__default.safeIndex(_697_v0, new BigNumber((_700_v3).length)), _716_v18), _module.__default.safeIndex(_module.__default.fm20(!(false), globalState), new BigNumber((_dafny.Seq.update(_700_v3, _module.__default.safeIndex(_697_v0, new BigNumber((_700_v3).length)), _716_v18)).length)), _716_v18);
          let _717_v19;
          _717_v19 = _dafny.Map.Empty.slice().updateUnsafe(true,_704_v6);
          let _718_v20;
          _718_v20 = _dafny.Map.Empty.slice().updateUnsafe(_704_v6,_dafny.Map.Empty.slice().updateUnsafe(_704_v6,_704_v6));
          let _719_v21;
          _719_v21 = _module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_704_v6,true));
          let _720_v22;
          _720_v22 = _dafny.Seq.of(_717_v19);
          let _721_v23;
          let _nw125 = Array((new BigNumber(14)).toNumber());
          _nw125[(_dafny.ZERO).toNumber()] = _717_v19;
          _nw125[(_dafny.ONE).toNumber()] = _717_v19;
          _nw125[(new BigNumber(2)).toNumber()] = (((_718_v20).contains(_704_v6)) ? ((_718_v20).get(_704_v6)) : (_dafny.Map.Empty.slice().updateUnsafe(_704_v6,_704_v6)));
          _nw125[(new BigNumber(3)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_704_v6,_704_v6);
          _nw125[(new BigNumber(5)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(6)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(7)).toNumber()] = (_719_v21).dtor_cf19;
          _nw125[(new BigNumber(8)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(9)).toNumber()] = _module.__default.fm21(_704_v6, globalState);
          _nw125[(new BigNumber(10)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(11)).toNumber()] = (_720_v22)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_720_v22).length))];
          _nw125[(new BigNumber(12)).toNumber()] = _717_v19;
          _nw125[(new BigNumber(13)).toNumber()] = _717_v19;
          _721_v23 = _nw125;
          let _722_v24;
          _722_v24 = _dafny.Map.Empty.slice().updateUnsafe(_721_v23,_704_v6);
          _722_v24 = (_722_v24).update(_721_v23, (_697_v0).isEqualTo(new BigNumber((_700_v3).length)));
          _697_v0 = _697_v0;
          (globalState).f3 = _697_v0;
          _708_v10 = _708_v10;
        }
        r0 = !(!((new BigNumber((_700_v3).length)).isLessThanOrEqualTo(_697_v0)));
      } else {
        let _723_v25;
        _723_v25 = false;
        (globalState).f3 = ((_723_v25) ? (_697_v0) : (_697_v0));
        let _724_v26;
        let _nw126 = Array((new BigNumber(20)).toNumber()).fill(false);
        _724_v26 = _nw126;
        let _index146 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length));
        (_724_v26)[_index146] = _723_v25;
        let _index147 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length));
        (_724_v26)[_index147] = _723_v25;
        let _725_v27;
        _725_v27 = _dafny.Seq.UnicodeFromString("ijeb");
        let _726_v28;
        _726_v28 = _dafny.Seq.of(new BigNumber(-273));
        let _727_v29;
        _727_v29 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_725_v27).length), _697_v0), (((_724_v26)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length))]) ? (_726_v28) : (_726_v28)));
        (globalState).f0 = (_727_v29)[_module.__default.safeIndex(_697_v0, new BigNumber((_727_v29).length))];
        _697_v0 = _697_v0;
        let _728_v30;
        let _nw127 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
        _728_v30 = _nw127;
        let _729_v31;
        _729_v31 = _dafny.Map.Empty.slice().updateUnsafe(_697_v0,_697_v0);
        let _730_v32;
        let _nw128 = Array((new BigNumber(11)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = _697_v0;
        _nw128[(_dafny.ONE).toNumber()] = _697_v0;
        _nw128[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_724_v26)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length))], _723_v25), _dafny.Seq.of((_724_v26)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length))]))).length);
        _nw128[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_697_v0, _module.__default.fm20((_724_v26)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length))], globalState));
        _nw128[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(_697_v0, _697_v0);
        _nw128[(new BigNumber(5)).toNumber()] = _697_v0;
        _nw128[(new BigNumber(6)).toNumber()] = (new BigNumber((_729_v31).length)).plus(new BigNumber(523));
        _nw128[(new BigNumber(7)).toNumber()] = _697_v0;
        _nw128[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_697_v0, _697_v0));
        _nw128[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(_697_v0, _697_v0);
        _nw128[(new BigNumber(10)).toNumber()] = _697_v0;
        _730_v32 = _nw128;
        let _index148 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_730_v32).length));
        (_730_v32)[_index148] = _module.__default.fm20((_724_v26)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_724_v26).length))], globalState);
        let _index149 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_730_v32).length));
        let _rhs123 = _728_v30;
        let _rhs124 = _module.__default.safeDivisionInt(_697_v0, _697_v0);
        let _rhs125 = _697_v0;
        let _rhs126 = _723_v25;
        let _rhs127 = (((_697_v0).isLessThan(_697_v0)) ? (_724_v26) : (_724_v26));
        let _lhs111 = globalState;
        let _lhs112 = _730_v32;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_730_v32).length));
        _728_v30 = _rhs123;
        _lhs111.f3 = _rhs124;
        _lhs112[_lhs113] = _rhs125;
        _723_v25 = _rhs126;
        _724_v26 = _rhs127;
      }
      let _731_v33;
      _731_v33 = false;
      let _732_v34;
      _732_v34 = _698_v2;
      let _733_v35;
      _733_v35 = _dafny.Map.Empty.slice().updateUnsafe(_697_v0,_dafny.MultiSet.fromElements(_697_v0, _697_v0));
      let _734_v36;
      _734_v36 = _dafny.Map.Empty.slice().updateUnsafe(false,_731_v33);
      let _735_v37;
      _735_v37 = _dafny.Seq.UnicodeFromString("xebnstm");
      let _736_v38;
      _736_v38 = _module.D4.create_DC11(_731_v33, !(true), _697_v0, _735_v37);
      let _737_v39;
      _737_v39 = _dafny.Seq.of(!(_731_v33));
      let _738_v40;
      let _nw129 = Array((new BigNumber(25)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = !(_731_v33);
      _nw129[(_dafny.ONE).toNumber()] = _731_v33;
      _nw129[(new BigNumber(2)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(3)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(4)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(5)).toNumber()] = !(true);
      _nw129[(new BigNumber(6)).toNumber()] = ((_732_v34)).IsDisjointFrom((((_733_v35).contains(new BigNumber(679))) ? ((_733_v35).get(new BigNumber(679))) : ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-318)), function (_739_i1) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_740_i2) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length);
      }))).update(new BigNumber((_734_v36).length), _module.__default.abs(_697_v0)))));
      _nw129[(new BigNumber(7)).toNumber()] = false;
      _nw129[(new BigNumber(8)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(9)).toNumber()] = (_736_v38).dtor_cf14;
      _nw129[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_731_v33, _731_v33), _737_v39);
      _nw129[(new BigNumber(11)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(12)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(13)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(14)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(15)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(16)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(17)).toNumber()] = !(_731_v33);
      _nw129[(new BigNumber(18)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(19)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(20)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(21)).toNumber()] = (_697_v0).isEqualTo(_697_v0);
      _nw129[(new BigNumber(22)).toNumber()] = true;
      _nw129[(new BigNumber(23)).toNumber()] = _731_v33;
      _nw129[(new BigNumber(24)).toNumber()] = _731_v33;
      _738_v40 = _nw129;
      let _index150 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length));
      (_738_v40)[_index150] = false;
      let _741_v41;
      _741_v41 = _dafny.Seq.of(_697_v0);
      let _index151 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length));
      (_738_v40)[_index151] = (_697_v0).isLessThanOrEqualTo((new BigNumber((_741_v41).length)).minus(_697_v0));
      let _742_v42;
      let _init21 = ((_743_v0) => function (_744_i3) {
        return _module.__default.safeDivisionInt(_744_i3, (_dafny.ZERO).minus(_743_v0));
      })(_697_v0);
      let _nw130 = Array((new BigNumber(19)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw130.length); _i0_21++) {
        _nw130[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _742_v42 = _nw130;
      let _index152 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length));
      (_742_v42)[_index152] = new BigNumber((_dafny.Seq.UnicodeFromString("l")).length);
      let _index153 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length));
      (_742_v42)[_index153] = _697_v0;
      let _745_i4;
      _745_i4 = _dafny.ZERO;
      L5: {
        while ((_738_v40)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length))]) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_745_i4)) {
              break L5;
            }
            _745_i4 = (_745_i4).plus(_dafny.ONE);
            r0 = _731_v33;
            (globalState).f8 = (_697_v0).minus((_742_v42)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length))]);
            let _746_v43;
            _746_v43 = _dafny.Map.Empty.slice().updateUnsafe(_742_v42,_741_v41);
            let _747_v44;
            _747_v44 = _dafny.MultiSet.fromElements((_738_v40)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length))], _731_v33, _731_v33);
            _741_v41 = (((_746_v43).contains(_742_v42)) ? ((_746_v43).get(_742_v42)) : (_dafny.Seq.update(_dafny.Seq.Concat(_741_v41, _741_v41), _module.__default.safeIndex((((_747_v44).contains((_738_v40)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length))])) ? ((_747_v44).get((_738_v40)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_738_v40).length))])) : (new BigNumber(690))), new BigNumber((_dafny.Seq.Concat(_741_v41, _741_v41)).length)), (_742_v42)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length))])));
            let _index154 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length));
            (_742_v42)[_index154] = ((_742_v42)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length))]).multipliedBy((_742_v42)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length))]);
          }
        }
      }
      let _748_v45;
      _748_v45 = new _dafny.CodePoint('h'.codePointAt(0));
      let _749_v46;
      _749_v46 = _dafny.Map.Empty.slice().updateUnsafe(_697_v0,_742_v42);
      let _index155 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length));
      let _rhs128 = (((_749_v46).contains(_697_v0)) ? ((_749_v46).get(_697_v0)) : (_742_v42));
      let _rhs129 = new _dafny.CodePoint('b'.codePointAt(0));
      let _rhs130 = _748_v45;
      let _rhs131 = ((_742_v42)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length))]).multipliedBy(_697_v0);
      let _lhs114 = _742_v42;
      let _lhs115 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_742_v42).length));
      _742_v42 = _rhs128;
      _748_v45 = _rhs129;
      _748_v45 = _rhs130;
      _lhs114[_lhs115] = _rhs131;
      r0 = _731_v33;
      let _pat_let_tv35 = _748_v45;
      let _pat_let_tv36 = _731_v33;
      let _pat_let_tv37 = _738_v40;
      let _pat_let_tv38 = _738_v40;
      r1 = function (_source14) {
        if (_source14.is_DC11) {
          let _750___mcc_h0 = (_source14).cf14;
          let _751___mcc_h1 = (_source14).cf15;
          let _752___mcc_h2 = (_source14).cf16;
          let _753___mcc_h3 = (_source14).cf17;
          let _754_cf17 = _753___mcc_h3;
          let _755_cf16 = _752___mcc_h2;
          let _756_cf15 = _751___mcc_h1;
          let _757_cf14 = _750___mcc_h0;
          return _dafny.Seq.IsProperPrefixOf(_754_cf17, _dafny.Seq.update(_754_cf17, _module.__default.safeIndex(_755_cf16, new BigNumber((_754_cf17).length)), _pat_let_tv35));
        } else if (_source14.is_DC10) {
          let _758___mcc_h4 = (_source14).cf13;
          let _759_cf13 = _758___mcc_h4;
          return _pat_let_tv36;
        } else {
          let _760___mcc_h5 = (_source14).cf18;
          let _761_cf18 = _760___mcc_h5;
          return (_pat_let_tv38)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_pat_let_tv37).length))];
        }
      }(_736_v38);
      return [r0, r1];
    }
    m11(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _762_v0;
      _762_v0 = true;
      (globalState).f4 = !(_762_v0);
      let _763_v1;
      _763_v1 = new BigNumber(159);
      let _rhs132 = _763_v1;
      let _rhs133 = _763_v1;
      let _lhs116 = globalState;
      _lhs116.f8 = _rhs132;
      r2 = _rhs133;
      if (_762_v0) {
        let _764_v2;
        _764_v2 = _dafny.Set.fromElements(_762_v0);
        _764_v2 = (_764_v2).Difference(((!(_762_v0)) ? (_dafny.Set.fromElements(_762_v0)) : (_764_v2)));
        let _765_v3;
        _765_v3 = _dafny.Seq.UnicodeFromString("cooouputi");
        let _766_v4;
        _766_v4 = _dafny.MultiSet.fromElements(_763_v1);
        let _767_v5;
        _767_v5 = _dafny.Map.Empty.slice().updateUnsafe(_762_v0,new BigNumber((_766_v4).cardinality()));
        let _768_v6;
        _768_v6 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ligquqw"), _765_v3)).length), (_763_v1).multipliedBy((((_766_v4).contains(_763_v1)) ? ((_766_v4).get(_763_v1)) : (new BigNumber((_767_v5).length)))), _module.__default.safeDivisionInt(_763_v1, _763_v1));
        (globalState).f2 = _766_v4;
        let _769_v7;
        _769_v7 = new _dafny.CodePoint('n'.codePointAt(0));
        let _770_v8;
        let _nw131 = Array((new BigNumber(9)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _769_v7;
        _nw131[(_dafny.ONE).toNumber()] = ((_762_v0) ? (_769_v7) : (_769_v7));
        _nw131[(new BigNumber(2)).toNumber()] = _769_v7;
        _nw131[(new BigNumber(3)).toNumber()] = _769_v7;
        _nw131[(new BigNumber(4)).toNumber()] = _769_v7;
        _nw131[(new BigNumber(5)).toNumber()] = _769_v7;
        _nw131[(new BigNumber(6)).toNumber()] = _module.__default.fm18(_762_v0, _762_v0, globalState);
        _nw131[(new BigNumber(7)).toNumber()] = _769_v7;
        _nw131[(new BigNumber(8)).toNumber()] = ((_762_v0) ? (new _dafny.CodePoint('t'.codePointAt(0))) : (_769_v7));
        _770_v8 = _nw131;
        let _index156 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_770_v8).length));
        (_770_v8)[_index156] = new _dafny.CodePoint('o'.codePointAt(0));
        let _index157 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_770_v8).length));
        (_770_v8)[_index157] = _769_v7;
        (globalState).f4 = (_763_v1).isEqualTo(_763_v1);
        let _771_v9;
        let _nw132 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _771_v9 = _nw132;
        let _index158 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_771_v9).length));
        (_771_v9)[_index158] = _763_v1;
        let _index159 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_771_v9).length));
        (_771_v9)[_index159] = _module.__default.safeModuloInt(_763_v1, new BigNumber(217));
      } else {
        let _772_v10;
        _772_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(346),_763_v1);
        let _773_v11;
        _773_v11 = _module.D3.create_DC8(_772_v10);
        let _774_v12;
        _774_v12 = _dafny.Seq.of(((_762_v0) ? (_773_v11) : (_773_v11)), _773_v11);
        _774_v12 = _dafny.Seq.Concat(_774_v12, _774_v12);
        let _775_v13;
        let _init22 = ((_776_v1) => function (_777_i0) {
          return _dafny.Seq.of(_776_v1, _776_v1);
        })(_763_v1);
        let _nw133 = Array((new BigNumber(16)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw133.length); _i0_22++) {
          _nw133[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _775_v13 = _nw133;
        let _778_v14;
        _778_v14 = _dafny.Seq.of(_763_v1);
        let _index160 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length));
        (_775_v13)[_index160] = _dafny.Seq.Concat(_778_v14, _778_v14);
        let _779_v15;
        _779_v15 = _dafny.MultiSet.fromElements(_762_v0);
        let _index161 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length));
        let _rhs134 = _dafny.Seq.of(_763_v1, _763_v1, _763_v1, (((_779_v15).contains(_762_v0)) ? ((_779_v15).get(_762_v0)) : (_763_v1)));
        let _rhs135 = (_762_v0) && (_762_v0);
        let _lhs117 = _775_v13;
        let _lhs118 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length));
        _lhs117[_lhs118] = _rhs134;
        _762_v0 = _rhs135;
        let _780_v16;
        _780_v16 = _dafny.Map.Empty.slice().updateUnsafe(_762_v0,_763_v1);
        _780_v16 = (_780_v16).update(_762_v0, new BigNumber(((p0).Merge(p0)).length));
        let _781_v17;
        _781_v17 = _dafny.Map.Empty.slice().updateUnsafe(_763_v1,(_this).fm3((_775_v13)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length))], (_dafny.MultiSet.fromElements(_763_v1)).update(_763_v1, _module.__default.abs(new BigNumber(((_775_v13)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length))]).length))), _762_v0, globalState));
        let _782_v18;
        _782_v18 = _module.D7.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_763_v1,_762_v0));
        _781_v17 = (_782_v18).dtor_cf23;
        let _783_v19;
        _783_v19 = _dafny.Seq.UnicodeFromString("rd");
        let _784_v20;
        _784_v20 = _module.D2.create_DC6((_763_v1).plus(new BigNumber((_783_v19).length)), !(_762_v0), !((_763_v1).isLessThan(_763_v1)));
        let _source15 = _784_v20;
        if (_source15.is_DC6) {
          let _785___mcc_h0 = (_source15).cf6;
          let _786___mcc_h1 = (_source15).cf7;
          let _787___mcc_h2 = (_source15).cf8;
          let _788_cf8 = _787___mcc_h2;
          let _789_cf7 = _786___mcc_h1;
          let _790_cf6 = _785___mcc_h0;
          let _791_v21;
          _791_v21 = _dafny.Seq.of(_762_v0, _762_v0, true, _762_v0, _788_cf8);
          let _792_v22;
          _792_v22 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_791_v21, _791_v21));
          _792_v22 = (_792_v22).update(_791_v21, _module.__default.abs(_763_v1));
          let _793_v23;
          let _nw134 = Array((new BigNumber(12)).toNumber()).fill(_module.D5.Default());
          _793_v23 = _nw134;
          let _index162 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_793_v23).length));
          (_793_v23)[_index162] = _module.D5.create_DC15();
          let _794_v24;
          _794_v24 = _module.D5.create_DC15();
          let _index163 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_793_v23).length));
          (_793_v23)[_index163] = _794_v24;
          let _795_v25;
          let _796_v26;
          let _797_v27;
          let _798_v28;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector4 = (_this).m12(globalState);
          _out17 = _outcollector4[0];
          _out18 = _outcollector4[1];
          _out19 = _outcollector4[2];
          _out20 = _outcollector4[3];
          _795_v25 = _out17;
          _796_v26 = _out18;
          _797_v27 = _out19;
          _798_v28 = _out20;
          (globalState).f8 = _796_v26;
        } else if (_source15.is_DC7) {
          let _799___mcc_h3 = (_source15).cf9;
          let _800___mcc_h4 = (_source15).cf10;
          let _801___mcc_h5 = (_source15).cf11;
          let _802_cf11 = _801___mcc_h5;
          let _803_cf10 = _800___mcc_h4;
          let _804_cf9 = _799___mcc_h3;
          let _805_v29;
          _805_v29 = _dafny.MultiSet.fromElements(new BigNumber(730));
          let _806_v30;
          _806_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_805_v29).cardinality()),_804_cf9);
          let _807_v31;
          _807_v31 = _dafny.MultiSet.fromElements(new BigNumber(770), new BigNumber((_806_v30).length));
          let _808_v32;
          _808_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(_762_v0),_762_v0);
          let _809_v33;
          _809_v33 = _dafny.Seq.of(_807_v31, _module.__default.fm22(_763_v1, globalState));
          let _810_v34;
          _810_v34 = _dafny.Seq.of(((_805_v29).update(new BigNumber(510), _module.__default.abs(new BigNumber((_808_v32).length)))).IsProperSubsetOf((_809_v33)[_module.__default.safeIndex(new BigNumber((_783_v19).length), new BigNumber((_809_v33).length))]));
          let _811_v36;
          _811_v36 = _dafny.Map.Empty.slice().updateUnsafe(_763_v1,(function () {
            let _coll70 = new _dafny.Map();
            for (const _compr_70 of (_772_v10).Keys.Elements) {
              let _812_v35 = _compr_70;
              if ((_772_v10).contains(_812_v35)) {
                _coll70.push([(_812_v35).minus(_763_v1),false]);
              }
            }
            return _coll70;
          }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(_802_cf11,_803_cf10)).update(_763_v1, false)));
          let _rhs136 = _module.__default.fm23(_783_v19, globalState);
          let _rhs137 = _762_v0;
          let _rhs138 = (_dafny.Map.Empty.slice().updateUnsafe(_802_cf11,_781_v17)).Merge(_811_v36);
          let _lhs119 = globalState;
          _810_v34 = _rhs136;
          _lhs119.f4 = _rhs137;
          _811_v36 = _rhs138;
          _804_cf9 = new _dafny.CodePoint('a'.codePointAt(0));
          _783_v19 = _module.__default.fm17(_804_cf9, globalState);
          let _pat_let_tv39 = _762_v0;
          let _813_v37;
          _813_v37 = _dafny.Map.Empty.slice().updateUnsafe(_802_cf11,_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-851)), ((_814_v1) => function (_815_i1) {
            return _814_v1;
          })(_763_v1)), _module.__default.safeIndex((function (_pat_let29_0) {
            return function (_816_dt__update__tmp_h0) {
              return function (_pat_let30_0) {
                return function (_817_dt__update_hcf7_h0) {
                  return function (_pat_let31_0) {
                    return function (_818_dt__update_hcf6_h0) {
                      return _module.D2.create_DC6(_818_dt__update_hcf6_h0, _817_dt__update_hcf7_h0, (_816_dt__update__tmp_h0).dtor_cf8);
                    }(_pat_let31_0);
                  }(new BigNumber(455));
                }(_pat_let30_0);
              }(_pat_let_tv39);
            }(_pat_let29_0);
          }(_784_v20)).dtor_cf6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-851)), ((_819_v1) => function (_820_i1) {
            return _819_v1;
          })(_763_v1))).length)), _763_v1), _778_v14));
          _813_v37 = (_813_v37).update((_dafny.ZERO).minus((_763_v1).plus(_763_v1)), (_775_v13)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length))]);
        } else {
          let _821___mcc_h6 = (_source15).cf5;
          let _822_cf5 = _821___mcc_h6;
          let _823_v38;
          _823_v38 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-644)), function (_824_i2) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length))).length));
          let _825_v39;
          _825_v39 = new _dafny.CodePoint('h'.codePointAt(0));
          let _826_v40;
          _826_v40 = _dafny.Set.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), _825_v39);
          let _827_v41;
          _827_v41 = _dafny.Map.Empty.slice().updateUnsafe(_823_v38,(new BigNumber((_826_v40).length)).plus(_763_v1));
          _827_v41 = (_827_v41).update((_module.__default.fm22(_763_v1, globalState)).update(_763_v1, _module.__default.abs(_763_v1)), new BigNumber(850));
          (globalState).f8 = (_763_v1).plus(_763_v1);
          let _index164 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_822_cf5).length));
          (_822_cf5)[_index164] = _762_v0;
          let _index165 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_822_cf5).length));
          (_822_cf5)[_index165] = true;
          let _rhs139 = !((_822_cf5)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_822_cf5).length))]);
          let _rhs140 = _762_v0;
          let _rhs141 = _dafny.Seq.Concat((_775_v13)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_775_v13).length))], _778_v14);
          let _rhs142 = _762_v0;
          let _lhs120 = globalState;
          let _lhs121 = globalState;
          r1 = _rhs139;
          _lhs120.f4 = _rhs140;
          _lhs121.f0 = _rhs141;
          _762_v0 = _rhs142;
        }
      }
      (globalState).f4 = true;
      let _828_v42;
      let _nw135 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _828_v42 = _nw135;
      let _829_v43;
      _829_v43 = _dafny.Set.fromElements(_828_v42, _828_v42, _828_v42, _828_v42);
      let _830_i3;
      _830_i3 = _dafny.ZERO;
      L6: {
        while (!((_829_v43).Difference(_829_v43)).contains(_828_v42)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_830_i3)) {
              break L6;
            }
            _830_i3 = (_830_i3).plus(_dafny.ONE);
            (globalState).f4 = _762_v0;
            let _831_v44;
            _831_v44 = _dafny.Seq.of(_762_v0);
            let _832_v45;
            _832_v45 = _dafny.Seq.of(_762_v0, _dafny.Seq.IsPrefixOf(_831_v44, _831_v44), _762_v0);
            _832_v45 = _dafny.Seq.Concat(_831_v44, _dafny.Seq.update(_831_v44, _module.__default.safeIndex(_763_v1, new BigNumber((_831_v44).length)), _762_v0));
            let _833_v46;
            let _nw136 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _833_v46 = _nw136;
            let _834_v47;
            _834_v47 = _dafny.Seq.of(_763_v1, new BigNumber((_832_v45).length));
            let _835_v48;
            _835_v48 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_763_v1));
            let _836_v49;
            let _nw137 = Array((new BigNumber(25)).toNumber());
            _nw137[(_dafny.ZERO).toNumber()] = !(_762_v0);
            _nw137[(_dafny.ONE).toNumber()] = _762_v0;
            _nw137[(new BigNumber(2)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(3)).toNumber()] = (_762_v0);
            _nw137[(new BigNumber(4)).toNumber()] = false;
            _nw137[(new BigNumber(5)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(6)).toNumber()] = !(_762_v0);
            _nw137[(new BigNumber(7)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(8)).toNumber()] = !(_762_v0);
            _nw137[(new BigNumber(9)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(10)).toNumber()] = true;
            _nw137[(new BigNumber(11)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(12)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(13)).toNumber()] = false;
            _nw137[(new BigNumber(14)).toNumber()] = !(_762_v0);
            _nw137[(new BigNumber(15)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(16)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(17)).toNumber()] = !(true);
            _nw137[(new BigNumber(18)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(19)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(20)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(21)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(22)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(23)).toNumber()] = _762_v0;
            _nw137[(new BigNumber(24)).toNumber()] = (_this).fm3(_834_v47, _835_v48, _762_v0, globalState);
            _836_v49 = _nw137;
            let _837_v50;
            _837_v50 = _module.D2.create_DC5(_836_v49);
            let _838_v51;
            _838_v51 = _dafny.MultiSet.fromElements(_762_v0, _762_v0);
            let _839_v52;
            _839_v52 = _dafny.Map.Empty.slice().updateUnsafe(_763_v1,_838_v51);
            let _840_v53;
            _840_v53 = _dafny.Set.fromElements(_763_v1);
            let _rhs143 = _833_v46;
            let _rhs144 = (_this).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), ((_841_v1) => function (_842_i4) {
              return _841_v1;
            })(_763_v1)), _835_v48, (_this).fm3(_dafny.Seq.update(_834_v47, _module.__default.safeIndex(_763_v1, new BigNumber((_834_v47).length)), _763_v1), _835_v48, false, globalState), globalState);
            let _rhs145 = _837_v50;
            let _rhs146 = _module.__default.safeDivisionInt((new BigNumber(101)).minus(new BigNumber(782)), _763_v1);
            let _rhs147 = (_this).fm3(_dafny.Seq.update(_834_v47, _module.__default.safeIndex(new BigNumber((_839_v52).length), new BigNumber((_834_v47).length)), _763_v1), _dafny.MultiSet.fromElements(new BigNumber((_840_v53).length), _763_v1), false, globalState);
            let _lhs122 = globalState;
            _833_v46 = _rhs143;
            _762_v0 = _rhs144;
            _837_v50 = _rhs145;
            _lhs122.f8 = _rhs146;
            r1 = _rhs147;
            if (_762_v0) {
              (globalState).f8 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_762_v0),_762_v0)).length), new BigNumber(945));
              let _843_v54;
              let _nw138 = new _module.C1();
              _nw138.__ctor();
              _843_v54 = _nw138;
              let _844_v55;
              _844_v55 = _dafny.Seq.of(_843_v54);
              let _845_v56;
              _845_v56 = _dafny.Seq.of(_844_v55);
              let _846_v57;
              _846_v57 = _dafny.Seq.of((_845_v56)[_module.__default.safeIndex(_763_v1, new BigNumber((_845_v56).length))]);
              (globalState).f8 = (_763_v1).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_845_v56, _846_v57)).length))));
              let _847_v58;
              _847_v58 = new _dafny.CodePoint('t'.codePointAt(0));
              let _848_v59;
              _848_v59 = _dafny.MultiSet.fromElements(_847_v58);
              _848_v59 = ((_848_v59).Intersect(_848_v59)).Union(_848_v59);
              let _849_v60;
              let _nw139 = new _module.C1();
              _nw139.__ctor();
              _849_v60 = _nw139;
              let _index166 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_828_v42).length));
              (_828_v42)[_index166] = new BigNumber((_832_v45).length);
              let _index167 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_828_v42).length));
              let _rhs148 = _763_v1;
              let _rhs149 = _836_v49;
              let _rhs150 = !(_762_v0);
              let _lhs123 = _828_v42;
              let _lhs124 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_828_v42).length));
              _lhs123[_lhs124] = _rhs148;
              _836_v49 = _rhs149;
              _762_v0 = _rhs150;
            } else {
              (globalState).f4 = (_dafny.Set.fromElements(!(!((_this).fm3(_dafny.Seq.of(_763_v1), _dafny.MultiSet.FromArray(_834_v47), _762_v0, globalState))), _762_v0)).IsProperSubsetOf(_dafny.Set.fromElements(_762_v0));
              let _850_v61;
              let _nw140 = new _module.C1();
              _nw140.__ctor();
              _850_v61 = _nw140;
              let _851_v62;
              _851_v62 = _dafny.Map.Empty.slice().updateUnsafe(_763_v1,_763_v1);
              _851_v62 = _dafny.Map.Empty.slice().updateUnsafe(_763_v1,_module.__default.safeModuloInt(new BigNumber((_840_v53).length), _763_v1));
              (globalState).f8 = (((_835_v48).contains(_763_v1)) ? ((_835_v48).get(_763_v1)) : (((_762_v0) ? (_763_v1) : (new BigNumber((_851_v62).length)))));
              _851_v62 = (_851_v62).Merge(_851_v62);
            }
          }
        }
      }
      let _852_v63;
      let _init23 = ((_853_v0) => function (_854_i5) {
        return _853_v0;
      })(_762_v0);
      let _nw141 = Array((new BigNumber(14)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw141.length); _i0_23++) {
        _nw141[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _852_v63 = _nw141;
      let _855_v64;
      _855_v64 = _dafny.Seq.UnicodeFromString("wdjl");
      let _index168 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_852_v63).length));
      (_852_v63)[_index168] = _dafny.Seq.IsPrefixOf(_855_v64, _855_v64);
      let _index169 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_852_v63).length));
      let _rhs151 = _762_v0;
      let _rhs152 = _762_v0;
      let _rhs153 = _762_v0;
      let _rhs154 = _module.__default.safeModuloInt((_763_v1).plus(_763_v1), _763_v1);
      let _rhs155 = _762_v0;
      let _lhs125 = _852_v63;
      let _lhs126 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_852_v63).length));
      _762_v0 = _rhs151;
      _762_v0 = _rhs152;
      r1 = _rhs153;
      _763_v1 = _rhs154;
      _lhs125[_lhs126] = _rhs155;
      let _856_v65;
      _856_v65 = _dafny.Map.Empty.slice().updateUnsafe(_762_v0,(_852_v63)[_module.__default.safeIndex(new BigNumber(775), new BigNumber((_852_v63).length))]);
      r0 = (((((_856_v65).contains(_762_v0)) ? ((_856_v65).get(_762_v0)) : (_762_v0))) ? (new BigNumber(704)) : (_763_v1));
      r1 = _762_v0;
      r2 = _763_v1;
      return [r0, r1, r2];
    }
    m12(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _857_v0;
      _857_v0 = new BigNumber(912);
      let _858_v1;
      _858_v1 = true;
      (globalState).f4 = !(_857_v0).isEqualTo(_module.__default.fm20(_858_v1, globalState));
      let _859_v2;
      let _init24 = ((_860_v1) => function (_861_i0) {
        return _860_v1;
      })(_858_v1);
      let _nw142 = Array((new BigNumber(6)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw142.length); _i0_24++) {
        _nw142[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _859_v2 = _nw142;
      let _index170 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length));
      (_859_v2)[_index170] = _858_v1;
      let _index171 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length));
      (_859_v2)[_index171] = _858_v1;
      let _862_v3;
      _862_v3 = _dafny.Seq.UnicodeFromString("pncslplhn");
      let _863_v4;
      _863_v4 = _dafny.Map.Empty.slice().updateUnsafe(_862_v3,_857_v0);
      _863_v4 = (_863_v4).update(_dafny.Seq.UnicodeFromString("momthrphd"), new BigNumber(92));
      let _864_v5;
      _864_v5 = _dafny.Seq.of(new BigNumber(-139));
      (globalState).f0 = _dafny.Seq.Concat(_864_v5, _864_v5);
      let _865_v6;
      _865_v6 = _dafny.Set.fromElements(_862_v3);
      let _866_v7;
      _866_v7 = _module.D13.create_DC29(_865_v6);
      if ((_865_v6).equals(((_866_v7).dtor_cf41).Difference(_865_v6))) {
        (globalState).f4 = (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))];
        let _nw143 = Array((new BigNumber(5)).toNumber()).fill(false);
        _859_v2 = _nw143;
        let _867_v8;
        _867_v8 = _dafny.Map.Empty.slice().updateUnsafe(_857_v0,_857_v0);
        let _868_v9;
        _868_v9 = _module.D3.create_DC8(_867_v8);
        let _869_v10;
        _869_v10 = _dafny.Seq.of(false);
        let _870_v11;
        _870_v11 = _dafny.MultiSet.fromElements(_869_v10, _869_v10, _869_v10, _869_v10);
        let _871_v12;
        _871_v12 = _dafny.MultiSet.fromElements(_module.__default.fm26(_868_v9, _857_v0, _870_v11, globalState));
        let _872_v13;
        _872_v13 = _871_v12;
        (globalState).f4 = !((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]) || (((_872_v13)).IsProperSubsetOf(_871_v12));
        let _873_v14;
        _873_v14 = _dafny.MultiSet.fromElements(_862_v3);
        let _874_v15;
        _874_v15 = _dafny.MultiSet.fromElements((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], _858_v1, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], _858_v1);
        let _875_v16;
        _875_v16 = _dafny.Seq.of(_871_v12, _dafny.MultiSet.fromElements(new BigNumber((_873_v14).cardinality()), (((_874_v15).contains(_858_v1)) ? ((_874_v15).get(_858_v1)) : (_857_v0))));
        (globalState).f8 = new BigNumber((_875_v16).length);
        let _876_v17;
        let _nw144 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _876_v17 = _nw144;
        let _index172 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_876_v17).length));
        (_876_v17)[_index172] = ((_858_v1) ? (_857_v0) : (_857_v0));
        let _877_v18;
        _877_v18 = _dafny.Set.fromElements(_858_v1);
        let _878_v19;
        _878_v19 = _module.D9.create_DC20(_877_v18);
        let _879_v20;
        _879_v20 = _dafny.Seq.of(_877_v18, _877_v18, _877_v18);
        let _pat_let_tv40 = _877_v18;
        let _880_v21;
        let _nw145 = Array((new BigNumber(26)).toNumber());
        _nw145[(_dafny.ZERO).toNumber()] = _878_v19;
        _nw145[(_dafny.ONE).toNumber()] = _module.D9.create_DC20(_877_v18);
        _nw145[(new BigNumber(2)).toNumber()] = _module.D9.create_DC20(_dafny.Set.fromElements(_858_v1, _858_v1, _858_v1, _858_v1, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]));
        _nw145[(new BigNumber(3)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(4)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(5)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(6)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(7)).toNumber()] = _module.D9.create_DC20(_877_v18);
        _nw145[(new BigNumber(8)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(9)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(10)).toNumber()] = _module.D9.create_DC20(_877_v18);
        _nw145[(new BigNumber(11)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(12)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(13)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(14)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(15)).toNumber()] = _module.D9.create_DC20(_module.__default.fm35(_858_v1, _857_v0, _858_v1, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], globalState));
        _nw145[(new BigNumber(16)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(17)).toNumber()] = function (_pat_let32_0) {
          return function (_881_dt__update__tmp_h0) {
            return function (_pat_let33_0) {
              return function (_882_dt__update_hcf25_h0) {
                return _module.D9.create_DC20(_882_dt__update_hcf25_h0);
              }(_pat_let33_0);
            }(_pat_let_tv40);
          }(_pat_let32_0);
        }(_878_v19);
        _nw145[(new BigNumber(18)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(19)).toNumber()] = _module.D9.create_DC20((_879_v20)[_module.__default.safeIndex(_857_v0, new BigNumber((_879_v20).length))]);
        _nw145[(new BigNumber(20)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(21)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(22)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(23)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(24)).toNumber()] = _878_v19;
        _nw145[(new BigNumber(25)).toNumber()] = _878_v19;
        _880_v21 = _nw145;
        let _index173 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_880_v21).length));
        (_880_v21)[_index173] = _878_v19;
        let _index174 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length));
        let _index175 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_876_v17).length));
        let _index176 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_880_v21).length));
        let _rhs156 = (_module.__default.fm26(_module.D3.create_DC8(_867_v8), _857_v0, _dafny.MultiSet.fromElements(_869_v10, _dafny.Seq.of(_858_v1), _869_v10), globalState)).multipliedBy((_857_v0).minus(new BigNumber(859)));
        let _rhs157 = _858_v1;
        let _rhs158 = _857_v0;
        let _rhs159 = _module.D9.create_DC20(_877_v18);
        let _lhs127 = globalState;
        let _lhs128 = _859_v2;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length));
        let _lhs130 = _876_v17;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_876_v17).length));
        let _lhs132 = _880_v21;
        let _lhs133 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_880_v21).length));
        _lhs127.f8 = _rhs156;
        _lhs128[_lhs129] = _rhs157;
        _lhs130[_lhs131] = _rhs158;
        _lhs132[_lhs133] = _rhs159;
      } else {
        let _883_v22;
        _883_v22 = _dafny.MultiSet.fromElements(_858_v1);
        let _884_v23;
        let _nw146 = new _module.C1();
        _nw146.__ctor();
        _884_v23 = _nw146;
        let _885_v24;
        _885_v24 = _dafny.Map.Empty.slice().updateUnsafe(_884_v23,(_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]);
        let _886_v25;
        _886_v25 = _dafny.Map.Empty.slice().updateUnsafe((_885_v24).update(_884_v23, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]),_dafny.MultiSet.fromElements(!((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]), (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]));
        let _rhs160 = new BigNumber((_862_v3).length);
        let _rhs161 = new BigNumber(16);
        let _rhs162 = ((_883_v22).Difference((((_886_v25).contains(_885_v24)) ? ((_886_v25).get(_885_v24)) : (_883_v22)))).IsSubsetOf(_883_v22);
        let _rhs163 = _dafny.Seq.Concat(_862_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_887_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
        r0 = _rhs160;
        r1 = _rhs161;
        r3 = _rhs162;
        _862_v3 = _rhs163;
        (globalState).f8 = _857_v0;
        let _888_v26;
        _888_v26 = new _dafny.CodePoint('n'.codePointAt(0));
        let _889_v27;
        _889_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_module.__default.fm17(_888_v26, globalState), _dafny.Seq.UnicodeFromString("isrckjn")),(_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]);
        _889_v27 = (_889_v27).update(_858_v1, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]);
        let _890_v28;
        let _nw147 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
        _890_v28 = _nw147;
        let _891_v29;
        _891_v29 = _dafny.Map.Empty.slice().updateUnsafe((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))],_857_v0);
        let _892_v30;
        _892_v30 = _dafny.Set.fromElements(_891_v29);
        let _index177 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_890_v28).length));
        (_890_v28)[_index177] = _892_v30;
        let _index178 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_890_v28).length));
        (_890_v28)[_index178] = ((_892_v30).Union(_892_v30)).Difference(_892_v30);
        let _893_v31;
        _893_v31 = _dafny.Map.Empty.slice().updateUnsafe(_857_v0,_858_v1);
        let _894_v32;
        _894_v32 = _module.D7.create_DC17((_893_v31).update(new BigNumber(-103), _858_v1));
        let _source16 = _894_v32;
        if (_source16.is_DC18) {
          (globalState).f8 = _module.__default.safeDivisionInt(_857_v0, _857_v0);
          let _895_v33;
          let _nw148 = new _module.C1();
          _nw148.__ctor();
          _895_v33 = _nw148;
          let _896_v35;
          _896_v35 = _dafny.Seq.of((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]);
          let _897_v36;
          _897_v36 = _dafny.MultiSet.fromElements(_896_v35);
          let _898_v37;
          _898_v37 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), function (_899_i2) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length));
          let _900_v38;
          _900_v38 = _dafny.Set.fromElements(new BigNumber((_898_v37).length), _857_v0, _857_v0, new BigNumber((_865_v6).length), _857_v0);
          let _901_v39;
          _901_v39 = _dafny.MultiSet.fromElements(_857_v0);
          let _902_v40;
          let _nw149 = Array((new BigNumber(23)).toNumber());
          _nw149[(_dafny.ZERO).toNumber()] = _857_v0;
          _nw149[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_857_v0).plus(_857_v0));
          _nw149[(new BigNumber(2)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(3)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(4)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_857_v0, _857_v0);
          _nw149[(new BigNumber(6)).toNumber()] = new BigNumber((function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of (_897_v36).Elements) {
              let _903_v34 = _compr_71;
              if ((_897_v36).contains(_903_v34)) {
                _coll71.push([_903_v34,_857_v0]);
              }
            }
            return _coll71;
          }()).length);
          _nw149[(new BigNumber(7)).toNumber()] = new BigNumber((_898_v37).length);
          _nw149[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_857_v0, new BigNumber(371));
          _nw149[(new BigNumber(9)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(10)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(11)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(12)).toNumber()] = new BigNumber((_891_v29).length);
          _nw149[(new BigNumber(13)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_857_v0,_858_v1)).length);
          _nw149[(new BigNumber(15)).toNumber()] = new BigNumber(-418);
          _nw149[(new BigNumber(16)).toNumber()] = (_857_v0).plus(new BigNumber((_dafny.Seq.of((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], true, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], _858_v1, _858_v1)).length));
          _nw149[(new BigNumber(17)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(18)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(19)).toNumber()] = (_864_v5)[_module.__default.safeIndex(_857_v0, new BigNumber((_864_v5).length))];
          _nw149[(new BigNumber(20)).toNumber()] = (((_901_v39).contains(_857_v0)) ? ((_901_v39).get(_857_v0)) : (_857_v0));
          _nw149[(new BigNumber(21)).toNumber()] = _857_v0;
          _nw149[(new BigNumber(22)).toNumber()] = (((_891_v29).contains((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))])) ? ((_891_v29).get((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))])) : (new BigNumber((_896_v35).length)));
          _902_v40 = _nw149;
          let _index179 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_902_v40).length));
          (_902_v40)[_index179] = _857_v0;
          let _904_v41;
          _904_v41 = _module.D5.create_DC14(new BigNumber((_module.__default.fm17(_888_v26, globalState)).length), _858_v1);
          let _index180 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_902_v40).length));
          (_902_v40)[_index180] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus((_857_v0).minus(new BigNumber((_dafny.Seq.update(_862_v3, _module.__default.safeIndex((_904_v41).dtor_cf20, new BigNumber((_862_v3).length)), _888_v26)).length))), _857_v0));
          let _905_v42;
          let _nw150 = new _module.C1();
          _nw150.__ctor();
          _905_v42 = _nw150;
        } else {
          let _906___mcc_h0 = (_source16).cf23;
          let _907_cf23 = _906___mcc_h0;
          _888_v26 = _888_v26;
          r1 = new BigNumber(-346);
          _893_v31 = _893_v31;
          let _908_v43;
          let _nw151 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _908_v43 = _nw151;
          let _909_v44;
          _909_v44 = _module.D10.create_DC24(new BigNumber(-90), _858_v1, _dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0))));
          let _910_v45;
          _910_v45 = _dafny.Set.fromElements(new BigNumber(238), new BigNumber((_dafny.MultiSet.fromElements(_857_v0)).cardinality()));
          let _911_v46;
          _911_v46 = _dafny.Set.fromElements(_857_v0, new BigNumber((_910_v45).length));
          let _912_v47;
          let _nw152 = new _module.C0();
          _nw152.__ctor();
          _912_v47 = _nw152;
          let _913_v48;
          _913_v48 = _module.D13.create_DC30(_858_v1, (_909_v44).dtor_cf32, _911_v46, _912_v47);
          let _pat_let_tv41 = _858_v1;
          let _index181 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_908_v43).length));
          (_908_v43)[_index181] = (function (_pat_let34_0) {
            return function (_914_dt__update__tmp_h1) {
              return function (_pat_let35_0) {
                return function (_915_dt__update_hcf42_h0) {
                  return _module.D13.create_DC30(_915_dt__update_hcf42_h0, (_914_dt__update__tmp_h1).dtor_cf43, (_914_dt__update__tmp_h1).dtor_cf44, (_914_dt__update__tmp_h1).dtor_cf45);
                }(_pat_let35_0);
              }(_pat_let_tv41);
            }(_pat_let34_0);
          }(_913_v48)).dtor_cf44;
          let _index182 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_908_v43).length));
          (_908_v43)[_index182] = _911_v46;
        }
      }
      let _916_v49;
      _916_v49 = _dafny.Seq.of((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], true, (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]);
      let _917_v50;
      _917_v50 = _module.D4.create_DC10(_916_v49);
      let _pat_let_tv42 = _916_v49;
      _917_v50 = function (_pat_let36_0) {
        return function (_918_dt__update__tmp_h2) {
          return function (_pat_let37_0) {
            return function (_919_dt__update_hcf13_h0) {
              return _module.D4.create_DC10(_919_dt__update_hcf13_h0);
            }(_pat_let37_0);
          }(_pat_let_tv42);
        }(_pat_let36_0);
      }(_917_v50);
      r0 = _857_v0;
      r1 = new BigNumber(((((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]) ? (_916_v49) : (_dafny.Seq.Concat(_dafny.Seq.of((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))], (_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]), _dafny.Seq.of((_859_v2)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_859_v2).length))]))))).length);
      let _920_v51;
      _920_v51 = _module.D0.create_DC0(_857_v0);
      let _pat_let_tv43 = _859_v2;
      let _pat_let_tv44 = _859_v2;
      let _pat_let_tv45 = _859_v2;
      let _pat_let_tv46 = _859_v2;
      let _pat_let_tv47 = _857_v0;
      let _pat_let_tv48 = _857_v0;
      let _pat_let_tv49 = _857_v0;
      r2 = function (_source17) {
        if (_source17.is_DC1) {
          let _921___mcc_h1 = (_source17).cf1;
          let _922___mcc_h2 = (_source17).cf2;
          let _923_cf2 = _922___mcc_h2;
          let _924_cf1 = _921___mcc_h1;
          return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv44)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_pat_let_tv43).length))],_923_cf2)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv46)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_pat_let_tv45).length))],new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_923_cf2,new BigNumber((_dafny.MultiSet.fromElements(_923_cf2, _923_cf2)).cardinality()))).length)))).length);
        } else if (_source17.is_DC2) {
          return _pat_let_tv47;
        } else if (_source17.is_DC0) {
          let _925___mcc_h3 = (_source17).cf0;
          let _926_cf0 = _925___mcc_h3;
          return (_926_cf0).multipliedBy((_dafny.ZERO).minus(new BigNumber(-441)));
        } else {
          let _927___mcc_h4 = (_source17).cf3;
          let _928_cf3 = _927___mcc_h4;
          return _pat_let_tv48;
        }
      }(function (_pat_let38_0) {
        return function (_929_dt__update__tmp_h3) {
          return function (_pat_let39_0) {
            return function (_930_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_930_dt__update_hcf0_h0);
            }(_pat_let39_0);
          }(_pat_let_tv49);
        }(_pat_let38_0);
      }(_920_v51));
      let _931_v52;
      let _nw153 = Array((new BigNumber(18)).toNumber());
      _931_v52 = _nw153;
      let _932_v53;
      _932_v53 = _dafny.Seq.of(_931_v52);
      r3 = _dafny.Seq.contains(_932_v53, (_932_v53)[_module.__default.safeIndex(_857_v0, new BigNumber((_932_v53).length))]);
      return [r0, r1, r2, r3];
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
    fm2(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-878),new BigNumber(740))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(663),new BigNumber((_dafny.Seq.of(!(true), false)).length)))).Merge(function () {
        let _coll72 = new _dafny.Map();
        for (const _compr_72 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll73 = new _dafny.Map();
          for (const _compr_73 of (_dafny.Set.fromElements(_dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC2())), _dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2())), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length))), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(603)))))).Elements) {
            let _933_v1 = _compr_73;
            if ((_dafny.Set.fromElements(_dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC2())), _dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2())), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length))), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(603)))))).contains(_933_v1)) {
              _coll73.push([_933_v1,false]);
            }
          }
          return _coll73;
        }()).length),true)).Keys.Elements) {
          let _934_v0 = _compr_72;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll74 = new _dafny.Map();
            for (const _compr_74 of (_dafny.Set.fromElements(_dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC2())), _dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2())), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length))), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(603)))))).Elements) {
              let _933_v1 = _compr_74;
              if ((_dafny.Set.fromElements(_dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC2())), _dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2())), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length))), _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(603)))))).contains(_933_v1)) {
                _coll74.push([_933_v1,false]);
              }
            }
            return _coll74;
          }()).length),true)).contains(_934_v0)) {
            _coll72.push([_module.__default.safeDivisionInt(_934_v0, new BigNumber((_dafny.Seq.UnicodeFromString("idrigq")).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-517)), function (_935_i0) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll72;
      }());
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return !(!(false));
    };
    fm15(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(401),new BigNumber(811))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ifkhv")).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()),false)).length))))).cardinality());
    };
    fm16(p0, globalState) {
      let _this = this;
      if (true) {
        return (true) === (!(true));
      } else {
        return !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))))).length)).isEqualTo(new BigNumber((_dafny.Seq.of(false)).length));
      }
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _936_v0;
      let _init25 = function (_937_i0) {
        return !(false);
      };
      let _nw154 = Array((new BigNumber(29)).toNumber());
      for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw154.length); _i0_25++) {
        _nw154[_i0_25] = _init25(new BigNumber(_i0_25));
      }
      _936_v0 = _nw154;
      let _938_v1;
      _938_v1 = true;
      let _index183 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
      (_936_v0)[_index183] = _938_v1;
      let _939_v2;
      _939_v2 = new BigNumber(370);
      let _index184 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
      (_936_v0)[_index184] = ((_dafny.ZERO).minus(_939_v2)).isLessThan(_939_v2);
      let _hi4 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-926)), function (_940_i2) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length);
      for (let _941_i1 = _939_v2; _941_i1.isLessThan(_hi4); _941_i1 = _941_i1.plus(_dafny.ONE)) {
        let _942_v3;
        _942_v3 = _dafny.Seq.UnicodeFromString("fm");
        let _943_v4;
        _943_v4 = _dafny.Seq.of(_942_v3);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_943_v4)[_module.__default.safeIndex(_939_v2, new BigNumber((_943_v4).length))], _942_v3), _942_v3)) {
          let _944_v5;
          _944_v5 = _module.D0.create_DC1(_941_i1, _941_i1);
          (globalState).f8 = (_dafny.ZERO).minus((_944_v5).dtor_cf1);
          let _945_v6;
          _945_v6 = new _dafny.CodePoint('w'.codePointAt(0));
          _945_v6 = _945_v6;
          _942_v3 = _dafny.Seq.update(_dafny.Seq.Concat(_942_v3, _942_v3), _module.__default.safeIndex(_939_v2, new BigNumber((_dafny.Seq.Concat(_942_v3, _942_v3)).length)), _945_v6);
          let _index185 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
          (_936_v0)[_index185] = !(_938_v1);
          let _946_v7;
          let _nw155 = new _module.C0();
          _nw155.__ctor();
          _946_v7 = _nw155;
        } else {
          let _947_v8;
          _947_v8 = _dafny.Map.Empty.slice().updateUnsafe(_941_i1,_939_v2);
          let _948_v9;
          _948_v9 = _module.D10.create_DC24((((_947_v8).contains(_941_i1)) ? ((_947_v8).get(_941_i1)) : (_939_v2)), (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], _942_v3);
          r0 = (_948_v9).dtor_cf32;
          (globalState).f8 = _939_v2;
          (globalState).f8 = _939_v2;
          let _949_v10;
          _949_v10 = _dafny.Set.fromElements(_941_i1, new BigNumber(903), (_dafny.ZERO).minus(_939_v2));
          let _950_v11;
          let _nw156 = new _module.C0();
          _nw156.__ctor();
          _950_v11 = _nw156;
          let _951_v12;
          _951_v12 = _module.D13.create_DC30(true, _938_v1, _949_v10, _950_v11);
          r1 = (_951_v12).dtor_cf42;
          let _952_v13;
          _952_v13 = _dafny.Map.Empty.slice().updateUnsafe((_939_v2).isLessThan(_939_v2),(_950_v11).fm19(globalState));
          _952_v13 = ((_952_v13).Merge(_952_v13)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_938_v1,_941_i1));
        }
        let _953_v14;
        _953_v14 = _dafny.Map.Empty.slice().updateUnsafe(_939_v2,_938_v1);
        let _954_v15;
        _954_v15 = _dafny.Seq.of(_953_v14);
        let _955_v16;
        _955_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_954_v15, _954_v15),_941_i1);
        _955_v16 = (_955_v16).update(false, (_939_v2).minus(_939_v2));
        let _index186 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
        let _rhs164 = !((_941_i1).isLessThan(_941_i1));
        let _rhs165 = _941_i1;
        let _rhs166 = _938_v1;
        let _rhs167 = _938_v1;
        let _rhs168 = (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))];
        let _lhs134 = _936_v0;
        let _lhs135 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
        let _lhs136 = globalState;
        let _lhs137 = globalState;
        let _lhs138 = globalState;
        _lhs134[_lhs135] = _rhs164;
        _lhs136.f8 = _rhs165;
        _lhs137.f4 = _rhs166;
        r0 = _rhs167;
        _lhs138.f4 = _rhs168;
        if ((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) {
          let _956_v17;
          _956_v17 = _dafny.Set.fromElements(_942_v3, _dafny.Seq.UnicodeFromString("dfqsa"), _dafny.Seq.Concat(_942_v3, _942_v3), _942_v3, _942_v3);
          _956_v17 = (_dafny.Set.fromElements(_942_v3)).Union(_956_v17);
          let _rhs169 = _941_i1;
          let _rhs170 = (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))];
          let _rhs171 = _939_v2;
          let _lhs139 = globalState;
          _939_v2 = _rhs169;
          _938_v1 = _rhs170;
          _lhs139.f8 = _rhs171;
          let _957_v18;
          _957_v18 = _dafny.Map.Empty.slice().updateUnsafe(_941_i1,_941_i1);
          let _958_v20;
          let _nw157 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _958_v20 = _nw157;
          let _959_v21;
          _959_v21 = _dafny.Map.Empty.slice().updateUnsafe(_958_v20,false);
          let _960_v22;
          _960_v22 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll75 = new _dafny.Set();
            for (const _compr_75 of (_957_v18).Keys.Elements) {
              let _961_v19 = _compr_75;
              if ((_957_v18).contains(_961_v19)) {
                _coll75.add((_961_v19).multipliedBy(new BigNumber(664)));
              }
            }
            return _coll75;
          }(),(((_959_v21).contains(_958_v20)) ? ((_959_v21).get(_958_v20)) : (_938_v1)));
          _960_v22 = _960_v22;
          let _962_v23;
          let _out21;
          _out21 = (_this).m10(_941_i1, (_953_v14).Merge(_953_v14), new BigNumber(329), globalState);
          _962_v23 = _out21;
          let _963_v24;
          _963_v24 = _dafny.Seq.of((_dafny.ZERO).minus(_939_v2), (_dafny.ZERO).minus((_dafny.ZERO).minus(_939_v2)), new BigNumber((_942_v3).length));
          let _964_v25;
          let _nw158 = Array((new BigNumber(6)).toNumber());
          _nw158[(_dafny.ZERO).toNumber()] = _939_v2;
          _nw158[(_dafny.ONE).toNumber()] = (_963_v24)[_module.__default.safeIndex(_941_i1, new BigNumber((_963_v24).length))];
          _nw158[(new BigNumber(2)).toNumber()] = _941_i1;
          _nw158[(new BigNumber(3)).toNumber()] = _941_i1;
          _nw158[(new BigNumber(4)).toNumber()] = _939_v2;
          _nw158[(new BigNumber(5)).toNumber()] = _939_v2;
          _964_v25 = _nw158;
          let _965_v26;
          _965_v26 = _module.D3.create_DC8(_957_v18);
          let _966_v27;
          _966_v27 = _dafny.Seq.of((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]);
          let _967_v28;
          _967_v28 = _dafny.MultiSet.fromElements(_966_v27);
          let _index187 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
          let _rhs172 = _module.__default.safeDivisionInt((_939_v2).multipliedBy(_module.__default.fm26(_965_v26, _939_v2, _967_v28, globalState)), new BigNumber(853));
          let _rhs173 = (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))];
          let _rhs174 = _964_v25;
          let _rhs175 = _dafny.Seq.Concat(_dafny.Seq.Concat(_942_v3, _942_v3), _942_v3);
          let _rhs176 = false;
          let _lhs140 = globalState;
          let _lhs141 = _936_v0;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
          _lhs140.f3 = _rhs172;
          _lhs141[_lhs142] = _rhs173;
          _964_v25 = _rhs174;
          _942_v3 = _rhs175;
          r0 = _rhs176;
        } else {
          _938_v1 = _938_v1;
          _939_v2 = _module.__default.fm20((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], globalState);
          let _968_v29;
          let _nw159 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _968_v29 = _nw159;
          let _index188 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_968_v29).length));
          (_968_v29)[_index188] = _941_i1;
          let _index189 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_968_v29).length));
          (_968_v29)[_index189] = new BigNumber((_942_v3).length);
          _968_v29 = _968_v29;
          let _969_v30;
          _969_v30 = new _dafny.CodePoint('r'.codePointAt(0));
          _969_v30 = _969_v30;
        }
      }
      if ((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) {
        (globalState).f8 = ((_939_v2).minus(_939_v2)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(625), new BigNumber(995)))));
        let _970_v31;
        _970_v31 = _dafny.MultiSet.fromElements(_939_v2, _939_v2);
        let _971_v32;
        _971_v32 = _dafny.Set.fromElements(_939_v2);
        let _972_v33;
        let _nw160 = new _module.C0();
        _nw160.__ctor();
        _972_v33 = _nw160;
        let _973_v34;
        _973_v34 = _module.D13.create_DC30(_938_v1, true, _971_v32, _972_v33);
        let _index190 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
        let _rhs177 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) ? (_939_v2) : (new BigNumber(889)))));
        let _rhs178 = (_970_v31).IsSubsetOf((_970_v31).update(_939_v2, _module.__default.abs(_939_v2)));
        let _rhs179 = (_973_v34).dtor_cf42;
        let _rhs180 = (_939_v2).isEqualTo(_939_v2);
        let _lhs143 = globalState;
        let _lhs144 = _936_v0;
        let _lhs145 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
        _939_v2 = _rhs177;
        _lhs143.f4 = _rhs178;
        _lhs144[_lhs145] = _rhs179;
        _938_v1 = _rhs180;
        let _974_v35;
        let _nw161 = Array((new BigNumber(18)).toNumber()).fill([]);
        _974_v35 = _nw161;
        let _975_v36;
        let _nw162 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _975_v36 = _nw162;
        let _index191 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length));
        (_974_v35)[_index191] = _975_v36;
        let _index192 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length));
        (_974_v35)[_index192] = _975_v36;
        let _976_v37;
        _976_v37 = _dafny.Set.fromElements(_938_v1);
        _976_v37 = (_976_v37).Difference(_dafny.Set.fromElements((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]));
        if ((_938_v1) || (_938_v1)) {
          (globalState).f4 = !((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]);
          r1 = (_this).fm16(_938_v1, globalState);
          let _977_v38;
          let _nw163 = new _module.C1();
          _nw163.__ctor();
          _977_v38 = _nw163;
          let _978_v39;
          _978_v39 = _module.D0.create_DC2();
          let _979_v40;
          _979_v40 = _module.D9.create_DC21((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], _dafny.Seq.of((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], _938_v1, (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]), _938_v1);
          let _rhs181 = _978_v39;
          let _rhs182 = (_dafny.ZERO).minus(_939_v2);
          let _rhs183 = ((_dafny.Seq.IsProperPrefixOf((_979_v40).dtor_cf27, _dafny.Seq.of(_938_v1, (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], false))) ? ((_939_v2).isLessThanOrEqualTo(_939_v2)) : (((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) && (false)));
          let _rhs184 = _938_v1;
          let _lhs146 = globalState;
          _978_v39 = _rhs181;
          _lhs146.f3 = _rhs182;
          r0 = _rhs183;
          r1 = _rhs184;
          let _980_v41;
          let _init26 = function (_981_i3) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          };
          let _nw164 = Array((new BigNumber(23)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw164.length); _i0_26++) {
            _nw164[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _980_v41 = _nw164;
          let _982_v42;
          _982_v42 = _dafny.Seq.of((((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) ? (_980_v41) : (_980_v41)), _980_v41, _980_v41);
          let _rhs185 = _dafny.Seq.Concat(_dafny.Seq.Concat(_982_v42, _982_v42), _982_v42);
          let _rhs186 = _module.__default.safeModuloInt(_939_v2, new BigNumber(776));
          let _lhs147 = globalState;
          _982_v42 = _rhs185;
          _lhs147.f8 = _rhs186;
        } else {
          let _983_v43;
          _983_v43 = _dafny.Seq.UnicodeFromString("ae");
          let _984_v44;
          _984_v44 = _dafny.Map.Empty.slice().updateUnsafe(_983_v43,_939_v2);
          _984_v44 = _984_v44;
          let _985_v45;
          _985_v45 = new _dafny.CodePoint('g'.codePointAt(0));
          let _986_v46;
          _986_v46 = _module.D3.create_DC9();
          let _987_v47;
          _987_v47 = _dafny.Map.Empty.slice().updateUnsafe(_985_v45,_986_v46);
          _987_v47 = (_987_v47).update(_985_v45, _986_v46);
          let _988_v48;
          _988_v48 = _dafny.Map.Empty.slice().updateUnsafe(!(_938_v1),_939_v2);
          let _989_v49;
          _989_v49 = _dafny.Map.Empty.slice().updateUnsafe(_939_v2,new BigNumber((_988_v48).length));
          let _990_v50;
          _990_v50 = _dafny.Seq.of(new BigNumber(((_988_v48).update((_this).fm16((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], globalState), new BigNumber((_983_v43).length))).length), new BigNumber((_989_v49).length));
          (globalState).f4 = (_this).fm3(_990_v50, (_970_v31).Difference(_970_v31), _938_v1, globalState);
          let _index193 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_975_v36).length));
          (_975_v36)[_index193] = new BigNumber((_989_v49).length);
          let _index194 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_975_v36).length));
          (_975_v36)[_index194] = _module.__default.safeModuloInt(_939_v2, (_990_v50)[_module.__default.safeIndex(_939_v2, new BigNumber((_990_v50).length))]);
          let _991_v51;
          _991_v51 = _dafny.Seq.of((_974_v35)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length))], ((_938_v1) ? ((_974_v35)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length))]) : ((_974_v35)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length))])), (_974_v35)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length))], (_974_v35)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_974_v35).length))]);
          _991_v51 = ((true) ? (_dafny.Seq.Concat(_991_v51, _991_v51)) : (_991_v51));
        }
      } else {
        let _992_v52;
        let _nw165 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _992_v52 = _nw165;
        let _993_v53;
        _993_v53 = _dafny.Map.Empty.slice().updateUnsafe(_939_v2,_992_v52);
        _993_v53 = (_993_v53).update(_939_v2, _992_v52);
        let _994_v54;
        _994_v54 = _dafny.Seq.UnicodeFromString("aeeyibtfx");
        let _995_v55;
        _995_v55 = _dafny.MultiSet.fromElements(_994_v54);
        let _996_v56;
        _996_v56 = _module.D10.create_DC23(_995_v55);
        let _pat_let_tv50 = _995_v55;
        _996_v56 = function (_pat_let40_0) {
          return function (_997_dt__update__tmp_h0) {
            return function (_pat_let41_0) {
              return function (_998_dt__update_hcf30_h0) {
                return _module.D10.create_DC23(_998_dt__update_hcf30_h0);
              }(_pat_let41_0);
            }(_pat_let_tv50);
          }(_pat_let40_0);
        }(((_938_v1) ? (_module.__default.fm36(new BigNumber((_994_v54).length), _dafny.Set.fromElements((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]), _938_v1, globalState)) : (_module.D10.create_DC23(_995_v55))));
        (globalState).f2 = _module.__default.fm27(_938_v1, globalState);
        let _999_v57;
        _999_v57 = _dafny.Seq.of(new BigNumber(-589));
        let _1000_v58;
        _1000_v58 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wc")).length), new BigNumber((_999_v57).length));
        let _1001_v59;
        _1001_v59 = _dafny.Seq.of((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]);
        let _1002_v60;
        _1002_v60 = _module.D10.create_DC24(_939_v2, _938_v1, _994_v54);
        let _1003_v61;
        _1003_v61 = _module.D9.create_DC21(_938_v1, _1001_v59, (_1002_v60).dtor_cf32);
        let _index195 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length));
        (_936_v0)[_index195] = ((((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) ? ((_this).fm3(_999_v57, _1000_v58, _938_v1, globalState)) : ((_1003_v61).dtor_cf26))) || (false);
        let _1004_v62;
        _1004_v62 = new _dafny.CodePoint('o'.codePointAt(0));
        (globalState).f3 = (_939_v2).plus(new BigNumber((_dafny.Seq.update(_994_v54, _module.__default.safeIndex(_939_v2, new BigNumber((_994_v54).length)), _1004_v62)).length));
      }
      (globalState).f3 = new BigNumber((_dafny.Seq.UnicodeFromString("h")).length);
      let _hi5 = _939_v2;
      for (let _1005_i4 = _939_v2; _1005_i4.isLessThan(_hi5); _1005_i4 = _1005_i4.plus(_dafny.ONE)) {
        (globalState).f3 = _module.__default.safeModuloInt(_1005_i4, new BigNumber((_dafny.MultiSet.fromElements(_938_v1)).cardinality()));
        r0 = (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))];
        let _1006_v63;
        let _nw166 = new _module.C2();
        _nw166.__ctor();
        _1006_v63 = _nw166;
        (globalState).f3 = _1005_i4;
      }
      let _1007_v64;
      _1007_v64 = _module.D2.create_DC7(new _dafny.CodePoint('l'.codePointAt(0)), (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], _module.__default.fm20((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], globalState));
      let _source18 = _1007_v64;
      if (_source18.is_DC6) {
        let _1008___mcc_h0 = (_source18).cf6;
        let _1009___mcc_h1 = (_source18).cf7;
        let _1010___mcc_h2 = (_source18).cf8;
        let _1011_cf8 = _1010___mcc_h2;
        let _1012_cf7 = _1009___mcc_h1;
        let _1013_cf6 = _1008___mcc_h0;
        let _1014_v65;
        _1014_v65 = _dafny.Map.Empty.slice().updateUnsafe(_938_v1,_938_v1);
        let _1015_v66;
        _1015_v66 = _dafny.MultiSet.fromElements(!(_938_v1), _1012_cf7, _1012_cf7, (((_1014_v65).contains(_1012_cf7)) ? ((_1014_v65).get(_1012_cf7)) : (!(false))));
        let _1016_v67;
        _1016_v67 = _dafny.Map.Empty.slice().updateUnsafe((_1015_v66).Union(_1015_v66),_1013_cf6);
        let _1017_v68;
        _1017_v68 = _module.D0.create_DC1(_1013_cf6, _1013_cf6);
        let _1018_v69;
        _1018_v69 = _module.D0.create_DC3(_1017_v68);
        let _1019_v70;
        _1019_v70 = _dafny.Seq.of(_1018_v69, _1018_v69, _module.D0.create_DC3(_1017_v68), _module.__default.fm38(_938_v1, globalState), _module.D0.create_DC3(_1017_v68));
        let _1020_v71;
        _1020_v71 = new _dafny.CodePoint('o'.codePointAt(0));
        _1016_v67 = _module.__default.fm37(_1019_v70, _module.__default.fm17(_1020_v71, globalState), _939_v2, (_939_v2).plus(_1013_cf6), globalState);
        let _1021_v73;
        _1021_v73 = _dafny.Map.Empty.slice().updateUnsafe(_939_v2,_1013_cf6);
        let _1022_v74;
        _1022_v74 = _dafny.Set.fromElements((function () {
          let _coll76 = new _dafny.Map();
          for (const _compr_76 of _dafny.IntegerRange(new BigNumber(143), new BigNumber(-189))) {
            let _1023_v72 = _compr_76;
            if (((new BigNumber(143)).isLessThanOrEqualTo(_1023_v72)) && ((_1023_v72).isLessThan(new BigNumber(-189)))) {
              _coll76.push([(_1023_v72).multipliedBy(new BigNumber((_dafny.Seq.of((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))], !(_1012_cf7), !(_938_v1), _1011_cf8)).length)),_939_v2]);
            }
          }
          return _coll76;
        }()).Merge(_1021_v73));
        let _1024_v75;
        _1024_v75 = _dafny.Seq.of(_1022_v74);
        _1022_v74 = (_1024_v75)[_module.__default.safeIndex(_1013_cf6, new BigNumber((_1024_v75).length))];
        let _1025_v76;
        let _nw167 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1025_v76 = _nw167;
        _1025_v76 = _1025_v76;
        let _1026_v77;
        _1026_v77 = _dafny.Seq.of((((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) ? (_1011_cf8) : ((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))])));
        let _1027_v78;
        let _nw168 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1027_v78 = _nw168;
        let _1028_v79;
        _1028_v79 = _dafny.Set.fromElements(_1027_v78);
        let _1029_v80;
        _1029_v80 = _dafny.Set.fromElements(_1012_cf7, (_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]);
        _1026_v77 = _module.__default.fm31((_1028_v79).IsProperSubsetOf(_dafny.Set.fromElements(_1027_v78, _1027_v78, _1027_v78)), (_dafny.Set.fromElements((_1026_v77)[_module.__default.safeIndex((_dafny.ZERO).minus(_1013_cf6), new BigNumber((_1026_v77).length))], _1011_cf8)).IsSubsetOf(_1029_v80), !(((_1012_cf7) ? ((_936_v0)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_936_v0).length))]) : (_938_v1))), globalState);
      } else if (_source18.is_DC7) {
        let _1030___mcc_h3 = (_source18).cf9;
        let _1031___mcc_h4 = (_source18).cf10;
        let _1032___mcc_h5 = (_source18).cf11;
        let _1033_cf11 = _1032___mcc_h5;
        let _1034_cf10 = _1031___mcc_h4;
        let _1035_cf9 = _1030___mcc_h3;
        let _1036_v81;
        _1036_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1034_cf10,_1035_cf9);
        _1036_v81 = (_1036_v81).update((_this).fm16(_1034_cf10, globalState), _1035_cf9);
        (globalState).f4 = false;
        let _1037_v82;
        let _nw169 = new _module.C1();
        _nw169.__ctor();
        _1037_v82 = _nw169;
        let _1038_v83;
        let _nw170 = new _module.C0();
        _nw170.__ctor();
        _1038_v83 = _nw170;
      } else {
        let _1039___mcc_h6 = (_source18).cf5;
        let _1040_cf5 = _1039___mcc_h6;
        let _1041_v84;
        _1041_v84 = _dafny.Map.Empty.slice().updateUnsafe(_939_v2,_939_v2);
        let _1042_v85;
        _1042_v85 = _module.D3.create_DC8(_1041_v84);
        _1042_v85 = _module.D3.create_DC8(_1041_v84);
        let _1043_v86;
        _1043_v86 = new _dafny.CodePoint('t'.codePointAt(0));
        _1043_v86 = _1043_v86;
        let _1044_v87;
        _1044_v87 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("exys"));
        let _1045_v88;
        _1045_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(934),(_1044_v87)[_module.__default.safeIndex(_939_v2, new BigNumber((_1044_v87).length))]);
        let _1046_v89;
        _1046_v89 = _dafny.Set.fromElements(_1043_v86);
        let _1047_v90;
        _1047_v90 = _dafny.Seq.UnicodeFromString("kjjgia");
        _1045_v88 = (_1045_v88).update((_939_v2).plus(new BigNumber((_1046_v89).length)), _1047_v90);
        (globalState).f8 = _939_v2;
      }
      let _1048_v91;
      _1048_v91 = _dafny.Set.fromElements(_938_v1);
      r0 = (_1048_v91).IsDisjointFrom(_1048_v91);
      r1 = _938_v1;
      return [r0, r1];
    }
    m10(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D3.Default();
      let _1049_v0;
      _1049_v0 = false;
      let _1050_v1;
      _1050_v1 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1051_v2;
      _1051_v2 = _dafny.Seq.of(_1050_v1, _1050_v1);
      let _1052_v3;
      _1052_v3 = _module.D10.create_DC24(p2, _1049_v0, _1051_v2);
      if ((_1052_v3).dtor_cf32) {
        if ((_1049_v0) || (false)) {
          let _1053_v4;
          let _nw171 = Array((new BigNumber(5)).toNumber()).fill(false);
          _1053_v4 = _nw171;
          let _index196 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length));
          (_1053_v4)[_index196] = _1049_v0;
          let _index197 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length));
          (_1053_v4)[_index197] = _1049_v0;
          let _1054_v5;
          let _nw172 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1054_v5 = _nw172;
          let _1055_v6;
          let _nw173 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1055_v6 = _nw173;
          let _index198 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1055_v6).length));
          (_1055_v6)[_index198] = p0;
          let _1056_v7;
          _1056_v7 = _dafny.Seq.of(p0);
          let _1057_v8;
          _1057_v8 = _dafny.MultiSet.fromElements(new BigNumber(-279));
          let _index199 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length));
          let _index200 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1055_v6).length));
          let _rhs187 = (_this).fm3(_dafny.Seq.Concat(_1056_v7, _1056_v7), _1057_v8, _1049_v0, globalState);
          let _rhs188 = _1054_v5;
          let _rhs189 = (_this).fm15(new BigNumber((_dafny.Seq.Concat(_1051_v2, _dafny.Seq.UnicodeFromString("x"))).length), globalState);
          let _rhs190 = p0;
          let _lhs148 = _1053_v4;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length));
          let _lhs150 = _1055_v6;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1055_v6).length));
          let _lhs152 = globalState;
          _lhs148[_lhs149] = _rhs187;
          _1054_v5 = _rhs188;
          _lhs150[_lhs151] = _rhs189;
          _lhs152.f3 = _rhs190;
          (globalState).f0 = _1056_v7;
          let _index201 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length));
          (_1053_v4)[_index201] = (_1053_v4)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length))];
          let _1058_v9;
          _1058_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1059_v10;
          _1059_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1058_v9,(_1053_v4)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1053_v4).length))]);
          (globalState).f0 = _dafny.Seq.Concat(_1056_v7, _dafny.Seq.update(_dafny.Seq.Concat(_1056_v7, _1056_v7), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1059_v10).length)), new BigNumber((_dafny.Seq.Concat(_1056_v7, _1056_v7)).length)), p0));
        } else {
          (globalState).f8 = _module.__default.safeDivisionInt((p2).minus(p0), p0);
          (globalState).f8 = p2;
          let _1060_v11;
          _1060_v11 = _dafny.Seq.of(_1049_v0, _1049_v0, !(_1049_v0));
          let _1061_v12;
          _1061_v12 = _module.D4.create_DC10(_1060_v11);
          let _1062_v13;
          _1062_v13 = _dafny.Seq.of(_module.D4.create_DC12(_1061_v12));
          let _1063_v15;
          _1063_v15 = _module.D9.create_DC22(_1060_v11);
          let _1064_v16;
          _1064_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1063_v15)).length),new BigNumber((_1051_v2).length));
          let _1065_v17;
          _1065_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1064_v16,(_this).fm16(_1049_v0, globalState));
          let _1066_v18;
          _1066_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v13,(function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of (_1065_v17).Keys.Elements) {
              let _1067_v14 = _compr_77;
              if ((_1065_v17).contains(_1067_v14)) {
                _coll77.push([_1067_v14,p2]);
              }
            }
            return _coll77;
          }()).update(_1064_v16, p0));
          let _1068_v19;
          _1068_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1064_v16,p0);
          _1066_v18 = (_1066_v18).update(_1062_v13, _1068_v19);
          let _1069_v20;
          _1069_v20 = _dafny.MultiSet.fromElements(_module.__default.fm20(_1049_v0, globalState));
          (globalState).f2 = (_1069_v20).Difference(_1069_v20);
          let _1070_v21;
          _1070_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,new BigNumber((_1051_v2).length));
          let _1071_v22;
          _1071_v22 = _module.D12.create_DC26(_1070_v21);
          let _pat_let_tv51 = _1049_v0;
          let _pat_let_tv52 = p2;
          let _pat_let_tv53 = _1070_v21;
          let _1072_v23;
          let _nw174 = Array((new BigNumber(18)).toNumber());
          _nw174[(_dafny.ZERO).toNumber()] = _1071_v22;
          _nw174[(_dafny.ONE).toNumber()] = _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(_1049_v0,p2));
          _nw174[(new BigNumber(2)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(3)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(4)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(5)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(6)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(7)).toNumber()] = ((_1049_v0) ? (_1071_v22) : (function (_pat_let42_0) {
            return function (_1073_dt__update__tmp_h0) {
              return function (_pat_let43_0) {
                return function (_1074_dt__update_hcf35_h0) {
                  return _module.D12.create_DC26(_1074_dt__update_hcf35_h0);
                }(_pat_let43_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv51,_pat_let_tv52));
            }(_pat_let42_0);
          }(_1071_v22)));
          _nw174[(new BigNumber(8)).toNumber()] = function (_pat_let44_0) {
            return function (_1075_dt__update__tmp_h1) {
              return function (_pat_let45_0) {
                return function (_1076_dt__update_hcf35_h1) {
                  return _module.D12.create_DC26(_1076_dt__update_hcf35_h1);
                }(_pat_let45_0);
              }(_pat_let_tv53);
            }(_pat_let44_0);
          }(_module.D12.create_DC26(_1070_v21));
          _nw174[(new BigNumber(9)).toNumber()] = _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(_1049_v0,p2));
          _nw174[(new BigNumber(10)).toNumber()] = _module.__default.fm39(p2, _1049_v0, _1049_v0, globalState);
          _nw174[(new BigNumber(11)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(12)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(13)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(14)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(15)).toNumber()] = _module.__default.fm39(new BigNumber(616), _1049_v0, _1049_v0, globalState);
          _nw174[(new BigNumber(16)).toNumber()] = _1071_v22;
          _nw174[(new BigNumber(17)).toNumber()] = _1071_v22;
          _1072_v23 = _nw174;
          let _index202 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_1072_v23).length));
          (_1072_v23)[_index202] = _1071_v22;
          let _index203 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_1072_v23).length));
          let _rhs191 = (p0).plus(new BigNumber((_1051_v2).length));
          let _rhs192 = _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(_1049_v0,p2));
          let _rhs193 = new BigNumber((_1051_v2).length);
          let _lhs153 = globalState;
          let _lhs154 = _1072_v23;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_1072_v23).length));
          let _lhs156 = globalState;
          _lhs153.f3 = _rhs191;
          _lhs154[_lhs155] = _rhs192;
          _lhs156.f8 = _rhs193;
        }
        _1049_v0 = _1049_v0;
        let _1077_v24;
        _1077_v24 = _dafny.MultiSet.fromElements(false, _1049_v0, _1049_v0, _1049_v0);
        let _1078_v25;
        _1078_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm28(p0, globalState),(_1077_v24).IsDisjointFrom(_1077_v24));
        let _1079_v26;
        _1079_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,new BigNumber(373));
        let _1080_v27;
        let _nw175 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1080_v27 = _nw175;
        let _1081_v28;
        _1081_v28 = _dafny.Set.fromElements(_1080_v27, _1080_v27);
        let _1082_v29;
        _1082_v29 = _dafny.Set.fromElements(new BigNumber(869), new BigNumber(951), p2, new BigNumber((_1079_v26).length), new BigNumber((_1081_v28).length));
        if ((((_1078_v25).contains(_1082_v29)) ? ((_1078_v25).get(_1082_v29)) : (_1049_v0))) {
          let _1083_v30;
          let _nw176 = new _module.C2();
          _nw176.__ctor();
          _1083_v30 = _nw176;
          let _1084_v31;
          _1084_v31 = _dafny.Seq.of(new BigNumber(659), p0);
          let _1085_v32;
          _1085_v32 = _dafny.MultiSet.fromElements(p0);
          _1049_v0 = (_this).fm3(_dafny.Seq.Concat(_1084_v31, _1084_v31), _1085_v32, _1049_v0, globalState);
          let _1086_v33;
          _1086_v33 = _dafny.Set.fromElements(_1049_v0, _1049_v0);
          let _1087_v34;
          _1087_v34 = _module.D9.create_DC20(_1086_v33);
          let _rhs194 = (_1049_v0) || (_1049_v0);
          let _rhs195 = ((_1087_v34).dtor_cf25).IsProperSubsetOf(_1086_v33);
          let _lhs157 = globalState;
          _1049_v0 = _rhs194;
          _lhs157.f4 = _rhs195;
          let _1088_v35;
          _1088_v35 = _1085_v32;
          let _1089_v36;
          let _nw177 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1089_v36 = _nw177;
          let _index204 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length));
          (_1089_v36)[_index204] = (_1082_v29).IsSubsetOf(_1082_v29);
          let _index205 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length));
          let _rhs196 = _1088_v35;
          let _rhs197 = false;
          let _rhs198 = _1049_v0;
          let _rhs199 = _1049_v0;
          let _lhs158 = _1089_v36;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length));
          _1088_v35 = _rhs196;
          _1049_v0 = _rhs197;
          _1049_v0 = _rhs198;
          _lhs158[_lhs159] = _rhs199;
          let _1090_v37;
          _1090_v37 = _dafny.Seq.of((_1089_v36)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length))]);
          let _1091_v38;
          _1091_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          _1090_v37 = _dafny.Seq.Concat(_dafny.Seq.of(_1049_v0), _dafny.Seq.Concat(_dafny.Seq.update(_1090_v37, _module.__default.safeIndex(new BigNumber((_module.__default.fm1(p0, _1050_v1, (_1089_v36)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length))], _1091_v38, globalState)).length), new BigNumber((_1090_v37).length)), (_1089_v36)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_1089_v36).length))]), _module.__default.fm31(true, false, _1049_v0, globalState)));
        } else {
          _1049_v0 = (_1077_v24).IsDisjointFrom((_1077_v24).Difference(_1077_v24));
          let _1092_v39;
          let _nw178 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1092_v39 = _nw178;
          let _1093_v40;
          _1093_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,_1092_v39);
          let _1094_v41;
          _1094_v41 = _dafny.Seq.of(_1092_v39);
          _1093_v40 = (_1093_v40).update(_1049_v0, (_1094_v41)[_module.__default.safeIndex(p2, new BigNumber((_1094_v41).length))]);
          (globalState).f4 = _1049_v0;
          _1049_v0 = !(new BigNumber(66)).isEqualTo(p2);
          let _1095_v42;
          _1095_v42 = _dafny.Set.fromElements(_1049_v0);
          let _1096_v43;
          let _nw179 = Array((new BigNumber(3)).toNumber());
          _nw179[(_dafny.ZERO).toNumber()] = _1095_v42;
          _nw179[(_dafny.ONE).toNumber()] = _1095_v42;
          _nw179[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_1049_v0);
          _1096_v43 = _nw179;
          let _1097_v44;
          _1097_v44 = _module.D4.create_DC11(true, _1049_v0, new BigNumber(-888), _1051_v2);
          let _1098_v45;
          _1098_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1096_v43,new BigNumber((_dafny.Seq.of(false, (_1097_v44).dtor_cf14, _1049_v0, _1049_v0)).length));
          _1098_v45 = (_1098_v45).update(_1096_v43, p2);
        }
        (globalState).f3 = p2;
        let _1099_v46;
        _1099_v46 = _module.D5.create_DC15();
        let _source19 = _1099_v46;
        if (_source19.is_DC14) {
          let _1100___mcc_h0 = (_source19).cf20;
          let _1101___mcc_h1 = (_source19).cf21;
          let _1102_cf21 = _1101___mcc_h1;
          let _1103_cf20 = _1100___mcc_h0;
          let _1104_v47;
          _1104_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1102_cf21,_1049_v0);
          let _1105_v48;
          _1105_v48 = _dafny.Seq.of(p2, (_dafny.ZERO).minus(p2));
          let _1106_v50;
          _1106_v50 = _module.D12.create_DC27(new BigNumber((_1051_v2).length), new BigNumber((_1104_v47).length), p0, new BigNumber((function () {
  let _coll78 = new _dafny.Set();
  for (const _compr_78 of (_1105_v48).Elements) {
    let _1107_v49 = _compr_78;
    if (_dafny.Seq.contains(_1105_v48, _1107_v49)) {
      _coll78.add((_1107_v49).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())));
    }
  }
  return _coll78;
}()).length));
          let _1108_v51;
          _1108_v51 = _module.D12.create_DC28(_1106_v50);
          let _1109_v52;
          _1109_v52 = _dafny.Set.fromElements(_1102_cf21);
          let _rhs200 = (_module.__default.safeModuloInt(p0, p0)).plus(_1103_cf20);
          let _rhs201 = _module.__default.safeModuloInt(new BigNumber((_1105_v48).length), _module.__default.safeDivisionInt(new BigNumber((_1109_v52).length), new BigNumber((_1077_v24).cardinality())));
          let _rhs202 = _module.__default.fm0(p2, globalState);
          let _rhs203 = (((new BigNumber(15)).isLessThan((_dafny.ZERO).minus((((_1079_v26).contains(_1049_v0)) ? ((_1079_v26).get(_1049_v0)) : (new BigNumber(461)))))) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(745)), ((_1110_v1) => function (_1111_i0) {
            return _1110_v1;
          })(_1050_v1))) : (_1051_v2));
          let _rhs204 = _1108_v51;
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          _lhs160.f3 = _rhs200;
          _lhs161.f8 = _rhs201;
          _1077_v24 = _rhs202;
          _1051_v2 = _rhs203;
          _1108_v51 = _rhs204;
          _1104_v47 = (_1104_v47).update(!(_1102_cf21) || (_1102_cf21), _1049_v0);
          (globalState).f8 = _module.__default.safeModuloInt(_1103_cf20, _module.__default.safeModuloInt(_1103_cf20, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length))));
          let _1112_v53;
          let _nw180 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1112_v53 = _nw180;
          let _1113_v54;
          _1113_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,_1112_v53);
          _1113_v54 = (_1113_v54).update(_1049_v0, _1112_v53);
        } else if (_source19.is_DC15) {
          let _index206 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1080_v27).length));
          (_1080_v27)[_index206] = p2;
          let _1114_v55;
          let _nw181 = Array((new BigNumber(4)).toNumber()).fill(false);
          _1114_v55 = _nw181;
          let _1115_v56;
          _1115_v56 = _dafny.Seq.of(_1114_v55);
          let _1116_v57;
          _1116_v57 = _dafny.Seq.of(_1114_v55, _1114_v55);
          let _1117_v58;
          _1117_v58 = _module.D2.create_DC7(_1050_v1, (_1049_v0) || (_1049_v0), new BigNumber((_dafny.Seq.Concat(_1115_v56, (_1116_v57))).length));
          let _index207 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1080_v27).length));
          let _rhs205 = p0;
          let _rhs206 = _dafny.Seq.Concat(_1051_v2, _dafny.Seq.Concat(_1051_v2, _1051_v2));
          let _rhs207 = ((_1049_v0) ? (_module.__default.fm40(!(_1049_v0), globalState)) : (_1117_v58));
          let _lhs162 = _1080_v27;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1080_v27).length));
          _lhs162[_lhs163] = _rhs205;
          _1051_v2 = _rhs206;
          _1117_v58 = _rhs207;
          let _1118_v59;
          _1118_v59 = _dafny.Seq.of((_1080_v27)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_1080_v27).length))]);
          let _1119_v60;
          let _nw182 = Array((new BigNumber(3)).toNumber());
          _nw182[(_dafny.ZERO).toNumber()] = _1118_v59;
          _nw182[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), ((_1120_p0) => function (_1121_i1) {
            return _1120_p0;
          })(p0));
          _nw182[(new BigNumber(2)).toNumber()] = _1118_v59;
          _1119_v60 = _nw182;
          let _index208 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1119_v60).length));
          (_1119_v60)[_index208] = _1118_v59;
          let _index209 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1119_v60).length));
          (_1119_v60)[_index209] = _1118_v59;
          let _1122_v61;
          _1122_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,_1118_v59);
          let _index210 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1119_v60).length));
          let _rhs208 = ((_1080_v27)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_1080_v27).length))]).isLessThanOrEqualTo(p0);
          let _rhs209 = _1050_v1;
          let _rhs210 = p0;
          let _rhs211 = (((_1122_v61).contains(_1049_v0)) ? ((_1122_v61).get(_1049_v0)) : ((_1119_v60)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_1119_v60).length))]));
          let _lhs164 = globalState;
          let _lhs165 = globalState;
          let _lhs166 = _1119_v60;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1119_v60).length));
          _lhs164.f4 = _rhs208;
          _1050_v1 = _rhs209;
          _lhs165.f3 = _rhs210;
          _lhs166[_lhs167] = _rhs211;
          let _1123_v62;
          let _nw183 = new _module.C1();
          _nw183.__ctor();
          _1123_v62 = _nw183;
        } else {
          let _1124___mcc_h2 = (_source19).cf19;
          let _1125_cf19 = _1124___mcc_h2;
          let _1126_v63;
          let _nw184 = Array((new BigNumber(13)).toNumber()).fill(_module.D12.Default());
          _1126_v63 = _nw184;
          let _1127_v64;
          _1127_v64 = _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe(false,p0));
          let _index211 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1126_v63).length));
          (_1126_v63)[_index211] = _1127_v64;
          let _index212 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1126_v63).length));
          (_1126_v63)[_index212] = _1127_v64;
          _1049_v0 = _1049_v0;
          let _1128_v65;
          let _nw185 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1128_v65 = _nw185;
          let _index213 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_1128_v65).length));
          (_1128_v65)[_index213] = _1049_v0;
          let _index214 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_1128_v65).length));
          (_1128_v65)[_index214] = (_1049_v0) && (_1049_v0);
          let _rhs212 = p0;
          let _rhs213 = (_1128_v65)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_1128_v65).length))];
          let _rhs214 = (new BigNumber((p1).length)).isLessThanOrEqualTo(p0);
          let _rhs215 = (_dafny.ZERO).minus(p0);
          let _lhs168 = globalState;
          let _lhs169 = globalState;
          _lhs168.f3 = _rhs212;
          _1049_v0 = _rhs213;
          _1049_v0 = _rhs214;
          _lhs169.f8 = _rhs215;
        }
      } else {
        let _1129_v66;
        _1129_v66 = _dafny.Map.Empty.slice().updateUnsafe(!(_1049_v0),p2);
        _1129_v66 = (_1129_v66).update(_1049_v0, p0);
        let _1130_v67;
        _1130_v67 = _dafny.Set.fromElements(_1049_v0, !((((p1).contains(p0)) ? ((p1).get(p0)) : (_1049_v0))), false, _1049_v0);
        let _1131_v68;
        _1131_v68 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1049_v0, _1049_v0, _1049_v0, _1049_v0),(_1130_v67).IsProperSubsetOf(_1130_v67));
        let _1132_v69;
        _1132_v69 = _module.D9.create_DC20(_1130_v67);
        _1131_v68 = (_1131_v68).update((_1132_v69).dtor_cf25, _1049_v0);
        let _1133_v70;
        _1133_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1051_v2).length),(_this).fm15(p0, globalState));
        let _1134_v71;
        _1134_v71 = _dafny.Seq.of(p0);
        let _1135_v72;
        let _nw186 = Array((new BigNumber(14)).toNumber()).fill(_module.D2.Default());
        _1135_v72 = _nw186;
        let _1136_v73;
        _1136_v73 = _dafny.Set.fromElements(_1135_v72);
        let _1137_v74;
        let _nw187 = Array((new BigNumber(25)).toNumber());
        _nw187[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1049_v0)).length);
        _nw187[(_dafny.ONE).toNumber()] = p0;
        _nw187[(new BigNumber(2)).toNumber()] = p0;
        _nw187[(new BigNumber(3)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p2)).length))).cardinality())).plus(p2);
        _nw187[(new BigNumber(4)).toNumber()] = (new BigNumber(-628)).plus(p2);
        _nw187[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p2), p2)).length);
        _nw187[(new BigNumber(6)).toNumber()] = p0;
        _nw187[(new BigNumber(7)).toNumber()] = p0;
        _nw187[(new BigNumber(8)).toNumber()] = p0;
        _nw187[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw187[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1050_v1, _1050_v1, new _dafny.CodePoint('r'.codePointAt(0)))).length);
        _nw187[(new BigNumber(11)).toNumber()] = (((_1133_v70).contains(p0)) ? ((_1133_v70).get(p0)) : (p2));
        _nw187[(new BigNumber(12)).toNumber()] = new BigNumber(-132);
        _nw187[(new BigNumber(13)).toNumber()] = p0;
        _nw187[(new BigNumber(14)).toNumber()] = p2;
        _nw187[(new BigNumber(15)).toNumber()] = (p0).multipliedBy(p2);
        _nw187[(new BigNumber(16)).toNumber()] = p0;
        _nw187[(new BigNumber(17)).toNumber()] = ((_1049_v0) ? (p2) : (new BigNumber(319)));
        _nw187[(new BigNumber(18)).toNumber()] = new BigNumber(607);
        _nw187[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus((p0).plus(p0));
        _nw187[(new BigNumber(20)).toNumber()] = p2;
        _nw187[(new BigNumber(21)).toNumber()] = p0;
        _nw187[(new BigNumber(22)).toNumber()] = new BigNumber((_1134_v71).length);
        _nw187[(new BigNumber(23)).toNumber()] = new BigNumber((_1136_v73).length);
        _nw187[(new BigNumber(24)).toNumber()] = p2;
        _1137_v74 = _nw187;
        _1137_v74 = _1137_v74;
        let _1138_v75;
        let _nw188 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1138_v75 = _nw188;
        let _index215 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length));
        (_1138_v75)[_index215] = _1049_v0;
        let _1139_v76;
        _1139_v76 = _dafny.MultiSet.fromElements(_1135_v72);
        let _1140_v77;
        _1140_v77 = _dafny.Set.fromElements(p0);
        let _index216 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length));
        (_1138_v75)[_index216] = !(_1140_v77).contains(new BigNumber(((_1139_v76).Intersect((_1139_v76))).cardinality()));
        if (_1049_v0) {
          let _1141_v78;
          let _nw189 = new _module.C0();
          _nw189.__ctor();
          _1141_v78 = _nw189;
          _1133_v70 = (_1133_v70).update((new BigNumber((_1130_v67).length)).plus(p2), p0);
          let _1142_v79;
          _1142_v79 = _dafny.Seq.of(false);
          _1049_v0 = (_1142_v79)[_module.__default.safeIndex(p2, new BigNumber((_1142_v79).length))];
          (globalState).f4 = (_1142_v79)[_module.__default.safeIndex(p0, new BigNumber((_1142_v79).length))];
          let _1143_v80;
          _1143_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))],(_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))]);
          let _rhs216 = p2;
          let _rhs217 = (_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))];
          let _rhs218 = (_module.D13.create_DC30((((_1143_v80).contains(_1049_v0)) ? ((_1143_v80).get(_1049_v0)) : ((_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))])), _1049_v0, _1140_v77, _1141_v78)).dtor_cf44;
          let _lhs170 = globalState;
          let _lhs171 = globalState;
          _lhs170.f3 = _rhs216;
          _lhs171.f4 = _rhs217;
          _1140_v77 = _rhs218;
        } else {
          _1133_v70 = (_1133_v70).Merge(_1133_v70);
          let _1144_v81;
          _1144_v81 = _module.__default.fm0((_dafny.ZERO).minus(p2), globalState);
          let _1145_v82;
          _1145_v82 = _dafny.MultiSet.fromElements((_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))]);
          _1144_v81 = (_1145_v82).Difference(_1145_v82);
          let _1146_v83;
          let _nw190 = new _module.C2();
          _nw190.__ctor();
          _1146_v83 = _nw190;
          let _1147_v84;
          _1147_v84 = _dafny.Seq.of((_this).fm16(false, globalState));
          let _1148_v85;
          _1148_v85 = _dafny.Seq.of(_1049_v0, (((_1147_v84)[_module.__default.safeIndex(_module.__default.fm20((_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))], globalState), new BigNumber((_1147_v84).length))]) ? (true) : ((_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))])), (_1138_v75)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_1138_v75).length))]);
          _1049_v0 = (_1147_v84)[_module.__default.safeIndex(p2, new BigNumber((_1147_v84).length))];
          _1051_v2 = _1051_v2;
        }
      }
      let _1149_v86;
      _1149_v86 = _dafny.Seq.of(_1049_v0);
      let _pat_let_tv54 = _1051_v2;
      let _pat_let_tv55 = _1149_v86;
      let _source20 = function (_pat_let46_0) {
        return function (_1150_dt__update__tmp_h2) {
          return function (_pat_let47_0) {
            return function (_1151_dt__update_hcf17_h0) {
              return function (_pat_let48_0) {
                return function (_1152_dt__update_hcf16_h0) {
                  return _module.D4.create_DC11((_1150_dt__update__tmp_h2).dtor_cf14, (_1150_dt__update__tmp_h2).dtor_cf15, _1152_dt__update_hcf16_h0, _1151_dt__update_hcf17_h0);
                }(_pat_let48_0);
              }(new BigNumber((_pat_let_tv55).length));
            }(_pat_let47_0);
          }(_pat_let_tv54);
        }(_pat_let46_0);
      }(_module.D4.create_DC11(_1049_v0, _1049_v0, p0, _1051_v2));
      if (_source20.is_DC11) {
        let _1153___mcc_h3 = (_source20).cf14;
        let _1154___mcc_h4 = (_source20).cf15;
        let _1155___mcc_h5 = (_source20).cf16;
        let _1156___mcc_h6 = (_source20).cf17;
        let _1157_cf17 = _1156___mcc_h6;
        let _1158_cf16 = _1155___mcc_h5;
        let _1159_cf15 = _1154___mcc_h4;
        let _1160_cf14 = _1153___mcc_h3;
        let _1161_v87;
        let _init27 = ((_1162_cf17) => function (_1163_i2) {
          return _1162_cf17;
        })(_1157_cf17);
        let _nw191 = Array((new BigNumber(5)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw191.length); _i0_27++) {
          _nw191[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1161_v87 = _nw191;
        let _index217 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1161_v87).length));
        (_1161_v87)[_index217] = _1051_v2;
        let _index218 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_1161_v87).length));
        (_1161_v87)[_index218] = _module.__default.fm25(globalState);
        let _1164_v88;
        _1164_v88 = _module.D9.create_DC22(_1149_v86);
        let _1165_v89;
        _1165_v89 = _dafny.Seq.of((_dafny.ZERO).minus(p0));
        let _1166_v90;
        _1166_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1158_cf16,p0);
        let _1167_v91;
        _1167_v91 = _dafny.Seq.of(_dafny.Seq.of(_1159_cf15));
        let _1168_v92;
        _1168_v92 = _dafny.MultiSet.fromElements(_1149_v86, (_1167_v91)[_module.__default.safeIndex(_module.__default.fm20(_1160_cf14, globalState), new BigNumber((_1167_v91).length))], _1149_v86);
        let _pat_let_tv56 = _1166_v90;
        let _rhs219 = _1165_v89;
        let _rhs220 = _module.__default.fm26(function (_pat_let49_0) {
          return function (_1169_dt__update__tmp_h3) {
            return function (_pat_let50_0) {
              return function (_1170_dt__update_hcf12_h0) {
                return _module.D3.create_DC8(_1170_dt__update_hcf12_h0);
              }(_pat_let50_0);
            }(_pat_let_tv56);
          }(_pat_let49_0);
        }(_module.D3.create_DC8(_1166_v90)), (_1158_cf16).plus(p0), _1168_v92, globalState);
        let _rhs221 = _1164_v88;
        let _lhs172 = globalState;
        let _lhs173 = globalState;
        _lhs172.f0 = _rhs219;
        _lhs173.f8 = _rhs220;
        _1164_v88 = _rhs221;
        let _1171_v93;
        let _nw192 = Array((new BigNumber(9)).toNumber()).fill(false);
        _1171_v93 = _nw192;
        let _1172_v94;
        _1172_v94 = _module.D4.create_DC11(_1160_cf14, _1159_cf15, new BigNumber((_1157_cf17).length), _1051_v2);
        let _index219 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1171_v93).length));
        (_1171_v93)[_index219] = (_1172_v94).dtor_cf14;
        let _index220 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1171_v93).length));
        (_1171_v93)[_index220] = ((_1049_v0) ? (!(_1160_cf14)) : (!(_dafny.Seq.IsPrefixOf(_1149_v86, _1149_v86))));
        let _1173_v95;
        _1173_v95 = _dafny.MultiSet.fromElements(_1049_v0);
        let _1174_v96;
        let _nw193 = new _module.C2();
        _nw193.__ctor();
        _1174_v96 = _nw193;
        let _1175_v97;
        _1175_v97 = _module.D16.create_DC34(_1174_v96);
        let _1176_v98;
        _1176_v98 = _dafny.Map.Empty.slice().updateUnsafe((_1175_v97).dtor_cf49,true);
        let _1177_v99;
        let _nw194 = new _module.C1();
        _nw194.__ctor();
        _1177_v99 = _nw194;
        let _1178_v100;
        _1178_v100 = _dafny.MultiSet.fromElements(p0);
        let _1179_v101;
        _1179_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1177_v99,_1178_v100);
        let _1180_v102;
        let _nw195 = Array((new BigNumber(27)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw195[(_dafny.ONE).toNumber()] = p2;
        _nw195[(new BigNumber(2)).toNumber()] = (((_1173_v95).contains(false)) ? ((_1173_v95).get(false)) : (new BigNumber((_1176_v98).length)));
        _nw195[(new BigNumber(3)).toNumber()] = _1158_cf16;
        _nw195[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(p0, p0)).length);
        _nw195[(new BigNumber(5)).toNumber()] = new BigNumber(141);
        _nw195[(new BigNumber(6)).toNumber()] = _1158_cf16;
        _nw195[(new BigNumber(7)).toNumber()] = (p2).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_1160_cf14)).cardinality()));
        _nw195[(new BigNumber(8)).toNumber()] = (p2).plus(p2);
        _nw195[(new BigNumber(9)).toNumber()] = p0;
        _nw195[(new BigNumber(10)).toNumber()] = p0;
        _nw195[(new BigNumber(11)).toNumber()] = (_1158_cf16).plus(_1158_cf16);
        _nw195[(new BigNumber(12)).toNumber()] = new BigNumber(-851);
        _nw195[(new BigNumber(13)).toNumber()] = new BigNumber(((((_1179_v101).contains(_1177_v99)) ? ((_1179_v101).get(_1177_v99)) : (_1178_v100))).cardinality());
        _nw195[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        _nw195[(new BigNumber(15)).toNumber()] = _1158_cf16;
        _nw195[(new BigNumber(16)).toNumber()] = (new BigNumber(-421)).minus(p2);
        _nw195[(new BigNumber(17)).toNumber()] = p0;
        _nw195[(new BigNumber(18)).toNumber()] = (new BigNumber((_1166_v90).length)).plus(new BigNumber(-24));
        _nw195[(new BigNumber(19)).toNumber()] = new BigNumber((_1165_v89).length);
        _nw195[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_1158_cf16).multipliedBy(p0));
        _nw195[(new BigNumber(21)).toNumber()] = p0;
        _nw195[(new BigNumber(22)).toNumber()] = p0;
        _nw195[(new BigNumber(23)).toNumber()] = _1158_cf16;
        _nw195[(new BigNumber(24)).toNumber()] = p2;
        _nw195[(new BigNumber(25)).toNumber()] = p2;
        _nw195[(new BigNumber(26)).toNumber()] = _1158_cf16;
        _1180_v102 = _nw195;
        let _index221 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1180_v102).length));
        (_1180_v102)[_index221] = (_1165_v89)[_module.__default.safeIndex(p0, new BigNumber((_1165_v89).length))];
        let _index222 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1180_v102).length));
        (_1180_v102)[_index222] = (p2).plus(p0);
      } else if (_source20.is_DC10) {
        let _1181___mcc_h7 = (_source20).cf13;
        let _1182_cf13 = _1181___mcc_h7;
        let _1183_v103;
        let _nw196 = new _module.C0();
        _nw196.__ctor();
        _1183_v103 = _nw196;
        let _1184_v104;
        _1184_v104 = _module.D0.create_DC2();
        let _source21 = _1184_v104;
        if (_source21.is_DC1) {
          let _1185___mcc_h9 = (_source21).cf1;
          let _1186___mcc_h10 = (_source21).cf2;
          let _1187_cf2 = _1186___mcc_h10;
          let _1188_cf1 = _1185___mcc_h9;
          let _1189_v105;
          let _nw197 = Array((new BigNumber(13)).toNumber());
          _1189_v105 = _nw197;
          let _1190_v106;
          _1190_v106 = _dafny.Set.fromElements(new BigNumber(274), _1187_cf2, p2, new BigNumber(845), _1187_cf2);
          let _1191_v107;
          _1191_v107 = _module.D13.create_DC30(false, _1049_v0, _1190_v106, _1183_v103);
          let _index223 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1189_v105).length));
          (_1189_v105)[_index223] = _1191_v107;
          let _1192_v108;
          _1192_v108 = _dafny.Seq.of(_1183_v103, _1183_v103);
          let _index224 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1189_v105).length));
          (_1189_v105)[_index224] = _module.D13.create_DC30(_1049_v0, _1049_v0, _dafny.Set.fromElements(new BigNumber(-260), p2), (_1192_v108)[_module.__default.safeIndex(_1187_cf2, new BigNumber((_1192_v108).length))]);
          let _1193_v109;
          let _nw198 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1193_v109 = _nw198;
          let _index225 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1193_v109).length));
          (_1193_v109)[_index225] = _1049_v0;
          let _index226 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1193_v109).length));
          (_1193_v109)[_index226] = !_dafny.Seq.contains(_1051_v2, _1050_v1);
          let _1194_v110;
          _1194_v110 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(812));
          let _1195_v111;
          _1195_v111 = _module.D3.create_DC8(_1194_v110);
          let _1196_v112;
          _1196_v112 = _dafny.MultiSet.fromElements(_1149_v86);
          let _pat_let_tv57 = _1194_v110;
          (globalState).f8 = _module.__default.fm26(function (_pat_let51_0) {
            return function (_1197_dt__update__tmp_h4) {
              return function (_pat_let52_0) {
                return function (_1198_dt__update_hcf12_h1) {
                  return _module.D3.create_DC8(_1198_dt__update_hcf12_h1);
                }(_pat_let52_0);
              }(_pat_let_tv57);
            }(_pat_let51_0);
          }(_1195_v111), new BigNumber((_1051_v2).length), _1196_v112, globalState);
          let _1199_v113;
          _1199_v113 = !((_1193_v109)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_1193_v109).length))]);
          let _1200_v114;
          _1200_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1199_v113,(_1193_v109)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_1193_v109).length))]);
          let _1201_v115;
          _1201_v115 = _dafny.Seq.of(new BigNumber((_1051_v2).length), p2, _1188_cf1);
          _1200_v114 = (_1200_v114).update(_1199_v113, _dafny.Seq.IsPrefixOf(_1201_v115, _dafny.Seq.of(new BigNumber((_1182_cf13).length), p2)));
        } else if (_source21.is_DC2) {
          let _1202_v116;
          let _nw199 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1202_v116 = _nw199;
          let _1203_v117;
          let _nw200 = Array((new BigNumber(5)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1202_v116;
          _nw200[(_dafny.ONE).toNumber()] = _1202_v116;
          _nw200[(new BigNumber(2)).toNumber()] = _1202_v116;
          _nw200[(new BigNumber(3)).toNumber()] = _1202_v116;
          _nw200[(new BigNumber(4)).toNumber()] = _1202_v116;
          _1203_v117 = _nw200;
          let _index227 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1203_v117).length));
          (_1203_v117)[_index227] = _1202_v116;
          let _1204_v118;
          _1204_v118 = _dafny.Seq.of(p0);
          let _1205_v119;
          _1205_v119 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p0).multipliedBy((_1204_v118)[_module.__default.safeIndex(new BigNumber((_1051_v2).length), new BigNumber((_1204_v118).length))]));
          let _1206_v120;
          let _init28 = ((_1207_v1) => function (_1208_i3) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), ((_1209_v1) => function (_1210_i4) {
              return _1209_v1;
            })(_1207_v1)), _dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), _1207_v1, _1207_v1, new _dafny.CodePoint('f'.codePointAt(0))));
          })(_1050_v1);
          let _nw201 = Array((new BigNumber(24)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw201.length); _i0_28++) {
            _nw201[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1206_v120 = _nw201;
          let _index228 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_1206_v120).length));
          (_1206_v120)[_index228] = _1051_v2;
          let _1211_v121;
          _1211_v121 = _1202_v116;
          let _index229 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1203_v117).length));
          let _index230 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_1206_v120).length));
          let _rhs222 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_module.__default.fm1(p0, _1050_v1, _1049_v0, _1205_v119, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm1(p0, _1050_v1, _1049_v0, _1205_v119, globalState)).length)), new BigNumber(973))).length));
          let _rhs223 = (new BigNumber((_1051_v2).length)).minus(new BigNumber((_1149_v86).length));
          let _rhs224 = (_1211_v121);
          let _rhs225 = _1205_v119;
          let _rhs226 = _1051_v2;
          let _lhs174 = globalState;
          let _lhs175 = globalState;
          let _lhs176 = _1203_v117;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1203_v117).length));
          let _lhs178 = _1206_v120;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_1206_v120).length));
          _lhs174.f3 = _rhs222;
          _lhs175.f8 = _rhs223;
          _lhs176[_lhs177] = _rhs224;
          _1205_v119 = _rhs225;
          _lhs178[_lhs179] = _rhs226;
          let _1212_v122;
          let _nw202 = Array((new BigNumber(13)).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = _1049_v0;
          _nw202[(_dafny.ONE).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(2)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(3)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(4)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(5)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(6)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(7)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(8)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(9)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(10)).toNumber()] = (_this).fm16(_1049_v0, globalState);
          _nw202[(new BigNumber(11)).toNumber()] = _1049_v0;
          _nw202[(new BigNumber(12)).toNumber()] = _1049_v0;
          _1212_v122 = _nw202;
          let _1213_v123;
          _1213_v123 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1212_v122);
          _1213_v123 = (_1213_v123).Merge(_1213_v123);
          _1051_v2 = (_1206_v120)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_1206_v120).length))];
          (globalState).f4 = _1049_v0;
        } else if (_source21.is_DC0) {
          let _1214___mcc_h11 = (_source21).cf0;
          let _1215_cf0 = _1214___mcc_h11;
          (globalState).f4 = _1049_v0;
          let _1216_v124;
          let _init29 = ((_1217_v2) => function (_1218_i5) {
            return _1217_v2;
          })(_1051_v2);
          let _nw203 = Array((new BigNumber(16)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw203.length); _i0_29++) {
            _nw203[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1216_v124 = _nw203;
          let _index231 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1216_v124).length));
          (_1216_v124)[_index231] = _dafny.Seq.UnicodeFromString("rwhy");
          let _1219_v125;
          _1219_v125 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,_1050_v1);
          let _1220_v126;
          _1220_v126 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v2,!(_1219_v125).contains(_1049_v0));
          let _1221_v127;
          let _nw204 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1221_v127 = _nw204;
          let _index232 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length));
          (_1221_v127)[_index232] = p0;
          let _1222_v128;
          _1222_v128 = _1221_v127;
          let _1223_v129;
          _1223_v129 = _dafny.Seq.of(_1222_v128);
          let _index233 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1216_v124).length));
          let _index234 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length));
          let _rhs227 = _1051_v2;
          let _rhs228 = _1220_v126;
          let _rhs229 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1149_v86, _dafny.Seq.Concat(_1182_cf13, _module.__default.fm31(_1049_v0, _1049_v0, _1049_v0, globalState))), _module.__default.safeIndex(new BigNumber((((_1049_v0) ? (_1223_v129) : (_1223_v129))).length), new BigNumber((_dafny.Seq.Concat(_1149_v86, _dafny.Seq.Concat(_1182_cf13, _module.__default.fm31(_1049_v0, _1049_v0, _1049_v0, globalState)))).length)), _1049_v0)).length);
          let _rhs230 = _1049_v0;
          let _lhs180 = _1216_v124;
          let _lhs181 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1216_v124).length));
          let _lhs182 = _1221_v127;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length));
          let _lhs184 = globalState;
          _lhs180[_lhs181] = _rhs227;
          _1220_v126 = _rhs228;
          _lhs182[_lhs183] = _rhs229;
          _lhs184.f4 = _rhs230;
          (globalState).f4 = _1049_v0;
          let _1224_v130;
          _1224_v130 = _dafny.Seq.of(_1215_cf0, (_1221_v127)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length))]);
          let _1225_v131;
          _1225_v131 = _dafny.MultiSet.fromElements(_1215_cf0);
          let _1226_v132;
          _1226_v132 = _dafny.MultiSet.fromElements(_1225_v131);
          let _1227_v133;
          _1227_v133 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-174)), ((_1228_p0) => function (_1229_i6) {
            return _1228_p0;
          })(p0)),(_dafny.ZERO).minus(_1215_cf0));
          let _1230_v134;
          _1230_v134 = _dafny.Map.Empty.slice().updateUnsafe(false,_1049_v0);
          let _1231_v135;
          _1231_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1049_v0,(((_1230_v134).contains(_1049_v0)) ? ((_1230_v134).get(_1049_v0)) : ((_this).fm3(_1224_v130, _1225_v131, _1049_v0, globalState))));
          let _1232_v136;
          _1232_v136 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1231_v135).length));
          let _1233_v137;
          _1233_v137 = _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_1215_cf0,p2)).update(new BigNumber((_1227_v133).length), p0), _1232_v136);
          let _1234_v139;
          let _nw205 = Array((new BigNumber(20)).toNumber());
          _nw205[(_dafny.ZERO).toNumber()] = _1049_v0;
          _nw205[(_dafny.ONE).toNumber()] = !(_1049_v0);
          _nw205[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1224_v130, _dafny.Seq.of((_1221_v127)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length))]));
          _nw205[(new BigNumber(3)).toNumber()] = (_1226_v132).IsDisjointFrom(_1226_v132);
          _nw205[(new BigNumber(4)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(5)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(6)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements(function () {
            let _coll79 = new _dafny.Map();
            for (const _compr_79 of _dafny.IntegerRange(new BigNumber(-765), new BigNumber(433))) {
              let _1235_v138 = _compr_79;
              if (((new BigNumber(-765)).isLessThanOrEqualTo(_1235_v138)) && ((_1235_v138).isLessThan(new BigNumber(433)))) {
                _coll79.push([(_1235_v138).minus((_1221_v127)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length))]),p0]);
              }
            }
            return _coll79;
          }(), _1232_v136)).IsSubsetOf(_1233_v137);
          _nw205[(new BigNumber(8)).toNumber()] = true;
          _nw205[(new BigNumber(9)).toNumber()] = !(!(_1049_v0));
          _nw205[(new BigNumber(10)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(11)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(12)).toNumber()] = (new BigNumber((_dafny.Seq.of(_1215_cf0, p0)).length)).isLessThanOrEqualTo((_1221_v127)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_1221_v127).length))]);
          _nw205[(new BigNumber(13)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(14)).toNumber()] = true;
          _nw205[(new BigNumber(15)).toNumber()] = !(!(_1049_v0));
          _nw205[(new BigNumber(16)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(17)).toNumber()] = _1049_v0;
          _nw205[(new BigNumber(18)).toNumber()] = (_dafny.Set.fromElements(_1215_cf0)).IsDisjointFrom(_dafny.Set.fromElements(p2, p2, new BigNumber(((_1216_v124)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_1216_v124).length))]).length)));
          _nw205[(new BigNumber(19)).toNumber()] = _1049_v0;
          _1234_v139 = _nw205;
          let _index235 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1234_v139).length));
          (_1234_v139)[_index235] = _1049_v0;
          let _1236_v140;
          _1236_v140 = _module.D12.create_DC27(p0, _1215_cf0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1049_v0,_dafny.Set.fromElements(p0, _1215_cf0, p2))).length), p0);
          let _1237_v141;
          _1237_v141 = _dafny.Seq.of(_1236_v140);
          let _index236 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1234_v139).length));
          (_1234_v139)[_index236] = !_dafny.areEqual(_1237_v141, _dafny.Seq.Concat(_1237_v141, _1237_v141));
        } else {
          let _1238___mcc_h12 = (_source21).cf3;
          let _1239_cf3 = _1238___mcc_h12;
          let _1240_v142;
          let _nw206 = Array((new BigNumber(10)).toNumber()).fill(false);
          _1240_v142 = _nw206;
          let _index237 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length));
          (_1240_v142)[_index237] = (_module.D5.create_DC14(p0, _1049_v0)).dtor_cf21;
          let _index238 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length));
          (_1240_v142)[_index238] = _1049_v0;
          _1049_v0 = (_1240_v142)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length))];
          let _1241_v143;
          _1241_v143 = _dafny.Set.fromElements((_1240_v142)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length))], (_1240_v142)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length))]);
          let _1242_v144;
          let _nw207 = Array((new BigNumber(10)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = p0;
          _nw207[(_dafny.ONE).toNumber()] = p2;
          _nw207[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), ((_1243_v1) => function (_1244_i7) {
            return _1243_v1;
          })(_1050_v1))).length);
          _nw207[(new BigNumber(3)).toNumber()] = p2;
          _nw207[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Set.fromElements(p0)).length);
          _nw207[(new BigNumber(5)).toNumber()] = p2;
          _nw207[(new BigNumber(6)).toNumber()] = p2;
          _nw207[(new BigNumber(7)).toNumber()] = new BigNumber((_1051_v2).length);
          _nw207[(new BigNumber(8)).toNumber()] = new BigNumber((_1182_cf13).length);
          _nw207[(new BigNumber(9)).toNumber()] = new BigNumber((_1241_v143).length);
          _1242_v144 = _nw207;
          let _1245_v145;
          _1245_v145 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142);
          let _1246_v146;
          _1246_v146 = _dafny.Seq.of(_1245_v145, _1245_v145);
          let _1247_v147;
          let _nw208 = Array((new BigNumber(26)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = ((_1049_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142)) : (_1245_v145));
          _nw208[(_dafny.ONE).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(2)).toNumber()] = (_1245_v145).Merge((_1246_v146)[_module.__default.safeIndex(p0, new BigNumber((_1246_v146).length))]);
          _nw208[(new BigNumber(3)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(4)).toNumber()] = (((_1240_v142)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length))]) ? (_1245_v145) : (_1245_v145));
          _nw208[(new BigNumber(5)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(6)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142)).Merge(_1245_v145);
          _nw208[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142);
          _nw208[(new BigNumber(9)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(10)).toNumber()] = ((_1245_v145).update(_1242_v144, _1240_v142)).Merge(_1245_v145);
          _nw208[(new BigNumber(11)).toNumber()] = ((_1245_v145).update(_1242_v144, _1240_v142)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142));
          _nw208[(new BigNumber(12)).toNumber()] = (_1245_v145).Merge(_1245_v145);
          _nw208[(new BigNumber(13)).toNumber()] = ((_1245_v145).update(_1242_v144, _1240_v142)).Merge(_1245_v145);
          _nw208[(new BigNumber(14)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(15)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(16)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142);
          _nw208[(new BigNumber(18)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(19)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(20)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(21)).toNumber()] = (_1245_v145).Merge(_1245_v145);
          _nw208[(new BigNumber(22)).toNumber()] = (_1245_v145).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142));
          _nw208[(new BigNumber(23)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(24)).toNumber()] = _1245_v145;
          _nw208[(new BigNumber(25)).toNumber()] = _1245_v145;
          _1247_v147 = _nw208;
          let _index239 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1247_v147).length));
          (_1247_v147)[_index239] = _dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142);
          let _1248_v148;
          _1248_v148 = _dafny.MultiSet.fromElements((_1240_v142)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length))]);
          let _1249_v149;
          _1249_v149 = _dafny.MultiSet.fromElements(p2, p2);
          let _1250_v150;
          _1250_v150 = _1249_v149;
          let _1251_v151;
          _1251_v151 = _dafny.Set.fromElements(_1250_v150);
          let _index240 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length));
          let _index241 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1247_v147).length));
          let _rhs231 = (new BigNumber(632)).isLessThan((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1248_v148).cardinality()), new BigNumber((_1251_v151).length))));
          let _rhs232 = (_1245_v145).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1242_v144,_1240_v142)));
          let _lhs185 = _1240_v142;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1240_v142).length));
          let _lhs187 = _1247_v147;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1247_v147).length));
          _lhs185[_lhs186] = _rhs231;
          _lhs187[_lhs188] = _rhs232;
          let _index242 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1242_v144).length));
          (_1242_v144)[_index242] = _module.__default.safeModuloInt(p0, p0);
          let _index243 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1242_v144).length));
          (_1242_v144)[_index243] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1051_v2, _1051_v2), _1051_v2)).length);
        }
        let _1252_v152;
        _1252_v152 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm27(_1049_v0, globalState)).cardinality()), _module.__default.safeModuloInt(p2, p2), ((_this).fm15(new BigNumber((p1).length), globalState)).multipliedBy(p2), (p2).minus(new BigNumber(603)), (p2).multipliedBy(new BigNumber(-108)));
        (globalState).f2 = _1252_v152;
        let _1253_v153;
        _1253_v153 = _dafny.MultiSet.fromElements((_1252_v152).Difference(_1252_v152));
        let _1254_v154;
        let _nw209 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1254_v154 = _nw209;
        let _index244 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1254_v154).length));
        (_1254_v154)[_index244] = p2;
        let _index245 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1254_v154).length));
        let _rhs233 = (_dafny.ZERO).minus(((_1049_v0) ? (p2) : ((_this).fm15(p2, globalState))));
        let _rhs234 = _dafny.Seq.Concat(_1051_v2, _1051_v2);
        let _rhs235 = (_1253_v153).update(_1252_v152, _module.__default.abs(p2));
        let _rhs236 = p2;
        let _lhs189 = globalState;
        let _lhs190 = _1254_v154;
        let _lhs191 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1254_v154).length));
        _lhs189.f3 = _rhs233;
        _1051_v2 = _rhs234;
        _1253_v153 = _rhs235;
        _lhs190[_lhs191] = _rhs236;
      } else {
        let _1255___mcc_h8 = (_source20).cf18;
        let _1256_cf18 = _1255___mcc_h8;
        (globalState).f3 = p0;
        let _1257_v155;
        _1257_v155 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0)),_dafny.Seq.Concat(_1149_v86, _dafny.Seq.of(_1049_v0, _1049_v0)));
        _1257_v155 = (_1257_v155).update(p2, _1149_v86);
        (globalState).f4 = _1049_v0;
        let _1258_v156;
        let _nw210 = new _module.C2();
        _nw210.__ctor();
        _1258_v156 = _nw210;
      }
      let _1259_v157;
      let _nw211 = Array((new BigNumber(9)).toNumber()).fill(false);
      _1259_v157 = _nw211;
      _1259_v157 = _1259_v157;
      let _1260_i8;
      _1260_i8 = _dafny.ZERO;
      L7: {
        while ((((p1).contains(p0)) ? ((p1).get(p0)) : ((_1149_v86)[_module.__default.safeIndex(p2, new BigNumber((_1149_v86).length))]))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1260_i8)) {
              break L7;
            }
            _1260_i8 = (_1260_i8).plus(_dafny.ONE);
            let _1261_v158;
            _1261_v158 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v2,new BigNumber((_1051_v2).length));
            _1261_v158 = (_1261_v158).update(_1051_v2, ((_1049_v0) ? (new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("bly"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("bly")).length)), _1050_v1)).length)) : ((_dafny.ZERO).minus(p2))));
            let _1262_v159;
            _1262_v159 = _dafny.Set.fromElements(p2, new BigNumber(-6), p0);
            let _1263_v160;
            _1263_v160 = _dafny.Map.Empty.slice().updateUnsafe(_1050_v1,_1262_v159);
            _1263_v160 = (_1263_v160).update(_1050_v1, (_1262_v159).Intersect(_dafny.Set.fromElements(p2)));
            let _1264_v161;
            let _nw212 = new _module.C0();
            _nw212.__ctor();
            _1264_v161 = _nw212;
            (globalState).f4 = (new BigNumber((p1).length)).isLessThanOrEqualTo(p2);
          }
        }
      }
      let _1265_v163;
      _1265_v163 = _module.D13.create_DC29(function () {
  let _coll80 = new _dafny.Set();
  for (const _compr_80 of (_module.__default.fm41(_1149_v86, _1050_v1, p0, globalState)).Elements) {
    let _1266_v162 = _compr_80;
    if ((_module.__default.fm41(_1149_v86, _1050_v1, p0, globalState)).contains(_1266_v162)) {
      _coll80.add(_1266_v162);
    }
  }
  return _coll80;
}());
      let _pat_let_tv58 = _1050_v1;
      let _pat_let_tv59 = p2;
      let _pat_let_tv60 = _1049_v0;
      let _pat_let_tv61 = p0;
      let _pat_let_tv62 = _1049_v0;
      let _source22 = function (_source23) {
        if (_source23.is_DC30) {
          let _1267___mcc_h16 = (_source23).cf42;
          let _1268___mcc_h17 = (_source23).cf43;
          let _1269___mcc_h18 = (_source23).cf44;
          let _1270___mcc_h19 = (_source23).cf45;
          let _1271_cf45 = _1270___mcc_h19;
          let _1272_cf44 = _1269___mcc_h18;
          let _1273_cf43 = _1268___mcc_h17;
          let _1274_cf42 = _1267___mcc_h16;
          return _module.D5.create_DC14(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_1275_v1) => function (_1276_i9) {
  return _1275_v1;
})(_pat_let_tv58))).length), !(_1274_cf42));
        } else if (_source23.is_DC29) {
          let _1277___mcc_h20 = (_source23).cf41;
          let _1278_cf41 = _1277___mcc_h20;
          return _module.D5.create_DC14(_pat_let_tv59, _pat_let_tv60);
        } else {
          let _1279___mcc_h21 = (_source23).cf46;
          let _1280_cf46 = _1279___mcc_h21;
          return _module.D5.create_DC14(_pat_let_tv61, _pat_let_tv62);
        }
      }(_1265_v163);
      if (_source22.is_DC14) {
        let _1281___mcc_h13 = (_source22).cf20;
        let _1282___mcc_h14 = (_source22).cf21;
        let _1283_cf21 = _1282___mcc_h14;
        let _1284_cf20 = _1281___mcc_h13;
        _1049_v0 = _1283_cf21;
        let _1285_v164;
        _1285_v164 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(909),!(false) || (_1283_cf21));
        _1285_v164 = (_1285_v164).update(_1284_cf20, false);
        let _1286_v165;
        _1286_v165 = _dafny.MultiSet.fromElements(new BigNumber(-720), p2, p0, _1284_cf20);
        (globalState).f2 = _1286_v165;
        let _1287_v166;
        _1287_v166 = _dafny.Set.fromElements((_this).fm16(_1049_v0, globalState));
        let _1288_v167;
        _1288_v167 = _dafny.Seq.of(p2, p2, new BigNumber((_1287_v166).length));
        let _1289_v168;
        _1289_v168 = _dafny.MultiSet.fromElements((((_this).fm3(_1288_v167, _1286_v165, _1049_v0, globalState)) ? (_1049_v0) : (_1283_cf21)), _1283_cf21, _1283_cf21, (_1149_v86)[_module.__default.safeIndex(p0, new BigNumber((_1149_v86).length))], _1049_v0);
        let _1290_v169;
        _1290_v169 = _module.D2.create_DC7(_1050_v1, false, p0);
        let _rhs237 = (_1289_v168).Difference(_1289_v168);
        let _rhs238 = (_1290_v169).dtor_cf11;
        let _rhs239 = (_this).fm15(_module.__default.safeModuloInt(p2, new BigNumber(824)), globalState);
        let _rhs240 = ((_1049_v0) ? (_1287_v166) : (_1287_v166));
        let _lhs192 = globalState;
        _1289_v168 = _rhs237;
        _lhs192.f3 = _rhs238;
        _1284_cf20 = _rhs239;
        _1287_v166 = _rhs240;
      } else if (_source22.is_DC15) {
        let _1291_v170;
        let _nw213 = new _module.C0();
        _nw213.__ctor();
        _1291_v170 = _nw213;
        let _1292_v171;
        _1292_v171 = _dafny.Seq.of(p0, p0);
        let _1293_v172;
        _1293_v172 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1149_v86);
        (globalState).f4 = !(_module.__default.safeDivisionInt(new BigNumber((_1292_v171).length), p2)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus((p0).multipliedBy(new BigNumber(((((_1293_v172).contains((_1291_v170).fm19(globalState))) ? ((_1293_v172).get((_1291_v170).fm19(globalState))) : (_1149_v86))).length)))));
        let _1294_v173;
        _1294_v173 = _module.D3.create_DC8((_this).fm2(globalState));
        let _1295_v174;
        _1295_v174 = _dafny.MultiSet.fromElements(_1149_v86, _1149_v86);
        (globalState).f3 = (_dafny.ZERO).minus(_module.__default.fm26(_1294_v173, p0, _1295_v174, globalState));
        let _1296_v175;
        let _nw214 = new _module.C2();
        _nw214.__ctor();
        _1296_v175 = _nw214;
      } else {
        let _1297___mcc_h15 = (_source22).cf19;
        let _1298_cf19 = _1297___mcc_h15;
        (globalState).f4 = _dafny.areEqual(_1051_v2, _dafny.Seq.UnicodeFromString("r"));
        let _1299_v176;
        let _nw215 = Array((new BigNumber(3)).toNumber()).fill(_module.D9.Default());
        _1299_v176 = _nw215;
        let _1300_v177;
        _1300_v177 = _dafny.Map.Empty.slice().updateUnsafe(_1299_v176,p0);
        _1300_v177 = ((_1300_v177).Merge(_1300_v177)).Merge(_1300_v177);
        let _1301_v178;
        let _init30 = ((_1302_p2) => function (_1303_i10) {
          return (_1303_i10).plus(_1302_p2);
        })(p2);
        let _nw216 = Array((new BigNumber(17)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw216.length); _i0_30++) {
          _nw216[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1301_v178 = _nw216;
        let _index246 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1301_v178).length));
        (_1301_v178)[_index246] = new BigNumber(899);
        let _index247 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1301_v178).length));
        let _rhs241 = _1049_v0;
        let _rhs242 = _1301_v178;
        let _rhs243 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((true) ? (p0) : (p0))));
        let _lhs193 = globalState;
        let _lhs194 = _1301_v178;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1301_v178).length));
        _lhs193.f4 = _rhs241;
        _1301_v178 = _rhs242;
        _lhs194[_lhs195] = _rhs243;
        let _1304_v179;
        _1304_v179 = _dafny.MultiSet.fromElements(_1049_v0, _1049_v0);
        let _1305_v180;
        _1305_v180 = _dafny.Map.Empty.slice().updateUnsafe(_1050_v1,_1049_v0);
        let _1306_v181;
        _1306_v181 = _dafny.Seq.of(_1305_v180);
        let _1307_v182;
        _1307_v182 = _dafny.MultiSet.fromElements(_1049_v0, _1049_v0, _1049_v0, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1050_v1,_1049_v0)), _module.__default.safeIndex((((_1304_v179).contains(_1049_v0)) ? ((_1304_v179).get(_1049_v0)) : ((_1301_v178)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1301_v178).length))])), new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1050_v1,_1049_v0))).length)), _1305_v180), _1306_v181));
        _1304_v179 = _module.__default.fm0(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("ldambf")).length), (_1301_v178)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_1301_v178).length))]), globalState);
      }
      let _1308_i11;
      _1308_i11 = _dafny.ZERO;
      L8: {
        while (!((_module.__default.fm27(_1049_v0, globalState)).contains(new BigNumber(387)))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1308_i11)) {
              break L8;
            }
            _1308_i11 = (_1308_i11).plus(_dafny.ONE);
            (globalState).f3 = p2;
            let _1309_v183;
            _1309_v183 = _dafny.MultiSet.fromElements(p1);
            (globalState).f8 = (new BigNumber((_1309_v183).cardinality())).plus(p0);
            let _index248 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1259_v157).length));
            (_1259_v157)[_index248] = true;
            let _1310_v184;
            _1310_v184 = _module.D9.create_DC21(_1049_v0, _1149_v86, _1049_v0);
            let _index249 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1259_v157).length));
            (_1259_v157)[_index249] = (_1310_v184).dtor_cf28;
            (globalState).f8 = p2;
          }
        }
      }
      let _1311_v185;
      _1311_v185 = _module.D3.create_DC9();
      r0 = _1311_v185;
      return r0;
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f12 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f12) {
      let _this = this;
      (_this)._f12 = f12;
      return;
    }
    fm2(globalState) {
      let _this = this;
      if (!(_dafny.Set.fromElements(false)).equals(_dafny.Set.fromElements(false, true, true))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), function (_1312_i0) {
          return new BigNumber(206);
        })).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0)))).length))).Merge(function () {
          let _coll81 = new _dafny.Map();
          for (const _compr_81 of _dafny.IntegerRange(new BigNumber(356), new BigNumber(2))) {
            let _1313_v0 = _compr_81;
            if (((new BigNumber(356)).isLessThanOrEqualTo(_1313_v0)) && ((_1313_v0).isLessThan(new BigNumber(2)))) {
              _coll81.push([(_1313_v0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), function (_1314_i1) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              })).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(742))).cardinality()))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(759))).length))).length)]);
            }
          }
          return _coll81;
        }());
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-686),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(959),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),!(true))).length),false)).length))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality()),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-167))).length), new BigNumber(717))).length)));
      }
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return (((true) ? (false) : (true))) || ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(264),true)).length)).isLessThan(new BigNumber(-407)));
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(825)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-928)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gpglxhpdq")).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bw")).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)).length)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("trm")).length));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      let _source24 = _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.MultiSet.fromElements(!(false), true)).cardinality())));
      if (_source24.is_DC1) {
        let _1315___mcc_h0 = (_source24).cf1;
        let _1316___mcc_h1 = (_source24).cf2;
        let _1317_cf2 = _1316___mcc_h1;
        let _1318_cf1 = _1315___mcc_h0;
        return (new BigNumber((_dafny.Set.fromElements(!(true), true, false, !(false))).length)).isLessThanOrEqualTo(_1318_cf1);
      } else if (_source24.is_DC2) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("kammnnur")).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true, !(!(true)), !(false)),false)).length));
      } else if (_source24.is_DC0) {
        let _1319___mcc_h2 = (_source24).cf0;
        let _1320_cf0 = _1319___mcc_h2;
        return !((_dafny.MultiSet.fromElements(_dafny.Seq.of(true))).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Seq.of((_module.D2.create_DC6(_1320_cf0, true, true)).dtor_cf8, false), _dafny.Seq.of(false, true, true, true))));
      } else {
        let _1321___mcc_h3 = (_source24).cf3;
        let _1322_cf3 = _1321___mcc_h3;
        return false;
      }
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1323_v0;
      let _init31 = function (_1324_i0) {
        return (_dafny.MultiSet.fromElements(true)).equals(_dafny.MultiSet.fromElements(true, true, false));
      };
      let _nw217 = Array((new BigNumber(23)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw217.length); _i0_31++) {
        _nw217[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1323_v0 = _nw217;
      let _1325_v1;
      _1325_v1 = false;
      let _index250 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length));
      (_1323_v0)[_index250] = (_1325_v1) === (!(!(!(_1325_v1))));
      let _1326_v2;
      _1326_v2 = _dafny.Seq.UnicodeFromString("upyxghc");
      let _1327_v3;
      _1327_v3 = new BigNumber(-263);
      let _1328_v4;
      _1328_v4 = _dafny.Seq.of(_module.__default.fm14(_1327_v3, globalState), _1326_v2, _dafny.Seq.UnicodeFromString("vjftgkl"));
      let _1329_v5;
      _1329_v5 = _dafny.MultiSet.fromElements(_1326_v2);
      let _index251 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length));
      (_1323_v0)[_index251] = ((_dafny.MultiSet.fromElements(_1326_v2, _1326_v2)).Union(_dafny.MultiSet.FromArray(_1328_v4))).IsProperSubsetOf(_1329_v5);
      _1326_v2 = _1326_v2;
      let _hi6 = _1327_v3;
      for (let _1330_i1 = (_1327_v3).minus(new BigNumber(-637)); _1330_i1.isLessThan(_hi6); _1330_i1 = _1330_i1.plus(_dafny.ONE)) {
        (globalState).f4 = _1325_v1;
        let _1331_v6;
        let _nw218 = new _module.C1();
        _nw218.__ctor();
        _1331_v6 = _nw218;
        (globalState).f4 = ((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]) || (_1325_v1);
        _1323_v0 = _1323_v0;
      }
      let _1332_v7;
      let _nw219 = new _module.C3();
      _nw219.__ctor();
      _1332_v7 = _nw219;
      if ((_1327_v3).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(_1327_v3, new BigNumber(663), (_dafny.ZERO).minus(_1327_v3))).cardinality()))) {
        let _1333_v8;
        let _nw220 = new _module.C1();
        _nw220.__ctor();
        _1333_v8 = _nw220;
        (globalState).f4 = _1325_v1;
        let _1334_v9;
        _1334_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1327_v3, _module.__default.fm20(_1325_v1, globalState)),_1325_v1);
        let _index252 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length));
        (_1323_v0)[_index252] = (((_1334_v9).contains(_1327_v3)) ? ((_1334_v9).get(_1327_v3)) : (true));
        let _1335_v10;
        _1335_v10 = _dafny.MultiSet.fromElements((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]);
        let _1336_v11;
        _1336_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1323_v0,_module.__default.safeDivisionInt(new BigNumber((_1335_v10).cardinality()), _1327_v3));
        let _1337_v12;
        _1337_v12 = _module.D2.create_DC5(_1323_v0);
        let _1338_v13;
        _1338_v13 = _dafny.Map.Empty.slice().updateUnsafe(!((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]),_1327_v3);
        let _1339_v14;
        let _init32 = ((_1340_v3) => function (_1341_i2) {
          return _module.__default.safeDivisionInt(_1341_i2, _1340_v3);
        })(_1327_v3);
        let _nw221 = Array((new BigNumber(19)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw221.length); _i0_32++) {
          _nw221[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1339_v14 = _nw221;
        let _1342_v15;
        _1342_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1338_v13).contains((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))])) ? ((_1338_v13).get((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))])) : (_1327_v3)),_1339_v14);
        _1336_v11 = (_1336_v11).update((_1337_v12).dtor_cf5, new BigNumber((_1342_v15).length));
        (globalState).f3 = _module.__default.safeModuloInt(new BigNumber((_1326_v2).length), _1327_v3);
      } else {
        let _1343_v16;
        let _nw222 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _1343_v16 = _nw222;
        let _1344_v17;
        _1344_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1327_v3,(_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]);
        let _index253 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length));
        (_1343_v16)[_index253] = new BigNumber((_1344_v17).length);
        let _index254 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length));
        (_1343_v16)[_index254] = (new BigNumber(975)).plus((_1327_v3).multipliedBy(_1327_v3));
        (globalState).f4 = ((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]) === (false);
        (globalState).f3 = _1327_v3;
        let _1345_v18;
        _1345_v18 = _dafny.Set.fromElements(new BigNumber(207));
        let _1346_v19;
        let _nw223 = new _module.C0();
        _nw223.__ctor();
        _1346_v19 = _nw223;
        let _1347_v20;
        _1347_v20 = _module.D13.create_DC30((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))], (_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))], _1345_v18, _1346_v19);
        let _1348_v21;
        _1348_v21 = _module.D13.create_DC31(_1347_v20);
        let _1349_v23;
        _1349_v23 = _module.D0.create_DC1(_1327_v3, new BigNumber((_1326_v2).length));
        _1348_v21 = (((function () {
          let _coll82 = new _dafny.Set();
          for (const _compr_82 of _dafny.IntegerRange(new BigNumber(756), new BigNumber(690))) {
            let _1350_v22 = _compr_82;
            if (((new BigNumber(756)).isLessThanOrEqualTo(_1350_v22)) && ((_1350_v22).isLessThan(new BigNumber(690)))) {
              _coll82.add((_1350_v22).multipliedBy(_1327_v3));
            }
          }
          return _coll82;
        }()).IsDisjointFrom(_dafny.Set.fromElements((_1349_v23).dtor_cf1))) ? (((_1325_v1) ? (_1348_v21) : (_1348_v21))) : (_1348_v21));
        if (true) {
          let _1351_v24;
          _1351_v24 = _dafny.Seq.of((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]);
          let _rhs244 = (_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))];
          let _rhs245 = _1327_v3;
          let _rhs246 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1351_v24, _1351_v24), _1351_v24);
          let _lhs196 = globalState;
          let _lhs197 = globalState;
          _lhs196.f4 = _rhs244;
          _lhs197.f3 = _rhs245;
          _1351_v24 = _rhs246;
          _1325_v1 = (new BigNumber((_1351_v24).length)).isLessThan(new BigNumber((_dafny.Seq.Concat(_1326_v2, _module.__default.fm25(globalState))).length));
          _1325_v1 = _1325_v1;
          let _1352_v25;
          _1352_v25 = _dafny.Seq.of(_1327_v3, (_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))]);
          let _1353_v26;
          _1353_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1323_v0,_1352_v25);
          _1353_v26 = (_1353_v26).update(_1323_v0, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_1354_v16) => function (_1355_i3) {
            return (_1354_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1354_v16).length))];
          })(_1343_v16)), _1352_v25));
          let _1356_v27;
          let _init33 = function (_1357_i4) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          };
          let _nw224 = Array((new BigNumber(19)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw224.length); _i0_33++) {
            _nw224[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1356_v27 = _nw224;
          let _1358_v28;
          _1358_v28 = new _dafny.CodePoint('w'.codePointAt(0));
          let _index255 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_1356_v27).length));
          (_1356_v27)[_index255] = _1358_v28;
          let _index256 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_1356_v27).length));
          (_1356_v27)[_index256] = _1358_v28;
        } else {
          let _index257 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length));
          (_1343_v16)[_index257] = (_module.__default.safeDivisionInt((_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))], (_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))])).minus((_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))]);
          let _1359_v29;
          _1359_v29 = _dafny.Seq.of(!(_1325_v1));
          let _rhs247 = _1323_v0;
          let _rhs248 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1359_v29, _dafny.Seq.update(_1359_v29, _module.__default.safeIndex(_1327_v3, new BigNumber((_1359_v29).length)), _1325_v1)), _dafny.Seq.Concat(_1359_v29, _1359_v29));
          let _rhs249 = _1326_v2;
          let _lhs198 = globalState;
          _1323_v0 = _rhs247;
          _lhs198.f4 = _rhs248;
          _1326_v2 = _rhs249;
          let _1360_v30;
          _1360_v30 = _1325_v1;
          let _index258 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length));
          (_1343_v16)[_index258] = (_this).fm12((new BigNumber((_1326_v2).length)).isLessThanOrEqualTo((_dafny.ZERO).minus((_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))])), _1327_v3, _1326_v2, _1360_v30, globalState);
          let _1361_v31;
          _1361_v31 = _dafny.Seq.of((_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))], _1327_v3);
          let _1362_v32;
          _1362_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1343_v16,new BigNumber((_1361_v31).length));
          let _1363_v33;
          _1363_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1327_v3,new BigNumber((_1362_v32).length));
          (globalState).f8 = _module.__default.safeDivisionInt((_1343_v16)[_module.__default.safeIndex(new BigNumber(625), new BigNumber((_1343_v16).length))], new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1327_v3,_1327_v3)).Merge(_1363_v33)).length));
          let _1364_v34;
          let _nw225 = new _module.C2();
          _nw225.__ctor();
          _1364_v34 = _nw225;
        }
      }
      let _1365_v35;
      _1365_v35 = _1325_v1;
      let _index259 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length));
      let _rhs250 = _1325_v1;
      let _rhs251 = (((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]) ? (_1326_v2) : (_dafny.Seq.UnicodeFromString("vtakwmgb")));
      let _rhs252 = (_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))];
      let _rhs253 = !((_1365_v35));
      let _rhs254 = ((_1323_v0)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length))]) || (_1325_v1);
      let _lhs199 = globalState;
      let _lhs200 = _1323_v0;
      let _lhs201 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_1323_v0).length));
      let _lhs202 = globalState;
      _lhs199.f4 = _rhs250;
      _1326_v2 = _rhs251;
      _lhs200[_lhs201] = _rhs252;
      r1 = _rhs253;
      _lhs202.f4 = _rhs254;
      r0 = _1325_v1;
      r1 = (_1327_v3).isLessThan(new BigNumber(-913));
      return [r0, r1];
    }
    m9(p0, p1, globalState) {
      let _this = this;
      let _1366_v0;
      let _init34 = function (_1367_i0) {
        return !(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("fjxvcpd"))).contains(false);
      };
      let _nw226 = Array((new BigNumber(2)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw226.length); _i0_34++) {
        _nw226[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1366_v0 = _nw226;
      let _index260 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
      (_1366_v0)[_index260] = !(false);
      let _1368_v1;
      _1368_v1 = true;
      let _index261 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
      (_1366_v0)[_index261] = _1368_v1;
      (globalState).f8 = new BigNumber((_module.__default.fm31(_1368_v1, (((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))]) ? ((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))]) : (!(_1368_v1))), _1368_v1, globalState)).length);
      let _hi7 = p1;
      for (let _1369_i1 = p0; _1369_i1.isLessThan(_hi7); _1369_i1 = _1369_i1.plus(_dafny.ONE)) {
        let _1370_v2;
        _1370_v2 = _dafny.Seq.UnicodeFromString("koegpi");
        let _1371_v3;
        _1371_v3 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1372_v4;
        _1372_v4 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("smk"));
        let _1373_v5;
        let _nw227 = new _module.C0();
        _nw227.__ctor();
        _1373_v5 = _nw227;
        let _1374_v6;
        _1374_v6 = _dafny.Seq.of(_1373_v5);
        let _1375_v7;
        _1375_v7 = _dafny.Set.fromElements(_1370_v2, (_1372_v4)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1374_v6)).cardinality()), new BigNumber((_1372_v4).length))]);
        let _1376_v8;
        _1376_v8 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.update(_1370_v2, _module.__default.safeIndex(p0, new BigNumber((_1370_v2).length)), _1371_v3), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1368_v1)).length), new BigNumber((_dafny.Seq.update(_1370_v2, _module.__default.safeIndex(p0, new BigNumber((_1370_v2).length)), _1371_v3)).length)), _1371_v3), _1370_v2),_1375_v7);
        let _1377_v9;
        _1377_v9 = _dafny.MultiSet.fromElements(new BigNumber((_1370_v2).length), _1369_i1, p0);
        let _1378_v10;
        _1378_v10 = _dafny.Set.fromElements(_1377_v9);
        _1376_v8 = (_1376_v8).update((_1378_v10).IsSubsetOf(_1378_v10), (_1375_v7).Difference(_1375_v7));
        let _1379_v11;
        _1379_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))],(_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))]);
        let _1380_v12;
        _1380_v12 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(_1368_v1),(_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))]), _1379_v11);
        let _index262 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        let _index263 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        let _index264 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        let _rhs255 = !((_1379_v11).equals((_1380_v12)[_module.__default.safeIndex(_1369_i1, new BigNumber((_1380_v12).length))]));
        let _rhs256 = _dafny.Seq.update(_1370_v2, _module.__default.safeIndex(_1369_i1, new BigNumber((_1370_v2).length)), _1371_v3);
        let _rhs257 = _1368_v1;
        let _rhs258 = (_1377_v9).IsDisjointFrom((_1377_v9).Intersect(_1377_v9));
        let _rhs259 = (new BigNumber((_1370_v2).length)).isLessThan(p1);
        let _lhs203 = globalState;
        let _lhs204 = _1366_v0;
        let _lhs205 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        let _lhs206 = _1366_v0;
        let _lhs207 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        let _lhs208 = _1366_v0;
        let _lhs209 = _module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length));
        _lhs203.f4 = _rhs255;
        _1370_v2 = _rhs256;
        _lhs204[_lhs205] = _rhs257;
        _lhs206[_lhs207] = _rhs258;
        _lhs208[_lhs209] = _rhs259;
        (globalState).f4 = (_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))];
        _1371_v3 = _1371_v3;
      }
      let _1381_v13;
      _1381_v13 = _module.D12.create_DC28(_module.D12.create_DC27(p0, p1, _dafny.ONE, (_dafny.ZERO).minus(p1)));
      let _1382_v14;
      _1382_v14 = _module.D12.create_DC28(_1381_v13);
      let _source25 = _module.D12.create_DC28(_1381_v13);
      if (_source25.is_DC27) {
        let _1383___mcc_h0 = (_source25).cf36;
        let _1384___mcc_h1 = (_source25).cf37;
        let _1385___mcc_h2 = (_source25).cf38;
        let _1386___mcc_h3 = (_source25).cf39;
        let _1387_cf39 = _1386___mcc_h3;
        let _1388_cf38 = _1385___mcc_h2;
        let _1389_cf37 = _1384___mcc_h1;
        let _1390_cf36 = _1383___mcc_h0;
        _1387_cf39 = (_dafny.ZERO).minus(_1390_cf36);
        (globalState).f8 = _module.__default.fm20((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], globalState);
        let _1391_v15;
        _1391_v15 = _dafny.Seq.UnicodeFromString("io");
        _1388_cf38 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_1387_cf39), new BigNumber((_dafny.Seq.Concat(_1391_v15, _1391_v15)).length));
        let _1392_v16;
        let _nw228 = new _module.C3();
        _nw228.__ctor();
        _1392_v16 = _nw228;
      } else if (_source25.is_DC26) {
        let _1393___mcc_h4 = (_source25).cf35;
        let _1394_cf35 = _1393___mcc_h4;
        let _1395_v17;
        let _init35 = ((_1396_v0) => function (_1397_i2) {
          return _module.__default.safeDivisionInt(_1397_i2, new BigNumber((_dafny.MultiSet.fromElements((_1396_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1396_v0).length))], (_1396_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1396_v0).length))])).cardinality()));
        })(_1366_v0);
        let _nw229 = Array((new BigNumber(26)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw229.length); _i0_35++) {
          _nw229[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1395_v17 = _nw229;
        _1395_v17 = _1395_v17;
        let _1398_v18;
        let _init36 = ((_1399_p1) => function (_1400_i3) {
          return _module.D0.create_DC0(_1399_p1);
        })(p1);
        let _nw230 = Array((new BigNumber(20)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw230.length); _i0_36++) {
          _nw230[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1398_v18 = _nw230;
        let _1401_v19;
        _1401_v19 = _dafny.Seq.UnicodeFromString("ag");
        let _1402_v20;
        _1402_v20 = _dafny.Seq.of(p0, new BigNumber((_1401_v19).length));
        let _1403_v21;
        _1403_v21 = _dafny.Seq.of(p0, p0, p0, (_1402_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), function (_1404_i4) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length), new BigNumber((_1402_v20).length))]);
        let _1405_v22;
        _1405_v22 = _module.D0.create_DC0(new BigNumber((_1403_v21).length));
        let _index265 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1398_v18).length));
        (_1398_v18)[_index265] = _1405_v22;
        let _index266 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1398_v18).length));
        (_1398_v18)[_index266] = _1405_v22;
        let _1406_v23;
        _1406_v23 = _dafny.MultiSet.fromElements((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], _1368_v1);
        let _1407_v24;
        _1407_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1406_v23);
        (globalState).f8 = (_dafny.ZERO).minus((((_1406_v23).contains((_1406_v23).IsDisjointFrom((((_1407_v24).contains(p1)) ? ((_1407_v24).get(p1)) : (_1406_v23))))) ? ((_1406_v23).get((_1406_v23).IsDisjointFrom((((_1407_v24).contains(p1)) ? ((_1407_v24).get(p1)) : (_1406_v23))))) : (p0)));
        _1395_v17 = _1395_v17;
      } else {
        let _1408___mcc_h5 = (_source25).cf40;
        let _1409_cf40 = _1408___mcc_h5;
        (globalState).f3 = p0;
        (globalState).f8 = new BigNumber(-323);
        (globalState).f8 = p0;
        let _1410_v25;
        _1410_v25 = _dafny.Seq.UnicodeFromString("sdikjo");
        let _1411_v26;
        _1411_v26 = new _dafny.CodePoint('w'.codePointAt(0));
        let _1412_v27;
        _1412_v27 = _dafny.Seq.of(_1410_v25);
        let _1413_v28;
        _1413_v28 = _dafny.Seq.of(p0);
        let _1414_v29;
        _1414_v29 = _module.D10.create_DC24(p1, !(_1368_v1), _1410_v25);
        let _1415_v30;
        let _nw231 = Array((new BigNumber(20)).toNumber());
        _nw231[(_dafny.ZERO).toNumber()] = _1410_v25;
        _nw231[(_dafny.ONE).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("i");
        _nw231[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("ceextbwbk");
        _nw231[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1410_v25, _dafny.Seq.Create(_module.__default.abs(new BigNumber(451)), function (_1416_i5) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }));
        _nw231[(new BigNumber(5)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(6)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(7)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(8)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lncxgmpd"), _1410_v25);
        _nw231[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1410_v25, _1410_v25);
        _nw231[(new BigNumber(11)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(12)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(13)).toNumber()] = _module.__default.fm17(_1411_v26, globalState);
        _nw231[(new BigNumber(14)).toNumber()] = (_1412_v27)[_module.__default.safeIndex((_1413_v28)[_module.__default.safeIndex(p0, new BigNumber((_1413_v28).length))], new BigNumber((_1412_v27).length))];
        _nw231[(new BigNumber(15)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-408)), ((_1417_v26) => function (_1418_i6) {
          return _1417_v26;
        })(_1411_v26));
        _nw231[(new BigNumber(17)).toNumber()] = (_1414_v29).dtor_cf33;
        _nw231[(new BigNumber(18)).toNumber()] = _1410_v25;
        _nw231[(new BigNumber(19)).toNumber()] = _1410_v25;
        _1415_v30 = _nw231;
        let _index267 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1415_v30).length));
        (_1415_v30)[_index267] = _dafny.Seq.Concat(_1410_v25, _1410_v25);
        let _index268 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1415_v30).length));
        (_1415_v30)[_index268] = _1410_v25;
      }
      let _1419_v31;
      let _init37 = ((_1420_p1, _1421_v1) => function (_1422_i7) {
        return _dafny.Map.Empty.slice().updateUnsafe(_1420_p1,_1421_v1);
      })(p1, _1368_v1);
      let _nw232 = Array((new BigNumber(5)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw232.length); _i0_37++) {
        _nw232[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1419_v31 = _nw232;
      let _1423_v32;
      _1423_v32 = _dafny.MultiSet.fromElements(p1, p0);
      let _index269 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1419_v31).length));
      (_1419_v31)[_index269] = (_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm3(_dafny.Seq.of(p0), _1423_v32, false, globalState))).Merge(function () {
        let _coll83 = new _dafny.Map();
        for (const _compr_83 of _dafny.IntegerRange(new BigNumber(58), new BigNumber(318))) {
          let _1424_v33 = _compr_83;
          if (((new BigNumber(58)).isLessThanOrEqualTo(_1424_v33)) && ((_1424_v33).isLessThan(new BigNumber(318)))) {
            _coll83.push([_module.__default.safeDivisionInt(_1424_v33, p1),_1368_v1]);
          }
        }
        return _coll83;
      }());
      let _1425_v35;
      _1425_v35 = _dafny.Seq.of(_1368_v1);
      let _1426_v36;
      _1426_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-494),new BigNumber((_dafny.MultiSet.FromArray(_1425_v35)).cardinality()));
      let _1427_v37;
      let _nw233 = new _module.C3();
      _nw233.__ctor();
      _1427_v37 = _nw233;
      let _1428_v38;
      _1428_v38 = _module.D0.create_DC1(p0, new BigNumber((_1425_v35).length));
      let _1429_v39;
      _1429_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1427_v37,(_1428_v38).dtor_cf1);
      let _1430_v40;
      _1430_v40 = _1427_v37;
      let _1431_v41;
      _1431_v41 = new _dafny.CodePoint('k'.codePointAt(0));
      let _index270 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1419_v31).length));
      (_1419_v31)[_index270] = (function () {
        let _coll84 = new _dafny.Map();
        for (const _compr_84 of (_1426_v36).Keys.Elements) {
          let _1432_v34 = _compr_84;
          if ((_1426_v36).contains(_1432_v34)) {
            _coll84.push([_module.__default.safeModuloInt(_1432_v34, p0),_1368_v1]);
          }
        }
        return _coll84;
      }()).update((((_1429_v39).contains((_1430_v40))) ? ((_1429_v39).get((_1430_v40))) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_module.__default.fm1(p0, _1431_v41, (_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], _1426_v36, globalState)).length))))), _1368_v1);
      let _1433_v42;
      _1433_v42 = _dafny.Seq.of(_module.__default.fm18((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], (_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], globalState), new _dafny.CodePoint('a'.codePointAt(0)), _1431_v41, _1431_v41);
      let _1434_v43;
      _1434_v43 = _module.D10.create_DC24(p0, _1368_v1, _1433_v42);
      let _1435_v44;
      _1435_v44 = _dafny.Map.Empty.slice().updateUnsafe((_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))],(_1434_v43).dtor_cf33);
      let _1436_v45;
      _1436_v45 = _module.D2.create_DC7(_1431_v41, (_1366_v0)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1366_v0).length))], p1);
      _1435_v44 = (_1435_v44).update((_1436_v45).dtor_cf10, _1433_v42);
      return;
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(globalState) {
      let _this = this;
      let _source26 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray((_module.D4.create_DC10(_dafny.Seq.of(!(false), true))).dtor_cf13)).cardinality()), new BigNumber((function () {
  let _coll85 = new _dafny.Set();
  for (const _compr_85 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),true)).Keys.Elements) {
    let _1437_v0 = _compr_85;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),true)).contains(_1437_v0)) {
      _coll85.add(_1437_v0);
    }
  }
  return _coll85;
}()).length), new BigNumber((_dafny.Set.fromElements(true, true)).length))).length), new BigNumber(277));
      if (_source26.is_DC1) {
        let _1438___mcc_h0 = (_source26).cf1;
        let _1439___mcc_h1 = (_source26).cf2;
        let _1440_cf2 = _1439___mcc_h1;
        let _1441_cf1 = _1438___mcc_h0;
        if (!(!(true))) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1440_cf2,_1441_cf1);
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, false)).length),_1441_cf1);
        }
      } else if (_source26.is_DC2) {
        return (function () {
          let _coll86 = new _dafny.Map();
          for (const _compr_86 of _dafny.IntegerRange(new BigNumber(532), new BigNumber(298))) {
            let _1442_v1 = _compr_86;
            if (((new BigNumber(532)).isLessThanOrEqualTo(_1442_v1)) && ((_1442_v1).isLessThan(new BigNumber(298)))) {
              _coll86.push([(_1442_v1).plus(new BigNumber(-161)),new BigNumber(-906)]);
            }
          }
          return _coll86;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(477),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), function (_1443_i0) {
          return new BigNumber(-366);
        })).length)));
      } else if (_source26.is_DC0) {
        let _1444___mcc_h2 = (_source26).cf0;
        let _1445_cf0 = _1444___mcc_h2;
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1445_cf0,_1445_cf0);
        } else {
          return function () {
            let _coll87 = new _dafny.Map();
            for (const _compr_87 of (_dafny.Set.fromElements(_1445_cf0, new BigNumber(771), _1445_cf0)).Elements) {
              let _1446_v2 = _compr_87;
              if ((_dafny.Set.fromElements(_1445_cf0, new BigNumber(771), _1445_cf0)).contains(_1446_v2)) {
                _coll87.push([(_1446_v2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('k'.codePointAt(0)))).length)),_1445_cf0]);
              }
            }
            return _coll87;
          }();
        }
      } else {
        let _1447___mcc_h3 = (_source26).cf3;
        let _1448_cf3 = _1447___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length),new BigNumber(-365))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length),new BigNumber(561)));
      }
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qvoets")).length), new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-690)), function (_1449_i0) {
        return new BigNumber(-419);
      })), _dafny.Seq.of(new BigNumber(994), new BigNumber(515), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), function (_1450_i1) {
        return new BigNumber(494);
      })).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), function (_1451_i2) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length))));
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1452_v0;
      let _nw234 = new _module.C1();
      _nw234.__ctor();
      _1452_v0 = _nw234;
      let _1453_v1;
      let _nw235 = Array((new BigNumber(22)).toNumber()).fill(_module.D4.Default());
      _1453_v1 = _nw235;
      let _1454_v2;
      _1454_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1453_v1,false);
      let _1455_v3;
      _1455_v3 = _dafny.Seq.UnicodeFromString("dyu");
      let _1456_v4;
      _1456_v4 = _dafny.Set.fromElements(_1455_v3);
      let _1457_v5;
      _1457_v5 = new BigNumber(112);
      let _1458_v6;
      _1458_v6 = new _dafny.CodePoint('o'.codePointAt(0));
      let _1459_v7;
      _1459_v7 = true;
      _1454_v2 = (_1454_v2).update(_1453_v1, !(_1456_v4).equals(function () {
        let _coll88 = new _dafny.Set();
        for (const _compr_88 of (_module.__default.fm42(_1457_v5, _1457_v5, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1458_v6,_1457_v5)).length), _1459_v7, globalState)).Keys.Elements) {
          let _1460_v8 = _compr_88;
          if ((_module.__default.fm42(_1457_v5, _1457_v5, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1458_v6,_1457_v5)).length), _1459_v7, globalState)).contains(_1460_v8)) {
            _coll88.add(_1460_v8);
          }
        }
        return _coll88;
      }()));
      let _1461_v9;
      _1461_v9 = _dafny.MultiSet.fromElements(_1459_v7);
      (globalState).f3 = ((new BigNumber(-68)).plus(_1457_v5)).multipliedBy(new BigNumber(((_1461_v9).Difference(_1461_v9)).cardinality()));
      _1459_v7 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("lxdhgpiq"), _1455_v3);
      let _1462_v10;
      _1462_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1459_v7,_1459_v7);
      let _1463_v11;
      _1463_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(87),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1462_v10).length),_1458_v6));
      let _1464_v12;
      _1464_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1457_v5,new _dafny.CodePoint('q'.codePointAt(0)));
      let _1465_v13;
      _1465_v13 = _dafny.Seq.of(_1464_v12, _1464_v12, _1464_v12, _1464_v12);
      let _1466_v15;
      _1466_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-257),new BigNumber(869));
      let _1467_v16;
      _1467_v16 = _dafny.Seq.of(_1459_v7, !(_1459_v7));
      let _1468_v17;
      _1468_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1459_v7,(_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm26(_module.D3.create_DC8(_1466_v15), new BigNumber((_1455_v3).length), _dafny.MultiSet.fromElements(_dafny.Seq.update(_dafny.Seq.of(_1459_v7, _1459_v7), _module.__default.safeIndex(_1457_v5, new BigNumber((_dafny.Seq.of(_1459_v7, _1459_v7)).length)), false), _1467_v16), globalState))));
      let _1469_v18;
      _1469_v18 = _dafny.Seq.of(new BigNumber((_1455_v3).length), _1457_v5, (_dafny.ZERO).minus(new BigNumber((_1468_v17).length)), _1457_v5, new BigNumber(340));
      let _1470_v19;
      _1470_v19 = _module.D19.create_DC38(_1464_v12);
      let _1471_v21;
      _1471_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1467_v16).length),_1459_v7);
      let _1472_v22;
      let _nw236 = Array((new BigNumber(20)).toNumber());
      _nw236[(_dafny.ZERO).toNumber()] = (((_1463_v11).contains(new BigNumber(8))) ? ((_1463_v11).get(new BigNumber(8))) : (_dafny.Map.Empty.slice().updateUnsafe(_1457_v5,_1458_v6)));
      _nw236[(_dafny.ONE).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(2)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(3)).toNumber()] = (_1464_v12).Merge(_1464_v12);
      _nw236[(new BigNumber(4)).toNumber()] = ((_1465_v13)[_module.__default.safeIndex(_1457_v5, new BigNumber((_1465_v13).length))]).Merge(_1464_v12);
      _nw236[(new BigNumber(5)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(6)).toNumber()] = (_1464_v12).Merge(_1464_v12);
      _nw236[(new BigNumber(7)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(8)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(9)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(10)).toNumber()] = _module.__default.fm43(!(_1459_v7), _1457_v5, _1458_v6, globalState);
      _nw236[(new BigNumber(11)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(12)).toNumber()] = function () {
        let _coll89 = new _dafny.Map();
        for (const _compr_89 of (_1469_v18).Elements) {
          let _1473_v14 = _compr_89;
          if (_dafny.Seq.contains(_1469_v18, _1473_v14)) {
            _coll89.push([(_1473_v14).multipliedBy(_1457_v5),_1458_v6]);
          }
        }
        return _coll89;
      }();
      _nw236[(new BigNumber(13)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(14)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(15)).toNumber()] = _1464_v12;
      _nw236[(new BigNumber(16)).toNumber()] = ((_1459_v7) ? (_dafny.Map.Empty.slice().updateUnsafe(_1457_v5,_1458_v6)) : ((_1470_v19).dtor_cf57));
      _nw236[(new BigNumber(17)).toNumber()] = (_1464_v12).Merge(_1464_v12);
      _nw236[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1457_v5,_1458_v6);
      _nw236[(new BigNumber(19)).toNumber()] = ((function () {
        let _coll90 = new _dafny.Map();
        for (const _compr_90 of (_1471_v21).Keys.Elements) {
          let _1474_v20 = _compr_90;
          if ((_1471_v21).contains(_1474_v20)) {
            _coll90.push([(_1474_v20).plus(new BigNumber((_1464_v12).length)),new _dafny.CodePoint('a'.codePointAt(0))]);
          }
        }
        return _coll90;
      }()).update(_1457_v5, _1458_v6)).Merge(_1464_v12);
      _1472_v22 = _nw236;
      _1472_v22 = _1472_v22;
      let _1475_v23;
      _1475_v23 = _module.D0.create_DC1(_1457_v5, _1457_v5);
      let _1476_v24;
      _1476_v24 = _dafny.Set.fromElements(_1475_v23, _1475_v23);
      let _1477_v25;
      _1477_v25 = _dafny.MultiSet.fromElements(_1457_v5);
      let _1478_v26;
      _1478_v26 = _module.D19.create_DC39(_1459_v7, _1457_v5);
      (_1452_v0).m13(new BigNumber((_1476_v24).length), ((_1459_v7) ? ((_1477_v25).update(new BigNumber((_1455_v3).length), _module.__default.abs(_1457_v5))) : (_module.__default.fm27((_1478_v26).dtor_cf58, globalState))), _1459_v7, globalState);
      r0 = !(true);
      r1 = !((((_1457_v5).isLessThanOrEqualTo(_1457_v5)) ? (_1459_v7) : ((_1452_v0).fm3(_1469_v18, _1477_v25, _1459_v7, globalState))));
      return [r0, r1];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f10, f11) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_module.D0.create_DC1((_this).f10, (_this).f10)).dtor_cf2, new BigNumber(547)),(_this).f10);
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.of(!(true)), _dafny.Seq.of(true, false));
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1479_v0;
      let _nw237 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _1479_v0 = _nw237;
      let _index271 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length));
      (_1479_v0)[_index271] = (_this).f10;
      let _1480_v1;
      _1480_v1 = false;
      let _1481_v2;
      _1481_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1480_v1,(_this).f10);
      let _index272 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length));
      (_1479_v0)[_index272] = ((_1480_v1) ? (new BigNumber(((_this).f11).length)) : ((((_1481_v2).contains(_1480_v1)) ? ((_1481_v2).get(_1480_v1)) : (new BigNumber(757)))));
      let _hi8 = new BigNumber(149);
      for (let _1482_i0 = (_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))]; _1482_i0.isLessThan(_hi8); _1482_i0 = _1482_i0.plus(_dafny.ONE)) {
        _1479_v0 = ((_1480_v1) ? (_1479_v0) : (_1479_v0));
        let _1483_v3;
        let _nw238 = new _module.C3();
        _nw238.__ctor();
        _1483_v3 = _nw238;
        let _1484_v4;
        _1484_v4 = _dafny.Set.fromElements(_1483_v3);
        _1484_v4 = (_1484_v4).Intersect(_1484_v4);
        let _1485_v5;
        _1485_v5 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1486_v6;
        _1486_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1482_i0);
        let _1487_v7;
        _1487_v7 = _dafny.MultiSet.fromElements((_this).f10);
        let _1488_v8;
        _1488_v8 = _dafny.Seq.of(_1480_v1, _1480_v1, _1480_v1, _1480_v1, _1480_v1);
        let _1489_v9;
        _1489_v9 = _dafny.Seq.of(new BigNumber(((_this).f11).length), (_dafny.ZERO).minus(_1482_i0));
        let _1490_v10;
        _1490_v10 = _dafny.Map.Empty.slice().updateUnsafe(!(_1480_v1),_1480_v1);
        let _1491_v11;
        let _nw239 = Array((new BigNumber(23)).toNumber());
        _nw239[(_dafny.ZERO).toNumber()] = _1480_v1;
        _nw239[(_dafny.ONE).toNumber()] = !(_1480_v1) || (_1480_v1);
        _nw239[(new BigNumber(2)).toNumber()] = true;
        _nw239[(new BigNumber(3)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(4)).toNumber()] = ((_1480_v1) ? (_1480_v1) : ((_1483_v3).fm3(_module.__default.fm1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),(_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))])).length), _1485_v5, _1480_v1, _1486_v6, globalState), _1487_v7, _1480_v1, globalState)));
        _nw239[(new BigNumber(5)).toNumber()] = !(_1480_v1);
        _nw239[(new BigNumber(6)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(7)).toNumber()] = (_1488_v8)[_module.__default.safeIndex((_this).f10, new BigNumber((_1488_v8).length))];
        _nw239[(new BigNumber(8)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(9)).toNumber()] = ((_this).f10).isLessThanOrEqualTo((_this).f10);
        _nw239[(new BigNumber(10)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(11)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(12)).toNumber()] = (_1483_v3).fm3(_1489_v9, _1487_v7, _1480_v1, globalState);
        _nw239[(new BigNumber(13)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(14)).toNumber()] = (_1480_v1) === (true);
        _nw239[(new BigNumber(15)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(16)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(17)).toNumber()] = !(_1490_v10).equals(_1490_v10);
        _nw239[(new BigNumber(18)).toNumber()] = ((_1480_v1) ? (_1480_v1) : (_1480_v1));
        _nw239[(new BigNumber(19)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(20)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(21)).toNumber()] = _1480_v1;
        _nw239[(new BigNumber(22)).toNumber()] = _1480_v1;
        _1491_v11 = _nw239;
        let _index273 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length));
        (_1491_v11)[_index273] = _1480_v1;
        let _index274 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length));
        (_1491_v11)[_index274] = !((!(_1487_v7).contains((_this).f10)) && (_1480_v1));
        if (((new BigNumber(515)).minus((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))])).isEqualTo(new BigNumber(-38))) {
          let _1492_v12;
          _1492_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1480_v1,_1485_v5);
          _1492_v12 = (_1492_v12).update(_1480_v1, _1485_v5);
          (globalState).f4 = (_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))];
          let _1493_v13;
          _1493_v13 = _dafny.MultiSet.fromElements(_1480_v1, false);
          let _1494_v14;
          _1494_v14 = _dafny.Set.fromElements(new BigNumber((_1493_v13).cardinality()), _module.__default.safeDivisionInt((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))], (_this).f10));
          let _1495_v15;
          _1495_v15 = _dafny.Seq.of(_1479_v0);
          let _index275 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length));
          let _rhs260 = _1494_v14;
          let _rhs261 = !(_dafny.areEqual(_1495_v15, _1495_v15));
          let _lhs210 = _1491_v11;
          let _lhs211 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length));
          _1494_v14 = _rhs260;
          _lhs210[_lhs211] = _rhs261;
          let _index276 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length));
          (_1491_v11)[_index276] = ((_this).f10).isLessThan(_1482_i0);
          let _1496_v16;
          _1496_v16 = _dafny.MultiSet.fromElements(_1488_v8);
          r1 = ((_1496_v16).update(_1488_v8, _module.__default.abs((_dafny.ZERO).minus((_this).f10)))).IsProperSubsetOf((_1496_v16).Intersect(_dafny.MultiSet.fromElements(_1488_v8, _1488_v8)));
        } else {
          let _1497_v17;
          _1497_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1485_v5,_1490_v10);
          let _1498_v18;
          _1498_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_1497_v17).contains(_1485_v5)) ? ((_1497_v17).get(_1485_v5)) : (_1490_v10)),_1479_v0);
          (globalState).f4 = (_1498_v18).contains(_1490_v10);
          r0 = !(_1480_v1);
          let _1499_v19;
          _1499_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))],_dafny.Seq.of(new BigNumber((_1489_v9).length), (_dafny.ZERO).minus((_this).f10)));
          let _rhs262 = (_1480_v1) && (_1480_v1);
          let _rhs263 = (_1499_v19).Merge(_1499_v19);
          let _lhs212 = globalState;
          _lhs212.f4 = _rhs262;
          _1499_v19 = _rhs263;
          let _1500_v20;
          _1500_v20 = _dafny.Set.fromElements((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))], (_dafny.ZERO).minus(new BigNumber(((_this).f11).length)), (_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))], (_1482_i0).minus((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))]));
          let _1501_v21;
          _1501_v21 = _dafny.MultiSet.fromElements(_1480_v1);
          let _1502_v22;
          let _init38 = ((_1503_v0, _1504_v1, _1505_v5) => function (_1506_i1) {
            return _dafny.Seq.Concat((_module.D10.create_DC24((_1503_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1503_v0).length))], _1504_v1, (_this).f11)).dtor_cf33, _dafny.Seq.Create(_module.__default.abs(new BigNumber(834)), ((_1507_v5) => function (_1508_i2) {
              return _1507_v5;
            })(_1505_v5)));
          })(_1479_v0, _1480_v1, _1485_v5);
          let _nw240 = Array((new BigNumber(21)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw240.length); _i0_38++) {
            _nw240[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1502_v22 = _nw240;
          let _index277 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_1502_v22).length));
          (_1502_v22)[_index277] = (_this).f11;
          let _1509_v23;
          _1509_v23 = _module.D19.create_DC39(true, (_this).f10);
          let _1510_v24;
          _1510_v24 = _module.D19.create_DC40(_1509_v23);
          let _1511_v25;
          _1511_v25 = _dafny.Set.fromElements(_1510_v24, _1510_v24, _module.D19.create_DC40(_1509_v23), _1510_v24);
          let _1512_v26;
          _1512_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1511_v25,(_dafny.MultiSet.fromElements((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))], _1480_v1, (_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))], false)).Union(_dafny.MultiSet.FromArray(_module.__default.fm31(_1480_v1, (_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))], _1480_v1, globalState))));
          let _1513_v27;
          _1513_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))],_1488_v8);
          let _index278 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_1502_v22).length));
          let _rhs264 = _1500_v20;
          let _rhs265 = (((_1512_v26).contains(_1511_v25)) ? ((_1512_v26).get(_1511_v25)) : ((_1501_v21).Union(_1501_v21)));
          let _rhs266 = new BigNumber(((((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))]) ? ((((_1513_v27).contains((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))])) ? ((_1513_v27).get((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))])) : (_1488_v8))) : (_1488_v8))).length);
          let _rhs267 = _module.__default.safeModuloInt(new BigNumber((_1487_v7).cardinality()), _1482_i0);
          let _rhs268 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-313)), ((_1514_v5) => function (_1515_i3) {
            return _1514_v5;
          })(_1485_v5)), (_this).f11);
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          let _lhs215 = _1502_v22;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_1502_v22).length));
          _1500_v20 = _rhs264;
          _1501_v21 = _rhs265;
          _lhs213.f3 = _rhs266;
          _lhs214.f8 = _rhs267;
          _lhs215[_lhs216] = _rhs268;
          let _1516_v28;
          _1516_v28 = _dafny.Set.fromElements((_1491_v11)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1491_v11).length))], true);
          _1516_v28 = ((_1516_v28).Intersect(_1516_v28)).Intersect(_dafny.Set.fromElements(true));
        }
      }
      let _1517_v29;
      let _nw241 = Array((new BigNumber(17)).toNumber());
      _nw241[(_dafny.ZERO).toNumber()] = false;
      _nw241[(_dafny.ONE).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(2)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(3)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(4)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(5)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(6)).toNumber()] = true;
      _nw241[(new BigNumber(7)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(8)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(9)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(10)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(11)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(12)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(13)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(14)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(15)).toNumber()] = _1480_v1;
      _nw241[(new BigNumber(16)).toNumber()] = _1480_v1;
      _1517_v29 = _nw241;
      let _1518_v30;
      _1518_v30 = _dafny.Set.fromElements(_1517_v29);
      let _1519_v31;
      _1519_v31 = _module.D2.create_DC5(_1517_v29);
      let _1520_v32;
      _1520_v32 = _dafny.Set.fromElements(_1480_v1, _1480_v1);
      let _1521_v34;
      _1521_v34 = _dafny.Set.fromElements((_this).f10, (_this).f10);
      let _1522_v35;
      _1522_v35 = _dafny.Seq.of(_1480_v1);
      let _1523_v36;
      _1523_v36 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_1522_v35).length)));
      let _1524_v37;
      _1524_v37 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1525_v38;
      _1525_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("wvptjdmb"), _module.__default.safeIndex(new BigNumber(((_1523_v36)[_module.__default.safeIndex((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))], new BigNumber((_1523_v36).length))]).length), new BigNumber((_dafny.Seq.UnicodeFromString("wvptjdmb")).length)), _1524_v37),false);
      let _1526_v39;
      let _nw242 = Array((new BigNumber(9)).toNumber());
      _nw242[(_dafny.ZERO).toNumber()] = ((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f10));
      _nw242[(_dafny.ONE).toNumber()] = (_1518_v30).IsDisjointFrom(_dafny.Set.fromElements(_1517_v29, (_1519_v31).dtor_cf5));
      _nw242[(new BigNumber(2)).toNumber()] = (_1520_v32).IsDisjointFrom(_dafny.Set.fromElements(_1480_v1));
      _nw242[(new BigNumber(3)).toNumber()] = _1480_v1;
      _nw242[(new BigNumber(4)).toNumber()] = _1480_v1;
      _nw242[(new BigNumber(5)).toNumber()] = _1480_v1;
      _nw242[(new BigNumber(6)).toNumber()] = _1480_v1;
      _nw242[(new BigNumber(7)).toNumber()] = !((_1521_v34).IsProperSubsetOf(function () {
        let _coll91 = new _dafny.Set();
        for (const _compr_91 of _dafny.IntegerRange(new BigNumber(964), new BigNumber(-933))) {
          let _1527_v33 = _compr_91;
          if (((new BigNumber(964)).isLessThanOrEqualTo(_1527_v33)) && ((_1527_v33).isLessThan(new BigNumber(-933)))) {
            _coll91.add((_1527_v33).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1480_v1)).length))));
          }
        }
        return _coll91;
      }()));
      _nw242[(new BigNumber(8)).toNumber()] = ((_1480_v1) ? (!(false)) : ((((_1525_v38).contains((_this).f11)) ? ((_1525_v38).get((_this).f11)) : (_1480_v1))));
      _1526_v39 = _nw242;
      let _index279 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length));
      (_1517_v29)[_index279] = true;
      let _index280 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length));
      (_1517_v29)[_index280] = false;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1479_v0).length))) {
        let _1528_i4 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1528_i4)) && ((_1528_i4).isLessThan(new BigNumber((_1479_v0).length))))) {
          (_1479_v0)[(_1528_i4)] = (_1528_i4).multipliedBy((_this).f10);
        }
      }
      let _hi9 = (_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))];
      for (let _1529_i5 = (_this).f10; _1529_i5.isLessThan(_hi9); _1529_i5 = _1529_i5.plus(_dafny.ONE)) {
        r0 = _1480_v1;
        let _1530_v40;
        _1530_v40 = _dafny.Seq.of(_1479_v0);
        _1479_v0 = ((_1480_v1) ? ((_1530_v40)[_module.__default.safeIndex((_this).f10, new BigNumber((_1530_v40).length))]) : (_1479_v0));
        let _1531_v41;
        let _nw243 = new _module.C0();
        _nw243.__ctor();
        _1531_v41 = _nw243;
        let _1532_v42;
        _1532_v42 = _module.D7.create_DC18();
        let _source27 = _1532_v42;
        if (_source27.is_DC18) {
          (globalState).f3 = _module.__default.fm20((_1517_v29)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length))], globalState);
          let _1533_v43;
          _1533_v43 = _dafny.Seq.of((_this).f10, _1529_i5);
          let _1534_v44;
          _1534_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1480_v1,(_1517_v29)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length))]);
          let _index281 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length));
          (_1517_v29)[_index281] = (_module.__default.safeModuloInt((_dafny.ZERO).minus(_1529_i5), new BigNumber((_dafny.Seq.update(_1533_v43, _module.__default.safeIndex(new BigNumber((_1522_v35).length), new BigNumber((_1533_v43).length)), new BigNumber((_1534_v44).length))).length))).isLessThanOrEqualTo((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))]);
          let _1535_v45;
          _1535_v45 = _module.D9.create_DC20(_1520_v32);
          _1535_v45 = _1535_v45;
          let _1536_v46;
          _1536_v46 = _dafny.Seq.UnicodeFromString("mn");
          _1536_v46 = _1536_v46;
        } else {
          let _1537___mcc_h0 = (_source27).cf23;
          let _1538_cf23 = _1537___mcc_h0;
          let _1539_v47;
          _1539_v47 = _dafny.MultiSet.fromElements((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))], new BigNumber(534));
          let _1540_v48;
          _1540_v48 = _dafny.MultiSet.fromElements((_this).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(267)), function (_1541_i6) {
            return new BigNumber(634);
          }), _1539_v47, (_1517_v29)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length))], globalState), _1480_v1, ((_this).f10).isLessThanOrEqualTo(new BigNumber(790)), (_1517_v29)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length))], _1480_v1);
          let _index282 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length));
          (_1479_v0)[_index282] = new BigNumber((_1540_v48).cardinality());
          let _index283 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length));
          (_1479_v0)[_index283] = (_dafny.ZERO).minus((((_1517_v29)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_1517_v29).length))]) ? (new BigNumber(817)) : (_module.__default.safeDivisionInt((_this).f10, _1529_i5))));
          (globalState).f8 = _1529_i5;
          (globalState).f3 = (_dafny.ZERO).minus((_1479_v0)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_1479_v0).length))]);
        }
      }
      _1524_v37 = _1524_v37;
      r0 = !_dafny.areEqual((_this).f11, (_this).f11);
      r1 = _1480_v1;
      return [r0, r1];
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let _1542_i0;
      _1542_i0 = _dafny.ZERO;
      L9: {
        while (p0) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1542_i0)) {
              break L9;
            }
            _1542_i0 = (_1542_i0).plus(_dafny.ONE);
            let _1543_v0;
            _1543_v0 = _dafny.Seq.of(p0, ((p0) ? (p0) : (p0)), p0, !(p0));
            let _1544_v1;
            let _init39 = function (_1545_i1) {
              return _dafny.Seq.update((_this).f11, _module.__default.safeIndex((_this).f10, new BigNumber(((_this).f11).length)), new _dafny.CodePoint('l'.codePointAt(0)));
            };
            let _nw244 = Array((new BigNumber(5)).toNumber());
            for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw244.length); _i0_39++) {
              _nw244[_i0_39] = _init39(new BigNumber(_i0_39));
            }
            _1544_v1 = _nw244;
            let _index284 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1544_v1).length));
            (_1544_v1)[_index284] = _dafny.Seq.Concat(_module.__default.fm14((_this).f10, globalState), (_this).f11);
            let _1546_v2;
            let _nw245 = new _module.C2();
            _nw245.__ctor();
            _1546_v2 = _nw245;
            let _1547_v3;
            _1547_v3 = _module.D16.create_DC34(_1546_v2);
            let _1548_v4;
            _1548_v4 = new _dafny.CodePoint('x'.codePointAt(0));
            let _1549_v5;
            _1549_v5 = _dafny.MultiSet.fromElements((_this).f10);
            let _1550_v6;
            _1550_v6 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm27(p0, globalState)).IsProperSubsetOf(_1549_v5),_1548_v4);
            let _index285 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1544_v1).length));
            let _rhs269 = _dafny.Seq.Concat(_1543_v0, _1543_v0);
            let _rhs270 = _dafny.Seq.Concat((_this).f11, (_this).f11);
            let _rhs271 = _1547_v3;
            let _rhs272 = (((_1550_v6).contains(((p0) ? (p0) : (p0)))) ? ((_1550_v6).get(((p0) ? (p0) : (p0)))) : (_1548_v4));
            let _rhs273 = p0;
            let _lhs217 = _1544_v1;
            let _lhs218 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1544_v1).length));
            let _lhs219 = globalState;
            _1543_v0 = _rhs269;
            _lhs217[_lhs218] = _rhs270;
            _1547_v3 = _rhs271;
            _1548_v4 = _rhs272;
            _lhs219.f4 = _rhs273;
            (globalState).f4 = p0;
            let _1551_v7;
            _1551_v7 = _dafny.Seq.of((_this).f10);
            if ((_1546_v2).fm3(_1551_v7, _1549_v5, p0, globalState)) {
              let _1552_v8;
              let _1553_v9;
              let _1554_v10;
              let _1555_v11;
              let _out22;
              let _out23;
              let _out24;
              let _out25;
              let _outcollector5 = (_1546_v2).m12(globalState);
              _out22 = _outcollector5[0];
              _out23 = _outcollector5[1];
              _out24 = _outcollector5[2];
              _out25 = _outcollector5[3];
              _1552_v8 = _out22;
              _1553_v9 = _out23;
              _1554_v10 = _out24;
              _1555_v11 = _out25;
              (globalState).f4 = _1555_v11;
              let _1556_v12;
              _1556_v12 = _dafny.Set.fromElements((_this).f10);
              _1556_v12 = (_1556_v12).Intersect(_1556_v12);
              (globalState).f8 = ((_dafny.ZERO).minus(_1553_v9)).multipliedBy(((_1551_v7)[_module.__default.safeIndex(_1554_v10, new BigNumber((_1551_v7).length))]).plus((_dafny.ZERO).minus(new BigNumber((_1549_v5).cardinality()))));
              let _index286 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1544_v1).length));
              (_1544_v1)[_index286] = (_this).f11;
            } else {
              _1548_v4 = _1548_v4;
              (globalState).f4 = p0;
              let _1557_v13;
              _1557_v13 = _dafny.MultiSet.fromElements(!(p0), p0);
              let _1558_v14;
              _1558_v14 = _module.D9.create_DC22(_dafny.Seq.update(_1543_v0, _module.__default.safeIndex(new BigNumber((_1557_v13).cardinality()), new BigNumber((_1543_v0).length)), p0));
              let _1559_v15;
              _1559_v15 = _dafny.Set.fromElements(p0);
              let _1560_v16;
              _1560_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1558_v14,((_this).f10).multipliedBy(new BigNumber((_1559_v15).length)));
              let _1561_v18;
              _1561_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(114),new BigNumber((function () {
                let _coll92 = new _dafny.Map();
                for (const _compr_92 of _dafny.IntegerRange(new BigNumber(-921), new BigNumber(719))) {
                  let _1562_v17 = _compr_92;
                  if (((new BigNumber(-921)).isLessThanOrEqualTo(_1562_v17)) && ((_1562_v17).isLessThan(new BigNumber(719)))) {
                    _coll92.push([_module.__default.safeModuloInt(_1562_v17, new BigNumber((_1559_v15).length)),_1548_v4]);
                  }
                }
                return _coll92;
              }()).length));
              let _1563_v19;
              _1563_v19 = _dafny.MultiSet.fromElements(_1543_v0);
              _1560_v16 = (_1560_v16).update(_1558_v14, _module.__default.fm26(_module.D3.create_DC8(_1561_v18), (_this).f10, _1563_v19, globalState));
              let _1564_v20;
              _1564_v20 = _dafny.Set.fromElements((_this).f10, (_this).f10, (_this).f10);
              (globalState).f8 = ((_this).f10).minus(new BigNumber((_1564_v20).length));
              (globalState).f4 = ((_this).f10).isLessThanOrEqualTo((((_1561_v18).contains((_this).f10)) ? ((_1561_v18).get((_this).f10)) : ((_this).f10)));
            }
            let _1565_v21;
            let _nw246 = new _module.C3();
            _nw246.__ctor();
            _1565_v21 = _nw246;
          }
        }
      }
      let _1566_v22;
      let _init40 = function (_1567_i3) {
        return _dafny.Seq.Concat(_dafny.Seq.update((_this).f11, _module.__default.safeIndex((_this).f10, new BigNumber(((_this).f11).length)), new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Seq.UnicodeFromString("kjr"));
      };
      let _nw247 = Array((new BigNumber(14)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw247.length); _i0_40++) {
        _nw247[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1566_v22 = _nw247;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1566_v22).length))) {
        let _1568_i2 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1568_i2)) && ((_1568_i2).isLessThan(new BigNumber((_1566_v22).length))))) {
          (_1566_v22)[(_1568_i2)] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-239)), function (_1569_i4) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hxtmbhm"), (_this).f11));
        }
      }
      let _1570_v23;
      let _init41 = ((_1571_p0) => function (_1572_i5) {
        return (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1571_p0, _1571_p0)).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_1573_i6) {
          return (_this).f10;
        }))).length))).IsSubsetOf(_dafny.Set.fromElements((_this).f10));
      })(p0);
      let _nw248 = Array((new BigNumber(12)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw248.length); _i0_41++) {
        _nw248[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1570_v23 = _nw248;
      let _index287 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
      (_1570_v23)[_index287] = p0;
      let _1574_v24;
      _1574_v24 = _dafny.Seq.of((_this).f10, (_this).f10);
      let _1575_v25;
      _1575_v25 = _dafny.Seq.of(new BigNumber((_1574_v24).length));
      let _index288 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
      let _rhs274 = (_this).f10;
      let _rhs275 = ((p0) ? (p0) : ((_this).fm3(_1575_v25, _dafny.MultiSet.fromElements((_this).f10), p0, globalState)));
      let _rhs276 = (p0) && (p0);
      let _rhs277 = _module.__default.safeModuloInt((_this).f10, (_this).f10);
      let _lhs220 = globalState;
      let _lhs221 = _1570_v23;
      let _lhs222 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
      let _lhs223 = globalState;
      let _lhs224 = globalState;
      _lhs220.f8 = _rhs274;
      _lhs221[_lhs222] = _rhs275;
      _lhs223.f4 = _rhs276;
      _lhs224.f8 = _rhs277;
      let _1576_v26;
      let _nw249 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1576_v26 = _nw249;
      let _1577_v27;
      let _nw250 = Array((new BigNumber(18)).toNumber());
      _nw250[(_dafny.ZERO).toNumber()] = _1576_v26;
      _nw250[(_dafny.ONE).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(2)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(3)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(4)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(5)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(6)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(7)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(8)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(9)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(10)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(11)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(12)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(13)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(14)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(15)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(16)).toNumber()] = _1576_v26;
      _nw250[(new BigNumber(17)).toNumber()] = _1576_v26;
      _1577_v27 = _nw250;
      let _1578_v28;
      _1578_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1577_v27,new BigNumber(179));
      let _hi10 = (_dafny.ZERO).minus(((_this).f10).plus((_this).f10));
      for (let _1579_i7 = (((_1578_v28).contains(_1577_v27)) ? ((_1578_v28).get(_1577_v27)) : (_module.__default.fm20((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))], globalState))); _1579_i7.isLessThan(_hi10); _1579_i7 = _1579_i7.plus(_dafny.ONE)) {
        let _1580_v29;
        _1580_v29 = _dafny.MultiSet.fromElements(p0);
        let _1581_v30;
        _1581_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))],(_1580_v29).IsSubsetOf(_1580_v29));
        let _1582_v31;
        _1582_v31 = _dafny.MultiSet.fromElements((_this).f10);
        let _index289 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
        let _rhs278 = ((_1579_i7).plus((_this).f10)).minus(_module.__default.safeDivisionInt(new BigNumber(114), new BigNumber((_1582_v31).cardinality())));
        let _rhs279 = (_this).f10;
        let _rhs280 = ((_1581_v30).Merge(_1581_v30)).Merge(_1581_v30);
        let _rhs281 = (_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))];
        let _rhs282 = (_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))];
        let _lhs225 = globalState;
        let _lhs226 = globalState;
        let _lhs227 = globalState;
        let _lhs228 = _1570_v23;
        let _lhs229 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
        _lhs225.f3 = _rhs278;
        _lhs226.f8 = _rhs279;
        _1581_v30 = _rhs280;
        _lhs227.f4 = _rhs281;
        _lhs228[_lhs229] = _rhs282;
        let _1583_v32;
        let _nw251 = new _module.C5();
        _nw251.__ctor();
        _1583_v32 = _nw251;
        let _1584_v33;
        _1584_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1581_v30,_1579_i7);
        let _1585_v34;
        _1585_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)))).length),(_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]);
        let _1586_v35;
        _1586_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1582_v31,!((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0)).equals(_1585_v34)));
        let _1587_v36;
        let _nw252 = new _module.C0();
        _nw252.__ctor();
        _1587_v36 = _nw252;
        let _1588_v37;
        _1588_v37 = _dafny.Set.fromElements(_1587_v36, _1587_v36);
        let _rhs283 = (_1584_v33).Merge((_1584_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1581_v30,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_1588_v37).length))).length))));
        let _rhs284 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_1579_i7)).Intersect(_1582_v31),(_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]);
        let _rhs285 = _1576_v26;
        let _rhs286 = _dafny.Seq.Concat(_1574_v24, _1575_v25);
        let _lhs230 = globalState;
        _1584_v33 = _rhs283;
        _1586_v35 = _rhs284;
        _1576_v26 = _rhs285;
        _lhs230.f0 = _rhs286;
        let _1589_v38;
        _1589_v38 = _dafny.Seq.of(_1576_v26);
        let _1590_v39;
        _1590_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1589_v38, _1589_v38),(_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]);
        let _1591_v40;
        _1591_v40 = new _dafny.CodePoint('l'.codePointAt(0));
        (globalState).f4 = (((_1590_v39).contains(_1589_v38)) ? ((_1590_v39).get(_1589_v38)) : ((new BigNumber((_dafny.Seq.update((_this).f11, _module.__default.safeIndex(_1579_i7, new BigNumber(((_this).f11).length)), _1591_v40)).length)).isLessThan((_this).f10)));
      }
      if (true) {
        let _1592_v41;
        _1592_v41 = _dafny.Seq.of((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]);
        let _1593_v42;
        _1593_v42 = _module.D4.create_DC10(_dafny.Seq.update(_1592_v41, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_this).f10), (_this).f10)).length), new BigNumber((_1592_v41).length)), p0));
        let _1594_v43;
        _1594_v43 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f10).multipliedBy((_this).f10),_1593_v42);
        _1594_v43 = (_1594_v43).update((_this).f10, _1593_v42);
        let _1595_v44;
        _1595_v44 = _dafny.Seq.UnicodeFromString("mjqmhkj");
        _1595_v44 = _1595_v44;
        let _1596_v45;
        _1596_v45 = _dafny.MultiSet.fromElements((_this).f10, (_dafny.ZERO).minus(new BigNumber((_1575_v25).length)), (_this).f10);
        let _index290 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length));
        (_1570_v23)[_index290] = ((_this).f10).isEqualTo(new BigNumber((_1596_v45).cardinality()));
        (globalState).f4 = p0;
        (globalState).f3 = (_this).f10;
      } else {
        let _1597_v46;
        _1597_v46 = new _dafny.CodePoint('i'.codePointAt(0));
        _1597_v46 = (((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]) ? (new _dafny.CodePoint('q'.codePointAt(0))) : (_1597_v46));
        let _index291 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_1566_v22).length));
        (_1566_v22)[_index291] = (_this).f11;
        let _index292 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_1566_v22).length));
        (_1566_v22)[_index292] = (_this).f11;
        let _1598_v47;
        _1598_v47 = _dafny.Set.fromElements((_this).f10);
        let _1599_v48;
        _1599_v48 = _module.D12.create_DC27((_this).f10, new BigNumber((_1598_v47).length), (_this).f10, (_this).f10);
        let _1600_v49;
        _1600_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f10)).cardinality()),(_1599_v48).dtor_cf39);
        let _1601_v50;
        _1601_v50 = _module.D3.create_DC8(_1600_v49);
        (globalState).f8 = _module.__default.fm26(_1601_v50, (_this).f10, _module.__default.fm44((_this).f10, globalState), globalState);
        (globalState).f4 = ((p0) ? ((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]) : ((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))]));
        let _1602_v51;
        _1602_v51 = _dafny.Seq.of((_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))], p0);
        let _1603_v52;
        _1603_v52 = _dafny.Seq.of(_1602_v51, _1602_v51, _1602_v51);
        let _1604_v53;
        _1604_v53 = _dafny.Set.fromElements((_1603_v52)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(false, p0, p0, (_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))], (_1570_v23)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1570_v23).length))])).length), new BigNumber((_1603_v52).length))]);
        (globalState).f3 = (new BigNumber((_1604_v53).length)).multipliedBy(new BigNumber(((_this).f11).length));
      }
      let _1605_v54;
      _1605_v54 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1606_v55;
      _1606_v55 = _module.D2.create_DC7(_1605_v54, true, (_this).f10);
      let _1607_v56;
      _1607_v56 = _dafny.Set.fromElements((_this).f10, (_1606_v55).dtor_cf11);
      (globalState).f8 = new BigNumber((_1607_v56).length);
      return;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _1608_v0;
      _1608_v0 = true;
      let _1609_v1;
      let _nw253 = new _module.C1();
      _nw253.__ctor();
      _1609_v1 = _nw253;
      let _1610_v2;
      _1610_v2 = _dafny.Seq.of(_1609_v1);
      let _1611_v3;
      _1611_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1608_v0,_1610_v2);
      (globalState).f8 = _module.__default.safeModuloInt(new BigNumber(((_this).f11).length), new BigNumber(((((_1611_v3).contains(_1608_v0)) ? ((_1611_v3).get(_1608_v0)) : ((_1610_v2)))).length));
      let _1612_v4;
      _1612_v4 = _dafny.Seq.of(p0);
      let _1613_v5;
      let _nw254 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1613_v5 = _nw254;
      let _1614_v6;
      _1614_v6 = _dafny.Seq.of(_1613_v5, _1613_v5);
      let _1615_v7;
      _1615_v7 = _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1614_v6).length),new BigNumber(((_this).f11).length)));
      let _1616_v8;
      _1616_v8 = _dafny.MultiSet.fromElements(_module.__default.fm31(_1608_v0, _1608_v0, _1608_v0, globalState));
      (globalState).f3 = (_dafny.ZERO).minus((_1612_v4)[_module.__default.safeIndex(_module.__default.fm26(_1615_v7, (_this).f10, _1616_v8, globalState), new BigNumber((_1612_v4).length))]);
      let _1617_v9;
      let _nw255 = new _module.C3();
      _nw255.__ctor();
      _1617_v9 = _nw255;
      let _1618_i0;
      _1618_i0 = _dafny.ZERO;
      L10: {
        while (((_this).f10).isEqualTo((_dafny.ZERO).minus((new BigNumber(389)).multipliedBy((_this).f10)))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1618_i0)) {
              break L10;
            }
            _1618_i0 = (_1618_i0).plus(_dafny.ONE);
            let _1619_v10;
            let _init42 = ((_1620_p0) => function (_1621_i1) {
              return _dafny.Map.Empty.slice().updateUnsafe(!(false),_1620_p0);
            })(p0);
            let _nw256 = Array((new BigNumber(29)).toNumber());
            for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw256.length); _i0_42++) {
              _nw256[_i0_42] = _init42(new BigNumber(_i0_42));
            }
            _1619_v10 = _nw256;
            let _1622_v11;
            let _nw257 = new _module.C4();
            _nw257.__ctor(_1619_v10);
            _1622_v11 = _nw257;
            let _1623_v12;
            _1623_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1608_v0,(_this).f10);
            (globalState).f3 = _module.__default.safeDivisionInt(p0, (((_1623_v12).contains(_1608_v0)) ? ((_1623_v12).get(_1608_v0)) : ((_this).f10)));
            let _1624_v13;
            _1624_v13 = new _dafny.CodePoint('f'.codePointAt(0));
            let _1625_v14;
            _1625_v14 = _dafny.Map.Empty.slice().updateUnsafe((_1622_v11).fm13(_1608_v0, (_this).f11, _1608_v0, globalState),_1624_v13);
            _1625_v14 = (_1625_v14).update(_1608_v0, _1624_v13);
            let _1626_v15;
            _1626_v15 = _dafny.MultiSet.fromElements((_this).f10);
            let _1627_v16;
            _1627_v16 = _dafny.Seq.of(!((_1622_v11).fm3(_1612_v4, _1626_v15, _1608_v0, globalState)));
            let _1628_v17;
            _1628_v17 = _1608_v0;
            let _1629_v18;
            _1629_v18 = _module.D0.create_DC1((_1617_v9).fm15(new BigNumber((_1627_v16).length), globalState), (_dafny.ZERO).minus(((_1622_v11).fm12(_1608_v0, p0, (_this).f11, _1628_v17, globalState)).minus(p0)));
            let _1630_v19;
            _1630_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
            let _1631_v20;
            _1631_v20 = _dafny.Set.fromElements(p0);
            let _1632_v21;
            _1632_v21 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements((((_1626_v15).contains(p0)) ? ((_1626_v15).get(p0)) : (p0)), (_this).f10, new BigNumber((_1630_v19).length)), _1631_v20, _1631_v20);
            let _1633_v22;
            _1633_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_1624_v13, _1608_v0, (_this).f10),new BigNumber((_1632_v21).cardinality()));
            let _1634_v23;
            _1634_v23 = _dafny.Set.fromElements(((_this).f10).minus((_this).f10), new BigNumber((_1633_v22).length));
            let _pat_let_tv63 = _1626_v15;
            let _pat_let_tv64 = p0;
            let _pat_let_tv65 = _1626_v15;
            let _rhs287 = _1608_v0;
            let _rhs288 = !(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("wcumpsjw"), _1624_v13));
            let _rhs289 = function (_pat_let53_0) {
              return function (_1635_dt__update__tmp_h0) {
                return function (_pat_let54_0) {
                  return function (_1636_dt__update_hcf1_h0) {
                    return _module.D0.create_DC1(_1636_dt__update_hcf1_h0, (_1635_dt__update__tmp_h0).dtor_cf2);
                  }(_pat_let54_0);
                }((_dafny.ZERO).minus((((_pat_let_tv65).contains((_this).f10)) ? ((_pat_let_tv63).get((_this).f10)) : (_pat_let_tv64))));
              }(_pat_let53_0);
            }(_1629_v18);
            let _rhs290 = _1634_v23;
            let _lhs231 = globalState;
            _1608_v0 = _rhs287;
            _lhs231.f4 = _rhs288;
            _1629_v18 = _rhs289;
            _1631_v20 = _rhs290;
          }
        }
      }
      let _1637_v24;
      _1637_v24 = _dafny.Set.fromElements(_1608_v0, _1608_v0);
      let _1638_v25;
      _1638_v25 = _dafny.Seq.of(!((_1637_v24).IsProperSubsetOf(_1637_v24)));
      let _1639_v26;
      _1639_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1608_v0,_1608_v0);
      let _1640_i2;
      _1640_i2 = _dafny.ZERO;
      L11: {
        while ((_1638_v25)[_module.__default.safeIndex(new BigNumber((_1639_v26).length), new BigNumber((_1638_v25).length))]) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1640_i2)) {
              break L11;
            }
            _1640_i2 = (_1640_i2).plus(_dafny.ONE);
            let _1641_v27;
            let _init43 = ((_1642_v0) => function (_1643_i3) {
              return _1642_v0;
            })(_1608_v0);
            let _nw258 = Array((new BigNumber(18)).toNumber());
            for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw258.length); _i0_43++) {
              _nw258[_i0_43] = _init43(new BigNumber(_i0_43));
            }
            _1641_v27 = _nw258;
            let _index293 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1641_v27).length));
            (_1641_v27)[_index293] = _1608_v0;
            let _1644_v28;
            _1644_v28 = _dafny.MultiSet.fromElements((_this).f10);
            let _index294 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1641_v27).length));
            (_1641_v27)[_index294] = ((_1644_v28).Intersect(_1644_v28)).IsDisjointFrom(_1644_v28);
            (globalState).f8 = _module.__default.safeModuloInt(new BigNumber(463), new BigNumber(((_this).f11).length));
            let _1645_v29;
            _1645_v29 = _dafny.Seq.of((_dafny.ZERO).minus(p0));
            let _source28 = _1645_v29;
            let _1646___mcc_h0 = _source28;
            let _1647_cf24 = _1646___mcc_h0;
            let _1648_v30;
            let _nw259 = Array((new BigNumber(9)).toNumber()).fill(_module.D2.Default());
            _1648_v30 = _nw259;
            let _1649_v31;
            _1649_v31 = _dafny.MultiSet.fromElements(_1648_v30);
            let _1650_v32;
            _1650_v32 = _1649_v31;
            _1650_v32 = _1650_v32;
            r0 = _1644_v28;
            let _1651_v33;
            _1651_v33 = _dafny.Seq.UnicodeFromString("uoxvu");
            let _1652_v34;
            _1652_v34 = new _dafny.CodePoint('e'.codePointAt(0));
            let _1653_v35;
            _1653_v35 = _module.D10.create_DC24(p0, (_1641_v27)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_1641_v27).length))], _dafny.Seq.of(_1652_v34));
            _1651_v33 = (_1653_v35).dtor_cf33;
            (globalState).f8 = (_dafny.ZERO).minus(((_this).f10).multipliedBy(_module.__default.safeModuloInt(new BigNumber(((_this).f11).length), new BigNumber((_1651_v33).length))));
            let _1654_v36;
            let _nw260 = Array((new BigNumber(18)).toNumber()).fill([]);
            _1654_v36 = _nw260;
            let _1655_v37;
            let _nw261 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _1655_v37 = _nw261;
            let _index295 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_1654_v36).length));
            (_1654_v36)[_index295] = _1655_v37;
            let _index296 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_1654_v36).length));
            (_1654_v36)[_index296] = _1655_v37;
          }
        }
      }
      let _1656_v38;
      _1656_v38 = _1610_v2;
      let _source29 = _1656_v38;
      let _1657___mcc_h1 = _source29;
      let _1658_cf61 = _1657___mcc_h1;
      if (_1608_v0) {
        let _1659_v39;
        _1659_v39 = new _dafny.CodePoint('n'.codePointAt(0));
        _1659_v39 = _1659_v39;
        (globalState).f3 = p0;
        _1612_v4 = _1612_v4;
        (globalState).f8 = ((_dafny.ZERO).minus((_this).f10)).plus((_this).f10);
        let _1660_v40;
        _1660_v40 = _dafny.Seq.UnicodeFromString("nujca");
        let _rhs291 = _module.__default.safeModuloInt(new BigNumber(-550), new BigNumber((_dafny.Seq.update((_this).f11, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1660_v40).length)), new BigNumber(((_this).f11).length)), _1659_v39)).length));
        let _rhs292 = _dafny.Seq.update((_this).f11, _module.__default.safeIndex(_module.__default.fm26(_1615_v7, p0, _module.__default.fm44(p0, globalState), globalState), new BigNumber(((_this).f11).length)), _1659_v39);
        let _lhs232 = globalState;
        _lhs232.f3 = _rhs291;
        _1660_v40 = _rhs292;
      } else {
        let _1661_v41;
        _1661_v41 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1662_v42;
        _1662_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1661_v41,_1608_v0);
        (globalState).f0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), ((_1663_v24) => function (_1664_i4) {
          return new BigNumber((_1663_v24).length);
        })(_1637_v24)), _dafny.Seq.of((_this).f10, new BigNumber((_1662_v42).length), new BigNumber((_1638_v25).length)));
        let _1665_v43;
        _1665_v43 = _module.D5.create_DC13((_1639_v26).Merge(_1639_v26));
        _1665_v43 = _1665_v43;
        let _1666_v44;
        _1666_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f11).length),_1608_v0);
        let _1667_v45;
        _1667_v45 = _dafny.Set.fromElements(p0);
        let _1668_v46;
        _1668_v46 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)));
        _1666_v44 = (_1666_v44).update(new BigNumber(((_1667_v45).Intersect(_1667_v45)).length), (_1668_v46).equals(_dafny.MultiSet.fromElements(_1661_v41)));
        let _1669_v47;
        _1669_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f11);
        let _1670_v48;
        _1670_v48 = _dafny.Map.Empty.slice().updateUnsafe((((_1669_v47).contains((_this).f10)) ? ((_1669_v47).get((_this).f10)) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("t"), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(!(true), true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)), _1661_v41))),(_dafny.ZERO).minus(((_this).f10).plus((_this).f10)));
        _1670_v48 = _1670_v48;
        let _1671_v49;
        _1671_v49 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10, p0);
        let _1672_v50;
        let _nw262 = Array((new BigNumber(2)).toNumber());
        _nw262[(_dafny.ZERO).toNumber()] = _1608_v0;
        _nw262[(_dafny.ONE).toNumber()] = (_1609_v1).fm3(_dafny.Seq.update(_1612_v4, _module.__default.safeIndex(p0, new BigNumber((_1612_v4).length)), (_this).f10), _1671_v49, _1608_v0, globalState);
        _1672_v50 = _nw262;
        let _index297 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1672_v50).length));
        (_1672_v50)[_index297] = (_1617_v9).fm16(_1608_v0, globalState);
        let _index298 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1672_v50).length));
        (_1672_v50)[_index298] = !_dafny.Seq.contains(_1638_v25, _dafny.Seq.IsProperPrefixOf((_this).f11, (_this).f11));
      }
      _1608_v0 = !(_1608_v0);
      let _1673_v51;
      let _init44 = ((_1674_v0) => function (_1675_i5) {
        return _1674_v0;
      })(_1608_v0);
      let _nw263 = Array((new BigNumber(17)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw263.length); _i0_44++) {
        _nw263[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1673_v51 = _nw263;
      let _index299 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1673_v51).length));
      (_1673_v51)[_index299] = _1608_v0;
      let _index300 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1673_v51).length));
      (_1673_v51)[_index300] = _1608_v0;
      let _1676_v52;
      let _nw264 = new _module.C1();
      _nw264.__ctor();
      _1676_v52 = _nw264;
      let _1677_v53;
      let _nw265 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _1677_v53 = _nw265;
      let _1678_v54;
      _1678_v54 = _dafny.Seq.of(_1677_v53, _1677_v53);
      let _1679_v55;
      _1679_v55 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt((_1617_v9).fm15(p0, globalState), new BigNumber((_1678_v54).length)));
      r0 = _1679_v55;
      return r0;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_1680_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length),new BigNumber(529))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ctfhl")).length),new BigNumber((_dafny.Seq.UnicodeFromString("w")).length)))).Merge((_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-955),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("awvtjnrl")).length))))).dtor_cf12);
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("no")).length),false)).length), (_module.D0.create_DC0(new BigNumber((function () {
  let _coll93 = new _dafny.Set();
  for (const _compr_93 of _dafny.IntegerRange(new BigNumber(385), new BigNumber(693))) {
    let _1681_v0 = _compr_93;
    if (((new BigNumber(385)).isLessThanOrEqualTo(_1681_v0)) && ((_1681_v0).isLessThan(new BigNumber(693)))) {
      _coll93.add((_1681_v0).multipliedBy(new BigNumber(-762)));
    }
  }
  return _coll93;
}()).length))).dtor_cf0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(501)), function (_1682_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_1683_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length));
      })).length))).cardinality()), new BigNumber(300), new BigNumber(524))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(795))).length))))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(470),true)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length)), new BigNumber(-52), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(813),_module.D3.create_DC8(function () {
  let _coll94 = new _dafny.Map();
  for (const _compr_94 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(646),new BigNumber((_dafny.Seq.UnicodeFromString("fsmppmgyx")).length))).length),new BigNumber(95))).Keys.Elements) {
    let _1684_v1 = _compr_94;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(646),new BigNumber((_dafny.Seq.UnicodeFromString("fsmppmgyx")).length))).length),new BigNumber(95))).contains(_1684_v1)) {
      _coll94.push([(_1684_v1).minus(new BigNumber(-781)),new BigNumber(797)]);
    }
  }
  return _coll94;
}()))).length))).length)));
    };
    fm10(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(15)), function (_1685_i0) {
        return new BigNumber(-113);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(!(true)))).length))), _dafny.Seq.of(new BigNumber(993)))).length);
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1686_v0;
      _1686_v0 = new BigNumber(905);
      let _1687_v1;
      _1687_v1 = _module.D2.create_DC6(_1686_v0, false, true);
      let _1688_v2;
      _1688_v2 = false;
      let _1689_v3;
      _1689_v3 = _dafny.MultiSet.fromElements(_1687_v1, _1687_v1, _1687_v1, _1687_v1, _module.D2.create_DC6(_1686_v0, _1688_v2, _1688_v2));
      _1689_v3 = _1689_v3;
      _1688_v2 = _1688_v2;
      (globalState).f8 = _1686_v0;
      let _1690_v4;
      _1690_v4 = _1688_v2;
      if (function (_source30) {
        let _1691___mcc_h6 = _source30;
        let _1692_cf4 = _1691___mcc_h6;
        return _1692_cf4;
      }(_1690_v4)) {
        if ((true) && (_1688_v2)) {
          (globalState).f3 = (_dafny.ZERO).minus(_1686_v0);
          let _1693_v5;
          _1693_v5 = _dafny.Map.Empty.slice().updateUnsafe(!(_1688_v2),_1686_v0);
          let _1694_v6;
          _1694_v6 = _dafny.Seq.of(_1686_v0, _1686_v0, _1686_v0);
          let _1695_v7;
          _1695_v7 = _dafny.MultiSet.fromElements(_1686_v0);
          let _1696_v8;
          _1696_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1693_v5,(_this).fm3(_1694_v6, _1695_v7, _1688_v2, globalState));
          let _1697_v9;
          _1697_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1696_v8,_1686_v0);
          _1686_v0 = ((((_1688_v2) ? (!(!(_1688_v2))) : (!(_1688_v2)))) ? (new BigNumber((_module.__default.fm11(_1688_v2, _1688_v2, globalState)).length)) : ((_dafny.ZERO).minus((((_1697_v9).contains(_1696_v8)) ? ((_1697_v9).get(_1696_v8)) : (new BigNumber(-971))))));
          let _1698_v10;
          let _nw266 = Array((new BigNumber(5)).toNumber());
          _nw266[(_dafny.ZERO).toNumber()] = _1693_v5;
          _nw266[(_dafny.ONE).toNumber()] = _1693_v5;
          _nw266[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,new BigNumber(518));
          _nw266[(new BigNumber(3)).toNumber()] = _1693_v5;
          _nw266[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1686_v0);
          _1698_v10 = _nw266;
          let _1699_v11;
          let _nw267 = new _module.C4();
          _nw267.__ctor(_1698_v10);
          _1699_v11 = _nw267;
          let _1700_v12;
          let _nw268 = new _module.C4();
          _nw268.__ctor(_1698_v10);
          _1700_v12 = _nw268;
          let _1701_v13;
          _1701_v13 = _dafny.Seq.UnicodeFromString("nph");
          let _1702_v14;
          _1702_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,(_1699_v11).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), ((_1703_v13) => function (_1704_i0) {
            return new BigNumber((_dafny.Seq.of(_1703_v13, _1703_v13, _dafny.Seq.UnicodeFromString("fmigiwsi"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(211)), function (_1705_i1) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            }), _1703_v13)).length);
          })(_1701_v13)), _1695_v7, _1688_v2, globalState));
          let _rhs293 = _1700_v12;
          let _rhs294 = !(_1688_v2);
          let _rhs295 = (_dafny.ZERO).minus(((_1686_v0).plus(new BigNumber((_1701_v13).length))).minus(new BigNumber(((_1702_v14).Merge(_module.__default.fm21(_1688_v2, globalState))).length)));
          let _rhs296 = _1686_v0;
          let _lhs233 = globalState;
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          _1700_v12 = _rhs293;
          _lhs233.f4 = _rhs294;
          _lhs234.f8 = _rhs295;
          _lhs235.f8 = _rhs296;
          let _1706_v16;
          let _init45 = ((_1707_v2, _1708_v0) => function (_1709_i2) {
            return _dafny.Set.fromElements(function () {
              let _coll95 = new _dafny.Set();
              for (const _compr_95 of (_dafny.Map.Empty.slice().updateUnsafe(_1708_v0,(_dafny.ZERO).minus(_1708_v0))).Keys.Elements) {
                let _1710_v15 = _compr_95;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_1708_v0,(_dafny.ZERO).minus(_1708_v0))).contains(_1710_v15)) {
                  _coll95.add(_module.__default.safeModuloInt(_1710_v15, new BigNumber(612)));
                }
              }
              return _coll95;
            }(), _dafny.Set.fromElements(_1708_v0), _dafny.Set.fromElements(_1708_v0, new BigNumber((_dafny.Seq.of(_1707_v2, _1707_v2)).length)));
          })(_1688_v2, _1686_v0);
          let _nw269 = Array((new BigNumber(7)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw269.length); _i0_45++) {
            _nw269[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1706_v16 = _nw269;
          let _1711_v17;
          _1711_v17 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), (_dafny.ZERO).minus(_1686_v0), new BigNumber((_1701_v13).length));
          let _1712_v18;
          _1712_v18 = _dafny.Set.fromElements(_1711_v17);
          let _index301 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_1706_v16).length));
          (_1706_v16)[_index301] = _1712_v18;
          let _1713_v19;
          let _nw270 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1713_v19 = _nw270;
          let _index302 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1713_v19).length));
          (_1713_v19)[_index302] = (_1700_v12).fm3(_1694_v6, _dafny.MultiSet.fromElements(_1686_v0, _1686_v0), _1688_v2, globalState);
          let _index303 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_1706_v16).length));
          let _index304 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1713_v19).length));
          let _rhs297 = (((_1702_v14).contains(_1688_v2)) ? ((_1702_v14).get(_1688_v2)) : (!(_1688_v2)));
          let _rhs298 = _1700_v12;
          let _rhs299 = _1712_v18;
          let _rhs300 = _1688_v2;
          let _lhs236 = _1706_v16;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_1706_v16).length));
          let _lhs238 = _1713_v19;
          let _lhs239 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1713_v19).length));
          r0 = _rhs297;
          _1700_v12 = _rhs298;
          _lhs236[_lhs237] = _rhs299;
          _lhs238[_lhs239] = _rhs300;
        } else {
          let _1714_v20;
          _1714_v20 = _dafny.Set.fromElements(_1686_v0);
          let _1715_v21;
          _1715_v21 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_1688_v2, _1688_v2)).cardinality()));
          let _1716_v22;
          _1716_v22 = _dafny.Seq.of((_1714_v20).IsSubsetOf(_1714_v20), !(!((_dafny.MultiSet.FromArray(_dafny.Seq.of(_1686_v0, _1686_v0))).IsSubsetOf(_1715_v21))), !(_1688_v2));
          _1716_v22 = _1716_v22;
          let _1717_v23;
          let _nw271 = new _module.C1();
          _nw271.__ctor();
          _1717_v23 = _nw271;
          let _1718_v24;
          _1718_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1688_v2);
          let _1719_v25;
          _1719_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1717_v23,(((_1718_v24).contains(_1688_v2)) ? ((_1718_v24).get(_1688_v2)) : (_1688_v2)));
          let _1720_v26;
          _1720_v26 = _dafny.Seq.of(_1717_v23, _1717_v23);
          _1719_v25 = (_1719_v25).update((_1720_v26)[_module.__default.safeIndex(new BigNumber((_1714_v20).length), new BigNumber((_1720_v26).length))], !(_1688_v2));
          let _1721_v27;
          _1721_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1686_v0);
          let _1722_v28;
          _1722_v28 = _module.D12.create_DC26(_1721_v27);
          let _1723_v29;
          _1723_v29 = _dafny.Seq.of(_1718_v24, _1718_v24, _1718_v24);
          let _1724_v30;
          _1724_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1722_v28,_module.__default.fm20(_1688_v2, globalState));
          let _1725_v31;
          _1725_v31 = _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(_1722_v28,new BigNumber(((_1723_v29)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1723_v29).length))]).length))).update(_1722_v28, _1686_v0), ((_1688_v2) ? (_1724_v30) : (_1724_v30)));
          let _1726_v32;
          _1726_v32 = _dafny.Seq.UnicodeFromString("xhgx");
          let _1727_v33;
          _1727_v33 = _dafny.Seq.of(_1686_v0, new BigNumber(841));
          let _1728_v34;
          _1728_v34 = _1727_v33;
          let _1729_v35;
          _1729_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(599)), function (_1730_i3) {
            return new BigNumber(940);
          }),_1686_v0);
          let _1731_v36;
          _1731_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1688_v2);
          _1725_v31 = _module.__default.fm45(_1726_v32, ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_1728_v34)).length)))).minus(new BigNumber((_1729_v35).length)), new BigNumber((_1731_v36).length), globalState);
          let _1732_v37;
          let _nw272 = Array((new BigNumber(4)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = false;
          _nw272[(_dafny.ONE).toNumber()] = _1688_v2;
          _nw272[(new BigNumber(2)).toNumber()] = _1688_v2;
          _nw272[(new BigNumber(3)).toNumber()] = _1688_v2;
          _1732_v37 = _nw272;
          let _1733_v38;
          _1733_v38 = _module.D2.create_DC5(_1732_v37);
          let _1734_v39;
          _1734_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1733_v38).dtor_cf5,_1688_v2);
          let _1735_v40;
          let _nw273 = Array((new BigNumber(10)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = _1688_v2;
          _nw273[(_dafny.ONE).toNumber()] = !(_1688_v2);
          _nw273[(new BigNumber(2)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(3)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(4)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(5)).toNumber()] = (((_1734_v39).contains(_1732_v37)) ? ((_1734_v39).get(_1732_v37)) : (_1688_v2));
          _nw273[(new BigNumber(6)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(7)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(8)).toNumber()] = _1688_v2;
          _nw273[(new BigNumber(9)).toNumber()] = false;
          _1735_v40 = _nw273;
          let _1736_v41;
          _1736_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1735_v40,_1686_v0);
          _1736_v41 = (_1736_v41).update(_1732_v37, new BigNumber(-116));
          let _1737_v42;
          let _nw274 = new _module.C5();
          _nw274.__ctor();
          _1737_v42 = _nw274;
        }
        let _1738_v43;
        let _1739_v44;
        let _1740_v45;
        let _out26;
        let _out27;
        let _out28;
        let _outcollector6 = (_this).m5(_1688_v2, _1686_v0, new BigNumber((_module.__default.fm46(globalState)).length), _1686_v0, globalState);
        _out26 = _outcollector6[0];
        _out27 = _outcollector6[1];
        _out28 = _outcollector6[2];
        _1738_v43 = _out26;
        _1739_v44 = _out27;
        _1740_v45 = _out28;
        _1740_v45 = _1688_v2;
        let _1741_v46;
        _1741_v46 = _dafny.Seq.of(_1740_v45, _1688_v2, _1740_v45);
        let _1742_v47;
        _1742_v47 = _module.D4.create_DC10(_1741_v46);
        let _1743_v48;
        _1743_v48 = _dafny.Seq.of(_1742_v47);
        r1 = (new BigNumber((_1743_v48).length)).isLessThanOrEqualTo(_1686_v0);
        let _1744_v49;
        _1744_v49 = _module.D4.create_DC10(_1741_v46);
        let _1745_v50;
        _1745_v50 = _module.D4.create_DC12(_1744_v49);
        let _1746_v51;
        _1746_v51 = _module.D4.create_DC12(_1744_v49);
        let _source31 = ((_1688_v2) ? (_module.D4.create_DC12(_module.D4.create_DC12(_1744_v49))) : (_1746_v51));
        if (_source31.is_DC11) {
          let _1747___mcc_h0 = (_source31).cf14;
          let _1748___mcc_h1 = (_source31).cf15;
          let _1749___mcc_h2 = (_source31).cf16;
          let _1750___mcc_h3 = (_source31).cf17;
          let _1751_cf17 = _1750___mcc_h3;
          let _1752_cf16 = _1749___mcc_h2;
          let _1753_cf15 = _1748___mcc_h1;
          let _1754_cf14 = _1747___mcc_h0;
          let _1755_v52;
          let _nw275 = new _module.C5();
          _nw275.__ctor();
          _1755_v52 = _nw275;
          (globalState).f3 = (_1686_v0).plus(_1686_v0);
          (globalState).f4 = _1688_v2;
          let _1756_v53;
          _1756_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,(_1752_cf16).plus(_1752_cf16));
          let _1757_v54;
          _1757_v54 = _dafny.MultiSet.fromElements(_1686_v0);
          _1756_v53 = (_1756_v53).update((_1755_v52).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), ((_1758_v45, _1759_cf17) => function (_1760_i4) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1758_v45,_1759_cf17)).length);
          })(_1740_v45, _1751_cf17)), _1757_v54, _1688_v2, globalState), _1686_v0);
        } else if (_source31.is_DC10) {
          let _1761___mcc_h4 = (_source31).cf13;
          let _1762_cf13 = _1761___mcc_h4;
          let _1763_v55;
          let _nw276 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _1763_v55 = _nw276;
          let _1764_v56;
          let _nw277 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1764_v56 = _nw277;
          let _1765_v57;
          _1765_v57 = _dafny.Seq.of(_1764_v56, _1764_v56, _1764_v56);
          let _index305 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1763_v55).length));
          (_1763_v55)[_index305] = _1765_v57;
          let _index306 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1763_v55).length));
          (_1763_v55)[_index306] = _dafny.Seq.update(_1765_v57, _module.__default.safeIndex(new BigNumber(-642), new BigNumber((_1765_v57).length)), _1764_v56);
          let _1766_v58;
          let _nw278 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1766_v58 = _nw278;
          let _index307 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1766_v58).length));
          (_1766_v58)[_index307] = _1764_v56;
          let _index308 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1766_v58).length));
          (_1766_v58)[_index308] = _1764_v56;
          let _1767_v59;
          _1767_v59 = _dafny.Seq.UnicodeFromString("tkabsjr");
          let _1768_v60;
          _1768_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1686_v0).isLessThan(new BigNumber((_1767_v59).length)),_1740_v45);
          _1768_v60 = (_1768_v60).update(_1688_v2, _1688_v2);
          let _1769_v61;
          let _nw279 = Array((new BigNumber(24)).toNumber());
          _nw279[(_dafny.ZERO).toNumber()] = _1688_v2;
          _nw279[(_dafny.ONE).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(2)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(3)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(4)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(5)).toNumber()] = false;
          _nw279[(new BigNumber(6)).toNumber()] = true;
          _nw279[(new BigNumber(7)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(8)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(9)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(10)).toNumber()] = false;
          _nw279[(new BigNumber(11)).toNumber()] = true;
          _nw279[(new BigNumber(12)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(13)).toNumber()] = !(_1740_v45);
          _nw279[(new BigNumber(14)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(15)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(16)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(17)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(18)).toNumber()] = !(true);
          _nw279[(new BigNumber(19)).toNumber()] = !(_1740_v45);
          _nw279[(new BigNumber(20)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(21)).toNumber()] = _1740_v45;
          _nw279[(new BigNumber(22)).toNumber()] = _1688_v2;
          _nw279[(new BigNumber(23)).toNumber()] = false;
          _1769_v61 = _nw279;
          (_this).m6((_this).fm10(globalState), _1769_v61, (_1686_v0).multipliedBy(new BigNumber(454)), globalState);
        } else {
          let _1770___mcc_h5 = (_source31).cf18;
          let _1771_cf18 = _1770___mcc_h5;
          let _1772_v62;
          _1772_v62 = _dafny.Seq.UnicodeFromString("fwbkil");
          let _1773_v63;
          _1773_v63 = _module.D4.create_DC11(_1740_v45, _1740_v45, _1686_v0, _1772_v62);
          let _1774_v64;
          _1774_v64 = _dafny.MultiSet.fromElements(_1686_v0, _1686_v0);
          let _1775_v65;
          _1775_v65 = _dafny.Set.fromElements(_1688_v2, false);
          let _1776_v66;
          _1776_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1774_v64,(_1775_v65).Difference(_1775_v65));
          let _1777_v67;
          _1777_v67 = _module.D9.create_DC21(_1740_v45, _1741_v46, _1740_v45);
          let _rhs301 = _module.__default.fm30((_this).fm10(globalState), _1777_v67, _1688_v2, _1686_v0, globalState);
          let _rhs302 = _1776_v66;
          let _rhs303 = _module.__default.fm25(globalState);
          let _rhs304 = ((_1686_v0).minus((_dafny.ZERO).minus(_1686_v0))).plus(new BigNumber((_dafny.Seq.update((_1773_v63).dtor_cf17, _module.__default.safeIndex(_1686_v0, new BigNumber(((_1773_v63).dtor_cf17).length)), _1738_v43)).length));
          let _rhs305 = _1686_v0;
          let _lhs240 = globalState;
          let _lhs241 = globalState;
          _1773_v63 = _rhs301;
          _1776_v66 = _rhs302;
          _1772_v62 = _rhs303;
          _lhs240.f8 = _rhs304;
          _lhs241.f8 = _rhs305;
          let _1778_v68;
          let _nw280 = new _module.C2();
          _nw280.__ctor();
          _1778_v68 = _nw280;
          let _rhs306 = (_1741_v46)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1741_v46).length))];
          let _rhs307 = _1688_v2;
          let _rhs308 = _module.__default.safeDivisionInt(new BigNumber((_1772_v62).length), _1686_v0);
          let _rhs309 = _1775_v65;
          let _rhs310 = ((false) ? (_1740_v45) : (_1740_v45));
          let _lhs242 = globalState;
          let _lhs243 = globalState;
          let _lhs244 = globalState;
          _lhs242.f4 = _rhs306;
          _1740_v45 = _rhs307;
          _lhs243.f8 = _rhs308;
          _1775_v65 = _rhs309;
          _lhs244.f4 = _rhs310;
          (globalState).f4 = (_dafny.Set.fromElements(_1772_v62, _1772_v62)).contains(_1772_v62);
        }
      } else {
        let _1779_v69;
        _1779_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1688_v2);
        let _1780_v70;
        _1780_v70 = _dafny.Seq.of(_1779_v69, _1779_v69);
        let _1781_v71;
        let _nw281 = Array((new BigNumber(3)).toNumber());
        _nw281[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_1779_v69, _1779_v69, _1779_v69);
        _nw281[(_dafny.ONE).toNumber()] = _1780_v70;
        _nw281[(new BigNumber(2)).toNumber()] = _1780_v70;
        _1781_v71 = _nw281;
        let _index309 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1781_v71).length));
        (_1781_v71)[_index309] = _1780_v70;
        let _index310 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1781_v71).length));
        (_1781_v71)[_index310] = _dafny.Seq.Concat(_dafny.Seq.Concat(_1780_v70, _1780_v70), _1780_v70);
        (globalState).f8 = _1686_v0;
        let _1782_v72;
        let _nw282 = new _module.C1();
        _nw282.__ctor();
        _1782_v72 = _nw282;
        let _1783_v73;
        _1783_v73 = _dafny.Seq.of(_1688_v2, _1688_v2, false);
        let _1784_v74;
        _1784_v74 = _dafny.Seq.of(_1686_v0);
        let _1785_v75;
        _1785_v75 = _dafny.Seq.UnicodeFromString("i");
        let _1786_v76;
        _1786_v76 = _dafny.MultiSet.fromElements(new BigNumber((_1785_v75).length));
        let _1787_v77;
        _1787_v77 = _dafny.Set.fromElements(!((_1783_v73)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1783_v73).length))]));
        let _1788_v78;
        _1788_v78 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1789_v79;
        _1789_v79 = _dafny.MultiSet.fromElements(_1788_v78);
        let _1790_v80;
        let _nw283 = Array((new BigNumber(23)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = (_1783_v73)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1783_v73).length))];
        _nw283[(_dafny.ONE).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(2)).toNumber()] = !(_1688_v2);
        _nw283[(new BigNumber(3)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(4)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(5)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(6)).toNumber()] = ((_this).fm10(globalState)).isLessThanOrEqualTo(_1686_v0);
        _nw283[(new BigNumber(7)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(8)).toNumber()] = (_1782_v72).fm3(_1784_v74, _1786_v76, _1688_v2, globalState);
        _nw283[(new BigNumber(9)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(10)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(11)).toNumber()] = (_1787_v77).IsSubsetOf(_1787_v77);
        _nw283[(new BigNumber(12)).toNumber()] = true;
        _nw283[(new BigNumber(13)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(14)).toNumber()] = !(_1789_v79).equals(_1789_v79);
        _nw283[(new BigNumber(15)).toNumber()] = !((_1690_v4));
        _nw283[(new BigNumber(16)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(17)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(18)).toNumber()] = !(_1688_v2);
        _nw283[(new BigNumber(19)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(20)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(21)).toNumber()] = _1688_v2;
        _nw283[(new BigNumber(22)).toNumber()] = _1688_v2;
        _1790_v80 = _nw283;
        let _index311 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1790_v80).length));
        (_1790_v80)[_index311] = (_1688_v2) === ((_1782_v72).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-537)), ((_1791_v0) => function (_1792_i5) {
          return _1791_v0;
        })(_1686_v0)), _1786_v76, _1688_v2, globalState));
        let _index312 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_1790_v80).length));
        (_1790_v80)[_index312] = _dafny.areEqual(_dafny.Seq.Concat(_1785_v75, _dafny.Seq.update(_1785_v75, _module.__default.safeIndex(_1686_v0, new BigNumber((_1785_v75).length)), new _dafny.CodePoint('i'.codePointAt(0)))), _dafny.Seq.Concat(_module.__default.fm17(_1788_v78, globalState), _1785_v75));
        _1688_v2 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1784_v74, _1784_v74))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.update(_1784_v74, _module.__default.safeIndex(new BigNumber(-860), new BigNumber((_1784_v74).length)), _1686_v0)));
      }
      if (((_1688_v2) ? (_1688_v2) : (true))) {
        let _1793_v81;
        _1793_v81 = _dafny.Seq.UnicodeFromString("cfusokk");
        let _1794_v82;
        _1794_v82 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1795_v83;
        _1795_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1794_v82,_1688_v2);
        let _1796_v84;
        _1796_v84 = _dafny.Seq.of(new BigNumber((_1795_v83).length));
        let _1797_v85;
        _1797_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1688_v2,_1688_v2);
        let _1798_v86;
        _1798_v86 = _dafny.MultiSet.fromElements(new BigNumber((_1797_v85).length));
        let _1799_v87;
        let _nw284 = Array((new BigNumber(18)).toNumber());
        _nw284[(_dafny.ZERO).toNumber()] = _1688_v2;
        _nw284[(_dafny.ONE).toNumber()] = true;
        _nw284[(new BigNumber(2)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(3)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(4)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(5)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(6)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(7)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(8)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(9)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(10)).toNumber()] = false;
        _nw284[(new BigNumber(11)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(12)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(13)).toNumber()] = false;
        _nw284[(new BigNumber(14)).toNumber()] = (_this).fm3(_1796_v84, _1798_v86, _1688_v2, globalState);
        _nw284[(new BigNumber(15)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(16)).toNumber()] = _1688_v2;
        _nw284[(new BigNumber(17)).toNumber()] = _1688_v2;
        _1799_v87 = _nw284;
        let _1800_v88;
        _1800_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1799_v87);
        let _1801_v89;
        _1801_v89 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1793_v81, _dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), function (_1802_i6) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })),_1800_v88);
        _1801_v89 = (_1801_v89).update(_1793_v81, _1800_v88);
        if (_1688_v2) {
          _1793_v81 = _dafny.Seq.Concat(_1793_v81, _1793_v81);
          (globalState).f3 = _1686_v0;
          _1793_v81 = _1793_v81;
          let _1803_v90;
          let _init46 = ((_1804_v82) => function (_1805_i7) {
            return _module.D10.create_DC23(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(93)), ((_1806_v82) => function (_1807_i8) {
  return _1806_v82;
})(_1804_v82)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-809)), ((_1808_v82) => function (_1809_i9) {
  return _1808_v82;
})(_1804_v82))));
          })(_1794_v82);
          let _nw285 = Array((new BigNumber(23)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw285.length); _i0_46++) {
            _nw285[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1803_v90 = _nw285;
          let _1810_v91;
          _1810_v91 = _dafny.MultiSet.fromElements(_1793_v81);
          let _1811_v92;
          _1811_v92 = _module.D10.create_DC23(_1810_v91);
          let _index313 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1803_v90).length));
          (_1803_v90)[_index313] = _1811_v92;
          let _index314 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1803_v90).length));
          (_1803_v90)[_index314] = _1811_v92;
          (globalState).f4 = (_1686_v0).isEqualTo(_module.__default.safeModuloInt(_1686_v0, new BigNumber((_module.__default.fm25(globalState)).length)));
        } else {
          (globalState).f3 = _1686_v0;
          (globalState).f8 = _1686_v0;
          let _1812_v93;
          _1812_v93 = _dafny.Set.fromElements(_1686_v0, _1686_v0, (new BigNumber((_1793_v81).length)).plus(_1686_v0), _1686_v0);
          (globalState).f8 = (_dafny.ZERO).minus(new BigNumber((_1812_v93).length));
          let _1813_v94;
          _1813_v94 = _dafny.Seq.of(_1688_v2);
          let _index315 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1799_v87).length));
          (_1799_v87)[_index315] = !(_dafny.Seq.contains(_1813_v94, (_1813_v94)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1813_v94).length))]));
          let _1814_v95;
          _1814_v95 = _dafny.Set.fromElements(_1688_v2, _1688_v2);
          let _index316 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1799_v87).length));
          let _rhs311 = ((_1688_v2) ? (_1688_v2) : (_1688_v2));
          let _rhs312 = _1814_v95;
          let _lhs245 = _1799_v87;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1799_v87).length));
          _lhs245[_lhs246] = _rhs311;
          _1814_v95 = _rhs312;
          _1797_v85 = _1797_v85;
        }
        let _rhs313 = true;
        let _rhs314 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1686_v0, (_1686_v0).minus((_dafny.ZERO).minus(_1686_v0))));
        let _lhs247 = globalState;
        _1688_v2 = _rhs313;
        _lhs247.f8 = _rhs314;
        let _1815_v96;
        let _nw286 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1815_v96 = _nw286;
        let _index317 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1815_v96).length));
        (_1815_v96)[_index317] = _1686_v0;
        let _index318 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1815_v96).length));
        (_1815_v96)[_index318] = _1686_v0;
        let _index319 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1815_v96).length));
        let _rhs315 = (_1815_v96)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_1815_v96).length))];
        let _rhs316 = _1815_v96;
        let _lhs248 = _1815_v96;
        let _lhs249 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1815_v96).length));
        _lhs248[_lhs249] = _rhs315;
        _1815_v96 = _rhs316;
      } else {
        let _1816_v97;
        let _nw287 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
        _1816_v97 = _nw287;
        let _1817_v98;
        _1817_v98 = _dafny.Set.fromElements(_1686_v0);
        let _1818_v99;
        _1818_v99 = _dafny.Seq.UnicodeFromString("xxcadue");
        let _1819_v100;
        _1819_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1817_v98,_1818_v99);
        let _index320 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1816_v97).length));
        (_1816_v97)[_index320] = _1819_v100;
        let _1820_v101;
        let _init47 = function (_1821_i10) {
          return _module.D7.create_DC18();
        };
        let _nw288 = Array((new BigNumber(17)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw288.length); _i0_47++) {
          _nw288[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1820_v101 = _nw288;
        let _1822_v102;
        _1822_v102 = _module.D7.create_DC18();
        let _index321 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1820_v101).length));
        (_1820_v101)[_index321] = ((_1688_v2) ? (_1822_v102) : (_1822_v102));
        let _index322 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1816_v97).length));
        let _index323 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1820_v101).length));
        let _rhs317 = ((_1819_v100).Merge(_module.__default.fm47(new BigNumber((_1818_v99).length), _1688_v2, _1688_v2, globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1817_v98,_dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), function (_1823_i11) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })));
        let _rhs318 = !(_1688_v2) || (_1688_v2);
        let _rhs319 = _1822_v102;
        let _lhs250 = _1816_v97;
        let _lhs251 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1816_v97).length));
        let _lhs252 = _1820_v101;
        let _lhs253 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1820_v101).length));
        _lhs250[_lhs251] = _rhs317;
        r0 = _rhs318;
        _lhs252[_lhs253] = _rhs319;
        let _1824_v103;
        let _nw289 = new _module.C1();
        _nw289.__ctor();
        _1824_v103 = _nw289;
        r1 = true;
        let _1825_v104;
        _1825_v104 = _dafny.Map.Empty.slice().updateUnsafe(!(_1688_v2) || (_1688_v2),_1686_v0);
        _1825_v104 = _1825_v104;
        let _1826_v105;
        _1826_v105 = _dafny.Seq.of(_1825_v104);
        let _1827_v106;
        _1827_v106 = _dafny.Seq.of(_1686_v0, _1686_v0, _1686_v0, _1686_v0, _1686_v0);
        let _1828_v107;
        _1828_v107 = _dafny.Seq.of((_1826_v105)[_module.__default.safeIndex(_1686_v0, new BigNumber((_1826_v105).length))], _1825_v104, _1825_v104, (_dafny.Map.Empty.slice().updateUnsafe(_1688_v2,new BigNumber((_1827_v106).length))).Merge(_1825_v104));
        let _1829_v108;
        _1829_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1686_v0,_1688_v2);
        _1828_v107 = _dafny.Seq.update(_1826_v105, _module.__default.safeIndex(new BigNumber(((_1829_v108).update(new BigNumber(507), _1688_v2)).length), new BigNumber((_1826_v105).length)), _module.__default.fm46(globalState));
      }
      let _1830_v109;
      let _init48 = function (_1831_i13) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      };
      let _nw290 = Array((new BigNumber(5)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw290.length); _i0_48++) {
        _nw290[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1830_v109 = _nw290;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1830_v109).length))) {
        let _1832_i12 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1832_i12)) && ((_1832_i12).isLessThan(new BigNumber((_1830_v109).length))))) {
          (_1830_v109)[(_1832_i12)] = new _dafny.CodePoint('h'.codePointAt(0));
        }
      }
      r0 = !(_1688_v2);
      let _1833_v110;
      _1833_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1687_v1,_1686_v0);
      let _1834_v111;
      _1834_v111 = _dafny.Set.fromElements(_1833_v110);
      r1 = (_1834_v111).IsProperSubsetOf(_1834_v111);
      return [r0, r1];
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.Set.Empty;
      let r2 = false;
      (globalState).f8 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), function (_1835_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length);
      let _1836_v0;
      _1836_v0 = _dafny.Set.fromElements(p2, p2);
      let _1837_v1;
      let _init49 = function (_1838_i2) {
        return true;
      };
      let _nw291 = Array((new BigNumber(5)).toNumber());
      for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw291.length); _i0_49++) {
        _nw291[_i0_49] = _init49(new BigNumber(_i0_49));
      }
      _1837_v1 = _nw291;
      let _1839_v2;
      _1839_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,(_dafny.ZERO).minus(p2));
      let _1840_i1;
      _1840_i1 = _dafny.ZERO;
      L12: {
        while ((_dafny.Set.fromElements(new BigNumber((_1839_v2).length))).IsSubsetOf((_1836_v0).Difference(_1836_v0))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1840_i1)) {
              break L12;
            }
            _1840_i1 = (_1840_i1).plus(_dafny.ONE);
            let _1841_v3;
            _1841_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(749),p3);
            let _1842_v4;
            _1842_v4 = _module.D3.create_DC8(_1841_v3);
            let _1843_v5;
            _1843_v5 = _dafny.Seq.of(p0, p0);
            let _1844_v6;
            _1844_v6 = _dafny.MultiSet.fromElements(_1843_v5);
            (globalState).f8 = (_module.__default.fm26(_1842_v4, p1, _1844_v6, globalState)).minus(p1);
            let _source32 = p0;
            let _1845___mcc_h0 = _source32;
            let _1846_cf4 = _1845___mcc_h0;
            let _1847_v7;
            _1847_v7 = _dafny.Seq.UnicodeFromString("cwtkmlf");
            let _1848_v8;
            _1848_v8 = _dafny.MultiSet.fromElements(_1847_v7);
            (globalState).f0 = _dafny.Seq.of(_module.__default.safeDivisionInt(p1, new BigNumber((_1848_v8).cardinality())), (((_1841_v3).contains(p1)) ? ((_1841_v3).get(p1)) : (p2)), _module.__default.fm20(p0, globalState), p3);
            (globalState).f4 = p0;
            let _nw292 = Array((new BigNumber(14)).toNumber()).fill(false);
            _1837_v1 = _nw292;
            _1847_v7 = _module.__default.fm14(p3, globalState);
            let _index324 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1837_v1).length));
            (_1837_v1)[_index324] = p0;
            let _index325 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1837_v1).length));
            (_1837_v1)[_index325] = p0;
            let _1849_v9;
            _1849_v9 = new _dafny.CodePoint('e'.codePointAt(0));
            let _1850_v10;
            _1850_v10 = _dafny.Seq.UnicodeFromString("jrrugumfh");
            let _1851_v11;
            _1851_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1849_v9,_1850_v10);
            let _1852_v12;
            _1852_v12 = _dafny.Seq.of(p2, p2, new BigNumber(((((_1851_v11).contains(_1849_v9)) ? ((_1851_v11).get(_1849_v9)) : (_1850_v10))).length), p2, new BigNumber((_1850_v10).length));
            (globalState).f0 = _dafny.Seq.Concat(_1852_v12, _1852_v12);
          }
        }
      }
      let _1853_v13;
      _1853_v13 = _dafny.Seq.UnicodeFromString("uba");
      let _1854_v14;
      _1854_v14 = _dafny.Set.fromElements(_1853_v13, _1853_v13);
      let _1855_v15;
      _1855_v15 = _module.D13.create_DC29(_1854_v14);
      let _1856_v16;
      _1856_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1855_v15);
      let _1857_v17;
      _1857_v17 = _module.D13.create_DC31((((_1856_v16).contains(p0)) ? ((_1856_v16).get(p0)) : (_1855_v15)));
      let _source33 = _1857_v17;
      if (_source33.is_DC30) {
        let _1858___mcc_h1 = (_source33).cf42;
        let _1859___mcc_h2 = (_source33).cf43;
        let _1860___mcc_h3 = (_source33).cf44;
        let _1861___mcc_h4 = (_source33).cf45;
        let _1862_cf45 = _1861___mcc_h4;
        let _1863_cf44 = _1860___mcc_h3;
        let _1864_cf43 = _1859___mcc_h2;
        let _1865_cf42 = _1858___mcc_h1;
        (globalState).f3 = ((p3).multipliedBy(p3)).plus(p1);
        _1853_v13 = _1853_v13;
        if (p0) {
          (globalState).f4 = !(p0) || (p0);
          let _1866_v18;
          _1866_v18 = _dafny.MultiSet.fromElements(new BigNumber(638));
          let _1867_v19;
          _1867_v19 = _dafny.Map.Empty.slice().updateUnsafe((p1).multipliedBy(p3),(_1866_v18).IsSubsetOf(_1866_v18));
          let _1868_v20;
          _1868_v20 = _dafny.Seq.of(p1);
          let _1869_v21;
          _1869_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1865_cf42,_dafny.MultiSet.fromElements(new BigNumber((_1868_v20).length)));
          _1867_v19 = (_1867_v19).update(_module.__default.safeDivisionInt(p3, p3), ((((_1869_v21).contains(_1864_cf43)) ? ((_1869_v21).get(_1864_cf43)) : (_dafny.MultiSet.FromArray(_1868_v20)))).IsProperSubsetOf(_1866_v18));
          let _1870_v22;
          _1870_v22 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),_1865_cf42);
          _1870_v22 = (_1870_v22).update(new _dafny.CodePoint('n'.codePointAt(0)), _1865_cf42);
          let _1871_v23;
          let _nw293 = new _module.C2();
          _nw293.__ctor();
          _1871_v23 = _nw293;
          let _1872_v24;
          _1872_v24 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm3(_1868_v20, _dafny.MultiSet.FromArray(_1868_v20), _1865_cf42, globalState)),new _dafny.CodePoint('s'.codePointAt(0)));
          let _1873_v25;
          _1873_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1865_cf42,_1853_v13);
          let _1874_v26;
          _1874_v26 = _module.D19.create_DC39(_1865_cf42, p1);
          let _1875_v27;
          let _nw294 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1875_v27 = _nw294;
          let _1876_v28;
          _1876_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1875_v27);
          let _1877_v29;
          _1877_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,p3);
          let _1878_v30;
          _1878_v30 = _dafny.Seq.of((((_1876_v28).contains((((_1877_v29).contains(_1864_cf43)) ? ((_1877_v29).get(_1864_cf43)) : (p3)))) ? ((_1876_v28).get((((_1877_v29).contains(_1864_cf43)) ? ((_1877_v29).get(_1864_cf43)) : (p3)))) : (_1875_v27)));
          let _1879_v31;
          _1879_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1874_v26).dtor_cf58,(_1878_v30)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,_1865_cf42)).length), new BigNumber((_1878_v30).length))]);
          let _1880_v32;
          _1880_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1879_v31,((p0) ? (_1871_v23) : (_1871_v23)));
          let _1881_v33;
          _1881_v33 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1882_v34;
          _1882_v34 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1881_v33);
          let _rhs320 = ((!(p2).isEqualTo(new BigNumber(((((_1873_v25).contains(!(p0))) ? ((_1873_v25).get(!(p0))) : (_1853_v13))).length))) ? (_1863_cf44) : (_1863_cf44));
          let _rhs321 = (((_1880_v32).contains(_1879_v31)) ? ((_1880_v32).get(_1879_v31)) : (((_1864_cf43) ? (_1871_v23) : (_1871_v23))));
          let _rhs322 = (_module.__default.fm43(_1864_cf43, p3, _1881_v33, globalState)).equals(_1882_v34);
          let _rhs323 = (_module.__default.fm11(_1865_cf42, _1864_cf43, globalState)).Merge(_1872_v24);
          let _rhs324 = p0;
          let _lhs254 = globalState;
          r1 = _rhs320;
          _1871_v23 = _rhs321;
          _lhs254.f4 = _rhs322;
          _1872_v24 = _rhs323;
          r2 = _rhs324;
          let _index326 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_1875_v27).length));
          (_1875_v27)[_index326] = new BigNumber(462);
          let _index327 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_1875_v27).length));
          (_1875_v27)[_index327] = ((false) ? ((new BigNumber(680)).minus(p3)) : (p2));
        } else {
          let _1883_v35;
          let _nw295 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1883_v35 = _nw295;
          let _1884_v36;
          _1884_v36 = _1883_v35;
          _1884_v36 = _1884_v36;
          let _1885_v37;
          let _nw296 = new _module.C2();
          _nw296.__ctor();
          _1885_v37 = _nw296;
          let _1886_v38;
          _1886_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1885_v37,_1853_v13);
          let _1887_v39;
          _1887_v39 = _dafny.MultiSet.fromElements(false);
          let _1888_v40;
          _1888_v40 = _dafny.Map.Empty.slice().updateUnsafe((p1).isLessThan(p1),(_1887_v39).IsProperSubsetOf(_1887_v39));
          let _index328 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length));
          (_1883_v35)[_index328] = (p2).multipliedBy(new BigNumber(-434));
          let _index329 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length));
          let _rhs325 = _1886_v38;
          let _rhs326 = p0;
          let _rhs327 = ((_1888_v40).Merge(_1888_v40)).update(!(_1864_cf43), (p2).isLessThan(p1));
          let _rhs328 = p2;
          let _lhs255 = globalState;
          let _lhs256 = _1883_v35;
          let _lhs257 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length));
          _1886_v38 = _rhs325;
          _lhs255.f4 = _rhs326;
          _1888_v40 = _rhs327;
          _lhs256[_lhs257] = _rhs328;
          _1865_cf42 = _1864_cf43;
          let _1889_v41;
          _1889_v41 = _dafny.Seq.of(new BigNumber(203), p2, p3);
          let _1890_v42;
          _1890_v42 = _dafny.MultiSet.fromElements(p2);
          let _index330 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length));
          (_1883_v35)[_index330] = ((_1889_v41)[_module.__default.safeIndex((((_1890_v42).contains((_1883_v35)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length))])) ? ((_1890_v42).get((_1883_v35)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length))])) : ((_1883_v35)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length))])), new BigNumber((_1889_v41).length))]).multipliedBy((new BigNumber(-356)).plus(p2));
          let _1891_v43;
          _1891_v43 = _1890_v42;
          let _1892_v44;
          _1892_v44 = _dafny.MultiSet.fromElements(_1891_v43, _1891_v43, _1891_v43);
          let _1893_v45;
          _1893_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1887_v39,new BigNumber((_1892_v44).cardinality()));
          let _1894_v46;
          _1894_v46 = _module.D12.create_DC27(p1, p2, new BigNumber(826), new BigNumber((_1893_v45).length));
          let _1895_v47;
          _1895_v47 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_1894_v46);
          let _1896_v48;
          _1896_v48 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1897_v49;
          _1897_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm31(_1865_cf42, p0, _1864_cf43, globalState)).length),p3);
          let _1898_v50;
          _1898_v50 = _module.D3.create_DC8(_1897_v49);
          _1895_v47 = (_1895_v47).update(_dafny.Seq.contains(_module.__default.fm1((_1883_v35)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_1883_v35).length))], _1896_v48, _1864_cf43, (_1898_v50).dtor_cf12, globalState), p2), _1894_v46);
        }
        (globalState).f8 = p2;
      } else if (_source33.is_DC29) {
        let _1899___mcc_h5 = (_source33).cf41;
        let _1900_cf41 = _1899___mcc_h5;
        (globalState).f8 = (_this).fm10(globalState);
        let _1901_v51;
        let _nw297 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1901_v51 = _nw297;
        let _1902_v52;
        let _nw298 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _1902_v52 = _nw298;
        let _index331 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1901_v51).length));
        (_1901_v51)[_index331] = _1902_v52;
        let _index332 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1901_v51).length));
        (_1901_v51)[_index332] = _1902_v52;
        let _1903_v53;
        let _init50 = ((_1904_p0) => function (_1905_i3) {
          return (_1905_i3).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_1904_p0)).length));
        })(p0);
        let _nw299 = Array((new BigNumber(23)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw299.length); _i0_50++) {
          _nw299[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1903_v53 = _nw299;
        let _1906_v54;
        _1906_v54 = _module.D21.create_DC42(_dafny.Map.Empty.slice().updateUnsafe(_1903_v53,new BigNumber(731)));
        (_this).m6(_module.__default.safeDivisionInt(new BigNumber(724), new BigNumber((_1853_v13).length)), _1837_v1, new BigNumber(((_1906_v54).dtor_cf62).length), globalState);
        (globalState).f8 = p1;
      } else {
        let _1907___mcc_h6 = (_source33).cf46;
        let _1908_cf46 = _1907___mcc_h6;
        _1853_v13 = _1853_v13;
        let _1909_v55;
        let _nw300 = new _module.C1();
        _nw300.__ctor();
        _1909_v55 = _nw300;
        let _1910_v56;
        _1910_v56 = _module.D4.create_DC11(p0, true, new BigNumber(-284), _1853_v13);
        let _1911_v57;
        _1911_v57 = _module.D4.create_DC12(_1910_v56);
        let _1912_v58;
        _1912_v58 = _module.D4.create_DC12(_1910_v56);
        let _1913_v59;
        _1913_v59 = _module.D4.create_DC12(_1910_v56);
        let _source34 = ((((p0) ? (false) : (p0))) ? (_1913_v59) : (_1913_v59));
        if (_source34.is_DC11) {
          let _1914___mcc_h7 = (_source34).cf14;
          let _1915___mcc_h8 = (_source34).cf15;
          let _1916___mcc_h9 = (_source34).cf16;
          let _1917___mcc_h10 = (_source34).cf17;
          let _1918_cf17 = _1917___mcc_h10;
          let _1919_cf16 = _1916___mcc_h9;
          let _1920_cf15 = _1915___mcc_h8;
          let _1921_cf14 = _1914___mcc_h7;
          let _index333 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1837_v1).length));
          (_1837_v1)[_index333] = false;
          let _index334 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1837_v1).length));
          (_1837_v1)[_index334] = true;
          let _1922_v60;
          let _init51 = ((_1923_cf14, _1924_p3, _1925_p2) => function (_1926_i4) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1923_cf14,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1924_p3,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_1925_p2)).length))).cardinality()))).length));
          })(_1921_cf14, p3, p2);
          let _nw301 = Array((new BigNumber(5)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw301.length); _i0_51++) {
            _nw301[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1922_v60 = _nw301;
          let _1927_v61;
          let _nw302 = new _module.C4();
          _nw302.__ctor(_1922_v60);
          _1927_v61 = _nw302;
          _1922_v60 = (((_1837_v1)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_1837_v1).length))]) ? (_1922_v60) : ((_1927_v61).f12));
          let _1928_v62;
          let _nw303 = new _module.C3();
          _nw303.__ctor();
          _1928_v62 = _nw303;
        } else if (_source34.is_DC10) {
          let _1929___mcc_h11 = (_source34).cf13;
          let _1930_cf13 = _1929___mcc_h11;
          let _1931_v63;
          _1931_v63 = new _dafny.CodePoint('k'.codePointAt(0));
          (globalState).f4 = !(new BigNumber((_dafny.Seq.update(_1853_v13, _module.__default.safeIndex(new BigNumber((_1853_v13).length), new BigNumber((_1853_v13).length)), _1931_v63)).length)).isEqualTo(p1);
          let _1932_v64;
          let _nw304 = new _module.C3();
          _nw304.__ctor();
          _1932_v64 = _nw304;
          let _1933_v65;
          _1933_v65 = _dafny.MultiSet.fromElements(_1932_v64);
          let _1934_v66;
          _1934_v66 = _dafny.MultiSet.fromElements(p0);
          let _1935_v67;
          _1935_v67 = _1934_v66;
          let _1936_v68;
          _1936_v68 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1935_v67);
          let _rhs329 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), function (_1937_i5) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          let _rhs330 = _1933_v65;
          let _rhs331 = ((_dafny.ZERO).minus(p3)).multipliedBy(new BigNumber((_1936_v68).length));
          let _lhs258 = globalState;
          _1853_v13 = _rhs329;
          _1933_v65 = _rhs330;
          _lhs258.f8 = _rhs331;
          let _1938_v69;
          _1938_v69 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
          let _1939_v70;
          _1939_v70 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0, p0));
          let _1940_v71;
          _1940_v71 = _dafny.MultiSet.fromElements(_module.__default.fm26(_module.D3.create_DC8(_1938_v69), p2, _1939_v70, globalState), p1, p1);
          (_1909_v55).m13(new BigNumber(768), _1940_v71, p0, globalState);
          let _rhs332 = !(p0);
          let _rhs333 = _module.__default.safeDivisionInt(p1, p3);
          let _rhs334 = _1931_v63;
          let _rhs335 = _module.__default.safeDivisionInt(p2, p3);
          let _lhs259 = globalState;
          let _lhs260 = globalState;
          let _lhs261 = globalState;
          _lhs259.f4 = _rhs332;
          _lhs260.f8 = _rhs333;
          r0 = _rhs334;
          _lhs261.f3 = _rhs335;
        } else {
          let _1941___mcc_h12 = (_source34).cf18;
          let _1942_cf18 = _1941___mcc_h12;
          let _1943_v72;
          let _nw305 = new _module.C1();
          _nw305.__ctor();
          _1943_v72 = _nw305;
          let _1944_v73;
          _1944_v73 = _module.D2.create_DC6(new BigNumber((_1836_v0).length), p0, p0);
          let _1945_v74;
          _1945_v74 = _dafny.Map.Empty.slice().updateUnsafe((_1944_v73).dtor_cf7,_1943_v72);
          _1943_v72 = ((p0) ? (_1943_v72) : ((((_1945_v74).contains(p0)) ? ((_1945_v74).get(p0)) : (_1943_v72))));
          let _1946_v75;
          let _nw306 = new _module.C3();
          _nw306.__ctor();
          _1946_v75 = _nw306;
          (globalState).f4 = ((p0) ? (p0) : (p0));
          let _1947_v76;
          _1947_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1909_v55,p1);
          let _1948_v77;
          _1948_v77 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
          let _1949_v78;
          _1949_v78 = _dafny.MultiSet.fromElements(p3, (_dafny.ZERO).minus(p2), p2, new BigNumber(934), new BigNumber((_1948_v77).length));
          let _1950_v79;
          let _nw307 = Array((new BigNumber(8)).toNumber());
          _nw307[(_dafny.ZERO).toNumber()] = (new BigNumber((_1947_v76).length)).plus(p2);
          _nw307[(_dafny.ONE).toNumber()] = p1;
          _nw307[(new BigNumber(2)).toNumber()] = p2;
          _nw307[(new BigNumber(3)).toNumber()] = (p1).minus(p3);
          _nw307[(new BigNumber(4)).toNumber()] = new BigNumber(206);
          _nw307[(new BigNumber(5)).toNumber()] = (new BigNumber(833)).minus(new BigNumber((_1853_v13).length));
          _nw307[(new BigNumber(6)).toNumber()] = new BigNumber((_1949_v78).cardinality());
          _nw307[(new BigNumber(7)).toNumber()] = p3;
          _1950_v79 = _nw307;
          let _index335 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1950_v79).length));
          (_1950_v79)[_index335] = p1;
          let _1951_v80;
          _1951_v80 = _dafny.Seq.of(p0);
          let _1952_v81;
          _1952_v81 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1951_v80).length));
          let _1953_v82;
          _1953_v82 = _dafny.Seq.of((_1946_v75).fm15(p2, globalState));
          let _1954_v83;
          _1954_v83 = _dafny.Seq.of((_1953_v82)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_1953_v82).length))], p3, new BigNumber(-919), (_dafny.ZERO).minus(p1), p1);
          let _index336 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1950_v79).length));
          (_1950_v79)[_index336] = (new BigNumber(((_1952_v81).update(new BigNumber((_1953_v82).length), new BigNumber(-642))).length)).minus((p1).plus(new BigNumber(160)));
        }
        let _1955_v84;
        _1955_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p0));
        _1955_v84 = (_1955_v84).update(true, p0);
      }
      let _1956_v85;
      _1956_v85 = _dafny.Seq.of(p1, new BigNumber(-829), p2, new BigNumber((_1853_v13).length));
      (globalState).f3 = (_1956_v85)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_1956_v85).length))];
      (globalState).f3 = new BigNumber(((_module.__default.fm28(p1, globalState)).Union(_1836_v0)).length);
      let _1957_v86;
      _1957_v86 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
      let _1958_v87;
      _1958_v87 = _module.D3.create_DC8(_1957_v86);
      let _1959_v89;
      _1959_v89 = _dafny.Seq.of(true);
      let _1960_v90;
      _1960_v90 = _dafny.MultiSet.fromElements(p1);
      let _1961_v91;
      _1961_v91 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_1959_v89, _module.__default.safeIndex(p3, new BigNumber((_1959_v89).length)), true), _module.__default.fm31(p0, p0, (_this).fm3(_1956_v85, _1960_v90, p0, globalState), globalState));
      let _1962_v92;
      _1962_v92 = new _dafny.CodePoint('v'.codePointAt(0));
      let _pat_let_tv66 = _1957_v86;
      let _1963_v93;
      let _nw308 = new _module.C6();
      _nw308.__ctor(p1, _dafny.Seq.update(_1853_v13, _module.__default.safeIndex(_module.__default.fm26(function (_pat_let55_0) {
        return function (_1964_dt__update__tmp_h0) {
          return function (_pat_let56_0) {
            return function (_1965_dt__update_hcf12_h0) {
              return _module.D3.create_DC8(_1965_dt__update_hcf12_h0);
            }(_pat_let56_0);
          }(_pat_let_tv66);
        }(_pat_let55_0);
      }(_1958_v87), new BigNumber((function () {
        let _coll96 = new _dafny.Map();
        for (const _compr_96 of _dafny.IntegerRange(new BigNumber(51), new BigNumber(-278))) {
          let _1966_v88 = _compr_96;
          if (((new BigNumber(51)).isLessThanOrEqualTo(_1966_v88)) && ((_1966_v88).isLessThan(new BigNumber(-278)))) {
            _coll96.push([_module.__default.safeModuloInt(_1966_v88, p2),true]);
          }
        }
        return _coll96;
      }()).length), _1961_v91, globalState), new BigNumber((_1853_v13).length)), _1962_v92));
      _1963_v93 = _nw308;
      r0 = _1962_v92;
      r1 = _1836_v0;
      r2 = (p3).isLessThan(p1);
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let _1967_v0;
      _1967_v0 = true;
      let _index337 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
      (p1)[_index337] = _1967_v0;
      let _index338 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
      (p1)[_index338] = false;
      let _1968_v1;
      _1968_v1 = new _dafny.CodePoint('m'.codePointAt(0));
      let _1969_v2;
      _1969_v2 = _dafny.MultiSet.fromElements(_1968_v1, _1968_v1, _1968_v1, _1968_v1);
      let _1970_v3;
      _1970_v3 = _dafny.Seq.of(_1968_v1);
      let _1971_v4;
      _1971_v4 = _1969_v2;
      let _1972_v5;
      _1972_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1970_v3).length),p2);
      let _1973_v6;
      _1973_v6 = _module.D3.create_DC8(_1972_v5);
      let _1974_v7;
      _1974_v7 = _dafny.Seq.of(_1967_v0, (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], _1967_v0);
      let _1975_v8;
      _1975_v8 = _dafny.MultiSet.fromElements(_1974_v7, _1974_v7);
      let _1976_v9;
      _1976_v9 = _module.D2.create_DC7(_1968_v1, false, p2);
      let _pat_let_tv67 = p1;
      let _pat_let_tv68 = p1;
      let _pat_let_tv69 = _1968_v1;
      let _index339 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
      let _rhs336 = (_1969_v2).Difference((_dafny.MultiSet.FromArray(_1970_v3)).Difference((_1971_v4)));
      let _rhs337 = (((_module.__default.fm26(_1973_v6, p0, _1975_v8, globalState)).isEqualTo(_module.__default.fm20((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState))) ? ((function (_pat_let57_0) {
        return function (_1977_dt__update__tmp_h0) {
          return function (_pat_let58_0) {
            return function (_1978_dt__update_hcf10_h0) {
              return function (_pat_let59_0) {
                return function (_1979_dt__update_hcf9_h0) {
                  return _module.D2.create_DC7(_1979_dt__update_hcf9_h0, _1978_dt__update_hcf10_h0, (_1977_dt__update__tmp_h0).dtor_cf11);
                }(_pat_let59_0);
              }(_pat_let_tv69);
            }(_pat_let58_0);
          }((_pat_let_tv68)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_pat_let_tv67).length))]);
        }(_pat_let57_0);
      }(_1976_v9)).dtor_cf10) : ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]));
      let _rhs338 = _1967_v0;
      let _lhs262 = p1;
      let _lhs263 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
      _1969_v2 = _rhs336;
      _1967_v0 = _rhs337;
      _lhs262[_lhs263] = _rhs338;
      let _index340 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
      (p1)[_index340] = !((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]);
      let _1980_v10;
      let _nw309 = new _module.C5();
      _nw309.__ctor();
      _1980_v10 = _nw309;
      let _1981_v11;
      _1981_v11 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))],_1980_v10);
      let _hi11 = p2;
      for (let _1982_i0 = new BigNumber((_1981_v11).length); _1982_i0.isLessThan(_hi11); _1982_i0 = _1982_i0.plus(_dafny.ONE)) {
        let _1983_v12;
        let _init52 = function (_1984_i1) {
          return (_1984_i1).minus(new BigNumber(310));
        };
        let _nw310 = Array((new BigNumber(18)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw310.length); _i0_52++) {
          _nw310[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1983_v12 = _nw310;
        let _index341 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1983_v12).length));
        (_1983_v12)[_index341] = new BigNumber(632);
        let _1985_v13;
        _1985_v13 = _module.D10.create_DC24(p0, _1967_v0, _1970_v3);
        let _index342 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1983_v12).length));
        (_1983_v12)[_index342] = _module.__default.safeModuloInt(_module.__default.fm26(_1973_v6, (_1985_v13).dtor_cf31, _1975_v8, globalState), p2);
        let _1986_v15;
        let _init53 = ((_1987_v1, _1988_p1) => function (_1989_i2) {
          return (function () {
            let _coll97 = new _dafny.Map();
            for (const _compr_97 of _dafny.IntegerRange(new BigNumber(-720), new BigNumber(104))) {
              let _1990_v14 = _compr_97;
              if (((new BigNumber(-720)).isLessThanOrEqualTo(_1990_v14)) && ((_1990_v14).isLessThan(new BigNumber(104)))) {
                _coll97.push([_module.__default.safeDivisionInt(_1990_v14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(858)), ((_1991_v1) => function (_1992_i3) {
                  return _1991_v1;
                })(_1987_v1))).length)),(_1988_p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_1988_p1).length))]]);
              }
            }
            return _coll97;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(27),(_1988_p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_1988_p1).length))]));
        })(_1968_v1, p1);
        let _nw311 = Array((new BigNumber(23)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw311.length); _i0_53++) {
          _nw311[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1986_v15 = _nw311;
        let _1993_v16;
        _1993_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1983_v12)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_1983_v12).length))],_1967_v0);
        let _1994_v17;
        _1994_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-945),(((_1993_v16).contains(new BigNumber((_1974_v7).length))) ? ((_1993_v16).get(new BigNumber((_1974_v7).length))) : (_1967_v0)));
        let _index343 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1986_v15).length));
        (_1986_v15)[_index343] = _1994_v17;
        let _index344 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1986_v15).length));
        let _index345 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
        let _rhs339 = _1993_v16;
        let _rhs340 = _1967_v0;
        let _rhs341 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_1995_v1) => function (_1996_i4) {
          return _1995_v1;
        })(_1968_v1)), ((_1967_v0) ? (_dafny.Seq.UnicodeFromString("fpl")) : (_1970_v3)));
        let _lhs264 = _1986_v15;
        let _lhs265 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1986_v15).length));
        let _lhs266 = p1;
        let _lhs267 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
        _lhs264[_lhs265] = _rhs339;
        _lhs266[_lhs267] = _rhs340;
        _1970_v3 = _rhs341;
        _1967_v0 = _1967_v0;
        let _1997_v18;
        let _nw312 = new _module.C3();
        _nw312.__ctor();
        _1997_v18 = _nw312;
      }
      let _1998_v19;
      _1998_v19 = _dafny.Set.fromElements(_1970_v3);
      if ((_1998_v19).IsDisjointFrom(_1998_v19)) {
        let _1999_v20;
        _1999_v20 = _dafny.Seq.of(p0);
        let _2000_v21;
        _2000_v21 = _dafny.MultiSet.fromElements(p2);
        let _2001_v22;
        _2001_v22 = _dafny.Set.fromElements(p2);
        let _2002_v23;
        let _nw313 = new _module.C0();
        _nw313.__ctor();
        _2002_v23 = _nw313;
        let _2003_v24;
        _2003_v24 = _module.D13.create_DC30((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], (_1980_v10).fm3(_1999_v20, _2000_v21, (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState), (_2001_v22).Difference(_2001_v22), _2002_v23);
        _2003_v24 = _module.D13.create_DC30((_2001_v22).IsSubsetOf(_2001_v22), (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], _2001_v22, _2002_v23);
        let _2004_v25;
        _2004_v25 = _dafny.MultiSet.fromElements(_1970_v3, _1970_v3);
        let _2005_v26;
        _2005_v26 = _module.D10.create_DC23(_2004_v25);
        let _source35 = _2005_v26;
        if (_source35.is_DC24) {
          let _2006___mcc_h0 = (_source35).cf31;
          let _2007___mcc_h1 = (_source35).cf32;
          let _2008___mcc_h2 = (_source35).cf33;
          let _2009_cf33 = _2008___mcc_h2;
          let _2010_cf32 = _2007___mcc_h1;
          let _2011_cf31 = _2006___mcc_h0;
          let _2012_v27;
          let _init54 = ((_2013_p2) => function (_2014_i5) {
            return (_2014_i5).multipliedBy(_2013_p2);
          })(p2);
          let _nw314 = Array((new BigNumber(2)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw314.length); _i0_54++) {
            _nw314[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _2012_v27 = _nw314;
          let _index346 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2012_v27).length));
          (_2012_v27)[_index346] = new BigNumber(-106);
          let _index347 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2012_v27).length));
          (_2012_v27)[_index347] = _2011_cf31;
          let _index348 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
          (p1)[_index348] = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
          (globalState).f8 = p2;
          let _2015_v28;
          let _init55 = ((_2016_v1) => function (_2017_i6) {
            return _2016_v1;
          })(_1968_v1);
          let _nw315 = Array((new BigNumber(19)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw315.length); _i0_55++) {
            _nw315[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2015_v28 = _nw315;
          let _2018_v29;
          _2018_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2015_v28,new BigNumber(459));
          _2018_v29 = _2018_v29;
        } else {
          let _2019___mcc_h3 = (_source35).cf30;
          let _2020_cf30 = _2019___mcc_h3;
          let _2021_v30;
          let _nw316 = new _module.C0();
          _nw316.__ctor();
          _2021_v30 = _nw316;
          _1967_v0 = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
          (globalState).f8 = p0;
          let _2022_v31;
          _2022_v31 = _module.D2.create_DC5(p1);
          _2022_v31 = _2022_v31;
        }
        let _2023_v32;
        let _nw317 = new _module.C1();
        _nw317.__ctor();
        _2023_v32 = _nw317;
        let _2024_v33;
        _2024_v33 = _module.D7.create_DC18();
        let _2025_v34;
        _2025_v34 = _dafny.Map.Empty.slice().updateUnsafe((_2003_v24).dtor_cf43,_2024_v33);
        let _2026_v35;
        _2026_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2001_v22).length),(p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]);
        if (!((_2025_v34).equals(_module.__default.fm48(_1999_v20, _2026_v35, globalState)))) {
          (globalState).f8 = p2;
          let _2027_v36;
          _2027_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, p0, p2, p2, new BigNumber((_1974_v7).length)),(_2000_v21).Intersect(_2000_v21));
          let _2028_v37;
          _2028_v37 = _dafny.Seq.of(_2000_v21);
          _2027_v36 = (_2027_v36).update(_1999_v20, (_2028_v37)[_module.__default.safeIndex(p0, new BigNumber((_2028_v37).length))]);
          let _index349 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
          (p1)[_index349] = ((_2023_v32).fm3(_1999_v20, _2000_v21, (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState));
          let _2029_v38;
          let _nw318 = Array((new BigNumber(27)).toNumber()).fill(false);
          _2029_v38 = _nw318;
          _2029_v38 = _2029_v38;
          let _2030_v39;
          let _nw319 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _2030_v39 = _nw319;
          let _index350 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_2030_v39).length));
          (_2030_v39)[_index350] = _1974_v7;
          let _index351 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_2030_v39).length));
          (_2030_v39)[_index351] = _1974_v7;
        } else {
          (globalState).f4 = true;
          let _2031_v40;
          _2031_v40 = _module.D19.create_DC39(false, new BigNumber((_1970_v3).length));
          let _2032_v41;
          _2032_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1968_v1,p2);
          (globalState).f8 = ((_dafny.ZERO).minus((_2031_v40).dtor_cf59)).multipliedBy((((_2032_v41).contains(_1968_v1)) ? ((_2032_v41).get(_1968_v1)) : (p2)));
          let _index352 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
          (p1)[_index352] = true;
          (globalState).f3 = (new BigNumber(39)).multipliedBy(_module.__default.fm20((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState));
          let _index353 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
          (p1)[_index353] = !(_2026_v35).contains(p2);
        }
        let _2033_v42;
        let _init56 = ((_2034_v0, _2035_p2) => function (_2036_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe(_2034_v0,_2035_p2);
        })(_1967_v0, p2);
        let _nw320 = Array((new BigNumber(14)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw320.length); _i0_56++) {
          _nw320[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _2033_v42 = _nw320;
        let _2037_v43;
        _2037_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1970_v3).length));
        let _index354 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_2033_v42).length));
        (_2033_v42)[_index354] = (_2037_v43).Merge(_2037_v43);
        let _index355 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_2033_v42).length));
        (_2033_v42)[_index355] = _2037_v43;
      } else {
        let _2038_v44;
        _2038_v44 = _dafny.Set.fromElements(_1967_v0, !(true), (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]);
        (globalState).f3 = new BigNumber(((_2038_v44).Union(_2038_v44)).length);
        let _2039_v45;
        let _nw321 = Array((new BigNumber(28)).toNumber()).fill([]);
        _2039_v45 = _nw321;
        let _2040_v46;
        _2040_v46 = _dafny.MultiSet.fromElements(new BigNumber(244));
        let _2041_v47;
        let _nw322 = Array((new BigNumber(21)).toNumber());
        _nw322[(_dafny.ZERO).toNumber()] = _2039_v45;
        _nw322[(_dafny.ONE).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(2)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(3)).toNumber()] = (((_this).fm3(_dafny.Seq.of(p2, p0, p0), _2040_v46, (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState)) ? (_2039_v45) : (_2039_v45));
        _nw322[(new BigNumber(4)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(5)).toNumber()] = ((!(_1967_v0)) ? (_2039_v45) : (_2039_v45));
        _nw322[(new BigNumber(6)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(7)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(8)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(9)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(10)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(11)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(12)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(13)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(14)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(15)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(16)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(17)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(18)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(19)).toNumber()] = _2039_v45;
        _nw322[(new BigNumber(20)).toNumber()] = _2039_v45;
        _2041_v47 = _nw322;
        let _index356 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2041_v47).length));
        (_2041_v47)[_index356] = _2039_v45;
        let _2042_v48;
        let _nw323 = Array((new BigNumber(11)).toNumber());
        _nw323[(_dafny.ZERO).toNumber()] = _1968_v1;
        _nw323[(_dafny.ONE).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(2)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(3)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(4)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(5)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(6)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(7)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(8)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(9)).toNumber()] = _1968_v1;
        _nw323[(new BigNumber(10)).toNumber()] = _1968_v1;
        _2042_v48 = _nw323;
        let _2043_v49;
        _2043_v49 = _dafny.Set.fromElements(p1);
        let _2044_v50;
        _2044_v50 = _dafny.Seq.of(_dafny.Set.fromElements(p1, p1), _2043_v49);
        let _2045_v51;
        _2045_v51 = _dafny.Seq.of(new BigNumber(((_2044_v50)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1974_v7)).cardinality()), new BigNumber((_2044_v50).length))]).length));
        let _index357 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2041_v47).length));
        let _rhs342 = _2039_v45;
        let _rhs343 = _dafny.Seq.update((((_1980_v10).fm3(_2045_v51, _dafny.MultiSet.fromElements(new BigNumber(436)), (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), ((_2046_v1) => function (_2047_i8) {
          return _2046_v1;
        })(_1968_v1))) : (_dafny.Seq.Concat(_1970_v3, _1970_v3))), _module.__default.safeIndex(new BigNumber((_2045_v51).length), new BigNumber(((((_1980_v10).fm3(_2045_v51, _dafny.MultiSet.fromElements(new BigNumber(436)), (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], globalState)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), ((_2048_v1) => function (_2049_i8) {
          return _2048_v1;
        })(_1968_v1))) : (_dafny.Seq.Concat(_1970_v3, _1970_v3)))).length)), _1968_v1);
        let _rhs344 = (_module.D23.create_DC45(_2042_v48)).dtor_cf66;
        let _lhs268 = _2041_v47;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2041_v47).length));
        _lhs268[_lhs269] = _rhs342;
        _1970_v3 = _rhs343;
        _2042_v48 = _rhs344;
        let _2050_v52;
        let _nw324 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _2050_v52 = _nw324;
        let _2051_v53;
        _2051_v53 = _2050_v52;
        let _source36 = _2051_v53;
        let _2052___mcc_h4 = _source36;
        let _2053_cf55 = _2052___mcc_h4;
        (globalState).f4 = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
        let _2054_v54;
        _2054_v54 = _dafny.Seq.of(p1);
        let _2055_v55;
        let _nw325 = Array((new BigNumber(16)).toNumber());
        _nw325[(_dafny.ZERO).toNumber()] = p1;
        _nw325[(_dafny.ONE).toNumber()] = p1;
        _nw325[(new BigNumber(2)).toNumber()] = p1;
        _nw325[(new BigNumber(3)).toNumber()] = p1;
        _nw325[(new BigNumber(4)).toNumber()] = p1;
        _nw325[(new BigNumber(5)).toNumber()] = p1;
        _nw325[(new BigNumber(6)).toNumber()] = p1;
        _nw325[(new BigNumber(7)).toNumber()] = p1;
        _nw325[(new BigNumber(8)).toNumber()] = (_2054_v54)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2054_v54).length))];
        _nw325[(new BigNumber(9)).toNumber()] = p1;
        _nw325[(new BigNumber(10)).toNumber()] = p1;
        _nw325[(new BigNumber(11)).toNumber()] = p1;
        _nw325[(new BigNumber(12)).toNumber()] = p1;
        _nw325[(new BigNumber(13)).toNumber()] = p1;
        _nw325[(new BigNumber(14)).toNumber()] = p1;
        _nw325[(new BigNumber(15)).toNumber()] = p1;
        _2055_v55 = _nw325;
        let _2056_v56;
        let _nw326 = Array((new BigNumber(13)).toNumber());
        _2056_v56 = _nw326;
        let _2057_v57;
        _2057_v57 = _dafny.Set.fromElements(_2056_v56, _2056_v56);
        let _rhs345 = _1970_v3;
        let _rhs346 = _2055_v55;
        let _rhs347 = (new BigNumber((_2057_v57).length)).minus((((_1972_v5).contains(p2)) ? ((_1972_v5).get(p2)) : (p2)));
        let _lhs270 = globalState;
        _1970_v3 = _rhs345;
        _2055_v55 = _rhs346;
        _lhs270.f8 = _rhs347;
        (globalState).f8 = p0;
        let _2058_v58;
        _2058_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sl"),_dafny.Seq.Create(_module.__default.abs(new BigNumber(711)), ((_2059_v1) => function (_2060_i9) {
          return _2059_v1;
        })(_1968_v1)));
        _1967_v0 = ((_1974_v7)[_module.__default.safeIndex(new BigNumber((_2058_v58).length), new BigNumber((_1974_v7).length))]) || (_1967_v0);
        (globalState).f4 = (_2040_v46).IsDisjointFrom(_2040_v46);
        let _2061_v59;
        _2061_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2038_v44);
        let _2062_v60;
        _2062_v60 = _dafny.Seq.of(p1);
        let _2063_v61;
        let _nw327 = Array((new BigNumber(22)).toNumber());
        _nw327[(_dafny.ZERO).toNumber()] = p1;
        _nw327[(_dafny.ONE).toNumber()] = p1;
        _nw327[(new BigNumber(2)).toNumber()] = p1;
        _nw327[(new BigNumber(3)).toNumber()] = p1;
        _nw327[(new BigNumber(4)).toNumber()] = p1;
        _nw327[(new BigNumber(5)).toNumber()] = p1;
        _nw327[(new BigNumber(6)).toNumber()] = (_2062_v60)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_2062_v60).length))];
        _nw327[(new BigNumber(7)).toNumber()] = p1;
        _nw327[(new BigNumber(8)).toNumber()] = (_2062_v60)[_module.__default.safeIndex(new BigNumber((_2045_v51).length), new BigNumber((_2062_v60).length))];
        _nw327[(new BigNumber(9)).toNumber()] = p1;
        _nw327[(new BigNumber(10)).toNumber()] = p1;
        _nw327[(new BigNumber(11)).toNumber()] = p1;
        _nw327[(new BigNumber(12)).toNumber()] = p1;
        _nw327[(new BigNumber(13)).toNumber()] = p1;
        _nw327[(new BigNumber(14)).toNumber()] = p1;
        _nw327[(new BigNumber(15)).toNumber()] = p1;
        _nw327[(new BigNumber(16)).toNumber()] = p1;
        _nw327[(new BigNumber(17)).toNumber()] = p1;
        _nw327[(new BigNumber(18)).toNumber()] = p1;
        _nw327[(new BigNumber(19)).toNumber()] = p1;
        _nw327[(new BigNumber(20)).toNumber()] = p1;
        _nw327[(new BigNumber(21)).toNumber()] = p1;
        _2063_v61 = _nw327;
        let _index358 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2063_v61).length));
        (_2063_v61)[_index358] = p1;
        let _2064_v63;
        _2064_v63 = _module.D12.create_DC27(new BigNumber((function () {
  let _coll98 = new _dafny.Map();
  for (const _compr_98 of _dafny.IntegerRange(new BigNumber(-216), new BigNumber(152))) {
    let _2065_v62 = _compr_98;
    if (((new BigNumber(-216)).isLessThanOrEqualTo(_2065_v62)) && ((_2065_v62).isLessThan(new BigNumber(152)))) {
      _coll98.push([(_2065_v62).minus(p2),(p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]]);
    }
  }
  return _coll98;
}()).length), p2, new BigNumber((_1970_v3).length), (_dafny.ZERO).minus(p0));
        let _index359 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2063_v61).length));
        let _rhs348 = (_2064_v63).dtor_cf39;
        let _rhs349 = ((_2061_v59).update(p0, _2038_v44)).Merge(_2061_v59);
        let _rhs350 = false;
        let _rhs351 = _1968_v1;
        let _rhs352 = p1;
        let _lhs271 = globalState;
        let _lhs272 = globalState;
        let _lhs273 = _2063_v61;
        let _lhs274 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2063_v61).length));
        _lhs271.f3 = _rhs348;
        _2061_v59 = _rhs349;
        _lhs272.f4 = _rhs350;
        _1968_v1 = _rhs351;
        _lhs273[_lhs274] = _rhs352;
      }
      let _2066_v65;
      _2066_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1968_v1,function () {
        let _coll99 = new _dafny.Map();
        for (const _compr_99 of (_1970_v3).Elements) {
          let _2067_v64 = _compr_99;
          if (_dafny.Seq.contains(_1970_v3, _2067_v64)) {
            _coll99.push([_2067_v64,(p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]]);
          }
        }
        return _coll99;
      }());
      let _2068_v66;
      _2068_v66 = _dafny.Seq.of(p2, p2, p2, new BigNumber((_2066_v65).length));
      if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), ((_2069_v66) => function (_2070_i10) {
        return _2069_v66;
      })(_2068_v66)), _2068_v66)) {
        (globalState).f8 = (_dafny.ZERO).minus(p2);
        let _2071_v67;
        _2071_v67 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))],p0);
        let _2072_v68;
        _2072_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2068_v66,_1980_v10);
        (globalState).f8 = (_module.__default.safeDivisionInt((((_2071_v67).contains(_1967_v0)) ? ((_2071_v67).get(_1967_v0)) : (_module.__default.fm26(_1973_v6, p2, _1975_v8, globalState))), new BigNumber((_1974_v7).length))).plus(_module.__default.safeDivisionInt((_this).fm10(globalState), new BigNumber((_2072_v68).length)));
        if ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]) {
          let _2073_v69;
          _2073_v69 = _module.D12.create_DC27(new BigNumber(852), p0, new BigNumber(886), p0);
          let _pat_let_tv70 = p0;
          let _2074_v70;
          let _nw328 = Array((new BigNumber(18)).toNumber());
          _nw328[(_dafny.ZERO).toNumber()] = _2073_v69;
          _nw328[(_dafny.ONE).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(2)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(3)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(4)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(5)).toNumber()] = _module.D12.create_DC27(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(201),p2)).length), new BigNumber((_1970_v3).length), _module.__default.fm20(true, globalState), p0);
          _nw328[(new BigNumber(6)).toNumber()] = _module.D12.create_DC27(new BigNumber((_1972_v5).length), new BigNumber((_1970_v3).length), p2, p2);
          _nw328[(new BigNumber(7)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(8)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(9)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(10)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(11)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(12)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(13)).toNumber()] = _module.D12.create_DC27(p2, new BigNumber((_2071_v67).length), p2, p0);
          _nw328[(new BigNumber(14)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(15)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(16)).toNumber()] = _2073_v69;
          _nw328[(new BigNumber(17)).toNumber()] = function (_pat_let60_0) {
            return function (_2075_dt__update__tmp_h2) {
              return function (_pat_let61_0) {
                return function (_2076_dt__update_hcf36_h0) {
                  return function (_pat_let62_0) {
                    return function (_2077_dt__update_hcf39_h0) {
                      return _module.D12.create_DC27(_2076_dt__update_hcf36_h0, (_2075_dt__update__tmp_h2).dtor_cf37, (_2075_dt__update__tmp_h2).dtor_cf38, _2077_dt__update_hcf39_h0);
                    }(_pat_let62_0);
                  }(_pat_let_tv70);
                }(_pat_let61_0);
              }(new BigNumber(183));
            }(_pat_let60_0);
          }(_2073_v69);
          _2074_v70 = _nw328;
          let _2078_v71;
          let _nw329 = Array((new BigNumber(14)).toNumber());
          _nw329[(_dafny.ZERO).toNumber()] = ((_1967_v0) ? (_2074_v70) : (_2074_v70));
          _nw329[(_dafny.ONE).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(2)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(3)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(4)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(5)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(6)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(7)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(8)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(9)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(10)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(11)).toNumber()] = _2074_v70;
          _nw329[(new BigNumber(12)).toNumber()] = (((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]) ? (_2074_v70) : (_2074_v70));
          _nw329[(new BigNumber(13)).toNumber()] = _2074_v70;
          _2078_v71 = _nw329;
          let _index360 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2078_v71).length));
          (_2078_v71)[_index360] = _2074_v70;
          let _index361 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2078_v71).length));
          (_2078_v71)[_index361] = _2074_v70;
          _2071_v67 = (_2071_v67).update(_1967_v0, p0);
          _1967_v0 = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
          (globalState).f8 = new BigNumber(((_dafny.Set.fromElements(_1968_v1)).Union(function () {
            let _coll100 = new _dafny.Set();
            for (const _compr_100 of (_1970_v3).Elements) {
              let _2079_v72 = _compr_100;
              if (_dafny.Seq.contains(_1970_v3, _2079_v72)) {
                _coll100.add(_2079_v72);
              }
            }
            return _coll100;
          }())).length);
          let _2080_v73;
          _2080_v73 = _dafny.Map.Empty.slice().updateUnsafe((p2).multipliedBy(new BigNumber((_2068_v66).length)),_module.__default.fm46(globalState));
          let _2081_v74;
          _2081_v74 = _dafny.MultiSet.fromElements(_1967_v0);
          _2080_v73 = (_2080_v73).update(_module.__default.safeDivisionInt(new BigNumber(600), new BigNumber((_2081_v74).cardinality())), (_2071_v67).update(_1967_v0, p0));
        } else {
          let _2082_v75;
          _2082_v75 = _dafny.MultiSet.fromElements(p2);
          let _2083_v76;
          _2083_v76 = _module.D4.create_DC10(_dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))], (_this).fm3(_2068_v66, _2082_v75, false, globalState), (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]));
          let _2084_v77;
          _2084_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1970_v3, _1970_v3),p0);
          let _rhs353 = _1970_v3;
          let _rhs354 = _module.D4.create_DC10(_dafny.Seq.update(_1974_v7, _module.__default.safeIndex(p0, new BigNumber((_1974_v7).length)), !(_1967_v0)));
          let _rhs355 = (((_2084_v77).contains(_1970_v3)) ? ((_2084_v77).get(_1970_v3)) : (p2));
          let _rhs356 = new BigNumber((((_1967_v0) ? (_dafny.Seq.Concat(_1970_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_2085_v1) => function (_2086_i11) {
            return _2085_v1;
          })(_1968_v1)))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), ((_2087_v1) => function (_2088_i12) {
            return _2087_v1;
          })(_1968_v1))))).length);
          let _lhs275 = globalState;
          let _lhs276 = globalState;
          _1970_v3 = _rhs353;
          _2083_v76 = _rhs354;
          _lhs275.f8 = _rhs355;
          _lhs276.f8 = _rhs356;
          _1976_v9 = _module.__default.fm40(_1967_v0, globalState);
          let _2089_v78;
          _2089_v78 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).fm10(globalState));
          _2089_v78 = _2089_v78;
          let _2090_v79;
          let _nw330 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2090_v79 = _nw330;
          let _index362 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_2090_v79).length));
          (_2090_v79)[_index362] = _1968_v1;
          let _2091_v80;
          let _nw331 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2091_v80 = _nw331;
          let _2092_v81;
          _2092_v81 = _dafny.Set.fromElements(p0, new BigNumber((_1970_v3).length), p0);
          let _index363 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_2090_v79).length));
          let _rhs357 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs358 = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
          let _rhs359 = p0;
          let _rhs360 = _module.__default.fm18((new BigNumber((_1970_v3).length)).isLessThanOrEqualTo(new BigNumber((_2092_v81).length)), _1967_v0, globalState);
          let _rhs361 = _2091_v80;
          let _lhs277 = globalState;
          let _lhs278 = _2090_v79;
          let _lhs279 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_2090_v79).length));
          _1968_v1 = _rhs357;
          _1967_v0 = _rhs358;
          _lhs277.f8 = _rhs359;
          _lhs278[_lhs279] = _rhs360;
          _2091_v80 = _rhs361;
          let _2093_v82;
          let _nw332 = new _module.C1();
          _nw332.__ctor();
          _2093_v82 = _nw332;
          _2093_v82 = _2093_v82;
        }
        (globalState).f4 = !(!_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dgljestqn"), _1970_v3), _dafny.Seq.Concat(_1970_v3, _1970_v3)));
        let _2094_v83;
        _2094_v83 = _dafny.Set.fromElements(!((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]), _1967_v0);
        let _2095_v84;
        _2095_v84 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1967_v0);
        let _2096_v85;
        _2096_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]);
        let _2097_v86;
        _2097_v86 = _dafny.Set.fromElements(p2);
        let _2098_v87;
        let _nw333 = new _module.C0();
        _nw333.__ctor();
        _2098_v87 = _nw333;
        let _2099_v88;
        _2099_v88 = _module.D13.create_DC30(_1967_v0, _1967_v0, _2097_v86, _2098_v87);
        let _2100_v89;
        _2100_v89 = _module.D13.create_DC31(_2099_v88);
        let _2101_v90;
        _2101_v90 = _dafny.Set.fromElements(_2100_v89);
        let _2102_v92;
        let _nw334 = Array((new BigNumber(20)).toNumber());
        _nw334[(_dafny.ZERO).toNumber()] = (_1967_v0) && (_1967_v0);
        _nw334[(_dafny.ONE).toNumber()] = (new BigNumber(-957)).isLessThanOrEqualTo(p2);
        _nw334[(new BigNumber(2)).toNumber()] = true;
        _nw334[(new BigNumber(3)).toNumber()] = (new BigNumber((_2094_v83).length)).isLessThanOrEqualTo(_module.__default.fm20(_1967_v0, globalState));
        _nw334[(new BigNumber(4)).toNumber()] = (((_2095_v84).contains(p1)) ? ((_2095_v84).get(p1)) : ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]));
        _nw334[(new BigNumber(5)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
        _nw334[(new BigNumber(6)).toNumber()] = _1967_v0;
        _nw334[(new BigNumber(7)).toNumber()] = (new BigNumber(695)).isLessThanOrEqualTo(p0);
        _nw334[(new BigNumber(8)).toNumber()] = true;
        _nw334[(new BigNumber(9)).toNumber()] = (((_2096_v85).contains(new BigNumber(-954))) ? ((_2096_v85).get(new BigNumber(-954))) : (_1967_v0));
        _nw334[(new BigNumber(10)).toNumber()] = false;
        _nw334[(new BigNumber(11)).toNumber()] = (p2).isLessThanOrEqualTo(p0);
        _nw334[(new BigNumber(12)).toNumber()] = (_2101_v90).IsSubsetOf(_2101_v90);
        _nw334[(new BigNumber(13)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
        _nw334[(new BigNumber(14)).toNumber()] = true;
        _nw334[(new BigNumber(15)).toNumber()] = _1967_v0;
        _nw334[(new BigNumber(16)).toNumber()] = true;
        _nw334[(new BigNumber(17)).toNumber()] = (new BigNumber((function () {
          let _coll101 = new _dafny.Set();
          for (const _compr_101 of _dafny.IntegerRange(new BigNumber(321), new BigNumber(692))) {
            let _2103_v91 = _compr_101;
            if (((new BigNumber(321)).isLessThanOrEqualTo(_2103_v91)) && ((_2103_v91).isLessThan(new BigNumber(692)))) {
              _coll101.add((_2103_v91).plus(p2));
            }
          }
          return _coll101;
        }()).length)).isLessThanOrEqualTo((_2068_v66)[_module.__default.safeIndex(p2, new BigNumber((_2068_v66).length))]);
        _nw334[(new BigNumber(18)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
        _nw334[(new BigNumber(19)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))];
        _2102_v92 = _nw334;
        let _2104_v93;
        _2104_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1970_v3,(p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]);
        let _2105_v94;
        _2105_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2104_v93,_1967_v0);
        let _index364 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
        let _rhs362 = (((_2105_v94).contains(_2104_v93)) ? ((_2105_v94).get(_2104_v93)) : ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]));
        let _rhs363 = _2102_v92;
        let _lhs280 = p1;
        let _lhs281 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length));
        _lhs280[_lhs281] = _rhs362;
        _2102_v92 = _rhs363;
      } else {
        _1967_v0 = _1967_v0;
        let _2106_v95;
        let _nw335 = new _module.C6();
        _nw335.__ctor(new BigNumber(966), _1970_v3);
        _2106_v95 = _nw335;
        (globalState).f8 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2106_v95).f10));
        let _2107_v96;
        _2107_v96 = _dafny.Set.fromElements(p0);
        (globalState).f3 = (p0).minus(_module.__default.safeDivisionInt(new BigNumber((_2107_v96).length), p2));
        let _2108_v97;
        let _nw336 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _2108_v97 = _nw336;
        let _2109_v98;
        _2109_v98 = _dafny.MultiSet.fromElements(_2108_v97);
        let _rhs364 = _dafny.MultiSet.fromElements(_2108_v97, _2108_v97, _2108_v97);
        let _rhs365 = _module.__default.fm20(((true) ? ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))]) : ((p1)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((p1).length))])), globalState);
        let _rhs366 = new BigNumber(-268);
        let _rhs367 = p2;
        let _lhs282 = globalState;
        let _lhs283 = globalState;
        let _lhs284 = globalState;
        _2109_v98 = _rhs364;
        _lhs282.f3 = _rhs365;
        _lhs283.f3 = _rhs366;
        _lhs284.f8 = _rhs367;
      }
      return;
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wsbmpw"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-960)), function (_2110_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(579), new BigNumber((_dafny.Seq.UnicodeFromString("tspj")).length))).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), function (_2111_i0) {
        return _dafny.Seq.UnicodeFromString("coabyxfb");
      })).length))).length)).plus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).length));
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let _hi12 = (p2).minus(p2);
      for (let _2112_i0 = (p2).plus(p2); _2112_i0.isLessThan(_hi12); _2112_i0 = _2112_i0.plus(_dafny.ONE)) {
        let _2113_v0;
        _2113_v0 = _module.D0.create_DC2();
        let _source37 = _2113_v0;
        if (_source37.is_DC1) {
          let _2114___mcc_h0 = (_source37).cf1;
          let _2115___mcc_h1 = (_source37).cf2;
          let _2116_cf2 = _2115___mcc_h1;
          let _2117_cf1 = _2114___mcc_h0;
          let _2118_v1;
          let _nw337 = Array((new BigNumber(4)).toNumber()).fill(_module.D0.Default());
          _2118_v1 = _nw337;
          let _2119_v2;
          _2119_v2 = _module.D0.create_DC0(new BigNumber(179));
          let _index365 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2118_v1).length));
          (_2118_v1)[_index365] = _2119_v2;
          let _2120_v3;
          _2120_v3 = new _dafny.CodePoint('n'.codePointAt(0));
          let _index366 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2118_v1).length));
          let _rhs368 = _2119_v2;
          let _rhs369 = _2120_v3;
          let _lhs285 = _2118_v1;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2118_v1).length));
          _lhs285[_lhs286] = _rhs368;
          _2120_v3 = _rhs369;
          let _2121_v4;
          let _nw338 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2121_v4 = _nw338;
          let _2122_v5;
          _2122_v5 = _dafny.Seq.UnicodeFromString("vvute");
          let _index367 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_2121_v4).length));
          (_2121_v4)[_index367] = _2122_v5;
          let _index368 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_2121_v4).length));
          (_2121_v4)[_index368] = _dafny.Seq.UnicodeFromString("rjcj");
          let _2123_v6;
          let _nw339 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2123_v6 = _nw339;
          let _index369 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_2123_v6).length));
          (_2123_v6)[_index369] = _2120_v3;
          let _index370 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_2123_v6).length));
          (_2123_v6)[_index370] = _2120_v3;
          let _2124_v7;
          _2124_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2116_cf2,p0);
          let _2125_v8;
          _2125_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v7,p0);
          (globalState).f4 = (p2).isLessThanOrEqualTo((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_2126_v3) => function (_2127_i1) {
            return _2126_v3;
          })(_2120_v3))).length)).multipliedBy(new BigNumber((_2125_v8).length)));
        } else if (_source37.is_DC2) {
          (globalState).f4 = (p0) || (p0);
          let _2128_v9;
          _2128_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _2129_v10;
          _2129_v10 = _dafny.MultiSet.fromElements(_2128_v9);
          let _2130_v11;
          _2130_v11 = _dafny.Set.fromElements(_2112_i0, _2112_i0);
          let _2131_v12;
          _2131_v12 = _dafny.MultiSet.fromElements(p0);
          let _2132_v13;
          _2132_v13 = _dafny.Seq.of(p2);
          let _2133_v14;
          let _nw340 = Array((new BigNumber(18)).toNumber());
          _nw340[(_dafny.ZERO).toNumber()] = true;
          _nw340[(_dafny.ONE).toNumber()] = (false) && (p0);
          _nw340[(new BigNumber(2)).toNumber()] = p0;
          _nw340[(new BigNumber(3)).toNumber()] = p0;
          _nw340[(new BigNumber(4)).toNumber()] = true;
          _nw340[(new BigNumber(5)).toNumber()] = p0;
          _nw340[(new BigNumber(6)).toNumber()] = (_2129_v10).contains(_2128_v9);
          _nw340[(new BigNumber(7)).toNumber()] = p0;
          _nw340[(new BigNumber(8)).toNumber()] = p0;
          _nw340[(new BigNumber(9)).toNumber()] = (_2130_v11).IsProperSubsetOf(_2130_v11);
          _nw340[(new BigNumber(10)).toNumber()] = (new BigNumber((_2130_v11).length)).isEqualTo(_2112_i0);
          _nw340[(new BigNumber(11)).toNumber()] = ((((_2131_v12).contains(false)) ? ((_2131_v12).get(false)) : (new BigNumber((_2132_v13).length)))).isLessThanOrEqualTo(_2112_i0);
          _nw340[(new BigNumber(12)).toNumber()] = p0;
          _nw340[(new BigNumber(13)).toNumber()] = !(p0) || (!(p0));
          _nw340[(new BigNumber(14)).toNumber()] = (p0) === (p0);
          _nw340[(new BigNumber(15)).toNumber()] = p0;
          _nw340[(new BigNumber(16)).toNumber()] = p0;
          _nw340[(new BigNumber(17)).toNumber()] = true;
          _2133_v14 = _nw340;
          let _index371 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_2133_v14).length));
          (_2133_v14)[_index371] = p0;
          let _2134_v15;
          _2134_v15 = p0;
          let _2135_v16;
          _2135_v16 = _dafny.Seq.of(p0, (_2134_v15));
          let _2136_v17;
          _2136_v17 = _module.D2.create_DC5(_2133_v14);
          let _2137_v18;
          _2137_v18 = _dafny.MultiSet.fromElements(new BigNumber((_2132_v13).length), new BigNumber(((_this).fm8(p2, p0, _2112_i0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), function (_2138_i2) {
            return _dafny.Seq.UnicodeFromString("iutevq");
          }), globalState)).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2133_v14,(_2136_v17).dtor_cf5)).update(_2133_v14, _2133_v14)).length));
          let _index372 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_2133_v14).length));
          (_2133_v14)[_index372] = (_2135_v16)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2137_v18).cardinality())), new BigNumber((_2135_v16).length))];
          let _index373 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_2133_v14).length));
          (_2133_v14)[_index373] = p0;
          (globalState).f3 = _module.__default.safeDivisionInt(((((_2137_v18).contains(p2)) ? ((_2137_v18).get(p2)) : (new BigNumber((_2132_v13).length)))).multipliedBy(_2112_i0), _2112_i0);
        } else if (_source37.is_DC0) {
          let _2139___mcc_h2 = (_source37).cf0;
          let _2140_cf0 = _2139___mcc_h2;
          let _2141_v19;
          _2141_v19 = _dafny.Seq.UnicodeFromString("dlnuduto");
          let _2142_v20;
          _2142_v20 = _module.D2.create_DC6(new BigNumber((_2141_v19).length), p0, p0);
          let _2143_v21;
          _2143_v21 = _module.D0.create_DC1((_2142_v20).dtor_cf6, new BigNumber(-464));
          let _2144_v22;
          _2144_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2143_v21);
          _2144_v22 = _2144_v22;
          (globalState).f3 = _2112_i0;
          let _2145_v23;
          let _nw341 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2145_v23 = _nw341;
          let _init57 = function (_2146_i3) {
            return _dafny.Seq.UnicodeFromString("irpd");
          };
          let _nw342 = Array((new BigNumber(29)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw342.length); _i0_57++) {
            _nw342[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2145_v23 = _nw342;
          let _2147_v24;
          _2147_v24 = _dafny.Set.fromElements(p0, p0);
          let _2148_v25;
          _2148_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2147_v24,(_this).fm9(p0, _2147_v24, globalState));
          _2148_v25 = (_2148_v25).update(_2147_v24, _2140_cf0);
        } else {
          let _2149___mcc_h3 = (_source37).cf3;
          let _2150_cf3 = _2149___mcc_h3;
          let _2151_v26;
          let _nw343 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _2151_v26 = _nw343;
          let _2152_v27;
          let _nw344 = new _module.C4();
          _nw344.__ctor(_2151_v26);
          _2152_v27 = _nw344;
          let _2153_v28;
          _2153_v28 = _dafny.MultiSet.fromElements(_2112_i0, new BigNumber(492));
          let _2154_v29;
          _2154_v29 = _dafny.Seq.UnicodeFromString("sjohkdo");
          let _2155_v30;
          _2155_v30 = _dafny.Seq.of(p0);
          let _2156_v31;
          _2156_v31 = _dafny.Seq.of(p2);
          let _2157_v32;
          _2157_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2153_v28).cardinality()),_2154_v29);
          let _2158_v33;
          let _nw345 = Array((new BigNumber(17)).toNumber());
          _nw345[(_dafny.ZERO).toNumber()] = (_2153_v28).contains(_2112_i0);
          _nw345[(_dafny.ONE).toNumber()] = p0;
          _nw345[(new BigNumber(2)).toNumber()] = (_2152_v27).fm13(p0, _2154_v29, p0, globalState);
          _nw345[(new BigNumber(3)).toNumber()] = !(!(!(p0)));
          _nw345[(new BigNumber(4)).toNumber()] = !((_2155_v30)[_module.__default.safeIndex(new BigNumber((_2156_v31).length), new BigNumber((_2155_v30).length))]) || (p0);
          _nw345[(new BigNumber(5)).toNumber()] = (p2).isEqualTo(_2112_i0);
          _nw345[(new BigNumber(6)).toNumber()] = (_2152_v27).fm13(p0, _2154_v29, p0, globalState);
          _nw345[(new BigNumber(7)).toNumber()] = true;
          _nw345[(new BigNumber(8)).toNumber()] = p0;
          _nw345[(new BigNumber(9)).toNumber()] = p0;
          _nw345[(new BigNumber(10)).toNumber()] = ((p0) ? (p0) : (p0));
          _nw345[(new BigNumber(11)).toNumber()] = p0;
          _nw345[(new BigNumber(12)).toNumber()] = p0;
          _nw345[(new BigNumber(13)).toNumber()] = p0;
          _nw345[(new BigNumber(14)).toNumber()] = p0;
          _nw345[(new BigNumber(15)).toNumber()] = p0;
          _nw345[(new BigNumber(16)).toNumber()] = (_2152_v27).fm13(p0, (((_2157_v32).contains(new BigNumber(823))) ? ((_2157_v32).get(new BigNumber(823))) : (_2154_v29)), p0, globalState);
          _2158_v33 = _nw345;
          let _2159_v34;
          _2159_v34 = _dafny.Set.fromElements(p0, true);
          let _rhs370 = new BigNumber(-61);
          let _rhs371 = (_dafny.Set.fromElements(p0, false, !(p0), p0, p0)).IsSubsetOf((_2159_v34).Intersect(_dafny.Set.fromElements((_2152_v27).fm3(_2156_v31, _dafny.MultiSet.fromElements(p2), p0, globalState))));
          let _rhs372 = _2158_v33;
          let _rhs373 = _2112_i0;
          let _lhs287 = globalState;
          let _lhs288 = globalState;
          let _lhs289 = globalState;
          _lhs287.f8 = _rhs370;
          _lhs288.f4 = _rhs371;
          _2158_v33 = _rhs372;
          _lhs289.f8 = _rhs373;
          let _2160_v35;
          _2160_v35 = _module.D4.create_DC10(_2155_v30);
          _2160_v35 = _2160_v35;
          let _2161_v36;
          _2161_v36 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p2, _2112_i0),p0);
          _2161_v36 = (_2161_v36).update(_2112_i0, !(p2).isEqualTo(_2112_i0));
        }
        let _2162_v37;
        _2162_v37 = new _dafny.CodePoint('s'.codePointAt(0));
        _2162_v37 = _2162_v37;
        _2162_v37 = _2162_v37;
        let _2163_v38;
        let _nw346 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
        _2163_v38 = _nw346;
        let _index374 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_2163_v38).length));
        (_2163_v38)[_index374] = _module.__default.fm31(false, !(p0), p0, globalState);
        let _2164_v39;
        _2164_v39 = _dafny.Seq.of(p0);
        let _2165_v40;
        _2165_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _index375 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_2163_v38).length));
        (_2163_v38)[_index375] = _dafny.Seq.Concat(_dafny.Seq.Concat(_2164_v39, _dafny.Seq.update(_dafny.Seq.update(_2164_v39, _module.__default.safeIndex(new BigNumber(-670), new BigNumber((_2164_v39).length)), (((_2165_v40).contains(p0)) ? ((_2165_v40).get(p0)) : (p0))), _module.__default.safeIndex(new BigNumber(330), new BigNumber((_dafny.Seq.update(_2164_v39, _module.__default.safeIndex(new BigNumber(-670), new BigNumber((_2164_v39).length)), (((_2165_v40).contains(p0)) ? ((_2165_v40).get(p0)) : (p0)))).length)), p0)), _dafny.Seq.Concat(_2164_v39, _2164_v39));
      }
      let _hi13 = (p2).multipliedBy(p2);
      for (let _2166_i4 = p2; _2166_i4.isLessThan(_hi13); _2166_i4 = _2166_i4.plus(_dafny.ONE)) {
        let _2167_v41;
        _2167_v41 = _dafny.Seq.of(p0);
        _2167_v41 = _2167_v41;
        (globalState).f8 = p2;
        let _2168_v42;
        _2168_v42 = new _dafny.CodePoint('a'.codePointAt(0));
        let _2169_v43;
        _2169_v43 = _dafny.Seq.UnicodeFromString("vstrg");
        (globalState).f4 = !_dafny.Seq.contains(_2169_v43, _2168_v42);
        let _2170_v44;
        let _nw347 = new _module.C6();
        _nw347.__ctor(_2166_i4, _dafny.Seq.Concat(_2169_v43, _2169_v43));
        _2170_v44 = _nw347;
      }
      (globalState).f4 = !(true);
      let _2171_v45;
      _2171_v45 = _dafny.Seq.of(p0, p0, p0, p0);
      let _2172_v46;
      _2172_v46 = _dafny.Seq.UnicodeFromString("am");
      let _2173_v47;
      _2173_v47 = new _dafny.CodePoint('b'.codePointAt(0));
      _2171_v45 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p0, p0), _2171_v45), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat((((p1).contains(p0)) ? ((p1).get(p0)) : (_2172_v46)), _2172_v46), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat((((p1).contains(p0)) ? ((p1).get(p0)) : (_2172_v46)), _2172_v46)).length)), _2173_v47)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, p0), _2171_v45)).length)), false);
      if (_module.__default.fm49(p0, globalState)) {
        let _2174_v48;
        let _init58 = function (_2175_i5) {
          return false;
        };
        let _nw348 = Array((new BigNumber(29)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw348.length); _i0_58++) {
          _nw348[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2174_v48 = _nw348;
        let _index376 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2174_v48).length));
        (_2174_v48)[_index376] = p0;
        let _index377 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_2174_v48).length));
        (_2174_v48)[_index377] = p0;
        let _2176_v49;
        _2176_v49 = _dafny.Set.fromElements(_2174_v48);
        let _index378 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2174_v48).length));
        let _index379 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_2174_v48).length));
        let _rhs374 = new BigNumber(((_dafny.Set.fromElements(_2174_v48)).Union(_2176_v49)).length);
        let _rhs375 = ((p0) ? (p2) : ((p2).minus(p2)));
        let _rhs376 = p0;
        let _rhs377 = p0;
        let _lhs290 = globalState;
        let _lhs291 = globalState;
        let _lhs292 = _2174_v48;
        let _lhs293 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2174_v48).length));
        let _lhs294 = _2174_v48;
        let _lhs295 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_2174_v48).length));
        _lhs290.f8 = _rhs374;
        _lhs291.f3 = _rhs375;
        _lhs292[_lhs293] = _rhs376;
        _lhs294[_lhs295] = _rhs377;
        let _2177_v50;
        let _init59 = ((_2178_p2) => function (_2179_i6) {
          return _module.__default.safeModuloInt(_2179_i6, _2178_p2);
        })(p2);
        let _nw349 = Array((new BigNumber(5)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw349.length); _i0_59++) {
          _nw349[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2177_v50 = _nw349;
        let _2180_v51;
        _2180_v51 = _2177_v50;
        let _2181_v52;
        let _nw350 = Array((new BigNumber(9)).toNumber());
        _nw350[(_dafny.ZERO).toNumber()] = _2180_v51;
        _nw350[(_dafny.ONE).toNumber()] = _2180_v51;
        _nw350[(new BigNumber(2)).toNumber()] = ((p0) ? (_2177_v50) : (_2180_v51));
        _nw350[(new BigNumber(3)).toNumber()] = _2180_v51;
        _nw350[(new BigNumber(4)).toNumber()] = _2177_v50;
        _nw350[(new BigNumber(5)).toNumber()] = _2177_v50;
        _nw350[(new BigNumber(6)).toNumber()] = _2180_v51;
        _nw350[(new BigNumber(7)).toNumber()] = _2180_v51;
        _nw350[(new BigNumber(8)).toNumber()] = _2180_v51;
        _2181_v52 = _nw350;
        let _index380 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_2181_v52).length));
        (_2181_v52)[_index380] = _2180_v51;
        let _index381 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_2181_v52).length));
        (_2181_v52)[_index381] = _2177_v50;
        let _2182_v53;
        _2182_v53 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        _2182_v53 = (_2182_v53).update(new BigNumber((_2171_v45).length), p2);
        let _2183_v54;
        _2183_v54 = _dafny.Set.fromElements(new BigNumber(-799), p2);
        let _2184_v55;
        _2184_v55 = _module.D3.create_DC9();
        let _2185_v56;
        _2185_v56 = _dafny.Set.fromElements(_module.__default.fm50(globalState), _2184_v55, _2184_v55);
        let _2186_v57;
        _2186_v57 = _dafny.Seq.of(_2185_v56, _2185_v56);
        (globalState).f8 = (new BigNumber(((_2183_v54).Union(_2183_v54)).length)).minus(((p0) ? (new BigNumber(((_2186_v57)[_module.__default.safeIndex(p2, new BigNumber((_2186_v57).length))]).length)) : (new BigNumber((_2171_v45).length))));
        let _2187_v58;
        let _nw351 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2187_v58 = _nw351;
        let _index382 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2187_v58).length));
        (_2187_v58)[_index382] = _dafny.Seq.Concat(_2172_v46, _2172_v46);
        let _index383 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2187_v58).length));
        (_2187_v58)[_index383] = _dafny.Seq.update(_2172_v46, _module.__default.safeIndex(p2, new BigNumber((_2172_v46).length)), _2173_v47);
      } else {
        let _2188_v59;
        _2188_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2172_v46,p2);
        let _2189_v60;
        _2189_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        _2188_v59 = (_2188_v59).update(_2172_v46, (_dafny.ZERO).minus((((_2189_v60).contains(p0)) ? ((_2189_v60).get(p0)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p2, p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(800)), ((_2190_v47) => function (_2191_i7) {
          return _2190_v47;
        })(_2173_v47))).length)), _module.__default.safeIndex(new BigNumber((_2172_v46).length), new BigNumber((_dafny.Seq.of(p2, p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(800)), ((_2192_v47) => function (_2193_i7) {
          return _2192_v47;
        })(_2173_v47))).length))).length)), p2)).length))))));
        let _2194_v61;
        let _nw352 = new _module.C2();
        _nw352.__ctor();
        _2194_v61 = _nw352;
        let _2195_v62;
        _2195_v62 = _dafny.Seq.of(_2194_v61);
        _2194_v61 = (_2195_v62)[_module.__default.safeIndex(p2, new BigNumber((_2195_v62).length))];
        _2172_v46 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tvyfq"), _2172_v46), _2172_v46);
        let _2196_v63;
        let _nw353 = new _module.C3();
        _nw353.__ctor();
        _2196_v63 = _nw353;
        let _2197_v64;
        let _nw354 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2197_v64 = _nw354;
        let _index384 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_2197_v64).length));
        (_2197_v64)[_index384] = _2172_v46;
        let _index385 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_2197_v64).length));
        (_2197_v64)[_index385] = _2172_v46;
      }
      let _2198_v65;
      let _nw355 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
      _2198_v65 = _nw355;
      let _2199_v66;
      let _nw356 = Array((new BigNumber(13)).toNumber());
      _nw356[(_dafny.ZERO).toNumber()] = _2173_v47;
      _nw356[(_dafny.ONE).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(2)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(3)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(4)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(5)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(6)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(7)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(8)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(9)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(10)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(11)).toNumber()] = _2173_v47;
      _nw356[(new BigNumber(12)).toNumber()] = _2173_v47;
      _2199_v66 = _nw356;
      let _2200_v67;
      _2200_v67 = _dafny.Set.fromElements(_2199_v66, _2199_v66);
      let _index386 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_2198_v65).length));
      (_2198_v65)[_index386] = _2200_v67;
      let _2201_v68;
      _2201_v68 = _module.D19.create_DC39(p0, new BigNumber(762));
      let _2202_v69;
      let _nw357 = Array((new BigNumber(6)).toNumber());
      _nw357[(_dafny.ZERO).toNumber()] = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("oxey"), _dafny.Seq.UnicodeFromString("brbh")));
      _nw357[(_dafny.ONE).toNumber()] = p0;
      _nw357[(new BigNumber(2)).toNumber()] = (_2201_v68).dtor_cf58;
      _nw357[(new BigNumber(3)).toNumber()] = p0;
      _nw357[(new BigNumber(4)).toNumber()] = !_dafny.areEqual(_module.__default.fm14(p2, globalState), _2172_v46);
      _nw357[(new BigNumber(5)).toNumber()] = false;
      _2202_v69 = _nw357;
      let _index387 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_2202_v69).length));
      (_2202_v69)[_index387] = p0;
      let _2203_v70;
      _2203_v70 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2200_v67);
      let _index388 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_2198_v65).length));
      let _index389 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_2202_v69).length));
      let _rhs378 = (_2200_v67).Difference((((_2203_v70).contains(p2)) ? ((_2203_v70).get(p2)) : (_2200_v67)));
      let _rhs379 = p0;
      let _rhs380 = (_module.__default.fm51(p2, globalState));
      let _lhs296 = _2198_v65;
      let _lhs297 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_2198_v65).length));
      let _lhs298 = globalState;
      let _lhs299 = _2202_v69;
      let _lhs300 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_2202_v69).length));
      _lhs296[_lhs297] = _rhs378;
      _lhs298.f4 = _rhs379;
      _lhs299[_lhs300] = _rhs380;
      return;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f9 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f9) {
      let _this = this;
      (_this)._f9 = f9;
      return;
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let _source38 = _module.__default.fm5(globalState);
      if (_source38.is_DC1) {
        let _2204___mcc_h0 = (_source38).cf1;
        let _2205___mcc_h1 = (_source38).cf2;
        let _2206_cf2 = _2205___mcc_h1;
        let _2207_cf1 = _2204___mcc_h0;
        let _2208_v0;
        let _nw358 = Array((new BigNumber(10)).toNumber()).fill(false);
        _2208_v0 = _nw358;
        _2208_v0 = _2208_v0;
        let _2209_v1;
        _2209_v1 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2210_v2;
        _2210_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2209_v1);
        let _2211_v3;
        _2211_v3 = _dafny.Map.Empty.slice().updateUnsafe((_2207_cf1).plus(_2207_cf1),_2210_v2);
        let _2212_v4;
        _2212_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2207_cf1);
        let _2213_v5;
        _2213_v5 = _dafny.Seq.of((_this).f9);
        _2211_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6((((_2212_v4).contains((_this).f9)) ? ((_2212_v4).get((_this).f9)) : (new BigNumber((_2213_v5).length))), globalState),_module.__default.fm7(_2209_v1, globalState));
        r1 = _2206_cf2;
        if ((_this).f9) {
          let _2214_v6;
          let _nw359 = new _module.C7();
          _nw359.__ctor();
          _2214_v6 = _nw359;
          let _2215_v7;
          let _nw360 = Array((new BigNumber(23)).toNumber());
          _2215_v7 = _nw360;
          let _2216_v8;
          let _nw361 = new _module.C0();
          _nw361.__ctor();
          _2216_v8 = _nw361;
          let _index390 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_2215_v7).length));
          (_2215_v7)[_index390] = _2216_v8;
          let _index391 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_2215_v7).length));
          (_2215_v7)[_index391] = _2216_v8;
          let _2217_v9;
          _2217_v9 = _dafny.Seq.of(_2207_cf1);
          (globalState).f4 = (_2214_v6).fm3(_2217_v9, _dafny.MultiSet.FromArray(_2217_v9), (_this).f9, globalState);
          let _2218_v10;
          let _nw362 = new _module.C8();
          _nw362.__ctor();
          _2218_v10 = _nw362;
          let _2219_v11;
          _2219_v11 = _module.D2.create_DC5(_2208_v0);
          let _2220_v12;
          let _nw363 = Array((new BigNumber(11)).toNumber());
          _nw363[(_dafny.ZERO).toNumber()] = _2208_v0;
          _nw363[(_dafny.ONE).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(2)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(3)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(4)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(5)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(6)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(7)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(8)).toNumber()] = _2208_v0;
          _nw363[(new BigNumber(9)).toNumber()] = (_2219_v11).dtor_cf5;
          _nw363[(new BigNumber(10)).toNumber()] = _2208_v0;
          _2220_v12 = _nw363;
          let _index392 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_2220_v12).length));
          (_2220_v12)[_index392] = _2208_v0;
          let _index393 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_2220_v12).length));
          (_2220_v12)[_index393] = _2208_v0;
        } else {
          (globalState).f8 = _module.__default.safeModuloInt(new BigNumber(142), (_dafny.ZERO).minus(_2206_cf2));
          let _2221_v13;
          let _init60 = ((_2222_cf1) => function (_2223_i0) {
            return (_2223_i0).multipliedBy(_2222_cf1);
          })(_2207_cf1);
          let _nw364 = Array((new BigNumber(23)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw364.length); _i0_60++) {
            _nw364[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2221_v13 = _nw364;
          let _2224_v14;
          _2224_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,false);
          let _2225_v15;
          _2225_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2221_v13,_2224_v14);
          let _2226_v16;
          _2226_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2207_cf1).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2225_v15).length)))),(_this).f9);
          let _2227_v18;
          _2227_v18 = _dafny.Set.fromElements(new BigNumber(-25));
          _2226_v16 = (_2226_v16).Merge((function () {
            let _coll102 = new _dafny.Map();
            for (const _compr_102 of (_2227_v18).Elements) {
              let _2228_v17 = _compr_102;
              if ((_2227_v18).contains(_2228_v17)) {
                _coll102.push([(_2228_v17).minus(_2207_cf1),(_this).f9]);
              }
            }
            return _coll102;
          }()).update(_2207_cf1, (_this).f9));
          let _2229_v19;
          _2229_v19 = _dafny.Seq.UnicodeFromString("fml");
          let _2230_v20;
          let _nw365 = new _module.C6();
          _nw365.__ctor(_2206_cf2, _dafny.Seq.Concat(_2229_v19, _2229_v19));
          _2230_v20 = _nw365;
          let _2231_v21;
          _2231_v21 = _dafny.Set.fromElements((_this).f9, (_this).f9, (_this).f9, (_2213_v5)[_module.__default.safeIndex((_2230_v20).f10, new BigNumber((_2213_v5).length))], (_this).f9);
          (globalState).f3 = _module.__default.safeDivisionInt(new BigNumber((_2231_v21).length), (new BigNumber(-741)).multipliedBy(_module.__default.fm6(_2206_cf2, globalState)));
          let _2232_v22;
          _2232_v22 = _module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe((((_2224_v14).contains((_this).f9)) ? ((_2224_v14).get((_this).f9)) : ((_this).f9)),_2206_cf2));
          let _2233_v23;
          _2233_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2232_v22,true);
          let _2234_v24;
          _2234_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2206_cf2,_2233_v23);
          let _rhs381 = new BigNumber(((_2234_v24).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2230_v20).f11).length),_2233_v23))).length);
          let _rhs382 = (_this).f9;
          let _lhs301 = globalState;
          _lhs301.f3 = _rhs381;
          r0 = _rhs382;
        }
      } else if (_source38.is_DC2) {
        let _2235_v25;
        _2235_v25 = _dafny.Seq.of((_this).f9, !((_this).f9));
        let _2236_v26;
        _2236_v26 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_2235_v25).length));
        let _2237_v27;
        _2237_v27 = new BigNumber(-901);
        let _2238_v29;
        _2238_v29 = _dafny.Seq.UnicodeFromString("ivrfg");
        let _2239_v30;
        _2239_v30 = _dafny.Seq.of(new BigNumber((_2238_v29).length));
        let _2240_v31;
        let _nw366 = Array((new BigNumber(18)).toNumber());
        _nw366[(_dafny.ZERO).toNumber()] = _2236_v26;
        _nw366[(_dafny.ONE).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(2)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(3)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(4)).toNumber()] = _module.__default.fm46(globalState);
        _nw366[(new BigNumber(5)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_2237_v27);
        _nw366[(new BigNumber(7)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(8)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(9)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2237_v27);
        _nw366[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2237_v27,(_this).f9)).update(new BigNumber((function () {
          let _coll103 = new _dafny.Map();
          for (const _compr_103 of _dafny.IntegerRange(new BigNumber(993), new BigNumber(539))) {
            let _2241_v28 = _compr_103;
            if (((new BigNumber(993)).isLessThanOrEqualTo(_2241_v28)) && ((_2241_v28).isLessThan(new BigNumber(539)))) {
              _coll103.push([(_2241_v28).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jkmh")).length)),(_this).f9]);
            }
          }
          return _coll103;
        }()).length), (_this).f9)).length));
        _nw366[(new BigNumber(12)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(13)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(14)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(_2237_v27, new BigNumber((_dafny.MultiSet.FromArray(_2239_v30)).cardinality()), _2237_v27, _2237_v27)).length));
        _nw366[(new BigNumber(16)).toNumber()] = _2236_v26;
        _nw366[(new BigNumber(17)).toNumber()] = _2236_v26;
        _2240_v31 = _nw366;
        let _2242_v32;
        let _nw367 = new _module.C4();
        _nw367.__ctor(_2240_v31);
        _2242_v32 = _nw367;
        let _2243_v33;
        _2243_v33 = _dafny.MultiSet.fromElements((_this).f9);
        let _2244_v34;
        _2244_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Seq.UnicodeFromString("fal"));
        let _2245_v35;
        _2245_v35 = _module.D4.create_DC11((new BigNumber(417)).isLessThan(new BigNumber((((_2243_v33).update(true, _module.__default.abs(new BigNumber(836)))).update((_this).f9, _module.__default.abs(_2237_v27))).cardinality())), (_2237_v27).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((((_2244_v34).contains((_this).f9)) ? ((_2244_v34).get((_this).f9)) : (_2238_v29))).length)))), _2237_v27, _2238_v29);
        let _2246_v36;
        _2246_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2235_v25,(_this).f9);
        let _2247_v37;
        _2247_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _2248_v38;
        _2248_v38 = _module.D12.create_DC27(new BigNumber(16), _2237_v27, _2237_v27, new BigNumber((_2247_v37).length));
        let _2249_v39;
        let _nw368 = Array((new BigNumber(20)).toNumber());
        _nw368[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw368[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(3)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(7)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(12)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(13)).toNumber()] = false;
        _nw368[(new BigNumber(14)).toNumber()] = false;
        _nw368[(new BigNumber(15)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(16)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(17)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(18)).toNumber()] = (_this).f9;
        _nw368[(new BigNumber(19)).toNumber()] = (_this).f9;
        _2249_v39 = _nw368;
        let _2250_v40;
        _2250_v40 = _dafny.MultiSet.fromElements(_2249_v39, _2249_v39, _2249_v39);
        let _2251_v41;
        _2251_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2237_v27,_2237_v27);
        let _2252_v42;
        _2252_v42 = _module.D3.create_DC8(_2251_v41);
        let _2253_v43;
        _2253_v43 = _dafny.MultiSet.fromElements(_2235_v25);
        let _2254_v44;
        let _nw369 = new _module.C6();
        _nw369.__ctor(_2237_v27, (((_2244_v34).contains((_this).f9)) ? ((_2244_v34).get((_this).f9)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), ((_2255_v29, _2256_v27) => function (_2257_i1) {
          return (_2255_v29)[_module.__default.safeIndex(_2256_v27, new BigNumber((_2255_v29).length))];
        })(_2238_v29, _2237_v27)))));
        _2254_v44 = _nw369;
        let _2258_v45;
        _2258_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2254_v44);
        let _2259_v46;
        let _nw370 = Array((new BigNumber(23)).toNumber());
        _nw370[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2246_v36).length));
        _nw370[(_dafny.ONE).toNumber()] = ((_2248_v38).dtor_cf39).minus(_2237_v27);
        _nw370[(new BigNumber(2)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(3)).toNumber()] = new BigNumber(((_2250_v40).update(_2249_v39, _module.__default.abs(_2237_v27))).cardinality());
        _nw370[(new BigNumber(4)).toNumber()] = (((_this).f9) ? (new BigNumber((_2235_v25).length)) : (_module.__default.fm6(_2237_v27, globalState)));
        _nw370[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2251_v41).length), _2237_v27);
        _nw370[(new BigNumber(6)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(7)).toNumber()] = new BigNumber(-926);
        _nw370[(new BigNumber(8)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(9)).toNumber()] = (((_2243_v33).contains((_this).f9)) ? ((_2243_v33).get((_this).f9)) : (new BigNumber(835)));
        _nw370[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(_2237_v27, _2237_v27);
        _nw370[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm26(_2252_v42, _2237_v27, _2253_v43, globalState), _2237_v27);
        _nw370[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((((_2258_v45).update((_this).f9, _2254_v44)).update((_module.D9.create_DC21((_this).f9, _2235_v25, true)).dtor_cf26, _2254_v44)).length)));
        _nw370[(new BigNumber(13)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(14)).toNumber()] = (_2254_v44).f10;
        _nw370[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_2237_v27);
        _nw370[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_2237_v27);
        _nw370[(new BigNumber(17)).toNumber()] = (_2237_v27).minus(new BigNumber(52));
        _nw370[(new BigNumber(18)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(19)).toNumber()] = _2237_v27;
        _nw370[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_2237_v27);
        _nw370[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.of(_2254_v44)).length);
        _nw370[(new BigNumber(22)).toNumber()] = _2237_v27;
        _2259_v46 = _nw370;
        let _2260_v47;
        _2260_v47 = _dafny.Seq.of(_2259_v46, _2259_v46, _2259_v46, _2259_v46);
        let _2261_v48;
        _2261_v48 = _dafny.Seq.of(_2259_v46, _2259_v46, _2259_v46, (_2260_v47)[_module.__default.safeIndex((_2254_v44).f10, new BigNumber((_2260_v47).length))]);
        let _rhs383 = _module.D4.create_DC11((_this).f9, true, new BigNumber(-501), _dafny.Seq.Concat((_2254_v44).f11, (_2254_v44).f11));
        let _rhs384 = (_2260_v47)[_module.__default.safeIndex(new BigNumber((_2235_v25).length), new BigNumber((_2260_v47).length))];
        let _rhs385 = (_this).f9;
        _2245_v35 = _rhs383;
        _2259_v46 = _rhs384;
        r0 = _rhs385;
        let _2262_v49;
        let _nw371 = new _module.C4();
        _nw371.__ctor((_2242_v32).f12);
        _2262_v49 = _nw371;
        r1 = _module.__default.fm20((_this).f9, globalState);
      } else if (_source38.is_DC0) {
        let _2263___mcc_h2 = (_source38).cf0;
        let _2264_cf0 = _2263___mcc_h2;
        let _2265_v50;
        _2265_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_2264_cf0).multipliedBy(new BigNumber(944)));
        let _2266_v51;
        _2266_v51 = _dafny.Seq.of(_2264_cf0);
        _2265_v50 = (_2265_v50).update((_this).f9, (_2264_cf0).plus(new BigNumber((_2266_v51).length)));
        (globalState).f8 = _2264_cf0;
        (globalState).f4 = (_2264_cf0).isLessThanOrEqualTo((_dafny.ZERO).minus(_2264_cf0));
        let _2267_v52;
        let _init61 = function (_2268_i2) {
          return (_this).f9;
        };
        let _nw372 = Array((new BigNumber(10)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw372.length); _i0_61++) {
          _nw372[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2267_v52 = _nw372;
        let _2269_v53;
        _2269_v53 = _dafny.Seq.UnicodeFromString("jwsx");
        let _2270_v54;
        _2270_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2269_v53).length),_2267_v52);
        let _2271_v55;
        _2271_v55 = _dafny.MultiSet.fromElements(_2264_cf0, new BigNumber((_dafny.Seq.of(new BigNumber(-984), new BigNumber(-528))).length), _2264_cf0);
        let _2272_v56;
        let _nw373 = Array((new BigNumber(14)).toNumber());
        _nw373[(_dafny.ZERO).toNumber()] = _2267_v52;
        _nw373[(_dafny.ONE).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(2)).toNumber()] = (((_2270_v54).contains(new BigNumber((_2271_v55).cardinality()))) ? ((_2270_v54).get(new BigNumber((_2271_v55).cardinality()))) : (_2267_v52));
        _nw373[(new BigNumber(3)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(4)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(5)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(6)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(7)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(8)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(9)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(10)).toNumber()] = ((true) ? (_2267_v52) : (_2267_v52));
        _nw373[(new BigNumber(11)).toNumber()] = (_module.D2.create_DC5(_2267_v52)).dtor_cf5;
        _nw373[(new BigNumber(12)).toNumber()] = _2267_v52;
        _nw373[(new BigNumber(13)).toNumber()] = _2267_v52;
        _2272_v56 = _nw373;
        let _index394 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_2272_v56).length));
        (_2272_v56)[_index394] = _2267_v52;
        let _index395 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_2272_v56).length));
        (_2272_v56)[_index395] = _2267_v52;
      } else {
        let _2273___mcc_h3 = (_source38).cf3;
        let _2274_cf3 = _2273___mcc_h3;
        if ((_this).f9) {
          let _2275_v57;
          let _nw374 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _2275_v57 = _nw374;
          _2275_v57 = _2275_v57;
          let _2276_v58;
          _2276_v58 = _dafny.Seq.UnicodeFromString("ljehk");
          let _2277_v59;
          _2277_v59 = _dafny.Seq.of(_2276_v58, _dafny.Seq.UnicodeFromString("pbuyasjcd"));
          let _2278_v60;
          _2278_v60 = _dafny.Seq.of(!_dafny.areEqual(_dafny.Seq.of(_2276_v58, _2276_v58, _2276_v58, _dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), function (_2279_i3) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })), _2277_v59));
          _2278_v60 = _2278_v60;
          let _2280_v61;
          let _nw375 = new _module.C5();
          _nw375.__ctor();
          _2280_v61 = _nw375;
          _2276_v58 = _dafny.Seq.UnicodeFromString("mrlflpdvu");
          let _2281_v62;
          _2281_v62 = new _dafny.CodePoint('m'.codePointAt(0));
          let _2282_v63;
          _2282_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2281_v62);
          let _2283_v64;
          _2283_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2282_v63,_2276_v58);
          _2283_v64 = (_2283_v64).update((_2282_v63).Merge(_2282_v63), _2276_v58);
        } else {
          let _2284_v65;
          _2284_v65 = _dafny.Seq.UnicodeFromString("kypi");
          let _2285_v66;
          _2285_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2284_v65).length),false);
          let _2286_v67;
          _2286_v67 = new BigNumber(831);
          _2285_v66 = (_2285_v66).update(_module.__default.fm6(_2286_v67, globalState), false);
          (globalState).f3 = (_dafny.ZERO).minus(_2286_v67);
          (globalState).f3 = _2286_v67;
          let _2287_v68;
          let _nw376 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2287_v68 = _nw376;
          let _2288_v69;
          _2288_v69 = _dafny.Seq.of((_this).f9);
          let _index396 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_2287_v68).length));
          (_2287_v68)[_index396] = _dafny.MultiSet.FromArray(_2288_v69);
          let _index397 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_2287_v68).length));
          (_2287_v68)[_index397] = (_dafny.MultiSet.fromElements((_this).f9, (_this).f9, !(!((_this).f9)))).update((_this).f9, _module.__default.abs(_2286_v67));
          let _2289_v70;
          _2289_v70 = _dafny.Set.fromElements(!((_this).f9));
          r2 = _dafny.Set.fromElements(_2286_v67, _module.__default.safeModuloInt(_2286_v67, _2286_v67), new BigNumber(((_dafny.Set.fromElements(true, (_this).f9, (_this).f9, (_this).f9, (_this).f9)).Intersect(_2289_v70)).length), (_dafny.ZERO).minus((_2286_v67).multipliedBy(new BigNumber(293))), _2286_v67);
        }
        let _2290_v71;
        _2290_v71 = _dafny.Seq.of(new BigNumber(390));
        let _2291_v72;
        _2291_v72 = new BigNumber(603);
        let _2292_v73;
        _2292_v73 = _dafny.Set.fromElements(_dafny.Seq.Concat(_2290_v71, _2290_v71), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_2293_i4) {
          return new BigNumber(-895);
        }), _module.__default.safeIndex(_2291_v72, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_2294_i4) {
          return new BigNumber(-895);
        })).length)), new BigNumber(206)), _2290_v71);
        _2292_v73 = _2292_v73;
        let _2295_v74;
        _2295_v74 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2296_v75;
        _2296_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2295_v74,_2295_v74);
        let _2297_v76;
        let _nw377 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _2297_v76 = _nw377;
        let _index398 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_2297_v76).length));
        (_2297_v76)[_index398] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2297_v76,(_this).f9)).length));
        let _2298_v77;
        let _init62 = function (_2299_i5) {
          return (_this).f9;
        };
        let _nw378 = Array((new BigNumber(3)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw378.length); _i0_62++) {
          _nw378[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _2298_v77 = _nw378;
        let _2300_v78;
        _2300_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2298_v77,_2296_v75);
        let _2301_v79;
        _2301_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2291_v72,_2291_v72);
        let _2302_v80;
        _2302_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2301_v79).length),true);
        let _2303_v81;
        _2303_v81 = _dafny.MultiSet.fromElements((_this).f9);
        let _index399 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_2297_v76).length));
        let _rhs386 = (((_2300_v78).contains(_2298_v77)) ? ((_2300_v78).get(_2298_v77)) : (_dafny.Map.Empty.slice().updateUnsafe(_2295_v74,_module.__default.fm18((((_2302_v80).contains(_2291_v72)) ? ((_2302_v80).get(_2291_v72)) : ((_this).f9)), (_this).f9, globalState))));
        let _rhs387 = (((_this).f9) ? ((_2291_v72).multipliedBy(new BigNumber((_2303_v81).cardinality()))) : ((_2290_v71)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_2290_v71).length))]));
        let _lhs302 = _2297_v76;
        let _lhs303 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_2297_v76).length));
        _2296_v75 = _rhs386;
        _lhs302[_lhs303] = _rhs387;
        let _2304_v82;
        _2304_v82 = _dafny.Seq.UnicodeFromString("ktuiokx");
        _2304_v82 = _module.__default.fm25(globalState);
      }
      let _2305_v83;
      let _nw379 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _2305_v83 = _nw379;
      let _index400 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length));
      (_2305_v83)[_index400] = _module.__default.fm20((_this).f9, globalState);
      let _2306_v84;
      _2306_v84 = new BigNumber(382);
      let _2307_v85;
      _2307_v85 = _dafny.Seq.of((_this).f9);
      let _2308_v86;
      _2308_v86 = _dafny.Seq.of(_2306_v84, _2306_v84, _2306_v84, new BigNumber((_2307_v85).length));
      let _2309_v87;
      _2309_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
      let _index401 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length));
      (_2305_v83)[_index401] = _module.__default.fm6(_module.__default.safeModuloInt(_2306_v84, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_2308_v86, _module.__default.safeIndex(_module.__default.fm6(new BigNumber((_2309_v87).length), globalState), new BigNumber((_2308_v86).length)), _2306_v84))).cardinality())), globalState);
      let _2310_v88;
      _2310_v88 = _dafny.Seq.UnicodeFromString("ojpl");
      let _2311_v89;
      _2311_v89 = _module.D4.create_DC11((_this).f9, (_this).f9, new BigNumber((_dafny.Seq.of(_2306_v84, (_2305_v83)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length))])).length), _2310_v88);
      let _hi14 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(722)), function (_2312_i7) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length);
      for (let _2313_i6 = (_2311_v89).dtor_cf16; _2313_i6.isLessThan(_hi14); _2313_i6 = _2313_i6.plus(_dafny.ONE)) {
        let _2314_v90;
        _2314_v90 = _module.D4.create_DC10(_dafny.Seq.of((_this).f9));
        let _2315_v91;
        _2315_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49((_this).f9, globalState),_dafny.Seq.Concat(_dafny.Seq.of((_this).f9), (_2314_v90).dtor_cf13));
        r1 = new BigNumber((_2315_v91).length);
        r0 = (_this).f9;
        (globalState).f8 = _2313_i6;
        _2310_v88 = _2310_v88;
      }
      let _index402 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length));
      let _rhs388 = ((_2305_v83)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length))]).plus((_2305_v83)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length))]);
      let _rhs389 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2307_v85, _dafny.Seq.update(_2307_v85, _module.__default.safeIndex(new BigNumber(421), new BigNumber((_2307_v85).length)), (_this).f9)), _dafny.Seq.of((_this).f9))).length);
      let _lhs304 = globalState;
      let _lhs305 = _2305_v83;
      let _lhs306 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length));
      _lhs304.f8 = _rhs388;
      _lhs305[_lhs306] = _rhs389;
      let _2316_v92;
      let _nw380 = new _module.C7();
      _nw380.__ctor();
      _2316_v92 = _nw380;
      let _2317_v93;
      _2317_v93 = _dafny.Seq.of(_2316_v92);
      let _2318_v94;
      _2318_v94 = _2317_v93;
      _2318_v94 = _2317_v93;
      (globalState).f3 = (_2305_v83)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length))];
      let _2319_v95;
      _2319_v95 = _module.D2.create_DC6((_2305_v83)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_2305_v83).length))], (_this).f9, (_this).f9);
      let _pat_let_tv71 = _2305_v83;
      let _pat_let_tv72 = _2305_v83;
      r0 = function (_source39) {
        if (_source39.is_DC6) {
          let _2320___mcc_h4 = (_source39).cf6;
          let _2321___mcc_h5 = (_source39).cf7;
          let _2322___mcc_h6 = (_source39).cf8;
          let _2323_cf8 = _2322___mcc_h6;
          let _2324_cf7 = _2321___mcc_h5;
          let _2325_cf6 = _2320___mcc_h4;
          if ((_this).f9) {
            return _2323_cf8;
          } else {
            return !(_2324_cf7);
          }
        } else if (_source39.is_DC7) {
          let _2326___mcc_h7 = (_source39).cf9;
          let _2327___mcc_h8 = (_source39).cf10;
          let _2328___mcc_h9 = (_source39).cf11;
          let _2329_cf11 = _2328___mcc_h9;
          let _2330_cf10 = _2327___mcc_h8;
          let _2331_cf9 = _2326___mcc_h7;
          return (_this).f9;
        } else {
          let _2332___mcc_h10 = (_source39).cf5;
          let _2333_cf5 = _2332___mcc_h10;
          return (_dafny.MultiSet.fromElements(new BigNumber(529))).IsDisjointFrom(_dafny.MultiSet.fromElements((_pat_let_tv72)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_pat_let_tv71).length))]));
        }
      }(_2319_v95);
      r1 = _2306_v84;
      let _2334_v96;
      _2334_v96 = _dafny.Set.fromElements(_2306_v84);
      r2 = _2334_v96;
      return [r0, r1, r2];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _2335_v0;
      _2335_v0 = _dafny.Seq.of((_this).f9);
      let _2336_v1;
      _2336_v1 = _dafny.Seq.of(_2335_v0, _dafny.Seq.of(false), _2335_v0, _2335_v0, _2335_v0);
      let _2337_v2;
      _2337_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2335_v0,(_dafny.ZERO).minus(new BigNumber(((_2336_v1)[_module.__default.safeIndex(p1, new BigNumber((_2336_v1).length))]).length)));
      let _2338_v3;
      _2338_v3 = _dafny.Seq.UnicodeFromString("julgoho");
      let _2339_v4;
      _2339_v4 = new _dafny.CodePoint('o'.codePointAt(0));
      let _2340_v5;
      _2340_v5 = _dafny.MultiSet.fromElements(((((_2337_v2).contains(_2335_v0)) ? ((_2337_v2).get(_2335_v0)) : (p0))).plus(p0), new BigNumber(-539), (((_this).f9) ? (new BigNumber((_2338_v3).length)) : ((_dafny.ZERO).minus(new BigNumber((_module.__default.fm17(_2339_v4, globalState)).length)))), (_dafny.ZERO).minus(p0), p0);
      (globalState).f3 = (((_2340_v5).contains(p0)) ? ((_2340_v5).get(p0)) : (p0));
      let _2341_v6;
      _2341_v6 = _dafny.Set.fromElements(p0);
      if (!((_this).f9) || ((_2341_v6).IsDisjointFrom(_2341_v6))) {
        let _2342_v7;
        _2342_v7 = _dafny.Seq.of(new BigNumber(989), p0);
        let _2343_v8;
        _2343_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(841));
        let _2344_v9;
        _2344_v9 = _module.D3.create_DC8(_2343_v8);
        let _2345_v10;
        _2345_v10 = _dafny.MultiSet.fromElements(_2335_v0);
        (globalState).f0 = _dafny.Seq.update(_dafny.Seq.Concat(_2342_v7, _2342_v7), _module.__default.safeIndex((p1).multipliedBy(p1), new BigNumber((_dafny.Seq.Concat(_2342_v7, _2342_v7)).length)), _module.__default.fm26(_2344_v9, new BigNumber(687), _2345_v10, globalState));
        let _2346_v11;
        let _nw381 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _2346_v11 = _nw381;
        let _2347_v12;
        _2347_v12 = _dafny.Set.fromElements(_2346_v11, _2346_v11);
        let _2348_v13;
        let _nw382 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _2348_v13 = _nw382;
        let _2349_v14;
        let _nw383 = new _module.C4();
        _nw383.__ctor(_2348_v13);
        _2349_v14 = _nw383;
        let _2350_v15;
        _2350_v15 = _dafny.Seq.of(_2346_v11);
        let _rhs390 = (_this).f9;
        let _rhs391 = (_this).f9;
        let _rhs392 = _dafny.Set.fromElements(_2346_v11, (((_this).f9) ? (_2346_v11) : (_2346_v11)), (_2350_v15)[_module.__default.safeIndex(p1, new BigNumber((_2350_v15).length))]);
        let _rhs393 = (_dafny.ZERO).minus(new BigNumber(-777));
        let _rhs394 = _2349_v14;
        let _lhs307 = globalState;
        let _lhs308 = globalState;
        r3 = _rhs390;
        _lhs307.f4 = _rhs391;
        _2347_v12 = _rhs392;
        _lhs308.f3 = _rhs393;
        _2349_v14 = _rhs394;
        if (((p0).plus(p0)).isLessThan(p0)) {
          (globalState).f3 = _module.__default.safeModuloInt(p1, (_module.__default.fm6(_module.__default.fm20((_this).f9, globalState), globalState)).minus(p0));
          let _2351_v16;
          _2351_v16 = _dafny.Set.fromElements(_2339_v4, _module.__default.fm18((_this).f9, (_this).f9, globalState), _module.__default.fm18((_this).f9, false, globalState), _2339_v4);
          _2351_v16 = (_2351_v16).Intersect((_dafny.Set.fromElements(_2339_v4)).Difference(_dafny.Set.fromElements((_module.D16.create_DC35(p0, _2339_v4, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-613)), ((_2352_v4) => function (_2353_i0) {
  return _2352_v4;
})(_2339_v4)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-613)), ((_2354_v4) => function (_2355_i0) {
  return _2354_v4;
})(_2339_v4))).length)), _2339_v4)).length), p0, p1)).dtor_cf51)));
          (globalState).f4 = (((_this).f9) ? ((_this).f9) : ((_this).f9));
          let _2356_v17;
          let _init63 = function (_2357_i1) {
            return (_this).f9;
          };
          let _nw384 = Array((new BigNumber(22)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw384.length); _i0_63++) {
            _nw384[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2356_v17 = _nw384;
          let _2358_v18;
          _2358_v18 = _dafny.Seq.of(_2356_v17, _2356_v17, _2356_v17);
          _2358_v18 = (((_this).f9) ? (_2358_v18) : (_2358_v18));
          _2335_v0 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_2335_v0, _module.__default.safeIndex(p0, new BigNumber((_2335_v0).length)), (_this).f9), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_2335_v0, _module.__default.safeIndex(p0, new BigNumber((_2335_v0).length)), (_this).f9)).length)), (_this).f9), _module.__default.safeIndex(_module.__default.safeModuloInt(_module.__default.fm20((_this).f9, globalState), p0), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_2335_v0, _module.__default.safeIndex(p0, new BigNumber((_2335_v0).length)), (_this).f9), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_2335_v0, _module.__default.safeIndex(p0, new BigNumber((_2335_v0).length)), (_this).f9)).length)), (_this).f9)).length)), (_this).f9);
        } else {
          let _2359_v19;
          let _nw385 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2359_v19 = _nw385;
          let _2360_v20;
          _2360_v20 = _module.D21.create_DC43(_2342_v7, (_this).f9);
          let _2361_v21;
          let _nw386 = Array((new BigNumber(21)).toNumber());
          _nw386[(_dafny.ZERO).toNumber()] = (_this).f9;
          _nw386[(_dafny.ONE).toNumber()] = (_2349_v14).fm3((_2360_v20).dtor_cf63, _2340_v5, (_this).f9, globalState);
          _nw386[(new BigNumber(2)).toNumber()] = false;
          _nw386[(new BigNumber(3)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(4)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(6)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(7)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(8)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(9)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(10)).toNumber()] = false;
          _nw386[(new BigNumber(11)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(12)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(13)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(14)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(15)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(16)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(17)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(18)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(19)).toNumber()] = (_this).f9;
          _nw386[(new BigNumber(20)).toNumber()] = (_this).f9;
          _2361_v21 = _nw386;
          let _2362_v22;
          _2362_v22 = _dafny.MultiSet.fromElements(_2361_v21);
          let _index403 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_2359_v19).length));
          (_2359_v19)[_index403] = _2362_v22;
          let _index404 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_2359_v19).length));
          (_2359_v19)[_index404] = _dafny.MultiSet.fromElements(_2361_v21, _2361_v21, _2361_v21);
          let _2363_v23;
          _2363_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2335_v0,_2342_v7);
          let _2364_v24;
          _2364_v24 = _dafny.Seq.of(_2363_v23);
          _2363_v23 = (_2364_v24)[_module.__default.safeIndex(new BigNumber(-356), new BigNumber((_2364_v24).length))];
          let _index405 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_2361_v21).length));
          (_2361_v21)[_index405] = _module.__default.fm49(_module.__default.fm49((_this).f9, globalState), globalState);
          let _index406 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_2361_v21).length));
          (_2361_v21)[_index406] = ((false) ? ((_this).f9) : ((_this).f9));
          (globalState).f3 = p1;
          (globalState).f4 = (_this).f9;
        }
        let _2365_v25;
        _2365_v25 = _dafny.Set.fromElements((_this).f9, (_this).f9);
        let _2366_v26;
        _2366_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2365_v25).IsDisjointFrom(_2365_v25));
        _2366_v26 = (_2366_v26).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2342_v7,_2346_v11)).length), (_this).f9);
        let _index407 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_2346_v11).length));
        (_2346_v11)[_index407] = p0;
        let _2367_v27;
        _2367_v27 = _dafny.Seq.of(_2338_v3);
        let _2368_v28;
        _2368_v28 = _module.D10.create_DC24(p1, !((_this).f9), (_2367_v27)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_2369_v3) => function (_2370_i2) {
  return _2369_v3;
})(_2338_v3))).length), new BigNumber((_2367_v27).length))]);
        let _index408 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_2346_v11).length));
        (_2346_v11)[_index408] = (_2368_v28).dtor_cf31;
      } else {
        _2338_v3 = _2338_v3;
        let _2371_v29;
        let _nw387 = Array((new BigNumber(5)).toNumber()).fill(false);
        _2371_v29 = _nw387;
        let _index409 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
        (_2371_v29)[_index409] = (_this).f9;
        let _index410 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
        (_2371_v29)[_index410] = !(_2341_v6).equals(_dafny.Set.fromElements(p0, p1, p1));
        if ((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))]) {
          let _2372_v30;
          _2372_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,!((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))]));
          (globalState).f8 = (p1).multipliedBy((_dafny.ZERO).minus(new BigNumber((_2372_v30).length)));
          (globalState).f4 = (_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))];
          let _2373_v31;
          _2373_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Concat(_2338_v3, _2338_v3));
          _2373_v31 = (_2373_v31).update(new BigNumber(((((_this).f9) ? (_2340_v5) : (_2340_v5))).cardinality()), _2338_v3);
          let _2374_v32;
          let _nw388 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _2374_v32 = _nw388;
          let _index411 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_2374_v32).length));
          (_2374_v32)[_index411] = (p0).multipliedBy(p1);
          let _index412 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_2374_v32).length));
          (_2374_v32)[_index412] = p0;
          let _2375_v33;
          let _nw389 = new _module.C6();
          _nw389.__ctor(p1, _2338_v3);
          _2375_v33 = _nw389;
          let _2376_v34;
          _2376_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2375_v33,!((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))]));
          _2376_v34 = (_2376_v34).update(_2375_v33, (_this).f9);
        } else {
          let _2377_v35;
          _2377_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _2378_v36;
          _2378_v36 = _dafny.Seq.of((_dafny.ZERO).minus(p1), p0);
          let _2379_v37;
          _2379_v37 = _module.D21.create_DC43(_2378_v36, (_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))]);
          let _2380_v38;
          _2380_v38 = _2378_v36;
          let _index413 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
          let _rhs395 = (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,p1));
          let _rhs396 = (_this).f9;
          let _rhs397 = _2371_v29;
          let _rhs398 = _module.D21.create_DC43(_dafny.Seq.update(_2378_v36, _module.__default.safeIndex(p1, new BigNumber((_2378_v36).length)), p0), _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_2381_p0) => function (_2382_i3) {
  return _2381_p0;
})(p0)), _2380_v38, _2380_v38), _dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), ((_2383_v38) => function (_2384_i4) {
  return _2383_v38;
})(_2380_v38))));
          let _rhs399 = (_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))];
          let _lhs309 = globalState;
          let _lhs310 = _2371_v29;
          let _lhs311 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
          _2377_v35 = _rhs395;
          _lhs309.f4 = _rhs396;
          _2371_v29 = _rhs397;
          _2379_v37 = _rhs398;
          _lhs310[_lhs311] = _rhs399;
          let _2385_v39;
          let _nw390 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
          _2385_v39 = _nw390;
          let _2386_v40;
          _2386_v40 = _dafny.Map.Empty.slice().updateUnsafe((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))],p1);
          let _index414 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_2385_v39).length));
          (_2385_v39)[_index414] = _2386_v40;
          let _2387_v41;
          let _nw391 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2387_v41 = _nw391;
          let _2388_v42;
          _2388_v42 = _dafny.MultiSet.fromElements((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))], (_this).f9);
          let _2389_v43;
          _2389_v43 = _dafny.Map.Empty.slice().updateUnsafe((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))],(_this).f9);
          let _index415 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_2385_v39).length));
          let _rhs400 = ((_2388_v42).Intersect(_2388_v42)).IsProperSubsetOf(_2388_v42);
          let _rhs401 = _dafny.Seq.update(_dafny.Seq.Concat(_2378_v36, _dafny.Seq.update(_2378_v36, _module.__default.safeIndex(p1, new BigNumber((_2378_v36).length)), p1)), _module.__default.safeIndex(new BigNumber((_2389_v43).length), new BigNumber((_dafny.Seq.Concat(_2378_v36, _dafny.Seq.update(_2378_v36, _module.__default.safeIndex(p1, new BigNumber((_2378_v36).length)), p1))).length)), p1);
          let _rhs402 = (_2386_v40).Merge((_2386_v40).Merge(_2386_v40));
          let _rhs403 = _2387_v41;
          let _lhs312 = globalState;
          let _lhs313 = globalState;
          let _lhs314 = _2385_v39;
          let _lhs315 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_2385_v39).length));
          _lhs312.f4 = _rhs400;
          _lhs313.f0 = _rhs401;
          _lhs314[_lhs315] = _rhs402;
          _2387_v41 = _rhs403;
          let _2390_v45;
          let _init64 = ((_2391_v6, _2392_p0) => function (_2393_i5) {
            return function () {
              let _coll104 = new _dafny.Map();
              for (const _compr_104 of (_2391_v6).Elements) {
                let _2394_v44 = _compr_104;
                if ((_2391_v6).contains(_2394_v44)) {
                  _coll104.push([(_2394_v44).minus(_2392_p0),_2392_p0]);
                }
              }
              return _coll104;
            }();
          })(_2341_v6, p0);
          let _nw392 = Array((new BigNumber(23)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw392.length); _i0_64++) {
            _nw392[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2390_v45 = _nw392;
          let _index416 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_2390_v45).length));
          (_2390_v45)[_index416] = _2377_v35;
          let _2395_v46;
          _2395_v46 = _module.D9.create_DC21((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))], _dafny.Seq.of((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))]), (_this).f9);
          let _2396_v47;
          _2396_v47 = _dafny.Set.fromElements(_module.__default.fm30(p1, _2395_v46, (_this).f9, p0, globalState));
          let _2397_v50;
          _2397_v50 = _module.D3.create_DC8(function () {
  let _coll105 = new _dafny.Map();
  for (const _compr_105 of (function () {
    let _coll106 = new _dafny.Set();
    for (const _compr_106 of _dafny.IntegerRange(new BigNumber(924), new BigNumber(161))) {
      let _2398_v49 = _compr_106;
      if (((new BigNumber(924)).isLessThanOrEqualTo(_2398_v49)) && ((_2398_v49).isLessThan(new BigNumber(161)))) {
        _coll106.add(_module.__default.safeModuloInt(_2398_v49, new BigNumber((_2378_v36).length)));
      }
    }
    return _coll106;
  }()).Elements) {
    let _2399_v48 = _compr_105;
    if ((function () {
      let _coll107 = new _dafny.Set();
      for (const _compr_107 of _dafny.IntegerRange(new BigNumber(924), new BigNumber(161))) {
        let _2400_v49 = _compr_107;
        if (((new BigNumber(924)).isLessThanOrEqualTo(_2400_v49)) && ((_2400_v49).isLessThan(new BigNumber(161)))) {
          _coll107.add(_module.__default.safeModuloInt(_2400_v49, new BigNumber((_2378_v36).length)));
        }
      }
      return _coll107;
    }()).contains(_2399_v48)) {
      _coll105.push([_module.__default.safeModuloInt(_2399_v48, p1),p1]);
    }
  }
  return _coll105;
}());
          let _2401_v51;
          _2401_v51 = _dafny.MultiSet.fromElements(_dafny.Seq.of(!((_this).f9)));
          let _2402_v52;
          _2402_v52 = _module.D24.create_DC48(_2396_v47);
          let _index417 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_2390_v45).length));
          let _index418 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
          let _rhs404 = _2377_v35;
          let _rhs405 = _module.__default.fm26(_2397_v50, _module.__default.safeDivisionInt(p1, p1), (_2401_v51).Difference(_2401_v51), globalState);
          let _rhs406 = _module.__default.fm49((_this).f9, globalState);
          let _rhs407 = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p1));
          let _rhs408 = ((_2396_v47).Difference((_2402_v52).dtor_cf71)).Union((_2396_v47).Union(_2396_v47));
          let _lhs316 = _2390_v45;
          let _lhs317 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_2390_v45).length));
          let _lhs318 = globalState;
          let _lhs319 = _2371_v29;
          let _lhs320 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length));
          let _lhs321 = globalState;
          _lhs316[_lhs317] = _rhs404;
          _lhs318.f3 = _rhs405;
          _lhs319[_lhs320] = _rhs406;
          _lhs321.f3 = _rhs407;
          _2396_v47 = _rhs408;
          (globalState).f4 = (_2335_v0)[_module.__default.safeIndex(p1, new BigNumber((_2335_v0).length))];
          let _rhs409 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("dhkfosxv")).length), ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f9)).length))).plus(p1));
          let _rhs410 = new BigNumber(-629);
          let _lhs322 = globalState;
          _lhs322.f3 = _rhs409;
          r2 = _rhs410;
        }
        let _2403_v53;
        _2403_v53 = _dafny.Map.Empty.slice().updateUnsafe((_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))],_2341_v6);
        let _2404_v54;
        _2404_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _rhs411 = (_2371_v29)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_2371_v29).length))];
        let _rhs412 = _module.__default.safeModuloInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2403_v53).length))).Merge(_2404_v54)).length), _module.__default.fm20(_module.__default.fm49(true, globalState), globalState));
        let _lhs323 = globalState;
        r3 = _rhs411;
        _lhs323.f3 = _rhs412;
        let _2405_v55;
        _2405_v55 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1(new BigNumber(630), p1)).dtor_cf2,_2340_v5);
        let _2406_v56;
        _2406_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2405_v55,(_dafny.ZERO).minus(p1));
        _2406_v56 = (_2406_v56).update((_2405_v55).Merge(_2405_v55), new BigNumber(652));
      }
      let _2407_v57;
      let _nw393 = Array((new BigNumber(23)).toNumber()).fill(false);
      _2407_v57 = _nw393;
      let _index419 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length));
      (_2407_v57)[_index419] = (((_this).f9) ? ((_this).f9) : ((_this).f9));
      let _index420 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length));
      (_2407_v57)[_index420] = false;
      let _2408_v58;
      let _nw394 = Array((new BigNumber(4)).toNumber());
      _nw394[(_dafny.ZERO).toNumber()] = _2339_v4;
      _nw394[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
      _nw394[(new BigNumber(2)).toNumber()] = _2339_v4;
      _nw394[(new BigNumber(3)).toNumber()] = _2339_v4;
      _2408_v58 = _nw394;
      let _index421 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_2408_v58).length));
      (_2408_v58)[_index421] = _2339_v4;
      let _index422 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_2408_v58).length));
      (_2408_v58)[_index422] = (((_2407_v57)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length))]) ? (_2339_v4) : (((!((_2407_v57)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length))])) ? (_2339_v4) : (new _dafny.CodePoint('r'.codePointAt(0))))));
      let _2409_v59;
      _2409_v59 = _dafny.Seq.of(_2407_v57);
      let _2410_v60;
      let _nw395 = Array((new BigNumber(22)).toNumber());
      _nw395[(_dafny.ZERO).toNumber()] = _2407_v57;
      _nw395[(_dafny.ONE).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(2)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(3)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(4)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(5)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(6)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(7)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(8)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(9)).toNumber()] = (((_2407_v57)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length))]) ? (_2407_v57) : (_2407_v57));
      _nw395[(new BigNumber(10)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(11)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(12)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(13)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(14)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(15)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(16)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(17)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(18)).toNumber()] = (_2409_v59)[_module.__default.safeIndex(new BigNumber((_2338_v3).length), new BigNumber((_2409_v59).length))];
      _nw395[(new BigNumber(19)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(20)).toNumber()] = _2407_v57;
      _nw395[(new BigNumber(21)).toNumber()] = _2407_v57;
      _2410_v60 = _nw395;
      _2410_v60 = _2410_v60;
      let _2411_v61;
      let _nw396 = Array((new BigNumber(10)).toNumber()).fill(_module.D19.Default());
      _2411_v61 = _nw396;
      let _2412_v62;
      _2412_v62 = _dafny.Set.fromElements(_2411_v61);
      let _2413_v63;
      _2413_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_2408_v58)[_module.__default.safeIndex(new BigNumber(802), new BigNumber((_2408_v58).length))]);
      let _2414_v64;
      let _nw397 = Array((new BigNumber(20)).toNumber());
      _nw397[(_dafny.ZERO).toNumber()] = p0;
      _nw397[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f9, (_this).f9)).length));
      _nw397[(new BigNumber(2)).toNumber()] = p0;
      _nw397[(new BigNumber(3)).toNumber()] = p0;
      _nw397[(new BigNumber(4)).toNumber()] = new BigNumber(889);
      _nw397[(new BigNumber(5)).toNumber()] = p1;
      _nw397[(new BigNumber(6)).toNumber()] = p0;
      _nw397[(new BigNumber(7)).toNumber()] = p1;
      _nw397[(new BigNumber(8)).toNumber()] = _module.__default.fm20((_2407_v57)[_module.__default.safeIndex(new BigNumber(570), new BigNumber((_2407_v57).length))], globalState);
      _nw397[(new BigNumber(9)).toNumber()] = p0;
      _nw397[(new BigNumber(10)).toNumber()] = p1;
      _nw397[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw397[(new BigNumber(12)).toNumber()] = p1;
      _nw397[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw397[(new BigNumber(14)).toNumber()] = p1;
      _nw397[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw397[(new BigNumber(16)).toNumber()] = p1;
      _nw397[(new BigNumber(17)).toNumber()] = new BigNumber((_2413_v63).length);
      _nw397[(new BigNumber(18)).toNumber()] = new BigNumber(240);
      _nw397[(new BigNumber(19)).toNumber()] = p0;
      _2414_v64 = _nw397;
      let _2415_v65;
      _2415_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2412_v62,_2414_v64);
      r0 = (((_2415_v65).contains((_2412_v62).Difference(_2412_v62))) ? ((_2415_v65).get((_2412_v62).Difference(_2412_v62))) : (_2414_v64));
      r0 = _2414_v64;
      r1 = p0;
      r2 = p1;
      r3 = false;
      return [r0, r1, r2, r3];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(645),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(663))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(81),false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jsdxerq"))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-368),_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality())))).length)), new BigNumber((function () {
        let _coll108 = new _dafny.Map();
        for (const _compr_108 of (_dafny.MultiSet.fromElements(new BigNumber(-85))).Elements) {
          let _2416_v0 = _compr_108;
          if ((_dafny.MultiSet.fromElements(new BigNumber(-85))).contains(_2416_v0)) {
            _coll108.push([(_2416_v0).minus(new BigNumber(898)),new _dafny.CodePoint('t'.codePointAt(0))]);
          }
        }
        return _coll108;
      }()).length))).cardinality()))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), function (_2417_i0) {
        return new BigNumber(23);
      })).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-29),new BigNumber((_dafny.Set.fromElements(new BigNumber(491))).length)));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    fm4(p0, globalState) {
      let _this = this;
      return (function () {
        let _coll109 = new _dafny.Map();
        for (const _compr_109 of (_dafny.Seq.of(_module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(true, true)).length)))).Elements) {
          let _2418_v0 = _compr_109;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(true, true)).length))), _2418_v0)) {
            _coll109.push([_2418_v0,true]);
          }
        }
        return _coll109;
      }()).Merge(function () {
        let _coll110 = new _dafny.Map();
        for (const _compr_110 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber(-672)),new _dafny.CodePoint('r'.codePointAt(0)))).Keys.Elements) {
          let _2419_v1 = _compr_110;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(new BigNumber(-672)),new _dafny.CodePoint('r'.codePointAt(0)))).contains(_2419_v1)) {
            _coll110.push([_2419_v1,false]);
          }
        }
        return _coll110;
      }());
    };
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _2420_v0;
      let _nw398 = new _module.C0();
      _nw398.__ctor();
      _2420_v0 = _nw398;
      let _2421_v1;
      _2421_v1 = true;
      let _2422_v2;
      _2422_v2 = _dafny.MultiSet.fromElements(_2421_v1);
      _2422_v2 = _2422_v2;
      let _2423_v3;
      _2423_v3 = new BigNumber(428);
      let _hi15 = (_2423_v3).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(_2423_v3)));
      for (let _2424_i0 = (_2423_v3).plus(_2423_v3); _2424_i0.isLessThan(_hi15); _2424_i0 = _2424_i0.plus(_dafny.ONE)) {
        _2423_v3 = new BigNumber(-567);
        let _2425_v4;
        _2425_v4 = _dafny.Seq.of(_2421_v1, _2421_v1);
        let _2426_v6;
        _2426_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_2424_i0)),_2421_v1);
        let _2427_v7;
        _2427_v7 = _dafny.Set.fromElements(new BigNumber((function () {
          let _coll111 = new _dafny.Map();
          for (const _compr_111 of (_2426_v6).Keys.Elements) {
            let _2428_v5 = _compr_111;
            if ((_2426_v6).contains(_2428_v5)) {
              _coll111.push([(_2428_v5).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(319)), function (_2429_i1) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              })).length)),_2421_v1]);
            }
          }
          return _coll111;
        }()).length));
        let _2430_v8;
        _2430_v8 = _module.D13.create_DC30(_2421_v1, (_2425_v4)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_2425_v4).length))], _2427_v7, _2420_v0);
        let _2431_v9;
        let _nw399 = Array((new BigNumber(3)).toNumber());
        _nw399[(_dafny.ZERO).toNumber()] = _2430_v8;
        _nw399[(_dafny.ONE).toNumber()] = _2430_v8;
        _nw399[(new BigNumber(2)).toNumber()] = _2430_v8;
        _2431_v9 = _nw399;
        let _index423 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_2431_v9).length));
        (_2431_v9)[_index423] = _2430_v8;
        let _2432_v10;
        let _nw400 = Array((new BigNumber(13)).toNumber()).fill(false);
        _2432_v10 = _nw400;
        let _index424 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_2431_v9).length));
        let _rhs413 = _2430_v8;
        let _rhs414 = _2423_v3;
        let _rhs415 = ((_dafny.areEqual(_2425_v4, _2425_v4)) ? (_2432_v10) : (_2432_v10));
        let _rhs416 = _2423_v3;
        let _lhs324 = _2431_v9;
        let _lhs325 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_2431_v9).length));
        let _lhs326 = globalState;
        let _lhs327 = globalState;
        _lhs324[_lhs325] = _rhs413;
        _lhs326.f8 = _rhs414;
        _2432_v10 = _rhs415;
        _lhs327.f8 = _rhs416;
        let _2433_v11;
        let _nw401 = Array((new BigNumber(17)).toNumber());
        _2433_v11 = _nw401;
        let _2434_v12;
        let _nw402 = new _module.C1();
        _nw402.__ctor();
        _2434_v12 = _nw402;
        let _index425 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2433_v11).length));
        (_2433_v11)[_index425] = _2434_v12;
        let _index426 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2433_v11).length));
        let _rhs417 = _2424_i0;
        let _rhs418 = _2434_v12;
        let _lhs328 = globalState;
        let _lhs329 = _2433_v11;
        let _lhs330 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2433_v11).length));
        _lhs328.f3 = _rhs417;
        _lhs329[_lhs330] = _rhs418;
        let _2435_v13;
        _2435_v13 = _dafny.Seq.UnicodeFromString("qajy");
        let _2436_v14;
        _2436_v14 = _dafny.Seq.of(new BigNumber(63));
        let _2437_v15;
        _2437_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_2435_v13).length), _2424_i0),(_2436_v14)[_module.__default.safeIndex(_2423_v3, new BigNumber((_2436_v14).length))]);
        _2437_v15 = (_2437_v15).update(_2423_v3, _2424_i0);
      }
      if ((((_2421_v1) ? (_2422_v2) : (_2422_v2))).IsProperSubsetOf(_2422_v2)) {
        (globalState).f8 = _2423_v3;
        if (_2421_v1) {
          let _2438_v16;
          let _init65 = ((_2439_v1) => function (_2440_i2) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_2439_v1,new _dafny.CodePoint('h'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2439_v1,new _dafny.CodePoint('b'.codePointAt(0))));
          })(_2421_v1);
          let _nw403 = Array((new BigNumber(21)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw403.length); _i0_65++) {
            _nw403[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2438_v16 = _nw403;
          _2438_v16 = _2438_v16;
          let _2441_v17;
          let _nw404 = new _module.C1();
          _nw404.__ctor();
          _2441_v17 = _nw404;
          let _2442_v18;
          _2442_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2441_v17,new BigNumber(761));
          (globalState).f3 = (_2423_v3).multipliedBy((((_2442_v18).contains(_2441_v17)) ? ((_2442_v18).get(_2441_v17)) : (_2423_v3)));
          let _rhs419 = _2421_v1;
          let _rhs420 = _2421_v1;
          let _lhs331 = globalState;
          r0 = _rhs419;
          _lhs331.f4 = _rhs420;
          let _2443_v19;
          let _nw405 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _2443_v19 = _nw405;
          let _2444_v20;
          let _nw406 = new _module.C4();
          _nw406.__ctor(_2443_v19);
          _2444_v20 = _nw406;
          let _2445_v21;
          _2445_v21 = _dafny.Set.fromElements(_2444_v20);
          let _2446_v22;
          _2446_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(898),_2445_v21);
          let _2447_v23;
          _2447_v23 = _dafny.Seq.of(_2445_v21);
          let _2448_v24;
          let _nw407 = Array((new BigNumber(22)).toNumber());
          _nw407[(_dafny.ZERO).toNumber()] = (_2445_v21).Difference(_dafny.Set.fromElements(_2444_v20, _2444_v20, _2444_v20));
          _nw407[(_dafny.ONE).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(2)).toNumber()] = (_2445_v21).Difference(_2445_v21);
          _nw407[(new BigNumber(3)).toNumber()] = (_2445_v21).Intersect(_2445_v21);
          _nw407[(new BigNumber(4)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(5)).toNumber()] = (_2445_v21).Union(_2445_v21);
          _nw407[(new BigNumber(6)).toNumber()] = (_2445_v21).Difference(_2445_v21);
          _nw407[(new BigNumber(7)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(8)).toNumber()] = (((_2446_v22).contains(_2423_v3)) ? ((_2446_v22).get(_2423_v3)) : (_2445_v21));
          _nw407[(new BigNumber(9)).toNumber()] = (_2445_v21).Difference(_dafny.Set.fromElements(_2444_v20, _2444_v20, _2444_v20));
          _nw407[(new BigNumber(10)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(11)).toNumber()] = (_2447_v23)[_module.__default.safeIndex(_2423_v3, new BigNumber((_2447_v23).length))];
          _nw407[(new BigNumber(12)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(13)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(14)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(15)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(_2444_v20);
          _nw407[(new BigNumber(17)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(18)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(19)).toNumber()] = _2445_v21;
          _nw407[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(_2444_v20);
          _nw407[(new BigNumber(21)).toNumber()] = (_2445_v21).Difference(_dafny.Set.fromElements(_2444_v20));
          _2448_v24 = _nw407;
          let _index427 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_2448_v24).length));
          (_2448_v24)[_index427] = _2445_v21;
          let _index428 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_2448_v24).length));
          (_2448_v24)[_index428] = _dafny.Set.fromElements(_2444_v20);
          let _2449_v25;
          _2449_v25 = new _dafny.CodePoint('h'.codePointAt(0));
          _2449_v25 = _2449_v25;
        } else {
          let _2450_v26;
          let _nw408 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2450_v26 = _nw408;
          let _2451_v27;
          _2451_v27 = _dafny.Seq.UnicodeFromString("sdej");
          let _2452_v28;
          _2452_v28 = new _dafny.CodePoint('w'.codePointAt(0));
          let _2453_v29;
          _2453_v29 = _module.D10.create_DC24(_2423_v3, _2421_v1, _2451_v27);
          let _2454_v30;
          let _nw409 = Array((new BigNumber(28)).toNumber());
          _nw409[(_dafny.ZERO).toNumber()] = _2451_v27;
          _nw409[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("uvvaguomr");
          _nw409[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("vv");
          _nw409[(new BigNumber(3)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_2455_i3) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), _module.__default.safeIndex(_2423_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_2456_i3) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })).length)), _2452_v28);
          _nw409[(new BigNumber(5)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(6)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(7)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("r");
          _nw409[(new BigNumber(9)).toNumber()] = _module.__default.fm25(globalState);
          _nw409[(new BigNumber(10)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(319)), ((_2457_v28) => function (_2458_i4) {
            return _2457_v28;
          })(_2452_v28));
          _nw409[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("gwsellrim");
          _nw409[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_2451_v27, _module.__default.safeIndex(_2423_v3, new BigNumber((_2451_v27).length)), _2452_v28);
          _nw409[(new BigNumber(14)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-182)), ((_2459_v28) => function (_2460_i5) {
            return _2459_v28;
          })(_2452_v28)), _module.__default.safeIndex(new BigNumber(352), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-182)), ((_2461_v28) => function (_2462_i5) {
            return _2461_v28;
          })(_2452_v28))).length)), _2452_v28);
          _nw409[(new BigNumber(16)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(17)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), ((_2463_v28) => function (_2464_i6) {
            return _2463_v28;
          })(_2452_v28));
          _nw409[(new BigNumber(19)).toNumber()] = (_2453_v29).dtor_cf33;
          _nw409[(new BigNumber(20)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("p");
          _nw409[(new BigNumber(22)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(23)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("fuxk");
          _nw409[(new BigNumber(25)).toNumber()] = _2451_v27;
          _nw409[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("lmxp");
          _nw409[(new BigNumber(27)).toNumber()] = _dafny.Seq.UnicodeFromString("gsdb");
          _2454_v30 = _nw409;
          let _index429 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2450_v26).length));
          (_2450_v26)[_index429] = _2454_v30;
          let _index430 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2450_v26).length));
          let _nw410 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          (_2450_v26)[_index430] = _nw410;
          let _2465_v31;
          _2465_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,true);
          let _2466_v32;
          _2466_v32 = _dafny.Seq.of(_2451_v27, _dafny.Seq.UnicodeFromString("tgmejkpo"), _2451_v27, _2451_v27, _2451_v27);
          let _2467_v33;
          _2467_v33 = _dafny.Set.fromElements(_2421_v1, _2421_v1);
          let _2468_v34;
          let _nw411 = Array((new BigNumber(26)).toNumber());
          _nw411[(_dafny.ZERO).toNumber()] = new BigNumber((_2422_v2).cardinality());
          _nw411[(_dafny.ONE).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2465_v31).length));
          _nw411[(new BigNumber(3)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2466_v32).length));
          _nw411[(new BigNumber(5)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(6)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(7)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(8)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_2423_v3);
          _nw411[(new BigNumber(10)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(11)).toNumber()] = new BigNumber((_2467_v33).length);
          _nw411[(new BigNumber(12)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_2469_v3) => function (_2470_i7) {
            return _2469_v3;
          })(_2423_v3))).length);
          _nw411[(new BigNumber(14)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(15)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(16)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(17)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(18)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(19)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(20)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(21)).toNumber()] = (_2420_v0).fm19(globalState);
          _nw411[(new BigNumber(22)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(23)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(24)).toNumber()] = _2423_v3;
          _nw411[(new BigNumber(25)).toNumber()] = new BigNumber((_2451_v27).length);
          _2468_v34 = _nw411;
          let _2471_v35;
          let _nw412 = Array((new BigNumber(10)).toNumber());
          _nw412[(_dafny.ZERO).toNumber()] = _2468_v34;
          _nw412[(_dafny.ONE).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(2)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(3)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(4)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(5)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(6)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(7)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(8)).toNumber()] = _2468_v34;
          _nw412[(new BigNumber(9)).toNumber()] = _2468_v34;
          _2471_v35 = _nw412;
          let _index431 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_2471_v35).length));
          (_2471_v35)[_index431] = _2468_v34;
          let _index432 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_2471_v35).length));
          let _rhs421 = _2468_v34;
          let _rhs422 = _2421_v1;
          let _rhs423 = (_2423_v3).isLessThan(_2423_v3);
          let _lhs332 = _2471_v35;
          let _lhs333 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_2471_v35).length));
          _lhs332[_lhs333] = _rhs421;
          r0 = _rhs422;
          _2421_v1 = _rhs423;
          let _2472_v36;
          let _nw413 = Array((new BigNumber(12)).toNumber()).fill(_module.D23.Default());
          _2472_v36 = _nw413;
          let _2473_v37;
          let _nw414 = Array((new BigNumber(4)).toNumber());
          _nw414[(_dafny.ZERO).toNumber()] = _2452_v28;
          _nw414[(_dafny.ONE).toNumber()] = _2452_v28;
          _nw414[(new BigNumber(2)).toNumber()] = _2452_v28;
          _nw414[(new BigNumber(3)).toNumber()] = _2452_v28;
          _2473_v37 = _nw414;
          let _2474_v38;
          _2474_v38 = _module.D23.create_DC45(_2473_v37);
          let _index433 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_2472_v36).length));
          (_2472_v36)[_index433] = _2474_v38;
          let _pat_let_tv73 = _2473_v37;
          let _index434 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_2472_v36).length));
          (_2472_v36)[_index434] = function (_pat_let63_0) {
            return function (_2475_dt__update__tmp_h0) {
              return function (_pat_let64_0) {
                return function (_2476_dt__update_hcf66_h0) {
                  return _module.D23.create_DC45(_2476_dt__update_hcf66_h0);
                }(_pat_let64_0);
              }(_pat_let_tv73);
            }(_pat_let63_0);
          }(_2474_v38);
          let _2477_v39;
          _2477_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,_2423_v3);
          let _2478_v40;
          _2478_v40 = _module.D2.create_DC7(_2452_v28, _2421_v1, new BigNumber(786));
          let _2479_v41;
          _2479_v41 = _dafny.Seq.of(_2467_v33);
          _2477_v39 = (_2477_v39).update((_2478_v40).dtor_cf11, new BigNumber(((_2467_v33).Intersect((_2479_v41)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("guqrkvlrc")).length), new BigNumber((_2479_v41).length))])).length));
          (globalState).f3 = (_dafny.ZERO).minus(_2423_v3);
        }
        let _2480_v42;
        let _init66 = ((_2481_v1) => function (_2482_i8) {
          return _2481_v1;
        })(_2421_v1);
        let _nw415 = Array((new BigNumber(12)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw415.length); _i0_66++) {
          _nw415[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2480_v42 = _nw415;
        _2480_v42 = _2480_v42;
        let _2483_v43;
        let _nw416 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _2483_v43 = _nw416;
        let _index435 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2483_v43).length));
        (_2483_v43)[_index435] = (_dafny.ZERO).minus(new BigNumber(-180));
        let _2484_v44;
        _2484_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,_2423_v3);
        let _2485_v45;
        _2485_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2484_v44,_2423_v3);
        let _index436 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2483_v43).length));
        (_2483_v43)[_index436] = (((_2485_v45).contains(_2484_v44)) ? ((_2485_v45).get(_2484_v44)) : (_2423_v3));
        let _2486_v46;
        _2486_v46 = new _dafny.CodePoint('d'.codePointAt(0));
        let _2487_v47;
        _2487_v47 = _dafny.Seq.of(_2486_v46, _2486_v46, _2486_v46);
        let _2488_v48;
        _2488_v48 = _dafny.Seq.of(_2486_v46, _2486_v46, _2486_v46, (_2487_v47)[_module.__default.safeIndex(_2423_v3, new BigNumber((_2487_v47).length))]);
        _2487_v47 = _2487_v47;
      } else {
        if (false) {
          let _2489_v49;
          let _init67 = ((_2490_v3) => function (_2491_i9) {
            return _module.__default.safeModuloInt(_2491_i9, _2490_v3);
          })(_2423_v3);
          let _nw417 = Array((new BigNumber(21)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw417.length); _i0_67++) {
            _nw417[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2489_v49 = _nw417;
          let _index437 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length));
          (_2489_v49)[_index437] = _2423_v3;
          let _index438 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length));
          (_2489_v49)[_index438] = _2423_v3;
          let _2492_v50;
          let _init68 = ((_2493_v1) => function (_2494_i10) {
            return _2493_v1;
          })(_2421_v1);
          let _nw418 = Array((new BigNumber(10)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw418.length); _i0_68++) {
            _nw418[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2492_v50 = _nw418;
          let _index439 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_2492_v50).length));
          (_2492_v50)[_index439] = _2421_v1;
          let _index440 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_2492_v50).length));
          (_2492_v50)[_index440] = _2421_v1;
          let _2495_v51;
          _2495_v51 = _dafny.Map.Empty.slice().updateUnsafe((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))],_2421_v1);
          let _2496_v52;
          _2496_v52 = _dafny.Seq.of((((_2495_v51).contains(_2423_v3)) ? ((_2495_v51).get(_2423_v3)) : (_2421_v1)));
          let _2497_v53;
          _2497_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v1,(_2420_v0).fm19(globalState));
          let _2498_v54;
          _2498_v54 = _dafny.MultiSet.fromElements((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))], _2423_v3, (_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))], new BigNumber((_2497_v53).length), new BigNumber((_dafny.Seq.UnicodeFromString("youkw")).length));
          let _2499_v55;
          _2499_v55 = _dafny.MultiSet.fromElements((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))], new BigNumber((_2498_v54).cardinality()), _2423_v3);
          let _2500_v56;
          _2500_v56 = _dafny.MultiSet.fromElements(_2499_v55, _2498_v54);
          let _2501_v57;
          _2501_v57 = _dafny.Map.Empty.slice().updateUnsafe((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))],_2500_v56);
          let _2502_v58;
          _2502_v58 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(455)), function (_2503_i11) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }));
          let _2504_v59;
          _2504_v59 = _dafny.Seq.of(new BigNumber(((((_2501_v57).contains((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))])) ? ((_2501_v57).get((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))])) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(_2498_v54, _2498_v54))))).cardinality()), (_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))], (new BigNumber((_2502_v58).length)).plus((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))]));
          let _index441 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_2492_v50).length));
          let _index442 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_2492_v50).length));
          let _rhs424 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-617)));
          let _rhs425 = _2421_v1;
          let _rhs426 = (_2496_v52)[_module.__default.safeIndex(_2423_v3, new BigNumber((_2496_v52).length))];
          let _rhs427 = _2421_v1;
          let _rhs428 = (_2504_v59)[_module.__default.safeIndex((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))], new BigNumber((_2504_v59).length))];
          let _lhs334 = _2492_v50;
          let _lhs335 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_2492_v50).length));
          let _lhs336 = _2492_v50;
          let _lhs337 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_2492_v50).length));
          let _lhs338 = globalState;
          _2423_v3 = _rhs424;
          _lhs334[_lhs335] = _rhs425;
          _lhs336[_lhs337] = _rhs426;
          _2421_v1 = _rhs427;
          _lhs338.f3 = _rhs428;
          let _2505_v60;
          _2505_v60 = new _dafny.CodePoint('n'.codePointAt(0));
          let _2506_v61;
          _2506_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v1,_2505_v60);
          let _2507_v62;
          _2507_v62 = _dafny.Set.fromElements(_2423_v3, new BigNumber((_2506_v61).length));
          let _2508_v63;
          _2508_v63 = _dafny.Set.fromElements(new BigNumber((_2507_v62).length), (_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))]);
          let _2509_v64;
          _2509_v64 = _dafny.Set.fromElements(_2507_v62, (_2508_v63).Union(_dafny.Set.fromElements(new BigNumber((_2496_v52).length))), _2507_v62);
          let _2510_v65;
          let _init69 = ((_2511_v51) => function (_2512_i12) {
            return _2511_v51;
          })(_2495_v51);
          let _nw419 = Array((new BigNumber(20)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw419.length); _i0_69++) {
            _nw419[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _2510_v65 = _nw419;
          let _index443 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2510_v65).length));
          (_2510_v65)[_index443] = function () {
            let _coll112 = new _dafny.Map();
            for (const _compr_112 of _dafny.IntegerRange(new BigNumber(137), new BigNumber(239))) {
              let _2513_v66 = _compr_112;
              if (((new BigNumber(137)).isLessThanOrEqualTo(_2513_v66)) && ((_2513_v66).isLessThan(new BigNumber(239)))) {
                _coll112.push([(_2513_v66).plus(new BigNumber(520)),_2421_v1]);
              }
            }
            return _coll112;
          }();
          let _index444 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2510_v65).length));
          let _rhs429 = _2509_v64;
          let _rhs430 = (_2495_v51).Merge(_2495_v51);
          let _rhs431 = (_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))];
          let _lhs339 = _2510_v65;
          let _lhs340 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2510_v65).length));
          let _lhs341 = globalState;
          _2509_v64 = _rhs429;
          _lhs339[_lhs340] = _rhs430;
          _lhs341.f8 = _rhs431;
          let _2514_v67;
          _2514_v67 = _dafny.Seq.UnicodeFromString("catkgef");
          let _2515_v68;
          _2515_v68 = _dafny.Map.Empty.slice().updateUnsafe((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))],_2514_v67);
          let _index445 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length));
          (_2489_v49)[_index445] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ds"), (((_2515_v68).contains((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))])) ? ((_2515_v68).get((_2489_v49)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2489_v49).length))])) : (_2514_v67)))).length);
          let _2516_v69;
          _2516_v69 = _module.D2.create_DC5(_2492_v50);
          let _2517_v70;
          let _nw420 = Array((new BigNumber(13)).toNumber());
          _nw420[(_dafny.ZERO).toNumber()] = _2492_v50;
          _nw420[(_dafny.ONE).toNumber()] = (((_2492_v50)[_module.__default.safeIndex(new BigNumber(256), new BigNumber((_2492_v50).length))]) ? (_2492_v50) : (_2492_v50));
          _nw420[(new BigNumber(2)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(3)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(4)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(5)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(6)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(7)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(8)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(9)).toNumber()] = (_2516_v69).dtor_cf5;
          _nw420[(new BigNumber(10)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(11)).toNumber()] = _2492_v50;
          _nw420[(new BigNumber(12)).toNumber()] = _2492_v50;
          _2517_v70 = _nw420;
          let _index446 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_2517_v70).length));
          (_2517_v70)[_index446] = _2492_v50;
          let _index447 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_2517_v70).length));
          (_2517_v70)[_index447] = _2492_v50;
        } else {
          let _2518_v71;
          _2518_v71 = new _dafny.CodePoint('t'.codePointAt(0));
          let _2519_v72;
          _2519_v72 = _dafny.Seq.UnicodeFromString("gi");
          let _2520_v73;
          let _nw421 = Array((new BigNumber(12)).toNumber());
          _nw421[(_dafny.ZERO).toNumber()] = _2518_v71;
          _nw421[(_dafny.ONE).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(2)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(3)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(4)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(5)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(6)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw421[(new BigNumber(8)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(9)).toNumber()] = (_2519_v72)[_module.__default.safeIndex(new BigNumber(650), new BigNumber((_2519_v72).length))];
          _nw421[(new BigNumber(10)).toNumber()] = _2518_v71;
          _nw421[(new BigNumber(11)).toNumber()] = ((_2421_v1) ? (_2518_v71) : (_2518_v71));
          _2520_v73 = _nw421;
          let _index448 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2520_v73).length));
          (_2520_v73)[_index448] = new _dafny.CodePoint('x'.codePointAt(0));
          let _2521_v74;
          _2521_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,(_2519_v72)[_module.__default.safeIndex(_2423_v3, new BigNumber((_2519_v72).length))]);
          let _index449 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2520_v73).length));
          (_2520_v73)[_index449] = (((_2521_v74).contains(_2423_v3)) ? ((_2521_v74).get(_2423_v3)) : (_2518_v71));
          let _2522_v75;
          let _init70 = ((_2523_v3, _2524_v2) => function (_2525_i13) {
            return _module.__default.safeDivisionInt(_2525_i13, new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2523_v3), new BigNumber(966), new BigNumber((_2524_v2).cardinality()), _2523_v3)).cardinality()));
          })(_2423_v3, _2422_v2);
          let _nw422 = Array((new BigNumber(26)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw422.length); _i0_70++) {
            _nw422[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2522_v75 = _nw422;
          let _index450 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2522_v75).length));
          (_2522_v75)[_index450] = _2423_v3;
          let _index451 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2522_v75).length));
          (_2522_v75)[_index451] = _module.__default.safeModuloInt(_module.__default.fm6(_2423_v3, globalState), (_dafny.ZERO).minus(new BigNumber((_2519_v72).length)));
          let _2526_v76;
          let _nw423 = new _module.C3();
          _nw423.__ctor();
          _2526_v76 = _nw423;
          let _2527_v77;
          _2527_v77 = _2526_v76;
          let _2528_v78;
          _2528_v78 = _dafny.Seq.of(_2527_v77, _2526_v76, _2527_v77, _2527_v77);
          _2528_v78 = _2528_v78;
          (globalState).f3 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), ((_2529_v73) => function (_2530_i14) {
            return (_2529_v73)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2529_v73).length))];
          })(_2520_v73))).length)).minus(_2423_v3);
          (globalState).f4 = _2421_v1;
        }
        (globalState).f3 = (_2423_v3).multipliedBy((_2423_v3).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2423_v3,_2421_v1)).length)));
        let _2531_v79;
        _2531_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(_2421_v1, globalState),_2421_v1);
        (globalState).f4 = !(_2531_v79).contains(_2421_v1);
        let _2532_v80;
        let _nw424 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
        _2532_v80 = _nw424;
        let _2533_v81;
        _2533_v81 = _dafny.Seq.of(_2421_v1);
        let _2534_v82;
        _2534_v82 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_2533_v81).length)), _2423_v3);
        let _2535_v83;
        _2535_v83 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v1,_2534_v82);
        let _index452 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2532_v80).length));
        (_2532_v80)[_index452] = (_2535_v83).update(_2421_v1, _2534_v82);
        let _index453 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_2532_v80).length));
        (_2532_v80)[_index453] = _dafny.Map.Empty.slice().updateUnsafe((_2421_v1) && (_2421_v1),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-944)), ((_2536_v3) => function (_2537_i15) {
          return _2536_v3;
        })(_2423_v3)));
        let _2538_v84;
        _2538_v84 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2539_v85;
        _2539_v85 = _dafny.Seq.UnicodeFromString("pudx");
        r1 = _dafny.Seq.contains(_2539_v85, _2538_v84);
      }
      let _2540_v86;
      _2540_v86 = new _dafny.CodePoint('b'.codePointAt(0));
      let _2541_v87;
      _2541_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-713),_2540_v86);
      let _2542_v88;
      _2542_v88 = _module.D19.create_DC38(_2541_v87);
      let _2543_v89;
      _2543_v89 = _module.D19.create_DC40(_2542_v88);
      let _2544_v90;
      _2544_v90 = _module.D19.create_DC40(_2542_v88);
      let _source40 = _2544_v90;
      if (_source40.is_DC39) {
        let _2545___mcc_h0 = (_source40).cf58;
        let _2546___mcc_h1 = (_source40).cf59;
        let _2547_cf59 = _2546___mcc_h1;
        let _2548_cf58 = _2545___mcc_h0;
        let _2549_v91;
        _2549_v91 = _dafny.Seq.UnicodeFromString("imuchiwj");
        let _2550_v92;
        let _nw425 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _2550_v92 = _nw425;
        let _2551_v93;
        _2551_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,_2550_v92);
        let _2552_v94;
        _2552_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v3,_2423_v3);
        _2423_v3 = (new BigNumber((_dafny.Seq.update(_2549_v91, _module.__default.safeIndex(new BigNumber((_2551_v93).length), new BigNumber((_2549_v91).length)), _2540_v86)).length)).multipliedBy(((((_2552_v94).contains(_module.__default.fm20(_2548_cf58, globalState))) ? ((_2552_v94).get(_module.__default.fm20(_2548_cf58, globalState))) : (_2423_v3))).minus(new BigNumber(890)));
        r0 = _2548_cf58;
        (globalState).f4 = _2421_v1;
        _2550_v92 = _2550_v92;
      } else if (_source40.is_DC38) {
        let _2553___mcc_h2 = (_source40).cf57;
        let _2554_cf57 = _2553___mcc_h2;
        let _2555_v95;
        _2555_v95 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2423_v3));
        let _2556_v96;
        _2556_v96 = _2555_v95;
        _2556_v96 = _2556_v96;
        let _2557_v97;
        let _nw426 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _2557_v97 = _nw426;
        let _2558_v98;
        _2558_v98 = _dafny.Map.Empty.slice().updateUnsafe(_2557_v97,_2423_v3);
        let _2559_v99;
        _2559_v99 = _module.D21.create_DC42(_2558_v98);
        let _2560_v100;
        _2560_v100 = _dafny.Set.fromElements(_2421_v1);
        let _2561_v101;
        _2561_v101 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v1,false);
        let _2562_v102;
        _2562_v102 = _module.D0.create_DC0(new BigNumber((_2561_v101).length));
        let _rhs432 = (((_2421_v1) ? (_2555_v95) : (_2555_v95))).Difference((_2555_v95).Difference(_dafny.MultiSet.fromElements((((_2422_v2).contains(_2421_v1)) ? ((_2422_v2).get(_2421_v1)) : (_2423_v3)), _2423_v3, new BigNumber((_2560_v100).length), _2423_v3, new BigNumber(309))));
        let _rhs433 = _2559_v99;
        let _rhs434 = (_2562_v102).dtor_cf0;
        let _lhs342 = globalState;
        let _lhs343 = globalState;
        _lhs342.f2 = _rhs432;
        _2559_v99 = _rhs433;
        _lhs343.f3 = _rhs434;
        let _2563_v103;
        _2563_v103 = _dafny.Seq.UnicodeFromString("jo");
        let _2564_v104;
        _2564_v104 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_2563_v103).length)).isEqualTo(_2423_v3),_2423_v3);
        (globalState).f3 = new BigNumber((_2564_v104).length);
        (globalState).f8 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_2423_v3, new BigNumber((_2561_v101).length)), new BigNumber(((_2422_v2).update(_2421_v1, _module.__default.abs(_2423_v3))).cardinality())));
      } else {
        let _2565___mcc_h3 = (_source40).cf60;
        let _2566_cf60 = _2565___mcc_h3;
        let _2567_v105;
        _2567_v105 = _dafny.Seq.UnicodeFromString("fgp");
        let _rhs435 = _2421_v1;
        let _rhs436 = _2421_v1;
        let _rhs437 = _dafny.Seq.IsProperPrefixOf(_2567_v105, _dafny.Seq.Concat(_dafny.Seq.update(_2567_v105, _module.__default.safeIndex(_2423_v3, new BigNumber((_2567_v105).length)), _2540_v86), _dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_2568_v86) => function (_2569_i16) {
          return _2568_v86;
        })(_2540_v86))));
        let _rhs438 = _2421_v1;
        let _lhs344 = globalState;
        let _lhs345 = globalState;
        _2421_v1 = _rhs435;
        _lhs344.f4 = _rhs436;
        r0 = _rhs437;
        _lhs345.f4 = _rhs438;
        let _2570_v106;
        let _nw427 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _2570_v106 = _nw427;
        _2570_v106 = _2570_v106;
        let _2571_v107;
        _2571_v107 = _dafny.Set.fromElements(_2540_v86);
        let _2572_v108;
        let _nw428 = new _module.C3();
        _nw428.__ctor();
        _2572_v108 = _nw428;
        let _2573_v109;
        _2573_v109 = _dafny.Map.Empty.slice().updateUnsafe((_2571_v107).Intersect(_2571_v107),_2572_v108);
        _2573_v109 = (_2573_v109).update(_2571_v107, _2572_v108);
        (globalState).f4 = _2421_v1;
      }
      let _2574_v110;
      _2574_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v1,_2540_v86);
      (globalState).f4 = !((_2574_v110).Merge(_2574_v110)).contains(!(_2421_v1));
      r0 = _2421_v1;
      r1 = _2421_v1;
      return [r0, r1];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2575_v0;
      _2575_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _2576_v1;
      _2576_v1 = _dafny.MultiSet.fromElements((((_2575_v0).contains(p0)) ? ((_2575_v0).get(p0)) : (p1)), p1, p1);
      let _source41 = _module.__default.fm40((_2576_v1).IsSubsetOf(_dafny.MultiSet.fromElements(p1)), globalState);
      if (_source41.is_DC6) {
        let _2577___mcc_h0 = (_source41).cf6;
        let _2578___mcc_h1 = (_source41).cf7;
        let _2579___mcc_h2 = (_source41).cf8;
        let _2580_cf8 = _2579___mcc_h2;
        let _2581_cf7 = _2578___mcc_h1;
        let _2582_cf6 = _2577___mcc_h0;
        let _2583_v2;
        let _nw429 = new _module.C3();
        _nw429.__ctor();
        _2583_v2 = _nw429;
        let _2584_v3;
        _2584_v3 = _dafny.MultiSet.fromElements(_2583_v2);
        let _2585_v4;
        _2585_v4 = _dafny.Set.fromElements((_2582_cf6).minus(p0), (_2582_cf6).plus(new BigNumber((_2584_v3).cardinality())), p0, _module.__default.safeDivisionInt(new BigNumber(293), _2582_cf6), _2582_cf6);
        let _2586_v5;
        let _nw430 = Array((new BigNumber(23)).toNumber()).fill(false);
        _2586_v5 = _nw430;
        let _index454 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length));
        (_2586_v5)[_index454] = !(p1);
        let _2587_v6;
        let _nw431 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2587_v6 = _nw431;
        let _2588_v7;
        _2588_v7 = _dafny.Seq.UnicodeFromString("wbp");
        let _index455 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length));
        (_2587_v6)[_index455] = _2588_v7;
        let _index456 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length));
        let _index457 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length));
        let _rhs439 = (_2585_v4).Difference(_2585_v4);
        let _rhs440 = !((p0).minus(p0)).isEqualTo((_dafny.ZERO).minus((p0).minus(_2582_cf6)));
        let _rhs441 = _dafny.Seq.Concat(_2588_v7, _2588_v7);
        let _rhs442 = _2580_cf8;
        let _lhs346 = _2586_v5;
        let _lhs347 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length));
        let _lhs348 = _2587_v6;
        let _lhs349 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length));
        let _lhs350 = globalState;
        _2585_v4 = _rhs439;
        _lhs346[_lhs347] = _rhs440;
        _lhs348[_lhs349] = _rhs441;
        _lhs350.f4 = _rhs442;
        let _2589_v8;
        _2589_v8 = new _dafny.CodePoint('y'.codePointAt(0));
        let _2590_v9;
        _2590_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.contains((_2587_v6)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length))], _2589_v8));
        _2590_v9 = _module.__default.fm21(!((_2586_v5)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length))]), globalState);
        let _index458 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length));
        (_2586_v5)[_index458] = !((_2586_v5)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_2586_v5).length))]);
        let _2591_v11;
        _2591_v11 = _dafny.Seq.of(new BigNumber((function () {
          let _coll113 = new _dafny.Map();
          for (const _compr_113 of _dafny.IntegerRange(new BigNumber(-461), new BigNumber(451))) {
            let _2592_v10 = _compr_113;
            if (((new BigNumber(-461)).isLessThanOrEqualTo(_2592_v10)) && ((_2592_v10).isLessThan(new BigNumber(451)))) {
              _coll113.push([(_2592_v10).minus(p0),_dafny.Seq.update((_2587_v6)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length))], _module.__default.safeIndex(p0, new BigNumber(((_2587_v6)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_2587_v6).length))]).length)), _2589_v8)]);
            }
          }
          return _coll113;
        }()).length));
        (globalState).f0 = _dafny.Seq.update(_2591_v11, _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2591_v11).length)), p0);
      } else if (_source41.is_DC7) {
        let _2593___mcc_h3 = (_source41).cf9;
        let _2594___mcc_h4 = (_source41).cf10;
        let _2595___mcc_h5 = (_source41).cf11;
        let _2596_cf11 = _2595___mcc_h5;
        let _2597_cf10 = _2594___mcc_h4;
        let _2598_cf9 = _2593___mcc_h3;
        let _2599_v12;
        let _nw432 = new _module.C5();
        _nw432.__ctor();
        _2599_v12 = _nw432;
        let _2600_v13;
        let _nw433 = new _module.C5();
        _nw433.__ctor();
        _2600_v13 = _nw433;
        (globalState).f8 = (_dafny.ZERO).minus(new BigNumber(-857));
        if (!(_2576_v1).equals(_2576_v1)) {
          let _2601_v14;
          let _init71 = ((_2602_cf11) => function (_2603_i0) {
            return _module.__default.safeDivisionInt(_2603_i0, _2602_cf11);
          })(_2596_cf11);
          let _nw434 = Array((new BigNumber(21)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw434.length); _i0_71++) {
            _nw434[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2601_v14 = _nw434;
          _2601_v14 = _2601_v14;
          let _2604_v15;
          let _nw435 = Array((new BigNumber(11)).toNumber());
          _2604_v15 = _nw435;
          let _2605_v16;
          let _nw436 = new _module.C1();
          _nw436.__ctor();
          _2605_v16 = _nw436;
          let _index459 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_2604_v15).length));
          (_2604_v15)[_index459] = _2605_v16;
          let _2606_v17;
          _2606_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2601_v14);
          let _2607_v18;
          let _nw437 = new _module.C3();
          _nw437.__ctor();
          _2607_v18 = _nw437;
          let _2608_v19;
          _2608_v19 = _dafny.Seq.of(_2606_v17, ((_2597_cf10) ? (_2606_v17) : (_2606_v17)), _2606_v17);
          let _2609_v20;
          _2609_v20 = _dafny.Seq.of(_2601_v14, (((_2606_v17).contains(p1)) ? ((_2606_v17).get(p1)) : (_2601_v14)));
          let _index460 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_2604_v15).length));
          let _rhs443 = _2605_v16;
          let _rhs444 = _2596_cf11;
          let _rhs445 = (_2608_v19)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, _2596_cf11), new BigNumber((_2608_v19).length))];
          let _rhs446 = _2607_v18;
          let _rhs447 = !_dafny.Seq.contains(_2609_v20, _2601_v14);
          let _lhs351 = _2604_v15;
          let _lhs352 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_2604_v15).length));
          let _lhs353 = globalState;
          _lhs351[_lhs352] = _rhs443;
          _lhs353.f3 = _rhs444;
          _2606_v17 = _rhs445;
          _2607_v18 = _rhs446;
          _2597_cf10 = _rhs447;
          let _2610_v21;
          _2610_v21 = _dafny.Map.Empty.slice().updateUnsafe((_2597_cf10) || (p1),(_dafny.ZERO).minus(p0));
          _2610_v21 = (_2610_v21).update(_2597_cf10, p0);
          let _index461 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_2601_v14).length));
          (_2601_v14)[_index461] = (_dafny.ZERO).minus(p0);
          let _index462 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_2601_v14).length));
          (_2601_v14)[_index462] = _module.__default.fm6(p0, globalState);
          r0 = new BigNumber(321);
        } else {
          let _2611_v22;
          _2611_v22 = _dafny.Seq.UnicodeFromString("k");
          _2611_v22 = _2611_v22;
          (globalState).f3 = p0;
          (globalState).f8 = (_module.__default.fm20(_2597_cf10, globalState)).plus((new BigNumber(212)).minus(_2596_cf11));
          (globalState).f4 = _2597_cf10;
          let _2612_v23;
          _2612_v23 = _module.D2.create_DC6(_2596_cf11, false, p1);
          _2612_v23 = _2612_v23;
        }
      } else {
        let _2613___mcc_h6 = (_source41).cf5;
        let _2614_cf5 = _2613___mcc_h6;
        let _2615_v24;
        let _nw438 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
        _2615_v24 = _nw438;
        let _2616_v25;
        _2616_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2615_v24);
        let _2617_v26;
        let _nw439 = new _module.C4();
        _nw439.__ctor((((_2616_v25).contains(p1)) ? ((_2616_v25).get(p1)) : (_2615_v24)));
        _2617_v26 = _nw439;
        let _2618_v27;
        let _init72 = ((_2619_p0) => function (_2620_i1) {
          return _module.__default.safeModuloInt(_2620_i1, _2619_p0);
        })(p0);
        let _nw440 = Array((new BigNumber(26)).toNumber());
        for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw440.length); _i0_72++) {
          _nw440[_i0_72] = _init72(new BigNumber(_i0_72));
        }
        _2618_v27 = _nw440;
        let _index463 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_2618_v27).length));
        (_2618_v27)[_index463] = (p0).minus(p0);
        let _2621_v28;
        let _nw441 = Array((new BigNumber(11)).toNumber());
        _nw441[(_dafny.ZERO).toNumber()] = _2618_v27;
        _nw441[(_dafny.ONE).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(2)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(3)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(4)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(5)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(6)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(7)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(8)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(9)).toNumber()] = _2618_v27;
        _nw441[(new BigNumber(10)).toNumber()] = _2618_v27;
        _2621_v28 = _nw441;
        let _index464 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_2621_v28).length));
        (_2621_v28)[_index464] = _2618_v27;
        let _index465 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_2618_v27).length));
        let _index466 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_2621_v28).length));
        let _rhs448 = p0;
        let _rhs449 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), new BigNumber(-463));
        let _rhs450 = (p0).minus(p0);
        let _rhs451 = _2618_v27;
        let _lhs354 = _2618_v27;
        let _lhs355 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_2618_v27).length));
        let _lhs356 = globalState;
        let _lhs357 = globalState;
        let _lhs358 = _2621_v28;
        let _lhs359 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_2621_v28).length));
        _lhs354[_lhs355] = _rhs448;
        _lhs356.f3 = _rhs449;
        _lhs357.f8 = _rhs450;
        _lhs358[_lhs359] = _rhs451;
        (globalState).f4 = p1;
        let _2622_v29;
        let _2623_v30;
        let _out29;
        let _out30;
        let _outcollector7 = (_2617_v26).m0(globalState);
        _out29 = _outcollector7[0];
        _out30 = _outcollector7[1];
        _2622_v29 = _out29;
        _2623_v30 = _out30;
      }
      let _hi16 = (p0).plus(_module.__default.fm6(p0, globalState));
      for (let _2624_i2 = p0; _2624_i2.isLessThan(_hi16); _2624_i2 = _2624_i2.plus(_dafny.ONE)) {
        let _2625_v31;
        let _nw442 = Array((new BigNumber(3)).toNumber()).fill([]);
        _2625_v31 = _nw442;
        let _nw443 = Array((_dafny.ONE).toNumber()).fill([]);
        _2625_v31 = _nw443;
        let _2626_v32;
        let _init73 = function (_2627_i3) {
          return true;
        };
        let _nw444 = Array((new BigNumber(14)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw444.length); _i0_73++) {
          _nw444[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2626_v32 = _nw444;
        _2626_v32 = ((!(p1)) ? (_2626_v32) : (_2626_v32));
        if (p1) {
          (globalState).f4 = p1;
          (globalState).f3 = p0;
          let _2628_v33;
          let _init74 = ((_2629_i2) => function (_2630_i4) {
            return _module.__default.safeDivisionInt(_2630_i4, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(745)), ((_2631_i2) => function (_2632_i5) {
              return _2631_i2;
            })(_2629_i2)))).cardinality()));
          })(_2624_i2);
          let _nw445 = Array((new BigNumber(26)).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw445.length); _i0_74++) {
            _nw445[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2628_v33 = _nw445;
          let _2633_v34;
          _2633_v34 = _dafny.Seq.UnicodeFromString("cqbcks");
          let _2634_v35;
          _2634_v35 = _dafny.Seq.of(_2633_v34);
          let _2635_v36;
          _2635_v36 = _2628_v33;
          let _rhs452 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_2633_v34), _2634_v35);
          let _rhs453 = (_2635_v36);
          let _lhs360 = globalState;
          _lhs360.f4 = _rhs452;
          _2628_v33 = _rhs453;
          let _2636_v37;
          let _nw446 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2636_v37 = _nw446;
          _2636_v37 = _2636_v37;
          (globalState).f4 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(586)), ((_2637_v34, _2638_p0) => function (_2639_i6) {
            return (_2637_v34)[_module.__default.safeIndex(_2638_p0, new BigNumber((_2637_v34).length))];
          })(_2633_v34, p0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(738)), function (_2640_i7) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }));
        } else {
          let _2641_v38;
          _2641_v38 = new _dafny.CodePoint('l'.codePointAt(0));
          let _2642_v39;
          _2642_v39 = _dafny.Seq.of(_2641_v38, _2641_v38);
          let _2643_v40;
          _2643_v40 = _module.D10.create_DC24(p0, p1, _2642_v39);
          let _2644_v41;
          _2644_v41 = _dafny.Seq.of(_2642_v39, _dafny.Seq.update((_2643_v40).dtor_cf33, _module.__default.safeIndex(_2624_i2, new BigNumber(((_2643_v40).dtor_cf33).length)), _2641_v38));
          let _2645_v42;
          _2645_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm20(p1, globalState),_dafny.Seq.Concat(_dafny.Seq.of(_2641_v38, _2641_v38, _2641_v38), (_2644_v41)[_module.__default.safeIndex(_2624_i2, new BigNumber((_2644_v41).length))]));
          _2645_v42 = (_2645_v42).update((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p0)), _2642_v39);
          (globalState).f4 = p1;
          _2641_v38 = ((p1) ? (_2641_v38) : (_2641_v38));
          let _2646_v43;
          _2646_v43 = _module.D25.create_DC52(_2625_v31);
          let _2647_v44;
          let _nw447 = Array((new BigNumber(27)).toNumber());
          _nw447[(_dafny.ZERO).toNumber()] = (_2646_v43).dtor_cf73;
          _nw447[(_dafny.ONE).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(2)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(3)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(4)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(5)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(6)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(7)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(8)).toNumber()] = ((p1) ? (_2625_v31) : (_2625_v31));
          _nw447[(new BigNumber(9)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(10)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(11)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(12)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(13)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(14)).toNumber()] = ((p1) ? (_2625_v31) : (_2625_v31));
          _nw447[(new BigNumber(15)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(16)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(17)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(18)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(19)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(20)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(21)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(22)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(23)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(24)).toNumber()] = (_2646_v43).dtor_cf73;
          _nw447[(new BigNumber(25)).toNumber()] = _2625_v31;
          _nw447[(new BigNumber(26)).toNumber()] = _2625_v31;
          _2647_v44 = _nw447;
          let _index467 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_2647_v44).length));
          (_2647_v44)[_index467] = _2625_v31;
          let _index468 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_2647_v44).length));
          (_2647_v44)[_index468] = _2625_v31;
          (globalState).f8 = (p0).plus(_2624_i2);
        }
        let _2648_v45;
        _2648_v45 = _dafny.Set.fromElements(p1);
        let _index469 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2626_v32).length));
        (_2626_v32)[_index469] = (_2648_v45).IsDisjointFrom(_dafny.Set.fromElements(p1, p1, false));
        let _index470 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2626_v32).length));
        (_2626_v32)[_index470] = p1;
      }
      r0 = _module.__default.safeDivisionInt(new BigNumber(357), (new BigNumber(663)).multipliedBy(p0));
      let _2649_v46;
      _2649_v46 = new _dafny.CodePoint('h'.codePointAt(0));
      let _2650_v47;
      _2650_v47 = _dafny.Map.Empty.slice().updateUnsafe((p0).plus(p0),_2649_v46);
      _2650_v47 = (_2650_v47).update(p0, ((p1) ? (new _dafny.CodePoint('a'.codePointAt(0))) : (_2649_v46)));
      let _2651_v48;
      let _nw448 = Array((new BigNumber(26)).toNumber());
      _2651_v48 = _nw448;
      let _2652_v49;
      _2652_v49 = _dafny.Seq.UnicodeFromString("msd");
      let _2653_v50;
      _2653_v50 = _dafny.Seq.of(_2652_v49);
      let _2654_v51;
      _2654_v51 = _module.D4.create_DC11(p1, p1, p0, _2652_v49);
      let _2655_v52;
      _2655_v52 = _module.D4.create_DC12(_2654_v51);
      let _2656_v53;
      _2656_v53 = _dafny.Seq.of(_2655_v52, _module.D4.create_DC12(_2654_v51));
      let _2657_v54;
      _2657_v54 = _dafny.Map.Empty.slice().updateUnsafe(_2651_v48,(_2653_v50)[_module.__default.safeIndex(new BigNumber((_2656_v53).length), new BigNumber((_2653_v50).length))]);
      _2657_v54 = (_2657_v54).Merge((_2657_v54).Merge(_2657_v54));
      (globalState).f4 = _module.__default.fm49(p1, globalState);
      let _2658_v55;
      _2658_v55 = _dafny.Set.fromElements(new BigNumber(167), p0);
      let _2659_v56;
      _2659_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(p1, globalState),new BigNumber((_2658_v55).length));
      let _2660_v57;
      _2660_v57 = _dafny.MultiSet.fromElements(_2659_v56, (_dafny.Map.Empty.slice().updateUnsafe(p1,p0)).update(p1, p0));
      let _2661_v58;
      _2661_v58 = _dafny.Seq.of(new BigNumber((_2660_v57).cardinality()));
      r0 = (p0).plus((_2661_v58)[_module.__default.safeIndex(p0, new BigNumber((_2661_v58).length))]);
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
